pragma solidity ^0.5.10;

import "crytic-export/flattening/Controller.sol";

contract EchidnaInterface{
    address payable internal echidna_owner = address(0x41414141);
    address internal echidna_controller_1 = address(0x424242421);
    address internal echidna_controller_2 = address(0x424242422);
    address internal echidna_attacker = address(0x43434343);
}

contract TEST is EchidnaInterface, Controller {

    constructor() Controller(echidna_owner)  public {
        addController(echidna_controller_1);
        addController(echidna_controller_2);
    }

    function isOwner(address x) internal view returns (bool) {
        return owner() == x;
    }

    function echidna_nontransferable_owner() public view returns (bool) {
        return isOwner(echidna_owner);
    }

    function echidna_no_admin_controller() public view returns (bool) {
        return this.controllerCount() == 2 && this.adminCount() == 0;
    }

    function echidna_owner_unique_rol() public view returns (bool) {
        return isOwner(echidna_owner) && !isAdmin(echidna_owner) && !isController(echidna_owner);
    }

    function echidna_controller_1_unique_rol() public view returns (bool) {
        return !isOwner(echidna_controller_1) && !isAdmin(echidna_controller_1) && isController(echidna_controller_1);
    }

    function echidna_controller_2_unique_rol() public view returns (bool) {
        return !isOwner(echidna_controller_2) && !isAdmin(echidna_controller_2) && isController(echidna_controller_2);
    }

    function echidna_attacker_no_rol() public view returns (bool) {
        return !isOwner(echidna_attacker) && !isAdmin(echidna_attacker) && !isController(echidna_attacker);
    }

}
