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

// WalletABI is the input ABI used to generate the binding from.
const WalletABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_ens_\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_controllerNode_\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"_status\",\"type\":\"bool\"}],\"name\":\"ConfirmedOperation\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"confirmOperation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"controllerNode\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ensRegistry\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"sendValue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// WalletBin is the compiled bytecode used for deploying new contracts.
var WalletBin = "0x6080604052603380546001600160a01b0319166e0c2e074ec69a0dfb2997ba6c7d2e1e1790557f7f2ce995617d2816b426c5c8698c5ec2952f7a34bb10f38326f74933d589369760345534801561005557600080fd5b5060405161071b38038061071b8339818101604052604081101561007857600080fd5b508051602090910151610093826001600160e01b036100ac16565b6100a5816001600160e01b0361018316565b505061023f565b600054610100900460ff16806100ce57506100ce6001600160e01b0361023916565b806100dc575060005460ff16155b6101175760405162461bcd60e51b815260040180806020018281038252602e8152602001806106ed602e913960400191505060405180910390fd5b600054610100900460ff16158015610142576000805460ff1961ff0019909116610100171660011790555b6001600160a01b0382161561016d57603380546001600160a01b0319166001600160a01b0384161790555b801561017f576000805461ff00191690555b5050565b600054610100900460ff16806101a557506101a56001600160e01b0361023916565b806101b3575060005460ff16155b6101ee5760405162461bcd60e51b815260040180806020018281038252602e8152602001806106ed602e913960400191505060405180910390fd5b600054610100900460ff16158015610219576000805460ff1961ff0019909116610100171660011790555b811561016d576034829055801561017f576000805461ff00191690555050565b303b1590565b61049f8061024e6000396000f3fe60806040526004361061004e5760003560e01c806324a084df1461005a5780636945341f146100955780637d73b231146100aa578063a9059cbb146100db578063e2b4ce971461011457610055565b3661005557005b600080fd5b34801561006657600080fd5b506100936004803603604081101561007d57600080fd5b506001600160a01b03813516906020013561013b565b005b3480156100a157600080fd5b506100936101d9565b3480156100b657600080fd5b506100bf610269565b604080516001600160a01b039092168252519081900360200190f35b3480156100e757600080fd5b50610093600480360360408110156100fe57600080fd5b506001600160a01b038135169060200135610278565b34801561012057600080fd5b506101296102ae565b60408051918252519081900360200190f35b6040516000906001600160a01b0384169083908381818185875af1925050503d8060008114610186576040519150601f19603f3d011682016040523d82523d6000602084013e61018b565b606091505b50509050806101d4576040805162461bcd60e51b815260206004820152601060248201526f1cd95b9915985b1d594819985a5b195960821b604482015290519081900360640190fd5b505050565b6101e2336102b4565b610233576040805162461bcd60e51b815260206004820152601a60248201527f73656e646572206973206e6f74206120636f6e74726f6c6c6572000000000000604482015290519081900360640190fd5b604080516001815290517f0e8fe5b3baccd13c21d6a20906753a2bbcd86f70d69f4bde83f4e38cf3aae12b9181900360200190a1565b6033546001600160a01b031690565b6040516001600160a01b0383169082156108fc029083906000818181858888f193505050501580156101d4573d6000803e3d6000fd5b60345490565b60006102c1603454610348565b6001600160a01b031663b429afeb836040518263ffffffff1660e01b815260040180826001600160a01b03166001600160a01b0316815260200191505060206040518083038186803b15801561031657600080fd5b505afa15801561032a573d6000803e3d6000fd5b505050506040513d602081101561034057600080fd5b505192915050565b6033546000906001600160a01b03166103a8576040805162461bcd60e51b815260206004820152601d60248201527f454e535265736f6c7661626c65206e6f7420696e697469616c697a6564000000604482015290519081900360640190fd5b60335460408051630178b8bf60e01b81526004810185905290516001600160a01b0390921691630178b8bf91602480820192602092909190829003018186803b1580156103f457600080fd5b505afa158015610408573d6000803e3d6000fd5b505050506040513d602081101561041e57600080fd5b505160408051631d9dabef60e11b81526004810185905290516001600160a01b0390921691633b3b57de91602480820192602092909190829003018186803b15801561031657600080fdfea2646970667358221220e99706be92494d8c24c504c1969f9a8731ea9dadf614e7103447fa47b910309064736f6c634300060b0033436f6e747261637420696e7374616e63652068617320616c7265616479206265656e20696e697469616c697a6564"

