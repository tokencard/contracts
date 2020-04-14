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

pragma solidity ^0.6.0;

import "./externals/base64.sol";
import "./externals/SafeMath.sol";
import "./externals/strings.sol";
import "./internals/controllable.sol";
import "./internals/date.sol";
import "./internals/ensResolvable.sol";
import "./internals/parseIntScientific.sol";
import "./internals/tokenWhitelistable.sol";


/// @title Oracle provides asset exchange rates and conversion functionality.
contract Oracle is ENSResolvable, Base64, Date, Controllable, ParseIntScientific, TokenWhitelistable {
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

    /// @notice Construct the oracle with multiple controllers, address resolver and custom gas price.
    /// @param _ens_ is the address of the ENS.
    /// @param _controllerNode_ is the ENS node corresponding to the Controller.
    /// @param _tokenWhitelistNode_ is the ENS corresponding to the Token Whitelist.
    constructor(
        address _ens_,
        bytes32 _controllerNode_,
        bytes32 _tokenWhitelistNode_
    ) public ENSResolvable(_ens_) Controllable(_controllerNode_) TokenWhitelistable(_tokenWhitelistNode_) {
        cryptoCompareAPIPublicKey = hex"a0f4f688350018ad1b9785991c0bde5f704b005dc79972b114dbed4a615a983710bfc647ebe5a320daa28771dce6a2d104f5efa2e4a85ba3760b76d46f8571ca";
    }

    /// @notice Updates the Crypto Compare public API key.
    /// @param _publicKey new Crypto Compare public API key
    function updateCryptoCompareAPIPublicKey(bytes calldata _publicKey) external onlyAdmin {
        cryptoCompareAPIPublicKey = _publicKey;
        emit SetCryptoComparePublicKey(msg.sender, _publicKey);
    }

    /// @notice Handle the off-chain oracle call and verifiy the provided origin proof.
    /// @param _result query result in JSON format.
    /// @param _proof origin proof from CryptoCompare.
    // solium-disable-next-line mixedcase
    function __callback(
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

    /// @notice Extracts JSON rate value from the response object.
    /// @param _json body of the JSON response from the CryptoCompare API.
    function parseRate(string memory _json) internal pure returns (string memory) {
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

    /*
     The following function has been written by Alex Beregszaszi, use it under the terms of the MIT license
    */
    function copyBytes(
        bytes memory _from,
        uint256 _fromOffset,
        uint256 _length,
        bytes memory _to,
        uint256 _toOffset
    ) internal pure returns (bytes memory _copiedBytes) {
        uint256 minLength = _length + _toOffset;
        require(_to.length >= minLength); // Buffer too small. Should be a better way?
        uint256 i = 32 + _fromOffset; // NOTE: the offset 32 is added to skip the `size` field of both bytes variables
        uint256 j = 32 + _toOffset;
        while (i < (32 + _fromOffset + _length)) {
            assembly {
                let tmp := mload(add(_from, i))
                mstore(add(_to, j), tmp)
            }
            i += 32;
            j += 32;
        }
        return _to;
    }

    /*
     The following function has been written by Alex Beregszaszi, use it under the terms of the MIT license
     Duplicate Solidity's ecrecover, but catching the CALL return value
    */
    function safer_ecrecover(
        bytes32 _hash,
        uint8 _v,
        bytes32 _r,
        bytes32 _s
    ) internal returns (bool _success, address _recoveredAddress) {
        /*
         We do our own memory management here. Solidity uses memory offset
         0x40 to store the current end of memory. We write past it (as
         writes are memory extensions), but don't update the offset so
         Solidity will reuse it. The memory used here is only needed for
         this context.
         FIXME: inline assembly can't access return values
        */
        bool ret;
        address addr;
        assembly {
            let size := mload(0x40)
            mstore(size, _hash)
            mstore(add(size, 32), _v)
            mstore(add(size, 64), _r)
            mstore(add(size, 96), _s)
            ret := call(3000, 1, 0, size, 128, size, 32) // NOTE: we can reuse the request memory because we deal with the return code.
            addr := mload(size)
        }
        return (ret, addr);
    }

    /*
     The following function has been written by Alex Beregszaszi, use it under the terms of the MIT license
    */
    function ecrecovery(bytes32 _hash, bytes memory _sig) internal returns (bool _success, address _recoveredAddress) {
        bytes32 r;
        bytes32 s;
        uint8 v;
        if (_sig.length != 65) {
            return (false, address(0));
        }
        /*
         The signature format is a compact form of:
           {bytes32 r}{bytes32 s}{uint8 v}
         Compact means, uint8 is not padded to 32 bytes.
        */
        assembly {
            r := mload(add(_sig, 32))
            s := mload(add(_sig, 64))
            /*
             Here we are loading the last 32 bytes. We exploit the fact that
             'mload' will pad with zeroes if we overread.
             There is no 'mload8' to do this, but that would be nicer.
            */
            v := byte(0, mload(add(_sig, 96)))
            /*
              Alternative solution:
              'byte' is not working due to the Solidity parser, so lets
              use the second best option, 'and'
              v := and(mload(add(_sig, 65)), 255)
            */
        }
        /*
         albeit non-transactional signatures are not specified by the YP, one would expect it
         to match the YP range of [27, 28]
         geth uses [0, 1] and some clients have followed. This might change, see:
         https://github.com/ethereum/go-ethereum/issues/2053
        */
        if (v < 27) {
            v += 27;
        }
        if (v != 27 && v != 28) {
            return (false, address(0));
        }
        return safer_ecrecover(_hash, v, r, s);
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
