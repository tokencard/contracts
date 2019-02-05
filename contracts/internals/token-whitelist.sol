/**
 *  The Consumer Contract Wallet
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
import "../externals/strings.sol";
import "../externals/ens/PublicResolver.sol";
import "../externals/SafeMath.sol";

/// @title The TokenWhitelist interface provides access to an external list of tokens.
interface ITokenWhitelist {
    function getTokenInfo(address _a) external view returns (string symbol, uint256 magnitude, uint256 rate, bool available, bool loadable, uint256 lastUpdate);
    function getTokenAddressArray() external view returns (address[]);
    function updateTokenRate(address _token, uint _rate, uint _updateDate) external;
}

/// @title TokenWhitelist stores a list of tokens used by the Consumer Contract Wallet, the Oracle, and the TKN Licence Contract
contract TokenWhitelist is Controllable, Ownable {
    using strings for *;
    using SafeMath for uint256;

    event UpdatedTokenRate(address _sender, address _token, uint _rate);

    event AddedToken(address _sender, address _token, string _symbol, uint _magnitude, bool _loadable);
    event RemovedToken(address _sender, address _token);

    struct Token {
        string symbol;    // Token symbol
        uint magnitude;   // 10^decimals
        uint rate;        // Token exchange rate in wei
        bool available;   // Flags if the token is available or not
        bool loadable;    // Flags if token is loadable to the TokenCard
        uint lastUpdate;  // Time of the last rate update
    }

    mapping(address => Token) private _tokenInfoMap;
    address[] private _tokenAddressArray;

    modifier onlyControllerOrOracle() {
        address oracleAddress = PublicResolver(_ENS.resolver(_oracleNode)).addr(_oracleNode);
        require (_isController(msg.sender) || msg.sender == oracleAddress);
        _;
    }

    /// @dev ENS points to the ENS registry smart contract.
    ENS internal _ENS;

    /// @dev Is the registered ENS name of the oracle contract.
    bytes32 private _oracleNode;

    /// @dev Constructor initializes ens and the oracle.
    /// @param _ens is the ENS public registry contract address.
    /// @param _oracleName is the ENS name of the Oracle.
    constructor(address _ens, bytes32 _oracleName) public {
        _ENS = ENS(_ens);
        _oracleNode = _oracleName;
    }

    /// @dev Add ERC20 tokens to the list of supported tokens.
    /// @param _tokens ERC20 token contract addresses.
    /// @param _symbols ERC20 token names.
    /// @param _magnitude 10 to the power of number of decimal places used by each ERC20 token.
    /// @param _loadable is a bool that states whether or not a token is loadable to the TokenCard.
    /// @param _updateDate date for the token updates. This will be compared to when oracle updates are received.
    function addTokens(address[] _tokens, bytes32[] _symbols, uint[] _magnitude, bool[] _loadable, uint _updateDate) external onlyController {
        // Require that all parameters have the same length.
        require(_tokens.length == _symbols.length && _tokens.length == _magnitude.length && _tokens.length == _loadable.length, "parameter lengths do not match");
        // Add each token to the list of supported tokens.
        for (uint i = 0; i < _tokens.length; i++) {
            // Require that the token isn't already available.
            address token = _tokens[i];
            require(!_tokenInfoMap[token].available, "token already available");
            // Store the intermediate values.
            string memory symbol = _symbols[i].toSliceB32().toString();
            uint magnitude = _magnitude[i];
            bool loadable = _loadable[i];
            // Add the token to the token list.
            _tokenInfoMap[token] = Token({
                symbol : symbol,
                magnitude : magnitude,
                rate : 0,
                available : true,
                loadable : loadable,
                lastUpdate : _updateDate
                });
            // Add the token address to the address list.
            _tokenAddressArray.push(token);
            // Emit token addition event.
            emit AddedToken(msg.sender, token, symbol, magnitude, loadable);
        }
    }

    /// @dev Remove ERC20 tokens from the list of supported tokens.
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

    /// @dev Update ERC20 token exchange rate manually.
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

    function getTokenInfo(address _a) external view returns (string, uint256, uint256, bool, bool, uint256) {
        return (_tokenInfoMap[_a].symbol, _tokenInfoMap[_a].magnitude, _tokenInfoMap[_a].rate, _tokenInfoMap[_a].available, _tokenInfoMap[_a].loadable, _tokenInfoMap[_a].lastUpdate);
    }

    function getTokenAddressArray() external view returns (address[]) {
        return _tokenAddressArray;
    }
}
