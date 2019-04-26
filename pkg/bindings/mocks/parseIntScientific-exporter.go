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

// ParseIntScientificExporterABI is the input ABI used to generate the binding from.
const ParseIntScientificExporterABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_a\",\"type\":\"string\"}],\"name\":\"parseIntScientificWei\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_a\",\"type\":\"string\"},{\"name\":\"_b\",\"type\":\"uint256\"}],\"name\":\"parseIntScientificDecimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_a\",\"type\":\"string\"}],\"name\":\"parseIntScientific\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// ParseIntScientificExporterBin is the compiled bytecode used for deploying new contracts.
const ParseIntScientificExporterBin = `608060405234801561001057600080fd5b50610e79806100206000396000f3006080604052600436106100565763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166361f7e0ac811461005b57806387c8da5e1461008d578063ba070695146100b1575b600080fd5b34801561006757600080fd5b5061007b60048035602481019101356100d1565b60408051918252519081900360200190f35b34801561009957600080fd5b5061007b602460048035828101929101359035610113565b3480156100bd57600080fd5b5061007b6004803560248101910135610158565b600061010c83838080601f01602080910402602001604051908101604052809392919081815260200183838082843750610193945050505050565b9392505050565b600061015084848080601f016020809104026020016040519081016040528093929190818152602001838380828437508894506101a69350505050565b949350505050565b600061010c83838080601f01602080910402602001604051908101604052809392919081815260200183838082843750610df5945050505050565b60006101a08260126101a6565b92915050565b60008281808080808080808080805b8b51811015610a22578b517f3000000000000000000000000000000000000000000000000000000000000000908d90839081106101ee57fe5b90602001015160f860020a900460f860020a02600160f860020a0319161015801561026457508b517f3900000000000000000000000000000000000000000000000000000000000000908d908390811061024457fe5b90602001015160f860020a900460f860020a02600160f860020a03191611155b801561026e575083155b156103385784156102d95761028a8a600a63ffffffff610e0216565b8c51909a506102cc906030908e90849081106102a257fe5b90602001015160f860020a900460f860020a0260f860020a9004038b610e3b90919063ffffffff16565b9950600190970196610333565b600195506102ee8b600a63ffffffff610e0216565b8c51909b50610330906030908e908490811061030657fe5b90602001015160f860020a900460f860020a0260f860020a9004038c610e3b90919063ffffffff16565b9a505b610a1a565b8b517f3000000000000000000000000000000000000000000000000000000000000000908d908390811061036857fe5b90602001015160f860020a900460f860020a02600160f860020a031916101580156103de57508b517f3900000000000000000000000000000000000000000000000000000000000000908d90839081106103be57fe5b90602001015160f860020a900460f860020a02600160f860020a03191611155b80156103e75750835b15610446576103fd89600a63ffffffff610e0216565b8c5190995061043f906030908e908490811061041557fe5b90602001015160f860020a900460f860020a0260f860020a9004038a610e3b90919063ffffffff16565b9850610a1a565b8b517f2e00000000000000000000000000000000000000000000000000000000000000908d908390811061047657fe5b90602001015160f860020a900460f860020a02600160f860020a03191614156105a5578515156104f0576040805160e560020a62461bcd02815260206004820152601560248201527f6d697373696e6720696e74656772616c20706172740000000000000000000000604482015290519081900360640190fd5b8415610546576040805160e560020a62461bcd02815260206004820152601760248201527f6475706c696361746520646563696d616c20706f696e74000000000000000000604482015290519081900360640190fd5b831561059c576040805160e560020a62461bcd02815260206004820152601660248201527f646563696d616c206166746572206578706f6e656e7400000000000000000000604482015290519081900360640190fd5b60019450610a1a565b8b517f2d00000000000000000000000000000000000000000000000000000000000000908d90839081106105d557fe5b90602001015160f860020a900460f860020a02600160f860020a031916141561070757821561064e576040805160e560020a62461bcd02815260206004820152600b60248201527f6475706c6963617465202d000000000000000000000000000000000000000000604482015290519081900360640190fd5b81156106a4576040805160e560020a62461bcd02815260206004820152600a60248201527f6578747261207369676e00000000000000000000000000000000000000000000604482015290519081900360640190fd5b6001870181146106fe576040805160e560020a62461bcd02815260206004820152601e60248201527f2d207369676e206e6f7420696d6d6564696174656c7920616674657220650000604482015290519081900360640190fd5b60019250610a1a565b8b517f2b00000000000000000000000000000000000000000000000000000000000000908d908390811061073757fe5b90602001015160f860020a900460f860020a02600160f860020a03191614156108695781156107b0576040805160e560020a62461bcd02815260206004820152600b60248201527f6475706c6963617465202b000000000000000000000000000000000000000000604482015290519081900360640190fd5b8215610806576040805160e560020a62461bcd02815260206004820152600a60248201527f6578747261207369676e00000000000000000000000000000000000000000000604482015290519081900360640190fd5b600187018114610860576040805160e560020a62461bcd02815260206004820152601e60248201527f2b207369676e206e6f7420696d6d6564696174656c7920616674657220650000604482015290519081900360640190fd5b60019150610a1a565b8b517f4500000000000000000000000000000000000000000000000000000000000000908d908390811061089957fe5b90602001015160f860020a900460f860020a02600160f860020a031916148061090c57508b517f6500000000000000000000000000000000000000000000000000000000000000908d90839081106108ed57fe5b90602001015160f860020a900460f860020a02600160f860020a031916145b156109ca57851515610968576040805160e560020a62461bcd02815260206004820152601560248201527f6d697373696e6720696e74656772616c20706172740000000000000000000000604482015290519081900360640190fd5b83156109be576040805160e560020a62461bcd02815260206004820152601960248201527f6475706c6963617465206578706f6e656e742073796d626f6c00000000000000604482015290519081900360640190fd5b60019350809650610a1a565b6040805160e560020a62461bcd02815260206004820152600d60248201527f696e76616c696420646967697400000000000000000000000000000000000000604482015290519081900360640190fd5b6001016101b5565b8280610a2b5750815b15610a4457600287018111610a3f57600080fd5b610a59565b8315610a5957600187018111610a5957600080fd5b8215610ae5578d8910610adb57604e8e8a0310610ac0576040805160e560020a62461bcd02815260206004820152600d60248201527f6578706f6e656e74203e20373700000000000000000000000000000000000000604482015290519081900360640190fd5b8d8903600a0a8b811515610ad057fe5b049a508a9c50610de3565b888e039d50610af8565b610af58e8a63ffffffff610e3b16565b9d505b878e10610c1e57604e8810610b7d576040805160e560020a62461bcd02815260206004820152602260248201527f6d6f7265207468616e20373720646563696d616c20646967697473207061727360448201527f6564000000000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b610b918b600a8a900a63ffffffff610e0216565b9a50610ba38b8b63ffffffff610e3b16565b9a50604e888f0310610bff576040805160e560020a62461bcd02815260206004820152600d60248201527f6578706f6e656e74203e20373700000000000000000000000000000000000000604482015290519081900360640190fd5b610c17888f03600a0a8c610e0290919063ffffffff16565b9a50610d5b565b968d900396604e8810610ca1576040805160e560020a62461bcd02815260206004820152602260248201527f6d6f7265207468616e20373720646563696d616c20646967697473207061727360448201527f6564000000000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b87600a0a8a811515610caf57fe5b049950604e8e10610d30576040805160e560020a62461bcd02815260206004820152602260248201527f6d6f7265207468616e20373720646563696d616c20646967697473207061727360448201527f6564000000000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b610d468e600a0a8c610e0290919063ffffffff16565b9a50610d588b8b63ffffffff610e3b16565b9a505b66400000000000008b10610ddf576040805160e560020a62461bcd02815260206004820152603c60248201527f6e756d626572206578636565646564206d6178696d756d20616c6c6f7765642060448201527f76616c756520666f722073616665206a736f6e206465636f64696e6700000000606482015290519081900360840190fd5b8a9c505b50505050505050505050505092915050565b60006101a08260006101a6565b600080831515610e155760009150610e34565b50828202828482811515610e2557fe5b0414610e3057600080fd5b8091505b5092915050565b600082820183811015610e3057600080fd00a165627a7a72305820dfdbb9969780d95a2559759f75357f98d5cfe31afa8b53596e84ccff971f5ab50029`

