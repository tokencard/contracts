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

// WalletABI is the input ABI used to generate the binding from.
const WalletABI = "[{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_owner_\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_transferable_\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"_ens_\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_tokenWhitelistNode_\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_controllerNode_\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_licenceNode_\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_spendLimit_\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"_addresses\",\"type\":\"address[]\"}],\"name\":\"AddedToWhitelist\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"_assets\",\"type\":\"address[]\"}],\"name\":\"BulkTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"CancelledWhitelistAddition\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"CancelledWhitelistRemoval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_returndata\",\"type\":\"bytes\"}],\"name\":\"ExecutedRelayedTransaction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_destination\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_returndata\",\"type\":\"bytes\"}],\"name\":\"ExecutedTransaction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_currentNonce\",\"type\":\"uint256\"}],\"name\":\"IncreasedRelayNonce\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"LoadedTokenCard\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_locked\",\"type\":\"address\"}],\"name\":\"LockedOwnership\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"Received\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"_addresses\",\"type\":\"address[]\"}],\"name\":\"RemovedFromWhitelist\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"SetGasTopUpLimit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"SetLoadLimit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"SetSpendLimit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"SubmittedGasTopUpLimitUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"SubmittedLoadLimitUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"SubmittedSpendLimitUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"_addresses\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"SubmittedWhitelistAddition\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"_addresses\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"SubmittedWhitelistRemoval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"ToppedUpGas\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"Transferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"TransferredOwnership\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"UpdatedAvailableLimit\",\"type\":\"event\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"constant\":true,\"inputs\":[],\"name\":\"WALLET_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_transactionBatch\",\"type\":\"bytes\"}],\"name\":\"batchExecuteTransaction\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"_assets\",\"type\":\"address[]\"}],\"name\":\"bulkTransfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_addresses\",\"type\":\"address[]\"}],\"name\":\"calculateHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"cancelWhitelistAddition\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"cancelWhitelistRemoval\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"confirmGasTopUpLimitUpdate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"confirmLoadLimitUpdate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"confirmSpendLimitUpdate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"confirmWhitelistAddition\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"confirmWhitelistRemoval\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"controllerNode\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"convertToEther\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"convertToStablecoin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ensRegistry\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"executeRelayedTransaction\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_destination\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"executeTransaction\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"gasTopUpLimitAvailable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"gasTopUpLimitPending\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"gasTopUpLimitUpdateable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"gasTopUpLimitValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"increaseRelayNonce\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isSetWhitelist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isTransferable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hashedData\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"isValidSignature\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"isValidSignature\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"licenceNode\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"loadLimitAvailable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"loadLimitPending\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"loadLimitUpdateable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"loadLimitValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"loadTokenCard\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingWhitelistAddition\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingWhitelistRemoval\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"relayNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"setGasTopUpLimit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"setLoadLimit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"setSpendLimit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_addresses\",\"type\":\"address[]\"}],\"name\":\"setWhitelist\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"spendLimitAvailable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"spendLimitPending\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"spendLimitUpdateable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"spendLimitValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"submitGasTopUpLimitUpdate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"submitLoadLimitUpdate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"submitSpendLimitUpdate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_addresses\",\"type\":\"address[]\"}],\"name\":\"submitWhitelistAddition\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_addresses\",\"type\":\"address[]\"}],\"name\":\"submitWhitelistRemoval\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"submittedWhitelistAddition\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"submittedWhitelistRemoval\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_interfaceID\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tokenWhitelistNode\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"topUpGas\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_transferable\",\"type\":\"bool\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"whitelistArray\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"whitelistMap\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// WalletBin is the compiled bytecode used for deploying new contracts.
var WalletBin = "0x60806040527f7f2ce995617d2816b426c5c8698c5ec2952f7a34bb10f38326f74933d58936976002553480156200003557600080fd5b5060405162005abd38038062005abd833981810160405260e08110156200005b57600080fd5b508051602082015160408301516060840151608085015160a086015160c090960151600180546001600160a01b038087166001600160a01b031992831617928390556000805490921692169190911790559495939492939192909190868685858482818686858015620000ce5760028190555b50600380546001600160a01b0319166001600160a01b0384161760ff60a01b1916600160a01b8315158102919091179182905560ff9104166200014857604080516001600160a01b038416815290517f808639ff9c8e4732d60b6c2330de498035416d229f27a77d259680895efec1229181900360200190a15b60408051600081526001600160a01b038416602082015281517f850b3df64837d7d518b45f5aa64d104652c3b80eb5b34a8e3d9eb666cb7cdea5929181900390910190a150506040805160a080820183528382526020808301859052428385018190526000606080860182905260809586018290526009889055600a97909755600b829055600c819055600d805460ff19908116909155600e98909855855193840186526706f05b59d3b20000808552928401839052948301819052948201849052910182905260108190556011556012919091556013819055601480549092169091559450620002479350506001600160e01b03620002be16915050565b50505050509150506000811162000295576040805162461bcd60e51b815260206004820152600d60248201526c37379039ba30b13632b1b7b4b760991b604482015290519081900360640190fd5b620002ad61271082026001600160e01b036200044016565b5050601b5550620005889350505050565b6060600080600080600080620002dc600e546200048e60201b60201c565b6001600160a01b0316633efec5e96040518163ffffffff1660e01b815260040160006040518083038186803b1580156200031557600080fd5b505afa1580156200032a573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405260e08110156200035457600080fd5b81019080805160405193929190846401000000008211156200037557600080fd5b9083019060208201858111156200038b57600080fd5b8251640100000000811182820188101715620003a657600080fd5b82525081516020918201929091019080838360005b83811015620003d5578181015183820152602001620003bb565b50505050905090810190601f168015620004035780820380516001836020036101000a031916815260200191505b5060409081526020820151908201516060830151608084015160a085015160c090950151979f939e50919c509a5098509096509294509192505050565b60158190556040805160a081018252828152602081018390524291810182905260006060820181905260809091018190526016839055601792909255601855601955601a805460ff19169055565b6000805460408051630178b8bf60e01b81526004810185905290516001600160a01b0390921691630178b8bf91602480820192602092909190829003018186803b158015620004dc57600080fd5b505afa158015620004f1573d6000803e3d6000fd5b505050506040513d60208110156200050857600080fd5b505160408051631d9dabef60e11b81526004810185905290516001600160a01b0390921691633b3b57de91602480820192602092909190829003018186803b1580156200055457600080fd5b505afa15801562000569573d6000803e3d6000fd5b505050506040513d60208110156200058057600080fd5b505192915050565b61552580620005986000396000f3fe60806040526004361061038c5760003560e01c80637d73b231116101dc578063cbd2ac6811610102578063de212bf3116100a0578063f36febda1161006f578063f36febda1461108c578063f40b51f8146110c5578063f41c4319146110ef578063f4217648146111195761038c565b8063de212bf31461100e578063e2b4ce9714611023578063e61c51ca14611038578063eadd3cea146110625761038c565b8063cd7958dd116100dc578063cd7958dd14610ef7578063ce0b5bd514610fa5578063d251fefc14610fcf578063da84b1ed14610ff95761038c565b8063cbd2ac6814610ea3578063cc0e7e5614610ecd578063cccdc55614610ee25761038c565b8063aaf1fc621161017a578063be40ba7911610149578063be40ba7914610e21578063beabacc814610e36578063c4856cd914610e79578063c785141814610e8e5761038c565b8063aaf1fc6214610cd2578063b221f31614610d83578063b242e53414610dad578063b87e21ef14610de85761038c565b8063877337b0116101b6578063877337b014610c7e5780638da5cb5b14610c9357806390e690c714610ca85780639b0dfd2714610cbd5761038c565b80637d73b23114610bbd5780637d7d004614610bee5780637fd004fa14610c035761038c565b8063294f4025116102c157806347b55a9d1161025f5780636137d6701161022e5780636137d67014610aee578063715018a614610b6957806374624c5514610b7e578063747c31d614610ba85761038c565b806347b55a9d14610a855780634aa46fde14610a9a5780635adc02ab14610aaf5780635d2362a814610ad95761038c565b80633bfec2541161029b5780633bfec254146108995780633c672eb7146108c35780633f579f42146108ed57806346efe0ed146109b35761038c565b8063294f4025146107d557806332531c3c1461083a5780633a43199f1461086d5761038c565b80631efd02991161032e57806321ce918d1161030857806321ce918d1461076c578063227149a3146107965780632587a6a2146107ab57806326d05ab2146107c05761038c565b80631efd02991461067757806320c13b0b1461068c5780632121dc75146107575761038c565b8063100f23fd1161036a578063100f23fd146104635780631127b57e1461048d5780631626ba7e146105175780631aa21fba146105ec5761038c565b806301ffc9a7146103c8578063027ef3eb146104105780630f3a85d814610437575b6040805133815234602082015281517f88a5966d370b9919b20f3e2c13ff65706f196a4e32cc2c12bf57088f88525874929181900390910190a1005b3480156103d457600080fd5b506103fc600480360360208110156103eb57600080fd5b50356001600160e01b031916611194565b604080519115158252519081900360200190f35b34801561041c57600080fd5b506104256111ae565b60408051918252519081900360200190f35b34801561044357600080fd5b506104616004803603602081101561045a57600080fd5b50356111b5565b005b34801561046f57600080fd5b506104616004803603602081101561048657600080fd5b50356112c1565b34801561049957600080fd5b506104a2611466565b6040805160208082528351818301528351919283929083019185019080838360005b838110156104dc5781810151838201526020016104c4565b50505050905090810190601f1680156105095780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561052357600080fd5b506105cf6004803603604081101561053a57600080fd5b81359190810190604081016020820135600160201b81111561055b57600080fd5b82018360208201111561056d57600080fd5b803590602001918460018302840111600160201b8311171561058e57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550611487945050505050565b604080516001600160e01b03199092168252519081900360200190f35b3480156105f857600080fd5b506104616004803603604081101561060f57600080fd5b6001600160a01b038235169190810190604081016020820135600160201b81111561063957600080fd5b82018360208201111561064b57600080fd5b803590602001918460208302840111600160201b8311171561066c57600080fd5b5090925090506114f5565b34801561068357600080fd5b5061042561167b565b34801561069857600080fd5b506105cf600480360360408110156106af57600080fd5b810190602081018135600160201b8111156106c957600080fd5b8201836020820111156106db57600080fd5b803590602001918460018302840111600160201b831117156106fc57600080fd5b919390929091602081019035600160201b81111561071957600080fd5b82018360208201111561072b57600080fd5b803590602001918460018302840111600160201b8311171561074c57600080fd5b50909250905061168c565b34801561076357600080fd5b506103fc611761565b34801561077857600080fd5b506104616004803603602081101561078f57600080fd5b5035611771565b3480156107a257600080fd5b506103fc61180f565b3480156107b757600080fd5b50610425611818565b3480156107cc57600080fd5b506103fc61181e565b3480156107e157600080fd5b506107ea611827565b60408051602080825283518183015283519192839290830191858101910280838360005b8381101561082657818101518382015260200161080e565b505050509050019250505060405180910390f35b34801561084657600080fd5b506103fc6004803603602081101561085d57600080fd5b50356001600160a01b0316611889565b6104616004803603604081101561088357600080fd5b506001600160a01b03813516906020013561189e565b3480156108a557600080fd5b50610461600480360360208110156108bc57600080fd5b5035611adc565b3480156108cf57600080fd5b50610461600480360360208110156108e657600080fd5b5035611be7565b3480156108f957600080fd5b506104a26004803603606081101561091057600080fd5b6001600160a01b0382351691602081013591810190606081016040820135600160201b81111561093f57600080fd5b82018360208201111561095157600080fd5b803590602001918460018302840111600160201b8311171561097257600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550611c8d945050505050565b3480156109bf57600080fd5b50610461600480360360608110156109d657600080fd5b81359190810190604081016020820135600160201b8111156109f757600080fd5b820183602082011115610a0957600080fd5b803590602001918460018302840111600160201b83111715610a2a57600080fd5b919390929091602081019035600160201b811115610a4757600080fd5b820183602082011115610a5957600080fd5b803590602001918460018302840111600160201b83111715610a7a57600080fd5b509092509050612184565b348015610a9157600080fd5b506107ea6124d3565b348015610aa657600080fd5b506103fc612533565b348015610abb57600080fd5b5061046160048036036020811015610ad257600080fd5b503561253c565b348015610ae557600080fd5b5061042561280c565b348015610afa57600080fd5b5061046160048036036020811015610b1157600080fd5b810190602081018135600160201b811115610b2b57600080fd5b820183602082011115610b3d57600080fd5b803590602001918460208302840111600160201b83111715610b5e57600080fd5b509092509050612818565b348015610b7557600080fd5b50610461612a3e565b348015610b8a57600080fd5b5061046160048036036020811015610ba157600080fd5b5035612b3c565b348015610bb457600080fd5b50610425612c40565b348015610bc957600080fd5b50610bd2612c46565b604080516001600160a01b039092168252519081900360200190f35b348015610bfa57600080fd5b50610425612c55565b348015610c0f57600080fd5b5061046160048036036020811015610c2657600080fd5b810190602081018135600160201b811115610c4057600080fd5b820183602082011115610c5257600080fd5b803590602001918460208302840111600160201b83111715610c7357600080fd5b509092509050612c61565b348015610c8a57600080fd5b50610425612fa3565b348015610c9f57600080fd5b50610bd2612fa9565b348015610cb457600080fd5b50610461612fb8565b348015610cc957600080fd5b50610425613053565b348015610cde57600080fd5b5061046160048036036020811015610cf557600080fd5b810190602081018135600160201b811115610d0f57600080fd5b820183602082011115610d2157600080fd5b803590602001918460018302840111600160201b83111715610d4257600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550613059945050505050565b348015610d8f57600080fd5b5061046160048036036020811015610da657600080fd5b5035613196565b348015610db957600080fd5b5061046160048036036040811015610dd057600080fd5b506001600160a01b0381351690602001351515613299565b348015610df457600080fd5b5061042560048036036040811015610e0b57600080fd5b506001600160a01b038135169060200135613453565b348015610e2d57600080fd5b506103fc6134e3565b348015610e4257600080fd5b5061046160048036036060811015610e5957600080fd5b506001600160a01b038135811691602081013590911690604001356134f2565b348015610e8557600080fd5b5061042561367c565b348015610e9a57600080fd5b506103fc613682565b348015610eaf57600080fd5b5061046160048036036020811015610ec657600080fd5b503561368b565b348015610ed957600080fd5b50610425613a08565b348015610eee57600080fd5b50610425613a0e565b348015610f0357600080fd5b5061042560048036036020811015610f1a57600080fd5b810190602081018135600160201b811115610f3457600080fd5b820183602082011115610f4657600080fd5b803590602001918460208302840111600160201b83111715610f6757600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929550613a14945050505050565b348015610fb157600080fd5b5061046160048036036020811015610fc857600080fd5b5035613a6e565b348015610fdb57600080fd5b50610bd260048036036020811015610ff257600080fd5b5035613c17565b34801561100557600080fd5b50610425613c3e565b34801561101a57600080fd5b506103fc613c44565b34801561102f57600080fd5b50610425613c52565b34801561104457600080fd5b506104616004803603602081101561105b57600080fd5b5035613c58565b34801561106e57600080fd5b506104616004803603602081101561108557600080fd5b5035613da2565b34801561109857600080fd5b50610425600480360360408110156110af57600080fd5b506001600160a01b038135169060200135613dfb565b3480156110d157600080fd5b50610461600480360360208110156110e857600080fd5b5035613fae565b3480156110fb57600080fd5b506104616004803603602081101561111257600080fd5b5035614007565b34801561112557600080fd5b506104616004803603602081101561113c57600080fd5b810190602081018135600160201b81111561115657600080fd5b82018360208201111561116857600080fd5b803590602001918460208302840111600160201b8311171561118957600080fd5b509092509050614060565b6001600160e01b031981166301ffc9a760e01b145b919050565b600c545b90565b6111be336143b2565b806111c857503330145b61120c576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b8066038d7ea4c680001115801561122b57506706f05b59d3b200008111155b611272576040805162461bcd60e51b815260206004820152601360248201527206f7574206f662072616e676520746f702d757606c1b604482015290519081900360640190fd5b61128360108263ffffffff6143c616565b604080513381526020810183905281517f41ff5d5ce3b7935893a4e7269ec5caae9cca5e3bf0eb4b21d2f443489667112e929181900390910190a150565b6112ca336143b2565b806112d957506112d933614432565b611323576040805162461bcd60e51b815260206004820152601660248201527537b7363c9037bbb732b93e3e31b7b73a3937b63632b960511b604482015290519081900360640190fd5b60085460ff16611372576040805162461bcd60e51b81526020600482015260156024820152743737903832b73234b7339039bab136b4b9b9b4b7b760591b604482015290519081900360640190fd5b6113d560068054806020026020016040519081016040528092919081815260200182805480156113cb57602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116113ad575b5050505050613a14565b81146114125760405162461bcd60e51b81526004018080602001828103825260238152602001806154986023913960400191505060405180910390fd5b61141e60066000615308565b6008805460ff19169055604080513381526020810183905281517f7794eff834d760583543e6e510e717a5e66d2c064e225f4db448343c3e66afcf929181900390910190a150565b604051806040016040528060058152602001640322e332e360dc1b81525081565b60008061149a848463ffffffff6144c616565b90506114a5816143b2565b6114e3576040805162461bcd60e51b815260206004820152600a60248201526937b7363c9037bbb732b960b11b604482015290519081900360640190fd5b50630b135d3f60e11b90505b92915050565b6114fe336143b2565b8061150857503330145b61154c576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b80611595576040805162461bcd60e51b8152602060048201526014602482015273617373657420617272617920697320656d70747960601b604482015290519081900360640190fd5b60005b818110156115f85760006115c7308585858181106115b257fe5b905060200201356001600160a01b03166145b4565b90506115ef858585858181106115d957fe5b905060200201356001600160a01b0316836134f2565b50600101611598565b507fd4f62f23021706247dcffea245d104ae7ddaec7f23acf3d11d7136d5de6a69ad83838360405180846001600160a01b03166001600160a01b03168152602001806020018281038252848482818152602001925060200280828437600083820152604051601f909101601f1916909201829003965090945050505050a1505050565b6000611687601661465f565b905090565b6000808585604051602001808383808284376040805191909301818103601f190182528084528151602092830120601f8b01839004830282018301909452898152929650630b135d3f60e11b9550611703945086935089915088908190840183828082843760009201919091525061148792505050565b6001600160e01b0319161461174f576040805162461bcd60e51b815260206004820152600d60248201526c1cda59c81b9bdd081d985b1a59609a1b604482015290519081900360640190fd5b506320c13b0b60e01b95945050505050565b600354600160a01b900460ff1690565b61177a336143b2565b8061178457503330145b6117c8576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b6117d960098263ffffffff61469416565b6040805182815290517f4b1b970c8a0fa761e7803ed70c13d7aca71904b13df60fbe03f981da1730da919181900360200190a150565b601a5460ff1690565b60105490565b60085460ff1681565b6060600780548060200260200160405190810160405280929190818152602001828054801561187f57602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311611861575b5050505050905090565b60046020526000908152604090205460ff1681565b6118a7336143b2565b806118b157503330145b6118f5576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b6118fe826146ee565b611944576040805162461bcd60e51b8152602060048201526012602482015271746f6b656e206e6f74206c6f616461626c6560701b604482015290519081900360640190fd5b60006119508383613dfb565b905061196360168263ffffffff61470816565b6000611970601b5461477e565b90506001600160a01b03841615611a185761199b6001600160a01b038516828563ffffffff61484016565b806001600160a01b0316631b3c96b485856040518363ffffffff1660e01b815260040180836001600160a01b03166001600160a01b0316815260200182815260200192505050600060405180830381600087803b1580156119fb57600080fd5b505af1158015611a0f573d6000803e3d6000fd5b50505050611a92565b806001600160a01b0316631b3c96b48486866040518463ffffffff1660e01b815260040180836001600160a01b03166001600160a01b03168152602001828152602001925050506000604051808303818588803b158015611a7857600080fd5b505af1158015611a8c573d6000803e3d6000fd5b50505050505b604080516001600160a01b03861681526020810185905281517f5f65674bec9af81f71be68674135a0ea3f163fb91984e3893d06da9f6ea2ce8a929181900390910190a150505050565b611ae5336143b2565b80611aef57503330145b611b33576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b8066038d7ea4c6800011158015611b4c57506015548111155b611b98576040805162461bcd60e51b81526020600482015260186024820152771bdd5d081bd9881c985b99d9481b1bd85908185b5bdd5b9d60421b604482015290519081900360640190fd5b611ba960168263ffffffff6143c616565b604080513381526020810183905281517f0b05243483e17c3f3377aee82b7d47e5700b48288695fc08b7ecc2759afa44ef929181900390910190a150565b611bf0336143b2565b80611bfa57503330145b611c3e576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b611c4f60098263ffffffff6143c616565b604080513381526020810183905281517f068f112e5ec923d412be64779fe69e0fcbb6784c6617e94cccc8fd348f2e0f21929181900390910190a150565b6060611c98336143b2565b80611ca257503330145b611ce6576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b6001600160a01b03841660009081526004602052604090205460ff16611d1757611d1760098463ffffffff61470816565b611d29846001600160a01b0316614958565b8015611d395750611d398461495e565b15611f2057600080611d4b8685614978565b6001600160a01b038216600090815260046020526040902054919350915060ff16611d91576000611d7c8783613453565b9050611d8f60098263ffffffff61470816565b505b611daa6001600160a01b0387168563ffffffff614a8216565b604080516020808252818301909252606091602082018180388339019050509050600160f81b81601f81518110611ddd57fe5b60200101906001600160f81b031916908160001a9053507ff77753fab406ecfff96d6ff2476c64a838fa9f6d37b1bf190f8546e395e3b6138787878460405180856001600160a01b03166001600160a01b031681526020018481526020018060200180602001838103835285818151815260200191508051906020019080838360005b83811015611e78578181015183820152602001611e60565b50505050905090810190601f168015611ea55780820380516001836020036101000a031916815260200191505b50838103825284518152845160209182019186019080838360005b83811015611ed8578181015183820152602001611ec0565b50505050905090810190601f168015611f055780820380516001836020036101000a031916815260200191505b50965050505050505060405180910390a1925061217d915050565b60006060856001600160a01b031685856040518082805190602001908083835b60208310611f5f5780518252601f199092019160209182019101611f40565b6001836020036101000a03801982511681845116808217855250505050505090500191505060006040518083038185875af1925050503d8060008114611fc1576040519150601f19603f3d011682016040523d82523d6000602084013e611fc6565b606091505b50915091508181906120565760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561201b578181015183820152602001612003565b50505050905090810190601f1680156120485780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b507ff77753fab406ecfff96d6ff2476c64a838fa9f6d37b1bf190f8546e395e3b6138686868460405180856001600160a01b03166001600160a01b031681526020018481526020018060200180602001838103835285818151815260200191508051906020019080838360005b838110156120db5781810151838201526020016120c3565b50505050905090810190601f1680156121085780820380516001836020036101000a031916815260200191505b50838103825284518152845160209182019186019080838360005b8381101561213b578181015183820152602001612123565b50505050905090810190601f1680156121685780820380516001836020036101000a031916815260200191505b50965050505050505060405180910390a19150505b9392505050565b61218d33614432565b6121cc576040805162461bcd60e51b815260206004820152601a602482015260008051602061540a833981519152604482015290519081900360640190fd5b600061221d86868660405160200180806339363c1d60e11b81525060040184815260200183838082843780830192505050935050505060405160208183030381529060405280519060200120614c40565b9050631626ba7e60e01b6001600160e01b0319166122718285858080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061148792505050565b6001600160e01b031916146122bd576040805162461bcd60e51b815260206004820152600d60248201526c1cda59c81b9bdd081d985b1a59609a1b604482015290519081900360640190fd5b600f5486146122ff576040805162461bcd60e51b81526020600482015260096024820152687478207265706c617960b81b604482015290519081900360640190fd5b600f80546001019081905560408051338152602081019290925280517fab0423a75986556234aecd171c46ce7f5e45607d8070bf5230f2735b50322bff9281900390910190a160006060306001600160a01b03168787604051808383808284376040519201945060009350909150508083038183865af19150503d80600081146123a5576040519150601f19603f3d011682016040523d82523d6000602084013e6123aa565b606091505b50915091508181906123fd5760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561201b578181015183820152602001612003565b507f823dbcf2b7b0f265871963ca65ac033f6b4c71e0d82cd123d2ff23d752dc21c187878360405180806020018060200183810383528686828181526020019250808284376000838201819052601f909101601f191690920185810384528651815286516020918201939188019250908190849084905b8381101561248c578181015183820152602001612474565b50505050905090810190601f1680156124b95780820380516001836020036101000a031916815260200191505b509550505050505060405180910390a15050505050505050565b6060600680548060200260200160405190810160405280929190818152602001828054801561187f576020028201919060005260206000209081546001600160a01b03168152600190910190602001808311611861575050505050905090565b600d5460ff1690565b61254533614432565b612584576040805162461bcd60e51b815260206004820152601a602482015260008051602061540a833981519152604482015290519081900360640190fd5b60085460ff166125d3576040805162461bcd60e51b81526020600482015260156024820152743737903832b73234b7339039bab136b4b9b9b4b7b760591b604482015290519081900360640190fd5b61263460068054806020026020016040519081016040528092919081815260200182805480156113cb576020028201919060005260206000209081546001600160a01b031681526001909101906020018083116113ad575050505050613a14565b81146126715760405162461bcd60e51b81526004018080602001828103825260238152602001806154986023913960400191505060405180910390fd5b60005b60065481101561275857600460006006838154811061268f57fe5b60009182526020808320909101546001600160a01b0316835282019290925260400190205460ff1661275057600160046000600684815481106126ce57fe5b6000918252602080832091909101546001600160a01b031683528201929092526040019020805460ff191691151591909117905560068054600591908390811061271457fe5b60009182526020808320909101548354600181018555938352912090910180546001600160a01b0319166001600160a01b039092169190911790555b600101612674565b507fb2f6cccee7a369e23e293c25aa19bef80af11eb26deba3ea0f2a02783f752e4a33600660405180836001600160a01b03166001600160a01b031681526020018060200182810382528381815481526020019150805480156127e457602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116127c6575b5050935050505060405180910390a16127ff60066000615308565b506008805460ff19169055565b6000611687600961465f565b612821336143b2565b8061282b57503330145b61286f576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b60085460ff1615801561288a5750600854610100900460ff16155b6128db576040805162461bcd60e51b815260206004820152601c60248201527f77686974656c6973742073756d62697373696f6e2070656e64696e6700000000604482015290519081900360640190fd5b60085462010000900460ff16612934576040805162461bcd60e51b81526020600482015260196024820152781dda1a5d195b1a5cdd081b9bdd081a5b9a5d1a585b1a5e9959603a1b604482015290519081900360640190fd5b80612978576040805162461bcd60e51b815260206004820152600f60248201526e195b5c1d1e481dda1a5d195b1a5cdd608a1b604482015290519081900360640190fd5b61298460078383615326565b506008805461ff00191661010017905560408051602080840282810182019093528382527ffbc0e5ca6c7e4858daf0fdb185ef5186203e74ec9c64737e93c0aeaec596e1d192859285926129f392859185918291850190849080828437600092019190915250613a1492505050565b60405180806020018381526020018281038252858582818152602001925060200280828437600083820152604051601f909101601f1916909201829003965090945050505050a15050565b612a47336143b2565b612a91576040805162461bcd60e51b815260206004820152601660248201527539b2b73232b91034b9903737ba1030b71037bbb732b960511b604482015290519081900360640190fd5b600354600160a01b900460ff16612aef576040805162461bcd60e51b815260206004820152601d60248201527f6f776e657273686970206973206e6f74207472616e7366657261626c65000000604482015290519081900360640190fd5b600380546001600160a01b0319169055604080516000808252602082015281517f850b3df64837d7d518b45f5aa64d104652c3b80eb5b34a8e3d9eb666cb7cdea5929181900390910190a1565b612b45336143b2565b80612b4f57503330145b612b93576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b8066038d7ea4c6800011158015612bb257506706f05b59d3b200008111155b612bf9576040805162461bcd60e51b815260206004820152601360248201527206f7574206f662072616e676520746f702d757606c1b604482015290519081900360640190fd5b612c0a60108263ffffffff61469416565b6040805182815290517faf2a77cd04c3cc155588dd3bf67b310ab4fb3b1da3cf6b8d7d4d2aa1d09b794c9181900360200190a150565b601b5490565b6001546001600160a01b031690565b6000611687601061465f565b612c6a336143b2565b80612c7457503330145b612cb8576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b60085460ff16158015612cd35750600854610100900460ff16155b612d24576040805162461bcd60e51b815260206004820152601c60248201527f77686974656c6973742073756d62697373696f6e2070656e64696e6700000000604482015290519081900360640190fd5b8181808060200260200160405190810160405280939291908181526020018383602002808284376000920182905250925050505b8151811015612e4057612d7d828281518110612d7057fe5b60200260200101516143b2565b15612dc8576040805162461bcd60e51b8152602060048201526016602482015275636f6e7461696e73206f776e6572206164647265737360501b604482015290519081900360640190fd5b60006001600160a01b0316828281518110612ddf57fe5b60200260200101516001600160a01b03161415612e38576040805162461bcd60e51b8152602060048201526012602482015271636f6e7461696e732030206164647265737360701b604482015290519081900360640190fd5b600101612d58565b5060085462010000900460ff16612e9a576040805162461bcd60e51b81526020600482015260196024820152781dda1a5d195b1a5cdd081b9bdd081a5b9a5d1a585b1a5e9959603a1b604482015290519081900360640190fd5b81612ede576040805162461bcd60e51b815260206004820152600f60248201526e195b5c1d1e481dda1a5d195b1a5cdd608a1b604482015290519081900360640190fd5b612eea60068484615326565b506008805460ff1916600117905560408051602080850282810182019093528482527f9c80b3b5f68b3e017766d59e8d09b34efe6462b05c398f35cab9e271d9bc3b9c9286928692612f5792859185918291850190849080828437600092019190915250613a1492505050565b60405180806020018381526020018281038252858582818152602001925060200280828437600083820152604051601f909101601f1916909201829003965090945050505050a1505050565b600e5490565b6003546001600160a01b031690565b612fc1336143b2565b61300b576040805162461bcd60e51b815260206004820152601660248201527539b2b73232b91034b9903737ba1030b71037bbb732b960511b604482015290519081900360640190fd5b600f80546001019081905560408051338152602081019290925280517fab0423a75986556234aecd171c46ce7f5e45607d8070bf5230f2735b50322bff9281900390910190a1565b60095490565b613062336143b2565b8061306c57503330145b6130b0576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b8051602080820191906000808060605b8685101561318c576130d986605463ffffffff614c9116565b888601805160148201516034909201805193995060609190911c9650909450909250905061311e6054613112878563ffffffff614cee16565b9063ffffffff614cee16565b945086851115613165576040805162461bcd60e51b815260206004820152600d60248201526c6f7574206f6620626f756e647360981b604482015290519081900360640190fd5b8161317b57506040805160208101909152600081525b613186848483611c8d565b506130c0565b5050505050505050565b61319f336143b2565b806131a957503330145b6131ed576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b8066038d7ea4c680001115801561320657506015548111155b613252576040805162461bcd60e51b81526020600482015260186024820152771bdd5d081bd9881c985b99d9481b1bd85908185b5bdd5b9d60421b604482015290519081900360640190fd5b61326360168263ffffffff61469416565b6040805182815290517fc178d379965e5657b6fc57494e392f121a14119215dfb422aad7db4cc03f2d109181900360200190a150565b6132a2336143b2565b6132ec576040805162461bcd60e51b815260206004820152601660248201527539b2b73232b91034b9903737ba1030b71037bbb732b960511b604482015290519081900360640190fd5b600354600160a01b900460ff1661334a576040805162461bcd60e51b815260206004820152601d60248201527f6f776e657273686970206973206e6f74207472616e7366657261626c65000000604482015290519081900360640190fd5b6001600160a01b03821661338f5760405162461bcd60e51b815260040180806020018281038252602381526020018061544b6023913960400191505060405180910390fd5b6003805460ff60a01b1916600160a01b83151502179055806133e857604080516001600160a01b038416815290517f808639ff9c8e4732d60b6c2330de498035416d229f27a77d259680895efec1229181900360200190a15b600354604080516001600160a01b039283168152918416602083015280517f850b3df64837d7d518b45f5aa64d104652c3b80eb5b34a8e3d9eb666cb7cdea59281900390910190a150600380546001600160a01b0319166001600160a01b0392909216919091179055565b60008060008061346286614d48565b5050509350935093505080156134d757816134ad576040805162461bcd60e51b81526020600482015260066024820152650726174653d360d41b604482015290519081900360640190fd5b6134cd836134c1878563ffffffff614eda16565b9063ffffffff614f3316565b93505050506114ef565b50600095945050505050565b60085462010000900460ff1681565b6134fb336143b2565b8061350557503330145b613549576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b8080613586576040805162461bcd60e51b8152602060048201526007602482015266076616c75653d360cc1b604482015290519081900360640190fd5b6001600160a01b0384166135d1576040805162461bcd60e51b815260206004820152600d60248201526c064657374696e6174696f6e3d3609c1b604482015290519081900360640190fd5b6001600160a01b03841660009081526004602052604090205460ff1661362157816001600160a01b0384161561360e5761360b8484613453565b90505b61361f60098263ffffffff61470816565b505b61362c848484614f9d565b604080516001600160a01b0380871682528516602082015280820184905290517fd1ba4ac2e2a11b5101f6cb4d978f514a155b421e8ec396d2d9abaf0bb02917ee9181900360600190a150505050565b60195490565b60145460ff1690565b61369433614432565b6136d3576040805162461bcd60e51b815260206004820152601a602482015260008051602061540a833981519152604482015290519081900360640190fd5b600854610100900460ff16613727576040805162461bcd60e51b81526020600482015260156024820152743737903832b73234b7339039bab136b4b9b9b4b7b760591b604482015290519081900360640190fd5b61378860078054806020026020016040519081016040528092919081815260200182805480156113cb576020028201919060005260206000209081546001600160a01b031681526001909101906020018083116113ad575050505050613a14565b81146137c55760405162461bcd60e51b81526004018080602001828103825260238152602001806154986023913960400191505060405180910390fd5b60005b6007548110156139535760046000600783815481106137e357fe5b60009182526020808320909101546001600160a01b0316835282019290925260400190205460ff161561394b576000600460006007848154811061382357fe5b6000918252602080832091909101546001600160a01b031683528201929092526040018120805460ff1916921515929092179091555b60055461386d90600163ffffffff614c9116565b811015613935576007828154811061388157fe5b600091825260209091200154600580546001600160a01b0390921691839081106138a757fe5b6000918252602090912001546001600160a01b0316141561392d576005805460001981019081106138d457fe5b600091825260209091200154600580546001600160a01b0390921691839081106138fa57fe5b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b03160217905550613935565b600101613859565b506005805490613949906000198301615389565b505b6001016137c8565b507fd218c430fa348f4ce67791021b6b89c0c3eacd4ead1d8f5b83c60038ec28249b33600760405180836001600160a01b03166001600160a01b031681526020018060200182810382528381815481526020019150805480156139df57602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116139c1575b5050935050505060405180910390a16139fa60076000615308565b506008805461ff0019169055565b60135490565b600f5481565b60008160405160200180828051906020019060200280838360005b83811015613a47578181015183820152602001613a2f565b50505050905001915050604051602081830303815290604052805190602001209050919050565b613a77336143b2565b80613a865750613a8633614432565b613ad0576040805162461bcd60e51b815260206004820152601660248201527537b7363c9037bbb732b93e3e31b7b73a3937b63632b960511b604482015290519081900360640190fd5b600854610100900460ff16613b24576040805162461bcd60e51b81526020600482015260156024820152743737903832b73234b7339039bab136b4b9b9b4b7b760591b604482015290519081900360640190fd5b613b8560078054806020026020016040519081016040528092919081815260200182805480156113cb576020028201919060005260206000209081546001600160a01b031681526001909101906020018083116113ad575050505050613a14565b8114613bc25760405162461bcd60e51b81526004018080602001828103825260238152602001806154986023913960400191505060405180910390fd5b613bce60076000615308565b6008805461ff0019169055604080513381526020810183905281517f13c935eb475aa0f6e931fece83e2ac44569ce2d53460d29a6dedab40b965c8a3929181900390910190a150565b60058181548110613c2457fe5b6000918252602090912001546001600160a01b0316905081565b60165490565b600854610100900460ff1681565b60025490565b8080613c95576040805162461bcd60e51b8152602060048201526007602482015266076616c75653d360cc1b604482015290519081900360640190fd5b613c9e336143b2565b80613cad5750613cad33614432565b613cf7576040805162461bcd60e51b815260206004820152601660248201527537b7363c9037bbb732b93e3e31b7b73a3937b63632b960511b604482015290519081900360640190fd5b613d0860108363ffffffff61470816565b613d10612fa9565b6001600160a01b03166108fc839081150290604051600060405180830381858888f19350505050158015613d48573d6000803e3d6000fd5b507f611b7c0d84fda988026215bef9b3e4d81cbceced7e679be6d5e044b588467c0e33613d73612fa9565b604080516001600160a01b03938416815291909216602082015280820185905290519081900360600190a15050565b613dab33614432565b613dea576040805162461bcd60e51b815260206004820152601a602482015260008051602061540a833981519152604482015290519081900360640190fd5b611c4f60098263ffffffff61500116565b6000613e05615055565b6001600160a01b0316836001600160a01b03161415613e255750806114ef565b816001600160a01b03841615613eea576000806000613e4387614d48565b5050509350935093505080613e95576040805162461bcd60e51b8152602060048201526013602482015272746f6b656e206e6f7420617661696c61626c6560681b604482015290519081900360640190fd5b81613ed0576040805162461bcd60e51b81526020600482015260066024820152650726174653d360d41b604482015290519081900360640190fd5b613ee4836134c1888563ffffffff614eda16565b93505050505b6000806000613ef76150cb565b5050509350935093505080613f49576040805162461bcd60e51b8152602060048201526013602482015272746f6b656e206e6f7420617661696c61626c6560681b604482015290519081900360640190fd5b81613f8f576040805162461bcd60e51b81526020600482015260116024820152700737461626c65636f696e20726174653d3607c1b604482015290519081900360640190fd5b613fa3826134c1868663ffffffff614eda16565b979650505050505050565b613fb733614432565b613ff6576040805162461bcd60e51b815260206004820152601a602482015260008051602061540a833981519152604482015290519081900360640190fd5b611ba960168263ffffffff61500116565b61401033614432565b61404f576040805162461bcd60e51b815260206004820152601a602482015260008051602061540a833981519152604482015290519081900360640190fd5b61128360108263ffffffff61500116565b614069336143b2565b8061407357503330145b6140b7576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b8181808060200260200160405190810160405280939291908181526020018383602002808284376000920182905250925050505b81518110156141c657614103828281518110612d7057fe5b1561414e576040805162461bcd60e51b8152602060048201526016602482015275636f6e7461696e73206f776e6572206164647265737360501b604482015290519081900360640190fd5b60006001600160a01b031682828151811061416557fe5b60200260200101516001600160a01b031614156141be576040805162461bcd60e51b8152602060048201526012602482015271636f6e7461696e732030206164647265737360701b604482015290519081900360640190fd5b6001016140eb565b5060085462010000900460ff161561421d576040805162461bcd60e51b81526020600482015260156024820152741dda1a5d195b1a5cdd081a5b9a5d1a585b1a5e9959605a1b604482015290519081900360640190fd5b60005b8281101561430e576004600085858481811061423857fe5b602090810292909201356001600160a01b03168352508101919091526040016000205460ff166143065760016004600086868581811061427457fe5b905060200201356001600160a01b03166001600160a01b03166001600160a01b0316815260200190815260200160002060006101000a81548160ff02191690831515021790555060058484838181106142c957fe5b835460018101855560009485526020948590200180546001600160a01b0319166001600160a01b0395909202939093013593909316929092179055505b600101614220565b506008805462ff0000191662010000179055604080513380825260208201838152600580549484018590527fb2f6cccee7a369e23e293c25aa19bef80af11eb26deba3ea0f2a02783f752e4a949293909290919060608301908490801561439e57602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311614380575b5050935050505060405180910390a1505050565b6003546001600160a01b0390811691161490565b600482015460ff1615614417576040805162461bcd60e51b81526020600482015260146024820152736c696d6974206e6f742075706461746561626c6560601b604482015290519081900360640190fd5b614421828261523a565b50600401805460ff19166001179055565b600061443f60025461477e565b6001600160a01b031663b429afeb836040518263ffffffff1660e01b815260040180826001600160a01b03166001600160a01b0316815260200191505060206040518083038186803b15801561449457600080fd5b505afa1580156144a8573d6000803e3d6000fd5b505050506040513d60208110156144be57600080fd5b505192915050565b600081516041146144d9575060006114ef565b60208201516040830151606084015160001a7f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a082111561451f57600093505050506114ef565b8060ff16601b1415801561453757508060ff16601c14155b1561454857600093505050506114ef565b6040805160008152602080820180845289905260ff8416828401526060820186905260808201859052915160019260a0808401939192601f1981019281900390910190855afa15801561459f573d6000803e3d6000fd5b5050604051601f190151979650505050505050565b60006001600160a01b0382161561464e57816001600160a01b03166370a08231846040518263ffffffff1660e01b815260040180826001600160a01b03166001600160a01b0316815260200191505060206040518083038186803b15801561461b57600080fd5b505afa15801561462f573d6000803e3d6000fd5b505050506040513d602081101561464557600080fd5b505190506114ef565b506001600160a01b038216316114ef565b600281015460009061467a906201518063ffffffff614cee16565b421115614689575080546111a9565b5060018101546111a9565b600482015460ff166146e6576040805162461bcd60e51b81526020600482015260166024820152756c696d6974207374696c6c2075706461746561626c6560501b604482015290519081900360640190fd5b600390910155565b6000806146fa83614d48565b509098975050505050505050565b6147118261525d565b808260010154101561475d576040805162461bcd60e51b815260206004820152601060248201526f185d985a5b18589b194f185b5bdd5b9d60821b604482015290519081900360640190fd5b6001820154614772908263ffffffff614c9116565b82600101819055505050565b6000805460408051630178b8bf60e01b81526004810185905290516001600160a01b0390921691630178b8bf91602480820192602092909190829003018186803b1580156147cb57600080fd5b505afa1580156147df573d6000803e3d6000fd5b505050506040513d60208110156147f557600080fd5b505160408051631d9dabef60e11b81526004810185905290516001600160a01b0390921691633b3b57de91602480820192602092909190829003018186803b15801561449457600080fd5b8015806148c6575060408051636eb1769f60e11b81523060048201526001600160a01b03848116602483015291519185169163dd62ed3e91604480820192602092909190829003018186803b15801561489857600080fd5b505afa1580156148ac573d6000803e3d6000fd5b505050506040513d60208110156148c257600080fd5b5051155b6149015760405162461bcd60e51b81526004018080602001828103825260368152602001806154bb6036913960400191505060405180910390fd5b604080516001600160a01b038416602482015260448082018490528251808303909101815260649091019091526020810180516001600160e01b031663095ea7b360e01b179052614953908490614a82565b505050565b3b151590565b60008061496a83614d48565b509198975050505050505050565b600080614986600e5461477e565b6001600160a01b031663afc72e9385856040518363ffffffff1660e01b815260040180836001600160a01b03166001600160a01b0316815260200180602001828103825283818151815260200191508051906020019080838360005b838110156149fa5781810151838201526020016149e2565b50505050905090810190601f168015614a275780820380516001836020036101000a031916815260200191505b509350505050604080518083038186803b158015614a4457600080fd5b505afa158015614a58573d6000803e3d6000fd5b505050506040513d6040811015614a6e57600080fd5b508051602090910151909590945092505050565b614a94826001600160a01b0316614958565b614ae5576040805162461bcd60e51b815260206004820152601f60248201527f5361666545524332303a2063616c6c20746f206e6f6e2d636f6e747261637400604482015290519081900360640190fd5b60006060836001600160a01b0316836040518082805190602001908083835b60208310614b235780518252601f199092019160209182019101614b04565b6001836020036101000a0380198251168184511680821785525050505050509050019150506000604051808303816000865af19150503d8060008114614b85576040519150601f19603f3d011682016040523d82523d6000602084013e614b8a565b606091505b509150915081614be1576040805162461bcd60e51b815260206004820181905260248201527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564604482015290519081900360640190fd5b805115614c3a57808060200190516020811015614bfd57600080fd5b5051614c3a5760405162461bcd60e51b815260040180806020018281038252602a81526020018061546e602a913960400191505060405180910390fd5b50505050565b604080517f19457468657265756d205369676e6564204d6573736167653a0a333200000000602080830191909152603c8083019490945282518083039094018452605c909101909152815191012090565b600082821115614ce8576040805162461bcd60e51b815260206004820152601e60248201527f536166654d6174683a207375627472616374696f6e206f766572666c6f770000604482015290519081900360640190fd5b50900390565b60008282018381101561217d576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b6060600080600080600080614d5e600e5461477e565b6001600160a01b0316631f69565f896040518263ffffffff1660e01b815260040180826001600160a01b03166001600160a01b0316815260200191505060006040518083038186803b158015614db357600080fd5b505afa158015614dc7573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405260e0811015614df057600080fd5b8101908080516040519392919084600160201b821115614e0f57600080fd5b908301906020820185811115614e2457600080fd5b8251600160201b811182820188101715614e3d57600080fd5b82525081516020918201929091019080838360005b83811015614e6a578181015183820152602001614e52565b50505050905090810190601f168015614e975780820380516001836020036101000a031916815260200191505b5060409081526020820151908201516060830151608084015160a085015160c090950151979e50929c50909a509850965094509192505050919395979092949650565b600082614ee9575060006114ef565b82820282848281614ef657fe5b041461217d5760405162461bcd60e51b815260040180806020018281038252602181526020018061542a6021913960400191505060405180910390fd5b6000808211614f89576040805162461bcd60e51b815260206004820152601a60248201527f536166654d6174683a206469766973696f6e206279207a65726f000000000000604482015290519081900360640190fd5b6000828481614f9457fe5b04949350505050565b6001600160a01b038216614fe7576040516001600160a01b0384169082156108fc029083906000818181858888f19350505050158015614fe1573d6000803e3d6000fd5b50614953565b6149536001600160a01b038316848363ffffffff6152b616565b808260030154146150435760405162461bcd60e51b81526004018080602001828103825260228152602001806153e86022913960400191505060405180910390fd5b61505182836003015461523a565b5050565b6000615062600e5461477e565b6001600160a01b031663e9cbd8226040518163ffffffff1660e01b815260040160206040518083038186803b15801561509a57600080fd5b505afa1580156150ae573d6000803e3d6000fd5b505050506040513d60208110156150c457600080fd5b5051905090565b60606000806000806000806150e1600e5461477e565b6001600160a01b0316633efec5e96040518163ffffffff1660e01b815260040160006040518083038186803b15801561511957600080fd5b505afa15801561512d573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405260e081101561515657600080fd5b8101908080516040519392919084600160201b82111561517557600080fd5b90830190602082018581111561518a57600080fd5b8251600160201b8111828201881017156151a357600080fd5b82525081516020918201929091019080838360005b838110156151d05781810151838201526020016151b8565b50505050905090810190601f1680156151fd5780820380516001836020036101000a031916815260200191505b5060409081526020820151908201516060830151608084015160a085015160c090950151979f939e50919c509a5098509096509294509192505050565b6152438261525d565b808255600182015481101561505157815460018301555050565b6002810154615275906201518063ffffffff614cee16565b4211156152b357426002820155805460018201556040517fe93bc25276d408d390778e7a8b926f2f67209c43ed540081b951fe128f0d3cd290600090a15b50565b604080516001600160a01b038416602482015260448082018490528251808303909101815260649091019091526020810180516001600160e01b031663a9059cbb60e01b179052614953908490614a82565b50805460008255906000526020600020908101906152b391906153a9565b828054828255906000526020600020908101928215615379579160200282015b828111156153795781546001600160a01b0319166001600160a01b03843516178255602090920191600190910190615346565b506153859291506153c3565b5090565b815481835581811115614953576000838152602090206149539181019083015b6111b291905b8082111561538557600081556001016153af565b6111b291905b808211156153855780546001600160a01b03191681556001016153c956fe636f6e6669726d65642f7375626d6974746564206c696d6974206d69736d6174636873656e646572206973206e6f74206120636f6e74726f6c6c6572000000000000536166654d6174683a206d756c7469706c69636174696f6e206f766572666c6f776f776e65722063616e6e6f742062652073657420746f207a65726f20616464726573735361666545524332303a204552433230206f7065726174696f6e20646964206e6f7420737563636565646e6f6e2d6d61746368696e672070656e64696e672077686974656c69737420686173685361666545524332303a20617070726f76652066726f6d206e6f6e2d7a65726f20746f206e6f6e2d7a65726f20616c6c6f77616e6365a265627a7a72315820fb531432806dc9cc3ad04e7caff8c310afc25de62936c0b055b0c5dfc6855cb264736f6c634300050f0032"

