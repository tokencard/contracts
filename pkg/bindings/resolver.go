// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

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

// ResolverABI is the input ABI used to generate the binding from.
const ResolverABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"getAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_controller\",\"type\":\"address\"}],\"name\":\"newController\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_target\",\"type\":\"address\"}],\"name\":\"setAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_target\",\"type\":\"address\"},{\"name\":\"_controller\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_old\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_new\",\"type\":\"address\"}],\"name\":\"ControllerResolverChanged\",\"type\":\"event\"}]"

// ResolverBin is the compiled bytecode used for deploying new contracts.
const ResolverBin = `608060405234801561001057600080fd5b5060405160408061039b83398101604052805160209091015160008054600160a060020a03928316600160a060020a03199182161790915560018054929093169116179055610337806100646000396000f3006080604052600436106100565763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166338cc4831811461005b578063ca57cb781461008c578063e30081a0146100af575b600080fd5b34801561006757600080fd5b506100706100d0565b60408051600160a060020a039092168252519081900360200190f35b34801561009857600080fd5b506100ad600160a060020a03600435166100df565b005b3480156100bb57600080fd5b506100ad600160a060020a03600435166101cc565b600154600160a060020a031690565b6100e833610271565b151561015557604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601a60248201527f73656e646572206973206e6f74206120636f6e74726f6c6c6572000000000000604482015290519081900360640190fd5b60005460408051600160a060020a039283168152918316602083015280517f95d5e19ddf857cd05ebf87d06d1bef8ceab83305a3f833ff72d6cfafbf76f8259281900390910190a16000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b6101d533610271565b151561024257604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601a60248201527f73656e646572206973206e6f74206120636f6e74726f6c6c6572000000000000604482015290519081900360640190fd5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b60008054604080517fb429afeb000000000000000000000000000000000000000000000000000000008152600160a060020a0385811660048301529151919092169163b429afeb91602480830192602092919082900301818787803b1580156102d957600080fd5b505af11580156102ed573d6000803e3d6000fd5b505050506040513d602081101561030357600080fd5b5051929150505600a165627a7a723058205d73ac1e00f965f2f4bc6e1546d3cc7e9acd96bc39fc3ba346eaedc7b604eac90029`

// DeployResolver deploys a new Ethereum contract, binding an instance of Resolver to it.
func DeployResolver(auth *bind.TransactOpts, backend bind.ContractBackend, _target common.Address, _controller common.Address) (common.Address, *types.Transaction, *Resolver, error) {
	parsed, err := abi.JSON(strings.NewReader(ResolverABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ResolverBin), backend, _target, _controller)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Resolver{ResolverCaller: ResolverCaller{contract: contract}, ResolverTransactor: ResolverTransactor{contract: contract}, ResolverFilterer: ResolverFilterer{contract: contract}}, nil
}

// Resolver is an auto generated Go binding around an Ethereum contract.
type Resolver struct {
	ResolverCaller     // Read-only binding to the contract
	ResolverTransactor // Write-only binding to the contract
	ResolverFilterer   // Log filterer for contract events
}

// ResolverCaller is an auto generated read-only Go binding around an Ethereum contract.
type ResolverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ResolverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ResolverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ResolverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ResolverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ResolverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ResolverSession struct {
	Contract     *Resolver         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ResolverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ResolverCallerSession struct {
	Contract *ResolverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ResolverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ResolverTransactorSession struct {
	Contract     *ResolverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ResolverRaw is an auto generated low-level Go binding around an Ethereum contract.
type ResolverRaw struct {
	Contract *Resolver // Generic contract binding to access the raw methods on
}

// ResolverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ResolverCallerRaw struct {
	Contract *ResolverCaller // Generic read-only contract binding to access the raw methods on
}

// ResolverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ResolverTransactorRaw struct {
	Contract *ResolverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewResolver creates a new instance of Resolver, bound to a specific deployed contract.
func NewResolver(address common.Address, backend bind.ContractBackend) (*Resolver, error) {
	contract, err := bindResolver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Resolver{ResolverCaller: ResolverCaller{contract: contract}, ResolverTransactor: ResolverTransactor{contract: contract}, ResolverFilterer: ResolverFilterer{contract: contract}}, nil
}

// NewResolverCaller creates a new read-only instance of Resolver, bound to a specific deployed contract.
func NewResolverCaller(address common.Address, caller bind.ContractCaller) (*ResolverCaller, error) {
	contract, err := bindResolver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ResolverCaller{contract: contract}, nil
}

// NewResolverTransactor creates a new write-only instance of Resolver, bound to a specific deployed contract.
func NewResolverTransactor(address common.Address, transactor bind.ContractTransactor) (*ResolverTransactor, error) {
	contract, err := bindResolver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ResolverTransactor{contract: contract}, nil
}

// NewResolverFilterer creates a new log filterer instance of Resolver, bound to a specific deployed contract.
func NewResolverFilterer(address common.Address, filterer bind.ContractFilterer) (*ResolverFilterer, error) {
	contract, err := bindResolver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ResolverFilterer{contract: contract}, nil
}

// bindResolver binds a generic wrapper to an already deployed contract.
func bindResolver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ResolverABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Resolver *ResolverRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Resolver.Contract.ResolverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Resolver *ResolverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Resolver.Contract.ResolverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Resolver *ResolverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Resolver.Contract.ResolverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Resolver *ResolverCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Resolver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Resolver *ResolverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Resolver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Resolver *ResolverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Resolver.Contract.contract.Transact(opts, method, params...)
}

// GetAddress is a free data retrieval call binding the contract method 0x38cc4831.
//
// Solidity: function getAddress() constant returns(address)
func (_Resolver *ResolverCaller) GetAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Resolver.contract.Call(opts, out, "getAddress")
	return *ret0, err
}

