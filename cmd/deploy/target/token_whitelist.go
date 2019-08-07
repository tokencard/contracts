package target

import (
	"context"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/contracts/ens"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/tokencard/contracts/v2/pkg/bindings"
)

func deployTokenWhitelist(ctx context.Context, ec *ethclient.Client, to *bind.TransactOpts, resolve Resolve) (common.Address, error) {

	ensAddress, err := resolve(ENSName)
	if err != nil {
		return common.Address{}, err
	}

	_, err = resolve(oracleName)
	if err != nil {
		return common.Address{}, err
	}

	_, err = resolve(controllerName)
	if err != nil {
		return common.Address{}, err
	}

	stablecoinAddress, err := resolve(stablecoinName)
	if err != nil {
		return common.Address{}, err
	}

	contractAddress, tx, _, err := bindings.DeployTokenWhitelist(to, ec, ensAddress, ens.EnsNode(oracleName), ens.EnsNode(controllerName), stablecoinAddress)
	if err != nil {
		return common.Address{}, err
	}

	err = waitForAndConfirmTransaction(ctx, ec, tx)
	if err != nil {
		return common.Address{}, err
	}

	return contractAddress, nil
}
