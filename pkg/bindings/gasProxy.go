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

// GasRefundableGasTokenParameters is an auto generated low-level Go binding around an user-defined struct.
type GasRefundableGasTokenParameters struct {
	FreeCallGasCost  *big.Int
	GasRefundPerUnit *big.Int
}

// GasProxyABI is the input ABI used to generate the binding from.
const GasProxyABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_ens_\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_controllerNode_\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_ensRegistry\",\"type\":\"address\"}],\"name\":\"ENSSetRegistry\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_destination\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_returnData\",\"type\":\"bytes\"}],\"name\":\"ExecutedTransaction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_gasTokenAddress\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"freeCallGasCost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasRefundPerUnit\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structGasRefundable.GasTokenParameters\",\"name\":\"_gasTokenParameters\",\"type\":\"tuple\"}],\"name\":\"SetGasToken\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"controllerNode\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ensRegistry\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_destination\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"executeTransaction\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"gasToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"gasTokenParameters\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"freeCallGasCost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasRefundPerUnit\",\"type\":\"uint256\"}],\"internalType\":\"structGasRefundable.GasTokenParameters\",\"name\":\"\",\"type\":\"tuple\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_gasTokenAddress\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"freeCallGasCost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasRefundPerUnit\",\"type\":\"uint256\"}],\"internalType\":\"structGasRefundable.GasTokenParameters\",\"name\":\"_parameters\",\"type\":\"tuple\"}],\"name\":\"setGasToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// GasProxyBin is the compiled bytecode used for deploying new contracts.
var GasProxyBin = "0x603380546001600160a01b03199081166e0c2e074ec69a0dfb2997ba6c7d2e1e179091557f7f2ce995617d2816b426c5c8698c5ec2952f7a34bb10f38326f74933d5893697603455603580549091166d4946c0e9f43f4dee607b0ef1fa1c17905560c060405261374a608081905261a0aa60a08190526036919091556037553480156200008b57600080fd5b5060405162000fd738038062000fd7833981016040819052620000ae9162000270565b620000c2826001600160e01b03620000de16565b620000d6816001600160e01b03620001a916565b50506200035d565b600054610100900460ff1680620001035750620001036001600160e01b036200024a16565b8062000112575060005460ff16155b6200013a5760405162461bcd60e51b8152600401620001319062000301565b60405180910390fd5b600054610100900460ff1615801562000166576000805460ff1961ff0019909116610100171660011790555b6001600160a01b038216156200019257603380546001600160a01b0319166001600160a01b0384161790555b8015620001a5576000805461ff00191690555b5050565b600054610100900460ff1680620001ce5750620001ce6001600160e01b036200024a16565b80620001dd575060005460ff16155b620001fc5760405162461bcd60e51b8152600401620001319062000301565b600054610100900460ff1615801562000228576000805460ff1961ff0019909116610100171660011790555b8115620001925760348290558015620001a5576000805461ff00191690555050565b303b1590565b80516200025d8162000338565b92915050565b80516200025d8162000352565b600080604083850312156200028457600080fd5b600062000292858562000250565b9250506020620002a58582860162000263565b9150509250929050565b6000620002be602e8362000313565b7f436f6e747261637420696e7374616e63652068617320616c726561647920626581526d195b881a5b9a5d1a585b1a5e995960921b602082015260400192915050565b602080825281016200025d81620002af565b90815260200190565b60006200025d826200032c565b90565b6001600160a01b031690565b62000343816200031c565b81146200034f57600080fd5b50565b620003438162000329565b610c6a806200036d6000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c80633686ba42146100675780633f579f4214610085578063719d5c6f146100a55780637d73b231146100ba578063c91d59fe146100cf578063e2b4ce97146100d7575b600080fd5b61006f6100ec565b60405161007c9190610b56565b60405180910390f35b610098610093366004610741565b61010f565b60405161007c9190610ad5565b6100b86100b3366004610707565b6102cc565b005b6100c261030d565b60405161007c9190610a4a565b6100c261031c565b6100df61032b565b60405161007c9190610ac7565b6100f46105e8565b50604080518082019091526036548152603754602082015290565b606061011a33610331565b61013f5760405162461bcd60e51b815260040161013690610b36565b60405180910390fd5b60005a905060006060876001600160a01b0316878787604051610163929190610a3d565b60006040518083038185875af1925050503d80600081146101a0576040519150601f19603f3d011682016040523d82523d6000602084013e6101a5565b606091505b5091509150816101c75760405162461bcd60e51b815260040161013690610b46565b7ff77753fab406ecfff96d6ff2476c64a838fa9f6d37b1bf190f8546e395e3b61388888888856040516101fe959493929190610a7a565b60405180910390a19250506000601036025a836152080103019050603560009054906101000a90046001600160a01b03166001600160a01b0316636366b93660366001015460366000015484018161025257fe5b046040518263ffffffff1660e01b815260040161026f9190610ac7565b602060405180830381600087803b15801561028957600080fd5b505af115801561029d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506102c191908101906107e5565b505050949350505050565b6102d5336103bf565b6102f15760405162461bcd60e51b815260040161013690610af6565b61030982610304368490038401846107c7565b6103f7565b5050565b6033546001600160a01b031690565b6035546001600160a01b031690565b60345490565b600061033e6034546104c2565b6001600160a01b031663b429afeb836040518263ffffffff1660e01b81526004016103699190610a4a565b60206040518083038186803b15801561038157600080fd5b505afa158015610395573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506103b991908101906107a9565b92915050565b60006103cc6034546104c2565b6001600160a01b03166324d7806c836040518263ffffffff1660e01b81526004016103699190610a4a565b6001600160a01b03821661041d5760405162461bcd60e51b815260040161013690610b16565b805161043b5760405162461bcd60e51b815260040161013690610b26565b602081015161045c5760405162461bcd60e51b815260040161013690610b06565b603580546001600160a01b0319166001600160a01b038416179055805160365560208101516037556040517fd82ec20581833e6b763037b2f1082d6b9c2ba3e8dbaaf4aab4c2b4c49d99fbb5906104b69084908490610a58565b60405180910390a15050565b6033546000906001600160a01b03166104ed5760405162461bcd60e51b815260040161013690610ae6565b603354604051630178b8bf60e01b81526001600160a01b0390911690630178b8bf9061051d908590600401610ac7565b60206040518083038186803b15801561053557600080fd5b505afa158015610549573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525061056d91908101906106e1565b6001600160a01b0316633b3b57de836040518263ffffffff1660e01b81526004016105989190610ac7565b60206040518083038186803b1580156105b057600080fd5b505afa1580156105c4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506103b991908101906106e1565b604051806040016040528060008152602001600081525090565b80356103b981610bfe565b80516103b981610bfe565b80516103b981610c15565b60008083601f84011261063557600080fd5b50813567ffffffffffffffff81111561064d57600080fd5b60208301915083600182028301111561066557600080fd5b9250929050565b60006040828403121561067e57600080fd5b50919050565b60006040828403121561069657600080fd5b6106a06040610b64565b905060006106ae84846106cb565b82525060206106bf848483016106cb565b60208301525092915050565b80356103b981610c1e565b80516103b981610c1e565b6000602082840312156106f357600080fd5b60006106ff848461060d565b949350505050565b6000806060838503121561071a57600080fd5b60006107268585610602565b92505060206107378582860161066c565b9150509250929050565b6000806000806060858703121561075757600080fd5b60006107638787610602565b9450506020610774878288016106cb565b935050604085013567ffffffffffffffff81111561079157600080fd5b61079d87828801610623565b95989497509550505050565b6000602082840312156107bb57600080fd5b60006106ff8484610618565b6000604082840312156107d957600080fd5b60006106ff8484610684565b6000602082840312156107f757600080fd5b60006106ff84846106d6565b61080c81610b9d565b82525050565b61080c81610bad565b60006108278385610b8f565b9350610834838584610bbc565b61083d83610bf4565b9093019392505050565b60006108538385610b98565b9350610860838584610bbc565b50500190565b600061087182610b8b565b61087b8185610b8f565b935061088b818560208601610bc8565b61083d81610bf4565b60006108a1601d83610b8f565b7f454e535265736f6c7661626c65206e6f7420696e697469616c697a6564000000815260200192915050565b60006108da601683610b8f565b7539b2b73232b91034b9903737ba1030b71030b236b4b760511b815260200192915050565b600061090c601883610b8f565b7f67617320726566756e642070657220756e697420697320300000000000000000815260200192915050565b6000610945601883610b8f565b7f67617320746f6b656e2061646472657373206973203078300000000000000000815260200192915050565b600061097e601783610b8f565b7f667265652063616c6c2067617320636f73742069732030000000000000000000815260200192915050565b60006109b7601a83610b8f565b7f73656e646572206973206e6f74206120636f6e74726f6c6c6572000000000000815260200192915050565b60006109f0601483610b8f565b73195e1d195c9b985b0818d85b1b0819985a5b195960621b815260200192915050565b80516040830190610a248482610812565b506020820151610a376020850182610812565b50505050565b60006106ff828486610847565b602081016103b98284610803565b60608101610a668285610803565b610a736020830184610a13565b9392505050565b60808101610a888288610803565b610a956020830187610812565b8181036040830152610aa881858761081b565b90508181036060830152610abc8184610866565b979650505050505050565b602081016103b98284610812565b60208082528101610a738184610866565b602080825281016103b981610894565b602080825281016103b9816108cd565b602080825281016103b9816108ff565b602080825281016103b981610938565b602080825281016103b981610971565b602080825281016103b9816109aa565b602080825281016103b9816109e3565b604081016103b98284610a13565b60405181810167ffffffffffffffff81118282101715610b8357600080fd5b604052919050565b5190565b90815260200190565b919050565b60006103b982610bb0565b151590565b90565b6001600160a01b031690565b82818337506000910152565b60005b83811015610be3578181015183820152602001610bcb565b83811115610a375750506000910152565b601f01601f191690565b610c0781610b9d565b8114610c1257600080fd5b50565b610c0781610ba8565b610c0781610bad56fea365627a7a723158200ee4c2c3b888259a99ad81fc27a75c8899643c1561ff14b47b292d845eb9cf4d6c6578706572696d656e74616cf564736f6c63430005110040"

