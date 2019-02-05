// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package internals

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

// TokenWhitelistableABI is the input ABI used to generate the binding from.
const TokenWhitelistableABI = "[{\"inputs\":[{\"name\":\"_ens\",\"type\":\"address\"},{\"name\":\"_tokenWhitelistName\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// TokenWhitelistable is an auto generated Go binding around an Ethereum contract.
type TokenWhitelistable struct {
	TokenWhitelistableCaller     // Read-only binding to the contract
	TokenWhitelistableTransactor // Write-only binding to the contract
	TokenWhitelistableFilterer   // Log filterer for contract events
}

// TokenWhitelistableCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenWhitelistableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenWhitelistableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenWhitelistableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenWhitelistableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenWhitelistableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenWhitelistableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenWhitelistableSession struct {
	Contract     *TokenWhitelistable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// TokenWhitelistableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenWhitelistableCallerSession struct {
	Contract *TokenWhitelistableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// TokenWhitelistableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenWhitelistableTransactorSession struct {
	Contract     *TokenWhitelistableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// TokenWhitelistableRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenWhitelistableRaw struct {
	Contract *TokenWhitelistable // Generic contract binding to access the raw methods on
}

// TokenWhitelistableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenWhitelistableCallerRaw struct {
	Contract *TokenWhitelistableCaller // Generic read-only contract binding to access the raw methods on
}

// TokenWhitelistableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenWhitelistableTransactorRaw struct {
	Contract *TokenWhitelistableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTokenWhitelistable creates a new instance of TokenWhitelistable, bound to a specific deployed contract.
func NewTokenWhitelistable(address common.Address, backend bind.ContractBackend) (*TokenWhitelistable, error) {
	contract, err := bindTokenWhitelistable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TokenWhitelistable{TokenWhitelistableCaller: TokenWhitelistableCaller{contract: contract}, TokenWhitelistableTransactor: TokenWhitelistableTransactor{contract: contract}, TokenWhitelistableFilterer: TokenWhitelistableFilterer{contract: contract}}, nil
}

// NewTokenWhitelistableCaller creates a new read-only instance of TokenWhitelistable, bound to a specific deployed contract.
func NewTokenWhitelistableCaller(address common.Address, caller bind.ContractCaller) (*TokenWhitelistableCaller, error) {
	contract, err := bindTokenWhitelistable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenWhitelistableCaller{contract: contract}, nil
}

// NewTokenWhitelistableTransactor creates a new write-only instance of TokenWhitelistable, bound to a specific deployed contract.
func NewTokenWhitelistableTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenWhitelistableTransactor, error) {
	contract, err := bindTokenWhitelistable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenWhitelistableTransactor{contract: contract}, nil
}

// NewTokenWhitelistableFilterer creates a new log filterer instance of TokenWhitelistable, bound to a specific deployed contract.
func NewTokenWhitelistableFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenWhitelistableFilterer, error) {
	contract, err := bindTokenWhitelistable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenWhitelistableFilterer{contract: contract}, nil
}

// bindTokenWhitelistable binds a generic wrapper to an already deployed contract.
func bindTokenWhitelistable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenWhitelistableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenWhitelistable *TokenWhitelistableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TokenWhitelistable.Contract.TokenWhitelistableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenWhitelistable *TokenWhitelistableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenWhitelistable.Contract.TokenWhitelistableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenWhitelistable *TokenWhitelistableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenWhitelistable.Contract.TokenWhitelistableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenWhitelistable *TokenWhitelistableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TokenWhitelistable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenWhitelistable *TokenWhitelistableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenWhitelistable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenWhitelistable *TokenWhitelistableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenWhitelistable.Contract.contract.Transact(opts, method, params...)
}
