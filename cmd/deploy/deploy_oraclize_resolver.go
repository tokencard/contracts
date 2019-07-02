package main

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
	"github.com/tokencard/contracts/pkg/bindings/mocks"
)

func (d *deployer) deployOraclizeResolver() error {

	oraclizeConnectorAddress, _, _, err := mocks.DeployOraclize(d.transactOpts, d.ethClient, d.controllerOwner)
	if err != nil {
		return errors.Wrap(err, "while deploying oraclize")
	}

	d.log.Infof("Oraclize Connector is at %s", oraclizeConnectorAddress.Hex())

	return d.deployContract(oracleName, func() (common.Address, *types.Transaction, error) {
		contractAddress, tx, _, err := mocks.DeployOraclizeAddrResolver(d.transactOpts, d.ethClient, oraclizeConnectorAddress)
		d.oraclizeResolverAddress = contractAddress
		return contractAddress, tx, err
	})
}
