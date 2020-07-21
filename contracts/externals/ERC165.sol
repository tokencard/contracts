// SPDX-License-Identifier: MIT

pragma solidity ^0.6.0;

/// @title ERC165 interface specifies a standard way of querying if a contract implements an interface.
interface ERC165 {
    function supportsInterface(bytes4) external view returns (bool);
}