// DeployParseIntScientificExporter deploys a new Ethereum contract, binding an instance of ParseIntScientificExporter to it.
func DeployParseIntScientificExporter(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ParseIntScientificExporter, error) {
	parsed, err := abi.JSON(strings.NewReader(ParseIntScientificExporterABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ParseIntScientificExporterBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ParseIntScientificExporter{ParseIntScientificExporterCaller: ParseIntScientificExporterCaller{contract: contract}, ParseIntScientificExporterTransactor: ParseIntScientificExporterTransactor{contract: contract}, ParseIntScientificExporterFilterer: ParseIntScientificExporterFilterer{contract: contract}}, nil
}

// ParseIntScientificExporter is an auto generated Go binding around an Ethereum contract.
type ParseIntScientificExporter struct {
	ParseIntScientificExporterCaller     // Read-only binding to the contract
	ParseIntScientificExporterTransactor // Write-only binding to the contract
	ParseIntScientificExporterFilterer   // Log filterer for contract events
}

// ParseIntScientificExporterCaller is an auto generated read-only Go binding around an Ethereum contract.
type ParseIntScientificExporterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ParseIntScientificExporterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ParseIntScientificExporterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ParseIntScientificExporterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ParseIntScientificExporterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ParseIntScientificExporterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ParseIntScientificExporterSession struct {
	Contract     *ParseIntScientificExporter // Generic contract binding to set the session for
	CallOpts     bind.CallOpts               // Call options to use throughout this session
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ParseIntScientificExporterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ParseIntScientificExporterCallerSession struct {
	Contract *ParseIntScientificExporterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                     // Call options to use throughout this session
}

// ParseIntScientificExporterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ParseIntScientificExporterTransactorSession struct {
	Contract     *ParseIntScientificExporterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                     // Transaction auth options to use throughout this session
}

// ParseIntScientificExporterRaw is an auto generated low-level Go binding around an Ethereum contract.
type ParseIntScientificExporterRaw struct {
	Contract *ParseIntScientificExporter // Generic contract binding to access the raw methods on
}

// ParseIntScientificExporterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ParseIntScientificExporterCallerRaw struct {
	Contract *ParseIntScientificExporterCaller // Generic read-only contract binding to access the raw methods on
}

// ParseIntScientificExporterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ParseIntScientificExporterTransactorRaw struct {
	Contract *ParseIntScientificExporterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewParseIntScientificExporter creates a new instance of ParseIntScientificExporter, bound to a specific deployed contract.
func NewParseIntScientificExporter(address common.Address, backend bind.ContractBackend) (*ParseIntScientificExporter, error) {
	contract, err := bindParseIntScientificExporter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ParseIntScientificExporter{ParseIntScientificExporterCaller: ParseIntScientificExporterCaller{contract: contract}, ParseIntScientificExporterTransactor: ParseIntScientificExporterTransactor{contract: contract}, ParseIntScientificExporterFilterer: ParseIntScientificExporterFilterer{contract: contract}}, nil
}

// NewParseIntScientificExporterCaller creates a new read-only instance of ParseIntScientificExporter, bound to a specific deployed contract.
func NewParseIntScientificExporterCaller(address common.Address, caller bind.ContractCaller) (*ParseIntScientificExporterCaller, error) {
	contract, err := bindParseIntScientificExporter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ParseIntScientificExporterCaller{contract: contract}, nil
}

// NewParseIntScientificExporterTransactor creates a new write-only instance of ParseIntScientificExporter, bound to a specific deployed contract.
func NewParseIntScientificExporterTransactor(address common.Address, transactor bind.ContractTransactor) (*ParseIntScientificExporterTransactor, error) {
	contract, err := bindParseIntScientificExporter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ParseIntScientificExporterTransactor{contract: contract}, nil
}

// NewParseIntScientificExporterFilterer creates a new log filterer instance of ParseIntScientificExporter, bound to a specific deployed contract.
func NewParseIntScientificExporterFilterer(address common.Address, filterer bind.ContractFilterer) (*ParseIntScientificExporterFilterer, error) {
	contract, err := bindParseIntScientificExporter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ParseIntScientificExporterFilterer{contract: contract}, nil
}

// bindParseIntScientificExporter binds a generic wrapper to an already deployed contract.
func bindParseIntScientificExporter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ParseIntScientificExporterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ParseIntScientificExporter *ParseIntScientificExporterRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ParseIntScientificExporter.Contract.ParseIntScientificExporterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ParseIntScientificExporter *ParseIntScientificExporterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ParseIntScientificExporter.Contract.ParseIntScientificExporterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ParseIntScientificExporter *ParseIntScientificExporterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ParseIntScientificExporter.Contract.ParseIntScientificExporterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ParseIntScientificExporter *ParseIntScientificExporterCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ParseIntScientificExporter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ParseIntScientificExporter *ParseIntScientificExporterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ParseIntScientificExporter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ParseIntScientificExporter *ParseIntScientificExporterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ParseIntScientificExporter.Contract.contract.Transact(opts, method, params...)
}

