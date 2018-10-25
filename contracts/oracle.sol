pragma solidity ^0.4.25;

import "./internal/controllable.sol";
import "./external/strings.sol";
import "./external/safe-math.sol";
import "./external/oraclize-api.sol";


/// @title Oracle converts ERC20 token amounts into equivalent ether amounts based on cryptocurrency exchange rates.
interface IOracle {
    function convert(address, uint) external view returns (uint);
}


/// @title JSON provides JSON parsing functionality.
contract JSON {
    using Strings for *;

    // @dev Extracts JSON rate value from the response object.
    // @param _json body of the JSON response from the CryptoCompare API.
    // @param _symbol asset symbol used to extract the correct json field.
    function parseRate(string _json, string _symbol) internal pure returns (string) {
        Strings.slice memory body = _json.toSlice();
        body.beyond("{".toSlice());
        body.until("}".toSlice());
        Strings.slice memory quoteMark = "\"".toSlice();
        body.find(quoteMark.concat(_symbol.toSlice()).toSlice().concat(quoteMark).toSlice());
        Strings.slice memory asset;
        body.split(",".toSlice(), asset);
        asset.split(":".toSlice());
        return asset.toString();
    }
}


/// @title Date provides date parsing functionality.
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

    // @return the number of the month based on its name.
    // @param _month the first three letters of a month's name e.g. "Jan".
    function monthToNumber(string _month) internal pure returns (uint8) {
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
            revert("not a valid month");
        }
    }
}


