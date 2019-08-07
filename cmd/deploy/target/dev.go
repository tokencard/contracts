package target

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"github.com/tokencard/contracts/v2/pkg/bindings"
	"github.com/tokencard/contracts/v2/pkg/bindings/mocks"
)

func deployDev(ctx context.Context, ec *ethclient.Client, to *bind.TransactOpts, resolve Resolve) (common.Address, error) {

	ca, err := resolve(controllerName)
	if err != nil {
		return zeroAddress, errors.Wrap(err, "whole resolving controller")
	}

	cont, err := bindings.NewController(ca, ec)
	if err != nil {
		return zeroAddress, errors.Wrap(err, "while getting controller instance")
	}

	devControllerAddress, err := resolve(DevControllerAddressName)
	if err != nil {
		return zeroAddress, errors.Wrap(err, "while getting dev controller address")
	}

	isController, err := cont.IsController(nil, devControllerAddress)
	if err != nil {
		return zeroAddress, errors.Wrap(err, "while checking if controller")
	}

	if !isController {
		tx, err := cont.AddController(to, devControllerAddress)
		if err != nil {
			return zeroAddress, errors.Wrap(err, "while adding controller")
		}

		err = waitForAndConfirmTransaction(ctx, ec, tx)
		if err != nil {
			return zeroAddress, errors.Wrap(err, "while waiting for adding controller tx")
		}

	}

	log.Println("controllers set up")

	_, err = resolve(WalletDeployerName)
	if err != nil {
		return zeroAddress, errors.Wrap(err, "while resolving wallet deployer")
	}
	log.Println("got wallet deployer")

	millionToken := &big.Int{}
	_, valid := millionToken.SetString("1000000000000000000000000", 10)
	if !valid {
		return zeroAddress, errors.New("could not parse 1M")
	}

	for _, tok := range plainTokens {
		ensName := fmt.Sprintf("%s.tokencard.eth", tok)
		tokenAddress, err := resolve(ensName)

		if tok == "ETH" {
			continue
		}

		if err != nil {
			return zeroAddress, errors.Wrapf(err, "while resolving %s", ensName)
		}

		token, err := mocks.NewToken(tokenAddress, ec)
		if err != nil {
			return zeroAddress, errors.Wrapf(err, "while getting token instance for %s", ensName)
		}

		tx, err := token.Credit(to, to.From, millionToken)
		if err != nil {
			return zeroAddress, err
		}

		fmt.Println("credited 1M")

		err = waitForAndConfirmTransaction(ctx, ec, tx)
		if err != nil {
			return zeroAddress, err
		}

	}

	_, err = resolve(addTokensToTokenWhitelistName)
	if err != nil {
		return zeroAddress, err
	}

	return zeroAddress, nil
}
