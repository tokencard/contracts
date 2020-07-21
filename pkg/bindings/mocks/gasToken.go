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

// GasTokenABI is the input ABI used to generate the binding from.
const GasTokenABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"free\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"freeUpTo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalBurned\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalMinted\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// GasTokenBin is the compiled bytecode used for deploying new contracts.
var GasTokenBin = "0x608060405234801561001057600080fd5b50610675806100206000396000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c8063a2309ff81161005b578063a2309ff814610109578063a9059cbb14610111578063d89135cd1461013d578063d8ccd0f31461014557610088565b806318160ddd1461008d5780636366b936146100a757806370a08231146100c4578063a0712d68146100ea575b600080fd5b610095610162565b60408051918252519081900360200190f35b610095600480360360208110156100bd57600080fd5b503561016c565b610095600480360360208110156100da57600080fd5b50356001600160a01b031661018e565b6101076004803603602081101561010057600080fd5b50356101a9565b005b61009561037e565b6101076004803603604081101561012757600080fd5b506001600160a01b038135169060200135610384565b6100956103f6565b6100956004803603602081101561015b57600080fd5b50356103fc565b6001546000540390565b60006101886101838361017e3361018e565b61041b565b6103fc565b92915050565b6001600160a01b031660009081526002602052604090205490565b600080547f766ffa233a79675b0530301caf58abcfa2eb3318585733ff60005260176009f3909152602082045b801561034957816020600080f550600182016020600080f550600282016020600080f550600382016020600080f550600482016020600080f550600582016020600080f550600682016020600080f550600782016020600080f550600882016020600080f550600982016020600080f550600a82016020600080f550600b82016020600080f550600c82016020600080f550600d82016020600080f550600e82016020600080f550600f82016020600080f550601082016020600080f550601182016020600080f550601282016020600080f550601382016020600080f550601482016020600080f550601582016020600080f550601682016020600080f550601782016020600080f550601882016020600080f550601982016020600080f550601a82016020600080f550601b82016020600080f550601c82016020600080f550601d82016020600080f550601e82016020600080f550601f82016020600080f55060209190910190600019016101d6565b50601f82165b801561036d57816020600080f550600191909101906000190161034f565b506103783383610433565b60005550565b60005481565b336000908152600260205260409020546103a4908263ffffffff61045c16565b33600090815260026020526040808220929092556001600160a01b038416815220546103d6908263ffffffff61049e16565b6001600160a01b0390921660009081526002602052604090209190915550565b60015481565b600081156104175761040e33836104f8565b61041782610521565b5090565b600081831061042a578161042c565b825b9392505050565b6001600160a01b0382166000908152600260205260409020546103d6908263ffffffff61049e16565b600061042c83836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f7700008152506105a8565b60008282018381101561042c576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b6001600160a01b0382166000908152600260205260409020546103d6908263ffffffff61045c16565b600154818101806001556040517fff00000000fa233a79675b0530301caf58abcfa2eb000000000000000000000081527f841da0d3b4b49d75c2a11068e21bceeb2e5d8c9e31ab7cea45c9ce114a2033dc6035820152601581015b828410156105a1578381526000806000806000605587205af15060018401935061057c565b5050505050565b600081848411156106375760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b838110156105fc5781810151838201526020016105e4565b50505050905090810190601f1680156106295780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b50505090039056fea26469706673582212208c384d4bc27576510e8787e2a026d77dee543423695bcc31bc89c0260804706764736f6c634300060b0033"

