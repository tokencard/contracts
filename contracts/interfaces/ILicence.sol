// SPDX-License-Identifier: GPLv3

pragma solidity ^0.6.0;

/// @title ILicence interface describes methods for loading a TokenCard and updating licence amount.
interface ILicence {
    function load(address, uint256) external payable;

    function updateLicenceAmount(uint256) external;
}