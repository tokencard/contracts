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

pragma solidity ^0.4.25;

import "./wallet.sol";
import "./internals/controllable.sol";


/// @title Wallet Deployer
/// @notice Factory method to create deploy a new wallet
/// @notice Allows for pre-fabrication / caching of wallets
contract WalletDeployer is Controllable {

    event CachedWallet(address wallet);
    event DeployedWallet(address wallet, address owner);
    event MigratedWallet(address wallet, address owner);

    mapping(address => address) public deployed;
    address[] public cached;

    address public ens;
    bytes32 public controllerName;
    bytes32 public oracleName;
    uint public defaultSpendLimit;

    /// @notice parameters are passed in so that they can be used to construct new instances of the wallet
    constructor(address _ens, bytes32 _oracleName, bytes32 _controllerName, uint _spendLimit) Controllable(_ens, _controllerName) public {
        ens = _ens;
        controllerName = _controllerName;
        oracleName = _oracleName;
        defaultSpendLimit = _spendLimit;
    }

    /// @notice This function is used to deploy a Wallet for a given owner address
    /// @param owner is the owner address for the new Wallet to be
    function deployWallet(address owner) external onlyController {
        if (cached.length < 1) {
            cacheWallet();
    	}

        address walletAddress = cached[cached.length-1];
        cached.length--;
        Wallet(walletAddress).transferOwnership(owner);
        deployed[owner] = walletAddress;

        emit DeployedWallet(walletAddress, owner);
    }

    /// @notice This function is used to migrate an owner's security settings from a previous version of the wallet
    /// @param owner is the owner address for the new Wallet to be
    /// @param _spendLimit is the user's set daily spend limit
    /// @param _topUpLimit is the user's set daily gas top-up limit
    /// @param _whitelistedAddresses is the set of the user's whitelisted addresses
    function migrateWallet(address owner, uint _spendLimit, uint _topUpLimit, bool _initializedWhitelist, address[] _whitelistedAddresses) external onlyController {
        if (cached.length < 1) {
            cacheWallet();
    	}

        address walletAddress = cached[cached.length-1];
        cached.length--;

        Wallet(walletAddress).initializeSpendLimit(_spendLimit);
        Wallet(walletAddress).initializeTopUpLimit(_topUpLimit);

        if (_initializedWhitelist) {
            Wallet(walletAddress).initializeWhitelist(_whitelistedAddresses);
        }

        Wallet(walletAddress).transferOwnership(owner);
        deployed[owner] = walletAddress;

        emit MigratedWallet(walletAddress, owner);
    }

    /// @notice returns the number of pre-cached wallets
    function cachedWalletCount() external view returns (uint) {
        return cached.length;
    }

    /// @notice This public method allows anyone to pre-cache wallets
    function cacheWallet() public {
        address walletAddress = new Wallet(address(this), true, ens, oracleName, controllerName, defaultSpendLimit);
        cached.push(walletAddress);
        emit CachedWallet(walletAddress);
    }
}
