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

// BytesUtilsExporterABI is the input ABI used to generate the binding from.
const BytesUtilsExporterABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_bts\",\"type\":\"bytes\"},{\"name\":\"_from\",\"type\":\"uint256\"}],\"name\":\"bytesToAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_bts\",\"type\":\"bytes\"},{\"name\":\"_from\",\"type\":\"uint256\"}],\"name\":\"bytesToUint256\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_bts\",\"type\":\"bytes\"},{\"name\":\"_from\",\"type\":\"uint256\"}],\"name\":\"bytesToBytes4\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// BytesUtilsExporterBin is the compiled bytecode used for deploying new contracts.
const BytesUtilsExporterBin = `608060405234801561001057600080fd5b50610522806100206000396000f3fe608060405234801561001057600080fd5b506004361061005d577c010000000000000000000000000000000000000000000000000000000060003504630638c0f4811461006257806355f94ce0146100fb578063cbd85c761461017d575b600080fd5b6100d26004803603604081101561007857600080fd5b81019060208101813564010000000081111561009357600080fd5b8201836020820111156100a557600080fd5b803590602001918460018302840111640100000000831117156100c757600080fd5b919350915035610222565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b61016b6004803603604081101561011157600080fd5b81019060208101813564010000000081111561012c57600080fd5b82018360208201111561013e57600080fd5b8035906020019184600183028401116401000000008311171561016057600080fd5b919350915035610275565b60408051918252519081900360200190f35b6101ed6004803603604081101561019357600080fd5b8101906020810181356401000000008111156101ae57600080fd5b8201836020820111156101c057600080fd5b803590602001918460018302840111640100000000831117156101e257600080fd5b9193509150356102c0565b604080517fffffffff000000000000000000000000000000000000000000000000000000009092168252519081900360200190f35b600061026d8285858080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929392505063ffffffff61030b169050565b949350505050565b600061026d8285858080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929392505063ffffffff6103bd169050565b600061026d8285858080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929392505063ffffffff61045f169050565b600061031e82601463ffffffff61047216565b8351101561038d57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f736c6963696e67206f7574206f662072616e6765000000000000000000000000604482015290519081900360640190fd5b6000806103a184602063ffffffff61047216565b909401516c010000000000000000000000009004949350505050565b60006103d082602063ffffffff61047216565b8351101561043f57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f736c6963696e67206f7574206f662072616e6765000000000000000000000000604482015290519081900360640190fd5b60008061045384602063ffffffff61047216565b90940151949350505050565b60006103d082600463ffffffff61047216565b6000828201838110156104e657604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b939250505056fea265627a7a723058200e523a7d098d489131ab93368229e267005f4151bccd2634759ecc0861cf5d8164736f6c634300050a0032`

