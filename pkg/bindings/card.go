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

// CardABI is the input ABI used to generate the binding from.
const CardABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"release\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"name\":\"available\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"overdraft\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"settle\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"debit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"hold\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unlock\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"o\",\"type\":\"address\"}],\"name\":\"changeOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"unlockAt\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"name\":\"balance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"controller\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"lock\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"c\",\"type\":\"address\"},{\"name\":\"o\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"}]"

// CardBin is the compiled bytecode used for deploying new contracts.
const CardBin = `6080604052600060025534801561001557600080fd5b506040516040806112288339810180604052810190808051906020019092919080519060200190929190505050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050611154806100d46000396000f3006080604052600436106100d0576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680630357371d146100d257806310098ad51461011f5780631218f3751461017657806315afd409146101cd578063636be27a1461021a5780638da5cb5b14610267578063977a5ec5146102be578063a69df4b51461030b578063a6f9dae114610322578063aa5dec6f14610365578063e3d670d714610390578063f3fef3a3146103e7578063f77c479114610434578063f83d08ba1461048b575b005b3480156100de57600080fd5b5061011d600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506104a2565b005b34801561012b57600080fd5b50610160600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506105e2565b6040518082815260200191505060405180910390f35b34801561018257600080fd5b506101b7600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061068f565b6040518082815260200191505060405180910390f35b3480156101d957600080fd5b50610218600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506106a7565b005b34801561022657600080fd5b50610265600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019092919050505061089e565b005b34801561027357600080fd5b5061027c610a25565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156102ca57600080fd5b50610309600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610a4b565b005b34801561031757600080fd5b50610320610b88565b005b34801561032e57600080fd5b50610363600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610bf4565b005b34801561037157600080fd5b5061037a610cef565b6040518082815260200191505060405180910390f35b34801561039c57600080fd5b506103d1600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610cf5565b6040518082815260200191505060405180910390f35b3480156103f357600080fd5b50610432600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610e13565b005b34801561044057600080fd5b50610449610ed8565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561049757600080fd5b506104a0610efd565b005b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff168073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156104fe57600080fd5b81600360008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205411156105975781600360008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825403925050819055506105dd565b6000600360008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505b505050565b6000806105ee83610cf5565b9050600360008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054811115156106415760009150610689565b600360008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054810390508091505b50919050565b60036020528060005260406000206000915090505481565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561070557600080fd5b600360008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205483111561078f57600360008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205492505b600083141561079d57610898565b6107a684610cf5565b915060008214156107b657610898565b818311156108135781600360008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254039250508190555081925061086b565b600083111561086a5782600360008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825403925050819055505b5b6108976000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff168585610fbe565b5b50505050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156108fc57600080fd5b610905846105e2565b9150818311156109f357600360008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054828403600360008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205401101515156109a057600080fd5b818303600360008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825401925050819055508192505b610a1f6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff168585610fbe565b50505050565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff168073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610aa757600080fd5b600360008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205482600360008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020540110151515610b3657600080fd5b81600360008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282540192505081905550505050565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610be557600080fd5b62013880430160028190555050565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161480610c9e57508073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16145b1515610ca957600080fd5b82600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505050565b60025481565b6000808273ffffffffffffffffffffffffffffffffffffffff16141515610df3578173ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050602060405180830381600087803b158015610db157600080fd5b505af1158015610dc5573d6000803e3d6000fd5b505050506040513d6020811015610ddb57600080fd5b81019080805190602001909291905050509050610e0e565b3073ffffffffffffffffffffffffffffffffffffffff163190505b919050565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610e7057600080fd5b6000600254118015610e8457506002544310155b1515610e8f57600080fd5b610e98836105e2565b8211151515610ea657600080fd5b610ed3600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168484610fbe565b505050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161480610fa757508073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16145b1515610fb257600080fd5b60006002819055505050565b6000811415610fcc57611123565b60008273ffffffffffffffffffffffffffffffffffffffff161415156110da578173ffffffffffffffffffffffffffffffffffffffff1663a9059cbb84836040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b15801561108f57600080fd5b505af11580156110a3573d6000803e3d6000fd5b505050506040513d60208110156110b957600080fd5b810190808051906020019092919050505015156110d557600080fd5b611122565b8273ffffffffffffffffffffffffffffffffffffffff166108fc829081150290604051600060405180830381858888f19350505050158015611120573d6000803e3d6000fd5b505b5b5050505600a165627a7a723058201af14fb3e37eee08fdbf3b18a0d93a3cd421ed58ba852b38d70c6e40014f5ef40029`

