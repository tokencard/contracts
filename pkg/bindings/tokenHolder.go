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
const TokenHolderABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"poolShare\",\"type\":\"uint256\"}],\"name\":\"payOut\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"burner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokens\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"t\",\"type\":\"address[]\"}],\"name\":\"setTokens\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"}],\"name\":\"changeOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"newOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"name\":\"balance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"burnerContract\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"}]"

// TokenHolderBin is the compiled bytecode used for deploying new contracts.
const TokenHolderBin = `608060405234801561001057600080fd5b5060405160208061068d833981016040525160008054600160a060020a0319908116331790915560028054600160a060020a039093169290911691909117905561062e8061005f6000396000f3006080604052600436106100985763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416630221038a811461009a57806327810b6e146100d25780634f64b2be14610103578063625adaf21461011b57806379ba50971461013b5780638da5cb5b14610150578063a6f9dae114610165578063d4ee1d9014610186578063e3d670d71461019b575b005b3480156100a657600080fd5b506100be600160a060020a03600435166024356101ce565b604080519115158252519081900360200190f35b3480156100de57600080fd5b506100e7610295565b60408051600160a060020a039092168252519081900360200190f35b34801561010f57600080fd5b506100e76004356102a4565b34801561012757600080fd5b5061009860048035602481019101356102cc565b34801561014757600080fd5b506100986102f6565b34801561015c57600080fd5b506100e761033c565b34801561017157600080fd5b50610098600160a060020a036004351661034b565b34801561019257600080fd5b506100e7610393565b3480156101a757600080fd5b506101bc600160a060020a03600435166103a2565b60408051918252519081900360200190f35b60025460009081908190600160a060020a03163381146101ed57600080fd5b8415156101fd576001935061028c565b600092505b6003548310156102875761023860038481548110151561021e57fe5b600091825260209091200154600160a060020a03166103a2565b9150600082111561027c5761027c8660038581548110151561025657fe5b600091825260209091200154600160a060020a0316878581151561027657fe5b04610450565b600190920191610202565b600193505b50505092915050565b600254600160a060020a031681565b60038054829081106102b257fe5b600091825260209091200154600160a060020a0316905081565b600054600160a060020a03163381146102e457600080fd5b6102f06003848461055e565b50505050565b600154600160a060020a031633811461030e57600080fd5b506000805473ffffffffffffffffffffffffffffffffffffffff199081163317909155600180549091169055565b600054600160a060020a031681565b600054600160a060020a031633811461036357600080fd5b506001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600154600160a060020a031681565b6000600160a060020a0382161561044757604080517f70a082310000000000000000000000000000000000000000000000000000000081523060048201529051600160a060020a038416916370a082319160248083019260209291908290030181600087803b15801561041457600080fd5b505af1158015610428573d6000803e3d6000fd5b505050506040513d602081101561043e57600080fd5b5051905061044b565b5030315b919050565b80151561045c57610559565b600160a060020a038216156105235781600160a060020a031663a9059cbb84836040518363ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018083600160a060020a0316600160a060020a0316815260200182815260200192505050602060405180830381600087803b1580156104e757600080fd5b505af11580156104fb573d6000803e3d6000fd5b505050506040513d602081101561051157600080fd5b5051151561051e57600080fd5b610559565b604051600160a060020a0384169082156108fc029083906000818181858888f193505050501580156102f0573d6000803e3d6000fd5b505050565b8280548282559060005260206000209081019282156105be579160200282015b828111156105be57815473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0384351617825560209092019160019091019061057e565b506105ca9291506105ce565b5090565b6105ff91905b808211156105ca57805473ffffffffffffffffffffffffffffffffffffffff191681556001016105d4565b905600a165627a7a723058201a9599596c1a74105114789f9af63f7cb322f34325ae100014481aae00e84cd90029`

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

// PayOut is a paid mutator transaction binding the contract method 0x0221038a.
//
// Solidity: function payOut(to address, poolShare uint256) returns(bool)
func (_TokenHolder *TokenHolderTransactor) PayOut(opts *bind.TransactOpts, to common.Address, poolShare *big.Int) (*types.Transaction, error) {
	return _TokenHolder.contract.Transact(opts, "payOut", to, poolShare)
}

// PayOut is a paid mutator transaction binding the contract method 0x0221038a.
//
// Solidity: function payOut(to address, poolShare uint256) returns(bool)
func (_TokenHolder *TokenHolderSession) PayOut(to common.Address, poolShare *big.Int) (*types.Transaction, error) {
	return _TokenHolder.Contract.PayOut(&_TokenHolder.TransactOpts, to, poolShare)
}

// PayOut is a paid mutator transaction binding the contract method 0x0221038a.
//
// Solidity: function payOut(to address, poolShare uint256) returns(bool)
func (_TokenHolder *TokenHolderTransactorSession) PayOut(to common.Address, poolShare *big.Int) (*types.Transaction, error) {
	return _TokenHolder.Contract.PayOut(&_TokenHolder.TransactOpts, to, poolShare)
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
