pragma solidity ^0.5.7;

import "./internals/tokenWhitelistable.sol";
import "./externals/SafeMath.sol";
import "./externals/ERC20.sol";
import "./mocks/burnerToken.sol";


// A TokenHolder holds token assets for a burnable token. When the contract
// calls the burn method, a share of the tokens held by this contract are
// disbursed to the burner.
contract Holder is ENSResolvable, TokenWhitelistable {

    using SafeMath for uint256;

    // Owner of this contract.
    address private _owner;

    // New owner in a two-phase ownership transfer.
    address private _newOwner;

    // Burner token which can be burned to redeem shares.
    address private _burner;

    // Only allow the given address to call the method.
    modifier only(address a) {
        require(msg.sender == a);
        _;
    }

    // Construct a TokenHolder for the given burner token with the sender
    // as the owner.
    constructor (address _burnerContract, address _ens, bytes32 _tokenWhitelistNameHash) ENSResolvable(_ens) TokenWhitelistable(_tokenWhitelistNameHash) public {
        _owner = msg.sender;
        _burner = _burnerContract;
    }

    // Ether may be sent from anywhere.
    function() external payable {}

    // Change owner in a two-phase ownership transfer.
    function changeOwner(address to) external only(_owner) {
        _newOwner = to;
    }

    // Accept ownership in a two-phase ownership transfer.
    function acceptOwnership() external only(_newOwner) {
        _owner = msg.sender;
        _newOwner = address(0);
    }

    // Burn handles disbursing a share of tokens to an address.
    function burn(address payable to, uint amount) external only(_burner) returns (bool) {
        if (amount == 0) {
            return true;
        }

        // The burner token deducts from the supply before calling.
        uint supply = Burner(_burner).currentSupply() + amount;

        require(amount <= supply);
        address[] memory tokenAddresses = _tokenAddressArray();
        for (uint i = 0; i < tokenAddresses.length; i++) {
            uint total = balance(tokenAddresses[i]);

            if (total > 0) {
                require((total * amount) / amount == total);
                // Overflow check.
                _transfer(to, tokenAddresses[i], (total * amount) / supply);
            }
        }

        return true;
    }

    function burner() external view returns (address) {
        return _burner;
    }

    // Balance returns the amount of a token owned by the contract.
    function balance(address token) public view returns (uint) {
        if (token != address(0)) {
            return ERC20(token).balanceOf(address(this));
        } else {
            return address(this).balance;
        }
    }

    // Transfer funds to an address.
    function _transfer(address payable to, address token, uint amount) private {
        if (amount == 0) {
            return;
        }

        if (token != address(0)) {
            require(ERC20(token).transfer(to, amount));
        } else {
            to.transfer(amount);
        }
    }
}
