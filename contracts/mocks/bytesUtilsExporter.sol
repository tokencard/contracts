// SPDX-License-Identifier: GPLv3

pragma solidity ^0.6.11;

import "../internals/bytesUtils.sol";


contract BytesUtilsExporter {
    using BytesUtils for bytes;

    /// @dev export _bytesToAddress() as an external function.
    function bytesToAddress(bytes calldata _bts, uint256 _from) external pure returns (address) {
        return _bts._bytesToAddress(_from);
    }

    /// @dev export _bytesToBytes4() as an external function.
    function bytesToBytes4(bytes calldata _bts, uint256 _from) external pure returns (bytes4) {
        return _bts._bytesToBytes4(_from);
    }

    /// @dev export _bytesToUint256() as an external function.
    function bytesToUint256(bytes calldata _bts, uint256 _from) external pure returns (uint256) {
        return _bts._bytesToUint256(_from);
    }
}
