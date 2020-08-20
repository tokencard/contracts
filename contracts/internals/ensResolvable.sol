/**
 *  ENSResolvable - The Consumer Contract Wallet
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

import "../externals/initializable.sol";
import "../interfaces/IENS.sol";
import "../interfaces/IPublicResolver.sol";

///@title ENSResolvable - Ethereum Name Service Resolver
///@notice contract should be used to get an address for an ENS node
contract ENSResolvable is Initializable {
    /// @dev Address of the ENS registry contract set to the default ENS registry address.
    address private _ensRegistry = address(0x00000000000C2E074eC69A0dFb2997BA6C7d2e1e);

    /// @notice Checks if the contract has been initialized succesfully i.e. the ENS registry has been set.
    modifier initialized() {
        require(_ensRegistry != address(0), "ENSResolvable not initialized");
        _;
    }

    /// @return Current address of the ENS registry contract.
    function ensRegistry() public view returns (address) {
        return _ensRegistry;
    }

    /// @notice Helper function used to get the address of a node.
    /// @param _node of the ENS entry that needs resolving.
    /// @return The address of the resolved ENS node.
    function _ensResolve(bytes32 _node) internal view initialized returns (address) {
        return IPublicResolver(IENS(_ensRegistry).resolver(_node)).addr(_node);
    }

    /// @param _ensReg is the ENS registry used.
    function _initializeENSResolvable(address _ensReg) internal initializer {
        // Set ENS registry or use default
        if (_ensReg != address(0)) {
            _ensRegistry = _ensReg;
        }
    }
}
