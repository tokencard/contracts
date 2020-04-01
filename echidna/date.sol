pragma solidity ^0.5.15;

import "contracts/internals/date.sol";
import "contracts/internals/parseIntScientific.sol";
import "contracts/externals/strings.sol";


contract Echidna {
    address internal echidna_deployer = address(0x1);
    address internal echidna_attacker = address(0x2);
    address internal echidna_owner = address(0x3);
}


contract TEST is Echidna, Date, ParseIntScientific {
    using strings for *;

    string example = "12_Sep_2018_15:18:14";
    string input = example;
    bool bool_out;
    uint256 timestamp_out;

    function input_date(string memory _date, uint256 _lastUpdate) public {
        input = _date;
        _verifyDate(_date, _lastUpdate);
    }

    function _verifyDate(string memory _dateHeader, uint256 _lastUpdate) private pure returns (bool, uint256) {
        // called by verifyProof(), _dateHeader is always a string of length = 20
        assert(abi.encodePacked(_dateHeader).length == 20);

        // Split the date string and get individual date components.
        strings.slice memory date = _dateHeader.toSlice();
        strings.slice memory timeDelimiter = ":".toSlice();
        strings.slice memory dateDelimiter = "_".toSlice();

        uint256 day = _parseIntScientific(date.split(dateDelimiter).toString());
        require(day > 0 && day < 32, "day error");

        uint256 month = _monthToNumber(date.split(dateDelimiter).toString());
        require(month > 0 && month < 13, "month error");

        uint256 year = _parseIntScientific(date.split(dateDelimiter).toString());
        require(year > 2017 && year < 3000, "year error");

        uint256 hour = _parseIntScientific(date.split(timeDelimiter).toString());
        require(hour < 25, "hour error");

        uint256 minute = _parseIntScientific(date.split(timeDelimiter).toString());
        require(minute < 60, "minute error");

        uint256 second = _parseIntScientific(date.split(timeDelimiter).toString());
        require(second < 60, "second error");

        uint256 timestamp = year * (10**10) + month * (10**8) + day * (10**6) + hour * (10**4) + minute * (10**2) + second;

        return (timestamp > _lastUpdate, timestamp);
    }

    function compare(string memory s1, string memory s2) internal returns (bool) {
        return keccak256(abi.encodePacked(s1)) == keccak256(abi.encodePacked(s2));
    }

    function echidna_valid_timestamp() public returns (bool) {
        return compare(input, example);
    }
}
