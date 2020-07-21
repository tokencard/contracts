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

// BurnerTokenABI is the input ABI used to generate the binding from.
const BurnerTokenABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"result\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseApproval\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseApproval\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_th\",\"type\":\"address\"}],\"name\":\"setTokenHolder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenholder\",\"outputs\":[{\"internalType\":\"contractTokenHolder\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// BurnerTokenBin is the compiled bytecode used for deploying new contracts.
var BurnerTokenBin = "0x608060405234801561001057600080fd5b50600180546001600160a01b0319163317905560408051808201909152600c8082526b26b7b737b634ba34102a25a760a11b602090920191825261005691600291610093565b506003805460ff1916600817815560408051808201909152818152622a25a760e91b602090910190815261008d9160049190610093565b5061012e565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106100d457805160ff1916838001178555610101565b82800160010185558215610101579182015b828111156101015782518255916020019190600101906100e6565b5061010d929150610111565b5090565b61012b91905b8082111561010d5760008155600101610117565b90565b610c538061013d6000396000f3fe608060405234801561001057600080fd5b506004361061010b5760003560e01c806370a08231116100a257806395d89b411161007157806395d89b411461030c578063a9059cbb14610314578063d73dd62314610340578063dd62ed3e1461036c578063f29d2f281461039a5761010b565b806370a08231146102b2578063771282f6146102d857806384eba00c146102e05780638da5cb5b146103045761010b565b8063313ce567116100de578063313ce5671461021d57806340c10f191461023b57806342966c681461026957806366188463146102865761010b565b806306fdde0314610110578063095ea7b31461018d57806318160ddd146101cd57806323b872dd146101e7575b600080fd5b6101186103c0565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561015257818101518382015260200161013a565b50505050905090810190601f16801561017f5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6101b9600480360360408110156101a357600080fd5b506001600160a01b03813516906020013561044b565b604080519115158252519081900360200190f35b6101d56104f3565b60408051918252519081900360200190f35b6101b9600480360360608110156101fd57600080fd5b506001600160a01b038135811691602081013590911690604001356104f9565b610225610652565b6040805160ff9092168252519081900360200190f35b6102676004803603604081101561025157600080fd5b506001600160a01b03813516906020013561065b565b005b6101b96004803603602081101561027f57600080fd5b50356106e8565b6101b96004803603604081101561029c57600080fd5b506001600160a01b038135169060200135610808565b6101d5600480360360208110156102c857600080fd5b50356001600160a01b03166108f8565b6101d561090a565b6102e8610910565b604080516001600160a01b039092168252519081900360200190f35b6102e861091f565b61011861092e565b6101b96004803603604081101561032a57600080fd5b506001600160a01b038135169060200135610989565b6101b96004803603604081101561035657600080fd5b506001600160a01b038135169060200135610a5c565b6101d56004803603604081101561038257600080fd5b506001600160a01b0381358116916020013516610a90565b610267600480360360208110156103b057600080fd5b50356001600160a01b0316610aad565b6002805460408051602060018416156101000260001901909316849004601f810184900484028201840190925281815292918301828280156104435780601f1061041857610100808354040283529160200191610443565b820191906000526020600020905b81548152906001019060200180831161042657829003601f168201915b505050505081565b6000811580159061047e57503360009081526006602090815260408083206001600160a01b038716845290915290205415155b1561048b575060006104ed565b3360008181526006602090815260408083206001600160a01b03881680855290835292819020869055805186815290519293927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929181900390910190a35060015b92915050565b60005490565b60006001600160a01b0383166105115750600061064b565b6001600160a01b0384166000908152600560205260409020548211156105395750600061064b565b6001600160a01b03841660009081526006602090815260408083203384529091529020548281101561056f57600091505061064b565b6001600160a01b038416600090815260056020526040902054610598908463ffffffff610acf16565b6001600160a01b0380861660009081526005602052604080822093909355908716815220546105cd908463ffffffff610b2916565b6001600160a01b0386166000908152600560205260409020556105f6818463ffffffff610b2916565b6001600160a01b0380871660008181526006602090815260408083203384528252918290209490945580518781529051928816939192600080516020610bfe833981519152929181900390910190a360019150505b9392505050565b60035460ff1681565b6001600160a01b038216600090815260056020526040902054610684908263ffffffff610acf16565b6001600160a01b038316600090815260056020526040812091909155546106b1908263ffffffff610acf16565b60009081556040805183815290516001600160a01b0385169291600080516020610bfe833981519152919081900360200190a35050565b3360009081526005602052604081205482111561070757506000610803565b33600090815260056020526040902054610727908363ffffffff610b2916565b336000908152600560205260408120919091555461074b908363ffffffff610b2916565b600090815560075460408051632770a7eb60e21b81523360048201526024810186905290516001600160a01b0390921692639dc29fac926044808401936020939083900390910190829087803b1580156107a457600080fd5b505af11580156107b8573d6000803e3d6000fd5b505050506040513d60208110156107ce57600080fd5b50519050806107dc57600080fd5b6040805183815290516000913391600080516020610bfe8339815191529181900360200190a35b919050565b3360009081526006602090815260408083206001600160a01b03861684529091528120548083111561085d573360009081526006602090815260408083206001600160a01b0388168452909152812055610892565b61086d818463ffffffff610b2916565b3360009081526006602090815260408083206001600160a01b03891684529091529020555b3360008181526006602090815260408083206001600160a01b0389168085529083529281902054815190815290519293927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929181900390910190a35060019392505050565b60056020526000908152604090205481565b60005481565b6007546001600160a01b031681565b6001546001600160a01b031681565b6004805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156104435780601f1061041857610100808354040283529160200191610443565b336000908152600560205260408120548211156109a8575060006104ed565b6001600160a01b0383166109be575060006104ed565b336000908152600560205260409020546109de908363ffffffff610b2916565b33600090815260056020526040808220929092556001600160a01b03851681522054610a10908363ffffffff610acf16565b6001600160a01b038416600081815260056020908152604091829020939093558051858152905191923392600080516020610bfe8339815191529281900390910190a350600192915050565b3360009081526006602090815260408083206001600160a01b038616845290915281205461086d818463ffffffff610acf16565b600660209081526000928352604080842090915290825290205481565b600780546001600160a01b0319166001600160a01b0392909216919091179055565b60008282018381101561064b576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b600061064b83836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f77000081525060008184841115610bf55760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b83811015610bba578181015183820152602001610ba2565b50505050905090810190601f168015610be75780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b50505090039056feddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3efa264697066735822122035dd362abec1c829d8a213367b32ed188c66d6174d8e0fff52da3d8b74d5b31564736f6c634300060b0033"

