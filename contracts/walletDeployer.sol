/**
 *  The Consumer Contract Wallet Deployer
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

    bytes32 public constant controllerNode = 0x7f2ce995617d2816b426c5c8698c5ec2952f7a34bb10f38326f74933d5893697;

    mapping(address => address) public deployedWallets;

    address public walletCache;

    /// @notice it needs to know to address of the wallet cache
    constructor(address _ens, address _walletCache) ENSResolvable(_ens) Controllable(controllerNode) public {
        walletCache = _walletCache;
    }

    /// @notice This function is used to deploy a Wallet for a given owner address
    /// @param _owner is the owner address for the new Wallet to be
    function deployWallet(address payable _owner) external onlyController {
        Wallet wallet = IWalletCache(walletCache).walletCachePop();
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
        Wallet wallet = IWalletCache(walletCache).walletCachePop();

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
