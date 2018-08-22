// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// MockOraclizeAddrResolverABI is the input ABI used to generate the binding from.
const MockOraclizeAddrResolverABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"getAddress\",\"outputs\":[{\"name\":\"_addr\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_oraclizedAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// MockOraclizeAddrResolverBin is the compiled bytecode used for deploying new contracts.
const MockOraclizeAddrResolverBin = `0x608060405234801561001057600080fd5b5060405160208061016f83398101806040528101908080519060200190929190505050806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505060ed806100826000396000f300608060405260043610603f576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806338cc4831146044575b600080fd5b348015604f57600080fd5b5060566098565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050905600a165627a7a72305820b041885935493ad488fe9ce431ef593aac654c506239bc2a027343c395fd4a1c0029`

// DeployMockOraclizeAddrResolver deploys a new Ethereum contract, binding an instance of MockOraclizeAddrResolver to it.
func DeployMockOraclizeAddrResolver(auth *bind.TransactOpts, backend bind.ContractBackend, _oraclizedAddress common.Address) (common.Address, *types.Transaction, *MockOraclizeAddrResolver, error) {
	parsed, err := abi.JSON(strings.NewReader(MockOraclizeAddrResolverABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MockOraclizeAddrResolverBin), backend, _oraclizedAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MockOraclizeAddrResolver{MockOraclizeAddrResolverCaller: MockOraclizeAddrResolverCaller{contract: contract}, MockOraclizeAddrResolverTransactor: MockOraclizeAddrResolverTransactor{contract: contract}, MockOraclizeAddrResolverFilterer: MockOraclizeAddrResolverFilterer{contract: contract}}, nil
}

// MockOraclizeAddrResolver is an auto generated Go binding around an Ethereum contract.
type MockOraclizeAddrResolver struct {
	MockOraclizeAddrResolverCaller     // Read-only binding to the contract
	MockOraclizeAddrResolverTransactor // Write-only binding to the contract
	MockOraclizeAddrResolverFilterer   // Log filterer for contract events
}

// MockOraclizeAddrResolverCaller is an auto generated read-only Go binding around an Ethereum contract.
type MockOraclizeAddrResolverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockOraclizeAddrResolverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MockOraclizeAddrResolverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockOraclizeAddrResolverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MockOraclizeAddrResolverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockOraclizeAddrResolverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MockOraclizeAddrResolverSession struct {
	Contract     *MockOraclizeAddrResolver // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// MockOraclizeAddrResolverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MockOraclizeAddrResolverCallerSession struct {
	Contract *MockOraclizeAddrResolverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// MockOraclizeAddrResolverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MockOraclizeAddrResolverTransactorSession struct {
	Contract     *MockOraclizeAddrResolverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// MockOraclizeAddrResolverRaw is an auto generated low-level Go binding around an Ethereum contract.
type MockOraclizeAddrResolverRaw struct {
	Contract *MockOraclizeAddrResolver // Generic contract binding to access the raw methods on
}

// MockOraclizeAddrResolverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MockOraclizeAddrResolverCallerRaw struct {
	Contract *MockOraclizeAddrResolverCaller // Generic read-only contract binding to access the raw methods on
}

// MockOraclizeAddrResolverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MockOraclizeAddrResolverTransactorRaw struct {
	Contract *MockOraclizeAddrResolverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMockOraclizeAddrResolver creates a new instance of MockOraclizeAddrResolver, bound to a specific deployed contract.
func NewMockOraclizeAddrResolver(address common.Address, backend bind.ContractBackend) (*MockOraclizeAddrResolver, error) {
	contract, err := bindMockOraclizeAddrResolver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MockOraclizeAddrResolver{MockOraclizeAddrResolverCaller: MockOraclizeAddrResolverCaller{contract: contract}, MockOraclizeAddrResolverTransactor: MockOraclizeAddrResolverTransactor{contract: contract}, MockOraclizeAddrResolverFilterer: MockOraclizeAddrResolverFilterer{contract: contract}}, nil
}

// NewMockOraclizeAddrResolverCaller creates a new read-only instance of MockOraclizeAddrResolver, bound to a specific deployed contract.
func NewMockOraclizeAddrResolverCaller(address common.Address, caller bind.ContractCaller) (*MockOraclizeAddrResolverCaller, error) {
	contract, err := bindMockOraclizeAddrResolver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MockOraclizeAddrResolverCaller{contract: contract}, nil
}

// NewMockOraclizeAddrResolverTransactor creates a new write-only instance of MockOraclizeAddrResolver, bound to a specific deployed contract.
func NewMockOraclizeAddrResolverTransactor(address common.Address, transactor bind.ContractTransactor) (*MockOraclizeAddrResolverTransactor, error) {
	contract, err := bindMockOraclizeAddrResolver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MockOraclizeAddrResolverTransactor{contract: contract}, nil
}

// NewMockOraclizeAddrResolverFilterer creates a new log filterer instance of MockOraclizeAddrResolver, bound to a specific deployed contract.
func NewMockOraclizeAddrResolverFilterer(address common.Address, filterer bind.ContractFilterer) (*MockOraclizeAddrResolverFilterer, error) {
	contract, err := bindMockOraclizeAddrResolver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MockOraclizeAddrResolverFilterer{contract: contract}, nil
}

// bindMockOraclizeAddrResolver binds a generic wrapper to an already deployed contract.
func bindMockOraclizeAddrResolver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MockOraclizeAddrResolverABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MockOraclizeAddrResolver *MockOraclizeAddrResolverRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MockOraclizeAddrResolver.Contract.MockOraclizeAddrResolverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MockOraclizeAddrResolver *MockOraclizeAddrResolverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockOraclizeAddrResolver.Contract.MockOraclizeAddrResolverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MockOraclizeAddrResolver *MockOraclizeAddrResolverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockOraclizeAddrResolver.Contract.MockOraclizeAddrResolverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MockOraclizeAddrResolver *MockOraclizeAddrResolverCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MockOraclizeAddrResolver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MockOraclizeAddrResolver *MockOraclizeAddrResolverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockOraclizeAddrResolver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MockOraclizeAddrResolver *MockOraclizeAddrResolverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockOraclizeAddrResolver.Contract.contract.Transact(opts, method, params...)
}

// GetAddress is a paid mutator transaction binding the contract method 0x38cc4831.
//
// Solidity: function getAddress() returns(_addr address)
func (_MockOraclizeAddrResolver *MockOraclizeAddrResolverTransactor) GetAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockOraclizeAddrResolver.contract.Transact(opts, "getAddress")
}

// GetAddress is a paid mutator transaction binding the contract method 0x38cc4831.
//
// Solidity: function getAddress() returns(_addr address)
func (_MockOraclizeAddrResolver *MockOraclizeAddrResolverSession) GetAddress() (*types.Transaction, error) {
	return _MockOraclizeAddrResolver.Contract.GetAddress(&_MockOraclizeAddrResolver.TransactOpts)
}

// GetAddress is a paid mutator transaction binding the contract method 0x38cc4831.
//
// Solidity: function getAddress() returns(_addr address)
func (_MockOraclizeAddrResolver *MockOraclizeAddrResolverTransactorSession) GetAddress() (*types.Transaction, error) {
	return _MockOraclizeAddrResolver.Contract.GetAddress(&_MockOraclizeAddrResolver.TransactOpts)
}
