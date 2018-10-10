pragma solidity ^0.4.25;

/// @title The Controller interface provides access to an external list of controllers.
interface IController {
    function isController(address) external view returns (bool);
}

/// @title Controller stores a list of controller addresses that can be used for authentication in other contracts.
contract Controller is IController {
    event ControllerAdded(address _sender, address _controller);
    event ControllerRemoved(address _sender, address _controller);

    mapping (address => bool) private _isController;
    uint private _controllerCount;

    /// @dev Constructor initializes the list of controllers with the provided address.
    /// @param _account address to add to the list of controllers.
    constructor(address _account) public {
        _addController(_account);
    }

    /// @dev Checks if message sender is a controller.
    modifier onlyController() {
        require(isController(msg.sender), "sender is not a controller");
        _;
    }

    /// @dev Add a new controller to the list of controllers.
    /// @param _account address to add to the list of controllers.
    function addController(address _account) external onlyController {
        _addController(_account);
    }

    /// @dev Remove a controller from the list of controllers.
    /// @param _account address to remove from the list of controllers.
    function removeController(address _account) external onlyController {
        _removeController(_account);
    }

    /// @return true if the provided account is a controller.
    function isController(address _account) public view returns (bool) {
        return _isController[_account];
    }

    /// @return the current number of controllers.
    function controllerCount() public view returns (uint) {
        return _controllerCount;
    }

    /// @dev Internal-only function that adds a new controller.
    function _addController(address _account) internal {
        require(!_isController[_account], "provided account is already a controller");
        _isController[_account] = true;
        _controllerCount++;
        emit ControllerAdded(msg.sender, _account);
    }

    /// @dev Internal-only function that removes an existing controller.
    function _removeController(address _account) internal {
        require(_isController[_account], "provided account is not a controller");
        require(_controllerCount > 1, "cannot remove the last controller");
        _isController[_account] = false;
        _controllerCount--;
        emit ControllerRemoved(msg.sender, _account);
    }
}
