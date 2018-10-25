package wallet_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tokencard/contracts/test/shared"
)

var _ = Describe("ERC165", func() {
	Context("When the contract has been deployed", func() {
		It("should support the ERC165 interface", func() {
			supported, err := Wallet.SupportsInterface(nil, [4]byte{0x01, 0xff, 0xc9, 0xa7})
			Expect(err).ToNot(HaveOccurred())
			Expect(supported).To(BeTrue())
		})
		It("should not support the 0xffffffff interface", func() {
			supported, err := Wallet.SupportsInterface(nil, [4]byte{0xff, 0xff, 0xff, 0xff})
			Expect(err).ToNot(HaveOccurred())
			Expect(supported).To(BeFalse())
		})
		It("should not support a random interface", func() {
			supported, err := Wallet.SupportsInterface(nil, [4]byte{0xab, 0xde, 0x12, 0x84})
			Expect(err).ToNot(HaveOccurred())
			Expect(supported).To(BeFalse())
		})

	})
})
