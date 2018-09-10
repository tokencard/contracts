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
			to := randomWallet.TransactOpts()
			to.GasLimit = 300000
			tx, err := oracle.Callback(to, [32]byte{}, "", []byte{})
			Expect(err).ToNot(HaveOccurred())
			be.Commit()
			Expect(isGasExhausted(tx, to.GasLimit)).To(BeFalse())
			Expect(isSuccessful(tx)).To(BeFalse())

		})
	})

})
