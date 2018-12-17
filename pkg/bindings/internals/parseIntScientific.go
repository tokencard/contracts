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

// ParseIntScientificABI is the input ABI used to generate the binding from.
const ParseIntScientificABI = "[]"

// ParseIntScientificBin is the compiled bytecode used for deploying new contracts.
const ParseIntScientificBin = `6080604052348015600f57600080fd5b50603580601d6000396000f3006080604052600080fd00a165627a7a72305820061ffa196e1380d4e53eb49cb4dc26b3e01f09d6536809c7cc73f9ff6c393e8f0029`

// DeployParseIntScientific deploys a new Ethereum contract, binding an instance of ParseIntScientific to it.
func DeployParseIntScientific(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ParseIntScientific, error) {
	parsed, err := abi.JSON(strings.NewReader(ParseIntScientificABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ParseIntScientificBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ParseIntScientific{ParseIntScientificCaller: ParseIntScientificCaller{contract: contract}, ParseIntScientificTransactor: ParseIntScientificTransactor{contract: contract}, ParseIntScientificFilterer: ParseIntScientificFilterer{contract: contract}}, nil
}

// ParseIntScientific is an auto generated Go binding around an Ethereum contract.
type ParseIntScientific struct {
	ParseIntScientificCaller     // Read-only binding to the contract
	ParseIntScientificTransactor // Write-only binding to the contract
	ParseIntScientificFilterer   // Log filterer for contract events
}

// ParseIntScientificCaller is an auto generated read-only Go binding around an Ethereum contract.
type ParseIntScientificCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ParseIntScientificTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ParseIntScientificTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ParseIntScientificFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ParseIntScientificFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ParseIntScientificSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ParseIntScientificSession struct {
	Contract     *ParseIntScientific // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ParseIntScientificCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ParseIntScientificCallerSession struct {
	Contract *ParseIntScientificCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// ParseIntScientificTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ParseIntScientificTransactorSession struct {
	Contract     *ParseIntScientificTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// ParseIntScientificRaw is an auto generated low-level Go binding around an Ethereum contract.
type ParseIntScientificRaw struct {
	Contract *ParseIntScientific // Generic contract binding to access the raw methods on
}

// ParseIntScientificCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ParseIntScientificCallerRaw struct {
	Contract *ParseIntScientificCaller // Generic read-only contract binding to access the raw methods on
}

// ParseIntScientificTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ParseIntScientificTransactorRaw struct {
	Contract *ParseIntScientificTransactor // Generic write-only contract binding to access the raw methods on
}

// NewParseIntScientific creates a new instance of ParseIntScientific, bound to a specific deployed contract.
func NewParseIntScientific(address common.Address, backend bind.ContractBackend) (*ParseIntScientific, error) {
	contract, err := bindParseIntScientific(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ParseIntScientific{ParseIntScientificCaller: ParseIntScientificCaller{contract: contract}, ParseIntScientificTransactor: ParseIntScientificTransactor{contract: contract}, ParseIntScientificFilterer: ParseIntScientificFilterer{contract: contract}}, nil
}

// NewParseIntScientificCaller creates a new read-only instance of ParseIntScientific, bound to a specific deployed contract.
func NewParseIntScientificCaller(address common.Address, caller bind.ContractCaller) (*ParseIntScientificCaller, error) {
	contract, err := bindParseIntScientific(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ParseIntScientificCaller{contract: contract}, nil
}

// NewParseIntScientificTransactor creates a new write-only instance of ParseIntScientific, bound to a specific deployed contract.
func NewParseIntScientificTransactor(address common.Address, transactor bind.ContractTransactor) (*ParseIntScientificTransactor, error) {
	contract, err := bindParseIntScientific(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ParseIntScientificTransactor{contract: contract}, nil
}

// NewParseIntScientificFilterer creates a new log filterer instance of ParseIntScientific, bound to a specific deployed contract.
func NewParseIntScientificFilterer(address common.Address, filterer bind.ContractFilterer) (*ParseIntScientificFilterer, error) {
	contract, err := bindParseIntScientific(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ParseIntScientificFilterer{contract: contract}, nil
}

// bindParseIntScientific binds a generic wrapper to an already deployed contract.
func bindParseIntScientific(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ParseIntScientificABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ParseIntScientific *ParseIntScientificRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ParseIntScientific.Contract.ParseIntScientificCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ParseIntScientific *ParseIntScientificRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ParseIntScientific.Contract.ParseIntScientificTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ParseIntScientific *ParseIntScientificRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ParseIntScientific.Contract.ParseIntScientificTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ParseIntScientific *ParseIntScientificCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ParseIntScientific.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ParseIntScientific *ParseIntScientificTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ParseIntScientific.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ParseIntScientific *ParseIntScientificTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ParseIntScientific.Contract.contract.Transact(opts, method, params...)
}
