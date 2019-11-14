pragma solidity ^0.5.10;

import "contracts/wallet.sol";
import "contracts/internals/controllable.sol";

contract CryticInterface{
    address internal crytic_owner = address(0x41414141);
    address internal crytic_user = address(0x42424242);
    address internal crytic_attacker = address(0x43434343);
}

contract TEST is CryticInterface, AddressWhitelist {

    address[] added_addresses;
    address[] removed_addresses;

    constructor() Ownable(crytic_user, true) Controllable(bytes32(0x0)) ENSResolvable(address(0x0)) public {
	    //address[] storage empty;
	    //empty.push(crytic_attacker);
	    //this.setWhitelist(empty);
    }

    modifier onlyController() {
        require (msg.sender == crytic_user || msg.sender == address(this), "");
        _;
    }

    function crytic_hasNoOwnerOrZeroAddress() public returns (bool) {
        for (uint i = 0; i < whitelistArray.length; i++) {
            if (_isOwner(whitelistArray[i])) {
                return false;
            }
            if (whitelistArray[i] == address(0)) {
                return false;
            }
        }
        return true;
    }

    function crytic_Whitelist_addition_length() public returns (bool) {
        if (!submittedWhitelistAddition) {
            return true;
        }

        uint length_before = whitelistArray.length;
        this.confirmWhitelistAddition(calculateHash(_pendingWhitelistAddition));
        uint length_after = whitelistArray.length;
        return (length_before < length_after);
    }

    function crytic_Whitelist_addition_map() public returns (bool) {
        if (!submittedWhitelistAddition) {
            return true;
        }

        for (uint i = 0; i < _pendingWhitelistAddition.length; i++) {
            added_addresses.push(_pendingWhitelistAddition[i]);
        }

        this.confirmWhitelistAddition(calculateHash(_pendingWhitelistAddition));

        bool checked_addresses = true;
        for (i = 0; i < added_addresses.length; i++) {
            checked_addresses = checked_addresses && whitelistMap[added_addresses[i]];
        }

        return checked_addresses;
    }

    function crytic_Whitelist_removal_length() public returns (bool) {
        if (!submittedWhitelistRemoval) {
            return true;
        }

        uint length_before = whitelistArray.length;
        this.confirmWhitelistRemoval(calculateHash(_pendingWhitelistRemoval));
        uint length_after = whitelistArray.length;
        return (length_before >= length_after);
    }

    function crytic_Whitelist_removal_map() public returns (bool) {
        if (!submittedWhitelistRemoval) {
            return true;
        }

        for (uint i = 0; i < _pendingWhitelistRemoval.length; i++) {
            removed_addresses.push(_pendingWhitelistRemoval[i]);
        }

        this.confirmWhitelistRemoval(calculateHash(_pendingWhitelistRemoval));

        bool checked_addresses = true;
        for (i = 0; i < removed_addresses.length; i++) {
            checked_addresses = checked_addresses && !whitelistMap[removed_addresses[i]];
        }

        return checked_addresses;
    }
}