// DeployBytesUtilsExporter deploys a new Ethereum contract, binding an instance of BytesUtilsExporter to it.
func DeployBytesUtilsExporter(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BytesUtilsExporter, error) {
	parsed, err := abi.JSON(strings.NewReader(BytesUtilsExporterABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BytesUtilsExporterBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BytesUtilsExporter{BytesUtilsExporterCaller: BytesUtilsExporterCaller{contract: contract}, BytesUtilsExporterTransactor: BytesUtilsExporterTransactor{contract: contract}, BytesUtilsExporterFilterer: BytesUtilsExporterFilterer{contract: contract}}, nil
}

// BytesUtilsExporter is an auto generated Go binding around an Ethereum contract.
type BytesUtilsExporter struct {
	BytesUtilsExporterCaller     // Read-only binding to the contract
	BytesUtilsExporterTransactor // Write-only binding to the contract
	BytesUtilsExporterFilterer   // Log filterer for contract events
}

// BytesUtilsExporterCaller is an auto generated read-only Go binding around an Ethereum contract.
type BytesUtilsExporterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BytesUtilsExporterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BytesUtilsExporterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BytesUtilsExporterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BytesUtilsExporterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BytesUtilsExporterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BytesUtilsExporterSession struct {
	Contract     *BytesUtilsExporter // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// BytesUtilsExporterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BytesUtilsExporterCallerSession struct {
	Contract *BytesUtilsExporterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// BytesUtilsExporterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BytesUtilsExporterTransactorSession struct {
	Contract     *BytesUtilsExporterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// BytesUtilsExporterRaw is an auto generated low-level Go binding around an Ethereum contract.
type BytesUtilsExporterRaw struct {
	Contract *BytesUtilsExporter // Generic contract binding to access the raw methods on
}

// BytesUtilsExporterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BytesUtilsExporterCallerRaw struct {
	Contract *BytesUtilsExporterCaller // Generic read-only contract binding to access the raw methods on
}

// BytesUtilsExporterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BytesUtilsExporterTransactorRaw struct {
	Contract *BytesUtilsExporterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBytesUtilsExporter creates a new instance of BytesUtilsExporter, bound to a specific deployed contract.
func NewBytesUtilsExporter(address common.Address, backend bind.ContractBackend) (*BytesUtilsExporter, error) {
	contract, err := bindBytesUtilsExporter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BytesUtilsExporter{BytesUtilsExporterCaller: BytesUtilsExporterCaller{contract: contract}, BytesUtilsExporterTransactor: BytesUtilsExporterTransactor{contract: contract}, BytesUtilsExporterFilterer: BytesUtilsExporterFilterer{contract: contract}}, nil
}

// NewBytesUtilsExporterCaller creates a new read-only instance of BytesUtilsExporter, bound to a specific deployed contract.
func NewBytesUtilsExporterCaller(address common.Address, caller bind.ContractCaller) (*BytesUtilsExporterCaller, error) {
	contract, err := bindBytesUtilsExporter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BytesUtilsExporterCaller{contract: contract}, nil
}

// NewBytesUtilsExporterTransactor creates a new write-only instance of BytesUtilsExporter, bound to a specific deployed contract.
func NewBytesUtilsExporterTransactor(address common.Address, transactor bind.ContractTransactor) (*BytesUtilsExporterTransactor, error) {
	contract, err := bindBytesUtilsExporter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BytesUtilsExporterTransactor{contract: contract}, nil
}

// NewBytesUtilsExporterFilterer creates a new log filterer instance of BytesUtilsExporter, bound to a specific deployed contract.
func NewBytesUtilsExporterFilterer(address common.Address, filterer bind.ContractFilterer) (*BytesUtilsExporterFilterer, error) {
	contract, err := bindBytesUtilsExporter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BytesUtilsExporterFilterer{contract: contract}, nil
}

// bindBytesUtilsExporter binds a generic wrapper to an already deployed contract.
func bindBytesUtilsExporter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BytesUtilsExporterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BytesUtilsExporter *BytesUtilsExporterRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BytesUtilsExporter.Contract.BytesUtilsExporterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BytesUtilsExporter *BytesUtilsExporterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BytesUtilsExporter.Contract.BytesUtilsExporterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BytesUtilsExporter *BytesUtilsExporterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BytesUtilsExporter.Contract.BytesUtilsExporterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BytesUtilsExporter *BytesUtilsExporterCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BytesUtilsExporter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BytesUtilsExporter *BytesUtilsExporterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BytesUtilsExporter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BytesUtilsExporter *BytesUtilsExporterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BytesUtilsExporter.Contract.contract.Transact(opts, method, params...)
}

// BytesToAddress is a free data retrieval call binding the contract method 0x0638c0f4.
//
// Solidity: function bytesToAddress(bytes _bts, uint256 _from) constant returns(address)
func (_BytesUtilsExporter *BytesUtilsExporterCaller) BytesToAddress(opts *bind.CallOpts, _bts []byte, _from *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BytesUtilsExporter.contract.Call(opts, out, "bytesToAddress", _bts, _from)
	return *ret0, err
}

// BytesToAddress is a free data retrieval call binding the contract method 0x0638c0f4.
//
// Solidity: function bytesToAddress(bytes _bts, uint256 _from) constant returns(address)
func (_BytesUtilsExporter *BytesUtilsExporterSession) BytesToAddress(_bts []byte, _from *big.Int) (common.Address, error) {
	return _BytesUtilsExporter.Contract.BytesToAddress(&_BytesUtilsExporter.CallOpts, _bts, _from)
}

// BytesToAddress is a free data retrieval call binding the contract method 0x0638c0f4.
//
// Solidity: function bytesToAddress(bytes _bts, uint256 _from) constant returns(address)
func (_BytesUtilsExporter *BytesUtilsExporterCallerSession) BytesToAddress(_bts []byte, _from *big.Int) (common.Address, error) {
	return _BytesUtilsExporter.Contract.BytesToAddress(&_BytesUtilsExporter.CallOpts, _bts, _from)
}

// BytesToBytes4 is a free data retrieval call binding the contract method 0xcbd85c76.
//
// Solidity: function bytesToBytes4(bytes _bts, uint256 _from) constant returns(bytes4)
func (_BytesUtilsExporter *BytesUtilsExporterCaller) BytesToBytes4(opts *bind.CallOpts, _bts []byte, _from *big.Int) ([4]byte, error) {
	var (
		ret0 = new([4]byte)
	)
	out := ret0
	err := _BytesUtilsExporter.contract.Call(opts, out, "bytesToBytes4", _bts, _from)
	return *ret0, err
}

// BytesToBytes4 is a free data retrieval call binding the contract method 0xcbd85c76.
//
// Solidity: function bytesToBytes4(bytes _bts, uint256 _from) constant returns(bytes4)
func (_BytesUtilsExporter *BytesUtilsExporterSession) BytesToBytes4(_bts []byte, _from *big.Int) ([4]byte, error) {
	return _BytesUtilsExporter.Contract.BytesToBytes4(&_BytesUtilsExporter.CallOpts, _bts, _from)
}

// BytesToBytes4 is a free data retrieval call binding the contract method 0xcbd85c76.
//
// Solidity: function bytesToBytes4(bytes _bts, uint256 _from) constant returns(bytes4)
func (_BytesUtilsExporter *BytesUtilsExporterCallerSession) BytesToBytes4(_bts []byte, _from *big.Int) ([4]byte, error) {
	return _BytesUtilsExporter.Contract.BytesToBytes4(&_BytesUtilsExporter.CallOpts, _bts, _from)
}

// BytesToUint256 is a free data retrieval call binding the contract method 0x55f94ce0.
//
// Solidity: function bytesToUint256(bytes _bts, uint256 _from) constant returns(uint256)
func (_BytesUtilsExporter *BytesUtilsExporterCaller) BytesToUint256(opts *bind.CallOpts, _bts []byte, _from *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BytesUtilsExporter.contract.Call(opts, out, "bytesToUint256", _bts, _from)
	return *ret0, err
}

// BytesToUint256 is a free data retrieval call binding the contract method 0x55f94ce0.
//
// Solidity: function bytesToUint256(bytes _bts, uint256 _from) constant returns(uint256)
func (_BytesUtilsExporter *BytesUtilsExporterSession) BytesToUint256(_bts []byte, _from *big.Int) (*big.Int, error) {
	return _BytesUtilsExporter.Contract.BytesToUint256(&_BytesUtilsExporter.CallOpts, _bts, _from)
}

// BytesToUint256 is a free data retrieval call binding the contract method 0x55f94ce0.
//
// Solidity: function bytesToUint256(bytes _bts, uint256 _from) constant returns(uint256)
func (_BytesUtilsExporter *BytesUtilsExporterCallerSession) BytesToUint256(_bts []byte, _from *big.Int) (*big.Int, error) {
	return _BytesUtilsExporter.Contract.BytesToUint256(&_BytesUtilsExporter.CallOpts, _bts, _from)
}
