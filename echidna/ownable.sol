pragma solidity ^0.5.10;

import "contracts/licence.sol";

contract CryticInterface{
    address internal crytic_owner = address(0x41414141);
    address internal crytic_user = address(0x42424242);
    address internal crytic_attacker = address(0x43434343);
}

contract TEST is CryticInterface, Licence {

    constructor() Licence(crytic_user, false, 10, address(0x0), address(0x0), address(0x0))  public {
    }

    function crytic_nontransferable_owner() public returns (bool) {
        return owner() == crytic_user;
    }

}
