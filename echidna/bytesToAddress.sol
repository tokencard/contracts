pragma solidity ^0.5.10;

contract CryticInterface{
    address internal crytic_owner = address(0x41414141);
    address internal crytic_user = address(0x42424242);
    address internal crytic_attacker = address(0x43434343);
}

contract TEST is CryticInterface {

    bytes example = "0xa462d983B4b8C855e1876e8c24889CBa466A67EB";
    bytes input = example;
    address output;

    constructor() public {
        //{CONSTRUCTOR}
    }

    function input_bytes(bytes memory _input) public {
        input = _input;
        output = _bytesToAddress(input, 0);
    }
	    
    /// @dev This function converts to an address
    /// @param _bts bytes
    /// @param _from start position
    function _bytesToAddress(bytes memory _bts, uint _from) private pure returns (address) {
        require(_bts.length >= _from + 20, "slicing out of range");

        uint160 m = 0;
        uint160 b = 0;

        for (uint8 i = 0; i < 20; i++) {
            m *= 256;
            b = uint160 (_bts[_from + i]);
            m += (b);
        }

        return address(m);
    }

    function compare(bytes memory s1, bytes memory s2) internal returns (bool) {
        return keccak256(abi.encodePacked(s1)) == keccak256(abi.encodePacked(s2));
    }
 
    function crytic_addresses() public returns (bool) {
        return compare(input, example);
    }

}
