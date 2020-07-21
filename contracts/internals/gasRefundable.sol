/**
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
pragma experimental ABIEncoderV2;


interface IGasToken {
    function freeUpTo(uint256) external returns (uint256);
}


contract GasRefundable {
    /// @notice Emits the new gas token information when it is set.
    event SetGasToken(address _gasTokenAddress, GasTokenParameters _gasTokenParameters);

    struct GasTokenParameters {
        uint256 freeCallGasCost;
        uint256 gasRefundPerUnit;
    }

    /// @dev Address of the gas token used to refund gas (default: CHI).
    IGasToken private _gasToken = IGasToken(0x0000000000004946c0e9F43F4Dee607b0eF1fA1c);
    /// @dev Gas token parameters parameters used in the gas refund calcualtion (default: CHI).
    GasTokenParameters private _gasTokenParameters = GasTokenParameters({freeCallGasCost: 14154, gasRefundPerUnit: 41130});

    /// @notice Refunds gas based on the amount of gas spent in the transaction and the gas token parameters.
    modifier refundGas() {
        uint256 gasStart = gasleft();
        _;
        uint256 gasSpent = 21000 + gasStart - gasleft() + 16 * msg.data.length;
        _gasToken.freeUpTo((gasSpent + _gasTokenParameters.freeCallGasCost) / _gasTokenParameters.gasRefundPerUnit);
    }

    /// @param _gasTokenAddress Address of the gas token used to refund gas.
    /// @param _parameters Gas cost of the gas token free method call and amount of gas refunded per unit of gas token.
    function _setGasToken(address _gasTokenAddress, GasTokenParameters memory _parameters) internal {
        require(_gasTokenAddress != address(0), "gas token address is 0x0");
        require(_parameters.freeCallGasCost != 0, "free call gas cost is 0");
        require(_parameters.gasRefundPerUnit != 0, "gas refund per unit is 0");
        _gasToken = IGasToken(_gasTokenAddress);
        _gasTokenParameters.freeCallGasCost = _parameters.freeCallGasCost;
        _gasTokenParameters.gasRefundPerUnit = _parameters.gasRefundPerUnit;
        emit SetGasToken(_gasTokenAddress, _parameters);
    }

    /// @return Address of the gas token used to refund gas.
    function gasToken() external view returns (address) {
        return address(_gasToken);
    }

    /// @return Gas cost of the gas token free method call/Amount of gas refunded per unit of gas token.
    function gasTokenParameters() external view returns (GasTokenParameters memory) {
        return _gasTokenParameters;
    }
}
