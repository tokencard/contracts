// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package upgradeability

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

// AdminUpgradeabilityProxyABI is the input ABI used to generate the binding from.
const AdminUpgradeabilityProxyABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_logic\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"constant\":true,\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"changeAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// AdminUpgradeabilityProxyBin is the compiled bytecode used for deploying new contracts.
var AdminUpgradeabilityProxyBin = "0x608060405260405161073f38038061073f8339818101604052606081101561002657600080fd5b8151602083015160408085018051915193959294830192918464010000000082111561005157600080fd5b90830190602082018581111561006657600080fd5b825164010000000081118282018810171561008057600080fd5b82525081516020918201929091019080838360005b838110156100ad578181015183820152602001610095565b50505050905090810190601f1680156100da5780820380516001836020036101000a031916815260200191505b5060408181527f656970313936372e70726f78792e696d706c656d656e746174696f6e0000000082525190819003601c0190208693508492506000805160206106e483398151915260001990910114905061013157fe5b610143826001600160e01b0361026516565b8051156101fb576000826001600160a01b0316826040518082805190602001908083835b602083106101865780518252601f199092019160209182019101610167565b6001836020036101000a038019825116818451168082178552505050505050905001915050600060405180830381855af49150503d80600081146101e6576040519150601f19603f3d011682016040523d82523d6000602084013e6101eb565b606091505b50509050806101f957600080fd5b505b5050604080517f656970313936372e70726f78792e61646d696e00000000000000000000000000815290519081900360130190206000805160206106c48339815191526000199091011461024b57fe5b61025d826001600160e01b036102c516565b5050506102dd565b610278816102d760201b61036e1760201c565b6102b35760405162461bcd60e51b815260040180806020018281038252603b815260200180610704603b913960400191505060405180910390fd5b6000805160206106e483398151915255565b6000805160206106c483398151915255565b3b151590565b6103d8806102ec6000396000f3fe6080604052600436106100345760003560e01c80635c60da1b1461003e5780638f2839701461006f578063f851a440146100a2575b61003c6100b7565b005b34801561004a57600080fd5b506100536100d1565b604080516001600160a01b039092168252519081900360200190f35b34801561007b57600080fd5b5061003c6004803603602081101561009257600080fd5b50356001600160a01b03166100e0565b3480156100ae57600080fd5b506100536102ca565b6100bf6102d4565b6100cf6100ca6102dc565b610301565b565b60006100db6102dc565b905090565b6100e8610325565b6001600160a01b0316336001600160a01b031614156102bf576001600160a01b0381166101465760405162461bcd60e51b815260040180806020018281038252602f815260200180610375602f913960400191505060405180910390fd5b600060606101526102dc565b6001600160a01b03166040518080632121dc7560e01b8152506004019050600060405180830381855af49150503d80600081146101ab576040519150601f19603f3d011682016040523d82523d6000602084013e6101b0565b606091505b5091509150816101fc576040805162461bcd60e51b815260206004820152601260248201527118da185b99d950591b5a5b8819985a5b195960721b604482015290519081900360640190fd5b80806020019051602081101561021157600080fd5b5051610264576040805162461bcd60e51b815260206004820152601c60248201527f6e6f6e2d7472616e7366657261626c652070726f78792061646d696e00000000604482015290519081900360640190fd5b7f7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f61028d610325565b604080516001600160a01b03928316815291861660208301528051918290030190a16102b88361034a565b50506102c7565b6102c76100b7565b50565b60006100db610325565b6100cf6100cf565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5490565b3660008037600080366000845af43d6000803e808015610320573d6000f35b3d6000fd5b7fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d61035490565b7fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d610355565b3b15159056fe43616e6e6f74206368616e6765207468652061646d696e206f6620612070726f787920746f20616464726573732030a265627a7a723158209677e7a767914536f878b2291af1bd414baa8e118b34229cffa2bbc7f8a4414464736f6c634300050f0032b53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d6103360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc43616e6e6f742073657420612070726f787920696d706c656d656e746174696f6e20746f2061206e6f6e2d636f6e74726163742061646472657373"

// DeployAdminUpgradeabilityProxy deploys a new Ethereum contract, binding an instance of AdminUpgradeabilityProxy to it.
func DeployAdminUpgradeabilityProxy(auth *bind.TransactOpts, backend bind.ContractBackend, _logic common.Address, _admin common.Address, _data []byte) (common.Address, *types.Transaction, *AdminUpgradeabilityProxy, error) {
	parsed, err := abi.JSON(strings.NewReader(AdminUpgradeabilityProxyABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AdminUpgradeabilityProxyBin), backend, _logic, _admin, _data)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AdminUpgradeabilityProxy{AdminUpgradeabilityProxyCaller: AdminUpgradeabilityProxyCaller{contract: contract}, AdminUpgradeabilityProxyTransactor: AdminUpgradeabilityProxyTransactor{contract: contract}, AdminUpgradeabilityProxyFilterer: AdminUpgradeabilityProxyFilterer{contract: contract}}, nil
}

