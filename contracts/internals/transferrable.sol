/**
 *  Transferrable - The Consumer Contract Wallet
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

import "../externals/Address.sol";
import "../externals/ERC20.sol";
import "../externals/SafeERC20.sol";


/// @title SafeTransfer, allowing contract to withdraw tokens accidentally sent to itself
contract Transferrable {
    using Address for address payable;
    using SafeERC20 for IERC20;

    /// @dev This function is used to move tokens sent accidentally to this contract method.
    /// @dev The owner can chose the new destination address
    /// @param _to is the recipient's address.
    /// @param _asset is the address of an ERC20 token or 0x0 for ether.
    /// @param _amount is the amount to be transferred in base units.
    function _safeTransfer(address payable _to, address _asset, uint256 _amount) internal {
        // address(0) is used to denote ETH
        if (_asset == address(0)) {
           _to.sendValue(_amount);
        } else {
            IERC20(_asset).safeTransfer(_to, _amount);
        }
    }
}
