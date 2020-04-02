pragma solidity ^0.5.15;

import "contracts/internals/ownable.sol";


contract Echidna {
    address payable internal echidna_deployer = address(0x1);
    address payable internal echidna_owner = address(0x2);
}


contract TEST is Echidna, Ownable {
    constructor() public Ownable(echidna_owner, false) {}

    function echidna_nonTransferableOwner() public view returns (bool) {
        return _isOwner(echidna_owner);
    }
}
