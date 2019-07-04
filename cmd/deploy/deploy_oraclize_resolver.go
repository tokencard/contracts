package main

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
	"github.com/tokencard/contracts/pkg/bindings/mocks"
)

func (d *deployer) deployOraclizeResolver() error {

	oraclizeAddress, tx, _, err := mocks.DeployOraclize(d.transactOpts, d.ethClient, d.controllerOwner)
	d.waitForAndConfirmTransaction(tx.Hash())
	if err != nil {
		return errors.Wrap(err, "while deploying oraclize")
	}

	return d.deployContract("oraclize-resolver.tokencard.eth", func() (common.Address, *types.Transaction, error) {
		contractAddress, tx, _, err := mocks.DeployOraclizeAddrResolver(d.transactOpts, d.ethClient, oraclizeAddress)
		d.oraclizeResolverAddress = contractAddress
		return contractAddress, tx, err
	})
}
