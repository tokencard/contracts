pragma solidity ^0.5.10;

import "contracts/externals/strings.sol";

contract EchidnaInterface{
    address internal echidna_owner = address(0x41414141);
    address internal echidna_user = address(0x42424242);
    address internal echidna_attacker = address(0x43434343);
}

contract TEST is EchidnaInterface {

    using strings for *;

    strings.slice s1 = "init".toSlice();

    function input(string memory x) internal {
        if (strings.equals(s1, "init".toSlice()))
            s1 = x.toSlice();
    }

    function concat(string memory x, string memory x1)  public {
        input(x);
        s1 = strings.concat(s1, x1.toSlice()).toSlice();
    }

    function beyond(string memory x, string memory x1)  public {
        input(x);
        s1 = strings.beyond(s1, x1.toSlice());
    }

    function until(string memory x, string memory x1)  public {
        input(x);
        s1 = strings.until(s1, x1.toSlice());
    }

    function find(string memory x, string memory x1)  public {
        input(x);
        s1 = strings.find(s1, x1.toSlice());
    }

    function rfind(string memory x, string memory x1)  public {
        input(x);
        s1 = strings.rfind(s1, x1.toSlice());
    }

    function split(string memory x, string memory x1, string memory x2)  public {
        input(x);
        s1 = strings.split(s1, x1.toSlice(), x2.toSlice());
    }

    function echidna_keccak() public view returns (bool) {
        return s1.keccak() == keccak256(abi.encodePacked(s1.toString()));
    }

    function echidna_length() public view returns (bool) {
        return s1.len() == bytes(s1.toString()).length;
    }

    function echidna_equality() public view returns (bool) {
        return s1.equals(s1);
    }

}