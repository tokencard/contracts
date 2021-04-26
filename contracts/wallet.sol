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
 *  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 *  GNU General Public License for more details.

 *  You should have received a copy of the GNU General Public License
 *  along with this program. If not, see <https://www.gnu.org/licenses/>.
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

/// @title OptOutableMonolith2FA is used a configurable 2FA.
/// @dev This provides the various modifiers and utility functions needed for 2FA.
/// @dev 2FA is needed to confirm changes to the security settings in the wallet.
contract OptOutableMonolith2FA is Controllable, Ownable {
    event SetPersonal2FA(address _sender, address _p2FA);
    event SetMonolith2FA(address _sender);

    /// @dev for accessing wallet in privileged mode.
    bool public privileged;
    bool public monolith2FA;
    address public personal2FA;

    function _initialize2FA() internal initializer {
        monolith2FA = true;
    }

    // @dev This modifier ensures that a method is only accessible to 2nd factor
    modifier only2FA() {
        if (monolith2FA) {
            require(_isController(msg.sender), "sender is not a Monolith 2FA");
        } else {
            require(msg.sender == personal2FA, "sender is not personal 2FA account");
        }
        _;
    }

    /// @dev Check if the sender is the Owner or 2FA
    modifier onlyOwnerOr2FA() {
        require(_isOwner(msg.sender) || _is2FA(msg.sender), "only owner or 2FA");
        _;
    }

    /// @dev set Monolith to be the 2FA
    function setMonolith2FA() external onlyOwner {
        require(!monolith2FA, "monolith2FA already enabled");
        monolith2FA = true;
        personal2FA = address(0);
        emit SetMonolith2FA(msg.sender);
    }

    /// @dev set personal 2FA to the address the user provided, needs to be called by a privileged relayed Tx
    function setPersonal2FA(address _p2FA) external onlyOwner {
        //require(privileged, "Setting a personal 2FA requires privileged mode");
        require(_p2FA != address(0), "2FA cannot be set to zero");
        require(_p2FA != personal2FA, "address already set");
        require(_p2FA != address(this), "2FA cannot be the wallet address");
        personal2FA = _p2FA;
        monolith2FA = false;
        emit SetPersonal2FA(msg.sender, _p2FA);
    }

    /// @dev utility function to check whether or not an address is valid 2FA'er
    function _is2FA(address _sender) private view returns (bool) {
        if (monolith2FA) {
            return _isController(_sender);
        } else {
            return (_sender == personal2FA);
        }
    }
}

/// @title SelfCallableOwnable allows either owner or the contract itself to call its functions
/// @dev providing an additional modifier to check if Owner or self is calling
/// @dev the "self" here is used for the meta transactions
contract SelfCallableOwnable is Ownable {
    /// @dev Check if the sender is the Owner or self
    modifier onlySelf() {
        require(msg.sender == address(this), "not self");
        _;
    }

    /// @dev Check if the sender is the Owner or self
    modifier onlyOwnerOrSelf() {
        require(_isOwner(msg.sender) || msg.sender == address(this), "Not owner or self");
        _;
    }
}

