/**
 *  The Consumer Contract Wallet
 *  Copyright (C) 2018 The Contract Wallet Company Limited
 *
 *  This program is free software: you can redistribute it and/or modify
 *  it under the terms of the GNU General Public License as published by
 *  the Free Software Foundation, either version 3 of the License, or
 *  (at your option) any later version.

 *  This program is distributed in the hope that it will be useful,
 *  but WITHOUT ANY WARRANTY; without even the implied warranty of
 *  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *  GNU General Public License for more details.

 *  You should have received a copy of the GNU General Public License
 *  along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

pragma solidity ^0.4.25;

import "./ownable.sol";

/// @title The Controller interface provides access to an external list of controllers.
interface IController {
    function isController(address) external view returns (bool);
}

/// @title Controller stores a list of controller addresses that can be used for authentication in other contracts.
contract Controller is IController, Ownable{
    event AddedController(address _sender, address _controller);
    event RemovedController(address _sender, address _controller);

    event AddedAdmin(address _sender, address _admin);
    event RemovedAdmin(address _sender, address _admin);

    mapping (address => bool) private _isAdmin;
    uint private _adminCount;

    mapping (address => bool) private _isController;
    uint private _controllerCount;

    /// @dev Constructor initializes the owner with the provided address.
    /// @param _ownerAddress address of the owner.
    /// @param _transferable indicates whether the contract ownership can be transferred.
    constructor(address _ownerAddress, bool _transferable) Ownable(_ownerAddress, _transferable) public {
    }

    /// @dev Checks if message sender is an admin.
    modifier onlyAdmin() {
        require(isAdmin(msg.sender), "sender is not an admin");
        _;
    }

    /// @dev Add a new admin to the list of admins.
    /// @param _account address to add to the list of admins.
    function addAdmin(address _account) external onlyOwner {
        _addAdmin(_account);
    }

    /// @dev Remove a admin from the list of admins.
    /// @param _account address to remove from the list of admins.
    function removeAdmin(address _account) external onlyOwner {
        _removeAdmin(_account);
    }

    /// @return true if the provided account is an admin.
    function isAdmin(address _account) public view returns (bool) {
        return _isAdmin[_account];
    }

    /// @return the current number of admins.
    function adminCount() public view returns (uint) {
        return _adminCount;
    }

    /// @dev Internal-only function that adds a new admin.
    function _addAdmin(address _account) internal {
        require(!_isAdmin[_account], "provided account is already an admin");
        require(!_isController[_account], "provided account is already a controller");
        require(_account != owner(), "provided account is already the owner");
        _isAdmin[_account] = true;
        _adminCount++;
        emit AddedAdmin(msg.sender, _account);
    }

    /// @dev Internal-only function that removes an existing admin.
    function _removeAdmin(address _account) internal {
        require(_isAdmin[_account], "provided account is not an admin");
        _isAdmin[_account] = false;
        _adminCount--;
        emit RemovedAdmin(msg.sender, _account);
    }

    /// @dev Checks if message sender is a controller.
    modifier onlyController() {
        require(isController(msg.sender), "sender is not a controller");
        _;
    }

    /// @dev Add a new controller to the list of controllers.
    /// @param _account address to add to the list of controllers.
    function addController(address _account) external onlyAdmin {
        _addController(_account);
    }

    /// @dev Remove a controller from the list of controllers.
    /// @param _account address to remove from the list of controllers.
    function removeController(address _account) external onlyAdmin {
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
        require(!_isAdmin[_account], "provided account is already an admin");
        require(_account != owner(), "provided account is already the owner");
        _isController[_account] = true;
        _controllerCount++;
        emit AddedController(msg.sender, _account);
    }

    /// @dev Internal-only function that removes an existing controller.
    function _removeController(address _account) internal {
        require(_isController[_account], "provided account is not a controller");
        _isController[_account] = false;
        _controllerCount--;
        emit RemovedController(msg.sender, _account);
    }
}
