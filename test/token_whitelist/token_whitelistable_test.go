package token_whitelist_test

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tokencard/contracts/v3/test/shared"
	"github.com/tokencard/ethertest"
)

var _ = Describe("tokenWhitelistable", func() {

	It("Should return the stablecoin's contract address", func() {
		sa, err := TokenWhitelist.Stablecoin(nil)
		Expect(err).ToNot(HaveOccurred())
		Expect(sa).To(Equal(StablecoinAddress))
	})

	It("Should return the orcale ENS-registered node", func() {
		sa, err := TokenWhitelist.OracleNode(nil)
		Expect(err).ToNot(HaveOccurred())
		Expect(sa).To(Equal([32]byte(OracleName)))
	})

	It("Should return the tokenWhitelist ENS-registered node", func() {
		sa, err := TokenWhitelistableExporter.TokenWhitelistNode(nil)
		Expect(err).ToNot(HaveOccurred())
		Expect(sa).To(Equal([32]byte(TokenWhitelistName)))
	})

	When("DAI is the stablecoin used", func() {

		BeforeEach(func() {
			tx, err := TokenWhitelist.AddTokens(
				ControllerAdmin.TransactOpts(),
				[]common.Address{StablecoinAddress},
				StringsToByte32("DAI"),
				[]*big.Int{DecimalsToMagnitude(big.NewInt(18))},
				[]bool{true},
				[]bool{true},
				big.NewInt(20180913153211),
			)
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeTrue())
		})

		It("Should update the token mapping", func() {
			symbol, magnitude, rate, available, loadable, redeemable, lastUpdate, err := TokenWhitelistableExporter.GetStablecoinInfo(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(symbol).To(Equal("DAI"))
			Expect(magnitude.String()).To(Equal(DecimalsToMagnitude(big.NewInt(18)).String()))
			Expect(rate.String()).To(Equal("0"))
			Expect(available).To(BeTrue())
			Expect(loadable).To(BeTrue())
			Expect(redeemable).To(BeTrue())
			Expect(lastUpdate.String()).To(Equal(big.NewInt(20180913153211).String()))
		})
	})

	When("a token is added", func() {
		BeforeEach(func() {
			tx, err := TokenWhitelist.AddTokens(
				ControllerAdmin.TransactOpts(),
				[]common.Address{common.HexToAddress("0x1")},
				StringsToByte32("ETH"),
				[]*big.Int{DecimalsToMagnitude(big.NewInt(18))},
				[]bool{true},
				[]bool{true},
				big.NewInt(20180913153211),
			)
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeTrue())
		})

		It("Should update the token mapping", func() {
			symbol, magnitude, rate, available, loadable, redeemable, lastUpdate, err := TokenWhitelistableExporter.GetTokenInfo(nil, common.HexToAddress("0x1"))
			Expect(err).ToNot(HaveOccurred())
			Expect(symbol).To(Equal("ETH"))
			Expect(magnitude.String()).To(Equal(DecimalsToMagnitude(big.NewInt(18)).String()))
			Expect(rate.String()).To(Equal("0"))
			Expect(available).To(BeTrue())
			Expect(loadable).To(BeTrue())
			Expect(redeemable).To(BeTrue())
			Expect(lastUpdate.String()).To(Equal(big.NewInt(20180913153211).String()))
		})

		It("Should update the TokenAddressArray", func() {
			arr, err := TokenWhitelistableExporter.TokenAddressArray(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(arr).To(Equal([]common.Address{common.HexToAddress("0x1")}))
		})

		It("Should return true", func() {
			av, err := TokenWhitelistableExporter.IsTokenAvailable(nil, common.HexToAddress("0x1"))
			Expect(err).ToNot(HaveOccurred())
			Expect(av).To(BeTrue())
		})

		It("Should return true", func() {
			ld, err := TokenWhitelistableExporter.IsTokenLoadable(nil, common.HexToAddress("0x1"))
			Expect(err).ToNot(HaveOccurred())
			Expect(ld).To(BeTrue())
		})

		It("Should return true", func() {
			ld, err := TokenWhitelistableExporter.IsTokenRedeemable(nil, common.HexToAddress("0x1"))
			Expect(err).ToNot(HaveOccurred())
			Expect(ld).To(BeTrue())
		})

		When("updateRate is called and the datetime is the same", func() {
			It("Should fail", func() {
				tx, err := TokenWhitelistableExporter.UpdateTokenRate(
					ControllerAdmin.TransactOpts(ethertest.WithGasLimit(100000)),
					common.HexToAddress("0x1"),
					big.NewInt(666),
					big.NewInt(20180913153211),
				)
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeFalse())
			})
		})

		When("called by the controller admin directly", func() {

			BeforeEach(func() {
				tx, err := TokenWhitelist.UpdateTokenRate(
					ControllerAdmin.TransactOpts(),
					common.HexToAddress("0x1"),
					big.NewInt(555),
					big.NewInt(20180913153211),
				)
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})

			It("Should update the token rate", func() {
				symbol, magnitude, rate, available, loadable, redeemable, lastUpdate, err := TokenWhitelist.GetTokenInfo(nil, common.HexToAddress("0x1"))
				Expect(err).ToNot(HaveOccurred())
				Expect(symbol).To(Equal("ETH"))
				Expect(magnitude.String()).To(Equal(DecimalsToMagnitude(big.NewInt(18)).String()))
				Expect(rate.String()).To(Equal(big.NewInt(555).String()))
				Expect(available).To(BeTrue())
				Expect(loadable).To(BeTrue())
				Expect(redeemable).To(BeTrue())
				Expect(lastUpdate.String()).To(Equal(big.NewInt(20180913153211).String()))
			})

			It("Should emit a TokenRateUpdate event", func() {
				it, err := TokenWhitelist.FilterUpdatedTokenRate(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(it.Next()).To(BeTrue())
				evt := it.Event
				Expect(it.Next()).To(BeFalse())
				Expect(evt.Sender).To(Equal(ControllerAdmin.Address()))
				Expect(evt.Token).To(Equal(common.HexToAddress("0x1")))
				Expect(evt.Rate.String()).To(Equal(big.NewInt(555).String()))
			})
		})

	})

	When("3 tokens are added: 2 loadable, 1 redeemable", func() {
		BeforeEach(func() {
			var err error
			tokens := []common.Address{common.HexToAddress("0x0"), common.HexToAddress("0x1"), common.HexToAddress("0x2")}
			tx, err := TokenWhitelist.AddTokens(
				ControllerAdmin.TransactOpts(),
				tokens,
				StringsToByte32(
					"BNT",
					"TKN",
					"YAN",
				),
				[]*big.Int{
					DecimalsToMagnitude(big.NewInt(18)),
					DecimalsToMagnitude(big.NewInt(8)),
					DecimalsToMagnitude(big.NewInt(18)),
				},
				[]bool{false, true, true},
				[]bool{false, false, true},
				big.NewInt(20180913153211),
			)
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeTrue())
		})

		It("Should return 1 redeemable Token", func() {
			arr, err := TokenWhitelistableExporter.RedeemableTokens(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(arr).To(Equal([]common.Address{common.HexToAddress("0x2")}))
		})

		It("Should return 1 redeemable Token", func() {
			arr, err := TokenWhitelist.RedeemableTokens(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(arr).To(Equal([]common.Address{common.HexToAddress("0x2")}))
		})

		It("Should update the TokenAddressArray", func() {
			arr, err := TokenWhitelistableExporter.TokenAddressArray(nil)
			tokens := []common.Address{common.HexToAddress("0x0"), common.HexToAddress("0x1"), common.HexToAddress("0x2")}
			Expect(err).ToNot(HaveOccurred())
			Expect(arr).To(Equal(tokens))
		})

		It("Should increase the redeemable counter by 1", func() {
			cnt, err := TokenWhitelist.RedeemableCounter(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(cnt.String()).To(Equal("1"))
		})
	})

})
