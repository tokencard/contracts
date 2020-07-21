/**
 *  Holder (aka Asset Contract) - The Consumer Contract Wallet
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

import "./externals/ERC20.sol";
import "./externals/SafeMath.sol";
import "./internals/transferrable.sol";
import "./internals/balanceable.sol";
import "./internals/burner.sol";
import "./internals/controllable.sol";
import "./internals/tokenWhitelistable.sol";

/// @title Holder - The TKN Asset Contract
/// @notice When the TKN contract calls the burn method, a share of the tokens held by this contract are disbursed to the burner.
contract Holder is Balanceable, ENSResolvable, Controllable, Transferrable, TokenWhitelistable {
    using SafeMath for uint256;

    event Received(address _from, uint256 _amount);
    event CashAndBurned(address _to, address _asset, uint256 _amount);
    event Claimed(address _to, address _asset, uint256 _amount);

    /// @dev Check if the sender is the burner contract
    modifier onlyBurner() {
        require(msg.sender == _burner, "burner contract is not the sender");
        _;
    }

    // Burner token which can be burned to redeem shares.
    address private _burner;

    /// @notice Constructor initializes the holder contract.
    /// @param _burnerContract_ is the address of the token contract TKN with burning functionality.
    /// @param _ens_ is the address of the ENS registry.
    /// @param _tokenWhitelistNode_ is the ENS node of the Token whitelist.
    /// @param _controllerNode_ is the ENS node of the Controller
    constructor(
        address _burnerContract_,
        address _ens_,
        bytes32 _tokenWhitelistNode_,
        bytes32 _controllerNode_
    ) public {
        _initializeENSResolvable(_ens_);
        _initializeControllable(_controllerNode_);
        _initializeTokenWhitelistable(_tokenWhitelistNode_);
        _burner = _burnerContract_;
    }

    /// @notice Ether may be sent from anywhere.
    receive() external payable {
        emit Received(msg.sender, msg.value);
    }

    /// @notice Burn handles disbursing a share of tokens in this contract to a given address.
    /// @param _to The address to disburse to
    /// @param _amount The amount of TKN that will be burned if this succeeds
    function burn(address payable _to, uint256 _amount) external onlyBurner returns (bool) {
        if (_amount == 0) {
            return true;
        }
        // The burner token deducts from the supply before calling.
        uint256 supply = IBurner(_burner).currentSupply().add(_amount);
        address[] memory redeemableAddresses = _redeemableTokens();
        for (uint256 i = 0; i < redeemableAddresses.length; i++) {
            uint256 redeemableBalance = _balance(redeemableAddresses[i]);
            if (redeemableBalance > 0) {
                uint256 redeemableAmount = redeemableBalance.mul(_amount).div(supply);
                _safeTransfer(_to, redeemableAddresses[i], redeemableAmount);
                emit CashAndBurned(_to, redeemableAddresses[i], redeemableAmount);
            }
        }

        return true;
    }

    /// @notice This allows for the admin to reclaim the non-redeemableTokens.
    /// @param _to this is the address which the reclaimed tokens will be sent to.
    /// @param _nonRedeemableAddresses this is the array of tokens to be claimed.
    function nonRedeemableTokenClaim(address payable _to, address[] calldata _nonRedeemableAddresses) external onlyAdmin returns (bool) {
        for (uint256 i = 0; i < _nonRedeemableAddresses.length; i++) {
            //revert if token is redeemable
            require(!_isTokenRedeemable(_nonRedeemableAddresses[i]), "redeemables cannot be claimed");
            uint256 claimBalance = _balance(_nonRedeemableAddresses[i]);
            if (claimBalance > 0) {
                _safeTransfer(_to, _nonRedeemableAddresses[i], claimBalance);
                emit Claimed(_to, _nonRedeemableAddresses[i], claimBalance);
            }
        }

        return true;
    }

    /// @notice Returned the address of the burner contract.
    /// @return the TKN address.
    function burner() external view returns (address) {
        return _burner;
    }
}
