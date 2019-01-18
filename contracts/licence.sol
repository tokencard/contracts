pragma solidity >=0.4.25;

/// @title A subset of the ERC-20 specification.
interface ERC20 {    
    function transferFrom(address _from, address _to, uint256 _value) public returns (bool success);
    function balanceOf(address) view external returns (uint);    
}    

/// @title This Contract is used to pay the TKN licence
contract Licence {

    event Received(address _from, uint _amount);
    event LockedDAO();
    event UpdatedDAO(address dao);
    event LoadedTokenCard(address sender, uint _amount, uint _fee, address _asset);
    event ChangedOwner(address owner);
    event ChangedLicenceFee(uint _newFee);

    // TOP UP goes here
    // Gnosis Wallet 
    address cryptoFloat;

    // licence fee goes here
    // TKN Holder Contract
    address tokenHolder;

    // Owner of this contract.
    address public owner;

    // New owner in a two-phase ownership transfer.
    address public newOwner;

    // Licence Fee
    uint public licenceFee = 1;

    // Target contract to manager percentages
    address public dao;

    // once locked, can no longer change the dao
    bool public lockedDAO;
 
    // lock the dao
    function lockDAO() only(owner) public {
        lockedDAO = true;
        emit LockedDAO();
    }
 
    // change the address of the dao
    function setDAO(address _dao) public only (owner) {
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
        licenceFee = _newFee; 
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
    function loadTokenCard(uint _amount, uint _fee, address _asset) external payable {
        require(_amount != 0 || _fee != 0);
        require(_amount == _fee * 100 / licenceFee, "these don't add up");

        if (_asset != address(0)) {
            require(ERC20(_asset).transferFrom(msg.sender, tokenHolder, _fee), "ERC20 token transfer was unsuccessful");
            require(ERC20(_asset).transferFrom(msg.sender, cryptoFloat, _amount), "ERC20 token transfer was unsuccessful");
        } else {
            require(msg.value == _amount + _fee, "eth sent it not equal to amount + fee");
            tokenHolder.transfer(_fee);
            cryptoFloat.transfer(_amount);
        }

        emit LoadedTokenCard(msg.sender, _amount, _fee, _asset);
    }
    
}
