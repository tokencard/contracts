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

// TokenWhitelistableExporterABI is the input ABI used to generate the binding from.
const TokenWhitelistableExporterABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_a\",\"type\":\"address\"}],\"name\":\"isTokenBurnable\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_a\",\"type\":\"address\"}],\"name\":\"isTokenLoadable\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_a\",\"type\":\"address\"}],\"name\":\"getTokenInfo\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"bool\"},{\"name\":\"\",\"type\":\"bool\"},{\"name\":\"\",\"type\":\"bool\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_a\",\"type\":\"address\"}],\"name\":\"isTokenAvailable\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getStablecoinInfo\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"bool\"},{\"name\":\"\",\"type\":\"bool\"},{\"name\":\"\",\"type\":\"bool\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tokenAddressArray\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ensRegistry\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tokenWhitelistNode\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"},{\"name\":\"_rate\",\"type\":\"uint256\"},{\"name\":\"_updateDate\",\"type\":\"uint256\"}],\"name\":\"updateTokenRate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_ens_\",\"type\":\"address\"},{\"name\":\"_tokenWhitelistName_\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// TokenWhitelistableExporterBin is the compiled bytecode used for deploying new contracts.
const TokenWhitelistableExporterBin = `608060405234801561001057600080fd5b5060405160408061092783398101604052805160209091015160018054600160a060020a03938416600160a060020a031991821617918290556000805490911691909316179091556002556108bd8061006a6000396000f30060806040526004361061007f5763ffffffff60e060020a6000350416630cde34b0811461008457806317e73794146100b95780631f69565f146100da5780632ff0fcaa146101a55780633efec5e9146101c6578063443dd2a4146101db5780637d73b23114610240578063877337b014610271578063d545782e14610298575b600080fd5b34801561009057600080fd5b506100a5600160a060020a03600435166102c1565b604080519115158252519081900360200190f35b3480156100c557600080fd5b506100a5600160a060020a03600435166102d2565b3480156100e657600080fd5b506100fb600160a060020a03600435166102dd565b6040805160208082018990529181018790528515156060820152841515608082015283151560a082015260c0810183905260e08082528951908201528851909182916101008301918b019080838360005b8381101561016457818101518382015260200161014c565b50505050905090810190601f1680156101915780820380516001836020036101000a031916815260200191505b509850505050505050505060405180910390f35b3480156101b157600080fd5b506100a5600160a060020a0360043516610308565b3480156101d257600080fd5b506100fb610313565b3480156101e757600080fd5b506101f061033c565b60408051602080825283518183015283519192839290830191858101910280838360005b8381101561022c578181015183820152602001610214565b505050509050019250505060405180910390f35b34801561024c57600080fd5b5061025561034b565b60408051600160a060020a039092168252519081900360200190f35b34801561027d57600080fd5b5061028661035a565b60408051918252519081900360200190f35b3480156102a457600080fd5b506102bf600160a060020a0360043516602435604435610360565b005b60006102cc82610370565b92915050565b60006102cc82610389565b60606000806000806000806102f1886103a3565b959e949d50929b5090995097509550909350915050565b60006102cc826104d5565b60606000806000806000806103266104ef565b959d949c50929a50909850965094509092509050565b60606103466105fc565b905090565b600154600160a060020a031690565b60025490565b61036b8383836106d6565b505050565b60008061037c836103a3565b5098975050505050505050565b600080610395836103a3565b509098975050505050505050565b60606000806000806000806103b9600254610768565b600160a060020a0316631f69565f896040518263ffffffff1660e060020a0281526004018082600160a060020a0316600160a060020a03168152602001915050600060405180830381600087803b15801561041357600080fd5b505af1158015610427573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405260e081101561045057600080fd5b81019080805164010000000081111561046857600080fd5b8201602081018481111561047b57600080fd5b815164010000000081118282018710171561049557600080fd5b5050602082015160408301516060840151608085015160a086015160c090960151949e50929c50909a509850965090945092505050919395979092949650565b6000806104e1836103a3565b509198975050505050505050565b6060600080600080600080610505600254610768565b600160a060020a0316633efec5e96040518163ffffffff1660e060020a028152600401600060405180830381600087803b15801561054257600080fd5b505af1158015610556573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405260e081101561057f57600080fd5b81019080805164010000000081111561059757600080fd5b820160208101848111156105aa57600080fd5b81516401000000008111828201871017156105c457600080fd5b5050602082015160408301516060840151608085015160a086015160c090960151949f939e50919c509a509850919650945092505050565b6060610609600254610768565b600160a060020a031663443dd2a46040518163ffffffff1660e060020a028152600401600060405180830381600087803b15801561064657600080fd5b505af115801561065a573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052602081101561068357600080fd5b81019080805164010000000081111561069b57600080fd5b820160208101848111156106ae57600080fd5b81518560208202830111640100000000821117156106cb57600080fd5b509094505050505090565b6106e1600254610768565b600160a060020a031663d545782e8484846040518463ffffffff1660e060020a0281526004018084600160a060020a0316600160a060020a031681526020018381526020018281526020019350505050600060405180830381600087803b15801561074b57600080fd5b505af115801561075f573d6000803e3d6000fd5b50505050505050565b60008054604080517f0178b8bf000000000000000000000000000000000000000000000000000000008152600481018590529051600160a060020a0390921691630178b8bf9160248082019260209290919082900301818787803b1580156107cf57600080fd5b505af11580156107e3573d6000803e3d6000fd5b505050506040513d60208110156107f957600080fd5b5051604080517f3b3b57de000000000000000000000000000000000000000000000000000000008152600481018590529051600160a060020a0390921691633b3b57de916024808201926020929091908290030181600087803b15801561085f57600080fd5b505af1158015610873573d6000803e3d6000fd5b505050506040513d602081101561088957600080fd5b5051929150505600a165627a7a72305820a200f44d223d63a0cda8b6ce77784eb74e9f901368a95cd7177c0a97137c32930029`

// DeployTokenWhitelistableExporter deploys a new Ethereum contract, binding an instance of TokenWhitelistableExporter to it.
func DeployTokenWhitelistableExporter(auth *bind.TransactOpts, backend bind.ContractBackend, _ens_ common.Address, _tokenWhitelistName_ [32]byte) (common.Address, *types.Transaction, *TokenWhitelistableExporter, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenWhitelistableExporterABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TokenWhitelistableExporterBin), backend, _ens_, _tokenWhitelistName_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TokenWhitelistableExporter{TokenWhitelistableExporterCaller: TokenWhitelistableExporterCaller{contract: contract}, TokenWhitelistableExporterTransactor: TokenWhitelistableExporterTransactor{contract: contract}, TokenWhitelistableExporterFilterer: TokenWhitelistableExporterFilterer{contract: contract}}, nil
}

