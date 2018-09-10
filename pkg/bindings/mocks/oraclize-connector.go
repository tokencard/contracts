// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package mocks

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// OraclizeABI is the input ABI used to generate the binding from.
const OraclizeABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_datasource\",\"type\":\"string\"},{\"name\":\"gaslimit\",\"type\":\"uint256\"}],\"name\":\"getPrice\",\"outputs\":[{\"name\":\"_dsprice\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_datasource\",\"type\":\"string\"}],\"name\":\"getPrice\",\"outputs\":[{\"name\":\"_dsprice\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_proofType\",\"type\":\"bytes1\"}],\"name\":\"setProofType\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_timestamp\",\"type\":\"uint256\"},{\"name\":\"_datasource\",\"type\":\"string\"},{\"name\":\"_arg1\",\"type\":\"string\"},{\"name\":\"_arg2\",\"type\":\"string\"}],\"name\":\"query2\",\"outputs\":[{\"name\":\"_id\",\"type\":\"bytes32\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_timestamp\",\"type\":\"uint256\"},{\"name\":\"_datasource\",\"type\":\"string\"},{\"name\":\"_argN\",\"type\":\"bytes\"}],\"name\":\"queryN\",\"outputs\":[{\"name\":\"_id\",\"type\":\"bytes32\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_timestamp\",\"type\":\"uint256\"},{\"name\":\"_datasource\",\"type\":\"string\"},{\"name\":\"_arg1\",\"type\":\"string\"},{\"name\":\"_arg2\",\"type\":\"string\"},{\"name\":\"_gaslimit\",\"type\":\"uint256\"}],\"name\":\"query2_withGasLimit\",\"outputs\":[{\"name\":\"_id\",\"type\":\"bytes32\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"randomDS_getSessionPubKeyHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_timestamp\",\"type\":\"uint256\"},{\"name\":\"_datasource\",\"type\":\"string\"},{\"name\":\"_arg\",\"type\":\"string\"}],\"name\":\"query\",\"outputs\":[{\"name\":\"_id\",\"type\":\"bytes32\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cbAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_timestamp\",\"type\":\"uint256\"},{\"name\":\"_datasource\",\"type\":\"string\"},{\"name\":\"_arg\",\"type\":\"string\"},{\"name\":\"_gaslimit\",\"type\":\"uint256\"}],\"name\":\"query_withGasLimit\",\"outputs\":[{\"name\":\"_id\",\"type\":\"bytes32\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_timestamp\",\"type\":\"uint256\"},{\"name\":\"_datasource\",\"type\":\"string\"},{\"name\":\"_argN\",\"type\":\"bytes\"},{\"name\":\"_gaslimit\",\"type\":\"uint256\"}],\"name\":\"queryN_withGasLimit\",\"outputs\":[{\"name\":\"_id\",\"type\":\"bytes32\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_gasPrice\",\"type\":\"uint256\"}],\"name\":\"setCustomGasPrice\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_cbAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// OraclizeBin is the compiled bytecode used for deploying new contracts.
const OraclizeBin = `608060405234801561001057600080fd5b506040516020806107e783398101806040528101908080519060200190929190505050806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050610764806100836000396000f3006080604052600436106100ba576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680632ef3accc146100bf578063524f388914610146578063688dcfd7146101c3578063772286591461021257806383eed3d51461032057806385dee34c146103e8578063abaa5f3e14610476578063adf59f99146104a9578063c281d19e14610515578063c51be90f1461056c578063c55c1cb6146105e2578063ca6ad1e414610658575b600080fd5b3480156100cb57600080fd5b50610130600480360381019080803590602001908201803590602001908080601f016020809104026020016040519081016040528093929190818152602001838380828437820191505050505050919291929080359060200190929190505050610685565b6040518082815260200191505060405180910390f35b34801561015257600080fd5b506101ad600480360381019080803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290505050610690565b6040518082815260200191505060405180910390f35b3480156101cf57600080fd5b5061021060048036038101908080357effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916906020019092919050505061069a565b005b61030260048036038101908080359060200190929190803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290803590602001908201803590602001908080601f016020809104026020016040519081016040528093929190818152602001838380828437820191505050505050919291929050505061069d565b60405180826000191660001916815260200191505060405180910390f35b6103ca60048036038101908080359060200190929190803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290803590602001908201803590602001908080601f01602080910402602001604051908101604052809392919081815260200183838082843782019150505050505091929192905050506106ad565b60405180826000191660001916815260200191505060405180910390f35b61045860048036038101908080359060200190929190803590602001908201803590602001919091929391929390803590602001908201803590602001919091929391929390803590602001908201803590602001919091929391929390803590602001909291905050506106bc565b60405180826000191660001916815260200191505060405180910390f35b34801561048257600080fd5b5061048b6106d0565b60405180826000191660001916815260200191505060405180910390f35b6104f7600480360381019080803590602001909291908035906020019082018035906020019190919293919293908035906020019082018035906020019190919293919293905050506106db565b60405180826000191660001916815260200191505060405180910390f35b34801561052157600080fd5b5061052a6106ec565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6105c46004803603810190808035906020019092919080359060200190820180359060200191909192939192939080359060200190820180359060200191909192939192939080359060200190929190505050610711565b60405180826000191660001916815260200191505060405180910390f35b61063a6004803603810190808035906020019092919080359060200190820180359060200191909192939192939080359060200190820180359060200191909192939192939080359060200190929190505050610723565b60405180826000191660001916815260200191505060405180910390f35b34801561066457600080fd5b5061068360048036038101908080359060200190929190505050610735565b005b600080905092915050565b6000809050919050565b50565b6000806001029050949350505050565b60008060010290509392505050565b600080600102905098975050505050505050565b600080600102905090565b600080600102905095945050505050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008060010290509695505050505050565b60008060010290509695505050505050565b505600a165627a7a72305820048be8a8b912587b2613b118a2b79cfd2c3fb7d3679b4968ae646c0a88d7d0950029`

