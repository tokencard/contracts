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
				it, err := oracle.FilterTokenRemoval(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(it.Next()).To(BeTrue())
				evt := it.Event
				Expect(it.Next()).To(BeFalse())
				Expect(evt.Token).To(Equal(common.HexToAddress("0x0")))
			})

		})
		Context("When removing an unsupported token", func() {

			It("Should fail", func() {
				to := controllerWallet.TransactOpts()
				to.GasLimit = 100000
				tx, err := oracle.RemoveTokens(to, []common.Address{common.HexToAddress("0x0")})
				Expect(err).ToNot(HaveOccurred())
				be.Commit()
				Expect(isSuccessful(tx)).To(BeFalse())
			})
		})
	})

	Context("When called by a random address", func() {
		It("Should fail", func() {
			to := randomWallet.TransactOpts()
			to.GasLimit = 300000
			tx, err := oracle.RemoveTokens(to, []common.Address{common.HexToAddress("0x0")})
			Expect(err).ToNot(HaveOccurred())
			be.Commit()
			Expect(isGasExhausted(tx, to.GasLimit)).To(BeFalse())
			Expect(isSuccessful(tx)).To(BeFalse())

		})
	})

})
