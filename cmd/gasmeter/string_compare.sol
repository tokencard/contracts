pragma solidity ^0.4.24;

contract Test {

    event Log(bool result);

    function asm_equal(string a, string b) private pure returns (bool res) {
        assembly {
            let lenA := mload(a)
            let lenB := mload(b)
            switch eq(lenA, lenB)
            case 1 {
                res := 1
                switch gt(lenA, 0)
                case 1 {
                    let words := div(add(lenA, 31), 32)
                    for { let i := 0 } lt(i, words) { i := add(i, 1) } {
                        let offset := mul(add(i, 1), 32)
                        switch eq(mload(add(a, offset)), mload(add(b, offset)))
                        case 1 {
                        }
                        default {
                            res := 0
                        }
                    }
                }
            }
            default {
                res := 0
            }
        }
    }

    function byte_equal(string a, string b) private pure returns (bool) {
        bytes memory _a = bytes(a);
        bytes memory _b = bytes(b);
        uint l = _a.length;
        if (l != _b.length) {
            return false;
        }
        for (uint i = 0; i < l; i++) {
            if (_a[i] != _b[i]) {
                return false;
            }
        }
        return true;
    }

    function hash_equal(string a, string b) private pure returns (bool) {
        return keccak256(bytes(a)) == keccak256(bytes(b));
    }

    function short_equal_asm() public pure {
        asm_equal("ETH", "ETH");
    }

    function short_equal_byte() public pure {
        byte_equal("ETH", "ETH");
    }

    function short_equal_hash() public pure {
        hash_equal("ETH", "ETH");
    }

    function short_not_equal_asm() public pure {
        asm_equal("ETH", "TKN");
    }

    function short_not_equal_byte() public pure {
        byte_equal("ETH", "TKN");
    }

    function short_not_equal_hash() public pure {
        hash_equal("ETH", "TKN");
    }

    function medium_equal_asm() public pure {
        asm_equal("hello world x hello world", "hello world y hello world");
    }

    function medium_equal_byte() public pure {
        byte_equal("hello world x hello world", "hello world y hello world");
    }

    function medium_equal_hash() public pure {
        hash_equal("hello world x hello world", "hello world y hello world");
    }

    function long_equal_asm() public pure {
        asm_equal("hello world x hello world hello world x hello world hello world x hello world hello world x hello world hello world x hello world hello world x hello world hello world x hello world hello world x hello world hello world x hello world hello world x hello world", "hello world x hello world hello world y hello world hello world x hello world hello world x hello world hello world x hello world hello world x hello world hello world x hello world hello world x hello world hello world x hello world hello world x hello world");
    }

    function long_equal_byte() public pure {
        byte_equal("hello world x hello world hello world x hello world hello world x hello world hello world x hello world hello world x hello world hello world x hello world hello world x hello world hello world x hello world hello world x hello world hello world x hello world", "hello world x hello world hello world y hello world hello world x hello world hello world x hello world hello world x hello world hello world x hello world hello world x hello world hello world x hello world hello world x hello world hello world x hello world");
    }

    function long_equal_hash() public pure {
        hash_equal("hello world x hello world hello world x hello world hello world x hello world hello world x hello world hello world x hello world hello world x hello world hello world x hello world hello world x hello world hello world x hello world hello world x hello world", "hello world x hello world hello world y hello world hello world x hello world hello world x hello world hello world x hello world hello world x hello world hello world x hello world hello world x hello world hello world x hello world hello world x hello world");
    }

}

