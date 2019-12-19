pragma solidity ^0.5.15;

import "contracts/wallet.sol";

contract EchidnaInterface{
    address internal echidna_owner = address(0x41414141);
    address payable internal echidna_user = address(0x42424242);
    address internal echidna_attacker = address(0x43434343);

    address internal hevm_cheatcode = address(0x7109709ECfa91a80626fF3989D68f67F5b1DD12D);
}

contract TEST is EchidnaInterface, GasTopUpLimit {

    uint initialLimit = 500 finney;
    constructor() GasTopUpLimit() Ownable(echidna_user, true) Controllable(bytes32(0x0)) ENSResolvable(address(0x0)) public {}

    // This function is used to simulate the user spending, so anyone can call it
    function spend(uint _amount) public {
        _gasTopUpLimit._enforceLimit(_amount);
    }

    // uses the hevm `warp` cheatcode to increment the block.timestamp by 24 hours + 1
    // see https://github.com/dapphub/dapptools/blob/79744a960979f1a593d8e04b04a4df9b6c533fa9/src/hevm/src/EVM.hs#L1333-L1338
    // for more info
    function increaseDayNow() public {
        hevm_cheatcode.call(abi.encodeWithSignature("warp(uint256)", now + 2 days));
    }

    // function echidna_daily_GasTopUpLimitAvailable() public returns (bool) {
    //     increaseDayNow();
    //     return this.gasTopUpLimitAvailable() == initialLimit;
    // }

    function echidna_constant_GasTopUpLimitAvailable() public returns (bool) {
        return this.gasTopUpLimitAvailable() <= initialLimit;
    }

}
