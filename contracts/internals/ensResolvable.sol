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
 *
 *  SPDX-License-Identifier: GPL-3.0
 */

pragma solidity ^0.5.17;

import "../externals/ens/ENS.sol";
import "../externals/ens/PublicResolver.sol";


///@title Provides an interface for the ENS registry and the corresponding Resolver contract.
///@dev It can be inherited by other contracts which want to make use of ENS.
contract ENSResolvable {
    /// @notice Emits the address of the new ENS registry when it is set.
    event ENSSetRegistry(address _ensRegistry);

    /// @notice Address of the ENS registry contract set to the default ENS registry address.
    address private _ensRegistry = 0x00000000000C2E074eC69A0dFb2997BA6C7d2e1e;

    /// @return Current address of the ENS registry contract.
    function ensRegistry() external view returns (address) {
        return _ensRegistry;
    }

    /// @notice Sets the ENS registry address to the provided address.
    /// @param _ensRegistry_ Address of the ENS registry smart contract.
    function _ensSetRegistry(address _ensRegistry_) internal {
        _ensRegistry = _ensRegistry_;
        emit ENSSetRegistry(_ensRegistry_);
    }

    /// @notice Resolves the address of an ENS node.
    /// @param _node The ENS node to be resolved.
    /// @return Address of the resolved ENS node.
    function _ensResolve(bytes32 _node) internal view returns (address) {
        return PublicResolver(ENS(_ensRegistry).resolver(_node)).addr(_node);
    }
}
