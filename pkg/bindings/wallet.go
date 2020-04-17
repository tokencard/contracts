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
const WalletABI = "[{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_owner_\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_transferable_\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"_ens_\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_tokenWhitelistNode_\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_controllerNode_\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_licenceNode_\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_spendLimit_\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"_addresses\",\"type\":\"address[]\"}],\"name\":\"AddedToWhitelist\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"_assets\",\"type\":\"address[]\"}],\"name\":\"BulkTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"CancelledWhitelistAddition\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"CancelledWhitelistRemoval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_returndata\",\"type\":\"bytes\"}],\"name\":\"ExecutedRelayedTransaction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_destination\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_returndata\",\"type\":\"bytes\"}],\"name\":\"ExecutedTransaction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_currentNonce\",\"type\":\"uint256\"}],\"name\":\"IncreasedRelayNonce\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"LoadedTokenCard\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_locked\",\"type\":\"address\"}],\"name\":\"LockedOwnership\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"Received\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"_addresses\",\"type\":\"address[]\"}],\"name\":\"RemovedFromWhitelist\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"SetGasTopUpLimit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"SetLoadLimit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"SetSpendLimit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"SubmittedGasTopUpLimitUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"SubmittedLoadLimitUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"SubmittedSpendLimitUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"_addresses\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"SubmittedWhitelistAddition\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"_addresses\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"SubmittedWhitelistRemoval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"ToppedUpGas\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"Transferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"TransferredOwnership\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"UpdatedAvailableLimit\",\"type\":\"event\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"constant\":true,\"inputs\":[],\"name\":\"WALLET_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_transactionBatch\",\"type\":\"bytes\"}],\"name\":\"batchExecuteTransaction\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"_assets\",\"type\":\"address[]\"}],\"name\":\"bulkTransfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_addresses\",\"type\":\"address[]\"}],\"name\":\"calculateHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"cancelWhitelistAddition\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"cancelWhitelistRemoval\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"confirmGasTopUpLimitUpdate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"confirmLoadLimitUpdate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"confirmSpendLimitUpdate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"confirmWhitelistAddition\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"confirmWhitelistRemoval\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"controllerNode\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"convertToEther\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"convertToStablecoin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ensRegistry\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_walletAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"executeRelayedTransaction\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_destination\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"executeTransaction\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"gasTopUpLimitAvailable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"gasTopUpLimitControllerConfirmationRequired\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"gasTopUpLimitPending\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"gasTopUpLimitValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"increaseRelayNonce\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isSetWhitelist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isTransferable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hashedData\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"isValidSignature\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"isValidSignature\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"licenceNode\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"loadLimitAvailable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"loadLimitControllerConfirmationRequired\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"loadLimitPending\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"loadLimitValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"loadTokenCard\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingWhitelistAddition\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingWhitelistRemoval\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"relayNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"setGasTopUpLimit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"setLoadLimit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"setSpendLimit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_addresses\",\"type\":\"address[]\"}],\"name\":\"setWhitelist\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"spendLimitAvailable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"spendLimitControllerConfirmationRequired\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"spendLimitPending\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"spendLimitValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"submitGasTopUpLimitUpdate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"submitLoadLimitUpdate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"submitSpendLimitUpdate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_addresses\",\"type\":\"address[]\"}],\"name\":\"submitWhitelistAddition\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_addresses\",\"type\":\"address[]\"}],\"name\":\"submitWhitelistRemoval\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"submittedWhitelistAddition\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"submittedWhitelistRemoval\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_interfaceID\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tokenWhitelistNode\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"topUpGas\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_transferable\",\"type\":\"bool\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"whitelistArray\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"whitelistMap\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// WalletBin is the compiled bytecode used for deploying new contracts.
var WalletBin = "0x60806040527f7f2ce995617d2816b426c5c8698c5ec2952f7a34bb10f38326f74933d58936976001553480156200003557600080fd5b5060405162005aca38038062005aca833981810160405260e08110156200005b57600080fd5b508051602082015160408301516060840151608085015160a086015160c090960151600080546001600160a01b0319166001600160a01b03861617905594959394929391929091908084808989878015620000b65760018190555b50600280546001600160a01b0319166001600160a01b0384161760ff60a01b1916600160a01b8315158102919091179182905560ff9104166200013057604080516001600160a01b038416815290517f808639ff9c8e4732d60b6c2330de498035416d229f27a77d259680895efec1229181900360200190a15b60408051600081526001600160a01b038416602082015281517f850b3df64837d7d518b45f5aa64d104652c3b80eb5b34a8e3d9eb666cb7cdea5929181900390910190a150506040805160a0810182526706f05b59d3b2000080825260208201819052429282018390526000606083018190526080909201829052600381905560045560059190915560068190556007805460ff19169055600891909155620001e16001600160e01b03620002d816565b5050505050915050600081116200022f576040805162461bcd60e51b815260206004820152600d60248201526c37379039ba30b13632b1b7b4b760991b604482015290519081900360640190fd5b6127100260098190556040805160a08082018352838252602080830185905242838501819052600060608086018290526080958601829052600a889055600b97909755600c829055600d819055600e805460ff1990811690915586519485018752898552928401899052948301819052948201849052910182905260148590556015949094556016919091556017555060188054909116905550601a5550620005549350505050565b6060600080600080600080620002f66008546200045a60201b60201c565b6001600160a01b0316633efec5e96040518163ffffffff1660e01b815260040160006040518083038186803b1580156200032f57600080fd5b505afa15801562000344573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405260e08110156200036e57600080fd5b81019080805160405193929190846401000000008211156200038f57600080fd5b908301906020820185811115620003a557600080fd5b8251640100000000811182820188101715620003c057600080fd5b82525081516020918201929091019080838360005b83811015620003ef578181015183820152602001620003d5565b50505050905090810190601f1680156200041d5780820380516001836020036101000a031916815260200191505b5060409081526020820151908201516060830151608084015160a085015160c090950151979f939e50919c509a5098509096509294509192505050565b6000805460408051630178b8bf60e01b81526004810185905290516001600160a01b0390921691630178b8bf91602480820192602092909190829003018186803b158015620004a857600080fd5b505afa158015620004bd573d6000803e3d6000fd5b505050506040513d6020811015620004d457600080fd5b505160408051631d9dabef60e11b81526004810185905290516001600160a01b0390921691633b3b57de91602480820192602092909190829003018186803b1580156200052057600080fd5b505afa15801562000535573d6000803e3d6000fd5b505050506040513d60208110156200054c57600080fd5b505192915050565b61556680620005646000396000f3fe60806040526004361061038c5760003560e01c80637fd004fa116101dc578063cc0e7e5611610102578063e2b4ce97116100a0578063f40b51f81161006f578063f40b51f8146110be578063f41c4319146110e8578063f421764814611112578063f776f5181461118d5761038c565b8063e2b4ce971461101c578063e61c51ca14611031578063eadd3cea1461105b578063f36febda146110855761038c565b8063ce0b5bd5116100dc578063ce0b5bd514610f9e578063d251fefc14610fc8578063da84b1ed14610ff2578063de212bf3146110075761038c565b8063cc0e7e5614610ec6578063cccdc55614610edb578063cd7958dd14610ef05761038c565b8063b221f3161161017a578063be40ba7911610149578063be40ba7914610e2f578063beabacc814610e44578063c4856cd914610e87578063cbd2ac6814610e9c5761038c565b8063b221f31614610d7c578063b242e53414610da6578063b87e21ef14610de1578063bcb8b74a14610e1a5761038c565b806390e690c7116101b657806390e690c714610c8c5780639b0dfd2714610ca1578063aaf1fc6214610cb6578063ab20599314610d675761038c565b80637fd004fa14610be7578063877337b014610c625780638da5cb5b14610c775761038c565b806332531c3c116102c15780635adc02ab1161025f57806374624c551161022e57806374624c5514610b62578063747c31d614610b8c5780637d73b23114610ba15780637d7d004614610bd25761038c565b80635adc02ab14610a935780635d2362a814610abd5780636137d67014610ad2578063715018a614610b4d5761038c565b80633c672eb71161029b5780633c672eb7146108ae5780633f579f42146108d8578063465b6ac71461099e57806347b55a9d14610a7e5761038c565b806332531c3c146108255780633a43199f146108585780633bfec254146108845761038c565b80631efd02991161032e57806321ce918d1161030857806321ce918d1461076c5780632587a6a21461079657806326d05ab2146107ab578063294f4025146107c05761038c565b80631efd02991461067757806320c13b0b1461068c5780632121dc75146107575761038c565b8063100f23fd1161036a578063100f23fd146104635780631127b57e1461048d5780631626ba7e146105175780631aa21fba146105ec5761038c565b806301ffc9a7146103c8578063027ef3eb146104105780630f3a85d814610437575b6040805133815234602082015281517f88a5966d370b9919b20f3e2c13ff65706f196a4e32cc2c12bf57088f88525874929181900390910190a1005b3480156103d457600080fd5b506103fc600480360360208110156103eb57600080fd5b50356001600160e01b0319166111a2565b604080519115158252519081900360200190f35b34801561041c57600080fd5b506104256111bc565b60408051918252519081900360200190f35b34801561044357600080fd5b506104616004803603602081101561045a57600080fd5b50356111c3565b005b34801561046f57600080fd5b506104616004803603602081101561048657600080fd5b50356112cf565b34801561049957600080fd5b506104a2611474565b6040805160208082528351818301528351919283929083019185019080838360005b838110156104dc5781810151838201526020016104c4565b50505050905090810190601f1680156105095780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561052357600080fd5b506105cf6004803603604081101561053a57600080fd5b81359190810190604081016020820135600160201b81111561055b57600080fd5b82018360208201111561056d57600080fd5b803590602001918460018302840111600160201b8311171561058e57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550611495945050505050565b604080516001600160e01b03199092168252519081900360200190f35b3480156105f857600080fd5b506104616004803603604081101561060f57600080fd5b6001600160a01b038235169190810190604081016020820135600160201b81111561063957600080fd5b82018360208201111561064b57600080fd5b803590602001918460208302840111600160201b8311171561066c57600080fd5b50909250905061150a565b34801561068357600080fd5b50610425611690565b34801561069857600080fd5b506105cf600480360360408110156106af57600080fd5b810190602081018135600160201b8111156106c957600080fd5b8201836020820111156106db57600080fd5b803590602001918460018302840111600160201b831117156106fc57600080fd5b919390929091602081019035600160201b81111561071957600080fd5b82018360208201111561072b57600080fd5b803590602001918460018302840111600160201b8311171561074c57600080fd5b5090925090506116a1565b34801561076357600080fd5b506103fc611776565b34801561077857600080fd5b506104616004803603602081101561078f57600080fd5b5035611786565b3480156107a257600080fd5b50610425611824565b3480156107b757600080fd5b506103fc61182a565b3480156107cc57600080fd5b506107d5611833565b60408051602080825283518183015283519192839290830191858101910280838360005b838110156108115781810151838201526020016107f9565b505050509050019250505060405180910390f35b34801561083157600080fd5b506103fc6004803603602081101561084857600080fd5b50356001600160a01b0316611895565b6104616004803603604081101561086e57600080fd5b506001600160a01b0381351690602001356118aa565b34801561089057600080fd5b50610461600480360360208110156108a757600080fd5b5035611ae8565b3480156108ba57600080fd5b50610461600480360360208110156108d157600080fd5b5035611be0565b3480156108e457600080fd5b506104a2600480360360608110156108fb57600080fd5b6001600160a01b0382351691602081013591810190606081016040820135600160201b81111561092a57600080fd5b82018360208201111561093c57600080fd5b803590602001918460018302840111600160201b8311171561095d57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550611c86945050505050565b3480156109aa57600080fd5b50610461600480360360808110156109c157600080fd5b8135916001600160a01b0360208201351691810190606081016040820135600160201b8111156109f057600080fd5b820183602082011115610a0257600080fd5b803590602001918460018302840111600160201b83111715610a2357600080fd5b919390929091602081019035600160201b811115610a4057600080fd5b820183602082011115610a5257600080fd5b803590602001918460018302840111600160201b83111715610a7357600080fd5b50909250905061217d565b348015610a8a57600080fd5b506107d5612510565b348015610a9f57600080fd5b5061046160048036036020811015610ab657600080fd5b5035612570565b348015610ac957600080fd5b50610425612840565b348015610ade57600080fd5b5061046160048036036020811015610af557600080fd5b810190602081018135600160201b811115610b0f57600080fd5b820183602082011115610b2157600080fd5b803590602001918460208302840111600160201b83111715610b4257600080fd5b50909250905061284c565b348015610b5957600080fd5b50610461612a72565b348015610b6e57600080fd5b5061046160048036036020811015610b8557600080fd5b5035612b70565b348015610b9857600080fd5b50610425612c74565b348015610bad57600080fd5b50610bb6612c7a565b604080516001600160a01b039092168252519081900360200190f35b348015610bde57600080fd5b50610425612c89565b348015610bf357600080fd5b5061046160048036036020811015610c0a57600080fd5b810190602081018135600160201b811115610c2457600080fd5b820183602082011115610c3657600080fd5b803590602001918460208302840111600160201b83111715610c5757600080fd5b509092509050612c95565b348015610c6e57600080fd5b50610425612fd7565b348015610c8357600080fd5b50610bb6612fdd565b348015610c9857600080fd5b50610461612fec565b348015610cad57600080fd5b50610425613049565b348015610cc257600080fd5b5061046160048036036020811015610cd957600080fd5b810190602081018135600160201b811115610cf357600080fd5b820183602082011115610d0557600080fd5b803590602001918460018302840111600160201b83111715610d2657600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955061304f945050505050565b348015610d7357600080fd5b506103fc61318c565b348015610d8857600080fd5b5061046160048036036020811015610d9f57600080fd5b5035613195565b348015610db257600080fd5b5061046160048036036040811015610dc957600080fd5b506001600160a01b0381351690602001351515613285565b348015610ded57600080fd5b5061042560048036036040811015610e0457600080fd5b506001600160a01b03813516906020013561343f565b348015610e2657600080fd5b506103fc6134cf565b348015610e3b57600080fd5b506103fc6134d8565b348015610e5057600080fd5b5061046160048036036060811015610e6757600080fd5b506001600160a01b038135811691602081013590911690604001356134e7565b348015610e9357600080fd5b50610425613671565b348015610ea857600080fd5b5061046160048036036020811015610ebf57600080fd5b5035613677565b348015610ed257600080fd5b506104256139f4565b348015610ee757600080fd5b506104256139fa565b348015610efc57600080fd5b5061042560048036036020811015610f1357600080fd5b810190602081018135600160201b811115610f2d57600080fd5b820183602082011115610f3f57600080fd5b803590602001918460208302840111600160201b83111715610f6057600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929550613a00945050505050565b348015610faa57600080fd5b5061046160048036036020811015610fc157600080fd5b5035613a5a565b348015610fd457600080fd5b50610bb660048036036020811015610feb57600080fd5b5035613c03565b348015610ffe57600080fd5b50610425613c2a565b34801561101357600080fd5b506103fc613c30565b34801561102857600080fd5b50610425613c3e565b34801561103d57600080fd5b506104616004803603602081101561105457600080fd5b5035613c44565b34801561106757600080fd5b506104616004803603602081101561107e57600080fd5b5035613d8e565b34801561109157600080fd5b50610425600480360360408110156110a857600080fd5b506001600160a01b038135169060200135613de7565b3480156110ca57600080fd5b50610461600480360360208110156110e157600080fd5b5035613f9a565b3480156110f457600080fd5b506104616004803603602081101561110b57600080fd5b5035613ff3565b34801561111e57600080fd5b506104616004803603602081101561113557600080fd5b810190602081018135600160201b81111561114f57600080fd5b82018360208201111561116157600080fd5b803590602001918460208302840111600160201b8311171561118257600080fd5b50909250905061404c565b34801561119957600080fd5b506103fc61439e565b6001600160e01b031981166301ffc9a760e01b145b919050565b6017545b90565b6111cc336143a7565b806111d657503330145b61121a576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b8066038d7ea4c680001115801561123957506706f05b59d3b200008111155b611280576040805162461bcd60e51b815260206004820152601360248201527206f7574206f662072616e676520746f702d757606c1b604482015290519081900360640190fd5b61129160038263ffffffff6143bb16565b604080513381526020810183905281517f41ff5d5ce3b7935893a4e7269ec5caae9cca5e3bf0eb4b21d2f443489667112e929181900390910190a150565b6112d8336143a7565b806112e757506112e733614424565b611331576040805162461bcd60e51b815260206004820152601660248201527537b7363c9037bbb732b93e3e31b7b73a3937b63632b960511b604482015290519081900360640190fd5b60135460ff16611380576040805162461bcd60e51b81526020600482015260156024820152743737903832b73234b7339039bab136b4b9b9b4b7b760591b604482015290519081900360640190fd5b6113e360118054806020026020016040519081016040528092919081815260200182805480156113d957602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116113bb575b5050505050613a00565b81146114205760405162461bcd60e51b81526004018080602001828103825260238152602001806154d96023913960400191505060405180910390fd5b61142c60116000615349565b6013805460ff19169055604080513381526020810183905281517f7794eff834d760583543e6e510e717a5e66d2c064e225f4db448343c3e66afcf929181900390910190a150565b60405180604001604052806005815260200164332e312e3160d81b81525081565b6000806114a8848463ffffffff6144b816565b90506114b3816143a7565b6114f8576040805162461bcd60e51b8152602060048201526011602482015270696e76616c6964207369676e617475726560781b604482015290519081900360640190fd5b50630b135d3f60e11b90505b92915050565b611513336143a7565b8061151d57503330145b611561576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b806115aa576040805162461bcd60e51b8152602060048201526014602482015273617373657420617272617920697320656d70747960601b604482015290519081900360640190fd5b60005b8181101561160d5760006115dc308585858181106115c757fe5b905060200201356001600160a01b03166145a6565b9050611604858585858181106115ee57fe5b905060200201356001600160a01b0316836134e7565b506001016115ad565b507fd4f62f23021706247dcffea245d104ae7ddaec7f23acf3d11d7136d5de6a69ad83838360405180846001600160a01b03166001600160a01b03168152602001806020018281038252848482818152602001925060200280828437600083820152604051601f909101601f1916909201829003965090945050505050a1505050565b600061169c600a614651565b905090565b6000808585604051602001808383808284376040805191909301818103601f190182528084528151602092830120601f8b01839004830282018301909452898152929650630b135d3f60e11b9550611718945086935089915088908190840183828082843760009201919091525061149592505050565b6001600160e01b03191614611764576040805162461bcd60e51b815260206004820152600d60248201526c1cda59c81b9bdd081d985b1a59609a1b604482015290519081900360640190fd5b506320c13b0b60e01b95945050505050565b600254600160a01b900460ff1690565b61178f336143a7565b8061179957503330145b6117dd576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b6117ee60148263ffffffff61468616565b6040805182815290517f4b1b970c8a0fa761e7803ed70c13d7aca71904b13df60fbe03f981da1730da919181900360200190a150565b60035490565b60135460ff1681565b6060601280548060200260200160405190810160405280929190818152602001828054801561188b57602002820191906000526020600020905b81546001600160a01b0316815260019091019060200180831161186d575b5050505050905090565b600f6020526000908152604090205460ff1681565b6118b3336143a7565b806118bd57503330145b611901576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b61190a826146e7565b611950576040805162461bcd60e51b8152602060048201526012602482015271746f6b656e206e6f74206c6f616461626c6560701b604482015290519081900360640190fd5b600061195c8383613de7565b905061196f600a8263ffffffff61470116565b600061197c601a54614777565b90506001600160a01b03841615611a24576119a76001600160a01b038516828563ffffffff61483916565b806001600160a01b0316631b3c96b485856040518363ffffffff1660e01b815260040180836001600160a01b03166001600160a01b0316815260200182815260200192505050600060405180830381600087803b158015611a0757600080fd5b505af1158015611a1b573d6000803e3d6000fd5b50505050611a9e565b806001600160a01b0316631b3c96b48486866040518463ffffffff1660e01b815260040180836001600160a01b03166001600160a01b03168152602001828152602001925050506000604051808303818588803b158015611a8457600080fd5b505af1158015611a98573d6000803e3d6000fd5b50505050505b604080516001600160a01b03861681526020810185905281517f5f65674bec9af81f71be68674135a0ea3f163fb91984e3893d06da9f6ea2ce8a929181900390910190a150505050565b611af1336143a7565b80611afb57503330145b611b3f576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b600954811115611b91576040805162461bcd60e51b81526020600482015260186024820152771bdd5d081bd9881c985b99d9481b1bd85908185b5bdd5b9d60421b604482015290519081900360640190fd5b611ba2600a8263ffffffff6143bb16565b604080513381526020810183905281517f0b05243483e17c3f3377aee82b7d47e5700b48288695fc08b7ecc2759afa44ef929181900390910190a150565b611be9336143a7565b80611bf357503330145b611c37576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b611c4860148263ffffffff6143bb16565b604080513381526020810183905281517f068f112e5ec923d412be64779fe69e0fcbb6784c6617e94cccc8fd348f2e0f21929181900390910190a150565b6060611c91336143a7565b80611c9b57503330145b611cdf576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b6001600160a01b0384166000908152600f602052604090205460ff16611d1057611d1060148463ffffffff61470116565b611d22846001600160a01b0316614951565b8015611d325750611d3284614957565b15611f1957600080611d448685614971565b6001600160a01b0382166000908152600f6020526040902054919350915060ff16611d8a576000611d75878361343f565b9050611d8860148263ffffffff61470116565b505b611da36001600160a01b0387168563ffffffff614a7b16565b604080516020808252818301909252606091602082018180388339019050509050600160f81b81601f81518110611dd657fe5b60200101906001600160f81b031916908160001a9053507ff77753fab406ecfff96d6ff2476c64a838fa9f6d37b1bf190f8546e395e3b6138787878460405180856001600160a01b03166001600160a01b031681526020018481526020018060200180602001838103835285818151815260200191508051906020019080838360005b83811015611e71578181015183820152602001611e59565b50505050905090810190601f168015611e9e5780820380516001836020036101000a031916815260200191505b50838103825284518152845160209182019186019080838360005b83811015611ed1578181015183820152602001611eb9565b50505050905090810190601f168015611efe5780820380516001836020036101000a031916815260200191505b50965050505050505060405180910390a19250612176915050565b60006060856001600160a01b031685856040518082805190602001908083835b60208310611f585780518252601f199092019160209182019101611f39565b6001836020036101000a03801982511681845116808217855250505050505090500191505060006040518083038185875af1925050503d8060008114611fba576040519150601f19603f3d011682016040523d82523d6000602084013e611fbf565b606091505b509150915081819061204f5760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b83811015612014578181015183820152602001611ffc565b50505050905090810190601f1680156120415780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b507ff77753fab406ecfff96d6ff2476c64a838fa9f6d37b1bf190f8546e395e3b6138686868460405180856001600160a01b03166001600160a01b031681526020018481526020018060200180602001838103835285818151815260200191508051906020019080838360005b838110156120d45781810151838201526020016120bc565b50505050905090810190601f1680156121015780820380516001836020036101000a031916815260200191505b50838103825284518152845160209182019186019080838360005b8381101561213457818101518382015260200161211c565b50505050905090810190601f1680156121615780820380516001836020036101000a031916815260200191505b50965050505050505060405180910390a19150505b9392505050565b61218633614424565b6121c5576040805162461bcd60e51b815260206004820152601a602482015260008051602061544b833981519152604482015290519081900360640190fd5b60006122338787878760405160200180806339363c1d60e11b815250600401858152602001846001600160a01b03166001600160a01b031660601b81526014018383808284378083019250505094505050505060405160208183030381529060405280519060200120614c39565b9050631626ba7e60e01b6001600160e01b0319166122878285858080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061149592505050565b6001600160e01b031916146122d3576040805162461bcd60e51b815260206004820152600d60248201526c1cda59c81b9bdd081d985b1a59609a1b604482015290519081900360640190fd5b601954871461231c576040805162461bcd60e51b815260206004820152601060248201526f5478207265706c61793a206e6f6e636560801b604482015290519081900360640190fd5b612324614c8a565b6001600160a01b0386163014612381576040805162461bcd60e51b815260206004820152601960248201527f5478207265706c61793a2077616c6c6574206164647265737300000000000000604482015290519081900360640190fd5b60006060306001600160a01b03168787604051808383808284376040519201945060009350909150508083038183865af19150503d80600081146123e1576040519150601f19603f3d011682016040523d82523d6000602084013e6123e6565b606091505b50915091508181906124395760405162461bcd60e51b8152602060048201818152835160248401528351909283926044909101919085019080838360008315612014578181015183820152602001611ffc565b507f823dbcf2b7b0f265871963ca65ac033f6b4c71e0d82cd123d2ff23d752dc21c187878360405180806020018060200183810383528686828181526020019250808284376000838201819052601f909101601f191690920185810384528651815286516020918201939188019250908190849084905b838110156124c85781810151838201526020016124b0565b50505050905090810190601f1680156124f55780820380516001836020036101000a031916815260200191505b509550505050505060405180910390a1505050505050505050565b6060601180548060200260200160405190810160405280929190818152602001828054801561188b576020028201919060005260206000209081546001600160a01b0316815260019091019060200180831161186d575050505050905090565b61257933614424565b6125b8576040805162461bcd60e51b815260206004820152601a602482015260008051602061544b833981519152604482015290519081900360640190fd5b60135460ff16612607576040805162461bcd60e51b81526020600482015260156024820152743737903832b73234b7339039bab136b4b9b9b4b7b760591b604482015290519081900360640190fd5b61266860118054806020026020016040519081016040528092919081815260200182805480156113d9576020028201919060005260206000209081546001600160a01b031681526001909101906020018083116113bb575050505050613a00565b81146126a55760405162461bcd60e51b81526004018080602001828103825260238152602001806154d96023913960400191505060405180910390fd5b60005b60115481101561278c57600f6000601183815481106126c357fe5b60009182526020808320909101546001600160a01b0316835282019290925260400190205460ff16612784576001600f60006011848154811061270257fe5b6000918252602080832091909101546001600160a01b031683528201929092526040019020805460ff191691151591909117905560118054601091908390811061274857fe5b60009182526020808320909101548354600181018555938352912090910180546001600160a01b0319166001600160a01b039092169190911790555b6001016126a8565b507fb2f6cccee7a369e23e293c25aa19bef80af11eb26deba3ea0f2a02783f752e4a33601160405180836001600160a01b03166001600160a01b0316815260200180602001828103825283818154815260200191508054801561281857602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116127fa575b5050935050505060405180910390a161283360116000615349565b506013805460ff19169055565b600061169c6014614651565b612855336143a7565b8061285f57503330145b6128a3576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b60135460ff161580156128be5750601354610100900460ff16155b61290f576040805162461bcd60e51b815260206004820152601c60248201527f77686974656c6973742073756d62697373696f6e2070656e64696e6700000000604482015290519081900360640190fd5b60135462010000900460ff16612968576040805162461bcd60e51b81526020600482015260196024820152781dda1a5d195b1a5cdd081b9bdd081a5b9a5d1a585b1a5e9959603a1b604482015290519081900360640190fd5b806129ac576040805162461bcd60e51b815260206004820152600f60248201526e195b5c1d1e481dda1a5d195b1a5cdd608a1b604482015290519081900360640190fd5b6129b860128383615367565b506013805461ff00191661010017905560408051602080840282810182019093528382527ffbc0e5ca6c7e4858daf0fdb185ef5186203e74ec9c64737e93c0aeaec596e1d19285928592612a2792859185918291850190849080828437600092019190915250613a0092505050565b60405180806020018381526020018281038252858582818152602001925060200280828437600083820152604051601f909101601f1916909201829003965090945050505050a15050565b612a7b336143a7565b612ac5576040805162461bcd60e51b815260206004820152601660248201527539b2b73232b91034b9903737ba1030b71037bbb732b960511b604482015290519081900360640190fd5b600254600160a01b900460ff16612b23576040805162461bcd60e51b815260206004820152601d60248201527f6f776e657273686970206973206e6f74207472616e7366657261626c65000000604482015290519081900360640190fd5b600280546001600160a01b0319169055604080516000808252602082015281517f850b3df64837d7d518b45f5aa64d104652c3b80eb5b34a8e3d9eb666cb7cdea5929181900390910190a1565b612b79336143a7565b80612b8357503330145b612bc7576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b8066038d7ea4c6800011158015612be657506706f05b59d3b200008111155b612c2d576040805162461bcd60e51b815260206004820152601360248201527206f7574206f662072616e676520746f702d757606c1b604482015290519081900360640190fd5b612c3e60038263ffffffff61468616565b6040805182815290517faf2a77cd04c3cc155588dd3bf67b310ab4fb3b1da3cf6b8d7d4d2aa1d09b794c9181900360200190a150565b601a5490565b6000546001600160a01b031690565b600061169c6003614651565b612c9e336143a7565b80612ca857503330145b612cec576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b60135460ff16158015612d075750601354610100900460ff16155b612d58576040805162461bcd60e51b815260206004820152601c60248201527f77686974656c6973742073756d62697373696f6e2070656e64696e6700000000604482015290519081900360640190fd5b8181808060200260200160405190810160405280939291908181526020018383602002808284376000920182905250925050505b8151811015612e7457612db1828281518110612da457fe5b60200260200101516143a7565b15612dfc576040805162461bcd60e51b8152602060048201526016602482015275636f6e7461696e73206f776e6572206164647265737360501b604482015290519081900360640190fd5b60006001600160a01b0316828281518110612e1357fe5b60200260200101516001600160a01b03161415612e6c576040805162461bcd60e51b8152602060048201526012602482015271636f6e7461696e732030206164647265737360701b604482015290519081900360640190fd5b600101612d8c565b5060135462010000900460ff16612ece576040805162461bcd60e51b81526020600482015260196024820152781dda1a5d195b1a5cdd081b9bdd081a5b9a5d1a585b1a5e9959603a1b604482015290519081900360640190fd5b81612f12576040805162461bcd60e51b815260206004820152600f60248201526e195b5c1d1e481dda1a5d195b1a5cdd608a1b604482015290519081900360640190fd5b612f1e60118484615367565b506013805460ff1916600117905560408051602080850282810182019093528482527f9c80b3b5f68b3e017766d59e8d09b34efe6462b05c398f35cab9e271d9bc3b9c9286928692612f8b92859185918291850190849080828437600092019190915250613a0092505050565b60405180806020018381526020018281038252858582818152602001925060200280828437600083820152604051601f909101601f1916909201829003965090945050505050a1505050565b60085490565b6002546001600160a01b031690565b612ff5336143a7565b61303f576040805162461bcd60e51b815260206004820152601660248201527539b2b73232b91034b9903737ba1030b71037bbb732b960511b604482015290519081900360640190fd5b613047614c8a565b565b60145490565b613058336143a7565b8061306257503330145b6130a6576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b8051602080820191906000808060605b86851015613182576130cf86605463ffffffff614cd216565b888601805160148201516034909201805193995060609190911c965090945090925090506131146054613108878563ffffffff614d2f16565b9063ffffffff614d2f16565b94508685111561315b576040805162461bcd60e51b815260206004820152600d60248201526c6f7574206f6620626f756e647360981b604482015290519081900360640190fd5b8161317157506040805160208101909152600081525b61317c848483611c86565b506130b6565b5050505050505050565b600e5460ff1690565b61319e336143a7565b806131a857503330145b6131ec576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b60095481111561323e576040805162461bcd60e51b81526020600482015260186024820152771bdd5d081bd9881c985b99d9481b1bd85908185b5bdd5b9d60421b604482015290519081900360640190fd5b61324f600a8263ffffffff61468616565b6040805182815290517fc178d379965e5657b6fc57494e392f121a14119215dfb422aad7db4cc03f2d109181900360200190a150565b61328e336143a7565b6132d8576040805162461bcd60e51b815260206004820152601660248201527539b2b73232b91034b9903737ba1030b71037bbb732b960511b604482015290519081900360640190fd5b600254600160a01b900460ff16613336576040805162461bcd60e51b815260206004820152601d60248201527f6f776e657273686970206973206e6f74207472616e7366657261626c65000000604482015290519081900360640190fd5b6001600160a01b03821661337b5760405162461bcd60e51b815260040180806020018281038252602381526020018061548c6023913960400191505060405180910390fd5b6002805460ff60a01b1916600160a01b83151502179055806133d457604080516001600160a01b038416815290517f808639ff9c8e4732d60b6c2330de498035416d229f27a77d259680895efec1229181900360200190a15b600254604080516001600160a01b039283168152918416602083015280517f850b3df64837d7d518b45f5aa64d104652c3b80eb5b34a8e3d9eb666cb7cdea59281900390910190a150600280546001600160a01b0319166001600160a01b0392909216919091179055565b60008060008061344e86614d89565b5050509350935093505080156134c35781613499576040805162461bcd60e51b81526020600482015260066024820152650726174653d360d41b604482015290519081900360640190fd5b6134b9836134ad878563ffffffff614f1b16565b9063ffffffff614f7416565b9350505050611504565b50600095945050505050565b60185460ff1690565b60135462010000900460ff1681565b6134f0336143a7565b806134fa57503330145b61353e576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b808061357b576040805162461bcd60e51b8152602060048201526007602482015266076616c75653d360cc1b604482015290519081900360640190fd5b6001600160a01b0384166135c6576040805162461bcd60e51b815260206004820152600d60248201526c064657374696e6174696f6e3d3609c1b604482015290519081900360640190fd5b6001600160a01b0384166000908152600f602052604090205460ff1661361657816001600160a01b0384161561360357613600848461343f565b90505b61361460148263ffffffff61470116565b505b613621848484614fde565b604080516001600160a01b0380871682528516602082015280820184905290517fd1ba4ac2e2a11b5101f6cb4d978f514a155b421e8ec396d2d9abaf0bb02917ee9181900360600190a150505050565b600d5490565b61368033614424565b6136bf576040805162461bcd60e51b815260206004820152601a602482015260008051602061544b833981519152604482015290519081900360640190fd5b601354610100900460ff16613713576040805162461bcd60e51b81526020600482015260156024820152743737903832b73234b7339039bab136b4b9b9b4b7b760591b604482015290519081900360640190fd5b61377460128054806020026020016040519081016040528092919081815260200182805480156113d9576020028201919060005260206000209081546001600160a01b031681526001909101906020018083116113bb575050505050613a00565b81146137b15760405162461bcd60e51b81526004018080602001828103825260238152602001806154d96023913960400191505060405180910390fd5b60005b60125481101561393f57600f6000601283815481106137cf57fe5b60009182526020808320909101546001600160a01b0316835282019290925260400190205460ff1615613937576000600f60006012848154811061380f57fe5b6000918252602080832091909101546001600160a01b031683528201929092526040018120805460ff1916921515929092179091555b60105461385990600163ffffffff614cd216565b811015613921576012828154811061386d57fe5b600091825260209091200154601080546001600160a01b03909216918390811061389357fe5b6000918252602090912001546001600160a01b03161415613919576010805460001981019081106138c057fe5b600091825260209091200154601080546001600160a01b0390921691839081106138e657fe5b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b03160217905550613921565b600101613845565b5060108054906139359060001983016153ca565b505b6001016137b4565b507fd218c430fa348f4ce67791021b6b89c0c3eacd4ead1d8f5b83c60038ec28249b33601260405180836001600160a01b03166001600160a01b031681526020018060200182810382528381815481526020019150805480156139cb57602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116139ad575b5050935050505060405180910390a16139e660126000615349565b506013805461ff0019169055565b60065490565b60195481565b60008160405160200180828051906020019060200280838360005b83811015613a33578181015183820152602001613a1b565b50505050905001915050604051602081830303815290604052805190602001209050919050565b613a63336143a7565b80613a725750613a7233614424565b613abc576040805162461bcd60e51b815260206004820152601660248201527537b7363c9037bbb732b93e3e31b7b73a3937b63632b960511b604482015290519081900360640190fd5b601354610100900460ff16613b10576040805162461bcd60e51b81526020600482015260156024820152743737903832b73234b7339039bab136b4b9b9b4b7b760591b604482015290519081900360640190fd5b613b7160128054806020026020016040519081016040528092919081815260200182805480156113d9576020028201919060005260206000209081546001600160a01b031681526001909101906020018083116113bb575050505050613a00565b8114613bae5760405162461bcd60e51b81526004018080602001828103825260238152602001806154d96023913960400191505060405180910390fd5b613bba60126000615349565b6013805461ff0019169055604080513381526020810183905281517f13c935eb475aa0f6e931fece83e2ac44569ce2d53460d29a6dedab40b965c8a3929181900390910190a150565b60108181548110613c1057fe5b6000918252602090912001546001600160a01b0316905081565b600a5490565b601354610100900460ff1681565b60015490565b8080613c81576040805162461bcd60e51b8152602060048201526007602482015266076616c75653d360cc1b604482015290519081900360640190fd5b613c8a336143a7565b80613c995750613c9933614424565b613ce3576040805162461bcd60e51b815260206004820152601660248201527537b7363c9037bbb732b93e3e31b7b73a3937b63632b960511b604482015290519081900360640190fd5b613cf460038363ffffffff61470116565b613cfc612fdd565b6001600160a01b03166108fc839081150290604051600060405180830381858888f19350505050158015613d34573d6000803e3d6000fd5b507f611b7c0d84fda988026215bef9b3e4d81cbceced7e679be6d5e044b588467c0e33613d5f612fdd565b604080516001600160a01b03938416815291909216602082015280820185905290519081900360600190a15050565b613d9733614424565b613dd6576040805162461bcd60e51b815260206004820152601a602482015260008051602061544b833981519152604482015290519081900360640190fd5b611c4860148263ffffffff61504216565b6000613df1615096565b6001600160a01b0316836001600160a01b03161415613e11575080611504565b816001600160a01b03841615613ed6576000806000613e2f87614d89565b5050509350935093505080613e81576040805162461bcd60e51b8152602060048201526013602482015272746f6b656e206e6f7420617661696c61626c6560681b604482015290519081900360640190fd5b81613ebc576040805162461bcd60e51b81526020600482015260066024820152650726174653d360d41b604482015290519081900360640190fd5b613ed0836134ad888563ffffffff614f1b16565b93505050505b6000806000613ee361510c565b5050509350935093505080613f35576040805162461bcd60e51b8152602060048201526013602482015272746f6b656e206e6f7420617661696c61626c6560681b604482015290519081900360640190fd5b81613f7b576040805162461bcd60e51b81526020600482015260116024820152700737461626c65636f696e20726174653d3607c1b604482015290519081900360640190fd5b613f8f826134ad868663ffffffff614f1b16565b979650505050505050565b613fa333614424565b613fe2576040805162461bcd60e51b815260206004820152601a602482015260008051602061544b833981519152604482015290519081900360640190fd5b611ba2600a8263ffffffff61504216565b613ffc33614424565b61403b576040805162461bcd60e51b815260206004820152601a602482015260008051602061544b833981519152604482015290519081900360640190fd5b61129160038263ffffffff61504216565b614055336143a7565b8061405f57503330145b6140a3576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b8181808060200260200160405190810160405280939291908181526020018383602002808284376000920182905250925050505b81518110156141b2576140ef828281518110612da457fe5b1561413a576040805162461bcd60e51b8152602060048201526016602482015275636f6e7461696e73206f776e6572206164647265737360501b604482015290519081900360640190fd5b60006001600160a01b031682828151811061415157fe5b60200260200101516001600160a01b031614156141aa576040805162461bcd60e51b8152602060048201526012602482015271636f6e7461696e732030206164647265737360701b604482015290519081900360640190fd5b6001016140d7565b5060135462010000900460ff1615614209576040805162461bcd60e51b81526020600482015260156024820152741dda1a5d195b1a5cdd081a5b9a5d1a585b1a5e9959605a1b604482015290519081900360640190fd5b60005b828110156142fa57600f600085858481811061422457fe5b602090810292909201356001600160a01b03168352508101919091526040016000205460ff166142f2576001600f600086868581811061426057fe5b905060200201356001600160a01b03166001600160a01b03166001600160a01b0316815260200190815260200160002060006101000a81548160ff02191690831515021790555060108484838181106142b557fe5b835460018101855560009485526020948590200180546001600160a01b0319166001600160a01b0395909202939093013593909316929092179055505b60010161420c565b506013805462ff0000191662010000179055604080513380825260208201838152601080549484018590527fb2f6cccee7a369e23e293c25aa19bef80af11eb26deba3ea0f2a02783f752e4a949293909290919060608301908490801561438a57602002820191906000526020600020905b81546001600160a01b0316815260019091019060200180831161436c575b5050935050505060405180910390a1505050565b60075460ff1690565b6002546001600160a01b0390811691161490565b600482015460ff1615614409576040805162461bcd60e51b81526020600482015260116024820152701b1a5b5a5d08185b1c9958591e481cd95d607a1b604482015290519081900360640190fd5b614413828261527b565b50600401805460ff19166001179055565b6000614431600154614777565b6001600160a01b031663b429afeb836040518263ffffffff1660e01b815260040180826001600160a01b03166001600160a01b0316815260200191505060206040518083038186803b15801561448657600080fd5b505afa15801561449a573d6000803e3d6000fd5b505050506040513d60208110156144b057600080fd5b505192915050565b600081516041146144cb57506000611504565b60208201516040830151606084015160001a7f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08211156145115760009350505050611504565b8060ff16601b1415801561452957508060ff16601c14155b1561453a5760009350505050611504565b6040805160008152602080820180845289905260ff8416828401526060820186905260808201859052915160019260a0808401939192601f1981019281900390910190855afa158015614591573d6000803e3d6000fd5b5050604051601f190151979650505050505050565b60006001600160a01b0382161561464057816001600160a01b03166370a08231846040518263ffffffff1660e01b815260040180826001600160a01b03166001600160a01b0316815260200191505060206040518083038186803b15801561460d57600080fd5b505afa158015614621573d6000803e3d6000fd5b505050506040513d602081101561463757600080fd5b50519050611504565b506001600160a01b03821631611504565b600281015460009061466c906201518063ffffffff614d2f16565b42111561467b575080546111b7565b5060018101546111b7565b600482015460ff166146df576040805162461bcd60e51b815260206004820152601960248201527f6c696d6974206861736e2774206265656e207365742079657400000000000000604482015290519081900360640190fd5b600390910155565b6000806146f383614d89565b509098975050505050505050565b61470a8261529e565b8082600101541015614756576040805162461bcd60e51b815260206004820152601060248201526f185d985a5b18589b194f185b5bdd5b9d60821b604482015290519081900360640190fd5b600182015461476b908263ffffffff614cd216565b82600101819055505050565b6000805460408051630178b8bf60e01b81526004810185905290516001600160a01b0390921691630178b8bf91602480820192602092909190829003018186803b1580156147c457600080fd5b505afa1580156147d8573d6000803e3d6000fd5b505050506040513d60208110156147ee57600080fd5b505160408051631d9dabef60e11b81526004810185905290516001600160a01b0390921691633b3b57de91602480820192602092909190829003018186803b15801561448657600080fd5b8015806148bf575060408051636eb1769f60e11b81523060048201526001600160a01b03848116602483015291519185169163dd62ed3e91604480820192602092909190829003018186803b15801561489157600080fd5b505afa1580156148a5573d6000803e3d6000fd5b505050506040513d60208110156148bb57600080fd5b5051155b6148fa5760405162461bcd60e51b81526004018080602001828103825260368152602001806154fc6036913960400191505060405180910390fd5b604080516001600160a01b038416602482015260448082018490528251808303909101815260649091019091526020810180516001600160e01b031663095ea7b360e01b17905261494c908490614a7b565b505050565b3b151590565b60008061496383614d89565b509198975050505050505050565b60008061497f600854614777565b6001600160a01b031663afc72e9385856040518363ffffffff1660e01b815260040180836001600160a01b03166001600160a01b0316815260200180602001828103825283818151815260200191508051906020019080838360005b838110156149f35781810151838201526020016149db565b50505050905090810190601f168015614a205780820380516001836020036101000a031916815260200191505b509350505050604080518083038186803b158015614a3d57600080fd5b505afa158015614a51573d6000803e3d6000fd5b505050506040513d6040811015614a6757600080fd5b508051602090910151909590945092505050565b614a8d826001600160a01b0316614951565b614ade576040805162461bcd60e51b815260206004820152601f60248201527f5361666545524332303a2063616c6c20746f206e6f6e2d636f6e747261637400604482015290519081900360640190fd5b60006060836001600160a01b0316836040518082805190602001908083835b60208310614b1c5780518252601f199092019160209182019101614afd565b6001836020036101000a0380198251168184511680821785525050505050509050019150506000604051808303816000865af19150503d8060008114614b7e576040519150601f19603f3d011682016040523d82523d6000602084013e614b83565b606091505b509150915081614bda576040805162461bcd60e51b815260206004820181905260248201527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564604482015290519081900360640190fd5b805115614c3357808060200190516020811015614bf657600080fd5b5051614c335760405162461bcd60e51b815260040180806020018281038252602a8152602001806154af602a913960400191505060405180910390fd5b50505050565b604080517f19457468657265756d205369676e6564204d6573736167653a0a333200000000602080830191909152603c8083019490945282518083039094018452605c909101909152815191012090565b601980546001019081905560408051338152602081019290925280517fab0423a75986556234aecd171c46ce7f5e45607d8070bf5230f2735b50322bff9281900390910190a1565b600082821115614d29576040805162461bcd60e51b815260206004820152601e60248201527f536166654d6174683a207375627472616374696f6e206f766572666c6f770000604482015290519081900360640190fd5b50900390565b600082820183811015612176576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b6060600080600080600080614d9f600854614777565b6001600160a01b0316631f69565f896040518263ffffffff1660e01b815260040180826001600160a01b03166001600160a01b0316815260200191505060006040518083038186803b158015614df457600080fd5b505afa158015614e08573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405260e0811015614e3157600080fd5b8101908080516040519392919084600160201b821115614e5057600080fd5b908301906020820185811115614e6557600080fd5b8251600160201b811182820188101715614e7e57600080fd5b82525081516020918201929091019080838360005b83811015614eab578181015183820152602001614e93565b50505050905090810190601f168015614ed85780820380516001836020036101000a031916815260200191505b5060409081526020820151908201516060830151608084015160a085015160c090950151979e50929c50909a509850965094509192505050919395979092949650565b600082614f2a57506000611504565b82820282848281614f3757fe5b04146121765760405162461bcd60e51b815260040180806020018281038252602181526020018061546b6021913960400191505060405180910390fd5b6000808211614fca576040805162461bcd60e51b815260206004820152601a60248201527f536166654d6174683a206469766973696f6e206279207a65726f000000000000604482015290519081900360640190fd5b6000828481614fd557fe5b04949350505050565b6001600160a01b038216615028576040516001600160a01b0384169082156108fc029083906000818181858888f19350505050158015615022573d6000803e3d6000fd5b5061494c565b61494c6001600160a01b038316848363ffffffff6152f716565b808260030154146150845760405162461bcd60e51b81526004018080602001828103825260228152602001806154296022913960400191505060405180910390fd5b61509282836003015461527b565b5050565b60006150a3600854614777565b6001600160a01b031663e9cbd8226040518163ffffffff1660e01b815260040160206040518083038186803b1580156150db57600080fd5b505afa1580156150ef573d6000803e3d6000fd5b505050506040513d602081101561510557600080fd5b5051905090565b6060600080600080600080615122600854614777565b6001600160a01b0316633efec5e96040518163ffffffff1660e01b815260040160006040518083038186803b15801561515a57600080fd5b505afa15801561516e573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405260e081101561519757600080fd5b8101908080516040519392919084600160201b8211156151b657600080fd5b9083019060208201858111156151cb57600080fd5b8251600160201b8111828201881017156151e457600080fd5b82525081516020918201929091019080838360005b838110156152115781810151838201526020016151f9565b50505050905090810190601f16801561523e5780820380516001836020036101000a031916815260200191505b5060409081526020820151908201516060830151608084015160a085015160c090950151979f939e50919c509a5098509096509294509192505050565b6152848261529e565b808255600182015481101561509257815460018301555050565b60028101546152b6906201518063ffffffff614d2f16565b4211156152f457426002820155805460018201556040517fe93bc25276d408d390778e7a8b926f2f67209c43ed540081b951fe128f0d3cd290600090a15b50565b604080516001600160a01b038416602482015260448082018490528251808303909101815260649091019091526020810180516001600160e01b031663a9059cbb60e01b17905261494c908490614a7b565b50805460008255906000526020600020908101906152f491906153ea565b8280548282559060005260206000209081019282156153ba579160200282015b828111156153ba5781546001600160a01b0319166001600160a01b03843516178255602090920191600190910190615387565b506153c6929150615404565b5090565b81548183558181111561494c5760008381526020902061494c9181019083015b6111c091905b808211156153c657600081556001016153f0565b6111c091905b808211156153c65780546001600160a01b031916815560010161540a56fe636f6e6669726d65642f7375626d6974746564206c696d6974206d69736d6174636873656e646572206973206e6f74206120636f6e74726f6c6c6572000000000000536166654d6174683a206d756c7469706c69636174696f6e206f766572666c6f776f776e65722063616e6e6f742062652073657420746f207a65726f20616464726573735361666545524332303a204552433230206f7065726174696f6e20646964206e6f7420737563636565646e6f6e2d6d61746368696e672070656e64696e672077686974656c69737420686173685361666545524332303a20617070726f76652066726f6d206e6f6e2d7a65726f20746f206e6f6e2d7a65726f20616c6c6f77616e6365a265627a7a72315820f786fd5f25d71e2b2e2bec6949e027f05528624a0bf6d5e44aadde987906425664736f6c63430005110032"

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

