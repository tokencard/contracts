pragma solidity ^0.4.24;


/// @title The Token interface is a subset of the ERC20 specification.
interface Token {
    function transfer(address, uint) external returns (bool);
    function balanceOf(address) external view returns (uint);
}

/// @title The Oracle interface provides exchange rates for ERC20 tokens in wei.
interface Oracle {
    function rate(address) external view returns (uint);
}

/// @title Control handles wallet access control.
contract Control {
    event AddController(address _account);
    event RemoveController(address _account);

    mapping (address => bool) public isController;
    address public owner;

    /// @dev Executable only by the owner.
    modifier onlyOwner {
        require(msg.sender == owner);
        _;
    }

    /// @dev Executable only by the controller.
    modifier onlyController {
        require(isController[msg.sender]);
        _;
    }

    /// @dev Executable by either owner or controller.
    modifier eitherOwnerOrController() {
        require(isController[msg.sender] || msg.sender == owner);
        _;
    }

    /// @dev Add a new controller to the list of controllers.
    function addController(address _account) public onlyController {
        isController[_account] = true;
        emit AddController(_account);
    }

    /// @dev Remove a controller from the list of controllers.
    function removeController(address _account) public onlyController {
        isController[_account] = false;
        emit RemoveController(_account);
    }
}

/// @title Whitelist provides payee-whitelist functionality.
contract Whitelist is Control {

    event WhitelistAddition(address[] _addresses);
    event WhitelistRemoval(address[] _addresses);

    mapping(address => bool) public isWhitelisted;
    address[] private _pendingAddition;
    address[] private _pendingRemoval;
    bool private _submittedAddition;
    bool private _submittedRemoval;
    bool private _initializedWhitelist;

    // @dev Getter for pending addition array.
    function pendingAddition() public view returns(address[]) {
        return _pendingAddition;
    }

    // @dev Getter for pending removal array.
    function pendingRemoval() public view returns(address[]) {
        return _pendingRemoval;
    }

    /// @dev Add initial addresses to the whitelist.
    /// @param _addresses are the Ethereum addresses to be whitelisted.
    function initializeWhitelist(address[] _addresses) public onlyOwner {
        require(!_initializedWhitelist);
        // Add each of the provided addresses to the whitelist.
        for (uint i = 0; i < _addresses.length; i++) {
            isWhitelisted[_addresses[i]] = true;
        }
        _initializedWhitelist = true;
        emit WhitelistAddition(_addresses);
    }

    /// @dev Add addresses to the whitelist.
    /// @param _addresses are the Ethereum addresses to be whitelisted.
    function addToWhitelist(address[] _addresses) public onlyOwner {
        // Check if this operation has been already submitted.
        require(!_submittedAddition);
        // Limit the amount of addresses that can be whitelisted at one time.
        require(_addresses.length <= 20);
        // Add each of the provided addresses to the pending addition list.
        for (uint i = 0; i < _addresses.length; i++) {
            _pendingAddition.push(_addresses[i]);
        }
        // Flag the operation as submitted.
        _submittedAddition = true;
    }

    /// @dev Confirm pending whitelist addition.
    function addToWhitelistConfirm() public onlyController {
        require(_pendingAddition.length > 0 && _submittedAddition);
        // Whitelist pending addresses.
        for (uint i = 0; i < _pendingAddition.length; i++) {
            isWhitelisted[_pendingAddition[i]] = true;
        }
        // Reset the submission flag.
        _submittedAddition = false;
        emit WhitelistAddition(_pendingAddition);
    }

    /// @dev Cancel pending whitelist addition.
    function addToWhitelistCancel() public onlyController {
        // Reset pending addresses.
        delete _pendingAddition;
        // Reset the submitted operation flag.
        _submittedAddition = false;
    }

    /// @dev Remove addresses from the whitelist.
    /// @param _addresses are the Ethereum addresses to be removed.
    function removeFromWhitelist(address[] _addresses) public onlyOwner {
        // Check if this operation has been already submitted.
        require(!_submittedRemoval);
        // Limit the amount of addresses that can be removed at one time.
        require(_addresses.length <= 20);
        // Add each of the addresses to the pending removal list.
        for (uint i = 0; i < _addresses.length; i++) {
            _pendingRemoval.push(_addresses[i]);
        }
        // Flag the operation as submitted.
        _submittedRemoval = true;
    }

    /// @dev Confirm pending removal of whitelisted addresses.
    function removeFromWhitelistConfirm() public onlyController {
        require(_pendingRemoval.length > 0 && _submittedRemoval);
        // Remove pending addresses.
        for (uint i = 0; i < _pendingRemoval.length; i++) {
            isWhitelisted[_pendingRemoval[i]] = false;
        }
        // Reset the submission flag.
        _submittedRemoval = false;
        emit WhitelistRemoval(_pendingRemoval);
    }

    /// @dev Cancel pending removal of whitelisted addresses.
    function removeFromWhitelistCancel() public onlyController {
        // Reset pending addresses.
        delete _pendingRemoval;
        // Reset the submitted operation flag.
        _submittedRemoval = false;
    }
}

