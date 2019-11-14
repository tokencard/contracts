pragma solidity ^0.5.10;

import "contracts/wallet.sol";

contract CryticInterface{
    address internal crytic_owner = address(0x41414141);
    address internal crytic_user = address(0x42424242);
    address internal crytic_attacker = address(0x43434343);

    address internal hevm_cheatcode = address(0x7109709ecfa91a80626ff3989d68f67f5b1dd12d);
}

contract TEST is CryticInterface, GasTopUpLimit {

    uint initialLimit = 500 finney;
    constructor() GasTopUpLimit() Ownable(crytic_user, true) Controllable(bytes32(0x0)) ENSResolvable(address(0x0)) public {
    }

    // This function is used to simulate the user spending, so anyone can call it
    function spend(uint _amount) public {
        _enforceLimit(_gasTopUpLimit,_amount);
    }

    // uses the hevm `warp` cheatcode to increment the block.timestamp by 24 hours + 1
    // see https://github.com/dapphub/dapptools/blob/79744a960979f1a593d8e04b04a4df9b6c533fa9/src/hevm/src/EVM.hs#L1333-L1338
    // for more info
    function increaseDayNow() public {
        hevm_cheatcode.call(bytes4(keccak256("warp(uint256)")), now + (1 days) + 1);
    }

    function crytic_daily_GasTopUpLimitAvailable() public returns (bool) {
        increaseDayNow();
        return this.gasTopUpLimitAvailable() == initialLimit;
    }

    function crytic_constant_GasTopUpLimitAvailable() public returns (bool) {
        return this.gasTopUpLimitAvailable() <= initialLimit;
    }

}
