pragma solidity ^0.4.25;

contract Owned {
    address public owner;

    modifier onlyOwner() {
        if (msg.sender != owner) revert();
        _;
    }

    address newOwner;

    function changeOwner(address _newOwner) public onlyOwner {
        newOwner = _newOwner;
    }

    function acceptOwnership() public {
        if (msg.sender == newOwner) {
            owner = newOwner;
        }
    }
}

//from Zeppelin
contract SafeMath {
    function safeMul(uint a, uint b) internal pure returns (uint) {
        uint c = a * b;
        assert(a == 0 || c / a == b);
        return c;
    }

    function safeSub(uint a, uint b) internal pure returns (uint) {
        assert(b <= a);
        return a - b;
    }

    function safeAdd(uint a, uint b) internal pure returns (uint) {
        uint c = a + b;
        assert(c>=a && c>=b);
        return c;
    }

    /* function assert(bool assertion) internal {
        if (!assertion) throw;
    } */
}

interface TokenHolder {
    function burn(address, uint) external returns (bool);
}

/// @title Token is a mock ERC20 token used for testing.
contract BurnerToken is SafeMath, Owned {

    event Transfer(address indexed from, address indexed to, uint256 amount);
    event Burn(address burner, uint amount);

    /// @dev Total supply of tokens in circulation.
    uint public totalSupply;

    /// @dev Balances for each account.
    mapping(address => uint) public balanceOf;

    /// @dev Transfer a token. This throws on insufficient balance.
    function transfer(address to, uint amount) public returns (bool) {
        require(balanceOf[msg.sender] >= amount);
        balanceOf[msg.sender] -= amount;
        balanceOf[to] += amount;
        emit Transfer(msg.sender, to, amount);
        return true;
    }

    /// @dev Credit an address.
    function credit(address to, uint amount) public returns (bool) {
        balanceOf[to] += amount;
        totalSupply += amount;
        return true;
    }

    /// @dev Debit an address.
    function debit(address from, uint amount) public {
        balanceOf[from] -= amount;
        totalSupply -= amount;
    }

    //Holds accumulated dividend tokens other than TKN
    TokenHolder tokenholder;

    //once locked, can no longer upgrade tokenholder
    bool lockedTokenHolder;

    function lockTokenHolder() public onlyOwner {
        lockedTokenHolder = true;
    }

    //while unlocked,
    //this gives owner lots of power over held dividend tokens
    //effectively can deny access to all accumulated tokens
    //thus crashing TKN value
    function setTokenHolder(address _th) public onlyOwner {
        if (lockedTokenHolder) revert();
        tokenholder = TokenHolder(_th);
    }

    function burn(uint _amount) public returns (bool result) {
        if (_amount > balanceOf[msg.sender]) return false;
        balanceOf[msg.sender] = safeSub(balanceOf[msg.sender], _amount);
        totalSupply = safeSub(totalSupply, _amount);
        result = tokenholder.burn(msg.sender, _amount);
        if (!result) revert();
        emit Burn(msg.sender, _amount);
    }

}

/* contract TokenHolder {
    function burn(address _burner, uint _amount)
    returns (bool result) {
        return false;
    }
} */
