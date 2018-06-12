// ABI encoding extensions for the generated contract bindings.

package bindings

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

// Balance is a free data retrieval call binding the contract method 0xe3d670d7.
// Solidity: function balance(token address) constant returns(uint256)
func (_Wallet *WalletCaller) BalanceEncode(asset common.Address) ([]byte, error) {
	walletABI, err := abi.JSON(strings.NewReader(WalletABI))
	if err != nil {
		return nil, err
	}
	return walletABI.Pack("balance", asset)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
// Solidity: function owner() constant returns(address)
func (_Wallet *WalletCaller) OwnerEncode() ([]byte, error) {
	walletABI, err := abi.JSON(strings.NewReader(WalletABI))
	if err != nil {
		return nil, err
	}
	return walletABI.Pack("owner")
}

func (_Wallet *WalletCaller) DailyLimitEncode() ([]byte, error) {
	walletABI, err := abi.JSON(strings.NewReader(WalletABI))
	if err != nil {
		return nil, err
	}
	return walletABI.Pack("dailyLimit")
}

func (_Wallet *WalletCaller) AvailableEncode() ([]byte, error) {
	walletABI, err := abi.JSON(strings.NewReader(WalletABI))
	if err != nil {
		return nil, err
	}
	return walletABI.Pack("available")
}
