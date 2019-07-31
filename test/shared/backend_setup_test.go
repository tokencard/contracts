package shared_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tokencard/contracts/v2/test/shared"
)

func TestWalletSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Contract Suite")
}

var _ = Describe("backend setup", func() {
	When("when the backend is initialized", func() {
		It("should succeed", func() {
			err := InitializeBackend()
			Expect(err).ToNot(HaveOccurred())
		})
	})
})

var _ = AfterEach(func() {
	err := Backend.Close()
	Expect(err).ToNot(HaveOccurred())
})
