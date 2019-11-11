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

pragma solidity ^0.5.10;

import "./wallet.sol";
import "./internals/ensResolvable.sol";


/// @title IWalletCache interface describes a method for poping an already cached wallet
interface IWalletCache {
    function walletCachePop() external returns(Wallet);
}

//// @title Wallet cache with wallet pre-caching functionality.
contract WalletCache is ENSResolvable {

    event CachedWallet(Wallet _wallet);

    /*****   Constants   *****/
    // CONTROLLER_NODE = controller.tokencard.eth
    bytes32 public constant CONTROLLER_NODE = 0x7f2ce995617d2816b426c5c8698c5ec2952f7a34bb10f38326f74933d5893697;
    // LICENCE_NODE = licence.tokencard.eth
    bytes32 public constant LICENCE_NODE = 0xd0ff8bd67f6e25e4e4b010df582a36a0ee9b78e49afe6cc1cff5dd5a83040330;
    // TOKEN_WHITELIST_NODE = token-whitelist.tokencard.eth
    bytes32 public constant TOKEN_WHITELIST_NODE = 0xe84f90570f13fe09f288f2411ff9cf50da611ed0c7db7f73d48053ffc974d396;
    // WALLET_DEPLOYER_NODE = wallet-deployer.tokencard.eth
    bytes32 public constant WALLET_DEPLOYER_NODE = 0xd21a61ac2e4de1319ef7c76dd03046ec2e67a92cfc9efb7c28f79a4c323a5b80;

    Wallet[] public cachedWallets;

    address public ens;
    uint public defaultSpendLimit;

    /// @notice parameters are passed in so that they can be used to construct new instances of the wallet
    constructor(address _ens, uint _defaultSpendLimit) ENSResolvable(_ens) public {
        ens = _ens;
        defaultSpendLimit = _defaultSpendLimit;
    }

    modifier onlyWalletDeployer() {
        require(msg.sender == _ensResolve(WALLET_DEPLOYER_NODE), "not called by wallet-deployer");
        _;
    }

    /// @notice This public method allows anyone to pre-cache wallets
    function cacheWallet() public {
        // the address(uint160()) cast is done as the Wallet owner (1st argument) needs to be payable
        Wallet wallet = new Wallet(address(uint160(_ensResolve(WALLET_DEPLOYER_NODE))), true, ens, TOKEN_WHITELIST_NODE, CONTROLLER_NODE, LICENCE_NODE, defaultSpendLimit);
        cachedWallets.push(wallet);

        emit CachedWallet(wallet);
    }

    /// @notice This public method allows only the wallet deployer to pop pre-cached wallets or create a new one in case there aren't any
    function walletCachePop() external onlyWalletDeployer returns (Wallet) {
        if (cachedWallets.length < 1) {
            cacheWallet();
        }

        Wallet wallet = cachedWallets[cachedWallets.length-1];
        cachedWallets.pop();

        return wallet;
    }

    /// @notice returns the number of pre-cached wallets
    function cachedWalletsCount() external view returns (uint) {
        return cachedWallets.length;
    }

}
