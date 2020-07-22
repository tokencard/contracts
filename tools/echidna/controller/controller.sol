// SPDX-License-Identifier: GPLv3

pragma solidity ^0.6.11;

import "crytic-export/flattening/Controller.sol";


contract Echidna {
    address payable internal echidnaDeployer = address(0x1);
    address internal echidnaController1 = address(0x2);
    address internal echidnaController2 = address(0x3);
    address internal echidnaAttacker = address(0x4);
}


contract TEST is Echidna, Controller {
    constructor() public Controller(echidnaDeployer) {
        addController(echidnaController1);
        addController(echidnaController2);
    }

    function echidna_nonTransferableOwner() public view returns (bool) {
        return _isOwner(echidnaDeployer);
    }

    function echidna_noAdminRole() public view returns (bool) {
        return controllerCount() == 2 && adminCount() == 0;
    }

    function echidna_ownerUniqueRole() public view returns (bool) {
        return _isOwner(echidnaDeployer) && !isAdmin(echidnaDeployer) && !isController(echidnaDeployer);
    }

    function echidna_controller1UniqueRole() public view returns (bool) {
        return !_isOwner(echidnaController1) && !isAdmin(echidnaController1) && isController(echidnaController1);
    }

    function echidna_controller2UniqueRole() public view returns (bool) {
        return !_isOwner(echidnaController2) && !isAdmin(echidnaController2) && isController(echidnaController2);
    }

    function echidna_attackerNoRole() public view returns (bool) {
        return !_isOwner(echidnaAttacker) && !isAdmin(echidnaAttacker) && !isController(echidnaAttacker);
    }
}
