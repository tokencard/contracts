// SPDX-License-Identifier: GPLv3

pragma solidity ^0.6.11;

interface IPublicResolver {

    function addr(bytes32) external view returns (address);

}
