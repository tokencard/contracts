pragma solidity ^0.4.25;

import "../externals/base64.sol";

contract Base64Exporter is Base64 {

    /// @dev export _base64decode() as an external function.
    function base64decode(bytes _encoded) external pure returns (bytes) {
        return _base64decode(_encoded);
    }

}
