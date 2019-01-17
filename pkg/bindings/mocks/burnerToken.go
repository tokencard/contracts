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

// BurnerTokenABI is the input ABI used to generate the binding from.
const BurnerTokenABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"lockTokenHolder\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[{\"name\":\"result\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"debit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"changeOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"credit\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_th\",\"type\":\"address\"}],\"name\":\"setTokenHolder\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"burner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Burn\",\"type\":\"event\"}]"

// BurnerTokenBin is the compiled bytecode used for deploying new contracts.
const BurnerTokenBin = `608060405234801561001057600080fd5b50610604806100206000396000f3006080604052600436106100ae5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166318160ddd81146100b35780632bbeac91146100da57806342966c68146100f1578063636be27a1461011d57806370a082311461014157806379ba5097146101625780638da5cb5b14610177578063a6f9dae1146101a8578063a9059cbb146101c9578063ef6506db146101ed578063f29d2f2814610211575b600080fd5b3480156100bf57600080fd5b506100c8610232565b60408051918252519081900360200190f35b3480156100e657600080fd5b506100ef610238565b005b3480156100fd57600080fd5b50610109600435610286565b604080519115158252519081900360200190f35b34801561012957600080fd5b506100ef600160a060020a03600435166024356103c8565b34801561014d57600080fd5b506100c8600160a060020a03600435166103f4565b34801561016e57600080fd5b506100ef610406565b34801561018357600080fd5b5061018c61044b565b60408051600160a060020a039092168252519081900360200190f35b3480156101b457600080fd5b506100ef600160a060020a036004351661045a565b3480156101d557600080fd5b50610109600160a060020a03600435166024356104a0565b3480156101f957600080fd5b50610109600160a060020a036004351660243561052a565b34801561021d57600080fd5b506100ef600160a060020a0360043516610558565b60025481565b600054600160a060020a0316331461024f57600080fd5b6004805474ff0000000000000000000000000000000000000000191674010000000000000000000000000000000000000000179055565b336000908152600360205260408120548211156102a5575060006103c3565b336000908152600360205260409020546102bf90836105c6565b336000908152600360205260409020556002546102dc90836105c6565b60025560048054604080517f9dc29fac00000000000000000000000000000000000000000000000000000000815233938101939093526024830185905251600160a060020a0390911691639dc29fac9160448083019260209291908290030181600087803b15801561034d57600080fd5b505af1158015610361573d6000803e3d6000fd5b505050506040513d602081101561037757600080fd5b5051905080151561038757600080fd5b604080513381526020810184905281517fcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5929181900390910190a15b919050565b600160a060020a0390911660009081526003602052604090208054829003905560028054919091039055565b60036020526000908152604090205481565b600154600160a060020a0316331415610449576001546000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a039092169190911790555b565b600054600160a060020a031681565b600054600160a060020a0316331461047157600080fd5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b336000908152600360205260408120548211156104bc57600080fd5b33600081815260036020908152604080832080548790039055600160a060020a03871680845292819020805487019055805186815290519293927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef929181900390910190a350600192915050565b600160a060020a03821660009081526003602052604090208054820190556002805482019055600192915050565b600054600160a060020a0316331461056f57600080fd5b60045474010000000000000000000000000000000000000000900460ff161561059757600080fd5b6004805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b6000828211156105d257fe5b509003905600a165627a7a723058206135797269ad1da4e3bce17059453ca1060626076064435ac16ebf8546e2282d0029`