// DeployWallet deploys a new Ethereum contract, binding an instance of Wallet to it.
func DeployWallet(auth *bind.TransactOpts, backend bind.ContractBackend, _owner_ common.Address, _transferable_ bool, _ens_ common.Address, _tokenWhitelistNode_ [32]byte, _controllerNode_ [32]byte, _licenceNode_ [32]byte, _spendLimit_ *big.Int) (common.Address, *types.Transaction, *Wallet, error) {
	parsed, err := abi.JSON(strings.NewReader(WalletABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(WalletBin), backend, _owner_, _transferable_, _ens_, _tokenWhitelistNode_, _controllerNode_, _licenceNode_, _spendLimit_)
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

// WALLETVERSION is a free data retrieval call binding the contract method 0x1127b57e.
//
// Solidity: function WALLET_VERSION() constant returns(string)
func (_Wallet *WalletCaller) WALLETVERSION(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Wallet.contract.Call(opts, out, "WALLET_VERSION")
	return *ret0, err
}

// WALLETVERSION is a free data retrieval call binding the contract method 0x1127b57e.
//
// Solidity: function WALLET_VERSION() constant returns(string)
func (_Wallet *WalletSession) WALLETVERSION() (string, error) {
	return _Wallet.Contract.WALLETVERSION(&_Wallet.CallOpts)
}

// WALLETVERSION is a free data retrieval call binding the contract method 0x1127b57e.
//
// Solidity: function WALLET_VERSION() constant returns(string)
func (_Wallet *WalletCallerSession) WALLETVERSION() (string, error) {
	return _Wallet.Contract.WALLETVERSION(&_Wallet.CallOpts)
}

// CalculateHash is a free data retrieval call binding the contract method 0xcd7958dd.
//
// Solidity: function calculateHash(address[] _addresses) constant returns(bytes32)
func (_Wallet *WalletCaller) CalculateHash(opts *bind.CallOpts, _addresses []common.Address) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Wallet.contract.Call(opts, out, "calculateHash", _addresses)
	return *ret0, err
}

// CalculateHash is a free data retrieval call binding the contract method 0xcd7958dd.
//
// Solidity: function calculateHash(address[] _addresses) constant returns(bytes32)
func (_Wallet *WalletSession) CalculateHash(_addresses []common.Address) ([32]byte, error) {
	return _Wallet.Contract.CalculateHash(&_Wallet.CallOpts, _addresses)
}

// CalculateHash is a free data retrieval call binding the contract method 0xcd7958dd.
//
// Solidity: function calculateHash(address[] _addresses) constant returns(bytes32)
func (_Wallet *WalletCallerSession) CalculateHash(_addresses []common.Address) ([32]byte, error) {
	return _Wallet.Contract.CalculateHash(&_Wallet.CallOpts, _addresses)
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

// ConvertToEther is a free data retrieval call binding the contract method 0xb87e21ef.
//
// Solidity: function convertToEther(address _token, uint256 _amount) constant returns(uint256)
func (_Wallet *WalletCaller) ConvertToEther(opts *bind.CallOpts, _token common.Address, _amount *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Wallet.contract.Call(opts, out, "convertToEther", _token, _amount)
	return *ret0, err
}

// ConvertToEther is a free data retrieval call binding the contract method 0xb87e21ef.
//
// Solidity: function convertToEther(address _token, uint256 _amount) constant returns(uint256)
func (_Wallet *WalletSession) ConvertToEther(_token common.Address, _amount *big.Int) (*big.Int, error) {
	return _Wallet.Contract.ConvertToEther(&_Wallet.CallOpts, _token, _amount)
}

// ConvertToEther is a free data retrieval call binding the contract method 0xb87e21ef.
//
// Solidity: function convertToEther(address _token, uint256 _amount) constant returns(uint256)
func (_Wallet *WalletCallerSession) ConvertToEther(_token common.Address, _amount *big.Int) (*big.Int, error) {
	return _Wallet.Contract.ConvertToEther(&_Wallet.CallOpts, _token, _amount)
}

// ConvertToStablecoin is a free data retrieval call binding the contract method 0xf36febda.
//
// Solidity: function convertToStablecoin(address _token, uint256 _amount) constant returns(uint256)
func (_Wallet *WalletCaller) ConvertToStablecoin(opts *bind.CallOpts, _token common.Address, _amount *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Wallet.contract.Call(opts, out, "convertToStablecoin", _token, _amount)
	return *ret0, err
}

// ConvertToStablecoin is a free data retrieval call binding the contract method 0xf36febda.
//
// Solidity: function convertToStablecoin(address _token, uint256 _amount) constant returns(uint256)
func (_Wallet *WalletSession) ConvertToStablecoin(_token common.Address, _amount *big.Int) (*big.Int, error) {
	return _Wallet.Contract.ConvertToStablecoin(&_Wallet.CallOpts, _token, _amount)
}

// ConvertToStablecoin is a free data retrieval call binding the contract method 0xf36febda.
//
// Solidity: function convertToStablecoin(address _token, uint256 _amount) constant returns(uint256)
func (_Wallet *WalletCallerSession) ConvertToStablecoin(_token common.Address, _amount *big.Int) (*big.Int, error) {
	return _Wallet.Contract.ConvertToStablecoin(&_Wallet.CallOpts, _token, _amount)
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

// GasTopUpLimitAvailable is a free data retrieval call binding the contract method 0x7d7d0046.
//
// Solidity: function gasTopUpLimitAvailable() constant returns(uint256)
func (_Wallet *WalletCaller) GasTopUpLimitAvailable(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Wallet.contract.Call(opts, out, "gasTopUpLimitAvailable")
	return *ret0, err
}

// GasTopUpLimitAvailable is a free data retrieval call binding the contract method 0x7d7d0046.
//
// Solidity: function gasTopUpLimitAvailable() constant returns(uint256)
func (_Wallet *WalletSession) GasTopUpLimitAvailable() (*big.Int, error) {
	return _Wallet.Contract.GasTopUpLimitAvailable(&_Wallet.CallOpts)
}

// GasTopUpLimitAvailable is a free data retrieval call binding the contract method 0x7d7d0046.
//
// Solidity: function gasTopUpLimitAvailable() constant returns(uint256)
func (_Wallet *WalletCallerSession) GasTopUpLimitAvailable() (*big.Int, error) {
	return _Wallet.Contract.GasTopUpLimitAvailable(&_Wallet.CallOpts)
}

// GasTopUpLimitPending is a free data retrieval call binding the contract method 0xcc0e7e56.
//
// Solidity: function gasTopUpLimitPending() constant returns(uint256)
func (_Wallet *WalletCaller) GasTopUpLimitPending(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Wallet.contract.Call(opts, out, "gasTopUpLimitPending")
	return *ret0, err
}

// GasTopUpLimitPending is a free data retrieval call binding the contract method 0xcc0e7e56.
//
// Solidity: function gasTopUpLimitPending() constant returns(uint256)
func (_Wallet *WalletSession) GasTopUpLimitPending() (*big.Int, error) {
	return _Wallet.Contract.GasTopUpLimitPending(&_Wallet.CallOpts)
}

// GasTopUpLimitPending is a free data retrieval call binding the contract method 0xcc0e7e56.
//
// Solidity: function gasTopUpLimitPending() constant returns(uint256)
func (_Wallet *WalletCallerSession) GasTopUpLimitPending() (*big.Int, error) {
	return _Wallet.Contract.GasTopUpLimitPending(&_Wallet.CallOpts)
}

// GasTopUpLimitUpdateable is a free data retrieval call binding the contract method 0xc7851418.
//
// Solidity: function gasTopUpLimitUpdateable() constant returns(bool)
func (_Wallet *WalletCaller) GasTopUpLimitUpdateable(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Wallet.contract.Call(opts, out, "gasTopUpLimitUpdateable")
	return *ret0, err
}

// GasTopUpLimitUpdateable is a free data retrieval call binding the contract method 0xc7851418.
//
// Solidity: function gasTopUpLimitUpdateable() constant returns(bool)
func (_Wallet *WalletSession) GasTopUpLimitUpdateable() (bool, error) {
	return _Wallet.Contract.GasTopUpLimitUpdateable(&_Wallet.CallOpts)
}

// GasTopUpLimitUpdateable is a free data retrieval call binding the contract method 0xc7851418.
//
// Solidity: function gasTopUpLimitUpdateable() constant returns(bool)
func (_Wallet *WalletCallerSession) GasTopUpLimitUpdateable() (bool, error) {
	return _Wallet.Contract.GasTopUpLimitUpdateable(&_Wallet.CallOpts)
}

// GasTopUpLimitValue is a free data retrieval call binding the contract method 0x2587a6a2.
//
// Solidity: function gasTopUpLimitValue() constant returns(uint256)
func (_Wallet *WalletCaller) GasTopUpLimitValue(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Wallet.contract.Call(opts, out, "gasTopUpLimitValue")
	return *ret0, err
}

// GasTopUpLimitValue is a free data retrieval call binding the contract method 0x2587a6a2.
//
// Solidity: function gasTopUpLimitValue() constant returns(uint256)
func (_Wallet *WalletSession) GasTopUpLimitValue() (*big.Int, error) {
	return _Wallet.Contract.GasTopUpLimitValue(&_Wallet.CallOpts)
}

// GasTopUpLimitValue is a free data retrieval call binding the contract method 0x2587a6a2.
//
// Solidity: function gasTopUpLimitValue() constant returns(uint256)
func (_Wallet *WalletCallerSession) GasTopUpLimitValue() (*big.Int, error) {
	return _Wallet.Contract.GasTopUpLimitValue(&_Wallet.CallOpts)
}

// IsSetWhitelist is a free data retrieval call binding the contract method 0xbe40ba79.
//
// Solidity: function isSetWhitelist() constant returns(bool)
func (_Wallet *WalletCaller) IsSetWhitelist(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Wallet.contract.Call(opts, out, "isSetWhitelist")
	return *ret0, err
}

// IsSetWhitelist is a free data retrieval call binding the contract method 0xbe40ba79.
//
// Solidity: function isSetWhitelist() constant returns(bool)
func (_Wallet *WalletSession) IsSetWhitelist() (bool, error) {
	return _Wallet.Contract.IsSetWhitelist(&_Wallet.CallOpts)
}

// IsSetWhitelist is a free data retrieval call binding the contract method 0xbe40ba79.
//
// Solidity: function isSetWhitelist() constant returns(bool)
func (_Wallet *WalletCallerSession) IsSetWhitelist() (bool, error) {
	return _Wallet.Contract.IsSetWhitelist(&_Wallet.CallOpts)
}

// IsTransferable is a free data retrieval call binding the contract method 0x2121dc75.
//
// Solidity: function isTransferable() constant returns(bool)
func (_Wallet *WalletCaller) IsTransferable(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Wallet.contract.Call(opts, out, "isTransferable")
	return *ret0, err
}

// IsTransferable is a free data retrieval call binding the contract method 0x2121dc75.
//
// Solidity: function isTransferable() constant returns(bool)
func (_Wallet *WalletSession) IsTransferable() (bool, error) {
	return _Wallet.Contract.IsTransferable(&_Wallet.CallOpts)
}

// IsTransferable is a free data retrieval call binding the contract method 0x2121dc75.
//
// Solidity: function isTransferable() constant returns(bool)
func (_Wallet *WalletCallerSession) IsTransferable() (bool, error) {
	return _Wallet.Contract.IsTransferable(&_Wallet.CallOpts)
}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 _hashedData, bytes _signature) constant returns(bytes4)
func (_Wallet *WalletCaller) IsValidSignature(opts *bind.CallOpts, _hashedData [32]byte, _signature []byte) ([4]byte, error) {
	var (
		ret0 = new([4]byte)
	)
	out := ret0
	err := _Wallet.contract.Call(opts, out, "isValidSignature", _hashedData, _signature)
	return *ret0, err
}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 _hashedData, bytes _signature) constant returns(bytes4)
func (_Wallet *WalletSession) IsValidSignature(_hashedData [32]byte, _signature []byte) ([4]byte, error) {
	return _Wallet.Contract.IsValidSignature(&_Wallet.CallOpts, _hashedData, _signature)
}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 _hashedData, bytes _signature) constant returns(bytes4)
func (_Wallet *WalletCallerSession) IsValidSignature(_hashedData [32]byte, _signature []byte) ([4]byte, error) {
	return _Wallet.Contract.IsValidSignature(&_Wallet.CallOpts, _hashedData, _signature)
}

