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
const BurnerTokenABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"lockedTokenHolder\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isTransferable\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"lockTokenHolder\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[{\"name\":\"result\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"debit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"credit\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_th\",\"type\":\"address\"}],\"name\":\"setTokenHolder\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_transferable\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"burner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"TransferredOwnership\",\"type\":\"event\"}]"

// BurnerTokenBin is the compiled bytecode used for deploying new contracts.
const BurnerTokenBin = `608060405234801561001057600080fd5b50604051604080610adb833981016040818152825160209384015160008054600160a060020a031916600160a060020a03841690811760a060020a60ff0219167401000000000000000000000000000000000000000084151502178255908552948401949094528151909392849284927f850b3df64837d7d518b45f5aa64d104652c3b80eb5b34a8e3d9eb666cb7cdea59281900390910190a150505050610a1e806100bd6000396000f3006080604052600436106100c45763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166317e139a781146100c957806318160ddd146100f25780632121dc75146101195780632bbeac911461012e57806342966c6814610145578063636be27a1461015d57806370a0823114610181578063715018a6146101a25780638da5cb5b146101b7578063a9059cbb146101e8578063ef6506db1461020c578063f29d2f2814610230578063f2fde38b14610251575b600080fd5b3480156100d557600080fd5b506100de610272565b604080519115158252519081900360200190f35b3480156100fe57600080fd5b5061010761027b565b60408051918252519081900360200190f35b34801561012557600080fd5b506100de610281565b34801561013a57600080fd5b506101436102a2565b005b34801561015157600080fd5b506100de6004356102fd565b34801561016957600080fd5b50610143600160a060020a03600435166024356104cc565b34801561018d57600080fd5b50610107600160a060020a03600435166104f8565b3480156101ae57600080fd5b5061014361050a565b3480156101c357600080fd5b506101cc610633565b60408051600160a060020a039092168252519081900360200190f35b3480156101f457600080fd5b506100de600160a060020a0360043516602435610642565b34801561021857600080fd5b506100de600160a060020a03600435166024356106cc565b34801561023c57600080fd5b50610143600160a060020a03600435166106fa565b34801561025d57600080fd5b50610143600160a060020a03600435166107d0565b60045460ff1681565b60025481565b60005474010000000000000000000000000000000000000000900460ff1690565b6102aa6109aa565b15156102ee576040805160e560020a62461bcd02815260206004820152601660248201526000805160206109d3833981519152604482015290519081900360640190fd5b6004805460ff19166001179055565b600080600060025411151561031157600080fd5b60025483111561032057600080fd5b3360009081526003602052604090205483111561034057600091506104c6565b8260025481151561034d57fe5b336000908152600360205260409020549190049150610372908463ffffffff6109bb16565b33600090815260036020526040902055600254610395908463ffffffff6109bb16565b600255600154604080517f0221038a000000000000000000000000000000000000000000000000000000008152336004820152602481018490529051600160a060020a0390921691630221038a916044808201926020929091908290030181600087803b15801561040557600080fd5b505af1158015610419573d6000803e3d6000fd5b505050506040513d602081101561042f57600080fd5b5051915081151561048a576040805160e560020a62461bcd02815260206004820152600b60248201527f6275726e206661696c6564000000000000000000000000000000000000000000604482015290519081900360640190fd5b604080513381526020810185905281517fcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5929181900390910190a15b50919050565b600160a060020a0390911660009081526003602052604090208054829003905560028054919091039055565b60036020526000908152604090205481565b6105126109aa565b1515610556576040805160e560020a62461bcd02815260206004820152601660248201526000805160206109d3833981519152604482015290519081900360640190fd5b60005474010000000000000000000000000000000000000000900460ff1615156105ca576040805160e560020a62461bcd02815260206004820152601d60248201527f6f776e657273686970206973206e6f74207472616e7366657261626c65000000604482015290519081900360640190fd5b6000805460408051600160a060020a039092168252602082019290925281517f850b3df64837d7d518b45f5aa64d104652c3b80eb5b34a8e3d9eb666cb7cdea5929181900390910190a16000805473ffffffffffffffffffffffffffffffffffffffff19169055565b600054600160a060020a031690565b3360009081526003602052604081205482111561065e57600080fd5b33600081815260036020908152604080832080548790039055600160a060020a03871680845292819020805487019055805186815290519293927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef929181900390910190a350600192915050565b600160a060020a03821660009081526003602052604090208054820190556002805482019055600192915050565b6107026109aa565b1515610746576040805160e560020a62461bcd02815260206004820152601660248201526000805160206109d3833981519152604482015290519081900360640190fd5b60045460ff16156107a1576040805160e560020a62461bcd02815260206004820152601560248201527f746f6b656e486f6c646572206973206c6f636b65640000000000000000000000604482015290519081900360640190fd5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b6107d86109aa565b151561081c576040805160e560020a62461bcd02815260206004820152601660248201526000805160206109d3833981519152604482015290519081900360640190fd5b60005474010000000000000000000000000000000000000000900460ff161515610890576040805160e560020a62461bcd02815260206004820152601d60248201527f6f776e657273686970206973206e6f74207472616e7366657261626c65000000604482015290519081900360640190fd5b600160a060020a0381161515610916576040805160e560020a62461bcd02815260206004820152602360248201527f6f776e65722063616e6e6f742062652073657420746f207a65726f206164647260448201527f6573730000000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b6000805474ff000000000000000000000000000000000000000019811690915560408051600160a060020a039283168152918316602083015280517f850b3df64837d7d518b45f5aa64d104652c3b80eb5b34a8e3d9eb666cb7cdea59281900390910190a16000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600054600160a060020a0316331490565b600080838311156109cb57600080fd5b5050900390560073656e646572206973206e6f7420616e206f776e657200000000000000000000a165627a7a723058204ebf6a86216070963bb495eb0c6b485609366cd6082dbfa3e94361eb94b3e7800029`

// DeployBurnerToken deploys a new Ethereum contract, binding an instance of BurnerToken to it.
func DeployBurnerToken(auth *bind.TransactOpts, backend bind.ContractBackend, _owner common.Address, _transferable bool) (common.Address, *types.Transaction, *BurnerToken, error) {
	parsed, err := abi.JSON(strings.NewReader(BurnerTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BurnerTokenBin), backend, _owner, _transferable)
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

// IsTransferable is a free data retrieval call binding the contract method 0x2121dc75.
//
// Solidity: function isTransferable() constant returns(bool)
func (_BurnerToken *BurnerTokenCaller) IsTransferable(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BurnerToken.contract.Call(opts, out, "isTransferable")
	return *ret0, err
}

// IsTransferable is a free data retrieval call binding the contract method 0x2121dc75.
//
// Solidity: function isTransferable() constant returns(bool)
func (_BurnerToken *BurnerTokenSession) IsTransferable() (bool, error) {
	return _BurnerToken.Contract.IsTransferable(&_BurnerToken.CallOpts)
}

// IsTransferable is a free data retrieval call binding the contract method 0x2121dc75.
//
// Solidity: function isTransferable() constant returns(bool)
func (_BurnerToken *BurnerTokenCallerSession) IsTransferable() (bool, error) {
	return _BurnerToken.Contract.IsTransferable(&_BurnerToken.CallOpts)
}

// LockedTokenHolder is a free data retrieval call binding the contract method 0x17e139a7.
//
// Solidity: function lockedTokenHolder() constant returns(bool)
func (_BurnerToken *BurnerTokenCaller) LockedTokenHolder(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BurnerToken.contract.Call(opts, out, "lockedTokenHolder")
	return *ret0, err
}

// LockedTokenHolder is a free data retrieval call binding the contract method 0x17e139a7.
//
// Solidity: function lockedTokenHolder() constant returns(bool)
func (_BurnerToken *BurnerTokenSession) LockedTokenHolder() (bool, error) {
	return _BurnerToken.Contract.LockedTokenHolder(&_BurnerToken.CallOpts)
}

// LockedTokenHolder is a free data retrieval call binding the contract method 0x17e139a7.
//
// Solidity: function lockedTokenHolder() constant returns(bool)
func (_BurnerToken *BurnerTokenCallerSession) LockedTokenHolder() (bool, error) {
	return _BurnerToken.Contract.LockedTokenHolder(&_BurnerToken.CallOpts)
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

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BurnerToken *BurnerTokenTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BurnerToken.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BurnerToken *BurnerTokenSession) RenounceOwnership() (*types.Transaction, error) {
	return _BurnerToken.Contract.RenounceOwnership(&_BurnerToken.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BurnerToken *BurnerTokenTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _BurnerToken.Contract.RenounceOwnership(&_BurnerToken.TransactOpts)
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

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_account address) returns()
func (_BurnerToken *BurnerTokenTransactor) TransferOwnership(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _BurnerToken.contract.Transact(opts, "transferOwnership", _account)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_account address) returns()
func (_BurnerToken *BurnerTokenSession) TransferOwnership(_account common.Address) (*types.Transaction, error) {
	return _BurnerToken.Contract.TransferOwnership(&_BurnerToken.TransactOpts, _account)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_account address) returns()
func (_BurnerToken *BurnerTokenTransactorSession) TransferOwnership(_account common.Address) (*types.Transaction, error) {
	return _BurnerToken.Contract.TransferOwnership(&_BurnerToken.TransactOpts, _account)
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

// BurnerTokenTransferredOwnershipIterator is returned from FilterTransferredOwnership and is used to iterate over the raw logs and unpacked data for TransferredOwnership events raised by the BurnerToken contract.
type BurnerTokenTransferredOwnershipIterator struct {
	Event *BurnerTokenTransferredOwnership // Event containing the contract specifics and raw log

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
func (it *BurnerTokenTransferredOwnershipIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnerTokenTransferredOwnership)
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
		it.Event = new(BurnerTokenTransferredOwnership)
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
func (it *BurnerTokenTransferredOwnershipIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BurnerTokenTransferredOwnershipIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BurnerTokenTransferredOwnership represents a TransferredOwnership event raised by the BurnerToken contract.
type BurnerTokenTransferredOwnership struct {
	From common.Address
	To   common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterTransferredOwnership is a free log retrieval operation binding the contract event 0x850b3df64837d7d518b45f5aa64d104652c3b80eb5b34a8e3d9eb666cb7cdea5.
//
// Solidity: e TransferredOwnership(_from address, _to address)
func (_BurnerToken *BurnerTokenFilterer) FilterTransferredOwnership(opts *bind.FilterOpts) (*BurnerTokenTransferredOwnershipIterator, error) {

	logs, sub, err := _BurnerToken.contract.FilterLogs(opts, "TransferredOwnership")
	if err != nil {
		return nil, err
	}
	return &BurnerTokenTransferredOwnershipIterator{contract: _BurnerToken.contract, event: "TransferredOwnership", logs: logs, sub: sub}, nil
}

// WatchTransferredOwnership is a free log subscription operation binding the contract event 0x850b3df64837d7d518b45f5aa64d104652c3b80eb5b34a8e3d9eb666cb7cdea5.
//
// Solidity: e TransferredOwnership(_from address, _to address)
func (_BurnerToken *BurnerTokenFilterer) WatchTransferredOwnership(opts *bind.WatchOpts, sink chan<- *BurnerTokenTransferredOwnership) (event.Subscription, error) {

	logs, sub, err := _BurnerToken.contract.WatchLogs(opts, "TransferredOwnership")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BurnerTokenTransferredOwnership)
				if err := _BurnerToken.contract.UnpackLog(event, "TransferredOwnership", log); err != nil {
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
