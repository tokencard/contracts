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
const ParseIntScientificExporterABI = "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_a\",\"type\":\"string\"}],\"name\":\"parseIntScientific\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_a\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_b\",\"type\":\"uint256\"}],\"name\":\"parseIntScientificDecimals\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_a\",\"type\":\"string\"}],\"name\":\"parseIntScientificWei\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// ParseIntScientificExporterBin is the compiled bytecode used for deploying new contracts.
var ParseIntScientificExporterBin = "0x608060405234801561001057600080fd5b50610c71806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c806361f7e0ac1461004657806387c8da5e146100c8578063ba07069514610138575b600080fd5b6100b66004803603602081101561005c57600080fd5b81019060208101813564010000000081111561007757600080fd5b82018360208201111561008957600080fd5b803590602001918460018302840111640100000000831117156100ab57600080fd5b5090925090506101a8565b60408051918252519081900360200190f35b6100b6600480360360408110156100de57600080fd5b8101906020810181356401000000008111156100f957600080fd5b82018360208201111561010b57600080fd5b8035906020019184600183028401116401000000008311171561012d57600080fd5b9193509150356101f2565b6100b66004803603602081101561014e57600080fd5b81019060208101813564010000000081111561016957600080fd5b82018360208201111561017b57600080fd5b8035906020019184600183028401116401000000008311171561019d57600080fd5b50909250905061023d565b60006101e983838080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061027e92505050565b90505b92915050565b600061023584848080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250869250610287915050565b949350505050565b60006101e983838080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250610b3892505050565b60006101ec8260125b60008281808080808080808080805b8b518110156108c9578b51600360fc1b908d90839081106102b357fe5b01602001516001600160f81b031916108015906102f157508b51603960f81b908d90839081106102df57fe5b01602001516001600160f81b03191611155b80156102fb575083155b156103ab578415610359576103178a600a63ffffffff610b4516565b995061034c603060f81b60f81c8d838151811061033057fe5b01602001518c9160f89190911c0360ff1663ffffffff610b9e16565b99506001909701966103a6565b6001955061036e8b600a63ffffffff610b4516565b9a506103a3603060f81b60f81c8d838151811061038757fe5b01602001518d9160f89190911c0360ff1663ffffffff610b9e16565b9a505b6108c1565b8b51600360fc1b908d90839081106103bf57fe5b01602001516001600160f81b031916108015906103fd57508b51603960f81b908d90839081106103eb57fe5b01602001516001600160f81b03191611155b80156104065750835b156104585761041c89600a63ffffffff610b4516565b9850610451603060f81b60f81c8d838151811061043557fe5b01602001518b9160f89190911c0360ff1663ffffffff610b9e16565b98506108c1565b8b51601760f91b908d908390811061046c57fe5b01602001516001600160f81b031916141561057357856104cb576040805162461bcd60e51b81526020600482015260156024820152741b5a5cdcda5b99c81a5b9d1959dc985b081c185c9d605a1b604482015290519081900360640190fd5b841561051e576040805162461bcd60e51b815260206004820152601760248201527f6475706c696361746520646563696d616c20706f696e74000000000000000000604482015290519081900360640190fd5b831561056a576040805162461bcd60e51b8152602060048201526016602482015275191958da5b585b0818599d195c88195e1c1bdb995b9d60521b604482015290519081900360640190fd5b600194506108c1565b8b51602d60f81b908d908390811061058757fe5b01602001516001600160f81b031916141561067d5782156105dd576040805162461bcd60e51b815260206004820152600b60248201526a6475706c6963617465202d60a81b604482015290519081900360640190fd5b811561061d576040805162461bcd60e51b815260206004820152600a60248201526932bc3a39309039b4b3b760b11b604482015290519081900360640190fd5b808760010114610674576040805162461bcd60e51b815260206004820152601e60248201527f2d207369676e206e6f7420696d6d6564696174656c7920616674657220650000604482015290519081900360640190fd5b600192506108c1565b8b51602b60f81b908d908390811061069157fe5b01602001516001600160f81b03191614156107875781156106e7576040805162461bcd60e51b815260206004820152600b60248201526a6475706c6963617465202b60a81b604482015290519081900360640190fd5b8215610727576040805162461bcd60e51b815260206004820152600a60248201526932bc3a39309039b4b3b760b11b604482015290519081900360640190fd5b80876001011461077e576040805162461bcd60e51b815260206004820152601e60248201527f2b207369676e206e6f7420696d6d6564696174656c7920616674657220650000604482015290519081900360640190fd5b600191506108c1565b8b51604560f81b908d908390811061079b57fe5b01602001516001600160f81b03191614806107d657508b51606560f81b908d90839081106107c557fe5b01602001516001600160f81b031916145b156108845785610825576040805162461bcd60e51b81526020600482015260156024820152741b5a5cdcda5b99c81a5b9d1959dc985b081c185c9d605a1b604482015290519081900360640190fd5b8315610878576040805162461bcd60e51b815260206004820152601960248201527f6475706c6963617465206578706f6e656e742073796d626f6c00000000000000604482015290519081900360640190fd5b600193508096506108c1565b6040805162461bcd60e51b815260206004820152600d60248201526c1a5b9d985b1a5908191a59da5d609a1b604482015290519081900360640190fd5b600101610296565b82806108d25750815b156108eb578660020181116108e657600080fd5b610900565b83156109005786600101811161090057600080fd5b8215610981578d891061097757604e8e8a0310610954576040805162461bcd60e51b815260206004820152600d60248201526c6578706f6e656e74203e20373760981b604482015290519081900360640190fd5b8d8903600a0a8b8161096257fe5b049c506101ec9b505050505050505050505050565b888e039d50610994565b6109918e8a63ffffffff610b9e16565b9d505b878e10610a6857604e88106109da5760405162461bcd60e51b8152600401808060200182810382526022815260200180610bf96022913960400191505060405180910390fd5b6109ee8b600a8a900a63ffffffff610b4516565b9a50610a008b8b63ffffffff610b9e16565b9a50604e888f0310610a49576040805162461bcd60e51b815260206004820152600d60248201526c6578706f6e656e74203e20373760981b604482015290519081900360640190fd5b610a61888f03600a0a8c610b4590919063ffffffff16565b9a50610b25565b8d88039750604e8810610aac5760405162461bcd60e51b8152600401808060200182810382526022815260200180610bf96022913960400191505060405180910390fd5b87600a0a8a81610ab857fe5b049950604e8e10610afa5760405162461bcd60e51b8152600401808060200182810382526022815260200180610bf96022913960400191505060405180910390fd5b610b108e600a0a8c610b4590919063ffffffff16565b9a50610b228b8b63ffffffff610b9e16565b9a505b50989d9c50505050505050505050505050565b60006101ec826000610287565b600082610b54575060006101ec565b82820282848281610b6157fe5b04146101e95760405162461bcd60e51b8152600401808060200182810382526021815260200180610c1b6021913960400191505060405180910390fd5b6000828201838110156101e9576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fdfe6d6f7265207468616e20373720646563696d616c2064696769747320706172736564536166654d6174683a206d756c7469706c69636174696f6e206f766572666c6f77a26469706673582212202a0468f823d4b0419383c11103d614b644150473bd9584c6a07ee419c535012e64736f6c634300060b0033"

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
