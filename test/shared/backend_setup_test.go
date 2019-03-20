package shared_test

import (
	"fmt"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tokencard/contracts/test/shared"
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

var _ = AfterEach(func() {
	td := CurrentGinkgoTestDescription()
	if td.Failed {
		fmt.Fprintf(GinkgoWriter, "\nLast Executed Smart Contract Line for %s:%d\n", td.FileName, td.LineNumber)
		fmt.Fprintln(GinkgoWriter, TestRig.LastExecuted())
	}
	err := Backend.Close()
	Expect(err).ToNot(HaveOccurred())
})
