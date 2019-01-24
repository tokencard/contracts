pragma solidity ^0.4.25;


contract OraclizeAddrResolver {

    address private oraclizedAddress;

    constructor(address _oraclizedAddress) public {
        oraclizedAddress = _oraclizedAddress;
    }

    function getAddress() public view returns (address) {
        return oraclizedAddress;
    }
}
