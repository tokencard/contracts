package wallet_deployer_test

import (
	"github.com/ethereum/go-ethereum/common"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/tokencard/contracts/v2/pkg/bindings"
	. "github.com/tokencard/contracts/v2/test/shared"
	"github.com/tokencard/ethertest"
)

var _ = Describe("Migrate Wallet", func() {

	var MigratedWallet *bindings.Wallet

	When("no wallets are cached and a controller migrates a Wallet with non-initialized security settings", func() {

		BeforeEach(func() {
			tx, err := WalletDeployer.MigrateWallet(Controller.TransactOpts(), Owner.Address(), RandomAccount.Address(), false, false, false, EthToWei(2), FinneyToWei(1), []common.Address{common.HexToAddress("0x1"), common.HexToAddress("0x2")})
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeTrue())

			MigratedWalletAddress, err := WalletDeployer.DeployedWallets(nil, Owner.Address())

			MigratedWallet, err = bindings.NewWallet(MigratedWalletAddress, Backend)
			Expect(err).ToNot(HaveOccurred())
		})

		It("should NOT set the whitelist flag", func() {
			initialized, err := MigratedWallet.IsSetWhitelist(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(initialized).To(BeFalse())
		})

		It("should NOT set the GasTopUpLimit flag", func() {
			initialized, err := MigratedWallet.GasTopUpLimitUpdateable(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(initialized).To(BeFalse())
		})

		It("should NOT set the initializedSpendLimit flag", func() {
			initialized, err := MigratedWallet.SpendLimitUpdateable(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(initialized).To(BeFalse())
		})

		It("should NOT add the whitelisted addresses to the whitelist", func() {
			isWhitelisted, err := MigratedWallet.WhitelistMap(nil, common.HexToAddress("0x1"))
			Expect(err).ToNot(HaveOccurred())
			Expect(isWhitelisted).To(BeFalse())

			isWhitelisted, err = MigratedWallet.WhitelistMap(nil, common.HexToAddress("0x2"))
			Expect(err).ToNot(HaveOccurred())
			Expect(isWhitelisted).To(BeFalse())
		})

		It("should NOT update the spend limit to 2 ETH", func() {
			sl, err := MigratedWallet.SpendLimitValue(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(sl.String()).To(Equal(EthToWei(1).String()))
		})

		It("should NOT update the top up limit to 1 finney", func() {
			sl, err := MigratedWallet.GasTopUpLimitValue(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(sl.String()).To(Equal(FinneyToWei(500).String()))
		})

	})

	When("no wallets are cached and a controller migrates a Wallet", func() {

		BeforeEach(func() {
			tx, err := WalletDeployer.MigrateWallet(Controller.TransactOpts(), Owner.Address(), RandomAccount.Address(), true, true, true, EthToWei(2), FinneyToWei(1), []common.Address{common.HexToAddress("0x1"), common.HexToAddress("0x2")})
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeTrue())
		})

		It("DeployedWallets should return the new Wallet address", func() {
			addr, err := WalletDeployer.DeployedWallets(nil, Owner.Address())
			Expect(err).ToNot(HaveOccurred())
			Expect(addr).ToNot(Equal(common.HexToAddress("0x0")))
		})

		It("should emit a MigratedWallet event", func() {

			addr, err := WalletDeployer.DeployedWallets(nil, Owner.Address())
			Expect(err).ToNot(HaveOccurred())

			it, err := WalletDeployer.FilterMigratedWallet(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(it.Next()).To(BeTrue())
			Expect(it.Event.Wallet).To(Equal(addr))
			Expect(it.Event.OldWallet).To(Equal(RandomAccount.Address()))
			Expect(it.Event.Owner).To(Equal(Owner.Address()))
		})

		It("should have cached Wallet count 0", func() {
			ccc, err := WalletDeployer.CachedWalletsCount(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(ccc.String()).To(Equal("0"))
		})

		It("should emit CachedWallet event", func() {

			addr, err := WalletDeployer.DeployedWallets(nil, Owner.Address())
			Expect(err).ToNot(HaveOccurred())

			it, err := WalletDeployer.FilterCachedWallet(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(it.Next()).To(BeTrue())
			Expect(it.Event.Wallet).To(Equal(addr))
		})

		When("a new Wallet is migrated", func() {

			BeforeEach(func() {
				MigratedWalletAddress, err := WalletDeployer.DeployedWallets(nil, Owner.Address())
				MigratedWallet, err = bindings.NewWallet(MigratedWalletAddress, Backend)
				Expect(err).ToNot(HaveOccurred())
			})

			It("should emit a spend limit set event", func() {
				it, err := MigratedWallet.FilterSetSpendLimit(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(it.Next()).To(BeTrue())
				evt := it.Event
				Expect(it.Next()).To(BeFalse())
				Expect(evt.Sender).To(Equal(WalletDeployerAddress))
				Expect(evt.Amount).To(Equal(EthToWei(2)))
			})

			It("should keep the available spend Limit  to 1 ETH", func() {
				av, err := MigratedWallet.SpendLimitAvailable(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(av.String()).To(Equal(EthToWei(1).String()))
			})

			It("should update the spend limit to 2 ETH", func() {
				sl, err := MigratedWallet.SpendLimitValue(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(sl.String()).To(Equal(EthToWei(2).String()))
			})

			It("should update the top up limit to 1 finney", func() {
				sl, err := MigratedWallet.GasTopUpLimitValue(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(sl.String()).To(Equal(FinneyToWei(1).String()))
			})

			It("should update the SpendLimit set flag", func() {
				initialized, err := MigratedWallet.SpendLimitUpdateable(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(initialized).To(BeTrue())
			})

			It("should emit a setTopUpLimit event", func() {
				it, err := MigratedWallet.FilterSetGasTopUpLimit(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(it.Next()).To(BeTrue())
				evt := it.Event
				Expect(it.Next()).To(BeFalse())
				Expect(evt.Sender).To(Equal(WalletDeployerAddress))
				Expect(evt.Amount).To(Equal(FinneyToWei(1)))
			})

			It("should update the Whitelist initializedWhitelist flag", func() {
				initialized, err := MigratedWallet.IsSetWhitelist(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(initialized).To(BeTrue())
			})

			It("should update the TopUpLimit initializedTopup flag", func() {
				initialized, err := MigratedWallet.GasTopUpLimitUpdateable(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(initialized).To(BeTrue())
			})

			It("should add the whitelisted addresses to the whitelist", func() {
				isWhitelisted, err := MigratedWallet.WhitelistMap(nil, common.HexToAddress("0x1"))
				Expect(err).ToNot(HaveOccurred())
				Expect(isWhitelisted).To(BeTrue())

				isWhitelisted, err = MigratedWallet.WhitelistMap(nil, common.HexToAddress("0x2"))
				Expect(err).ToNot(HaveOccurred())
				Expect(isWhitelisted).To(BeTrue())
			})

			It("should emit AddedToWhitelist event", func() {
				it, err := MigratedWallet.FilterAddedToWhitelist(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(it.Next()).To(BeTrue())
				evt := it.Event
				Expect(it.Next()).To(BeFalse())
				Expect(evt.Sender).To(Equal(WalletDeployerAddress))
				Expect(evt.Addresses).To(Equal([]common.Address{common.HexToAddress("0x1"), common.HexToAddress("0x2")}))
			})
		})

	}) //no wallets chached

	When("a random account tries to migrate a Wallet", func() {
		It("should fail", func() {
			tx, err := WalletDeployer.MigrateWallet(RandomAccount.TransactOpts(ethertest.WithGasLimit(5000000)), Owner.Address(), RandomAccount.Address(), false, false, false, EthToWei(1), FinneyToWei(2), []common.Address{common.HexToAddress("0x1"), common.HexToAddress("0x2")})
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeFalse())
		})
	})

	When("one wallet is cached", func() {
		BeforeEach(func() {
			tx, err := WalletDeployer.CacheWallet(Controller.TransactOpts())
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeTrue())
		})

		When("a controller migrates a Wallet", func() {
			BeforeEach(func() {
				tx, err := WalletDeployer.MigrateWallet(Controller.TransactOpts(), Owner.Address(), RandomAccount.Address(), false, false, false, EthToWei(1), FinneyToWei(2), []common.Address{common.HexToAddress("0x1"), common.HexToAddress("0x2")})
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})

			It("should emit MigratedWallet event", func() {

				addr, err := WalletDeployer.DeployedWallets(nil, Owner.Address())
				Expect(err).ToNot(HaveOccurred())

				it, err := WalletDeployer.FilterMigratedWallet(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(it.Next()).To(BeTrue())
				Expect(it.Event.Wallet).To(Equal(addr))
				Expect(it.Event.OldWallet).To(Equal(RandomAccount.Address()))
				Expect(it.Event.Owner).To(Equal(Owner.Address()))
			})

			It("should have cached Wallet count 0", func() {
				ccc, err := WalletDeployer.CachedWalletsCount(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(ccc.String()).To(Equal("0"))
			})

			It("should emit CachedWallet event", func() {

				addr, err := WalletDeployer.DeployedWallets(nil, Owner.Address())
				Expect(err).ToNot(HaveOccurred())

				it, err := WalletDeployer.FilterCachedWallet(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(it.Next()).To(BeTrue())
				Expect(it.Event.Wallet).To(Equal(addr))
			})

		})
	})

})
