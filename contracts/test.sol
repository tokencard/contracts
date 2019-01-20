/**
 *  The Consumer Contract Wallet
 *  Copyright (C) 2018 Token Group Ltd
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

import "./licence.sol";

/// @title ERC20 interface is a subset of the ERC20 specification.
interface ERC20 {
    function approve(address, uint256) external returns (bool);
    function balanceOf(address) external view returns (uint);
    function transfer(address, uint) external returns (bool);
    function transferFrom(address _from, address _to, uint256 _value) external returns (bool success);
}

//// @title Asset wallet with extra security features and gas top up management.
contract Test {

    event LoadedTokenCard(uint _amount, uint _licenceFee, address _asset);

    constructor (address licenseAddress) public {
        licence = Licence(licenseAddress);
        
    }

    // Needs this
    Licence public licence;

    function loadTokenCard(uint _amount, uint _licenceFee, address _asset) payable public {
            
        if (_asset != address(0)) {
            require(ERC20(_asset).approve(licence, _amount + _licenceFee), "Failed to approve the proposed ERC20 approval");
            //require(licence.loadTokenCard.value(_amount + _licenceFee)(_amount, _licenceFee, _asset), "failed to send enough ETH");
            require(licence.load(_amount, _licenceFee, _asset), "lame");
        } else {
            require(licence.load.value(_amount + _licenceFee)(_amount, _licenceFee, _asset), "lame");
        }  

        emit LoadedTokenCard(_amount, _licenceFee, _asset);
    }
}

