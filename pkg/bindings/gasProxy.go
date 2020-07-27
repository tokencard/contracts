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

// GasProxyABI is the input ABI used to generate the binding from.
const GasProxyABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_ens_\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_controllerNode_\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_gasTokenAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_ensRegistry\",\"type\":\"address\"}],\"name\":\"ENSSetRegistry\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_destination\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_returnData\",\"type\":\"bytes\"}],\"name\":\"ExecutedTransaction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_previous\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_new\",\"type\":\"uint256\"}],\"name\":\"SetFreeCallGasCost\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_previous\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_new\",\"type\":\"uint256\"}],\"name\":\"SetGasRefundPerUnit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_previous\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_new\",\"type\":\"address\"}],\"name\":\"SetGasTokenAddress\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"controllerNode\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ensRegistry\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_destination\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"executeTransaction\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"freeCallGasCost\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"gasRefundPerUnit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"gasToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_gasCost\",\"type\":\"uint256\"}],\"name\":\"setFreeCallGasCost\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_gasRefund\",\"type\":\"uint256\"}],\"name\":\"setGasRefundPerUnit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_gasTokenAddress\",\"type\":\"address\"}],\"name\":\"setGasToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// GasProxyBin is the compiled bytecode used for deploying new contracts.
var GasProxyBin = "0x6080604052603380546001600160a01b03199081166e0c2e074ec69a0dfb2997ba6c7d2e1e179091557f7f2ce995617d2816b426c5c8698c5ec2952f7a34bb10f38326f74933d5893697603455603580549091166d4946c0e9f43f4dee607b0ef1fa1c17905561374a60365561a0aa60375534801561007d57600080fd5b5060405162000cd538038062000cd5833981810160405260608110156100a257600080fd5b5080516020820151604090920151909190806001600160a01b0381161561012757603554604080516001600160a01b039283168152918316602083015280517fefc05aee327f08c8963b481afe3719d7322b797f06f5841b1fb61bd11fe507a29281900390910190a1603580546001600160a01b0319166001600160a01b0383161790555b5061013a836001600160e01b0361015416565b61014c826001600160e01b0361022c16565b5050506102e9565b600054610100900460ff168061017657506101766001600160e01b036102e316565b80610184575060005460ff16155b6101c05760405162461bcd60e51b815260040180806020018281038252602e81526020018062000ca7602e913960400191505060405180910390fd5b600054610100900460ff161580156101eb576000805460ff1961ff0019909116610100171660011790555b6001600160a01b0382161561021657603380546001600160a01b0319166001600160a01b0384161790555b8015610228576000805461ff00191690555b5050565b600054610100900460ff168061024e575061024e6001600160e01b036102e316565b8061025c575060005460ff16155b6102985760405162461bcd60e51b815260040180806020018281038252602e81526020018062000ca7602e913960400191505060405180910390fd5b600054610100900460ff161580156102c3576000805460ff1961ff0019909116610100171660011790555b81156102165760348290558015610228576000805461ff00191690555050565b303b1590565b6109ae80620002f96000396000f3fe608060405234801561001057600080fd5b50600436106100935760003560e01c8063838b1ba811610066578063838b1ba8146101835780638f3ed12a146101a0578063c91d59fe146101a8578063dc74a6d6146101b0578063e2b4ce97146101cd57610093565b80631e9f8aad146100985780633d5bc088146100c05780633f579f42146100da5780637d73b2311461015f575b600080fd5b6100be600480360360208110156100ae57600080fd5b50356001600160a01b03166101d5565b005b6100c8610238565b60408051918252519081900360200190f35b6100be600480360360608110156100f057600080fd5b6001600160a01b038235169160208101359181019060608101604082013564010000000081111561012057600080fd5b82018360208201111561013257600080fd5b8035906020019184600183028401116401000000008311171561015457600080fd5b50909250905061023e565b6101676104ed565b604080516001600160a01b039092168252519081900360200190f35b6100be6004803603602081101561019957600080fd5b50356104fc565b6100c861055c565b610167610562565b6100be600480360360208110156101c657600080fd5b5035610571565b6100c86105d1565b6101de336105d7565b61022c576040805162461bcd60e51b815260206004820152601a60248201527939b2b73232b91034b9903737ba10309031b7b73a3937b63632b960311b604482015290519081900360640190fd5b6102358161066b565b50565b60375490565b610247336105d7565b610295576040805162461bcd60e51b815260206004820152601a60248201527939b2b73232b91034b9903737ba10309031b7b73a3937b63632b960311b604482015290519081900360640190fd5b60005a905060006060866001600160a01b0316868686604051808383808284376040519201945060009350909150508083038185875af1925050503d80600081146102fc576040519150601f19603f3d011682016040523d82523d6000602084013e610301565b606091505b50915091508161034f576040805162461bcd60e51b8152602060048201526014602482015273195e1d195c9b985b0818d85b1b0819985a5b195960621b604482015290519081900360640190fd5b7ff77753fab406ecfff96d6ff2476c64a838fa9f6d37b1bf190f8546e395e3b613878787878560405180866001600160a01b03166001600160a01b03168152602001858152602001806020018060200183810383528686828181526020019250808284376000838201819052601f909101601f191690920185810384528651815286516020918201939188019250908190849084905b838110156103fd5781810151838201526020016103e5565b50505050905090810190601f16801561042a5780820380516001836020036101000a031916815260200191505b5097505050505050505060405180910390a15060009050601036025a836152080103019050603560009054906101000a90046001600160a01b03166001600160a01b0316636366b93660375460365484018161048257fe5b046040518263ffffffff1660e01b815260040180828152602001915050602060405180830381600087803b1580156104b957600080fd5b505af11580156104cd573d6000803e3d6000fd5b505050506040513d60208110156104e357600080fd5b5050505050505050565b6033546001600160a01b031690565b610505336105d7565b610553576040805162461bcd60e51b815260206004820152601a60248201527939b2b73232b91034b9903737ba10309031b7b73a3937b63632b960311b604482015290519081900360640190fd5b61023581610730565b60365490565b6035546001600160a01b031690565b61057a336105d7565b6105c8576040805162461bcd60e51b815260206004820152601a60248201527939b2b73232b91034b9903737ba10309031b7b73a3937b63632b960311b604482015290519081900360640190fd5b610235816107c4565b60345490565b60006105e4603454610858565b6001600160a01b031663b429afeb836040518263ffffffff1660e01b815260040180826001600160a01b03166001600160a01b0316815260200191505060206040518083038186803b15801561063957600080fd5b505afa15801561064d573d6000803e3d6000fd5b505050506040513d602081101561066357600080fd5b505192915050565b6001600160a01b0381166106c6576040805162461bcd60e51b815260206004820152601860248201527f67617320746f6b656e2061646472657373206973203078300000000000000000604482015290519081900360640190fd5b603554604080516001600160a01b039283168152918316602083015280517fefc05aee327f08c8963b481afe3719d7322b797f06f5841b1fb61bd11fe507a29281900390910190a1603580546001600160a01b0319166001600160a01b0392909216919091179055565b80610782576040805162461bcd60e51b815260206004820152601760248201527f667265652063616c6c2067617320636f73742069732030000000000000000000604482015290519081900360640190fd5b603654604080519182526020820183905280517f681a47ddfadc97a7d738fac5ae18e4c65b64823e3f7e357843ffe2d9771741d69281900390910190a1603655565b80610816576040805162461bcd60e51b815260206004820152601860248201527f67617320726566756e642070657220756e697420697320300000000000000000604482015290519081900360640190fd5b603754604080519182526020820183905280517f2614bf3c020c024b81abd430f0573d060dadfd8225586fd5a1e7749d8b1a278f9281900390910190a1603755565b6033546000906001600160a01b03166108b8576040805162461bcd60e51b815260206004820152601d60248201527f454e535265736f6c7661626c65206e6f7420696e697469616c697a6564000000604482015290519081900360640190fd5b60335460408051630178b8bf60e01b81526004810185905290516001600160a01b0390921691630178b8bf91602480820192602092909190829003018186803b15801561090457600080fd5b505afa158015610918573d6000803e3d6000fd5b505050506040513d602081101561092e57600080fd5b505160408051631d9dabef60e11b81526004810185905290516001600160a01b0390921691633b3b57de91602480820192602092909190829003018186803b15801561063957600080fdfea265627a7a72315820af48ca3ba60b85c827234bf78437797a8db535959f42201558be482604f9109364736f6c63430005110032436f6e747261637420696e7374616e63652068617320616c7265616479206265656e20696e697469616c697a6564"

