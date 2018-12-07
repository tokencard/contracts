/**
 *  The Consumer Contract Wallet
 *  Copyright (C) 2018 Token Group Ltd
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

import "./controllable.sol";


/// @title OraclizeResolver implements on-chain contract address resolution.
contract OraclizeResolver is Controllable {

    address private _address;

    /// @dev Constructor initializes the resolver target and sets up controller authentication.
    /// @param _target initial address of the target contract.
    /// @param _ens is the address of the ENS registry.
    /// @param _controllerName is the registered ENS name for the controller contract.
    constructor(address _target, address _ens, bytes32 _controllerName) Controllable(_ens, _controllerName) public {
        _address = _target;
    }

    /// @return the address of the target contract.
    /// @notice this function is compatible with the oraclize contract API.
    function getAddress() public view returns (address) {
        return _address;
    }

    /// @dev Sets the address of the target contract.
    /// @param _target new address of the target contract.
    function setAddress(address _target) public onlyController {
        _address = _target;
    }
}
