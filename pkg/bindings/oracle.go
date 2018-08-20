// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// OracleABI is the input ABI used to generate the binding from.
const OracleABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"rate\",\"outputs\":[{\"name\":\"value\",\"type\":\"uint256\"},{\"name\":\"decimal\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"convert\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tokens\",\"type\":\"address[]\"},{\"name\":\"_values\",\"type\":\"uint256[]\"},{\"name\":\"_decimals\",\"type\":\"uint256[]\"}],\"name\":\"set\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"controller\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// OracleBin is the compiled bytecode used for deploying new contracts.
const OracleBin = `0x608060405234801561001057600080fd5b50336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506104db806100606000396000f300608060405260043610610062576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680630ba9d8ca1461006757806367c6e39c146100c5578063b5c5d47e14610126578063f77c479114610212575b600080fd5b34801561007357600080fd5b506100a8600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610269565b604051808381526020018281526020019250505060405180910390f35b3480156100d157600080fd5b50610110600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019092919050505061028d565b6040518082815260200191505060405180910390f35b34801561013257600080fd5b50610210600480360381019080803590602001908201803590602001908080602002602001604051908101604052809392919081815260200183836020028082843782019150505050505091929192908035906020019082018035906020019080806020026020016040519081016040528093929190818152602001838360200280828437820191505050505050919291929080359060200190820180359060200190808060200260200160405190810160405280939291908181526020018383602002808284378201915050505050509192919290505050610332565b005b34801561021e57600080fd5b50610227610470565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b60016020528060005260406000206000915090508060000154908060010154905082565b6000610297610495565b6000600160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206040805190810160405290816000820154815260200160018201548152505091508160200151600a0a8260000151850281151561031157fe5b0490506000811415610326576001925061032a565b8092505b505092915050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561038f57600080fd5b8251845114151561039f57600080fd5b600090505b835181101561046a57604080519081016040528084838151811015156103c657fe5b90602001906020020151815260200183838151811015156103e357fe5b9060200190602002015181525060016000868481518110151561040257fe5b9060200190602002015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082015181600001556020820151816001015590505080806001019150506103a4565b50505050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6040805190810160405280600081526020016000815250905600a165627a7a723058202644d43023e16f7ea479b505899d1ad1376d760667b95f15ed4fb6c88ad99a610029`

// DeployOracle deploys a new Ethereum contract, binding an instance of Oracle to it.
func DeployOracle(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Oracle, error) {
	parsed, err := abi.JSON(strings.NewReader(OracleABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OracleBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Oracle{OracleCaller: OracleCaller{contract: contract}, OracleTransactor: OracleTransactor{contract: contract}, OracleFilterer: OracleFilterer{contract: contract}}, nil
}

// Oracle is an auto generated Go binding around an Ethereum contract.
type Oracle struct {
	OracleCaller     // Read-only binding to the contract
	OracleTransactor // Write-only binding to the contract
	OracleFilterer   // Log filterer for contract events
}

// OracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type OracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OracleSession struct {
	Contract     *Oracle           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OracleCallerSession struct {
	Contract *OracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// OracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OracleTransactorSession struct {
	Contract     *OracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type OracleRaw struct {
	Contract *Oracle // Generic contract binding to access the raw methods on
}

// OracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OracleCallerRaw struct {
	Contract *OracleCaller // Generic read-only contract binding to access the raw methods on
}

// OracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OracleTransactorRaw struct {
	Contract *OracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOracle creates a new instance of Oracle, bound to a specific deployed contract.
func NewOracle(address common.Address, backend bind.ContractBackend) (*Oracle, error) {
	contract, err := bindOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Oracle{OracleCaller: OracleCaller{contract: contract}, OracleTransactor: OracleTransactor{contract: contract}, OracleFilterer: OracleFilterer{contract: contract}}, nil
}

// NewOracleCaller creates a new read-only instance of Oracle, bound to a specific deployed contract.
func NewOracleCaller(address common.Address, caller bind.ContractCaller) (*OracleCaller, error) {
	contract, err := bindOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OracleCaller{contract: contract}, nil
}

// NewOracleTransactor creates a new write-only instance of Oracle, bound to a specific deployed contract.
func NewOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*OracleTransactor, error) {
	contract, err := bindOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OracleTransactor{contract: contract}, nil
}

// NewOracleFilterer creates a new log filterer instance of Oracle, bound to a specific deployed contract.
func NewOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*OracleFilterer, error) {
	contract, err := bindOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OracleFilterer{contract: contract}, nil
}

