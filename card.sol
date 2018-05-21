pragma solidity 0.4.23;

// The Token interface is a subset of the ERC20 specification.
interface Token {
    function balanceOf(address) constant external returns (uint);
    function transfer(address, uint) external returns (bool);
}

// A Card contract holds funds and allows a controller to place holds on it.
contract Card {
    // Controller contract address.
    address public controller;

    // Owner of the card contract.
    address public owner;

    // Unlock the card contract at this block. If zero, the card contract is locked.
    uint public unlockAt = 0;

    // Overdraft amounts by token owed to the controller.
    mapping (address => uint) public overdraft;

    // Only allow the given address to call the method.
    modifier only(address a) {
        require(msg.sender == a);
        _;
    }

    // Either of these addresses may call the method.
    modifier either(address a, address b) {
        require(msg.sender == a || msg.sender == b);
        _;
    }

    // Construct a card contract owned by the sender with the given controller.
    constructor(address c) public {
        controller = c;
        owner = msg.sender;
    }

    // Card contracts may be topped up from any source, so this contract must be payable by anyone.
    function () external payable { }

    // Change the owner to the given address.
    function changeOwner(address o) public either (owner, controller) {
        owner = o;
    }

    // Unlock the card contract, allowing the owner to withdraw funds.
    function unlock() public only (owner) {
        unlockAt = block.number + 80000;
    }

    // Lock the card contract, allowing funds to be spent with swipes.
    function lock() public either (owner, controller) {
        unlockAt = 0;
    }

    // Balance returns the amount of a token owned by the contract. This
    // does not take overdraft into consideration.
    function balance(address token) constant public returns (uint) {
        if (token != 0x0) {
            return Token(token).balanceOf(this);
        } else {
            return address(this).balance;
        }
    }

    // Available returns the amount of a token available for spending.
    function available(address token) constant public returns (uint) {
        uint avail = balance(token);

        if (avail <= overdraft[token]) {
            return 0;
        }

        avail -= overdraft[token];

        return avail;
    }

    // Debit the card contract, adding to the overdraft amount if this exceeds the
    // balance of the card contract.
    function debit(address token, uint amount) public only (controller) {
        uint avail = available(token);

        // Only the available amount can be transferred, so anything
        // greater is overdraft.
        if (amount > avail) {
            require(overdraft[token] + (amount - avail) >= overdraft[token]);
            overdraft[token] += amount - avail;
            amount = avail;
        }

        transfer(controller, token, amount);
    }

    // Hold an amount on the card contract, adding to the overdraft balance
    // without settling.
    function hold(address token, uint amount) public only (controller) {
        require(overdraft[token] + amount >= overdraft[token]);
        overdraft[token] += amount;
    }

    // Release a hold, reducing the overdraft balance without settling.
    // This will not reduce overdraft below zero.
    function release(address token, uint amount) public only (controller) {
        if (overdraft[token] > amount) {
            overdraft[token] -= amount;
        } else {
            overdraft[token] = 0;
        }
    }

    // Settle the overdraft balance for the given token up to the given
    // amount.
    function settle(address token, uint amount) public only (controller) {
        if (amount > overdraft[token]) {
            amount = overdraft[token];
        }

        if (amount == 0) {
            return;
        }

        uint avail = balance(token);

        if (avail == 0) {
            return;
        }

        if (amount > avail) {
            overdraft[token] -= avail;
            amount = avail;
        } else if (amount > 0) {
            overdraft[token] -= amount;
        }

        transfer(controller, token, amount);
    }

    // Withdraw available funds to the owner. The card contract must be unlocked.
    function withdraw(address token, uint amount) public only (owner) {
        require(unlockAt > 0 && block.number >= unlockAt);
        require(amount <= available(token));

        transfer(owner, token, amount);
    }

    // Transfer funds to an address.
    function transfer(address to, address token, uint amount) private {
        if (amount == 0) {
            return;
        }

        if (token != 0x0) {
            require(Token(token).transfer(to, amount));
        } else {
            to.transfer(amount);
        }
    }
}
