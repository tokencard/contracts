package externals_test

import (
	// "errors"
	"math/big"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ParseIntScientific", func() {

	Context("When a valid encoded string is passed", func() {

		When("'0.0123' is passed", func() {
			It("Should return 0", func() {
				res, err := ParseIntScientificExporter.ParseIntScientific(nil, "0.0123")
				Expect(err).ToNot(HaveOccurred())
				Expect(res.String()).To(Equal("0"))
			})
		})

		When("('0.0123', 5) is passed", func() {
			It("Should return 2130", func() {
				res, err := ParseIntScientificExporter.ParseIntScientificDecimals(nil, "0.0123", big.NewInt(5))
				Expect(err).ToNot(HaveOccurred())
				Expect(res.String()).To(Equal("1230"))
			})
		})

		When("('0.0123e-1', 5) is passed", func() {
			It("Should return 213", func() {
				res, err := ParseIntScientificExporter.ParseIntScientificDecimals(nil, "0.0123e-1", big.NewInt(5))
				Expect(err).ToNot(HaveOccurred())
				Expect(res.String()).To(Equal("123"))
			})
		})

		})

})
