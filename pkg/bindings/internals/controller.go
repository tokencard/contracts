// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package internals

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

// ControllerABI is the input ABI used to generate the binding from.
const ControllerABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"controllerCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"removeAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isTransferable\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"isAdmin\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"adminCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"addAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_asset\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"claim\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"addController\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_account\",\"type\":\"address\"},{\"name\":\"_transferable\",\"type\":\"bool\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"isController\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"removeController\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_ownerAddress_\",\"type\":\"address\"},{\"name\":\"_transferable_\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_controller\",\"type\":\"address\"}],\"name\":\"AddedController\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_controller\",\"type\":\"address\"}],\"name\":\"RemovedController\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_admin\",\"type\":\"address\"}],\"name\":\"AddedAdmin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_admin\",\"type\":\"address\"}],\"name\":\"RemovedAdmin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_asset\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"Claimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"TransferredOwnership\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_locked\",\"type\":\"address\"}],\"name\":\"LockedOwnership\",\"type\":\"event\"}]"

// ControllerBin is the compiled bytecode used for deploying new contracts.
const ControllerBin = `608060405234801561001057600080fd5b506040516040806110688339810180604052604081101561003057600080fd5b508051602090910151600080546001600160a01b0319166001600160a01b03841617600160a01b60ff02191674010000000000000000000000000000000000000000831515810291909117918290558391839160ff9104166100c957604080516001600160a01b038416815290517f808639ff9c8e4732d60b6c2330de498035416d229f27a77d259680895efec1229181900360200190a15b60408051600081526001600160a01b038416602082015281517f850b3df64837d7d518b45f5aa64d104652c3b80eb5b34a8e3d9eb666cb7cdea5929181900390910190a150505050610f48806101206000396000f3fe608060405234801561001057600080fd5b50600436106100cf5760003560e01c8063715018a61161008c578063a7fc7a0711610066578063a7fc7a07146101e8578063b242e5341461020e578063b429afeb1461023c578063f6a74ed714610262576100cf565b8063715018a6146101865780638da5cb5b1461018e578063996cba68146101b2576100cf565b806315b9a8b8146100d45780631785f53c146100ee5780632121dc751461011657806324d7806c146101325780632b7832b3146101585780637048027514610160575b600080fd5b6100dc610288565b60408051918252519081900360200190f35b6101146004803603602081101561010457600080fd5b50356001600160a01b031661028e565b005b61011e6102e5565b604080519115158252519081900360200190f35b61011e6004803603602081101561014857600080fd5b50356001600160a01b03166102f5565b6100dc610313565b6101146004803603602081101561017657600080fd5b50356001600160a01b0316610319565b61011461036d565b610196610466565b604080516001600160a01b039092168252519081900360200190f35b610114600480360360608110156101c857600080fd5b506001600160a01b03813581169160208101359091169060400135610475565b610114600480360360208110156101fe57600080fd5b50356001600160a01b03166104d0565b6101146004803603604081101561022457600080fd5b506001600160a01b0381351690602001351515610536565b61011e6004803603602081101561025257600080fd5b50356001600160a01b03166106ff565b6101146004803603602081101561027857600080fd5b50356001600160a01b031661071d565b60045490565b61029733610783565b6102d95760408051600160e51b62461bcd0281526020600482015260166024820152600080516020610e45833981519152604482015290519081900360640190fd5b6102e281610797565b50565b600054600160a01b900460ff1690565b6001600160a01b031660009081526001602052604090205460ff1690565b60025490565b61032233610783565b6103645760408051600160e51b62461bcd0281526020600482015260166024820152600080516020610e45833981519152604482015290519081900360640190fd5b6102e281610870565b61037633610783565b6103b85760408051600160e51b62461bcd0281526020600482015260166024820152600080516020610e45833981519152604482015290519081900360640190fd5b600054600160a01b900460ff166104195760408051600160e51b62461bcd02815260206004820152601d60248201527f6f776e657273686970206973206e6f74207472616e7366657261626c65000000604482015290519081900360640190fd5b600080546001600160a01b031916815560408051828152602081019290925280517f850b3df64837d7d518b45f5aa64d104652c3b80eb5b34a8e3d9eb666cb7cdea59281900390910190a1565b6000546001600160a01b031690565b61047e33610783565b6104c05760408051600160e51b62461bcd0281526020600482015260166024820152600080516020610e45833981519152604482015290519081900360640190fd5b6104cb838383610a21565b505050565b6104d9336102f5565b61052d5760408051600160e51b62461bcd02815260206004820152601660248201527f73656e646572206973206e6f7420616e2061646d696e00000000000000000000604482015290519081900360640190fd5b6102e281610b84565b61053f33610783565b6105815760408051600160e51b62461bcd0281526020600482015260166024820152600080516020610e45833981519152604482015290519081900360640190fd5b600054600160a01b900460ff166105e25760408051600160e51b62461bcd02815260206004820152601d60248201527f6f776e657273686970206973206e6f74207472616e7366657261626c65000000604482015290519081900360640190fd5b6001600160a01b03821661062a57604051600160e51b62461bcd028152600401808060200182810382526023815260200180610e896023913960400191505060405180910390fd5b6000805474ff00000000000000000000000000000000000000001916600160a01b831515021790558061069457604080516001600160a01b038416815290517f808639ff9c8e4732d60b6c2330de498035416d229f27a77d259680895efec1229181900360200190a15b600054604080516001600160a01b039283168152918416602083015280517f850b3df64837d7d518b45f5aa64d104652c3b80eb5b34a8e3d9eb666cb7cdea59281900390910190a150600080546001600160a01b0319166001600160a01b0392909216919091179055565b6001600160a01b031660009081526003602052604090205460ff1690565b610726336102f5565b61077a5760408051600160e51b62461bcd02815260206004820152601660248201527f73656e646572206973206e6f7420616e2061646d696e00000000000000000000604482015290519081900360640190fd5b6102e281610d38565b6000546001600160a01b0390811691161490565b6001600160a01b03811660009081526001602052604090205460ff166108075760408051600160e51b62461bcd02815260206004820181905260248201527f70726f7669646564206163636f756e74206973206e6f7420616e2061646d696e604482015290519081900360640190fd5b6001600160a01b038116600081815260016020908152604091829020805460ff191690556002805460001901905581513381529081019290925280517f787a2e12f4a55b658b8f573c32432ee11a5e8b51677d1e1e937aaf6a0bb5776e9281900390910190a150565b6001600160a01b03811660009081526001602052604090205460ff16156108cb57604051600160e51b62461bcd028152600401808060200182810382526024815260200180610e216024913960400191505060405180910390fd5b6001600160a01b03811660009081526003602052604090205460ff161561092657604051600160e51b62461bcd028152600401808060200182810382526028815260200180610ef56028913960400191505060405180910390fd5b61092f81610783565b1561096e57604051600160e51b62461bcd028152600401808060200182810382526025815260200180610dfc6025913960400191505060405180910390fd5b6001600160a01b0381166109b657604051600160e51b62461bcd028152600401808060200182810382526024815260200180610ed16024913960400191505060405180910390fd5b6001600160a01b038116600081815260016020818152604092839020805460ff1916831790556002805490920190915581513381529081019290925280517fc58b647b8ba5a8cab2f11f32673636cc1061324240972ed05e8cc005b81a4b7a9281900390910190a150565b6001600160a01b038216610a6b576040516001600160a01b0384169082156108fc029083906000818181858888f19350505050158015610a65573d6000803e3d6000fd5b50610b35565b816001600160a01b031663a9059cbb84836040518363ffffffff1660e01b815260040180836001600160a01b03166001600160a01b0316815260200182815260200192505050602060405180830381600087803b158015610acb57600080fd5b505af1158015610adf573d6000803e3d6000fd5b505050506040513d6020811015610af557600080fd5b5051610b3557604051600160e51b62461bcd028152600401808060200182810382526025815260200180610eac6025913960400191505060405180910390fd5b604080516001600160a01b0380861682528416602082015280820183905290517ff7a40077ff7a04c7e61f6f26fb13774259ddf1b6bce9ecf26a8276cdd39926839181900360600190a1505050565b6001600160a01b03811660009081526001602052604090205460ff1615610bdf57604051600160e51b62461bcd028152600401808060200182810382526024815260200180610e216024913960400191505060405180910390fd5b6001600160a01b03811660009081526003602052604090205460ff1615610c3a57604051600160e51b62461bcd028152600401808060200182810382526028815260200180610ef56028913960400191505060405180910390fd5b610c4381610783565b15610c8257604051600160e51b62461bcd028152600401808060200182810382526025815260200180610dfc6025913960400191505060405180910390fd5b6001600160a01b038116610cca57604051600160e51b62461bcd028152600401808060200182810382526024815260200180610ed16024913960400191505060405180910390fd5b6001600160a01b038116600081815260036020908152604091829020805460ff1916600190811790915560048054909101905581513381529081019290925280517fb890d5abdcd5c2b61ce8bbc2cf6af9b6d7f7451830cbc85037cbdd182c86fe1d9281900390910190a150565b6001600160a01b03811660009081526003602052604090205460ff16610d9257604051600160e51b62461bcd028152600401808060200182810382526024815260200180610e656024913960400191505060405180910390fd5b6001600160a01b038116600081815260036020908152604091829020805460ff191690556004805460001901905581513381529081019290925280517fb6a283aaede08e15ef55c74e3014e30eb0c0040d4b156cccb77391268ea373949281900390910190a15056fe70726f7669646564206163636f756e7420697320616c726561647920746865206f776e657270726f7669646564206163636f756e7420697320616c726561647920616e2061646d696e73656e646572206973206e6f7420616e206f776e65720000000000000000000070726f7669646564206163636f756e74206973206e6f74206120636f6e74726f6c6c65726f776e65722063616e6e6f742062652073657420746f207a65726f2061646472657373455243323020746f6b656e207472616e736665722077617320756e7375636365737366756c70726f7669646564206163636f756e7420697320746865207a65726f206164647265737370726f7669646564206163636f756e7420697320616c7265616479206120636f6e74726f6c6c6572a165627a7a723058207a5768e28fc354e0f74e8d034b1a287c1b299b093b516ad167c39b6c637ba0820029`

