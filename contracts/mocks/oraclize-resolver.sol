pragma solidity ^0.4.24;


contract OraclizeAddrResolver {

    address private oraclizedAddress;

    constructor(address _oraclizedAddress) public {
        oraclizedAddress = _oraclizedAddress;
    }

    function getAddress() public returns (address) {
        return oraclizedAddress;
    }
}
