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
const WalletABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"_addresses\",\"type\":\"address[]\"}],\"name\":\"AddedToWhitelist\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"_assets\",\"type\":\"address[]\"}],\"name\":\"BulkTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"CancelledWhitelistAddition\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"CancelledWhitelistRemoval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_returndata\",\"type\":\"bytes\"}],\"name\":\"ExecutedRelayedTransaction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_destination\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_returndata\",\"type\":\"bytes\"}],\"name\":\"ExecutedTransaction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_currentNonce\",\"type\":\"uint256\"}],\"name\":\"IncreasedRelayNonce\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"LoadedTokenCard\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_locked\",\"type\":\"address\"}],\"name\":\"LockedOwnership\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"_addresses\",\"type\":\"address[]\"}],\"name\":\"RemovedFromWhitelist\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"SetGasTopUpLimit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"SetLoadLimit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"SetSpendLimit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"SubmittedGasTopUpLimitUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"SubmittedLoadLimitUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"SubmittedSpendLimitUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"_addresses\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"SubmittedWhitelistAddition\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"_addresses\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"SubmittedWhitelistRemoval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"ToppedUpGas\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"Transferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"TransferredOwnership\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"UpdatedAvailableLimit\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"WALLET_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_transactionBatch\",\"type\":\"bytes\"}],\"name\":\"batchExecuteTransaction\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"_assets\",\"type\":\"address[]\"}],\"name\":\"bulkTransfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_addresses\",\"type\":\"address[]\"}],\"name\":\"calculateHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"cancelWhitelistAddition\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"cancelWhitelistRemoval\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"confirmGasTopUpLimitUpdate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"confirmLoadLimitUpdate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"confirmSpendLimitUpdate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"confirmWhitelistAddition\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"confirmWhitelistRemoval\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"controllerNode\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"convertToEther\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"convertToStablecoin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ensRegistry\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"executeRelayedTransaction\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_destination\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"executeTransaction\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"gasTopUpLimitAvailable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"gasTopUpLimitControllerConfirmationRequired\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"gasTopUpLimitPending\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"gasTopUpLimitValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"}],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"increaseRelayNonce\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_owner_\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_transferable_\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"_ens_\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_tokenWhitelistNode_\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_controllerNode_\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_licenceNode_\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_spendLimit_\",\"type\":\"uint256\"}],\"name\":\"initializeWallet\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isSetWhitelist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isTransferable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hashedData\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"isValidSignature\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"isValidSignature\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"licenceNode\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"loadLimitAvailable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"loadLimitControllerConfirmationRequired\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"loadLimitPending\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"loadLimitValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"loadTokenCard\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingWhitelistAddition\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingWhitelistRemoval\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"relayNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"setGasTopUpLimit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"setLoadLimit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"setSpendLimit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_addresses\",\"type\":\"address[]\"}],\"name\":\"setWhitelist\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"spendLimitAvailable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"spendLimitControllerConfirmationRequired\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"spendLimitPending\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"spendLimitValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"submitGasTopUpLimitUpdate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"submitLoadLimitUpdate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"submitSpendLimitUpdate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_addresses\",\"type\":\"address[]\"}],\"name\":\"submitWhitelistAddition\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_addresses\",\"type\":\"address[]\"}],\"name\":\"submitWhitelistRemoval\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"submittedWhitelistAddition\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"submittedWhitelistRemoval\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_interfaceID\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tokenWhitelistNode\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"topUpGas\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_transferable\",\"type\":\"bool\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"whitelistArray\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"whitelistMap\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// WalletBin is the compiled bytecode used for deploying new contracts.
var WalletBin = "0x60806040527f7f2ce995617d2816b426c5c8698c5ec2952f7a34bb10f38326f74933d589369760345534801561003457600080fd5b50615e1a80620000456000396000f3fe6080604052600436106103a25760003560e01c80637fd004fa116101e7578063cc0e7e561161010d578063e61c51ca116100a0578063f41c43191161006f578063f41c431914611114578063f42176481461113e578063f776f518146111b9578063f8b2cb4f146111ce576103a2565b8063e61c51ca1461105d578063eadd3cea14611087578063f36febda146110b1578063f40b51f8146110ea576103a2565b8063d251fefc116100dc578063d251fefc14610ff4578063da84b1ed1461101e578063de212bf314611033578063e2b4ce9714611048576103a2565b8063cc0e7e5614610ef2578063cccdc55614610f07578063cd7958dd14610f1c578063ce0b5bd514610fca576103a2565b8063b221f31611610185578063be40ba7911610154578063be40ba7914610e5b578063beabacc814610e70578063c4856cd914610eb3578063cbd2ac6814610ec8576103a2565b8063b221f31614610da8578063b242e53414610dd2578063b87e21ef14610e0d578063bcb8b74a14610e46576103a2565b806390e690c7116101c157806390e690c714610cb85780639b0dfd2714610ccd578063aaf1fc6214610ce2578063ab20599314610d93576103a2565b80637fd004fa14610c13578063877337b014610c8e5780638da5cb5b14610ca3576103a2565b80633a43199f116102cc5780635d2362a81161026a57806374624c551161023957806374624c5514610b8e578063747c31d614610bb85780637d73b23114610bcd5780637d7d004614610bfe576103a2565b80635d2362a814610a8e5780636137d67014610aa357806369efdfc014610b1e578063715018a614610b79576103a2565b80633f579f42116102a65780633f579f42146108b757806346efe0ed1461097d57806347b55a9d14610a4f5780635adc02ab14610a64576103a2565b80633a43199f146108375780633bfec254146108635780633c672eb71461088d576103a2565b80631efd0299116103445780632587a6a2116103135780632587a6a21461077557806326d05ab21461078a578063294f40251461079f57806332531c3c14610804576103a2565b80631efd02991461065657806320c13b0b1461066b5780632121dc751461073657806321ce918d1461074b576103a2565b8063100f23fd11610380578063100f23fd146104425780631127b57e1461046c5780631626ba7e146104f65780631aa21fba146105cb576103a2565b806301ffc9a7146103a7578063027ef3eb146103ef5780630f3a85d814610416575b600080fd5b3480156103b357600080fd5b506103db600480360360208110156103ca57600080fd5b50356001600160e01b031916611201565b604080519115158252519081900360200190f35b3480156103fb57600080fd5b5061040461121b565b60408051918252519081900360200190f35b34801561042257600080fd5b506104406004803603602081101561043957600080fd5b5035611222565b005b34801561044e57600080fd5b506104406004803603602081101561046557600080fd5b503561132e565b34801561047857600080fd5b506104816114d3565b6040805160208082528351818301528351919283929083019185019080838360005b838110156104bb5781810151838201526020016104a3565b50505050905090810190601f1680156104e85780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561050257600080fd5b506105ae6004803603604081101561051957600080fd5b81359190810190604081016020820135600160201b81111561053a57600080fd5b82018360208201111561054c57600080fd5b803590602001918460018302840111600160201b8311171561056d57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506114f4945050505050565b604080516001600160e01b03199092168252519081900360200190f35b3480156105d757600080fd5b50610440600480360360408110156105ee57600080fd5b6001600160a01b038235169190810190604081016020820135600160201b81111561061857600080fd5b82018360208201111561062a57600080fd5b803590602001918460208302840111600160201b8311171561064b57600080fd5b509092509050611569565b34801561066257600080fd5b506104046116ee565b34801561067757600080fd5b506105ae6004803603604081101561068e57600080fd5b810190602081018135600160201b8111156106a857600080fd5b8201836020820111156106ba57600080fd5b803590602001918460018302840111600160201b831117156106db57600080fd5b919390929091602081019035600160201b8111156106f857600080fd5b82018360208201111561070a57600080fd5b803590602001918460018302840111600160201b8311171561072b57600080fd5b5090925090506116ff565b34801561074257600080fd5b506103db6117d4565b34801561075757600080fd5b506104406004803603602081101561076e57600080fd5b50356117e4565b34801561078157600080fd5b50610404611882565b34801561079657600080fd5b506103db611888565b3480156107ab57600080fd5b506107b4611891565b60408051602080825283518183015283519192839290830191858101910280838360005b838110156107f05781810151838201526020016107d8565b505050509050019250505060405180910390f35b34801561081057600080fd5b506103db6004803603602081101561082757600080fd5b50356001600160a01b03166118f3565b6104406004803603604081101561084d57600080fd5b506001600160a01b038135169060200135611908565b34801561086f57600080fd5b506104406004803603602081101561088657600080fd5b5035611b46565b34801561089957600080fd5b50610440600480360360208110156108b057600080fd5b5035611c3e565b3480156108c357600080fd5b50610481600480360360608110156108da57600080fd5b6001600160a01b0382351691602081013591810190606081016040820135600160201b81111561090957600080fd5b82018360208201111561091b57600080fd5b803590602001918460018302840111600160201b8311171561093c57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550611ce4945050505050565b34801561098957600080fd5b50610440600480360360608110156109a057600080fd5b81359190810190604081016020820135600160201b8111156109c157600080fd5b8201836020820111156109d357600080fd5b803590602001918460018302840111600160201b831117156109f457600080fd5b919390929091602081019035600160201b811115610a1157600080fd5b820183602082011115610a2357600080fd5b803590602001918460018302840111600160201b83111715610a4457600080fd5b5090925090506121db565b348015610a5b57600080fd5b506107b461251c565b348015610a7057600080fd5b5061044060048036036020811015610a8757600080fd5b503561257c565b348015610a9a57600080fd5b5061040461284c565b348015610aaf57600080fd5b5061044060048036036020811015610ac657600080fd5b810190602081018135600160201b811115610ae057600080fd5b820183602082011115610af257600080fd5b803590602001918460208302840111600160201b83111715610b1357600080fd5b509092509050612858565b348015610b2a57600080fd5b50610440600480360360e0811015610b4157600080fd5b506001600160a01b03813581169160208101351515916040820135169060608101359060808101359060a08101359060c00135612a7e565b348015610b8557600080fd5b50610440612b62565b348015610b9a57600080fd5b5061044060048036036020811015610bb157600080fd5b5035612c60565b348015610bc457600080fd5b50610404612d64565b348015610bd957600080fd5b50610be2612d6a565b604080516001600160a01b039092168252519081900360200190f35b348015610c0a57600080fd5b50610404612d79565b348015610c1f57600080fd5b5061044060048036036020811015610c3657600080fd5b810190602081018135600160201b811115610c5057600080fd5b820183602082011115610c6257600080fd5b803590602001918460208302840111600160201b83111715610c8357600080fd5b509092509050612d85565b348015610c9a57600080fd5b506104046130c7565b348015610caf57600080fd5b50610be26130cd565b348015610cc457600080fd5b506104406130dc565b348015610cd957600080fd5b50610404613139565b348015610cee57600080fd5b5061044060048036036020811015610d0557600080fd5b810190602081018135600160201b811115610d1f57600080fd5b820183602082011115610d3157600080fd5b803590602001918460018302840111600160201b83111715610d5257600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955061313f945050505050565b348015610d9f57600080fd5b506103db613272565b348015610db457600080fd5b5061044060048036036020811015610dcb57600080fd5b503561327b565b348015610dde57600080fd5b5061044060048036036040811015610df557600080fd5b506001600160a01b038135169060200135151561336b565b348015610e1957600080fd5b5061040460048036036040811015610e3057600080fd5b506001600160a01b038135169060200135613525565b348015610e5257600080fd5b506103db6135b5565b348015610e6757600080fd5b506103db6135be565b348015610e7c57600080fd5b5061044060048036036060811015610e9357600080fd5b506001600160a01b038135811691602081013590911690604001356135cd565b348015610ebf57600080fd5b50610404613757565b348015610ed457600080fd5b5061044060048036036020811015610eeb57600080fd5b503561375d565b348015610efe57600080fd5b50610404613ada565b348015610f1357600080fd5b50610404613ae0565b348015610f2857600080fd5b5061040460048036036020811015610f3f57600080fd5b810190602081018135600160201b811115610f5957600080fd5b820183602082011115610f6b57600080fd5b803590602001918460208302840111600160201b83111715610f8c57600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929550613ae6945050505050565b348015610fd657600080fd5b5061044060048036036020811015610fed57600080fd5b5035613b40565b34801561100057600080fd5b50610be26004803603602081101561101757600080fd5b5035613ce9565b34801561102a57600080fd5b50610404613d10565b34801561103f57600080fd5b506103db613d16565b34801561105457600080fd5b50610404613d24565b34801561106957600080fd5b506104406004803603602081101561108057600080fd5b5035613d2a565b34801561109357600080fd5b50610440600480360360208110156110aa57600080fd5b5035613e74565b3480156110bd57600080fd5b50610404600480360360408110156110d457600080fd5b506001600160a01b038135169060200135613ecd565b3480156110f657600080fd5b506104406004803603602081101561110d57600080fd5b5035614080565b34801561112057600080fd5b506104406004803603602081101561113757600080fd5b50356140d9565b34801561114a57600080fd5b506104406004803603602081101561116157600080fd5b810190602081018135600160201b81111561117b57600080fd5b82018360208201111561118d57600080fd5b803590602001918460208302840111600160201b831117156111ae57600080fd5b509092509050614132565b3480156111c557600080fd5b506103db614484565b3480156111da57600080fd5b50610404600480360360208110156111f157600080fd5b50356001600160a01b031661448d565b6001600160e01b031981166301ffc9a760e01b145b919050565b603e545b90565b61122b33614498565b8061123557503330145b611279576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b8066038d7ea4c680001115801561129857506706f05b59d3b200008111155b6112df576040805162461bcd60e51b815260206004820152601360248201527206f7574206f662072616e676520746f702d757606c1b604482015290519081900360640190fd5b6112f060408263ffffffff6144ac16565b604080513381526020810183905281517f41ff5d5ce3b7935893a4e7269ec5caae9cca5e3bf0eb4b21d2f443489667112e929181900390910190a150565b61133733614498565b80611346575061134633614515565b611390576040805162461bcd60e51b815260206004820152601660248201527537b7363c9037bbb732b93e3e31b7b73a3937b63632b960511b604482015290519081900360640190fd5b603a5460ff166113df576040805162461bcd60e51b81526020600482015260156024820152743737903832b73234b7339039bab136b4b9b9b4b7b760591b604482015290519081900360640190fd5b611442603880548060200260200160405190810160405280929190818152602001828054801561143857602002820191906000526020600020905b81546001600160a01b0316815260019091019060200180831161141a575b5050505050613ae6565b811461147f5760405162461bcd60e51b8152600401808060200182810382526023815260200180615d8d6023913960400191505060405180910390fd5b61148b60386000615bcf565b603a805460ff19169055604080513381526020810183905281517f7794eff834d760583543e6e510e717a5e66d2c064e225f4db448343c3e66afcf929181900390910190a150565b60405180604001604052806005815260200164332e332e3160d81b81525081565b600080611507848463ffffffff6145a916565b905061151281614498565b611557576040805162461bcd60e51b8152602060048201526011602482015270696e76616c6964207369676e617475726560781b604482015290519081900360640190fd5b50630b135d3f60e11b90505b92915050565b61157233614498565b8061157c57503330145b6115c0576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b80611609576040805162461bcd60e51b8152602060048201526014602482015273617373657420617272617920697320656d70747960601b604482015290519081900360640190fd5b60005b8181101561166b57600061163a84848481811061162557fe5b905060200201356001600160a01b0316614697565b90506116628585858581811061164c57fe5b905060200201356001600160a01b0316836135cd565b5060010161160c565b507fd4f62f23021706247dcffea245d104ae7ddaec7f23acf3d11d7136d5de6a69ad83838360405180846001600160a01b03166001600160a01b03168152602001806020018281038252848482818152602001925060200280828437600083820152604051601f909101601f1916909201829003965090945050505050a1505050565b60006116fa6047614728565b905090565b6000808585604051602001808383808284376040805191909301818103601f190182528084528151602092830120601f8b01839004830282018301909452898152929650630b135d3f60e11b955061177694508693508991508890819084018382808284376000920191909152506114f492505050565b6001600160e01b031916146117c2576040805162461bcd60e51b815260206004820152600d60248201526c1cda59c81b9bdd081d985b1a59609a1b604482015290519081900360640190fd5b506320c13b0b60e01b95945050505050565b603554600160a01b900460ff1690565b6117ed33614498565b806117f757503330145b61183b576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b61184c603b8263ffffffff61475d16565b6040805182815290517f4b1b970c8a0fa761e7803ed70c13d7aca71904b13df60fbe03f981da1730da919181900360200190a150565b60405490565b603a5460ff1681565b606060398054806020026020016040519081016040528092919081815260200182805480156118e957602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116118cb575b5050505050905090565b60366020526000908152604090205460ff1681565b61191133614498565b8061191b57503330145b61195f576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b611968826147be565b6119ae576040805162461bcd60e51b8152602060048201526012602482015271746f6b656e206e6f74206c6f616461626c6560701b604482015290519081900360640190fd5b60006119ba8383613ecd565b90506119cd60478263ffffffff6147d816565b60006119da604d5461484e565b90506001600160a01b03841615611a8257611a056001600160a01b038516828563ffffffff61496f16565b806001600160a01b0316631b3c96b485856040518363ffffffff1660e01b815260040180836001600160a01b03166001600160a01b0316815260200182815260200192505050600060405180830381600087803b158015611a6557600080fd5b505af1158015611a79573d6000803e3d6000fd5b50505050611afc565b806001600160a01b0316631b3c96b48486866040518463ffffffff1660e01b815260040180836001600160a01b03166001600160a01b03168152602001828152602001925050506000604051808303818588803b158015611ae257600080fd5b505af1158015611af6573d6000803e3d6000fd5b50505050505b604080516001600160a01b03861681526020810185905281517f5f65674bec9af81f71be68674135a0ea3f163fb91984e3893d06da9f6ea2ce8a929181900390910190a150505050565b611b4f33614498565b80611b5957503330145b611b9d576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b604654811115611bef576040805162461bcd60e51b81526020600482015260186024820152771bdd5d081bd9881c985b99d9481b1bd85908185b5bdd5b9d60421b604482015290519081900360640190fd5b611c0060478263ffffffff6144ac16565b604080513381526020810183905281517f0b05243483e17c3f3377aee82b7d47e5700b48288695fc08b7ecc2759afa44ef929181900390910190a150565b611c4733614498565b80611c5157503330145b611c95576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b611ca6603b8263ffffffff6144ac16565b604080513381526020810183905281517f068f112e5ec923d412be64779fe69e0fcbb6784c6617e94cccc8fd348f2e0f21929181900390910190a150565b6060611cef33614498565b80611cf957503330145b611d3d576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b6001600160a01b03841660009081526036602052604090205460ff16611d6e57611d6e603b8463ffffffff6147d816565b611d80846001600160a01b0316614a87565b8015611d905750611d9084614a8d565b15611f7757600080611da28685614aa7565b6001600160a01b038216600090815260366020526040902054919350915060ff16611de8576000611dd38783613525565b9050611de6603b8263ffffffff6147d816565b505b611e016001600160a01b0387168563ffffffff614bb116565b604080516020808252818301909252606091602082018180388339019050509050600160f81b81601f81518110611e3457fe5b60200101906001600160f81b031916908160001a9053507ff77753fab406ecfff96d6ff2476c64a838fa9f6d37b1bf190f8546e395e3b6138787878460405180856001600160a01b03166001600160a01b031681526020018481526020018060200180602001838103835285818151815260200191508051906020019080838360005b83811015611ecf578181015183820152602001611eb7565b50505050905090810190601f168015611efc5780820380516001836020036101000a031916815260200191505b50838103825284518152845160209182019186019080838360005b83811015611f2f578181015183820152602001611f17565b50505050905090810190601f168015611f5c5780820380516001836020036101000a031916815260200191505b50965050505050505060405180910390a192506121d4915050565b60006060856001600160a01b031685856040518082805190602001908083835b60208310611fb65780518252601f199092019160209182019101611f97565b6001836020036101000a03801982511681845116808217855250505050505090500191505060006040518083038185875af1925050503d8060008114612018576040519150601f19603f3d011682016040523d82523d6000602084013e61201d565b606091505b50915091508181906120ad5760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561207257818101518382015260200161205a565b50505050905090810190601f16801561209f5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b507ff77753fab406ecfff96d6ff2476c64a838fa9f6d37b1bf190f8546e395e3b6138686868460405180856001600160a01b03166001600160a01b031681526020018481526020018060200180602001838103835285818151815260200191508051906020019080838360005b8381101561213257818101518382015260200161211a565b50505050905090810190601f16801561215f5780820380516001836020036101000a031916815260200191505b50838103825284518152845160209182019186019080838360005b8381101561219257818101518382015260200161217a565b50505050905090810190601f1680156121bf5780820380516001836020036101000a031916815260200191505b50965050505050505060405180910390a19150505b9392505050565b6121e433614515565b612223576040805162461bcd60e51b815260206004820152601a6024820152600080516020615cd1833981519152604482015290519081900360640190fd5b600046905060006122a3823089898960405160200180806836b7b737b634ba341d60b91b815250600901868152602001856001600160a01b03166001600160a01b031660601b8152601401848152602001838380828437808301925050509550505050505060405160208183030381529060405280519060200120614d6f565b9050631626ba7e60e01b6001600160e01b0319166122f78286868080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506114f492505050565b6001600160e01b03191614612343576040805162461bcd60e51b815260206004820152600d60248201526c1cda59c81b9bdd081d985b1a59609a1b604482015290519081900360640190fd5b604c548714612385576040805162461bcd60e51b81526020600482015260096024820152687478207265706c617960b81b604482015290519081900360640190fd5b61238d614dc0565b60006060306001600160a01b03168888604051808383808284376040519201945060009350909150508083038183865af19150503d80600081146123ed576040519150601f19603f3d011682016040523d82523d6000602084013e6123f2565b606091505b50915091508181906124455760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561207257818101518382015260200161205a565b507f823dbcf2b7b0f265871963ca65ac033f6b4c71e0d82cd123d2ff23d752dc21c188888360405180806020018060200183810383528686828181526020019250808284376000838201819052601f909101601f191690920185810384528651815286516020918201939188019250908190849084905b838110156124d45781810151838201526020016124bc565b50505050905090810190601f1680156125015780820380516001836020036101000a031916815260200191505b509550505050505060405180910390a1505050505050505050565b606060388054806020026020016040519081016040528092919081815260200182805480156118e9576020028201919060005260206000209081546001600160a01b031681526001909101906020018083116118cb575050505050905090565b61258533614515565b6125c4576040805162461bcd60e51b815260206004820152601a6024820152600080516020615cd1833981519152604482015290519081900360640190fd5b603a5460ff16612613576040805162461bcd60e51b81526020600482015260156024820152743737903832b73234b7339039bab136b4b9b9b4b7b760591b604482015290519081900360640190fd5b6126746038805480602002602001604051908101604052809291908181526020018280548015611438576020028201919060005260206000209081546001600160a01b0316815260019091019060200180831161141a575050505050613ae6565b81146126b15760405162461bcd60e51b8152600401808060200182810382526023815260200180615d8d6023913960400191505060405180910390fd5b60005b6038548110156127985760366000603883815481106126cf57fe5b60009182526020808320909101546001600160a01b0316835282019290925260400190205460ff16612790576001603660006038848154811061270e57fe5b6000918252602080832091909101546001600160a01b031683528201929092526040019020805460ff191691151591909117905560388054603791908390811061275457fe5b60009182526020808320909101548354600181018555938352912090910180546001600160a01b0319166001600160a01b039092169190911790555b6001016126b4565b507fb2f6cccee7a369e23e293c25aa19bef80af11eb26deba3ea0f2a02783f752e4a33603860405180836001600160a01b03166001600160a01b0316815260200180602001828103825283818154815260200191508054801561282457602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311612806575b5050935050505060405180910390a161283f60386000615bcf565b50603a805460ff19169055565b60006116fa603b614728565b61286133614498565b8061286b57503330145b6128af576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b603a5460ff161580156128ca5750603a54610100900460ff16155b61291b576040805162461bcd60e51b815260206004820152601c60248201527f77686974656c6973742073756d62697373696f6e2070656e64696e6700000000604482015290519081900360640190fd5b603a5462010000900460ff16612974576040805162461bcd60e51b81526020600482015260196024820152781dda1a5d195b1a5cdd081b9bdd081a5b9a5d1a585b1a5e9959603a1b604482015290519081900360640190fd5b806129b8576040805162461bcd60e51b815260206004820152600f60248201526e195b5c1d1e481dda1a5d195b1a5cdd608a1b604482015290519081900360640190fd5b6129c460398383615bed565b50603a805461ff00191661010017905560408051602080840282810182019093528382527ffbc0e5ca6c7e4858daf0fdb185ef5186203e74ec9c64737e93c0aeaec596e1d19285928592612a3392859185918291850190849080828437600092019190915250613ae692505050565b60405180806020018381526020018281038252858582818152602001925060200280828437600083820152604051601f909101601f1916909201829003965090945050505050a15050565b600054610100900460ff1680612a975750612a97614e08565b80612aa5575060005460ff16155b612ae05760405162461bcd60e51b815260040180806020018281038252602e815260200180615d12602e913960400191505060405180910390fd5b600054610100900460ff16158015612b0b576000805460ff1961ff0019909116610100171660011790555b612b1486614e0e565b612b1d84614f15565b612b278888614fc3565b612b3082615122565b612b3861520d565b612b41856152fe565b604d8390558015612b58576000805461ff00191690555b5050505050505050565b612b6b33614498565b612bb5576040805162461bcd60e51b815260206004820152601660248201527539b2b73232b91034b9903737ba1030b71037bbb732b960511b604482015290519081900360640190fd5b603554600160a01b900460ff16612c13576040805162461bcd60e51b815260206004820152601d60248201527f6f776e657273686970206973206e6f74207472616e7366657261626c65000000604482015290519081900360640190fd5b603580546001600160a01b0319169055604080516000808252602082015281517f850b3df64837d7d518b45f5aa64d104652c3b80eb5b34a8e3d9eb666cb7cdea5929181900390910190a1565b612c6933614498565b80612c7357503330145b612cb7576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b8066038d7ea4c6800011158015612cd657506706f05b59d3b200008111155b612d1d576040805162461bcd60e51b815260206004820152601360248201527206f7574206f662072616e676520746f702d757606c1b604482015290519081900360640190fd5b612d2e60408263ffffffff61475d16565b6040805182815290517faf2a77cd04c3cc155588dd3bf67b310ab4fb3b1da3cf6b8d7d4d2aa1d09b794c9181900360200190a150565b604d5490565b6033546001600160a01b031690565b60006116fa6040614728565b612d8e33614498565b80612d9857503330145b612ddc576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b603a5460ff16158015612df75750603a54610100900460ff16155b612e48576040805162461bcd60e51b815260206004820152601c60248201527f77686974656c6973742073756d62697373696f6e2070656e64696e6700000000604482015290519081900360640190fd5b8181808060200260200160405190810160405280939291908181526020018383602002808284376000920182905250925050505b8151811015612f6457612ea1828281518110612e9457fe5b6020026020010151614498565b15612eec576040805162461bcd60e51b8152602060048201526016602482015275636f6e7461696e73206f776e6572206164647265737360501b604482015290519081900360640190fd5b60006001600160a01b0316828281518110612f0357fe5b60200260200101516001600160a01b03161415612f5c576040805162461bcd60e51b8152602060048201526012602482015271636f6e7461696e732030206164647265737360701b604482015290519081900360640190fd5b600101612e7c565b50603a5462010000900460ff16612fbe576040805162461bcd60e51b81526020600482015260196024820152781dda1a5d195b1a5cdd081b9bdd081a5b9a5d1a585b1a5e9959603a1b604482015290519081900360640190fd5b81613002576040805162461bcd60e51b815260206004820152600f60248201526e195b5c1d1e481dda1a5d195b1a5cdd608a1b604482015290519081900360640190fd5b61300e60388484615bed565b50603a805460ff1916600117905560408051602080850282810182019093528482527f9c80b3b5f68b3e017766d59e8d09b34efe6462b05c398f35cab9e271d9bc3b9c928692869261307b92859185918291850190849080828437600092019190915250613ae692505050565b60405180806020018381526020018281038252858582818152602001925060200280828437600083820152604051601f909101601f1916909201829003965090945050505050a1505050565b60455490565b6035546001600160a01b031690565b6130e533614498565b61312f576040805162461bcd60e51b815260206004820152601660248201527539b2b73232b91034b9903737ba1030b71037bbb732b960511b604482015290519081900360640190fd5b613137614dc0565b565b603b5490565b61314833614498565b8061315257503330145b613196576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b8051602080820191906000808060605b86851015612b58576131bf86605463ffffffff61545016565b888601805160148201516034909201805193995060609190911c9650909450909250905061320460546131f8878563ffffffff6154ad16565b9063ffffffff6154ad16565b94508685111561324b576040805162461bcd60e51b815260206004820152600d60248201526c6f7574206f6620626f756e647360981b604482015290519081900360640190fd5b8161326157506040805160208101909152600081525b61326c848483611ce4565b506131a6565b604b5460ff1690565b61328433614498565b8061328e57503330145b6132d2576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b604654811115613324576040805162461bcd60e51b81526020600482015260186024820152771bdd5d081bd9881c985b99d9481b1bd85908185b5bdd5b9d60421b604482015290519081900360640190fd5b61333560478263ffffffff61475d16565b6040805182815290517fc178d379965e5657b6fc57494e392f121a14119215dfb422aad7db4cc03f2d109181900360200190a150565b61337433614498565b6133be576040805162461bcd60e51b815260206004820152601660248201527539b2b73232b91034b9903737ba1030b71037bbb732b960511b604482015290519081900360640190fd5b603554600160a01b900460ff1661341c576040805162461bcd60e51b815260206004820152601d60248201527f6f776e657273686970206973206e6f74207472616e7366657261626c65000000604482015290519081900360640190fd5b6001600160a01b0382166134615760405162461bcd60e51b8152600401808060200182810382526023815260200180615d406023913960400191505060405180910390fd5b6035805460ff60a01b1916600160a01b83151502179055806134ba57604080516001600160a01b038416815290517f808639ff9c8e4732d60b6c2330de498035416d229f27a77d259680895efec1229181900360200190a15b603554604080516001600160a01b039283168152918416602083015280517f850b3df64837d7d518b45f5aa64d104652c3b80eb5b34a8e3d9eb666cb7cdea59281900390910190a150603580546001600160a01b0319166001600160a01b0392909216919091179055565b60008060008061353486615507565b5050509350935093505080156135a9578161357f576040805162461bcd60e51b81526020600482015260066024820152650726174653d360d41b604482015290519081900360640190fd5b61359f83613593878563ffffffff61569916565b9063ffffffff6156f216565b9350505050611563565b50600095945050505050565b603f5460ff1690565b603a5462010000900460ff1681565b6135d633614498565b806135e057503330145b613624576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b8080613661576040805162461bcd60e51b8152602060048201526007602482015266076616c75653d360cc1b604482015290519081900360640190fd5b6001600160a01b0384166136ac576040805162461bcd60e51b815260206004820152600d60248201526c064657374696e6174696f6e3d3609c1b604482015290519081900360640190fd5b6001600160a01b03841660009081526036602052604090205460ff166136fc57816001600160a01b038416156136e9576136e68484613525565b90505b6136fa603b8263ffffffff6147d816565b505b61370784848461575c565b604080516001600160a01b0380871682528516602082015280820184905290517fd1ba4ac2e2a11b5101f6cb4d978f514a155b421e8ec396d2d9abaf0bb02917ee9181900360600190a150505050565b604a5490565b61376633614515565b6137a5576040805162461bcd60e51b815260206004820152601a6024820152600080516020615cd1833981519152604482015290519081900360640190fd5b603a54610100900460ff166137f9576040805162461bcd60e51b81526020600482015260156024820152743737903832b73234b7339039bab136b4b9b9b4b7b760591b604482015290519081900360640190fd5b61385a6039805480602002602001604051908101604052809291908181526020018280548015611438576020028201919060005260206000209081546001600160a01b0316815260019091019060200180831161141a575050505050613ae6565b81146138975760405162461bcd60e51b8152600401808060200182810382526023815260200180615d8d6023913960400191505060405180910390fd5b60005b603954811015613a255760366000603983815481106138b557fe5b60009182526020808320909101546001600160a01b0316835282019290925260400190205460ff1615613a1d57600060366000603984815481106138f557fe5b6000918252602080832091909101546001600160a01b031683528201929092526040018120805460ff1916921515929092179091555b60375461393f90600163ffffffff61545016565b811015613a07576039828154811061395357fe5b600091825260209091200154603780546001600160a01b03909216918390811061397957fe5b6000918252602090912001546001600160a01b031614156139ff576037805460001981019081106139a657fe5b600091825260209091200154603780546001600160a01b0390921691839081106139cc57fe5b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b03160217905550613a07565b60010161392b565b506037805490613a1b906000198301615c50565b505b60010161389a565b507fd218c430fa348f4ce67791021b6b89c0c3eacd4ead1d8f5b83c60038ec28249b33603960405180836001600160a01b03166001600160a01b03168152602001806020018281038252838181548152602001915080548015613ab157602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311613a93575b5050935050505060405180910390a1613acc60396000615bcf565b50603a805461ff0019169055565b60435490565b604c5481565b60008160405160200180828051906020019060200280838360005b83811015613b19578181015183820152602001613b01565b50505050905001915050604051602081830303815290604052805190602001209050919050565b613b4933614498565b80613b585750613b5833614515565b613ba2576040805162461bcd60e51b815260206004820152601660248201527537b7363c9037bbb732b93e3e31b7b73a3937b63632b960511b604482015290519081900360640190fd5b603a54610100900460ff16613bf6576040805162461bcd60e51b81526020600482015260156024820152743737903832b73234b7339039bab136b4b9b9b4b7b760591b604482015290519081900360640190fd5b613c576039805480602002602001604051908101604052809291908181526020018280548015611438576020028201919060005260206000209081546001600160a01b0316815260019091019060200180831161141a575050505050613ae6565b8114613c945760405162461bcd60e51b8152600401808060200182810382526023815260200180615d8d6023913960400191505060405180910390fd5b613ca060396000615bcf565b603a805461ff0019169055604080513381526020810183905281517f13c935eb475aa0f6e931fece83e2ac44569ce2d53460d29a6dedab40b965c8a3929181900390910190a150565b60378181548110613cf657fe5b6000918252602090912001546001600160a01b0316905081565b60475490565b603a54610100900460ff1681565b60345490565b8080613d67576040805162461bcd60e51b8152602060048201526007602482015266076616c75653d360cc1b604482015290519081900360640190fd5b613d7033614498565b80613d7f5750613d7f33614515565b613dc9576040805162461bcd60e51b815260206004820152601660248201527537b7363c9037bbb732b93e3e31b7b73a3937b63632b960511b604482015290519081900360640190fd5b613dda60408363ffffffff6147d816565b613de26130cd565b6001600160a01b03166108fc839081150290604051600060405180830381858888f19350505050158015613e1a573d6000803e3d6000fd5b507f611b7c0d84fda988026215bef9b3e4d81cbceced7e679be6d5e044b588467c0e33613e456130cd565b604080516001600160a01b03938416815291909216602082015280820185905290519081900360600190a15050565b613e7d33614515565b613ebc576040805162461bcd60e51b815260206004820152601a6024820152600080516020615cd1833981519152604482015290519081900360640190fd5b611ca6603b8263ffffffff61582616565b6000613ed7615876565b6001600160a01b0316836001600160a01b03161415613ef7575080611563565b816001600160a01b03841615613fbc576000806000613f1587615507565b5050509350935093505080613f67576040805162461bcd60e51b8152602060048201526013602482015272746f6b656e206e6f7420617661696c61626c6560681b604482015290519081900360640190fd5b81613fa2576040805162461bcd60e51b81526020600482015260066024820152650726174653d360d41b604482015290519081900360640190fd5b613fb683613593888563ffffffff61569916565b93505050505b6000806000613fc96158ec565b505050935093509350508061401b576040805162461bcd60e51b8152602060048201526013602482015272746f6b656e206e6f7420617661696c61626c6560681b604482015290519081900360640190fd5b81614061576040805162461bcd60e51b81526020600482015260116024820152700737461626c65636f696e20726174653d3607c1b604482015290519081900360640190fd5b61407582613593868663ffffffff61569916565b979650505050505050565b61408933614515565b6140c8576040805162461bcd60e51b815260206004820152601a6024820152600080516020615cd1833981519152604482015290519081900360640190fd5b611c0060478263ffffffff61582616565b6140e233614515565b614121576040805162461bcd60e51b815260206004820152601a6024820152600080516020615cd1833981519152604482015290519081900360640190fd5b6112f060408263ffffffff61582616565b61413b33614498565b8061414557503330145b614189576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b8181808060200260200160405190810160405280939291908181526020018383602002808284376000920182905250925050505b8151811015614298576141d5828281518110612e9457fe5b15614220576040805162461bcd60e51b8152602060048201526016602482015275636f6e7461696e73206f776e6572206164647265737360501b604482015290519081900360640190fd5b60006001600160a01b031682828151811061423757fe5b60200260200101516001600160a01b03161415614290576040805162461bcd60e51b8152602060048201526012602482015271636f6e7461696e732030206164647265737360701b604482015290519081900360640190fd5b6001016141bd565b50603a5462010000900460ff16156142ef576040805162461bcd60e51b81526020600482015260156024820152741dda1a5d195b1a5cdd081a5b9a5d1a585b1a5e9959605a1b604482015290519081900360640190fd5b60005b828110156143e0576036600085858481811061430a57fe5b602090810292909201356001600160a01b03168352508101919091526040016000205460ff166143d85760016036600086868581811061434657fe5b905060200201356001600160a01b03166001600160a01b03166001600160a01b0316815260200190815260200160002060006101000a81548160ff021916908315150217905550603784848381811061439b57fe5b835460018101855560009485526020948590200180546001600160a01b0319166001600160a01b0395909202939093013593909316929092179055505b6001016142f2565b50603a805462ff0000191662010000179055604080513380825260208201838152603780549484018590527fb2f6cccee7a369e23e293c25aa19bef80af11eb26deba3ea0f2a02783f752e4a949293909290919060608301908490801561447057602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311614452575b5050935050505060405180910390a1505050565b60445460ff1690565b600061156382614697565b6035546001600160a01b0390811691161490565b600482015460ff16156144fa576040805162461bcd60e51b81526020600482015260116024820152701b1a5b5a5d08185b1c9958591e481cd95d607a1b604482015290519081900360640190fd5b6145048282615a5b565b50600401805460ff19166001179055565b600061452260345461484e565b6001600160a01b031663b429afeb836040518263ffffffff1660e01b815260040180826001600160a01b03166001600160a01b0316815260200191505060206040518083038186803b15801561457757600080fd5b505afa15801561458b573d6000803e3d6000fd5b505050506040513d60208110156145a157600080fd5b505192915050565b600081516041146145bc57506000611563565b60208201516040830151606084015160001a7f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08211156146025760009350505050611563565b8060ff16601b1415801561461a57508060ff16601c14155b1561462b5760009350505050611563565b6040805160008152602080820180845289905260ff8416828401526060820186905260808201859052915160019260a0808401939192601f1981019281900390910190855afa158015614682573d6000803e3d6000fd5b5050604051601f190151979650505050505050565b60006001600160a01b0382161561472157604080516370a0823160e01b815230600482015290516001600160a01b038416916370a08231916024808301926020929190829003018186803b1580156146ee57600080fd5b505afa158015614702573d6000803e3d6000fd5b505050506040513d602081101561471857600080fd5b50519050611216565b5047611216565b6002810154600090614743906201518063ffffffff6154ad16565b42111561475257508054611216565b506001810154611216565b600482015460ff166147b6576040805162461bcd60e51b815260206004820152601960248201527f6c696d6974206861736e2774206265656e207365742079657400000000000000604482015290519081900360640190fd5b600390910155565b6000806147ca83615507565b509098975050505050505050565b6147e182615a7e565b808260010154101561482d576040805162461bcd60e51b815260206004820152601060248201526f185d985a5b18589b194f185b5bdd5b9d60821b604482015290519081900360640190fd5b6001820154614842908263ffffffff61545016565b82600101819055505050565b6033546000906001600160a01b03166148ae576040805162461bcd60e51b815260206004820152601d60248201527f454e535265736f6c7661626c65206e6f7420696e697469616c697a6564000000604482015290519081900360640190fd5b60335460408051630178b8bf60e01b81526004810185905290516001600160a01b0390921691630178b8bf91602480820192602092909190829003018186803b1580156148fa57600080fd5b505afa15801561490e573d6000803e3d6000fd5b505050506040513d602081101561492457600080fd5b505160408051631d9dabef60e11b81526004810185905290516001600160a01b0390921691633b3b57de91602480820192602092909190829003018186803b15801561457757600080fd5b8015806149f5575060408051636eb1769f60e11b81523060048201526001600160a01b03848116602483015291519185169163dd62ed3e91604480820192602092909190829003018186803b1580156149c757600080fd5b505afa1580156149db573d6000803e3d6000fd5b505050506040513d60208110156149f157600080fd5b5051155b614a305760405162461bcd60e51b8152600401808060200182810382526036815260200180615db06036913960400191505060405180910390fd5b604080516001600160a01b038416602482015260448082018490528251808303909101815260649091019091526020810180516001600160e01b031663095ea7b360e01b179052614a82908490614bb1565b505050565b3b151590565b600080614a9983615507565b509198975050505050505050565b600080614ab560455461484e565b6001600160a01b031663afc72e9385856040518363ffffffff1660e01b815260040180836001600160a01b03166001600160a01b0316815260200180602001828103825283818151815260200191508051906020019080838360005b83811015614b29578181015183820152602001614b11565b50505050905090810190601f168015614b565780820380516001836020036101000a031916815260200191505b509350505050604080518083038186803b158015614b7357600080fd5b505afa158015614b87573d6000803e3d6000fd5b505050506040513d6040811015614b9d57600080fd5b508051602090910151909590945092505050565b614bc3826001600160a01b0316614a87565b614c14576040805162461bcd60e51b815260206004820152601f60248201527f5361666545524332303a2063616c6c20746f206e6f6e2d636f6e747261637400604482015290519081900360640190fd5b60006060836001600160a01b0316836040518082805190602001908083835b60208310614c525780518252601f199092019160209182019101614c33565b6001836020036101000a0380198251168184511680821785525050505050509050019150506000604051808303816000865af19150503d8060008114614cb4576040519150601f19603f3d011682016040523d82523d6000602084013e614cb9565b606091505b509150915081614d10576040805162461bcd60e51b815260206004820181905260248201527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564604482015290519081900360640190fd5b805115614d6957808060200190516020811015614d2c57600080fd5b5051614d695760405162461bcd60e51b815260040180806020018281038252602a815260200180615d63602a913960400191505060405180910390fd5b50505050565b604080517f19457468657265756d205369676e6564204d6573736167653a0a333200000000602080830191909152603c8083019490945282518083039094018452605c909101909152815191012090565b604c80546001019081905560408051338152602081019290925280517fab0423a75986556234aecd171c46ce7f5e45607d8070bf5230f2735b50322bff9281900390910190a1565b303b1590565b600054610100900460ff1680614e275750614e27614e08565b80614e35575060005460ff16155b614e705760405162461bcd60e51b815260040180806020018281038252602e815260200180615d12602e913960400191505060405180910390fd5b600054610100900460ff16158015614e9b576000805460ff1961ff0019909116610100171660011790555b6001600160a01b038216614ee4576040805162461bcd60e51b815260206004820152600b60248201526a0656e7352656720697320360ac1b604482015290519081900360640190fd5b603380546001600160a01b0319166001600160a01b0384161790558015614f11576000805461ff00191690555b5050565b600054610100900460ff1680614f2e5750614f2e614e08565b80614f3c575060005460ff16155b614f775760405162461bcd60e51b815260040180806020018281038252602e815260200180615d12602e913960400191505060405180910390fd5b600054610100900460ff16158015614fa2576000805460ff1961ff0019909116610100171660011790555b8115614fae5760348290555b8015614f11576000805461ff00191690555050565b600054610100900460ff1680614fdc5750614fdc614e08565b80614fea575060005460ff16155b6150255760405162461bcd60e51b815260040180806020018281038252602e815260200180615d12602e913960400191505060405180910390fd5b600054610100900460ff16158015615050576000805460ff1961ff0019909116610100171660011790555b603580546001600160a01b0319166001600160a01b0385161760ff60a01b1916600160a01b8415158102919091179182905560ff9104166150c857604080516001600160a01b038516815290517f808639ff9c8e4732d60b6c2330de498035416d229f27a77d259680895efec1229181900360200190a15b60408051600081526001600160a01b038516602082015281517f850b3df64837d7d518b45f5aa64d104652c3b80eb5b34a8e3d9eb666cb7cdea5929181900390910190a18015614a82576000805461ff0019169055505050565b600054610100900460ff168061513b575061513b614e08565b80615149575060005460ff16155b6151845760405162461bcd60e51b815260040180806020018281038252602e815260200180615d12602e913960400191505060405180910390fd5b600054610100900460ff161580156151af576000805460ff1961ff0019909116610100171660011790555b6040805160a08101825283815260208101849052429181018290526000606082018190526080909101819052603b849055603c849055603d91909155603e55603f805460ff191690558015614f11576000805461ff00191690555050565b600054610100900460ff16806152265750615226614e08565b80615234575060005460ff16155b61526f5760405162461bcd60e51b815260040180806020018281038252602e815260200180615d12602e913960400191505060405180910390fd5b600054610100900460ff1615801561529a576000805460ff1961ff0019909116610100171660011790555b6040805160a0810182526706f05b59d3b2000080825260208201819052428284018190526000606084018190526080909301839052928190556041556042919091556043556044805460ff1916905580156152fb576000805461ff00191690555b50565b600054610100900460ff16806153175750615317614e08565b80615325575060005460ff16155b6153605760405162461bcd60e51b815260040180806020018281038252602e815260200180615d12602e913960400191505060405180910390fd5b600054610100900460ff1615801561538b576000805460ff1961ff0019909116610100171660011790555b61539482615ad6565b600061539e6158ec565b5050505050915050600081116153eb576040805162461bcd60e51b815260206004820152600d60248201526c37379039ba30b13632b1b7b4b760991b604482015290519081900360640190fd5b6127100260468190556040805160a081018252828152602081018390524291810182905260006060820181905260809091018190526047839055604892909255604955604a55604b805460ff191690558015614f11576000805461ff00191690555050565b6000828211156154a7576040805162461bcd60e51b815260206004820152601e60248201527f536166654d6174683a207375627472616374696f6e206f766572666c6f770000604482015290519081900360640190fd5b50900390565b6000828201838110156121d4576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b606060008060008060008061551d60455461484e565b6001600160a01b0316631f69565f896040518263ffffffff1660e01b815260040180826001600160a01b03166001600160a01b0316815260200191505060006040518083038186803b15801561557257600080fd5b505afa158015615586573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405260e08110156155af57600080fd5b8101908080516040519392919084600160201b8211156155ce57600080fd5b9083019060208201858111156155e357600080fd5b8251600160201b8111828201881017156155fc57600080fd5b82525081516020918201929091019080838360005b83811015615629578181015183820152602001615611565b50505050905090810190601f1680156156565780820380516001836020036101000a031916815260200191505b5060409081526020820151908201516060830151608084015160a085015160c090950151979e50929c50909a509850965094509192505050919395979092949650565b6000826156a857506000611563565b828202828482816156b557fe5b04146121d45760405162461bcd60e51b8152600401808060200182810382526021815260200180615cf16021913960400191505060405180910390fd5b6000808211615748576040805162461bcd60e51b815260206004820152601a60248201527f536166654d6174683a206469766973696f6e206279207a65726f000000000000604482015290519081900360640190fd5b600082848161575357fe5b04949350505050565b6001600160a01b03821661580c576040516000906001600160a01b0385169083908381818185875af1925050503d80600081146157b5576040519150601f19603f3d011682016040523d82523d6000602084013e6157ba565b606091505b5050905080615806576040805162461bcd60e51b81526020600482015260136024820152721cd85999551c985b9cd9995c8819985a5b1959606a1b604482015290519081900360640190fd5b50614a82565b614a826001600160a01b038316848363ffffffff615b7d16565b808260030154146158685760405162461bcd60e51b8152600401808060200182810382526022815260200180615caf6022913960400191505060405180910390fd5b614f11828360030154615a5b565b600061588360455461484e565b6001600160a01b031663e9cbd8226040518163ffffffff1660e01b815260040160206040518083038186803b1580156158bb57600080fd5b505afa1580156158cf573d6000803e3d6000fd5b505050506040513d60208110156158e557600080fd5b5051905090565b606060008060008060008061590260455461484e565b6001600160a01b0316633efec5e96040518163ffffffff1660e01b815260040160006040518083038186803b15801561593a57600080fd5b505afa15801561594e573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405260e081101561597757600080fd5b8101908080516040519392919084600160201b82111561599657600080fd5b9083019060208201858111156159ab57600080fd5b8251600160201b8111828201881017156159c457600080fd5b82525081516020918201929091019080838360005b838110156159f15781810151838201526020016159d9565b50505050905090810190601f168015615a1e5780820380516001836020036101000a031916815260200191505b5060409081526020820151908201516060830151608084015160a085015160c090950151979f939e50919c509a5098509096509294509192505050565b615a6482615a7e565b8082556001820154811015614f1157815460018301555050565b6002810154615a96906201518063ffffffff6154ad16565b4211156152fb57426002820155805460018201556040517fe93bc25276d408d390778e7a8b926f2f67209c43ed540081b951fe128f0d3cd290600090a150565b600054610100900460ff1680615aef5750615aef614e08565b80615afd575060005460ff16155b615b385760405162461bcd60e51b815260040180806020018281038252602e815260200180615d12602e913960400191505060405180910390fd5b600054610100900460ff16158015615b63576000805460ff1961ff0019909116610100171660011790555b60458290558015614f11576000805461ff00191690555050565b604080516001600160a01b038416602482015260448082018490528251808303909101815260649091019091526020810180516001600160e01b031663a9059cbb60e01b179052614a82908490614bb1565b50805460008255906000526020600020908101906152fb9190615c70565b828054828255906000526020600020908101928215615c40579160200282015b82811115615c405781546001600160a01b0319166001600160a01b03843516178255602090920191600190910190615c0d565b50615c4c929150615c8a565b5090565b815481835581811115614a8257600083815260209020614a829181019083015b61121f91905b80821115615c4c5760008155600101615c76565b61121f91905b80821115615c4c5780546001600160a01b0319168155600101615c9056fe636f6e6669726d65642f7375626d6974746564206c696d6974206d69736d6174636873656e646572206973206e6f74206120636f6e74726f6c6c6572000000000000536166654d6174683a206d756c7469706c69636174696f6e206f766572666c6f77436f6e747261637420696e7374616e63652068617320616c7265616479206265656e20696e697469616c697a65646f776e65722063616e6e6f742062652073657420746f207a65726f20616464726573735361666545524332303a204552433230206f7065726174696f6e20646964206e6f7420737563636565646e6f6e2d6d61746368696e672070656e64696e672077686974656c69737420686173685361666545524332303a20617070726f76652066726f6d206e6f6e2d7a65726f20746f206e6f6e2d7a65726f20616c6c6f77616e6365a265627a7a72315820f41e568cc3d2f22651e48c9958c51441796c3f589bfd42d1220a454a9a86961264736f6c63430005110032"

