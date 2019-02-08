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

import "./tokenWhitelist.sol";
import "../externals/ens/ENS.sol";
import "./resolver.sol";


/// @title Controllable implements access control functionality based on a controller set in ENS.
contract TokenWhitelistable {
    /// @dev _ENS points to the ENS registry smart contract.
    ENS private _ENS;
    /// @dev Is the registered ENS name of the controller contract.
    bytes32 private _node;

    /// @dev Constructor initializes the controller contract object.
    /// @param _ens is the address of the ENS.
    /// @param _tokenWhitelistName is the ENS name of the Controller.
    constructor(address _ens, bytes32 _tokenWhitelistName) internal {
        _ENS = ENS(_ens);
        _node = _tokenWhitelistName;
    }

    function _getTokenInfo(address _a) internal view returns (string _symbol, uint256 _magnitude, uint256 _rate, bool _available, bool _loadable, uint256 _lastUpdate) {
        return ITokenWhitelist(IResolver(_ENS.resolver(_node)).addr(_node)).getTokenInfo(_a);
    }

    function _getTokenAddressArray() internal view returns (address[] _tokenAddresses) {
        return ITokenWhitelist(IResolver(_ENS.resolver(_node)).addr(_node)).getTokenAddressArray();
    }

    function _updateTokenRate(address _token, uint _rate, uint _updateDate) internal {
        ITokenWhitelist(IResolver(_ENS.resolver(_node)).addr(_node)).updateTokenRate(_token, _rate, _updateDate);
    }

    function _isTokenLoadable(address _a) internal view returns (bool _loadable) {
        ( , , , , bool loadable, ) = _getTokenInfo(_a);
        return loadable;
    }

    function _isTokenAvailable(address _a) internal view returns (bool _exists) {
        ( , , , bool available, , ) = _getTokenInfo(_a);
        return available;
    }
}
