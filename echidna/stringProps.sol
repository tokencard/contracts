pragma solidity ^0.5.10;

import "contracts/externals/strings.sol";

contract CryticInterface{
    address internal crytic_owner = address(0x41414141);
    address internal crytic_user = address(0x42424242);
    address internal crytic_attacker = address(0x43434343);
}

contract TEST is CryticInterface {

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

    function crytic_keccak() public view returns (bool) {
        return s1.keccak() == keccak256(abi.encodePacked(s1.toString()));
    }

    function crytic_length() public view returns (bool) {
        return s1.len() == bytes(s1.toString()).length;
    }

    function crytic_equality() public view returns (bool) {
        return s1.equals(s1);
    }

}