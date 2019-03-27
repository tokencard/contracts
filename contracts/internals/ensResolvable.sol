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

import "../externals/ens/ENS.sol";
import "../externals/ens/PublicResolver.sol";

///@title ENSResolvable - Ethereum Name Service Resolver helper
///@dev This contract should be used to get addresses for ENS things
contract ENSResolvable {
    /// @dev _ens points to the ENS registry smart contract.
    ENS private _ens;

    address private _ensRegistry;

    constructor(address _ensReg) internal {
      _ensRegistry = _ensReg;
      _ens = ENS(_ensRegistry);
    }

    function _ensResolve(bytes32 _nodeHash) internal view returns (address) {
        return PublicResolver(_ens.resolver(_nodeHash)).addr(_nodeHash);
    }

    function ensRegistry() public view returns (address) {
        return _ensRegistry;
    }

}
