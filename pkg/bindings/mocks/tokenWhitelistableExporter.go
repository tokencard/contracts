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
const TokenWhitelistableExporterABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_a\",\"type\":\"address\"}],\"name\":\"isTokenBurnable\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_a\",\"type\":\"address\"}],\"name\":\"isTokenLoadable\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_a\",\"type\":\"address\"}],\"name\":\"getTokenInfo\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"bool\"},{\"name\":\"\",\"type\":\"bool\"},{\"name\":\"\",\"type\":\"bool\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_a\",\"type\":\"address\"}],\"name\":\"isTokenAvailable\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getStablecoinInfo\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"bool\"},{\"name\":\"\",\"type\":\"bool\"},{\"name\":\"\",\"type\":\"bool\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tokenAddressArray\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ensRegistry\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tokenWhitelistNode\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"burnableTokens\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"},{\"name\":\"_rate\",\"type\":\"uint256\"},{\"name\":\"_updateDate\",\"type\":\"uint256\"}],\"name\":\"updateTokenRate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_ens_\",\"type\":\"address\"},{\"name\":\"_tokenWhitelistName_\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// TokenWhitelistableExporterBin is the compiled bytecode used for deploying new contracts.
const TokenWhitelistableExporterBin = `608060405234801561001057600080fd5b5060405161095d38038061095d8339818101604052604081101561003357600080fd5b508051602090910151600180546001600160a01b039384166001600160a01b031991821617918290556000805490911691909316179091556002556108e08061007d6000396000f3fe608060405234801561001057600080fd5b506004361061009e5760003560e01c8063443dd2a411610066578063443dd2a4146102015780637d73b23114610259578063877337b01461027d578063bf25811f14610297578063d545782e1461029f5761009e565b80630cde34b0146100a357806317e73794146100dd5780631f69565f146101035780632ff0fcaa146101d35780633efec5e9146101f9575b600080fd5b6100c9600480360360208110156100b957600080fd5b50356001600160a01b03166102d3565b604080519115158252519081900360200190f35b6100c9600480360360208110156100f357600080fd5b50356001600160a01b03166102e4565b6101296004803603602081101561011957600080fd5b50356001600160a01b03166102ef565b6040805160208082018990529181018790528515156060820152841515608082015283151560a082015260c0810183905260e08082528951908201528851909182916101008301918b019080838360005b8381101561019257818101518382015260200161017a565b50505050905090810190601f1680156101bf5780820380516001836020036101000a031916815260200191505b509850505050505050505060405180910390f35b6100c9600480360360208110156101e957600080fd5b50356001600160a01b031661031a565b610129610325565b61020961034e565b60408051602080825283518183015283519192839290830191858101910280838360005b8381101561024557818101518382015260200161022d565b505050509050019250505060405180910390f35b61026161035d565b604080516001600160a01b039092168252519081900360200190f35b61028561036c565b60408051918252519081900360200190f35b610209610372565b6102d1600480360360608110156102b557600080fd5b506001600160a01b03813516906020810135906040013561037c565b005b60006102de8261038c565b92915050565b60006102de826103a5565b6060600080600080600080610303886103bf565b959e949d50929b5090995097509550909350915050565b60006102de826104ec565b6060600080600080600080610338610506565b959d949c50929a50909850965094509092509050565b606061035861060e565b905090565b6001546001600160a01b031690565b60025490565b60606103586106e3565b610387838383610728565b505050565b600080610398836103bf565b5098975050505050505050565b6000806103b1836103bf565b509098975050505050505050565b60606000806000806000806103d56002546107b7565b6001600160a01b0316631f69565f896040518263ffffffff1660e01b815260040180826001600160a01b03166001600160a01b0316815260200191505060006040518083038186803b15801561042a57600080fd5b505afa15801561043e573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405260e081101561046757600080fd5b81019080805164010000000081111561047f57600080fd5b8201602081018481111561049257600080fd5b81516401000000008111828201871017156104ac57600080fd5b5050602082015160408301516060840151608085015160a086015160c090960151949e50929c50909a509850965090945092505050919395979092949650565b6000806104f8836103bf565b509198975050505050505050565b606060008060008060008061051c6002546107b7565b6001600160a01b0316633efec5e96040518163ffffffff1660e01b815260040160006040518083038186803b15801561055457600080fd5b505afa158015610568573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405260e081101561059157600080fd5b8101908080516401000000008111156105a957600080fd5b820160208101848111156105bc57600080fd5b81516401000000008111828201871017156105d657600080fd5b5050602082015160408301516060840151608085015160a086015160c090960151949f939e50919c509a509850919650945092505050565b606061061b6002546107b7565b6001600160a01b031663443dd2a46040518163ffffffff1660e01b815260040160006040518083038186803b15801561065357600080fd5b505afa158015610667573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052602081101561069057600080fd5b8101908080516401000000008111156106a857600080fd5b820160208101848111156106bb57600080fd5b81518560208202830111640100000000821117156106d857600080fd5b509094505050505090565b60606106f06002546107b7565b6001600160a01b031663bf25811f6040518163ffffffff1660e01b815260040160006040518083038186803b15801561065357600080fd5b6107336002546107b7565b6001600160a01b031663d545782e8484846040518463ffffffff1660e01b815260040180846001600160a01b03166001600160a01b031681526020018381526020018281526020019350505050600060405180830381600087803b15801561079a57600080fd5b505af11580156107ae573d6000803e3d6000fd5b50505050505050565b6000805460408051630178b8bf60e01b81526004810185905290516001600160a01b0390921691630178b8bf91602480820192602092909190829003018186803b15801561080457600080fd5b505afa158015610818573d6000803e3d6000fd5b505050506040513d602081101561082e57600080fd5b505160408051631d9dabef60e11b81526004810185905290516001600160a01b0390921691633b3b57de91602480820192602092909190829003018186803b15801561087957600080fd5b505afa15801561088d573d6000803e3d6000fd5b505050506040513d60208110156108a357600080fd5b50519291505056fea265627a7a723058203bb5903bc7f6566a16c2da19075d8402f731561bb19f4206888dfe063b593fa064736f6c634300050a0032`

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

// BurnableTokens is a free data retrieval call binding the contract method 0xbf25811f.
//
// Solidity: function burnableTokens() constant returns(address[])
func (_TokenWhitelistableExporter *TokenWhitelistableExporterCaller) BurnableTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _TokenWhitelistableExporter.contract.Call(opts, out, "burnableTokens")
	return *ret0, err
}

// BurnableTokens is a free data retrieval call binding the contract method 0xbf25811f.
//
// Solidity: function burnableTokens() constant returns(address[])
func (_TokenWhitelistableExporter *TokenWhitelistableExporterSession) BurnableTokens() ([]common.Address, error) {
	return _TokenWhitelistableExporter.Contract.BurnableTokens(&_TokenWhitelistableExporter.CallOpts)
}

// BurnableTokens is a free data retrieval call binding the contract method 0xbf25811f.
//
// Solidity: function burnableTokens() constant returns(address[])
func (_TokenWhitelistableExporter *TokenWhitelistableExporterCallerSession) BurnableTokens() ([]common.Address, error) {
	return _TokenWhitelistableExporter.Contract.BurnableTokens(&_TokenWhitelistableExporter.CallOpts)
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
