pragma solidity ^0.5.10;

import "contracts/externals/base64.sol";

contract CryticInterface{
    address internal crytic_owner = address(0x41414141);
    address internal crytic_user = address(0x42424242);
    address internal crytic_attacker = address(0x43434343);
}

contract TEST is CryticInterface, Base64 {

    bytes b;
    function crytic_random_base64() public view returns (bool) {
        // it is very unlikely that random bytes are valid base 64
        return b.length == 0;
    }

    function base64decode(bytes memory _b) private {
        b = _base64decode(_b);
    }
}
