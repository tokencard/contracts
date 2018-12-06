package externals_test

import (
	"errors"
	"math/big"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	// . "github.com/tokencard/contracts/test/shared"
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

		When("'123' is passed", func() {
			It("Should return 0", func() {
				res, err := ParseIntScientificExporter.ParseIntScientific(nil, "123")
				Expect(err).ToNot(HaveOccurred())
				Expect(res.String()).To(Equal("123"))
			})
		})

		When("'123.0123' is passed", func() {
			It("Should return 0", func() {
				res, err := ParseIntScientificExporter.ParseIntScientific(nil, "123")
				Expect(err).ToNot(HaveOccurred())
				Expect(res.String()).To(Equal("123"))
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

		When("'12345e+5' is passed", func() {
			It("Should return 12345", func() {
				res, err := ParseIntScientificExporter.ParseIntScientific(nil, "12345e+5")
				Expect(err).ToNot(HaveOccurred())
				Expect(res.String()).To(Equal("1234500000"))
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

		When("there is an extra '.'", func() {
			It("Should revert", func() {
				_, err := ParseIntScientificExporter.ParseIntScientificDecimals(nil, "1.01.23e-3", big.NewInt(2))
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError(errors.New("abi: improperly formatted output")))
				// Expect(TestRig.LastExecuted()).To(MatchRegexp(`.*monthToNumber.*`))
			})
		})

		When("there is an extra 'e'", func() {
			It("Should revert", func() {
				_, err := ParseIntScientificExporter.ParseIntScientificDecimals(nil, "1.0123e-3e", big.NewInt(2))
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError(errors.New("abi: improperly formatted output")))
				// Expect(TestRig.LastExecuted()).To(MatchRegexp(`.*monthToNumber.*`))
			})
		})

		When("there is a '+' before 'e-'", func() {
			It("Should revert", func() {
				_, err := ParseIntScientificExporter.ParseIntScientificDecimals(nil, "1.01+23e-3", big.NewInt(2))
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError(errors.New("abi: improperly formatted output")))
				// Expect(TestRig.LastExecuted()).To(MatchRegexp(`.*monthToNumber.*`))
			})
		})

		When("there is a '+' before 'e+'", func() {
			It("Should revert", func() {
				_, err := ParseIntScientificExporter.ParseIntScientificDecimals(nil, "1.01+23e+3", big.NewInt(2))
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError(errors.New("abi: improperly formatted output")))
				// Expect(TestRig.LastExecuted()).To(MatchRegexp(`.*monthToNumber.*`))
			})
		})

		When("there is a '-' before 'e-'", func() {
			It("Should revert", func() {
				_, err := ParseIntScientificExporter.ParseIntScientificDecimals(nil, "1.01-23e-3", big.NewInt(2))
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError(errors.New("abi: improperly formatted output")))
				// Expect(TestRig.LastExecuted()).To(MatchRegexp(`.*monthToNumber.*`))
			})
		})

		When("there is a '-' before 'e+'", func() {
			It("Should revert", func() {
				_, err := ParseIntScientificExporter.ParseIntScientificDecimals(nil, "1.01-23e+3", big.NewInt(2))
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError(errors.New("abi: improperly formatted output")))
				// Expect(TestRig.LastExecuted()).To(MatchRegexp(`.*monthToNumber.*`))
			})
		})

		When("there is an extra '-' after '-'", func() {
			It("Should revert", func() {
				_, err := ParseIntScientificExporter.ParseIntScientificDecimals(nil, "1.0123e--3", big.NewInt(2))
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError(errors.New("abi: improperly formatted output")))
				// Expect(TestRig.LastExecuted()).To(MatchRegexp(`.*monthToNumber.*`))
			})
		})

		When("there is an extra '-' after 'e'", func() {
			It("Should revert", func() {
				_, err := ParseIntScientificExporter.ParseIntScientificDecimals(nil, "1.0123e-3-", big.NewInt(2))
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError(errors.New("abi: improperly formatted output")))
				// Expect(TestRig.LastExecuted()).To(MatchRegexp(`.*monthToNumber.*`))
			})
		})

		When("there is an extra '+' after '-'", func() {
			It("Should revert", func() {
				_, err := ParseIntScientificExporter.ParseIntScientificDecimals(nil, "1.0123e-+3", big.NewInt(2))
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError(errors.New("abi: improperly formatted output")))
				// Expect(TestRig.LastExecuted()).To(MatchRegexp(`.*monthToNumber.*`))
			})
		})

		When("there is an extra '+' after 'e'", func() {
			It("Should revert", func() {
				_, err := ParseIntScientificExporter.ParseIntScientificDecimals(nil, "1.0123e-3+", big.NewInt(2))
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError(errors.New("abi: improperly formatted output")))
				// Expect(TestRig.LastExecuted()).To(MatchRegexp(`.*monthToNumber.*`))
			})
		})

		When("there is no sign after exp", func() {
			It("Should revert", func() {
				_, err := ParseIntScientificExporter.ParseIntScientificDecimals(nil, "1.0123e3", big.NewInt(2))
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError(errors.New("abi: improperly formatted output")))
				// Expect(TestRig.LastExecuted()).To(MatchRegexp(`.*monthToNumber.*`))
			})
		})

		When("the exponent is not specified", func() {
			It("Should revert", func() {
				_, err := ParseIntScientificExporter.ParseIntScientificDecimals(nil, "1.0123e-", big.NewInt(2))
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError(errors.New("abi: improperly formatted output")))
				// Expect(TestRig.LastExecuted()).To(MatchRegexp(`.*monthToNumber.*`))
			})
		})

		When("the exponent is not specified", func() {
			It("Should revert", func() {
				_, err := ParseIntScientificExporter.ParseIntScientificDecimals(nil, "1.0123e+", big.NewInt(2))
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError(errors.New("abi: improperly formatted output")))
				// Expect(TestRig.LastExecuted()).To(MatchRegexp(`.*monthToNumber.*`))
			})
		})

		When("there is an illegal character (non-digit,'e','-','+')", func() {
			It("Should revert", func() {
				_, err := ParseIntScientificExporter.ParseIntScientificDecimals(nil, "1.01t23e+", big.NewInt(2))
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError(errors.New("abi: improperly formatted output")))
				// Expect(TestRig.LastExecuted()).To(MatchRegexp(`.*monthToNumber.*`))
			})
		})

		})

})
