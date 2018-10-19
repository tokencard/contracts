// Package contracts provides extra functionality in addition to the generated wallet bindings.

package contracts

import (
	"context"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/tokencard/contracts/pkg/bindings"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
)

var walletABI abi.ABI

func init() {
	parsed, err := abi.JSON(strings.NewReader(bindings.WalletABI))
	if err != nil {
		panic(err)
	}
	walletABI = parsed
}

var (
	ErrFailedContractCall = errors.New("calling smart contract failed")
	ErrInvalidEventData   = errors.New("event data could not be parsed")
	ErrNoCode             = errors.New("no contract code at given address")
)

// Event represents a raw smart contract event.
type Event struct {
	Address   common.Address
	TxHash    common.Hash
	BlockHash common.Hash
	Data      [][]byte
}

type SignTransaction func(context.Context, *types.Transaction) (*types.Transaction, error)

type ConstructOpts struct {
	From     common.Address  // Ethereum account to send the transaction from
	Nonce    uint64          // Nonce to use for the transaction execution
	Value    *big.Int        // Funds to transfer along along the transaction (nil = 0 = no funds)
	GasPrice *big.Int        // Gas price to use for the transaction execution (nil = gas price oracle)
	GasLimit uint64          // Gas limit to set for the transaction execution (0 = estimate)
	Sign     SignTransaction // Method to use for signing the transaction (nil = return unsigned transaction)
	Context  context.Context // Network context to support cancellation and timeouts
}

func NewWallet(ethereum *ethclient.Client, address common.Address) (*Wallet, error) {
	contractBindings, err := bindings.NewWallet(address, ethereum)
	if err != nil {
		return nil, err
	}
	contractABI, err := abi.JSON(strings.NewReader(bindings.WalletABI))
	if err != nil {
		return nil, err
	}
	return &Wallet{
		address:  address,
		bindings: contractBindings,
		abi:      contractABI,
		ethereum: ethereum,
	}, nil
}

type Wallet struct {
	address  common.Address
	bindings *bindings.Wallet
	abi      abi.ABI
	ethereum *ethclient.Client
}

func DeployWallet(opts *ConstructOpts, eth *ethclient.Client, owner common.Address, transferable bool, ens common.Address, oracleName [32]byte, controllerName [32]byte) (common.Address, *types.Transaction, error) {
	contractABI, err := abi.JSON(strings.NewReader(bindings.WalletABI))
	if err != nil {
		return common.Address{}, nil, err
	}
	data, err := contractABI.Pack("", owner, transferable, ens, oracleName, controllerName)
	if err != nil {
		return common.Address{}, nil, err
	}
	transaction, err := construct(opts, eth, nil, append(common.FromHex(bindings.WalletBin), data...))
	if err != nil {
		return common.Address{}, nil, err
	}
	return crypto.CreateAddress(opts.From, transaction.Nonce()), transaction, err
}

func (w *Wallet) Balance(ctx context.Context, block *big.Int, asset common.Address) (*big.Int, error) {
	data, err := w.abi.Pack("balance", asset)
	if err != nil {
		return nil, err
	}
	rsp, err := w.ethereum.CallContract(ctx, ethereum.CallMsg{
		To:   &w.address,
		Data: data,
	}, block)
	if err != nil {
		return nil, err
	}
	if len(rsp) != 32 {
		return nil, errors.Wrap(ErrFailedContractCall, "balance")
	}
	return new(big.Int).SetBytes(rsp), nil
}

func (w *Wallet) InitializedSpendLimit(ctx context.Context, block *big.Int) (bool, error) {
	data, err := w.abi.Pack("initializedSpendLimit")
	if err != nil {
		return false, err
	}
	rsp, err := w.ethereum.CallContract(ctx, ethereum.CallMsg{
		To:   &w.address,
		Data: data,
	}, block)
	if err != nil {
		return false, err
	}
	if len(rsp) != 32 {
		return false, errors.Wrap(ErrFailedContractCall, "initializedSpendLimit")
	}
	var result bool
	if new(big.Int).SetBytes(rsp).Uint64() == 1 {
		result = true
	}
	return result, nil
}

