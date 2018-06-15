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

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
	"github.com/tokencard/assets"
)

var (
	ErrFailedContractCall = errors.New("calling smart contract failed")
	ErrInvalidEventData   = errors.New("event data could not be parsed")
)

const (
	depositTopic           = "0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c"
	transferTopic          = "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"
	dailyLimitTopic        = "0x4c6fb4d469797cfdfb4ac382c2fe84a494711f5b9676334ff9fac5f7601aef70"
	whitelistAdditionTopic = "0xc76dd62bd7d0b2212e0d3445c1703a522dd816a749fe499b3bcb0f51b2500434"
	whitelistRemovalTopic  = "0x4b089aff1cdd9a6984aa832d4a013996b3acd3d6244ce3de5e07e6ab050d2b94"
	topUpGasTopic          = "0x2235de9f3363e464311d990c51aeef966703c87d1c77e80737831d6944d87c86"
	ERC20Topic             = "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"
)

func New(ethereum *ethclient.Client, assets []*assets.Asset, address common.Address) (*Wallet, error) {
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

type Event struct {
	Address   common.Address
	TxHash    common.Hash
	BlockHash common.Hash
	Data      [][]byte
}

func DeployWallet(opts *bind.TransactOpts, eth *ethclient.Client, owner common.Address, oracle common.Address, controllers []common.Address) (common.Address, *types.Transaction, *bindings.Wallet, error) {
	return bindings.DeployWallet(opts, eth, owner, oracle, controllers)
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
	var res *big.Int
	if len(rsp) == 32 {
		res = new(big.Int).SetBytes(rsp)
	} else {
		return nil, errors.Wrap(ErrFailedContractCall, "balance")
	}
	return res, nil
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
	var res *big.Int
	if len(rsp) == 32 {
		res = new(big.Int).SetBytes(rsp)
	} else {
		return nil, errors.Wrap(ErrFailedContractCall, "currentDay")
	}
	return res, nil
}

func (w *Wallet) DailyLimit(ctx context.Context, block *big.Int) (*big.Int, error) {
	data, err := w.abi.Pack("dailyLimit")
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
	var res *big.Int
	if len(rsp) == 32 {
		res = new(big.Int).SetBytes(rsp)
	} else {
		return nil, errors.Wrap(ErrFailedContractCall, "dailyLimit")
	}
	return res, nil
}

func (w *Wallet) DailyAvailable(ctx context.Context, block *big.Int) (*big.Int, error) {
	data, err := w.abi.Pack("dailyAvailable")
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
	var res *big.Int
	if len(rsp) == 32 {
		res = new(big.Int).SetBytes(rsp)
	} else {
		return nil, errors.Wrap(ErrFailedContractCall, "dailyAvailable")
	}
	return res, nil
}

func (w *Wallet) GasLimit(ctx context.Context, block *big.Int) (*big.Int, error) {
	data, err := w.abi.Pack("gasLimit")
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
	var res *big.Int
	if len(rsp) == 32 {
		res = new(big.Int).SetBytes(rsp)
	} else {
		return nil, errors.Wrap(ErrFailedContractCall, "gasLimit")
	}
	return res, nil
}

func (w *Wallet) GasAvailable(ctx context.Context, block *big.Int) (*big.Int, error) {
	data, err := w.abi.Pack("gasAvailable")
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
	var res *big.Int
	if len(rsp) == 32 {
		res = new(big.Int).SetBytes(rsp)
	} else {
		return nil, errors.Wrap(ErrFailedContractCall, "gasAvailable")
	}
	return res, nil
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
	var res bool
	if len(rsp) == 32 {
		switch new(big.Int).SetBytes(rsp).Uint64() {
		case 0:
			res = false
		case 1:
			res = true
		}
	} else {
		return false, errors.Wrap(ErrFailedContractCall, "isController")
	}
	return res, nil
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
	var res bool
	if len(rsp) == 32 {
		switch new(big.Int).SetBytes(rsp).Uint64() {
		case 0:
			res = false
		case 1:
			res = true
		}
	} else {
		return false, errors.Wrap(ErrFailedContractCall, "isWhitelisted")
	}
	return res, nil
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
	var res common.Address
	if len(rsp) == 32 {
		res = common.BytesToAddress(rsp)
	} else {
		return common.Address{}, errors.Wrap(ErrFailedContractCall, "oracle")
	}
	return res, nil
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
	var res common.Address
	if len(rsp) == 32 {
		res = common.BytesToAddress(rsp)
	} else {
		return common.Address{}, errors.Wrap(ErrFailedContractCall, "owner")
	}
	return res, nil
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

func (w *Wallet) PendingDailyLimit(ctx context.Context, block *big.Int) (*big.Int, error) {
	data, err := w.abi.Pack("pendingDailyLimit")
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
	var res *big.Int
	if len(rsp) == 32 {
		res = new(big.Int).SetBytes(rsp)
	} else {
		return nil, errors.Wrap(ErrFailedContractCall, "pendingDailyLimit")
	}
	return res, nil
}

func (w *Wallet) PendingGasLimit(ctx context.Context, block *big.Int) (*big.Int, error) {
	data, err := w.abi.Pack("pendingGasLimit")
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
	var res *big.Int
	if len(rsp) == 32 {
		res = new(big.Int).SetBytes(rsp)
	} else {
		return nil, errors.Wrap(ErrFailedContractCall, "pendingGasLimit")
	}
	return res, nil
}

func (w *Wallet) AddController(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return w.bindings.AddController(opts, account)
}

func (w *Wallet) RemoveController(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return w.bindings.RemoveController(opts, account)
}

func (w *Wallet) InitializeWhitelist(opts *bind.TransactOpts, addresses []common.Address) (*types.Transaction, error) {
	return w.bindings.InitializeWhitelist(opts, addresses)
}

func (w *Wallet) AddToWhitelist(opts *bind.TransactOpts, addresses []common.Address) (*types.Transaction, error) {
	return w.bindings.AddToWhitelist(opts, addresses)
}

func (w *Wallet) AddToWhitelistConfirm(opts *bind.TransactOpts) (*types.Transaction, error) {
	return w.bindings.AddToWhitelistConfirm(opts)
}

func (w *Wallet) AddToWhitelistCancel(opts *bind.TransactOpts) (*types.Transaction, error) {
	return w.bindings.AddToWhitelistCancel(opts)
}

func (w *Wallet) RemoveFromWhitelist(opts *bind.TransactOpts, addresses []common.Address) (*types.Transaction, error) {
	return w.bindings.RemoveFromWhitelist(opts, addresses)
}

func (w *Wallet) RemoveFromWhitelistConfirm(opts *bind.TransactOpts) (*types.Transaction, error) {
	return w.bindings.RemoveFromWhitelistConfirm(opts)
}

func (w *Wallet) RemoveFromWhitelistCancel(opts *bind.TransactOpts) (*types.Transaction, error) {
	return w.bindings.RemoveFromWhitelistCancel(opts)
}

func (w *Wallet) InitializeDailyLimit(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return w.bindings.InitializeDailyLimit(opts, amount)
}

func (w *Wallet) SetDailyLimit(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return w.bindings.SetDailyLimit(opts, amount)
}

func (w *Wallet) SetDailyLimitConfirm(opts *bind.TransactOpts) (*types.Transaction, error) {
	return w.bindings.SetDailyLimitConfirm(opts)
}

func (w *Wallet) SetDailyLimitCancel(opts *bind.TransactOpts) (*types.Transaction, error) {
	return w.bindings.SetDailyLimitCancel(opts)
}

func (w *Wallet) InitializeGasLimit(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return w.bindings.InitializeGasLimit(opts, amount)
}

func (w *Wallet) SetGasLimit(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return w.bindings.SetGasLimit(opts, amount)
}

func (w *Wallet) SetGasLimitConfirm(opts *bind.TransactOpts) (*types.Transaction, error) {
	return w.bindings.SetGasLimitConfirm(opts)
}

func (w *Wallet) SetGasLimitCancel(opts *bind.TransactOpts) (*types.Transaction, error) {
	return w.bindings.SetGasLimitCancel(opts)
}

func (w *Wallet) Transfer(opts *bind.TransactOpts, to common.Address, asset common.Address, amount *big.Int) (*types.Transaction, error) {
	return w.bindings.Transfer(opts, to, asset, amount)
}

func (w *Wallet) TopUpGas(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return w.bindings.TopUpGas(opts, amount)
}

func (w *Wallet) WhitelistAdditionEvents(ctx context.Context, block *big.Int) ([]*Event, error) {
	// Create a query for top up gas events.
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
	// Create a query for top up gas events.
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

func (w *Wallet) SetDailyLimitEvents(ctx context.Context, block *big.Int) ([]*Event, error) {
	// Create a query for set daily limit events.
	query := ethereum.FilterQuery{
		FromBlock: nil,
		ToBlock:   block,
		Addresses: []common.Address{w.address},
		Topics:    [][]common.Hash{{common.HexToHash(dailyLimitTopic)}},
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

func (w *Wallet) TopUpGasEvents(ctx context.Context, block *big.Int) ([]*Event, error) {
	// Create a query for top up gas events.
	query := ethereum.FilterQuery{
		FromBlock: nil,
		ToBlock:   block,
		Addresses: []common.Address{w.address},
		Topics:    [][]common.Hash{{common.HexToHash(topUpGasTopic)}},
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
	// Create a query for top up gas events.
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
	// Create a query for top up gas events.
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

func (w *Wallet) ERC20DepositEvents(ctx context.Context, block *big.Int) ([]*Event, error) {
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
			Topics:    [][]common.Hash{{common.HexToHash(ERC20Topic)}, nil, {w.address.Hash()}},
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
