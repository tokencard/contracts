pragma solidity ^0.4.25;

import "../internals/parseIntScientific.sol";

contract ParseIntScientificExporter is ParseIntScientific {

    /// @dev exports _parseIntScientific(string) as an external function.
    function parseIntScientific(string _a) external pure returns (uint) {
        return _parseIntScientific(_a);
    }

    /// @dev exports _parseIntScientific(string, uint) as an external function.
    function parseIntScientificDecimals(string _a, uint _b) external pure returns (uint) {
        return _parseIntScientific(_a, _b);
    }

    /// @dev exports _parseIntScientificWei(string) as an external function.
    function parseIntScientificWei(string _a) external pure returns (uint) {
        return _parseIntScientificWei(_a);
    }

}
