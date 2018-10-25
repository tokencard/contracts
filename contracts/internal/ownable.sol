pragma solidity ^0.4.25;


/// @title Ownable has an owner address and provides basic authorization control functions.
contract Ownable {
    event TransferredOwnership(address _to);

    address private _owner;
    bool private _isTransferable;


    /// @dev Constructor sets the original owner of the contract and the maximum number of ownership transfers.
    //todo mischa ... boolean only
    constructor(address _account, bool _transferable) internal {
        _owner = _account;
        _isTransferable = _transferable;
    }

    /// @dev Reverts if called by any account other than the owner.
    modifier onlyOwner() {
        require(isOwner(), "sender is not an owner");
        _;
    }

    /// @dev Allows the current owner to transfer control of the contract to a new address.
    /// @param _account address to transfer ownership to.
    function transferOwnership(address _account) external onlyOwner {
        // Require that the ownership is transferable.
        require(_isTransferable, "ownership is not transferable");
        // Require that the new owner is not the zero address.
        require(_account != address(0), "owner cannot be set to 0x0");
        // Set the transferable flag to false.
        _isTransferable = false;
        // Set the owner to the provided address.
        _owner = _account;
        // Emit the ownership transfer event.
        emit TransferredOwnership(_account);
    }

    /// @return the address of the owner.
    function owner() public view returns (address) {
        return _owner;
    }

    /// @return true if the ownership is transferable.
    function isTransferable() public view returns (bool) {
        return _isTransferable;
    }

    /// @return true if sender is the owner of the contract.
    function isOwner() internal view returns (bool) {
        return msg.sender == _owner;
    }
}
