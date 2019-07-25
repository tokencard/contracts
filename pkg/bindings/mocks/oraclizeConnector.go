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

// OraclizeConnectorABI is the input ABI used to generate the binding from.
const OraclizeConnectorABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"proofType\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes1\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_datasource\",\"type\":\"string\"},{\"name\":\"gaslimit\",\"type\":\"uint256\"}],\"name\":\"getPrice\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_datasource\",\"type\":\"string\"}],\"name\":\"getPrice\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_proofType\",\"type\":\"bytes1\"}],\"name\":\"setProofType\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_timestamp\",\"type\":\"uint256\"},{\"name\":\"_datasource\",\"type\":\"string\"},{\"name\":\"_arg\",\"type\":\"string\"}],\"name\":\"query\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cbAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_timestamp\",\"type\":\"uint256\"},{\"name\":\"_datasource\",\"type\":\"string\"},{\"name\":\"_arg\",\"type\":\"string\"},{\"name\":\"_gaslimit\",\"type\":\"uint256\"}],\"name\":\"query_withGasLimit\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_gasPrice\",\"type\":\"uint256\"}],\"name\":\"setCustomGasPrice\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_cbAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// OraclizeConnectorBin is the compiled bytecode used for deploying new contracts.
const OraclizeConnectorBin = `608060405234801561001057600080fd5b506040516020806107108339810180604052602081101561003057600080fd5b505160008054600160a060020a03909216600160a060020a03199092169190911790556106ae806100626000396000f3fe608060405260043610610098576000357c010000000000000000000000000000000000000000000000000000000090048063adf59f991161006b578063adf59f99146102af578063c281d19e14610378578063c51be90f146103b6578063ca6ad1e41461047f57610098565b806305af5d351461009d5780632ef3accc146100e7578063524f3889146101ae578063688dcfd714610261575b600080fd5b3480156100a957600080fd5b506100b26104a9565b604080517fff000000000000000000000000000000000000000000000000000000000000009092168252519081900360200190f35b3480156100f357600080fd5b5061019c6004803603604081101561010a57600080fd5b81019060208101813564010000000081111561012557600080fd5b82018360208201111561013757600080fd5b8035906020019184600183028401116401000000008311171561015957600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955050913592506104d1915050565b60408051918252519081900360200190f35b3480156101ba57600080fd5b5061019c600480360360208110156101d157600080fd5b8101906020810181356401000000008111156101ec57600080fd5b8201836020820111156101fe57600080fd5b8035906020019184600183028401116401000000008311171561022057600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506104f6945050505050565b34801561026d57600080fd5b506102ad6004803603602081101561028457600080fd5b50357fff0000000000000000000000000000000000000000000000000000000000000016610517565b005b61019c600480360360608110156102c557600080fd5b813591908101906040810160208201356401000000008111156102e757600080fd5b8201836020820111156102f957600080fd5b8035906020019184600183028401116401000000008311171561031b57600080fd5b91939092909160208101903564010000000081111561033957600080fd5b82018360208201111561034b57600080fd5b8035906020019184600183028401116401000000008311171561036d57600080fd5b50909250905061054b565b34801561038457600080fd5b5061038d6105cb565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b61019c600480360360808110156103cc57600080fd5b813591908101906040810160208201356401000000008111156103ee57600080fd5b82018360208201111561040057600080fd5b8035906020019184600183028401116401000000008311171561042257600080fd5b91939092909160208101903564010000000081111561044057600080fd5b82018360208201111561045257600080fd5b8035906020019184600183028401116401000000008311171561047457600080fd5b9193509150356105e7565b34801561048b57600080fd5b506102ad600480360360208110156104a257600080fd5b5035610672565b6002547f01000000000000000000000000000000000000000000000000000000000000000281565b60008115156104ea576104e3836104f6565b90506104f0565b50620f42405b92915050565b600081516000141561050c5750620f4240610512565b50620f42405b919050565b6002805460ff19167f0100000000000000000000000000000000000000000000000000000000000000909204919091179055565b600085158015610559575083155b156105a45761059d83838080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061067792505050565b90506105c2565b82826040518083838082843760405192018290039091209450505050505b95945050505050565b60005473ffffffffffffffffffffffffffffffffffffffff1681565b6000861580156105f5575084155b80156105ff575081155b1561064a5761064384848080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061067792505050565b9050610668565b83836040518083838082843760405192018290039091209450505050505b9695505050505050565b600155565b80516020909101209056fea165627a7a72305820c2b85e9dcaa292f44641c9dd6f71b0822e7281ab52d4335cc2df323b8669c9cd0029`