// ParseIntScientific is a free data retrieval call binding the contract method 0xba070695.
//
// Solidity: function parseIntScientific(string _a) constant returns(uint256)
func (_ParseIntScientificExporter *ParseIntScientificExporterCaller) ParseIntScientific(opts *bind.CallOpts, _a string) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ParseIntScientificExporter.contract.Call(opts, out, "parseIntScientific", _a)
	return *ret0, err
}

// ParseIntScientific is a free data retrieval call binding the contract method 0xba070695.
//
// Solidity: function parseIntScientific(string _a) constant returns(uint256)
func (_ParseIntScientificExporter *ParseIntScientificExporterSession) ParseIntScientific(_a string) (*big.Int, error) {
	return _ParseIntScientificExporter.Contract.ParseIntScientific(&_ParseIntScientificExporter.CallOpts, _a)
}

// ParseIntScientific is a free data retrieval call binding the contract method 0xba070695.
//
// Solidity: function parseIntScientific(string _a) constant returns(uint256)
func (_ParseIntScientificExporter *ParseIntScientificExporterCallerSession) ParseIntScientific(_a string) (*big.Int, error) {
	return _ParseIntScientificExporter.Contract.ParseIntScientific(&_ParseIntScientificExporter.CallOpts, _a)
}

// ParseIntScientificDecimals is a free data retrieval call binding the contract method 0x87c8da5e.
//
// Solidity: function parseIntScientificDecimals(string _a, uint256 _b) constant returns(uint256)
func (_ParseIntScientificExporter *ParseIntScientificExporterCaller) ParseIntScientificDecimals(opts *bind.CallOpts, _a string, _b *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ParseIntScientificExporter.contract.Call(opts, out, "parseIntScientificDecimals", _a, _b)
	return *ret0, err
}

