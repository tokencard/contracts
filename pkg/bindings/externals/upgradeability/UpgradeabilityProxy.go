// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package upgradeability

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

// AdminUpgradeabilityProxyABI is the input ABI used to generate the binding from.
const AdminUpgradeabilityProxyABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_logic\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"}]"

// AdminUpgradeabilityProxyBin is the compiled bytecode used for deploying new contracts.
var AdminUpgradeabilityProxyBin = "0x60806040526040516103673803806103678339818101604052604081101561002657600080fd5b81516020830180516040519294929383019291908464010000000082111561004d57600080fd5b90830190602082018581111561006257600080fd5b825164010000000081118282018810171561007c57600080fd5b82525081516020918201929091019080838360005b838110156100a9578181015183820152602001610091565b50505050905090810190601f1680156100d65780820380516001836020036101000a031916815260200191505b5060408181527f656970313936372e70726f78792e696d706c656d656e746174696f6e0000000082525190819003601c01902060008051602061030c83398151915260001990910114925061012a91505057fe5b61013c826001600160e01b036101fb16565b8051156101f4576000826001600160a01b0316826040518082805190602001908083835b6020831061017f5780518252601f199092019160209182019101610160565b6001836020036101000a038019825116818451168082178552505050505050905001915050600060405180830381855af49150503d80600081146101df576040519150601f19603f3d011682016040523d82523d6000602084013e6101e4565b606091505b50509050806101f257600080fd5b505b5050610261565b61020e8161025b60201b6100621760201c565b6102495760405162461bcd60e51b815260040180806020018281038252603b81526020018061032c603b913960400191505060405180910390fd5b60008051602061030c83398151915255565b3b151590565b609d8061026f6000396000f3fe6080604052600a600c565b005b60186014601a565b603f565b565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5490565b3660008037600080366000845af43d6000803e808015605d573d6000f35b3d6000fd5b3b15159056fea265627a7a72315820e9b4dd5f8ba87fcb046690cdf759c7cd3d55b9b9bf32532973275a36a7a97a5264736f6c63430005110032360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc43616e6e6f742073657420612070726f787920696d706c656d656e746174696f6e20746f2061206e6f6e2d636f6e74726163742061646472657373"

// DeployAdminUpgradeabilityProxy deploys a new Ethereum contract, binding an instance of AdminUpgradeabilityProxy to it.
func DeployAdminUpgradeabilityProxy(auth *bind.TransactOpts, backend bind.ContractBackend, _logic common.Address, _data []byte) (common.Address, *types.Transaction, *AdminUpgradeabilityProxy, error) {
	parsed, err := abi.JSON(strings.NewReader(AdminUpgradeabilityProxyABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AdminUpgradeabilityProxyBin), backend, _logic, _data)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AdminUpgradeabilityProxy{AdminUpgradeabilityProxyCaller: AdminUpgradeabilityProxyCaller{contract: contract}, AdminUpgradeabilityProxyTransactor: AdminUpgradeabilityProxyTransactor{contract: contract}, AdminUpgradeabilityProxyFilterer: AdminUpgradeabilityProxyFilterer{contract: contract}}, nil
}

// AdminUpgradeabilityProxy is an auto generated Go binding around an Ethereum contract.
type AdminUpgradeabilityProxy struct {
	AdminUpgradeabilityProxyCaller     // Read-only binding to the contract
	AdminUpgradeabilityProxyTransactor // Write-only binding to the contract
	AdminUpgradeabilityProxyFilterer   // Log filterer for contract events
}

// AdminUpgradeabilityProxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type AdminUpgradeabilityProxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AdminUpgradeabilityProxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AdminUpgradeabilityProxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AdminUpgradeabilityProxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AdminUpgradeabilityProxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AdminUpgradeabilityProxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AdminUpgradeabilityProxySession struct {
	Contract     *AdminUpgradeabilityProxy // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// AdminUpgradeabilityProxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AdminUpgradeabilityProxyCallerSession struct {
	Contract *AdminUpgradeabilityProxyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// AdminUpgradeabilityProxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AdminUpgradeabilityProxyTransactorSession struct {
	Contract     *AdminUpgradeabilityProxyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// AdminUpgradeabilityProxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type AdminUpgradeabilityProxyRaw struct {
	Contract *AdminUpgradeabilityProxy // Generic contract binding to access the raw methods on
}

// AdminUpgradeabilityProxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AdminUpgradeabilityProxyCallerRaw struct {
	Contract *AdminUpgradeabilityProxyCaller // Generic read-only contract binding to access the raw methods on
}

// AdminUpgradeabilityProxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AdminUpgradeabilityProxyTransactorRaw struct {
	Contract *AdminUpgradeabilityProxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAdminUpgradeabilityProxy creates a new instance of AdminUpgradeabilityProxy, bound to a specific deployed contract.
func NewAdminUpgradeabilityProxy(address common.Address, backend bind.ContractBackend) (*AdminUpgradeabilityProxy, error) {
	contract, err := bindAdminUpgradeabilityProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AdminUpgradeabilityProxy{AdminUpgradeabilityProxyCaller: AdminUpgradeabilityProxyCaller{contract: contract}, AdminUpgradeabilityProxyTransactor: AdminUpgradeabilityProxyTransactor{contract: contract}, AdminUpgradeabilityProxyFilterer: AdminUpgradeabilityProxyFilterer{contract: contract}}, nil
}

// NewAdminUpgradeabilityProxyCaller creates a new read-only instance of AdminUpgradeabilityProxy, bound to a specific deployed contract.
func NewAdminUpgradeabilityProxyCaller(address common.Address, caller bind.ContractCaller) (*AdminUpgradeabilityProxyCaller, error) {
	contract, err := bindAdminUpgradeabilityProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AdminUpgradeabilityProxyCaller{contract: contract}, nil
}

// NewAdminUpgradeabilityProxyTransactor creates a new write-only instance of AdminUpgradeabilityProxy, bound to a specific deployed contract.
func NewAdminUpgradeabilityProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*AdminUpgradeabilityProxyTransactor, error) {
	contract, err := bindAdminUpgradeabilityProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AdminUpgradeabilityProxyTransactor{contract: contract}, nil
}

// NewAdminUpgradeabilityProxyFilterer creates a new log filterer instance of AdminUpgradeabilityProxy, bound to a specific deployed contract.
func NewAdminUpgradeabilityProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*AdminUpgradeabilityProxyFilterer, error) {
	contract, err := bindAdminUpgradeabilityProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AdminUpgradeabilityProxyFilterer{contract: contract}, nil
}

// bindAdminUpgradeabilityProxy binds a generic wrapper to an already deployed contract.
func bindAdminUpgradeabilityProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AdminUpgradeabilityProxyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AdminUpgradeabilityProxy *AdminUpgradeabilityProxyRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AdminUpgradeabilityProxy.Contract.AdminUpgradeabilityProxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AdminUpgradeabilityProxy *AdminUpgradeabilityProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AdminUpgradeabilityProxy.Contract.AdminUpgradeabilityProxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AdminUpgradeabilityProxy *AdminUpgradeabilityProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AdminUpgradeabilityProxy.Contract.AdminUpgradeabilityProxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AdminUpgradeabilityProxy *AdminUpgradeabilityProxyCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AdminUpgradeabilityProxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AdminUpgradeabilityProxy *AdminUpgradeabilityProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AdminUpgradeabilityProxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AdminUpgradeabilityProxy *AdminUpgradeabilityProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AdminUpgradeabilityProxy.Contract.contract.Transact(opts, method, params...)
}
