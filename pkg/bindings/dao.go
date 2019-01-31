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

// DaoABI is the input ABI used to generate the binding from.
const DaoABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"isTransferable\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"licence\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isLocked\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_account\",\"type\":\"address\"},{\"name\":\"_transferable\",\"type\":\"bool\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newFee\",\"type\":\"uint256\"}],\"name\":\"updateLicenceFee\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"lock\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_transferable\",\"type\":\"bool\"},{\"name\":\"_licence\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"TransferredOwnership\",\"type\":\"event\"}]"

// DaoBin is the compiled bytecode used for deploying new contracts.
const DaoBin = `608060405234801561001057600080fd5b5060405160608061079283398101604081815282516020808501519483015160008054600160a060020a031916600160a060020a03851690811760a060020a60ff02191674010000000000000000000000000000000000000000891515021782559086529185019190915282519194939092859285927f850b3df64837d7d518b45f5aa64d104652c3b80eb5b34a8e3d9eb666cb7cdea592908290030190a1505060018054600160a060020a03909216600160a060020a031990921691909117905550506106af806100e36000396000f30060806040526004361061008d5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416632121dc75811461009257806340ace468146100bb578063715018a6146100ec5780638da5cb5b14610103578063a4e2d63414610118578063b242e5341461012d578063de3d80d314610153578063f83d08ba1461016b575b600080fd5b34801561009e57600080fd5b506100a7610180565b604080519115158252519081900360200190f35b3480156100c757600080fd5b506100d06101a1565b60408051600160a060020a039092168252519081900360200190f35b3480156100f857600080fd5b506101016101b0565b005b34801561010f57600080fd5b506100d06102d9565b34801561012457600080fd5b506100a76102e8565b34801561013957600080fd5b50610101600160a060020a0360043516602435151561030a565b34801561015f57600080fd5b50610101600435610500565b34801561017757600080fd5b506101016105cd565b60005474010000000000000000000000000000000000000000900460ff1690565b600154600160a060020a031681565b6101b8610652565b15156101fc576040805160e560020a62461bcd0281526020600482015260166024820152600080516020610664833981519152604482015290519081900360640190fd5b60005474010000000000000000000000000000000000000000900460ff161515610270576040805160e560020a62461bcd02815260206004820152601d60248201527f6f776e657273686970206973206e6f74207472616e7366657261626c65000000604482015290519081900360640190fd5b6000805460408051600160a060020a039092168252602082019290925281517f850b3df64837d7d518b45f5aa64d104652c3b80eb5b34a8e3d9eb666cb7cdea5929181900390910190a16000805473ffffffffffffffffffffffffffffffffffffffff19169055565b600054600160a060020a031690565b6000547501000000000000000000000000000000000000000000900460ff1690565b610312610652565b1515610356576040805160e560020a62461bcd0281526020600482015260166024820152600080516020610664833981519152604482015290519081900360640190fd5b60005474010000000000000000000000000000000000000000900460ff1615156103ca576040805160e560020a62461bcd02815260206004820152601d60248201527f6f776e657273686970206973206e6f74207472616e7366657261626c65000000604482015290519081900360640190fd5b600160a060020a0382161515610450576040805160e560020a62461bcd02815260206004820152602360248201527f6f776e65722063616e6e6f742062652073657420746f207a65726f206164647260448201527f6573730000000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b6000805474ff000000000000000000000000000000000000000019167401000000000000000000000000000000000000000083151502179081905560408051600160a060020a039283168152918416602083015280517f850b3df64837d7d518b45f5aa64d104652c3b80eb5b34a8e3d9eb666cb7cdea59281900390910190a1506000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b610508610652565b151561054c576040805160e560020a62461bcd0281526020600482015260166024820152600080516020610664833981519152604482015290519081900360640190fd5b600154604080517f9012c4a8000000000000000000000000000000000000000000000000000000008152600481018490529051600160a060020a0390921691639012c4a89160248082019260009290919082900301818387803b1580156105b257600080fd5b505af11580156105c6573d6000803e3d6000fd5b5050505050565b6105d5610652565b1515610619576040805160e560020a62461bcd0281526020600482015260166024820152600080516020610664833981519152604482015290519081900360640190fd5b6000805475ff00000000000000000000000000000000000000000019167501000000000000000000000000000000000000000000179055565b600054600160a060020a0316331490560073656e646572206973206e6f7420616e206f776e657200000000000000000000a165627a7a72305820eb7f08a545c5cb9859fc7b413617562472e949c20fe344d861e3d0ab60749beb0029`

