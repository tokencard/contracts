package contracts

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/tokencard/contracts/pkg/bindings"

	"context"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
)

func NewCard(ethereum *ethclient.Client, address common.Address) (*Card, error) {
	contractBindings, err := bindings.NewCard(address, ethereum)
	if err != nil {
		return nil, err
	}
	contractABI, err := abi.JSON(strings.NewReader(bindings.CardABI))
	if err != nil {
		return nil, err
	}
	return &Card{
		address:  address,
		bindings: contractBindings,
		abi:      contractABI,
		ethereum: ethereum,
	}, nil
}

type Card struct {
	address  common.Address
	bindings *bindings.Card
	abi      abi.ABI
	ethereum *ethclient.Client
}

func DeployCard(opts *bind.TransactOpts, eth *ethclient.Client, controller common.Address, owner common.Address) (common.Address, *types.Transaction, *bindings.Card, error) {
	return bindings.DeployCard(opts, eth, controller, owner)
}

func (c *Card) Owner(ctx context.Context, block *big.Int) (common.Address, error) {
	data, err := c.abi.Pack("owner")
	if err != nil {
		return common.Address{}, err
	}
	rsp, err := c.ethereum.CallContract(ctx, ethereum.CallMsg{
		To:   &c.address,
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

func (c *Card) Controller(ctx context.Context, block *big.Int) (common.Address, error) {
	data, err := c.abi.Pack("controller")
	if err != nil {
		return common.Address{}, err
	}
	rsp, err := c.ethereum.CallContract(ctx, ethereum.CallMsg{
		To:   &c.address,
		Data: data,
	}, block)
	if err != nil {
		return common.Address{}, err
	}
	if len(rsp) != 32 {
		return common.Address{}, errors.Wrap(ErrFailedContractCall, "controller")
	}
	return common.BytesToAddress(rsp), nil
}

func (c *Card) UnlockAt(ctx context.Context, block *big.Int) (*big.Int, error) {
	data, err := c.abi.Pack("unlockAt")
	if err != nil {
		return nil, err
	}
	rsp, err := c.ethereum.CallContract(ctx, ethereum.CallMsg{
		To:   &c.address,
		Data: data,
	}, block)
	if err != nil {
		return nil, err
	}
	if len(rsp) != 32 {
		return nil, errors.Wrap(ErrFailedContractCall, "unlockAt")
	}
	return new(big.Int).SetBytes(rsp), nil
}

func (c *Card) Balance(ctx context.Context, block *big.Int, asset common.Address) (*big.Int, error) {
	data, err := c.abi.Pack("balance", asset)
	if err != nil {
		return nil, err
	}
	rsp, err := c.ethereum.CallContract(ctx, ethereum.CallMsg{
		To:   &c.address,
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

func (c *Card) Overdraft(ctx context.Context, block *big.Int, asset common.Address) (*big.Int, error) {
	data, err := c.abi.Pack("overdraft", asset)
	if err != nil {
		return nil, err
	}
	rsp, err := c.ethereum.CallContract(ctx, ethereum.CallMsg{
		To:   &c.address,
		Data: data,
	}, block)
	if err != nil {
		return nil, err
	}
	if len(rsp) != 32 {
		return nil, errors.Wrap(ErrFailedContractCall, "overdraft")
	}
	return new(big.Int).SetBytes(rsp), nil
}

func (c *Card) Available(ctx context.Context, block *big.Int, asset common.Address) (*big.Int, error) {
	data, err := c.abi.Pack("available", asset)
	if err != nil {
		return nil, err
	}
	rsp, err := c.ethereum.CallContract(ctx, ethereum.CallMsg{
		To:   &c.address,
		Data: data,
	}, block)
	if err != nil {
		return nil, err
	}
	if len(rsp) != 32 {
		return nil, errors.Wrap(ErrFailedContractCall, "available")
	}
	return new(big.Int).SetBytes(rsp), nil
}
