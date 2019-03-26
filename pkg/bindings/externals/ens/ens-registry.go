// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ens

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

// ENSRegistryABI is the input ABI used to generate the binding from.
const ENSRegistryABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"resolver\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"},{\"name\":\"label\",\"type\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"setSubnodeOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"},{\"name\":\"ttl\",\"type\":\"uint64\"}],\"name\":\"setTTL\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"ttl\",\"outputs\":[{\"name\":\"\",\"type\":\"uint64\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"},{\"name\":\"resolver\",\"type\":\"address\"}],\"name\":\"setResolver\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"label\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"NewOwner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"resolver\",\"type\":\"address\"}],\"name\":\"NewResolver\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"ttl\",\"type\":\"uint64\"}],\"name\":\"NewTTL\",\"type\":\"event\"}]"

// ENSRegistryBin is the compiled bytecode used for deploying new contracts.
const ENSRegistryBin = `608060405234801561001057600080fd5b5060008080526020527fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb58054600160a060020a03191633179055610562806100596000396000f3006080604052600436106100825763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416630178b8bf811461008757806302571be3146100bb57806306ab5923146100d357806314ab9038146100fc57806316a25cbd146101215780631896f70a146101565780635b0fc9c31461017a575b600080fd5b34801561009357600080fd5b5061009f60043561019e565b60408051600160a060020a039092168252519081900360200190f35b3480156100c757600080fd5b5061009f6004356101bc565b3480156100df57600080fd5b506100fa600435602435600160a060020a03604435166101d7565b005b34801561010857600080fd5b506100fa60043567ffffffffffffffff602435166102f3565b34801561012d57600080fd5b506101396004356103bc565b6040805167ffffffffffffffff9092168252519081900360200190f35b34801561016257600080fd5b506100fa600435600160a060020a03602435166103f3565b34801561018657600080fd5b506100fa600435600160a060020a0360243516610496565b600090815260208190526040902060010154600160a060020a031690565b600090815260208190526040902054600160a060020a031690565b6000838152602081905260408120548490600160a060020a031633146101fc57600080fd5b604080516020808201889052818301879052825180830384018152606090920192839052815191929182918401908083835b6020831061024d5780518252601f19909201916020918201910161022e565b51815160209384036101000a600019018019909216911617905260408051929094018290038220600160a060020a038a16835293519397508995508a94507fce0457fe73731f824cc272376169235128c118b49d344817417c6d108d155e829391829003019150a3506000908152602081905260409020805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03929092169190911790555050565b6000828152602081905260409020548290600160a060020a0316331461031857600080fd5b6040805167ffffffffffffffff84168152905184917f1d4f9bbfc9cab89d66e1a1562f2233ccbf1308cb4f63de2ead5787adddb8fa68919081900360200190a250600091825260208290526040909120600101805467ffffffffffffffff90921674010000000000000000000000000000000000000000027fffffffff0000000000000000ffffffffffffffffffffffffffffffffffffffff909216919091179055565b60009081526020819052604090206001015474010000000000000000000000000000000000000000900467ffffffffffffffff1690565b6000828152602081905260409020548290600160a060020a0316331461041857600080fd5b60408051600160a060020a0384168152905184917f335721b01866dc23fbee8b6b2c7b1e14d6f05c28cd35a2c934239f94095602a0919081900360200190a250600091825260208290526040909120600101805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03909216919091179055565b6000828152602081905260409020548290600160a060020a031633146104bb57600080fd5b60408051600160a060020a0384168152905184917fd4735d920b0f87494915f556dd9b54c8f309026070caea5c737245152564d266919081900360200190a250600091825260208290526040909120805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a039092169190911790555600a165627a7a72305820d985bfb1e8b3f0b3bc012998e7025377a5f0b7a5c40fdb057c29212cb7b1fd150029`

// DeployENSRegistry deploys a new Ethereum contract, binding an instance of ENSRegistry to it.
func DeployENSRegistry(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ENSRegistry, error) {
	parsed, err := abi.JSON(strings.NewReader(ENSRegistryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ENSRegistryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ENSRegistry{ENSRegistryCaller: ENSRegistryCaller{contract: contract}, ENSRegistryTransactor: ENSRegistryTransactor{contract: contract}, ENSRegistryFilterer: ENSRegistryFilterer{contract: contract}}, nil
}

// ENSRegistry is an auto generated Go binding around an Ethereum contract.
type ENSRegistry struct {
	ENSRegistryCaller     // Read-only binding to the contract
	ENSRegistryTransactor // Write-only binding to the contract
	ENSRegistryFilterer   // Log filterer for contract events
}

// ENSRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ENSRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ENSRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ENSRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ENSRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ENSRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ENSRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ENSRegistrySession struct {
	Contract     *ENSRegistry      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ENSRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ENSRegistryCallerSession struct {
	Contract *ENSRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// ENSRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ENSRegistryTransactorSession struct {
	Contract     *ENSRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ENSRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ENSRegistryRaw struct {
	Contract *ENSRegistry // Generic contract binding to access the raw methods on
}

// ENSRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ENSRegistryCallerRaw struct {
	Contract *ENSRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// ENSRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ENSRegistryTransactorRaw struct {
	Contract *ENSRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewENSRegistry creates a new instance of ENSRegistry, bound to a specific deployed contract.
func NewENSRegistry(address common.Address, backend bind.ContractBackend) (*ENSRegistry, error) {
	contract, err := bindENSRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ENSRegistry{ENSRegistryCaller: ENSRegistryCaller{contract: contract}, ENSRegistryTransactor: ENSRegistryTransactor{contract: contract}, ENSRegistryFilterer: ENSRegistryFilterer{contract: contract}}, nil
}

// NewENSRegistryCaller creates a new read-only instance of ENSRegistry, bound to a specific deployed contract.
func NewENSRegistryCaller(address common.Address, caller bind.ContractCaller) (*ENSRegistryCaller, error) {
	contract, err := bindENSRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ENSRegistryCaller{contract: contract}, nil
}

// NewENSRegistryTransactor creates a new write-only instance of ENSRegistry, bound to a specific deployed contract.
func NewENSRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*ENSRegistryTransactor, error) {
	contract, err := bindENSRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ENSRegistryTransactor{contract: contract}, nil
}

// NewENSRegistryFilterer creates a new log filterer instance of ENSRegistry, bound to a specific deployed contract.
func NewENSRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*ENSRegistryFilterer, error) {
	contract, err := bindENSRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ENSRegistryFilterer{contract: contract}, nil
}

// bindENSRegistry binds a generic wrapper to an already deployed contract.
func bindENSRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ENSRegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ENSRegistry *ENSRegistryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ENSRegistry.Contract.ENSRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ENSRegistry *ENSRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ENSRegistry.Contract.ENSRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ENSRegistry *ENSRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ENSRegistry.Contract.ENSRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ENSRegistry *ENSRegistryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ENSRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ENSRegistry *ENSRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ENSRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ENSRegistry *ENSRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ENSRegistry.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x02571be3.
//
// Solidity: function owner(bytes32 node) constant returns(address)
func (_ENSRegistry *ENSRegistryCaller) Owner(opts *bind.CallOpts, node [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ENSRegistry.contract.Call(opts, out, "owner", node)
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x02571be3.
//
// Solidity: function owner(bytes32 node) constant returns(address)
func (_ENSRegistry *ENSRegistrySession) Owner(node [32]byte) (common.Address, error) {
	return _ENSRegistry.Contract.Owner(&_ENSRegistry.CallOpts, node)
}

// Owner is a free data retrieval call binding the contract method 0x02571be3.
//
// Solidity: function owner(bytes32 node) constant returns(address)
func (_ENSRegistry *ENSRegistryCallerSession) Owner(node [32]byte) (common.Address, error) {
	return _ENSRegistry.Contract.Owner(&_ENSRegistry.CallOpts, node)
}

// Resolver is a free data retrieval call binding the contract method 0x0178b8bf.
//
// Solidity: function resolver(bytes32 node) constant returns(address)
func (_ENSRegistry *ENSRegistryCaller) Resolver(opts *bind.CallOpts, node [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ENSRegistry.contract.Call(opts, out, "resolver", node)
	return *ret0, err
}

// Resolver is a free data retrieval call binding the contract method 0x0178b8bf.
//
// Solidity: function resolver(bytes32 node) constant returns(address)
func (_ENSRegistry *ENSRegistrySession) Resolver(node [32]byte) (common.Address, error) {
	return _ENSRegistry.Contract.Resolver(&_ENSRegistry.CallOpts, node)
}

// Resolver is a free data retrieval call binding the contract method 0x0178b8bf.
//
// Solidity: function resolver(bytes32 node) constant returns(address)
func (_ENSRegistry *ENSRegistryCallerSession) Resolver(node [32]byte) (common.Address, error) {
	return _ENSRegistry.Contract.Resolver(&_ENSRegistry.CallOpts, node)
}

// Ttl is a free data retrieval call binding the contract method 0x16a25cbd.
//
// Solidity: function ttl(bytes32 node) constant returns(uint64)
func (_ENSRegistry *ENSRegistryCaller) Ttl(opts *bind.CallOpts, node [32]byte) (uint64, error) {
	var (
		ret0 = new(uint64)
	)
	out := ret0
	err := _ENSRegistry.contract.Call(opts, out, "ttl", node)
	return *ret0, err
}

// Ttl is a free data retrieval call binding the contract method 0x16a25cbd.
//
// Solidity: function ttl(bytes32 node) constant returns(uint64)
func (_ENSRegistry *ENSRegistrySession) Ttl(node [32]byte) (uint64, error) {
	return _ENSRegistry.Contract.Ttl(&_ENSRegistry.CallOpts, node)
}

// Ttl is a free data retrieval call binding the contract method 0x16a25cbd.
//
// Solidity: function ttl(bytes32 node) constant returns(uint64)
func (_ENSRegistry *ENSRegistryCallerSession) Ttl(node [32]byte) (uint64, error) {
	return _ENSRegistry.Contract.Ttl(&_ENSRegistry.CallOpts, node)
}

// SetOwner is a paid mutator transaction binding the contract method 0x5b0fc9c3.
//
// Solidity: function setOwner(bytes32 node, address owner) returns()
func (_ENSRegistry *ENSRegistryTransactor) SetOwner(opts *bind.TransactOpts, node [32]byte, owner common.Address) (*types.Transaction, error) {
	return _ENSRegistry.contract.Transact(opts, "setOwner", node, owner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x5b0fc9c3.
//
// Solidity: function setOwner(bytes32 node, address owner) returns()
func (_ENSRegistry *ENSRegistrySession) SetOwner(node [32]byte, owner common.Address) (*types.Transaction, error) {
	return _ENSRegistry.Contract.SetOwner(&_ENSRegistry.TransactOpts, node, owner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x5b0fc9c3.
//
// Solidity: function setOwner(bytes32 node, address owner) returns()
func (_ENSRegistry *ENSRegistryTransactorSession) SetOwner(node [32]byte, owner common.Address) (*types.Transaction, error) {
	return _ENSRegistry.Contract.SetOwner(&_ENSRegistry.TransactOpts, node, owner)
}

// SetResolver is a paid mutator transaction binding the contract method 0x1896f70a.
//
// Solidity: function setResolver(bytes32 node, address resolver) returns()
func (_ENSRegistry *ENSRegistryTransactor) SetResolver(opts *bind.TransactOpts, node [32]byte, resolver common.Address) (*types.Transaction, error) {
	return _ENSRegistry.contract.Transact(opts, "setResolver", node, resolver)
}

// SetResolver is a paid mutator transaction binding the contract method 0x1896f70a.
//
// Solidity: function setResolver(bytes32 node, address resolver) returns()
func (_ENSRegistry *ENSRegistrySession) SetResolver(node [32]byte, resolver common.Address) (*types.Transaction, error) {
	return _ENSRegistry.Contract.SetResolver(&_ENSRegistry.TransactOpts, node, resolver)
}

// SetResolver is a paid mutator transaction binding the contract method 0x1896f70a.
//
// Solidity: function setResolver(bytes32 node, address resolver) returns()
func (_ENSRegistry *ENSRegistryTransactorSession) SetResolver(node [32]byte, resolver common.Address) (*types.Transaction, error) {
	return _ENSRegistry.Contract.SetResolver(&_ENSRegistry.TransactOpts, node, resolver)
}

// SetSubnodeOwner is a paid mutator transaction binding the contract method 0x06ab5923.
//
// Solidity: function setSubnodeOwner(bytes32 node, bytes32 label, address owner) returns()
func (_ENSRegistry *ENSRegistryTransactor) SetSubnodeOwner(opts *bind.TransactOpts, node [32]byte, label [32]byte, owner common.Address) (*types.Transaction, error) {
	return _ENSRegistry.contract.Transact(opts, "setSubnodeOwner", node, label, owner)
}

// SetSubnodeOwner is a paid mutator transaction binding the contract method 0x06ab5923.
//
// Solidity: function setSubnodeOwner(bytes32 node, bytes32 label, address owner) returns()
func (_ENSRegistry *ENSRegistrySession) SetSubnodeOwner(node [32]byte, label [32]byte, owner common.Address) (*types.Transaction, error) {
	return _ENSRegistry.Contract.SetSubnodeOwner(&_ENSRegistry.TransactOpts, node, label, owner)
}

// SetSubnodeOwner is a paid mutator transaction binding the contract method 0x06ab5923.
//
// Solidity: function setSubnodeOwner(bytes32 node, bytes32 label, address owner) returns()
func (_ENSRegistry *ENSRegistryTransactorSession) SetSubnodeOwner(node [32]byte, label [32]byte, owner common.Address) (*types.Transaction, error) {
	return _ENSRegistry.Contract.SetSubnodeOwner(&_ENSRegistry.TransactOpts, node, label, owner)
}

// SetTTL is a paid mutator transaction binding the contract method 0x14ab9038.
//
// Solidity: function setTTL(bytes32 node, uint64 ttl) returns()
func (_ENSRegistry *ENSRegistryTransactor) SetTTL(opts *bind.TransactOpts, node [32]byte, ttl uint64) (*types.Transaction, error) {
	return _ENSRegistry.contract.Transact(opts, "setTTL", node, ttl)
}

// SetTTL is a paid mutator transaction binding the contract method 0x14ab9038.
//
// Solidity: function setTTL(bytes32 node, uint64 ttl) returns()
func (_ENSRegistry *ENSRegistrySession) SetTTL(node [32]byte, ttl uint64) (*types.Transaction, error) {
	return _ENSRegistry.Contract.SetTTL(&_ENSRegistry.TransactOpts, node, ttl)
}

// SetTTL is a paid mutator transaction binding the contract method 0x14ab9038.
//
// Solidity: function setTTL(bytes32 node, uint64 ttl) returns()
func (_ENSRegistry *ENSRegistryTransactorSession) SetTTL(node [32]byte, ttl uint64) (*types.Transaction, error) {
	return _ENSRegistry.Contract.SetTTL(&_ENSRegistry.TransactOpts, node, ttl)
}

// ENSRegistryNewOwnerIterator is returned from FilterNewOwner and is used to iterate over the raw logs and unpacked data for NewOwner events raised by the ENSRegistry contract.
type ENSRegistryNewOwnerIterator struct {
	Event *ENSRegistryNewOwner // Event containing the contract specifics and raw log

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
func (it *ENSRegistryNewOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ENSRegistryNewOwner)
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
		it.Event = new(ENSRegistryNewOwner)
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
func (it *ENSRegistryNewOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ENSRegistryNewOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ENSRegistryNewOwner represents a NewOwner event raised by the ENSRegistry contract.
type ENSRegistryNewOwner struct {
	Node  [32]byte
	Label [32]byte
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterNewOwner is a free log retrieval operation binding the contract event 0xce0457fe73731f824cc272376169235128c118b49d344817417c6d108d155e82.
//
// Solidity: event NewOwner(bytes32 indexed node, bytes32 indexed label, address owner)
func (_ENSRegistry *ENSRegistryFilterer) FilterNewOwner(opts *bind.FilterOpts, node [][32]byte, label [][32]byte) (*ENSRegistryNewOwnerIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}
	var labelRule []interface{}
	for _, labelItem := range label {
		labelRule = append(labelRule, labelItem)
	}

	logs, sub, err := _ENSRegistry.contract.FilterLogs(opts, "NewOwner", nodeRule, labelRule)
	if err != nil {
		return nil, err
	}
	return &ENSRegistryNewOwnerIterator{contract: _ENSRegistry.contract, event: "NewOwner", logs: logs, sub: sub}, nil
}

// WatchNewOwner is a free log subscription operation binding the contract event 0xce0457fe73731f824cc272376169235128c118b49d344817417c6d108d155e82.
//
// Solidity: event NewOwner(bytes32 indexed node, bytes32 indexed label, address owner)
func (_ENSRegistry *ENSRegistryFilterer) WatchNewOwner(opts *bind.WatchOpts, sink chan<- *ENSRegistryNewOwner, node [][32]byte, label [][32]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}
	var labelRule []interface{}
	for _, labelItem := range label {
		labelRule = append(labelRule, labelItem)
	}

	logs, sub, err := _ENSRegistry.contract.WatchLogs(opts, "NewOwner", nodeRule, labelRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ENSRegistryNewOwner)
				if err := _ENSRegistry.contract.UnpackLog(event, "NewOwner", log); err != nil {
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

// ENSRegistryNewResolverIterator is returned from FilterNewResolver and is used to iterate over the raw logs and unpacked data for NewResolver events raised by the ENSRegistry contract.
type ENSRegistryNewResolverIterator struct {
	Event *ENSRegistryNewResolver // Event containing the contract specifics and raw log

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
func (it *ENSRegistryNewResolverIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ENSRegistryNewResolver)
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
		it.Event = new(ENSRegistryNewResolver)
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
func (it *ENSRegistryNewResolverIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ENSRegistryNewResolverIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ENSRegistryNewResolver represents a NewResolver event raised by the ENSRegistry contract.
type ENSRegistryNewResolver struct {
	Node     [32]byte
	Resolver common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewResolver is a free log retrieval operation binding the contract event 0x335721b01866dc23fbee8b6b2c7b1e14d6f05c28cd35a2c934239f94095602a0.
//
// Solidity: event NewResolver(bytes32 indexed node, address resolver)
func (_ENSRegistry *ENSRegistryFilterer) FilterNewResolver(opts *bind.FilterOpts, node [][32]byte) (*ENSRegistryNewResolverIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _ENSRegistry.contract.FilterLogs(opts, "NewResolver", nodeRule)
	if err != nil {
		return nil, err
	}
	return &ENSRegistryNewResolverIterator{contract: _ENSRegistry.contract, event: "NewResolver", logs: logs, sub: sub}, nil
}

// WatchNewResolver is a free log subscription operation binding the contract event 0x335721b01866dc23fbee8b6b2c7b1e14d6f05c28cd35a2c934239f94095602a0.
//
// Solidity: event NewResolver(bytes32 indexed node, address resolver)
func (_ENSRegistry *ENSRegistryFilterer) WatchNewResolver(opts *bind.WatchOpts, sink chan<- *ENSRegistryNewResolver, node [][32]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _ENSRegistry.contract.WatchLogs(opts, "NewResolver", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ENSRegistryNewResolver)
				if err := _ENSRegistry.contract.UnpackLog(event, "NewResolver", log); err != nil {
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

// ENSRegistryNewTTLIterator is returned from FilterNewTTL and is used to iterate over the raw logs and unpacked data for NewTTL events raised by the ENSRegistry contract.
type ENSRegistryNewTTLIterator struct {
	Event *ENSRegistryNewTTL // Event containing the contract specifics and raw log

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
func (it *ENSRegistryNewTTLIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ENSRegistryNewTTL)
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
		it.Event = new(ENSRegistryNewTTL)
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
func (it *ENSRegistryNewTTLIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ENSRegistryNewTTLIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ENSRegistryNewTTL represents a NewTTL event raised by the ENSRegistry contract.
type ENSRegistryNewTTL struct {
	Node [32]byte
	Ttl  uint64
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterNewTTL is a free log retrieval operation binding the contract event 0x1d4f9bbfc9cab89d66e1a1562f2233ccbf1308cb4f63de2ead5787adddb8fa68.
//
// Solidity: event NewTTL(bytes32 indexed node, uint64 ttl)
func (_ENSRegistry *ENSRegistryFilterer) FilterNewTTL(opts *bind.FilterOpts, node [][32]byte) (*ENSRegistryNewTTLIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _ENSRegistry.contract.FilterLogs(opts, "NewTTL", nodeRule)
	if err != nil {
		return nil, err
	}
	return &ENSRegistryNewTTLIterator{contract: _ENSRegistry.contract, event: "NewTTL", logs: logs, sub: sub}, nil
}

// WatchNewTTL is a free log subscription operation binding the contract event 0x1d4f9bbfc9cab89d66e1a1562f2233ccbf1308cb4f63de2ead5787adddb8fa68.
//
// Solidity: event NewTTL(bytes32 indexed node, uint64 ttl)
func (_ENSRegistry *ENSRegistryFilterer) WatchNewTTL(opts *bind.WatchOpts, sink chan<- *ENSRegistryNewTTL, node [][32]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _ENSRegistry.contract.WatchLogs(opts, "NewTTL", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ENSRegistryNewTTL)
				if err := _ENSRegistry.contract.UnpackLog(event, "NewTTL", log); err != nil {
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

// ENSRegistryTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ENSRegistry contract.
type ENSRegistryTransferIterator struct {
	Event *ENSRegistryTransfer // Event containing the contract specifics and raw log

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
func (it *ENSRegistryTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ENSRegistryTransfer)
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
		it.Event = new(ENSRegistryTransfer)
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
func (it *ENSRegistryTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ENSRegistryTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ENSRegistryTransfer represents a Transfer event raised by the ENSRegistry contract.
type ENSRegistryTransfer struct {
	Node  [32]byte
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xd4735d920b0f87494915f556dd9b54c8f309026070caea5c737245152564d266.
//
// Solidity: event Transfer(bytes32 indexed node, address owner)
func (_ENSRegistry *ENSRegistryFilterer) FilterTransfer(opts *bind.FilterOpts, node [][32]byte) (*ENSRegistryTransferIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _ENSRegistry.contract.FilterLogs(opts, "Transfer", nodeRule)
	if err != nil {
		return nil, err
	}
	return &ENSRegistryTransferIterator{contract: _ENSRegistry.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xd4735d920b0f87494915f556dd9b54c8f309026070caea5c737245152564d266.
//
// Solidity: event Transfer(bytes32 indexed node, address owner)
func (_ENSRegistry *ENSRegistryFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ENSRegistryTransfer, node [][32]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _ENSRegistry.contract.WatchLogs(opts, "Transfer", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ENSRegistryTransfer)
				if err := _ENSRegistry.contract.UnpackLog(event, "Transfer", log); err != nil {
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
