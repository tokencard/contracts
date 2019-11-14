pragma solidity ^0.5.10;

import "contracts/controller.sol";

contract CryticInterface{
    address payable internal crytic_owner = address(0x41414141);
    address internal crytic_controller_1 = address(0x424242421);
    address internal crytic_controller_2 = address(0x424242422);
    address internal crytic_attacker = address(0x43434343);
}

contract TEST is CryticInterface, Controller {

    constructor() Controller(crytic_owner)  public {}

    function initializeControllers() internal returns (bool) {
        this.addController(crytic_controller_1);
        this.addController(crytic_controller_2);
    }

    function isOwner(address x) internal view returns (bool) {
        return owner() == x;
    }

    function crytic_nontransferable_owner() public view returns (bool) {
        return isOwner(crytic_owner);
    }

    function crytic_no_admin_controller() public view returns (bool) {
        return this.controllerCount() == 2 && this.adminCount() == 0;
    }

    function crytic_owner_unique_rol() public view returns (bool) {
        return isOwner(crytic_owner) && !isAdmin(crytic_owner) && !isController(crytic_owner);
    }

    function crytic_controller_1_unique_rol() public view returns (bool) {
        return !isOwner(crytic_controller_1) && !isAdmin(crytic_controller_1) && isController(crytic_controller_1);
    }

    function crytic_controller_2_unique_rol() public view returns (bool) {
        return !isOwner(crytic_controller_2) && !isAdmin(crytic_controller_2) && isController(crytic_controller_2);
    }

    function crytic_attacker_no_rol() public view returns (bool) {
        return !isOwner(crytic_attacker) && !isAdmin(crytic_attacker) && !isController(crytic_attacker);
    }

}
