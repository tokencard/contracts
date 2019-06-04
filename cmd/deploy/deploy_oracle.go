package main

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/contracts/ens"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/tokencard/contracts/pkg/bindings"
)

const oracleName = "oracle.tokencard.eth"

func (d *deployer) deployOracle() error {
	return d.deployContract(oracleName, func() (common.Address, *types.Transaction, error) {
		contractAddress, tx, _, err := bindings.DeployOracle(d.transactOpts, d.ethClient, d.oraclizeResolverAddress, d.ensAddress, ens.EnsNode(controllerName))
		return contractAddress, tx, err
	})
}
