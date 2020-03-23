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

var _ = Describe("addTokens", func() {

	Context("When called by the controller admin", func() {

		Context("When the added tokens are not already supported", func() {
			var tx *types.Transaction
			BeforeEach(func() {
				var err error
				tokens := []common.Address{common.HexToAddress("0x0"), common.HexToAddress("0x1")}
				tx, err = TokenWhitelist.AddTokens(
					ControllerAdmin.TransactOpts(),
					tokens,
					StringsToByte32(
						"BNT",
						"TKN",
					),
					[]*big.Int{
						DecimalsToMagnitude(big.NewInt(18)),
						DecimalsToMagnitude(big.NewInt(8)),
					},
					[]bool{false, true},
					[]bool{true, true},
					big.NewInt(20180913153211),
				)
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})

			It("Should update the tokens map", func() {
				symbol, magnitude, _, available, loadable, redeemable, lastUpdate, err := TokenWhitelist.GetTokenInfo(nil, common.HexToAddress("0x0"))
				Expect(err).ToNot(HaveOccurred())
				Expect(symbol).To(Equal("BNT"))
				Expect(magnitude).To(Equal(DecimalsToMagnitude(big.NewInt(18))))
				Expect(available).To(BeTrue())
				Expect(loadable).To(BeFalse())
				Expect(redeemable).To(BeTrue())
				Expect(lastUpdate.String()).To(Equal("20180913153211"))

				symbol, magnitude, _, available, loadable, redeemable, lastUpdate, err = TokenWhitelist.GetTokenInfo(nil, common.HexToAddress("0x1"))
				Expect(err).ToNot(HaveOccurred())
				Expect(symbol).To(Equal("TKN"))
				Expect(magnitude.String()).To(Equal("100000000"))
				Expect(available).To(BeTrue())
				Expect(loadable).To(BeTrue())
				Expect(redeemable).To(BeTrue())
				Expect(lastUpdate.String()).To(Equal("20180913153211"))
			})

			It("Should emit 2 AddedToken events", func() {
				it, err := TokenWhitelist.FilterAddedToken(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(it.Next()).To(BeTrue())
				evt := it.Event
				Expect(it.Next()).To(BeTrue())
				Expect(evt.Token).To(Equal(common.HexToAddress("0x0")))
				Expect(evt.Symbol).To(Equal("BNT"))
				Expect(evt.Magnitude.String()).To(Equal(DecimalsToMagnitude(big.NewInt(18)).String()))
				evt = it.Event
				Expect(it.Next()).To(BeFalse())
				Expect(evt.Token).To(Equal(common.HexToAddress("0x1")))
				Expect(evt.Symbol).To(Equal("TKN"))
				Expect(evt.Magnitude.String()).To(Equal("100000000"))
			})

			It("Should increase the redeemable counter by 2", func() {
				cnt, err := TokenWhitelist.RedeemableCounter(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(cnt.String()).To(Equal("2"))
			})

			Context("When at least one of the added tokens is already supported", func() {
				It("Should fail", func() {
					var err error
					tokens := []common.Address{common.HexToAddress("0x0"), common.HexToAddress("0x2")}
					tx, err = TokenWhitelist.AddTokens(
						ControllerAdmin.TransactOpts(ethertest.WithGasLimit(200000)),
						tokens, StringsToByte32(
							"BNT",
							"OMG",
						),
						[]*big.Int{
							DecimalsToMagnitude(big.NewInt(18)),
							DecimalsToMagnitude(big.NewInt(18)),
						},
						[]bool{true, true},
						[]bool{true, true},
						big.NewInt(20180913153211),
					)
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isGasExhausted(tx, 200000)).To(BeFalse())
					Expect(isSuccessful(tx)).To(BeFalse())
				})
			})

			Context("When none of the added tokens is supported", func() {
				BeforeEach(func() {
					var err error
					tokens := []common.Address{common.HexToAddress("0x4"), common.HexToAddress("0x5")}
					tx, err = TokenWhitelist.AddTokens(
						ControllerAdmin.TransactOpts(ethertest.WithGasLimit(300000)),
						tokens,
						StringsToByte32(
							"MKR",
							"MTL",
						),
						[]*big.Int{
							DecimalsToMagnitude(big.NewInt(18)),
							DecimalsToMagnitude(big.NewInt(8)),
						},
						[]bool{true, true},
						[]bool{true, true},
						big.NewInt(20180913153211),
					)
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isGasExhausted(tx, 300000)).To(BeFalse())
					Expect(isSuccessful(tx)).To(BeTrue())
				})

				It("Should increase the redeemable counter by 2", func() {
					cnt, err := TokenWhitelist.RedeemableCounter(nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(cnt.String()).To(Equal("4"))
				})
			})

		})

		Context("When the parameters have different lengths", func() {
			It("Should fail", func() {
				tokens := []common.Address{common.HexToAddress("0x4"), common.HexToAddress("0x5")}
				tx, err := TokenWhitelist.AddTokens(
					ControllerAdmin.TransactOpts(ethertest.WithGasLimit(300000)),
					tokens,
					StringsToByte32(
						"BNT",
						"TKN",
						"ZRX",
					),
					[]*big.Int{
						DecimalsToMagnitude(big.NewInt(18)),
						DecimalsToMagnitude(big.NewInt(8)),
					},
					[]bool{true, true},
					[]bool{true, true},
					big.NewInt(20180913153211),
				)
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isGasExhausted(tx, 300000)).To(BeFalse())
				Expect(isSuccessful(tx)).To(BeFalse())
			})
		})

		Context("When trying to add #0 tokens (empty list)", func() {

			var tx *types.Transaction
			BeforeEach(func() {
				var err error
				tx, err = TokenWhitelist.AddTokens(ControllerAdmin.TransactOpts(), nil, nil, nil, nil, nil, big.NewInt(20180913153211))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
			})

			It("Should succeed", func() {
				Expect(isSuccessful(tx)).To(BeTrue())
			})

			It("Should emit no AddedToken event", func() {
				it, err := TokenWhitelist.FilterAddedToken(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(it.Next()).To(BeFalse())
			})

			It("Should not increase the redeemable counter", func() {
				cnt, err := TokenWhitelist.RedeemableCounter(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(cnt.String()).To(Equal("0"))
			})
		})

	})

	Context("When called by a random address", func() {
		It("Should fail", func() {
			tokens := []common.Address{common.HexToAddress("0x0"), common.HexToAddress("0x1")}
			tx, err := TokenWhitelist.AddTokens(
				RandomAccount.TransactOpts(ethertest.WithGasLimit(300000)),
				tokens,
				StringsToByte32(
					"BNT",
					"TKN",
				),
				[]*big.Int{
					DecimalsToMagnitude(big.NewInt(18)),
					DecimalsToMagnitude(big.NewInt(8)),
				},
				[]bool{true, true},
				[]bool{true, true},
				big.NewInt(20180913153211),
			)
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isGasExhausted(tx, 300000)).To(BeFalse())
			Expect(isSuccessful(tx)).To(BeFalse())
		})
	})

})
