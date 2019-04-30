/**
 * The MIT License (MIT)
 *
 * Copyright (c) 2016-2019 zOS Global Limited
 *
 * Permission is hereby granted, free of charge, to any person obtaining
 * a copy of this software and associated documentation files (the
 * "Software"), to deal in the Software without restriction, including
 * without limitation the rights to use, copy, modify, merge, publish,
 * distribute, sublicense, and/or sell copies of the Software, and to
 * permit persons to whom the Software is furnished to do so, subject to
 * the following conditions:
 *
 * The above copyright notice and this permission notice shall be included
 * in all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS
 * OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
 * MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
 * IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY
 * CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,
 * TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE
 * SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
 */

pragma solidity ^0.4.25;

import "./ERC20.sol";
import "./SafeMath.sol";
import "./Address.sol";

/**
 * @title SafeERC20
 * @dev Wrappers around ERC20 operations that throw on failure (when the token
 * contract returns false). Tokens that return no value (and instead revert or
 * throw on failure) are also supported, non-reverting calls are assumed to be
 * successful.
 * To use this library you can add a `using SafeERC20 for ERC20;` statement to your contract,
 * which allows you to call the safe operations as `token.safeTransfer(...)`, etc.
 * This is a modified version of the OpenZepplin SafeERC20 ^v0.5.0
 * https://github.com/OpenZeppelin/openzeppelin-solidity/blob/master/contracts/token/ERC20/SafeERC20.sol
 */

library SafeERC20 {
    using SafeMath for uint256;
    using Address for address;

    function safeTransfer(ERC20 _token, address _to, uint256 _value) internal {
        callOptionalReturn(_token, abi.encodeWithSignature("transfer(address,uint256)", _to, _value));
    }

    function safeTransferFrom(ERC20 _token, address _from, address _to, uint256 _value) internal {
        callOptionalReturn(_token, abi.encodeWithSignature("transferFrom(address,address,uint256)", _from, _to, _value));
    }

    function safeApprove(ERC20 _token, address _spender, uint256 _value) internal {
        // safeApprove should only be called when setting an initial allowance,
        // or when resetting it to zero. To increase and decrease it, use
        // 'safeIncreaseAllowance' and 'safeDecreaseAllowance'
        // solhint-disable-next-line max-line-length
        require((_value == 0) || (_token.allowance(address(this), _spender) == 0), "SafeERC20: approve from non-zero to non-zero allowance");
        callOptionalReturn(_token, abi.encodeWithSignature("approve(address,uint256)", _spender, _value));
    }

    function safeIncreaseAllowance(ERC20 _token, address _spender, uint256 _value) internal {
        uint256 newAllowance = _token.allowance(address(this), _spender).add(_value);
        callOptionalReturn(_token, abi.encodeWithSignature("approve(address,uint256)", _spender, newAllowance));
    }

    function safeDecreaseAllowance(ERC20 _token, address _spender, uint256 _value) internal {
        uint256 newAllowance = _token.allowance(address(this), _spender).sub(_value);
        callOptionalReturn(_token, abi.encodeWithSignature("approve(address,uint256)", _spender, newAllowance));
    }

    /**
     * @dev Imitates a Solidity high-level call (i.e. a regular function call to a contract), relaxing the requirement
     * on the return value: the return value is optional (but if data is returned, it must not be false).
     * @param _token The token targeted by the call.
     * @param _data The call data (encoded using abi.encode or one of its variants).
     */
     function callOptionalReturn(ERC20 _token, bytes memory _data) private {
         // We need to perform a low level call here, to bypass Solidity's return data size checking mechanism, since
         // we're implementing it ourselves.

         // A Solidity high level call has three parts:
         //  1. The target address is checked to verify it contains contract code
         //  2. The call itself is made, and success asserted
         //  3. The return value is decoded, which in turn checks the size of the returned data.
         // solhint-disable-next-line max-line-length
         require(address(_token).isContract(), "SafeERC20: call to non-contract");
         // solhint-disable-next-line avoid-low-level-calls
         require(address(_token).call(_data), "SafeERC20: low-level call failed");
         /* address tokenAddress = address(token);
         uint dataLength = data.length; */

         assembly {
             /* if iszero(call(gas, tokenAddress, 0, add(data, 0x20), dataLength, 0, 0)) { revert(0, 0) } */
             let size := returndatasize
             switch size
             case 0 {} //non-compliant with the ERC20 interface, ignore
             case 32 {
                 let ptr := mload(0x40)
                 returndatacopy(ptr, 0, 32)
                 let rvalue := mload(ptr)
                 if iszero(rvalue) { revert(0, 0) }
             }
             default { revert(0, 0) }
             }
     }
}
