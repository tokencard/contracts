// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// ControllerABI is the input ABI used to generate the binding from.
const ControllerABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"sequenceNumber\",\"type\":\"uint256\"},{\"name\":\"operations\",\"type\":\"bytes\"},{\"name\":\"cards\",\"type\":\"address[]\"},{\"name\":\"tokens\",\"type\":\"address[]\"},{\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"process\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"next\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"card\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"}],\"name\":\"reownCard\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"}],\"name\":\"changeOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"newOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"card\",\"type\":\"address\"}],\"name\":\"lock\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"}]"

// ControllerBin is the compiled bytecode used for deploying new contracts.
const ControllerBin = `0x6080604052600160025534801561001557600080fd5b50336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506112db806100656000396000f30060806040526004361061008e576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680634ad27ff7146100905780634c8fe5261461011d5780635a4423e51461014857806379ba5097146101ab5780638da5cb5b146101c2578063a6f9dae114610219578063d4ee1d901461025c578063f435f5a7146102b3575b005b34801561009c57600080fd5b5061011b600480360381019080803590602001909291908035906020019082018035906020019190919293919293908035906020019082018035906020019190919293919293908035906020019082018035906020019190919293919293908035906020019082018035906020019190919293919293905050506102f6565b005b34801561012957600080fd5b50610132610d90565b6040518082815260200191505060405180910390f35b34801561015457600080fd5b506101a9600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610d96565b005b3480156101b757600080fd5b506101c0610eaa565b005b3480156101ce57600080fd5b506101d7610f8c565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561022557600080fd5b5061025a600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610fb1565b005b34801561026857600080fd5b50610271611052565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156102bf57600080fd5b506102f4600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611078565b005b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561035457600080fd5b6002548b1115801561036c57506002548a8a90508c01115b151561037757600080fd5b878790508a8a905014151561038b57600080fd5b858590508a8a905014151561039f57600080fd5b838390508a8a90501415156103b357600080fd5b8a6002540391505b89899050821015610d775760007f0100000000000000000000000000000000000000000000000000000000000000027effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff19168a8a84818110151561041a57fe5b905001357f010000000000000000000000000000000000000000000000000000000000000090047f0100000000000000000000000000000000000000000000000000000000000000027effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff19167effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191614156105d55787878381811015156104bb57fe5b9050602002013573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663636be27a878785818110151561050157fe5b9050602002013573ffffffffffffffffffffffffffffffffffffffff16868686818110151561052c57fe5b905060200201356040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050600060405180830381600087803b1580156105b857600080fd5b505af11580156105cc573d6000803e3d6000fd5b50505050610d6a565b60017f0100000000000000000000000000000000000000000000000000000000000000027effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff19168a8a84818110151561062957fe5b905001357f010000000000000000000000000000000000000000000000000000000000000090047f0100000000000000000000000000000000000000000000000000000000000000027effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff19167effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191614156107345761072f88888481811015156106cd57fe5b9050602002013573ffffffffffffffffffffffffffffffffffffffff1687878581811015156106f857fe5b9050602002013573ffffffffffffffffffffffffffffffffffffffff16868686818110151561072357fe5b90506020020135611154565b610d69565b60027f0100000000000000000000000000000000000000000000000000000000000000027effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff19168a8a84818110151561078857fe5b905001357f010000000000000000000000000000000000000000000000000000000000000090047f0100000000000000000000000000000000000000000000000000000000000000027effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff19167effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916141561094357878783818110151561082957fe5b9050602002013573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663977a5ec5878785818110151561086f57fe5b9050602002013573ffffffffffffffffffffffffffffffffffffffff16868686818110151561089a57fe5b905060200201356040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050600060405180830381600087803b15801561092657600080fd5b505af115801561093a573d6000803e3d6000fd5b50505050610d68565b60037f0100000000000000000000000000000000000000000000000000000000000000027effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff19168a8a84818110151561099757fe5b905001357f010000000000000000000000000000000000000000000000000000000000000090047f0100000000000000000000000000000000000000000000000000000000000000027effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff19167effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff19161415610b52578787838181101515610a3857fe5b9050602002013573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16630357371d8787858181101515610a7e57fe5b9050602002013573ffffffffffffffffffffffffffffffffffffffff168686868181101515610aa957fe5b905060200201356040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050600060405180830381600087803b158015610b3557600080fd5b505af1158015610b49573d6000803e3d6000fd5b50505050610d67565b60047f0100000000000000000000000000000000000000000000000000000000000000027effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff19168a8a848181101515610ba657fe5b905001357f010000000000000000000000000000000000000000000000000000000000000090047f0100000000000000000000000000000000000000000000000000000000000000027effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff19167effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff19161415610d61578787838181101515610c4757fe5b9050602002013573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166315afd4098787858181101515610c8d57fe5b9050602002013573ffffffffffffffffffffffffffffffffffffffff168686868181101515610cb857fe5b905060200201356040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050600060405180830381600087803b158015610d4457600080fd5b505af1158015610d58573d6000803e3d6000fd5b50505050610d66565b600080fd5b5b5b5b5b81806001019250506103bb565b898990508b016002819055505050505050505050505050565b60025481565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff168073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610df257600080fd5b8273ffffffffffffffffffffffffffffffffffffffff1663a6f9dae1836040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050600060405180830381600087803b158015610e8d57600080fd5b505af1158015610ea1573d6000803e3d6000fd5b50505050505050565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610f0757600080fd5b336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff168073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561100d57600080fd5b81600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff168073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156110d457600080fd5b8173ffffffffffffffffffffffffffffffffffffffff1663f83d08ba6040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401600060405180830381600087803b15801561113857600080fd5b505af115801561114c573d6000803e3d6000fd5b505050505050565b60008273ffffffffffffffffffffffffffffffffffffffff16141515611262578173ffffffffffffffffffffffffffffffffffffffff1663a9059cbb84836040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b15801561121757600080fd5b505af115801561122b573d6000803e3d6000fd5b505050506040513d602081101561124157600080fd5b8101908080519060200190929190505050151561125d57600080fd5b6112aa565b8273ffffffffffffffffffffffffffffffffffffffff166108fc829081150290604051600060405180830381858888f193505050501580156112a8573d6000803e3d6000fd5b505b5050505600a165627a7a723058202149ef93a75e21886fd9b2abeb852f89a4ebe3d16672d8705b3882ba96db7da10029`

