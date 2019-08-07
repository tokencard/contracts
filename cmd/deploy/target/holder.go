package target

import (
	"context"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/contracts/ens"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/tokencard/contracts/v2/pkg/bindings"
)

func deployHolder(ctx context.Context, ec *ethclient.Client, to *bind.TransactOpts, resolve Resolve) (common.Address, error) {
	burnerAddress, err := resolve(tknName)
	if err != nil {
		return common.Address{}, err
	}

	ensAddress, err := resolve(ENSName)
	if err != nil {
		return common.Address{}, err
	}

	_, err = resolve(tokenWhitelistName)
	if err != nil {
		return common.Address{}, err
	}

	_, err = resolve(controllerName)
	if err != nil {
		return common.Address{}, err
	}

	address, tx, _, err := bindings.DeployHolder(to, ec, burnerAddress, ensAddress, ens.EnsNode(tokenWhitelistName), ens.EnsNode(controllerName))
	if err != nil {
		return common.Address{}, err
	}

	err = waitForAndConfirmTransaction(ctx, ec, tx)
	if err != nil {
		return common.Address{}, err
	}

	return address, nil
}