// DeployBurnerToken deploys a new Ethereum contract, binding an instance of BurnerToken to it.
func DeployBurnerToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BurnerToken, error) {
	parsed, err := abi.JSON(strings.NewReader(BurnerTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BurnerTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BurnerToken{BurnerTokenCaller: BurnerTokenCaller{contract: contract}, BurnerTokenTransactor: BurnerTokenTransactor{contract: contract}, BurnerTokenFilterer: BurnerTokenFilterer{contract: contract}}, nil
}

// BurnerToken is an auto generated Go binding around an Ethereum contract.
type BurnerToken struct {
	BurnerTokenCaller     // Read-only binding to the contract
	BurnerTokenTransactor // Write-only binding to the contract
	BurnerTokenFilterer   // Log filterer for contract events
}

// BurnerTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type BurnerTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BurnerTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BurnerTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BurnerTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BurnerTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BurnerTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BurnerTokenSession struct {
	Contract     *BurnerToken      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BurnerTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BurnerTokenCallerSession struct {
	Contract *BurnerTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// BurnerTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BurnerTokenTransactorSession struct {
	Contract     *BurnerTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// BurnerTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type BurnerTokenRaw struct {
	Contract *BurnerToken // Generic contract binding to access the raw methods on
}

// BurnerTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BurnerTokenCallerRaw struct {
	Contract *BurnerTokenCaller // Generic read-only contract binding to access the raw methods on
}

// BurnerTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BurnerTokenTransactorRaw struct {
	Contract *BurnerTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBurnerToken creates a new instance of BurnerToken, bound to a specific deployed contract.
func NewBurnerToken(address common.Address, backend bind.ContractBackend) (*BurnerToken, error) {
	contract, err := bindBurnerToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BurnerToken{BurnerTokenCaller: BurnerTokenCaller{contract: contract}, BurnerTokenTransactor: BurnerTokenTransactor{contract: contract}, BurnerTokenFilterer: BurnerTokenFilterer{contract: contract}}, nil
}

// NewBurnerTokenCaller creates a new read-only instance of BurnerToken, bound to a specific deployed contract.
func NewBurnerTokenCaller(address common.Address, caller bind.ContractCaller) (*BurnerTokenCaller, error) {
	contract, err := bindBurnerToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BurnerTokenCaller{contract: contract}, nil
}

// NewBurnerTokenTransactor creates a new write-only instance of BurnerToken, bound to a specific deployed contract.
func NewBurnerTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*BurnerTokenTransactor, error) {
	contract, err := bindBurnerToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BurnerTokenTransactor{contract: contract}, nil
}

