/**
 *  The Consumer Contract Wallet - Wallet Deployer
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

pragma solidity ^0.5.10;

import "./wallet.sol";
import "./walletCache.sol";
import "./internals/controllable.sol";

//// @title Wallet deployer with pre-caching if wallets functionality.
contract WalletDeployer is ENSResolvable, Controllable {

    event DeployedWallet(Wallet _wallet, address _owner);
    event MigratedWallet(Wallet _wallet, Wallet _oldWallet, address _owner);

    // CONTROLLER_NODE = controller.tokencard.eth
    bytes32 public constant CONTROLLER_NODE = 0x7f2ce995617d2816b426c5c8698c5ec2952f7a34bb10f38326f74933d5893697;
    // walletCacheNode = wallet-cache.tokencard.eth
    bytes32 public constant WALLET_CACHE_NODE = 0x7eee9c3927d17f70ce19de05f73d05dbda3449d450ba9a4c64f24c24bfb9d7ac;

    mapping(address => address) public deployedWallets;

    /// @notice it needs to know to address of the wallet cache
    constructor(address _ens) ENSResolvable(_ens) Controllable(CONTROLLER_NODE) public {
    }

    /// @notice This function is used to deploy a Wallet for a given owner address
    /// @param _owner is the owner address for the new Wallet to be
    function deployWallet(address payable _owner) external onlyController {
        Wallet wallet = IWalletCache(_ensResolve(WALLET_CACHE_NODE)).walletCachePop();
        // This sets the changeableOwner boolean to false, i.e. no more change ownership
        wallet.transferOwnership(_owner, false);
        deployedWallets[_owner] = address(wallet);

        emit DeployedWallet(wallet, _owner);
    }

    /// @notice This function is used to migrate an owner's security settings from a previous version of the wallet
    /// @param _owner is the owner address for the new Wallet to be
    /// @param _spendLimit is the user's set daily spend limit
    /// @param _gasTopUpLimit is the user's set daily gas top-up limit
    /// @param _whitelistedAddresses is the set of the user's whitelisted addresses
    function migrateWallet(address payable _owner, Wallet _oldWallet, bool _initializedSpendLimit, bool _initializedGasTopUpLimit, bool _initializedWhitelist, uint _spendLimit, uint _gasTopUpLimit, address[] calldata _whitelistedAddresses) external onlyController {
        Wallet wallet = IWalletCache(_ensResolve(WALLET_CACHE_NODE)).walletCachePop();

        // Sets up the security settings as per the _oldWallet 
        if (_initializedSpendLimit) {
            wallet.setSpendLimit(_spendLimit);
        }
        if (_initializedGasTopUpLimit) {
            wallet.setGasTopUpLimit(_gasTopUpLimit);
        }
        if (_initializedWhitelist) {
            wallet.setWhitelist(_whitelistedAddresses);
        }

        wallet.transferOwnership(_owner, false);
        deployedWallets[_owner] = address(wallet);

        emit MigratedWallet(wallet, _oldWallet, _owner);
    }

}
