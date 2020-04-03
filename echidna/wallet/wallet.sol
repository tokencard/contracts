pragma solidity ^0.5.15;

import "crytic-export/flattening/Wallet.sol";


contract Echidna {
    address internal echidnaDeployer = address(0x1);
    address payable internal echidnaOwner = address(0x2);
    address payable internal echidnaAttacker = address(0x3);
}


contract TEST is Echidna, Wallet {
    bytes32 private testNode = "test";

    constructor() public Wallet(echidnaOwner, false, address(0), bytes32(0), bytes32(0), bytes32(0), 10000) {}

    function echidna_licenceNodeStatic() public view returns (bool) {
        return licenceNode() == testNode;
    }
}
