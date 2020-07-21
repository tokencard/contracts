// SPDX-License-Identifier: GPLv3

pragma solidity ^0.6.11;

import "../externals/SafeMath.sol";


contract GasToken {
    using SafeMath for uint256;

    uint256 public totalMinted;
    uint256 public totalBurned;

    mapping(address => uint256) private _balances;

    function balanceOf(address account) public view returns (uint256) {
        return _balances[account];
    }

    function totalSupply() public view returns (uint256) {
        return totalMinted - totalBurned;
    }

    function mint(uint256 value) public {
        uint256 offset = totalMinted;
        assembly {
            mstore(0, 0x766ffa233a79675b0530301caf58abcfa2eb3318585733ff60005260176009f3)

            for {
                let i := div(value, 32)
            } i {
                i := sub(i, 1)
            } {
                pop(create2(0, 0, 32, add(offset, 0)))
                pop(create2(0, 0, 32, add(offset, 1)))
                pop(create2(0, 0, 32, add(offset, 2)))
                pop(create2(0, 0, 32, add(offset, 3)))
                pop(create2(0, 0, 32, add(offset, 4)))
                pop(create2(0, 0, 32, add(offset, 5)))
                pop(create2(0, 0, 32, add(offset, 6)))
                pop(create2(0, 0, 32, add(offset, 7)))
                pop(create2(0, 0, 32, add(offset, 8)))
                pop(create2(0, 0, 32, add(offset, 9)))
                pop(create2(0, 0, 32, add(offset, 10)))
                pop(create2(0, 0, 32, add(offset, 11)))
                pop(create2(0, 0, 32, add(offset, 12)))
                pop(create2(0, 0, 32, add(offset, 13)))
                pop(create2(0, 0, 32, add(offset, 14)))
                pop(create2(0, 0, 32, add(offset, 15)))
                pop(create2(0, 0, 32, add(offset, 16)))
                pop(create2(0, 0, 32, add(offset, 17)))
                pop(create2(0, 0, 32, add(offset, 18)))
                pop(create2(0, 0, 32, add(offset, 19)))
                pop(create2(0, 0, 32, add(offset, 20)))
                pop(create2(0, 0, 32, add(offset, 21)))
                pop(create2(0, 0, 32, add(offset, 22)))
                pop(create2(0, 0, 32, add(offset, 23)))
                pop(create2(0, 0, 32, add(offset, 24)))
                pop(create2(0, 0, 32, add(offset, 25)))
                pop(create2(0, 0, 32, add(offset, 26)))
                pop(create2(0, 0, 32, add(offset, 27)))
                pop(create2(0, 0, 32, add(offset, 28)))
                pop(create2(0, 0, 32, add(offset, 29)))
                pop(create2(0, 0, 32, add(offset, 30)))
                pop(create2(0, 0, 32, add(offset, 31)))
                offset := add(offset, 32)
            }

            for {
                let i := and(value, 0x1F)
            } i {
                i := sub(i, 1)
            } {
                pop(create2(0, 0, 32, offset))
                offset := add(offset, 1)
            }
        }

        _mint(msg.sender, value);
        totalMinted = offset;
    }

    function _destroyChildren(uint256 value) internal {
        assembly {
            let i := sload(totalBurned_slot)
            let end := add(i, value)
            sstore(totalBurned_slot, end)

            let data := mload(0x40)
            mstore(data, 0xff00000000fa233a79675b0530301caf58abcfa2eb0000000000000000000000)
            mstore(add(data, 53), 0x841da0d3b4b49d75c2a11068e21bceeb2e5d8c9e31ab7cea45c9ce114a2033dc)
            let ptr := add(data, 21)
            for {

            } lt(i, end) {
                i := add(i, 1)
            } {
                mstore(ptr, i)
                pop(call(gas(), keccak256(data, 85), 0, 0, 0, 0, 0))
            }
        }
    }

    function free(uint256 value) public returns (uint256) {
        if (value > 0) {
            _burn(msg.sender, value);
            _destroyChildren(value);
        }
        return value;
    }

    function freeUpTo(uint256 value) public returns (uint256) {
        return free(_min(value, balanceOf(msg.sender)));
    }

    function transfer(address recipient, uint256 amount) public {
        _balances[msg.sender] = _balances[msg.sender].sub(amount);
        _balances[recipient] = _balances[recipient].add(amount);
    }

    function _min(uint256 a, uint256 b) private pure returns (uint256) {
        return a < b ? a : b;
    }

    function _mint(address account, uint256 amount) private {
        _balances[account] = _balances[account].add(amount);
    }

    function _burn(address account, uint256 amount) private {
        _balances[account] = _balances[account].sub(amount);
    }
}
