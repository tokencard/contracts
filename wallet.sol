pragma solidity ^0.4.24;


// The Token interface is a subset of the ERC20 specification.
interface Token {
    function transfer(address, uint) external returns (bool);
    function balanceOf(address) external constant returns (uint);
}


contract Wallet {
    // Emitted contract events.
    event TopUpGas(address sender, address owner, uint amount);
    event Transfer(address to, address token, uint amount);
    event Deposit(address from, uint amount);

    uint private constant GAS_TOPUP_LIMIT = 10 finney;

    address public controller;
    address public owner;

    // Construct a wallet with an owner and a controller.
    constructor(address o, address c) public {
        owner = o;
        controller = c;
    }

    // Executable only by the specified address.
    modifier only(address a) {
        require(msg.sender == a);
        _;
    }

    // Executable by either of the two addresses.
    modifier either(address a, address b) {
        require(msg.sender == a || msg.sender == b);
        _;
    }

    // Ether can be deposited from any source, so this contract must be payable by anyone.
    function() external payable {
        if (msg.value > 0) {
            emit Deposit(msg.sender, msg.value);
        }
    }

    // Returns the amount of a token owned by the contract.
    function balance(address token) public constant returns (uint) {
        return Token(token).balanceOf(this);
    }
    // Returns the amount of ether owned by the contract.
    function balance() public constant returns (uint) {
        return address(this).balance;
    }

    // Transfer asset to an address.
    function transfer(address to, address token, uint amount) public only(owner) returns (bool) {
        if (amount == 0) {
            return;
        }
        require(Token(token).transfer(to, amount));

        emit Transfer(to, token, amount);
        return true;
    }

    // Transfer ether to an address.
    function transfer(address to, uint amount) public only(owner) returns (bool) {
        if (amount == 0) {
            return;
        }
        to.transfer(amount);

        emit Transfer(to, 0x0, amount);
        return true;
    }

    // Refill owner's gas balance.
    function topUpGas(uint amount) public either(owner, controller) returns (bool) {
        // Require that the amount not exceed the cap.
        require(amount <= GAS_TOPUP_LIMIT);

        // Transfer ether to owner's account
        owner.transfer(amount);

        emit TopUpGas(tx.origin, owner, amount);
        return true;
    }
}
