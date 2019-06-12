package main

import (
	"github.com/pkg/errors"
	"github.com/tokencard/contracts/pkg/bindings"
)

func (d *deployer) deployWallet() error {

	isController, err := d.isController(d.controllerOwner)
	if err != nil {
		return err
	}

	d.log.Infof("Is controller: %t", isController)

	wda, err := d.ens.Addr(walletDeployerName)
	if err != nil {
		return errors.Wrapf(err, "while looking up %s in ENS", walletDeployerName)
	}

	if wda == zeroAddress {
		return errors.Errorf("%s has zero address", walletDeployerName)
	}

	wd, err := bindings.NewWalletDeployer(wda, d.ethClient)
	if err != nil {
		return errors.Wrap(err, "while creating an instance of wallet deployer")
	}

	// walletDeployerABI, err := abi.JSON(strings.NewReader(bindings.WalletDeployerABI))
	// if err != nil {
	// 	return err
	// }

	// data, err := walletDeployerABI.Pack("deployWallet", d.controllerOwner)
	// if err != nil {
	// 	return err
	// }

	//
	tx, err := wd.DeployWallet(d.transactOpts, d.controllerOwner)
	if err != nil {
		return errors.Wrap(err, "while creating an instance of wallet")
	}

	err = d.waitForAndConfirmTransaction(tx.Hash())
	if err != nil {
		return errors.Wrap(err, "while waiting for the transaction")
	}

	contractAddress, err := wd.Deployed(nil, d.controllerOwner)
	if err != nil {
		return errors.Wrap(err, "while getting deployed contract address")
	}

	d.log.Infof("wallet address: %s", contractAddress.Hex())
	return nil

}
