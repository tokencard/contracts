/**
 *  The Consumer Contract Wallet Deployer
 *  Copyright (C) 2018 The Contract Wallet Company Limited
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

//// @title Wallet deployer with pre-caching if wallets functionality.
contract WalletDeployer is Controllable {

    event WalletCached(address wallet);
    event WalletDeployed(address wallet, address owner);
    event WalletMigrated(address wallet, address newOwner);

    mapping(address => address) public deployed;
    address[] public cached;

    address private ens;
    bytes32 private controllerName;
    bytes32 private oracleName;
    uint private spendLimit;

    constructor(address _ens, bytes32 _oracleName, bytes32 _controllerName, uint _spendLimit) Controllable(_ens, _controllerName) public {
        ens = _ens;
        controllerName = _controllerName;
        oracleName = _oracleName;
        spendLimit = _spendLimit;
    }

    function cacheWallet() public {
        address walletAddress = new Wallet(address(this), true, ens, oracleName, controllerName, spendLimit);
        cached.push(walletAddress);
        emit WalletCached(walletAddress);
    }

    function deployWallet(address owner) external onlyController {
        if (cached.length < 1) {
            cacheWallet();
    	}
        address walletAddress = cached[cached.length-1];
        cached.length--;
        Wallet(walletAddress).transferOwnership(owner);
        deployed[owner] = walletAddress;
        emit WalletDeployed(walletAddress, owner);
    }

    function migrateWallet(address newOwner, uint _spendLimit, address[] _whitelistedAddresses) external onlyController {
        if (cached.length < 1) {
            cacheWallet();
    	}
        address walletAddress = cached[cached.length-1];
        cached.length--;
        deployed[newOwner] = walletAddress;
        IWallet(walletAddress).initializeWhitelist(_whitelistedAddresses);
        IWallet(walletAddress).initializeSpendLimit(_spendLimit);
        Wallet(walletAddress).transferOwnership(newOwner);
        emit WalletMigrated(walletAddress, newOwner);
    }

    function cachedContractCount() external view returns (uint) {
        return cached.length;
    }

}
