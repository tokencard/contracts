// SPDX-License-Identifier: GPLv3

pragma solidity 0.6.11;

import "crytic-export/flattening/Wallet.sol";


contract Echidna {
    address internal echidnaDeployer = address(0x1);
    address payable internal echidnaOwner = address(0x2);
    address payable internal echidnaAttacker = address(0x3);
}


contract TEST is Echidna, Wallet {
    bytes32 private testNode = "test";

    constructor() public {
        initializeWallet(echidnaOwner, false, address(0), bytes32(0), bytes32(0), bytes32(0), 10000);
    }

    function echidna_licenceNodeStatic() public view returns (bool) {
        return licenceNode() == testNode;
    }
}
