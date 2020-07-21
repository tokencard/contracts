/**
 *  The Consumer Contract Wallet - Wallet Deployer Cache
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

import "./externals/upgradeability/UpgradeabilityProxy.sol";
import "./interfaces/IWallet.sol";
import "./internals/ensResolvable.sol";
import "./internals/controllable.sol";


/// @title IWalletCache interface describes a method for poping an already cached wallet
interface IWalletCache {
    function walletCachePop() external returns (address payable);
}


//// @title Wallet cache with wallet pre-caching functionality.
contract WalletCache is ENSResolvable, Controllable {
    event CachedWallet(address payable _wallet);

    /*****   Constants   *****/
    // Default values for mainnet ENS
    // licence.tokencard.eth
    bytes32 private constant _DEFAULT_LICENCE_NODE = 0xd0ff8bd67f6e25e4e4b010df582a36a0ee9b78e49afe6cc1cff5dd5a83040330;
    // token-whitelist.tokencard.eth
    bytes32 private constant _DEFAULT_TOKEN_WHITELIST_NODE = 0xe84f90570f13fe09f288f2411ff9cf50da611ed0c7db7f73d48053ffc974d396;
    // wallet-deployer.v3.tokencard.eth
    bytes32 private constant _DEFAULT_WALLET_DEPLOYER_NODE = 0x1d0c0adbe6addd93659446311e0767a56b67d41ef38f0cb66dcf7560d28a5a38;

    bytes32 public licenceNode = _DEFAULT_LICENCE_NODE;
    bytes32 public tokenWhitelistNode = _DEFAULT_TOKEN_WHITELIST_NODE;
    bytes32 public walletDeployerNode = _DEFAULT_WALLET_DEPLOYER_NODE;

    address public walletImplementation;
    uint256 public defaultSpendLimit;

    address payable[] public cachedWallets;

    /// @notice parameters are passed in so that they can be used to construct new instances of the wallet
    /// @dev pass in bytes32 to use the default, production node labels for ENS
    constructor(
        address _walletImplementation_,
        address _ens_,
        uint256 _defaultSpendLimit_,
        bytes32 _controllerNode_,
        bytes32 _licenceNode_,
        bytes32 _tokenWhitelistNode_,
        bytes32 _walletDeployerNode_
    ) public {
        _initializeENSResolvable(_ens_);
        _initializeControllable(_controllerNode_);
        walletImplementation = _walletImplementation_;
        defaultSpendLimit = _defaultSpendLimit_;

        // Set licenceNode or use default
        if (_licenceNode_ != bytes32(0)) {
            licenceNode = _licenceNode_;
        }
        // Set tokenWhitelistNode or use default
        if (_tokenWhitelistNode_ != bytes32(0)) {
            tokenWhitelistNode = _tokenWhitelistNode_;
        }
        // Set walletDeployerNode or use default
        if (_walletDeployerNode_ != bytes32(0)) {
            walletDeployerNode = _walletDeployerNode_;
        }
    }

    modifier onlyWalletDeployer() {
        require(msg.sender == _ensResolve(walletDeployerNode), "not called by wallet-deployer");
        _;
    }

    /// @notice returns the number of pre-cached wallets.
    function cachedWalletsCount() external view returns (uint256) {
        return cachedWallets.length;
    }

    /// @notice This public method allows only the wallet deployer to pop pre-cached wallets or create a new one in case there aren't any
    function walletCachePop() external onlyWalletDeployer returns (address payable) {
        if (cachedWallets.length < 1) {
            cacheWallet();
        }

        address payable wallet = cachedWallets[cachedWallets.length - 1];
        cachedWallets.pop();

        return wallet;
    }

    /// @notice This public method allows anyone to pre-cache wallets
    function cacheWallet() public {
        address walletDeployerAddress = _ensResolve(walletDeployerNode);
        address payable wallet = address(new UpgradeabilityProxy(walletImplementation, ""));
        IWallet(wallet).initializeWallet(
            address(uint160(walletDeployerAddress)), // the address(uint160()) cast is done as the Wallet owner (1st argument) needs to be payable
            true,
            ensRegistry(),
            tokenWhitelistNode,
            controllerNode(),
            licenceNode,
            defaultSpendLimit
        );
        cachedWallets.push(wallet);

        emit CachedWallet(wallet);
    }
}
