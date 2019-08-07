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
const ParseIntScientificExporterBin = `608060405234801561001057600080fd5b50610f90806100206000396000f3fe608060405234801561001057600080fd5b506004361061005d577c0100000000000000000000000000000000000000000000000000000000600035046361f7e0ac811461006257806387c8da5e146100e4578063ba07069514610154575b600080fd5b6100d26004803603602081101561007857600080fd5b81019060208101813564010000000081111561009357600080fd5b8201836020820111156100a557600080fd5b803590602001918460018302840111640100000000831117156100c757600080fd5b5090925090506101c4565b60408051918252519081900360200190f35b6100d2600480360360408110156100fa57600080fd5b81019060208101813564010000000081111561011557600080fd5b82018360208201111561012757600080fd5b8035906020019184600183028401116401000000008311171561014957600080fd5b91935091503561020e565b6100d26004803603602081101561016a57600080fd5b81019060208101813564010000000081111561018557600080fd5b82018360208201111561019757600080fd5b803590602001918460018302840111640100000000831117156101b957600080fd5b509092509050610259565b600061020583838080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061029a92505050565b90505b92915050565b600061025184848080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152508692506102a3915050565b949350505050565b600061020583838080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250610e1692505050565b60006102088260125b60008281808080808080808080805b8b51811015610b30578b517f3000000000000000000000000000000000000000000000000000000000000000908d90839081106102eb57fe5b602001015160f860020a900460f860020a02600160f860020a0319161015801561035f57508b517f3900000000000000000000000000000000000000000000000000000000000000908d908390811061034057fe5b602001015160f860020a900460f860020a02600160f860020a03191611155b8015610369575083155b156104475784156103de576103858a600a63ffffffff610e2316565b99506103d1603060f860020a0260f860020a90048d83815181106103a557fe5b602001015160f860020a900460f860020a0260f860020a90040360ff168b610e7f90919063ffffffff16565b9950600190970196610442565b600195506103f38b600a63ffffffff610e2316565b9a5061043f603060f860020a0260f860020a90048d838151811061041357fe5b602001015160f860020a900460f860020a0260f860020a90040360ff168c610e7f90919063ffffffff16565b9a505b610b28565b8b517f3000000000000000000000000000000000000000000000000000000000000000908d908390811061047757fe5b602001015160f860020a900460f860020a02600160f860020a031916101580156104eb57508b517f3900000000000000000000000000000000000000000000000000000000000000908d90839081106104cc57fe5b602001015160f860020a900460f860020a02600160f860020a03191611155b80156104f45750835b1561055d5761050a89600a63ffffffff610e2316565b9850610556603060f860020a0260f860020a90048d838151811061052a57fe5b602001015160f860020a900460f860020a0260f860020a90040360ff168a610e7f90919063ffffffff16565b9850610b28565b8b517f2e00000000000000000000000000000000000000000000000000000000000000908d908390811061058d57fe5b602001015160f860020a900460f860020a02600160f860020a03191614156106b95785610604576040805160e560020a62461bcd02815260206004820152601560248201527f6d697373696e6720696e74656772616c20706172740000000000000000000000604482015290519081900360640190fd5b841561065a576040805160e560020a62461bcd02815260206004820152601760248201527f6475706c696361746520646563696d616c20706f696e74000000000000000000604482015290519081900360640190fd5b83156106b0576040805160e560020a62461bcd02815260206004820152601660248201527f646563696d616c206166746572206578706f6e656e7400000000000000000000604482015290519081900360640190fd5b60019450610b28565b8b517f2d00000000000000000000000000000000000000000000000000000000000000908d90839081106106e957fe5b602001015160f860020a900460f860020a02600160f860020a031916141561081a578215610761576040805160e560020a62461bcd02815260206004820152600b60248201527f6475706c6963617465202d000000000000000000000000000000000000000000604482015290519081900360640190fd5b81156107b7576040805160e560020a62461bcd02815260206004820152600a60248201527f6578747261207369676e00000000000000000000000000000000000000000000604482015290519081900360640190fd5b808760010114610811576040805160e560020a62461bcd02815260206004820152601e60248201527f2d207369676e206e6f7420696d6d6564696174656c7920616674657220650000604482015290519081900360640190fd5b60019250610b28565b8b517f2b00000000000000000000000000000000000000000000000000000000000000908d908390811061084a57fe5b602001015160f860020a900460f860020a02600160f860020a031916141561097b5781156108c2576040805160e560020a62461bcd02815260206004820152600b60248201527f6475706c6963617465202b000000000000000000000000000000000000000000604482015290519081900360640190fd5b8215610918576040805160e560020a62461bcd02815260206004820152600a60248201527f6578747261207369676e00000000000000000000000000000000000000000000604482015290519081900360640190fd5b808760010114610972576040805160e560020a62461bcd02815260206004820152601e60248201527f2b207369676e206e6f7420696d6d6564696174656c7920616674657220650000604482015290519081900360640190fd5b60019150610b28565b8b517f4500000000000000000000000000000000000000000000000000000000000000908d90839081106109ab57fe5b602001015160f860020a900460f860020a02600160f860020a0319161480610a1c57508b517f6500000000000000000000000000000000000000000000000000000000000000908d90839081106109fe57fe5b602001015160f860020a900460f860020a02600160f860020a031916145b15610ad85785610a76576040805160e560020a62461bcd02815260206004820152601560248201527f6d697373696e6720696e74656772616c20706172740000000000000000000000604482015290519081900360640190fd5b8315610acc576040805160e560020a62461bcd02815260206004820152601960248201527f6475706c6963617465206578706f6e656e742073796d626f6c00000000000000604482015290519081900360640190fd5b60019350809650610b28565b6040805160e560020a62461bcd02815260206004820152600d60248201527f696e76616c696420646967697400000000000000000000000000000000000000604482015290519081900360640190fd5b6001016102b2565b8280610b395750815b15610b5257866002018111610b4d57600080fd5b610b67565b8315610b6757866001018111610b6757600080fd5b8215610bfb578d8910610bf157604e8e8a0310610bce576040805160e560020a62461bcd02815260206004820152600d60248201527f6578706f6e656e74203e20373700000000000000000000000000000000000000604482015290519081900360640190fd5b8d8903600a0a8b81610bdc57fe5b049c506102089b505050505050505050505050565b888e039d50610c0e565b610c0b8e8a63ffffffff610e7f16565b9d505b878e10610cf857604e8810610c575760405160e560020a62461bcd028152600401808060200182810382526022815260200180610edd6022913960400191505060405180910390fd5b610c6b8b600a8a900a63ffffffff610e2316565b9a50610c7d8b8b63ffffffff610e7f16565b9a50604e888f0310610cd9576040805160e560020a62461bcd02815260206004820152600d60248201527f6578706f6e656e74203e20373700000000000000000000000000000000000000604482015290519081900360640190fd5b610cf1888f03600a0a8c610e2390919063ffffffff16565b9a50610dbb565b8d88039750604e8810610d3f5760405160e560020a62461bcd028152600401808060200182810382526022815260200180610edd6022913960400191505060405180910390fd5b87600a0a8a81610d4b57fe5b049950604e8e10610d905760405160e560020a62461bcd028152600401808060200182810382526022815260200180610edd6022913960400191505060405180910390fd5b610da68e600a0a8c610e2390919063ffffffff16565b9a50610db88b8b63ffffffff610e7f16565b9a505b66400000000000008b10610e035760405160e560020a62461bcd02815260040180806020018281038252603c815260200180610f20603c913960400191505060405180910390fd5b50989d9c50505050505050505050505050565b60006102088260006102a3565b600082610e3257506000610208565b82820282848281610e3f57fe5b04146102055760405160e560020a62461bcd028152600401808060200182810382526021815260200180610eff6021913960400191505060405180910390fd5b600082820183811015610205576040805160e560020a62461bcd02815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fdfe6d6f7265207468616e20373720646563696d616c2064696769747320706172736564536166654d6174683a206d756c7469706c69636174696f6e206f766572666c6f776e756d626572206578636565646564206d6178696d756d20616c6c6f7765642076616c756520666f722073616665206a736f6e206465636f64696e67a265627a7a723058205f81b7ec8e3f1940d843438d4f5703c7550cb32a3854813bc478a5dcfa539ba464736f6c634300050a0032`

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
