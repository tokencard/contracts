pragma solidity ^0.5.15;

import "crytic-export/flattening/Wallet.sol";


contract Echidna {
    address internal echidna_deployer = address(0x1);
    address payable internal echidna_owner = address(0x2);
    address payable internal echidna_attacker = address(0x3);
}


contract TEST is Echidna, Wallet {
    bytes32 private testNode = "test";

    constructor() public Wallet(echidna_owner, false, address(0), bytes32(0), bytes32(0), bytes32(0), 10000) {}

    function echidna_licenceNodeStatic() public view returns (bool) {
        return licenceNode() == testNode;
    }
}
