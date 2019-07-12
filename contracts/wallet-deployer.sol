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

    event CachedWallet(address _wallet);
    event DeployedWallet(address _wallet, address _owner);
    event MigratedWallet(address _wallet, address _owner, address _oldWallet);

    mapping(address => address) public deployed;
    address[] public cached;

    address public ens;
    bytes32 public controllerName;
    bytes32 public oracleName;
    uint public defaultSpendLimit;

    /// @notice parameters are passed in so that they can be used to construct new instances of the wallet
    constructor(address _ens, bytes32 _controllerName, uint _spendLimit) Controllable(_ens, _controllerName) public {
        ens = _ens;
        controllerName = _controllerName;
        oracleName = subdomainNodeHash("oracle");
        controllerName = subdomainNodeHash("controller");
        require(controllerName == _controllerName, "wrong controller namehash");
        defaultSpendLimit = _spendLimit;
    }

    /// @notice This function is used to deploy a Wallet for a given owner address
    /// @param _owner is the owner address for the new Wallet to be
    function deployWallet(address _owner) external onlyController {
        if (cached.length < 1) {
            cacheWallet();
    	}

        address walletAddress = cached[cached.length-1];
        cached.length--;
        Wallet(walletAddress).transferOwnership(_owner);
        deployed[_owner] = walletAddress;

        emit DeployedWallet(walletAddress, _owner);
    }

    /// @notice This function is used to migrate an owner's security settings from a previous version of the wallet
    /// @param _owner is the owner address for the new Wallet to be
    /// @param _initializedSpendLimit is boolean to state whether or not spendLimit has been initialized
    /// @param _initializedTopUpLimit is boolean to state whether or not topUpLImit has been initialized
    /// @param _initializedWhitelist is boolean to state whether or not the whitelisted addresses has been initialized
    /// @param _spendLimit is the user's set daily spend limit
    /// @param _topUpLimit is the user's set daily gas top-up limit
    /// @param _whitelistedAddresses is the set of the user's whitelisted addresses
    function migrateWallet(address _owner, address _oldWallet, bool _initializedSpendLimit, bool _initializedTopUpLimit, bool _initializedWhitelist, uint _spendLimit, uint _topUpLimit, address[] _whitelistedAddresses) external onlyController {
        if (cached.length < 1) {
            cacheWallet();
    	}

        address walletAddress = cached[cached.length-1];
        cached.length--;

        if (_initializedSpendLimit) {
            Wallet(walletAddress).initializeSpendLimit(_spendLimit);
        }
        if (_initializedTopUpLimit) {
            Wallet(walletAddress).initializeTopUpLimit(_topUpLimit);
        }
        if (_initializedWhitelist) {
            Wallet(walletAddress).initializeWhitelist(_whitelistedAddresses);
        }

        Wallet(walletAddress).transferOwnership(_owner);
        deployed[_owner] = walletAddress;

        emit MigratedWallet(walletAddress, _owner, _oldWallet);
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

    function subdomainNodeHash(string memory _label) public pure returns (bytes32) {
        string memory tokencardLabel = "tokencard";
        string memory ethLabel = "eth";
        bytes32 zeroHash;
        bytes32 node = keccak256(abi.encodePacked(zeroHash, keccak256(bytes(ethLabel))));
        node = keccak256(abi.encodePacked(node, keccak256(bytes(tokencardLabel))));
        node = keccak256(abi.encodePacked(node, keccak256(bytes(_label))));
        return node;
    }
}
