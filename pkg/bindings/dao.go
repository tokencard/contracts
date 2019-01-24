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
const DaoABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"isLocked\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"locked\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"lock\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// DaoBin is the compiled bytecode used for deploying new contracts.
const DaoBin = `608060405234801561001057600080fd5b5060f08061001f6000396000f30060806040526004361060525763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663a4e2d63481146057578063cf30901214607d578063f83d08ba14608f575b600080fd5b348015606257600080fd5b50606960a3565b604080519115158252519081900360200190f35b348015608857600080fd5b50606960ac565b348015609a57600080fd5b5060a160b5565b005b60005460ff1690565b60005460ff1681565b6000805460ff191660011790555600a165627a7a72305820785c2a90abc9ad01521e5101e6f37cdb0c5a311bec47c79886bf746df69888150029`

// DeployDao deploys a new Ethereum contract, binding an instance of Dao to it.
func DeployDao(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Dao, error) {
	parsed, err := abi.JSON(strings.NewReader(DaoABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DaoBin), backend)
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

// Locked is a free data retrieval call binding the contract method 0xcf309012.
//
// Solidity: function locked() constant returns(bool)
func (_Dao *DaoCaller) Locked(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Dao.contract.Call(opts, out, "locked")
	return *ret0, err
}

// Locked is a free data retrieval call binding the contract method 0xcf309012.
//
// Solidity: function locked() constant returns(bool)
func (_Dao *DaoSession) Locked() (bool, error) {
	return _Dao.Contract.Locked(&_Dao.CallOpts)
}

// Locked is a free data retrieval call binding the contract method 0xcf309012.
//
// Solidity: function locked() constant returns(bool)
func (_Dao *DaoCallerSession) Locked() (bool, error) {
	return _Dao.Contract.Locked(&_Dao.CallOpts)
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
