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

// OracleABI is the input ABI used to generate the binding from.
const OracleABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"myid\",\"type\":\"bytes32\"},{\"name\":\"result\",\"type\":\"string\"}],\"name\":\"__callback\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_queryID\",\"type\":\"bytes32\"},{\"name\":\"_result\",\"type\":\"string\"},{\"name\":\"_proof\",\"type\":\"bytes\"}],\"name\":\"__callback\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"convert\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_gasLimit\",\"type\":\"uint256\"},{\"name\":\"_tokenList\",\"type\":\"address[]\"}],\"name\":\"updateTokenRatesList\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_publicKey\",\"type\":\"bytes\"}],\"name\":\"updateAPIPublicKey\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_json\",\"type\":\"string\"}],\"name\":\"parseRate\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"name\":\"updateTokenRates\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_gasPrice\",\"type\":\"uint256\"}],\"name\":\"setCustomGasPrice\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"APIPublicKey\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_resolver\",\"type\":\"address\"},{\"name\":\"_ens\",\"type\":\"address\"},{\"name\":\"_controllerName\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_gasPrice\",\"type\":\"uint256\"}],\"name\":\"SetGasPrice\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_token\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_ether\",\"type\":\"uint256\"}],\"name\":\"Converted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_symbol\",\"type\":\"string\"}],\"name\":\"RequestedUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_reason\",\"type\":\"string\"}],\"name\":\"FailedUpdateRequest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_publicKey\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"_result\",\"type\":\"string\"}],\"name\":\"VerifiedProof\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_publicKey\",\"type\":\"bytes\"}],\"name\":\"SetCryptoComparePublicKey\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_asset\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"Claimed\",\"type\":\"event\"}]"

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

// APIPublicKey is a free data retrieval call binding the contract method 0xcc204119.
//
// Solidity: function APIPublicKey() constant returns(bytes)
func (_Oracle *OracleCaller) APIPublicKey(opts *bind.CallOpts) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Oracle.contract.Call(opts, out, "APIPublicKey")
	return *ret0, err
}

// APIPublicKey is a free data retrieval call binding the contract method 0xcc204119.
//
// Solidity: function APIPublicKey() constant returns(bytes)
func (_Oracle *OracleSession) APIPublicKey() ([]byte, error) {
	return _Oracle.Contract.APIPublicKey(&_Oracle.CallOpts)
}

// APIPublicKey is a free data retrieval call binding the contract method 0xcc204119.
//
// Solidity: function APIPublicKey() constant returns(bytes)
func (_Oracle *OracleCallerSession) APIPublicKey() ([]byte, error) {
	return _Oracle.Contract.APIPublicKey(&_Oracle.CallOpts)
}

// Convert is a free data retrieval call binding the contract method 0x67c6e39c.
//
// Solidity: function convert(address _token, uint256 _amount) constant returns(bool, uint256)
func (_Oracle *OracleCaller) Convert(opts *bind.CallOpts, _token common.Address, _amount *big.Int) (bool, *big.Int, error) {
	var (
		ret0 = new(bool)
		ret1 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Oracle.contract.Call(opts, out, "convert", _token, _amount)
	return *ret0, *ret1, err
}

// Convert is a free data retrieval call binding the contract method 0x67c6e39c.
//
// Solidity: function convert(address _token, uint256 _amount) constant returns(bool, uint256)
func (_Oracle *OracleSession) Convert(_token common.Address, _amount *big.Int) (bool, *big.Int, error) {
	return _Oracle.Contract.Convert(&_Oracle.CallOpts, _token, _amount)
}

// Convert is a free data retrieval call binding the contract method 0x67c6e39c.
//
// Solidity: function convert(address _token, uint256 _amount) constant returns(bool, uint256)
func (_Oracle *OracleCallerSession) Convert(_token common.Address, _amount *big.Int) (bool, *big.Int, error) {
	return _Oracle.Contract.Convert(&_Oracle.CallOpts, _token, _amount)
}

// ParseRate is a free data retrieval call binding the contract method 0xa780b2f3.
//
// Solidity: function parseRate(string _json) constant returns(string)
func (_Oracle *OracleCaller) ParseRate(opts *bind.CallOpts, _json string) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Oracle.contract.Call(opts, out, "parseRate", _json)
	return *ret0, err
}

