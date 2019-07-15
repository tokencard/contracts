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
const HolderABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"isTransferable\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"burner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_nonRedeemableAddresses\",\"type\":\"address[]\"}],\"name\":\"nonRedeemableTokenClaim\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ensRegistry\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tokenWhitelistNode\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"},{\"name\":\"_asset\",\"type\":\"address\"}],\"name\":\"_balance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_account\",\"type\":\"address\"},{\"name\":\"_transferable\",\"type\":\"bool\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_transferable\",\"type\":\"bool\"},{\"name\":\"_burnerContract\",\"type\":\"address\"},{\"name\":\"_ens\",\"type\":\"address\"},{\"name\":\"_tokenWhitelistNameHash\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"Received\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_asset\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"CashAndBurned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_asset\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"Claimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"TransferredOwnership\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_locked\",\"type\":\"address\"}],\"name\":\"LockedOwnership\",\"type\":\"event\"}]"

// HolderBin is the compiled bytecode used for deploying new contracts.
const HolderBin = `608060405234801561001057600080fd5b50604051611321380380611321833981810160405260a081101561003357600080fd5b508051602082015160408301516060840151608090940151600080546001600160a01b0319166001600160a01b0386161760ff60a01b1916740100000000000000000000000000000000000000008515158102919091179182905594959394929392829184918891889160ff9104166100e357604080516001600160a01b038416815290517f808639ff9c8e4732d60b6c2330de498035416d229f27a77d259680895efec1229181900360200190a15b60408051600081526001600160a01b038416602082015281517f850b3df64837d7d518b45f5aa64d104652c3b80eb5b34a8e3d9eb666cb7cdea5929181900390910190a15050600280546001600160a01b039283166001600160a01b0319918216179182905560018054821692841692909217909155600392909255600480549590911694909116939093179092555050505061119c806101856000396000f3fe6080604052600436106100915760003560e01c8063877337b011610059578063877337b0146101e057806389b05656146102075780638da5cb5b146102425780639dc29fac14610257578063b242e5341461029057610091565b80632121dc75146100cd57806327810b6e146100f657806340f6a70f14610127578063715018a6146101b45780637d73b231146101cb575b6040805133815234602082015281517f88a5966d370b9919b20f3e2c13ff65706f196a4e32cc2c12bf57088f88525874929181900390910190a1005b3480156100d957600080fd5b506100e26102cb565b604080519115158252519081900360200190f35b34801561010257600080fd5b5061010b6102db565b604080516001600160a01b039092168252519081900360200190f35b34801561013357600080fd5b506100e26004803603604081101561014a57600080fd5b6001600160a01b03823516919081019060408101602082013564010000000081111561017557600080fd5b82018360208201111561018757600080fd5b803590602001918460208302840111640100000000831117156101a957600080fd5b5090925090506102ea565b3480156101c057600080fd5b506101c961048d565b005b3480156101d757600080fd5b5061010b61058b565b3480156101ec57600080fd5b506101f561059a565b60408051918252519081900360200190f35b34801561021357600080fd5b506101f56004803603604081101561022a57600080fd5b506001600160a01b03813581169160200135166105a0565b34801561024e57600080fd5b5061010b61064d565b34801561026357600080fd5b506100e26004803603604081101561027a57600080fd5b506001600160a01b03813516906020013561065c565b34801561029c57600080fd5b506101c9600480360360408110156102b357600080fd5b506001600160a01b0381351690602001351515610859565b600054600160a01b900460ff1690565b6004546001600160a01b031690565b60006102f533610a13565b61033f576040805162461bcd60e51b815260206004820152601660248201527539b2b73232b91034b9903737ba1030b71037bbb732b960511b604482015290519081900360640190fd5b60005b828110156104825761036e84848381811061035957fe5b905060200201356001600160a01b0316610a27565b156103c0576040805162461bcd60e51b815260206004820152601d60248201527f72656465656d61626c65732063616e6e6f7420626520636c61696d6564000000604482015290519081900360640190fd5b60006103e7308686858181106103d257fe5b905060200201356001600160a01b03166105a0565b9050801561047957610415868686858181106103ff57fe5b905060200201356001600160a01b031683610a40565b7ff7a40077ff7a04c7e61f6f26fb13774259ddf1b6bce9ecf26a8276cdd39926838686868581811061044357fe5b604080516001600160a01b0395861681526020928302949094013594909416908301525080820184905290519081900360600190a15b50600101610342565b506001949350505050565b61049633610a13565b6104e0576040805162461bcd60e51b815260206004820152601660248201527539b2b73232b91034b9903737ba1030b71037bbb732b960511b604482015290519081900360640190fd5b600054600160a01b900460ff1661053e576040805162461bcd60e51b815260206004820152601d60248201527f6f776e657273686970206973206e6f74207472616e7366657261626c65000000604482015290519081900360640190fd5b600080546001600160a01b031916815560408051828152602081019290925280517f850b3df64837d7d518b45f5aa64d104652c3b80eb5b34a8e3d9eb666cb7cdea59281900390910190a1565b6002546001600160a01b031690565b60035490565b60006001600160a01b0382161561063a57816001600160a01b03166370a08231846040518263ffffffff1660e01b815260040180826001600160a01b03166001600160a01b0316815260200191505060206040518083038186803b15801561060757600080fd5b505afa15801561061b573d6000803e3d6000fd5b505050506040513d602081101561063157600080fd5b50519050610647565b506001600160a01b038216315b92915050565b6000546001600160a01b031690565b6004546000906001600160a01b031633146106a85760405162461bcd60e51b81526004018080602001828103825260218152602001806111476021913960400191505060405180910390fd5b816106b557506001610647565b600061074183600460009054906101000a90046001600160a01b03166001600160a01b031663771282f66040518163ffffffff1660e01b815260040160206040518083038186803b15801561070957600080fd5b505afa15801561071d573d6000803e3d6000fd5b505050506040513d602081101561073357600080fd5b50519063ffffffff610aa916565b9050606061074d610b0a565b905060005b815181101561084d57600061077a3084848151811061076d57fe5b60200260200101516105a0565b905080156108445760006107a485610798848a63ffffffff610bdf16565b9063ffffffff610c3816565b90506107c4888585815181106107b657fe5b602002602001015183610a40565b7f43e074e3351faae8657cc314cf10440a8e7a87ce5092ee4bf9baf56f73fe6c56888585815181106107f257fe5b60200260200101518360405180846001600160a01b03166001600160a01b03168152602001836001600160a01b03166001600160a01b03168152602001828152602001935050505060405180910390a1505b50600101610752565b50600195945050505050565b61086233610a13565b6108ac576040805162461bcd60e51b815260206004820152601660248201527539b2b73232b91034b9903737ba1030b71037bbb732b960511b604482015290519081900360640190fd5b600054600160a01b900460ff1661090a576040805162461bcd60e51b815260206004820152601d60248201527f6f776e657273686970206973206e6f74207472616e7366657261626c65000000604482015290519081900360640190fd5b6001600160a01b03821661094f5760405162461bcd60e51b81526004018080602001828103825260238152602001806110fa6023913960400191505060405180910390fd5b6000805460ff60a01b1916600160a01b83151502179055806109a857604080516001600160a01b038416815290517f808639ff9c8e4732d60b6c2330de498035416d229f27a77d259680895efec1229181900360200190a15b600054604080516001600160a01b039283168152918416602083015280517f850b3df64837d7d518b45f5aa64d104652c3b80eb5b34a8e3d9eb666cb7cdea59281900390910190a150600080546001600160a01b0319166001600160a01b0392909216919091179055565b6000546001600160a01b0390811691161490565b600080610a3383610ca2565b5098975050505050505050565b6001600160a01b038216610a8a576040516001600160a01b0384169082156108fc029083906000818181858888f19350505050158015610a84573d6000803e3d6000fd5b50610aa4565b610aa46001600160a01b038316848363ffffffff610dcf16565b505050565b600082820183811015610b03576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b9392505050565b6060610b17600354610e21565b6001600160a01b03166344b049bc6040518163ffffffff1660e01b815260040160006040518083038186803b158015610b4f57600080fd5b505afa158015610b63573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526020811015610b8c57600080fd5b810190808051640100000000811115610ba457600080fd5b82016020810184811115610bb757600080fd5b8151856020820283011164010000000082111715610bd457600080fd5b509094505050505090565b600082610bee57506000610647565b82820282848281610bfb57fe5b0414610b035760405162461bcd60e51b81526004018080602001828103825260218152602001806110d96021913960400191505060405180910390fd5b6000808211610c8e576040805162461bcd60e51b815260206004820152601a60248201527f536166654d6174683a206469766973696f6e206279207a65726f000000000000604482015290519081900360640190fd5b6000828481610c9957fe5b04949350505050565b6060600080600080600080610cb8600354610e21565b6001600160a01b0316631f69565f896040518263ffffffff1660e01b815260040180826001600160a01b03166001600160a01b0316815260200191505060006040518083038186803b158015610d0d57600080fd5b505afa158015610d21573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405260e0811015610d4a57600080fd5b810190808051640100000000811115610d6257600080fd5b82016020810184811115610d7557600080fd5b8151640100000000811182820187101715610d8f57600080fd5b5050602082015160408301516060840151608085015160a086015160c090960151949e50929c50909a509850965090945092505050919395979092949650565b604080516001600160a01b038416602482015260448082018490528251808303909101815260649091019091526020810180516001600160e01b031663a9059cbb60e01b179052610aa4908490610f14565b60015460408051630178b8bf60e01b81526004810184905290516000926001600160a01b031691630178b8bf916024808301926020929190829003018186803b158015610e6d57600080fd5b505afa158015610e81573d6000803e3d6000fd5b505050506040513d6020811015610e9757600080fd5b505160408051631d9dabef60e11b81526004810185905290516001600160a01b0390921691633b3b57de91602480820192602092909190829003018186803b158015610ee257600080fd5b505afa158015610ef6573d6000803e3d6000fd5b505050506040513d6020811015610f0c57600080fd5b505192915050565b610f26826001600160a01b03166110d2565b610f77576040805162461bcd60e51b815260206004820152601f60248201527f5361666545524332303a2063616c6c20746f206e6f6e2d636f6e747261637400604482015290519081900360640190fd5b60006060836001600160a01b0316836040518082805190602001908083835b60208310610fb55780518252601f199092019160209182019101610f96565b6001836020036101000a0380198251168184511680821785525050505050509050019150506000604051808303816000865af19150503d8060008114611017576040519150601f19603f3d011682016040523d82523d6000602084013e61101c565b606091505b509150915081611073576040805162461bcd60e51b815260206004820181905260248201527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564604482015290519081900360640190fd5b8051156110cc5780806020019051602081101561108f57600080fd5b50516110cc5760405162461bcd60e51b815260040180806020018281038252602a81526020018061111d602a913960400191505060405180910390fd5b50505050565b3b15159056fe536166654d6174683a206d756c7469706c69636174696f6e206f766572666c6f776f776e65722063616e6e6f742062652073657420746f207a65726f20616464726573735361666545524332303a204552433230206f7065726174696f6e20646964206e6f7420737563636565646275726e657220636f6e7472616374206973206e6f74207468652073656e646572a265627a7a723058201b8bb09940d5805d37a02109491de28b7944fda7096f62eecf63c8134885938a64736f6c634300050a0032`

// DeployHolder deploys a new Ethereum contract, binding an instance of Holder to it.
func DeployHolder(auth *bind.TransactOpts, backend bind.ContractBackend, _owner common.Address, _transferable bool, _burnerContract common.Address, _ens common.Address, _tokenWhitelistNameHash [32]byte) (common.Address, *types.Transaction, *Holder, error) {
	parsed, err := abi.JSON(strings.NewReader(HolderABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(HolderBin), backend, _owner, _transferable, _burnerContract, _ens, _tokenWhitelistNameHash)
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

// Balance is a free data retrieval call binding the contract method 0x89b05656.
//
// Solidity: function _balance(address _address, address _asset) constant returns(uint256)
func (_Holder *HolderCaller) Balance(opts *bind.CallOpts, _address common.Address, _asset common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Holder.contract.Call(opts, out, "_balance", _address, _asset)
	return *ret0, err
}

// Balance is a free data retrieval call binding the contract method 0x89b05656.
//
// Solidity: function _balance(address _address, address _asset) constant returns(uint256)
func (_Holder *HolderSession) Balance(_address common.Address, _asset common.Address) (*big.Int, error) {
	return _Holder.Contract.Balance(&_Holder.CallOpts, _address, _asset)
}

// Balance is a free data retrieval call binding the contract method 0x89b05656.
//
// Solidity: function _balance(address _address, address _asset) constant returns(uint256)
func (_Holder *HolderCallerSession) Balance(_address common.Address, _asset common.Address) (*big.Int, error) {
	return _Holder.Contract.Balance(&_Holder.CallOpts, _address, _asset)
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

// IsTransferable is a free data retrieval call binding the contract method 0x2121dc75.
//
// Solidity: function isTransferable() constant returns(bool)
func (_Holder *HolderCaller) IsTransferable(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Holder.contract.Call(opts, out, "isTransferable")
	return *ret0, err
}

// IsTransferable is a free data retrieval call binding the contract method 0x2121dc75.
//
// Solidity: function isTransferable() constant returns(bool)
func (_Holder *HolderSession) IsTransferable() (bool, error) {
	return _Holder.Contract.IsTransferable(&_Holder.CallOpts)
}

// IsTransferable is a free data retrieval call binding the contract method 0x2121dc75.
//
// Solidity: function isTransferable() constant returns(bool)
func (_Holder *HolderCallerSession) IsTransferable() (bool, error) {
	return _Holder.Contract.IsTransferable(&_Holder.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Holder *HolderCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Holder.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Holder *HolderSession) Owner() (common.Address, error) {
	return _Holder.Contract.Owner(&_Holder.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Holder *HolderCallerSession) Owner() (common.Address, error) {
	return _Holder.Contract.Owner(&_Holder.CallOpts)
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

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Holder *HolderTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Holder.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Holder *HolderSession) RenounceOwnership() (*types.Transaction, error) {
	return _Holder.Contract.RenounceOwnership(&_Holder.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Holder *HolderTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Holder.Contract.RenounceOwnership(&_Holder.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xb242e534.
//
// Solidity: function transferOwnership(address _account, bool _transferable) returns()
func (_Holder *HolderTransactor) TransferOwnership(opts *bind.TransactOpts, _account common.Address, _transferable bool) (*types.Transaction, error) {
	return _Holder.contract.Transact(opts, "transferOwnership", _account, _transferable)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xb242e534.
//
// Solidity: function transferOwnership(address _account, bool _transferable) returns()
func (_Holder *HolderSession) TransferOwnership(_account common.Address, _transferable bool) (*types.Transaction, error) {
	return _Holder.Contract.TransferOwnership(&_Holder.TransactOpts, _account, _transferable)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xb242e534.
//
// Solidity: function transferOwnership(address _account, bool _transferable) returns()
func (_Holder *HolderTransactorSession) TransferOwnership(_account common.Address, _transferable bool) (*types.Transaction, error) {
	return _Holder.Contract.TransferOwnership(&_Holder.TransactOpts, _account, _transferable)
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

// HolderLockedOwnershipIterator is returned from FilterLockedOwnership and is used to iterate over the raw logs and unpacked data for LockedOwnership events raised by the Holder contract.
type HolderLockedOwnershipIterator struct {
	Event *HolderLockedOwnership // Event containing the contract specifics and raw log

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
func (it *HolderLockedOwnershipIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HolderLockedOwnership)
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
		it.Event = new(HolderLockedOwnership)
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
func (it *HolderLockedOwnershipIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HolderLockedOwnershipIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HolderLockedOwnership represents a LockedOwnership event raised by the Holder contract.
type HolderLockedOwnership struct {
	Locked common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLockedOwnership is a free log retrieval operation binding the contract event 0x808639ff9c8e4732d60b6c2330de498035416d229f27a77d259680895efec122.
//
// Solidity: event LockedOwnership(address _locked)
func (_Holder *HolderFilterer) FilterLockedOwnership(opts *bind.FilterOpts) (*HolderLockedOwnershipIterator, error) {

	logs, sub, err := _Holder.contract.FilterLogs(opts, "LockedOwnership")
	if err != nil {
		return nil, err
	}
	return &HolderLockedOwnershipIterator{contract: _Holder.contract, event: "LockedOwnership", logs: logs, sub: sub}, nil
}

// WatchLockedOwnership is a free log subscription operation binding the contract event 0x808639ff9c8e4732d60b6c2330de498035416d229f27a77d259680895efec122.
//
// Solidity: event LockedOwnership(address _locked)
func (_Holder *HolderFilterer) WatchLockedOwnership(opts *bind.WatchOpts, sink chan<- *HolderLockedOwnership) (event.Subscription, error) {

	logs, sub, err := _Holder.contract.WatchLogs(opts, "LockedOwnership")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HolderLockedOwnership)
				if err := _Holder.contract.UnpackLog(event, "LockedOwnership", log); err != nil {
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

// HolderTransferredOwnershipIterator is returned from FilterTransferredOwnership and is used to iterate over the raw logs and unpacked data for TransferredOwnership events raised by the Holder contract.
type HolderTransferredOwnershipIterator struct {
	Event *HolderTransferredOwnership // Event containing the contract specifics and raw log

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
func (it *HolderTransferredOwnershipIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HolderTransferredOwnership)
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
		it.Event = new(HolderTransferredOwnership)
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
func (it *HolderTransferredOwnershipIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HolderTransferredOwnershipIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HolderTransferredOwnership represents a TransferredOwnership event raised by the Holder contract.
type HolderTransferredOwnership struct {
	From common.Address
	To   common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterTransferredOwnership is a free log retrieval operation binding the contract event 0x850b3df64837d7d518b45f5aa64d104652c3b80eb5b34a8e3d9eb666cb7cdea5.
//
// Solidity: event TransferredOwnership(address _from, address _to)
func (_Holder *HolderFilterer) FilterTransferredOwnership(opts *bind.FilterOpts) (*HolderTransferredOwnershipIterator, error) {

	logs, sub, err := _Holder.contract.FilterLogs(opts, "TransferredOwnership")
	if err != nil {
		return nil, err
	}
	return &HolderTransferredOwnershipIterator{contract: _Holder.contract, event: "TransferredOwnership", logs: logs, sub: sub}, nil
}

// WatchTransferredOwnership is a free log subscription operation binding the contract event 0x850b3df64837d7d518b45f5aa64d104652c3b80eb5b34a8e3d9eb666cb7cdea5.
//
// Solidity: event TransferredOwnership(address _from, address _to)
func (_Holder *HolderFilterer) WatchTransferredOwnership(opts *bind.WatchOpts, sink chan<- *HolderTransferredOwnership) (event.Subscription, error) {

	logs, sub, err := _Holder.contract.WatchLogs(opts, "TransferredOwnership")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HolderTransferredOwnership)
				if err := _Holder.contract.UnpackLog(event, "TransferredOwnership", log); err != nil {
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
