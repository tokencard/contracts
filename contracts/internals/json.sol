/**
 *  TokenWallet - The Consumer Contract Wallet
 *  Copyright (C) 2018 Token Group Ltd
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

/// @title JSON provides JSON parsing functionality.
contract JSON {
    using strings for *;

    /// @dev Extracts JSON rate value from the response object.
    /// @param _json body of the JSON response from the CryptoCompare API.
    /// @param _symbol asset symbol used to extract the correct json field.
    function parseRate(string _json, string _symbol) internal pure returns (string) {
        strings.slice memory body = _json.toSlice();
        body.beyond("{".toSlice());
        body.until("}".toSlice());
        strings.slice memory _quote_mark = "\"".toSlice();
        body.find(_quote_mark.concat(_symbol.toSlice()).toSlice().concat(_quote_mark).toSlice());
        strings.slice memory asset;
        body.split(",".toSlice(), asset);
        asset.split(":".toSlice());
        return asset.toString();
    }
}

