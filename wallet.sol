pragma solidity ^0.4.24;


// The Token interface is a subset of the ERC20 specification.
interface Token {
    function transfer(address, uint) external returns (bool);
    function balanceOf(address) external constant returns (uint);
}

/// @title Asset wallet with extra security features.
/// @author TokenCard
contract Wallet {
    // Events
    event Deposit(address _from, uint _amount);
    event Transfer(address _to, address _token, uint _amount);
    event AddToWhitelist(address[] _addresses);
    event RemoveFromWhitelist(address[] _addresses);
    event SetDailyLimit(uint _amount);
    event TopUpGas(address _sender, address _owner, uint _amount);

    // Constants
    uint private constant GAS_TOPUP_LIMIT = 10 finney;

    // Storage
    address public controller;
    address public owner;

    uint public dailyLimit;
    uint public currentDay;
    uint private availableToday;
    uint private pendingSetDailyLimit;

    mapping(address => bool) public whitelisted;
    address[] private pendingAddToWhitelist;
    address[] private pendingRemoveFromWhitelist;

    /// @dev Construct a wallet with an owner and a controller.
    constructor(address _owner, address _controller) public {
        owner = _owner;
        controller = _controller;
        currentDay = now;
    }

    /// @dev Checks if the value is not zero.
    modifier notZero(uint _value) {
        require(_value != 0);
        _;
    }

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

    /// @dev Returns the available daily limit - accounts for daily reset.
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
        if (!whitelisted[_to]) {
            // Convert token amount to ether value.
            uint etherValue = _tokenToEther(_token, _amount);
            // Require that the value is under remaining limit.
            if (now > currentDay + 24 hours) {
                uint extraDays = (now - currentDay) / 24 hours;
                currentDay += extraDays * 24 hours;
                availableToday = dailyLimit;
            }
            require(etherValue <= availableToday);
            // Update the remaining limit.
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

    /// @dev Convert tokens to the corresponding ether value in wei.
    /// @param _token address of an ERC20 token.
    /// @param _amount is the amount of tokens in base units.
    /// @return token value converted to ether (wei).
    function _tokenToEther(address _token, uint _amount) private returns (uint) {
        // TODO: Use Bancor to get current Token/ETH rate.
        return 1000;
    }

    /// @dev Add pending Ethereum addresses to be whitelisted.
    /// @param _addresses are the Ethereum addresses to be whitelisted.
    function addToWhitelist(address[] _addresses) public only(owner) {
        for (uint i = 0; i < _addresses.length; i++) {
            pendingAddToWhitelist.push(_addresses[i]);
        }
    }

    /// @dev Confirm pending whitelist addresses and update the whitelist (controller only).
    function confirmAddToWhitelist() public only(controller) {
        require(pendingAddToWhitelist.length > 0);
        for (uint i = 0; i < pendingAddToWhitelist.length; i++) {
            whitelisted[pendingAddToWhitelist[i]] = true;
        }
        emit AddToWhitelist(pendingAddToWhitelist);
    }

    /// @dev Add a pending removal of whitelisted addresses.
    /// @param _addresses are the Ethereum addresses to be removed.
    function removeFromWhitelist(address[] _addresses) public only(owner) {
        for (uint i = 0; i < _addresses.length; i++) {
            pendingRemoveFromWhitelist.push(_addresses[i]);
        }
    }

    /// @dev Confirm pending removal of whitelisted addresses (controller only).
    function confirmRemoveFromWhitelist() public only(controller) {
        require(pendingRemoveFromWhitelist.length > 0);
        for (uint i = 0; i < pendingRemoveFromWhitelist.length; i++) {
            whitelisted[pendingRemoveFromWhitelist[i]] = false;
        }
        emit RemoveFromWhitelist(pendingRemoveFromWhitelist);
    }

    /// @dev Set a daily transfer limit for non-whitelisted addresses.
    /// @param _amount is the daily limit amount in wei.
    function setDailyLimit(uint _amount) public only(owner) {
        pendingSetDailyLimit = _amount;
    }

    /// @dev Confirm the pending set daily limit operation.
    function confirmSetDailyLimit() public only(controller) {
        dailyLimit = pendingSetDailyLimit;
        emit SetDailyLimit(pendingSetDailyLimit);
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
