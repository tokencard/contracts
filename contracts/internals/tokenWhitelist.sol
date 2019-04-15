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

pragma solidity ^0.4.25;

import "./ownable.sol";
import "./controllable.sol";
import "./claimable.sol";
import "../externals/strings.sol";
import "../externals/SafeMath.sol";

/// @title The ITokenWhitelist interface provides access to a whitelist of tokens.
interface ITokenWhitelist {
    function getTokenInfo(address) external view returns (string, uint256, uint256, bool, bool, bool, uint256);
    function getStablecoinInfo() external view returns (string, uint256, uint256, bool, bool, bool, uint256);
    function tokenAddressArray() external view returns (address[]);
    function updateTokenRate(address, uint, uint) external;
    function stablecoin() external view returns (address);
}


/// @title TokenWhitelist stores a list of tokens used by the Consumer Contract Wallet, the Oracle, the TKN Holder and the TKN Licence Contract
contract TokenWhitelist is ENSResolvable, Controllable, Ownable, Claimable {
    using strings for *;
    using SafeMath for uint256;

    event UpdatedTokenRate(address _sender, address _token, uint _rate);

    event AddedToken(address _sender, address _token, string _symbol, uint _magnitude, bool _loadable, bool _burnable);
    event RemovedToken(address _sender, address _token);

    struct Token {
        string symbol;    // Token symbol
        uint magnitude;   // 10^decimals
        uint rate;        // Token exchange rate in wei
        bool available;   // Flags if the token is available or not
        bool loadable;    // Flags if token is loadable to the TokenCard
        bool burnable;    // Flags if token is burnable in the TKN Holder contract
        uint lastUpdate;  // Time of the last rate update
    }

    mapping(address => Token) private _tokenInfoMap;
    address[] private _tokenAddressArray;

    modifier onlyControllerOrOracle() {
        address oracleAddress = _ensResolve(_oracleNode);
        require (_isController(msg.sender) || msg.sender == oracleAddress, "either oracle or controller");
        _;
    }

    /// @notice Address of the stablecoin.
    address private _stablecoin;

    /// @notice is registered ENS node hash identifying the oracle contract.
    bytes32 private _oracleNode;

    /// @notice Constructor initializes ENSResolvable, Controllable, and Ownable
    /// @param _oracleNameHash_ is the ENS name hash of the Oracle.
    /// @param _stabelcoinAddress_ is the address of the stablecoint used by the wallet
    constructor(address _ens_, bytes32 _oracleNameHash_, bytes32 _controllerNameHash_, address _owner_, bool _transferable_, address _stabelcoinAddress_) ENSResolvable(_ens_) Controllable(_controllerNameHash_) Ownable(_owner_, _transferable_) public {
        _oracleNode = _oracleNameHash_;
        _stablecoin = _stabelcoinAddress_;
    }

    /// @notice Add ERC20 tokens to the list of whitelisted tokens.
    /// @param _tokens ERC20 token contract addresses.
    /// @param _symbols ERC20 token names.
    /// @param _magnitude 10 to the power of number of decimal places used by each ERC20 token.
    /// @param _loadable is a bool that states whether or not a token is loadable to the TokenCard.
    /// @param _burnable is a bool that states whether or not a token is burnable in the TKN Holder Contract.
    /// @param _lastUpdate is a unit representing an ISO datetime e.g. 20180913153211
    function addTokens(address[] _tokens, bytes32[] _symbols, uint[] _magnitude, bool[] _loadable, bool[] _burnable, uint _lastUpdate) external onlyController {
        // Require that all parameters have the same length.
        require(_tokens.length == _symbols.length && _tokens.length == _magnitude.length && _tokens.length == _loadable.length && _tokens.length == _loadable.length, "parameter lengths do not match");
        // Add each token to the list of supported tokens.
        for (uint i = 0; i < _tokens.length; i++) {
            // Require that the token isn't already available.
            require(!_tokenInfoMap[_tokens[i]].available, "token already available");
            // Store the intermediate values.
            string memory symbol = _symbols[i].toSliceB32().toString();
            // Add the token to the token list.
            _tokenInfoMap[_tokens[i]] = Token({
                symbol : symbol,
                magnitude : _magnitude[i],
                rate : 0,
                available : true,
                loadable : _loadable[i],
                burnable: _burnable[i],
                lastUpdate : _lastUpdate
                });
            // Add the token address to the address list.
            _tokenAddressArray.push(_tokens[i]);
            // Emit token addition event.
            emit AddedToken(msg.sender, _tokens[i], symbol, _magnitude[i], _loadable[i], _burnable[i]);
        }
    }

    /// @notice Remove ERC20 tokens from the whitelist of tokens.
    /// @param _tokens ERC20 token contract addresses.
    function removeTokens(address[] _tokens) external onlyController {
        // Delete each token object from the list of supported tokens based on the addresses provided.
        for (uint i = 0; i < _tokens.length; i++) {
            //token must be available, reverts on duplicates as well
            require(_tokenInfoMap[_tokens[i]].available, "token is not available");
            // Store the token address.
            address token = _tokens[i];
            // Delete the token object.
            delete _tokenInfoMap[token];
            // Remove the token address from the address list.
            for (uint j = 0; j < _tokenAddressArray.length.sub(1); j++) {
                if (_tokenAddressArray[j] == token) {
                    _tokenAddressArray[j] = _tokenAddressArray[_tokenAddressArray.length.sub(1)];
                    break;
                }
            }
            _tokenAddressArray.length--;
            // Emit token removal event.
            emit RemovedToken(msg.sender, token);
        }
    }

    /// @notice Update ERC20 token exchange rate.
    /// @param _token ERC20 token contract address.
    /// @param _rate ERC20 token exchange rate in wei.
    /// @param _updateDate date for the token updates. This will be compared to when oracle updates are received.
    function updateTokenRate(address _token, uint _rate, uint _updateDate) external onlyControllerOrOracle {
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
    function claim(address _to, address _asset, uint _amount) external onlyOwner {
        _claim(_to, _asset, _amount);
    }

    /// @notice This returns all of the fields for a given token
    /// @param _a is the address of a given token
    /// @return string of the token's symbol
    /// @return uint of the token's magnitude
    /// @return uint of the token's exchange rate to ETH
    /// @return bool whether the token is available
    /// @return bool whether the token is loadable to the TokenCard
    /// @return bool whether the token is burnable to the TKN Holder Contract
    /// @return uint of the lastUpdated time of the token's exchange rate
    function getTokenInfo(address _a) external view returns (string, uint256, uint256, bool, bool, bool, uint256) {
        Token storage tokenInfo = _tokenInfoMap[_a];
        return (tokenInfo.symbol, tokenInfo.magnitude, tokenInfo.rate, tokenInfo.available, tokenInfo.loadable, tokenInfo.burnable, tokenInfo.lastUpdate);
    }

    /// @notice This returns all of the fields for our StableCoin
    /// @return string of the token's symbol
    /// @return uint of the token's magnitude
    /// @return uint of the token's exchange rate to ETH
    /// @return bool whether the token is available
    /// @return bool whether the token is loadable to the TokenCard
    /// @return bool whether the token is burnable to the TKN Holder Contract
    /// @return uint of the lastUpdated time of the token's exchange rate
    function getStablecoinInfo() external view returns (string, uint256, uint256, bool, bool, bool, uint256) {
        Token storage stablecoinInfo = _tokenInfoMap[_stablecoin];
        return (stablecoinInfo.symbol, stablecoinInfo.magnitude, stablecoinInfo.rate, stablecoinInfo.available, stablecoinInfo.loadable, stablecoinInfo.burnable, stablecoinInfo.lastUpdate);
    }

    /// @notice This returns an array of our whitelisted address
    /// @return address[] of our whitelisted tokens
    function tokenAddressArray() external view returns (address[]) {
        return _tokenAddressArray;
    }

    /// @notice This returns the address of our stablecoin of choice
    /// @return the address of the stablecoin contract.
    function stablecoin() external view returns (address) {
        return _stablecoin;
    }

    /// @notice this returns the node hash of our Oracle
    /// @return the oracle node registered in ENS.
    function oracleNode() external view returns (bytes32) {
        return _oracleNode;
    }
}
