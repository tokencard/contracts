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
const WalletDeployerABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_ens_\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_controllerNode_\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_walletCacheNode_\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractWallet\",\"name\":\"_wallet\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"DeployedWallet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractWallet\",\"name\":\"_wallet\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractWallet\",\"name\":\"_oldWallet\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_paid\",\"type\":\"uint256\"}],\"name\":\"MigratedWallet\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"controllerNode\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"deployWallet\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"deployedWallets\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ensRegistry\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"contractWallet\",\"name\":\"_oldWallet\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_initializedSpendLimit\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"_initializedGasTopUpLimit\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"_initializedWhitelist\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"_spendLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasTopUpLimit\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_whitelistedAddresses\",\"type\":\"address[]\"}],\"name\":\"migrateWallet\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"walletCacheNode\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// WalletDeployerBin is the compiled bytecode used for deploying new contracts.
var WalletDeployerBin = "0x60806040527f7f2ce995617d2816b426c5c8698c5ec2952f7a34bb10f38326f74933d58936976002557f7eee9c3927d17f70ce19de05f73d05dbda3449d450ba9a4c64f24c24bfb9d7ac60035534801561005857600080fd5b50604051610aba380380610aba8339818101604052606081101561007b57600080fd5b5080516020820151604090920151600180546001600160a01b038085166001600160a01b031992831617928390556000805490921692169190911790559091908180156100c85760028190555b5080156100d55760038190555b5050506109d3806100e76000396000f3fe6080604052600436106100555760003560e01c806317ca93051461005a5780637d73b2311461010c57806380a12c0e1461013d578063a6ede3d414610170578063c8cc2fc214610197578063e2b4ce97146101ca575b600080fd5b61010a600480360361010081101561007157600080fd5b6001600160a01b03823581169260208101359091169160408201351515916060810135151591608082013515159160a08101359160c08201359190810190610100810160e08201356401000000008111156100cb57600080fd5b8201836020820111156100dd57600080fd5b803590602001918460208302840111640100000000831117156100ff57600080fd5b5090925090506101df565b005b34801561011857600080fd5b5061012161064d565b604080516001600160a01b039092168252519081900360200190f35b34801561014957600080fd5b506101216004803603602081101561016057600080fd5b50356001600160a01b031661065c565b34801561017c57600080fd5b50610185610677565b60408051918252519081900360200190f35b3480156101a357600080fd5b5061010a600480360360208110156101ba57600080fd5b50356001600160a01b031661067d565b3480156101d657600080fd5b50610185610821565b6101e833610827565b610239576040805162461bcd60e51b815260206004820152601a60248201527f73656e646572206973206e6f74206120636f6e74726f6c6c6572000000000000604482015290519081900360640190fd5b6001600160a01b0389811660009081526004602052604090205416156102905760405162461bcd60e51b815260040180806020018281038252602181526020018061097e6021913960400191505060405180910390fd5b886001600160a01b0316886001600160a01b0316638da5cb5b6040518163ffffffff1660e01b815260040160206040518083038186803b1580156102d357600080fd5b505afa1580156102e7573d6000803e3d6000fd5b505050506040513d60208110156102fd57600080fd5b50516001600160a01b03161461034b576040805162461bcd60e51b815260206004820152600e60248201526d0deeedccae440dad2e6dac2e8c6d60931b604482015290519081900360640190fd5b60006103586003546108bb565b6001600160a01b031663a4570e516040518163ffffffff1660e01b8152600401602060405180830381600087803b15801561039257600080fd5b505af11580156103a6573d6000803e3d6000fd5b505050506040513d60208110156103bc57600080fd5b50519050871561042557806001600160a01b0316633c672eb7866040518263ffffffff1660e01b815260040180828152602001915050600060405180830381600087803b15801561040c57600080fd5b505af1158015610420573d6000803e3d6000fd5b505050505b861561048a57806001600160a01b0316630f3a85d8856040518263ffffffff1660e01b815260040180828152602001915050600060405180830381600087803b15801561047157600080fd5b505af1158015610485573d6000803e3d6000fd5b505050505b851561051e57806001600160a01b031663f421764884846040518363ffffffff1660e01b815260040180806020018281038252848482818152602001925060200280828437600081840152601f19601f8201169050808301925050509350505050600060405180830381600087803b15801561050557600080fd5b505af1158015610519573d6000803e3d6000fd5b505050505b60408051632c90b94d60e21b81526001600160a01b038c8116600483015260006024830181905292519084169263b242e534926044808201939182900301818387803b15801561056d57600080fd5b505af1158015610581573d6000803e3d6000fd5b505050506001600160a01b038a8116600090815260046020526040902080546001600160a01b03191691831691909117905534156105f0576040516001600160a01b038b16903480156108fc02916000818181858888f193505050501580156105ee573d6000803e3d6000fd5b505b604080516001600160a01b038084168252808c1660208301528c168183015234606082015290517fc65d6ee9571556236e352151c95c79b6589474ad814195aaac7d5ab8d88ba2dd9181900360800190a150505050505050505050565b6001546001600160a01b031690565b6004602052600090815260409020546001600160a01b031681565b60035481565b61068633610827565b6106d7576040805162461bcd60e51b815260206004820152601a60248201527f73656e646572206973206e6f74206120636f6e74726f6c6c6572000000000000604482015290519081900360640190fd5b60006106e46003546108bb565b6001600160a01b031663a4570e516040518163ffffffff1660e01b8152600401602060405180830381600087803b15801561071e57600080fd5b505af1158015610732573d6000803e3d6000fd5b505050506040513d602081101561074857600080fd5b505160408051632c90b94d60e21b81526001600160a01b038581166004830152600060248301819052925193945084169263b242e5349260448084019391929182900301818387803b15801561079d57600080fd5b505af11580156107b1573d6000803e3d6000fd5b5050506001600160a01b0380841660008181526004602090815260409182902080549487166001600160a01b031990951685179055815193845283019190915280517fc02db5f4164f89d90905928336769906e16d79c4a77342126eb647ca9440d0789350918290030190a15050565b60025490565b60006108346002546108bb565b6001600160a01b031663b429afeb836040518263ffffffff1660e01b815260040180826001600160a01b03166001600160a01b0316815260200191505060206040518083038186803b15801561088957600080fd5b505afa15801561089d573d6000803e3d6000fd5b505050506040513d60208110156108b357600080fd5b505192915050565b6000805460408051630178b8bf60e01b81526004810185905290516001600160a01b0390921691630178b8bf91602480820192602092909190829003018186803b15801561090857600080fd5b505afa15801561091c573d6000803e3d6000fd5b505050506040513d602081101561093257600080fd5b505160408051631d9dabef60e11b81526004810185905290516001600160a01b0390921691633b3b57de91602480820192602092909190829003018186803b15801561088957600080fdfe77616c6c657420616c7265616479206465706c6f79656420666f72206f776e6572a265627a7a723158208ef5c55225a1f1d6c9b52763d2d5df6f7b66dba0758eff12aab4545bab911ae164736f6c634300050f0032"

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