// bindOracle binds a generic wrapper to an already deployed contract.
func bindOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OracleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Oracle *OracleRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Oracle.Contract.OracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Oracle *OracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Oracle.Contract.OracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Oracle *OracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Oracle.Contract.OracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Oracle *OracleCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Oracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Oracle *OracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Oracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Oracle *OracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Oracle.Contract.contract.Transact(opts, method, params...)
}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() constant returns(address)
func (_Oracle *OracleCaller) Controller(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Oracle.contract.Call(opts, out, "controller")
	return *ret0, err
}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() constant returns(address)
func (_Oracle *OracleSession) Controller() (common.Address, error) {
	return _Oracle.Contract.Controller(&_Oracle.CallOpts)
}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() constant returns(address)
func (_Oracle *OracleCallerSession) Controller() (common.Address, error) {
	return _Oracle.Contract.Controller(&_Oracle.CallOpts)
}

// Convert is a free data retrieval call binding the contract method 0x67c6e39c.
//
// Solidity: function convert(_token address, _amount uint256) constant returns(uint256)
func (_Oracle *OracleCaller) Convert(opts *bind.CallOpts, _token common.Address, _amount *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Oracle.contract.Call(opts, out, "convert", _token, _amount)
	return *ret0, err
}

// Convert is a free data retrieval call binding the contract method 0x67c6e39c.
//
// Solidity: function convert(_token address, _amount uint256) constant returns(uint256)
func (_Oracle *OracleSession) Convert(_token common.Address, _amount *big.Int) (*big.Int, error) {
	return _Oracle.Contract.Convert(&_Oracle.CallOpts, _token, _amount)
}

// Convert is a free data retrieval call binding the contract method 0x67c6e39c.
//
// Solidity: function convert(_token address, _amount uint256) constant returns(uint256)
func (_Oracle *OracleCallerSession) Convert(_token common.Address, _amount *big.Int) (*big.Int, error) {
	return _Oracle.Contract.Convert(&_Oracle.CallOpts, _token, _amount)
}

// Rate is a free data retrieval call binding the contract method 0x0ba9d8ca.
//
// Solidity: function rate( address) constant returns(value uint256, decimal uint256)
func (_Oracle *OracleCaller) Rate(opts *bind.CallOpts, arg0 common.Address) (struct {
	Value   *big.Int
	Decimal *big.Int
}, error) {
	ret := new(struct {
		Value   *big.Int
		Decimal *big.Int
	})
	out := ret
	err := _Oracle.contract.Call(opts, out, "rate", arg0)
	return *ret, err
}

// Rate is a free data retrieval call binding the contract method 0x0ba9d8ca.
//
// Solidity: function rate( address) constant returns(value uint256, decimal uint256)
func (_Oracle *OracleSession) Rate(arg0 common.Address) (struct {
	Value   *big.Int
	Decimal *big.Int
}, error) {
	return _Oracle.Contract.Rate(&_Oracle.CallOpts, arg0)
}

// Rate is a free data retrieval call binding the contract method 0x0ba9d8ca.
//
// Solidity: function rate( address) constant returns(value uint256, decimal uint256)
func (_Oracle *OracleCallerSession) Rate(arg0 common.Address) (struct {
	Value   *big.Int
	Decimal *big.Int
}, error) {
	return _Oracle.Contract.Rate(&_Oracle.CallOpts, arg0)
}

// Set is a paid mutator transaction binding the contract method 0xb5c5d47e.
//
// Solidity: function set(_tokens address[], _values uint256[], _decimals uint256[]) returns()
func (_Oracle *OracleTransactor) Set(opts *bind.TransactOpts, _tokens []common.Address, _values []*big.Int, _decimals []*big.Int) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "set", _tokens, _values, _decimals)
}

// Set is a paid mutator transaction binding the contract method 0xb5c5d47e.
//
// Solidity: function set(_tokens address[], _values uint256[], _decimals uint256[]) returns()
func (_Oracle *OracleSession) Set(_tokens []common.Address, _values []*big.Int, _decimals []*big.Int) (*types.Transaction, error) {
	return _Oracle.Contract.Set(&_Oracle.TransactOpts, _tokens, _values, _decimals)
}

// Set is a paid mutator transaction binding the contract method 0xb5c5d47e.
//
// Solidity: function set(_tokens address[], _values uint256[], _decimals uint256[]) returns()
func (_Oracle *OracleTransactorSession) Set(_tokens []common.Address, _values []*big.Int, _decimals []*big.Int) (*types.Transaction, error) {
	return _Oracle.Contract.Set(&_Oracle.TransactOpts, _tokens, _values, _decimals)
}
