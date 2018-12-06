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
const ParseIntScientificExporterABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_a\",\"type\":\"string\"},{\"name\":\"_b\",\"type\":\"uint256\"}],\"name\":\"parseIntScientificDecimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_a\",\"type\":\"string\"}],\"name\":\"parseIntScientific\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// ParseIntScientificExporterBin is the compiled bytecode used for deploying new contracts.
const ParseIntScientificExporterBin = `608060405234801561001057600080fd5b5061079f806100206000396000f30060806040526004361061004b5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166387c8da5e8114610050578063ba070695146100bd575b600080fd5b34801561005c57600080fd5b506040805160206004803580820135601f81018490048402850184019095528484526100ab94369492936024939284019190819084018382808284375094975050933594506101169350505050565b60408051918252519081900360200190f35b3480156100c957600080fd5b506040805160206004803580820135601f81018490048402850184019095528484526100ab9436949293602493928401919081908401838280828437509497506101299650505050505050565b6000610122838361013c565b9392505050565b600061013682600061013c565b92915050565b60008281808080808080808080805b8b5181101561058c578b517f3000000000000000000000000000000000000000000000000000000000000000908d908390811061018457fe5b90602001015160f860020a900460f860020a02600160f860020a031916101580156101fa57508b517f3900000000000000000000000000000000000000000000000000000000000000908d90839081106101da57fe5b90602001015160f860020a900460f860020a02600160f860020a03191611155b8015610204575083155b15610286578415610249578b51600a9a909a02996030908d908390811061022757fe5b016020015160f860020a90819004810204039990990198600190970196610281565b8b51600a9b909b029a6030908d908390811061026157fe5b90602001015160f860020a900460f860020a0260f860020a9004038b019a505b610584565b8b517f3000000000000000000000000000000000000000000000000000000000000000908d90839081106102b657fe5b90602001015160f860020a900460f860020a02600160f860020a0319161015801561032c57508b517f3900000000000000000000000000000000000000000000000000000000000000908d908390811061030c57fe5b90602001015160f860020a900460f860020a02600160f860020a03191611155b801561033c5750828061033c5750815b1561037d578b51600a99909902986030908d908390811061035957fe5b90602001015160f860020a900460f860020a0260f860020a90040389019850610584565b8b517f2e00000000000000000000000000000000000000000000000000000000000000908d90839081106103ad57fe5b90602001015160f860020a900460f860020a02600160f860020a0319161480156103d5575084155b156103e35760019450610584565b8b517f2d00000000000000000000000000000000000000000000000000000000000000908d908390811061041357fe5b90602001015160f860020a900460f860020a02600160f860020a03191614801561043b575082155b801561044957508087600101145b156104575760019250610584565b8b517f2b00000000000000000000000000000000000000000000000000000000000000908d908390811061048757fe5b90602001015160f860020a900460f860020a02600160f860020a0319161480156104af575081155b80156104bd57508087600101145b156104cb5760019150610584565b8b517f6500000000000000000000000000000000000000000000000000000000000000908d90839081106104fb57fe5b90602001015160f860020a900460f860020a02600160f860020a031916148015610523575083155b156105345760019350809650610584565b6040805160e560020a62461bcd02815260206004820152600c60248201527f77726f6e6720666f726d61740000000000000000000000000000000000000000604482015290519081900360640190fd5b60010161014b565b838015610597575082155b80156105a1575081155b156105f6576040805160e560020a62461bcd02815260206004820152601d60248201527f6578706f6e656e74206e6f7420666f6c6c6f776564206279207369676e000000604482015290519081900360640190fd5b8380156106005750825b801561060e57508660020181145b15610663576040805160e560020a62461bcd02815260206004820152601660248201527f6578706f6e656e74206e6f742073706563696669656400000000000000000000604482015290519081900360640190fd5b83801561066d5750815b801561067b57508660020181145b156106d0576040805160e560020a62461bcd02815260206004820152601660248201527f6578706f6e656e74206e6f742073706563696669656400000000000000000000604482015290519081900360640190fd5b8215610743578d89106106f5578d8903600a0a8b8115156106ed57fe5b049a5061073e565b888e03955087861061071d5787600a0a8b029a50898b019a50878603600a0a8b029a5061073e565b858803600a0a8a81151561072d57fe5b04995085600a0a8b029a50898b019a505b610760565b888e01955087600a0a8b029a50898b019a50878603600a0a8b029a505b50989d9c505050505050505050505050505600a165627a7a723058204d4e01865ec3836be216476be639796e65c079bd79adebc4baaed9c97a33d7600029`

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
// Solidity: function parseIntScientific(_a string) constant returns(uint256)
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
// Solidity: function parseIntScientific(_a string) constant returns(uint256)
func (_ParseIntScientificExporter *ParseIntScientificExporterSession) ParseIntScientific(_a string) (*big.Int, error) {
	return _ParseIntScientificExporter.Contract.ParseIntScientific(&_ParseIntScientificExporter.CallOpts, _a)
}

// ParseIntScientific is a free data retrieval call binding the contract method 0xba070695.
//
// Solidity: function parseIntScientific(_a string) constant returns(uint256)
func (_ParseIntScientificExporter *ParseIntScientificExporterCallerSession) ParseIntScientific(_a string) (*big.Int, error) {
	return _ParseIntScientificExporter.Contract.ParseIntScientific(&_ParseIntScientificExporter.CallOpts, _a)
}

// ParseIntScientificDecimals is a free data retrieval call binding the contract method 0x87c8da5e.
//
// Solidity: function parseIntScientificDecimals(_a string, _b uint256) constant returns(uint256)
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
// Solidity: function parseIntScientificDecimals(_a string, _b uint256) constant returns(uint256)
func (_ParseIntScientificExporter *ParseIntScientificExporterSession) ParseIntScientificDecimals(_a string, _b *big.Int) (*big.Int, error) {
	return _ParseIntScientificExporter.Contract.ParseIntScientificDecimals(&_ParseIntScientificExporter.CallOpts, _a, _b)
}

// ParseIntScientificDecimals is a free data retrieval call binding the contract method 0x87c8da5e.
//
// Solidity: function parseIntScientificDecimals(_a string, _b uint256) constant returns(uint256)
func (_ParseIntScientificExporter *ParseIntScientificExporterCallerSession) ParseIntScientificDecimals(_a string, _b *big.Int) (*big.Int, error) {
	return _ParseIntScientificExporter.Contract.ParseIntScientificDecimals(&_ParseIntScientificExporter.CallOpts, _a, _b)
}
