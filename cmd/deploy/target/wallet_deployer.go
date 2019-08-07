package target

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/tokencard/contracts/v2/pkg/bindings"
)

var oneETH = big.NewInt(1000000000000000000)

func deployWalletDeployer(ctx context.Context, ec *ethclient.Client, to *bind.TransactOpts, resolve Resolve) (common.Address, error) {

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

	_, err = resolve(licenceName)
	if err != nil {
		return common.Address{}, err
	}

	contractAddress, tx, _, err := bindings.DeployWalletDeployer(to, ec, ensAddress, oneETH)
	if err != nil {
		return common.Address{}, err
	}

	err = waitForAndConfirmTransaction(ctx, ec, tx)
	if err != nil {
		return common.Address{}, err
	}

	return contractAddress, nil
}
