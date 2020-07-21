// SPDX-License-Identifier: GPLv3

pragma solidity ^0.6.0;

/**
 * This method was modified from the GPLv3 solidity code found in this repository
 * https://github.com/vcealicu/melonport-price-feed/blob/master/pricefeed/PriceFeed.sol
 */

/// @title Base64 provides base 64 decoding functionality.
contract Base64 {
    bytes constant BASE64_DECODE_CHAR = hex"000000000000000000000000000000000000000000000000000000000000000000000000000000000000003e003e003f3435363738393a3b3c3d00000000000000000102030405060708090a0b0c0d0e0f10111213141516171819000000003f001a1b1c1d1e1f202122232425262728292a2b2c2d2e2f30313233";

    /// @return decoded array of bytes.
    /// @param _encoded base 64 encoded array of bytes.
    function _base64decode(bytes memory _encoded) internal pure returns (bytes memory) {
        byte v1;
        byte v2;
        byte v3;
        byte v4;
        uint length = _encoded.length;
        bytes memory result = new bytes(length);
        uint index;

        // base64 encoded strings can't be length 0 and they must be divisble by 4
        require(length > 0  && length % 4 == 0, "invalid base64 encoding");

        if (keccak256(abi.encodePacked(_encoded[length - 2])) == keccak256("=")) {
            length -= 2;
        } else if (keccak256(abi.encodePacked(_encoded[length - 1])) == keccak256("=")) {
            length -= 1;
        }

        uint count = length >> 2 << 2;
        uint i;

        for (i = 0; i < count;) {
            v1 = BASE64_DECODE_CHAR[uint8(_encoded[i++])];
            v2 = BASE64_DECODE_CHAR[uint8(_encoded[i++])];
            v3 = BASE64_DECODE_CHAR[uint8(_encoded[i++])];
            v4 = BASE64_DECODE_CHAR[uint8(_encoded[i++])];

            result[index++] = (v1 << 2 | v2 >> 4) & 0xff;
            result[index++] = (v2 << 4 | v3 >> 2) & 0xff;
            result[index++] = (v3 << 6 | v4) & 0xff;
        }

        if (length - count == 2) {
            v1 = BASE64_DECODE_CHAR[uint8(_encoded[i++])];
            v2 = BASE64_DECODE_CHAR[uint8(_encoded[i++])];

            result[index++] = (v1 << 2 | v2 >> 4) & 0xff;
        } else if (length - count == 3) {
            v1 = BASE64_DECODE_CHAR[uint8(_encoded[i++])];
            v2 = BASE64_DECODE_CHAR[uint8(_encoded[i++])];
            v3 = BASE64_DECODE_CHAR[uint8(_encoded[i++])];

            result[index++] = (v1 << 2 | v2 >> 4) & 0xff;
            result[index++] = (v2 << 4 | v3 >> 2) & 0xff;
        }

        // Set to correct length.
        assembly {
            mstore(result, index)
        }

        return result;
    }
}
