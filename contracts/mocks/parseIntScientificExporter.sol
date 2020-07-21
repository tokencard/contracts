// SPDX-License-Identifier: GPLv3

pragma solidity ^0.6.11;

import "../internals/parseIntScientific.sol";


contract ParseIntScientificExporter is ParseIntScientific {
    /// @dev exports _parseIntScientific(string) as an external function.
    function parseIntScientific(string calldata _a) external pure returns (uint256) {
        return _parseIntScientific(_a);
    }

    /// @dev exports _parseIntScientific(string, uint) as an external function.
    function parseIntScientificDecimals(string calldata _a, uint256 _b) external pure returns (uint256) {
        return _parseIntScientific(_a, _b);
    }

    /// @dev exports _parseIntScientificWei(string) as an external function.
    function parseIntScientificWei(string calldata _a) external pure returns (uint256) {
        return _parseIntScientificWei(_a);
    }
}
