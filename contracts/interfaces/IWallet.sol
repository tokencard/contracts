// SPDX-License-Identifier: GPLv3

pragma solidity ^0.6.11;

interface IWallet {

function initializeWallet(
        address payable _owner_,
        bool _transferable_,
        address _ens_,
        bytes32 _tokenWhitelistNode_,
        bytes32 _controllerNode_,
        bytes32 _licenceNode_,
        uint256 _spendLimit_
    ) external;

}