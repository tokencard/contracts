/**
 *  The Consumer Contract Wallet
 *  Copyright (C) 2018 Token Group Ltd
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


/// @title Ownable has an owner address and provides basic authorization control functions.
/// This contract is modified version of the MIT OpenZepplin Ownable contract 
/// This contract doesn't allow for multiple changeOwner operations
/// https://github.com/OpenZeppelin/openzeppelin-solidity/blob/master/contracts/ownership/Ownable.sol
contract Ownable {
    event TransferredOwnership(address _from, address _to);

    address private _owner;
    bool private _isTransferable;

    /// @dev Constructor sets the original owner of the contract and whether or not it is one time transferable.
    constructor(address _account, bool _transferable) internal {
        _owner = _account;
        _isTransferable = _transferable;
        emit TransferredOwnership(address(0), _account);
    }

    /// @dev Reverts if called by any account other than the owner.
    modifier onlyOwner() {
        require(_isOwner(), "sender is not an owner");
        _;
    }

    /// @dev Allows the current owner to transfer control of the contract to a new address.
    /// @param _account address to transfer ownership to.
    /// @param _transferable indicates whether to keep the ownership transferable.
    function transferOwnership(address _account, bool _transferable) external onlyOwner {
        // Require that the ownership is transferable.
        require(_isTransferable, "ownership is not transferable");
        // Require that the new owner is not the zero address.
        require(_account != address(0), "owner cannot be set to zero address");
        // Set the transferable flag.
        _isTransferable = _transferable;
        // Emit the ownership transfer event.
        emit TransferredOwnership(_owner, _account);
        // Set the owner to the provided address.
        _owner = _account;
    }

    /// @dev Allows the current owner to relinquish control of the contract. 
    /// @notice Renouncing to ownership will leave the contract without an owner and unusable.
    /// It will not be possible to call the functions with the `onlyOwner` modifier anymore.
    function renounceOwnership() public onlyOwner {
        // Require that the ownership is transferable.
        require(_isTransferable, "ownership is not transferable");
        emit TransferredOwnership(_owner, address(0));
        // note that this could be terminal
        _owner = address(0);
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
    function _isOwner() internal view returns (bool) {
        return msg.sender == _owner;
    }
}
