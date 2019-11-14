pragma solidity ^0.5.10;

import "contracts/wallet.sol";

contract CryticInterface{
    address internal crytic_owner = address(0x41414141);
    address internal crytic_attacker = address(0x43434343);

    address internal hevm_cheatcode = address(0x7109709ecfa91a80626ff3989d68f67f5b1dd12d);
}

contract TEST is CryticInterface, SpendLimit {

    uint initialLimit = 10000;

    constructor() SpendLimit(initialLimit) Ownable(crytic_owner, false) Controllable(bytes32(0x0)) ENSResolvable(address(0x0)) public {
        // this simulates calling to setSpendLimit (which is an external function)
        _setLimit(_spendLimit, initialLimit);
    }

    // This function is used to simulate the user spending, so anyone can call it
    function spend(uint _amount) public {
        _enforceLimit(_spendLimit,_amount);
    }

    // uses the hevm `warp` cheatcode to increment the block.timestamp by 24 hours + 1
    // see https://github.com/dapphub/dapptools/blob/79744a960979f1a593d8e04b04a4df9b6c533fa9/src/hevm/src/EVM.hs#L1333-L1338
    // for more info
    function increaseDayNow() public {
        hevm_cheatcode.call(bytes4(keccak256("warp(uint256)")), now + (1 days) + 1);
    }

    // The owner cannot be changed
    function crytic_fixed_owner() public returns (bool) {
        return owner() == crytic_owner;
    }
    
    // spend limit available is bounded by the spend limit value
    function crytic_bounded_spendLimitAvailable() public returns (bool) {
        return this.spendLimitAvailable() <= this.spendLimitValue();
    }

    // spend limit value cannot be changed (by an external user)
    function crytic_fixed_spendLimitValue() public returns (bool) {
        return this.spendLimitValue() == initialLimit;
    }

    // after a day, your limit available should reset
    function crytic_daily_spendLimitAvailable() public returns (bool) {
        increaseDayNow();
        spend(0);
        return this.spendLimitAvailable() == this.spendLimitValue();
    }
}
