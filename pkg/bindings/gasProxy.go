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
const GasProxyABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_ens_\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_controllerNode_\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_destination\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_returnData\",\"type\":\"bytes\"}],\"name\":\"ExecutedTransaction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_gasTokenAddress\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"freeCallGasCost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasRefundPerUnit\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structGasRefundable.GasTokenParameters\",\"name\":\"_gasTokenParameters\",\"type\":\"tuple\"}],\"name\":\"SetGasToken\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"controllerNode\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ensRegistry\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_destination\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"executeTransaction\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasTokenParameters\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"freeCallGasCost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasRefundPerUnit\",\"type\":\"uint256\"}],\"internalType\":\"structGasRefundable.GasTokenParameters\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_gasTokenAddress\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"freeCallGasCost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasRefundPerUnit\",\"type\":\"uint256\"}],\"internalType\":\"structGasRefundable.GasTokenParameters\",\"name\":\"_parameters\",\"type\":\"tuple\"}],\"name\":\"setGasToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// GasProxyBin is the compiled bytecode used for deploying new contracts.
var GasProxyBin = "0x603380546001600160a01b03199081166e0c2e074ec69a0dfb2997ba6c7d2e1e179091557f7f2ce995617d2816b426c5c8698c5ec2952f7a34bb10f38326f74933d5893697603455603580549091166d4946c0e9f43f4dee607b0ef1fa1c17905560c060405261374a608081905261a0aa60a08190526036919091556037553480156200008b57600080fd5b5060405162000d3838038062000d38833981016040819052620000ae9162000250565b620000c2826001600160e01b03620000de16565b620000d6816001600160e01b03620001a916565b5050620002d8565b600054610100900460ff1680620001035750620001036001600160e01b036200024a16565b8062000112575060005460ff16155b6200013a5760405162461bcd60e51b815260040162000131906200028a565b60405180910390fd5b600054610100900460ff1615801562000166576000805460ff1961ff0019909116610100171660011790555b6001600160a01b038216156200019257603380546001600160a01b0319166001600160a01b0384161790555b8015620001a5576000805461ff00191690555b5050565b600054610100900460ff1680620001ce5750620001ce6001600160e01b036200024a16565b80620001dd575060005460ff16155b620001fc5760405162461bcd60e51b815260040162000131906200028a565b600054610100900460ff1615801562000228576000805460ff1961ff0019909116610100171660011790555b8115620001925760348290558015620001a5576000805461ff00191690555050565b303b1590565b6000806040838503121562000263578182fd5b82516001600160a01b03811681146200027a578283fd5b6020939093015192949293505050565b6020808252602e908201527f436f6e747261637420696e7374616e63652068617320616c726561647920626560408201526d195b881a5b9a5d1a585b1a5e995960921b606082015260800190565b610a5080620002e86000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c80633686ba42146100675780633f579f4214610085578063719d5c6f146100a55780637d73b231146100ba578063c91d59fe146100cf578063e2b4ce97146100d7575b600080fd5b61006f6100ec565b60405161007c91906109f4565b60405180910390f35b610098610093366004610664565b61010f565b60405161007c9190610870565b6100b86100b3366004610625565b6102cc565b005b6100c261030d565b60405161007c91906107d6565b6100c261031c565b6100df61032b565b60405161007c9190610867565b6100f46105e8565b50604080518082019091526036548152603754602082015290565b606061011a33610331565b61013f5760405162461bcd60e51b81526004016101369061098f565b60405180910390fd5b60005a905060006060876001600160a01b03168787876040516101639291906107c6565b60006040518083038185875af1925050503d80600081146101a0576040519150601f19603f3d011682016040523d82523d6000602084013e6101a5565b606091505b5091509150816101c75760405162461bcd60e51b8152600401610136906109c6565b7ff77753fab406ecfff96d6ff2476c64a838fa9f6d37b1bf190f8546e395e3b61388888888856040516101fe959493929190610807565b60405180910390a19250506000601036025a836152080103019050603560009054906101000a90046001600160a01b03166001600160a01b0316636366b93660366001015460366000015484018161025257fe5b046040518263ffffffff1660e01b815260040161026f9190610867565b602060405180830381600087803b15801561028957600080fd5b505af115801561029d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102c19190610754565b505050949350505050565b6102d5336103bf565b6102f15760405162461bcd60e51b8152600401610136906108ba565b610309826103043684900384018461070a565b6103f7565b5050565b6033546001600160a01b031690565b6035546001600160a01b031690565b60345490565b600061033e6034546104c2565b6001600160a01b031663b429afeb836040518263ffffffff1660e01b815260040161036991906107d6565b60206040518083038186803b15801561038157600080fd5b505afa158015610395573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103b991906106ea565b92915050565b60006103cc6034546104c2565b6001600160a01b03166324d7806c836040518263ffffffff1660e01b815260040161036991906107d6565b6001600160a01b03821661041d5760405162461bcd60e51b815260040161013690610921565b805161043b5760405162461bcd60e51b815260040161013690610958565b602081015161045c5760405162461bcd60e51b8152600401610136906108ea565b603580546001600160a01b0319166001600160a01b038416179055805160365560208101516037556040517fd82ec20581833e6b763037b2f1082d6b9c2ba3e8dbaaf4aab4c2b4c49d99fbb5906104b690849084906107ea565b60405180910390a15050565b6033546000906001600160a01b03166104ed5760405162461bcd60e51b815260040161013690610883565b603354604051630178b8bf60e01b81526001600160a01b0390911690630178b8bf9061051d908590600401610867565b60206040518083038186803b15801561053557600080fd5b505afa158015610549573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061056d9190610602565b6001600160a01b0316633b3b57de836040518263ffffffff1660e01b81526004016105989190610867565b60206040518083038186803b1580156105b057600080fd5b505afa1580156105c4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103b99190610602565b604051806040016040528060008152602001600081525090565b600060208284031215610613578081fd5b815161061e81610a02565b9392505050565b6000808284036060811215610638578182fd5b833561064381610a02565b92506040601f1982011215610656578182fd5b506020830190509250929050565b60008060008060608587031215610679578182fd5b843561068481610a02565b935060208501359250604085013567ffffffffffffffff808211156106a7578384fd5b81870188601f8201126106b8578485fd5b80359250818311156106c8578485fd5b8860208483010111156106d9578485fd5b959894975050602090940194505050565b6000602082840312156106fb578081fd5b8151801515811461061e578182fd5b60006040828403121561071b578081fd5b6040516040810181811067ffffffffffffffff8211171561073a578283fd5b604052823581526020928301359281019290925250919050565b600060208284031215610765578081fd5b5051919050565b60008151808452815b8181101561079157602081850181015186830182015201610775565b818111156107a25782602083870101525b50601f01601f19169290920160200192915050565b80518252602090810151910152565b6000828483379101908152919050565b6001600160a01b0391909116815260200190565b6001600160a01b03831681526060810161061e60208301846107b7565b6001600160a01b03861681526020810185905260806040820181905281018390526000838560a08401378060a08584010152601f19601f850116820160a083820301606084015261085b60a082018561076c565b98975050505050505050565b90815260200190565b60006020825261061e602083018461076c565b6020808252601d908201527f454e535265736f6c7661626c65206e6f7420696e697469616c697a6564000000604082015260600190565b60208082526016908201527539b2b73232b91034b9903737ba1030b71030b236b4b760511b604082015260600190565b60208082526018908201527f67617320726566756e642070657220756e697420697320300000000000000000604082015260600190565b60208082526018908201527f67617320746f6b656e2061646472657373206973203078300000000000000000604082015260600190565b60208082526017908201527f667265652063616c6c2067617320636f73742069732030000000000000000000604082015260600190565b6020808252601a908201527f73656e646572206973206e6f74206120636f6e74726f6c6c6572000000000000604082015260600190565b602080825260149082015273195e1d195c9b985b0818d85b1b0819985a5b195960621b604082015260600190565b604081016103b982846107b7565b6001600160a01b0381168114610a1757600080fd5b5056fea2646970667358221220c9420927419dae5d3e58da7d76b93f6092e16a2eb34e54f66acd38cfce298f2f64736f6c634300060b0033"

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
