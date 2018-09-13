pragma solidity 0.4.24;


/// @title The Token interface is a subset of the ERC20 specification.
interface Token {
    function transfer(address, uint) external returns (bool);
    function balanceOf(address) external view returns (uint);
}

/// @title The Oracle interface provides exchange rates for ERC20 tokens in wei.
interface Oracle {
    function convert(address, uint) external view returns (uint);
}

/// @title Control handles wallet access control.
contract Control {
    event AddController(address _sender, address _account);
    event RemoveController(address _sender, address _account);

    mapping (address => bool) public isController;
    uint public controllerCount;
    address public owner;

    /// @dev Executable only by the owner.
    modifier onlyOwner() {
        require(msg.sender == owner);
        _;
    }

    /// @dev Executable only by the controller.
    modifier onlyController() {
        require(isController[msg.sender]);
        _;
    }

    /// @dev Executable by either owner or controller.
    modifier eitherOwnerOrController() {
        require(isController[msg.sender] || msg.sender == owner);
        _;
    }

    /// @dev Add a new controller to the list of controllers.
    function addController(address _account) external onlyController {
        addControllerInternal(_account);
    }

    /// @dev Re-usable internal function that adds a new controller.
    function addControllerInternal(address _account) internal {
        require(!isController[_account]);
        isController[_account] = true;
        controllerCount++;
        emit AddController(msg.sender, _account);
    }

    /// @dev Remove a controller from the list of controllers.
    function removeController(address _account) external onlyController {
        removeControllerInternal(_account);
    }

    /// @dev Re-usable internal function that removes an existing controller.
    function removeControllerInternal(address _account) internal {
        require(isController[_account] && controllerCount > 1);
        isController[_account] = false;
        controllerCount--;
        emit RemoveController(msg.sender, _account);
    }
}

