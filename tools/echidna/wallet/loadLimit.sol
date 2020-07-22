// SPDX-License-Identifier: GPLv3

pragma solidity 0.6.11;

import "crytic-export/flattening/Wallet.sol";


contract Echidna {
    address payable internal echidnaOwner = address(0x1);
    address internal echidnaAttacker = address(0x2);
    address internal hevmCheatcode = address(0x7109709ECfa91a80626fF3989D68f67F5b1DD12D);
}


contract TEST is Echidna, LoadLimit {
    uint256 initialLimit = 10000000000;

    constructor() public {
        _initializeLoadLimit(bytes32(0x0));
        _initializeOwnable(echidnaOwner, false);
        _initializeControllable(bytes32(0x0));
        _initializeENSResolvable(address(0x0));
        _loadLimit = DailyLimitTrait.DailyLimit(initialLimit, initialLimit, now, 0, false);
    }

    // This function is used to simulate the user loading a card, so anyone can call it
    function load(uint256 _amount) public {
        _loadLimit._enforceLimit(_amount);
    }

    // Uses the hevm `warp` cheatcode to increment the block.timestamp by 24 hours + 1.
    function increaseDayNow() public {
        hevmCheatcode.call(abi.encodeWithSignature("warp(uint256)", now + 2 days));
    }

    // The owner cannot be changed
    function echidna_fixedOwner() public view returns (bool) {
        return owner() == echidnaOwner;
    }

    // Load limit available is bounded by the spend limit value.
    function echidna_boundedLoadLimitAvailable() public view returns (bool) {
        return loadLimitAvailable() <= loadLimitValue();
    }

    // Load limit value cannot be changed (by an external user).
    function echidna_fixedLoadLimitValue() public view returns (bool) {
        return loadLimitValue() == initialLimit;
    }

    // After a day, your limit available should reset.
    function echidna_dailyLoadLimitAvailable() public returns (bool) {
        increaseDayNow();
        load(0);
        return loadLimitAvailable() == loadLimitValue();
    }
}