// IsValidSignature0 is a free data retrieval call binding the contract method 0x20c13b0b.
//
// Solidity: function isValidSignature(bytes _data, bytes _signature) constant returns(bytes4)
func (_Wallet *WalletCaller) IsValidSignature0(opts *bind.CallOpts, _data []byte, _signature []byte) ([4]byte, error) {
	var (
		ret0 = new([4]byte)
	)
	out := ret0
	err := _Wallet.contract.Call(opts, out, "isValidSignature0", _data, _signature)
	return *ret0, err
}

// IsValidSignature0 is a free data retrieval call binding the contract method 0x20c13b0b.
//
// Solidity: function isValidSignature(bytes _data, bytes _signature) constant returns(bytes4)
func (_Wallet *WalletSession) IsValidSignature0(_data []byte, _signature []byte) ([4]byte, error) {
	return _Wallet.Contract.IsValidSignature0(&_Wallet.CallOpts, _data, _signature)
}

// IsValidSignature0 is a free data retrieval call binding the contract method 0x20c13b0b.
//
// Solidity: function isValidSignature(bytes _data, bytes _signature) constant returns(bytes4)
func (_Wallet *WalletCallerSession) IsValidSignature0(_data []byte, _signature []byte) ([4]byte, error) {
	return _Wallet.Contract.IsValidSignature0(&_Wallet.CallOpts, _data, _signature)
}

// LicenceNode is a free data retrieval call binding the contract method 0x747c31d6.
//
// Solidity: function licenceNode() constant returns(bytes32)
func (_Wallet *WalletCaller) LicenceNode(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Wallet.contract.Call(opts, out, "licenceNode")
	return *ret0, err
}

// LicenceNode is a free data retrieval call binding the contract method 0x747c31d6.
//
// Solidity: function licenceNode() constant returns(bytes32)
func (_Wallet *WalletSession) LicenceNode() ([32]byte, error) {
	return _Wallet.Contract.LicenceNode(&_Wallet.CallOpts)
}

