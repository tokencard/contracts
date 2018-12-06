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
			It("Should return 1230", func() {
				res, err := ParseIntScientificExporter.ParseIntScientificDecimals(nil, "0.0123", big.NewInt(5))
				Expect(err).ToNot(HaveOccurred())
				Expect(res.String()).To(Equal("1230"))
			})
		})

		When("('0.0123e-1', 5) is passed", func() {
			It("Should return 123", func() {
				res, err := ParseIntScientificExporter.ParseIntScientificDecimals(nil, "0.0123e-1", big.NewInt(5))
				Expect(err).ToNot(HaveOccurred())
				Expect(res.String()).To(Equal("123"))
			})
		})

		When("('0.0123e-1', 6) is passed", func() {
			It("Should return 1230", func() {
				res, err := ParseIntScientificExporter.ParseIntScientificDecimals(nil, "0.0123e-1", big.NewInt(6))
				Expect(err).ToNot(HaveOccurred())
				Expect(res.String()).To(Equal("1230"))
			})
		})

		When("('0.0123e-3', 6) is passed", func() {
			It("Should return 12", func() {
				res, err := ParseIntScientificExporter.ParseIntScientificDecimals(nil, "0.0123e-3", big.NewInt(6))
				Expect(err).ToNot(HaveOccurred())
				Expect(res.String()).To(Equal("12"))
			})
		})

		When("('1.0123e-3', 2) is passed", func() {
			It("Should return 0", func() {
				res, err := ParseIntScientificExporter.ParseIntScientificDecimals(nil, "1.0123e-3", big.NewInt(2))
				Expect(err).ToNot(HaveOccurred())
				Expect(res.String()).To(Equal("0"))
			})
		})

		When("('12345e-10', 5) is passed", func() {
			It("Should return 0", func() {
				res, err := ParseIntScientificExporter.ParseIntScientificDecimals(nil, "12345e-10", big.NewInt(5))
				Expect(err).ToNot(HaveOccurred())
				Expect(res.String()).To(Equal("0"))
			})
		})

		When("'0.12345e+5' is passed", func() {
			It("Should return 12345", func() {
				res, err := ParseIntScientificExporter.ParseIntScientific(nil, "0.12345e+5")
				Expect(err).ToNot(HaveOccurred())
				Expect(res.String()).To(Equal("12345"))
			})
		})

		When("('0.00123e+3', 2) is passed", func() {
			It("Should return 123", func() {
				res, err := ParseIntScientificExporter.ParseIntScientificDecimals(nil, "0.00123e+3", big.NewInt(2))
				Expect(err).ToNot(HaveOccurred())
				Expect(res.String()).To(Equal("123"))
			})
		})

		When("('0.00123e+2', 3) is passed", func() {
			It("Should return 123", func() {
				res, err := ParseIntScientificExporter.ParseIntScientificDecimals(nil, "0.00123e+2", big.NewInt(3))
				Expect(err).ToNot(HaveOccurred())
				Expect(res.String()).To(Equal("123"))
			})
		})

		When("('123.00123e+3', 4) is passed", func() {
			It("Should return 1230012300", func() {
				res, err := ParseIntScientificExporter.ParseIntScientificDecimals(nil, "123.00123e+3", big.NewInt(4))
				Expect(err).ToNot(HaveOccurred())
				Expect(res.String()).To(Equal("1230012300"))
			})
		})

		// When("('1.01.23e-3', 2) is passed", func() {
		// 	FIt("Should revert", func() {
		// 		res, err := ParseIntScientificExporter.ParseIntScientificDecimals(nil, "1.01.23e-3", big.NewInt(2))
		// 		Expect(err).ToNot(HaveOccurred())
		// 		Expect(res.String()).To(Equal("0"))
		// 	})
		// })


		})

})
