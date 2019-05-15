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

// HolderABI is the input ABI used to generate the binding from.
const HolderABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"burner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tokens\",\"type\":\"address[]\"}],\"name\":\"addTokens\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tokens\",\"type\":\"address[]\"}],\"name\":\"removeTokens\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"}],\"name\":\"changeOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokenExists\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"name\":\"balance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"burnerContract\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"AddedToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"RemovedToken\",\"type\":\"event\"}]"

// HolderBin is the compiled bytecode used for deploying new contracts.
const HolderBin = `608060405234801561001057600080fd5b50604051602080610ac38339810180604052602081101561003057600080fd5b5051600080546001600160a01b03199081163317909155600280546001600160a01b0390931692909116919091179055610a548061006f6000396000f3fe60806040526004361061007b5760003560e01c80639dc29fac1161004e5780639dc29fac146101bd578063a6f9dae11461020a578063b33f78ca1461023d578063e3d670d7146102705761007b565b806327810b6e1461007d5780634ae05c7d146100ae5780636c3824ef1461012b57806379ba5097146101a8575b005b34801561008957600080fd5b506100926102b5565b604080516001600160a01b039092168252519081900360200190f35b3480156100ba57600080fd5b5061007b600480360360208110156100d157600080fd5b8101906020810181356401000000008111156100ec57600080fd5b8201836020820111156100fe57600080fd5b8035906020019184602083028401116401000000008311171561012057600080fd5b5090925090506102c5565b34801561013757600080fd5b5061007b6004803603602081101561014e57600080fd5b81019060208101813564010000000081111561016957600080fd5b82018360208201111561017b57600080fd5b8035906020019184602083028401116401000000008311171561019d57600080fd5b509092509050610424565b3480156101b457600080fd5b5061007b610622565b3480156101c957600080fd5b506101f6600480360360408110156101e057600080fd5b506001600160a01b03813516906020013561065b565b604080519115158252519081900360200190f35b34801561021657600080fd5b5061007b6004803603602081101561022d57600080fd5b50356001600160a01b03166107b7565b34801561024957600080fd5b506101f66004803603602081101561026057600080fd5b50356001600160a01b03166107f2565b34801561027c57600080fd5b506102a36004803603602081101561029357600080fd5b50356001600160a01b0316610807565b60408051918252519081900360200190f35b6002546001600160a01b03165b90565b6000546001600160a01b03163381146102dd57600080fd5b60005b8281101561041e5760008484838181106102f657fe5b602090810292909201356001600160a01b0316600081815260049093526040909220549192505060ff16156103755760408051600160e51b62461bcd02815260206004820152601460248201527f746f6b656e20616c726561647920657869737473000000000000000000000000604482015290519081900360640190fd5b6003805460018082019092557fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b0180546001600160a01b0319166001600160a01b038416908117909155600081815260046020908152604091829020805460ff191690941790935580513381529283019190915280517fb688caad5e1d3e1a0fe69a509504a4849096cc519f7c542e07b8a9dbc7993e569281900390910190a1506001016102e0565b50505050565b6000546001600160a01b031633811461043c57600080fd5b60005b8281101561041e57600084848381811061045557fe5b602090810292909201356001600160a01b0316600081815260049093526040909220549192505060ff166104d35760408051600160e51b62461bcd02815260206004820152601460248201527f746f6b656e20646f6573206e6f74206578697374000000000000000000000000604482015290519081900360640190fd5b6001600160a01b0381166000908152600460205260408120805460ff191690555b60035461050890600163ffffffff61089d16565b8110156105c157816001600160a01b03166003828154811061052657fe5b6000918252602090912001546001600160a01b031614156105b9576003805461055690600163ffffffff61089d16565b8154811061056057fe5b600091825260209091200154600380546001600160a01b03909216918390811061058657fe5b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b031602179055506105c1565b6001016104f4565b5060038054906105d59060001983016109eb565b50604080513381526001600160a01b038316602082015281517f703f7e3f084d5b8dcc12fddcfd9a70d65b6b21ec7659e4608dbaf4419ede3ad0929181900390910190a15060010161043f565b6001546001600160a01b031633811461063a57600080fd5b50600080546001600160a01b03199081163317909155600180549091169055565b6002546000906001600160a01b031633811461067657600080fd5b8261068457600191506107b0565b600083600260009054906101000a90046001600160a01b03166001600160a01b031663771282f66040518163ffffffff1660e01b815260040160206040518083038186803b1580156106d557600080fd5b505afa1580156106e9573d6000803e3d6000fd5b505050506040513d60208110156106ff57600080fd5b50510190508084111561071157600080fd5b60005b6003548110156107a957600061074a6003838154811061073057fe5b6000918252602090912001546001600160a01b0316610807565b905080156107a05780868783028161075e57fe5b041461076957600080fd5b6107a0876003848154811061077a57fe5b6000918252602090912001546001600160a01b031685848a028161079a57fe5b046108fd565b50600101610714565b5060019250505b5092915050565b6000546001600160a01b03163381146107cf57600080fd5b50600180546001600160a01b0319166001600160a01b0392909216919091179055565b60046020526000908152604090205460ff1681565b60006001600160a01b038216156108945760408051600160e01b6370a0823102815230600482015290516001600160a01b038416916370a08231916024808301926020929190829003018186803b15801561086157600080fd5b505afa158015610875573d6000803e3d6000fd5b505050506040513d602081101561088b57600080fd5b50519050610898565b5030315b919050565b6000828211156108f75760408051600160e51b62461bcd02815260206004820152601e60248201527f536166654d6174683a207375627472616374696f6e206f766572666c6f770000604482015290519081900360640190fd5b50900390565b80610907576109e6565b6001600160a01b038216156109b057816001600160a01b031663a9059cbb84836040518363ffffffff1660e01b815260040180836001600160a01b03166001600160a01b0316815260200182815260200192505050602060405180830381600087803b15801561097657600080fd5b505af115801561098a573d6000803e3d6000fd5b505050506040513d60208110156109a057600080fd5b50516109ab57600080fd5b6109e6565b6040516001600160a01b0384169082156108fc029083906000818181858888f1935050505015801561041e573d6000803e3d6000fd5b505050565b8154818355818111156109e6576000838152602090206109e69181019083016102c291905b80821115610a245760008155600101610a10565b509056fea165627a7a72305820c249a00aec226f8a1ce15cc21fe6336631677d811aa169f15999245759e3192f0029`

