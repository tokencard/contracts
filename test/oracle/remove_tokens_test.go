package oracle_test

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tokencard/ethertest"
)

var _ = Describe("removeTokens", func() {

	Context("When called by the controller", func() {
		Context("When removing a supported token", func() {
			var tx *types.Transaction
			base := big.NewInt(10)
			expDec := big.NewInt(1)
			for i := 0; i < 18; i++ {
				expDec.Mul(expDec, base)
			}
			BeforeEach(func() {
				var err error
				tx, err = oracle.AddTokens(controller.TransactOpts(), []common.Address{common.HexToAddress("0x0")}, stringsToByte32("ETH"), []*big.Int{expDec})
				Expect(err).ToNot(HaveOccurred())
				be.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})

			BeforeEach(func() {
				var err error
				tx, err = oracle.RemoveTokens(controller.TransactOpts(), []common.Address{common.HexToAddress("0x0")})
				Expect(err).ToNot(HaveOccurred())
				be.Commit()
			})

			It("Should succeed", func() {
				Expect(isSuccessful(tx)).To(BeTrue())
			})

			It("Should emit a TokenRemoval event", func() {
				it, err := oracle.FilterRemovedToken(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(it.Next()).To(BeTrue())
				evt := it.Event
				Expect(it.Next()).To(BeFalse())
				Expect(evt.Token).To(Equal(common.HexToAddress("0x0")))
			})

		})
		Context("When removing an unsupported token", func() {

			It("Should fail", func() {
				tx, err := oracle.RemoveTokens(controller.TransactOpts(WithGasLimit(100000)), []common.Address{common.HexToAddress("0x0")})
				Expect(err).ToNot(HaveOccurred())
				be.Commit()
				Expect(isSuccessful(tx)).To(BeFalse())
			})
		})
	})

	Context("When called by a random address", func() {
		It("Should fail", func() {
			tx, err := oracle.RemoveTokens(randomAccount.TransactOpts(WithGasLimit(300000)), []common.Address{common.HexToAddress("0x0")})
			Expect(err).ToNot(HaveOccurred())
			be.Commit()
			Expect(isGasExhausted(tx, 300000)).To(BeFalse())
			Expect(isSuccessful(tx)).To(BeFalse())

		})
	})

})