func (w *Wallet) InitializedWhitelist(ctx context.Context, block *big.Int) (bool, error) {
	data, err := w.abi.Pack("initializedWhitelist")
	if err != nil {
		return false, err
	}
	rsp, err := w.ethereum.CallContract(ctx, ethereum.CallMsg{
		To:   &w.address,
		Data: data,
	}, block)
	if err != nil {
		return false, err
	}
	if len(rsp) != 32 {
		return false, errors.Wrap(ErrFailedContractCall, "initializedWhitelist")
	}
	var result bool
	if new(big.Int).SetBytes(rsp).Uint64() == 1 {
		result = true
	}
	return result, nil
}

func (w *Wallet) InitializedTopUpLimit(ctx context.Context, block *big.Int) (bool, error) {
	data, err := w.abi.Pack("initializedTopUpLimit")
	if err != nil {
		return false, err
	}
	rsp, err := w.ethereum.CallContract(ctx, ethereum.CallMsg{
		To:   &w.address,
		Data: data,
	}, block)
	if err != nil {
		return false, err
	}
	if len(rsp) != 32 {
		return false, errors.Wrap(ErrFailedContractCall, "initializedTopUpLimit")
	}
	var result bool
	if new(big.Int).SetBytes(rsp).Uint64() == 1 {
		result = true
	}
	return result, nil
}

func (w *Wallet) SpendLimit(ctx context.Context, block *big.Int) (*big.Int, error) {
	data, err := w.abi.Pack("spendLimit")
	if err != nil {
		return nil, err
	}
	rsp, err := w.ethereum.CallContract(ctx, ethereum.CallMsg{
		To:   &w.address,
		Data: data,
	}, block)
	if err != nil {
		return nil, err
	}
	if len(rsp) != 32 {
		return nil, errors.Wrap(ErrFailedContractCall, "spendLimit")
	}
	return new(big.Int).SetBytes(rsp), nil
}

func (w *Wallet) SpendAvailable(ctx context.Context, block *big.Int) (*big.Int, error) {
	data, err := w.abi.Pack("spendAvailable")
	if err != nil {
		return nil, err
	}
	rsp, err := w.ethereum.CallContract(ctx, ethereum.CallMsg{
		To:   &w.address,
		Data: data,
	}, block)
	if err != nil {
		return nil, err
	}
	if len(rsp) != 32 {
		return nil, errors.Wrap(ErrFailedContractCall, "spendAvailable")
	}
	return new(big.Int).SetBytes(rsp), nil
}

func (w *Wallet) TopUpLimit(ctx context.Context, block *big.Int) (*big.Int, error) {
	data, err := w.abi.Pack("topUpLimit")
	if err != nil {
		return nil, err
	}
	rsp, err := w.ethereum.CallContract(ctx, ethereum.CallMsg{
		To:   &w.address,
		Data: data,
	}, block)
	if err != nil {
		return nil, err
	}
	if len(rsp) != 32 {
		return nil, errors.Wrap(ErrFailedContractCall, "topUpLimit")
	}
	return new(big.Int).SetBytes(rsp), nil
}

func (w *Wallet) TopUpAvailable(ctx context.Context, block *big.Int) (*big.Int, error) {
	data, err := w.abi.Pack("topUpAvailable")
	if err != nil {
		return nil, err
	}
	rsp, err := w.ethereum.CallContract(ctx, ethereum.CallMsg{
		To:   &w.address,
		Data: data,
	}, block)
	if err != nil {
		return nil, err
	}
	if len(rsp) != 32 {
		return nil, errors.Wrap(ErrFailedContractCall, "topUpAvailable")
	}
	return new(big.Int).SetBytes(rsp), nil
}

func (w *Wallet) IsWhitelisted(ctx context.Context, block *big.Int, address common.Address) (bool, error) {
	data, err := w.abi.Pack("isWhitelisted", address)
	if err != nil {
		return false, err
	}
	rsp, err := w.ethereum.CallContract(ctx, ethereum.CallMsg{
		To:   &w.address,
		Data: data,
	}, block)
	if err != nil {
		return false, err
	}
	if len(rsp) != 32 {
		return false, errors.Wrap(ErrFailedContractCall, "isWhitelisted")
	}
	var result bool
	if new(big.Int).SetBytes(rsp).Uint64() == 1 {
		result = true
	}
	return result, nil
}

func (w *Wallet) Owner(ctx context.Context, block *big.Int) (common.Address, error) {
	data, err := w.abi.Pack("owner")
	if err != nil {
		return common.Address{}, err
	}
	rsp, err := w.ethereum.CallContract(ctx, ethereum.CallMsg{
		To:   &w.address,
		Data: data,
	}, block)
	if err != nil {
		return common.Address{}, err
	}
	if len(rsp) != 32 {
		return common.Address{}, errors.Wrap(ErrFailedContractCall, "owner")
	}
	return common.BytesToAddress(rsp), nil
}

