// SPDX-License-Identifier: GPLv3

pragma solidity ^0.6.11;

import "crytic-export/flattening/Wallet.sol";


contract Echidna {
    address internal echidnaDeployer = address(0x1);
    address payable internal echidnaOwner = address(0x2);
    address payable internal echidnaAttacker = address(0x3);
    address internal echidnaRecipient = address(0x4);
}


contract TEST is Echidna, AddressWhitelist {
    constructor() public {
        _initializeOwnable(echidnaOwner, false);
        _initializeControllable(bytes32(0x0));
        _initializeENSResolvable(address(0x0));
        isSetWhitelist = true;
    }

    modifier onlyController() {
        require(msg.sender == msg.sender, "");
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
        uint256 lengthBefore = whitelistArray.length;
        address[] memory pendingWhitelist = pendingWhitelistAddition();
        confirmWhitelistAddition(calculateHash(pendingWhitelist));
        uint256 lengthAfter = whitelistArray.length;
        return (lengthBefore <= lengthAfter);
    }

    function echidna_whitelistAdditionMap() public returns (bool) {
        if (!submittedWhitelistAddition) {
            return true;
        }
        address[] memory pendingWhitelist = pendingWhitelistAddition();
        address[] memory addedAddresses = new address[](pendingWhitelist.length);
        for (uint256 i = 0; i < pendingWhitelist.length; i++) {
            addedAddresses[i] = pendingWhitelist[i];
        }
        confirmWhitelistAddition(calculateHash(pendingWhitelist));
        bool missingAddress = false;
        for (uint256 i = 0; i < addedAddresses.length; i++) {
            if (!whitelistMap[addedAddresses[i]]) {
                missingAddress = true;
            }
        }
        return !missingAddress;
    }

    function echidna_whitelistRemovalLength() public returns (bool) {
        if (!submittedWhitelistRemoval) {
            return true;
        }
        uint256 lengthBefore = whitelistArray.length;
        address[] memory pendingRemoval = pendingWhitelistRemoval();
        confirmWhitelistRemoval(calculateHash(pendingRemoval));
        uint256 lengthAfter = whitelistArray.length;
        return (lengthBefore >= lengthAfter);
    }

    function echidna_whitelistRemovalMap() public returns (bool) {
        if (!submittedWhitelistRemoval) {
            return true;
        }
        address[] memory pendingRemoval = pendingWhitelistRemoval();
        address[] memory removedAddresses = new address[](pendingRemoval.length);
        for (uint256 i = 0; i < pendingRemoval.length; i++) {
            removedAddresses[i] = pendingRemoval[i];
        }
        confirmWhitelistRemoval(calculateHash(pendingRemoval));
        bool addressNotRemoved = false;
        for (uint256 i = 0; i < removedAddresses.length; i++) {
            if (whitelistMap[removedAddresses[i]]) {
                addressNotRemoved = true;
            }
        }
        return !addressNotRemoved;
    }
}
