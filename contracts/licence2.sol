pragma solidity ^0.4.25;

//import "./wallet.sol";
import "./test.sol";


/// @title This Contract is used to pay the TKN licence
contract Licence {

    event Received(address _from, uint _amount);
    event LockedDAO();
    event UpdatedDAO(address _dao);
    event ChangedOwner(address _owner);
    event ChangedLicenceFee(uint _fee);

    // TOP UP goes here
    // Gnosis Wallet 
    address public cryptoFloat;

    // licence fee goes here
    // TKN Holder Contract
    address public tokenHolder;

    // Owner of this contract.
    address public owner;

    // New owner in a two-phase ownership transfer.
    address public newOwner;

    // Licence Fee
    uint public fee = 1;

    // Target contract to manager percentages
    address public dao;

    // once locked, can no longer change the dao
    bool public lockedDAO;
 
    // lock the dao
    function lockDAO() external only(owner) {
        lockedDAO = true;
        emit LockedDAO();
    }
 
    // change the address of the dao
    function setDAO(address _dao) external only (owner) {
        require(!lockedDAO, "DAO is locked");
        dao = _dao;
        emit UpdatedDAO(_dao);
    }

    // Only allow the given address to call the method.
    modifier only(address a) {
        require(msg.sender == a);
        _;
    }

    /// @dev Constructor for the TokenCard Licence Contract
    constructor (address _dao, address _cryptoFloat, address _tokenHolder) public {
        owner = msg.sender;
        dao = _dao;
        cryptoFloat = _cryptoFloat;
        tokenHolder = _tokenHolder;
    }

    /// @dev Ether can be deposited from any source, so this contract must be payable by anyone.
    function() public payable {
        //TODO question: Why is this check here, is it necessary or are we building into a corner?
        require(msg.data.length == 0);
        emit Received(msg.sender, msg.value);
    }

    function changeLicenceFee(uint _newFee) external only (dao) {
        require(1 <= _newFee && _newFee <= 100, "percent fee out of range");
        fee = _newFee; 
        emit ChangedLicenceFee(_newFee);
    }

    // Change owner in a two-phase ownership transfer.
    function changeOwner(address to) external only (owner) {
        // TODO This is kinda nasty i know
        if (to == address(0)) {
            owner = to; 
            emit ChangedOwner(owner);
        }
        newOwner = to;
    }

    // Accept ownership in a two-phase ownership transfer.
    function acceptOwnership() external only (newOwner) {
        owner = msg.sender;
        newOwner = address(0);
        emit ChangedOwner(newOwner);
    }

    /// @dev this is the function that we need to call to do the top-up 
    /// TODO should this method have a bool return?
    function load(uint _amount, uint _licenceFee, address _asset) payable public returns (bool) {
        require(_amount != 0 || _licenceFee != 0);
        require(_amount == _licenceFee * 100 / fee, "these don't add up");

        if (_asset != address(0)) {
            require(ERC20(_asset).transferFrom(msg.sender, tokenHolder, _licenceFee), "ERC20 token transfer was unsuccessful");
            require(ERC20(_asset).transferFrom(msg.sender, cryptoFloat, _amount), "ERC20 token transfer was unsuccessful");
        } else {
            require(msg.value == _amount + _licenceFee, "eth sent it not equal to amount + fee");
            tokenHolder.transfer(_licenceFee);
            cryptoFloat.transfer(_amount);
            // TODO do i need put requires or something up there ?
        }

        return true;
    }
}

