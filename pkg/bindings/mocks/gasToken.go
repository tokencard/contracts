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

// GasTokenABI is the input ABI used to generate the binding from.
const GasTokenABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"}],\"name\":\"computeAddress2\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"child\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"free\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"freeUpTo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalBurned\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalMinted\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// GasTokenBin is the compiled bytecode used for deploying new contracts.
var GasTokenBin = "0x608060405234801561001057600080fd5b506106a0806100206000396000f3fe608060405234801561001057600080fd5b50600436106100935760003560e01c8063a2309ff811610066578063a2309ff814610114578063a9059cbb1461011c578063b0ac19a014610148578063d89135cd14610181578063d8ccd0f31461018957610093565b806318160ddd146100985780636366b936146100b257806370a08231146100cf578063a0712d68146100f5575b600080fd5b6100a06101a6565b60408051918252519081900360200190f35b6100a0600480360360208110156100c857600080fd5b50356101b0565b6100a0600480360360208110156100e557600080fd5b50356001600160a01b03166101d2565b6101126004803603602081101561010b57600080fd5b50356101ed565b005b6100a06103c2565b6101126004803603604081101561013257600080fd5b506001600160a01b0381351690602001356103c8565b6101656004803603602081101561015e57600080fd5b503561043a565b604080516001600160a01b039092168252519081900360200190f35b6100a061049e565b6100a06004803603602081101561019f57600080fd5b50356104a4565b6001546000540390565b60006101cc6101c7836101c2336101d2565b6104c3565b6104a4565b92915050565b6001600160a01b031660009081526002602052604090205490565b600080547f746d11d3d93b3c965b59e6f453412ed064933318585733ff6000526015600bf3909152602082045b801561038d5781601e600080f55060018201601e600080f55060028201601e600080f55060038201601e600080f55060048201601e600080f55060058201601e600080f55060068201601e600080f55060078201601e600080f55060088201601e600080f55060098201601e600080f550600a8201601e600080f550600b8201601e600080f550600c8201601e600080f550600d8201601e600080f550600e8201601e600080f550600f8201601e600080f55060108201601e600080f55060118201601e600080f55060128201601e600080f55060138201601e600080f55060148201601e600080f55060158201601e600080f55060168201601e600080f55060178201601e600080f55060188201601e600080f55060198201601e600080f550601a8201601e600080f550601b8201601e600080f550601c8201601e600080f550601d8201601e600080f550601e8201601e600080f550601f8201601e600080f550602091909101906000190161021a565b50601f82165b80156103b15781601e600080f5506001919091019060001901610393565b506103bc33836104db565b60005550565b60005481565b336000908152600260205260409020546103e8908263ffffffff61050416565b33600090815260026020526040808220929092556001600160a01b0384168152205461041a908263ffffffff61056116565b6001600160a01b0390921660009081526002602052604090209190915550565b6040517fff0000000011d3d93b3c965b59e6f453412ed064930000000000000000000000815260158101919091527f66f87d73d6f32ab19f388f5b6fff433c292dc352e46467b9c0cc77fab425cb5d6035820152605590206001600160a01b031690565b60015481565b600081156104bf576104b633836105bb565b6104bf826105e4565b5090565b60008183106104d257816104d4565b825b9392505050565b6001600160a01b03821660009081526002602052604090205461041a908263ffffffff61056116565b60008282111561055b576040805162461bcd60e51b815260206004820152601e60248201527f536166654d6174683a207375627472616374696f6e206f766572666c6f770000604482015290519081900360640190fd5b50900390565b6000828201838110156104d4576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b6001600160a01b03821660009081526002602052604090205461041a908263ffffffff61050416565b600154818101806001556040517fff0000000011d3d93b3c965b59e6f453412ed06493000000000000000000000081527f66f87d73d6f32ab19f388f5b6fff433c292dc352e46467b9c0cc77fab425cb5d6035820152601581015b82841015610664578381526000806000806000605587205af15060018401935061063f565b505050505056fea265627a7a72315820873804b45e230ec7ba86b63b9b0496f0b450cd91c16f0419949b5e186657395364736f6c63430005110032"

