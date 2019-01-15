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

// TokenHolderABI is the input ABI used to generate the binding from.
const TokenHolderABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"burner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokens\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"t\",\"type\":\"address[]\"}],\"name\":\"setTokens\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"}],\"name\":\"changeOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"newOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"name\":\"balance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"burnerContract\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"}]"

// TokenHolderBin is the compiled bytecode used for deploying new contracts.
const TokenHolderBin = `608060405234801561001057600080fd5b50604051602080610746833981016040525160008054600160a060020a0319908116331790915560028054600160a060020a03909316929091169190911790556106e78061005f6000396000f3006080604052600436106100985763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166327810b6e811461009a5780634f64b2be146100cb578063625adaf2146100e357806379ba5097146101035780638da5cb5b146101185780639dc29fac1461012d578063a6f9dae114610165578063d4ee1d9014610186578063e3d670d71461019b575b005b3480156100a657600080fd5b506100af6101ce565b60408051600160a060020a039092168252519081900360200190f35b3480156100d757600080fd5b506100af6004356101dd565b3480156100ef57600080fd5b506100986004803560248101910135610205565b34801561010f57600080fd5b5061009861022f565b34801561012457600080fd5b506100af610275565b34801561013957600080fd5b50610151600160a060020a0360043516602435610284565b604080519115158252519081900360200190f35b34801561017157600080fd5b50610098600160a060020a0360043516610404565b34801561019257600080fd5b506100af61044c565b3480156101a757600080fd5b506101bc600160a060020a036004351661045b565b60408051918252519081900360200190f35b600254600160a060020a031681565b60038054829081106101eb57fe5b600091825260209091200154600160a060020a0316905081565b600054600160a060020a031633811461021d57600080fd5b61022960038484610617565b50505050565b600154600160a060020a031633811461024757600080fd5b506000805473ffffffffffffffffffffffffffffffffffffffff199081163317909155600180549091169055565b600054600160a060020a031681565b600254600090819081908190600160a060020a03163381146102a557600080fd5b8515156102b557600194506103fa565b600254604080517f771282f600000000000000000000000000000000000000000000000000000000815290518892600160a060020a03169163771282f69160048083019260209291908290030181600087803b15801561031457600080fd5b505af1158015610328573d6000803e3d6000fd5b505050506040513d602081101561033e57600080fd5b50510193508386111561035057600080fd5b600092505b6003548310156103f55761038b60038481548110151561037157fe5b600091825260209091200154600160a060020a031661045b565b915060008211156103ea5781868784028115156103a457fe5b04146103af57600080fd5b6103ea876003858154811015156103c257fe5b600091825260209091200154600160a060020a031686858a028115156103e457fe5b04610509565b600190920191610355565b600194505b5050505092915050565b600054600160a060020a031633811461041c57600080fd5b506001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600154600160a060020a031681565b6000600160a060020a0382161561050057604080517f70a082310000000000000000000000000000000000000000000000000000000081523060048201529051600160a060020a038416916370a082319160248083019260209291908290030181600087803b1580156104cd57600080fd5b505af11580156104e1573d6000803e3d6000fd5b505050506040513d60208110156104f757600080fd5b50519050610504565b5030315b919050565b80151561051557610612565b600160a060020a038216156105dc5781600160a060020a031663a9059cbb84836040518363ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018083600160a060020a0316600160a060020a0316815260200182815260200192505050602060405180830381600087803b1580156105a057600080fd5b505af11580156105b4573d6000803e3d6000fd5b505050506040513d60208110156105ca57600080fd5b505115156105d757600080fd5b610612565b604051600160a060020a0384169082156108fc029083906000818181858888f19350505050158015610229573d6000803e3d6000fd5b505050565b828054828255906000526020600020908101928215610677579160200282015b8281111561067757815473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03843516178255602090920191600190910190610637565b50610683929150610687565b5090565b6106b891905b8082111561068357805473ffffffffffffffffffffffffffffffffffffffff1916815560010161068d565b905600a165627a7a72305820104d0caca3f9f8b3da2372886ae9d63c283bb4e25590c072263ea7998d17a40b0029`

// DeployTokenHolder deploys a new Ethereum contract, binding an instance of TokenHolder to it.
func DeployTokenHolder(auth *bind.TransactOpts, backend bind.ContractBackend, burnerContract common.Address) (common.Address, *types.Transaction, *TokenHolder, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenHolderABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TokenHolderBin), backend, burnerContract)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TokenHolder{TokenHolderCaller: TokenHolderCaller{contract: contract}, TokenHolderTransactor: TokenHolderTransactor{contract: contract}, TokenHolderFilterer: TokenHolderFilterer{contract: contract}}, nil
}