// ParseRate is a free data retrieval call binding the contract method 0xa780b2f3.
//
// Solidity: function parseRate(string _json) constant returns(string)
func (_Oracle *OracleSession) ParseRate(_json string) (string, error) {
	return _Oracle.Contract.ParseRate(&_Oracle.CallOpts, _json)
}

// ParseRate is a free data retrieval call binding the contract method 0xa780b2f3.
//
// Solidity: function parseRate(string _json) constant returns(string)
func (_Oracle *OracleCallerSession) ParseRate(_json string) (string, error) {
	return _Oracle.Contract.ParseRate(&_Oracle.CallOpts, _json)
}

// Callback is a paid mutator transaction binding the contract method 0x38bbfa50.
//
// Solidity: function __callback(bytes32 _queryID, string _result, bytes _proof) returns()
func (_Oracle *OracleTransactor) Callback(opts *bind.TransactOpts, _queryID [32]byte, _result string, _proof []byte) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "__callback", _queryID, _result, _proof)
}

// Callback is a paid mutator transaction binding the contract method 0x38bbfa50.
//
// Solidity: function __callback(bytes32 _queryID, string _result, bytes _proof) returns()
func (_Oracle *OracleSession) Callback(_queryID [32]byte, _result string, _proof []byte) (*types.Transaction, error) {
	return _Oracle.Contract.Callback(&_Oracle.TransactOpts, _queryID, _result, _proof)
}

// Callback is a paid mutator transaction binding the contract method 0x38bbfa50.
//
// Solidity: function __callback(bytes32 _queryID, string _result, bytes _proof) returns()
func (_Oracle *OracleTransactorSession) Callback(_queryID [32]byte, _result string, _proof []byte) (*types.Transaction, error) {
	return _Oracle.Contract.Callback(&_Oracle.TransactOpts, _queryID, _result, _proof)
}

// SetCustomGasPrice is a paid mutator transaction binding the contract method 0xca6ad1e4.
//
// Solidity: function setCustomGasPrice(uint256 _gasPrice) returns()
func (_Oracle *OracleTransactor) SetCustomGasPrice(opts *bind.TransactOpts, _gasPrice *big.Int) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "setCustomGasPrice", _gasPrice)
}

// SetCustomGasPrice is a paid mutator transaction binding the contract method 0xca6ad1e4.
//
// Solidity: function setCustomGasPrice(uint256 _gasPrice) returns()
func (_Oracle *OracleSession) SetCustomGasPrice(_gasPrice *big.Int) (*types.Transaction, error) {
	return _Oracle.Contract.SetCustomGasPrice(&_Oracle.TransactOpts, _gasPrice)
}

// SetCustomGasPrice is a paid mutator transaction binding the contract method 0xca6ad1e4.
//
// Solidity: function setCustomGasPrice(uint256 _gasPrice) returns()
func (_Oracle *OracleTransactorSession) SetCustomGasPrice(_gasPrice *big.Int) (*types.Transaction, error) {
	return _Oracle.Contract.SetCustomGasPrice(&_Oracle.TransactOpts, _gasPrice)
}

// UpdateAPIPublicKey is a paid mutator transaction binding the contract method 0x9f6f99ee.
//
// Solidity: function updateAPIPublicKey(bytes _publicKey) returns()
func (_Oracle *OracleTransactor) UpdateAPIPublicKey(opts *bind.TransactOpts, _publicKey []byte) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "updateAPIPublicKey", _publicKey)
}

// UpdateAPIPublicKey is a paid mutator transaction binding the contract method 0x9f6f99ee.
//
// Solidity: function updateAPIPublicKey(bytes _publicKey) returns()
func (_Oracle *OracleSession) UpdateAPIPublicKey(_publicKey []byte) (*types.Transaction, error) {
	return _Oracle.Contract.UpdateAPIPublicKey(&_Oracle.TransactOpts, _publicKey)
}

// UpdateAPIPublicKey is a paid mutator transaction binding the contract method 0x9f6f99ee.
//
// Solidity: function updateAPIPublicKey(bytes _publicKey) returns()
func (_Oracle *OracleTransactorSession) UpdateAPIPublicKey(_publicKey []byte) (*types.Transaction, error) {
	return _Oracle.Contract.UpdateAPIPublicKey(&_Oracle.TransactOpts, _publicKey)
}

