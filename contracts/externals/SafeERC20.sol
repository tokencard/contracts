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

 pragma solidity ^0.5.0;

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
  */
 library SafeERC20 {
     using SafeMath for uint256;
     using Address for address;

     function safeTransfer(ERC20 token, address to, uint256 value) internal {
         callOptionalReturn(token, abi.encodeWithSelector(token.transfer.selector, to, value));
     }

     function safeTransferFrom(ERC20 token, address from, address to, uint256 value) internal {
         callOptionalReturn(token, abi.encodeWithSelector(token.transferFrom.selector, from, to, value));
     }

     function safeApprove(ERC20 token, address spender, uint256 value) internal {
         // safeApprove should only be called when setting an initial allowance,
         // or when resetting it to zero. To increase and decrease it, use
         // 'safeIncreaseAllowance' and 'safeDecreaseAllowance'
         // solhint-disable-next-line max-line-length
         require((value == 0) || (token.allowance(address(this), spender) == 0),
             "SafeERC20: approve from non-zero to non-zero allowance"
         );
         callOptionalReturn(token, abi.encodeWithSelector(token.approve.selector, spender, value));
     }

     function safeIncreaseAllowance(ERC20 token, address spender, uint256 value) internal {
         uint256 newAllowance = token.allowance(address(this), spender).add(value);
         callOptionalReturn(token, abi.encodeWithSelector(token.approve.selector, spender, newAllowance));
     }

     function safeDecreaseAllowance(ERC20 token, address spender, uint256 value) internal {
         uint256 newAllowance = token.allowance(address(this), spender).sub(value);
         callOptionalReturn(token, abi.encodeWithSelector(token.approve.selector, spender, newAllowance));
     }

     /**
      * @dev Imitates a Solidity high-level call (i.e. a regular function call to a contract), relaxing the requirement
      * on the return value: the return value is optional (but if data is returned, it must not be false).
      * @param token The token targeted by the call.
      * @param data The call data (encoded using abi.encode or one of its variants).
      */
     function callOptionalReturn(ERC20 token, bytes memory data) private {
         // We need to perform a low level call here, to bypass Solidity's return data size checking mechanism, since
         // we're implementing it ourselves.

         // A Solidity high level call has three parts:
         //  1. The target address is checked to verify it contains contract code
         //  2. The call itself is made, and success asserted
         //  3. The return value is decoded, which in turn checks the size of the returned data.
         // solhint-disable-next-line max-line-length
         require(address(token).isContract(), "SafeERC20: call to non-contract");

         // solhint-disable-next-line avoid-low-level-calls
         (bool success, bytes memory returndata) = address(token).call(data);
         require(success, "SafeERC20: low-level call failed");

         if (returndata.length > 0) { // Return data is optional
             // solhint-disable-next-line max-line-length
             require(abi.decode(returndata, (bool)), "SafeERC20: ERC20 operation did not succeed");
         }
     }
 }
