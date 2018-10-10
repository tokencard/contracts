pragma solidity ^0.4.25;


/// @title Token is a mock ERC20 token used for testing.
contract Token {
    event Transfer(address indexed from, address indexed to, uint256 amount);
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
}