// DeployGasToken deploys a new Ethereum contract, binding an instance of GasToken to it.
func DeployGasToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *GasToken, error) {
	parsed, err := abi.JSON(strings.NewReader(GasTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(GasTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &GasToken{GasTokenCaller: GasTokenCaller{contract: contract}, GasTokenTransactor: GasTokenTransactor{contract: contract}, GasTokenFilterer: GasTokenFilterer{contract: contract}}, nil
}

// GasToken is an auto generated Go binding around an Ethereum contract.
type GasToken struct {
	GasTokenCaller     // Read-only binding to the contract
	GasTokenTransactor // Write-only binding to the contract
	GasTokenFilterer   // Log filterer for contract events
}

// GasTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type GasTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GasTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GasTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GasTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GasTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GasTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GasTokenSession struct {
	Contract     *GasToken         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GasTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GasTokenCallerSession struct {
	Contract *GasTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// GasTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GasTokenTransactorSession struct {
	Contract     *GasTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// GasTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type GasTokenRaw struct {
	Contract *GasToken // Generic contract binding to access the raw methods on
}

// GasTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GasTokenCallerRaw struct {
	Contract *GasTokenCaller // Generic read-only contract binding to access the raw methods on
}

// GasTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GasTokenTransactorRaw struct {
	Contract *GasTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGasToken creates a new instance of GasToken, bound to a specific deployed contract.
func NewGasToken(address common.Address, backend bind.ContractBackend) (*GasToken, error) {
	contract, err := bindGasToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GasToken{GasTokenCaller: GasTokenCaller{contract: contract}, GasTokenTransactor: GasTokenTransactor{contract: contract}, GasTokenFilterer: GasTokenFilterer{contract: contract}}, nil
}

// NewGasTokenCaller creates a new read-only instance of GasToken, bound to a specific deployed contract.
func NewGasTokenCaller(address common.Address, caller bind.ContractCaller) (*GasTokenCaller, error) {
	contract, err := bindGasToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GasTokenCaller{contract: contract}, nil
}

// NewGasTokenTransactor creates a new write-only instance of GasToken, bound to a specific deployed contract.
func NewGasTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*GasTokenTransactor, error) {
	contract, err := bindGasToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GasTokenTransactor{contract: contract}, nil
}

// NewGasTokenFilterer creates a new log filterer instance of GasToken, bound to a specific deployed contract.
func NewGasTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*GasTokenFilterer, error) {
	contract, err := bindGasToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GasTokenFilterer{contract: contract}, nil
}

// bindGasToken binds a generic wrapper to an already deployed contract.
func bindGasToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GasTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GasToken *GasTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _GasToken.Contract.GasTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GasToken *GasTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GasToken.Contract.GasTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GasToken *GasTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GasToken.Contract.GasTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GasToken *GasTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _GasToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GasToken *GasTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GasToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GasToken *GasTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GasToken.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_GasToken *GasTokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _GasToken.contract.Call(opts, out, "balanceOf", account)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_GasToken *GasTokenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _GasToken.Contract.BalanceOf(&_GasToken.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_GasToken *GasTokenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _GasToken.Contract.BalanceOf(&_GasToken.CallOpts, account)
}

// ComputeAddress2 is a free data retrieval call binding the contract method 0xb0ac19a0.
//
// Solidity: function computeAddress2(uint256 salt) constant returns(address child)
func (_GasToken *GasTokenCaller) ComputeAddress2(opts *bind.CallOpts, salt *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _GasToken.contract.Call(opts, out, "computeAddress2", salt)
	return *ret0, err
}

// ComputeAddress2 is a free data retrieval call binding the contract method 0xb0ac19a0.
//
// Solidity: function computeAddress2(uint256 salt) constant returns(address child)
func (_GasToken *GasTokenSession) ComputeAddress2(salt *big.Int) (common.Address, error) {
	return _GasToken.Contract.ComputeAddress2(&_GasToken.CallOpts, salt)
}

// ComputeAddress2 is a free data retrieval call binding the contract method 0xb0ac19a0.
//
// Solidity: function computeAddress2(uint256 salt) constant returns(address child)
func (_GasToken *GasTokenCallerSession) ComputeAddress2(salt *big.Int) (common.Address, error) {
	return _GasToken.Contract.ComputeAddress2(&_GasToken.CallOpts, salt)
}

// TotalBurned is a free data retrieval call binding the contract method 0xd89135cd.
//
// Solidity: function totalBurned() constant returns(uint256)
func (_GasToken *GasTokenCaller) TotalBurned(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _GasToken.contract.Call(opts, out, "totalBurned")
	return *ret0, err
}

// TotalBurned is a free data retrieval call binding the contract method 0xd89135cd.
//
// Solidity: function totalBurned() constant returns(uint256)
func (_GasToken *GasTokenSession) TotalBurned() (*big.Int, error) {
	return _GasToken.Contract.TotalBurned(&_GasToken.CallOpts)
}

