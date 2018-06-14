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

	"github.com/pkg/errors"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
)

var ErrFailedContractCall = errors.New("calling smart contract failed")

type Wallet interface {
}

func New(ethereum *ethclient.Client, address common.Address) (Wallet, error) {
	walletBindings, err := bindings.NewWallet(address, ethereum)
	if err != nil {
		return nil, err
	}
	walletABI, err := abi.JSON(strings.NewReader(bindings.WalletABI))
	if err != nil {
		return nil, err
	}
	return &wallet{
		address:  address,
		bindings: walletBindings,
		abi:      walletABI,
		ethereum: ethereum,
	}, nil
}

type wallet struct {
	address  common.Address
	bindings *bindings.Wallet
	abi      abi.ABI
	ethereum *ethclient.Client
}

func (w *wallet) Deploy(auth *bind.TransactOpts, _owner common.Address, _oracle common.Address, _controllers []common.Address) (common.Address, *types.Transaction, *bindings.Wallet, error) {
	return bindings.DeployWallet(auth, w.ethereum,  _owner, _oracle, _controllers)
}

func (w *wallet) Balance(ctx context.Context, block *big.Int, asset common.Address) (*big.Int, error) {
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

func (w *wallet) CurrentDay(ctx context.Context, block *big.Int) (*big.Int, error) {
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

func (w *wallet) DailyLimit(ctx context.Context, block *big.Int) (*big.Int, error) {
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

func (w *wallet) DailyAvailable(ctx context.Context, block *big.Int) (*big.Int, error) {
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

func (w *wallet) GasLimit(ctx context.Context, block *big.Int) (*big.Int, error) {
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

func (w *wallet) GasAvailable(ctx context.Context, block *big.Int) (*big.Int, error) {
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

func (w *wallet) IsController(ctx context.Context, block *big.Int, address common.Address) (bool, error) {
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

func (w *wallet) IsWhitelisted(ctx context.Context, block *big.Int, address common.Address) (bool, error) {
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

func (w *wallet) Oracle(ctx context.Context, block *big.Int) (common.Address, error) {
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

func (w *wallet) Owner(ctx context.Context, block *big.Int) (common.Address, error) {
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

func (w *wallet) PendingAddition(ctx context.Context, block *big.Int) ([]common.Address, error) {
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

func (w *wallet) PendingRemoval(ctx context.Context, block *big.Int) ([]common.Address, error) {
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

func (w *wallet) PendingDailyLimit(ctx context.Context, block *big.Int) (*big.Int, error) {
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

func (w *wallet) PendingGasLimit(ctx context.Context, block *big.Int) (*big.Int, error) {
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

func (w *wallet) AddController(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return w.bindings.AddController(opts, account)
}

func (w *wallet) AddToWhitelist(opts *bind.TransactOpts, addresses []common.Address) (*types.Transaction, error) {
	return w.bindings.AddToWhitelist(opts, addresses)
}