// TokenWhitelistableExporter is an auto generated Go binding around an Ethereum contract.
type TokenWhitelistableExporter struct {
	TokenWhitelistableExporterCaller     // Read-only binding to the contract
	TokenWhitelistableExporterTransactor // Write-only binding to the contract
	TokenWhitelistableExporterFilterer   // Log filterer for contract events
}

// TokenWhitelistableExporterCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenWhitelistableExporterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenWhitelistableExporterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenWhitelistableExporterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenWhitelistableExporterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenWhitelistableExporterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenWhitelistableExporterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenWhitelistableExporterSession struct {
	Contract     *TokenWhitelistableExporter // Generic contract binding to set the session for
	CallOpts     bind.CallOpts               // Call options to use throughout this session
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// TokenWhitelistableExporterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenWhitelistableExporterCallerSession struct {
	Contract *TokenWhitelistableExporterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                     // Call options to use throughout this session
}

// TokenWhitelistableExporterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenWhitelistableExporterTransactorSession struct {
	Contract     *TokenWhitelistableExporterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                     // Transaction auth options to use throughout this session
}

// TokenWhitelistableExporterRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenWhitelistableExporterRaw struct {
	Contract *TokenWhitelistableExporter // Generic contract binding to access the raw methods on
}

// TokenWhitelistableExporterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenWhitelistableExporterCallerRaw struct {
	Contract *TokenWhitelistableExporterCaller // Generic read-only contract binding to access the raw methods on
}

// TokenWhitelistableExporterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenWhitelistableExporterTransactorRaw struct {
	Contract *TokenWhitelistableExporterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTokenWhitelistableExporter creates a new instance of TokenWhitelistableExporter, bound to a specific deployed contract.
func NewTokenWhitelistableExporter(address common.Address, backend bind.ContractBackend) (*TokenWhitelistableExporter, error) {
	contract, err := bindTokenWhitelistableExporter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TokenWhitelistableExporter{TokenWhitelistableExporterCaller: TokenWhitelistableExporterCaller{contract: contract}, TokenWhitelistableExporterTransactor: TokenWhitelistableExporterTransactor{contract: contract}, TokenWhitelistableExporterFilterer: TokenWhitelistableExporterFilterer{contract: contract}}, nil
}

// NewTokenWhitelistableExporterCaller creates a new read-only instance of TokenWhitelistableExporter, bound to a specific deployed contract.
func NewTokenWhitelistableExporterCaller(address common.Address, caller bind.ContractCaller) (*TokenWhitelistableExporterCaller, error) {
	contract, err := bindTokenWhitelistableExporter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenWhitelistableExporterCaller{contract: contract}, nil
}

// NewTokenWhitelistableExporterTransactor creates a new write-only instance of TokenWhitelistableExporter, bound to a specific deployed contract.
func NewTokenWhitelistableExporterTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenWhitelistableExporterTransactor, error) {
	contract, err := bindTokenWhitelistableExporter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenWhitelistableExporterTransactor{contract: contract}, nil
}

// NewTokenWhitelistableExporterFilterer creates a new log filterer instance of TokenWhitelistableExporter, bound to a specific deployed contract.
func NewTokenWhitelistableExporterFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenWhitelistableExporterFilterer, error) {
	contract, err := bindTokenWhitelistableExporter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenWhitelistableExporterFilterer{contract: contract}, nil
}

// bindTokenWhitelistableExporter binds a generic wrapper to an already deployed contract.
func bindTokenWhitelistableExporter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenWhitelistableExporterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenWhitelistableExporter *TokenWhitelistableExporterRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TokenWhitelistableExporter.Contract.TokenWhitelistableExporterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenWhitelistableExporter *TokenWhitelistableExporterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenWhitelistableExporter.Contract.TokenWhitelistableExporterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenWhitelistableExporter *TokenWhitelistableExporterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenWhitelistableExporter.Contract.TokenWhitelistableExporterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenWhitelistableExporter *TokenWhitelistableExporterCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TokenWhitelistableExporter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenWhitelistableExporter *TokenWhitelistableExporterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenWhitelistableExporter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenWhitelistableExporter *TokenWhitelistableExporterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenWhitelistableExporter.Contract.contract.Transact(opts, method, params...)
}

// EnsRegistry is a free data retrieval call binding the contract method 0x7d73b231.
//
// Solidity: function ensRegistry() constant returns(address)
func (_TokenWhitelistableExporter *TokenWhitelistableExporterCaller) EnsRegistry(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TokenWhitelistableExporter.contract.Call(opts, out, "ensRegistry")
	return *ret0, err
}

// EnsRegistry is a free data retrieval call binding the contract method 0x7d73b231.
//
// Solidity: function ensRegistry() constant returns(address)
func (_TokenWhitelistableExporter *TokenWhitelistableExporterSession) EnsRegistry() (common.Address, error) {
	return _TokenWhitelistableExporter.Contract.EnsRegistry(&_TokenWhitelistableExporter.CallOpts)
}

// EnsRegistry is a free data retrieval call binding the contract method 0x7d73b231.
//
// Solidity: function ensRegistry() constant returns(address)
func (_TokenWhitelistableExporter *TokenWhitelistableExporterCallerSession) EnsRegistry() (common.Address, error) {
	return _TokenWhitelistableExporter.Contract.EnsRegistry(&_TokenWhitelistableExporter.CallOpts)
}

