/**
 *  The Consumer Contract Wallet
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

// SPDX-License-Identifier: GPLv3

pragma solidity ^0.6.11;

import "./externals/Address.sol";
import "./externals/ECDSA.sol";
import "./externals/initializable.sol";
import "./externals/SafeMath.sol";
import "./externals/SafeERC20.sol";
import "./interfaces/IERC20.sol";
import "./interfaces/IERC165.sol";
import "./interfaces/ILicence.sol";
import "./internals/balanceable.sol";
import "./internals/controllable.sol";
import "./internals/ensResolvable.sol";
import "./internals/ownable.sol";
import "./internals/tokenWhitelistable.sol";
import "./internals/transferrable.sol";

<<<<<<< HEAD

=======
>>>>>>> 68763a47... Upgrade prettier and remove wallet.sol from ingore list
/// @title SelfCallableOwnable allows either owner or the contract itself to call its functions
/// @dev providing an additional modifier to check if Owner or self is calling
/// @dev the "self" here is used for the meta transactions
abstract contract SelfCallableOwnable is Ownable {
    /// @dev Check if the sender is the Owner or self
    modifier onlySelf() {
        require(msg.sender == address(this), "Only self");
        _;
    }

    /// @dev Check if the sender is the Owner or self
    modifier onlyOwnerOrSelf() {
        require(_isOwner(msg.sender) || msg.sender == address(this), "Only owner or self");
        _;
    }
}

/// @title OptOutableMonolith2FA is used a configurable 2FA.
/// @dev This provides the various modifiers and utility functions needed for 2FA.
/// @dev 2FA is needed to confirm changes to the security settings in the wallet.
abstract contract OptOutableMonolith2FA is Controllable, SelfCallableOwnable {
    event SetMonolith2FA(address _sender);
    event SetPersonal2FA(address _sender, address _p2FA);

    /// @dev the operation can access the wallet in "privileged" mode i.e. sensitive operation
    bool public privileged;
    bool public monolith2FA;
    address public personal2FA;

    constructor() public {
        monolith2FA = true;
    }

    // @dev This modifier ensures that a method is only accessible to 2nd factor
    modifier only2FA() {
        if (monolith2FA) {
            require(_isController(msg.sender), "Only Monolith 2FA");
        } else {
            require(msg.sender == personal2FA, "Only personal 2FA");
        }
        _;
    }

    /// @dev set Monolith to be the 2FA
    function setMonolith2FA() external onlyOwner {
        require(!monolith2FA, "Monolith 2FA already enabled");

        monolith2FA = true;
        personal2FA = address(0);

        emit SetMonolith2FA(msg.sender);
    }

    /// @dev set personal 2FA to the address the user provided, needs to be called by a privileged relayed Tx
    function setPersonal2FA(address _p2FA) external onlySelf {
        require(privileged, "Setting a personal 2FA requires privileged mode");
        require(_p2FA != address(0), "2FA cannot be the 0 address");
        require(_p2FA != address(this), "2FA cannot be the wallet address");

        personal2FA = _p2FA;
        monolith2FA = false;

        emit SetPersonal2FA(msg.sender, _p2FA);
    }
}

abstract contract WalletDeployable is ENSResolvable {
    // wallet-deployer.v4.tokencard.eth
    bytes32 private constant _DEFAULT_WALLET_DEPLOYER_NODE = 0x23349faba58c4a8622c88e7d3ba4a01da2f0764900bfb876898ac21e573273c9;

    bytes32 public walletDeployerNode = _DEFAULT_WALLET_DEPLOYER_NODE;

    constructor(bytes32 _walletDeployerNode_) public {
        // Set walletDeployerNode or use default
        if (_walletDeployerNode_ != bytes32(0)) {
            walletDeployerNode = _walletDeployerNode_;
        }
    }

    /// @dev Check if the sender is the wallet-deployer
    modifier onlyWalletDeployer() {
        require(msg.sender == _ensResolve(walletDeployerNode), "Only walletDeployer");
        _;
    }
}

/// @title Recoverable provides wallet recovery functionality.
/// @dev This contract will allow the user revover their wallet i.e. change the owner in case they lose their back-up seed.
abstract contract Recoverable is SelfCallableOwnable, OptOutableMonolith2FA {
    /// @dev Recovers wallet by transfering ownership to a new address, needs privileged access.
    /// @param _newOwner address to transfer ownership to.
    /// @param _transferable indicates whether to keep the ownership transferable or not.
    function recoverWallet(address payable _newOwner, bool _transferable) external onlySelf {
        require(privileged, "Recovery requires privileged mode");
        _transferOwnership(_newOwner, _transferable);
    }
}

/// @title AddressWhitelist provides payee-whitelist functionality.
/// @dev This contract will allow the user to maintain a whitelist of addresses.
/// @dev These addresses will live outside of the daily limit.
abstract contract AddressWhitelist is WalletDeployable, SelfCallableOwnable, OptOutableMonolith2FA {
    using SafeMath for uint256;

    event AddedToWhitelist(address _sender, address[] _addresses);
    event RemovedFromWhitelist(address _sender, address[] _addresses);

    mapping(address => bool) public whitelistMap;
    address[] public whitelistArray;
    // This variable ensures that migrateWhitelist() can only be called once by the walletDeployer.
    bool private _migratedWhitelist;

    /// @dev Check if the provided addresses contain the owner or the zero-address address.
    modifier hasNoOwnerOrZeroAddress(address[] memory _addresses) {
        for (uint256 i = 0; i < _addresses.length; i++) {
            require(!_isOwner(_addresses[i]), "Contains owner address");
            require(_addresses[i] != address(0), "Contains 0 address");
        }
        _;
    }

    /// @dev Adds addresses to AddressWhitelist.
    /// @param _addresses The addresses to be whitelisted.
    function addToWhitelist(address[] calldata _addresses) external onlySelf {
        require(privileged, "Adding to whitelist requires privileged mode");
        _addToWhitelist(_addresses);
    }

    /// @dev This is used to restore the the address whitelist, called by the walletDeployer during wallet migration.
    /// @param _addresses The addresses whitelisted by the owner before migration.
    function migrateWhitelist(address[] calldata _addresses) external onlyWalletDeployer {
        require(!_migratedWhitelist, "Whitelist already migrated");
        _migratedWhitelist = true;
        _addToWhitelist(_addresses);
    }

    /// @dev Removes addresses from AddressWhitelist.
    /// @param _addresses The addresses to be removed from the whitelist.
    function removeFromWhitelist(address[] calldata _addresses) external onlySelf {
        require(privileged, "Removing from whitelist requires privileged mode");
        // Require that the input list is not empty.
        require(_addresses.length != 0, "Empty list to be removed from whitelist");
        // Require that the whiteList is not empty.
        require(whitelistArray.length != 0, "Whitelist is empty");
        // Remove provided addresses.
        for (uint256 i = 0; i < _addresses.length; i++) {
            // Require addresses to have been already whitelisted.
            require(whitelistMap[_addresses[i]], "Address not whitelisted");
            whitelistMap[_addresses[i]] = false;
            for (uint256 j = 0; j < whitelistArray.length.sub(1); j++) {
                if (whitelistArray[j] == _addresses[i]) {
                    whitelistArray[j] = whitelistArray[whitelistArray.length - 1];
                    break;
                }
            }
            whitelistArray.pop();
        }
        // Emit the removal event.
        emit RemovedFromWhitelist(msg.sender, _addresses);
    }

    /// @dev Adds addresses to AddressWhitelist, called by addToWhitelist() and migrateWhitelist().
    /// @param _addresses The addresses to be whitelisted.
    function _addToWhitelist(address[] memory _addresses) private hasNoOwnerOrZeroAddress(_addresses) {
        // Require that the list is not empty.
        require(_addresses.length != 0, "Empty list to be added to whitelist");
        // Add each of the provided addresses to the whitelist.
        for (uint256 i = 0; i < _addresses.length; i++) {
            // Require addresses not be already whitelisted.
            require(!whitelistMap[_addresses[i]], "Address already whitelisted");
            // adds to the whitelist mapping
            whitelistMap[_addresses[i]] = true;
            // adds to the whitelist array
            whitelistArray.push(_addresses[i]);
        }
        // Emit the addition event.
        emit AddedToWhitelist(msg.sender, _addresses);
    }
}

/// @title DailyLimit provides daily spend limit functionality.
abstract contract DailyLimit is WalletDeployable, SelfCallableOwnable, OptOutableMonolith2FA, TokenWhitelistable {
    using SafeMath for uint256;

    event SetDailyLimit(address _sender, uint256 _amount);
    event UpdatedAvailableDailyLimit(uint256 _amount, uint256 _nextReset);

    uint256 private _available; // The remaining amount available for spending in the current 24-hour window.
    uint256 private _dailyLimit; // The daily limit amount.
    bool private _migratedDailyLimit; // This variable ensures that migratedDailyLimit() can only be called once by the walletDeployer.
    uint256 private _resetTimestamp; // Denotes the future timestamp when the available daily limit is due to reset.

    /// @dev Constructor initializes the daily limit in wei.
    constructor(uint256 _limit, bytes32 _tokenWhitelistNode) internal TokenWhitelistable(_tokenWhitelistNode) {
        (, uint256 stablecoinMagnitude, , , , , ) = _getStablecoinInfo();
        require(stablecoinMagnitude > 0, "Stablecoin not set");
        uint256 limitBaseUnits = _limit * stablecoinMagnitude;
        _dailyLimit = limitBaseUnits;
        _available = limitBaseUnits;
        _resetTimestamp = now.add(24 hours);
        emit UpdatedAvailableDailyLimit(limitBaseUnits, _resetTimestamp);
    }

    /// @dev View the limit value
    function dailyLimitValue() external view returns (uint256) {
        return _dailyLimit;
    }

    /// @dev Returns the available daily limit/balance, accounts for daily limit reset.
    /// @return amount of available to spend within the current day in base units.
    function dailyLimitAvailable() external view returns (uint256) {
        if (now > _resetTimestamp) {
            return _dailyLimit;
        } else {
            return _available;
        }
    }

    /// @dev Restores the daily limit, called by the walletDeployer duirng wallet migration.
    /// @param _amount The limit set by the owner before migration.
    function migrateDailyLimit(uint256 _amount) external onlyWalletDeployer {
        require(!_migratedDailyLimit, "DailyLimit already migrated");
        _migratedDailyLimit = true;
        // Set the daily limit...
        _setDailyLimit(_amount);
    }

    /// @dev Set daily limit operation.
    /// @param _amount the new limit amount.
    function setDailyLimit(uint256 _amount) external onlySelf {
        require(privileged, "Setting the daily limit requires privileged mode");
        _setDailyLimit(_amount);
    }

    /// @dev Use up amount within the daily limit. Will fail if amount is larger than available limit.
    function _enforceDailyLimit(uint256 _amount) internal {
        // Account for the limit daily reset.
        if (now > _resetTimestamp) {
            // Set the available limit to the current daily limit.
            _available = _dailyLimit;
            // Update the current timestamp.
            _resetTimestamp = now.add(24 hours);
        }
        require(_available >= _amount, "Spend amount exceeds available limit");
        _available = _available.sub(_amount);
        emit UpdatedAvailableDailyLimit(_available, _resetTimestamp);
    }

    /// @dev Modify the daily limit and spend available based on the provided value.
    /// @dev _amount is the daily limit amount in stablecoin base units.
    function _setDailyLimit(uint256 _amount) private {
        // Set the daily limit to the provided amount.
        _dailyLimit = _amount;
        emit SetDailyLimit(msg.sender, _dailyLimit);
        // update the available amount...
        _available = _amount;
        // ...and reset the 24-hour window
        _resetTimestamp = now.add(24 hours);
        emit UpdatedAvailableDailyLimit(_available, _resetTimestamp);
    }
<<<<<<< HEAD

    /// @dev Update available limit based on the daily reset.
    function _updateAvailableDailyLimit() private {
        if (now > _resetTimestamp) {
            // Update the current timestamp.
            _resetTimestamp = now.add(24 hours);
            // Set the available limit to the current daily limit.
            _available = _dailyLimit;
        }
    }

    function _initializeLoadLimit(bytes32 _tokenWhitelistNode) internal initializer {
        _initializeTokenWhitelistable(_tokenWhitelistNode);
        (, uint256 stablecoinMagnitude, , , , , ) = _getStablecoinInfo();
        require(stablecoinMagnitude > 0, "no stablecoin");
        _maximumLoadLimit = _MAXIMUM_STABLECOIN_LOAD_LIMIT * stablecoinMagnitude;
        _loadLimit = DailyLimitTrait.DailyLimit(_maximumLoadLimit, _maximumLoadLimit, now, 0, false);
    }
}

=======
}

>>>>>>> 68763a47... Upgrade prettier and remove wallet.sol from ingore list
/// @title Asset wallet with extra security features and card integration.
contract Wallet is ENSResolvable, WalletDeployable, AddressWhitelist, DailyLimit, IERC165, Transferrable, Balanceable {
    using Address for address;
    using ECDSA for bytes32;
    using SafeERC20 for IERC20;
    using SafeMath for uint256;

    event BulkTransferred(address _to, address[] _assets);
    event ExecutedRelayedTransaction(bytes _data, bool _privileged);
    event ExecutedTransaction(address _destination, uint256 _value, bytes _data, bytes _returndata);
    event IncreasedRelayNonce(address _sender, uint256 _currentNonce);
    event LoadedTokenCard(address _asset, uint256 _amount);
    event ToppedUpGas(address _sender, address _owner, uint256 _amount);
    event Transferred(address _to, address _asset, uint256 _amount);
    event UpdatedAvailableLimit(); // This is here because our tests don't inherit events from a library

    string public constant WALLET_VERSION = "3.4.1";

    // keccak256("isValidSignature(bytes,bytes)") = 20c13b0bc670c284a9f19cdf7a533ca249404190f8dc132aac33e733b965269e
    bytes4 private constant _EIP_1271 = 0x20c13b0b;
    // keccak256("isValidSignature(bytes32,bytes)") = 1626ba7e356f5979dd355a3d2bfb43e80420a480c3b854edce286a82d7496869
    bytes4 private constant _EIP_1654 = 0x1626ba7e;

    /// @dev Supported ERC165 interface ID.
    bytes4 private constant _ERC165_INTERFACE_ID = 0x01ffc9a7; // solium-disable-line uppercase

    /// @dev This is an internal nonce to prevent replay attacks from relayer
    uint256 public relayNonce;

    /// @dev This ensures that transferOwnership() can only be called once by the walletDeployer during deployment or migration.
    bool private _transferredOwnership;

    /// @dev Is the registered ENS node identifying the licence contract.
    bytes32 private _licenceNode;

    /// @dev Constructor initializes the wallet with an owner address and daily limit. It also sets up the controllable and tokenWhitelist contracts with the right name registered in ENS.
    /// @param _owner_ is the owner account of the wallet contract.
    /// @param _transferable_ indicates whether the contract ownership can be transferred.
    /// @param _ens_ is the address of the ENS registry.
    /// @param _tokenWhitelistNode_ is the ENS name node of the Token whitelist.
    /// @param _controllerNode_ is the ENS name node of the Controller contract.
    /// @param _licenceNode_ is the ENS name node of the Licence contract.
    /// @param _dailyLimit_ is the initial spend limit.
    function initializeWallet(
        address payable _owner_,
        bool _transferable_,
        address _ens_,
        bytes32 _tokenWhitelistNode_,
        bytes32 _controllerNode_,
        bytes32 _licenceNode_,
        bytes32 _walletDeployerNode_,
        uint256 _dailyLimit_
<<<<<<< HEAD
    ) external initializer {
        _initializeENSResolvable(_ens_);
        _initializeControllable(_controllerNode_);
        _initializeOwnable(_owner_, _transferable_);
        _initializeSpendLimit(_dailyLimit_);
        _initializeLoadLimit(_tokenWhitelistNode_);
=======
    )
        public
        ENSResolvable(_ens_)
        WalletDeployable(_walletDeployerNode_)
        DailyLimit(_dailyLimit_, _tokenWhitelistNode_)
        Ownable(_owner_, _transferable_)
        Controllable(_controllerNode_)
    {
>>>>>>> 68763a47... Upgrade prettier and remove wallet.sol from ingore list
        _licenceNode = _licenceNode_;
    }

    /// @dev Checks if the value is not zero.
    modifier isNotZero(uint256 _value) {
        require(_value != 0, "value=0");
        _;
    }

    /// @dev Ether can be deposited from any source, so this contract must be payable by anyone.
    receive() external payable {
        emit Received(msg.sender, msg.value);
    }

    /// @dev This is a bulk transfer convenience function, used to migrate contracts.
    /// @notice If any of the transfers fail, this will revert.
    /// @param _to is the recipient's address, can't be the zero (0x0) address: transfer() will revert.
    /// @param _assets is an array of addresses of ERC20 tokens or 0x0 for ether.
    function bulkTransfer(address payable _to, address[] calldata _assets) external onlyOwnerOrSelf {
        // check to make sure that _assets isn't empty
        require(_assets.length != 0, "asset array is empty");
        // This loops through all of the transfers to be made
        for (uint256 i = 0; i < _assets.length; i++) {
            uint256 amount = _balance(address(this), _assets[i]);
            // use our safe, daily limit protected transfer
            transfer(_to, _assets[i], amount);
        }

        emit BulkTransferred(_to, _assets);
    }

    /// Meta-transaction
    function executeRelayedTransaction(
        uint256 _nonce,
        bytes calldata _data,
        bytes calldata _signature
    ) external onlyController {
        return _executeRelayedTransaction(_nonce, _data, _signature, false);
    }

    /// Privileged functionality
    function executePrivilegedRelayedTransaction(
        uint256 _nonce,
        bytes calldata _data,
        bytes calldata _signature
    ) external only2FA {
        return _executeRelayedTransaction(_nonce, _data, _signature, true);
    }

    /// @dev Enables the wallet-deployer to transfer ownership to the new owner, can be used only once.
    /// @param _newOwner address to transfer ownership to.
    /// @param _transferable indicates whether to keep the ownership transferable or not.
    function transferOwnership(address payable _newOwner, bool _transferable) external onlyWalletDeployer {
        require(!_transferredOwnership, "Ownership already transferred");
        _transferredOwnership = true;
        _transferOwnership(_newOwner, _transferable);
    }

    /// @dev This function allows for the controller to relay transactions on the owner's behalf,
    /// the relayed message has to be signed by the owner.
    /// @param _nonce only used for relayed transactions, must match the wallet's relayNonce.
    /// @param _data abi encoded data payload.
    /// @param _signature signed prefix + data.
    function _executeRelayedTransaction(
        uint256 _nonce,
        bytes memory _data,
        bytes memory _signature,
        bool _privileged
    ) private {
        // expecting prefixed data ("rlx:") indicating relayed transaction...
        // ...and an Ethereum Signed Message to protect user from signing an actual Tx
        uint256 id;
        assembly {
            id := chainid() //1 for Ethereum mainnet, > 1 for public testnets.
        }
        bytes32 dataHash = keccak256(abi.encodePacked("monolith:", id, address(this), _nonce, _data)).toEthSignedMessageHash();
        // Verify signature validity i.e. signer == owner
        require(isValidSignature(dataHash, _signature) == _EIP_1654, "sig not valid");
        // Verify and increase relayNonce to prevent replay attacks from the relayer
        require(_nonce == relayNonce, "tx replay");
        _increaseRelayNonce();

        // TO DO: an "if(_privileged) {privileged = _privileged}  should be less expensive
        privileged = _privileged;
        _batchExecuteTransaction(_data);
        privileged = false;

        emit ExecutedRelayedTransaction(_data, _privileged);
    }

    /// @dev This allows the user to cancel a transaction that was unexpectedly delayed by the relayer
    function increaseRelayNonce() external onlyOwner {
        _increaseRelayNonce();
    }

    /// @dev This bumps the relayNonce and emits an event accordingly
    function _increaseRelayNonce() private {
        relayNonce++;
        emit IncreasedRelayNonce(msg.sender, relayNonce);
    }

    /// @dev Implements EIP-1271: receives the raw data (bytes)
    ///      https://github.com/ethereum/EIPs/blob/master/EIPS/eip-1271.md
    /// @param _data Arbitrary length data signed on the behalf of address(this)
    /// @param _signature Signature byte array associated with _data
    function isValidSignature(bytes calldata _data, bytes calldata _signature) external view returns (bytes4) {
        bytes32 dataHash = keccak256(abi.encodePacked(_data));
        // isValidSignature call reverts if not valid.
        require(isValidSignature(dataHash, _signature) == _EIP_1654, "sig not valid");
        return _EIP_1271;
    }

    /// @return licence contract node registered in ENS.
    function licenceNode() external view returns (bytes32) {
        return _licenceNode;
    }

    /// @dev Load a token card with the specified asset amount.
    /// @dev the amount send should be inclusive of the percent licence.
    /// @param _asset is the address of an ERC20 token or 0x0 for ether.
    /// @param _amount is the amount of assets to be transferred in base units.
    function loadTokenCard(address _asset, uint256 _amount) external payable onlyOwnerOrSelf {
        // check if token is allowed to be used for loading the card
        require(_isTokenLoadable(_asset), "token not loadable");
        // if it is not privileged enforce the limit
        if (!privileged) {
            // Convert token amount to stablecoin value.
            // If the asset is not available (2nd return value) then revertstablecoinValue will be zero
            uint256 stablecoinValue = convertToStablecoin(_asset, _amount);
            // Check against the daily spent limit and update accordingly, require that the value is under remaining limit.
            _enforceDailyLimit(stablecoinValue);
        }
        // Get the TKN licenceAddress from ENS
        address licenceAddress = _ensResolve(_licenceNode);
        if (_asset != address(0)) {
            IERC20(_asset).safeApprove(licenceAddress, _amount);
            ILicence(licenceAddress).load(_asset, _amount);
        } else {
            require(address(this).balance >= _amount, "Load: not sufficient balance");
            ILicence(licenceAddress).load{value: _amount}(_asset, _amount);
        }

        emit LoadedTokenCard(_asset, _amount);
    }

    /// @dev Checks for interface support based on ERC165.
    function supportsInterface(bytes4 _interfaceID) external override view returns (bool) {
        return _interfaceID == _ERC165_INTERFACE_ID;
    }

    /// @dev This function allows for the wallet to send a batch of transactions instead of one,
    /// it calls executeTransaction() so that the daily limit is enforced.
    /// @param _transactionBatch data encoding the transactions to be sent,
    /// following executeTransaction's format i.e. (destination, value, data)
    function _batchExecuteTransaction(bytes memory _transactionBatch) private {
        uint256 batchLength = _transactionBatch.length + 32; // because the pos starts from 32
        uint256 remainingBytesLength = _transactionBatch.length; // remaining bytes to be processed
        uint256 pos = 32; //the first 32 bytes denote the byte array length, start from actual data

        address destination; // destination address
        uint256 value; // trasanction value
        uint256 dataLength; // externall call data length
        bytes memory data; // call data

        while (pos < batchLength) {
            // there should always be at least 84 bytes remaining: the minimun #bytes required to encode a Tx
            remainingBytesLength = remainingBytesLength.sub(84);
            assembly {
                // shift right by 96 bits (256 - 160) to get the destination address (and zero the excessive bytes)
                destination := shr(96, mload(add(_transactionBatch, pos)))
                // get value: pos + 20 bytes (destinnation address)
                value := mload(add(_transactionBatch, add(pos, 20)))
                // get data: pos  + 20 (destination address) + 32 (value) bytes
                // the first 32 bytes denote the byte array length
                dataLength := mload(add(_transactionBatch, add(pos, 52)))
                data := add(_transactionBatch, add(pos, 52))
            }
            // pos += 20 + 32 + 32 + dataLength, reverts in case of overflow
            pos = pos.add(dataLength).add(84);
            // revert in case the encoded dataLength is gonna cause an out of bound access
            require(pos <= batchLength, "out of bounds");

            // if length is 0 ignore the data field
            if (dataLength == 0) {
                data = bytes("");
            }

            // call executeTransaction(), if one of them reverts then the whole batch reverts.
            _executeTransaction(destination, value, data);
        }
    }

    /// @dev Convert ether or ERC20 token amount to the corresponding stablecoin amount.
    /// @param _token ERC20 token contract address.
    /// @param _amount amount of token in base units.
    /// @return the eqivalent amount in stablecoin base units & 0 if the token is not available.
    function convertToStablecoin(address _token, uint256 _amount) public view returns (uint256) {
        // avoid the unnecessary calculations if the token to be loaded is the stablecoin itself
        if (_token == _stablecoin()) {
            return _amount;
        }

        uint256 amountToSend = _amount;
        // convert token amount to ETH first (the 0x0 address represents ether)
        if (_token != address(0)) {
            // Store the token in memory to save map entry lookup gas.
            (, uint256 magnitude, uint256 rate, bool available, , , ) = _getTokenInfo(_token);
            // if the token does NOT exist in the whitelist then return 0
            if (!available) {
                return 0;
            }
            // if it exists then require that its rate is not zero.
            require(rate != 0, "rate=0");
            // Safely convert the token amount to ether based on the exchange rate.
            amountToSend = _amount.mul(rate).div(magnitude);
        }
        // _amountToSend now is in ether
        // Get the stablecoin's magnitude and its current rate.
        (, uint256 stablecoinMagnitude, uint256 stablecoinRate, bool stablecoinAvailable, , , ) = _getStablecoinInfo();
        // Check if the stablecoin rate is set.
        require(stablecoinAvailable, "token not available");
        require(stablecoinRate != 0, "stablecoin rate=0");
        // Safely convert the token amount to stablecoin based on its exchange rate and the stablecoin exchange rate.
        return amountToSend.mul(stablecoinMagnitude).div(stablecoinRate);
    }

    function executeTransaction(
        address _destination,
        uint256 _value,
        bytes calldata _data
    ) external onlyOwner returns (bytes memory) {
        _executeTransaction(_destination, _value, _data);
    }

    /// @dev This function allows for the owner to send any transaction from the Wallet to arbitrary addresses
    /// @param _destination address of the transaction
    /// @param _value ETH amount in wei
    /// @param _data transaction payload binary
    function _executeTransaction(
        address _destination,
        uint256 _value,
        bytes memory _data
    ) private returns (bytes memory) {
        // If value is send across as a part of this executeTransaction, this will be sent to any payable
        // destination. As a result enforceLimit if destination is not whitelisted.

        if (!whitelistMap[_destination] && !privileged) {
            // Convert ETH value to stablecoin, 0x0 denotes ETH.
            uint256 stablecoinValue = convertToStablecoin(address(0), _value);
            _enforceDailyLimit(stablecoinValue);
        }
        // Check if the destination is a Contract and it is one of our supported tokens
        if (address(_destination).isContract() && _isTokenAvailable(_destination)) {
            // to is the recipient's address and amount is the value to be transferred
            address to;
            uint256 amount;
            (to, amount) = _getERC20RecipientAndAmount(_destination, _data);
            if (!whitelistMap[to] && !privileged) {
                // Convert token amount to stablecoin value.
                // If the address (of the token contract) is not in the TokenWhitelist used by the convert method...
                // ...then stablecoinValue will be zero
                uint256 stablecoinValue = convertToStablecoin(_destination, amount);
                _enforceDailyLimit(stablecoinValue);
            }
            // use callOptionalReturn provided in SafeERC20 in case the ERC20 method
            // returns false instead of reverting!
            IERC20(_destination)._callOptionalReturn(_data);

            // if ERC20 call completes, return a boolean "true" as bytes emulating ERC20
            bytes memory b = new bytes(32);
            b[31] = 0x01;

            emit ExecutedTransaction(_destination, _value, _data, b);
            return b;
        }

        (bool success, bytes memory returndata) = _destination.call{value: _value}(_data);
        require(success, string(returndata));

        emit ExecutedTransaction(_destination, _value, _data, returndata);
        // returns all of the bytes returned by _destination contract
        return returndata;
    }

    /// @dev Implements EIP-1654: receives the hashed message(bytes32)
    /// https://github.com/ethereum/EIPs/issues/1654.md
    /// @param _hashedData Hashed data signed on the behalf of address(this)
    /// @param _signature Signature byte array associated with _dataHash
    function isValidSignature(bytes32 _hashedData, bytes memory _signature) public view returns (bytes4) {
        address from = _hashedData.recover(_signature);
        require(_isOwner(from), "invalid signature");
        return _EIP_1654;
    }

    /// @dev Transfers the specified asset to the recipient's address.
    /// @param _to is the recipient's address.
    /// @param _asset is the address of an ERC20 token or 0x0 for ether.
    /// @param _amount is the amount of assets to be transferred in base units.
    function transfer(
        address payable _to,
        address _asset,
        uint256 _amount
    ) public onlyOwnerOrSelf isNotZero(_amount) {
        // Checks if the _to address is not the zero-address
        require(_to != address(0), "destination=0");

        // If address is not whitelisted, take daily limit into account.
        if (!whitelistMap[_to] && !privileged) {
            // Convert token amount to stablecoin value.
            // If the address (of the token contract) is not in the TokenWhitelist used by the convert method...
            // ...then stablecoinValue will be zero
            uint256 stablecoinValue = convertToStablecoin(_asset, _amount);
            // Check against the daily spent limit and update accordingly, require that the value is under remaining limit.
            _enforceDailyLimit(stablecoinValue);
        }
        // Transfer token or ether based on the provided address.
        _safeTransfer(_to, _asset, _amount);
        // Emit the transfer event.
        emit Transferred(_to, _asset, _amount);
    }
}
