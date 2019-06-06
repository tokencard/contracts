package main

import (
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
	"github.com/tokencard/contracts/pkg/bindings/externals/ens"
	externalens "github.com/tokencard/contracts/pkg/bindings/externals/ens"
)

func (d *deployer) deployENS() error {

	// name := "ENS"

	ensRegistryAddress, tx, ensRegistry, err := externalens.DeployENSRegistry(d.transactOpts, d.ethClient)
	if err != nil {
		return err
	}

	err = d.waitForAndConfirmTransaction(tx.Hash())
	if err != nil {
		return errors.Wrapf(err, "while waiting for transaction")
	}

	d.log.Infof("ens registry address: %s", ensRegistryAddress.Hex())

	{
		tx, err = ensRegistry.SetSubnodeOwner(d.transactOpts, EnsNode(""), LabelHash("eth"), d.controllerOwner)
		if err != nil {
			return err
		}

		err = d.waitForAndConfirmTransaction(tx.Hash())
		if err != nil {
			return errors.Wrapf(err, "while waiting for transaction")
		}
	}

	{
		tx, err = ensRegistry.SetSubnodeOwner(d.transactOpts, EnsNode("eth"), LabelHash("tokencard"), d.controllerOwner)
		if err != nil {
			return err
		}

		err = d.waitForAndConfirmTransaction(tx.Hash())
		if err != nil {
			return errors.Wrapf(err, "while waiting for transaction")
		}
	}

	{
		ensResolverAddress, tx, _, err := ens.DeployPublicResolver(d.transactOpts, d.ethClient, ensRegistryAddress)
		if err != nil {
			return errors.Wrap(err, "while deploying resolver")
		}

		err = d.waitForAndConfirmTransaction(tx.Hash())

		if err != nil {
			return errors.Wrapf(err, "while waiting for transaction")
		}

		d.log.Infof("ens resolver address: %s", ensResolverAddress.Hex())

		tx, err = ensRegistry.SetResolver(d.transactOpts, EnsNode("tokencard.eth"), ensResolverAddress)

		if err != nil {
			return errors.Wrap(err, "while setting resolver")
		}

		err = d.waitForAndConfirmTransaction(tx.Hash())

		if err != nil {
			return errors.Wrapf(err, "while waiting for transaction")
		}

	}

	return nil

}

func (d *deployer) setAddress(name string, address common.Address) error {
	er, err := externalens.NewENSRegistry(d.ensAddress, d.ethClient)
	if err != nil {
		return errors.Wrap(err, "while creating ens registry instance")
	}

	resolverAddress, err := er.Resolver(nil, EnsNode(name))

	if err != nil {
		return errors.Wrapf(err, "while getting resolver address of %s", name)
	}

	if resolverAddress == zeroAddress {
		d.log.Info("resolver not found, looking for parent's resolver...")

		parent, label := EnsParentNode(name)

		tx, err := er.SetSubnodeOwner(d.transactOpts, parent, label, d.controllerOwner)
		if err != nil {
			return errors.Wrap(err, "while setting sub-domain resolver")
		}

		err = d.waitForAndConfirmTransaction(tx.Hash())
		if err != nil {
			return errors.Wrapf(err, "while waiting for transaction")
		}

		resolverAddress, err = er.Resolver(nil, parent)
		if err != nil {
			return errors.Wrapf(err, "while getting resolver address for parent of %q", name)
		}

		if resolverAddress == zeroAddress {
			d.log.Info("resolver of parent is not set, deploying a new resolver ...")

			resolverAddress, tx, _, err = ens.DeployPublicResolver(d.transactOpts, d.ethClient, d.ensAddress)
			if err != nil {
				return errors.Wrap(err, "while deploying resolver")
			}

			err = d.waitForAndConfirmTransaction(tx.Hash())

			if err != nil {
				return errors.Wrapf(err, "while waiting for transaction")
			}

		}

		tx, err = er.SetResolver(d.transactOpts, EnsNode(name), resolverAddress)
		if err != nil {
			return errors.Wrapf(err, "while setting resolver for %q", name)
		}

		err = d.waitForAndConfirmTransaction(tx.Hash())
		if err != nil {
			return errors.Wrapf(err, "while waiting for transaction")
		}

	}

	d.log.Infof("resolver address %s", resolverAddress.Hex())

	re, err := externalens.NewPublicResolver(resolverAddress, d.ethClient)
	if err != nil {
		return errors.Wrap(err, "while creating resolver instance")
	}

	tx, err := re.SetAddr(d.transactOpts, EnsNode(name), address)
	if err != nil {
		return errors.Wrap(err, "while setting address in resolver")
	}

	err = d.waitForAndConfirmTransaction(tx.Hash())
	if err != nil {
		return errors.Wrapf(err, "while waiting for transaction")
	}

	return nil

}

// Lifted from https://github.com/tronprotocol/tron-demo/blob/master/demo/go-client-api/vendor/github.com/ethereum/go-ethereum/contracts/ens/ens.go
func EnsParentNode(name string) (common.Hash, common.Hash) {
	parts := strings.SplitN(name, ".", 2)
	label := crypto.Keccak256Hash([]byte(parts[0]))
	if len(parts) == 1 {
		return [32]byte{}, label
	} else {
		parentNode, parentLabel := EnsParentNode(parts[1])
		return crypto.Keccak256Hash(parentNode[:], parentLabel[:]), label
	}
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
