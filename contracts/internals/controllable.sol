pragma solidity ^0.4.25;

import "./controller.sol";
import "../externals/ens/ENS.sol";

/// @title Resolver returns the controller contract address.
interface IResolver {
    function addr(bytes32) external view returns (address);
}

/// @title Controllable implements access control functionality based on a controller set in ENS.
contract Controllable {
    // @dev _ENS points to the ENS registry smart contract.
    ENS private _ENS;
    // @dev Is the registered ENS name of the controller contract.
    bytes32 private _node;

    // @dev Constructor initializes the controller contract object.
    // @param _ens is the address of the ENS.
    // @param _controllerName is the ENS name of the Controller.
    constructor(address _ens, bytes32 _controllerName) internal {
      _ENS = ENS(_ens);
      _node = _controllerName;
    }

    // @dev Checks if message sender is the controller.
    modifier onlyController() {
        require(isController(msg.sender), "sender is not a controller");
        _;
    }

    // @return true if the provided account is the controller.
    function isController(address _account) internal view returns (bool) {
        return IController(IResolver(_ENS.resolver(_node)).addr(_node)).isController(_account);
    }
}
