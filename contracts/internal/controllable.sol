pragma solidity ^0.4.24;


/// @title The Controller interface provides access to an external list of controllers.
interface Controller {
    function isController(address) public view returns (bool);
}


/// @title Controllable implements access control functionality based on an external list of controllers.
contract Controllable {
    /// @dev Controller points to the contract that contains the list of controllers.
    Controller private _c;

    /// @dev Constructor initializes the controller contract object.
    /// @param _controller address of the controller contract.
    constructor(address _controller) internal {
        _c = Controller(_controller);
    }

    /// @dev Checks if message sender is a controller.
    modifier onlyController() {
        require(isController(msg.sender), "sender is not a controller");
        _;
    }

    /// @return true if the provided account is a controller.
    function isController(address _account) internal view returns (bool) {
        return _c.isController(_account);
    }
}
