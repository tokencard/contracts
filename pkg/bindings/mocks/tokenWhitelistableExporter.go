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
const TokenWhitelistableExporterABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_ens_\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_tokenWhitelistName_\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ensRegistry\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_destination\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"getERC20RecipientAndAmount\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStablecoinInfo\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"}],\"name\":\"getTokenInfo\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"}],\"name\":\"isTokenAvailable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"}],\"name\":\"isTokenLoadable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"}],\"name\":\"isTokenRedeemable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"redeemableTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenAddressArray\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenWhitelistNode\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_updateDate\",\"type\":\"uint256\"}],\"name\":\"updateTokenRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// TokenWhitelistableExporterBin is the compiled bytecode used for deploying new contracts.
var TokenWhitelistableExporterBin = "0x608060405234801561001057600080fd5b50604051610c6b380380610c6b8339818101604052604081101561003357600080fd5b508051602090910151600180546001600160a01b039384166001600160a01b03199182161791829055600080549091169190931617909155600255610bee8061007d6000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c806344b049bc1161007157806344b049bc1461023e5780637d73b23114610246578063877337b01461026a578063afc72e9314610284578063d545782e14610325578063dcc922bc14610359576100a9565b806317e73794146100ae5780631f69565f146100e85780632ff0fcaa146101b85780633efec5e9146101de578063443dd2a4146101e6575b600080fd5b6100d4600480360360208110156100c457600080fd5b50356001600160a01b031661037f565b604080519115158252519081900360200190f35b61010e600480360360208110156100fe57600080fd5b50356001600160a01b0316610390565b6040805160208082018990529181018790528515156060820152841515608082015283151560a082015260c0810183905260e08082528951908201528851909182916101008301918b019080838360005b8381101561017757818101518382015260200161015f565b50505050905090810190601f1680156101a45780820380516001836020036101000a031916815260200191505b509850505050505050505060405180910390f35b6100d4600480360360208110156101ce57600080fd5b50356001600160a01b03166103bb565b61010e6103c6565b6101ee6103ef565b60408051602080825283518183015283519192839290830191858101910280838360005b8381101561022a578181015183820152602001610212565b505050509050019250505060405180910390f35b6101ee6103fe565b61024e610408565b604080516001600160a01b039092168252519081900360200190f35b610272610417565b60408051918252519081900360200190f35b6103026004803603604081101561029a57600080fd5b6001600160a01b038235169190810190604081016020820135600160201b8111156102c457600080fd5b8201836020820111156102d657600080fd5b803590602001918460018302840111600160201b831117156102f757600080fd5b50909250905061041d565b604080516001600160a01b03909316835260208301919091528051918290030190f35b6103576004803603606081101561033b57600080fd5b506001600160a01b03813516906020810135906040013561046c565b005b6100d46004803603602081101561036f57600080fd5b50356001600160a01b031661047c565b600061038a82610487565b92915050565b60606000806000806000806103a4886104a1565b959e949d50929b5090995097509550909350915050565b600061038a82610633565b60606000806000806000806103d961064d565b959d949c50929a50909850965094509092509050565b60606103f96107bc565b905090565b60606103f96108cd565b6001546001600160a01b031690565b60025490565b6000806104608585858080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061091292505050565b91509150935093915050565b610477838383610a1c565b505050565b600061038a82610aab565b600080610493836104a1565b509098975050505050505050565b60606000806000806000806104b7600254610ac4565b6001600160a01b0316631f69565f896040518263ffffffff1660e01b815260040180826001600160a01b03166001600160a01b0316815260200191505060006040518083038186803b15801561050c57600080fd5b505afa158015610520573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405260e081101561054957600080fd5b8101908080516040519392919084600160201b82111561056857600080fd5b90830190602082018581111561057d57600080fd5b8251600160201b81118282018810171561059657600080fd5b82525081516020918201929091019080838360005b838110156105c35781810151838201526020016105ab565b50505050905090810190601f1680156105f05780820380516001836020036101000a031916815260200191505b5060409081526020820151908201516060830151608084015160a085015160c090950151979e50929c50909a509850965094509192505050919395979092949650565b60008061063f836104a1565b509198975050505050505050565b6060600080600080600080610663600254610ac4565b6001600160a01b0316633efec5e96040518163ffffffff1660e01b815260040160006040518083038186803b15801561069b57600080fd5b505afa1580156106af573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405260e08110156106d857600080fd5b8101908080516040519392919084600160201b8211156106f757600080fd5b90830190602082018581111561070c57600080fd5b8251600160201b81118282018810171561072557600080fd5b82525081516020918201929091019080838360005b8381101561075257818101518382015260200161073a565b50505050905090810190601f16801561077f5780820380516001836020036101000a031916815260200191505b5060409081526020820151908201516060830151608084015160a085015160c090950151979f939e50919c509a5098509096509294509192505050565b60606107c9600254610ac4565b6001600160a01b031663443dd2a46040518163ffffffff1660e01b815260040160006040518083038186803b15801561080157600080fd5b505afa158015610815573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052602081101561083e57600080fd5b8101908080516040519392919084600160201b82111561085d57600080fd5b90830190602082018581111561087257600080fd5b82518660208202830111600160201b8211171561088e57600080fd5b82525081516020918201928201910280838360005b838110156108bb5781810151838201526020016108a3565b50505050905001604052505050905090565b60606108da600254610ac4565b6001600160a01b03166344b049bc6040518163ffffffff1660e01b815260040160006040518083038186803b15801561080157600080fd5b600080610920600254610ac4565b6001600160a01b031663afc72e9385856040518363ffffffff1660e01b815260040180836001600160a01b03166001600160a01b0316815260200180602001828103825283818151815260200191508051906020019080838360005b8381101561099457818101518382015260200161097c565b50505050905090810190601f1680156109c15780820380516001836020036101000a031916815260200191505b509350505050604080518083038186803b1580156109de57600080fd5b505afa1580156109f2573d6000803e3d6000fd5b505050506040513d6040811015610a0857600080fd5b508051602090910151909590945092505050565b610a27600254610ac4565b6001600160a01b031663d545782e8484846040518463ffffffff1660e01b815260040180846001600160a01b03166001600160a01b031681526020018381526020018281526020019350505050600060405180830381600087803b158015610a8e57600080fd5b505af1158015610aa2573d6000803e3d6000fd5b50505050505050565b600080610ab7836104a1565b5098975050505050505050565b6000805460408051630178b8bf60e01b81526004810185905290516001600160a01b0390921691630178b8bf91602480820192602092909190829003018186803b158015610b1157600080fd5b505afa158015610b25573d6000803e3d6000fd5b505050506040513d6020811015610b3b57600080fd5b505160408051631d9dabef60e11b81526004810185905290516001600160a01b0390921691633b3b57de91602480820192602092909190829003018186803b158015610b8657600080fd5b505afa158015610b9a573d6000803e3d6000fd5b505050506040513d6020811015610bb057600080fd5b50519291505056fea264697066735822122010737a78c002f82d8bcb149a4a720443f4cb5a25a5354a6a1490cb4406ce910064736f6c63430006040033"

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