// UpdateTokenRates is a paid mutator transaction binding the contract method 0xb598f882.
//
// Solidity: function updateTokenRates(uint256 _gasLimit) returns()
func (_Oracle *OracleTransactor) UpdateTokenRates(opts *bind.TransactOpts, _gasLimit *big.Int) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "updateTokenRates", _gasLimit)
}

// UpdateTokenRates is a paid mutator transaction binding the contract method 0xb598f882.
//
// Solidity: function updateTokenRates(uint256 _gasLimit) returns()
func (_Oracle *OracleSession) UpdateTokenRates(_gasLimit *big.Int) (*types.Transaction, error) {
	return _Oracle.Contract.UpdateTokenRates(&_Oracle.TransactOpts, _gasLimit)
}

// UpdateTokenRates is a paid mutator transaction binding the contract method 0xb598f882.
//
// Solidity: function updateTokenRates(uint256 _gasLimit) returns()
func (_Oracle *OracleTransactorSession) UpdateTokenRates(_gasLimit *big.Int) (*types.Transaction, error) {
	return _Oracle.Contract.UpdateTokenRates(&_Oracle.TransactOpts, _gasLimit)
}

// UpdateTokenRatesList is a paid mutator transaction binding the contract method 0x937f54a4.
//
// Solidity: function updateTokenRatesList(uint256 _gasLimit, address[] _tokenList) returns()
func (_Oracle *OracleTransactor) UpdateTokenRatesList(opts *bind.TransactOpts, _gasLimit *big.Int, _tokenList []common.Address) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "updateTokenRatesList", _gasLimit, _tokenList)
}

// UpdateTokenRatesList is a paid mutator transaction binding the contract method 0x937f54a4.
//
// Solidity: function updateTokenRatesList(uint256 _gasLimit, address[] _tokenList) returns()
func (_Oracle *OracleSession) UpdateTokenRatesList(_gasLimit *big.Int, _tokenList []common.Address) (*types.Transaction, error) {
	return _Oracle.Contract.UpdateTokenRatesList(&_Oracle.TransactOpts, _gasLimit, _tokenList)
}

// UpdateTokenRatesList is a paid mutator transaction binding the contract method 0x937f54a4.
//
// Solidity: function updateTokenRatesList(uint256 _gasLimit, address[] _tokenList) returns()
func (_Oracle *OracleTransactorSession) UpdateTokenRatesList(_gasLimit *big.Int, _tokenList []common.Address) (*types.Transaction, error) {
	return _Oracle.Contract.UpdateTokenRatesList(&_Oracle.TransactOpts, _gasLimit, _tokenList)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address _to, uint256 _amount) returns()
