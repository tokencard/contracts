/**
 * The MIT License (MIT)
 *
 * Copyright (c) 2016 Smart Contract Solutions, Inc.
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

 /**
  * @title SafeMath
  * @dev Unsigned math operations with safety checks that revert on error.
  */
 library SafeMath {
     /**
      * @dev Multiplies two unsigned integers, reverts on overflow.
      */
     function mul(uint256 a, uint256 b) internal pure returns (uint256) {
         // Gas optimization: this is cheaper than requiring 'a' not being zero, but the
         // benefit is lost if 'b' is also tested.
         // See: https://github.com/OpenZeppelin/openzeppelin-solidity/pull/522
         if (a == 0) {
             return 0;
         }

         uint256 c = a * b;
         require(c / a == b, "SafeMath: multiplication overflow");

         return c;
     }

     /**
      * @dev Integer division of two unsigned integers truncating the quotient, reverts on division by zero.
      */
     function div(uint256 a, uint256 b) internal pure returns (uint256) {
         // Solidity only automatically asserts when dividing by 0
         require(b > 0, "SafeMath: division by zero");
         uint256 c = a / b;
         // assert(a == b * c + a % b); // There is no case in which this doesn't hold

         return c;
     }

     /**
      * @dev Subtracts two unsigned integers, reverts on overflow (i.e. if subtrahend is greater than minuend).
      */
     function sub(uint256 a, uint256 b) internal pure returns (uint256) {
         require(b <= a, "SafeMath: subtraction overflow");
         uint256 c = a - b;

         return c;
     }

     /**
      * @dev Adds two unsigned integers, reverts on overflow.
      */
     function add(uint256 a, uint256 b) internal pure returns (uint256) {
         uint256 c = a + b;
         require(c >= a, "SafeMath: addition overflow");

         return c;
     }

     /**
      * @dev Divides two unsigned integers and returns the remainder (unsigned integer modulo),
      * reverts when dividing by zero.
      */
     function mod(uint256 a, uint256 b) internal pure returns (uint256) {
         require(b != 0, "SafeMath: modulo by zero");
         return a % b;
     }
 }