/// @title DailyLimit provides daily spend limit functionality.
contract DailyLimit is Control {

    event SetDailyLimit(uint _amount);

    uint public currentDay;
    uint public dailyLimit;
    uint internal _dailyAvailable;

    uint public pendingDailyLimit;
    bool private _submittedDailyLimit;
    bool private _initializedDailyLimit;

    /// @dev Returns the available daily balance - accounts for daily limit reset.
    /// @return amount of ether in wei.
    function dailyAvailable() public view returns (uint) {
        if (now > currentDay + 24 hours) {
            return dailyLimit;
        } else {
            return _dailyAvailable;
        }
    }
    /// @dev Initialize a daily transfer limit for non-whitelisted addresses.
    /// @param _amount is the daily limit amount in wei.
    function initializeDailyLimit(uint _amount) public onlyOwner {
        require(!_initializedDailyLimit);
        // Set the daily limit to the provided amount.
        dailyLimit = _amount;
        _initializedDailyLimit = true;
        emit SetDailyLimit(_amount);
    }

    /// @dev Set a daily transfer limit for non-whitelisted addresses.
    /// @param _amount is the daily limit amount in wei.
    function setDailyLimit(uint _amount) public onlyOwner {
        // Check if this operation has been already submitted.
        require(!_submittedDailyLimit);
        // Assign the provided amount to pending daily limit change.
        pendingDailyLimit = _amount;
        // Flag the operation as submitted.
        _submittedDailyLimit = true;
    }

    /// @dev Confirm pending set daily limit operation.
    function setDailyLimitConfirm() public onlyController {
        require(_submittedDailyLimit);
        // Set the daily limit to the pending amount.
        dailyLimit = pendingDailyLimit;
        // Reset the submission flag.
        _submittedDailyLimit = false;
        emit SetDailyLimit(pendingDailyLimit);
    }

    /// @dev Cancel pending set daily limit operation.
    function setDailyLimitCancel() public onlyController {
        // Reset pending daily limit change.
        pendingDailyLimit = 0;
        // Reset the submitted operation flag.
        _submittedDailyLimit = false;
    }
}

