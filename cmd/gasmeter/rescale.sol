pragma solidity ^0.4.24;

contract Test {

    function rescale_via_exp(uint amount, uint scale) private pure returns (uint) {
        return amount * (10 ** scale);
    }

    function rescale_via_mul(uint amount, uint scale) private pure returns (uint) {
        uint _amount = amount;
        for (uint i = 0; i < scale; i++) {
            _amount *= 10;
        }
        return _amount;
    }

    function via_exp() public pure {
        rescale_via_exp(1234, 18);
    }

    function via_mul() public pure {
        rescale_via_mul(1234, 18);
    }

}