// DeployController deploys a new Ethereum contract, binding an instance of Controller to it.
func DeployController(auth *bind.TransactOpts, backend bind.ContractBackend, _ownerAddress_ common.Address, _transferable_ bool) (common.Address, *types.Transaction, *Controller, error) {
	parsed, err := abi.JSON(strings.NewReader(ControllerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ControllerBin), backend, _ownerAddress_, _transferable_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Controller{ControllerCaller: ControllerCaller{contract: contract}, ControllerTransactor: ControllerTransactor{contract: contract}, ControllerFilterer: ControllerFilterer{contract: contract}}, nil
}

// Controller is an auto generated Go binding around an Ethereum contract.
type Controller struct {
	ControllerCaller     // Read-only binding to the contract
	ControllerTransactor // Write-only binding to the contract
	ControllerFilterer   // Log filterer for contract events
}

// ControllerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ControllerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ControllerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ControllerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ControllerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ControllerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ControllerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ControllerSession struct {
	Contract     *Controller       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ControllerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ControllerCallerSession struct {
	Contract *ControllerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ControllerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ControllerTransactorSession struct {
	Contract     *ControllerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ControllerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ControllerRaw struct {
	Contract *Controller // Generic contract binding to access the raw methods on
}

// ControllerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ControllerCallerRaw struct {
	Contract *ControllerCaller // Generic read-only contract binding to access the raw methods on
}

// ControllerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ControllerTransactorRaw struct {
	Contract *ControllerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewController creates a new instance of Controller, bound to a specific deployed contract.
func NewController(address common.Address, backend bind.ContractBackend) (*Controller, error) {
	contract, err := bindController(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Controller{ControllerCaller: ControllerCaller{contract: contract}, ControllerTransactor: ControllerTransactor{contract: contract}, ControllerFilterer: ControllerFilterer{contract: contract}}, nil
}

// NewControllerCaller creates a new read-only instance of Controller, bound to a specific deployed contract.
func NewControllerCaller(address common.Address, caller bind.ContractCaller) (*ControllerCaller, error) {
	contract, err := bindController(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ControllerCaller{contract: contract}, nil
}

// NewControllerTransactor creates a new write-only instance of Controller, bound to a specific deployed contract.
func NewControllerTransactor(address common.Address, transactor bind.ContractTransactor) (*ControllerTransactor, error) {
	contract, err := bindController(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ControllerTransactor{contract: contract}, nil
}

// NewControllerFilterer creates a new log filterer instance of Controller, bound to a specific deployed contract.
func NewControllerFilterer(address common.Address, filterer bind.ContractFilterer) (*ControllerFilterer, error) {
	contract, err := bindController(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ControllerFilterer{contract: contract}, nil
}

// bindController binds a generic wrapper to an already deployed contract.
func bindController(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ControllerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Controller *ControllerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Controller.Contract.ControllerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Controller *ControllerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Controller.Contract.ControllerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Controller *ControllerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Controller.Contract.ControllerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Controller *ControllerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Controller.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Controller *ControllerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Controller.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Controller *ControllerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Controller.Contract.contract.Transact(opts, method, params...)
}

// AdminCount is a free data retrieval call binding the contract method 0x2b7832b3.
//
// Solidity: function adminCount() constant returns(uint256)
func (_Controller *ControllerCaller) AdminCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Controller.contract.Call(opts, out, "adminCount")
	return *ret0, err
}

// AdminCount is a free data retrieval call binding the contract method 0x2b7832b3.
//
// Solidity: function adminCount() constant returns(uint256)
func (_Controller *ControllerSession) AdminCount() (*big.Int, error) {
	return _Controller.Contract.AdminCount(&_Controller.CallOpts)
}

// AdminCount is a free data retrieval call binding the contract method 0x2b7832b3.
//
// Solidity: function adminCount() constant returns(uint256)
func (_Controller *ControllerCallerSession) AdminCount() (*big.Int, error) {
	return _Controller.Contract.AdminCount(&_Controller.CallOpts)
}

// ControllerCount is a free data retrieval call binding the contract method 0x15b9a8b8.
//
// Solidity: function controllerCount() constant returns(uint256)
func (_Controller *ControllerCaller) ControllerCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Controller.contract.Call(opts, out, "controllerCount")
	return *ret0, err
}

// ControllerCount is a free data retrieval call binding the contract method 0x15b9a8b8.
//
// Solidity: function controllerCount() constant returns(uint256)
func (_Controller *ControllerSession) ControllerCount() (*big.Int, error) {
	return _Controller.Contract.ControllerCount(&_Controller.CallOpts)
}

// ControllerCount is a free data retrieval call binding the contract method 0x15b9a8b8.
//
// Solidity: function controllerCount() constant returns(uint256)
func (_Controller *ControllerCallerSession) ControllerCount() (*big.Int, error) {
	return _Controller.Contract.ControllerCount(&_Controller.CallOpts)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address _account) constant returns(bool)
func (_Controller *ControllerCaller) IsAdmin(opts *bind.CallOpts, _account common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Controller.contract.Call(opts, out, "isAdmin", _account)
	return *ret0, err
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address _account) constant returns(bool)
func (_Controller *ControllerSession) IsAdmin(_account common.Address) (bool, error) {
	return _Controller.Contract.IsAdmin(&_Controller.CallOpts, _account)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address _account) constant returns(bool)
func (_Controller *ControllerCallerSession) IsAdmin(_account common.Address) (bool, error) {
	return _Controller.Contract.IsAdmin(&_Controller.CallOpts, _account)
}

// IsController is a free data retrieval call binding the contract method 0xb429afeb.
//
// Solidity: function isController(address _account) constant returns(bool)
func (_Controller *ControllerCaller) IsController(opts *bind.CallOpts, _account common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Controller.contract.Call(opts, out, "isController", _account)
	return *ret0, err
}

// IsController is a free data retrieval call binding the contract method 0xb429afeb.
//
// Solidity: function isController(address _account) constant returns(bool)
func (_Controller *ControllerSession) IsController(_account common.Address) (bool, error) {
	return _Controller.Contract.IsController(&_Controller.CallOpts, _account)
}

// IsController is a free data retrieval call binding the contract method 0xb429afeb.
//
// Solidity: function isController(address _account) constant returns(bool)
func (_Controller *ControllerCallerSession) IsController(_account common.Address) (bool, error) {
	return _Controller.Contract.IsController(&_Controller.CallOpts, _account)
}

// IsTransferable is a free data retrieval call binding the contract method 0x2121dc75.
//
// Solidity: function isTransferable() constant returns(bool)
func (_Controller *ControllerCaller) IsTransferable(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Controller.contract.Call(opts, out, "isTransferable")
	return *ret0, err
}

// IsTransferable is a free data retrieval call binding the contract method 0x2121dc75.
//
// Solidity: function isTransferable() constant returns(bool)
func (_Controller *ControllerSession) IsTransferable() (bool, error) {
	return _Controller.Contract.IsTransferable(&_Controller.CallOpts)
}

// IsTransferable is a free data retrieval call binding the contract method 0x2121dc75.
//
// Solidity: function isTransferable() constant returns(bool)
func (_Controller *ControllerCallerSession) IsTransferable() (bool, error) {
	return _Controller.Contract.IsTransferable(&_Controller.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Controller *ControllerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Controller.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Controller *ControllerSession) Owner() (common.Address, error) {
	return _Controller.Contract.Owner(&_Controller.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Controller *ControllerCallerSession) Owner() (common.Address, error) {
	return _Controller.Contract.Owner(&_Controller.CallOpts)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address _account) returns()
func (_Controller *ControllerTransactor) AddAdmin(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "addAdmin", _account)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address _account) returns()
func (_Controller *ControllerSession) AddAdmin(_account common.Address) (*types.Transaction, error) {
	return _Controller.Contract.AddAdmin(&_Controller.TransactOpts, _account)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address _account) returns()
func (_Controller *ControllerTransactorSession) AddAdmin(_account common.Address) (*types.Transaction, error) {
	return _Controller.Contract.AddAdmin(&_Controller.TransactOpts, _account)
}

// AddController is a paid mutator transaction binding the contract method 0xa7fc7a07.
//
// Solidity: function addController(address _account) returns()
func (_Controller *ControllerTransactor) AddController(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "addController", _account)
}

// AddController is a paid mutator transaction binding the contract method 0xa7fc7a07.
//
// Solidity: function addController(address _account) returns()
func (_Controller *ControllerSession) AddController(_account common.Address) (*types.Transaction, error) {
	return _Controller.Contract.AddController(&_Controller.TransactOpts, _account)
}

// AddController is a paid mutator transaction binding the contract method 0xa7fc7a07.
//
// Solidity: function addController(address _account) returns()
func (_Controller *ControllerTransactorSession) AddController(_account common.Address) (*types.Transaction, error) {
	return _Controller.Contract.AddController(&_Controller.TransactOpts, _account)
}

// Claim is a paid mutator transaction binding the contract method 0x996cba68.
//
// Solidity: function claim(address _to, address _asset, uint256 _amount) returns()
func (_Controller *ControllerTransactor) Claim(opts *bind.TransactOpts, _to common.Address, _asset common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "claim", _to, _asset, _amount)
}

// Claim is a paid mutator transaction binding the contract method 0x996cba68.
//
// Solidity: function claim(address _to, address _asset, uint256 _amount) returns()
func (_Controller *ControllerSession) Claim(_to common.Address, _asset common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Controller.Contract.Claim(&_Controller.TransactOpts, _to, _asset, _amount)
}

// Claim is a paid mutator transaction binding the contract method 0x996cba68.
//
// Solidity: function claim(address _to, address _asset, uint256 _amount) returns()
func (_Controller *ControllerTransactorSession) Claim(_to common.Address, _asset common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Controller.Contract.Claim(&_Controller.TransactOpts, _to, _asset, _amount)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address _account) returns()
func (_Controller *ControllerTransactor) RemoveAdmin(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "removeAdmin", _account)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address _account) returns()
func (_Controller *ControllerSession) RemoveAdmin(_account common.Address) (*types.Transaction, error) {
	return _Controller.Contract.RemoveAdmin(&_Controller.TransactOpts, _account)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address _account) returns()
func (_Controller *ControllerTransactorSession) RemoveAdmin(_account common.Address) (*types.Transaction, error) {
	return _Controller.Contract.RemoveAdmin(&_Controller.TransactOpts, _account)
}

// RemoveController is a paid mutator transaction binding the contract method 0xf6a74ed7.
//
// Solidity: function removeController(address _account) returns()
func (_Controller *ControllerTransactor) RemoveController(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "removeController", _account)
}

// RemoveController is a paid mutator transaction binding the contract method 0xf6a74ed7.
//
// Solidity: function removeController(address _account) returns()
func (_Controller *ControllerSession) RemoveController(_account common.Address) (*types.Transaction, error) {
	return _Controller.Contract.RemoveController(&_Controller.TransactOpts, _account)
}

// RemoveController is a paid mutator transaction binding the contract method 0xf6a74ed7.
//
// Solidity: function removeController(address _account) returns()
func (_Controller *ControllerTransactorSession) RemoveController(_account common.Address) (*types.Transaction, error) {
	return _Controller.Contract.RemoveController(&_Controller.TransactOpts, _account)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Controller *ControllerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Controller *ControllerSession) RenounceOwnership() (*types.Transaction, error) {
	return _Controller.Contract.RenounceOwnership(&_Controller.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Controller *ControllerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Controller.Contract.RenounceOwnership(&_Controller.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xb242e534.
//
// Solidity: function transferOwnership(address _account, bool _transferable) returns()
func (_Controller *ControllerTransactor) TransferOwnership(opts *bind.TransactOpts, _account common.Address, _transferable bool) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "transferOwnership", _account, _transferable)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xb242e534.
//
// Solidity: function transferOwnership(address _account, bool _transferable) returns()
func (_Controller *ControllerSession) TransferOwnership(_account common.Address, _transferable bool) (*types.Transaction, error) {
	return _Controller.Contract.TransferOwnership(&_Controller.TransactOpts, _account, _transferable)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xb242e534.
//
// Solidity: function transferOwnership(address _account, bool _transferable) returns()
func (_Controller *ControllerTransactorSession) TransferOwnership(_account common.Address, _transferable bool) (*types.Transaction, error) {
	return _Controller.Contract.TransferOwnership(&_Controller.TransactOpts, _account, _transferable)
}

// ControllerAddedAdminIterator is returned from FilterAddedAdmin and is used to iterate over the raw logs and unpacked data for AddedAdmin events raised by the Controller contract.
type ControllerAddedAdminIterator struct {
	Event *ControllerAddedAdmin // Event containing the contract specifics and raw log

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
func (it *ControllerAddedAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ControllerAddedAdmin)
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
		it.Event = new(ControllerAddedAdmin)
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
func (it *ControllerAddedAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ControllerAddedAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ControllerAddedAdmin represents a AddedAdmin event raised by the Controller contract.
type ControllerAddedAdmin struct {
	Sender common.Address
	Admin  common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAddedAdmin is a free log retrieval operation binding the contract event 0xc58b647b8ba5a8cab2f11f32673636cc1061324240972ed05e8cc005b81a4b7a.
//
// Solidity: event AddedAdmin(address _sender, address _admin)
func (_Controller *ControllerFilterer) FilterAddedAdmin(opts *bind.FilterOpts) (*ControllerAddedAdminIterator, error) {

	logs, sub, err := _Controller.contract.FilterLogs(opts, "AddedAdmin")
	if err != nil {
		return nil, err
	}
	return &ControllerAddedAdminIterator{contract: _Controller.contract, event: "AddedAdmin", logs: logs, sub: sub}, nil
}

// WatchAddedAdmin is a free log subscription operation binding the contract event 0xc58b647b8ba5a8cab2f11f32673636cc1061324240972ed05e8cc005b81a4b7a.
//
// Solidity: event AddedAdmin(address _sender, address _admin)
func (_Controller *ControllerFilterer) WatchAddedAdmin(opts *bind.WatchOpts, sink chan<- *ControllerAddedAdmin) (event.Subscription, error) {

	logs, sub, err := _Controller.contract.WatchLogs(opts, "AddedAdmin")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ControllerAddedAdmin)
				if err := _Controller.contract.UnpackLog(event, "AddedAdmin", log); err != nil {
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

// ControllerAddedControllerIterator is returned from FilterAddedController and is used to iterate over the raw logs and unpacked data for AddedController events raised by the Controller contract.
type ControllerAddedControllerIterator struct {
	Event *ControllerAddedController // Event containing the contract specifics and raw log

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
func (it *ControllerAddedControllerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ControllerAddedController)
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
		it.Event = new(ControllerAddedController)
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
func (it *ControllerAddedControllerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ControllerAddedControllerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ControllerAddedController represents a AddedController event raised by the Controller contract.
type ControllerAddedController struct {
	Sender     common.Address
	Controller common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAddedController is a free log retrieval operation binding the contract event 0xb890d5abdcd5c2b61ce8bbc2cf6af9b6d7f7451830cbc85037cbdd182c86fe1d.
//
// Solidity: event AddedController(address _sender, address _controller)
func (_Controller *ControllerFilterer) FilterAddedController(opts *bind.FilterOpts) (*ControllerAddedControllerIterator, error) {

	logs, sub, err := _Controller.contract.FilterLogs(opts, "AddedController")
	if err != nil {
		return nil, err
	}
	return &ControllerAddedControllerIterator{contract: _Controller.contract, event: "AddedController", logs: logs, sub: sub}, nil
}

// WatchAddedController is a free log subscription operation binding the contract event 0xb890d5abdcd5c2b61ce8bbc2cf6af9b6d7f7451830cbc85037cbdd182c86fe1d.
//
// Solidity: event AddedController(address _sender, address _controller)
func (_Controller *ControllerFilterer) WatchAddedController(opts *bind.WatchOpts, sink chan<- *ControllerAddedController) (event.Subscription, error) {

	logs, sub, err := _Controller.contract.WatchLogs(opts, "AddedController")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ControllerAddedController)
				if err := _Controller.contract.UnpackLog(event, "AddedController", log); err != nil {
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

// ControllerClaimedIterator is returned from FilterClaimed and is used to iterate over the raw logs and unpacked data for Claimed events raised by the Controller contract.
type ControllerClaimedIterator struct {
	Event *ControllerClaimed // Event containing the contract specifics and raw log

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
func (it *ControllerClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ControllerClaimed)
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
		it.Event = new(ControllerClaimed)
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
func (it *ControllerClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ControllerClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ControllerClaimed represents a Claimed event raised by the Controller contract.
type ControllerClaimed struct {
	To     common.Address
	Asset  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterClaimed is a free log retrieval operation binding the contract event 0xf7a40077ff7a04c7e61f6f26fb13774259ddf1b6bce9ecf26a8276cdd3992683.
//
// Solidity: event Claimed(address _to, address _asset, uint256 _amount)
func (_Controller *ControllerFilterer) FilterClaimed(opts *bind.FilterOpts) (*ControllerClaimedIterator, error) {

	logs, sub, err := _Controller.contract.FilterLogs(opts, "Claimed")
	if err != nil {
		return nil, err
	}
	return &ControllerClaimedIterator{contract: _Controller.contract, event: "Claimed", logs: logs, sub: sub}, nil
}

// WatchClaimed is a free log subscription operation binding the contract event 0xf7a40077ff7a04c7e61f6f26fb13774259ddf1b6bce9ecf26a8276cdd3992683.
//
// Solidity: event Claimed(address _to, address _asset, uint256 _amount)
func (_Controller *ControllerFilterer) WatchClaimed(opts *bind.WatchOpts, sink chan<- *ControllerClaimed) (event.Subscription, error) {

	logs, sub, err := _Controller.contract.WatchLogs(opts, "Claimed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ControllerClaimed)
				if err := _Controller.contract.UnpackLog(event, "Claimed", log); err != nil {
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

// ControllerLockedOwnershipIterator is returned from FilterLockedOwnership and is used to iterate over the raw logs and unpacked data for LockedOwnership events raised by the Controller contract.
type ControllerLockedOwnershipIterator struct {
	Event *ControllerLockedOwnership // Event containing the contract specifics and raw log

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
func (it *ControllerLockedOwnershipIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ControllerLockedOwnership)
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
		it.Event = new(ControllerLockedOwnership)
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
func (it *ControllerLockedOwnershipIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ControllerLockedOwnershipIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ControllerLockedOwnership represents a LockedOwnership event raised by the Controller contract.
type ControllerLockedOwnership struct {
	Locked common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLockedOwnership is a free log retrieval operation binding the contract event 0x808639ff9c8e4732d60b6c2330de498035416d229f27a77d259680895efec122.
//
// Solidity: event LockedOwnership(address _locked)
func (_Controller *ControllerFilterer) FilterLockedOwnership(opts *bind.FilterOpts) (*ControllerLockedOwnershipIterator, error) {

	logs, sub, err := _Controller.contract.FilterLogs(opts, "LockedOwnership")
	if err != nil {
		return nil, err
	}
	return &ControllerLockedOwnershipIterator{contract: _Controller.contract, event: "LockedOwnership", logs: logs, sub: sub}, nil
}

// WatchLockedOwnership is a free log subscription operation binding the contract event 0x808639ff9c8e4732d60b6c2330de498035416d229f27a77d259680895efec122.
//
// Solidity: event LockedOwnership(address _locked)
func (_Controller *ControllerFilterer) WatchLockedOwnership(opts *bind.WatchOpts, sink chan<- *ControllerLockedOwnership) (event.Subscription, error) {

	logs, sub, err := _Controller.contract.WatchLogs(opts, "LockedOwnership")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ControllerLockedOwnership)
				if err := _Controller.contract.UnpackLog(event, "LockedOwnership", log); err != nil {
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

// ControllerRemovedAdminIterator is returned from FilterRemovedAdmin and is used to iterate over the raw logs and unpacked data for RemovedAdmin events raised by the Controller contract.
type ControllerRemovedAdminIterator struct {
	Event *ControllerRemovedAdmin // Event containing the contract specifics and raw log

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
func (it *ControllerRemovedAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ControllerRemovedAdmin)
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
		it.Event = new(ControllerRemovedAdmin)
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
func (it *ControllerRemovedAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ControllerRemovedAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ControllerRemovedAdmin represents a RemovedAdmin event raised by the Controller contract.
type ControllerRemovedAdmin struct {
	Sender common.Address
	Admin  common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRemovedAdmin is a free log retrieval operation binding the contract event 0x787a2e12f4a55b658b8f573c32432ee11a5e8b51677d1e1e937aaf6a0bb5776e.
//
// Solidity: event RemovedAdmin(address _sender, address _admin)
func (_Controller *ControllerFilterer) FilterRemovedAdmin(opts *bind.FilterOpts) (*ControllerRemovedAdminIterator, error) {

	logs, sub, err := _Controller.contract.FilterLogs(opts, "RemovedAdmin")
	if err != nil {
		return nil, err
	}
	return &ControllerRemovedAdminIterator{contract: _Controller.contract, event: "RemovedAdmin", logs: logs, sub: sub}, nil
}

// WatchRemovedAdmin is a free log subscription operation binding the contract event 0x787a2e12f4a55b658b8f573c32432ee11a5e8b51677d1e1e937aaf6a0bb5776e.
//
// Solidity: event RemovedAdmin(address _sender, address _admin)
func (_Controller *ControllerFilterer) WatchRemovedAdmin(opts *bind.WatchOpts, sink chan<- *ControllerRemovedAdmin) (event.Subscription, error) {

	logs, sub, err := _Controller.contract.WatchLogs(opts, "RemovedAdmin")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ControllerRemovedAdmin)
				if err := _Controller.contract.UnpackLog(event, "RemovedAdmin", log); err != nil {
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

// ControllerRemovedControllerIterator is returned from FilterRemovedController and is used to iterate over the raw logs and unpacked data for RemovedController events raised by the Controller contract.
type ControllerRemovedControllerIterator struct {
	Event *ControllerRemovedController // Event containing the contract specifics and raw log

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
func (it *ControllerRemovedControllerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ControllerRemovedController)
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
		it.Event = new(ControllerRemovedController)
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
func (it *ControllerRemovedControllerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ControllerRemovedControllerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ControllerRemovedController represents a RemovedController event raised by the Controller contract.
type ControllerRemovedController struct {
	Sender     common.Address
	Controller common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRemovedController is a free log retrieval operation binding the contract event 0xb6a283aaede08e15ef55c74e3014e30eb0c0040d4b156cccb77391268ea37394.
//
// Solidity: event RemovedController(address _sender, address _controller)
func (_Controller *ControllerFilterer) FilterRemovedController(opts *bind.FilterOpts) (*ControllerRemovedControllerIterator, error) {

	logs, sub, err := _Controller.contract.FilterLogs(opts, "RemovedController")
	if err != nil {
		return nil, err
	}
	return &ControllerRemovedControllerIterator{contract: _Controller.contract, event: "RemovedController", logs: logs, sub: sub}, nil
}

// WatchRemovedController is a free log subscription operation binding the contract event 0xb6a283aaede08e15ef55c74e3014e30eb0c0040d4b156cccb77391268ea37394.
//
// Solidity: event RemovedController(address _sender, address _controller)
func (_Controller *ControllerFilterer) WatchRemovedController(opts *bind.WatchOpts, sink chan<- *ControllerRemovedController) (event.Subscription, error) {

	logs, sub, err := _Controller.contract.WatchLogs(opts, "RemovedController")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ControllerRemovedController)
				if err := _Controller.contract.UnpackLog(event, "RemovedController", log); err != nil {
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

// ControllerTransferredOwnershipIterator is returned from FilterTransferredOwnership and is used to iterate over the raw logs and unpacked data for TransferredOwnership events raised by the Controller contract.
type ControllerTransferredOwnershipIterator struct {
	Event *ControllerTransferredOwnership // Event containing the contract specifics and raw log

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
func (it *ControllerTransferredOwnershipIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ControllerTransferredOwnership)
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
		it.Event = new(ControllerTransferredOwnership)
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
func (it *ControllerTransferredOwnershipIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ControllerTransferredOwnershipIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ControllerTransferredOwnership represents a TransferredOwnership event raised by the Controller contract.
type ControllerTransferredOwnership struct {
	From common.Address
	To   common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterTransferredOwnership is a free log retrieval operation binding the contract event 0x850b3df64837d7d518b45f5aa64d104652c3b80eb5b34a8e3d9eb666cb7cdea5.
//
// Solidity: event TransferredOwnership(address _from, address _to)
func (_Controller *ControllerFilterer) FilterTransferredOwnership(opts *bind.FilterOpts) (*ControllerTransferredOwnershipIterator, error) {

	logs, sub, err := _Controller.contract.FilterLogs(opts, "TransferredOwnership")
	if err != nil {
		return nil, err
	}
	return &ControllerTransferredOwnershipIterator{contract: _Controller.contract, event: "TransferredOwnership", logs: logs, sub: sub}, nil
}

// WatchTransferredOwnership is a free log subscription operation binding the contract event 0x850b3df64837d7d518b45f5aa64d104652c3b80eb5b34a8e3d9eb666cb7cdea5.
//
// Solidity: event TransferredOwnership(address _from, address _to)
func (_Controller *ControllerFilterer) WatchTransferredOwnership(opts *bind.WatchOpts, sink chan<- *ControllerTransferredOwnership) (event.Subscription, error) {

	logs, sub, err := _Controller.contract.WatchLogs(opts, "TransferredOwnership")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ControllerTransferredOwnership)
				if err := _Controller.contract.UnpackLog(event, "TransferredOwnership", log); err != nil {
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