// DeployWallet deploys a new Ethereum contract, binding an instance of Wallet to it.
func DeployWallet(auth *bind.TransactOpts, backend bind.ContractBackend, _ens_ common.Address, _controllerNode_ [32]byte) (common.Address, *types.Transaction, *Wallet, error) {
	parsed, err := abi.JSON(strings.NewReader(WalletABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(WalletBin), backend, _ens_, _controllerNode_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Wallet{WalletCaller: WalletCaller{contract: contract}, WalletTransactor: WalletTransactor{contract: contract}, WalletFilterer: WalletFilterer{contract: contract}}, nil
}

// Wallet is an auto generated Go binding around an Ethereum contract.
type Wallet struct {
	WalletCaller     // Read-only binding to the contract
	WalletTransactor // Write-only binding to the contract
	WalletFilterer   // Log filterer for contract events
}

// WalletCaller is an auto generated read-only Go binding around an Ethereum contract.
type WalletCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WalletTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WalletFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WalletSession struct {
	Contract     *Wallet           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WalletCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WalletCallerSession struct {
	Contract *WalletCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// WalletTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WalletTransactorSession struct {
	Contract     *WalletTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WalletRaw is an auto generated low-level Go binding around an Ethereum contract.
type WalletRaw struct {
	Contract *Wallet // Generic contract binding to access the raw methods on
}

// WalletCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WalletCallerRaw struct {
	Contract *WalletCaller // Generic read-only contract binding to access the raw methods on
}

// WalletTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WalletTransactorRaw struct {
	Contract *WalletTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWallet creates a new instance of Wallet, bound to a specific deployed contract.
func NewWallet(address common.Address, backend bind.ContractBackend) (*Wallet, error) {
	contract, err := bindWallet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Wallet{WalletCaller: WalletCaller{contract: contract}, WalletTransactor: WalletTransactor{contract: contract}, WalletFilterer: WalletFilterer{contract: contract}}, nil
}

// NewWalletCaller creates a new read-only instance of Wallet, bound to a specific deployed contract.
func NewWalletCaller(address common.Address, caller bind.ContractCaller) (*WalletCaller, error) {
	contract, err := bindWallet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WalletCaller{contract: contract}, nil
}

// NewWalletTransactor creates a new write-only instance of Wallet, bound to a specific deployed contract.
func NewWalletTransactor(address common.Address, transactor bind.ContractTransactor) (*WalletTransactor, error) {
	contract, err := bindWallet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WalletTransactor{contract: contract}, nil
}

// NewWalletFilterer creates a new log filterer instance of Wallet, bound to a specific deployed contract.
func NewWalletFilterer(address common.Address, filterer bind.ContractFilterer) (*WalletFilterer, error) {
	contract, err := bindWallet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WalletFilterer{contract: contract}, nil
}

// bindWallet binds a generic wrapper to an already deployed contract.
func bindWallet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(WalletABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Wallet *WalletRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Wallet.Contract.WalletCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Wallet *WalletRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Wallet.Contract.WalletTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Wallet *WalletRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Wallet.Contract.WalletTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Wallet *WalletCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Wallet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Wallet *WalletTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Wallet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Wallet *WalletTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Wallet.Contract.contract.Transact(opts, method, params...)
}

// ControllerNode is a free data retrieval call binding the contract method 0xe2b4ce97.
//
// Solidity: function controllerNode() constant returns(bytes32)
func (_Wallet *WalletCaller) ControllerNode(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Wallet.contract.Call(opts, out, "controllerNode")
	return *ret0, err
}

// ControllerNode is a free data retrieval call binding the contract method 0xe2b4ce97.
//
// Solidity: function controllerNode() constant returns(bytes32)
func (_Wallet *WalletSession) ControllerNode() ([32]byte, error) {
	return _Wallet.Contract.ControllerNode(&_Wallet.CallOpts)
}

// ControllerNode is a free data retrieval call binding the contract method 0xe2b4ce97.
//
// Solidity: function controllerNode() constant returns(bytes32)
func (_Wallet *WalletCallerSession) ControllerNode() ([32]byte, error) {
	return _Wallet.Contract.ControllerNode(&_Wallet.CallOpts)
}

// EnsRegistry is a free data retrieval call binding the contract method 0x7d73b231.
//
// Solidity: function ensRegistry() constant returns(address)
func (_Wallet *WalletCaller) EnsRegistry(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Wallet.contract.Call(opts, out, "ensRegistry")
	return *ret0, err
}

// EnsRegistry is a free data retrieval call binding the contract method 0x7d73b231.
//
// Solidity: function ensRegistry() constant returns(address)
func (_Wallet *WalletSession) EnsRegistry() (common.Address, error) {
	return _Wallet.Contract.EnsRegistry(&_Wallet.CallOpts)
}

// EnsRegistry is a free data retrieval call binding the contract method 0x7d73b231.
//
// Solidity: function ensRegistry() constant returns(address)
func (_Wallet *WalletCallerSession) EnsRegistry() (common.Address, error) {
	return _Wallet.Contract.EnsRegistry(&_Wallet.CallOpts)
}

// ConfirmOperation is a paid mutator transaction binding the contract method 0x6945341f.
//
// Solidity: function confirmOperation() returns()
func (_Wallet *WalletTransactor) ConfirmOperation(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "confirmOperation")
}

// ConfirmOperation is a paid mutator transaction binding the contract method 0x6945341f.
//
// Solidity: function confirmOperation() returns()
func (_Wallet *WalletSession) ConfirmOperation() (*types.Transaction, error) {
	return _Wallet.Contract.ConfirmOperation(&_Wallet.TransactOpts)
}

// ConfirmOperation is a paid mutator transaction binding the contract method 0x6945341f.
//
// Solidity: function confirmOperation() returns()
func (_Wallet *WalletTransactorSession) ConfirmOperation() (*types.Transaction, error) {
	return _Wallet.Contract.ConfirmOperation(&_Wallet.TransactOpts)
}

// SendValue is a paid mutator transaction binding the contract method 0x24a084df.
//
// Solidity: function sendValue(address _to, uint256 _amount) returns()
func (_Wallet *WalletTransactor) SendValue(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "sendValue", _to, _amount)
}

// SendValue is a paid mutator transaction binding the contract method 0x24a084df.
//
// Solidity: function sendValue(address _to, uint256 _amount) returns()
func (_Wallet *WalletSession) SendValue(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Wallet.Contract.SendValue(&_Wallet.TransactOpts, _to, _amount)
}

// SendValue is a paid mutator transaction binding the contract method 0x24a084df.
//
// Solidity: function sendValue(address _to, uint256 _amount) returns()
func (_Wallet *WalletTransactorSession) SendValue(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Wallet.Contract.SendValue(&_Wallet.TransactOpts, _to, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _amount) returns()
func (_Wallet *WalletTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "transfer", _to, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _amount) returns()
func (_Wallet *WalletSession) Transfer(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Wallet.Contract.Transfer(&_Wallet.TransactOpts, _to, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _amount) returns()
func (_Wallet *WalletTransactorSession) Transfer(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Wallet.Contract.Transfer(&_Wallet.TransactOpts, _to, _amount)
}

// WalletConfirmedOperationIterator is returned from FilterConfirmedOperation and is used to iterate over the raw logs and unpacked data for ConfirmedOperation events raised by the Wallet contract.
type WalletConfirmedOperationIterator struct {
	Event *WalletConfirmedOperation // Event containing the contract specifics and raw log

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
func (it *WalletConfirmedOperationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletConfirmedOperation)
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
		it.Event = new(WalletConfirmedOperation)
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
func (it *WalletConfirmedOperationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletConfirmedOperationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletConfirmedOperation represents a ConfirmedOperation event raised by the Wallet contract.
type WalletConfirmedOperation struct {
	Status bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterConfirmedOperation is a free log retrieval operation binding the contract event 0x0e8fe5b3baccd13c21d6a20906753a2bbcd86f70d69f4bde83f4e38cf3aae12b.
//
// Solidity: event ConfirmedOperation(bool _status)
func (_Wallet *WalletFilterer) FilterConfirmedOperation(opts *bind.FilterOpts) (*WalletConfirmedOperationIterator, error) {

	logs, sub, err := _Wallet.contract.FilterLogs(opts, "ConfirmedOperation")
	if err != nil {
		return nil, err
	}
	return &WalletConfirmedOperationIterator{contract: _Wallet.contract, event: "ConfirmedOperation", logs: logs, sub: sub}, nil
}

// WatchConfirmedOperation is a free log subscription operation binding the contract event 0x0e8fe5b3baccd13c21d6a20906753a2bbcd86f70d69f4bde83f4e38cf3aae12b.
//
// Solidity: event ConfirmedOperation(bool _status)
func (_Wallet *WalletFilterer) WatchConfirmedOperation(opts *bind.WatchOpts, sink chan<- *WalletConfirmedOperation) (event.Subscription, error) {

	logs, sub, err := _Wallet.contract.WatchLogs(opts, "ConfirmedOperation")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletConfirmedOperation)
				if err := _Wallet.contract.UnpackLog(event, "ConfirmedOperation", log); err != nil {
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

// ParseConfirmedOperation is a log parse operation binding the contract event 0x0e8fe5b3baccd13c21d6a20906753a2bbcd86f70d69f4bde83f4e38cf3aae12b.
//
// Solidity: event ConfirmedOperation(bool _status)
func (_Wallet *WalletFilterer) ParseConfirmedOperation(log types.Log) (*WalletConfirmedOperation, error) {
	event := new(WalletConfirmedOperation)
	if err := _Wallet.contract.UnpackLog(event, "ConfirmedOperation", log); err != nil {
		return nil, err
	}
	return event, nil
}
