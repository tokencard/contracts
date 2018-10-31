// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package externals

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

// Base64ABI is the input ABI used to generate the binding from.
const Base64ABI = "[]"

// Base64Bin is the compiled bytecode used for deploying new contracts.
const Base64Bin = `6080604052348015600f57600080fd5b50603580601d6000396000f3006080604052600080fd00a165627a7a72305820ad83891fca1f85e22696a33105abe39a6c67f809643162cc031d4af5479d42370029`

// DeployBase64 deploys a new Ethereum contract, binding an instance of Base64 to it.
func DeployBase64(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Base64, error) {
	parsed, err := abi.JSON(strings.NewReader(Base64ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(Base64Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Base64{Base64Caller: Base64Caller{contract: contract}, Base64Transactor: Base64Transactor{contract: contract}, Base64Filterer: Base64Filterer{contract: contract}}, nil
}

// Base64 is an auto generated Go binding around an Ethereum contract.
type Base64 struct {
	Base64Caller     // Read-only binding to the contract
	Base64Transactor // Write-only binding to the contract
	Base64Filterer   // Log filterer for contract events
}

// Base64Caller is an auto generated read-only Go binding around an Ethereum contract.
type Base64Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Base64Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Base64Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Base64Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Base64Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Base64Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Base64Session struct {
	Contract     *Base64           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Base64CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Base64CallerSession struct {
	Contract *Base64Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// Base64TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Base64TransactorSession struct {
	Contract     *Base64Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Base64Raw is an auto generated low-level Go binding around an Ethereum contract.
type Base64Raw struct {
	Contract *Base64 // Generic contract binding to access the raw methods on
}

// Base64CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Base64CallerRaw struct {
	Contract *Base64Caller // Generic read-only contract binding to access the raw methods on
}

// Base64TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Base64TransactorRaw struct {
	Contract *Base64Transactor // Generic write-only contract binding to access the raw methods on
}

// NewBase64 creates a new instance of Base64, bound to a specific deployed contract.
func NewBase64(address common.Address, backend bind.ContractBackend) (*Base64, error) {
	contract, err := bindBase64(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Base64{Base64Caller: Base64Caller{contract: contract}, Base64Transactor: Base64Transactor{contract: contract}, Base64Filterer: Base64Filterer{contract: contract}}, nil
}

// NewBase64Caller creates a new read-only instance of Base64, bound to a specific deployed contract.
func NewBase64Caller(address common.Address, caller bind.ContractCaller) (*Base64Caller, error) {
	contract, err := bindBase64(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Base64Caller{contract: contract}, nil
}

// NewBase64Transactor creates a new write-only instance of Base64, bound to a specific deployed contract.
func NewBase64Transactor(address common.Address, transactor bind.ContractTransactor) (*Base64Transactor, error) {
	contract, err := bindBase64(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Base64Transactor{contract: contract}, nil
}

// NewBase64Filterer creates a new log filterer instance of Base64, bound to a specific deployed contract.
func NewBase64Filterer(address common.Address, filterer bind.ContractFilterer) (*Base64Filterer, error) {
	contract, err := bindBase64(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Base64Filterer{contract: contract}, nil
}

// bindBase64 binds a generic wrapper to an already deployed contract.
func bindBase64(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Base64ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Base64 *Base64Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Base64.Contract.Base64Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Base64 *Base64Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Base64.Contract.Base64Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Base64 *Base64Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Base64.Contract.Base64Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Base64 *Base64CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Base64.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Base64 *Base64TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Base64.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Base64 *Base64TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Base64.Contract.contract.Transact(opts, method, params...)
}
