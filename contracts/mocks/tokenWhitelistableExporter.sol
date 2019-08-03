pragma solidity ^0.5.10;

import "../internals/tokenWhitelistable.sol";
import "../internals/ensResolvable.sol";


contract TokenWhitelistableExporter is ENSResolvable, TokenWhitelistable {

    constructor(address _ens_, bytes32 _tokenWhitelistName_) ENSResolvable(_ens_) TokenWhitelistable(_tokenWhitelistName_) public {}

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

    function updateTokenRate(address _token, uint _rate, uint _updateDate) external {
        return _updateTokenRate(_token, _rate, _updateDate);
    }

    function isMethodIdWhitelisted(uint32 _methodId) external view returns (bool) {
        return _isMethodIdWhitelisted(_methodId);
    }

    function getERC20RecipientAndAmount(bytes calldata _data) external view returns (address, uint) {
        return _getERC20RecipientAndAmount(_data);
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
