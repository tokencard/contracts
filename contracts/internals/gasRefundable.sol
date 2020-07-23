/**
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

pragma solidity ^0.5.17;

import "./ensResolvable.sol";


interface IGasToken {
    function free(uint256) external returns (bool);

    function freeUpTo(uint256) external returns (uint256);

    function freeFrom(address, uint256) external returns (bool);

    function freeFromUpTo(address, uint256) external returns (uint256);
}


contract GasRefundable is ENSResolvable {
    /// @notice Emits the gas token ENS node when set.
    event GasTokenSetNode(bytes32 _gasTokenNode);

    /// @dev ENS node for the gas token contract set to the default value (gst2.gastokenio.eth).
    bytes32 private _gasTokenNode = 0x09eda238d4d37e8ac546190722df083c57c5adeae8a32c990781b796ca0886c0;

    /// @param _gasTokenNode_ ENS node for the gas token contract.
    constructor(bytes32 _gasTokenNode_) internal {
        if (_gasTokenNode_ != bytes32(0)) {
            _gasTokenNode = _gasTokenNode_;
            emit GasTokenSetNode(_gasTokenNode_);
        }
    }

    /// @notice Set the ENS node for the gas token.
    /// @param _gasTokenNode_ A new ENS node for the gas token contract.
    function _gasTokenSetNode(bytes32 _gasTokenNode_) internal {
        _gasTokenNode = _gasTokenNode_;
        emit GasTokenSetNode(_gasTokenNode_);
    }

    /// @param _amount Amount of gas tokens to be freed.
    function _freeGas(uint256 _amount) internal {
        IGasToken(_ensResolve(_gasTokenNode)).free(_amount);
    }

    /// @param _amount Amount of gas tokens to be freed.
    function _freeGasUpTo(uint256 _amount) internal {
        IGasToken(_ensResolve(_gasTokenNode)).freeUpTo(_amount);
    }
}
