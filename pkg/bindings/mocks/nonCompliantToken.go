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

// NonCompliantTokenABI is the input ABI used to generate the binding from.
const NonCompliantTokenABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"credit\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"debit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// NonCompliantTokenBin is the compiled bytecode used for deploying new contracts.
var NonCompliantTokenBin = "0x608060405234801561001057600080fd5b50610693806100206000396000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c806370a082311161005b57806370a0823114610137578063a9059cbb1461015d578063dd62ed3e14610189578063ef6506db146101b757610088565b8063095ea7b31461008d57806318160ddd146100bb57806323b872dd146100d5578063636be27a1461010b575b600080fd5b6100b9600480360360408110156100a357600080fd5b506001600160a01b0381351690602001356101f7565b005b6100c3610294565b60408051918252519081900360200190f35b6100b9600480360360608110156100eb57600080fd5b506001600160a01b0381358116916020810135909116906040013561029a565b6100b96004803603604081101561012157600080fd5b506001600160a01b0381351690602001356103df565b6100c36004803603602081101561014d57600080fd5b50356001600160a01b0316610409565b6100b96004803603604081101561017357600080fd5b506001600160a01b03813516906020013561041b565b6100c36004803603604081101561019f57600080fd5b506001600160a01b03813581169160200135166104df565b6101e3600480360360408110156101cd57600080fd5b506001600160a01b0381351690602001356104fc565b604080519115158252519081900360200190f35b801580159061022857503360009081526002602090815260408083206001600160a01b038616845290915290205415155b1561023257600080fd5b3360008181526002602090815260408083206001600160a01b03871680855290835292819020859055805185815290519293927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929181900390910190a35050565b60005481565b6001600160a01b0382166102ad57600080fd5b6001600160a01b0383166000908152600160205260409020548111156102d257600080fd5b6001600160a01b03831660009081526002602090815260408083203384529091529020548181101561030357600080fd5b6001600160a01b0383166000908152600160205260409020546103269083610528565b6001600160a01b0380851660009081526001602052604080822093909355908616815220546103559083610589565b6001600160a01b0385166000908152600160205260409020556103788183610589565b6001600160a01b03808616600081815260026020908152604080832033845282529182902094909455805186815290519287169391927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef929181900390910190a350505050565b6001600160a01b039091166000908152600160205260408120805483900390558054919091039055565b60016020526000908152604090205481565b33600090815260016020526040902054811115610476576040805162461bcd60e51b8152602060048201526014602482015273696e73756666696369656e742062616c616e636560601b604482015290519081900360640190fd5b336000818152600160209081526040808320805486900390556001600160a01b03861680845292819020805486019055805185815290519293927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef929181900390910190a35050565b600260209081526000928352604080842090915290825290205481565b6001600160a01b0391909116600090815260016020819052604082208054840190558154909201905590565b600082820183811015610582576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b9392505050565b600061058283836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f770000815250600081848411156106555760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561061a578181015183820152602001610602565b50505050905090810190601f1680156106475780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b50505090039056fea264697066735822122024c7544145647512e08559464ad9ec71d20c2b42b9336582709768bc2eb8ec1164736f6c634300060b0033"

