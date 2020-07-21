/**
 *  Date - The Consumer Contract Wallet
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

// SPDX-License-Identifier: GPLv3

pragma solidity ^0.6.11;

/// @title Date provides redimentary date parsing functionality.
/// @notice This method parses months found in an ISO date to a number
contract Date {
    bytes32 private constant _JANUARY = keccak256("Jan");
    bytes32 private constant _FEBRUARY = keccak256("Feb");
    bytes32 private constant _MARCH = keccak256("Mar");
    bytes32 private constant _APRIL = keccak256("Apr");
    bytes32 private constant _MAY = keccak256("May");
    bytes32 private constant _JUNE = keccak256("Jun");
    bytes32 private constant _JULY = keccak256("Jul");
    bytes32 private constant _AUGUST = keccak256("Aug");
    bytes32 private constant _SEPTEMBER = keccak256("Sep");
    bytes32 private constant _OCTOBER = keccak256("Oct");
    bytes32 private constant _NOVEMBER = keccak256("Nov");
    bytes32 private constant _DECEMBER = keccak256("Dec");

    /// @return the number of the month based on its name.
    /// @param _month the first three letters of a month's name e.g. "Jan".
    function _monthToNumber(string memory _month) internal pure returns (uint8) {
        bytes32 month = keccak256(abi.encodePacked(_month));
        if (month == _JANUARY) {
            return 1;
        } else if (month == _FEBRUARY) {
            return 2;
        } else if (month == _MARCH) {
            return 3;
        } else if (month == _APRIL) {
            return 4;
        } else if (month == _MAY) {
            return 5;
        } else if (month == _JUNE) {
            return 6;
        } else if (month == _JULY) {
            return 7;
        } else if (month == _AUGUST) {
            return 8;
        } else if (month == _SEPTEMBER) {
            return 9;
        } else if (month == _OCTOBER) {
            return 10;
        } else if (month == _NOVEMBER) {
            return 11;
        } else if (month == _DECEMBER) {
            return 12;
        } else {
            revert("not a valid month");
        }
    }
}
