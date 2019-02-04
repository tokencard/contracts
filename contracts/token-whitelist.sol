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

import "./internals/ownable.sol";
import "./internals/controller.sol";
import "./externals/strings.sol";

/// @title The Token Whitelist interface provides access to an external list of tokens.
interface ITokenWhitelist {
    function isController(address) external view returns (bool);
}

/// @title TokenWhiteli stores a list of tokens used by the Consumer Contract Wallet, the Oracle, and the TKN Licence Contract
contract TokenWhitelist is ITokenWhitelist, Controller , Ownable {

    event AddedToken(address _sender, address _token, string _symbol, uint _magnitude, bool _loadable);
    event RemovedToken(address _sender, address _token);

    struct Token {
        string symbol;    // Token symbol
        uint magnitude;   // 10^decimals
        uint rate;        // Token exchange rate in wei
        bool exists;      // Flags if the struct is empty or not
        bool loadable;    // Flags if token is loadable to the TokenCard
        uint lastUpdate;  // Time of the last rate update
    }

    mapping(address => Token) public tokens;
    address[] private _tokenAddresses;

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
            // Require that the token doesn't already exist.
            address token = _tokens[i];
            require(!tokens[token].exists, "token already exists");
            // Store the intermediate values.
            string memory symbol = _symbols[i].toSliceB32().toString();
            uint magnitude = _magnitude[i];
            bool loadable = _loadable[i];
            // Add the token to the token list.
            tokens[token] = Token({
                symbol : symbol,
                magnitude : magnitude,
                rate : 0,
                exists : true,
                loadable : loadable,
                lastUpdate : _updateDate
                });
            // Add the token address to the address list.
            _tokenAddresses.push(token);
            // Emit token addition event.
            emit AddedToken(msg.sender, token, symbol, magnitude, loadable);
        }
    }

    /// @dev Remove ERC20 tokens from the list of supported tokens.
    /// @param _tokens ERC20 token contract addresses.
    function removeTokens(address[] _tokens) external onlyController {
        // Delete each token object from the list of supported tokens based on the addresses provided.
        for (uint i = 0; i < _tokens.length; i++) {
            //token must exist, reverts on duplicates as well
            require(tokens[_tokens[i]].exists, "token does not exist");
            // Store the token address.
            address token = _tokens[i];
            // Delete the token object.
            delete tokens[token];
            // Remove the token address from the address list.
            for (uint j = 0; j < _tokenAddresses.length.sub(1); j++) {
                if (_tokenAddresses[j] == token) {
                    _tokenAddresses[j] = _tokenAddresses[_tokenAddresses.length.sub(1)];
                    break;
                }
            }
            _tokenAddresses.length--;
            // Emit token removal event.
            emit RemovedToken(msg.sender, token);
        }
    }


/*
 *    event AddedController(address _sender, address _controller);
 *    event RemovedController(address _sender, address _controller);
 *
 *    mapping(address => bool) private _isController;
 *    uint private _controllerCount;
 *
 *    /// @dev Constructor initializes the list of controllers with the provided address.
 *    /// @param _account address to add to the list of controllers.
 *    constructor(address _account) public {
 *        _addController(_account);
 *    }
 *
 *    /// @dev Checks if message sender is a controller.
 *    modifier onlyController() {
 *        require(isController(msg.sender), "sender is not a controller");
 *        _;
 *    }
 *
 *    /// @dev Add a new controller to the list of controllers.
 *    /// @param _account address to add to the list of controllers.
 *    function addController(address _account) external onlyController {
 *        _addController(_account);
 *    }
 *
 *    /// @dev Remove a controller from the list of controllers.
 *    /// @param _account address to remove from the list of controllers.
 *    function removeController(address _account) external onlyController {
 *        _removeController(_account);
 *    }
 *
 *    /// @return true if the provided account is a controller.
 *    function isController(address _account) public view returns (bool) {
 *        return _isController[_account];
 *    }
 *
 *    /// @return the current number of controllers.
 *    function controllerCount() public view returns (uint) {
 *        return _controllerCount;
 *    }
 *
 *    /// @dev Internal-only function that adds a new controller.
 *    function _addController(address _account) internal {
 *        require(!_isController[_account], "provided account is already a controller");
 *        _isController[_account] = true;
 *        _controllerCount++;
 *        emit AddedController(msg.sender, _account);
 *    }
 *
 *    /// @dev Internal-only function that removes an existing controller.
 *    function _removeController(address _account) internal {
 *        require(_isController[_account], "provided account is not a controller");
 *        require(_controllerCount > 1, "cannot remove the last controller");
 *        _isController[_account] = false;
 *        _controllerCount--;
 *        emit RemovedController(msg.sender, _account);
 *    }
 */
}
