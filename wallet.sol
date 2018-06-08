pragma solidity ^0.4.24;


/// @title The Token interface is a subset of the ERC20 specification.
interface Token {
    function transfer(address, uint) external returns (bool);
    function balanceOf(address) external view returns (uint);
}

/// @title The Oracle interface provides exchange rates for ERC20 tokens in wei.
interface Oracle {
    function rates(address) external view returns (uint);
}

/// @title Control handles wallet access control.
contract Control {
    address public controller;
    address public owner;

    /// @dev Executable only by the specified address.
    modifier only(address _a) {
        require(msg.sender == _a);
        _;
    }

    /// @dev Executable by either of the two addresses.
    modifier either(address _a, address _b) {
        require(msg.sender == _a || msg.sender == _b);
        _;
    }
}

/// @title Whitelist provides payee-whitelist functionality.
contract Whitelist is Control {

    event WhitelistAddition(address[] _addresses);
    event WhitelistRemoval(address[] _addresses);

    mapping(address => bool) public whitelist;
    address[] public pendingAddition;
    address[] public pendingRemoval;
    bool internal submittedAddition;
    bool internal submittedRemoval;

    /// @dev Add addresses to the whitelist.
    /// @param _addresses are the Ethereum addresses to be whitelisted.
    function addToWhitelist(address[] _addresses) public only(owner) {
        // Check if this operation has been already submitted.
        require(!submittedAddition);
        // Limit the amount of addresses that can be whitelisted at one time.
        require(_addresses.length <= 20);
        // Add each of the addresses to the pending addition list.
        for (uint i = 0; i < _addresses.length; i++) {
            pendingAddition.push(_addresses[i]);
        }
        // Flag the operation as submitted.
        submittedAddition = true;
    }

    /// @dev Confirm pending whitelist addition.
    function addToWhitelistConfirm() public only(controller) {
        require(pendingAddition.length > 0 && submittedAddition);
        // Whitelist pending addresses.
        for (uint i = 0; i < pendingAddition.length; i++) {
            whitelist[pendingAddition[i]] = true;
        }
        // Reset the submission flag.
        submittedAddition = false;
        emit WhitelistAddition(pendingAddition);
    }

    /// @dev Cancel pending whitelist addition.
    function addToWhitelistCancel() public only(controller) {
        // Reset pending addresses.
        delete pendingAddition;
        // Reset the submitted operation flag.
        submittedAddition = false;
    }

    /// @dev Remove addresses from the whitelist.
    /// @param _addresses are the Ethereum addresses to be removed.
    function removeFromWhitelist(address[] _addresses) public only(owner) {
        // Check if this operation has been already submitted.
        require(!submittedRemoval);
        // Limit the amount of addresses that can be removed at one time.
        require(_addresses.length <= 20);
        // Add each of the addresses to the pending removal list.
        for (uint i = 0; i < _addresses.length; i++) {
            pendingRemoval.push(_addresses[i]);
        }
        // Flag the operation as submitted.
        submittedRemoval = true;
    }

    /// @dev Confirm pending removal of whitelisted addresses.
    function removeFromWhitelistConfirm() public only(controller) {
        require(pendingRemoval.length > 0 && submittedRemoval);
        // Remove pending addresses.
        for (uint i = 0; i < pendingRemoval.length; i++) {
            whitelist[pendingRemoval[i]] = false;
        }
        // Reset the submission flag.
        submittedRemoval = false;
        emit WhitelistRemoval(pendingRemoval);
    }

    /// @dev Cancel pending removal of whitelisted addresses.
    function removeFromWhitelistCancel() public only(controller) {
        // Reset pending addresses.
        delete pendingRemoval;
        // Reset the submitted operation flag.
        submittedRemoval = false;
    }
}

