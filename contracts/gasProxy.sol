/**
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
 */

 pragma solidity ^0.5.17;


import "./internals/ownable.sol";
import "./internals/controllable.sol";
import "./internals/gasRefundable.sol";


contract GasProxy is Ownable, Controllable, GasRefundable {

	/// @dev Controllers are used by the application.	    
	event ExecutedTransaction(address _destination, uint256 _value, bytes _data, bytes _returnData);


    /// @param _ownerAddress Address of the contract owner.
	/// @param _controllerNode ENS node of the controller contract.
	/// @param _gasTokenNode ENS node of the gas token contract.
    constructor(address payable _ownerAddress, bytes32 _controllerNode, bytes32 _gasTokenNode) public GasRefundable(_gasTokenNode) {
		_initializeControllable(_controllerNode);
		_initializeOwnable(_ownerAddress, false);
    }

    /// @notice Sets the ENS registry address to a different one.
    function ensSetRegistry(address _ensRegistry) external onlyOwner {
        _ensSetRegistry(_ensRegistry);
    }

	/// @notice Sets the ENS node for the GST token.
    function gstSetNode(bytes32 _gstNode) external onlyOwner {
        _gasTokenSetNode(_gstNode);
    }

	/// @notice Executes a controller operation and refunds gas using gas tokens.
    /// @notice Only controllers can call this method when the controller isn't stopped.
    /// @param _destination Destination address of the executed transaction.
    /// @param _value Amount of ETH (wei) to be sent together with the transaction.
    /// @param _data Data payload of the controller transaction.
    /// @param _gasRefund Amount of gas to be refunded using gas tokens.
    function executeTransaction(address _destination, uint256 _value, bytes calldata _data, uint256 _gasRefund) external onlyController {
        if (_gasRefund != 0) {
            _freeGas(_gasRefund);
        }
        (bool success, bytes memory returnData) = _destination.call.value(_value)(_data);
        require(success, "external call failed");

        emit ExecutedTransaction(_destination, _value, _data, returnData);
    }
}
