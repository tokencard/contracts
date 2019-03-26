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

// TokenWhitelistableExporterABI is the input ABI used to generate the binding from.
const TokenWhitelistableExporterABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_a\",\"type\":\"address\"}],\"name\":\"isTokenLoadable\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_a\",\"type\":\"address\"}],\"name\":\"getTokenInfo\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"bool\"},{\"name\":\"\",\"type\":\"bool\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_a\",\"type\":\"address\"}],\"name\":\"isTokenAvailable\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getStablecoinInfo\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"bool\"},{\"name\":\"\",\"type\":\"bool\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"},{\"name\":\"_rate\",\"type\":\"uint256\"},{\"name\":\"_updateDate\",\"type\":\"uint256\"}],\"name\":\"updateTokenRate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTokenAddressArray\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_ens\",\"type\":\"address\"},{\"name\":\"_tokenWhitelistName\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// TokenWhitelistableExporterBin is the compiled bytecode used for deploying new contracts.
const TokenWhitelistableExporterBin = `608060405234801561001057600080fd5b50604051604080610b2983398101604052805160209091015160008054600160a060020a031916600160a060020a03909316929092178255600155610ace90819061005b90396000f3006080604052600436106100775763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166317e73794811461007c5780631f69565f146100b15780632ff0fcaa146101725780633efec5e914610193578063d545782e146101a8578063e81239ac146101d1575b600080fd5b34801561008857600080fd5b5061009d600160a060020a0360043516610236565b604080519115158252519081900360200190f35b3480156100bd57600080fd5b506100d2600160a060020a0360043516610247565b6040805160208082018890529181018690528415156060820152831515608082015260a0810183905260c080825288519082015287519091829160e08301918a019080838360005b8381101561013257818101518382015260200161011a565b50505050905090810190601f16801561015f5780820380516001836020036101000a031916815260200191505b5097505050505050505060405180910390f35b34801561017e57600080fd5b5061009d600160a060020a036004351661026e565b34801561019f57600080fd5b506100d2610279565b3480156101b457600080fd5b506101cf600160a060020a036004351660243560443561029e565b005b3480156101dd57600080fd5b506101e66102ae565b60408051602080825283518183015283519192839290830191858101910280838360005b8381101561022257818101518382015260200161020a565b505050509050019250505060405180910390f35b6000610241826102bd565b92915050565b6060600080600080600061025a876102d5565b949c939b5091995097509550909350915050565b6000610241826104f6565b6060600080600080600061028b61050f565b949b939a50919850965094509092509050565b6102a9838383610727565b505050565b60606102b86108b4565b905090565b6000806102c9836102d5565b50979650505050505050565b600080546001546040805160e060020a630178b8bf028152600481019290925251606093928392839283928392600160a060020a0390911691630178b8bf9160248082019260209290919082900301818787803b15801561033557600080fd5b505af1158015610349573d6000803e3d6000fd5b505050506040513d602081101561035f57600080fd5b50516001546040805160e160020a631d9dabef028152600481019290925251600160a060020a0390921691633b3b57de916024808201926020929091908290030181600087803b1580156103b257600080fd5b505af11580156103c6573d6000803e3d6000fd5b505050506040513d60208110156103dc57600080fd5b5051604080517f1f69565f000000000000000000000000000000000000000000000000000000008152600160a060020a038a8116600483015291519190921691631f69565f91602480830192600092919082900301818387803b15801561044257600080fd5b505af1158015610456573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405260c081101561047f57600080fd5b81019080805164010000000081111561049757600080fd5b820160208101848111156104aa57600080fd5b81516401000000008111828201871017156104c457600080fd5b5050602082015160408301516060840151608085015160a090950151939e929d50909b50995091975095509350505050565b600080610502836102d5565b5090979650505050505050565b600080546001546040805160e060020a630178b8bf028152600481019290925251606093928392839283928392600160a060020a0390911691630178b8bf9160248082019260209290919082900301818787803b15801561056f57600080fd5b505af1158015610583573d6000803e3d6000fd5b505050506040513d602081101561059957600080fd5b50516001546040805160e160020a631d9dabef028152600481019290925251600160a060020a0390921691633b3b57de916024808201926020929091908290030181600087803b1580156105ec57600080fd5b505af1158015610600573d6000803e3d6000fd5b505050506040513d602081101561061657600080fd5b5051604080517f3efec5e90000000000000000000000000000000000000000000000000000000081529051600160a060020a0390921691633efec5e99160048082019260009290919082900301818387803b15801561067457600080fd5b505af1158015610688573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405260c08110156106b157600080fd5b8101908080516401000000008111156106c957600080fd5b820160208101848111156106dc57600080fd5b81516401000000008111828201871017156106f657600080fd5b5050602082015160408301516060840151608085015160a090950151939d929c50909a509850919650945092505050565b600080546001546040805160e060020a630178b8bf028152600481019290925251600160a060020a0390921692630178b8bf926024808401936020939083900390910190829087803b15801561077c57600080fd5b505af1158015610790573d6000803e3d6000fd5b505050506040513d60208110156107a657600080fd5b50516001546040805160e160020a631d9dabef028152600481019290925251600160a060020a0390921691633b3b57de916024808201926020929091908290030181600087803b1580156107f957600080fd5b505af115801561080d573d6000803e3d6000fd5b505050506040513d602081101561082357600080fd5b5051604080517fd545782e000000000000000000000000000000000000000000000000000000008152600160a060020a03868116600483015260248201869052604482018590529151919092169163d545782e91606480830192600092919082900301818387803b15801561089757600080fd5b505af11580156108ab573d6000803e3d6000fd5b50505050505050565b600080546001546040805160e060020a630178b8bf028152600481019290925251606093600160a060020a0390931692630178b8bf92602480820193602093909283900390910190829087803b15801561090d57600080fd5b505af1158015610921573d6000803e3d6000fd5b505050506040513d602081101561093757600080fd5b50516001546040805160e160020a631d9dabef028152600481019290925251600160a060020a0390921691633b3b57de916024808201926020929091908290030181600087803b15801561098a57600080fd5b505af115801561099e573d6000803e3d6000fd5b505050506040513d60208110156109b457600080fd5b5051604080517fe81239ac0000000000000000000000000000000000000000000000000000000081529051600160a060020a039092169163e81239ac9160048082019260009290919082900301818387803b158015610a1257600080fd5b505af1158015610a26573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526020811015610a4f57600080fd5b810190808051640100000000811115610a6757600080fd5b82016020810184811115610a7a57600080fd5b8151856020820283011164010000000082111715610a9757600080fd5b5090945050505050905600a165627a7a723058207c819862c7635eadb8537a124c315542755de6f46e919099d284be026c3ed9910029`

// DeployTokenWhitelistableExporter deploys a new Ethereum contract, binding an instance of TokenWhitelistableExporter to it.
func DeployTokenWhitelistableExporter(auth *bind.TransactOpts, backend bind.ContractBackend, _ens common.Address, _tokenWhitelistName [32]byte) (common.Address, *types.Transaction, *TokenWhitelistableExporter, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenWhitelistableExporterABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TokenWhitelistableExporterBin), backend, _ens, _tokenWhitelistName)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TokenWhitelistableExporter{TokenWhitelistableExporterCaller: TokenWhitelistableExporterCaller{contract: contract}, TokenWhitelistableExporterTransactor: TokenWhitelistableExporterTransactor{contract: contract}, TokenWhitelistableExporterFilterer: TokenWhitelistableExporterFilterer{contract: contract}}, nil
}

