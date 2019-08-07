package target

import (
	"context"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/tokencard/contracts/v2/pkg/bindings/mocks"
)

func deployOraclizeConnector(ctx context.Context, ec *ethclient.Client, to *bind.TransactOpts, resolve Resolve) (common.Address, error) {

	oraclizeConnectorAddress, tx, _, err := mocks.DeployOraclizeConnector(to, ec, to.From)
	if err != nil {
		return common.Address{}, err
	}

	err = waitForAndConfirmTransaction(ctx, ec, tx)
	if err != nil {
		return common.Address{}, err
	}

	return oraclizeConnectorAddress, nil
}
