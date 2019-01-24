pragma solidity ^0.4.25;

/// @title The DAO interface (WIP stub).
interface IDAO {
    function lock() external;
    function isLocked() external returns (bool);
}


contract Dao {

  bool public locked;

  function lock() external{

    locked = true;

  }

  function isLocked() external view returns (bool){

    return locked;

  }

}