// DeployGasProxy deploys a new Ethereum contract, binding an instance of GasProxy to it.
func DeployGasProxy(auth *bind.TransactOpts, backend bind.ContractBackend, _ens_ common.Address, _controllerNode_ [32]byte, _gasTokenAddress common.Address) (common.Address, *types.Transaction, *GasProxy, error) {
	parsed, err := abi.JSON(strings.NewReader(GasProxyABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(GasProxyBin), backend, _ens_, _controllerNode_, _gasTokenAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &GasProxy{GasProxyCaller: GasProxyCaller{contract: contract}, GasProxyTransactor: GasProxyTransactor{contract: contract}, GasProxyFilterer: GasProxyFilterer{contract: contract}}, nil
}

// GasProxy is an auto generated Go binding around an Ethereum contract.
type GasProxy struct {
	GasProxyCaller     // Read-only binding to the contract
	GasProxyTransactor // Write-only binding to the contract
	GasProxyFilterer   // Log filterer for contract events
}

// GasProxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type GasProxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GasProxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GasProxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GasProxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GasProxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GasProxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GasProxySession struct {
	Contract     *GasProxy         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GasProxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GasProxyCallerSession struct {
	Contract *GasProxyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// GasProxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GasProxyTransactorSession struct {
	Contract     *GasProxyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// GasProxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type GasProxyRaw struct {
	Contract *GasProxy // Generic contract binding to access the raw methods on
}

// GasProxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GasProxyCallerRaw struct {
	Contract *GasProxyCaller // Generic read-only contract binding to access the raw methods on
}

// GasProxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GasProxyTransactorRaw struct {
	Contract *GasProxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGasProxy creates a new instance of GasProxy, bound to a specific deployed contract.
func NewGasProxy(address common.Address, backend bind.ContractBackend) (*GasProxy, error) {
	contract, err := bindGasProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GasProxy{GasProxyCaller: GasProxyCaller{contract: contract}, GasProxyTransactor: GasProxyTransactor{contract: contract}, GasProxyFilterer: GasProxyFilterer{contract: contract}}, nil
}

// NewGasProxyCaller creates a new read-only instance of GasProxy, bound to a specific deployed contract.
func NewGasProxyCaller(address common.Address, caller bind.ContractCaller) (*GasProxyCaller, error) {
	contract, err := bindGasProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GasProxyCaller{contract: contract}, nil
}

// NewGasProxyTransactor creates a new write-only instance of GasProxy, bound to a specific deployed contract.
func NewGasProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*GasProxyTransactor, error) {
	contract, err := bindGasProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GasProxyTransactor{contract: contract}, nil
}

