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
    address public controller;
    address public owner;
    address public newOwner;
    address[] public whitelist;

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

    // Transfer contract ownership to another Ethereum address.
    function transferOwnership(address account) public only(owner) {
        newOwner = account;
    }

    // Accept contract ownership.
    function acceptOwnership() public only(newOwner) {
        owner = newOwner;
    }

    // Add an address to a whitelist.
    // TODO: This should be multi sig to prevent attacker changing the addresses.
    function addWhitelistAddress(address address_) public only(owner) returns (bool) {
        return true;
    }

    // TODO: This should be multi sig to prevent attacker changing the addresses.
    function removeWhitelistAddress(address address_) public only(owner) returns (bool) {
        return true;
    }

    // Transfer asset to an address.
    function transfer(address to, address token, uint amount) public only(owner) {
        if (amount == 0) {
            return;
        }
        require(Token(token).transfer(to, amount));

        emit Transfer(to, token, amount);
    }

    // Transfer ether to an address.
    function transfer(address to, uint amount) public only(owner) {
        if (amount == 0) {
            return;
        }
        to.transfer(amount);

        emit Transfer(to, 0x0, amount);
    }

    // TODO: When we have a list of private addresses, these addresses should be able to top up user's gas and cover the transaction cost.
    function changeController(address newController) public only(controller) {
        controller = newController;
    }


    // Refill owner's gas balance.
    function refillGas(uint256 amount) public either(owner, controller) {
        // TODO: Set cap to how much ether can be transferred.



        // TODO: Transfer ether to owner's account

    }
}
