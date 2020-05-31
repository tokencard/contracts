/**
 *  Ownable - The Consumer Contract Wallet
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


/// @title Ownable has an owner address and provides basic authorization control functions.
/// @notice This contract allows for the transferOwnership operation to be made impossible.
/// @dev This contract is a modified version of the MIT OpenZepplin Ownable contract
/// @dev https://github.com/OpenZeppelin/openzeppelin-solidity/blob/master/contracts/ownership/Ownable.sol
contract Ownable {
    /// @notice Emits from and to addresses when ownership is transferred.
    event TransferredOwnership(address _from, address _to, bool _transferable);

    address payable private _owner;
    bool private _isTransferable;

    /// @notice Constructor sets the original owner of the contract and whether or not it is one time transferable.
    /// @param _account Address to transfer ownership to.
    /// @param _transferable Whether to keep the ownership transferable.
    constructor(address payable _account, bool _transferable) internal {
        require(_account != address(0), "owner cannot be set to zero address");
        require(_account != address(this), "owner cannot be set to this contract");
        _owner = _account;
        _isTransferable = _transferable;
        emit TransferredOwnership(address(0), _account, _transferable);
    }

    /// @notice Reverts if called by any account other than the owner.
    modifier onlyOwner() {
        require(_isOwner(msg.sender), "sender is not an owner");
        _;
    }

    /// @notice Allows the current owner to transfer control of the contract to a new address.
    /// @param _account Address to transfer ownership to.
    /// @param _transferable Whether to keep the ownership transferable.
    function transferOwnership(address payable _account, bool _transferable) external onlyOwner {
        require(_isTransferable, "ownership is not transferable");
        require(_account != address(0), "owner cannot be set to zero address");
        require(_account != address(this), "owner cannot be set to this contract");
        _isTransferable = _transferable;
        emit TransferredOwnership(_owner, _account, _transferable);
        _owner = _account;
    }

    /// @notice Checks if the ownership is transferable.
    /// @return Ownership transferable status.
    function isTransferable() external view returns (bool) {
        return _isTransferable;
    }

    /// @notice Allows the current owner to relinquish control of the contract.
    /// @notice Renouncing to ownership will leave the contract without an owner and unusable.
    /// @notice It will not be possible to call the functions with the `onlyOwner` modifier anymore.
    function renounceOwnership() external onlyOwner {
        require(_isTransferable, "ownership is not transferable");
        _isTransferable = false;
        emit TransferredOwnership(_owner, address(0), false);
        _owner = address(0);
    }

    /// @return Address of the owner.
    function owner() public view returns (address payable) {
        return _owner;
    }

    /// @return Status of whether the provided address is an owner.
    function _isOwner(address _address) internal view returns (bool) {
        return _address == _owner;
    }
}
