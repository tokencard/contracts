package wallet_deployer_test

import (
	"context"
	"fmt"
	"os"
	"testing"
    . "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
    . "github.com/tokencard/contracts/v2/test/shared"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/tokencard/contracts/v2/pkg/bindings"
	"github.com/tokencard/ethertest"
)

func TestWalletDeployer(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "WalletDeployer Suite")
}

var _ = BeforeSuite(func() {
	TestRig.AddCoverageForContracts(
		"../../build/walletDeployer/combined.json",
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
		TestRig.ExpectMinimumCoverage("walletDeployer.sol", 100.00)
		TestRig.PrintGasUsage(os.Stdout)
	}
})

var _ = AfterEach(func() {
	td := CurrentGinkgoTestDescription()

	if td.Failed {
		fmt.Fprintf(GinkgoWriter, "\nLast Executed Smart Contract Line for %s:%d\n", td.FileName, td.LineNumber)
		fmt.Fprintln(GinkgoWriter, TestRig.LastExecuted())
	}
})

var _ = AfterEach(func() {
	err := Backend.Close()
	Expect(err).ToNot(HaveOccurred())

})

var WalletCacheAddress common.Address
var WalletCache *bindings.WalletCache

var WalletDeployerAddress common.Address
var WalletDeployer *bindings.WalletDeployer

func isSuccessful(tx *types.Transaction) bool {
	r, err := Backend.TransactionReceipt(context.Background(), tx.Hash())
	Expect(err).ToNot(HaveOccurred())
	return r.Status == types.ReceiptStatusSuccessful
}

var WalletDeployerNode = EnsNode("wallet-deployer.tokencard.eth")

var _ = BeforeEach(func() {
	err := InitializeBackend()
	Expect(err).ToNot(HaveOccurred())
	var tx *types.Transaction
	WalletCacheAddress, tx, WalletCache, err = bindings.DeployWalletCache(BankAccount.TransactOpts(), Backend, ENSRegistryAddress, EthToWei(1))
	Expect(err).ToNot(HaveOccurred())
    WalletDeployerAddress, tx, WalletDeployer, err = bindings.DeployWalletDeployer(BankAccount.TransactOpts(), Backend, ENSRegistryAddress, WalletCacheAddress)
	Expect(err).ToNot(HaveOccurred())
	Backend.Commit()
	Expect(isSuccessful(tx)).To(BeTrue())

    {
        // set WalletDeployer node owner
        tx, err = ENSRegistry.SetSubnodeOwner(BankAccount.TransactOpts(), EnsNode("tokencard.eth"), LabelHash("wallet-deployer"), BankAccount.Address())
        Expect(err).ToNot(HaveOccurred())
		Backend.Commit()
		// Register WalletDeployer with ENS
		tx, err = ENSRegistry.SetResolver(BankAccount.TransactOpts(), WalletDeployerNode, ENSResolverAddress)
		Expect(err).ToNot(HaveOccurred())
		Backend.Commit()
		Expect(isSuccessful(tx)).To(BeTrue())
		tx, err = ENSResolver.SetAddr(BankAccount.TransactOpts(), WalletDeployerNode, WalletDeployerAddress)
		Expect(err).ToNot(HaveOccurred())
		Backend.Commit()
		Expect(isSuccessful(tx)).To(BeTrue())
	}
})