// GetStablecoinInfo is a free data retrieval call binding the contract method 0x3efec5e9.
//
// Solidity: function getStablecoinInfo() constant returns(string, uint256, uint256, bool, bool, bool, uint256)
func (_TokenWhitelistableExporter *TokenWhitelistableExporterCaller) GetStablecoinInfo(opts *bind.CallOpts) (string, *big.Int, *big.Int, bool, bool, bool, *big.Int, error) {
	var (
		ret0 = new(string)
		ret1 = new(*big.Int)
		ret2 = new(*big.Int)
		ret3 = new(bool)
		ret4 = new(bool)
		ret5 = new(bool)
		ret6 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
		ret5,
		ret6,
	}
	err := _TokenWhitelistableExporter.contract.Call(opts, out, "getStablecoinInfo")
	return *ret0, *ret1, *ret2, *ret3, *ret4, *ret5, *ret6, err
}

// GetStablecoinInfo is a free data retrieval call binding the contract method 0x3efec5e9.
//
// Solidity: function getStablecoinInfo() constant returns(string, uint256, uint256, bool, bool, bool, uint256)
func (_TokenWhitelistableExporter *TokenWhitelistableExporterSession) GetStablecoinInfo() (string, *big.Int, *big.Int, bool, bool, bool, *big.Int, error) {
	return _TokenWhitelistableExporter.Contract.GetStablecoinInfo(&_TokenWhitelistableExporter.CallOpts)
}

// GetStablecoinInfo is a free data retrieval call binding the contract method 0x3efec5e9.
//
// Solidity: function getStablecoinInfo() constant returns(string, uint256, uint256, bool, bool, bool, uint256)
func (_TokenWhitelistableExporter *TokenWhitelistableExporterCallerSession) GetStablecoinInfo() (string, *big.Int, *big.Int, bool, bool, bool, *big.Int, error) {
	return _TokenWhitelistableExporter.Contract.GetStablecoinInfo(&_TokenWhitelistableExporter.CallOpts)
}

// GetTokenInfo is a free data retrieval call binding the contract method 0x1f69565f.
//
// Solidity: function getTokenInfo(address _a) constant returns(string, uint256, uint256, bool, bool, bool, uint256)
func (_TokenWhitelistableExporter *TokenWhitelistableExporterCaller) GetTokenInfo(opts *bind.CallOpts, _a common.Address) (string, *big.Int, *big.Int, bool, bool, bool, *big.Int, error) {
	var (
		ret0 = new(string)
		ret1 = new(*big.Int)
		ret2 = new(*big.Int)
		ret3 = new(bool)
		ret4 = new(bool)
		ret5 = new(bool)
		ret6 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
		ret5,
		ret6,
	}
	err := _TokenWhitelistableExporter.contract.Call(opts, out, "getTokenInfo", _a)
	return *ret0, *ret1, *ret2, *ret3, *ret4, *ret5, *ret6, err
}

// GetTokenInfo is a free data retrieval call binding the contract method 0x1f69565f.
//
// Solidity: function getTokenInfo(address _a) constant returns(string, uint256, uint256, bool, bool, bool, uint256)
func (_TokenWhitelistableExporter *TokenWhitelistableExporterSession) GetTokenInfo(_a common.Address) (string, *big.Int, *big.Int, bool, bool, bool, *big.Int, error) {
	return _TokenWhitelistableExporter.Contract.GetTokenInfo(&_TokenWhitelistableExporter.CallOpts, _a)
}

// GetTokenInfo is a free data retrieval call binding the contract method 0x1f69565f.
//
// Solidity: function getTokenInfo(address _a) constant returns(string, uint256, uint256, bool, bool, bool, uint256)
func (_TokenWhitelistableExporter *TokenWhitelistableExporterCallerSession) GetTokenInfo(_a common.Address) (string, *big.Int, *big.Int, bool, bool, bool, *big.Int, error) {
	return _TokenWhitelistableExporter.Contract.GetTokenInfo(&_TokenWhitelistableExporter.CallOpts, _a)
}

