pragma solidity ^0.4.25;

/// @title The DAO interface (WIP stub).
interface IDAO {
    function lock() external;
    function isLocked() external returns (bool);
}
