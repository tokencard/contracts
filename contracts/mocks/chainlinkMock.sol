pragma solidity ^0.6.0;

import "../interfaces/IChainlink.sol";


contract ChainlinkMock is IChainlink {
    int256 public rate;

    function setLatestAnswer(int256 _rate) public returns (int256) {
        rate = _rate;
    }

    function latestAnswer() public override view returns (int256) {
        return rate;
    }

    function latestTimestamp() public override view returns (uint256) {
        return block.timestamp;
    }
}
