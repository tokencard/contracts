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





