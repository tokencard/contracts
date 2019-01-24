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

// LicenceABI is the input ABI used to generate the binding from.
const LicenceABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"isTransferable\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tokenHolder\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_fee\",\"type\":\"uint256\"},{\"name\":\"_asset\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"load\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newFee\",\"type\":\"uint256\"}],\"name\":\"updateFee\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"DAO\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cryptoFloat\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_account\",\"type\":\"address\"},{\"name\":\"_transferable\",\"type\":\"bool\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newDAO\",\"type\":\"address\"}],\"name\":\"updateDAO\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"fee\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_transferable\",\"type\":\"bool\"},{\"name\":\"_dao\",\"type\":\"address\"},{\"name\":\"_cryptoFloat\",\"type\":\"address\"},{\"name\":\"_tokenHolder\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"Received\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_dao\",\"type\":\"address\"}],\"name\":\"UpdatedDAO\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_fee\",\"type\":\"address\"}],\"name\":\"UpdatedFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_contract\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_asset\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"TransferredToTokenHolder\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_contract\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_asset\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"TransferredToCryptoFloat\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_dao\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_newFee\",\"type\":\"uint256\"}],\"name\":\"ChangedLicenceFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"TransferredOwnership\",\"type\":\"event\"}]"

// LicenceBin is the compiled bytecode used for deploying new contracts.
const LicenceBin = `608060405234801561001057600080fd5b5060405160a080610ec3833981016040818152825160208085015183860151606087015160809097015160008054600160a060020a031916600160a060020a03871690811760a060020a60ff02191674010000000000000000000000000000000000000000861515021782559088529387019390935284519396919590949193879287927f850b3df64837d7d518b45f5aa64d104652c3b80eb5b34a8e3d9eb666cb7cdea5929081900390910190a1505060048054600160a060020a03948516600160a060020a0319918216179091556001805493851693821693909317909255600280549190931691161790555050610db48061010f6000396000f3006080604052600436106100ae5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416632121dc7581146100f5578063420a83e71461011e578063715018a61461014f578063755677b9146101665780638da5cb5b146101805780639012c4a81461019557806398fabd3a146101ad578063a036ba60146101c2578063b242e534146101d7578063bcb6c0b5146101fd578063ddca3f431461021e575b36156100b957600080fd5b6040805133815234602082015281517f88a5966d370b9919b20f3e2c13ff65706f196a4e32cc2c12bf57088f88525874929181900390910190a1005b34801561010157600080fd5b5061010a610245565b604080519115158252519081900360200190f35b34801561012a57600080fd5b50610133610266565b60408051600160a060020a039092168252519081900360200190f35b34801561015b57600080fd5b50610164610275565b005b61010a600435600160a060020a03602435166044356103b0565b34801561018c57600080fd5b506101336108a3565b3480156101a157600080fd5b506101646004356108b2565b3480156101b957600080fd5b50610133610980565b3480156101ce57600080fd5b5061013361098f565b3480156101e357600080fd5b50610164600160a060020a0360043516602435151561099e565b34801561020957600080fd5b50610164600160a060020a0360043516610ba6565b34801561022a57600080fd5b50610233610d58565b60408051918252519081900360200190f35b60005474010000000000000000000000000000000000000000900460ff1690565b600254600160a060020a031690565b61027d610d5e565b15156102d3576040805160e560020a62461bcd02815260206004820152601660248201527f73656e646572206973206e6f7420616e206f776e657200000000000000000000604482015290519081900360640190fd5b60005474010000000000000000000000000000000000000000900460ff161515610347576040805160e560020a62461bcd02815260206004820152601d60248201527f6f776e657273686970206973206e6f74207472616e7366657261626c65000000604482015290519081900360640190fd5b6000805460408051600160a060020a039092168252602082019290925281517f850b3df64837d7d518b45f5aa64d104652c3b80eb5b34a8e3d9eb666cb7cdea5929181900390910190a16000805473ffffffffffffffffffffffffffffffffffffffff19169055565b6000811515806103bf57508315155b151561043b576040805160e560020a62461bcd02815260206004820152602260248201527f66656520616e6420616d6f756e742063616e6e6f7420626f7468206265207a6560448201527f726f000000000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b6003548460640281151561044b57fe5b0482146104a2576040805160e560020a62461bcd02815260206004820152601260248201527f746865736520646f6e2774206164642075700000000000000000000000000000604482015290519081900360640190fd5b600160a060020a038316156106f057600254604080517f23b872dd000000000000000000000000000000000000000000000000000000008152336004820152600160a060020a039283166024820152604481018790529051918516916323b872dd916064808201926020929091908290030181600087803b15801561052657600080fd5b505af115801561053a573d6000803e3d6000fd5b505050506040513d602081101561055057600080fd5b505115156105ce576040805160e560020a62461bcd02815260206004820152603960248201527f455243323020666565207472616e736665722066726f6d2065787465726e616c60448201527f206163636f756e742077617320756e7375636365737366756c00000000000000606482015290519081900360840190fd5b600154604080517f23b872dd000000000000000000000000000000000000000000000000000000008152336004820152600160a060020a039283166024820152604481018590529051918516916323b872dd916064808201926020929091908290030181600087803b15801561064357600080fd5b505af1158015610657573d6000803e3d6000fd5b505050506040513d602081101561066d57600080fd5b505115156106eb576040805160e560020a62461bcd02815260206004820152603b60248201527f455243323020746f6b656e207472616e736665722066726f6d2065787465726e60448201527f616c206163636f756e742077617320756e7375636365737366756c0000000000606482015290519081900360840190fd5b6107f3565b610700828563ffffffff610d6f16565b341461077c576040805160e560020a62461bcd02815260206004820152602760248201527f65746865722073656e74206974206e6f7420657175616c20746f20616d6f756e60448201527f74202b2066656500000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b600254604051600160a060020a039091169085156108fc029086906000818181858888f193505050501580156107b6573d6000803e3d6000fd5b50600154604051600160a060020a039091169083156108fc029084906000818181858888f193505050501580156107f1573d6000803e3d6000fd5b505b60025460408051338152600160a060020a0392831660208201529185168282015260608201869052517fdd9dfad7b30d6b224e235f89565871419d3dec3b563a4e231f12d2cc97f9acfc9181900360800190a160015460408051338152600160a060020a0392831660208201529185168282015260608201849052517fc8a7b0bd71097b47b2cad75e4e939d2aeb7fae88110e68f93b83fed08e9d3c389181900360800190a15060019392505050565b600054600160a060020a031690565b600454600160a060020a031633146108c957600080fd5b806001111580156108db575060648111155b1515610931576040805160e560020a62461bcd02815260206004820152601860248201527f70657263656e7420666565206f7574206f662072616e67650000000000000000604482015290519081900360640190fd5b600381905560045460408051600160a060020a0390921682526020820183905280517fed263c6ee1194d6a8ba3bee297ab0af0dd4b3b24f89ea68fea5c25e1642a975d9281900390910190a150565b600454600160a060020a031681565b600154600160a060020a031690565b6109a6610d5e565b15156109fc576040805160e560020a62461bcd02815260206004820152601660248201527f73656e646572206973206e6f7420616e206f776e657200000000000000000000604482015290519081900360640190fd5b60005474010000000000000000000000000000000000000000900460ff161515610a70576040805160e560020a62461bcd02815260206004820152601d60248201527f6f776e657273686970206973206e6f74207472616e7366657261626c65000000604482015290519081900360640190fd5b600160a060020a0382161515610af6576040805160e560020a62461bcd02815260206004820152602360248201527f6f776e65722063616e6e6f742062652073657420746f207a65726f206164647260448201527f6573730000000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b6000805474ff000000000000000000000000000000000000000019167401000000000000000000000000000000000000000083151502179081905560408051600160a060020a039283168152918416602083015280517f850b3df64837d7d518b45f5aa64d104652c3b80eb5b34a8e3d9eb666cb7cdea59281900390910190a1506000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b610bae610d5e565b1515610c04576040805160e560020a62461bcd02815260206004820152601660248201527f73656e646572206973206e6f7420616e206f776e657200000000000000000000604482015290519081900360640190fd5b60048054604080517fa4e2d6340000000000000000000000000000000000000000000000000000000081529051600160a060020a039092169263a4e2d6349282820192602092908290030181600087803b158015610c6157600080fd5b505af1158015610c75573d6000803e3d6000fd5b505050506040513d6020811015610c8b57600080fd5b505115610ce2576040805160e560020a62461bcd02815260206004820152600d60248201527f44414f206973206c6f636b656400000000000000000000000000000000000000604482015290519081900360640190fd5b6004805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0383161790557f97ad44d5f57e52b66c8d238e30b7cccdc1e11ca21e52f3838a6d1d4fccc41ec3610d336108a3565b60408051600160a060020a03928316815291841660208301528051918290030190a150565b60035481565b600054600160a060020a0316331490565b600082820183811015610d8157600080fd5b93925050505600a165627a7a72305820dc6927f91999e15443f9c83197b25c09afc0f443680805fc38f98db0aa21492c0029`

