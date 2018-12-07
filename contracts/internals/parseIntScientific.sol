/**
 *  The Consumer Contract Wallet
 *  Copyright (C) 2018 Token Group Ltd
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

/// @title ParseIntScientific provides floating point in scientific notation (e.g. e-5) parsing functionality.
contract ParseIntScientific {

    byte constant private PLUS_ASCII = byte(43); //decimal value of '+'
    byte constant private DASH_ASCII = byte(45); //decimal value of '-'
    byte constant private DOT_ASCII = byte(46); //decimal value of '.'
    byte constant private ZERO_ASCII = byte(48); //decimal value of '0'
    byte constant private NINE_ASCII = byte(57); //decimal value of '9'
    byte constant private E_ASCII = byte(69); //decimal value of 'E'
    byte constant private e_ASCII = byte(101); //decimal value of 'e'
  
    /// @dev ParseIntScientific delegates the call to _parseIntScientific(string, uint) with the 2nd argument being 0.
    function _parseIntScientific(string _inString) internal pure returns (uint) {
        return _parseIntScientific(_inString, 0);
    }
  
    /// @dev ParseIntScientificWei parses a rate expressed in ETH and returns its wei denomination
    function _parseIntScientificWei(string _inString) internal pure returns (uint) {
        return _parseIntScientific(_inString, 18);
    }
  
    /// @dev ParseIntScientific parses a JSON standard - floating point number.
    /// @param _inString is input string.
    /// @param _magnitudeMult multiplies the number with 10^_magnitudeMult.
    function _parseIntScientific(string _inString, uint _magnitudeMult) internal pure returns (uint) {
  
        bytes memory inBytes = bytes(_inString);
        uint mint = 0; // the final uint returned
        uint mintDec = 0; // the uint following the decimal point
        uint mintExp = 0; // the exponent
        uint decMinted = 0; // how many decimals were 'minted'.
        uint expIndex = 0; // the position in the byte array that 'e' was found (if found)
        uint shifts; // how many times the final number has to be shifted (left or right) i.e. 10^shifts
        bool decimals = false; // indicates a decimal number, set to true if '.' is found
        bool exp = false; // indicates if the number being parsed has an exponential representation
        bool minus = false; // indicated if the exponent is negative
        bool plus = false; // indicated if the exponent is positive
  
        for (uint i = 0; i < inBytes.length; i++) {
            if ((inBytes[i] >= ZERO_ASCII) && (inBytes[i] <= NINE_ASCII) && (!exp)) {
                // 'e' not encountered yet, minting integer part or decimals
                if (decimals) {
                    // '.' encountered
                    mintDec *= 10;
                    mintDec += uint(inBytes[i]) - uint(ZERO_ASCII);
                    decMinted++; //keep track of how many decimals the input number had
                } else {
                    // integer part (before '.')
                    mint *= 10;
                    mint += uint(inBytes[i]) - uint(ZERO_ASCII);
                }
            } else if ((inBytes[i] >= ZERO_ASCII) && (inBytes[i] <= NINE_ASCII) && (exp)) {
                //exponential notation (e-/+) has been detected, mint the exponent
                mintExp *= 10;
                mintExp += uint(inBytes[i]) - uint(ZERO_ASCII);
            } else if (inBytes[i] == DOT_ASCII) {
                // an extra decimal point makes the format invalid
                require(!decimals, "duplicate decimal point");
                decimals = true;
            } else if (inBytes[i] == DASH_ASCII) {
                // an extra '-' should be considered an invalid character
                require(!minus, "duplicate -");
                require(!plus, "extra sign");
                require(expIndex + 1 == i, "- sign not immediately after e");
                minus = true;
            } else if (inBytes[i] == PLUS_ASCII) {
                // an extra '+' should be considered an invalid character
                require(!plus, "duplicate +");
                require(!minus, "extra sign");
                require(expIndex + 1 == i, "+ sign not immediately after e");
                plus = true;
            } else if ((inBytes[i] == E_ASCII) || (inBytes[i] == e_ASCII)) {
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
                mint /= 10 ** (mintExp - _magnitudeMult);
            } else {
                // the (negative) exponent is smaller than the given parameter for "shifting left".
                shifts = _magnitudeMult - mintExp;
                if (shifts >= decMinted) {
                    // the decimals are fewer or equal than the shifts: use all of them
                    // shift number and add the decimals at the end
                    mint *= 10 ** (decMinted);
                    mint += mintDec;
                    // add zeros at the end if needed
                    mint *= 10 ** (shifts - decMinted);
                } else {
                    // the decimals are more than the shifts
                    // use only the ones needed, discard the rest
                    mintDec /= 10 ** (decMinted-shifts);
                    // shift number and add the decimals at the end
                    mint *= 10 ** (shifts);
                    mint += mintDec;
                }
            }
        } else {
            // e^(+x), positive exponent or no exponent
            // just shift left as many times as indicated by the exponent and the shift parameter
            shifts = _magnitudeMult + mintExp;
            // include decimals if present in the original input
            mint *= 10 ** (decMinted);
            mint += mintDec;
            //'shift' again if the decimals were fewer that the combined shifts
            mint *= 10 ** (shifts - decMinted);
        }
        return mint;
    }
}
