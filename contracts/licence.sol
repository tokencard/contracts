pragma solidity >=0.4.25;

 // The Token interface is a subset of the ERC20 specification.    
interface Token {    
    function transfer(address, uint) external returns (bool);    
    function balanceOf(address) view external returns (uint);    
}    

/// @title The DAO interface to changing TKN licence fee
contract DAO {
    function setLicenceFee(uint _amount) public returns (uint result) {
        return 1;
    }
}

contract Licence {
    // Owner of this contract.
    address cryptoFloat;

    address tokenHolder;

    address public owner;

    // New owner in a two-phase ownership transfer.
    address public newOwner;

    // Licence Fee
    uint public licenceFee = 1;

    // Burner token which can be burned to redeem shares.
    address public dao;

    // once locked, can no longer change the dao
    bool public lockedDAO;
 
    // lock the dao
    function lockDAO() only(owner) public {
        lockedDAO = true;
    }
 
    // change the address of the dao
    function setDAO(address _dao) public only (owner) {
        require(!lockedDAO, "DAO is not locked in");
        dao = _dao;
    }

    // Only allow the given address to call the method.
    modifier only(address a) {
        require(msg.sender == a);
        _;
    }

    // Only allow one of the two address to call the method.
    modifier either(address a, address b) {
        require(msg.sender == a || msg.sender == b);
        _;
    }

    // Construct a TokenHolder for the given burner token with the sender
    // as the owner.
    constructor (address _dao, address _cryptoFloat, address _tokenHolder) public {
        owner = msg.sender;
        dao = _dao;
        cryptoFloat = _cryptoFloat;
        tokenHolder = _tokenHolder;
    }

    // Ether may be sent from anywhere.
    function () external payable { }

    function changeLicenceFee(uint _newFee) external only (dao) {
        require(1 <= _newFee && _newFee <= 100, "percent fee out of range");
        licenceFee = _newFee; 
    }

    // Change owner in a two-phase ownership transfer.
    function changeOwner(address to) external only (owner) {
        newOwner = to;
    }

    // Accept ownership in a two-phase ownership transfer.
    function acceptOwnership() external only (newOwner) {
        owner = msg.sender;
        newOwner = address(0);
    }

    // this only works with ETH
    function processTransaction(uint _amount, uint _fee) external payable {
        require(_amount == 0 || _fee == 0);
        require(msg.value == _amount + _fee, "eth sent it not equal to amount + fee");
        require(_amount == _fee * 100 / licenceFee, "these don't add up");

        tokenHolder.transfer(_fee);
        cryptoFloat.transfer(_amount);

    }
    
}
