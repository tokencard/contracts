pragma solidity ^0.5.15;

import "contracts/wallet.sol";


contract Echidna {
    address internal echidna_deployer = address(0x1);
    address internal echidna_attacker = address(0x2);
    address payable internal echidna_owner = address(0x3);
    address internal hevm_cheatcode = address(0x7109709ECfa91a80626fF3989D68f67F5b1DD12D);
}


contract TEST is Echidna, GasTopUpLimit {
    uint256 initialLimit = 500 finney;

    constructor() public GasTopUpLimit() Ownable(echidna_owner, false) Controllable(bytes32(0x0)) ENSResolvable(address(0x0)) {}

    // This function is used to simulate the user spending, so anyone can call it
    function spend(uint256 _amount) public {
        _gasTopUpLimit._enforceLimit(_amount);
    }

    // Uses the hevm `warp` cheatcode to increment the block.timestamp by 24 hours + 1
    function increaseDayNow() public {
        hevm_cheatcode.call(abi.encodeWithSignature("warp(uint256)", now + 2 days));
    }

    function echidna_DailyGasTopUpLimitAvailable() public returns (bool) {
        increaseDayNow();
        return this.gasTopUpLimitAvailable() <= initialLimit;
    }

    function echidna_ConstantGasTopUpLimitAvailable() public returns (bool) {
        return this.gasTopUpLimitAvailable() <= initialLimit;
    }
}