// DeployBurnerToken deploys a new Ethereum contract, binding an instance of BurnerToken to it.
func DeployBurnerToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BurnerToken, error) {
	parsed, err := abi.JSON(strings.NewReader(BurnerTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BurnerTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BurnerToken{BurnerTokenCaller: BurnerTokenCaller{contract: contract}, BurnerTokenTransactor: BurnerTokenTransactor{contract: contract}, BurnerTokenFilterer: BurnerTokenFilterer{contract: contract}}, nil
}

// BurnerToken is an auto generated Go binding around an Ethereum contract.
type BurnerToken struct {
	BurnerTokenCaller     // Read-only binding to the contract
	BurnerTokenTransactor // Write-only binding to the contract
	BurnerTokenFilterer   // Log filterer for contract events
}

// BurnerTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type BurnerTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BurnerTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BurnerTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BurnerTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BurnerTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BurnerTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BurnerTokenSession struct {
	Contract     *BurnerToken      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BurnerTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BurnerTokenCallerSession struct {
	Contract *BurnerTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// BurnerTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BurnerTokenTransactorSession struct {
	Contract     *BurnerTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// BurnerTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type BurnerTokenRaw struct {
	Contract *BurnerToken // Generic contract binding to access the raw methods on
}

// BurnerTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BurnerTokenCallerRaw struct {
	Contract *BurnerTokenCaller // Generic read-only contract binding to access the raw methods on
}

// BurnerTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BurnerTokenTransactorRaw struct {
	Contract *BurnerTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBurnerToken creates a new instance of BurnerToken, bound to a specific deployed contract.
func NewBurnerToken(address common.Address, backend bind.ContractBackend) (*BurnerToken, error) {
	contract, err := bindBurnerToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BurnerToken{BurnerTokenCaller: BurnerTokenCaller{contract: contract}, BurnerTokenTransactor: BurnerTokenTransactor{contract: contract}, BurnerTokenFilterer: BurnerTokenFilterer{contract: contract}}, nil
}

// NewBurnerTokenCaller creates a new read-only instance of BurnerToken, bound to a specific deployed contract.
func NewBurnerTokenCaller(address common.Address, caller bind.ContractCaller) (*BurnerTokenCaller, error) {
	contract, err := bindBurnerToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BurnerTokenCaller{contract: contract}, nil
}

// NewBurnerTokenTransactor creates a new write-only instance of BurnerToken, bound to a specific deployed contract.
func NewBurnerTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*BurnerTokenTransactor, error) {
	contract, err := bindBurnerToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BurnerTokenTransactor{contract: contract}, nil
}

// NewBurnerTokenFilterer creates a new log filterer instance of BurnerToken, bound to a specific deployed contract.
func NewBurnerTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*BurnerTokenFilterer, error) {
	contract, err := bindBurnerToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BurnerTokenFilterer{contract: contract}, nil
}

