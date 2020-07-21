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

// WalletCacheABI is the input ABI used to generate the binding from.
const WalletCacheABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_walletImplementation_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_ens_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_defaultSpendLimit_\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_controllerNode_\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_licenceNode_\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_tokenWhitelistNode_\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_walletDeployerNode_\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"addresspayable\",\"name\":\"_wallet\",\"type\":\"address\"}],\"name\":\"CachedWallet\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"cacheWallet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"cachedWallets\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cachedWalletsCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"controllerNode\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultSpendLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ensRegistry\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"licenceNode\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenWhitelistNode\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"walletCachePop\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"walletDeployerNode\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"walletImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// WalletCacheBin is the compiled bytecode used for deploying new contracts.
var WalletCacheBin = "0x6080604052603380546001600160a01b0319166e0c2e074ec69a0dfb2997ba6c7d2e1e1790557f7f2ce995617d2816b426c5c8698c5ec2952f7a34bb10f38326f74933d58936976034557fd0ff8bd67f6e25e4e4b010df582a36a0ee9b78e49afe6cc1cff5dd5a830403306035557fe84f90570f13fe09f288f2411ff9cf50da611ed0c7db7f73d48053ffc974d3966036557f1d0c0adbe6addd93659446311e0767a56b67d41ef38f0cb66dcf7560d28a5a386037553480156100c157600080fd5b50604051610cce380380610cce833981810160405260e08110156100e457600080fd5b508051602082015160408301516060840151608085015160a086015160c0909601519495939492939192909190610123866001600160e01b0361018516565b610135846001600160e01b0361025c16565b603880546001600160a01b0319166001600160a01b038916179055603985905582156101615760358390555b811561016d5760368290555b80156101795760378190555b50505050505050610318565b600054610100900460ff16806101a757506101a76001600160e01b0361031216565b806101b5575060005460ff16155b6101f05760405162461bcd60e51b815260040180806020018281038252602e815260200180610ca0602e913960400191505060405180910390fd5b600054610100900460ff1615801561021b576000805460ff1961ff0019909116610100171660011790555b6001600160a01b0382161561024657603380546001600160a01b0319166001600160a01b0384161790555b8015610258576000805461ff00191690555b5050565b600054610100900460ff168061027e575061027e6001600160e01b0361031216565b8061028c575060005460ff16155b6102c75760405162461bcd60e51b815260040180806020018281038252602e815260200180610ca0602e913960400191505060405180910390fd5b600054610100900460ff161580156102f2576000805460ff1961ff0019909116610100171660011790555b81156102465760348290558015610258576000805461ff00191690555050565b303b1590565b610979806103276000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c8063877337b011610071578063877337b01461010657806387e8bed11461010e578063a4570e5114610116578063a7a7d3bc1461011e578063e2b4ce9714610126578063fc6cabe71461012e576100a9565b806360dbc5de146100ae57806371b9076a146100b8578063747c31d6146100d25780637d73b231146100da5780638117abc1146100fe575b600080fd5b6100b661014b565b005b6100c06102f3565b60408051918252519081900360200190f35b6100c06102f9565b6100e26102ff565b604080516001600160a01b039092168252519081900360200190f35b6100e261030e565b6100c061031d565b6100c0610323565b6100e2610329565b6100c0610410565b6100c0610416565b6100e26004803603602081101561014457600080fd5b503561041c565b6000610158603754610443565b6038546040519192506000916001600160a01b039091169061017990610596565b6001600160a01b03909116815260406020820181905260008183018190529051918290036080019190f0801580156101b5573d6000803e3d6000fd5b509050806001600160a01b03166369efdfc08360016101d26102ff565b6036546101dd610416565b603554603954604080516001600160e01b031960e08b901b1681526001600160a01b03988916600482015296151560248801529490961660448601526064850192909252608484015260a483015260c4820192909252905160e480830192600092919082900301818387803b15801561025557600080fd5b505af1158015610269573d6000803e3d6000fd5b5050603a80546001810182556000919091527fa2999d817b6757290b50e8ecf3fa939673403dd35c97de392fdb343b4015ce9e0180546001600160a01b0385166001600160a01b0319909116811790915560408051918252517f9ede7876a6b2454072ceeaff4b6b4e6eaa5381db241b850f2a46034136fc2e6e9350908190036020019150a15050565b603a5490565b60355481565b6033546001600160a01b031690565b6038546001600160a01b031681565b60365481565b60375481565b6000610336603754610443565b6001600160a01b0316336001600160a01b03161461039b576040805162461bcd60e51b815260206004820152601d60248201527f6e6f742063616c6c65642062792077616c6c65742d6465706c6f796572000000604482015290519081900360640190fd5b603a54600111156103ae576103ae61014b565b603a80546000919060001981019081106103c457fe5b600091825260209091200154603a80546001600160a01b03909216925090806103e957fe5b600082815260209020810160001990810180546001600160a01b0319169055019055905090565b60395481565b60345490565b603a818154811061042957fe5b6000918252602090912001546001600160a01b0316905081565b6033546000906001600160a01b03166104a3576040805162461bcd60e51b815260206004820152601d60248201527f454e535265736f6c7661626c65206e6f7420696e697469616c697a6564000000604482015290519081900360640190fd5b60335460408051630178b8bf60e01b81526004810185905290516001600160a01b0390921691630178b8bf91602480820192602092909190829003018186803b1580156104ef57600080fd5b505afa158015610503573d6000803e3d6000fd5b505050506040513d602081101561051957600080fd5b505160408051631d9dabef60e11b81526004810185905290516001600160a01b0390921691633b3b57de91602480820192602092909190829003018186803b15801561056457600080fd5b505afa158015610578573d6000803e3d6000fd5b505050506040513d602081101561058e57600080fd5b505192915050565b6103a0806105a48339019056fe60806040526040516103a03803806103a08339818101604052604081101561002657600080fd5b81516020830180516040519294929383019291908464010000000082111561004d57600080fd5b90830190602082018581111561006257600080fd5b825164010000000081118282018810171561007c57600080fd5b82525081516020918201929091019080838360005b838110156100a9578181015183820152602001610091565b50505050905090810190601f1680156100d65780820380516001836020036101000a031916815260200191505b5060408181527f656970313936372e70726f78792e696d706c656d656e746174696f6e0000000082525190819003601c01902060008051602061034583398151915260001990910114925061012a91505057fe5b61013c826001600160e01b036101fb16565b8051156101f4576000826001600160a01b0316826040518082805190602001908083835b6020831061017f5780518252601f199092019160209182019101610160565b6001836020036101000a038019825116818451168082178552505050505050905001915050600060405180830381855af49150503d80600081146101df576040519150601f19603f3d011682016040523d82523d6000602084013e6101e4565b606091505b50509050806101f257600080fd5b505b5050610261565b61020e8161025b60201b61009a1760201c565b6102495760405162461bcd60e51b815260040180806020018281038252603b815260200180610365603b913960400191505060405180910390fd5b60008051602061034583398151915255565b3b151590565b60d68061026f6000396000f3fe6080604052366044576040805133815234602082015281517f88a5966d370b9919b20f3e2c13ff65706f196a4e32cc2c12bf57088f88525874929181900390910190a1005b6050604c6052565b6077565b005b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5490565b3660008037600080366000845af43d6000803e8080156095573d6000f35b3d6000fd5b3b15159056fea26469706673582212208bfb9225a6c250950d4189db972d83a09a36faf516dfbe43c19678125cc0b49a64736f6c634300060b0033360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc43616e6e6f742073657420612070726f787920696d706c656d656e746174696f6e20746f2061206e6f6e2d636f6e74726163742061646472657373a2646970667358221220a624ec591bdbac61d983da0ea66006880f040e12f189e66f4daefb2f6f50e00864736f6c634300060b0033436f6e747261637420696e7374616e63652068617320616c7265616479206265656e20696e697469616c697a6564"

