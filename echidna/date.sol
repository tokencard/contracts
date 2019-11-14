pragma solidity ^0.5.10;

import "contracts/internals/date.sol";
import "contracts/internals/parseIntScientific.sol";

import "contracts/externals/strings.sol";


contract EchidnaInterface{
    address internal echidna_owner = address(0x41414141);
    address internal echidna_user = address(0x42424242);
    address internal echidna_attacker = address(0x43434343);
}

contract TEST is EchidnaInterface, Date, ParseIntScientific {

    using strings for *;

    string example = "12_Sep_2018_15:18:14";
    string input = example;
    bool bool_out;
    uint timestamp_out;

    function input_date(string memory _date, uint _lastUpdate) public {
        input = _date;
        _verifyDate(_date, _lastUpdate);
    }

    function _verifyDate(string memory _dateHeader, uint _lastUpdate) private pure returns (bool, uint) {

        // called by verifyProof(), _dateHeader is always a string of length = 20
        assert(abi.encodePacked(_dateHeader).length == 20);

        // Split the date string and get individual date components.
        strings.slice memory date = _dateHeader.toSlice();
        strings.slice memory timeDelimiter = ":".toSlice();
        strings.slice memory dateDelimiter = "_".toSlice();

        uint day = _parseIntScientific(date.split(dateDelimiter).toString());
        require(day > 0 && day < 32, "day error");

        uint month = _monthToNumber(date.split(dateDelimiter).toString());
        require(month > 0 && month < 13, "month error");

        uint year = _parseIntScientific(date.split(dateDelimiter).toString());
        require(year > 2017 && year < 3000, "year error");

        uint hour = _parseIntScientific(date.split(timeDelimiter).toString());
        require(hour < 25, "hour error");

        uint minute = _parseIntScientific(date.split(timeDelimiter).toString());
        require(minute < 60, "minute error");

        uint second = _parseIntScientific(date.split(timeDelimiter).toString());
        require(second < 60, "second error");

        uint timestamp = year * (10 ** 10) + month * (10 ** 8) + day * (10 ** 6) + hour * (10 ** 4) + minute * (10 ** 2) + second;

        return (timestamp > _lastUpdate, timestamp);
    }

    function compare(string memory s1, string memory s2) internal returns (bool) {
        return keccak256(abi.encodePacked(s1)) == keccak256(abi.encodePacked(s2));
    }

    function echidna_valid_timestamp() public returns (bool) {
        return compare(input, example);
    }
}
