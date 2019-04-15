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

pragma solidity ^0.4.25;

import "./controller.sol";
import "../internals/ensResolvable.sol";


/// @title Controllable implements access control functionality of the Controller found via ENS.
contract Controllable is ENSResolvable {
    /// @dev Is the registered ENS node identifying the controller contract.
    bytes32 private _controllerNode;

    /// @notice Constructor initializes the controller contract object.
    /// @param _controllerNameHash_ is the ENS name hash of the Controller.
    constructor(bytes32 _controllerNameHash_) internal {
        _controllerNode = _controllerNameHash_;
    }

    /// @notice Checks if message sender is the controller.
    modifier onlyController() {
        require(_isController(msg.sender), "sender is not a controller");
        _;
    }

    /// @return the controller name hash registered in ENS.
    function controllerNode() external view returns (bytes32) {
        return _controllerNode;
    }

    /// @return true if the provided account is the controller.
    function _isController(address _account) internal view returns (bool) {
        return IController(_ensResolve(_controllerNode)).isController(_account);
    }

}