// DeployLicence deploys a new Ethereum contract, binding an instance of Licence to it.
func DeployLicence(auth *bind.TransactOpts, backend bind.ContractBackend, _owner common.Address, _transferable bool, _dao common.Address, _cryptoFloat common.Address, _tokenHolder common.Address) (common.Address, *types.Transaction, *Licence, error) {
	parsed, err := abi.JSON(strings.NewReader(LicenceABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(LicenceBin), backend, _owner, _transferable, _dao, _cryptoFloat, _tokenHolder)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Licence{LicenceCaller: LicenceCaller{contract: contract}, LicenceTransactor: LicenceTransactor{contract: contract}, LicenceFilterer: LicenceFilterer{contract: contract}}, nil
}

// Licence is an auto generated Go binding around an Ethereum contract.
type Licence struct {
	LicenceCaller     // Read-only binding to the contract
	LicenceTransactor // Write-only binding to the contract
	LicenceFilterer   // Log filterer for contract events
}

// LicenceCaller is an auto generated read-only Go binding around an Ethereum contract.
type LicenceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LicenceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LicenceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LicenceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LicenceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LicenceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LicenceSession struct {
	Contract     *Licence          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LicenceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LicenceCallerSession struct {
	Contract *LicenceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// LicenceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LicenceTransactorSession struct {
	Contract     *LicenceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// LicenceRaw is an auto generated low-level Go binding around an Ethereum contract.
type LicenceRaw struct {
	Contract *Licence // Generic contract binding to access the raw methods on
}

// LicenceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LicenceCallerRaw struct {
	Contract *LicenceCaller // Generic read-only contract binding to access the raw methods on
}

// LicenceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LicenceTransactorRaw struct {
	Contract *LicenceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLicence creates a new instance of Licence, bound to a specific deployed contract.
func NewLicence(address common.Address, backend bind.ContractBackend) (*Licence, error) {
	contract, err := bindLicence(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Licence{LicenceCaller: LicenceCaller{contract: contract}, LicenceTransactor: LicenceTransactor{contract: contract}, LicenceFilterer: LicenceFilterer{contract: contract}}, nil
}

// NewLicenceCaller creates a new read-only instance of Licence, bound to a specific deployed contract.
func NewLicenceCaller(address common.Address, caller bind.ContractCaller) (*LicenceCaller, error) {
	contract, err := bindLicence(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LicenceCaller{contract: contract}, nil
}

// NewLicenceTransactor creates a new write-only instance of Licence, bound to a specific deployed contract.
func NewLicenceTransactor(address common.Address, transactor bind.ContractTransactor) (*LicenceTransactor, error) {
	contract, err := bindLicence(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LicenceTransactor{contract: contract}, nil
}

// NewLicenceFilterer creates a new log filterer instance of Licence, bound to a specific deployed contract.
func NewLicenceFilterer(address common.Address, filterer bind.ContractFilterer) (*LicenceFilterer, error) {
	contract, err := bindLicence(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LicenceFilterer{contract: contract}, nil
}

// bindLicence binds a generic wrapper to an already deployed contract.
func bindLicence(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LicenceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Licence *LicenceRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Licence.Contract.LicenceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Licence *LicenceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Licence.Contract.LicenceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Licence *LicenceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Licence.Contract.LicenceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Licence *LicenceCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Licence.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Licence *LicenceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Licence.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Licence *LicenceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Licence.Contract.contract.Transact(opts, method, params...)
}

// DAO is a free data retrieval call binding the contract method 0x98fabd3a.
//
// Solidity: function DAO() constant returns(address)
func (_Licence *LicenceCaller) DAO(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Licence.contract.Call(opts, out, "DAO")
	return *ret0, err
}

// DAO is a free data retrieval call binding the contract method 0x98fabd3a.
//
// Solidity: function DAO() constant returns(address)
func (_Licence *LicenceSession) DAO() (common.Address, error) {
	return _Licence.Contract.DAO(&_Licence.CallOpts)
}

// DAO is a free data retrieval call binding the contract method 0x98fabd3a.
//
// Solidity: function DAO() constant returns(address)
func (_Licence *LicenceCallerSession) DAO() (common.Address, error) {
	return _Licence.Contract.DAO(&_Licence.CallOpts)
}

// CryptoFloat is a free data retrieval call binding the contract method 0xa036ba60.
//
// Solidity: function cryptoFloat() constant returns(address)
func (_Licence *LicenceCaller) CryptoFloat(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Licence.contract.Call(opts, out, "cryptoFloat")
	return *ret0, err
}

// CryptoFloat is a free data retrieval call binding the contract method 0xa036ba60.
//
// Solidity: function cryptoFloat() constant returns(address)
func (_Licence *LicenceSession) CryptoFloat() (common.Address, error) {
	return _Licence.Contract.CryptoFloat(&_Licence.CallOpts)
}

// CryptoFloat is a free data retrieval call binding the contract method 0xa036ba60.
//
// Solidity: function cryptoFloat() constant returns(address)
func (_Licence *LicenceCallerSession) CryptoFloat() (common.Address, error) {
	return _Licence.Contract.CryptoFloat(&_Licence.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() constant returns(uint256)
func (_Licence *LicenceCaller) Fee(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Licence.contract.Call(opts, out, "fee")
	return *ret0, err
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() constant returns(uint256)
func (_Licence *LicenceSession) Fee() (*big.Int, error) {
	return _Licence.Contract.Fee(&_Licence.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() constant returns(uint256)
func (_Licence *LicenceCallerSession) Fee() (*big.Int, error) {
	return _Licence.Contract.Fee(&_Licence.CallOpts)
}

// IsTransferable is a free data retrieval call binding the contract method 0x2121dc75.
//
// Solidity: function isTransferable() constant returns(bool)
func (_Licence *LicenceCaller) IsTransferable(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Licence.contract.Call(opts, out, "isTransferable")
	return *ret0, err
}

// IsTransferable is a free data retrieval call binding the contract method 0x2121dc75.
//
// Solidity: function isTransferable() constant returns(bool)
func (_Licence *LicenceSession) IsTransferable() (bool, error) {
	return _Licence.Contract.IsTransferable(&_Licence.CallOpts)
}

// IsTransferable is a free data retrieval call binding the contract method 0x2121dc75.
//
// Solidity: function isTransferable() constant returns(bool)
func (_Licence *LicenceCallerSession) IsTransferable() (bool, error) {
	return _Licence.Contract.IsTransferable(&_Licence.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Licence *LicenceCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Licence.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Licence *LicenceSession) Owner() (common.Address, error) {
	return _Licence.Contract.Owner(&_Licence.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Licence *LicenceCallerSession) Owner() (common.Address, error) {
	return _Licence.Contract.Owner(&_Licence.CallOpts)
}

// TokenHolder is a free data retrieval call binding the contract method 0x420a83e7.
//
// Solidity: function tokenHolder() constant returns(address)
func (_Licence *LicenceCaller) TokenHolder(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Licence.contract.Call(opts, out, "tokenHolder")
	return *ret0, err
}

// TokenHolder is a free data retrieval call binding the contract method 0x420a83e7.
//
// Solidity: function tokenHolder() constant returns(address)
func (_Licence *LicenceSession) TokenHolder() (common.Address, error) {
	return _Licence.Contract.TokenHolder(&_Licence.CallOpts)
}

// TokenHolder is a free data retrieval call binding the contract method 0x420a83e7.
//
// Solidity: function tokenHolder() constant returns(address)
func (_Licence *LicenceCallerSession) TokenHolder() (common.Address, error) {
	return _Licence.Contract.TokenHolder(&_Licence.CallOpts)
}

// Load is a paid mutator transaction binding the contract method 0x755677b9.
//
// Solidity: function load(_fee uint256, _asset address, _amount uint256) returns(bool)
func (_Licence *LicenceTransactor) Load(opts *bind.TransactOpts, _fee *big.Int, _asset common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Licence.contract.Transact(opts, "load", _fee, _asset, _amount)
}

// Load is a paid mutator transaction binding the contract method 0x755677b9.
//
// Solidity: function load(_fee uint256, _asset address, _amount uint256) returns(bool)
func (_Licence *LicenceSession) Load(_fee *big.Int, _asset common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Licence.Contract.Load(&_Licence.TransactOpts, _fee, _asset, _amount)
}

// Load is a paid mutator transaction binding the contract method 0x755677b9.
//
// Solidity: function load(_fee uint256, _asset address, _amount uint256) returns(bool)
func (_Licence *LicenceTransactorSession) Load(_fee *big.Int, _asset common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Licence.Contract.Load(&_Licence.TransactOpts, _fee, _asset, _amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Licence *LicenceTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Licence.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Licence *LicenceSession) RenounceOwnership() (*types.Transaction, error) {
	return _Licence.Contract.RenounceOwnership(&_Licence.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Licence *LicenceTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Licence.Contract.RenounceOwnership(&_Licence.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xb242e534.
//
// Solidity: function transferOwnership(_account address, _transferable bool) returns()
func (_Licence *LicenceTransactor) TransferOwnership(opts *bind.TransactOpts, _account common.Address, _transferable bool) (*types.Transaction, error) {
	return _Licence.contract.Transact(opts, "transferOwnership", _account, _transferable)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xb242e534.
//
// Solidity: function transferOwnership(_account address, _transferable bool) returns()
func (_Licence *LicenceSession) TransferOwnership(_account common.Address, _transferable bool) (*types.Transaction, error) {
	return _Licence.Contract.TransferOwnership(&_Licence.TransactOpts, _account, _transferable)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xb242e534.
//
// Solidity: function transferOwnership(_account address, _transferable bool) returns()
func (_Licence *LicenceTransactorSession) TransferOwnership(_account common.Address, _transferable bool) (*types.Transaction, error) {
	return _Licence.Contract.TransferOwnership(&_Licence.TransactOpts, _account, _transferable)
}

// UpdateDAO is a paid mutator transaction binding the contract method 0xbcb6c0b5.
//
// Solidity: function updateDAO(_newDAO address) returns()
func (_Licence *LicenceTransactor) UpdateDAO(opts *bind.TransactOpts, _newDAO common.Address) (*types.Transaction, error) {
	return _Licence.contract.Transact(opts, "updateDAO", _newDAO)
}

// UpdateDAO is a paid mutator transaction binding the contract method 0xbcb6c0b5.
//
// Solidity: function updateDAO(_newDAO address) returns()
func (_Licence *LicenceSession) UpdateDAO(_newDAO common.Address) (*types.Transaction, error) {
	return _Licence.Contract.UpdateDAO(&_Licence.TransactOpts, _newDAO)
}

// UpdateDAO is a paid mutator transaction binding the contract method 0xbcb6c0b5.
//
// Solidity: function updateDAO(_newDAO address) returns()
func (_Licence *LicenceTransactorSession) UpdateDAO(_newDAO common.Address) (*types.Transaction, error) {
	return _Licence.Contract.UpdateDAO(&_Licence.TransactOpts, _newDAO)
}

// UpdateFee is a paid mutator transaction binding the contract method 0x9012c4a8.
//
// Solidity: function updateFee(_newFee uint256) returns()
func (_Licence *LicenceTransactor) UpdateFee(opts *bind.TransactOpts, _newFee *big.Int) (*types.Transaction, error) {
	return _Licence.contract.Transact(opts, "updateFee", _newFee)
}

// UpdateFee is a paid mutator transaction binding the contract method 0x9012c4a8.
//
// Solidity: function updateFee(_newFee uint256) returns()
func (_Licence *LicenceSession) UpdateFee(_newFee *big.Int) (*types.Transaction, error) {
	return _Licence.Contract.UpdateFee(&_Licence.TransactOpts, _newFee)
}

// UpdateFee is a paid mutator transaction binding the contract method 0x9012c4a8.
//
// Solidity: function updateFee(_newFee uint256) returns()
func (_Licence *LicenceTransactorSession) UpdateFee(_newFee *big.Int) (*types.Transaction, error) {
	return _Licence.Contract.UpdateFee(&_Licence.TransactOpts, _newFee)
}

// LicenceChangedLicenceFeeIterator is returned from FilterChangedLicenceFee and is used to iterate over the raw logs and unpacked data for ChangedLicenceFee events raised by the Licence contract.
type LicenceChangedLicenceFeeIterator struct {
	Event *LicenceChangedLicenceFee // Event containing the contract specifics and raw log

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
func (it *LicenceChangedLicenceFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LicenceChangedLicenceFee)
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
		it.Event = new(LicenceChangedLicenceFee)
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
func (it *LicenceChangedLicenceFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LicenceChangedLicenceFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LicenceChangedLicenceFee represents a ChangedLicenceFee event raised by the Licence contract.
type LicenceChangedLicenceFee struct {
	Dao    common.Address
	NewFee *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterChangedLicenceFee is a free log retrieval operation binding the contract event 0xed263c6ee1194d6a8ba3bee297ab0af0dd4b3b24f89ea68fea5c25e1642a975d.
//
// Solidity: e ChangedLicenceFee(_dao address, _newFee uint256)
func (_Licence *LicenceFilterer) FilterChangedLicenceFee(opts *bind.FilterOpts) (*LicenceChangedLicenceFeeIterator, error) {

	logs, sub, err := _Licence.contract.FilterLogs(opts, "ChangedLicenceFee")
	if err != nil {
		return nil, err
	}
	return &LicenceChangedLicenceFeeIterator{contract: _Licence.contract, event: "ChangedLicenceFee", logs: logs, sub: sub}, nil
}

// WatchChangedLicenceFee is a free log subscription operation binding the contract event 0xed263c6ee1194d6a8ba3bee297ab0af0dd4b3b24f89ea68fea5c25e1642a975d.
//
// Solidity: e ChangedLicenceFee(_dao address, _newFee uint256)
func (_Licence *LicenceFilterer) WatchChangedLicenceFee(opts *bind.WatchOpts, sink chan<- *LicenceChangedLicenceFee) (event.Subscription, error) {

	logs, sub, err := _Licence.contract.WatchLogs(opts, "ChangedLicenceFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LicenceChangedLicenceFee)
				if err := _Licence.contract.UnpackLog(event, "ChangedLicenceFee", log); err != nil {
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

// LicenceReceivedIterator is returned from FilterReceived and is used to iterate over the raw logs and unpacked data for Received events raised by the Licence contract.
type LicenceReceivedIterator struct {
	Event *LicenceReceived // Event containing the contract specifics and raw log

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
func (it *LicenceReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LicenceReceived)
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
		it.Event = new(LicenceReceived)
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
func (it *LicenceReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LicenceReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LicenceReceived represents a Received event raised by the Licence contract.
type LicenceReceived struct {
	From   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterReceived is a free log retrieval operation binding the contract event 0x88a5966d370b9919b20f3e2c13ff65706f196a4e32cc2c12bf57088f88525874.
//
// Solidity: e Received(_from address, _amount uint256)
func (_Licence *LicenceFilterer) FilterReceived(opts *bind.FilterOpts) (*LicenceReceivedIterator, error) {

	logs, sub, err := _Licence.contract.FilterLogs(opts, "Received")
	if err != nil {
		return nil, err
	}
	return &LicenceReceivedIterator{contract: _Licence.contract, event: "Received", logs: logs, sub: sub}, nil
}

// WatchReceived is a free log subscription operation binding the contract event 0x88a5966d370b9919b20f3e2c13ff65706f196a4e32cc2c12bf57088f88525874.
//
// Solidity: e Received(_from address, _amount uint256)
func (_Licence *LicenceFilterer) WatchReceived(opts *bind.WatchOpts, sink chan<- *LicenceReceived) (event.Subscription, error) {

	logs, sub, err := _Licence.contract.WatchLogs(opts, "Received")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LicenceReceived)
				if err := _Licence.contract.UnpackLog(event, "Received", log); err != nil {
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

// LicenceTransferredOwnershipIterator is returned from FilterTransferredOwnership and is used to iterate over the raw logs and unpacked data for TransferredOwnership events raised by the Licence contract.
type LicenceTransferredOwnershipIterator struct {
	Event *LicenceTransferredOwnership // Event containing the contract specifics and raw log

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
func (it *LicenceTransferredOwnershipIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LicenceTransferredOwnership)
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
		it.Event = new(LicenceTransferredOwnership)
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
func (it *LicenceTransferredOwnershipIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LicenceTransferredOwnershipIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LicenceTransferredOwnership represents a TransferredOwnership event raised by the Licence contract.
type LicenceTransferredOwnership struct {
	From common.Address
	To   common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterTransferredOwnership is a free log retrieval operation binding the contract event 0x850b3df64837d7d518b45f5aa64d104652c3b80eb5b34a8e3d9eb666cb7cdea5.
//
// Solidity: e TransferredOwnership(_from address, _to address)
func (_Licence *LicenceFilterer) FilterTransferredOwnership(opts *bind.FilterOpts) (*LicenceTransferredOwnershipIterator, error) {

	logs, sub, err := _Licence.contract.FilterLogs(opts, "TransferredOwnership")
	if err != nil {
		return nil, err
	}
	return &LicenceTransferredOwnershipIterator{contract: _Licence.contract, event: "TransferredOwnership", logs: logs, sub: sub}, nil
}

// WatchTransferredOwnership is a free log subscription operation binding the contract event 0x850b3df64837d7d518b45f5aa64d104652c3b80eb5b34a8e3d9eb666cb7cdea5.
//
// Solidity: e TransferredOwnership(_from address, _to address)
func (_Licence *LicenceFilterer) WatchTransferredOwnership(opts *bind.WatchOpts, sink chan<- *LicenceTransferredOwnership) (event.Subscription, error) {

	logs, sub, err := _Licence.contract.WatchLogs(opts, "TransferredOwnership")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LicenceTransferredOwnership)
				if err := _Licence.contract.UnpackLog(event, "TransferredOwnership", log); err != nil {
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

// LicenceTransferredToCryptoFloatIterator is returned from FilterTransferredToCryptoFloat and is used to iterate over the raw logs and unpacked data for TransferredToCryptoFloat events raised by the Licence contract.
type LicenceTransferredToCryptoFloatIterator struct {
	Event *LicenceTransferredToCryptoFloat // Event containing the contract specifics and raw log

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
func (it *LicenceTransferredToCryptoFloatIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LicenceTransferredToCryptoFloat)
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
		it.Event = new(LicenceTransferredToCryptoFloat)
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
func (it *LicenceTransferredToCryptoFloatIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LicenceTransferredToCryptoFloatIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LicenceTransferredToCryptoFloat represents a TransferredToCryptoFloat event raised by the Licence contract.
type LicenceTransferredToCryptoFloat struct {
	From     common.Address
	Contract common.Address
	Asset    common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferredToCryptoFloat is a free log retrieval operation binding the contract event 0xc8a7b0bd71097b47b2cad75e4e939d2aeb7fae88110e68f93b83fed08e9d3c38.
//
// Solidity: e TransferredToCryptoFloat(_from address, _contract address, _asset address, _amount uint256)
func (_Licence *LicenceFilterer) FilterTransferredToCryptoFloat(opts *bind.FilterOpts) (*LicenceTransferredToCryptoFloatIterator, error) {

	logs, sub, err := _Licence.contract.FilterLogs(opts, "TransferredToCryptoFloat")
	if err != nil {
		return nil, err
	}
	return &LicenceTransferredToCryptoFloatIterator{contract: _Licence.contract, event: "TransferredToCryptoFloat", logs: logs, sub: sub}, nil
}

// WatchTransferredToCryptoFloat is a free log subscription operation binding the contract event 0xc8a7b0bd71097b47b2cad75e4e939d2aeb7fae88110e68f93b83fed08e9d3c38.
//
// Solidity: e TransferredToCryptoFloat(_from address, _contract address, _asset address, _amount uint256)
func (_Licence *LicenceFilterer) WatchTransferredToCryptoFloat(opts *bind.WatchOpts, sink chan<- *LicenceTransferredToCryptoFloat) (event.Subscription, error) {

	logs, sub, err := _Licence.contract.WatchLogs(opts, "TransferredToCryptoFloat")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LicenceTransferredToCryptoFloat)
				if err := _Licence.contract.UnpackLog(event, "TransferredToCryptoFloat", log); err != nil {
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

// LicenceTransferredToTokenHolderIterator is returned from FilterTransferredToTokenHolder and is used to iterate over the raw logs and unpacked data for TransferredToTokenHolder events raised by the Licence contract.
type LicenceTransferredToTokenHolderIterator struct {
	Event *LicenceTransferredToTokenHolder // Event containing the contract specifics and raw log

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
func (it *LicenceTransferredToTokenHolderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LicenceTransferredToTokenHolder)
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
		it.Event = new(LicenceTransferredToTokenHolder)
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
func (it *LicenceTransferredToTokenHolderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LicenceTransferredToTokenHolderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LicenceTransferredToTokenHolder represents a TransferredToTokenHolder event raised by the Licence contract.
type LicenceTransferredToTokenHolder struct {
	From     common.Address
	Contract common.Address
	Asset    common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferredToTokenHolder is a free log retrieval operation binding the contract event 0xdd9dfad7b30d6b224e235f89565871419d3dec3b563a4e231f12d2cc97f9acfc.
//
// Solidity: e TransferredToTokenHolder(_from address, _contract address, _asset address, _amount uint256)
func (_Licence *LicenceFilterer) FilterTransferredToTokenHolder(opts *bind.FilterOpts) (*LicenceTransferredToTokenHolderIterator, error) {

	logs, sub, err := _Licence.contract.FilterLogs(opts, "TransferredToTokenHolder")
	if err != nil {
		return nil, err
	}
	return &LicenceTransferredToTokenHolderIterator{contract: _Licence.contract, event: "TransferredToTokenHolder", logs: logs, sub: sub}, nil
}

// WatchTransferredToTokenHolder is a free log subscription operation binding the contract event 0xdd9dfad7b30d6b224e235f89565871419d3dec3b563a4e231f12d2cc97f9acfc.
//
// Solidity: e TransferredToTokenHolder(_from address, _contract address, _asset address, _amount uint256)
func (_Licence *LicenceFilterer) WatchTransferredToTokenHolder(opts *bind.WatchOpts, sink chan<- *LicenceTransferredToTokenHolder) (event.Subscription, error) {

	logs, sub, err := _Licence.contract.WatchLogs(opts, "TransferredToTokenHolder")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LicenceTransferredToTokenHolder)
				if err := _Licence.contract.UnpackLog(event, "TransferredToTokenHolder", log); err != nil {
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

// LicenceUpdatedDAOIterator is returned from FilterUpdatedDAO and is used to iterate over the raw logs and unpacked data for UpdatedDAO events raised by the Licence contract.
type LicenceUpdatedDAOIterator struct {
	Event *LicenceUpdatedDAO // Event containing the contract specifics and raw log

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
func (it *LicenceUpdatedDAOIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LicenceUpdatedDAO)
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
		it.Event = new(LicenceUpdatedDAO)
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
func (it *LicenceUpdatedDAOIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LicenceUpdatedDAOIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LicenceUpdatedDAO represents a UpdatedDAO event raised by the Licence contract.
type LicenceUpdatedDAO struct {
	Owner common.Address
	Dao   common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterUpdatedDAO is a free log retrieval operation binding the contract event 0x97ad44d5f57e52b66c8d238e30b7cccdc1e11ca21e52f3838a6d1d4fccc41ec3.
//
// Solidity: e UpdatedDAO(_owner address, _dao address)
func (_Licence *LicenceFilterer) FilterUpdatedDAO(opts *bind.FilterOpts) (*LicenceUpdatedDAOIterator, error) {

	logs, sub, err := _Licence.contract.FilterLogs(opts, "UpdatedDAO")
	if err != nil {
		return nil, err
	}
	return &LicenceUpdatedDAOIterator{contract: _Licence.contract, event: "UpdatedDAO", logs: logs, sub: sub}, nil
}

// WatchUpdatedDAO is a free log subscription operation binding the contract event 0x97ad44d5f57e52b66c8d238e30b7cccdc1e11ca21e52f3838a6d1d4fccc41ec3.
//
// Solidity: e UpdatedDAO(_owner address, _dao address)
func (_Licence *LicenceFilterer) WatchUpdatedDAO(opts *bind.WatchOpts, sink chan<- *LicenceUpdatedDAO) (event.Subscription, error) {

	logs, sub, err := _Licence.contract.WatchLogs(opts, "UpdatedDAO")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LicenceUpdatedDAO)
				if err := _Licence.contract.UnpackLog(event, "UpdatedDAO", log); err != nil {
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

// LicenceUpdatedFeeIterator is returned from FilterUpdatedFee and is used to iterate over the raw logs and unpacked data for UpdatedFee events raised by the Licence contract.
type LicenceUpdatedFeeIterator struct {
	Event *LicenceUpdatedFee // Event containing the contract specifics and raw log

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
func (it *LicenceUpdatedFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LicenceUpdatedFee)
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
		it.Event = new(LicenceUpdatedFee)
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
func (it *LicenceUpdatedFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LicenceUpdatedFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LicenceUpdatedFee represents a UpdatedFee event raised by the Licence contract.
type LicenceUpdatedFee struct {
	Owner common.Address
	Fee   common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterUpdatedFee is a free log retrieval operation binding the contract event 0x05467857a672883f86395f63a4360793abece3cc5d6df6b03292f1087d8a5c35.
//
// Solidity: e UpdatedFee(_owner address, _fee address)
func (_Licence *LicenceFilterer) FilterUpdatedFee(opts *bind.FilterOpts) (*LicenceUpdatedFeeIterator, error) {

	logs, sub, err := _Licence.contract.FilterLogs(opts, "UpdatedFee")
	if err != nil {
		return nil, err
	}
	return &LicenceUpdatedFeeIterator{contract: _Licence.contract, event: "UpdatedFee", logs: logs, sub: sub}, nil
}

// WatchUpdatedFee is a free log subscription operation binding the contract event 0x05467857a672883f86395f63a4360793abece3cc5d6df6b03292f1087d8a5c35.
//
// Solidity: e UpdatedFee(_owner address, _fee address)
func (_Licence *LicenceFilterer) WatchUpdatedFee(opts *bind.WatchOpts, sink chan<- *LicenceUpdatedFee) (event.Subscription, error) {

	logs, sub, err := _Licence.contract.WatchLogs(opts, "UpdatedFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LicenceUpdatedFee)
				if err := _Licence.contract.UnpackLog(event, "UpdatedFee", log); err != nil {
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
