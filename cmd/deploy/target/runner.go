package target

import (
	"context"
	"log"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/contracts/ens"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type runner struct {
	ec     *ethclient.Client
	en     *ens.ENS
	params map[string]common.Address
}

var zeroAddress = common.Address{}

func (r runner) resolve(name string) (common.Address, error) {

	if strings.HasSuffix(name, ".eth") {
		addr, err := r.en.Addr(name)
		if err != nil || addr == zeroAddress {
			// TODO run dep
		}

		return addr, nil
	}

	return zeroAddress, errors.Errorf("Unknown task %s", name)
}

var forceDeployNames = []string{
	WalletDeployerName,
	tokenWhitelistName,
	oracleName,
	controllerName,
}

func forceDeploy(name string) bool {
	log.Println("pre pre force deploying", name)
	for _, n := range forceDeployNames {
		log.Println("pre force deploying", name)
		if n == name {
			log.Println("force deploying", n)
			return true
		}
	}
	return false
}

// Execute will create destination and all transivite dependencies
func Execute(
	ctx context.Context,
	log logrus.FieldLogger,
	ec *ethclient.Client,
	to *bind.TransactOpts,
	ensAddress common.Address,
	destination string,
	params map[string]common.Address,
) error {

	en, err := ens.NewENS(to, ensAddress, ec)
	if err != nil {
		return errors.Wrap(err, "while creating ens instance")
	}

	var resolve func(name string) (common.Address, error)

	resolve = func(name string) (common.Address, error) {

		if name == "ETH.tokencard.eth" {
			return common.Address{}, nil
		}

		addr, ok := newENSAddresses[name]
		if ok {
			return addr, nil
		}

		if strings.HasSuffix(name, ".eth") {
			addr, err := en.Addr(name)
			if err != nil || addr == zeroAddress || forceDeploy(name) {
				target, ok := targets[name]
				if !ok {
					return zeroAddress, errors.Errorf("could not find target %s", name)
				}

				log.Infof("Not found %s in ENS, creating", name)
				addr, err = target(ctx, ec, to, resolve)
				if err != nil {
					return zeroAddress, errors.Wrapf(err, "while resolving %s", name)
				}
				log.Infof("Created %s at %s, setting ENS", name, addr.Hex())

				err = setENSAddress(ctx, ec, to, ensAddress, name, addr)
				if err != nil {
					return zeroAddress, err
				}
			}

			log.Infof("Found %s at %s", name, addr.Hex())
			return addr, nil
		}

		{
			addr, ok := params[name]
			if ok {
				return addr, nil
			}
		}

		{
			target, ok := targets[name]
			if !ok {
				return zeroAddress, errors.Errorf("could not find target %s", name)
			}
			if ok {
				log.Infof("Executing target %s", name)
				addr, err := target(ctx, ec, to, resolve)
				if err != nil {
					return zeroAddress, errors.Wrapf(err, "while resolving %s", name)
				}
				return addr, err
			}
		}

		return zeroAddress, errors.Errorf("not yet implemented: %s", name)
	}

	_, err = resolve(destination)
	return err

}
