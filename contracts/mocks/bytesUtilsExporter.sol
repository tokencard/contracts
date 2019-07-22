pragma solidity ^0.5.10;

import "../internals/bytesUtils.sol";

contract BytesUtilsExporter {

    using BytesUtils for bytes;

    /// @dev export _bytesToAddress() as an external function.
    function bytesToAddress(bytes calldata _bts, uint _from) external pure returns (address) {
        return _bts._bytesToAddress(_from);
    }

    /// @dev export _bytesToUint32() as an external function.
    function bytesToUint32(bytes calldata _bts, uint _from) external pure returns (uint32) {
        return _bts._bytesToUint32(_from);
    }

    /// @dev export _bytesToUint256() as an external function.
    function bytesToUint256(bytes calldata _bts, uint _from) external pure returns (uint) {
        return _bts._bytesToUint256(_from);
    }
}