// bindBurnerToken binds a generic wrapper to an already deployed contract.
func bindBurnerToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BurnerTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BurnerToken *BurnerTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BurnerToken.Contract.BurnerTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BurnerToken *BurnerTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BurnerToken.Contract.BurnerTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BurnerToken *BurnerTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BurnerToken.Contract.BurnerTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BurnerToken *BurnerTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BurnerToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BurnerToken *BurnerTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BurnerToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BurnerToken *BurnerTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BurnerToken.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf( address) constant returns(uint256)
func (_BurnerToken *BurnerTokenCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BurnerToken.contract.Call(opts, out, "balanceOf", arg0)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf( address) constant returns(uint256)
func (_BurnerToken *BurnerTokenSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _BurnerToken.Contract.BalanceOf(&_BurnerToken.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf( address) constant returns(uint256)
func (_BurnerToken *BurnerTokenCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _BurnerToken.Contract.BalanceOf(&_BurnerToken.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BurnerToken *BurnerTokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BurnerToken.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BurnerToken *BurnerTokenSession) Owner() (common.Address, error) {
	return _BurnerToken.Contract.Owner(&_BurnerToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BurnerToken *BurnerTokenCallerSession) Owner() (common.Address, error) {
	return _BurnerToken.Contract.Owner(&_BurnerToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_BurnerToken *BurnerTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BurnerToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_BurnerToken *BurnerTokenSession) TotalSupply() (*big.Int, error) {
	return _BurnerToken.Contract.TotalSupply(&_BurnerToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_BurnerToken *BurnerTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _BurnerToken.Contract.TotalSupply(&_BurnerToken.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_BurnerToken *BurnerTokenTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BurnerToken.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_BurnerToken *BurnerTokenSession) AcceptOwnership() (*types.Transaction, error) {
	return _BurnerToken.Contract.AcceptOwnership(&_BurnerToken.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_BurnerToken *BurnerTokenTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _BurnerToken.Contract.AcceptOwnership(&_BurnerToken.TransactOpts)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(_amount uint256) returns(result bool)
func (_BurnerToken *BurnerTokenTransactor) Burn(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _BurnerToken.contract.Transact(opts, "burn", _amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(_amount uint256) returns(result bool)
func (_BurnerToken *BurnerTokenSession) Burn(_amount *big.Int) (*types.Transaction, error) {
	return _BurnerToken.Contract.Burn(&_BurnerToken.TransactOpts, _amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(_amount uint256) returns(result bool)
func (_BurnerToken *BurnerTokenTransactorSession) Burn(_amount *big.Int) (*types.Transaction, error) {
	return _BurnerToken.Contract.Burn(&_BurnerToken.TransactOpts, _amount)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(_newOwner address) returns()
func (_BurnerToken *BurnerTokenTransactor) ChangeOwner(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _BurnerToken.contract.Transact(opts, "changeOwner", _newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(_newOwner address) returns()
func (_BurnerToken *BurnerTokenSession) ChangeOwner(_newOwner common.Address) (*types.Transaction, error) {
	return _BurnerToken.Contract.ChangeOwner(&_BurnerToken.TransactOpts, _newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(_newOwner address) returns()
func (_BurnerToken *BurnerTokenTransactorSession) ChangeOwner(_newOwner common.Address) (*types.Transaction, error) {
	return _BurnerToken.Contract.ChangeOwner(&_BurnerToken.TransactOpts, _newOwner)
}

// Credit is a paid mutator transaction binding the contract method 0xef6506db.
//
// Solidity: function credit(to address, amount uint256) returns(bool)
func (_BurnerToken *BurnerTokenTransactor) Credit(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnerToken.contract.Transact(opts, "credit", to, amount)
}

// Credit is a paid mutator transaction binding the contract method 0xef6506db.
//
// Solidity: function credit(to address, amount uint256) returns(bool)
func (_BurnerToken *BurnerTokenSession) Credit(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnerToken.Contract.Credit(&_BurnerToken.TransactOpts, to, amount)
}

// Credit is a paid mutator transaction binding the contract method 0xef6506db.
//
// Solidity: function credit(to address, amount uint256) returns(bool)
func (_BurnerToken *BurnerTokenTransactorSession) Credit(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnerToken.Contract.Credit(&_BurnerToken.TransactOpts, to, amount)
}

// Debit is a paid mutator transaction binding the contract method 0x636be27a.
//
// Solidity: function debit(from address, amount uint256) returns()
func (_BurnerToken *BurnerTokenTransactor) Debit(opts *bind.TransactOpts, from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnerToken.contract.Transact(opts, "debit", from, amount)
}

// Debit is a paid mutator transaction binding the contract method 0x636be27a.
//
// Solidity: function debit(from address, amount uint256) returns()
func (_BurnerToken *BurnerTokenSession) Debit(from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnerToken.Contract.Debit(&_BurnerToken.TransactOpts, from, amount)
}

// Debit is a paid mutator transaction binding the contract method 0x636be27a.
//
// Solidity: function debit(from address, amount uint256) returns()
func (_BurnerToken *BurnerTokenTransactorSession) Debit(from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnerToken.Contract.Debit(&_BurnerToken.TransactOpts, from, amount)
}

// LockTokenHolder is a paid mutator transaction binding the contract method 0x2bbeac91.
//
// Solidity: function lockTokenHolder() returns()
func (_BurnerToken *BurnerTokenTransactor) LockTokenHolder(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BurnerToken.contract.Transact(opts, "lockTokenHolder")
}

// LockTokenHolder is a paid mutator transaction binding the contract method 0x2bbeac91.
//
// Solidity: function lockTokenHolder() returns()
func (_BurnerToken *BurnerTokenSession) LockTokenHolder() (*types.Transaction, error) {
	return _BurnerToken.Contract.LockTokenHolder(&_BurnerToken.TransactOpts)
}

// LockTokenHolder is a paid mutator transaction binding the contract method 0x2bbeac91.
//
// Solidity: function lockTokenHolder() returns()
func (_BurnerToken *BurnerTokenTransactorSession) LockTokenHolder() (*types.Transaction, error) {
	return _BurnerToken.Contract.LockTokenHolder(&_BurnerToken.TransactOpts)
}

// SetTokenHolder is a paid mutator transaction binding the contract method 0xf29d2f28.
//
// Solidity: function setTokenHolder(_th address) returns()
func (_BurnerToken *BurnerTokenTransactor) SetTokenHolder(opts *bind.TransactOpts, _th common.Address) (*types.Transaction, error) {
	return _BurnerToken.contract.Transact(opts, "setTokenHolder", _th)
}

// SetTokenHolder is a paid mutator transaction binding the contract method 0xf29d2f28.
//
// Solidity: function setTokenHolder(_th address) returns()
func (_BurnerToken *BurnerTokenSession) SetTokenHolder(_th common.Address) (*types.Transaction, error) {
	return _BurnerToken.Contract.SetTokenHolder(&_BurnerToken.TransactOpts, _th)
}

// SetTokenHolder is a paid mutator transaction binding the contract method 0xf29d2f28.
//
// Solidity: function setTokenHolder(_th address) returns()
func (_BurnerToken *BurnerTokenTransactorSession) SetTokenHolder(_th common.Address) (*types.Transaction, error) {
	return _BurnerToken.Contract.SetTokenHolder(&_BurnerToken.TransactOpts, _th)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, amount uint256) returns(bool)
func (_BurnerToken *BurnerTokenTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnerToken.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, amount uint256) returns(bool)
func (_BurnerToken *BurnerTokenSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnerToken.Contract.Transfer(&_BurnerToken.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, amount uint256) returns(bool)
func (_BurnerToken *BurnerTokenTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnerToken.Contract.Transfer(&_BurnerToken.TransactOpts, to, amount)
}

// BurnerTokenBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the BurnerToken contract.
type BurnerTokenBurnIterator struct {
	Event *BurnerTokenBurn // Event containing the contract specifics and raw log

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
func (it *BurnerTokenBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnerTokenBurn)
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
		it.Event = new(BurnerTokenBurn)
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
func (it *BurnerTokenBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BurnerTokenBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BurnerTokenBurn represents a Burn event raised by the BurnerToken contract.
type BurnerTokenBurn struct {
	Burner common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: e Burn(burner address, amount uint256)
func (_BurnerToken *BurnerTokenFilterer) FilterBurn(opts *bind.FilterOpts) (*BurnerTokenBurnIterator, error) {

	logs, sub, err := _BurnerToken.contract.FilterLogs(opts, "Burn")
	if err != nil {
		return nil, err
	}
	return &BurnerTokenBurnIterator{contract: _BurnerToken.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: e Burn(burner address, amount uint256)
func (_BurnerToken *BurnerTokenFilterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *BurnerTokenBurn) (event.Subscription, error) {

	logs, sub, err := _BurnerToken.contract.WatchLogs(opts, "Burn")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BurnerTokenBurn)
				if err := _BurnerToken.contract.UnpackLog(event, "Burn", log); err != nil {
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

// BurnerTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the BurnerToken contract.
type BurnerTokenTransferIterator struct {
	Event *BurnerTokenTransfer // Event containing the contract specifics and raw log

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
func (it *BurnerTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnerTokenTransfer)
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
		it.Event = new(BurnerTokenTransfer)
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
func (it *BurnerTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BurnerTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BurnerTokenTransfer represents a Transfer event raised by the BurnerToken contract.
type BurnerTokenTransfer struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, amount uint256)
func (_BurnerToken *BurnerTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BurnerTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BurnerToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BurnerTokenTransferIterator{contract: _BurnerToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, amount uint256)
func (_BurnerToken *BurnerTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *BurnerTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BurnerToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BurnerTokenTransfer)
				if err := _BurnerToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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
