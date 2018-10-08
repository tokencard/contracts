package oracle_test

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("removeTokens", func() {

	Context("When called by the controller", func() {
		Context("When removing a supported token", func() {
			var tx *types.Transaction

			BeforeEach(func() {
				var err error
				tx, err = oracle.AddTokens(controllerWallet.TransactOpts(), []common.Address{common.HexToAddress("0x0")}, stringsToByte32("ETH"), []byte{18})
				Expect(err).ToNot(HaveOccurred())
				be.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})

			BeforeEach(func() {
				var err error
				tx, err = oracle.RemoveTokens(controllerWallet.TransactOpts(), []common.Address{common.HexToAddress("0x0")})
				Expect(err).ToNot(HaveOccurred())
				be.Commit()
			})

			It("Should succeed", func() {
				Expect(isSuccessful(tx)).To(BeTrue())
			})

			It("Should emit a TokenRemoval event", func() {
				it, err := oracle.FilterTokenRemoved(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(it.Next()).To(BeTrue())
				evt := it.Event
				Expect(it.Next()).To(BeFalse())
				Expect(evt.Token).To(Equal(common.HexToAddress("0x0")))
			})

		})
		Context("When removing an unsupported token", func() {

			It("Should fail", func() {
				tx, err := oracle.RemoveTokens(controllerWallet.TransactOptsWithGasLimit(100000), []common.Address{common.HexToAddress("0x0")})
				Expect(err).ToNot(HaveOccurred())
				be.Commit()
				Expect(isSuccessful(tx)).To(BeFalse())
			})
		})
	})

	Context("When called by a random address", func() {
		It("Should fail", func() {
			tx, err := oracle.RemoveTokens(randomWallet.TransactOptsWithGasLimit(300000), []common.Address{common.HexToAddress("0x0")})
			Expect(err).ToNot(HaveOccurred())
			be.Commit()
			Expect(isGasExhausted(tx, 300000)).To(BeFalse())
			Expect(isSuccessful(tx)).To(BeFalse())

		})
	})

})