func (w *Wallet) PendingWhitelistAddition(ctx context.Context, block *big.Int) ([]common.Address, error) {
	data, err := w.abi.Pack("pendingWhitelistAddition")
	if err != nil {
		return nil, err
	}
	rsp, err := w.ethereum.CallContract(ctx, ethereum.CallMsg{
		To:   &w.address,
		Data: data,
	}, block)
	if err != nil {
		return nil, err
	}
	if len(rsp) < 64 {
		return nil, errors.Wrap(ErrFailedContractCall, "pendingWhitelistAddition")
	}
	// Get the response array length.
	length := new(big.Int).SetBytes(rsp[32:64])
	// Get the list of whitelisted addresses.
	var res []common.Address
	for i := uint64(0); i < length.Uint64(); i++ {
		res = append(res, common.BytesToAddress(rsp[64+32*i:96+32*i]))
	}
	return res, nil
}

func (w *Wallet) PendingWhitelistRemoval(ctx context.Context, block *big.Int) ([]common.Address, error) {
	data, err := w.abi.Pack("pendingWhitelistRemoval")
	if err != nil {
		return nil, err
	}
	rsp, err := w.ethereum.CallContract(ctx, ethereum.CallMsg{
		To:   &w.address,
		Data: data,
	}, block)
	if err != nil {
		return nil, err
	}
	if len(rsp) < 64 {
		return nil, errors.Wrap(ErrFailedContractCall, "pendingWhitelistRemoval")
	}
	// Get the response array length.
	length := new(big.Int).SetBytes(rsp[32:64])
	// Get the list of whitelisted addresses.
	var res []common.Address
	for i := uint64(0); i < length.Uint64(); i++ {
		res = append(res, common.BytesToAddress(rsp[64+32*i:96+32*i]))
	}
	return res, nil
}

func (w *Wallet) PendingSpendLimit(ctx context.Context, block *big.Int) (*big.Int, error) {
	data, err := w.abi.Pack("pendingSpendLimit")
	if err != nil {
		return nil, err
	}
	rsp, err := w.ethereum.CallContract(ctx, ethereum.CallMsg{
		To:   &w.address,
		Data: data,
	}, block)
	if err != nil {
		return nil, err
	}
	if len(rsp) != 32 {
		return nil, errors.Wrap(ErrFailedContractCall, "pendingSpendLimit")
	}
	return new(big.Int).SetBytes(rsp), nil
}

func (w *Wallet) PendingTopUpLimit(ctx context.Context, block *big.Int) (*big.Int, error) {
	data, err := w.abi.Pack("pendingTopUpLimit")
	if err != nil {
		return nil, err
	}
	rsp, err := w.ethereum.CallContract(ctx, ethereum.CallMsg{
		To:   &w.address,
		Data: data,
	}, block)
	if err != nil {
		return nil, err
	}
	if len(rsp) != 32 {
		return nil, errors.Wrap(ErrFailedContractCall, "pendingTopUpLimit")
	}
	return new(big.Int).SetBytes(rsp), nil
}

func (w *Wallet) InitializeWhitelist(opts *ConstructOpts, addresses []common.Address) (*types.Transaction, error) {
	data, err := w.abi.Pack("initializeWhitelist", addresses)
	if err != nil {
		return nil, err
	}
	return construct(opts, w.ethereum, &w.address, data)
}

func (w *Wallet) SubmitWhitelistAddition(opts *ConstructOpts, addresses []common.Address) (*types.Transaction, error) {
	data, err := w.abi.Pack("submitWhitelistAddition", addresses)
	if err != nil {
		return nil, err
	}
	return construct(opts, w.ethereum, &w.address, data)
}

func (w *Wallet) ConfirmWhitelistAddition(opts *ConstructOpts) (*types.Transaction, error) {
	data, err := w.abi.Pack("confirmWhitelistAddition")
	if err != nil {
		return nil, err
	}
	return construct(opts, w.ethereum, &w.address, data)
}

func (w *Wallet) CancelWhitelistAddition(opts *ConstructOpts) (*types.Transaction, error) {
	data, err := w.abi.Pack("cancelWhitelistAddition")
	if err != nil {
		return nil, err
	}
	return construct(opts, w.ethereum, &w.address, data)
}

