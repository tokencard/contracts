package main

import (
	"log"
	"math/big"
	"os"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
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

	addCommand(app, "setup-dev", func(d *deployer, args []string) error {

		ethOwners := []common.Address{
			common.HexToAddress("0x00442c3994155d098f4ffc2e28f76d810ce9209c"),
			common.HexToAddress("0x0076ed2dd9f7dc78e3f336141329f8784d8cd564"),
			common.HexToAddress("0x00f137e9bfe37cc015f11cec8339cc2f1a3ae3a6"),
			common.HexToAddress("0x6b0c56d1ad5144b4d37fa6e27dc9afd5c2435c3b"),
			common.HexToAddress("0x7464a376e9c27c8f9cbe1e38443cf994169f0839"),
			common.HexToAddress("0xa04a110c983ee867d70a450ab29464bb04bf5763"),
		}

		millionETH := &big.Int{}
		_, valid := millionETH.SetString("1000000000000000000000000", 10)
		if !valid {
			return errors.New("could not parse 1M ETH to Wei")
		}

		for _, eo := range ethOwners {
			d.transferETH(eo, millionETH)
		}

		d.log.Info("Deploying ENS ...")
		err := d.deployENS()
		if err != nil {
			return err
		}

		err = d.deployController()
		if err != nil {
			return err
		}

		err = d.addController(common.HexToAddress("0x00442c3994155d098f4ffc2e28f76d810ce9209c"))
		if err != nil {
			return err
		}

		err = d.deployOraclizeResolver()
		if err != nil {
			return err
		}

		err = d.deployOracle()
		if err != nil {
			return err
		}

		err = d.deployWalletDeployer()
		if err != nil {
			return err
		}

		return nil
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
