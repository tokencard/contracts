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
)

func addTokensToTokenWhitelist(ctx context.Context, ec *ethclient.Client, to *bind.TransactOpts, resolve Resolve) (common.Address, error) {

	whitelistAddress, err := resolve(tokenWhitelistName)
	if err != nil {
		return zeroAddress, err
	}

	wl, err := bindings.NewTokenWhitelist(whitelistAddress, ec)
	if err != nil {
		return zeroAddress, err
	}

	notAvailable := []string{}

	for _, tok := range append(plainTokens, "TKN") {
		addr, err := resolve(fmt.Sprintf("%s.tokencard.eth", tok))
		if err != nil {
			return zeroAddress, err
		}

		_, _, _, available, _, _, _, err := wl.GetTokenInfo(nil, addr)
		if err != nil {
			return zeroAddress, err
		}

		if !available {
			notAvailable = append(notAvailable, tok)
		}
	}

	if len(notAvailable) > 0 {
		addresses := make([]common.Address, len(notAvailable))
		loadable := make([]bool, len(notAvailable))
		redeemable := make([]bool, len(notAvailable))
		magnitudes := make([]*big.Int, len(notAvailable))

		for i, tkn := range notAvailable {
			addr, err := resolve(fmt.Sprintf("%s.tokencard.eth", tkn))
			if err != nil {
				return zeroAddress, err
			}
			addresses[i] = addr
			loadable[i] = true
			redeemable[i] = true
			magnitudes[i] = decimalsToMagnitude(18)

		}

		tx, err := wl.AddTokens(
			to,
			addresses,
			stringsToByte32(notAvailable...),
			magnitudes,
			loadable,
			redeemable,
			big.NewInt(20180913153211),
		)

		err = waitForAndConfirmTransaction(ctx, ec, tx)
		if err != nil {
			return common.Address{}, err
		}

		flatRate := new(big.Int)
		_, ok := flatRate.SetString("1000000000000000000", 10)
		if !ok {
			return zeroAddress, errors.New("could not parse flat rate")
		}

		for _, ta := range addresses {
			log.Println("updating token rate for ", ta.Hex())
			tx, err = wl.UpdateTokenRate(to, ta, flatRate, big.NewInt(20180913153212))
			err = waitForAndConfirmTransaction(ctx, ec, tx)
			if err != nil {
				return common.Address{}, err
			}
		}

	}

	return zeroAddress, nil
}

func stringsToByte32(names ...string) [][32]byte {
	r := [][32]byte{}
	for _, n := range names {
		nb := [32]byte{}
		copy(nb[:], []byte(n))
		r = append(r, nb)
	}
	return r
}

func decimalsToMagnitude(decimals int64) *big.Int {
	return new(big.Int).Exp(big.NewInt(10), big.NewInt(decimals), nil)
}