// MigrateWallet is a paid mutator transaction binding the contract method 0x17ca9305.
//
// Solidity: function migrateWallet(address _owner, address _oldWallet, bool _initializedSpendLimit, bool _initializedGasTopUpLimit, bool _initializedWhitelist, uint256 _spendLimit, uint256 _gasTopUpLimit, address[] _whitelistedAddresses) returns()
func (_WalletDeployer *WalletDeployerTransactor) MigrateWallet(opts *bind.TransactOpts, _owner common.Address, _oldWallet common.Address, _initializedSpendLimit bool, _initializedGasTopUpLimit bool, _initializedWhitelist bool, _spendLimit *big.Int, _gasTopUpLimit *big.Int, _whitelistedAddresses []common.Address) (*types.Transaction, error) {
	return _WalletDeployer.contract.Transact(opts, "migrateWallet", _owner, _oldWallet, _initializedSpendLimit, _initializedGasTopUpLimit, _initializedWhitelist, _spendLimit, _gasTopUpLimit, _whitelistedAddresses)
}

// MigrateWallet is a paid mutator transaction binding the contract method 0x17ca9305.
//
// Solidity: function migrateWallet(address _owner, address _oldWallet, bool _initializedSpendLimit, bool _initializedGasTopUpLimit, bool _initializedWhitelist, uint256 _spendLimit, uint256 _gasTopUpLimit, address[] _whitelistedAddresses) returns()
func (_WalletDeployer *WalletDeployerSession) MigrateWallet(_owner common.Address, _oldWallet common.Address, _initializedSpendLimit bool, _initializedGasTopUpLimit bool, _initializedWhitelist bool, _spendLimit *big.Int, _gasTopUpLimit *big.Int, _whitelistedAddresses []common.Address) (*types.Transaction, error) {
	return _WalletDeployer.Contract.MigrateWallet(&_WalletDeployer.TransactOpts, _owner, _oldWallet, _initializedSpendLimit, _initializedGasTopUpLimit, _initializedWhitelist, _spendLimit, _gasTopUpLimit, _whitelistedAddresses)
}

// MigrateWallet is a paid mutator transaction binding the contract method 0x17ca9305.
//
// Solidity: function migrateWallet(address _owner, address _oldWallet, bool _initializedSpendLimit, bool _initializedGasTopUpLimit, bool _initializedWhitelist, uint256 _spendLimit, uint256 _gasTopUpLimit, address[] _whitelistedAddresses) returns()
func (_WalletDeployer *WalletDeployerTransactorSession) MigrateWallet(_owner common.Address, _oldWallet common.Address, _initializedSpendLimit bool, _initializedGasTopUpLimit bool, _initializedWhitelist bool, _spendLimit *big.Int, _gasTopUpLimit *big.Int, _whitelistedAddresses []common.Address) (*types.Transaction, error) {
	return _WalletDeployer.Contract.MigrateWallet(&_WalletDeployer.TransactOpts, _owner, _oldWallet, _initializedSpendLimit, _initializedGasTopUpLimit, _initializedWhitelist, _spendLimit, _gasTopUpLimit, _whitelistedAddresses)
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
