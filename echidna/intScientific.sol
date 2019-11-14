pragma solidity ^0.5.10;

import "contracts/internals/parseIntScientific.sol";

contract CryticInterface{
    address internal crytic_owner = address(0x41414141);
    address internal crytic_user = address(0x42424242);
    address internal crytic_attacker = address(0x43434343);
}

contract TEST is CryticInterface, ParseIntScientific {

    uint u1;
    uint u2;

    // compares _parseIntScientific's output to that of NodeJS 10 and jq 1.5
    // (see TOB-TokenCard-002 for more info)
    function crytic_javascript_artifacts() public returns (bool) {
        return _parseIntScientific("1152921504606846976",0) == 1152921504606847000;
    }

    function crytic_diff() public returns (bool) {
        return u1 == u2;
    }

    function parseDiff(string _a, uint _b) public {
        u1 = _parseIntScientific(_a, _b);
        u2 = parseInt(_a, _b);
    }

    // taken from contracts/externals/oraclizeAPI_0.4.25.sol
    // parseInt(parseFloat*10^_b)
    function parseInt(string _a, uint _b) internal pure returns (uint) {
        string a = _a;
        uint b = _b;
        bytes memory bresult = bytes(a);
        uint mint = 0;
        bool decimals = false;
        for (uint i = 0; i<bresult.length; i++){
            if ((bresult[i] >= 48)&&(bresult[i] <= 57)){
                if (decimals){
                    if (b == 0) break;
                    else b--;
                }
                mint *= 10;
                mint += uint(bresult[i]) - 48;
            } else if (bresult[i] == 46) decimals = true;
        }
        if (_b > 0) mint *= 10**_b;
        return mint;
    }

    function uint2str(uint _i) internal pure returns (string){
        uint i = _i;
        if (i == 0) return "0";
        uint j = i;
        uint len;
        while (j != 0){
            len++;
            j /= 10;
        }
        bytes memory bstr = new bytes(len);
        uint k = len - 1;
        while (i != 0){
            bstr[k--] = byte(48 + i % 10);
            i /= 10;
        }
        return string(bstr);
    }

}
