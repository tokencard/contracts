pragma solidity ^0.4.25;

import "./controllable.sol";


/// @title OraclizeResolver implements on-chain contract address resolution.
contract OraclizeResolver is Controllable {

    address private _address;

    /// @dev Constructor initializes the resolver target and sets up controller authentication.
    /// @param _target initial address of the target contract.
    /// @param _controller is the controller contract address.
    constructor(address _target, address _ens, byte32 _controller) Controllable(_ens, _controller) public {
        _address = _target;
    }

    /// @return the address of the target contract.
    /// @notice this function is compatible with the oraclize contract API.
    function getAddress() public view returns (address) {
        return _address;
    }

    /// @dev Sets the address of the target contract.
    /// @param _target new address of the target contract.
    function setAddress(address _target) public onlyController {
        _address = _target;
    }
}
