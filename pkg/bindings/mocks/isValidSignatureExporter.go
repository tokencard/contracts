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

// IsValidSignatureExporterABI is the input ABI used to generate the binding from.
const IsValidSignatureExporterABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_wallet\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"isValidSignature\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// IsValidSignatureExporterBin is the compiled bytecode used for deploying new contracts.
var IsValidSignatureExporterBin = "0x608060405234801561001057600080fd5b506040516102833803806102838339818101604052602081101561003357600080fd5b5051600080546001600160a01b039092166001600160a01b031990921691909117905561021e806100656000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c806320c13b0b14610030575b600080fd5b6100f26004803603604081101561004657600080fd5b81019060208101813564010000000081111561006157600080fd5b82018360208201111561007357600080fd5b8035906020019184600183028401116401000000008311171561009557600080fd5b9193909290916020810190356401000000008111156100b357600080fd5b8201836020820111156100c557600080fd5b803590602001918460018302840111640100000000831117156100e757600080fd5b50909250905061010f565b604080516001600160e01b03199092168252519081900360200190f35b60008054604080516320c13b0b60e01b815260048101918252604481018790526001600160a01b03909216916320c13b0b9188918891889188919081906024810190606401878780828437600083820152601f01601f191690910184810383528581526020019050858580828437600081840152601f19601f820116905080830192505050965050505050505060206040518083038186803b1580156101b457600080fd5b505afa1580156101c8573d6000803e3d6000fd5b505050506040513d60208110156101de57600080fd5b50519594505050505056fea265627a7a72315820abdde881545f5e96fccd7300631368d650ba16dae6009abe72ad52c81ed67e3464736f6c634300050f0032"