// DeployNonCompliantToken deploys a new Ethereum contract, binding an instance of NonCompliantToken to it.
func DeployNonCompliantToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *NonCompliantToken, error) {
	parsed, err := abi.JSON(strings.NewReader(NonCompliantTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(NonCompliantTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &NonCompliantToken{NonCompliantTokenCaller: NonCompliantTokenCaller{contract: contract}, NonCompliantTokenTransactor: NonCompliantTokenTransactor{contract: contract}, NonCompliantTokenFilterer: NonCompliantTokenFilterer{contract: contract}}, nil
}

// NonCompliantToken is an auto generated Go binding around an Ethereum contract.
type NonCompliantToken struct {
	NonCompliantTokenCaller     // Read-only binding to the contract
	NonCompliantTokenTransactor // Write-only binding to the contract
	NonCompliantTokenFilterer   // Log filterer for contract events
}

// NonCompliantTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type NonCompliantTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NonCompliantTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NonCompliantTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NonCompliantTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NonCompliantTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NonCompliantTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NonCompliantTokenSession struct {
	Contract     *NonCompliantToken // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// NonCompliantTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NonCompliantTokenCallerSession struct {
	Contract *NonCompliantTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// NonCompliantTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NonCompliantTokenTransactorSession struct {
	Contract     *NonCompliantTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// NonCompliantTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type NonCompliantTokenRaw struct {
	Contract *NonCompliantToken // Generic contract binding to access the raw methods on
}

// NonCompliantTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NonCompliantTokenCallerRaw struct {
	Contract *NonCompliantTokenCaller // Generic read-only contract binding to access the raw methods on
}

// NonCompliantTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NonCompliantTokenTransactorRaw struct {
	Contract *NonCompliantTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNonCompliantToken creates a new instance of NonCompliantToken, bound to a specific deployed contract.
func NewNonCompliantToken(address common.Address, backend bind.ContractBackend) (*NonCompliantToken, error) {
	contract, err := bindNonCompliantToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NonCompliantToken{NonCompliantTokenCaller: NonCompliantTokenCaller{contract: contract}, NonCompliantTokenTransactor: NonCompliantTokenTransactor{contract: contract}, NonCompliantTokenFilterer: NonCompliantTokenFilterer{contract: contract}}, nil
}

// NewNonCompliantTokenCaller creates a new read-only instance of NonCompliantToken, bound to a specific deployed contract.
func NewNonCompliantTokenCaller(address common.Address, caller bind.ContractCaller) (*NonCompliantTokenCaller, error) {
	contract, err := bindNonCompliantToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NonCompliantTokenCaller{contract: contract}, nil
}

// NewNonCompliantTokenTransactor creates a new write-only instance of NonCompliantToken, bound to a specific deployed contract.
func NewNonCompliantTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*NonCompliantTokenTransactor, error) {
	contract, err := bindNonCompliantToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NonCompliantTokenTransactor{contract: contract}, nil
}

// NewNonCompliantTokenFilterer creates a new log filterer instance of NonCompliantToken, bound to a specific deployed contract.
func NewNonCompliantTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*NonCompliantTokenFilterer, error) {
	contract, err := bindNonCompliantToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NonCompliantTokenFilterer{contract: contract}, nil
}

