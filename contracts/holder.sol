pragma solidity >=0.4.18;

import "./externals/SafeMath.sol";

// The Token interface is a subset of the ERC20 specification.
interface Token {
    function transfer(address, uint) external returns (bool);
    function balanceOf(address) constant external returns (uint);
}

// The BurnerToken interface is the interface to a token contract which
// provides the total burnable supply for the TokenHolder contract.
interface BurnerToken {
    function currentSupply() constant external returns (uint);
}

// A TokenHolder holds token assets for a burnable token. When the contract
// calls the burn method, a share of the tokens held by this contract are
// disbursed to the burner.
contract Holder {

  using SafeMath for uint256;

  /*******************/
  /*     Events     */
  /*****************/

  event AddedToken(address _sender, address _token);
  event RemovedToken(address _sender, address _token);

  // Owner of this contract.
  address public owner;

  // New owner in a two-phase ownership transfer.
  address public newOwner;

  // Burner token which can be burned to redeem shares.
  address public burner;

  // Tokens known to this contract. When the burn method is called, only
  // shares of these tokens will be disbursed.
  address[] private _tokenAddresses;
  mapping(address => bool) public tokenExists;

  // Only allow the given address to call the method.
  modifier only(address a) {
      require(msg.sender == a);
      _;
  }

    // Construct a TokenHolder for the given burner token with the sender
    // as the owner.
    constructor (address burnerContract) public {
        owner = msg.sender;
        burner = burnerContract;
    }

    // Ether may be sent from anywhere.
    function () external payable { }

    // Change owner in a two-phase ownership transfer.
    function changeOwner(address to) external only (owner) {
        newOwner = to;
    }

    // Accept ownership in a two-phase ownership transfer.
    function acceptOwnership() external only (newOwner) {
        owner = msg.sender;
        newOwner = 0x0;
    }

    /// @dev Add ERC20 tokens to the list of supported tokens.
    /// @param _tokens ERC20 token contract addresses.
    function addTokens(address[] _tokens) external only (owner) {
        // Add each token to the list of supported tokens.
        for (uint i = 0; i < _tokens.length; i++) {
            // Require that the token doesn't already exist.
            address token = _tokens[i];
            require(!tokenExists[token], "token already exists");
            // Add the token address to the address list.
            _tokenAddresses.push(token);
            tokenExists[token] = true;
            // Emit token addition event.
            emit AddedToken(msg.sender, token);
        }
    }

    /// @dev Remove ERC20 tokens from the list of supported tokens.
    /// @param _tokens ERC20 token contract addresses.
    function removeTokens(address[] _tokens) external only (owner) {
        // Delete each token object from the list of supported tokens based on the addresses provided.
        for (uint i = 0; i < _tokens.length; i++) {
            // Store the token address.
            address token = _tokens[i];
            //token must exist, reverts on duplicates as well
            require(tokenExists[token], "token does not exist");
            // Delete the token entry.
            delete tokenExists[token];
            // Remove the token address from the address list.
            for (uint j = 0; j < _tokenAddresses.length.sub(1); j++) {
                if (_tokenAddresses[j] == token) {
                    _tokenAddresses[j] = _tokenAddresses[_tokenAddresses.length.sub(1)];
                    break;
                }
            }
            _tokenAddresses.length--;
            // Emit token removal event.
            emit RemovedToken(msg.sender, token);
        }
    }

    // Burn handles disbursing a share of tokens to an address.
    function burn(address to, uint amount) external only (burner) returns (bool) {
        if (amount == 0) {
            return true;
        }

        // The burner token deducts from the supply before calling.
        uint supply = BurnerToken(burner).currentSupply() + amount;

        require(amount <= supply);

        for (uint i = 0; i < _tokenAddresses.length; i++) {
            uint total = balance(_tokenAddresses[i]);

            if (total > 0) {
                require((total * amount) / amount == total); // Overflow check.
                transfer(to, _tokenAddresses[i], (total * amount) / supply);
            }
        }

        return true;
    }

    // Balance returns the amount of a token owned by the contract.
    function balance(address token) constant public returns (uint) {
        if (token != 0x0) {
            return Token(token).balanceOf(this);
        } else {
            return address(this).balance;
        }
    }

    // Transfer funds to an address.
    function transfer(address to, address token, uint amount) private {
        if (amount == 0) {
            return;
        }

        if (token != 0x0) {
            require(Token(token).transfer(to, amount));
        } else {
            to.transfer(amount);
        }
    }
}
