pragma solidity ^0.5.15;

import "contracts/controller.sol";


contract Echidna {
    address payable internal echidna_deployer = address(0x1);
    address internal echidna_controller1 = address(0x2);
    address internal echidna_controller2 = address(0x3);
    address internal echidna_attacker = address(0x4);
}


contract TEST is Echidna, Controller {
    constructor() public Controller(echidna_deployer) {
        this.addController(echidna_controller1);
        this.addController(echidna_controller2);
    }

    function echidna_nonTransferableOwner() public view returns (bool) {
        return _isOwner(echidna_deployer);
    }

    function echidna_noAdminRole() public view returns (bool) {
        return this.controllerCount() == 2 && this.adminCount() == 0;
    }

    function echidna_ownerUniqueRole() public view returns (bool) {
        return _isOwner(echidna_deployer) && !this.isAdmin(echidna_deployer) && !this.isController(echidna_deployer);
    }

    function echidna_controller1UniqueRole() public view returns (bool) {
        return !_isOwner(echidna_controller1) && !this.isAdmin(echidna_controller1) && this.isController(echidna_controller1);
    }

    function echidna_controller2UniqueRole() public view returns (bool) {
        return !_isOwner(echidna_controller2) && !this.isAdmin(echidna_controller2) && this.isController(echidna_controller2);
    }

    function echidna_attackerNoRole() public view returns (bool) {
        return !_isOwner(echidna_attacker) && !this.isAdmin(echidna_attacker) && !this.isController(echidna_attacker);
    }
}
