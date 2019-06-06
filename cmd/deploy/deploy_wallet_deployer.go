package main

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/contracts/ens"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/tokencard/contracts/pkg/bindings"
)

const walletDeployerName = "wallet-deployer.tokencard.eth"

var oneETH = big.NewInt(1000000000000000000)

func (d *deployer) deployWalletDeployer() error {
	return d.deployContract(walletDeployerName, func() (common.Address, *types.Transaction, error) {
		contractAddress, tx, _, err := bindings.DeployWalletDeployer(d.transactOpts, d.ethClient, d.ensAddress, ens.EnsNode(oracleName), ens.EnsNode(controllerName), oneETH)
		return contractAddress, tx, err
	})
}
