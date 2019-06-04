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

func main() {
	app := cli.NewApp()

	app.Name = "deploy"
	app.Usage = "deploy infrastructure contracts to an ethereum network"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "ethereum",
			Usage:  "Ethereum node's address.",
			EnvVar: "ETHEREUM_ADDRESS",
		},
		cli.StringFlag{
			Name:   "ens-address",
			Usage:  "Ethereum name service contract address.",
			EnvVar: "ENS_ADDRESS",
		},
		cli.StringFlag{
			Name:   "oraclize-resolver-address",
			Usage:  "Oraclize resolver address.",
			EnvVar: "ORACLIZE_RESOLVER_ADDRESS",
		},
		cli.StringFlag{
			Name:   "key-file",
			Usage:  "JSON key file location.",
			EnvVar: "KEY_FILE",
		},
		cli.StringFlag{
			Name:   "passphrase",
			Usage:  "Keystore file passphrase.",
			EnvVar: "PASSPHRASE",
		},
	}

	app.Action = run

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

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

	// signer := types.NewEIP155Signer(new(big.Int).SetUint64(c.Uint64("chain-id")))

	// ens.
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

	en, err := ens.NewENS(txOpt, ensAddress, ec)
	if err != nil {
		return err
	}

	d := &deployer{
		transactOpts:    txOpt,
		ens:             en,
		ensAddress:      ensAddress,
		controllerOwner: decrypted.Address,
		ctx:             context.Background(),
		ethClient:       ec,
		log:             logrus.New(),
	}

	err = d.deployController()
	if err != nil {
		return err
	}

	err = d.deployOracle()
	if err != nil {
		return err
	}

	return nil
}
