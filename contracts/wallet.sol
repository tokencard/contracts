pragma solidity ^0.4.24;

import "./internal/ownable.sol";
import "./internal/controllable.sol";


/// @title ERC20 is a subset of the ERC20 specification.
interface ERC20 {
    function transfer(address, uint) external returns (bool);
    function balanceOf(address) external view returns (uint);
}


/// @title ERC165 specifies a standard way of querying if a contract implements an interface.
interface ERC165 {
    function supportsInterface(bytes4) external view returns (bool);
}


/// @title Oracle converts ERC20 token amounts into equivalent ether amounts based on cryptocurrency exchange rates.
interface Oracle {
    function convert(address, uint) external view returns (uint);
}


/// @title Resolver returns the controller contract address.
interface Resolver {
    function getAddress() external returns (address);
}


/// @title Whitelist provides payee-whitelist functionality.
contract Whitelist is Controllable, Ownable {
    event WhitelistAdded(address _sender, address[] _addresses);
    event WhitelistAdditionSubmitted(address[] _addresses);
    event WhitelistAdditionCancelled(address _sender);

    event WhitelistRemoved(address _sender, address[] _addresses);
    event WhitelistRemovalSubmitted(address[] _addresses);
    event WhitelistRemovalCancelled(address _sender);

    mapping(address => bool) public isWhitelisted;
    address[] private _pendingWhitelistAddition;
    address[] private _pendingWhitelistRemoval;
    bool public submittedWhitelistAddition;
    bool public submittedWhitelistRemoval;
    bool public initializedWhitelist;

    /// @dev Check if the provided addresses contain the owner address.
    modifier hasNoOwner(address[] _addresses) {
        for (uint i = 0; i < _addresses.length; i++) {
            require(_addresses[i] != owner(), "provided whitelist contains the owner address");
        }
        _;
    }

    // @dev Getter for pending addition array.
    function pendingWhitelistAddition() external view returns(address[]) {
        return _pendingWhitelistAddition;
    }

    // @dev Getter for pending removal array.
    function pendingWhitelistRemoval() external view returns(address[]) {
        return _pendingWhitelistRemoval;
    }

    /// @dev Add initial addresses to the whitelist.
    /// @param _addresses are the Ethereum addresses to be whitelisted.
    function initializeWhitelist(address[] _addresses) external onlyOwner hasNoOwner(_addresses) {
        // Require that the whitelist has not been initialized.
        require(!initializedWhitelist, "whitelist has already been initialized");
        // Add each of the provided addresses to the whitelist.
        for (uint i = 0; i < _addresses.length; i++) {
            isWhitelisted[_addresses[i]] = true;
        }
        initializedWhitelist = true;
        // Emit the addition event.
        emit WhitelistAdded(msg.sender, _addresses);
    }

    /// @dev Add addresses to the whitelist.
    /// @param _addresses are the Ethereum addresses to be whitelisted.
    function submitWhitelistAddition(address[] _addresses) external onlyOwner hasNoOwner(_addresses) {
        // Require that either addition or removal operations have not been already submitted.
        require(!submittedWhitelistAddition && !submittedWhitelistRemoval, "whitelist operation has already been submitted");
        // Add the provided addresses to the pending addition list.
        _pendingWhitelistAddition = _addresses;
        // Flag the operation as submitted.
        submittedWhitelistAddition = true;
        // Flag operation as initialized if not initialized already.
        if (!initializedWhitelist) {
            initializedWhitelist = true;
        }
        // Emit the submission event.
        emit WhitelistAdditionSubmitted(_addresses);
    }

    /// @dev Confirm pending whitelist addition.
    function confirmWhitelistAddition() external onlyController {
        // Require that the pending whitelist is not empty and the operation has been submitted.
        require(submittedWhitelistAddition, "whitelist addition has not been submitted");
        require(_pendingWhitelistAddition.length > 0, "pending whitelist addition is empty");
        // Whitelist pending addresses.
        for (uint i = 0; i < _pendingWhitelistAddition.length; i++) {
            isWhitelisted[_pendingWhitelistAddition[i]] = true;
        }
        // Emit the addition event.
        emit WhitelistAdded(msg.sender, _pendingWhitelistAddition);
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
        emit WhitelistAdditionCancelled(msg.sender);
    }

    /// @dev Remove addresses from the whitelist.
    /// @param _addresses are the Ethereum addresses to be removed.
    function submitWhitelistRemoval(address[] _addresses) external onlyOwner {
        // Require that either addition or removal operations have not been already submitted.
        require(!submittedWhitelistRemoval && !submittedWhitelistAddition, "whitelist operation has already been submitted");
        // Add the provided addresses to the pending addition list.
        _pendingWhitelistRemoval = _addresses;
        // Flag the operation as submitted.
        submittedWhitelistRemoval = true;
        // Emit the submission event.
        emit WhitelistRemovalSubmitted(_addresses);
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
        emit WhitelistRemoved(msg.sender, _pendingWhitelistRemoval);
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
        emit WhitelistRemovalCancelled(msg.sender);
    }
}


