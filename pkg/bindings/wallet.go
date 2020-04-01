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
<<<<<<< HEAD
const WalletABI = "[{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_owner_\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_transferable_\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"_ens_\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_tokenWhitelistNode_\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_controllerNode_\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_licenceNode_\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_dailyLimit_\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"_addresses\",\"type\":\"address[]\"}],\"name\":\"AddedToWhitelist\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"_assets\",\"type\":\"address[]\"}],\"name\":\"BulkTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"CancelledWhitelistAddition\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"CancelledWhitelistRemoval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"_privileged\",\"type\":\"bool\"}],\"name\":\"ExecutedRelayedTransaction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_destination\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_returndata\",\"type\":\"bytes\"}],\"name\":\"ExecutedTransaction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_currentNonce\",\"type\":\"uint256\"}],\"name\":\"IncreasedRelayNonce\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_nextReset\",\"type\":\"uint256\"}],\"name\":\"InitializedDailyLimit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"LoadedTokenCard\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_locked\",\"type\":\"address\"}],\"name\":\"LockedOwnership\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"Received\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"_addresses\",\"type\":\"address[]\"}],\"name\":\"RemovedFromWhitelist\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"SetDailyLimit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"SetMonolith2FA\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_p2FA\",\"type\":\"address\"}],\"name\":\"SetPersonal2FA\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"SubmittedDailyLimitUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"_addresses\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"SubmittedWhitelistAddition\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"_addresses\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"SubmittedWhitelistRemoval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"ToppedUpGas\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"Transferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"TransferredOwnership\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_nextReset\",\"type\":\"uint256\"}],\"name\":\"UpdatedAvailableDailyLimit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"UpdatedAvailableLimit\",\"type\":\"event\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"constant\":true,\"inputs\":[],\"name\":\"WALLET_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"_assets\",\"type\":\"address[]\"}],\"name\":\"bulkTransfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_addresses\",\"type\":\"address[]\"}],\"name\":\"calculateHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"cancelWhitelistAddition\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"cancelWhitelistRemoval\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"confirmDailyLimitUpdate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"confirmWhitelistAddition\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"confirmWhitelistRemoval\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"controllerNode\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"convertToStablecoin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"dailyLimitAvailable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"dailyLimitPending\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"dailyLimitValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ensRegistry\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"executePrivilegedRelayedTransaction\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"executeRelayedTransaction\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_destination\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"executeTransaction\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"increaseRelayNonce\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isSetWhitelist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isTransferable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hashedData\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"isValidSignature\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"isValidSignature\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"licenceNode\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"loadTokenCard\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"monolith2FA\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingWhitelistAddition\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingWhitelistRemoval\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"personal2FA\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"privileged\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"relayNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"setMonolith2FA\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_p2FA\",\"type\":\"address\"}],\"name\":\"setPersonal2FA\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_addresses\",\"type\":\"address[]\"}],\"name\":\"setWhitelist\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"submitDailyLimitUpdate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_addresses\",\"type\":\"address[]\"}],\"name\":\"submitWhitelistAddition\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_addresses\",\"type\":\"address[]\"}],\"name\":\"submitWhitelistRemoval\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"submittedWhitelistAddition\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"submittedWhitelistRemoval\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_interfaceID\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tokenWhitelistNode\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_transferable\",\"type\":\"bool\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"whitelistArray\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"whitelistMap\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"
||||||| parent of 64ba3a32... Upgrade to solc 0.6.4 for Wallet
const WalletABI = "[{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_owner_\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_transferable_\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"_ens_\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_tokenWhitelistNode_\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_controllerNode_\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_licenceNode_\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_spendLimit_\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"_addresses\",\"type\":\"address[]\"}],\"name\":\"AddedToWhitelist\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"_assets\",\"type\":\"address[]\"}],\"name\":\"BulkTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"CancelledWhitelistAddition\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"CancelledWhitelistRemoval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_returndata\",\"type\":\"bytes\"}],\"name\":\"ExecutedRelayedTransaction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_destination\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_returndata\",\"type\":\"bytes\"}],\"name\":\"ExecutedTransaction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_currentNonce\",\"type\":\"uint256\"}],\"name\":\"IncreasedRelayNonce\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"LoadedTokenCard\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_locked\",\"type\":\"address\"}],\"name\":\"LockedOwnership\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"Received\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"_addresses\",\"type\":\"address[]\"}],\"name\":\"RemovedFromWhitelist\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"SetGasTopUpLimit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"SetLoadLimit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"SetSpendLimit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"SubmittedGasTopUpLimitUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"SubmittedLoadLimitUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"SubmittedSpendLimitUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"_addresses\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"SubmittedWhitelistAddition\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"_addresses\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"SubmittedWhitelistRemoval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"ToppedUpGas\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"Transferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"TransferredOwnership\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"UpdatedAvailableLimit\",\"type\":\"event\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"constant\":true,\"inputs\":[],\"name\":\"WALLET_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_transactionBatch\",\"type\":\"bytes\"}],\"name\":\"batchExecuteTransaction\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"_assets\",\"type\":\"address[]\"}],\"name\":\"bulkTransfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_addresses\",\"type\":\"address[]\"}],\"name\":\"calculateHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"cancelWhitelistAddition\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"cancelWhitelistRemoval\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"confirmGasTopUpLimitUpdate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"confirmLoadLimitUpdate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"confirmSpendLimitUpdate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"confirmWhitelistAddition\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"confirmWhitelistRemoval\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"controllerNode\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"convertToEther\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"convertToStablecoin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ensRegistry\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"executeRelayedTransaction\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_destination\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"executeTransaction\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"gasTopUpLimitAvailable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"gasTopUpLimitControllerConfirmationRequired\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"gasTopUpLimitPending\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"gasTopUpLimitValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"increaseRelayNonce\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isSetWhitelist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isTransferable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hashedData\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"isValidSignature\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"isValidSignature\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"licenceNode\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"loadLimitAvailable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"loadLimitControllerConfirmationRequired\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"loadLimitPending\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"loadLimitValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"loadTokenCard\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingWhitelistAddition\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingWhitelistRemoval\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"relayNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"setGasTopUpLimit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"setLoadLimit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"setSpendLimit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_addresses\",\"type\":\"address[]\"}],\"name\":\"setWhitelist\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"spendLimitAvailable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"spendLimitControllerConfirmationRequired\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"spendLimitPending\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"spendLimitValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"submitGasTopUpLimitUpdate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"submitLoadLimitUpdate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"submitSpendLimitUpdate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_addresses\",\"type\":\"address[]\"}],\"name\":\"submitWhitelistAddition\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_addresses\",\"type\":\"address[]\"}],\"name\":\"submitWhitelistRemoval\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"submittedWhitelistAddition\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"submittedWhitelistRemoval\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_interfaceID\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tokenWhitelistNode\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"topUpGas\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_transferable\",\"type\":\"bool\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"whitelistArray\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"whitelistMap\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"
=======
const WalletABI = "[{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_owner_\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_transferable_\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"_ens_\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_tokenWhitelistNode_\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_controllerNode_\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_licenceNode_\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_spendLimit_\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"_addresses\",\"type\":\"address[]\"}],\"name\":\"AddedToWhitelist\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"_assets\",\"type\":\"address[]\"}],\"name\":\"BulkTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"CancelledWhitelistAddition\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"CancelledWhitelistRemoval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_returndata\",\"type\":\"bytes\"}],\"name\":\"ExecutedRelayedTransaction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_destination\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_returndata\",\"type\":\"bytes\"}],\"name\":\"ExecutedTransaction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_currentNonce\",\"type\":\"uint256\"}],\"name\":\"IncreasedRelayNonce\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"LoadedTokenCard\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_locked\",\"type\":\"address\"}],\"name\":\"LockedOwnership\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"Received\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"_addresses\",\"type\":\"address[]\"}],\"name\":\"RemovedFromWhitelist\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"SetGasTopUpLimit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"SetLoadLimit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"SetSpendLimit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"SubmittedGasTopUpLimitUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"SubmittedLoadLimitUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"SubmittedSpendLimitUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"_addresses\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"SubmittedWhitelistAddition\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"_addresses\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"SubmittedWhitelistRemoval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"ToppedUpGas\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"Transferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"TransferredOwnership\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"UpdatedAvailableLimit\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"WALLET_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_transactionBatch\",\"type\":\"bytes\"}],\"name\":\"batchExecuteTransaction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"_assets\",\"type\":\"address[]\"}],\"name\":\"bulkTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_addresses\",\"type\":\"address[]\"}],\"name\":\"calculateHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"cancelWhitelistAddition\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"cancelWhitelistRemoval\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"confirmGasTopUpLimitUpdate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"confirmLoadLimitUpdate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"confirmSpendLimitUpdate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"confirmWhitelistAddition\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"confirmWhitelistRemoval\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"controllerNode\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"convertToEther\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"convertToStablecoin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ensRegistry\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"executeRelayedTransaction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_destination\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"executeTransaction\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasTopUpLimitAvailable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasTopUpLimitControllerConfirmationRequired\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasTopUpLimitPending\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasTopUpLimitValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"increaseRelayNonce\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isSetWhitelist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isTransferable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hashedData\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"isValidSignature\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"isValidSignature\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"licenceNode\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"loadLimitAvailable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"loadLimitControllerConfirmationRequired\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"loadLimitPending\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"loadLimitValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"loadTokenCard\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingWhitelistAddition\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingWhitelistRemoval\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"relayNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"setGasTopUpLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"setLoadLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"setSpendLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_addresses\",\"type\":\"address[]\"}],\"name\":\"setWhitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"spendLimitAvailable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"spendLimitControllerConfirmationRequired\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"spendLimitPending\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"spendLimitValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"submitGasTopUpLimitUpdate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"submitLoadLimitUpdate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"submitSpendLimitUpdate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_addresses\",\"type\":\"address[]\"}],\"name\":\"submitWhitelistAddition\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_addresses\",\"type\":\"address[]\"}],\"name\":\"submitWhitelistRemoval\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"submittedWhitelistAddition\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"submittedWhitelistRemoval\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_interfaceID\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenWhitelistNode\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"topUpGas\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_transferable\",\"type\":\"bool\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"whitelistArray\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"whitelistMap\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"
>>>>>>> 64ba3a32... Upgrade to solc 0.6.4 for Wallet

// WalletBin is the compiled bytecode used for deploying new contracts.
<<<<<<< HEAD
var WalletBin = "0x60806040527f7f2ce995617d2816b426c5c8698c5ec2952f7a34bb10f38326f74933d58936976001553480156200003557600080fd5b50604051620055a4380380620055a4833981810160405260e08110156200005b57600080fd5b508051602082015160408301516060840151608085015160a086015160c090960151600080546001600160a01b0319166001600160a01b03861617905594959394929391929091908084808989878015620000b65760018190555b50600280546001600160a01b0319166001600160a01b0384161760ff60a01b1916600160a01b8315158102919091179182905560ff9104166200013057604080516001600160a01b038416815290517f808639ff9c8e4732d60b6c2330de498035416d229f27a77d259680895efec1229181900360200190a15b60408051600081526001600160a01b038416602082015281517f850b3df64837d7d518b45f5aa64d104652c3b80eb5b34a8e3d9eb666cb7cdea5929181900390910190a150506002805460ff60b01b1916600160b01b1790556009556000620001a16001600160e01b036200027316565b505050505091505060008111620001ef576040805162461bcd60e51b815260206004820152600d60248201526c37379039ba30b13632b1b7b4b760991b604482015290519081900360640190fd5b808302600a819055600b819055600c8190556200021c4262015180620003f5602090811b620043a617901c565b600d81905560408051838152602081019290925280517fb8d7171194501073e2d8151eeccf1398143c5df9acfb9868d0539a256164f6ca9281900390910190a1505050600f92909255506200055195505050505050565b6060600080600080600080620002916009546200045760201b60201c565b6001600160a01b0316633efec5e96040518163ffffffff1660e01b815260040160006040518083038186803b158015620002ca57600080fd5b505afa158015620002df573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405260e08110156200030957600080fd5b81019080805160405193929190846401000000008211156200032a57600080fd5b9083019060208201858111156200034057600080fd5b82516401000000008111828201881017156200035b57600080fd5b82525081516020918201929091019080838360005b838110156200038a57818101518382015260200162000370565b50505050905090810190601f168015620003b85780820380516001836020036101000a031916815260200191505b5060409081526020820151908201516060830151608084015160a085015160c090950151979f939e50919c509a5098509096509294509192505050565b60008282018381101562000450576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b9392505050565b6000805460408051630178b8bf60e01b81526004810185905290516001600160a01b0390921691630178b8bf91602480820192602092909190829003018186803b158015620004a557600080fd5b505afa158015620004ba573d6000803e3d6000fd5b505050506040513d6020811015620004d157600080fd5b505160408051631d9dabef60e11b81526004810185905290516001600160a01b0390921691633b3b57de91602480820192602092909190829003018186803b1580156200051d57600080fd5b505afa15801562000532573d6000803e3d6000fd5b505050506040513d60208110156200054957600080fd5b505192915050565b61504380620005616000396000f3fe60806040526004361061027d5760003560e01c8063747c31d61161014f578063c1e559a3116100c1578063cf0a866b1161007a578063cf0a866b14610d94578063d251fefc14610da9578063de212bf314610dd3578063e2b4ce9714610de8578063f36febda14610dfd578063f421764814610e365761027d565b8063c1e559a314610b96578063cbd2ac6814610c68578063cccdc55614610c92578063cd7958dd14610ca7578063ce0b5bd514610d55578063ced99cce14610d7f5761027d565b80638da5cb5b116101135780638da5cb5b14610ac457806390e690c714610ad9578063ad95580b14610aee578063b242e53414610b03578063be40ba7914610b3e578063beabacc814610b535761027d565b8063747c31d6146109d75780637b580e75146109ec5780637d73b23114610a1f5780637fd004fa14610a34578063877337b014610aaf5761027d565b80633b8252fa116101f357806347d125af116101ac57806347d125af146108ad5780634d9aa248146108de5780635adc02ab146108f35780636137d6701461091d5780636c37a7e614610998578063715018a6146109c25761027d565b80633b8252fa146106d05780633f579f42146106f7578063458d07f21461078757806345b12efc146107b157806346efe0ed146107c657806347b55a9d146108985761027d565b806320c13b0b1161024557806320c13b0b146105175780632121dc75146105e257806326d05ab2146105f7578063294f40251461060c57806332531c3c146106715780633a43199f146106a45761027d565b806301ffc9a7146102b9578063100f23fd146103015780631127b57e1461032d5780631626ba7e146103b75780631aa21fba1461048c575b6040805133815234602082015281517f88a5966d370b9919b20f3e2c13ff65706f196a4e32cc2c12bf57088f88525874929181900390910190a1005b3480156102c557600080fd5b506102ed600480360360208110156102dc57600080fd5b50356001600160e01b031916610eb1565b604080519115158252519081900360200190f35b34801561030d57600080fd5b5061032b6004803603602081101561032457600080fd5b5035610ecb565b005b34801561033957600080fd5b5061034261106b565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561037c578181015183820152602001610364565b50505050905090810190601f1680156103a95780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156103c357600080fd5b5061046f600480360360408110156103da57600080fd5b81359190810190604081016020820135600160201b8111156103fb57600080fd5b82018360208201111561040d57600080fd5b803590602001918460018302840111600160201b8311171561042e57600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955061108c945050505050565b604080516001600160e01b03199092168252519081900360200190f35b34801561049857600080fd5b5061032b600480360360408110156104af57600080fd5b6001600160a01b038235169190810190604081016020820135600160201b8111156104d957600080fd5b8201836020820111156104eb57600080fd5b803590602001918460208302840111600160201b8311171561050c57600080fd5b5090925090506110fa565b34801561052357600080fd5b5061046f6004803603604081101561053a57600080fd5b810190602081018135600160201b81111561055457600080fd5b82018360208201111561056657600080fd5b803590602001918460018302840111600160201b8311171561058757600080fd5b919390929091602081019035600160201b8111156105a457600080fd5b8201836020820111156105b657600080fd5b803590602001918460018302840111600160201b831117156105d757600080fd5b509092509050611281565b3480156105ee57600080fd5b506102ed611356565b34801561060357600080fd5b506102ed611367565b34801561061857600080fd5b50610621611370565b60408051602080825283518183015283519192839290830191858101910280838360005b8381101561065d578181015183820152602001610645565b505050509050019250505060405180910390f35b34801561067d57600080fd5b506102ed6004803603602081101561069457600080fd5b50356001600160a01b03166113d2565b61032b600480360360408110156106ba57600080fd5b506001600160a01b0381351690602001356113e7565b3480156106dc57600080fd5b506106e5611630565b60408051918252519081900360200190f35b34801561070357600080fd5b506103426004803603606081101561071a57600080fd5b6001600160a01b0382351691602081013591810190606081016040820135600160201b81111561074957600080fd5b82018360208201111561075b57600080fd5b803590602001918460018302840111600160201b8311171561077c57600080fd5b509092509050611636565b34801561079357600080fd5b5061032b600480360360208110156107aa57600080fd5b50356116d2565b3480156107bd57600080fd5b506106e56117cb565b3480156107d257600080fd5b5061032b600480360360608110156107e957600080fd5b81359190810190604081016020820135600160201b81111561080a57600080fd5b82018360208201111561081c57600080fd5b803590602001918460018302840111600160201b8311171561083d57600080fd5b919390929091602081019035600160201b81111561085a57600080fd5b82018360208201111561086c57600080fd5b803590602001918460018302840111600160201b8311171561088d57600080fd5b5090925090506117e9565b3480156108a457600080fd5b506106216118be565b3480156108b957600080fd5b506108c261191e565b604080516001600160a01b039092168252519081900360200190f35b3480156108ea57600080fd5b506106e561192d565b3480156108ff57600080fd5b5061032b6004803603602081101561091657600080fd5b5035611933565b34801561092957600080fd5b5061032b6004803603602081101561094057600080fd5b810190602081018135600160201b81111561095a57600080fd5b82018360208201111561096c57600080fd5b803590602001918460208302840111600160201b8311171561098d57600080fd5b509092509050611c76565b3480156109a457600080fd5b5061032b600480360360208110156109bb57600080fd5b5035611e9d565b3480156109ce57600080fd5b5061032b61203f565b3480156109e357600080fd5b506106e561213a565b3480156109f857600080fd5b5061032b60048036036020811015610a0f57600080fd5b50356001600160a01b0316612140565b348015610a2b57600080fd5b506108c26122f9565b348015610a4057600080fd5b5061032b60048036036020811015610a5757600080fd5b810190602081018135600160201b811115610a7157600080fd5b820183602082011115610a8357600080fd5b803590602001918460208302840111600160201b83111715610aa457600080fd5b509092509050612308565b348015610abb57600080fd5b506106e561264b565b348015610ad057600080fd5b506108c2612651565b348015610ae557600080fd5b5061032b612660565b348015610afa57600080fd5b5061032b6126ba565b348015610b0f57600080fd5b5061032b60048036036040811015610b2657600080fd5b506001600160a01b03813516906020013515156127c1565b348015610b4a57600080fd5b506102ed612978565b348015610b5f57600080fd5b5061032b60048036036060811015610b7657600080fd5b506001600160a01b03813581169160208101359091169060400135612987565b348015610ba257600080fd5b5061032b60048036036060811015610bb957600080fd5b81359190810190604081016020820135600160201b811115610bda57600080fd5b820183602082011115610bec57600080fd5b803590602001918460018302840111600160201b83111715610c0d57600080fd5b919390929091602081019035600160201b811115610c2a57600080fd5b820183602082011115610c3c57600080fd5b803590602001918460018302840111600160201b83111715610c5d57600080fd5b509092509050612b13565b348015610c7457600080fd5b5061032b60048036036020811015610c8b57600080fd5b5035612c45565b348015610c9e57600080fd5b506106e5613035565b348015610cb357600080fd5b506106e560048036036020811015610cca57600080fd5b810190602081018135600160201b811115610ce457600080fd5b820183602082011115610cf657600080fd5b803590602001918460208302840111600160201b83111715610d1757600080fd5b91908080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525092955061303b945050505050565b348015610d6157600080fd5b5061032b60048036036020811015610d7857600080fd5b5035613095565b348015610d8b57600080fd5b506102ed613239565b348015610da057600080fd5b506102ed613249565b348015610db557600080fd5b506108c260048036036020811015610dcc57600080fd5b5035613259565b348015610ddf57600080fd5b506102ed613280565b348015610df457600080fd5b506106e561328e565b348015610e0957600080fd5b506106e560048036036040811015610e2057600080fd5b506001600160a01b038135169060200135613294565b348015610e4257600080fd5b5061032b60048036036020811015610e5957600080fd5b810190602081018135600160201b811115610e7357600080fd5b820183602082011115610e8557600080fd5b803590602001918460208302840111600160201b83111715610ea657600080fd5b50909250905061341d565b6001600160e01b031981166301ffc9a760e01b145b919050565b610ed433613770565b80610ee35750610ee333613784565b610f28576040805162461bcd60e51b81526020600482015260116024820152706f6e6c79206f776e6572206f722032464160781b604482015290519081900360640190fd5b60085460ff16610f77576040805162461bcd60e51b81526020600482015260156024820152743737903832b73234b7339039bab136b4b9b9b4b7b760591b604482015290519081900360640190fd5b610fda6006805480602002602001604051908101604052809291908181526020018280548015610fd057602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311610fb2575b505050505061303b565b81146110175760405162461bcd60e51b8152600401808060200182810382526023815260200180614fb66023913960400191505060405180910390fd5b61102360066000614ddc565b6008805460ff19169055604080513381526020810183905281517f7794eff834d760583543e6e510e717a5e66d2c064e225f4db448343c3e66afcf929181900390910190a150565b60405180604001604052806005815260200164332e312e3160d81b81525081565b60008061109f848463ffffffff6137c016565b90506110aa81613770565b6110e8576040805162461bcd60e51b815260206004820152600a60248201526937b7363c9037bbb732b960b11b604482015290519081900360640190fd5b50630b135d3f60e11b90505b92915050565b61110333613770565b8061110d57503330145b611152576040805162461bcd60e51b81526020600482015260116024820152702737ba1037bbb732b91037b91039b2b63360791b604482015290519081900360640190fd5b8061119b576040805162461bcd60e51b8152602060048201526014602482015273617373657420617272617920697320656d70747960601b604482015290519081900360640190fd5b60005b818110156111fe5760006111cd308585858181106111b857fe5b905060200201356001600160a01b03166138ae565b90506111f5858585858181106111df57fe5b905060200201356001600160a01b031683612987565b5060010161119e565b507fd4f62f23021706247dcffea245d104ae7ddaec7f23acf3d11d7136d5de6a69ad83838360405180846001600160a01b03166001600160a01b03168152602001806020018281038252848482818152602001925060200280828437600083820152604051601f909101601f1916909201829003965090945050505050a1505050565b6000808585604051602001808383808284376040805191909301818103601f190182528084528151602092830120601f8b01839004830282018301909452898152929650630b135d3f60e11b95506112f8945086935089915088908190840183828082843760009201919091525061108c92505050565b6001600160e01b03191614611344576040805162461bcd60e51b815260206004820152600d60248201526c1cda59c81b9bdd081d985b1a59609a1b604482015290519081900360640190fd5b506320c13b0b60e01b95945050505050565b600254600160a01b900460ff165b90565b60085460ff1681565b606060078054806020026020016040519081016040528092919081815260200182805480156113c857602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116113aa575b5050505050905090565b60046020526000908152604090205460ff1681565b6113f033613770565b806113fa57503330145b61143f576040805162461bcd60e51b81526020600482015260116024820152702737ba1037bbb732b91037b91039b2b63360791b604482015290519081900360640190fd5b61144882613959565b61148e576040805162461bcd60e51b8152602060048201526012602482015271746f6b656e206e6f74206c6f616461626c6560701b604482015290519081900360640190fd5b600254600160a81b900460ff166114b85760006114ab8383613294565b90506114b681613973565b505b60006114c5600f54613a1e565b90506001600160a01b0383161561156d576114f06001600160a01b038416828463ffffffff613b1216565b806001600160a01b0316631b3c96b484846040518363ffffffff1660e01b815260040180836001600160a01b03166001600160a01b0316815260200182815260200192505050600060405180830381600087803b15801561155057600080fd5b505af1158015611564573d6000803e3d6000fd5b505050506115e7565b806001600160a01b0316631b3c96b48385856040518463ffffffff1660e01b815260040180836001600160a01b03166001600160a01b03168152602001828152602001925050506000604051808303818588803b1580156115cd57600080fd5b505af11580156115e1573d6000803e3d6000fd5b50505050505b604080516001600160a01b03851681526020810184905281517f5f65674bec9af81f71be68674135a0ea3f163fb91984e3893d06da9f6ea2ce8a929181900390910190a1505050565b600c5490565b606061164133613770565b611688576040805162461bcd60e51b815260206004820152601360248201527239b2b73232b91034b9903737ba1037bbb732b960691b604482015290519081900360640190fd5b6116c9858585858080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250613c2a92505050565b50949350505050565b6116db33613770565b806116e557503330145b61172a576040805162461bcd60e51b81526020600482015260116024820152702737ba1037bbb732b91037b91039b2b63360791b604482015290519081900360640190fd5b600c819055600a54811161179457600b5481101561178657600b819055600d5460408051838152602081019290925280517fb8d7171194501073e2d8151eeccf1398143c5df9acfb9868d0539a256164f6ca9281900390910190a15b61178f816140eb565b6117c8565b6040805182815290517f065b9ade648867cf901516060dd7a78fad8ab1aec5eb80ee57acbb30badf86ec9181900360200190a15b50565b6000600d544211156117e05750600a54611364565b50600b54611364565b6117f23361412e565b611843576040805162461bcd60e51b815260206004820152601a60248201527f73656e646572206973206e6f74206120636f6e74726f6c6c6572000000000000604482015290519081900360640190fd5b6118b78585858080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050604080516020601f8901819004810282018101909252878152925087915086908190840183828082843760009201829052509250614190915050565b5050505050565b606060068054806020026020016040519081016040528092919081815260200182805480156113c8576020028201919060005260206000209081546001600160a01b031681526001909101906020018083116113aa575050505050905090565b6003546001600160a01b031681565b600a5490565b600254600160b01b900460ff16156119925761194e3361412e565b61198d576040805162461bcd60e51b815260206004820152601c6024820152600080516020614ee4833981519152604482015290519081900360640190fd5b6119ee565b6003546001600160a01b031633146119ee576040805162461bcd60e51b815260206004820152601a60248201527973656e646572206973206e6f7420706572736f6e616c2032464160301b604482015290519081900360640190fd5b60085460ff16611a3d576040805162461bcd60e51b81526020600482015260156024820152743737903832b73234b7339039bab136b4b9b9b4b7b760591b604482015290519081900360640190fd5b611a9e6006805480602002602001604051908101604052809291908181526020018280548015610fd0576020028201919060005260206000209081546001600160a01b03168152600190910190602001808311610fb257505050505061303b565b8114611adb5760405162461bcd60e51b8152600401808060200182810382526023815260200180614fb66023913960400191505060405180910390fd5b60005b600654811015611bc2576004600060068381548110611af957fe5b60009182526020808320909101546001600160a01b0316835282019290925260400190205460ff16611bba5760016004600060068481548110611b3857fe5b6000918252602080832091909101546001600160a01b031683528201929092526040019020805460ff1916911515919091179055600680546005919083908110611b7e57fe5b60009182526020808320909101548354600181018555938352912090910180546001600160a01b0319166001600160a01b039092169190911790555b600101611ade565b507fb2f6cccee7a369e23e293c25aa19bef80af11eb26deba3ea0f2a02783f752e4a33600660405180836001600160a01b03166001600160a01b03168152602001806020018281038252838181548152602001915080548015611c4e57602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311611c30575b5050935050505060405180910390a1611c6960066000614ddc565b506008805460ff19169055565b611c7f33613770565b80611c8957503330145b611cce576040805162461bcd60e51b81526020600482015260116024820152702737ba1037bbb732b91037b91039b2b63360791b604482015290519081900360640190fd5b60085460ff16158015611ce95750600854610100900460ff16155b611d3a576040805162461bcd60e51b815260206004820152601c60248201527f77686974656c6973742073756d62697373696f6e2070656e64696e6700000000604482015290519081900360640190fd5b60085462010000900460ff16611d93576040805162461bcd60e51b81526020600482015260196024820152781dda1a5d195b1a5cdd081b9bdd081a5b9a5d1a585b1a5e9959603a1b604482015290519081900360640190fd5b80611dd7576040805162461bcd60e51b815260206004820152600f60248201526e195b5c1d1e481dda1a5d195b1a5cdd608a1b604482015290519081900360640190fd5b611de360078383614dfa565b506008805461ff00191661010017905560408051602080840282810182019093528382527ffbc0e5ca6c7e4858daf0fdb185ef5186203e74ec9c64737e93c0aeaec596e1d19285928592611e529285918591829185019084908082843760009201919091525061303b92505050565b60405180806020018381526020018281038252858582818152602001925060200280828437600083820152604051601f909101601f1916909201829003965090945050505050a15050565b600254600160b01b900460ff1615611efc57611eb83361412e565b611ef7576040805162461bcd60e51b815260206004820152601c6024820152600080516020614ee4833981519152604482015290519081900360640190fd5b611f58565b6003546001600160a01b03163314611f58576040805162461bcd60e51b815260206004820152601a60248201527973656e646572206973206e6f7420706572736f6e616c2032464160301b604482015290519081900360640190fd5b80600c5414611f985760405162461bcd60e51b8152600401808060200182810382526022815260200180614f046022913960400191505060405180910390fd5b600a548111611fd85760405162461bcd60e51b8152600401808060200182810382526028815260200180614ebc6028913960400191505060405180910390fd5b600b819055611ff0426201518063ffffffff6143a616565b600d819055600b5460408051918252602082019290925281517fb8d7171194501073e2d8151eeccf1398143c5df9acfb9868d0539a256164f6ca929181900390910190a16117c8600c546140eb565b61204833613770565b61208f576040805162461bcd60e51b815260206004820152601360248201527239b2b73232b91034b9903737ba1037bbb732b960691b604482015290519081900360640190fd5b600254600160a01b900460ff166120ed576040805162461bcd60e51b815260206004820152601d60248201527f6f776e657273686970206973206e6f74207472616e7366657261626c65000000604482015290519081900360640190fd5b600280546001600160a01b0319169055604080516000808252602082015281517f850b3df64837d7d518b45f5aa64d104652c3b80eb5b34a8e3d9eb666cb7cdea5929181900390910190a1565b600f5490565b61214933613770565b8061215357503330145b61218f576040805162461bcd60e51b81526020600482015260086024820152673737ba1039b2b63360c11b604482015290519081900360640190fd5b600254600160a81b900460ff166121ed576040805162461bcd60e51b815260206004820152601d60248201527f53657420324641206e656564732070726976696c65676564206d6f6465000000604482015290519081900360640190fd5b6001600160a01b038116612248576040805162461bcd60e51b815260206004820152601960248201527f3246412063616e6e6f742062652073657420746f207a65726f00000000000000604482015290519081900360640190fd5b6001600160a01b0381163014156122905760405162461bcd60e51b8152600401808060200182810382526022815260200180614f476022913960400191505060405180910390fd5b600380546001600160a01b0319166001600160a01b0383169081179091556002805460ff60b01b1916905560408051338152602081019290925280517f33635a1d7938fa110d60d48b4ecbefc4afcc07e782ca013f11de948ee7949d1c9281900390910190a150565b6000546001600160a01b031690565b61231133613770565b8061231b57503330145b612360576040805162461bcd60e51b81526020600482015260116024820152702737ba1037bbb732b91037b91039b2b63360791b604482015290519081900360640190fd5b60085460ff1615801561237b5750600854610100900460ff16155b6123cc576040805162461bcd60e51b815260206004820152601c60248201527f77686974656c6973742073756d62697373696f6e2070656e64696e6700000000604482015290519081900360640190fd5b8181808060200260200160405190810160405280939291908181526020018383602002808284376000920182905250925050505b81518110156124e85761242582828151811061241857fe5b6020026020010151613770565b15612470576040805162461bcd60e51b8152602060048201526016602482015275636f6e7461696e73206f776e6572206164647265737360501b604482015290519081900360640190fd5b60006001600160a01b031682828151811061248757fe5b60200260200101516001600160a01b031614156124e0576040805162461bcd60e51b8152602060048201526012602482015271636f6e7461696e732030206164647265737360701b604482015290519081900360640190fd5b600101612400565b5060085462010000900460ff16612542576040805162461bcd60e51b81526020600482015260196024820152781dda1a5d195b1a5cdd081b9bdd081a5b9a5d1a585b1a5e9959603a1b604482015290519081900360640190fd5b81612586576040805162461bcd60e51b815260206004820152600f60248201526e195b5c1d1e481dda1a5d195b1a5cdd608a1b604482015290519081900360640190fd5b61259260068484614dfa565b506008805460ff1916600117905560408051602080850282810182019093528482527f9c80b3b5f68b3e017766d59e8d09b34efe6462b05c398f35cab9e271d9bc3b9c92869286926125ff9285918591829185019084908082843760009201919091525061303b92505050565b60405180806020018381526020018281038252858582818152602001925060200280828437600083820152604051601f909101601f1916909201829003965090945050505050a1505050565b60095490565b6002546001600160a01b031690565b61266933613770565b6126b0576040805162461bcd60e51b815260206004820152601360248201527239b2b73232b91034b9903737ba1037bbb732b960691b604482015290519081900360640190fd5b6126b8614400565b565b6126c333613770565b61270a576040805162461bcd60e51b815260206004820152601360248201527239b2b73232b91034b9903737ba1037bbb732b960691b604482015290519081900360640190fd5b600254600160b01b900460ff1615612769576040805162461bcd60e51b815260206004820152601b60248201527f6d6f6e6f6c69746832464120616c726561647920656e61626c65640000000000604482015290519081900360640190fd5b6002805460ff60b01b1916600160b01b179055600380546001600160a01b03191690556040805133815290517fea0cce48757f6b222f91e711f59a5a8ec05a3ed3c0a4328fe36cb48f31a869699181900360200190a1565b6127ca33613770565b612811576040805162461bcd60e51b815260206004820152601360248201527239b2b73232b91034b9903737ba1037bbb732b960691b604482015290519081900360640190fd5b600254600160a01b900460ff1661286f576040805162461bcd60e51b815260206004820152601d60248201527f6f776e657273686970206973206e6f74207472616e7366657261626c65000000604482015290519081900360640190fd5b6001600160a01b0382166128b45760405162461bcd60e51b8152600401808060200182810382526023815260200180614f696023913960400191505060405180910390fd5b6002805460ff60a01b1916600160a01b831515021790558061290d57604080516001600160a01b038416815290517f808639ff9c8e4732d60b6c2330de498035416d229f27a77d259680895efec1229181900360200190a15b600254604080516001600160a01b039283168152918416602083015280517f850b3df64837d7d518b45f5aa64d104652c3b80eb5b34a8e3d9eb666cb7cdea59281900390910190a150600280546001600160a01b0319166001600160a01b0392909216919091179055565b60085462010000900460ff1681565b61299033613770565b8061299a57503330145b6129df576040805162461bcd60e51b81526020600482015260116024820152702737ba1037bbb732b91037b91039b2b63360791b604482015290519081900360640190fd5b8080612a1c576040805162461bcd60e51b8152602060048201526007602482015266076616c75653d360cc1b604482015290519081900360640190fd5b6001600160a01b038416612a67576040805162461bcd60e51b815260206004820152600d60248201526c064657374696e6174696f6e3d3609c1b604482015290519081900360640190fd5b6001600160a01b03841660009081526004602052604090205460ff16158015612a9a5750600254600160a81b900460ff16155b15612ab8576000612aab8484613294565b9050612ab681613973565b505b612ac3848484614448565b604080516001600160a01b0380871682528516602082015280820184905290517fd1ba4ac2e2a11b5101f6cb4d978f514a155b421e8ec396d2d9abaf0bb02917ee9181900360600190a150505050565b600254600160b01b900460ff1615612b7257612b2e3361412e565b612b6d576040805162461bcd60e51b815260206004820152601c6024820152600080516020614ee4833981519152604482015290519081900360640190fd5b612bce565b6003546001600160a01b03163314612bce576040805162461bcd60e51b815260206004820152601a60248201527973656e646572206973206e6f7420706572736f6e616c2032464160301b604482015290519081900360640190fd5b6118b78585858080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050604080516020601f8901819004810282018101909252878152925087915086908190840183828082843760009201919091525060019250614190915050565b600254600160b01b900460ff1615612ca457612c603361412e565b612c9f576040805162461bcd60e51b815260206004820152601c6024820152600080516020614ee4833981519152604482015290519081900360640190fd5b612d00565b6003546001600160a01b03163314612d00576040805162461bcd60e51b815260206004820152601a60248201527973656e646572206973206e6f7420706572736f6e616c2032464160301b604482015290519081900360640190fd5b600854610100900460ff16612d54576040805162461bcd60e51b81526020600482015260156024820152743737903832b73234b7339039bab136b4b9b9b4b7b760591b604482015290519081900360640190fd5b612db56007805480602002602001604051908101604052809291908181526020018280548015610fd0576020028201919060005260206000209081546001600160a01b03168152600190910190602001808311610fb257505050505061303b565b8114612df25760405162461bcd60e51b8152600401808060200182810382526023815260200180614fb66023913960400191505060405180910390fd5b60005b600754811015612f80576004600060078381548110612e1057fe5b60009182526020808320909101546001600160a01b0316835282019290925260400190205460ff1615612f785760006004600060078481548110612e5057fe5b6000918252602080832091909101546001600160a01b031683528201929092526040018120805460ff1916921515929092179091555b600554612e9a90600163ffffffff6144ac16565b811015612f625760078281548110612eae57fe5b600091825260209091200154600580546001600160a01b039092169183908110612ed457fe5b6000918252602090912001546001600160a01b03161415612f5a57600580546000198101908110612f0157fe5b600091825260209091200154600580546001600160a01b039092169183908110612f2757fe5b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b03160217905550612f62565b600101612e86565b506005805490612f76906000198301614e5d565b505b600101612df5565b507fd218c430fa348f4ce67791021b6b89c0c3eacd4ead1d8f5b83c60038ec28249b33600760405180836001600160a01b03166001600160a01b0316815260200180602001828103825283818154815260200191508054801561300c57602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311612fee575b5050935050505060405180910390a161302760076000614ddc565b506008805461ff0019169055565b600e5481565b60008160405160200180828051906020019060200280838360005b8381101561306e578181015183820152602001613056565b50505050905001915050604051602081830303815290604052805190602001209050919050565b61309e33613770565b806130ad57506130ad33613784565b6130f2576040805162461bcd60e51b81526020600482015260116024820152706f6e6c79206f776e6572206f722032464160781b604482015290519081900360640190fd5b600854610100900460ff16613146576040805162461bcd60e51b81526020600482015260156024820152743737903832b73234b7339039bab136b4b9b9b4b7b760591b604482015290519081900360640190fd5b6131a76007805480602002602001604051908101604052809291908181526020018280548015610fd0576020028201919060005260206000209081546001600160a01b03168152600190910190602001808311610fb257505050505061303b565b81146131e45760405162461bcd60e51b8152600401808060200182810382526023815260200180614fb66023913960400191505060405180910390fd5b6131f060076000614ddc565b6008805461ff0019169055604080513381526020810183905281517f13c935eb475aa0f6e931fece83e2ac44569ce2d53460d29a6dedab40b965c8a3929181900390910190a150565b600254600160a81b900460ff1681565b600254600160b01b900460ff1681565b6005818154811061326657fe5b6000918252602090912001546001600160a01b0316905081565b600854610100900460ff1681565b60015490565b600061329e614509565b6001600160a01b0316836001600160a01b031614156132be5750806110f4565b816001600160a01b038416156133595760008060006132dc8761457f565b50505093509350935050806132f85760009450505050506110f4565b81613333576040805162461bcd60e51b81526020600482015260066024820152650726174653d360d41b604482015290519081900360640190fd5b61335383613347888563ffffffff61471116565b9063ffffffff61476a16565b93505050505b60008060006133666147d4565b50505093509350935050806133b8576040805162461bcd60e51b8152602060048201526013602482015272746f6b656e206e6f7420617661696c61626c6560681b604482015290519081900360640190fd5b816133fe576040805162461bcd60e51b81526020600482015260116024820152700737461626c65636f696e20726174653d3607c1b604482015290519081900360640190fd5b61341282613347868663ffffffff61471116565b979650505050505050565b61342633613770565b8061343057503330145b613475576040805162461bcd60e51b81526020600482015260116024820152702737ba1037bbb732b91037b91039b2b63360791b604482015290519081900360640190fd5b8181808060200260200160405190810160405280939291908181526020018383602002808284376000920182905250925050505b8151811015613584576134c182828151811061241857fe5b1561350c576040805162461bcd60e51b8152602060048201526016602482015275636f6e7461696e73206f776e6572206164647265737360501b604482015290519081900360640190fd5b60006001600160a01b031682828151811061352357fe5b60200260200101516001600160a01b0316141561357c576040805162461bcd60e51b8152602060048201526012602482015271636f6e7461696e732030206164647265737360701b604482015290519081900360640190fd5b6001016134a9565b5060085462010000900460ff16156135db576040805162461bcd60e51b81526020600482015260156024820152741dda1a5d195b1a5cdd081a5b9a5d1a585b1a5e9959605a1b604482015290519081900360640190fd5b60005b828110156136cc57600460008585848181106135f657fe5b602090810292909201356001600160a01b03168352508101919091526040016000205460ff166136c45760016004600086868581811061363257fe5b905060200201356001600160a01b03166001600160a01b03166001600160a01b0316815260200190815260200160002060006101000a81548160ff021916908315150217905550600584848381811061368757fe5b835460018101855560009485526020948590200180546001600160a01b0319166001600160a01b0395909202939093013593909316929092179055505b6001016135de565b506008805462ff0000191662010000179055604080513380825260208201838152600580549484018590527fb2f6cccee7a369e23e293c25aa19bef80af11eb26deba3ea0f2a02783f752e4a949293909290919060608301908490801561375c57602002820191906000526020600020905b81546001600160a01b0316815260019091019060200180831161373e575b5050935050505060405180910390a1505050565b6002546001600160a01b0390811691161490565b600254600090600160b01b900460ff16156137a9576137a28261412e565b9050610ec6565b506003546001600160a01b03828116911614610ec6565b600081516041146137d3575060006110f4565b60208201516040830151606084015160001a7f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a082111561381957600093505050506110f4565b8060ff16601b1415801561383157508060ff16601c14155b1561384257600093505050506110f4565b6040805160008152602080820180845289905260ff8416828401526060820186905260808201859052915160019260a0808401939192601f1981019281900390910190855afa158015613899573d6000803e3d6000fd5b5050604051601f190151979650505050505050565b60006001600160a01b0382161561394857816001600160a01b03166370a08231846040518263ffffffff1660e01b815260040180826001600160a01b03166001600160a01b0316815260200191505060206040518083038186803b15801561391557600080fd5b505afa158015613929573d6000803e3d6000fd5b505050506040513d602081101561393f57600080fd5b505190506110f4565b506001600160a01b038216316110f4565b6000806139658361457f565b509098975050505050505050565b61397b614943565b80600b5410156139c5576040805162461bcd60e51b815260206004820152601060248201526f185d985a5b18589b194f185b5bdd5b9d60821b604482015290519081900360640190fd5b600b546139d8908263ffffffff6144ac16565b600b819055600d5460408051928352602083019190915280517fb8d7171194501073e2d8151eeccf1398143c5df9acfb9868d0539a256164f6ca9281900390910190a150565b6000805460408051630178b8bf60e01b81526004810185905290516001600160a01b0390921691630178b8bf91602480820192602092909190829003018186803b158015613a6b57600080fd5b505afa158015613a7f573d6000803e3d6000fd5b505050506040513d6020811015613a9557600080fd5b505160408051631d9dabef60e11b81526004810185905290516001600160a01b0390921691633b3b57de91602480820192602092909190829003018186803b158015613ae057600080fd5b505afa158015613af4573d6000803e3d6000fd5b505050506040513d6020811015613b0a57600080fd5b505192915050565b801580613b98575060408051636eb1769f60e11b81523060048201526001600160a01b03848116602483015291519185169163dd62ed3e91604480820192602092909190829003018186803b158015613b6a57600080fd5b505afa158015613b7e573d6000803e3d6000fd5b505050506040513d6020811015613b9457600080fd5b5051155b613bd35760405162461bcd60e51b8152600401808060200182810382526036815260200180614fd96036913960400191505060405180910390fd5b604080516001600160a01b038416602482015260448082018490528251808303909101815260649091019091526020810180516001600160e01b031663095ea7b360e01b179052613c2590849061496b565b505050565b6001600160a01b03831660009081526004602052604090205460609060ff16158015613c605750600254600160a81b900460ff16155b15613c6e57613c6e83613973565b613c80846001600160a01b0316614b29565b8015613c905750613c9084614b2f565b15613e8757600080613ca28685614b49565b6001600160a01b038216600090815260046020526040902054919350915060ff16158015613cda5750600254600160a81b900460ff16155b15613cf8576000613ceb8783613294565b9050613cf681613973565b505b613d116001600160a01b0387168563ffffffff61496b16565b604080516020808252818301909252606091602082018180388339019050509050600160f81b81601f81518110613d4457fe5b60200101906001600160f81b031916908160001a9053507ff77753fab406ecfff96d6ff2476c64a838fa9f6d37b1bf190f8546e395e3b6138787878460405180856001600160a01b03166001600160a01b031681526020018481526020018060200180602001838103835285818151815260200191508051906020019080838360005b83811015613ddf578181015183820152602001613dc7565b50505050905090810190601f168015613e0c5780820380516001836020036101000a031916815260200191505b50838103825284518152845160209182019186019080838360005b83811015613e3f578181015183820152602001613e27565b50505050905090810190601f168015613e6c5780820380516001836020036101000a031916815260200191505b50965050505050505060405180910390a192506140e4915050565b60006060856001600160a01b031685856040518082805190602001908083835b60208310613ec65780518252601f199092019160209182019101613ea7565b6001836020036101000a03801982511681845116808217855250505050505090500191505060006040518083038185875af1925050503d8060008114613f28576040519150601f19603f3d011682016040523d82523d6000602084013e613f2d565b606091505b5091509150818190613fbd5760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b83811015613f82578181015183820152602001613f6a565b50505050905090810190601f168015613faf5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b507ff77753fab406ecfff96d6ff2476c64a838fa9f6d37b1bf190f8546e395e3b6138686868460405180856001600160a01b03166001600160a01b031681526020018481526020018060200180602001838103835285818151815260200191508051906020019080838360005b8381101561404257818101518382015260200161402a565b50505050905090810190601f16801561406f5780820380516001836020036101000a031916815260200191505b50838103825284518152845160209182019186019080838360005b838110156140a257818101518382015260200161408a565b50505050905090810190601f1680156140cf5780820380516001836020036101000a031916815260200191505b50965050505050505060405180910390a19150505b9392505050565b600a819055604080513381526020810183905281517f2a843f39f13315c4c1a9bc53a1a32162858f272f3b2d0c656f409431251b6768929181900390910190a150565b600061413b600154613a1e565b6001600160a01b031663b429afeb836040518263ffffffff1660e01b815260040180826001600160a01b03166001600160a01b0316815260200191505060206040518083038186803b158015613ae057600080fd5b600061421f858560405160200180806339363c1d60e11b81525060040183815260200182805190602001908083835b602083106141de5780518252601f1990920191602091820191016141bf565b6001836020036101000a0380198251168184511680821785525050505050509050019250505060405160208183030381529060405280519060200120614c53565b9050630b135d3f60e11b614233828561108c565b6001600160e01b0319161461427f576040805162461bcd60e51b815260206004820152600d60248201526c1cda59c81b9bdd081d985b1a59609a1b604482015290519081900360640190fd5b600e5485146142c1576040805162461bcd60e51b81526020600482015260096024820152687478207265706c617960b81b604482015290519081900360640190fd5b6142c9614400565b6002805460ff60a81b1916600160a81b841515021790556142e984614ca4565b6002805460ff60a81b191690556040805183151560208083019190915282825286519282019290925285517f1e67acdbe17d73f10c3c1cec8dba9c0bca6d8fcd7f326d2a00e6520026b21585928792869290918291606083019186019080838360005b8381101561436457818101518382015260200161434c565b50505050905090810190601f1680156143915780820380516001836020036101000a031916815260200191505b50935050505060405180910390a15050505050565b6000828201838110156140e4576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b600e80546001019081905560408051338152602081019290925280517fab0423a75986556234aecd171c46ce7f5e45607d8070bf5230f2735b50322bff9281900390910190a1565b6001600160a01b038216614492576040516001600160a01b0384169082156108fc029083906000818181858888f1935050505015801561448c573d6000803e3d6000fd5b50613c25565b613c256001600160a01b038316848363ffffffff614d8a16565b600082821115614503576040805162461bcd60e51b815260206004820152601e60248201527f536166654d6174683a207375627472616374696f6e206f766572666c6f770000604482015290519081900360640190fd5b50900390565b6000614516600954613a1e565b6001600160a01b031663e9cbd8226040518163ffffffff1660e01b815260040160206040518083038186803b15801561454e57600080fd5b505afa158015614562573d6000803e3d6000fd5b505050506040513d602081101561457857600080fd5b5051905090565b6060600080600080600080614595600954613a1e565b6001600160a01b0316631f69565f896040518263ffffffff1660e01b815260040180826001600160a01b03166001600160a01b0316815260200191505060006040518083038186803b1580156145ea57600080fd5b505afa1580156145fe573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405260e081101561462757600080fd5b8101908080516040519392919084600160201b82111561464657600080fd5b90830190602082018581111561465b57600080fd5b8251600160201b81118282018810171561467457600080fd5b82525081516020918201929091019080838360005b838110156146a1578181015183820152602001614689565b50505050905090810190601f1680156146ce5780820380516001836020036101000a031916815260200191505b5060409081526020820151908201516060830151608084015160a085015160c090950151979e50929c50909a509850965094509192505050919395979092949650565b600082614720575060006110f4565b8282028284828161472d57fe5b04146140e45760405162461bcd60e51b8152600401808060200182810382526021815260200180614f266021913960400191505060405180910390fd5b60008082116147c0576040805162461bcd60e51b815260206004820152601a60248201527f536166654d6174683a206469766973696f6e206279207a65726f000000000000604482015290519081900360640190fd5b60008284816147cb57fe5b04949350505050565b60606000806000806000806147ea600954613a1e565b6001600160a01b0316633efec5e96040518163ffffffff1660e01b815260040160006040518083038186803b15801561482257600080fd5b505afa158015614836573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405260e081101561485f57600080fd5b8101908080516040519392919084600160201b82111561487e57600080fd5b90830190602082018581111561489357600080fd5b8251600160201b8111828201881017156148ac57600080fd5b82525081516020918201929091019080838360005b838110156148d95781810151838201526020016148c1565b50505050905090810190601f1680156149065780820380516001836020036101000a031916815260200191505b5060409081526020820151908201516060830151608084015160a085015160c090950151979f939e50919c509a5098509096509294509192505050565b600d544211156126b857614960426201518063ffffffff6143a616565b600d55600a54600b55565b61497d826001600160a01b0316614b29565b6149ce576040805162461bcd60e51b815260206004820152601f60248201527f5361666545524332303a2063616c6c20746f206e6f6e2d636f6e747261637400604482015290519081900360640190fd5b60006060836001600160a01b0316836040518082805190602001908083835b60208310614a0c5780518252601f1990920191602091820191016149ed565b6001836020036101000a0380198251168184511680821785525050505050509050019150506000604051808303816000865af19150503d8060008114614a6e576040519150601f19603f3d011682016040523d82523d6000602084013e614a73565b606091505b509150915081614aca576040805162461bcd60e51b815260206004820181905260248201527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564604482015290519081900360640190fd5b805115614b2357808060200190516020811015614ae657600080fd5b5051614b235760405162461bcd60e51b815260040180806020018281038252602a815260200180614f8c602a913960400191505060405180910390fd5b50505050565b3b151590565b600080614b3b8361457f565b509198975050505050505050565b600080614b57600954613a1e565b6001600160a01b031663afc72e9385856040518363ffffffff1660e01b815260040180836001600160a01b03166001600160a01b0316815260200180602001828103825283818151815260200191508051906020019080838360005b83811015614bcb578181015183820152602001614bb3565b50505050905090810190601f168015614bf85780820380516001836020036101000a031916815260200191505b509350505050604080518083038186803b158015614c1557600080fd5b505afa158015614c29573d6000803e3d6000fd5b505050506040513d6040811015614c3f57600080fd5b508051602090910151909590945092505050565b604080517f19457468657265756d205369676e6564204d6573736167653a0a333200000000602080830191909152603c8083019490945282518083039094018452605c909101909152815191012090565b8051602080820191906000808060605b86851015614d8057614ccd86605463ffffffff6144ac16565b888601805160148201516034909201805193995060609190911c96509094509092509050614d126054614d06878563ffffffff6143a616565b9063ffffffff6143a616565b945086851115614d59576040805162461bcd60e51b815260206004820152600d60248201526c6f7574206f6620626f756e647360981b604482015290519081900360640190fd5b81614d6f57506040805160208101909152600081525b614d7a848483613c2a565b50614cb4565b5050505050505050565b604080516001600160a01b038416602482015260448082018490528251808303909101815260649091019091526020810180516001600160e01b031663a9059cbb60e01b179052613c2590849061496b565b50805460008255906000526020600020908101906117c89190614e7d565b828054828255906000526020600020908101928215614e4d579160200282015b82811115614e4d5781546001600160a01b0319166001600160a01b03843516178255602090920191600190910190614e1a565b50614e59929150614e97565b5090565b815481835581811115613c2557600083815260209020613c259181019083015b61136491905b80821115614e595760008155600101614e83565b61136491905b80821115614e595780546001600160a01b0319168155600101614e9d56fe6c696d69742073686f756c642062652067726561746572207468616e2063757272656e74206f6e6573656e646572206973206e6f742061204d6f6e6f6c6974682032464100000000636f6e6669726d65642f7375626d6974746564206c696d6974206d69736d61746368536166654d6174683a206d756c7469706c69636174696f6e206f766572666c6f773246412063616e6e6f742062652074686520636f6e747261637420616464726573736f776e65722063616e6e6f742062652073657420746f207a65726f20616464726573735361666545524332303a204552433230206f7065726174696f6e20646964206e6f7420737563636565646e6f6e2d6d61746368696e672070656e64696e672077686974656c69737420686173685361666545524332303a20617070726f76652066726f6d206e6f6e2d7a65726f20746f206e6f6e2d7a65726f20616c6c6f77616e6365a265627a7a72315820692db7efb16913f05b684c3547ffa7664fa402ea0829aa03c31f1824b2ac946564736f6c63430005110032"
||||||| parent of 64ba3a32... Upgrade to solc 0.6.4 for Wallet
var WalletBin = "0x60806040527f7f2ce995617d2816b426c5c8698c5ec2952f7a34bb10f38326f74933d58936976001553480156200003557600080fd5b5060405162005a3338038062005a33833981810160405260e08110156200005b57600080fd5b508051602082015160408301516060840151608085015160a086015160c090960151600080546001600160a01b0319166001600160a01b03861617905594959394929391929091908084808989878015620000b65760018190555b50600280546001600160a01b0319166001600160a01b0384161760ff60a01b1916600160a01b8315158102919091179182905560ff9104166200013057604080516001600160a01b038416815290517f808639ff9c8e4732d60b6c2330de498035416d229f27a77d259680895efec1229181900360200190a15b60408051600081526001600160a01b038416602082015281517f850b3df64837d7d518b45f5aa64d104652c3b80eb5b34a8e3d9eb666cb7cdea5929181900390910190a150506040805160a0810182526706f05b59d3b2000080825260208201819052429282018390526000606083018190526080909201829052600381905560045560059190915560068190556007805460ff19169055600891909155620001e16001600160e01b03620002d816565b5050505050915050600081116200022f576040805162461bcd60e51b815260206004820152600d60248201526c37379039ba30b13632b1b7b4b760991b604482015290519081900360640190fd5b6127100260098190556040805160a08082018352838252602080830185905242838501819052600060608086018290526080958601829052600a889055600b97909755600c829055600d819055600e805460ff1990811690915586519485018752898552928401899052948301819052948201849052910182905260148590556015949094556016919091556017555060188054909116905550601a5550620005549350505050565b6060600080600080600080620002f66008546200045a60201b60201c565b6001600160a01b0316633efec5e96040518163ffffffff1660e01b815260040160006040518083038186803b1580156200032f57600080fd5b505afa15801562000344573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405260e08110156200036e57600080fd5b81019080805160405193929190846401000000008211156200038f57600080fd5b908301906020820185811115620003a557600080fd5b8251640100000000811182820188101715620003c057600080fd5b82525081516020918201929091019080838360005b83811015620003ef578181015183820152602001620003d5565b50505050905090810190601f1680156200041d5780820380516001836020036101000a031916815260200191505b5060409081526020820151908201516060830151608084015160a085015160c090950151979f939e50919c509a5098509096509294509192505050565b6000805460408051630178b8bf60e01b81526004810185905290516001600160a01b0390921691630178b8bf91602480820192602092909190829003018186803b158015620004a857600080fd5b505afa158015620004bd573d6000803e3d6000fd5b505050506040513d6020811015620004d457600080fd5b505160408051631d9dabef60e11b81526004810185905290516001600160a01b0390921691633b3b57de91602480820192602092909190829003018186803b1580156200052057600080fd5b505afa15801562000535573d6000803e3d6000fd5b505050506040513d60208110156200054c57600080fd5b505192915050565b6154cf80620005646000396000f3fe60806040526004361061038c5760003560e01c80637fd004fa116101dc578063cc0e7e5611610102578063e2b4ce97116100a0578063f40b51f81161006f578063f40b51f8146110b0578063f41c4319146110da578063f421764814611104578063f776f5181461117f5761038c565b8063e2b4ce971461100e578063e61c51ca14611023578063eadd3cea1461104d578063f36febda146110775761038c565b8063ce0b5bd5116100dc578063ce0b5bd514610f90578063d251fefc14610fba578063da84b1ed14610fe4578063de212bf314610ff95761038c565b8063cc0e7e5614610eb8578063cccdc55614610ecd578063cd7958dd14610ee25761038c565b8063b221f3161161017a578063be40ba7911610149578063be40ba7914610e21578063beabacc814610e36578063c4856cd914610e79578063cbd2ac6814610e8e5761038c565b8063b221f31614610d6e578063b242e53414610d98578063b87e21ef14610dd3578063bcb8b74a14610e0c5761038c565b806390e690c7116101b657806390e690c714610c7e5780639b0dfd2714610c93578063aaf1fc6214610ca8578063ab20599314610d595761038c565b80637fd004fa14610bd9578063877337b014610c545780638da5cb5b14610c695761038c565b806332531c3c116102c15780635adc02ab1161025f57806374624c551161022e57806374624c5514610b54578063747c31d614610b7e5780637d73b23114610b935780637d7d004614610bc45761038c565b80635adc02ab14610a855780635d2362a814610aaf5780636137d67014610ac4578063715018a614610b3f5761038c565b80633c672eb71161029b5780633c672eb7146108ae5780633f579f42146108d857806346efe0ed1461099e57806347b55a9d14610a705761038c565b806332531c3c146108255780633a43199f146108585780633bfec254146108845761038c565b80631efd02991161032e57806321ce918d1161030857806321ce918d1461076c5780632587a6a21461079657806326d05ab2146107ab578063294f4025146107c05761038c565b80631efd02991461067757806320c13b0b1461068c5780632121dc75146107575761038c565b8063100f23fd1161036a578063100f23fd146104635780631127b57e1461048d5780631626ba7e146105175780631aa21fba146105ec5761038c565b806301ffc9a7146103c8578063027ef3eb146104105780630f3a85d814610437575b6040805133815234602082015281517f88a5966d370b9919b20f3e2c13ff65706f196a4e32cc2c12bf57088f88525874929181900390910190a1005b3480156103d457600080fd5b506103fc600480360360208110156103eb57600080fd5b50356001600160e01b031916611194565b604080519115158252519081900360200190f35b34801561041c57600080fd5b506104256111ae565b60408051918252519081900360200190f35b34801561044357600080fd5b506104616004803603602081101561045a57600080fd5b50356111b5565b005b34801561046f57600080fd5b506104616004803603602081101561048657600080fd5b50356112c1565b34801561049957600080fd5b506104a2611466565b6040805160208082528351818301528351919283929083019185019080838360005b838110156104dc5781810151838201526020016104c4565b50505050905090810190601f1680156105095780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561052357600080fd5b506105cf6004803603604081101561053a57600080fd5b81359190810190604081016020820135600160201b81111561055b57600080fd5b82018360208201111561056d57600080fd5b803590602001918460018302840111600160201b8311171561058e57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550611487945050505050565b604080516001600160e01b03199092168252519081900360200190f35b3480156105f857600080fd5b506104616004803603604081101561060f57600080fd5b6001600160a01b038235169190810190604081016020820135600160201b81111561063957600080fd5b82018360208201111561064b57600080fd5b803590602001918460208302840111600160201b8311171561066c57600080fd5b5090925090506114f5565b34801561068357600080fd5b5061042561167b565b34801561069857600080fd5b506105cf600480360360408110156106af57600080fd5b810190602081018135600160201b8111156106c957600080fd5b8201836020820111156106db57600080fd5b803590602001918460018302840111600160201b831117156106fc57600080fd5b919390929091602081019035600160201b81111561071957600080fd5b82018360208201111561072b57600080fd5b803590602001918460018302840111600160201b8311171561074c57600080fd5b50909250905061168c565b34801561076357600080fd5b506103fc611761565b34801561077857600080fd5b506104616004803603602081101561078f57600080fd5b5035611771565b3480156107a257600080fd5b5061042561180f565b3480156107b757600080fd5b506103fc611815565b3480156107cc57600080fd5b506107d561181e565b60408051602080825283518183015283519192839290830191858101910280838360005b838110156108115781810151838201526020016107f9565b505050509050019250505060405180910390f35b34801561083157600080fd5b506103fc6004803603602081101561084857600080fd5b50356001600160a01b0316611880565b6104616004803603604081101561086e57600080fd5b506001600160a01b038135169060200135611895565b34801561089057600080fd5b50610461600480360360208110156108a757600080fd5b5035611ad3565b3480156108ba57600080fd5b50610461600480360360208110156108d157600080fd5b5035611bcb565b3480156108e457600080fd5b506104a2600480360360608110156108fb57600080fd5b6001600160a01b0382351691602081013591810190606081016040820135600160201b81111561092a57600080fd5b82018360208201111561093c57600080fd5b803590602001918460018302840111600160201b8311171561095d57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550611c71945050505050565b3480156109aa57600080fd5b50610461600480360360608110156109c157600080fd5b81359190810190604081016020820135600160201b8111156109e257600080fd5b8201836020820111156109f457600080fd5b803590602001918460018302840111600160201b83111715610a1557600080fd5b919390929091602081019035600160201b811115610a3257600080fd5b820183602082011115610a4457600080fd5b803590602001918460018302840111600160201b83111715610a6557600080fd5b509092509050612168565b348015610a7c57600080fd5b506107d5612479565b348015610a9157600080fd5b5061046160048036036020811015610aa857600080fd5b50356124d9565b348015610abb57600080fd5b506104256127a9565b348015610ad057600080fd5b5061046160048036036020811015610ae757600080fd5b810190602081018135600160201b811115610b0157600080fd5b820183602082011115610b1357600080fd5b803590602001918460208302840111600160201b83111715610b3457600080fd5b5090925090506127b5565b348015610b4b57600080fd5b506104616129db565b348015610b6057600080fd5b5061046160048036036020811015610b7757600080fd5b5035612ad9565b348015610b8a57600080fd5b50610425612bdd565b348015610b9f57600080fd5b50610ba8612be3565b604080516001600160a01b039092168252519081900360200190f35b348015610bd057600080fd5b50610425612bf2565b348015610be557600080fd5b5061046160048036036020811015610bfc57600080fd5b810190602081018135600160201b811115610c1657600080fd5b820183602082011115610c2857600080fd5b803590602001918460208302840111600160201b83111715610c4957600080fd5b509092509050612bfe565b348015610c6057600080fd5b50610425612f40565b348015610c7557600080fd5b50610ba8612f46565b348015610c8a57600080fd5b50610461612f55565b348015610c9f57600080fd5b50610425612fb2565b348015610cb457600080fd5b5061046160048036036020811015610ccb57600080fd5b810190602081018135600160201b811115610ce557600080fd5b820183602082011115610cf757600080fd5b803590602001918460018302840111600160201b83111715610d1857600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550612fb8945050505050565b348015610d6557600080fd5b506103fc6130f5565b348015610d7a57600080fd5b5061046160048036036020811015610d9157600080fd5b50356130fe565b348015610da457600080fd5b5061046160048036036040811015610dbb57600080fd5b506001600160a01b03813516906020013515156131ee565b348015610ddf57600080fd5b5061042560048036036040811015610df657600080fd5b506001600160a01b0381351690602001356133a8565b348015610e1857600080fd5b506103fc613438565b348015610e2d57600080fd5b506103fc613441565b348015610e4257600080fd5b5061046160048036036060811015610e5957600080fd5b506001600160a01b03813581169160208101359091169060400135613450565b348015610e8557600080fd5b506104256135da565b348015610e9a57600080fd5b5061046160048036036020811015610eb157600080fd5b50356135e0565b348015610ec457600080fd5b5061042561395d565b348015610ed957600080fd5b50610425613963565b348015610eee57600080fd5b5061042560048036036020811015610f0557600080fd5b810190602081018135600160201b811115610f1f57600080fd5b820183602082011115610f3157600080fd5b803590602001918460208302840111600160201b83111715610f5257600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929550613969945050505050565b348015610f9c57600080fd5b5061046160048036036020811015610fb357600080fd5b50356139c3565b348015610fc657600080fd5b50610ba860048036036020811015610fdd57600080fd5b5035613b6c565b348015610ff057600080fd5b50610425613b93565b34801561100557600080fd5b506103fc613b99565b34801561101a57600080fd5b50610425613ba7565b34801561102f57600080fd5b506104616004803603602081101561104657600080fd5b5035613bad565b34801561105957600080fd5b506104616004803603602081101561107057600080fd5b5035613cf7565b34801561108357600080fd5b506104256004803603604081101561109a57600080fd5b506001600160a01b038135169060200135613d50565b3480156110bc57600080fd5b50610461600480360360208110156110d357600080fd5b5035613f03565b3480156110e657600080fd5b50610461600480360360208110156110fd57600080fd5b5035613f5c565b34801561111057600080fd5b506104616004803603602081101561112757600080fd5b810190602081018135600160201b81111561114157600080fd5b82018360208201111561115357600080fd5b803590602001918460208302840111600160201b8311171561117457600080fd5b509092509050613fb5565b34801561118b57600080fd5b506103fc614307565b6001600160e01b031981166301ffc9a760e01b145b919050565b6017545b90565b6111be33614310565b806111c857503330145b61120c576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b8066038d7ea4c680001115801561122b57506706f05b59d3b200008111155b611272576040805162461bcd60e51b815260206004820152601360248201527206f7574206f662072616e676520746f702d757606c1b604482015290519081900360640190fd5b61128360038263ffffffff61432416565b604080513381526020810183905281517f41ff5d5ce3b7935893a4e7269ec5caae9cca5e3bf0eb4b21d2f443489667112e929181900390910190a150565b6112ca33614310565b806112d957506112d93361438d565b611323576040805162461bcd60e51b815260206004820152601660248201527537b7363c9037bbb732b93e3e31b7b73a3937b63632b960511b604482015290519081900360640190fd5b60135460ff16611372576040805162461bcd60e51b81526020600482015260156024820152743737903832b73234b7339039bab136b4b9b9b4b7b760591b604482015290519081900360640190fd5b6113d560118054806020026020016040519081016040528092919081815260200182805480156113cb57602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116113ad575b5050505050613969565b81146114125760405162461bcd60e51b81526004018080602001828103825260238152602001806154426023913960400191505060405180910390fd5b61141e601160006152b2565b6013805460ff19169055604080513381526020810183905281517f7794eff834d760583543e6e510e717a5e66d2c064e225f4db448343c3e66afcf929181900390910190a150565b60405180604001604052806005815260200164332e312e3160d81b81525081565b60008061149a848463ffffffff61442116565b90506114a581614310565b6114e3576040805162461bcd60e51b815260206004820152600a60248201526937b7363c9037bbb732b960b11b604482015290519081900360640190fd5b50630b135d3f60e11b90505b92915050565b6114fe33614310565b8061150857503330145b61154c576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b80611595576040805162461bcd60e51b8152602060048201526014602482015273617373657420617272617920697320656d70747960601b604482015290519081900360640190fd5b60005b818110156115f85760006115c7308585858181106115b257fe5b905060200201356001600160a01b031661450f565b90506115ef858585858181106115d957fe5b905060200201356001600160a01b031683613450565b50600101611598565b507fd4f62f23021706247dcffea245d104ae7ddaec7f23acf3d11d7136d5de6a69ad83838360405180846001600160a01b03166001600160a01b03168152602001806020018281038252848482818152602001925060200280828437600083820152604051601f909101601f1916909201829003965090945050505050a1505050565b6000611687600a6145ba565b905090565b6000808585604051602001808383808284376040805191909301818103601f190182528084528151602092830120601f8b01839004830282018301909452898152929650630b135d3f60e11b9550611703945086935089915088908190840183828082843760009201919091525061148792505050565b6001600160e01b0319161461174f576040805162461bcd60e51b815260206004820152600d60248201526c1cda59c81b9bdd081d985b1a59609a1b604482015290519081900360640190fd5b506320c13b0b60e01b95945050505050565b600254600160a01b900460ff1690565b61177a33614310565b8061178457503330145b6117c8576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b6117d960148263ffffffff6145ef16565b6040805182815290517f4b1b970c8a0fa761e7803ed70c13d7aca71904b13df60fbe03f981da1730da919181900360200190a150565b60035490565b60135460ff1681565b6060601280548060200260200160405190810160405280929190818152602001828054801561187657602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311611858575b5050505050905090565b600f6020526000908152604090205460ff1681565b61189e33614310565b806118a857503330145b6118ec576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b6118f582614650565b61193b576040805162461bcd60e51b8152602060048201526012602482015271746f6b656e206e6f74206c6f616461626c6560701b604482015290519081900360640190fd5b60006119478383613d50565b905061195a600a8263ffffffff61466a16565b6000611967601a546146e0565b90506001600160a01b03841615611a0f576119926001600160a01b038516828563ffffffff6147a216565b806001600160a01b0316631b3c96b485856040518363ffffffff1660e01b815260040180836001600160a01b03166001600160a01b0316815260200182815260200192505050600060405180830381600087803b1580156119f257600080fd5b505af1158015611a06573d6000803e3d6000fd5b50505050611a89565b806001600160a01b0316631b3c96b48486866040518463ffffffff1660e01b815260040180836001600160a01b03166001600160a01b03168152602001828152602001925050506000604051808303818588803b158015611a6f57600080fd5b505af1158015611a83573d6000803e3d6000fd5b50505050505b604080516001600160a01b03861681526020810185905281517f5f65674bec9af81f71be68674135a0ea3f163fb91984e3893d06da9f6ea2ce8a929181900390910190a150505050565b611adc33614310565b80611ae657503330145b611b2a576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b600954811115611b7c576040805162461bcd60e51b81526020600482015260186024820152771bdd5d081bd9881c985b99d9481b1bd85908185b5bdd5b9d60421b604482015290519081900360640190fd5b611b8d600a8263ffffffff61432416565b604080513381526020810183905281517f0b05243483e17c3f3377aee82b7d47e5700b48288695fc08b7ecc2759afa44ef929181900390910190a150565b611bd433614310565b80611bde57503330145b611c22576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b611c3360148263ffffffff61432416565b604080513381526020810183905281517f068f112e5ec923d412be64779fe69e0fcbb6784c6617e94cccc8fd348f2e0f21929181900390910190a150565b6060611c7c33614310565b80611c8657503330145b611cca576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b6001600160a01b0384166000908152600f602052604090205460ff16611cfb57611cfb60148463ffffffff61466a16565b611d0d846001600160a01b03166148ba565b8015611d1d5750611d1d846148c0565b15611f0457600080611d2f86856148da565b6001600160a01b0382166000908152600f6020526040902054919350915060ff16611d75576000611d6087836133a8565b9050611d7360148263ffffffff61466a16565b505b611d8e6001600160a01b0387168563ffffffff6149e416565b604080516020808252818301909252606091602082018180388339019050509050600160f81b81601f81518110611dc157fe5b60200101906001600160f81b031916908160001a9053507ff77753fab406ecfff96d6ff2476c64a838fa9f6d37b1bf190f8546e395e3b6138787878460405180856001600160a01b03166001600160a01b031681526020018481526020018060200180602001838103835285818151815260200191508051906020019080838360005b83811015611e5c578181015183820152602001611e44565b50505050905090810190601f168015611e895780820380516001836020036101000a031916815260200191505b50838103825284518152845160209182019186019080838360005b83811015611ebc578181015183820152602001611ea4565b50505050905090810190601f168015611ee95780820380516001836020036101000a031916815260200191505b50965050505050505060405180910390a19250612161915050565b60006060856001600160a01b031685856040518082805190602001908083835b60208310611f435780518252601f199092019160209182019101611f24565b6001836020036101000a03801982511681845116808217855250505050505090500191505060006040518083038185875af1925050503d8060008114611fa5576040519150601f19603f3d011682016040523d82523d6000602084013e611faa565b606091505b509150915081819061203a5760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b83811015611fff578181015183820152602001611fe7565b50505050905090810190601f16801561202c5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b507ff77753fab406ecfff96d6ff2476c64a838fa9f6d37b1bf190f8546e395e3b6138686868460405180856001600160a01b03166001600160a01b031681526020018481526020018060200180602001838103835285818151815260200191508051906020019080838360005b838110156120bf5781810151838201526020016120a7565b50505050905090810190601f1680156120ec5780820380516001836020036101000a031916815260200191505b50838103825284518152845160209182019186019080838360005b8381101561211f578181015183820152602001612107565b50505050905090810190601f16801561214c5780820380516001836020036101000a031916815260200191505b50965050505050505060405180910390a19150505b9392505050565b6121713361438d565b6121b0576040805162461bcd60e51b815260206004820152601a60248201526000805160206153b4833981519152604482015290519081900360640190fd5b600061220186868660405160200180806339363c1d60e11b81525060040184815260200183838082843780830192505050935050505060405160208183030381529060405280519060200120614ba2565b9050631626ba7e60e01b6001600160e01b0319166122558285858080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061148792505050565b6001600160e01b031916146122a1576040805162461bcd60e51b815260206004820152600d60248201526c1cda59c81b9bdd081d985b1a59609a1b604482015290519081900360640190fd5b60195486146122e3576040805162461bcd60e51b81526020600482015260096024820152687478207265706c617960b81b604482015290519081900360640190fd5b6122eb614bf3565b60006060306001600160a01b03168787604051808383808284376040519201945060009350909150508083038183865af19150503d806000811461234b576040519150601f19603f3d011682016040523d82523d6000602084013e612350565b606091505b50915091508181906123a35760405162461bcd60e51b8152602060048201818152835160248401528351909283926044909101919085019080838360008315611fff578181015183820152602001611fe7565b507f823dbcf2b7b0f265871963ca65ac033f6b4c71e0d82cd123d2ff23d752dc21c187878360405180806020018060200183810383528686828181526020019250808284376000838201819052601f909101601f191690920185810384528651815286516020918201939188019250908190849084905b8381101561243257818101518382015260200161241a565b50505050905090810190601f16801561245f5780820380516001836020036101000a031916815260200191505b509550505050505060405180910390a15050505050505050565b60606011805480602002602001604051908101604052809291908181526020018280548015611876576020028201919060005260206000209081546001600160a01b03168152600190910190602001808311611858575050505050905090565b6124e23361438d565b612521576040805162461bcd60e51b815260206004820152601a60248201526000805160206153b4833981519152604482015290519081900360640190fd5b60135460ff16612570576040805162461bcd60e51b81526020600482015260156024820152743737903832b73234b7339039bab136b4b9b9b4b7b760591b604482015290519081900360640190fd5b6125d160118054806020026020016040519081016040528092919081815260200182805480156113cb576020028201919060005260206000209081546001600160a01b031681526001909101906020018083116113ad575050505050613969565b811461260e5760405162461bcd60e51b81526004018080602001828103825260238152602001806154426023913960400191505060405180910390fd5b60005b6011548110156126f557600f60006011838154811061262c57fe5b60009182526020808320909101546001600160a01b0316835282019290925260400190205460ff166126ed576001600f60006011848154811061266b57fe5b6000918252602080832091909101546001600160a01b031683528201929092526040019020805460ff19169115159190911790556011805460109190839081106126b157fe5b60009182526020808320909101548354600181018555938352912090910180546001600160a01b0319166001600160a01b039092169190911790555b600101612611565b507fb2f6cccee7a369e23e293c25aa19bef80af11eb26deba3ea0f2a02783f752e4a33601160405180836001600160a01b03166001600160a01b0316815260200180602001828103825283818154815260200191508054801561278157602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311612763575b5050935050505060405180910390a161279c601160006152b2565b506013805460ff19169055565b600061168760146145ba565b6127be33614310565b806127c857503330145b61280c576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b60135460ff161580156128275750601354610100900460ff16155b612878576040805162461bcd60e51b815260206004820152601c60248201527f77686974656c6973742073756d62697373696f6e2070656e64696e6700000000604482015290519081900360640190fd5b60135462010000900460ff166128d1576040805162461bcd60e51b81526020600482015260196024820152781dda1a5d195b1a5cdd081b9bdd081a5b9a5d1a585b1a5e9959603a1b604482015290519081900360640190fd5b80612915576040805162461bcd60e51b815260206004820152600f60248201526e195b5c1d1e481dda1a5d195b1a5cdd608a1b604482015290519081900360640190fd5b612921601283836152d0565b506013805461ff00191661010017905560408051602080840282810182019093528382527ffbc0e5ca6c7e4858daf0fdb185ef5186203e74ec9c64737e93c0aeaec596e1d192859285926129909285918591829185019084908082843760009201919091525061396992505050565b60405180806020018381526020018281038252858582818152602001925060200280828437600083820152604051601f909101601f1916909201829003965090945050505050a15050565b6129e433614310565b612a2e576040805162461bcd60e51b815260206004820152601660248201527539b2b73232b91034b9903737ba1030b71037bbb732b960511b604482015290519081900360640190fd5b600254600160a01b900460ff16612a8c576040805162461bcd60e51b815260206004820152601d60248201527f6f776e657273686970206973206e6f74207472616e7366657261626c65000000604482015290519081900360640190fd5b600280546001600160a01b0319169055604080516000808252602082015281517f850b3df64837d7d518b45f5aa64d104652c3b80eb5b34a8e3d9eb666cb7cdea5929181900390910190a1565b612ae233614310565b80612aec57503330145b612b30576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b8066038d7ea4c6800011158015612b4f57506706f05b59d3b200008111155b612b96576040805162461bcd60e51b815260206004820152601360248201527206f7574206f662072616e676520746f702d757606c1b604482015290519081900360640190fd5b612ba760038263ffffffff6145ef16565b6040805182815290517faf2a77cd04c3cc155588dd3bf67b310ab4fb3b1da3cf6b8d7d4d2aa1d09b794c9181900360200190a150565b601a5490565b6000546001600160a01b031690565b600061168760036145ba565b612c0733614310565b80612c1157503330145b612c55576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b60135460ff16158015612c705750601354610100900460ff16155b612cc1576040805162461bcd60e51b815260206004820152601c60248201527f77686974656c6973742073756d62697373696f6e2070656e64696e6700000000604482015290519081900360640190fd5b8181808060200260200160405190810160405280939291908181526020018383602002808284376000920182905250925050505b8151811015612ddd57612d1a828281518110612d0d57fe5b6020026020010151614310565b15612d65576040805162461bcd60e51b8152602060048201526016602482015275636f6e7461696e73206f776e6572206164647265737360501b604482015290519081900360640190fd5b60006001600160a01b0316828281518110612d7c57fe5b60200260200101516001600160a01b03161415612dd5576040805162461bcd60e51b8152602060048201526012602482015271636f6e7461696e732030206164647265737360701b604482015290519081900360640190fd5b600101612cf5565b5060135462010000900460ff16612e37576040805162461bcd60e51b81526020600482015260196024820152781dda1a5d195b1a5cdd081b9bdd081a5b9a5d1a585b1a5e9959603a1b604482015290519081900360640190fd5b81612e7b576040805162461bcd60e51b815260206004820152600f60248201526e195b5c1d1e481dda1a5d195b1a5cdd608a1b604482015290519081900360640190fd5b612e87601184846152d0565b506013805460ff1916600117905560408051602080850282810182019093528482527f9c80b3b5f68b3e017766d59e8d09b34efe6462b05c398f35cab9e271d9bc3b9c9286928692612ef49285918591829185019084908082843760009201919091525061396992505050565b60405180806020018381526020018281038252858582818152602001925060200280828437600083820152604051601f909101601f1916909201829003965090945050505050a1505050565b60085490565b6002546001600160a01b031690565b612f5e33614310565b612fa8576040805162461bcd60e51b815260206004820152601660248201527539b2b73232b91034b9903737ba1030b71037bbb732b960511b604482015290519081900360640190fd5b612fb0614bf3565b565b60145490565b612fc133614310565b80612fcb57503330145b61300f576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b8051602080820191906000808060605b868510156130eb5761303886605463ffffffff614c3b16565b888601805160148201516034909201805193995060609190911c9650909450909250905061307d6054613071878563ffffffff614c9816565b9063ffffffff614c9816565b9450868511156130c4576040805162461bcd60e51b815260206004820152600d60248201526c6f7574206f6620626f756e647360981b604482015290519081900360640190fd5b816130da57506040805160208101909152600081525b6130e5848483611c71565b5061301f565b5050505050505050565b600e5460ff1690565b61310733614310565b8061311157503330145b613155576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b6009548111156131a7576040805162461bcd60e51b81526020600482015260186024820152771bdd5d081bd9881c985b99d9481b1bd85908185b5bdd5b9d60421b604482015290519081900360640190fd5b6131b8600a8263ffffffff6145ef16565b6040805182815290517fc178d379965e5657b6fc57494e392f121a14119215dfb422aad7db4cc03f2d109181900360200190a150565b6131f733614310565b613241576040805162461bcd60e51b815260206004820152601660248201527539b2b73232b91034b9903737ba1030b71037bbb732b960511b604482015290519081900360640190fd5b600254600160a01b900460ff1661329f576040805162461bcd60e51b815260206004820152601d60248201527f6f776e657273686970206973206e6f74207472616e7366657261626c65000000604482015290519081900360640190fd5b6001600160a01b0382166132e45760405162461bcd60e51b81526004018080602001828103825260238152602001806153f56023913960400191505060405180910390fd5b6002805460ff60a01b1916600160a01b831515021790558061333d57604080516001600160a01b038416815290517f808639ff9c8e4732d60b6c2330de498035416d229f27a77d259680895efec1229181900360200190a15b600254604080516001600160a01b039283168152918416602083015280517f850b3df64837d7d518b45f5aa64d104652c3b80eb5b34a8e3d9eb666cb7cdea59281900390910190a150600280546001600160a01b0319166001600160a01b0392909216919091179055565b6000806000806133b786614cf2565b50505093509350935050801561342c5781613402576040805162461bcd60e51b81526020600482015260066024820152650726174653d360d41b604482015290519081900360640190fd5b61342283613416878563ffffffff614e8416565b9063ffffffff614edd16565b93505050506114ef565b50600095945050505050565b60185460ff1690565b60135462010000900460ff1681565b61345933614310565b8061346357503330145b6134a7576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b80806134e4576040805162461bcd60e51b8152602060048201526007602482015266076616c75653d360cc1b604482015290519081900360640190fd5b6001600160a01b03841661352f576040805162461bcd60e51b815260206004820152600d60248201526c064657374696e6174696f6e3d3609c1b604482015290519081900360640190fd5b6001600160a01b0384166000908152600f602052604090205460ff1661357f57816001600160a01b0384161561356c5761356984846133a8565b90505b61357d60148263ffffffff61466a16565b505b61358a848484614f47565b604080516001600160a01b0380871682528516602082015280820184905290517fd1ba4ac2e2a11b5101f6cb4d978f514a155b421e8ec396d2d9abaf0bb02917ee9181900360600190a150505050565b600d5490565b6135e93361438d565b613628576040805162461bcd60e51b815260206004820152601a60248201526000805160206153b4833981519152604482015290519081900360640190fd5b601354610100900460ff1661367c576040805162461bcd60e51b81526020600482015260156024820152743737903832b73234b7339039bab136b4b9b9b4b7b760591b604482015290519081900360640190fd5b6136dd60128054806020026020016040519081016040528092919081815260200182805480156113cb576020028201919060005260206000209081546001600160a01b031681526001909101906020018083116113ad575050505050613969565b811461371a5760405162461bcd60e51b81526004018080602001828103825260238152602001806154426023913960400191505060405180910390fd5b60005b6012548110156138a857600f60006012838154811061373857fe5b60009182526020808320909101546001600160a01b0316835282019290925260400190205460ff16156138a0576000600f60006012848154811061377857fe5b6000918252602080832091909101546001600160a01b031683528201929092526040018120805460ff1916921515929092179091555b6010546137c290600163ffffffff614c3b16565b81101561388a57601282815481106137d657fe5b600091825260209091200154601080546001600160a01b0390921691839081106137fc57fe5b6000918252602090912001546001600160a01b031614156138825760108054600019810190811061382957fe5b600091825260209091200154601080546001600160a01b03909216918390811061384f57fe5b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555061388a565b6001016137ae565b50601080549061389e906000198301615333565b505b60010161371d565b507fd218c430fa348f4ce67791021b6b89c0c3eacd4ead1d8f5b83c60038ec28249b33601260405180836001600160a01b03166001600160a01b0316815260200180602001828103825283818154815260200191508054801561393457602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311613916575b5050935050505060405180910390a161394f601260006152b2565b506013805461ff0019169055565b60065490565b60195481565b60008160405160200180828051906020019060200280838360005b8381101561399c578181015183820152602001613984565b50505050905001915050604051602081830303815290604052805190602001209050919050565b6139cc33614310565b806139db57506139db3361438d565b613a25576040805162461bcd60e51b815260206004820152601660248201527537b7363c9037bbb732b93e3e31b7b73a3937b63632b960511b604482015290519081900360640190fd5b601354610100900460ff16613a79576040805162461bcd60e51b81526020600482015260156024820152743737903832b73234b7339039bab136b4b9b9b4b7b760591b604482015290519081900360640190fd5b613ada60128054806020026020016040519081016040528092919081815260200182805480156113cb576020028201919060005260206000209081546001600160a01b031681526001909101906020018083116113ad575050505050613969565b8114613b175760405162461bcd60e51b81526004018080602001828103825260238152602001806154426023913960400191505060405180910390fd5b613b23601260006152b2565b6013805461ff0019169055604080513381526020810183905281517f13c935eb475aa0f6e931fece83e2ac44569ce2d53460d29a6dedab40b965c8a3929181900390910190a150565b60108181548110613b7957fe5b6000918252602090912001546001600160a01b0316905081565b600a5490565b601354610100900460ff1681565b60015490565b8080613bea576040805162461bcd60e51b8152602060048201526007602482015266076616c75653d360cc1b604482015290519081900360640190fd5b613bf333614310565b80613c025750613c023361438d565b613c4c576040805162461bcd60e51b815260206004820152601660248201527537b7363c9037bbb732b93e3e31b7b73a3937b63632b960511b604482015290519081900360640190fd5b613c5d60038363ffffffff61466a16565b613c65612f46565b6001600160a01b03166108fc839081150290604051600060405180830381858888f19350505050158015613c9d573d6000803e3d6000fd5b507f611b7c0d84fda988026215bef9b3e4d81cbceced7e679be6d5e044b588467c0e33613cc8612f46565b604080516001600160a01b03938416815291909216602082015280820185905290519081900360600190a15050565b613d003361438d565b613d3f576040805162461bcd60e51b815260206004820152601a60248201526000805160206153b4833981519152604482015290519081900360640190fd5b611c3360148263ffffffff614fab16565b6000613d5a614fff565b6001600160a01b0316836001600160a01b03161415613d7a5750806114ef565b816001600160a01b03841615613e3f576000806000613d9887614cf2565b5050509350935093505080613dea576040805162461bcd60e51b8152602060048201526013602482015272746f6b656e206e6f7420617661696c61626c6560681b604482015290519081900360640190fd5b81613e25576040805162461bcd60e51b81526020600482015260066024820152650726174653d360d41b604482015290519081900360640190fd5b613e3983613416888563ffffffff614e8416565b93505050505b6000806000613e4c615075565b5050509350935093505080613e9e576040805162461bcd60e51b8152602060048201526013602482015272746f6b656e206e6f7420617661696c61626c6560681b604482015290519081900360640190fd5b81613ee4576040805162461bcd60e51b81526020600482015260116024820152700737461626c65636f696e20726174653d3607c1b604482015290519081900360640190fd5b613ef882613416868663ffffffff614e8416565b979650505050505050565b613f0c3361438d565b613f4b576040805162461bcd60e51b815260206004820152601a60248201526000805160206153b4833981519152604482015290519081900360640190fd5b611b8d600a8263ffffffff614fab16565b613f653361438d565b613fa4576040805162461bcd60e51b815260206004820152601a60248201526000805160206153b4833981519152604482015290519081900360640190fd5b61128360038263ffffffff614fab16565b613fbe33614310565b80613fc857503330145b61400c576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b8181808060200260200160405190810160405280939291908181526020018383602002808284376000920182905250925050505b815181101561411b57614058828281518110612d0d57fe5b156140a3576040805162461bcd60e51b8152602060048201526016602482015275636f6e7461696e73206f776e6572206164647265737360501b604482015290519081900360640190fd5b60006001600160a01b03168282815181106140ba57fe5b60200260200101516001600160a01b03161415614113576040805162461bcd60e51b8152602060048201526012602482015271636f6e7461696e732030206164647265737360701b604482015290519081900360640190fd5b600101614040565b5060135462010000900460ff1615614172576040805162461bcd60e51b81526020600482015260156024820152741dda1a5d195b1a5cdd081a5b9a5d1a585b1a5e9959605a1b604482015290519081900360640190fd5b60005b8281101561426357600f600085858481811061418d57fe5b602090810292909201356001600160a01b03168352508101919091526040016000205460ff1661425b576001600f60008686858181106141c957fe5b905060200201356001600160a01b03166001600160a01b03166001600160a01b0316815260200190815260200160002060006101000a81548160ff021916908315150217905550601084848381811061421e57fe5b835460018101855560009485526020948590200180546001600160a01b0319166001600160a01b0395909202939093013593909316929092179055505b600101614175565b506013805462ff0000191662010000179055604080513380825260208201838152601080549484018590527fb2f6cccee7a369e23e293c25aa19bef80af11eb26deba3ea0f2a02783f752e4a94929390929091906060830190849080156142f357602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116142d5575b5050935050505060405180910390a1505050565b60075460ff1690565b6002546001600160a01b0390811691161490565b600482015460ff1615614372576040805162461bcd60e51b81526020600482015260116024820152701b1a5b5a5d08185b1c9958591e481cd95d607a1b604482015290519081900360640190fd5b61437c82826151e4565b50600401805460ff19166001179055565b600061439a6001546146e0565b6001600160a01b031663b429afeb836040518263ffffffff1660e01b815260040180826001600160a01b03166001600160a01b0316815260200191505060206040518083038186803b1580156143ef57600080fd5b505afa158015614403573d6000803e3d6000fd5b505050506040513d602081101561441957600080fd5b505192915050565b60008151604114614434575060006114ef565b60208201516040830151606084015160001a7f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a082111561447a57600093505050506114ef565b8060ff16601b1415801561449257508060ff16601c14155b156144a357600093505050506114ef565b6040805160008152602080820180845289905260ff8416828401526060820186905260808201859052915160019260a0808401939192601f1981019281900390910190855afa1580156144fa573d6000803e3d6000fd5b5050604051601f190151979650505050505050565b60006001600160a01b038216156145a957816001600160a01b03166370a08231846040518263ffffffff1660e01b815260040180826001600160a01b03166001600160a01b0316815260200191505060206040518083038186803b15801561457657600080fd5b505afa15801561458a573d6000803e3d6000fd5b505050506040513d60208110156145a057600080fd5b505190506114ef565b506001600160a01b038216316114ef565b60028101546000906145d5906201518063ffffffff614c9816565b4211156145e4575080546111a9565b5060018101546111a9565b600482015460ff16614648576040805162461bcd60e51b815260206004820152601960248201527f6c696d6974206861736e2774206265656e207365742079657400000000000000604482015290519081900360640190fd5b600390910155565b60008061465c83614cf2565b509098975050505050505050565b61467382615207565b80826001015410156146bf576040805162461bcd60e51b815260206004820152601060248201526f185d985a5b18589b194f185b5bdd5b9d60821b604482015290519081900360640190fd5b60018201546146d4908263ffffffff614c3b16565b82600101819055505050565b6000805460408051630178b8bf60e01b81526004810185905290516001600160a01b0390921691630178b8bf91602480820192602092909190829003018186803b15801561472d57600080fd5b505afa158015614741573d6000803e3d6000fd5b505050506040513d602081101561475757600080fd5b505160408051631d9dabef60e11b81526004810185905290516001600160a01b0390921691633b3b57de91602480820192602092909190829003018186803b1580156143ef57600080fd5b801580614828575060408051636eb1769f60e11b81523060048201526001600160a01b03848116602483015291519185169163dd62ed3e91604480820192602092909190829003018186803b1580156147fa57600080fd5b505afa15801561480e573d6000803e3d6000fd5b505050506040513d602081101561482457600080fd5b5051155b6148635760405162461bcd60e51b81526004018080602001828103825260368152602001806154656036913960400191505060405180910390fd5b604080516001600160a01b038416602482015260448082018490528251808303909101815260649091019091526020810180516001600160e01b031663095ea7b360e01b1790526148b59084906149e4565b505050565b3b151590565b6000806148cc83614cf2565b509198975050505050505050565b6000806148e86008546146e0565b6001600160a01b031663afc72e9385856040518363ffffffff1660e01b815260040180836001600160a01b03166001600160a01b0316815260200180602001828103825283818151815260200191508051906020019080838360005b8381101561495c578181015183820152602001614944565b50505050905090810190601f1680156149895780820380516001836020036101000a031916815260200191505b509350505050604080518083038186803b1580156149a657600080fd5b505afa1580156149ba573d6000803e3d6000fd5b505050506040513d60408110156149d057600080fd5b508051602090910151909590945092505050565b6149f6826001600160a01b03166148ba565b614a47576040805162461bcd60e51b815260206004820152601f60248201527f5361666545524332303a2063616c6c20746f206e6f6e2d636f6e747261637400604482015290519081900360640190fd5b60006060836001600160a01b0316836040518082805190602001908083835b60208310614a855780518252601f199092019160209182019101614a66565b6001836020036101000a0380198251168184511680821785525050505050509050019150506000604051808303816000865af19150503d8060008114614ae7576040519150601f19603f3d011682016040523d82523d6000602084013e614aec565b606091505b509150915081614b43576040805162461bcd60e51b815260206004820181905260248201527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564604482015290519081900360640190fd5b805115614b9c57808060200190516020811015614b5f57600080fd5b5051614b9c5760405162461bcd60e51b815260040180806020018281038252602a815260200180615418602a913960400191505060405180910390fd5b50505050565b604080517f19457468657265756d205369676e6564204d6573736167653a0a333200000000602080830191909152603c8083019490945282518083039094018452605c909101909152815191012090565b601980546001019081905560408051338152602081019290925280517fab0423a75986556234aecd171c46ce7f5e45607d8070bf5230f2735b50322bff9281900390910190a1565b600082821115614c92576040805162461bcd60e51b815260206004820152601e60248201527f536166654d6174683a207375627472616374696f6e206f766572666c6f770000604482015290519081900360640190fd5b50900390565b600082820183811015612161576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b6060600080600080600080614d086008546146e0565b6001600160a01b0316631f69565f896040518263ffffffff1660e01b815260040180826001600160a01b03166001600160a01b0316815260200191505060006040518083038186803b158015614d5d57600080fd5b505afa158015614d71573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405260e0811015614d9a57600080fd5b8101908080516040519392919084600160201b821115614db957600080fd5b908301906020820185811115614dce57600080fd5b8251600160201b811182820188101715614de757600080fd5b82525081516020918201929091019080838360005b83811015614e14578181015183820152602001614dfc565b50505050905090810190601f168015614e415780820380516001836020036101000a031916815260200191505b5060409081526020820151908201516060830151608084015160a085015160c090950151979e50929c50909a509850965094509192505050919395979092949650565b600082614e93575060006114ef565b82820282848281614ea057fe5b04146121615760405162461bcd60e51b81526004018080602001828103825260218152602001806153d46021913960400191505060405180910390fd5b6000808211614f33576040805162461bcd60e51b815260206004820152601a60248201527f536166654d6174683a206469766973696f6e206279207a65726f000000000000604482015290519081900360640190fd5b6000828481614f3e57fe5b04949350505050565b6001600160a01b038216614f91576040516001600160a01b0384169082156108fc029083906000818181858888f19350505050158015614f8b573d6000803e3d6000fd5b506148b5565b6148b56001600160a01b038316848363ffffffff61526016565b80826003015414614fed5760405162461bcd60e51b81526004018080602001828103825260228152602001806153926022913960400191505060405180910390fd5b614ffb8283600301546151e4565b5050565b600061500c6008546146e0565b6001600160a01b031663e9cbd8226040518163ffffffff1660e01b815260040160206040518083038186803b15801561504457600080fd5b505afa158015615058573d6000803e3d6000fd5b505050506040513d602081101561506e57600080fd5b5051905090565b606060008060008060008061508b6008546146e0565b6001600160a01b0316633efec5e96040518163ffffffff1660e01b815260040160006040518083038186803b1580156150c357600080fd5b505afa1580156150d7573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405260e081101561510057600080fd5b8101908080516040519392919084600160201b82111561511f57600080fd5b90830190602082018581111561513457600080fd5b8251600160201b81118282018810171561514d57600080fd5b82525081516020918201929091019080838360005b8381101561517a578181015183820152602001615162565b50505050905090810190601f1680156151a75780820380516001836020036101000a031916815260200191505b5060409081526020820151908201516060830151608084015160a085015160c090950151979f939e50919c509a5098509096509294509192505050565b6151ed82615207565b8082556001820154811015614ffb57815460018301555050565b600281015461521f906201518063ffffffff614c9816565b42111561525d57426002820155805460018201556040517fe93bc25276d408d390778e7a8b926f2f67209c43ed540081b951fe128f0d3cd290600090a15b50565b604080516001600160a01b038416602482015260448082018490528251808303909101815260649091019091526020810180516001600160e01b031663a9059cbb60e01b1790526148b59084906149e4565b508054600082559060005260206000209081019061525d9190615353565b828054828255906000526020600020908101928215615323579160200282015b828111156153235781546001600160a01b0319166001600160a01b038435161782556020909201916001909101906152f0565b5061532f92915061536d565b5090565b8154818355818111156148b5576000838152602090206148b59181019083015b6111b291905b8082111561532f5760008155600101615359565b6111b291905b8082111561532f5780546001600160a01b031916815560010161537356fe636f6e6669726d65642f7375626d6974746564206c696d6974206d69736d6174636873656e646572206973206e6f74206120636f6e74726f6c6c6572000000000000536166654d6174683a206d756c7469706c69636174696f6e206f766572666c6f776f776e65722063616e6e6f742062652073657420746f207a65726f20616464726573735361666545524332303a204552433230206f7065726174696f6e20646964206e6f7420737563636565646e6f6e2d6d61746368696e672070656e64696e672077686974656c69737420686173685361666545524332303a20617070726f76652066726f6d206e6f6e2d7a65726f20746f206e6f6e2d7a65726f20616c6c6f77616e6365a265627a7a72315820a7321f3cf8984c18d49b59bfb92e2c20d256e7ae3b1df51d97e62da513f6daf164736f6c63430005110032"
=======
<<<<<<< HEAD
var WalletBin = "0x60806040527f7f2ce995617d2816b426c5c8698c5ec2952f7a34bb10f38326f74933d58936976001553480156200003557600080fd5b5060405162005a3338038062005a33833981810160405260e08110156200005b57600080fd5b508051602082015160408301516060840151608085015160a086015160c090960151600080546001600160a01b0319166001600160a01b03861617905594959394929391929091908084808989878015620000b65760018190555b50600280546001600160a01b0319166001600160a01b0384161760ff60a01b1916600160a01b8315158102919091179182905560ff9104166200013057604080516001600160a01b038416815290517f808639ff9c8e4732d60b6c2330de498035416d229f27a77d259680895efec1229181900360200190a15b60408051600081526001600160a01b038416602082015281517f850b3df64837d7d518b45f5aa64d104652c3b80eb5b34a8e3d9eb666cb7cdea5929181900390910190a150506040805160a0810182526706f05b59d3b2000080825260208201819052429282018390526000606083018190526080909201829052600381905560045560059190915560068190556007805460ff19169055600891909155620001e16001600160e01b03620002d816565b5050505050915050600081116200022f576040805162461bcd60e51b815260206004820152600d60248201526c37379039ba30b13632b1b7b4b760991b604482015290519081900360640190fd5b6127100260098190556040805160a08082018352838252602080830185905242838501819052600060608086018290526080958601829052600a889055600b97909755600c829055600d819055600e805460ff1990811690915586519485018752898552928401899052948301819052948201849052910182905260148590556015949094556016919091556017555060188054909116905550601a5550620005549350505050565b6060600080600080600080620002f66008546200045a60201b60201c565b6001600160a01b0316633efec5e96040518163ffffffff1660e01b815260040160006040518083038186803b1580156200032f57600080fd5b505afa15801562000344573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405260e08110156200036e57600080fd5b81019080805160405193929190846401000000008211156200038f57600080fd5b908301906020820185811115620003a557600080fd5b8251640100000000811182820188101715620003c057600080fd5b82525081516020918201929091019080838360005b83811015620003ef578181015183820152602001620003d5565b50505050905090810190601f1680156200041d5780820380516001836020036101000a031916815260200191505b5060409081526020820151908201516060830151608084015160a085015160c090950151979f939e50919c509a5098509096509294509192505050565b6000805460408051630178b8bf60e01b81526004810185905290516001600160a01b0390921691630178b8bf91602480820192602092909190829003018186803b158015620004a857600080fd5b505afa158015620004bd573d6000803e3d6000fd5b505050506040513d6020811015620004d457600080fd5b505160408051631d9dabef60e11b81526004810185905290516001600160a01b0390921691633b3b57de91602480820192602092909190829003018186803b1580156200052057600080fd5b505afa15801562000535573d6000803e3d6000fd5b505050506040513d60208110156200054c57600080fd5b505192915050565b6154cf80620005646000396000f3fe60806040526004361061038c5760003560e01c80637fd004fa116101dc578063cc0e7e5611610102578063e2b4ce97116100a0578063f40b51f81161006f578063f40b51f8146110b0578063f41c4319146110da578063f421764814611104578063f776f5181461117f5761038c565b8063e2b4ce971461100e578063e61c51ca14611023578063eadd3cea1461104d578063f36febda146110775761038c565b8063ce0b5bd5116100dc578063ce0b5bd514610f90578063d251fefc14610fba578063da84b1ed14610fe4578063de212bf314610ff95761038c565b8063cc0e7e5614610eb8578063cccdc55614610ecd578063cd7958dd14610ee25761038c565b8063b221f3161161017a578063be40ba7911610149578063be40ba7914610e21578063beabacc814610e36578063c4856cd914610e79578063cbd2ac6814610e8e5761038c565b8063b221f31614610d6e578063b242e53414610d98578063b87e21ef14610dd3578063bcb8b74a14610e0c5761038c565b806390e690c7116101b657806390e690c714610c7e5780639b0dfd2714610c93578063aaf1fc6214610ca8578063ab20599314610d595761038c565b80637fd004fa14610bd9578063877337b014610c545780638da5cb5b14610c695761038c565b806332531c3c116102c15780635adc02ab1161025f57806374624c551161022e57806374624c5514610b54578063747c31d614610b7e5780637d73b23114610b935780637d7d004614610bc45761038c565b80635adc02ab14610a855780635d2362a814610aaf5780636137d67014610ac4578063715018a614610b3f5761038c565b80633c672eb71161029b5780633c672eb7146108ae5780633f579f42146108d857806346efe0ed1461099e57806347b55a9d14610a705761038c565b806332531c3c146108255780633a43199f146108585780633bfec254146108845761038c565b80631efd02991161032e57806321ce918d1161030857806321ce918d1461076c5780632587a6a21461079657806326d05ab2146107ab578063294f4025146107c05761038c565b80631efd02991461067757806320c13b0b1461068c5780632121dc75146107575761038c565b8063100f23fd1161036a578063100f23fd146104635780631127b57e1461048d5780631626ba7e146105175780631aa21fba146105ec5761038c565b806301ffc9a7146103c8578063027ef3eb146104105780630f3a85d814610437575b6040805133815234602082015281517f88a5966d370b9919b20f3e2c13ff65706f196a4e32cc2c12bf57088f88525874929181900390910190a1005b3480156103d457600080fd5b506103fc600480360360208110156103eb57600080fd5b50356001600160e01b031916611194565b604080519115158252519081900360200190f35b34801561041c57600080fd5b506104256111ae565b60408051918252519081900360200190f35b34801561044357600080fd5b506104616004803603602081101561045a57600080fd5b50356111b5565b005b34801561046f57600080fd5b506104616004803603602081101561048657600080fd5b50356112c1565b34801561049957600080fd5b506104a2611466565b6040805160208082528351818301528351919283929083019185019080838360005b838110156104dc5781810151838201526020016104c4565b50505050905090810190601f1680156105095780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561052357600080fd5b506105cf6004803603604081101561053a57600080fd5b81359190810190604081016020820135600160201b81111561055b57600080fd5b82018360208201111561056d57600080fd5b803590602001918460018302840111600160201b8311171561058e57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550611487945050505050565b604080516001600160e01b03199092168252519081900360200190f35b3480156105f857600080fd5b506104616004803603604081101561060f57600080fd5b6001600160a01b038235169190810190604081016020820135600160201b81111561063957600080fd5b82018360208201111561064b57600080fd5b803590602001918460208302840111600160201b8311171561066c57600080fd5b5090925090506114f5565b34801561068357600080fd5b5061042561167b565b34801561069857600080fd5b506105cf600480360360408110156106af57600080fd5b810190602081018135600160201b8111156106c957600080fd5b8201836020820111156106db57600080fd5b803590602001918460018302840111600160201b831117156106fc57600080fd5b919390929091602081019035600160201b81111561071957600080fd5b82018360208201111561072b57600080fd5b803590602001918460018302840111600160201b8311171561074c57600080fd5b50909250905061168c565b34801561076357600080fd5b506103fc611761565b34801561077857600080fd5b506104616004803603602081101561078f57600080fd5b5035611771565b3480156107a257600080fd5b5061042561180f565b3480156107b757600080fd5b506103fc611815565b3480156107cc57600080fd5b506107d561181e565b60408051602080825283518183015283519192839290830191858101910280838360005b838110156108115781810151838201526020016107f9565b505050509050019250505060405180910390f35b34801561083157600080fd5b506103fc6004803603602081101561084857600080fd5b50356001600160a01b0316611880565b6104616004803603604081101561086e57600080fd5b506001600160a01b038135169060200135611895565b34801561089057600080fd5b50610461600480360360208110156108a757600080fd5b5035611ad3565b3480156108ba57600080fd5b50610461600480360360208110156108d157600080fd5b5035611bcb565b3480156108e457600080fd5b506104a2600480360360608110156108fb57600080fd5b6001600160a01b0382351691602081013591810190606081016040820135600160201b81111561092a57600080fd5b82018360208201111561093c57600080fd5b803590602001918460018302840111600160201b8311171561095d57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550611c71945050505050565b3480156109aa57600080fd5b50610461600480360360608110156109c157600080fd5b81359190810190604081016020820135600160201b8111156109e257600080fd5b8201836020820111156109f457600080fd5b803590602001918460018302840111600160201b83111715610a1557600080fd5b919390929091602081019035600160201b811115610a3257600080fd5b820183602082011115610a4457600080fd5b803590602001918460018302840111600160201b83111715610a6557600080fd5b509092509050612168565b348015610a7c57600080fd5b506107d5612479565b348015610a9157600080fd5b5061046160048036036020811015610aa857600080fd5b50356124d9565b348015610abb57600080fd5b506104256127a9565b348015610ad057600080fd5b5061046160048036036020811015610ae757600080fd5b810190602081018135600160201b811115610b0157600080fd5b820183602082011115610b1357600080fd5b803590602001918460208302840111600160201b83111715610b3457600080fd5b5090925090506127b5565b348015610b4b57600080fd5b506104616129db565b348015610b6057600080fd5b5061046160048036036020811015610b7757600080fd5b5035612ad9565b348015610b8a57600080fd5b50610425612bdd565b348015610b9f57600080fd5b50610ba8612be3565b604080516001600160a01b039092168252519081900360200190f35b348015610bd057600080fd5b50610425612bf2565b348015610be557600080fd5b5061046160048036036020811015610bfc57600080fd5b810190602081018135600160201b811115610c1657600080fd5b820183602082011115610c2857600080fd5b803590602001918460208302840111600160201b83111715610c4957600080fd5b509092509050612bfe565b348015610c6057600080fd5b50610425612f40565b348015610c7557600080fd5b50610ba8612f46565b348015610c8a57600080fd5b50610461612f55565b348015610c9f57600080fd5b50610425612fb2565b348015610cb457600080fd5b5061046160048036036020811015610ccb57600080fd5b810190602081018135600160201b811115610ce557600080fd5b820183602082011115610cf757600080fd5b803590602001918460018302840111600160201b83111715610d1857600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550612fb8945050505050565b348015610d6557600080fd5b506103fc6130f5565b348015610d7a57600080fd5b5061046160048036036020811015610d9157600080fd5b50356130fe565b348015610da457600080fd5b5061046160048036036040811015610dbb57600080fd5b506001600160a01b03813516906020013515156131ee565b348015610ddf57600080fd5b5061042560048036036040811015610df657600080fd5b506001600160a01b0381351690602001356133a8565b348015610e1857600080fd5b506103fc613438565b348015610e2d57600080fd5b506103fc613441565b348015610e4257600080fd5b5061046160048036036060811015610e5957600080fd5b506001600160a01b03813581169160208101359091169060400135613450565b348015610e8557600080fd5b506104256135da565b348015610e9a57600080fd5b5061046160048036036020811015610eb157600080fd5b50356135e0565b348015610ec457600080fd5b5061042561395d565b348015610ed957600080fd5b50610425613963565b348015610eee57600080fd5b5061042560048036036020811015610f0557600080fd5b810190602081018135600160201b811115610f1f57600080fd5b820183602082011115610f3157600080fd5b803590602001918460208302840111600160201b83111715610f5257600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929550613969945050505050565b348015610f9c57600080fd5b5061046160048036036020811015610fb357600080fd5b50356139c3565b348015610fc657600080fd5b50610ba860048036036020811015610fdd57600080fd5b5035613b6c565b348015610ff057600080fd5b50610425613b93565b34801561100557600080fd5b506103fc613b99565b34801561101a57600080fd5b50610425613ba7565b34801561102f57600080fd5b506104616004803603602081101561104657600080fd5b5035613bad565b34801561105957600080fd5b506104616004803603602081101561107057600080fd5b5035613cf7565b34801561108357600080fd5b506104256004803603604081101561109a57600080fd5b506001600160a01b038135169060200135613d50565b3480156110bc57600080fd5b50610461600480360360208110156110d357600080fd5b5035613f03565b3480156110e657600080fd5b50610461600480360360208110156110fd57600080fd5b5035613f5c565b34801561111057600080fd5b506104616004803603602081101561112757600080fd5b810190602081018135600160201b81111561114157600080fd5b82018360208201111561115357600080fd5b803590602001918460208302840111600160201b8311171561117457600080fd5b509092509050613fb5565b34801561118b57600080fd5b506103fc614307565b6001600160e01b031981166301ffc9a760e01b145b919050565b6017545b90565b6111be33614310565b806111c857503330145b61120c576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b8066038d7ea4c680001115801561122b57506706f05b59d3b200008111155b611272576040805162461bcd60e51b815260206004820152601360248201527206f7574206f662072616e676520746f702d757606c1b604482015290519081900360640190fd5b61128360038263ffffffff61432416565b604080513381526020810183905281517f41ff5d5ce3b7935893a4e7269ec5caae9cca5e3bf0eb4b21d2f443489667112e929181900390910190a150565b6112ca33614310565b806112d957506112d93361438d565b611323576040805162461bcd60e51b815260206004820152601660248201527537b7363c9037bbb732b93e3e31b7b73a3937b63632b960511b604482015290519081900360640190fd5b60135460ff16611372576040805162461bcd60e51b81526020600482015260156024820152743737903832b73234b7339039bab136b4b9b9b4b7b760591b604482015290519081900360640190fd5b6113d560118054806020026020016040519081016040528092919081815260200182805480156113cb57602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116113ad575b5050505050613969565b81146114125760405162461bcd60e51b81526004018080602001828103825260238152602001806154426023913960400191505060405180910390fd5b61141e601160006152b2565b6013805460ff19169055604080513381526020810183905281517f7794eff834d760583543e6e510e717a5e66d2c064e225f4db448343c3e66afcf929181900390910190a150565b60405180604001604052806005815260200164332e312e3160d81b81525081565b60008061149a848463ffffffff61442116565b90506114a581614310565b6114e3576040805162461bcd60e51b815260206004820152600a60248201526937b7363c9037bbb732b960b11b604482015290519081900360640190fd5b50630b135d3f60e11b90505b92915050565b6114fe33614310565b8061150857503330145b61154c576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b80611595576040805162461bcd60e51b8152602060048201526014602482015273617373657420617272617920697320656d70747960601b604482015290519081900360640190fd5b60005b818110156115f85760006115c7308585858181106115b257fe5b905060200201356001600160a01b031661450f565b90506115ef858585858181106115d957fe5b905060200201356001600160a01b031683613450565b50600101611598565b507fd4f62f23021706247dcffea245d104ae7ddaec7f23acf3d11d7136d5de6a69ad83838360405180846001600160a01b03166001600160a01b03168152602001806020018281038252848482818152602001925060200280828437600083820152604051601f909101601f1916909201829003965090945050505050a1505050565b6000611687600a6145ba565b905090565b6000808585604051602001808383808284376040805191909301818103601f190182528084528151602092830120601f8b01839004830282018301909452898152929650630b135d3f60e11b9550611703945086935089915088908190840183828082843760009201919091525061148792505050565b6001600160e01b0319161461174f576040805162461bcd60e51b815260206004820152600d60248201526c1cda59c81b9bdd081d985b1a59609a1b604482015290519081900360640190fd5b506320c13b0b60e01b95945050505050565b600254600160a01b900460ff1690565b61177a33614310565b8061178457503330145b6117c8576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b6117d960148263ffffffff6145ef16565b6040805182815290517f4b1b970c8a0fa761e7803ed70c13d7aca71904b13df60fbe03f981da1730da919181900360200190a150565b60035490565b60135460ff1681565b6060601280548060200260200160405190810160405280929190818152602001828054801561187657602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311611858575b5050505050905090565b600f6020526000908152604090205460ff1681565b61189e33614310565b806118a857503330145b6118ec576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b6118f582614650565b61193b576040805162461bcd60e51b8152602060048201526012602482015271746f6b656e206e6f74206c6f616461626c6560701b604482015290519081900360640190fd5b60006119478383613d50565b905061195a600a8263ffffffff61466a16565b6000611967601a546146e0565b90506001600160a01b03841615611a0f576119926001600160a01b038516828563ffffffff6147a216565b806001600160a01b0316631b3c96b485856040518363ffffffff1660e01b815260040180836001600160a01b03166001600160a01b0316815260200182815260200192505050600060405180830381600087803b1580156119f257600080fd5b505af1158015611a06573d6000803e3d6000fd5b50505050611a89565b806001600160a01b0316631b3c96b48486866040518463ffffffff1660e01b815260040180836001600160a01b03166001600160a01b03168152602001828152602001925050506000604051808303818588803b158015611a6f57600080fd5b505af1158015611a83573d6000803e3d6000fd5b50505050505b604080516001600160a01b03861681526020810185905281517f5f65674bec9af81f71be68674135a0ea3f163fb91984e3893d06da9f6ea2ce8a929181900390910190a150505050565b611adc33614310565b80611ae657503330145b611b2a576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b600954811115611b7c576040805162461bcd60e51b81526020600482015260186024820152771bdd5d081bd9881c985b99d9481b1bd85908185b5bdd5b9d60421b604482015290519081900360640190fd5b611b8d600a8263ffffffff61432416565b604080513381526020810183905281517f0b05243483e17c3f3377aee82b7d47e5700b48288695fc08b7ecc2759afa44ef929181900390910190a150565b611bd433614310565b80611bde57503330145b611c22576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b611c3360148263ffffffff61432416565b604080513381526020810183905281517f068f112e5ec923d412be64779fe69e0fcbb6784c6617e94cccc8fd348f2e0f21929181900390910190a150565b6060611c7c33614310565b80611c8657503330145b611cca576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b6001600160a01b0384166000908152600f602052604090205460ff16611cfb57611cfb60148463ffffffff61466a16565b611d0d846001600160a01b03166148ba565b8015611d1d5750611d1d846148c0565b15611f0457600080611d2f86856148da565b6001600160a01b0382166000908152600f6020526040902054919350915060ff16611d75576000611d6087836133a8565b9050611d7360148263ffffffff61466a16565b505b611d8e6001600160a01b0387168563ffffffff6149e416565b604080516020808252818301909252606091602082018180388339019050509050600160f81b81601f81518110611dc157fe5b60200101906001600160f81b031916908160001a9053507ff77753fab406ecfff96d6ff2476c64a838fa9f6d37b1bf190f8546e395e3b6138787878460405180856001600160a01b03166001600160a01b031681526020018481526020018060200180602001838103835285818151815260200191508051906020019080838360005b83811015611e5c578181015183820152602001611e44565b50505050905090810190601f168015611e895780820380516001836020036101000a031916815260200191505b50838103825284518152845160209182019186019080838360005b83811015611ebc578181015183820152602001611ea4565b50505050905090810190601f168015611ee95780820380516001836020036101000a031916815260200191505b50965050505050505060405180910390a19250612161915050565b60006060856001600160a01b031685856040518082805190602001908083835b60208310611f435780518252601f199092019160209182019101611f24565b6001836020036101000a03801982511681845116808217855250505050505090500191505060006040518083038185875af1925050503d8060008114611fa5576040519150601f19603f3d011682016040523d82523d6000602084013e611faa565b606091505b509150915081819061203a5760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b83811015611fff578181015183820152602001611fe7565b50505050905090810190601f16801561202c5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b507ff77753fab406ecfff96d6ff2476c64a838fa9f6d37b1bf190f8546e395e3b6138686868460405180856001600160a01b03166001600160a01b031681526020018481526020018060200180602001838103835285818151815260200191508051906020019080838360005b838110156120bf5781810151838201526020016120a7565b50505050905090810190601f1680156120ec5780820380516001836020036101000a031916815260200191505b50838103825284518152845160209182019186019080838360005b8381101561211f578181015183820152602001612107565b50505050905090810190601f16801561214c5780820380516001836020036101000a031916815260200191505b50965050505050505060405180910390a19150505b9392505050565b6121713361438d565b6121b0576040805162461bcd60e51b815260206004820152601a60248201526000805160206153b4833981519152604482015290519081900360640190fd5b600061220186868660405160200180806339363c1d60e11b81525060040184815260200183838082843780830192505050935050505060405160208183030381529060405280519060200120614ba2565b9050631626ba7e60e01b6001600160e01b0319166122558285858080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061148792505050565b6001600160e01b031916146122a1576040805162461bcd60e51b815260206004820152600d60248201526c1cda59c81b9bdd081d985b1a59609a1b604482015290519081900360640190fd5b60195486146122e3576040805162461bcd60e51b81526020600482015260096024820152687478207265706c617960b81b604482015290519081900360640190fd5b6122eb614bf3565b60006060306001600160a01b03168787604051808383808284376040519201945060009350909150508083038183865af19150503d806000811461234b576040519150601f19603f3d011682016040523d82523d6000602084013e612350565b606091505b50915091508181906123a35760405162461bcd60e51b8152602060048201818152835160248401528351909283926044909101919085019080838360008315611fff578181015183820152602001611fe7565b507f823dbcf2b7b0f265871963ca65ac033f6b4c71e0d82cd123d2ff23d752dc21c187878360405180806020018060200183810383528686828181526020019250808284376000838201819052601f909101601f191690920185810384528651815286516020918201939188019250908190849084905b8381101561243257818101518382015260200161241a565b50505050905090810190601f16801561245f5780820380516001836020036101000a031916815260200191505b509550505050505060405180910390a15050505050505050565b60606011805480602002602001604051908101604052809291908181526020018280548015611876576020028201919060005260206000209081546001600160a01b03168152600190910190602001808311611858575050505050905090565b6124e23361438d565b612521576040805162461bcd60e51b815260206004820152601a60248201526000805160206153b4833981519152604482015290519081900360640190fd5b60135460ff16612570576040805162461bcd60e51b81526020600482015260156024820152743737903832b73234b7339039bab136b4b9b9b4b7b760591b604482015290519081900360640190fd5b6125d160118054806020026020016040519081016040528092919081815260200182805480156113cb576020028201919060005260206000209081546001600160a01b031681526001909101906020018083116113ad575050505050613969565b811461260e5760405162461bcd60e51b81526004018080602001828103825260238152602001806154426023913960400191505060405180910390fd5b60005b6011548110156126f557600f60006011838154811061262c57fe5b60009182526020808320909101546001600160a01b0316835282019290925260400190205460ff166126ed576001600f60006011848154811061266b57fe5b6000918252602080832091909101546001600160a01b031683528201929092526040019020805460ff19169115159190911790556011805460109190839081106126b157fe5b60009182526020808320909101548354600181018555938352912090910180546001600160a01b0319166001600160a01b039092169190911790555b600101612611565b507fb2f6cccee7a369e23e293c25aa19bef80af11eb26deba3ea0f2a02783f752e4a33601160405180836001600160a01b03166001600160a01b0316815260200180602001828103825283818154815260200191508054801561278157602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311612763575b5050935050505060405180910390a161279c601160006152b2565b506013805460ff19169055565b600061168760146145ba565b6127be33614310565b806127c857503330145b61280c576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b60135460ff161580156128275750601354610100900460ff16155b612878576040805162461bcd60e51b815260206004820152601c60248201527f77686974656c6973742073756d62697373696f6e2070656e64696e6700000000604482015290519081900360640190fd5b60135462010000900460ff166128d1576040805162461bcd60e51b81526020600482015260196024820152781dda1a5d195b1a5cdd081b9bdd081a5b9a5d1a585b1a5e9959603a1b604482015290519081900360640190fd5b80612915576040805162461bcd60e51b815260206004820152600f60248201526e195b5c1d1e481dda1a5d195b1a5cdd608a1b604482015290519081900360640190fd5b612921601283836152d0565b506013805461ff00191661010017905560408051602080840282810182019093528382527ffbc0e5ca6c7e4858daf0fdb185ef5186203e74ec9c64737e93c0aeaec596e1d192859285926129909285918591829185019084908082843760009201919091525061396992505050565b60405180806020018381526020018281038252858582818152602001925060200280828437600083820152604051601f909101601f1916909201829003965090945050505050a15050565b6129e433614310565b612a2e576040805162461bcd60e51b815260206004820152601660248201527539b2b73232b91034b9903737ba1030b71037bbb732b960511b604482015290519081900360640190fd5b600254600160a01b900460ff16612a8c576040805162461bcd60e51b815260206004820152601d60248201527f6f776e657273686970206973206e6f74207472616e7366657261626c65000000604482015290519081900360640190fd5b600280546001600160a01b0319169055604080516000808252602082015281517f850b3df64837d7d518b45f5aa64d104652c3b80eb5b34a8e3d9eb666cb7cdea5929181900390910190a1565b612ae233614310565b80612aec57503330145b612b30576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b8066038d7ea4c6800011158015612b4f57506706f05b59d3b200008111155b612b96576040805162461bcd60e51b815260206004820152601360248201527206f7574206f662072616e676520746f702d757606c1b604482015290519081900360640190fd5b612ba760038263ffffffff6145ef16565b6040805182815290517faf2a77cd04c3cc155588dd3bf67b310ab4fb3b1da3cf6b8d7d4d2aa1d09b794c9181900360200190a150565b601a5490565b6000546001600160a01b031690565b600061168760036145ba565b612c0733614310565b80612c1157503330145b612c55576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b60135460ff16158015612c705750601354610100900460ff16155b612cc1576040805162461bcd60e51b815260206004820152601c60248201527f77686974656c6973742073756d62697373696f6e2070656e64696e6700000000604482015290519081900360640190fd5b8181808060200260200160405190810160405280939291908181526020018383602002808284376000920182905250925050505b8151811015612ddd57612d1a828281518110612d0d57fe5b6020026020010151614310565b15612d65576040805162461bcd60e51b8152602060048201526016602482015275636f6e7461696e73206f776e6572206164647265737360501b604482015290519081900360640190fd5b60006001600160a01b0316828281518110612d7c57fe5b60200260200101516001600160a01b03161415612dd5576040805162461bcd60e51b8152602060048201526012602482015271636f6e7461696e732030206164647265737360701b604482015290519081900360640190fd5b600101612cf5565b5060135462010000900460ff16612e37576040805162461bcd60e51b81526020600482015260196024820152781dda1a5d195b1a5cdd081b9bdd081a5b9a5d1a585b1a5e9959603a1b604482015290519081900360640190fd5b81612e7b576040805162461bcd60e51b815260206004820152600f60248201526e195b5c1d1e481dda1a5d195b1a5cdd608a1b604482015290519081900360640190fd5b612e87601184846152d0565b506013805460ff1916600117905560408051602080850282810182019093528482527f9c80b3b5f68b3e017766d59e8d09b34efe6462b05c398f35cab9e271d9bc3b9c9286928692612ef49285918591829185019084908082843760009201919091525061396992505050565b60405180806020018381526020018281038252858582818152602001925060200280828437600083820152604051601f909101601f1916909201829003965090945050505050a1505050565b60085490565b6002546001600160a01b031690565b612f5e33614310565b612fa8576040805162461bcd60e51b815260206004820152601660248201527539b2b73232b91034b9903737ba1030b71037bbb732b960511b604482015290519081900360640190fd5b612fb0614bf3565b565b60145490565b612fc133614310565b80612fcb57503330145b61300f576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b8051602080820191906000808060605b868510156130eb5761303886605463ffffffff614c3b16565b888601805160148201516034909201805193995060609190911c9650909450909250905061307d6054613071878563ffffffff614c9816565b9063ffffffff614c9816565b9450868511156130c4576040805162461bcd60e51b815260206004820152600d60248201526c6f7574206f6620626f756e647360981b604482015290519081900360640190fd5b816130da57506040805160208101909152600081525b6130e5848483611c71565b5061301f565b5050505050505050565b600e5460ff1690565b61310733614310565b8061311157503330145b613155576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b6009548111156131a7576040805162461bcd60e51b81526020600482015260186024820152771bdd5d081bd9881c985b99d9481b1bd85908185b5bdd5b9d60421b604482015290519081900360640190fd5b6131b8600a8263ffffffff6145ef16565b6040805182815290517fc178d379965e5657b6fc57494e392f121a14119215dfb422aad7db4cc03f2d109181900360200190a150565b6131f733614310565b613241576040805162461bcd60e51b815260206004820152601660248201527539b2b73232b91034b9903737ba1030b71037bbb732b960511b604482015290519081900360640190fd5b600254600160a01b900460ff1661329f576040805162461bcd60e51b815260206004820152601d60248201527f6f776e657273686970206973206e6f74207472616e7366657261626c65000000604482015290519081900360640190fd5b6001600160a01b0382166132e45760405162461bcd60e51b81526004018080602001828103825260238152602001806153f56023913960400191505060405180910390fd5b6002805460ff60a01b1916600160a01b831515021790558061333d57604080516001600160a01b038416815290517f808639ff9c8e4732d60b6c2330de498035416d229f27a77d259680895efec1229181900360200190a15b600254604080516001600160a01b039283168152918416602083015280517f850b3df64837d7d518b45f5aa64d104652c3b80eb5b34a8e3d9eb666cb7cdea59281900390910190a150600280546001600160a01b0319166001600160a01b0392909216919091179055565b6000806000806133b786614cf2565b50505093509350935050801561342c5781613402576040805162461bcd60e51b81526020600482015260066024820152650726174653d360d41b604482015290519081900360640190fd5b61342283613416878563ffffffff614e8416565b9063ffffffff614edd16565b93505050506114ef565b50600095945050505050565b60185460ff1690565b60135462010000900460ff1681565b61345933614310565b8061346357503330145b6134a7576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b80806134e4576040805162461bcd60e51b8152602060048201526007602482015266076616c75653d360cc1b604482015290519081900360640190fd5b6001600160a01b03841661352f576040805162461bcd60e51b815260206004820152600d60248201526c064657374696e6174696f6e3d3609c1b604482015290519081900360640190fd5b6001600160a01b0384166000908152600f602052604090205460ff1661357f57816001600160a01b0384161561356c5761356984846133a8565b90505b61357d60148263ffffffff61466a16565b505b61358a848484614f47565b604080516001600160a01b0380871682528516602082015280820184905290517fd1ba4ac2e2a11b5101f6cb4d978f514a155b421e8ec396d2d9abaf0bb02917ee9181900360600190a150505050565b600d5490565b6135e93361438d565b613628576040805162461bcd60e51b815260206004820152601a60248201526000805160206153b4833981519152604482015290519081900360640190fd5b601354610100900460ff1661367c576040805162461bcd60e51b81526020600482015260156024820152743737903832b73234b7339039bab136b4b9b9b4b7b760591b604482015290519081900360640190fd5b6136dd60128054806020026020016040519081016040528092919081815260200182805480156113cb576020028201919060005260206000209081546001600160a01b031681526001909101906020018083116113ad575050505050613969565b811461371a5760405162461bcd60e51b81526004018080602001828103825260238152602001806154426023913960400191505060405180910390fd5b60005b6012548110156138a857600f60006012838154811061373857fe5b60009182526020808320909101546001600160a01b0316835282019290925260400190205460ff16156138a0576000600f60006012848154811061377857fe5b6000918252602080832091909101546001600160a01b031683528201929092526040018120805460ff1916921515929092179091555b6010546137c290600163ffffffff614c3b16565b81101561388a57601282815481106137d657fe5b600091825260209091200154601080546001600160a01b0390921691839081106137fc57fe5b6000918252602090912001546001600160a01b031614156138825760108054600019810190811061382957fe5b600091825260209091200154601080546001600160a01b03909216918390811061384f57fe5b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555061388a565b6001016137ae565b50601080549061389e906000198301615333565b505b60010161371d565b507fd218c430fa348f4ce67791021b6b89c0c3eacd4ead1d8f5b83c60038ec28249b33601260405180836001600160a01b03166001600160a01b0316815260200180602001828103825283818154815260200191508054801561393457602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311613916575b5050935050505060405180910390a161394f601260006152b2565b506013805461ff0019169055565b60065490565b60195481565b60008160405160200180828051906020019060200280838360005b8381101561399c578181015183820152602001613984565b50505050905001915050604051602081830303815290604052805190602001209050919050565b6139cc33614310565b806139db57506139db3361438d565b613a25576040805162461bcd60e51b815260206004820152601660248201527537b7363c9037bbb732b93e3e31b7b73a3937b63632b960511b604482015290519081900360640190fd5b601354610100900460ff16613a79576040805162461bcd60e51b81526020600482015260156024820152743737903832b73234b7339039bab136b4b9b9b4b7b760591b604482015290519081900360640190fd5b613ada60128054806020026020016040519081016040528092919081815260200182805480156113cb576020028201919060005260206000209081546001600160a01b031681526001909101906020018083116113ad575050505050613969565b8114613b175760405162461bcd60e51b81526004018080602001828103825260238152602001806154426023913960400191505060405180910390fd5b613b23601260006152b2565b6013805461ff0019169055604080513381526020810183905281517f13c935eb475aa0f6e931fece83e2ac44569ce2d53460d29a6dedab40b965c8a3929181900390910190a150565b60108181548110613b7957fe5b6000918252602090912001546001600160a01b0316905081565b600a5490565b601354610100900460ff1681565b60015490565b8080613bea576040805162461bcd60e51b8152602060048201526007602482015266076616c75653d360cc1b604482015290519081900360640190fd5b613bf333614310565b80613c025750613c023361438d565b613c4c576040805162461bcd60e51b815260206004820152601660248201527537b7363c9037bbb732b93e3e31b7b73a3937b63632b960511b604482015290519081900360640190fd5b613c5d60038363ffffffff61466a16565b613c65612f46565b6001600160a01b03166108fc839081150290604051600060405180830381858888f19350505050158015613c9d573d6000803e3d6000fd5b507f611b7c0d84fda988026215bef9b3e4d81cbceced7e679be6d5e044b588467c0e33613cc8612f46565b604080516001600160a01b03938416815291909216602082015280820185905290519081900360600190a15050565b613d003361438d565b613d3f576040805162461bcd60e51b815260206004820152601a60248201526000805160206153b4833981519152604482015290519081900360640190fd5b611c3360148263ffffffff614fab16565b6000613d5a614fff565b6001600160a01b0316836001600160a01b03161415613d7a5750806114ef565b816001600160a01b03841615613e3f576000806000613d9887614cf2565b5050509350935093505080613dea576040805162461bcd60e51b8152602060048201526013602482015272746f6b656e206e6f7420617661696c61626c6560681b604482015290519081900360640190fd5b81613e25576040805162461bcd60e51b81526020600482015260066024820152650726174653d360d41b604482015290519081900360640190fd5b613e3983613416888563ffffffff614e8416565b93505050505b6000806000613e4c615075565b5050509350935093505080613e9e576040805162461bcd60e51b8152602060048201526013602482015272746f6b656e206e6f7420617661696c61626c6560681b604482015290519081900360640190fd5b81613ee4576040805162461bcd60e51b81526020600482015260116024820152700737461626c65636f696e20726174653d3607c1b604482015290519081900360640190fd5b613ef882613416868663ffffffff614e8416565b979650505050505050565b613f0c3361438d565b613f4b576040805162461bcd60e51b815260206004820152601a60248201526000805160206153b4833981519152604482015290519081900360640190fd5b611b8d600a8263ffffffff614fab16565b613f653361438d565b613fa4576040805162461bcd60e51b815260206004820152601a60248201526000805160206153b4833981519152604482015290519081900360640190fd5b61128360038263ffffffff614fab16565b613fbe33614310565b80613fc857503330145b61400c576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b8181808060200260200160405190810160405280939291908181526020018383602002808284376000920182905250925050505b815181101561411b57614058828281518110612d0d57fe5b156140a3576040805162461bcd60e51b8152602060048201526016602482015275636f6e7461696e73206f776e6572206164647265737360501b604482015290519081900360640190fd5b60006001600160a01b03168282815181106140ba57fe5b60200260200101516001600160a01b03161415614113576040805162461bcd60e51b8152602060048201526012602482015271636f6e7461696e732030206164647265737360701b604482015290519081900360640190fd5b600101614040565b5060135462010000900460ff1615614172576040805162461bcd60e51b81526020600482015260156024820152741dda1a5d195b1a5cdd081a5b9a5d1a585b1a5e9959605a1b604482015290519081900360640190fd5b60005b8281101561426357600f600085858481811061418d57fe5b602090810292909201356001600160a01b03168352508101919091526040016000205460ff1661425b576001600f60008686858181106141c957fe5b905060200201356001600160a01b03166001600160a01b03166001600160a01b0316815260200190815260200160002060006101000a81548160ff021916908315150217905550601084848381811061421e57fe5b835460018101855560009485526020948590200180546001600160a01b0319166001600160a01b0395909202939093013593909316929092179055505b600101614175565b506013805462ff0000191662010000179055604080513380825260208201838152601080549484018590527fb2f6cccee7a369e23e293c25aa19bef80af11eb26deba3ea0f2a02783f752e4a94929390929091906060830190849080156142f357602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116142d5575b5050935050505060405180910390a1505050565b60075460ff1690565b6002546001600160a01b0390811691161490565b600482015460ff1615614372576040805162461bcd60e51b81526020600482015260116024820152701b1a5b5a5d08185b1c9958591e481cd95d607a1b604482015290519081900360640190fd5b61437c82826151e4565b50600401805460ff19166001179055565b600061439a6001546146e0565b6001600160a01b031663b429afeb836040518263ffffffff1660e01b815260040180826001600160a01b03166001600160a01b0316815260200191505060206040518083038186803b1580156143ef57600080fd5b505afa158015614403573d6000803e3d6000fd5b505050506040513d602081101561441957600080fd5b505192915050565b60008151604114614434575060006114ef565b60208201516040830151606084015160001a7f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a082111561447a57600093505050506114ef565b8060ff16601b1415801561449257508060ff16601c14155b156144a357600093505050506114ef565b6040805160008152602080820180845289905260ff8416828401526060820186905260808201859052915160019260a0808401939192601f1981019281900390910190855afa1580156144fa573d6000803e3d6000fd5b5050604051601f190151979650505050505050565b60006001600160a01b038216156145a957816001600160a01b03166370a08231846040518263ffffffff1660e01b815260040180826001600160a01b03166001600160a01b0316815260200191505060206040518083038186803b15801561457657600080fd5b505afa15801561458a573d6000803e3d6000fd5b505050506040513d60208110156145a057600080fd5b505190506114ef565b506001600160a01b038216316114ef565b60028101546000906145d5906201518063ffffffff614c9816565b4211156145e4575080546111a9565b5060018101546111a9565b600482015460ff16614648576040805162461bcd60e51b815260206004820152601960248201527f6c696d6974206861736e2774206265656e207365742079657400000000000000604482015290519081900360640190fd5b600390910155565b60008061465c83614cf2565b509098975050505050505050565b61467382615207565b80826001015410156146bf576040805162461bcd60e51b815260206004820152601060248201526f185d985a5b18589b194f185b5bdd5b9d60821b604482015290519081900360640190fd5b60018201546146d4908263ffffffff614c3b16565b82600101819055505050565b6000805460408051630178b8bf60e01b81526004810185905290516001600160a01b0390921691630178b8bf91602480820192602092909190829003018186803b15801561472d57600080fd5b505afa158015614741573d6000803e3d6000fd5b505050506040513d602081101561475757600080fd5b505160408051631d9dabef60e11b81526004810185905290516001600160a01b0390921691633b3b57de91602480820192602092909190829003018186803b1580156143ef57600080fd5b801580614828575060408051636eb1769f60e11b81523060048201526001600160a01b03848116602483015291519185169163dd62ed3e91604480820192602092909190829003018186803b1580156147fa57600080fd5b505afa15801561480e573d6000803e3d6000fd5b505050506040513d602081101561482457600080fd5b5051155b6148635760405162461bcd60e51b81526004018080602001828103825260368152602001806154656036913960400191505060405180910390fd5b604080516001600160a01b038416602482015260448082018490528251808303909101815260649091019091526020810180516001600160e01b031663095ea7b360e01b1790526148b59084906149e4565b505050565b3b151590565b6000806148cc83614cf2565b509198975050505050505050565b6000806148e86008546146e0565b6001600160a01b031663afc72e9385856040518363ffffffff1660e01b815260040180836001600160a01b03166001600160a01b0316815260200180602001828103825283818151815260200191508051906020019080838360005b8381101561495c578181015183820152602001614944565b50505050905090810190601f1680156149895780820380516001836020036101000a031916815260200191505b509350505050604080518083038186803b1580156149a657600080fd5b505afa1580156149ba573d6000803e3d6000fd5b505050506040513d60408110156149d057600080fd5b508051602090910151909590945092505050565b6149f6826001600160a01b03166148ba565b614a47576040805162461bcd60e51b815260206004820152601f60248201527f5361666545524332303a2063616c6c20746f206e6f6e2d636f6e747261637400604482015290519081900360640190fd5b60006060836001600160a01b0316836040518082805190602001908083835b60208310614a855780518252601f199092019160209182019101614a66565b6001836020036101000a0380198251168184511680821785525050505050509050019150506000604051808303816000865af19150503d8060008114614ae7576040519150601f19603f3d011682016040523d82523d6000602084013e614aec565b606091505b509150915081614b43576040805162461bcd60e51b815260206004820181905260248201527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564604482015290519081900360640190fd5b805115614b9c57808060200190516020811015614b5f57600080fd5b5051614b9c5760405162461bcd60e51b815260040180806020018281038252602a815260200180615418602a913960400191505060405180910390fd5b50505050565b604080517f19457468657265756d205369676e6564204d6573736167653a0a333200000000602080830191909152603c8083019490945282518083039094018452605c909101909152815191012090565b601980546001019081905560408051338152602081019290925280517fab0423a75986556234aecd171c46ce7f5e45607d8070bf5230f2735b50322bff9281900390910190a1565b600082821115614c92576040805162461bcd60e51b815260206004820152601e60248201527f536166654d6174683a207375627472616374696f6e206f766572666c6f770000604482015290519081900360640190fd5b50900390565b600082820183811015612161576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b6060600080600080600080614d086008546146e0565b6001600160a01b0316631f69565f896040518263ffffffff1660e01b815260040180826001600160a01b03166001600160a01b0316815260200191505060006040518083038186803b158015614d5d57600080fd5b505afa158015614d71573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405260e0811015614d9a57600080fd5b8101908080516040519392919084600160201b821115614db957600080fd5b908301906020820185811115614dce57600080fd5b8251600160201b811182820188101715614de757600080fd5b82525081516020918201929091019080838360005b83811015614e14578181015183820152602001614dfc565b50505050905090810190601f168015614e415780820380516001836020036101000a031916815260200191505b5060409081526020820151908201516060830151608084015160a085015160c090950151979e50929c50909a509850965094509192505050919395979092949650565b600082614e93575060006114ef565b82820282848281614ea057fe5b04146121615760405162461bcd60e51b81526004018080602001828103825260218152602001806153d46021913960400191505060405180910390fd5b6000808211614f33576040805162461bcd60e51b815260206004820152601a60248201527f536166654d6174683a206469766973696f6e206279207a65726f000000000000604482015290519081900360640190fd5b6000828481614f3e57fe5b04949350505050565b6001600160a01b038216614f91576040516001600160a01b0384169082156108fc029083906000818181858888f19350505050158015614f8b573d6000803e3d6000fd5b506148b5565b6148b56001600160a01b038316848363ffffffff61526016565b80826003015414614fed5760405162461bcd60e51b81526004018080602001828103825260228152602001806153926022913960400191505060405180910390fd5b614ffb8283600301546151e4565b5050565b600061500c6008546146e0565b6001600160a01b031663e9cbd8226040518163ffffffff1660e01b815260040160206040518083038186803b15801561504457600080fd5b505afa158015615058573d6000803e3d6000fd5b505050506040513d602081101561506e57600080fd5b5051905090565b606060008060008060008061508b6008546146e0565b6001600160a01b0316633efec5e96040518163ffffffff1660e01b815260040160006040518083038186803b1580156150c357600080fd5b505afa1580156150d7573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405260e081101561510057600080fd5b8101908080516040519392919084600160201b82111561511f57600080fd5b90830190602082018581111561513457600080fd5b8251600160201b81118282018810171561514d57600080fd5b82525081516020918201929091019080838360005b8381101561517a578181015183820152602001615162565b50505050905090810190601f1680156151a75780820380516001836020036101000a031916815260200191505b5060409081526020820151908201516060830151608084015160a085015160c090950151979f939e50919c509a5098509096509294509192505050565b6151ed82615207565b8082556001820154811015614ffb57815460018301555050565b600281015461521f906201518063ffffffff614c9816565b42111561525d57426002820155805460018201556040517fe93bc25276d408d390778e7a8b926f2f67209c43ed540081b951fe128f0d3cd290600090a15b50565b604080516001600160a01b038416602482015260448082018490528251808303909101815260649091019091526020810180516001600160e01b031663a9059cbb60e01b1790526148b59084906149e4565b508054600082559060005260206000209081019061525d9190615353565b828054828255906000526020600020908101928215615323579160200282015b828111156153235781546001600160a01b0319166001600160a01b038435161782556020909201916001909101906152f0565b5061532f92915061536d565b5090565b8154818355818111156148b5576000838152602090206148b59181019083015b6111b291905b8082111561532f5760008155600101615359565b6111b291905b8082111561532f5780546001600160a01b031916815560010161537356fe636f6e6669726d65642f7375626d6974746564206c696d6974206d69736d6174636873656e646572206973206e6f74206120636f6e74726f6c6c6572000000000000536166654d6174683a206d756c7469706c69636174696f6e206f766572666c6f776f776e65722063616e6e6f742062652073657420746f207a65726f20616464726573735361666545524332303a204552433230206f7065726174696f6e20646964206e6f7420737563636565646e6f6e2d6d61746368696e672070656e64696e672077686974656c69737420686173685361666545524332303a20617070726f76652066726f6d206e6f6e2d7a65726f20746f206e6f6e2d7a65726f20616c6c6f77616e6365a265627a7a72315820a7321f3cf8984c18d49b59bfb92e2c20d256e7ae3b1df51d97e62da513f6daf164736f6c63430005110032"
||||||| constructed merge base
var WalletBin = "0x60806040527f7f2ce995617d2816b426c5c8698c5ec2952f7a34bb10f38326f74933d58936976002553480156200003557600080fd5b5060405162005a4738038062005a47833981810160405260e08110156200005b57600080fd5b508051602082015160408301516060840151608085015160a086015160c090960151600180546001600160a01b038087166001600160a01b0319928316179283905560008054909216921691909117905594959394929391929091908084808989878015620000ca5760028190555b50600380546001600160a01b0319166001600160a01b0384161760ff60a01b1916600160a01b8315158102919091179182905560ff9104166200014457604080516001600160a01b038416815290517f808639ff9c8e4732d60b6c2330de498035416d229f27a77d259680895efec1229181900360200190a15b60408051600081526001600160a01b038416602082015281517f850b3df64837d7d518b45f5aa64d104652c3b80eb5b34a8e3d9eb666cb7cdea5929181900390910190a150506040805160a0810182526706f05b59d3b2000080825260208201819052429282018390526000606083018190526080909201829052600481905560055560069190915560078190556008805460ff19169055600991909155620001f56001600160e01b03620002ec16565b50505050509150506000811162000243576040805162461bcd60e51b815260206004820152600d60248201526c37379039ba30b13632b1b7b4b760991b604482015290519081900360640190fd5b61271002600a8190556040805160a08082018352838252602080830185905242838501819052600060608086018290526080958601829052600b889055600c97909755600d829055600e819055600f805460ff1990811690915586519485018752898552928401899052948301819052948201849052910182905260158590556016949094556017919091556018555060198054909116905550601b5550620005689350505050565b60606000806000806000806200030a6009546200046e60201b60201c565b6001600160a01b0316633efec5e96040518163ffffffff1660e01b815260040160006040518083038186803b1580156200034357600080fd5b505afa15801562000358573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405260e08110156200038257600080fd5b8101908080516040519392919084640100000000821115620003a357600080fd5b908301906020820185811115620003b957600080fd5b8251640100000000811182820188101715620003d457600080fd5b82525081516020918201929091019080838360005b8381101562000403578181015183820152602001620003e9565b50505050905090810190601f168015620004315780820380516001836020036101000a031916815260200191505b5060409081526020820151908201516060830151608084015160a085015160c090950151979f939e50919c509a5098509096509294509192505050565b6000805460408051630178b8bf60e01b81526004810185905290516001600160a01b0390921691630178b8bf91602480820192602092909190829003018186803b158015620004bc57600080fd5b505afa158015620004d1573d6000803e3d6000fd5b505050506040513d6020811015620004e857600080fd5b505160408051631d9dabef60e11b81526004810185905290516001600160a01b0390921691633b3b57de91602480820192602092909190829003018186803b1580156200053457600080fd5b505afa15801562000549573d6000803e3d6000fd5b505050506040513d60208110156200056057600080fd5b505192915050565b6154cf80620005786000396000f3fe60806040526004361061038c5760003560e01c80637fd004fa116101dc578063cc0e7e5611610102578063e2b4ce97116100a0578063f40b51f81161006f578063f40b51f8146110b0578063f41c4319146110da578063f421764814611104578063f776f5181461117f5761038c565b8063e2b4ce971461100e578063e61c51ca14611023578063eadd3cea1461104d578063f36febda146110775761038c565b8063ce0b5bd5116100dc578063ce0b5bd514610f90578063d251fefc14610fba578063da84b1ed14610fe4578063de212bf314610ff95761038c565b8063cc0e7e5614610eb8578063cccdc55614610ecd578063cd7958dd14610ee25761038c565b8063b221f3161161017a578063be40ba7911610149578063be40ba7914610e21578063beabacc814610e36578063c4856cd914610e79578063cbd2ac6814610e8e5761038c565b8063b221f31614610d6e578063b242e53414610d98578063b87e21ef14610dd3578063bcb8b74a14610e0c5761038c565b806390e690c7116101b657806390e690c714610c7e5780639b0dfd2714610c93578063aaf1fc6214610ca8578063ab20599314610d595761038c565b80637fd004fa14610bd9578063877337b014610c545780638da5cb5b14610c695761038c565b806332531c3c116102c15780635adc02ab1161025f57806374624c551161022e57806374624c5514610b54578063747c31d614610b7e5780637d73b23114610b935780637d7d004614610bc45761038c565b80635adc02ab14610a855780635d2362a814610aaf5780636137d67014610ac4578063715018a614610b3f5761038c565b80633c672eb71161029b5780633c672eb7146108ae5780633f579f42146108d857806346efe0ed1461099e57806347b55a9d14610a705761038c565b806332531c3c146108255780633a43199f146108585780633bfec254146108845761038c565b80631efd02991161032e57806321ce918d1161030857806321ce918d1461076c5780632587a6a21461079657806326d05ab2146107ab578063294f4025146107c05761038c565b80631efd02991461067757806320c13b0b1461068c5780632121dc75146107575761038c565b8063100f23fd1161036a578063100f23fd146104635780631127b57e1461048d5780631626ba7e146105175780631aa21fba146105ec5761038c565b806301ffc9a7146103c8578063027ef3eb146104105780630f3a85d814610437575b6040805133815234602082015281517f88a5966d370b9919b20f3e2c13ff65706f196a4e32cc2c12bf57088f88525874929181900390910190a1005b3480156103d457600080fd5b506103fc600480360360208110156103eb57600080fd5b50356001600160e01b031916611194565b604080519115158252519081900360200190f35b34801561041c57600080fd5b506104256111ae565b60408051918252519081900360200190f35b34801561044357600080fd5b506104616004803603602081101561045a57600080fd5b50356111b5565b005b34801561046f57600080fd5b506104616004803603602081101561048657600080fd5b50356112c1565b34801561049957600080fd5b506104a2611466565b6040805160208082528351818301528351919283929083019185019080838360005b838110156104dc5781810151838201526020016104c4565b50505050905090810190601f1680156105095780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561052357600080fd5b506105cf6004803603604081101561053a57600080fd5b81359190810190604081016020820135600160201b81111561055b57600080fd5b82018360208201111561056d57600080fd5b803590602001918460018302840111600160201b8311171561058e57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550611487945050505050565b604080516001600160e01b03199092168252519081900360200190f35b3480156105f857600080fd5b506104616004803603604081101561060f57600080fd5b6001600160a01b038235169190810190604081016020820135600160201b81111561063957600080fd5b82018360208201111561064b57600080fd5b803590602001918460208302840111600160201b8311171561066c57600080fd5b5090925090506114f5565b34801561068357600080fd5b5061042561167b565b34801561069857600080fd5b506105cf600480360360408110156106af57600080fd5b810190602081018135600160201b8111156106c957600080fd5b8201836020820111156106db57600080fd5b803590602001918460018302840111600160201b831117156106fc57600080fd5b919390929091602081019035600160201b81111561071957600080fd5b82018360208201111561072b57600080fd5b803590602001918460018302840111600160201b8311171561074c57600080fd5b50909250905061168c565b34801561076357600080fd5b506103fc611761565b34801561077857600080fd5b506104616004803603602081101561078f57600080fd5b5035611771565b3480156107a257600080fd5b5061042561180f565b3480156107b757600080fd5b506103fc611815565b3480156107cc57600080fd5b506107d561181e565b60408051602080825283518183015283519192839290830191858101910280838360005b838110156108115781810151838201526020016107f9565b505050509050019250505060405180910390f35b34801561083157600080fd5b506103fc6004803603602081101561084857600080fd5b50356001600160a01b0316611880565b6104616004803603604081101561086e57600080fd5b506001600160a01b038135169060200135611895565b34801561089057600080fd5b50610461600480360360208110156108a757600080fd5b5035611ad3565b3480156108ba57600080fd5b50610461600480360360208110156108d157600080fd5b5035611bcb565b3480156108e457600080fd5b506104a2600480360360608110156108fb57600080fd5b6001600160a01b0382351691602081013591810190606081016040820135600160201b81111561092a57600080fd5b82018360208201111561093c57600080fd5b803590602001918460018302840111600160201b8311171561095d57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550611c71945050505050565b3480156109aa57600080fd5b50610461600480360360608110156109c157600080fd5b81359190810190604081016020820135600160201b8111156109e257600080fd5b8201836020820111156109f457600080fd5b803590602001918460018302840111600160201b83111715610a1557600080fd5b919390929091602081019035600160201b811115610a3257600080fd5b820183602082011115610a4457600080fd5b803590602001918460018302840111600160201b83111715610a6557600080fd5b509092509050612168565b348015610a7c57600080fd5b506107d5612479565b348015610a9157600080fd5b5061046160048036036020811015610aa857600080fd5b50356124d9565b348015610abb57600080fd5b506104256127a9565b348015610ad057600080fd5b5061046160048036036020811015610ae757600080fd5b810190602081018135600160201b811115610b0157600080fd5b820183602082011115610b1357600080fd5b803590602001918460208302840111600160201b83111715610b3457600080fd5b5090925090506127b5565b348015610b4b57600080fd5b506104616129db565b348015610b6057600080fd5b5061046160048036036020811015610b7757600080fd5b5035612ad9565b348015610b8a57600080fd5b50610425612bdd565b348015610b9f57600080fd5b50610ba8612be3565b604080516001600160a01b039092168252519081900360200190f35b348015610bd057600080fd5b50610425612bf2565b348015610be557600080fd5b5061046160048036036020811015610bfc57600080fd5b810190602081018135600160201b811115610c1657600080fd5b820183602082011115610c2857600080fd5b803590602001918460208302840111600160201b83111715610c4957600080fd5b509092509050612bfe565b348015610c6057600080fd5b50610425612f40565b348015610c7557600080fd5b50610ba8612f46565b348015610c8a57600080fd5b50610461612f55565b348015610c9f57600080fd5b50610425612fb2565b348015610cb457600080fd5b5061046160048036036020811015610ccb57600080fd5b810190602081018135600160201b811115610ce557600080fd5b820183602082011115610cf757600080fd5b803590602001918460018302840111600160201b83111715610d1857600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550612fb8945050505050565b348015610d6557600080fd5b506103fc6130f5565b348015610d7a57600080fd5b5061046160048036036020811015610d9157600080fd5b50356130fe565b348015610da457600080fd5b5061046160048036036040811015610dbb57600080fd5b506001600160a01b03813516906020013515156131ee565b348015610ddf57600080fd5b5061042560048036036040811015610df657600080fd5b506001600160a01b0381351690602001356133a8565b348015610e1857600080fd5b506103fc613438565b348015610e2d57600080fd5b506103fc613441565b348015610e4257600080fd5b5061046160048036036060811015610e5957600080fd5b506001600160a01b03813581169160208101359091169060400135613450565b348015610e8557600080fd5b506104256135da565b348015610e9a57600080fd5b5061046160048036036020811015610eb157600080fd5b50356135e0565b348015610ec457600080fd5b5061042561395d565b348015610ed957600080fd5b50610425613963565b348015610eee57600080fd5b5061042560048036036020811015610f0557600080fd5b810190602081018135600160201b811115610f1f57600080fd5b820183602082011115610f3157600080fd5b803590602001918460208302840111600160201b83111715610f5257600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929550613969945050505050565b348015610f9c57600080fd5b5061046160048036036020811015610fb357600080fd5b50356139c3565b348015610fc657600080fd5b50610ba860048036036020811015610fdd57600080fd5b5035613b6c565b348015610ff057600080fd5b50610425613b93565b34801561100557600080fd5b506103fc613b99565b34801561101a57600080fd5b50610425613ba7565b34801561102f57600080fd5b506104616004803603602081101561104657600080fd5b5035613bad565b34801561105957600080fd5b506104616004803603602081101561107057600080fd5b5035613cf7565b34801561108357600080fd5b506104256004803603604081101561109a57600080fd5b506001600160a01b038135169060200135613d50565b3480156110bc57600080fd5b50610461600480360360208110156110d357600080fd5b5035613f03565b3480156110e657600080fd5b50610461600480360360208110156110fd57600080fd5b5035613f5c565b34801561111057600080fd5b506104616004803603602081101561112757600080fd5b810190602081018135600160201b81111561114157600080fd5b82018360208201111561115357600080fd5b803590602001918460208302840111600160201b8311171561117457600080fd5b509092509050613fb5565b34801561118b57600080fd5b506103fc614307565b6001600160e01b031981166301ffc9a760e01b145b919050565b6018545b90565b6111be33614310565b806111c857503330145b61120c576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b8066038d7ea4c680001115801561122b57506706f05b59d3b200008111155b611272576040805162461bcd60e51b815260206004820152601360248201527206f7574206f662072616e676520746f702d757606c1b604482015290519081900360640190fd5b61128360048263ffffffff61432416565b604080513381526020810183905281517f41ff5d5ce3b7935893a4e7269ec5caae9cca5e3bf0eb4b21d2f443489667112e929181900390910190a150565b6112ca33614310565b806112d957506112d93361438d565b611323576040805162461bcd60e51b815260206004820152601660248201527537b7363c9037bbb732b93e3e31b7b73a3937b63632b960511b604482015290519081900360640190fd5b60145460ff16611372576040805162461bcd60e51b81526020600482015260156024820152743737903832b73234b7339039bab136b4b9b9b4b7b760591b604482015290519081900360640190fd5b6113d560128054806020026020016040519081016040528092919081815260200182805480156113cb57602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116113ad575b5050505050613969565b81146114125760405162461bcd60e51b81526004018080602001828103825260238152602001806154426023913960400191505060405180910390fd5b61141e601260006152b2565b6014805460ff19169055604080513381526020810183905281517f7794eff834d760583543e6e510e717a5e66d2c064e225f4db448343c3e66afcf929181900390910190a150565b60405180604001604052806005815260200164332e312e3160d81b81525081565b60008061149a848463ffffffff61442116565b90506114a581614310565b6114e3576040805162461bcd60e51b815260206004820152600a60248201526937b7363c9037bbb732b960b11b604482015290519081900360640190fd5b50630b135d3f60e11b90505b92915050565b6114fe33614310565b8061150857503330145b61154c576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b80611595576040805162461bcd60e51b8152602060048201526014602482015273617373657420617272617920697320656d70747960601b604482015290519081900360640190fd5b60005b818110156115f85760006115c7308585858181106115b257fe5b905060200201356001600160a01b031661450f565b90506115ef858585858181106115d957fe5b905060200201356001600160a01b031683613450565b50600101611598565b507fd4f62f23021706247dcffea245d104ae7ddaec7f23acf3d11d7136d5de6a69ad83838360405180846001600160a01b03166001600160a01b03168152602001806020018281038252848482818152602001925060200280828437600083820152604051601f909101601f1916909201829003965090945050505050a1505050565b6000611687600b6145ba565b905090565b6000808585604051602001808383808284376040805191909301818103601f190182528084528151602092830120601f8b01839004830282018301909452898152929650630b135d3f60e11b9550611703945086935089915088908190840183828082843760009201919091525061148792505050565b6001600160e01b0319161461174f576040805162461bcd60e51b815260206004820152600d60248201526c1cda59c81b9bdd081d985b1a59609a1b604482015290519081900360640190fd5b506320c13b0b60e01b95945050505050565b600354600160a01b900460ff1690565b61177a33614310565b8061178457503330145b6117c8576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b6117d960158263ffffffff6145ef16565b6040805182815290517f4b1b970c8a0fa761e7803ed70c13d7aca71904b13df60fbe03f981da1730da919181900360200190a150565b60045490565b60145460ff1681565b6060601380548060200260200160405190810160405280929190818152602001828054801561187657602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311611858575b5050505050905090565b60106020526000908152604090205460ff1681565b61189e33614310565b806118a857503330145b6118ec576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b6118f582614650565b61193b576040805162461bcd60e51b8152602060048201526012602482015271746f6b656e206e6f74206c6f616461626c6560701b604482015290519081900360640190fd5b60006119478383613d50565b905061195a600b8263ffffffff61466a16565b6000611967601b546146e0565b90506001600160a01b03841615611a0f576119926001600160a01b038516828563ffffffff6147a216565b806001600160a01b0316631b3c96b485856040518363ffffffff1660e01b815260040180836001600160a01b03166001600160a01b0316815260200182815260200192505050600060405180830381600087803b1580156119f257600080fd5b505af1158015611a06573d6000803e3d6000fd5b50505050611a89565b806001600160a01b0316631b3c96b48486866040518463ffffffff1660e01b815260040180836001600160a01b03166001600160a01b03168152602001828152602001925050506000604051808303818588803b158015611a6f57600080fd5b505af1158015611a83573d6000803e3d6000fd5b50505050505b604080516001600160a01b03861681526020810185905281517f5f65674bec9af81f71be68674135a0ea3f163fb91984e3893d06da9f6ea2ce8a929181900390910190a150505050565b611adc33614310565b80611ae657503330145b611b2a576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b600a54811115611b7c576040805162461bcd60e51b81526020600482015260186024820152771bdd5d081bd9881c985b99d9481b1bd85908185b5bdd5b9d60421b604482015290519081900360640190fd5b611b8d600b8263ffffffff61432416565b604080513381526020810183905281517f0b05243483e17c3f3377aee82b7d47e5700b48288695fc08b7ecc2759afa44ef929181900390910190a150565b611bd433614310565b80611bde57503330145b611c22576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b611c3360158263ffffffff61432416565b604080513381526020810183905281517f068f112e5ec923d412be64779fe69e0fcbb6784c6617e94cccc8fd348f2e0f21929181900390910190a150565b6060611c7c33614310565b80611c8657503330145b611cca576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b6001600160a01b03841660009081526010602052604090205460ff16611cfb57611cfb60158463ffffffff61466a16565b611d0d846001600160a01b03166148ba565b8015611d1d5750611d1d846148c0565b15611f0457600080611d2f86856148da565b6001600160a01b038216600090815260106020526040902054919350915060ff16611d75576000611d6087836133a8565b9050611d7360158263ffffffff61466a16565b505b611d8e6001600160a01b0387168563ffffffff6149e416565b604080516020808252818301909252606091602082018180388339019050509050600160f81b81601f81518110611dc157fe5b60200101906001600160f81b031916908160001a9053507ff77753fab406ecfff96d6ff2476c64a838fa9f6d37b1bf190f8546e395e3b6138787878460405180856001600160a01b03166001600160a01b031681526020018481526020018060200180602001838103835285818151815260200191508051906020019080838360005b83811015611e5c578181015183820152602001611e44565b50505050905090810190601f168015611e895780820380516001836020036101000a031916815260200191505b50838103825284518152845160209182019186019080838360005b83811015611ebc578181015183820152602001611ea4565b50505050905090810190601f168015611ee95780820380516001836020036101000a031916815260200191505b50965050505050505060405180910390a19250612161915050565b60006060856001600160a01b031685856040518082805190602001908083835b60208310611f435780518252601f199092019160209182019101611f24565b6001836020036101000a03801982511681845116808217855250505050505090500191505060006040518083038185875af1925050503d8060008114611fa5576040519150601f19603f3d011682016040523d82523d6000602084013e611faa565b606091505b509150915081819061203a5760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b83811015611fff578181015183820152602001611fe7565b50505050905090810190601f16801561202c5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b507ff77753fab406ecfff96d6ff2476c64a838fa9f6d37b1bf190f8546e395e3b6138686868460405180856001600160a01b03166001600160a01b031681526020018481526020018060200180602001838103835285818151815260200191508051906020019080838360005b838110156120bf5781810151838201526020016120a7565b50505050905090810190601f1680156120ec5780820380516001836020036101000a031916815260200191505b50838103825284518152845160209182019186019080838360005b8381101561211f578181015183820152602001612107565b50505050905090810190601f16801561214c5780820380516001836020036101000a031916815260200191505b50965050505050505060405180910390a19150505b9392505050565b6121713361438d565b6121b0576040805162461bcd60e51b815260206004820152601a60248201526000805160206153b4833981519152604482015290519081900360640190fd5b600061220186868660405160200180806339363c1d60e11b81525060040184815260200183838082843780830192505050935050505060405160208183030381529060405280519060200120614ba2565b9050631626ba7e60e01b6001600160e01b0319166122558285858080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061148792505050565b6001600160e01b031916146122a1576040805162461bcd60e51b815260206004820152600d60248201526c1cda59c81b9bdd081d985b1a59609a1b604482015290519081900360640190fd5b601a5486146122e3576040805162461bcd60e51b81526020600482015260096024820152687478207265706c617960b81b604482015290519081900360640190fd5b6122eb614bf3565b60006060306001600160a01b03168787604051808383808284376040519201945060009350909150508083038183865af19150503d806000811461234b576040519150601f19603f3d011682016040523d82523d6000602084013e612350565b606091505b50915091508181906123a35760405162461bcd60e51b8152602060048201818152835160248401528351909283926044909101919085019080838360008315611fff578181015183820152602001611fe7565b507f823dbcf2b7b0f265871963ca65ac033f6b4c71e0d82cd123d2ff23d752dc21c187878360405180806020018060200183810383528686828181526020019250808284376000838201819052601f909101601f191690920185810384528651815286516020918201939188019250908190849084905b8381101561243257818101518382015260200161241a565b50505050905090810190601f16801561245f5780820380516001836020036101000a031916815260200191505b509550505050505060405180910390a15050505050505050565b60606012805480602002602001604051908101604052809291908181526020018280548015611876576020028201919060005260206000209081546001600160a01b03168152600190910190602001808311611858575050505050905090565b6124e23361438d565b612521576040805162461bcd60e51b815260206004820152601a60248201526000805160206153b4833981519152604482015290519081900360640190fd5b60145460ff16612570576040805162461bcd60e51b81526020600482015260156024820152743737903832b73234b7339039bab136b4b9b9b4b7b760591b604482015290519081900360640190fd5b6125d160128054806020026020016040519081016040528092919081815260200182805480156113cb576020028201919060005260206000209081546001600160a01b031681526001909101906020018083116113ad575050505050613969565b811461260e5760405162461bcd60e51b81526004018080602001828103825260238152602001806154426023913960400191505060405180910390fd5b60005b6012548110156126f557601060006012838154811061262c57fe5b60009182526020808320909101546001600160a01b0316835282019290925260400190205460ff166126ed576001601060006012848154811061266b57fe5b6000918252602080832091909101546001600160a01b031683528201929092526040019020805460ff19169115159190911790556012805460119190839081106126b157fe5b60009182526020808320909101548354600181018555938352912090910180546001600160a01b0319166001600160a01b039092169190911790555b600101612611565b507fb2f6cccee7a369e23e293c25aa19bef80af11eb26deba3ea0f2a02783f752e4a33601260405180836001600160a01b03166001600160a01b0316815260200180602001828103825283818154815260200191508054801561278157602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311612763575b5050935050505060405180910390a161279c601260006152b2565b506014805460ff19169055565b600061168760156145ba565b6127be33614310565b806127c857503330145b61280c576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b60145460ff161580156128275750601454610100900460ff16155b612878576040805162461bcd60e51b815260206004820152601c60248201527f77686974656c6973742073756d62697373696f6e2070656e64696e6700000000604482015290519081900360640190fd5b60145462010000900460ff166128d1576040805162461bcd60e51b81526020600482015260196024820152781dda1a5d195b1a5cdd081b9bdd081a5b9a5d1a585b1a5e9959603a1b604482015290519081900360640190fd5b80612915576040805162461bcd60e51b815260206004820152600f60248201526e195b5c1d1e481dda1a5d195b1a5cdd608a1b604482015290519081900360640190fd5b612921601383836152d0565b506014805461ff00191661010017905560408051602080840282810182019093528382527ffbc0e5ca6c7e4858daf0fdb185ef5186203e74ec9c64737e93c0aeaec596e1d192859285926129909285918591829185019084908082843760009201919091525061396992505050565b60405180806020018381526020018281038252858582818152602001925060200280828437600083820152604051601f909101601f1916909201829003965090945050505050a15050565b6129e433614310565b612a2e576040805162461bcd60e51b815260206004820152601660248201527539b2b73232b91034b9903737ba1030b71037bbb732b960511b604482015290519081900360640190fd5b600354600160a01b900460ff16612a8c576040805162461bcd60e51b815260206004820152601d60248201527f6f776e657273686970206973206e6f74207472616e7366657261626c65000000604482015290519081900360640190fd5b600380546001600160a01b0319169055604080516000808252602082015281517f850b3df64837d7d518b45f5aa64d104652c3b80eb5b34a8e3d9eb666cb7cdea5929181900390910190a1565b612ae233614310565b80612aec57503330145b612b30576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b8066038d7ea4c6800011158015612b4f57506706f05b59d3b200008111155b612b96576040805162461bcd60e51b815260206004820152601360248201527206f7574206f662072616e676520746f702d757606c1b604482015290519081900360640190fd5b612ba760048263ffffffff6145ef16565b6040805182815290517faf2a77cd04c3cc155588dd3bf67b310ab4fb3b1da3cf6b8d7d4d2aa1d09b794c9181900360200190a150565b601b5490565b6001546001600160a01b031690565b600061168760046145ba565b612c0733614310565b80612c1157503330145b612c55576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b60145460ff16158015612c705750601454610100900460ff16155b612cc1576040805162461bcd60e51b815260206004820152601c60248201527f77686974656c6973742073756d62697373696f6e2070656e64696e6700000000604482015290519081900360640190fd5b8181808060200260200160405190810160405280939291908181526020018383602002808284376000920182905250925050505b8151811015612ddd57612d1a828281518110612d0d57fe5b6020026020010151614310565b15612d65576040805162461bcd60e51b8152602060048201526016602482015275636f6e7461696e73206f776e6572206164647265737360501b604482015290519081900360640190fd5b60006001600160a01b0316828281518110612d7c57fe5b60200260200101516001600160a01b03161415612dd5576040805162461bcd60e51b8152602060048201526012602482015271636f6e7461696e732030206164647265737360701b604482015290519081900360640190fd5b600101612cf5565b5060145462010000900460ff16612e37576040805162461bcd60e51b81526020600482015260196024820152781dda1a5d195b1a5cdd081b9bdd081a5b9a5d1a585b1a5e9959603a1b604482015290519081900360640190fd5b81612e7b576040805162461bcd60e51b815260206004820152600f60248201526e195b5c1d1e481dda1a5d195b1a5cdd608a1b604482015290519081900360640190fd5b612e87601284846152d0565b506014805460ff1916600117905560408051602080850282810182019093528482527f9c80b3b5f68b3e017766d59e8d09b34efe6462b05c398f35cab9e271d9bc3b9c9286928692612ef49285918591829185019084908082843760009201919091525061396992505050565b60405180806020018381526020018281038252858582818152602001925060200280828437600083820152604051601f909101601f1916909201829003965090945050505050a1505050565b60095490565b6003546001600160a01b031690565b612f5e33614310565b612fa8576040805162461bcd60e51b815260206004820152601660248201527539b2b73232b91034b9903737ba1030b71037bbb732b960511b604482015290519081900360640190fd5b612fb0614bf3565b565b60155490565b612fc133614310565b80612fcb57503330145b61300f576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b8051602080820191906000808060605b868510156130eb5761303886605463ffffffff614c3b16565b888601805160148201516034909201805193995060609190911c9650909450909250905061307d6054613071878563ffffffff614c9816565b9063ffffffff614c9816565b9450868511156130c4576040805162461bcd60e51b815260206004820152600d60248201526c6f7574206f6620626f756e647360981b604482015290519081900360640190fd5b816130da57506040805160208101909152600081525b6130e5848483611c71565b5061301f565b5050505050505050565b600f5460ff1690565b61310733614310565b8061311157503330145b613155576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b600a548111156131a7576040805162461bcd60e51b81526020600482015260186024820152771bdd5d081bd9881c985b99d9481b1bd85908185b5bdd5b9d60421b604482015290519081900360640190fd5b6131b8600b8263ffffffff6145ef16565b6040805182815290517fc178d379965e5657b6fc57494e392f121a14119215dfb422aad7db4cc03f2d109181900360200190a150565b6131f733614310565b613241576040805162461bcd60e51b815260206004820152601660248201527539b2b73232b91034b9903737ba1030b71037bbb732b960511b604482015290519081900360640190fd5b600354600160a01b900460ff1661329f576040805162461bcd60e51b815260206004820152601d60248201527f6f776e657273686970206973206e6f74207472616e7366657261626c65000000604482015290519081900360640190fd5b6001600160a01b0382166132e45760405162461bcd60e51b81526004018080602001828103825260238152602001806153f56023913960400191505060405180910390fd5b6003805460ff60a01b1916600160a01b831515021790558061333d57604080516001600160a01b038416815290517f808639ff9c8e4732d60b6c2330de498035416d229f27a77d259680895efec1229181900360200190a15b600354604080516001600160a01b039283168152918416602083015280517f850b3df64837d7d518b45f5aa64d104652c3b80eb5b34a8e3d9eb666cb7cdea59281900390910190a150600380546001600160a01b0319166001600160a01b0392909216919091179055565b6000806000806133b786614cf2565b50505093509350935050801561342c5781613402576040805162461bcd60e51b81526020600482015260066024820152650726174653d360d41b604482015290519081900360640190fd5b61342283613416878563ffffffff614e8416565b9063ffffffff614edd16565b93505050506114ef565b50600095945050505050565b60195460ff1690565b60145462010000900460ff1681565b61345933614310565b8061346357503330145b6134a7576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b80806134e4576040805162461bcd60e51b8152602060048201526007602482015266076616c75653d360cc1b604482015290519081900360640190fd5b6001600160a01b03841661352f576040805162461bcd60e51b815260206004820152600d60248201526c064657374696e6174696f6e3d3609c1b604482015290519081900360640190fd5b6001600160a01b03841660009081526010602052604090205460ff1661357f57816001600160a01b0384161561356c5761356984846133a8565b90505b61357d60158263ffffffff61466a16565b505b61358a848484614f47565b604080516001600160a01b0380871682528516602082015280820184905290517fd1ba4ac2e2a11b5101f6cb4d978f514a155b421e8ec396d2d9abaf0bb02917ee9181900360600190a150505050565b600e5490565b6135e93361438d565b613628576040805162461bcd60e51b815260206004820152601a60248201526000805160206153b4833981519152604482015290519081900360640190fd5b601454610100900460ff1661367c576040805162461bcd60e51b81526020600482015260156024820152743737903832b73234b7339039bab136b4b9b9b4b7b760591b604482015290519081900360640190fd5b6136dd60138054806020026020016040519081016040528092919081815260200182805480156113cb576020028201919060005260206000209081546001600160a01b031681526001909101906020018083116113ad575050505050613969565b811461371a5760405162461bcd60e51b81526004018080602001828103825260238152602001806154426023913960400191505060405180910390fd5b60005b6013548110156138a857601060006013838154811061373857fe5b60009182526020808320909101546001600160a01b0316835282019290925260400190205460ff16156138a0576000601060006013848154811061377857fe5b6000918252602080832091909101546001600160a01b031683528201929092526040018120805460ff1916921515929092179091555b6011546137c290600163ffffffff614c3b16565b81101561388a57601382815481106137d657fe5b600091825260209091200154601180546001600160a01b0390921691839081106137fc57fe5b6000918252602090912001546001600160a01b031614156138825760118054600019810190811061382957fe5b600091825260209091200154601180546001600160a01b03909216918390811061384f57fe5b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555061388a565b6001016137ae565b50601180549061389e906000198301615333565b505b60010161371d565b507fd218c430fa348f4ce67791021b6b89c0c3eacd4ead1d8f5b83c60038ec28249b33601360405180836001600160a01b03166001600160a01b0316815260200180602001828103825283818154815260200191508054801561393457602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311613916575b5050935050505060405180910390a161394f601360006152b2565b506014805461ff0019169055565b60075490565b601a5481565b60008160405160200180828051906020019060200280838360005b8381101561399c578181015183820152602001613984565b50505050905001915050604051602081830303815290604052805190602001209050919050565b6139cc33614310565b806139db57506139db3361438d565b613a25576040805162461bcd60e51b815260206004820152601660248201527537b7363c9037bbb732b93e3e31b7b73a3937b63632b960511b604482015290519081900360640190fd5b601454610100900460ff16613a79576040805162461bcd60e51b81526020600482015260156024820152743737903832b73234b7339039bab136b4b9b9b4b7b760591b604482015290519081900360640190fd5b613ada60138054806020026020016040519081016040528092919081815260200182805480156113cb576020028201919060005260206000209081546001600160a01b031681526001909101906020018083116113ad575050505050613969565b8114613b175760405162461bcd60e51b81526004018080602001828103825260238152602001806154426023913960400191505060405180910390fd5b613b23601360006152b2565b6014805461ff0019169055604080513381526020810183905281517f13c935eb475aa0f6e931fece83e2ac44569ce2d53460d29a6dedab40b965c8a3929181900390910190a150565b60118181548110613b7957fe5b6000918252602090912001546001600160a01b0316905081565b600b5490565b601454610100900460ff1681565b60025490565b8080613bea576040805162461bcd60e51b8152602060048201526007602482015266076616c75653d360cc1b604482015290519081900360640190fd5b613bf333614310565b80613c025750613c023361438d565b613c4c576040805162461bcd60e51b815260206004820152601660248201527537b7363c9037bbb732b93e3e31b7b73a3937b63632b960511b604482015290519081900360640190fd5b613c5d60048363ffffffff61466a16565b613c65612f46565b6001600160a01b03166108fc839081150290604051600060405180830381858888f19350505050158015613c9d573d6000803e3d6000fd5b507f611b7c0d84fda988026215bef9b3e4d81cbceced7e679be6d5e044b588467c0e33613cc8612f46565b604080516001600160a01b03938416815291909216602082015280820185905290519081900360600190a15050565b613d003361438d565b613d3f576040805162461bcd60e51b815260206004820152601a60248201526000805160206153b4833981519152604482015290519081900360640190fd5b611c3360158263ffffffff614fab16565b6000613d5a614fff565b6001600160a01b0316836001600160a01b03161415613d7a5750806114ef565b816001600160a01b03841615613e3f576000806000613d9887614cf2565b5050509350935093505080613dea576040805162461bcd60e51b8152602060048201526013602482015272746f6b656e206e6f7420617661696c61626c6560681b604482015290519081900360640190fd5b81613e25576040805162461bcd60e51b81526020600482015260066024820152650726174653d360d41b604482015290519081900360640190fd5b613e3983613416888563ffffffff614e8416565b93505050505b6000806000613e4c615075565b5050509350935093505080613e9e576040805162461bcd60e51b8152602060048201526013602482015272746f6b656e206e6f7420617661696c61626c6560681b604482015290519081900360640190fd5b81613ee4576040805162461bcd60e51b81526020600482015260116024820152700737461626c65636f696e20726174653d3607c1b604482015290519081900360640190fd5b613ef882613416868663ffffffff614e8416565b979650505050505050565b613f0c3361438d565b613f4b576040805162461bcd60e51b815260206004820152601a60248201526000805160206153b4833981519152604482015290519081900360640190fd5b611b8d600b8263ffffffff614fab16565b613f653361438d565b613fa4576040805162461bcd60e51b815260206004820152601a60248201526000805160206153b4833981519152604482015290519081900360640190fd5b61128360048263ffffffff614fab16565b613fbe33614310565b80613fc857503330145b61400c576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b8181808060200260200160405190810160405280939291908181526020018383602002808284376000920182905250925050505b815181101561411b57614058828281518110612d0d57fe5b156140a3576040805162461bcd60e51b8152602060048201526016602482015275636f6e7461696e73206f776e6572206164647265737360501b604482015290519081900360640190fd5b60006001600160a01b03168282815181106140ba57fe5b60200260200101516001600160a01b03161415614113576040805162461bcd60e51b8152602060048201526012602482015271636f6e7461696e732030206164647265737360701b604482015290519081900360640190fd5b600101614040565b5060145462010000900460ff1615614172576040805162461bcd60e51b81526020600482015260156024820152741dda1a5d195b1a5cdd081a5b9a5d1a585b1a5e9959605a1b604482015290519081900360640190fd5b60005b82811015614263576010600085858481811061418d57fe5b602090810292909201356001600160a01b03168352508101919091526040016000205460ff1661425b576001601060008686858181106141c957fe5b905060200201356001600160a01b03166001600160a01b03166001600160a01b0316815260200190815260200160002060006101000a81548160ff021916908315150217905550601184848381811061421e57fe5b835460018101855560009485526020948590200180546001600160a01b0319166001600160a01b0395909202939093013593909316929092179055505b600101614175565b506014805462ff0000191662010000179055604080513380825260208201838152601180549484018590527fb2f6cccee7a369e23e293c25aa19bef80af11eb26deba3ea0f2a02783f752e4a94929390929091906060830190849080156142f357602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116142d5575b5050935050505060405180910390a1505050565b60085460ff1690565b6003546001600160a01b0390811691161490565b600482015460ff1615614372576040805162461bcd60e51b81526020600482015260116024820152701b1a5b5a5d08185b1c9958591e481cd95d607a1b604482015290519081900360640190fd5b61437c82826151e4565b50600401805460ff19166001179055565b600061439a6002546146e0565b6001600160a01b031663b429afeb836040518263ffffffff1660e01b815260040180826001600160a01b03166001600160a01b0316815260200191505060206040518083038186803b1580156143ef57600080fd5b505afa158015614403573d6000803e3d6000fd5b505050506040513d602081101561441957600080fd5b505192915050565b60008151604114614434575060006114ef565b60208201516040830151606084015160001a7f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a082111561447a57600093505050506114ef565b8060ff16601b1415801561449257508060ff16601c14155b156144a357600093505050506114ef565b6040805160008152602080820180845289905260ff8416828401526060820186905260808201859052915160019260a0808401939192601f1981019281900390910190855afa1580156144fa573d6000803e3d6000fd5b5050604051601f190151979650505050505050565b60006001600160a01b038216156145a957816001600160a01b03166370a08231846040518263ffffffff1660e01b815260040180826001600160a01b03166001600160a01b0316815260200191505060206040518083038186803b15801561457657600080fd5b505afa15801561458a573d6000803e3d6000fd5b505050506040513d60208110156145a057600080fd5b505190506114ef565b506001600160a01b038216316114ef565b60028101546000906145d5906201518063ffffffff614c9816565b4211156145e4575080546111a9565b5060018101546111a9565b600482015460ff16614648576040805162461bcd60e51b815260206004820152601960248201527f6c696d6974206861736e2774206265656e207365742079657400000000000000604482015290519081900360640190fd5b600390910155565b60008061465c83614cf2565b509098975050505050505050565b61467382615207565b80826001015410156146bf576040805162461bcd60e51b815260206004820152601060248201526f185d985a5b18589b194f185b5bdd5b9d60821b604482015290519081900360640190fd5b60018201546146d4908263ffffffff614c3b16565b82600101819055505050565b6000805460408051630178b8bf60e01b81526004810185905290516001600160a01b0390921691630178b8bf91602480820192602092909190829003018186803b15801561472d57600080fd5b505afa158015614741573d6000803e3d6000fd5b505050506040513d602081101561475757600080fd5b505160408051631d9dabef60e11b81526004810185905290516001600160a01b0390921691633b3b57de91602480820192602092909190829003018186803b1580156143ef57600080fd5b801580614828575060408051636eb1769f60e11b81523060048201526001600160a01b03848116602483015291519185169163dd62ed3e91604480820192602092909190829003018186803b1580156147fa57600080fd5b505afa15801561480e573d6000803e3d6000fd5b505050506040513d602081101561482457600080fd5b5051155b6148635760405162461bcd60e51b81526004018080602001828103825260368152602001806154656036913960400191505060405180910390fd5b604080516001600160a01b038416602482015260448082018490528251808303909101815260649091019091526020810180516001600160e01b031663095ea7b360e01b1790526148b59084906149e4565b505050565b3b151590565b6000806148cc83614cf2565b509198975050505050505050565b6000806148e86009546146e0565b6001600160a01b031663afc72e9385856040518363ffffffff1660e01b815260040180836001600160a01b03166001600160a01b0316815260200180602001828103825283818151815260200191508051906020019080838360005b8381101561495c578181015183820152602001614944565b50505050905090810190601f1680156149895780820380516001836020036101000a031916815260200191505b509350505050604080518083038186803b1580156149a657600080fd5b505afa1580156149ba573d6000803e3d6000fd5b505050506040513d60408110156149d057600080fd5b508051602090910151909590945092505050565b6149f6826001600160a01b03166148ba565b614a47576040805162461bcd60e51b815260206004820152601f60248201527f5361666545524332303a2063616c6c20746f206e6f6e2d636f6e747261637400604482015290519081900360640190fd5b60006060836001600160a01b0316836040518082805190602001908083835b60208310614a855780518252601f199092019160209182019101614a66565b6001836020036101000a0380198251168184511680821785525050505050509050019150506000604051808303816000865af19150503d8060008114614ae7576040519150601f19603f3d011682016040523d82523d6000602084013e614aec565b606091505b509150915081614b43576040805162461bcd60e51b815260206004820181905260248201527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564604482015290519081900360640190fd5b805115614b9c57808060200190516020811015614b5f57600080fd5b5051614b9c5760405162461bcd60e51b815260040180806020018281038252602a815260200180615418602a913960400191505060405180910390fd5b50505050565b604080517f19457468657265756d205369676e6564204d6573736167653a0a333200000000602080830191909152603c8083019490945282518083039094018452605c909101909152815191012090565b601a80546001019081905560408051338152602081019290925280517fab0423a75986556234aecd171c46ce7f5e45607d8070bf5230f2735b50322bff9281900390910190a1565b600082821115614c92576040805162461bcd60e51b815260206004820152601e60248201527f536166654d6174683a207375627472616374696f6e206f766572666c6f770000604482015290519081900360640190fd5b50900390565b600082820183811015612161576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b6060600080600080600080614d086009546146e0565b6001600160a01b0316631f69565f896040518263ffffffff1660e01b815260040180826001600160a01b03166001600160a01b0316815260200191505060006040518083038186803b158015614d5d57600080fd5b505afa158015614d71573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405260e0811015614d9a57600080fd5b8101908080516040519392919084600160201b821115614db957600080fd5b908301906020820185811115614dce57600080fd5b8251600160201b811182820188101715614de757600080fd5b82525081516020918201929091019080838360005b83811015614e14578181015183820152602001614dfc565b50505050905090810190601f168015614e415780820380516001836020036101000a031916815260200191505b5060409081526020820151908201516060830151608084015160a085015160c090950151979e50929c50909a509850965094509192505050919395979092949650565b600082614e93575060006114ef565b82820282848281614ea057fe5b04146121615760405162461bcd60e51b81526004018080602001828103825260218152602001806153d46021913960400191505060405180910390fd5b6000808211614f33576040805162461bcd60e51b815260206004820152601a60248201527f536166654d6174683a206469766973696f6e206279207a65726f000000000000604482015290519081900360640190fd5b6000828481614f3e57fe5b04949350505050565b6001600160a01b038216614f91576040516001600160a01b0384169082156108fc029083906000818181858888f19350505050158015614f8b573d6000803e3d6000fd5b506148b5565b6148b56001600160a01b038316848363ffffffff61526016565b80826003015414614fed5760405162461bcd60e51b81526004018080602001828103825260228152602001806153926022913960400191505060405180910390fd5b614ffb8283600301546151e4565b5050565b600061500c6009546146e0565b6001600160a01b031663e9cbd8226040518163ffffffff1660e01b815260040160206040518083038186803b15801561504457600080fd5b505afa158015615058573d6000803e3d6000fd5b505050506040513d602081101561506e57600080fd5b5051905090565b606060008060008060008061508b6009546146e0565b6001600160a01b0316633efec5e96040518163ffffffff1660e01b815260040160006040518083038186803b1580156150c357600080fd5b505afa1580156150d7573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405260e081101561510057600080fd5b8101908080516040519392919084600160201b82111561511f57600080fd5b90830190602082018581111561513457600080fd5b8251600160201b81118282018810171561514d57600080fd5b82525081516020918201929091019080838360005b8381101561517a578181015183820152602001615162565b50505050905090810190601f1680156151a75780820380516001836020036101000a031916815260200191505b5060409081526020820151908201516060830151608084015160a085015160c090950151979f939e50919c509a5098509096509294509192505050565b6151ed82615207565b8082556001820154811015614ffb57815460018301555050565b600281015461521f906201518063ffffffff614c9816565b42111561525d57426002820155805460018201556040517fe93bc25276d408d390778e7a8b926f2f67209c43ed540081b951fe128f0d3cd290600090a15b50565b604080516001600160a01b038416602482015260448082018490528251808303909101815260649091019091526020810180516001600160e01b031663a9059cbb60e01b1790526148b59084906149e4565b508054600082559060005260206000209081019061525d9190615353565b828054828255906000526020600020908101928215615323579160200282015b828111156153235781546001600160a01b0319166001600160a01b038435161782556020909201916001909101906152f0565b5061532f92915061536d565b5090565b8154818355818111156148b5576000838152602090206148b59181019083015b6111b291905b8082111561532f5760008155600101615359565b6111b291905b8082111561532f5780546001600160a01b031916815560010161537356fe636f6e6669726d65642f7375626d6974746564206c696d6974206d69736d6174636873656e646572206973206e6f74206120636f6e74726f6c6c6572000000000000536166654d6174683a206d756c7469706c69636174696f6e206f766572666c6f776f776e65722063616e6e6f742062652073657420746f207a65726f20616464726573735361666545524332303a204552433230206f7065726174696f6e20646964206e6f7420737563636565646e6f6e2d6d61746368696e672070656e64696e672077686974656c69737420686173685361666545524332303a20617070726f76652066726f6d206e6f6e2d7a65726f20746f206e6f6e2d7a65726f20616c6c6f77616e6365a265627a7a7231582026e1c268487cb71f6386771c751057a76de75c808456898744d851d2655c5dcc64736f6c634300050f0032"
=======
<<<<<<< HEAD
var WalletBin = "0x60806040527f7f2ce995617d2816b426c5c8698c5ec2952f7a34bb10f38326f74933d58936976002553480156200003557600080fd5b5060405162005a4738038062005a47833981810160405260e08110156200005b57600080fd5b508051602082015160408301516060840151608085015160a086015160c090960151600180546001600160a01b038087166001600160a01b0319928316179283905560008054909216921691909117905594959394929391929091908084808989878015620000ca5760028190555b50600380546001600160a01b0319166001600160a01b0384161760ff60a01b1916600160a01b8315158102919091179182905560ff9104166200014457604080516001600160a01b038416815290517f808639ff9c8e4732d60b6c2330de498035416d229f27a77d259680895efec1229181900360200190a15b60408051600081526001600160a01b038416602082015281517f850b3df64837d7d518b45f5aa64d104652c3b80eb5b34a8e3d9eb666cb7cdea5929181900390910190a150506040805160a0810182526706f05b59d3b2000080825260208201819052429282018390526000606083018190526080909201829052600481905560055560069190915560078190556008805460ff19169055600991909155620001f56001600160e01b03620002ec16565b50505050509150506000811162000243576040805162461bcd60e51b815260206004820152600d60248201526c37379039ba30b13632b1b7b4b760991b604482015290519081900360640190fd5b61271002600a8190556040805160a08082018352838252602080830185905242838501819052600060608086018290526080958601829052600b889055600c97909755600d829055600e819055600f805460ff1990811690915586519485018752898552928401899052948301819052948201849052910182905260158590556016949094556017919091556018555060198054909116905550601b5550620005689350505050565b60606000806000806000806200030a6009546200046e60201b60201c565b6001600160a01b0316633efec5e96040518163ffffffff1660e01b815260040160006040518083038186803b1580156200034357600080fd5b505afa15801562000358573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405260e08110156200038257600080fd5b8101908080516040519392919084640100000000821115620003a357600080fd5b908301906020820185811115620003b957600080fd5b8251640100000000811182820188101715620003d457600080fd5b82525081516020918201929091019080838360005b8381101562000403578181015183820152602001620003e9565b50505050905090810190601f168015620004315780820380516001836020036101000a031916815260200191505b5060409081526020820151908201516060830151608084015160a085015160c090950151979f939e50919c509a5098509096509294509192505050565b6000805460408051630178b8bf60e01b81526004810185905290516001600160a01b0390921691630178b8bf91602480820192602092909190829003018186803b158015620004bc57600080fd5b505afa158015620004d1573d6000803e3d6000fd5b505050506040513d6020811015620004e857600080fd5b505160408051631d9dabef60e11b81526004810185905290516001600160a01b0390921691633b3b57de91602480820192602092909190829003018186803b1580156200053457600080fd5b505afa15801562000549573d6000803e3d6000fd5b505050506040513d60208110156200056057600080fd5b505192915050565b6154cf80620005786000396000f3fe60806040526004361061038c5760003560e01c80637fd004fa116101dc578063cc0e7e5611610102578063e2b4ce97116100a0578063f40b51f81161006f578063f40b51f8146110b0578063f41c4319146110da578063f421764814611104578063f776f5181461117f5761038c565b8063e2b4ce971461100e578063e61c51ca14611023578063eadd3cea1461104d578063f36febda146110775761038c565b8063ce0b5bd5116100dc578063ce0b5bd514610f90578063d251fefc14610fba578063da84b1ed14610fe4578063de212bf314610ff95761038c565b8063cc0e7e5614610eb8578063cccdc55614610ecd578063cd7958dd14610ee25761038c565b8063b221f3161161017a578063be40ba7911610149578063be40ba7914610e21578063beabacc814610e36578063c4856cd914610e79578063cbd2ac6814610e8e5761038c565b8063b221f31614610d6e578063b242e53414610d98578063b87e21ef14610dd3578063bcb8b74a14610e0c5761038c565b806390e690c7116101b657806390e690c714610c7e5780639b0dfd2714610c93578063aaf1fc6214610ca8578063ab20599314610d595761038c565b80637fd004fa14610bd9578063877337b014610c545780638da5cb5b14610c695761038c565b806332531c3c116102c15780635adc02ab1161025f57806374624c551161022e57806374624c5514610b54578063747c31d614610b7e5780637d73b23114610b935780637d7d004614610bc45761038c565b80635adc02ab14610a855780635d2362a814610aaf5780636137d67014610ac4578063715018a614610b3f5761038c565b80633c672eb71161029b5780633c672eb7146108ae5780633f579f42146108d857806346efe0ed1461099e57806347b55a9d14610a705761038c565b806332531c3c146108255780633a43199f146108585780633bfec254146108845761038c565b80631efd02991161032e57806321ce918d1161030857806321ce918d1461076c5780632587a6a21461079657806326d05ab2146107ab578063294f4025146107c05761038c565b80631efd02991461067757806320c13b0b1461068c5780632121dc75146107575761038c565b8063100f23fd1161036a578063100f23fd146104635780631127b57e1461048d5780631626ba7e146105175780631aa21fba146105ec5761038c565b806301ffc9a7146103c8578063027ef3eb146104105780630f3a85d814610437575b6040805133815234602082015281517f88a5966d370b9919b20f3e2c13ff65706f196a4e32cc2c12bf57088f88525874929181900390910190a1005b3480156103d457600080fd5b506103fc600480360360208110156103eb57600080fd5b50356001600160e01b031916611194565b604080519115158252519081900360200190f35b34801561041c57600080fd5b506104256111ae565b60408051918252519081900360200190f35b34801561044357600080fd5b506104616004803603602081101561045a57600080fd5b50356111b5565b005b34801561046f57600080fd5b506104616004803603602081101561048657600080fd5b50356112c1565b34801561049957600080fd5b506104a2611466565b6040805160208082528351818301528351919283929083019185019080838360005b838110156104dc5781810151838201526020016104c4565b50505050905090810190601f1680156105095780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561052357600080fd5b506105cf6004803603604081101561053a57600080fd5b81359190810190604081016020820135600160201b81111561055b57600080fd5b82018360208201111561056d57600080fd5b803590602001918460018302840111600160201b8311171561058e57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550611487945050505050565b604080516001600160e01b03199092168252519081900360200190f35b3480156105f857600080fd5b506104616004803603604081101561060f57600080fd5b6001600160a01b038235169190810190604081016020820135600160201b81111561063957600080fd5b82018360208201111561064b57600080fd5b803590602001918460208302840111600160201b8311171561066c57600080fd5b5090925090506114f5565b34801561068357600080fd5b5061042561167b565b34801561069857600080fd5b506105cf600480360360408110156106af57600080fd5b810190602081018135600160201b8111156106c957600080fd5b8201836020820111156106db57600080fd5b803590602001918460018302840111600160201b831117156106fc57600080fd5b919390929091602081019035600160201b81111561071957600080fd5b82018360208201111561072b57600080fd5b803590602001918460018302840111600160201b8311171561074c57600080fd5b50909250905061168c565b34801561076357600080fd5b506103fc611761565b34801561077857600080fd5b506104616004803603602081101561078f57600080fd5b5035611771565b3480156107a257600080fd5b5061042561180f565b3480156107b757600080fd5b506103fc611815565b3480156107cc57600080fd5b506107d561181e565b60408051602080825283518183015283519192839290830191858101910280838360005b838110156108115781810151838201526020016107f9565b505050509050019250505060405180910390f35b34801561083157600080fd5b506103fc6004803603602081101561084857600080fd5b50356001600160a01b0316611880565b6104616004803603604081101561086e57600080fd5b506001600160a01b038135169060200135611895565b34801561089057600080fd5b50610461600480360360208110156108a757600080fd5b5035611ad3565b3480156108ba57600080fd5b50610461600480360360208110156108d157600080fd5b5035611bcb565b3480156108e457600080fd5b506104a2600480360360608110156108fb57600080fd5b6001600160a01b0382351691602081013591810190606081016040820135600160201b81111561092a57600080fd5b82018360208201111561093c57600080fd5b803590602001918460018302840111600160201b8311171561095d57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550611c71945050505050565b3480156109aa57600080fd5b50610461600480360360608110156109c157600080fd5b81359190810190604081016020820135600160201b8111156109e257600080fd5b8201836020820111156109f457600080fd5b803590602001918460018302840111600160201b83111715610a1557600080fd5b919390929091602081019035600160201b811115610a3257600080fd5b820183602082011115610a4457600080fd5b803590602001918460018302840111600160201b83111715610a6557600080fd5b509092509050612168565b348015610a7c57600080fd5b506107d5612479565b348015610a9157600080fd5b5061046160048036036020811015610aa857600080fd5b50356124d9565b348015610abb57600080fd5b506104256127a9565b348015610ad057600080fd5b5061046160048036036020811015610ae757600080fd5b810190602081018135600160201b811115610b0157600080fd5b820183602082011115610b1357600080fd5b803590602001918460208302840111600160201b83111715610b3457600080fd5b5090925090506127b5565b348015610b4b57600080fd5b506104616129db565b348015610b6057600080fd5b5061046160048036036020811015610b7757600080fd5b5035612ad9565b348015610b8a57600080fd5b50610425612bdd565b348015610b9f57600080fd5b50610ba8612be3565b604080516001600160a01b039092168252519081900360200190f35b348015610bd057600080fd5b50610425612bf2565b348015610be557600080fd5b5061046160048036036020811015610bfc57600080fd5b810190602081018135600160201b811115610c1657600080fd5b820183602082011115610c2857600080fd5b803590602001918460208302840111600160201b83111715610c4957600080fd5b509092509050612bfe565b348015610c6057600080fd5b50610425612f40565b348015610c7557600080fd5b50610ba8612f46565b348015610c8a57600080fd5b50610461612f55565b348015610c9f57600080fd5b50610425612fb2565b348015610cb457600080fd5b5061046160048036036020811015610ccb57600080fd5b810190602081018135600160201b811115610ce557600080fd5b820183602082011115610cf757600080fd5b803590602001918460018302840111600160201b83111715610d1857600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550612fb8945050505050565b348015610d6557600080fd5b506103fc6130f5565b348015610d7a57600080fd5b5061046160048036036020811015610d9157600080fd5b50356130fe565b348015610da457600080fd5b5061046160048036036040811015610dbb57600080fd5b506001600160a01b03813516906020013515156131ee565b348015610ddf57600080fd5b5061042560048036036040811015610df657600080fd5b506001600160a01b0381351690602001356133a8565b348015610e1857600080fd5b506103fc613438565b348015610e2d57600080fd5b506103fc613441565b348015610e4257600080fd5b5061046160048036036060811015610e5957600080fd5b506001600160a01b03813581169160208101359091169060400135613450565b348015610e8557600080fd5b506104256135da565b348015610e9a57600080fd5b5061046160048036036020811015610eb157600080fd5b50356135e0565b348015610ec457600080fd5b5061042561395d565b348015610ed957600080fd5b50610425613963565b348015610eee57600080fd5b5061042560048036036020811015610f0557600080fd5b810190602081018135600160201b811115610f1f57600080fd5b820183602082011115610f3157600080fd5b803590602001918460208302840111600160201b83111715610f5257600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929550613969945050505050565b348015610f9c57600080fd5b5061046160048036036020811015610fb357600080fd5b50356139c3565b348015610fc657600080fd5b50610ba860048036036020811015610fdd57600080fd5b5035613b6c565b348015610ff057600080fd5b50610425613b93565b34801561100557600080fd5b506103fc613b99565b34801561101a57600080fd5b50610425613ba7565b34801561102f57600080fd5b506104616004803603602081101561104657600080fd5b5035613bad565b34801561105957600080fd5b506104616004803603602081101561107057600080fd5b5035613cf7565b34801561108357600080fd5b506104256004803603604081101561109a57600080fd5b506001600160a01b038135169060200135613d50565b3480156110bc57600080fd5b50610461600480360360208110156110d357600080fd5b5035613f03565b3480156110e657600080fd5b50610461600480360360208110156110fd57600080fd5b5035613f5c565b34801561111057600080fd5b506104616004803603602081101561112757600080fd5b810190602081018135600160201b81111561114157600080fd5b82018360208201111561115357600080fd5b803590602001918460208302840111600160201b8311171561117457600080fd5b509092509050613fb5565b34801561118b57600080fd5b506103fc614307565b6001600160e01b031981166301ffc9a760e01b145b919050565b6018545b90565b6111be33614310565b806111c857503330145b61120c576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b8066038d7ea4c680001115801561122b57506706f05b59d3b200008111155b611272576040805162461bcd60e51b815260206004820152601360248201527206f7574206f662072616e676520746f702d757606c1b604482015290519081900360640190fd5b61128360048263ffffffff61432416565b604080513381526020810183905281517f41ff5d5ce3b7935893a4e7269ec5caae9cca5e3bf0eb4b21d2f443489667112e929181900390910190a150565b6112ca33614310565b806112d957506112d93361438d565b611323576040805162461bcd60e51b815260206004820152601660248201527537b7363c9037bbb732b93e3e31b7b73a3937b63632b960511b604482015290519081900360640190fd5b60145460ff16611372576040805162461bcd60e51b81526020600482015260156024820152743737903832b73234b7339039bab136b4b9b9b4b7b760591b604482015290519081900360640190fd5b6113d560128054806020026020016040519081016040528092919081815260200182805480156113cb57602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116113ad575b5050505050613969565b81146114125760405162461bcd60e51b81526004018080602001828103825260238152602001806154426023913960400191505060405180910390fd5b61141e601260006152b2565b6014805460ff19169055604080513381526020810183905281517f7794eff834d760583543e6e510e717a5e66d2c064e225f4db448343c3e66afcf929181900390910190a150565b60405180604001604052806005815260200164332e312e3160d81b81525081565b60008061149a848463ffffffff61442116565b90506114a581614310565b6114e3576040805162461bcd60e51b815260206004820152600a60248201526937b7363c9037bbb732b960b11b604482015290519081900360640190fd5b50630b135d3f60e11b90505b92915050565b6114fe33614310565b8061150857503330145b61154c576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b80611595576040805162461bcd60e51b8152602060048201526014602482015273617373657420617272617920697320656d70747960601b604482015290519081900360640190fd5b60005b818110156115f85760006115c7308585858181106115b257fe5b905060200201356001600160a01b031661450f565b90506115ef858585858181106115d957fe5b905060200201356001600160a01b031683613450565b50600101611598565b507fd4f62f23021706247dcffea245d104ae7ddaec7f23acf3d11d7136d5de6a69ad83838360405180846001600160a01b03166001600160a01b03168152602001806020018281038252848482818152602001925060200280828437600083820152604051601f909101601f1916909201829003965090945050505050a1505050565b6000611687600b6145ba565b905090565b6000808585604051602001808383808284376040805191909301818103601f190182528084528151602092830120601f8b01839004830282018301909452898152929650630b135d3f60e11b9550611703945086935089915088908190840183828082843760009201919091525061148792505050565b6001600160e01b0319161461174f576040805162461bcd60e51b815260206004820152600d60248201526c1cda59c81b9bdd081d985b1a59609a1b604482015290519081900360640190fd5b506320c13b0b60e01b95945050505050565b600354600160a01b900460ff1690565b61177a33614310565b8061178457503330145b6117c8576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b6117d960158263ffffffff6145ef16565b6040805182815290517f4b1b970c8a0fa761e7803ed70c13d7aca71904b13df60fbe03f981da1730da919181900360200190a150565b60045490565b60145460ff1681565b6060601380548060200260200160405190810160405280929190818152602001828054801561187657602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311611858575b5050505050905090565b60106020526000908152604090205460ff1681565b61189e33614310565b806118a857503330145b6118ec576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b6118f582614650565b61193b576040805162461bcd60e51b8152602060048201526012602482015271746f6b656e206e6f74206c6f616461626c6560701b604482015290519081900360640190fd5b60006119478383613d50565b905061195a600b8263ffffffff61466a16565b6000611967601b546146e0565b90506001600160a01b03841615611a0f576119926001600160a01b038516828563ffffffff6147a216565b806001600160a01b0316631b3c96b485856040518363ffffffff1660e01b815260040180836001600160a01b03166001600160a01b0316815260200182815260200192505050600060405180830381600087803b1580156119f257600080fd5b505af1158015611a06573d6000803e3d6000fd5b50505050611a89565b806001600160a01b0316631b3c96b48486866040518463ffffffff1660e01b815260040180836001600160a01b03166001600160a01b03168152602001828152602001925050506000604051808303818588803b158015611a6f57600080fd5b505af1158015611a83573d6000803e3d6000fd5b50505050505b604080516001600160a01b03861681526020810185905281517f5f65674bec9af81f71be68674135a0ea3f163fb91984e3893d06da9f6ea2ce8a929181900390910190a150505050565b611adc33614310565b80611ae657503330145b611b2a576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b600a54811115611b7c576040805162461bcd60e51b81526020600482015260186024820152771bdd5d081bd9881c985b99d9481b1bd85908185b5bdd5b9d60421b604482015290519081900360640190fd5b611b8d600b8263ffffffff61432416565b604080513381526020810183905281517f0b05243483e17c3f3377aee82b7d47e5700b48288695fc08b7ecc2759afa44ef929181900390910190a150565b611bd433614310565b80611bde57503330145b611c22576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b611c3360158263ffffffff61432416565b604080513381526020810183905281517f068f112e5ec923d412be64779fe69e0fcbb6784c6617e94cccc8fd348f2e0f21929181900390910190a150565b6060611c7c33614310565b80611c8657503330145b611cca576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b6001600160a01b03841660009081526010602052604090205460ff16611cfb57611cfb60158463ffffffff61466a16565b611d0d846001600160a01b03166148ba565b8015611d1d5750611d1d846148c0565b15611f0457600080611d2f86856148da565b6001600160a01b038216600090815260106020526040902054919350915060ff16611d75576000611d6087836133a8565b9050611d7360158263ffffffff61466a16565b505b611d8e6001600160a01b0387168563ffffffff6149e416565b604080516020808252818301909252606091602082018180388339019050509050600160f81b81601f81518110611dc157fe5b60200101906001600160f81b031916908160001a9053507ff77753fab406ecfff96d6ff2476c64a838fa9f6d37b1bf190f8546e395e3b6138787878460405180856001600160a01b03166001600160a01b031681526020018481526020018060200180602001838103835285818151815260200191508051906020019080838360005b83811015611e5c578181015183820152602001611e44565b50505050905090810190601f168015611e895780820380516001836020036101000a031916815260200191505b50838103825284518152845160209182019186019080838360005b83811015611ebc578181015183820152602001611ea4565b50505050905090810190601f168015611ee95780820380516001836020036101000a031916815260200191505b50965050505050505060405180910390a19250612161915050565b60006060856001600160a01b031685856040518082805190602001908083835b60208310611f435780518252601f199092019160209182019101611f24565b6001836020036101000a03801982511681845116808217855250505050505090500191505060006040518083038185875af1925050503d8060008114611fa5576040519150601f19603f3d011682016040523d82523d6000602084013e611faa565b606091505b509150915081819061203a5760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b83811015611fff578181015183820152602001611fe7565b50505050905090810190601f16801561202c5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b507ff77753fab406ecfff96d6ff2476c64a838fa9f6d37b1bf190f8546e395e3b6138686868460405180856001600160a01b03166001600160a01b031681526020018481526020018060200180602001838103835285818151815260200191508051906020019080838360005b838110156120bf5781810151838201526020016120a7565b50505050905090810190601f1680156120ec5780820380516001836020036101000a031916815260200191505b50838103825284518152845160209182019186019080838360005b8381101561211f578181015183820152602001612107565b50505050905090810190601f16801561214c5780820380516001836020036101000a031916815260200191505b50965050505050505060405180910390a19150505b9392505050565b6121713361438d565b6121b0576040805162461bcd60e51b815260206004820152601a60248201526000805160206153b4833981519152604482015290519081900360640190fd5b600061220186868660405160200180806339363c1d60e11b81525060040184815260200183838082843780830192505050935050505060405160208183030381529060405280519060200120614ba2565b9050631626ba7e60e01b6001600160e01b0319166122558285858080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061148792505050565b6001600160e01b031916146122a1576040805162461bcd60e51b815260206004820152600d60248201526c1cda59c81b9bdd081d985b1a59609a1b604482015290519081900360640190fd5b601a5486146122e3576040805162461bcd60e51b81526020600482015260096024820152687478207265706c617960b81b604482015290519081900360640190fd5b6122eb614bf3565b60006060306001600160a01b03168787604051808383808284376040519201945060009350909150508083038183865af19150503d806000811461234b576040519150601f19603f3d011682016040523d82523d6000602084013e612350565b606091505b50915091508181906123a35760405162461bcd60e51b8152602060048201818152835160248401528351909283926044909101919085019080838360008315611fff578181015183820152602001611fe7565b507f823dbcf2b7b0f265871963ca65ac033f6b4c71e0d82cd123d2ff23d752dc21c187878360405180806020018060200183810383528686828181526020019250808284376000838201819052601f909101601f191690920185810384528651815286516020918201939188019250908190849084905b8381101561243257818101518382015260200161241a565b50505050905090810190601f16801561245f5780820380516001836020036101000a031916815260200191505b509550505050505060405180910390a15050505050505050565b60606012805480602002602001604051908101604052809291908181526020018280548015611876576020028201919060005260206000209081546001600160a01b03168152600190910190602001808311611858575050505050905090565b6124e23361438d565b612521576040805162461bcd60e51b815260206004820152601a60248201526000805160206153b4833981519152604482015290519081900360640190fd5b60145460ff16612570576040805162461bcd60e51b81526020600482015260156024820152743737903832b73234b7339039bab136b4b9b9b4b7b760591b604482015290519081900360640190fd5b6125d160128054806020026020016040519081016040528092919081815260200182805480156113cb576020028201919060005260206000209081546001600160a01b031681526001909101906020018083116113ad575050505050613969565b811461260e5760405162461bcd60e51b81526004018080602001828103825260238152602001806154426023913960400191505060405180910390fd5b60005b6012548110156126f557601060006012838154811061262c57fe5b60009182526020808320909101546001600160a01b0316835282019290925260400190205460ff166126ed576001601060006012848154811061266b57fe5b6000918252602080832091909101546001600160a01b031683528201929092526040019020805460ff19169115159190911790556012805460119190839081106126b157fe5b60009182526020808320909101548354600181018555938352912090910180546001600160a01b0319166001600160a01b039092169190911790555b600101612611565b507fb2f6cccee7a369e23e293c25aa19bef80af11eb26deba3ea0f2a02783f752e4a33601260405180836001600160a01b03166001600160a01b0316815260200180602001828103825283818154815260200191508054801561278157602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311612763575b5050935050505060405180910390a161279c601260006152b2565b506014805460ff19169055565b600061168760156145ba565b6127be33614310565b806127c857503330145b61280c576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b60145460ff161580156128275750601454610100900460ff16155b612878576040805162461bcd60e51b815260206004820152601c60248201527f77686974656c6973742073756d62697373696f6e2070656e64696e6700000000604482015290519081900360640190fd5b60145462010000900460ff166128d1576040805162461bcd60e51b81526020600482015260196024820152781dda1a5d195b1a5cdd081b9bdd081a5b9a5d1a585b1a5e9959603a1b604482015290519081900360640190fd5b80612915576040805162461bcd60e51b815260206004820152600f60248201526e195b5c1d1e481dda1a5d195b1a5cdd608a1b604482015290519081900360640190fd5b612921601383836152d0565b506014805461ff00191661010017905560408051602080840282810182019093528382527ffbc0e5ca6c7e4858daf0fdb185ef5186203e74ec9c64737e93c0aeaec596e1d192859285926129909285918591829185019084908082843760009201919091525061396992505050565b60405180806020018381526020018281038252858582818152602001925060200280828437600083820152604051601f909101601f1916909201829003965090945050505050a15050565b6129e433614310565b612a2e576040805162461bcd60e51b815260206004820152601660248201527539b2b73232b91034b9903737ba1030b71037bbb732b960511b604482015290519081900360640190fd5b600354600160a01b900460ff16612a8c576040805162461bcd60e51b815260206004820152601d60248201527f6f776e657273686970206973206e6f74207472616e7366657261626c65000000604482015290519081900360640190fd5b600380546001600160a01b0319169055604080516000808252602082015281517f850b3df64837d7d518b45f5aa64d104652c3b80eb5b34a8e3d9eb666cb7cdea5929181900390910190a1565b612ae233614310565b80612aec57503330145b612b30576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b8066038d7ea4c6800011158015612b4f57506706f05b59d3b200008111155b612b96576040805162461bcd60e51b815260206004820152601360248201527206f7574206f662072616e676520746f702d757606c1b604482015290519081900360640190fd5b612ba760048263ffffffff6145ef16565b6040805182815290517faf2a77cd04c3cc155588dd3bf67b310ab4fb3b1da3cf6b8d7d4d2aa1d09b794c9181900360200190a150565b601b5490565b6001546001600160a01b031690565b600061168760046145ba565b612c0733614310565b80612c1157503330145b612c55576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b60145460ff16158015612c705750601454610100900460ff16155b612cc1576040805162461bcd60e51b815260206004820152601c60248201527f77686974656c6973742073756d62697373696f6e2070656e64696e6700000000604482015290519081900360640190fd5b8181808060200260200160405190810160405280939291908181526020018383602002808284376000920182905250925050505b8151811015612ddd57612d1a828281518110612d0d57fe5b6020026020010151614310565b15612d65576040805162461bcd60e51b8152602060048201526016602482015275636f6e7461696e73206f776e6572206164647265737360501b604482015290519081900360640190fd5b60006001600160a01b0316828281518110612d7c57fe5b60200260200101516001600160a01b03161415612dd5576040805162461bcd60e51b8152602060048201526012602482015271636f6e7461696e732030206164647265737360701b604482015290519081900360640190fd5b600101612cf5565b5060145462010000900460ff16612e37576040805162461bcd60e51b81526020600482015260196024820152781dda1a5d195b1a5cdd081b9bdd081a5b9a5d1a585b1a5e9959603a1b604482015290519081900360640190fd5b81612e7b576040805162461bcd60e51b815260206004820152600f60248201526e195b5c1d1e481dda1a5d195b1a5cdd608a1b604482015290519081900360640190fd5b612e87601284846152d0565b506014805460ff1916600117905560408051602080850282810182019093528482527f9c80b3b5f68b3e017766d59e8d09b34efe6462b05c398f35cab9e271d9bc3b9c9286928692612ef49285918591829185019084908082843760009201919091525061396992505050565b60405180806020018381526020018281038252858582818152602001925060200280828437600083820152604051601f909101601f1916909201829003965090945050505050a1505050565b60095490565b6003546001600160a01b031690565b612f5e33614310565b612fa8576040805162461bcd60e51b815260206004820152601660248201527539b2b73232b91034b9903737ba1030b71037bbb732b960511b604482015290519081900360640190fd5b612fb0614bf3565b565b60155490565b612fc133614310565b80612fcb57503330145b61300f576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b8051602080820191906000808060605b868510156130eb5761303886605463ffffffff614c3b16565b888601805160148201516034909201805193995060609190911c9650909450909250905061307d6054613071878563ffffffff614c9816565b9063ffffffff614c9816565b9450868511156130c4576040805162461bcd60e51b815260206004820152600d60248201526c6f7574206f6620626f756e647360981b604482015290519081900360640190fd5b816130da57506040805160208101909152600081525b6130e5848483611c71565b5061301f565b5050505050505050565b600f5460ff1690565b61310733614310565b8061311157503330145b613155576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b600a548111156131a7576040805162461bcd60e51b81526020600482015260186024820152771bdd5d081bd9881c985b99d9481b1bd85908185b5bdd5b9d60421b604482015290519081900360640190fd5b6131b8600b8263ffffffff6145ef16565b6040805182815290517fc178d379965e5657b6fc57494e392f121a14119215dfb422aad7db4cc03f2d109181900360200190a150565b6131f733614310565b613241576040805162461bcd60e51b815260206004820152601660248201527539b2b73232b91034b9903737ba1030b71037bbb732b960511b604482015290519081900360640190fd5b600354600160a01b900460ff1661329f576040805162461bcd60e51b815260206004820152601d60248201527f6f776e657273686970206973206e6f74207472616e7366657261626c65000000604482015290519081900360640190fd5b6001600160a01b0382166132e45760405162461bcd60e51b81526004018080602001828103825260238152602001806153f56023913960400191505060405180910390fd5b6003805460ff60a01b1916600160a01b831515021790558061333d57604080516001600160a01b038416815290517f808639ff9c8e4732d60b6c2330de498035416d229f27a77d259680895efec1229181900360200190a15b600354604080516001600160a01b039283168152918416602083015280517f850b3df64837d7d518b45f5aa64d104652c3b80eb5b34a8e3d9eb666cb7cdea59281900390910190a150600380546001600160a01b0319166001600160a01b0392909216919091179055565b6000806000806133b786614cf2565b50505093509350935050801561342c5781613402576040805162461bcd60e51b81526020600482015260066024820152650726174653d360d41b604482015290519081900360640190fd5b61342283613416878563ffffffff614e8416565b9063ffffffff614edd16565b93505050506114ef565b50600095945050505050565b60195460ff1690565b60145462010000900460ff1681565b61345933614310565b8061346357503330145b6134a7576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b80806134e4576040805162461bcd60e51b8152602060048201526007602482015266076616c75653d360cc1b604482015290519081900360640190fd5b6001600160a01b03841661352f576040805162461bcd60e51b815260206004820152600d60248201526c064657374696e6174696f6e3d3609c1b604482015290519081900360640190fd5b6001600160a01b03841660009081526010602052604090205460ff1661357f57816001600160a01b0384161561356c5761356984846133a8565b90505b61357d60158263ffffffff61466a16565b505b61358a848484614f47565b604080516001600160a01b0380871682528516602082015280820184905290517fd1ba4ac2e2a11b5101f6cb4d978f514a155b421e8ec396d2d9abaf0bb02917ee9181900360600190a150505050565b600e5490565b6135e93361438d565b613628576040805162461bcd60e51b815260206004820152601a60248201526000805160206153b4833981519152604482015290519081900360640190fd5b601454610100900460ff1661367c576040805162461bcd60e51b81526020600482015260156024820152743737903832b73234b7339039bab136b4b9b9b4b7b760591b604482015290519081900360640190fd5b6136dd60138054806020026020016040519081016040528092919081815260200182805480156113cb576020028201919060005260206000209081546001600160a01b031681526001909101906020018083116113ad575050505050613969565b811461371a5760405162461bcd60e51b81526004018080602001828103825260238152602001806154426023913960400191505060405180910390fd5b60005b6013548110156138a857601060006013838154811061373857fe5b60009182526020808320909101546001600160a01b0316835282019290925260400190205460ff16156138a0576000601060006013848154811061377857fe5b6000918252602080832091909101546001600160a01b031683528201929092526040018120805460ff1916921515929092179091555b6011546137c290600163ffffffff614c3b16565b81101561388a57601382815481106137d657fe5b600091825260209091200154601180546001600160a01b0390921691839081106137fc57fe5b6000918252602090912001546001600160a01b031614156138825760118054600019810190811061382957fe5b600091825260209091200154601180546001600160a01b03909216918390811061384f57fe5b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555061388a565b6001016137ae565b50601180549061389e906000198301615333565b505b60010161371d565b507fd218c430fa348f4ce67791021b6b89c0c3eacd4ead1d8f5b83c60038ec28249b33601360405180836001600160a01b03166001600160a01b0316815260200180602001828103825283818154815260200191508054801561393457602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311613916575b5050935050505060405180910390a161394f601360006152b2565b506014805461ff0019169055565b60075490565b601a5481565b60008160405160200180828051906020019060200280838360005b8381101561399c578181015183820152602001613984565b50505050905001915050604051602081830303815290604052805190602001209050919050565b6139cc33614310565b806139db57506139db3361438d565b613a25576040805162461bcd60e51b815260206004820152601660248201527537b7363c9037bbb732b93e3e31b7b73a3937b63632b960511b604482015290519081900360640190fd5b601454610100900460ff16613a79576040805162461bcd60e51b81526020600482015260156024820152743737903832b73234b7339039bab136b4b9b9b4b7b760591b604482015290519081900360640190fd5b613ada60138054806020026020016040519081016040528092919081815260200182805480156113cb576020028201919060005260206000209081546001600160a01b031681526001909101906020018083116113ad575050505050613969565b8114613b175760405162461bcd60e51b81526004018080602001828103825260238152602001806154426023913960400191505060405180910390fd5b613b23601360006152b2565b6014805461ff0019169055604080513381526020810183905281517f13c935eb475aa0f6e931fece83e2ac44569ce2d53460d29a6dedab40b965c8a3929181900390910190a150565b60118181548110613b7957fe5b6000918252602090912001546001600160a01b0316905081565b600b5490565b601454610100900460ff1681565b60025490565b8080613bea576040805162461bcd60e51b8152602060048201526007602482015266076616c75653d360cc1b604482015290519081900360640190fd5b613bf333614310565b80613c025750613c023361438d565b613c4c576040805162461bcd60e51b815260206004820152601660248201527537b7363c9037bbb732b93e3e31b7b73a3937b63632b960511b604482015290519081900360640190fd5b613c5d60048363ffffffff61466a16565b613c65612f46565b6001600160a01b03166108fc839081150290604051600060405180830381858888f19350505050158015613c9d573d6000803e3d6000fd5b507f611b7c0d84fda988026215bef9b3e4d81cbceced7e679be6d5e044b588467c0e33613cc8612f46565b604080516001600160a01b03938416815291909216602082015280820185905290519081900360600190a15050565b613d003361438d565b613d3f576040805162461bcd60e51b815260206004820152601a60248201526000805160206153b4833981519152604482015290519081900360640190fd5b611c3360158263ffffffff614fab16565b6000613d5a614fff565b6001600160a01b0316836001600160a01b03161415613d7a5750806114ef565b816001600160a01b03841615613e3f576000806000613d9887614cf2565b5050509350935093505080613dea576040805162461bcd60e51b8152602060048201526013602482015272746f6b656e206e6f7420617661696c61626c6560681b604482015290519081900360640190fd5b81613e25576040805162461bcd60e51b81526020600482015260066024820152650726174653d360d41b604482015290519081900360640190fd5b613e3983613416888563ffffffff614e8416565b93505050505b6000806000613e4c615075565b5050509350935093505080613e9e576040805162461bcd60e51b8152602060048201526013602482015272746f6b656e206e6f7420617661696c61626c6560681b604482015290519081900360640190fd5b81613ee4576040805162461bcd60e51b81526020600482015260116024820152700737461626c65636f696e20726174653d3607c1b604482015290519081900360640190fd5b613ef882613416868663ffffffff614e8416565b979650505050505050565b613f0c3361438d565b613f4b576040805162461bcd60e51b815260206004820152601a60248201526000805160206153b4833981519152604482015290519081900360640190fd5b611b8d600b8263ffffffff614fab16565b613f653361438d565b613fa4576040805162461bcd60e51b815260206004820152601a60248201526000805160206153b4833981519152604482015290519081900360640190fd5b61128360048263ffffffff614fab16565b613fbe33614310565b80613fc857503330145b61400c576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b8181808060200260200160405190810160405280939291908181526020018383602002808284376000920182905250925050505b815181101561411b57614058828281518110612d0d57fe5b156140a3576040805162461bcd60e51b8152602060048201526016602482015275636f6e7461696e73206f776e6572206164647265737360501b604482015290519081900360640190fd5b60006001600160a01b03168282815181106140ba57fe5b60200260200101516001600160a01b03161415614113576040805162461bcd60e51b8152602060048201526012602482015271636f6e7461696e732030206164647265737360701b604482015290519081900360640190fd5b600101614040565b5060145462010000900460ff1615614172576040805162461bcd60e51b81526020600482015260156024820152741dda1a5d195b1a5cdd081a5b9a5d1a585b1a5e9959605a1b604482015290519081900360640190fd5b60005b82811015614263576010600085858481811061418d57fe5b602090810292909201356001600160a01b03168352508101919091526040016000205460ff1661425b576001601060008686858181106141c957fe5b905060200201356001600160a01b03166001600160a01b03166001600160a01b0316815260200190815260200160002060006101000a81548160ff021916908315150217905550601184848381811061421e57fe5b835460018101855560009485526020948590200180546001600160a01b0319166001600160a01b0395909202939093013593909316929092179055505b600101614175565b506014805462ff0000191662010000179055604080513380825260208201838152601180549484018590527fb2f6cccee7a369e23e293c25aa19bef80af11eb26deba3ea0f2a02783f752e4a94929390929091906060830190849080156142f357602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116142d5575b5050935050505060405180910390a1505050565b60085460ff1690565b6003546001600160a01b0390811691161490565b600482015460ff1615614372576040805162461bcd60e51b81526020600482015260116024820152701b1a5b5a5d08185b1c9958591e481cd95d607a1b604482015290519081900360640190fd5b61437c82826151e4565b50600401805460ff19166001179055565b600061439a6002546146e0565b6001600160a01b031663b429afeb836040518263ffffffff1660e01b815260040180826001600160a01b03166001600160a01b0316815260200191505060206040518083038186803b1580156143ef57600080fd5b505afa158015614403573d6000803e3d6000fd5b505050506040513d602081101561441957600080fd5b505192915050565b60008151604114614434575060006114ef565b60208201516040830151606084015160001a7f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a082111561447a57600093505050506114ef565b8060ff16601b1415801561449257508060ff16601c14155b156144a357600093505050506114ef565b6040805160008152602080820180845289905260ff8416828401526060820186905260808201859052915160019260a0808401939192601f1981019281900390910190855afa1580156144fa573d6000803e3d6000fd5b5050604051601f190151979650505050505050565b60006001600160a01b038216156145a957816001600160a01b03166370a08231846040518263ffffffff1660e01b815260040180826001600160a01b03166001600160a01b0316815260200191505060206040518083038186803b15801561457657600080fd5b505afa15801561458a573d6000803e3d6000fd5b505050506040513d60208110156145a057600080fd5b505190506114ef565b506001600160a01b038216316114ef565b60028101546000906145d5906201518063ffffffff614c9816565b4211156145e4575080546111a9565b5060018101546111a9565b600482015460ff16614648576040805162461bcd60e51b815260206004820152601960248201527f6c696d6974206861736e2774206265656e207365742079657400000000000000604482015290519081900360640190fd5b600390910155565b60008061465c83614cf2565b509098975050505050505050565b61467382615207565b80826001015410156146bf576040805162461bcd60e51b815260206004820152601060248201526f185d985a5b18589b194f185b5bdd5b9d60821b604482015290519081900360640190fd5b60018201546146d4908263ffffffff614c3b16565b82600101819055505050565b6000805460408051630178b8bf60e01b81526004810185905290516001600160a01b0390921691630178b8bf91602480820192602092909190829003018186803b15801561472d57600080fd5b505afa158015614741573d6000803e3d6000fd5b505050506040513d602081101561475757600080fd5b505160408051631d9dabef60e11b81526004810185905290516001600160a01b0390921691633b3b57de91602480820192602092909190829003018186803b1580156143ef57600080fd5b801580614828575060408051636eb1769f60e11b81523060048201526001600160a01b03848116602483015291519185169163dd62ed3e91604480820192602092909190829003018186803b1580156147fa57600080fd5b505afa15801561480e573d6000803e3d6000fd5b505050506040513d602081101561482457600080fd5b5051155b6148635760405162461bcd60e51b81526004018080602001828103825260368152602001806154656036913960400191505060405180910390fd5b604080516001600160a01b038416602482015260448082018490528251808303909101815260649091019091526020810180516001600160e01b031663095ea7b360e01b1790526148b59084906149e4565b505050565b3b151590565b6000806148cc83614cf2565b509198975050505050505050565b6000806148e86009546146e0565b6001600160a01b031663afc72e9385856040518363ffffffff1660e01b815260040180836001600160a01b03166001600160a01b0316815260200180602001828103825283818151815260200191508051906020019080838360005b8381101561495c578181015183820152602001614944565b50505050905090810190601f1680156149895780820380516001836020036101000a031916815260200191505b509350505050604080518083038186803b1580156149a657600080fd5b505afa1580156149ba573d6000803e3d6000fd5b505050506040513d60408110156149d057600080fd5b508051602090910151909590945092505050565b6149f6826001600160a01b03166148ba565b614a47576040805162461bcd60e51b815260206004820152601f60248201527f5361666545524332303a2063616c6c20746f206e6f6e2d636f6e747261637400604482015290519081900360640190fd5b60006060836001600160a01b0316836040518082805190602001908083835b60208310614a855780518252601f199092019160209182019101614a66565b6001836020036101000a0380198251168184511680821785525050505050509050019150506000604051808303816000865af19150503d8060008114614ae7576040519150601f19603f3d011682016040523d82523d6000602084013e614aec565b606091505b509150915081614b43576040805162461bcd60e51b815260206004820181905260248201527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564604482015290519081900360640190fd5b805115614b9c57808060200190516020811015614b5f57600080fd5b5051614b9c5760405162461bcd60e51b815260040180806020018281038252602a815260200180615418602a913960400191505060405180910390fd5b50505050565b604080517f19457468657265756d205369676e6564204d6573736167653a0a333200000000602080830191909152603c8083019490945282518083039094018452605c909101909152815191012090565b601a80546001019081905560408051338152602081019290925280517fab0423a75986556234aecd171c46ce7f5e45607d8070bf5230f2735b50322bff9281900390910190a1565b600082821115614c92576040805162461bcd60e51b815260206004820152601e60248201527f536166654d6174683a207375627472616374696f6e206f766572666c6f770000604482015290519081900360640190fd5b50900390565b600082820183811015612161576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b6060600080600080600080614d086009546146e0565b6001600160a01b0316631f69565f896040518263ffffffff1660e01b815260040180826001600160a01b03166001600160a01b0316815260200191505060006040518083038186803b158015614d5d57600080fd5b505afa158015614d71573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405260e0811015614d9a57600080fd5b8101908080516040519392919084600160201b821115614db957600080fd5b908301906020820185811115614dce57600080fd5b8251600160201b811182820188101715614de757600080fd5b82525081516020918201929091019080838360005b83811015614e14578181015183820152602001614dfc565b50505050905090810190601f168015614e415780820380516001836020036101000a031916815260200191505b5060409081526020820151908201516060830151608084015160a085015160c090950151979e50929c50909a509850965094509192505050919395979092949650565b600082614e93575060006114ef565b82820282848281614ea057fe5b04146121615760405162461bcd60e51b81526004018080602001828103825260218152602001806153d46021913960400191505060405180910390fd5b6000808211614f33576040805162461bcd60e51b815260206004820152601a60248201527f536166654d6174683a206469766973696f6e206279207a65726f000000000000604482015290519081900360640190fd5b6000828481614f3e57fe5b04949350505050565b6001600160a01b038216614f91576040516001600160a01b0384169082156108fc029083906000818181858888f19350505050158015614f8b573d6000803e3d6000fd5b506148b5565b6148b56001600160a01b038316848363ffffffff61526016565b80826003015414614fed5760405162461bcd60e51b81526004018080602001828103825260228152602001806153926022913960400191505060405180910390fd5b614ffb8283600301546151e4565b5050565b600061500c6009546146e0565b6001600160a01b031663e9cbd8226040518163ffffffff1660e01b815260040160206040518083038186803b15801561504457600080fd5b505afa158015615058573d6000803e3d6000fd5b505050506040513d602081101561506e57600080fd5b5051905090565b606060008060008060008061508b6009546146e0565b6001600160a01b0316633efec5e96040518163ffffffff1660e01b815260040160006040518083038186803b1580156150c357600080fd5b505afa1580156150d7573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405260e081101561510057600080fd5b8101908080516040519392919084600160201b82111561511f57600080fd5b90830190602082018581111561513457600080fd5b8251600160201b81118282018810171561514d57600080fd5b82525081516020918201929091019080838360005b8381101561517a578181015183820152602001615162565b50505050905090810190601f1680156151a75780820380516001836020036101000a031916815260200191505b5060409081526020820151908201516060830151608084015160a085015160c090950151979f939e50919c509a5098509096509294509192505050565b6151ed82615207565b8082556001820154811015614ffb57815460018301555050565b600281015461521f906201518063ffffffff614c9816565b42111561525d57426002820155805460018201556040517fe93bc25276d408d390778e7a8b926f2f67209c43ed540081b951fe128f0d3cd290600090a15b50565b604080516001600160a01b038416602482015260448082018490528251808303909101815260649091019091526020810180516001600160e01b031663a9059cbb60e01b1790526148b59084906149e4565b508054600082559060005260206000209081019061525d9190615353565b828054828255906000526020600020908101928215615323579160200282015b828111156153235781546001600160a01b0319166001600160a01b038435161782556020909201916001909101906152f0565b5061532f92915061536d565b5090565b8154818355818111156148b5576000838152602090206148b59181019083015b6111b291905b8082111561532f5760008155600101615359565b6111b291905b8082111561532f5780546001600160a01b031916815560010161537356fe636f6e6669726d65642f7375626d6974746564206c696d6974206d69736d6174636873656e646572206973206e6f74206120636f6e74726f6c6c6572000000000000536166654d6174683a206d756c7469706c69636174696f6e206f766572666c6f776f776e65722063616e6e6f742062652073657420746f207a65726f20616464726573735361666545524332303a204552433230206f7065726174696f6e20646964206e6f7420737563636565646e6f6e2d6d61746368696e672070656e64696e672077686974656c69737420686173685361666545524332303a20617070726f76652066726f6d206e6f6e2d7a65726f20746f206e6f6e2d7a65726f20616c6c6f77616e6365a265627a7a7231582026e1c268487cb71f6386771c751057a76de75c808456898744d851d2655c5dcc64736f6c634300050f0032"
||||||| constructed merge base
var WalletBin = "0x60806040527f7f2ce995617d2816b426c5c8698c5ec2952f7a34bb10f38326f74933d58936976002553480156200003557600080fd5b5060405162005a6738038062005a67833981810160405260e08110156200005b57600080fd5b508051602082015160408301516060840151608085015160a086015160c090960151600180546001600160a01b038087166001600160a01b031992831617928390556000805490921692169190911790559495939492939192909190868685858482818686858015620000ce5760028190555b50600380546001600160a01b0319166001600160a01b0384161760ff60a01b1916600160a01b8315158102919091179182905560ff9104166200014857604080516001600160a01b038416815290517f808639ff9c8e4732d60b6c2330de498035416d229f27a77d259680895efec1229181900360200190a15b60408051600081526001600160a01b038416602082015281517f850b3df64837d7d518b45f5aa64d104652c3b80eb5b34a8e3d9eb666cb7cdea5929181900390910190a150506040805160a080820183528382526020808301859052428385018190526000606080860182905260809586018290526009889055600a97909755600b829055600c819055600d805460ff19908116909155600e98909855855193840186526706f05b59d3b20000808552928401839052948301819052948201849052910182905260108190556011556012919091556013819055601480549092169091559450620002479350506001600160e01b03620002be16915050565b50505050509150506000811162000295576040805162461bcd60e51b815260206004820152600d60248201526c37379039ba30b13632b1b7b4b760991b604482015290519081900360640190fd5b620002ad61271082026001600160e01b036200044016565b5050601b5550620005889350505050565b6060600080600080600080620002dc600e546200048e60201b60201c565b6001600160a01b0316633efec5e96040518163ffffffff1660e01b815260040160006040518083038186803b1580156200031557600080fd5b505afa1580156200032a573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405260e08110156200035457600080fd5b81019080805160405193929190846401000000008211156200037557600080fd5b9083019060208201858111156200038b57600080fd5b8251640100000000811182820188101715620003a657600080fd5b82525081516020918201929091019080838360005b83811015620003d5578181015183820152602001620003bb565b50505050905090810190601f168015620004035780820380516001836020036101000a031916815260200191505b5060409081526020820151908201516060830151608084015160a085015160c090950151979f939e50919c509a5098509096509294509192505050565b60158190556040805160a081018252828152602081018390524291810182905260006060820181905260809091018190526016839055601792909255601855601955601a805460ff19169055565b6000805460408051630178b8bf60e01b81526004810185905290516001600160a01b0390921691630178b8bf91602480820192602092909190829003018186803b158015620004dc57600080fd5b505afa158015620004f1573d6000803e3d6000fd5b505050506040513d60208110156200050857600080fd5b505160408051631d9dabef60e11b81526004810185905290516001600160a01b0390921691633b3b57de91602480820192602092909190829003018186803b1580156200055457600080fd5b505afa15801562000569573d6000803e3d6000fd5b505050506040513d60208110156200058057600080fd5b505192915050565b6154cf80620005986000396000f3fe60806040526004361061038c5760003560e01c80637fd004fa116101dc578063cc0e7e5611610102578063e2b4ce97116100a0578063f40b51f81161006f578063f40b51f8146110b0578063f41c4319146110da578063f421764814611104578063f776f5181461117f5761038c565b8063e2b4ce971461100e578063e61c51ca14611023578063eadd3cea1461104d578063f36febda146110775761038c565b8063ce0b5bd5116100dc578063ce0b5bd514610f90578063d251fefc14610fba578063da84b1ed14610fe4578063de212bf314610ff95761038c565b8063cc0e7e5614610eb8578063cccdc55614610ecd578063cd7958dd14610ee25761038c565b8063b221f3161161017a578063be40ba7911610149578063be40ba7914610e21578063beabacc814610e36578063c4856cd914610e79578063cbd2ac6814610e8e5761038c565b8063b221f31614610d6e578063b242e53414610d98578063b87e21ef14610dd3578063bcb8b74a14610e0c5761038c565b806390e690c7116101b657806390e690c714610c7e5780639b0dfd2714610c93578063aaf1fc6214610ca8578063ab20599314610d595761038c565b80637fd004fa14610bd9578063877337b014610c545780638da5cb5b14610c695761038c565b806332531c3c116102c15780635adc02ab1161025f57806374624c551161022e57806374624c5514610b54578063747c31d614610b7e5780637d73b23114610b935780637d7d004614610bc45761038c565b80635adc02ab14610a855780635d2362a814610aaf5780636137d67014610ac4578063715018a614610b3f5761038c565b80633c672eb71161029b5780633c672eb7146108ae5780633f579f42146108d857806346efe0ed1461099e57806347b55a9d14610a705761038c565b806332531c3c146108255780633a43199f146108585780633bfec254146108845761038c565b80631efd02991161032e57806321ce918d1161030857806321ce918d1461076c5780632587a6a21461079657806326d05ab2146107ab578063294f4025146107c05761038c565b80631efd02991461067757806320c13b0b1461068c5780632121dc75146107575761038c565b8063100f23fd1161036a578063100f23fd146104635780631127b57e1461048d5780631626ba7e146105175780631aa21fba146105ec5761038c565b806301ffc9a7146103c8578063027ef3eb146104105780630f3a85d814610437575b6040805133815234602082015281517f88a5966d370b9919b20f3e2c13ff65706f196a4e32cc2c12bf57088f88525874929181900390910190a1005b3480156103d457600080fd5b506103fc600480360360208110156103eb57600080fd5b50356001600160e01b031916611194565b604080519115158252519081900360200190f35b34801561041c57600080fd5b506104256111ae565b60408051918252519081900360200190f35b34801561044357600080fd5b506104616004803603602081101561045a57600080fd5b50356111b5565b005b34801561046f57600080fd5b506104616004803603602081101561048657600080fd5b50356112c1565b34801561049957600080fd5b506104a2611466565b6040805160208082528351818301528351919283929083019185019080838360005b838110156104dc5781810151838201526020016104c4565b50505050905090810190601f1680156105095780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561052357600080fd5b506105cf6004803603604081101561053a57600080fd5b81359190810190604081016020820135600160201b81111561055b57600080fd5b82018360208201111561056d57600080fd5b803590602001918460018302840111600160201b8311171561058e57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550611487945050505050565b604080516001600160e01b03199092168252519081900360200190f35b3480156105f857600080fd5b506104616004803603604081101561060f57600080fd5b6001600160a01b038235169190810190604081016020820135600160201b81111561063957600080fd5b82018360208201111561064b57600080fd5b803590602001918460208302840111600160201b8311171561066c57600080fd5b5090925090506114f5565b34801561068357600080fd5b5061042561167b565b34801561069857600080fd5b506105cf600480360360408110156106af57600080fd5b810190602081018135600160201b8111156106c957600080fd5b8201836020820111156106db57600080fd5b803590602001918460018302840111600160201b831117156106fc57600080fd5b919390929091602081019035600160201b81111561071957600080fd5b82018360208201111561072b57600080fd5b803590602001918460018302840111600160201b8311171561074c57600080fd5b50909250905061168c565b34801561076357600080fd5b506103fc611761565b34801561077857600080fd5b506104616004803603602081101561078f57600080fd5b5035611771565b3480156107a257600080fd5b5061042561180f565b3480156107b757600080fd5b506103fc611815565b3480156107cc57600080fd5b506107d561181e565b60408051602080825283518183015283519192839290830191858101910280838360005b838110156108115781810151838201526020016107f9565b505050509050019250505060405180910390f35b34801561083157600080fd5b506103fc6004803603602081101561084857600080fd5b50356001600160a01b0316611880565b6104616004803603604081101561086e57600080fd5b506001600160a01b038135169060200135611895565b34801561089057600080fd5b50610461600480360360208110156108a757600080fd5b5035611ad3565b3480156108ba57600080fd5b50610461600480360360208110156108d157600080fd5b5035611bcb565b3480156108e457600080fd5b506104a2600480360360608110156108fb57600080fd5b6001600160a01b0382351691602081013591810190606081016040820135600160201b81111561092a57600080fd5b82018360208201111561093c57600080fd5b803590602001918460018302840111600160201b8311171561095d57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550611c71945050505050565b3480156109aa57600080fd5b50610461600480360360608110156109c157600080fd5b81359190810190604081016020820135600160201b8111156109e257600080fd5b8201836020820111156109f457600080fd5b803590602001918460018302840111600160201b83111715610a1557600080fd5b919390929091602081019035600160201b811115610a3257600080fd5b820183602082011115610a4457600080fd5b803590602001918460018302840111600160201b83111715610a6557600080fd5b509092509050612168565b348015610a7c57600080fd5b506107d5612479565b348015610a9157600080fd5b5061046160048036036020811015610aa857600080fd5b50356124d9565b348015610abb57600080fd5b506104256127a9565b348015610ad057600080fd5b5061046160048036036020811015610ae757600080fd5b810190602081018135600160201b811115610b0157600080fd5b820183602082011115610b1357600080fd5b803590602001918460208302840111600160201b83111715610b3457600080fd5b5090925090506127b5565b348015610b4b57600080fd5b506104616129db565b348015610b6057600080fd5b5061046160048036036020811015610b7757600080fd5b5035612ad9565b348015610b8a57600080fd5b50610425612bdd565b348015610b9f57600080fd5b50610ba8612be3565b604080516001600160a01b039092168252519081900360200190f35b348015610bd057600080fd5b50610425612bf2565b348015610be557600080fd5b5061046160048036036020811015610bfc57600080fd5b810190602081018135600160201b811115610c1657600080fd5b820183602082011115610c2857600080fd5b803590602001918460208302840111600160201b83111715610c4957600080fd5b509092509050612bfe565b348015610c6057600080fd5b50610425612f40565b348015610c7557600080fd5b50610ba8612f46565b348015610c8a57600080fd5b50610461612f55565b348015610c9f57600080fd5b50610425612fb2565b348015610cb457600080fd5b5061046160048036036020811015610ccb57600080fd5b810190602081018135600160201b811115610ce557600080fd5b820183602082011115610cf757600080fd5b803590602001918460018302840111600160201b83111715610d1857600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550612fb8945050505050565b348015610d6557600080fd5b506103fc6130f5565b348015610d7a57600080fd5b5061046160048036036020811015610d9157600080fd5b50356130fe565b348015610da457600080fd5b5061046160048036036040811015610dbb57600080fd5b506001600160a01b03813516906020013515156131ee565b348015610ddf57600080fd5b5061042560048036036040811015610df657600080fd5b506001600160a01b0381351690602001356133a8565b348015610e1857600080fd5b506103fc613438565b348015610e2d57600080fd5b506103fc613441565b348015610e4257600080fd5b5061046160048036036060811015610e5957600080fd5b506001600160a01b03813581169160208101359091169060400135613450565b348015610e8557600080fd5b506104256135da565b348015610e9a57600080fd5b5061046160048036036020811015610eb157600080fd5b50356135e0565b348015610ec457600080fd5b5061042561395d565b348015610ed957600080fd5b50610425613963565b348015610eee57600080fd5b5061042560048036036020811015610f0557600080fd5b810190602081018135600160201b811115610f1f57600080fd5b820183602082011115610f3157600080fd5b803590602001918460208302840111600160201b83111715610f5257600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929550613969945050505050565b348015610f9c57600080fd5b5061046160048036036020811015610fb357600080fd5b50356139c3565b348015610fc657600080fd5b50610ba860048036036020811015610fdd57600080fd5b5035613b6c565b348015610ff057600080fd5b50610425613b93565b34801561100557600080fd5b506103fc613b99565b34801561101a57600080fd5b50610425613ba7565b34801561102f57600080fd5b506104616004803603602081101561104657600080fd5b5035613bad565b34801561105957600080fd5b506104616004803603602081101561107057600080fd5b5035613cf7565b34801561108357600080fd5b506104256004803603604081101561109a57600080fd5b506001600160a01b038135169060200135613d50565b3480156110bc57600080fd5b50610461600480360360208110156110d357600080fd5b5035613f03565b3480156110e657600080fd5b50610461600480360360208110156110fd57600080fd5b5035613f5c565b34801561111057600080fd5b506104616004803603602081101561112757600080fd5b810190602081018135600160201b81111561114157600080fd5b82018360208201111561115357600080fd5b803590602001918460208302840111600160201b8311171561117457600080fd5b509092509050613fb5565b34801561118b57600080fd5b506103fc614307565b6001600160e01b031981166301ffc9a760e01b145b919050565b600c545b90565b6111be33614310565b806111c857503330145b61120c576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b8066038d7ea4c680001115801561122b57506706f05b59d3b200008111155b611272576040805162461bcd60e51b815260206004820152601360248201527206f7574206f662072616e676520746f702d757606c1b604482015290519081900360640190fd5b61128360108263ffffffff61432416565b604080513381526020810183905281517f41ff5d5ce3b7935893a4e7269ec5caae9cca5e3bf0eb4b21d2f443489667112e929181900390910190a150565b6112ca33614310565b806112d957506112d93361438d565b611323576040805162461bcd60e51b815260206004820152601660248201527537b7363c9037bbb732b93e3e31b7b73a3937b63632b960511b604482015290519081900360640190fd5b60085460ff16611372576040805162461bcd60e51b81526020600482015260156024820152743737903832b73234b7339039bab136b4b9b9b4b7b760591b604482015290519081900360640190fd5b6113d560068054806020026020016040519081016040528092919081815260200182805480156113cb57602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116113ad575b5050505050613969565b81146114125760405162461bcd60e51b81526004018080602001828103825260238152602001806154426023913960400191505060405180910390fd5b61141e600660006152b2565b6008805460ff19169055604080513381526020810183905281517f7794eff834d760583543e6e510e717a5e66d2c064e225f4db448343c3e66afcf929181900390910190a150565b604051806040016040528060058152602001640332e312e360dc1b81525081565b60008061149a848463ffffffff61442116565b90506114a581614310565b6114e3576040805162461bcd60e51b815260206004820152600a60248201526937b7363c9037bbb732b960b11b604482015290519081900360640190fd5b50630b135d3f60e11b90505b92915050565b6114fe33614310565b8061150857503330145b61154c576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b80611595576040805162461bcd60e51b8152602060048201526014602482015273617373657420617272617920697320656d70747960601b604482015290519081900360640190fd5b60005b818110156115f85760006115c7308585858181106115b257fe5b905060200201356001600160a01b031661450f565b90506115ef858585858181106115d957fe5b905060200201356001600160a01b031683613450565b50600101611598565b507fd4f62f23021706247dcffea245d104ae7ddaec7f23acf3d11d7136d5de6a69ad83838360405180846001600160a01b03166001600160a01b03168152602001806020018281038252848482818152602001925060200280828437600083820152604051601f909101601f1916909201829003965090945050505050a1505050565b600061168760166145ba565b905090565b6000808585604051602001808383808284376040805191909301818103601f190182528084528151602092830120601f8b01839004830282018301909452898152929650630b135d3f60e11b9550611703945086935089915088908190840183828082843760009201919091525061148792505050565b6001600160e01b0319161461174f576040805162461bcd60e51b815260206004820152600d60248201526c1cda59c81b9bdd081d985b1a59609a1b604482015290519081900360640190fd5b506320c13b0b60e01b95945050505050565b600354600160a01b900460ff1690565b61177a33614310565b8061178457503330145b6117c8576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b6117d960098263ffffffff6145ef16565b6040805182815290517f4b1b970c8a0fa761e7803ed70c13d7aca71904b13df60fbe03f981da1730da919181900360200190a150565b60105490565b60085460ff1681565b6060600780548060200260200160405190810160405280929190818152602001828054801561187657602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311611858575b5050505050905090565b60046020526000908152604090205460ff1681565b61189e33614310565b806118a857503330145b6118ec576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b6118f582614650565b61193b576040805162461bcd60e51b8152602060048201526012602482015271746f6b656e206e6f74206c6f616461626c6560701b604482015290519081900360640190fd5b60006119478383613d50565b905061195a60168263ffffffff61466a16565b6000611967601b546146e0565b90506001600160a01b03841615611a0f576119926001600160a01b038516828563ffffffff6147a216565b806001600160a01b0316631b3c96b485856040518363ffffffff1660e01b815260040180836001600160a01b03166001600160a01b0316815260200182815260200192505050600060405180830381600087803b1580156119f257600080fd5b505af1158015611a06573d6000803e3d6000fd5b50505050611a89565b806001600160a01b0316631b3c96b48486866040518463ffffffff1660e01b815260040180836001600160a01b03166001600160a01b03168152602001828152602001925050506000604051808303818588803b158015611a6f57600080fd5b505af1158015611a83573d6000803e3d6000fd5b50505050505b604080516001600160a01b03861681526020810185905281517f5f65674bec9af81f71be68674135a0ea3f163fb91984e3893d06da9f6ea2ce8a929181900390910190a150505050565b611adc33614310565b80611ae657503330145b611b2a576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b601554811115611b7c576040805162461bcd60e51b81526020600482015260186024820152771bdd5d081bd9881c985b99d9481b1bd85908185b5bdd5b9d60421b604482015290519081900360640190fd5b611b8d60168263ffffffff61432416565b604080513381526020810183905281517f0b05243483e17c3f3377aee82b7d47e5700b48288695fc08b7ecc2759afa44ef929181900390910190a150565b611bd433614310565b80611bde57503330145b611c22576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b611c3360098263ffffffff61432416565b604080513381526020810183905281517f068f112e5ec923d412be64779fe69e0fcbb6784c6617e94cccc8fd348f2e0f21929181900390910190a150565b6060611c7c33614310565b80611c8657503330145b611cca576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b6001600160a01b03841660009081526004602052604090205460ff16611cfb57611cfb60098463ffffffff61466a16565b611d0d846001600160a01b03166148ba565b8015611d1d5750611d1d846148c0565b15611f0457600080611d2f86856148da565b6001600160a01b038216600090815260046020526040902054919350915060ff16611d75576000611d6087836133a8565b9050611d7360098263ffffffff61466a16565b505b611d8e6001600160a01b0387168563ffffffff6149e416565b604080516020808252818301909252606091602082018180388339019050509050600160f81b81601f81518110611dc157fe5b60200101906001600160f81b031916908160001a9053507ff77753fab406ecfff96d6ff2476c64a838fa9f6d37b1bf190f8546e395e3b6138787878460405180856001600160a01b03166001600160a01b031681526020018481526020018060200180602001838103835285818151815260200191508051906020019080838360005b83811015611e5c578181015183820152602001611e44565b50505050905090810190601f168015611e895780820380516001836020036101000a031916815260200191505b50838103825284518152845160209182019186019080838360005b83811015611ebc578181015183820152602001611ea4565b50505050905090810190601f168015611ee95780820380516001836020036101000a031916815260200191505b50965050505050505060405180910390a19250612161915050565b60006060856001600160a01b031685856040518082805190602001908083835b60208310611f435780518252601f199092019160209182019101611f24565b6001836020036101000a03801982511681845116808217855250505050505090500191505060006040518083038185875af1925050503d8060008114611fa5576040519150601f19603f3d011682016040523d82523d6000602084013e611faa565b606091505b509150915081819061203a5760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b83811015611fff578181015183820152602001611fe7565b50505050905090810190601f16801561202c5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b507ff77753fab406ecfff96d6ff2476c64a838fa9f6d37b1bf190f8546e395e3b6138686868460405180856001600160a01b03166001600160a01b031681526020018481526020018060200180602001838103835285818151815260200191508051906020019080838360005b838110156120bf5781810151838201526020016120a7565b50505050905090810190601f1680156120ec5780820380516001836020036101000a031916815260200191505b50838103825284518152845160209182019186019080838360005b8381101561211f578181015183820152602001612107565b50505050905090810190601f16801561214c5780820380516001836020036101000a031916815260200191505b50965050505050505060405180910390a19150505b9392505050565b6121713361438d565b6121b0576040805162461bcd60e51b815260206004820152601a60248201526000805160206153b4833981519152604482015290519081900360640190fd5b600061220186868660405160200180806339363c1d60e11b81525060040184815260200183838082843780830192505050935050505060405160208183030381529060405280519060200120614ba2565b9050631626ba7e60e01b6001600160e01b0319166122558285858080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061148792505050565b6001600160e01b031916146122a1576040805162461bcd60e51b815260206004820152600d60248201526c1cda59c81b9bdd081d985b1a59609a1b604482015290519081900360640190fd5b600f5486146122e3576040805162461bcd60e51b81526020600482015260096024820152687478207265706c617960b81b604482015290519081900360640190fd5b6122eb614bf3565b60006060306001600160a01b03168787604051808383808284376040519201945060009350909150508083038183865af19150503d806000811461234b576040519150601f19603f3d011682016040523d82523d6000602084013e612350565b606091505b50915091508181906123a35760405162461bcd60e51b8152602060048201818152835160248401528351909283926044909101919085019080838360008315611fff578181015183820152602001611fe7565b507f823dbcf2b7b0f265871963ca65ac033f6b4c71e0d82cd123d2ff23d752dc21c187878360405180806020018060200183810383528686828181526020019250808284376000838201819052601f909101601f191690920185810384528651815286516020918201939188019250908190849084905b8381101561243257818101518382015260200161241a565b50505050905090810190601f16801561245f5780820380516001836020036101000a031916815260200191505b509550505050505060405180910390a15050505050505050565b60606006805480602002602001604051908101604052809291908181526020018280548015611876576020028201919060005260206000209081546001600160a01b03168152600190910190602001808311611858575050505050905090565b6124e23361438d565b612521576040805162461bcd60e51b815260206004820152601a60248201526000805160206153b4833981519152604482015290519081900360640190fd5b60085460ff16612570576040805162461bcd60e51b81526020600482015260156024820152743737903832b73234b7339039bab136b4b9b9b4b7b760591b604482015290519081900360640190fd5b6125d160068054806020026020016040519081016040528092919081815260200182805480156113cb576020028201919060005260206000209081546001600160a01b031681526001909101906020018083116113ad575050505050613969565b811461260e5760405162461bcd60e51b81526004018080602001828103825260238152602001806154426023913960400191505060405180910390fd5b60005b6006548110156126f557600460006006838154811061262c57fe5b60009182526020808320909101546001600160a01b0316835282019290925260400190205460ff166126ed576001600460006006848154811061266b57fe5b6000918252602080832091909101546001600160a01b031683528201929092526040019020805460ff19169115159190911790556006805460059190839081106126b157fe5b60009182526020808320909101548354600181018555938352912090910180546001600160a01b0319166001600160a01b039092169190911790555b600101612611565b507fb2f6cccee7a369e23e293c25aa19bef80af11eb26deba3ea0f2a02783f752e4a33600660405180836001600160a01b03166001600160a01b0316815260200180602001828103825283818154815260200191508054801561278157602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311612763575b5050935050505060405180910390a161279c600660006152b2565b506008805460ff19169055565b600061168760096145ba565b6127be33614310565b806127c857503330145b61280c576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b60085460ff161580156128275750600854610100900460ff16155b612878576040805162461bcd60e51b815260206004820152601c60248201527f77686974656c6973742073756d62697373696f6e2070656e64696e6700000000604482015290519081900360640190fd5b60085462010000900460ff166128d1576040805162461bcd60e51b81526020600482015260196024820152781dda1a5d195b1a5cdd081b9bdd081a5b9a5d1a585b1a5e9959603a1b604482015290519081900360640190fd5b80612915576040805162461bcd60e51b815260206004820152600f60248201526e195b5c1d1e481dda1a5d195b1a5cdd608a1b604482015290519081900360640190fd5b612921600783836152d0565b506008805461ff00191661010017905560408051602080840282810182019093528382527ffbc0e5ca6c7e4858daf0fdb185ef5186203e74ec9c64737e93c0aeaec596e1d192859285926129909285918591829185019084908082843760009201919091525061396992505050565b60405180806020018381526020018281038252858582818152602001925060200280828437600083820152604051601f909101601f1916909201829003965090945050505050a15050565b6129e433614310565b612a2e576040805162461bcd60e51b815260206004820152601660248201527539b2b73232b91034b9903737ba1030b71037bbb732b960511b604482015290519081900360640190fd5b600354600160a01b900460ff16612a8c576040805162461bcd60e51b815260206004820152601d60248201527f6f776e657273686970206973206e6f74207472616e7366657261626c65000000604482015290519081900360640190fd5b600380546001600160a01b0319169055604080516000808252602082015281517f850b3df64837d7d518b45f5aa64d104652c3b80eb5b34a8e3d9eb666cb7cdea5929181900390910190a1565b612ae233614310565b80612aec57503330145b612b30576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b8066038d7ea4c6800011158015612b4f57506706f05b59d3b200008111155b612b96576040805162461bcd60e51b815260206004820152601360248201527206f7574206f662072616e676520746f702d757606c1b604482015290519081900360640190fd5b612ba760108263ffffffff6145ef16565b6040805182815290517faf2a77cd04c3cc155588dd3bf67b310ab4fb3b1da3cf6b8d7d4d2aa1d09b794c9181900360200190a150565b601b5490565b6001546001600160a01b031690565b600061168760106145ba565b612c0733614310565b80612c1157503330145b612c55576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b60085460ff16158015612c705750600854610100900460ff16155b612cc1576040805162461bcd60e51b815260206004820152601c60248201527f77686974656c6973742073756d62697373696f6e2070656e64696e6700000000604482015290519081900360640190fd5b8181808060200260200160405190810160405280939291908181526020018383602002808284376000920182905250925050505b8151811015612ddd57612d1a828281518110612d0d57fe5b6020026020010151614310565b15612d65576040805162461bcd60e51b8152602060048201526016602482015275636f6e7461696e73206f776e6572206164647265737360501b604482015290519081900360640190fd5b60006001600160a01b0316828281518110612d7c57fe5b60200260200101516001600160a01b03161415612dd5576040805162461bcd60e51b8152602060048201526012602482015271636f6e7461696e732030206164647265737360701b604482015290519081900360640190fd5b600101612cf5565b5060085462010000900460ff16612e37576040805162461bcd60e51b81526020600482015260196024820152781dda1a5d195b1a5cdd081b9bdd081a5b9a5d1a585b1a5e9959603a1b604482015290519081900360640190fd5b81612e7b576040805162461bcd60e51b815260206004820152600f60248201526e195b5c1d1e481dda1a5d195b1a5cdd608a1b604482015290519081900360640190fd5b612e87600684846152d0565b506008805460ff1916600117905560408051602080850282810182019093528482527f9c80b3b5f68b3e017766d59e8d09b34efe6462b05c398f35cab9e271d9bc3b9c9286928692612ef49285918591829185019084908082843760009201919091525061396992505050565b60405180806020018381526020018281038252858582818152602001925060200280828437600083820152604051601f909101601f1916909201829003965090945050505050a1505050565b600e5490565b6003546001600160a01b031690565b612f5e33614310565b612fa8576040805162461bcd60e51b815260206004820152601660248201527539b2b73232b91034b9903737ba1030b71037bbb732b960511b604482015290519081900360640190fd5b612fb0614bf3565b565b60095490565b612fc133614310565b80612fcb57503330145b61300f576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b8051602080820191906000808060605b868510156130eb5761303886605463ffffffff614c3b16565b888601805160148201516034909201805193995060609190911c9650909450909250905061307d6054613071878563ffffffff614c9816565b9063ffffffff614c9816565b9450868511156130c4576040805162461bcd60e51b815260206004820152600d60248201526c6f7574206f6620626f756e647360981b604482015290519081900360640190fd5b816130da57506040805160208101909152600081525b6130e5848483611c71565b5061301f565b5050505050505050565b601a5460ff1690565b61310733614310565b8061311157503330145b613155576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b6015548111156131a7576040805162461bcd60e51b81526020600482015260186024820152771bdd5d081bd9881c985b99d9481b1bd85908185b5bdd5b9d60421b604482015290519081900360640190fd5b6131b860168263ffffffff6145ef16565b6040805182815290517fc178d379965e5657b6fc57494e392f121a14119215dfb422aad7db4cc03f2d109181900360200190a150565b6131f733614310565b613241576040805162461bcd60e51b815260206004820152601660248201527539b2b73232b91034b9903737ba1030b71037bbb732b960511b604482015290519081900360640190fd5b600354600160a01b900460ff1661329f576040805162461bcd60e51b815260206004820152601d60248201527f6f776e657273686970206973206e6f74207472616e7366657261626c65000000604482015290519081900360640190fd5b6001600160a01b0382166132e45760405162461bcd60e51b81526004018080602001828103825260238152602001806153f56023913960400191505060405180910390fd5b6003805460ff60a01b1916600160a01b831515021790558061333d57604080516001600160a01b038416815290517f808639ff9c8e4732d60b6c2330de498035416d229f27a77d259680895efec1229181900360200190a15b600354604080516001600160a01b039283168152918416602083015280517f850b3df64837d7d518b45f5aa64d104652c3b80eb5b34a8e3d9eb666cb7cdea59281900390910190a150600380546001600160a01b0319166001600160a01b0392909216919091179055565b6000806000806133b786614cf2565b50505093509350935050801561342c5781613402576040805162461bcd60e51b81526020600482015260066024820152650726174653d360d41b604482015290519081900360640190fd5b61342283613416878563ffffffff614e8416565b9063ffffffff614edd16565b93505050506114ef565b50600095945050505050565b600d5460ff1690565b60085462010000900460ff1681565b61345933614310565b8061346357503330145b6134a7576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b80806134e4576040805162461bcd60e51b8152602060048201526007602482015266076616c75653d360cc1b604482015290519081900360640190fd5b6001600160a01b03841661352f576040805162461bcd60e51b815260206004820152600d60248201526c064657374696e6174696f6e3d3609c1b604482015290519081900360640190fd5b6001600160a01b03841660009081526004602052604090205460ff1661357f57816001600160a01b0384161561356c5761356984846133a8565b90505b61357d60098263ffffffff61466a16565b505b61358a848484614f47565b604080516001600160a01b0380871682528516602082015280820184905290517fd1ba4ac2e2a11b5101f6cb4d978f514a155b421e8ec396d2d9abaf0bb02917ee9181900360600190a150505050565b60195490565b6135e93361438d565b613628576040805162461bcd60e51b815260206004820152601a60248201526000805160206153b4833981519152604482015290519081900360640190fd5b600854610100900460ff1661367c576040805162461bcd60e51b81526020600482015260156024820152743737903832b73234b7339039bab136b4b9b9b4b7b760591b604482015290519081900360640190fd5b6136dd60078054806020026020016040519081016040528092919081815260200182805480156113cb576020028201919060005260206000209081546001600160a01b031681526001909101906020018083116113ad575050505050613969565b811461371a5760405162461bcd60e51b81526004018080602001828103825260238152602001806154426023913960400191505060405180910390fd5b60005b6007548110156138a857600460006007838154811061373857fe5b60009182526020808320909101546001600160a01b0316835282019290925260400190205460ff16156138a0576000600460006007848154811061377857fe5b6000918252602080832091909101546001600160a01b031683528201929092526040018120805460ff1916921515929092179091555b6005546137c290600163ffffffff614c3b16565b81101561388a57600782815481106137d657fe5b600091825260209091200154600580546001600160a01b0390921691839081106137fc57fe5b6000918252602090912001546001600160a01b031614156138825760058054600019810190811061382957fe5b600091825260209091200154600580546001600160a01b03909216918390811061384f57fe5b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555061388a565b6001016137ae565b50600580549061389e906000198301615333565b505b60010161371d565b507fd218c430fa348f4ce67791021b6b89c0c3eacd4ead1d8f5b83c60038ec28249b33600760405180836001600160a01b03166001600160a01b0316815260200180602001828103825283818154815260200191508054801561393457602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311613916575b5050935050505060405180910390a161394f600760006152b2565b506008805461ff0019169055565b60135490565b600f5481565b60008160405160200180828051906020019060200280838360005b8381101561399c578181015183820152602001613984565b50505050905001915050604051602081830303815290604052805190602001209050919050565b6139cc33614310565b806139db57506139db3361438d565b613a25576040805162461bcd60e51b815260206004820152601660248201527537b7363c9037bbb732b93e3e31b7b73a3937b63632b960511b604482015290519081900360640190fd5b600854610100900460ff16613a79576040805162461bcd60e51b81526020600482015260156024820152743737903832b73234b7339039bab136b4b9b9b4b7b760591b604482015290519081900360640190fd5b613ada60078054806020026020016040519081016040528092919081815260200182805480156113cb576020028201919060005260206000209081546001600160a01b031681526001909101906020018083116113ad575050505050613969565b8114613b175760405162461bcd60e51b81526004018080602001828103825260238152602001806154426023913960400191505060405180910390fd5b613b23600760006152b2565b6008805461ff0019169055604080513381526020810183905281517f13c935eb475aa0f6e931fece83e2ac44569ce2d53460d29a6dedab40b965c8a3929181900390910190a150565b60058181548110613b7957fe5b6000918252602090912001546001600160a01b0316905081565b60165490565b600854610100900460ff1681565b60025490565b8080613bea576040805162461bcd60e51b8152602060048201526007602482015266076616c75653d360cc1b604482015290519081900360640190fd5b613bf333614310565b80613c025750613c023361438d565b613c4c576040805162461bcd60e51b815260206004820152601660248201527537b7363c9037bbb732b93e3e31b7b73a3937b63632b960511b604482015290519081900360640190fd5b613c5d60108363ffffffff61466a16565b613c65612f46565b6001600160a01b03166108fc839081150290604051600060405180830381858888f19350505050158015613c9d573d6000803e3d6000fd5b507f611b7c0d84fda988026215bef9b3e4d81cbceced7e679be6d5e044b588467c0e33613cc8612f46565b604080516001600160a01b03938416815291909216602082015280820185905290519081900360600190a15050565b613d003361438d565b613d3f576040805162461bcd60e51b815260206004820152601a60248201526000805160206153b4833981519152604482015290519081900360640190fd5b611c3360098263ffffffff614fab16565b6000613d5a614fff565b6001600160a01b0316836001600160a01b03161415613d7a5750806114ef565b816001600160a01b03841615613e3f576000806000613d9887614cf2565b5050509350935093505080613dea576040805162461bcd60e51b8152602060048201526013602482015272746f6b656e206e6f7420617661696c61626c6560681b604482015290519081900360640190fd5b81613e25576040805162461bcd60e51b81526020600482015260066024820152650726174653d360d41b604482015290519081900360640190fd5b613e3983613416888563ffffffff614e8416565b93505050505b6000806000613e4c615075565b5050509350935093505080613e9e576040805162461bcd60e51b8152602060048201526013602482015272746f6b656e206e6f7420617661696c61626c6560681b604482015290519081900360640190fd5b81613ee4576040805162461bcd60e51b81526020600482015260116024820152700737461626c65636f696e20726174653d3607c1b604482015290519081900360640190fd5b613ef882613416868663ffffffff614e8416565b979650505050505050565b613f0c3361438d565b613f4b576040805162461bcd60e51b815260206004820152601a60248201526000805160206153b4833981519152604482015290519081900360640190fd5b611b8d60168263ffffffff614fab16565b613f653361438d565b613fa4576040805162461bcd60e51b815260206004820152601a60248201526000805160206153b4833981519152604482015290519081900360640190fd5b61128360108263ffffffff614fab16565b613fbe33614310565b80613fc857503330145b61400c576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b8181808060200260200160405190810160405280939291908181526020018383602002808284376000920182905250925050505b815181101561411b57614058828281518110612d0d57fe5b156140a3576040805162461bcd60e51b8152602060048201526016602482015275636f6e7461696e73206f776e6572206164647265737360501b604482015290519081900360640190fd5b60006001600160a01b03168282815181106140ba57fe5b60200260200101516001600160a01b03161415614113576040805162461bcd60e51b8152602060048201526012602482015271636f6e7461696e732030206164647265737360701b604482015290519081900360640190fd5b600101614040565b5060085462010000900460ff1615614172576040805162461bcd60e51b81526020600482015260156024820152741dda1a5d195b1a5cdd081a5b9a5d1a585b1a5e9959605a1b604482015290519081900360640190fd5b60005b82811015614263576004600085858481811061418d57fe5b602090810292909201356001600160a01b03168352508101919091526040016000205460ff1661425b576001600460008686858181106141c957fe5b905060200201356001600160a01b03166001600160a01b03166001600160a01b0316815260200190815260200160002060006101000a81548160ff021916908315150217905550600584848381811061421e57fe5b835460018101855560009485526020948590200180546001600160a01b0319166001600160a01b0395909202939093013593909316929092179055505b600101614175565b506008805462ff0000191662010000179055604080513380825260208201838152600580549484018590527fb2f6cccee7a369e23e293c25aa19bef80af11eb26deba3ea0f2a02783f752e4a94929390929091906060830190849080156142f357602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116142d5575b5050935050505060405180910390a1505050565b60145460ff1690565b6003546001600160a01b0390811691161490565b600482015460ff1615614372576040805162461bcd60e51b81526020600482015260116024820152701b1a5b5a5d08185b1c9958591e481cd95d607a1b604482015290519081900360640190fd5b61437c82826151e4565b50600401805460ff19166001179055565b600061439a6002546146e0565b6001600160a01b031663b429afeb836040518263ffffffff1660e01b815260040180826001600160a01b03166001600160a01b0316815260200191505060206040518083038186803b1580156143ef57600080fd5b505afa158015614403573d6000803e3d6000fd5b505050506040513d602081101561441957600080fd5b505192915050565b60008151604114614434575060006114ef565b60208201516040830151606084015160001a7f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a082111561447a57600093505050506114ef565b8060ff16601b1415801561449257508060ff16601c14155b156144a357600093505050506114ef565b6040805160008152602080820180845289905260ff8416828401526060820186905260808201859052915160019260a0808401939192601f1981019281900390910190855afa1580156144fa573d6000803e3d6000fd5b5050604051601f190151979650505050505050565b60006001600160a01b038216156145a957816001600160a01b03166370a08231846040518263ffffffff1660e01b815260040180826001600160a01b03166001600160a01b0316815260200191505060206040518083038186803b15801561457657600080fd5b505afa15801561458a573d6000803e3d6000fd5b505050506040513d60208110156145a057600080fd5b505190506114ef565b506001600160a01b038216316114ef565b60028101546000906145d5906201518063ffffffff614c9816565b4211156145e4575080546111a9565b5060018101546111a9565b600482015460ff16614648576040805162461bcd60e51b815260206004820152601960248201527f6c696d6974206861736e2774206265656e207365742079657400000000000000604482015290519081900360640190fd5b600390910155565b60008061465c83614cf2565b509098975050505050505050565b61467382615207565b80826001015410156146bf576040805162461bcd60e51b815260206004820152601060248201526f185d985a5b18589b194f185b5bdd5b9d60821b604482015290519081900360640190fd5b60018201546146d4908263ffffffff614c3b16565b82600101819055505050565b6000805460408051630178b8bf60e01b81526004810185905290516001600160a01b0390921691630178b8bf91602480820192602092909190829003018186803b15801561472d57600080fd5b505afa158015614741573d6000803e3d6000fd5b505050506040513d602081101561475757600080fd5b505160408051631d9dabef60e11b81526004810185905290516001600160a01b0390921691633b3b57de91602480820192602092909190829003018186803b1580156143ef57600080fd5b801580614828575060408051636eb1769f60e11b81523060048201526001600160a01b03848116602483015291519185169163dd62ed3e91604480820192602092909190829003018186803b1580156147fa57600080fd5b505afa15801561480e573d6000803e3d6000fd5b505050506040513d602081101561482457600080fd5b5051155b6148635760405162461bcd60e51b81526004018080602001828103825260368152602001806154656036913960400191505060405180910390fd5b604080516001600160a01b038416602482015260448082018490528251808303909101815260649091019091526020810180516001600160e01b031663095ea7b360e01b1790526148b59084906149e4565b505050565b3b151590565b6000806148cc83614cf2565b509198975050505050505050565b6000806148e8600e546146e0565b6001600160a01b031663afc72e9385856040518363ffffffff1660e01b815260040180836001600160a01b03166001600160a01b0316815260200180602001828103825283818151815260200191508051906020019080838360005b8381101561495c578181015183820152602001614944565b50505050905090810190601f1680156149895780820380516001836020036101000a031916815260200191505b509350505050604080518083038186803b1580156149a657600080fd5b505afa1580156149ba573d6000803e3d6000fd5b505050506040513d60408110156149d057600080fd5b508051602090910151909590945092505050565b6149f6826001600160a01b03166148ba565b614a47576040805162461bcd60e51b815260206004820152601f60248201527f5361666545524332303a2063616c6c20746f206e6f6e2d636f6e747261637400604482015290519081900360640190fd5b60006060836001600160a01b0316836040518082805190602001908083835b60208310614a855780518252601f199092019160209182019101614a66565b6001836020036101000a0380198251168184511680821785525050505050509050019150506000604051808303816000865af19150503d8060008114614ae7576040519150601f19603f3d011682016040523d82523d6000602084013e614aec565b606091505b509150915081614b43576040805162461bcd60e51b815260206004820181905260248201527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564604482015290519081900360640190fd5b805115614b9c57808060200190516020811015614b5f57600080fd5b5051614b9c5760405162461bcd60e51b815260040180806020018281038252602a815260200180615418602a913960400191505060405180910390fd5b50505050565b604080517f19457468657265756d205369676e6564204d6573736167653a0a333200000000602080830191909152603c8083019490945282518083039094018452605c909101909152815191012090565b600f80546001019081905560408051338152602081019290925280517fab0423a75986556234aecd171c46ce7f5e45607d8070bf5230f2735b50322bff9281900390910190a1565b600082821115614c92576040805162461bcd60e51b815260206004820152601e60248201527f536166654d6174683a207375627472616374696f6e206f766572666c6f770000604482015290519081900360640190fd5b50900390565b600082820183811015612161576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b6060600080600080600080614d08600e546146e0565b6001600160a01b0316631f69565f896040518263ffffffff1660e01b815260040180826001600160a01b03166001600160a01b0316815260200191505060006040518083038186803b158015614d5d57600080fd5b505afa158015614d71573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405260e0811015614d9a57600080fd5b8101908080516040519392919084600160201b821115614db957600080fd5b908301906020820185811115614dce57600080fd5b8251600160201b811182820188101715614de757600080fd5b82525081516020918201929091019080838360005b83811015614e14578181015183820152602001614dfc565b50505050905090810190601f168015614e415780820380516001836020036101000a031916815260200191505b5060409081526020820151908201516060830151608084015160a085015160c090950151979e50929c50909a509850965094509192505050919395979092949650565b600082614e93575060006114ef565b82820282848281614ea057fe5b04146121615760405162461bcd60e51b81526004018080602001828103825260218152602001806153d46021913960400191505060405180910390fd5b6000808211614f33576040805162461bcd60e51b815260206004820152601a60248201527f536166654d6174683a206469766973696f6e206279207a65726f000000000000604482015290519081900360640190fd5b6000828481614f3e57fe5b04949350505050565b6001600160a01b038216614f91576040516001600160a01b0384169082156108fc029083906000818181858888f19350505050158015614f8b573d6000803e3d6000fd5b506148b5565b6148b56001600160a01b038316848363ffffffff61526016565b80826003015414614fed5760405162461bcd60e51b81526004018080602001828103825260228152602001806153926022913960400191505060405180910390fd5b614ffb8283600301546151e4565b5050565b600061500c600e546146e0565b6001600160a01b031663e9cbd8226040518163ffffffff1660e01b815260040160206040518083038186803b15801561504457600080fd5b505afa158015615058573d6000803e3d6000fd5b505050506040513d602081101561506e57600080fd5b5051905090565b606060008060008060008061508b600e546146e0565b6001600160a01b0316633efec5e96040518163ffffffff1660e01b815260040160006040518083038186803b1580156150c357600080fd5b505afa1580156150d7573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405260e081101561510057600080fd5b8101908080516040519392919084600160201b82111561511f57600080fd5b90830190602082018581111561513457600080fd5b8251600160201b81118282018810171561514d57600080fd5b82525081516020918201929091019080838360005b8381101561517a578181015183820152602001615162565b50505050905090810190601f1680156151a75780820380516001836020036101000a031916815260200191505b5060409081526020820151908201516060830151608084015160a085015160c090950151979f939e50919c509a5098509096509294509192505050565b6151ed82615207565b8082556001820154811015614ffb57815460018301555050565b600281015461521f906201518063ffffffff614c9816565b42111561525d57426002820155805460018201556040517fe93bc25276d408d390778e7a8b926f2f67209c43ed540081b951fe128f0d3cd290600090a15b50565b604080516001600160a01b038416602482015260448082018490528251808303909101815260649091019091526020810180516001600160e01b031663a9059cbb60e01b1790526148b59084906149e4565b508054600082559060005260206000209081019061525d9190615353565b828054828255906000526020600020908101928215615323579160200282015b828111156153235781546001600160a01b0319166001600160a01b038435161782556020909201916001909101906152f0565b5061532f92915061536d565b5090565b8154818355818111156148b5576000838152602090206148b59181019083015b6111b291905b8082111561532f5760008155600101615359565b6111b291905b8082111561532f5780546001600160a01b031916815560010161537356fe636f6e6669726d65642f7375626d6974746564206c696d6974206d69736d6174636873656e646572206973206e6f74206120636f6e74726f6c6c6572000000000000536166654d6174683a206d756c7469706c69636174696f6e206f766572666c6f776f776e65722063616e6e6f742062652073657420746f207a65726f20616464726573735361666545524332303a204552433230206f7065726174696f6e20646964206e6f7420737563636565646e6f6e2d6d61746368696e672070656e64696e672077686974656c69737420686173685361666545524332303a20617070726f76652066726f6d206e6f6e2d7a65726f20746f206e6f6e2d7a65726f20616c6c6f77616e6365a265627a7a72315820d21b75f8dfe861434a6413379a74c79eea3f7663c1173157de37be9ac3435e6764736f6c634300050f0032"
=======
var WalletBin = "0x60806040527f7f2ce995617d2816b426c5c8698c5ec2952f7a34bb10f38326f74933d58936976002553480156200003557600080fd5b5060405162005c5e38038062005c5e833981810160405260e08110156200005b57600080fd5b508051602082015160408301516060840151608085015160a086015160c090960151600180546001600160a01b038087166001600160a01b031992831617928390556000805490921692169190911790559495939492939192909190868685858482818686858015620000ce5760028190555b50600380546001600160a01b0319166001600160a01b0384161760ff60a01b1916600160a01b8315158102919091179182905560ff9104166200014857604080516001600160a01b038416815290517f808639ff9c8e4732d60b6c2330de498035416d229f27a77d259680895efec1229181900360200190a15b60408051600081526001600160a01b038416602082015281517f850b3df64837d7d518b45f5aa64d104652c3b80eb5b34a8e3d9eb666cb7cdea5929181900390910190a150506040805160a080820183528382526020808301859052428385018190526000606080860182905260809586018290526009889055600a97909755600b829055600c819055600d805460ff19908116909155600e98909855855193840186526706f05b59d3b20000808552928401839052948301819052948201849052910182905260108190556011556012919091556013819055601480549092169091559450620002479350506001600160e01b03620002be16915050565b50505050509150506000811162000295576040805162461bcd60e51b815260206004820152600d60248201526c37379039ba30b13632b1b7b4b760991b604482015290519081900360640190fd5b620002ad61271082026001600160e01b036200044016565b5050601b5550620005889350505050565b6060600080600080600080620002dc600e546200048e60201b60201c565b6001600160a01b0316633efec5e96040518163ffffffff1660e01b815260040160006040518083038186803b1580156200031557600080fd5b505afa1580156200032a573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405260e08110156200035457600080fd5b81019080805160405193929190846401000000008211156200037557600080fd5b9083019060208201858111156200038b57600080fd5b8251640100000000811182820188101715620003a657600080fd5b82525081516020918201929091019080838360005b83811015620003d5578181015183820152602001620003bb565b50505050905090810190601f168015620004035780820380516001836020036101000a031916815260200191505b5060409081526020820151908201516060830151608084015160a085015160c090950151979f939e50919c509a5098509096509294509192505050565b60158190556040805160a081018252828152602081018390524291810182905260006060820181905260809091018190526016839055601792909255601855601955601a805460ff19169055565b6000805460408051630178b8bf60e01b81526004810185905290516001600160a01b0390921691630178b8bf91602480820192602092909190829003018186803b158015620004dc57600080fd5b505afa158015620004f1573d6000803e3d6000fd5b505050506040513d60208110156200050857600080fd5b505160408051631d9dabef60e11b81526004810185905290516001600160a01b0390921691633b3b57de91602480820192602092909190829003018186803b1580156200055457600080fd5b505afa15801562000569573d6000803e3d6000fd5b505050506040513d60208110156200058057600080fd5b505192915050565b6156c680620005986000396000f3fe6080604052600436106103905760003560e01c80637fd004fa116101dc578063cc0e7e5611610102578063e2b4ce97116100a0578063f40b51f81161006f578063f40b51f8146110be578063f41c4319146110e8578063f421764814611112578063f776f5181461118d576103d1565b8063e2b4ce971461101c578063e61c51ca14611031578063eadd3cea1461105b578063f36febda14611085576103d1565b8063ce0b5bd5116100dc578063ce0b5bd514610f9e578063d251fefc14610fc8578063da84b1ed14610ff2578063de212bf314611007576103d1565b8063cc0e7e5614610ec6578063cccdc55614610edb578063cd7958dd14610ef0576103d1565b8063b221f3161161017a578063be40ba7911610149578063be40ba7914610e2f578063beabacc814610e44578063c4856cd914610e87578063cbd2ac6814610e9c576103d1565b8063b221f31614610d7c578063b242e53414610da6578063b87e21ef14610de1578063bcb8b74a14610e1a576103d1565b806390e690c7116101b657806390e690c714610c8c5780639b0dfd2714610ca1578063aaf1fc6214610cb6578063ab20599314610d67576103d1565b80637fd004fa14610be7578063877337b014610c625780638da5cb5b14610c77576103d1565b806332531c3c116102c15780635adc02ab1161025f57806374624c551161022e57806374624c5514610b62578063747c31d614610b8c5780637d73b23114610ba15780637d7d004614610bd2576103d1565b80635adc02ab14610a935780635d2362a814610abd5780636137d67014610ad2578063715018a614610b4d576103d1565b80633c672eb71161029b5780633c672eb7146108bc5780633f579f42146108e657806346efe0ed146109ac57806347b55a9d14610a7e576103d1565b806332531c3c146108335780633a43199f146108665780633bfec25414610892576103d1565b80631efd02991161032e57806321ce918d1161030857806321ce918d1461077a5780632587a6a2146107a457806326d05ab2146107b9578063294f4025146107ce576103d1565b80631efd02991461068557806320c13b0b1461069a5780632121dc7514610765576103d1565b8063100f23fd1161036a578063100f23fd146104715780631127b57e1461049b5780631626ba7e146105255780631aa21fba146105fa576103d1565b806301ffc9a7146103d6578063027ef3eb1461041e5780630f3a85d814610445576103d1565b366103d1576040805133815234602082015281517f88a5966d370b9919b20f3e2c13ff65706f196a4e32cc2c12bf57088f88525874929181900390910190a1005b600080fd5b3480156103e257600080fd5b5061040a600480360360208110156103f957600080fd5b50356001600160e01b0319166111a2565b604080519115158252519081900360200190f35b34801561042a57600080fd5b506104336111bc565b60408051918252519081900360200190f35b34801561045157600080fd5b5061046f6004803603602081101561046857600080fd5b50356111c3565b005b34801561047d57600080fd5b5061046f6004803603602081101561049457600080fd5b50356112cf565b3480156104a757600080fd5b506104b0611474565b6040805160208082528351818301528351919283929083019185019080838360005b838110156104ea5781810151838201526020016104d2565b50505050905090810190601f1680156105175780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561053157600080fd5b506105dd6004803603604081101561054857600080fd5b81359190810190604081016020820135600160201b81111561056957600080fd5b82018360208201111561057b57600080fd5b803590602001918460018302840111600160201b8311171561059c57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550611495945050505050565b604080516001600160e01b03199092168252519081900360200190f35b34801561060657600080fd5b5061046f6004803603604081101561061d57600080fd5b6001600160a01b038235169190810190604081016020820135600160201b81111561064757600080fd5b82018360208201111561065957600080fd5b803590602001918460208302840111600160201b8311171561067a57600080fd5b509092509050611503565b34801561069157600080fd5b50610433611689565b3480156106a657600080fd5b506105dd600480360360408110156106bd57600080fd5b810190602081018135600160201b8111156106d757600080fd5b8201836020820111156106e957600080fd5b803590602001918460018302840111600160201b8311171561070a57600080fd5b919390929091602081019035600160201b81111561072757600080fd5b82018360208201111561073957600080fd5b803590602001918460018302840111600160201b8311171561075a57600080fd5b50909250905061169a565b34801561077157600080fd5b5061040a61176f565b34801561078657600080fd5b5061046f6004803603602081101561079d57600080fd5b503561177f565b3480156107b057600080fd5b5061043361181d565b3480156107c557600080fd5b5061040a611823565b3480156107da57600080fd5b506107e361182c565b60408051602080825283518183015283519192839290830191858101910280838360005b8381101561081f578181015183820152602001610807565b505050509050019250505060405180910390f35b34801561083f57600080fd5b5061040a6004803603602081101561085657600080fd5b50356001600160a01b031661188e565b61046f6004803603604081101561087c57600080fd5b506001600160a01b0381351690602001356118a3565b34801561089e57600080fd5b5061046f600480360360208110156108b557600080fd5b5035611ae1565b3480156108c857600080fd5b5061046f600480360360208110156108df57600080fd5b5035611bd9565b3480156108f257600080fd5b506104b06004803603606081101561090957600080fd5b6001600160a01b0382351691602081013591810190606081016040820135600160201b81111561093857600080fd5b82018360208201111561094a57600080fd5b803590602001918460018302840111600160201b8311171561096b57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550611c7f945050505050565b3480156109b857600080fd5b5061046f600480360360608110156109cf57600080fd5b81359190810190604081016020820135600160201b8111156109f057600080fd5b820183602082011115610a0257600080fd5b803590602001918460018302840111600160201b83111715610a2357600080fd5b919390929091602081019035600160201b811115610a4057600080fd5b820183602082011115610a5257600080fd5b803590602001918460018302840111600160201b83111715610a7357600080fd5b509092509050612176565b348015610a8a57600080fd5b506107e3612487565b348015610a9f57600080fd5b5061046f60048036036020811015610ab657600080fd5b50356124e7565b348015610ac957600080fd5b506104336127b7565b348015610ade57600080fd5b5061046f60048036036020811015610af557600080fd5b810190602081018135600160201b811115610b0f57600080fd5b820183602082011115610b2157600080fd5b803590602001918460208302840111600160201b83111715610b4257600080fd5b5090925090506127c3565b348015610b5957600080fd5b5061046f6129e9565b348015610b6e57600080fd5b5061046f60048036036020811015610b8557600080fd5b5035612ae7565b348015610b9857600080fd5b50610433612beb565b348015610bad57600080fd5b50610bb6612bf1565b604080516001600160a01b039092168252519081900360200190f35b348015610bde57600080fd5b50610433612c00565b348015610bf357600080fd5b5061046f60048036036020811015610c0a57600080fd5b810190602081018135600160201b811115610c2457600080fd5b820183602082011115610c3657600080fd5b803590602001918460208302840111600160201b83111715610c5757600080fd5b509092509050612c0c565b348015610c6e57600080fd5b50610433612f4e565b348015610c8357600080fd5b50610bb6612f54565b348015610c9857600080fd5b5061046f612f63565b348015610cad57600080fd5b50610433612fc0565b348015610cc257600080fd5b5061046f60048036036020811015610cd957600080fd5b810190602081018135600160201b811115610cf357600080fd5b820183602082011115610d0557600080fd5b803590602001918460018302840111600160201b83111715610d2657600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550612fc6945050505050565b348015610d7357600080fd5b5061040a613103565b348015610d8857600080fd5b5061046f60048036036020811015610d9f57600080fd5b503561310c565b348015610db257600080fd5b5061046f60048036036040811015610dc957600080fd5b506001600160a01b03813516906020013515156131fc565b348015610ded57600080fd5b5061043360048036036040811015610e0457600080fd5b506001600160a01b0381351690602001356133b6565b348015610e2657600080fd5b5061040a613446565b348015610e3b57600080fd5b5061040a61344f565b348015610e5057600080fd5b5061046f60048036036060811015610e6757600080fd5b506001600160a01b0381358116916020810135909116906040013561345e565b348015610e9357600080fd5b506104336135e8565b348015610ea857600080fd5b5061046f60048036036020811015610ebf57600080fd5b50356135ee565b348015610ed257600080fd5b50610433613984565b348015610ee757600080fd5b5061043361398a565b348015610efc57600080fd5b5061043360048036036020811015610f1357600080fd5b810190602081018135600160201b811115610f2d57600080fd5b820183602082011115610f3f57600080fd5b803590602001918460208302840111600160201b83111715610f6057600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929550613990945050505050565b348015610faa57600080fd5b5061046f60048036036020811015610fc157600080fd5b50356139ea565b348015610fd457600080fd5b50610bb660048036036020811015610feb57600080fd5b5035613b93565b348015610ffe57600080fd5b50610433613bba565b34801561101357600080fd5b5061040a613bc0565b34801561102857600080fd5b50610433613bce565b34801561103d57600080fd5b5061046f6004803603602081101561105457600080fd5b5035613bd4565b34801561106757600080fd5b5061046f6004803603602081101561107e57600080fd5b5035613d1e565b34801561109157600080fd5b50610433600480360360408110156110a857600080fd5b506001600160a01b038135169060200135613d77565b3480156110ca57600080fd5b5061046f600480360360208110156110e157600080fd5b5035613f2a565b3480156110f457600080fd5b5061046f6004803603602081101561110b57600080fd5b5035613f83565b34801561111e57600080fd5b5061046f6004803603602081101561113557600080fd5b810190602081018135600160201b81111561114f57600080fd5b82018360208201111561116157600080fd5b803590602001918460208302840111600160201b8311171561118257600080fd5b509092509050613fdc565b34801561119957600080fd5b5061040a61432e565b6001600160e01b031981166301ffc9a760e01b145b919050565b600c545b90565b6111cc33614337565b806111d657503330145b61121a576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b8066038d7ea4c680001115801561123957506706f05b59d3b200008111155b611280576040805162461bcd60e51b815260206004820152601360248201527206f7574206f662072616e676520746f702d757606c1b604482015290519081900360640190fd5b61129160108263ffffffff61434b16565b604080513381526020810183905281517f41ff5d5ce3b7935893a4e7269ec5caae9cca5e3bf0eb4b21d2f443489667112e929181900390910190a150565b6112d833614337565b806112e757506112e7336143b4565b611331576040805162461bcd60e51b815260206004820152601660248201527537b7363c9037bbb732b93e3e31b7b73a3937b63632b960511b604482015290519081900360640190fd5b60085460ff16611380576040805162461bcd60e51b81526020600482015260156024820152743737903832b73234b7339039bab136b4b9b9b4b7b760591b604482015290519081900360640190fd5b6113e360068054806020026020016040519081016040528092919081815260200182805480156113d957602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116113bb575b5050505050613990565b81146114205760405162461bcd60e51b81526004018080602001828103825260238152602001806156386023913960400191505060405180910390fd5b61142c60066000615484565b6008805460ff19169055604080513381526020810183905281517f7794eff834d760583543e6e510e717a5e66d2c064e225f4db448343c3e66afcf929181900390910190a150565b604051806040016040528060058152602001640332e312e360dc1b81525081565b6000806114a8848463ffffffff61444816565b90506114b381614337565b6114f1576040805162461bcd60e51b815260206004820152600a60248201526937b7363c9037bbb732b960b11b604482015290519081900360640190fd5b50630b135d3f60e11b90505b92915050565b61150c33614337565b8061151657503330145b61155a576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b806115a3576040805162461bcd60e51b8152602060048201526014602482015273617373657420617272617920697320656d70747960601b604482015290519081900360640190fd5b60005b818110156116065760006115d5308585858181106115c057fe5b905060200201356001600160a01b031661462f565b90506115fd858585858181106115e757fe5b905060200201356001600160a01b03168361345e565b506001016115a6565b507fd4f62f23021706247dcffea245d104ae7ddaec7f23acf3d11d7136d5de6a69ad83838360405180846001600160a01b03166001600160a01b03168152602001806020018281038252848482818152602001925060200280828437600083820152604051601f909101601f1916909201829003965090945050505050a1505050565b600061169560166146da565b905090565b6000808585604051602001808383808284376040805191909301818103601f190182528084528151602092830120601f8b01839004830282018301909452898152929650630b135d3f60e11b9550611711945086935089915088908190840183828082843760009201919091525061149592505050565b6001600160e01b0319161461175d576040805162461bcd60e51b815260206004820152600d60248201526c1cda59c81b9bdd081d985b1a59609a1b604482015290519081900360640190fd5b506320c13b0b60e01b95945050505050565b600354600160a01b900460ff1690565b61178833614337565b8061179257503330145b6117d6576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b6117e760098263ffffffff61470f16565b6040805182815290517f4b1b970c8a0fa761e7803ed70c13d7aca71904b13df60fbe03f981da1730da919181900360200190a150565b60105490565b60085460ff1681565b6060600780548060200260200160405190810160405280929190818152602001828054801561188457602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311611866575b5050505050905090565b60046020526000908152604090205460ff1681565b6118ac33614337565b806118b657503330145b6118fa576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b61190382614770565b611949576040805162461bcd60e51b8152602060048201526012602482015271746f6b656e206e6f74206c6f616461626c6560701b604482015290519081900360640190fd5b60006119558383613d77565b905061196860168263ffffffff61478a16565b6000611975601b54614800565b90506001600160a01b03841615611a1d576119a06001600160a01b038516828563ffffffff6148c216565b806001600160a01b0316631b3c96b485856040518363ffffffff1660e01b815260040180836001600160a01b03166001600160a01b0316815260200182815260200192505050600060405180830381600087803b158015611a0057600080fd5b505af1158015611a14573d6000803e3d6000fd5b50505050611a97565b806001600160a01b0316631b3c96b48486866040518463ffffffff1660e01b815260040180836001600160a01b03166001600160a01b03168152602001828152602001925050506000604051808303818588803b158015611a7d57600080fd5b505af1158015611a91573d6000803e3d6000fd5b50505050505b604080516001600160a01b03861681526020810185905281517f5f65674bec9af81f71be68674135a0ea3f163fb91984e3893d06da9f6ea2ce8a929181900390910190a150505050565b611aea33614337565b80611af457503330145b611b38576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b601554811115611b8a576040805162461bcd60e51b81526020600482015260186024820152771bdd5d081bd9881c985b99d9481b1bd85908185b5bdd5b9d60421b604482015290519081900360640190fd5b611b9b60168263ffffffff61434b16565b604080513381526020810183905281517f0b05243483e17c3f3377aee82b7d47e5700b48288695fc08b7ecc2759afa44ef929181900390910190a150565b611be233614337565b80611bec57503330145b611c30576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b611c4160098263ffffffff61434b16565b604080513381526020810183905281517f068f112e5ec923d412be64779fe69e0fcbb6784c6617e94cccc8fd348f2e0f21929181900390910190a150565b6060611c8a33614337565b80611c9457503330145b611cd8576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b6001600160a01b03841660009081526004602052604090205460ff16611d0957611d0960098463ffffffff61478a16565b611d1b846001600160a01b03166149da565b8015611d2b5750611d2b84614a16565b15611f1257600080611d3d8685614a30565b6001600160a01b038216600090815260046020526040902054919350915060ff16611d83576000611d6e87836133b6565b9050611d8160098263ffffffff61478a16565b505b611d9c6001600160a01b0387168563ffffffff614b3a16565b604080516020808252818301909252606091602082018180368337019050509050600160f81b81601f81518110611dcf57fe5b60200101906001600160f81b031916908160001a9053507ff77753fab406ecfff96d6ff2476c64a838fa9f6d37b1bf190f8546e395e3b6138787878460405180856001600160a01b03166001600160a01b031681526020018481526020018060200180602001838103835285818151815260200191508051906020019080838360005b83811015611e6a578181015183820152602001611e52565b50505050905090810190601f168015611e975780820380516001836020036101000a031916815260200191505b50838103825284518152845160209182019186019080838360005b83811015611eca578181015183820152602001611eb2565b50505050905090810190601f168015611ef75780820380516001836020036101000a031916815260200191505b50965050505050505060405180910390a1925061216f915050565b60006060856001600160a01b031685856040518082805190602001908083835b60208310611f515780518252601f199092019160209182019101611f32565b6001836020036101000a03801982511681845116808217855250505050505090500191505060006040518083038185875af1925050503d8060008114611fb3576040519150601f19603f3d011682016040523d82523d6000602084013e611fb8565b606091505b50915091508181906120485760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561200d578181015183820152602001611ff5565b50505050905090810190601f16801561203a5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b507ff77753fab406ecfff96d6ff2476c64a838fa9f6d37b1bf190f8546e395e3b6138686868460405180856001600160a01b03166001600160a01b031681526020018481526020018060200180602001838103835285818151815260200191508051906020019080838360005b838110156120cd5781810151838201526020016120b5565b50505050905090810190601f1680156120fa5780820380516001836020036101000a031916815260200191505b50838103825284518152845160209182019186019080838360005b8381101561212d578181015183820152602001612115565b50505050905090810190601f16801561215a5780820380516001836020036101000a031916815260200191505b50965050505050505060405180910390a19150505b9392505050565b61217f336143b4565b6121be576040805162461bcd60e51b815260206004820152601a6024820152600080516020615588833981519152604482015290519081900360640190fd5b600061220f86868660405160200180806339363c1d60e11b81525060040184815260200183838082843780830192505050935050505060405160208183030381529060405280519060200120614cf8565b9050631626ba7e60e01b6001600160e01b0319166122638285858080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061149592505050565b6001600160e01b031916146122af576040805162461bcd60e51b815260206004820152600d60248201526c1cda59c81b9bdd081d985b1a59609a1b604482015290519081900360640190fd5b600f5486146122f1576040805162461bcd60e51b81526020600482015260096024820152687478207265706c617960b81b604482015290519081900360640190fd5b6122f9614d49565b60006060306001600160a01b03168787604051808383808284376040519201945060009350909150508083038183865af19150503d8060008114612359576040519150601f19603f3d011682016040523d82523d6000602084013e61235e565b606091505b50915091508181906123b15760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561200d578181015183820152602001611ff5565b507f823dbcf2b7b0f265871963ca65ac033f6b4c71e0d82cd123d2ff23d752dc21c187878360405180806020018060200183810383528686828181526020019250808284376000838201819052601f909101601f191690920185810384528651815286516020918201939188019250908190849084905b83811015612440578181015183820152602001612428565b50505050905090810190601f16801561246d5780820380516001836020036101000a031916815260200191505b509550505050505060405180910390a15050505050505050565b60606006805480602002602001604051908101604052809291908181526020018280548015611884576020028201919060005260206000209081546001600160a01b03168152600190910190602001808311611866575050505050905090565b6124f0336143b4565b61252f576040805162461bcd60e51b815260206004820152601a6024820152600080516020615588833981519152604482015290519081900360640190fd5b60085460ff1661257e576040805162461bcd60e51b81526020600482015260156024820152743737903832b73234b7339039bab136b4b9b9b4b7b760591b604482015290519081900360640190fd5b6125df60068054806020026020016040519081016040528092919081815260200182805480156113d9576020028201919060005260206000209081546001600160a01b031681526001909101906020018083116113bb575050505050613990565b811461261c5760405162461bcd60e51b81526004018080602001828103825260238152602001806156386023913960400191505060405180910390fd5b60005b60065481101561270357600460006006838154811061263a57fe5b60009182526020808320909101546001600160a01b0316835282019290925260400190205460ff166126fb576001600460006006848154811061267957fe5b6000918252602080832091909101546001600160a01b031683528201929092526040019020805460ff19169115159190911790556006805460059190839081106126bf57fe5b60009182526020808320909101548354600181018555938352912090910180546001600160a01b0319166001600160a01b039092169190911790555b60010161261f565b507fb2f6cccee7a369e23e293c25aa19bef80af11eb26deba3ea0f2a02783f752e4a33600660405180836001600160a01b03166001600160a01b0316815260200180602001828103825283818154815260200191508054801561278f57602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311612771575b5050935050505060405180910390a16127aa60066000615484565b506008805460ff19169055565b600061169560096146da565b6127cc33614337565b806127d657503330145b61281a576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b60085460ff161580156128355750600854610100900460ff16155b612886576040805162461bcd60e51b815260206004820152601c60248201527f77686974656c6973742073756d62697373696f6e2070656e64696e6700000000604482015290519081900360640190fd5b60085462010000900460ff166128df576040805162461bcd60e51b81526020600482015260196024820152781dda1a5d195b1a5cdd081b9bdd081a5b9a5d1a585b1a5e9959603a1b604482015290519081900360640190fd5b80612923576040805162461bcd60e51b815260206004820152600f60248201526e195b5c1d1e481dda1a5d195b1a5cdd608a1b604482015290519081900360640190fd5b61292f600783836154a2565b506008805461ff00191661010017905560408051602080840282810182019093528382527ffbc0e5ca6c7e4858daf0fdb185ef5186203e74ec9c64737e93c0aeaec596e1d1928592859261299e9285918591829185019084908082843760009201919091525061399092505050565b60405180806020018381526020018281038252858582818152602001925060200280828437600083820152604051601f909101601f1916909201829003965090945050505050a15050565b6129f233614337565b612a3c576040805162461bcd60e51b815260206004820152601660248201527539b2b73232b91034b9903737ba1030b71037bbb732b960511b604482015290519081900360640190fd5b600354600160a01b900460ff16612a9a576040805162461bcd60e51b815260206004820152601d60248201527f6f776e657273686970206973206e6f74207472616e7366657261626c65000000604482015290519081900360640190fd5b600380546001600160a01b0319169055604080516000808252602082015281517f850b3df64837d7d518b45f5aa64d104652c3b80eb5b34a8e3d9eb666cb7cdea5929181900390910190a1565b612af033614337565b80612afa57503330145b612b3e576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b8066038d7ea4c6800011158015612b5d57506706f05b59d3b200008111155b612ba4576040805162461bcd60e51b815260206004820152601360248201527206f7574206f662072616e676520746f702d757606c1b604482015290519081900360640190fd5b612bb560108263ffffffff61470f16565b6040805182815290517faf2a77cd04c3cc155588dd3bf67b310ab4fb3b1da3cf6b8d7d4d2aa1d09b794c9181900360200190a150565b601b5490565b6001546001600160a01b031690565b600061169560106146da565b612c1533614337565b80612c1f57503330145b612c63576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b60085460ff16158015612c7e5750600854610100900460ff16155b612ccf576040805162461bcd60e51b815260206004820152601c60248201527f77686974656c6973742073756d62697373696f6e2070656e64696e6700000000604482015290519081900360640190fd5b8181808060200260200160405190810160405280939291908181526020018383602002808284376000920182905250925050505b8151811015612deb57612d28828281518110612d1b57fe5b6020026020010151614337565b15612d73576040805162461bcd60e51b8152602060048201526016602482015275636f6e7461696e73206f776e6572206164647265737360501b604482015290519081900360640190fd5b60006001600160a01b0316828281518110612d8a57fe5b60200260200101516001600160a01b03161415612de3576040805162461bcd60e51b8152602060048201526012602482015271636f6e7461696e732030206164647265737360701b604482015290519081900360640190fd5b600101612d03565b5060085462010000900460ff16612e45576040805162461bcd60e51b81526020600482015260196024820152781dda1a5d195b1a5cdd081b9bdd081a5b9a5d1a585b1a5e9959603a1b604482015290519081900360640190fd5b81612e89576040805162461bcd60e51b815260206004820152600f60248201526e195b5c1d1e481dda1a5d195b1a5cdd608a1b604482015290519081900360640190fd5b612e95600684846154a2565b506008805460ff1916600117905560408051602080850282810182019093528482527f9c80b3b5f68b3e017766d59e8d09b34efe6462b05c398f35cab9e271d9bc3b9c9286928692612f029285918591829185019084908082843760009201919091525061399092505050565b60405180806020018381526020018281038252858582818152602001925060200280828437600083820152604051601f909101601f1916909201829003965090945050505050a1505050565b600e5490565b6003546001600160a01b031690565b612f6c33614337565b612fb6576040805162461bcd60e51b815260206004820152601660248201527539b2b73232b91034b9903737ba1030b71037bbb732b960511b604482015290519081900360640190fd5b612fbe614d49565b565b60095490565b612fcf33614337565b80612fd957503330145b61301d576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b8051602080820191906000808060605b868510156130f95761304686605463ffffffff614d9116565b888601805160148201516034909201805193995060609190911c9650909450909250905061308b605461307f878563ffffffff614dd316565b9063ffffffff614dd316565b9450868511156130d2576040805162461bcd60e51b815260206004820152600d60248201526c6f7574206f6620626f756e647360981b604482015290519081900360640190fd5b816130e857506040805160208101909152600081525b6130f3848483611c7f565b5061302d565b5050505050505050565b601a5460ff1690565b61311533614337565b8061311f57503330145b613163576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b6015548111156131b5576040805162461bcd60e51b81526020600482015260186024820152771bdd5d081bd9881c985b99d9481b1bd85908185b5bdd5b9d60421b604482015290519081900360640190fd5b6131c660168263ffffffff61470f16565b6040805182815290517fc178d379965e5657b6fc57494e392f121a14119215dfb422aad7db4cc03f2d109181900360200190a150565b61320533614337565b61324f576040805162461bcd60e51b815260206004820152601660248201527539b2b73232b91034b9903737ba1030b71037bbb732b960511b604482015290519081900360640190fd5b600354600160a01b900460ff166132ad576040805162461bcd60e51b815260206004820152601d60248201527f6f776e657273686970206973206e6f74207472616e7366657261626c65000000604482015290519081900360640190fd5b6001600160a01b0382166132f25760405162461bcd60e51b81526004018080602001828103825260238152602001806155eb6023913960400191505060405180910390fd5b6003805460ff60a01b1916600160a01b831515021790558061334b57604080516001600160a01b038416815290517f808639ff9c8e4732d60b6c2330de498035416d229f27a77d259680895efec1229181900360200190a15b600354604080516001600160a01b039283168152918416602083015280517f850b3df64837d7d518b45f5aa64d104652c3b80eb5b34a8e3d9eb666cb7cdea59281900390910190a150600380546001600160a01b0319166001600160a01b0392909216919091179055565b6000806000806133c586614e2d565b50505093509350935050801561343a5781613410576040805162461bcd60e51b81526020600482015260066024820152650726174653d360d41b604482015290519081900360640190fd5b61343083613424878563ffffffff614fbf16565b9063ffffffff61501816565b93505050506114fd565b50600095945050505050565b600d5460ff1690565b60085462010000900460ff1681565b61346733614337565b8061347157503330145b6134b5576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b80806134f2576040805162461bcd60e51b8152602060048201526007602482015266076616c75653d360cc1b604482015290519081900360640190fd5b6001600160a01b03841661353d576040805162461bcd60e51b815260206004820152600d60248201526c064657374696e6174696f6e3d3609c1b604482015290519081900360640190fd5b6001600160a01b03841660009081526004602052604090205460ff1661358d57816001600160a01b0384161561357a5761357784846133b6565b90505b61358b60098263ffffffff61478a16565b505b61359884848461505a565b604080516001600160a01b0380871682528516602082015280820184905290517fd1ba4ac2e2a11b5101f6cb4d978f514a155b421e8ec396d2d9abaf0bb02917ee9181900360600190a150505050565b60195490565b6135f7336143b4565b613636576040805162461bcd60e51b815260206004820152601a6024820152600080516020615588833981519152604482015290519081900360640190fd5b600854610100900460ff1661368a576040805162461bcd60e51b81526020600482015260156024820152743737903832b73234b7339039bab136b4b9b9b4b7b760591b604482015290519081900360640190fd5b6136eb60078054806020026020016040519081016040528092919081815260200182805480156113d9576020028201919060005260206000209081546001600160a01b031681526001909101906020018083116113bb575050505050613990565b81146137285760405162461bcd60e51b81526004018080602001828103825260238152602001806156386023913960400191505060405180910390fd5b60005b6007548110156138cf57600460006007838154811061374657fe5b60009182526020808320909101546001600160a01b0316835282019290925260400190205460ff16156138c7576000600460006007848154811061378657fe5b6000918252602080832091909101546001600160a01b031683528201929092526040018120805460ff1916921515929092179091555b6005546137d090600163ffffffff614d9116565b81101561389857600782815481106137e457fe5b600091825260209091200154600580546001600160a01b03909216918390811061380a57fe5b6000918252602090912001546001600160a01b031614156138905760058054600019810190811061383757fe5b600091825260209091200154600580546001600160a01b03909216918390811061385d57fe5b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b03160217905550613898565b6001016137bc565b5060058054806138a457fe5b600082815260209020810160001990810180546001600160a01b03191690550190555b60010161372b565b507fd218c430fa348f4ce67791021b6b89c0c3eacd4ead1d8f5b83c60038ec28249b33600760405180836001600160a01b03166001600160a01b0316815260200180602001828103825283818154815260200191508054801561395b57602002820191906000526020600020905b81546001600160a01b0316815260019091019060200180831161393d575b5050935050505060405180910390a161397660076000615484565b506008805461ff0019169055565b60135490565b600f5481565b60008160405160200180828051906020019060200280838360005b838110156139c35781810151838201526020016139ab565b50505050905001915050604051602081830303815290604052805190602001209050919050565b6139f333614337565b80613a025750613a02336143b4565b613a4c576040805162461bcd60e51b815260206004820152601660248201527537b7363c9037bbb732b93e3e31b7b73a3937b63632b960511b604482015290519081900360640190fd5b600854610100900460ff16613aa0576040805162461bcd60e51b81526020600482015260156024820152743737903832b73234b7339039bab136b4b9b9b4b7b760591b604482015290519081900360640190fd5b613b0160078054806020026020016040519081016040528092919081815260200182805480156113d9576020028201919060005260206000209081546001600160a01b031681526001909101906020018083116113bb575050505050613990565b8114613b3e5760405162461bcd60e51b81526004018080602001828103825260238152602001806156386023913960400191505060405180910390fd5b613b4a60076000615484565b6008805461ff0019169055604080513381526020810183905281517f13c935eb475aa0f6e931fece83e2ac44569ce2d53460d29a6dedab40b965c8a3929181900390910190a150565b60058181548110613ba057fe5b6000918252602090912001546001600160a01b0316905081565b60165490565b600854610100900460ff1681565b60025490565b8080613c11576040805162461bcd60e51b8152602060048201526007602482015266076616c75653d360cc1b604482015290519081900360640190fd5b613c1a33614337565b80613c295750613c29336143b4565b613c73576040805162461bcd60e51b815260206004820152601660248201527537b7363c9037bbb732b93e3e31b7b73a3937b63632b960511b604482015290519081900360640190fd5b613c8460108363ffffffff61478a16565b613c8c612f54565b6001600160a01b03166108fc839081150290604051600060405180830381858888f19350505050158015613cc4573d6000803e3d6000fd5b507f611b7c0d84fda988026215bef9b3e4d81cbceced7e679be6d5e044b588467c0e33613cef612f54565b604080516001600160a01b03938416815291909216602082015280820185905290519081900360600190a15050565b613d27336143b4565b613d66576040805162461bcd60e51b815260206004820152601a6024820152600080516020615588833981519152604482015290519081900360640190fd5b611c4160098263ffffffff6150be16565b6000613d81615112565b6001600160a01b0316836001600160a01b03161415613da15750806114fd565b816001600160a01b03841615613e66576000806000613dbf87614e2d565b5050509350935093505080613e11576040805162461bcd60e51b8152602060048201526013602482015272746f6b656e206e6f7420617661696c61626c6560681b604482015290519081900360640190fd5b81613e4c576040805162461bcd60e51b81526020600482015260066024820152650726174653d360d41b604482015290519081900360640190fd5b613e6083613424888563ffffffff614fbf16565b93505050505b6000806000613e73615188565b5050509350935093505080613ec5576040805162461bcd60e51b8152602060048201526013602482015272746f6b656e206e6f7420617661696c61626c6560681b604482015290519081900360640190fd5b81613f0b576040805162461bcd60e51b81526020600482015260116024820152700737461626c65636f696e20726174653d3607c1b604482015290519081900360640190fd5b613f1f82613424868663ffffffff614fbf16565b979650505050505050565b613f33336143b4565b613f72576040805162461bcd60e51b815260206004820152601a6024820152600080516020615588833981519152604482015290519081900360640190fd5b611b9b60168263ffffffff6150be16565b613f8c336143b4565b613fcb576040805162461bcd60e51b815260206004820152601a6024820152600080516020615588833981519152604482015290519081900360640190fd5b61129160108263ffffffff6150be16565b613fe533614337565b80613fef57503330145b614033576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b8181808060200260200160405190810160405280939291908181526020018383602002808284376000920182905250925050505b81518110156141425761407f828281518110612d1b57fe5b156140ca576040805162461bcd60e51b8152602060048201526016602482015275636f6e7461696e73206f776e6572206164647265737360501b604482015290519081900360640190fd5b60006001600160a01b03168282815181106140e157fe5b60200260200101516001600160a01b0316141561413a576040805162461bcd60e51b8152602060048201526012602482015271636f6e7461696e732030206164647265737360701b604482015290519081900360640190fd5b600101614067565b5060085462010000900460ff1615614199576040805162461bcd60e51b81526020600482015260156024820152741dda1a5d195b1a5cdd081a5b9a5d1a585b1a5e9959605a1b604482015290519081900360640190fd5b60005b8281101561428a57600460008585848181106141b457fe5b602090810292909201356001600160a01b03168352508101919091526040016000205460ff16614282576001600460008686858181106141f057fe5b905060200201356001600160a01b03166001600160a01b03166001600160a01b0316815260200190815260200160002060006101000a81548160ff021916908315150217905550600584848381811061424557fe5b835460018101855560009485526020948590200180546001600160a01b0319166001600160a01b0395909202939093013593909316929092179055505b60010161419c565b506008805462ff0000191662010000179055604080513380825260208201838152600580549484018590527fb2f6cccee7a369e23e293c25aa19bef80af11eb26deba3ea0f2a02783f752e4a949293909290919060608301908490801561431a57602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116142fc575b5050935050505060405180910390a1505050565b60145460ff1690565b6003546001600160a01b0390811691161490565b600482015460ff1615614399576040805162461bcd60e51b81526020600482015260116024820152701b1a5b5a5d08185b1c9958591e481cd95d607a1b604482015290519081900360640190fd5b6143a382826152f7565b50600401805460ff19166001179055565b60006143c1600254614800565b6001600160a01b031663b429afeb836040518263ffffffff1660e01b815260040180826001600160a01b03166001600160a01b0316815260200191505060206040518083038186803b15801561441657600080fd5b505afa15801561442a573d6000803e3d6000fd5b505050506040513d602081101561444057600080fd5b505192915050565b600081516041146144a0576040805162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e67746800604482015290519081900360640190fd5b60208201516040830151606084015160001a7f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08211156145115760405162461bcd60e51b81526004018080602001828103825260228152602001806155666022913960400191505060405180910390fd5b8060ff16601b1415801561452957508060ff16601c14155b156145655760405162461bcd60e51b81526004018080602001828103825260228152602001806155a86022913960400191505060405180910390fd5b60408051600080825260208083018085528a905260ff85168385015260608301879052608083018690529251909260019260a080820193601f1981019281900390910190855afa1580156145bd573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b038116614625576040805162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e61747572650000000000000000604482015290519081900360640190fd5b9695505050505050565b60006001600160a01b038216156146c957816001600160a01b03166370a08231846040518263ffffffff1660e01b815260040180826001600160a01b03166001600160a01b0316815260200191505060206040518083038186803b15801561469657600080fd5b505afa1580156146aa573d6000803e3d6000fd5b505050506040513d60208110156146c057600080fd5b505190506114fd565b506001600160a01b038216316114fd565b60028101546000906146f5906201518063ffffffff614dd316565b421115614704575080546111b7565b5060018101546111b7565b600482015460ff16614768576040805162461bcd60e51b815260206004820152601960248201527f6c696d6974206861736e2774206265656e207365742079657400000000000000604482015290519081900360640190fd5b600390910155565b60008061477c83614e2d565b509098975050505050505050565b6147938261531a565b80826001015410156147df576040805162461bcd60e51b815260206004820152601060248201526f185d985a5b18589b194f185b5bdd5b9d60821b604482015290519081900360640190fd5b60018201546147f4908263ffffffff614d9116565b82600101819055505050565b6000805460408051630178b8bf60e01b81526004810185905290516001600160a01b0390921691630178b8bf91602480820192602092909190829003018186803b15801561484d57600080fd5b505afa158015614861573d6000803e3d6000fd5b505050506040513d602081101561487757600080fd5b505160408051631d9dabef60e11b81526004810185905290516001600160a01b0390921691633b3b57de91602480820192602092909190829003018186803b15801561441657600080fd5b801580614948575060408051636eb1769f60e11b81523060048201526001600160a01b03848116602483015291519185169163dd62ed3e91604480820192602092909190829003018186803b15801561491a57600080fd5b505afa15801561492e573d6000803e3d6000fd5b505050506040513d602081101561494457600080fd5b5051155b6149835760405162461bcd60e51b815260040180806020018281038252603681526020018061565b6036913960400191505060405180910390fd5b604080516001600160a01b038416602482015260448082018490528251808303909101815260649091019091526020810180516001600160e01b031663095ea7b360e01b1790526149d5908490614b3a565b505050565b6000813f7fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470818114801590614a0e57508115155b949350505050565b600080614a2283614e2d565b509198975050505050505050565b600080614a3e600e54614800565b6001600160a01b031663afc72e9385856040518363ffffffff1660e01b815260040180836001600160a01b03166001600160a01b0316815260200180602001828103825283818151815260200191508051906020019080838360005b83811015614ab2578181015183820152602001614a9a565b50505050905090810190601f168015614adf5780820380516001836020036101000a031916815260200191505b509350505050604080518083038186803b158015614afc57600080fd5b505afa158015614b10573d6000803e3d6000fd5b505050506040513d6040811015614b2657600080fd5b508051602090910151909590945092505050565b614b4c826001600160a01b03166149da565b614b9d576040805162461bcd60e51b815260206004820152601f60248201527f5361666545524332303a2063616c6c20746f206e6f6e2d636f6e747261637400604482015290519081900360640190fd5b60006060836001600160a01b0316836040518082805190602001908083835b60208310614bdb5780518252601f199092019160209182019101614bbc565b6001836020036101000a0380198251168184511680821785525050505050509050019150506000604051808303816000865af19150503d8060008114614c3d576040519150601f19603f3d011682016040523d82523d6000602084013e614c42565b606091505b509150915081614c99576040805162461bcd60e51b815260206004820181905260248201527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564604482015290519081900360640190fd5b805115614cf257808060200190516020811015614cb557600080fd5b5051614cf25760405162461bcd60e51b815260040180806020018281038252602a81526020018061560e602a913960400191505060405180910390fd5b50505050565b604080517f19457468657265756d205369676e6564204d6573736167653a0a333200000000602080830191909152603c8083019490945282518083039094018452605c909101909152815191012090565b600f80546001019081905560408051338152602081019290925280517fab0423a75986556234aecd171c46ce7f5e45607d8070bf5230f2735b50322bff9281900390910190a1565b600061216f83836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f770000815250615373565b60008282018381101561216f576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b6060600080600080600080614e43600e54614800565b6001600160a01b0316631f69565f896040518263ffffffff1660e01b815260040180826001600160a01b03166001600160a01b0316815260200191505060006040518083038186803b158015614e9857600080fd5b505afa158015614eac573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405260e0811015614ed557600080fd5b8101908080516040519392919084600160201b821115614ef457600080fd5b908301906020820185811115614f0957600080fd5b8251600160201b811182820188101715614f2257600080fd5b82525081516020918201929091019080838360005b83811015614f4f578181015183820152602001614f37565b50505050905090810190601f168015614f7c5780820380516001836020036101000a031916815260200191505b5060409081526020820151908201516060830151608084015160a085015160c090950151979e50929c50909a509850965094509192505050919395979092949650565b600082614fce575060006114fd565b82820282848281614fdb57fe5b041461216f5760405162461bcd60e51b81526004018080602001828103825260218152602001806155ca6021913960400191505060405180910390fd5b600061216f83836040518060400160405280601a81526020017f536166654d6174683a206469766973696f6e206279207a65726f0000000000008152506153cd565b6001600160a01b0382166150a4576040516001600160a01b0384169082156108fc029083906000818181858888f1935050505015801561509e573d6000803e3d6000fd5b506149d5565b6149d56001600160a01b038316848363ffffffff61543216565b808260030154146151005760405162461bcd60e51b81526004018080602001828103825260228152602001806155446022913960400191505060405180910390fd5b61510e8283600301546152f7565b5050565b600061511f600e54614800565b6001600160a01b031663e9cbd8226040518163ffffffff1660e01b815260040160206040518083038186803b15801561515757600080fd5b505afa15801561516b573d6000803e3d6000fd5b505050506040513d602081101561518157600080fd5b5051905090565b606060008060008060008061519e600e54614800565b6001600160a01b0316633efec5e96040518163ffffffff1660e01b815260040160006040518083038186803b1580156151d657600080fd5b505afa1580156151ea573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405260e081101561521357600080fd5b8101908080516040519392919084600160201b82111561523257600080fd5b90830190602082018581111561524757600080fd5b8251600160201b81118282018810171561526057600080fd5b82525081516020918201929091019080838360005b8381101561528d578181015183820152602001615275565b50505050905090810190601f1680156152ba5780820380516001836020036101000a031916815260200191505b5060409081526020820151908201516060830151608084015160a085015160c090950151979f939e50919c509a5098509096509294509192505050565b6153008261531a565b808255600182015481101561510e57815460018301555050565b6002810154615332906201518063ffffffff614dd316565b42111561537057426002820155805460018201556040517fe93bc25276d408d390778e7a8b926f2f67209c43ed540081b951fe128f0d3cd290600090a15b50565b600081848411156153c55760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561200d578181015183820152602001611ff5565b505050900390565b6000818361541c5760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561200d578181015183820152602001611ff5565b50600083858161542857fe5b0495945050505050565b604080516001600160a01b038416602482015260448082018490528251808303909101815260649091019091526020810180516001600160e01b031663a9059cbb60e01b1790526149d5908490614b3a565b50805460008255906000526020600020908101906153709190615505565b8280548282559060005260206000209081019282156154f5579160200282015b828111156154f55781546001600160a01b0319166001600160a01b038435161782556020909201916001909101906154c2565b5061550192915061551f565b5090565b6111c091905b80821115615501576000815560010161550b565b6111c091905b808211156155015780546001600160a01b031916815560010161552556fe636f6e6669726d65642f7375626d6974746564206c696d6974206d69736d6174636845434453413a20696e76616c6964207369676e6174757265202773272076616c756573656e646572206973206e6f74206120636f6e74726f6c6c657200000000000045434453413a20696e76616c6964207369676e6174757265202776272076616c7565536166654d6174683a206d756c7469706c69636174696f6e206f766572666c6f776f776e65722063616e6e6f742062652073657420746f207a65726f20616464726573735361666545524332303a204552433230206f7065726174696f6e20646964206e6f7420737563636565646e6f6e2d6d61746368696e672070656e64696e672077686974656c69737420686173685361666545524332303a20617070726f76652066726f6d206e6f6e2d7a65726f20746f206e6f6e2d7a65726f20616c6c6f77616e6365a2646970667358221220b1d50d593928fada47c90252ba02bf12000bc7dbf918bacab891aa9c2d981c0f64736f6c63430006040033"
>>>>>>> Upgrade to solc 0.6.4 for Wallet
>>>>>>> Upgrade to solc 0.6.4 for Wallet
>>>>>>> 64ba3a32... Upgrade to solc 0.6.4 for Wallet

// DeployWallet deploys a new Ethereum contract, binding an instance of Wallet to it.
func DeployWallet(auth *bind.TransactOpts, backend bind.ContractBackend, _owner_ common.Address, _transferable_ bool, _ens_ common.Address, _tokenWhitelistNode_ [32]byte, _controllerNode_ [32]byte, _licenceNode_ [32]byte, _dailyLimit_ *big.Int) (common.Address, *types.Transaction, *Wallet, error) {
	parsed, err := abi.JSON(strings.NewReader(WalletABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(WalletBin), backend, _owner_, _transferable_, _ens_, _tokenWhitelistNode_, _controllerNode_, _licenceNode_, _dailyLimit_)
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

// WALLETVERSION is a paid mutator transaction binding the contract method 0x1127b57e.
//
// Solidity: function WALLET_VERSION() returns(string)
func (_Wallet *WalletTransactor) WALLETVERSION(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "WALLET_VERSION")
}

// WALLETVERSION is a paid mutator transaction binding the contract method 0x1127b57e.
//
// Solidity: function WALLET_VERSION() returns(string)
func (_Wallet *WalletSession) WALLETVERSION() (*types.Transaction, error) {
	return _Wallet.Contract.WALLETVERSION(&_Wallet.TransactOpts)
}

// WALLETVERSION is a paid mutator transaction binding the contract method 0x1127b57e.
//
// Solidity: function WALLET_VERSION() returns(string)
func (_Wallet *WalletTransactorSession) WALLETVERSION() (*types.Transaction, error) {
	return _Wallet.Contract.WALLETVERSION(&_Wallet.TransactOpts)
}

<<<<<<< HEAD
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

// DailyLimitAvailable is a free data retrieval call binding the contract method 0x45b12efc.
//
// Solidity: function dailyLimitAvailable() constant returns(uint256)
func (_Wallet *WalletCaller) DailyLimitAvailable(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Wallet.contract.Call(opts, out, "dailyLimitAvailable")
	return *ret0, err
}

// DailyLimitAvailable is a free data retrieval call binding the contract method 0x45b12efc.
//
// Solidity: function dailyLimitAvailable() constant returns(uint256)
func (_Wallet *WalletSession) DailyLimitAvailable() (*big.Int, error) {
	return _Wallet.Contract.DailyLimitAvailable(&_Wallet.CallOpts)
}

// DailyLimitAvailable is a free data retrieval call binding the contract method 0x45b12efc.
//
// Solidity: function dailyLimitAvailable() constant returns(uint256)
func (_Wallet *WalletCallerSession) DailyLimitAvailable() (*big.Int, error) {
	return _Wallet.Contract.DailyLimitAvailable(&_Wallet.CallOpts)
}

// DailyLimitPending is a free data retrieval call binding the contract method 0x3b8252fa.
//
// Solidity: function dailyLimitPending() constant returns(uint256)
func (_Wallet *WalletCaller) DailyLimitPending(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Wallet.contract.Call(opts, out, "dailyLimitPending")
	return *ret0, err
}

// DailyLimitPending is a free data retrieval call binding the contract method 0x3b8252fa.
//
// Solidity: function dailyLimitPending() constant returns(uint256)
func (_Wallet *WalletSession) DailyLimitPending() (*big.Int, error) {
	return _Wallet.Contract.DailyLimitPending(&_Wallet.CallOpts)
}

// DailyLimitPending is a free data retrieval call binding the contract method 0x3b8252fa.
//
// Solidity: function dailyLimitPending() constant returns(uint256)
func (_Wallet *WalletCallerSession) DailyLimitPending() (*big.Int, error) {
	return _Wallet.Contract.DailyLimitPending(&_Wallet.CallOpts)
}

// DailyLimitValue is a free data retrieval call binding the contract method 0x4d9aa248.
//
// Solidity: function dailyLimitValue() constant returns(uint256)
func (_Wallet *WalletCaller) DailyLimitValue(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Wallet.contract.Call(opts, out, "dailyLimitValue")
	return *ret0, err
}

// DailyLimitValue is a free data retrieval call binding the contract method 0x4d9aa248.
//
// Solidity: function dailyLimitValue() constant returns(uint256)
func (_Wallet *WalletSession) DailyLimitValue() (*big.Int, error) {
	return _Wallet.Contract.DailyLimitValue(&_Wallet.CallOpts)
}

// DailyLimitValue is a free data retrieval call binding the contract method 0x4d9aa248.
//
// Solidity: function dailyLimitValue() constant returns(uint256)
func (_Wallet *WalletCallerSession) DailyLimitValue() (*big.Int, error) {
	return _Wallet.Contract.DailyLimitValue(&_Wallet.CallOpts)
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
||||||| parent of 64ba3a32... Upgrade to solc 0.6.4 for Wallet
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
=======
// BatchExecuteTransaction is a paid mutator transaction binding the contract method 0xaaf1fc62.
>>>>>>> 64ba3a32... Upgrade to solc 0.6.4 for Wallet
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

// CalculateHash is a paid mutator transaction binding the contract method 0xcd7958dd.
//
// Solidity: function calculateHash(address[] _addresses) returns(bytes32)
func (_Wallet *WalletTransactor) CalculateHash(opts *bind.TransactOpts, _addresses []common.Address) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "calculateHash", _addresses)
}

// CalculateHash is a paid mutator transaction binding the contract method 0xcd7958dd.
//
// Solidity: function calculateHash(address[] _addresses) returns(bytes32)
func (_Wallet *WalletSession) CalculateHash(_addresses []common.Address) (*types.Transaction, error) {
	return _Wallet.Contract.CalculateHash(&_Wallet.TransactOpts, _addresses)
}

// CalculateHash is a paid mutator transaction binding the contract method 0xcd7958dd.
//
// Solidity: function calculateHash(address[] _addresses) returns(bytes32)
func (_Wallet *WalletTransactorSession) CalculateHash(_addresses []common.Address) (*types.Transaction, error) {
	return _Wallet.Contract.CalculateHash(&_Wallet.TransactOpts, _addresses)
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

<<<<<<< HEAD
// Monolith2FA is a free data retrieval call binding the contract method 0xcf0a866b.
||||||| parent of 64ba3a32... Upgrade to solc 0.6.4 for Wallet
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
=======
// CancelWhitelistRemoval is a paid mutator transaction binding the contract method 0xce0b5bd5.
//
// Solidity: function cancelWhitelistRemoval(bytes32 _hash) returns()
func (_Wallet *WalletTransactor) CancelWhitelistRemoval(opts *bind.TransactOpts, _hash [32]byte) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "cancelWhitelistRemoval", _hash)
}

// CancelWhitelistRemoval is a paid mutator transaction binding the contract method 0xce0b5bd5.
>>>>>>> 64ba3a32... Upgrade to solc 0.6.4 for Wallet
//
<<<<<<< HEAD
// Solidity: function monolith2FA() constant returns(bool)
func (_Wallet *WalletCaller) Monolith2FA(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Wallet.contract.Call(opts, out, "monolith2FA")
	return *ret0, err
||||||| parent of 64ba3a32... Upgrade to solc 0.6.4 for Wallet
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
=======
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
>>>>>>> 64ba3a32... Upgrade to solc 0.6.4 for Wallet
}

<<<<<<< HEAD
// Monolith2FA is a free data retrieval call binding the contract method 0xcf0a866b.
||||||| parent of 64ba3a32... Upgrade to solc 0.6.4 for Wallet
// LoadLimitValue is a free data retrieval call binding the contract method 0xda84b1ed.
=======
// ConfirmSpendLimitUpdate is a paid mutator transaction binding the contract method 0xeadd3cea.
>>>>>>> 64ba3a32... Upgrade to solc 0.6.4 for Wallet
//
<<<<<<< HEAD
// Solidity: function monolith2FA() constant returns(bool)
func (_Wallet *WalletSession) Monolith2FA() (bool, error) {
	return _Wallet.Contract.Monolith2FA(&_Wallet.CallOpts)
||||||| parent of 64ba3a32... Upgrade to solc 0.6.4 for Wallet
// Solidity: function loadLimitValue() constant returns(uint256)
func (_Wallet *WalletSession) LoadLimitValue() (*big.Int, error) {
	return _Wallet.Contract.LoadLimitValue(&_Wallet.CallOpts)
=======
// Solidity: function confirmSpendLimitUpdate(uint256 _amount) returns()
func (_Wallet *WalletSession) ConfirmSpendLimitUpdate(_amount *big.Int) (*types.Transaction, error) {
	return _Wallet.Contract.ConfirmSpendLimitUpdate(&_Wallet.TransactOpts, _amount)
>>>>>>> 64ba3a32... Upgrade to solc 0.6.4 for Wallet
}

<<<<<<< HEAD
// Monolith2FA is a free data retrieval call binding the contract method 0xcf0a866b.
||||||| parent of 64ba3a32... Upgrade to solc 0.6.4 for Wallet
// LoadLimitValue is a free data retrieval call binding the contract method 0xda84b1ed.
=======
// ConfirmSpendLimitUpdate is a paid mutator transaction binding the contract method 0xeadd3cea.
>>>>>>> 64ba3a32... Upgrade to solc 0.6.4 for Wallet
//
<<<<<<< HEAD
// Solidity: function monolith2FA() constant returns(bool)
func (_Wallet *WalletCallerSession) Monolith2FA() (bool, error) {
	return _Wallet.Contract.Monolith2FA(&_Wallet.CallOpts)
||||||| parent of 64ba3a32... Upgrade to solc 0.6.4 for Wallet
// Solidity: function loadLimitValue() constant returns(uint256)
func (_Wallet *WalletCallerSession) LoadLimitValue() (*big.Int, error) {
	return _Wallet.Contract.LoadLimitValue(&_Wallet.CallOpts)
=======
// Solidity: function confirmSpendLimitUpdate(uint256 _amount) returns()
func (_Wallet *WalletTransactorSession) ConfirmSpendLimitUpdate(_amount *big.Int) (*types.Transaction, error) {
	return _Wallet.Contract.ConfirmSpendLimitUpdate(&_Wallet.TransactOpts, _amount)
>>>>>>> 64ba3a32... Upgrade to solc 0.6.4 for Wallet
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

// ControllerNode is a paid mutator transaction binding the contract method 0xe2b4ce97.
//
// Solidity: function controllerNode() returns(bytes32)
func (_Wallet *WalletTransactor) ControllerNode(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "controllerNode")
}

// ControllerNode is a paid mutator transaction binding the contract method 0xe2b4ce97.
//
// Solidity: function controllerNode() returns(bytes32)
func (_Wallet *WalletSession) ControllerNode() (*types.Transaction, error) {
	return _Wallet.Contract.ControllerNode(&_Wallet.TransactOpts)
}

// ControllerNode is a paid mutator transaction binding the contract method 0xe2b4ce97.
//
// Solidity: function controllerNode() returns(bytes32)
func (_Wallet *WalletTransactorSession) ControllerNode() (*types.Transaction, error) {
	return _Wallet.Contract.ControllerNode(&_Wallet.TransactOpts)
}

<<<<<<< HEAD
// Personal2FA is a free data retrieval call binding the contract method 0x47d125af.
||||||| parent of 64ba3a32... Upgrade to solc 0.6.4 for Wallet
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
=======
// ConvertToEther is a paid mutator transaction binding the contract method 0xb87e21ef.
//
// Solidity: function convertToEther(address _token, uint256 _amount) returns(uint256)
func (_Wallet *WalletTransactor) ConvertToEther(opts *bind.TransactOpts, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "convertToEther", _token, _amount)
}

// ConvertToEther is a paid mutator transaction binding the contract method 0xb87e21ef.
//
// Solidity: function convertToEther(address _token, uint256 _amount) returns(uint256)
func (_Wallet *WalletSession) ConvertToEther(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Wallet.Contract.ConvertToEther(&_Wallet.TransactOpts, _token, _amount)
}

// ConvertToEther is a paid mutator transaction binding the contract method 0xb87e21ef.
//
// Solidity: function convertToEther(address _token, uint256 _amount) returns(uint256)
func (_Wallet *WalletTransactorSession) ConvertToEther(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Wallet.Contract.ConvertToEther(&_Wallet.TransactOpts, _token, _amount)
}

// ConvertToStablecoin is a paid mutator transaction binding the contract method 0xf36febda.
>>>>>>> 64ba3a32... Upgrade to solc 0.6.4 for Wallet
//
<<<<<<< HEAD
// Solidity: function personal2FA() constant returns(address)
func (_Wallet *WalletCaller) Personal2FA(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Wallet.contract.Call(opts, out, "personal2FA")
	return *ret0, err
||||||| parent of 64ba3a32... Upgrade to solc 0.6.4 for Wallet
// Solidity: function spendLimitAvailable() constant returns(uint256)
func (_Wallet *WalletCaller) SpendLimitAvailable(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Wallet.contract.Call(opts, out, "spendLimitAvailable")
	return *ret0, err
=======
// Solidity: function convertToStablecoin(address _token, uint256 _amount) returns(uint256)
func (_Wallet *WalletTransactor) ConvertToStablecoin(opts *bind.TransactOpts, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "convertToStablecoin", _token, _amount)
>>>>>>> 64ba3a32... Upgrade to solc 0.6.4 for Wallet
}

<<<<<<< HEAD
// Personal2FA is a free data retrieval call binding the contract method 0x47d125af.
||||||| parent of 64ba3a32... Upgrade to solc 0.6.4 for Wallet
// SpendLimitAvailable is a free data retrieval call binding the contract method 0x5d2362a8.
=======
// ConvertToStablecoin is a paid mutator transaction binding the contract method 0xf36febda.
>>>>>>> 64ba3a32... Upgrade to solc 0.6.4 for Wallet
//
<<<<<<< HEAD
// Solidity: function personal2FA() constant returns(address)
func (_Wallet *WalletSession) Personal2FA() (common.Address, error) {
	return _Wallet.Contract.Personal2FA(&_Wallet.CallOpts)
||||||| parent of 64ba3a32... Upgrade to solc 0.6.4 for Wallet
// Solidity: function spendLimitAvailable() constant returns(uint256)
func (_Wallet *WalletSession) SpendLimitAvailable() (*big.Int, error) {
	return _Wallet.Contract.SpendLimitAvailable(&_Wallet.CallOpts)
=======
// Solidity: function convertToStablecoin(address _token, uint256 _amount) returns(uint256)
func (_Wallet *WalletSession) ConvertToStablecoin(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Wallet.Contract.ConvertToStablecoin(&_Wallet.TransactOpts, _token, _amount)
>>>>>>> 64ba3a32... Upgrade to solc 0.6.4 for Wallet
}

<<<<<<< HEAD
// Personal2FA is a free data retrieval call binding the contract method 0x47d125af.
||||||| parent of 64ba3a32... Upgrade to solc 0.6.4 for Wallet
// SpendLimitAvailable is a free data retrieval call binding the contract method 0x5d2362a8.
=======
// ConvertToStablecoin is a paid mutator transaction binding the contract method 0xf36febda.
>>>>>>> 64ba3a32... Upgrade to solc 0.6.4 for Wallet
//
<<<<<<< HEAD
// Solidity: function personal2FA() constant returns(address)
func (_Wallet *WalletCallerSession) Personal2FA() (common.Address, error) {
	return _Wallet.Contract.Personal2FA(&_Wallet.CallOpts)
||||||| parent of 64ba3a32... Upgrade to solc 0.6.4 for Wallet
// Solidity: function spendLimitAvailable() constant returns(uint256)
func (_Wallet *WalletCallerSession) SpendLimitAvailable() (*big.Int, error) {
	return _Wallet.Contract.SpendLimitAvailable(&_Wallet.CallOpts)
=======
// Solidity: function convertToStablecoin(address _token, uint256 _amount) returns(uint256)
func (_Wallet *WalletTransactorSession) ConvertToStablecoin(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Wallet.Contract.ConvertToStablecoin(&_Wallet.TransactOpts, _token, _amount)
>>>>>>> 64ba3a32... Upgrade to solc 0.6.4 for Wallet
}

<<<<<<< HEAD
// Privileged is a free data retrieval call binding the contract method 0xced99cce.
||||||| parent of 64ba3a32... Upgrade to solc 0.6.4 for Wallet
// SpendLimitControllerConfirmationRequired is a free data retrieval call binding the contract method 0xbcb8b74a.
=======
// EnsRegistry is a paid mutator transaction binding the contract method 0x7d73b231.
>>>>>>> 64ba3a32... Upgrade to solc 0.6.4 for Wallet
//
<<<<<<< HEAD
// Solidity: function privileged() constant returns(bool)
func (_Wallet *WalletCaller) Privileged(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Wallet.contract.Call(opts, out, "privileged")
	return *ret0, err
||||||| parent of 64ba3a32... Upgrade to solc 0.6.4 for Wallet
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
=======
// Solidity: function ensRegistry() returns(address)
func (_Wallet *WalletTransactor) EnsRegistry(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "ensRegistry")
}

// EnsRegistry is a paid mutator transaction binding the contract method 0x7d73b231.
//
// Solidity: function ensRegistry() returns(address)
func (_Wallet *WalletSession) EnsRegistry() (*types.Transaction, error) {
	return _Wallet.Contract.EnsRegistry(&_Wallet.TransactOpts)
}

// EnsRegistry is a paid mutator transaction binding the contract method 0x7d73b231.
//
// Solidity: function ensRegistry() returns(address)
func (_Wallet *WalletTransactorSession) EnsRegistry() (*types.Transaction, error) {
	return _Wallet.Contract.EnsRegistry(&_Wallet.TransactOpts)
}

// ExecuteRelayedTransaction is a paid mutator transaction binding the contract method 0x46efe0ed.
//
// Solidity: function executeRelayedTransaction(uint256 _nonce, bytes _data, bytes _signature) returns()
func (_Wallet *WalletTransactor) ExecuteRelayedTransaction(opts *bind.TransactOpts, _nonce *big.Int, _data []byte, _signature []byte) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "executeRelayedTransaction", _nonce, _data, _signature)
>>>>>>> 64ba3a32... Upgrade to solc 0.6.4 for Wallet
}

<<<<<<< HEAD
// Privileged is a free data retrieval call binding the contract method 0xced99cce.
||||||| parent of 64ba3a32... Upgrade to solc 0.6.4 for Wallet
// SpendLimitPending is a free data retrieval call binding the contract method 0x027ef3eb.
=======
// ExecuteRelayedTransaction is a paid mutator transaction binding the contract method 0x46efe0ed.
>>>>>>> 64ba3a32... Upgrade to solc 0.6.4 for Wallet
//
<<<<<<< HEAD
// Solidity: function privileged() constant returns(bool)
func (_Wallet *WalletSession) Privileged() (bool, error) {
	return _Wallet.Contract.Privileged(&_Wallet.CallOpts)
||||||| parent of 64ba3a32... Upgrade to solc 0.6.4 for Wallet
// Solidity: function spendLimitPending() constant returns(uint256)
func (_Wallet *WalletSession) SpendLimitPending() (*big.Int, error) {
	return _Wallet.Contract.SpendLimitPending(&_Wallet.CallOpts)
=======
// Solidity: function executeRelayedTransaction(uint256 _nonce, bytes _data, bytes _signature) returns()
func (_Wallet *WalletSession) ExecuteRelayedTransaction(_nonce *big.Int, _data []byte, _signature []byte) (*types.Transaction, error) {
	return _Wallet.Contract.ExecuteRelayedTransaction(&_Wallet.TransactOpts, _nonce, _data, _signature)
>>>>>>> 64ba3a32... Upgrade to solc 0.6.4 for Wallet
}

<<<<<<< HEAD
// Privileged is a free data retrieval call binding the contract method 0xced99cce.
||||||| parent of 64ba3a32... Upgrade to solc 0.6.4 for Wallet
// SpendLimitPending is a free data retrieval call binding the contract method 0x027ef3eb.
=======
// ExecuteRelayedTransaction is a paid mutator transaction binding the contract method 0x46efe0ed.
>>>>>>> 64ba3a32... Upgrade to solc 0.6.4 for Wallet
//
<<<<<<< HEAD
// Solidity: function privileged() constant returns(bool)
func (_Wallet *WalletCallerSession) Privileged() (bool, error) {
	return _Wallet.Contract.Privileged(&_Wallet.CallOpts)
||||||| parent of 64ba3a32... Upgrade to solc 0.6.4 for Wallet
// Solidity: function spendLimitPending() constant returns(uint256)
func (_Wallet *WalletCallerSession) SpendLimitPending() (*big.Int, error) {
	return _Wallet.Contract.SpendLimitPending(&_Wallet.CallOpts)
=======
// Solidity: function executeRelayedTransaction(uint256 _nonce, bytes _data, bytes _signature) returns()
func (_Wallet *WalletTransactorSession) ExecuteRelayedTransaction(_nonce *big.Int, _data []byte, _signature []byte) (*types.Transaction, error) {
	return _Wallet.Contract.ExecuteRelayedTransaction(&_Wallet.TransactOpts, _nonce, _data, _signature)
>>>>>>> 64ba3a32... Upgrade to solc 0.6.4 for Wallet
}

<<<<<<< HEAD
// RelayNonce is a free data retrieval call binding the contract method 0xcccdc556.
||||||| parent of 64ba3a32... Upgrade to solc 0.6.4 for Wallet
// SpendLimitValue is a free data retrieval call binding the contract method 0x9b0dfd27.
=======
// ExecuteTransaction is a paid mutator transaction binding the contract method 0x3f579f42.
>>>>>>> 64ba3a32... Upgrade to solc 0.6.4 for Wallet
//
<<<<<<< HEAD
// Solidity: function relayNonce() constant returns(uint256)
func (_Wallet *WalletCaller) RelayNonce(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Wallet.contract.Call(opts, out, "relayNonce")
	return *ret0, err
||||||| parent of 64ba3a32... Upgrade to solc 0.6.4 for Wallet
// Solidity: function spendLimitValue() constant returns(uint256)
func (_Wallet *WalletCaller) SpendLimitValue(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Wallet.contract.Call(opts, out, "spendLimitValue")
	return *ret0, err
=======
// Solidity: function executeTransaction(address _destination, uint256 _value, bytes _data) returns(bytes)
func (_Wallet *WalletTransactor) ExecuteTransaction(opts *bind.TransactOpts, _destination common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "executeTransaction", _destination, _value, _data)
>>>>>>> 64ba3a32... Upgrade to solc 0.6.4 for Wallet
}

<<<<<<< HEAD
// RelayNonce is a free data retrieval call binding the contract method 0xcccdc556.
||||||| parent of 64ba3a32... Upgrade to solc 0.6.4 for Wallet
// SpendLimitValue is a free data retrieval call binding the contract method 0x9b0dfd27.
=======
// ExecuteTransaction is a paid mutator transaction binding the contract method 0x3f579f42.
>>>>>>> 64ba3a32... Upgrade to solc 0.6.4 for Wallet
//
<<<<<<< HEAD
// Solidity: function relayNonce() constant returns(uint256)
func (_Wallet *WalletSession) RelayNonce() (*big.Int, error) {
	return _Wallet.Contract.RelayNonce(&_Wallet.CallOpts)
||||||| parent of 64ba3a32... Upgrade to solc 0.6.4 for Wallet
// Solidity: function spendLimitValue() constant returns(uint256)
func (_Wallet *WalletSession) SpendLimitValue() (*big.Int, error) {
	return _Wallet.Contract.SpendLimitValue(&_Wallet.CallOpts)
=======
// Solidity: function executeTransaction(address _destination, uint256 _value, bytes _data) returns(bytes)
func (_Wallet *WalletSession) ExecuteTransaction(_destination common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _Wallet.Contract.ExecuteTransaction(&_Wallet.TransactOpts, _destination, _value, _data)
>>>>>>> 64ba3a32... Upgrade to solc 0.6.4 for Wallet
}

<<<<<<< HEAD
// RelayNonce is a free data retrieval call binding the contract method 0xcccdc556.
||||||| parent of 64ba3a32... Upgrade to solc 0.6.4 for Wallet
// SpendLimitValue is a free data retrieval call binding the contract method 0x9b0dfd27.
=======
// ExecuteTransaction is a paid mutator transaction binding the contract method 0x3f579f42.
>>>>>>> 64ba3a32... Upgrade to solc 0.6.4 for Wallet
//
<<<<<<< HEAD
// Solidity: function relayNonce() constant returns(uint256)
func (_Wallet *WalletCallerSession) RelayNonce() (*big.Int, error) {
	return _Wallet.Contract.RelayNonce(&_Wallet.CallOpts)
||||||| parent of 64ba3a32... Upgrade to solc 0.6.4 for Wallet
// Solidity: function spendLimitValue() constant returns(uint256)
func (_Wallet *WalletCallerSession) SpendLimitValue() (*big.Int, error) {
	return _Wallet.Contract.SpendLimitValue(&_Wallet.CallOpts)
=======
// Solidity: function executeTransaction(address _destination, uint256 _value, bytes _data) returns(bytes)
func (_Wallet *WalletTransactorSession) ExecuteTransaction(_destination common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _Wallet.Contract.ExecuteTransaction(&_Wallet.TransactOpts, _destination, _value, _data)
>>>>>>> 64ba3a32... Upgrade to solc 0.6.4 for Wallet
}

// GasTopUpLimitAvailable is a paid mutator transaction binding the contract method 0x7d7d0046.
//
// Solidity: function gasTopUpLimitAvailable() returns(uint256)
func (_Wallet *WalletTransactor) GasTopUpLimitAvailable(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "gasTopUpLimitAvailable")
}

// GasTopUpLimitAvailable is a paid mutator transaction binding the contract method 0x7d7d0046.
//
// Solidity: function gasTopUpLimitAvailable() returns(uint256)
func (_Wallet *WalletSession) GasTopUpLimitAvailable() (*types.Transaction, error) {
	return _Wallet.Contract.GasTopUpLimitAvailable(&_Wallet.TransactOpts)
}

// GasTopUpLimitAvailable is a paid mutator transaction binding the contract method 0x7d7d0046.
//
// Solidity: function gasTopUpLimitAvailable() returns(uint256)
func (_Wallet *WalletTransactorSession) GasTopUpLimitAvailable() (*types.Transaction, error) {
	return _Wallet.Contract.GasTopUpLimitAvailable(&_Wallet.TransactOpts)
}

// GasTopUpLimitControllerConfirmationRequired is a paid mutator transaction binding the contract method 0xf776f518.
//
// Solidity: function gasTopUpLimitControllerConfirmationRequired() returns(bool)
func (_Wallet *WalletTransactor) GasTopUpLimitControllerConfirmationRequired(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "gasTopUpLimitControllerConfirmationRequired")
}

// GasTopUpLimitControllerConfirmationRequired is a paid mutator transaction binding the contract method 0xf776f518.
//
// Solidity: function gasTopUpLimitControllerConfirmationRequired() returns(bool)
func (_Wallet *WalletSession) GasTopUpLimitControllerConfirmationRequired() (*types.Transaction, error) {
	return _Wallet.Contract.GasTopUpLimitControllerConfirmationRequired(&_Wallet.TransactOpts)
}

// GasTopUpLimitControllerConfirmationRequired is a paid mutator transaction binding the contract method 0xf776f518.
//
// Solidity: function gasTopUpLimitControllerConfirmationRequired() returns(bool)
func (_Wallet *WalletTransactorSession) GasTopUpLimitControllerConfirmationRequired() (*types.Transaction, error) {
	return _Wallet.Contract.GasTopUpLimitControllerConfirmationRequired(&_Wallet.TransactOpts)
}

// GasTopUpLimitPending is a paid mutator transaction binding the contract method 0xcc0e7e56.
//
// Solidity: function gasTopUpLimitPending() returns(uint256)
func (_Wallet *WalletTransactor) GasTopUpLimitPending(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "gasTopUpLimitPending")
}

// GasTopUpLimitPending is a paid mutator transaction binding the contract method 0xcc0e7e56.
//
// Solidity: function gasTopUpLimitPending() returns(uint256)
func (_Wallet *WalletSession) GasTopUpLimitPending() (*types.Transaction, error) {
	return _Wallet.Contract.GasTopUpLimitPending(&_Wallet.TransactOpts)
}

// GasTopUpLimitPending is a paid mutator transaction binding the contract method 0xcc0e7e56.
//
// Solidity: function gasTopUpLimitPending() returns(uint256)
func (_Wallet *WalletTransactorSession) GasTopUpLimitPending() (*types.Transaction, error) {
	return _Wallet.Contract.GasTopUpLimitPending(&_Wallet.TransactOpts)
}

// GasTopUpLimitValue is a paid mutator transaction binding the contract method 0x2587a6a2.
//
// Solidity: function gasTopUpLimitValue() returns(uint256)
func (_Wallet *WalletTransactor) GasTopUpLimitValue(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "gasTopUpLimitValue")
}

// GasTopUpLimitValue is a paid mutator transaction binding the contract method 0x2587a6a2.
//
// Solidity: function gasTopUpLimitValue() returns(uint256)
func (_Wallet *WalletSession) GasTopUpLimitValue() (*types.Transaction, error) {
	return _Wallet.Contract.GasTopUpLimitValue(&_Wallet.TransactOpts)
}

// GasTopUpLimitValue is a paid mutator transaction binding the contract method 0x2587a6a2.
//
// Solidity: function gasTopUpLimitValue() returns(uint256)
func (_Wallet *WalletTransactorSession) GasTopUpLimitValue() (*types.Transaction, error) {
	return _Wallet.Contract.GasTopUpLimitValue(&_Wallet.TransactOpts)
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

// IsSetWhitelist is a paid mutator transaction binding the contract method 0xbe40ba79.
//
// Solidity: function isSetWhitelist() returns(bool)
func (_Wallet *WalletTransactor) IsSetWhitelist(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "isSetWhitelist")
}

// IsSetWhitelist is a paid mutator transaction binding the contract method 0xbe40ba79.
//
// Solidity: function isSetWhitelist() returns(bool)
func (_Wallet *WalletSession) IsSetWhitelist() (*types.Transaction, error) {
	return _Wallet.Contract.IsSetWhitelist(&_Wallet.TransactOpts)
}

// IsSetWhitelist is a paid mutator transaction binding the contract method 0xbe40ba79.
//
// Solidity: function isSetWhitelist() returns(bool)
func (_Wallet *WalletTransactorSession) IsSetWhitelist() (*types.Transaction, error) {
	return _Wallet.Contract.IsSetWhitelist(&_Wallet.TransactOpts)
}

<<<<<<< HEAD
// BulkTransfer is a paid mutator transaction binding the contract method 0x1aa21fba.
||||||| parent of 64ba3a32... Upgrade to solc 0.6.4 for Wallet
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
=======
// IsTransferable is a paid mutator transaction binding the contract method 0x2121dc75.
//
// Solidity: function isTransferable() returns(bool)
func (_Wallet *WalletTransactor) IsTransferable(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "isTransferable")
}

// IsTransferable is a paid mutator transaction binding the contract method 0x2121dc75.
//
// Solidity: function isTransferable() returns(bool)
func (_Wallet *WalletSession) IsTransferable() (*types.Transaction, error) {
	return _Wallet.Contract.IsTransferable(&_Wallet.TransactOpts)
}

// IsTransferable is a paid mutator transaction binding the contract method 0x2121dc75.
//
// Solidity: function isTransferable() returns(bool)
func (_Wallet *WalletTransactorSession) IsTransferable() (*types.Transaction, error) {
	return _Wallet.Contract.IsTransferable(&_Wallet.TransactOpts)
}

// IsValidSignature is a paid mutator transaction binding the contract method 0x1626ba7e.
>>>>>>> 64ba3a32... Upgrade to solc 0.6.4 for Wallet
//
// Solidity: function isValidSignature(bytes32 _hashedData, bytes _signature) returns(bytes4)
func (_Wallet *WalletTransactor) IsValidSignature(opts *bind.TransactOpts, _hashedData [32]byte, _signature []byte) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "isValidSignature", _hashedData, _signature)
}

// IsValidSignature is a paid mutator transaction binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 _hashedData, bytes _signature) returns(bytes4)
func (_Wallet *WalletSession) IsValidSignature(_hashedData [32]byte, _signature []byte) (*types.Transaction, error) {
	return _Wallet.Contract.IsValidSignature(&_Wallet.TransactOpts, _hashedData, _signature)
}

// IsValidSignature is a paid mutator transaction binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 _hashedData, bytes _signature) returns(bytes4)
func (_Wallet *WalletTransactorSession) IsValidSignature(_hashedData [32]byte, _signature []byte) (*types.Transaction, error) {
	return _Wallet.Contract.IsValidSignature(&_Wallet.TransactOpts, _hashedData, _signature)
}

// IsValidSignature0 is a paid mutator transaction binding the contract method 0x20c13b0b.
//
// Solidity: function isValidSignature(bytes _data, bytes _signature) returns(bytes4)
func (_Wallet *WalletTransactor) IsValidSignature0(opts *bind.TransactOpts, _data []byte, _signature []byte) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "isValidSignature0", _data, _signature)
}

// IsValidSignature0 is a paid mutator transaction binding the contract method 0x20c13b0b.
//
// Solidity: function isValidSignature(bytes _data, bytes _signature) returns(bytes4)
func (_Wallet *WalletSession) IsValidSignature0(_data []byte, _signature []byte) (*types.Transaction, error) {
	return _Wallet.Contract.IsValidSignature0(&_Wallet.TransactOpts, _data, _signature)
}

// IsValidSignature0 is a paid mutator transaction binding the contract method 0x20c13b0b.
//
// Solidity: function isValidSignature(bytes _data, bytes _signature) returns(bytes4)
func (_Wallet *WalletTransactorSession) IsValidSignature0(_data []byte, _signature []byte) (*types.Transaction, error) {
	return _Wallet.Contract.IsValidSignature0(&_Wallet.TransactOpts, _data, _signature)
}

// LicenceNode is a paid mutator transaction binding the contract method 0x747c31d6.
//
// Solidity: function licenceNode() returns(bytes32)
func (_Wallet *WalletTransactor) LicenceNode(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "licenceNode")
}

// LicenceNode is a paid mutator transaction binding the contract method 0x747c31d6.
//
// Solidity: function licenceNode() returns(bytes32)
func (_Wallet *WalletSession) LicenceNode() (*types.Transaction, error) {
	return _Wallet.Contract.LicenceNode(&_Wallet.TransactOpts)
}

// LicenceNode is a paid mutator transaction binding the contract method 0x747c31d6.
//
// Solidity: function licenceNode() returns(bytes32)
func (_Wallet *WalletTransactorSession) LicenceNode() (*types.Transaction, error) {
	return _Wallet.Contract.LicenceNode(&_Wallet.TransactOpts)
}

<<<<<<< HEAD
// ConfirmDailyLimitUpdate is a paid mutator transaction binding the contract method 0x6c37a7e6.
||||||| parent of 64ba3a32... Upgrade to solc 0.6.4 for Wallet
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
=======
// LoadLimitAvailable is a paid mutator transaction binding the contract method 0x1efd0299.
//
// Solidity: function loadLimitAvailable() returns(uint256)
func (_Wallet *WalletTransactor) LoadLimitAvailable(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "loadLimitAvailable")
}

// LoadLimitAvailable is a paid mutator transaction binding the contract method 0x1efd0299.
//
// Solidity: function loadLimitAvailable() returns(uint256)
func (_Wallet *WalletSession) LoadLimitAvailable() (*types.Transaction, error) {
	return _Wallet.Contract.LoadLimitAvailable(&_Wallet.TransactOpts)
}

// LoadLimitAvailable is a paid mutator transaction binding the contract method 0x1efd0299.
//
// Solidity: function loadLimitAvailable() returns(uint256)
func (_Wallet *WalletTransactorSession) LoadLimitAvailable() (*types.Transaction, error) {
	return _Wallet.Contract.LoadLimitAvailable(&_Wallet.TransactOpts)
}

// LoadLimitControllerConfirmationRequired is a paid mutator transaction binding the contract method 0xab205993.
//
// Solidity: function loadLimitControllerConfirmationRequired() returns(bool)
func (_Wallet *WalletTransactor) LoadLimitControllerConfirmationRequired(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "loadLimitControllerConfirmationRequired")
}

// LoadLimitControllerConfirmationRequired is a paid mutator transaction binding the contract method 0xab205993.
//
// Solidity: function loadLimitControllerConfirmationRequired() returns(bool)
func (_Wallet *WalletSession) LoadLimitControllerConfirmationRequired() (*types.Transaction, error) {
	return _Wallet.Contract.LoadLimitControllerConfirmationRequired(&_Wallet.TransactOpts)
}

// LoadLimitControllerConfirmationRequired is a paid mutator transaction binding the contract method 0xab205993.
>>>>>>> 64ba3a32... Upgrade to solc 0.6.4 for Wallet
//
<<<<<<< HEAD
// Solidity: function confirmDailyLimitUpdate(uint256 _amount) returns()
func (_Wallet *WalletTransactor) ConfirmDailyLimitUpdate(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "confirmDailyLimitUpdate", _amount)
||||||| parent of 64ba3a32... Upgrade to solc 0.6.4 for Wallet
// Solidity: function confirmLoadLimitUpdate(uint256 _amount) returns()
func (_Wallet *WalletTransactorSession) ConfirmLoadLimitUpdate(_amount *big.Int) (*types.Transaction, error) {
	return _Wallet.Contract.ConfirmLoadLimitUpdate(&_Wallet.TransactOpts, _amount)
=======
// Solidity: function loadLimitControllerConfirmationRequired() returns(bool)
func (_Wallet *WalletTransactorSession) LoadLimitControllerConfirmationRequired() (*types.Transaction, error) {
	return _Wallet.Contract.LoadLimitControllerConfirmationRequired(&_Wallet.TransactOpts)
>>>>>>> 64ba3a32... Upgrade to solc 0.6.4 for Wallet
}

<<<<<<< HEAD
// ConfirmDailyLimitUpdate is a paid mutator transaction binding the contract method 0x6c37a7e6.
||||||| parent of 64ba3a32... Upgrade to solc 0.6.4 for Wallet
// ConfirmSpendLimitUpdate is a paid mutator transaction binding the contract method 0xeadd3cea.
=======
// LoadLimitPending is a paid mutator transaction binding the contract method 0xc4856cd9.
>>>>>>> 64ba3a32... Upgrade to solc 0.6.4 for Wallet
//
<<<<<<< HEAD
// Solidity: function confirmDailyLimitUpdate(uint256 _amount) returns()
func (_Wallet *WalletSession) ConfirmDailyLimitUpdate(_amount *big.Int) (*types.Transaction, error) {
	return _Wallet.Contract.ConfirmDailyLimitUpdate(&_Wallet.TransactOpts, _amount)
||||||| parent of 64ba3a32... Upgrade to solc 0.6.4 for Wallet
// Solidity: function confirmSpendLimitUpdate(uint256 _amount) returns()
func (_Wallet *WalletTransactor) ConfirmSpendLimitUpdate(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "confirmSpendLimitUpdate", _amount)
=======
// Solidity: function loadLimitPending() returns(uint256)
func (_Wallet *WalletTransactor) LoadLimitPending(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "loadLimitPending")
>>>>>>> 64ba3a32... Upgrade to solc 0.6.4 for Wallet
}

<<<<<<< HEAD
// ConfirmDailyLimitUpdate is a paid mutator transaction binding the contract method 0x6c37a7e6.
||||||| parent of 64ba3a32... Upgrade to solc 0.6.4 for Wallet
// ConfirmSpendLimitUpdate is a paid mutator transaction binding the contract method 0xeadd3cea.
=======
// LoadLimitPending is a paid mutator transaction binding the contract method 0xc4856cd9.
>>>>>>> 64ba3a32... Upgrade to solc 0.6.4 for Wallet
//
<<<<<<< HEAD
// Solidity: function confirmDailyLimitUpdate(uint256 _amount) returns()
func (_Wallet *WalletTransactorSession) ConfirmDailyLimitUpdate(_amount *big.Int) (*types.Transaction, error) {
	return _Wallet.Contract.ConfirmDailyLimitUpdate(&_Wallet.TransactOpts, _amount)
||||||| parent of 64ba3a32... Upgrade to solc 0.6.4 for Wallet
// Solidity: function confirmSpendLimitUpdate(uint256 _amount) returns()
func (_Wallet *WalletSession) ConfirmSpendLimitUpdate(_amount *big.Int) (*types.Transaction, error) {
	return _Wallet.Contract.ConfirmSpendLimitUpdate(&_Wallet.TransactOpts, _amount)
}

// ConfirmSpendLimitUpdate is a paid mutator transaction binding the contract method 0xeadd3cea.
//
// Solidity: function confirmSpendLimitUpdate(uint256 _amount) returns()
func (_Wallet *WalletTransactorSession) ConfirmSpendLimitUpdate(_amount *big.Int) (*types.Transaction, error) {
	return _Wallet.Contract.ConfirmSpendLimitUpdate(&_Wallet.TransactOpts, _amount)
=======
// Solidity: function loadLimitPending() returns(uint256)
func (_Wallet *WalletSession) LoadLimitPending() (*types.Transaction, error) {
	return _Wallet.Contract.LoadLimitPending(&_Wallet.TransactOpts)
}

// LoadLimitPending is a paid mutator transaction binding the contract method 0xc4856cd9.
//
// Solidity: function loadLimitPending() returns(uint256)
func (_Wallet *WalletTransactorSession) LoadLimitPending() (*types.Transaction, error) {
	return _Wallet.Contract.LoadLimitPending(&_Wallet.TransactOpts)
>>>>>>> 64ba3a32... Upgrade to solc 0.6.4 for Wallet
}

// LoadLimitValue is a paid mutator transaction binding the contract method 0xda84b1ed.
//
// Solidity: function loadLimitValue() returns(uint256)
func (_Wallet *WalletTransactor) LoadLimitValue(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "loadLimitValue")
}

// LoadLimitValue is a paid mutator transaction binding the contract method 0xda84b1ed.
//
// Solidity: function loadLimitValue() returns(uint256)
func (_Wallet *WalletSession) LoadLimitValue() (*types.Transaction, error) {
	return _Wallet.Contract.LoadLimitValue(&_Wallet.TransactOpts)
}

// LoadLimitValue is a paid mutator transaction binding the contract method 0xda84b1ed.
//
// Solidity: function loadLimitValue() returns(uint256)
func (_Wallet *WalletTransactorSession) LoadLimitValue() (*types.Transaction, error) {
	return _Wallet.Contract.LoadLimitValue(&_Wallet.TransactOpts)
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

<<<<<<< HEAD
// ExecutePrivilegedRelayedTransaction is a paid mutator transaction binding the contract method 0xc1e559a3.
//
// Solidity: function executePrivilegedRelayedTransaction(uint256 _nonce, bytes _data, bytes _signature) returns()
func (_Wallet *WalletTransactor) ExecutePrivilegedRelayedTransaction(opts *bind.TransactOpts, _nonce *big.Int, _data []byte, _signature []byte) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "executePrivilegedRelayedTransaction", _nonce, _data, _signature)
}

// ExecutePrivilegedRelayedTransaction is a paid mutator transaction binding the contract method 0xc1e559a3.
//
// Solidity: function executePrivilegedRelayedTransaction(uint256 _nonce, bytes _data, bytes _signature) returns()
func (_Wallet *WalletSession) ExecutePrivilegedRelayedTransaction(_nonce *big.Int, _data []byte, _signature []byte) (*types.Transaction, error) {
	return _Wallet.Contract.ExecutePrivilegedRelayedTransaction(&_Wallet.TransactOpts, _nonce, _data, _signature)
}

// ExecutePrivilegedRelayedTransaction is a paid mutator transaction binding the contract method 0xc1e559a3.
//
// Solidity: function executePrivilegedRelayedTransaction(uint256 _nonce, bytes _data, bytes _signature) returns()
func (_Wallet *WalletTransactorSession) ExecutePrivilegedRelayedTransaction(_nonce *big.Int, _data []byte, _signature []byte) (*types.Transaction, error) {
	return _Wallet.Contract.ExecutePrivilegedRelayedTransaction(&_Wallet.TransactOpts, _nonce, _data, _signature)
}

// ExecuteRelayedTransaction is a paid mutator transaction binding the contract method 0x46efe0ed.
||||||| parent of 64ba3a32... Upgrade to solc 0.6.4 for Wallet
// ExecuteRelayedTransaction is a paid mutator transaction binding the contract method 0x46efe0ed.
=======
// Owner is a paid mutator transaction binding the contract method 0x8da5cb5b.
>>>>>>> 64ba3a32... Upgrade to solc 0.6.4 for Wallet
//
// Solidity: function owner() returns(address)
func (_Wallet *WalletTransactor) Owner(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "owner")
}

// Owner is a paid mutator transaction binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() returns(address)
func (_Wallet *WalletSession) Owner() (*types.Transaction, error) {
	return _Wallet.Contract.Owner(&_Wallet.TransactOpts)
}

// Owner is a paid mutator transaction binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() returns(address)
func (_Wallet *WalletTransactorSession) Owner() (*types.Transaction, error) {
	return _Wallet.Contract.Owner(&_Wallet.TransactOpts)
}

// PendingWhitelistAddition is a paid mutator transaction binding the contract method 0x47b55a9d.
//
// Solidity: function pendingWhitelistAddition() returns(address[])
func (_Wallet *WalletTransactor) PendingWhitelistAddition(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "pendingWhitelistAddition")
}

// PendingWhitelistAddition is a paid mutator transaction binding the contract method 0x47b55a9d.
//
// Solidity: function pendingWhitelistAddition() returns(address[])
func (_Wallet *WalletSession) PendingWhitelistAddition() (*types.Transaction, error) {
	return _Wallet.Contract.PendingWhitelistAddition(&_Wallet.TransactOpts)
}

// PendingWhitelistAddition is a paid mutator transaction binding the contract method 0x47b55a9d.
//
// Solidity: function pendingWhitelistAddition() returns(address[])
func (_Wallet *WalletTransactorSession) PendingWhitelistAddition() (*types.Transaction, error) {
	return _Wallet.Contract.PendingWhitelistAddition(&_Wallet.TransactOpts)
}

// PendingWhitelistRemoval is a paid mutator transaction binding the contract method 0x294f4025.
//
// Solidity: function pendingWhitelistRemoval() returns(address[])
func (_Wallet *WalletTransactor) PendingWhitelistRemoval(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "pendingWhitelistRemoval")
}

// PendingWhitelistRemoval is a paid mutator transaction binding the contract method 0x294f4025.
//
// Solidity: function pendingWhitelistRemoval() returns(address[])
func (_Wallet *WalletSession) PendingWhitelistRemoval() (*types.Transaction, error) {
	return _Wallet.Contract.PendingWhitelistRemoval(&_Wallet.TransactOpts)
}

// PendingWhitelistRemoval is a paid mutator transaction binding the contract method 0x294f4025.
//
// Solidity: function pendingWhitelistRemoval() returns(address[])
func (_Wallet *WalletTransactorSession) PendingWhitelistRemoval() (*types.Transaction, error) {
	return _Wallet.Contract.PendingWhitelistRemoval(&_Wallet.TransactOpts)
}

// RelayNonce is a paid mutator transaction binding the contract method 0xcccdc556.
//
// Solidity: function relayNonce() returns(uint256)
func (_Wallet *WalletTransactor) RelayNonce(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "relayNonce")
}

// RelayNonce is a paid mutator transaction binding the contract method 0xcccdc556.
//
// Solidity: function relayNonce() returns(uint256)
func (_Wallet *WalletSession) RelayNonce() (*types.Transaction, error) {
	return _Wallet.Contract.RelayNonce(&_Wallet.TransactOpts)
}

// RelayNonce is a paid mutator transaction binding the contract method 0xcccdc556.
//
// Solidity: function relayNonce() returns(uint256)
func (_Wallet *WalletTransactorSession) RelayNonce() (*types.Transaction, error) {
	return _Wallet.Contract.RelayNonce(&_Wallet.TransactOpts)
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

// SetMonolith2FA is a paid mutator transaction binding the contract method 0xad95580b.
//
// Solidity: function setMonolith2FA() returns()
func (_Wallet *WalletTransactor) SetMonolith2FA(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "setMonolith2FA")
}

// SetMonolith2FA is a paid mutator transaction binding the contract method 0xad95580b.
//
// Solidity: function setMonolith2FA() returns()
func (_Wallet *WalletSession) SetMonolith2FA() (*types.Transaction, error) {
	return _Wallet.Contract.SetMonolith2FA(&_Wallet.TransactOpts)
}

// SetMonolith2FA is a paid mutator transaction binding the contract method 0xad95580b.
//
// Solidity: function setMonolith2FA() returns()
func (_Wallet *WalletTransactorSession) SetMonolith2FA() (*types.Transaction, error) {
	return _Wallet.Contract.SetMonolith2FA(&_Wallet.TransactOpts)
}

// SetPersonal2FA is a paid mutator transaction binding the contract method 0x7b580e75.
//
// Solidity: function setPersonal2FA(address _p2FA) returns()
func (_Wallet *WalletTransactor) SetPersonal2FA(opts *bind.TransactOpts, _p2FA common.Address) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "setPersonal2FA", _p2FA)
}

// SetPersonal2FA is a paid mutator transaction binding the contract method 0x7b580e75.
//
// Solidity: function setPersonal2FA(address _p2FA) returns()
func (_Wallet *WalletSession) SetPersonal2FA(_p2FA common.Address) (*types.Transaction, error) {
	return _Wallet.Contract.SetPersonal2FA(&_Wallet.TransactOpts, _p2FA)
}

// SetPersonal2FA is a paid mutator transaction binding the contract method 0x7b580e75.
//
// Solidity: function setPersonal2FA(address _p2FA) returns()
func (_Wallet *WalletTransactorSession) SetPersonal2FA(_p2FA common.Address) (*types.Transaction, error) {
	return _Wallet.Contract.SetPersonal2FA(&_Wallet.TransactOpts, _p2FA)
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

<<<<<<< HEAD
// SubmitDailyLimitUpdate is a paid mutator transaction binding the contract method 0x458d07f2.
||||||| parent of 64ba3a32... Upgrade to solc 0.6.4 for Wallet
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
=======
// SpendLimitAvailable is a paid mutator transaction binding the contract method 0x5d2362a8.
//
// Solidity: function spendLimitAvailable() returns(uint256)
func (_Wallet *WalletTransactor) SpendLimitAvailable(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "spendLimitAvailable")
}

// SpendLimitAvailable is a paid mutator transaction binding the contract method 0x5d2362a8.
//
// Solidity: function spendLimitAvailable() returns(uint256)
func (_Wallet *WalletSession) SpendLimitAvailable() (*types.Transaction, error) {
	return _Wallet.Contract.SpendLimitAvailable(&_Wallet.TransactOpts)
}

// SpendLimitAvailable is a paid mutator transaction binding the contract method 0x5d2362a8.
//
// Solidity: function spendLimitAvailable() returns(uint256)
func (_Wallet *WalletTransactorSession) SpendLimitAvailable() (*types.Transaction, error) {
	return _Wallet.Contract.SpendLimitAvailable(&_Wallet.TransactOpts)
}

// SpendLimitControllerConfirmationRequired is a paid mutator transaction binding the contract method 0xbcb8b74a.
//
// Solidity: function spendLimitControllerConfirmationRequired() returns(bool)
func (_Wallet *WalletTransactor) SpendLimitControllerConfirmationRequired(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "spendLimitControllerConfirmationRequired")
}

// SpendLimitControllerConfirmationRequired is a paid mutator transaction binding the contract method 0xbcb8b74a.
//
// Solidity: function spendLimitControllerConfirmationRequired() returns(bool)
func (_Wallet *WalletSession) SpendLimitControllerConfirmationRequired() (*types.Transaction, error) {
	return _Wallet.Contract.SpendLimitControllerConfirmationRequired(&_Wallet.TransactOpts)
}

// SpendLimitControllerConfirmationRequired is a paid mutator transaction binding the contract method 0xbcb8b74a.
//
// Solidity: function spendLimitControllerConfirmationRequired() returns(bool)
func (_Wallet *WalletTransactorSession) SpendLimitControllerConfirmationRequired() (*types.Transaction, error) {
	return _Wallet.Contract.SpendLimitControllerConfirmationRequired(&_Wallet.TransactOpts)
}

// SpendLimitPending is a paid mutator transaction binding the contract method 0x027ef3eb.
//
// Solidity: function spendLimitPending() returns(uint256)
func (_Wallet *WalletTransactor) SpendLimitPending(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "spendLimitPending")
}

// SpendLimitPending is a paid mutator transaction binding the contract method 0x027ef3eb.
//
// Solidity: function spendLimitPending() returns(uint256)
func (_Wallet *WalletSession) SpendLimitPending() (*types.Transaction, error) {
	return _Wallet.Contract.SpendLimitPending(&_Wallet.TransactOpts)
}

// SpendLimitPending is a paid mutator transaction binding the contract method 0x027ef3eb.
//
// Solidity: function spendLimitPending() returns(uint256)
func (_Wallet *WalletTransactorSession) SpendLimitPending() (*types.Transaction, error) {
	return _Wallet.Contract.SpendLimitPending(&_Wallet.TransactOpts)
}

// SpendLimitValue is a paid mutator transaction binding the contract method 0x9b0dfd27.
//
// Solidity: function spendLimitValue() returns(uint256)
func (_Wallet *WalletTransactor) SpendLimitValue(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "spendLimitValue")
}

// SpendLimitValue is a paid mutator transaction binding the contract method 0x9b0dfd27.
//
// Solidity: function spendLimitValue() returns(uint256)
func (_Wallet *WalletSession) SpendLimitValue() (*types.Transaction, error) {
	return _Wallet.Contract.SpendLimitValue(&_Wallet.TransactOpts)
}

// SpendLimitValue is a paid mutator transaction binding the contract method 0x9b0dfd27.
//
// Solidity: function spendLimitValue() returns(uint256)
func (_Wallet *WalletTransactorSession) SpendLimitValue() (*types.Transaction, error) {
	return _Wallet.Contract.SpendLimitValue(&_Wallet.TransactOpts)
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
>>>>>>> 64ba3a32... Upgrade to solc 0.6.4 for Wallet
//
// Solidity: function submitDailyLimitUpdate(uint256 _amount) returns()
func (_Wallet *WalletTransactor) SubmitDailyLimitUpdate(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "submitDailyLimitUpdate", _amount)
}

// SubmitDailyLimitUpdate is a paid mutator transaction binding the contract method 0x458d07f2.
//
// Solidity: function submitDailyLimitUpdate(uint256 _amount) returns()
func (_Wallet *WalletSession) SubmitDailyLimitUpdate(_amount *big.Int) (*types.Transaction, error) {
	return _Wallet.Contract.SubmitDailyLimitUpdate(&_Wallet.TransactOpts, _amount)
}

// SubmitDailyLimitUpdate is a paid mutator transaction binding the contract method 0x458d07f2.
//
// Solidity: function submitDailyLimitUpdate(uint256 _amount) returns()
func (_Wallet *WalletTransactorSession) SubmitDailyLimitUpdate(_amount *big.Int) (*types.Transaction, error) {
	return _Wallet.Contract.SubmitDailyLimitUpdate(&_Wallet.TransactOpts, _amount)
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

<<<<<<< HEAD
||||||| parent of 64ba3a32... Upgrade to solc 0.6.4 for Wallet
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

=======
// SubmittedWhitelistAddition is a paid mutator transaction binding the contract method 0x26d05ab2.
//
// Solidity: function submittedWhitelistAddition() returns(bool)
func (_Wallet *WalletTransactor) SubmittedWhitelistAddition(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "submittedWhitelistAddition")
}

// SubmittedWhitelistAddition is a paid mutator transaction binding the contract method 0x26d05ab2.
//
// Solidity: function submittedWhitelistAddition() returns(bool)
func (_Wallet *WalletSession) SubmittedWhitelistAddition() (*types.Transaction, error) {
	return _Wallet.Contract.SubmittedWhitelistAddition(&_Wallet.TransactOpts)
}

// SubmittedWhitelistAddition is a paid mutator transaction binding the contract method 0x26d05ab2.
//
// Solidity: function submittedWhitelistAddition() returns(bool)
func (_Wallet *WalletTransactorSession) SubmittedWhitelistAddition() (*types.Transaction, error) {
	return _Wallet.Contract.SubmittedWhitelistAddition(&_Wallet.TransactOpts)
}

// SubmittedWhitelistRemoval is a paid mutator transaction binding the contract method 0xde212bf3.
//
// Solidity: function submittedWhitelistRemoval() returns(bool)
func (_Wallet *WalletTransactor) SubmittedWhitelistRemoval(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "submittedWhitelistRemoval")
}

// SubmittedWhitelistRemoval is a paid mutator transaction binding the contract method 0xde212bf3.
//
// Solidity: function submittedWhitelistRemoval() returns(bool)
func (_Wallet *WalletSession) SubmittedWhitelistRemoval() (*types.Transaction, error) {
	return _Wallet.Contract.SubmittedWhitelistRemoval(&_Wallet.TransactOpts)
}

// SubmittedWhitelistRemoval is a paid mutator transaction binding the contract method 0xde212bf3.
//
// Solidity: function submittedWhitelistRemoval() returns(bool)
func (_Wallet *WalletTransactorSession) SubmittedWhitelistRemoval() (*types.Transaction, error) {
	return _Wallet.Contract.SubmittedWhitelistRemoval(&_Wallet.TransactOpts)
}

// SupportsInterface is a paid mutator transaction binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceID) returns(bool)
func (_Wallet *WalletTransactor) SupportsInterface(opts *bind.TransactOpts, _interfaceID [4]byte) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "supportsInterface", _interfaceID)
}

// SupportsInterface is a paid mutator transaction binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceID) returns(bool)
func (_Wallet *WalletSession) SupportsInterface(_interfaceID [4]byte) (*types.Transaction, error) {
	return _Wallet.Contract.SupportsInterface(&_Wallet.TransactOpts, _interfaceID)
}

// SupportsInterface is a paid mutator transaction binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceID) returns(bool)
func (_Wallet *WalletTransactorSession) SupportsInterface(_interfaceID [4]byte) (*types.Transaction, error) {
	return _Wallet.Contract.SupportsInterface(&_Wallet.TransactOpts, _interfaceID)
}

// TokenWhitelistNode is a paid mutator transaction binding the contract method 0x877337b0.
//
// Solidity: function tokenWhitelistNode() returns(bytes32)
func (_Wallet *WalletTransactor) TokenWhitelistNode(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "tokenWhitelistNode")
}

// TokenWhitelistNode is a paid mutator transaction binding the contract method 0x877337b0.
//
// Solidity: function tokenWhitelistNode() returns(bytes32)
func (_Wallet *WalletSession) TokenWhitelistNode() (*types.Transaction, error) {
	return _Wallet.Contract.TokenWhitelistNode(&_Wallet.TransactOpts)
}

// TokenWhitelistNode is a paid mutator transaction binding the contract method 0x877337b0.
//
// Solidity: function tokenWhitelistNode() returns(bytes32)
func (_Wallet *WalletTransactorSession) TokenWhitelistNode() (*types.Transaction, error) {
	return _Wallet.Contract.TokenWhitelistNode(&_Wallet.TransactOpts)
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

>>>>>>> 64ba3a32... Upgrade to solc 0.6.4 for Wallet
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

// WhitelistArray is a paid mutator transaction binding the contract method 0xd251fefc.
//
// Solidity: function whitelistArray(uint256 ) returns(address)
func (_Wallet *WalletTransactor) WhitelistArray(opts *bind.TransactOpts, arg0 *big.Int) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "whitelistArray", arg0)
}

// WhitelistArray is a paid mutator transaction binding the contract method 0xd251fefc.
//
// Solidity: function whitelistArray(uint256 ) returns(address)
func (_Wallet *WalletSession) WhitelistArray(arg0 *big.Int) (*types.Transaction, error) {
	return _Wallet.Contract.WhitelistArray(&_Wallet.TransactOpts, arg0)
}

// WhitelistArray is a paid mutator transaction binding the contract method 0xd251fefc.
//
// Solidity: function whitelistArray(uint256 ) returns(address)
func (_Wallet *WalletTransactorSession) WhitelistArray(arg0 *big.Int) (*types.Transaction, error) {
	return _Wallet.Contract.WhitelistArray(&_Wallet.TransactOpts, arg0)
}

// WhitelistMap is a paid mutator transaction binding the contract method 0x32531c3c.
//
// Solidity: function whitelistMap(address ) returns(bool)
func (_Wallet *WalletTransactor) WhitelistMap(opts *bind.TransactOpts, arg0 common.Address) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "whitelistMap", arg0)
}

// WhitelistMap is a paid mutator transaction binding the contract method 0x32531c3c.
//
// Solidity: function whitelistMap(address ) returns(bool)
func (_Wallet *WalletSession) WhitelistMap(arg0 common.Address) (*types.Transaction, error) {
	return _Wallet.Contract.WhitelistMap(&_Wallet.TransactOpts, arg0)
}

// WhitelistMap is a paid mutator transaction binding the contract method 0x32531c3c.
//
// Solidity: function whitelistMap(address ) returns(bool)
func (_Wallet *WalletTransactorSession) WhitelistMap(arg0 common.Address) (*types.Transaction, error) {
	return _Wallet.Contract.WhitelistMap(&_Wallet.TransactOpts, arg0)
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
	Privileged bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterExecutedRelayedTransaction is a free log retrieval operation binding the contract event 0x1e67acdbe17d73f10c3c1cec8dba9c0bca6d8fcd7f326d2a00e6520026b21585.
//
// Solidity: event ExecutedRelayedTransaction(bytes _data, bool _privileged)
func (_Wallet *WalletFilterer) FilterExecutedRelayedTransaction(opts *bind.FilterOpts) (*WalletExecutedRelayedTransactionIterator, error) {

	logs, sub, err := _Wallet.contract.FilterLogs(opts, "ExecutedRelayedTransaction")
	if err != nil {
		return nil, err
	}
	return &WalletExecutedRelayedTransactionIterator{contract: _Wallet.contract, event: "ExecutedRelayedTransaction", logs: logs, sub: sub}, nil
}

// WatchExecutedRelayedTransaction is a free log subscription operation binding the contract event 0x1e67acdbe17d73f10c3c1cec8dba9c0bca6d8fcd7f326d2a00e6520026b21585.
//
// Solidity: event ExecutedRelayedTransaction(bytes _data, bool _privileged)
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

// ParseExecutedRelayedTransaction is a log parse operation binding the contract event 0x1e67acdbe17d73f10c3c1cec8dba9c0bca6d8fcd7f326d2a00e6520026b21585.
//
// Solidity: event ExecutedRelayedTransaction(bytes _data, bool _privileged)
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

// WalletInitializedDailyLimitIterator is returned from FilterInitializedDailyLimit and is used to iterate over the raw logs and unpacked data for InitializedDailyLimit events raised by the Wallet contract.
type WalletInitializedDailyLimitIterator struct {
	Event *WalletInitializedDailyLimit // Event containing the contract specifics and raw log

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
func (it *WalletInitializedDailyLimitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletInitializedDailyLimit)
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
		it.Event = new(WalletInitializedDailyLimit)
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
func (it *WalletInitializedDailyLimitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletInitializedDailyLimitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletInitializedDailyLimit represents a InitializedDailyLimit event raised by the Wallet contract.
type WalletInitializedDailyLimit struct {
	Amount    *big.Int
	NextReset *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterInitializedDailyLimit is a free log retrieval operation binding the contract event 0x47dcdfb9e867032608a011a7b2438466389c6eba4d51bf520b402083def29caa.
//
// Solidity: event InitializedDailyLimit(uint256 _amount, uint256 _nextReset)
func (_Wallet *WalletFilterer) FilterInitializedDailyLimit(opts *bind.FilterOpts) (*WalletInitializedDailyLimitIterator, error) {

	logs, sub, err := _Wallet.contract.FilterLogs(opts, "InitializedDailyLimit")
	if err != nil {
		return nil, err
	}
	return &WalletInitializedDailyLimitIterator{contract: _Wallet.contract, event: "InitializedDailyLimit", logs: logs, sub: sub}, nil
}

// WatchInitializedDailyLimit is a free log subscription operation binding the contract event 0x47dcdfb9e867032608a011a7b2438466389c6eba4d51bf520b402083def29caa.
//
// Solidity: event InitializedDailyLimit(uint256 _amount, uint256 _nextReset)
func (_Wallet *WalletFilterer) WatchInitializedDailyLimit(opts *bind.WatchOpts, sink chan<- *WalletInitializedDailyLimit) (event.Subscription, error) {

	logs, sub, err := _Wallet.contract.WatchLogs(opts, "InitializedDailyLimit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletInitializedDailyLimit)
				if err := _Wallet.contract.UnpackLog(event, "InitializedDailyLimit", log); err != nil {
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

// ParseInitializedDailyLimit is a log parse operation binding the contract event 0x47dcdfb9e867032608a011a7b2438466389c6eba4d51bf520b402083def29caa.
//
// Solidity: event InitializedDailyLimit(uint256 _amount, uint256 _nextReset)
func (_Wallet *WalletFilterer) ParseInitializedDailyLimit(log types.Log) (*WalletInitializedDailyLimit, error) {
	event := new(WalletInitializedDailyLimit)
	if err := _Wallet.contract.UnpackLog(event, "InitializedDailyLimit", log); err != nil {
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

// WalletSetDailyLimitIterator is returned from FilterSetDailyLimit and is used to iterate over the raw logs and unpacked data for SetDailyLimit events raised by the Wallet contract.
type WalletSetDailyLimitIterator struct {
	Event *WalletSetDailyLimit // Event containing the contract specifics and raw log

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
func (it *WalletSetDailyLimitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletSetDailyLimit)
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
		it.Event = new(WalletSetDailyLimit)
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
func (it *WalletSetDailyLimitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletSetDailyLimitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletSetDailyLimit represents a SetDailyLimit event raised by the Wallet contract.
type WalletSetDailyLimit struct {
	Sender common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetDailyLimit is a free log retrieval operation binding the contract event 0x2a843f39f13315c4c1a9bc53a1a32162858f272f3b2d0c656f409431251b6768.
//
// Solidity: event SetDailyLimit(address _sender, uint256 _amount)
func (_Wallet *WalletFilterer) FilterSetDailyLimit(opts *bind.FilterOpts) (*WalletSetDailyLimitIterator, error) {

	logs, sub, err := _Wallet.contract.FilterLogs(opts, "SetDailyLimit")
	if err != nil {
		return nil, err
	}
	return &WalletSetDailyLimitIterator{contract: _Wallet.contract, event: "SetDailyLimit", logs: logs, sub: sub}, nil
}

// WatchSetDailyLimit is a free log subscription operation binding the contract event 0x2a843f39f13315c4c1a9bc53a1a32162858f272f3b2d0c656f409431251b6768.
//
// Solidity: event SetDailyLimit(address _sender, uint256 _amount)
func (_Wallet *WalletFilterer) WatchSetDailyLimit(opts *bind.WatchOpts, sink chan<- *WalletSetDailyLimit) (event.Subscription, error) {

	logs, sub, err := _Wallet.contract.WatchLogs(opts, "SetDailyLimit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletSetDailyLimit)
				if err := _Wallet.contract.UnpackLog(event, "SetDailyLimit", log); err != nil {
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

// ParseSetDailyLimit is a log parse operation binding the contract event 0x2a843f39f13315c4c1a9bc53a1a32162858f272f3b2d0c656f409431251b6768.
//
// Solidity: event SetDailyLimit(address _sender, uint256 _amount)
func (_Wallet *WalletFilterer) ParseSetDailyLimit(log types.Log) (*WalletSetDailyLimit, error) {
	event := new(WalletSetDailyLimit)
	if err := _Wallet.contract.UnpackLog(event, "SetDailyLimit", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WalletSetMonolith2FAIterator is returned from FilterSetMonolith2FA and is used to iterate over the raw logs and unpacked data for SetMonolith2FA events raised by the Wallet contract.
type WalletSetMonolith2FAIterator struct {
	Event *WalletSetMonolith2FA // Event containing the contract specifics and raw log

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
func (it *WalletSetMonolith2FAIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletSetMonolith2FA)
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
		it.Event = new(WalletSetMonolith2FA)
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
func (it *WalletSetMonolith2FAIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletSetMonolith2FAIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletSetMonolith2FA represents a SetMonolith2FA event raised by the Wallet contract.
type WalletSetMonolith2FA struct {
	Sender common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetMonolith2FA is a free log retrieval operation binding the contract event 0xea0cce48757f6b222f91e711f59a5a8ec05a3ed3c0a4328fe36cb48f31a86969.
//
// Solidity: event SetMonolith2FA(address _sender)
func (_Wallet *WalletFilterer) FilterSetMonolith2FA(opts *bind.FilterOpts) (*WalletSetMonolith2FAIterator, error) {

	logs, sub, err := _Wallet.contract.FilterLogs(opts, "SetMonolith2FA")
	if err != nil {
		return nil, err
	}
	return &WalletSetMonolith2FAIterator{contract: _Wallet.contract, event: "SetMonolith2FA", logs: logs, sub: sub}, nil
}

// WatchSetMonolith2FA is a free log subscription operation binding the contract event 0xea0cce48757f6b222f91e711f59a5a8ec05a3ed3c0a4328fe36cb48f31a86969.
//
// Solidity: event SetMonolith2FA(address _sender)
func (_Wallet *WalletFilterer) WatchSetMonolith2FA(opts *bind.WatchOpts, sink chan<- *WalletSetMonolith2FA) (event.Subscription, error) {

	logs, sub, err := _Wallet.contract.WatchLogs(opts, "SetMonolith2FA")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletSetMonolith2FA)
				if err := _Wallet.contract.UnpackLog(event, "SetMonolith2FA", log); err != nil {
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

// ParseSetMonolith2FA is a log parse operation binding the contract event 0xea0cce48757f6b222f91e711f59a5a8ec05a3ed3c0a4328fe36cb48f31a86969.
//
// Solidity: event SetMonolith2FA(address _sender)
func (_Wallet *WalletFilterer) ParseSetMonolith2FA(log types.Log) (*WalletSetMonolith2FA, error) {
	event := new(WalletSetMonolith2FA)
	if err := _Wallet.contract.UnpackLog(event, "SetMonolith2FA", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WalletSetPersonal2FAIterator is returned from FilterSetPersonal2FA and is used to iterate over the raw logs and unpacked data for SetPersonal2FA events raised by the Wallet contract.
type WalletSetPersonal2FAIterator struct {
	Event *WalletSetPersonal2FA // Event containing the contract specifics and raw log

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
func (it *WalletSetPersonal2FAIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletSetPersonal2FA)
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
		it.Event = new(WalletSetPersonal2FA)
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
func (it *WalletSetPersonal2FAIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletSetPersonal2FAIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletSetPersonal2FA represents a SetPersonal2FA event raised by the Wallet contract.
type WalletSetPersonal2FA struct {
	Sender common.Address
	P2FA   common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetPersonal2FA is a free log retrieval operation binding the contract event 0x33635a1d7938fa110d60d48b4ecbefc4afcc07e782ca013f11de948ee7949d1c.
//
// Solidity: event SetPersonal2FA(address _sender, address _p2FA)
func (_Wallet *WalletFilterer) FilterSetPersonal2FA(opts *bind.FilterOpts) (*WalletSetPersonal2FAIterator, error) {

	logs, sub, err := _Wallet.contract.FilterLogs(opts, "SetPersonal2FA")
	if err != nil {
		return nil, err
	}
	return &WalletSetPersonal2FAIterator{contract: _Wallet.contract, event: "SetPersonal2FA", logs: logs, sub: sub}, nil
}

// WatchSetPersonal2FA is a free log subscription operation binding the contract event 0x33635a1d7938fa110d60d48b4ecbefc4afcc07e782ca013f11de948ee7949d1c.
//
// Solidity: event SetPersonal2FA(address _sender, address _p2FA)
func (_Wallet *WalletFilterer) WatchSetPersonal2FA(opts *bind.WatchOpts, sink chan<- *WalletSetPersonal2FA) (event.Subscription, error) {

	logs, sub, err := _Wallet.contract.WatchLogs(opts, "SetPersonal2FA")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletSetPersonal2FA)
				if err := _Wallet.contract.UnpackLog(event, "SetPersonal2FA", log); err != nil {
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

// ParseSetPersonal2FA is a log parse operation binding the contract event 0x33635a1d7938fa110d60d48b4ecbefc4afcc07e782ca013f11de948ee7949d1c.
//
// Solidity: event SetPersonal2FA(address _sender, address _p2FA)
func (_Wallet *WalletFilterer) ParseSetPersonal2FA(log types.Log) (*WalletSetPersonal2FA, error) {
	event := new(WalletSetPersonal2FA)
	if err := _Wallet.contract.UnpackLog(event, "SetPersonal2FA", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WalletSubmittedDailyLimitUpdateIterator is returned from FilterSubmittedDailyLimitUpdate and is used to iterate over the raw logs and unpacked data for SubmittedDailyLimitUpdate events raised by the Wallet contract.
type WalletSubmittedDailyLimitUpdateIterator struct {
	Event *WalletSubmittedDailyLimitUpdate // Event containing the contract specifics and raw log

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
func (it *WalletSubmittedDailyLimitUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletSubmittedDailyLimitUpdate)
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
		it.Event = new(WalletSubmittedDailyLimitUpdate)
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
func (it *WalletSubmittedDailyLimitUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletSubmittedDailyLimitUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletSubmittedDailyLimitUpdate represents a SubmittedDailyLimitUpdate event raised by the Wallet contract.
type WalletSubmittedDailyLimitUpdate struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSubmittedDailyLimitUpdate is a free log retrieval operation binding the contract event 0x065b9ade648867cf901516060dd7a78fad8ab1aec5eb80ee57acbb30badf86ec.
//
// Solidity: event SubmittedDailyLimitUpdate(uint256 _amount)
func (_Wallet *WalletFilterer) FilterSubmittedDailyLimitUpdate(opts *bind.FilterOpts) (*WalletSubmittedDailyLimitUpdateIterator, error) {

	logs, sub, err := _Wallet.contract.FilterLogs(opts, "SubmittedDailyLimitUpdate")
	if err != nil {
		return nil, err
	}
	return &WalletSubmittedDailyLimitUpdateIterator{contract: _Wallet.contract, event: "SubmittedDailyLimitUpdate", logs: logs, sub: sub}, nil
}

// WatchSubmittedDailyLimitUpdate is a free log subscription operation binding the contract event 0x065b9ade648867cf901516060dd7a78fad8ab1aec5eb80ee57acbb30badf86ec.
//
// Solidity: event SubmittedDailyLimitUpdate(uint256 _amount)
func (_Wallet *WalletFilterer) WatchSubmittedDailyLimitUpdate(opts *bind.WatchOpts, sink chan<- *WalletSubmittedDailyLimitUpdate) (event.Subscription, error) {

	logs, sub, err := _Wallet.contract.WatchLogs(opts, "SubmittedDailyLimitUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletSubmittedDailyLimitUpdate)
				if err := _Wallet.contract.UnpackLog(event, "SubmittedDailyLimitUpdate", log); err != nil {
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

// ParseSubmittedDailyLimitUpdate is a log parse operation binding the contract event 0x065b9ade648867cf901516060dd7a78fad8ab1aec5eb80ee57acbb30badf86ec.
//
// Solidity: event SubmittedDailyLimitUpdate(uint256 _amount)
func (_Wallet *WalletFilterer) ParseSubmittedDailyLimitUpdate(log types.Log) (*WalletSubmittedDailyLimitUpdate, error) {
	event := new(WalletSubmittedDailyLimitUpdate)
	if err := _Wallet.contract.UnpackLog(event, "SubmittedDailyLimitUpdate", log); err != nil {
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

// WalletUpdatedAvailableDailyLimitIterator is returned from FilterUpdatedAvailableDailyLimit and is used to iterate over the raw logs and unpacked data for UpdatedAvailableDailyLimit events raised by the Wallet contract.
type WalletUpdatedAvailableDailyLimitIterator struct {
	Event *WalletUpdatedAvailableDailyLimit // Event containing the contract specifics and raw log

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
func (it *WalletUpdatedAvailableDailyLimitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletUpdatedAvailableDailyLimit)
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
		it.Event = new(WalletUpdatedAvailableDailyLimit)
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
func (it *WalletUpdatedAvailableDailyLimitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletUpdatedAvailableDailyLimitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletUpdatedAvailableDailyLimit represents a UpdatedAvailableDailyLimit event raised by the Wallet contract.
type WalletUpdatedAvailableDailyLimit struct {
	Amount    *big.Int
	NextReset *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUpdatedAvailableDailyLimit is a free log retrieval operation binding the contract event 0xb8d7171194501073e2d8151eeccf1398143c5df9acfb9868d0539a256164f6ca.
//
// Solidity: event UpdatedAvailableDailyLimit(uint256 _amount, uint256 _nextReset)
func (_Wallet *WalletFilterer) FilterUpdatedAvailableDailyLimit(opts *bind.FilterOpts) (*WalletUpdatedAvailableDailyLimitIterator, error) {

	logs, sub, err := _Wallet.contract.FilterLogs(opts, "UpdatedAvailableDailyLimit")
	if err != nil {
		return nil, err
	}
	return &WalletUpdatedAvailableDailyLimitIterator{contract: _Wallet.contract, event: "UpdatedAvailableDailyLimit", logs: logs, sub: sub}, nil
}

// WatchUpdatedAvailableDailyLimit is a free log subscription operation binding the contract event 0xb8d7171194501073e2d8151eeccf1398143c5df9acfb9868d0539a256164f6ca.
//
// Solidity: event UpdatedAvailableDailyLimit(uint256 _amount, uint256 _nextReset)
func (_Wallet *WalletFilterer) WatchUpdatedAvailableDailyLimit(opts *bind.WatchOpts, sink chan<- *WalletUpdatedAvailableDailyLimit) (event.Subscription, error) {

	logs, sub, err := _Wallet.contract.WatchLogs(opts, "UpdatedAvailableDailyLimit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletUpdatedAvailableDailyLimit)
				if err := _Wallet.contract.UnpackLog(event, "UpdatedAvailableDailyLimit", log); err != nil {
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

// ParseUpdatedAvailableDailyLimit is a log parse operation binding the contract event 0xb8d7171194501073e2d8151eeccf1398143c5df9acfb9868d0539a256164f6ca.
//
// Solidity: event UpdatedAvailableDailyLimit(uint256 _amount, uint256 _nextReset)
func (_Wallet *WalletFilterer) ParseUpdatedAvailableDailyLimit(log types.Log) (*WalletUpdatedAvailableDailyLimit, error) {
	event := new(WalletUpdatedAvailableDailyLimit)
	if err := _Wallet.contract.UnpackLog(event, "UpdatedAvailableDailyLimit", log); err != nil {
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