// EnsRegistry is a paid mutator transaction binding the contract method 0x7d73b231.
//
// Solidity: function ensRegistry() returns(address)
func (_TokenWhitelistableExporter *TokenWhitelistableExporterTransactor) EnsRegistry(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenWhitelistableExporter.contract.Transact(opts, "ensRegistry")
}

// EnsRegistry is a paid mutator transaction binding the contract method 0x7d73b231.
//
// Solidity: function ensRegistry() returns(address)
func (_TokenWhitelistableExporter *TokenWhitelistableExporterSession) EnsRegistry() (*types.Transaction, error) {
	return _TokenWhitelistableExporter.Contract.EnsRegistry(&_TokenWhitelistableExporter.TransactOpts)
}

// EnsRegistry is a paid mutator transaction binding the contract method 0x7d73b231.
//
// Solidity: function ensRegistry() returns(address)
func (_TokenWhitelistableExporter *TokenWhitelistableExporterTransactorSession) EnsRegistry() (*types.Transaction, error) {
	return _TokenWhitelistableExporter.Contract.EnsRegistry(&_TokenWhitelistableExporter.TransactOpts)
}

// GetERC20RecipientAndAmount is a paid mutator transaction binding the contract method 0xafc72e93.
//
// Solidity: function getERC20RecipientAndAmount(address _destination, bytes _data) returns(address, uint256)
func (_TokenWhitelistableExporter *TokenWhitelistableExporterTransactor) GetERC20RecipientAndAmount(opts *bind.TransactOpts, _destination common.Address, _data []byte) (*types.Transaction, error) {
	return _TokenWhitelistableExporter.contract.Transact(opts, "getERC20RecipientAndAmount", _destination, _data)
}