// DeployGasProxy deploys a new Ethereum contract, binding an instance of GasProxy to it.
func DeployGasProxy(auth *bind.TransactOpts, backend bind.ContractBackend, _ens_ common.Address, _controllerNode_ [32]byte) (common.Address, *types.Transaction, *GasProxy, error) {
	parsed, err := abi.JSON(strings.NewReader(GasProxyABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(GasProxyBin), backend, _ens_, _controllerNode_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &GasProxy{GasProxyCaller: GasProxyCaller{contract: contract}, GasProxyTransactor: GasProxyTransactor{contract: contract}, GasProxyFilterer: GasProxyFilterer{contract: contract}}, nil
}

// GasProxy is an auto generated Go binding around an Ethereum contract.
type GasProxy struct {
	GasProxyCaller     // Read-only binding to the contract
	GasProxyTransactor // Write-only binding to the contract
	GasProxyFilterer   // Log filterer for contract events
}

// GasProxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type GasProxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GasProxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GasProxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GasProxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GasProxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GasProxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GasProxySession struct {
	Contract     *GasProxy         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GasProxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GasProxyCallerSession struct {
	Contract *GasProxyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// GasProxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GasProxyTransactorSession struct {
	Contract     *GasProxyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// GasProxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type GasProxyRaw struct {
	Contract *GasProxy // Generic contract binding to access the raw methods on
}

// GasProxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GasProxyCallerRaw struct {
	Contract *GasProxyCaller // Generic read-only contract binding to access the raw methods on
}

// GasProxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GasProxyTransactorRaw struct {
	Contract *GasProxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGasProxy creates a new instance of GasProxy, bound to a specific deployed contract.
func NewGasProxy(address common.Address, backend bind.ContractBackend) (*GasProxy, error) {
	contract, err := bindGasProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GasProxy{GasProxyCaller: GasProxyCaller{contract: contract}, GasProxyTransactor: GasProxyTransactor{contract: contract}, GasProxyFilterer: GasProxyFilterer{contract: contract}}, nil
}

// NewGasProxyCaller creates a new read-only instance of GasProxy, bound to a specific deployed contract.
func NewGasProxyCaller(address common.Address, caller bind.ContractCaller) (*GasProxyCaller, error) {
	contract, err := bindGasProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GasProxyCaller{contract: contract}, nil
}

// NewGasProxyTransactor creates a new write-only instance of GasProxy, bound to a specific deployed contract.
func NewGasProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*GasProxyTransactor, error) {
	contract, err := bindGasProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GasProxyTransactor{contract: contract}, nil
}