// NewGasProxyFilterer creates a new log filterer instance of GasProxy, bound to a specific deployed contract.
func NewGasProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*GasProxyFilterer, error) {
	contract, err := bindGasProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GasProxyFilterer{contract: contract}, nil
}

// bindGasProxy binds a generic wrapper to an already deployed contract.
func bindGasProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GasProxyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GasProxy *GasProxyRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _GasProxy.Contract.GasProxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GasProxy *GasProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GasProxy.Contract.GasProxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GasProxy *GasProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GasProxy.Contract.GasProxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GasProxy *GasProxyCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _GasProxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GasProxy *GasProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GasProxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GasProxy *GasProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GasProxy.Contract.contract.Transact(opts, method, params...)
}

// ControllerNode is a free data retrieval call binding the contract method 0xe2b4ce97.
//
// Solidity: function controllerNode() constant returns(bytes32)
func (_GasProxy *GasProxyCaller) ControllerNode(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _GasProxy.contract.Call(opts, out, "controllerNode")
	return *ret0, err
}

// ControllerNode is a free data retrieval call binding the contract method 0xe2b4ce97.
//
// Solidity: function controllerNode() constant returns(bytes32)
func (_GasProxy *GasProxySession) ControllerNode() ([32]byte, error) {
	return _GasProxy.Contract.ControllerNode(&_GasProxy.CallOpts)
}

// ControllerNode is a free data retrieval call binding the contract method 0xe2b4ce97.
//
// Solidity: function controllerNode() constant returns(bytes32)
func (_GasProxy *GasProxyCallerSession) ControllerNode() ([32]byte, error) {
	return _GasProxy.Contract.ControllerNode(&_GasProxy.CallOpts)
}

// EnsRegistry is a free data retrieval call binding the contract method 0x7d73b231.
//
// Solidity: function ensRegistry() constant returns(address)
func (_GasProxy *GasProxyCaller) EnsRegistry(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _GasProxy.contract.Call(opts, out, "ensRegistry")
	return *ret0, err
}

// EnsRegistry is a free data retrieval call binding the contract method 0x7d73b231.
//
// Solidity: function ensRegistry() constant returns(address)
func (_GasProxy *GasProxySession) EnsRegistry() (common.Address, error) {
	return _GasProxy.Contract.EnsRegistry(&_GasProxy.CallOpts)
}

