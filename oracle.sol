pragma solidity ^0.4.24;

/// @title The Oracle keeps track of exchange rates for a set of ERC20 tokens.
/// @author TokenCard
contract Oracle {

    address public controller;
    mapping(address => uint) public rates;

    constructor() public {
        controller = msg.sender;
    }

    /// @dev Executable only by the controller.
    modifier onlyController() {
        require(msg.sender == controller);
        _;
    }

    /// @dev Sets the exchange rates for the selected tokens.
    /// @param _tokens a list of ERC20 tokens (matched with _values by index).
    /// @param _values a list of corresponding exchange values in wei per base unit of token (e.g. 25000000 wei / 0.00000001 TKN).
    function setRates(address[] _tokens, uint[] _values) public onlyController {
        require(_tokens.length == _values.length);
        for (uint i = 0; i < _tokens.length; i++) {
            rates[_tokens[i]] = _values[i];
        }
    }
}
