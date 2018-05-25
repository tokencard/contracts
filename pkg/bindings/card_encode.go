// ABI encoding extensions for the generated contract bindings.

package bindings

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

// Balance is a free data retrieval call binding the contract method 0xe3d670d7.
// Solidity: function balance(token address) constant returns(uint256)
func (_Card *CardCaller) BalanceEncode(asset common.Address) ([]byte, error) {
	cardABI, err := abi.JSON(strings.NewReader(CardABI))
	if err != nil {
		return nil, err
	}
	return cardABI.Pack("balance", asset)
}

func (_Card *CardCaller) OverdraftEncode(asset common.Address) ([]byte, error) {
	cardABI, err := abi.JSON(strings.NewReader(CardABI))
	if err != nil {
		return nil, err
	}
	return cardABI.Pack("overdraft", asset)
}

func (_Card *CardCaller) AvailableEncode(asset common.Address) ([]byte, error) {
	cardABI, err := abi.JSON(strings.NewReader(CardABI))
	if err != nil {
		return nil, err
	}
	return cardABI.Pack("available", asset)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
// Solidity: function owner() constant returns(address)
func (_Card *CardCaller) OwnerEncode() ([]byte, error) {
	cardABI, err := abi.JSON(strings.NewReader(CardABI))
	if err != nil {
		return nil, err
	}
	return cardABI.Pack("owner")
}

// UnlockAt is a free data retrieval call binding the contract method 0xaa5dec6f.
// Solidity: function unlockAt() constant returns(uint256)
func (_Card *CardCaller) UnlockAtEncode() ([]byte, error) {
	cardABI, err := abi.JSON(strings.NewReader(CardABI))
	if err != nil {
		return nil, err
	}
	return cardABI.Pack("unlockAt")
}
