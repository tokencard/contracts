package target

import (
	"context"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	externalens "github.com/tokencard/contracts/v2/pkg/bindings/externals/ens"
)

var newENSAddresses = map[string]common.Address{}

func setENSAddress(
	ctx context.Context,
	ec *ethclient.Client,
	to *bind.TransactOpts,
	ensAddress common.Address,
	name string,
	address common.Address,
) error {

	er, err := externalens.NewENSRegistry(ensAddress, ec)
	if err != nil {
		return errors.Wrap(err, "while creating ens registry instance")
	}

	resolverAddress, err := er.Resolver(nil, EnsNode(name))

	if err != nil {
		return errors.Wrapf(err, "while getting resolver address of %s", name)
	}

	if resolverAddress == zeroAddress {

		parent, label := EnsParentNode(name)

		tx, err := er.SetSubnodeOwner(to, parent, label, to.From)
		if err != nil {
			return errors.Wrap(err, "while setting sub-domain resolver")
		}

		err = waitForAndConfirmTransaction(ctx, ec, tx)
		if err != nil {
			return errors.Wrapf(err, "while waiting for transaction")
		}

		resolverAddress, err = er.Resolver(nil, parent)
		if err != nil {
			return errors.Wrapf(err, "while getting resolver address for parent of %q", name)
		}

		if resolverAddress == zeroAddress {

			resolverAddress, tx, _, err = externalens.DeployPublicResolver(to, ec, ensAddress)
			if err != nil {
				return errors.Wrap(err, "while deploying resolver")
			}

			err = waitForAndConfirmTransaction(ctx, ec, tx)

			if err != nil {
				return errors.Wrapf(err, "while waiting for transaction")
			}

		}

		tx, err = er.SetResolver(to, EnsNode(name), resolverAddress)
		if err != nil {
			return errors.Wrapf(err, "while setting resolver for %q", name)
		}

		err = waitForAndConfirmTransaction(ctx, ec, tx)
		if err != nil {
			return errors.Wrapf(err, "while waiting for transaction")
		}

	}

	re, err := externalens.NewPublicResolver(resolverAddress, ec)
	if err != nil {
		return errors.Wrap(err, "while creating resolver instance")
	}

	tx, err := re.SetAddr(to, EnsNode(name), address)
	if err != nil {
		return errors.Wrap(err, "while setting address in resolver")
	}

	err = waitForAndConfirmTransaction(ctx, ec, tx)
	if err != nil {
		return errors.Wrapf(err, "while waiting for transaction")
	}

	newENSAddresses[name] = address

	return nil

}

// Lifted from https://github.com/tronprotocol/tron-demo/blob/master/demo/go-client-api/vendor/github.com/ethereum/go-ethereum/contracts/ens/ens.go
func EnsParentNode(name string) (common.Hash, common.Hash) {
	parts := strings.SplitN(name, ".", 2)
	label := crypto.Keccak256Hash([]byte(parts[0]))
	if len(parts) == 1 {
		return [32]byte{}, label
	}

	parentNode, parentLabel := EnsParentNode(parts[1])
	return crypto.Keccak256Hash(parentNode[:], parentLabel[:]), label
}

func LabelHash(label string) common.Hash {
	return crypto.Keccak256Hash([]byte(label))
}

func EnsNode(name string) common.Hash {
	if name == "" {
		return common.Hash{}
	}
	parentNode, parentLabel := EnsParentNode(name)
	return crypto.Keccak256Hash(parentNode[:], parentLabel[:])
}
