package wrappers

import (
	"context"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/tokencard/contracts/pkg/bindings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
)

// NewController creates a new controller contract, if provided address is `nil` the contract will have no bindings.
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

func DeployController(opts *ConstructOpts, eth *ethclient.Client) (common.Address, *types.Transaction, error) {
	transaction, err := construct(opts, eth, nil, common.FromHex(bindings.ControllerBin))
	if err != nil {
		return common.Address{}, nil, err
	}
	return crypto.CreateAddress(opts.From, transaction.Nonce()), transaction, err
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

func (c *Controller) Process(opts *ConstructOpts, sequenceNumber *big.Int, operations []byte, cards []common.Address, tokens []common.Address, values []*big.Int) (*types.Transaction, error) {
	data, err := c.abi.Pack("process", sequenceNumber, operations, cards, tokens, values)
	if err != nil {
		return nil, err
	}
	return construct(opts, c.ethereum, &c.address, data)
}
