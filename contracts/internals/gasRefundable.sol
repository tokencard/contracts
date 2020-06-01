/**
 *  Controller - The Consumer Contract Wallet
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
 *
 *  SPDX-License-Identifier: GPL-3.0
 */

pragma solidity ^0.5.17;

import "./ensResolvable.sol";


interface IGasToken {
    function free(uint256) external returns (bool);
    function freeUpTo(uint256) external returns (uint256);
    function freeFrom(address, uint256) external returns (bool);
    function freeFromUpTo(address, uint256) external returns (uint256);
}

contract GasRefundable is ENSResolvable {
    /// @notice Emits the GST ENS node when set.
    event GSTSetNode(bytes32 _gstNode);

    /// @dev ENS node for the GST token contract set to the default value (gst2.gastokenio.eth).
    bytes32 private _gstNode = 0x09eda238d4d37e8ac546190722df083c57c5adeae8a32c990781b796ca0886c0;

    /// @param _gstNode_ ENS node for the GST token contract.
	constructor(bytes32 _gstNode_) internal {
        if (_gstNode_ != bytes32(0)) {
            _gstNode = _gstNode_;
            emit GSTSetNode(_gstNode_);
        }
    }

	/// @notice Set the ENS node for the GST token.
    /// @param _gstNode_ A new ENS node for the GST token contract.
    function _gstSetNode(bytes32 _gstNode_) internal {
        _gstNode = _gstNode_;
        emit GSTSetNode(_gstNode_);
    }

    /// @param _amount Amount of GST tokens to be freed.
    function _freeGas(uint256 _amount) internal {
        IGasToken(_ensResolve(_gstNode)).free(_amount);
    }

    /// @param _amount Amount of GST tokens to be freed.
    function _freeGasUpTo(uint256 _amount) internal {
        IGasToken(_ensResolve(_gstNode)).freeUpTo(_amount);
    }
}