// GetERC20RecipientAndAmount is a paid mutator transaction binding the contract method 0xafc72e93.
//
// Solidity: function getERC20RecipientAndAmount(address _destination, bytes _data) returns(address, uint256)
func (_TokenWhitelistableExporter *TokenWhitelistableExporterSession) GetERC20RecipientAndAmount(_destination common.Address, _data []byte) (*types.Transaction, error) {
	return _TokenWhitelistableExporter.Contract.GetERC20RecipientAndAmount(&_TokenWhitelistableExporter.TransactOpts, _destination, _data)
}

// GetERC20RecipientAndAmount is a paid mutator transaction binding the contract method 0xafc72e93.
//
// Solidity: function getERC20RecipientAndAmount(address _destination, bytes _data) returns(address, uint256)
func (_TokenWhitelistableExporter *TokenWhitelistableExporterTransactorSession) GetERC20RecipientAndAmount(_destination common.Address, _data []byte) (*types.Transaction, error) {
	return _TokenWhitelistableExporter.Contract.GetERC20RecipientAndAmount(&_TokenWhitelistableExporter.TransactOpts, _destination, _data)
}

// GetStablecoinInfo is a paid mutator transaction binding the contract method 0x3efec5e9.
//
// Solidity: function getStablecoinInfo() returns(string, uint256, uint256, bool, bool, bool, uint256)
func (_TokenWhitelistableExporter *TokenWhitelistableExporterTransactor) GetStablecoinInfo(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenWhitelistableExporter.contract.Transact(opts, "getStablecoinInfo")
}

// GetStablecoinInfo is a paid mutator transaction binding the contract method 0x3efec5e9.
//
// Solidity: function getStablecoinInfo() returns(string, uint256, uint256, bool, bool, bool, uint256)
func (_TokenWhitelistableExporter *TokenWhitelistableExporterSession) GetStablecoinInfo() (*types.Transaction, error) {
	return _TokenWhitelistableExporter.Contract.GetStablecoinInfo(&_TokenWhitelistableExporter.TransactOpts)
}

// GetStablecoinInfo is a paid mutator transaction binding the contract method 0x3efec5e9.
//
// Solidity: function getStablecoinInfo() returns(string, uint256, uint256, bool, bool, bool, uint256)
func (_TokenWhitelistableExporter *TokenWhitelistableExporterTransactorSession) GetStablecoinInfo() (*types.Transaction, error) {
	return _TokenWhitelistableExporter.Contract.GetStablecoinInfo(&_TokenWhitelistableExporter.TransactOpts)
}

// GetTokenInfo is a paid mutator transaction binding the contract method 0x1f69565f.
//
// Solidity: function getTokenInfo(address _a) returns(string, uint256, uint256, bool, bool, bool, uint256)
func (_TokenWhitelistableExporter *TokenWhitelistableExporterTransactor) GetTokenInfo(opts *bind.TransactOpts, _a common.Address) (*types.Transaction, error) {
	return _TokenWhitelistableExporter.contract.Transact(opts, "getTokenInfo", _a)
}

// GetTokenInfo is a paid mutator transaction binding the contract method 0x1f69565f.
//
// Solidity: function getTokenInfo(address _a) returns(string, uint256, uint256, bool, bool, bool, uint256)
func (_TokenWhitelistableExporter *TokenWhitelistableExporterSession) GetTokenInfo(_a common.Address) (*types.Transaction, error) {
	return _TokenWhitelistableExporter.Contract.GetTokenInfo(&_TokenWhitelistableExporter.TransactOpts, _a)
}

