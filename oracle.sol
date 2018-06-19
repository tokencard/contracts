pragma solidity ^0.4.24;

/// @title The Oracle keeps track of exchange rates for a set of ERC20 tokens.
/// @author TokenCard
contract Oracle {

    address public controller;
    mapping(address => Rate) public rate;

    struct Rate {
        uint value;
        uint decimal;
    }

    constructor() public {
        controller = msg.sender;
    }

    /// @dev Executable only by the controller.
    modifier onlyController() {
        require(msg.sender == controller);
        _;
    }

    // TODO: Account for uint overflow.
    /// @dev Converts token amount to the corresponding ether amount.
    /// @param _token an ERC20 token contract address.
    /// @param _amount amount of token in base units.
    function convert(address _token, uint _amount) public view returns (uint) {
        Rate memory r = rate[_token];
        uint result = _amount * r.value / 10**r.decimal;
        if (result == 0) {
            return 1;
        }
        return result;
    }

    /// @dev Sets the exchange rate for the selected tokens.
    /// @param _tokens a list of ERC20 tokens (matched with _values by index).
    /// @param _decimals a list of corresponding decimal places for the token.
    /// @param _values a list of corresponding exchange values in wei per unit of token (e.g. 25000000 wei / 1 TKN).
    function set(address[] _tokens, uint[] _values, uint[] _decimals) public onlyController {
        require(_tokens.length == _values.length);
        for (uint i = 0; i < _tokens.length; i++) {
            rate[_tokens[i]] = Rate(_values[i], _decimals[i]);

        }
    }
}
