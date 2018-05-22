// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package wallet

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

// WalletABI is the input ABI used to generate the binding from.
const WalletABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"newController\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"a\",\"type\":\"address\"}],\"name\":\"transferControl\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"addGas\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"whitelist\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"a\",\"type\":\"address\"}],\"name\":\"addWhitelistAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"balance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"a\",\"type\":\"address\"}],\"name\":\"removeWhitelistAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"newOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"acceptControl\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"name\":\"balance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"controller\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"a\",\"type\":\"address\"},{\"name\":\"c\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"}]"

// WalletBin is the compiled bytecode used for deploying new contracts.
const WalletBin = `0x608060405234801561001057600080fd5b506040516040806110cc833981018060405281019080805190602001909291908051906020019092919050505081600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050610ffd806100cf6000396000f3006080604052600436106100e6576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680633a02263c1461015d5780636d16fa41146101b457806375b01b8f146101f757806379ba50971461020e5780637ebd1b30146102255780638da5cb5b1461029257806394a7ef15146102e9578063a9059cbb1461032c578063b69ef8a814610379578063b7ecbaae146103a4578063beabacc8146103e7578063d4ee1d9014610454578063d9cba399146104ab578063e3d670d7146104c2578063f2fde38b14610519578063f77c47911461055c575b600034111561015b577fe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c3334604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019250505060405180910390a15b005b34801561016957600080fd5b506101726105b3565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156101c057600080fd5b506101f5600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506105d9565b005b34801561020357600080fd5b5061020c61067a565b005b34801561021a57600080fd5b50610223610733565b005b34801561023157600080fd5b50610250600480360381019080803590602001909291905050506107f6565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561029e57600080fd5b506102a7610834565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156102f557600080fd5b5061032a600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061085a565b005b34801561033857600080fd5b50610377600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610921565b005b34801561038557600080fd5b5061038e610a63565b6040518082815260200191505060405180910390f35b3480156103b057600080fd5b506103e5600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610a82565b005b3480156103f357600080fd5b50610452600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610b49565b005b34801561046057600080fd5b50610469610d43565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156104b757600080fd5b506104c0610d69565b005b3480156104ce57600080fd5b50610503600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610e2b565b6040518082815260200191505060405180910390f35b34801561052557600080fd5b5061055a600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610f0a565b005b34801561056857600080fd5b50610571610fac565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff168073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561063557600080fd5b81600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16148061072457508073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16145b151561072f57600080fd5b5050565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561079057600080fd5b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60048181548110151561080557fe5b906000526020600020016000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156108b757600080fd5b60048290806001815401808255809150509060018203906000526020600020016000909192909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505050565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561097e57600080fd5b600082141561098c57610a5e565b8273ffffffffffffffffffffffffffffffffffffffff166108fc839081150290604051600060405180830381858888f193505050501580156109d2573d6000803e3d6000fd5b507fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef83600084604051808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff168152602001828152602001935050505060405180910390a15b505050565b60003073ffffffffffffffffffffffffffffffffffffffff1631905090565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610adf57600080fd5b60048290806001815401808255809150509060018203906000526020600020016000909192909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505050565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610ba657600080fd5b6000821415610bb457610d3d565b8273ffffffffffffffffffffffffffffffffffffffff1663a9059cbb85846040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b158015610c5757600080fd5b505af1158015610c6b573d6000803e3d6000fd5b505050506040513d6020811015610c8157600080fd5b81019080805190602001909291905050501515610c9d57600080fd5b7fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef848484604051808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001828152602001935050505060405180910390a15b50505050565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610dc657600080fd5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60008173ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050602060405180830381600087803b158015610ec857600080fd5b505af1158015610edc573d6000803e3d6000fd5b505050506040513d6020811015610ef257600080fd5b81019080805190602001909291905050509050919050565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610f6757600080fd5b81600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff16815600a165627a7a72305820e1ff604d8257c078afc48e5fe57b477636e15633a713b6a3df81fdda3302cfae0029`

