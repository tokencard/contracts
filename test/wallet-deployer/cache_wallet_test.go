package wallet_deployer_test

import (
    "math/big"
    "context"
	"github.com/ethereum/go-ethereum/common"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tokencard/contracts/v2/test/shared"
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

        It("should have a cached Wallet count of 0", func() {
			ccc, err := WalletCache.CachedWalletsCount(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(ccc.String()).To(Equal("0"))
		})

		When("ETH is sent", func() {

			BeforeEach(func() {
				BankAccount.MustTransfer(Backend, WalletCacheAddress, EthToWei(1))
			})

			It("should receive it (payable)", func() {
				b, e := Backend.BalanceAt(context.Background(), WalletCacheAddress, nil)
				Expect(e).ToNot(HaveOccurred())
				Expect(b.String()).To(Equal(EthToWei(1).String()))
			})

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

            When("someone pops", func() {
                BeforeEach(func() {
        			tx, err := WalletCache.WalletCachePop(RandomAccount.TransactOpts())
        			Expect(err).ToNot(HaveOccurred())
        			Backend.Commit()
        			Expect(isSuccessful(tx)).To(BeTrue())
        		})

                It("should have a cached Wallet count of 1", func() {
        			cwc, err := WalletCache.CachedWalletsCount(nil)
        			Expect(err).ToNot(HaveOccurred())
        			Expect(cwc.String()).To(Equal("1"))
        		})
            })
        })

        When("someone pops", func() {
            BeforeEach(func() {
    			tx, err := WalletCache.WalletCachePop(RandomAccount.TransactOpts())
    			Expect(err).ToNot(HaveOccurred())
    			Backend.Commit()
    			Expect(isSuccessful(tx)).To(BeTrue())
    		})

            It("should have a cached Wallet count of 0", func() {
    			cwc, err := WalletCache.CachedWalletsCount(nil)
    			Expect(err).ToNot(HaveOccurred())
    			Expect(cwc.String()).To(Equal("0"))
    		})

            It("should emit a CachedWallet event", func() {
    			it, err := WalletCache.FilterCachedWallet(nil)
    			Expect(err).ToNot(HaveOccurred())
    			Expect(it.Next()).To(BeTrue())
                Expect(it.Next()).To(BeFalse())
    			Expect(it.Event.Wallet).ToNot(Equal(common.HexToAddress("0x0")))
    		})
        })
	})

    When("someone pops without having cached a wallet first", func() {
        BeforeEach(func() {
            tx, err := WalletCache.WalletCachePop(RandomAccount.TransactOpts())
            Expect(err).ToNot(HaveOccurred())
            Backend.Commit()
            Expect(isSuccessful(tx)).To(BeTrue())
        })

        It("should have a cached Wallet count of 0", func() {
            cwc, err := WalletCache.CachedWalletsCount(nil)
            Expect(err).ToNot(HaveOccurred())
            Expect(cwc.String()).To(Equal("0"))
        })

        It("should emit a CachedWallet event", func() {
			it, err := WalletCache.FilterCachedWallet(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(it.Next()).To(BeTrue())
            Expect(it.Next()).To(BeFalse())
		})
    })

})