// LicenceNode is a free data retrieval call binding the contract method 0x747c31d6.
//
// Solidity: function licenceNode() constant returns(bytes32)
func (_Wallet *WalletCallerSession) LicenceNode() ([32]byte, error) {
	return _Wallet.Contract.LicenceNode(&_Wallet.CallOpts)
}

// LoadLimitAvailable is a free data retrieval call binding the contract method 0x1efd0299.
//
// Solidity: function loadLimitAvailable() constant returns(uint256)
func (_Wallet *WalletCaller) LoadLimitAvailable(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Wallet.contract.Call(opts, out, "loadLimitAvailable")
	return *ret0, err
}

// LoadLimitAvailable is a free data retrieval call binding the contract method 0x1efd0299.
//
// Solidity: function loadLimitAvailable() constant returns(uint256)
func (_Wallet *WalletSession) LoadLimitAvailable() (*big.Int, error) {
	return _Wallet.Contract.LoadLimitAvailable(&_Wallet.CallOpts)
}

// LoadLimitAvailable is a free data retrieval call binding the contract method 0x1efd0299.
//
// Solidity: function loadLimitAvailable() constant returns(uint256)
func (_Wallet *WalletCallerSession) LoadLimitAvailable() (*big.Int, error) {
	return _Wallet.Contract.LoadLimitAvailable(&_Wallet.CallOpts)
}

// LoadLimitPending is a free data retrieval call binding the contract method 0xc4856cd9.
//
// Solidity: function loadLimitPending() constant returns(uint256)
func (_Wallet *WalletCaller) LoadLimitPending(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Wallet.contract.Call(opts, out, "loadLimitPending")
	return *ret0, err
}

// LoadLimitPending is a free data retrieval call binding the contract method 0xc4856cd9.
//
// Solidity: function loadLimitPending() constant returns(uint256)
func (_Wallet *WalletSession) LoadLimitPending() (*big.Int, error) {
	return _Wallet.Contract.LoadLimitPending(&_Wallet.CallOpts)
}

// LoadLimitPending is a free data retrieval call binding the contract method 0xc4856cd9.
//
// Solidity: function loadLimitPending() constant returns(uint256)
func (_Wallet *WalletCallerSession) LoadLimitPending() (*big.Int, error) {
	return _Wallet.Contract.LoadLimitPending(&_Wallet.CallOpts)
}

// LoadLimitUpdateable is a free data retrieval call binding the contract method 0x227149a3.
//
// Solidity: function loadLimitUpdateable() constant returns(bool)
func (_Wallet *WalletCaller) LoadLimitUpdateable(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Wallet.contract.Call(opts, out, "loadLimitUpdateable")
	return *ret0, err
}

// LoadLimitUpdateable is a free data retrieval call binding the contract method 0x227149a3.
//
// Solidity: function loadLimitUpdateable() constant returns(bool)
func (_Wallet *WalletSession) LoadLimitUpdateable() (bool, error) {
	return _Wallet.Contract.LoadLimitUpdateable(&_Wallet.CallOpts)
}

// LoadLimitUpdateable is a free data retrieval call binding the contract method 0x227149a3.
//
// Solidity: function loadLimitUpdateable() constant returns(bool)
func (_Wallet *WalletCallerSession) LoadLimitUpdateable() (bool, error) {
	return _Wallet.Contract.LoadLimitUpdateable(&_Wallet.CallOpts)
}

// LoadLimitValue is a free data retrieval call binding the contract method 0xda84b1ed.
//
// Solidity: function loadLimitValue() constant returns(uint256)
func (_Wallet *WalletCaller) LoadLimitValue(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Wallet.contract.Call(opts, out, "loadLimitValue")
	return *ret0, err
}

// LoadLimitValue is a free data retrieval call binding the contract method 0xda84b1ed.
//
// Solidity: function loadLimitValue() constant returns(uint256)
func (_Wallet *WalletSession) LoadLimitValue() (*big.Int, error) {
	return _Wallet.Contract.LoadLimitValue(&_Wallet.CallOpts)
}

// LoadLimitValue is a free data retrieval call binding the contract method 0xda84b1ed.
//
// Solidity: function loadLimitValue() constant returns(uint256)
func (_Wallet *WalletCallerSession) LoadLimitValue() (*big.Int, error) {
	return _Wallet.Contract.LoadLimitValue(&_Wallet.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Wallet *WalletCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Wallet.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Wallet *WalletSession) Owner() (common.Address, error) {
	return _Wallet.Contract.Owner(&_Wallet.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Wallet *WalletCallerSession) Owner() (common.Address, error) {
	return _Wallet.Contract.Owner(&_Wallet.CallOpts)
}

// PendingWhitelistAddition is a free data retrieval call binding the contract method 0x47b55a9d.
//
// Solidity: function pendingWhitelistAddition() constant returns(address[])
func (_Wallet *WalletCaller) PendingWhitelistAddition(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _Wallet.contract.Call(opts, out, "pendingWhitelistAddition")
	return *ret0, err
}

// PendingWhitelistAddition is a free data retrieval call binding the contract method 0x47b55a9d.
//
// Solidity: function pendingWhitelistAddition() constant returns(address[])
func (_Wallet *WalletSession) PendingWhitelistAddition() ([]common.Address, error) {
	return _Wallet.Contract.PendingWhitelistAddition(&_Wallet.CallOpts)
}

// PendingWhitelistAddition is a free data retrieval call binding the contract method 0x47b55a9d.
//
// Solidity: function pendingWhitelistAddition() constant returns(address[])
func (_Wallet *WalletCallerSession) PendingWhitelistAddition() ([]common.Address, error) {
	return _Wallet.Contract.PendingWhitelistAddition(&_Wallet.CallOpts)
}

// PendingWhitelistRemoval is a free data retrieval call binding the contract method 0x294f4025.
//
// Solidity: function pendingWhitelistRemoval() constant returns(address[])
func (_Wallet *WalletCaller) PendingWhitelistRemoval(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _Wallet.contract.Call(opts, out, "pendingWhitelistRemoval")
	return *ret0, err
}

// PendingWhitelistRemoval is a free data retrieval call binding the contract method 0x294f4025.
//
// Solidity: function pendingWhitelistRemoval() constant returns(address[])
func (_Wallet *WalletSession) PendingWhitelistRemoval() ([]common.Address, error) {
	return _Wallet.Contract.PendingWhitelistRemoval(&_Wallet.CallOpts)
}

// PendingWhitelistRemoval is a free data retrieval call binding the contract method 0x294f4025.
//
// Solidity: function pendingWhitelistRemoval() constant returns(address[])
func (_Wallet *WalletCallerSession) PendingWhitelistRemoval() ([]common.Address, error) {
	return _Wallet.Contract.PendingWhitelistRemoval(&_Wallet.CallOpts)
}

// RelayNonce is a free data retrieval call binding the contract method 0xcccdc556.
//
// Solidity: function relayNonce() constant returns(uint256)
func (_Wallet *WalletCaller) RelayNonce(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Wallet.contract.Call(opts, out, "relayNonce")
	return *ret0, err
}

// RelayNonce is a free data retrieval call binding the contract method 0xcccdc556.
//
// Solidity: function relayNonce() constant returns(uint256)
func (_Wallet *WalletSession) RelayNonce() (*big.Int, error) {
	return _Wallet.Contract.RelayNonce(&_Wallet.CallOpts)
}

// RelayNonce is a free data retrieval call binding the contract method 0xcccdc556.
//
// Solidity: function relayNonce() constant returns(uint256)
func (_Wallet *WalletCallerSession) RelayNonce() (*big.Int, error) {
	return _Wallet.Contract.RelayNonce(&_Wallet.CallOpts)
}

// SpendLimitAvailable is a free data retrieval call binding the contract method 0x5d2362a8.
//
// Solidity: function spendLimitAvailable() constant returns(uint256)
func (_Wallet *WalletCaller) SpendLimitAvailable(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Wallet.contract.Call(opts, out, "spendLimitAvailable")
	return *ret0, err
}

// SpendLimitAvailable is a free data retrieval call binding the contract method 0x5d2362a8.
//
// Solidity: function spendLimitAvailable() constant returns(uint256)
func (_Wallet *WalletSession) SpendLimitAvailable() (*big.Int, error) {
	return _Wallet.Contract.SpendLimitAvailable(&_Wallet.CallOpts)
}

// SpendLimitAvailable is a free data retrieval call binding the contract method 0x5d2362a8.
//
// Solidity: function spendLimitAvailable() constant returns(uint256)
func (_Wallet *WalletCallerSession) SpendLimitAvailable() (*big.Int, error) {
	return _Wallet.Contract.SpendLimitAvailable(&_Wallet.CallOpts)
}

// SpendLimitPending is a free data retrieval call binding the contract method 0x027ef3eb.
//
// Solidity: function spendLimitPending() constant returns(uint256)
func (_Wallet *WalletCaller) SpendLimitPending(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Wallet.contract.Call(opts, out, "spendLimitPending")
	return *ret0, err
}

// SpendLimitPending is a free data retrieval call binding the contract method 0x027ef3eb.
//
// Solidity: function spendLimitPending() constant returns(uint256)
func (_Wallet *WalletSession) SpendLimitPending() (*big.Int, error) {
	return _Wallet.Contract.SpendLimitPending(&_Wallet.CallOpts)
}

// SpendLimitPending is a free data retrieval call binding the contract method 0x027ef3eb.
//
// Solidity: function spendLimitPending() constant returns(uint256)
func (_Wallet *WalletCallerSession) SpendLimitPending() (*big.Int, error) {
	return _Wallet.Contract.SpendLimitPending(&_Wallet.CallOpts)
}

// SpendLimitUpdateable is a free data retrieval call binding the contract method 0x4aa46fde.
//
// Solidity: function spendLimitUpdateable() constant returns(bool)
func (_Wallet *WalletCaller) SpendLimitUpdateable(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Wallet.contract.Call(opts, out, "spendLimitUpdateable")
	return *ret0, err
}

// SpendLimitUpdateable is a free data retrieval call binding the contract method 0x4aa46fde.
//
// Solidity: function spendLimitUpdateable() constant returns(bool)
func (_Wallet *WalletSession) SpendLimitUpdateable() (bool, error) {
	return _Wallet.Contract.SpendLimitUpdateable(&_Wallet.CallOpts)
}

// SpendLimitUpdateable is a free data retrieval call binding the contract method 0x4aa46fde.
//
// Solidity: function spendLimitUpdateable() constant returns(bool)
func (_Wallet *WalletCallerSession) SpendLimitUpdateable() (bool, error) {
	return _Wallet.Contract.SpendLimitUpdateable(&_Wallet.CallOpts)
}

// SpendLimitValue is a free data retrieval call binding the contract method 0x9b0dfd27.
//
// Solidity: function spendLimitValue() constant returns(uint256)
func (_Wallet *WalletCaller) SpendLimitValue(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Wallet.contract.Call(opts, out, "spendLimitValue")
	return *ret0, err
}

// SpendLimitValue is a free data retrieval call binding the contract method 0x9b0dfd27.
//
// Solidity: function spendLimitValue() constant returns(uint256)
func (_Wallet *WalletSession) SpendLimitValue() (*big.Int, error) {
	return _Wallet.Contract.SpendLimitValue(&_Wallet.CallOpts)
}

// SpendLimitValue is a free data retrieval call binding the contract method 0x9b0dfd27.
//
// Solidity: function spendLimitValue() constant returns(uint256)
func (_Wallet *WalletCallerSession) SpendLimitValue() (*big.Int, error) {
	return _Wallet.Contract.SpendLimitValue(&_Wallet.CallOpts)
}

// SubmittedWhitelistAddition is a free data retrieval call binding the contract method 0x26d05ab2.
//
// Solidity: function submittedWhitelistAddition() constant returns(bool)
func (_Wallet *WalletCaller) SubmittedWhitelistAddition(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Wallet.contract.Call(opts, out, "submittedWhitelistAddition")
	return *ret0, err
}

// SubmittedWhitelistAddition is a free data retrieval call binding the contract method 0x26d05ab2.
//
// Solidity: function submittedWhitelistAddition() constant returns(bool)
func (_Wallet *WalletSession) SubmittedWhitelistAddition() (bool, error) {
	return _Wallet.Contract.SubmittedWhitelistAddition(&_Wallet.CallOpts)
}

// SubmittedWhitelistAddition is a free data retrieval call binding the contract method 0x26d05ab2.
//
// Solidity: function submittedWhitelistAddition() constant returns(bool)
func (_Wallet *WalletCallerSession) SubmittedWhitelistAddition() (bool, error) {
	return _Wallet.Contract.SubmittedWhitelistAddition(&_Wallet.CallOpts)
}

// SubmittedWhitelistRemoval is a free data retrieval call binding the contract method 0xde212bf3.
//
// Solidity: function submittedWhitelistRemoval() constant returns(bool)
func (_Wallet *WalletCaller) SubmittedWhitelistRemoval(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Wallet.contract.Call(opts, out, "submittedWhitelistRemoval")
	return *ret0, err
}

// SubmittedWhitelistRemoval is a free data retrieval call binding the contract method 0xde212bf3.
//
// Solidity: function submittedWhitelistRemoval() constant returns(bool)
func (_Wallet *WalletSession) SubmittedWhitelistRemoval() (bool, error) {
	return _Wallet.Contract.SubmittedWhitelistRemoval(&_Wallet.CallOpts)
}

// SubmittedWhitelistRemoval is a free data retrieval call binding the contract method 0xde212bf3.
//
// Solidity: function submittedWhitelistRemoval() constant returns(bool)
func (_Wallet *WalletCallerSession) SubmittedWhitelistRemoval() (bool, error) {
	return _Wallet.Contract.SubmittedWhitelistRemoval(&_Wallet.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceID) constant returns(bool)
func (_Wallet *WalletCaller) SupportsInterface(opts *bind.CallOpts, _interfaceID [4]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Wallet.contract.Call(opts, out, "supportsInterface", _interfaceID)
	return *ret0, err
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceID) constant returns(bool)
func (_Wallet *WalletSession) SupportsInterface(_interfaceID [4]byte) (bool, error) {
	return _Wallet.Contract.SupportsInterface(&_Wallet.CallOpts, _interfaceID)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceID) constant returns(bool)
func (_Wallet *WalletCallerSession) SupportsInterface(_interfaceID [4]byte) (bool, error) {
	return _Wallet.Contract.SupportsInterface(&_Wallet.CallOpts, _interfaceID)
}

// TokenWhitelistNode is a free data retrieval call binding the contract method 0x877337b0.
//
// Solidity: function tokenWhitelistNode() constant returns(bytes32)
func (_Wallet *WalletCaller) TokenWhitelistNode(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Wallet.contract.Call(opts, out, "tokenWhitelistNode")
	return *ret0, err
}

// TokenWhitelistNode is a free data retrieval call binding the contract method 0x877337b0.
//
// Solidity: function tokenWhitelistNode() constant returns(bytes32)
func (_Wallet *WalletSession) TokenWhitelistNode() ([32]byte, error) {
	return _Wallet.Contract.TokenWhitelistNode(&_Wallet.CallOpts)
}

// TokenWhitelistNode is a free data retrieval call binding the contract method 0x877337b0.
//
// Solidity: function tokenWhitelistNode() constant returns(bytes32)
func (_Wallet *WalletCallerSession) TokenWhitelistNode() ([32]byte, error) {
	return _Wallet.Contract.TokenWhitelistNode(&_Wallet.CallOpts)
}

// WhitelistArray is a free data retrieval call binding the contract method 0xd251fefc.
//
// Solidity: function whitelistArray(uint256 ) constant returns(address)
func (_Wallet *WalletCaller) WhitelistArray(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Wallet.contract.Call(opts, out, "whitelistArray", arg0)
	return *ret0, err
}

// WhitelistArray is a free data retrieval call binding the contract method 0xd251fefc.
//
// Solidity: function whitelistArray(uint256 ) constant returns(address)
func (_Wallet *WalletSession) WhitelistArray(arg0 *big.Int) (common.Address, error) {
	return _Wallet.Contract.WhitelistArray(&_Wallet.CallOpts, arg0)
}

// WhitelistArray is a free data retrieval call binding the contract method 0xd251fefc.
//
// Solidity: function whitelistArray(uint256 ) constant returns(address)
func (_Wallet *WalletCallerSession) WhitelistArray(arg0 *big.Int) (common.Address, error) {
	return _Wallet.Contract.WhitelistArray(&_Wallet.CallOpts, arg0)
}

// WhitelistMap is a free data retrieval call binding the contract method 0x32531c3c.
//
// Solidity: function whitelistMap(address ) constant returns(bool)
func (_Wallet *WalletCaller) WhitelistMap(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Wallet.contract.Call(opts, out, "whitelistMap", arg0)
	return *ret0, err
}

// WhitelistMap is a free data retrieval call binding the contract method 0x32531c3c.
//
// Solidity: function whitelistMap(address ) constant returns(bool)
func (_Wallet *WalletSession) WhitelistMap(arg0 common.Address) (bool, error) {
	return _Wallet.Contract.WhitelistMap(&_Wallet.CallOpts, arg0)
}

// WhitelistMap is a free data retrieval call binding the contract method 0x32531c3c.
//
// Solidity: function whitelistMap(address ) constant returns(bool)
func (_Wallet *WalletCallerSession) WhitelistMap(arg0 common.Address) (bool, error) {
	return _Wallet.Contract.WhitelistMap(&_Wallet.CallOpts, arg0)
}

// BatchExecuteTransaction is a paid mutator transaction binding the contract method 0xaaf1fc62.
//
// Solidity: function batchExecuteTransaction(bytes _transactionBatch) returns()
func (_Wallet *WalletTransactor) BatchExecuteTransaction(opts *bind.TransactOpts, _transactionBatch []byte) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "batchExecuteTransaction", _transactionBatch)
}

// BatchExecuteTransaction is a paid mutator transaction binding the contract method 0xaaf1fc62.
//
// Solidity: function batchExecuteTransaction(bytes _transactionBatch) returns()
func (_Wallet *WalletSession) BatchExecuteTransaction(_transactionBatch []byte) (*types.Transaction, error) {
	return _Wallet.Contract.BatchExecuteTransaction(&_Wallet.TransactOpts, _transactionBatch)
}

// BatchExecuteTransaction is a paid mutator transaction binding the contract method 0xaaf1fc62.
//
// Solidity: function batchExecuteTransaction(bytes _transactionBatch) returns()
func (_Wallet *WalletTransactorSession) BatchExecuteTransaction(_transactionBatch []byte) (*types.Transaction, error) {
	return _Wallet.Contract.BatchExecuteTransaction(&_Wallet.TransactOpts, _transactionBatch)
}

// BulkTransfer is a paid mutator transaction binding the contract method 0x1aa21fba.
//
// Solidity: function bulkTransfer(address _to, address[] _assets) returns()
func (_Wallet *WalletTransactor) BulkTransfer(opts *bind.TransactOpts, _to common.Address, _assets []common.Address) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "bulkTransfer", _to, _assets)
}

// BulkTransfer is a paid mutator transaction binding the contract method 0x1aa21fba.
//
// Solidity: function bulkTransfer(address _to, address[] _assets) returns()
func (_Wallet *WalletSession) BulkTransfer(_to common.Address, _assets []common.Address) (*types.Transaction, error) {
	return _Wallet.Contract.BulkTransfer(&_Wallet.TransactOpts, _to, _assets)
}

// BulkTransfer is a paid mutator transaction binding the contract method 0x1aa21fba.
//
// Solidity: function bulkTransfer(address _to, address[] _assets) returns()
func (_Wallet *WalletTransactorSession) BulkTransfer(_to common.Address, _assets []common.Address) (*types.Transaction, error) {
	return _Wallet.Contract.BulkTransfer(&_Wallet.TransactOpts, _to, _assets)
}

// CancelWhitelistAddition is a paid mutator transaction binding the contract method 0x100f23fd.
//
// Solidity: function cancelWhitelistAddition(bytes32 _hash) returns()
func (_Wallet *WalletTransactor) CancelWhitelistAddition(opts *bind.TransactOpts, _hash [32]byte) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "cancelWhitelistAddition", _hash)
}

// CancelWhitelistAddition is a paid mutator transaction binding the contract method 0x100f23fd.
//
// Solidity: function cancelWhitelistAddition(bytes32 _hash) returns()
func (_Wallet *WalletSession) CancelWhitelistAddition(_hash [32]byte) (*types.Transaction, error) {
	return _Wallet.Contract.CancelWhitelistAddition(&_Wallet.TransactOpts, _hash)
}

// CancelWhitelistAddition is a paid mutator transaction binding the contract method 0x100f23fd.
//
// Solidity: function cancelWhitelistAddition(bytes32 _hash) returns()
func (_Wallet *WalletTransactorSession) CancelWhitelistAddition(_hash [32]byte) (*types.Transaction, error) {
	return _Wallet.Contract.CancelWhitelistAddition(&_Wallet.TransactOpts, _hash)
}

// CancelWhitelistRemoval is a paid mutator transaction binding the contract method 0xce0b5bd5.
//
// Solidity: function cancelWhitelistRemoval(bytes32 _hash) returns()
func (_Wallet *WalletTransactor) CancelWhitelistRemoval(opts *bind.TransactOpts, _hash [32]byte) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "cancelWhitelistRemoval", _hash)
}

