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

pragma solidity 0.5.15;

import "../externals/ens/ENS.sol";
import "../externals/initializable.sol";
import "../externals/ens/PublicResolver.sol";


///@title ENSResolvable - Ethereum Name Service Resolver
///@notice contract should be used to get an address for an ENS node
contract ENSResolvable is Initializable {
    /// @notice _ensRegistry points to the ENS registry smart contract.
    address private _ensRegistry;

    /// @notice Checks if the contract has been initialized succesfully i.e. the ENS registry has been set.
    modifier initialized() {
        require(_ensRegistry != address(0), "ENSResolvable not initialized");
        _;
    }

    /// @notice this is used to that one can observe which ENS registry is being used
    function ensRegistry() external view returns (address) {
        return _ensRegistry;
    }

    /// @notice helper function used to get the address of a node
    /// @param _node of the ENS entry that needs resolving
    /// @return the address of the said node
    function _ensResolve(bytes32 _node) internal view initialized returns (address) {
        return PublicResolver(ENS(_ensRegistry).resolver(_node)).addr(_node);
    }

    /// @param _ensReg is the ENS registry used
    function _initializeENSResolvable(address _ensReg) internal initializer {
        require(_ensReg != address(0), "ensReg is 0");
        _ensRegistry = _ensReg;
    }
}
