pragma solidity 0.4.24;

import "./internal/control.sol";
import "./external/strings.sol";
import "./external/oraclize-api.sol";

contract Date {
    uint constant DAY_IN_SECONDS = 86400;
    uint constant YEAR_IN_SECONDS = 31536000;
    uint constant LEAP_YEAR_IN_SECONDS = 31622400;
    uint constant HOUR_IN_SECONDS = 3600;
    uint constant MINUTE_IN_SECONDS = 60;
    uint16 constant ORIGIN_YEAR = 1970;

    bytes32 constant private JANUARY = keccak256("Jan");
    bytes32 constant private FEBRUARY = keccak256("Feb");
    bytes32 constant private MARCH = keccak256("Mar");
    bytes32 constant private APRIL = keccak256("Apr");
    bytes32 constant private MAY = keccak256("May");
    bytes32 constant private JUNE = keccak256("Jun");
    bytes32 constant private JULY = keccak256("Jul");
    bytes32 constant private AUGUST = keccak256("Aug");
    bytes32 constant private SEPTEMBER = keccak256("Sep");
    bytes32 constant private OCTOBER = keccak256("Oct");
    bytes32 constant private NOVEMBER = keccak256("Nov");
    bytes32 constant private DECEMBER = keccak256("Dec");

    function verifyDate(string _date) internal view returns (bool) {
        strings.slice memory date = _date.toSlice();
        strings.slice memory timeDelimiter = ":".toSlice();
        strings.slice memory dateDelimiter = " ".toSlice();

        // Split the week day.
        date.split(",".toSlice());
        // Split beyond the first space
        date.split(dateDelimiter);

        // Get individual date components.
        uint8 day = uint8(parseInt(date.split(dateDelimiter).toString(), 0));
        uint8 month = monthNumber(date.split(dateDelimiter).toString());
        uint16 year = uint16(parseInt(date.split(dateDelimiter).toString(), 0));
        uint8 hour = uint8(parseInt(date.split(timeDelimiter).toString(), 0));
        uint8 minute = uint8(parseInt(date.split(timeDelimiter).toString(), 0));
        uint8 second = uint8(parseInt(date.split(timeDelimiter).toString(), 0));

        // Convert date components to a unix timestamp.
        uint timestamp;
        // Year
        for (uint16 i = ORIGIN_YEAR; i < year; i++) {
            if (isLeapYear(i)) {
                timestamp += LEAP_YEAR_IN_SECONDS;
            } else {
                timestamp += YEAR_IN_SECONDS;
            }
        }
        // Month
        uint8[12] memory monthDayCounts;

        monthDayCounts[0] = 31;

        if (isLeapYear(year)) {
            monthDayCounts[1] = 29;
        } else {
            monthDayCounts[1] = 28;
        }
        monthDayCounts[2] = 31;
        monthDayCounts[3] = 30;
        monthDayCounts[4] = 31;
        monthDayCounts[5] = 30;
        monthDayCounts[6] = 31;
        monthDayCounts[7] = 31;
        monthDayCounts[8] = 30;
        monthDayCounts[9] = 31;
        monthDayCounts[10] = 30;
        monthDayCounts[11] = 31;

        for (i = 1; i < month; i++) {
            timestamp += DAY_IN_SECONDS * monthDayCounts[i - 1];
        }
        // Day
        timestamp += DAY_IN_SECONDS * (day - 1);

        // Hour
        timestamp += HOUR_IN_SECONDS * (hour);

        // Minute
        timestamp += MINUTE_IN_SECONDS * (minute);

        // Second
        timestamp += second;
        __timestamp = timestamp;
        return timestamp > lastCallbackTime;
    }

    function isLeapYear(uint16 year) private pure returns (bool) {
        if (year % 4 != 0) {
            return false;
        }
        if (year % 100 != 0) {
            return true;
        }
        if (year % 400 != 0) {
            return false;
        }
        return true;
    }

    function monthNumber(string _month) private pure returns (uint8) {
        bytes32 month = keccak256(_month);
        if (month == JANUARY) {
            return 1;
        } else if (month == FEBRUARY) {
            return 2;
        } else if (month == MARCH) {
            return 3;
        } else if (month == APRIL) {
            return 4;
        } else if (month == MAY) {
            return 5;
        } else if (month == JUNE) {
            return 6;
        } else if (month == JULY) {
            return 7;
        } else if (month == AUGUST) {
            return 8;
        } else if (month == SEPTEMBER) {
            return 9;
        } else if (month == OCTOBER) {
            return 10;
        } else if (month == NOVEMBER) {
            return 11;
        } else if (month == DECEMBER) {
            return 12;
        } else {
            revert();
        }
    }
}

