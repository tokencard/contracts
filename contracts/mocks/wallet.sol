// SPDX-License-Identifier: GPLv3

pragma solidity ^0.6.11;

import "../internals/controllable.sol";


contract Wallet is Controllable {

    event ConfirmedOperation(bool _status);

    /// @dev Ether can be deposited from any source, so this contract must be payable by anyone.
    receive() external payable {}

    constructor(address _ens_, bytes32 _controllerNode_) public {
        _initializeENSResolvable(_ens_);
        _initializeControllable(_controllerNode_);
    }

    function transfer(address payable _to, uint256 _amount) external {
        _to.transfer(_amount);
    }

    function sendValue(address payable _to, uint256 _amount) external {
        (bool success, ) = _to.call{value: _amount}("");
        require(success, "sendValue failed");
    }

    function confirmOperation() external onlyController {
        emit ConfirmedOperation(true);
    }
}
