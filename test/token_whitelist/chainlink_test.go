package token_whitelist_test

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/tokencard/contracts/v3/pkg/bindings/mocks"
	. "github.com/tokencard/contracts/v3/test/shared"
	"github.com/tokencard/ethertest"
)

var _ = Describe("Chainlink", func() {

	var ChainlinkMock *mocks.ChainlinkMock
	var ChainlinkMockAddress common.Address

	When("not called by a controller admin", func() {
		It("Should fail", func() {
			tx, err := TokenWhitelist.UpdateChainlinkFeeds(
				Controller.TransactOpts(ethertest.WithGasLimit(100000)),
				[]common.Address{common.HexToAddress("0x1")},
				[]common.Address{common.HexToAddress("0x0")},
			)
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeFalse())
			returnData, _ := ethCall(tx)
			Expect(string(returnData[len(returnData)-64:])).To(ContainSubstring("sender is not an admin"))
		})
	})

	When("the token is not supported", func() {
		It("Should fail", func() {
			tx, err := TokenWhitelist.UpdateChainlinkFeeds(
				ControllerAdmin.TransactOpts(ethertest.WithGasLimit(100000)),
				[]common.Address{common.HexToAddress("0x1")},
				[]common.Address{common.HexToAddress("0x0")},
			)
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeFalse())
			returnData, _ := ethCall(tx)
			Expect(string(returnData[len(returnData)-64:])).To(ContainSubstring("Token not available"))
		})
	})

	When("parameter lengths don't match", func() {
		It("Should fail", func() {
			tx, err := TokenWhitelist.UpdateChainlinkFeeds(
				ControllerAdmin.TransactOpts(ethertest.WithGasLimit(100000)),
				[]common.Address{common.HexToAddress("0x1"), common.HexToAddress("0x2")},
				[]common.Address{common.HexToAddress("0x0")},
			)
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeFalse())
			returnData, _ := ethCall(tx)
			Expect(string(returnData[len(returnData)-64:])).To(ContainSubstring("Parameter length mismatch"))
		})
	})

	When("the tokens are already supported and the rates are updated", func() {
		BeforeEach(func() {
			tx, err := TokenWhitelist.AddTokens(
				ControllerAdmin.TransactOpts(),
				[]common.Address{common.HexToAddress("0x1"), StablecoinAddress},
				StringsToByte32("MKR", "USDC"),
				[]*big.Int{DecimalsToMagnitude(big.NewInt(18)), DecimalsToMagnitude(big.NewInt(18))},
				[]bool{true, true},
				[]bool{true, true},
				big.NewInt(20180913153211),
			)
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeTrue())
			tx, err = TokenWhitelist.UpdateTokenRate(
				ControllerAdmin.TransactOpts(),
				common.HexToAddress("0x1"),
				big.NewInt(555),
				big.NewInt(20180913153211),
			)
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeTrue())
		})

		It("Should return the token rate set in the whitelist", func() {
			_, _, rate, _, _, _, _, err := TokenWhitelist.GetTokenInfo(nil, common.HexToAddress("0x1"))
			Expect(err).ToNot(HaveOccurred())
			Expect(rate.String()).To(Equal("555"))
			_, _, rate, _, _, _, _, err = TokenWhitelist.GetStablecoinInfo(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(rate.String()).To(Equal("0"))
		})

		When("a chainlink ref contract is set", func() {
			var tx *types.Transaction
			var err error
			BeforeEach(func() {
				ChainlinkMockAddress, tx, ChainlinkMock, err = mocks.DeployChainlinkMock(BankAccount.TransactOpts(), Backend)
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())

				tx, err = TokenWhitelist.UpdateChainlinkFeeds(
					ControllerAdmin.TransactOpts(),
					[]common.Address{common.HexToAddress("0x1"), StablecoinAddress},
					[]common.Address{ChainlinkMockAddress, ChainlinkMockAddress},
				)
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})

			It("Should emit 2 UpdatedChainlinkFeed events", func() {
				it, err := TokenWhitelist.FilterUpdatedChainlinkFeed(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(it.Next()).To(BeTrue())
				evt := it.Event
				Expect(it.Next()).To(BeTrue())
				Expect(evt.Token).To(Equal(common.HexToAddress("0x1")))
				Expect(evt.PreviousContract).To(Equal(common.HexToAddress("0x0")))
				Expect(evt.NewContract).To(Equal(ChainlinkMockAddress))
				evt = it.Event
				Expect(it.Next()).To(BeFalse())
				Expect(evt.Token).To(Equal(StablecoinAddress))
				Expect(evt.PreviousContract).To(Equal(common.HexToAddress("0x0")))
				Expect(evt.NewContract).To(Equal(ChainlinkMockAddress))
			})

			It("Should return 0", func() {
				_, _, rate, _, _, _, _, err := TokenWhitelist.GetTokenInfo(nil, common.HexToAddress("0x1"))
				Expect(err).ToNot(HaveOccurred())
				Expect(rate.String()).To(Equal("0"))
				_, _, rate, _, _, _, _, err = TokenWhitelist.GetStablecoinInfo(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(rate.String()).To(Equal("0"))
			})

			When("the chainlink rate positive", func() {
				BeforeEach(func() {
					tx, err := ChainlinkMock.SetLatestAnswer(Controller.TransactOpts(), big.NewInt(666))
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeTrue())
				})

				It("Should return the chainlink rate", func() {
					_, _, rate, _, _, _, _, err := TokenWhitelist.GetTokenInfo(nil, common.HexToAddress("0x1"))
					Expect(err).ToNot(HaveOccurred())
					Expect(rate.String()).To(Equal("666"))
					_, _, rate, _, _, _, _, err = TokenWhitelist.GetStablecoinInfo(nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(rate.String()).To(Equal("666"))
				})
			})

			When("the chainlink rate is negative", func() {
				BeforeEach(func() {
					tx, err := ChainlinkMock.SetLatestAnswer(Controller.TransactOpts(), big.NewInt(-1))
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeTrue())
				})

				It("Should return a negative rate rate", func() {
					rate, err := ChainlinkMock.LatestAnswer(nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(rate.String()).To(Equal("-1"))
				})

				It("Should return the chainlink rate", func() {
					_, _, rate, _, _, _, _, err := TokenWhitelist.GetTokenInfo(nil, common.HexToAddress("0x1"))
					Expect(err).ToNot(HaveOccurred())
					Expect(rate.String()).To(Equal("0"))
					_, _, rate, _, _, _, _, err = TokenWhitelist.GetStablecoinInfo(nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(rate.String()).To(Equal("0"))
				})
			})

			When("the chainlink ref contract is removed", func() {
				BeforeEach(func() {
					tx, err = TokenWhitelist.UpdateChainlinkFeeds(
						ControllerAdmin.TransactOpts(),
						[]common.Address{common.HexToAddress("0x1"), StablecoinAddress},
						[]common.Address{common.HexToAddress("0x0"), common.HexToAddress("0x0")},
					)
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeTrue())
				})

				It("Should emit 2 extra UpdatedChainlinkFeed event", func() {
					it, err := TokenWhitelist.FilterUpdatedChainlinkFeed(nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(it.Next()).To(BeTrue())
					Expect(it.Next()).To(BeTrue())
					Expect(it.Next()).To(BeTrue())
					evt := it.Event
					Expect(it.Next()).To(BeTrue())
					Expect(evt.Token).To(Equal(common.HexToAddress("0x1")))
					Expect(evt.PreviousContract).To(Equal(ChainlinkMockAddress))
					Expect(evt.NewContract).To(Equal(common.HexToAddress("0x0")))
					evt = it.Event
					Expect(it.Next()).To(BeFalse())
					Expect(evt.Token).To(Equal(StablecoinAddress))
					Expect(evt.PreviousContract).To(Equal(ChainlinkMockAddress))
					Expect(evt.NewContract).To(Equal(common.HexToAddress("0x0")))
				})

				It("Should return the rate stored in the whitelist", func() {
					_, _, rate, _, _, _, _, err := TokenWhitelist.GetTokenInfo(nil, common.HexToAddress("0x1"))
					Expect(err).ToNot(HaveOccurred())
					Expect(rate.String()).To(Equal("555"))
					_, _, rate, _, _, _, _, err = TokenWhitelist.GetStablecoinInfo(nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(rate.String()).To(Equal("0"))
				})
			})

		})

	})

})
