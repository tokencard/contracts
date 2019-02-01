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
  event UpdatedCryptoFloat(address _sender, address _newFloat);
  event UpdatedTokenHolder(address _sender, address _newHolder);
  event UpdatedFee(address _sender, uint _newFee);

  event TransferredToTokenHolder(address _from, address _to, address _asset, uint _amount);
  event TransferredToCryptoFloat(address _from, address _to, address _asset, uint _amount);

  /// @dev This is 100% scaled up by a factor of 10 to give us 1 decimal place of precision
  uint constant public MAX_AMOUNT_SCALED = 1000;

  address constant private TKN = 0xaaaf91d9b90df800df4f55c205fd6989c977e73a; // solium-disable-line uppercase

  address private _cryptoFloat;
  address private _tokenHolder;

  bool private _lockedCryptoFloat;
  bool private _lockedTokenHolder;

  /// @dev This is the _licenceAmountScaled by a factor of 10
  /// i.e. 1% is 10 _licenceAmountScaled, 0.1% is 1 _licenceAmountScaled
  uint private _licenceAmountScaled;

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
      _licenceAmountScaled = _fee;
      _cryptoFloat = _float;
      _tokenHolder = _holder;
  }

  /// @dev Ether can be deposited from any source, so this contract should be payable by anyone.
  function() public payable {
      require(msg.data.length == 0);
      emit Received(msg.sender, msg.value);
  }

  /// @return the scaled licence amount, used to calculate the split when loading.
  function licenceAmountScaled() external view returns (uint) {
    return _licenceAmountScaled;
  }

  /// @return the address of the multi-sig cryptocurrency float contract.
  function cryptoFloat() external view returns (address) {
      return _cryptoFloat;
  }

  /// @return the address of the token holder contract.
  function tokenHolder() external view returns (address) {
      return _tokenHolder;
  }

  function lockFloat() external onlyOwner{
    _lockedCryptoFloat = true;
  }

  function lockHolder() external onlyOwner{
    _lockedTokenHolder = true;
  }

  function floatLocked() public view returns (bool){
    return _lockedCryptoFloat;
  }

  function holderLocked() public view returns (bool){
    return _lockedTokenHolder;
  }

  /// @dev Updates the address of the DAO contract.
  function updateDAO(address _newDAO) external onlyOwner {
      if (address(DAO) != address(0))
        require(!DAO.isLocked(), "DAO is locked");
      DAO = IDAO(_newDAO);
      emit UpdatedDAO(msg.sender, _newDAO);
  }

  /// @dev Updates the address of the cyptoFloat.
  function updateFloat(address _newFloat) external onlyOwner {
      require(!floatLocked(), "float is locked");
      _cryptoFloat = _newFloat;
      emit UpdatedCryptoFloat(msg.sender, _newFloat);
  }

  /// @dev Updates the address of the Holder contract.
  function updateHolder(address _newHolder) external onlyOwner {
    require(!holderLocked(), "holder contract is locked");
    _tokenHolder = _newHolder;
    emit UpdatedTokenHolder(msg.sender, _newHolder);
}

  /// @dev Updates the card licence fee.
  function updateFee(uint _newFee) external onlyDAO {
      require(1 <= _newFee && _newFee <= MAX_AMOUNT_SCALED, "fee out of range"); // TODO(daniel): same as below, not sure if using percentages inside solidity is the optimal solution.
      _licenceAmountScaled = _newFee;
      emit UpdatedFee(address(DAO), _newFee);
  }

  /// @dev Load the holder and float contracts based on the licence fee and asset amount.
  /// @param _asset is the address of an ERC20 token or 0x0 for ether.
  /// @param _amount is the amount of assets to be transferred including the fee.
  function load(address _asset, uint _amount) external payable {
      uint loadAmount = _amount;    
      // If TKN then no licence to be paid
      if (_asset == TKN) {
          require(ERC20(_asset).transferFrom(msg.sender, _cryptoFloat, _amount), "TKN transfer from external account was unsuccessful");
      } else {
          loadAmount = _amount.mul(MAX_AMOUNT_SCALED).div(_licenceAmountScaled + MAX_AMOUNT_SCALED);
          uint fee = _amount.sub(loadAmount);
    
          if (_asset != address(0)) {
              require(ERC20(_asset).transferFrom(msg.sender, _tokenHolder, fee), "ERC20 fee transfer from external account was unsuccessful");
              require(ERC20(_asset).transferFrom(msg.sender, _cryptoFloat, loadAmount), "ERC20 token transfer from external account was unsuccessful");
          } else {
              require(msg.value == _amount, "ether sent is not equal to amount");
              _tokenHolder.transfer(fee);
              _cryptoFloat.transfer(loadAmount);
          }

          emit TransferredToTokenHolder(msg.sender, _tokenHolder, _asset, fee);
      }
      emit TransferredToCryptoFloat(msg.sender, _cryptoFloat, _asset, loadAmount);
  }
}