// NewGasProxyFilterer creates a new log filterer instance of GasProxy, bound to a specific deployed contract.
func NewGasProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*GasProxyFilterer, error) {
	contract, err := bindGasProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GasProxyFilterer{contract: contract}, nil
}

// bindGasProxy binds a generic wrapper to an already deployed contract.
func bindGasProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GasProxyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GasProxy *GasProxyRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _GasProxy.Contract.GasProxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GasProxy *GasProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GasProxy.Contract.GasProxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GasProxy *GasProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GasProxy.Contract.GasProxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GasProxy *GasProxyCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _GasProxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GasProxy *GasProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GasProxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GasProxy *GasProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GasProxy.Contract.contract.Transact(opts, method, params...)
}

// ControllerNode is a free data retrieval call binding the contract method 0xe2b4ce97.
//
// Solidity: function controllerNode() constant returns(bytes32)
func (_GasProxy *GasProxyCaller) ControllerNode(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _GasProxy.contract.Call(opts, out, "controllerNode")
	return *ret0, err
}

// ControllerNode is a free data retrieval call binding the contract method 0xe2b4ce97.
//
// Solidity: function controllerNode() constant returns(bytes32)
func (_GasProxy *GasProxySession) ControllerNode() ([32]byte, error) {
	return _GasProxy.Contract.ControllerNode(&_GasProxy.CallOpts)
}

