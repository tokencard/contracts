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
import "./externals/ECDSA.sol";
import "./externals/SafeMath.sol";
import "./externals/strings.sol";
import "./internals/controllable.sol";
import "./internals/date.sol";
import "./internals/ensResolvable.sol";
import "./internals/parseIntScientific.sol";
import "./internals/tokenWhitelistable.sol";

/// @title Oracle provides asset exchange rates and conversion functionality.
contract Oracle is ENSResolvable, Base64, Date, Controllable, ParseIntScientific, TokenWhitelistable {
    using ECDSA for bytes32;
    using SafeMath for uint256;
    using strings for *;

    /*******************/
    /*     Events     */
    /*****************/

    event SetCryptoComparePublicKey(address _sender, bytes _publicKey);
    event VerifiedProof(bytes _publicKey, string _result);

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

    /// @dev Construct the oracle with multiple controllers, address resolver and custom gas price.
    /// @param _ens_ is the address of the ENS.
    /// @param _controllerNode_ is the ENS node corresponding to the Controller.
    /// @param _tokenWhitelistNode_ is the ENS corresponding to the Token Whitelist.
    constructor(
        address _ens_,
        bytes32 _controllerNode_,
        bytes32 _tokenWhitelistNode_
    ) public {
        _initializeENSResolvable(_ens_);
        _initializeControllable(_controllerNode_);
        _initializeTokenWhitelistable(_tokenWhitelistNode_);
        cryptoCompareAPIPublicKey = hex"a0f4f688350018ad1b9785991c0bde5f704b005dc79972b114dbed4a615a983710bfc647ebe5a320daa28771dce6a2d104f5efa2e4a85ba3760b76d46f8571ca";
    }

    /// @dev Updates the Crypto Compare public API key.
    /// @param _publicKey new Crypto Compare public API key
    function updateCryptoCompareAPIPublicKey(bytes calldata _publicKey) external onlyAdmin {
        cryptoCompareAPIPublicKey = _publicKey;
        emit SetCryptoComparePublicKey(msg.sender, _publicKey);
    }

    /// @dev Verifiy the provided origin proof and update token's rate.
    /// @param _token the address of the token to be updated.
    /// @param _result query result in JSON format.
    /// @param _proof origin proof from CryptoCompare.
    function UpdateTokenRate(
        address _token,
        string calldata _result,
        bytes calldata _proof
    ) external {
        // Get the corresponding token object.
        (, , , bool available, , , uint256 lastUpdate) = _getTokenInfo(_token);
        require(available, "token must be available");

        bool valid;
        uint256 timestamp;
        (valid, timestamp) = _verifyProof(_result, _proof, cryptoCompareAPIPublicKey, lastUpdate);

        // Require that the proof is valid.
        if (valid) {
            // Parse the JSON result to get the rate in wei.
            uint256 parsedRate = _parseIntScientificWei(parseRate(_result));
            // Set the update time of the token rate.
            uint256 parsedLastUpdate = timestamp;

            _updateTokenRate(_token, parsedRate, parsedLastUpdate);
        }
    }

    /// @dev Extracts JSON rate value from the response object.
    /// @param _json body of the JSON response from the CryptoCompare API.
    function parseRate(string memory _json) internal pure returns (string memory) {
        uint256 jsonLen = abi.encodePacked(_json).length;
        //{"ETH":}.length = 8, assuming a (maximum of) 18 digit prevision
        require(jsonLen > 8 && jsonLen <= 28, "misformatted input");

        bytes memory jsonPrefix = new bytes(7);
        _copyBytes(jsonPrefix, abi.encodePacked(_json), 0, 7);
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

    /// @dev Copies bytes from source array to destination array.
    /// @param _dst the bytes array that we want to copy data to.
    /// @param _src the bytes array that we want to copy data from.
    /// @param _srcOffset the offset of the source array that we want to start copying data from.
    /// @param _len the length of the data we want to copy from source to destination.
    function _copyBytes(
        bytes memory _dst,
        bytes memory _src,
        uint256 _srcOffset,
        uint256 _len
    ) private pure {
        uint256 dstPtr;
        uint256 srcPtr = _srcOffset;
        assembly {
            dstPtr := add(_dst, 32)
            srcPtr := add(_src, add(32, _srcOffset))
        }
        dstPtr.memcpy(srcPtr, _len);
    }

    /// @dev Verify the origin proof returned by the cryptocompare API.
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

        _copyBytes(signature, _proof, 2, _ECDSA_SIG_LEN);
        // Extract the headers, big endian encoding of headers length
        if (
            uint256(uint8(_proof[_ENCODING_BYTES + _ECDSA_SIG_LEN])) * _MAX_BYTE_SIZE + uint256(uint8(_proof[_ENCODING_BYTES + _ECDSA_SIG_LEN + 1])) !=
            _HEADERS_LEN
        ) {
            revert("invalid headers length");
        }

        bytes memory headers = new bytes(_HEADERS_LEN);
        _copyBytes(headers, _proof, 2 * _ENCODING_BYTES + _ECDSA_SIG_LEN, _HEADERS_LEN);

        // Check if the signature is valid and if the signer address is matching.
        if (!_verifySignature(headers, signature, _publicKey)) {
            revert("invalid signature");
        }

        // Check if the date is valid.
        bytes memory dateHeader = new bytes(20);
        // keep only the relevant string(e.g. "16 Nov 2018 16:22:18")
        _copyBytes(dateHeader, headers, 11, 20);

        bool dateValid;
        uint256 timestamp;
        (dateValid, timestamp) = _verifyDate(string(dateHeader), _lastUpdate);

        // Check whether the date returned is valid or not
        if (!dateValid) {
            revert("invalid date");
        }

        // Check if the signed digest hash matches the result hash.
        bytes memory digest = new bytes(_DIGEST_BASE64_LEN);
        _copyBytes(digest, headers, _DIGEST_OFFSET, _DIGEST_BASE64_LEN);

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
    function _verifySignature(
        bytes memory _headers,
        bytes memory _signature,
        bytes memory _publicKey
    ) private pure returns (bool) {
        address signer;
        // Checks if the signature is valid by hashing the headers
        bytes32 dataHash = sha256(_headers);
        signer = dataHash.recoverNonMalleable(_signature);
        return signer == address(uint256(keccak256(_publicKey)));
    }

    /// @dev Verify the signed HTTP date header.
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
