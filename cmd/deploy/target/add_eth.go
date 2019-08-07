package target

import (
	"context"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func addETH(ctx context.Context, ec *ethclient.Client, to *bind.TransactOpts, resolve Resolve) (common.Address, error) {

	// twa, err := resolve(tokenWhitelistName)
	// if err != nil {
	// 	return common.Address{}, err
	// }

	// tw, err := bindings.NewTokenWhitelist(twa, ec)
	// if err != nil {
	// 	return zeroAddress, err
	// }

	// for _, tok := range append(plainTokens, "TKN") {
	// 	tokenName := fmt.Sprintf("%s.tokencard.eth")

	// 	sca, err := resolve(stablecoinName)
	// 	if err != nil {
	// 		return zeroAddress, err
	// 	}

	// 	tx, err := tw.UpdateTokenRate(to, sca, decimalsToMagnitude(18), big.NewInt(20181215153212))

	// }

	// err = waitForAndConfirmTransaction(ctx, ec, tx)
	// if err != nil {
	// 	return common.Address{}, err
	// }

	return zeroAddress, nil
}
