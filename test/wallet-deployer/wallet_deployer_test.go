package wallet_deployer_test

import (
	"github.com/ethereum/go-ethereum/common"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/tokencard/contracts/v2/pkg/bindings"
	. "github.com/tokencard/contracts/v2/test/shared"
	"github.com/tokencard/ethertest"
)

var _ = Describe("Wallet Deployer", func() {

	It("should point to the right controller node", func() {
		wcn, err := WalletDeployer.WalletCacheNode(nil)
		Expect(err).ToNot(HaveOccurred())
		Expect(common.Hash(wcn)).To(Equal(EnsNode("wallet-cache.tokencard.eth")))
	})

	It("should point to the right wallet cache node", func() {
		on, err := WalletDeployer.ControllerNode(nil)
		Expect(err).ToNot(HaveOccurred())
		Expect(common.Hash(on)).To(Equal(EnsNode("controller.tokencard.eth")))
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

			When("new Wallet owner sets the daily limit", func() {
				var NewWallet *bindings.Wallet

				BeforeEach(func() {
					NewWalletAddress, err := WalletDeployer.DeployedWallets(nil, Owner.Address())
					Expect(err).ToNot(HaveOccurred())

					NewWallet, err = bindings.NewWallet(NewWalletAddress, Backend)
					Expect(err).ToNot(HaveOccurred())

				})

				It("should succeed", func() {
					tx, err := NewWallet.SetDailyLimit(Owner.TransactOpts(), EthToWei(1))
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

			When("New Wallet owner sets the daily limit", func() {
				var NewWallet *bindings.Wallet

				BeforeEach(func() {
					NewWalletAddress, err := WalletDeployer.DeployedWallets(nil, Owner.Address())

					NewWallet, err = bindings.NewWallet(NewWalletAddress, Backend)
					Expect(err).ToNot(HaveOccurred())

					tx, err := NewWallet.SetDailyLimit(Owner.TransactOpts(), FinneyToWei(500))
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeTrue())
				})

				It("should lower the spend available to 500 Finney", func() {
					av, err := NewWallet.DailyLimitAvailable(nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(av.String()).To(Equal(FinneyToWei(500).String()))
				})

				It("should have a daily limit of 500 Finney", func() {
					sl, err := NewWallet.DailyLimitValue(nil)
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
