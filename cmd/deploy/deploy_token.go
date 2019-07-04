package main

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/tokencard/contracts/pkg/bindings"
	"github.com/tokencard/contracts/pkg/bindings/mocks"
)

func (d *deployer) deployToken(name string) error {

	d.log.Infof("Deploying Token %s", name)

	addr, err := d.ens.Addr(name)

	if err == nil && addr != zeroAddress {
		d.log.Infof("%s found, with address %s", name, addr.Hex())
		return nil
	}

	millionToken := &big.Int{}
	_, valid := millionToken.SetString("1000000000000000000000000", 10)
	if !valid {
		return errors.New("could not parse 1M")
	}

	tokenAddress, tx, token, err := mocks.DeployToken(d.transactOpts, d.ethClient)
	d.log.Infof("%s address is %s", name, tokenAddress.Hex())
	err = d.waitForAndConfirmTransaction(tx.Hash())
	if err != nil {
		return errors.Wrap(err, "while waiting for transaction")
	}

	tx, err = token.Credit(d.transactOpts, d.controllerOwner, millionToken)
	if err != nil {
		return errors.Wrap(err, "while waiting for transaction")
	}

	err = d.waitForAndConfirmTransaction(tx.Hash())
	if err != nil {
		return errors.Wrap(err, "while waiting for transaction")
	}

	oracleAddress, err := d.ens.Addr(oracleName)

	if err != nil {
		return errors.Wrapf(err, "while resolving ENS %s", oracleAddress)
	}

	d.log.Infof("Oracle address %s", oracleAddress.Hex())

	oracle, err := bindings.NewOracle(oracleAddress, d.ethClient)
	if err != nil {
		return errors.Wrap(err, "while creating Oracle binding")
	}

	tx, err = oracle.AddTokens(d.transactOpts, []common.Address{tokenAddress}, stringsToByte32(name), []*big.Int{decimalsToMagnitude(9)}, big.NewInt(20180913153211))
	if err != nil {
		return errors.Wrap(err, "while sending AddTokens transaction")
	}

	err = d.waitForAndConfirmTransaction(tx.Hash())
	if err != nil {
		return errors.Wrap(err, "while waiting for transaction")
	}

	ensName := fmt.Sprintf("%s.tokencard.eth", name)

	d.log.Infof("Setting ENS for %s to %s", ensName, tokenAddress.Hex())

	err = d.setAddress(ensName, tokenAddress)
	if err != nil {
		return errors.Wrapf(err, "while setting ENS enry for %s", name)
	}

	return nil

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
