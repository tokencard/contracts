package oracle_test

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tokencard/contracts/test/shared"
	"github.com/tokencard/ethertest"
)

var _ = Describe("updateTokenRate", func() {

	Context("When the token is already supported", func() {
		BeforeEach(func() {
			tx, err := Oracle.AddTokens(
				Controller.TransactOpts(),
				[]common.Address{common.HexToAddress("0x1")},
				StringsToByte32("ETH"),
				[]*big.Int{DecimalsToMagnitude(big.NewInt(18))},
				big.NewInt(20180913153211),
			)
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeTrue())
		})
		Context("When called by the controller", func() {
			var tx *types.Transaction

			BeforeEach(func() {
				var err error
				tx, err = Oracle.UpdateTokenRate(
					Controller.TransactOpts(),
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
				token, err := Oracle.Tokens(nil, common.HexToAddress("0x1"))
				Expect(err).ToNot(HaveOccurred())
				Expect(token.Exists).To(BeTrue())
				Expect(token.Rate).To(Equal(big.NewInt(555)))
			})
			It("Should emit a TokenRateUpdate event", func() {
				it, err := Oracle.FilterUpdatedTokenRate(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(it.Next()).To(BeTrue())
				evt := it.Event
				Expect(it.Next()).To(BeFalse())
				Expect(evt.Token).To(Equal(common.HexToAddress("0x1")))
				// Expect(evt.Rate).To(Equal(big.NewInt(666)))
				Expect(evt.Rate.String()).To(Equal(big.NewInt(555).String()))
			})
		})
		Context("When not called by the controller", func() {
			It("Should fail", func() {
				tx, err := Oracle.UpdateTokenRate(
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
		Context("When called by the controller", func() {
			It("Should fail", func() {
				tx, err := Oracle.UpdateTokenRate(
					Controller.TransactOpts(ethertest.WithGasLimit(100000)),
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
				tx, err := Oracle.UpdateTokenRate(
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
