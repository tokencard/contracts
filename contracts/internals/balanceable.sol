/**
 *  Balanceable - The Consumer Contract Wallet
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

import "../interfaces/IERC20.sol";

/// @title Balanceable - This is a contract used to get a balance
contract Balanceable {
    /// @dev This function is used to get a balance.
    /// @param _asset is the address of an ERC20 token or 0x0 for ether.
    /// @return balance associated with an address, for any token, in the wei equivalent
    function _balance(address _asset) internal view returns (uint256) {
        if (_asset != address(0)) {
            return IERC20(_asset).balanceOf(address(this));
        } else {
            return address(this).balance;
        }
    }
}
