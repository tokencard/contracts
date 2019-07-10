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
import "./internals/ensResolvable.sol";

//// @title Wallet deployer with pre-caching if wallets functionality.
contract WalletDeployer is Controllable {

    event WalletCached(address wallet);
    event WalletDeployed(address wallet, address owner);

    mapping(address => address) public deployed;
    address[] public cached;

    address private ens;
    bytes32 private tokenWhitelistNameHash;
    bytes32 private controllerNameHash;
    bytes32 private licenceNameHash;
    uint private spendLimit;

    constructor(address _ens, bytes32 _tokenWhitelistNameHash, bytes32 _controllerNameHash, bytes32 _licenceNameHash, uint _spendLimit) ENSResolvable(_ens) Controllable(_controllerNameHash) public {
        ens = _ens;
        tokenWhitelistNameHash = _tokenWhitelistNameHash;
        controllerNameHash = _controllerNameHash;
        licenceNameHash = _licenceNameHash;
        spendLimit = _spendLimit;
    }

    function deployWallet(address owner) external onlyController {
        if (cached.length < 1) {
            cacheWallet();
        }
        address walletAddress = cached[cached.length-1];
        cached.length--;
        Wallet(walletAddress).transferOwnership(owner, false);
        deployed[owner] = walletAddress;
        emit WalletDeployed(walletAddress, owner);
    }

    function cacheWallet() public onlyController {
        address walletAddress = new Wallet(address(this), true, ens, tokenWhitelistNameHash, controllerNameHash, licenceNameHash, spendLimit);
        cached.push(walletAddress);
        emit WalletCached(walletAddress);
    }
}
