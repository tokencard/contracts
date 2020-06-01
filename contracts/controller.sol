/**
 *  Controller - The Consumer Contract Wallet
 *  Copyright (C) 2019 The Contract Wallet Company Limited
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
 *
 *  SPDX-License-Identifier: GPL-3.0
 */

pragma solidity ^0.5.17;

import "./internals/ownable.sol";
import "./internals/gasRefundable.sol";


/// @title The IController interface provides access to the isController and isAdmin checks.
interface IController {
    function isController(address) external view returns (bool);

    function isAdmin(address) external view returns (bool);
}


/// @title Controller stores a list of controller addresses that can be used for authentication in other contracts.
/// @notice The Controller contract implements a hierarchy of concepts: owner, admin and signer.
/// @notice Owner can modify admins, admins can modify the signers, signers are used to execute transactions.
contract Controller is Ownable, GasRefundable {
    event ExecutedTransaction(address _destination, uint256 _value, bytes _data, bytes _returnData);

    event AddedSigner(address _sender, address _signer);
    event RemovedSigner(address _sender, address _signer);

    event AddedAdmin(address _sender, address _admin);
    event RemovedAdmin(address _sender, address _admin);

    event Stopped(address _sender);
    event Started(address _sender);

    mapping(address => bool) private _isAdmin;
    uint256 private _adminCount;

    mapping(address => bool) private _isSigner;
    uint256 private _signerCount;

    mapping(address => bool) private _isController;
    uint256 private _controllerCount;

    bool private _stopped;

    /// @dev The controller contract is added to the list of controllers to be backwards compatible with the v1 interface.
    /// @param _ownerAddress Address of the contract owner.
    constructor(address payable _ownerAddress, bytes32 _gasTokenNode) public Ownable(_ownerAddress, false) GasRefundable(_gasTokenNode) {
        _isController[address(this)] = true;
        _controllerCount++;
    }

    /// @notice Checks if message sender is an admin.
    modifier onlyAdmin() {
        require(_isAdmin[msg.sender], "sender is not an admin");
        _;
    }

    /// @notice Checks if message sender is a signer.
    modifier onlySigner() {
        require(_isSigner[msg.sender], "sender is not a signer");
        _;
    }

    /// @notice Checks if message sender is an owner or admin.
    modifier onlyAdminOrOwner() {
        require(_isOwner(msg.sender) || _isAdmin[msg.sender], "sender is not admin or owner");
        _;
    }

    /// @notice Checks if the controller is stopped.
    modifier notStopped() {
        require(!_stopped, "controller is stopped");
        _;
    }

    /// @notice Sets the ENS registry address to a different one.
    function ensSetRegistry(address _ensRegistry) external onlyOwner {
        _ensSetRegistry(_ensRegistry);
    }

    /// @notice Sets the ENS node for the GST token.
    function gstSetNode(bytes32 _gstNode) external onlyOwner {
        _gstSetNode(_gstNode);
    }

    /// @notice Adds a new admin to the list of admins.
    /// @param _account Address to add to the list of admins.
    function addAdmin(address _account) external onlyOwner {
        _addAdmin(_account);
    }

    /// @notice Remove a admin from the list of admins.
    /// @param _account Address to remove from the list of admins.
    function removeAdmin(address _account) external onlyOwner {
        _removeAdmin(_account);
    }

    /// @return The current number of admins.
    function adminCount() external view returns (uint256) {
        return _adminCount;
    }

    /// @notice Add a new signer to the list of signers.
    /// @param _account Address to be added to the list of signers.
    function addSigner(address _account) external onlyAdminOrOwner {
        _addSigner(_account);
    }

    /// @notice Remove a signer from the list of signers.
    /// @param _account Address to remove from the list of signers.
    function removeSigner(address _account) external onlyAdminOrOwner {
        _removeSigner(_account);
    }

    /// @return The current number of signers.
    function signerCount() external view returns (uint256) {
        return _signerCount;
    }

    /// @return The current number of controllers.
    function controllerCount() external view returns (uint256) {
        return _controllerCount;
    }

    /// @return Status of whether the provided account is an admin.
    function isAdmin(address _account) external view returns (bool) {
        return _isAdmin[_account];
    }

    /// @return Status of whether the provided account is a signer.
    function isSigner(address _account) external view returns (bool) {
        return _isSigner[_account];
    }

    /// @return Status of whether the provided account is a controller.
    function isController(address _account) external view returns (bool) {
        return _isController[_account];
    }

    /// @return Status of whether the controller contract is stopped.
    function isStopped() external view returns (bool) {
        return _stopped;
    }

    /// @dev Private method that adds a new admin.
    function _addAdmin(address _account) private {
        require(!_isAdmin[_account], "provided account is already an admin");
        require(!_isSigner[_account], "provided account is already a controller");
        require(!_isOwner(_account), "provided account is already the owner");
        require(_account != address(0), "provided account is the zero address");
        require(_account != address(this), "provided account is this contract");
        _isAdmin[_account] = true;
        _adminCount++;
        emit AddedAdmin(msg.sender, _account);
    }

    /// @dev Private method that removes an existing admin.
    function _removeAdmin(address _account) private {
        require(_isAdmin[_account], "provided account is not an admin");
        _isAdmin[_account] = false;
        _adminCount--;
        emit RemovedAdmin(msg.sender, _account);
    }

    /// @dev Private method that adds a new signer.
    function _addSigner(address _account) private {
        require(!_isAdmin[_account], "provided account is already an admin");
        require(!_isSigner[_account], "provided account is already a signer");
        require(!_isOwner(_account), "provided account is already the owner");
        require(_account != address(0), "provided account is the zero address");
        require(_account != address(this), "provided account is this contract");
        _isSigner[_account] = true;
        _signerCount++;
        emit AddedSigner(msg.sender, _account);
    }

    /// @dev Private method that removes an existing signer.
    function _removeSigner(address _account) private {
        require(_isSigner[_account], "provided account is not a signer");
        _isSigner[_account] = false;
        _signerCount--;
        emit RemovedSigner(msg.sender, _account);
    }

    /// @notice Stop the controller contract stopping the execution of transactions.
    function stop() external onlyAdminOrOwner {
        _stopped = true;
        emit Stopped(msg.sender);
    }

    /// @notice Start the controller if it was previously stopped.
    function start() external onlyOwner {
        _stopped = false;
        emit Started(msg.sender);
    }

    /// @notice Executes a controller operation and refunds gas using GST tokens.
    /// @notice Only signers can call this method when the controller isn't stopped.
    /// @param _destination Destination address of the executed transaction.
    /// @param _value Amount of ETH (wei) to be sent together with the transaction.
    /// @param _data Data payload of the controller transaction.
    /// @param _gasRefund Amount of gas to be refunded using GST tokens.
    function executeTransaction(address _destination, uint256 _value, bytes calldata _data, uint256 _gasRefund) external onlySigner notStopped {
        if (_gasRefund != 0) {
            _freeGas(_gasRefund);
        }
        (bool success, bytes memory returnData) = _destination.call.value(_value)(_data);
        require(success, "external call failed");

        emit ExecutedTransaction(_destination, _value, _data, returnData);
    }
}
