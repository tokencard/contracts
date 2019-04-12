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
import "../internals/ensResolvable.sol";


/// @title Controllable implements access control functionality based on a controller set in ENS.
contract TokenWhitelistable is ENSResolvable {
    /// @dev Is the registered ENS node identifying the tokenWhitelist contract.
    bytes32 private _tokenWhitelistNode;

    /// @dev Constructor initializes the controller contract object.
    /// @param _tokenWhitelistNameHash is the ENS name hash of the Controller.
    constructor(bytes32 _tokenWhitelistNameHash) internal {
        _tokenWhitelistNode = _tokenWhitelistNameHash;
    }

    function _getTokenInfo(address _a) internal view returns (string, uint256, uint256, bool, bool, bool, uint256) {
        return ITokenWhitelist(_ensResolve(_tokenWhitelistNode)).getTokenInfo(_a);
    }

    function _getStablecoinInfo() internal view returns (string, uint256, uint256, bool, bool, bool, uint256) {
        return ITokenWhitelist(_ensResolve(_tokenWhitelistNode)).getStablecoinInfo();
    }

    function _getTokenAddressArray() internal view returns (address[]) {
        return ITokenWhitelist(_ensResolve(_tokenWhitelistNode)).getTokenAddressArray();
    }

    function _updateTokenRate(address _token, uint _rate, uint _updateDate) internal {
        ITokenWhitelist(_ensResolve(_tokenWhitelistNode)).updateTokenRate(_token, _rate, _updateDate);
    }

    function _isTokenLoadable(address _a) internal view returns (bool) {
        ( , , , , bool loadable, , ) = _getTokenInfo(_a);
        return loadable;
    }

    function _isTokenBurnable(address _a) internal view returns (bool) {
        ( , , , , , bool burnable, ) = _getTokenInfo(_a);
        return burnable;
    }

    function _isTokenAvailable(address _a) internal view returns (bool) {
        ( , , , bool available, , , ) = _getTokenInfo(_a);
        return available;
    }

    /// @return the address of the stablecoin contract.
    function _stablecoin() internal view returns (address) {
        return ITokenWhitelist(_ensResolve(_tokenWhitelistNode)).stablecoin();
    }

    /// @return the maximum card loadLimit.
    function _maxStablecoinLoadLimit() internal view returns (uint) {
        return ITokenWhitelist(_ensResolve(_tokenWhitelistNode)).maxStablecoinLoadLimit();
    }

    /// @dev updates the maximum card loadLimit.
    function _updateMaxStablecoinLoadLimit(uint newLimit) internal {
        ITokenWhitelist(_ensResolve(_tokenWhitelistNode)).updateMaxStablecoinLoadLimit(newLimit);
    }

    /// @return the token whitelist node registered in ENS.
    function tokenWhitelistNode() public view returns (bytes32) {
        return _tokenWhitelistNode;
    }
}