// DeployCard deploys a new Ethereum contract, binding an instance of Card to it.
func DeployCard(auth *bind.TransactOpts, backend bind.ContractBackend, c common.Address, o common.Address) (common.Address, *types.Transaction, *Card, error) {
	parsed, err := abi.JSON(strings.NewReader(CardABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CardBin), backend, c, o)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Card{CardCaller: CardCaller{contract: contract}, CardTransactor: CardTransactor{contract: contract}, CardFilterer: CardFilterer{contract: contract}}, nil
}

// Card is an auto generated Go binding around an Ethereum contract.
type Card struct {
	CardCaller     // Read-only binding to the contract
	CardTransactor // Write-only binding to the contract
	CardFilterer   // Log filterer for contract events
}

// CardCaller is an auto generated read-only Go binding around an Ethereum contract.
type CardCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CardTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CardTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CardFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CardFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CardSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CardSession struct {
	Contract     *Card             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CardCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CardCallerSession struct {
	Contract *CardCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// CardTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CardTransactorSession struct {
	Contract     *CardTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CardRaw is an auto generated low-level Go binding around an Ethereum contract.
type CardRaw struct {
	Contract *Card // Generic contract binding to access the raw methods on
}

// CardCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CardCallerRaw struct {
	Contract *CardCaller // Generic read-only contract binding to access the raw methods on
}

// CardTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CardTransactorRaw struct {
	Contract *CardTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCard creates a new instance of Card, bound to a specific deployed contract.
func NewCard(address common.Address, backend bind.ContractBackend) (*Card, error) {
	contract, err := bindCard(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Card{CardCaller: CardCaller{contract: contract}, CardTransactor: CardTransactor{contract: contract}, CardFilterer: CardFilterer{contract: contract}}, nil
}

// NewCardCaller creates a new read-only instance of Card, bound to a specific deployed contract.
func NewCardCaller(address common.Address, caller bind.ContractCaller) (*CardCaller, error) {
	contract, err := bindCard(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CardCaller{contract: contract}, nil
}

// NewCardTransactor creates a new write-only instance of Card, bound to a specific deployed contract.
func NewCardTransactor(address common.Address, transactor bind.ContractTransactor) (*CardTransactor, error) {
	contract, err := bindCard(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CardTransactor{contract: contract}, nil
}

// NewCardFilterer creates a new log filterer instance of Card, bound to a specific deployed contract.
func NewCardFilterer(address common.Address, filterer bind.ContractFilterer) (*CardFilterer, error) {
	contract, err := bindCard(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CardFilterer{contract: contract}, nil
}

// bindCard binds a generic wrapper to an already deployed contract.
func bindCard(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CardABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Card *CardRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Card.Contract.CardCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Card *CardRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Card.Contract.CardTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Card *CardRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Card.Contract.CardTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Card *CardCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Card.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Card *CardTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Card.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Card *CardTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Card.Contract.contract.Transact(opts, method, params...)
}

// Available is a free data retrieval call binding the contract method 0x10098ad5.
//
// Solidity: function available(token address) constant returns(uint256)
func (_Card *CardCaller) Available(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Card.contract.Call(opts, out, "available", token)
	return *ret0, err
}

// Available is a free data retrieval call binding the contract method 0x10098ad5.
//
// Solidity: function available(token address) constant returns(uint256)
func (_Card *CardSession) Available(token common.Address) (*big.Int, error) {
	return _Card.Contract.Available(&_Card.CallOpts, token)
}

// Available is a free data retrieval call binding the contract method 0x10098ad5.
//
// Solidity: function available(token address) constant returns(uint256)
func (_Card *CardCallerSession) Available(token common.Address) (*big.Int, error) {
	return _Card.Contract.Available(&_Card.CallOpts, token)
}

// Balance is a free data retrieval call binding the contract method 0xe3d670d7.
//
// Solidity: function balance(token address) constant returns(uint256)
func (_Card *CardCaller) Balance(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Card.contract.Call(opts, out, "balance", token)
	return *ret0, err
}

// Balance is a free data retrieval call binding the contract method 0xe3d670d7.
//
// Solidity: function balance(token address) constant returns(uint256)
func (_Card *CardSession) Balance(token common.Address) (*big.Int, error) {
	return _Card.Contract.Balance(&_Card.CallOpts, token)
}

// Balance is a free data retrieval call binding the contract method 0xe3d670d7.
//
// Solidity: function balance(token address) constant returns(uint256)
func (_Card *CardCallerSession) Balance(token common.Address) (*big.Int, error) {
	return _Card.Contract.Balance(&_Card.CallOpts, token)
}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() constant returns(address)
func (_Card *CardCaller) Controller(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Card.contract.Call(opts, out, "controller")
	return *ret0, err
}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() constant returns(address)
func (_Card *CardSession) Controller() (common.Address, error) {
	return _Card.Contract.Controller(&_Card.CallOpts)
}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() constant returns(address)
func (_Card *CardCallerSession) Controller() (common.Address, error) {
	return _Card.Contract.Controller(&_Card.CallOpts)
}

// Overdraft is a free data retrieval call binding the contract method 0x1218f375.
//
// Solidity: function overdraft( address) constant returns(uint256)
func (_Card *CardCaller) Overdraft(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Card.contract.Call(opts, out, "overdraft", arg0)
	return *ret0, err
}

// Overdraft is a free data retrieval call binding the contract method 0x1218f375.
//
// Solidity: function overdraft( address) constant returns(uint256)
func (_Card *CardSession) Overdraft(arg0 common.Address) (*big.Int, error) {
	return _Card.Contract.Overdraft(&_Card.CallOpts, arg0)
}

// Overdraft is a free data retrieval call binding the contract method 0x1218f375.
//
// Solidity: function overdraft( address) constant returns(uint256)
func (_Card *CardCallerSession) Overdraft(arg0 common.Address) (*big.Int, error) {
	return _Card.Contract.Overdraft(&_Card.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Card *CardCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Card.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Card *CardSession) Owner() (common.Address, error) {
	return _Card.Contract.Owner(&_Card.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Card *CardCallerSession) Owner() (common.Address, error) {
	return _Card.Contract.Owner(&_Card.CallOpts)
}

// UnlockAt is a free data retrieval call binding the contract method 0xaa5dec6f.
//
// Solidity: function unlockAt() constant returns(uint256)
func (_Card *CardCaller) UnlockAt(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Card.contract.Call(opts, out, "unlockAt")
	return *ret0, err
}

// UnlockAt is a free data retrieval call binding the contract method 0xaa5dec6f.
//
// Solidity: function unlockAt() constant returns(uint256)
func (_Card *CardSession) UnlockAt() (*big.Int, error) {
	return _Card.Contract.UnlockAt(&_Card.CallOpts)
}

// UnlockAt is a free data retrieval call binding the contract method 0xaa5dec6f.
//
// Solidity: function unlockAt() constant returns(uint256)
func (_Card *CardCallerSession) UnlockAt() (*big.Int, error) {
	return _Card.Contract.UnlockAt(&_Card.CallOpts)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(o address) returns()
func (_Card *CardTransactor) ChangeOwner(opts *bind.TransactOpts, o common.Address) (*types.Transaction, error) {
	return _Card.contract.Transact(opts, "changeOwner", o)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(o address) returns()
func (_Card *CardSession) ChangeOwner(o common.Address) (*types.Transaction, error) {
	return _Card.Contract.ChangeOwner(&_Card.TransactOpts, o)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(o address) returns()
func (_Card *CardTransactorSession) ChangeOwner(o common.Address) (*types.Transaction, error) {
	return _Card.Contract.ChangeOwner(&_Card.TransactOpts, o)
}

// Debit is a paid mutator transaction binding the contract method 0x636be27a.
//
// Solidity: function debit(token address, amount uint256) returns()
func (_Card *CardTransactor) Debit(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Card.contract.Transact(opts, "debit", token, amount)
}

// Debit is a paid mutator transaction binding the contract method 0x636be27a.
//
// Solidity: function debit(token address, amount uint256) returns()
func (_Card *CardSession) Debit(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Card.Contract.Debit(&_Card.TransactOpts, token, amount)
}

// Debit is a paid mutator transaction binding the contract method 0x636be27a.
//
// Solidity: function debit(token address, amount uint256) returns()
func (_Card *CardTransactorSession) Debit(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Card.Contract.Debit(&_Card.TransactOpts, token, amount)
}

// Hold is a paid mutator transaction binding the contract method 0x977a5ec5.
//
// Solidity: function hold(token address, amount uint256) returns()
func (_Card *CardTransactor) Hold(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Card.contract.Transact(opts, "hold", token, amount)
}

// Hold is a paid mutator transaction binding the contract method 0x977a5ec5.
//
// Solidity: function hold(token address, amount uint256) returns()
func (_Card *CardSession) Hold(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Card.Contract.Hold(&_Card.TransactOpts, token, amount)
}

// Hold is a paid mutator transaction binding the contract method 0x977a5ec5.
//
// Solidity: function hold(token address, amount uint256) returns()
func (_Card *CardTransactorSession) Hold(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Card.Contract.Hold(&_Card.TransactOpts, token, amount)
}

// Lock is a paid mutator transaction binding the contract method 0xf83d08ba.
//
// Solidity: function lock() returns()
func (_Card *CardTransactor) Lock(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Card.contract.Transact(opts, "lock")
}

// Lock is a paid mutator transaction binding the contract method 0xf83d08ba.
//
// Solidity: function lock() returns()
func (_Card *CardSession) Lock() (*types.Transaction, error) {
	return _Card.Contract.Lock(&_Card.TransactOpts)
}

// Lock is a paid mutator transaction binding the contract method 0xf83d08ba.
//
// Solidity: function lock() returns()
func (_Card *CardTransactorSession) Lock() (*types.Transaction, error) {
	return _Card.Contract.Lock(&_Card.TransactOpts)
}

// Release is a paid mutator transaction binding the contract method 0x0357371d.
//
// Solidity: function release(token address, amount uint256) returns()
func (_Card *CardTransactor) Release(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Card.contract.Transact(opts, "release", token, amount)
}

// Release is a paid mutator transaction binding the contract method 0x0357371d.
//
// Solidity: function release(token address, amount uint256) returns()
func (_Card *CardSession) Release(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Card.Contract.Release(&_Card.TransactOpts, token, amount)
}

// Release is a paid mutator transaction binding the contract method 0x0357371d.
//
// Solidity: function release(token address, amount uint256) returns()
func (_Card *CardTransactorSession) Release(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Card.Contract.Release(&_Card.TransactOpts, token, amount)
}

// Settle is a paid mutator transaction binding the contract method 0x15afd409.
//
// Solidity: function settle(token address, amount uint256) returns()
func (_Card *CardTransactor) Settle(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Card.contract.Transact(opts, "settle", token, amount)
}

// Settle is a paid mutator transaction binding the contract method 0x15afd409.
//
// Solidity: function settle(token address, amount uint256) returns()
func (_Card *CardSession) Settle(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Card.Contract.Settle(&_Card.TransactOpts, token, amount)
}

// Settle is a paid mutator transaction binding the contract method 0x15afd409.
//
// Solidity: function settle(token address, amount uint256) returns()
func (_Card *CardTransactorSession) Settle(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Card.Contract.Settle(&_Card.TransactOpts, token, amount)
}

// Unlock is a paid mutator transaction binding the contract method 0xa69df4b5.
//
// Solidity: function unlock() returns()
func (_Card *CardTransactor) Unlock(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Card.contract.Transact(opts, "unlock")
}

// Unlock is a paid mutator transaction binding the contract method 0xa69df4b5.
//
// Solidity: function unlock() returns()
func (_Card *CardSession) Unlock() (*types.Transaction, error) {
	return _Card.Contract.Unlock(&_Card.TransactOpts)
}

// Unlock is a paid mutator transaction binding the contract method 0xa69df4b5.
//
// Solidity: function unlock() returns()
func (_Card *CardTransactorSession) Unlock() (*types.Transaction, error) {
	return _Card.Contract.Unlock(&_Card.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(token address, amount uint256) returns()
func (_Card *CardTransactor) Withdraw(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Card.contract.Transact(opts, "withdraw", token, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(token address, amount uint256) returns()
func (_Card *CardSession) Withdraw(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Card.Contract.Withdraw(&_Card.TransactOpts, token, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(token address, amount uint256) returns()
func (_Card *CardTransactorSession) Withdraw(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Card.Contract.Withdraw(&_Card.TransactOpts, token, amount)
}
