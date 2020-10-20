pragma solidity ^0.6.0;


contract ChainlinkMock {


  int256 public rate;

  function setLatestAnswer(int256 _rate) public returns (int256) {
    rate = _rate;
  }

  function getLatestAnswer() public returns (int256) {
    return rate;
  }

  function getLatestTimestamp() public returns (uint256) {
    return block.timestamp;
  }


}
