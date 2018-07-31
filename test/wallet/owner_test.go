package wallet_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("owner", func() {
	Context("When the contract has been deployed", func() {
		It("should have an owner", func() {
			o, err := w.Owner(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(o).To(Equal(owner.Address()))
		})
	})
})
