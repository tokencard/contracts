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

pragma solidity ^0.6.0;

import "./internals/controllable.sol";
import "./wallet.sol";
import "./walletCache.sol";


//// @title Wallet deployer with pre-caching if wallets functionality.
contract WalletDeployer is ENSResolvable, Controllable {
    event DeployedWallet(Wallet _wallet, address _owner);
    event MigratedWallet(Wallet _wallet, Wallet _oldWallet, address _owner, uint256 _paid);

    /*****   Constants   *****/
    // Default values for mainnet ENS
    // wallet-cache.v3.tokencard.eth
    bytes32 private constant _DEFAULT_WALLET_CACHE_NODE = 0xaf553cb0d77690819f9d6fbaa04416e1fdcfa01b2a9a833c7a11e6ae0bc1be88;
    bytes32 public walletCacheNode = _DEFAULT_WALLET_CACHE_NODE;

    mapping(address => address) public deployedWallets;

    /// @notice it needs to know to address of the wallet cache

    constructor(
        address _ens_,
        bytes32 _controllerNode_,
        bytes32 _walletCacheNode_
    ) public ENSResolvable(_ens_) Controllable(_controllerNode_) {
        // Set walletCacheNode or use default
        if (_walletCacheNode_ != bytes32(0)) {
            walletCacheNode = _walletCacheNode_;
        }
    }

    /// @notice This function is used to deploy a Wallet for a given owner address
    /// @param _owner is the owner address for the new Wallet to be
    function deployWallet(address payable _owner) external onlyController {
        Wallet wallet = IWalletCache(_ensResolve(walletCacheNode)).walletCachePop();
        emit DeployedWallet(wallet, _owner);

        deployedWallets[_owner] = address(wallet);

        // This sets the changeableOwner boolean to false, i.e. no more change ownership
        wallet.transferOwnership(_owner, false);
    }

    /// @notice This function is used to migrate an owner's security settings from a previous version of the wallet
    /// @param _owner is the owner address for the new Wallet to be
    /// @param _spendLimit is the user's set daily spend limit
    /// @param _gasTopUpLimit is the user's set daily gas top-up limit
    /// @param _whitelistedAddresses is the set of the user's whitelisted addresses
    function migrateWallet(
        address payable _owner,
        Wallet _oldWallet,
        bool _initializedSpendLimit,
        bool _initializedGasTopUpLimit,
        bool _initializedLoadLimit,
        bool _initializedWhitelist,
        uint256 _spendLimit,
        uint256 _gasTopUpLimit,
        uint256 _loadLimit,
        address[] calldata _whitelistedAddresses
    ) external payable onlyController {
        require(deployedWallets[_owner] == address(0x0), "wallet already deployed for owner");
        require(_oldWallet.owner() == _owner, "owner mismatch");

        Wallet wallet = IWalletCache(_ensResolve(walletCacheNode)).walletCachePop();
        emit MigratedWallet(wallet, _oldWallet, _owner, msg.value);

        deployedWallets[_owner] = address(wallet);

        // Sets up the security settings as per the _oldWallet
        if (_initializedSpendLimit) {
            wallet.setSpendLimit(_spendLimit);
        }
        if (_initializedGasTopUpLimit) {
            wallet.setGasTopUpLimit(_gasTopUpLimit);
        }
        if (_initializedLoadLimit) {
            wallet.setLoadLimit(_loadLimit);
        }
        if (_initializedWhitelist) {
            wallet.setWhitelist(_whitelistedAddresses);
        }

        wallet.transferOwnership(_owner, false);

        if (msg.value > 0) {
            _owner.transfer(msg.value);
        }
    }
}
