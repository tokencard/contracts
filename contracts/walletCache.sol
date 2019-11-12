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
import "./internals/controllable.sol";


/// @title IWalletCache interface describes a method for poping an already cached wallet
interface IWalletCache {
    function walletCachePop() external returns(Wallet);
}

//// @title Wallet cache with wallet pre-caching functionality.
contract WalletCache is ENSResolvable, Controllable {

    event CachedWallet(Wallet _wallet);

    /*****   Constants   *****/
    // Default values for mainnet ENS
    // licence.tokencard.eth
    bytes32 private constant _DEFAULT_LICENCE_NODE = 0xd0ff8bd67f6e25e4e4b010df582a36a0ee9b78e49afe6cc1cff5dd5a83040330;
    // token-whitelist.tokencard.eth
    bytes32 private constant _DEFAULT_TOKEN_WHITELIST_NODE = 0xe84f90570f13fe09f288f2411ff9cf50da611ed0c7db7f73d48053ffc974d396;
    // wallet-deployer.tokencard.eth
    bytes32 private constant _DEFAULT_WALLET_DEPLOYER_NODE = 0xd21a61ac2e4de1319ef7c76dd03046ec2e67a92cfc9efb7c28f79a4c323a5b80;

    bytes32 public licenceNode = _DEFAULT_LICENCE_NODE;
    bytes32 public tokenWhitelistNode= _DEFAULT_TOKEN_WHITELIST_NODE;
    bytes32 public walletDeployerNode = _DEFAULT_WALLET_DEPLOYER_NODE;

    Wallet[] public cachedWallets;

    address public ens;
    uint public defaultSpendLimit;

    /// @notice parameters are passed in so that they can be used to construct new instances of the wallet
    /// @dev pass in bytes32 to use the default, production node labels for ENS
    constructor(address _ens_, uint _defaultSpendLimit_, bytes32 _controllerNode_, bytes32 _licenceNode_, bytes32 _tokenWhitelistNode_, bytes32 _walletDeployerNode_) ENSResolvable(_ens_) Controllable(_controllerNode_) public {
        ens = _ens_;
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

    /// @notice This public method allows anyone to pre-cache wallets
    function cacheWallet() public {
        // the address(uint160()) cast is done as the Wallet owner (1st argument) needs to be payable
        Wallet wallet = new Wallet(address(uint160(_ensResolve(walletDeployerNode))), true, ens, tokenWhitelistNode, controllerNode(), licenceNode, defaultSpendLimit);
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
