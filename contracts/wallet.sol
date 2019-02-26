/**
 *  The Consumer Contract Wallet
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

import "./licence.sol";
import "./internals/ownable.sol";
import "./internals/controllable.sol";
import "./internals/tokenWhitelistable.sol";
import "./externals/ens/PublicResolver.sol";
import "./externals/SafeMath.sol";
import "./externals/ERC20.sol";
import "./externals/ERC165.sol";

/// @title Whitelist provides payee-whitelist functionality.
contract Whitelist is Controllable, Ownable {
    event AddedToWhitelist(address _sender, address[] _addresses);
    event SubmittedWhitelistAddition(address[] _addresses, bytes32 _hash);
    event CancelledWhitelistAddition(address _sender, bytes32 _hash);

    event RemovedFromWhitelist(address _sender, address[] _addresses);
    event SubmittedWhitelistRemoval(address[] _addresses, bytes32 _hash);
    event CancelledWhitelistRemoval(address _sender, bytes32 _hash);

    mapping(address => bool) public isWhitelisted;
    address[] private _pendingWhitelistAddition;
    address[] private _pendingWhitelistRemoval;
    bool public submittedWhitelistAddition;
    bool public submittedWhitelistRemoval;
    bool public initializedWhitelist;

    /// @dev Check if the provided addresses contain the owner or the zero-address address.
    modifier hasNoOwnerOrZeroAddress(address[] _addresses) {
        for (uint i = 0; i < _addresses.length; i++) {
            require(_addresses[i] != owner(), "provided whitelist contains the owner address");
            require(_addresses[i] != address(0), "provided whitelist contains the zero address");
        }
        _;
    }

    /// @dev Check that neither addition nor removal operations have already been submitted.
    modifier noActiveSubmission() {
        require(!submittedWhitelistAddition && !submittedWhitelistRemoval, "whitelist operation has already been submitted");
        _;
    }

    /// @dev Getter for pending addition array.
    function pendingWhitelistAddition() external view returns (address[]) {
        return _pendingWhitelistAddition;
    }

    /// @dev Getter for pending removal array.
    function pendingWhitelistRemoval() external view returns (address[]) {
        return _pendingWhitelistRemoval;
    }

    /// @dev Getter for pending addition/removal array hash.
    function pendingWhitelistHash(address[] _pendingWhitelist) public pure returns (bytes32) {
        return keccak256(abi.encodePacked(_pendingWhitelist));
    }

    /// @dev Add initial addresses to the whitelist.
    /// @param _addresses are the Ethereum addresses to be whitelisted.
    function initializeWhitelist(address[] _addresses) external onlyOwner hasNoOwnerOrZeroAddress(_addresses) {
        // Require that the whitelist has not been initialized.
        require(!initializedWhitelist, "whitelist has already been initialized");
        // Add each of the provided addresses to the whitelist.
        for (uint i = 0; i < _addresses.length; i++) {
            isWhitelisted[_addresses[i]] = true;
        }
        initializedWhitelist = true;
        // Emit the addition event.
        emit AddedToWhitelist(msg.sender, _addresses);
    }

    /// @dev Add addresses to the whitelist.
    /// @param _addresses are the Ethereum addresses to be whitelisted.
    function submitWhitelistAddition(address[] _addresses) external onlyOwner noActiveSubmission hasNoOwnerOrZeroAddress(_addresses) {
        // Require that the whitelist has been initialized.
        require(initializedWhitelist, "whitelist has not been initialized");
        // Set the provided addresses to the pending addition addresses.
        _pendingWhitelistAddition = _addresses;
        // Flag the operation as submitted.
        submittedWhitelistAddition = true;
        // Emit the submission event.
        emit SubmittedWhitelistAddition(_addresses, pendingWhitelistHash(_pendingWhitelistAddition));
    }

    /// @dev Confirm pending whitelist addition.
    function confirmWhitelistAddition(bytes32 _hash) external onlyController {
        // Require that the whitelist addition has been submitted.
        require(submittedWhitelistAddition, "whitelist addition has not been submitted");

        // Require that confirmation hash and the hash of the pending whitelist addition match
        require(_hash == pendingWhitelistHash(_pendingWhitelistAddition), "hash of the pending whitelist addition do not match");

        // Whitelist pending addresses.
        for (uint i = 0; i < _pendingWhitelistAddition.length; i++) {
            isWhitelisted[_pendingWhitelistAddition[i]] = true;
        }
        // Emit the addition event.
        emit AddedToWhitelist(msg.sender, _pendingWhitelistAddition);
        // Reset pending addresses.
        delete _pendingWhitelistAddition;
        // Reset the submission flag.
        submittedWhitelistAddition = false;
    }

    /// @dev Cancel pending whitelist addition.
    function cancelWhitelistAddition(bytes32 _hash) external onlyController {
        // Require that confirmation hash and the hash of the pending whitelist addition match
        require(_hash == pendingWhitelistHash(_pendingWhitelistAddition), "hash of the pending whitelist addition does not match");
        // Reset pending addresses.
        delete _pendingWhitelistAddition;
        // Reset the submitted operation flag.
        submittedWhitelistAddition = false;
        // Emit the cancellation event.
        emit CancelledWhitelistAddition(msg.sender, _hash);
    }

    /// @dev Remove addresses from the whitelist.
    /// @param _addresses are the Ethereum addresses to be removed.
    function submitWhitelistRemoval(address[] _addresses) external onlyOwner noActiveSubmission {
        // Add the provided addresses to the pending addition list.
        _pendingWhitelistRemoval = _addresses;
        // Flag the operation as submitted.
        submittedWhitelistRemoval = true;
        // Emit the submission event.
        emit SubmittedWhitelistRemoval(_addresses, pendingWhitelistHash(_pendingWhitelistRemoval));
    }

    /// @dev Confirm pending removal of whitelisted addresses.
    function confirmWhitelistRemoval(bytes32 _hash) external onlyController {
        // Require that the pending whitelist is not empty and the operation has been submitted.
        require(submittedWhitelistRemoval, "whitelist removal has not been submitted");
        require(_pendingWhitelistRemoval.length > 0, "pending whitelist removal is empty");
        // Require that confirmation hash and the hash of the pending whitelist removal match
        require(_hash == pendingWhitelistHash(_pendingWhitelistRemoval), "hash of the pending whitelist removal does not match the confirmed hash");
        // Remove pending addresses.
        for (uint i = 0; i < _pendingWhitelistRemoval.length; i++) {
            isWhitelisted[_pendingWhitelistRemoval[i]] = false;
        }
        // Emit the removal event.
        emit RemovedFromWhitelist(msg.sender, _pendingWhitelistRemoval);
        // Reset pending addresses.
        delete _pendingWhitelistRemoval;
        // Reset the submission flag.
        submittedWhitelistRemoval = false;
    }

    /// @dev Cancel pending removal of whitelisted addresses.
    function cancelWhitelistRemoval(bytes32 _hash) external onlyController {
        // Require that confirmation hash and the hash of the pending whitelist removal match
        require(_hash == pendingWhitelistHash(_pendingWhitelistRemoval), "hash of the pending whitelist removal does not match");
        // Reset pending addresses.
        delete _pendingWhitelistRemoval;
        // Reset the submitted operation flag.
        submittedWhitelistRemoval = false;
        // Emit the cancellation event.
        emit CancelledWhitelistRemoval(msg.sender, _hash);
    }
}

contract DailyLimitable {
    using SafeMath for uint256;

    struct DailyLimit {
      uint dailyLimit;
      uint available;
      uint limitDay;

      uint pending;
      bool submitted;
      bool initialized;
    }

    /// @dev Returns the available daily balance - accounts for daily limit reset.
    /// @return amount of ether in wei.
    function dl_available(DailyLimit storage dl) internal view returns (uint) {
        if (now > dl.limitDay + 24 hours) {
            return dl.dailyLimit;
        } else {
            return dl.available;
        }
    }

    // @dev Use up amount within the daily limit. Will fail if amount is larger than daily limit.
    function dl_useAmount(DailyLimit storage dl, uint _amount) internal {
        dl_updateAvailable(dl);
        require(dl.available >= _amount, "available has to be greater or equal to use amount");
        dl.available = dl.available.sub(_amount);
    }


    /// @dev Update available spend limit based on the daily reset.
    function dl_updateAvailable(DailyLimit storage dl) internal {
        if (now > dl.limitDay.add(24 hours)) {
            // Advance the current day by how many days have passed.
            uint extraDays = now.sub(dl.limitDay).div(24 hours);
            dl.limitDay = dl.limitDay.add(extraDays.mul(24 hours));
            // Set the available limit to the current spend limit.
            dl.available = dl.dailyLimit;
            //TODO i guess we need to do this twice
        }
    }

    /// @dev Modify the spend limit and spend available based on the provided value.
    /// @dev _amount is the daily limit amount in wei.
    function dl_modifyLimit(DailyLimit storage dl, uint _amount) private {
        // Account for the spend limit daily reset.
        dl_updateAvailable(dl);
        // Set the daily limit to the provided amount.
        dl.dailyLimit = _amount;
        // Lower the available limit if it's higher than the new daily limit.
        if (dl.available > dl.dailyLimit) {
            dl.available = dl.dailyLimit;
        }
    }

    /// @dev Initialize a daily limit.
    /// @param _amount is the daily limit amount in wei.
    function dl_initialize(DailyLimit storage dl, uint _amount) internal {
        // Require that the spend limit has not been initialized.
        require(!dl.initialized, "limit has already been initialized");
        // Modify spend limit based on the provided value.
        dl_modifyLimit(dl, _amount);
        // Flag the operation as initialized.
        dl.initialized = true;
    }


    /// @dev Submit a daily limit change, needs to be confirmed.
    /// @param _amount is the daily limit amount in wei.
    function dl_submit(DailyLimit storage dl, uint _amount) internal {
        // Require that the spend limit has been initialized.
        require(dl.initialized, "limit has not been initialized");
        // Require that the operation has been submitted.
        require(!dl.submitted, "limit has already been submitted");
        // Assign the provided amount to pending daily limit change.
        dl.pending = _amount;
        // Flag the operation as submitted.
        dl.submitted = true;
    }

    /// @dev Confirm pending set daily limit operation.
    function dl_confirm(DailyLimit storage dl, uint _amount) internal {
        // Require that the operation has been submitted.
        require(dl.submitted, "limit has not been submitted");
        // Require that pending and confirmed spend limit are the same
        require(dl.pending == _amount, "confirmed and submitted limits dont match");
        // Modify spend limit based on the pending value.
        dl_modifyLimit(dl, dl.pending);
        // Reset the submission flag.
        dl.submitted = false;
        // Reset pending daily limit.
        dl.pending = 0;
    }

    /// @dev Cancel pending set daily limit operation.
    function dl_cancel(DailyLimit storage dl, uint _amount) internal {
        // Require that pending and confirmed spend limit are the same
        require(dl.pending == _amount, "confirmed and cancelled limits dont match");
        // Reset pending daily limit.
        dl.pending = 0;
        // Reset the submitted operation flag.
        dl.submitted = false;
    }

}


//// @title SpendLimit provides daily spend limit functionality.
contract SpendLimit is Controllable, Ownable, DailyLimitable {
    event SetSpendLimit(address _sender, uint _amount);
    event SubmittedSpendLimitChange(uint _amount);
    event CancelledSpendLimitChange(address _sender, uint _amount);

    DailyLimit internal _spendLimit;

    /// @dev Constructor initializes the daily spend limit in wei.
    constructor(uint spendLimit) internal {
        _spendLimit = DailyLimit(spendLimit, spendLimit, now, 0, false, false);
    }

    /// @dev Initialize a daily spend (aka transfer) limit for non-whitelisted addresses.
    /// @param _amount is the daily limit amount in wei.
    function initializeSpendLimit(uint _amount) external onlyOwner {
        dl_initialize(_spendLimit, _amount);
        emit SetSpendLimit(msg.sender, _amount);
    }

    /// @dev Set a daily transfer limit for non-whitelisted addresses.
    /// @param _amount is the daily limit amount in wei.
    function submitSpendLimit(uint _amount) external onlyOwner {
        dl_submit(_spendLimit, _amount);
        emit SubmittedSpendLimitChange(_amount);
    }

    /// @dev Confirm pending set daily limit operation.
    function confirmSpendLimit(uint _amount) external onlyController {
        dl_confirm(_spendLimit, _amount);
        emit SetSpendLimit(msg.sender, _amount);
    }

    /// @dev Cancel pending set daily limit operation.
    function cancelSpendLimit(uint _amount) external onlyController {
        dl_cancel(_spendLimit, _amount);
        emit CancelledSpendLimitChange(msg.sender, _amount);
    }

    function spendAvailable() public view returns (uint) {
        return dl_available(_spendLimit);
    }

    function spendLimit() public view returns (uint) {
        return _spendLimit.dailyLimit;
    }

    function initializedSpendLimit() public view returns (bool) {
        return _spendLimit.initialized;
    }

    function submittedSpendLimit() public view returns (bool) {
      return _spendLimit.submitted;
    }

    function pendingSpendLimit() public view returns (uint) {
      return _spendLimit.pending;
    }

}


//// @title Asset store with extra security features.
contract Vault is Whitelist, SpendLimit, ERC165, TokenWhitelistable {

    using SafeMath for uint256;

    event Received(address _from, uint _amount);
    event Transferred(address _to, address _asset, uint _amount);

    /// @dev Supported ERC165 interface ID.
    bytes4 private constant _ERC165_INTERFACE_ID = 0x01ffc9a7; // solium-disable-line uppercase

    /// @dev Constructor initializes the vault with an owner address and spend limit. It also sets up the oracle and controller contracts.
    /// @param _owner is the owner account of the wallet contract.
    /// @param _transferable indicates whether the contract ownership can be transferred.
    /// @param _ens is the ENS public registry contract address.
    /// @param _tokenWhitelistName is the ENS name of the Token whitelist.
    /// @param _controllerName is the ENS name of the controller.
    /// @param _spendLimit is the initial spend limit.
    constructor(address _owner, bool _transferable, address _ens, bytes32 _tokenWhitelistName, bytes32 _controllerName, uint _spendLimit) SpendLimit(_spendLimit) Ownable(_owner, _transferable) Controllable(_ens, _controllerName) TokenWhitelistable(_ens, _tokenWhitelistName) public {
    }

    /// @dev Checks if the value is not zero.
    modifier isNotZero(uint _value) {
        require(_value != 0, "provided value cannot be zero");
        _;
    }

    /// @dev Ether can be deposited from any source, so this contract must be payable by anyone.
    function() public payable {
        require(msg.data.length == 0);
        emit Received(msg.sender, msg.value);
    }

    /// @dev Returns the amount of an asset owned by the contract.
    /// @param _asset address of an ERC20 token or 0x0 for ether.
    /// @return balance associated with the wallet address in wei.
    function balance(address _asset) external view returns (uint) {
        if (_asset != address(0)) {
            return ERC20(_asset).balanceOf(this);
        } else {
            return address(this).balance;
        }
    }

    /// @dev Convert ERC20 token amount to the corresponding ether amount (used by the wallet contract).
    /// @param _token ERC20 token contract address.
    /// @param _amount amount of token in base units.
    function convert(address _token, uint _amount) public view returns (uint) {
        // Store the token in memory to save map entry lookup gas.
        ( , uint256 magnitude, uint256 rate, bool available, , ) = _getTokenInfo(_token);
        // If the token exists require that its rate is not zero
        if (available) {
            require(rate != 0, "token rate is 0");
            // Safely convert the token amount to ether based on the exchange rate.
            // return the value, the token is, AT LEAST, protected
            return _amount.mul(rate).div(magnitude);
        }
        // this returns a 0/'false' to imply that the token is not protected
        return 0;
    }


    /// @dev Transfers the specified asset to the recipient's address.
    /// @param _to is the recipient's address.
    /// @param _asset is the address of an ERC20 token or 0x0 for ether.
    /// @param _amount is the amount of assets to be transferred in base units.
    function transfer(address _to, address _asset, uint _amount) external onlyOwner isNotZero(_amount) {
        // Checks if the _to address is not the zero-address
        require(_to != address(0), "_to address cannot be set to 0x0");

        // If address is not whitelisted, take daily limit into account.
        if (!isWhitelisted[_to]) {
            //initialize ether value in case the asset is ETH
            uint etherValue = _amount;
            // Convert token amount to ether value if asset is an ERC20 token.
            if (_asset != address(0)) {
                etherValue = convert(_asset, _amount);
            }
            // Check against the daily spent limit and update accordingly
            // Require that the value is under remaining limit.
            dl_useAmount(_spendLimit, etherValue);
        }
        // Transfer token or ether based on the provided address.
        if (_asset != address(0)) {
            require(ERC20(_asset).transfer(_to, _amount), "ERC20 token transfer was unsuccessful");
        } else {
            _to.transfer(_amount);
        }
        // Emit the transfer event.
        emit Transferred(_to, _asset, _amount);
    }

    /// @dev Checks for interface support based on ERC165.
    function supportsInterface(bytes4 interfaceID) external view returns (bool) {
        return interfaceID == _ERC165_INTERFACE_ID;
    }
}


//// @title Asset wallet with extra security features, gas top up management and card integration.
contract Wallet is Vault {
    event SetGasTopUpLimit(address _sender, uint _amount);
    event SubmittedGasTopUpLimitChange(uint _amount);
    event CancelledGasTopUpLimitChange(address _sender, uint _amount);

    event ToppedUpGas(address _sender, address _owner, uint _amount);

    event SetLoadLimit(address _sender, uint _amount);
    event SubmittedLoadLimitChange(uint _amount);
    event CancelledLoadLimitChange(address _sender, uint _amount);

    event LoadedTokenCard(address _asset, uint _amount);

    uint constant private MINIMUM_GAS_TOPUP_LIMIT = 1 finney;
    uint constant private MAXIMUM_GAS_TOPUP_LIMIT = 500 finney;

    uint constant private MINIMUM_LOAD_LIMIT = 1 finney;
    uint constant private MAXIMUM_LOAD_LIMIT = 101 ether;

    /// @dev Is the registered ENS name of the oracle contract.
    bytes32 private _licenceNode;

    /// @dev ENS points to the ENS registry smart contract.
    ENS internal _ENS;

    DailyLimit internal _gasTopUpLimit;
    DailyLimit internal _loadLimit;

    /// @dev Constructor initializes the wallet top up limit and the vault contract.
    /// @param _owner is the owner account of the wallet contract.
    /// @param _transferable indicates whether the contract ownership can be transferred.
    /// @param _ens is the address of the ENS.
    /// @param _oracleName is the ENS name of the Oracle.
    /// @param _controllerName is the ENS name of the Controller.
    /// @param _licenceName is the ENS name of the licence.
    /// @param _spendLimit is the initial spend limit.
    constructor(address _owner, bool _transferable, address _ens, bytes32 _oracleName, bytes32 _controllerName, bytes32 _licenceName, uint _spendLimit) Vault(_owner, _transferable, _ens, _oracleName, _controllerName, _spendLimit) public {
        _gasTopUpLimit = DailyLimit(MAXIMUM_GAS_TOPUP_LIMIT, MAXIMUM_GAS_TOPUP_LIMIT, now, 0, false, false);
        _loadLimit = DailyLimit(MAXIMUM_LOAD_LIMIT, MAXIMUM_LOAD_LIMIT, now, 0, false, false);
        _licenceNode = _licenceName;
        _ENS = ENS(_ens);
    }

    /// @dev Initialize a daily gas top up limit.
    /// @param _amount is the top up gas amount in wei.
    function initializeGasTopUpLimit(uint _amount) external onlyOwner {
        require(MINIMUM_GAS_TOPUP_LIMIT <= _amount && _amount <= MAXIMUM_GAS_TOPUP_LIMIT, "gas top up amount is outside the min/max range");
        dl_initialize(_gasTopUpLimit, _amount);
        emit SetGasTopUpLimit(msg.sender, _amount);
    }

    /// @dev Initialize a daily card load limit.
    /// @param _amount is the card load amount in wei.
    function initializeLoadLimit(uint _amount) external onlyOwner {
        require(MINIMUM_LOAD_LIMIT <= _amount && _amount <= MAXIMUM_LOAD_LIMIT, "card load amount is outside the min/max range");
        dl_initialize(_loadLimit, _amount);
        emit SetLoadLimit(msg.sender, _amount);
    }

    /// @dev Set a daily top up gas limit.
    /// @param _amount is the daily top up gas limit amount in wei.
    function submitGasTopUpLimit(uint _amount) external onlyOwner {
        require(MINIMUM_GAS_TOPUP_LIMIT <= _amount && _amount <= MAXIMUM_GAS_TOPUP_LIMIT, "gas top up amount is outside the min/max range");
        dl_submit(_gasTopUpLimit, _amount);
        emit SubmittedGasTopUpLimitChange(_amount);
    }

    /// @dev Set a daily load limit.
    /// @param _amount is the daily load limit amount in wei.
    function submitLoadLimit(uint _amount) external onlyOwner {
        require(MINIMUM_LOAD_LIMIT <= _amount && _amount <= MAXIMUM_LOAD_LIMIT, "card load amount is outside the min/max range");
        dl_submit(_loadLimit, _amount);
        emit SubmittedLoadLimitChange(_amount);
    }

    /// @dev Confirm pending set top up gas limit operation.
    function confirmGasTopUpLimit(uint _amount) external onlyController {
        dl_confirm(_gasTopUpLimit, _amount);
        emit SetGasTopUpLimit(msg.sender, _amount);
    }

    /// @dev Confirm pending set load limit operation.
    function confirmLoadLimit(uint _amount) external onlyController {
        dl_confirm(_loadLimit, _amount);
        emit SetLoadLimit(msg.sender, _amount);
    }

    /// @dev Cancel pending set top up gas limit operation.
    function cancelGasTopUpLimit(uint _amount) external onlyController {
        dl_cancel(_gasTopUpLimit, _amount);
        emit CancelledGasTopUpLimitChange(msg.sender, _amount);
    }

    /// @dev Cancel pending set load limit operation.
    function cancelLoadLimit(uint _amount) external onlyController {
        dl_cancel(_loadLimit, _amount);
        emit CancelledLoadLimitChange(msg.sender, _amount);
    }

    /// @dev Refill owner's gas balance, revert if the transaction amount is too large
    /// @param _amount is the amount of ether to transfer to the owner account in wei.
    function topUpGas(uint _amount) external isNotZero(_amount) {
        // Require that the sender is either the owner or a controller.
        require(_isOwner() || _isController(msg.sender), "sender is neither an owner nor a controller");

        dl_useAmount(_gasTopUpLimit, _amount);

        owner().transfer(_amount);

        // Emit the gas top up event.
        emit ToppedUpGas(msg.sender, owner(), _amount);
    }

    /// @dev Load a token card with the specified asset amount.
    /// the amount send should be inclusive of the percent licence.
    /// @param _asset is the address of an ERC20 token or 0x0 for ether.
    /// @param _amount is the amount of assets to be transferred in base units.
    function loadTokenCard(address _asset, uint _amount) external payable onlyOwner {

      address licenceAddress = PublicResolver(_ENS.resolver(_licenceNode)).addr(_licenceNode);

      if (_asset != address(0)) {
          //check if token is allowed to be used for loading the card
          require(_isTokenLoadable(_asset), "token not loadable");
          // Convert token amount to ether value.
          uint etherValue = convert(_asset, _amount);
          // Check against the daily spent limit and update accordingly, require that the value is under remaining limit.
          dl_useAmount(_loadLimit, etherValue);
          require(ERC20(_asset).approve(licenceAddress, _amount), "ERC20 token approval was unsuccessful");
          ILicence(licenceAddress).load(_asset, _amount);
      } else {
          //_amount is in wei, require that the value is under remaining limit.
          dl_useAmount(_loadLimit, _amount);
          ILicence(licenceAddress).load.value(_amount)(_asset, _amount);
      }

        emit LoadedTokenCard(_asset, _amount);
    }

    function gasTopUpLimit() public view returns (uint) {
        return _gasTopUpLimit.dailyLimit;
    }

    function gasTopUpAvailable() public view returns (uint) {
        return dl_available(_gasTopUpLimit);
    }

    function initializedGasTopUpLimit() public view returns (bool) {
        return _gasTopUpLimit.initialized;
    }

    function submittedGasTopUpLimit() public view returns (bool) {
      return _gasTopUpLimit.submitted;
    }

    function pendingGasTopUpLimit() public view returns (uint) {
      return _gasTopUpLimit.pending;
    }


    function loadLimit() public view returns (uint) {
        return _loadLimit.dailyLimit;
    }

    function loadAvailable() public view returns (uint) {
        return dl_available(_loadLimit);
    }

    function initializedLoadLimit() public view returns (bool) {
        return _loadLimit.initialized;
    }

    function submittedLoadLimit() public view returns (bool) {
      return _loadLimit.submitted;
    }

    function pendingLoadLimit() public view returns (uint) {
      return _loadLimit.pending;
    }

}
