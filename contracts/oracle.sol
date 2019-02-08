/**
 *  The Consumer Contract Wallet
 *  Copyright (C) 2018 Token Group Ltd
 *
 *  This program is free software: you can redistribute it and/or modify
 *  it under the terms of the GNU General Public License as published by
 *  the Free Software Foundation, either version 3 of the License, or
 *  (at your option) any later version.
 *
 *  This program is distributed in the hope that it will be useful,
 *  but WITHOUT ANY WARRANTY; without even the implied warranty of
 *  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *  GNU General Public License for more details.
 *
 *  You should have received a copy of the GNU General Public License
 *  along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

pragma solidity ^0.4.25;

import "./internals/controllable.sol";
import "./internals/claimable.sol";
import "./internals/date.sol";
import "./internals/json.sol";
import "./internals/parseIntScientific.sol";
import "./internals/tokenWhitelistable.sol";
import "./externals/SafeMath.sol";
import "./externals/base64.sol";


/// @title Oracle converts ERC20 token amounts into equivalent ether amounts based on cryptocurrency exchange rates.
interface IOracle {
    function convert(address, uint) external view returns (bool, uint);
}


/// @title Oracle provides asset exchange rates and conversion functionality.
contract Oracle is usingOraclize, Claimable, Base64, Date, JSON, Controllable, ParseIntScientific, TokenWhitelistable, IOracle {
    using strings for *;
    using SafeMath for uint256;


    /*******************/
    /*     Events     */
    /*****************/

    event SetGasPrice(address _sender, uint _gasPrice);
    event Converted(address _sender, address _token, uint _amount, uint _ether);

    event RequestedUpdate(string _symbol);
    event FailedUpdateRequest(string _reason);

    event VerifiedProof(bytes _publicKey, string _result);

    event SetCryptoComparePublicKey(address _sender, bytes _publicKey);

    /**********************/
    /*     Constants     */
    /********************/

    uint constant private PROOF_LEN = 165;
    uint constant private ECDSA_SIG_LEN = 65;
    uint constant private ENCODING_BYTES = 2;
    uint constant private HEADERS_LEN = PROOF_LEN - 2 * ENCODING_BYTES - ECDSA_SIG_LEN; // 2 bytes encoding headers length + 2 for signature.
    uint constant private DIGEST_BASE64_LEN = 44; //base64 encoding of the SHA256 hash (32-bytes) of the result: fixed length.
    uint constant private DIGEST_OFFSET = HEADERS_LEN - DIGEST_BASE64_LEN; // the starting position of the result hash in the headers string.

    uint constant private MAX_BYTE_SIZE = 256; //for calculating length encoding

    bytes public APIPublicKey;
    mapping(bytes32 => address) private _queryToToken;

    /// @dev Construct the oracle with multiple controllers, address resolver and custom gas price.
    /// @dev _resolver is the oraclize address resolver contract address.
    /// @param _ens is the address of the ENS.
    /// @param _controllerName is the ENS name of the Controller.
    constructor(address _resolver, address _ens, bytes32 _controllerName, bytes32 _tokenWhitelistName) Controllable(_ens, _controllerName) TokenWhitelistable(_ens, _tokenWhitelistName) public {
        APIPublicKey = hex"a0f4f688350018ad1b9785991c0bde5f704b005dc79972b114dbed4a615a983710bfc647ebe5a320daa28771dce6a2d104f5efa2e4a85ba3760b76d46f8571ca";
        OAR = OraclizeAddrResolverI(_resolver);
        oraclize_setCustomGasPrice(10000000000);
        oraclize_setProof(proofType_Native);
    }

    /// @dev Updates the Crypto Compare public API key.
    function updateAPIPublicKey(bytes _publicKey) external onlyController {
        APIPublicKey = _publicKey;
        emit SetCryptoComparePublicKey(msg.sender, _publicKey);
    }

    /// @dev Sets the gas price used by oraclize query.
    function setCustomGasPrice(uint _gasPrice) external onlyController {
        oraclize_setCustomGasPrice(_gasPrice);
        emit SetGasPrice(msg.sender, _gasPrice);
    }

    /// @dev Convert ERC20 token amount to the corresponding ether amount (used by the wallet contract).
    /// @param _token ERC20 token contract address.
    /// @param _amount amount of token in base units.
    function convert(address _token, uint _amount) external view returns (bool, uint) {
        // Store the token in memory to save map entry lookup gas.
        ( , uint256 magnitude, uint256 rate, bool available, , ) = _getTokenInfo(_token);
        // If the token exists require that its rate is not zero
        if (available) {
            require(rate != 0, "token rate is 0");
            // Safely convert the token amount to ether based on the exchange rate.
            // return the value and a 'true' implying that the token is protected
            return (true, _amount.mul(rate).div(magnitude));
        }
        // this returns a 'false' to imply that a card is not protected
        return (false, 0);

    }

    /// @dev Update ERC20 token exchange rates for all supported tokens.
    //// @param _gasLimit the gas limit is passed, this is used for the Oraclize callback
    function updateTokenRates(uint _gasLimit) external payable onlyController {
        _updateTokenRates(_gasLimit);
    }

    //// @dev Withdraw tokens from the smart contract to the specified account.
    function claim(address _to, address _asset, uint _amount) external onlyController {
        _claim(_to, _asset, _amount);
    }

    //// @dev Handle Oraclize query callback and verifiy the provided origin proof.
    //// @param _queryID Oraclize query ID.
    //// @param _result query result in JSON format.
    //// @param _proof origin proof from crypto compare.
    // solium-disable-next-line mixedcase
    function __callback(bytes32 _queryID, string _result, bytes _proof) public {
        // Require that the caller is the Oraclize contract.
        require(msg.sender == oraclize_cbAddress(), "sender is not oraclize");
        // Use the query ID to find the matching token address.
        address _token = _queryToToken[_queryID];
        require(_token != address(0), "queryID matches to address 0");
        // Get the corresponding token object.
        ( , , , , , uint256 lastUpdate) = _getTokenInfo(_token);

        bool valid;
        uint timestamp;
        (valid, timestamp) = _verifyProof(_result, _proof, APIPublicKey, lastUpdate);

        // Require that the proof is valid.
        if (valid) {
            // Parse the JSON result to get the rate in wei.
            uint256 parsedRate = _parseIntScientificWei(parseRate(_result));
            // Set the update time of the token rate.
            uint256 parsedLastUpdate = timestamp;
            // Remove query from the list.
            delete _queryToToken[_queryID];

            _updateTokenRate(_token, parsedRate, parsedLastUpdate);
        }
    }

    /// @dev Re-usable helper function that performs the Oraclize Query.
    //// @param _gasLimit the gas limit is passed, this is used for the Oraclize callback
    function _updateTokenRates(uint _gasLimit) private {
        address[] memory tokenAddresses = _getTokenAddressArray();
        // Check if there are any existing tokens.
        if (tokenAddresses.length == 0) {
            // Emit a query failure event.
            emit FailedUpdateRequest("no tokens");
            // Check if the contract has enough Ether to pay for the query.
        } else if (oraclize_getPrice("URL") * tokenAddresses.length > address(this).balance) {
            // Emit a query failure event.
            emit FailedUpdateRequest("insufficient balance");
        } else {
            // Set up the cryptocompare API query strings.
            strings.slice memory apiPrefix = "https://min-api.cryptocompare.com/data/price?fsym=".toSlice();
            strings.slice memory apiSuffix = "&tsyms=ETH&sign=true".toSlice();

            // Create a new oraclize query for each supported token.
            for (uint i = 0; i < tokenAddresses.length; i++) {
                // Store the token symbol used in the query.
                (string memory symbol, , , , , ) = _getTokenInfo(tokenAddresses[i]);

                strings.slice memory sym = symbol.toSlice();
                // Create a new oraclize query from the component strings.
                bytes32 queryID = oraclize_query("URL", apiPrefix.concat(sym).toSlice().concat(apiSuffix), _gasLimit);
                // Store the query ID together with the associated token address.
                _queryToToken[queryID] = tokenAddresses[i];
                // Emit the query success event.
                emit RequestedUpdate(sym.toString());
            }
        }
    }

    /// @dev Verify the origin proof returned by the cryptocompare API.
    /// @param _result query result in JSON format.
    /// @param _proof origin proof from cryptocompare.
    /// @param _publicKey cryptocompare public key.
    /// @param _lastUpdate timestamp of the last time the requested token was updated.
    function _verifyProof(string _result, bytes _proof, bytes _publicKey, uint _lastUpdate) private returns (bool, uint) {

        //expecting fixed length proofs
        if (_proof.length != PROOF_LEN) {
            revert("invalid proof length");
        }

        //proof should be 65 bytes long: R (32 bytes) + S (32 bytes) + v (1 byte)
        if (uint(_proof[1]) != ECDSA_SIG_LEN) {
            revert("invalid signature length");
        }

        bytes memory signature = new bytes(ECDSA_SIG_LEN);

        signature = copyBytes(_proof, 2, ECDSA_SIG_LEN, signature, 0);

        // Extract the headers, big endian encoding of headers length
        if (uint(_proof[ENCODING_BYTES + ECDSA_SIG_LEN]) * MAX_BYTE_SIZE + uint(_proof[ENCODING_BYTES + ECDSA_SIG_LEN + 1]) != HEADERS_LEN) {
            revert("invalid headers length");
        }

        bytes memory headers = new bytes(HEADERS_LEN);
        headers = copyBytes(_proof, 2 * ENCODING_BYTES + ECDSA_SIG_LEN, HEADERS_LEN, headers, 0);

        // Check if the signature is valid and if the signer address is matching.
        if (!_verifySignature(headers, signature, _publicKey)) {
            revert("invalid signature");
        }

        // Check if the date is valid.
        bytes memory dateHeader = new bytes(20);
        //keep only the relevant string(e.g. "16 Nov 2018 16:22:18")
        dateHeader = copyBytes(headers, 11, 20, dateHeader, 0);

        bool dateValid;
        uint timestamp;
        (dateValid, timestamp) = _verifyDate(string(dateHeader), _lastUpdate);

        // Check whether the date returned is valid or not
        if (!dateValid) {
            revert("invalid date");
        }

        // Check if the signed digest hash matches the result hash.
        bytes memory digest = new bytes(DIGEST_BASE64_LEN);
        digest = copyBytes(headers, DIGEST_OFFSET, DIGEST_BASE64_LEN, digest, 0);

        if (keccak256(abi.encodePacked(sha256(abi.encodePacked(_result)))) != keccak256(_base64decode(digest))) {
            revert("result hash not matching");
        }

        emit VerifiedProof(_publicKey, _result);
        return (true, timestamp);
    }

    /// @dev Verify the HTTP headers and the signature
    /// @param _headers HTTP headers provided by the cryptocompare api
    /// @param _signature signature provided by the cryptocompare api
    /// @param _publicKey cryptocompare public key.
    function _verifySignature(bytes _headers, bytes _signature, bytes _publicKey) private returns (bool) {
        address signer;
        bool signatureOK;

        // Checks if the signature is valid by hashing the headers
        (signatureOK, signer) = ecrecovery(sha256(_headers), _signature);
        return signatureOK && signer == address(keccak256(_publicKey));
    }

    /// @dev Verify the signed HTTP date header.
    /// @param _dateHeader extracted date string e.g. Wed, 12 Sep 2018 15:18:14 GMT.
    /// @param _lastUpdate timestamp of the last time the requested token was updated.
    function _verifyDate(string _dateHeader, uint _lastUpdate) private pure returns (bool, uint) {

        //called by verifyProof(), _dateHeader is always a string of length = 20
        assert(abi.encodePacked(_dateHeader).length == 20);

        //Split the date string and get individual date components.
        strings.slice memory date = _dateHeader.toSlice();
        strings.slice memory timeDelimiter = ":".toSlice();
        strings.slice memory dateDelimiter = " ".toSlice();

        uint day = _parseIntScientific(date.split(dateDelimiter).toString());
        require(day > 0 && day < 32, "day error");

        uint month = _monthToNumber(date.split(dateDelimiter).toString());
        require(month > 0 && month < 13, "month error");

        uint year = _parseIntScientific(date.split(dateDelimiter).toString());
        require(year > 2017 && year < 3000, "year error");

        uint hour = _parseIntScientific(date.split(timeDelimiter).toString());
        require(hour < 25, "hour error");

        uint minute = _parseIntScientific(date.split(timeDelimiter).toString());
        require(minute < 60, "minute error");

        uint second = _parseIntScientific(date.split(timeDelimiter).toString());
        require(second < 60, "second error");

        uint timestamp = year * (10 ** 10) + month * (10 ** 8) + day * (10 ** 6) + hour * (10 ** 4) + minute * (10 ** 2) + second;

        return (timestamp > _lastUpdate, timestamp);
    }
}