// NewBurnerTokenFilterer creates a new log filterer instance of BurnerToken, bound to a specific deployed contract.
func NewBurnerTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*BurnerTokenFilterer, error) {
	contract, err := bindBurnerToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BurnerTokenFilterer{contract: contract}, nil
}

// bindBurnerToken binds a generic wrapper to an already deployed contract.
func bindBurnerToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BurnerTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BurnerToken *BurnerTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BurnerToken.Contract.BurnerTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BurnerToken *BurnerTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BurnerToken.Contract.BurnerTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BurnerToken *BurnerTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BurnerToken.Contract.BurnerTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BurnerToken *BurnerTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BurnerToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BurnerToken *BurnerTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BurnerToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BurnerToken *BurnerTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BurnerToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) constant returns(uint256)
func (_BurnerToken *BurnerTokenCaller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BurnerToken.contract.Call(opts, out, "allowance", arg0, arg1)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) constant returns(uint256)
func (_BurnerToken *BurnerTokenSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _BurnerToken.Contract.Allowance(&_BurnerToken.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) constant returns(uint256)
func (_BurnerToken *BurnerTokenCallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _BurnerToken.Contract.Allowance(&_BurnerToken.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) constant returns(uint256)
func (_BurnerToken *BurnerTokenCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BurnerToken.contract.Call(opts, out, "balanceOf", arg0)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) constant returns(uint256)
func (_BurnerToken *BurnerTokenSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _BurnerToken.Contract.BalanceOf(&_BurnerToken.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) constant returns(uint256)
func (_BurnerToken *BurnerTokenCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _BurnerToken.Contract.BalanceOf(&_BurnerToken.CallOpts, arg0)
}

// CurrentSupply is a free data retrieval call binding the contract method 0x771282f6.
//
// Solidity: function currentSupply() constant returns(uint256)
func (_BurnerToken *BurnerTokenCaller) CurrentSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BurnerToken.contract.Call(opts, out, "currentSupply")
	return *ret0, err
}

// CurrentSupply is a free data retrieval call binding the contract method 0x771282f6.
//
// Solidity: function currentSupply() constant returns(uint256)
func (_BurnerToken *BurnerTokenSession) CurrentSupply() (*big.Int, error) {
	return _BurnerToken.Contract.CurrentSupply(&_BurnerToken.CallOpts)
}

// CurrentSupply is a free data retrieval call binding the contract method 0x771282f6.
//
// Solidity: function currentSupply() constant returns(uint256)
func (_BurnerToken *BurnerTokenCallerSession) CurrentSupply() (*big.Int, error) {
	return _BurnerToken.Contract.CurrentSupply(&_BurnerToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_BurnerToken *BurnerTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _BurnerToken.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_BurnerToken *BurnerTokenSession) Decimals() (uint8, error) {
	return _BurnerToken.Contract.Decimals(&_BurnerToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_BurnerToken *BurnerTokenCallerSession) Decimals() (uint8, error) {
	return _BurnerToken.Contract.Decimals(&_BurnerToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_BurnerToken *BurnerTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _BurnerToken.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_BurnerToken *BurnerTokenSession) Name() (string, error) {
	return _BurnerToken.Contract.Name(&_BurnerToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_BurnerToken *BurnerTokenCallerSession) Name() (string, error) {
	return _BurnerToken.Contract.Name(&_BurnerToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BurnerToken *BurnerTokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BurnerToken.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BurnerToken *BurnerTokenSession) Owner() (common.Address, error) {
	return _BurnerToken.Contract.Owner(&_BurnerToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BurnerToken *BurnerTokenCallerSession) Owner() (common.Address, error) {
	return _BurnerToken.Contract.Owner(&_BurnerToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_BurnerToken *BurnerTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _BurnerToken.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_BurnerToken *BurnerTokenSession) Symbol() (string, error) {
	return _BurnerToken.Contract.Symbol(&_BurnerToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_BurnerToken *BurnerTokenCallerSession) Symbol() (string, error) {
	return _BurnerToken.Contract.Symbol(&_BurnerToken.CallOpts)
}

// Tokenholder is a free data retrieval call binding the contract method 0x84eba00c.
//
// Solidity: function tokenholder() constant returns(address)
func (_BurnerToken *BurnerTokenCaller) Tokenholder(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BurnerToken.contract.Call(opts, out, "tokenholder")
	return *ret0, err
}

// Tokenholder is a free data retrieval call binding the contract method 0x84eba00c.
//
// Solidity: function tokenholder() constant returns(address)
func (_BurnerToken *BurnerTokenSession) Tokenholder() (common.Address, error) {
	return _BurnerToken.Contract.Tokenholder(&_BurnerToken.CallOpts)
}

// Tokenholder is a free data retrieval call binding the contract method 0x84eba00c.
//
// Solidity: function tokenholder() constant returns(address)
func (_BurnerToken *BurnerTokenCallerSession) Tokenholder() (common.Address, error) {
	return _BurnerToken.Contract.Tokenholder(&_BurnerToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_BurnerToken *BurnerTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BurnerToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_BurnerToken *BurnerTokenSession) TotalSupply() (*big.Int, error) {
	return _BurnerToken.Contract.TotalSupply(&_BurnerToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_BurnerToken *BurnerTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _BurnerToken.Contract.TotalSupply(&_BurnerToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool success)
func (_BurnerToken *BurnerTokenTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BurnerToken.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool success)
func (_BurnerToken *BurnerTokenSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BurnerToken.Contract.Approve(&_BurnerToken.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool success)
func (_BurnerToken *BurnerTokenTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BurnerToken.Contract.Approve(&_BurnerToken.TransactOpts, _spender, _value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 _amount) returns(bool result)
func (_BurnerToken *BurnerTokenTransactor) Burn(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _BurnerToken.contract.Transact(opts, "burn", _amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 _amount) returns(bool result)
func (_BurnerToken *BurnerTokenSession) Burn(_amount *big.Int) (*types.Transaction, error) {
	return _BurnerToken.Contract.Burn(&_BurnerToken.TransactOpts, _amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 _amount) returns(bool result)
func (_BurnerToken *BurnerTokenTransactorSession) Burn(_amount *big.Int) (*types.Transaction, error) {
	return _BurnerToken.Contract.Burn(&_BurnerToken.TransactOpts, _amount)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(address _spender, uint256 _subtractedValue) returns(bool success)
func (_BurnerToken *BurnerTokenTransactor) DecreaseApproval(opts *bind.TransactOpts, _spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _BurnerToken.contract.Transact(opts, "decreaseApproval", _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(address _spender, uint256 _subtractedValue) returns(bool success)
func (_BurnerToken *BurnerTokenSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _BurnerToken.Contract.DecreaseApproval(&_BurnerToken.TransactOpts, _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(address _spender, uint256 _subtractedValue) returns(bool success)
func (_BurnerToken *BurnerTokenTransactorSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _BurnerToken.Contract.DecreaseApproval(&_BurnerToken.TransactOpts, _spender, _subtractedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(address _spender, uint256 _addedValue) returns(bool success)
func (_BurnerToken *BurnerTokenTransactor) IncreaseApproval(opts *bind.TransactOpts, _spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _BurnerToken.contract.Transact(opts, "increaseApproval", _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(address _spender, uint256 _addedValue) returns(bool success)
func (_BurnerToken *BurnerTokenSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _BurnerToken.Contract.IncreaseApproval(&_BurnerToken.TransactOpts, _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(address _spender, uint256 _addedValue) returns(bool success)
func (_BurnerToken *BurnerTokenTransactorSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _BurnerToken.Contract.IncreaseApproval(&_BurnerToken.TransactOpts, _spender, _addedValue)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address addr, uint256 amount) returns()
func (_BurnerToken *BurnerTokenTransactor) Mint(opts *bind.TransactOpts, addr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnerToken.contract.Transact(opts, "mint", addr, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address addr, uint256 amount) returns()
func (_BurnerToken *BurnerTokenSession) Mint(addr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnerToken.Contract.Mint(&_BurnerToken.TransactOpts, addr, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address addr, uint256 amount) returns()
func (_BurnerToken *BurnerTokenTransactorSession) Mint(addr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnerToken.Contract.Mint(&_BurnerToken.TransactOpts, addr, amount)
}

// SetTokenHolder is a paid mutator transaction binding the contract method 0xf29d2f28.
//
// Solidity: function setTokenHolder(address _th) returns()
func (_BurnerToken *BurnerTokenTransactor) SetTokenHolder(opts *bind.TransactOpts, _th common.Address) (*types.Transaction, error) {
	return _BurnerToken.contract.Transact(opts, "setTokenHolder", _th)
}

// SetTokenHolder is a paid mutator transaction binding the contract method 0xf29d2f28.
//
// Solidity: function setTokenHolder(address _th) returns()
func (_BurnerToken *BurnerTokenSession) SetTokenHolder(_th common.Address) (*types.Transaction, error) {
	return _BurnerToken.Contract.SetTokenHolder(&_BurnerToken.TransactOpts, _th)
}

// SetTokenHolder is a paid mutator transaction binding the contract method 0xf29d2f28.
//
// Solidity: function setTokenHolder(address _th) returns()
func (_BurnerToken *BurnerTokenTransactorSession) SetTokenHolder(_th common.Address) (*types.Transaction, error) {
	return _BurnerToken.Contract.SetTokenHolder(&_BurnerToken.TransactOpts, _th)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool success)
func (_BurnerToken *BurnerTokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BurnerToken.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool success)
func (_BurnerToken *BurnerTokenSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BurnerToken.Contract.Transfer(&_BurnerToken.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool success)
func (_BurnerToken *BurnerTokenTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BurnerToken.Contract.Transfer(&_BurnerToken.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool success)
func (_BurnerToken *BurnerTokenTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BurnerToken.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool success)
func (_BurnerToken *BurnerTokenSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BurnerToken.Contract.TransferFrom(&_BurnerToken.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool success)
func (_BurnerToken *BurnerTokenTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BurnerToken.Contract.TransferFrom(&_BurnerToken.TransactOpts, _from, _to, _value)
}

// BurnerTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the BurnerToken contract.
type BurnerTokenApprovalIterator struct {
	Event *BurnerTokenApproval // Event containing the contract specifics and raw log

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
func (it *BurnerTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnerTokenApproval)
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
		it.Event = new(BurnerTokenApproval)
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
func (it *BurnerTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BurnerTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BurnerTokenApproval represents a Approval event raised by the BurnerToken contract.
type BurnerTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_BurnerToken *BurnerTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*BurnerTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _BurnerToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &BurnerTokenApprovalIterator{contract: _BurnerToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_BurnerToken *BurnerTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *BurnerTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _BurnerToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BurnerTokenApproval)
				if err := _BurnerToken.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_BurnerToken *BurnerTokenFilterer) ParseApproval(log types.Log) (*BurnerTokenApproval, error) {
	event := new(BurnerTokenApproval)
	if err := _BurnerToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BurnerTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the BurnerToken contract.
type BurnerTokenTransferIterator struct {
	Event *BurnerTokenTransfer // Event containing the contract specifics and raw log

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
func (it *BurnerTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnerTokenTransfer)
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
		it.Event = new(BurnerTokenTransfer)
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
func (it *BurnerTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BurnerTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BurnerTokenTransfer represents a Transfer event raised by the BurnerToken contract.
type BurnerTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_BurnerToken *BurnerTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BurnerTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BurnerToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BurnerTokenTransferIterator{contract: _BurnerToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_BurnerToken *BurnerTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *BurnerTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BurnerToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BurnerTokenTransfer)
				if err := _BurnerToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_BurnerToken *BurnerTokenFilterer) ParseTransfer(log types.Log) (*BurnerTokenTransfer, error) {
	event := new(BurnerTokenTransfer)
	if err := _BurnerToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}
