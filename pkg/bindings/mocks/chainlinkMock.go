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

// ChainlinkMockABI is the input ABI used to generate the binding from.
const ChainlinkMockABI = "[{\"inputs\":[],\"name\":\"latestAnswer\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rate\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"_rate\",\"type\":\"int256\"}],\"name\":\"setLatestAnswer\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ChainlinkMockBin is the compiled bytecode used for deploying new contracts.
var ChainlinkMockBin = "0x608060405234801561001057600080fd5b5060d78061001f6000396000f3fe6080604052348015600f57600080fd5b506004361060465760003560e01c806304ea97b014604b5780632c4e722e14607757806350d25bcd14607d5780638205bf6a146083575b600080fd5b606560048036036020811015605f57600080fd5b50356089565b60408051918252519081900360200190f35b60656091565b60656097565b6065609d565b600090815590565b60005481565b60005490565b429056fea2646970667358221220b2c73f280dcaac6ccfc939515acbea2b796b5cfe441dc892bb2f197477f9336364736f6c63430006040033"

// DeployChainlinkMock deploys a new Ethereum contract, binding an instance of ChainlinkMock to it.
func DeployChainlinkMock(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ChainlinkMock, error) {
	parsed, err := abi.JSON(strings.NewReader(ChainlinkMockABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ChainlinkMockBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ChainlinkMock{ChainlinkMockCaller: ChainlinkMockCaller{contract: contract}, ChainlinkMockTransactor: ChainlinkMockTransactor{contract: contract}, ChainlinkMockFilterer: ChainlinkMockFilterer{contract: contract}}, nil
}

// ChainlinkMock is an auto generated Go binding around an Ethereum contract.
type ChainlinkMock struct {
	ChainlinkMockCaller     // Read-only binding to the contract
	ChainlinkMockTransactor // Write-only binding to the contract
	ChainlinkMockFilterer   // Log filterer for contract events
}

// ChainlinkMockCaller is an auto generated read-only Go binding around an Ethereum contract.
type ChainlinkMockCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChainlinkMockTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ChainlinkMockTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChainlinkMockFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ChainlinkMockFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChainlinkMockSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ChainlinkMockSession struct {
	Contract     *ChainlinkMock    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ChainlinkMockCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ChainlinkMockCallerSession struct {
	Contract *ChainlinkMockCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// ChainlinkMockTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ChainlinkMockTransactorSession struct {
	Contract     *ChainlinkMockTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ChainlinkMockRaw is an auto generated low-level Go binding around an Ethereum contract.
type ChainlinkMockRaw struct {
	Contract *ChainlinkMock // Generic contract binding to access the raw methods on
}

// ChainlinkMockCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ChainlinkMockCallerRaw struct {
	Contract *ChainlinkMockCaller // Generic read-only contract binding to access the raw methods on
}

// ChainlinkMockTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ChainlinkMockTransactorRaw struct {
	Contract *ChainlinkMockTransactor // Generic write-only contract binding to access the raw methods on
}

// NewChainlinkMock creates a new instance of ChainlinkMock, bound to a specific deployed contract.
func NewChainlinkMock(address common.Address, backend bind.ContractBackend) (*ChainlinkMock, error) {
	contract, err := bindChainlinkMock(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ChainlinkMock{ChainlinkMockCaller: ChainlinkMockCaller{contract: contract}, ChainlinkMockTransactor: ChainlinkMockTransactor{contract: contract}, ChainlinkMockFilterer: ChainlinkMockFilterer{contract: contract}}, nil
}

// NewChainlinkMockCaller creates a new read-only instance of ChainlinkMock, bound to a specific deployed contract.
func NewChainlinkMockCaller(address common.Address, caller bind.ContractCaller) (*ChainlinkMockCaller, error) {
	contract, err := bindChainlinkMock(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ChainlinkMockCaller{contract: contract}, nil
}

// NewChainlinkMockTransactor creates a new write-only instance of ChainlinkMock, bound to a specific deployed contract.
func NewChainlinkMockTransactor(address common.Address, transactor bind.ContractTransactor) (*ChainlinkMockTransactor, error) {
	contract, err := bindChainlinkMock(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ChainlinkMockTransactor{contract: contract}, nil
}

// NewChainlinkMockFilterer creates a new log filterer instance of ChainlinkMock, bound to a specific deployed contract.
func NewChainlinkMockFilterer(address common.Address, filterer bind.ContractFilterer) (*ChainlinkMockFilterer, error) {
	contract, err := bindChainlinkMock(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ChainlinkMockFilterer{contract: contract}, nil
}

// bindChainlinkMock binds a generic wrapper to an already deployed contract.
func bindChainlinkMock(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ChainlinkMockABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChainlinkMock *ChainlinkMockRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ChainlinkMock.Contract.ChainlinkMockCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChainlinkMock *ChainlinkMockRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChainlinkMock.Contract.ChainlinkMockTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChainlinkMock *ChainlinkMockRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChainlinkMock.Contract.ChainlinkMockTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChainlinkMock *ChainlinkMockCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ChainlinkMock.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChainlinkMock *ChainlinkMockTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChainlinkMock.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChainlinkMock *ChainlinkMockTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChainlinkMock.Contract.contract.Transact(opts, method, params...)
}

// LatestAnswer is a free data retrieval call binding the contract method 0x50d25bcd.
//
// Solidity: function latestAnswer() constant returns(int256)
func (_ChainlinkMock *ChainlinkMockCaller) LatestAnswer(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ChainlinkMock.contract.Call(opts, out, "latestAnswer")
	return *ret0, err
}

// LatestAnswer is a free data retrieval call binding the contract method 0x50d25bcd.
//
// Solidity: function latestAnswer() constant returns(int256)
func (_ChainlinkMock *ChainlinkMockSession) LatestAnswer() (*big.Int, error) {
	return _ChainlinkMock.Contract.LatestAnswer(&_ChainlinkMock.CallOpts)
}

// LatestAnswer is a free data retrieval call binding the contract method 0x50d25bcd.
//
// Solidity: function latestAnswer() constant returns(int256)
func (_ChainlinkMock *ChainlinkMockCallerSession) LatestAnswer() (*big.Int, error) {
	return _ChainlinkMock.Contract.LatestAnswer(&_ChainlinkMock.CallOpts)
}

// LatestTimestamp is a free data retrieval call binding the contract method 0x8205bf6a.
//
// Solidity: function latestTimestamp() constant returns(uint256)
func (_ChainlinkMock *ChainlinkMockCaller) LatestTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ChainlinkMock.contract.Call(opts, out, "latestTimestamp")
	return *ret0, err
}

// LatestTimestamp is a free data retrieval call binding the contract method 0x8205bf6a.
//
// Solidity: function latestTimestamp() constant returns(uint256)
func (_ChainlinkMock *ChainlinkMockSession) LatestTimestamp() (*big.Int, error) {
	return _ChainlinkMock.Contract.LatestTimestamp(&_ChainlinkMock.CallOpts)
}

// LatestTimestamp is a free data retrieval call binding the contract method 0x8205bf6a.
//
// Solidity: function latestTimestamp() constant returns(uint256)
func (_ChainlinkMock *ChainlinkMockCallerSession) LatestTimestamp() (*big.Int, error) {
	return _ChainlinkMock.Contract.LatestTimestamp(&_ChainlinkMock.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() constant returns(int256)
func (_ChainlinkMock *ChainlinkMockCaller) Rate(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ChainlinkMock.contract.Call(opts, out, "rate")
	return *ret0, err
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() constant returns(int256)
func (_ChainlinkMock *ChainlinkMockSession) Rate() (*big.Int, error) {
	return _ChainlinkMock.Contract.Rate(&_ChainlinkMock.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() constant returns(int256)
func (_ChainlinkMock *ChainlinkMockCallerSession) Rate() (*big.Int, error) {
	return _ChainlinkMock.Contract.Rate(&_ChainlinkMock.CallOpts)
}

// SetLatestAnswer is a paid mutator transaction binding the contract method 0x04ea97b0.
//
// Solidity: function setLatestAnswer(int256 _rate) returns(int256)
func (_ChainlinkMock *ChainlinkMockTransactor) SetLatestAnswer(opts *bind.TransactOpts, _rate *big.Int) (*types.Transaction, error) {
	return _ChainlinkMock.contract.Transact(opts, "setLatestAnswer", _rate)
}

// SetLatestAnswer is a paid mutator transaction binding the contract method 0x04ea97b0.
//
// Solidity: function setLatestAnswer(int256 _rate) returns(int256)
func (_ChainlinkMock *ChainlinkMockSession) SetLatestAnswer(_rate *big.Int) (*types.Transaction, error) {
	return _ChainlinkMock.Contract.SetLatestAnswer(&_ChainlinkMock.TransactOpts, _rate)
}

// SetLatestAnswer is a paid mutator transaction binding the contract method 0x04ea97b0.
//
// Solidity: function setLatestAnswer(int256 _rate) returns(int256)
func (_ChainlinkMock *ChainlinkMockTransactorSession) SetLatestAnswer(_rate *big.Int) (*types.Transaction, error) {
	return _ChainlinkMock.Contract.SetLatestAnswer(&_ChainlinkMock.TransactOpts, _rate)
}