// TokenWhitelistableExporter is an auto generated Go binding around an Ethereum contract.
type TokenWhitelistableExporter struct {
	TokenWhitelistableExporterCaller     // Read-only binding to the contract
	TokenWhitelistableExporterTransactor // Write-only binding to the contract
	TokenWhitelistableExporterFilterer   // Log filterer for contract events
}

// TokenWhitelistableExporterCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenWhitelistableExporterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenWhitelistableExporterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenWhitelistableExporterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenWhitelistableExporterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenWhitelistableExporterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenWhitelistableExporterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenWhitelistableExporterSession struct {
	Contract     *TokenWhitelistableExporter // Generic contract binding to set the session for
	CallOpts     bind.CallOpts               // Call options to use throughout this session
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// TokenWhitelistableExporterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenWhitelistableExporterCallerSession struct {
	Contract *TokenWhitelistableExporterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                     // Call options to use throughout this session
}

// TokenWhitelistableExporterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenWhitelistableExporterTransactorSession struct {
	Contract     *TokenWhitelistableExporterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                     // Transaction auth options to use throughout this session
}

// TokenWhitelistableExporterRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenWhitelistableExporterRaw struct {
	Contract *TokenWhitelistableExporter // Generic contract binding to access the raw methods on
}

// TokenWhitelistableExporterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenWhitelistableExporterCallerRaw struct {
	Contract *TokenWhitelistableExporterCaller // Generic read-only contract binding to access the raw methods on
}

// TokenWhitelistableExporterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenWhitelistableExporterTransactorRaw struct {
	Contract *TokenWhitelistableExporterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTokenWhitelistableExporter creates a new instance of TokenWhitelistableExporter, bound to a specific deployed contract.
func NewTokenWhitelistableExporter(address common.Address, backend bind.ContractBackend) (*TokenWhitelistableExporter, error) {
	contract, err := bindTokenWhitelistableExporter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TokenWhitelistableExporter{TokenWhitelistableExporterCaller: TokenWhitelistableExporterCaller{contract: contract}, TokenWhitelistableExporterTransactor: TokenWhitelistableExporterTransactor{contract: contract}, TokenWhitelistableExporterFilterer: TokenWhitelistableExporterFilterer{contract: contract}}, nil
}

// NewTokenWhitelistableExporterCaller creates a new read-only instance of TokenWhitelistableExporter, bound to a specific deployed contract.
func NewTokenWhitelistableExporterCaller(address common.Address, caller bind.ContractCaller) (*TokenWhitelistableExporterCaller, error) {
	contract, err := bindTokenWhitelistableExporter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenWhitelistableExporterCaller{contract: contract}, nil
}

// NewTokenWhitelistableExporterTransactor creates a new write-only instance of TokenWhitelistableExporter, bound to a specific deployed contract.
func NewTokenWhitelistableExporterTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenWhitelistableExporterTransactor, error) {
	contract, err := bindTokenWhitelistableExporter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenWhitelistableExporterTransactor{contract: contract}, nil
}

// NewTokenWhitelistableExporterFilterer creates a new log filterer instance of TokenWhitelistableExporter, bound to a specific deployed contract.
func NewTokenWhitelistableExporterFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenWhitelistableExporterFilterer, error) {
	contract, err := bindTokenWhitelistableExporter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenWhitelistableExporterFilterer{contract: contract}, nil
}

// bindTokenWhitelistableExporter binds a generic wrapper to an already deployed contract.
func bindTokenWhitelistableExporter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenWhitelistableExporterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenWhitelistableExporter *TokenWhitelistableExporterRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TokenWhitelistableExporter.Contract.TokenWhitelistableExporterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenWhitelistableExporter *TokenWhitelistableExporterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenWhitelistableExporter.Contract.TokenWhitelistableExporterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenWhitelistableExporter *TokenWhitelistableExporterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenWhitelistableExporter.Contract.TokenWhitelistableExporterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenWhitelistableExporter *TokenWhitelistableExporterCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TokenWhitelistableExporter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenWhitelistableExporter *TokenWhitelistableExporterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenWhitelistableExporter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenWhitelistableExporter *TokenWhitelistableExporterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenWhitelistableExporter.Contract.contract.Transact(opts, method, params...)
}

// GetStablecoinInfo is a free data retrieval call binding the contract method 0x3efec5e9.
//
// Solidity: function getStablecoinInfo() constant returns(string, uint256, uint256, bool, bool, uint256)
func (_TokenWhitelistableExporter *TokenWhitelistableExporterCaller) GetStablecoinInfo(opts *bind.CallOpts) (string, *big.Int, *big.Int, bool, bool, *big.Int, error) {
	var (
		ret0 = new(string)
		ret1 = new(*big.Int)
		ret2 = new(*big.Int)
		ret3 = new(bool)
		ret4 = new(bool)
		ret5 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
		ret5,
	}
	err := _TokenWhitelistableExporter.contract.Call(opts, out, "getStablecoinInfo")
	return *ret0, *ret1, *ret2, *ret3, *ret4, *ret5, err
}