// EnsRegistry is a free data retrieval call binding the contract method 0x7d73b231.
//
// Solidity: function ensRegistry() constant returns(address)
func (_GasProxy *GasProxyCallerSession) EnsRegistry() (common.Address, error) {
	return _GasProxy.Contract.EnsRegistry(&_GasProxy.CallOpts)
}

// FreeCallGasCost is a free data retrieval call binding the contract method 0x8f3ed12a.
//
// Solidity: function freeCallGasCost() constant returns(uint256)
func (_GasProxy *GasProxyCaller) FreeCallGasCost(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _GasProxy.contract.Call(opts, out, "freeCallGasCost")
	return *ret0, err
}

// FreeCallGasCost is a free data retrieval call binding the contract method 0x8f3ed12a.
//
// Solidity: function freeCallGasCost() constant returns(uint256)
func (_GasProxy *GasProxySession) FreeCallGasCost() (*big.Int, error) {
	return _GasProxy.Contract.FreeCallGasCost(&_GasProxy.CallOpts)
}

// FreeCallGasCost is a free data retrieval call binding the contract method 0x8f3ed12a.
//
// Solidity: function freeCallGasCost() constant returns(uint256)
func (_GasProxy *GasProxyCallerSession) FreeCallGasCost() (*big.Int, error) {
	return _GasProxy.Contract.FreeCallGasCost(&_GasProxy.CallOpts)
}

// GasRefundPerUnit is a free data retrieval call binding the contract method 0x3d5bc088.
//
// Solidity: function gasRefundPerUnit() constant returns(uint256)
func (_GasProxy *GasProxyCaller) GasRefundPerUnit(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _GasProxy.contract.Call(opts, out, "gasRefundPerUnit")
	return *ret0, err
}

// GasRefundPerUnit is a free data retrieval call binding the contract method 0x3d5bc088.
//
// Solidity: function gasRefundPerUnit() constant returns(uint256)
func (_GasProxy *GasProxySession) GasRefundPerUnit() (*big.Int, error) {
	return _GasProxy.Contract.GasRefundPerUnit(&_GasProxy.CallOpts)
}

// GasRefundPerUnit is a free data retrieval call binding the contract method 0x3d5bc088.
//
// Solidity: function gasRefundPerUnit() constant returns(uint256)
func (_GasProxy *GasProxyCallerSession) GasRefundPerUnit() (*big.Int, error) {
	return _GasProxy.Contract.GasRefundPerUnit(&_GasProxy.CallOpts)
}

// GasToken is a free data retrieval call binding the contract method 0xc91d59fe.
//
// Solidity: function gasToken() constant returns(address)
func (_GasProxy *GasProxyCaller) GasToken(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _GasProxy.contract.Call(opts, out, "gasToken")
	return *ret0, err
}

// GasToken is a free data retrieval call binding the contract method 0xc91d59fe.
//
// Solidity: function gasToken() constant returns(address)
func (_GasProxy *GasProxySession) GasToken() (common.Address, error) {
	return _GasProxy.Contract.GasToken(&_GasProxy.CallOpts)
}

