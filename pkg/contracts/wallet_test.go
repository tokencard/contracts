package contracts_test

import (
	"context"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Whitelist methods", func() {
	Context("When the InitializedWhitelist method is called", func() {
		It("should return false", func() {
			initialized, err := wallet.InitializedWhitelist(context.Background(), nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(initialized).To(BeFalse())
		})
	})


})