// DeployOraclizeConnector deploys a new Ethereum contract, binding an instance of OraclizeConnector to it.
func DeployOraclizeConnector(auth *bind.TransactOpts, backend bind.ContractBackend, _cbAddress common.Address) (common.Address, *types.Transaction, *OraclizeConnector, error) {
	parsed, err := abi.JSON(strings.NewReader(OraclizeConnectorABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OraclizeConnectorBin), backend, _cbAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OraclizeConnector{OraclizeConnectorCaller: OraclizeConnectorCaller{contract: contract}, OraclizeConnectorTransactor: OraclizeConnectorTransactor{contract: contract}, OraclizeConnectorFilterer: OraclizeConnectorFilterer{contract: contract}}, nil
}

// OraclizeConnector is an auto generated Go binding around an Ethereum contract.
type OraclizeConnector struct {
	OraclizeConnectorCaller     // Read-only binding to the contract
	OraclizeConnectorTransactor // Write-only binding to the contract
	OraclizeConnectorFilterer   // Log filterer for contract events
}

// OraclizeConnectorCaller is an auto generated read-only Go binding around an Ethereum contract.
type OraclizeConnectorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OraclizeConnectorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OraclizeConnectorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OraclizeConnectorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OraclizeConnectorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OraclizeConnectorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OraclizeConnectorSession struct {
	Contract     *OraclizeConnector // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// OraclizeConnectorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OraclizeConnectorCallerSession struct {
	Contract *OraclizeConnectorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// OraclizeConnectorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OraclizeConnectorTransactorSession struct {
	Contract     *OraclizeConnectorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// OraclizeConnectorRaw is an auto generated low-level Go binding around an Ethereum contract.
type OraclizeConnectorRaw struct {
	Contract *OraclizeConnector // Generic contract binding to access the raw methods on
}

// OraclizeConnectorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OraclizeConnectorCallerRaw struct {
	Contract *OraclizeConnectorCaller // Generic read-only contract binding to access the raw methods on
}

// OraclizeConnectorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OraclizeConnectorTransactorRaw struct {
	Contract *OraclizeConnectorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOraclizeConnector creates a new instance of OraclizeConnector, bound to a specific deployed contract.
func NewOraclizeConnector(address common.Address, backend bind.ContractBackend) (*OraclizeConnector, error) {
	contract, err := bindOraclizeConnector(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OraclizeConnector{OraclizeConnectorCaller: OraclizeConnectorCaller{contract: contract}, OraclizeConnectorTransactor: OraclizeConnectorTransactor{contract: contract}, OraclizeConnectorFilterer: OraclizeConnectorFilterer{contract: contract}}, nil
}

// NewOraclizeConnectorCaller creates a new read-only instance of OraclizeConnector, bound to a specific deployed contract.
func NewOraclizeConnectorCaller(address common.Address, caller bind.ContractCaller) (*OraclizeConnectorCaller, error) {
	contract, err := bindOraclizeConnector(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OraclizeConnectorCaller{contract: contract}, nil
}

// NewOraclizeConnectorTransactor creates a new write-only instance of OraclizeConnector, bound to a specific deployed contract.
func NewOraclizeConnectorTransactor(address common.Address, transactor bind.ContractTransactor) (*OraclizeConnectorTransactor, error) {
	contract, err := bindOraclizeConnector(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OraclizeConnectorTransactor{contract: contract}, nil
}

// NewOraclizeConnectorFilterer creates a new log filterer instance of OraclizeConnector, bound to a specific deployed contract.
func NewOraclizeConnectorFilterer(address common.Address, filterer bind.ContractFilterer) (*OraclizeConnectorFilterer, error) {
	contract, err := bindOraclizeConnector(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OraclizeConnectorFilterer{contract: contract}, nil
}

// bindOraclizeConnector binds a generic wrapper to an already deployed contract.
func bindOraclizeConnector(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OraclizeConnectorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OraclizeConnector *OraclizeConnectorRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _OraclizeConnector.Contract.OraclizeConnectorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OraclizeConnector *OraclizeConnectorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OraclizeConnector.Contract.OraclizeConnectorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OraclizeConnector *OraclizeConnectorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OraclizeConnector.Contract.OraclizeConnectorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OraclizeConnector *OraclizeConnectorCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _OraclizeConnector.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OraclizeConnector *OraclizeConnectorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OraclizeConnector.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OraclizeConnector *OraclizeConnectorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OraclizeConnector.Contract.contract.Transact(opts, method, params...)
}

// CbAddress is a free data retrieval call binding the contract method 0xc281d19e.
//
// Solidity: function cbAddress() constant returns(address)
func (_OraclizeConnector *OraclizeConnectorCaller) CbAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _OraclizeConnector.contract.Call(opts, out, "cbAddress")
	return *ret0, err
}

// CbAddress is a free data retrieval call binding the contract method 0xc281d19e.
//
// Solidity: function cbAddress() constant returns(address)
func (_OraclizeConnector *OraclizeConnectorSession) CbAddress() (common.Address, error) {
	return _OraclizeConnector.Contract.CbAddress(&_OraclizeConnector.CallOpts)
}

// CbAddress is a free data retrieval call binding the contract method 0xc281d19e.
//
// Solidity: function cbAddress() constant returns(address)
func (_OraclizeConnector *OraclizeConnectorCallerSession) CbAddress() (common.Address, error) {
	return _OraclizeConnector.Contract.CbAddress(&_OraclizeConnector.CallOpts)
}

// GetPrice is a free data retrieval call binding the contract method 0x524f3889.
//
// Solidity: function getPrice(string _datasource) constant returns(uint256)
func (_OraclizeConnector *OraclizeConnectorCaller) GetPrice(opts *bind.CallOpts, _datasource string) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _OraclizeConnector.contract.Call(opts, out, "getPrice", _datasource)
	return *ret0, err
}

// GetPrice is a free data retrieval call binding the contract method 0x524f3889.
//
// Solidity: function getPrice(string _datasource) constant returns(uint256)
func (_OraclizeConnector *OraclizeConnectorSession) GetPrice(_datasource string) (*big.Int, error) {
	return _OraclizeConnector.Contract.GetPrice(&_OraclizeConnector.CallOpts, _datasource)
}

// GetPrice is a free data retrieval call binding the contract method 0x524f3889.
//
// Solidity: function getPrice(string _datasource) constant returns(uint256)
func (_OraclizeConnector *OraclizeConnectorCallerSession) GetPrice(_datasource string) (*big.Int, error) {
	return _OraclizeConnector.Contract.GetPrice(&_OraclizeConnector.CallOpts, _datasource)
}

// ProofType is a free data retrieval call binding the contract method 0x05af5d35.
//
// Solidity: function proofType() constant returns(bytes1)
func (_OraclizeConnector *OraclizeConnectorCaller) ProofType(opts *bind.CallOpts) ([1]byte, error) {
	var (
		ret0 = new([1]byte)
	)
	out := ret0
	err := _OraclizeConnector.contract.Call(opts, out, "proofType")
	return *ret0, err
}

// ProofType is a free data retrieval call binding the contract method 0x05af5d35.
//
// Solidity: function proofType() constant returns(bytes1)
func (_OraclizeConnector *OraclizeConnectorSession) ProofType() ([1]byte, error) {
	return _OraclizeConnector.Contract.ProofType(&_OraclizeConnector.CallOpts)
}

// ProofType is a free data retrieval call binding the contract method 0x05af5d35.
//
// Solidity: function proofType() constant returns(bytes1)
func (_OraclizeConnector *OraclizeConnectorCallerSession) ProofType() ([1]byte, error) {
	return _OraclizeConnector.Contract.ProofType(&_OraclizeConnector.CallOpts)
}

// Query is a paid mutator transaction binding the contract method 0xadf59f99.
//
// Solidity: function query(uint256 _timestamp, string _datasource, string _arg) returns(bytes32)
func (_OraclizeConnector *OraclizeConnectorTransactor) Query(opts *bind.TransactOpts, _timestamp *big.Int, _datasource string, _arg string) (*types.Transaction, error) {
	return _OraclizeConnector.contract.Transact(opts, "query", _timestamp, _datasource, _arg)
}

// Query is a paid mutator transaction binding the contract method 0xadf59f99.
//
// Solidity: function query(uint256 _timestamp, string _datasource, string _arg) returns(bytes32)
func (_OraclizeConnector *OraclizeConnectorSession) Query(_timestamp *big.Int, _datasource string, _arg string) (*types.Transaction, error) {
	return _OraclizeConnector.Contract.Query(&_OraclizeConnector.TransactOpts, _timestamp, _datasource, _arg)
}

// Query is a paid mutator transaction binding the contract method 0xadf59f99.
//
// Solidity: function query(uint256 _timestamp, string _datasource, string _arg) returns(bytes32)
func (_OraclizeConnector *OraclizeConnectorTransactorSession) Query(_timestamp *big.Int, _datasource string, _arg string) (*types.Transaction, error) {
	return _OraclizeConnector.Contract.Query(&_OraclizeConnector.TransactOpts, _timestamp, _datasource, _arg)
}

// QueryWithGasLimit is a paid mutator transaction binding the contract method 0xc51be90f.
//
// Solidity: function query_withGasLimit(uint256 _timestamp, string _datasource, string _arg, uint256 _gaslimit) returns(bytes32)
func (_OraclizeConnector *OraclizeConnectorTransactor) QueryWithGasLimit(opts *bind.TransactOpts, _timestamp *big.Int, _datasource string, _arg string, _gaslimit *big.Int) (*types.Transaction, error) {
	return _OraclizeConnector.contract.Transact(opts, "query_withGasLimit", _timestamp, _datasource, _arg, _gaslimit)
}

// QueryWithGasLimit is a paid mutator transaction binding the contract method 0xc51be90f.
//
// Solidity: function query_withGasLimit(uint256 _timestamp, string _datasource, string _arg, uint256 _gaslimit) returns(bytes32)
func (_OraclizeConnector *OraclizeConnectorSession) QueryWithGasLimit(_timestamp *big.Int, _datasource string, _arg string, _gaslimit *big.Int) (*types.Transaction, error) {
	return _OraclizeConnector.Contract.QueryWithGasLimit(&_OraclizeConnector.TransactOpts, _timestamp, _datasource, _arg, _gaslimit)
}

// QueryWithGasLimit is a paid mutator transaction binding the contract method 0xc51be90f.
//
// Solidity: function query_withGasLimit(uint256 _timestamp, string _datasource, string _arg, uint256 _gaslimit) returns(bytes32)
func (_OraclizeConnector *OraclizeConnectorTransactorSession) QueryWithGasLimit(_timestamp *big.Int, _datasource string, _arg string, _gaslimit *big.Int) (*types.Transaction, error) {
	return _OraclizeConnector.Contract.QueryWithGasLimit(&_OraclizeConnector.TransactOpts, _timestamp, _datasource, _arg, _gaslimit)
}

// SetCustomGasPrice is a paid mutator transaction binding the contract method 0xca6ad1e4.
//
// Solidity: function setCustomGasPrice(uint256 _gasPrice) returns()
func (_OraclizeConnector *OraclizeConnectorTransactor) SetCustomGasPrice(opts *bind.TransactOpts, _gasPrice *big.Int) (*types.Transaction, error) {
	return _OraclizeConnector.contract.Transact(opts, "setCustomGasPrice", _gasPrice)
}

// SetCustomGasPrice is a paid mutator transaction binding the contract method 0xca6ad1e4.
//
// Solidity: function setCustomGasPrice(uint256 _gasPrice) returns()
func (_OraclizeConnector *OraclizeConnectorSession) SetCustomGasPrice(_gasPrice *big.Int) (*types.Transaction, error) {
	return _OraclizeConnector.Contract.SetCustomGasPrice(&_OraclizeConnector.TransactOpts, _gasPrice)
}

// SetCustomGasPrice is a paid mutator transaction binding the contract method 0xca6ad1e4.
//
// Solidity: function setCustomGasPrice(uint256 _gasPrice) returns()
func (_OraclizeConnector *OraclizeConnectorTransactorSession) SetCustomGasPrice(_gasPrice *big.Int) (*types.Transaction, error) {
	return _OraclizeConnector.Contract.SetCustomGasPrice(&_OraclizeConnector.TransactOpts, _gasPrice)
}

// SetProofType is a paid mutator transaction binding the contract method 0x688dcfd7.
//
// Solidity: function setProofType(bytes1 _proofType) returns()
func (_OraclizeConnector *OraclizeConnectorTransactor) SetProofType(opts *bind.TransactOpts, _proofType [1]byte) (*types.Transaction, error) {
	return _OraclizeConnector.contract.Transact(opts, "setProofType", _proofType)
}

// SetProofType is a paid mutator transaction binding the contract method 0x688dcfd7.
//
// Solidity: function setProofType(bytes1 _proofType) returns()
func (_OraclizeConnector *OraclizeConnectorSession) SetProofType(_proofType [1]byte) (*types.Transaction, error) {
	return _OraclizeConnector.Contract.SetProofType(&_OraclizeConnector.TransactOpts, _proofType)
}

// SetProofType is a paid mutator transaction binding the contract method 0x688dcfd7.
//
// Solidity: function setProofType(bytes1 _proofType) returns()
func (_OraclizeConnector *OraclizeConnectorTransactorSession) SetProofType(_proofType [1]byte) (*types.Transaction, error) {
	return _OraclizeConnector.Contract.SetProofType(&_OraclizeConnector.TransactOpts, _proofType)
}
