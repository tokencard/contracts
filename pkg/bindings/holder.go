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
const HolderABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_burnerContract_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_ens_\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_tokenWhitelistNode_\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_controllerNode_\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"CashAndBurned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"Claimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"Received\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"burner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"controllerNode\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ensRegistry\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"_nonRedeemableAddresses\",\"type\":\"address[]\"}],\"name\":\"nonRedeemableTokenClaim\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenWhitelistNode\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// HolderBin is the compiled bytecode used for deploying new contracts.
var HolderBin = "0x6080604052603380546001600160a01b0319166e0c2e074ec69a0dfb2997ba6c7d2e1e1790557f7f2ce995617d2816b426c5c8698c5ec2952f7a34bb10f38326f74933d58936976034557fe84f90570f13fe09f288f2411ff9cf50da611ed0c7db7f73d48053ffc974d3966035553480156200007a57600080fd5b506040516200145c3803806200145c83398181016040526080811015620000a057600080fd5b5080516020820151604083015160609093015191929091620000c28362000102565b620000cd81620001d9565b620000d8826200028f565b5050603680546001600160a01b0319166001600160a01b039390931692909217909155506200034b565b600054610100900460ff16806200011e57506200011e62000345565b806200012d575060005460ff16155b6200016a5760405162461bcd60e51b815260040180806020018281038252602e8152602001806200142e602e913960400191505060405180910390fd5b600054610100900460ff1615801562000196576000805460ff1961ff0019909116610100171660011790555b6001600160a01b03821615620001c257603380546001600160a01b0319166001600160a01b0384161790555b8015620001d5576000805461ff00191690555b5050565b600054610100900460ff1680620001f55750620001f562000345565b8062000204575060005460ff16155b620002415760405162461bcd60e51b815260040180806020018281038252602e8152602001806200142e602e913960400191505060405180910390fd5b600054610100900460ff161580156200026d576000805460ff1961ff0019909116610100171660011790555b8115620001c25760348290558015620001d5576000805461ff00191690555050565b600054610100900460ff1680620002ab5750620002ab62000345565b80620002ba575060005460ff16155b620002f75760405162461bcd60e51b815260040180806020018281038252602e8152602001806200142e602e913960400191505060405180910390fd5b600054610100900460ff1615801562000323576000805460ff1961ff0019909116610100171660011790555b8115620001c25760358290558015620001d5576000805461ff00191690555050565b303b1590565b6110d3806200035b6000396000f3fe6080604052600436106100595760003560e01c806327810b6e1461009f57806340f6a70f146100d05780637d73b23114610171578063877337b0146101865780639dc29fac146101ad578063e2b4ce97146101e65761009a565b3661009a576040805133815234602082015281517f88a5966d370b9919b20f3e2c13ff65706f196a4e32cc2c12bf57088f88525874929181900390910190a1005b600080fd5b3480156100ab57600080fd5b506100b46101fb565b604080516001600160a01b039092168252519081900360200190f35b3480156100dc57600080fd5b5061015d600480360360408110156100f357600080fd5b6001600160a01b03823516919081019060408101602082013564010000000081111561011e57600080fd5b82018360208201111561013057600080fd5b8035906020019184602083028401116401000000008311171561015257600080fd5b50909250905061020a565b604080519115158252519081900360200190f35b34801561017d57600080fd5b506100b46103bd565b34801561019257600080fd5b5061019b6103cc565b60408051918252519081900360200190f35b3480156101b957600080fd5b5061015d600480360360408110156101d057600080fd5b506001600160a01b0381351690602001356103d2565b3480156101f257600080fd5b5061019b6105ac565b6036546001600160a01b031690565b6000610215336105b2565b61025f576040805162461bcd60e51b815260206004820152601660248201527539b2b73232b91034b9903737ba1030b71030b236b4b760511b604482015290519081900360640190fd5b60005b828110156103b25761028e84848381811061027957fe5b905060200201356001600160a01b031661063f565b156102e0576040805162461bcd60e51b815260206004820152601d60248201527f72656465656d61626c65732063616e6e6f7420626520636c61696d6564000000604482015290519081900360640190fd5b60006103068585848181106102f157fe5b905060200201356001600160a01b0316610658565b905080156103a9576103348686868581811061031e57fe5b905060200201356001600160a01b0316836106e9565b7ff7a40077ff7a04c7e61f6f26fb13774259ddf1b6bce9ecf26a8276cdd39926838686868581811061036257fe5b905060200201356001600160a01b03168360405180846001600160a01b03168152602001836001600160a01b03168152602001828152602001935050505060405180910390a15b50600101610262565b506001949350505050565b6033546001600160a01b031690565b60355490565b6036546000906001600160a01b0316331461041e5760405162461bcd60e51b815260040180806020018281038252602181526020018061107d6021913960400191505060405180910390fd5b8161042b575060016105a6565b60006104b183603660009054906101000a90046001600160a01b03166001600160a01b031663771282f66040518163ffffffff1660e01b815260040160206040518083038186803b15801561047f57600080fd5b505afa158015610493573d6000803e3d6000fd5b505050506040513d60208110156104a957600080fd5b505190610728565b905060606104bd610789565b905060005b815181101561059e5760006104e98383815181106104dc57fe5b6020026020010151610658565b9050801561059557600061050785610501848a61089c565b906108f5565b90506105278885858151811061051957fe5b6020026020010151836106e9565b7f43e074e3351faae8657cc314cf10440a8e7a87ce5092ee4bf9baf56f73fe6c568885858151811061055557fe5b60200260200101518360405180846001600160a01b03168152602001836001600160a01b03168152602001828152602001935050505060405180910390a1505b506001016104c2565b506001925050505b92915050565b60345490565b60006105bf603454610937565b6001600160a01b03166324d7806c836040518263ffffffff1660e01b815260040180826001600160a01b0316815260200191505060206040518083038186803b15801561060b57600080fd5b505afa15801561061f573d6000803e3d6000fd5b505050506040513d602081101561063557600080fd5b505190505b919050565b60008061064b83610a58565b5098975050505050505050565b60006001600160a01b038216156106e257604080516370a0823160e01b815230600482015290516001600160a01b038416916370a08231916024808301926020929190829003018186803b1580156106af57600080fd5b505afa1580156106c3573d6000803e3d6000fd5b505050506040513d60208110156106d957600080fd5b5051905061063a565b504761063a565b6001600160a01b03821661070f5761070a6001600160a01b03841682610be3565b610723565b6107236001600160a01b0383168483610cc8565b505050565b600082820183811015610782576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b9392505050565b6060610796603554610937565b6001600160a01b03166344b049bc6040518163ffffffff1660e01b815260040160006040518083038186803b1580156107ce57600080fd5b505afa1580156107e2573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052602081101561080b57600080fd5b810190808051604051939291908464010000000082111561082b57600080fd5b90830190602082018581111561084057600080fd5b825186602082028301116401000000008211171561085d57600080fd5b82525081516020918201928201910280838360005b8381101561088a578181015183820152602001610872565b50505050905001604052505050905090565b6000826108ab575060006105a6565b828202828482816108b857fe5b04146107825760405162461bcd60e51b81526004018080602001828103825260218152602001806110326021913960400191505060405180910390fd5b600061078283836040518060400160405280601a81526020017f536166654d6174683a206469766973696f6e206279207a65726f000000000000815250610d1a565b6033546000906001600160a01b0316610997576040805162461bcd60e51b815260206004820152601d60248201527f454e535265736f6c7661626c65206e6f7420696e697469616c697a6564000000604482015290519081900360640190fd5b60335460408051630178b8bf60e01b81526004810185905290516001600160a01b0390921691630178b8bf91602480820192602092909190829003018186803b1580156109e357600080fd5b505afa1580156109f7573d6000803e3d6000fd5b505050506040513d6020811015610a0d57600080fd5b505160408051631d9dabef60e11b81526004810185905290516001600160a01b0390921691633b3b57de91602480820192602092909190829003018186803b15801561060b57600080fd5b6060600080600080600080610a6e603554610937565b6001600160a01b0316631f69565f896040518263ffffffff1660e01b815260040180826001600160a01b0316815260200191505060006040518083038186803b158015610aba57600080fd5b505afa158015610ace573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405260e0811015610af757600080fd5b8101908080516040519392919084640100000000821115610b1757600080fd5b908301906020820185811115610b2c57600080fd5b8251640100000000811182820188101715610b4657600080fd5b82525081516020918201929091019080838360005b83811015610b73578181015183820152602001610b5b565b50505050905090810190601f168015610ba05780820380516001836020036101000a031916815260200191505b5060409081526020820151908201516060830151608084015160a085015160c090950151979e50929c50909a509850965094509192505050919395979092949650565b80471015610c38576040805162461bcd60e51b815260206004820152601d60248201527f416464726573733a20696e73756666696369656e742062616c616e6365000000604482015290519081900360640190fd5b6040516000906001600160a01b0384169083908381818185875af1925050503d8060008114610c83576040519150601f19603f3d011682016040523d82523d6000602084013e610c88565b606091505b50509050806107235760405162461bcd60e51b815260040180806020018281038252603a815260200180610ff8603a913960400191505060405180910390fd5b604080516001600160a01b038416602482015260448082018490528251808303909101815260649091019091526020810180516001600160e01b031663a9059cbb60e01b179052610723908490610dbc565b60008183610da65760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b83811015610d6b578181015183820152602001610d53565b50505050905090810190601f168015610d985780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b506000838581610db257fe5b0495945050505050565b6060610e11826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b0316610e6d9092919063ffffffff16565b80519091501561072357808060200190516020811015610e3057600080fd5b50516107235760405162461bcd60e51b815260040180806020018281038252602a815260200180611053602a913960400191505060405180910390fd5b6060610e7c8484600085610e84565b949350505050565b6060610e8f85610ff1565b610ee0576040805162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604482015290519081900360640190fd5b60006060866001600160a01b031685876040518082805190602001908083835b60208310610f1f5780518252601f199092019160209182019101610f00565b6001836020036101000a03801982511681845116808217855250505050505090500191505060006040518083038185875af1925050503d8060008114610f81576040519150601f19603f3d011682016040523d82523d6000602084013e610f86565b606091505b50915091508115610f9a579150610e7c9050565b805115610faa5780518082602001fd5b60405162461bcd60e51b8152602060048201818152865160248401528651879391928392604401919085019080838360008315610d6b578181015183820152602001610d53565b3b15159056fe416464726573733a20756e61626c6520746f2073656e642076616c75652c20726563697069656e74206d61792068617665207265766572746564536166654d6174683a206d756c7469706c69636174696f6e206f766572666c6f775361666545524332303a204552433230206f7065726174696f6e20646964206e6f7420737563636565646275726e657220636f6e7472616374206973206e6f74207468652073656e646572a26469706673582212202ae039c8a58b2fd61d2031ca056f15bad3873ccdd362ec172233ab3d10b0406e64736f6c634300060c0033436f6e747261637420696e7374616e63652068617320616c7265616479206265656e20696e697469616c697a6564"