// DeployController deploys a new Ethereum contract, binding an instance of Controller to it.
func DeployController(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Controller, error) {
	parsed, err := abi.JSON(strings.NewReader(ControllerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ControllerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Controller{ControllerCaller: ControllerCaller{contract: contract}, ControllerTransactor: ControllerTransactor{contract: contract}, ControllerFilterer: ControllerFilterer{contract: contract}}, nil
}

// Controller is an auto generated Go binding around an Ethereum contract.
type Controller struct {
	ControllerCaller     // Read-only binding to the contract
	ControllerTransactor // Write-only binding to the contract
	ControllerFilterer   // Log filterer for contract events
}

// ControllerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ControllerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ControllerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ControllerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ControllerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ControllerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ControllerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ControllerSession struct {
	Contract     *Controller       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ControllerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ControllerCallerSession struct {
	Contract *ControllerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ControllerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ControllerTransactorSession struct {
	Contract     *ControllerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ControllerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ControllerRaw struct {
	Contract *Controller // Generic contract binding to access the raw methods on
}

// ControllerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ControllerCallerRaw struct {
	Contract *ControllerCaller // Generic read-only contract binding to access the raw methods on
}

// ControllerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ControllerTransactorRaw struct {
	Contract *ControllerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewController creates a new instance of Controller, bound to a specific deployed contract.
func NewController(address common.Address, backend bind.ContractBackend) (*Controller, error) {
	contract, err := bindController(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Controller{ControllerCaller: ControllerCaller{contract: contract}, ControllerTransactor: ControllerTransactor{contract: contract}, ControllerFilterer: ControllerFilterer{contract: contract}}, nil
}

// NewControllerCaller creates a new read-only instance of Controller, bound to a specific deployed contract.
func NewControllerCaller(address common.Address, caller bind.ContractCaller) (*ControllerCaller, error) {
	contract, err := bindController(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ControllerCaller{contract: contract}, nil
}

// NewControllerTransactor creates a new write-only instance of Controller, bound to a specific deployed contract.
func NewControllerTransactor(address common.Address, transactor bind.ContractTransactor) (*ControllerTransactor, error) {
	contract, err := bindController(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ControllerTransactor{contract: contract}, nil
}

// NewControllerFilterer creates a new log filterer instance of Controller, bound to a specific deployed contract.
func NewControllerFilterer(address common.Address, filterer bind.ContractFilterer) (*ControllerFilterer, error) {
	contract, err := bindController(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ControllerFilterer{contract: contract}, nil
}

// bindController binds a generic wrapper to an already deployed contract.
func bindController(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ControllerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Controller *ControllerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Controller.Contract.ControllerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Controller *ControllerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Controller.Contract.ControllerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Controller *ControllerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Controller.Contract.ControllerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Controller *ControllerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Controller.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Controller *ControllerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Controller.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Controller *ControllerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Controller.Contract.contract.Transact(opts, method, params...)
}

// NewOwner is a free data retrieval call binding the contract method 0xd4ee1d90.
//
// Solidity: function newOwner() constant returns(address)
func (_Controller *ControllerCaller) NewOwner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Controller.contract.Call(opts, out, "newOwner")
	return *ret0, err
}

// NewOwner is a free data retrieval call binding the contract method 0xd4ee1d90.
//
// Solidity: function newOwner() constant returns(address)
func (_Controller *ControllerSession) NewOwner() (common.Address, error) {
	return _Controller.Contract.NewOwner(&_Controller.CallOpts)
}

// NewOwner is a free data retrieval call binding the contract method 0xd4ee1d90.
//
// Solidity: function newOwner() constant returns(address)
func (_Controller *ControllerCallerSession) NewOwner() (common.Address, error) {
	return _Controller.Contract.NewOwner(&_Controller.CallOpts)
}

// Next is a free data retrieval call binding the contract method 0x4c8fe526.
//
// Solidity: function next() constant returns(uint256)
func (_Controller *ControllerCaller) Next(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Controller.contract.Call(opts, out, "next")
	return *ret0, err
}

// Next is a free data retrieval call binding the contract method 0x4c8fe526.
//
// Solidity: function next() constant returns(uint256)
func (_Controller *ControllerSession) Next() (*big.Int, error) {
	return _Controller.Contract.Next(&_Controller.CallOpts)
}

// Next is a free data retrieval call binding the contract method 0x4c8fe526.
//
// Solidity: function next() constant returns(uint256)
func (_Controller *ControllerCallerSession) Next() (*big.Int, error) {
	return _Controller.Contract.Next(&_Controller.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Controller *ControllerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Controller.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Controller *ControllerSession) Owner() (common.Address, error) {
	return _Controller.Contract.Owner(&_Controller.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Controller *ControllerCallerSession) Owner() (common.Address, error) {
	return _Controller.Contract.Owner(&_Controller.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Controller *ControllerTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Controller *ControllerSession) AcceptOwnership() (*types.Transaction, error) {
	return _Controller.Contract.AcceptOwnership(&_Controller.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Controller *ControllerTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _Controller.Contract.AcceptOwnership(&_Controller.TransactOpts)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(to address) returns()
func (_Controller *ControllerTransactor) ChangeOwner(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "changeOwner", to)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(to address) returns()
func (_Controller *ControllerSession) ChangeOwner(to common.Address) (*types.Transaction, error) {
	return _Controller.Contract.ChangeOwner(&_Controller.TransactOpts, to)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(to address) returns()
func (_Controller *ControllerTransactorSession) ChangeOwner(to common.Address) (*types.Transaction, error) {
	return _Controller.Contract.ChangeOwner(&_Controller.TransactOpts, to)
}

// Lock is a paid mutator transaction binding the contract method 0xf435f5a7.
//
// Solidity: function lock(card address) returns()
func (_Controller *ControllerTransactor) Lock(opts *bind.TransactOpts, card common.Address) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "lock", card)
}

// Lock is a paid mutator transaction binding the contract method 0xf435f5a7.
//
// Solidity: function lock(card address) returns()
func (_Controller *ControllerSession) Lock(card common.Address) (*types.Transaction, error) {
	return _Controller.Contract.Lock(&_Controller.TransactOpts, card)
}

// Lock is a paid mutator transaction binding the contract method 0xf435f5a7.
//
// Solidity: function lock(card address) returns()
func (_Controller *ControllerTransactorSession) Lock(card common.Address) (*types.Transaction, error) {
	return _Controller.Contract.Lock(&_Controller.TransactOpts, card)
}

// Process is a paid mutator transaction binding the contract method 0x4ad27ff7.
//
// Solidity: function process(sequenceNumber uint256, operations bytes, cards address[], tokens address[], amounts uint256[]) returns()
func (_Controller *ControllerTransactor) Process(opts *bind.TransactOpts, sequenceNumber *big.Int, operations []byte, cards []common.Address, tokens []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "process", sequenceNumber, operations, cards, tokens, amounts)
}

// Process is a paid mutator transaction binding the contract method 0x4ad27ff7.
//
// Solidity: function process(sequenceNumber uint256, operations bytes, cards address[], tokens address[], amounts uint256[]) returns()
func (_Controller *ControllerSession) Process(sequenceNumber *big.Int, operations []byte, cards []common.Address, tokens []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _Controller.Contract.Process(&_Controller.TransactOpts, sequenceNumber, operations, cards, tokens, amounts)
}

// Process is a paid mutator transaction binding the contract method 0x4ad27ff7.
//
// Solidity: function process(sequenceNumber uint256, operations bytes, cards address[], tokens address[], amounts uint256[]) returns()
func (_Controller *ControllerTransactorSession) Process(sequenceNumber *big.Int, operations []byte, cards []common.Address, tokens []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _Controller.Contract.Process(&_Controller.TransactOpts, sequenceNumber, operations, cards, tokens, amounts)
}

// ReownCard is a paid mutator transaction binding the contract method 0x5a4423e5.
//
// Solidity: function reownCard(card address, to address) returns()
func (_Controller *ControllerTransactor) ReownCard(opts *bind.TransactOpts, card common.Address, to common.Address) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "reownCard", card, to)
}

// ReownCard is a paid mutator transaction binding the contract method 0x5a4423e5.
//
// Solidity: function reownCard(card address, to address) returns()
func (_Controller *ControllerSession) ReownCard(card common.Address, to common.Address) (*types.Transaction, error) {
	return _Controller.Contract.ReownCard(&_Controller.TransactOpts, card, to)
}

// ReownCard is a paid mutator transaction binding the contract method 0x5a4423e5.
//
// Solidity: function reownCard(card address, to address) returns()
func (_Controller *ControllerTransactorSession) ReownCard(card common.Address, to common.Address) (*types.Transaction, error) {
	return _Controller.Contract.ReownCard(&_Controller.TransactOpts, card, to)
}
