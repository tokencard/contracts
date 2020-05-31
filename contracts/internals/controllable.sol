/**
 *  Controllable - The Consumer Contract Wallet
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

import "../controller.sol";
import "./ensResolvable.sol";

/// @title Controllable implements access control functionality of the Controller smart contract.
contract Controllable is ENSResolvable {
    /// @notice Emits the controller ENS node when set.
    event SetControllerNode(bytes32 _controllerNode);

    /// @notice ENS node for the Controller smart contract, set to the default value of `controller.tokencard.eth`.
    bytes32 private _controllerNode = 0x7f2ce995617d2816b426c5c8698c5ec2952f7a34bb10f38326f74933d5893697;

    /// @notice Constructor initializes the Controller contract object.
    /// @param _controllerNode_ ENS node for the Controller smart contract.
    constructor(bytes32 _controllerNode_) internal {
        if (_controllerNode_ != bytes32(0)) {
            _controllerNode = _controllerNode_;
            emit SetControllerNode(_controllerNode_);
        }
    }

    /// @notice Checks if message sender is a controller.
    modifier onlyController() {
        require(_isController(msg.sender), "sender is not a controller");
        _;
    }

    /// @notice Checks if message sender is an admin.
    modifier onlyAdmin() {
        require(_isAdmin(msg.sender), "sender is not an admin");
        _;
    }

    /// @return Controller node registered in the ENS registry.
    function controllerNode() public view returns (bytes32) {
        return _controllerNode;
    }

    /// @notice Set the controller ENS node to a different one.
    function setControllerNode(bytes32 _controllerNode_) internal {
        _controllerNode = _controllerNode_;
        emit SetControllerNode(_controllerNode_);
    }

    /// @return Result of whether the provided address is the same as the controller contract address.
    function _isController(address _account) internal view returns (bool) {
        return IController(_ensResolve(_controllerNode)).isController(_account);
    }

    /// @return Result of whether the provided address is an admin address.
    function _isAdmin(address _account) internal view returns (bool) {
        return IController(_ensResolve(_controllerNode)).isAdmin(_account);
    }
}