// DeployOraclize deploys a new Ethereum contract, binding an instance of Oraclize to it.
func DeployOraclize(auth *bind.TransactOpts, backend bind.ContractBackend, _cbAddress common.Address) (common.Address, *types.Transaction, *Oraclize, error) {
	parsed, err := abi.JSON(strings.NewReader(OraclizeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OraclizeBin), backend, _cbAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Oraclize{OraclizeCaller: OraclizeCaller{contract: contract}, OraclizeTransactor: OraclizeTransactor{contract: contract}, OraclizeFilterer: OraclizeFilterer{contract: contract}}, nil
}

// Oraclize is an auto generated Go binding around an Ethereum contract.
type Oraclize struct {
	OraclizeCaller     // Read-only binding to the contract
	OraclizeTransactor // Write-only binding to the contract
	OraclizeFilterer   // Log filterer for contract events
}

// OraclizeCaller is an auto generated read-only Go binding around an Ethereum contract.
type OraclizeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OraclizeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OraclizeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OraclizeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OraclizeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OraclizeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OraclizeSession struct {
	Contract     *Oraclize         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OraclizeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OraclizeCallerSession struct {
	Contract *OraclizeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// OraclizeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OraclizeTransactorSession struct {
	Contract     *OraclizeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// OraclizeRaw is an auto generated low-level Go binding around an Ethereum contract.
type OraclizeRaw struct {
	Contract *Oraclize // Generic contract binding to access the raw methods on
}

// OraclizeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OraclizeCallerRaw struct {
	Contract *OraclizeCaller // Generic read-only contract binding to access the raw methods on
}

// OraclizeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OraclizeTransactorRaw struct {
	Contract *OraclizeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOraclize creates a new instance of Oraclize, bound to a specific deployed contract.
func NewOraclize(address common.Address, backend bind.ContractBackend) (*Oraclize, error) {
	contract, err := bindOraclize(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Oraclize{OraclizeCaller: OraclizeCaller{contract: contract}, OraclizeTransactor: OraclizeTransactor{contract: contract}, OraclizeFilterer: OraclizeFilterer{contract: contract}}, nil
}

// NewOraclizeCaller creates a new read-only instance of Oraclize, bound to a specific deployed contract.
func NewOraclizeCaller(address common.Address, caller bind.ContractCaller) (*OraclizeCaller, error) {
	contract, err := bindOraclize(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OraclizeCaller{contract: contract}, nil
}

// NewOraclizeTransactor creates a new write-only instance of Oraclize, bound to a specific deployed contract.
func NewOraclizeTransactor(address common.Address, transactor bind.ContractTransactor) (*OraclizeTransactor, error) {
	contract, err := bindOraclize(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OraclizeTransactor{contract: contract}, nil
}

// NewOraclizeFilterer creates a new log filterer instance of Oraclize, bound to a specific deployed contract.
func NewOraclizeFilterer(address common.Address, filterer bind.ContractFilterer) (*OraclizeFilterer, error) {
	contract, err := bindOraclize(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OraclizeFilterer{contract: contract}, nil
}

// bindOraclize binds a generic wrapper to an already deployed contract.
func bindOraclize(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OraclizeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Oraclize *OraclizeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Oraclize.Contract.OraclizeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Oraclize *OraclizeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Oraclize.Contract.OraclizeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Oraclize *OraclizeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Oraclize.Contract.OraclizeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Oraclize *OraclizeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Oraclize.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Oraclize *OraclizeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Oraclize.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Oraclize *OraclizeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Oraclize.Contract.contract.Transact(opts, method, params...)
}

// CbAddress is a free data retrieval call binding the contract method 0xc281d19e.
//
// Solidity: function cbAddress() constant returns(address)
func (_Oraclize *OraclizeCaller) CbAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Oraclize.contract.Call(opts, out, "cbAddress")
	return *ret0, err
}

// CbAddress is a free data retrieval call binding the contract method 0xc281d19e.
//
// Solidity: function cbAddress() constant returns(address)
func (_Oraclize *OraclizeSession) CbAddress() (common.Address, error) {
	return _Oraclize.Contract.CbAddress(&_Oraclize.CallOpts)
}

// CbAddress is a free data retrieval call binding the contract method 0xc281d19e.
//
// Solidity: function cbAddress() constant returns(address)
func (_Oraclize *OraclizeCallerSession) CbAddress() (common.Address, error) {
	return _Oraclize.Contract.CbAddress(&_Oraclize.CallOpts)
}

// RandomDSGetSessionPubKeyHash is a free data retrieval call binding the contract method 0xabaa5f3e.
//
// Solidity: function randomDS_getSessionPubKeyHash() constant returns(bytes32)
func (_Oraclize *OraclizeCaller) RandomDSGetSessionPubKeyHash(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Oraclize.contract.Call(opts, out, "randomDS_getSessionPubKeyHash")
	return *ret0, err
}

// RandomDSGetSessionPubKeyHash is a free data retrieval call binding the contract method 0xabaa5f3e.
//
// Solidity: function randomDS_getSessionPubKeyHash() constant returns(bytes32)
func (_Oraclize *OraclizeSession) RandomDSGetSessionPubKeyHash() ([32]byte, error) {
	return _Oraclize.Contract.RandomDSGetSessionPubKeyHash(&_Oraclize.CallOpts)
}

// RandomDSGetSessionPubKeyHash is a free data retrieval call binding the contract method 0xabaa5f3e.
//
// Solidity: function randomDS_getSessionPubKeyHash() constant returns(bytes32)
func (_Oraclize *OraclizeCallerSession) RandomDSGetSessionPubKeyHash() ([32]byte, error) {
	return _Oraclize.Contract.RandomDSGetSessionPubKeyHash(&_Oraclize.CallOpts)
}

// GetPrice is a paid mutator transaction binding the contract method 0x524f3889.
//
// Solidity: function getPrice(_datasource string) returns(_dsprice uint256)
func (_Oraclize *OraclizeTransactor) GetPrice(opts *bind.TransactOpts, _datasource string) (*types.Transaction, error) {
	return _Oraclize.contract.Transact(opts, "getPrice", _datasource)
}

// GetPrice is a paid mutator transaction binding the contract method 0x524f3889.
//
// Solidity: function getPrice(_datasource string) returns(_dsprice uint256)
func (_Oraclize *OraclizeSession) GetPrice(_datasource string) (*types.Transaction, error) {
	return _Oraclize.Contract.GetPrice(&_Oraclize.TransactOpts, _datasource)
}

// GetPrice is a paid mutator transaction binding the contract method 0x524f3889.
//
// Solidity: function getPrice(_datasource string) returns(_dsprice uint256)
func (_Oraclize *OraclizeTransactorSession) GetPrice(_datasource string) (*types.Transaction, error) {
	return _Oraclize.Contract.GetPrice(&_Oraclize.TransactOpts, _datasource)
}

// Query is a paid mutator transaction binding the contract method 0xadf59f99.
//
// Solidity: function query(_timestamp uint256, _datasource string, _arg string) returns(_id bytes32)
func (_Oraclize *OraclizeTransactor) Query(opts *bind.TransactOpts, _timestamp *big.Int, _datasource string, _arg string) (*types.Transaction, error) {
	return _Oraclize.contract.Transact(opts, "query", _timestamp, _datasource, _arg)
}

// Query is a paid mutator transaction binding the contract method 0xadf59f99.
//
// Solidity: function query(_timestamp uint256, _datasource string, _arg string) returns(_id bytes32)
func (_Oraclize *OraclizeSession) Query(_timestamp *big.Int, _datasource string, _arg string) (*types.Transaction, error) {
	return _Oraclize.Contract.Query(&_Oraclize.TransactOpts, _timestamp, _datasource, _arg)
}

// Query is a paid mutator transaction binding the contract method 0xadf59f99.
//
// Solidity: function query(_timestamp uint256, _datasource string, _arg string) returns(_id bytes32)
func (_Oraclize *OraclizeTransactorSession) Query(_timestamp *big.Int, _datasource string, _arg string) (*types.Transaction, error) {
	return _Oraclize.Contract.Query(&_Oraclize.TransactOpts, _timestamp, _datasource, _arg)
}

// Query2 is a paid mutator transaction binding the contract method 0x77228659.
//
// Solidity: function query2(_timestamp uint256, _datasource string, _arg1 string, _arg2 string) returns(_id bytes32)
func (_Oraclize *OraclizeTransactor) Query2(opts *bind.TransactOpts, _timestamp *big.Int, _datasource string, _arg1 string, _arg2 string) (*types.Transaction, error) {
	return _Oraclize.contract.Transact(opts, "query2", _timestamp, _datasource, _arg1, _arg2)
}

// Query2 is a paid mutator transaction binding the contract method 0x77228659.
//
// Solidity: function query2(_timestamp uint256, _datasource string, _arg1 string, _arg2 string) returns(_id bytes32)
func (_Oraclize *OraclizeSession) Query2(_timestamp *big.Int, _datasource string, _arg1 string, _arg2 string) (*types.Transaction, error) {
	return _Oraclize.Contract.Query2(&_Oraclize.TransactOpts, _timestamp, _datasource, _arg1, _arg2)
}

// Query2 is a paid mutator transaction binding the contract method 0x77228659.
//
// Solidity: function query2(_timestamp uint256, _datasource string, _arg1 string, _arg2 string) returns(_id bytes32)
func (_Oraclize *OraclizeTransactorSession) Query2(_timestamp *big.Int, _datasource string, _arg1 string, _arg2 string) (*types.Transaction, error) {
	return _Oraclize.Contract.Query2(&_Oraclize.TransactOpts, _timestamp, _datasource, _arg1, _arg2)
}

// Query2WithGasLimit is a paid mutator transaction binding the contract method 0x85dee34c.
//
// Solidity: function query2_withGasLimit(_timestamp uint256, _datasource string, _arg1 string, _arg2 string, _gaslimit uint256) returns(_id bytes32)
func (_Oraclize *OraclizeTransactor) Query2WithGasLimit(opts *bind.TransactOpts, _timestamp *big.Int, _datasource string, _arg1 string, _arg2 string, _gaslimit *big.Int) (*types.Transaction, error) {
	return _Oraclize.contract.Transact(opts, "query2_withGasLimit", _timestamp, _datasource, _arg1, _arg2, _gaslimit)
}

// Query2WithGasLimit is a paid mutator transaction binding the contract method 0x85dee34c.
//
// Solidity: function query2_withGasLimit(_timestamp uint256, _datasource string, _arg1 string, _arg2 string, _gaslimit uint256) returns(_id bytes32)
func (_Oraclize *OraclizeSession) Query2WithGasLimit(_timestamp *big.Int, _datasource string, _arg1 string, _arg2 string, _gaslimit *big.Int) (*types.Transaction, error) {
	return _Oraclize.Contract.Query2WithGasLimit(&_Oraclize.TransactOpts, _timestamp, _datasource, _arg1, _arg2, _gaslimit)
}

// Query2WithGasLimit is a paid mutator transaction binding the contract method 0x85dee34c.
//
// Solidity: function query2_withGasLimit(_timestamp uint256, _datasource string, _arg1 string, _arg2 string, _gaslimit uint256) returns(_id bytes32)
func (_Oraclize *OraclizeTransactorSession) Query2WithGasLimit(_timestamp *big.Int, _datasource string, _arg1 string, _arg2 string, _gaslimit *big.Int) (*types.Transaction, error) {
	return _Oraclize.Contract.Query2WithGasLimit(&_Oraclize.TransactOpts, _timestamp, _datasource, _arg1, _arg2, _gaslimit)
}

// QueryN is a paid mutator transaction binding the contract method 0x83eed3d5.
//
// Solidity: function queryN(_timestamp uint256, _datasource string, _argN bytes) returns(_id bytes32)
func (_Oraclize *OraclizeTransactor) QueryN(opts *bind.TransactOpts, _timestamp *big.Int, _datasource string, _argN []byte) (*types.Transaction, error) {
	return _Oraclize.contract.Transact(opts, "queryN", _timestamp, _datasource, _argN)
}

// QueryN is a paid mutator transaction binding the contract method 0x83eed3d5.
//
// Solidity: function queryN(_timestamp uint256, _datasource string, _argN bytes) returns(_id bytes32)
func (_Oraclize *OraclizeSession) QueryN(_timestamp *big.Int, _datasource string, _argN []byte) (*types.Transaction, error) {
	return _Oraclize.Contract.QueryN(&_Oraclize.TransactOpts, _timestamp, _datasource, _argN)
}

// QueryN is a paid mutator transaction binding the contract method 0x83eed3d5.
//
// Solidity: function queryN(_timestamp uint256, _datasource string, _argN bytes) returns(_id bytes32)
func (_Oraclize *OraclizeTransactorSession) QueryN(_timestamp *big.Int, _datasource string, _argN []byte) (*types.Transaction, error) {
	return _Oraclize.Contract.QueryN(&_Oraclize.TransactOpts, _timestamp, _datasource, _argN)
}

// QueryNWithGasLimit is a paid mutator transaction binding the contract method 0xc55c1cb6.
//
// Solidity: function queryN_withGasLimit(_timestamp uint256, _datasource string, _argN bytes, _gaslimit uint256) returns(_id bytes32)
func (_Oraclize *OraclizeTransactor) QueryNWithGasLimit(opts *bind.TransactOpts, _timestamp *big.Int, _datasource string, _argN []byte, _gaslimit *big.Int) (*types.Transaction, error) {
	return _Oraclize.contract.Transact(opts, "queryN_withGasLimit", _timestamp, _datasource, _argN, _gaslimit)
}

// QueryNWithGasLimit is a paid mutator transaction binding the contract method 0xc55c1cb6.
//
// Solidity: function queryN_withGasLimit(_timestamp uint256, _datasource string, _argN bytes, _gaslimit uint256) returns(_id bytes32)
func (_Oraclize *OraclizeSession) QueryNWithGasLimit(_timestamp *big.Int, _datasource string, _argN []byte, _gaslimit *big.Int) (*types.Transaction, error) {
	return _Oraclize.Contract.QueryNWithGasLimit(&_Oraclize.TransactOpts, _timestamp, _datasource, _argN, _gaslimit)
}

// QueryNWithGasLimit is a paid mutator transaction binding the contract method 0xc55c1cb6.
//
// Solidity: function queryN_withGasLimit(_timestamp uint256, _datasource string, _argN bytes, _gaslimit uint256) returns(_id bytes32)
func (_Oraclize *OraclizeTransactorSession) QueryNWithGasLimit(_timestamp *big.Int, _datasource string, _argN []byte, _gaslimit *big.Int) (*types.Transaction, error) {
	return _Oraclize.Contract.QueryNWithGasLimit(&_Oraclize.TransactOpts, _timestamp, _datasource, _argN, _gaslimit)
}

// QueryWithGasLimit is a paid mutator transaction binding the contract method 0xc51be90f.
//
// Solidity: function query_withGasLimit(_timestamp uint256, _datasource string, _arg string, _gaslimit uint256) returns(_id bytes32)
func (_Oraclize *OraclizeTransactor) QueryWithGasLimit(opts *bind.TransactOpts, _timestamp *big.Int, _datasource string, _arg string, _gaslimit *big.Int) (*types.Transaction, error) {
	return _Oraclize.contract.Transact(opts, "query_withGasLimit", _timestamp, _datasource, _arg, _gaslimit)
}

// QueryWithGasLimit is a paid mutator transaction binding the contract method 0xc51be90f.
//
// Solidity: function query_withGasLimit(_timestamp uint256, _datasource string, _arg string, _gaslimit uint256) returns(_id bytes32)
func (_Oraclize *OraclizeSession) QueryWithGasLimit(_timestamp *big.Int, _datasource string, _arg string, _gaslimit *big.Int) (*types.Transaction, error) {
	return _Oraclize.Contract.QueryWithGasLimit(&_Oraclize.TransactOpts, _timestamp, _datasource, _arg, _gaslimit)
}

// QueryWithGasLimit is a paid mutator transaction binding the contract method 0xc51be90f.
//
// Solidity: function query_withGasLimit(_timestamp uint256, _datasource string, _arg string, _gaslimit uint256) returns(_id bytes32)
func (_Oraclize *OraclizeTransactorSession) QueryWithGasLimit(_timestamp *big.Int, _datasource string, _arg string, _gaslimit *big.Int) (*types.Transaction, error) {
	return _Oraclize.Contract.QueryWithGasLimit(&_Oraclize.TransactOpts, _timestamp, _datasource, _arg, _gaslimit)
}

// SetCustomGasPrice is a paid mutator transaction binding the contract method 0xca6ad1e4.
//
// Solidity: function setCustomGasPrice(_gasPrice uint256) returns()
func (_Oraclize *OraclizeTransactor) SetCustomGasPrice(opts *bind.TransactOpts, _gasPrice *big.Int) (*types.Transaction, error) {
	return _Oraclize.contract.Transact(opts, "setCustomGasPrice", _gasPrice)
}

// SetCustomGasPrice is a paid mutator transaction binding the contract method 0xca6ad1e4.
//
// Solidity: function setCustomGasPrice(_gasPrice uint256) returns()
func (_Oraclize *OraclizeSession) SetCustomGasPrice(_gasPrice *big.Int) (*types.Transaction, error) {
	return _Oraclize.Contract.SetCustomGasPrice(&_Oraclize.TransactOpts, _gasPrice)
}

// SetCustomGasPrice is a paid mutator transaction binding the contract method 0xca6ad1e4.
//
// Solidity: function setCustomGasPrice(_gasPrice uint256) returns()
func (_Oraclize *OraclizeTransactorSession) SetCustomGasPrice(_gasPrice *big.Int) (*types.Transaction, error) {
	return _Oraclize.Contract.SetCustomGasPrice(&_Oraclize.TransactOpts, _gasPrice)
}

// SetProofType is a paid mutator transaction binding the contract method 0x688dcfd7.
//
// Solidity: function setProofType(_proofType bytes1) returns()
func (_Oraclize *OraclizeTransactor) SetProofType(opts *bind.TransactOpts, _proofType [1]byte) (*types.Transaction, error) {
	return _Oraclize.contract.Transact(opts, "setProofType", _proofType)
}

// SetProofType is a paid mutator transaction binding the contract method 0x688dcfd7.
//
// Solidity: function setProofType(_proofType bytes1) returns()
func (_Oraclize *OraclizeSession) SetProofType(_proofType [1]byte) (*types.Transaction, error) {
	return _Oraclize.Contract.SetProofType(&_Oraclize.TransactOpts, _proofType)
}

// SetProofType is a paid mutator transaction binding the contract method 0x688dcfd7.
//
// Solidity: function setProofType(_proofType bytes1) returns()
func (_Oraclize *OraclizeTransactorSession) SetProofType(_proofType [1]byte) (*types.Transaction, error) {
	return _Oraclize.Contract.SetProofType(&_Oraclize.TransactOpts, _proofType)
}