// GetAddress is a free data retrieval call binding the contract method 0x38cc4831.
//
// Solidity: function getAddress() constant returns(address)
func (_Resolver *ResolverSession) GetAddress() (common.Address, error) {
	return _Resolver.Contract.GetAddress(&_Resolver.CallOpts)
}

// GetAddress is a free data retrieval call binding the contract method 0x38cc4831.
//
// Solidity: function getAddress() constant returns(address)
func (_Resolver *ResolverCallerSession) GetAddress() (common.Address, error) {
	return _Resolver.Contract.GetAddress(&_Resolver.CallOpts)
}

// NewController is a paid mutator transaction binding the contract method 0xca57cb78.
//
// Solidity: function newController(_controller address) returns()
func (_Resolver *ResolverTransactor) NewController(opts *bind.TransactOpts, _controller common.Address) (*types.Transaction, error) {
	return _Resolver.contract.Transact(opts, "newController", _controller)
}

// NewController is a paid mutator transaction binding the contract method 0xca57cb78.
//
// Solidity: function newController(_controller address) returns()
func (_Resolver *ResolverSession) NewController(_controller common.Address) (*types.Transaction, error) {
	return _Resolver.Contract.NewController(&_Resolver.TransactOpts, _controller)
}

// NewController is a paid mutator transaction binding the contract method 0xca57cb78.
//
// Solidity: function newController(_controller address) returns()
func (_Resolver *ResolverTransactorSession) NewController(_controller common.Address) (*types.Transaction, error) {
	return _Resolver.Contract.NewController(&_Resolver.TransactOpts, _controller)
}

// SetAddress is a paid mutator transaction binding the contract method 0xe30081a0.
//
// Solidity: function setAddress(_target address) returns()
func (_Resolver *ResolverTransactor) SetAddress(opts *bind.TransactOpts, _target common.Address) (*types.Transaction, error) {
	return _Resolver.contract.Transact(opts, "setAddress", _target)
}

// SetAddress is a paid mutator transaction binding the contract method 0xe30081a0.
//
// Solidity: function setAddress(_target address) returns()
func (_Resolver *ResolverSession) SetAddress(_target common.Address) (*types.Transaction, error) {
	return _Resolver.Contract.SetAddress(&_Resolver.TransactOpts, _target)
}

// SetAddress is a paid mutator transaction binding the contract method 0xe30081a0.
//
// Solidity: function setAddress(_target address) returns()
func (_Resolver *ResolverTransactorSession) SetAddress(_target common.Address) (*types.Transaction, error) {
	return _Resolver.Contract.SetAddress(&_Resolver.TransactOpts, _target)
}

// ResolverControllerResolverChangedIterator is returned from FilterControllerResolverChanged and is used to iterate over the raw logs and unpacked data for ControllerResolverChanged events raised by the Resolver contract.
type ResolverControllerResolverChangedIterator struct {
	Event *ResolverControllerResolverChanged // Event containing the contract specifics and raw log

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
func (it *ResolverControllerResolverChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ResolverControllerResolverChanged)
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
		it.Event = new(ResolverControllerResolverChanged)
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
func (it *ResolverControllerResolverChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ResolverControllerResolverChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ResolverControllerResolverChanged represents a ControllerResolverChanged event raised by the Resolver contract.
type ResolverControllerResolverChanged struct {
	Old common.Address
	New common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterControllerResolverChanged is a free log retrieval operation binding the contract event 0x95d5e19ddf857cd05ebf87d06d1bef8ceab83305a3f833ff72d6cfafbf76f825.
//
// Solidity: e ControllerResolverChanged(_old address, _new address)
func (_Resolver *ResolverFilterer) FilterControllerResolverChanged(opts *bind.FilterOpts) (*ResolverControllerResolverChangedIterator, error) {

	logs, sub, err := _Resolver.contract.FilterLogs(opts, "ControllerResolverChanged")
	if err != nil {
		return nil, err
	}
	return &ResolverControllerResolverChangedIterator{contract: _Resolver.contract, event: "ControllerResolverChanged", logs: logs, sub: sub}, nil
}

// WatchControllerResolverChanged is a free log subscription operation binding the contract event 0x95d5e19ddf857cd05ebf87d06d1bef8ceab83305a3f833ff72d6cfafbf76f825.
//
// Solidity: e ControllerResolverChanged(_old address, _new address)
func (_Resolver *ResolverFilterer) WatchControllerResolverChanged(opts *bind.WatchOpts, sink chan<- *ResolverControllerResolverChanged) (event.Subscription, error) {

	logs, sub, err := _Resolver.contract.WatchLogs(opts, "ControllerResolverChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ResolverControllerResolverChanged)
				if err := _Resolver.contract.UnpackLog(event, "ControllerResolverChanged", log); err != nil {
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
