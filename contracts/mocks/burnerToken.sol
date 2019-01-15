pragma solidity ^0.4.25;

import "../internals/ownable.sol";
import "../externals/SafeMath.sol";

/// @title ERC20 interface is a subset of the ERC20 specification.
interface TokenHolder {
    function payOut(address, uint) external returns (bool);
}

/// @title burnerToken is a mock ERC20 token with token burning capabilities used for testing.
contract BurnerToken is Ownable {

    using SafeMath for uint256;

    event Transfer(address indexed from, address indexed to, uint amount);
    event Burn(address burner, uint amount);

    //Holds accumulated dividend tokens other than TKN
    TokenHolder tokenholder;

    /// @dev Total supply of tokens in circulation.
    uint public totalSupply;

    /// @dev Balances for each account.
    mapping(address => uint) public balanceOf;

    //once locked, can no longer upgrade tokenholder
    bool public lockedTokenHolder;

    constructor(address _owner, bool _transferable) Ownable(_owner, _transferable) public {

    }

    function lockTokenHolder() external onlyOwner {
        lockedTokenHolder = true;
    }

    //while unlocked,
    //this gives owner lots of power over held dividend tokens
    //effectively can deny access to all accumulated tokens
    //thus crashing TKN value
    function setTokenHolder(address _th) external onlyOwner {
        require(!lockedTokenHolder, "tokenHolder is locked");
        tokenholder = TokenHolder(_th);
    }

    /// @dev Transfer a token. This reverts on insufficient balance.
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

    /// @dev Burn tokens.
    function burn(uint _amount) external returns (bool result) {
        require(totalSupply > 0);
        require(_amount <= totalSupply);
        //assert(_amount <= balanceOf[msg.sender]); //"punish" the sender
        //require(balanceOf[msg.sender] >= _amount); // FIXE ME: use require instead of returning false????
        if (_amount > balanceOf[msg.sender])
          return false;

        uint poolShare = totalSupply/_amount; //totalSupply and _amount != 0
        balanceOf[msg.sender] = balanceOf[msg.sender].sub(_amount);
        totalSupply = totalSupply.sub(_amount);

        result = tokenholder.payOut(msg.sender, poolShare);
        require(result, "burn failed");
        emit Burn(msg.sender, _amount);
    }

}