// AdminUpgradeabilityProxy is an auto generated Go binding around an Ethereum contract.
type AdminUpgradeabilityProxy struct {
	AdminUpgradeabilityProxyCaller     // Read-only binding to the contract
	AdminUpgradeabilityProxyTransactor // Write-only binding to the contract
	AdminUpgradeabilityProxyFilterer   // Log filterer for contract events
}

// AdminUpgradeabilityProxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type AdminUpgradeabilityProxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AdminUpgradeabilityProxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AdminUpgradeabilityProxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AdminUpgradeabilityProxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AdminUpgradeabilityProxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AdminUpgradeabilityProxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AdminUpgradeabilityProxySession struct {
	Contract     *AdminUpgradeabilityProxy // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// AdminUpgradeabilityProxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AdminUpgradeabilityProxyCallerSession struct {
	Contract *AdminUpgradeabilityProxyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// AdminUpgradeabilityProxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AdminUpgradeabilityProxyTransactorSession struct {
	Contract     *AdminUpgradeabilityProxyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// AdminUpgradeabilityProxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type AdminUpgradeabilityProxyRaw struct {
	Contract *AdminUpgradeabilityProxy // Generic contract binding to access the raw methods on
}

// AdminUpgradeabilityProxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AdminUpgradeabilityProxyCallerRaw struct {
	Contract *AdminUpgradeabilityProxyCaller // Generic read-only contract binding to access the raw methods on
}

// AdminUpgradeabilityProxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AdminUpgradeabilityProxyTransactorRaw struct {
	Contract *AdminUpgradeabilityProxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAdminUpgradeabilityProxy creates a new instance of AdminUpgradeabilityProxy, bound to a specific deployed contract.
func NewAdminUpgradeabilityProxy(address common.Address, backend bind.ContractBackend) (*AdminUpgradeabilityProxy, error) {
	contract, err := bindAdminUpgradeabilityProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AdminUpgradeabilityProxy{AdminUpgradeabilityProxyCaller: AdminUpgradeabilityProxyCaller{contract: contract}, AdminUpgradeabilityProxyTransactor: AdminUpgradeabilityProxyTransactor{contract: contract}, AdminUpgradeabilityProxyFilterer: AdminUpgradeabilityProxyFilterer{contract: contract}}, nil
}

// NewAdminUpgradeabilityProxyCaller creates a new read-only instance of AdminUpgradeabilityProxy, bound to a specific deployed contract.
func NewAdminUpgradeabilityProxyCaller(address common.Address, caller bind.ContractCaller) (*AdminUpgradeabilityProxyCaller, error) {
	contract, err := bindAdminUpgradeabilityProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AdminUpgradeabilityProxyCaller{contract: contract}, nil
}

// NewAdminUpgradeabilityProxyTransactor creates a new write-only instance of AdminUpgradeabilityProxy, bound to a specific deployed contract.
func NewAdminUpgradeabilityProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*AdminUpgradeabilityProxyTransactor, error) {
	contract, err := bindAdminUpgradeabilityProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AdminUpgradeabilityProxyTransactor{contract: contract}, nil
}

// NewAdminUpgradeabilityProxyFilterer creates a new log filterer instance of AdminUpgradeabilityProxy, bound to a specific deployed contract.
func NewAdminUpgradeabilityProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*AdminUpgradeabilityProxyFilterer, error) {
	contract, err := bindAdminUpgradeabilityProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AdminUpgradeabilityProxyFilterer{contract: contract}, nil
}

