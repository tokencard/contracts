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

pragma solidity ^0.5.10;

import "./wallet.sol";
import "./internals/controllable.sol";
import "./internals/ensResolvable.sol";

//// @title Wallet deployer with pre-caching if wallets functionality.
contract WalletDeployer is Controllable {

    event Received(address _from, uint _value);
    event CachedWallet(address _wallet);
    event DeployedWallet(address _wallet, address _owner);
    event MigratedWallet(address _wallet, address _owner);

    mapping(address => address) public deployed;
    address payable[] public cachedWallets;

    address public ens;
    bytes32 public tokenWhitelistNameHash;
    bytes32 public controllerNameHash;
    bytes32 public licenceNameHash;
    uint public defaultSpendLimit;

    /// @notice parameters are passed in so that they can be used to construct new instances of the wallet
    constructor(address _ens, bytes32 _tokenWhitelistNameHash, bytes32 _controllerNameHash, bytes32 _licenceNameHash, uint _defaultSpendLimit) ENSResolvable(_ens) Controllable(_controllerNameHash) public {
        ens = _ens;
        tokenWhitelistNameHash = _tokenWhitelistNameHash;
        controllerNameHash = _controllerNameHash;
        licenceNameHash = _licenceNameHash;
        defaultSpendLimit = _defaultSpendLimit;
    }

    /// @dev Ether can be deposited from any source, so this contract must be payable by anyone.
    function() external payable {
        emit Received(msg.sender, msg.value);
    }

    /// @notice This public method allows anyone to pre-cache wallets
    function cacheWallet() public {
        address payable walletAddress = address(new Wallet(address(this), true, ens, tokenWhitelistNameHash, controllerNameHash, licenceNameHash, defaultSpendLimit));
        cachedWallets.push(walletAddress);
        emit CachedWallet(walletAddress);
    }

    /// @notice This function is used to deploy a Wallet for a given owner address
    /// @param owner is the owner address for the new Wallet to be
    function deployWallet(address payable owner) external onlyController {
        if (cachedWallets.length < 1) {
            cacheWallet();
        }
        address payable walletAddress = cachedWallets[cachedWallets.length-1];
        cachedWallets.pop();
        Wallet(walletAddress).transferOwnership(owner, false);
        deployed[owner] = walletAddress;
        emit DeployedWallet(walletAddress, owner);
    }

    /// @notice This function is used to migrate an owner's security settings from a previous version of the wallet
    /// @param owner is the owner address for the new Wallet to be
    /// @param _spendLimit is the user's set daily spend limit
    /// @param _topUpLimit is the user's set daily gas top-up limit
    /// @param _whitelistedAddresses is the set of the user's whitelisted addresses
    function migrateWallet(address payable owner, uint _spendLimit, uint _topUpLimit, bool _initializedWhitelist, address[] calldata _whitelistedAddresses) external onlyController {
        if (cachedWallets.length < 1) {
            cacheWallet();
    	}

        address payable walletAddress = cachedWallets[cachedWallets.length-1];
        cachedWallets.pop();

        Wallet(walletAddress).setSpendLimit(_spendLimit);
        Wallet(walletAddress).setGasTopUpLimit(_topUpLimit);

        if (_initializedWhitelist) {
            Wallet(walletAddress).setWhitelist(_whitelistedAddresses);
        }

        Wallet(walletAddress).transferOwnership(owner, false);
        deployed[owner] = walletAddress;

        emit MigratedWallet(walletAddress, owner);
    }

    /// @notice returns the number of pre-cached wallets
    function cachedWalletCount() external view returns (uint) {
        return cachedWallets.length;
    }
}
