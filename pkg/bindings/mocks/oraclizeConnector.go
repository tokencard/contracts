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
const OraclizeConnectorABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_cbAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"cbAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_datasource\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"gaslimit\",\"type\":\"uint256\"}],\"name\":\"getPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_datasource\",\"type\":\"string\"}],\"name\":\"getPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proofType\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"\",\"type\":\"bytes1\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_datasource\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_arg\",\"type\":\"string\"}],\"name\":\"query\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_datasource\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_arg\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_gaslimit\",\"type\":\"uint256\"}],\"name\":\"query_withGasLimit\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_gasPrice\",\"type\":\"uint256\"}],\"name\":\"setCustomGasPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes1\",\"name\":\"_proofType\",\"type\":\"bytes1\"}],\"name\":\"setProofType\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// OraclizeConnectorBin is the compiled bytecode used for deploying new contracts.
var OraclizeConnectorBin = "0x608060405234801561001057600080fd5b5060405161066b38038061066b8339818101604052602081101561003357600080fd5b5051600080546001600160a01b039092166001600160a01b0319909216919091179055610606806100656000396000f3fe60806040526004361061007b5760003560e01c8063adf59f991161004e578063adf59f991461025e578063c281d19e14610323578063c51be90f14610354578063ca6ad1e4146104195761007b565b806305af5d35146100805780632ef3accc146100b2578063524f388914610177578063688dcfd714610228575b600080fd5b34801561008c57600080fd5b50610095610443565b604080516001600160f81b03199092168252519081900360200190f35b3480156100be57600080fd5b50610165600480360360408110156100d557600080fd5b810190602081018135600160201b8111156100ef57600080fd5b82018360208201111561010157600080fd5b803590602001918460018302840111600160201b8311171561012257600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550509135925061044c915050565b60408051918252519081900360200190f35b34801561018357600080fd5b506101656004803603602081101561019a57600080fd5b810190602081018135600160201b8111156101b457600080fd5b8201836020820111156101c657600080fd5b803590602001918460018302840111600160201b831117156101e757600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955061046f945050505050565b34801561023457600080fd5b5061025c6004803603602081101561024b57600080fd5b50356001600160f81b031916610490565b005b6101656004803603606081101561027457600080fd5b81359190810190604081016020820135600160201b81111561029557600080fd5b8201836020820111156102a757600080fd5b803590602001918460018302840111600160201b831117156102c857600080fd5b919390929091602081019035600160201b8111156102e557600080fd5b8201836020820111156102f757600080fd5b803590602001918460018302840111600160201b8311171561031857600080fd5b5090925090506104a6565b34801561032f57600080fd5b50610338610526565b604080516001600160a01b039092168252519081900360200190f35b6101656004803603608081101561036a57600080fd5b81359190810190604081016020820135600160201b81111561038b57600080fd5b82018360208201111561039d57600080fd5b803590602001918460018302840111600160201b831117156103be57600080fd5b919390929091602081019035600160201b8111156103db57600080fd5b8201836020820111156103ed57600080fd5b803590602001918460018302840111600160201b8311171561040e57600080fd5b919350915035610535565b34801561042557600080fd5b5061025c6004803603602081101561043c57600080fd5b50356105c0565b60025460f81b81565b6000816104635761045c8361046f565b9050610469565b50620f42405b92915050565b60008151600014156104855750620f424061048b565b50620f42405b919050565b6002805460ff191660f89290921c919091179055565b6000851580156104b4575083155b156104ff576104f883838080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506105c592505050565b905061051d565b82826040518083838082843760405192018290039091209450505050505b95945050505050565b6000546001600160a01b031681565b600086158015610543575084155b801561054d575081155b156105985761059184848080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506105c592505050565b90506105b6565b83836040518083838082843760405192018290039091209450505050505b9695505050505050565b600155565b80516020909101209056fea26469706673582212202a781f5281ae72473eca77e685639d0e8c9bc848586e3fa610dfa8447ab345ce64736f6c634300060b0033"

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

// GetPrice is a free data retrieval call binding the contract method 0x2ef3accc.
//
// Solidity: function getPrice(string _datasource, uint256 gaslimit) constant returns(uint256)
func (_OraclizeConnector *OraclizeConnectorCaller) GetPrice(opts *bind.CallOpts, _datasource string, gaslimit *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _OraclizeConnector.contract.Call(opts, out, "getPrice", _datasource, gaslimit)
	return *ret0, err
}

// GetPrice is a free data retrieval call binding the contract method 0x2ef3accc.
//
// Solidity: function getPrice(string _datasource, uint256 gaslimit) constant returns(uint256)
func (_OraclizeConnector *OraclizeConnectorSession) GetPrice(_datasource string, gaslimit *big.Int) (*big.Int, error) {
	return _OraclizeConnector.Contract.GetPrice(&_OraclizeConnector.CallOpts, _datasource, gaslimit)
}

// GetPrice is a free data retrieval call binding the contract method 0x2ef3accc.
//
// Solidity: function getPrice(string _datasource, uint256 gaslimit) constant returns(uint256)
func (_OraclizeConnector *OraclizeConnectorCallerSession) GetPrice(_datasource string, gaslimit *big.Int) (*big.Int, error) {
	return _OraclizeConnector.Contract.GetPrice(&_OraclizeConnector.CallOpts, _datasource, gaslimit)
}

// GetPrice0 is a free data retrieval call binding the contract method 0x524f3889.
//
// Solidity: function getPrice(string _datasource) constant returns(uint256)
func (_OraclizeConnector *OraclizeConnectorCaller) GetPrice0(opts *bind.CallOpts, _datasource string) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _OraclizeConnector.contract.Call(opts, out, "getPrice0", _datasource)
	return *ret0, err
}

// GetPrice0 is a free data retrieval call binding the contract method 0x524f3889.
//
// Solidity: function getPrice(string _datasource) constant returns(uint256)
func (_OraclizeConnector *OraclizeConnectorSession) GetPrice0(_datasource string) (*big.Int, error) {
	return _OraclizeConnector.Contract.GetPrice0(&_OraclizeConnector.CallOpts, _datasource)
}

// GetPrice0 is a free data retrieval call binding the contract method 0x524f3889.
//
// Solidity: function getPrice(string _datasource) constant returns(uint256)
func (_OraclizeConnector *OraclizeConnectorCallerSession) GetPrice0(_datasource string) (*big.Int, error) {
	return _OraclizeConnector.Contract.GetPrice0(&_OraclizeConnector.CallOpts, _datasource)
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