// GasToken is a free data retrieval call binding the contract method 0xc91d59fe.
//
// Solidity: function gasToken() constant returns(address)
func (_GasProxy *GasProxyCallerSession) GasToken() (common.Address, error) {
	return _GasProxy.Contract.GasToken(&_GasProxy.CallOpts)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0x3f579f42.
//
// Solidity: function executeTransaction(address _destination, uint256 _value, bytes _data) returns()
func (_GasProxy *GasProxyTransactor) ExecuteTransaction(opts *bind.TransactOpts, _destination common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _GasProxy.contract.Transact(opts, "executeTransaction", _destination, _value, _data)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0x3f579f42.
//
// Solidity: function executeTransaction(address _destination, uint256 _value, bytes _data) returns()
func (_GasProxy *GasProxySession) ExecuteTransaction(_destination common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _GasProxy.Contract.ExecuteTransaction(&_GasProxy.TransactOpts, _destination, _value, _data)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0x3f579f42.
//
// Solidity: function executeTransaction(address _destination, uint256 _value, bytes _data) returns()
func (_GasProxy *GasProxyTransactorSession) ExecuteTransaction(_destination common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _GasProxy.Contract.ExecuteTransaction(&_GasProxy.TransactOpts, _destination, _value, _data)
}

// SetFreeCallGasCost is a paid mutator transaction binding the contract method 0x838b1ba8.
//
// Solidity: function setFreeCallGasCost(uint256 _gasCost) returns()
func (_GasProxy *GasProxyTransactor) SetFreeCallGasCost(opts *bind.TransactOpts, _gasCost *big.Int) (*types.Transaction, error) {
	return _GasProxy.contract.Transact(opts, "setFreeCallGasCost", _gasCost)
}

// SetFreeCallGasCost is a paid mutator transaction binding the contract method 0x838b1ba8.
//
// Solidity: function setFreeCallGasCost(uint256 _gasCost) returns()
func (_GasProxy *GasProxySession) SetFreeCallGasCost(_gasCost *big.Int) (*types.Transaction, error) {
	return _GasProxy.Contract.SetFreeCallGasCost(&_GasProxy.TransactOpts, _gasCost)
}

// SetFreeCallGasCost is a paid mutator transaction binding the contract method 0x838b1ba8.
//
// Solidity: function setFreeCallGasCost(uint256 _gasCost) returns()
func (_GasProxy *GasProxyTransactorSession) SetFreeCallGasCost(_gasCost *big.Int) (*types.Transaction, error) {
	return _GasProxy.Contract.SetFreeCallGasCost(&_GasProxy.TransactOpts, _gasCost)
}

// SetGasRefundPerUnit is a paid mutator transaction binding the contract method 0xdc74a6d6.
//
// Solidity: function setGasRefundPerUnit(uint256 _gasRefund) returns()
func (_GasProxy *GasProxyTransactor) SetGasRefundPerUnit(opts *bind.TransactOpts, _gasRefund *big.Int) (*types.Transaction, error) {
	return _GasProxy.contract.Transact(opts, "setGasRefundPerUnit", _gasRefund)
}

// SetGasRefundPerUnit is a paid mutator transaction binding the contract method 0xdc74a6d6.
//
// Solidity: function setGasRefundPerUnit(uint256 _gasRefund) returns()
func (_GasProxy *GasProxySession) SetGasRefundPerUnit(_gasRefund *big.Int) (*types.Transaction, error) {
	return _GasProxy.Contract.SetGasRefundPerUnit(&_GasProxy.TransactOpts, _gasRefund)
}

// SetGasRefundPerUnit is a paid mutator transaction binding the contract method 0xdc74a6d6.
//
// Solidity: function setGasRefundPerUnit(uint256 _gasRefund) returns()
func (_GasProxy *GasProxyTransactorSession) SetGasRefundPerUnit(_gasRefund *big.Int) (*types.Transaction, error) {
	return _GasProxy.Contract.SetGasRefundPerUnit(&_GasProxy.TransactOpts, _gasRefund)
}

// SetGasToken is a paid mutator transaction binding the contract method 0x1e9f8aad.
//
// Solidity: function setGasToken(address _gasTokenAddress) returns()
func (_GasProxy *GasProxyTransactor) SetGasToken(opts *bind.TransactOpts, _gasTokenAddress common.Address) (*types.Transaction, error) {
	return _GasProxy.contract.Transact(opts, "setGasToken", _gasTokenAddress)
}

// SetGasToken is a paid mutator transaction binding the contract method 0x1e9f8aad.
//
// Solidity: function setGasToken(address _gasTokenAddress) returns()
func (_GasProxy *GasProxySession) SetGasToken(_gasTokenAddress common.Address) (*types.Transaction, error) {
	return _GasProxy.Contract.SetGasToken(&_GasProxy.TransactOpts, _gasTokenAddress)
}

// SetGasToken is a paid mutator transaction binding the contract method 0x1e9f8aad.
//
// Solidity: function setGasToken(address _gasTokenAddress) returns()
func (_GasProxy *GasProxyTransactorSession) SetGasToken(_gasTokenAddress common.Address) (*types.Transaction, error) {
	return _GasProxy.Contract.SetGasToken(&_GasProxy.TransactOpts, _gasTokenAddress)
}

// GasProxyENSSetRegistryIterator is returned from FilterENSSetRegistry and is used to iterate over the raw logs and unpacked data for ENSSetRegistry events raised by the GasProxy contract.
type GasProxyENSSetRegistryIterator struct {
	Event *GasProxyENSSetRegistry // Event containing the contract specifics and raw log

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
func (it *GasProxyENSSetRegistryIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GasProxyENSSetRegistry)
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
		it.Event = new(GasProxyENSSetRegistry)
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
func (it *GasProxyENSSetRegistryIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GasProxyENSSetRegistryIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GasProxyENSSetRegistry represents a ENSSetRegistry event raised by the GasProxy contract.
type GasProxyENSSetRegistry struct {
	EnsRegistry common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterENSSetRegistry is a free log retrieval operation binding the contract event 0x186a5507043ec82d7cac3080e4bf22b3f781d96b198adf165a2df507bda90773.
//
// Solidity: event ENSSetRegistry(address _ensRegistry)
func (_GasProxy *GasProxyFilterer) FilterENSSetRegistry(opts *bind.FilterOpts) (*GasProxyENSSetRegistryIterator, error) {

	logs, sub, err := _GasProxy.contract.FilterLogs(opts, "ENSSetRegistry")
	if err != nil {
		return nil, err
	}
	return &GasProxyENSSetRegistryIterator{contract: _GasProxy.contract, event: "ENSSetRegistry", logs: logs, sub: sub}, nil
}

// WatchENSSetRegistry is a free log subscription operation binding the contract event 0x186a5507043ec82d7cac3080e4bf22b3f781d96b198adf165a2df507bda90773.
//
// Solidity: event ENSSetRegistry(address _ensRegistry)
func (_GasProxy *GasProxyFilterer) WatchENSSetRegistry(opts *bind.WatchOpts, sink chan<- *GasProxyENSSetRegistry) (event.Subscription, error) {

	logs, sub, err := _GasProxy.contract.WatchLogs(opts, "ENSSetRegistry")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GasProxyENSSetRegistry)
				if err := _GasProxy.contract.UnpackLog(event, "ENSSetRegistry", log); err != nil {
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

// ParseENSSetRegistry is a log parse operation binding the contract event 0x186a5507043ec82d7cac3080e4bf22b3f781d96b198adf165a2df507bda90773.
//
// Solidity: event ENSSetRegistry(address _ensRegistry)
func (_GasProxy *GasProxyFilterer) ParseENSSetRegistry(log types.Log) (*GasProxyENSSetRegistry, error) {
	event := new(GasProxyENSSetRegistry)
	if err := _GasProxy.contract.UnpackLog(event, "ENSSetRegistry", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GasProxyExecutedTransactionIterator is returned from FilterExecutedTransaction and is used to iterate over the raw logs and unpacked data for ExecutedTransaction events raised by the GasProxy contract.
type GasProxyExecutedTransactionIterator struct {
	Event *GasProxyExecutedTransaction // Event containing the contract specifics and raw log

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
func (it *GasProxyExecutedTransactionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GasProxyExecutedTransaction)
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
		it.Event = new(GasProxyExecutedTransaction)
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
func (it *GasProxyExecutedTransactionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GasProxyExecutedTransactionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GasProxyExecutedTransaction represents a ExecutedTransaction event raised by the GasProxy contract.
type GasProxyExecutedTransaction struct {
	Destination common.Address
	Value       *big.Int
	Data        []byte
	ReturnData  []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterExecutedTransaction is a free log retrieval operation binding the contract event 0xf77753fab406ecfff96d6ff2476c64a838fa9f6d37b1bf190f8546e395e3b613.
//
// Solidity: event ExecutedTransaction(address _destination, uint256 _value, bytes _data, bytes _returnData)
func (_GasProxy *GasProxyFilterer) FilterExecutedTransaction(opts *bind.FilterOpts) (*GasProxyExecutedTransactionIterator, error) {

	logs, sub, err := _GasProxy.contract.FilterLogs(opts, "ExecutedTransaction")
	if err != nil {
		return nil, err
	}
	return &GasProxyExecutedTransactionIterator{contract: _GasProxy.contract, event: "ExecutedTransaction", logs: logs, sub: sub}, nil
}

// WatchExecutedTransaction is a free log subscription operation binding the contract event 0xf77753fab406ecfff96d6ff2476c64a838fa9f6d37b1bf190f8546e395e3b613.
//
// Solidity: event ExecutedTransaction(address _destination, uint256 _value, bytes _data, bytes _returnData)
func (_GasProxy *GasProxyFilterer) WatchExecutedTransaction(opts *bind.WatchOpts, sink chan<- *GasProxyExecutedTransaction) (event.Subscription, error) {

	logs, sub, err := _GasProxy.contract.WatchLogs(opts, "ExecutedTransaction")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GasProxyExecutedTransaction)
				if err := _GasProxy.contract.UnpackLog(event, "ExecutedTransaction", log); err != nil {
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

// ParseExecutedTransaction is a log parse operation binding the contract event 0xf77753fab406ecfff96d6ff2476c64a838fa9f6d37b1bf190f8546e395e3b613.
//
// Solidity: event ExecutedTransaction(address _destination, uint256 _value, bytes _data, bytes _returnData)
func (_GasProxy *GasProxyFilterer) ParseExecutedTransaction(log types.Log) (*GasProxyExecutedTransaction, error) {
	event := new(GasProxyExecutedTransaction)
	if err := _GasProxy.contract.UnpackLog(event, "ExecutedTransaction", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GasProxySetFreeCallGasCostIterator is returned from FilterSetFreeCallGasCost and is used to iterate over the raw logs and unpacked data for SetFreeCallGasCost events raised by the GasProxy contract.
type GasProxySetFreeCallGasCostIterator struct {
	Event *GasProxySetFreeCallGasCost // Event containing the contract specifics and raw log

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
func (it *GasProxySetFreeCallGasCostIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GasProxySetFreeCallGasCost)
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
		it.Event = new(GasProxySetFreeCallGasCost)
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
func (it *GasProxySetFreeCallGasCostIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GasProxySetFreeCallGasCostIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GasProxySetFreeCallGasCost represents a SetFreeCallGasCost event raised by the GasProxy contract.
type GasProxySetFreeCallGasCost struct {
	Previous *big.Int
	New      *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSetFreeCallGasCost is a free log retrieval operation binding the contract event 0x681a47ddfadc97a7d738fac5ae18e4c65b64823e3f7e357843ffe2d9771741d6.
//
// Solidity: event SetFreeCallGasCost(uint256 _previous, uint256 _new)
func (_GasProxy *GasProxyFilterer) FilterSetFreeCallGasCost(opts *bind.FilterOpts) (*GasProxySetFreeCallGasCostIterator, error) {

	logs, sub, err := _GasProxy.contract.FilterLogs(opts, "SetFreeCallGasCost")
	if err != nil {
		return nil, err
	}
	return &GasProxySetFreeCallGasCostIterator{contract: _GasProxy.contract, event: "SetFreeCallGasCost", logs: logs, sub: sub}, nil
}

// WatchSetFreeCallGasCost is a free log subscription operation binding the contract event 0x681a47ddfadc97a7d738fac5ae18e4c65b64823e3f7e357843ffe2d9771741d6.
//
// Solidity: event SetFreeCallGasCost(uint256 _previous, uint256 _new)
func (_GasProxy *GasProxyFilterer) WatchSetFreeCallGasCost(opts *bind.WatchOpts, sink chan<- *GasProxySetFreeCallGasCost) (event.Subscription, error) {

	logs, sub, err := _GasProxy.contract.WatchLogs(opts, "SetFreeCallGasCost")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GasProxySetFreeCallGasCost)
				if err := _GasProxy.contract.UnpackLog(event, "SetFreeCallGasCost", log); err != nil {
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

// ParseSetFreeCallGasCost is a log parse operation binding the contract event 0x681a47ddfadc97a7d738fac5ae18e4c65b64823e3f7e357843ffe2d9771741d6.
//
// Solidity: event SetFreeCallGasCost(uint256 _previous, uint256 _new)
func (_GasProxy *GasProxyFilterer) ParseSetFreeCallGasCost(log types.Log) (*GasProxySetFreeCallGasCost, error) {
	event := new(GasProxySetFreeCallGasCost)
	if err := _GasProxy.contract.UnpackLog(event, "SetFreeCallGasCost", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GasProxySetGasRefundPerUnitIterator is returned from FilterSetGasRefundPerUnit and is used to iterate over the raw logs and unpacked data for SetGasRefundPerUnit events raised by the GasProxy contract.
type GasProxySetGasRefundPerUnitIterator struct {
	Event *GasProxySetGasRefundPerUnit // Event containing the contract specifics and raw log

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
func (it *GasProxySetGasRefundPerUnitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GasProxySetGasRefundPerUnit)
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
		it.Event = new(GasProxySetGasRefundPerUnit)
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
func (it *GasProxySetGasRefundPerUnitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GasProxySetGasRefundPerUnitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GasProxySetGasRefundPerUnit represents a SetGasRefundPerUnit event raised by the GasProxy contract.
type GasProxySetGasRefundPerUnit struct {
	Previous *big.Int
	New      *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSetGasRefundPerUnit is a free log retrieval operation binding the contract event 0x2614bf3c020c024b81abd430f0573d060dadfd8225586fd5a1e7749d8b1a278f.
//
// Solidity: event SetGasRefundPerUnit(uint256 _previous, uint256 _new)
func (_GasProxy *GasProxyFilterer) FilterSetGasRefundPerUnit(opts *bind.FilterOpts) (*GasProxySetGasRefundPerUnitIterator, error) {

	logs, sub, err := _GasProxy.contract.FilterLogs(opts, "SetGasRefundPerUnit")
	if err != nil {
		return nil, err
	}
	return &GasProxySetGasRefundPerUnitIterator{contract: _GasProxy.contract, event: "SetGasRefundPerUnit", logs: logs, sub: sub}, nil
}

// WatchSetGasRefundPerUnit is a free log subscription operation binding the contract event 0x2614bf3c020c024b81abd430f0573d060dadfd8225586fd5a1e7749d8b1a278f.
//
// Solidity: event SetGasRefundPerUnit(uint256 _previous, uint256 _new)
func (_GasProxy *GasProxyFilterer) WatchSetGasRefundPerUnit(opts *bind.WatchOpts, sink chan<- *GasProxySetGasRefundPerUnit) (event.Subscription, error) {

	logs, sub, err := _GasProxy.contract.WatchLogs(opts, "SetGasRefundPerUnit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GasProxySetGasRefundPerUnit)
				if err := _GasProxy.contract.UnpackLog(event, "SetGasRefundPerUnit", log); err != nil {
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

// ParseSetGasRefundPerUnit is a log parse operation binding the contract event 0x2614bf3c020c024b81abd430f0573d060dadfd8225586fd5a1e7749d8b1a278f.
//
// Solidity: event SetGasRefundPerUnit(uint256 _previous, uint256 _new)
func (_GasProxy *GasProxyFilterer) ParseSetGasRefundPerUnit(log types.Log) (*GasProxySetGasRefundPerUnit, error) {
	event := new(GasProxySetGasRefundPerUnit)
	if err := _GasProxy.contract.UnpackLog(event, "SetGasRefundPerUnit", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GasProxySetGasTokenAddressIterator is returned from FilterSetGasTokenAddress and is used to iterate over the raw logs and unpacked data for SetGasTokenAddress events raised by the GasProxy contract.
type GasProxySetGasTokenAddressIterator struct {
	Event *GasProxySetGasTokenAddress // Event containing the contract specifics and raw log

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
func (it *GasProxySetGasTokenAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GasProxySetGasTokenAddress)
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
		it.Event = new(GasProxySetGasTokenAddress)
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
func (it *GasProxySetGasTokenAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GasProxySetGasTokenAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GasProxySetGasTokenAddress represents a SetGasTokenAddress event raised by the GasProxy contract.
type GasProxySetGasTokenAddress struct {
	Previous common.Address
	New      common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSetGasTokenAddress is a free log retrieval operation binding the contract event 0xefc05aee327f08c8963b481afe3719d7322b797f06f5841b1fb61bd11fe507a2.
//
// Solidity: event SetGasTokenAddress(address _previous, address _new)
func (_GasProxy *GasProxyFilterer) FilterSetGasTokenAddress(opts *bind.FilterOpts) (*GasProxySetGasTokenAddressIterator, error) {

	logs, sub, err := _GasProxy.contract.FilterLogs(opts, "SetGasTokenAddress")
	if err != nil {
		return nil, err
	}
	return &GasProxySetGasTokenAddressIterator{contract: _GasProxy.contract, event: "SetGasTokenAddress", logs: logs, sub: sub}, nil
}

// WatchSetGasTokenAddress is a free log subscription operation binding the contract event 0xefc05aee327f08c8963b481afe3719d7322b797f06f5841b1fb61bd11fe507a2.
//
// Solidity: event SetGasTokenAddress(address _previous, address _new)
func (_GasProxy *GasProxyFilterer) WatchSetGasTokenAddress(opts *bind.WatchOpts, sink chan<- *GasProxySetGasTokenAddress) (event.Subscription, error) {

	logs, sub, err := _GasProxy.contract.WatchLogs(opts, "SetGasTokenAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GasProxySetGasTokenAddress)
				if err := _GasProxy.contract.UnpackLog(event, "SetGasTokenAddress", log); err != nil {
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

// ParseSetGasTokenAddress is a log parse operation binding the contract event 0xefc05aee327f08c8963b481afe3719d7322b797f06f5841b1fb61bd11fe507a2.
//
// Solidity: event SetGasTokenAddress(address _previous, address _new)
func (_GasProxy *GasProxyFilterer) ParseSetGasTokenAddress(log types.Log) (*GasProxySetGasTokenAddress, error) {
	event := new(GasProxySetGasTokenAddress)
	if err := _GasProxy.contract.UnpackLog(event, "SetGasTokenAddress", log); err != nil {
		return nil, err
	}
	return event, nil
}
