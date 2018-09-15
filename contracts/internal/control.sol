pragma solidity ^0.4.24;


/// @title Control handles oracle access control.
contract Control {
    event AddController(address _sender, address _account);
    event RemoveController(address _sender, address _account);

    mapping (address => bool) public isController;
    uint public controllerCount;
    address public owner;

    /// @dev Executable only by the owner.
    modifier onlyOwner() {
        require(msg.sender == owner, "sender is not an owner");
        _;
    }

    /// @dev Executable only by the controller.
    modifier onlyController() {
        require(isController[msg.sender], "sender is not a controller");
        _;
    }

    /// @dev Executable by either owner or controller.
    modifier eitherOwnerOrController() {
        require(isController[msg.sender] || msg.sender == owner, "sender is neither a controller nor an owner");
        _;
    }

    /// @dev Add a new controller to the list of controllers.
    function addController(address _account) external onlyController {
        _addController(_account);
    }

    /// @dev Remove a controller from the list of controllers.
    function removeController(address _account) external onlyController {
        _removeController(_account);
    }

    /// @dev Re-usable internal function that adds a new controller.
    function _addController(address _account) internal {
        require(!isController[_account], "provided account is already a controller");
        isController[_account] = true;
        controllerCount++;
        emit AddController(msg.sender, _account);
    }

    /// @dev Re-usable internal function that removes an existing controller.
    function _removeController(address _account) internal {
        require(isController[_account], "provided account is not a controller");
        require(controllerCount > 1, "last controller account cannot be removed");
        isController[_account] = false;
        controllerCount--;
        emit RemoveController(msg.sender, _account);
    }
}
