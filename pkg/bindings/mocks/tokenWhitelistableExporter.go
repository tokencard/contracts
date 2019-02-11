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
const TokenWhitelistableExporterABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_a\",\"type\":\"address\"}],\"name\":\"isTokenLoadable\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_a\",\"type\":\"address\"}],\"name\":\"getTokenInfo\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"bool\"},{\"name\":\"\",\"type\":\"bool\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_a\",\"type\":\"address\"}],\"name\":\"isTokenAvailable\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"},{\"name\":\"_rate\",\"type\":\"uint256\"},{\"name\":\"_updateDate\",\"type\":\"uint256\"}],\"name\":\"updateTokenRate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTokenAddressArray\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_ens\",\"type\":\"address\"},{\"name\":\"_tokenWhitelistName\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// TokenWhitelistableExporterBin is the compiled bytecode used for deploying new contracts.
const TokenWhitelistableExporterBin = `608060405234801561001057600080fd5b5060405160408061095683398101604052805160209091015160008054600160a060020a031916600160a060020a039093169290921782556001556108fb90819061005b90396000f30060806040526004361061006c5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166317e7379481146100715780631f69565f146100a65780632ff0fcaa14610167578063d545782e14610188578063e81239ac146101b1575b600080fd5b34801561007d57600080fd5b50610092600160a060020a0360043516610216565b604080519115158252519081900360200190f35b3480156100b257600080fd5b506100c7600160a060020a0360043516610227565b6040805160208082018890529181018690528415156060820152831515608082015260a0810183905260c080825288519082015287519091829160e08301918a019080838360005b8381101561012757818101518382015260200161010f565b50505050905090810190601f1680156101545780820380516001836020036101000a031916815260200191505b5097505050505050505060405180910390f35b34801561017357600080fd5b50610092600160a060020a036004351661024e565b34801561019457600080fd5b506101af600160a060020a036004351660243560443561025f565b005b3480156101bd57600080fd5b506101c661026f565b60408051602080825283518183015283519192839290830191858101910280838360005b838110156102025781810151838201526020016101ea565b505050509050019250505060405180910390f35b60006102218261027e565b92915050565b6060600080600080600061023a87610296565b949c939b5091995097509550909350915050565b6000610259826104e3565b50919050565b61026a8383836104fc565b505050565b60606102796106b5565b905090565b60008061028a83610296565b50979650505050505050565b60008054600154604080517f0178b8bf000000000000000000000000000000000000000000000000000000008152600481019290925251606093928392839283928392600160a060020a0390911691630178b8bf9160248082019260209290919082900301818787803b15801561030c57600080fd5b505af1158015610320573d6000803e3d6000fd5b505050506040513d602081101561033657600080fd5b5051600154604080517f3b3b57de000000000000000000000000000000000000000000000000000000008152600481019290925251600160a060020a0390921691633b3b57de916024808201926020929091908290030181600087803b15801561039f57600080fd5b505af11580156103b3573d6000803e3d6000fd5b505050506040513d60208110156103c957600080fd5b5051604080517f1f69565f000000000000000000000000000000000000000000000000000000008152600160a060020a038a8116600483015291519190921691631f69565f91602480830192600092919082900301818387803b15801561042f57600080fd5b505af1158015610443573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405260c081101561046c57600080fd5b81019080805164010000000081111561048457600080fd5b8201602081018481111561049757600080fd5b81516401000000008111828201871017156104b157600080fd5b5050602082015160408301516060840151608085015160a090950151939e929d50909b50995091975095509350505050565b6000806104ef83610296565b5090979650505050505050565b60008054600154604080517f0178b8bf000000000000000000000000000000000000000000000000000000008152600481019290925251600160a060020a0390921692630178b8bf926024808401936020939083900390910190829087803b15801561056757600080fd5b505af115801561057b573d6000803e3d6000fd5b505050506040513d602081101561059157600080fd5b5051600154604080517f3b3b57de000000000000000000000000000000000000000000000000000000008152600481019290925251600160a060020a0390921691633b3b57de916024808201926020929091908290030181600087803b1580156105fa57600080fd5b505af115801561060e573d6000803e3d6000fd5b505050506040513d602081101561062457600080fd5b5051604080517fd545782e000000000000000000000000000000000000000000000000000000008152600160a060020a03868116600483015260248201869052604482018590529151919092169163d545782e91606480830192600092919082900301818387803b15801561069857600080fd5b505af11580156106ac573d6000803e3d6000fd5b50505050505050565b60008054600154604080517f0178b8bf000000000000000000000000000000000000000000000000000000008152600481019290925251606093600160a060020a0390931692630178b8bf92602480820193602093909283900390910190829087803b15801561072457600080fd5b505af1158015610738573d6000803e3d6000fd5b505050506040513d602081101561074e57600080fd5b5051600154604080517f3b3b57de000000000000000000000000000000000000000000000000000000008152600481019290925251600160a060020a0390921691633b3b57de916024808201926020929091908290030181600087803b1580156107b757600080fd5b505af11580156107cb573d6000803e3d6000fd5b505050506040513d60208110156107e157600080fd5b5051604080517fe81239ac0000000000000000000000000000000000000000000000000000000081529051600160a060020a039092169163e81239ac9160048082019260009290919082900301818387803b15801561083f57600080fd5b505af1158015610853573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052602081101561087c57600080fd5b81019080805164010000000081111561089457600080fd5b820160208101848111156108a757600080fd5b81518560208202830111640100000000821117156108c457600080fd5b5090945050505050905600a165627a7a723058201c8fe21525841fc1916fa0529bc3c0ebde575e088689e88b87f2b5b3dbdc25c40029`

// DeployTokenWhitelistableExporter deploys a new Ethereum contract, binding an instance of TokenWhitelistableExporter to it.
func DeployTokenWhitelistableExporter(auth *bind.TransactOpts, backend bind.ContractBackend, _ens common.Address, _tokenWhitelistName [32]byte) (common.Address, *types.Transaction, *TokenWhitelistableExporter, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenWhitelistableExporterABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TokenWhitelistableExporterBin), backend, _ens, _tokenWhitelistName)
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

