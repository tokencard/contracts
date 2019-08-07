package target

import (
	"context"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
)

const waitForMiningTimeout = 2 * 60 * time.Second

func waitForTransactionToBeMined(ctx context.Context, ethClient *ethclient.Client, txHash common.Hash) error {
	ctx, cancel := context.WithTimeout(ctx, waitForMiningTimeout)
	defer cancel()

	for {
		var pending bool
		_, pending, err := ethClient.TransactionByHash(ctx, txHash)
		if err != nil {
			return errors.Wrapf(err, "while getting transaction status of %s", txHash.Hex())
		}

		if !pending {
			break
		}

		time.Sleep(time.Second)
	}

	return nil
}

func ensureTransactionSuccess(ctx context.Context, ethClient *ethclient.Client, txHash common.Hash) error {
	rcpt, err := ethClient.TransactionReceipt(ctx, txHash)
	if err != nil {
		return err
	}

	if rcpt.Status != types.ReceiptStatusSuccessful {
		return errors.Errorf("transaction %s failed", txHash)
	}

	return nil

}

func waitForAndConfirmTransaction(ctx context.Context, ethClient *ethclient.Client, tx *types.Transaction) error {
	err := waitForTransactionToBeMined(ctx, ethClient, tx.Hash())
	if err != nil {
		return err
	}
	return ensureTransactionSuccess(ctx, ethClient, tx.Hash())
}
