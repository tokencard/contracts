pragma solidity ^0.4.25;

import "../internals/tokenWhitelistable.sol";

contract TokenWhitelistableExporter is TokenWhitelistable {

  constructor(address _ens, bytes32 _tokenWhitelistName) TokenWhitelistable(_ens, _tokenWhitelistName) public {}

  function getTokenInfo(address _a) external view returns (string, uint256, uint256, bool, bool, uint256) {
      return _getTokenInfo(_a);
  }

  function getStablecoinInfo() external view returns (string, uint256, uint256, bool, bool, uint256) {
      return _getStablecoinInfo();
  }

  function getTokenAddressArray() external view returns (address[]) {
      return _getTokenAddressArray();
  }

  function updateTokenRate(address _token, uint _rate, uint _updateDate) external {
      return _updateTokenRate(_token, _rate, _updateDate);
  }

  function isTokenLoadable(address _a) external view returns (bool) {
      return _isTokenLoadable(_a);
  }

  function isTokenAvailable(address _a) external view returns (bool) {
      return _isTokenAvailable(_a);
  }

}