// ParseIntScientificDecimals is a free data retrieval call binding the contract method 0x87c8da5e.
//
// Solidity: function parseIntScientificDecimals(string _a, uint256 _b) constant returns(uint256)
func (_ParseIntScientificExporter *ParseIntScientificExporterSession) ParseIntScientificDecimals(_a string, _b *big.Int) (*big.Int, error) {
	return _ParseIntScientificExporter.Contract.ParseIntScientificDecimals(&_ParseIntScientificExporter.CallOpts, _a, _b)
}

// ParseIntScientificDecimals is a free data retrieval call binding the contract method 0x87c8da5e.
//
// Solidity: function parseIntScientificDecimals(string _a, uint256 _b) constant returns(uint256)
func (_ParseIntScientificExporter *ParseIntScientificExporterCallerSession) ParseIntScientificDecimals(_a string, _b *big.Int) (*big.Int, error) {
	return _ParseIntScientificExporter.Contract.ParseIntScientificDecimals(&_ParseIntScientificExporter.CallOpts, _a, _b)
}

// ParseIntScientificWei is a free data retrieval call binding the contract method 0x61f7e0ac.
//
// Solidity: function parseIntScientificWei(string _a) constant returns(uint256)
func (_ParseIntScientificExporter *ParseIntScientificExporterCaller) ParseIntScientificWei(opts *bind.CallOpts, _a string) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ParseIntScientificExporter.contract.Call(opts, out, "parseIntScientificWei", _a)
	return *ret0, err
}

// ParseIntScientificWei is a free data retrieval call binding the contract method 0x61f7e0ac.
//
// Solidity: function parseIntScientificWei(string _a) constant returns(uint256)
func (_ParseIntScientificExporter *ParseIntScientificExporterSession) ParseIntScientificWei(_a string) (*big.Int, error) {
	return _ParseIntScientificExporter.Contract.ParseIntScientificWei(&_ParseIntScientificExporter.CallOpts, _a)
}

// ParseIntScientificWei is a free data retrieval call binding the contract method 0x61f7e0ac.
//
// Solidity: function parseIntScientificWei(string _a) constant returns(uint256)
func (_ParseIntScientificExporter *ParseIntScientificExporterCallerSession) ParseIntScientificWei(_a string) (*big.Int, error) {
	return _ParseIntScientificExporter.Contract.ParseIntScientificWei(&_ParseIntScientificExporter.CallOpts, _a)
}
