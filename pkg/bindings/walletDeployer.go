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

// WalletDeployerABI is the input ABI used to generate the binding from.
const WalletDeployerABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_ens_\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_controllerNode_\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_walletCacheNode_\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractWallet\",\"name\":\"_wallet\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"DeployedWallet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractWallet\",\"name\":\"_wallet\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractWallet\",\"name\":\"_oldWallet\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_paid\",\"type\":\"uint256\"}],\"name\":\"MigratedWallet\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"controllerNode\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"deployWallet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"deployedWallets\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ensRegistry\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"contractWallet\",\"name\":\"_oldWallet\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_initializedDailyLimit\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"_initializedWhitelist\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"_dailyLimit\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_whitelistedAddresses\",\"type\":\"address[]\"}],\"name\":\"migrateWallet\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"walletCacheNode\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// WalletDeployerBin is the compiled bytecode used for deploying new contracts.
var WalletDeployerBin = "0x60806040527f7f2ce995617d2816b426c5c8698c5ec2952f7a34bb10f38326f74933d58936976001557faf553cb0d77690819f9d6fbaa04416e1fdcfa01b2a9a833c7a11e6ae0bc1be8860025534801561005857600080fd5b50604051610a2f380380610a2f8339818101604052606081101561007b57600080fd5b5080516020820151604090920151600080546001600160a01b0319166001600160a01b0384161790559091908180156100b45760018190555b5080156100c15760028190555b50505061095c806100d36000396000f3fe6080604052600436106100555760003560e01c80634b316f7a1461005a5780637d73b231146100fc57806380a12c0e1461012d578063a6ede3d414610160578063c8cc2fc214610187578063e2b4ce97146101ba575b600080fd5b6100fa600480360360c081101561007057600080fd5b6001600160a01b038235811692602081013590911691604082013515159160608101351515916080820135919081019060c0810160a08201356401000000008111156100bb57600080fd5b8201836020820111156100cd57600080fd5b803590602001918460208302840111640100000000831117156100ef57600080fd5b5090925090506101cf565b005b34801561010857600080fd5b506101116105db565b604080516001600160a01b039092168252519081900360200190f35b34801561013957600080fd5b506101116004803603602081101561015057600080fd5b50356001600160a01b03166105ea565b34801561016c57600080fd5b50610175610605565b60408051918252519081900360200190f35b34801561019357600080fd5b506100fa600480360360208110156101aa57600080fd5b50356001600160a01b031661060b565b3480156101c657600080fd5b506101756107a9565b6101d8336107af565b610229576040805162461bcd60e51b815260206004820152601a60248201527f73656e646572206973206e6f74206120636f6e74726f6c6c6572000000000000604482015290519081900360640190fd5b6001600160a01b0387811660009081526003602052604090205416156102805760405162461bcd60e51b81526004018080602001828103825260218152602001806109066021913960400191505060405180910390fd5b866001600160a01b0316866001600160a01b0316638da5cb5b6040518163ffffffff1660e01b815260040160206040518083038186803b1580156102c357600080fd5b505afa1580156102d7573d6000803e3d6000fd5b505050506040513d60208110156102ed57600080fd5b50516001600160a01b03161461033b576040805162461bcd60e51b815260206004820152600e60248201526d0deeedccae440dad2e6dac2e8c6d60931b604482015290519081900360640190fd5b6000610348600254610843565b6001600160a01b031663a4570e516040518163ffffffff1660e01b8152600401602060405180830381600087803b15801561038257600080fd5b505af1158015610396573d6000803e3d6000fd5b505050506040513d60208110156103ac57600080fd5b5051604080516001600160a01b038084168252808b1660208301528b168183015234606082015290519192507fc65d6ee9571556236e352151c95c79b6589474ad814195aaac7d5ab8d88ba2dd919081900360800190a16001600160a01b03888116600090815260036020526040902080546001600160a01b031916918316919091179055851561049657806001600160a01b031663458d07f2856040518263ffffffff1660e01b815260040180828152602001915050600060405180830381600087803b15801561047d57600080fd5b505af1158015610491573d6000803e3d6000fd5b505050505b841561052a57806001600160a01b031663f421764884846040518363ffffffff1660e01b815260040180806020018281038252848482818152602001925060200280828437600081840152601f19601f8201169050808301925050509350505050600060405180830381600087803b15801561051157600080fd5b505af1158015610525573d6000803e3d6000fd5b505050505b60408051632c90b94d60e21b81526001600160a01b038a8116600483015260006024830181905292519084169263b242e534926044808201939182900301818387803b15801561057957600080fd5b505af115801561058d573d6000803e3d6000fd5b5050505060003411156105d1576040516001600160a01b038916903480156108fc02916000818181858888f193505050501580156105cf573d6000803e3d6000fd5b505b5050505050505050565b6000546001600160a01b031690565b6003602052600090815260409020546001600160a01b031681565b60025481565b610614336107af565b610665576040805162461bcd60e51b815260206004820152601a60248201527f73656e646572206973206e6f74206120636f6e74726f6c6c6572000000000000604482015290519081900360640190fd5b6000610672600254610843565b6001600160a01b031663a4570e516040518163ffffffff1660e01b8152600401602060405180830381600087803b1580156106ac57600080fd5b505af11580156106c0573d6000803e3d6000fd5b505050506040513d60208110156106d657600080fd5b5051604080516001600160a01b0380841682528516602082015281519293507fc02db5f4164f89d90905928336769906e16d79c4a77342126eb647ca9440d078929081900390910190a16001600160a01b0382811660008181526003602052604080822080546001600160a01b03191694861694851790558051632c90b94d60e21b81526004810193909352602483018290525163b242e53492604480820193929182900301818387803b15801561078d57600080fd5b505af11580156107a1573d6000803e3d6000fd5b505050505050565b60015490565b60006107bc600154610843565b6001600160a01b031663b429afeb836040518263ffffffff1660e01b815260040180826001600160a01b03166001600160a01b0316815260200191505060206040518083038186803b15801561081157600080fd5b505afa158015610825573d6000803e3d6000fd5b505050506040513d602081101561083b57600080fd5b505192915050565b6000805460408051630178b8bf60e01b81526004810185905290516001600160a01b0390921691630178b8bf91602480820192602092909190829003018186803b15801561089057600080fd5b505afa1580156108a4573d6000803e3d6000fd5b505050506040513d60208110156108ba57600080fd5b505160408051631d9dabef60e11b81526004810185905290516001600160a01b0390921691633b3b57de91602480820192602092909190829003018186803b15801561081157600080fdfe77616c6c657420616c7265616479206465706c6f79656420666f72206f776e6572a2646970667358221220cf0d8b07e9ce80546af2f623e1e32e849a2592969491bc3ad0289371152c2aac64736f6c63430006040033"