// DeployDao deploys a new Ethereum contract, binding an instance of Dao to it.
func DeployDao(auth *bind.TransactOpts, backend bind.ContractBackend, _owner common.Address, _transferable bool, _licence common.Address) (common.Address, *types.Transaction, *Dao, error) {
	parsed, err := abi.JSON(strings.NewReader(DaoABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DaoBin), backend, _owner, _transferable, _licence)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Dao{DaoCaller: DaoCaller{contract: contract}, DaoTransactor: DaoTransactor{contract: contract}, DaoFilterer: DaoFilterer{contract: contract}}, nil
}

// Dao is an auto generated Go binding around an Ethereum contract.
type Dao struct {
	DaoCaller     // Read-only binding to the contract
	DaoTransactor // Write-only binding to the contract
	DaoFilterer   // Log filterer for contract events
}

// DaoCaller is an auto generated read-only Go binding around an Ethereum contract.
type DaoCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DaoTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DaoTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DaoFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DaoFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DaoSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DaoSession struct {
	Contract     *Dao              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DaoCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DaoCallerSession struct {
	Contract *DaoCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// DaoTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DaoTransactorSession struct {
	Contract     *DaoTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DaoRaw is an auto generated low-level Go binding around an Ethereum contract.
type DaoRaw struct {
	Contract *Dao // Generic contract binding to access the raw methods on
}

// DaoCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DaoCallerRaw struct {
	Contract *DaoCaller // Generic read-only contract binding to access the raw methods on
}

// DaoTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DaoTransactorRaw struct {
	Contract *DaoTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDao creates a new instance of Dao, bound to a specific deployed contract.
func NewDao(address common.Address, backend bind.ContractBackend) (*Dao, error) {
	contract, err := bindDao(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Dao{DaoCaller: DaoCaller{contract: contract}, DaoTransactor: DaoTransactor{contract: contract}, DaoFilterer: DaoFilterer{contract: contract}}, nil
}

// NewDaoCaller creates a new read-only instance of Dao, bound to a specific deployed contract.
func NewDaoCaller(address common.Address, caller bind.ContractCaller) (*DaoCaller, error) {
	contract, err := bindDao(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DaoCaller{contract: contract}, nil
}

// NewDaoTransactor creates a new write-only instance of Dao, bound to a specific deployed contract.
func NewDaoTransactor(address common.Address, transactor bind.ContractTransactor) (*DaoTransactor, error) {
	contract, err := bindDao(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DaoTransactor{contract: contract}, nil
}

// NewDaoFilterer creates a new log filterer instance of Dao, bound to a specific deployed contract.
func NewDaoFilterer(address common.Address, filterer bind.ContractFilterer) (*DaoFilterer, error) {
	contract, err := bindDao(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DaoFilterer{contract: contract}, nil
}

// bindDao binds a generic wrapper to an already deployed contract.
func bindDao(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DaoABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dao *DaoRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Dao.Contract.DaoCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dao *DaoRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dao.Contract.DaoTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dao *DaoRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dao.Contract.DaoTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dao *DaoCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Dao.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dao *DaoTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dao.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dao *DaoTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dao.Contract.contract.Transact(opts, method, params...)
}

// IsLocked is a free data retrieval call binding the contract method 0xa4e2d634.
//
// Solidity: function isLocked() constant returns(bool)
func (_Dao *DaoCaller) IsLocked(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Dao.contract.Call(opts, out, "isLocked")
	return *ret0, err
}

// IsLocked is a free data retrieval call binding the contract method 0xa4e2d634.
//
// Solidity: function isLocked() constant returns(bool)
func (_Dao *DaoSession) IsLocked() (bool, error) {
	return _Dao.Contract.IsLocked(&_Dao.CallOpts)
}

// IsLocked is a free data retrieval call binding the contract method 0xa4e2d634.
//
// Solidity: function isLocked() constant returns(bool)
func (_Dao *DaoCallerSession) IsLocked() (bool, error) {
	return _Dao.Contract.IsLocked(&_Dao.CallOpts)
}

// IsTransferable is a free data retrieval call binding the contract method 0x2121dc75.
//
// Solidity: function isTransferable() constant returns(bool)
func (_Dao *DaoCaller) IsTransferable(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Dao.contract.Call(opts, out, "isTransferable")
	return *ret0, err
}

// IsTransferable is a free data retrieval call binding the contract method 0x2121dc75.
//
// Solidity: function isTransferable() constant returns(bool)
func (_Dao *DaoSession) IsTransferable() (bool, error) {
	return _Dao.Contract.IsTransferable(&_Dao.CallOpts)
}

// IsTransferable is a free data retrieval call binding the contract method 0x2121dc75.
//
// Solidity: function isTransferable() constant returns(bool)
func (_Dao *DaoCallerSession) IsTransferable() (bool, error) {
	return _Dao.Contract.IsTransferable(&_Dao.CallOpts)
}

// Licence is a free data retrieval call binding the contract method 0x40ace468.
//
// Solidity: function licence() constant returns(address)
func (_Dao *DaoCaller) Licence(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dao.contract.Call(opts, out, "licence")
	return *ret0, err
}

// Licence is a free data retrieval call binding the contract method 0x40ace468.
//
// Solidity: function licence() constant returns(address)
func (_Dao *DaoSession) Licence() (common.Address, error) {
	return _Dao.Contract.Licence(&_Dao.CallOpts)
}

// Licence is a free data retrieval call binding the contract method 0x40ace468.
//
// Solidity: function licence() constant returns(address)
func (_Dao *DaoCallerSession) Licence() (common.Address, error) {
	return _Dao.Contract.Licence(&_Dao.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Dao *DaoCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dao.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Dao *DaoSession) Owner() (common.Address, error) {
	return _Dao.Contract.Owner(&_Dao.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Dao *DaoCallerSession) Owner() (common.Address, error) {
	return _Dao.Contract.Owner(&_Dao.CallOpts)
}

// Lock is a paid mutator transaction binding the contract method 0xf83d08ba.
//
// Solidity: function lock() returns()
func (_Dao *DaoTransactor) Lock(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dao.contract.Transact(opts, "lock")
}

// Lock is a paid mutator transaction binding the contract method 0xf83d08ba.
//
// Solidity: function lock() returns()
func (_Dao *DaoSession) Lock() (*types.Transaction, error) {
	return _Dao.Contract.Lock(&_Dao.TransactOpts)
}

// Lock is a paid mutator transaction binding the contract method 0xf83d08ba.
//
// Solidity: function lock() returns()
func (_Dao *DaoTransactorSession) Lock() (*types.Transaction, error) {
	return _Dao.Contract.Lock(&_Dao.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Dao *DaoTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dao.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Dao *DaoSession) RenounceOwnership() (*types.Transaction, error) {
	return _Dao.Contract.RenounceOwnership(&_Dao.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Dao *DaoTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Dao.Contract.RenounceOwnership(&_Dao.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xb242e534.
//
// Solidity: function transferOwnership(address _account, bool _transferable) returns()
func (_Dao *DaoTransactor) TransferOwnership(opts *bind.TransactOpts, _account common.Address, _transferable bool) (*types.Transaction, error) {
	return _Dao.contract.Transact(opts, "transferOwnership", _account, _transferable)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xb242e534.
//
// Solidity: function transferOwnership(address _account, bool _transferable) returns()
func (_Dao *DaoSession) TransferOwnership(_account common.Address, _transferable bool) (*types.Transaction, error) {
	return _Dao.Contract.TransferOwnership(&_Dao.TransactOpts, _account, _transferable)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xb242e534.
//
// Solidity: function transferOwnership(address _account, bool _transferable) returns()
func (_Dao *DaoTransactorSession) TransferOwnership(_account common.Address, _transferable bool) (*types.Transaction, error) {
	return _Dao.Contract.TransferOwnership(&_Dao.TransactOpts, _account, _transferable)
}

// UpdateLicenceFee is a paid mutator transaction binding the contract method 0xde3d80d3.
//
// Solidity: function updateLicenceFee(uint256 _newFee) returns()
func (_Dao *DaoTransactor) UpdateLicenceFee(opts *bind.TransactOpts, _newFee *big.Int) (*types.Transaction, error) {
	return _Dao.contract.Transact(opts, "updateLicenceFee", _newFee)
}

// UpdateLicenceFee is a paid mutator transaction binding the contract method 0xde3d80d3.
//
// Solidity: function updateLicenceFee(uint256 _newFee) returns()
func (_Dao *DaoSession) UpdateLicenceFee(_newFee *big.Int) (*types.Transaction, error) {
	return _Dao.Contract.UpdateLicenceFee(&_Dao.TransactOpts, _newFee)
}

// UpdateLicenceFee is a paid mutator transaction binding the contract method 0xde3d80d3.
//
// Solidity: function updateLicenceFee(uint256 _newFee) returns()
func (_Dao *DaoTransactorSession) UpdateLicenceFee(_newFee *big.Int) (*types.Transaction, error) {
	return _Dao.Contract.UpdateLicenceFee(&_Dao.TransactOpts, _newFee)
}

// DaoTransferredOwnershipIterator is returned from FilterTransferredOwnership and is used to iterate over the raw logs and unpacked data for TransferredOwnership events raised by the Dao contract.
type DaoTransferredOwnershipIterator struct {
	Event *DaoTransferredOwnership // Event containing the contract specifics and raw log

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
func (it *DaoTransferredOwnershipIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DaoTransferredOwnership)
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
		it.Event = new(DaoTransferredOwnership)
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
func (it *DaoTransferredOwnershipIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DaoTransferredOwnershipIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DaoTransferredOwnership represents a TransferredOwnership event raised by the Dao contract.
type DaoTransferredOwnership struct {
	From common.Address
	To   common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterTransferredOwnership is a free log retrieval operation binding the contract event 0x850b3df64837d7d518b45f5aa64d104652c3b80eb5b34a8e3d9eb666cb7cdea5.
//
// Solidity: event TransferredOwnership(address _from, address _to)
func (_Dao *DaoFilterer) FilterTransferredOwnership(opts *bind.FilterOpts) (*DaoTransferredOwnershipIterator, error) {

	logs, sub, err := _Dao.contract.FilterLogs(opts, "TransferredOwnership")
	if err != nil {
		return nil, err
	}
	return &DaoTransferredOwnershipIterator{contract: _Dao.contract, event: "TransferredOwnership", logs: logs, sub: sub}, nil
}

// WatchTransferredOwnership is a free log subscription operation binding the contract event 0x850b3df64837d7d518b45f5aa64d104652c3b80eb5b34a8e3d9eb666cb7cdea5.
//
// Solidity: event TransferredOwnership(address _from, address _to)
func (_Dao *DaoFilterer) WatchTransferredOwnership(opts *bind.WatchOpts, sink chan<- *DaoTransferredOwnership) (event.Subscription, error) {

	logs, sub, err := _Dao.contract.WatchLogs(opts, "TransferredOwnership")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DaoTransferredOwnership)
				if err := _Dao.contract.UnpackLog(event, "TransferredOwnership", log); err != nil {
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
