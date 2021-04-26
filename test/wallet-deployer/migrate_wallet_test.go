package wallet_deployer_test

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/tokencard/contracts/v3/pkg/bindings"
	. "github.com/tokencard/contracts/v3/test/shared"
	"github.com/tokencard/ethertest"
)

var _ = Describe("Migrate Wallet", func() {

	var MigratedWallet *bindings.Wallet

	When("no wallets are cached and a controller migrates a Wallet with non-initialized security settings", func() {

		BeforeEach(func() {

			RandomProxyAddress := deployInitProxy(Owner.Address(), EthToWei(2))

			tx, err := WalletDeployer.MigrateWallet(Controller.TransactOpts(), Owner.Address(), RandomProxyAddress, false, false, EthToWei(2), []common.Address{common.HexToAddress("0x1"), common.HexToAddress("0x2")})
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

		It("should NOT add the whitelisted addresses to the whitelist", func() {
			isWhitelisted, err := MigratedWallet.WhitelistMap(nil, common.HexToAddress("0x1"))
			Expect(err).ToNot(HaveOccurred())
			Expect(isWhitelisted).To(BeFalse())

			isWhitelisted, err = MigratedWallet.WhitelistMap(nil, common.HexToAddress("0x2"))
			Expect(err).ToNot(HaveOccurred())
			Expect(isWhitelisted).To(BeFalse())
		})

		It("should NOT update the daily limit", func() {
			sl, err := MigratedWallet.DailyLimitValue(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(sl.String()).To(Equal(EthToWei(10000).String()))
		})
	})

	When("no wallets are cached and a controller migrates a Wallet and sends 1000 wei", func() {

		var RandomProxyAddress common.Address
		var RandomOwner common.Address
		var tx *types.Transaction
		var err error

		BeforeEach(func() {
			privateKey, _ := crypto.GenerateKey()
			RandomOwner = crypto.PubkeyToAddress(privateKey.PublicKey)

			RandomProxyAddress = deployInitProxy(RandomOwner, EthToWei(2))

			tx, err = WalletDeployer.MigrateWallet(
				Controller.TransactOpts(ethertest.WithValue(big.NewInt(1000))),
				RandomOwner,
				RandomProxyAddress,
				true,
				true,
				MweiToWei(5000),
				[]common.Address{common.HexToAddress("0x1"),
					common.HexToAddress("0x2")},
			)
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeTrue())
		})

		It("DeployedWallets should return the new Wallet address", func() {
			addr, err := WalletDeployer.DeployedWallets(nil, RandomOwner)
			Expect(err).ToNot(HaveOccurred())
			Expect(addr).ToNot(Equal(common.HexToAddress("0x0")))
		})

		It("should emit a MigratedWallet event", func() {

			addr, err := WalletDeployer.DeployedWallets(nil, RandomOwner)
			Expect(err).ToNot(HaveOccurred())

			it, err := WalletDeployer.FilterMigratedWallet(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(it.Next()).To(BeTrue())
			Expect(it.Event.Wallet).To(Equal(addr))
			Expect(it.Event.OldWallet).To(Equal(RandomProxyAddress))
			Expect(it.Event.Owner).To(Equal(RandomOwner))
			Expect(it.Event.Paid.String()).To(Equal("1000"))
		})

		It("should have cached Wallet count 0", func() {
			ccc, err := WalletCache.CachedWalletsCount(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(ccc.String()).To(Equal("0"))
		})

		It("should increase the owner's balance by 1000 wei", func() {
			b, e := Backend.BalanceAt(context.Background(), RandomOwner, nil)
			Expect(e).ToNot(HaveOccurred())
			Expect(b.String()).To(Equal("1000"))
		})

		It("should fail if a wallet is already deployed/migrated for this owner", func() {
			tx, err = WalletDeployer.MigrateWallet(Controller.TransactOpts(ethertest.WithGasLimit(5000000)), RandomOwner, RandomProxyAddress, true, true, EthToWei(2), []common.Address{common.HexToAddress("0x1"), common.HexToAddress("0x2")})
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeFalse())
			returnData, _ := ethCall(tx)
			Expect(string(returnData[len(returnData)-64:])).To(ContainSubstring("wallet already deployed for owner"))
		})

		It("should emit CachedWallet event", func() {

			addr, err := WalletDeployer.DeployedWallets(nil, RandomOwner)
			Expect(err).ToNot(HaveOccurred())

			it, err := WalletCache.FilterCachedWallet(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(it.Next()).To(BeTrue())
			Expect(it.Event.Wallet).To(Equal(addr))
		})

		When("a new Wallet is migrated", func() {

			BeforeEach(func() {
				MigratedWalletAddress, err := WalletDeployer.DeployedWallets(nil, RandomOwner)
				MigratedWallet, err = bindings.NewWallet(MigratedWalletAddress, Backend)
				Expect(err).ToNot(HaveOccurred())
			})

			It("should emit a daily limit set event", func() {
				it, err := MigratedWallet.FilterSetDailyLimit(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(it.Next()).To(BeTrue())
				evt := it.Event
				Expect(it.Next()).To(BeFalse())
				Expect(evt.Sender).To(Equal(WalletDeployerAddress))
				Expect(evt.Amount).To(Equal(MweiToWei(5000)))
			})

			It("should decrease the available amount to 5k USD", func() {
				av, err := MigratedWallet.DailyLimitAvailable(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(av.String()).To(Equal(MweiToWei(5000).String()))
			})

			It("should decrement the daily limit to 5k USD", func() {
				sl, err := MigratedWallet.DailyLimitValue(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(sl.String()).To(Equal(MweiToWei(5000).String()))
			})

			It("should update the Whitelist initializedWhitelist flag", func() {
				initialized, err := MigratedWallet.IsSetWhitelist(nil)
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

	}) //no wallets cached

	When("the wrong owner is passed in", func() {
		It("should fail", func() {
			RandomProxyAddress := deployInitProxy(Owner.Address(), EthToWei(2))

			tx, err := WalletDeployer.MigrateWallet(Controller.TransactOpts(ethertest.WithGasLimit(5000000)), Controller.Address(), RandomProxyAddress, false, false, EthToWei(2), []common.Address{common.HexToAddress("0x1"), common.HexToAddress("0x2")})
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeFalse())
			returnData, _ := ethCall(tx)
			Expect(string(returnData[len(returnData)-64:])).To(ContainSubstring("owner mismatch"))
		})
	})

	When("a random account tries to migrate a Wallet", func() {
		It("should fail", func() {
			tx, err := WalletDeployer.MigrateWallet(RandomAccount.TransactOpts(ethertest.WithGasLimit(5000000)), Owner.Address(), RandomAccount.Address(), false, false, EthToWei(1), []common.Address{common.HexToAddress("0x1"), common.HexToAddress("0x2")})
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeFalse())
		})
	})

	When("one wallet is cached", func() {
		BeforeEach(func() {
			tx, err := WalletCache.CacheWallet(Controller.TransactOpts())
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeTrue())
		})

		When("a controller migrates a Wallet", func() {

			var RandomProxyAddress common.Address
			var tx *types.Transaction
			var err error

			BeforeEach(func() {

				RandomProxyAddress = deployInitProxy(Owner.Address(), EthToWei(2))

				tx, err = WalletDeployer.MigrateWallet(Controller.TransactOpts(), Owner.Address(), RandomProxyAddress, false, false, EthToWei(1), []common.Address{common.HexToAddress("0x1"), common.HexToAddress("0x2")})
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
				Expect(it.Event.OldWallet).To(Equal(RandomProxyAddress))
				Expect(it.Event.Owner).To(Equal(Owner.Address()))
			})

			It("should have cached Wallet count 0", func() {
				ccc, err := WalletCache.CachedWalletsCount(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(ccc.String()).To(Equal("0"))
			})

			It("should emit CachedWallet event", func() {

				addr, err := WalletDeployer.DeployedWallets(nil, Owner.Address())
				Expect(err).ToNot(HaveOccurred())

				it, err := WalletCache.FilterCachedWallet(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(it.Next()).To(BeTrue())
				Expect(it.Event.Wallet).To(Equal(addr))
			})

		})
	})

})
