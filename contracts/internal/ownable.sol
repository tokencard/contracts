pragma solidity ^0.4.24;


/// @title Ownable has an owner address and provides basic authorization control functions.
contract Ownable {
    event OwnershipTransferred(address _to);
    event OwnershipAccepted(address _from);

    address private _owner;
    address private _newOwner;

    /// @dev Constructor sets the original owner of the contract to the provided account.
    constructor(address _account) internal {
        _owner = _account;
    }

    /// @dev Reverts if called by any account other than the owner.
    modifier onlyOwner() {
        require(isOwner(), "sender is not an owner");
        _;
    }

    /// @dev Reverts if called by any account other than the new owner.
    modifier onlyNewOwner() {
        require(isNewOwner(), "sender is not the new owner");
        _;
    }

    /// @dev Allows the current owner to transfer control of the contract to a new address.
    /// @notice The ownership transfer has to be accepted by the new proposed owner.
    /// @param _account address to transfer ownership to.
    function transferOwnership(address _account) external onlyOwner {
        // Set the new owner to the provided address.
        _newOwner = _account;
        // Emit the ownership transfer event.
        emit OwnershipTransferred(_newOwner);
    }

    /// @dev Allows the new owner to accept ownership of the contract.
    function acceptOwnership() external onlyNewOwner {
        // Require that the new owner is not the zero address.
        require(_newOwner != address(0), "owner cannot be set to 0x0000000000000000000000000000000000000000");
        // Emit the acceptance event with the previous owner address.
        emit OwnershipAccepted(_owner);
        // Change the owner to the new address.
        _owner = _newOwner;
        // Reset the new owner address variable.
        delete _newOwner;
    }

    /// @return the address of the owner.
    function owner() public view returns (address) {
        return _owner;
    }

    /// @return the address of the new owner.
    function newOwner() public view returns (address) {
        return _newOwner;
    }

    /// @return true if sender is the owner of the contract.
    function isOwner() public view returns (bool) {
        return msg.sender == _owner;
    }

    /// @return true if sender is the new owner of the contract.
    function isNewOwner() public view returns (bool) {
        return msg.sender == _newOwner;
    }
}