// TokenHolder is an auto generated Go binding around an Ethereum contract.
type TokenHolder struct {
	TokenHolderCaller     // Read-only binding to the contract
	TokenHolderTransactor // Write-only binding to the contract
	TokenHolderFilterer   // Log filterer for contract events
}

// TokenHolderCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenHolderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenHolderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenHolderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenHolderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenHolderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenHolderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenHolderSession struct {
	Contract     *TokenHolder      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenHolderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenHolderCallerSession struct {
	Contract *TokenHolderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// TokenHolderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenHolderTransactorSession struct {
	Contract     *TokenHolderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// TokenHolderRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenHolderRaw struct {
	Contract *TokenHolder // Generic contract binding to access the raw methods on
}

// TokenHolderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenHolderCallerRaw struct {
	Contract *TokenHolderCaller // Generic read-only contract binding to access the raw methods on
}

// TokenHolderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenHolderTransactorRaw struct {
	Contract *TokenHolderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTokenHolder creates a new instance of TokenHolder, bound to a specific deployed contract.
func NewTokenHolder(address common.Address, backend bind.ContractBackend) (*TokenHolder, error) {
	contract, err := bindTokenHolder(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TokenHolder{TokenHolderCaller: TokenHolderCaller{contract: contract}, TokenHolderTransactor: TokenHolderTransactor{contract: contract}, TokenHolderFilterer: TokenHolderFilterer{contract: contract}}, nil
}

// NewTokenHolderCaller creates a new read-only instance of TokenHolder, bound to a specific deployed contract.
func NewTokenHolderCaller(address common.Address, caller bind.ContractCaller) (*TokenHolderCaller, error) {
	contract, err := bindTokenHolder(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenHolderCaller{contract: contract}, nil
}

// NewTokenHolderTransactor creates a new write-only instance of TokenHolder, bound to a specific deployed contract.
func NewTokenHolderTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenHolderTransactor, error) {
	contract, err := bindTokenHolder(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenHolderTransactor{contract: contract}, nil
}

// NewTokenHolderFilterer creates a new log filterer instance of TokenHolder, bound to a specific deployed contract.
func NewTokenHolderFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenHolderFilterer, error) {
	contract, err := bindTokenHolder(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenHolderFilterer{contract: contract}, nil
}

// bindTokenHolder binds a generic wrapper to an already deployed contract.
func bindTokenHolder(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenHolderABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenHolder *TokenHolderRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TokenHolder.Contract.TokenHolderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenHolder *TokenHolderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenHolder.Contract.TokenHolderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenHolder *TokenHolderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenHolder.Contract.TokenHolderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenHolder *TokenHolderCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TokenHolder.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenHolder *TokenHolderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenHolder.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenHolder *TokenHolderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenHolder.Contract.contract.Transact(opts, method, params...)
}

// Balance is a free data retrieval call binding the contract method 0xe3d670d7.
//
// Solidity: function balance(token address) constant returns(uint256)
func (_TokenHolder *TokenHolderCaller) Balance(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TokenHolder.contract.Call(opts, out, "balance", token)
	return *ret0, err
}

// Balance is a free data retrieval call binding the contract method 0xe3d670d7.
//
// Solidity: function balance(token address) constant returns(uint256)
func (_TokenHolder *TokenHolderSession) Balance(token common.Address) (*big.Int, error) {
	return _TokenHolder.Contract.Balance(&_TokenHolder.CallOpts, token)
}

// Balance is a free data retrieval call binding the contract method 0xe3d670d7.
//
// Solidity: function balance(token address) constant returns(uint256)
func (_TokenHolder *TokenHolderCallerSession) Balance(token common.Address) (*big.Int, error) {
	return _TokenHolder.Contract.Balance(&_TokenHolder.CallOpts, token)
}

// Burner is a free data retrieval call binding the contract method 0x27810b6e.
//
// Solidity: function burner() constant returns(address)
func (_TokenHolder *TokenHolderCaller) Burner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TokenHolder.contract.Call(opts, out, "burner")
	return *ret0, err
}

// Burner is a free data retrieval call binding the contract method 0x27810b6e.
//
// Solidity: function burner() constant returns(address)
func (_TokenHolder *TokenHolderSession) Burner() (common.Address, error) {
	return _TokenHolder.Contract.Burner(&_TokenHolder.CallOpts)
}

// Burner is a free data retrieval call binding the contract method 0x27810b6e.
//
// Solidity: function burner() constant returns(address)
func (_TokenHolder *TokenHolderCallerSession) Burner() (common.Address, error) {
	return _TokenHolder.Contract.Burner(&_TokenHolder.CallOpts)
}

// NewOwner is a free data retrieval call binding the contract method 0xd4ee1d90.
//
// Solidity: function newOwner() constant returns(address)
func (_TokenHolder *TokenHolderCaller) NewOwner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TokenHolder.contract.Call(opts, out, "newOwner")
	return *ret0, err
}

