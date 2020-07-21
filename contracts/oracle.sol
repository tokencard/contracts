/**
 *  Oracle - The Consumer Contract Wallet
 *  Copyright (C) 2019 The Contract Wallet Company Limited
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

// SPDX-License-Identifier: GPLv3

pragma solidity ^0.6.11;

import "./externals/base64.sol";
import "./externals/provableAPI_0.6.sol";
import "./externals/SafeMath.sol";
import "./internals/controllable.sol";
import "./internals/date.sol";
import "./internals/ensResolvable.sol";
import "./internals/parseIntScientific.sol";
import "./internals/tokenWhitelistable.sol";
import "./internals/transferrable.sol";

/// @title Oracle provides asset exchange rates and conversion functionality.
contract Oracle is ENSResolvable, usingProvable, Transferrable, Base64, Date, Controllable, ParseIntScientific, TokenWhitelistable {
    using strings for *;
    using SafeMath for uint256;

    /*******************/
    /*     Events     */
    /*****************/

    event SetGasPrice(address _sender, uint256 _gasPrice);

    event RequestedUpdate(string _symbol, bytes32 _queryID);
    event FailedUpdateRequest(string _reason);

    event VerifiedProof(bytes _publicKey, string _result);

    event SetCryptoComparePublicKey(address _sender, bytes _publicKey);

    event Claimed(address _to, address _asset, uint256 _amount);

    /**********************/
    /*     Constants     */
    /********************/

    uint256 private constant _PROOF_LEN = 165;
    uint256 private constant _ECDSA_SIG_LEN = 65;
    uint256 private constant _ENCODING_BYTES = 2;
    uint256 private constant _HEADERS_LEN = _PROOF_LEN - 2 * _ENCODING_BYTES - _ECDSA_SIG_LEN; // 2 bytes encoding headers length + 2 for signature.
    uint256 private constant _DIGEST_BASE64_LEN = 44; //base64 encoding of the SHA256 hash (32-bytes) of the result: fixed length.
    uint256 private constant _DIGEST_OFFSET = _HEADERS_LEN - _DIGEST_BASE64_LEN; // the starting position of the result hash in the headers string.

    uint256 private constant _MAX_BYTE_SIZE = 256; //for calculating length encoding

    // This is how the cryptocompare json begins
    bytes32 private constant _PREFIX_HASH = keccak256('{"ETH":');

    bytes public cryptoCompareAPIPublicKey;
    mapping(bytes32 => address) private _queryToToken;

    /// @notice Construct the oracle with multiple controllers, address resolver and custom gas price.
    /// @param _resolver_ is the address of the Provable resolver
    /// @param _ens_ is the address of the ENS.
    /// @param _controllerNode_ is the ENS node corresponding to the Controller.
    /// @param _tokenWhitelistNode_ is the ENS corresponding to the Token Whitelist.
    constructor(
        address _resolver_,
        address _ens_,
        bytes32 _controllerNode_,
        bytes32 _tokenWhitelistNode_
    ) public {
        _initializeENSResolvable(_ens_);
        _initializeControllable(_controllerNode_);
        _initializeTokenWhitelistable(_tokenWhitelistNode_);
        cryptoCompareAPIPublicKey = hex"a0f4f688350018ad1b9785991c0bde5f704b005dc79972b114dbed4a615a983710bfc647ebe5a320daa28771dce6a2d104f5efa2e4a85ba3760b76d46f8571ca";
        OAR = OracleAddrResolverI(_resolver_);
        provable_setCustomGasPrice(10000000000);
        provable_setProof(proofType_Native);
    }

    /// @notice Updates the Crypto Compare public API key.
    /// @param _publicKey new Crypto Compare public API key
    function updateCryptoCompareAPIPublicKey(bytes calldata _publicKey) external onlyAdmin {
        cryptoCompareAPIPublicKey = _publicKey;
        emit SetCryptoComparePublicKey(msg.sender, _publicKey);
    }

    /// @notice Sets the gas price used by Provable query.
    /// @param _gasPrice in wei for Provable
    function setCustomGasPrice(uint256 _gasPrice) external onlyController {
        provable_setCustomGasPrice(_gasPrice);
        emit SetGasPrice(msg.sender, _gasPrice);
    }

    /// @notice Update ERC20 token exchange rates for all supported tokens.
    /// @param _gasLimit the gas limit is passed, this is used for the Provable callback
    function updateTokenRates(uint256 _gasLimit) external payable onlyController {
        _updateTokenRates(_gasLimit);
    }

    /// @notice Update ERC20 token exchange rates for the list of tokens provided.
    /// @param _gasLimit the gas limit is passed, this is used for the Provable callback
    /// @param _tokenList the list of tokens that need to be updated
    function updateTokenRatesList(uint256 _gasLimit, address[] calldata _tokenList) external payable onlyController {
        _updateTokenRatesList(_gasLimit, _tokenList);
    }

    /// @notice Withdraw tokens from the smart contract to the specified account.
    function claim(
        address payable _to,
        address _asset,
        uint256 _amount
    ) external onlyAdmin {
        _safeTransfer(_to, _asset, _amount);
        emit Claimed(_to, _asset, _amount);
    }

    /// @notice Handle Provable query callback and verifiy the provided origin proof.
    /// @param _queryID Provable query ID.
    /// @param _result query result in JSON format.
    /// @param _proof origin proof from crypto compare.
    // solium-disable-next-line mixedcase
    function __callback(
        bytes32 _queryID,
        string memory _result,
        bytes memory _proof
    ) public override {
        // Require that the caller is the Provable contract.
        require(msg.sender == provable_cbAddress(), "sender is not provable");
        // Use the query ID to find the matching token address.
        address token = _queryToToken[_queryID];
        // Get the corresponding token object.
        (, , , bool available, , , uint256 lastUpdate) = _getTokenInfo(token);
        require(available, "token must be available");

        bool valid;
        uint256 timestamp;
        (valid, timestamp) = _verifyProof(_result, _proof, cryptoCompareAPIPublicKey, lastUpdate);

        // Require that the proof is valid.
        if (valid) {
            // Parse the JSON result to get the rate in wei.
            uint256 parsedRate = _parseIntScientificWei(_parseRate(_result));
            // Set the update time of the token rate.
            uint256 parsedLastUpdate = timestamp;
            // Remove query from the list.
            delete _queryToToken[_queryID];

            _updateTokenRate(token, parsedRate, parsedLastUpdate);
        }
    }

    /// @notice Extracts JSON rate value from the response object.
    /// @param _json body of the JSON response from the CryptoCompare API.
    function _parseRate(string memory _json) internal pure returns (string memory) {
        uint256 jsonLen = abi.encodePacked(_json).length;
        //{"ETH":}.length = 8, assuming a (maximum of) 18 digit prevision
        require(jsonLen > 8 && jsonLen <= 28, "misformatted input");

        bytes memory jsonPrefix = new bytes(7);
        copyBytes(abi.encodePacked(_json), 0, 7, jsonPrefix, 0);
        require(keccak256(jsonPrefix) == _PREFIX_HASH, "prefix mismatch");

        strings.slice memory body = _json.toSlice();
        body.split(":".toSlice());
        //we are sure that ':' is included in the string, body now contains the rate+'}'
        jsonLen = body._len;
        body.until("}".toSlice());
        require(body._len == jsonLen - 1, "not json format");
        //ensure that the json is properly terminated with a '}'
        return body.toString();
    }

    /// @notice Re-usable helper function that performs the Provable Query.
    /// @param _gasLimit the gas limit is passed, this is used for the Provable callback
    function _updateTokenRates(uint256 _gasLimit) private {
        address[] memory tokenAddresses = _tokenAddressArray();
        // Check if there are any existing tokens.
        if (tokenAddresses.length == 0) {
            // Emit a query failure event.
            emit FailedUpdateRequest("no tokens");
            // Check if the contract has enough Ether to pay for the query.
        } else if (provable_getPrice("URL") * tokenAddresses.length > address(this).balance) {
            // Emit a query failure event.
            emit FailedUpdateRequest("insufficient balance");
        } else {
            // Set up the cryptocompare API query strings.
            strings.slice memory apiPrefix = "https://min-api.cryptocompare.com/data/price?fsym=".toSlice();
            strings.slice memory apiSuffix = "&tsyms=ETH&sign=true".toSlice();

            // Create a new Provable query for each supported token.
            for (uint256 i = 0; i < tokenAddresses.length; i++) {
                // Store the token symbol used in the query.
                (string memory symbol, , , , , , ) = _getTokenInfo(tokenAddresses[i]);

                strings.slice memory sym = symbol.toSlice();
                // Create a new Provable query from the component strings.
                bytes32 queryID = provable_query("URL", apiPrefix.concat(sym).toSlice().concat(apiSuffix), _gasLimit);
                // Store the query ID together with the associated token address.
                _queryToToken[queryID] = tokenAddresses[i];
                // Emit the query success event.
                emit RequestedUpdate(sym.toString(), queryID);
            }
        }
    }

    /// @notice Re-usable helper function that performs the Provable Query for a specific list of tokens.
    /// @param _gasLimit the gas limit is passed, this is used for the Provable callback.
    /// @param _tokenList the list of tokens that need to be updated.
    function _updateTokenRatesList(uint256 _gasLimit, address[] memory _tokenList) private {
        // Check if there are any existing tokens.
        if (_tokenList.length == 0) {
            // Emit a query failure event.
            emit FailedUpdateRequest("empty token list");
            // Check if the contract has enough Ether to pay for the query.
        } else if (provable_getPrice("URL") * _tokenList.length > address(this).balance) {
            // Emit a query failure event.
            emit FailedUpdateRequest("insufficient balance");
        } else {
            // Set up the cryptocompare API query strings.
            strings.slice memory apiPrefix = "https://min-api.cryptocompare.com/data/price?fsym=".toSlice();
            strings.slice memory apiSuffix = "&tsyms=ETH&sign=true".toSlice();

            // Create a new Provable query for each supported token.
            for (uint256 i = 0; i < _tokenList.length; i++) {
                //token must exist, revert if it doesn't
                (string memory tokenSymbol, , , bool available, , , ) = _getTokenInfo(_tokenList[i]);
                require(available, "token must be available");
                // Store the token symbol used in the query.
                strings.slice memory symbol = tokenSymbol.toSlice();
                // Create a new Provable query from the component strings.
                bytes32 queryID = provable_query("URL", apiPrefix.concat(symbol).toSlice().concat(apiSuffix), _gasLimit);
                // Store the query ID together with the associated token address.
                _queryToToken[queryID] = _tokenList[i];
                // Emit the query success event.
                emit RequestedUpdate(symbol.toString(), queryID);
            }
        }
    }

    /// @notice Verify the origin proof returned by the cryptocompare API.
    /// @param _result query result in JSON format.
    /// @param _proof origin proof from cryptocompare.
    /// @param _publicKey cryptocompare public key.
    /// @param _lastUpdate timestamp of the last time the requested token was updated.
    function _verifyProof(
        string memory _result,
        bytes memory _proof,
        bytes memory _publicKey,
        uint256 _lastUpdate
    ) private returns (bool, uint256) {
        // expecting fixed length proofs
        if (_proof.length != _PROOF_LEN) {
            revert("invalid proof length");
        }

        // proof should be 65 bytes long: R (32 bytes) + S (32 bytes) + v (1 byte)
        if (uint256(uint8(_proof[1])) != _ECDSA_SIG_LEN) {
            revert("invalid signature length");
        }

        bytes memory signature = new bytes(_ECDSA_SIG_LEN);

        signature = copyBytes(_proof, 2, _ECDSA_SIG_LEN, signature, 0);

        // Extract the headers, big endian encoding of headers length
        if (
            uint256(uint8(_proof[_ENCODING_BYTES + _ECDSA_SIG_LEN])) * _MAX_BYTE_SIZE + uint256(uint8(_proof[_ENCODING_BYTES + _ECDSA_SIG_LEN + 1])) !=
            _HEADERS_LEN
        ) {
            revert("invalid headers length");
        }

        bytes memory headers = new bytes(_HEADERS_LEN);
        headers = copyBytes(_proof, 2 * _ENCODING_BYTES + _ECDSA_SIG_LEN, _HEADERS_LEN, headers, 0);

        // Check if the signature is valid and if the signer address is matching.
        if (!_verifySignature(headers, signature, _publicKey)) {
            revert("invalid signature");
        }

        // Check if the date is valid.
        bytes memory dateHeader = new bytes(20);
        // keep only the relevant string(e.g. "16 Nov 2018 16:22:18")
        dateHeader = copyBytes(headers, 11, 20, dateHeader, 0);

        bool dateValid;
        uint256 timestamp;
        (dateValid, timestamp) = _verifyDate(string(dateHeader), _lastUpdate);

        // Check whether the date returned is valid or not
        if (!dateValid) {
            revert("invalid date");
        }

        // Check if the signed digest hash matches the result hash.
        bytes memory digest = new bytes(_DIGEST_BASE64_LEN);
        digest = copyBytes(headers, _DIGEST_OFFSET, _DIGEST_BASE64_LEN, digest, 0);

        if (keccak256(abi.encodePacked(sha256(abi.encodePacked(_result)))) != keccak256(_base64decode(digest))) {
            revert("result hash not matching");
        }

        emit VerifiedProof(_publicKey, _result);
        return (true, timestamp);
    }

    /// @notice Verify the HTTP headers and the signature
    /// @param _headers HTTP headers provided by the cryptocompare api
    /// @param _signature signature provided by the cryptocompare api
    /// @param _publicKey cryptocompare public key.
    function _verifySignature(
        bytes memory _headers,
        bytes memory _signature,
        bytes memory _publicKey
    ) private returns (bool) {
        address signer;
        bool signatureOK;

        // Checks if the signature is valid by hashing the headers
        (signatureOK, signer) = ecrecovery(sha256(_headers), _signature);
        return signatureOK && signer == address(uint160(uint256(keccak256(_publicKey))));
    }

    /// @notice Verify the signed HTTP date header.
    /// @param _dateHeader extracted date string e.g. Wed, 12 Sep 2018 15:18:14 GMT.
    /// @param _lastUpdate timestamp of the last time the requested token was updated.
    function _verifyDate(string memory _dateHeader, uint256 _lastUpdate) private pure returns (bool, uint256) {
        // called by verifyProof(), _dateHeader is always a string of length = 20
        assert(abi.encodePacked(_dateHeader).length == 20);

        // Split the date string and get individual date components.
        strings.slice memory date = _dateHeader.toSlice();
        strings.slice memory timeDelimiter = ":".toSlice();
        strings.slice memory dateDelimiter = " ".toSlice();

        uint256 day = _parseIntScientific(date.split(dateDelimiter).toString());
        require(day > 0 && day < 32, "day error");

        uint256 month = _monthToNumber(date.split(dateDelimiter).toString());
        require(month > 0 && month < 13, "month error");

        uint256 year = _parseIntScientific(date.split(dateDelimiter).toString());
        require(year > 2017 && year < 3000, "year error");

        uint256 hour = _parseIntScientific(date.split(timeDelimiter).toString());
        require(hour < 25, "hour error");

        uint256 minute = _parseIntScientific(date.split(timeDelimiter).toString());
        require(minute < 60, "minute error");

        uint256 second = _parseIntScientific(date.split(timeDelimiter).toString());
        require(second < 60, "second error");

        uint256 timestamp = year * (10**10) + month * (10**8) + day * (10**6) + hour * (10**4) + minute * (10**2) + second;

        return (timestamp > _lastUpdate, timestamp);
    }
}
