pragma solidity ^0.6.0;

import "../tmp_0_6/ensResolvable.sol";
import "../tmp_0_6/tokenWhitelistable.sol";


contract TokenWhitelistableExporter is ENSResolvable, TokenWhitelistable {
    constructor(address _ens_, bytes32 _tokenWhitelistName_) public ENSResolvable(_ens_) TokenWhitelistable(_tokenWhitelistName_) {}

    function getTokenInfo(address _a) external view returns (string memory, uint256, uint256, bool, bool, bool, uint256) {
        return _getTokenInfo(_a);
    }

    function getStablecoinInfo() external view returns (string memory, uint256, uint256, bool, bool, bool, uint256) {
        return _getStablecoinInfo();
    }

    function tokenAddressArray() external view returns (address[] memory) {
        return _tokenAddressArray();
    }

    function redeemableTokens() external view returns (address[] memory) {
        return _redeemableTokens();
    }

    function updateTokenRate(address _token, uint256 _rate, uint256 _updateDate) external {
        return _updateTokenRate(_token, _rate, _updateDate);
    }

    function getERC20RecipientAndAmount(address _destination, bytes calldata _data) external view returns (address, uint256) {
        return _getERC20RecipientAndAmount(_destination, _data);
    }

    function isTokenLoadable(address _a) external view returns (bool) {
        return _isTokenLoadable(_a);
    }

    function isTokenAvailable(address _a) external view returns (bool) {
        return _isTokenAvailable(_a);
    }

    function isTokenRedeemable(address _a) external view returns (bool) {
        return _isTokenRedeemable(_a);
    }
}
