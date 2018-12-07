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

/// @title Date provides date parsing functionality.
contract Date {

    bytes32 constant private JANUARY = keccak256("Jan");
    bytes32 constant private FEBRUARY = keccak256("Feb");
    bytes32 constant private MARCH = keccak256("Mar");
    bytes32 constant private APRIL = keccak256("Apr");
    bytes32 constant private MAY = keccak256("May");
    bytes32 constant private JUNE = keccak256("Jun");
    bytes32 constant private JULY = keccak256("Jul");
    bytes32 constant private AUGUST = keccak256("Aug");
    bytes32 constant private SEPTEMBER = keccak256("Sep");
    bytes32 constant private OCTOBER = keccak256("Oct");
    bytes32 constant private NOVEMBER = keccak256("Nov");
    bytes32 constant private DECEMBER = keccak256("Dec");

    /// @return the number of the month based on its name.
    /// @param _month the first three letters of a month's name e.g. "Jan".
    function monthToNumber(string _month) internal pure returns (uint8) {
        bytes32 month = keccak256(abi.encodePacked(_month));
        if (month == JANUARY) {
            return 1;
        } else if (month == FEBRUARY) {
            return 2;
        } else if (month == MARCH) {
            return 3;
        } else if (month == APRIL) {
            return 4;
        } else if (month == MAY) {
            return 5;
        } else if (month == JUNE) {
            return 6;
        } else if (month == JULY) {
            return 7;
        } else if (month == AUGUST) {
            return 8;
        } else if (month == SEPTEMBER) {
            return 9;
        } else if (month == OCTOBER) {
            return 10;
        } else if (month == NOVEMBER) {
            return 11;
        } else if (month == DECEMBER) {
            return 12;
        } else {
            revert("not a valid month");
        }
    }
}
