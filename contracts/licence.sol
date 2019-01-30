pragma solidity ^0.4.25;

import "./externals/SafeMath.sol";
import "./externals/ERC20.sol";
import "./internals/ownable.sol";
import "./dao.sol";

/// @title ILicence interface describes methods for loading a TokenCard inclusive of licence fees.
interface ILicence {
    function load(address, uint) external payable returns (bool);
    function updateFee(uint) external;
}

/// @title Licence loads the TokenCard and transfers the licence fee to the token holder contract.
contract Licence is Ownable {

  using SafeMath for uint256;

  /*******************/
  /*     Events     */
  /*****************/

  event Received(address _from, uint _amount);

  event UpdatedDAO(address _sender, address _newDAO);
  event UpdatedFee(address _sender, uint _newFee);

  event TransferredToTokenHolder(address _from, address _to, address _asset, uint _amount);
  event TransferredToCryptoFloat(address _from, address _to, address _asset, uint _amount);

  uint constant public MAX_FEE_FACTOR = 1000;

  address private _cryptoFloat;
  address private _tokenHolder;

  uint private _feeFactor;

  /// @dev IDAO is an interface to a DAO contract that can update the licence contract.
  IDAO public DAO;

  /// @dev Reverts if called by any address other than the DAO contract.
  modifier onlyDAO() {
      require(msg.sender == address(DAO));
      _;
  }

  /// @dev Constructor initializes the card licence contract.
  /// @param _owner is the owner account of the wallet contract.
  /// @param _transferable indicates whether the contract ownership can be transferred.
  /// @param _fee is the initial card licence fee.
  /// @param _float is the address of the multi-sig cryptocurrency float contract.
  /// @param _holder is the address of the token holder contract
  constructor(address _owner, bool _transferable, uint _fee, address _float, address _holder) Ownable(_owner, _transferable) public {
      _feeFactor = _fee;
      _cryptoFloat = _float;
      _tokenHolder = _holder;
  }

  /// @dev Ether can be deposited from any source, so this contract should be payable by anyone.
  function() public payable {
      require(msg.data.length == 0);
      emit Received(msg.sender, msg.value);
  }

  /// @return the address of the multi-sig cryptocurrency float contract.
  function cryptoFloat() external view returns (address) {
      return _cryptoFloat;
  }

  /// @return the address of the token holder contract.
  function tokenHolder() external view returns (address) {
      return _tokenHolder;
  }

  /// @return the fee factor used to calculate how to split the amount sent to load() function.
  function feeFactor() external view returns (uint) {
    return _feeFactor;
  }

  /// @dev Updates the address of the DAO contract.
  function updateDAO(address _newDAO) external onlyOwner {
      if (address(DAO) != address(0))
        require(!DAO.isLocked(), "DAO is locked");
      DAO = IDAO(_newDAO);
      emit UpdatedDAO(msg.sender, _newDAO);
  }

  /// @dev Updates the card licence fee.
  function updateFee(uint _newFee) external onlyDAO {
      require(1 <= _newFee && _newFee <= MAX_FEE_FACTOR, "fee out of range"); // TODO(daniel): same as below, not sure if using percentages inside solidity is the optimal solution.
      _feeFactor = _newFee;
      emit UpdatedFee(address(DAO), _newFee);
  }

  /// @dev Load the holder and float contracts based on the licence fee and asset amount.
  /// @param _asset is the address of an ERC20 token or 0x0 for ether.
  /// @param _amount is the amount of assets to be transferred including the fee.
  function load(address _asset, uint _amount) external payable {

      uint transferAmount = _amount.mul(MAX_FEE_FACTOR).div(_feeFactor + MAX_FEE_FACTOR);
      uint fee = _amount.sub(transferAmount); //transferAmount is always <= amount

    //  assert(_amount == transferAmount + fee);
      if (_asset != address(0)) {
          require(ERC20(_asset).transferFrom(msg.sender, _tokenHolder, fee), "ERC20 fee transfer from external account was unsuccessful");
          require(ERC20(_asset).transferFrom(msg.sender, _cryptoFloat, transferAmount), "ERC20 token transfer from external account was unsuccessful");
      } else {
          require(msg.value == _amount, "ether sent is not equal to amount");
          _tokenHolder.transfer(fee);
          _cryptoFloat.transfer(transferAmount);
      }
      emit TransferredToTokenHolder(msg.sender, _tokenHolder, _asset, fee);
      emit TransferredToCryptoFloat(msg.sender, _cryptoFloat, _asset, transferAmount);

  }
}
