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

pragma solidity ^0.6.0;

import "./ensResolvable.sol";
import "../interfaces/ITokenWhitelist.sol";


/// @title TokenWhitelistable implements access to the TokenWhitelist located behind ENS.
abstract contract TokenWhitelistable is ENSResolvable {
    /// @notice Is the registered ENS node identifying the tokenWhitelist contract
    bytes32 private _tokenWhitelistNode;

    /// @notice Constructor initializes the TokenWhitelistable object.
    /// @param _tokenWhitelistNode_ is the ENS node of the TokenWhitelist.
    constructor(bytes32 _tokenWhitelistNode_) internal {
        _tokenWhitelistNode = _tokenWhitelistNode_;
    }

    /// @notice This shows what TokenWhitelist is being used
    /// @return TokenWhitelist's node registered in ENS.
    function tokenWhitelistNode() external view returns (bytes32) {
        return _tokenWhitelistNode;
    }

    /// @notice This returns all of the fields for a given token.
    /// @param _a is the address of a given token.
    /// @return string of the token's symbol.
    /// @return uint of the token's magnitude.
    /// @return uint of the token's exchange rate to ETH.
    /// @return bool whether the token is available.
    /// @return bool whether the token is loadable to the TokenCard.
    /// @return bool whether the token is redeemable to the TKN Holder Contract.
    /// @return uint of the lastUpdated time of the token's exchange rate.
    function _getTokenInfo(address _a)
        internal
        view
        returns (
            string memory,
            uint256,
            uint256,
            bool,
            bool,
            bool,
            uint256
        )
    {
        return ITokenWhitelist(_ensResolve(_tokenWhitelistNode)).getTokenInfo(_a);
    }

    /// @notice This returns all of the fields for our stablecoin token.
    /// @return string of the token's symbol.
    /// @return uint of the token's magnitude.
    /// @return uint of the token's exchange rate to ETH.
    /// @return bool whether the token is available.
    /// @return bool whether the token is loadable to the TokenCard.
    /// @return bool whether the token is redeemable to the TKN Holder Contract.
    /// @return uint of the lastUpdated time of the token's exchange rate.
    function _getStablecoinInfo()
        internal
        view
        returns (
            string memory,
            uint256,
            uint256,
            bool,
            bool,
            bool,
            uint256
        )
    {
        return ITokenWhitelist(_ensResolve(_tokenWhitelistNode)).getStablecoinInfo();
    }

    /// @notice This returns an array of our whitelisted addresses.
    /// @return address[] of our whitelisted tokens.
    function _tokenAddressArray() internal view returns (address[] memory) {
        return ITokenWhitelist(_ensResolve(_tokenWhitelistNode)).tokenAddressArray();
    }

    /// @notice This returns an array of all redeemable token addresses.
    /// @return address[] of redeemable tokens.
    function _redeemableTokens() internal view returns (address[] memory) {
        return ITokenWhitelist(_ensResolve(_tokenWhitelistNode)).redeemableTokens();
    }

    /// @notice Update ERC20 token exchange rate.
    /// @param _token ERC20 token contract address.
    /// @param _rate ERC20 token exchange rate in wei.
    /// @param _updateDate date for the token updates. This will be compared to when oracle updates are received.
    function _updateTokenRate(
        address _token,
        uint256 _rate,
        uint256 _updateDate
    ) internal {
        ITokenWhitelist(_ensResolve(_tokenWhitelistNode)).updateTokenRate(_token, _rate, _updateDate);
    }

    /// @notice based on the method it returns the recipient address and amount/value, ERC20 specific.
    /// @param _data is the transaction payload.
    function _getERC20RecipientAndAmount(address _destination, bytes memory _data) internal view returns (address, uint256) {
        return ITokenWhitelist(_ensResolve(_tokenWhitelistNode)).getERC20RecipientAndAmount(_destination, _data);
    }

    /// @notice Checks whether a token is available.
    /// @return bool available or not.
    function _isTokenAvailable(address _a) internal view returns (bool) {
        (, , , bool available, , , ) = _getTokenInfo(_a);
        return available;
    }

    /// @notice Checks whether a token is redeemable.
    /// @return bool redeemable or not.
    function _isTokenRedeemable(address _a) internal view returns (bool) {
        (, , , , , bool redeemable, ) = _getTokenInfo(_a);
        return redeemable;
    }

    /// @notice Checks whether a token is loadable.
    /// @return bool loadable or not.
    function _isTokenLoadable(address _a) internal view returns (bool) {
        (, , , , bool loadable, , ) = _getTokenInfo(_a);
        return loadable;
    }

    /// @notice This gets the address of the stablecoin.
    /// @return the address of the stablecoin contract.
    function _stablecoin() internal view returns (address) {
        return ITokenWhitelist(_ensResolve(_tokenWhitelistNode)).stablecoin();
    }
}
