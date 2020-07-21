// SPDX-License-Identifier: GPLv3

pragma solidity ^0.6.11;

import "../internals/tokenWhitelistable.sol";
import "../internals/ensResolvable.sol";


contract TokenWhitelistableExporter is ENSResolvable, TokenWhitelistable {
    constructor(address _ens_, bytes32 _tokenWhitelistNode_) public {
        _initializeENSResolvable(_ens_);
        _initializeTokenWhitelistable(_tokenWhitelistNode_);
    }

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
