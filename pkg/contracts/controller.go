package contracts

import (
	"context"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/tokencard/contracts/pkg/bindings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
)

const (
	ControllerABI = bindings.ControllerABI
	ControllerBin = bindings.ControllerBin
)

func NewController(ethereum *ethclient.Client, address common.Address) (*Controller, error) {
	contractBindings, err := bindings.NewController(address, ethereum)
	if err != nil {
		return nil, err
	}
	contractABI, err := abi.JSON(strings.NewReader(bindings.ControllerABI))
	if err != nil {
		return nil, err
	}
	return &Controller{
		address:  address,
		bindings: contractBindings,
		abi:      contractABI,
		ethereum: ethereum,
	}, nil
}

type Controller struct {
	address  common.Address
	bindings *bindings.Controller
	abi      abi.ABI
	ethereum *ethclient.Client
}

func DeployController(opts *bind.TransactOpts, eth *ethclient.Client) (common.Address, *types.Transaction, *bindings.Controller, error) {
	return bindings.DeployController(opts, eth)
}

func (c *Controller) Next(ctx context.Context, block *big.Int) (*big.Int, error) {
	data, err := c.abi.Pack("next")
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
		return nil, errors.Wrap(ErrFailedContractCall, "next")
	}
	return new(big.Int).SetBytes(rsp), nil
}

func (c *Controller) Process(opts *bind.TransactOpts, sequenceNumber *big.Int, operations []byte, cards []common.Address, tokens []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return c.bindings.Process(opts, sequenceNumber, operations, cards, tokens, amounts)
}