func (w *Wallet) SubmitWhitelistRemoval(opts *ConstructOpts, addresses []common.Address) (*types.Transaction, error) {
	data, err := w.abi.Pack("submitWhitelistRemoval", addresses)
	if err != nil {
		return nil, err
	}
	return construct(opts, w.ethereum, &w.address, data)
}

func (w *Wallet) ConfirmWhitelistRemoval(opts *ConstructOpts) (*types.Transaction, error) {
	data, err := w.abi.Pack("confirmWhitelistRemoval")
	if err != nil {
		return nil, err
	}
	return construct(opts, w.ethereum, &w.address, data)
}

func (w *Wallet) CancelWhitelistRemoval(opts *ConstructOpts) (*types.Transaction, error) {
	data, err := w.abi.Pack("cancelWhitelistRemoval")
	if err != nil {
		return nil, err
	}
	return construct(opts, w.ethereum, &w.address, data)
}

func (w *Wallet) InitializeSpendLimit(opts *ConstructOpts, amount *big.Int) (*types.Transaction, error) {
	data, err := w.abi.Pack("initializeSpendLimit", amount)
	if err != nil {
		return nil, err
	}
	return construct(opts, w.ethereum, &w.address, data)
}

func (w *Wallet) SubmitSpendLimit(opts *ConstructOpts, amount *big.Int) (*types.Transaction, error) {
	data, err := w.abi.Pack("submitSpendLimit", amount)
	if err != nil {
		return nil, err
	}
	return construct(opts, w.ethereum, &w.address, data)
}

func (w *Wallet) ConfirmSpendLimit(opts *ConstructOpts) (*types.Transaction, error) {
	data, err := w.abi.Pack("confirmSpendLimit")
	if err != nil {
		return nil, err
	}
	return construct(opts, w.ethereum, &w.address, data)
}

func (w *Wallet) CancelSpendLimit(opts *ConstructOpts) (*types.Transaction, error) {
	data, err := w.abi.Pack("cancelSpendLimit")
	if err != nil {
		return nil, err
	}
	return construct(opts, w.ethereum, &w.address, data)
}

func (w *Wallet) InitializeTopUpLimit(opts *ConstructOpts, amount *big.Int) (*types.Transaction, error) {
	data, err := w.abi.Pack("initializeTopUpLimit", amount)
	if err != nil {
		return nil, err
	}
	return construct(opts, w.ethereum, &w.address, data)
}

func (w *Wallet) SubmitTopUpLimit(opts *ConstructOpts, amount *big.Int) (*types.Transaction, error) {
	data, err := w.abi.Pack("submitTopUpLimit", amount)
	if err != nil {
		return nil, err
	}
	return construct(opts, w.ethereum, &w.address, data)
}

func (w *Wallet) ConfirmTopUpLimit(opts *ConstructOpts) (*types.Transaction, error) {
	data, err := w.abi.Pack("confirmTopUpLimit")
	if err != nil {
		return nil, err
	}
	return construct(opts, w.ethereum, &w.address, data)
}

func (w *Wallet) CancelTopUpLimit(opts *ConstructOpts) (*types.Transaction, error) {
	data, err := w.abi.Pack("cancelTopUpLimit")
	if err != nil {
		return nil, err
	}
	return construct(opts, w.ethereum, &w.address, data)
}

func (w *Wallet) Transfer(opts *ConstructOpts, to common.Address, asset common.Address, amount *big.Int) (*types.Transaction, error) {
	data, err := w.abi.Pack("transfer", to, asset, amount)
	if err != nil {
		return nil, err
	}
	return construct(opts, w.ethereum, &w.address, data)
}

func (w *Wallet) TopUpGas(opts *ConstructOpts, amount *big.Int) (*types.Transaction, error) {
	data, err := w.abi.Pack("topUpGas", amount)
	if err != nil {
		return nil, err
	}
	return construct(opts, w.ethereum, &w.address, data)
}

