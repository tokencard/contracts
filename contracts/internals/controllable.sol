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
 */

// SPDX-License-Identifier: GPLv3

pragma solidity ^0.6.11;

import "./ensResolvable.sol";
import "../controller.sol";
import "../externals/initializable.sol";


/// @title Controllable implements access control functionality of the Controller found via ENS.
contract Controllable is ENSResolvable {
    // Default values for mainnet ENS
    // controller.tokencard.eth
    bytes32 private constant _DEFAULT_CONTROLLER_NODE = 0x7f2ce995617d2816b426c5c8698c5ec2952f7a34bb10f38326f74933d5893697;

    /// @dev Is the registered ENS node identifying the controller contract.
    bytes32 private _controllerNode = _DEFAULT_CONTROLLER_NODE;

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

    /// @return the controller node registered in ENS.
    function controllerNode() public view returns (bytes32) {
        return _controllerNode;
    }

    /// @notice Initializes the controller contract object.
    /// @param _controllerNode_ is the ENS node of the Controller.
    /// @dev pass in bytes32(0) to use the default, production node labels for ENS
    function _initializeControllable(bytes32 _controllerNode_) internal initializer {
        // Set controllerNode or use default
        if (_controllerNode_ != bytes32(0)) {
            _controllerNode = _controllerNode_;
        }
    }

    /// @return true if the provided account is a controller.
    function _isController(address _account) internal view returns (bool) {
        return IController(_ensResolve(_controllerNode)).isController(_account);
    }

    /// @return true if the provided account is an admin.
    function _isAdmin(address _account) internal view returns (bool) {
        return IController(_ensResolve(_controllerNode)).isAdmin(_account);
    }
}
