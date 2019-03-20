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

// ProxyWalletABI is the input ABI used to generate the binding from.
const ProxyWalletABI = "[{\"inputs\":[{\"name\":\"_masterCopy\",\"type\":\"address\"},{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_transferable\",\"type\":\"bool\"},{\"name\":\"_ens\",\"type\":\"address\"},{\"name\":\"_oracleName\",\"type\":\"bytes32\"},{\"name\":\"_controllerName\",\"type\":\"bytes32\"},{\"name\":\"_licenceName\",\"type\":\"bytes32\"},{\"name\":\"_spendLimit\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"}]"

// ProxyWalletBin is the compiled bytecode used for deploying new contracts.
const ProxyWalletBin = `608060405234801561001057600080fd5b506040516101008061019383398101604090815281516020830151918301516060840151608085015160a086015160c087015160e0909701519496939492939192909190600160a060020a03881615156100f057604080517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f496e76616c6964206d617374657220636f707920616464726573732070726f7660448201527f6964656400000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b505060008054600160a060020a03909716600160a060020a0319909716969096179095555050505050606c806101276000396000f300608060405273ffffffffffffffffffffffffffffffffffffffff600054163660008037600080366000845af43d6000803e801515603b573d6000fd5b3d6000f300a165627a7a72305820de8aecb6183ffc1e631c71adab9d2009289ac47364eba695f689268f7a52d0d00029`

// DeployProxyWallet deploys a new Ethereum contract, binding an instance of ProxyWallet to it.
func DeployProxyWallet(auth *bind.TransactOpts, backend bind.ContractBackend, _masterCopy common.Address, _owner common.Address, _transferable bool, _ens common.Address, _oracleName [32]byte, _controllerName [32]byte, _licenceName [32]byte, _spendLimit *big.Int) (common.Address, *types.Transaction, *ProxyWallet, error) {
	parsed, err := abi.JSON(strings.NewReader(ProxyWalletABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ProxyWalletBin), backend, _masterCopy, _owner, _transferable, _ens, _oracleName, _controllerName, _licenceName, _spendLimit)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ProxyWallet{ProxyWalletCaller: ProxyWalletCaller{contract: contract}, ProxyWalletTransactor: ProxyWalletTransactor{contract: contract}, ProxyWalletFilterer: ProxyWalletFilterer{contract: contract}}, nil
}

// ProxyWallet is an auto generated Go binding around an Ethereum contract.
type ProxyWallet struct {
	ProxyWalletCaller     // Read-only binding to the contract
	ProxyWalletTransactor // Write-only binding to the contract
	ProxyWalletFilterer   // Log filterer for contract events
}

// ProxyWalletCaller is an auto generated read-only Go binding around an Ethereum contract.
type ProxyWalletCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProxyWalletTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ProxyWalletTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProxyWalletFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ProxyWalletFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProxyWalletSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ProxyWalletSession struct {
	Contract     *ProxyWallet      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProxyWalletCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ProxyWalletCallerSession struct {
	Contract *ProxyWalletCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// ProxyWalletTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ProxyWalletTransactorSession struct {
	Contract     *ProxyWalletTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ProxyWalletRaw is an auto generated low-level Go binding around an Ethereum contract.
type ProxyWalletRaw struct {
	Contract *ProxyWallet // Generic contract binding to access the raw methods on
}

// ProxyWalletCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ProxyWalletCallerRaw struct {
	Contract *ProxyWalletCaller // Generic read-only contract binding to access the raw methods on
}

// ProxyWalletTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ProxyWalletTransactorRaw struct {
	Contract *ProxyWalletTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProxyWallet creates a new instance of ProxyWallet, bound to a specific deployed contract.
func NewProxyWallet(address common.Address, backend bind.ContractBackend) (*ProxyWallet, error) {
	contract, err := bindProxyWallet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ProxyWallet{ProxyWalletCaller: ProxyWalletCaller{contract: contract}, ProxyWalletTransactor: ProxyWalletTransactor{contract: contract}, ProxyWalletFilterer: ProxyWalletFilterer{contract: contract}}, nil
}

// NewProxyWalletCaller creates a new read-only instance of ProxyWallet, bound to a specific deployed contract.
func NewProxyWalletCaller(address common.Address, caller bind.ContractCaller) (*ProxyWalletCaller, error) {
	contract, err := bindProxyWallet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProxyWalletCaller{contract: contract}, nil
}

// NewProxyWalletTransactor creates a new write-only instance of ProxyWallet, bound to a specific deployed contract.
func NewProxyWalletTransactor(address common.Address, transactor bind.ContractTransactor) (*ProxyWalletTransactor, error) {
	contract, err := bindProxyWallet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProxyWalletTransactor{contract: contract}, nil
}

// NewProxyWalletFilterer creates a new log filterer instance of ProxyWallet, bound to a specific deployed contract.
func NewProxyWalletFilterer(address common.Address, filterer bind.ContractFilterer) (*ProxyWalletFilterer, error) {
	contract, err := bindProxyWallet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProxyWalletFilterer{contract: contract}, nil
}

// bindProxyWallet binds a generic wrapper to an already deployed contract.
func bindProxyWallet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ProxyWalletABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProxyWallet *ProxyWalletRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ProxyWallet.Contract.ProxyWalletCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProxyWallet *ProxyWalletRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProxyWallet.Contract.ProxyWalletTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProxyWallet *ProxyWalletRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProxyWallet.Contract.ProxyWalletTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProxyWallet *ProxyWalletCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ProxyWallet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProxyWallet *ProxyWalletTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProxyWallet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProxyWallet *ProxyWalletTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProxyWallet.Contract.contract.Transact(opts, method, params...)
}