// ControllerNode is a free data retrieval call binding the contract method 0xe2b4ce97.
//
// Solidity: function controllerNode() constant returns(bytes32)
func (_GasProxy *GasProxyCallerSession) ControllerNode() ([32]byte, error) {
	return _GasProxy.Contract.ControllerNode(&_GasProxy.CallOpts)
}

// EnsRegistry is a free data retrieval call binding the contract method 0x7d73b231.
//
// Solidity: function ensRegistry() constant returns(address)
func (_GasProxy *GasProxyCaller) EnsRegistry(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _GasProxy.contract.Call(opts, out, "ensRegistry")
	return *ret0, err
}

// EnsRegistry is a free data retrieval call binding the contract method 0x7d73b231.
//
// Solidity: function ensRegistry() constant returns(address)
func (_GasProxy *GasProxySession) EnsRegistry() (common.Address, error) {
	return _GasProxy.Contract.EnsRegistry(&_GasProxy.CallOpts)
}

// EnsRegistry is a free data retrieval call binding the contract method 0x7d73b231.
//
// Solidity: function ensRegistry() constant returns(address)
func (_GasProxy *GasProxyCallerSession) EnsRegistry() (common.Address, error) {
	return _GasProxy.Contract.EnsRegistry(&_GasProxy.CallOpts)
}

// GasToken is a free data retrieval call binding the contract method 0xc91d59fe.
//
// Solidity: function gasToken() constant returns(address)
func (_GasProxy *GasProxyCaller) GasToken(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _GasProxy.contract.Call(opts, out, "gasToken")
	return *ret0, err
}

// GasToken is a free data retrieval call binding the contract method 0xc91d59fe.
//
// Solidity: function gasToken() constant returns(address)
func (_GasProxy *GasProxySession) GasToken() (common.Address, error) {
	return _GasProxy.Contract.GasToken(&_GasProxy.CallOpts)
}

// GasToken is a free data retrieval call binding the contract method 0xc91d59fe.
//
// Solidity: function gasToken() constant returns(address)
func (_GasProxy *GasProxyCallerSession) GasToken() (common.Address, error) {
	return _GasProxy.Contract.GasToken(&_GasProxy.CallOpts)
}

