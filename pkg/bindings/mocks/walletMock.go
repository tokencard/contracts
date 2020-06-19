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

// WalletMockABI is the input ABI used to generate the binding from.
const WalletMockABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"_addresses\",\"type\":\"address[]\"}],\"name\":\"AddedToWhitelist\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"_assets\",\"type\":\"address[]\"}],\"name\":\"BulkTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"CancelledWhitelistAddition\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"CancelledWhitelistRemoval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_returndata\",\"type\":\"bytes\"}],\"name\":\"ExecutedRelayedTransaction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_destination\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_returndata\",\"type\":\"bytes\"}],\"name\":\"ExecutedTransaction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_currentNonce\",\"type\":\"uint256\"}],\"name\":\"IncreasedRelayNonce\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"LoadedTokenCard\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_locked\",\"type\":\"address\"}],\"name\":\"LockedOwnership\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"Received\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"_addresses\",\"type\":\"address[]\"}],\"name\":\"RemovedFromWhitelist\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"SetGasTopUpLimit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"SetLoadLimit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"SetSpendLimit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"SubmittedGasTopUpLimitUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"SubmittedLoadLimitUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"SubmittedSpendLimitUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"_addresses\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"SubmittedWhitelistAddition\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"_addresses\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"SubmittedWhitelistRemoval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"ToppedUpGas\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"Transferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"TransferredOwnership\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"UpdatedAvailableLimit\",\"type\":\"event\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"constant\":true,\"inputs\":[],\"name\":\"WALLET_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_transactionBatch\",\"type\":\"bytes\"}],\"name\":\"batchExecuteTransaction\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"_assets\",\"type\":\"address[]\"}],\"name\":\"bulkTransfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_addresses\",\"type\":\"address[]\"}],\"name\":\"calculateHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"cancelWhitelistAddition\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"cancelWhitelistRemoval\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"confirmGasTopUpLimitUpdate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"confirmLoadLimitUpdate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"confirmSpendLimitUpdate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"confirmWhitelistAddition\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"confirmWhitelistRemoval\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"controllerNode\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"convertToEther\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"convertToStablecoin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ensRegistry\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"executeRelayedTransaction\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_destination\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"executeTransaction\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"gasTopUpLimitAvailable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"gasTopUpLimitControllerConfirmationRequired\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"gasTopUpLimitPending\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"gasTopUpLimitValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"increaseRelayNonce\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_owner_\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_transferable_\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"_ens_\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_tokenWhitelistNode_\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_controllerNode_\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_licenceNode_\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_spendLimit_\",\"type\":\"uint256\"}],\"name\":\"initializeWallet\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isSetWhitelist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isTransferable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hashedData\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"isValidSignature\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"isValidSignature\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"licenceNode\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"loadLimitAvailable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"loadLimitControllerConfirmationRequired\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"loadLimitPending\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"loadLimitValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"loadTokenCard\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingWhitelistAddition\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingWhitelistRemoval\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"relayNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"setGasTopUpLimit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"setLoadLimit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"setSpendLimit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_addresses\",\"type\":\"address[]\"}],\"name\":\"setWhitelist\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"spendLimitAvailable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"spendLimitControllerConfirmationRequired\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"spendLimitPending\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"spendLimitValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"submitGasTopUpLimitUpdate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"submitLoadLimitUpdate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"submitSpendLimitUpdate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_addresses\",\"type\":\"address[]\"}],\"name\":\"submitWhitelistAddition\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_addresses\",\"type\":\"address[]\"}],\"name\":\"submitWhitelistRemoval\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"submittedWhitelistAddition\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"submittedWhitelistRemoval\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_interfaceID\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tokenWhitelistNode\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"topUpGas\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_transferable\",\"type\":\"bool\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"whitelistArray\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"whitelistMap\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// WalletMockBin is the compiled bytecode used for deploying new contracts.
var WalletMockBin = "0x60806040527f7f2ce995617d2816b426c5c8698c5ec2952f7a34bb10f38326f74933d589369760015534801561003457600080fd5b5061596380620000456000396000f3fe6080604052600436106103975760003560e01c80637fd004fa116101dc578063cc0e7e5611610102578063e2b4ce97116100a0578063f40b51f81161006f578063f40b51f814611116578063f41c431914611140578063f42176481461116a578063f776f518146111e557610397565b8063e2b4ce9714611074578063e61c51ca14611089578063eadd3cea146110b3578063f36febda146110dd57610397565b8063ce0b5bd5116100dc578063ce0b5bd514610ff6578063d251fefc14611020578063da84b1ed1461104a578063de212bf31461105f57610397565b8063cc0e7e5614610f1e578063cccdc55614610f33578063cd7958dd14610f4857610397565b8063b221f3161161017a578063be40ba7911610149578063be40ba7914610e87578063beabacc814610e9c578063c4856cd914610edf578063cbd2ac6814610ef457610397565b8063b221f31614610dd4578063b242e53414610dfe578063b87e21ef14610e39578063bcb8b74a14610e7257610397565b806390e690c7116101b657806390e690c714610ce45780639b0dfd2714610cf9578063aaf1fc6214610d0e578063ab20599314610dbf57610397565b80637fd004fa14610c3f578063877337b014610cba5780638da5cb5b14610ccf57610397565b80633a43199f116102c15780635d2362a81161025f57806374624c551161022e57806374624c5514610bba578063747c31d614610be45780637d73b23114610bf95780637d7d004614610c2a57610397565b80635d2362a814610aba5780636137d67014610acf57806369efdfc014610b4a578063715018a614610ba557610397565b80633f579f421161029b5780633f579f42146108e357806346efe0ed146109a957806347b55a9d14610a7b5780635adc02ab14610a9057610397565b80633a43199f146108635780633bfec2541461088f5780633c672eb7146108b957610397565b80631efd0299116103395780632587a6a2116103085780632587a6a2146107a157806326d05ab2146107b6578063294f4025146107cb57806332531c3c1461083057610397565b80631efd02991461068257806320c13b0b146106975780632121dc751461076257806321ce918d1461077757610397565b8063100f23fd11610375578063100f23fd1461046e5780631127b57e146104985780631626ba7e146105225780631aa21fba146105f757610397565b806301ffc9a7146103d3578063027ef3eb1461041b5780630f3a85d814610442575b6040805133815234602082015281517f88a5966d370b9919b20f3e2c13ff65706f196a4e32cc2c12bf57088f88525874929181900390910190a1005b3480156103df57600080fd5b50610407600480360360208110156103f657600080fd5b50356001600160e01b0319166111fa565b604080519115158252519081900360200190f35b34801561042757600080fd5b50610430611214565b60408051918252519081900360200190f35b34801561044e57600080fd5b5061046c6004803603602081101561046557600080fd5b503561121b565b005b34801561047a57600080fd5b5061046c6004803603602081101561049157600080fd5b5035611327565b3480156104a457600080fd5b506104ad6114cc565b6040805160208082528351818301528351919283929083019185019080838360005b838110156104e75781810151838201526020016104cf565b50505050905090810190601f1680156105145780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561052e57600080fd5b506105da6004803603604081101561054557600080fd5b81359190810190604081016020820135600160201b81111561056657600080fd5b82018360208201111561057857600080fd5b803590602001918460018302840111600160201b8311171561059957600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506114ed945050505050565b604080516001600160e01b03199092168252519081900360200190f35b34801561060357600080fd5b5061046c6004803603604081101561061a57600080fd5b6001600160a01b038235169190810190604081016020820135600160201b81111561064457600080fd5b82018360208201111561065657600080fd5b803590602001918460208302840111600160201b8311171561067757600080fd5b509092509050611562565b34801561068e57600080fd5b506104306116e8565b3480156106a357600080fd5b506105da600480360360408110156106ba57600080fd5b810190602081018135600160201b8111156106d457600080fd5b8201836020820111156106e657600080fd5b803590602001918460018302840111600160201b8311171561070757600080fd5b919390929091602081019035600160201b81111561072457600080fd5b82018360208201111561073657600080fd5b803590602001918460018302840111600160201b8311171561075757600080fd5b5090925090506116f9565b34801561076e57600080fd5b506104076117ce565b34801561078357600080fd5b5061046c6004803603602081101561079a57600080fd5b50356117de565b3480156107ad57600080fd5b5061043061187c565b3480156107c257600080fd5b50610407611882565b3480156107d757600080fd5b506107e061188b565b60408051602080825283518183015283519192839290830191858101910280838360005b8381101561081c578181015183820152602001610804565b505050509050019250505060405180910390f35b34801561083c57600080fd5b506104076004803603602081101561085357600080fd5b50356001600160a01b03166118ed565b61046c6004803603604081101561087957600080fd5b506001600160a01b038135169060200135611902565b34801561089b57600080fd5b5061046c600480360360208110156108b257600080fd5b5035611b40565b3480156108c557600080fd5b5061046c600480360360208110156108dc57600080fd5b5035611c38565b3480156108ef57600080fd5b506104ad6004803603606081101561090657600080fd5b6001600160a01b0382351691602081013591810190606081016040820135600160201b81111561093557600080fd5b82018360208201111561094757600080fd5b803590602001918460018302840111600160201b8311171561096857600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550611cde945050505050565b3480156109b557600080fd5b5061046c600480360360608110156109cc57600080fd5b81359190810190604081016020820135600160201b8111156109ed57600080fd5b8201836020820111156109ff57600080fd5b803590602001918460018302840111600160201b83111715610a2057600080fd5b919390929091602081019035600160201b811115610a3d57600080fd5b820183602082011115610a4f57600080fd5b803590602001918460018302840111600160201b83111715610a7057600080fd5b5090925090506121d5565b348015610a8757600080fd5b506107e0612516565b348015610a9c57600080fd5b5061046c60048036036020811015610ab357600080fd5b5035612576565b348015610ac657600080fd5b50610430612846565b348015610adb57600080fd5b5061046c60048036036020811015610af257600080fd5b810190602081018135600160201b811115610b0c57600080fd5b820183602082011115610b1e57600080fd5b803590602001918460208302840111600160201b83111715610b3f57600080fd5b509092509050612852565b348015610b5657600080fd5b5061046c600480360360e0811015610b6d57600080fd5b506001600160a01b03813581169160208101351515916040820135169060608101359060808101359060a08101359060c00135612a78565b348015610bb157600080fd5b5061046c612b5c565b348015610bc657600080fd5b5061046c60048036036020811015610bdd57600080fd5b5035612c5a565b348015610bf057600080fd5b50610430612d5e565b348015610c0557600080fd5b50610c0e612d64565b604080516001600160a01b039092168252519081900360200190f35b348015610c3657600080fd5b50610430612d73565b348015610c4b57600080fd5b5061046c60048036036020811015610c6257600080fd5b810190602081018135600160201b811115610c7c57600080fd5b820183602082011115610c8e57600080fd5b803590602001918460208302840111600160201b83111715610caf57600080fd5b509092509050612d7f565b348015610cc657600080fd5b506104306130c1565b348015610cdb57600080fd5b50610c0e6130c7565b348015610cf057600080fd5b5061046c6130d6565b348015610d0557600080fd5b50610430613133565b348015610d1a57600080fd5b5061046c60048036036020811015610d3157600080fd5b810190602081018135600160201b811115610d4b57600080fd5b820183602082011115610d5d57600080fd5b803590602001918460018302840111600160201b83111715610d7e57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550613139945050505050565b348015610dcb57600080fd5b5061040761326c565b348015610de057600080fd5b5061046c60048036036020811015610df757600080fd5b5035613275565b348015610e0a57600080fd5b5061046c60048036036040811015610e2157600080fd5b506001600160a01b0381351690602001351515613365565b348015610e4557600080fd5b5061043060048036036040811015610e5c57600080fd5b506001600160a01b03813516906020013561351f565b348015610e7e57600080fd5b506104076135af565b348015610e9357600080fd5b506104076135b8565b348015610ea857600080fd5b5061046c60048036036060811015610ebf57600080fd5b506001600160a01b038135811691602081013590911690604001356135c7565b348015610eeb57600080fd5b50610430613751565b348015610f0057600080fd5b5061046c60048036036020811015610f1757600080fd5b5035613757565b348015610f2a57600080fd5b50610430613ad4565b348015610f3f57600080fd5b50610430613ada565b348015610f5457600080fd5b5061043060048036036020811015610f6b57600080fd5b810190602081018135600160201b811115610f8557600080fd5b820183602082011115610f9757600080fd5b803590602001918460208302840111600160201b83111715610fb857600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929550613ae0945050505050565b34801561100257600080fd5b5061046c6004803603602081101561101957600080fd5b5035613b3a565b34801561102c57600080fd5b50610c0e6004803603602081101561104357600080fd5b5035613ce3565b34801561105657600080fd5b50610430613d0a565b34801561106b57600080fd5b50610407613d10565b34801561108057600080fd5b50610430613d1e565b34801561109557600080fd5b5061046c600480360360208110156110ac57600080fd5b5035613d24565b3480156110bf57600080fd5b5061046c600480360360208110156110d657600080fd5b5035613e6e565b3480156110e957600080fd5b506104306004803603604081101561110057600080fd5b506001600160a01b038135169060200135613ec7565b34801561112257600080fd5b5061046c6004803603602081101561113957600080fd5b503561407a565b34801561114c57600080fd5b5061046c6004803603602081101561116357600080fd5b50356140d3565b34801561117657600080fd5b5061046c6004803603602081101561118d57600080fd5b810190602081018135600160201b8111156111a757600080fd5b8201836020820111156111b957600080fd5b803590602001918460208302840111600160201b831117156111da57600080fd5b50909250905061412c565b3480156111f157600080fd5b5061040761447e565b6001600160e01b031981166301ffc9a760e01b145b919050565b600b545b90565b61122433614487565b8061122e57503330145b611272576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b8066038d7ea4c680001115801561129157506706f05b59d3b200008111155b6112d8576040805162461bcd60e51b815260206004820152601360248201527206f7574206f662072616e676520746f702d757606c1b604482015290519081900360640190fd5b6112e9600d8263ffffffff61449b16565b604080513381526020810183905281517f41ff5d5ce3b7935893a4e7269ec5caae9cca5e3bf0eb4b21d2f443489667112e929181900390910190a150565b61133033614487565b8061133f575061133f33614504565b611389576040805162461bcd60e51b815260206004820152601660248201527537b7363c9037bbb732b93e3e31b7b73a3937b63632b960511b604482015290519081900360640190fd5b60075460ff166113d8576040805162461bcd60e51b81526020600482015260156024820152743737903832b73234b7339039bab136b4b9b9b4b7b760591b604482015290519081900360640190fd5b61143b600580548060200260200160405190810160405280929190818152602001828054801561143157602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311611413575b5050505050613ae0565b81146114785760405162461bcd60e51b81526004018080602001828103825260238152602001806158d66023913960400191505060405180910390fd5b61148460056000615718565b6007805460ff19169055604080513381526020810183905281517f7794eff834d760583543e6e510e717a5e66d2c064e225f4db448343c3e66afcf929181900390910190a150565b604051806040016040528060058152602001640332e322e360dc1b81525081565b600080611500848463ffffffff61459816565b905061150b81614487565b611550576040805162461bcd60e51b8152602060048201526011602482015270696e76616c6964207369676e617475726560781b604482015290519081900360640190fd5b50630b135d3f60e11b90505b92915050565b61156b33614487565b8061157557503330145b6115b9576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b80611602576040805162461bcd60e51b8152602060048201526014602482015273617373657420617272617920697320656d70747960601b604482015290519081900360640190fd5b60005b818110156116655760006116343085858581811061161f57fe5b905060200201356001600160a01b0316614686565b905061165c8585858581811061164657fe5b905060200201356001600160a01b0316836135c7565b50600101611605565b507fd4f62f23021706247dcffea245d104ae7ddaec7f23acf3d11d7136d5de6a69ad83838360405180846001600160a01b03166001600160a01b03168152602001806020018281038252848482818152602001925060200280828437600083820152604051601f909101601f1916909201829003965090945050505050a1505050565b60006116f46014614731565b905090565b6000808585604051602001808383808284376040805191909301818103601f190182528084528151602092830120601f8b01839004830282018301909452898152929650630b135d3f60e11b955061177094508693508991508890819084018382808284376000920191909152506114ed92505050565b6001600160e01b031916146117bc576040805162461bcd60e51b815260206004820152600d60248201526c1cda59c81b9bdd081d985b1a59609a1b604482015290519081900360640190fd5b506320c13b0b60e01b95945050505050565b600254600160a01b900460ff1690565b6117e733614487565b806117f157503330145b611835576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b61184660088263ffffffff61476616565b6040805182815290517f4b1b970c8a0fa761e7803ed70c13d7aca71904b13df60fbe03f981da1730da919181900360200190a150565b600d5490565b60075460ff1681565b606060068054806020026020016040519081016040528092919081815260200182805480156118e357602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116118c5575b5050505050905090565b60036020526000908152604090205460ff1681565b61190b33614487565b8061191557503330145b611959576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b611962826147c7565b6119a8576040805162461bcd60e51b8152602060048201526012602482015271746f6b656e206e6f74206c6f616461626c6560701b604482015290519081900360640190fd5b60006119b48383613ec7565b90506119c760148263ffffffff6147e116565b60006119d4604d54614857565b90506001600160a01b03841615611a7c576119ff6001600160a01b038516828563ffffffff61497616565b806001600160a01b0316631b3c96b485856040518363ffffffff1660e01b815260040180836001600160a01b03166001600160a01b0316815260200182815260200192505050600060405180830381600087803b158015611a5f57600080fd5b505af1158015611a73573d6000803e3d6000fd5b50505050611af6565b806001600160a01b0316631b3c96b48486866040518463ffffffff1660e01b815260040180836001600160a01b03166001600160a01b03168152602001828152602001925050506000604051808303818588803b158015611adc57600080fd5b505af1158015611af0573d6000803e3d6000fd5b50505050505b604080516001600160a01b03861681526020810185905281517f5f65674bec9af81f71be68674135a0ea3f163fb91984e3893d06da9f6ea2ce8a929181900390910190a150505050565b611b4933614487565b80611b5357503330145b611b97576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b601354811115611be9576040805162461bcd60e51b81526020600482015260186024820152771bdd5d081bd9881c985b99d9481b1bd85908185b5bdd5b9d60421b604482015290519081900360640190fd5b611bfa60148263ffffffff61449b16565b604080513381526020810183905281517f0b05243483e17c3f3377aee82b7d47e5700b48288695fc08b7ecc2759afa44ef929181900390910190a150565b611c4133614487565b80611c4b57503330145b611c8f576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b611ca060088263ffffffff61449b16565b604080513381526020810183905281517f068f112e5ec923d412be64779fe69e0fcbb6784c6617e94cccc8fd348f2e0f21929181900390910190a150565b6060611ce933614487565b80611cf357503330145b611d37576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b6001600160a01b03841660009081526003602052604090205460ff16611d6857611d6860088463ffffffff6147e116565b611d7a846001600160a01b0316614a8e565b8015611d8a5750611d8a84614a94565b15611f7157600080611d9c8685614aae565b6001600160a01b038216600090815260036020526040902054919350915060ff16611de2576000611dcd878361351f565b9050611de060088263ffffffff6147e116565b505b611dfb6001600160a01b0387168563ffffffff614bb816565b604080516020808252818301909252606091602082018180388339019050509050600160f81b81601f81518110611e2e57fe5b60200101906001600160f81b031916908160001a9053507ff77753fab406ecfff96d6ff2476c64a838fa9f6d37b1bf190f8546e395e3b6138787878460405180856001600160a01b03166001600160a01b031681526020018481526020018060200180602001838103835285818151815260200191508051906020019080838360005b83811015611ec9578181015183820152602001611eb1565b50505050905090810190601f168015611ef65780820380516001836020036101000a031916815260200191505b50838103825284518152845160209182019186019080838360005b83811015611f29578181015183820152602001611f11565b50505050905090810190601f168015611f565780820380516001836020036101000a031916815260200191505b50965050505050505060405180910390a192506121ce915050565b60006060856001600160a01b031685856040518082805190602001908083835b60208310611fb05780518252601f199092019160209182019101611f91565b6001836020036101000a03801982511681845116808217855250505050505090500191505060006040518083038185875af1925050503d8060008114612012576040519150601f19603f3d011682016040523d82523d6000602084013e612017565b606091505b50915091508181906120a75760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561206c578181015183820152602001612054565b50505050905090810190601f1680156120995780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b507ff77753fab406ecfff96d6ff2476c64a838fa9f6d37b1bf190f8546e395e3b6138686868460405180856001600160a01b03166001600160a01b031681526020018481526020018060200180602001838103835285818151815260200191508051906020019080838360005b8381101561212c578181015183820152602001612114565b50505050905090810190601f1680156121595780820380516001836020036101000a031916815260200191505b50838103825284518152845160209182019186019080838360005b8381101561218c578181015183820152602001612174565b50505050905090810190601f1680156121b95780820380516001836020036101000a031916815260200191505b50965050505050505060405180910390a19150505b9392505050565b6121de33614504565b61221d576040805162461bcd60e51b815260206004820152601a602482015260008051602061581a833981519152604482015290519081900360640190fd5b6000469050600061229d823089898960405160200180806836b7b737b634ba341d60b91b815250600901868152602001856001600160a01b03166001600160a01b031660601b8152601401848152602001838380828437808301925050509550505050505060405160208183030381529060405280519060200120614d76565b9050631626ba7e60e01b6001600160e01b0319166122f18286868080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506114ed92505050565b6001600160e01b0319161461233d576040805162461bcd60e51b815260206004820152600d60248201526c1cda59c81b9bdd081d985b1a59609a1b604482015290519081900360640190fd5b604c54871461237f576040805162461bcd60e51b81526020600482015260096024820152687478207265706c617960b81b604482015290519081900360640190fd5b612387614dc7565b60006060306001600160a01b03168888604051808383808284376040519201945060009350909150508083038183865af19150503d80600081146123e7576040519150601f19603f3d011682016040523d82523d6000602084013e6123ec565b606091505b509150915081819061243f5760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561206c578181015183820152602001612054565b507f823dbcf2b7b0f265871963ca65ac033f6b4c71e0d82cd123d2ff23d752dc21c188888360405180806020018060200183810383528686828181526020019250808284376000838201819052601f909101601f191690920185810384528651815286516020918201939188019250908190849084905b838110156124ce5781810151838201526020016124b6565b50505050905090810190601f1680156124fb5780820380516001836020036101000a031916815260200191505b509550505050505060405180910390a1505050505050505050565b606060058054806020026020016040519081016040528092919081815260200182805480156118e3576020028201919060005260206000209081546001600160a01b031681526001909101906020018083116118c5575050505050905090565b61257f33614504565b6125be576040805162461bcd60e51b815260206004820152601a602482015260008051602061581a833981519152604482015290519081900360640190fd5b60075460ff1661260d576040805162461bcd60e51b81526020600482015260156024820152743737903832b73234b7339039bab136b4b9b9b4b7b760591b604482015290519081900360640190fd5b61266e6005805480602002602001604051908101604052809291908181526020018280548015611431576020028201919060005260206000209081546001600160a01b03168152600190910190602001808311611413575050505050613ae0565b81146126ab5760405162461bcd60e51b81526004018080602001828103825260238152602001806158d66023913960400191505060405180910390fd5b60005b6005548110156127925760036000600583815481106126c957fe5b60009182526020808320909101546001600160a01b0316835282019290925260400190205460ff1661278a576001600360006005848154811061270857fe5b6000918252602080832091909101546001600160a01b031683528201929092526040019020805460ff191691151591909117905560058054600491908390811061274e57fe5b60009182526020808320909101548354600181018555938352912090910180546001600160a01b0319166001600160a01b039092169190911790555b6001016126ae565b507fb2f6cccee7a369e23e293c25aa19bef80af11eb26deba3ea0f2a02783f752e4a33600560405180836001600160a01b03166001600160a01b0316815260200180602001828103825283818154815260200191508054801561281e57602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311612800575b5050935050505060405180910390a161283960056000615718565b506007805460ff19169055565b60006116f46008614731565b61285b33614487565b8061286557503330145b6128a9576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b60075460ff161580156128c45750600754610100900460ff16155b612915576040805162461bcd60e51b815260206004820152601c60248201527f77686974656c6973742073756d62697373696f6e2070656e64696e6700000000604482015290519081900360640190fd5b60075462010000900460ff1661296e576040805162461bcd60e51b81526020600482015260196024820152781dda1a5d195b1a5cdd081b9bdd081a5b9a5d1a585b1a5e9959603a1b604482015290519081900360640190fd5b806129b2576040805162461bcd60e51b815260206004820152600f60248201526e195b5c1d1e481dda1a5d195b1a5cdd608a1b604482015290519081900360640190fd5b6129be60068383615736565b506007805461ff00191661010017905560408051602080840282810182019093528382527ffbc0e5ca6c7e4858daf0fdb185ef5186203e74ec9c64737e93c0aeaec596e1d19285928592612a2d92859185918291850190849080828437600092019190915250613ae092505050565b60405180806020018381526020018281038252858582818152602001925060200280828437600083820152604051601f909101601f1916909201829003965090945050505050a15050565b601954610100900460ff1680612a915750612a91614e0f565b80612a9f575060195460ff16155b612ada5760405162461bcd60e51b815260040180806020018281038252602e81526020018061585b602e913960400191505060405180910390fd5b601954610100900460ff16158015612b05576019805460ff1961ff0019909116610100171660011790555b612b0e86614e15565b612b1784614e80565b612b218888614e8f565b612b2a82614f4f565b612b32614f98565b612b3b85614fea565b604d8390558015612b52576019805461ff00191690555b5050505050505050565b612b6533614487565b612baf576040805162461bcd60e51b815260206004820152601660248201527539b2b73232b91034b9903737ba1030b71037bbb732b960511b604482015290519081900360640190fd5b600254600160a01b900460ff16612c0d576040805162461bcd60e51b815260206004820152601d60248201527f6f776e657273686970206973206e6f74207472616e7366657261626c65000000604482015290519081900360640190fd5b600280546001600160a01b0319169055604080516000808252602082015281517f850b3df64837d7d518b45f5aa64d104652c3b80eb5b34a8e3d9eb666cb7cdea5929181900390910190a1565b612c6333614487565b80612c6d57503330145b612cb1576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b8066038d7ea4c6800011158015612cd057506706f05b59d3b200008111155b612d17576040805162461bcd60e51b815260206004820152601360248201527206f7574206f662072616e676520746f702d757606c1b604482015290519081900360640190fd5b612d28600d8263ffffffff61476616565b6040805182815290517faf2a77cd04c3cc155588dd3bf67b310ab4fb3b1da3cf6b8d7d4d2aa1d09b794c9181900360200190a150565b604d5490565b6000546001600160a01b031690565b60006116f4600d614731565b612d8833614487565b80612d9257503330145b612dd6576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b60075460ff16158015612df15750600754610100900460ff16155b612e42576040805162461bcd60e51b815260206004820152601c60248201527f77686974656c6973742073756d62697373696f6e2070656e64696e6700000000604482015290519081900360640190fd5b8181808060200260200160405190810160405280939291908181526020018383602002808284376000920182905250925050505b8151811015612f5e57612e9b828281518110612e8e57fe5b6020026020010151614487565b15612ee6576040805162461bcd60e51b8152602060048201526016602482015275636f6e7461696e73206f776e6572206164647265737360501b604482015290519081900360640190fd5b60006001600160a01b0316828281518110612efd57fe5b60200260200101516001600160a01b03161415612f56576040805162461bcd60e51b8152602060048201526012602482015271636f6e7461696e732030206164647265737360701b604482015290519081900360640190fd5b600101612e76565b5060075462010000900460ff16612fb8576040805162461bcd60e51b81526020600482015260196024820152781dda1a5d195b1a5cdd081b9bdd081a5b9a5d1a585b1a5e9959603a1b604482015290519081900360640190fd5b81612ffc576040805162461bcd60e51b815260206004820152600f60248201526e195b5c1d1e481dda1a5d195b1a5cdd608a1b604482015290519081900360640190fd5b61300860058484615736565b506007805460ff1916600117905560408051602080850282810182019093528482527f9c80b3b5f68b3e017766d59e8d09b34efe6462b05c398f35cab9e271d9bc3b9c928692869261307592859185918291850190849080828437600092019190915250613ae092505050565b60405180806020018381526020018281038252858582818152602001925060200280828437600083820152604051601f909101601f1916909201829003965090945050505050a1505050565b60125490565b6002546001600160a01b031690565b6130df33614487565b613129576040805162461bcd60e51b815260206004820152601660248201527539b2b73232b91034b9903737ba1030b71037bbb732b960511b604482015290519081900360640190fd5b613131614dc7565b565b60085490565b61314233614487565b8061314c57503330145b613190576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b8051602080820191906000808060605b86851015612b52576131b986605463ffffffff61509d16565b888601805160148201516034909201805193995060609190911c965090945090925090506131fe60546131f2878563ffffffff6150fa16565b9063ffffffff6150fa16565b945086851115613245576040805162461bcd60e51b815260206004820152600d60248201526c6f7574206f6620626f756e647360981b604482015290519081900360640190fd5b8161325b57506040805160208101909152600081525b613266848483611cde565b506131a0565b60185460ff1690565b61327e33614487565b8061328857503330145b6132cc576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b60135481111561331e576040805162461bcd60e51b81526020600482015260186024820152771bdd5d081bd9881c985b99d9481b1bd85908185b5bdd5b9d60421b604482015290519081900360640190fd5b61332f60148263ffffffff61476616565b6040805182815290517fc178d379965e5657b6fc57494e392f121a14119215dfb422aad7db4cc03f2d109181900360200190a150565b61336e33614487565b6133b8576040805162461bcd60e51b815260206004820152601660248201527539b2b73232b91034b9903737ba1030b71037bbb732b960511b604482015290519081900360640190fd5b600254600160a01b900460ff16613416576040805162461bcd60e51b815260206004820152601d60248201527f6f776e657273686970206973206e6f74207472616e7366657261626c65000000604482015290519081900360640190fd5b6001600160a01b03821661345b5760405162461bcd60e51b81526004018080602001828103825260238152602001806158896023913960400191505060405180910390fd5b6002805460ff60a01b1916600160a01b83151502179055806134b457604080516001600160a01b038416815290517f808639ff9c8e4732d60b6c2330de498035416d229f27a77d259680895efec1229181900360200190a15b600254604080516001600160a01b039283168152918416602083015280517f850b3df64837d7d518b45f5aa64d104652c3b80eb5b34a8e3d9eb666cb7cdea59281900390910190a150600280546001600160a01b0319166001600160a01b0392909216919091179055565b60008060008061352e86615154565b5050509350935093505080156135a35781613579576040805162461bcd60e51b81526020600482015260066024820152650726174653d360d41b604482015290519081900360640190fd5b6135998361358d878563ffffffff6152e616565b9063ffffffff61533f16565b935050505061155c565b50600095945050505050565b600c5460ff1690565b60075462010000900460ff1681565b6135d033614487565b806135da57503330145b61361e576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b808061365b576040805162461bcd60e51b8152602060048201526007602482015266076616c75653d360cc1b604482015290519081900360640190fd5b6001600160a01b0384166136a6576040805162461bcd60e51b815260206004820152600d60248201526c064657374696e6174696f6e3d3609c1b604482015290519081900360640190fd5b6001600160a01b03841660009081526003602052604090205460ff166136f657816001600160a01b038416156136e3576136e0848461351f565b90505b6136f460088263ffffffff6147e116565b505b6137018484846153a9565b604080516001600160a01b0380871682528516602082015280820184905290517fd1ba4ac2e2a11b5101f6cb4d978f514a155b421e8ec396d2d9abaf0bb02917ee9181900360600190a150505050565b60175490565b61376033614504565b61379f576040805162461bcd60e51b815260206004820152601a602482015260008051602061581a833981519152604482015290519081900360640190fd5b600754610100900460ff166137f3576040805162461bcd60e51b81526020600482015260156024820152743737903832b73234b7339039bab136b4b9b9b4b7b760591b604482015290519081900360640190fd5b6138546006805480602002602001604051908101604052809291908181526020018280548015611431576020028201919060005260206000209081546001600160a01b03168152600190910190602001808311611413575050505050613ae0565b81146138915760405162461bcd60e51b81526004018080602001828103825260238152602001806158d66023913960400191505060405180910390fd5b60005b600654811015613a1f5760036000600683815481106138af57fe5b60009182526020808320909101546001600160a01b0316835282019290925260400190205460ff1615613a1757600060036000600684815481106138ef57fe5b6000918252602080832091909101546001600160a01b031683528201929092526040018120805460ff1916921515929092179091555b60045461393990600163ffffffff61509d16565b811015613a01576006828154811061394d57fe5b600091825260209091200154600480546001600160a01b03909216918390811061397357fe5b6000918252602090912001546001600160a01b031614156139f9576004805460001981019081106139a057fe5b600091825260209091200154600480546001600160a01b0390921691839081106139c657fe5b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b03160217905550613a01565b600101613925565b506004805490613a15906000198301615799565b505b600101613894565b507fd218c430fa348f4ce67791021b6b89c0c3eacd4ead1d8f5b83c60038ec28249b33600660405180836001600160a01b03166001600160a01b03168152602001806020018281038252838181548152602001915080548015613aab57602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311613a8d575b5050935050505060405180910390a1613ac660066000615718565b506007805461ff0019169055565b60105490565b604c5481565b60008160405160200180828051906020019060200280838360005b83811015613b13578181015183820152602001613afb565b50505050905001915050604051602081830303815290604052805190602001209050919050565b613b4333614487565b80613b525750613b5233614504565b613b9c576040805162461bcd60e51b815260206004820152601660248201527537b7363c9037bbb732b93e3e31b7b73a3937b63632b960511b604482015290519081900360640190fd5b600754610100900460ff16613bf0576040805162461bcd60e51b81526020600482015260156024820152743737903832b73234b7339039bab136b4b9b9b4b7b760591b604482015290519081900360640190fd5b613c516006805480602002602001604051908101604052809291908181526020018280548015611431576020028201919060005260206000209081546001600160a01b03168152600190910190602001808311611413575050505050613ae0565b8114613c8e5760405162461bcd60e51b81526004018080602001828103825260238152602001806158d66023913960400191505060405180910390fd5b613c9a60066000615718565b6007805461ff0019169055604080513381526020810183905281517f13c935eb475aa0f6e931fece83e2ac44569ce2d53460d29a6dedab40b965c8a3929181900390910190a150565b60048181548110613cf057fe5b6000918252602090912001546001600160a01b0316905081565b60145490565b600754610100900460ff1681565b60015490565b8080613d61576040805162461bcd60e51b8152602060048201526007602482015266076616c75653d360cc1b604482015290519081900360640190fd5b613d6a33614487565b80613d795750613d7933614504565b613dc3576040805162461bcd60e51b815260206004820152601660248201527537b7363c9037bbb732b93e3e31b7b73a3937b63632b960511b604482015290519081900360640190fd5b613dd4600d8363ffffffff6147e116565b613ddc6130c7565b6001600160a01b03166108fc839081150290604051600060405180830381858888f19350505050158015613e14573d6000803e3d6000fd5b507f611b7c0d84fda988026215bef9b3e4d81cbceced7e679be6d5e044b588467c0e33613e3f6130c7565b604080516001600160a01b03938416815291909216602082015280820185905290519081900360600190a15050565b613e7733614504565b613eb6576040805162461bcd60e51b815260206004820152601a602482015260008051602061581a833981519152604482015290519081900360640190fd5b611ca060088263ffffffff61540d16565b6000613ed1615461565b6001600160a01b0316836001600160a01b03161415613ef157508061155c565b816001600160a01b03841615613fb6576000806000613f0f87615154565b5050509350935093505080613f61576040805162461bcd60e51b8152602060048201526013602482015272746f6b656e206e6f7420617661696c61626c6560681b604482015290519081900360640190fd5b81613f9c576040805162461bcd60e51b81526020600482015260066024820152650726174653d360d41b604482015290519081900360640190fd5b613fb08361358d888563ffffffff6152e616565b93505050505b6000806000613fc36154d7565b5050509350935093505080614015576040805162461bcd60e51b8152602060048201526013602482015272746f6b656e206e6f7420617661696c61626c6560681b604482015290519081900360640190fd5b8161405b576040805162461bcd60e51b81526020600482015260116024820152700737461626c65636f696e20726174653d3607c1b604482015290519081900360640190fd5b61406f8261358d868663ffffffff6152e616565b979650505050505050565b61408333614504565b6140c2576040805162461bcd60e51b815260206004820152601a602482015260008051602061581a833981519152604482015290519081900360640190fd5b611bfa60148263ffffffff61540d16565b6140dc33614504565b61411b576040805162461bcd60e51b815260206004820152601a602482015260008051602061581a833981519152604482015290519081900360640190fd5b6112e9600d8263ffffffff61540d16565b61413533614487565b8061413f57503330145b614183576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037bbb732b93e3e39b2b63360811b604482015290519081900360640190fd5b8181808060200260200160405190810160405280939291908181526020018383602002808284376000920182905250925050505b8151811015614292576141cf828281518110612e8e57fe5b1561421a576040805162461bcd60e51b8152602060048201526016602482015275636f6e7461696e73206f776e6572206164647265737360501b604482015290519081900360640190fd5b60006001600160a01b031682828151811061423157fe5b60200260200101516001600160a01b0316141561428a576040805162461bcd60e51b8152602060048201526012602482015271636f6e7461696e732030206164647265737360701b604482015290519081900360640190fd5b6001016141b7565b5060075462010000900460ff16156142e9576040805162461bcd60e51b81526020600482015260156024820152741dda1a5d195b1a5cdd081a5b9a5d1a585b1a5e9959605a1b604482015290519081900360640190fd5b60005b828110156143da576003600085858481811061430457fe5b602090810292909201356001600160a01b03168352508101919091526040016000205460ff166143d25760016003600086868581811061434057fe5b905060200201356001600160a01b03166001600160a01b03166001600160a01b0316815260200190815260200160002060006101000a81548160ff021916908315150217905550600484848381811061439557fe5b835460018101855560009485526020948590200180546001600160a01b0319166001600160a01b0395909202939093013593909316929092179055505b6001016142ec565b506007805462ff0000191662010000179055604080513380825260208201838152600480549484018590527fb2f6cccee7a369e23e293c25aa19bef80af11eb26deba3ea0f2a02783f752e4a949293909290919060608301908490801561446a57602002820191906000526020600020905b81546001600160a01b0316815260019091019060200180831161444c575b5050935050505060405180910390a1505050565b60115460ff1690565b6002546001600160a01b0390811691161490565b600482015460ff16156144e9576040805162461bcd60e51b81526020600482015260116024820152701b1a5b5a5d08185b1c9958591e481cd95d607a1b604482015290519081900360640190fd5b6144f38282615646565b50600401805460ff19166001179055565b6000614511600154614857565b6001600160a01b031663b429afeb836040518263ffffffff1660e01b815260040180826001600160a01b03166001600160a01b0316815260200191505060206040518083038186803b15801561456657600080fd5b505afa15801561457a573d6000803e3d6000fd5b505050506040513d602081101561459057600080fd5b505192915050565b600081516041146145ab5750600061155c565b60208201516040830151606084015160001a7f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08211156145f1576000935050505061155c565b8060ff16601b1415801561460957508060ff16601c14155b1561461a576000935050505061155c565b6040805160008152602080820180845289905260ff8416828401526060820186905260808201859052915160019260a0808401939192601f1981019281900390910190855afa158015614671573d6000803e3d6000fd5b5050604051601f190151979650505050505050565b60006001600160a01b0382161561472057816001600160a01b03166370a08231846040518263ffffffff1660e01b815260040180826001600160a01b03166001600160a01b0316815260200191505060206040518083038186803b1580156146ed57600080fd5b505afa158015614701573d6000803e3d6000fd5b505050506040513d602081101561471757600080fd5b5051905061155c565b506001600160a01b0382163161155c565b600281015460009061474c906201518063ffffffff6150fa16565b42111561475b5750805461120f565b50600181015461120f565b600482015460ff166147bf576040805162461bcd60e51b815260206004820152601960248201527f6c696d6974206861736e2774206265656e207365742079657400000000000000604482015290519081900360640190fd5b600390910155565b6000806147d383615154565b509098975050505050505050565b6147ea82615669565b8082600101541015614836576040805162461bcd60e51b815260206004820152601060248201526f185d985a5b18589b194f185b5bdd5b9d60821b604482015290519081900360640190fd5b600182015461484b908263ffffffff61509d16565b82600101819055505050565b600080546001600160a01b03166148b5576040805162461bcd60e51b815260206004820152601d60248201527f454e535265736f6c7661626c65206e6f7420696e697469616c697a6564000000604482015290519081900360640190fd5b60005460408051630178b8bf60e01b81526004810185905290516001600160a01b0390921691630178b8bf91602480820192602092909190829003018186803b15801561490157600080fd5b505afa158015614915573d6000803e3d6000fd5b505050506040513d602081101561492b57600080fd5b505160408051631d9dabef60e11b81526004810185905290516001600160a01b0390921691633b3b57de91602480820192602092909190829003018186803b15801561456657600080fd5b8015806149fc575060408051636eb1769f60e11b81523060048201526001600160a01b03848116602483015291519185169163dd62ed3e91604480820192602092909190829003018186803b1580156149ce57600080fd5b505afa1580156149e2573d6000803e3d6000fd5b505050506040513d60208110156149f857600080fd5b5051155b614a375760405162461bcd60e51b81526004018080602001828103825260368152602001806158f96036913960400191505060405180910390fd5b604080516001600160a01b038416602482015260448082018490528251808303909101815260649091019091526020810180516001600160e01b031663095ea7b360e01b179052614a89908490614bb8565b505050565b3b151590565b600080614aa083615154565b509198975050505050505050565b600080614abc601254614857565b6001600160a01b031663afc72e9385856040518363ffffffff1660e01b815260040180836001600160a01b03166001600160a01b0316815260200180602001828103825283818151815260200191508051906020019080838360005b83811015614b30578181015183820152602001614b18565b50505050905090810190601f168015614b5d5780820380516001836020036101000a031916815260200191505b509350505050604080518083038186803b158015614b7a57600080fd5b505afa158015614b8e573d6000803e3d6000fd5b505050506040513d6040811015614ba457600080fd5b508051602090910151909590945092505050565b614bca826001600160a01b0316614a8e565b614c1b576040805162461bcd60e51b815260206004820152601f60248201527f5361666545524332303a2063616c6c20746f206e6f6e2d636f6e747261637400604482015290519081900360640190fd5b60006060836001600160a01b0316836040518082805190602001908083835b60208310614c595780518252601f199092019160209182019101614c3a565b6001836020036101000a0380198251168184511680821785525050505050509050019150506000604051808303816000865af19150503d8060008114614cbb576040519150601f19603f3d011682016040523d82523d6000602084013e614cc0565b606091505b509150915081614d17576040805162461bcd60e51b815260206004820181905260248201527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564604482015290519081900360640190fd5b805115614d7057808060200190516020811015614d3357600080fd5b5051614d705760405162461bcd60e51b815260040180806020018281038252602a8152602001806158ac602a913960400191505060405180910390fd5b50505050565b604080517f19457468657265756d205369676e6564204d6573736167653a0a333200000000602080830191909152603c8083019490945282518083039094018452605c909101909152815191012090565b604c80546002019081905560408051338152602081019290925280517fab0423a75986556234aecd171c46ce7f5e45607d8070bf5230f2735b50322bff9281900390910190a1565b303b1590565b6001600160a01b038116614e5e576040805162461bcd60e51b815260206004820152600b60248201526a0656e7352656720697320360ac1b604482015290519081900360640190fd5b600080546001600160a01b0319166001600160a01b0392909216919091179055565b8015614e8c5760018190555b50565b600280546001600160a01b0319166001600160a01b0384161760ff60a01b1916600160a01b8315158102919091179182905560ff910416614f0757604080516001600160a01b038416815290517f808639ff9c8e4732d60b6c2330de498035416d229f27a77d259680895efec1229181900360200190a15b60408051600081526001600160a01b038416602082015281517f850b3df64837d7d518b45f5aa64d104652c3b80eb5b34a8e3d9eb666cb7cdea5929181900390910190a15050565b6040805160a081018252828152602081018390524291810182905260006060820181905260809091018190526008839055600992909255600a55600b55600c805460ff19169055565b6040805160a0810182526706f05b59d3b2000080825260208201819052429282018390526000606083018190526080909201829052600d819055600e55600f919091556010556011805460ff19169055565b614ff3816156c1565b6000614ffd6154d7565b50505050509150506000811161504a576040805162461bcd60e51b815260206004820152600d60248201526c37379039ba30b13632b1b7b4b760991b604482015290519081900360640190fd5b6127100260138190556040805160a081018252828152602081018390524291810182905260006060820181905260809091018190526014839055601592909255601655601755506018805460ff19169055565b6000828211156150f4576040805162461bcd60e51b815260206004820152601e60248201527f536166654d6174683a207375627472616374696f6e206f766572666c6f770000604482015290519081900360640190fd5b50900390565b6000828201838110156121ce576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b606060008060008060008061516a601254614857565b6001600160a01b0316631f69565f896040518263ffffffff1660e01b815260040180826001600160a01b03166001600160a01b0316815260200191505060006040518083038186803b1580156151bf57600080fd5b505afa1580156151d3573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405260e08110156151fc57600080fd5b8101908080516040519392919084600160201b82111561521b57600080fd5b90830190602082018581111561523057600080fd5b8251600160201b81118282018810171561524957600080fd5b82525081516020918201929091019080838360005b8381101561527657818101518382015260200161525e565b50505050905090810190601f1680156152a35780820380516001836020036101000a031916815260200191505b5060409081526020820151908201516060830151608084015160a085015160c090950151979e50929c50909a509850965094509192505050919395979092949650565b6000826152f55750600061155c565b8282028284828161530257fe5b04146121ce5760405162461bcd60e51b815260040180806020018281038252602181526020018061583a6021913960400191505060405180910390fd5b6000808211615395576040805162461bcd60e51b815260206004820152601a60248201527f536166654d6174683a206469766973696f6e206279207a65726f000000000000604482015290519081900360640190fd5b60008284816153a057fe5b04949350505050565b6001600160a01b0382166153f3576040516001600160a01b0384169082156108fc029083906000818181858888f193505050501580156153ed573d6000803e3d6000fd5b50614a89565b614a896001600160a01b038316848363ffffffff6156c616565b8082600301541461544f5760405162461bcd60e51b81526004018080602001828103825260228152602001806157f86022913960400191505060405180910390fd5b61545d828360030154615646565b5050565b600061546e601254614857565b6001600160a01b031663e9cbd8226040518163ffffffff1660e01b815260040160206040518083038186803b1580156154a657600080fd5b505afa1580156154ba573d6000803e3d6000fd5b505050506040513d60208110156154d057600080fd5b5051905090565b60606000806000806000806154ed601254614857565b6001600160a01b0316633efec5e96040518163ffffffff1660e01b815260040160006040518083038186803b15801561552557600080fd5b505afa158015615539573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405260e081101561556257600080fd5b8101908080516040519392919084600160201b82111561558157600080fd5b90830190602082018581111561559657600080fd5b8251600160201b8111828201881017156155af57600080fd5b82525081516020918201929091019080838360005b838110156155dc5781810151838201526020016155c4565b50505050905090810190601f1680156156095780820380516001836020036101000a031916815260200191505b5060409081526020820151908201516060830151608084015160a085015160c090950151979f939e50919c509a5098509096509294509192505050565b61564f82615669565b808255600182015481101561545d57815460018301555050565b6002810154615681906201518063ffffffff6150fa16565b421115614e8c57426002820155805460018201556040517fe93bc25276d408d390778e7a8b926f2f67209c43ed540081b951fe128f0d3cd290600090a150565b601255565b604080516001600160a01b038416602482015260448082018490528251808303909101815260649091019091526020810180516001600160e01b031663a9059cbb60e01b179052614a89908490614bb8565b5080546000825590600052602060002090810190614e8c91906157b9565b828054828255906000526020600020908101928215615789579160200282015b828111156157895781546001600160a01b0319166001600160a01b03843516178255602090920191600190910190615756565b506157959291506157d3565b5090565b815481835581811115614a8957600083815260209020614a899181019083015b61121891905b8082111561579557600081556001016157bf565b61121891905b808211156157955780546001600160a01b03191681556001016157d956fe636f6e6669726d65642f7375626d6974746564206c696d6974206d69736d6174636873656e646572206973206e6f74206120636f6e74726f6c6c6572000000000000536166654d6174683a206d756c7469706c69636174696f6e206f766572666c6f77436f6e747261637420696e7374616e63652068617320616c7265616479206265656e20696e697469616c697a65646f776e65722063616e6e6f742062652073657420746f207a65726f20616464726573735361666545524332303a204552433230206f7065726174696f6e20646964206e6f7420737563636565646e6f6e2d6d61746368696e672070656e64696e672077686974656c69737420686173685361666545524332303a20617070726f76652066726f6d206e6f6e2d7a65726f20746f206e6f6e2d7a65726f20616c6c6f77616e6365a265627a7a72315820828a21ea8ff2f0cf794a63717e5d87d2d8e9c76120811b43d647df092f4afe7a64736f6c634300050f0032"

