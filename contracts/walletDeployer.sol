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
import "./internals/controllable.sol";
import "./internals/ensResolvable.sol";

//// @title Wallet deployer with pre-caching if wallets functionality.
contract WalletDeployer is ENSResolvable, Controllable {

    event CachedWallet(Wallet _wallet);
    event DeployedWallet(Wallet _wallet, address _owner);
    event MigratedWallet(Wallet _wallet, Wallet _oldWallet, address _owner);

    /*****   Constants   *****/
    bytes32 public constant controllerNode = 0x7f2ce995617d2816b426c5c8698c5ec2952f7a34bb10f38326f74933d5893697;
    bytes32 public constant licenceNode = 0xd0ff8bd67f6e25e4e4b010df582a36a0ee9b78e49afe6cc1cff5dd5a83040330;
    bytes32 public constant tokenWhitelistNode = 0xe84f90570f13fe09f288f2411ff9cf50da611ed0c7db7f73d48053ffc974d396;

    mapping(address => address) public deployedWallets;
    Wallet[] public cachedWallets;

    address public ens;
    uint public defaultSpendLimit;

    /// @notice parameters are passed in so that they can be used to construct new instances of the wallet
    constructor(address _ens, uint _defaultSpendLimit) ENSResolvable(_ens) Controllable(controllerNode) public {
        ens = _ens;
        defaultSpendLimit = _defaultSpendLimit;
    }

    /// @dev This contract must be payable by anyone, as the Wallet Owner needs to be payable
    function() external payable {}

    /// @notice This public method allows anyone to pre-cache wallets
    function cacheWallet() public {
        Wallet wallet = new Wallet(address(this), true, ens, tokenWhitelistNode, controllerNode, licenceNode, defaultSpendLimit);
        cachedWallets.push(wallet);
        emit CachedWallet(wallet);
    }

    /// @notice This function is used to deploy a Wallet for a given owner address
    /// @param _owner is the owner address for the new Wallet to be
    function deployWallet(address payable _owner) external onlyController {
        if (cachedWallets.length < 1) {
            cacheWallet();
        }
        Wallet wallet = cachedWallets[cachedWallets.length-1];
        cachedWallets.pop();
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
        if (cachedWallets.length < 1) {
            cacheWallet();
        }

        Wallet  wallet = cachedWallets[cachedWallets.length-1];
        cachedWallets.pop();

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

    /// @notice returns the number of pre-cached wallets
    function cachedWalletsCount() external view returns (uint) {
        return cachedWallets.length;
    }

}
