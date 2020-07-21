/**
 *  BytesUtils - The Consumer Contract Wallet
 *  Copyright (C) 2019 The Contract Wallet Company Limited
 *
 *  This program is free software: you can redistribute it and/or modify
 *  it under the terms of the GNU General Public License as published by
 *  the Free Software Foundation, either version 3 of the License, or
 *  (at your option) any later version.

 *  This program is distributed in the hope that it will be useful,
 *  but WITHOUT ANY WARRANTY; without even the implied warranty of
 *  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *  GNU General Public License for more details.

 *  You should have received a copy of the GNU General Public License
 *  along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

// SPDX-License-Identifier: GPLv3

pragma solidity ^0.6.11;

import "../externals/SafeMath.sol";

/// @title BytesUtils provides basic byte slicing and casting functionality.
library BytesUtils {
    using SafeMath for uint256;

    /// @dev This function converts to an address
    /// @param _bts bytes
    /// @param _from start position
    function _bytesToAddress(bytes memory _bts, uint256 _from) internal pure returns (address) {
        require(_bts.length >= _from.add(20), "slicing out of range");

        bytes20 convertedAddress;
        uint256 startByte = _from.add(32); //first 32 bytes denote the array length

        assembly {
            convertedAddress := mload(add(_bts, startByte))
        }

        return address(convertedAddress);
    }

    /// @dev This function slices bytes into bytes4
    /// @param _bts some bytes
    /// @param _from start position
    function _bytesToBytes4(bytes memory _bts, uint256 _from) internal pure returns (bytes4) {
        require(_bts.length >= _from.add(4), "slicing out of range");

        bytes4 slicedBytes4;
        uint256 startByte = _from.add(32); //first 32 bytes denote the array length

        assembly {
            slicedBytes4 := mload(add(_bts, startByte))
        }

        return slicedBytes4;
    }

    /// @dev This function slices a uint
    /// @param _bts some bytes
    /// @param _from start position
    // credit to https://ethereum.stackexchange.com/questions/51229/how-to-convert-bytes-to-uint-in-solidity
    // and Nick Johnson https://ethereum.stackexchange.com/questions/4170/how-to-convert-a-uint-to-bytes-in-solidity/4177#4177
    function _bytesToUint256(bytes memory _bts, uint256 _from) internal pure returns (uint256) {
        require(_bts.length >= _from.add(32), "slicing out of range");

        uint256 convertedUint256;
        uint256 startByte = _from.add(32); //first 32 bytes denote the array length

        assembly {
            convertedUint256 := mload(add(_bts, startByte))
        }

        return convertedUint256;
    }
}