// DeployWallet deploys a new Ethereum contract, binding an instance of Wallet to it.
func DeployWallet(auth *bind.TransactOpts, backend bind.ContractBackend, a common.Address, c common.Address) (common.Address, *types.Transaction, *Wallet, error) {
	parsed, err := abi.JSON(strings.NewReader(WalletABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(WalletBin), backend, a, c)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Wallet{WalletCaller: WalletCaller{contract: contract}, WalletTransactor: WalletTransactor{contract: contract}, WalletFilterer: WalletFilterer{contract: contract}}, nil
}

// Wallet is an auto generated Go binding around an Ethereum contract.
type Wallet struct {
	WalletCaller     // Read-only binding to the contract
	WalletTransactor // Write-only binding to the contract
	WalletFilterer   // Log filterer for contract events
}

// WalletCaller is an auto generated read-only Go binding around an Ethereum contract.
type WalletCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WalletTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WalletFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WalletSession struct {
	Contract     *Wallet           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WalletCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WalletCallerSession struct {
	Contract *WalletCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// WalletTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WalletTransactorSession struct {
	Contract     *WalletTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WalletRaw is an auto generated low-level Go binding around an Ethereum contract.
type WalletRaw struct {
	Contract *Wallet // Generic contract binding to access the raw methods on
}

// WalletCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WalletCallerRaw struct {
	Contract *WalletCaller // Generic read-only contract binding to access the raw methods on
}

// WalletTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WalletTransactorRaw struct {
	Contract *WalletTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWallet creates a new instance of Wallet, bound to a specific deployed contract.
func NewWallet(address common.Address, backend bind.ContractBackend) (*Wallet, error) {
	contract, err := bindWallet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Wallet{WalletCaller: WalletCaller{contract: contract}, WalletTransactor: WalletTransactor{contract: contract}, WalletFilterer: WalletFilterer{contract: contract}}, nil
}

// NewWalletCaller creates a new read-only instance of Wallet, bound to a specific deployed contract.
func NewWalletCaller(address common.Address, caller bind.ContractCaller) (*WalletCaller, error) {
	contract, err := bindWallet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WalletCaller{contract: contract}, nil
}

// NewWalletTransactor creates a new write-only instance of Wallet, bound to a specific deployed contract.
func NewWalletTransactor(address common.Address, transactor bind.ContractTransactor) (*WalletTransactor, error) {
	contract, err := bindWallet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WalletTransactor{contract: contract}, nil
}

// NewWalletFilterer creates a new log filterer instance of Wallet, bound to a specific deployed contract.
func NewWalletFilterer(address common.Address, filterer bind.ContractFilterer) (*WalletFilterer, error) {
	contract, err := bindWallet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WalletFilterer{contract: contract}, nil
}

// bindWallet binds a generic wrapper to an already deployed contract.
func bindWallet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(WalletABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Wallet *WalletRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Wallet.Contract.WalletCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Wallet *WalletRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Wallet.Contract.WalletTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Wallet *WalletRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Wallet.Contract.WalletTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Wallet *WalletCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Wallet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Wallet *WalletTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Wallet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Wallet *WalletTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Wallet.Contract.contract.Transact(opts, method, params...)
}

// Balance is a free data retrieval call binding the contract method 0xe3d670d7.
//
// Solidity: function balance(token address) constant returns(uint256)
func (_Wallet *WalletCaller) Balance(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Wallet.contract.Call(opts, out, "balance", token)
	return *ret0, err
}

// Balance is a free data retrieval call binding the contract method 0xe3d670d7.
//
// Solidity: function balance(token address) constant returns(uint256)
func (_Wallet *WalletSession) Balance(token common.Address) (*big.Int, error) {
	return _Wallet.Contract.Balance(&_Wallet.CallOpts, token)
}

// Balance is a free data retrieval call binding the contract method 0xe3d670d7.
//
// Solidity: function balance(token address) constant returns(uint256)
func (_Wallet *WalletCallerSession) Balance(token common.Address) (*big.Int, error) {
	return _Wallet.Contract.Balance(&_Wallet.CallOpts, token)
}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() constant returns(address)
func (_Wallet *WalletCaller) Controller(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Wallet.contract.Call(opts, out, "controller")
	return *ret0, err
}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() constant returns(address)
func (_Wallet *WalletSession) Controller() (common.Address, error) {
	return _Wallet.Contract.Controller(&_Wallet.CallOpts)
}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() constant returns(address)
func (_Wallet *WalletCallerSession) Controller() (common.Address, error) {
	return _Wallet.Contract.Controller(&_Wallet.CallOpts)
}

// NewController is a free data retrieval call binding the contract method 0x3a02263c.
//
// Solidity: function newController() constant returns(address)
func (_Wallet *WalletCaller) NewController(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Wallet.contract.Call(opts, out, "newController")
	return *ret0, err
}

// NewController is a free data retrieval call binding the contract method 0x3a02263c.
//
// Solidity: function newController() constant returns(address)
func (_Wallet *WalletSession) NewController() (common.Address, error) {
	return _Wallet.Contract.NewController(&_Wallet.CallOpts)
}

// NewController is a free data retrieval call binding the contract method 0x3a02263c.
//
// Solidity: function newController() constant returns(address)
func (_Wallet *WalletCallerSession) NewController() (common.Address, error) {
	return _Wallet.Contract.NewController(&_Wallet.CallOpts)
}

// NewOwner is a free data retrieval call binding the contract method 0xd4ee1d90.
//
// Solidity: function newOwner() constant returns(address)
func (_Wallet *WalletCaller) NewOwner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Wallet.contract.Call(opts, out, "newOwner")
	return *ret0, err
}

// NewOwner is a free data retrieval call binding the contract method 0xd4ee1d90.
//
// Solidity: function newOwner() constant returns(address)
func (_Wallet *WalletSession) NewOwner() (common.Address, error) {
	return _Wallet.Contract.NewOwner(&_Wallet.CallOpts)
}

// NewOwner is a free data retrieval call binding the contract method 0xd4ee1d90.
//
// Solidity: function newOwner() constant returns(address)
func (_Wallet *WalletCallerSession) NewOwner() (common.Address, error) {
	return _Wallet.Contract.NewOwner(&_Wallet.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Wallet *WalletCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Wallet.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Wallet *WalletSession) Owner() (common.Address, error) {
	return _Wallet.Contract.Owner(&_Wallet.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Wallet *WalletCallerSession) Owner() (common.Address, error) {
	return _Wallet.Contract.Owner(&_Wallet.CallOpts)
}

// Whitelist is a free data retrieval call binding the contract method 0x7ebd1b30.
//
// Solidity: function whitelist( uint256) constant returns(address)
func (_Wallet *WalletCaller) Whitelist(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Wallet.contract.Call(opts, out, "whitelist", arg0)
	return *ret0, err
}

// Whitelist is a free data retrieval call binding the contract method 0x7ebd1b30.
//
// Solidity: function whitelist( uint256) constant returns(address)
func (_Wallet *WalletSession) Whitelist(arg0 *big.Int) (common.Address, error) {
	return _Wallet.Contract.Whitelist(&_Wallet.CallOpts, arg0)
}

// Whitelist is a free data retrieval call binding the contract method 0x7ebd1b30.
//
// Solidity: function whitelist( uint256) constant returns(address)
func (_Wallet *WalletCallerSession) Whitelist(arg0 *big.Int) (common.Address, error) {
	return _Wallet.Contract.Whitelist(&_Wallet.CallOpts, arg0)
}

// AcceptControl is a paid mutator transaction binding the contract method 0xd9cba399.
//
// Solidity: function acceptControl() returns()
func (_Wallet *WalletTransactor) AcceptControl(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "acceptControl")
}

// AcceptControl is a paid mutator transaction binding the contract method 0xd9cba399.
//
// Solidity: function acceptControl() returns()
func (_Wallet *WalletSession) AcceptControl() (*types.Transaction, error) {
	return _Wallet.Contract.AcceptControl(&_Wallet.TransactOpts)
}

// AcceptControl is a paid mutator transaction binding the contract method 0xd9cba399.
//
// Solidity: function acceptControl() returns()
func (_Wallet *WalletTransactorSession) AcceptControl() (*types.Transaction, error) {
	return _Wallet.Contract.AcceptControl(&_Wallet.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Wallet *WalletTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Wallet *WalletSession) AcceptOwnership() (*types.Transaction, error) {
	return _Wallet.Contract.AcceptOwnership(&_Wallet.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Wallet *WalletTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _Wallet.Contract.AcceptOwnership(&_Wallet.TransactOpts)
}

// AddGas is a paid mutator transaction binding the contract method 0x75b01b8f.
//
// Solidity: function addGas() returns()
func (_Wallet *WalletTransactor) AddGas(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "addGas")
}

// AddGas is a paid mutator transaction binding the contract method 0x75b01b8f.
//
// Solidity: function addGas() returns()
func (_Wallet *WalletSession) AddGas() (*types.Transaction, error) {
	return _Wallet.Contract.AddGas(&_Wallet.TransactOpts)
}

// AddGas is a paid mutator transaction binding the contract method 0x75b01b8f.
//
// Solidity: function addGas() returns()
func (_Wallet *WalletTransactorSession) AddGas() (*types.Transaction, error) {
	return _Wallet.Contract.AddGas(&_Wallet.TransactOpts)
}

// AddWhitelistAddress is a paid mutator transaction binding the contract method 0x94a7ef15.
//
// Solidity: function addWhitelistAddress(a address) returns()
func (_Wallet *WalletTransactor) AddWhitelistAddress(opts *bind.TransactOpts, a common.Address) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "addWhitelistAddress", a)
}

// AddWhitelistAddress is a paid mutator transaction binding the contract method 0x94a7ef15.
//
// Solidity: function addWhitelistAddress(a address) returns()
func (_Wallet *WalletSession) AddWhitelistAddress(a common.Address) (*types.Transaction, error) {
	return _Wallet.Contract.AddWhitelistAddress(&_Wallet.TransactOpts, a)
}

// AddWhitelistAddress is a paid mutator transaction binding the contract method 0x94a7ef15.
//
// Solidity: function addWhitelistAddress(a address) returns()
func (_Wallet *WalletTransactorSession) AddWhitelistAddress(a common.Address) (*types.Transaction, error) {
	return _Wallet.Contract.AddWhitelistAddress(&_Wallet.TransactOpts, a)
}

// RemoveWhitelistAddress is a paid mutator transaction binding the contract method 0xb7ecbaae.
//
// Solidity: function removeWhitelistAddress(a address) returns()
func (_Wallet *WalletTransactor) RemoveWhitelistAddress(opts *bind.TransactOpts, a common.Address) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "removeWhitelistAddress", a)
}