/// @title AddressWhitelist provides payee-whitelist functionality.
/// @dev This contract will allow the user to maintain a whitelist of addresses
/// @dev These addresses will live outside of the daily limit
contract AddressWhitelist is OptOutableMonolith2FA, SelfCallableOwnable {
    using SafeMath for uint256;

    event AddedToWhitelist(address _sender, address[] _addresses);
    event CancelledWhitelistAddition(address _sender, bytes32 _hash);
    event SubmittedWhitelistAddition(address[] _addresses, bytes32 _hash);

    event CancelledWhitelistRemoval(address _sender, bytes32 _hash);
    event RemovedFromWhitelist(address _sender, address[] _addresses);
    event SubmittedWhitelistRemoval(address[] _addresses, bytes32 _hash);

    mapping(address => bool) public whitelistMap;
    address[] public whitelistArray;
    address[] private _pendingWhitelistAddition;
    address[] private _pendingWhitelistRemoval;
    bool public submittedWhitelistAddition;
    bool public submittedWhitelistRemoval;
    bool public isSetWhitelist;

    /// @dev Check if the provided addresses contain the owner or the zero-address address.
    modifier hasNoOwnerOrZeroAddress(address[] memory _addresses) {
        for (uint256 i = 0; i < _addresses.length; i++) {
            require(!_isOwner(_addresses[i]), "contains owner address");
            require(_addresses[i] != address(0), "contains 0 address");
        }
        _;
    }

    /// @dev Check that neither addition nor removal operations have already been submitted.
    modifier noActiveSubmission() {
        require(!submittedWhitelistAddition && !submittedWhitelistRemoval, "whitelist submission pending");
        _;
    }

    /// @dev Cancel pending whitelist addition.
    function cancelWhitelistAddition(bytes32 _hash) external onlyOwnerOr2FA {
        // Check if operation has been submitted.
        require(submittedWhitelistAddition, "no pending submission");
        // Require that confirmation hash and the hash of the pending whitelist addition match
        require(_hash == calculateHash(_pendingWhitelistAddition), "non-matching pending whitelist hash");
        // Reset pending addresses.
        delete _pendingWhitelistAddition;
        // Reset the submitted operation flag.
        submittedWhitelistAddition = false;
        // Emit the cancellation event.
        emit CancelledWhitelistAddition(msg.sender, _hash);
    }

    /// @dev Cancel pending removal of whitelisted addresses.
    function cancelWhitelistRemoval(bytes32 _hash) external onlyOwnerOr2FA {
        // Check if operation has been submitted.
        require(submittedWhitelistRemoval, "no pending submission");
        // Require that confirmation hash and the hash of the pending whitelist removal match
        require(_hash == calculateHash(_pendingWhitelistRemoval), "non-matching pending whitelist hash");
        // Reset pending addresses.
        delete _pendingWhitelistRemoval;
        // Reset pending addresses.
        submittedWhitelistRemoval = false;
        // Emit the cancellation event.
        emit CancelledWhitelistRemoval(msg.sender, _hash);
    }

    /// @dev Confirm pending whitelist addition.
    /// @dev This will only ever be applied post 2FA, by one of the Controllers
    /// @param _hash is the hash of the pending whitelist array, a form of Lamport's lock
    function confirmWhitelistAddition(bytes32 _hash) external only2FA {
        // Require that the whitelist addition has been submitted.
        require(submittedWhitelistAddition, "no pending submission");
        // Require that confirmation hash and the hash of the pending whitelist addition match
        require(_hash == calculateHash(_pendingWhitelistAddition), "non-matching pending whitelist hash");
        // Whitelist pending addresses.
        for (uint256 i = 0; i < _pendingWhitelistAddition.length; i++) {
            // check if it doesn't exist already.
            if (!whitelistMap[_pendingWhitelistAddition[i]]) {
                // add to the Map and the Array
                whitelistMap[_pendingWhitelistAddition[i]] = true;
                whitelistArray.push(_pendingWhitelistAddition[i]);
            }
        }
        // Emit the addition event.
        emit AddedToWhitelist(msg.sender, _pendingWhitelistAddition);
        // Reset pending addresses.
        delete _pendingWhitelistAddition;
        // Reset the submission flag.
        submittedWhitelistAddition = false;
    }

    /// @dev Confirm pending removal of whitelisted addresses.
    function confirmWhitelistRemoval(bytes32 _hash) external only2FA {
        // Require that the pending whitelist is not empty and the operation has been submitted.
        require(submittedWhitelistRemoval, "no pending submission");
        // Require that confirmation hash and the hash of the pending whitelist removal match
        require(_hash == calculateHash(_pendingWhitelistRemoval), "non-matching pending whitelist hash");
        // Remove pending addresses.
        for (uint256 i = 0; i < _pendingWhitelistRemoval.length; i++) {
            // check if it exists
            if (whitelistMap[_pendingWhitelistRemoval[i]]) {
                whitelistMap[_pendingWhitelistRemoval[i]] = false;
                for (uint256 j = 0; j < whitelistArray.length.sub(1); j++) {
                    if (whitelistArray[j] == _pendingWhitelistRemoval[i]) {
                        whitelistArray[j] = whitelistArray[whitelistArray.length - 1];
                        break;
                    }
                }
                whitelistArray.pop();
            }
        }
        // Emit the removal event.
        emit RemovedFromWhitelist(msg.sender, _pendingWhitelistRemoval);
        // Reset pending addresses.
        delete _pendingWhitelistRemoval;
        // Reset the submission flag.
        submittedWhitelistRemoval = false;
    }

    /// @dev Getter for pending addition array.
    function pendingWhitelistAddition() external view returns (address[] memory) {
        return _pendingWhitelistAddition;
    }

    /// @dev Getter for pending removal array.
    function pendingWhitelistRemoval() external view returns (address[] memory) {
        return _pendingWhitelistRemoval;
    }

    /// @dev Add initial addresses to the whitelist.
    /// @param _addresses are the Ethereum addresses to be whitelisted.
    function setWhitelist(address[] calldata _addresses) external onlyOwnerOrSelf hasNoOwnerOrZeroAddress(_addresses) {
        // Require that the whitelist has not been initialized.
        require(!isSetWhitelist, "whitelist initialized");
        // Add each of the provided addresses to the whitelist.
        for (uint256 i = 0; i < _addresses.length; i++) {
            // Dedup addresses before writing to the whitelist
            if (!whitelistMap[_addresses[i]]) {
                // adds to the whitelist mapping
                whitelistMap[_addresses[i]] = true;
                // adds to the whitelist array
                whitelistArray.push(_addresses[i]);
            }
        }
        isSetWhitelist = true;
        // Emit the addition event.
        emit AddedToWhitelist(msg.sender, whitelistArray);
    }

    /// @dev Add addresses to the whitelist.
    /// @param _addresses are the Ethereum addresses to be whitelisted.
    function submitWhitelistAddition(address[] calldata _addresses) external onlyOwnerOrSelf noActiveSubmission hasNoOwnerOrZeroAddress(_addresses) {
        // Require that the whitelist has been initialized.
        require(isSetWhitelist, "whitelist not initialized");
        // Require this array of addresses not empty
        require(_addresses.length > 0, "empty whitelist");
        // Set the provided addresses to the pending addition addresses.
        _pendingWhitelistAddition = _addresses;
        // Flag the operation as submitted.
        submittedWhitelistAddition = true;
        // Emit the submission event.
        emit SubmittedWhitelistAddition(_addresses, calculateHash(_addresses));
    }

    /// @dev Remove addresses from the whitelist.
    /// @param _addresses are the Ethereum addresses to be removed.
    function submitWhitelistRemoval(address[] calldata _addresses) external onlyOwnerOrSelf noActiveSubmission {
        // Require that the whitelist has been initialized.
        require(isSetWhitelist, "whitelist not initialized");
        // Require that the array of addresses is not empty
        require(_addresses.length > 0, "empty whitelist");
        // Add the provided addresses to the pending addition list.
        _pendingWhitelistRemoval = _addresses;
        // Flag the operation as submitted.
        submittedWhitelistRemoval = true;
        // Emit the submission event.
        emit SubmittedWhitelistRemoval(_addresses, calculateHash(_addresses));
    }

    /// @dev Method used to hash our whitelist address arrays.
    function calculateHash(address[] memory _addresses) public pure returns (bytes32) {
        return keccak256(abi.encodePacked(_addresses));
    }
}