// GetStablecoinInfo is a free data retrieval call binding the contract method 0x3efec5e9.
//
// Solidity: function getStablecoinInfo() constant returns(string, uint256, uint256, bool, bool, uint256)
func (_TokenWhitelistableExporter *TokenWhitelistableExporterSession) GetStablecoinInfo() (string, *big.Int, *big.Int, bool, bool, *big.Int, error) {
	return _TokenWhitelistableExporter.Contract.GetStablecoinInfo(&_TokenWhitelistableExporter.CallOpts)
}

// GetStablecoinInfo is a free data retrieval call binding the contract method 0x3efec5e9.
//
// Solidity: function getStablecoinInfo() constant returns(string, uint256, uint256, bool, bool, uint256)
func (_TokenWhitelistableExporter *TokenWhitelistableExporterCallerSession) GetStablecoinInfo() (string, *big.Int, *big.Int, bool, bool, *big.Int, error) {
	return _TokenWhitelistableExporter.Contract.GetStablecoinInfo(&_TokenWhitelistableExporter.CallOpts)
}

// GetTokenAddressArray is a free data retrieval call binding the contract method 0xe81239ac.
//
// Solidity: function getTokenAddressArray() constant returns(address[])
func (_TokenWhitelistableExporter *TokenWhitelistableExporterCaller) GetTokenAddressArray(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _TokenWhitelistableExporter.contract.Call(opts, out, "getTokenAddressArray")
	return *ret0, err
}

// GetTokenAddressArray is a free data retrieval call binding the contract method 0xe81239ac.
//
// Solidity: function getTokenAddressArray() constant returns(address[])
func (_TokenWhitelistableExporter *TokenWhitelistableExporterSession) GetTokenAddressArray() ([]common.Address, error) {
	return _TokenWhitelistableExporter.Contract.GetTokenAddressArray(&_TokenWhitelistableExporter.CallOpts)
}

// GetTokenAddressArray is a free data retrieval call binding the contract method 0xe81239ac.
//
// Solidity: function getTokenAddressArray() constant returns(address[])
func (_TokenWhitelistableExporter *TokenWhitelistableExporterCallerSession) GetTokenAddressArray() ([]common.Address, error) {
	return _TokenWhitelistableExporter.Contract.GetTokenAddressArray(&_TokenWhitelistableExporter.CallOpts)
}

// GetTokenInfo is a free data retrieval call binding the contract method 0x1f69565f.
//
// Solidity: function getTokenInfo(address _a) constant returns(string, uint256, uint256, bool, bool, uint256)
func (_TokenWhitelistableExporter *TokenWhitelistableExporterCaller) GetTokenInfo(opts *bind.CallOpts, _a common.Address) (string, *big.Int, *big.Int, bool, bool, *big.Int, error) {
	var (
		ret0 = new(string)
		ret1 = new(*big.Int)
		ret2 = new(*big.Int)
		ret3 = new(bool)
		ret4 = new(bool)
		ret5 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
		ret5,
	}
	err := _TokenWhitelistableExporter.contract.Call(opts, out, "getTokenInfo", _a)
	return *ret0, *ret1, *ret2, *ret3, *ret4, *ret5, err
}

// GetTokenInfo is a free data retrieval call binding the contract method 0x1f69565f.
//
// Solidity: function getTokenInfo(address _a) constant returns(string, uint256, uint256, bool, bool, uint256)
func (_TokenWhitelistableExporter *TokenWhitelistableExporterSession) GetTokenInfo(_a common.Address) (string, *big.Int, *big.Int, bool, bool, *big.Int, error) {
	return _TokenWhitelistableExporter.Contract.GetTokenInfo(&_TokenWhitelistableExporter.CallOpts, _a)
}

// GetTokenInfo is a free data retrieval call binding the contract method 0x1f69565f.
//
// Solidity: function getTokenInfo(address _a) constant returns(string, uint256, uint256, bool, bool, uint256)
func (_TokenWhitelistableExporter *TokenWhitelistableExporterCallerSession) GetTokenInfo(_a common.Address) (string, *big.Int, *big.Int, bool, bool, *big.Int, error) {
	return _TokenWhitelistableExporter.Contract.GetTokenInfo(&_TokenWhitelistableExporter.CallOpts, _a)
}

