package oracle_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("addToken", func() {

	Context("When called by oraclize", func() {
		It("Should succeed", func() {

		})
	})

	Context("When called by a random address", func() {
		It("Should fail", func() {
			tx, err := oracle.Callback(randomWallet.TransactOptsWithGasLimit(300000), [32]byte{}, "", []byte{})
			Expect(err).ToNot(HaveOccurred())
			be.Commit()
			Expect(isGasExhausted(tx, 300000)).To(BeFalse())
			Expect(isSuccessful(tx)).To(BeFalse())

		})
	})

})
