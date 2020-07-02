pragma solidity ^0.6.0;

import "../externals/base64.sol";

contract Base64Exporter is Base64 {
    /// @dev export _base64decode() as an external function.
    function base64decode(bytes calldata _encoded) external pure returns (bytes memory) {
        return _base64decode(_encoded);
    }
}
