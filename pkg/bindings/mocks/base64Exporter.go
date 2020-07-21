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

// Base64ExporterABI is the input ABI used to generate the binding from.
const Base64ExporterABI = "[{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_encoded\",\"type\":\"bytes\"}],\"name\":\"base64decode\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// Base64ExporterBin is the compiled bytecode used for deploying new contracts.
var Base64ExporterBin = "0x608060405234801561001057600080fd5b50610842806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c80633a718bfb14610030575b600080fd5b6100a06004803603602081101561004657600080fd5b81019060208101813564010000000081111561006157600080fd5b82018360208201111561007357600080fd5b8035906020019184600183028401116401000000008311171561009557600080fd5b509092509050610115565b6040805160208082528351818301528351919283929083019185019080838360005b838110156100da5781810151838201526020016100c2565b50505050905090810190601f1680156101075780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b606061015683838080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061015d92505050565b9392505050565b606060008060008060008651905060608167ffffffffffffffff8111801561018457600080fd5b506040519080825280601f01601f1916602001820160405280156101af576020820181803683370190505b509050600080831180156101c4575060048306155b610215576040805162461bcd60e51b815260206004820152601760248201527f696e76616c69642062617365363420656e636f64696e67000000000000000000604482015290519081900360640190fd5b60408051603d60f81b8152905190819003600101902089518a90600119860190811061023d57fe5b01602090810151604080516001600160f81b03199092168284015280518083036001018152602190920190528051910120141561027f576002830392506102e5565b60408051603d60f81b8152905190819003600101902089518a9060001986019081106102a757fe5b01602090810151604080516001600160f81b0319909216828401528051808303600101815260219092019052805191012014156102e5576001830392505b600319831660005b81811015610503576040518060a00160405280607b8152602001610792607b91398b5160018301928d91811061031f57fe5b0160200151815160f89190911c90811061033557fe5b602001015160f81c60f81b98506040518060a00160405280607b8152602001610792607b91398b5160018301928d91811061036c57fe5b0160200151815160f89190911c90811061038257fe5b602001015160f81c60f81b97506040518060a00160405280607b8152602001610792607b91398b5160018301928d9181106103b957fe5b0160200151815160f89190911c9081106103cf57fe5b602001015160f81c60f81b96506040518060a00160405280607b8152602001610792607b91398b5160018301928d91811061040657fe5b0160200151815160f89190911c90811061041c57fe5b016020015184516001600160f81b031991821697506001850194603f60fa1b60028d901b1660ff60f41b60048d901c1617909216918691811061045b57fe5b60200101906001600160f81b031916908160001a90535083516001840193600f60fc1b60048b901b1660ff60f61b60028b901c16176001600160f81b0319169186919081106104a657fe5b60200101906001600160f81b031916908160001a90535083516001840193600360fe1b60068a901b1688176001600160f81b0319169186919081106104e757fe5b60200101906001600160f81b031916908160001a9053506102ed565b818503600214156105fb576040518060a00160405280607b8152602001610792607b91398b5160018301928d91811061053857fe5b0160200151815160f89190911c90811061054e57fe5b602001015160f81c60f81b98506040518060a00160405280607b8152602001610792607b91398b5160018301928d91811061058557fe5b0160200151815160f89190911c90811061059b57fe5b602001015160f81c60f81b97506004886001600160f81b031916901c60028a6001600160f81b031916901b1760ff60f81b168484806001019550815181106105df57fe5b60200101906001600160f81b031916908160001a905350610782565b81850360031415610782576040518060a00160405280607b8152602001610792607b91398b5160018301928d91811061063057fe5b0160200151815160f89190911c90811061064657fe5b602001015160f81c60f81b98506040518060a00160405280607b8152602001610792607b91398b5160018301928d91811061067d57fe5b0160200151815160f89190911c90811061069357fe5b602001015160f81c60f81b97506040518060a00160405280607b8152602001610792607b91398b5160018301928d9181106106ca57fe5b0160200151815160f89190911c9081106106e057fe5b016020015184516001600160f81b031991821698506001850194603f60fa1b60028d901b1660ff60f41b60048d901c1617909216918691811061071f57fe5b60200101906001600160f81b031916908160001a90535083516001840193600f60fc1b60048b901b1660ff60f61b60028b901c16176001600160f81b03191691869190811061076a57fe5b60200101906001600160f81b031916908160001a9053505b5050815297965050505050505056fe000000000000000000000000000000000000000000000000000000000000000000000000000000000000003e003e003f3435363738393a3b3c3d00000000000000000102030405060708090a0b0c0d0e0f10111213141516171819000000003f001a1b1c1d1e1f202122232425262728292a2b2c2d2e2f30313233a2646970667358221220b28e0962f4dd8b5b48ed955984cb32efdf725280b9ff4a7a91cdc9048859856a64736f6c634300060b0033"