contract Base64 {
    bytes constant BASE64_DECODE_CHAR = hex"000000000000000000000000000000000000000000000000000000000000000000000000000000000000003e003e003f3435363738393a3b3c3d00000000000000000102030405060708090a0b0c0d0e0f10111213141516171819000000003f001a1b1c1d1e1f202122232425262728292a2b2c2d2e2f30313233";

    function base64decode(bytes _encoded) internal pure returns (bytes) {
        byte v1;
        byte v2;
        byte v3;
        byte v4;

        uint length = _encoded.length;
        bytes memory result = new bytes(length);

        uint index;
        if (keccak256(_encoded[length - 2]) == keccak256('=')) {
            length -= 2;
        } else if (keccak256(_encoded[length - 1]) == keccak256('=')) {
            length -= 1;
        }
        uint count = length >> 2 << 2;
        for (uint i = 0; i < count;) {
            v1 = BASE64_DECODE_CHAR[uint(_encoded[i++])];
            v2 = BASE64_DECODE_CHAR[uint(_encoded[i++])];
            v3 = BASE64_DECODE_CHAR[uint(_encoded[i++])];
            v4 = BASE64_DECODE_CHAR[uint(_encoded[i++])];

            result[index++] = (v1 << 2 | v2 >> 4) & 255;
            result[index++] = (v2 << 4 | v3 >> 2) & 255;
            result[index++] = (v3 << 6 | v4) & 255;
        }
        if (length - count == 2) {
            v1 = BASE64_DECODE_CHAR[uint(_encoded[i++])];
            v2 = BASE64_DECODE_CHAR[uint(_encoded[i++])];
            result[index++] = (v1 << 2 | v2 >> 4) & 255;
        }
        else if (length - count == 3) {
            v1 = BASE64_DECODE_CHAR[uint(_encoded[i++])];
            v2 = BASE64_DECODE_CHAR[uint(_encoded[i++])];
            v3 = BASE64_DECODE_CHAR[uint(_encoded[i++])];

            result[index++] = (v1 << 2 | v2 >> 4) & 255;
            result[index++] = (v2 << 4 | v3 >> 2) & 255;
        }
        // Set to correct length.
        assembly {
            mstore(result, index)
        }
        return result;
    }
}


