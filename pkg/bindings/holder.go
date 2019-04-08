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
const HolderBin = `608060405234801561001057600080fd5b50604051602080610a47833981016040525160008054600160a060020a0319908116331790915560028054600160a060020a03909316929091169190911790556109e88061005f6000396000f30060806040526004361061008d5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166327810b6e811461008f5780634ae05c7d146100c05780636c3824ef146100e057806379ba5097146101005780639dc29fac14610115578063a6f9dae11461014d578063b33f78ca1461016e578063e3d670d71461018f575b005b34801561009b57600080fd5b506100a46101c2565b60408051600160a060020a039092168252519081900360200190f35b3480156100cc57600080fd5b5061008d60048035602481019101356101d2565b3480156100ec57600080fd5b5061008d600480356024810191013561035b565b34801561010c57600080fd5b5061008d610587565b34801561012157600080fd5b50610139600160a060020a03600435166024356105cd565b604080519115158252519081900360200190f35b34801561015957600080fd5b5061008d600160a060020a036004351661074d565b34801561017a57600080fd5b50610139600160a060020a0360043516610795565b34801561019b57600080fd5b506101b0600160a060020a03600435166107aa565b60408051918252519081900360200190f35b600254600160a060020a03165b90565b600080548190600160a060020a03163381146101ed57600080fd5b600092505b838310156103545784848481811061020657fe5b60209081029290920135600160a060020a0316600081815260049093526040909220549193505060ff161561029c57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f746f6b656e20616c726561647920657869737473000000000000000000000000604482015290519081900360640190fd5b6003805460018082019092557fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b01805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a038516908117909155600081815260046020908152604091829020805460ff191690941790935580513381529283019190915280517fb688caad5e1d3e1a0fe69a509504a4849096cc519f7c542e07b8a9dbc7993e569281900390910190a16001909201916101f2565b5050505050565b6000805481908190600160a060020a031633811461037857600080fd5b600093505b8484101561057f5785858581811061039157fe5b60209081029290920135600160a060020a0316600081815260049093526040909220549194505060ff16151561042857604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f746f6b656e20646f6573206e6f74206578697374000000000000000000000000604482015290519081900360640190fd5b600160a060020a0383166000908152600460205260408120805460ff1916905591505b60035461045f90600163ffffffff61085816565b82101561051d5782600160a060020a031660038381548110151561047f57fe5b600091825260209091200154600160a060020a0316141561051257600380546104af90600163ffffffff61085816565b815481106104b957fe5b60009182526020909120015460038054600160a060020a0390921691849081106104df57fe5b9060005260206000200160006101000a815481600160a060020a030219169083600160a060020a0316021790555061051d565b60019091019061044b565b600380549061053090600019830161097f565b5060408051338152600160a060020a038516602082015281517f703f7e3f084d5b8dcc12fddcfd9a70d65b6b21ec7659e4608dbaf4419ede3ad0929181900390910190a160019093019261037d565b505050505050565b600154600160a060020a031633811461059f57600080fd5b506000805473ffffffffffffffffffffffffffffffffffffffff199081163317909155600180549091169055565b600254600090819081908190600160a060020a03163381146105ee57600080fd5b8515156105fe5760019450610743565b600254604080517f771282f600000000000000000000000000000000000000000000000000000000815290518892600160a060020a03169163771282f69160048083019260209291908290030181600087803b15801561065d57600080fd5b505af1158015610671573d6000803e3d6000fd5b505050506040513d602081101561068757600080fd5b50510193508386111561069957600080fd5b600092505b60035483101561073e576106d46003848154811015156106ba57fe5b600091825260209091200154600160a060020a03166107aa565b915060008211156107335781868784028115156106ed57fe5b04146106f857600080fd5b6107338760038581548110151561070b57fe5b600091825260209091200154600160a060020a031686858a0281151561072d57fe5b0461086f565b60019092019161069e565b600194505b5050505092915050565b600054600160a060020a031633811461076557600080fd5b506001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b60046020526000908152604090205460ff1681565b6000600160a060020a0382161561084f57604080517f70a082310000000000000000000000000000000000000000000000000000000081523060048201529051600160a060020a038416916370a082319160248083019260209291908290030181600087803b15801561081c57600080fd5b505af1158015610830573d6000803e3d6000fd5b505050506040513d602081101561084657600080fd5b50519050610853565b5030315b919050565b6000808383111561086857600080fd5b5050900390565b80151561087b5761097a565b600160a060020a038216156109425781600160a060020a031663a9059cbb84836040518363ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018083600160a060020a0316600160a060020a0316815260200182815260200192505050602060405180830381600087803b15801561090657600080fd5b505af115801561091a573d6000803e3d6000fd5b505050506040513d602081101561093057600080fd5b5051151561093d57600080fd5b61097a565b604051600160a060020a0384169082156108fc029083906000818181858888f19350505050158015610978573d6000803e3d6000fd5b505b505050565b81548183558181111561097a5760008381526020902061097a9181019083016101cf91905b808211156109b857600081556001016109a4565b50905600a165627a7a7230582093cf4344198987565144bcb0b0034008f21832afaf5c73355fbd7f9e57179d9b0029`

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