// IsTokenAvailable is a free data retrieval call binding the contract method 0x2ff0fcaa.
//
// Solidity: function isTokenAvailable(address _a) constant returns(bool)
func (_TokenWhitelistableExporter *TokenWhitelistableExporterCaller) IsTokenAvailable(opts *bind.CallOpts, _a common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TokenWhitelistableExporter.contract.Call(opts, out, "isTokenAvailable", _a)
	return *ret0, err
}

// IsTokenAvailable is a free data retrieval call binding the contract method 0x2ff0fcaa.
//
// Solidity: function isTokenAvailable(address _a) constant returns(bool)
func (_TokenWhitelistableExporter *TokenWhitelistableExporterSession) IsTokenAvailable(_a common.Address) (bool, error) {
	return _TokenWhitelistableExporter.Contract.IsTokenAvailable(&_TokenWhitelistableExporter.CallOpts, _a)
}

// IsTokenAvailable is a free data retrieval call binding the contract method 0x2ff0fcaa.
//
// Solidity: function isTokenAvailable(address _a) constant returns(bool)
func (_TokenWhitelistableExporter *TokenWhitelistableExporterCallerSession) IsTokenAvailable(_a common.Address) (bool, error) {
	return _TokenWhitelistableExporter.Contract.IsTokenAvailable(&_TokenWhitelistableExporter.CallOpts, _a)
}

// IsTokenLoadable is a free data retrieval call binding the contract method 0x17e73794.
//
// Solidity: function isTokenLoadable(address _a) constant returns(bool)
func (_TokenWhitelistableExporter *TokenWhitelistableExporterCaller) IsTokenLoadable(opts *bind.CallOpts, _a common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TokenWhitelistableExporter.contract.Call(opts, out, "isTokenLoadable", _a)
	return *ret0, err
}

// IsTokenLoadable is a free data retrieval call binding the contract method 0x17e73794.
//
// Solidity: function isTokenLoadable(address _a) constant returns(bool)
func (_TokenWhitelistableExporter *TokenWhitelistableExporterSession) IsTokenLoadable(_a common.Address) (bool, error) {
	return _TokenWhitelistableExporter.Contract.IsTokenLoadable(&_TokenWhitelistableExporter.CallOpts, _a)
}

// IsTokenLoadable is a free data retrieval call binding the contract method 0x17e73794.
//
// Solidity: function isTokenLoadable(address _a) constant returns(bool)
func (_TokenWhitelistableExporter *TokenWhitelistableExporterCallerSession) IsTokenLoadable(_a common.Address) (bool, error) {
	return _TokenWhitelistableExporter.Contract.IsTokenLoadable(&_TokenWhitelistableExporter.CallOpts, _a)
}

// UpdateTokenRate is a paid mutator transaction binding the contract method 0xd545782e.
//
// Solidity: function updateTokenRate(address _token, uint256 _rate, uint256 _updateDate) returns()
func (_TokenWhitelistableExporter *TokenWhitelistableExporterTransactor) UpdateTokenRate(opts *bind.TransactOpts, _token common.Address, _rate *big.Int, _updateDate *big.Int) (*types.Transaction, error) {
	return _TokenWhitelistableExporter.contract.Transact(opts, "updateTokenRate", _token, _rate, _updateDate)
}

// UpdateTokenRate is a paid mutator transaction binding the contract method 0xd545782e.
//
// Solidity: function updateTokenRate(address _token, uint256 _rate, uint256 _updateDate) returns()
func (_TokenWhitelistableExporter *TokenWhitelistableExporterSession) UpdateTokenRate(_token common.Address, _rate *big.Int, _updateDate *big.Int) (*types.Transaction, error) {
	return _TokenWhitelistableExporter.Contract.UpdateTokenRate(&_TokenWhitelistableExporter.TransactOpts, _token, _rate, _updateDate)
}

// UpdateTokenRate is a paid mutator transaction binding the contract method 0xd545782e.
//
// Solidity: function updateTokenRate(address _token, uint256 _rate, uint256 _updateDate) returns()
func (_TokenWhitelistableExporter *TokenWhitelistableExporterTransactorSession) UpdateTokenRate(_token common.Address, _rate *big.Int, _updateDate *big.Int) (*types.Transaction, error) {
	return _TokenWhitelistableExporter.Contract.UpdateTokenRate(&_TokenWhitelistableExporter.TransactOpts, _token, _rate, _updateDate)
}