// RemoveWhitelistAddress is a paid mutator transaction binding the contract method 0xb7ecbaae.
//
// Solidity: function removeWhitelistAddress(a address) returns()
func (_Wallet *WalletSession) RemoveWhitelistAddress(a common.Address) (*types.Transaction, error) {
	return _Wallet.Contract.RemoveWhitelistAddress(&_Wallet.TransactOpts, a)
}

// RemoveWhitelistAddress is a paid mutator transaction binding the contract method 0xb7ecbaae.
//
// Solidity: function removeWhitelistAddress(a address) returns()
func (_Wallet *WalletTransactorSession) RemoveWhitelistAddress(a common.Address) (*types.Transaction, error) {
	return _Wallet.Contract.RemoveWhitelistAddress(&_Wallet.TransactOpts, a)
}

// Transfer is a paid mutator transaction binding the contract method 0xbeabacc8.
//
// Solidity: function transfer(to address, token address, amount uint256) returns()
func (_Wallet *WalletTransactor) Transfer(opts *bind.TransactOpts, to common.Address, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "transfer", to, token, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xbeabacc8.
//
// Solidity: function transfer(to address, token address, amount uint256) returns()
func (_Wallet *WalletSession) Transfer(to common.Address, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Wallet.Contract.Transfer(&_Wallet.TransactOpts, to, token, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xbeabacc8.
//
// Solidity: function transfer(to address, token address, amount uint256) returns()
func (_Wallet *WalletTransactorSession) Transfer(to common.Address, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Wallet.Contract.Transfer(&_Wallet.TransactOpts, to, token, amount)
}

// TransferControl is a paid mutator transaction binding the contract method 0x6d16fa41.
//
// Solidity: function transferControl(a address) returns()
func (_Wallet *WalletTransactor) TransferControl(opts *bind.TransactOpts, a common.Address) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "transferControl", a)
}

// TransferControl is a paid mutator transaction binding the contract method 0x6d16fa41.
//
// Solidity: function transferControl(a address) returns()
func (_Wallet *WalletSession) TransferControl(a common.Address) (*types.Transaction, error) {
	return _Wallet.Contract.TransferControl(&_Wallet.TransactOpts, a)
}

// TransferControl is a paid mutator transaction binding the contract method 0x6d16fa41.
//
// Solidity: function transferControl(a address) returns()
func (_Wallet *WalletTransactorSession) TransferControl(a common.Address) (*types.Transaction, error) {
	return _Wallet.Contract.TransferControl(&_Wallet.TransactOpts, a)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(account address) returns()
func (_Wallet *WalletTransactor) TransferOwnership(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "transferOwnership", account)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(account address) returns()
func (_Wallet *WalletSession) TransferOwnership(account common.Address) (*types.Transaction, error) {
	return _Wallet.Contract.TransferOwnership(&_Wallet.TransactOpts, account)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(account address) returns()
func (_Wallet *WalletTransactorSession) TransferOwnership(account common.Address) (*types.Transaction, error) {
	return _Wallet.Contract.TransferOwnership(&_Wallet.TransactOpts, account)
}

// WalletDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Wallet contract.
type WalletDepositIterator struct {
	Event *WalletDeposit // Event containing the contract specifics and raw log

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
func (it *WalletDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletDeposit)
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
		it.Event = new(WalletDeposit)
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
func (it *WalletDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletDeposit represents a Deposit event raised by the Wallet contract.
type WalletDeposit struct {
	From   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(from address, amount uint256)
func (_Wallet *WalletFilterer) FilterDeposit(opts *bind.FilterOpts) (*WalletDepositIterator, error) {

	logs, sub, err := _Wallet.contract.FilterLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return &WalletDepositIterator{contract: _Wallet.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(from address, amount uint256)
func (_Wallet *WalletFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *WalletDeposit) (event.Subscription, error) {

	logs, sub, err := _Wallet.contract.WatchLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletDeposit)
				if err := _Wallet.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// WalletTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Wallet contract.
type WalletTransferIterator struct {
	Event *WalletTransfer // Event containing the contract specifics and raw log

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
func (it *WalletTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletTransfer)
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
		it.Event = new(WalletTransfer)
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
func (it *WalletTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletTransfer represents a Transfer event raised by the Wallet contract.
type WalletTransfer struct {
	To     common.Address
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(to address, token address, amount uint256)
func (_Wallet *WalletFilterer) FilterTransfer(opts *bind.FilterOpts) (*WalletTransferIterator, error) {

	logs, sub, err := _Wallet.contract.FilterLogs(opts, "Transfer")
	if err != nil {
		return nil, err
	}
	return &WalletTransferIterator{contract: _Wallet.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(to address, token address, amount uint256)
func (_Wallet *WalletFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *WalletTransfer) (event.Subscription, error) {

	logs, sub, err := _Wallet.contract.WatchLogs(opts, "Transfer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletTransfer)
				if err := _Wallet.contract.UnpackLog(event, "Transfer", log); err != nil {
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
