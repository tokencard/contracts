pragma solidity ^0.4.25;

import "./externals/SafeMath.sol";
import "./internals/ownable.sol";
import "./dao.sol";

/// @title ERC20 interface is a subset of the ERC20 specification.
interface ERC20 {
    function approve(address, uint256) external returns (bool);
    function balanceOf(address) external view returns (uint);
    function transfer(address, uint) external returns (bool);
    function transferFrom(address _from, address _to, uint256 _value) external returns (bool success);
}

/// @title ILicence interface describes methods for loading a TokenCard inclusive of licence fees.
interface ILicence {
    function load(address, uint, uint) external payable returns (bool);
    function updateFee(uint) external;
}

/// @title Licence loads the TokenCard and transfers the licence fee to the token holder contract.
contract Licence is Ownable {

    event Received(address _from, uint _amount);

    event UpdatedDAO(address _sender, address _newDAO);
    event UpdatedFee(address _sender, uint _newFee);

    event TransferredToTokenHolder(address _from, address _contract, address _asset, uint _amount);
    event TransferredToCryptoFloat(address _from, address _contract, address _asset, uint _amount);

    using SafeMath for uint256;

    address public cryptoFloat;
    address public tokenHolder;

    uint public fee;

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
    /// @param _cryptoFloat is the address of the multi-sig cryptocurrency float contract.
    /// @param _tokenHolder is the address of the token holder contract
    constructor(address _owner, bool _transferable, uint _fee, address _cryptoFloat, address _tokenHolder) Ownable(_owner, _transferable) public {
        fee = _fee;
        cryptoFloat = _cryptoFloat;
        tokenHolder = _tokenHolder;
    }

    /// @dev Ether can be deposited from any source, so this contract should be payable by anyone.
    function() public payable {
        require(msg.data.length == 0);
        emit Received(msg.sender, msg.value);
    }

    /// @return the address of the multi-sig cryptocurrency float contract.
    function cryptoFloat() external view returns (address) {
        return cryptoFloat;
    }

    /// @return the address of the token holder contract.
    function tokenHolder() external view returns (address) {
        return tokenHolder;
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
        require(1 <= _newFee && _newFee <= 100, "percent fee out of range"); // TODO(daniel): same as below, not sure if using percentages inside solidity is the optimal solution.
        fee = _newFee;
        emit UpdatedFee(address(DAO), _newFee);
    }

    /// @dev Load the holder and float contracts based on the licence fee and asset amount.
    /// @param _fee is the card licence fee in wei.
    /// @param _asset is the address of an ERC20 token or 0x0 for ether.
    /// @param _amount is the amount of assets to be transferred in base units.
    function load(address _asset, uint _amount, uint _fee) external payable returns (bool) {
        require(_amount != 0 || _fee != 0, "fee and amount cannot both be zero");

        require(_amount == _fee * 100 / fee, "these do not add up"); // TODO(daniel): don't know if it's a good idea to calculate percentages inside a smart contract, best to stick to uints. Also should use safe math here. I guess we could calculate the fee outside the chain and just submit the amount in wei?

        if (_asset != address(0)) {
            require(ERC20(_asset).transferFrom(msg.sender, tokenHolder, _fee), "ERC20 fee transfer from external account was unsuccessful");
            require(ERC20(_asset).transferFrom(msg.sender, cryptoFloat, _amount), "ERC20 token transfer from external account was unsuccessful");
        } else {
            require(msg.value == _amount.add(_fee), "ether sent it not equal to amount + fee");
            tokenHolder.transfer(_fee);
            cryptoFloat.transfer(_amount);
        }
        emit TransferredToTokenHolder(msg.sender, tokenHolder, _asset, _fee);
        emit TransferredToCryptoFloat(msg.sender, cryptoFloat, _asset, _amount);

        return true;
    }
}
