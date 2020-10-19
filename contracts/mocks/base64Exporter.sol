<<<<<<< HEAD
// SPDX-License-Identifier: GPLv3
=======
pragma solidity ^0.6.0;
>>>>>>> f9e39260... Use solc 0.6.4 for oracle contract and remove the oraclize service

pragma solidity ^0.6.11;

import "../externals/base64.sol";

contract Base64Exporter is Base64 {
    /// @dev export _base64decode() as an external function.
    function base64decode(bytes calldata _encoded) external pure returns (bytes memory) {
        return _base64decode(_encoded);
    }
}
