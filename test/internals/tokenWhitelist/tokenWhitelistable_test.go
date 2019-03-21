package tokenWhitelist_test

import (
	"math/big"
	"github.com/ethereum/go-ethereum/common"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tokencard/contracts/test/shared"
  "github.com/tokencard/ethertest"
)

var _ = Describe("tokenWhitelistable", func() {

	When("DAI is the stablecoin used", func() {

		BeforeEach(func() {
			tx, err := TokenWhitelist.AddTokens(
				Controller.TransactOpts(),
				[]common.Address{StablecoinAddress},
				StringsToByte32("DAI"),
				[]*big.Int{DecimalsToMagnitude(big.NewInt(18))},
				[]bool{true},
				big.NewInt(20180913153211),
			)
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeTrue())
		})

		It("Should update the token mapping", func() {
			symbol, magnitude, rate, available, loadable, lastUpdate, err := TokenWhitelistableExporter.GetStablecoinInfo(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(symbol).To(Equal("DAI"))
			Expect(magnitude.String()).To(Equal(DecimalsToMagnitude(big.NewInt(18)).String()))
			Expect(rate.String()).To(Equal("0"))
			Expect(available).To(BeTrue())
			Expect(loadable).To(BeTrue())
			Expect(lastUpdate.String()).To(Equal(big.NewInt(20180913153211).String()))
		})
	})

  Context("When a token is added", func() {
		BeforeEach(func() {
			tx, err := TokenWhitelist.AddTokens(
				Controller.TransactOpts(),
				[]common.Address{common.HexToAddress("0x1")},
				StringsToByte32("ETH"),
				[]*big.Int{DecimalsToMagnitude(big.NewInt(18))},
				[]bool{true},
				big.NewInt(20180913153211),
			)
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeTrue())
		})

    It("Should update the token mapping", func() {
      symbol, magnitude, rate, available, loadable, lastUpdate, err := TokenWhitelistableExporter.GetTokenInfo(nil, common.HexToAddress("0x1"))
      Expect(err).ToNot(HaveOccurred())
      Expect(symbol).To(Equal("ETH"))
      Expect(magnitude.String()).To(Equal(DecimalsToMagnitude(big.NewInt(18)).String()))
      Expect(rate.String()).To(Equal("0"))
      Expect(available).To(BeTrue())
      Expect(loadable).To(BeTrue())
      Expect(lastUpdate.String()).To(Equal(big.NewInt(20180913153211).String()))
    })


    It("Should update the TokenAddressArray", func() {
      arr, err := TokenWhitelistableExporter.GetTokenAddressArray(nil)
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


    Context("When updateRate is called", func() {
			It("Should fail", func() {
				tx, err := TokenWhitelistableExporter.UpdateTokenRate(
					Controller.TransactOpts(ethertest.WithGasLimit(100000)),
					common.HexToAddress("0x1"),
					big.NewInt(666),
					big.NewInt(20180913153211),
				)
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isGasExhausted(tx, 100000)).To(BeFalse())
				Expect(isSuccessful(tx)).To(BeFalse())
			})
		})

    Context("When called by the controller directly", func() {

			BeforeEach(func() {
				tx, err := TokenWhitelist.UpdateTokenRate(
					Controller.TransactOpts(),
					common.HexToAddress("0x1"),
					big.NewInt(555),
					big.NewInt(20180913153211),
				)
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
        Expect(isSuccessful(tx)).To(BeTrue())
			})

			It("Should update the token rate", func() {
				symbol, magnitude, rate, available, loadable, lastUpdate, err := TokenWhitelist.GetTokenInfo(nil, common.HexToAddress("0x1"))
				Expect(err).ToNot(HaveOccurred())
				Expect(symbol).To(Equal("ETH"))
				Expect(magnitude.String()).To(Equal(DecimalsToMagnitude(big.NewInt(18)).String()))
				Expect(rate.String()).To(Equal(big.NewInt(555).String()))
				Expect(available).To(BeTrue())
				Expect(loadable).To(BeTrue())
				Expect(lastUpdate.String()).To(Equal(big.NewInt(20180913153211).String()))
			})
			It("Should emit a TokenRateUpdate event", func() {
				it, err := TokenWhitelist.FilterUpdatedTokenRate(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(it.Next()).To(BeTrue())
				evt := it.Event
				Expect(it.Next()).To(BeFalse())
				Expect(evt.Sender).To(Equal(Controller.Address()))
				Expect(evt.Token).To(Equal(common.HexToAddress("0x1")))
				Expect(evt.Rate.String()).To(Equal(big.NewInt(555).String()))
			})
		})

  })

})
