pragma solidity ^0.5.10;

import "crytic-export/flattening/Wallet.sol";

contract EchidnaInterface{
    address internal echidna_owner = address(0x41414141);
    address payable internal echidna_user = address(0x42424242);
    address internal echidna_attacker = address(0x43434343);
}

contract TEST is EchidnaInterface, AddressWhitelist {

    address[] added_addresses;
    address[] removed_addresses;

    constructor() Ownable(echidna_user, true) Controllable(bytes32(0x0)) ENSResolvable(address(0x0)) public {
	    //address[] storage empty;
	    //empty.push(echidna_attacker);
	    //this.setWhitelist(empty);
    }

    modifier onlyController() {
        require (msg.sender == echidna_user || msg.sender == address(this), "");
        _;
    }

    function echidna_hasNoOwnerOrZeroAddress() public returns (bool) {
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

    function echidna_Whitelist_addition_length() public returns (bool) {
        if (!submittedWhitelistAddition) {
            return true;
        }

        uint length_before = whitelistArray.length;
        this.confirmWhitelistAddition(calculateHash(pendingWhitelistAddition()));
        uint length_after = whitelistArray.length;
        return (length_before < length_after);
    }

    function echidna_Whitelist_addition_map() public returns (bool) {
        if (!submittedWhitelistAddition) {
            return true;
        }

        for (uint i = 0; i < pendingWhitelistAddition().length; i++) {
            added_addresses.push(pendingWhitelistAddition()[i]);
        }

        this.confirmWhitelistAddition(calculateHash(pendingWhitelistAddition()));

        bool checked_addresses = true;
        for (uint i = 0; i < added_addresses.length; i++) {
            checked_addresses = checked_addresses && whitelistMap[added_addresses[i]];
        }

        return checked_addresses;
    }

    function echidna_Whitelist_removal_length() public returns (bool) {
        if (!submittedWhitelistRemoval) {
            return true;
        }

        uint length_before = whitelistArray.length;
        this.confirmWhitelistRemoval(calculateHash(pendingWhitelistRemoval()));
        uint length_after = whitelistArray.length;
        return (length_before >= length_after);
    }

    function echidna_Whitelist_removal_map() public returns (bool) {
        if (!submittedWhitelistRemoval) {
            return true;
        }

        for (uint i = 0; i < pendingWhitelistRemoval().length; i++) {
            removed_addresses.push(pendingWhitelistRemoval()[i]);
        }

        this.confirmWhitelistRemoval(calculateHash(pendingWhitelistRemoval()));

        bool checked_addresses = true;
        for (uint i = 0; i < removed_addresses.length; i++) {
            checked_addresses = checked_addresses && !whitelistMap[removed_addresses[i]];
        }

        return checked_addresses;
    }
}
