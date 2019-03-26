/**
 *  The Consumer Contract Wallet
 *  Copyright (C) 2018 The Contract Wallet Company Limited
 *
 *  This program is free software: you can redistribute it and/or modify
 *  it under the terms of the GNU General Public License as published by
 *  the Free Software Foundation, either version 3 of the License, or
 *  (at your option) any later version.

 *  This program is distributed in the hope that it will be useful,
 *  but WITHOUT ANY WARRANTY; without even the implied warranty of
 *  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *  GNU General Public License for more details.

 *  You should have received a copy of the GNU General Public License
 *  along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

pragma solidity ^0.4.25;

import "../externals/strings.sol";
import "../externals/oraclizeAPI_0.4.25.sol";
/// @title JSON provides JSON parsing functionality.
contract JSON is usingOraclize {
    using strings for *;

    bytes32 constant private prefixHash = keccak256("{\"ETH\":");

    /// @dev Extracts JSON rate value from the response object.
    /// @param _json body of the JSON response from the CryptoCompare API.
    function parseRate(string _json) public pure returns (string) {

        uint json_len = abi.encodePacked(_json).length;
        //{"ETH":}.length = 8, assuming a (maximum of) 18 digit prevision
        require(json_len > 8 && json_len <= 28, "misformatted input");

        bytes memory jsonPrefix = new bytes(7);
        copyBytes(abi.encodePacked(_json), 0, 7, jsonPrefix, 0);
        require(keccak256(jsonPrefix) == prefixHash, "prefix mismatch");

        strings.slice memory body = _json.toSlice();
        body.split(":".toSlice());
        //we are sure that ':' is included in the string, body now contains the rate+'}'
        json_len = body._len;
        body.until("}".toSlice());
        require(body._len == json_len - 1, "not json format");
        //ensure that the json is properly terminated with a '}'
        return body.toString();

    }
}
