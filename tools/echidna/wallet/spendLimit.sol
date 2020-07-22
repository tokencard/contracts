// SPDX-License-Identifier: GPLv3

pragma solidity ^0.6.11;

import "crytic-export/flattening/Wallet.sol";


contract Echidna {
    address payable internal echidnaOwner = address(0x1);
    address internal echidnaAttacker = address(0x2);
    address internal hevmCheatcode = address(0x7109709ECfa91a80626fF3989D68f67F5b1DD12D);
}


contract TEST is Echidna, SpendLimit {
    uint256 initialLimit = 10000;

    constructor() public {
        _initializeSpendLimit(initialLimit);
        _initializeOwnable(echidnaOwner, false);
        _initializeControllable(bytes32(0x0));
        _initializeENSResolvable(address(0x0));
        // This simulates calling to setSpendLimit (which is an external function).
        _spendLimit._setLimit(initialLimit);
    }

    // This function is used to simulate the user spending, so anyone can call it.
    function spend(uint256 _amount) public {
        _spendLimit._enforceLimit(_amount);
    }

    // Uses the hevm `warp` cheatcode to increment the block.timestamp by 24 hours + 1.
    function increaseDayNow() public {
        hevmCheatcode.call(abi.encodeWithSignature("warp(uint256)", now + 2 days));
    }

    // The owner cannot be changed.
    function echidna_fixedOwner() public view returns (bool) {
        return owner() == echidnaOwner;
    }

    // Spend limit available is bounded by the spend limit value.
    function echidna_boundedSpendLimitAvailable() public view returns (bool) {
        return spendLimitAvailable() <= spendLimitValue();
    }

    // Spend limit value cannot be changed (by an external user).
    function echidna_fixedSpendLimitValue() public view returns (bool) {
        return spendLimitValue() == initialLimit;
    }

    // After a day, your limit available should reset.
    function echidna_dailySpendLimitAvailable() public returns (bool) {
        increaseDayNow();
        spend(0);
        return spendLimitAvailable() == spendLimitValue();
    }
}
