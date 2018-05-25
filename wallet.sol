pragma solidity ^0.4.24;


// The Token interface is a subset of the ERC20 specification.
interface Token {
    function transfer(address, uint) external returns (bool);
    function balanceOf(address) external constant returns (uint);
}


contract Wallet {
    // Emitted contract events.
    event Transfer(address to, address token, uint amount);
    event Deposit(address from, uint amount);

    // Owner of the wallet.
    address public owner;

    // Construct a wallet owned by the sender with the given controller.
    constructor() public {
        owner = msg.sender;
    }

    // Wallets may be topped up from any source, so this contract must be payable by anyone.
    function () external payable {
        if (msg.value > 0) {
            emit Deposit(msg.sender, msg.value);
        }
    }

    // Balance returns the amount of a token owned by the contract.
    function balance(address token) public constant returns (uint) {
        if (token != 0x0) {
            return Token(token).balanceOf(this);
        } else {
            return address(this).balance;
        }
    }

    // Change the owner to the given address.
    function changeOwner(address o) public {
        owner = o;
    }

    // Transfer funds to an address.
    function transfer(address to, address token, uint amount) public {
        if (amount == 0) {
            return;
        }
        if (token != 0x0) {
            require(Token(token).transfer(to, amount));
        } else {
            to.transfer(amount);
        }
        emit Transfer(to, token, amount);
    }
}