// NewOwner is a free data retrieval call binding the contract method 0xd4ee1d90.
//
// Solidity: function newOwner() constant returns(address)
func (_TokenHolder *TokenHolderSession) NewOwner() (common.Address, error) {
	return _TokenHolder.Contract.NewOwner(&_TokenHolder.CallOpts)
}

// NewOwner is a free data retrieval call binding the contract method 0xd4ee1d90.
//
// Solidity: function newOwner() constant returns(address)
func (_TokenHolder *TokenHolderCallerSession) NewOwner() (common.Address, error) {
	return _TokenHolder.Contract.NewOwner(&_TokenHolder.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_TokenHolder *TokenHolderCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TokenHolder.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_TokenHolder *TokenHolderSession) Owner() (common.Address, error) {
	return _TokenHolder.Contract.Owner(&_TokenHolder.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_TokenHolder *TokenHolderCallerSession) Owner() (common.Address, error) {
	return _TokenHolder.Contract.Owner(&_TokenHolder.CallOpts)
}

// Tokens is a free data retrieval call binding the contract method 0x4f64b2be.
//
// Solidity: function tokens( uint256) constant returns(address)
func (_TokenHolder *TokenHolderCaller) Tokens(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TokenHolder.contract.Call(opts, out, "tokens", arg0)
	return *ret0, err
}

// Tokens is a free data retrieval call binding the contract method 0x4f64b2be.
//
// Solidity: function tokens( uint256) constant returns(address)
func (_TokenHolder *TokenHolderSession) Tokens(arg0 *big.Int) (common.Address, error) {
	return _TokenHolder.Contract.Tokens(&_TokenHolder.CallOpts, arg0)
}

// Tokens is a free data retrieval call binding the contract method 0x4f64b2be.
//
// Solidity: function tokens( uint256) constant returns(address)
func (_TokenHolder *TokenHolderCallerSession) Tokens(arg0 *big.Int) (common.Address, error) {
	return _TokenHolder.Contract.Tokens(&_TokenHolder.CallOpts, arg0)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_TokenHolder *TokenHolderTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenHolder.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_TokenHolder *TokenHolderSession) AcceptOwnership() (*types.Transaction, error) {
	return _TokenHolder.Contract.AcceptOwnership(&_TokenHolder.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_TokenHolder *TokenHolderTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _TokenHolder.Contract.AcceptOwnership(&_TokenHolder.TransactOpts)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(to address, amount uint256) returns(bool)
func (_TokenHolder *TokenHolderTransactor) Burn(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TokenHolder.contract.Transact(opts, "burn", to, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(to address, amount uint256) returns(bool)
func (_TokenHolder *TokenHolderSession) Burn(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TokenHolder.Contract.Burn(&_TokenHolder.TransactOpts, to, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(to address, amount uint256) returns(bool)
func (_TokenHolder *TokenHolderTransactorSession) Burn(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TokenHolder.Contract.Burn(&_TokenHolder.TransactOpts, to, amount)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(to address) returns()
func (_TokenHolder *TokenHolderTransactor) ChangeOwner(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _TokenHolder.contract.Transact(opts, "changeOwner", to)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(to address) returns()
func (_TokenHolder *TokenHolderSession) ChangeOwner(to common.Address) (*types.Transaction, error) {
	return _TokenHolder.Contract.ChangeOwner(&_TokenHolder.TransactOpts, to)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(to address) returns()
func (_TokenHolder *TokenHolderTransactorSession) ChangeOwner(to common.Address) (*types.Transaction, error) {
	return _TokenHolder.Contract.ChangeOwner(&_TokenHolder.TransactOpts, to)
}

// SetTokens is a paid mutator transaction binding the contract method 0x625adaf2.
//
// Solidity: function setTokens(t address[]) returns()
func (_TokenHolder *TokenHolderTransactor) SetTokens(opts *bind.TransactOpts, t []common.Address) (*types.Transaction, error) {
	return _TokenHolder.contract.Transact(opts, "setTokens", t)
}

// SetTokens is a paid mutator transaction binding the contract method 0x625adaf2.
//
// Solidity: function setTokens(t address[]) returns()
func (_TokenHolder *TokenHolderSession) SetTokens(t []common.Address) (*types.Transaction, error) {
	return _TokenHolder.Contract.SetTokens(&_TokenHolder.TransactOpts, t)
}

// SetTokens is a paid mutator transaction binding the contract method 0x625adaf2.
//
// Solidity: function setTokens(t address[]) returns()
func (_TokenHolder *TokenHolderTransactorSession) SetTokens(t []common.Address) (*types.Transaction, error) {
	return _TokenHolder.Contract.SetTokens(&_TokenHolder.TransactOpts, t)
}
