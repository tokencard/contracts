package wallet_deployer_test

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/tokencard/contracts/pkg/bindings"
	. "github.com/tokencard/contracts/test/shared"
	"github.com/tokencard/ethertest"
)

func TestWalletDeployer(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "WalletDeployer Suite")
}

var _ = BeforeSuite(func() {
	TestRig.AddCoverageForContracts(
		"../../build/wallet-deployer/combined.json",
		"../../contracts",
	)
})

var allPassed = true

var _ = AfterEach(func() {
	td := CurrentGinkgoTestDescription()
	if td.Failed {
		allPassed = false
	}

})

var _ = AfterSuite(func() {
	if allPassed {
		TestRig.ExpectMinimumCoverage("wallet-deployer.sol", 100.00)
		TestRig.PrintGasUsage(os.Stdout)
	}
})

var _ = AfterEach(func() {
	td := CurrentGinkgoTestDescription()

	if td.Failed {
		fmt.Fprintf(GinkgoWriter, "\nLast Executed Smart Contract Line for %s:%d\n", td.FileName, td.LineNumber)
		fmt.Fprintln(GinkgoWriter, TestRig.LastExecuted())
	}
	err := Backend.Close()
	Expect(err).ToNot(HaveOccurred())
})

var walletDeployerAddress common.Address
var walletDeployer *bindings.WalletDeployer

func isSuccessful(tx *types.Transaction) bool {
	r, err := Backend.TransactionReceipt(context.Background(), tx.Hash())
	Expect(err).ToNot(HaveOccurred())
	return r.Status == types.ReceiptStatusSuccessful
}

var _ = BeforeEach(func() {
	err := InitializeBackend()
	Expect(err).ToNot(HaveOccurred())

	var tx *types.Transaction
	walletDeployerAddress, tx, walletDeployer, err = bindings.DeployWalletDeployer(BankAccount.TransactOpts(), Backend, ENSRegistryAddress, TokenWhitelistName, ControllerName, LicenceName, EthToWei(100))
	Expect(err).ToNot(HaveOccurred())
	Backend.Commit()
	Expect(isSuccessful(tx)).To(BeTrue())

})

var _ = Describe("WalletDeployer", func() {
	When("no contracts are cached", func() {
		When("a controller deploys a wallet", func() {
			BeforeEach(func() {
				tx, err := walletDeployer.DeployWallet(Controller.TransactOpts(), Owner.Address())
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})

			It("should return new wallet address", func() {
				addr, err := walletDeployer.Deployed(nil, Owner.Address())
				Expect(err).ToNot(HaveOccurred())
				Expect(addr).ToNot(Equal(common.HexToAddress("0x0")))
			})

			It("should emit ContractDeployed event", func() {

				addr, err := walletDeployer.Deployed(nil, Owner.Address())
				Expect(err).ToNot(HaveOccurred())

				it, err := walletDeployer.FilterWalletDeployed(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(it.Next()).To(BeTrue())
				Expect(it.Event.Wallet).To(Equal(addr))
				Expect(it.Event.Owner).To(Equal(Owner.Address()))
			})

			It("should emit ContractCached event", func() {

				addr, err := walletDeployer.Deployed(nil, Owner.Address())
				Expect(err).ToNot(HaveOccurred())

				it, err := walletDeployer.FilterWalletCached(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(it.Next()).To(BeTrue())
				Expect(it.Event.Wallet).To(Equal(addr))
			})

			When("new wallet owner sets the spend limit", func() {
				var wallet *bindings.Wallet
				var tx *types.Transaction

				BeforeEach(func() {
					walletAddress, err := walletDeployer.Deployed(nil, Owner.Address())

					wallet, err = bindings.NewWallet(walletAddress, Backend)
					Expect(err).ToNot(HaveOccurred())

					tx, err = wallet.SetSpendLimit(Owner.TransactOpts(), EthToWei(1))
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
				})

				It("should succeed", func() {
					Expect(isSuccessful(tx)).To(BeTrue())
				})
			})
		})

		When("a random account tries to deploy a wallet", func() {

			var tx *types.Transaction

			BeforeEach(func() {
				var err error
				tx, err = walletDeployer.DeployWallet(RandomAccount.TransactOpts(ethertest.WithGasLimit(5000000)), Owner.Address())
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
			})

			It("should fail", func() {
				Expect(isSuccessful(tx)).To(BeFalse())
			})
		})
	})

	When("a random account tries to cache a wallet", func() {

		var tx *types.Transaction

		BeforeEach(func() {
			var err error
			tx, err = walletDeployer.CacheWallet(RandomAccount.TransactOpts(ethertest.WithGasLimit(5000000)))
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
		})

		It("should fail", func() {
			Expect(isSuccessful(tx)).To(BeFalse())
		})
	})

	When("one contract is cached", func() {
		BeforeEach(func() {
			tx, err := walletDeployer.CacheWallet(Controller.TransactOpts())
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeTrue())
		})

		It("should cache the wallet", func() {
			addr, err := walletDeployer.Cached(nil, big.NewInt(0))
			Expect(err).ToNot(HaveOccurred())
			Expect(addr).ToNot(Equal(common.HexToAddress("0x0")))
		})

		It("should emit ContractCached event", func() {
			addr, err := walletDeployer.Cached(nil, big.NewInt(0))
			Expect(err).ToNot(HaveOccurred())

			it, err := walletDeployer.FilterWalletCached(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(it.Next()).To(BeTrue())
			Expect(it.Event.Wallet).To(Equal(addr))
		})

		When("a controller deploys a wallet", func() {
			BeforeEach(func() {
				tx, err := walletDeployer.DeployWallet(Controller.TransactOpts(), Owner.Address())
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})

			It("should return new wallet address", func() {
				addr, err := walletDeployer.Deployed(nil, Owner.Address())
				Expect(err).ToNot(HaveOccurred())
				Expect(addr).ToNot(Equal(common.HexToAddress("0x0")))
			})

			When("New wallet owner sets the spend limit", func() {
				var wallet *bindings.Wallet
				var tx *types.Transaction

				BeforeEach(func() {
					walletAddress, err := walletDeployer.Deployed(nil, Owner.Address())

					wallet, err = bindings.NewWallet(walletAddress, Backend)
					Expect(err).ToNot(HaveOccurred())

					tx, err = wallet.SetSpendLimit(Owner.TransactOpts(), EthToWei(1))
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
				})

				It("should succeed", func() {
					Expect(isSuccessful(tx)).To(BeTrue())
				})
			})

			It("should emit ContractDeployed event", func() {

				addr, err := walletDeployer.Deployed(nil, Owner.Address())
				Expect(err).ToNot(HaveOccurred())

				it, err := walletDeployer.FilterWalletDeployed(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(it.Next()).To(BeTrue())
				Expect(it.Event.Wallet).To(Equal(addr))
				Expect(it.Event.Owner).To(Equal(Owner.Address()))
			})


		})
	})

})
