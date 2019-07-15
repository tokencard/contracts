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
})

var WalletDeployerAddress common.Address
var WalletDeployer *bindings.WalletDeployer

func isSuccessful(tx *types.Transaction) bool {
	r, err := Backend.TransactionReceipt(context.Background(), tx.Hash())
	Expect(err).ToNot(HaveOccurred())
	return r.Status == types.ReceiptStatusSuccessful
}

var _ = BeforeEach(func() {
	err := InitializeBackend()
	Expect(err).ToNot(HaveOccurred())

	var tx *types.Transaction
	WalletDeployerAddress, tx, WalletDeployer, err = bindings.DeployWalletDeployer(BankAccount.TransactOpts(), Backend, ENSRegistryAddress, ControllerName, EthToWei(1))
	Expect(err).ToNot(HaveOccurred())
	Backend.Commit()
	Expect(isSuccessful(tx)).To(BeTrue())

})

var _ = Describe("WalletDeployer", func() {
	When("no Wallets are cached", func() {

		It("should have a default spend limit of 1 ETH", func() {
			sl, err := WalletDeployer.DefaultSpendLimit(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(sl.String()).To(Equal(EthToWei(1).String()))
		})

		It("should point to the right oracle node name", func() {
			on, err := WalletDeployer.OracleName(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(common.Hash(on)).To(Equal(EnsNode("oracle.tokencard.eth")))
		})

        It("should point to the right controller node name", func() {
			on, err := WalletDeployer.ControllerName(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(common.Hash(on)).To(Equal(EnsNode("controller.tokencard.eth")))
		})

		It("should have cached Wallet count 0", func() {
			ccc, err := WalletDeployer.CachedWalletCount(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(ccc.String()).To(Equal("0"))
		})

		When("a controller deploys a Wallet", func() {
			BeforeEach(func() {
				tx, err := WalletDeployer.DeployWallet(Controller.TransactOpts(), Owner.Address())
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})

			It("should return new Wallet address", func() {
				addr, err := WalletDeployer.Deployed(nil, Owner.Address())
				Expect(err).ToNot(HaveOccurred())
				Expect(addr).ToNot(Equal(common.HexToAddress("0x0")))
			})

			It("should emit DeployedWallet event", func() {

				addr, err := WalletDeployer.Deployed(nil, Owner.Address())
				Expect(err).ToNot(HaveOccurred())

				it, err := WalletDeployer.FilterDeployedWallet(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(it.Next()).To(BeTrue())
				Expect(it.Event.Wallet).To(Equal(addr))
				Expect(it.Event.Owner).To(Equal(Owner.Address()))
			})

			It("should have cached Wallet count 0", func() {
				ccc, err := WalletDeployer.CachedWalletCount(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(ccc.String()).To(Equal("0"))
			})

			It("should emit CachedWallet event", func() {

				addr, err := WalletDeployer.Deployed(nil, Owner.Address())
				Expect(err).ToNot(HaveOccurred())

				it, err := WalletDeployer.FilterCachedWallet(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(it.Next()).To(BeTrue())
				Expect(it.Event.Wallet).To(Equal(addr))
			})

			When("new Wallet owner sets the spend limit", func() {
				var NewWallet *bindings.Wallet

				BeforeEach(func() {
					NewWalletAddress, err := WalletDeployer.Deployed(nil, Owner.Address())

					NewWallet, err = bindings.NewWallet(NewWalletAddress, Backend)
					Expect(err).ToNot(HaveOccurred())
				})

				It("should succeed", func() {
                    tx, err := NewWallet.InitializeSpendLimit(Owner.TransactOpts(), EthToWei(1))
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

	When("a random account tries to cache a Wallet", func() {

		It("should succeed", func() {
			tx, err := WalletDeployer.CacheWallet(RandomAccount.TransactOpts(ethertest.WithGasLimit(5000000)))
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
            Expect(isSuccessful(tx)).To(BeTrue())
		})
	})

	When("one Wallet is cached", func() {
		BeforeEach(func() {
			tx, err := WalletDeployer.CacheWallet(Controller.TransactOpts())
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeTrue())
		})

		It("should cache the Wallet", func() {
			addr, err := WalletDeployer.Cached(nil, big.NewInt(0))
			Expect(err).ToNot(HaveOccurred())
			Expect(addr).ToNot(Equal(common.HexToAddress("0x0")))
		})

		It("should have cached Wallet count 1", func() {
			ccc, err := WalletDeployer.CachedWalletCount(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(ccc.String()).To(Equal("1"))
		})

		It("should emit CachedWallet event", func() {
			addr, err := WalletDeployer.Cached(nil, big.NewInt(0))
			Expect(err).ToNot(HaveOccurred())

			it, err := WalletDeployer.FilterCachedWallet(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(it.Next()).To(BeTrue())
			Expect(it.Event.Wallet).To(Equal(addr))
		})

		When("a controller deploys a Wallet", func() {
			BeforeEach(func() {
				tx, err := WalletDeployer.DeployWallet(Controller.TransactOpts(), Owner.Address())
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})

			It("should return new Wallet address", func() {
				addr, err := WalletDeployer.Deployed(nil, Owner.Address())
				Expect(err).ToNot(HaveOccurred())
				Expect(addr).ToNot(Equal(common.HexToAddress("0x0")))
			})

			When("New Wallet owner sets the spend limit", func() {

				var NewWallet *bindings.Wallet

				BeforeEach(func() {
					NewWalletAddress, err := WalletDeployer.Deployed(nil, Owner.Address())
					NewWallet, err = bindings.NewWallet(NewWalletAddress, Backend)
					Expect(err).ToNot(HaveOccurred())

                    tx, err := NewWallet.InitializeSpendLimit(Owner.TransactOpts(), FinneyToWei(500))
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeTrue())
				})

				It("should lower the spend available to 500 Finney", func() {
					av, err := NewWallet.SpendAvailable(nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(av.String()).To(Equal(FinneyToWei(500).String()))
				})

				It("should have a spend limit of 500 Finney", func() {
					sl, err := NewWallet.SpendLimit(nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(sl.String()).To(Equal(FinneyToWei(500).String()))
				})
			})

			It("should emit DeployedWallet event", func() {

				addr, err := WalletDeployer.Deployed(nil, Owner.Address())
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
