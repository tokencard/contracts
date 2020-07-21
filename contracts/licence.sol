/**
 *  Licence - The Consumer Contract Wallet
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

// SPDX-License-Identifier: GPLv3

pragma solidity ^0.6.11;

import "./externals/SafeMath.sol";
import "./externals/SafeERC20.sol";
import "./internals/controllable.sol";
import "./internals/ensResolvable.sol";
import "./internals/transferrable.sol";


/// @title ILicence interface describes methods for loading a TokenCard and updating licence amount.
interface ILicence {
    function load(address, uint256) external payable;

    function updateLicenceAmount(uint256) external;
}


/// @title Licence loads the TokenCard and transfers the licence amout to the TKN Holder Contract.
/// @dev the rest of the amount gets sent to the CryptoFloat
contract Licence is Transferrable, ENSResolvable, Controllable {
    using SafeMath for uint256;
    using SafeERC20 for IERC20;

    /*******************/
    /*     Events     */
    /*****************/

    event UpdatedLicenceDAO(address _newDAO);
    event UpdatedCryptoFloat(address _newFloat);
    event UpdatedTokenHolder(address _newHolder);
    event UpdatedTKNContractAddress(address _newTKN);
    event UpdatedLicenceAmount(uint256 _newAmount);

    event TransferredToTokenHolder(address _from, address _to, address _asset, uint256 _amount);
    event TransferredToCryptoFloat(address _from, address _to, address _asset, uint256 _amount);

    event Claimed(address _to, address _asset, uint256 _amount);

    /// @dev This is 100% scaled up by a factor of 10 to give us an extra 1 decimal place of precision
    uint256 public constant MAX_AMOUNT_SCALE = 1000;
    uint256 public constant MIN_AMOUNT_SCALE = 1;

    address private _tknContractAddress = 0xaAAf91D9b90dF800Df4F55c205fd6989c977E73a; // solium-disable-line uppercase

    address payable private _cryptoFloat;
    address payable private _tokenHolder;
    address private _licenceDAO;

    bool private _lockedCryptoFloat;
    bool private _lockedTokenHolder;
    bool private _lockedLicenceDAO;
    bool private _lockedTKNContractAddress;

    /// @dev This is the _licenceAmountScaled by a factor of 10
    /// @dev i.e. 1% is 10 _licenceAmountScaled, 0.1% is 1 _licenceAmountScaled
    uint256 private _licenceAmountScaled;

    /// @dev Reverts if called by any address other than the DAO contract.
    modifier onlyDAO() {
        require(msg.sender == _licenceDAO, "the sender isn't the DAO");
        _;
    }

    /// @dev Constructor initializes the card licence contract.
    /// @param _licence_ is the initial card licence amount. this number is scaled 10 = 1%, 9 = 0.9%
    /// @param _float_ is the address of the multi-sig cryptocurrency float contract.
    /// @param _holder_ is the address of the token holder contract
    /// @param _tknAddress_ is the address of the TKN ERC20 contract
    /// @param _ens_ is the address of the ENS Registry
    /// @param _controllerNode_ is the ENS node corresponding to the controller
    constructor(uint256 _licence_, address payable _float_, address payable _holder_, address _tknAddress_, address _ens_, bytes32 _controllerNode_) public {
        require(MIN_AMOUNT_SCALE <= _licence_ && _licence_ <= MAX_AMOUNT_SCALE, "licence amount out of range");
        _initializeENSResolvable(_ens_);
        _initializeControllable(_controllerNode_);
        _licenceAmountScaled = _licence_;
        _cryptoFloat = _float_;
        _tokenHolder = _holder_;
        if (_tknAddress_ != address(0)) {
            _tknContractAddress = _tknAddress_;
        }
    }

    /// @dev Ether can be deposited from any source, so this contract should be payable by anyone.
    receive() external payable {}

    /// @dev this allows for people to see the scaled licence amount
    /// @return the scaled licence amount, used to calculate the split when loading.
    function licenceAmountScaled() external view returns (uint256) {
        return _licenceAmountScaled;
    }

    /// @dev allows one to see the address of the CryptoFloat
    /// @return the address of the multi-sig cryptocurrency float contract.
    function cryptoFloat() external view returns (address) {
        return _cryptoFloat;
    }

    /// @dev allows one to see the address TKN holder contract
    /// @return the address of the token holder contract.
    function tokenHolder() external view returns (address) {
        return _tokenHolder;
    }

    /// @dev allows one to see the address of the DAO
    /// @return the address of the DAO contract.
    function licenceDAO() external view returns (address) {
        return _licenceDAO;
    }

    /// @dev The address of the TKN token
    /// @return the address of the TKN contract.
    function tknContractAddress() external view returns (address) {
        return _tknContractAddress;
    }

    /// @dev This locks the cryptoFloat address
    /// @dev so that it can no longer be updated
    function lockFloat() external onlyAdmin {
        _lockedCryptoFloat = true;
    }

    /// @dev This locks the TokenHolder address
    /// @dev so that it can no longer be updated
    function lockHolder() external onlyAdmin {
        _lockedTokenHolder = true;
    }

    /// @dev This locks the DAO address
    /// @dev so that it can no longer be updated
    function lockLicenceDAO() external onlyAdmin {
        _lockedLicenceDAO = true;
    }

    /// @dev This locks the TKN address
    /// @dev so that it can no longer be updated
    function lockTKNContractAddress() external onlyAdmin {
        _lockedTKNContractAddress = true;
    }

    /// @dev Updates the address of the cyptoFloat.
    /// @param _newFloat This is the new address for the CryptoFloat
    function updateFloat(address payable _newFloat) external onlyAdmin {
        require(!floatLocked(), "float is locked");
        _cryptoFloat = _newFloat;
        emit UpdatedCryptoFloat(_newFloat);
    }

    /// @dev Updates the address of the Holder contract.
    /// @param _newHolder This is the new address for the TokenHolder
    function updateHolder(address payable _newHolder) external onlyAdmin {
        require(!holderLocked(), "holder contract is locked");
        _tokenHolder = _newHolder;
        emit UpdatedTokenHolder(_newHolder);
    }

    /// @dev Updates the address of the DAO contract.
    /// @param _newDAO This is the new address for the Licence DAO
    function updateLicenceDAO(address _newDAO) external onlyAdmin {
        require(!licenceDAOLocked(), "DAO is locked");
        _licenceDAO = _newDAO;
        emit UpdatedLicenceDAO(_newDAO);
    }

    /// @dev Updates the address of the TKN contract.
    /// @param _newTKN This is the new address for the TKN contract
    function updateTKNContractAddress(address _newTKN) external onlyAdmin {
        require(!tknContractAddressLocked(), "TKN is locked");
        _tknContractAddress = _newTKN;
        emit UpdatedTKNContractAddress(_newTKN);
    }

    /// @dev Updates the TKN licence amount
    /// @param _newAmount is a number between MIN_AMOUNT_SCALE (1) and MAX_AMOUNT_SCALE
    function updateLicenceAmount(uint256 _newAmount) external onlyDAO {
        require(MIN_AMOUNT_SCALE <= _newAmount && _newAmount <= MAX_AMOUNT_SCALE, "licence amount out of range");
        _licenceAmountScaled = _newAmount;
        emit UpdatedLicenceAmount(_newAmount);
    }

    /// @dev Load the holder and float contracts based on the licence amount and asset amount.
    /// @param _asset is the address of an ERC20 token or 0x0 for ether.
    /// @param _amount is the amount of assets to be transferred including the licence amount.
    function load(address _asset, uint256 _amount) external payable {
        uint256 loadAmount = _amount;
        // If TKN then no licence to be paid
        if (_asset == _tknContractAddress) {
            IERC20(_asset).safeTransferFrom(msg.sender, _cryptoFloat, loadAmount);
        } else {
            loadAmount = _amount.mul(MAX_AMOUNT_SCALE).div(_licenceAmountScaled + MAX_AMOUNT_SCALE);
            uint256 licenceAmount = _amount.sub(loadAmount);

            if (_asset != address(0)) {
                IERC20(_asset).safeTransferFrom(msg.sender, _tokenHolder, licenceAmount);
                IERC20(_asset).safeTransferFrom(msg.sender, _cryptoFloat, loadAmount);
            } else {
                require(msg.value == _amount, "ETH sent is not equal to amount");
                _tokenHolder.transfer(licenceAmount);
                _cryptoFloat.transfer(loadAmount);
            }

            emit TransferredToTokenHolder(msg.sender, _tokenHolder, _asset, licenceAmount);
        }

        emit TransferredToCryptoFloat(msg.sender, _cryptoFloat, _asset, loadAmount);
    }

    //// @dev Withdraw tokens from the smart contract to the specified account.
    function claim(address payable _to, address _asset, uint256 _amount) external onlyAdmin {
        _safeTransfer(_to, _asset, _amount);
        emit Claimed(_to, _asset, _amount);
    }

    /// @dev returns whether or not the CryptoFloat address is locked
    function floatLocked() public view returns (bool) {
        return _lockedCryptoFloat;
    }

    /// @dev returns whether or not the TokenHolder address is locked
    function holderLocked() public view returns (bool) {
        return _lockedTokenHolder;
    }

    /// @dev returns whether or not the Licence DAO address is locked
    function licenceDAOLocked() public view returns (bool) {
        return _lockedLicenceDAO;
    }

    /// @dev returns whether or not the TKN address is locked
    function tknContractAddressLocked() public view returns (bool) {
        return _lockedTKNContractAddress;
    }
}
