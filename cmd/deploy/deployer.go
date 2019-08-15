package main

import (
	"context"
	"io/ioutil"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/contracts/ens"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/tyler-smith/go-bip39"
	"gopkg.in/urfave/cli.v1"
)

type deployer struct {
	transactOpts *bind.TransactOpts
	ethClient    *ethclient.Client
	ens          *ens.ENS

	ensAddress         common.Address
	ensResolverAddress common.Address

	log                     logrus.FieldLogger
	controllerOwner         common.Address
	initialNonce            int64
	gasPrice                *big.Int
	ctx                     context.Context
	oraclizeResolverAddress common.Address
}

func runWithDeployer(fn func(*deployer, []string) error) func(c *cli.Context) error {

	return func(c *cli.Context) error {
		var txOpt *bind.TransactOpts

		if c.IsSet("key-mnemonic") {

			logrus.Info("using provided mnemonic")

			mnemonic := c.String("key-mnemonic")
			seed := bip39.NewSeed(mnemonic, "")

			wallet, err := hdwallet.NewFromSeed(seed)
			if err != nil {
				return err
			}

			path, err := hdwallet.ParseDerivationPath("m/44'/60'/0'/0/0")
			if err != nil {
				return err
			}

			account, err := wallet.Derive(path, false)
			if err != nil {
				return err
			}

			txOpt = &bind.TransactOpts{
				Signer: func(signer types.Signer, addr common.Address, tx *types.Transaction) (*types.Transaction, error) {
					pk, err := wallet.PrivateKey(account)
					if err != nil {
						return nil, err
					}
					return types.SignTx(tx, signer, pk)
				},
				From: account.Address,
			}
		} else if c.IsSet("key-file") {

			logrus.Infof("using keystore at %s", c.String("key-file"))

			keyJSON, err := ioutil.ReadFile(c.String("key-file"))
			if err != nil {
				return err
			}

			decrypted, err := keystore.DecryptKey(keyJSON, c.String("passphrase"))
			if err != nil {
				return err
			}

			txOpt = &bind.TransactOpts{
				Signer: func(signer types.Signer, addr common.Address, tx *types.Transaction) (*types.Transaction, error) {
					return types.SignTx(tx, signer, decrypted.PrivateKey)
				},
				From: decrypted.Address,
			}

		} else {
			return errors.New("neither key file nor key mnemonic used")
		}

		ec, err := ethclient.Dial(c.String("ethereum"))
		if err != nil {
			return err
		}

		defer ec.Close()

		ensAddress := common.HexToAddress(c.String("ens-address"))

		logrus.Infof("using ENS address %s", ensAddress.Hex())
		logrus.Infof("sending from address %s", txOpt.From.Hex())

		en, err := ens.NewENS(txOpt, ensAddress, ec)
		if err != nil {
			return err
		}

		var gasPrice *big.Int

		if c.IsSet("gas-price") {
			gasPrice = big.NewInt(c.Int64("gas-price"))
			gasPrice.Mul(big.NewInt(1000000000))
		}

		d := &deployer{
			transactOpts:            txOpt,
			ens:                     en,
			ensAddress:              ensAddress,
			controllerOwner:         txOpt.From,
			ctx:                     context.Background(),
			ethClient:               ec,
			log:                     logrus.New(),
			oraclizeResolverAddress: common.HexToAddress(c.String("oraclize-resolver-address")),
			initialNonce:            c.Int64("initial-nonce"),
			gasPrice:                gasPrice,
		}

		balance, err := ec.BalanceAt(d.ctx, d.transactOpts.From, nil)
		if err != nil {
			return errors.Wrap(err, "while getting balance")
		}

		d.log.Infof("Balance of %s: %s", txOpt.From.Hex(), balance.String())

		return fn(d, c.Args())

	}
}

var zeroAddress = common.HexToAddress("0x0")

const waitForMiningTimeout = 2 * 60 * time.Second

func (d *deployer) waitForTransactionToBeMined(txHash common.Hash) error {
	ctx, cancel := context.WithTimeout(d.ctx, waitForMiningTimeout)
	defer cancel()

	for {
		var pending bool
		_, pending, err := d.ethClient.TransactionByHash(ctx, txHash)
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

func (d *deployer) ensureTransactionSuccess(txHash common.Hash) error {
	rcpt, err := d.ethClient.TransactionReceipt(d.ctx, txHash)
	if err != nil {
		return err
	}

	if rcpt.Status != types.ReceiptStatusSuccessful {
		return errors.Errorf("transaction %s failed", txHash.Hex())
	}

	return nil

}

func (d *deployer) waitForAndConfirmTransaction(txHash common.Hash) error {
	err := d.waitForTransactionToBeMined(txHash)
	if err != nil {
		return err
	}
	return d.ensureTransactionSuccess(txHash)
}
