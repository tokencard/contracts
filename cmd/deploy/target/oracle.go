package target

import (
	"context"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/contracts/ens"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/tokencard/contracts/v2/pkg/bindings"
)

func deployOracle(ctx context.Context, ec *ethclient.Client, to *bind.TransactOpts, resolve Resolve) (common.Address, error) {
	oraclizeResolverAddress, err := resolve(oraclizeResovlerName)
	if err != nil {
		return common.Address{}, err
	}

	ensAddress, err := resolve(ENSName)
	if err != nil {
		return common.Address{}, err
	}

	// STUPID CYCLIC DEPENDENCY!
	// _, err = resolve(tokenWhitelistName)
	// if err != nil {
	// 	return common.Address{}, err
	// }

	_, err = resolve(controllerName)
	if err != nil {
		return common.Address{}, err
	}

	contractAddress, tx, _, err := bindings.DeployOracle(to, ec, oraclizeResolverAddress, ensAddress, ens.EnsNode(controllerName), ens.EnsNode(tokenWhitelistName))
	if err != nil {
		return common.Address{}, err
	}

	err = waitForAndConfirmTransaction(ctx, ec, tx)
	if err != nil {
		return common.Address{}, err
	}

	return contractAddress, nil
}
