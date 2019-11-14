/* This test requires echidna with hevm-0.29 (branch `dev-hevm-0.29`) */
pragma solidity ^0.5.10;

contract TEST {
    bytes32 hash;
    uint8 v;
    bytes32 r;
    bytes32 s;
    address raddr1;
    address raddr2;
    bool rbool;

    function input(bytes32 _hash, uint8 _v, bytes32 _r, bytes32 _s) public {
        hash = _hash;
        v = _v;
        r = _r;
        s = _s;
        raddr1 = ecrecover(hash, v, r, s);
        (rbool, raddr2) = safer_ecrecover(hash, v, r, s);
    }

    function safer_ecrecover(bytes32 _hash, uint8 _v, bytes32 _r, bytes32 _s) public returns (bool, address) {
        // We do our own memory management here. Solidity uses memory offset
        // 0x40 to store the current end of memory. We write past it (as
        // writes are memory extensions), but don't update the offset so
        // Solidity will reuse it. The memory used here is only needed for
        // this context.

        // FIXME: inline assembly can't access return values
        bool ret;
        address addr;

        assembly {
            let size := mload(0x40)
            mstore(size, _hash)
            mstore(add(size, 32), _v)
            mstore(add(size, 64), _r)
            mstore(add(size, 96), _s)

            // NOTE: we can reuse the request memory because we deal with
            //       the return code
            ret := call(3000, 1, 0, size, 128, size, 32)
            addr := mload(size)
        }

        return (ret, addr);
    }

    function crytic_ecrecovers_equal() public view returns (bool) {
        return (raddr1 == raddr2);
    }
}