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
            if (_addresses[i] != owner) {
                _pendingAddition.push(_addresses[i]);
            }
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

/// @title SpendLimit provides daily spend limit functionality.
contract SpendLimit is Control {

    event SetSpendLimit(uint _amount);

    uint public currentDay;
    uint public spendLimit;
    uint internal _spendAvailable;

    uint public pendingSpendLimit;
    bool private _submittedSpendLimit;
    bool private _initializedSpendLimit;

    /// @dev Returns the available daily balance - accounts for daily limit reset.
    /// @return amount of ether in wei.
    function spendAvailable() public view returns (uint) {
        if (now > currentDay + 24 hours) {
            return spendLimit;
        } else {
            return _spendAvailable;
        }
    }
    /// @dev Initialize a daily transfer limit for non-whitelisted addresses.
    /// @param _amount is the daily limit amount in wei.
    function initializeSpendLimit(uint _amount) public onlyOwner {
        require(!_initializedSpendLimit);
        // Set the daily limit to the provided amount.
        spendLimit = _amount;
        _initializedSpendLimit = true;
        emit SetSpendLimit(_amount);
    }

    /// @dev Set a daily transfer limit for non-whitelisted addresses.
    /// @param _amount is the daily limit amount in wei.
    function setSpendLimit(uint _amount) public onlyOwner {
        // Check if this operation has been already submitted.
        require(!_submittedSpendLimit);
        // Assign the provided amount to pending daily limit change.
        pendingSpendLimit = _amount;
        // Flag the operation as submitted.
        _submittedSpendLimit = true;
    }

    /// @dev Confirm pending set daily limit operation.
    function setSpendLimitConfirm() public onlyController {
        require(_submittedSpendLimit);
        // Set the daily limit to the pending amount.
        spendLimit = pendingSpendLimit;
        // Reset the submission flag.
        _submittedSpendLimit = false;
        emit SetSpendLimit(pendingSpendLimit);
    }

    /// @dev Cancel pending set daily limit operation.
    function setSpendLimitCancel() public onlyController {
        // Reset pending daily limit change.
        pendingSpendLimit = 0;
        // Reset the submitted operation flag.
        _submittedSpendLimit = false;
    }
}

/// @title Asset wallet with extra security features.
/// @author TokenCard
contract Vault is Whitelist, SpendLimit {
    // Events
    event Deposit(address _from, uint _amount);
    event Transfer(address _to, address _asset, uint _amount);

    // Storage
    address public oracle;

    /// @dev Construct a wallet with an owner and a controller.
    constructor(address _owner, address _oracle, address[] _controllers) public {
        currentDay = now;
        spendLimit = 1 ether;
        _spendAvailable = spendLimit;
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
                etherValue = Oracle(oracle).convert(_asset, _amount);
            } else {
                etherValue = _amount;
            }
            // Require that the value is under remaining limit.
            if (now > currentDay + 24 hours) {
                uint extraDays = (now - currentDay) / 24 hours;
                currentDay += extraDays * 24 hours;
                _spendAvailable = spendLimit;
            }
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
        emit Transfer(_to, _asset, _amount);
    }
}

contract Wallet is Vault {

    event SetTopupLimit(uint _amount);
    event TopupGas(address _sender, address _owner, uint _amount);

    uint constant private MINIMUM_TOPUP_LIMIT = 1 finney;
    uint constant private MAXIMUM_TOPUP_LIMIT = 100 finney;

    uint public topupLimit;
    uint private _topupAvailable;

    uint public pendingTopupLimit;
    bool private _submittedTopupLimit;
    bool private _initializedTopupLimit;

    constructor(address _owner, address _oracle, address[] _controllers) Vault(_owner, _oracle, _controllers) public {
        topupLimit = MAXIMUM_TOPUP_LIMIT;
        _topupAvailable = topupLimit;
    }

    /// @dev Returns the available daily gas top up balance - accounts for daily limit reset.
    /// @return amount of gas in wei.
    function topupAvailable() public view returns (uint) {
        if (now > currentDay + 24 hours) {
            return topupLimit;
        } else {
            return _topupAvailable;
        }
    }

    /// @dev Initialize a daily gas top up limit.
    /// @param _amount is the gas top up amount in wei.
    function initializeTopupLimit(uint _amount) public onlyOwner {
        require(!_initializedTopupLimit);
        // Require that the limit amount is within the acceptable range.
        require(MINIMUM_TOPUP_LIMIT <= _amount && _amount <= MAXIMUM_TOPUP_LIMIT);
        // Set the topup limit to the provided amount.
        topupLimit = _amount;
        _initializedTopupLimit = true;
        emit SetTopupLimit(_amount);
    }

    /// @dev Set a daily topup top up limit.
    /// @param _amount is the daily topup limit amount in wei.
    function setTopupLimit(uint _amount) public onlyOwner {
        // Check if this operation has been already submitted.
        require(!_submittedTopupLimit);
        // Require that the limit amount is within the acceptable range.
        require(MINIMUM_TOPUP_LIMIT <= _amount && _amount <= MAXIMUM_TOPUP_LIMIT);
        // Assign the provided amount to pending daily limit change.
        pendingSpendLimit = _amount;
        // Flag the operation as submitted.
        _submittedTopupLimit = true;
    }

    /// @dev Confirm pending set top up limit operation.
    function setTopupLimitConfirm() public onlyController {
        // That a set topup limit operation has been submitted.
        require(_submittedTopupLimit);
        // Assert that the pending topup limit amount is within the acceptable range.
        assert(MINIMUM_TOPUP_LIMIT <= pendingSpendLimit && pendingSpendLimit <= MAXIMUM_TOPUP_LIMIT);
        // Set the daily limit to the pending amount.
        topupLimit = pendingSpendLimit;
        // Reset the submission flag.
        _submittedTopupLimit = false;
        emit SetTopupLimit(pendingSpendLimit);
    }

    /// @dev Cancel pending set top up limit operation.
    function setTopupLimitCancel() public onlyController {
        // Reset pending daily limit change.
        pendingSpendLimit = 0;
        // Reset the submitted operation flag.
        _submittedTopupLimit = false;
    }

    /// @dev Refill owner's gas balance.
    /// @param _amount the amount of ether to transfer to the owner account in wei.
    function topupGas(uint _amount) public eitherOwnerOrController notZero(_amount) {
        // Account for the topup limit daily reset.
        if (now > currentDay + 24 hours) {
            uint extraDays = (now - currentDay) / 24 hours;
            currentDay += extraDays * 24 hours;
            _topupAvailable = topupLimit;
        }
        // Make sure the available topup is not zero.
        require(_topupAvailable > 0);
        // If amount is above available balance, use the entire balance.
        if (_amount > _topupAvailable) {
            _amount = _topupAvailable;
        }
        // Deduce the top up amount from available balance and transfer corresponding
        // ether to the owner's account.
        _topupAvailable -= _amount;
        owner.transfer(_amount);
        emit TopupGas(tx.origin, owner, _amount);
    }
}
