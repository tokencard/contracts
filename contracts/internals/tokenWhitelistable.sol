/**
 *  TokenWhitelistable - The Consumer Contract Wallet
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

import "./tokenWhitelist.sol";
import "../internals/ensResolvable.sol";


/// @title TokenWhitelistable implements access to the TokenWhitelist located behind ENS.
contract TokenWhitelistable is ENSResolvable {

    /// @notice Is the registered ENS node identifying the tokenWhitelist contract.
    bytes32 private _tokenWhitelistNode;

    /// @notice Constructor initializes the TokenWhitelistable object.
    /// @param _tokenWhitelistNameHash_ is the ENS name hash of the TokenWhitelist.
    constructor(bytes32 _tokenWhitelistNameHash_) internal {
        _tokenWhitelistNode = _tokenWhitelistNameHash_;
    }

    /// @notice This shows what TokenWhitelist is being used
    /// @return the TokenWhitelist's name hash registered in ENS.
    function tokenWhitelistNode() external view returns (bytes32) {
        return _tokenWhitelistNode;
    }

    /// @notice This returns all of the fields for a given token
    /// @param _a is the address of a given token
    /// @return string of the token's symbol
    /// @return uint of the token's magnitude
    /// @return uint of the token's exchange rate to ETH
    /// @return bool whether the token is available
    /// @return bool whether the token is loadable to the TokenCard
    /// @return bool whether the token is burnable to the TKN Holder Contract
    /// @return uint of the lastUpdated time of the token's exchange rate
    function _getTokenInfo(address _a) internal view returns (string, uint256, uint256, bool, bool, bool, uint256) {
        return ITokenWhitelist(_ensResolve(_tokenWhitelistNode)).getTokenInfo(_a);
    }

    /// @notice This returns all of the fields for our stablecoin token
    /// @return string of the token's symbol
    /// @return uint of the token's magnitude
    /// @return uint of the token's exchange rate to ETH
    /// @return bool whether the token is available
    /// @return bool whether the token is loadable to the TokenCard
    /// @return bool whether the token is burnable to the TKN Holder Contract
    /// @return uint of the lastUpdated time of the token's exchange rate
    function _getStablecoinInfo() internal view returns (string, uint256, uint256, bool, bool, bool, uint256) {
        return ITokenWhitelist(_ensResolve(_tokenWhitelistNode)).getStablecoinInfo();
    }

    /// @notice This returns an array of our whitelisted address
    /// @return address[] of our whitelisted tokens
    function _tokenAddressArray() internal view returns (address[]) {
        return ITokenWhitelist(_ensResolve(_tokenWhitelistNode)).tokenAddressArray();
    }

    /// @notice Update ERC20 token exchange rate.
    /// @param _token ERC20 token contract address.
    /// @param _rate ERC20 token exchange rate in wei.
    /// @param _updateDate date for the token updates. This will be compared to when oracle updates are received.
    function _updateTokenRate(address _token, uint _rate, uint _updateDate) internal {
        ITokenWhitelist(_ensResolve(_tokenWhitelistNode)).updateTokenRate(_token, _rate, _updateDate);
    }

    /// @notice Checks whether a token is available
    /// @return bool available or not
    function _isTokenAvailable(address _a) internal view returns (bool) {
        ( , , , bool available, , , ) = _getTokenInfo(_a);
        return available;
    }

    /// @notice Checks whether a token is burnable
    /// @return bool burnable or not
    function _isTokenBurnable(address _a) internal view returns (bool) {
        ( , , , , , bool burnable, ) = _getTokenInfo(_a);
        return burnable;
    }

    /// @notice Checks whether a token is loadable
    /// @return bool loadable or not
    function _isTokenLoadable(address _a) internal view returns (bool) {
        ( , , , , bool loadable, , ) = _getTokenInfo(_a);
        return loadable;
    }

    /// @notice This gets the address of the stablecoin
    /// @return the address of the stablecoin contract.
    function _stablecoin() internal view returns (address) {
        return ITokenWhitelist(_ensResolve(_tokenWhitelistNode)).stablecoin();
    }

}
