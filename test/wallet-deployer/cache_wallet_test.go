package wallet_deployer_test

import (
	"math/big"
    "github.com/ethereum/go-ethereum/common"
    "github.com/tokencard/ethertest"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tokencard/contracts/v3/test/shared"
)

var _ = Describe("Cache Wallet", func() {

	When("no Wallets are cached", func() {

		It("should have a default spend limit of 1 ETH", func() {
			sl, err := WalletCache.DefaultSpendLimit(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(sl.String()).To(Equal(EthToWei(1).String()))
		})

		It("should point to the right tokenwhitelist node name", func() {
			on, err := WalletCache.TokenWhitelistNode(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(common.Hash(on)).To(Equal(EnsNode("token-whitelist.tokencard.eth")))
		})

		It("should point to the right controller node name", func() {
			on, err := WalletCache.ControllerNode(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(common.Hash(on)).To(Equal(EnsNode("controller.tokencard.eth")))
		})

		It("should point to the right licence node name", func() {
			on, err := WalletCache.LicenceNode(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(common.Hash(on)).To(Equal(EnsNode("licence.tokencard.eth")))
		})

		It("should point to the right wallet deployer node name", func() {
			on, err := WalletCache.WalletDeployerNode(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(common.Hash(on)).To(Equal(EnsNode("wallet-deployer.tokencard.eth")))
		})

		It("should have a cached Wallet count of 0", func() {
			ccc, err := WalletCache.CachedWalletsCount(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(ccc.String()).To(Equal("0"))
		})
	})

	When("no wallets are cached and a random account caches a wallet", func() {

		BeforeEach(func() {
			tx, err := WalletCache.CacheWallet(RandomAccount.TransactOpts())
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeTrue())
		})

		It("should cache the Wallet", func() {
			addr, err := WalletCache.CachedWallets(nil, big.NewInt(0))
			Expect(err).ToNot(HaveOccurred())
			Expect(addr).ToNot(Equal(common.HexToAddress("0x0")))
		})

		It("should have a cached Wallet count of 1", func() {
			cwc, err := WalletCache.CachedWalletsCount(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(cwc.String()).To(Equal("1"))
		})

		It("should emit a CachedWallet event", func() {
			addr, err := WalletCache.CachedWallets(nil, big.NewInt(0))
			Expect(err).ToNot(HaveOccurred())

			it, err := WalletCache.FilterCachedWallet(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(it.Next()).To(BeTrue())
			Expect(it.Next()).To(BeFalse())
			Expect(it.Event.Wallet).To(Equal(addr))
		})

		When("a random account caches again", func() {
			BeforeEach(func() {
				tx, err := WalletCache.CacheWallet(RandomAccount.TransactOpts())
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})

			It("should cache the Wallet", func() {
				addr, err := WalletCache.CachedWallets(nil, big.NewInt(1))
				Expect(err).ToNot(HaveOccurred())
				Expect(addr).ToNot(Equal(common.HexToAddress("0x0")))
			})

			It("should have a cached Wallet count of 2", func() {
				cwc, err := WalletCache.CachedWalletsCount(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(cwc.String()).To(Equal("2"))
			})

			It("should emit 2 CachedWallet events", func() {
				addr, err := WalletCache.CachedWallets(nil, big.NewInt(0))
				Expect(err).ToNot(HaveOccurred())

				it, err := WalletCache.FilterCachedWallet(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(it.Next()).To(BeTrue())
				Expect(it.Event.Wallet).To(Equal(addr))

				addr, err = WalletCache.CachedWallets(nil, big.NewInt(1))
				Expect(err).ToNot(HaveOccurred())
				Expect(it.Next()).To(BeTrue())
				Expect(it.Next()).To(BeFalse())
				Expect(it.Event.Wallet).To(Equal(addr))
			})

			When("someone tries to pop", func() {
				It("should fail", func() {
					tx, err := WalletCache.WalletCachePop(Controller.TransactOpts(ethertest.WithGasLimit(6000000)))
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeFalse())
				})
			})
		})

		When("a random account tries to pop", func() {
			It("should fail", func() {
				tx, err := WalletCache.WalletCachePop(RandomAccount.TransactOpts(ethertest.WithGasLimit(6000000)))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeFalse())
			})
		})
	})

	When("someone tries to pop without having cached a wallet first", func() {
		It("should fail", func() {
			tx, err := WalletCache.WalletCachePop(RandomAccount.TransactOpts(ethertest.WithGasLimit(6000000)))
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeFalse())
		})
	})

})