// DeployGasToken deploys a new Ethereum contract, binding an instance of GasToken to it.
func DeployGasToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *GasToken, error) {
	parsed, err := abi.JSON(strings.NewReader(GasTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(GasTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &GasToken{GasTokenCaller: GasTokenCaller{contract: contract}, GasTokenTransactor: GasTokenTransactor{contract: contract}, GasTokenFilterer: GasTokenFilterer{contract: contract}}, nil
}

// GasToken is an auto generated Go binding around an Ethereum contract.
type GasToken struct {
	GasTokenCaller     // Read-only binding to the contract
	GasTokenTransactor // Write-only binding to the contract
	GasTokenFilterer   // Log filterer for contract events
}

// GasTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type GasTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GasTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GasTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GasTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GasTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GasTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GasTokenSession struct {
	Contract     *GasToken         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GasTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GasTokenCallerSession struct {
	Contract *GasTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// GasTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GasTokenTransactorSession struct {
	Contract     *GasTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// GasTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type GasTokenRaw struct {
	Contract *GasToken // Generic contract binding to access the raw methods on
}

// GasTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GasTokenCallerRaw struct {
	Contract *GasTokenCaller // Generic read-only contract binding to access the raw methods on
}

// GasTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GasTokenTransactorRaw struct {
	Contract *GasTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGasToken creates a new instance of GasToken, bound to a specific deployed contract.
func NewGasToken(address common.Address, backend bind.ContractBackend) (*GasToken, error) {
	contract, err := bindGasToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GasToken{GasTokenCaller: GasTokenCaller{contract: contract}, GasTokenTransactor: GasTokenTransactor{contract: contract}, GasTokenFilterer: GasTokenFilterer{contract: contract}}, nil
}

// NewGasTokenCaller creates a new read-only instance of GasToken, bound to a specific deployed contract.
func NewGasTokenCaller(address common.Address, caller bind.ContractCaller) (*GasTokenCaller, error) {
	contract, err := bindGasToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GasTokenCaller{contract: contract}, nil
}

// NewGasTokenTransactor creates a new write-only instance of GasToken, bound to a specific deployed contract.
func NewGasTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*GasTokenTransactor, error) {
	contract, err := bindGasToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GasTokenTransactor{contract: contract}, nil
}

// NewGasTokenFilterer creates a new log filterer instance of GasToken, bound to a specific deployed contract.
func NewGasTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*GasTokenFilterer, error) {
	contract, err := bindGasToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GasTokenFilterer{contract: contract}, nil
}

// bindGasToken binds a generic wrapper to an already deployed contract.
func bindGasToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GasTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GasToken *GasTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _GasToken.Contract.GasTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GasToken *GasTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GasToken.Contract.GasTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GasToken *GasTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GasToken.Contract.GasTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GasToken *GasTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _GasToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GasToken *GasTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GasToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GasToken *GasTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GasToken.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_GasToken *GasTokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _GasToken.contract.Call(opts, out, "balanceOf", account)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_GasToken *GasTokenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _GasToken.Contract.BalanceOf(&_GasToken.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_GasToken *GasTokenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _GasToken.Contract.BalanceOf(&_GasToken.CallOpts, account)
}

// TotalBurned is a free data retrieval call binding the contract method 0xd89135cd.
//
// Solidity: function totalBurned() constant returns(uint256)
func (_GasToken *GasTokenCaller) TotalBurned(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _GasToken.contract.Call(opts, out, "totalBurned")
	return *ret0, err
}

// TotalBurned is a free data retrieval call binding the contract method 0xd89135cd.
//
// Solidity: function totalBurned() constant returns(uint256)
func (_GasToken *GasTokenSession) TotalBurned() (*big.Int, error) {
	return _GasToken.Contract.TotalBurned(&_GasToken.CallOpts)
}

// TotalBurned is a free data retrieval call binding the contract method 0xd89135cd.
//
// Solidity: function totalBurned() constant returns(uint256)
func (_GasToken *GasTokenCallerSession) TotalBurned() (*big.Int, error) {
	return _GasToken.Contract.TotalBurned(&_GasToken.CallOpts)
}

// TotalMinted is a free data retrieval call binding the contract method 0xa2309ff8.
//
// Solidity: function totalMinted() constant returns(uint256)
func (_GasToken *GasTokenCaller) TotalMinted(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _GasToken.contract.Call(opts, out, "totalMinted")
	return *ret0, err
}

