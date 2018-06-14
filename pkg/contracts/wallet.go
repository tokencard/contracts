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

func (w *wallet) Available(ctx context.Context, block *big.Int) (*big.Int, error) {
	data, err := w.abi.Pack("available")
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
		return nil, errors.Wrap(ErrFailedContractCall, "available")
	}
	return res, nil
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

func (w *wallet) IsController(ctx context.Context, block *big.Int, address common.Address) (bool, error) {
	// TODO: How is bool return value represented.
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
		res = new(big.Int).SetBytes(rsp)
	} else {
		return false, errors.Wrap(ErrFailedContractCall, "isController")
	}
	return true, nil
}

func (w *wallet) IsWhitelisted(ctx context.Context, block *big.Int, address common.Address) (bool, error) {
	// TODO: How is bool return value represented.
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
		res = new(big.Int).SetBytes(rsp)
	} else {
		return false, errors.Wrap(ErrFailedContractCall, "isWhitelisted")
	}
	return true, nil
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
	// TODO: PendingAddition, what happens when index out of range (how to break the loop)
	var pending []common.Address
	for i := 0; true; i++ {
		data, err := w.abi.Pack("pendingAddition", i)
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
		if len(rsp) == 32 {
			pending = append(pending, common.BytesToAddress(rsp))
		} else {
			return nil, errors.Wrap(ErrFailedContractCall, "pendingAddition")
		}
	}
	return pending, nil
}

func (w *wallet) PendingRemoval(ctx context.Context, block *big.Int) ([]common.Address, error) {
	// TODO: PendingAddition, what happens when index out of range (how to break the loop)
	var pending []common.Address
	for i := 0; true; i++ {
		data, err := w.abi.Pack("pendingRemoval", i)
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
		if len(rsp) == 32 {
			pending = append(pending, common.BytesToAddress(rsp))
		} else {
			return nil, errors.Wrap(ErrFailedContractCall, "pendingRemoval")
		}
	}
	return pending, nil
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

func (w *wallet) AddController(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return w.bindings.AddController(opts, _account)
}

func (w *wallet) AddToWhitelist(opts *bind.TransactOpts, _addresses []common.Address) (*types.Transaction, error) {
	return w.bindings.AddToWhitelist(opts, _addresses)
}