// DeployBase64Exporter deploys a new Ethereum contract, binding an instance of Base64Exporter to it.
func DeployBase64Exporter(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Base64Exporter, error) {
	parsed, err := abi.JSON(strings.NewReader(Base64ExporterABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(Base64ExporterBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Base64Exporter{Base64ExporterCaller: Base64ExporterCaller{contract: contract}, Base64ExporterTransactor: Base64ExporterTransactor{contract: contract}, Base64ExporterFilterer: Base64ExporterFilterer{contract: contract}}, nil
}

// Base64Exporter is an auto generated Go binding around an Ethereum contract.
type Base64Exporter struct {
	Base64ExporterCaller     // Read-only binding to the contract
	Base64ExporterTransactor // Write-only binding to the contract
	Base64ExporterFilterer   // Log filterer for contract events
}

// Base64ExporterCaller is an auto generated read-only Go binding around an Ethereum contract.
type Base64ExporterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Base64ExporterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Base64ExporterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Base64ExporterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Base64ExporterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Base64ExporterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Base64ExporterSession struct {
	Contract     *Base64Exporter   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Base64ExporterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Base64ExporterCallerSession struct {
	Contract *Base64ExporterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// Base64ExporterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Base64ExporterTransactorSession struct {
	Contract     *Base64ExporterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// Base64ExporterRaw is an auto generated low-level Go binding around an Ethereum contract.
type Base64ExporterRaw struct {
	Contract *Base64Exporter // Generic contract binding to access the raw methods on
}

// Base64ExporterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Base64ExporterCallerRaw struct {
	Contract *Base64ExporterCaller // Generic read-only contract binding to access the raw methods on
}

// Base64ExporterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Base64ExporterTransactorRaw struct {
	Contract *Base64ExporterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBase64Exporter creates a new instance of Base64Exporter, bound to a specific deployed contract.
func NewBase64Exporter(address common.Address, backend bind.ContractBackend) (*Base64Exporter, error) {
	contract, err := bindBase64Exporter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Base64Exporter{Base64ExporterCaller: Base64ExporterCaller{contract: contract}, Base64ExporterTransactor: Base64ExporterTransactor{contract: contract}, Base64ExporterFilterer: Base64ExporterFilterer{contract: contract}}, nil
}

// NewBase64ExporterCaller creates a new read-only instance of Base64Exporter, bound to a specific deployed contract.
func NewBase64ExporterCaller(address common.Address, caller bind.ContractCaller) (*Base64ExporterCaller, error) {
	contract, err := bindBase64Exporter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Base64ExporterCaller{contract: contract}, nil
}

// NewBase64ExporterTransactor creates a new write-only instance of Base64Exporter, bound to a specific deployed contract.
func NewBase64ExporterTransactor(address common.Address, transactor bind.ContractTransactor) (*Base64ExporterTransactor, error) {
	contract, err := bindBase64Exporter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Base64ExporterTransactor{contract: contract}, nil
}

// NewBase64ExporterFilterer creates a new log filterer instance of Base64Exporter, bound to a specific deployed contract.
func NewBase64ExporterFilterer(address common.Address, filterer bind.ContractFilterer) (*Base64ExporterFilterer, error) {
	contract, err := bindBase64Exporter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Base64ExporterFilterer{contract: contract}, nil
}

// bindBase64Exporter binds a generic wrapper to an already deployed contract.
func bindBase64Exporter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Base64ExporterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Base64Exporter *Base64ExporterRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Base64Exporter.Contract.Base64ExporterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Base64Exporter *Base64ExporterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Base64Exporter.Contract.Base64ExporterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Base64Exporter *Base64ExporterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Base64Exporter.Contract.Base64ExporterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Base64Exporter *Base64ExporterCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Base64Exporter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Base64Exporter *Base64ExporterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Base64Exporter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Base64Exporter *Base64ExporterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Base64Exporter.Contract.contract.Transact(opts, method, params...)
}

// Base64decode is a free data retrieval call binding the contract method 0x3a718bfb.
//
// Solidity: function base64decode(bytes _encoded) constant returns(bytes)
func (_Base64Exporter *Base64ExporterCaller) Base64decode(opts *bind.CallOpts, _encoded []byte) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Base64Exporter.contract.Call(opts, out, "base64decode", _encoded)
	return *ret0, err
}

// Base64decode is a free data retrieval call binding the contract method 0x3a718bfb.
//
// Solidity: function base64decode(bytes _encoded) constant returns(bytes)
func (_Base64Exporter *Base64ExporterSession) Base64decode(_encoded []byte) ([]byte, error) {
	return _Base64Exporter.Contract.Base64decode(&_Base64Exporter.CallOpts, _encoded)
}

// Base64decode is a free data retrieval call binding the contract method 0x3a718bfb.
//
// Solidity: function base64decode(bytes _encoded) constant returns(bytes)
func (_Base64Exporter *Base64ExporterCallerSession) Base64decode(_encoded []byte) ([]byte, error) {
	return _Base64Exporter.Contract.Base64decode(&_Base64Exporter.CallOpts, _encoded)
}
