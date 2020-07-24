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

pragma solidity ^0.5.17;


interface IGasToken {
    function freeUpTo(uint256) external returns (uint256);
}


contract GasRefundable {
    /// @notice Emits the old and new gas token parameter when changed.
    event SetGasTokenAddress(address _previous, address _new);
    event SetFreeCallGasCost(uint256 _previous, uint256 _new);
    event SetGasRefundPerUnit(uint256 _previous, uint256 _new);

    /// @notice Address of the gas token used to refund gas (default: CHI).
    IGasToken private _gasToken = IGasToken(0x0000000000004946c0e9F43F4Dee607b0eF1fA1c);
    /// @notice Gas token cost and refund parameters used in the gas refund calcualtion (default: CHI).
    uint256 private _freeCallGasCost = 14154;
    uint256 private _gasRefundPerUnit = 41130;

    /// @notice Refunds gas based on the amount of gas spent in the transaction and the gas token parameters.
    modifier refundGas {
        uint256 gasStart = gasleft();
        _;
        uint256 gasSpent = 21000 + gasStart - gasleft() + 16 * msg.data.length;
        _gasToken.freeUpTo((gasSpent + _freeCallGasCost) / _gasRefundPerUnit);
    }

    /// @param _gasTokenAddress Address of the gas token used to refund gas.
    constructor(address _gasTokenAddress) internal {
        if (_gasTokenAddress != address(0)) {
            emit SetGasTokenAddress(address(_gasToken), _gasTokenAddress);
            _gasToken = IGasToken(_gasTokenAddress);
        }
    }

    /// @param _gasTokenAddress Address of the gas token used to refund gas.
    function _setGasToken(address _gasTokenAddress) internal {
        require(_gasTokenAddress != address(0), "gas token address is 0x0");
        emit SetGasTokenAddress(address(_gasToken), _gasTokenAddress);
        _gasToken = IGasToken(_gasTokenAddress);
    }

    /// @param _gasCost Gas cost of the gas token free method call.
    function _setFreeCallGasCost(uint256 _gasCost) internal {
        require(_gasCost != 0, "free call gas cost is 0");
        emit SetFreeCallGasCost(_freeCallGasCost, _gasCost);
        _freeCallGasCost = _gasCost;
    }

    /// @param _gasRefund Amount of gas refunded per unit of gas token.
    function _setGasRefundPerUnit(uint256 _gasRefund) internal {
        require(_gasRefund != 0, "gas refund per unit is 0");
        emit SetGasRefundPerUnit(_gasRefundPerUnit, _gasRefund);
        _gasRefundPerUnit = _gasRefund;
    }

    /// @return Address of the gas token used to refund gas.
    function gasToken() external view returns (address) {
        return address(_gasToken);
    }

    /// @return Gas cost of the gas token free method call.
    function freeCallGasCost() external view returns (uint256) {
        return _freeCallGasCost;
    }

    /// @return Amount of gas refunded per unit of gas token.
    function gasRefundPerUnit() external view returns (uint256) {
        return _gasRefundPerUnit;
    }
}
