package main

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
)

type createContractDeploymentTransaction func() (common.Address, *types.Transaction, error)

func (d *deployer) deployContract(name string, createTX createContractDeploymentTransaction) error {
	contractAddr, err := d.ens.Addr(name)
	if err != nil {
		return errors.Wrap(err, "while looking up "+name)
	}
	if contractAddr != zeroAddress {
		d.log.Info("Found already deployed %s on %s", name, contractAddr.String())
		return nil
	}

	d.log.Infof("%s not found, deploying ...", name)
	contractAddress, tx, err := createTX()
	if err != nil {
		return errors.Wrapf(err, "while deploying %s", name)
	}

	d.log.Info("Waiting for transaction to be mined")
	err = d.waitForTransactionToBeMined(tx.Hash())
	if err != nil {
		return errors.Wrapf(err, "while waiting for deployment transaction of %s", name)
	}

	err = d.ensureTransactionSuccess(tx.Hash())
	if err != nil {
		return errors.Wrapf(err, "while ensuring transaction success for deployment of %s", name)
	}

	d.log.Infof("Setting ENS for %s", name)

	tx, err = d.ens.SetAddr(name, contractAddress)
	if err != nil {
		return errors.Wrap(err, "while deploying %s")
	}

	d.log.Info("Waiting for ENS transaction to be mined")
	err = d.waitForTransactionToBeMined(tx.Hash())
	if err != nil {
		return errors.Wrap(err, "while waiting for setting ENS transaction to be deployed")
	}

	err = d.ensureTransactionSuccess(tx.Hash())
	if err != nil {
		return errors.Wrapf(err, "while ensuring transaction success for setting ENS of %s", name)
	}

	return nil

}