/// @title DailyLimit provides daily limit functionality
contract DailyLimit is OptOutableMonolith2FA, SelfCallableOwnable, TokenWhitelistable {
    using SafeMath for uint256;

    event InitializedDailyLimit(uint256 _amount, uint256 _nextReset);
    event SetDailyLimit(address _sender, uint256 _amount);
    event SubmittedDailyLimitUpdate(uint256 _amount);
    event UpdatedAvailableDailyLimit(uint256 _amount, uint256 _nextReset);

    uint256 private _dailyLimit; // The daily limit amount.
    uint256 private _available; // The remaining amount available for spending in the current 24-hour window.
    uint256 private _pendingLimit; // The pending new limit value requested in the latest limit update submission.
    uint256 private _resetTimestamp; // Denotes the future timestamp when the available daily limit is due to reset again.

    // @dev This initializes the daily spend limit using the stablecoin defined in the tokenWhitelist
    // @param _limit is base units of the stablecoin defined in the tokenWhitelist
    // @param _tokenWhitelistNode is the node that points to our tokenWhitelist
    function _initializeDailyLimit(uint256 _limit, bytes32 _tokenWhitelistNode) internal initializer {
        _initializeTokenWhitelistable(_tokenWhitelistNode);
        (, uint256 stablecoinMagnitude, , , , , ) = _getStablecoinInfo();
        require(stablecoinMagnitude > 0, "no stablecoin");
        uint256 limitBaseUnits = _limit * stablecoinMagnitude;
        _dailyLimit = limitBaseUnits;
        _available = limitBaseUnits;
        _pendingLimit = limitBaseUnits;
        _resetTimestamp = now.add(24 hours);
        emit InitializedDailyLimit(limitBaseUnits, _resetTimestamp);
    }

    /// @dev Confirm pending set daily limit operation.
    function confirmDailyLimitUpdate(uint256 _amount) external only2FA {
        // Require that pending and confirmed limits are the same.
        require(_pendingLimit == _amount, "confirmed or submitted limit mismatch");
        // The new limit should be always higher then the current one otherwise no 2FA would be needed
        // this is done to avoid abuse e.g. resetting the current daily limit and thus resetting the available amount
        require(_amount > _dailyLimit, "limit should be greater than current one");
        // Increase the available amount...
        _available = _amount;
        // ...and reset the 24-hour window
        _resetTimestamp = now.add(24 hours);
        emit UpdatedAvailableDailyLimit(_available, _resetTimestamp);
        // Set daily limit based on the pending value.
        _setLimit(_pendingLimit);
    }

    /// @dev Returns the active daily limit change.
    function dailyLimitPending() external view returns (uint256) {
        return _pendingLimit;
    }

    /// @dev Get the daily limit value.
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

    /// @dev Submit a daily transfer limit update for non-whitelisted addresses.
    /// @param _amount is the daily limit amount in stablecoin base units.
    function submitDailyLimitUpdate(uint256 _amount) external onlyOwnerOrSelf {
        // Assign the provided amount to pending daily limit.
        _pendingLimit = _amount;
        // If the new limit is lower, then there is no need for 2FA.
        if (_amount <= _dailyLimit) {
            // Decrease the available amount if the new limit is lower than it
            if (_amount < _available) {
                _available = _amount;
                emit UpdatedAvailableDailyLimit(_available, _resetTimestamp);
            }
            _setLimit(_amount);
        } else {
            emit SubmittedDailyLimitUpdate(_amount);
        }
    }

    /// @dev Use up amount within the daily limit. Will fail if amount is larger than available limit.
    function _enforceDailyLimit(address _asset, uint256 _amount) internal {
        // Convert token amount to stablecoin value.
        uint256 stablecoinValue = convertToStablecoin(_asset, _amount);

        // Account for the limit daily reset.
        _updateAvailableDailyLimit();
        require(_available >= stablecoinValue, "available smaller than amount");
        _available = _available.sub(stablecoinValue);
        emit UpdatedAvailableDailyLimit(_available, _resetTimestamp);
    }

    /// @dev Modify the daily and available limits based on the provided value.
    /// @dev _amount is the daily limit amount in stablecoin base units.
    function _setLimit(uint256 _amount) private {
        // Set the daily limit to the provided amount.
        _dailyLimit = _amount;
        emit SetDailyLimit(msg.sender, _dailyLimit);
    }

    /// @dev Update available limit based on the daily reset.
    function _updateAvailableDailyLimit() private {
        if (now > _resetTimestamp) {
            // Update the current timestamp.
            _resetTimestamp = now.add(24 hours);
            // Set the available limit to the current daily limit.
            _available = _dailyLimit;
            emit UpdatedAvailableDailyLimit(_available, _resetTimestamp);
        }
    }

    /// @dev Convert ether or ERC20 token amount to the corresponding stablecoin amount.
    /// @param _token ERC20 token contract address.
    /// @param _amount amount of token in base units.
    /// @return the equivalent amount in stablecoin base units & 0 if the token is not available.
    function convertToStablecoin(address _token, uint256 _amount) public view returns (uint256) {
        // avoid the unnecessary calculations if the token to be loaded is the stablecoin itself
        if (_token == _stablecoin()) {
            return _amount;
        }

        uint256 amountToSend = _amount;

        // convert token amount to ETH first (0x0 represents ether)
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
}

/// @title Asset wallet with extra security features, gas top up management and card integration.
contract Wallet is ENSResolvable, AddressWhitelist, DailyLimit, IERC165, Transferrable, Balanceable {
    using Address for address;
    using ECDSA for bytes32;
    using SafeERC20 for IERC20;
    using SafeMath for uint256;

    event ExecutedRelayedTransaction(bytes _data, bytes _returnData, bool _privileged);
    event ExecutedTransaction(address _destination, uint256 _value, bytes _data, bytes _returnData);
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

    /// @dev this is an internal nonce to prevent replay attacks from relayer
    uint256 public relayNonce;

    /// @dev Is the registered ENS node identifying the licence contract.
    bytes32 private _licenceNode;

    /// @dev Initializes the wallet top up limit and the vault contract.
    /// @param _owner_ is the owner account of the wallet contract.
    /// @param _transferable_ indicates whether the contract ownership can be transferred.
    /// @param _ens_ is the address of the ENS registry.
    /// @param _tokenWhitelistNode_ is the ENS name node of the Token whitelist.
    /// @param _controllerNode_ is the ENS name node of the Controller contract.
    /// @param _licenceNode_ is the ENS name node of the Licence contract.
    /// @param _dailyLimit_ is the initial daily limit.
    function initializeWallet(
        address payable _owner_,
        bool _transferable_,
        address _ens_,
        bytes32 _tokenWhitelistNode_,
        bytes32 _controllerNode_,
        bytes32 _licenceNode_,
        uint256 _dailyLimit_
    ) external initializer {
        _initialize2FA();
        _initializeENSResolvable(_ens_);
        _initializeControllable(_controllerNode_);
        _initializeOwnable(_owner_, _transferable_);
        _initializeDailyLimit(_dailyLimit_, _tokenWhitelistNode_);
        _licenceNode = _licenceNode_;
    }

    /// @dev Checks if the value is not zero.
    modifier isNotZero(uint256 _value) {
        require(_value != 0, "value=0");
        _;
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

    /// @dev This function allows for the controller to relay transactions on the owner's behalf,
    /// the relayed message has to be signed by the owner.
    /// @param _nonce only used for relayed transactions, must match the wallet's relayNonce.
    /// @param _data abi encoded data payload.
    /// @param _signature signed prefix + data.
    function _executeRelayedTransaction(
        uint256 _nonce,
        bytes calldata _data,
        bytes calldata _signature,
        bool _privileged
    ) private {
        // Expecting prefixed data ("monolith:") indicating relayed transaction...
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
        batchExecuteTransaction(_data);
        privileged = false;

        emit ExecutedRelayedTransaction(_data, "", _privileged);
    }

    /// @dev This returns the balance of the contract for any ERC20 token or ETH.
    /// @param _asset is the address of an ERC20 token or 0x0 for ETH.
    function getBalance(address _asset) external view returns (uint256) {
        return _balance(_asset);
    }

    /// @dev This allows the user to cancel a transaction that was unexpectedly delayed by the relayer
    function increaseRelayNonce() external onlyOwner {
        _increaseRelayNonce();
    }

    /// @dev This bumps the relayNonce and emits an event accordingly
    function _increaseRelayNonce() internal {
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
        // Check against the daily spent limit and update accordingly, require that the value is under remaining limit.
        _enforceDailyLimit(_asset, _amount);
        // Get the TKN licenceAddress from ENS
        address licenceAddress = _ensResolve(_licenceNode);
        if (_asset != address(0)) {
            IERC20(_asset).safeApprove(licenceAddress, _amount);
            ILicence(licenceAddress).load(_asset, _amount);
        } else {
            ILicence(licenceAddress).load{value: _amount}(_asset, _amount);
        }

        emit LoadedTokenCard(_asset, _amount);
    }

    /// @dev Checks for interface support based on ERC165.
    function supportsInterface(bytes4 _interfaceID) external override view returns (bool) {
        return _interfaceID == _ERC165_INTERFACE_ID;
    }

    /// @dev Refill owner's gas balance, revert if the transaction amount is too large
    /// @param _amount is the amount of ether to transfer to the owner account in wei.
    function topUpGas(uint256 _amount) external isNotZero(_amount) onlyOwnerOr2FA {
        // Check contract balance is sufficient for the operation
        require(address(this).balance > _amount, "balance not sufficient");
        // Check against the daily spent limit and update accordingly, require that the value is under remaining limit.
        _enforceDailyLimit(address(0), _amount);
        // Then perform the transfer
        owner().transfer(_amount);
        // Emit the gas top up event.
        emit ToppedUpGas(msg.sender, owner(), _amount);
    }

    /// @dev This function allows for the wallet to send a batch of transactions instead of one,
    /// it calls executeTransaction() so that the daily limit is enforced.
    /// @param _transactionBatch data encoding the transactions to be sent,
    /// following executeTransaction's format i.e. (destination, value, data)
    function batchExecuteTransaction(bytes memory _transactionBatch) public onlyOwnerOr2FA {
        uint256 batchLength = _transactionBatch.length + 32; // because the pos starts from 32
        uint256 remainingBytesLength = _transactionBatch.length; // remaining bytes to be processed
        uint256 pos = 32; //the first 32 bytes denote the byte array length, start from actual data

        address destination; // destination address
        uint256 value; // transaction value
        uint256 dataLength; // external call data length
        bytes memory data; // call data

        while (pos < batchLength) {
            // there should always be at least 84 bytes remaining: the minimum #bytes required to encode a Tx
            remainingBytesLength = remainingBytesLength.sub(84);
            assembly {
                // shift right by 96 bits (256 - 160) to get the destination address (and zero the excessive bytes)
                destination := shr(96, mload(add(_transactionBatch, pos)))
                // get value: pos + 20 bytes (destination address)
                value := mload(add(_transactionBatch, add(pos, 20)))
                // get data: pos + 20 (destination address) + 32 (value) bytes
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
            executeTransaction(destination, value, data);
        }
    }

    /// @dev This function allows for the owner to send any transaction from the Wallet to arbitrary addresses
    /// @param _destination address of the transaction
    /// @param _value ETH amount in wei
    /// @param _data transaction payload binary
    function executeTransaction(
        address _destination,
        uint256 _value,
        bytes memory _data
    ) public onlyOwnerOr2FA returns (bytes memory) {
        // If value is send across as a part of this executeTransaction, this will be sent to any payable
        // destination. As a result enforceLimit if destination is not whitelisted.
        if (!whitelistMap[_destination] && !privileged) {
            // enforce daily limit, 0x0 denotes ETH.
            _enforceDailyLimit(address(0), _value);
        }

        // Check if the destination is a Contract and it is one of our supported tokens
        if (address(_destination).isContract() && _isTokenAvailable(_destination)) {
            // to is the recipient's address and amount is the value to be transferred
            address to;
            uint256 amount;
            (to, amount) = _getERC20RecipientAndAmount(_destination, _data);
            if (!whitelistMap[to] && !privileged) {
                // Convert token amount to stablecoin value.
                // If the address (of the token contract, e.g) is not in the TokenWhitelist used by the convert method
                // ...then stablecoinValue will be zero
                _enforceDailyLimit(_destination, amount);
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

        (bool success, bytes memory returnData) = _destination.call{value: _value}(_data);
        require(success, string(returnData));

        emit ExecutedTransaction(_destination, _value, _data, returnData);
        // returns all of the bytes returned by _destination contract
        return returnData;
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
            _enforceDailyLimit(_asset, _amount);
        }
        // Transfer token or ether based on the provided address.
        _safeTransfer(_to, _asset, _amount);
        // Emit the transfer event.
        emit Transferred(_to, _asset, _amount);
    }
}
