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

// OraclizeAddrResolverABI is the input ABI used to generate the binding from.
const OraclizeAddrResolverABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oraclizedAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"getAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// OraclizeAddrResolverBin is the compiled bytecode used for deploying new contracts.
var OraclizeAddrResolverBin = "0x608060405234801561001057600080fd5b506040516100f83803806100f88339818101604052602081101561003357600080fd5b5051600080546001600160a01b039092166001600160a01b03199092169190911790556094806100646000396000f3fe6080604052348015600f57600080fd5b506004361060285760003560e01c806338cc483114602d575b600080fd5b6033604f565b604080516001600160a01b039092168252519081900360200190f35b6000546001600160a01b03169056fea26469706673582212209709156c371a46905a4870e73c244b9c9020067422234abfd8ee692b1081139d64736f6c634300060b0033"

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

// GetAddress is a free data retrieval call binding the contract method 0x38cc4831.
//
// Solidity: function getAddress() constant returns(address)
func (_OraclizeAddrResolver *OraclizeAddrResolverCaller) GetAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _OraclizeAddrResolver.contract.Call(opts, out, "getAddress")
	return *ret0, err
}

// GetAddress is a free data retrieval call binding the contract method 0x38cc4831.
//
// Solidity: function getAddress() constant returns(address)
func (_OraclizeAddrResolver *OraclizeAddrResolverSession) GetAddress() (common.Address, error) {
	return _OraclizeAddrResolver.Contract.GetAddress(&_OraclizeAddrResolver.CallOpts)
}

// GetAddress is a free data retrieval call binding the contract method 0x38cc4831.
//
// Solidity: function getAddress() constant returns(address)
func (_OraclizeAddrResolver *OraclizeAddrResolverCallerSession) GetAddress() (common.Address, error) {
	return _OraclizeAddrResolver.Contract.GetAddress(&_OraclizeAddrResolver.CallOpts)
}
