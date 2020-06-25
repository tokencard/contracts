pragma solidity 0.5.17;


contract WalletMock {


    /// @dev Ether can be deposited from any source, so this contract must be payable by anyone.
    function() external payable {
    }

    
   
    function transfer(address payable _to, uint256 _amount) external {
       _to.transfer(_amount);
    }

    function sendValue(address payable _to, uint256 _amount) external {
        (bool success, ) = _to.call.value(_amount)("");
       require(success, "sendValue failed");
    }

}
