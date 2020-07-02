package wallet_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tokencard/contracts/v3/test/shared"
)

var _ = Describe("ownable", func() {

	When("the contract with transferable ownership has been deployed", func() {

		It("should have an owner", func() {
			o, err := Wallet.Owner(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(o).To(Equal(Owner.Address()))
		})

		It("should have an owner", func() {
			t, err := Wallet.IsTransferable(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(t).To(BeTrue())
		})

	})

})
