package main

import (
	"context"
	"io/ioutil"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/contracts/ens"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sirupsen/logrus"
	"gopkg.in/urfave/cli.v1"
)

// Running test chain: docker run --rm -p 8545:8545 parity/parity:v2.4.6 --config dev --jsonrpc-interface=0.0.0.0 --jsonrpc-port=8545

func main() {
	app := cli.NewApp()

	app.Name = "deploy"
	app.Usage = "deploy infrastructure contracts to an ethereum network"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "ethereum",
			Usage:  "Ethereum node's address.",
			EnvVar: "ETHEREUM_ADDRESS",
			Value:  "http://localhost:8545",
		},
		cli.StringFlag{
			Name:   "ens-address",
			Usage:  "Ethereum name service contract address.",
			EnvVar: "ENS_ADDRESS",
		},
		cli.StringFlag{
			Name:   "ens-resolver-address",
			Usage:  "Tokencard's ENS resolver address.",
			EnvVar: "ENS_RESOLVER_ADDRESS",
		},
		cli.StringFlag{
			Name:   "oraclize-resolver-address",
			Usage:  "Oraclize resolver address.",
			EnvVar: "ORACLIZE_RESOLVER_ADDRESS",
			Value:  "0x1",
		},
		cli.StringFlag{
			Name:   "key-file",
			Usage:  "JSON key file location.",
			EnvVar: "KEY_FILE",
			Value:  "dev.key.json",
		},
		cli.StringFlag{
			Name:   "passphrase",
			Usage:  "Keystore file passphrase.",
			EnvVar: "PASSPHRASE",
			Value:  "",
		},
	}

	app.Action = run

	app.Commands = []cli.Command{
		cli.Command{
			Name:   "deploy-ens",
			Action: deployENS,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:   "ethereum",
					Usage:  "Ethereum node's address.",
					EnvVar: "ETHEREUM_ADDRESS",
					Value:  "http://localhost:8545",
				},
				cli.StringFlag{
					Name:   "key-file",
					Usage:  "JSON key file location.",
					EnvVar: "KEY_FILE",
					Value:  "dev.key.json",
				},
				cli.StringFlag{
					Name:   "passphrase",
					Usage:  "Keystore file passphrase.",
					EnvVar: "PASSPHRASE",
					Value:  "",
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}

func deployENS(c *cli.Context) error {
	keyJSON, err := ioutil.ReadFile(c.String("key-file"))
	if err != nil {
		return err
	}

	decrypted, err := keystore.DecryptKey(keyJSON, c.String("passphrase"))
	if err != nil {
		return err
	}

	ec, err := ethclient.Dial(c.String("ethereum"))
	if err != nil {
		return err
	}

	defer ec.Close()

	txOpt := &bind.TransactOpts{
		Signer: func(signer types.Signer, addr common.Address, tx *types.Transaction) (*types.Transaction, error) {
			return types.SignTx(tx, signer, decrypted.PrivateKey)
		},
		From: decrypted.Address,
	}

	d := &deployer{
		transactOpts:    txOpt,
		controllerOwner: decrypted.Address,
		ctx:             context.Background(),
		ethClient:       ec,
		log:             logrus.New(),
	}

	// bal := ec.BalanceAt(context.Background(), decrypted.PrivateKey, nil)

	return d.deployENS()

}

func run(c *cli.Context) error {

	keyJSON, err := ioutil.ReadFile(c.String("key-file"))
	if err != nil {
		return err
	}

	decrypted, err := keystore.DecryptKey(keyJSON, c.String("passphrase"))
	if err != nil {
		return err
	}

	ec, err := ethclient.Dial(c.String("ethereum"))
	if err != nil {
		return err
	}

	defer ec.Close()

	txOpt := &bind.TransactOpts{
		Signer: func(signer types.Signer, addr common.Address, tx *types.Transaction) (*types.Transaction, error) {
			return types.SignTx(tx, signer, decrypted.PrivateKey)
		},
		From: decrypted.Address,
	}

	ensAddress := common.HexToAddress(c.String("ens-address"))

	logrus.Info("ens address", ensAddress.Hex())

	en, err := ens.NewENS(txOpt, ensAddress, ec)
	if err != nil {
		return err
	}

	d := &deployer{
		transactOpts:            txOpt,
		ens:                     en,
		ensAddress:              ensAddress,
		controllerOwner:         decrypted.Address,
		ctx:                     context.Background(),
		ethClient:               ec,
		log:                     logrus.New(),
		oraclizeResolverAddress: common.HexToAddress(c.String("oraclize-resolver-address")),
	}

	err = d.deployController()
	if err != nil {
		return err
	}

	// TODO deploy oracle - needs fixing, ATM it's failing
	// err = d.deployOracle()
	// if err != nil {
	// 	return err
	// }

	err = d.deployWalletDeployer()
	if err != nil {
		return err
	}

	return nil
}
