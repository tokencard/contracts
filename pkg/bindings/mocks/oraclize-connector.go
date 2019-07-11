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

// OraclizeABI is the input ABI used to generate the binding from.
const OraclizeABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_datasource\",\"type\":\"string\"},{\"name\":\"gaslimit\",\"type\":\"uint256\"}],\"name\":\"getPrice\",\"outputs\":[{\"name\":\"_dsprice\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_datasource\",\"type\":\"string\"}],\"name\":\"getPrice\",\"outputs\":[{\"name\":\"_dsprice\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_proofType\",\"type\":\"bytes1\"}],\"name\":\"setProofType\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_timestamp\",\"type\":\"uint256\"},{\"name\":\"_datasource\",\"type\":\"string\"},{\"name\":\"_arg1\",\"type\":\"string\"},{\"name\":\"_arg2\",\"type\":\"string\"}],\"name\":\"query2\",\"outputs\":[{\"name\":\"_id\",\"type\":\"bytes32\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_timestamp\",\"type\":\"uint256\"},{\"name\":\"_datasource\",\"type\":\"string\"},{\"name\":\"_argN\",\"type\":\"bytes\"}],\"name\":\"queryN\",\"outputs\":[{\"name\":\"_id\",\"type\":\"bytes32\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_timestamp\",\"type\":\"uint256\"},{\"name\":\"_datasource\",\"type\":\"string\"},{\"name\":\"_arg1\",\"type\":\"string\"},{\"name\":\"_arg2\",\"type\":\"string\"},{\"name\":\"_gaslimit\",\"type\":\"uint256\"}],\"name\":\"query2_withGasLimit\",\"outputs\":[{\"name\":\"_id\",\"type\":\"bytes32\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"randomDS_getSessionPubKeyHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_timestamp\",\"type\":\"uint256\"},{\"name\":\"_datasource\",\"type\":\"string\"},{\"name\":\"_arg\",\"type\":\"string\"}],\"name\":\"query\",\"outputs\":[{\"name\":\"_id\",\"type\":\"bytes32\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cbAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_timestamp\",\"type\":\"uint256\"},{\"name\":\"_datasource\",\"type\":\"string\"},{\"name\":\"_arg\",\"type\":\"string\"},{\"name\":\"_gaslimit\",\"type\":\"uint256\"}],\"name\":\"query_withGasLimit\",\"outputs\":[{\"name\":\"_id\",\"type\":\"bytes32\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_timestamp\",\"type\":\"uint256\"},{\"name\":\"_datasource\",\"type\":\"string\"},{\"name\":\"_argN\",\"type\":\"bytes\"},{\"name\":\"_gaslimit\",\"type\":\"uint256\"}],\"name\":\"queryN_withGasLimit\",\"outputs\":[{\"name\":\"_id\",\"type\":\"bytes32\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_gasPrice\",\"type\":\"uint256\"}],\"name\":\"setCustomGasPrice\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_cbAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// OraclizeBin is the compiled bytecode used for deploying new contracts.
const OraclizeBin = `608060405234801561001057600080fd5b50604051602080610558833981016040525160008054600160a060020a03909216600160a060020a0319909216919091179055610506806100526000396000f3006080604052600436106100b95763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416632ef3accc81146100be578063524f38891461012b578063688dcfd71461018457806377228659146101c057806383eed3d51461028d57806385dee34c1461031c578063abaa5f3e1461034e578063adf59f9914610363578063c281d19e14610386578063c51be90f146103c4578063c55c1cb6146103ea578063ca6ad1e414610410575b600080fd5b3480156100ca57600080fd5b506040805160206004803580820135601f810184900484028501840190955284845261011994369492936024939284019190819084018382808284375094975050933594506104289350505050565b60408051918252519081900360200190f35b34801561013757600080fd5b506040805160206004803580820135601f81018490048402850184019095528484526101199436949293602493928401919081908401838280828437509497506104329650505050505050565b34801561019057600080fd5b506101be7fff000000000000000000000000000000000000000000000000000000000000006004351661043a565b005b60408051602060046024803582810135601f810185900485028601850190965285855261011995833595369560449491939091019190819084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a99988101979196509182019450925082915084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a99988101979196509182019450925082915084018382808284375094975061043d9650505050505050565b60408051602060046024803582810135601f810185900485028601850190965285855261011995833595369560449491939091019190819084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a9998810197919650918201945092508291508401838280828437509497506104479650505050505050565b610119600480359060248035808201929081013591604435808201929081013591606435908101910135608435610450565b34801561035a57600080fd5b5061011961045e565b610119600480359060248035808201929081013591604435908101910135610463565b34801561039257600080fd5b5061039b61048a565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b6101196004803590602480358082019290810135916044359081019101356064356104a6565b6101196004803590602480358082019290810135916044359081019101356064356104ce565b34801561041c57600080fd5b506101be60043561043a565b620f424092915050565b50620f424090565b50565b6000949350505050565b60009392505050565b600098975050505050505050565b600090565b600082826040518083838082843760405192018290039091209a9950505050505050505050565b60005473ffffffffffffffffffffffffffffffffffffffff1681565b600083836040518083838082843760405192018290039091209b9a5050505050505050505050565b600096955050505050505600a165627a7a72305820b43fc7ac283606d98759c8bad8a5823049b31552f4528161ad9fad30fa5d7d550029`

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
