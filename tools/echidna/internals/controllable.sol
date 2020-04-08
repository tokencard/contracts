pragma solidity ^0.5.15;

import "crytic-export/flattening/Controllable.sol";


contract Echidna {
    address payable internal echidnaDeployer = address(0x1);
    address payable internal echidnaOwner = address(0x2);
}


contract TEST is Echidna, Controllable {
    bytes32 constant _DEFAULT_CONTROLLER_NODE = 0x7f2ce995617d2816b426c5c8698c5ec2952f7a34bb10f38326f74933d5893697;

    constructor() public Controllable(bytes32(0)) {}

    function echidna_controllerNode() public view returns (bool) {
        return controllerNode() == _DEFAULT_CONTROLLER_NODE;
    }
}
