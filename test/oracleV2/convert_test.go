package oracle_v2_test

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
			tx, err := oracle.AddToken(controllerWallet.TransactOpts(), common.HexToAddress("0xfe209bdE5CA32fa20E6728A005F26D651FFF5982"), "TKN", 8)
			Expect(err).ToNot(HaveOccurred())
			be.Commit()
			Expect(isSuccessful(tx)).To(BeTrue())
		})
		Context("When exchange rate is 0", func() {
			It("Should trigger require() (empty output)", func() {
				_, err := oracle.Convert(nil, common.HexToAddress("0xfe209bdE5CA32fa20E6728A005F26D651FFF5982"), big.NewInt(100))
				Expect(err).To(MatchError(errors.New("abi: unmarshalling empty output")))
			})
		})
		Context("When exchange rate is NOT 0", func() {
			BeforeEach(func() {
				tx, err := oracle.UpdateRateManual(controllerWallet.TransactOpts(), common.HexToAddress("0xfe209bdE5CA32fa20E6728A005F26D651FFF5982"), big.NewInt(int64(0.001633*math.Pow10(18))))
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
				It("Should return 1 (100*100/10^4)", func() {
					value, err := oracle.Convert(nil, common.HexToAddress("0xfe209bdE5CA32fa20E6728A005F26D651FFF5982"), big.NewInt(int64(0.01*math.Pow10(8))))
					Expect(err).ToNot(HaveOccurred())
					Expect(value.String()).To(Equal(big.NewInt(int64(0.00001633 * math.Pow10(18))).String()))
				})
			})
		})
	})

	Context("When the token is not supported", func() {
		//the subsequent BeforeEach ensure that the the token fields are initialized but it is not supported longer
		BeforeEach(func() {
			tx, err := oracle.AddToken(controllerWallet.TransactOpts(), common.HexToAddress("0x1"), "ETH", 4)
			Expect(err).ToNot(HaveOccurred())
			be.Commit()
			Expect(isSuccessful(tx)).To(BeTrue())
		})
		BeforeEach(func() {
			tx, err := oracle.UpdateRateManual(controllerWallet.TransactOpts(), common.HexToAddress("0x1"), big.NewInt(100))
			Expect(err).ToNot(HaveOccurred())
			be.Commit()
			Expect(isSuccessful(tx)).To(BeTrue())
		})
		BeforeEach(func() {
			tx, err := oracle.RemoveToken(controllerWallet.TransactOpts(), common.HexToAddress("0x1"))
			Expect(err).ToNot(HaveOccurred())
			be.Commit()
			Expect(isSuccessful(tx)).To(BeTrue())
		})
		It("Should trigger require() (empty output)", func() {
			_, err := oracle.Convert(nil, common.HexToAddress("0x1"), big.NewInt(100))
			Expect(err).To(MatchError(errors.New("abi: unmarshalling empty output")))
		})
	})

})
