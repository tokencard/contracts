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
const WalletCacheABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_walletImplementation_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_ens_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_defaultSpendLimit_\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_controllerNode_\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_licenceNode_\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_tokenWhitelistNode_\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_walletDeployerNode_\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"addresspayable\",\"name\":\"_wallet\",\"type\":\"address\"}],\"name\":\"CachedWallet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_newWalletImplementation\",\"type\":\"address\"}],\"name\":\"setNewWalletImplementation\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[],\"name\":\"cacheWallet\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"cachedWallets\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cachedWalletsCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"controllerNode\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"defaultSpendLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ens\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ensRegistry\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"licenceNode\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newWalletImplementation\",\"type\":\"address\"}],\"name\":\"setNewWalletImplementaton\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tokenWhitelistNode\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"walletCachePop\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"walletDeployerNode\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"walletImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// WalletCacheBin is the compiled bytecode used for deploying new contracts.
var WalletCacheBin = "0x60806040527f7f2ce995617d2816b426c5c8698c5ec2952f7a34bb10f38326f74933d58936976034557fd0ff8bd67f6e25e4e4b010df582a36a0ee9b78e49afe6cc1cff5dd5a830403306035557fe84f90570f13fe09f288f2411ff9cf50da611ed0c7db7f73d48053ffc974d3966036557f1d0c0adbe6addd93659446311e0767a56b67d41ef38f0cb66dcf7560d28a5a386037553480156100a057600080fd5b5060405161128d38038061128d833981810160405260e08110156100c357600080fd5b508051602082015160408301516060840151608085015160a086015160c0909601519495939492939192909190610102866001600160e01b0361017916565b610114846001600160e01b0361028916565b603880546001600160a01b03808a166001600160a01b0319928316179092556039805492891692909116919091179055603a85905582156101555760358390555b81156101615760368290555b801561016d5760378190555b50505050505050610346565b600054610100900460ff168061019b575061019b6001600160e01b0361034016565b806101a9575060005460ff16155b6101e45760405162461bcd60e51b815260040180806020018281038252602e81526020018061125f602e913960400191505060405180910390fd5b600054610100900460ff1615801561020f576000805460ff1961ff0019909116610100171660011790555b6001600160a01b038216610258576040805162461bcd60e51b815260206004820152600b60248201526a0656e7352656720697320360ac1b604482015290519081900360640190fd5b603380546001600160a01b0319166001600160a01b0384161790558015610285576000805461ff00191690555b5050565b600054610100900460ff16806102ab57506102ab6001600160e01b0361034016565b806102b9575060005460ff16155b6102f45760405162461bcd60e51b815260040180806020018281038252602e81526020018061125f602e913960400191505060405180910390fd5b600054610100900460ff1615801561031f576000805460ff1961ff0019909116610100171660011790555b811561032b5760348290555b8015610285576000805461ff00191690555050565b303b1590565b610f0a806103556000396000f3fe608060405234801561001057600080fd5b50600436106100cf5760003560e01c8063877337b01161008c578063a7a7d3bc11610066578063a7a7d3bc1461014c578063d30a1c0d14610154578063e2b4ce971461017a578063fc6cabe714610182576100cf565b8063877337b01461013457806387e8bed11461013c578063a4570e5114610144576100cf565b80633f15457f146100d457806360dbc5de146100f857806371b9076a14610102578063747c31d61461011c5780637d73b231146101245780638117abc11461012c575b600080fd5b6100dc61019f565b604080516001600160a01b039092168252519081900360200190f35b6101006101ae565b005b61010a61036d565b60408051918252519081900360200190f35b61010a610373565b6100dc610379565b6100dc610388565b61010a610397565b61010a61039d565b6100dc6103a3565b61010a61048a565b6101006004803603602081101561016a57600080fd5b50356001600160a01b0316610490565b61010a6105a7565b6100dc6004803603602081101561019857600080fd5b50356105ad565b6039546001600160a01b031681565b60006101bb6037546105d4565b6038546040519192506000916001600160a01b039091169083906101de90610789565b6001600160a01b03928316815291166020820152606060408083018290526000918301829052519182900360a0019190f080158015610221573d6000803e3d6000fd5b509050806001600160a01b03166369efdfc0836001603960009054906101000a90046001600160a01b03166036546102576105a7565b603554603a54604080516001600160e01b031960e08b901b1681526001600160a01b03988916600482015296151560248801529490961660448601526064850192909252608484015260a483015260c4820192909252905160e480830192600092919082900301818387803b1580156102cf57600080fd5b505af11580156102e3573d6000803e3d6000fd5b5050603b80546001810182556000919091527fbbe3212124853f8b0084a66a2d057c2966e251e132af3691db153ab65f0d1a4d0180546001600160a01b0385166001600160a01b0319909116811790915560408051918252517f9ede7876a6b2454072ceeaff4b6b4e6eaa5381db241b850f2a46034136fc2e6e9350908190036020019150a15050565b603b5490565b60355481565b6033546001600160a01b031690565b6038546001600160a01b031681565b60365481565b60375481565b60006103b06037546105d4565b6001600160a01b0316336001600160a01b031614610415576040805162461bcd60e51b815260206004820152601d60248201527f6e6f742063616c6c65642062792077616c6c65742d6465706c6f796572000000604482015290519081900360640190fd5b603b5460011115610428576104286101ae565b603b805460009190600019810190811061043e57fe5b600091825260209091200154603b80546001600160a01b039092169250908061046357fe5b600082815260209020810160001990810180546001600160a01b0319169055019055905090565b603a5481565b61049933610727565b6104e3576040805162461bcd60e51b815260206004820152601660248201527539b2b73232b91034b9903737ba1030b71030b236b4b760511b604482015290519081900360640190fd5b6001600160a01b0381161580159061050957506038546001600160a01b03828116911614155b610553576040805162461bcd60e51b815260206004820152601660248201527534b73b30b634b21034b6b83632b6b2b73a30ba34b7b760511b604482015290519081900360640190fd5b603880546001600160a01b0383166001600160a01b0319909116811790915560408051918252517fed6a15d6ec63befef25c99a6aa87c7723e4c73e1ffde18a399c6a8abb75172649181900360200190a150565b60345490565b603b81815481106105ba57fe5b6000918252602090912001546001600160a01b0316905081565b6033546000906001600160a01b0316610634576040805162461bcd60e51b815260206004820152601d60248201527f454e535265736f6c7661626c65206e6f7420696e697469616c697a6564000000604482015290519081900360640190fd5b60335460408051630178b8bf60e01b81526004810185905290516001600160a01b0390921691630178b8bf91602480820192602092909190829003018186803b15801561068057600080fd5b505afa158015610694573d6000803e3d6000fd5b505050506040513d60208110156106aa57600080fd5b505160408051631d9dabef60e11b81526004810185905290516001600160a01b0390921691633b3b57de91602480820192602092909190829003018186803b1580156106f557600080fd5b505afa158015610709573d6000803e3d6000fd5b505050506040513d602081101561071f57600080fd5b505192915050565b60006107346034546105d4565b6001600160a01b03166324d7806c836040518263ffffffff1660e01b815260040180826001600160a01b03166001600160a01b0316815260200191505060206040518083038186803b1580156106f557600080fd5b61073f806107978339019056fe608060405260405161073f38038061073f8339818101604052606081101561002657600080fd5b8151602083015160408085018051915193959294830192918464010000000082111561005157600080fd5b90830190602082018581111561006657600080fd5b825164010000000081118282018810171561008057600080fd5b82525081516020918201929091019080838360005b838110156100ad578181015183820152602001610095565b50505050905090810190601f1680156100da5780820380516001836020036101000a031916815260200191505b5060408181527f656970313936372e70726f78792e696d706c656d656e746174696f6e0000000082525190819003601c0190208693508492506000805160206106e483398151915260001990910114905061013157fe5b610143826001600160e01b0361026516565b8051156101fb576000826001600160a01b0316826040518082805190602001908083835b602083106101865780518252601f199092019160209182019101610167565b6001836020036101000a038019825116818451168082178552505050505050905001915050600060405180830381855af49150503d80600081146101e6576040519150601f19603f3d011682016040523d82523d6000602084013e6101eb565b606091505b50509050806101f957600080fd5b505b5050604080517f656970313936372e70726f78792e61646d696e00000000000000000000000000815290519081900360130190206000805160206106c48339815191526000199091011461024b57fe5b61025d826001600160e01b036102c516565b5050506102dd565b610278816102d760201b61036e1760201c565b6102b35760405162461bcd60e51b815260040180806020018281038252603b815260200180610704603b913960400191505060405180910390fd5b6000805160206106e483398151915255565b6000805160206106c483398151915255565b3b151590565b6103d8806102ec6000396000f3fe6080604052600436106100345760003560e01c80635c60da1b1461003e5780638f2839701461006f578063f851a440146100a2575b61003c6100b7565b005b34801561004a57600080fd5b506100536100d1565b604080516001600160a01b039092168252519081900360200190f35b34801561007b57600080fd5b5061003c6004803603602081101561009257600080fd5b50356001600160a01b03166100e0565b3480156100ae57600080fd5b506100536102ca565b6100bf6102d4565b6100cf6100ca6102dc565b610301565b565b60006100db6102dc565b905090565b6100e8610325565b6001600160a01b0316336001600160a01b031614156102bf576001600160a01b0381166101465760405162461bcd60e51b815260040180806020018281038252602f815260200180610375602f913960400191505060405180910390fd5b600060606101526102dc565b6001600160a01b03166040518080632121dc7560e01b8152506004019050600060405180830381855af49150503d80600081146101ab576040519150601f19603f3d011682016040523d82523d6000602084013e6101b0565b606091505b5091509150816101fc576040805162461bcd60e51b815260206004820152601260248201527118da185b99d950591b5a5b8819985a5b195960721b604482015290519081900360640190fd5b80806020019051602081101561021157600080fd5b5051610264576040805162461bcd60e51b815260206004820152601c60248201527f6e6f6e2d7472616e7366657261626c652070726f78792061646d696e00000000604482015290519081900360640190fd5b7f7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f61028d610325565b604080516001600160a01b03928316815291861660208301528051918290030190a16102b88361034a565b50506102c7565b6102c76100b7565b50565b60006100db610325565b6100cf6100cf565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5490565b3660008037600080366000845af43d6000803e808015610320573d6000f35b3d6000fd5b7fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d61035490565b7fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d610355565b3b15159056fe43616e6e6f74206368616e6765207468652061646d696e206f6620612070726f787920746f20616464726573732030a265627a7a723158201be2cdd95b5f17ede0f0a2f47551e1702ee9ec87810031845615fbb1bf0b622f64736f6c63430005110032b53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d6103360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc43616e6e6f742073657420612070726f787920696d706c656d656e746174696f6e20746f2061206e6f6e2d636f6e74726163742061646472657373a265627a7a723158208481e66053a069072039a4f7c2a23d1d4090822e373fe502e5639ef5530d3b7d64736f6c63430005110032436f6e747261637420696e7374616e63652068617320616c7265616479206265656e20696e697469616c697a6564"

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

