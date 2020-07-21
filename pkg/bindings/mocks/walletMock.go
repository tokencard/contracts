// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package mocks

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// WalletMockABI is the input ABI used to generate the binding from.
const WalletMockABI = "[{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"sendValue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// WalletMockBin is the compiled bytecode used for deploying new contracts.
var WalletMockBin = "0x608060405234801561001057600080fd5b506101b7806100206000396000f3fe60806040526004361061002d5760003560e01c806324a084df14610039578063a9059cbb1461007457610034565b3661003457005b600080fd5b34801561004557600080fd5b506100726004803603604081101561005c57600080fd5b506001600160a01b0381351690602001356100ad565b005b34801561008057600080fd5b506100726004803603604081101561009757600080fd5b506001600160a01b03813516906020013561014b565b6040516000906001600160a01b0384169083908381818185875af1925050503d80600081146100f8576040519150601f19603f3d011682016040523d82523d6000602084013e6100fd565b606091505b5050905080610146576040805162461bcd60e51b815260206004820152601060248201526f1cd95b9915985b1d594819985a5b195960821b604482015290519081900360640190fd5b505050565b6040516001600160a01b0383169082156108fc029083906000818181858888f19350505050158015610146573d6000803e3d6000fdfea2646970667358221220ed15540c52d6a01316c70cfeba227cbd6573233ac82a7c9a4aea6e52f3809ad664736f6c634300060b0033"

// DeployWalletMock deploys a new Ethereum contract, binding an instance of WalletMock to it.
func DeployWalletMock(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *WalletMock, error) {
	parsed, err := abi.JSON(strings.NewReader(WalletMockABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(WalletMockBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &WalletMock{WalletMockCaller: WalletMockCaller{contract: contract}, WalletMockTransactor: WalletMockTransactor{contract: contract}, WalletMockFilterer: WalletMockFilterer{contract: contract}}, nil
}

// WalletMock is an auto generated Go binding around an Ethereum contract.
type WalletMock struct {
	WalletMockCaller     // Read-only binding to the contract
	WalletMockTransactor // Write-only binding to the contract
	WalletMockFilterer   // Log filterer for contract events
}

// WalletMockCaller is an auto generated read-only Go binding around an Ethereum contract.
type WalletMockCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletMockTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WalletMockTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletMockFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WalletMockFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletMockSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WalletMockSession struct {
	Contract     *WalletMock       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WalletMockCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WalletMockCallerSession struct {
	Contract *WalletMockCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// WalletMockTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WalletMockTransactorSession struct {
	Contract     *WalletMockTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// WalletMockRaw is an auto generated low-level Go binding around an Ethereum contract.
type WalletMockRaw struct {
	Contract *WalletMock // Generic contract binding to access the raw methods on
}

// WalletMockCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WalletMockCallerRaw struct {
	Contract *WalletMockCaller // Generic read-only contract binding to access the raw methods on
}

// WalletMockTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WalletMockTransactorRaw struct {
	Contract *WalletMockTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWalletMock creates a new instance of WalletMock, bound to a specific deployed contract.
func NewWalletMock(address common.Address, backend bind.ContractBackend) (*WalletMock, error) {
	contract, err := bindWalletMock(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WalletMock{WalletMockCaller: WalletMockCaller{contract: contract}, WalletMockTransactor: WalletMockTransactor{contract: contract}, WalletMockFilterer: WalletMockFilterer{contract: contract}}, nil
}

// NewWalletMockCaller creates a new read-only instance of WalletMock, bound to a specific deployed contract.
func NewWalletMockCaller(address common.Address, caller bind.ContractCaller) (*WalletMockCaller, error) {
	contract, err := bindWalletMock(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WalletMockCaller{contract: contract}, nil
}

// NewWalletMockTransactor creates a new write-only instance of WalletMock, bound to a specific deployed contract.
func NewWalletMockTransactor(address common.Address, transactor bind.ContractTransactor) (*WalletMockTransactor, error) {
	contract, err := bindWalletMock(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WalletMockTransactor{contract: contract}, nil
}

// NewWalletMockFilterer creates a new log filterer instance of WalletMock, bound to a specific deployed contract.
func NewWalletMockFilterer(address common.Address, filterer bind.ContractFilterer) (*WalletMockFilterer, error) {
	contract, err := bindWalletMock(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WalletMockFilterer{contract: contract}, nil
}

// bindWalletMock binds a generic wrapper to an already deployed contract.
func bindWalletMock(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(WalletMockABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WalletMock *WalletMockRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _WalletMock.Contract.WalletMockCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WalletMock *WalletMockRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WalletMock.Contract.WalletMockTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WalletMock *WalletMockRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WalletMock.Contract.WalletMockTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WalletMock *WalletMockCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _WalletMock.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WalletMock *WalletMockTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WalletMock.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WalletMock *WalletMockTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WalletMock.Contract.contract.Transact(opts, method, params...)
}

// SendValue is a paid mutator transaction binding the contract method 0x24a084df.
//
// Solidity: function sendValue(address _to, uint256 _amount) returns()
func (_WalletMock *WalletMockTransactor) SendValue(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _WalletMock.contract.Transact(opts, "sendValue", _to, _amount)
}

// SendValue is a paid mutator transaction binding the contract method 0x24a084df.
//
// Solidity: function sendValue(address _to, uint256 _amount) returns()
func (_WalletMock *WalletMockSession) SendValue(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _WalletMock.Contract.SendValue(&_WalletMock.TransactOpts, _to, _amount)
}

// SendValue is a paid mutator transaction binding the contract method 0x24a084df.
//
// Solidity: function sendValue(address _to, uint256 _amount) returns()
func (_WalletMock *WalletMockTransactorSession) SendValue(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _WalletMock.Contract.SendValue(&_WalletMock.TransactOpts, _to, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _amount) returns()
func (_WalletMock *WalletMockTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _WalletMock.contract.Transact(opts, "transfer", _to, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _amount) returns()
func (_WalletMock *WalletMockSession) Transfer(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _WalletMock.Contract.Transfer(&_WalletMock.TransactOpts, _to, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _amount) returns()
func (_WalletMock *WalletMockTransactorSession) Transfer(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _WalletMock.Contract.Transfer(&_WalletMock.TransactOpts, _to, _amount)
}