/// @title Base64 provides base 64 decoding functionality.
contract Base64 {
    bytes constant BASE64_DECODE_CHAR = hex"000000000000000000000000000000000000000000000000000000000000000000000000000000000000003e003e003f3435363738393a3b3c3d00000000000000000102030405060708090a0b0c0d0e0f10111213141516171819000000003f001a1b1c1d1e1f202122232425262728292a2b2c2d2e2f30313233";

    // @return decoded array of bytes.
    // @param _encoded base 64 encoded array of bytes.
    function base64decode(bytes _encoded) internal pure returns (bytes) {
        byte v1;
        byte v2;
        byte v3;
        byte v4;
        uint length = _encoded.length;
        bytes memory result = new bytes(length);
        uint index;
        if (keccak256(_encoded[length - 2]) == keccak256("=")) {
            length -= 2;
        } else if (keccak256(_encoded[length - 1]) == keccak256("=")) {
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
        } else if (length - count == 3) {
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


/// @title Oracle provides asset exchange rates and conversion functionality.
contract Oracle is UsingOraclize, Base64, Date, JSON, Controllable, IOracle {
    using Strings for *;
    using SafeMath for uint256;

    //todo mischa why don't we have our controller address here so we know who ?
    event AddedToken(address _token, string _symbol, uint _magnitude);
    event RemovedToken(address _token);
    event UpdatedTokenRate(address _token, uint _rate);

    event SetGasPrice(uint _gasPrice);
    //todo mischa both of these guys as well 
    event Converted(address _token, uint _amount, uint _ether);

    event RequestedUpdate(string _symbol);
    event FailedUpdateRequest(string _reason);

    event VerifiedProof(bytes _publicKey, string _result);
    event FailedProofVerification(bytes _publicKey, string _result, string _reason);

    event SetCryptoComparePublicKey(bytes _publicKey);

    struct Token {
        string symbol;    // Token symbol
        uint magnitude;   // 10^decimals
        uint rate;        // Token exchange rate in wei
        uint lastUpdate;  // Time of the last rate update
        bool exists;      // Flags if the struct is empty or not
    }

    mapping(address => Token) public tokens;
    address[] private _tokenAddresses;

    bytes public APIPublicKey;
    mapping(bytes32 => address) private _queryToToken;

    // @dev Construct the oracle with multiple controllers, address resolver and custom gas price.
    // @dev _resolver is the oraclize address resolver contract address.
    // @param _ens is the address of the ENS.
    // @param _controllerName is the ENS name of the Controller.
    constructor(address _resolver, address _ens, bytes32 _controllerName) Controllable(_ens, _controllerName) public {
        APIPublicKey = hex"a0f4f688350018ad1b9785991c0bde5f704b005dc79972b114dbed4a615a983710bfc647ebe5a320daa28771dce6a2d104f5efa2e4a85ba3760b76d46f8571ca";
        OAR = OraclizeAddrResolverI(_resolver);
        //todo mischa is this the right thing?
        oraclize_setCustomGasPrice(10000000000);
        oraclize_setProof(proofType_Native);
    }

    // @dev Updates the Crypto Compare public API key.
    function updateAPIPublicKey(bytes _publicKey) external onlyController {
        APIPublicKey = _publicKey;
        emit SetCryptoComparePublicKey(_publicKey);
    }

    // @dev Sets the gas price used by oraclize query.
    function setCustomGasPrice(uint _gasPrice) external onlyController {
        oraclize_setCustomGasPrice(_gasPrice);
        emit SetGasPrice(_gasPrice);
    }

    // @dev Convert ERC20 token amount to the corresponding ether amount (used by the wallet contract).
    // @param _token ERC20 token contract address.
    // @param _amount amount of token in base units.
    function convert(address _token, uint _amount) external view returns (uint) {
        // Store the token in memory to save map entry lookup gas.
        Token storage token = tokens[_token];
        // Require that the token exists and that its rate is not zero.
        require(token.exists && token.rate != 0, "token does not exist");
        // Safely convert the token amount to ether based on the exchange rate.
        return _amount.mul(token.rate).div(token.magnitude);
    }

    // @dev Add ERC20 tokens to the list of supported tokens.
    // @param _tokens ERC20 token contract addresses.
    // @param _symbols ERC20 token names.
    // @param _magnitude 10 to the power of number of decimal places used by each ERC20 token.
    // @param _updateDate date for the token updates. This will be compared to when oracle updates are received.
    function addTokens(address[] _tokens, bytes32[] _symbols, uint[] _magnitude, uint _updateDate) external onlyController {
        // Require that all parameters have the same length.
        require(_tokens.length == _symbols.length && _tokens.length == _magnitude.length, "parameter lengths do not match");
        // Add each token to the list of supported tokens.
        for (uint i = 0; i < _tokens.length; i++) {
            // Require that the token doesn't already exist.
            address token = _tokens[i];
            require(!tokens[token].exists, "token already exists");
            // Store the intermediate values.
            string memory symbol = _symbols[i].toSliceB32().toString();
            uint magnitude = _magnitude[i];
            // Add the token to the token list.
            tokens[token] = Token({
                symbol : symbol,
                magnitude : magnitude,
                rate : 0,
                exists : true,
                lastUpdate: _updateDate
            });
            // Add the token address to the address list.
            _tokenAddresses.push(token);
            // Emit token addition event.
            emit AddedToken(token, symbol, magnitude);
        }
    }

    // @dev Remove ERC20 tokens from the list of supported tokens.
    // @param _tokens ERC20 token contract addresses.
    function removeTokens(address[] _tokens) external onlyController {
        // Delete each token object from the list of supported tokens based on the addresses provided.
        for (uint i = 0; i < _tokens.length; i++) {
            //token must exist, reverts on duplicates as well
            require(tokens[_tokens[i]].exists, "token does not exist");
            // Store the token address.
            address token = _tokens[i];
            // Delete the token object.
            delete tokens[token];
            // Remove the token address from the address list.
            for (uint j = 0; j < _tokenAddresses.length.sub(1); j++) {
                if (_tokenAddresses[j] == token) {
                    _tokenAddresses[j] = _tokenAddresses[_tokenAddresses.length.sub(1)];
                    break;
                }
            }
            _tokenAddresses.length--;
            // Emit token removal event.
            emit RemovedToken(token);
        }
    }

    // @dev Update ERC20 token exchange rate manually.
    // @param _token ERC20 token contract address.
    // @param _rate ERC20 token exchange rate in wei.
    // @param _updateDate date for the token updates. This will be compared to when oracle updates are received.
    function updateTokenRate(address _token, uint _rate, uint _updateDate) external onlyController {
        // Require that the token exists.
        require(tokens[_token].exists, "token does not exist");
        // Update the token's rate.
        tokens[_token].rate = _rate;
        // Update the token's last update timestamp.
        tokens[_token].lastUpdate = _updateDate;
        // Emit the rate update event.
        emit UpdatedTokenRate(_token, _rate);
    }

    // @dev Update ERC20 token exchange rates for all supported tokens.
    function updateTokenRates() external payable onlyController {
        _updateTokenRates();
    }

    /// @dev Withdraw ether from the smart contract to the specified account.
    function withdraw(address _to, uint _amount) external onlyController {
        _to.transfer(_amount);
    }

    /// @dev Handle Oraclize query callback and verifiy the provided origin proof.
    /// @param _queryID Oraclize query ID.
    /// @param _result query result in JSON format.
    /// @param _proof origin proof from crypto compare.
    // solium-disable-next-line mixedcase
    function __callback(bytes32 _queryID, string _result, bytes _proof) public {
        // Require that the caller is the Oraclize contract.
        require(msg.sender == oraclize_cbAddress(), "sender is not oraclize");
        // Use the query ID to find the matching token address.
        address _token = _queryToToken[_queryID];
        require(_token != address(0), "queryID matches to address 0");
        // Get the corresponding token object.
        Token storage token = tokens[_token];

        bool valid;
        uint timestamp;
        (valid, timestamp) = verifyProof(_result, _proof, APIPublicKey, token.lastUpdate);

        // Require that the proof is valid.
        if (valid) {
            // Parse the JSON result to get the rate in wei.
            token.rate = parseInt(parseRate(_result, "ETH"), 18);
            // Emit the rate update event.
            emit UpdatedTokenRate(_token, token.rate);
            // Set the update time of the token rate.
            token.lastUpdate = timestamp;
            // Remove query from the list.
            delete _queryToToken[_queryID];
        } 
    }

    // @dev Re-usable helper function that performs the Oraclize Query.
    function _updateTokenRates() private {
        // Check if there are any existing tokens.
        if (_tokenAddresses.length == 0) {
            // Emit a query failure event.
            emit FailedUpdateRequest("no tokens");
            return;
        // Check if the contract has enough Ether to pay for the query.
        } else if (oraclize_getPrice("URL") * _tokenAddresses.length > address(this).balance) {
            // Emit a query failure event.
            emit FailedUpdateRequest("insufficient balance");
            return;
        } else {
            // Set up the crypto compare API query strings.
            Strings.slice memory apiPrefix = "https://min-api.cryptocompare.com/data/price?fsym=".toSlice();
            Strings.slice memory apiSuffix = "&tsyms=ETH&sign=true".toSlice();

            uint gaslimit = 2000000;
            // Create a new oraclize query for each supported token.
            for (uint i = 0; i < _tokenAddresses.length; i++) {
                // Store the token symbol used in the query.
                Strings.slice memory symbol = tokens[_tokenAddresses[i]].symbol.toSlice();
                // Create a new oraclize query from the component strings.
                bytes32 queryID = oraclize_query("URL", apiPrefix.concat(symbol).toSlice().concat(apiSuffix), gaslimit);
                // Store the query ID together with the associated token address.
                _queryToToken[queryID] = _tokenAddresses[i];
                // Emit the query success event.
                emit RequestedUpdate(symbol.toString());
            }
            return;
        }
    }

    // @dev Verify the origin proof returned by the crypto compare API.
    // @param _result query result in JSON format.
    // @param _proof origin proof from crypto compare.
    // @param _publicKey crypto compare public key.
    // @param _lastUpdate timestamp of the last time the requested token was updated.
    function verifyProof(string _result, bytes _proof, bytes _publicKey, uint _lastUpdate) private returns (bool, uint) {
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

        bool dateValid;
        uint timestamp;
        (dateValid, timestamp) = verifyDate(string(dateHeader), _lastUpdate);

        // Check if the Date returned is valid or not
        if (!dateValid) {
            emit FailedProofVerification(_publicKey, _result, "date");
            return (false, 0);
        }

        // Check if the signed digest hash matches the result hash.
        bytes memory digest = new bytes(headersLength - 52);
        digest = copyBytes(headers, 52, headersLength - 52, digest, 0);
        if (keccak256(sha256(_result)) != keccak256(base64decode(digest))) {
            emit FailedProofVerification(_publicKey, _result, "hash");
            return (false, 0);
        }

        //todo mischa ... this is ugly i would have it the other way round :)
        // Check if the signature is valid and if the signer addresses match.
        if (verifySignature(headers, signature, _publicKey)) {
            emit VerifiedProof(_publicKey, _result);
            return (true, timestamp);
        }

        emit FailedProofVerification(_publicKey, _result, "signature");
        return (false, 0);
    }

    function verifySignature(bytes _headers, bytes _signature, bytes _publicKey) private returns (bool) {
        address signer;
        bool signatureOK;
        (signatureOK, signer) = ecrecovery(sha256(_headers), _signature);
        return signatureOK && signer == address(keccak256(_publicKey));
    }

    // @dev Verify the signed HTTP date header.
    // @param _dateHeader extracted date string e.g. Wed, 12 Sep 2018 15:18:14 GMT.
    // @param _lastUpdate timestamp of the last time the requested token was updated.
    function verifyDate(string _dateHeader, uint _lastUpdate) private pure returns (bool, uint) {
        Strings.slice memory date = _dateHeader.toSlice();
        Strings.slice memory timeDelimiter = ":".toSlice();
        Strings.slice memory dateDelimiter = " ".toSlice();
        // Split the date string.
        date.split(",".toSlice());
        date.split(dateDelimiter);
        // Get individual date components.
        uint day = parseInt(date.split(dateDelimiter).toString());
        uint month = monthToNumber(date.split(dateDelimiter).toString());
        uint year = parseInt(date.split(dateDelimiter).toString());
        uint hour = parseInt(date.split(timeDelimiter).toString());
        uint minute = parseInt(date.split(timeDelimiter).toString());
        uint second = parseInt(date.split(timeDelimiter).toString());

        if (day > 31 || day < 1) {
            return (false, 0);
        }

        if (month > 12 || month < 1) {
            return (false, 0);
        }

        if (year < 2018 || year > 3000) {
            return (false, 0);
        }
    
        if (hour >= 24) {
            return (false, 0);
        }

        if (minute >= 60) {
            return (false, 0);
        }

        if (second >= 60) {
            return (false, 0);
        }

        uint timestamp = year * (10 ** 10) + month * (10 ** 8) + day * (10 ** 6) + hour * (10 ** 4) + minute * (10 ** 2) + second;

        return (timestamp > _lastUpdate, timestamp);
    }
}