/// @title SpendLimit provides daily spend limit functionality.
contract SpendLimit is Controllable, Ownable {
    event SpendLimitSet(address _sender, uint _amount);
    event SpendLimitSubmitted(uint _amount);
    event SpendLimitCancelled(address _sender);

    uint public spendLimit;
    uint internal _spendLimitDay;
    uint internal _spendAvailable;

    uint public pendingSpendLimit;
    bool public submittedSpendLimit;
    bool public initializedSpendLimit;

    /// @dev Constructor initializes the daily spend limit.
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

    /// @dev Initialize a daily transfer limit for non-whitelisted addresses.
    /// @param _amount is the daily limit amount in wei.
    function initializeSpendLimit(uint _amount) external onlyOwner {
        // Require that the spend limit has not been initialized.
        require(!initializedSpendLimit, "spend limit has already been initialized");
        // Modify spend limit based on the provided value.
        modifySpendLimit(_amount);
        // Flag the operation as initialized.
        initializedSpendLimit = true;
        // Emit the set limit event.
        emit SpendLimitSet(msg.sender, _amount);
    }

    /// @dev Set a daily transfer limit for non-whitelisted addresses.
    /// @param _amount is the daily limit amount in wei.
    function submitSpendLimit(uint _amount) external onlyOwner {
        // Require that the operation has been submitted.
        require(!submittedSpendLimit, "spend limit has already been submitted");
        // Assign the provided amount to pending daily limit change.
        pendingSpendLimit = _amount;
        // Flag the operation as submitted.
        submittedSpendLimit = true;
        // Flag operation as initialized if not initialized already.
        if (!initializedSpendLimit) {
            initializedSpendLimit = true;
        }
        // Emit the submission event.
        emit SpendLimitSubmitted(_amount);
    }

    /// @dev Confirm pending set daily limit operation.
    function confirmSpendLimit() external onlyController {
        // Require that the operation has been submitted.
        require(submittedSpendLimit, "spend limit has not been submitted");
        // Modify spend limit based on the pending value.
        modifySpendLimit(pendingSpendLimit);
        // Emit the set limit event.
        emit SpendLimitSet(msg.sender, pendingSpendLimit);
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
        emit SpendLimitCancelled(msg.sender);
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


/// @title Asset store with extra security features.
contract Vault is Whitelist, SpendLimit, ERC165 {
    event DepositReceived(address _from, uint _amount);
    event Transferred(address _to, address _asset, uint _amount);

    //// @dev Supported ERC165 interface ID.
    bytes4 private constant _ERC165_INTERFACE_ID = 0x01ffc9a7; // solium-disable-line uppercase

    /// @dev Resolver points to the oracle address resolver.
    Resolver private _OR; // solium-disable-line mixedcase

    /// @dev Constructor initializes the vault with an owner address and spend limit. It also sets up the oracle and controller contracts.
    /// @param _owner is the owner account of the wallet contract.
    /// @param _transferable indicates whether the contract ownership can be transferred.
    /// @param _resolver is the oracle resolver contract address.
    /// @param _controller is the controller contract address.
    constructor(address _owner, bool _transferable, address _resolver, address _controller, uint _spendLimit) SpendLimit(_spendLimit) Ownable(_owner, _transferable) Controllable(_controller) public {
        _OR = Resolver(_resolver);
    }

    /// @dev Checks if the value is not zero.
    modifier isNotZero(uint _value) {
        require(_value != 0, "provided value cannot be zero");
        _;
    }

    /// @dev Ether can be deposited from any source, so this contract must be payable by anyone.
    function() public payable {
        emit DepositReceived(msg.sender, msg.value);
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
    function transfer(address _to, address _asset, uint _amount) external onlyOwner isNotZero(_amount) {
        // If address is not whitelisted, take daily limit into account.
        if (!isWhitelisted[_to]) {
            // Update the available spend limit.
            updateSpendAvailable();
            // Convert token amount to ether value.
            uint etherValue;
            if (_asset != 0x0) {
                etherValue = Oracle(_OR.getAddress()).convert(_asset, _amount);
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


/// @title Asset wallet with extra security features and gas topup management.
contract Wallet is Vault {
    event TopupLimitSet(address _sender, uint _amount);
    event TopupLimitSubmitted(uint _amount);
    event TopupLimitCancelled(address _sender);

    event TopupGas(address _sender, address _owner, uint _amount);

    uint constant private MINIMUM_TOPUP_LIMIT = 1 finney;
    uint constant private MAXIMUM_TOPUP_LIMIT = 500 finney;

    uint public topupLimit;
    uint private _topupLimitDay;
    uint private _topupAvailable;

    uint public pendingTopupLimit;
    bool public submittedTopupLimit;
    bool public initializedTopupLimit;

    /// @dev Constructor initializes the wallet top up limit and the vault contract.
    /// @param _owner is the owner account of the wallet contract.
    /// @param _transferable indicates whether the contract ownership can be transferred.
    /// @param _resolver is the oracle resolver contract address.
    /// @param _controller is the controller contract address.
    constructor(address _owner, bool _transferable, address _resolver, address _controller, uint _spendLimit) Vault(_owner, _transferable, _resolver, _controller, _spendLimit) public {
        _topupLimitDay = now;
        topupLimit = MAXIMUM_TOPUP_LIMIT;
        _topupAvailable = topupLimit;
    }

    /// @dev Returns the available daily gas top up balance - accounts for daily limit reset.
    /// @return amount of gas in wei.
    function topupAvailable() external view returns (uint) {
        if (now > _topupLimitDay + 24 hours) {
            return topupLimit;
        } else {
            return _topupAvailable;
        }
    }

    /// @dev Initialize a daily gas top up limit.
    /// @param _amount is the gas top up amount in wei.
    function initializeTopupLimit(uint _amount) external onlyOwner {
        // Require that the topup limit has not been initialized.
        require(!initializedTopupLimit, "topup limit has already been initialized");
        // Require that the limit amount is within the acceptable range.
        require(MINIMUM_TOPUP_LIMIT <= _amount && _amount <= MAXIMUM_TOPUP_LIMIT, "topup amount is outside of the min/max range");
        // Modify spend limit based on the provided value.
        modifyTopupLimit(_amount);
        // Flag operation as initialized.
        initializedTopupLimit = true;
        // Emit the set limit event.
        emit TopupLimitSet(msg.sender, _amount);
    }

    /// @dev Set a daily topup top up limit.
    /// @param _amount is the daily topup limit amount in wei.
    function submitTopupLimit(uint _amount) external onlyOwner {
        // Require that the operation has not been submitted.
        require(!submittedTopupLimit, "topup limit has already been submitted");
        // Require that the limit amount is within the acceptable range.
        require(MINIMUM_TOPUP_LIMIT <= _amount && _amount <= MAXIMUM_TOPUP_LIMIT, "topup amount is outside of the min/max range");
        // Assign the provided amount to pending daily limit change.
        pendingTopupLimit = _amount;
        // Flag the operation as submitted.
        submittedTopupLimit = true;
        // Flag operation as initialized if not initialized already.
        if (!initializedTopupLimit) {
            initializedTopupLimit = true;
        }
        // Emit the submission event.
        emit TopupLimitSubmitted(_amount);
    }

    /// @dev Confirm pending set top up limit operation.
    function confirmTopupLimit() external onlyController {
        // Require that the operation has been submitted.
        require(submittedTopupLimit, "topup limit has not been submitted");
        // Assert that the pending topup limit amount is within the acceptable range.
        assert(MINIMUM_TOPUP_LIMIT <= pendingTopupLimit && pendingTopupLimit <= MAXIMUM_TOPUP_LIMIT);
        // Modify topup limit based on the pending value.
        modifyTopupLimit(pendingTopupLimit);
        // Emit the set limit event.
        emit TopupLimitSet(msg.sender, pendingTopupLimit);
        // Reset pending daily limit.
        pendingTopupLimit = 0;
        // Reset the submission flag.
        submittedTopupLimit = false;
    }

    /// @dev Cancel pending set top up limit operation.
    function cancelTopupLimit() external onlyController {
        // Reset pending daily limit.
        pendingTopupLimit = 0;
        // Reset the submitted operation flag.
        submittedTopupLimit = false;
        // Emit the cancellation event.
        emit TopupLimitCancelled(msg.sender);
    }

    /// @dev Refill owner's gas balance.
    /// @param _amount is the amount of ether to transfer to the owner account in wei.
    function topupGas(uint _amount) external isNotZero(_amount) {
        // Require that the sender is either the owner or a controller.
        require(isOwner() || isController(msg.sender), "sender is neither an owner nor a controller");
        // Account for the topup limit daily reset.
        updateTopupAvailable();
        // Make sure the available topup amount is not zero.
        require(_topupAvailable != 0, "available topup limit cannot be zero");
        // Limit topup amount to the available topup level.
        uint amount = _amount;
        if (amount > _topupAvailable) {
            amount = _topupAvailable;
        }
        // Reduce the top up amount from available balance and transfer corresponding
        // ether to the owner's account.
        _topupAvailable -= amount;
        owner().transfer(amount);
        // Emit the gas topup event.
        emit TopupGas(tx.origin, owner(), amount);
    }

    /// @dev Update available topup limit based on the daily reset.
    function updateTopupAvailable() private {
        if (now > _topupLimitDay + 24 hours) {
            // Advance the current day by how many days have passed.
            uint extraDays = (now - _topupLimitDay) / 24 hours;
            _topupLimitDay += extraDays * 24 hours;
            // Set the available limit to the current topup limit.
            _topupAvailable = topupLimit;
        }
    }

    /// @dev Modify the topup limit and topup available based on the provided value.
    /// @dev _amount is the daily limit amount in wei.
    function modifyTopupLimit(uint _amount) private {
        // Account for the topup limit daily reset.
        updateTopupAvailable();
        // Set the daily limit to the provided amount.
        topupLimit = _amount;
        // Lower the available limit if it's higher than the new daily limit.
        if (_topupAvailable > topupLimit) {
            _topupAvailable = topupLimit;
        }
    }
}