// DeployIsValidSignatureExporter deploys a new Ethereum contract, binding an instance of IsValidSignatureExporter to it.
func DeployIsValidSignatureExporter(auth *bind.TransactOpts, backend bind.ContractBackend, _wallet common.Address) (common.Address, *types.Transaction, *IsValidSignatureExporter, error) {
	parsed, err := abi.JSON(strings.NewReader(IsValidSignatureExporterABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(IsValidSignatureExporterBin), backend, _wallet)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &IsValidSignatureExporter{IsValidSignatureExporterCaller: IsValidSignatureExporterCaller{contract: contract}, IsValidSignatureExporterTransactor: IsValidSignatureExporterTransactor{contract: contract}, IsValidSignatureExporterFilterer: IsValidSignatureExporterFilterer{contract: contract}}, nil
}

// IsValidSignatureExporter is an auto generated Go binding around an Ethereum contract.
type IsValidSignatureExporter struct {
	IsValidSignatureExporterCaller     // Read-only binding to the contract
	IsValidSignatureExporterTransactor // Write-only binding to the contract
	IsValidSignatureExporterFilterer   // Log filterer for contract events
}

// IsValidSignatureExporterCaller is an auto generated read-only Go binding around an Ethereum contract.
type IsValidSignatureExporterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IsValidSignatureExporterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IsValidSignatureExporterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IsValidSignatureExporterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IsValidSignatureExporterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IsValidSignatureExporterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IsValidSignatureExporterSession struct {
	Contract     *IsValidSignatureExporter // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IsValidSignatureExporterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IsValidSignatureExporterCallerSession struct {
	Contract *IsValidSignatureExporterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// IsValidSignatureExporterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IsValidSignatureExporterTransactorSession struct {
	Contract     *IsValidSignatureExporterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// IsValidSignatureExporterRaw is an auto generated low-level Go binding around an Ethereum contract.
type IsValidSignatureExporterRaw struct {
	Contract *IsValidSignatureExporter // Generic contract binding to access the raw methods on
}

// IsValidSignatureExporterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IsValidSignatureExporterCallerRaw struct {
	Contract *IsValidSignatureExporterCaller // Generic read-only contract binding to access the raw methods on
}

// IsValidSignatureExporterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IsValidSignatureExporterTransactorRaw struct {
	Contract *IsValidSignatureExporterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIsValidSignatureExporter creates a new instance of IsValidSignatureExporter, bound to a specific deployed contract.
func NewIsValidSignatureExporter(address common.Address, backend bind.ContractBackend) (*IsValidSignatureExporter, error) {
	contract, err := bindIsValidSignatureExporter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IsValidSignatureExporter{IsValidSignatureExporterCaller: IsValidSignatureExporterCaller{contract: contract}, IsValidSignatureExporterTransactor: IsValidSignatureExporterTransactor{contract: contract}, IsValidSignatureExporterFilterer: IsValidSignatureExporterFilterer{contract: contract}}, nil
}

// NewIsValidSignatureExporterCaller creates a new read-only instance of IsValidSignatureExporter, bound to a specific deployed contract.
func NewIsValidSignatureExporterCaller(address common.Address, caller bind.ContractCaller) (*IsValidSignatureExporterCaller, error) {
	contract, err := bindIsValidSignatureExporter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IsValidSignatureExporterCaller{contract: contract}, nil
}

// NewIsValidSignatureExporterTransactor creates a new write-only instance of IsValidSignatureExporter, bound to a specific deployed contract.
func NewIsValidSignatureExporterTransactor(address common.Address, transactor bind.ContractTransactor) (*IsValidSignatureExporterTransactor, error) {
	contract, err := bindIsValidSignatureExporter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IsValidSignatureExporterTransactor{contract: contract}, nil
}

// NewIsValidSignatureExporterFilterer creates a new log filterer instance of IsValidSignatureExporter, bound to a specific deployed contract.
func NewIsValidSignatureExporterFilterer(address common.Address, filterer bind.ContractFilterer) (*IsValidSignatureExporterFilterer, error) {
	contract, err := bindIsValidSignatureExporter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IsValidSignatureExporterFilterer{contract: contract}, nil
}

// bindIsValidSignatureExporter binds a generic wrapper to an already deployed contract.
func bindIsValidSignatureExporter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IsValidSignatureExporterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IsValidSignatureExporter *IsValidSignatureExporterRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IsValidSignatureExporter.Contract.IsValidSignatureExporterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IsValidSignatureExporter *IsValidSignatureExporterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IsValidSignatureExporter.Contract.IsValidSignatureExporterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IsValidSignatureExporter *IsValidSignatureExporterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IsValidSignatureExporter.Contract.IsValidSignatureExporterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IsValidSignatureExporter *IsValidSignatureExporterCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IsValidSignatureExporter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IsValidSignatureExporter *IsValidSignatureExporterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IsValidSignatureExporter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IsValidSignatureExporter *IsValidSignatureExporterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IsValidSignatureExporter.Contract.contract.Transact(opts, method, params...)
}

// IsValidSignature is a free data retrieval call binding the contract method 0x20c13b0b.
//
// Solidity: function isValidSignature(bytes _data, bytes _signature) constant returns(bytes4)
func (_IsValidSignatureExporter *IsValidSignatureExporterCaller) IsValidSignature(opts *bind.CallOpts, _data []byte, _signature []byte) ([4]byte, error) {
	var (
		ret0 = new([4]byte)
	)
	out := ret0
	err := _IsValidSignatureExporter.contract.Call(opts, out, "isValidSignature", _data, _signature)
	return *ret0, err
}

// IsValidSignature is a free data retrieval call binding the contract method 0x20c13b0b.
//
// Solidity: function isValidSignature(bytes _data, bytes _signature) constant returns(bytes4)
func (_IsValidSignatureExporter *IsValidSignatureExporterSession) IsValidSignature(_data []byte, _signature []byte) ([4]byte, error) {
	return _IsValidSignatureExporter.Contract.IsValidSignature(&_IsValidSignatureExporter.CallOpts, _data, _signature)
}

// IsValidSignature is a free data retrieval call binding the contract method 0x20c13b0b.
//
// Solidity: function isValidSignature(bytes _data, bytes _signature) constant returns(bytes4)
func (_IsValidSignatureExporter *IsValidSignatureExporterCallerSession) IsValidSignature(_data []byte, _signature []byte) ([4]byte, error) {
	return _IsValidSignatureExporter.Contract.IsValidSignature(&_IsValidSignatureExporter.CallOpts, _data, _signature)
}
