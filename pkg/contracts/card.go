package contracts

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/tokencard/contracts/pkg/bindings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
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