func (_Oracle *OracleTransactor) Withdraw(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "withdraw", _to, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address _to, uint256 _amount) returns()
func (_Oracle *OracleSession) Withdraw(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Oracle.Contract.Withdraw(&_Oracle.TransactOpts, _to, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address _to, uint256 _amount) returns()
func (_Oracle *OracleTransactorSession) Withdraw(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Oracle.Contract.Withdraw(&_Oracle.TransactOpts, _to, _amount)
}

// OracleClaimedIterator is returned from FilterClaimed and is used to iterate over the raw logs and unpacked data for Claimed events raised by the Oracle contract.
type OracleClaimedIterator struct {
	Event *OracleClaimed // Event containing the contract specifics and raw log

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
func (it *OracleClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleClaimed)
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
		it.Event = new(OracleClaimed)
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
func (it *OracleClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleClaimed represents a Claimed event raised by the Oracle contract.
type OracleClaimed struct {
	To     common.Address
	Asset  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterClaimed is a free log retrieval operation binding the contract event 0xf7a40077ff7a04c7e61f6f26fb13774259ddf1b6bce9ecf26a8276cdd3992683.
//
// Solidity: event Claimed(address _to, address _asset, uint256 _amount)
func (_Oracle *OracleFilterer) FilterClaimed(opts *bind.FilterOpts) (*OracleClaimedIterator, error) {

	logs, sub, err := _Oracle.contract.FilterLogs(opts, "Claimed")
	if err != nil {
		return nil, err
	}
	return &OracleClaimedIterator{contract: _Oracle.contract, event: "Claimed", logs: logs, sub: sub}, nil
}

// WatchClaimed is a free log subscription operation binding the contract event 0xf7a40077ff7a04c7e61f6f26fb13774259ddf1b6bce9ecf26a8276cdd3992683.
//
// Solidity: event Claimed(address _to, address _asset, uint256 _amount)
func (_Oracle *OracleFilterer) WatchClaimed(opts *bind.WatchOpts, sink chan<- *OracleClaimed) (event.Subscription, error) {

	logs, sub, err := _Oracle.contract.WatchLogs(opts, "Claimed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleClaimed)
				if err := _Oracle.contract.UnpackLog(event, "Claimed", log); err != nil {
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

// OracleConvertedIterator is returned from FilterConverted and is used to iterate over the raw logs and unpacked data for Converted events raised by the Oracle contract.
type OracleConvertedIterator struct {
	Event *OracleConverted // Event containing the contract specifics and raw log

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
func (it *OracleConvertedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleConverted)
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
		it.Event = new(OracleConverted)
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
func (it *OracleConvertedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleConvertedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleConverted represents a Converted event raised by the Oracle contract.
type OracleConverted struct {
	Sender common.Address
	Token  common.Address
	Amount *big.Int
	Ether  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterConverted is a free log retrieval operation binding the contract event 0x3eb9c827477658b168595969c2b83c99d882f7741df2545eb56754b796a968be.
//
// Solidity: event Converted(address _sender, address _token, uint256 _amount, uint256 _ether)
func (_Oracle *OracleFilterer) FilterConverted(opts *bind.FilterOpts) (*OracleConvertedIterator, error) {

	logs, sub, err := _Oracle.contract.FilterLogs(opts, "Converted")
	if err != nil {
		return nil, err
	}
	return &OracleConvertedIterator{contract: _Oracle.contract, event: "Converted", logs: logs, sub: sub}, nil
}

// WatchConverted is a free log subscription operation binding the contract event 0x3eb9c827477658b168595969c2b83c99d882f7741df2545eb56754b796a968be.
//
// Solidity: event Converted(address _sender, address _token, uint256 _amount, uint256 _ether)
func (_Oracle *OracleFilterer) WatchConverted(opts *bind.WatchOpts, sink chan<- *OracleConverted) (event.Subscription, error) {

	logs, sub, err := _Oracle.contract.WatchLogs(opts, "Converted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleConverted)
				if err := _Oracle.contract.UnpackLog(event, "Converted", log); err != nil {
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

// OracleFailedUpdateRequestIterator is returned from FilterFailedUpdateRequest and is used to iterate over the raw logs and unpacked data for FailedUpdateRequest events raised by the Oracle contract.
type OracleFailedUpdateRequestIterator struct {
	Event *OracleFailedUpdateRequest // Event containing the contract specifics and raw log

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
func (it *OracleFailedUpdateRequestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleFailedUpdateRequest)
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
		it.Event = new(OracleFailedUpdateRequest)
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
func (it *OracleFailedUpdateRequestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleFailedUpdateRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleFailedUpdateRequest represents a FailedUpdateRequest event raised by the Oracle contract.
type OracleFailedUpdateRequest struct {
	Reason string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFailedUpdateRequest is a free log retrieval operation binding the contract event 0x4eb5629fd8501532aeb93b1b6a5b5b2ae398561e56514ed4b4b0c5ac2d381b6e.
//
// Solidity: event FailedUpdateRequest(string _reason)
func (_Oracle *OracleFilterer) FilterFailedUpdateRequest(opts *bind.FilterOpts) (*OracleFailedUpdateRequestIterator, error) {

	logs, sub, err := _Oracle.contract.FilterLogs(opts, "FailedUpdateRequest")
	if err != nil {
		return nil, err
	}
	return &OracleFailedUpdateRequestIterator{contract: _Oracle.contract, event: "FailedUpdateRequest", logs: logs, sub: sub}, nil
}

// WatchFailedUpdateRequest is a free log subscription operation binding the contract event 0x4eb5629fd8501532aeb93b1b6a5b5b2ae398561e56514ed4b4b0c5ac2d381b6e.
//
// Solidity: event FailedUpdateRequest(string _reason)
func (_Oracle *OracleFilterer) WatchFailedUpdateRequest(opts *bind.WatchOpts, sink chan<- *OracleFailedUpdateRequest) (event.Subscription, error) {

	logs, sub, err := _Oracle.contract.WatchLogs(opts, "FailedUpdateRequest")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleFailedUpdateRequest)
				if err := _Oracle.contract.UnpackLog(event, "FailedUpdateRequest", log); err != nil {
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

// OracleRequestedUpdateIterator is returned from FilterRequestedUpdate and is used to iterate over the raw logs and unpacked data for RequestedUpdate events raised by the Oracle contract.
type OracleRequestedUpdateIterator struct {
	Event *OracleRequestedUpdate // Event containing the contract specifics and raw log

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
func (it *OracleRequestedUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleRequestedUpdate)
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
		it.Event = new(OracleRequestedUpdate)
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
func (it *OracleRequestedUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleRequestedUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleRequestedUpdate represents a RequestedUpdate event raised by the Oracle contract.
type OracleRequestedUpdate struct {
	Symbol string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRequestedUpdate is a free log retrieval operation binding the contract event 0xd28a52bc81f387be057d69ee89db644f94b7064c4acec0d5bc33e201a2031c05.
//
// Solidity: event RequestedUpdate(string _symbol)
func (_Oracle *OracleFilterer) FilterRequestedUpdate(opts *bind.FilterOpts) (*OracleRequestedUpdateIterator, error) {

	logs, sub, err := _Oracle.contract.FilterLogs(opts, "RequestedUpdate")
	if err != nil {
		return nil, err
	}
	return &OracleRequestedUpdateIterator{contract: _Oracle.contract, event: "RequestedUpdate", logs: logs, sub: sub}, nil
}

// WatchRequestedUpdate is a free log subscription operation binding the contract event 0xd28a52bc81f387be057d69ee89db644f94b7064c4acec0d5bc33e201a2031c05.
//
// Solidity: event RequestedUpdate(string _symbol)
func (_Oracle *OracleFilterer) WatchRequestedUpdate(opts *bind.WatchOpts, sink chan<- *OracleRequestedUpdate) (event.Subscription, error) {

	logs, sub, err := _Oracle.contract.WatchLogs(opts, "RequestedUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleRequestedUpdate)
				if err := _Oracle.contract.UnpackLog(event, "RequestedUpdate", log); err != nil {
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

// OracleSetCryptoComparePublicKeyIterator is returned from FilterSetCryptoComparePublicKey and is used to iterate over the raw logs and unpacked data for SetCryptoComparePublicKey events raised by the Oracle contract.
type OracleSetCryptoComparePublicKeyIterator struct {
	Event *OracleSetCryptoComparePublicKey // Event containing the contract specifics and raw log

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
func (it *OracleSetCryptoComparePublicKeyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleSetCryptoComparePublicKey)
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
		it.Event = new(OracleSetCryptoComparePublicKey)
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
func (it *OracleSetCryptoComparePublicKeyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleSetCryptoComparePublicKeyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleSetCryptoComparePublicKey represents a SetCryptoComparePublicKey event raised by the Oracle contract.
type OracleSetCryptoComparePublicKey struct {
	Sender    common.Address
	PublicKey []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSetCryptoComparePublicKey is a free log retrieval operation binding the contract event 0xc6b0860ba9f580e9c5b6ba4e0954fe82827096a99d92e8c2d73009539ea8d9fa.
//
// Solidity: event SetCryptoComparePublicKey(address _sender, bytes _publicKey)
func (_Oracle *OracleFilterer) FilterSetCryptoComparePublicKey(opts *bind.FilterOpts) (*OracleSetCryptoComparePublicKeyIterator, error) {

	logs, sub, err := _Oracle.contract.FilterLogs(opts, "SetCryptoComparePublicKey")
	if err != nil {
		return nil, err
	}
	return &OracleSetCryptoComparePublicKeyIterator{contract: _Oracle.contract, event: "SetCryptoComparePublicKey", logs: logs, sub: sub}, nil
}

// WatchSetCryptoComparePublicKey is a free log subscription operation binding the contract event 0xc6b0860ba9f580e9c5b6ba4e0954fe82827096a99d92e8c2d73009539ea8d9fa.
//
// Solidity: event SetCryptoComparePublicKey(address _sender, bytes _publicKey)
func (_Oracle *OracleFilterer) WatchSetCryptoComparePublicKey(opts *bind.WatchOpts, sink chan<- *OracleSetCryptoComparePublicKey) (event.Subscription, error) {

	logs, sub, err := _Oracle.contract.WatchLogs(opts, "SetCryptoComparePublicKey")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleSetCryptoComparePublicKey)
				if err := _Oracle.contract.UnpackLog(event, "SetCryptoComparePublicKey", log); err != nil {
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

// OracleSetGasPriceIterator is returned from FilterSetGasPrice and is used to iterate over the raw logs and unpacked data for SetGasPrice events raised by the Oracle contract.
type OracleSetGasPriceIterator struct {
	Event *OracleSetGasPrice // Event containing the contract specifics and raw log

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
func (it *OracleSetGasPriceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleSetGasPrice)
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
		it.Event = new(OracleSetGasPrice)
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
func (it *OracleSetGasPriceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleSetGasPriceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleSetGasPrice represents a SetGasPrice event raised by the Oracle contract.
type OracleSetGasPrice struct {
	Sender   common.Address
	GasPrice *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSetGasPrice is a free log retrieval operation binding the contract event 0xfbd406825addb09beef160afc17bb80ba28df4a3533dcd23592b82658a1c5ab4.
//
// Solidity: event SetGasPrice(address _sender, uint256 _gasPrice)
func (_Oracle *OracleFilterer) FilterSetGasPrice(opts *bind.FilterOpts) (*OracleSetGasPriceIterator, error) {

	logs, sub, err := _Oracle.contract.FilterLogs(opts, "SetGasPrice")
	if err != nil {
		return nil, err
	}
	return &OracleSetGasPriceIterator{contract: _Oracle.contract, event: "SetGasPrice", logs: logs, sub: sub}, nil
}

// WatchSetGasPrice is a free log subscription operation binding the contract event 0xfbd406825addb09beef160afc17bb80ba28df4a3533dcd23592b82658a1c5ab4.
//
// Solidity: event SetGasPrice(address _sender, uint256 _gasPrice)
func (_Oracle *OracleFilterer) WatchSetGasPrice(opts *bind.WatchOpts, sink chan<- *OracleSetGasPrice) (event.Subscription, error) {

	logs, sub, err := _Oracle.contract.WatchLogs(opts, "SetGasPrice")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleSetGasPrice)
				if err := _Oracle.contract.UnpackLog(event, "SetGasPrice", log); err != nil {
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

// OracleVerifiedProofIterator is returned from FilterVerifiedProof and is used to iterate over the raw logs and unpacked data for VerifiedProof events raised by the Oracle contract.
type OracleVerifiedProofIterator struct {
	Event *OracleVerifiedProof // Event containing the contract specifics and raw log

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
func (it *OracleVerifiedProofIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleVerifiedProof)
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
		it.Event = new(OracleVerifiedProof)
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
func (it *OracleVerifiedProofIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleVerifiedProofIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleVerifiedProof represents a VerifiedProof event raised by the Oracle contract.
type OracleVerifiedProof struct {
	PublicKey []byte
	Result    string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterVerifiedProof is a free log retrieval operation binding the contract event 0x0902fdd015aa1e56f7e6026b69c0595e82155dcbd83a83a23b40f9fe96babbd9.
//
// Solidity: event VerifiedProof(bytes _publicKey, string _result)
func (_Oracle *OracleFilterer) FilterVerifiedProof(opts *bind.FilterOpts) (*OracleVerifiedProofIterator, error) {

	logs, sub, err := _Oracle.contract.FilterLogs(opts, "VerifiedProof")
	if err != nil {
		return nil, err
	}
	return &OracleVerifiedProofIterator{contract: _Oracle.contract, event: "VerifiedProof", logs: logs, sub: sub}, nil
}

// WatchVerifiedProof is a free log subscription operation binding the contract event 0x0902fdd015aa1e56f7e6026b69c0595e82155dcbd83a83a23b40f9fe96babbd9.
//
// Solidity: event VerifiedProof(bytes _publicKey, string _result)
func (_Oracle *OracleFilterer) WatchVerifiedProof(opts *bind.WatchOpts, sink chan<- *OracleVerifiedProof) (event.Subscription, error) {

	logs, sub, err := _Oracle.contract.WatchLogs(opts, "VerifiedProof")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleVerifiedProof)
				if err := _Oracle.contract.UnpackLog(event, "VerifiedProof", log); err != nil {
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