// bindNonCompliantToken binds a generic wrapper to an already deployed contract.
func bindNonCompliantToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NonCompliantTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NonCompliantToken *NonCompliantTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _NonCompliantToken.Contract.NonCompliantTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NonCompliantToken *NonCompliantTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NonCompliantToken.Contract.NonCompliantTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NonCompliantToken *NonCompliantTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NonCompliantToken.Contract.NonCompliantTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NonCompliantToken *NonCompliantTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _NonCompliantToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NonCompliantToken *NonCompliantTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NonCompliantToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NonCompliantToken *NonCompliantTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NonCompliantToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) constant returns(uint256)
func (_NonCompliantToken *NonCompliantTokenCaller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _NonCompliantToken.contract.Call(opts, out, "allowance", arg0, arg1)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) constant returns(uint256)
func (_NonCompliantToken *NonCompliantTokenSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _NonCompliantToken.Contract.Allowance(&_NonCompliantToken.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) constant returns(uint256)
func (_NonCompliantToken *NonCompliantTokenCallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _NonCompliantToken.Contract.Allowance(&_NonCompliantToken.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) constant returns(uint256)
func (_NonCompliantToken *NonCompliantTokenCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _NonCompliantToken.contract.Call(opts, out, "balanceOf", arg0)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) constant returns(uint256)
func (_NonCompliantToken *NonCompliantTokenSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _NonCompliantToken.Contract.BalanceOf(&_NonCompliantToken.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) constant returns(uint256)
func (_NonCompliantToken *NonCompliantTokenCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _NonCompliantToken.Contract.BalanceOf(&_NonCompliantToken.CallOpts, arg0)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_NonCompliantToken *NonCompliantTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _NonCompliantToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_NonCompliantToken *NonCompliantTokenSession) TotalSupply() (*big.Int, error) {
	return _NonCompliantToken.Contract.TotalSupply(&_NonCompliantToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_NonCompliantToken *NonCompliantTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _NonCompliantToken.Contract.TotalSupply(&_NonCompliantToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns()
func (_NonCompliantToken *NonCompliantTokenTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _NonCompliantToken.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns()
func (_NonCompliantToken *NonCompliantTokenSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _NonCompliantToken.Contract.Approve(&_NonCompliantToken.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns()
func (_NonCompliantToken *NonCompliantTokenTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _NonCompliantToken.Contract.Approve(&_NonCompliantToken.TransactOpts, _spender, _value)
}

// Credit is a paid mutator transaction binding the contract method 0xef6506db.
//
// Solidity: function credit(address to, uint256 amount) returns(bool)
func (_NonCompliantToken *NonCompliantTokenTransactor) Credit(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NonCompliantToken.contract.Transact(opts, "credit", to, amount)
}

// Credit is a paid mutator transaction binding the contract method 0xef6506db.
//
// Solidity: function credit(address to, uint256 amount) returns(bool)
func (_NonCompliantToken *NonCompliantTokenSession) Credit(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NonCompliantToken.Contract.Credit(&_NonCompliantToken.TransactOpts, to, amount)
}

// Credit is a paid mutator transaction binding the contract method 0xef6506db.
//
// Solidity: function credit(address to, uint256 amount) returns(bool)
func (_NonCompliantToken *NonCompliantTokenTransactorSession) Credit(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NonCompliantToken.Contract.Credit(&_NonCompliantToken.TransactOpts, to, amount)
}

// Debit is a paid mutator transaction binding the contract method 0x636be27a.
//
// Solidity: function debit(address from, uint256 amount) returns()
func (_NonCompliantToken *NonCompliantTokenTransactor) Debit(opts *bind.TransactOpts, from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NonCompliantToken.contract.Transact(opts, "debit", from, amount)
}

// Debit is a paid mutator transaction binding the contract method 0x636be27a.
//
// Solidity: function debit(address from, uint256 amount) returns()
func (_NonCompliantToken *NonCompliantTokenSession) Debit(from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NonCompliantToken.Contract.Debit(&_NonCompliantToken.TransactOpts, from, amount)
}

// Debit is a paid mutator transaction binding the contract method 0x636be27a.
//
// Solidity: function debit(address from, uint256 amount) returns()
func (_NonCompliantToken *NonCompliantTokenTransactorSession) Debit(from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NonCompliantToken.Contract.Debit(&_NonCompliantToken.TransactOpts, from, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns()
func (_NonCompliantToken *NonCompliantTokenTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NonCompliantToken.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns()
func (_NonCompliantToken *NonCompliantTokenSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NonCompliantToken.Contract.Transfer(&_NonCompliantToken.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns()
func (_NonCompliantToken *NonCompliantTokenTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NonCompliantToken.Contract.Transfer(&_NonCompliantToken.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns()
func (_NonCompliantToken *NonCompliantTokenTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _NonCompliantToken.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns()
func (_NonCompliantToken *NonCompliantTokenSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _NonCompliantToken.Contract.TransferFrom(&_NonCompliantToken.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns()
func (_NonCompliantToken *NonCompliantTokenTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _NonCompliantToken.Contract.TransferFrom(&_NonCompliantToken.TransactOpts, _from, _to, _value)
}

// NonCompliantTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the NonCompliantToken contract.
type NonCompliantTokenApprovalIterator struct {
	Event *NonCompliantTokenApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NonCompliantTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NonCompliantTokenApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NonCompliantTokenApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NonCompliantTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NonCompliantTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NonCompliantTokenApproval represents a Approval event raised by the NonCompliantToken contract.
type NonCompliantTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_NonCompliantToken *NonCompliantTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*NonCompliantTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _NonCompliantToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &NonCompliantTokenApprovalIterator{contract: _NonCompliantToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_NonCompliantToken *NonCompliantTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *NonCompliantTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _NonCompliantToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NonCompliantTokenApproval)
				if err := _NonCompliantToken.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_NonCompliantToken *NonCompliantTokenFilterer) ParseApproval(log types.Log) (*NonCompliantTokenApproval, error) {
	event := new(NonCompliantTokenApproval)
	if err := _NonCompliantToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// NonCompliantTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the NonCompliantToken contract.
type NonCompliantTokenTransferIterator struct {
	Event *NonCompliantTokenTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NonCompliantTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NonCompliantTokenTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NonCompliantTokenTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NonCompliantTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NonCompliantTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NonCompliantTokenTransfer represents a Transfer event raised by the NonCompliantToken contract.
type NonCompliantTokenTransfer struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_NonCompliantToken *NonCompliantTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*NonCompliantTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _NonCompliantToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &NonCompliantTokenTransferIterator{contract: _NonCompliantToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_NonCompliantToken *NonCompliantTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *NonCompliantTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _NonCompliantToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NonCompliantTokenTransfer)
				if err := _NonCompliantToken.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_NonCompliantToken *NonCompliantTokenFilterer) ParseTransfer(log types.Log) (*NonCompliantTokenTransfer, error) {
	event := new(NonCompliantTokenTransfer)
	if err := _NonCompliantToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}