// TotalBurned is a free data retrieval call binding the contract method 0xd89135cd.
//
// Solidity: function totalBurned() constant returns(uint256)
func (_GasToken *GasTokenCallerSession) TotalBurned() (*big.Int, error) {
	return _GasToken.Contract.TotalBurned(&_GasToken.CallOpts)
}

// TotalMinted is a free data retrieval call binding the contract method 0xa2309ff8.
//
// Solidity: function totalMinted() constant returns(uint256)
func (_GasToken *GasTokenCaller) TotalMinted(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _GasToken.contract.Call(opts, out, "totalMinted")
	return *ret0, err
}

// TotalMinted is a free data retrieval call binding the contract method 0xa2309ff8.
//
// Solidity: function totalMinted() constant returns(uint256)
func (_GasToken *GasTokenSession) TotalMinted() (*big.Int, error) {
	return _GasToken.Contract.TotalMinted(&_GasToken.CallOpts)
}

// TotalMinted is a free data retrieval call binding the contract method 0xa2309ff8.
//
// Solidity: function totalMinted() constant returns(uint256)
func (_GasToken *GasTokenCallerSession) TotalMinted() (*big.Int, error) {
	return _GasToken.Contract.TotalMinted(&_GasToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_GasToken *GasTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _GasToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_GasToken *GasTokenSession) TotalSupply() (*big.Int, error) {
	return _GasToken.Contract.TotalSupply(&_GasToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_GasToken *GasTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _GasToken.Contract.TotalSupply(&_GasToken.CallOpts)
}

// Free is a paid mutator transaction binding the contract method 0xd8ccd0f3.
//
// Solidity: function free(uint256 value) returns(uint256)
func (_GasToken *GasTokenTransactor) Free(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _GasToken.contract.Transact(opts, "free", value)
}

// Free is a paid mutator transaction binding the contract method 0xd8ccd0f3.
//
// Solidity: function free(uint256 value) returns(uint256)
func (_GasToken *GasTokenSession) Free(value *big.Int) (*types.Transaction, error) {
	return _GasToken.Contract.Free(&_GasToken.TransactOpts, value)
}

// Free is a paid mutator transaction binding the contract method 0xd8ccd0f3.
//
// Solidity: function free(uint256 value) returns(uint256)
func (_GasToken *GasTokenTransactorSession) Free(value *big.Int) (*types.Transaction, error) {
	return _GasToken.Contract.Free(&_GasToken.TransactOpts, value)
}

// FreeUpTo is a paid mutator transaction binding the contract method 0x6366b936.
//
// Solidity: function freeUpTo(uint256 value) returns(uint256)
func (_GasToken *GasTokenTransactor) FreeUpTo(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _GasToken.contract.Transact(opts, "freeUpTo", value)
}

// FreeUpTo is a paid mutator transaction binding the contract method 0x6366b936.
//
// Solidity: function freeUpTo(uint256 value) returns(uint256)
func (_GasToken *GasTokenSession) FreeUpTo(value *big.Int) (*types.Transaction, error) {
	return _GasToken.Contract.FreeUpTo(&_GasToken.TransactOpts, value)
}

// FreeUpTo is a paid mutator transaction binding the contract method 0x6366b936.
//
// Solidity: function freeUpTo(uint256 value) returns(uint256)
func (_GasToken *GasTokenTransactorSession) FreeUpTo(value *big.Int) (*types.Transaction, error) {
	return _GasToken.Contract.FreeUpTo(&_GasToken.TransactOpts, value)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 value) returns()
func (_GasToken *GasTokenTransactor) Mint(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _GasToken.contract.Transact(opts, "mint", value)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 value) returns()
func (_GasToken *GasTokenSession) Mint(value *big.Int) (*types.Transaction, error) {
	return _GasToken.Contract.Mint(&_GasToken.TransactOpts, value)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 value) returns()
func (_GasToken *GasTokenTransactorSession) Mint(value *big.Int) (*types.Transaction, error) {
	return _GasToken.Contract.Mint(&_GasToken.TransactOpts, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns()
func (_GasToken *GasTokenTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _GasToken.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns()
func (_GasToken *GasTokenSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _GasToken.Contract.Transfer(&_GasToken.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns()
func (_GasToken *GasTokenTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _GasToken.Contract.Transfer(&_GasToken.TransactOpts, recipient, amount)
}
