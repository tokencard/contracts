pragma solidity ^0.5.10;

/// @title BytesUtils provides basic byte slicing and casting functionality.
library BytesUtils {

    /// @dev This function converts to an address
    /// @param _bts bytes
    /// @param _from start position
    function _bytesToAddress(bytes memory _bts, uint _from) internal pure returns (address) {
        require(_bts.length >= _from + 20, "slicing out of range");

        bytes20 convertedAddress;
        uint startByte = _from + 32; //first 32 bytes denote the array length

        assembly {
            convertedAddress := mload(add(_bts, startByte))
        }

        return address(convertedAddress);
    }

    /// @dev This function slicing bytes into uint32
    /// @param _bts some bytes
    /// @param _from start position
    function _bytesToUint32(bytes memory _bts, uint _from) internal pure returns (uint32) {
        require(_bts.length >= _from + 4, "slicing out of range");

        bytes4 convertedUint32;
        uint startByte = _from + 32; //first 32 bytes denote the array length

        assembly {
            convertedUint32 := mload(add(_bts, startByte))
        }

        return uint32(convertedUint32);

    }

    /// @dev This function slices a uint
    /// @param _bts some bytes
    /// @param _from start position
    // credit to https://ethereum.stackexchange.com/questions/51229/how-to-convert-bytes-to-uint-in-solidity
    // and Nick Johnson https://ethereum.stackexchange.com/questions/4170/how-to-convert-a-uint-to-bytes-in-solidity/4177#4177
    function _bytesToUint256(bytes memory _bts, uint _from) internal pure returns (uint) {
        require(_bts.length >= _from + 32, "slicing out of range");

        uint x;
        assembly {
            x := mload(add(_bts, add(0x20, _from)))
        }

        return x;
    }
}
