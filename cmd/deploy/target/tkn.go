package target

import (
	"context"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/tokencard/contracts/v2/pkg/bindings/mocks"
)

func deployTkn(ctx context.Context, ec *ethclient.Client, to *bind.TransactOpts, resolve Resolve) (common.Address, error) {

	contractAddress, tx, _, err := mocks.DeployBurnerToken(to, ec)
	if err != nil {
		return common.Address{}, err
	}

	err = waitForAndConfirmTransaction(ctx, ec, tx)
	if err != nil {
		return common.Address{}, err
	}

	return contractAddress, nil
}