/// @title Asset wallet with extra security features.
/// @author TokenCard
contract Vault is Whitelist, DailyLimit {
    // Events
    event Deposit(address _from, uint _amount);
    event Transfer(address _to, address _asset, uint _amount);

    // Storage
    address public oracle;

    /// @dev Construct a wallet with an owner and a controller.
    constructor(address _owner, address _oracle, address[] _controllers) public {
        currentDay = now;
        dailyLimit = 1 ether;
        _dailyAvailable = dailyLimit;
        owner = _owner;
        oracle = _oracle;
        for (uint i = 0; i < _controllers.length; i++) {
            isController[_controllers[i]] = true;
        }
    }

    /// @dev Checks if the value is not zero.
    modifier notZero(uint _value) {
        require(_value != 0);
        _;
    }

    /// @dev Ether can be deposited from any source, so this contract must be payable by anyone.
    function() public payable {
        if (msg.value > 0) {
            emit Deposit(msg.sender, msg.value);
        }
    }

    /// @dev Returns the amount of an asset owned by the contract.
    /// @param _asset address of an ERC20 token or 0x0 for ether.
    /// @return balance associated with the wallet address in wei.
    function balance(address _asset) public view returns (uint) {
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
    function transfer(address _to, address _asset, uint _amount) public onlyOwner notZero(_amount) {
        // If address is not whitelisted, take daily limit into account.
        if (!isWhitelisted[_to]) {
            // Convert token amount to ether value.
            uint etherValue;
            if (_asset != 0x0) {
                etherValue = Oracle(oracle).rate(_asset) * _amount;
                assert(etherValue != 0);
            } else {
                etherValue = _amount;
            }
            // Require that the value is under remaining limit.
            if (now > currentDay + 24 hours) {
                uint extraDays = (now - currentDay) / 24 hours;
                currentDay += extraDays * 24 hours;
                _dailyAvailable = dailyLimit;
            }
            require(etherValue <= _dailyAvailable);
            // Update the available limit.
            _dailyAvailable -= etherValue;
        }
        // Transfer token or ether based on the provided address.
        if (_asset != 0x0) {
            require(Token(_asset).transfer(_to, _amount));
        } else {
            _to.transfer(_amount);
        }
        emit Transfer(_to, _asset, _amount);
    }
}

contract Wallet is Vault {

    event SetGasLimit(uint _amount);
    event TopUpGas(address _sender, address _owner, uint _amount);

    uint constant private MINIMUM_GAS_LIMIT = 1 finney;
    uint constant private MAXIMUM_GAS_LIMIT = 100 finney;

    uint public gasLimit;
    uint private _gasAvailable;

    uint public pendingGasLimit;
    bool private _submittedGasLimit;
    bool private _initializedGasLimit;

    constructor(address _owner, address _oracle, address[] _controllers) Vault(_owner, _oracle, _controllers) public {
        gasLimit = MAXIMUM_GAS_LIMIT;
        _gasAvailable = gasLimit;
    }

    /// @dev Returns the available daily gas top up balance - accounts for daily limit reset.
    /// @return amount of gas in wei.
    function gasAvailable() public view returns (uint) {
        if (now > currentDay + 24 hours) {
            return gasLimit;
        } else {
            return _gasAvailable;
        }
    }

    /// @dev Initialize a daily gas top up limit.
    /// @param _amount is the gas top up amount in wei.
    function initializeGasLimit(uint _amount) public onlyOwner {
        require(!_initializedGasLimit);
        // Require that the limit amount is within the acceptable range.
        require(MINIMUM_GAS_LIMIT <= _amount && _amount <= MAXIMUM_GAS_LIMIT);
        // Set the gas limit to the provided amount.
        gasLimit = _amount;
        _initializedGasLimit = true;
        emit SetGasLimit(_amount);
    }

    /// @dev Set a daily gas top up limit.
    /// @param _amount is the daily gas limit amount in wei.
    function setGasLimit(uint _amount) public onlyOwner {
        // Check if this operation has been already submitted.
        require(!_submittedGasLimit);
        // Require that the limit amount is within the acceptable range.
        require(MINIMUM_GAS_LIMIT <= _amount && _amount <= MAXIMUM_GAS_LIMIT);
        // Assign the provided amount to pending daily limit change.
        pendingDailyLimit = _amount;
        // Flag the operation as submitted.
        _submittedGasLimit = true;
    }

    /// @dev Confirm pending set gas limit operation.
    function setGasLimitConfirm() public onlyController {
        // That a set gas limit operation has been submitted.
        require(_submittedGasLimit);
        // Assert that the pending gas limit amount is within the acceptable range.
        assert(MINIMUM_GAS_LIMIT <= pendingDailyLimit && pendingDailyLimit <= MAXIMUM_GAS_LIMIT);
        // Set the daily limit to the pending amount.
        gasLimit = pendingDailyLimit;
        // Reset the submission flag.
        _submittedGasLimit = false;
        emit SetGasLimit(pendingDailyLimit);
    }

    /// @dev Cancel pending set gas limit operation.
    function setGasLimitCancel() public onlyController {
        // Reset pending daily limit change.
        pendingDailyLimit = 0;
        // Reset the submitted operation flag.
        _submittedGasLimit = false;
    }

    /// @dev Refill owner's gas balance.
    /// @param _amount the amount of ether to transfer to the owner account in wei.
    function topUpGas(uint _amount) public eitherOwnerOrController notZero(_amount) {
        // Make sure the available gas is not zero.
        require(_gasAvailable > 0);
        // Account for the gas limit daily reset.
        if (now > currentDay + 24 hours) {
            uint extraDays = (now - currentDay) / 24 hours;
            currentDay += extraDays * 24 hours;
            _gasAvailable = gasLimit;
        }
        // If amount is above available balance, use the entire balance.
        if (_amount > _gasAvailable) {
            _amount = _gasAvailable;
        }
        // Deduce the top up amount from available balance and transfer corresponding
        // ether to the owner's account.
        _gasAvailable -= _amount;
        owner.transfer(_amount);
        emit TopUpGas(tx.origin, owner, _amount);
    }
}
