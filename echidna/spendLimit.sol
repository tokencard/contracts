pragma solidity ^0.5.15;

import "contracts/wallet.sol";


contract Echidna {
    address payable internal echidna_owner = address(0x1);
    address internal echidna_attacker = address(0x2);
    address internal hevm_cheatcode = address(0x7109709ECfa91a80626fF3989D68f67F5b1DD12D);
}


contract TEST is Echidna, SpendLimit {
    uint256 initialLimit = 10000;

    constructor() public SpendLimit(initialLimit) Ownable(echidna_owner, false) Controllable(bytes32(0x0)) ENSResolvable(address(0x0)) {
        // this simulates calling to setSpendLimit (which is an external function)
        _spendLimit._setLimit(initialLimit);
    }

    // This function is used to simulate the user spending, so anyone can call it
    function spend(uint256 _amount) public {
        _spendLimit._enforceLimit(_amount);
    }

    // Uses the hevm `warp` cheatcode to increment the block.timestamp by 24 hours + 1
    function increaseDayNow() public {
        hevm_cheatcode.call(abi.encodeWithSignature("warp(uint256)", now + 2 days));
    }

    // The owner cannot be changed
    function echidna_fixedOwner() public returns (bool) {
        return owner() == echidna_owner;
    }

    // spend limit available is bounded by the spend limit value
    function echidna_boundedSpendLimitAvailable() public returns (bool) {
        return this.spendLimitAvailable() <= this.spendLimitValue();
    }

    // spend limit value cannot be changed (by an external user)
    function echidna_fixedSpendLimitValue() public returns (bool) {
        return this.spendLimitValue() == initialLimit;
    }

    // after a day, your limit available should reset
    function echidna_dailySpendLimitAvailable() public returns (bool) {
        increaseDayNow();
        spend(0);
        return this.spendLimitAvailable() == this.spendLimitValue();
    }
}
