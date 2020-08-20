// SPDX-License-Identifier: GPLv3

pragma solidity ^0.6.11;

contract GasBurner {
    function dummy() public pure {
        assembly {
            invalid()
        }
    }

    function burnGas(uint256 burn) public {
        // Calls self.dummy() to burn gas.
        assembly {
            mstore(0x0, 0x32e43a1100000000000000000000000000000000000000000000000000000000)
            let ret := call(burn, address(), 0, 0x0, 0x04, 0x0, 0)
        }
    }
}
