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

pragma solidity ^0.6.0;

import "../interfaces/IENS.sol";
import "../interfaces/IPublicResolver.sol";


///@title ENSResolvable - Ethereum Name Service Resolver
///@notice contract should be used to get an address for an ENS node
contract ENSResolvable {

    /// @notice _ensRegistry points to the ENS registry smart contract.
    address private _ensRegistry;

    /// @param _ensRegistry_ is the ENS registry used
    constructor(address _ensRegistry_) internal {
        _ensRegistry = _ensRegistry_;
    }

    /// @notice this is used to that one can observe which ENS registry is being used
    function ensRegistry() external view returns (address) {
        return _ensRegistry;
    }

    /// @notice helper function used to get the address of a node
    /// @param _node of the ENS entry that needs resolving
    /// @return the address of the said node
    function _ensResolve(bytes32 _node) internal view returns (address) {
        return IPublicResolver(IENS(_ensRegistry).resolver(_node)).addr(_node);
    }
}