// DeployHolder deploys a new Ethereum contract, binding an instance of Holder to it.
func DeployHolder(auth *bind.TransactOpts, backend bind.ContractBackend, burnerContract common.Address) (common.Address, *types.Transaction, *Holder, error) {
	parsed, err := abi.JSON(strings.NewReader(HolderABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(HolderBin), backend, burnerContract)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Holder{HolderCaller: HolderCaller{contract: contract}, HolderTransactor: HolderTransactor{contract: contract}, HolderFilterer: HolderFilterer{contract: contract}}, nil
}

// Holder is an auto generated Go binding around an Ethereum contract.
type Holder struct {
	HolderCaller     // Read-only binding to the contract
	HolderTransactor // Write-only binding to the contract
	HolderFilterer   // Log filterer for contract events
}

// HolderCaller is an auto generated read-only Go binding around an Ethereum contract.
type HolderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HolderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HolderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HolderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HolderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HolderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HolderSession struct {
	Contract     *Holder           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HolderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HolderCallerSession struct {
	Contract *HolderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// HolderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HolderTransactorSession struct {
	Contract     *HolderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HolderRaw is an auto generated low-level Go binding around an Ethereum contract.
type HolderRaw struct {
	Contract *Holder // Generic contract binding to access the raw methods on
}

// HolderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HolderCallerRaw struct {
	Contract *HolderCaller // Generic read-only contract binding to access the raw methods on
}

// HolderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HolderTransactorRaw struct {
	Contract *HolderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHolder creates a new instance of Holder, bound to a specific deployed contract.
func NewHolder(address common.Address, backend bind.ContractBackend) (*Holder, error) {
	contract, err := bindHolder(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Holder{HolderCaller: HolderCaller{contract: contract}, HolderTransactor: HolderTransactor{contract: contract}, HolderFilterer: HolderFilterer{contract: contract}}, nil
}

// NewHolderCaller creates a new read-only instance of Holder, bound to a specific deployed contract.
func NewHolderCaller(address common.Address, caller bind.ContractCaller) (*HolderCaller, error) {
	contract, err := bindHolder(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HolderCaller{contract: contract}, nil
}

// NewHolderTransactor creates a new write-only instance of Holder, bound to a specific deployed contract.
func NewHolderTransactor(address common.Address, transactor bind.ContractTransactor) (*HolderTransactor, error) {
	contract, err := bindHolder(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HolderTransactor{contract: contract}, nil
}

// NewHolderFilterer creates a new log filterer instance of Holder, bound to a specific deployed contract.
func NewHolderFilterer(address common.Address, filterer bind.ContractFilterer) (*HolderFilterer, error) {
	contract, err := bindHolder(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HolderFilterer{contract: contract}, nil
}

// bindHolder binds a generic wrapper to an already deployed contract.
func bindHolder(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HolderABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Holder *HolderRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Holder.Contract.HolderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Holder *HolderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Holder.Contract.HolderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Holder *HolderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Holder.Contract.HolderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Holder *HolderCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Holder.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Holder *HolderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Holder.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Holder *HolderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Holder.Contract.contract.Transact(opts, method, params...)
}

// Balance is a free data retrieval call binding the contract method 0xe3d670d7.
//
// Solidity: function balance(address token) constant returns(uint256)
func (_Holder *HolderCaller) Balance(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Holder.contract.Call(opts, out, "balance", token)
	return *ret0, err
}

// Balance is a free data retrieval call binding the contract method 0xe3d670d7.
//
// Solidity: function balance(address token) constant returns(uint256)
func (_Holder *HolderSession) Balance(token common.Address) (*big.Int, error) {
	return _Holder.Contract.Balance(&_Holder.CallOpts, token)
}

// Balance is a free data retrieval call binding the contract method 0xe3d670d7.
//
// Solidity: function balance(address token) constant returns(uint256)
func (_Holder *HolderCallerSession) Balance(token common.Address) (*big.Int, error) {
	return _Holder.Contract.Balance(&_Holder.CallOpts, token)
}

// Burner is a free data retrieval call binding the contract method 0x27810b6e.
//
// Solidity: function burner() constant returns(address)
func (_Holder *HolderCaller) Burner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Holder.contract.Call(opts, out, "burner")
	return *ret0, err
}

// Burner is a free data retrieval call binding the contract method 0x27810b6e.
//
// Solidity: function burner() constant returns(address)
func (_Holder *HolderSession) Burner() (common.Address, error) {
	return _Holder.Contract.Burner(&_Holder.CallOpts)
}

// Burner is a free data retrieval call binding the contract method 0x27810b6e.
//
// Solidity: function burner() constant returns(address)
func (_Holder *HolderCallerSession) Burner() (common.Address, error) {
	return _Holder.Contract.Burner(&_Holder.CallOpts)
}

// TokenExists is a free data retrieval call binding the contract method 0xb33f78ca.
//
// Solidity: function tokenExists(address ) constant returns(bool)
func (_Holder *HolderCaller) TokenExists(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Holder.contract.Call(opts, out, "tokenExists", arg0)
	return *ret0, err
}

// TokenExists is a free data retrieval call binding the contract method 0xb33f78ca.
//
// Solidity: function tokenExists(address ) constant returns(bool)
func (_Holder *HolderSession) TokenExists(arg0 common.Address) (bool, error) {
	return _Holder.Contract.TokenExists(&_Holder.CallOpts, arg0)
}

// TokenExists is a free data retrieval call binding the contract method 0xb33f78ca.
//
// Solidity: function tokenExists(address ) constant returns(bool)
func (_Holder *HolderCallerSession) TokenExists(arg0 common.Address) (bool, error) {
	return _Holder.Contract.TokenExists(&_Holder.CallOpts, arg0)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Holder *HolderTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Holder.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Holder *HolderSession) AcceptOwnership() (*types.Transaction, error) {
	return _Holder.Contract.AcceptOwnership(&_Holder.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Holder *HolderTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _Holder.Contract.AcceptOwnership(&_Holder.TransactOpts)
}

// AddTokens is a paid mutator transaction binding the contract method 0x4ae05c7d.
//
// Solidity: function addTokens(address[] _tokens) returns()
func (_Holder *HolderTransactor) AddTokens(opts *bind.TransactOpts, _tokens []common.Address) (*types.Transaction, error) {
	return _Holder.contract.Transact(opts, "addTokens", _tokens)
}

// AddTokens is a paid mutator transaction binding the contract method 0x4ae05c7d.
//
// Solidity: function addTokens(address[] _tokens) returns()
func (_Holder *HolderSession) AddTokens(_tokens []common.Address) (*types.Transaction, error) {
	return _Holder.Contract.AddTokens(&_Holder.TransactOpts, _tokens)
}

// AddTokens is a paid mutator transaction binding the contract method 0x4ae05c7d.
//
// Solidity: function addTokens(address[] _tokens) returns()
func (_Holder *HolderTransactorSession) AddTokens(_tokens []common.Address) (*types.Transaction, error) {
	return _Holder.Contract.AddTokens(&_Holder.TransactOpts, _tokens)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address to, uint256 amount) returns(bool)
func (_Holder *HolderTransactor) Burn(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Holder.contract.Transact(opts, "burn", to, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address to, uint256 amount) returns(bool)
func (_Holder *HolderSession) Burn(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Holder.Contract.Burn(&_Holder.TransactOpts, to, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address to, uint256 amount) returns(bool)
func (_Holder *HolderTransactorSession) Burn(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Holder.Contract.Burn(&_Holder.TransactOpts, to, amount)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address to) returns()
func (_Holder *HolderTransactor) ChangeOwner(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _Holder.contract.Transact(opts, "changeOwner", to)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address to) returns()
func (_Holder *HolderSession) ChangeOwner(to common.Address) (*types.Transaction, error) {
	return _Holder.Contract.ChangeOwner(&_Holder.TransactOpts, to)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address to) returns()
func (_Holder *HolderTransactorSession) ChangeOwner(to common.Address) (*types.Transaction, error) {
	return _Holder.Contract.ChangeOwner(&_Holder.TransactOpts, to)
}

// RemoveTokens is a paid mutator transaction binding the contract method 0x6c3824ef.
//
// Solidity: function removeTokens(address[] _tokens) returns()
func (_Holder *HolderTransactor) RemoveTokens(opts *bind.TransactOpts, _tokens []common.Address) (*types.Transaction, error) {
	return _Holder.contract.Transact(opts, "removeTokens", _tokens)
}

// RemoveTokens is a paid mutator transaction binding the contract method 0x6c3824ef.
//
// Solidity: function removeTokens(address[] _tokens) returns()
func (_Holder *HolderSession) RemoveTokens(_tokens []common.Address) (*types.Transaction, error) {
	return _Holder.Contract.RemoveTokens(&_Holder.TransactOpts, _tokens)
}

// RemoveTokens is a paid mutator transaction binding the contract method 0x6c3824ef.
//
// Solidity: function removeTokens(address[] _tokens) returns()
func (_Holder *HolderTransactorSession) RemoveTokens(_tokens []common.Address) (*types.Transaction, error) {
	return _Holder.Contract.RemoveTokens(&_Holder.TransactOpts, _tokens)
}

// HolderAddedTokenIterator is returned from FilterAddedToken and is used to iterate over the raw logs and unpacked data for AddedToken events raised by the Holder contract.
type HolderAddedTokenIterator struct {
	Event *HolderAddedToken // Event containing the contract specifics and raw log

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
func (it *HolderAddedTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HolderAddedToken)
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
		it.Event = new(HolderAddedToken)
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
func (it *HolderAddedTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HolderAddedTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HolderAddedToken represents a AddedToken event raised by the Holder contract.
type HolderAddedToken struct {
	Sender common.Address
	Token  common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAddedToken is a free log retrieval operation binding the contract event 0xb688caad5e1d3e1a0fe69a509504a4849096cc519f7c542e07b8a9dbc7993e56.
//
// Solidity: event AddedToken(address _sender, address _token)
func (_Holder *HolderFilterer) FilterAddedToken(opts *bind.FilterOpts) (*HolderAddedTokenIterator, error) {

	logs, sub, err := _Holder.contract.FilterLogs(opts, "AddedToken")
	if err != nil {
		return nil, err
	}
	return &HolderAddedTokenIterator{contract: _Holder.contract, event: "AddedToken", logs: logs, sub: sub}, nil
}

// WatchAddedToken is a free log subscription operation binding the contract event 0xb688caad5e1d3e1a0fe69a509504a4849096cc519f7c542e07b8a9dbc7993e56.
//
// Solidity: event AddedToken(address _sender, address _token)
func (_Holder *HolderFilterer) WatchAddedToken(opts *bind.WatchOpts, sink chan<- *HolderAddedToken) (event.Subscription, error) {

	logs, sub, err := _Holder.contract.WatchLogs(opts, "AddedToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HolderAddedToken)
				if err := _Holder.contract.UnpackLog(event, "AddedToken", log); err != nil {
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

// HolderRemovedTokenIterator is returned from FilterRemovedToken and is used to iterate over the raw logs and unpacked data for RemovedToken events raised by the Holder contract.
type HolderRemovedTokenIterator struct {
	Event *HolderRemovedToken // Event containing the contract specifics and raw log

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
func (it *HolderRemovedTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HolderRemovedToken)
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
		it.Event = new(HolderRemovedToken)
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
func (it *HolderRemovedTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HolderRemovedTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HolderRemovedToken represents a RemovedToken event raised by the Holder contract.
type HolderRemovedToken struct {
	Sender common.Address
	Token  common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRemovedToken is a free log retrieval operation binding the contract event 0x703f7e3f084d5b8dcc12fddcfd9a70d65b6b21ec7659e4608dbaf4419ede3ad0.
//
// Solidity: event RemovedToken(address _sender, address _token)
func (_Holder *HolderFilterer) FilterRemovedToken(opts *bind.FilterOpts) (*HolderRemovedTokenIterator, error) {

	logs, sub, err := _Holder.contract.FilterLogs(opts, "RemovedToken")
	if err != nil {
		return nil, err
	}
	return &HolderRemovedTokenIterator{contract: _Holder.contract, event: "RemovedToken", logs: logs, sub: sub}, nil
}

// WatchRemovedToken is a free log subscription operation binding the contract event 0x703f7e3f084d5b8dcc12fddcfd9a70d65b6b21ec7659e4608dbaf4419ede3ad0.
//
// Solidity: event RemovedToken(address _sender, address _token)
func (_Holder *HolderFilterer) WatchRemovedToken(opts *bind.WatchOpts, sink chan<- *HolderRemovedToken) (event.Subscription, error) {

	logs, sub, err := _Holder.contract.WatchLogs(opts, "RemovedToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HolderRemovedToken)
				if err := _Holder.contract.UnpackLog(event, "RemovedToken", log); err != nil {
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
