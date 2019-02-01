pragma solidity ^0.4.25;

import "./externals/SafeMath.sol";
import "./externals/ERC20.sol";
import "./internals/ownable.sol";

/// @title ILicence interface describes methods for loading a TokenCard and updating licence amount.
interface ILicence {
    function load(address, uint) external payable returns (bool);
    function updateLicenceAmount(uint) external;
}

/// @title Licence loads the TokenCard and transfers the licence amout to the token holder contract.
contract Licence is Ownable {

    using SafeMath for uint256;

    /*******************/
    /*     Events     */
    /*****************/

    event Received(address _from, uint _amount);

    event UpdatedLicenceDAO(address _sender, address _newDAO);
    event UpdatedCryptoFloat(address _sender, address _newFloat);
    event UpdatedTokenHolder(address _sender, address _newHolder);
    event UpdatedLicenceAmount(address _sender, uint _newAmount);

    event TransferredToTokenHolder(address _from, address _to, address _asset, uint _amount);
    event TransferredToCryptoFloat(address _from, address _to, address _asset, uint _amount);

    /// @dev This is 100% scaled up by a factor of 10 to give us an extra 1 decimal place of precision
    uint constant public MAX_AMOUNT_SCALE = 1000;

    address constant private TKN = 0xaAAf91D9b90dF800Df4F55c205fd6989c977E73a; // solium-disable-line uppercase

    address private _cryptoFloat;
    address private _tokenHolder;
    address private _licenceDAO;

    bool private _lockedCryptoFloat;
    bool private _lockedTokenHolder;
    bool private _lockedLicenceDAO;

    /// @dev This is the _licenceAmountScaled by a factor of 10
    /// i.e. 1% is 10 _licenceAmountScaled, 0.1% is 1 _licenceAmountScaled
    uint private _licenceAmountScaled;

    /// @dev Reverts if called by any address other than the DAO contract.
    modifier onlyDAO() {
        require(msg.sender == _licenceDAO);
        _;
    }

    /// @dev Constructor initializes the card licence contract.
    /// @param _owner is the owner account of the wallet contract.
    /// @param _transferable indicates whether the contract ownership can be transferred.
    /// @param _licence is the initial card licence amount. this number is scaled 10 = 1%, 9 = 0.9%
    /// @param _float is the address of the multi-sig cryptocurrency float contract.
    /// @param _holder is the address of the token holder contract
    constructor(address _owner, bool _transferable, uint _licence, address _float, address _holder) Ownable(_owner, _transferable) public {
        _licenceAmountScaled = _licence;
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

    /// @return the address of the DAO contract.
    function licenceDAO() external view returns (address) {
        return _licenceDAO;
    }

    function lockFloat() external onlyOwner {
        _lockedCryptoFloat = true;
    }

    function lockHolder() external onlyOwner {
        _lockedTokenHolder = true;
    }

    function lockLicenceDAO() external onlyOwner {
        _lockedLicenceDAO = true;
    }

    function floatLocked() public view returns (bool) {
        return _lockedCryptoFloat;
    }

    function holderLocked() public view returns (bool) {
        return _lockedTokenHolder;
    }

    function licenceDAOLocked() public view returns (bool) {
        return _lockedLicenceDAO;
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

    /// @dev Updates the address of the DAO contract.
    function updateLicenceDAO(address _newDAO) external onlyOwner {
        require(!licenceDAOLocked(), "DAO is locked");
        _licenceDAO = _newDAO;
        emit UpdatedLicenceDAO(msg.sender, _newDAO);
    }

    /// @dev Updates the TKN licence amount
    /// @param _newAmount is a number between 1 and MAX_AMOUNT_SCALE
    function updateLicenceAmount(uint _newAmount) external onlyDAO {
        require(1 <= _newAmount && _newAmount <= MAX_AMOUNT_SCALE, "licence amount out of range");
        _licenceAmountScaled = _newAmount;
        emit UpdatedLicenceAmount(msg.sender, _newAmount);
    }

    /// @dev Load the holder and float contracts based on the licence amount and asset amount.
    /// @param _asset is the address of an ERC20 token or 0x0 for ether.
    /// @param _amount is the amount of assets to be transferred including the licence amount.
    function load(address _asset, uint _amount) external payable {
        uint loadAmount = _amount;
        // If TKN then no licence to be paid
        if (_asset == TKN) {
            require(ERC20(_asset).transferFrom(msg.sender, _cryptoFloat, loadAmount), "TKN transfer from external account was unsuccessful");
        } else {
            loadAmount = _amount.mul(MAX_AMOUNT_SCALE).div(_licenceAmountScaled + MAX_AMOUNT_SCALE);
            uint licenceAmount = _amount.sub(loadAmount);

            if (_asset != address(0)) {
                require(ERC20(_asset).transferFrom(msg.sender, _tokenHolder, licenceAmount), "ERC20 licenceAmount transfer from external account was unsuccessful");
                require(ERC20(_asset).transferFrom(msg.sender, _cryptoFloat, loadAmount), "ERC20 token transfer from external account was unsuccessful");
            } else {
                require(msg.value == _amount, "ether sent is not equal to amount");
                _tokenHolder.transfer(licenceAmount);
                _cryptoFloat.transfer(loadAmount);
            }

            emit TransferredToTokenHolder(msg.sender, _tokenHolder, _asset, licenceAmount);
        }

        emit TransferredToCryptoFloat(msg.sender, _cryptoFloat, _asset, loadAmount);
    }
}