// bindAdminUpgradeabilityProxy binds a generic wrapper to an already deployed contract.
func bindAdminUpgradeabilityProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AdminUpgradeabilityProxyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AdminUpgradeabilityProxy *AdminUpgradeabilityProxyRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AdminUpgradeabilityProxy.Contract.AdminUpgradeabilityProxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AdminUpgradeabilityProxy *AdminUpgradeabilityProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AdminUpgradeabilityProxy.Contract.AdminUpgradeabilityProxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AdminUpgradeabilityProxy *AdminUpgradeabilityProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AdminUpgradeabilityProxy.Contract.AdminUpgradeabilityProxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AdminUpgradeabilityProxy *AdminUpgradeabilityProxyCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AdminUpgradeabilityProxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AdminUpgradeabilityProxy *AdminUpgradeabilityProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AdminUpgradeabilityProxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AdminUpgradeabilityProxy *AdminUpgradeabilityProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AdminUpgradeabilityProxy.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_AdminUpgradeabilityProxy *AdminUpgradeabilityProxyCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _AdminUpgradeabilityProxy.contract.Call(opts, out, "admin")
	return *ret0, err
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_AdminUpgradeabilityProxy *AdminUpgradeabilityProxySession) Admin() (common.Address, error) {
	return _AdminUpgradeabilityProxy.Contract.Admin(&_AdminUpgradeabilityProxy.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_AdminUpgradeabilityProxy *AdminUpgradeabilityProxyCallerSession) Admin() (common.Address, error) {
	return _AdminUpgradeabilityProxy.Contract.Admin(&_AdminUpgradeabilityProxy.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() constant returns(address)
func (_AdminUpgradeabilityProxy *AdminUpgradeabilityProxyCaller) Implementation(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _AdminUpgradeabilityProxy.contract.Call(opts, out, "implementation")
	return *ret0, err
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() constant returns(address)
func (_AdminUpgradeabilityProxy *AdminUpgradeabilityProxySession) Implementation() (common.Address, error) {
	return _AdminUpgradeabilityProxy.Contract.Implementation(&_AdminUpgradeabilityProxy.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() constant returns(address)
func (_AdminUpgradeabilityProxy *AdminUpgradeabilityProxyCallerSession) Implementation() (common.Address, error) {
	return _AdminUpgradeabilityProxy.Contract.Implementation(&_AdminUpgradeabilityProxy.CallOpts)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address newAdmin) returns()
func (_AdminUpgradeabilityProxy *AdminUpgradeabilityProxyTransactor) ChangeAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _AdminUpgradeabilityProxy.contract.Transact(opts, "changeAdmin", newAdmin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address newAdmin) returns()
func (_AdminUpgradeabilityProxy *AdminUpgradeabilityProxySession) ChangeAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _AdminUpgradeabilityProxy.Contract.ChangeAdmin(&_AdminUpgradeabilityProxy.TransactOpts, newAdmin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address newAdmin) returns()
func (_AdminUpgradeabilityProxy *AdminUpgradeabilityProxyTransactorSession) ChangeAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _AdminUpgradeabilityProxy.Contract.ChangeAdmin(&_AdminUpgradeabilityProxy.TransactOpts, newAdmin)
}

// AdminUpgradeabilityProxyAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the AdminUpgradeabilityProxy contract.
type AdminUpgradeabilityProxyAdminChangedIterator struct {
	Event *AdminUpgradeabilityProxyAdminChanged // Event containing the contract specifics and raw log

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
func (it *AdminUpgradeabilityProxyAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AdminUpgradeabilityProxyAdminChanged)
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
		it.Event = new(AdminUpgradeabilityProxyAdminChanged)
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
func (it *AdminUpgradeabilityProxyAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AdminUpgradeabilityProxyAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AdminUpgradeabilityProxyAdminChanged represents a AdminChanged event raised by the AdminUpgradeabilityProxy contract.
type AdminUpgradeabilityProxyAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_AdminUpgradeabilityProxy *AdminUpgradeabilityProxyFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*AdminUpgradeabilityProxyAdminChangedIterator, error) {

	logs, sub, err := _AdminUpgradeabilityProxy.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &AdminUpgradeabilityProxyAdminChangedIterator{contract: _AdminUpgradeabilityProxy.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_AdminUpgradeabilityProxy *AdminUpgradeabilityProxyFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *AdminUpgradeabilityProxyAdminChanged) (event.Subscription, error) {

	logs, sub, err := _AdminUpgradeabilityProxy.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AdminUpgradeabilityProxyAdminChanged)
				if err := _AdminUpgradeabilityProxy.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_AdminUpgradeabilityProxy *AdminUpgradeabilityProxyFilterer) ParseAdminChanged(log types.Log) (*AdminUpgradeabilityProxyAdminChanged, error) {
	event := new(AdminUpgradeabilityProxyAdminChanged)
	if err := _AdminUpgradeabilityProxy.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AdminUpgradeabilityProxyUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the AdminUpgradeabilityProxy contract.
type AdminUpgradeabilityProxyUpgradedIterator struct {
	Event *AdminUpgradeabilityProxyUpgraded // Event containing the contract specifics and raw log

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
func (it *AdminUpgradeabilityProxyUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AdminUpgradeabilityProxyUpgraded)
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
		it.Event = new(AdminUpgradeabilityProxyUpgraded)
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
func (it *AdminUpgradeabilityProxyUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AdminUpgradeabilityProxyUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AdminUpgradeabilityProxyUpgraded represents a Upgraded event raised by the AdminUpgradeabilityProxy contract.
type AdminUpgradeabilityProxyUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_AdminUpgradeabilityProxy *AdminUpgradeabilityProxyFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*AdminUpgradeabilityProxyUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _AdminUpgradeabilityProxy.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &AdminUpgradeabilityProxyUpgradedIterator{contract: _AdminUpgradeabilityProxy.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_AdminUpgradeabilityProxy *AdminUpgradeabilityProxyFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *AdminUpgradeabilityProxyUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _AdminUpgradeabilityProxy.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AdminUpgradeabilityProxyUpgraded)
				if err := _AdminUpgradeabilityProxy.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_AdminUpgradeabilityProxy *AdminUpgradeabilityProxyFilterer) ParseUpgraded(log types.Log) (*AdminUpgradeabilityProxyUpgraded, error) {
	event := new(AdminUpgradeabilityProxyUpgraded)
	if err := _AdminUpgradeabilityProxy.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	return event, nil
}
