pragma solidity ^0.4.24;

import "controllable.sol";


/// @title Resolver implements on-chain contract address resolution.
contract Resolver is Controllable {

    address private _target;

    /// @dev Constructor initializes the resolver target and sets up controller authentication.
    /// @param _address initial address of the target contract.
    /// @param _controller address of the controller contract.
    constructor(address _address, address _controller) Controllable(_controller) public {
        _target = _address;
    }

    /// @return the address of the target contract.
    /// @notice this function is compatible with the oraclize contract API.
    function getAddress() public returns (address) {
        return _target;
    }

    /// @dev Sets the address of the target contract.
    /// @param _address new address of the target contract.
    function setAddress(address _address) public onlyController {
        _target = _address;
    }
}
