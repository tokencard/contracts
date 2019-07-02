package main

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
)

type createContractDeploymentTransaction func() (common.Address, *types.Transaction, error)

func (d *deployer) deployContract(name string, createTX createContractDeploymentTransaction) error {

	d.log.Infof("Deploying %s", name)

	addr, err := d.ens.Addr(name)

	if err == nil && addr != zeroAddress {
		d.log.Infof("%s found, with address %s", name, addr.Hex())
		return nil
	}

	d.log.Infof("%s not found, deploying ...", name)
	contractAddress, tx, err := createTX()
	if err != nil {
		return errors.Wrapf(err, "while deploying %s", name)
	}

	d.log.Info("Waiting for transaction to be mined")
	err = d.waitForAndConfirmTransaction(tx.Hash())
	if err != nil {
		return errors.Wrap(err, "while waiting for transaction")
	}

	d.log.Infof("Setting ENS for %s to %s", name, contractAddress.Hex())

	err = d.setAddress(name, contractAddress)
	if err != nil {
		return errors.Wrapf(err, "while setting ENS enry for %s", name)
	}

	return nil

}
