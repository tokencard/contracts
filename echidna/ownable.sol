pragma solidity ^0.5.15;

import "contracts/licence.sol";

contract EchidnaInterface{
    address internal echidna_owner = address(0x41414141);
    address payable internal echidna_user = address(0x42424242);
    address internal echidna_attacker = address(0x43434343);
}

contract TEST is EchidnaInterface, Controller {

    constructor() Controller(echidna_user) public {}

    function echidna_nontransferable_owner() public returns (bool) {
        return owner() == echidna_user;
    }

}
