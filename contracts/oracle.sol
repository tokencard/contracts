pragma solidity 0.4.24;

import "internal/control.sol";
import "external/strings.sol";
import "external/safe-math.sol";
import "external/oraclize-api.sol";

contract Proof {
    using Strings for *;

    function fromJSON(string _json, string _label) internal returns (string) {
        Strings.slice memory slice = _json.toSlice();
        slice.split(":".toSlice()).toString();
        return slice.until("}".toSlice()).toString();
    }

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

    // Date and time
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

    function isLeapYear(uint16 year) internal pure returns (bool) {
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

    function monthNumber(string _month) internal pure returns (uint8) {
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
            revert("invalid month name");
        }
    }
}

contract Oracle is UsingOraclize, Proof, Control {
    using Strings for *;
    using SafeMath for uint256;

    event TokenAddition(address _token, string _label, uint8 _decimals);
    event TokenOmission(address _token, string _label, uint8 _decimals);
    event TokenRemoval(address _token);
    event TokenRateUpdate(address _token, uint _rate);

    event OraclizeQuerySuccess(string _label);
    event OraclizeQueryFailure(string _reason);

    event Conversion(address _token, uint _rate, uint _result);
    event ProofVerification(bytes _publicKey, string _result, bool _success);

    struct Token {
        string label;   // Token symbol
        uint8 decimals; // Number of decimal places
        bool exists;    // Token addition flag
        uint rate;      // Token exchange rate in wei
    }

    mapping(address => Token) public tokens;
    address[] private _tokenAddresses;

    bytes public cryptoComparePublicKey = hex"a0f4f688350018ad1b9785991c0bde5f704b005dc79972b114dbed4a615a983710bfc647ebe5a320daa28771dce6a2d104f5efa2e4a85ba3760b76d46f8571ca";
    mapping(bytes32 => address) private _queryToAddress;
    uint private _callbackTimestamp;

    /// @dev Check if the parameter's length is within a valid range.
    modifier hasValidLength(address[] _addresses) {
        require(_addresses.length >= 1 && _addresses.length <= 20, "invalid parameter length");
        _;
    }

    /// @dev Construct the oracle with multiple controllers, address resolver and custom gas price.
    constructor(address[] _controllers, address _resolver, uint _gasPrice) public {
        OAR = OraclizeAddrResolverI(_resolver);
        oraclize_setCustomGasPrice(_gasPrice);
        oraclize_setProof(proofType_Native);
        for (uint i = 0; i < _controllers.length; i++) {
            _addController(_controllers[i]);
        }
    }

    /// @dev Convert ERC20 token amount to the corresponding ether amount (used by the wallet contract).
    /// @param _token ERC20 token contract address.
    /// @param _amount amount of token in base units.
    function convert(address _token, uint _amount) external view returns (uint) {
        // Store the token in memory to save map entry lookup gas.
        Token memory token = tokens[_token];
        // Require that the token exists and that its rate is not zero.
        require(token.exists && token.decimals != 0);
        // Check for overflow (TODO: add safe exponentiation) and return the ether amount.
        return _amount.mul(token.rate).div(10 ** token.decimals);
    }

    /// @dev Add ERC20 tokens to the list of supported tokens.
    /// @param _tokens ERC20 token contract addresses.
    /// @param _labels ERC20 token names.
    /// @param _decimals number of decimal places used by each ERC20 token.
    function addTokens(address[] _tokens, bytes32[] _labels, uint8[] _decimals) external onlyController hasValidLength(_tokens) {
        // Require that all parameters have the same length.
        require(_tokens.length == _labels.length && _tokens.length == _decimals.length, "parameter lengths do not match");
        // Add each token to the list of supported tokens.
        for (uint i = 0; i < _tokens.length; i++) {
            // Store the intermediate values to save gas on array element lookup.
            string memory label = _labels[i].toSliceB32().toString();
            address token = _tokens[i];
            uint8 decimals = _decimals[i];
            // Add the token if it doesn't exist, otherwise emit omission event.
            if (!tokens[token].exists) {
                // Add the token to the token list.
                tokens[token] = Token({
                    label : label,
                    decimals : decimals,
                    rate : 0,
                    exists : true
                    });
                // Add the token address to the address list.
                _tokenAddresses.push(token);
                // Emit token addition event.
                emit TokenAddition(token, label, decimals);
            } else {
                // Emit token omission event.
                emit TokenOmission(token, label, decimals);
            }
        }
    }

    /// @dev Remove ERC20 tokens from the list of supported tokens.
    /// @param _tokens ERC20 token contract addresses.
    function removeTokens(address[] _tokens) external onlyController hasValidLength(_tokens) {
        // Delete each token object from the list of supported tokens based on the addresses provided.
        for (uint i = 0; i < _tokens.length; i++) {
            // Store the token address.
            address token = _tokens[i];
            // Delete the token object.
            delete tokens[token];
            // Remove the token address from the address list.
            for (uint ii = 0; ii < _tokenAddresses.length.sub(1); ii++) {
                if (_tokenAddresses[ii] == token) {
                    _tokenAddresses[ii] = _tokenAddresses[_tokenAddresses.length.sub(1)];
                    break;
                }
            }
            _tokenAddresses.length--;
            // Emit token removal event.
            emit TokenRemoval(token);
        }
    }

    /// @dev Update ERC20 token exchange rate manually.
    /// @param _token ERC20 token contract address.
    /// @param _rate ERC20 token exchange rate in wei.
    function updateTokenRate(address _token, uint _rate) external onlyController {
        // Require that the token exists.
        require(tokens[_token].exists);
        // Update the token's rate.
        tokens[_token].rate = _rate;
        // Emit the rate update event.
        emit TokenRateUpdate(_token, _rate);
    }

    /// @dev Update ERC20 token exchange rates for all supported tokens.
    function updateTokenRates() external payable onlyController {
        _updateTokenRates();
    }

    /// @dev Re-usable helper function that performs the Oraclize Query.
    function _updateTokenRates() private {
        // Check if the contract has enough Ether to pay for the query.
        if (oraclize_getPrice("URL") * _tokenAddresses.length > address(this).balance) {
            // Emit the query failure event.
            emit OraclizeQueryFailure("not enough balance to pay for the query");
        } else {
            // Set up the crypto compare API query strings.
            Strings.slice memory apiPrefix = "https://min-api.cryptocompare.com/data/price?fsym=".toSlice();
            Strings.slice memory apiSuffix = "&tsyms=ETH&sign=true".toSlice();
            // Create a new oraclize query for each supported token.
            for (uint i = 0; i < _tokenAddresses.length; i++) {
                // Store the token label used in the query.
                Strings.slice memory label = tokens[_tokenAddresses[i]].label.toSlice();
                // Create a new oraclize query from the component strings.
                bytes32 queryID = oraclize_query(5, "URL", apiPrefix.concat(label).toSlice().concat(apiSuffix), 2000000);
                // Store the query ID together with the associated token address.
                _queryToAddress[queryID] = _tokenAddresses[i];
                // Emit the query success event.
                emit OraclizeQuerySuccess(label.toString());
            }
        }
    }

    /// @dev Handle Oraclize query callback and verifiy the provided origin proof.
    /// @param _queryID Oraclize query ID.
    /// @param _result query result in JSON format.
    /// @param _proof origin proof from crypto compare.
    function __callback(bytes32 _queryID, string _result, bytes _proof) public {
        // Require that the caller is the Oraclize contract.
        require(msg.sender == oraclize_cbAddress(), "sender is not oraclize");
        // Require that the proof is valid.
        require(verifyProof(_result, _proof, cryptoComparePublicKey), "provided origin proof is invalid");
        // Use the query ID to find the matching token.
        address _address = _queryToAddress[_queryID];
        Token memory token = tokens[_address];
        // Parse the JSON result to get the rate in wei.
        token.rate = parseInt(fromJSON(_result, token.label), 18);
        // Emit the rate update event.
        emit TokenRateUpdate(_address, token.rate);
        // Remove query from the list.
        delete _queryToAddress[_queryID];
    }

    /// @dev Verify the origin proof returned by the crypto compare API.
    /// @param _result query result in JSON format.
    /// @param _proof origin proof from crypto compare.
    /// @param _publicKey crypto compare public key.
    function verifyProof(string _result, bytes _proof, bytes _publicKey) private returns (bool) {
        // Extract the signature.
        uint signatureLength = uint(_proof[1]);
        bytes memory signature = new bytes(signatureLength);
        signature = copyBytes(_proof, 2, signatureLength, signature, 0);

        // Extract the headers.
        uint headersLength = uint(_proof[2 + signatureLength]) * 256 + uint(_proof[2 + signatureLength + 1]);
        bytes memory headers = new bytes(headersLength);
        headers = copyBytes(_proof, 4 + signatureLength, headersLength, headers, 0);

        // Check if the date is valid.
        bytes memory dateHeader = new bytes(30);
        dateHeader = copyBytes(headers, 5, 30, dateHeader, 0);
        bool dateOK = verifyDate(string(dateHeader));
        if (!dateOK) return false;

        // Check if the signed digest hash matches the result hash.
        bytes memory digest = new bytes(headersLength - 52);
        digest = copyBytes(headers, 52, headersLength - 52, digest, 0);
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

    /// @dev Verify the signed HTTP date header.
    /// @param _date extracted date string e.g. Wed, 12 Sep 2018 15:18:14 GMT.
    function verifyDate(string _date) private view returns (bool) {
        Strings.slice memory date = _date.toSlice();
        Strings.slice memory timeDelimiter = ":".toSlice();
        Strings.slice memory dateDelimiter = " ".toSlice();
        // Split the date string.
        date.split(",".toSlice());
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
        return timestamp > _callbackTimestamp;
    }
}