// DeployWalletDeployer deploys a new Ethereum contract, binding an instance of WalletDeployer to it.
func DeployWalletDeployer(auth *bind.TransactOpts, backend bind.ContractBackend, _ens_ common.Address, _controllerNode_ [32]byte, _walletCacheNode_ [32]byte) (common.Address, *types.Transaction, *WalletDeployer, error) {
	parsed, err := abi.JSON(strings.NewReader(WalletDeployerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(WalletDeployerBin), backend, _ens_, _controllerNode_, _walletCacheNode_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &WalletDeployer{WalletDeployerCaller: WalletDeployerCaller{contract: contract}, WalletDeployerTransactor: WalletDeployerTransactor{contract: contract}, WalletDeployerFilterer: WalletDeployerFilterer{contract: contract}}, nil
}

// WalletDeployer is an auto generated Go binding around an Ethereum contract.
type WalletDeployer struct {
	WalletDeployerCaller     // Read-only binding to the contract
	WalletDeployerTransactor // Write-only binding to the contract
	WalletDeployerFilterer   // Log filterer for contract events
}

// WalletDeployerCaller is an auto generated read-only Go binding around an Ethereum contract.
type WalletDeployerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletDeployerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WalletDeployerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletDeployerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WalletDeployerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletDeployerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WalletDeployerSession struct {
	Contract     *WalletDeployer   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WalletDeployerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WalletDeployerCallerSession struct {
	Contract *WalletDeployerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// WalletDeployerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WalletDeployerTransactorSession struct {
	Contract     *WalletDeployerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// WalletDeployerRaw is an auto generated low-level Go binding around an Ethereum contract.
type WalletDeployerRaw struct {
	Contract *WalletDeployer // Generic contract binding to access the raw methods on
}

// WalletDeployerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WalletDeployerCallerRaw struct {
	Contract *WalletDeployerCaller // Generic read-only contract binding to access the raw methods on
}

// WalletDeployerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WalletDeployerTransactorRaw struct {
	Contract *WalletDeployerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWalletDeployer creates a new instance of WalletDeployer, bound to a specific deployed contract.
func NewWalletDeployer(address common.Address, backend bind.ContractBackend) (*WalletDeployer, error) {
	contract, err := bindWalletDeployer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WalletDeployer{WalletDeployerCaller: WalletDeployerCaller{contract: contract}, WalletDeployerTransactor: WalletDeployerTransactor{contract: contract}, WalletDeployerFilterer: WalletDeployerFilterer{contract: contract}}, nil
}

// NewWalletDeployerCaller creates a new read-only instance of WalletDeployer, bound to a specific deployed contract.
func NewWalletDeployerCaller(address common.Address, caller bind.ContractCaller) (*WalletDeployerCaller, error) {
	contract, err := bindWalletDeployer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WalletDeployerCaller{contract: contract}, nil
}

// NewWalletDeployerTransactor creates a new write-only instance of WalletDeployer, bound to a specific deployed contract.
func NewWalletDeployerTransactor(address common.Address, transactor bind.ContractTransactor) (*WalletDeployerTransactor, error) {
	contract, err := bindWalletDeployer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WalletDeployerTransactor{contract: contract}, nil
}

// NewWalletDeployerFilterer creates a new log filterer instance of WalletDeployer, bound to a specific deployed contract.
func NewWalletDeployerFilterer(address common.Address, filterer bind.ContractFilterer) (*WalletDeployerFilterer, error) {
	contract, err := bindWalletDeployer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WalletDeployerFilterer{contract: contract}, nil
}

// bindWalletDeployer binds a generic wrapper to an already deployed contract.
func bindWalletDeployer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(WalletDeployerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WalletDeployer *WalletDeployerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _WalletDeployer.Contract.WalletDeployerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WalletDeployer *WalletDeployerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WalletDeployer.Contract.WalletDeployerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WalletDeployer *WalletDeployerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WalletDeployer.Contract.WalletDeployerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WalletDeployer *WalletDeployerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _WalletDeployer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WalletDeployer *WalletDeployerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WalletDeployer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WalletDeployer *WalletDeployerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WalletDeployer.Contract.contract.Transact(opts, method, params...)
}

// ControllerNode is a free data retrieval call binding the contract method 0xe2b4ce97.
//
// Solidity: function controllerNode() constant returns(bytes32)
func (_WalletDeployer *WalletDeployerCaller) ControllerNode(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _WalletDeployer.contract.Call(opts, out, "controllerNode")
	return *ret0, err
}

// ControllerNode is a free data retrieval call binding the contract method 0xe2b4ce97.
//
// Solidity: function controllerNode() constant returns(bytes32)
func (_WalletDeployer *WalletDeployerSession) ControllerNode() ([32]byte, error) {
	return _WalletDeployer.Contract.ControllerNode(&_WalletDeployer.CallOpts)
}

// ControllerNode is a free data retrieval call binding the contract method 0xe2b4ce97.
//
// Solidity: function controllerNode() constant returns(bytes32)
func (_WalletDeployer *WalletDeployerCallerSession) ControllerNode() ([32]byte, error) {
	return _WalletDeployer.Contract.ControllerNode(&_WalletDeployer.CallOpts)
}

// DeployedWallets is a free data retrieval call binding the contract method 0x80a12c0e.
//
// Solidity: function deployedWallets(address ) constant returns(address)
func (_WalletDeployer *WalletDeployerCaller) DeployedWallets(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _WalletDeployer.contract.Call(opts, out, "deployedWallets", arg0)
	return *ret0, err
}

// DeployedWallets is a free data retrieval call binding the contract method 0x80a12c0e.
//
// Solidity: function deployedWallets(address ) constant returns(address)
func (_WalletDeployer *WalletDeployerSession) DeployedWallets(arg0 common.Address) (common.Address, error) {
	return _WalletDeployer.Contract.DeployedWallets(&_WalletDeployer.CallOpts, arg0)
}

// DeployedWallets is a free data retrieval call binding the contract method 0x80a12c0e.
//
// Solidity: function deployedWallets(address ) constant returns(address)
func (_WalletDeployer *WalletDeployerCallerSession) DeployedWallets(arg0 common.Address) (common.Address, error) {
	return _WalletDeployer.Contract.DeployedWallets(&_WalletDeployer.CallOpts, arg0)
}

// EnsRegistry is a free data retrieval call binding the contract method 0x7d73b231.
//
// Solidity: function ensRegistry() constant returns(address)
func (_WalletDeployer *WalletDeployerCaller) EnsRegistry(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _WalletDeployer.contract.Call(opts, out, "ensRegistry")
	return *ret0, err
}

// EnsRegistry is a free data retrieval call binding the contract method 0x7d73b231.
//
// Solidity: function ensRegistry() constant returns(address)
func (_WalletDeployer *WalletDeployerSession) EnsRegistry() (common.Address, error) {
	return _WalletDeployer.Contract.EnsRegistry(&_WalletDeployer.CallOpts)
}

// EnsRegistry is a free data retrieval call binding the contract method 0x7d73b231.
//
// Solidity: function ensRegistry() constant returns(address)
func (_WalletDeployer *WalletDeployerCallerSession) EnsRegistry() (common.Address, error) {
	return _WalletDeployer.Contract.EnsRegistry(&_WalletDeployer.CallOpts)
}

// WalletCacheNode is a free data retrieval call binding the contract method 0xa6ede3d4.
//
// Solidity: function walletCacheNode() constant returns(bytes32)
func (_WalletDeployer *WalletDeployerCaller) WalletCacheNode(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _WalletDeployer.contract.Call(opts, out, "walletCacheNode")
	return *ret0, err
}

// WalletCacheNode is a free data retrieval call binding the contract method 0xa6ede3d4.
//
// Solidity: function walletCacheNode() constant returns(bytes32)
func (_WalletDeployer *WalletDeployerSession) WalletCacheNode() ([32]byte, error) {
	return _WalletDeployer.Contract.WalletCacheNode(&_WalletDeployer.CallOpts)
}

// WalletCacheNode is a free data retrieval call binding the contract method 0xa6ede3d4.
//
// Solidity: function walletCacheNode() constant returns(bytes32)
func (_WalletDeployer *WalletDeployerCallerSession) WalletCacheNode() ([32]byte, error) {
	return _WalletDeployer.Contract.WalletCacheNode(&_WalletDeployer.CallOpts)
}

// DeployWallet is a paid mutator transaction binding the contract method 0xc8cc2fc2.
//
// Solidity: function deployWallet(address _owner) returns()
func (_WalletDeployer *WalletDeployerTransactor) DeployWallet(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _WalletDeployer.contract.Transact(opts, "deployWallet", _owner)
}

// DeployWallet is a paid mutator transaction binding the contract method 0xc8cc2fc2.
//
// Solidity: function deployWallet(address _owner) returns()
func (_WalletDeployer *WalletDeployerSession) DeployWallet(_owner common.Address) (*types.Transaction, error) {
	return _WalletDeployer.Contract.DeployWallet(&_WalletDeployer.TransactOpts, _owner)
}

// DeployWallet is a paid mutator transaction binding the contract method 0xc8cc2fc2.
//
// Solidity: function deployWallet(address _owner) returns()
func (_WalletDeployer *WalletDeployerTransactorSession) DeployWallet(_owner common.Address) (*types.Transaction, error) {
	return _WalletDeployer.Contract.DeployWallet(&_WalletDeployer.TransactOpts, _owner)
}

// MigrateWallet is a paid mutator transaction binding the contract method 0x4b316f7a.
//
// Solidity: function migrateWallet(address _owner, address _oldWallet, bool _initializedDailyLimit, bool _initializedWhitelist, uint256 _dailyLimit, address[] _whitelistedAddresses) returns()
func (_WalletDeployer *WalletDeployerTransactor) MigrateWallet(opts *bind.TransactOpts, _owner common.Address, _oldWallet common.Address, _initializedDailyLimit bool, _initializedWhitelist bool, _dailyLimit *big.Int, _whitelistedAddresses []common.Address) (*types.Transaction, error) {
	return _WalletDeployer.contract.Transact(opts, "migrateWallet", _owner, _oldWallet, _initializedDailyLimit, _initializedWhitelist, _dailyLimit, _whitelistedAddresses)
}

// MigrateWallet is a paid mutator transaction binding the contract method 0x4b316f7a.
//
// Solidity: function migrateWallet(address _owner, address _oldWallet, bool _initializedDailyLimit, bool _initializedWhitelist, uint256 _dailyLimit, address[] _whitelistedAddresses) returns()
func (_WalletDeployer *WalletDeployerSession) MigrateWallet(_owner common.Address, _oldWallet common.Address, _initializedDailyLimit bool, _initializedWhitelist bool, _dailyLimit *big.Int, _whitelistedAddresses []common.Address) (*types.Transaction, error) {
	return _WalletDeployer.Contract.MigrateWallet(&_WalletDeployer.TransactOpts, _owner, _oldWallet, _initializedDailyLimit, _initializedWhitelist, _dailyLimit, _whitelistedAddresses)
}

// MigrateWallet is a paid mutator transaction binding the contract method 0x4b316f7a.
//
// Solidity: function migrateWallet(address _owner, address _oldWallet, bool _initializedDailyLimit, bool _initializedWhitelist, uint256 _dailyLimit, address[] _whitelistedAddresses) returns()
func (_WalletDeployer *WalletDeployerTransactorSession) MigrateWallet(_owner common.Address, _oldWallet common.Address, _initializedDailyLimit bool, _initializedWhitelist bool, _dailyLimit *big.Int, _whitelistedAddresses []common.Address) (*types.Transaction, error) {
	return _WalletDeployer.Contract.MigrateWallet(&_WalletDeployer.TransactOpts, _owner, _oldWallet, _initializedDailyLimit, _initializedWhitelist, _dailyLimit, _whitelistedAddresses)
}

// WalletDeployerDeployedWalletIterator is returned from FilterDeployedWallet and is used to iterate over the raw logs and unpacked data for DeployedWallet events raised by the WalletDeployer contract.
type WalletDeployerDeployedWalletIterator struct {
	Event *WalletDeployerDeployedWallet // Event containing the contract specifics and raw log

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
func (it *WalletDeployerDeployedWalletIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletDeployerDeployedWallet)
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
		it.Event = new(WalletDeployerDeployedWallet)
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
func (it *WalletDeployerDeployedWalletIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletDeployerDeployedWalletIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletDeployerDeployedWallet represents a DeployedWallet event raised by the WalletDeployer contract.
type WalletDeployerDeployedWallet struct {
	Wallet common.Address
	Owner  common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeployedWallet is a free log retrieval operation binding the contract event 0xc02db5f4164f89d90905928336769906e16d79c4a77342126eb647ca9440d078.
//
// Solidity: event DeployedWallet(address _wallet, address _owner)
func (_WalletDeployer *WalletDeployerFilterer) FilterDeployedWallet(opts *bind.FilterOpts) (*WalletDeployerDeployedWalletIterator, error) {

	logs, sub, err := _WalletDeployer.contract.FilterLogs(opts, "DeployedWallet")
	if err != nil {
		return nil, err
	}
	return &WalletDeployerDeployedWalletIterator{contract: _WalletDeployer.contract, event: "DeployedWallet", logs: logs, sub: sub}, nil
}

// WatchDeployedWallet is a free log subscription operation binding the contract event 0xc02db5f4164f89d90905928336769906e16d79c4a77342126eb647ca9440d078.
//
// Solidity: event DeployedWallet(address _wallet, address _owner)
func (_WalletDeployer *WalletDeployerFilterer) WatchDeployedWallet(opts *bind.WatchOpts, sink chan<- *WalletDeployerDeployedWallet) (event.Subscription, error) {

	logs, sub, err := _WalletDeployer.contract.WatchLogs(opts, "DeployedWallet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletDeployerDeployedWallet)
				if err := _WalletDeployer.contract.UnpackLog(event, "DeployedWallet", log); err != nil {
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

// ParseDeployedWallet is a log parse operation binding the contract event 0xc02db5f4164f89d90905928336769906e16d79c4a77342126eb647ca9440d078.
//
// Solidity: event DeployedWallet(address _wallet, address _owner)
func (_WalletDeployer *WalletDeployerFilterer) ParseDeployedWallet(log types.Log) (*WalletDeployerDeployedWallet, error) {
	event := new(WalletDeployerDeployedWallet)
	if err := _WalletDeployer.contract.UnpackLog(event, "DeployedWallet", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WalletDeployerMigratedWalletIterator is returned from FilterMigratedWallet and is used to iterate over the raw logs and unpacked data for MigratedWallet events raised by the WalletDeployer contract.
type WalletDeployerMigratedWalletIterator struct {
	Event *WalletDeployerMigratedWallet // Event containing the contract specifics and raw log

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
func (it *WalletDeployerMigratedWalletIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletDeployerMigratedWallet)
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
		it.Event = new(WalletDeployerMigratedWallet)
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
func (it *WalletDeployerMigratedWalletIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletDeployerMigratedWalletIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletDeployerMigratedWallet represents a MigratedWallet event raised by the WalletDeployer contract.
type WalletDeployerMigratedWallet struct {
	Wallet    common.Address
	OldWallet common.Address
	Owner     common.Address
	Paid      *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMigratedWallet is a free log retrieval operation binding the contract event 0xc65d6ee9571556236e352151c95c79b6589474ad814195aaac7d5ab8d88ba2dd.
//
// Solidity: event MigratedWallet(address _wallet, address _oldWallet, address _owner, uint256 _paid)
func (_WalletDeployer *WalletDeployerFilterer) FilterMigratedWallet(opts *bind.FilterOpts) (*WalletDeployerMigratedWalletIterator, error) {

	logs, sub, err := _WalletDeployer.contract.FilterLogs(opts, "MigratedWallet")
	if err != nil {
		return nil, err
	}
	return &WalletDeployerMigratedWalletIterator{contract: _WalletDeployer.contract, event: "MigratedWallet", logs: logs, sub: sub}, nil
}

// WatchMigratedWallet is a free log subscription operation binding the contract event 0xc65d6ee9571556236e352151c95c79b6589474ad814195aaac7d5ab8d88ba2dd.
//
// Solidity: event MigratedWallet(address _wallet, address _oldWallet, address _owner, uint256 _paid)
func (_WalletDeployer *WalletDeployerFilterer) WatchMigratedWallet(opts *bind.WatchOpts, sink chan<- *WalletDeployerMigratedWallet) (event.Subscription, error) {

	logs, sub, err := _WalletDeployer.contract.WatchLogs(opts, "MigratedWallet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletDeployerMigratedWallet)
				if err := _WalletDeployer.contract.UnpackLog(event, "MigratedWallet", log); err != nil {
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

// ParseMigratedWallet is a log parse operation binding the contract event 0xc65d6ee9571556236e352151c95c79b6589474ad814195aaac7d5ab8d88ba2dd.
//
// Solidity: event MigratedWallet(address _wallet, address _oldWallet, address _owner, uint256 _paid)
func (_WalletDeployer *WalletDeployerFilterer) ParseMigratedWallet(log types.Log) (*WalletDeployerMigratedWallet, error) {
	event := new(WalletDeployerMigratedWallet)
	if err := _WalletDeployer.contract.UnpackLog(event, "MigratedWallet", log); err != nil {
		return nil, err
	}
	return event, nil
}
