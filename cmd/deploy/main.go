package main

import (
	"context"
	"fmt"
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/tokencard/contracts/v2/cmd/deploy/target"
	"github.com/tokencard/contracts/v2/pkg/bindings"
	"gopkg.in/urfave/cli.v1"
)

// Running test chain: docker run --rm -p 8545:8545 parity/parity:v2.5.5-stable --config dev --jsonrpc-interface=0.0.0.0 --jsonrpc-port=8545

func main() {
	app := cli.NewApp()

	app.Name = "deploy"
	app.Usage = "deploy infrastructure contracts to an ethereum network"

	addCommand(app, "dev", func(d *deployer, args []string) error {
		return target.Execute(
			context.Background(),
			d.log,
			d.ethClient,
			d.transactOpts,
			d.ensAddress,
			target.DevEnvironment,
			map[string]common.Address{
				target.ENSName:                  d.ensAddress,
				target.DevControllerAddressName: common.HexToAddress("0x00442c3994155d098f4ffc2e28f76d810ce9209c"),
			},
		)
	})

	addCommand(app, "cache-contracts", func(d *deployer, args []string) error {

		if len(args) != 1 {
			return errors.New("number of contracts not provided")
		}

		numberOfContracts, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			return errors.Wrap(err, "while parsing number of contracts")
		}

		fmt.Println("number of contracts to cache", numberOfContracts)

		wda, err := d.ens.Addr(target.WalletDeployerName)
		if err != nil {
			return errors.Wrapf(err, "while looking up %s", target.WalletDeployerName)
		}

		wd, err := bindings.NewWalletDeployer(wda, d.ethClient)

		nonce, err := d.ethClient.PendingNonceAt(d.ctx, d.transactOpts.From)
		if err != nil {
			return err
		}

		txHashes := []common.Hash{}

		for i := 0; i < int(numberOfContracts); i++ {
			to := *d.transactOpts
			to.Nonce = big.NewInt(int64(nonce + uint64(i)))
			to.GasPrice = d.gasPrice

			tx, err := wd.CacheWallet(&to)
			if err != nil {
				return errors.Wrap(err, "while creating cache wallet transaction")
			}

			fmt.Println(tx.Hash().Hex())
			txHashes = append(txHashes, tx.Hash())
		}

		return nil
	})

	addCommand(app, "deploy-ens", func(d *deployer, args []string) error {
		d.log.Info("Deploying ENS ...")
		return d.deployENS()
	})

	app.RunAndExitOnError()

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
				Name:   "gas-price",
				Usage:  "gas price for the transaction (in GWEI)",
				EnvVar: "GAS_PRICE",
				Value:  0,
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
