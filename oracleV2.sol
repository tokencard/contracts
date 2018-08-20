pragma solidity ^0.4.18;

//"0xfe209bde5ca32fa20e6728a005f26d651fff5982","TKN",8

import "github.com/oraclize/ethereum-api/oraclizeAPI.sol";
import "github.com/Arachnid/solidity-stringutils/strings.sol";

contract Oracle is usingOraclize{
    
    using strings for *;

    /*
     *  Events
     */

    event TokenRemoval(address indexed tokenID);
    event LogNewOraclizeQuery(string description);
    event RateUpdated(address tokenAddress,string rate);
    
    struct Erc20Token {
        string label;
        uint8 decimals; //maximum number of decimals
        uint rate; // ratio: value over ETH
        bool supported;

    }
    
    mapping (address => Erc20Token) public tokens;
    address[] public contractAddresses;
    
    address public controller;
    
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
    
    modifier onlyOraclize {
        require(msg.sender == oraclize_cbAddress());
        _;
    }
    
    constructor() public {
        controller = msg.sender;
        OAR = OraclizeAddrResolverI(0x6f485c8bf6fc43ea212e93bbf8ce046c7f1cb475);
        // oraclize_setProof(proofType_Android);
    }
    
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
    }
    
    function addTokenBatch (address[] tokenIDs, string labels, uint8[] decimals) 
        public
        onlyController
    {
        require(tokenIDs.length == decimals.length);
        
        strings.slice  memory labelSlice = labels.toSlice();
        strings.slice memory delim = ".".toSlice();
        
        uint numTokenLabels = labelSlice.count(delim) + 1;
        require(numTokenLabels == decimals.length);
    
        for (uint i = 0; i < numTokenLabels; i++) {
            if(!tokens[tokenIDs[i]].supported){
                contractAddresses.push(tokenIDs[i]);
                tokens[tokenIDs[i]].label = labelSlice.split(delim).toString();
                tokens[tokenIDs[i]].decimals = decimals[i];
                tokens[tokenIDs[i]].rate = 0;
                tokens[tokenIDs[i]].supported = true;
            }
            
        }
    }
    
    function removeToken(address tokenID)
        public
        onlyController
        tokenSupported(tokenID)
    {
        delete tokens[tokenID].supported;
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
       
        uint contractAddressesLength = contractAddresses.length;
        
        if (oraclize_getPrice("URL") * contractAddressesLength > address(this).balance){
            emit LogNewOraclizeQuery("Oraclize query was NOT sent, please add some ETH to cover for the query fee");
        }
        else{
            
            strings.slice  memory apiPrefix = "json(https://min-api.cryptocompare.com/data/price?fsym=".toSlice();
            strings.slice  memory apiSuffix = "&tsyms=ETH&sign=true).ETH".toSlice();
            
            for (uint i=0; i<contractAddressesLength; i++){
                strings.slice memory tokenLabel = tokens[contractAddresses[i]].label.toSlice();
                string memory apiString = apiPrefix.concat(tokenLabel).toSlice().concat(apiSuffix); //assigned for clarity
                bytes32 queryId = oraclize_query("URL",apiString);
                validIDs[queryId] = contractAddresses[i];
                emit LogNewOraclizeQuery("Oraclize query was sent, standing by for the answer...");
            }
        }
    }
    
    function __callback(bytes32 queryId, string result, bytes proof) public onlyOraclize{
        
        require(tokens[validIDs[queryId]].supported);
        
        /* if ((proof.length > 0) && (nativeProof_verify(result, proof, cc_pubkey))) {
          TKNETH = parseInt(result, 6);
          delete validIDs[queryId];
          last_update_timestamp = now;
        } */

        tokens[validIDs[queryId]].rate = parseInt(result, 18);
        
        delete validIDs[queryId];
        emit RateUpdated(validIDs[queryId],result);
    }
    
    function mockTransaction(address tknLabel, string amount) public view returns (uint){
        
        uint amountInt =  parseInt(amount, tokens[tknLabel].decimals)*tokens[tknLabel].rate;
        
        return amountInt/tokens[tknLabel].decimals;
    }
    
    
}
