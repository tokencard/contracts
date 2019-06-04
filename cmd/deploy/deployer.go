package main

import (
	"context"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/contracts/ens"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type deployer struct {
	transactOpts            *bind.TransactOpts
	ethClient               *ethclient.Client
	ens                     *ens.ENS
	ensAddress              common.Address
	log                     logrus.FieldLogger
	controllerOwner         common.Address
	ctx                     context.Context
	oraclizeResolverAddress common.Address
}

var zeroAddress = common.HexToAddress("0x0")

const waitForMiningTimeout = 2 * 60 * time.Second

func (d *deployer) waitForTransactionToBeMined(txHash common.Hash) error {
	ctx, cancel := context.WithTimeout(d.ctx, waitForMiningTimeout)
	defer cancel()
	for {
		var pending bool
		_, pending, err := d.ethClient.TransactionByHash(ctx, txHash)
		if err != nil {
			return errors.Wrap(err, "while getting transaction status")
		}

		if !pending {
			continue
		}

		time.Sleep(time.Second)
	}
}

func (d *deployer) ensureTransactionSuccess(txHash common.Hash) error {
	rcpt, err := d.ethClient.TransactionReceipt(d.ctx, txHash)
	if err != nil {
		return err
	}

	if rcpt.Status != types.ReceiptStatusSuccessful {
		return errors.New("controller deployment transaction failed")
	}

	return nil

}