// DeployWalletMock deploys a new Ethereum contract, binding an instance of WalletMock to it.
func DeployWalletMock(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *WalletMock, error) {
	parsed, err := abi.JSON(strings.NewReader(WalletMockABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(WalletMockBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &WalletMock{WalletMockCaller: WalletMockCaller{contract: contract}, WalletMockTransactor: WalletMockTransactor{contract: contract}, WalletMockFilterer: WalletMockFilterer{contract: contract}}, nil
}

// WalletMock is an auto generated Go binding around an Ethereum contract.
type WalletMock struct {
	WalletMockCaller     // Read-only binding to the contract
	WalletMockTransactor // Write-only binding to the contract
	WalletMockFilterer   // Log filterer for contract events
}

// WalletMockCaller is an auto generated read-only Go binding around an Ethereum contract.
type WalletMockCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletMockTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WalletMockTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletMockFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WalletMockFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletMockSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WalletMockSession struct {
	Contract     *WalletMock       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WalletMockCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WalletMockCallerSession struct {
	Contract *WalletMockCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// WalletMockTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WalletMockTransactorSession struct {
	Contract     *WalletMockTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// WalletMockRaw is an auto generated low-level Go binding around an Ethereum contract.
type WalletMockRaw struct {
	Contract *WalletMock // Generic contract binding to access the raw methods on
}

// WalletMockCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WalletMockCallerRaw struct {
	Contract *WalletMockCaller // Generic read-only contract binding to access the raw methods on
}

// WalletMockTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WalletMockTransactorRaw struct {
	Contract *WalletMockTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWalletMock creates a new instance of WalletMock, bound to a specific deployed contract.
func NewWalletMock(address common.Address, backend bind.ContractBackend) (*WalletMock, error) {
	contract, err := bindWalletMock(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WalletMock{WalletMockCaller: WalletMockCaller{contract: contract}, WalletMockTransactor: WalletMockTransactor{contract: contract}, WalletMockFilterer: WalletMockFilterer{contract: contract}}, nil
}

// NewWalletMockCaller creates a new read-only instance of WalletMock, bound to a specific deployed contract.
func NewWalletMockCaller(address common.Address, caller bind.ContractCaller) (*WalletMockCaller, error) {
	contract, err := bindWalletMock(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WalletMockCaller{contract: contract}, nil
}

// NewWalletMockTransactor creates a new write-only instance of WalletMock, bound to a specific deployed contract.
func NewWalletMockTransactor(address common.Address, transactor bind.ContractTransactor) (*WalletMockTransactor, error) {
	contract, err := bindWalletMock(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WalletMockTransactor{contract: contract}, nil
}

// NewWalletMockFilterer creates a new log filterer instance of WalletMock, bound to a specific deployed contract.
func NewWalletMockFilterer(address common.Address, filterer bind.ContractFilterer) (*WalletMockFilterer, error) {
	contract, err := bindWalletMock(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WalletMockFilterer{contract: contract}, nil
}

// bindWalletMock binds a generic wrapper to an already deployed contract.
func bindWalletMock(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(WalletMockABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WalletMock *WalletMockRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _WalletMock.Contract.WalletMockCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WalletMock *WalletMockRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WalletMock.Contract.WalletMockTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WalletMock *WalletMockRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WalletMock.Contract.WalletMockTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WalletMock *WalletMockCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _WalletMock.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WalletMock *WalletMockTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WalletMock.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WalletMock *WalletMockTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WalletMock.Contract.contract.Transact(opts, method, params...)
}

// WALLETVERSION is a free data retrieval call binding the contract method 0x1127b57e.
//
// Solidity: function WALLET_VERSION() constant returns(string)
func (_WalletMock *WalletMockCaller) WALLETVERSION(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _WalletMock.contract.Call(opts, out, "WALLET_VERSION")
	return *ret0, err
}

// WALLETVERSION is a free data retrieval call binding the contract method 0x1127b57e.
//
// Solidity: function WALLET_VERSION() constant returns(string)
func (_WalletMock *WalletMockSession) WALLETVERSION() (string, error) {
	return _WalletMock.Contract.WALLETVERSION(&_WalletMock.CallOpts)
}

// WALLETVERSION is a free data retrieval call binding the contract method 0x1127b57e.
//
// Solidity: function WALLET_VERSION() constant returns(string)
func (_WalletMock *WalletMockCallerSession) WALLETVERSION() (string, error) {
	return _WalletMock.Contract.WALLETVERSION(&_WalletMock.CallOpts)
}

// CalculateHash is a free data retrieval call binding the contract method 0xcd7958dd.
//
// Solidity: function calculateHash(address[] _addresses) constant returns(bytes32)
func (_WalletMock *WalletMockCaller) CalculateHash(opts *bind.CallOpts, _addresses []common.Address) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _WalletMock.contract.Call(opts, out, "calculateHash", _addresses)
	return *ret0, err
}

// CalculateHash is a free data retrieval call binding the contract method 0xcd7958dd.
//
// Solidity: function calculateHash(address[] _addresses) constant returns(bytes32)
func (_WalletMock *WalletMockSession) CalculateHash(_addresses []common.Address) ([32]byte, error) {
	return _WalletMock.Contract.CalculateHash(&_WalletMock.CallOpts, _addresses)
}

// CalculateHash is a free data retrieval call binding the contract method 0xcd7958dd.
//
// Solidity: function calculateHash(address[] _addresses) constant returns(bytes32)
func (_WalletMock *WalletMockCallerSession) CalculateHash(_addresses []common.Address) ([32]byte, error) {
	return _WalletMock.Contract.CalculateHash(&_WalletMock.CallOpts, _addresses)
}

// ControllerNode is a free data retrieval call binding the contract method 0xe2b4ce97.
//
// Solidity: function controllerNode() constant returns(bytes32)
func (_WalletMock *WalletMockCaller) ControllerNode(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _WalletMock.contract.Call(opts, out, "controllerNode")
	return *ret0, err
}

// ControllerNode is a free data retrieval call binding the contract method 0xe2b4ce97.
//
// Solidity: function controllerNode() constant returns(bytes32)
func (_WalletMock *WalletMockSession) ControllerNode() ([32]byte, error) {
	return _WalletMock.Contract.ControllerNode(&_WalletMock.CallOpts)
}

// ControllerNode is a free data retrieval call binding the contract method 0xe2b4ce97.
//
// Solidity: function controllerNode() constant returns(bytes32)
func (_WalletMock *WalletMockCallerSession) ControllerNode() ([32]byte, error) {
	return _WalletMock.Contract.ControllerNode(&_WalletMock.CallOpts)
}

// ConvertToEther is a free data retrieval call binding the contract method 0xb87e21ef.
//
// Solidity: function convertToEther(address _token, uint256 _amount) constant returns(uint256)
func (_WalletMock *WalletMockCaller) ConvertToEther(opts *bind.CallOpts, _token common.Address, _amount *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _WalletMock.contract.Call(opts, out, "convertToEther", _token, _amount)
	return *ret0, err
}

// ConvertToEther is a free data retrieval call binding the contract method 0xb87e21ef.
//
// Solidity: function convertToEther(address _token, uint256 _amount) constant returns(uint256)
func (_WalletMock *WalletMockSession) ConvertToEther(_token common.Address, _amount *big.Int) (*big.Int, error) {
	return _WalletMock.Contract.ConvertToEther(&_WalletMock.CallOpts, _token, _amount)
}

// ConvertToEther is a free data retrieval call binding the contract method 0xb87e21ef.
//
// Solidity: function convertToEther(address _token, uint256 _amount) constant returns(uint256)
func (_WalletMock *WalletMockCallerSession) ConvertToEther(_token common.Address, _amount *big.Int) (*big.Int, error) {
	return _WalletMock.Contract.ConvertToEther(&_WalletMock.CallOpts, _token, _amount)
}

// ConvertToStablecoin is a free data retrieval call binding the contract method 0xf36febda.
//
// Solidity: function convertToStablecoin(address _token, uint256 _amount) constant returns(uint256)
func (_WalletMock *WalletMockCaller) ConvertToStablecoin(opts *bind.CallOpts, _token common.Address, _amount *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _WalletMock.contract.Call(opts, out, "convertToStablecoin", _token, _amount)
	return *ret0, err
}

// ConvertToStablecoin is a free data retrieval call binding the contract method 0xf36febda.
//
// Solidity: function convertToStablecoin(address _token, uint256 _amount) constant returns(uint256)
func (_WalletMock *WalletMockSession) ConvertToStablecoin(_token common.Address, _amount *big.Int) (*big.Int, error) {
	return _WalletMock.Contract.ConvertToStablecoin(&_WalletMock.CallOpts, _token, _amount)
}

// ConvertToStablecoin is a free data retrieval call binding the contract method 0xf36febda.
//
// Solidity: function convertToStablecoin(address _token, uint256 _amount) constant returns(uint256)
func (_WalletMock *WalletMockCallerSession) ConvertToStablecoin(_token common.Address, _amount *big.Int) (*big.Int, error) {
	return _WalletMock.Contract.ConvertToStablecoin(&_WalletMock.CallOpts, _token, _amount)
}

// EnsRegistry is a free data retrieval call binding the contract method 0x7d73b231.
//
// Solidity: function ensRegistry() constant returns(address)
func (_WalletMock *WalletMockCaller) EnsRegistry(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _WalletMock.contract.Call(opts, out, "ensRegistry")
	return *ret0, err
}

// EnsRegistry is a free data retrieval call binding the contract method 0x7d73b231.
//
// Solidity: function ensRegistry() constant returns(address)
func (_WalletMock *WalletMockSession) EnsRegistry() (common.Address, error) {
	return _WalletMock.Contract.EnsRegistry(&_WalletMock.CallOpts)
}

// EnsRegistry is a free data retrieval call binding the contract method 0x7d73b231.
//
// Solidity: function ensRegistry() constant returns(address)
func (_WalletMock *WalletMockCallerSession) EnsRegistry() (common.Address, error) {
	return _WalletMock.Contract.EnsRegistry(&_WalletMock.CallOpts)
}

// GasTopUpLimitAvailable is a free data retrieval call binding the contract method 0x7d7d0046.
//
// Solidity: function gasTopUpLimitAvailable() constant returns(uint256)
func (_WalletMock *WalletMockCaller) GasTopUpLimitAvailable(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _WalletMock.contract.Call(opts, out, "gasTopUpLimitAvailable")
	return *ret0, err
}

// GasTopUpLimitAvailable is a free data retrieval call binding the contract method 0x7d7d0046.
//
// Solidity: function gasTopUpLimitAvailable() constant returns(uint256)
func (_WalletMock *WalletMockSession) GasTopUpLimitAvailable() (*big.Int, error) {
	return _WalletMock.Contract.GasTopUpLimitAvailable(&_WalletMock.CallOpts)
}

// GasTopUpLimitAvailable is a free data retrieval call binding the contract method 0x7d7d0046.
//
// Solidity: function gasTopUpLimitAvailable() constant returns(uint256)
func (_WalletMock *WalletMockCallerSession) GasTopUpLimitAvailable() (*big.Int, error) {
	return _WalletMock.Contract.GasTopUpLimitAvailable(&_WalletMock.CallOpts)
}

// GasTopUpLimitControllerConfirmationRequired is a free data retrieval call binding the contract method 0xf776f518.
//
// Solidity: function gasTopUpLimitControllerConfirmationRequired() constant returns(bool)
func (_WalletMock *WalletMockCaller) GasTopUpLimitControllerConfirmationRequired(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _WalletMock.contract.Call(opts, out, "gasTopUpLimitControllerConfirmationRequired")
	return *ret0, err
}

// GasTopUpLimitControllerConfirmationRequired is a free data retrieval call binding the contract method 0xf776f518.
//
// Solidity: function gasTopUpLimitControllerConfirmationRequired() constant returns(bool)
func (_WalletMock *WalletMockSession) GasTopUpLimitControllerConfirmationRequired() (bool, error) {
	return _WalletMock.Contract.GasTopUpLimitControllerConfirmationRequired(&_WalletMock.CallOpts)
}

// GasTopUpLimitControllerConfirmationRequired is a free data retrieval call binding the contract method 0xf776f518.
//
// Solidity: function gasTopUpLimitControllerConfirmationRequired() constant returns(bool)
func (_WalletMock *WalletMockCallerSession) GasTopUpLimitControllerConfirmationRequired() (bool, error) {
	return _WalletMock.Contract.GasTopUpLimitControllerConfirmationRequired(&_WalletMock.CallOpts)
}

// GasTopUpLimitPending is a free data retrieval call binding the contract method 0xcc0e7e56.
//
// Solidity: function gasTopUpLimitPending() constant returns(uint256)
func (_WalletMock *WalletMockCaller) GasTopUpLimitPending(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _WalletMock.contract.Call(opts, out, "gasTopUpLimitPending")
	return *ret0, err
}

// GasTopUpLimitPending is a free data retrieval call binding the contract method 0xcc0e7e56.
//
// Solidity: function gasTopUpLimitPending() constant returns(uint256)
func (_WalletMock *WalletMockSession) GasTopUpLimitPending() (*big.Int, error) {
	return _WalletMock.Contract.GasTopUpLimitPending(&_WalletMock.CallOpts)
}

// GasTopUpLimitPending is a free data retrieval call binding the contract method 0xcc0e7e56.
//
// Solidity: function gasTopUpLimitPending() constant returns(uint256)
func (_WalletMock *WalletMockCallerSession) GasTopUpLimitPending() (*big.Int, error) {
	return _WalletMock.Contract.GasTopUpLimitPending(&_WalletMock.CallOpts)
}

// GasTopUpLimitValue is a free data retrieval call binding the contract method 0x2587a6a2.
//
// Solidity: function gasTopUpLimitValue() constant returns(uint256)
func (_WalletMock *WalletMockCaller) GasTopUpLimitValue(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _WalletMock.contract.Call(opts, out, "gasTopUpLimitValue")
	return *ret0, err
}

// GasTopUpLimitValue is a free data retrieval call binding the contract method 0x2587a6a2.
//
// Solidity: function gasTopUpLimitValue() constant returns(uint256)
func (_WalletMock *WalletMockSession) GasTopUpLimitValue() (*big.Int, error) {
	return _WalletMock.Contract.GasTopUpLimitValue(&_WalletMock.CallOpts)
}

// GasTopUpLimitValue is a free data retrieval call binding the contract method 0x2587a6a2.
//
// Solidity: function gasTopUpLimitValue() constant returns(uint256)
func (_WalletMock *WalletMockCallerSession) GasTopUpLimitValue() (*big.Int, error) {
	return _WalletMock.Contract.GasTopUpLimitValue(&_WalletMock.CallOpts)
}

// IsSetWhitelist is a free data retrieval call binding the contract method 0xbe40ba79.
//
// Solidity: function isSetWhitelist() constant returns(bool)
func (_WalletMock *WalletMockCaller) IsSetWhitelist(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _WalletMock.contract.Call(opts, out, "isSetWhitelist")
	return *ret0, err
}

// IsSetWhitelist is a free data retrieval call binding the contract method 0xbe40ba79.
//
// Solidity: function isSetWhitelist() constant returns(bool)
func (_WalletMock *WalletMockSession) IsSetWhitelist() (bool, error) {
	return _WalletMock.Contract.IsSetWhitelist(&_WalletMock.CallOpts)
}

// IsSetWhitelist is a free data retrieval call binding the contract method 0xbe40ba79.
//
// Solidity: function isSetWhitelist() constant returns(bool)
func (_WalletMock *WalletMockCallerSession) IsSetWhitelist() (bool, error) {
	return _WalletMock.Contract.IsSetWhitelist(&_WalletMock.CallOpts)
}

// IsTransferable is a free data retrieval call binding the contract method 0x2121dc75.
//
// Solidity: function isTransferable() constant returns(bool)
func (_WalletMock *WalletMockCaller) IsTransferable(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _WalletMock.contract.Call(opts, out, "isTransferable")
	return *ret0, err
}

// IsTransferable is a free data retrieval call binding the contract method 0x2121dc75.
//
// Solidity: function isTransferable() constant returns(bool)
func (_WalletMock *WalletMockSession) IsTransferable() (bool, error) {
	return _WalletMock.Contract.IsTransferable(&_WalletMock.CallOpts)
}

// IsTransferable is a free data retrieval call binding the contract method 0x2121dc75.
//
// Solidity: function isTransferable() constant returns(bool)
func (_WalletMock *WalletMockCallerSession) IsTransferable() (bool, error) {
	return _WalletMock.Contract.IsTransferable(&_WalletMock.CallOpts)
}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 _hashedData, bytes _signature) constant returns(bytes4)
func (_WalletMock *WalletMockCaller) IsValidSignature(opts *bind.CallOpts, _hashedData [32]byte, _signature []byte) ([4]byte, error) {
	var (
		ret0 = new([4]byte)
	)
	out := ret0
	err := _WalletMock.contract.Call(opts, out, "isValidSignature", _hashedData, _signature)
	return *ret0, err
}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 _hashedData, bytes _signature) constant returns(bytes4)
func (_WalletMock *WalletMockSession) IsValidSignature(_hashedData [32]byte, _signature []byte) ([4]byte, error) {
	return _WalletMock.Contract.IsValidSignature(&_WalletMock.CallOpts, _hashedData, _signature)
}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 _hashedData, bytes _signature) constant returns(bytes4)
func (_WalletMock *WalletMockCallerSession) IsValidSignature(_hashedData [32]byte, _signature []byte) ([4]byte, error) {
	return _WalletMock.Contract.IsValidSignature(&_WalletMock.CallOpts, _hashedData, _signature)
}

// IsValidSignature0 is a free data retrieval call binding the contract method 0x20c13b0b.
//
// Solidity: function isValidSignature(bytes _data, bytes _signature) constant returns(bytes4)
func (_WalletMock *WalletMockCaller) IsValidSignature0(opts *bind.CallOpts, _data []byte, _signature []byte) ([4]byte, error) {
	var (
		ret0 = new([4]byte)
	)
	out := ret0
	err := _WalletMock.contract.Call(opts, out, "isValidSignature0", _data, _signature)
	return *ret0, err
}

// IsValidSignature0 is a free data retrieval call binding the contract method 0x20c13b0b.
//
// Solidity: function isValidSignature(bytes _data, bytes _signature) constant returns(bytes4)
func (_WalletMock *WalletMockSession) IsValidSignature0(_data []byte, _signature []byte) ([4]byte, error) {
	return _WalletMock.Contract.IsValidSignature0(&_WalletMock.CallOpts, _data, _signature)
}

// IsValidSignature0 is a free data retrieval call binding the contract method 0x20c13b0b.
//
// Solidity: function isValidSignature(bytes _data, bytes _signature) constant returns(bytes4)
func (_WalletMock *WalletMockCallerSession) IsValidSignature0(_data []byte, _signature []byte) ([4]byte, error) {
	return _WalletMock.Contract.IsValidSignature0(&_WalletMock.CallOpts, _data, _signature)
}

// LicenceNode is a free data retrieval call binding the contract method 0x747c31d6.
//
// Solidity: function licenceNode() constant returns(bytes32)
func (_WalletMock *WalletMockCaller) LicenceNode(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _WalletMock.contract.Call(opts, out, "licenceNode")
	return *ret0, err
}

// LicenceNode is a free data retrieval call binding the contract method 0x747c31d6.
//
// Solidity: function licenceNode() constant returns(bytes32)
func (_WalletMock *WalletMockSession) LicenceNode() ([32]byte, error) {
	return _WalletMock.Contract.LicenceNode(&_WalletMock.CallOpts)
}

// LicenceNode is a free data retrieval call binding the contract method 0x747c31d6.
//
// Solidity: function licenceNode() constant returns(bytes32)
func (_WalletMock *WalletMockCallerSession) LicenceNode() ([32]byte, error) {
	return _WalletMock.Contract.LicenceNode(&_WalletMock.CallOpts)
}

// LoadLimitAvailable is a free data retrieval call binding the contract method 0x1efd0299.
//
// Solidity: function loadLimitAvailable() constant returns(uint256)
func (_WalletMock *WalletMockCaller) LoadLimitAvailable(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _WalletMock.contract.Call(opts, out, "loadLimitAvailable")
	return *ret0, err
}

// LoadLimitAvailable is a free data retrieval call binding the contract method 0x1efd0299.
//
// Solidity: function loadLimitAvailable() constant returns(uint256)
func (_WalletMock *WalletMockSession) LoadLimitAvailable() (*big.Int, error) {
	return _WalletMock.Contract.LoadLimitAvailable(&_WalletMock.CallOpts)
}

// LoadLimitAvailable is a free data retrieval call binding the contract method 0x1efd0299.
//
// Solidity: function loadLimitAvailable() constant returns(uint256)
func (_WalletMock *WalletMockCallerSession) LoadLimitAvailable() (*big.Int, error) {
	return _WalletMock.Contract.LoadLimitAvailable(&_WalletMock.CallOpts)
}

// LoadLimitControllerConfirmationRequired is a free data retrieval call binding the contract method 0xab205993.
//
// Solidity: function loadLimitControllerConfirmationRequired() constant returns(bool)
func (_WalletMock *WalletMockCaller) LoadLimitControllerConfirmationRequired(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _WalletMock.contract.Call(opts, out, "loadLimitControllerConfirmationRequired")
	return *ret0, err
}

// LoadLimitControllerConfirmationRequired is a free data retrieval call binding the contract method 0xab205993.
//
// Solidity: function loadLimitControllerConfirmationRequired() constant returns(bool)
func (_WalletMock *WalletMockSession) LoadLimitControllerConfirmationRequired() (bool, error) {
	return _WalletMock.Contract.LoadLimitControllerConfirmationRequired(&_WalletMock.CallOpts)
}

// LoadLimitControllerConfirmationRequired is a free data retrieval call binding the contract method 0xab205993.
//
// Solidity: function loadLimitControllerConfirmationRequired() constant returns(bool)
func (_WalletMock *WalletMockCallerSession) LoadLimitControllerConfirmationRequired() (bool, error) {
	return _WalletMock.Contract.LoadLimitControllerConfirmationRequired(&_WalletMock.CallOpts)
}

// LoadLimitPending is a free data retrieval call binding the contract method 0xc4856cd9.
//
// Solidity: function loadLimitPending() constant returns(uint256)
func (_WalletMock *WalletMockCaller) LoadLimitPending(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _WalletMock.contract.Call(opts, out, "loadLimitPending")
	return *ret0, err
}

// LoadLimitPending is a free data retrieval call binding the contract method 0xc4856cd9.
//
// Solidity: function loadLimitPending() constant returns(uint256)
func (_WalletMock *WalletMockSession) LoadLimitPending() (*big.Int, error) {
	return _WalletMock.Contract.LoadLimitPending(&_WalletMock.CallOpts)
}

// LoadLimitPending is a free data retrieval call binding the contract method 0xc4856cd9.
//
// Solidity: function loadLimitPending() constant returns(uint256)
func (_WalletMock *WalletMockCallerSession) LoadLimitPending() (*big.Int, error) {
	return _WalletMock.Contract.LoadLimitPending(&_WalletMock.CallOpts)
}

// LoadLimitValue is a free data retrieval call binding the contract method 0xda84b1ed.
//
// Solidity: function loadLimitValue() constant returns(uint256)
func (_WalletMock *WalletMockCaller) LoadLimitValue(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _WalletMock.contract.Call(opts, out, "loadLimitValue")
	return *ret0, err
}

// LoadLimitValue is a free data retrieval call binding the contract method 0xda84b1ed.
//
// Solidity: function loadLimitValue() constant returns(uint256)
func (_WalletMock *WalletMockSession) LoadLimitValue() (*big.Int, error) {
	return _WalletMock.Contract.LoadLimitValue(&_WalletMock.CallOpts)
}

// LoadLimitValue is a free data retrieval call binding the contract method 0xda84b1ed.
//
// Solidity: function loadLimitValue() constant returns(uint256)
func (_WalletMock *WalletMockCallerSession) LoadLimitValue() (*big.Int, error) {
	return _WalletMock.Contract.LoadLimitValue(&_WalletMock.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_WalletMock *WalletMockCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _WalletMock.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_WalletMock *WalletMockSession) Owner() (common.Address, error) {
	return _WalletMock.Contract.Owner(&_WalletMock.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_WalletMock *WalletMockCallerSession) Owner() (common.Address, error) {
	return _WalletMock.Contract.Owner(&_WalletMock.CallOpts)
}

// PendingWhitelistAddition is a free data retrieval call binding the contract method 0x47b55a9d.
//
// Solidity: function pendingWhitelistAddition() constant returns(address[])
func (_WalletMock *WalletMockCaller) PendingWhitelistAddition(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _WalletMock.contract.Call(opts, out, "pendingWhitelistAddition")
	return *ret0, err
}

// PendingWhitelistAddition is a free data retrieval call binding the contract method 0x47b55a9d.
//
// Solidity: function pendingWhitelistAddition() constant returns(address[])
func (_WalletMock *WalletMockSession) PendingWhitelistAddition() ([]common.Address, error) {
	return _WalletMock.Contract.PendingWhitelistAddition(&_WalletMock.CallOpts)
}

// PendingWhitelistAddition is a free data retrieval call binding the contract method 0x47b55a9d.
//
// Solidity: function pendingWhitelistAddition() constant returns(address[])
func (_WalletMock *WalletMockCallerSession) PendingWhitelistAddition() ([]common.Address, error) {
	return _WalletMock.Contract.PendingWhitelistAddition(&_WalletMock.CallOpts)
}

// PendingWhitelistRemoval is a free data retrieval call binding the contract method 0x294f4025.
//
// Solidity: function pendingWhitelistRemoval() constant returns(address[])
func (_WalletMock *WalletMockCaller) PendingWhitelistRemoval(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _WalletMock.contract.Call(opts, out, "pendingWhitelistRemoval")
	return *ret0, err
}

// PendingWhitelistRemoval is a free data retrieval call binding the contract method 0x294f4025.
//
// Solidity: function pendingWhitelistRemoval() constant returns(address[])
func (_WalletMock *WalletMockSession) PendingWhitelistRemoval() ([]common.Address, error) {
	return _WalletMock.Contract.PendingWhitelistRemoval(&_WalletMock.CallOpts)
}

// PendingWhitelistRemoval is a free data retrieval call binding the contract method 0x294f4025.
//
// Solidity: function pendingWhitelistRemoval() constant returns(address[])
func (_WalletMock *WalletMockCallerSession) PendingWhitelistRemoval() ([]common.Address, error) {
	return _WalletMock.Contract.PendingWhitelistRemoval(&_WalletMock.CallOpts)
}

// RelayNonce is a free data retrieval call binding the contract method 0xcccdc556.
//
// Solidity: function relayNonce() constant returns(uint256)
func (_WalletMock *WalletMockCaller) RelayNonce(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _WalletMock.contract.Call(opts, out, "relayNonce")
	return *ret0, err
}

// RelayNonce is a free data retrieval call binding the contract method 0xcccdc556.
//
// Solidity: function relayNonce() constant returns(uint256)
func (_WalletMock *WalletMockSession) RelayNonce() (*big.Int, error) {
	return _WalletMock.Contract.RelayNonce(&_WalletMock.CallOpts)
}

// RelayNonce is a free data retrieval call binding the contract method 0xcccdc556.
//
// Solidity: function relayNonce() constant returns(uint256)
func (_WalletMock *WalletMockCallerSession) RelayNonce() (*big.Int, error) {
	return _WalletMock.Contract.RelayNonce(&_WalletMock.CallOpts)
}

// SpendLimitAvailable is a free data retrieval call binding the contract method 0x5d2362a8.
//
// Solidity: function spendLimitAvailable() constant returns(uint256)
func (_WalletMock *WalletMockCaller) SpendLimitAvailable(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _WalletMock.contract.Call(opts, out, "spendLimitAvailable")
	return *ret0, err
}

// SpendLimitAvailable is a free data retrieval call binding the contract method 0x5d2362a8.
//
// Solidity: function spendLimitAvailable() constant returns(uint256)
func (_WalletMock *WalletMockSession) SpendLimitAvailable() (*big.Int, error) {
	return _WalletMock.Contract.SpendLimitAvailable(&_WalletMock.CallOpts)
}

// SpendLimitAvailable is a free data retrieval call binding the contract method 0x5d2362a8.
//
// Solidity: function spendLimitAvailable() constant returns(uint256)
func (_WalletMock *WalletMockCallerSession) SpendLimitAvailable() (*big.Int, error) {
	return _WalletMock.Contract.SpendLimitAvailable(&_WalletMock.CallOpts)
}

// SpendLimitControllerConfirmationRequired is a free data retrieval call binding the contract method 0xbcb8b74a.
//
// Solidity: function spendLimitControllerConfirmationRequired() constant returns(bool)
func (_WalletMock *WalletMockCaller) SpendLimitControllerConfirmationRequired(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _WalletMock.contract.Call(opts, out, "spendLimitControllerConfirmationRequired")
	return *ret0, err
}

// SpendLimitControllerConfirmationRequired is a free data retrieval call binding the contract method 0xbcb8b74a.
//
// Solidity: function spendLimitControllerConfirmationRequired() constant returns(bool)
func (_WalletMock *WalletMockSession) SpendLimitControllerConfirmationRequired() (bool, error) {
	return _WalletMock.Contract.SpendLimitControllerConfirmationRequired(&_WalletMock.CallOpts)
}

// SpendLimitControllerConfirmationRequired is a free data retrieval call binding the contract method 0xbcb8b74a.
//
// Solidity: function spendLimitControllerConfirmationRequired() constant returns(bool)
func (_WalletMock *WalletMockCallerSession) SpendLimitControllerConfirmationRequired() (bool, error) {
	return _WalletMock.Contract.SpendLimitControllerConfirmationRequired(&_WalletMock.CallOpts)
}

// SpendLimitPending is a free data retrieval call binding the contract method 0x027ef3eb.
//
// Solidity: function spendLimitPending() constant returns(uint256)
func (_WalletMock *WalletMockCaller) SpendLimitPending(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _WalletMock.contract.Call(opts, out, "spendLimitPending")
	return *ret0, err
}

// SpendLimitPending is a free data retrieval call binding the contract method 0x027ef3eb.
//
// Solidity: function spendLimitPending() constant returns(uint256)
func (_WalletMock *WalletMockSession) SpendLimitPending() (*big.Int, error) {
	return _WalletMock.Contract.SpendLimitPending(&_WalletMock.CallOpts)
}

// SpendLimitPending is a free data retrieval call binding the contract method 0x027ef3eb.
//
// Solidity: function spendLimitPending() constant returns(uint256)
func (_WalletMock *WalletMockCallerSession) SpendLimitPending() (*big.Int, error) {
	return _WalletMock.Contract.SpendLimitPending(&_WalletMock.CallOpts)
}

// SpendLimitValue is a free data retrieval call binding the contract method 0x9b0dfd27.
//
// Solidity: function spendLimitValue() constant returns(uint256)
func (_WalletMock *WalletMockCaller) SpendLimitValue(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _WalletMock.contract.Call(opts, out, "spendLimitValue")
	return *ret0, err
}

// SpendLimitValue is a free data retrieval call binding the contract method 0x9b0dfd27.
//
// Solidity: function spendLimitValue() constant returns(uint256)
func (_WalletMock *WalletMockSession) SpendLimitValue() (*big.Int, error) {
	return _WalletMock.Contract.SpendLimitValue(&_WalletMock.CallOpts)
}

// SpendLimitValue is a free data retrieval call binding the contract method 0x9b0dfd27.
//
// Solidity: function spendLimitValue() constant returns(uint256)
func (_WalletMock *WalletMockCallerSession) SpendLimitValue() (*big.Int, error) {
	return _WalletMock.Contract.SpendLimitValue(&_WalletMock.CallOpts)
}

// SubmittedWhitelistAddition is a free data retrieval call binding the contract method 0x26d05ab2.
//
// Solidity: function submittedWhitelistAddition() constant returns(bool)
func (_WalletMock *WalletMockCaller) SubmittedWhitelistAddition(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _WalletMock.contract.Call(opts, out, "submittedWhitelistAddition")
	return *ret0, err
}

// SubmittedWhitelistAddition is a free data retrieval call binding the contract method 0x26d05ab2.
//
// Solidity: function submittedWhitelistAddition() constant returns(bool)
func (_WalletMock *WalletMockSession) SubmittedWhitelistAddition() (bool, error) {
	return _WalletMock.Contract.SubmittedWhitelistAddition(&_WalletMock.CallOpts)
}

// SubmittedWhitelistAddition is a free data retrieval call binding the contract method 0x26d05ab2.
//
// Solidity: function submittedWhitelistAddition() constant returns(bool)
func (_WalletMock *WalletMockCallerSession) SubmittedWhitelistAddition() (bool, error) {
	return _WalletMock.Contract.SubmittedWhitelistAddition(&_WalletMock.CallOpts)
}

// SubmittedWhitelistRemoval is a free data retrieval call binding the contract method 0xde212bf3.
//
// Solidity: function submittedWhitelistRemoval() constant returns(bool)
func (_WalletMock *WalletMockCaller) SubmittedWhitelistRemoval(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _WalletMock.contract.Call(opts, out, "submittedWhitelistRemoval")
	return *ret0, err
}

// SubmittedWhitelistRemoval is a free data retrieval call binding the contract method 0xde212bf3.
//
// Solidity: function submittedWhitelistRemoval() constant returns(bool)
func (_WalletMock *WalletMockSession) SubmittedWhitelistRemoval() (bool, error) {
	return _WalletMock.Contract.SubmittedWhitelistRemoval(&_WalletMock.CallOpts)
}

// SubmittedWhitelistRemoval is a free data retrieval call binding the contract method 0xde212bf3.
//
// Solidity: function submittedWhitelistRemoval() constant returns(bool)
func (_WalletMock *WalletMockCallerSession) SubmittedWhitelistRemoval() (bool, error) {
	return _WalletMock.Contract.SubmittedWhitelistRemoval(&_WalletMock.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceID) constant returns(bool)
func (_WalletMock *WalletMockCaller) SupportsInterface(opts *bind.CallOpts, _interfaceID [4]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _WalletMock.contract.Call(opts, out, "supportsInterface", _interfaceID)
	return *ret0, err
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceID) constant returns(bool)
func (_WalletMock *WalletMockSession) SupportsInterface(_interfaceID [4]byte) (bool, error) {
	return _WalletMock.Contract.SupportsInterface(&_WalletMock.CallOpts, _interfaceID)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceID) constant returns(bool)
func (_WalletMock *WalletMockCallerSession) SupportsInterface(_interfaceID [4]byte) (bool, error) {
	return _WalletMock.Contract.SupportsInterface(&_WalletMock.CallOpts, _interfaceID)
}

// TokenWhitelistNode is a free data retrieval call binding the contract method 0x877337b0.
//
// Solidity: function tokenWhitelistNode() constant returns(bytes32)
func (_WalletMock *WalletMockCaller) TokenWhitelistNode(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _WalletMock.contract.Call(opts, out, "tokenWhitelistNode")
	return *ret0, err
}

// TokenWhitelistNode is a free data retrieval call binding the contract method 0x877337b0.
//
// Solidity: function tokenWhitelistNode() constant returns(bytes32)
func (_WalletMock *WalletMockSession) TokenWhitelistNode() ([32]byte, error) {
	return _WalletMock.Contract.TokenWhitelistNode(&_WalletMock.CallOpts)
}

// TokenWhitelistNode is a free data retrieval call binding the contract method 0x877337b0.
//
// Solidity: function tokenWhitelistNode() constant returns(bytes32)
func (_WalletMock *WalletMockCallerSession) TokenWhitelistNode() ([32]byte, error) {
	return _WalletMock.Contract.TokenWhitelistNode(&_WalletMock.CallOpts)
}

// WhitelistArray is a free data retrieval call binding the contract method 0xd251fefc.
//
// Solidity: function whitelistArray(uint256 ) constant returns(address)
func (_WalletMock *WalletMockCaller) WhitelistArray(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _WalletMock.contract.Call(opts, out, "whitelistArray", arg0)
	return *ret0, err
}

// WhitelistArray is a free data retrieval call binding the contract method 0xd251fefc.
//
// Solidity: function whitelistArray(uint256 ) constant returns(address)
func (_WalletMock *WalletMockSession) WhitelistArray(arg0 *big.Int) (common.Address, error) {
	return _WalletMock.Contract.WhitelistArray(&_WalletMock.CallOpts, arg0)
}

// WhitelistArray is a free data retrieval call binding the contract method 0xd251fefc.
//
// Solidity: function whitelistArray(uint256 ) constant returns(address)
func (_WalletMock *WalletMockCallerSession) WhitelistArray(arg0 *big.Int) (common.Address, error) {
	return _WalletMock.Contract.WhitelistArray(&_WalletMock.CallOpts, arg0)
}

// WhitelistMap is a free data retrieval call binding the contract method 0x32531c3c.
//
// Solidity: function whitelistMap(address ) constant returns(bool)
func (_WalletMock *WalletMockCaller) WhitelistMap(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _WalletMock.contract.Call(opts, out, "whitelistMap", arg0)
	return *ret0, err
}

// WhitelistMap is a free data retrieval call binding the contract method 0x32531c3c.
//
// Solidity: function whitelistMap(address ) constant returns(bool)
func (_WalletMock *WalletMockSession) WhitelistMap(arg0 common.Address) (bool, error) {
	return _WalletMock.Contract.WhitelistMap(&_WalletMock.CallOpts, arg0)
}

// WhitelistMap is a free data retrieval call binding the contract method 0x32531c3c.
//
// Solidity: function whitelistMap(address ) constant returns(bool)
func (_WalletMock *WalletMockCallerSession) WhitelistMap(arg0 common.Address) (bool, error) {
	return _WalletMock.Contract.WhitelistMap(&_WalletMock.CallOpts, arg0)
}

// BatchExecuteTransaction is a paid mutator transaction binding the contract method 0xaaf1fc62.
//
// Solidity: function batchExecuteTransaction(bytes _transactionBatch) returns()
func (_WalletMock *WalletMockTransactor) BatchExecuteTransaction(opts *bind.TransactOpts, _transactionBatch []byte) (*types.Transaction, error) {
	return _WalletMock.contract.Transact(opts, "batchExecuteTransaction", _transactionBatch)
}

// BatchExecuteTransaction is a paid mutator transaction binding the contract method 0xaaf1fc62.
//
// Solidity: function batchExecuteTransaction(bytes _transactionBatch) returns()
func (_WalletMock *WalletMockSession) BatchExecuteTransaction(_transactionBatch []byte) (*types.Transaction, error) {
	return _WalletMock.Contract.BatchExecuteTransaction(&_WalletMock.TransactOpts, _transactionBatch)
}

// BatchExecuteTransaction is a paid mutator transaction binding the contract method 0xaaf1fc62.
//
// Solidity: function batchExecuteTransaction(bytes _transactionBatch) returns()
func (_WalletMock *WalletMockTransactorSession) BatchExecuteTransaction(_transactionBatch []byte) (*types.Transaction, error) {
	return _WalletMock.Contract.BatchExecuteTransaction(&_WalletMock.TransactOpts, _transactionBatch)
}

// BulkTransfer is a paid mutator transaction binding the contract method 0x1aa21fba.
//
// Solidity: function bulkTransfer(address _to, address[] _assets) returns()
func (_WalletMock *WalletMockTransactor) BulkTransfer(opts *bind.TransactOpts, _to common.Address, _assets []common.Address) (*types.Transaction, error) {
	return _WalletMock.contract.Transact(opts, "bulkTransfer", _to, _assets)
}

// BulkTransfer is a paid mutator transaction binding the contract method 0x1aa21fba.
//
// Solidity: function bulkTransfer(address _to, address[] _assets) returns()
func (_WalletMock *WalletMockSession) BulkTransfer(_to common.Address, _assets []common.Address) (*types.Transaction, error) {
	return _WalletMock.Contract.BulkTransfer(&_WalletMock.TransactOpts, _to, _assets)
}

// BulkTransfer is a paid mutator transaction binding the contract method 0x1aa21fba.
//
// Solidity: function bulkTransfer(address _to, address[] _assets) returns()
func (_WalletMock *WalletMockTransactorSession) BulkTransfer(_to common.Address, _assets []common.Address) (*types.Transaction, error) {
	return _WalletMock.Contract.BulkTransfer(&_WalletMock.TransactOpts, _to, _assets)
}

// CancelWhitelistAddition is a paid mutator transaction binding the contract method 0x100f23fd.
//
// Solidity: function cancelWhitelistAddition(bytes32 _hash) returns()
func (_WalletMock *WalletMockTransactor) CancelWhitelistAddition(opts *bind.TransactOpts, _hash [32]byte) (*types.Transaction, error) {
	return _WalletMock.contract.Transact(opts, "cancelWhitelistAddition", _hash)
}

// CancelWhitelistAddition is a paid mutator transaction binding the contract method 0x100f23fd.
//
// Solidity: function cancelWhitelistAddition(bytes32 _hash) returns()
func (_WalletMock *WalletMockSession) CancelWhitelistAddition(_hash [32]byte) (*types.Transaction, error) {
	return _WalletMock.Contract.CancelWhitelistAddition(&_WalletMock.TransactOpts, _hash)
}

// CancelWhitelistAddition is a paid mutator transaction binding the contract method 0x100f23fd.
//
// Solidity: function cancelWhitelistAddition(bytes32 _hash) returns()
func (_WalletMock *WalletMockTransactorSession) CancelWhitelistAddition(_hash [32]byte) (*types.Transaction, error) {
	return _WalletMock.Contract.CancelWhitelistAddition(&_WalletMock.TransactOpts, _hash)
}

// CancelWhitelistRemoval is a paid mutator transaction binding the contract method 0xce0b5bd5.
//
// Solidity: function cancelWhitelistRemoval(bytes32 _hash) returns()
func (_WalletMock *WalletMockTransactor) CancelWhitelistRemoval(opts *bind.TransactOpts, _hash [32]byte) (*types.Transaction, error) {
	return _WalletMock.contract.Transact(opts, "cancelWhitelistRemoval", _hash)
}

// CancelWhitelistRemoval is a paid mutator transaction binding the contract method 0xce0b5bd5.
//
// Solidity: function cancelWhitelistRemoval(bytes32 _hash) returns()
func (_WalletMock *WalletMockSession) CancelWhitelistRemoval(_hash [32]byte) (*types.Transaction, error) {
	return _WalletMock.Contract.CancelWhitelistRemoval(&_WalletMock.TransactOpts, _hash)
}

// CancelWhitelistRemoval is a paid mutator transaction binding the contract method 0xce0b5bd5.
//
// Solidity: function cancelWhitelistRemoval(bytes32 _hash) returns()
func (_WalletMock *WalletMockTransactorSession) CancelWhitelistRemoval(_hash [32]byte) (*types.Transaction, error) {
	return _WalletMock.Contract.CancelWhitelistRemoval(&_WalletMock.TransactOpts, _hash)
}

// ConfirmGasTopUpLimitUpdate is a paid mutator transaction binding the contract method 0xf41c4319.
//
// Solidity: function confirmGasTopUpLimitUpdate(uint256 _amount) returns()
func (_WalletMock *WalletMockTransactor) ConfirmGasTopUpLimitUpdate(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _WalletMock.contract.Transact(opts, "confirmGasTopUpLimitUpdate", _amount)
}

// ConfirmGasTopUpLimitUpdate is a paid mutator transaction binding the contract method 0xf41c4319.
//
// Solidity: function confirmGasTopUpLimitUpdate(uint256 _amount) returns()
func (_WalletMock *WalletMockSession) ConfirmGasTopUpLimitUpdate(_amount *big.Int) (*types.Transaction, error) {
	return _WalletMock.Contract.ConfirmGasTopUpLimitUpdate(&_WalletMock.TransactOpts, _amount)
}

// ConfirmGasTopUpLimitUpdate is a paid mutator transaction binding the contract method 0xf41c4319.
//
// Solidity: function confirmGasTopUpLimitUpdate(uint256 _amount) returns()
func (_WalletMock *WalletMockTransactorSession) ConfirmGasTopUpLimitUpdate(_amount *big.Int) (*types.Transaction, error) {
	return _WalletMock.Contract.ConfirmGasTopUpLimitUpdate(&_WalletMock.TransactOpts, _amount)
}

// ConfirmLoadLimitUpdate is a paid mutator transaction binding the contract method 0xf40b51f8.
//
// Solidity: function confirmLoadLimitUpdate(uint256 _amount) returns()
func (_WalletMock *WalletMockTransactor) ConfirmLoadLimitUpdate(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _WalletMock.contract.Transact(opts, "confirmLoadLimitUpdate", _amount)
}

// ConfirmLoadLimitUpdate is a paid mutator transaction binding the contract method 0xf40b51f8.
//
// Solidity: function confirmLoadLimitUpdate(uint256 _amount) returns()
func (_WalletMock *WalletMockSession) ConfirmLoadLimitUpdate(_amount *big.Int) (*types.Transaction, error) {
	return _WalletMock.Contract.ConfirmLoadLimitUpdate(&_WalletMock.TransactOpts, _amount)
}

// ConfirmLoadLimitUpdate is a paid mutator transaction binding the contract method 0xf40b51f8.
//
// Solidity: function confirmLoadLimitUpdate(uint256 _amount) returns()
func (_WalletMock *WalletMockTransactorSession) ConfirmLoadLimitUpdate(_amount *big.Int) (*types.Transaction, error) {
	return _WalletMock.Contract.ConfirmLoadLimitUpdate(&_WalletMock.TransactOpts, _amount)
}

// ConfirmSpendLimitUpdate is a paid mutator transaction binding the contract method 0xeadd3cea.
//
// Solidity: function confirmSpendLimitUpdate(uint256 _amount) returns()
func (_WalletMock *WalletMockTransactor) ConfirmSpendLimitUpdate(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _WalletMock.contract.Transact(opts, "confirmSpendLimitUpdate", _amount)
}

// ConfirmSpendLimitUpdate is a paid mutator transaction binding the contract method 0xeadd3cea.
//
// Solidity: function confirmSpendLimitUpdate(uint256 _amount) returns()
func (_WalletMock *WalletMockSession) ConfirmSpendLimitUpdate(_amount *big.Int) (*types.Transaction, error) {
	return _WalletMock.Contract.ConfirmSpendLimitUpdate(&_WalletMock.TransactOpts, _amount)
}

// ConfirmSpendLimitUpdate is a paid mutator transaction binding the contract method 0xeadd3cea.
//
// Solidity: function confirmSpendLimitUpdate(uint256 _amount) returns()
func (_WalletMock *WalletMockTransactorSession) ConfirmSpendLimitUpdate(_amount *big.Int) (*types.Transaction, error) {
	return _WalletMock.Contract.ConfirmSpendLimitUpdate(&_WalletMock.TransactOpts, _amount)
}

// ConfirmWhitelistAddition is a paid mutator transaction binding the contract method 0x5adc02ab.
//
// Solidity: function confirmWhitelistAddition(bytes32 _hash) returns()
func (_WalletMock *WalletMockTransactor) ConfirmWhitelistAddition(opts *bind.TransactOpts, _hash [32]byte) (*types.Transaction, error) {
	return _WalletMock.contract.Transact(opts, "confirmWhitelistAddition", _hash)
}

// ConfirmWhitelistAddition is a paid mutator transaction binding the contract method 0x5adc02ab.
//
// Solidity: function confirmWhitelistAddition(bytes32 _hash) returns()
func (_WalletMock *WalletMockSession) ConfirmWhitelistAddition(_hash [32]byte) (*types.Transaction, error) {
	return _WalletMock.Contract.ConfirmWhitelistAddition(&_WalletMock.TransactOpts, _hash)
}

// ConfirmWhitelistAddition is a paid mutator transaction binding the contract method 0x5adc02ab.
//
// Solidity: function confirmWhitelistAddition(bytes32 _hash) returns()
func (_WalletMock *WalletMockTransactorSession) ConfirmWhitelistAddition(_hash [32]byte) (*types.Transaction, error) {
	return _WalletMock.Contract.ConfirmWhitelistAddition(&_WalletMock.TransactOpts, _hash)
}

// ConfirmWhitelistRemoval is a paid mutator transaction binding the contract method 0xcbd2ac68.
//
// Solidity: function confirmWhitelistRemoval(bytes32 _hash) returns()
func (_WalletMock *WalletMockTransactor) ConfirmWhitelistRemoval(opts *bind.TransactOpts, _hash [32]byte) (*types.Transaction, error) {
	return _WalletMock.contract.Transact(opts, "confirmWhitelistRemoval", _hash)
}

// ConfirmWhitelistRemoval is a paid mutator transaction binding the contract method 0xcbd2ac68.
//
// Solidity: function confirmWhitelistRemoval(bytes32 _hash) returns()
func (_WalletMock *WalletMockSession) ConfirmWhitelistRemoval(_hash [32]byte) (*types.Transaction, error) {
	return _WalletMock.Contract.ConfirmWhitelistRemoval(&_WalletMock.TransactOpts, _hash)
}

// ConfirmWhitelistRemoval is a paid mutator transaction binding the contract method 0xcbd2ac68.
//
// Solidity: function confirmWhitelistRemoval(bytes32 _hash) returns()
func (_WalletMock *WalletMockTransactorSession) ConfirmWhitelistRemoval(_hash [32]byte) (*types.Transaction, error) {
	return _WalletMock.Contract.ConfirmWhitelistRemoval(&_WalletMock.TransactOpts, _hash)
}

// ExecuteRelayedTransaction is a paid mutator transaction binding the contract method 0x46efe0ed.
//
// Solidity: function executeRelayedTransaction(uint256 _nonce, bytes _data, bytes _signature) returns()
func (_WalletMock *WalletMockTransactor) ExecuteRelayedTransaction(opts *bind.TransactOpts, _nonce *big.Int, _data []byte, _signature []byte) (*types.Transaction, error) {
	return _WalletMock.contract.Transact(opts, "executeRelayedTransaction", _nonce, _data, _signature)
}

// ExecuteRelayedTransaction is a paid mutator transaction binding the contract method 0x46efe0ed.
//
// Solidity: function executeRelayedTransaction(uint256 _nonce, bytes _data, bytes _signature) returns()
func (_WalletMock *WalletMockSession) ExecuteRelayedTransaction(_nonce *big.Int, _data []byte, _signature []byte) (*types.Transaction, error) {
	return _WalletMock.Contract.ExecuteRelayedTransaction(&_WalletMock.TransactOpts, _nonce, _data, _signature)
}

// ExecuteRelayedTransaction is a paid mutator transaction binding the contract method 0x46efe0ed.
//
// Solidity: function executeRelayedTransaction(uint256 _nonce, bytes _data, bytes _signature) returns()
func (_WalletMock *WalletMockTransactorSession) ExecuteRelayedTransaction(_nonce *big.Int, _data []byte, _signature []byte) (*types.Transaction, error) {
	return _WalletMock.Contract.ExecuteRelayedTransaction(&_WalletMock.TransactOpts, _nonce, _data, _signature)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0x3f579f42.
//
// Solidity: function executeTransaction(address _destination, uint256 _value, bytes _data) returns(bytes)
func (_WalletMock *WalletMockTransactor) ExecuteTransaction(opts *bind.TransactOpts, _destination common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _WalletMock.contract.Transact(opts, "executeTransaction", _destination, _value, _data)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0x3f579f42.
//
// Solidity: function executeTransaction(address _destination, uint256 _value, bytes _data) returns(bytes)
func (_WalletMock *WalletMockSession) ExecuteTransaction(_destination common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _WalletMock.Contract.ExecuteTransaction(&_WalletMock.TransactOpts, _destination, _value, _data)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0x3f579f42.
//
// Solidity: function executeTransaction(address _destination, uint256 _value, bytes _data) returns(bytes)
func (_WalletMock *WalletMockTransactorSession) ExecuteTransaction(_destination common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _WalletMock.Contract.ExecuteTransaction(&_WalletMock.TransactOpts, _destination, _value, _data)
}

// IncreaseRelayNonce is a paid mutator transaction binding the contract method 0x90e690c7.
//
// Solidity: function increaseRelayNonce() returns()
func (_WalletMock *WalletMockTransactor) IncreaseRelayNonce(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WalletMock.contract.Transact(opts, "increaseRelayNonce")
}

// IncreaseRelayNonce is a paid mutator transaction binding the contract method 0x90e690c7.
//
// Solidity: function increaseRelayNonce() returns()
func (_WalletMock *WalletMockSession) IncreaseRelayNonce() (*types.Transaction, error) {
	return _WalletMock.Contract.IncreaseRelayNonce(&_WalletMock.TransactOpts)
}

// IncreaseRelayNonce is a paid mutator transaction binding the contract method 0x90e690c7.
//
// Solidity: function increaseRelayNonce() returns()
func (_WalletMock *WalletMockTransactorSession) IncreaseRelayNonce() (*types.Transaction, error) {
	return _WalletMock.Contract.IncreaseRelayNonce(&_WalletMock.TransactOpts)
}

// InitializeWallet is a paid mutator transaction binding the contract method 0x69efdfc0.
//
// Solidity: function initializeWallet(address _owner_, bool _transferable_, address _ens_, bytes32 _tokenWhitelistNode_, bytes32 _controllerNode_, bytes32 _licenceNode_, uint256 _spendLimit_) returns()
func (_WalletMock *WalletMockTransactor) InitializeWallet(opts *bind.TransactOpts, _owner_ common.Address, _transferable_ bool, _ens_ common.Address, _tokenWhitelistNode_ [32]byte, _controllerNode_ [32]byte, _licenceNode_ [32]byte, _spendLimit_ *big.Int) (*types.Transaction, error) {
	return _WalletMock.contract.Transact(opts, "initializeWallet", _owner_, _transferable_, _ens_, _tokenWhitelistNode_, _controllerNode_, _licenceNode_, _spendLimit_)
}

// InitializeWallet is a paid mutator transaction binding the contract method 0x69efdfc0.
//
// Solidity: function initializeWallet(address _owner_, bool _transferable_, address _ens_, bytes32 _tokenWhitelistNode_, bytes32 _controllerNode_, bytes32 _licenceNode_, uint256 _spendLimit_) returns()
func (_WalletMock *WalletMockSession) InitializeWallet(_owner_ common.Address, _transferable_ bool, _ens_ common.Address, _tokenWhitelistNode_ [32]byte, _controllerNode_ [32]byte, _licenceNode_ [32]byte, _spendLimit_ *big.Int) (*types.Transaction, error) {
	return _WalletMock.Contract.InitializeWallet(&_WalletMock.TransactOpts, _owner_, _transferable_, _ens_, _tokenWhitelistNode_, _controllerNode_, _licenceNode_, _spendLimit_)
}

// InitializeWallet is a paid mutator transaction binding the contract method 0x69efdfc0.
//
// Solidity: function initializeWallet(address _owner_, bool _transferable_, address _ens_, bytes32 _tokenWhitelistNode_, bytes32 _controllerNode_, bytes32 _licenceNode_, uint256 _spendLimit_) returns()
func (_WalletMock *WalletMockTransactorSession) InitializeWallet(_owner_ common.Address, _transferable_ bool, _ens_ common.Address, _tokenWhitelistNode_ [32]byte, _controllerNode_ [32]byte, _licenceNode_ [32]byte, _spendLimit_ *big.Int) (*types.Transaction, error) {
	return _WalletMock.Contract.InitializeWallet(&_WalletMock.TransactOpts, _owner_, _transferable_, _ens_, _tokenWhitelistNode_, _controllerNode_, _licenceNode_, _spendLimit_)
}

// LoadTokenCard is a paid mutator transaction binding the contract method 0x3a43199f.
//
// Solidity: function loadTokenCard(address _asset, uint256 _amount) returns()
func (_WalletMock *WalletMockTransactor) LoadTokenCard(opts *bind.TransactOpts, _asset common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _WalletMock.contract.Transact(opts, "loadTokenCard", _asset, _amount)
}

// LoadTokenCard is a paid mutator transaction binding the contract method 0x3a43199f.
//
// Solidity: function loadTokenCard(address _asset, uint256 _amount) returns()
func (_WalletMock *WalletMockSession) LoadTokenCard(_asset common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _WalletMock.Contract.LoadTokenCard(&_WalletMock.TransactOpts, _asset, _amount)
}

// LoadTokenCard is a paid mutator transaction binding the contract method 0x3a43199f.
//
// Solidity: function loadTokenCard(address _asset, uint256 _amount) returns()
func (_WalletMock *WalletMockTransactorSession) LoadTokenCard(_asset common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _WalletMock.Contract.LoadTokenCard(&_WalletMock.TransactOpts, _asset, _amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_WalletMock *WalletMockTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WalletMock.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_WalletMock *WalletMockSession) RenounceOwnership() (*types.Transaction, error) {
	return _WalletMock.Contract.RenounceOwnership(&_WalletMock.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_WalletMock *WalletMockTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _WalletMock.Contract.RenounceOwnership(&_WalletMock.TransactOpts)
}

// SetGasTopUpLimit is a paid mutator transaction binding the contract method 0x0f3a85d8.
//
// Solidity: function setGasTopUpLimit(uint256 _amount) returns()
func (_WalletMock *WalletMockTransactor) SetGasTopUpLimit(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _WalletMock.contract.Transact(opts, "setGasTopUpLimit", _amount)
}

// SetGasTopUpLimit is a paid mutator transaction binding the contract method 0x0f3a85d8.
//
// Solidity: function setGasTopUpLimit(uint256 _amount) returns()
func (_WalletMock *WalletMockSession) SetGasTopUpLimit(_amount *big.Int) (*types.Transaction, error) {
	return _WalletMock.Contract.SetGasTopUpLimit(&_WalletMock.TransactOpts, _amount)
}

// SetGasTopUpLimit is a paid mutator transaction binding the contract method 0x0f3a85d8.
//
// Solidity: function setGasTopUpLimit(uint256 _amount) returns()
func (_WalletMock *WalletMockTransactorSession) SetGasTopUpLimit(_amount *big.Int) (*types.Transaction, error) {
	return _WalletMock.Contract.SetGasTopUpLimit(&_WalletMock.TransactOpts, _amount)
}

// SetLoadLimit is a paid mutator transaction binding the contract method 0x3bfec254.
//
// Solidity: function setLoadLimit(uint256 _amount) returns()
func (_WalletMock *WalletMockTransactor) SetLoadLimit(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _WalletMock.contract.Transact(opts, "setLoadLimit", _amount)
}

// SetLoadLimit is a paid mutator transaction binding the contract method 0x3bfec254.
//
// Solidity: function setLoadLimit(uint256 _amount) returns()
func (_WalletMock *WalletMockSession) SetLoadLimit(_amount *big.Int) (*types.Transaction, error) {
	return _WalletMock.Contract.SetLoadLimit(&_WalletMock.TransactOpts, _amount)
}

// SetLoadLimit is a paid mutator transaction binding the contract method 0x3bfec254.
//
// Solidity: function setLoadLimit(uint256 _amount) returns()
func (_WalletMock *WalletMockTransactorSession) SetLoadLimit(_amount *big.Int) (*types.Transaction, error) {
	return _WalletMock.Contract.SetLoadLimit(&_WalletMock.TransactOpts, _amount)
}

// SetSpendLimit is a paid mutator transaction binding the contract method 0x3c672eb7.
//
// Solidity: function setSpendLimit(uint256 _amount) returns()
func (_WalletMock *WalletMockTransactor) SetSpendLimit(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _WalletMock.contract.Transact(opts, "setSpendLimit", _amount)
}

// SetSpendLimit is a paid mutator transaction binding the contract method 0x3c672eb7.
//
// Solidity: function setSpendLimit(uint256 _amount) returns()
func (_WalletMock *WalletMockSession) SetSpendLimit(_amount *big.Int) (*types.Transaction, error) {
	return _WalletMock.Contract.SetSpendLimit(&_WalletMock.TransactOpts, _amount)
}

// SetSpendLimit is a paid mutator transaction binding the contract method 0x3c672eb7.
//
// Solidity: function setSpendLimit(uint256 _amount) returns()
func (_WalletMock *WalletMockTransactorSession) SetSpendLimit(_amount *big.Int) (*types.Transaction, error) {
	return _WalletMock.Contract.SetSpendLimit(&_WalletMock.TransactOpts, _amount)
}

// SetWhitelist is a paid mutator transaction binding the contract method 0xf4217648.
//
// Solidity: function setWhitelist(address[] _addresses) returns()
func (_WalletMock *WalletMockTransactor) SetWhitelist(opts *bind.TransactOpts, _addresses []common.Address) (*types.Transaction, error) {
	return _WalletMock.contract.Transact(opts, "setWhitelist", _addresses)
}

// SetWhitelist is a paid mutator transaction binding the contract method 0xf4217648.
//
// Solidity: function setWhitelist(address[] _addresses) returns()
func (_WalletMock *WalletMockSession) SetWhitelist(_addresses []common.Address) (*types.Transaction, error) {
	return _WalletMock.Contract.SetWhitelist(&_WalletMock.TransactOpts, _addresses)
}

// SetWhitelist is a paid mutator transaction binding the contract method 0xf4217648.
//
// Solidity: function setWhitelist(address[] _addresses) returns()
func (_WalletMock *WalletMockTransactorSession) SetWhitelist(_addresses []common.Address) (*types.Transaction, error) {
	return _WalletMock.Contract.SetWhitelist(&_WalletMock.TransactOpts, _addresses)
}

// SubmitGasTopUpLimitUpdate is a paid mutator transaction binding the contract method 0x74624c55.
//
// Solidity: function submitGasTopUpLimitUpdate(uint256 _amount) returns()
func (_WalletMock *WalletMockTransactor) SubmitGasTopUpLimitUpdate(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _WalletMock.contract.Transact(opts, "submitGasTopUpLimitUpdate", _amount)
}

// SubmitGasTopUpLimitUpdate is a paid mutator transaction binding the contract method 0x74624c55.
//
// Solidity: function submitGasTopUpLimitUpdate(uint256 _amount) returns()
func (_WalletMock *WalletMockSession) SubmitGasTopUpLimitUpdate(_amount *big.Int) (*types.Transaction, error) {
	return _WalletMock.Contract.SubmitGasTopUpLimitUpdate(&_WalletMock.TransactOpts, _amount)
}

// SubmitGasTopUpLimitUpdate is a paid mutator transaction binding the contract method 0x74624c55.
//
// Solidity: function submitGasTopUpLimitUpdate(uint256 _amount) returns()
func (_WalletMock *WalletMockTransactorSession) SubmitGasTopUpLimitUpdate(_amount *big.Int) (*types.Transaction, error) {
	return _WalletMock.Contract.SubmitGasTopUpLimitUpdate(&_WalletMock.TransactOpts, _amount)
}

// SubmitLoadLimitUpdate is a paid mutator transaction binding the contract method 0xb221f316.
//
// Solidity: function submitLoadLimitUpdate(uint256 _amount) returns()
func (_WalletMock *WalletMockTransactor) SubmitLoadLimitUpdate(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _WalletMock.contract.Transact(opts, "submitLoadLimitUpdate", _amount)
}

// SubmitLoadLimitUpdate is a paid mutator transaction binding the contract method 0xb221f316.
//
// Solidity: function submitLoadLimitUpdate(uint256 _amount) returns()
func (_WalletMock *WalletMockSession) SubmitLoadLimitUpdate(_amount *big.Int) (*types.Transaction, error) {
	return _WalletMock.Contract.SubmitLoadLimitUpdate(&_WalletMock.TransactOpts, _amount)
}

// SubmitLoadLimitUpdate is a paid mutator transaction binding the contract method 0xb221f316.
//
// Solidity: function submitLoadLimitUpdate(uint256 _amount) returns()
func (_WalletMock *WalletMockTransactorSession) SubmitLoadLimitUpdate(_amount *big.Int) (*types.Transaction, error) {
	return _WalletMock.Contract.SubmitLoadLimitUpdate(&_WalletMock.TransactOpts, _amount)
}

// SubmitSpendLimitUpdate is a paid mutator transaction binding the contract method 0x21ce918d.
//
// Solidity: function submitSpendLimitUpdate(uint256 _amount) returns()
func (_WalletMock *WalletMockTransactor) SubmitSpendLimitUpdate(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _WalletMock.contract.Transact(opts, "submitSpendLimitUpdate", _amount)
}

// SubmitSpendLimitUpdate is a paid mutator transaction binding the contract method 0x21ce918d.
//
// Solidity: function submitSpendLimitUpdate(uint256 _amount) returns()
func (_WalletMock *WalletMockSession) SubmitSpendLimitUpdate(_amount *big.Int) (*types.Transaction, error) {
	return _WalletMock.Contract.SubmitSpendLimitUpdate(&_WalletMock.TransactOpts, _amount)
}

// SubmitSpendLimitUpdate is a paid mutator transaction binding the contract method 0x21ce918d.
//
// Solidity: function submitSpendLimitUpdate(uint256 _amount) returns()
func (_WalletMock *WalletMockTransactorSession) SubmitSpendLimitUpdate(_amount *big.Int) (*types.Transaction, error) {
	return _WalletMock.Contract.SubmitSpendLimitUpdate(&_WalletMock.TransactOpts, _amount)
}

// SubmitWhitelistAddition is a paid mutator transaction binding the contract method 0x7fd004fa.
//
// Solidity: function submitWhitelistAddition(address[] _addresses) returns()
func (_WalletMock *WalletMockTransactor) SubmitWhitelistAddition(opts *bind.TransactOpts, _addresses []common.Address) (*types.Transaction, error) {
	return _WalletMock.contract.Transact(opts, "submitWhitelistAddition", _addresses)
}

// SubmitWhitelistAddition is a paid mutator transaction binding the contract method 0x7fd004fa.
//
// Solidity: function submitWhitelistAddition(address[] _addresses) returns()
func (_WalletMock *WalletMockSession) SubmitWhitelistAddition(_addresses []common.Address) (*types.Transaction, error) {
	return _WalletMock.Contract.SubmitWhitelistAddition(&_WalletMock.TransactOpts, _addresses)
}

// SubmitWhitelistAddition is a paid mutator transaction binding the contract method 0x7fd004fa.
//
// Solidity: function submitWhitelistAddition(address[] _addresses) returns()
func (_WalletMock *WalletMockTransactorSession) SubmitWhitelistAddition(_addresses []common.Address) (*types.Transaction, error) {
	return _WalletMock.Contract.SubmitWhitelistAddition(&_WalletMock.TransactOpts, _addresses)
}

// SubmitWhitelistRemoval is a paid mutator transaction binding the contract method 0x6137d670.
//
// Solidity: function submitWhitelistRemoval(address[] _addresses) returns()
func (_WalletMock *WalletMockTransactor) SubmitWhitelistRemoval(opts *bind.TransactOpts, _addresses []common.Address) (*types.Transaction, error) {
	return _WalletMock.contract.Transact(opts, "submitWhitelistRemoval", _addresses)
}

// SubmitWhitelistRemoval is a paid mutator transaction binding the contract method 0x6137d670.
//
// Solidity: function submitWhitelistRemoval(address[] _addresses) returns()
func (_WalletMock *WalletMockSession) SubmitWhitelistRemoval(_addresses []common.Address) (*types.Transaction, error) {
	return _WalletMock.Contract.SubmitWhitelistRemoval(&_WalletMock.TransactOpts, _addresses)
}

// SubmitWhitelistRemoval is a paid mutator transaction binding the contract method 0x6137d670.
//
// Solidity: function submitWhitelistRemoval(address[] _addresses) returns()
func (_WalletMock *WalletMockTransactorSession) SubmitWhitelistRemoval(_addresses []common.Address) (*types.Transaction, error) {
	return _WalletMock.Contract.SubmitWhitelistRemoval(&_WalletMock.TransactOpts, _addresses)
}

// TopUpGas is a paid mutator transaction binding the contract method 0xe61c51ca.
//
// Solidity: function topUpGas(uint256 _amount) returns()
func (_WalletMock *WalletMockTransactor) TopUpGas(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _WalletMock.contract.Transact(opts, "topUpGas", _amount)
}

// TopUpGas is a paid mutator transaction binding the contract method 0xe61c51ca.
//
// Solidity: function topUpGas(uint256 _amount) returns()
func (_WalletMock *WalletMockSession) TopUpGas(_amount *big.Int) (*types.Transaction, error) {
	return _WalletMock.Contract.TopUpGas(&_WalletMock.TransactOpts, _amount)
}

// TopUpGas is a paid mutator transaction binding the contract method 0xe61c51ca.
//
// Solidity: function topUpGas(uint256 _amount) returns()
func (_WalletMock *WalletMockTransactorSession) TopUpGas(_amount *big.Int) (*types.Transaction, error) {
	return _WalletMock.Contract.TopUpGas(&_WalletMock.TransactOpts, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xbeabacc8.
//
// Solidity: function transfer(address _to, address _asset, uint256 _amount) returns()
func (_WalletMock *WalletMockTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _asset common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _WalletMock.contract.Transact(opts, "transfer", _to, _asset, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xbeabacc8.
//
// Solidity: function transfer(address _to, address _asset, uint256 _amount) returns()
func (_WalletMock *WalletMockSession) Transfer(_to common.Address, _asset common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _WalletMock.Contract.Transfer(&_WalletMock.TransactOpts, _to, _asset, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xbeabacc8.
//
// Solidity: function transfer(address _to, address _asset, uint256 _amount) returns()
func (_WalletMock *WalletMockTransactorSession) Transfer(_to common.Address, _asset common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _WalletMock.Contract.Transfer(&_WalletMock.TransactOpts, _to, _asset, _amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xb242e534.
//
// Solidity: function transferOwnership(address _account, bool _transferable) returns()
func (_WalletMock *WalletMockTransactor) TransferOwnership(opts *bind.TransactOpts, _account common.Address, _transferable bool) (*types.Transaction, error) {
	return _WalletMock.contract.Transact(opts, "transferOwnership", _account, _transferable)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xb242e534.
//
// Solidity: function transferOwnership(address _account, bool _transferable) returns()
func (_WalletMock *WalletMockSession) TransferOwnership(_account common.Address, _transferable bool) (*types.Transaction, error) {
	return _WalletMock.Contract.TransferOwnership(&_WalletMock.TransactOpts, _account, _transferable)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xb242e534.
//
// Solidity: function transferOwnership(address _account, bool _transferable) returns()
func (_WalletMock *WalletMockTransactorSession) TransferOwnership(_account common.Address, _transferable bool) (*types.Transaction, error) {
	return _WalletMock.Contract.TransferOwnership(&_WalletMock.TransactOpts, _account, _transferable)
}

// WalletMockAddedToWhitelistIterator is returned from FilterAddedToWhitelist and is used to iterate over the raw logs and unpacked data for AddedToWhitelist events raised by the WalletMock contract.
type WalletMockAddedToWhitelistIterator struct {
	Event *WalletMockAddedToWhitelist // Event containing the contract specifics and raw log

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
func (it *WalletMockAddedToWhitelistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletMockAddedToWhitelist)
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
		it.Event = new(WalletMockAddedToWhitelist)
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
func (it *WalletMockAddedToWhitelistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletMockAddedToWhitelistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletMockAddedToWhitelist represents a AddedToWhitelist event raised by the WalletMock contract.
type WalletMockAddedToWhitelist struct {
	Sender    common.Address
	Addresses []common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAddedToWhitelist is a free log retrieval operation binding the contract event 0xb2f6cccee7a369e23e293c25aa19bef80af11eb26deba3ea0f2a02783f752e4a.
//
// Solidity: event AddedToWhitelist(address _sender, address[] _addresses)
func (_WalletMock *WalletMockFilterer) FilterAddedToWhitelist(opts *bind.FilterOpts) (*WalletMockAddedToWhitelistIterator, error) {

	logs, sub, err := _WalletMock.contract.FilterLogs(opts, "AddedToWhitelist")
	if err != nil {
		return nil, err
	}
	return &WalletMockAddedToWhitelistIterator{contract: _WalletMock.contract, event: "AddedToWhitelist", logs: logs, sub: sub}, nil
}

// WatchAddedToWhitelist is a free log subscription operation binding the contract event 0xb2f6cccee7a369e23e293c25aa19bef80af11eb26deba3ea0f2a02783f752e4a.
//
// Solidity: event AddedToWhitelist(address _sender, address[] _addresses)
func (_WalletMock *WalletMockFilterer) WatchAddedToWhitelist(opts *bind.WatchOpts, sink chan<- *WalletMockAddedToWhitelist) (event.Subscription, error) {

	logs, sub, err := _WalletMock.contract.WatchLogs(opts, "AddedToWhitelist")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletMockAddedToWhitelist)
				if err := _WalletMock.contract.UnpackLog(event, "AddedToWhitelist", log); err != nil {
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
func (_WalletMock *WalletMockFilterer) ParseAddedToWhitelist(log types.Log) (*WalletMockAddedToWhitelist, error) {
	event := new(WalletMockAddedToWhitelist)
	if err := _WalletMock.contract.UnpackLog(event, "AddedToWhitelist", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WalletMockBulkTransferredIterator is returned from FilterBulkTransferred and is used to iterate over the raw logs and unpacked data for BulkTransferred events raised by the WalletMock contract.
type WalletMockBulkTransferredIterator struct {
	Event *WalletMockBulkTransferred // Event containing the contract specifics and raw log

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
func (it *WalletMockBulkTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletMockBulkTransferred)
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
		it.Event = new(WalletMockBulkTransferred)
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
func (it *WalletMockBulkTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletMockBulkTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletMockBulkTransferred represents a BulkTransferred event raised by the WalletMock contract.
type WalletMockBulkTransferred struct {
	To     common.Address
	Assets []common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBulkTransferred is a free log retrieval operation binding the contract event 0xd4f62f23021706247dcffea245d104ae7ddaec7f23acf3d11d7136d5de6a69ad.
//
// Solidity: event BulkTransferred(address _to, address[] _assets)
func (_WalletMock *WalletMockFilterer) FilterBulkTransferred(opts *bind.FilterOpts) (*WalletMockBulkTransferredIterator, error) {

	logs, sub, err := _WalletMock.contract.FilterLogs(opts, "BulkTransferred")
	if err != nil {
		return nil, err
	}
	return &WalletMockBulkTransferredIterator{contract: _WalletMock.contract, event: "BulkTransferred", logs: logs, sub: sub}, nil
}

// WatchBulkTransferred is a free log subscription operation binding the contract event 0xd4f62f23021706247dcffea245d104ae7ddaec7f23acf3d11d7136d5de6a69ad.
//
// Solidity: event BulkTransferred(address _to, address[] _assets)
func (_WalletMock *WalletMockFilterer) WatchBulkTransferred(opts *bind.WatchOpts, sink chan<- *WalletMockBulkTransferred) (event.Subscription, error) {

	logs, sub, err := _WalletMock.contract.WatchLogs(opts, "BulkTransferred")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletMockBulkTransferred)
				if err := _WalletMock.contract.UnpackLog(event, "BulkTransferred", log); err != nil {
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
func (_WalletMock *WalletMockFilterer) ParseBulkTransferred(log types.Log) (*WalletMockBulkTransferred, error) {
	event := new(WalletMockBulkTransferred)
	if err := _WalletMock.contract.UnpackLog(event, "BulkTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WalletMockCancelledWhitelistAdditionIterator is returned from FilterCancelledWhitelistAddition and is used to iterate over the raw logs and unpacked data for CancelledWhitelistAddition events raised by the WalletMock contract.
type WalletMockCancelledWhitelistAdditionIterator struct {
	Event *WalletMockCancelledWhitelistAddition // Event containing the contract specifics and raw log

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
func (it *WalletMockCancelledWhitelistAdditionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletMockCancelledWhitelistAddition)
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
		it.Event = new(WalletMockCancelledWhitelistAddition)
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
func (it *WalletMockCancelledWhitelistAdditionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletMockCancelledWhitelistAdditionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletMockCancelledWhitelistAddition represents a CancelledWhitelistAddition event raised by the WalletMock contract.
type WalletMockCancelledWhitelistAddition struct {
	Sender common.Address
	Hash   [32]byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCancelledWhitelistAddition is a free log retrieval operation binding the contract event 0x7794eff834d760583543e6e510e717a5e66d2c064e225f4db448343c3e66afcf.
//
// Solidity: event CancelledWhitelistAddition(address _sender, bytes32 _hash)
func (_WalletMock *WalletMockFilterer) FilterCancelledWhitelistAddition(opts *bind.FilterOpts) (*WalletMockCancelledWhitelistAdditionIterator, error) {

	logs, sub, err := _WalletMock.contract.FilterLogs(opts, "CancelledWhitelistAddition")
	if err != nil {
		return nil, err
	}
	return &WalletMockCancelledWhitelistAdditionIterator{contract: _WalletMock.contract, event: "CancelledWhitelistAddition", logs: logs, sub: sub}, nil
}

// WatchCancelledWhitelistAddition is a free log subscription operation binding the contract event 0x7794eff834d760583543e6e510e717a5e66d2c064e225f4db448343c3e66afcf.
//
// Solidity: event CancelledWhitelistAddition(address _sender, bytes32 _hash)
func (_WalletMock *WalletMockFilterer) WatchCancelledWhitelistAddition(opts *bind.WatchOpts, sink chan<- *WalletMockCancelledWhitelistAddition) (event.Subscription, error) {

	logs, sub, err := _WalletMock.contract.WatchLogs(opts, "CancelledWhitelistAddition")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletMockCancelledWhitelistAddition)
				if err := _WalletMock.contract.UnpackLog(event, "CancelledWhitelistAddition", log); err != nil {
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
func (_WalletMock *WalletMockFilterer) ParseCancelledWhitelistAddition(log types.Log) (*WalletMockCancelledWhitelistAddition, error) {
	event := new(WalletMockCancelledWhitelistAddition)
	if err := _WalletMock.contract.UnpackLog(event, "CancelledWhitelistAddition", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WalletMockCancelledWhitelistRemovalIterator is returned from FilterCancelledWhitelistRemoval and is used to iterate over the raw logs and unpacked data for CancelledWhitelistRemoval events raised by the WalletMock contract.
type WalletMockCancelledWhitelistRemovalIterator struct {
	Event *WalletMockCancelledWhitelistRemoval // Event containing the contract specifics and raw log

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
func (it *WalletMockCancelledWhitelistRemovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletMockCancelledWhitelistRemoval)
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
		it.Event = new(WalletMockCancelledWhitelistRemoval)
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
func (it *WalletMockCancelledWhitelistRemovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletMockCancelledWhitelistRemovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletMockCancelledWhitelistRemoval represents a CancelledWhitelistRemoval event raised by the WalletMock contract.
type WalletMockCancelledWhitelistRemoval struct {
	Sender common.Address
	Hash   [32]byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCancelledWhitelistRemoval is a free log retrieval operation binding the contract event 0x13c935eb475aa0f6e931fece83e2ac44569ce2d53460d29a6dedab40b965c8a3.
//
// Solidity: event CancelledWhitelistRemoval(address _sender, bytes32 _hash)
func (_WalletMock *WalletMockFilterer) FilterCancelledWhitelistRemoval(opts *bind.FilterOpts) (*WalletMockCancelledWhitelistRemovalIterator, error) {

	logs, sub, err := _WalletMock.contract.FilterLogs(opts, "CancelledWhitelistRemoval")
	if err != nil {
		return nil, err
	}
	return &WalletMockCancelledWhitelistRemovalIterator{contract: _WalletMock.contract, event: "CancelledWhitelistRemoval", logs: logs, sub: sub}, nil
}

// WatchCancelledWhitelistRemoval is a free log subscription operation binding the contract event 0x13c935eb475aa0f6e931fece83e2ac44569ce2d53460d29a6dedab40b965c8a3.
//
// Solidity: event CancelledWhitelistRemoval(address _sender, bytes32 _hash)
func (_WalletMock *WalletMockFilterer) WatchCancelledWhitelistRemoval(opts *bind.WatchOpts, sink chan<- *WalletMockCancelledWhitelistRemoval) (event.Subscription, error) {

	logs, sub, err := _WalletMock.contract.WatchLogs(opts, "CancelledWhitelistRemoval")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletMockCancelledWhitelistRemoval)
				if err := _WalletMock.contract.UnpackLog(event, "CancelledWhitelistRemoval", log); err != nil {
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
func (_WalletMock *WalletMockFilterer) ParseCancelledWhitelistRemoval(log types.Log) (*WalletMockCancelledWhitelistRemoval, error) {
	event := new(WalletMockCancelledWhitelistRemoval)
	if err := _WalletMock.contract.UnpackLog(event, "CancelledWhitelistRemoval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WalletMockExecutedRelayedTransactionIterator is returned from FilterExecutedRelayedTransaction and is used to iterate over the raw logs and unpacked data for ExecutedRelayedTransaction events raised by the WalletMock contract.
type WalletMockExecutedRelayedTransactionIterator struct {
	Event *WalletMockExecutedRelayedTransaction // Event containing the contract specifics and raw log

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
func (it *WalletMockExecutedRelayedTransactionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletMockExecutedRelayedTransaction)
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
		it.Event = new(WalletMockExecutedRelayedTransaction)
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
func (it *WalletMockExecutedRelayedTransactionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletMockExecutedRelayedTransactionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletMockExecutedRelayedTransaction represents a ExecutedRelayedTransaction event raised by the WalletMock contract.
type WalletMockExecutedRelayedTransaction struct {
	Data       []byte
	Returndata []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterExecutedRelayedTransaction is a free log retrieval operation binding the contract event 0x823dbcf2b7b0f265871963ca65ac033f6b4c71e0d82cd123d2ff23d752dc21c1.
//
// Solidity: event ExecutedRelayedTransaction(bytes _data, bytes _returndata)
func (_WalletMock *WalletMockFilterer) FilterExecutedRelayedTransaction(opts *bind.FilterOpts) (*WalletMockExecutedRelayedTransactionIterator, error) {

	logs, sub, err := _WalletMock.contract.FilterLogs(opts, "ExecutedRelayedTransaction")
	if err != nil {
		return nil, err
	}
	return &WalletMockExecutedRelayedTransactionIterator{contract: _WalletMock.contract, event: "ExecutedRelayedTransaction", logs: logs, sub: sub}, nil
}

// WatchExecutedRelayedTransaction is a free log subscription operation binding the contract event 0x823dbcf2b7b0f265871963ca65ac033f6b4c71e0d82cd123d2ff23d752dc21c1.
//
// Solidity: event ExecutedRelayedTransaction(bytes _data, bytes _returndata)
func (_WalletMock *WalletMockFilterer) WatchExecutedRelayedTransaction(opts *bind.WatchOpts, sink chan<- *WalletMockExecutedRelayedTransaction) (event.Subscription, error) {

	logs, sub, err := _WalletMock.contract.WatchLogs(opts, "ExecutedRelayedTransaction")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletMockExecutedRelayedTransaction)
				if err := _WalletMock.contract.UnpackLog(event, "ExecutedRelayedTransaction", log); err != nil {
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
func (_WalletMock *WalletMockFilterer) ParseExecutedRelayedTransaction(log types.Log) (*WalletMockExecutedRelayedTransaction, error) {
	event := new(WalletMockExecutedRelayedTransaction)
	if err := _WalletMock.contract.UnpackLog(event, "ExecutedRelayedTransaction", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WalletMockExecutedTransactionIterator is returned from FilterExecutedTransaction and is used to iterate over the raw logs and unpacked data for ExecutedTransaction events raised by the WalletMock contract.
type WalletMockExecutedTransactionIterator struct {
	Event *WalletMockExecutedTransaction // Event containing the contract specifics and raw log

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
func (it *WalletMockExecutedTransactionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletMockExecutedTransaction)
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
		it.Event = new(WalletMockExecutedTransaction)
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
func (it *WalletMockExecutedTransactionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletMockExecutedTransactionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletMockExecutedTransaction represents a ExecutedTransaction event raised by the WalletMock contract.
type WalletMockExecutedTransaction struct {
	Destination common.Address
	Value       *big.Int
	Data        []byte
	Returndata  []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterExecutedTransaction is a free log retrieval operation binding the contract event 0xf77753fab406ecfff96d6ff2476c64a838fa9f6d37b1bf190f8546e395e3b613.
//
// Solidity: event ExecutedTransaction(address _destination, uint256 _value, bytes _data, bytes _returndata)
func (_WalletMock *WalletMockFilterer) FilterExecutedTransaction(opts *bind.FilterOpts) (*WalletMockExecutedTransactionIterator, error) {

	logs, sub, err := _WalletMock.contract.FilterLogs(opts, "ExecutedTransaction")
	if err != nil {
		return nil, err
	}
	return &WalletMockExecutedTransactionIterator{contract: _WalletMock.contract, event: "ExecutedTransaction", logs: logs, sub: sub}, nil
}

// WatchExecutedTransaction is a free log subscription operation binding the contract event 0xf77753fab406ecfff96d6ff2476c64a838fa9f6d37b1bf190f8546e395e3b613.
//
// Solidity: event ExecutedTransaction(address _destination, uint256 _value, bytes _data, bytes _returndata)
func (_WalletMock *WalletMockFilterer) WatchExecutedTransaction(opts *bind.WatchOpts, sink chan<- *WalletMockExecutedTransaction) (event.Subscription, error) {

	logs, sub, err := _WalletMock.contract.WatchLogs(opts, "ExecutedTransaction")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletMockExecutedTransaction)
				if err := _WalletMock.contract.UnpackLog(event, "ExecutedTransaction", log); err != nil {
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
func (_WalletMock *WalletMockFilterer) ParseExecutedTransaction(log types.Log) (*WalletMockExecutedTransaction, error) {
	event := new(WalletMockExecutedTransaction)
	if err := _WalletMock.contract.UnpackLog(event, "ExecutedTransaction", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WalletMockIncreasedRelayNonceIterator is returned from FilterIncreasedRelayNonce and is used to iterate over the raw logs and unpacked data for IncreasedRelayNonce events raised by the WalletMock contract.
type WalletMockIncreasedRelayNonceIterator struct {
	Event *WalletMockIncreasedRelayNonce // Event containing the contract specifics and raw log

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
func (it *WalletMockIncreasedRelayNonceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletMockIncreasedRelayNonce)
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
		it.Event = new(WalletMockIncreasedRelayNonce)
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
func (it *WalletMockIncreasedRelayNonceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletMockIncreasedRelayNonceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletMockIncreasedRelayNonce represents a IncreasedRelayNonce event raised by the WalletMock contract.
type WalletMockIncreasedRelayNonce struct {
	Sender       common.Address
	CurrentNonce *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterIncreasedRelayNonce is a free log retrieval operation binding the contract event 0xab0423a75986556234aecd171c46ce7f5e45607d8070bf5230f2735b50322bff.
//
// Solidity: event IncreasedRelayNonce(address _sender, uint256 _currentNonce)
func (_WalletMock *WalletMockFilterer) FilterIncreasedRelayNonce(opts *bind.FilterOpts) (*WalletMockIncreasedRelayNonceIterator, error) {

	logs, sub, err := _WalletMock.contract.FilterLogs(opts, "IncreasedRelayNonce")
	if err != nil {
		return nil, err
	}
	return &WalletMockIncreasedRelayNonceIterator{contract: _WalletMock.contract, event: "IncreasedRelayNonce", logs: logs, sub: sub}, nil
}

// WatchIncreasedRelayNonce is a free log subscription operation binding the contract event 0xab0423a75986556234aecd171c46ce7f5e45607d8070bf5230f2735b50322bff.
//
// Solidity: event IncreasedRelayNonce(address _sender, uint256 _currentNonce)
func (_WalletMock *WalletMockFilterer) WatchIncreasedRelayNonce(opts *bind.WatchOpts, sink chan<- *WalletMockIncreasedRelayNonce) (event.Subscription, error) {

	logs, sub, err := _WalletMock.contract.WatchLogs(opts, "IncreasedRelayNonce")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletMockIncreasedRelayNonce)
				if err := _WalletMock.contract.UnpackLog(event, "IncreasedRelayNonce", log); err != nil {
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
func (_WalletMock *WalletMockFilterer) ParseIncreasedRelayNonce(log types.Log) (*WalletMockIncreasedRelayNonce, error) {
	event := new(WalletMockIncreasedRelayNonce)
	if err := _WalletMock.contract.UnpackLog(event, "IncreasedRelayNonce", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WalletMockLoadedTokenCardIterator is returned from FilterLoadedTokenCard and is used to iterate over the raw logs and unpacked data for LoadedTokenCard events raised by the WalletMock contract.
type WalletMockLoadedTokenCardIterator struct {
	Event *WalletMockLoadedTokenCard // Event containing the contract specifics and raw log

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
func (it *WalletMockLoadedTokenCardIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletMockLoadedTokenCard)
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
		it.Event = new(WalletMockLoadedTokenCard)
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
func (it *WalletMockLoadedTokenCardIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletMockLoadedTokenCardIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletMockLoadedTokenCard represents a LoadedTokenCard event raised by the WalletMock contract.
type WalletMockLoadedTokenCard struct {
	Asset  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLoadedTokenCard is a free log retrieval operation binding the contract event 0x5f65674bec9af81f71be68674135a0ea3f163fb91984e3893d06da9f6ea2ce8a.
//
// Solidity: event LoadedTokenCard(address _asset, uint256 _amount)
func (_WalletMock *WalletMockFilterer) FilterLoadedTokenCard(opts *bind.FilterOpts) (*WalletMockLoadedTokenCardIterator, error) {

	logs, sub, err := _WalletMock.contract.FilterLogs(opts, "LoadedTokenCard")
	if err != nil {
		return nil, err
	}
	return &WalletMockLoadedTokenCardIterator{contract: _WalletMock.contract, event: "LoadedTokenCard", logs: logs, sub: sub}, nil
}

// WatchLoadedTokenCard is a free log subscription operation binding the contract event 0x5f65674bec9af81f71be68674135a0ea3f163fb91984e3893d06da9f6ea2ce8a.
//
// Solidity: event LoadedTokenCard(address _asset, uint256 _amount)
func (_WalletMock *WalletMockFilterer) WatchLoadedTokenCard(opts *bind.WatchOpts, sink chan<- *WalletMockLoadedTokenCard) (event.Subscription, error) {

	logs, sub, err := _WalletMock.contract.WatchLogs(opts, "LoadedTokenCard")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletMockLoadedTokenCard)
				if err := _WalletMock.contract.UnpackLog(event, "LoadedTokenCard", log); err != nil {
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
func (_WalletMock *WalletMockFilterer) ParseLoadedTokenCard(log types.Log) (*WalletMockLoadedTokenCard, error) {
	event := new(WalletMockLoadedTokenCard)
	if err := _WalletMock.contract.UnpackLog(event, "LoadedTokenCard", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WalletMockLockedOwnershipIterator is returned from FilterLockedOwnership and is used to iterate over the raw logs and unpacked data for LockedOwnership events raised by the WalletMock contract.
type WalletMockLockedOwnershipIterator struct {
	Event *WalletMockLockedOwnership // Event containing the contract specifics and raw log

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
func (it *WalletMockLockedOwnershipIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletMockLockedOwnership)
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
		it.Event = new(WalletMockLockedOwnership)
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
func (it *WalletMockLockedOwnershipIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletMockLockedOwnershipIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletMockLockedOwnership represents a LockedOwnership event raised by the WalletMock contract.
type WalletMockLockedOwnership struct {
	Locked common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLockedOwnership is a free log retrieval operation binding the contract event 0x808639ff9c8e4732d60b6c2330de498035416d229f27a77d259680895efec122.
//
// Solidity: event LockedOwnership(address _locked)
func (_WalletMock *WalletMockFilterer) FilterLockedOwnership(opts *bind.FilterOpts) (*WalletMockLockedOwnershipIterator, error) {

	logs, sub, err := _WalletMock.contract.FilterLogs(opts, "LockedOwnership")
	if err != nil {
		return nil, err
	}
	return &WalletMockLockedOwnershipIterator{contract: _WalletMock.contract, event: "LockedOwnership", logs: logs, sub: sub}, nil
}

// WatchLockedOwnership is a free log subscription operation binding the contract event 0x808639ff9c8e4732d60b6c2330de498035416d229f27a77d259680895efec122.
//
// Solidity: event LockedOwnership(address _locked)
func (_WalletMock *WalletMockFilterer) WatchLockedOwnership(opts *bind.WatchOpts, sink chan<- *WalletMockLockedOwnership) (event.Subscription, error) {

	logs, sub, err := _WalletMock.contract.WatchLogs(opts, "LockedOwnership")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletMockLockedOwnership)
				if err := _WalletMock.contract.UnpackLog(event, "LockedOwnership", log); err != nil {
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
func (_WalletMock *WalletMockFilterer) ParseLockedOwnership(log types.Log) (*WalletMockLockedOwnership, error) {
	event := new(WalletMockLockedOwnership)
	if err := _WalletMock.contract.UnpackLog(event, "LockedOwnership", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WalletMockReceivedIterator is returned from FilterReceived and is used to iterate over the raw logs and unpacked data for Received events raised by the WalletMock contract.
type WalletMockReceivedIterator struct {
	Event *WalletMockReceived // Event containing the contract specifics and raw log

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
func (it *WalletMockReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletMockReceived)
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
		it.Event = new(WalletMockReceived)
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
func (it *WalletMockReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletMockReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletMockReceived represents a Received event raised by the WalletMock contract.
type WalletMockReceived struct {
	From   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterReceived is a free log retrieval operation binding the contract event 0x88a5966d370b9919b20f3e2c13ff65706f196a4e32cc2c12bf57088f88525874.
//
// Solidity: event Received(address _from, uint256 _amount)
func (_WalletMock *WalletMockFilterer) FilterReceived(opts *bind.FilterOpts) (*WalletMockReceivedIterator, error) {

	logs, sub, err := _WalletMock.contract.FilterLogs(opts, "Received")
	if err != nil {
		return nil, err
	}
	return &WalletMockReceivedIterator{contract: _WalletMock.contract, event: "Received", logs: logs, sub: sub}, nil
}

// WatchReceived is a free log subscription operation binding the contract event 0x88a5966d370b9919b20f3e2c13ff65706f196a4e32cc2c12bf57088f88525874.
//
// Solidity: event Received(address _from, uint256 _amount)
func (_WalletMock *WalletMockFilterer) WatchReceived(opts *bind.WatchOpts, sink chan<- *WalletMockReceived) (event.Subscription, error) {

	logs, sub, err := _WalletMock.contract.WatchLogs(opts, "Received")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletMockReceived)
				if err := _WalletMock.contract.UnpackLog(event, "Received", log); err != nil {
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
func (_WalletMock *WalletMockFilterer) ParseReceived(log types.Log) (*WalletMockReceived, error) {
	event := new(WalletMockReceived)
	if err := _WalletMock.contract.UnpackLog(event, "Received", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WalletMockRemovedFromWhitelistIterator is returned from FilterRemovedFromWhitelist and is used to iterate over the raw logs and unpacked data for RemovedFromWhitelist events raised by the WalletMock contract.
type WalletMockRemovedFromWhitelistIterator struct {
	Event *WalletMockRemovedFromWhitelist // Event containing the contract specifics and raw log

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
func (it *WalletMockRemovedFromWhitelistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletMockRemovedFromWhitelist)
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
		it.Event = new(WalletMockRemovedFromWhitelist)
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
func (it *WalletMockRemovedFromWhitelistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletMockRemovedFromWhitelistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletMockRemovedFromWhitelist represents a RemovedFromWhitelist event raised by the WalletMock contract.
type WalletMockRemovedFromWhitelist struct {
	Sender    common.Address
	Addresses []common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRemovedFromWhitelist is a free log retrieval operation binding the contract event 0xd218c430fa348f4ce67791021b6b89c0c3eacd4ead1d8f5b83c60038ec28249b.
//
// Solidity: event RemovedFromWhitelist(address _sender, address[] _addresses)
func (_WalletMock *WalletMockFilterer) FilterRemovedFromWhitelist(opts *bind.FilterOpts) (*WalletMockRemovedFromWhitelistIterator, error) {

	logs, sub, err := _WalletMock.contract.FilterLogs(opts, "RemovedFromWhitelist")
	if err != nil {
		return nil, err
	}
	return &WalletMockRemovedFromWhitelistIterator{contract: _WalletMock.contract, event: "RemovedFromWhitelist", logs: logs, sub: sub}, nil
}

// WatchRemovedFromWhitelist is a free log subscription operation binding the contract event 0xd218c430fa348f4ce67791021b6b89c0c3eacd4ead1d8f5b83c60038ec28249b.
//
// Solidity: event RemovedFromWhitelist(address _sender, address[] _addresses)
func (_WalletMock *WalletMockFilterer) WatchRemovedFromWhitelist(opts *bind.WatchOpts, sink chan<- *WalletMockRemovedFromWhitelist) (event.Subscription, error) {

	logs, sub, err := _WalletMock.contract.WatchLogs(opts, "RemovedFromWhitelist")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletMockRemovedFromWhitelist)
				if err := _WalletMock.contract.UnpackLog(event, "RemovedFromWhitelist", log); err != nil {
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
func (_WalletMock *WalletMockFilterer) ParseRemovedFromWhitelist(log types.Log) (*WalletMockRemovedFromWhitelist, error) {
	event := new(WalletMockRemovedFromWhitelist)
	if err := _WalletMock.contract.UnpackLog(event, "RemovedFromWhitelist", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WalletMockSetGasTopUpLimitIterator is returned from FilterSetGasTopUpLimit and is used to iterate over the raw logs and unpacked data for SetGasTopUpLimit events raised by the WalletMock contract.
type WalletMockSetGasTopUpLimitIterator struct {
	Event *WalletMockSetGasTopUpLimit // Event containing the contract specifics and raw log

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
func (it *WalletMockSetGasTopUpLimitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletMockSetGasTopUpLimit)
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
		it.Event = new(WalletMockSetGasTopUpLimit)
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
func (it *WalletMockSetGasTopUpLimitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletMockSetGasTopUpLimitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletMockSetGasTopUpLimit represents a SetGasTopUpLimit event raised by the WalletMock contract.
type WalletMockSetGasTopUpLimit struct {
	Sender common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetGasTopUpLimit is a free log retrieval operation binding the contract event 0x41ff5d5ce3b7935893a4e7269ec5caae9cca5e3bf0eb4b21d2f443489667112e.
//
// Solidity: event SetGasTopUpLimit(address _sender, uint256 _amount)
func (_WalletMock *WalletMockFilterer) FilterSetGasTopUpLimit(opts *bind.FilterOpts) (*WalletMockSetGasTopUpLimitIterator, error) {

	logs, sub, err := _WalletMock.contract.FilterLogs(opts, "SetGasTopUpLimit")
	if err != nil {
		return nil, err
	}
	return &WalletMockSetGasTopUpLimitIterator{contract: _WalletMock.contract, event: "SetGasTopUpLimit", logs: logs, sub: sub}, nil
}

// WatchSetGasTopUpLimit is a free log subscription operation binding the contract event 0x41ff5d5ce3b7935893a4e7269ec5caae9cca5e3bf0eb4b21d2f443489667112e.
//
// Solidity: event SetGasTopUpLimit(address _sender, uint256 _amount)
func (_WalletMock *WalletMockFilterer) WatchSetGasTopUpLimit(opts *bind.WatchOpts, sink chan<- *WalletMockSetGasTopUpLimit) (event.Subscription, error) {

	logs, sub, err := _WalletMock.contract.WatchLogs(opts, "SetGasTopUpLimit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletMockSetGasTopUpLimit)
				if err := _WalletMock.contract.UnpackLog(event, "SetGasTopUpLimit", log); err != nil {
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
func (_WalletMock *WalletMockFilterer) ParseSetGasTopUpLimit(log types.Log) (*WalletMockSetGasTopUpLimit, error) {
	event := new(WalletMockSetGasTopUpLimit)
	if err := _WalletMock.contract.UnpackLog(event, "SetGasTopUpLimit", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WalletMockSetLoadLimitIterator is returned from FilterSetLoadLimit and is used to iterate over the raw logs and unpacked data for SetLoadLimit events raised by the WalletMock contract.
type WalletMockSetLoadLimitIterator struct {
	Event *WalletMockSetLoadLimit // Event containing the contract specifics and raw log

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
func (it *WalletMockSetLoadLimitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletMockSetLoadLimit)
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
		it.Event = new(WalletMockSetLoadLimit)
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
func (it *WalletMockSetLoadLimitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletMockSetLoadLimitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletMockSetLoadLimit represents a SetLoadLimit event raised by the WalletMock contract.
type WalletMockSetLoadLimit struct {
	Sender common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetLoadLimit is a free log retrieval operation binding the contract event 0x0b05243483e17c3f3377aee82b7d47e5700b48288695fc08b7ecc2759afa44ef.
//
// Solidity: event SetLoadLimit(address _sender, uint256 _amount)
func (_WalletMock *WalletMockFilterer) FilterSetLoadLimit(opts *bind.FilterOpts) (*WalletMockSetLoadLimitIterator, error) {

	logs, sub, err := _WalletMock.contract.FilterLogs(opts, "SetLoadLimit")
	if err != nil {
		return nil, err
	}
	return &WalletMockSetLoadLimitIterator{contract: _WalletMock.contract, event: "SetLoadLimit", logs: logs, sub: sub}, nil
}

// WatchSetLoadLimit is a free log subscription operation binding the contract event 0x0b05243483e17c3f3377aee82b7d47e5700b48288695fc08b7ecc2759afa44ef.
//
// Solidity: event SetLoadLimit(address _sender, uint256 _amount)
func (_WalletMock *WalletMockFilterer) WatchSetLoadLimit(opts *bind.WatchOpts, sink chan<- *WalletMockSetLoadLimit) (event.Subscription, error) {

	logs, sub, err := _WalletMock.contract.WatchLogs(opts, "SetLoadLimit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletMockSetLoadLimit)
				if err := _WalletMock.contract.UnpackLog(event, "SetLoadLimit", log); err != nil {
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
func (_WalletMock *WalletMockFilterer) ParseSetLoadLimit(log types.Log) (*WalletMockSetLoadLimit, error) {
	event := new(WalletMockSetLoadLimit)
	if err := _WalletMock.contract.UnpackLog(event, "SetLoadLimit", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WalletMockSetSpendLimitIterator is returned from FilterSetSpendLimit and is used to iterate over the raw logs and unpacked data for SetSpendLimit events raised by the WalletMock contract.
type WalletMockSetSpendLimitIterator struct {
	Event *WalletMockSetSpendLimit // Event containing the contract specifics and raw log

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
func (it *WalletMockSetSpendLimitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletMockSetSpendLimit)
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
		it.Event = new(WalletMockSetSpendLimit)
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
func (it *WalletMockSetSpendLimitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletMockSetSpendLimitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletMockSetSpendLimit represents a SetSpendLimit event raised by the WalletMock contract.
type WalletMockSetSpendLimit struct {
	Sender common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetSpendLimit is a free log retrieval operation binding the contract event 0x068f112e5ec923d412be64779fe69e0fcbb6784c6617e94cccc8fd348f2e0f21.
//
// Solidity: event SetSpendLimit(address _sender, uint256 _amount)
func (_WalletMock *WalletMockFilterer) FilterSetSpendLimit(opts *bind.FilterOpts) (*WalletMockSetSpendLimitIterator, error) {

	logs, sub, err := _WalletMock.contract.FilterLogs(opts, "SetSpendLimit")
	if err != nil {
		return nil, err
	}
	return &WalletMockSetSpendLimitIterator{contract: _WalletMock.contract, event: "SetSpendLimit", logs: logs, sub: sub}, nil
}

// WatchSetSpendLimit is a free log subscription operation binding the contract event 0x068f112e5ec923d412be64779fe69e0fcbb6784c6617e94cccc8fd348f2e0f21.
//
// Solidity: event SetSpendLimit(address _sender, uint256 _amount)
func (_WalletMock *WalletMockFilterer) WatchSetSpendLimit(opts *bind.WatchOpts, sink chan<- *WalletMockSetSpendLimit) (event.Subscription, error) {

	logs, sub, err := _WalletMock.contract.WatchLogs(opts, "SetSpendLimit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletMockSetSpendLimit)
				if err := _WalletMock.contract.UnpackLog(event, "SetSpendLimit", log); err != nil {
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
func (_WalletMock *WalletMockFilterer) ParseSetSpendLimit(log types.Log) (*WalletMockSetSpendLimit, error) {
	event := new(WalletMockSetSpendLimit)
	if err := _WalletMock.contract.UnpackLog(event, "SetSpendLimit", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WalletMockSubmittedGasTopUpLimitUpdateIterator is returned from FilterSubmittedGasTopUpLimitUpdate and is used to iterate over the raw logs and unpacked data for SubmittedGasTopUpLimitUpdate events raised by the WalletMock contract.
type WalletMockSubmittedGasTopUpLimitUpdateIterator struct {
	Event *WalletMockSubmittedGasTopUpLimitUpdate // Event containing the contract specifics and raw log

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
func (it *WalletMockSubmittedGasTopUpLimitUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletMockSubmittedGasTopUpLimitUpdate)
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
		it.Event = new(WalletMockSubmittedGasTopUpLimitUpdate)
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
func (it *WalletMockSubmittedGasTopUpLimitUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletMockSubmittedGasTopUpLimitUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletMockSubmittedGasTopUpLimitUpdate represents a SubmittedGasTopUpLimitUpdate event raised by the WalletMock contract.
type WalletMockSubmittedGasTopUpLimitUpdate struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSubmittedGasTopUpLimitUpdate is a free log retrieval operation binding the contract event 0xaf2a77cd04c3cc155588dd3bf67b310ab4fb3b1da3cf6b8d7d4d2aa1d09b794c.
//
// Solidity: event SubmittedGasTopUpLimitUpdate(uint256 _amount)
func (_WalletMock *WalletMockFilterer) FilterSubmittedGasTopUpLimitUpdate(opts *bind.FilterOpts) (*WalletMockSubmittedGasTopUpLimitUpdateIterator, error) {

	logs, sub, err := _WalletMock.contract.FilterLogs(opts, "SubmittedGasTopUpLimitUpdate")
	if err != nil {
		return nil, err
	}
	return &WalletMockSubmittedGasTopUpLimitUpdateIterator{contract: _WalletMock.contract, event: "SubmittedGasTopUpLimitUpdate", logs: logs, sub: sub}, nil
}

// WatchSubmittedGasTopUpLimitUpdate is a free log subscription operation binding the contract event 0xaf2a77cd04c3cc155588dd3bf67b310ab4fb3b1da3cf6b8d7d4d2aa1d09b794c.
//
// Solidity: event SubmittedGasTopUpLimitUpdate(uint256 _amount)
func (_WalletMock *WalletMockFilterer) WatchSubmittedGasTopUpLimitUpdate(opts *bind.WatchOpts, sink chan<- *WalletMockSubmittedGasTopUpLimitUpdate) (event.Subscription, error) {

	logs, sub, err := _WalletMock.contract.WatchLogs(opts, "SubmittedGasTopUpLimitUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletMockSubmittedGasTopUpLimitUpdate)
				if err := _WalletMock.contract.UnpackLog(event, "SubmittedGasTopUpLimitUpdate", log); err != nil {
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
func (_WalletMock *WalletMockFilterer) ParseSubmittedGasTopUpLimitUpdate(log types.Log) (*WalletMockSubmittedGasTopUpLimitUpdate, error) {
	event := new(WalletMockSubmittedGasTopUpLimitUpdate)
	if err := _WalletMock.contract.UnpackLog(event, "SubmittedGasTopUpLimitUpdate", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WalletMockSubmittedLoadLimitUpdateIterator is returned from FilterSubmittedLoadLimitUpdate and is used to iterate over the raw logs and unpacked data for SubmittedLoadLimitUpdate events raised by the WalletMock contract.
type WalletMockSubmittedLoadLimitUpdateIterator struct {
	Event *WalletMockSubmittedLoadLimitUpdate // Event containing the contract specifics and raw log

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
func (it *WalletMockSubmittedLoadLimitUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletMockSubmittedLoadLimitUpdate)
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
		it.Event = new(WalletMockSubmittedLoadLimitUpdate)
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
func (it *WalletMockSubmittedLoadLimitUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletMockSubmittedLoadLimitUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletMockSubmittedLoadLimitUpdate represents a SubmittedLoadLimitUpdate event raised by the WalletMock contract.
type WalletMockSubmittedLoadLimitUpdate struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSubmittedLoadLimitUpdate is a free log retrieval operation binding the contract event 0xc178d379965e5657b6fc57494e392f121a14119215dfb422aad7db4cc03f2d10.
//
// Solidity: event SubmittedLoadLimitUpdate(uint256 _amount)
func (_WalletMock *WalletMockFilterer) FilterSubmittedLoadLimitUpdate(opts *bind.FilterOpts) (*WalletMockSubmittedLoadLimitUpdateIterator, error) {

	logs, sub, err := _WalletMock.contract.FilterLogs(opts, "SubmittedLoadLimitUpdate")
	if err != nil {
		return nil, err
	}
	return &WalletMockSubmittedLoadLimitUpdateIterator{contract: _WalletMock.contract, event: "SubmittedLoadLimitUpdate", logs: logs, sub: sub}, nil
}

// WatchSubmittedLoadLimitUpdate is a free log subscription operation binding the contract event 0xc178d379965e5657b6fc57494e392f121a14119215dfb422aad7db4cc03f2d10.
//
// Solidity: event SubmittedLoadLimitUpdate(uint256 _amount)
func (_WalletMock *WalletMockFilterer) WatchSubmittedLoadLimitUpdate(opts *bind.WatchOpts, sink chan<- *WalletMockSubmittedLoadLimitUpdate) (event.Subscription, error) {

	logs, sub, err := _WalletMock.contract.WatchLogs(opts, "SubmittedLoadLimitUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletMockSubmittedLoadLimitUpdate)
				if err := _WalletMock.contract.UnpackLog(event, "SubmittedLoadLimitUpdate", log); err != nil {
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
func (_WalletMock *WalletMockFilterer) ParseSubmittedLoadLimitUpdate(log types.Log) (*WalletMockSubmittedLoadLimitUpdate, error) {
	event := new(WalletMockSubmittedLoadLimitUpdate)
	if err := _WalletMock.contract.UnpackLog(event, "SubmittedLoadLimitUpdate", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WalletMockSubmittedSpendLimitUpdateIterator is returned from FilterSubmittedSpendLimitUpdate and is used to iterate over the raw logs and unpacked data for SubmittedSpendLimitUpdate events raised by the WalletMock contract.
type WalletMockSubmittedSpendLimitUpdateIterator struct {
	Event *WalletMockSubmittedSpendLimitUpdate // Event containing the contract specifics and raw log

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
func (it *WalletMockSubmittedSpendLimitUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletMockSubmittedSpendLimitUpdate)
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
		it.Event = new(WalletMockSubmittedSpendLimitUpdate)
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
func (it *WalletMockSubmittedSpendLimitUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletMockSubmittedSpendLimitUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletMockSubmittedSpendLimitUpdate represents a SubmittedSpendLimitUpdate event raised by the WalletMock contract.
type WalletMockSubmittedSpendLimitUpdate struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSubmittedSpendLimitUpdate is a free log retrieval operation binding the contract event 0x4b1b970c8a0fa761e7803ed70c13d7aca71904b13df60fbe03f981da1730da91.
//
// Solidity: event SubmittedSpendLimitUpdate(uint256 _amount)
func (_WalletMock *WalletMockFilterer) FilterSubmittedSpendLimitUpdate(opts *bind.FilterOpts) (*WalletMockSubmittedSpendLimitUpdateIterator, error) {

	logs, sub, err := _WalletMock.contract.FilterLogs(opts, "SubmittedSpendLimitUpdate")
	if err != nil {
		return nil, err
	}
	return &WalletMockSubmittedSpendLimitUpdateIterator{contract: _WalletMock.contract, event: "SubmittedSpendLimitUpdate", logs: logs, sub: sub}, nil
}

// WatchSubmittedSpendLimitUpdate is a free log subscription operation binding the contract event 0x4b1b970c8a0fa761e7803ed70c13d7aca71904b13df60fbe03f981da1730da91.
//
// Solidity: event SubmittedSpendLimitUpdate(uint256 _amount)
func (_WalletMock *WalletMockFilterer) WatchSubmittedSpendLimitUpdate(opts *bind.WatchOpts, sink chan<- *WalletMockSubmittedSpendLimitUpdate) (event.Subscription, error) {

	logs, sub, err := _WalletMock.contract.WatchLogs(opts, "SubmittedSpendLimitUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletMockSubmittedSpendLimitUpdate)
				if err := _WalletMock.contract.UnpackLog(event, "SubmittedSpendLimitUpdate", log); err != nil {
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
func (_WalletMock *WalletMockFilterer) ParseSubmittedSpendLimitUpdate(log types.Log) (*WalletMockSubmittedSpendLimitUpdate, error) {
	event := new(WalletMockSubmittedSpendLimitUpdate)
	if err := _WalletMock.contract.UnpackLog(event, "SubmittedSpendLimitUpdate", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WalletMockSubmittedWhitelistAdditionIterator is returned from FilterSubmittedWhitelistAddition and is used to iterate over the raw logs and unpacked data for SubmittedWhitelistAddition events raised by the WalletMock contract.
type WalletMockSubmittedWhitelistAdditionIterator struct {
	Event *WalletMockSubmittedWhitelistAddition // Event containing the contract specifics and raw log

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
func (it *WalletMockSubmittedWhitelistAdditionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletMockSubmittedWhitelistAddition)
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
		it.Event = new(WalletMockSubmittedWhitelistAddition)
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
func (it *WalletMockSubmittedWhitelistAdditionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletMockSubmittedWhitelistAdditionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletMockSubmittedWhitelistAddition represents a SubmittedWhitelistAddition event raised by the WalletMock contract.
type WalletMockSubmittedWhitelistAddition struct {
	Addresses []common.Address
	Hash      [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSubmittedWhitelistAddition is a free log retrieval operation binding the contract event 0x9c80b3b5f68b3e017766d59e8d09b34efe6462b05c398f35cab9e271d9bc3b9c.
//
// Solidity: event SubmittedWhitelistAddition(address[] _addresses, bytes32 _hash)
func (_WalletMock *WalletMockFilterer) FilterSubmittedWhitelistAddition(opts *bind.FilterOpts) (*WalletMockSubmittedWhitelistAdditionIterator, error) {

	logs, sub, err := _WalletMock.contract.FilterLogs(opts, "SubmittedWhitelistAddition")
	if err != nil {
		return nil, err
	}
	return &WalletMockSubmittedWhitelistAdditionIterator{contract: _WalletMock.contract, event: "SubmittedWhitelistAddition", logs: logs, sub: sub}, nil
}

// WatchSubmittedWhitelistAddition is a free log subscription operation binding the contract event 0x9c80b3b5f68b3e017766d59e8d09b34efe6462b05c398f35cab9e271d9bc3b9c.
//
// Solidity: event SubmittedWhitelistAddition(address[] _addresses, bytes32 _hash)
func (_WalletMock *WalletMockFilterer) WatchSubmittedWhitelistAddition(opts *bind.WatchOpts, sink chan<- *WalletMockSubmittedWhitelistAddition) (event.Subscription, error) {

	logs, sub, err := _WalletMock.contract.WatchLogs(opts, "SubmittedWhitelistAddition")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletMockSubmittedWhitelistAddition)
				if err := _WalletMock.contract.UnpackLog(event, "SubmittedWhitelistAddition", log); err != nil {
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
func (_WalletMock *WalletMockFilterer) ParseSubmittedWhitelistAddition(log types.Log) (*WalletMockSubmittedWhitelistAddition, error) {
	event := new(WalletMockSubmittedWhitelistAddition)
	if err := _WalletMock.contract.UnpackLog(event, "SubmittedWhitelistAddition", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WalletMockSubmittedWhitelistRemovalIterator is returned from FilterSubmittedWhitelistRemoval and is used to iterate over the raw logs and unpacked data for SubmittedWhitelistRemoval events raised by the WalletMock contract.
type WalletMockSubmittedWhitelistRemovalIterator struct {
	Event *WalletMockSubmittedWhitelistRemoval // Event containing the contract specifics and raw log

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
func (it *WalletMockSubmittedWhitelistRemovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletMockSubmittedWhitelistRemoval)
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
		it.Event = new(WalletMockSubmittedWhitelistRemoval)
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
func (it *WalletMockSubmittedWhitelistRemovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletMockSubmittedWhitelistRemovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletMockSubmittedWhitelistRemoval represents a SubmittedWhitelistRemoval event raised by the WalletMock contract.
type WalletMockSubmittedWhitelistRemoval struct {
	Addresses []common.Address
	Hash      [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSubmittedWhitelistRemoval is a free log retrieval operation binding the contract event 0xfbc0e5ca6c7e4858daf0fdb185ef5186203e74ec9c64737e93c0aeaec596e1d1.
//
// Solidity: event SubmittedWhitelistRemoval(address[] _addresses, bytes32 _hash)
func (_WalletMock *WalletMockFilterer) FilterSubmittedWhitelistRemoval(opts *bind.FilterOpts) (*WalletMockSubmittedWhitelistRemovalIterator, error) {

	logs, sub, err := _WalletMock.contract.FilterLogs(opts, "SubmittedWhitelistRemoval")
	if err != nil {
		return nil, err
	}
	return &WalletMockSubmittedWhitelistRemovalIterator{contract: _WalletMock.contract, event: "SubmittedWhitelistRemoval", logs: logs, sub: sub}, nil
}

// WatchSubmittedWhitelistRemoval is a free log subscription operation binding the contract event 0xfbc0e5ca6c7e4858daf0fdb185ef5186203e74ec9c64737e93c0aeaec596e1d1.
//
// Solidity: event SubmittedWhitelistRemoval(address[] _addresses, bytes32 _hash)
func (_WalletMock *WalletMockFilterer) WatchSubmittedWhitelistRemoval(opts *bind.WatchOpts, sink chan<- *WalletMockSubmittedWhitelistRemoval) (event.Subscription, error) {

	logs, sub, err := _WalletMock.contract.WatchLogs(opts, "SubmittedWhitelistRemoval")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletMockSubmittedWhitelistRemoval)
				if err := _WalletMock.contract.UnpackLog(event, "SubmittedWhitelistRemoval", log); err != nil {
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
func (_WalletMock *WalletMockFilterer) ParseSubmittedWhitelistRemoval(log types.Log) (*WalletMockSubmittedWhitelistRemoval, error) {
	event := new(WalletMockSubmittedWhitelistRemoval)
	if err := _WalletMock.contract.UnpackLog(event, "SubmittedWhitelistRemoval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WalletMockToppedUpGasIterator is returned from FilterToppedUpGas and is used to iterate over the raw logs and unpacked data for ToppedUpGas events raised by the WalletMock contract.
type WalletMockToppedUpGasIterator struct {
	Event *WalletMockToppedUpGas // Event containing the contract specifics and raw log

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
func (it *WalletMockToppedUpGasIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletMockToppedUpGas)
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
		it.Event = new(WalletMockToppedUpGas)
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
func (it *WalletMockToppedUpGasIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletMockToppedUpGasIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletMockToppedUpGas represents a ToppedUpGas event raised by the WalletMock contract.
type WalletMockToppedUpGas struct {
	Sender common.Address
	Owner  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterToppedUpGas is a free log retrieval operation binding the contract event 0x611b7c0d84fda988026215bef9b3e4d81cbceced7e679be6d5e044b588467c0e.
//
// Solidity: event ToppedUpGas(address _sender, address _owner, uint256 _amount)
func (_WalletMock *WalletMockFilterer) FilterToppedUpGas(opts *bind.FilterOpts) (*WalletMockToppedUpGasIterator, error) {

	logs, sub, err := _WalletMock.contract.FilterLogs(opts, "ToppedUpGas")
	if err != nil {
		return nil, err
	}
	return &WalletMockToppedUpGasIterator{contract: _WalletMock.contract, event: "ToppedUpGas", logs: logs, sub: sub}, nil
}

// WatchToppedUpGas is a free log subscription operation binding the contract event 0x611b7c0d84fda988026215bef9b3e4d81cbceced7e679be6d5e044b588467c0e.
//
// Solidity: event ToppedUpGas(address _sender, address _owner, uint256 _amount)
func (_WalletMock *WalletMockFilterer) WatchToppedUpGas(opts *bind.WatchOpts, sink chan<- *WalletMockToppedUpGas) (event.Subscription, error) {

	logs, sub, err := _WalletMock.contract.WatchLogs(opts, "ToppedUpGas")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletMockToppedUpGas)
				if err := _WalletMock.contract.UnpackLog(event, "ToppedUpGas", log); err != nil {
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
func (_WalletMock *WalletMockFilterer) ParseToppedUpGas(log types.Log) (*WalletMockToppedUpGas, error) {
	event := new(WalletMockToppedUpGas)
	if err := _WalletMock.contract.UnpackLog(event, "ToppedUpGas", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WalletMockTransferredIterator is returned from FilterTransferred and is used to iterate over the raw logs and unpacked data for Transferred events raised by the WalletMock contract.
type WalletMockTransferredIterator struct {
	Event *WalletMockTransferred // Event containing the contract specifics and raw log

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
func (it *WalletMockTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletMockTransferred)
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
		it.Event = new(WalletMockTransferred)
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
func (it *WalletMockTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletMockTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletMockTransferred represents a Transferred event raised by the WalletMock contract.
type WalletMockTransferred struct {
	To     common.Address
	Asset  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransferred is a free log retrieval operation binding the contract event 0xd1ba4ac2e2a11b5101f6cb4d978f514a155b421e8ec396d2d9abaf0bb02917ee.
//
// Solidity: event Transferred(address _to, address _asset, uint256 _amount)
func (_WalletMock *WalletMockFilterer) FilterTransferred(opts *bind.FilterOpts) (*WalletMockTransferredIterator, error) {

	logs, sub, err := _WalletMock.contract.FilterLogs(opts, "Transferred")
	if err != nil {
		return nil, err
	}
	return &WalletMockTransferredIterator{contract: _WalletMock.contract, event: "Transferred", logs: logs, sub: sub}, nil
}

// WatchTransferred is a free log subscription operation binding the contract event 0xd1ba4ac2e2a11b5101f6cb4d978f514a155b421e8ec396d2d9abaf0bb02917ee.
//
// Solidity: event Transferred(address _to, address _asset, uint256 _amount)
func (_WalletMock *WalletMockFilterer) WatchTransferred(opts *bind.WatchOpts, sink chan<- *WalletMockTransferred) (event.Subscription, error) {

	logs, sub, err := _WalletMock.contract.WatchLogs(opts, "Transferred")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletMockTransferred)
				if err := _WalletMock.contract.UnpackLog(event, "Transferred", log); err != nil {
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
func (_WalletMock *WalletMockFilterer) ParseTransferred(log types.Log) (*WalletMockTransferred, error) {
	event := new(WalletMockTransferred)
	if err := _WalletMock.contract.UnpackLog(event, "Transferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WalletMockTransferredOwnershipIterator is returned from FilterTransferredOwnership and is used to iterate over the raw logs and unpacked data for TransferredOwnership events raised by the WalletMock contract.
type WalletMockTransferredOwnershipIterator struct {
	Event *WalletMockTransferredOwnership // Event containing the contract specifics and raw log

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
func (it *WalletMockTransferredOwnershipIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletMockTransferredOwnership)
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
		it.Event = new(WalletMockTransferredOwnership)
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
func (it *WalletMockTransferredOwnershipIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletMockTransferredOwnershipIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletMockTransferredOwnership represents a TransferredOwnership event raised by the WalletMock contract.
type WalletMockTransferredOwnership struct {
	From common.Address
	To   common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterTransferredOwnership is a free log retrieval operation binding the contract event 0x850b3df64837d7d518b45f5aa64d104652c3b80eb5b34a8e3d9eb666cb7cdea5.
//
// Solidity: event TransferredOwnership(address _from, address _to)
func (_WalletMock *WalletMockFilterer) FilterTransferredOwnership(opts *bind.FilterOpts) (*WalletMockTransferredOwnershipIterator, error) {

	logs, sub, err := _WalletMock.contract.FilterLogs(opts, "TransferredOwnership")
	if err != nil {
		return nil, err
	}
	return &WalletMockTransferredOwnershipIterator{contract: _WalletMock.contract, event: "TransferredOwnership", logs: logs, sub: sub}, nil
}

// WatchTransferredOwnership is a free log subscription operation binding the contract event 0x850b3df64837d7d518b45f5aa64d104652c3b80eb5b34a8e3d9eb666cb7cdea5.
//
// Solidity: event TransferredOwnership(address _from, address _to)
func (_WalletMock *WalletMockFilterer) WatchTransferredOwnership(opts *bind.WatchOpts, sink chan<- *WalletMockTransferredOwnership) (event.Subscription, error) {

	logs, sub, err := _WalletMock.contract.WatchLogs(opts, "TransferredOwnership")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletMockTransferredOwnership)
				if err := _WalletMock.contract.UnpackLog(event, "TransferredOwnership", log); err != nil {
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
func (_WalletMock *WalletMockFilterer) ParseTransferredOwnership(log types.Log) (*WalletMockTransferredOwnership, error) {
	event := new(WalletMockTransferredOwnership)
	if err := _WalletMock.contract.UnpackLog(event, "TransferredOwnership", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WalletMockUpdatedAvailableLimitIterator is returned from FilterUpdatedAvailableLimit and is used to iterate over the raw logs and unpacked data for UpdatedAvailableLimit events raised by the WalletMock contract.
type WalletMockUpdatedAvailableLimitIterator struct {
	Event *WalletMockUpdatedAvailableLimit // Event containing the contract specifics and raw log

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
func (it *WalletMockUpdatedAvailableLimitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletMockUpdatedAvailableLimit)
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
		it.Event = new(WalletMockUpdatedAvailableLimit)
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
func (it *WalletMockUpdatedAvailableLimitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletMockUpdatedAvailableLimitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletMockUpdatedAvailableLimit represents a UpdatedAvailableLimit event raised by the WalletMock contract.
type WalletMockUpdatedAvailableLimit struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUpdatedAvailableLimit is a free log retrieval operation binding the contract event 0xe93bc25276d408d390778e7a8b926f2f67209c43ed540081b951fe128f0d3cd2.
//
// Solidity: event UpdatedAvailableLimit()
func (_WalletMock *WalletMockFilterer) FilterUpdatedAvailableLimit(opts *bind.FilterOpts) (*WalletMockUpdatedAvailableLimitIterator, error) {

	logs, sub, err := _WalletMock.contract.FilterLogs(opts, "UpdatedAvailableLimit")
	if err != nil {
		return nil, err
	}
	return &WalletMockUpdatedAvailableLimitIterator{contract: _WalletMock.contract, event: "UpdatedAvailableLimit", logs: logs, sub: sub}, nil
}

// WatchUpdatedAvailableLimit is a free log subscription operation binding the contract event 0xe93bc25276d408d390778e7a8b926f2f67209c43ed540081b951fe128f0d3cd2.
//
// Solidity: event UpdatedAvailableLimit()
func (_WalletMock *WalletMockFilterer) WatchUpdatedAvailableLimit(opts *bind.WatchOpts, sink chan<- *WalletMockUpdatedAvailableLimit) (event.Subscription, error) {

	logs, sub, err := _WalletMock.contract.WatchLogs(opts, "UpdatedAvailableLimit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletMockUpdatedAvailableLimit)
				if err := _WalletMock.contract.UnpackLog(event, "UpdatedAvailableLimit", log); err != nil {
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
func (_WalletMock *WalletMockFilterer) ParseUpdatedAvailableLimit(log types.Log) (*WalletMockUpdatedAvailableLimit, error) {
	event := new(WalletMockUpdatedAvailableLimit)
	if err := _WalletMock.contract.UnpackLog(event, "UpdatedAvailableLimit", log); err != nil {
		return nil, err
	}
	return event, nil
}
