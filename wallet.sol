pragma solidity ^0.4.24;


// The Token interface is a subset of the ERC20 specification.
interface Token {
    function transfer(address, uint) external returns (bool);
    function balanceOf(address) external constant returns (uint);
}


contract Wallet {
    // Emitted contract events.
    event TopUpGas(address _sender, address _owner, uint _amount);
    event Transfer(address _to, address _token, uint _amount);
    event Deposit(address _from, uint _amount);

    uint private constant GAS_TOPUP_LIMIT = 10 finney;

    address public controller;
    address public owner;

    // Construct a wallet with an owner and a controller.
    constructor(address _owner, address _controller) public {
        owner = _owner;
        controller = _controller;
    }

    // Executable only by the specified address.
    modifier only(address _a) {
        require(msg.sender == _a);
        _;
    }

    // Executable by either of the two addresses.
    modifier either(address _a, address _b) {
        require(msg.sender == _a || msg.sender == _b);
        _;
    }

    // Ether can be deposited from any source, so this contract must be payable by anyone.
    function() external payable {
        if (msg.value > 0) {
            emit Deposit(msg.sender, msg.value);
        }
    }

    // Returns the amount of a token owned by the contract.
    function balance(address _token) public view returns (uint) {
        return Token(_token).balanceOf(this);
    }
    // Returns the amount of ether owned by the contract.
    function balance() public view returns (uint) {
        return address(this).balance;
    }

    // Transfer asset to an address.
    function transfer(address _to, uint _amount, address _token) public only(owner) {
        require(_amount != 0);
        require(Token(_token).transfer(_to, _amount));

        emit Transfer(_to, _token, _amount);
    }

    // Transfer ether to an address.
    function transfer(address _to, uint _amount) public only(owner) {
        require(_amount != 0);
        _to.transfer(_amount);
        emit Transfer(_to, 0x0, _amount);
    }

    // Refill owner's gas balance.
    function topUpGas(uint _amount) public either(owner, controller) {
        // Require that the amount not exceed the top up limit.
        require(_amount <= GAS_TOPUP_LIMIT);

        // Transfer ether to owner's account
        owner.transfer(_amount);

        emit TopUpGas(tx.origin, owner, _amount);
    }
}
