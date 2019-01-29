pragma solidity ^0.4.25;

import "./internals/ownable.sol";
import "./licence.sol";

/// @title The DAO interface (WIP stub).
interface IDAO {
    function lock() external;
    function isLocked() external returns (bool);
}

contract Dao is Ownable{

  bool private _locked;

  /// @dev ILicence is an interface to the licence contract that manages licence fees.
  ILicence public licence;

  /// @dev Constructor initializes the DAO contract.
  /// @param _owner is the owner account of the Dao contract.
  /// @param _transferable indicates whether the contract ownership can be transferred.
  /// @param _licence is the address of the licence contract.
  constructor(address _owner, bool _transferable, address _licence) Ownable(_owner, _transferable) public {
      licence = ILicence(_licence);
  }

  function lock() external onlyOwner{

    _locked = true;

  }

  function isLocked() external view returns (bool){

    return _locked;

  }

  function updateLicenceFee(uint _newFee) external onlyOwner {

    ILicence(licence).updateFee(_newFee);

  }

}
