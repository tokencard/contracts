// ABI encoding extensions for the generated contract bindings.

package bindings

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"strings"
)

// Next is a free data retrieval call binding the contract method 0x4c8fe526.
// Solidity: function next() constant returns(uint256)
func (_Controller *ControllerCaller) NextEncode() ([]byte, error) {
	controllerABI, err := abi.JSON(strings.NewReader(ControllerABI))
	if err != nil {
		return nil, err
	}
	return controllerABI.Pack("next")
}
