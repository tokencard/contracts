// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package mocks

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// OraclizeAddrResolverABI is the input ABI used to generate the binding from.
const OraclizeAddrResolverABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"getAddress\",\"outputs\":[{\"name\":\"_addr\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_oraclizedAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// OraclizeAddrResolverBin is the compiled bytecode used for deploying new contracts.
const OraclizeAddrResolverBin = `608060405234801561001057600080fd5b5060405160208061016f83398101806040528101908080519060200190929190505050806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505060ed806100826000396000f300608060405260043610603f576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806338cc4831146044575b600080fd5b348015604f57600080fd5b5060566098565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050905600a165627a7a72305820a536e0d556238e7fefb0e102085018d3f14bf444292ff49a0494a55a3dd72ee80029`

// DeployOraclizeAddrResolver deploys a new Ethereum contract, binding an instance of OraclizeAddrResolver to it.
func DeployOraclizeAddrResolver(auth *bind.TransactOpts, backend bind.ContractBackend, _oraclizedAddress common.Address) (common.Address, *types.Transaction, *OraclizeAddrResolver, error) {
	parsed, err := abi.JSON(strings.NewReader(OraclizeAddrResolverABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OraclizeAddrResolverBin), backend, _oraclizedAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OraclizeAddrResolver{OraclizeAddrResolverCaller: OraclizeAddrResolverCaller{contract: contract}, OraclizeAddrResolverTransactor: OraclizeAddrResolverTransactor{contract: contract}, OraclizeAddrResolverFilterer: OraclizeAddrResolverFilterer{contract: contract}}, nil
}

// OraclizeAddrResolver is an auto generated Go binding around an Ethereum contract.
type OraclizeAddrResolver struct {
	OraclizeAddrResolverCaller     // Read-only binding to the contract
	OraclizeAddrResolverTransactor // Write-only binding to the contract
	OraclizeAddrResolverFilterer   // Log filterer for contract events
}

// OraclizeAddrResolverCaller is an auto generated read-only Go binding around an Ethereum contract.
type OraclizeAddrResolverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OraclizeAddrResolverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OraclizeAddrResolverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OraclizeAddrResolverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OraclizeAddrResolverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OraclizeAddrResolverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OraclizeAddrResolverSession struct {
	Contract     *OraclizeAddrResolver // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// OraclizeAddrResolverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OraclizeAddrResolverCallerSession struct {
	Contract *OraclizeAddrResolverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// OraclizeAddrResolverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OraclizeAddrResolverTransactorSession struct {
	Contract     *OraclizeAddrResolverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// OraclizeAddrResolverRaw is an auto generated low-level Go binding around an Ethereum contract.
type OraclizeAddrResolverRaw struct {
	Contract *OraclizeAddrResolver // Generic contract binding to access the raw methods on
}

// OraclizeAddrResolverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OraclizeAddrResolverCallerRaw struct {
	Contract *OraclizeAddrResolverCaller // Generic read-only contract binding to access the raw methods on
}

// OraclizeAddrResolverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OraclizeAddrResolverTransactorRaw struct {
	Contract *OraclizeAddrResolverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOraclizeAddrResolver creates a new instance of OraclizeAddrResolver, bound to a specific deployed contract.
func NewOraclizeAddrResolver(address common.Address, backend bind.ContractBackend) (*OraclizeAddrResolver, error) {
	contract, err := bindOraclizeAddrResolver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OraclizeAddrResolver{OraclizeAddrResolverCaller: OraclizeAddrResolverCaller{contract: contract}, OraclizeAddrResolverTransactor: OraclizeAddrResolverTransactor{contract: contract}, OraclizeAddrResolverFilterer: OraclizeAddrResolverFilterer{contract: contract}}, nil
}

// NewOraclizeAddrResolverCaller creates a new read-only instance of OraclizeAddrResolver, bound to a specific deployed contract.
func NewOraclizeAddrResolverCaller(address common.Address, caller bind.ContractCaller) (*OraclizeAddrResolverCaller, error) {
	contract, err := bindOraclizeAddrResolver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OraclizeAddrResolverCaller{contract: contract}, nil
}

// NewOraclizeAddrResolverTransactor creates a new write-only instance of OraclizeAddrResolver, bound to a specific deployed contract.
func NewOraclizeAddrResolverTransactor(address common.Address, transactor bind.ContractTransactor) (*OraclizeAddrResolverTransactor, error) {
	contract, err := bindOraclizeAddrResolver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OraclizeAddrResolverTransactor{contract: contract}, nil
}

// NewOraclizeAddrResolverFilterer creates a new log filterer instance of OraclizeAddrResolver, bound to a specific deployed contract.
func NewOraclizeAddrResolverFilterer(address common.Address, filterer bind.ContractFilterer) (*OraclizeAddrResolverFilterer, error) {
	contract, err := bindOraclizeAddrResolver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OraclizeAddrResolverFilterer{contract: contract}, nil
}

// bindOraclizeAddrResolver binds a generic wrapper to an already deployed contract.
func bindOraclizeAddrResolver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OraclizeAddrResolverABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OraclizeAddrResolver *OraclizeAddrResolverRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _OraclizeAddrResolver.Contract.OraclizeAddrResolverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OraclizeAddrResolver *OraclizeAddrResolverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OraclizeAddrResolver.Contract.OraclizeAddrResolverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OraclizeAddrResolver *OraclizeAddrResolverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OraclizeAddrResolver.Contract.OraclizeAddrResolverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OraclizeAddrResolver *OraclizeAddrResolverCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _OraclizeAddrResolver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OraclizeAddrResolver *OraclizeAddrResolverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OraclizeAddrResolver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OraclizeAddrResolver *OraclizeAddrResolverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OraclizeAddrResolver.Contract.contract.Transact(opts, method, params...)
}

// GetAddress is a paid mutator transaction binding the contract method 0x38cc4831.
//
// Solidity: function getAddress() returns(_addr address)
func (_OraclizeAddrResolver *OraclizeAddrResolverTransactor) GetAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OraclizeAddrResolver.contract.Transact(opts, "getAddress")
}

// GetAddress is a paid mutator transaction binding the contract method 0x38cc4831.
//
// Solidity: function getAddress() returns(_addr address)
func (_OraclizeAddrResolver *OraclizeAddrResolverSession) GetAddress() (*types.Transaction, error) {
	return _OraclizeAddrResolver.Contract.GetAddress(&_OraclizeAddrResolver.TransactOpts)
}

// GetAddress is a paid mutator transaction binding the contract method 0x38cc4831.
//
// Solidity: function getAddress() returns(_addr address)
func (_OraclizeAddrResolver *OraclizeAddrResolverTransactorSession) GetAddress() (*types.Transaction, error) {
	return _OraclizeAddrResolver.Contract.GetAddress(&_OraclizeAddrResolver.TransactOpts)
}
