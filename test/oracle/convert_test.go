package oracle_test

import (
	"errors"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("convert", func() {
	Context("When the token is already supported", func() {
		BeforeEach(func() {
			tx, err := oracle.AddTokens(
				controller.TransactOpts(),
				[]common.Address{common.HexToAddress("0xfe209bdE5CA32fa20E6728A005F26D651FFF5982")},
				stringsToByte32("TKN"),
				[]*big.Int{calculateMagnitude(big.NewInt(8))},
				big.NewInt(20180913153211),
			)
			Expect(err).ToNot(HaveOccurred())
			be.Commit()
			Expect(isSuccessful(tx)).To(BeTrue())
		})
		Context("When exchange rate is 0", func() {
			It("Should fail", func() {
				_, err := oracle.Convert(nil, common.HexToAddress("0xfe209bdE5CA32fa20E6728A005F26D651FFF5982"), big.NewInt(100))
				Expect(err).To(HaveOccurred())
			})
		})
		Context("When exchange rate is NOT 0", func() {
			BeforeEach(func() {
				tx, err := oracle.UpdateTokenRate(
					controller.TransactOpts(),
					common.HexToAddress("0xfe209bdE5CA32fa20E6728A005F26D651FFF5982"),
					big.NewInt(int64(0.001633*math.Pow10(18))),
					big.NewInt(20180913153211),
				)
				Expect(err).ToNot(HaveOccurred())
				be.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})
			Context("When overflow occurs", func() {
				It("Should trigger assert() (empty output)", func() {
					_, err := oracle.Convert(nil, common.HexToAddress("0xfe209bdE5CA32fa20E6728A005F26D651FFF5982"), big.NewInt(-1))
					Expect(err).To(MatchError(errors.New("abi: unmarshalling empty output")))
				})
			})
			Context("When overflow does not occur", func() {
				It("Should return 0.01(amount)*0.001633(rate)*10^18(in wei)", func() {
					value, err := oracle.Convert(nil, common.HexToAddress("0xfe209bdE5CA32fa20E6728A005F26D651FFF5982"), big.NewInt(int64(0.01*math.Pow10(8))))
					Expect(err).ToNot(HaveOccurred())
					Expect(value.String()).To(Equal("16330000000000"))
				})
			})
		})
	})

	Context("When the token is not supported", func() {
		//the subsequent BeforeEach ensure that the the token fields are initialized but it is not supported longer
		BeforeEach(func() {
			tx, err := oracle.AddTokens(
				controller.TransactOpts(),
				[]common.Address{common.HexToAddress("0x1")},
				stringsToByte32("ETH"),
				[]*big.Int{calculateMagnitude(big.NewInt(18))},
				big.NewInt(20180913153211),
			)
			Expect(err).ToNot(HaveOccurred())
			be.Commit()
			Expect(isSuccessful(tx)).To(BeTrue())
		})
		BeforeEach(func() {
			tx, err := oracle.UpdateTokenRate(
				controller.TransactOpts(),
				common.HexToAddress("0x1"),
				big.NewInt(100),
				big.NewInt(20180913153211),
			)
			Expect(err).ToNot(HaveOccurred())
			be.Commit()
			Expect(isSuccessful(tx)).To(BeTrue())
		})
		BeforeEach(func() {
			tx, err := oracle.RemoveTokens(controller.TransactOpts(), []common.Address{common.HexToAddress("0x1")})
			Expect(err).ToNot(HaveOccurred())
			be.Commit()
			Expect(isSuccessful(tx)).To(BeTrue())
		})
		It("Should trigger require() (empty output)", func() {
			_, err := oracle.Convert(nil, common.HexToAddress("0x1"), big.NewInt(100))
			Expect(err).To(MatchError(errors.New("abi: improperly formatted output")))
		})
	})

})