// GetTokenAddressArray is a free data retrieval call binding the contract method 0xe81239ac.
//
// Solidity: function getTokenAddressArray() constant returns(address[])
func (_TokenWhitelistableExporter *TokenWhitelistableExporterCaller) GetTokenAddressArray(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _TokenWhitelistableExporter.contract.Call(opts, out, "getTokenAddressArray")
	return *ret0, err
}

// GetTokenAddressArray is a free data retrieval call binding the contract method 0xe81239ac.
//
// Solidity: function getTokenAddressArray() constant returns(address[])
func (_TokenWhitelistableExporter *TokenWhitelistableExporterSession) GetTokenAddressArray() ([]common.Address, error) {
	return _TokenWhitelistableExporter.Contract.GetTokenAddressArray(&_TokenWhitelistableExporter.CallOpts)
}

// GetTokenAddressArray is a free data retrieval call binding the contract method 0xe81239ac.
//
// Solidity: function getTokenAddressArray() constant returns(address[])
func (_TokenWhitelistableExporter *TokenWhitelistableExporterCallerSession) GetTokenAddressArray() ([]common.Address, error) {
	return _TokenWhitelistableExporter.Contract.GetTokenAddressArray(&_TokenWhitelistableExporter.CallOpts)
}

// GetTokenInfo is a free data retrieval call binding the contract method 0x1f69565f.
//
// Solidity: function getTokenInfo(address _a) constant returns(string, uint256, uint256, bool, bool, uint256)
func (_TokenWhitelistableExporter *TokenWhitelistableExporterCaller) GetTokenInfo(opts *bind.CallOpts, _a common.Address) (string, *big.Int, *big.Int, bool, bool, *big.Int, error) {
	var (
		ret0 = new(string)
		ret1 = new(*big.Int)
		ret2 = new(*big.Int)
		ret3 = new(bool)
		ret4 = new(bool)
		ret5 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
		ret5,
	}
	err := _TokenWhitelistableExporter.contract.Call(opts, out, "getTokenInfo", _a)
	return *ret0, *ret1, *ret2, *ret3, *ret4, *ret5, err
}

// GetTokenInfo is a free data retrieval call binding the contract method 0x1f69565f.
//
// Solidity: function getTokenInfo(address _a) constant returns(string, uint256, uint256, bool, bool, uint256)
func (_TokenWhitelistableExporter *TokenWhitelistableExporterSession) GetTokenInfo(_a common.Address) (string, *big.Int, *big.Int, bool, bool, *big.Int, error) {
	return _TokenWhitelistableExporter.Contract.GetTokenInfo(&_TokenWhitelistableExporter.CallOpts, _a)
}

// GetTokenInfo is a free data retrieval call binding the contract method 0x1f69565f.
//
// Solidity: function getTokenInfo(address _a) constant returns(string, uint256, uint256, bool, bool, uint256)
func (_TokenWhitelistableExporter *TokenWhitelistableExporterCallerSession) GetTokenInfo(_a common.Address) (string, *big.Int, *big.Int, bool, bool, *big.Int, error) {
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
