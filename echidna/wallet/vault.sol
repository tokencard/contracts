pragma solidity ^0.5.15;

import "crytic-export/flattening/Wallet.sol";


contract Echidna {
    address internal echidna_deployer = address(0x1);
    address payable internal echidna_owner = address(0x2);
    address payable internal echidna_attacker = address(0x3);
}


contract TEST is Echidna, Vault {
    constructor() public {}
}