contract Oracle is UsingOraclize, Date, Base64, Control {

    using strings for *;

    event TokenRemoval(address indexed token);
    event TokenAddition(address indexed token);
    event ProofVerification(bytes publicKey, string result, bool verified);
    event TokenAlreadySupported(address indexed token, string label);
    event RateUpdated(address indexed token, uint rate);
    event OraclizeSucceeded(string label);
    event OraclizeFailedNotEnoughEth(uint balance);

    struct Erc20Token {
        string label;   // Token symbol
        uint8 decimals; // Maximum number of decimals
        uint rate;      // Rate of token in wei
        bool supported; // Denotes whether the token is inserted in the array
    }

    mapping (address => Erc20Token) public tokens;
    address[] private _contractAddresses; // keep track of the existing tokens, a way to enumerating them

    address public controller;
    uint public lastCallbackTime;
    bytes public cryptoComparePublicKey;

    /// @dev unique id returned from Oraclize, mapped to a token address so we can understand in the callback which rate to update.
    mapping (bytes32 => address) private _queryIDToToken;

    modifier tokenSupported(address _token) {
        require(tokens[_token].supported);
        _;
    }

    modifier tokenNotSupported(address _token) {
        require(!tokens[_token].supported);
        _;
    }

    /// @dev Executable only by the controller.
    modifier onlyController() {
        require(msg.sender == controller);
        _;
    }

    /// @dev Executable only by Oraclize (used in the __callback function).
    modifier onlyOraclize() {
        require(msg.sender == oraclize_cbAddress());
        _;
    }

    constructor(address _oraclizeResolver, uint _gasPrice) public {
        controller = msg.sender;
        cryptoComparePublicKey = hex"a0f4f688350018ad1b9785991c0bde5f704b005dc79972b114dbed4a615a983710bfc647ebe5a320daa28771dce6a2d104f5efa2e4a85ba3760b76d46f8571ca";
        OAR = OraclizeAddrResolverI(_oraclizeResolver);
        oraclize_setCustomGasPrice(_gasPrice);
        oraclize_setProof(proofType_Native);
    }

    /// @dev Getter for contract addresses array.
    function contractAddresses() public view returns (address[]) {
        return _contractAddresses;
    }


    /// @dev add a new token to the list and mapping
    /// @param _token token contract addresses
    /// @param _label the symbol/abbreviation used to represent the token (a '.' separated string)
    /// @param _decimals the precision of the token value(maximum number of decimal points)
    function addToken(address _token, string _label, uint8 _decimals) public onlyController tokenNotSupported(_token) {
        _contractAddresses.push(_token);
        tokens[_token] = Erc20Token({
            label : _label,
            decimals : _decimals,
            rate : 0,
            supported: true
        });
        emit TokenAddition(_token);
    }


    /// @dev add new tokens to the list and mapping
    /// @param _tokens token contract addresses
    /// @param _labels the symbol/abbreviation used to represent the token (a '.' separated string)
    /// @param _decimals the precision of the token value(maximum number of decimal points)
    function addTokenBatch (address[] _tokens, string _labels, uint8[] _decimals) public onlyController {
        require(_tokens.length == _decimals.length);

        // Convert strings into the library's 'slice' format.
        strings.slice memory labelSlice = _labels.toSlice();
        strings.slice memory delimiter = ".".toSlice();

        uint numTokenLabels = labelSlice.count(delimiter) + 1; // the number of labels is +1 of the number of '.' ["t1.t2.t3"] string expected
        require(numTokenLabels == _decimals.length);

        for (uint i = 0; i < numTokenLabels; i++) {
            string memory tempLabel = labelSlice.split(delimiter).toString();//split the string with a '.' delimiter
            address t = _tokens[i];
            uint8 d = _decimals[i];
            if (!tokens[t].supported) {
                addToken(t, tempLabel, d);
            } else {
                emit TokenAlreadySupported(t, tempLabel);
            }
        }
    }


    /// @dev remove a token from the list of supported ones
    /// @param _token token contract addresses
    function removeToken(address _token) public onlyController tokenSupported(_token) {
        delete tokens[_token].supported;

        // Check if the address matches up to one token before the last one
        // the tokenSupported() modifier ensures that the token address actually exists.
        // If no match is found in the loop, it means that the last address was the desired one, simply reduce the size by one in any case.
        uint contractAddressesLength = _contractAddresses.length - 1;
        for (uint i=0; i<contractAddressesLength; i++)
            if (_contractAddresses[i] == _token) {
                _contractAddresses[i] = _contractAddresses[contractAddressesLength];
                break;
            }
        _contractAddresses.length--;

        emit TokenRemoval(_token);
    }

    function convert(address _token, uint amount) external view tokenSupported(_token) returns (uint) {
        require(tokens[_token].rate != 0);
        assert((amount * tokens[_token].rate) / tokens[_token].rate == amount); // Overflow check, returns 0
        return amount*tokens[_token].rate/(uint(10)**tokens[_token].decimals);
    }

    function updateRateManual(address _token, uint rate) external onlyController tokenSupported(_token) {
        tokens[_token].rate = rate;

        emit RateUpdated(_token, rate);
    }

    function getContractAddressesLength() private view returns (uint) {
        return _contractAddresses.length;
    }

    function updateRates() public payable {

        uint contractAddressesLength = _contractAddresses.length; //number of supported tokens
        uint currentBalance = address(this).balance;

        // Check if the address has enough ETH to pay for the oraclize_getPrice method
        if (oraclize_getPrice("URL") * contractAddressesLength > address(this).balance) {
            emit OraclizeFailedNotEnoughEth(currentBalance);
        } else {
            // The fixed strings required to access the Cryptocompare api.
            strings.slice memory apiPrefix = "https://min-api.cryptocompare.com/data/price?fsym=".toSlice();
            strings.slice memory apiSuffix = "&tsyms=ETH&sign=true".toSlice();

            for (uint i = 0; i < contractAddressesLength; i++) {
                strings.slice memory tokenLabel = tokens[_contractAddresses[i]].label.toSlice(); //the token label to be inserted in the api.
                string memory apiString = apiPrefix.concat(tokenLabel).toSlice().concat(apiSuffix); //assigned for clarity
                bytes32 queryID = oraclize_query(5, "URL", apiString, 2000000);
                _queryIDToToken[queryID] = _contractAddresses[i];// map queryId to token contract address, to be used in the callback.

                emit OraclizeSucceeded(tokenLabel.toString());
            }
        }
    }

    bytes32 public __id;
    string public __result;
    bytes public __proof;
    bool public __verified;
    uint public __timestamp;

    function __callback(bytes32 _queryID, string _result, bytes _proof) public onlyOraclize {

        __id = _queryID;
        __result = _result;
        __proof = _proof;
        __verified = verifyProof(_result, _proof, cryptoComparePublicKey);



        //address tokenAddress = _queryIDToToken[queryID];
        //Erc20Token token = tokens[tokenAddress];

        //require(token.supported); // must be a valid token.




        //uint rate = parseInt(result, token.decimals);

        //token.rate = rate; // transform rate(string) to uint (wei precision)

        //emit RateUpdated(tokenAddress, rate);
        //delete tokenAddress; // remove mapping

    }

    function verifyProof(string _result, bytes _proof, bytes _publicKey) private returns (bool) {
        // Extract the signature.
        uint signatureLength = uint(_proof[1]);
        bytes memory signature = new bytes(signatureLength);
        signature = copyBytes(_proof, 2, signatureLength, signature, 0);

        // Extract the headers.
        uint headersLength = uint(_proof[2+signatureLength])*256 + uint(_proof[2+signatureLength+1]);
        bytes memory headers = new bytes(headersLength);
        headers = copyBytes(_proof, 4+signatureLength, headersLength, headers, 0);

        // Check if the date is valid.
        bytes memory dateHeader = new bytes(30);
        dateHeader = copyBytes(headers, 5, 30, dateHeader, 0);
        bool dateOK = verifyDate(string(dateHeader));
        if (!dateOK) return false;

        // Check if the signed digest hash matches the derived one.
        bytes memory digest = new bytes(headersLength-52);
        digest = copyBytes(headers, 52, headersLength-52, digest, 0);
        bool digestOK = keccak256(sha256(_result)) == keccak256(base64decode(digest));
        if (!digestOK) return false;

        // Check if the signature is valid and if the signer addresses match.
        address signer;
        bool signatureOK;
        (signatureOK, signer) = ecrecovery(sha256(headers), signature);

        if (signatureOK && signer == address(keccak256(_publicKey))) {
            emit ProofVerification(_publicKey, _result, true);
            return true;
        } else {
            emit ProofVerification(_publicKey, _result, false);
            return false;
        }
    }
}