// DeployHolder deploys a new Ethereum contract, binding an instance of Holder to it.
func DeployHolder(auth *bind.TransactOpts, backend bind.ContractBackend, _burnerContract_ common.Address, _ens_ common.Address, _tokenWhitelistNode_ [32]byte, _controllerNode_ [32]byte) (common.Address, *types.Transaction, *Holder, error) {
	parsed, err := abi.JSON(strings.NewReader(HolderABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(HolderBin), backend, _burnerContract_, _ens_, _tokenWhitelistNode_, _controllerNode_)
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

// ControllerNode is a free data retrieval call binding the contract method 0xe2b4ce97.
//
// Solidity: function controllerNode() constant returns(bytes32)
func (_Holder *HolderCaller) ControllerNode(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Holder.contract.Call(opts, out, "controllerNode")
	return *ret0, err
}

// ControllerNode is a free data retrieval call binding the contract method 0xe2b4ce97.
//
// Solidity: function controllerNode() constant returns(bytes32)
func (_Holder *HolderSession) ControllerNode() ([32]byte, error) {
	return _Holder.Contract.ControllerNode(&_Holder.CallOpts)
}

// ControllerNode is a free data retrieval call binding the contract method 0xe2b4ce97.
//
// Solidity: function controllerNode() constant returns(bytes32)
func (_Holder *HolderCallerSession) ControllerNode() ([32]byte, error) {
	return _Holder.Contract.ControllerNode(&_Holder.CallOpts)
}

// EnsRegistry is a free data retrieval call binding the contract method 0x7d73b231.
//
// Solidity: function ensRegistry() constant returns(address)
func (_Holder *HolderCaller) EnsRegistry(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Holder.contract.Call(opts, out, "ensRegistry")
	return *ret0, err
}

// EnsRegistry is a free data retrieval call binding the contract method 0x7d73b231.
//
// Solidity: function ensRegistry() constant returns(address)
func (_Holder *HolderSession) EnsRegistry() (common.Address, error) {
	return _Holder.Contract.EnsRegistry(&_Holder.CallOpts)
}

// EnsRegistry is a free data retrieval call binding the contract method 0x7d73b231.
//
// Solidity: function ensRegistry() constant returns(address)
func (_Holder *HolderCallerSession) EnsRegistry() (common.Address, error) {
	return _Holder.Contract.EnsRegistry(&_Holder.CallOpts)
}

// TokenWhitelistNode is a free data retrieval call binding the contract method 0x877337b0.
//
// Solidity: function tokenWhitelistNode() constant returns(bytes32)
func (_Holder *HolderCaller) TokenWhitelistNode(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Holder.contract.Call(opts, out, "tokenWhitelistNode")
	return *ret0, err
}

// TokenWhitelistNode is a free data retrieval call binding the contract method 0x877337b0.
//
// Solidity: function tokenWhitelistNode() constant returns(bytes32)
func (_Holder *HolderSession) TokenWhitelistNode() ([32]byte, error) {
	return _Holder.Contract.TokenWhitelistNode(&_Holder.CallOpts)
}

// TokenWhitelistNode is a free data retrieval call binding the contract method 0x877337b0.
//
// Solidity: function tokenWhitelistNode() constant returns(bytes32)
func (_Holder *HolderCallerSession) TokenWhitelistNode() ([32]byte, error) {
	return _Holder.Contract.TokenWhitelistNode(&_Holder.CallOpts)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address _to, uint256 _amount) returns(bool)
func (_Holder *HolderTransactor) Burn(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Holder.contract.Transact(opts, "burn", _to, _amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address _to, uint256 _amount) returns(bool)
func (_Holder *HolderSession) Burn(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Holder.Contract.Burn(&_Holder.TransactOpts, _to, _amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address _to, uint256 _amount) returns(bool)
func (_Holder *HolderTransactorSession) Burn(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Holder.Contract.Burn(&_Holder.TransactOpts, _to, _amount)
}

// NonRedeemableTokenClaim is a paid mutator transaction binding the contract method 0x40f6a70f.
//
// Solidity: function nonRedeemableTokenClaim(address _to, address[] _nonRedeemableAddresses) returns(bool)
func (_Holder *HolderTransactor) NonRedeemableTokenClaim(opts *bind.TransactOpts, _to common.Address, _nonRedeemableAddresses []common.Address) (*types.Transaction, error) {
	return _Holder.contract.Transact(opts, "nonRedeemableTokenClaim", _to, _nonRedeemableAddresses)
}

// NonRedeemableTokenClaim is a paid mutator transaction binding the contract method 0x40f6a70f.
//
// Solidity: function nonRedeemableTokenClaim(address _to, address[] _nonRedeemableAddresses) returns(bool)
func (_Holder *HolderSession) NonRedeemableTokenClaim(_to common.Address, _nonRedeemableAddresses []common.Address) (*types.Transaction, error) {
	return _Holder.Contract.NonRedeemableTokenClaim(&_Holder.TransactOpts, _to, _nonRedeemableAddresses)
}

// NonRedeemableTokenClaim is a paid mutator transaction binding the contract method 0x40f6a70f.
//
// Solidity: function nonRedeemableTokenClaim(address _to, address[] _nonRedeemableAddresses) returns(bool)
func (_Holder *HolderTransactorSession) NonRedeemableTokenClaim(_to common.Address, _nonRedeemableAddresses []common.Address) (*types.Transaction, error) {
	return _Holder.Contract.NonRedeemableTokenClaim(&_Holder.TransactOpts, _to, _nonRedeemableAddresses)
}

// HolderCashAndBurnedIterator is returned from FilterCashAndBurned and is used to iterate over the raw logs and unpacked data for CashAndBurned events raised by the Holder contract.
type HolderCashAndBurnedIterator struct {
	Event *HolderCashAndBurned // Event containing the contract specifics and raw log

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
func (it *HolderCashAndBurnedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HolderCashAndBurned)
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
		it.Event = new(HolderCashAndBurned)
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
func (it *HolderCashAndBurnedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HolderCashAndBurnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HolderCashAndBurned represents a CashAndBurned event raised by the Holder contract.
type HolderCashAndBurned struct {
	To     common.Address
	Asset  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCashAndBurned is a free log retrieval operation binding the contract event 0x43e074e3351faae8657cc314cf10440a8e7a87ce5092ee4bf9baf56f73fe6c56.
//
// Solidity: event CashAndBurned(address _to, address _asset, uint256 _amount)
func (_Holder *HolderFilterer) FilterCashAndBurned(opts *bind.FilterOpts) (*HolderCashAndBurnedIterator, error) {

	logs, sub, err := _Holder.contract.FilterLogs(opts, "CashAndBurned")
	if err != nil {
		return nil, err
	}
	return &HolderCashAndBurnedIterator{contract: _Holder.contract, event: "CashAndBurned", logs: logs, sub: sub}, nil
}

// WatchCashAndBurned is a free log subscription operation binding the contract event 0x43e074e3351faae8657cc314cf10440a8e7a87ce5092ee4bf9baf56f73fe6c56.
//
// Solidity: event CashAndBurned(address _to, address _asset, uint256 _amount)
func (_Holder *HolderFilterer) WatchCashAndBurned(opts *bind.WatchOpts, sink chan<- *HolderCashAndBurned) (event.Subscription, error) {

	logs, sub, err := _Holder.contract.WatchLogs(opts, "CashAndBurned")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HolderCashAndBurned)
				if err := _Holder.contract.UnpackLog(event, "CashAndBurned", log); err != nil {
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

// ParseCashAndBurned is a log parse operation binding the contract event 0x43e074e3351faae8657cc314cf10440a8e7a87ce5092ee4bf9baf56f73fe6c56.
//
// Solidity: event CashAndBurned(address _to, address _asset, uint256 _amount)
func (_Holder *HolderFilterer) ParseCashAndBurned(log types.Log) (*HolderCashAndBurned, error) {
	event := new(HolderCashAndBurned)
	if err := _Holder.contract.UnpackLog(event, "CashAndBurned", log); err != nil {
		return nil, err
	}
	return event, nil
}

// HolderClaimedIterator is returned from FilterClaimed and is used to iterate over the raw logs and unpacked data for Claimed events raised by the Holder contract.
type HolderClaimedIterator struct {
	Event *HolderClaimed // Event containing the contract specifics and raw log

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
func (it *HolderClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HolderClaimed)
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
		it.Event = new(HolderClaimed)
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
func (it *HolderClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HolderClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HolderClaimed represents a Claimed event raised by the Holder contract.
type HolderClaimed struct {
	To     common.Address
	Asset  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterClaimed is a free log retrieval operation binding the contract event 0xf7a40077ff7a04c7e61f6f26fb13774259ddf1b6bce9ecf26a8276cdd3992683.
//
// Solidity: event Claimed(address _to, address _asset, uint256 _amount)
func (_Holder *HolderFilterer) FilterClaimed(opts *bind.FilterOpts) (*HolderClaimedIterator, error) {

	logs, sub, err := _Holder.contract.FilterLogs(opts, "Claimed")
	if err != nil {
		return nil, err
	}
	return &HolderClaimedIterator{contract: _Holder.contract, event: "Claimed", logs: logs, sub: sub}, nil
}

// WatchClaimed is a free log subscription operation binding the contract event 0xf7a40077ff7a04c7e61f6f26fb13774259ddf1b6bce9ecf26a8276cdd3992683.
//
// Solidity: event Claimed(address _to, address _asset, uint256 _amount)
func (_Holder *HolderFilterer) WatchClaimed(opts *bind.WatchOpts, sink chan<- *HolderClaimed) (event.Subscription, error) {

	logs, sub, err := _Holder.contract.WatchLogs(opts, "Claimed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HolderClaimed)
				if err := _Holder.contract.UnpackLog(event, "Claimed", log); err != nil {
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

// ParseClaimed is a log parse operation binding the contract event 0xf7a40077ff7a04c7e61f6f26fb13774259ddf1b6bce9ecf26a8276cdd3992683.
//
// Solidity: event Claimed(address _to, address _asset, uint256 _amount)
func (_Holder *HolderFilterer) ParseClaimed(log types.Log) (*HolderClaimed, error) {
	event := new(HolderClaimed)
	if err := _Holder.contract.UnpackLog(event, "Claimed", log); err != nil {
		return nil, err
	}
	return event, nil
}

// HolderReceivedIterator is returned from FilterReceived and is used to iterate over the raw logs and unpacked data for Received events raised by the Holder contract.
type HolderReceivedIterator struct {
	Event *HolderReceived // Event containing the contract specifics and raw log

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
func (it *HolderReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HolderReceived)
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
		it.Event = new(HolderReceived)
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
func (it *HolderReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HolderReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HolderReceived represents a Received event raised by the Holder contract.
type HolderReceived struct {
	From   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterReceived is a free log retrieval operation binding the contract event 0x88a5966d370b9919b20f3e2c13ff65706f196a4e32cc2c12bf57088f88525874.
//
// Solidity: event Received(address _from, uint256 _amount)
func (_Holder *HolderFilterer) FilterReceived(opts *bind.FilterOpts) (*HolderReceivedIterator, error) {

	logs, sub, err := _Holder.contract.FilterLogs(opts, "Received")
	if err != nil {
		return nil, err
	}
	return &HolderReceivedIterator{contract: _Holder.contract, event: "Received", logs: logs, sub: sub}, nil
}

// WatchReceived is a free log subscription operation binding the contract event 0x88a5966d370b9919b20f3e2c13ff65706f196a4e32cc2c12bf57088f88525874.
//
// Solidity: event Received(address _from, uint256 _amount)
func (_Holder *HolderFilterer) WatchReceived(opts *bind.WatchOpts, sink chan<- *HolderReceived) (event.Subscription, error) {

	logs, sub, err := _Holder.contract.WatchLogs(opts, "Received")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HolderReceived)
				if err := _Holder.contract.UnpackLog(event, "Received", log); err != nil {
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

// ParseReceived is a log parse operation binding the contract event 0x88a5966d370b9919b20f3e2c13ff65706f196a4e32cc2c12bf57088f88525874.
//
// Solidity: event Received(address _from, uint256 _amount)
func (_Holder *HolderFilterer) ParseReceived(log types.Log) (*HolderReceived, error) {
	event := new(HolderReceived)
	if err := _Holder.contract.UnpackLog(event, "Received", log); err != nil {
		return nil, err
	}
	return event, nil
}