// CancelWhitelistRemoval is a paid mutator transaction binding the contract method 0xce0b5bd5.
//
// Solidity: function cancelWhitelistRemoval(bytes32 _hash) returns()
func (_Wallet *WalletSession) CancelWhitelistRemoval(_hash [32]byte) (*types.Transaction, error) {
	return _Wallet.Contract.CancelWhitelistRemoval(&_Wallet.TransactOpts, _hash)
}

// CancelWhitelistRemoval is a paid mutator transaction binding the contract method 0xce0b5bd5.
//
// Solidity: function cancelWhitelistRemoval(bytes32 _hash) returns()
func (_Wallet *WalletTransactorSession) CancelWhitelistRemoval(_hash [32]byte) (*types.Transaction, error) {
	return _Wallet.Contract.CancelWhitelistRemoval(&_Wallet.TransactOpts, _hash)
}

// ConfirmGasTopUpLimitUpdate is a paid mutator transaction binding the contract method 0xf41c4319.
//
// Solidity: function confirmGasTopUpLimitUpdate(uint256 _amount) returns()
func (_Wallet *WalletTransactor) ConfirmGasTopUpLimitUpdate(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "confirmGasTopUpLimitUpdate", _amount)
}

// ConfirmGasTopUpLimitUpdate is a paid mutator transaction binding the contract method 0xf41c4319.
//
// Solidity: function confirmGasTopUpLimitUpdate(uint256 _amount) returns()
func (_Wallet *WalletSession) ConfirmGasTopUpLimitUpdate(_amount *big.Int) (*types.Transaction, error) {
	return _Wallet.Contract.ConfirmGasTopUpLimitUpdate(&_Wallet.TransactOpts, _amount)
}

// ConfirmGasTopUpLimitUpdate is a paid mutator transaction binding the contract method 0xf41c4319.
//
// Solidity: function confirmGasTopUpLimitUpdate(uint256 _amount) returns()
func (_Wallet *WalletTransactorSession) ConfirmGasTopUpLimitUpdate(_amount *big.Int) (*types.Transaction, error) {
	return _Wallet.Contract.ConfirmGasTopUpLimitUpdate(&_Wallet.TransactOpts, _amount)
}

// ConfirmLoadLimitUpdate is a paid mutator transaction binding the contract method 0xf40b51f8.
//
// Solidity: function confirmLoadLimitUpdate(uint256 _amount) returns()
func (_Wallet *WalletTransactor) ConfirmLoadLimitUpdate(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "confirmLoadLimitUpdate", _amount)
}

// ConfirmLoadLimitUpdate is a paid mutator transaction binding the contract method 0xf40b51f8.
//
// Solidity: function confirmLoadLimitUpdate(uint256 _amount) returns()
func (_Wallet *WalletSession) ConfirmLoadLimitUpdate(_amount *big.Int) (*types.Transaction, error) {
	return _Wallet.Contract.ConfirmLoadLimitUpdate(&_Wallet.TransactOpts, _amount)
}

// ConfirmLoadLimitUpdate is a paid mutator transaction binding the contract method 0xf40b51f8.
//
// Solidity: function confirmLoadLimitUpdate(uint256 _amount) returns()
func (_Wallet *WalletTransactorSession) ConfirmLoadLimitUpdate(_amount *big.Int) (*types.Transaction, error) {
	return _Wallet.Contract.ConfirmLoadLimitUpdate(&_Wallet.TransactOpts, _amount)
}

// ConfirmSpendLimitUpdate is a paid mutator transaction binding the contract method 0xeadd3cea.
//
// Solidity: function confirmSpendLimitUpdate(uint256 _amount) returns()
func (_Wallet *WalletTransactor) ConfirmSpendLimitUpdate(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "confirmSpendLimitUpdate", _amount)
}

// ConfirmSpendLimitUpdate is a paid mutator transaction binding the contract method 0xeadd3cea.
//
// Solidity: function confirmSpendLimitUpdate(uint256 _amount) returns()
func (_Wallet *WalletSession) ConfirmSpendLimitUpdate(_amount *big.Int) (*types.Transaction, error) {
	return _Wallet.Contract.ConfirmSpendLimitUpdate(&_Wallet.TransactOpts, _amount)
}

// ConfirmSpendLimitUpdate is a paid mutator transaction binding the contract method 0xeadd3cea.
//
// Solidity: function confirmSpendLimitUpdate(uint256 _amount) returns()
func (_Wallet *WalletTransactorSession) ConfirmSpendLimitUpdate(_amount *big.Int) (*types.Transaction, error) {
	return _Wallet.Contract.ConfirmSpendLimitUpdate(&_Wallet.TransactOpts, _amount)
}

// ConfirmWhitelistAddition is a paid mutator transaction binding the contract method 0x5adc02ab.
//
// Solidity: function confirmWhitelistAddition(bytes32 _hash) returns()
func (_Wallet *WalletTransactor) ConfirmWhitelistAddition(opts *bind.TransactOpts, _hash [32]byte) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "confirmWhitelistAddition", _hash)
}

// ConfirmWhitelistAddition is a paid mutator transaction binding the contract method 0x5adc02ab.
//
// Solidity: function confirmWhitelistAddition(bytes32 _hash) returns()
func (_Wallet *WalletSession) ConfirmWhitelistAddition(_hash [32]byte) (*types.Transaction, error) {
	return _Wallet.Contract.ConfirmWhitelistAddition(&_Wallet.TransactOpts, _hash)
}

// ConfirmWhitelistAddition is a paid mutator transaction binding the contract method 0x5adc02ab.
//
// Solidity: function confirmWhitelistAddition(bytes32 _hash) returns()
func (_Wallet *WalletTransactorSession) ConfirmWhitelistAddition(_hash [32]byte) (*types.Transaction, error) {
	return _Wallet.Contract.ConfirmWhitelistAddition(&_Wallet.TransactOpts, _hash)
}

// ConfirmWhitelistRemoval is a paid mutator transaction binding the contract method 0xcbd2ac68.
//
// Solidity: function confirmWhitelistRemoval(bytes32 _hash) returns()
func (_Wallet *WalletTransactor) ConfirmWhitelistRemoval(opts *bind.TransactOpts, _hash [32]byte) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "confirmWhitelistRemoval", _hash)
}

// ConfirmWhitelistRemoval is a paid mutator transaction binding the contract method 0xcbd2ac68.
//
// Solidity: function confirmWhitelistRemoval(bytes32 _hash) returns()
func (_Wallet *WalletSession) ConfirmWhitelistRemoval(_hash [32]byte) (*types.Transaction, error) {
	return _Wallet.Contract.ConfirmWhitelistRemoval(&_Wallet.TransactOpts, _hash)
}

// ConfirmWhitelistRemoval is a paid mutator transaction binding the contract method 0xcbd2ac68.
//
// Solidity: function confirmWhitelistRemoval(bytes32 _hash) returns()
func (_Wallet *WalletTransactorSession) ConfirmWhitelistRemoval(_hash [32]byte) (*types.Transaction, error) {
	return _Wallet.Contract.ConfirmWhitelistRemoval(&_Wallet.TransactOpts, _hash)
}

// ExecuteRelayedTransaction is a paid mutator transaction binding the contract method 0x46efe0ed.
//
// Solidity: function executeRelayedTransaction(uint256 _nonce, bytes _data, bytes _signature) returns()
func (_Wallet *WalletTransactor) ExecuteRelayedTransaction(opts *bind.TransactOpts, _nonce *big.Int, _data []byte, _signature []byte) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "executeRelayedTransaction", _nonce, _data, _signature)
}

// ExecuteRelayedTransaction is a paid mutator transaction binding the contract method 0x46efe0ed.
//
// Solidity: function executeRelayedTransaction(uint256 _nonce, bytes _data, bytes _signature) returns()
func (_Wallet *WalletSession) ExecuteRelayedTransaction(_nonce *big.Int, _data []byte, _signature []byte) (*types.Transaction, error) {
	return _Wallet.Contract.ExecuteRelayedTransaction(&_Wallet.TransactOpts, _nonce, _data, _signature)
}

// ExecuteRelayedTransaction is a paid mutator transaction binding the contract method 0x46efe0ed.
//
// Solidity: function executeRelayedTransaction(uint256 _nonce, bytes _data, bytes _signature) returns()
func (_Wallet *WalletTransactorSession) ExecuteRelayedTransaction(_nonce *big.Int, _data []byte, _signature []byte) (*types.Transaction, error) {
	return _Wallet.Contract.ExecuteRelayedTransaction(&_Wallet.TransactOpts, _nonce, _data, _signature)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0x3f579f42.
//
// Solidity: function executeTransaction(address _destination, uint256 _value, bytes _data) returns(bytes)
func (_Wallet *WalletTransactor) ExecuteTransaction(opts *bind.TransactOpts, _destination common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "executeTransaction", _destination, _value, _data)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0x3f579f42.
//
// Solidity: function executeTransaction(address _destination, uint256 _value, bytes _data) returns(bytes)
func (_Wallet *WalletSession) ExecuteTransaction(_destination common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _Wallet.Contract.ExecuteTransaction(&_Wallet.TransactOpts, _destination, _value, _data)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0x3f579f42.
//
// Solidity: function executeTransaction(address _destination, uint256 _value, bytes _data) returns(bytes)
func (_Wallet *WalletTransactorSession) ExecuteTransaction(_destination common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _Wallet.Contract.ExecuteTransaction(&_Wallet.TransactOpts, _destination, _value, _data)
}

// IncreaseRelayNonce is a paid mutator transaction binding the contract method 0x90e690c7.
//
// Solidity: function increaseRelayNonce() returns()
func (_Wallet *WalletTransactor) IncreaseRelayNonce(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "increaseRelayNonce")
}

// IncreaseRelayNonce is a paid mutator transaction binding the contract method 0x90e690c7.
//
// Solidity: function increaseRelayNonce() returns()
func (_Wallet *WalletSession) IncreaseRelayNonce() (*types.Transaction, error) {
	return _Wallet.Contract.IncreaseRelayNonce(&_Wallet.TransactOpts)
}

// IncreaseRelayNonce is a paid mutator transaction binding the contract method 0x90e690c7.
//
// Solidity: function increaseRelayNonce() returns()
func (_Wallet *WalletTransactorSession) IncreaseRelayNonce() (*types.Transaction, error) {
	return _Wallet.Contract.IncreaseRelayNonce(&_Wallet.TransactOpts)
}

// LoadTokenCard is a paid mutator transaction binding the contract method 0x3a43199f.
//
// Solidity: function loadTokenCard(address _asset, uint256 _amount) returns()
func (_Wallet *WalletTransactor) LoadTokenCard(opts *bind.TransactOpts, _asset common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "loadTokenCard", _asset, _amount)
}

// LoadTokenCard is a paid mutator transaction binding the contract method 0x3a43199f.
//
// Solidity: function loadTokenCard(address _asset, uint256 _amount) returns()
func (_Wallet *WalletSession) LoadTokenCard(_asset common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Wallet.Contract.LoadTokenCard(&_Wallet.TransactOpts, _asset, _amount)
}

// LoadTokenCard is a paid mutator transaction binding the contract method 0x3a43199f.
//
// Solidity: function loadTokenCard(address _asset, uint256 _amount) returns()
func (_Wallet *WalletTransactorSession) LoadTokenCard(_asset common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Wallet.Contract.LoadTokenCard(&_Wallet.TransactOpts, _asset, _amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Wallet *WalletTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Wallet *WalletSession) RenounceOwnership() (*types.Transaction, error) {
	return _Wallet.Contract.RenounceOwnership(&_Wallet.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Wallet *WalletTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Wallet.Contract.RenounceOwnership(&_Wallet.TransactOpts)
}

// SetGasTopUpLimit is a paid mutator transaction binding the contract method 0x0f3a85d8.
//
// Solidity: function setGasTopUpLimit(uint256 _amount) returns()
func (_Wallet *WalletTransactor) SetGasTopUpLimit(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "setGasTopUpLimit", _amount)
}

// SetGasTopUpLimit is a paid mutator transaction binding the contract method 0x0f3a85d8.
//
// Solidity: function setGasTopUpLimit(uint256 _amount) returns()
func (_Wallet *WalletSession) SetGasTopUpLimit(_amount *big.Int) (*types.Transaction, error) {
	return _Wallet.Contract.SetGasTopUpLimit(&_Wallet.TransactOpts, _amount)
}

// SetGasTopUpLimit is a paid mutator transaction binding the contract method 0x0f3a85d8.
//
// Solidity: function setGasTopUpLimit(uint256 _amount) returns()
func (_Wallet *WalletTransactorSession) SetGasTopUpLimit(_amount *big.Int) (*types.Transaction, error) {
	return _Wallet.Contract.SetGasTopUpLimit(&_Wallet.TransactOpts, _amount)
}

// SetLoadLimit is a paid mutator transaction binding the contract method 0x3bfec254.
//
// Solidity: function setLoadLimit(uint256 _amount) returns()
func (_Wallet *WalletTransactor) SetLoadLimit(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "setLoadLimit", _amount)
}

// SetLoadLimit is a paid mutator transaction binding the contract method 0x3bfec254.
//
// Solidity: function setLoadLimit(uint256 _amount) returns()
func (_Wallet *WalletSession) SetLoadLimit(_amount *big.Int) (*types.Transaction, error) {
	return _Wallet.Contract.SetLoadLimit(&_Wallet.TransactOpts, _amount)
}

// SetLoadLimit is a paid mutator transaction binding the contract method 0x3bfec254.
//
// Solidity: function setLoadLimit(uint256 _amount) returns()
func (_Wallet *WalletTransactorSession) SetLoadLimit(_amount *big.Int) (*types.Transaction, error) {
	return _Wallet.Contract.SetLoadLimit(&_Wallet.TransactOpts, _amount)
}

// SetSpendLimit is a paid mutator transaction binding the contract method 0x3c672eb7.
//
// Solidity: function setSpendLimit(uint256 _amount) returns()
func (_Wallet *WalletTransactor) SetSpendLimit(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "setSpendLimit", _amount)
}

// SetSpendLimit is a paid mutator transaction binding the contract method 0x3c672eb7.
//
// Solidity: function setSpendLimit(uint256 _amount) returns()
func (_Wallet *WalletSession) SetSpendLimit(_amount *big.Int) (*types.Transaction, error) {
	return _Wallet.Contract.SetSpendLimit(&_Wallet.TransactOpts, _amount)
}

// SetSpendLimit is a paid mutator transaction binding the contract method 0x3c672eb7.
//
// Solidity: function setSpendLimit(uint256 _amount) returns()
func (_Wallet *WalletTransactorSession) SetSpendLimit(_amount *big.Int) (*types.Transaction, error) {
	return _Wallet.Contract.SetSpendLimit(&_Wallet.TransactOpts, _amount)
}

// SetWhitelist is a paid mutator transaction binding the contract method 0xf4217648.
//
// Solidity: function setWhitelist(address[] _addresses) returns()
func (_Wallet *WalletTransactor) SetWhitelist(opts *bind.TransactOpts, _addresses []common.Address) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "setWhitelist", _addresses)
}

// SetWhitelist is a paid mutator transaction binding the contract method 0xf4217648.
//
// Solidity: function setWhitelist(address[] _addresses) returns()
func (_Wallet *WalletSession) SetWhitelist(_addresses []common.Address) (*types.Transaction, error) {
	return _Wallet.Contract.SetWhitelist(&_Wallet.TransactOpts, _addresses)
}

// SetWhitelist is a paid mutator transaction binding the contract method 0xf4217648.
//
// Solidity: function setWhitelist(address[] _addresses) returns()
func (_Wallet *WalletTransactorSession) SetWhitelist(_addresses []common.Address) (*types.Transaction, error) {
	return _Wallet.Contract.SetWhitelist(&_Wallet.TransactOpts, _addresses)
}

// SubmitGasTopUpLimitUpdate is a paid mutator transaction binding the contract method 0x74624c55.
//
// Solidity: function submitGasTopUpLimitUpdate(uint256 _amount) returns()
func (_Wallet *WalletTransactor) SubmitGasTopUpLimitUpdate(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "submitGasTopUpLimitUpdate", _amount)
}

// SubmitGasTopUpLimitUpdate is a paid mutator transaction binding the contract method 0x74624c55.
//
// Solidity: function submitGasTopUpLimitUpdate(uint256 _amount) returns()
func (_Wallet *WalletSession) SubmitGasTopUpLimitUpdate(_amount *big.Int) (*types.Transaction, error) {
	return _Wallet.Contract.SubmitGasTopUpLimitUpdate(&_Wallet.TransactOpts, _amount)
}

// SubmitGasTopUpLimitUpdate is a paid mutator transaction binding the contract method 0x74624c55.
//
// Solidity: function submitGasTopUpLimitUpdate(uint256 _amount) returns()
func (_Wallet *WalletTransactorSession) SubmitGasTopUpLimitUpdate(_amount *big.Int) (*types.Transaction, error) {
	return _Wallet.Contract.SubmitGasTopUpLimitUpdate(&_Wallet.TransactOpts, _amount)
}

// SubmitLoadLimitUpdate is a paid mutator transaction binding the contract method 0xb221f316.
//
// Solidity: function submitLoadLimitUpdate(uint256 _amount) returns()
func (_Wallet *WalletTransactor) SubmitLoadLimitUpdate(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "submitLoadLimitUpdate", _amount)
}

// SubmitLoadLimitUpdate is a paid mutator transaction binding the contract method 0xb221f316.
//
// Solidity: function submitLoadLimitUpdate(uint256 _amount) returns()
func (_Wallet *WalletSession) SubmitLoadLimitUpdate(_amount *big.Int) (*types.Transaction, error) {
	return _Wallet.Contract.SubmitLoadLimitUpdate(&_Wallet.TransactOpts, _amount)
}

// SubmitLoadLimitUpdate is a paid mutator transaction binding the contract method 0xb221f316.
//
// Solidity: function submitLoadLimitUpdate(uint256 _amount) returns()
func (_Wallet *WalletTransactorSession) SubmitLoadLimitUpdate(_amount *big.Int) (*types.Transaction, error) {
	return _Wallet.Contract.SubmitLoadLimitUpdate(&_Wallet.TransactOpts, _amount)
}

// SubmitSpendLimitUpdate is a paid mutator transaction binding the contract method 0x21ce918d.
//
// Solidity: function submitSpendLimitUpdate(uint256 _amount) returns()
func (_Wallet *WalletTransactor) SubmitSpendLimitUpdate(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "submitSpendLimitUpdate", _amount)
}

// SubmitSpendLimitUpdate is a paid mutator transaction binding the contract method 0x21ce918d.
//
// Solidity: function submitSpendLimitUpdate(uint256 _amount) returns()
func (_Wallet *WalletSession) SubmitSpendLimitUpdate(_amount *big.Int) (*types.Transaction, error) {
	return _Wallet.Contract.SubmitSpendLimitUpdate(&_Wallet.TransactOpts, _amount)
}

// SubmitSpendLimitUpdate is a paid mutator transaction binding the contract method 0x21ce918d.
//
// Solidity: function submitSpendLimitUpdate(uint256 _amount) returns()
func (_Wallet *WalletTransactorSession) SubmitSpendLimitUpdate(_amount *big.Int) (*types.Transaction, error) {
	return _Wallet.Contract.SubmitSpendLimitUpdate(&_Wallet.TransactOpts, _amount)
}

// SubmitWhitelistAddition is a paid mutator transaction binding the contract method 0x7fd004fa.
//
// Solidity: function submitWhitelistAddition(address[] _addresses) returns()
func (_Wallet *WalletTransactor) SubmitWhitelistAddition(opts *bind.TransactOpts, _addresses []common.Address) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "submitWhitelistAddition", _addresses)
}

// SubmitWhitelistAddition is a paid mutator transaction binding the contract method 0x7fd004fa.
//
// Solidity: function submitWhitelistAddition(address[] _addresses) returns()
func (_Wallet *WalletSession) SubmitWhitelistAddition(_addresses []common.Address) (*types.Transaction, error) {
	return _Wallet.Contract.SubmitWhitelistAddition(&_Wallet.TransactOpts, _addresses)
}

// SubmitWhitelistAddition is a paid mutator transaction binding the contract method 0x7fd004fa.
//
// Solidity: function submitWhitelistAddition(address[] _addresses) returns()
func (_Wallet *WalletTransactorSession) SubmitWhitelistAddition(_addresses []common.Address) (*types.Transaction, error) {
	return _Wallet.Contract.SubmitWhitelistAddition(&_Wallet.TransactOpts, _addresses)
}

// SubmitWhitelistRemoval is a paid mutator transaction binding the contract method 0x6137d670.
//
// Solidity: function submitWhitelistRemoval(address[] _addresses) returns()
func (_Wallet *WalletTransactor) SubmitWhitelistRemoval(opts *bind.TransactOpts, _addresses []common.Address) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "submitWhitelistRemoval", _addresses)
}

// SubmitWhitelistRemoval is a paid mutator transaction binding the contract method 0x6137d670.
//
// Solidity: function submitWhitelistRemoval(address[] _addresses) returns()
func (_Wallet *WalletSession) SubmitWhitelistRemoval(_addresses []common.Address) (*types.Transaction, error) {
	return _Wallet.Contract.SubmitWhitelistRemoval(&_Wallet.TransactOpts, _addresses)
}

