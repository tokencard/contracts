package target

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/contracts/ens"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/tokencard/contracts/v2/pkg/bindings"
)

func deployLicence(ctx context.Context, ec *ethclient.Client, to *bind.TransactOpts, resolve Resolve) (common.Address, error) {
	// floatAddress, err := resolve(floatName)
	// if err != nil {
	// 	return common.Address{}, err
	// }

	holderAddress, err := resolve(holderName)
	if err != nil {
		return common.Address{}, err
	}

	tknAddress, err := resolve(tknName)
	if err != nil {
		return common.Address{}, err
	}

	ensAddress, err := resolve(ENSName)
	if err != nil {
		return common.Address{}, err
	}

	_, err = resolve(controllerName)
	if err != nil {
		return common.Address{}, err
	}

	floatAddress := to.From

	licenceAddress, tx, _, err := bindings.DeployLicence(to, ec, big.NewInt(10), floatAddress, holderAddress, tknAddress, ensAddress, ens.EnsNode(controllerName))
	if err != nil {
		return common.Address{}, err
	}

	err = waitForAndConfirmTransaction(ctx, ec, tx)
	if err != nil {
		return common.Address{}, err
	}

	return licenceAddress, nil
}
