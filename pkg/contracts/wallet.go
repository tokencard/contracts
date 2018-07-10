package contracts

import (
	"context"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/tokencard/contracts/pkg/bindings"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
	"github.com/tokencard/assets"
)

const (
	depositTopic           = "0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c"
	transferTopic          = "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"
	spendLimitTopic        = "0x21e1049325acc99b4f885709c6ca1a70281b586f585ef03485b62f7ad0a1e253"
	topupLimitTopic        = "0x19ec72a595b8aab321636cc55d51478ac78e93a69b5c4a07ae548eb29e40c0a0"
	whitelistAdditionTopic = "0xc76dd62bd7d0b2212e0d3445c1703a522dd816a749fe499b3bcb0f51b2500434"
	whitelistRemovalTopic  = "0x4b089aff1cdd9a6984aa832d4a013996b3acd3d6244ce3de5e07e6ab050d2b94"
	topupGasTopic          = "0x11bb310b94280c15845698b8ce945817e14456a5d1582e387e6e4a01ef2c6742"
	tokenTopic             = "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"
)

func NewWallet(ethereum *ethclient.Client, assets []*assets.Asset, address common.Address) (*Wallet, error) {
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
		assets:   assets,
	}, nil
}

type Wallet struct {
	address  common.Address
	bindings *bindings.Wallet
	abi      abi.ABI
	ethereum *ethclient.Client
	assets   []*assets.Asset
}

func DeployWallet(opts *ConstructOpts, eth *ethclient.Client, owner common.Address, oracle common.Address, controllers []common.Address) (common.Address, *types.Transaction, error) {
	contractABI, err := abi.JSON(strings.NewReader(bindings.WalletABI))
	if err != nil {
		return common.Address{}, nil, err
	}
	data, err := contractABI.Pack("", owner, oracle, controllers)
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

func (w *Wallet) CurrentDay(ctx context.Context, block *big.Int) (*big.Int, error) {
	data, err := w.abi.Pack("currentDay")
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
		return nil, errors.Wrap(ErrFailedContractCall, "currentDay")
	}
	return new(big.Int).SetBytes(rsp), nil
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

func (w *Wallet) TopupLimit(ctx context.Context, block *big.Int) (*big.Int, error) {
	data, err := w.abi.Pack("topupLimit")
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
		return nil, errors.Wrap(ErrFailedContractCall, "topupLimit")
	}
	return new(big.Int).SetBytes(rsp), nil
}

func (w *Wallet) TopupAvailable(ctx context.Context, block *big.Int) (*big.Int, error) {
	data, err := w.abi.Pack("topupAvailable")
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
		return nil, errors.Wrap(ErrFailedContractCall, "topupAvailable")
	}
	return new(big.Int).SetBytes(rsp), nil
}

func (w *Wallet) IsController(ctx context.Context, block *big.Int, address common.Address) (bool, error) {
	data, err := w.abi.Pack("isController", address)
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
		return false, errors.Wrap(ErrFailedContractCall, "isController")
	}
	var result bool
	if new(big.Int).SetBytes(rsp).Uint64() == 1 {
		result = true
	}
	return result, nil
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

func (w *Wallet) Oracle(ctx context.Context, block *big.Int) (common.Address, error) {
	data, err := w.abi.Pack("oracle")
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
		return common.Address{}, errors.Wrap(ErrFailedContractCall, "oracle")
	}
	return common.BytesToAddress(rsp), nil
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

