/**
 *  TokenWallet - The Consumer Contract Wallet
 *  Copyright (C) 2018 Token Group Ltd
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

import "./oracle.sol";
import "./internals/ownable.sol";
import "./internals/controllable.sol";
import "./externals/ens/PublicResolver.sol";

/// @title ERC20 interface is a subset of the ERC20 specification.
interface ERC20 {
    function transfer(address, uint) external returns (bool);
    function balanceOf(address) external view returns (uint);
}


/// @title ERC165 interface specifies a standard way of querying if a contract implements an interface.
interface ERC165 {
    function supportsInterface(bytes4) external view returns (bool);
}


/// @title Whitelist provides payee-whitelist functionality.
contract Whitelist is Controllable, Ownable {
    event AddedToWhitelist(address _sender, address[] _addresses);
    event SubmittedWhitelistAddition(address[] _addresses);
    event CancelledWhitelistAddition(address _sender);

    event RemovedFromWhitelist(address _sender, address[] _addresses);
    event SubmittedWhitelistRemoval(address[] _addresses);
    event CancelledWhitelistRemoval(address _sender);

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
            require(_addresses[i] != address(0), "provided whitelist contains the 0x0 zero address");
        }
        _;
    }

    /// @dev Check that neither addition nor removal operations have already been submitted.
    modifier noActiveSubmission() {
        require(!submittedWhitelistAddition && !submittedWhitelistRemoval, "whitelist operation has already been submitted");
        _;
    }

    /// @dev Getter for pending addition array.
    function pendingWhitelistAddition() external view returns(address[]) {
        return _pendingWhitelistAddition;
    }

    /// @dev Getter for pending removal array.
    function pendingWhitelistRemoval() external view returns(address[]) {
        return _pendingWhitelistRemoval;
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
    function submitWhitelistAddition(address[] _addresses) external onlyOwner noActiveSubmission hasNoOwnerOrZeroAddress(_addresses)  {
        // Require that the whitelist has been initialized.
        require(initializedWhitelist, "whitelist has not been initialized");
        // Set the provided addresses to the pending addition addresses.
        _pendingWhitelistAddition = _addresses;
        // Flag the operation as submitted.
        submittedWhitelistAddition = true;
        // Emit the submission event.
        emit SubmittedWhitelistAddition(_addresses);
    }

    /// @dev Confirm pending whitelist addition.
    function confirmWhitelistAddition() external onlyController {
        // Require that the whitelist addition has been submitted.
        require(submittedWhitelistAddition, "whitelist addition has not been submitted");
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
    function cancelWhitelistAddition() external onlyController {
        // Reset pending addresses.
        delete _pendingWhitelistAddition;
        // Reset the submitted operation flag.
        submittedWhitelistAddition = false;
        // Emit the cancellation event.
        emit CancelledWhitelistAddition(msg.sender);
    }

    /// @dev Remove addresses from the whitelist.
    /// @param _addresses are the Ethereum addresses to be removed.
    function submitWhitelistRemoval(address[] _addresses) external onlyOwner noActiveSubmission {
        // Add the provided addresses to the pending addition list.
        _pendingWhitelistRemoval = _addresses;
        // Flag the operation as submitted.
        submittedWhitelistRemoval = true;
        // Emit the submission event.
        emit SubmittedWhitelistRemoval(_addresses);
    }

    /// @dev Confirm pending removal of whitelisted addresses.
    function confirmWhitelistRemoval() external onlyController {
        // Require that the pending whitelist is not empty and the operation has been submitted.
        require(submittedWhitelistRemoval, "whitelist removal has not been submitted");
        require(_pendingWhitelistRemoval.length > 0, "pending whitelist removal is empty");
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
    function cancelWhitelistRemoval() external onlyController {
        // Reset pending addresses.
        delete _pendingWhitelistRemoval;
        // Reset the submitted operation flag.
        submittedWhitelistRemoval = false;
        // Emit the cancellation event.
        emit CancelledWhitelistRemoval(msg.sender);
    }
}


//// @title SpendLimit provides daily spend limit functionality.
contract SpendLimit is Controllable, Ownable {
    event SetSpendLimit(address _sender, uint _amount);
    event SubmittedSpendLimitChange(uint _amount);
    event CancelledSpendLimitChange(address _sender);

    uint public spendLimit;
    uint internal _spendLimitDay;
    uint internal _spendAvailable;

    uint public pendingSpendLimit;
    bool public submittedSpendLimit;
    bool public initializedSpendLimit;

    /// @dev Constructor initializes the daily spend limit in wei.
    constructor(uint _spendLimit) internal {
        spendLimit = _spendLimit;
        _spendLimitDay = now;
        _spendAvailable = spendLimit;
    }

    /// @dev Returns the available daily balance - accounts for daily limit reset.
    /// @return amount of ether in wei.
    function spendAvailable() external view returns (uint) {
        if (now > _spendLimitDay + 24 hours) {
            return spendLimit;
        } else {
            return _spendAvailable;
        }
    }

    /// @dev Initialize a daily spend (aka transfer) limit for non-whitelisted addresses.
    /// @param _amount is the daily limit amount in wei.
    function initializeSpendLimit(uint _amount) external onlyOwner {
        // Require that the spend limit has not been initialized.
        require(!initializedSpendLimit, "spend limit has already been initialized");
        // Modify spend limit based on the provided value.
        modifySpendLimit(_amount);
        // Flag the operation as initialized.
        initializedSpendLimit = true;
        // Emit the set limit event.
        emit SetSpendLimit(msg.sender, _amount);
    }

    /// @dev Set a daily transfer limit for non-whitelisted addresses.
    /// @param _amount is the daily limit amount in wei.
    function submitSpendLimit(uint _amount) external onlyOwner {
        // Require that the spend limit has been initialized.
        require(initializedSpendLimit, "spend limit has not been initialized");
        // Require that the operation has been submitted.
        require(!submittedSpendLimit, "spend limit has already been submitted");
        // Assign the provided amount to pending daily limit change.
        pendingSpendLimit = _amount;
        // Flag the operation as submitted.
        submittedSpendLimit = true;
        // Emit the submission event.
        emit SubmittedSpendLimitChange(_amount);
    }

    /// @dev Confirm pending set daily limit operation.
    function confirmSpendLimit() external onlyController {
        // Require that the operation has been submitted.
        require(submittedSpendLimit, "spend limit has not been submitted");
        // Modify spend limit based on the pending value.
        modifySpendLimit(pendingSpendLimit);
        // Emit the set limit event.
        emit SetSpendLimit(msg.sender, pendingSpendLimit);
        // Reset the submission flag.
        submittedSpendLimit = false;
        // Reset pending daily limit.
        pendingSpendLimit = 0;
    }

    /// @dev Cancel pending set daily limit operation.
    function cancelSpendLimit() external onlyController {
        // Reset pending daily limit.
        pendingSpendLimit = 0;
        // Reset the submitted operation flag.
        submittedSpendLimit = false;
        // Emit the cancellation event.
        emit CancelledSpendLimitChange(msg.sender);
    }

    /// @dev Update available spend limit based on the daily reset.
    function updateSpendAvailable() internal {
        if (now > _spendLimitDay + 24 hours) {
            // Advance the current day by how many days have passed.
            uint extraDays = (now - _spendLimitDay) / 24 hours;
            _spendLimitDay += extraDays * 24 hours;
            // Set the available limit to the current spend limit.
            _spendAvailable = spendLimit;
        }
    }

    /// @dev Modify the spend limit and spend available based on the provided value.
    /// @dev _amount is the daily limit amount in wei.
    function modifySpendLimit(uint _amount) private {
        // Account for the spend limit daily reset.
        updateSpendAvailable();
        // Set the daily limit to the provided amount.
        spendLimit = _amount;
        // Lower the available limit if it's higher than the new daily limit.
        if (_spendAvailable > spendLimit) {
            _spendAvailable = spendLimit;
        }
    }
}


//// @title Asset store with extra security features.
contract Vault is Whitelist, SpendLimit, ERC165 {
    event Received(address _from, uint _amount);
    event Transferred(address _to, address _asset, uint _amount);

    /// @dev Supported ERC165 interface ID.
    bytes4 private constant _ERC165_INTERFACE_ID = 0x01ffc9a7; // solium-disable-line uppercase

    /// @dev ENS points to the ENS registry smart contract.
    ENS private _ENS;
    /// @dev Is the registered ENS name of the oracle contract.
    bytes32 private _node;

    /// @dev Constructor initializes the vault with an owner address and spend limit. It also sets up the oracle and controller contracts.
    /// @param _owner is the owner account of the wallet contract.
    /// @param _transferable indicates whether the contract ownership can be transferred.
    /// @param _ens is the ENS public registry contract address.
    /// @param _oracleName is the ENS name of the Oracle.
    /// @param _controllerName is the ENS name of the controller.
    /// @param _spendLimit is the initial spend limit.
    constructor(address _owner, bool _transferable, address _ens, bytes32 _oracleName, bytes32 _controllerName, uint _spendLimit) SpendLimit(_spendLimit) Ownable(_owner, _transferable) Controllable(_ens, _controllerName) public {
        _ENS = ENS(_ens);
        _node = _oracleName;
    }

    /// @dev Checks if the value is not zero.
    modifier isNotZero(uint _value) {
        require(_value != 0, "provided value cannot be zero");
        _;
    }

    /// @dev Checks if the to address is not the zero-address.
    modifier isNotZeroAddress(address _to) {
        require(_to != address(0), "'to' address cannot be set to 0x0");     
        _;
    }

    /// @dev Ether can be deposited from any source, so this contract must be payable by anyone.
    function() public payable {
        //TODO question: Why is this check here, is it necessary or are we building into a corner?
        require(msg.data.length == 0);
        emit Received(msg.sender, msg.value);
    }

    /// @dev Returns the amount of an asset owned by the contract.
    /// @param _asset address of an ERC20 token or 0x0 for ether.
    /// @return balance associated with the wallet address in wei.
    function balance(address _asset) external view returns (uint) {
        if (_asset != 0x0) {
            return ERC20(_asset).balanceOf(this);
        } else {
            return address(this).balance;
        }
    }

    /// @dev Transfers the specified asset to the recipient's address.
    /// @param _to is the recipient's address.
    /// @param _asset is the address of an ERC20 token or 0x0 for ether.
    /// @param _amount is the amount of tokens to be transferred in base units.
    function transfer(address _to, address _asset, uint _amount) external onlyOwner isNotZero(_amount) isNotZeroAddress(_to) {
        // If address is not whitelisted, take daily limit into account.
        if (!isWhitelisted[_to]) {
            // Update the available spend limit.
            updateSpendAvailable();
            // Convert token amount to ether value.
            uint etherValue;
            if (_asset != 0x0) {
                etherValue = IOracle(PublicResolver(_ENS.resolver(_node)).addr(_node)).convert(_asset, _amount);
            } else {
                etherValue = _amount;
            }
            // Require that the value is under remaining limit.
            require(etherValue <= _spendAvailable, "transfer amount exceeds available spend limit");
            // Update the available limit.
            _spendAvailable -= etherValue;
        }
        // Transfer token or ether based on the provided address.
        if (_asset != 0x0) {
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


//// @title Asset wallet with extra security features and gas top up management.
contract Wallet is Vault {
    event SetTopUpLimit(address _sender, uint _amount);
    event SubmittedTopUpLimitChange(uint _amount);
    event CancelledTopUpLimitChange(address _sender);

    event ToppedUpGas(address _sender, address _owner, uint _amount);

    uint constant private MINIMUM_TOPUP_LIMIT = 1 finney; // solium-disable-line uppercase
    uint constant private MAXIMUM_TOPUP_LIMIT = 500 finney; // solium-disable-line uppercase

    uint public topUpLimit;
    uint private _topUpLimitDay;
    uint private _topUpAvailable;

    uint public pendingTopUpLimit;
    bool public submittedTopUpLimit;
    bool public initializedTopUpLimit;

    /// @dev Constructor initializes the wallet top up limit and the vault contract.
    /// @param _owner is the owner account of the wallet contract.
    /// @param _transferable indicates whether the contract ownership can be transferred.
    /// @param _ens is the address of the ENS.
    /// @param _oracleName is the ENS name of the Oracle.
    /// @param _controllerName is the ENS name of the Controller.
    /// @param _spendLimit is the initial spend limit.
    constructor(address _owner, bool _transferable, address _ens, bytes32 _oracleName, bytes32 _controllerName, uint _spendLimit) Vault(_owner, _transferable, _ens, _oracleName, _controllerName, _spendLimit) public {
        _topUpLimitDay = now;
        topUpLimit = MAXIMUM_TOPUP_LIMIT;
        _topUpAvailable = topUpLimit;
    }

    /// @dev Returns the available daily gas top up balance - accounts for daily limit reset.
    /// @return amount of gas in wei.
    function topUpAvailable() external view returns (uint) {
        if (now > _topUpLimitDay + 24 hours) {
            return topUpLimit;
        } else {
            return _topUpAvailable;
        }
    }

    /// @dev Initialize a daily gas top up limit.
    /// @param _amount is the gas top up amount in wei.
    function initializeTopUpLimit(uint _amount) external onlyOwner {
        // Require that the top up limit has not been initialized.
        require(!initializedTopUpLimit, "top up limit has already been initialized");
        // Require that the limit amount is within the acceptable range.
        require(MINIMUM_TOPUP_LIMIT <= _amount && _amount <= MAXIMUM_TOPUP_LIMIT, "top up amount is outside of the min/max range");
        // Modify spend limit based on the provided value.
        modifyTopUpLimit(_amount);
        // Flag operation as initialized.
        initializedTopUpLimit = true;
        // Emit the set limit event.
        emit SetTopUpLimit(msg.sender, _amount);
    }

    /// @dev Set a daily top up top up limit.
    /// @param _amount is the daily top up limit amount in wei.
    function submitTopUpLimit(uint _amount) external onlyOwner {
        // Require that the top up limit has been initialized.
        require(initializedTopUpLimit, "top up limit has not been initialized");
        // Require that the operation has not been submitted.
        require(!submittedTopUpLimit, "top up limit has already been submitted");
        // Require that the limit amount is within the acceptable range.
        require(MINIMUM_TOPUP_LIMIT <= _amount && _amount <= MAXIMUM_TOPUP_LIMIT, "top up amount is outside of the min/max range");
        // Assign the provided amount to pending daily limit change.
        pendingTopUpLimit = _amount;
        // Flag the operation as submitted.
        submittedTopUpLimit = true;
        // Emit the submission event.
        emit SubmittedTopUpLimitChange(_amount);
    }

    /// @dev Confirm pending set top up limit operation.
    function confirmTopUpLimit() external onlyController {
        // Require that the operation has been submitted.
        require(submittedTopUpLimit, "top up limit has not been submitted");
        // Assert that the pending top up limit amount is within the acceptable range.
        require(MINIMUM_TOPUP_LIMIT <= pendingTopUpLimit && pendingTopUpLimit <= MAXIMUM_TOPUP_LIMIT, "top up amount is outside the min/max range");
        // Modify top up limit based on the pending value.
        modifyTopUpLimit(pendingTopUpLimit);
        // Emit the set limit event.
        emit SetTopUpLimit(msg.sender, pendingTopUpLimit);
        // Reset pending daily limit.
        pendingTopUpLimit = 0;
        // Reset the submission flag.
        submittedTopUpLimit = false;
    }

    /// @dev Cancel pending set top up limit operation.
    function cancelTopUpLimit() external onlyController {
        // Reset pending daily limit.
        pendingTopUpLimit = 0;
        // Reset the submitted operation flag.
        submittedTopUpLimit = false;
        // Emit the cancellation event.
        emit CancelledTopUpLimitChange(msg.sender);
    }

    /// @dev Refill owner's gas balance.
    /// @param _amount is the amount of ether to transfer to the owner account in wei.
    function topUpGas(uint _amount) external isNotZero(_amount) {
        // Require that the sender is either the owner or a controller.
        require(isOwner() || isController(msg.sender), "sender is neither an owner nor a controller");
        // Account for the top up limit daily reset.
        updateTopUpAvailable();
        // Make sure the available top up amount is not zero.
        require(_topUpAvailable != 0, "available top up limit cannot be zero");
        // Limit top up amount to the available top up level.
        uint amount = _amount;
        if (amount > _topUpAvailable) {
            amount = _topUpAvailable;
        }
        // Reduce the top up amount from available balance and transfer corresponding
        // ether to the owner's account.
        _topUpAvailable -= amount;
        owner().transfer(amount);
        // Emit the gas top up event.
        emit ToppedUpGas(tx.origin, owner(), amount);
    }

    /// @dev Modify the top up limit and top up available based on the provided value.
    /// @dev _amount is the daily limit amount in wei.
    function modifyTopUpLimit(uint _amount) private {
        // Account for the top up limit daily reset.
        updateTopUpAvailable();
        // Set the daily limit to the provided amount.
        topUpLimit = _amount;
        // Lower the available limit if it's higher than the new daily limit.
        if (_topUpAvailable > topUpLimit) {
            _topUpAvailable = topUpLimit;
        }
    }

    /// @dev Update available top up limit based on the daily reset.
    function updateTopUpAvailable() private {
        if (now > _topUpLimitDay + 24 hours) {
            // Advance the current day by how many days have passed.
            uint extraDays = (now - _topUpLimitDay) / 24 hours;
            _topUpLimitDay += extraDays * 24 hours;
            // Set the available limit to the current top up limit.
            _topUpAvailable = topUpLimit;
        }
    }
}
