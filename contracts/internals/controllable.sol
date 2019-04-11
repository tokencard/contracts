/**
 *  Controller Interface
 *  Copyright (C) 2018 The Contract Wallet Company Limited
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
import "../externals/ens/ENS.sol";

/// @title Resolver returns the controller contract address.
interface IResolver {
    function addr(bytes32) external view returns (address);
}

/// @title Controllable implements access control functionality based on a controller set in ENS.
contract Controllable {
    /// @dev _ENS points to the ENS registry smart contract.
    ENS private _ENS;
    /// @dev Is the registered ENS name of the controller contract.
    bytes32 private _node;

    /// @dev Constructor initializes the controller contract object.
    /// @param _ens is the address of the ENS.
    /// @param _controllerName is the ENS name of the Controller.
    constructor(address _ens, bytes32 _controllerName) internal {
      _ENS = ENS(_ens);
      _node = _controllerName;
    }

    /// @dev Checks if message sender is the controller.
    modifier onlyController() {
        require(_isController(msg.sender), "sender is not a controller");
        _;
    }

    /// @return true if the provided account is the controller.
    function _isController(address _account) internal view returns (bool) {
        return IController(IResolver(_ENS.resolver(_node)).addr(_node)).isController(_account);
    }
}