func (w *Wallet) AddedToWhitelistEvents(ctx context.Context, block *big.Int) ([]*Event, error) {
	// Create a query for whitelist addition events.
	query := ethereum.FilterQuery{
		FromBlock: nil,
		ToBlock:   block,
		Addresses: []common.Address{w.address},
		Topics:    [][]common.Hash{{walletABI.Events["AddedToWhitelist"].Id()}},
	}
	// Get the contract logs.
	logs, err := w.ethereum.FilterLogs(ctx, query)
	if err != nil {
		return nil, err
	}
	// Create a list of logged outgoing events.
	var events []*Event
	for _, v := range logs {
		// Decode event parameters.
		if len(v.Data) < 64 {
			return nil, ErrInvalidEventData
		}
		var data [][]byte
		for i := 0; i < len(v.Data); i += 32 {
			data = append(data, v.Data[i:i+32])
		}
		events = append(events, &Event{
			Address:   v.Address,
			BlockHash: v.BlockHash,
			Data:      data,
			TxHash:    v.TxHash,
		})
	}
	return events, nil
}

func (w *Wallet) RemovedFromWhitelistEvents(ctx context.Context, block *big.Int) ([]*Event, error) {
	// Create a query for whitelist removal events.
	query := ethereum.FilterQuery{
		FromBlock: nil,
		ToBlock:   block,
		Addresses: []common.Address{w.address},
		Topics:    [][]common.Hash{{walletABI.Events["RemovedFromWhitelist"].Id()}},
	}
	// Get the contract logs.
	logs, err := w.ethereum.FilterLogs(ctx, query)
	if err != nil {
		return nil, err
	}
	// Create a list of logged outgoing events.
	var events []*Event
	for _, v := range logs {
		// Decode event parameters.
		if len(v.Data) < 64 {
			return nil, ErrInvalidEventData
		}
		var data [][]byte
		for i := 0; i < len(v.Data); i += 32 {
			data = append(data, v.Data[i:i+32])
		}
		events = append(events, &Event{
			Address:   v.Address,
			BlockHash: v.BlockHash,
			Data:      data,
			TxHash:    v.TxHash,
		})
	}
	return events, nil
}

func (w *Wallet) SetSpendLimitEvents(ctx context.Context, block *big.Int) ([]*Event, error) {
	// Create a query for set daily limit events.
	query := ethereum.FilterQuery{
		FromBlock: nil,
		ToBlock:   block,
		Addresses: []common.Address{w.address},
		Topics:    [][]common.Hash{{walletABI.Events["SetSpendLimit"].Id()}},
	}
	// Get the contract logs.
	logs, err := w.ethereum.FilterLogs(ctx, query)
	if err != nil {
		return nil, err
	}
	// Create a list of logged outgoing events.
	var events []*Event
	for _, v := range logs {
		// Decode event parameters.
		if len(v.Data) != 32 {
			return nil, ErrInvalidEventData
		}
		var data [][]byte
		for i := 0; i < len(v.Data); i += 32 {
			data = append(data, v.Data[i:i+32])
		}
		events = append(events, &Event{
			Address:   v.Address,
			BlockHash: v.BlockHash,
			Data:      data,
			TxHash:    v.TxHash,
		})
	}
	return events, nil
}

func (w *Wallet) SetTopUpLimitEvents(ctx context.Context, block *big.Int) ([]*Event, error) {
	// Create a query for set daily limit events.
	query := ethereum.FilterQuery{
		FromBlock: nil,
		ToBlock:   block,
		Addresses: []common.Address{w.address},
		Topics:    [][]common.Hash{{walletABI.Events["SetTopUpLimit"].Id()}},
	}
	// Get the contract logs.
	logs, err := w.ethereum.FilterLogs(ctx, query)
	if err != nil {
		return nil, err
	}
	// Create a list of logged outgoing events.
	var events []*Event
	for _, v := range logs {
		// Decode event parameters.
		if len(v.Data) != 32 {
			return nil, ErrInvalidEventData
		}
		var data [][]byte
		for i := 0; i < len(v.Data); i += 32 {
			data = append(data, v.Data[i:i+32])
		}
		events = append(events, &Event{
			Address:   v.Address,
			BlockHash: v.BlockHash,
			Data:      data,
			TxHash:    v.TxHash,
		})
	}
	return events, nil
}

func (w *Wallet) ToppedUpGasEvents(ctx context.Context, block *big.Int) ([]*Event, error) {
	// Create a query for gas top up events.
	query := ethereum.FilterQuery{
		FromBlock: nil,
		ToBlock:   block,
		Addresses: []common.Address{w.address},
		Topics:    [][]common.Hash{{walletABI.Events["ToppedUpGas"].Id()}},
	}
	// Get the contract logs.
	logs, err := w.ethereum.FilterLogs(ctx, query)
	if err != nil {
		return nil, err
	}
	// Create a list of logged outgoing events.
	var events []*Event
	for _, v := range logs {
		// Decode event parameters.
		if len(v.Data) != 96 {
			return nil, ErrInvalidEventData
		}
		var data [][]byte
		for i := 0; i < len(v.Data); i += 32 {
			data = append(data, v.Data[i:i+32])
		}
		events = append(events, &Event{
			Address:   v.Address,
			BlockHash: v.BlockHash,
			Data:      data,
			TxHash:    v.TxHash,
		})
	}
	return events, nil
}