func (w *Wallet) PendingAddition(ctx context.Context, block *big.Int) ([]common.Address, error) {
	data, err := w.abi.Pack("pendingAddition")
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
		return nil, errors.Wrap(ErrFailedContractCall, "pendingAddition")
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

func (w *Wallet) PendingRemoval(ctx context.Context, block *big.Int) ([]common.Address, error) {
	data, err := w.abi.Pack("pendingRemoval")
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
		return nil, errors.Wrap(ErrFailedContractCall, "pendingRemoval")
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

func (w *Wallet) PendingTopupLimit(ctx context.Context, block *big.Int) (*big.Int, error) {
	data, err := w.abi.Pack("pendingTopupLimit")
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
		return nil, errors.Wrap(ErrFailedContractCall, "pendingTopupLimit")
	}
	return new(big.Int).SetBytes(rsp), nil
}

func (w *Wallet) AddController(opts *ConstructOpts, account common.Address) (*types.Transaction, error) {
	data, err := w.abi.Pack("addController", account)
	if err != nil {
		return nil, err
	}
	return construct(opts, w.ethereum, &w.address, data)
}

func (w *Wallet) RemoveController(opts *ConstructOpts, account common.Address) (*types.Transaction, error) {
	data, err := w.abi.Pack("removeController", account)
	if err != nil {
		return nil, err
	}
	return construct(opts, w.ethereum, &w.address, data)
}

func (w *Wallet) InitializeWhitelist(opts *ConstructOpts, addresses []common.Address) (*types.Transaction, error) {
	data, err := w.abi.Pack("initializeWhitelist", addresses)
	if err != nil {
		return nil, err
	}
	return construct(opts, w.ethereum, &w.address, data)
}

func (w *Wallet) AddToWhitelist(opts *ConstructOpts, addresses []common.Address) (*types.Transaction, error) {
	data, err := w.abi.Pack("addToWhitelist", addresses)
	if err != nil {
		return nil, err
	}
	return construct(opts, w.ethereum, &w.address, data)
}

func (w *Wallet) AddToWhitelistConfirm(opts *ConstructOpts) (*types.Transaction, error) {
	data, err := w.abi.Pack("addToWhitelistConfirm")
	if err != nil {
		return nil, err
	}
	return construct(opts, w.ethereum, &w.address, data)
}

func (w *Wallet) AddToWhitelistCancel(opts *ConstructOpts) (*types.Transaction, error) {
	data, err := w.abi.Pack("addToWhitelistCancel")
	if err != nil {
		return nil, err
	}
	return construct(opts, w.ethereum, &w.address, data)
}

func (w *Wallet) RemoveFromWhitelist(opts *ConstructOpts, addresses []common.Address) (*types.Transaction, error) {
	data, err := w.abi.Pack("removeFromWhitelist", addresses)
	if err != nil {
		return nil, err
	}
	return construct(opts, w.ethereum, &w.address, data)
}

func (w *Wallet) RemoveFromWhitelistConfirm(opts *ConstructOpts) (*types.Transaction, error) {
	data, err := w.abi.Pack("removeFromWhitelistConfirm")
	if err != nil {
		return nil, err
	}
	return construct(opts, w.ethereum, &w.address, data)
}

func (w *Wallet) RemoveFromWhitelistCancel(opts *ConstructOpts) (*types.Transaction, error) {
	data, err := w.abi.Pack("removeFromWhitelistCancel")
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

func (w *Wallet) SetSpendLimit(opts *ConstructOpts, amount *big.Int) (*types.Transaction, error) {
	data, err := w.abi.Pack("setSpendLimit", amount)
	if err != nil {
		return nil, err
	}
	return construct(opts, w.ethereum, &w.address, data)
}

func (w *Wallet) SetSpendLimitConfirm(opts *ConstructOpts) (*types.Transaction, error) {
	data, err := w.abi.Pack("setSpendLimitConfirm")
	if err != nil {
		return nil, err
	}
	return construct(opts, w.ethereum, &w.address, data)
}

func (w *Wallet) SetSpendLimitCancel(opts *ConstructOpts) (*types.Transaction, error) {
	data, err := w.abi.Pack("setSpendLimitCancel")
	if err != nil {
		return nil, err
	}
	return construct(opts, w.ethereum, &w.address, data)
}

func (w *Wallet) InitializeTopupLimit(opts *ConstructOpts, amount *big.Int) (*types.Transaction, error) {
	data, err := w.abi.Pack("initializeTopupLimit", amount)
	if err != nil {
		return nil, err
	}
	return construct(opts, w.ethereum, &w.address, data)
}

func (w *Wallet) SetTopupLimit(opts *ConstructOpts, amount *big.Int) (*types.Transaction, error) {
	data, err := w.abi.Pack("setTopupLimit", amount)
	if err != nil {
		return nil, err
	}
	return construct(opts, w.ethereum, &w.address, data)
}

func (w *Wallet) SetTopupLimitConfirm(opts *ConstructOpts) (*types.Transaction, error) {
	data, err := w.abi.Pack("setTopupLimitConfirm")
	if err != nil {
		return nil, err
	}
	return construct(opts, w.ethereum, &w.address, data)
}

func (w *Wallet) SetTopupLimitCancel(opts *ConstructOpts) (*types.Transaction, error) {
	data, err := w.abi.Pack("setTopupLimitCancel")
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

func (w *Wallet) TopupGas(opts *ConstructOpts, amount *big.Int) (*types.Transaction, error) {
	data, err := w.abi.Pack("topupGas", amount)
	if err != nil {
		return nil, err
	}
	return construct(opts, w.ethereum, &w.address, data)
}

func (w *Wallet) WhitelistAdditionEvents(ctx context.Context, block *big.Int) ([]*Event, error) {
	// Create a query for whitelist addition events.
	query := ethereum.FilterQuery{
		FromBlock: nil,
		ToBlock:   block,
		Addresses: []common.Address{w.address},
		Topics:    [][]common.Hash{{common.HexToHash(whitelistAdditionTopic)}},
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

func (w *Wallet) WhitelistRemovalEvents(ctx context.Context, block *big.Int) ([]*Event, error) {
	// Create a query for whitelist removal events.
	query := ethereum.FilterQuery{
		FromBlock: nil,
		ToBlock:   block,
		Addresses: []common.Address{w.address},
		Topics:    [][]common.Hash{{common.HexToHash(whitelistRemovalTopic)}},
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
		Topics:    [][]common.Hash{{common.HexToHash(spendLimitTopic)}},
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

func (w *Wallet) SetTopupLimitEvents(ctx context.Context, block *big.Int) ([]*Event, error) {
	// Create a query for set daily limit events.
	query := ethereum.FilterQuery{
		FromBlock: nil,
		ToBlock:   block,
		Addresses: []common.Address{w.address},
		Topics:    [][]common.Hash{{common.HexToHash(topupLimitTopic)}},
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

func (w *Wallet) TopupGasEvents(ctx context.Context, block *big.Int) ([]*Event, error) {
	// Create a query for gas top up events.
	query := ethereum.FilterQuery{
		FromBlock: nil,
		ToBlock:   block,
		Addresses: []common.Address{w.address},
		Topics:    [][]common.Hash{{common.HexToHash(topupGasTopic)}},
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

func (w *Wallet) TransferEvents(ctx context.Context, block *big.Int) ([]*Event, error) {
	// Create a query for transfer events.
	query := ethereum.FilterQuery{
		FromBlock: nil,
		ToBlock:   block,
		Addresses: []common.Address{w.address},
		Topics:    [][]common.Hash{{common.HexToHash(transferTopic)}},
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

func (w *Wallet) DepositEvents(ctx context.Context, block *big.Int) ([]*Event, error) {
	// Create a query for deposit events.
	query := ethereum.FilterQuery{
		FromBlock: nil,
		ToBlock:   block,
		Addresses: []common.Address{w.address},
		Topics:    [][]common.Hash{{common.HexToHash(depositTopic)}},
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

func (w *Wallet) TokenDepositEvents(ctx context.Context, block *big.Int) ([]*Event, error) {
	var events []*Event
	// Scan supported tokens for incoming events.
	for _, asset := range w.assets {
		if asset.Id == assets.ETHER {
			continue
		}
		// Create a log filter query.
		query := ethereum.FilterQuery{
			FromBlock: nil,
			ToBlock:   block,
			Addresses: []common.Address{common.HexToAddress(asset.Id)},
			Topics:    [][]common.Hash{{common.HexToHash(tokenTopic)}, nil, {w.address.Hash()}},
		}
		// Get the contract logs.
		logs, err := w.ethereum.FilterLogs(ctx, query)
		if err != nil {
			return nil, err
		}
		// Create a list of incoming asset transfer events.
		for _, v := range logs {
			// Decode event parameters.
			if len(v.Data) < 32 {
				return nil, ErrInvalidEventData
			}
			var data [][]byte
			data = append(data, v.Topics[len(v.Topics)-1].Bytes())
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
	}
	return events, nil
}
