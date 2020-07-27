package gas_proxy_test

import (
	. "github.com/onsi/ginkgo"
	// . "github.com/onsi/gomega"
	// . "github.com/tokencard/contracts/v3/test/shared"
)

var _ = Describe("ExecuteTransaction", func() {
	When("a transaction that burns 1,000,000 gas is executed by the controller", func() {
		BeforeEach(func() {})

		It("should cost at least 500,000 less gas when called via the gas proxy", func() {

		})

	})
})
