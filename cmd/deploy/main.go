package main

import (
	"log"
	"math/big"
	"os"
	"strconv"

	"github.com/ethereum/go-ethereum/contracts/ens"
	"github.com/pkg/errors"
	"github.com/tokencard/contracts/pkg/bindings"
	"gopkg.in/urfave/cli.v1"
)

// Running test chain: docker run --rm -p 8545:8545 parity/parity:v2.4.6 --config dev --jsonrpc-interface=0.0.0.0 --jsonrpc-port=8545

func main() {
	app := cli.NewApp()

	app.Name = "deploy"
	app.Usage = "deploy infrastructure contracts to an ethereum network"

	addCommand(app, "wallet-deployer", func(d *deployer, args []string) error {
		return d.deployWalletDeployer()
	})

	addCommand(app, "wallet-deployer-only", func(d *deployer, args []string) error {
		contractAddress, tx, _, err := bindings.DeployWalletDeployer(d.transactOpts, d.ethClient, d.ensAddress, ens.EnsNode(oracleName), ens.EnsNode(controllerName), oneETH)
		if err != nil {
			return errors.Wrap(err, "while deploying wallet deployer")
		}

		d.log.Infof("Wallet deployer address: %s", contractAddress.Hex())
		d.log.Infof("Transaction hash: %s", tx.Hash().Hex())
		return nil
	})

	addCommand(app, "deploy-ens", func(d *deployer, args []string) error {
		d.log.Info("Deploying ENS ...")
		return d.deployENS()
	})

	addCommand(app, "cache-wallets", func(d *deployer, args []string) error {

		toCache := 0

		if len(args) > 1 {
			return errors.New("only one argument (number of instances) is accepted")
		}

		if len(args) == 1 {
			var err error
			toCache, err = strconv.Atoi(args[0])
			if err != nil {
				return errors.Wrapf(err, "while parsing integer %q", args[0])
			}
		}

		d.log.Infof("Caching %d instances of wallet", toCache)

		addr, err := d.ens.Addr(walletDeployerName)
		if err != nil {
			return err
		}

		if addr == zeroAddress {
			return errors.New("deployer not found")
		}

		wd, err := bindings.NewWalletDeployer(addr, d.ethClient)
		if err != nil {
			return err
		}

		txOpts := *d.transactOpts

		// use 1 GWEI for contract caching
		txOpts.GasPrice = big.NewInt(1000000000)

		nonce, err := d.ethClient.PendingNonceAt(d.ctx, txOpts.From)
		if err != nil {
			return errors.Wrapf(err, "while getting pending nonce for %s", txOpts.From.Hex())
		}

		if d.initialNonce >= 0 {
			nonce = uint64(d.initialNonce)
		}

		for i := 0; i < toCache; i++ {
			txOpts.Nonce = big.NewInt(int64(nonce) + int64(i))
			tx, err := wd.CacheWallet(&txOpts)
			if err != nil {
				return errors.Wrap(err, "while sending transaction")
			}
			d.log.Infof("Created transaction %s", tx.Hash().Hex())
		}

		return nil
	})

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}

func addCommand(app *cli.App, name string, fn func(d *deployer, args []string) error) {
	app.Commands = append(app.Commands, cli.Command{
		Name: name,
		Action: runWithDeployer(func(d *deployer, args []string) error {
			return fn(d, args)
		}),
		Flags: []cli.Flag{
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
				Value:  "0x314159265dd8dbb310642f98f50c066173c1259b",
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
			cli.StringFlag{
				Name:   "key-mnemonic",
				Usage:  "Mnemonic (BIP 39) to derive the key.",
				EnvVar: "KEY_MNEMONIC",
				Value:  "",
			},
			cli.Int64Flag{
				Name:   "initial-nonce",
				Usage:  "Initial nonce to be used for transaction",
				EnvVar: "INITIAL_NONCE",
				Value:  -1,
			},
		},
	})
}
