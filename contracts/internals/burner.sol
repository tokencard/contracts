pragma solidity ^0.5.10;

// The BurnerToken interface is the interface to a token contract which
// provides the total burnable supply for the TokenHolder contract.
interface IBurner {
    function currentSupply() external view returns (uint);
}
