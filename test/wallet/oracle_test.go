package wallet_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("oracle", func() {
	Context("When the contract has been deployed", func() {
		It("should have an oracle", func() {
			o, err := w.Oracle(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(o).To(Equal(oracleAddress))
		})
	})
})
