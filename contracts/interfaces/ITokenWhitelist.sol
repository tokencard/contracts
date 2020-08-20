// SPDX-License-Identifier: GPLv3

pragma solidity ^0.6.0;

/// @title The ITokenWhitelist interface provides access to a whitelist of tokens.
interface ITokenWhitelist {
    function getTokenInfo(address) external view returns (string memory, uint256, uint256, bool, bool, bool, uint256);

    function getStablecoinInfo() external view returns (string memory, uint256, uint256, bool, bool, bool, uint256);

    function tokenAddressArray() external view returns (address[] memory);

    function redeemableTokens() external view returns (address[] memory);

    function methodIdWhitelist(bytes4) external view returns (bool);

    function getERC20RecipientAndAmount(address, bytes calldata) external view returns (address, uint256);

    function stablecoin() external view returns (address);

    function updateTokenRate(address, uint256, uint256) external;
}