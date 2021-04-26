package wallet_test

import (
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tokencard/contracts/v3/test/shared"
)

var _ = Describe("convertToStablecoin", func() {

	When("the token is already supported", func() {
		BeforeEach(func() {
			tx, err := TokenWhitelist.AddTokens(
				ControllerAdmin.TransactOpts(),
				[]common.Address{common.HexToAddress("0xfe209bdE5CA32fa20E6728A005F26D651FFF5982")},
				StringsToByte32("TKN"),
				[]*big.Int{DecimalsToMagnitude(big.NewInt(8))},
				[]bool{true},
				[]bool{true},
				big.NewInt(20180913153211),
			)
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeTrue())
		})

		When("the exchange rate is 0", func() {
			It("Should revert", func() {
				_, err := WalletProxy.ConvertToStablecoin(nil, common.HexToAddress("0xfe209bdE5CA32fa20E6728A005F26D651FFF5982"), big.NewInt(100))
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("rate=0"))
			})
		})

		When("the stablecoin rate is 0.1", func() {

			BeforeEach(func() {
				tx, err := TokenWhitelist.UpdateTokenRate(
					ControllerAdmin.TransactOpts(),
					StablecoinAddress,
					big.NewInt(int64(0.1*math.Pow10(18))),
					big.NewInt(20180913153211),
				)
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})

			When("the token rate is 0.1", func() {

				BeforeEach(func() {
					tx, err := TokenWhitelist.UpdateTokenRate(
						ControllerAdmin.TransactOpts(),
						common.HexToAddress("0xfe209bdE5CA32fa20E6728A005F26D651FFF5982"),
						big.NewInt(int64(0.1*math.Pow10(18))),
						big.NewInt(20180913153211),
					)
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeTrue())
				})

				When("overflow occurs", func() {
					It("Should trigger assert() (empty output)", func() {
						_, err := WalletProxy.ConvertToStablecoin(nil, common.HexToAddress("0xfe209bdE5CA32fa20E6728A005F26D651FFF5982"), big.NewInt(-1))
						Expect(err.Error()).To(ContainSubstring("SafeMath: multiplication overflow"))
					})
				})

				When("overflow does not occur", func() {
					It("Should return 0.1(amount)*0.1(rate)/0.1(stablecoin rate)*10^18(in Eth)", func() {
						value, err := WalletProxy.ConvertToStablecoin(nil, common.HexToAddress("0xfe209bdE5CA32fa20E6728A005F26D651FFF5982"), big.NewInt(int64(0.1*math.Pow10(8))))
						Expect(err).ToNot(HaveOccurred())
						finalAmount := EthToWei(1)
						finalAmount.Div(finalAmount, big.NewInt(10)) //the final amount should be 0.1*0.1*10*1DAI => 1/10 EthToWei , DAI decimals = 18
						Expect(value.String()).To(Equal(finalAmount.String()))
					})
				})
			})

			When("the token rate is 5.09", func() { //MKR
				BeforeEach(func() {
					tx, err := TokenWhitelist.UpdateTokenRate(
						ControllerAdmin.TransactOpts(),
						common.HexToAddress("0xfe209bdE5CA32fa20E6728A005F26D651FFF5982"),
						big.NewInt(int64(5.09*math.Pow10(18))),
						big.NewInt(20180913153211),
					)
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeTrue())
				})

				It("Should return 0.1(amount)*5.09(rate)/0.1(stablecoin rate)*10^6(in Mwei)", func() {
					value, err := WalletProxy.ConvertToStablecoin(nil, common.HexToAddress("0xfe209bdE5CA32fa20E6728A005F26D651FFF5982"), big.NewInt(int64(0.1*math.Pow10(8))))
					Expect(err).ToNot(HaveOccurred())
					finalAmount := EthToWei(509)
					finalAmount.Div(finalAmount, big.NewInt(100)) //the final amount should be 0.1*5.09*10*1ETH => 5.09 ETH => 509/100
					Expect(value.String()).To(Equal(finalAmount.String()))
				})

			})

		})
	})

	Context("When the token is not available", func() {
		It("Should return 0", func() {
			value, err := WalletProxy.ConvertToStablecoin(nil, common.HexToAddress("0x1"), big.NewInt(100))
			Expect(err).ToNot(HaveOccurred())
			Expect(value.String()).To(Equal("0"))
		})
	})

	When("eth is converted", func() {
		BeforeEach(func() {
			tx, err := TokenWhitelist.AddTokens(
				ControllerAdmin.TransactOpts(),
				[]common.Address{common.HexToAddress("0x0")},
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

		//rate has to be 1, i.e. 1 eth = > 10^18 wei
		BeforeEach(func() {
			tx, err := TokenWhitelist.UpdateTokenRate(
				ControllerAdmin.TransactOpts(),
				common.HexToAddress("0x0"),
				big.NewInt(int64(math.Pow10(18))),
				big.NewInt(20180913153211),
			)
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeTrue())
		})

		When("the stablecoin rate is 0.001", func() {
			BeforeEach(func() {
				tx, err := TokenWhitelist.UpdateTokenRate(
					ControllerAdmin.TransactOpts(),
					StablecoinAddress,
					big.NewInt(int64(0.001*math.Pow10(18))),
					big.NewInt(20180913153211),
				)
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})

			It("Should return 0.1(amount)/0.001(stablecoin rate)*10^18(in Eth)", func() {
				value, err := WalletProxy.ConvertToStablecoin(nil, common.HexToAddress("0x0"), big.NewInt(int64(0.1*math.Pow10(18))))
				Expect(err).ToNot(HaveOccurred())
				finalAmount := EthToWei(1)
				finalAmount.Mul(finalAmount, big.NewInt(100)) //the final amount should be 0.1*1000*1ETH => 1*100 ETH
				Expect(value.String()).To(Equal(finalAmount.String()))
			})
		})

		When("the stablecoin rate is 0", func() {
			BeforeEach(func() {
				tx, err := TokenWhitelist.UpdateTokenRate(
					ControllerAdmin.TransactOpts(),
					StablecoinAddress,
					big.NewInt(0),
					big.NewInt(20180913153211),
				)
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})

			It("Should revert", func() {
				_, err := WalletProxy.ConvertToStablecoin(nil, common.HexToAddress("0x0"), big.NewInt(100))
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("stablecoin rate=0"))
			})

		})
	})

	When("stablecoin is converted", func() {

		When("the stablecoin rate is 0.007", func() {
			BeforeEach(func() {
				tx, err := TokenWhitelist.UpdateTokenRate(
					ControllerAdmin.TransactOpts(),
					StablecoinAddress,
					big.NewInt(int64(0.007*math.Pow10(18))),
					big.NewInt(20180913153211),
				)
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})

			It("Should return the same amount", func() {
				value, err := WalletProxy.ConvertToStablecoin(nil, StablecoinAddress, big.NewInt(1000))
				Expect(err).ToNot(HaveOccurred())
				Expect(value.String()).To(Equal("1000"))
			})

			It("Should return the same amount", func() {
				amount := new(big.Int).Mul(EthToWei(1), big.NewInt(int64(1232.76546365*math.Pow10(8))))
				value, err := WalletProxy.ConvertToStablecoin(nil, StablecoinAddress, amount)
				Expect(err).ToNot(HaveOccurred())
				Expect(value.String()).To(Equal(amount.String()))
			})

		})
	})

})
