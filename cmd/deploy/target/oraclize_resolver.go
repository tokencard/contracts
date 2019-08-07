package target

import (
	"context"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/tokencard/contracts/v2/pkg/bindings/mocks"
)

func deployOraclizeResolver(ctx context.Context, ec *ethclient.Client, to *bind.TransactOpts, resolve Resolve) (common.Address, error) {
	oraclizeConnectorAddress, err := resolve(oraclizeConnectorName)
	if err != nil {
		return common.Address{}, err
	}

	oraclizeResolverAddress, tx, _, err := mocks.DeployOraclizeAddrResolver(to, ec, oraclizeConnectorAddress)
	if err != nil {
		return common.Address{}, err
	}

	err = waitForAndConfirmTransaction(ctx, ec, tx)
	if err != nil {
		return common.Address{}, err
	}

	return oraclizeResolverAddress, nil
}