// IsTokenAvailable is a free data retrieval call binding the contract method 0x2ff0fcaa.
//
// Solidity: function isTokenAvailable(address _a) constant returns(bool)
func (_TokenWhitelistableExporter *TokenWhitelistableExporterCaller) IsTokenAvailable(opts *bind.CallOpts, _a common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TokenWhitelistableExporter.contract.Call(opts, out, "isTokenAvailable", _a)
	return *ret0, err
}

// IsTokenAvailable is a free data retrieval call binding the contract method 0x2ff0fcaa.
//
// Solidity: function isTokenAvailable(address _a) constant returns(bool)
func (_TokenWhitelistableExporter *TokenWhitelistableExporterSession) IsTokenAvailable(_a common.Address) (bool, error) {
	return _TokenWhitelistableExporter.Contract.IsTokenAvailable(&_TokenWhitelistableExporter.CallOpts, _a)
}

// IsTokenAvailable is a free data retrieval call binding the contract method 0x2ff0fcaa.
//
// Solidity: function isTokenAvailable(address _a) constant returns(bool)
func (_TokenWhitelistableExporter *TokenWhitelistableExporterCallerSession) IsTokenAvailable(_a common.Address) (bool, error) {
	return _TokenWhitelistableExporter.Contract.IsTokenAvailable(&_TokenWhitelistableExporter.CallOpts, _a)
}

// IsTokenBurnable is a free data retrieval call binding the contract method 0x0cde34b0.
//
// Solidity: function isTokenBurnable(address _a) constant returns(bool)
func (_TokenWhitelistableExporter *TokenWhitelistableExporterCaller) IsTokenBurnable(opts *bind.CallOpts, _a common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TokenWhitelistableExporter.contract.Call(opts, out, "isTokenBurnable", _a)
	return *ret0, err
}

// IsTokenBurnable is a free data retrieval call binding the contract method 0x0cde34b0.
//
// Solidity: function isTokenBurnable(address _a) constant returns(bool)
func (_TokenWhitelistableExporter *TokenWhitelistableExporterSession) IsTokenBurnable(_a common.Address) (bool, error) {
	return _TokenWhitelistableExporter.Contract.IsTokenBurnable(&_TokenWhitelistableExporter.CallOpts, _a)
}

// IsTokenBurnable is a free data retrieval call binding the contract method 0x0cde34b0.
//
// Solidity: function isTokenBurnable(address _a) constant returns(bool)
func (_TokenWhitelistableExporter *TokenWhitelistableExporterCallerSession) IsTokenBurnable(_a common.Address) (bool, error) {
	return _TokenWhitelistableExporter.Contract.IsTokenBurnable(&_TokenWhitelistableExporter.CallOpts, _a)
}

// IsTokenLoadable is a free data retrieval call binding the contract method 0x17e73794.
//
// Solidity: function isTokenLoadable(address _a) constant returns(bool)
func (_TokenWhitelistableExporter *TokenWhitelistableExporterCaller) IsTokenLoadable(opts *bind.CallOpts, _a common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TokenWhitelistableExporter.contract.Call(opts, out, "isTokenLoadable", _a)
	return *ret0, err
}

// IsTokenLoadable is a free data retrieval call binding the contract method 0x17e73794.
//
// Solidity: function isTokenLoadable(address _a) constant returns(bool)
func (_TokenWhitelistableExporter *TokenWhitelistableExporterSession) IsTokenLoadable(_a common.Address) (bool, error) {
	return _TokenWhitelistableExporter.Contract.IsTokenLoadable(&_TokenWhitelistableExporter.CallOpts, _a)
}

// IsTokenLoadable is a free data retrieval call binding the contract method 0x17e73794.
//
// Solidity: function isTokenLoadable(address _a) constant returns(bool)
func (_TokenWhitelistableExporter *TokenWhitelistableExporterCallerSession) IsTokenLoadable(_a common.Address) (bool, error) {
	return _TokenWhitelistableExporter.Contract.IsTokenLoadable(&_TokenWhitelistableExporter.CallOpts, _a)
}

