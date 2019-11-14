pragma solidity ^0.5.10;

import "contracts/externals/strings.sol";

contract CryticInterface{
    address internal crytic_owner = address(0x41414141);
    address internal crytic_user = address(0x42424242);
    address internal crytic_attacker = address(0x43434343);
}

contract TEST is CryticInterface {

    using strings for *;

    string example = "{\"ETH\":1.0}";
    string input = example;
    string out = "";
    bytes32 constant private _PREFIX_HASH = keccak256("{\"ETH\":");

    function input_json(string memory _json) public {
        input = _json;
        out = _parseRate(_json);
    }

    // the following function has been written by Alex Beregszaszi (@axic), use it under the terms of the MIT license
    function copyBytes(bytes memory from, uint fromOffset, uint length, bytes memory to, uint toOffset) internal pure returns (bytes memory) {
        uint minLength = length + toOffset;

        // Buffer too small
        require(to.length >= minLength, ""); // Should be a better way?

        // NOTE: the offset 32 is added to skip the `size` field of both bytes variables
        uint i = 32 + fromOffset;
        uint j = 32 + toOffset;

        while (i < (32 + fromOffset + length)) {
            assembly {
                let tmp := mload(add(from, i))
                mstore(add(to, j), tmp)
            }
            i += 32;
            j += 32;
        }

        return to;
    }

    /// @notice Extracts JSON rate value from the response object.
    /// @param _json body of the JSON response from the CryptoCompare API.
    function _parseRate(string memory _json) internal pure returns (string memory) {

        uint jsonLen = abi.encodePacked(_json).length;
        //{"ETH":}.length = 8, assuming a (maximum of) 18 digit prevision
        require(jsonLen > 8 && jsonLen <= 28, "misformatted input");

        bytes memory jsonPrefix = new bytes(7);
        copyBytes(abi.encodePacked(_json), 0, 7, jsonPrefix, 0);
        require(keccak256(jsonPrefix) == _PREFIX_HASH, "prefix mismatch");

        strings.slice memory body = _json.toSlice();
        body.split(":".toSlice());
        //we are sure that ':' is included in the string, body now contains the rate+'}'
        jsonLen = body._len;
        body.until("}".toSlice());
        require(body._len == jsonLen - 1, "not json format");
        //ensure that the json is properly terminated with a '}'
        return body.toString();
    }

    function compare(string memory s1, string memory s2) internal pure returns (bool) {
        return keccak256(abi.encodePacked(s1)) == keccak256(abi.encodePacked(s2));
    }

    function crytic_json() public view returns (bool) {
        return compare(input, example);
    }

}
