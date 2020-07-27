pragma solidity ^0.5.17;

import "../externals/SafeMath.sol";

interface IGasToken {
    function free(uint256 value) external returns (uint256);
    function freeUpTo(uint256 value) external returns (uint256);
}

contract GasToken is IGasToken {
	using SafeMath for uint256;

    uint256 public totalMinted;
    uint256 public totalBurned;

    mapping(address => uint256) private _balances;

    function balanceOf(address account) public view returns (uint256) {
        return _balances[account];
    }

    function totalSupply() public view returns(uint256) {
        return totalMinted.sub(totalBurned);
    }

    function mint(uint256 value) public {
        uint256 offset = totalMinted;
        assembly {
            mstore(0, 0x746d4946c0e9F43F4Dee607b0eF1fA1c3318585733ff6000526015600bf30000)

            for {let i := div(value, 32)} i {i := sub(i, 1)} {
                pop(create2(0, 0, 30, add(offset, 0))) pop(create2(0, 0, 30, add(offset, 1)))
                pop(create2(0, 0, 30, add(offset, 2))) pop(create2(0, 0, 30, add(offset, 3)))
                pop(create2(0, 0, 30, add(offset, 4))) pop(create2(0, 0, 30, add(offset, 5)))
                pop(create2(0, 0, 30, add(offset, 6))) pop(create2(0, 0, 30, add(offset, 7)))
                pop(create2(0, 0, 30, add(offset, 8))) pop(create2(0, 0, 30, add(offset, 9)))
                pop(create2(0, 0, 30, add(offset, 10))) pop(create2(0, 0, 30, add(offset, 11)))
                pop(create2(0, 0, 30, add(offset, 12))) pop(create2(0, 0, 30, add(offset, 13)))
                pop(create2(0, 0, 30, add(offset, 14))) pop(create2(0, 0, 30, add(offset, 15)))
                pop(create2(0, 0, 30, add(offset, 16))) pop(create2(0, 0, 30, add(offset, 17)))
                pop(create2(0, 0, 30, add(offset, 18))) pop(create2(0, 0, 30, add(offset, 19)))
                pop(create2(0, 0, 30, add(offset, 20))) pop(create2(0, 0, 30, add(offset, 21)))
                pop(create2(0, 0, 30, add(offset, 22))) pop(create2(0, 0, 30, add(offset, 23)))
                pop(create2(0, 0, 30, add(offset, 24))) pop(create2(0, 0, 30, add(offset, 25)))
                pop(create2(0, 0, 30, add(offset, 26))) pop(create2(0, 0, 30, add(offset, 27)))
                pop(create2(0, 0, 30, add(offset, 28))) pop(create2(0, 0, 30, add(offset, 29)))
                pop(create2(0, 0, 30, add(offset, 30))) pop(create2(0, 0, 30, add(offset, 31)))
                offset := add(offset, 32)
            }

            for {let i := and(value, 0x1F)} i {i := sub(i, 1)} {
                pop(create2(0, 0, 30, offset))
                offset := add(offset, 1)
            }
        }

        _mint(msg.sender, value);
        totalMinted = offset;
    }

    function computeAddress2(uint256 salt) public view returns (address) {
        bytes32 _data = keccak256(
            abi.encodePacked(bytes1(0xff), address(this), salt, bytes32(0x3c1644c68e5d6cb380c36d1bf847fdbc0c7ac28030025a2fc5e63cce23c16348))
        );
        return address(uint256(_data));
    }

    function _destroyChildren(uint256 value) private {
        uint256 _totalBurned = totalBurned;
        for (uint256 i = 0; i < value; i++) {
            computeAddress2(_totalBurned + i).call("");
        }
        totalBurned = _totalBurned + value;
    }

    function free(uint256 value) public returns (uint256)  {
        _burn(msg.sender, value);
        _destroyChildren(value);
        return value;
    }

    function freeUpTo(uint256 value) public returns (uint256) {
        return free(min(value, balanceOf(msg.sender)));
    }

    function transfer(address recipient, uint256 amount) public {
        _balances[msg.sender] = _balances[msg.sender].sub(amount);
        _balances[recipient] = _balances[recipient].add(amount);
    }

    function min(uint256 a, uint256 b) private pure returns (uint256) {
        return a < b ? a : b;
    }

    function _mint(address account, uint256 amount) private {
        _balances[account] = _balances[account].add(amount);
    }

    function _burn(address account, uint256 amount) private {
        _balances[account] = _balances[account].sub(amount);
    }
}
