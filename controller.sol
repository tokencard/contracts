pragma solidity 0.4.23;

// The Token interface is a subset of the ERC20 specification.
interface Token {
    function transfer(address, uint) external returns (bool);
}

// The Card interface is the interface to a card smart contract.
interface Card {
    function changeOwner(address) external;
    function debit(address, uint) external;
    function hold(address, uint) external;
    function lock() external;
    function release(address, uint) external;
    function settle(address, uint) external;
}

// A Controller performs batch transfers to and from cards.
contract Controller {
    // Batch opcodes.
    byte private constant OP_DEBIT   = 0x00;
    byte private constant OP_CREDIT  = 0x01;
    byte private constant OP_HOLD    = 0x02;
    byte private constant OP_RELEASE = 0x03;
    byte private constant OP_SETTLE  = 0x04;

    // Owner of this contract.
    address public owner;

    // New owner in a two-phase ownership transfer.
    address public newOwner;

    // Next transfer ID to process.
    uint public next = 1;

    // Only allow the given address to call the method.
    modifier only(address a) {
        require(msg.sender == a);
        _;
    }

    // Construct a controller with the sender as the owner.
    constructor() public {
        owner = msg.sender;
    }

    // Ether sent to this contract may not be returned.
    function () external payable { }

    // Change owner in a two-phase ownership transfer.
    function changeOwner(address to) external only (owner) {
        newOwner = to;
    }

    // Accept ownership in a two-phase ownership transfer.
    function acceptOwnership() external only (newOwner) {
        owner = msg.sender;
        newOwner = 0x0;
    }

    // Process a transfer batch.
    function process(uint sequenceNumber, bytes operations, address[] cards, address[] tokens, uint[] amounts) external only (owner) {
        require(sequenceNumber <= next && sequenceNumber + operations.length > next);

        require(operations.length == cards.length);
        require(operations.length == tokens.length);
        require(operations.length == amounts.length);

        for (uint i = next - sequenceNumber; i < operations.length; i++) {
            if (operations[i] == OP_DEBIT) {
                Card(cards[i]).debit(tokens[i], amounts[i]);
            } else if (operations[i] == OP_CREDIT) {
                transfer(cards[i], tokens[i], amounts[i]);
            } else if (operations[i] == OP_HOLD) {
                Card(cards[i]).hold(tokens[i], amounts[i]);
            } else if (operations[i] == OP_RELEASE) {
                Card(cards[i]).release(tokens[i], amounts[i]);
            } else if (operations[i] == OP_SETTLE) {
                Card(cards[i]).settle(tokens[i], amounts[i]);
            } else {
                revert();
            }
        }

        next = sequenceNumber + operations.length;
    }

    // Transfer funds to an address.
    function transfer(address to, address token, uint amount) private {
        if (token != 0x0) {
            require(Token(token).transfer(to, amount));
        } else {
            to.transfer(amount);
        }
    }

    // Change the owner of a card smart contract on the owner's behalf.
    function reownCard(address card, address to) external only (owner) {
        Card(card).changeOwner(to);
    }

    // Lock a card.
    function lock(address card) external only (owner) {
        Card(card).lock();
    }
}