// GetTokenInfo is a paid mutator transaction binding the contract method 0x1f69565f.
//
// Solidity: function getTokenInfo(address _a) returns(string, uint256, uint256, bool, bool, bool, uint256)
func (_TokenWhitelistableExporter *TokenWhitelistableExporterTransactorSession) GetTokenInfo(_a common.Address) (*types.Transaction, error) {
	return _TokenWhitelistableExporter.Contract.GetTokenInfo(&_TokenWhitelistableExporter.TransactOpts, _a)
}

// IsTokenAvailable is a paid mutator transaction binding the contract method 0x2ff0fcaa.
//
// Solidity: function isTokenAvailable(address _a) returns(bool)
func (_TokenWhitelistableExporter *TokenWhitelistableExporterTransactor) IsTokenAvailable(opts *bind.TransactOpts, _a common.Address) (*types.Transaction, error) {
	return _TokenWhitelistableExporter.contract.Transact(opts, "isTokenAvailable", _a)
}

// IsTokenAvailable is a paid mutator transaction binding the contract method 0x2ff0fcaa.
//
// Solidity: function isTokenAvailable(address _a) returns(bool)
func (_TokenWhitelistableExporter *TokenWhitelistableExporterSession) IsTokenAvailable(_a common.Address) (*types.Transaction, error) {
	return _TokenWhitelistableExporter.Contract.IsTokenAvailable(&_TokenWhitelistableExporter.TransactOpts, _a)
}

// IsTokenAvailable is a paid mutator transaction binding the contract method 0x2ff0fcaa.
//
// Solidity: function isTokenAvailable(address _a) returns(bool)
func (_TokenWhitelistableExporter *TokenWhitelistableExporterTransactorSession) IsTokenAvailable(_a common.Address) (*types.Transaction, error) {
	return _TokenWhitelistableExporter.Contract.IsTokenAvailable(&_TokenWhitelistableExporter.TransactOpts, _a)
}

// IsTokenLoadable is a paid mutator transaction binding the contract method 0x17e73794.
//
// Solidity: function isTokenLoadable(address _a) returns(bool)
func (_TokenWhitelistableExporter *TokenWhitelistableExporterTransactor) IsTokenLoadable(opts *bind.TransactOpts, _a common.Address) (*types.Transaction, error) {
	return _TokenWhitelistableExporter.contract.Transact(opts, "isTokenLoadable", _a)
}

// IsTokenLoadable is a paid mutator transaction binding the contract method 0x17e73794.
//
// Solidity: function isTokenLoadable(address _a) returns(bool)
func (_TokenWhitelistableExporter *TokenWhitelistableExporterSession) IsTokenLoadable(_a common.Address) (*types.Transaction, error) {
	return _TokenWhitelistableExporter.Contract.IsTokenLoadable(&_TokenWhitelistableExporter.TransactOpts, _a)
}

// IsTokenLoadable is a paid mutator transaction binding the contract method 0x17e73794.
//
// Solidity: function isTokenLoadable(address _a) returns(bool)
func (_TokenWhitelistableExporter *TokenWhitelistableExporterTransactorSession) IsTokenLoadable(_a common.Address) (*types.Transaction, error) {
	return _TokenWhitelistableExporter.Contract.IsTokenLoadable(&_TokenWhitelistableExporter.TransactOpts, _a)
}

// IsTokenRedeemable is a paid mutator transaction binding the contract method 0xdcc922bc.
//
// Solidity: function isTokenRedeemable(address _a) returns(bool)
func (_TokenWhitelistableExporter *TokenWhitelistableExporterTransactor) IsTokenRedeemable(opts *bind.TransactOpts, _a common.Address) (*types.Transaction, error) {
	return _TokenWhitelistableExporter.contract.Transact(opts, "isTokenRedeemable", _a)
}

// IsTokenRedeemable is a paid mutator transaction binding the contract method 0xdcc922bc.
//
// Solidity: function isTokenRedeemable(address _a) returns(bool)
func (_TokenWhitelistableExporter *TokenWhitelistableExporterSession) IsTokenRedeemable(_a common.Address) (*types.Transaction, error) {
	return _TokenWhitelistableExporter.Contract.IsTokenRedeemable(&_TokenWhitelistableExporter.TransactOpts, _a)
}

