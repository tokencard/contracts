pragma solidity ^0.5.15;

import "crytic-export/flattening/Wallet.sol";


contract Echidna {
    address internal echidna_deployer = address(0x1);
    address payable internal echidna_owner = address(0x2);
    address payable internal echidna_attacker = address(0x3);
}


contract TEST is Echidna, AddressWhitelist {
    address[] added_addresses;
    address[] removed_addresses;

    constructor() public Ownable(echidna_owner, false) Controllable(bytes32(0x0)) ENSResolvable(address(0x0)) {}

    modifier onlyController() {
        require(msg.sender == echidna_owner || msg.sender == address(this), "");
        _;
    }

    function echidna_hasNoOwnerOrZeroAddress() public view returns (bool) {
        for (uint256 i = 0; i < whitelistArray.length; i++) {
            if (_isOwner(whitelistArray[i])) {
                return false;
            }
            if (whitelistArray[i] == address(0)) {
                return false;
            }
        }
        return true;
    }

    function echidna_whitelistAdditionLength() public returns (bool) {
        if (!submittedWhitelistAddition) {
            return true;
        }

        uint256 length_before = whitelistArray.length;
        this.confirmWhitelistAddition(
            calculateHash(pendingWhitelistAddition())
        );
        uint256 length_after = whitelistArray.length;
        return (length_before <= length_after);
    }

    function echidna_whitelistAdditionMap() public returns (bool) {
        if (!submittedWhitelistAddition) {
            return true;
        }

        for (uint256 i = 0; i < pendingWhitelistAddition().length; i++) {
            added_addresses.push(pendingWhitelistAddition()[i]);
        }

        this.confirmWhitelistAddition(
            calculateHash(pendingWhitelistAddition())
        );

        bool checked_addresses = true;
        for (uint256 i = 0; i < added_addresses.length; i++) {
            checked_addresses = checked_addresses && whitelistMap[added_addresses[i]];
        }

        return checked_addresses;
    }

    function echidna_whitelistRemovalLength() public returns (bool) {
        if (!submittedWhitelistRemoval) {
            return true;
        }

        uint256 length_before = whitelistArray.length;
        this.confirmWhitelistRemoval(calculateHash(pendingWhitelistRemoval()));
        uint256 length_after = whitelistArray.length;
        return (length_before >= length_after);
    }

    function echidna_whitelistRemovalMap() public returns (bool) {
        if (!submittedWhitelistRemoval) {
            return true;
        }

        for (uint256 i = 0; i < pendingWhitelistRemoval().length; i++) {
            removed_addresses.push(pendingWhitelistRemoval()[i]);
        }

        this.confirmWhitelistRemoval(calculateHash(pendingWhitelistRemoval()));

        bool checked_addresses = true;
        for (uint256 i = 0; i < removed_addresses.length; i++) {
            checked_addresses = checked_addresses && !whitelistMap[removed_addresses[i]];
        }

        return checked_addresses;
    }
}