// DeployWallet deploys a new Ethereum contract, binding an instance of Wallet to it.
func DeployWallet(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Wallet, error) {
	parsed, err := abi.JSON(strings.NewReader(WalletABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(WalletBin), backend)
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

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address _asset) constant returns(uint256)
func (_Wallet *WalletCaller) GetBalance(opts *bind.CallOpts, _asset common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Wallet.contract.Call(opts, out, "getBalance", _asset)
	return *ret0, err
}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address _asset) constant returns(uint256)
func (_Wallet *WalletSession) GetBalance(_asset common.Address) (*big.Int, error) {
	return _Wallet.Contract.GetBalance(&_Wallet.CallOpts, _asset)
}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address _asset) constant returns(uint256)
func (_Wallet *WalletCallerSession) GetBalance(_asset common.Address) (*big.Int, error) {
	return _Wallet.Contract.GetBalance(&_Wallet.CallOpts, _asset)
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

// InitializeWallet is a paid mutator transaction binding the contract method 0x69efdfc0.
//
// Solidity: function initializeWallet(address _owner_, bool _transferable_, address _ens_, bytes32 _tokenWhitelistNode_, bytes32 _controllerNode_, bytes32 _licenceNode_, uint256 _spendLimit_) returns()
func (_Wallet *WalletTransactor) InitializeWallet(opts *bind.TransactOpts, _owner_ common.Address, _transferable_ bool, _ens_ common.Address, _tokenWhitelistNode_ [32]byte, _controllerNode_ [32]byte, _licenceNode_ [32]byte, _spendLimit_ *big.Int) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "initializeWallet", _owner_, _transferable_, _ens_, _tokenWhitelistNode_, _controllerNode_, _licenceNode_, _spendLimit_)
}