/// @title DailyLimit provides daily spend limit functionality.
contract DailyLimit is Control {

    event SetDailyLimit(uint _amount);

    uint public dailyLimit;
    uint public currentDay;
    uint internal availableToday;

    uint public pendingDailyLimit;
    bool internal submittedDailyLimit;

    /// @dev Set a daily transfer limit for non-whitelisted addresses.
    /// @param _amount is the daily limit amount in wei.
    function setDailyLimit(uint _amount) public only(owner) {
        // Check if this operation has been already submitted.
        require(!submittedDailyLimit);
        // Assign the provided amount to pending daily limit change.
        pendingDailyLimit = _amount;
        // Flag the operation as submitted.
        submittedDailyLimit = true;
    }

    /// @dev Confirm pending set daily limit operation.
    function setDailyLimitConfirm() public only(controller) {
        require(submittedDailyLimit);
        // Set the daily limit to the pending amount.
        dailyLimit = pendingDailyLimit;
        // Lower the available balance if it exceeds the new daily limit.
        if (availableToday > dailyLimit) {
            availableToday = dailyLimit;
        }
        // Reset the submission flag.
        submittedDailyLimit = false;
        emit SetDailyLimit(pendingDailyLimit);
    }

    /// @dev Cancel pending set daily limit operation.
    function setDailyLimitCancel() public only(controller) {
        // Reset pending daily limit change.
        pendingDailyLimit = 0;
        // Reset the submitted operation flag.
        submittedDailyLimit = false;
    }
}

/// @title Asset wallet with extra security features.
/// @author TokenCard
contract Wallet is Whitelist, DailyLimit {
    // Events
    event Deposit(address _from, uint _amount);
    event Transfer(address _to, address _token, uint _amount);
    event TopUpGas(address _sender, address _owner, uint _amount);

    // Constants
    uint private constant GAS_TOPUP_LIMIT = 10 finney;

    // Storage
    address public oracle;

    /// @dev Construct a wallet with an owner and a controller.
    constructor(address _owner, address _controller, address _oracle) public {
        owner = _owner;
        controller = _controller;
        currentDay = now;
        dailyLimit = 2 ether;
        availableToday = dailyLimit;
        oracle = _oracle;
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
    /// @param _token address of an ERC20 token, 0x0 is used for ether.
    /// @return balance associated with the wallet address in wei.
    function balance(address _token) public view returns (uint) {
        if (_token != 0x0) {
            return Token(_token).balanceOf(this);
        } else {
            return address(this).balance;
        }
    }

    /// @dev Returns the available daily balance - accounts for daily limit reset.
    /// @return amount of ether in wei.
    function available() public view returns (uint) {
        if (now > currentDay + 24 hours) {
            return dailyLimit;
        } else {
            return availableToday;
        }
    }

    /// @dev Transfer asset to an address.
    /// @param _to recipient address.
    /// @param _token address of an ERC20 token, 0x0 is used for ether.
    /// @param _amount is the amount of tokens to be transferred in base units.
    function transfer(address _to, address _token, uint _amount) public only(owner) notZero(_amount) {
        // If address is not whitelisted, take daily limit into account.
        if (!whitelist[_to]) {
            // Convert token amount to ether value.
            uint etherValue;
            if (_token != 0x0) {
                etherValue = Oracle(oracle).rates(_token) * _amount;
            } else {
                etherValue = _amount;
            }
            // Require that the value is under remaining limit.
            if (now > currentDay + 24 hours) {
                uint extraDays = (now - currentDay) / 24 hours;
                currentDay += extraDays * 24 hours;
                availableToday = dailyLimit;
            }
            require(etherValue <= availableToday);
            // Update the available limit.
            availableToday -= etherValue;
        }
        // Transfer token or ether based on the provided address.
        if (_token != 0x0) {
            require(Token(_token).transfer(_to, _amount));
        } else {
            _to.transfer(_amount);
        }

        emit Transfer(_to, _token, _amount);
    }

    /// @dev Refill owner's gas balance.
    /// @param _amount the amount of ether to transfer to the owner account in wei.
    function topUpGas(uint _amount) public either(owner, controller) notZero(_amount) {
        // Require that the amount not exceed the top up limit.
        require(_amount <= GAS_TOPUP_LIMIT);
        // Transfer ether to owner's account
        owner.transfer(_amount);

        emit TopUpGas(tx.origin, owner, _amount);
    }
}
