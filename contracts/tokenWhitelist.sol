/**
 *  TokenWhitelist - The Consumer Contract Wallet
 *  Copyright (C) 2019 The Contract Wallet Company Limited
 *
 *  This program is free software: you can redistribute it and/or modify
 *  it under the terms of the GNU General Public License as published by
 *  the Free Software Foundation, either version 3 of the License, or
 *  (at your option) any later version.

 *  This program is distributed in the hope that it will be useful,
 *  but WITHOUT ANY WARRANTY; without even the implied warranty of
 *  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *  GNU General Public License for more details.

 *  You should have received a copy of the GNU General Public License
 *  along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

// SPDX-License-Identifier: GPLv3

pragma solidity ^0.6.11;

import "./internals/bytesUtils.sol";
import "./internals/controllable.sol";
import "./internals/transferrable.sol";
import "./externals/SafeMath.sol";
import "./externals/strings.sol";

/// @title TokenWhitelist stores a list of tokens used by the Consumer Contract Wallet, the Oracle, the TKN Holder and the TKN Licence Contract
contract TokenWhitelist is ENSResolvable, Controllable, Transferrable {
    using strings for *;
    using SafeMath for uint256;
    using BytesUtils for bytes;

    event UpdatedTokenRate(address _sender, address _token, uint256 _rate);

    event UpdatedTokenLoadable(address _sender, address _token, bool _loadable);
    event UpdatedTokenRedeemable(address _sender, address _token, bool _redeemable);

    event AddedToken(address _sender, address _token, string _symbol, uint256 _magnitude, bool _loadable, bool _redeemable);
    event RemovedToken(address _sender, address _token);

    event AddedMethodId(bytes4 _methodId);
    event RemovedMethodId(bytes4 _methodId);
    event AddedExclusiveMethod(address _token, bytes4 _methodId);
    event RemovedExclusiveMethod(address _token, bytes4 _methodId);

    event Claimed(address _to, address _asset, uint256 _amount);

    /// @dev these are the methods whitelisted by default in executeTransaction() for protected tokens
    bytes4 private constant _APPROVE = 0x095ea7b3; // keccak256(approve(address,uint256)) => 0x095ea7b3
    bytes4 private constant _BURN = 0x42966c68; // keccak256(burn(uint256)) => 0x42966c68
    bytes4 private constant _TRANSFER = 0xa9059cbb; // keccak256(transfer(address,uint256)) => 0xa9059cbb
    bytes4 private constant _TRANSFER_FROM = 0x23b872dd; // keccak256(transferFrom(address,address,uint256)) => 0x23b872dd

    struct Token {
        string symbol; // Token symbol
        uint256 magnitude; // 10^decimals
        uint256 rate; // Token exchange rate in wei
        bool available; // Flags if the token is available or not
        bool loadable; // Flags if token is loadable to the TokenCard
        bool redeemable; // Flags if token is redeemable in the TKN Holder contract
        uint256 lastUpdate; // Time of the last rate update
    }

    mapping(address => Token) private _tokenInfoMap;

    // @notice specifies whitelisted methodIds for protected tokens in wallet's excuteTranaction() e.g. keccak256(transfer(address,uint256)) => 0xa9059cbb
    mapping(bytes4 => bool) private _methodIdWhitelist;

    address[] private _tokenAddressArray;

    /// @dev keeping track of how many redeemable tokens are in the tokenWhitelist
    uint256 private _redeemableCounter;

    /// @dev Address of the stablecoin.
    address private _stablecoin;

    /// @dev is registered ENS node identifying the oracle contract.
    bytes32 private _oracleNode;

    /// @notice Constructor initializes ENSResolvable, and Controllable.
    /// @param _ens_ is the ENS registry address.
    /// @param _oracleNode_ is the ENS node of the Oracle.
    /// @param _controllerNode_ is our Controllers node.
    /// @param _stablecoinAddress_ is the address of the stablecoint used by the wallet for the card load limit.
    constructor(
        address _ens_,
        bytes32 _oracleNode_,
        bytes32 _controllerNode_,
        address _stablecoinAddress_
    ) public {
        _initializeENSResolvable(_ens_);
        _initializeControllable(_controllerNode_);
        _oracleNode = _oracleNode_;
        _stablecoin = _stablecoinAddress_;
        //a priori ERC20 whitelisted methods
        _methodIdWhitelist[_APPROVE] = true;
        _methodIdWhitelist[_BURN] = true;
        _methodIdWhitelist[_TRANSFER] = true;
        _methodIdWhitelist[_TRANSFER_FROM] = true;
    }

    modifier onlyAdminOrOracle() {
        address oracleAddress = _ensResolve(_oracleNode);
        require(_isAdmin(msg.sender) || msg.sender == oracleAddress, "either oracle or admin");
        _;
    }

    /// @notice Add ERC20 tokens to the list of whitelisted tokens.
    /// @param _tokens ERC20 token contract addresses.
    /// @param _symbols ERC20 token names.
    /// @param _magnitude 10 to the power of number of decimal places used by each ERC20 token.
    /// @param _loadable is a bool that states whether or not a token is loadable to the TokenCard.
    /// @param _redeemable is a bool that states whether or not a token is redeemable in the TKN Holder Contract.
    /// @param _lastUpdate is a unit representing an ISO datetime e.g. 20180913153211.
    function addTokens(
        address[] calldata _tokens,
        bytes32[] calldata _symbols,
        uint256[] calldata _magnitude,
        bool[] calldata _loadable,
        bool[] calldata _redeemable,
        uint256 _lastUpdate
    ) external onlyAdmin {
        // Require that all parameters have the same length.
        require(
            _tokens.length == _symbols.length &&
                _tokens.length == _magnitude.length &&
                _tokens.length == _loadable.length &&
                _tokens.length == _loadable.length,
            "parameter lengths do not match"
        );
        // Add each token to the list of supported tokens.
        for (uint256 i = 0; i < _tokens.length; i++) {
            // Require that the token isn't already available.
            require(!_tokenInfoMap[_tokens[i]].available, "token already available");
            // Store the intermediate values.
            string memory symbol = _symbols[i].toSliceB32().toString();
            // Add the token to the token list.
            _tokenInfoMap[_tokens[i]] = Token({
                symbol: symbol,
                magnitude: _magnitude[i],
                rate: 0,
                available: true,
                loadable: _loadable[i],
                redeemable: _redeemable[i],
                lastUpdate: _lastUpdate
            });
            // Add the token address to the address list.
            _tokenAddressArray.push(_tokens[i]);
            //if the token is redeemable, increase the redeemableCounter
            if (_redeemable[i]) {
                _redeemableCounter = _redeemableCounter.add(1);
            }
            // Emit token addition event.
            emit AddedToken(msg.sender, _tokens[i], symbol, _magnitude[i], _loadable[i], _redeemable[i]);
        }
    }

    /// @notice Remove ERC20 tokens from the whitelist of tokens.
    /// @param _tokens ERC20 token contract addresses.
    function removeTokens(address[] calldata _tokens) external onlyAdmin {
        // Delete each token object from the list of supported tokens based on the addresses provided.
        for (uint256 i = 0; i < _tokens.length; i++) {
            // Store the token address.
            address token = _tokens[i];
            //token must be available, reverts on duplicates as well
            require(_tokenInfoMap[token].available, "token is not available");
            //if the token is redeemable decrease the redeemableCounter
            if (_tokenInfoMap[token].redeemable) {
                _redeemableCounter = _redeemableCounter.sub(1);
            }
            // Delete the token object.
            delete _tokenInfoMap[token];
            // Remove the token address from the address list.
            for (uint256 j = 0; j < _tokenAddressArray.length.sub(1); j++) {
                if (_tokenAddressArray[j] == token) {
                    _tokenAddressArray[j] = _tokenAddressArray[_tokenAddressArray.length.sub(1)];
                    break;
                }
            }
            _tokenAddressArray.pop();
            // Emit token removal event.
            emit RemovedToken(msg.sender, token);
        }
    }

    /// @notice based on the method it returns the recipient address and amount/value, ERC20 specific.
    /// @param _token ERC20 token contract address.
    /// @param _data is the transaction payload.
    function getERC20RecipientAndAmount(address _token, bytes calldata _data) external view returns (address, uint256) {
        // Require that there exist enough bytes for encoding at least a method signature + data in the transaction payload:
        // 4 (signature)  + 32(address or uint256)
        require(_data.length >= 4 + 32, "not enough method-encoding bytes");
        // Get the method signature
        bytes4 signature = _data._bytesToBytes4(0);
        // Check if method Id is supported
        require(isERC20MethodSupported(_token, signature), "unsupported method");
        // returns the recipient's address and amount is the value to be transferred
        if (signature == _BURN) {
            // 4 (signature) + 32(uint256)
            return (_token, _data._bytesToUint256(4));
        } else if (signature == _TRANSFER_FROM) {
            // 4 (signature) + 32(address) + 32(address) + 32(uint256)
            require(_data.length >= 4 + 32 + 32 + 32, "not enough data for transferFrom");
            return (_data._bytesToAddress(4 + 32 + 12), _data._bytesToUint256(4 + 32 + 32));
        } else {
            //transfer or approve
            // 4 (signature) + 32(address) + 32(uint)
            require(_data.length >= 4 + 32 + 32, "not enough data for transfer/appprove");
            return (_data._bytesToAddress(4 + 12), _data._bytesToUint256(4 + 32));
        }
    }

    /// @notice Toggles whether or not a token is loadable.
    /// @param _token ERC20 token contract address.
    /// @param _loadable bool that toggles the corresponding parameter.
    function setTokenLoadable(address _token, bool _loadable) external onlyAdmin {
        // Require that the token exists.
        require(_tokenInfoMap[_token].available, "loadable: token is not available");

        // this sets the loadable flag to the value passed in
        _tokenInfoMap[_token].loadable = _loadable;

        emit UpdatedTokenLoadable(msg.sender, _token, _loadable);
    }

    /// @notice Toggles whether or not a token is redeemable.
    /// @param _token ERC20 token contract address.
    /// @param _redeemable bool that toggles the corresponding parameter
    function setTokenRedeemable(address _token, bool _redeemable) external onlyAdmin {
        // Require that the token exists.
        require(_tokenInfoMap[_token].available, "redeemable: token not available");
        // If it does not change the current state i.e. _tokenInfoMap[_token].redeemable == _redeemable, revert!
        require(_tokenInfoMap[_token].redeemable != _redeemable, "redeemable: no state change");
        if (_redeemable) {
            _redeemableCounter = _redeemableCounter.add(1);
        } else {
            _redeemableCounter = _redeemableCounter.sub(1);
        }
        // This sets the redeemable flag to the value passed in
        _tokenInfoMap[_token].redeemable = _redeemable;

        emit UpdatedTokenRedeemable(msg.sender, _token, _redeemable);
    }

    /// @notice Update ERC20 token exchange rate.
    /// @param _token ERC20 token contract address.
    /// @param _rate ERC20 token exchange rate in wei.
    /// @param _updateDate date for the token updates. This will be compared to when oracle updates are received.
    function updateTokenRate(
        address _token,
        uint256 _rate,
        uint256 _updateDate
    ) external onlyAdminOrOracle {
        // Require that the token exists.
        require(_tokenInfoMap[_token].available, "token is not available");
        // Update the token's rate.
        _tokenInfoMap[_token].rate = _rate;
        // Update the token's last update timestamp.
        _tokenInfoMap[_token].lastUpdate = _updateDate;
        // Emit the rate update event.
        emit UpdatedTokenRate(msg.sender, _token, _rate);
    }

    //// @notice Withdraw tokens from the smart contract to the specified account.
    function claim(
        address payable _to,
        address _asset,
        uint256 _amount
    ) external onlyAdmin {
        _safeTransfer(_to, _asset, _amount);
        emit Claimed(_to, _asset, _amount);
    }

    /// @notice This returns all of the fields for a given token.
    /// @param _a is the address of a given token.
    /// @return string of the token's symbol.
    /// @return uint of the token's magnitude.
    /// @return uint of the token's exchange rate to ETH.
    /// @return bool whether the token is available.
    /// @return bool whether the token is loadable to the TokenCard.
    /// @return bool whether the token is redeemable to the TKN Holder Contract.
    /// @return uint of the lastUpdated time of the token's exchange rate.
    function getTokenInfo(address _a)
        external
        view
        returns (
            string memory,
            uint256,
            uint256,
            bool,
            bool,
            bool,
            uint256
        )
    {
        Token storage tokenInfo = _tokenInfoMap[_a];
        return (tokenInfo.symbol, tokenInfo.magnitude, tokenInfo.rate, tokenInfo.available, tokenInfo.loadable, tokenInfo.redeemable, tokenInfo.lastUpdate);
    }

    /// @notice This returns all of the fields for our StableCoin.
    /// @return string of the token's symbol.
    /// @return uint of the token's magnitude.
    /// @return uint of the token's exchange rate to ETH.
    /// @return bool whether the token is available.
    /// @return bool whether the token is loadable to the TokenCard.
    /// @return bool whether the token is redeemable to the TKN Holder Contract.
    /// @return uint of the lastUpdated time of the token's exchange rate.
    function getStablecoinInfo()
        external
        view
        returns (
            string memory,
            uint256,
            uint256,
            bool,
            bool,
            bool,
            uint256
        )
    {
        Token storage stablecoinInfo = _tokenInfoMap[_stablecoin];
        return (
            stablecoinInfo.symbol,
            stablecoinInfo.magnitude,
            stablecoinInfo.rate,
            stablecoinInfo.available,
            stablecoinInfo.loadable,
            stablecoinInfo.redeemable,
            stablecoinInfo.lastUpdate
        );
    }

    /// @notice This returns an array of all whitelisted token addresses.
    /// @return address[] of whitelisted tokens.
    function tokenAddressArray() external view returns (address[] memory) {
        return _tokenAddressArray;
    }

    /// @notice This returns an array of all redeemable token addresses.
    /// @return address[] of redeemable tokens.
    function redeemableTokens() external view returns (address[] memory) {
        address[] memory redeemableAddresses = new address[](_redeemableCounter);
        uint256 redeemableIndex = 0;
        for (uint256 i = 0; i < _tokenAddressArray.length; i++) {
            address token = _tokenAddressArray[i];
            if (_tokenInfoMap[token].redeemable) {
                redeemableAddresses[redeemableIndex] = token;
                redeemableIndex += 1;
            }
        }
        return redeemableAddresses;
    }

    /// @notice This returns true if a method Id is supported for the specific token.
    /// @param _token ERC20 token contract address.
    /// @param _methodId The first four bytes of the Keccak256 hash of the function signature...
    /// ...i.e the canonical expression of the basic prototype without data location specifier.
    /// @return True if _methodId is supported in general or just for the specific token.
    function isERC20MethodSupported(address _token, bytes4 _methodId) public view returns (bool) {
        require(_tokenInfoMap[_token].available, "non-existing token");
        return (_methodIdWhitelist[_methodId]);
    }

    /// @notice This returns true if the method is supported for all protected tokens.
    /// @return true if _methodId is in the method whitelist.
    function isERC20MethodWhitelisted(bytes4 _methodId) external view returns (bool) {
        return (_methodIdWhitelist[_methodId]);
    }

    /// @notice This returns the number of redeemable tokens.
    /// @return current # of redeemables.
    function redeemableCounter() external view returns (uint256) {
        return _redeemableCounter;
    }

    /// @notice This returns the address of our stablecoin of choice.
    /// @return the address of the stablecoin contract.
    function stablecoin() external view returns (address) {
        return _stablecoin;
    }

    /// @notice this returns the node hash of our Oracle.
    /// @return the oracle node registered in ENS.
    function oracleNode() external view returns (bytes32) {
        return _oracleNode;
    }
}
