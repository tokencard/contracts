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

    event CachedWallet(address wallet);
    event DeployedWallet(address wallet, address owner);
    event MigratedWallet(address wallet, address owner);

    mapping(address => address) public deployed;
    address[] public cached;

    address public ens;
    bytes32 public controllerName;
    bytes32 public oracleName;
    uint public spendLimit;

    constructor(address _ens, bytes32 _oracleName, bytes32 _controllerName, uint _spendLimit) Controllable(_ens, _controllerName) public {
        ens = _ens;
        controllerName = _controllerName;
        oracleName = _oracleName;
        spendLimit = _spendLimit;
    }

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

    function migrateWallet(address owner, uint _spendLimit, address[] _whitelistedAddresses) external onlyController {
        if (cached.length < 1) {
            cacheWallet();
    	}
        address walletAddress = cached[cached.length-1];
        cached.length--;
        deployed[owner] = walletAddress;
        Wallet(walletAddress).initializeSpendLimit(_spendLimit);
        Wallet(walletAddress).initializeWhitelist(_whitelistedAddresses);
        Wallet(walletAddress).transferOwnership(owner);
        emit MigratedWallet(walletAddress, owner);
    }

    function cachedWalletCount() external view returns (uint) {
        return cached.length;
    }

    function cacheWallet() public {
        address walletAddress = new Wallet(address(this), true, ens, oracleName, controllerName, spendLimit);
        cached.push(walletAddress);
        emit CachedWallet(walletAddress);
    }


}