/// @title Whitelist provides payee-whitelist functionality.
contract Whitelist is Control {
    event WhitelistAddition(address _sender, address[] _addresses);
    event SubmitWhitelistAddition(address[] _addresses);
    event CancelWhitelistAddition(address _sender);

    event WhitelistRemoval(address _sender, address[] _addresses);
    event SubmitWhitelistRemoval(address[] _addresses);
    event CancelWhitelistRemoval(address _sender);

    mapping(address => bool) public isWhitelisted;
    address[] private _pendingWhitelistAddition;
    address[] private _pendingWhitelistRemoval;
    bool public submittedWhitelistAddition;
    bool public submittedWhitelistRemoval;
    bool public initializedWhitelist;

    /// @dev Check if the provided addresses array has a valid length.
    modifier validLength(address[] _addresses) {
        require(_addresses.length >= 1 && _addresses.length <= 20);
        _;
    }

    /// @dev Check if the provided addresses contain the owner address.
    modifier hasNoOwner(address[] _addresses) {
        for (uint i = 0; i < _addresses.length; i++) {
            require(_addresses[i] != owner);
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
    function initializeWhitelist(address[] _addresses) external onlyOwner validLength(_addresses) hasNoOwner(_addresses) {
        // Require that the whitelist has not been initialized.
        require(!initializedWhitelist);
        // Add each of the provided addresses to the whitelist.
        for (uint i = 0; i < _addresses.length; i++) {
            isWhitelisted[_addresses[i]] = true;
        }
        initializedWhitelist = true;
        // Emit the addition event.
        emit WhitelistAddition(msg.sender, _addresses);
    }

    /// @dev Add addresses to the whitelist.
    /// @param _addresses are the Ethereum addresses to be whitelisted.
    function submitWhitelistAddition(address[] _addresses) external onlyOwner validLength(_addresses) hasNoOwner(_addresses) {
        // Require that either addition or removal operations have not been already submitted.
        require(!submittedWhitelistAddition && !submittedWhitelistRemoval);
        // Add the provided addresses to the pending addition list.
        _pendingWhitelistAddition = _addresses;
        // Flag the operation as submitted.
        submittedWhitelistAddition = true;
        // Flag operation as initialized if not initialized already.
        if (!initializedWhitelist) {
            initializedWhitelist = true;
        }
        // Emit the submission event.
        emit SubmitWhitelistAddition(_addresses);
    }

    /// @dev Confirm pending whitelist addition.
    function confirmWhitelistAddition() external onlyController {
        // Require that the pending whitelist is not empty and the operation has been submitted.
        require(_pendingWhitelistAddition.length > 0 && submittedWhitelistAddition);
        // Whitelist pending addresses.
        for (uint i = 0; i < _pendingWhitelistAddition.length; i++) {
            isWhitelisted[_pendingWhitelistAddition[i]] = true;
        }
        // Emit the addition event.
        emit WhitelistAddition(msg.sender, _pendingWhitelistAddition);
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
        emit CancelWhitelistAddition(msg.sender);
    }

    /// @dev Remove addresses from the whitelist.
    /// @param _addresses are the Ethereum addresses to be removed.
    function submitWhitelistRemoval(address[] _addresses) external onlyOwner validLength(_addresses) {
        // Require that either addition or removal operations have not been already submitted.
        require(!submittedWhitelistRemoval && !submittedWhitelistAddition);
        // Add the provided addresses to the pending addition list.
        _pendingWhitelistRemoval = _addresses;
        // Flag the operation as submitted.
        submittedWhitelistRemoval = true;
        // Emit the submission event.
        emit SubmitWhitelistRemoval(_addresses);
    }

    /// @dev Confirm pending removal of whitelisted addresses.
    function confirmWhitelistRemoval() external onlyController {
        // Require that the pending whitelist is not empty and the operation has been submitted.
        require(_pendingWhitelistRemoval.length > 0 && submittedWhitelistRemoval);
        // Remove pending addresses.
        for (uint i = 0; i < _pendingWhitelistRemoval.length; i++) {
            isWhitelisted[_pendingWhitelistRemoval[i]] = false;
        }
        // Emit the removal event.
        emit WhitelistRemoval(msg.sender, _pendingWhitelistRemoval);
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
        emit CancelWhitelistRemoval(msg.sender);
    }
}

/// @title SpendLimit provides daily spend limit functionality.
contract SpendLimit is Control {
    event SetSpendLimit(address _sender, uint _amount);
    event SubmitSpendLimit(uint _amount);
    event CancelSpendLimit(address _sender);

    uint public spendLimit;
    uint internal _spendLimitDay;
    uint internal _spendAvailable;

    uint public pendingSpendLimit;
    bool public submittedSpendLimit;
    bool public initializedSpendLimit;

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
        require(!initializedSpendLimit);
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
        // Require that the operation has been submitted.
        require(!submittedSpendLimit);
        // Assign the provided amount to pending daily limit change.
        pendingSpendLimit = _amount;
        // Flag the operation as submitted.
        submittedSpendLimit = true;
        // Flag operation as initialized if not initialized already.
        if (!initializedSpendLimit) {
            initializedSpendLimit = true;
        }
        // Emit the submission event.
        emit SubmitSpendLimit(_amount);
    }

    /// @dev Confirm pending set daily limit operation.
    function confirmSpendLimit() external onlyController {
        // Require that the operation has been submitted.
        require(submittedSpendLimit);
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
        emit CancelSpendLimit(msg.sender);
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
}

/// @title Asset store with extra security features.
contract Vault is Whitelist, SpendLimit {
    event Deposit(address _from, uint _amount);
    event Transfer(address _to, address _asset, uint _amount);

    address public oracle;

    /// @dev Construct a wallet with an owner and a controller.
    constructor(address _owner, address _oracle, address[] _controllers) public {
        spendLimit = 100 ether;
        _spendLimitDay = now;
        _spendAvailable = spendLimit;
        owner = _owner;
        oracle = _oracle;
        for (uint i = 0; i < _controllers.length; i++) {
            addControllerInternal(_controllers[i]);
        }
    }

    /// @dev Check if the value is not zero.
    modifier notZero(uint _value) {
        require(_value != 0);
        _;
    }

    /// @dev Ether can be deposited from any source, so this contract must be payable by anyone.
    function() public payable {
        if (msg.value > 0) {
            // Emit the deposit event.
            emit Deposit(msg.sender, msg.value);
        }
    }

    /// @dev Returns the amount of an asset owned by the contract.
    /// @param _asset address of an ERC20 token or 0x0 for ether.
    /// @return balance associated with the wallet address in wei.
    function balance(address _asset) external view returns (uint) {
        if (_asset != 0x0) {
            return Token(_asset).balanceOf(this);
        } else {
            return address(this).balance;
        }
    }

    /// @dev Transfer asset to an address.
    /// @param _to recipient address.
    /// @param _asset address of an ERC20 token or 0x0 for ether.
    /// @param _amount is the amount of tokens to be transferred in base units.
    function transfer(address _to, address _asset, uint _amount) external onlyOwner notZero(_amount) {
        // If address is not whitelisted, take daily limit into account.
        if (!isWhitelisted[_to]) {
            // Update the available spend limit.
            updateSpendAvailable();
            // Convert token amount to ether value.
            uint etherValue;
            if (_asset != 0x0) {
                etherValue = Oracle(oracle).convert(_asset, _amount);
            } else {
                etherValue = _amount;
            }
            // Require that the value is under remaining limit.
            require(etherValue <= _spendAvailable);
            // Update the available limit.
            _spendAvailable -= etherValue;
        }
        // Transfer token or ether based on the provided address.
        if (_asset != 0x0) {
            require(Token(_asset).transfer(_to, _amount));
        } else {
            _to.transfer(_amount);
        }
        // Emit the transfer event.
        emit Transfer(_to, _asset, _amount);
    }
}

/// @title Asset wallet with extra security features and gas topup management.
contract Wallet is Vault {
    event SetTopupLimit(address _sender, uint _amount);
    event SubmitTopupLimit(uint _amount);
    event CancelTopupLimit(address _sender);

    event TopupGas(address _sender, address _owner, uint _amount);

    uint constant private MINIMUM_TOPUP_LIMIT = 1 finney;
    uint constant private MAXIMUM_TOPUP_LIMIT = 500 finney;

    uint public topupLimit;
    uint private _topupLimitDay;
    uint private _topupAvailable;

    uint public pendingTopupLimit;
    bool public submittedTopupLimit;
    bool public initializedTopupLimit;

    constructor(address _owner, address _oracle, address[] _controllers) Vault(_owner, _oracle, _controllers) public {
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
        require(!initializedTopupLimit);
        // Require that the limit amount is within the acceptable range.
        require(MINIMUM_TOPUP_LIMIT <= _amount && _amount <= MAXIMUM_TOPUP_LIMIT);
        // Modify spend limit based on the provided value.
        modifyTopupLimit(_amount);
        // Flag operation as initialized.
        initializedTopupLimit = true;
        // Emit the set limit event.
        emit SetTopupLimit(msg.sender, _amount);
    }

    /// @dev Set a daily topup top up limit.
    /// @param _amount is the daily topup limit amount in wei.
    function submitTopupLimit(uint _amount) external onlyOwner {
        // Require that the operation has not been submitted.
        require(!submittedTopupLimit);
        // Require that the limit amount is within the acceptable range.
        require(MINIMUM_TOPUP_LIMIT <= _amount && _amount <= MAXIMUM_TOPUP_LIMIT);
        // Assign the provided amount to pending daily limit change.
        pendingTopupLimit = _amount;
        // Flag the operation as submitted.
        submittedTopupLimit = true;
        // Flag operation as initialized if not initialized already.
        if (!initializedTopupLimit) {
            initializedTopupLimit = true;
        }
        // Emit the submission event.
        emit SubmitTopupLimit(_amount);
    }

    /// @dev Confirm pending set top up limit operation.
    function confirmTopupLimit() external onlyController {
        // Require that the operation has been submitted.
        require(submittedTopupLimit);
        // Assert that the pending topup limit amount is within the acceptable range.
        assert(MINIMUM_TOPUP_LIMIT <= pendingTopupLimit && pendingTopupLimit <= MAXIMUM_TOPUP_LIMIT);
        // Modify topup limit based on the pending value.
        modifyTopupLimit(pendingTopupLimit);
        // Emit the set limit event.
        emit SetTopupLimit(msg.sender, pendingTopupLimit);
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
        emit CancelTopupLimit(msg.sender);
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

    /// @dev Refill owner's gas balance.
    /// @param _amount the amount of ether to transfer to the owner account in wei.
    function topupGas(uint _amount) external eitherOwnerOrController notZero(_amount) {
        // Account for the topup limit daily reset.
        updateTopupAvailable();
        // Make sure the available topup is not zero.
        require(_topupAvailable > 0);
        // If amount is above available balance, use the entire balance.
        uint amount = _amount;
        if (amount > _topupAvailable) {
            amount = _topupAvailable;
        }
        // Reduce the top up amount from available balance and transfer corresponding
        // ether to the owner's account.
        _topupAvailable -= amount;
        owner.transfer(amount);
        // Emit the gas topup event.
        emit TopupGas(tx.origin, owner, amount);
    }
}
