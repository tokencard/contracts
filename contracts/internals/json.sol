pragma solidity ^0.4.25;

import "../externals/strings.sol";

/// @title JSON provides JSON parsing functionality.
contract JSON {
    using strings for *;

    // @dev Extracts JSON rate value from the response object.
    // @param _json body of the JSON response from the CryptoCompare API.
    // @param _symbol asset symbol used to extract the correct json field.
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