var _ = Describe("Wallet Deployer", func() {

    It("should return walletCache addres", func() {
        wca, err := WalletDeployer.WalletCache(nil)
        Expect(err).ToNot(HaveOccurred())
        Expect(wca).To(Equal(WalletCacheAddress))
    })

	When("no Wallets are cached", func() {

		When("a controller deploys a Wallet", func() {
			BeforeEach(func() {
				tx, err := WalletDeployer.DeployWallet(Controller.TransactOpts(), Owner.Address())
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})

			It("should return a new Wallet address", func() {
				addr, err := WalletDeployer.DeployedWallets(nil, Owner.Address())
				Expect(err).ToNot(HaveOccurred())
				Expect(addr).ToNot(Equal(common.HexToAddress("0x0")))
			})

			It("should emit a DeployedWallet event", func() {

				addr, err := WalletDeployer.DeployedWallets(nil, Owner.Address())
				Expect(err).ToNot(HaveOccurred())

				it, err := WalletDeployer.FilterDeployedWallet(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(it.Next()).To(BeTrue())
				Expect(it.Event.Wallet).To(Equal(addr))
				Expect(it.Event.Owner).To(Equal(Owner.Address()))
			})

			It("should have a cached Wallet count of 0", func() {
				ccc, err := WalletCache.CachedWalletsCount(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(ccc.String()).To(Equal("0"))
			})

			It("should emit a CachedWallet event", func() {

				addr, err := WalletDeployer.DeployedWallets(nil, Owner.Address())
				Expect(err).ToNot(HaveOccurred())

				it, err := WalletCache.FilterCachedWallet(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(it.Next()).To(BeTrue())
				Expect(it.Event.Wallet).To(Equal(addr))
			})

			When("new Wallet owner sets the spend limit", func() {
				var NewWallet *bindings.Wallet

				BeforeEach(func() {
					NewWalletAddress, err := WalletDeployer.DeployedWallets(nil, Owner.Address())

					NewWallet, err = bindings.NewWallet(NewWalletAddress, Backend)
					Expect(err).ToNot(HaveOccurred())

				})

				It("should succeed", func() {
					tx, err := NewWallet.SetSpendLimit(Owner.TransactOpts(), EthToWei(1))
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeTrue())
				})
			})
		})

		When("a random account tries to deploy a Wallet", func() {

			It("should fail", func() {
				tx, err := WalletDeployer.DeployWallet(RandomAccount.TransactOpts(ethertest.WithGasLimit(5000000)), Owner.Address())
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeFalse())
			})
		})
	})

	When("one Wallet is cached", func() {
		BeforeEach(func() {
			tx, err := WalletCache.CacheWallet(Controller.TransactOpts())
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeTrue())
		})

		When("a controller deploys a Wallet", func() {
			BeforeEach(func() {
				tx, err := WalletDeployer.DeployWallet(Controller.TransactOpts(), Owner.Address())
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})

			It("should return new Wallet address", func() {
				addr, err := WalletDeployer.DeployedWallets(nil, Owner.Address())
				Expect(err).ToNot(HaveOccurred())
				Expect(addr).ToNot(Equal(common.HexToAddress("0x0")))
			})

			When("New Wallet owner sets the spend limit", func() {
				var NewWallet *bindings.Wallet

				BeforeEach(func() {
					NewWalletAddress, err := WalletDeployer.DeployedWallets(nil, Owner.Address())

					NewWallet, err = bindings.NewWallet(NewWalletAddress, Backend)
					Expect(err).ToNot(HaveOccurred())

					tx, err := NewWallet.SetSpendLimit(Owner.TransactOpts(), FinneyToWei(500))
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeTrue())
				})

				It("should lower the spend available to 500 Finney", func() {
					av, err := NewWallet.SpendLimitAvailable(nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(av.String()).To(Equal(FinneyToWei(500).String()))
				})

				It("should have a spend limit of 500 Finney", func() {
					sl, err := NewWallet.SpendLimitValue(nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(sl.String()).To(Equal(FinneyToWei(500).String()))
				})
			})

			It("should emit a DeployedWallet event", func() {

				addr, err := WalletDeployer.DeployedWallets(nil, Owner.Address())
				Expect(err).ToNot(HaveOccurred())

				it, err := WalletDeployer.FilterDeployedWallet(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(it.Next()).To(BeTrue())
				Expect(it.Event.Wallet).To(Equal(addr))
				Expect(it.Event.Owner).To(Equal(Owner.Address()))
			})

		})
	})

})
