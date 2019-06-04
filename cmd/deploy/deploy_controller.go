package main

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/tokencard/contracts/pkg/bindings/internals"
)

const controllerName = "controller.tokencard.eth"

func (d *deployer) deployController() error {
	return d.deployContract(controllerName, func() (common.Address, *types.Transaction, error) {
		contractAddress, tx, _, err := internals.DeployController(d.transactOpts, d.ethClient, d.controllerOwner)
		return contractAddress, tx, err
	})
}
