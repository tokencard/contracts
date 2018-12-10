pragma solidity >=0.4.18;

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
contract TokenHolder {
    // Owner of this contract.
    address public owner;

    // New owner in a two-phase ownership transfer.
    address public newOwner;

    // Burner token which can be burned to redeem shares.
    address public burner;

    // Tokens known to this contract. When the burn method is called, only
    // shares of these tokens will be disbursed.
    address[] public tokens;

    // Only allow the given address to call the method.
    modifier only(address a) {
        require(msg.sender == a);
        _;
    }

    // Construct a TokenHolder for the given burner token with the sender
    // as the owner.
    function TokenHolder(address b) public {
        owner = msg.sender;
        burner = b;
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

    // Set the tokens known to this contract.
    function setTokens(address[] t) external only (owner) {
        tokens = t;
    }

    // Burn handles disbursing a share of tokens to an address.
    function burn(address to, uint amount) external only (burner) returns (bool) {
        if (amount == 0) {
            return true;
        }

        // The burner token deducts from the supply before calling.
        uint supply = BurnerToken(burner).currentSupply() + amount;

        require(amount <= supply);

        for (uint i = 0; i < tokens.length; i++) {
            uint total = balance(tokens[i]);

            if (total > 0) {
                require((total * amount) / amount == total); // Overflow check.
                transfer(to, tokens[i], (total * amount) / supply);
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