// GasTopUpLimitControllerConfirmationRequired is a free data retrieval call binding the contract method 0xf776f518.
//
// Solidity: function gasTopUpLimitControllerConfirmationRequired() constant returns(bool)
func (_Wallet *WalletCaller) GasTopUpLimitControllerConfirmationRequired(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Wallet.contract.Call(opts, out, "gasTopUpLimitControllerConfirmationRequired")
	return *ret0, err
}

// GasTopUpLimitControllerConfirmationRequired is a free data retrieval call binding the contract method 0xf776f518.
//
// Solidity: function gasTopUpLimitControllerConfirmationRequired() constant returns(bool)
func (_Wallet *WalletSession) GasTopUpLimitControllerConfirmationRequired() (bool, error) {
	return _Wallet.Contract.GasTopUpLimitControllerConfirmationRequired(&_Wallet.CallOpts)
}

// GasTopUpLimitControllerConfirmationRequired is a free data retrieval call binding the contract method 0xf776f518.
//
// Solidity: function gasTopUpLimitControllerConfirmationRequired() constant returns(bool)
func (_Wallet *WalletCallerSession) GasTopUpLimitControllerConfirmationRequired() (bool, error) {
	return _Wallet.Contract.GasTopUpLimitControllerConfirmationRequired(&_Wallet.CallOpts)
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

// LoadLimitControllerConfirmationRequired is a free data retrieval call binding the contract method 0xab205993.
//
// Solidity: function loadLimitControllerConfirmationRequired() constant returns(bool)
func (_Wallet *WalletCaller) LoadLimitControllerConfirmationRequired(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Wallet.contract.Call(opts, out, "loadLimitControllerConfirmationRequired")
	return *ret0, err
}

// LoadLimitControllerConfirmationRequired is a free data retrieval call binding the contract method 0xab205993.
//
// Solidity: function loadLimitControllerConfirmationRequired() constant returns(bool)
func (_Wallet *WalletSession) LoadLimitControllerConfirmationRequired() (bool, error) {
	return _Wallet.Contract.LoadLimitControllerConfirmationRequired(&_Wallet.CallOpts)
}

// LoadLimitControllerConfirmationRequired is a free data retrieval call binding the contract method 0xab205993.
//
// Solidity: function loadLimitControllerConfirmationRequired() constant returns(bool)
func (_Wallet *WalletCallerSession) LoadLimitControllerConfirmationRequired() (bool, error) {
	return _Wallet.Contract.LoadLimitControllerConfirmationRequired(&_Wallet.CallOpts)
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

// SpendLimitControllerConfirmationRequired is a free data retrieval call binding the contract method 0xbcb8b74a.
//
// Solidity: function spendLimitControllerConfirmationRequired() constant returns(bool)
func (_Wallet *WalletCaller) SpendLimitControllerConfirmationRequired(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Wallet.contract.Call(opts, out, "spendLimitControllerConfirmationRequired")
	return *ret0, err
}

// SpendLimitControllerConfirmationRequired is a free data retrieval call binding the contract method 0xbcb8b74a.
//
// Solidity: function spendLimitControllerConfirmationRequired() constant returns(bool)
func (_Wallet *WalletSession) SpendLimitControllerConfirmationRequired() (bool, error) {
	return _Wallet.Contract.SpendLimitControllerConfirmationRequired(&_Wallet.CallOpts)
}

// SpendLimitControllerConfirmationRequired is a free data retrieval call binding the contract method 0xbcb8b74a.
//
// Solidity: function spendLimitControllerConfirmationRequired() constant returns(bool)
func (_Wallet *WalletCallerSession) SpendLimitControllerConfirmationRequired() (bool, error) {
	return _Wallet.Contract.SpendLimitControllerConfirmationRequired(&_Wallet.CallOpts)
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

// ExecuteRelayedTransaction is a paid mutator transaction binding the contract method 0x465b6ac7.
//
// Solidity: function executeRelayedTransaction(uint256 _nonce, address _walletAddress, bytes _data, bytes _signature) returns()
func (_Wallet *WalletTransactor) ExecuteRelayedTransaction(opts *bind.TransactOpts, _nonce *big.Int, _walletAddress common.Address, _data []byte, _signature []byte) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "executeRelayedTransaction", _nonce, _walletAddress, _data, _signature)
}

// ExecuteRelayedTransaction is a paid mutator transaction binding the contract method 0x465b6ac7.
//
// Solidity: function executeRelayedTransaction(uint256 _nonce, address _walletAddress, bytes _data, bytes _signature) returns()
func (_Wallet *WalletSession) ExecuteRelayedTransaction(_nonce *big.Int, _walletAddress common.Address, _data []byte, _signature []byte) (*types.Transaction, error) {
	return _Wallet.Contract.ExecuteRelayedTransaction(&_Wallet.TransactOpts, _nonce, _walletAddress, _data, _signature)
}

// ExecuteRelayedTransaction is a paid mutator transaction binding the contract method 0x465b6ac7.
//
// Solidity: function executeRelayedTransaction(uint256 _nonce, address _walletAddress, bytes _data, bytes _signature) returns()
func (_Wallet *WalletTransactorSession) ExecuteRelayedTransaction(_nonce *big.Int, _walletAddress common.Address, _data []byte, _signature []byte) (*types.Transaction, error) {
	return _Wallet.Contract.ExecuteRelayedTransaction(&_Wallet.TransactOpts, _nonce, _walletAddress, _data, _signature)
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
