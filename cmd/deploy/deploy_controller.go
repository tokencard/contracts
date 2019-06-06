package main

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
	"github.com/tokencard/contracts/pkg/bindings/internals"
)

const controllerName = "controller.tokencard.eth"

func (d *deployer) deployController() error {
	return d.deployContract(controllerName, func() (common.Address, *types.Transaction, error) {
		contractAddress, tx, _, err := internals.DeployController(d.transactOpts, d.ethClient, d.controllerOwner)
		return contractAddress, tx, err
	})
}

func (d *deployer) isController(addr common.Address) (bool, error) {
	cona, err := d.ens.Addr(controllerName)
	if err != nil {
		return false, errors.Wrapf(err, "while looking up %s in ENS", controllerName)
	}

	if cona == zeroAddress {
		return false, errors.Errorf("%s has zero address", controllerName)
	}

	con, err := internals.NewController(cona, d.ethClient)

	if err != nil {
		return false, errors.Wrap(err, "while creating controller instance")
	}

	return con.IsController(nil, addr)
}
