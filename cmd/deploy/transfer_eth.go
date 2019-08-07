package main

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
)

func (d *deployer) transferETH(to common.Address, amount *big.Int) error {

	d.log.Infof("Transferring %s Wei to %s", amount.String(), to.Hex())

	nonce, err := d.ethClient.PendingNonceAt(d.ctx, d.controllerOwner)
	if err != nil {
		return errors.Wrap(err, "while getting pending nonce")
	}

	unsigned := types.NewTransaction(nonce, to, amount, 30000, big.NewInt(1000000000000), nil)

	tx, err := d.transactOpts.Signer(types.HomesteadSigner{}, d.controllerOwner, unsigned)
	if err != nil {
		return errors.Wrap(err, "while signing transaction")
	}

	err = d.ethClient.SendTransaction(d.ctx, tx)
	if err != nil {
		return errors.Wrap(err, "while sending transaction")
	}

	d.log.Info("Waiting for transaction to be mined")
	err = d.waitForAndConfirmTransaction(tx.Hash())
	if err != nil {
		return errors.Wrap(err, "while waiting for transaction")
	}

	return nil

}
