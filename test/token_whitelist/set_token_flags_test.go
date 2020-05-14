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

var _ = Describe("setTokenLoadable/Redeemable", func() {

	When("the token is both loadable and redeemable", func() {
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

		It("The overal token counter should be equal to 1", func() {
			cnt, err := TokenWhitelist.TokenCounter(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(cnt.String()).To(Equal("1"))
		})

		It("The loadable counter should be equal to 1", func() {
			cnt, err := TokenWhitelist.LoadableCounter(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(cnt.String()).To(Equal("1"))
		})

		It("The redeemable counter should be equal to 1", func() {
			cnt, err := TokenWhitelist.RedeemableCounter(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(cnt.String()).To(Equal("1"))
		})

		When("trying to set it to the current value (loadable = true)", func() {
			It("Should fail", func() {
				tx, err := TokenWhitelist.SetTokenLoadable(ControllerAdmin.TransactOpts(ethertest.WithGasLimit(100000)), common.HexToAddress("0x1"), true)
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeFalse())
				returnData, _ := ethCall(tx)
				Expect(string(returnData[len(returnData)-64:])).To(ContainSubstring("loadable: no state change"))
			})
		})

		When("trying to set it to the current value (redeemable = true)", func() {
			It("Should fail", func() {
				tx, err := TokenWhitelist.SetTokenRedeemable(ControllerAdmin.TransactOpts(ethertest.WithGasLimit(100000)), common.HexToAddress("0x1"), true)
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeFalse())
				returnData, _ := ethCall(tx)
				Expect(string(returnData[len(returnData)-64:])).To(ContainSubstring("redeemable: no state change"))
			})
		})

		When("SetTokenLoadable is called by the controller admin", func() {
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

			It("Should decrease the loadable counter", func() {
				cnt, err := TokenWhitelist.LoadableCounter(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(cnt.String()).To(Equal("0"))
			})

			When("trying to set it to the current value (loadable = false)", func() {
				It("Should fail", func() {
					tx, err := TokenWhitelist.SetTokenLoadable(ControllerAdmin.TransactOpts(ethertest.WithGasLimit(100000)), common.HexToAddress("0x1"), false)
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeFalse())
					returnData, _ := ethCall(tx)
					Expect(string(returnData[len(returnData)-64:])).To(ContainSubstring("loadable: no state change"))
				})
			})

			When("trying to reset it", func() {
				It("Should succeed", func() {
					tx, err := TokenWhitelist.SetTokenLoadable(ControllerAdmin.TransactOpts(), common.HexToAddress("0x1"), true)
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeTrue())
				})
			})
		})

		When("SetTokenRedeemable is called by the controller admin", func() {
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

			It("Should decrease the redeemable counter", func() {
				cnt, err := TokenWhitelist.RedeemableCounter(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(cnt.String()).To(Equal("0"))
			})

			When("trying to set it to the current value (false)", func() {
				It("Should fail", func() {
					tx, err := TokenWhitelist.SetTokenRedeemable(ControllerAdmin.TransactOpts(ethertest.WithGasLimit(100000)), common.HexToAddress("0x1"), false)
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeFalse())
					returnData, _ := ethCall(tx)
					Expect(string(returnData[len(returnData)-64:])).To(ContainSubstring("redeemable: no state change"))
				})
			})

			When("trying to reset it", func() {
				It("Should succeed", func() {
					tx, err := TokenWhitelist.SetTokenRedeemable(ControllerAdmin.TransactOpts(), common.HexToAddress("0x1"), true)
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeTrue())
				})
			})
		})

		When("SetTokenLoadable is not called by the controller admin", func() {
			It("Should fail", func() {
				tx, err := TokenWhitelist.SetTokenLoadable(
					RandomAccount.TransactOpts(ethertest.WithGasLimit(100000)),
					common.HexToAddress("0x1"),
					true,
				)
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeFalse())
				returnData, _ := ethCall(tx)
				Expect(string(returnData[len(returnData)-64:])).To(ContainSubstring("sender is not an admin"))
			})
		})

		When("SetTokenRedeemable is not called by the controller admin", func() {
			It("Should fail", func() {
				tx, err := TokenWhitelist.SetTokenRedeemable(
					RandomAccount.TransactOpts(ethertest.WithGasLimit(100000)),
					common.HexToAddress("0x1"),
					true,
				)
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeFalse())
				returnData, _ := ethCall(tx)
				Expect(string(returnData[len(returnData)-64:])).To(ContainSubstring("sender is not an admin"))
			})
		})
	})

	When("the token is not supported", func() {

		When("SetTokenRedeemable is called by the controller admin", func() {
			It("Should fail", func() {
				tx, err := TokenWhitelist.SetTokenRedeemable(
					ControllerAdmin.TransactOpts(ethertest.WithGasLimit(100000)),
					common.HexToAddress("0x2"),
					false,
				)
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeFalse())
				returnData, _ := ethCall(tx)
				Expect(string(returnData[len(returnData)-64:])).To(ContainSubstring("redeemable: token not available"))
			})
		})

		When("SetTokenLoadable is called by the controller admin", func() {
			It("Should fail", func() {
				tx, err := TokenWhitelist.SetTokenLoadable(
					ControllerAdmin.TransactOpts(ethertest.WithGasLimit(100000)),
					common.HexToAddress("0x2"),
					false,
				)
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeFalse())
				returnData, _ := ethCall(tx)
				Expect(string(returnData[len(returnData)-64:])).To(ContainSubstring("loadable: token not available"))
			})
		})

	})

})
