pragma solidity ^0.4.24;

contract Oracle {

    address public controller;

    mapping(address => uint) public rates;

    modifier onlyController() {
        require(msg.sender == controller);
        _;
    }

    constructor() public {
        controller = msg.sender;
    }

    function setRates(address[] _tokens, uint[] _values) public onlyController {
        require(_tokens.length == _values.length);
        for (uint i = 0; i < _tokens.length; i++) {
            rates[_tokens[i]] = _values[i];
        }
    }
}