func (w *Wallet) TransferredEvents(ctx context.Context, block *big.Int) ([]*Event, error) {
	// Create a query for transfer events.
	query := ethereum.FilterQuery{
		FromBlock: nil,
		ToBlock:   block,
		Addresses: []common.Address{w.address},
		Topics:    [][]common.Hash{{walletABI.Events["Transferred"].Id()}},
	}
	// Get the contract logs.
	logs, err := w.ethereum.FilterLogs(ctx, query)
	if err != nil {
		return nil, err
	}
	// Create a list of logged outgoing events.
	var events []*Event
	for _, v := range logs {
		// Decode event parameters.
		if len(v.Data) != 96 {
			return nil, ErrInvalidEventData
		}
		var data [][]byte
		for i := 0; i < len(v.Data); i += 32 {
			data = append(data, v.Data[i:i+32])
		}
		events = append(events, &Event{
			Address:   v.Address,
			BlockHash: v.BlockHash,
			Data:      data,
			TxHash:    v.TxHash,
		})
	}
	return events, nil
}

func (w *Wallet) ReceivedEvents(ctx context.Context, block *big.Int) ([]*Event, error) {
	// Create a query for deposit events.
	query := ethereum.FilterQuery{
		FromBlock: nil,
		ToBlock:   block,
		Addresses: []common.Address{w.address},
		Topics:    [][]common.Hash{{walletABI.Events["Received"].Id()}},
	}
	// Get the contract logs.
	logs, err := w.ethereum.FilterLogs(ctx, query)
	if err != nil {
		return nil, err
	}
	// Create a list of logged outgoing events.
	var events []*Event
	for _, v := range logs {
		// Decode event parameters.
		if len(v.Data) != 64 {
			return nil, ErrInvalidEventData
		}
		var data [][]byte
		for i := 0; i < len(v.Data); i += 32 {
			data = append(data, v.Data[i:i+32])
		}
		events = append(events, &Event{
			Address:   v.Address,
			BlockHash: v.BlockHash,
			Data:      data,
			TxHash:    v.TxHash,
		})
	}
	return events, nil
}

// construct creates a signed transaction object from the provided parameters.
func construct(opts *ConstructOpts, eth *ethclient.Client, address *common.Address, data []byte) (*types.Transaction, error) {
	value := opts.Value
	if value == nil {
		value = new(big.Int)
	}
	var err error
	gasPrice := opts.GasPrice
	if opts.GasPrice == nil {
		gasPrice, err = eth.SuggestGasPrice(opts.Context)
		if err != nil {
			return nil, err
		}
	}
	gasLimit := opts.GasLimit
	if opts.GasLimit == 0 {
		if address != nil {
			// Check if contract has code.
			code, err := eth.PendingCodeAt(opts.Context, *address)
			if err != nil {
				return nil, err
			} else if len(code) == 0 {
				return nil, ErrNoCode
			}
		}
		gasLimit, err = eth.EstimateGas(opts.Context, ethereum.CallMsg{
			From:  opts.From,
			To:    address,
			Value: value,
			Data:  data,
		})
		if err != nil {
			return nil, err
		}
	}
	if address == nil && opts.Sign == nil {
		return types.NewContractCreation(opts.Nonce, new(big.Int), gasLimit+gasLimit/10, gasPrice, data), nil
	} else if address != nil && opts.Sign == nil {
		return types.NewTransaction(opts.Nonce, *address, new(big.Int), gasLimit+gasLimit/10, gasPrice, data), nil
	} else if address == nil && opts.Sign != nil {
		return opts.Sign(opts.Context, types.NewContractCreation(opts.Nonce, new(big.Int), gasLimit+gasLimit/10, gasPrice, data))
	} else {
		return opts.Sign(opts.Context, types.NewTransaction(opts.Nonce, *address, new(big.Int), gasLimit+gasLimit/10, gasPrice, data))
	}
}
