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

var _ = Describe("removeTokens", func() {

	When("tokens are supported", func() {
		BeforeEach(func() {

			tokens := []common.Address{common.HexToAddress("0x1"), common.HexToAddress("0x2"), common.HexToAddress("0x3")}
			tx, err := TokenWhitelist.AddTokens(
				ControllerAdmin.TransactOpts(),
				tokens,
				StringsToByte32(
					"OMG",
					"EOS",
					"TKN",
				),
				[]*big.Int{
					DecimalsToMagnitude(big.NewInt(18)),
					DecimalsToMagnitude(big.NewInt(18)),
					DecimalsToMagnitude(big.NewInt(8)),
				},
				[]bool{true, true, true},
				[]bool{true, true, true},
				big.NewInt(20180913153211),
			)
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeTrue())
		})

		It("The overal token counter should be equal to 3", func() {
			cnt, err := TokenWhitelist.TokenCounter(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(cnt.String()).To(Equal("3"))
		})

		When("called by the controller admin", func() {

			When("removing a supported token", func() {

				BeforeEach(func() {
					tx, err := TokenWhitelist.RemoveTokens(ControllerAdmin.TransactOpts(), []common.Address{common.HexToAddress("0x2")})
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeTrue())
				})

				It("Should emit a TokenRemoval event", func() {
					it, err := TokenWhitelist.FilterRemovedToken(nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(it.Next()).To(BeTrue())
					evt := it.Event
					Expect(it.Next()).To(BeFalse())
					Expect(evt.Token).To(Equal(common.HexToAddress("0x2")))
				})

				It("Should update the tokens map", func() {
					symbol, magnitude, rate, available, loadable, redeemable, lastUpdate, err := TokenWhitelist.GetTokenInfo(nil, common.HexToAddress("0x2"))
					Expect(err).ToNot(HaveOccurred())
					Expect(symbol).To(Equal(""))
					Expect(magnitude.String()).To(Equal("0"))
					Expect(rate.String()).To(Equal(big.NewInt(0).String()))
					Expect(available).To(BeFalse())
					Expect(loadable).To(BeFalse())
					Expect(redeemable).To(BeFalse())
					Expect(lastUpdate.String()).To(Equal("0"))

					//the other tokens should remain unchanged
					symbol, magnitude, rate, available, loadable, redeemable, lastUpdate, err = TokenWhitelist.GetTokenInfo(nil, common.HexToAddress("0x1"))
					Expect(err).ToNot(HaveOccurred())
					Expect(symbol).To(Equal("OMG"))
					Expect(magnitude.String()).To(Equal(DecimalsToMagnitude(big.NewInt(18)).String()))
					Expect(rate.String()).To(Equal(big.NewInt(0).String()))
					Expect(available).To(BeTrue())
					Expect(loadable).To(BeTrue())
					Expect(redeemable).To(BeTrue())
					Expect(lastUpdate.String()).To(Equal(big.NewInt(20180913153211).String()))

					symbol, magnitude, rate, available, loadable, redeemable, lastUpdate, err = TokenWhitelist.GetTokenInfo(nil, common.HexToAddress("0x3"))
					Expect(err).ToNot(HaveOccurred())
					Expect(symbol).To(Equal("TKN"))
					Expect(magnitude.String()).To(Equal(DecimalsToMagnitude(big.NewInt(8)).String()))
					Expect(rate.String()).To(Equal(big.NewInt(0).String()))
					Expect(available).To(BeTrue())
					Expect(loadable).To(BeTrue())
					Expect(redeemable).To(BeTrue())
					Expect(lastUpdate.String()).To(Equal(big.NewInt(20180913153211).String()))
				})

				It("Should decrease the overal token counter by 1", func() {
					cnt, err := TokenWhitelist.TokenCounter(nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(cnt.String()).To(Equal("2"))
				})

				It("Should decrease the loadable counter by 1", func() {
					cnt, err := TokenWhitelist.LoadableCounter(nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(cnt.String()).To(Equal("2"))
				})

				It("Should decrease the redeemable counter by 1", func() {
					cnt, err := TokenWhitelist.RedeemableCounter(nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(cnt.String()).To(Equal("2"))
				})

				When("removing another supported token", func() {

					BeforeEach(func() {
						tx, err := TokenWhitelist.RemoveTokens(ControllerAdmin.TransactOpts(), []common.Address{common.HexToAddress("0x3")})
						Expect(err).ToNot(HaveOccurred())
						Backend.Commit()
						Expect(isSuccessful(tx)).To(BeTrue())
					})

					It("Should decrease the overal token counter by 1", func() {
						cnt, err := TokenWhitelist.TokenCounter(nil)
						Expect(err).ToNot(HaveOccurred())
						Expect(cnt.String()).To(Equal("1"))
					})

					It("Should decrease the loadable counter by 1", func() {
						cnt, err := TokenWhitelist.LoadableCounter(nil)
						Expect(err).ToNot(HaveOccurred())
						Expect(cnt.String()).To(Equal("1"))
					})

					It("Should decrease the redeemable counter by 1", func() {
						cnt, err := TokenWhitelist.RedeemableCounter(nil)
						Expect(err).ToNot(HaveOccurred())
						Expect(cnt.String()).To(Equal("1"))
					})
				})

				When("removing an unsupported token", func() {

					BeforeEach(func() {
						tx, err := TokenWhitelist.RemoveTokens(ControllerAdmin.TransactOpts(ethertest.WithGasLimit(100000)), []common.Address{common.HexToAddress("0x5")})
						Expect(err).ToNot(HaveOccurred())
						Backend.Commit()
						Expect(isSuccessful(tx)).To(BeFalse())
						returnData, _ := ethCall(tx)
						Expect(string(returnData[len(returnData)-64:])).To(ContainSubstring("token is not available"))
					})

					It("Should NOT decrease the redeemable counter", func() {
						cnt, err := TokenWhitelist.RedeemableCounter(nil)
						Expect(err).ToNot(HaveOccurred())
						Expect(cnt.String()).To(Equal("2"))
					})
				})
			})

			When("removing all supported tokens", func() {

				var tx *types.Transaction
				BeforeEach(func() {
					var err error
					tokens := []common.Address{common.HexToAddress("0x1"), common.HexToAddress("0x2"), common.HexToAddress("0x3")}
					tx, err = TokenWhitelist.RemoveTokens(ControllerAdmin.TransactOpts(), tokens)
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
				})

				It("Should succeed", func() {
					Expect(isSuccessful(tx)).To(BeTrue())
				})

				It("Should emit 3 TokenRemoval events", func() {
					it, err := TokenWhitelist.FilterRemovedToken(nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(it.Next()).To(BeTrue())
					evt := it.Event
					Expect(it.Next()).To(BeTrue())
					Expect(evt.Token).To(Equal(common.HexToAddress("0x1")))
					evt = it.Event
					Expect(it.Next()).To(BeTrue())
					Expect(evt.Token).To(Equal(common.HexToAddress("0x2")))
					evt = it.Event
					Expect(it.Next()).To(BeFalse())
					Expect(evt.Token).To(Equal(common.HexToAddress("0x3")))
				})

				It("Should decrease the overal token counter down to 0", func() {
					cnt, err := TokenWhitelist.TokenCounter(nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(cnt.String()).To(Equal("0"))
				})

				It("Should decrease the loadable counter down to 0", func() {
					cnt, err := TokenWhitelist.LoadableCounter(nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(cnt.String()).To(Equal("0"))
				})

				It("Should decrease the redeemable counter down to 0", func() {
					cnt, err := TokenWhitelist.RedeemableCounter(nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(cnt.String()).To(Equal("0"))
				})
			})

			When("removing all supported tokens but a duplicate is passed", func() {
				It("Should fail", func() {
					tokens := []common.Address{common.HexToAddress("0x3"), common.HexToAddress("0x2"), common.HexToAddress("0x1"), common.HexToAddress("0x2")}
					tx, err := TokenWhitelist.RemoveTokens(ControllerAdmin.TransactOpts(ethertest.WithGasLimit(300000)), tokens)
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeFalse())
					returnData, _ := ethCall(tx)
					Expect(string(returnData[len(returnData)-64:])).To(ContainSubstring("token is not available"))
				})
			})

			When("removing at least one unsupported token", func() {
				It("Should fail", func() {
					tokens := []common.Address{common.HexToAddress("0x1"), common.HexToAddress("0x0"), common.HexToAddress("0x2")}
					tx, err := TokenWhitelist.RemoveTokens(ControllerAdmin.TransactOpts(ethertest.WithGasLimit(300000)), tokens)
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeFalse())
					returnData, _ := ethCall(tx)
					Expect(string(returnData[len(returnData)-64:])).To(ContainSubstring("token is not available"))
				})
			})

			When("trying to remove #0 tokens (empty list)", func() {

				var tx *types.Transaction
				BeforeEach(func() {
					var err error
					tx, err = TokenWhitelist.RemoveTokens(ControllerAdmin.TransactOpts(), nil)
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
				})

				It("Should succeed", func() {
					Expect(isSuccessful(tx)).To(BeTrue())
				})

				It("Should emit no Removed Token event", func() {
					it, err := TokenWhitelist.FilterRemovedToken(nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(it.Next()).To(BeFalse())
				})

				It("Should leave the the overal token counter intact", func() {
					cnt, err := TokenWhitelist.TokenCounter(nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(cnt.String()).To(Equal("3"))
				})

				It("Should leave the the loadable counter intact", func() {
					cnt, err := TokenWhitelist.LoadableCounter(nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(cnt.String()).To(Equal("3"))
				})

				It("Should leave the redeemable counter intact", func() {
					cnt, err := TokenWhitelist.RedeemableCounter(nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(cnt.String()).To(Equal("3"))
				})
			})

		})

		When("called by a random address", func() {
			It("Should fail", func() {
				tx, err := TokenWhitelist.RemoveTokens(RandomAccount.TransactOpts(ethertest.WithGasLimit(300000)), []common.Address{common.HexToAddress("0x1")})
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeFalse())
				returnData, _ := ethCall(tx)
				Expect(string(returnData[len(returnData)-64:])).To(ContainSubstring("sender is not an admin"))
			})
		})
	}) //when tokens are supported

})
