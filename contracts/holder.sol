pragma solidity ^0.5.7;

import "./externals/ERC20.sol";
import "./externals/SafeMath.sol";
import "./internals/claimable.sol";
import "./internals/ownable.sol";
import "./internals/tokenWhitelistable.sol";
import "./mocks/burnerToken.sol";


// A TokenHolder holds token assets for a burnable token. When the contract
// calls the burn method, a share of the tokens held by this contract are
// disbursed to the burner.
contract Holder is Ownable, Claimable, ENSResolvable, TokenWhitelistable {

    using SafeMath for uint256;

    /// @dev Check if the sender is the burner contract
    modifier onlyBurner() {
        require (msg.sender == _burner, "burner contract is not the sender");
        _;
    }

    // Burner token which can be burned to redeem shares.
    address private _burner;

    /// @notice Constructor  initializes the holder contract.
    /// @param _burnerContract is the address of the token contract with burning functionality.
    /// @param _ens is the address of the ENS registry.
    /// @param _tokenWhitelistNameHash is the ENS name hash of the Token whitelist.
    constructor (address payable _owner, bool _transferable, address _burnerContract, address _ens, bytes32 _tokenWhitelistNameHash) Ownable(_owner, _transferable) ENSResolvable(_ens) TokenWhitelistable(_tokenWhitelistNameHash) public {
        _burner = _burnerContract;
    }

    // Ether may be sent from anywhere.
    function() external payable {}

    // Burn handles disbursing a share of tokens to an address.
    function burn(address payable to, uint amount) external onlyBurner returns (bool) {
        if (amount == 0) {
            return true;
        }

        // The burner token deducts from the supply before calling.
        uint supply = Burner(_burner).currentSupply() + amount;

        require(amount <= supply, "amount greater that total supply!");
        address[] memory burnableAddresses = _burnableTokens();
        for (uint i = 0; i < burnableAddresses.length; i++) {
            uint burnableBalance = balance(burnableAddresses[i]);
            if (burnableBalance > 0) {
                _claim(to, burnableAddresses[i], burnableBalance.mul(amount).div(supply));
            }
        }
        return true;
    }

    function nonBurnableTokenClaim(address payable _to, address[] calldata _nonBurnableAddresses) external onlyOwner returns (bool) {
        for (uint i = 0; i < _nonBurnableAddresses.length; i++) {
            //revert if token is burnable
            require(!_isTokenBurnable(_nonBurnableAddresses[i]), "burnables cannot be claimed");
            uint claimBalance = balance(_nonBurnableAddresses[i]);
            if (claimBalance > 0) {
                _claim(_to, _nonBurnableAddresses[i], claimBalance);
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


}