// TotalMinted is a free data retrieval call binding the contract method 0xa2309ff8.
//
// Solidity: function totalMinted() constant returns(uint256)
func (_GasToken *GasTokenSession) TotalMinted() (*big.Int, error) {
	return _GasToken.Contract.TotalMinted(&_GasToken.CallOpts)
}

// TotalMinted is a free data retrieval call binding the contract method 0xa2309ff8.
//
// Solidity: function totalMinted() constant returns(uint256)
func (_GasToken *GasTokenCallerSession) TotalMinted() (*big.Int, error) {
	return _GasToken.Contract.TotalMinted(&_GasToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_GasToken *GasTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _GasToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_GasToken *GasTokenSession) TotalSupply() (*big.Int, error) {
	return _GasToken.Contract.TotalSupply(&_GasToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_GasToken *GasTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _GasToken.Contract.TotalSupply(&_GasToken.CallOpts)
}

// Free is a paid mutator transaction binding the contract method 0xd8ccd0f3.
//
// Solidity: function free(uint256 value) returns(uint256)
func (_GasToken *GasTokenTransactor) Free(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _GasToken.contract.Transact(opts, "free", value)
}

// Free is a paid mutator transaction binding the contract method 0xd8ccd0f3.
//
// Solidity: function free(uint256 value) returns(uint256)
func (_GasToken *GasTokenSession) Free(value *big.Int) (*types.Transaction, error) {
	return _GasToken.Contract.Free(&_GasToken.TransactOpts, value)
}

// Free is a paid mutator transaction binding the contract method 0xd8ccd0f3.
//
// Solidity: function free(uint256 value) returns(uint256)
func (_GasToken *GasTokenTransactorSession) Free(value *big.Int) (*types.Transaction, error) {
	return _GasToken.Contract.Free(&_GasToken.TransactOpts, value)
}

// FreeUpTo is a paid mutator transaction binding the contract method 0x6366b936.
//
// Solidity: function freeUpTo(uint256 value) returns(uint256)
func (_GasToken *GasTokenTransactor) FreeUpTo(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _GasToken.contract.Transact(opts, "freeUpTo", value)
}

// FreeUpTo is a paid mutator transaction binding the contract method 0x6366b936.
//
// Solidity: function freeUpTo(uint256 value) returns(uint256)
func (_GasToken *GasTokenSession) FreeUpTo(value *big.Int) (*types.Transaction, error) {
	return _GasToken.Contract.FreeUpTo(&_GasToken.TransactOpts, value)
}

// FreeUpTo is a paid mutator transaction binding the contract method 0x6366b936.
//
// Solidity: function freeUpTo(uint256 value) returns(uint256)
func (_GasToken *GasTokenTransactorSession) FreeUpTo(value *big.Int) (*types.Transaction, error) {
	return _GasToken.Contract.FreeUpTo(&_GasToken.TransactOpts, value)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 value) returns()
func (_GasToken *GasTokenTransactor) Mint(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _GasToken.contract.Transact(opts, "mint", value)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 value) returns()
func (_GasToken *GasTokenSession) Mint(value *big.Int) (*types.Transaction, error) {
	return _GasToken.Contract.Mint(&_GasToken.TransactOpts, value)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 value) returns()
func (_GasToken *GasTokenTransactorSession) Mint(value *big.Int) (*types.Transaction, error) {
	return _GasToken.Contract.Mint(&_GasToken.TransactOpts, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns()
func (_GasToken *GasTokenTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _GasToken.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns()
func (_GasToken *GasTokenSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _GasToken.Contract.Transfer(&_GasToken.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns()
func (_GasToken *GasTokenTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _GasToken.Contract.Transfer(&_GasToken.TransactOpts, recipient, amount)
}