// InitializeWallet is a paid mutator transaction binding the contract method 0x69efdfc0.
//
// Solidity: function initializeWallet(address _owner_, bool _transferable_, address _ens_, bytes32 _tokenWhitelistNode_, bytes32 _controllerNode_, bytes32 _licenceNode_, uint256 _spendLimit_) returns()
func (_Wallet *WalletSession) InitializeWallet(_owner_ common.Address, _transferable_ bool, _ens_ common.Address, _tokenWhitelistNode_ [32]byte, _controllerNode_ [32]byte, _licenceNode_ [32]byte, _spendLimit_ *big.Int) (*types.Transaction, error) {
	return _Wallet.Contract.InitializeWallet(&_Wallet.TransactOpts, _owner_, _transferable_, _ens_, _tokenWhitelistNode_, _controllerNode_, _licenceNode_, _spendLimit_)
}

// InitializeWallet is a paid mutator transaction binding the contract method 0x69efdfc0.
//
// Solidity: function initializeWallet(address _owner_, bool _transferable_, address _ens_, bytes32 _tokenWhitelistNode_, bytes32 _controllerNode_, bytes32 _licenceNode_, uint256 _spendLimit_) returns()
func (_Wallet *WalletTransactorSession) InitializeWallet(_owner_ common.Address, _transferable_ bool, _ens_ common.Address, _tokenWhitelistNode_ [32]byte, _controllerNode_ [32]byte, _licenceNode_ [32]byte, _spendLimit_ *big.Int) (*types.Transaction, error) {
	return _Wallet.Contract.InitializeWallet(&_Wallet.TransactOpts, _owner_, _transferable_, _ens_, _tokenWhitelistNode_, _controllerNode_, _licenceNode_, _spendLimit_)
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
