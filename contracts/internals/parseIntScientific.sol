/**
 *  ParseIntScientific - The Consumer Contract Wallet
 *  Copyright (C) 2019 The Contract Wallet Company Limited
 *
 *  This program is free software: you can redistribute it and/or modify
 *  it under the terms of the GNU General Public License as published by
 *  the Free Software Foundation, either version 3 of the License, or
 *  (at your option) any later version.

 *  This program is distributed in the hope that it will be useful,
 *  but WITHOUT ANY WARRANTY; without even the implied warranty of
 *  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *  GNU General Public License for more details.

 *  You should have received a copy of the GNU General Public License
 *  along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

pragma solidity ^0.4.25;

import "../externals/SafeMath.sol";


/// @title ParseIntScientific provides floating point in scientific notation (e.g. e-5) parsing functionality.
contract ParseIntScientific {

    using SafeMath for uint256;

    byte constant private _PLUS_ASCII = byte(43); //decimal value of '+'
    byte constant private _DASH_ASCII = byte(45); //decimal value of '-'
    byte constant private _DOT_ASCII = byte(46); //decimal value of '.'
    byte constant private _ZERO_ASCII = byte(48); //decimal value of '0'
    byte constant private _NINE_ASCII = byte(57); //decimal value of '9'
    byte constant private _E_ASCII = byte(69); //decimal value of 'E'
    byte constant private _LOWERCASE_E_ASCII = byte(101); //decimal value of 'e'

    /// @notice ParseIntScientific delegates the call to _parseIntScientific(string, uint) with the 2nd argument being 0.
    function _parseIntScientific(string _inString) internal pure returns (uint) {
        return _parseIntScientific(_inString, 0);
    }

    /// @notice ParseIntScientificWei parses a rate expressed in ETH and returns its wei denomination
    function _parseIntScientificWei(string _inString) internal pure returns (uint) {
        return _parseIntScientific(_inString, 18);
    }

    /// @notice ParseIntScientific parses a JSON standard - floating point number.
    /// @param _inString is input string.
    /// @param _magnitudeMult multiplies the number with 10^_magnitudeMult.
    function _parseIntScientific(string _inString, uint _magnitudeMult) internal pure returns (uint) {

        bytes memory inBytes = bytes(_inString);
        uint mint = 0; // the final uint returned
        uint mintDec = 0; // the uint following the decimal point
        uint mintExp = 0; // the exponent
        uint decMinted = 0; // how many decimals were 'minted'.
        uint expIndex = 0; // the position in the byte array that 'e' was found (if found)
        bool integral = false; // indicates the existence of the integral part, it should always exist (even if 0) e.g. 'e+1'  or '.1' is not valid
        bool decimals = false; // indicates a decimal number, set to true if '.' is found
        bool exp = false; // indicates if the number being parsed has an exponential representation
        bool minus = false; // indicated if the exponent is negative
        bool plus = false; // indicated if the exponent is positive

        for (uint i = 0; i < inBytes.length; i++) {
            if ((inBytes[i] >= _ZERO_ASCII) && (inBytes[i] <= _NINE_ASCII) && (!exp)) {
                // 'e' not encountered yet, minting integer part or decimals
                if (decimals) {
                    // '.' encountered
                    // use safeMath in case there is an overflow
                    mintDec = mintDec.mul(10);
                    mintDec = mintDec.add(uint(inBytes[i]) - uint(_ZERO_ASCII));
                    decMinted++; //keep track of the #decimals
                } else {
                    // integral part (before '.')
                    integral = true;
                    // use safeMath in case there is an overflow
                    mint = mint.mul(10);
                    mint = mint.add(uint(inBytes[i]) - uint(_ZERO_ASCII));
                }
            } else if ((inBytes[i] >= _ZERO_ASCII) && (inBytes[i] <= _NINE_ASCII) && (exp)) {
                //exponential notation (e-/+) has been detected, mint the exponent
                mintExp = mintExp.mul(10);
                mintExp = mintExp.add(uint(inBytes[i]) - uint(_ZERO_ASCII));
            } else if (inBytes[i] == _DOT_ASCII) {
                //an integral part before should always exist before '.'
                require(integral, "missing integral part");
                // an extra decimal point makes the format invalid
                require(!decimals, "duplicate decimal point");
                //the decimal point should always be before the exponent
                require(!exp, "decimal after exponent");
                decimals = true;
            } else if (inBytes[i] == _DASH_ASCII) {
                // an extra '-' should be considered an invalid character
                require(!minus, "duplicate -");
                require(!plus, "extra sign");
                require(expIndex + 1 == i, "- sign not immediately after e");
                minus = true;
            } else if (inBytes[i] == _PLUS_ASCII) {
                // an extra '+' should be considered an invalid character
                require(!plus, "duplicate +");
                require(!minus, "extra sign");
                require(expIndex + 1 == i, "+ sign not immediately after e");
                plus = true;
            } else if ((inBytes[i] == _E_ASCII) || (inBytes[i] == _LOWERCASE_E_ASCII)) {
                //an integral part before should always exist before 'e'
                require(integral, "missing integral part");
                // an extra 'e' or 'E' should be considered an invalid character
                require(!exp, "duplicate exponent symbol");
                exp = true;
                expIndex = i;
            } else {
                revert("invalid digit");
            }
        }

        if (minus || plus) {
            // end of string e[x|-] without specifying the exponent
            require(i > expIndex + 2);
        } else if (exp) {
            // end of string (e) without specifying the exponent
            require(i > expIndex + 1);
        }

        if (minus) {
            // e^(-x)
            if (mintExp >= _magnitudeMult) {
                // the (negative) exponent is bigger than the given parameter for "shifting left".
                // use integer division to reduce the precision.
                require(mintExp - _magnitudeMult < 78, "exponent > 77"); //
                mint /= 10 ** (mintExp - _magnitudeMult);
                return mint;

            } else {
                // the (negative) exponent is smaller than the given parameter for "shifting left".
                //no need for underflow check
                _magnitudeMult = _magnitudeMult - mintExp;
            }
        } else {
            // e^(+x), positive exponent or no exponent
            // just shift left as many times as indicated by the exponent and the shift parameter
            _magnitudeMult = _magnitudeMult.add(mintExp);
        }

        if (_magnitudeMult >= decMinted) {
            // the decimals are fewer or equal than the shifts: use all of them
            // shift number and add the decimals at the end
            // include decimals if present in the original input
            require(decMinted < 78, "more than 77 decimal digits parsed"); //
            mint = mint.mul(10 ** (decMinted));
            mint = mint.add(mintDec);
            //// add zeros at the end if the decimals were fewer than #_magnitudeMult
            require(_magnitudeMult - decMinted < 78, "exponent > 77"); //
            mint = mint.mul(10 ** (_magnitudeMult - decMinted));
        } else {
            // the decimals are more than the #_magnitudeMult shifts
            // use only the ones needed, discard the rest
            decMinted -= _magnitudeMult;
            require(decMinted < 78, "more than 77 decimal digits parsed"); //
            mintDec /= 10 ** (decMinted);
            // shift number and add the decimals at the end
            require(_magnitudeMult < 78, "more than 77 decimal digits parsed"); //
            mint = mint.mul(10 ** (_magnitudeMult));
            mint = mint.add(mintDec);
        }

        return mint;
    }
}