// GasTokenParameters is a free data retrieval call binding the contract method 0x3686ba42.
//
// Solidity: function gasTokenParameters() constant returns(GasRefundableGasTokenParameters)
func (_GasProxy *GasProxyCaller) GasTokenParameters(opts *bind.CallOpts) (GasRefundableGasTokenParameters, error) {
	var (
		ret0 = new(GasRefundableGasTokenParameters)
	)
	out := ret0
	err := _GasProxy.contract.Call(opts, out, "gasTokenParameters")
	return *ret0, err
}

// GasTokenParameters is a free data retrieval call binding the contract method 0x3686ba42.
//
// Solidity: function gasTokenParameters() constant returns(GasRefundableGasTokenParameters)
func (_GasProxy *GasProxySession) GasTokenParameters() (GasRefundableGasTokenParameters, error) {
	return _GasProxy.Contract.GasTokenParameters(&_GasProxy.CallOpts)
}

// GasTokenParameters is a free data retrieval call binding the contract method 0x3686ba42.
//
// Solidity: function gasTokenParameters() constant returns(GasRefundableGasTokenParameters)
func (_GasProxy *GasProxyCallerSession) GasTokenParameters() (GasRefundableGasTokenParameters, error) {
	return _GasProxy.Contract.GasTokenParameters(&_GasProxy.CallOpts)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0x3f579f42.
//
// Solidity: function executeTransaction(address _destination, uint256 _value, bytes _data) returns(bytes)
func (_GasProxy *GasProxyTransactor) ExecuteTransaction(opts *bind.TransactOpts, _destination common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _GasProxy.contract.Transact(opts, "executeTransaction", _destination, _value, _data)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0x3f579f42.
//
// Solidity: function executeTransaction(address _destination, uint256 _value, bytes _data) returns(bytes)
func (_GasProxy *GasProxySession) ExecuteTransaction(_destination common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _GasProxy.Contract.ExecuteTransaction(&_GasProxy.TransactOpts, _destination, _value, _data)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0x3f579f42.
//
// Solidity: function executeTransaction(address _destination, uint256 _value, bytes _data) returns(bytes)
func (_GasProxy *GasProxyTransactorSession) ExecuteTransaction(_destination common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _GasProxy.Contract.ExecuteTransaction(&_GasProxy.TransactOpts, _destination, _value, _data)
}

// SetGasToken is a paid mutator transaction binding the contract method 0x719d5c6f.
//
// Solidity: function setGasToken(address _gasTokenAddress, GasRefundableGasTokenParameters _parameters) returns()
func (_GasProxy *GasProxyTransactor) SetGasToken(opts *bind.TransactOpts, _gasTokenAddress common.Address, _parameters GasRefundableGasTokenParameters) (*types.Transaction, error) {
	return _GasProxy.contract.Transact(opts, "setGasToken", _gasTokenAddress, _parameters)
}

// SetGasToken is a paid mutator transaction binding the contract method 0x719d5c6f.
//
// Solidity: function setGasToken(address _gasTokenAddress, GasRefundableGasTokenParameters _parameters) returns()
func (_GasProxy *GasProxySession) SetGasToken(_gasTokenAddress common.Address, _parameters GasRefundableGasTokenParameters) (*types.Transaction, error) {
	return _GasProxy.Contract.SetGasToken(&_GasProxy.TransactOpts, _gasTokenAddress, _parameters)
}

// SetGasToken is a paid mutator transaction binding the contract method 0x719d5c6f.
//
// Solidity: function setGasToken(address _gasTokenAddress, GasRefundableGasTokenParameters _parameters) returns()
func (_GasProxy *GasProxyTransactorSession) SetGasToken(_gasTokenAddress common.Address, _parameters GasRefundableGasTokenParameters) (*types.Transaction, error) {
	return _GasProxy.Contract.SetGasToken(&_GasProxy.TransactOpts, _gasTokenAddress, _parameters)
}

// GasProxyENSSetRegistryIterator is returned from FilterENSSetRegistry and is used to iterate over the raw logs and unpacked data for ENSSetRegistry events raised by the GasProxy contract.
type GasProxyENSSetRegistryIterator struct {
	Event *GasProxyENSSetRegistry // Event containing the contract specifics and raw log

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
func (it *GasProxyENSSetRegistryIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GasProxyENSSetRegistry)
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
		it.Event = new(GasProxyENSSetRegistry)
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
func (it *GasProxyENSSetRegistryIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GasProxyENSSetRegistryIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GasProxyENSSetRegistry represents a ENSSetRegistry event raised by the GasProxy contract.
type GasProxyENSSetRegistry struct {
	EnsRegistry common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterENSSetRegistry is a free log retrieval operation binding the contract event 0x186a5507043ec82d7cac3080e4bf22b3f781d96b198adf165a2df507bda90773.
//
// Solidity: event ENSSetRegistry(address _ensRegistry)
func (_GasProxy *GasProxyFilterer) FilterENSSetRegistry(opts *bind.FilterOpts) (*GasProxyENSSetRegistryIterator, error) {

	logs, sub, err := _GasProxy.contract.FilterLogs(opts, "ENSSetRegistry")
	if err != nil {
		return nil, err
	}
	return &GasProxyENSSetRegistryIterator{contract: _GasProxy.contract, event: "ENSSetRegistry", logs: logs, sub: sub}, nil
}

// WatchENSSetRegistry is a free log subscription operation binding the contract event 0x186a5507043ec82d7cac3080e4bf22b3f781d96b198adf165a2df507bda90773.
//
// Solidity: event ENSSetRegistry(address _ensRegistry)
func (_GasProxy *GasProxyFilterer) WatchENSSetRegistry(opts *bind.WatchOpts, sink chan<- *GasProxyENSSetRegistry) (event.Subscription, error) {

	logs, sub, err := _GasProxy.contract.WatchLogs(opts, "ENSSetRegistry")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GasProxyENSSetRegistry)
				if err := _GasProxy.contract.UnpackLog(event, "ENSSetRegistry", log); err != nil {
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

// ParseENSSetRegistry is a log parse operation binding the contract event 0x186a5507043ec82d7cac3080e4bf22b3f781d96b198adf165a2df507bda90773.
//
// Solidity: event ENSSetRegistry(address _ensRegistry)
func (_GasProxy *GasProxyFilterer) ParseENSSetRegistry(log types.Log) (*GasProxyENSSetRegistry, error) {
	event := new(GasProxyENSSetRegistry)
	if err := _GasProxy.contract.UnpackLog(event, "ENSSetRegistry", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GasProxyExecutedTransactionIterator is returned from FilterExecutedTransaction and is used to iterate over the raw logs and unpacked data for ExecutedTransaction events raised by the GasProxy contract.
type GasProxyExecutedTransactionIterator struct {
	Event *GasProxyExecutedTransaction // Event containing the contract specifics and raw log

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
func (it *GasProxyExecutedTransactionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GasProxyExecutedTransaction)
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
		it.Event = new(GasProxyExecutedTransaction)
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
func (it *GasProxyExecutedTransactionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GasProxyExecutedTransactionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GasProxyExecutedTransaction represents a ExecutedTransaction event raised by the GasProxy contract.
type GasProxyExecutedTransaction struct {
	Destination common.Address
	Value       *big.Int
	Data        []byte
	ReturnData  []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterExecutedTransaction is a free log retrieval operation binding the contract event 0xf77753fab406ecfff96d6ff2476c64a838fa9f6d37b1bf190f8546e395e3b613.
//
// Solidity: event ExecutedTransaction(address _destination, uint256 _value, bytes _data, bytes _returnData)
func (_GasProxy *GasProxyFilterer) FilterExecutedTransaction(opts *bind.FilterOpts) (*GasProxyExecutedTransactionIterator, error) {

	logs, sub, err := _GasProxy.contract.FilterLogs(opts, "ExecutedTransaction")
	if err != nil {
		return nil, err
	}
	return &GasProxyExecutedTransactionIterator{contract: _GasProxy.contract, event: "ExecutedTransaction", logs: logs, sub: sub}, nil
}

// WatchExecutedTransaction is a free log subscription operation binding the contract event 0xf77753fab406ecfff96d6ff2476c64a838fa9f6d37b1bf190f8546e395e3b613.
//
// Solidity: event ExecutedTransaction(address _destination, uint256 _value, bytes _data, bytes _returnData)
func (_GasProxy *GasProxyFilterer) WatchExecutedTransaction(opts *bind.WatchOpts, sink chan<- *GasProxyExecutedTransaction) (event.Subscription, error) {

	logs, sub, err := _GasProxy.contract.WatchLogs(opts, "ExecutedTransaction")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GasProxyExecutedTransaction)
				if err := _GasProxy.contract.UnpackLog(event, "ExecutedTransaction", log); err != nil {
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

// ParseExecutedTransaction is a log parse operation binding the contract event 0xf77753fab406ecfff96d6ff2476c64a838fa9f6d37b1bf190f8546e395e3b613.
//
// Solidity: event ExecutedTransaction(address _destination, uint256 _value, bytes _data, bytes _returnData)
func (_GasProxy *GasProxyFilterer) ParseExecutedTransaction(log types.Log) (*GasProxyExecutedTransaction, error) {
	event := new(GasProxyExecutedTransaction)
	if err := _GasProxy.contract.UnpackLog(event, "ExecutedTransaction", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GasProxySetGasTokenIterator is returned from FilterSetGasToken and is used to iterate over the raw logs and unpacked data for SetGasToken events raised by the GasProxy contract.
type GasProxySetGasTokenIterator struct {
	Event *GasProxySetGasToken // Event containing the contract specifics and raw log

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
func (it *GasProxySetGasTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GasProxySetGasToken)
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
		it.Event = new(GasProxySetGasToken)
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
func (it *GasProxySetGasTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GasProxySetGasTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GasProxySetGasToken represents a SetGasToken event raised by the GasProxy contract.
type GasProxySetGasToken struct {
	GasTokenAddress    common.Address
	GasTokenParameters GasRefundableGasTokenParameters
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterSetGasToken is a free log retrieval operation binding the contract event 0xd82ec20581833e6b763037b2f1082d6b9c2ba3e8dbaaf4aab4c2b4c49d99fbb5.
//
// Solidity: event SetGasToken(address _gasTokenAddress, GasRefundableGasTokenParameters _gasTokenParameters)
func (_GasProxy *GasProxyFilterer) FilterSetGasToken(opts *bind.FilterOpts) (*GasProxySetGasTokenIterator, error) {

	logs, sub, err := _GasProxy.contract.FilterLogs(opts, "SetGasToken")
	if err != nil {
		return nil, err
	}
	return &GasProxySetGasTokenIterator{contract: _GasProxy.contract, event: "SetGasToken", logs: logs, sub: sub}, nil
}

// WatchSetGasToken is a free log subscription operation binding the contract event 0xd82ec20581833e6b763037b2f1082d6b9c2ba3e8dbaaf4aab4c2b4c49d99fbb5.
//
// Solidity: event SetGasToken(address _gasTokenAddress, GasRefundableGasTokenParameters _gasTokenParameters)
func (_GasProxy *GasProxyFilterer) WatchSetGasToken(opts *bind.WatchOpts, sink chan<- *GasProxySetGasToken) (event.Subscription, error) {

	logs, sub, err := _GasProxy.contract.WatchLogs(opts, "SetGasToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GasProxySetGasToken)
				if err := _GasProxy.contract.UnpackLog(event, "SetGasToken", log); err != nil {
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

// ParseSetGasToken is a log parse operation binding the contract event 0xd82ec20581833e6b763037b2f1082d6b9c2ba3e8dbaaf4aab4c2b4c49d99fbb5.
//
// Solidity: event SetGasToken(address _gasTokenAddress, GasRefundableGasTokenParameters _gasTokenParameters)
func (_GasProxy *GasProxyFilterer) ParseSetGasToken(log types.Log) (*GasProxySetGasToken, error) {
	event := new(GasProxySetGasToken)
	if err := _GasProxy.contract.UnpackLog(event, "SetGasToken", log); err != nil {
		return nil, err
	}
	return event, nil
}
