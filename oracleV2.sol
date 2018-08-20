pragma solidity ^0.4.18;

import "github.com/oraclize/ethereum-api/oraclizeAPI.sol";
import "github.com/Arachnid/solidity-stringutils/strings.sol";

contract Oracle is usingOraclize{

    using strings for *;

    /*
     *  Events
     */

    event TokenRemoval(address indexed tokenID);
    event TokenAddition(address indexed tokenID, string label);
    event LogNewOraclizeQuery(string description);
    event RateUpdated(address tokenAddress,string rate);

    struct Erc20Token {
        string label;  //token symbol
        uint8 decimals; //maximum number of decimals
        uint rate; // ratio: value over ETH
        bool supported; //denotes whether the token is inserted in the array

    }

    mapping (address => Erc20Token) public tokens;
    address[] public contractAddresses; //keep track of the existing tokens, a way to enumerating them

    address public controller;

    /// @dev unique id returned from Oraclize, mapped to a token address so we can undertand in the callback which rate to update.
    mapping (bytes32 => address) validIDs;

    modifier tokenSupported(address tokenID){
        require(tokens[tokenID].supported);
        _;
    }

    modifier tokenNotSupported(address tokenID){
        require(!tokens[tokenID].supported);
        _;
    }

    /// @dev Executable only by the controller.
    modifier onlyController() {
        require(msg.sender == controller);
        _;
    }

    /// @dev Executable only by Oraclize (used in the __callback function).
    modifier onlyOraclize {
        require(msg.sender == oraclize_cbAddress());
        _;
    }

    constructor() public {
        controller = msg.sender;
        OAR = OraclizeAddrResolverI(0x6f485C8BF6fc43eA212E93BBF8ce046C7f1cb475);
        // oraclize_setProof(proofType_Android);
    }

    /**
    * @dev add a new token to the list and mapping
    * @param tokenID token contract addresses
    * @param label the symbol/abbreviation used to represent the token (a '.' seperated string)
    * @param decimals the precision of the token value(maximum number of decimal points)
    */
    function addToken(address tokenID, string label, uint8 decimals)
        public
        onlyController
        tokenNotSupported(tokenID)
        {
            contractAddresses.push(tokenID);
            tokens[tokenID] = Erc20Token({
                label : label,
                decimals : decimals,
                rate : 0,
                supported: true
            });

        emit TokenAddition(tokenID,label);
    }

    /**
    * @dev add new tokens to the list and mapping
    * @param tokenIDs token contract addresses
    * @param labels the symbol/abbreviation used to represent the token (a '.' seperated string)
    * @param decimals the precision of the token value(maximum number of decimal points)
    */
    function addTokenBatch (address[] tokenIDs, string labels, uint8[] decimals)
        public
        onlyController
    {
        require(tokenIDs.length == decimals.length);

        //convert strings into the library's 'slice' format
        strings.slice  memory labelSlice = labels.toSlice();
        strings.slice memory delim = ".".toSlice();

        uint numTokenLabels = labelSlice.count(delim) + 1; //the number of labels is +1 of thenumber of '.' ["t1.t2.t3"] string expected
        require(numTokenLabels == decimals.length);

        for (uint i = 0; i < numTokenLabels; i++) {
            if(!tokens[tokenIDs[i]].supported){
                contractAddresses.push(tokenIDs[i]); //push token to the array
                tokens[tokenIDs[i]].label = labelSlice.split(delim).toString();//split the string with a '.' delimiter
                tokens[tokenIDs[i]].decimals = decimals[i];
                tokens[tokenIDs[i]].rate = 0; //to be updated later
                tokens[tokenIDs[i]].supported = true;
            }

        }
    }

     /**
    * @dev remove a token from the list of supported ones
    * @param tokenID token contract addresses
    */
    function removeToken(address tokenID)
        public
        onlyController
        tokenSupported(tokenID)
    {
        delete tokens[tokenID].supported;

        //check if the address matches up to one token before the last one
        //the tokenSupported() modifier ensures that the token address actually exists.
        //if no match is found in the loop, it means that the last address was the desired one, simply reduce the size by one in any case.
        uint contractAddressesLength = contractAddresses.length - 1;
        for (uint i=0; i<contractAddressesLength; i++)
            if (contractAddresses[i] == tokenID) {
                contractAddresses[i] = contractAddresses[contractAddressesLength];
                break;
            }
        contractAddresses.length--;

        emit TokenRemoval(tokenID);
    }

    function updateRates() public payable{

        uint contractAddressesLength = contractAddresses.length; //number of supported tokens

        if (oraclize_getPrice("URL") * contractAddressesLength > address(this).balance){
            emit LogNewOraclizeQuery("Oraclize query was NOT sent, please add some ETH to cover for the query fee");
        }
        else{

            //the fixed strings required to access the Cryptocompare api.
            strings.slice  memory apiPrefix = "json(https://min-api.cryptocompare.com/data/price?fsym=".toSlice();
            strings.slice  memory apiSuffix = "&tsyms=ETH&sign=true).ETH".toSlice();

            for (uint i=0; i<contractAddressesLength; i++){
                strings.slice memory tokenLabel = tokens[contractAddresses[i]].label.toSlice();//the token label to be inserted in the api.
                string memory apiString = apiPrefix.concat(tokenLabel).toSlice().concat(apiSuffix); //assigned for clarity
                bytes32 queryId = oraclize_query("URL",apiString);
                validIDs[queryId] = contractAddresses[i];// map queryId to token contract address, to be used in the callback.

                emit LogNewOraclizeQuery("Oraclize query was sent, standing by for the answer...");
            }
        }
    }

    function __callback(bytes32 queryId, string result, bytes proof) public onlyOraclize{

        require(tokens[validIDs[queryId]].supported);//must be a valid token.

        /* if ((proof.length > 0) && (nativeProof_verify(result, proof, cc_pubkey))) {
          TKNETH = parseInt(result, 6);
          delete validIDs[queryId];
          last_update_timestamp = now;
        } */

        tokens[validIDs[queryId]].rate = parseInt(result, 18); //transform rate(string) to uint (wei precision)
        delete validIDs[queryId]; //remove mapping

        emit RateUpdated(validIDs[queryId],result);
    }

    function mockTransaction(address tknLabel, string amount) public view returns (uint){

        uint amountInt =  parseInt(amount, tokens[tknLabel].decimals)*tokens[tknLabel].rate;

        return amountInt/tokens[tknLabel].decimals;
    }


}
