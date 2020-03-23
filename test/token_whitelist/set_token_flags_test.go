package token_whitelist_test

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tokencard/contracts/v3/test/shared"
	"github.com/tokencard/ethertest"
)

var _ = Describe("setTokenLoadable", func() {

	Context("When the token is already loadable", func() {
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
				tx, err = TokenWhitelist.SetTokenLoadable(ControllerAdmin.TransactOpts(), common.HexToAddress("0x1"), false)
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
			})

			It("Should succeed", func() {
				Expect(isSuccessful(tx)).To(BeTrue())
			})

			It("Should update the token loadable", func() {
				symbol, magnitude, _, available, loadable, redeemable, lastUpdate, err := TokenWhitelist.GetTokenInfo(nil, common.HexToAddress("0x1"))
				Expect(err).ToNot(HaveOccurred())
				Expect(symbol).To(Equal("ETH"))
				Expect(magnitude.String()).To(Equal(DecimalsToMagnitude(big.NewInt(18)).String()))
				Expect(available).To(BeTrue())
				Expect(loadable).To(BeFalse())
				Expect(redeemable).To(BeTrue())
				Expect(lastUpdate.String()).To(Equal(big.NewInt(20180913153211).String()))
			})

			It("Should emit a UpdatedTokenLoadable event", func() {
				it, err := TokenWhitelist.FilterUpdatedTokenLoadable(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(it.Next()).To(BeTrue())
				evt := it.Event
				Expect(it.Next()).To(BeFalse())
				Expect(evt.Sender).To(Equal(ControllerAdmin.Address()))
				Expect(evt.Token).To(Equal(common.HexToAddress("0x1")))
				Expect(evt.Loadable).To(BeFalse())
			})
		})
		Context("When not called by the controller admin", func() {
			It("Should fail", func() {
				tx, err := TokenWhitelist.SetTokenLoadable(
					RandomAccount.TransactOpts(ethertest.WithGasLimit(100000)),
					common.HexToAddress("0x1"),
					true,
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
				tx, err := TokenWhitelist.SetTokenLoadable(
					ControllerAdmin.TransactOpts(ethertest.WithGasLimit(100000)),
					common.HexToAddress("0x2"),
					false,
				)
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isGasExhausted(tx, 100000)).To(BeFalse())
				Expect(isSuccessful(tx)).To(BeFalse())
			})
		})

		Context("When not called by the controller", func() {
			It("Should fail", func() {
				tx, err := TokenWhitelist.SetTokenLoadable(
					RandomAccount.TransactOpts(ethertest.WithGasLimit(100000)),
					common.HexToAddress("0x2"),
					false,
				)
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isGasExhausted(tx, 100000)).To(BeFalse())
				Expect(isSuccessful(tx)).To(BeFalse())
			})
		})
	})

})

var _ = Describe("setTokenRedeemable", func() {

	Context("When the token is already redeemable", func() {
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
				tx, err = TokenWhitelist.SetTokenRedeemable(ControllerAdmin.TransactOpts(), common.HexToAddress("0x1"), false)
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
			})

			It("Should succeed", func() {
				Expect(isSuccessful(tx)).To(BeTrue())
			})

			It("Should update the token redeemable", func() {
				symbol, magnitude, _, available, loadable, redeemable, lastUpdate, err := TokenWhitelist.GetTokenInfo(nil, common.HexToAddress("0x1"))
				Expect(err).ToNot(HaveOccurred())
				Expect(symbol).To(Equal("ETH"))
				Expect(magnitude.String()).To(Equal(DecimalsToMagnitude(big.NewInt(18)).String()))
				Expect(available).To(BeTrue())
				Expect(loadable).To(BeTrue())
				Expect(redeemable).To(BeFalse())
				Expect(lastUpdate.String()).To(Equal(big.NewInt(20180913153211).String()))
			})

			It("Should emit a UpdatedTokenRedeemable event", func() {
				it, err := TokenWhitelist.FilterUpdatedTokenRedeemable(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(it.Next()).To(BeTrue())
				evt := it.Event
				Expect(it.Next()).To(BeFalse())
				Expect(evt.Sender).To(Equal(ControllerAdmin.Address()))
				Expect(evt.Token).To(Equal(common.HexToAddress("0x1")))
				Expect(evt.Redeemable).To(BeFalse())
			})
		})
		Context("When not called by the controller", func() {
			It("Should fail", func() {
				tx, err := TokenWhitelist.SetTokenRedeemable(
					RandomAccount.TransactOpts(ethertest.WithGasLimit(100000)),
					common.HexToAddress("0x1"),
					true,
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
				tx, err := TokenWhitelist.SetTokenRedeemable(
					ControllerAdmin.TransactOpts(ethertest.WithGasLimit(100000)),
					common.HexToAddress("0x2"),
					false,
				)
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isGasExhausted(tx, 100000)).To(BeFalse())
				Expect(isSuccessful(tx)).To(BeFalse())
			})
		})

		Context("When not called by the controller", func() {
			It("Should fail", func() {
				tx, err := TokenWhitelist.SetTokenRedeemable(
					RandomAccount.TransactOpts(ethertest.WithGasLimit(100000)),
					common.HexToAddress("0x2"),
					false,
				)
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isGasExhausted(tx, 100000)).To(BeFalse())
				Expect(isSuccessful(tx)).To(BeFalse())
			})
		})
	})

})
