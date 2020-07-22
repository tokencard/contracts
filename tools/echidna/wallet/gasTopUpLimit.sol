// SPDX-License-Identifier: GPLv3

pragma solidity ^0.6.11;

import "crytic-export/flattening/Wallet.sol";


contract Echidna {
    address internal echidnaDeployer = address(0x1);
    address internal echidnaAttacker = address(0x2);
    address payable internal echidnaOwner = address(0x3);
    address internal hevmCheatcode = address(0x7109709ECfa91a80626fF3989D68f67F5b1DD12D);
}


contract TEST is Echidna, GasTopUpLimit {
    uint256 initialLimit = 500 finney;

    constructor() public {
        _initializeGasTopUpLimit();
        _initializeOwnable(echidnaOwner, false);
        _initializeControllable(bytes32(0x0));
        _initializeENSResolvable(address(0x0));
    }

    // This function is used to simulate the user spending, so anyone can call it
    function spend(uint256 _amount) public {
        _gasTopUpLimit._enforceLimit(_amount);
    }

    // Uses the hevm `warp` cheatcode to increment the block.timestamp by 24 hours + 1
    function increaseDayNow() public {
        hevmCheatcode.call(abi.encodeWithSignature("warp(uint256)", now + 2 days));
    }

    function echidna_dailyGasTopUpLimitAvailable() public returns (bool) {
        increaseDayNow();
        return gasTopUpLimitAvailable() <= initialLimit;
    }

    function echidna_constantGasTopUpLimitAvailable() public view returns (bool) {
        return gasTopUpLimitAvailable() <= initialLimit;
    }
}