// IsTokenRedeemable is a paid mutator transaction binding the contract method 0xdcc922bc.
//
// Solidity: function isTokenRedeemable(address _a) returns(bool)
func (_TokenWhitelistableExporter *TokenWhitelistableExporterTransactorSession) IsTokenRedeemable(_a common.Address) (*types.Transaction, error) {
	return _TokenWhitelistableExporter.Contract.IsTokenRedeemable(&_TokenWhitelistableExporter.TransactOpts, _a)
}

// RedeemableTokens is a paid mutator transaction binding the contract method 0x44b049bc.
//
// Solidity: function redeemableTokens() returns(address[])
func (_TokenWhitelistableExporter *TokenWhitelistableExporterTransactor) RedeemableTokens(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenWhitelistableExporter.contract.Transact(opts, "redeemableTokens")
}

// RedeemableTokens is a paid mutator transaction binding the contract method 0x44b049bc.
//
// Solidity: function redeemableTokens() returns(address[])
func (_TokenWhitelistableExporter *TokenWhitelistableExporterSession) RedeemableTokens() (*types.Transaction, error) {
	return _TokenWhitelistableExporter.Contract.RedeemableTokens(&_TokenWhitelistableExporter.TransactOpts)
}

// RedeemableTokens is a paid mutator transaction binding the contract method 0x44b049bc.
//
// Solidity: function redeemableTokens() returns(address[])
func (_TokenWhitelistableExporter *TokenWhitelistableExporterTransactorSession) RedeemableTokens() (*types.Transaction, error) {
	return _TokenWhitelistableExporter.Contract.RedeemableTokens(&_TokenWhitelistableExporter.TransactOpts)
}

// TokenAddressArray is a paid mutator transaction binding the contract method 0x443dd2a4.
//
// Solidity: function tokenAddressArray() returns(address[])
func (_TokenWhitelistableExporter *TokenWhitelistableExporterTransactor) TokenAddressArray(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenWhitelistableExporter.contract.Transact(opts, "tokenAddressArray")
}

// TokenAddressArray is a paid mutator transaction binding the contract method 0x443dd2a4.
//
// Solidity: function tokenAddressArray() returns(address[])
func (_TokenWhitelistableExporter *TokenWhitelistableExporterSession) TokenAddressArray() (*types.Transaction, error) {
	return _TokenWhitelistableExporter.Contract.TokenAddressArray(&_TokenWhitelistableExporter.TransactOpts)
}

// TokenAddressArray is a paid mutator transaction binding the contract method 0x443dd2a4.
//
// Solidity: function tokenAddressArray() returns(address[])
func (_TokenWhitelistableExporter *TokenWhitelistableExporterTransactorSession) TokenAddressArray() (*types.Transaction, error) {
	return _TokenWhitelistableExporter.Contract.TokenAddressArray(&_TokenWhitelistableExporter.TransactOpts)
}

// TokenWhitelistNode is a paid mutator transaction binding the contract method 0x877337b0.
//
// Solidity: function tokenWhitelistNode() returns(bytes32)
func (_TokenWhitelistableExporter *TokenWhitelistableExporterTransactor) TokenWhitelistNode(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenWhitelistableExporter.contract.Transact(opts, "tokenWhitelistNode")
}

// TokenWhitelistNode is a paid mutator transaction binding the contract method 0x877337b0.
//
// Solidity: function tokenWhitelistNode() returns(bytes32)
func (_TokenWhitelistableExporter *TokenWhitelistableExporterSession) TokenWhitelistNode() (*types.Transaction, error) {
	return _TokenWhitelistableExporter.Contract.TokenWhitelistNode(&_TokenWhitelistableExporter.TransactOpts)
}

// TokenWhitelistNode is a paid mutator transaction binding the contract method 0x877337b0.
//
// Solidity: function tokenWhitelistNode() returns(bytes32)
func (_TokenWhitelistableExporter *TokenWhitelistableExporterTransactorSession) TokenWhitelistNode() (*types.Transaction, error) {
	return _TokenWhitelistableExporter.Contract.TokenWhitelistNode(&_TokenWhitelistableExporter.TransactOpts)
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