// DeployWalletCache deploys a new Ethereum contract, binding an instance of WalletCache to it.
func DeployWalletCache(auth *bind.TransactOpts, backend bind.ContractBackend, _walletImplementation_ common.Address, _ens_ common.Address, _defaultSpendLimit_ *big.Int, _controllerNode_ [32]byte, _licenceNode_ [32]byte, _tokenWhitelistNode_ [32]byte, _walletDeployerNode_ [32]byte) (common.Address, *types.Transaction, *WalletCache, error) {
	parsed, err := abi.JSON(strings.NewReader(WalletCacheABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(WalletCacheBin), backend, _walletImplementation_, _ens_, _defaultSpendLimit_, _controllerNode_, _licenceNode_, _tokenWhitelistNode_, _walletDeployerNode_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &WalletCache{WalletCacheCaller: WalletCacheCaller{contract: contract}, WalletCacheTransactor: WalletCacheTransactor{contract: contract}, WalletCacheFilterer: WalletCacheFilterer{contract: contract}}, nil
}

// WalletCache is an auto generated Go binding around an Ethereum contract.
type WalletCache struct {
	WalletCacheCaller     // Read-only binding to the contract
	WalletCacheTransactor // Write-only binding to the contract
	WalletCacheFilterer   // Log filterer for contract events
}

// WalletCacheCaller is an auto generated read-only Go binding around an Ethereum contract.
type WalletCacheCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletCacheTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WalletCacheTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletCacheFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WalletCacheFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletCacheSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WalletCacheSession struct {
	Contract     *WalletCache      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WalletCacheCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WalletCacheCallerSession struct {
	Contract *WalletCacheCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// WalletCacheTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WalletCacheTransactorSession struct {
	Contract     *WalletCacheTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// WalletCacheRaw is an auto generated low-level Go binding around an Ethereum contract.
type WalletCacheRaw struct {
	Contract *WalletCache // Generic contract binding to access the raw methods on
}

// WalletCacheCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WalletCacheCallerRaw struct {
	Contract *WalletCacheCaller // Generic read-only contract binding to access the raw methods on
}

// WalletCacheTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WalletCacheTransactorRaw struct {
	Contract *WalletCacheTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWalletCache creates a new instance of WalletCache, bound to a specific deployed contract.
func NewWalletCache(address common.Address, backend bind.ContractBackend) (*WalletCache, error) {
	contract, err := bindWalletCache(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WalletCache{WalletCacheCaller: WalletCacheCaller{contract: contract}, WalletCacheTransactor: WalletCacheTransactor{contract: contract}, WalletCacheFilterer: WalletCacheFilterer{contract: contract}}, nil
}

// NewWalletCacheCaller creates a new read-only instance of WalletCache, bound to a specific deployed contract.
func NewWalletCacheCaller(address common.Address, caller bind.ContractCaller) (*WalletCacheCaller, error) {
	contract, err := bindWalletCache(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WalletCacheCaller{contract: contract}, nil
}

// NewWalletCacheTransactor creates a new write-only instance of WalletCache, bound to a specific deployed contract.
func NewWalletCacheTransactor(address common.Address, transactor bind.ContractTransactor) (*WalletCacheTransactor, error) {
	contract, err := bindWalletCache(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WalletCacheTransactor{contract: contract}, nil
}

// NewWalletCacheFilterer creates a new log filterer instance of WalletCache, bound to a specific deployed contract.
func NewWalletCacheFilterer(address common.Address, filterer bind.ContractFilterer) (*WalletCacheFilterer, error) {
	contract, err := bindWalletCache(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WalletCacheFilterer{contract: contract}, nil
}

// bindWalletCache binds a generic wrapper to an already deployed contract.
func bindWalletCache(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(WalletCacheABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WalletCache *WalletCacheRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _WalletCache.Contract.WalletCacheCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WalletCache *WalletCacheRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WalletCache.Contract.WalletCacheTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WalletCache *WalletCacheRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WalletCache.Contract.WalletCacheTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WalletCache *WalletCacheCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _WalletCache.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WalletCache *WalletCacheTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WalletCache.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WalletCache *WalletCacheTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WalletCache.Contract.contract.Transact(opts, method, params...)
}

// CachedWallets is a free data retrieval call binding the contract method 0xfc6cabe7.
//
// Solidity: function cachedWallets(uint256 ) constant returns(address)
func (_WalletCache *WalletCacheCaller) CachedWallets(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _WalletCache.contract.Call(opts, out, "cachedWallets", arg0)
	return *ret0, err
}

// CachedWallets is a free data retrieval call binding the contract method 0xfc6cabe7.
//
// Solidity: function cachedWallets(uint256 ) constant returns(address)
func (_WalletCache *WalletCacheSession) CachedWallets(arg0 *big.Int) (common.Address, error) {
	return _WalletCache.Contract.CachedWallets(&_WalletCache.CallOpts, arg0)
}

// CachedWallets is a free data retrieval call binding the contract method 0xfc6cabe7.
//
// Solidity: function cachedWallets(uint256 ) constant returns(address)
func (_WalletCache *WalletCacheCallerSession) CachedWallets(arg0 *big.Int) (common.Address, error) {
	return _WalletCache.Contract.CachedWallets(&_WalletCache.CallOpts, arg0)
}

// CachedWalletsCount is a free data retrieval call binding the contract method 0x71b9076a.
//
// Solidity: function cachedWalletsCount() constant returns(uint256)
func (_WalletCache *WalletCacheCaller) CachedWalletsCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _WalletCache.contract.Call(opts, out, "cachedWalletsCount")
	return *ret0, err
}

// CachedWalletsCount is a free data retrieval call binding the contract method 0x71b9076a.
//
// Solidity: function cachedWalletsCount() constant returns(uint256)
func (_WalletCache *WalletCacheSession) CachedWalletsCount() (*big.Int, error) {
	return _WalletCache.Contract.CachedWalletsCount(&_WalletCache.CallOpts)
}

// CachedWalletsCount is a free data retrieval call binding the contract method 0x71b9076a.
//
// Solidity: function cachedWalletsCount() constant returns(uint256)
func (_WalletCache *WalletCacheCallerSession) CachedWalletsCount() (*big.Int, error) {
	return _WalletCache.Contract.CachedWalletsCount(&_WalletCache.CallOpts)
}

// ControllerNode is a free data retrieval call binding the contract method 0xe2b4ce97.
//
// Solidity: function controllerNode() constant returns(bytes32)
func (_WalletCache *WalletCacheCaller) ControllerNode(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _WalletCache.contract.Call(opts, out, "controllerNode")
	return *ret0, err
}

// ControllerNode is a free data retrieval call binding the contract method 0xe2b4ce97.
//
// Solidity: function controllerNode() constant returns(bytes32)
func (_WalletCache *WalletCacheSession) ControllerNode() ([32]byte, error) {
	return _WalletCache.Contract.ControllerNode(&_WalletCache.CallOpts)
}

// ControllerNode is a free data retrieval call binding the contract method 0xe2b4ce97.
//
// Solidity: function controllerNode() constant returns(bytes32)
func (_WalletCache *WalletCacheCallerSession) ControllerNode() ([32]byte, error) {
	return _WalletCache.Contract.ControllerNode(&_WalletCache.CallOpts)
}

// DefaultSpendLimit is a free data retrieval call binding the contract method 0xa7a7d3bc.
//
// Solidity: function defaultSpendLimit() constant returns(uint256)
func (_WalletCache *WalletCacheCaller) DefaultSpendLimit(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _WalletCache.contract.Call(opts, out, "defaultSpendLimit")
	return *ret0, err
}

// DefaultSpendLimit is a free data retrieval call binding the contract method 0xa7a7d3bc.
//
// Solidity: function defaultSpendLimit() constant returns(uint256)
func (_WalletCache *WalletCacheSession) DefaultSpendLimit() (*big.Int, error) {
	return _WalletCache.Contract.DefaultSpendLimit(&_WalletCache.CallOpts)
}

// DefaultSpendLimit is a free data retrieval call binding the contract method 0xa7a7d3bc.
//
// Solidity: function defaultSpendLimit() constant returns(uint256)
func (_WalletCache *WalletCacheCallerSession) DefaultSpendLimit() (*big.Int, error) {
	return _WalletCache.Contract.DefaultSpendLimit(&_WalletCache.CallOpts)
}

// EnsRegistry is a free data retrieval call binding the contract method 0x7d73b231.
//
// Solidity: function ensRegistry() constant returns(address)
func (_WalletCache *WalletCacheCaller) EnsRegistry(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _WalletCache.contract.Call(opts, out, "ensRegistry")
	return *ret0, err
}

// EnsRegistry is a free data retrieval call binding the contract method 0x7d73b231.
//
// Solidity: function ensRegistry() constant returns(address)
func (_WalletCache *WalletCacheSession) EnsRegistry() (common.Address, error) {
	return _WalletCache.Contract.EnsRegistry(&_WalletCache.CallOpts)
}

// EnsRegistry is a free data retrieval call binding the contract method 0x7d73b231.
//
// Solidity: function ensRegistry() constant returns(address)
func (_WalletCache *WalletCacheCallerSession) EnsRegistry() (common.Address, error) {
	return _WalletCache.Contract.EnsRegistry(&_WalletCache.CallOpts)
}

// LicenceNode is a free data retrieval call binding the contract method 0x747c31d6.
//
// Solidity: function licenceNode() constant returns(bytes32)
func (_WalletCache *WalletCacheCaller) LicenceNode(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _WalletCache.contract.Call(opts, out, "licenceNode")
	return *ret0, err
}

// LicenceNode is a free data retrieval call binding the contract method 0x747c31d6.
//
// Solidity: function licenceNode() constant returns(bytes32)
func (_WalletCache *WalletCacheSession) LicenceNode() ([32]byte, error) {
	return _WalletCache.Contract.LicenceNode(&_WalletCache.CallOpts)
}

// LicenceNode is a free data retrieval call binding the contract method 0x747c31d6.
//
// Solidity: function licenceNode() constant returns(bytes32)
func (_WalletCache *WalletCacheCallerSession) LicenceNode() ([32]byte, error) {
	return _WalletCache.Contract.LicenceNode(&_WalletCache.CallOpts)
}

// TokenWhitelistNode is a free data retrieval call binding the contract method 0x877337b0.
//
// Solidity: function tokenWhitelistNode() constant returns(bytes32)
func (_WalletCache *WalletCacheCaller) TokenWhitelistNode(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _WalletCache.contract.Call(opts, out, "tokenWhitelistNode")
	return *ret0, err
}

// TokenWhitelistNode is a free data retrieval call binding the contract method 0x877337b0.
//
// Solidity: function tokenWhitelistNode() constant returns(bytes32)
func (_WalletCache *WalletCacheSession) TokenWhitelistNode() ([32]byte, error) {
	return _WalletCache.Contract.TokenWhitelistNode(&_WalletCache.CallOpts)
}

// TokenWhitelistNode is a free data retrieval call binding the contract method 0x877337b0.
//
// Solidity: function tokenWhitelistNode() constant returns(bytes32)
func (_WalletCache *WalletCacheCallerSession) TokenWhitelistNode() ([32]byte, error) {
	return _WalletCache.Contract.TokenWhitelistNode(&_WalletCache.CallOpts)
}

// WalletDeployerNode is a free data retrieval call binding the contract method 0x87e8bed1.
//
// Solidity: function walletDeployerNode() constant returns(bytes32)
func (_WalletCache *WalletCacheCaller) WalletDeployerNode(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _WalletCache.contract.Call(opts, out, "walletDeployerNode")
	return *ret0, err
}

// WalletDeployerNode is a free data retrieval call binding the contract method 0x87e8bed1.
//
// Solidity: function walletDeployerNode() constant returns(bytes32)
func (_WalletCache *WalletCacheSession) WalletDeployerNode() ([32]byte, error) {
	return _WalletCache.Contract.WalletDeployerNode(&_WalletCache.CallOpts)
}

// WalletDeployerNode is a free data retrieval call binding the contract method 0x87e8bed1.
//
// Solidity: function walletDeployerNode() constant returns(bytes32)
func (_WalletCache *WalletCacheCallerSession) WalletDeployerNode() ([32]byte, error) {
	return _WalletCache.Contract.WalletDeployerNode(&_WalletCache.CallOpts)
}

// WalletImplementation is a free data retrieval call binding the contract method 0x8117abc1.
//
// Solidity: function walletImplementation() constant returns(address)
func (_WalletCache *WalletCacheCaller) WalletImplementation(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _WalletCache.contract.Call(opts, out, "walletImplementation")
	return *ret0, err
}

// WalletImplementation is a free data retrieval call binding the contract method 0x8117abc1.
//
// Solidity: function walletImplementation() constant returns(address)
func (_WalletCache *WalletCacheSession) WalletImplementation() (common.Address, error) {
	return _WalletCache.Contract.WalletImplementation(&_WalletCache.CallOpts)
}

// WalletImplementation is a free data retrieval call binding the contract method 0x8117abc1.
//
// Solidity: function walletImplementation() constant returns(address)
func (_WalletCache *WalletCacheCallerSession) WalletImplementation() (common.Address, error) {
	return _WalletCache.Contract.WalletImplementation(&_WalletCache.CallOpts)
}

// CacheWallet is a paid mutator transaction binding the contract method 0x60dbc5de.
//
// Solidity: function cacheWallet() returns()
func (_WalletCache *WalletCacheTransactor) CacheWallet(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WalletCache.contract.Transact(opts, "cacheWallet")
}

// CacheWallet is a paid mutator transaction binding the contract method 0x60dbc5de.
//
// Solidity: function cacheWallet() returns()
func (_WalletCache *WalletCacheSession) CacheWallet() (*types.Transaction, error) {
	return _WalletCache.Contract.CacheWallet(&_WalletCache.TransactOpts)
}

// CacheWallet is a paid mutator transaction binding the contract method 0x60dbc5de.
//
// Solidity: function cacheWallet() returns()
func (_WalletCache *WalletCacheTransactorSession) CacheWallet() (*types.Transaction, error) {
	return _WalletCache.Contract.CacheWallet(&_WalletCache.TransactOpts)
}

// WalletCachePop is a paid mutator transaction binding the contract method 0xa4570e51.
//
// Solidity: function walletCachePop() returns(address)
func (_WalletCache *WalletCacheTransactor) WalletCachePop(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WalletCache.contract.Transact(opts, "walletCachePop")
}

// WalletCachePop is a paid mutator transaction binding the contract method 0xa4570e51.
//
// Solidity: function walletCachePop() returns(address)
func (_WalletCache *WalletCacheSession) WalletCachePop() (*types.Transaction, error) {
	return _WalletCache.Contract.WalletCachePop(&_WalletCache.TransactOpts)
}

// WalletCachePop is a paid mutator transaction binding the contract method 0xa4570e51.
//
// Solidity: function walletCachePop() returns(address)
func (_WalletCache *WalletCacheTransactorSession) WalletCachePop() (*types.Transaction, error) {
	return _WalletCache.Contract.WalletCachePop(&_WalletCache.TransactOpts)
}

// WalletCacheCachedWalletIterator is returned from FilterCachedWallet and is used to iterate over the raw logs and unpacked data for CachedWallet events raised by the WalletCache contract.
type WalletCacheCachedWalletIterator struct {
	Event *WalletCacheCachedWallet // Event containing the contract specifics and raw log

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
func (it *WalletCacheCachedWalletIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletCacheCachedWallet)
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
		it.Event = new(WalletCacheCachedWallet)
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
func (it *WalletCacheCachedWalletIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletCacheCachedWalletIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletCacheCachedWallet represents a CachedWallet event raised by the WalletCache contract.
type WalletCacheCachedWallet struct {
	Wallet common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCachedWallet is a free log retrieval operation binding the contract event 0x9ede7876a6b2454072ceeaff4b6b4e6eaa5381db241b850f2a46034136fc2e6e.
//
// Solidity: event CachedWallet(address _wallet)
func (_WalletCache *WalletCacheFilterer) FilterCachedWallet(opts *bind.FilterOpts) (*WalletCacheCachedWalletIterator, error) {

	logs, sub, err := _WalletCache.contract.FilterLogs(opts, "CachedWallet")
	if err != nil {
		return nil, err
	}
	return &WalletCacheCachedWalletIterator{contract: _WalletCache.contract, event: "CachedWallet", logs: logs, sub: sub}, nil
}

// WatchCachedWallet is a free log subscription operation binding the contract event 0x9ede7876a6b2454072ceeaff4b6b4e6eaa5381db241b850f2a46034136fc2e6e.
//
// Solidity: event CachedWallet(address _wallet)
func (_WalletCache *WalletCacheFilterer) WatchCachedWallet(opts *bind.WatchOpts, sink chan<- *WalletCacheCachedWallet) (event.Subscription, error) {

	logs, sub, err := _WalletCache.contract.WatchLogs(opts, "CachedWallet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletCacheCachedWallet)
				if err := _WalletCache.contract.UnpackLog(event, "CachedWallet", log); err != nil {
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

// ParseCachedWallet is a log parse operation binding the contract event 0x9ede7876a6b2454072ceeaff4b6b4e6eaa5381db241b850f2a46034136fc2e6e.
//
// Solidity: event CachedWallet(address _wallet)
func (_WalletCache *WalletCacheFilterer) ParseCachedWallet(log types.Log) (*WalletCacheCachedWallet, error) {
	event := new(WalletCacheCachedWallet)
	if err := _WalletCache.contract.UnpackLog(event, "CachedWallet", log); err != nil {
		return nil, err
	}
	return event, nil
}