// SubmitWhitelistRemoval is a paid mutator transaction binding the contract method 0x6137d670.
//
// Solidity: function submitWhitelistRemoval(address[] _addresses) returns()
func (_Wallet *WalletTransactorSession) SubmitWhitelistRemoval(_addresses []common.Address) (*types.Transaction, error) {
	return _Wallet.Contract.SubmitWhitelistRemoval(&_Wallet.TransactOpts, _addresses)
}

// TopUpGas is a paid mutator transaction binding the contract method 0xe61c51ca.
//
// Solidity: function topUpGas(uint256 _amount) returns()
func (_Wallet *WalletTransactor) TopUpGas(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "topUpGas", _amount)
}

// TopUpGas is a paid mutator transaction binding the contract method 0xe61c51ca.
//
// Solidity: function topUpGas(uint256 _amount) returns()
func (_Wallet *WalletSession) TopUpGas(_amount *big.Int) (*types.Transaction, error) {
	return _Wallet.Contract.TopUpGas(&_Wallet.TransactOpts, _amount)
}

// TopUpGas is a paid mutator transaction binding the contract method 0xe61c51ca.
//
// Solidity: function topUpGas(uint256 _amount) returns()
func (_Wallet *WalletTransactorSession) TopUpGas(_amount *big.Int) (*types.Transaction, error) {
	return _Wallet.Contract.TopUpGas(&_Wallet.TransactOpts, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xbeabacc8.
//
// Solidity: function transfer(address _to, address _asset, uint256 _amount) returns()
func (_Wallet *WalletTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _asset common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "transfer", _to, _asset, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xbeabacc8.
//
// Solidity: function transfer(address _to, address _asset, uint256 _amount) returns()
func (_Wallet *WalletSession) Transfer(_to common.Address, _asset common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Wallet.Contract.Transfer(&_Wallet.TransactOpts, _to, _asset, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xbeabacc8.
//
// Solidity: function transfer(address _to, address _asset, uint256 _amount) returns()
func (_Wallet *WalletTransactorSession) Transfer(_to common.Address, _asset common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Wallet.Contract.Transfer(&_Wallet.TransactOpts, _to, _asset, _amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xb242e534.
//
// Solidity: function transferOwnership(address _account, bool _transferable) returns()
func (_Wallet *WalletTransactor) TransferOwnership(opts *bind.TransactOpts, _account common.Address, _transferable bool) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "transferOwnership", _account, _transferable)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xb242e534.
//
// Solidity: function transferOwnership(address _account, bool _transferable) returns()
func (_Wallet *WalletSession) TransferOwnership(_account common.Address, _transferable bool) (*types.Transaction, error) {
	return _Wallet.Contract.TransferOwnership(&_Wallet.TransactOpts, _account, _transferable)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xb242e534.
//
// Solidity: function transferOwnership(address _account, bool _transferable) returns()
func (_Wallet *WalletTransactorSession) TransferOwnership(_account common.Address, _transferable bool) (*types.Transaction, error) {
	return _Wallet.Contract.TransferOwnership(&_Wallet.TransactOpts, _account, _transferable)
}

// WalletAddedToWhitelistIterator is returned from FilterAddedToWhitelist and is used to iterate over the raw logs and unpacked data for AddedToWhitelist events raised by the Wallet contract.
type WalletAddedToWhitelistIterator struct {
	Event *WalletAddedToWhitelist // Event containing the contract specifics and raw log

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
func (it *WalletAddedToWhitelistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletAddedToWhitelist)
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
		it.Event = new(WalletAddedToWhitelist)
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
func (it *WalletAddedToWhitelistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletAddedToWhitelistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletAddedToWhitelist represents a AddedToWhitelist event raised by the Wallet contract.
type WalletAddedToWhitelist struct {
	Sender    common.Address
	Addresses []common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAddedToWhitelist is a free log retrieval operation binding the contract event 0xb2f6cccee7a369e23e293c25aa19bef80af11eb26deba3ea0f2a02783f752e4a.
//
// Solidity: event AddedToWhitelist(address _sender, address[] _addresses)
func (_Wallet *WalletFilterer) FilterAddedToWhitelist(opts *bind.FilterOpts) (*WalletAddedToWhitelistIterator, error) {

	logs, sub, err := _Wallet.contract.FilterLogs(opts, "AddedToWhitelist")
	if err != nil {
		return nil, err
	}
	return &WalletAddedToWhitelistIterator{contract: _Wallet.contract, event: "AddedToWhitelist", logs: logs, sub: sub}, nil
}

// WatchAddedToWhitelist is a free log subscription operation binding the contract event 0xb2f6cccee7a369e23e293c25aa19bef80af11eb26deba3ea0f2a02783f752e4a.
//
// Solidity: event AddedToWhitelist(address _sender, address[] _addresses)
func (_Wallet *WalletFilterer) WatchAddedToWhitelist(opts *bind.WatchOpts, sink chan<- *WalletAddedToWhitelist) (event.Subscription, error) {

	logs, sub, err := _Wallet.contract.WatchLogs(opts, "AddedToWhitelist")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletAddedToWhitelist)
				if err := _Wallet.contract.UnpackLog(event, "AddedToWhitelist", log); err != nil {
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

// ParseAddedToWhitelist is a log parse operation binding the contract event 0xb2f6cccee7a369e23e293c25aa19bef80af11eb26deba3ea0f2a02783f752e4a.
//
// Solidity: event AddedToWhitelist(address _sender, address[] _addresses)
func (_Wallet *WalletFilterer) ParseAddedToWhitelist(log types.Log) (*WalletAddedToWhitelist, error) {
	event := new(WalletAddedToWhitelist)
	if err := _Wallet.contract.UnpackLog(event, "AddedToWhitelist", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WalletBulkTransferredIterator is returned from FilterBulkTransferred and is used to iterate over the raw logs and unpacked data for BulkTransferred events raised by the Wallet contract.
type WalletBulkTransferredIterator struct {
	Event *WalletBulkTransferred // Event containing the contract specifics and raw log

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
func (it *WalletBulkTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletBulkTransferred)
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
		it.Event = new(WalletBulkTransferred)
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
func (it *WalletBulkTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletBulkTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletBulkTransferred represents a BulkTransferred event raised by the Wallet contract.
type WalletBulkTransferred struct {
	To     common.Address
	Assets []common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBulkTransferred is a free log retrieval operation binding the contract event 0xd4f62f23021706247dcffea245d104ae7ddaec7f23acf3d11d7136d5de6a69ad.
//
// Solidity: event BulkTransferred(address _to, address[] _assets)
func (_Wallet *WalletFilterer) FilterBulkTransferred(opts *bind.FilterOpts) (*WalletBulkTransferredIterator, error) {

	logs, sub, err := _Wallet.contract.FilterLogs(opts, "BulkTransferred")
	if err != nil {
		return nil, err
	}
	return &WalletBulkTransferredIterator{contract: _Wallet.contract, event: "BulkTransferred", logs: logs, sub: sub}, nil
}

// WatchBulkTransferred is a free log subscription operation binding the contract event 0xd4f62f23021706247dcffea245d104ae7ddaec7f23acf3d11d7136d5de6a69ad.
//
// Solidity: event BulkTransferred(address _to, address[] _assets)
func (_Wallet *WalletFilterer) WatchBulkTransferred(opts *bind.WatchOpts, sink chan<- *WalletBulkTransferred) (event.Subscription, error) {

	logs, sub, err := _Wallet.contract.WatchLogs(opts, "BulkTransferred")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletBulkTransferred)
				if err := _Wallet.contract.UnpackLog(event, "BulkTransferred", log); err != nil {
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

// ParseBulkTransferred is a log parse operation binding the contract event 0xd4f62f23021706247dcffea245d104ae7ddaec7f23acf3d11d7136d5de6a69ad.
//
// Solidity: event BulkTransferred(address _to, address[] _assets)
func (_Wallet *WalletFilterer) ParseBulkTransferred(log types.Log) (*WalletBulkTransferred, error) {
	event := new(WalletBulkTransferred)
	if err := _Wallet.contract.UnpackLog(event, "BulkTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WalletCancelledWhitelistAdditionIterator is returned from FilterCancelledWhitelistAddition and is used to iterate over the raw logs and unpacked data for CancelledWhitelistAddition events raised by the Wallet contract.
type WalletCancelledWhitelistAdditionIterator struct {
	Event *WalletCancelledWhitelistAddition // Event containing the contract specifics and raw log

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
func (it *WalletCancelledWhitelistAdditionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletCancelledWhitelistAddition)
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
		it.Event = new(WalletCancelledWhitelistAddition)
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
func (it *WalletCancelledWhitelistAdditionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletCancelledWhitelistAdditionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletCancelledWhitelistAddition represents a CancelledWhitelistAddition event raised by the Wallet contract.
type WalletCancelledWhitelistAddition struct {
	Sender common.Address
	Hash   [32]byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCancelledWhitelistAddition is a free log retrieval operation binding the contract event 0x7794eff834d760583543e6e510e717a5e66d2c064e225f4db448343c3e66afcf.
//
// Solidity: event CancelledWhitelistAddition(address _sender, bytes32 _hash)
func (_Wallet *WalletFilterer) FilterCancelledWhitelistAddition(opts *bind.FilterOpts) (*WalletCancelledWhitelistAdditionIterator, error) {

	logs, sub, err := _Wallet.contract.FilterLogs(opts, "CancelledWhitelistAddition")
	if err != nil {
		return nil, err
	}
	return &WalletCancelledWhitelistAdditionIterator{contract: _Wallet.contract, event: "CancelledWhitelistAddition", logs: logs, sub: sub}, nil
}

// WatchCancelledWhitelistAddition is a free log subscription operation binding the contract event 0x7794eff834d760583543e6e510e717a5e66d2c064e225f4db448343c3e66afcf.
//
// Solidity: event CancelledWhitelistAddition(address _sender, bytes32 _hash)
func (_Wallet *WalletFilterer) WatchCancelledWhitelistAddition(opts *bind.WatchOpts, sink chan<- *WalletCancelledWhitelistAddition) (event.Subscription, error) {

	logs, sub, err := _Wallet.contract.WatchLogs(opts, "CancelledWhitelistAddition")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletCancelledWhitelistAddition)
				if err := _Wallet.contract.UnpackLog(event, "CancelledWhitelistAddition", log); err != nil {
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

// ParseCancelledWhitelistAddition is a log parse operation binding the contract event 0x7794eff834d760583543e6e510e717a5e66d2c064e225f4db448343c3e66afcf.
//
// Solidity: event CancelledWhitelistAddition(address _sender, bytes32 _hash)
func (_Wallet *WalletFilterer) ParseCancelledWhitelistAddition(log types.Log) (*WalletCancelledWhitelistAddition, error) {
	event := new(WalletCancelledWhitelistAddition)
	if err := _Wallet.contract.UnpackLog(event, "CancelledWhitelistAddition", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WalletCancelledWhitelistRemovalIterator is returned from FilterCancelledWhitelistRemoval and is used to iterate over the raw logs and unpacked data for CancelledWhitelistRemoval events raised by the Wallet contract.
type WalletCancelledWhitelistRemovalIterator struct {
	Event *WalletCancelledWhitelistRemoval // Event containing the contract specifics and raw log

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
func (it *WalletCancelledWhitelistRemovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletCancelledWhitelistRemoval)
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
		it.Event = new(WalletCancelledWhitelistRemoval)
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
func (it *WalletCancelledWhitelistRemovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletCancelledWhitelistRemovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletCancelledWhitelistRemoval represents a CancelledWhitelistRemoval event raised by the Wallet contract.
type WalletCancelledWhitelistRemoval struct {
	Sender common.Address
	Hash   [32]byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCancelledWhitelistRemoval is a free log retrieval operation binding the contract event 0x13c935eb475aa0f6e931fece83e2ac44569ce2d53460d29a6dedab40b965c8a3.
//
// Solidity: event CancelledWhitelistRemoval(address _sender, bytes32 _hash)
func (_Wallet *WalletFilterer) FilterCancelledWhitelistRemoval(opts *bind.FilterOpts) (*WalletCancelledWhitelistRemovalIterator, error) {

	logs, sub, err := _Wallet.contract.FilterLogs(opts, "CancelledWhitelistRemoval")
	if err != nil {
		return nil, err
	}
	return &WalletCancelledWhitelistRemovalIterator{contract: _Wallet.contract, event: "CancelledWhitelistRemoval", logs: logs, sub: sub}, nil
}

// WatchCancelledWhitelistRemoval is a free log subscription operation binding the contract event 0x13c935eb475aa0f6e931fece83e2ac44569ce2d53460d29a6dedab40b965c8a3.
//
// Solidity: event CancelledWhitelistRemoval(address _sender, bytes32 _hash)
func (_Wallet *WalletFilterer) WatchCancelledWhitelistRemoval(opts *bind.WatchOpts, sink chan<- *WalletCancelledWhitelistRemoval) (event.Subscription, error) {

	logs, sub, err := _Wallet.contract.WatchLogs(opts, "CancelledWhitelistRemoval")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletCancelledWhitelistRemoval)
				if err := _Wallet.contract.UnpackLog(event, "CancelledWhitelistRemoval", log); err != nil {
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

// ParseCancelledWhitelistRemoval is a log parse operation binding the contract event 0x13c935eb475aa0f6e931fece83e2ac44569ce2d53460d29a6dedab40b965c8a3.
//
// Solidity: event CancelledWhitelistRemoval(address _sender, bytes32 _hash)
func (_Wallet *WalletFilterer) ParseCancelledWhitelistRemoval(log types.Log) (*WalletCancelledWhitelistRemoval, error) {
	event := new(WalletCancelledWhitelistRemoval)
	if err := _Wallet.contract.UnpackLog(event, "CancelledWhitelistRemoval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WalletExecutedRelayedTransactionIterator is returned from FilterExecutedRelayedTransaction and is used to iterate over the raw logs and unpacked data for ExecutedRelayedTransaction events raised by the Wallet contract.
type WalletExecutedRelayedTransactionIterator struct {
	Event *WalletExecutedRelayedTransaction // Event containing the contract specifics and raw log

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
func (it *WalletExecutedRelayedTransactionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletExecutedRelayedTransaction)
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
		it.Event = new(WalletExecutedRelayedTransaction)
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
func (it *WalletExecutedRelayedTransactionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletExecutedRelayedTransactionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletExecutedRelayedTransaction represents a ExecutedRelayedTransaction event raised by the Wallet contract.
type WalletExecutedRelayedTransaction struct {
	Data       []byte
	Returndata []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterExecutedRelayedTransaction is a free log retrieval operation binding the contract event 0x823dbcf2b7b0f265871963ca65ac033f6b4c71e0d82cd123d2ff23d752dc21c1.
//
// Solidity: event ExecutedRelayedTransaction(bytes _data, bytes _returndata)
func (_Wallet *WalletFilterer) FilterExecutedRelayedTransaction(opts *bind.FilterOpts) (*WalletExecutedRelayedTransactionIterator, error) {

	logs, sub, err := _Wallet.contract.FilterLogs(opts, "ExecutedRelayedTransaction")
	if err != nil {
		return nil, err
	}
	return &WalletExecutedRelayedTransactionIterator{contract: _Wallet.contract, event: "ExecutedRelayedTransaction", logs: logs, sub: sub}, nil
}

// WatchExecutedRelayedTransaction is a free log subscription operation binding the contract event 0x823dbcf2b7b0f265871963ca65ac033f6b4c71e0d82cd123d2ff23d752dc21c1.
//
// Solidity: event ExecutedRelayedTransaction(bytes _data, bytes _returndata)
func (_Wallet *WalletFilterer) WatchExecutedRelayedTransaction(opts *bind.WatchOpts, sink chan<- *WalletExecutedRelayedTransaction) (event.Subscription, error) {

	logs, sub, err := _Wallet.contract.WatchLogs(opts, "ExecutedRelayedTransaction")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletExecutedRelayedTransaction)
				if err := _Wallet.contract.UnpackLog(event, "ExecutedRelayedTransaction", log); err != nil {
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

// ParseExecutedRelayedTransaction is a log parse operation binding the contract event 0x823dbcf2b7b0f265871963ca65ac033f6b4c71e0d82cd123d2ff23d752dc21c1.
//
// Solidity: event ExecutedRelayedTransaction(bytes _data, bytes _returndata)
func (_Wallet *WalletFilterer) ParseExecutedRelayedTransaction(log types.Log) (*WalletExecutedRelayedTransaction, error) {
	event := new(WalletExecutedRelayedTransaction)
	if err := _Wallet.contract.UnpackLog(event, "ExecutedRelayedTransaction", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WalletExecutedTransactionIterator is returned from FilterExecutedTransaction and is used to iterate over the raw logs and unpacked data for ExecutedTransaction events raised by the Wallet contract.
type WalletExecutedTransactionIterator struct {
	Event *WalletExecutedTransaction // Event containing the contract specifics and raw log

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
func (it *WalletExecutedTransactionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletExecutedTransaction)
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
		it.Event = new(WalletExecutedTransaction)
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
func (it *WalletExecutedTransactionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletExecutedTransactionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletExecutedTransaction represents a ExecutedTransaction event raised by the Wallet contract.
type WalletExecutedTransaction struct {
	Destination common.Address
	Value       *big.Int
	Data        []byte
	Returndata  []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterExecutedTransaction is a free log retrieval operation binding the contract event 0xf77753fab406ecfff96d6ff2476c64a838fa9f6d37b1bf190f8546e395e3b613.
//
// Solidity: event ExecutedTransaction(address _destination, uint256 _value, bytes _data, bytes _returndata)
func (_Wallet *WalletFilterer) FilterExecutedTransaction(opts *bind.FilterOpts) (*WalletExecutedTransactionIterator, error) {

	logs, sub, err := _Wallet.contract.FilterLogs(opts, "ExecutedTransaction")
	if err != nil {
		return nil, err
	}
	return &WalletExecutedTransactionIterator{contract: _Wallet.contract, event: "ExecutedTransaction", logs: logs, sub: sub}, nil
}

// WatchExecutedTransaction is a free log subscription operation binding the contract event 0xf77753fab406ecfff96d6ff2476c64a838fa9f6d37b1bf190f8546e395e3b613.
//
// Solidity: event ExecutedTransaction(address _destination, uint256 _value, bytes _data, bytes _returndata)
func (_Wallet *WalletFilterer) WatchExecutedTransaction(opts *bind.WatchOpts, sink chan<- *WalletExecutedTransaction) (event.Subscription, error) {

	logs, sub, err := _Wallet.contract.WatchLogs(opts, "ExecutedTransaction")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletExecutedTransaction)
				if err := _Wallet.contract.UnpackLog(event, "ExecutedTransaction", log); err != nil {
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
// Solidity: event ExecutedTransaction(address _destination, uint256 _value, bytes _data, bytes _returndata)
func (_Wallet *WalletFilterer) ParseExecutedTransaction(log types.Log) (*WalletExecutedTransaction, error) {
	event := new(WalletExecutedTransaction)
	if err := _Wallet.contract.UnpackLog(event, "ExecutedTransaction", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WalletIncreasedRelayNonceIterator is returned from FilterIncreasedRelayNonce and is used to iterate over the raw logs and unpacked data for IncreasedRelayNonce events raised by the Wallet contract.
type WalletIncreasedRelayNonceIterator struct {
	Event *WalletIncreasedRelayNonce // Event containing the contract specifics and raw log

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
func (it *WalletIncreasedRelayNonceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletIncreasedRelayNonce)
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
		it.Event = new(WalletIncreasedRelayNonce)
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
func (it *WalletIncreasedRelayNonceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletIncreasedRelayNonceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletIncreasedRelayNonce represents a IncreasedRelayNonce event raised by the Wallet contract.
type WalletIncreasedRelayNonce struct {
	Sender       common.Address
	CurrentNonce *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterIncreasedRelayNonce is a free log retrieval operation binding the contract event 0xab0423a75986556234aecd171c46ce7f5e45607d8070bf5230f2735b50322bff.
//
// Solidity: event IncreasedRelayNonce(address _sender, uint256 _currentNonce)
func (_Wallet *WalletFilterer) FilterIncreasedRelayNonce(opts *bind.FilterOpts) (*WalletIncreasedRelayNonceIterator, error) {

	logs, sub, err := _Wallet.contract.FilterLogs(opts, "IncreasedRelayNonce")
	if err != nil {
		return nil, err
	}
	return &WalletIncreasedRelayNonceIterator{contract: _Wallet.contract, event: "IncreasedRelayNonce", logs: logs, sub: sub}, nil
}

// WatchIncreasedRelayNonce is a free log subscription operation binding the contract event 0xab0423a75986556234aecd171c46ce7f5e45607d8070bf5230f2735b50322bff.
//
// Solidity: event IncreasedRelayNonce(address _sender, uint256 _currentNonce)
func (_Wallet *WalletFilterer) WatchIncreasedRelayNonce(opts *bind.WatchOpts, sink chan<- *WalletIncreasedRelayNonce) (event.Subscription, error) {

	logs, sub, err := _Wallet.contract.WatchLogs(opts, "IncreasedRelayNonce")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletIncreasedRelayNonce)
				if err := _Wallet.contract.UnpackLog(event, "IncreasedRelayNonce", log); err != nil {
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

// ParseIncreasedRelayNonce is a log parse operation binding the contract event 0xab0423a75986556234aecd171c46ce7f5e45607d8070bf5230f2735b50322bff.
//
// Solidity: event IncreasedRelayNonce(address _sender, uint256 _currentNonce)
func (_Wallet *WalletFilterer) ParseIncreasedRelayNonce(log types.Log) (*WalletIncreasedRelayNonce, error) {
	event := new(WalletIncreasedRelayNonce)
	if err := _Wallet.contract.UnpackLog(event, "IncreasedRelayNonce", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WalletLoadedTokenCardIterator is returned from FilterLoadedTokenCard and is used to iterate over the raw logs and unpacked data for LoadedTokenCard events raised by the Wallet contract.
type WalletLoadedTokenCardIterator struct {
	Event *WalletLoadedTokenCard // Event containing the contract specifics and raw log

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
func (it *WalletLoadedTokenCardIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletLoadedTokenCard)
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
		it.Event = new(WalletLoadedTokenCard)
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
func (it *WalletLoadedTokenCardIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletLoadedTokenCardIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletLoadedTokenCard represents a LoadedTokenCard event raised by the Wallet contract.
type WalletLoadedTokenCard struct {
	Asset  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLoadedTokenCard is a free log retrieval operation binding the contract event 0x5f65674bec9af81f71be68674135a0ea3f163fb91984e3893d06da9f6ea2ce8a.
//
// Solidity: event LoadedTokenCard(address _asset, uint256 _amount)
func (_Wallet *WalletFilterer) FilterLoadedTokenCard(opts *bind.FilterOpts) (*WalletLoadedTokenCardIterator, error) {

	logs, sub, err := _Wallet.contract.FilterLogs(opts, "LoadedTokenCard")
	if err != nil {
		return nil, err
	}
	return &WalletLoadedTokenCardIterator{contract: _Wallet.contract, event: "LoadedTokenCard", logs: logs, sub: sub}, nil
}

// WatchLoadedTokenCard is a free log subscription operation binding the contract event 0x5f65674bec9af81f71be68674135a0ea3f163fb91984e3893d06da9f6ea2ce8a.
//
// Solidity: event LoadedTokenCard(address _asset, uint256 _amount)
func (_Wallet *WalletFilterer) WatchLoadedTokenCard(opts *bind.WatchOpts, sink chan<- *WalletLoadedTokenCard) (event.Subscription, error) {

	logs, sub, err := _Wallet.contract.WatchLogs(opts, "LoadedTokenCard")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletLoadedTokenCard)
				if err := _Wallet.contract.UnpackLog(event, "LoadedTokenCard", log); err != nil {
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

// ParseLoadedTokenCard is a log parse operation binding the contract event 0x5f65674bec9af81f71be68674135a0ea3f163fb91984e3893d06da9f6ea2ce8a.
//
// Solidity: event LoadedTokenCard(address _asset, uint256 _amount)
func (_Wallet *WalletFilterer) ParseLoadedTokenCard(log types.Log) (*WalletLoadedTokenCard, error) {
	event := new(WalletLoadedTokenCard)
	if err := _Wallet.contract.UnpackLog(event, "LoadedTokenCard", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WalletLockedOwnershipIterator is returned from FilterLockedOwnership and is used to iterate over the raw logs and unpacked data for LockedOwnership events raised by the Wallet contract.
type WalletLockedOwnershipIterator struct {
	Event *WalletLockedOwnership // Event containing the contract specifics and raw log

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
func (it *WalletLockedOwnershipIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletLockedOwnership)
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
		it.Event = new(WalletLockedOwnership)
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
func (it *WalletLockedOwnershipIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletLockedOwnershipIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletLockedOwnership represents a LockedOwnership event raised by the Wallet contract.
type WalletLockedOwnership struct {
	Locked common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLockedOwnership is a free log retrieval operation binding the contract event 0x808639ff9c8e4732d60b6c2330de498035416d229f27a77d259680895efec122.
//
// Solidity: event LockedOwnership(address _locked)
func (_Wallet *WalletFilterer) FilterLockedOwnership(opts *bind.FilterOpts) (*WalletLockedOwnershipIterator, error) {

	logs, sub, err := _Wallet.contract.FilterLogs(opts, "LockedOwnership")
	if err != nil {
		return nil, err
	}
	return &WalletLockedOwnershipIterator{contract: _Wallet.contract, event: "LockedOwnership", logs: logs, sub: sub}, nil
}

// WatchLockedOwnership is a free log subscription operation binding the contract event 0x808639ff9c8e4732d60b6c2330de498035416d229f27a77d259680895efec122.
//
// Solidity: event LockedOwnership(address _locked)
func (_Wallet *WalletFilterer) WatchLockedOwnership(opts *bind.WatchOpts, sink chan<- *WalletLockedOwnership) (event.Subscription, error) {

	logs, sub, err := _Wallet.contract.WatchLogs(opts, "LockedOwnership")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletLockedOwnership)
				if err := _Wallet.contract.UnpackLog(event, "LockedOwnership", log); err != nil {
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

// ParseLockedOwnership is a log parse operation binding the contract event 0x808639ff9c8e4732d60b6c2330de498035416d229f27a77d259680895efec122.
//
// Solidity: event LockedOwnership(address _locked)
func (_Wallet *WalletFilterer) ParseLockedOwnership(log types.Log) (*WalletLockedOwnership, error) {
	event := new(WalletLockedOwnership)
	if err := _Wallet.contract.UnpackLog(event, "LockedOwnership", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WalletReceivedIterator is returned from FilterReceived and is used to iterate over the raw logs and unpacked data for Received events raised by the Wallet contract.
type WalletReceivedIterator struct {
	Event *WalletReceived // Event containing the contract specifics and raw log

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
func (it *WalletReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletReceived)
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
		it.Event = new(WalletReceived)
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
func (it *WalletReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletReceived represents a Received event raised by the Wallet contract.
type WalletReceived struct {
	From   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterReceived is a free log retrieval operation binding the contract event 0x88a5966d370b9919b20f3e2c13ff65706f196a4e32cc2c12bf57088f88525874.
//
// Solidity: event Received(address _from, uint256 _amount)
func (_Wallet *WalletFilterer) FilterReceived(opts *bind.FilterOpts) (*WalletReceivedIterator, error) {

	logs, sub, err := _Wallet.contract.FilterLogs(opts, "Received")
	if err != nil {
		return nil, err
	}
	return &WalletReceivedIterator{contract: _Wallet.contract, event: "Received", logs: logs, sub: sub}, nil
}

// WatchReceived is a free log subscription operation binding the contract event 0x88a5966d370b9919b20f3e2c13ff65706f196a4e32cc2c12bf57088f88525874.
//
// Solidity: event Received(address _from, uint256 _amount)
func (_Wallet *WalletFilterer) WatchReceived(opts *bind.WatchOpts, sink chan<- *WalletReceived) (event.Subscription, error) {

	logs, sub, err := _Wallet.contract.WatchLogs(opts, "Received")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletReceived)
				if err := _Wallet.contract.UnpackLog(event, "Received", log); err != nil {
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
func (_Wallet *WalletFilterer) ParseReceived(log types.Log) (*WalletReceived, error) {
	event := new(WalletReceived)
	if err := _Wallet.contract.UnpackLog(event, "Received", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WalletRemovedFromWhitelistIterator is returned from FilterRemovedFromWhitelist and is used to iterate over the raw logs and unpacked data for RemovedFromWhitelist events raised by the Wallet contract.
type WalletRemovedFromWhitelistIterator struct {
	Event *WalletRemovedFromWhitelist // Event containing the contract specifics and raw log

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
func (it *WalletRemovedFromWhitelistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletRemovedFromWhitelist)
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
		it.Event = new(WalletRemovedFromWhitelist)
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
func (it *WalletRemovedFromWhitelistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletRemovedFromWhitelistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletRemovedFromWhitelist represents a RemovedFromWhitelist event raised by the Wallet contract.
type WalletRemovedFromWhitelist struct {
	Sender    common.Address
	Addresses []common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRemovedFromWhitelist is a free log retrieval operation binding the contract event 0xd218c430fa348f4ce67791021b6b89c0c3eacd4ead1d8f5b83c60038ec28249b.
//
// Solidity: event RemovedFromWhitelist(address _sender, address[] _addresses)
func (_Wallet *WalletFilterer) FilterRemovedFromWhitelist(opts *bind.FilterOpts) (*WalletRemovedFromWhitelistIterator, error) {

	logs, sub, err := _Wallet.contract.FilterLogs(opts, "RemovedFromWhitelist")
	if err != nil {
		return nil, err
	}
	return &WalletRemovedFromWhitelistIterator{contract: _Wallet.contract, event: "RemovedFromWhitelist", logs: logs, sub: sub}, nil
}

// WatchRemovedFromWhitelist is a free log subscription operation binding the contract event 0xd218c430fa348f4ce67791021b6b89c0c3eacd4ead1d8f5b83c60038ec28249b.
//
// Solidity: event RemovedFromWhitelist(address _sender, address[] _addresses)
func (_Wallet *WalletFilterer) WatchRemovedFromWhitelist(opts *bind.WatchOpts, sink chan<- *WalletRemovedFromWhitelist) (event.Subscription, error) {

	logs, sub, err := _Wallet.contract.WatchLogs(opts, "RemovedFromWhitelist")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletRemovedFromWhitelist)
				if err := _Wallet.contract.UnpackLog(event, "RemovedFromWhitelist", log); err != nil {
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

// ParseRemovedFromWhitelist is a log parse operation binding the contract event 0xd218c430fa348f4ce67791021b6b89c0c3eacd4ead1d8f5b83c60038ec28249b.
//
// Solidity: event RemovedFromWhitelist(address _sender, address[] _addresses)
func (_Wallet *WalletFilterer) ParseRemovedFromWhitelist(log types.Log) (*WalletRemovedFromWhitelist, error) {
	event := new(WalletRemovedFromWhitelist)
	if err := _Wallet.contract.UnpackLog(event, "RemovedFromWhitelist", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WalletSetGasTopUpLimitIterator is returned from FilterSetGasTopUpLimit and is used to iterate over the raw logs and unpacked data for SetGasTopUpLimit events raised by the Wallet contract.
type WalletSetGasTopUpLimitIterator struct {
	Event *WalletSetGasTopUpLimit // Event containing the contract specifics and raw log

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
func (it *WalletSetGasTopUpLimitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletSetGasTopUpLimit)
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
		it.Event = new(WalletSetGasTopUpLimit)
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
func (it *WalletSetGasTopUpLimitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletSetGasTopUpLimitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletSetGasTopUpLimit represents a SetGasTopUpLimit event raised by the Wallet contract.
type WalletSetGasTopUpLimit struct {
	Sender common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetGasTopUpLimit is a free log retrieval operation binding the contract event 0x41ff5d5ce3b7935893a4e7269ec5caae9cca5e3bf0eb4b21d2f443489667112e.
//
// Solidity: event SetGasTopUpLimit(address _sender, uint256 _amount)
func (_Wallet *WalletFilterer) FilterSetGasTopUpLimit(opts *bind.FilterOpts) (*WalletSetGasTopUpLimitIterator, error) {

	logs, sub, err := _Wallet.contract.FilterLogs(opts, "SetGasTopUpLimit")
	if err != nil {
		return nil, err
	}
	return &WalletSetGasTopUpLimitIterator{contract: _Wallet.contract, event: "SetGasTopUpLimit", logs: logs, sub: sub}, nil
}

// WatchSetGasTopUpLimit is a free log subscription operation binding the contract event 0x41ff5d5ce3b7935893a4e7269ec5caae9cca5e3bf0eb4b21d2f443489667112e.
//
// Solidity: event SetGasTopUpLimit(address _sender, uint256 _amount)
func (_Wallet *WalletFilterer) WatchSetGasTopUpLimit(opts *bind.WatchOpts, sink chan<- *WalletSetGasTopUpLimit) (event.Subscription, error) {

	logs, sub, err := _Wallet.contract.WatchLogs(opts, "SetGasTopUpLimit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletSetGasTopUpLimit)
				if err := _Wallet.contract.UnpackLog(event, "SetGasTopUpLimit", log); err != nil {
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

// ParseSetGasTopUpLimit is a log parse operation binding the contract event 0x41ff5d5ce3b7935893a4e7269ec5caae9cca5e3bf0eb4b21d2f443489667112e.
//
// Solidity: event SetGasTopUpLimit(address _sender, uint256 _amount)
func (_Wallet *WalletFilterer) ParseSetGasTopUpLimit(log types.Log) (*WalletSetGasTopUpLimit, error) {
	event := new(WalletSetGasTopUpLimit)
	if err := _Wallet.contract.UnpackLog(event, "SetGasTopUpLimit", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WalletSetLoadLimitIterator is returned from FilterSetLoadLimit and is used to iterate over the raw logs and unpacked data for SetLoadLimit events raised by the Wallet contract.
type WalletSetLoadLimitIterator struct {
	Event *WalletSetLoadLimit // Event containing the contract specifics and raw log

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
func (it *WalletSetLoadLimitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletSetLoadLimit)
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
		it.Event = new(WalletSetLoadLimit)
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
func (it *WalletSetLoadLimitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletSetLoadLimitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletSetLoadLimit represents a SetLoadLimit event raised by the Wallet contract.
type WalletSetLoadLimit struct {
	Sender common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetLoadLimit is a free log retrieval operation binding the contract event 0x0b05243483e17c3f3377aee82b7d47e5700b48288695fc08b7ecc2759afa44ef.
//
// Solidity: event SetLoadLimit(address _sender, uint256 _amount)
func (_Wallet *WalletFilterer) FilterSetLoadLimit(opts *bind.FilterOpts) (*WalletSetLoadLimitIterator, error) {

	logs, sub, err := _Wallet.contract.FilterLogs(opts, "SetLoadLimit")
	if err != nil {
		return nil, err
	}
	return &WalletSetLoadLimitIterator{contract: _Wallet.contract, event: "SetLoadLimit", logs: logs, sub: sub}, nil
}

// WatchSetLoadLimit is a free log subscription operation binding the contract event 0x0b05243483e17c3f3377aee82b7d47e5700b48288695fc08b7ecc2759afa44ef.
//
// Solidity: event SetLoadLimit(address _sender, uint256 _amount)
func (_Wallet *WalletFilterer) WatchSetLoadLimit(opts *bind.WatchOpts, sink chan<- *WalletSetLoadLimit) (event.Subscription, error) {

	logs, sub, err := _Wallet.contract.WatchLogs(opts, "SetLoadLimit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletSetLoadLimit)
				if err := _Wallet.contract.UnpackLog(event, "SetLoadLimit", log); err != nil {
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

// ParseSetLoadLimit is a log parse operation binding the contract event 0x0b05243483e17c3f3377aee82b7d47e5700b48288695fc08b7ecc2759afa44ef.
//
// Solidity: event SetLoadLimit(address _sender, uint256 _amount)
func (_Wallet *WalletFilterer) ParseSetLoadLimit(log types.Log) (*WalletSetLoadLimit, error) {
	event := new(WalletSetLoadLimit)
	if err := _Wallet.contract.UnpackLog(event, "SetLoadLimit", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WalletSetSpendLimitIterator is returned from FilterSetSpendLimit and is used to iterate over the raw logs and unpacked data for SetSpendLimit events raised by the Wallet contract.
type WalletSetSpendLimitIterator struct {
	Event *WalletSetSpendLimit // Event containing the contract specifics and raw log

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
func (it *WalletSetSpendLimitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletSetSpendLimit)
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
		it.Event = new(WalletSetSpendLimit)
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
func (it *WalletSetSpendLimitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletSetSpendLimitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletSetSpendLimit represents a SetSpendLimit event raised by the Wallet contract.
type WalletSetSpendLimit struct {
	Sender common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetSpendLimit is a free log retrieval operation binding the contract event 0x068f112e5ec923d412be64779fe69e0fcbb6784c6617e94cccc8fd348f2e0f21.
//
// Solidity: event SetSpendLimit(address _sender, uint256 _amount)
func (_Wallet *WalletFilterer) FilterSetSpendLimit(opts *bind.FilterOpts) (*WalletSetSpendLimitIterator, error) {

	logs, sub, err := _Wallet.contract.FilterLogs(opts, "SetSpendLimit")
	if err != nil {
		return nil, err
	}
	return &WalletSetSpendLimitIterator{contract: _Wallet.contract, event: "SetSpendLimit", logs: logs, sub: sub}, nil
}

// WatchSetSpendLimit is a free log subscription operation binding the contract event 0x068f112e5ec923d412be64779fe69e0fcbb6784c6617e94cccc8fd348f2e0f21.
//
// Solidity: event SetSpendLimit(address _sender, uint256 _amount)
func (_Wallet *WalletFilterer) WatchSetSpendLimit(opts *bind.WatchOpts, sink chan<- *WalletSetSpendLimit) (event.Subscription, error) {

	logs, sub, err := _Wallet.contract.WatchLogs(opts, "SetSpendLimit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletSetSpendLimit)
				if err := _Wallet.contract.UnpackLog(event, "SetSpendLimit", log); err != nil {
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

// ParseSetSpendLimit is a log parse operation binding the contract event 0x068f112e5ec923d412be64779fe69e0fcbb6784c6617e94cccc8fd348f2e0f21.
//
// Solidity: event SetSpendLimit(address _sender, uint256 _amount)
func (_Wallet *WalletFilterer) ParseSetSpendLimit(log types.Log) (*WalletSetSpendLimit, error) {
	event := new(WalletSetSpendLimit)
	if err := _Wallet.contract.UnpackLog(event, "SetSpendLimit", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WalletSubmittedGasTopUpLimitUpdateIterator is returned from FilterSubmittedGasTopUpLimitUpdate and is used to iterate over the raw logs and unpacked data for SubmittedGasTopUpLimitUpdate events raised by the Wallet contract.
type WalletSubmittedGasTopUpLimitUpdateIterator struct {
	Event *WalletSubmittedGasTopUpLimitUpdate // Event containing the contract specifics and raw log

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
func (it *WalletSubmittedGasTopUpLimitUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletSubmittedGasTopUpLimitUpdate)
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
		it.Event = new(WalletSubmittedGasTopUpLimitUpdate)
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
func (it *WalletSubmittedGasTopUpLimitUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletSubmittedGasTopUpLimitUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletSubmittedGasTopUpLimitUpdate represents a SubmittedGasTopUpLimitUpdate event raised by the Wallet contract.
type WalletSubmittedGasTopUpLimitUpdate struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSubmittedGasTopUpLimitUpdate is a free log retrieval operation binding the contract event 0xaf2a77cd04c3cc155588dd3bf67b310ab4fb3b1da3cf6b8d7d4d2aa1d09b794c.
//
// Solidity: event SubmittedGasTopUpLimitUpdate(uint256 _amount)
func (_Wallet *WalletFilterer) FilterSubmittedGasTopUpLimitUpdate(opts *bind.FilterOpts) (*WalletSubmittedGasTopUpLimitUpdateIterator, error) {

	logs, sub, err := _Wallet.contract.FilterLogs(opts, "SubmittedGasTopUpLimitUpdate")
	if err != nil {
		return nil, err
	}
	return &WalletSubmittedGasTopUpLimitUpdateIterator{contract: _Wallet.contract, event: "SubmittedGasTopUpLimitUpdate", logs: logs, sub: sub}, nil
}

// WatchSubmittedGasTopUpLimitUpdate is a free log subscription operation binding the contract event 0xaf2a77cd04c3cc155588dd3bf67b310ab4fb3b1da3cf6b8d7d4d2aa1d09b794c.
//
// Solidity: event SubmittedGasTopUpLimitUpdate(uint256 _amount)
func (_Wallet *WalletFilterer) WatchSubmittedGasTopUpLimitUpdate(opts *bind.WatchOpts, sink chan<- *WalletSubmittedGasTopUpLimitUpdate) (event.Subscription, error) {

	logs, sub, err := _Wallet.contract.WatchLogs(opts, "SubmittedGasTopUpLimitUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletSubmittedGasTopUpLimitUpdate)
				if err := _Wallet.contract.UnpackLog(event, "SubmittedGasTopUpLimitUpdate", log); err != nil {
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

// ParseSubmittedGasTopUpLimitUpdate is a log parse operation binding the contract event 0xaf2a77cd04c3cc155588dd3bf67b310ab4fb3b1da3cf6b8d7d4d2aa1d09b794c.
//
// Solidity: event SubmittedGasTopUpLimitUpdate(uint256 _amount)
func (_Wallet *WalletFilterer) ParseSubmittedGasTopUpLimitUpdate(log types.Log) (*WalletSubmittedGasTopUpLimitUpdate, error) {
	event := new(WalletSubmittedGasTopUpLimitUpdate)
	if err := _Wallet.contract.UnpackLog(event, "SubmittedGasTopUpLimitUpdate", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WalletSubmittedLoadLimitUpdateIterator is returned from FilterSubmittedLoadLimitUpdate and is used to iterate over the raw logs and unpacked data for SubmittedLoadLimitUpdate events raised by the Wallet contract.
type WalletSubmittedLoadLimitUpdateIterator struct {
	Event *WalletSubmittedLoadLimitUpdate // Event containing the contract specifics and raw log

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
func (it *WalletSubmittedLoadLimitUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletSubmittedLoadLimitUpdate)
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
		it.Event = new(WalletSubmittedLoadLimitUpdate)
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
func (it *WalletSubmittedLoadLimitUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletSubmittedLoadLimitUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletSubmittedLoadLimitUpdate represents a SubmittedLoadLimitUpdate event raised by the Wallet contract.
type WalletSubmittedLoadLimitUpdate struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSubmittedLoadLimitUpdate is a free log retrieval operation binding the contract event 0xc178d379965e5657b6fc57494e392f121a14119215dfb422aad7db4cc03f2d10.
//
// Solidity: event SubmittedLoadLimitUpdate(uint256 _amount)
func (_Wallet *WalletFilterer) FilterSubmittedLoadLimitUpdate(opts *bind.FilterOpts) (*WalletSubmittedLoadLimitUpdateIterator, error) {

	logs, sub, err := _Wallet.contract.FilterLogs(opts, "SubmittedLoadLimitUpdate")
	if err != nil {
		return nil, err
	}
	return &WalletSubmittedLoadLimitUpdateIterator{contract: _Wallet.contract, event: "SubmittedLoadLimitUpdate", logs: logs, sub: sub}, nil
}

// WatchSubmittedLoadLimitUpdate is a free log subscription operation binding the contract event 0xc178d379965e5657b6fc57494e392f121a14119215dfb422aad7db4cc03f2d10.
//
// Solidity: event SubmittedLoadLimitUpdate(uint256 _amount)
func (_Wallet *WalletFilterer) WatchSubmittedLoadLimitUpdate(opts *bind.WatchOpts, sink chan<- *WalletSubmittedLoadLimitUpdate) (event.Subscription, error) {

	logs, sub, err := _Wallet.contract.WatchLogs(opts, "SubmittedLoadLimitUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletSubmittedLoadLimitUpdate)
				if err := _Wallet.contract.UnpackLog(event, "SubmittedLoadLimitUpdate", log); err != nil {
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

// ParseSubmittedLoadLimitUpdate is a log parse operation binding the contract event 0xc178d379965e5657b6fc57494e392f121a14119215dfb422aad7db4cc03f2d10.
//
// Solidity: event SubmittedLoadLimitUpdate(uint256 _amount)
func (_Wallet *WalletFilterer) ParseSubmittedLoadLimitUpdate(log types.Log) (*WalletSubmittedLoadLimitUpdate, error) {
	event := new(WalletSubmittedLoadLimitUpdate)
	if err := _Wallet.contract.UnpackLog(event, "SubmittedLoadLimitUpdate", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WalletSubmittedSpendLimitUpdateIterator is returned from FilterSubmittedSpendLimitUpdate and is used to iterate over the raw logs and unpacked data for SubmittedSpendLimitUpdate events raised by the Wallet contract.
type WalletSubmittedSpendLimitUpdateIterator struct {
	Event *WalletSubmittedSpendLimitUpdate // Event containing the contract specifics and raw log

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
func (it *WalletSubmittedSpendLimitUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletSubmittedSpendLimitUpdate)
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
		it.Event = new(WalletSubmittedSpendLimitUpdate)
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
func (it *WalletSubmittedSpendLimitUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletSubmittedSpendLimitUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletSubmittedSpendLimitUpdate represents a SubmittedSpendLimitUpdate event raised by the Wallet contract.
type WalletSubmittedSpendLimitUpdate struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSubmittedSpendLimitUpdate is a free log retrieval operation binding the contract event 0x4b1b970c8a0fa761e7803ed70c13d7aca71904b13df60fbe03f981da1730da91.
//
// Solidity: event SubmittedSpendLimitUpdate(uint256 _amount)
func (_Wallet *WalletFilterer) FilterSubmittedSpendLimitUpdate(opts *bind.FilterOpts) (*WalletSubmittedSpendLimitUpdateIterator, error) {

	logs, sub, err := _Wallet.contract.FilterLogs(opts, "SubmittedSpendLimitUpdate")
	if err != nil {
		return nil, err
	}
	return &WalletSubmittedSpendLimitUpdateIterator{contract: _Wallet.contract, event: "SubmittedSpendLimitUpdate", logs: logs, sub: sub}, nil
}

// WatchSubmittedSpendLimitUpdate is a free log subscription operation binding the contract event 0x4b1b970c8a0fa761e7803ed70c13d7aca71904b13df60fbe03f981da1730da91.
//
// Solidity: event SubmittedSpendLimitUpdate(uint256 _amount)
func (_Wallet *WalletFilterer) WatchSubmittedSpendLimitUpdate(opts *bind.WatchOpts, sink chan<- *WalletSubmittedSpendLimitUpdate) (event.Subscription, error) {

	logs, sub, err := _Wallet.contract.WatchLogs(opts, "SubmittedSpendLimitUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletSubmittedSpendLimitUpdate)
				if err := _Wallet.contract.UnpackLog(event, "SubmittedSpendLimitUpdate", log); err != nil {
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

// ParseSubmittedSpendLimitUpdate is a log parse operation binding the contract event 0x4b1b970c8a0fa761e7803ed70c13d7aca71904b13df60fbe03f981da1730da91.
//
// Solidity: event SubmittedSpendLimitUpdate(uint256 _amount)
func (_Wallet *WalletFilterer) ParseSubmittedSpendLimitUpdate(log types.Log) (*WalletSubmittedSpendLimitUpdate, error) {
	event := new(WalletSubmittedSpendLimitUpdate)
	if err := _Wallet.contract.UnpackLog(event, "SubmittedSpendLimitUpdate", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WalletSubmittedWhitelistAdditionIterator is returned from FilterSubmittedWhitelistAddition and is used to iterate over the raw logs and unpacked data for SubmittedWhitelistAddition events raised by the Wallet contract.
type WalletSubmittedWhitelistAdditionIterator struct {
	Event *WalletSubmittedWhitelistAddition // Event containing the contract specifics and raw log

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
func (it *WalletSubmittedWhitelistAdditionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletSubmittedWhitelistAddition)
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
		it.Event = new(WalletSubmittedWhitelistAddition)
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
func (it *WalletSubmittedWhitelistAdditionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletSubmittedWhitelistAdditionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletSubmittedWhitelistAddition represents a SubmittedWhitelistAddition event raised by the Wallet contract.
type WalletSubmittedWhitelistAddition struct {
	Addresses []common.Address
	Hash      [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSubmittedWhitelistAddition is a free log retrieval operation binding the contract event 0x9c80b3b5f68b3e017766d59e8d09b34efe6462b05c398f35cab9e271d9bc3b9c.
//
// Solidity: event SubmittedWhitelistAddition(address[] _addresses, bytes32 _hash)
func (_Wallet *WalletFilterer) FilterSubmittedWhitelistAddition(opts *bind.FilterOpts) (*WalletSubmittedWhitelistAdditionIterator, error) {

	logs, sub, err := _Wallet.contract.FilterLogs(opts, "SubmittedWhitelistAddition")
	if err != nil {
		return nil, err
	}
	return &WalletSubmittedWhitelistAdditionIterator{contract: _Wallet.contract, event: "SubmittedWhitelistAddition", logs: logs, sub: sub}, nil
}

// WatchSubmittedWhitelistAddition is a free log subscription operation binding the contract event 0x9c80b3b5f68b3e017766d59e8d09b34efe6462b05c398f35cab9e271d9bc3b9c.
//
// Solidity: event SubmittedWhitelistAddition(address[] _addresses, bytes32 _hash)
func (_Wallet *WalletFilterer) WatchSubmittedWhitelistAddition(opts *bind.WatchOpts, sink chan<- *WalletSubmittedWhitelistAddition) (event.Subscription, error) {

	logs, sub, err := _Wallet.contract.WatchLogs(opts, "SubmittedWhitelistAddition")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletSubmittedWhitelistAddition)
				if err := _Wallet.contract.UnpackLog(event, "SubmittedWhitelistAddition", log); err != nil {
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

// ParseSubmittedWhitelistAddition is a log parse operation binding the contract event 0x9c80b3b5f68b3e017766d59e8d09b34efe6462b05c398f35cab9e271d9bc3b9c.
//
// Solidity: event SubmittedWhitelistAddition(address[] _addresses, bytes32 _hash)
func (_Wallet *WalletFilterer) ParseSubmittedWhitelistAddition(log types.Log) (*WalletSubmittedWhitelistAddition, error) {
	event := new(WalletSubmittedWhitelistAddition)
	if err := _Wallet.contract.UnpackLog(event, "SubmittedWhitelistAddition", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WalletSubmittedWhitelistRemovalIterator is returned from FilterSubmittedWhitelistRemoval and is used to iterate over the raw logs and unpacked data for SubmittedWhitelistRemoval events raised by the Wallet contract.
type WalletSubmittedWhitelistRemovalIterator struct {
	Event *WalletSubmittedWhitelistRemoval // Event containing the contract specifics and raw log

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
func (it *WalletSubmittedWhitelistRemovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletSubmittedWhitelistRemoval)
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
		it.Event = new(WalletSubmittedWhitelistRemoval)
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
func (it *WalletSubmittedWhitelistRemovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletSubmittedWhitelistRemovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletSubmittedWhitelistRemoval represents a SubmittedWhitelistRemoval event raised by the Wallet contract.
type WalletSubmittedWhitelistRemoval struct {
	Addresses []common.Address
	Hash      [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSubmittedWhitelistRemoval is a free log retrieval operation binding the contract event 0xfbc0e5ca6c7e4858daf0fdb185ef5186203e74ec9c64737e93c0aeaec596e1d1.
//
// Solidity: event SubmittedWhitelistRemoval(address[] _addresses, bytes32 _hash)
func (_Wallet *WalletFilterer) FilterSubmittedWhitelistRemoval(opts *bind.FilterOpts) (*WalletSubmittedWhitelistRemovalIterator, error) {

	logs, sub, err := _Wallet.contract.FilterLogs(opts, "SubmittedWhitelistRemoval")
	if err != nil {
		return nil, err
	}
	return &WalletSubmittedWhitelistRemovalIterator{contract: _Wallet.contract, event: "SubmittedWhitelistRemoval", logs: logs, sub: sub}, nil
}

// WatchSubmittedWhitelistRemoval is a free log subscription operation binding the contract event 0xfbc0e5ca6c7e4858daf0fdb185ef5186203e74ec9c64737e93c0aeaec596e1d1.
//
// Solidity: event SubmittedWhitelistRemoval(address[] _addresses, bytes32 _hash)
func (_Wallet *WalletFilterer) WatchSubmittedWhitelistRemoval(opts *bind.WatchOpts, sink chan<- *WalletSubmittedWhitelistRemoval) (event.Subscription, error) {

	logs, sub, err := _Wallet.contract.WatchLogs(opts, "SubmittedWhitelistRemoval")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletSubmittedWhitelistRemoval)
				if err := _Wallet.contract.UnpackLog(event, "SubmittedWhitelistRemoval", log); err != nil {
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

// ParseSubmittedWhitelistRemoval is a log parse operation binding the contract event 0xfbc0e5ca6c7e4858daf0fdb185ef5186203e74ec9c64737e93c0aeaec596e1d1.
//
// Solidity: event SubmittedWhitelistRemoval(address[] _addresses, bytes32 _hash)
func (_Wallet *WalletFilterer) ParseSubmittedWhitelistRemoval(log types.Log) (*WalletSubmittedWhitelistRemoval, error) {
	event := new(WalletSubmittedWhitelistRemoval)
	if err := _Wallet.contract.UnpackLog(event, "SubmittedWhitelistRemoval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WalletToppedUpGasIterator is returned from FilterToppedUpGas and is used to iterate over the raw logs and unpacked data for ToppedUpGas events raised by the Wallet contract.
type WalletToppedUpGasIterator struct {
	Event *WalletToppedUpGas // Event containing the contract specifics and raw log

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
func (it *WalletToppedUpGasIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletToppedUpGas)
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
		it.Event = new(WalletToppedUpGas)
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
func (it *WalletToppedUpGasIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletToppedUpGasIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletToppedUpGas represents a ToppedUpGas event raised by the Wallet contract.
type WalletToppedUpGas struct {
	Sender common.Address
	Owner  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterToppedUpGas is a free log retrieval operation binding the contract event 0x611b7c0d84fda988026215bef9b3e4d81cbceced7e679be6d5e044b588467c0e.
//
// Solidity: event ToppedUpGas(address _sender, address _owner, uint256 _amount)
func (_Wallet *WalletFilterer) FilterToppedUpGas(opts *bind.FilterOpts) (*WalletToppedUpGasIterator, error) {

	logs, sub, err := _Wallet.contract.FilterLogs(opts, "ToppedUpGas")
	if err != nil {
		return nil, err
	}
	return &WalletToppedUpGasIterator{contract: _Wallet.contract, event: "ToppedUpGas", logs: logs, sub: sub}, nil
}

// WatchToppedUpGas is a free log subscription operation binding the contract event 0x611b7c0d84fda988026215bef9b3e4d81cbceced7e679be6d5e044b588467c0e.
//
// Solidity: event ToppedUpGas(address _sender, address _owner, uint256 _amount)
func (_Wallet *WalletFilterer) WatchToppedUpGas(opts *bind.WatchOpts, sink chan<- *WalletToppedUpGas) (event.Subscription, error) {

	logs, sub, err := _Wallet.contract.WatchLogs(opts, "ToppedUpGas")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletToppedUpGas)
				if err := _Wallet.contract.UnpackLog(event, "ToppedUpGas", log); err != nil {
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

// ParseToppedUpGas is a log parse operation binding the contract event 0x611b7c0d84fda988026215bef9b3e4d81cbceced7e679be6d5e044b588467c0e.
//
// Solidity: event ToppedUpGas(address _sender, address _owner, uint256 _amount)
func (_Wallet *WalletFilterer) ParseToppedUpGas(log types.Log) (*WalletToppedUpGas, error) {
	event := new(WalletToppedUpGas)
	if err := _Wallet.contract.UnpackLog(event, "ToppedUpGas", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WalletTransferredIterator is returned from FilterTransferred and is used to iterate over the raw logs and unpacked data for Transferred events raised by the Wallet contract.
type WalletTransferredIterator struct {
	Event *WalletTransferred // Event containing the contract specifics and raw log

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
func (it *WalletTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletTransferred)
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
		it.Event = new(WalletTransferred)
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
func (it *WalletTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletTransferred represents a Transferred event raised by the Wallet contract.
type WalletTransferred struct {
	To     common.Address
	Asset  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransferred is a free log retrieval operation binding the contract event 0xd1ba4ac2e2a11b5101f6cb4d978f514a155b421e8ec396d2d9abaf0bb02917ee.
//
// Solidity: event Transferred(address _to, address _asset, uint256 _amount)
func (_Wallet *WalletFilterer) FilterTransferred(opts *bind.FilterOpts) (*WalletTransferredIterator, error) {

	logs, sub, err := _Wallet.contract.FilterLogs(opts, "Transferred")
	if err != nil {
		return nil, err
	}
	return &WalletTransferredIterator{contract: _Wallet.contract, event: "Transferred", logs: logs, sub: sub}, nil
}

// WatchTransferred is a free log subscription operation binding the contract event 0xd1ba4ac2e2a11b5101f6cb4d978f514a155b421e8ec396d2d9abaf0bb02917ee.
//
// Solidity: event Transferred(address _to, address _asset, uint256 _amount)
func (_Wallet *WalletFilterer) WatchTransferred(opts *bind.WatchOpts, sink chan<- *WalletTransferred) (event.Subscription, error) {

	logs, sub, err := _Wallet.contract.WatchLogs(opts, "Transferred")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletTransferred)
				if err := _Wallet.contract.UnpackLog(event, "Transferred", log); err != nil {
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

// ParseTransferred is a log parse operation binding the contract event 0xd1ba4ac2e2a11b5101f6cb4d978f514a155b421e8ec396d2d9abaf0bb02917ee.
//
// Solidity: event Transferred(address _to, address _asset, uint256 _amount)
func (_Wallet *WalletFilterer) ParseTransferred(log types.Log) (*WalletTransferred, error) {
	event := new(WalletTransferred)
	if err := _Wallet.contract.UnpackLog(event, "Transferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WalletTransferredOwnershipIterator is returned from FilterTransferredOwnership and is used to iterate over the raw logs and unpacked data for TransferredOwnership events raised by the Wallet contract.
type WalletTransferredOwnershipIterator struct {
	Event *WalletTransferredOwnership // Event containing the contract specifics and raw log

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
func (it *WalletTransferredOwnershipIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletTransferredOwnership)
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
		it.Event = new(WalletTransferredOwnership)
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
func (it *WalletTransferredOwnershipIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletTransferredOwnershipIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletTransferredOwnership represents a TransferredOwnership event raised by the Wallet contract.
type WalletTransferredOwnership struct {
	From common.Address
	To   common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterTransferredOwnership is a free log retrieval operation binding the contract event 0x850b3df64837d7d518b45f5aa64d104652c3b80eb5b34a8e3d9eb666cb7cdea5.
//
// Solidity: event TransferredOwnership(address _from, address _to)
func (_Wallet *WalletFilterer) FilterTransferredOwnership(opts *bind.FilterOpts) (*WalletTransferredOwnershipIterator, error) {

	logs, sub, err := _Wallet.contract.FilterLogs(opts, "TransferredOwnership")
	if err != nil {
		return nil, err
	}
	return &WalletTransferredOwnershipIterator{contract: _Wallet.contract, event: "TransferredOwnership", logs: logs, sub: sub}, nil
}

// WatchTransferredOwnership is a free log subscription operation binding the contract event 0x850b3df64837d7d518b45f5aa64d104652c3b80eb5b34a8e3d9eb666cb7cdea5.
//
// Solidity: event TransferredOwnership(address _from, address _to)
func (_Wallet *WalletFilterer) WatchTransferredOwnership(opts *bind.WatchOpts, sink chan<- *WalletTransferredOwnership) (event.Subscription, error) {

	logs, sub, err := _Wallet.contract.WatchLogs(opts, "TransferredOwnership")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletTransferredOwnership)
				if err := _Wallet.contract.UnpackLog(event, "TransferredOwnership", log); err != nil {
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

// ParseTransferredOwnership is a log parse operation binding the contract event 0x850b3df64837d7d518b45f5aa64d104652c3b80eb5b34a8e3d9eb666cb7cdea5.
//
// Solidity: event TransferredOwnership(address _from, address _to)
func (_Wallet *WalletFilterer) ParseTransferredOwnership(log types.Log) (*WalletTransferredOwnership, error) {
	event := new(WalletTransferredOwnership)
	if err := _Wallet.contract.UnpackLog(event, "TransferredOwnership", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WalletUpdatedAvailableLimitIterator is returned from FilterUpdatedAvailableLimit and is used to iterate over the raw logs and unpacked data for UpdatedAvailableLimit events raised by the Wallet contract.
type WalletUpdatedAvailableLimitIterator struct {
	Event *WalletUpdatedAvailableLimit // Event containing the contract specifics and raw log

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
func (it *WalletUpdatedAvailableLimitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletUpdatedAvailableLimit)
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
		it.Event = new(WalletUpdatedAvailableLimit)
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
func (it *WalletUpdatedAvailableLimitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletUpdatedAvailableLimitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletUpdatedAvailableLimit represents a UpdatedAvailableLimit event raised by the Wallet contract.
type WalletUpdatedAvailableLimit struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUpdatedAvailableLimit is a free log retrieval operation binding the contract event 0xe93bc25276d408d390778e7a8b926f2f67209c43ed540081b951fe128f0d3cd2.
//
// Solidity: event UpdatedAvailableLimit()
func (_Wallet *WalletFilterer) FilterUpdatedAvailableLimit(opts *bind.FilterOpts) (*WalletUpdatedAvailableLimitIterator, error) {

	logs, sub, err := _Wallet.contract.FilterLogs(opts, "UpdatedAvailableLimit")
	if err != nil {
		return nil, err
	}
	return &WalletUpdatedAvailableLimitIterator{contract: _Wallet.contract, event: "UpdatedAvailableLimit", logs: logs, sub: sub}, nil
}

// WatchUpdatedAvailableLimit is a free log subscription operation binding the contract event 0xe93bc25276d408d390778e7a8b926f2f67209c43ed540081b951fe128f0d3cd2.
//
// Solidity: event UpdatedAvailableLimit()
func (_Wallet *WalletFilterer) WatchUpdatedAvailableLimit(opts *bind.WatchOpts, sink chan<- *WalletUpdatedAvailableLimit) (event.Subscription, error) {

	logs, sub, err := _Wallet.contract.WatchLogs(opts, "UpdatedAvailableLimit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletUpdatedAvailableLimit)
				if err := _Wallet.contract.UnpackLog(event, "UpdatedAvailableLimit", log); err != nil {
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

// ParseUpdatedAvailableLimit is a log parse operation binding the contract event 0xe93bc25276d408d390778e7a8b926f2f67209c43ed540081b951fe128f0d3cd2.
//
// Solidity: event UpdatedAvailableLimit()
func (_Wallet *WalletFilterer) ParseUpdatedAvailableLimit(log types.Log) (*WalletUpdatedAvailableLimit, error) {
	event := new(WalletUpdatedAvailableLimit)
	if err := _Wallet.contract.UnpackLog(event, "UpdatedAvailableLimit", log); err != nil {
		return nil, err
	}
	return event, nil
}
