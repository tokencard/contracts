pragma solidity ^0.5.10;

import "contracts/wallet.sol";

contract EchidnaInterface{
    address payable internal echidna_owner = address(0x41414141);
    address internal echidna_attacker = address(0x43434343);

    address internal hevm_cheatcode = address(0x7109709ECfa91a80626fF3989D68f67F5b1DD12D);
}

contract TEST is EchidnaInterface, SpendLimit {

    uint initialLimit = 10000;

    constructor() SpendLimit(initialLimit) Ownable(echidna_owner, false) Controllable(bytes32(0x0)) ENSResolvable(address(0x0)) public {
        // this simulates calling to setSpendLimit (which is an external function)
        _spendLimit._setLimit(initialLimit);
    }

    // This function is used to simulate the user spending, so anyone can call it
    function spend(uint _amount) public {
        _spendLimit._enforceLimit(_amount);
    }

    // uses the hevm `warp` cheatcode to increment the block.timestamp by 24 hours + 1
    // see https://github.com/dapphub/dapptools/blob/79744a960979f1a593d8e04b04a4df9b6c533fa9/src/hevm/src/EVM.hs#L1333-L1338
    // for more info
    function increaseDayNow() public {
        hevm_cheatcode.call(abi.encodeWithSignature("warp(uint256)", now + 2 days));
    }

    // The owner cannot be changed
    function echidna_fixed_owner() public returns (bool) {
        return owner() == echidna_owner;
    }
    
    // spend limit available is bounded by the spend limit value
    function echidna_bounded_spendLimitAvailable() public returns (bool) {
        return this.spendLimitAvailable() <= this.spendLimitValue();
    }

    // spend limit value cannot be changed (by an external user)
    function echidna_fixed_spendLimitValue() public returns (bool) {
        return this.spendLimitValue() == initialLimit;
    }

    // after a day, your limit available should reset
    function echidna_daily_spendLimitAvailable() public returns (bool) {
        increaseDayNow();
        spend(0);
        return this.spendLimitAvailable() == this.spendLimitValue();
    }
}