// Ens is a free data retrieval call binding the contract method 0x3f15457f.
//
// Solidity: function ens() constant returns(address)
func (_WalletCache *WalletCacheCaller) Ens(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _WalletCache.contract.Call(opts, out, "ens")
	return *ret0, err
}

// Ens is a free data retrieval call binding the contract method 0x3f15457f.
//
// Solidity: function ens() constant returns(address)
func (_WalletCache *WalletCacheSession) Ens() (common.Address, error) {
	return _WalletCache.Contract.Ens(&_WalletCache.CallOpts)
}

// Ens is a free data retrieval call binding the contract method 0x3f15457f.
//
// Solidity: function ens() constant returns(address)
func (_WalletCache *WalletCacheCallerSession) Ens() (common.Address, error) {
	return _WalletCache.Contract.Ens(&_WalletCache.CallOpts)
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

// SetNewWalletImplementaton is a paid mutator transaction binding the contract method 0xd30a1c0d.
//
// Solidity: function setNewWalletImplementaton(address _newWalletImplementation) returns()
func (_WalletCache *WalletCacheTransactor) SetNewWalletImplementaton(opts *bind.TransactOpts, _newWalletImplementation common.Address) (*types.Transaction, error) {
	return _WalletCache.contract.Transact(opts, "setNewWalletImplementaton", _newWalletImplementation)
}

// SetNewWalletImplementaton is a paid mutator transaction binding the contract method 0xd30a1c0d.
//
// Solidity: function setNewWalletImplementaton(address _newWalletImplementation) returns()
func (_WalletCache *WalletCacheSession) SetNewWalletImplementaton(_newWalletImplementation common.Address) (*types.Transaction, error) {
	return _WalletCache.Contract.SetNewWalletImplementaton(&_WalletCache.TransactOpts, _newWalletImplementation)
}

// SetNewWalletImplementaton is a paid mutator transaction binding the contract method 0xd30a1c0d.
//
// Solidity: function setNewWalletImplementaton(address _newWalletImplementation) returns()
func (_WalletCache *WalletCacheTransactorSession) SetNewWalletImplementaton(_newWalletImplementation common.Address) (*types.Transaction, error) {
	return _WalletCache.Contract.SetNewWalletImplementaton(&_WalletCache.TransactOpts, _newWalletImplementation)
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

// WalletCacheSetNewWalletImplementationIterator is returned from FilterSetNewWalletImplementation and is used to iterate over the raw logs and unpacked data for SetNewWalletImplementation events raised by the WalletCache contract.
type WalletCacheSetNewWalletImplementationIterator struct {
	Event *WalletCacheSetNewWalletImplementation // Event containing the contract specifics and raw log

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
func (it *WalletCacheSetNewWalletImplementationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletCacheSetNewWalletImplementation)
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
		it.Event = new(WalletCacheSetNewWalletImplementation)
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
func (it *WalletCacheSetNewWalletImplementationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletCacheSetNewWalletImplementationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletCacheSetNewWalletImplementation represents a SetNewWalletImplementation event raised by the WalletCache contract.
type WalletCacheSetNewWalletImplementation struct {
	NewWalletImplementation common.Address
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterSetNewWalletImplementation is a free log retrieval operation binding the contract event 0xed6a15d6ec63befef25c99a6aa87c7723e4c73e1ffde18a399c6a8abb7517264.
//
// Solidity: event setNewWalletImplementation(address _newWalletImplementation)
func (_WalletCache *WalletCacheFilterer) FilterSetNewWalletImplementation(opts *bind.FilterOpts) (*WalletCacheSetNewWalletImplementationIterator, error) {

	logs, sub, err := _WalletCache.contract.FilterLogs(opts, "setNewWalletImplementation")
	if err != nil {
		return nil, err
	}
	return &WalletCacheSetNewWalletImplementationIterator{contract: _WalletCache.contract, event: "setNewWalletImplementation", logs: logs, sub: sub}, nil
}

// WatchSetNewWalletImplementation is a free log subscription operation binding the contract event 0xed6a15d6ec63befef25c99a6aa87c7723e4c73e1ffde18a399c6a8abb7517264.
//
// Solidity: event setNewWalletImplementation(address _newWalletImplementation)
func (_WalletCache *WalletCacheFilterer) WatchSetNewWalletImplementation(opts *bind.WatchOpts, sink chan<- *WalletCacheSetNewWalletImplementation) (event.Subscription, error) {

	logs, sub, err := _WalletCache.contract.WatchLogs(opts, "setNewWalletImplementation")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletCacheSetNewWalletImplementation)
				if err := _WalletCache.contract.UnpackLog(event, "setNewWalletImplementation", log); err != nil {
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

// ParseSetNewWalletImplementation is a log parse operation binding the contract event 0xed6a15d6ec63befef25c99a6aa87c7723e4c73e1ffde18a399c6a8abb7517264.
//
// Solidity: event setNewWalletImplementation(address _newWalletImplementation)
func (_WalletCache *WalletCacheFilterer) ParseSetNewWalletImplementation(log types.Log) (*WalletCacheSetNewWalletImplementation, error) {
	event := new(WalletCacheSetNewWalletImplementation)
	if err := _WalletCache.contract.UnpackLog(event, "setNewWalletImplementation", log); err != nil {
		return nil, err
	}
	return event, nil
}
