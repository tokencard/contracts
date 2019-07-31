package token_whitelist_test

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tokencard/contracts/v2/test/shared"
	"github.com/tokencard/ethertest"
)

var _ = Describe("updateTokenRate", func() {

	Context("When the token is already supported", func() {
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
		Context("When called by the controller admin", func() {
			var tx *types.Transaction

			BeforeEach(func() {
				var err error
				tx, err = TokenWhitelist.UpdateTokenRate(
					ControllerAdmin.TransactOpts(),
					common.HexToAddress("0x1"),
					big.NewInt(555),
					big.NewInt(20180913153211),
				)
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
			})

			It("Should succeed", func() {
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
		Context("When not called by the controller", func() {
			It("Should fail", func() {
				tx, err := TokenWhitelist.UpdateTokenRate(
					RandomAccount.TransactOpts(ethertest.WithGasLimit(100000)),
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
	})

	Context("When the token is not supported", func() {
		Context("When called by the controller admin", func() {
			It("Should fail", func() {
				tx, err := TokenWhitelist.UpdateTokenRate(
					ControllerAdmin.TransactOpts(ethertest.WithGasLimit(100000)),
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
		Context("When not called by the controller", func() {
			It("Should fail", func() {
				tx, err := TokenWhitelist.UpdateTokenRate(
					RandomAccount.TransactOpts(ethertest.WithGasLimit(100000)),
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
	})

})
