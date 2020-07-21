// SPDX-License-Identifier: MIT

pragma solidity ^0.6.0;

/// @title ERC20 interface is a subset of the ERC20 specification.
/// @notice see https://github.com/ethereum/EIPs/issues/20
interface IERC20 {
    function allowance(address _owner, address _spender) external view returns (uint256);
    function approve(address _spender, uint256 _value) external returns (bool);
    function balanceOf(address _who) external view returns (uint256);
    function totalSupply() external view returns (uint256);
    function transfer(address _to, uint256 _value) external returns (bool);
    function transferFrom(address _from, address _to, uint256 _value) external returns (bool);
}