// TokenAddressArray is a free data retrieval call binding the contract method 0x443dd2a4.
//
// Solidity: function tokenAddressArray() constant returns(address[])
func (_TokenWhitelistableExporter *TokenWhitelistableExporterCaller) TokenAddressArray(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _TokenWhitelistableExporter.contract.Call(opts, out, "tokenAddressArray")
	return *ret0, err
}

// TokenAddressArray is a free data retrieval call binding the contract method 0x443dd2a4.
//
// Solidity: function tokenAddressArray() constant returns(address[])
func (_TokenWhitelistableExporter *TokenWhitelistableExporterSession) TokenAddressArray() ([]common.Address, error) {
	return _TokenWhitelistableExporter.Contract.TokenAddressArray(&_TokenWhitelistableExporter.CallOpts)
}

// TokenAddressArray is a free data retrieval call binding the contract method 0x443dd2a4.
//
// Solidity: function tokenAddressArray() constant returns(address[])
func (_TokenWhitelistableExporter *TokenWhitelistableExporterCallerSession) TokenAddressArray() ([]common.Address, error) {
	return _TokenWhitelistableExporter.Contract.TokenAddressArray(&_TokenWhitelistableExporter.CallOpts)
}

// TokenWhitelistNode is a free data retrieval call binding the contract method 0x877337b0.
//
// Solidity: function tokenWhitelistNode() constant returns(bytes32)
func (_TokenWhitelistableExporter *TokenWhitelistableExporterCaller) TokenWhitelistNode(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _TokenWhitelistableExporter.contract.Call(opts, out, "tokenWhitelistNode")
	return *ret0, err
}

// TokenWhitelistNode is a free data retrieval call binding the contract method 0x877337b0.
//
// Solidity: function tokenWhitelistNode() constant returns(bytes32)
func (_TokenWhitelistableExporter *TokenWhitelistableExporterSession) TokenWhitelistNode() ([32]byte, error) {
	return _TokenWhitelistableExporter.Contract.TokenWhitelistNode(&_TokenWhitelistableExporter.CallOpts)
}

// TokenWhitelistNode is a free data retrieval call binding the contract method 0x877337b0.
//
// Solidity: function tokenWhitelistNode() constant returns(bytes32)
func (_TokenWhitelistableExporter *TokenWhitelistableExporterCallerSession) TokenWhitelistNode() ([32]byte, error) {
	return _TokenWhitelistableExporter.Contract.TokenWhitelistNode(&_TokenWhitelistableExporter.CallOpts)
}

// UpdateTokenRate is a paid mutator transaction binding the contract method 0xd545782e.
//
// Solidity: function updateTokenRate(address _token, uint256 _rate, uint256 _updateDate) returns()
func (_TokenWhitelistableExporter *TokenWhitelistableExporterTransactor) UpdateTokenRate(opts *bind.TransactOpts, _token common.Address, _rate *big.Int, _updateDate *big.Int) (*types.Transaction, error) {
	return _TokenWhitelistableExporter.contract.Transact(opts, "updateTokenRate", _token, _rate, _updateDate)
}

// UpdateTokenRate is a paid mutator transaction binding the contract method 0xd545782e.
//
// Solidity: function updateTokenRate(address _token, uint256 _rate, uint256 _updateDate) returns()
func (_TokenWhitelistableExporter *TokenWhitelistableExporterSession) UpdateTokenRate(_token common.Address, _rate *big.Int, _updateDate *big.Int) (*types.Transaction, error) {
	return _TokenWhitelistableExporter.Contract.UpdateTokenRate(&_TokenWhitelistableExporter.TransactOpts, _token, _rate, _updateDate)
}

// UpdateTokenRate is a paid mutator transaction binding the contract method 0xd545782e.
//
// Solidity: function updateTokenRate(address _token, uint256 _rate, uint256 _updateDate) returns()
func (_TokenWhitelistableExporter *TokenWhitelistableExporterTransactorSession) UpdateTokenRate(_token common.Address, _rate *big.Int, _updateDate *big.Int) (*types.Transaction, error) {
	return _TokenWhitelistableExporter.Contract.UpdateTokenRate(&_TokenWhitelistableExporter.TransactOpts, _token, _rate, _updateDate)
}
