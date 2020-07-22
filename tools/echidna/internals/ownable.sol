// SPDX-License-Identifier: GPLv3

pragma solidity ^0.6.11;

import "crytic-export/flattening/Ownable.sol";


contract Echidna {
    address payable internal echidnaDeployer = address(0x1);
    address payable internal echidnaOwner = address(0x2);
}


contract TEST is Echidna, Ownable {
    constructor() public {
        _initializeOwnable(echidnaOwner, false);
    }

    function echidna_nonTransferable() public view returns (bool) {
        return !isTransferable();
    }

    function echidna_nonTransferableOwner() public view returns (bool) {
        return _isOwner(echidnaOwner);
    }
}
