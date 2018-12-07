pragma solidity ^0.4.25;

import "../internals/parseIntScientific.sol";

contract ParseIntScientificExporter is ParseIntScientific {

    /// @dev export _parseIntScientific(string) as a public function.
    function parseIntScientific(string _a) public pure returns (uint) {
        return _parseIntScientific(_a, 0);
    }

    /// @dev export _parseIntScientific(string, uint) as a public function.
    function parseIntScientificDecimals(string _a, uint _b) public pure returns (uint) {
        return _parseIntScientific(_a, _b);
    }

}
