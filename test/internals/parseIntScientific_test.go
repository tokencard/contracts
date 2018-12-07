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

		When("'123e3' is passed", func() {
			It("Should return 123000", func() {
				res, err := ParseIntScientificExporter.ParseIntScientific(nil, "123e3")
				Expect(err).ToNot(HaveOccurred())
				Expect(res.String()).To(Equal("123000"))
			})
		})

		When("'123E3' is passed", func() {
			It("Should return 123000", func() {
				res, err := ParseIntScientificExporter.ParseIntScientific(nil, "123E3")
				Expect(err).ToNot(HaveOccurred())
				Expect(res.String()).To(Equal("123000"))
			})
		})

		When("'123e-3' is passed", func() {
			It("Should return 0", func() {
				res, err := ParseIntScientificExporter.ParseIntScientific(nil, "123e-3")
				Expect(err).ToNot(HaveOccurred())
				Expect(res.String()).To(Equal("0"))
			})
		})

		When("'123E-3' is passed", func() {
			It("Should return 0", func() {
				res, err := ParseIntScientificExporter.ParseIntScientific(nil, "123E-3")
				Expect(err).ToNot(HaveOccurred())
				Expect(res.String()).To(Equal("0"))
			})
		})

		When("'123e-2' is passed", func() {
			It("Should return 1", func() {
				res, err := ParseIntScientificExporter.ParseIntScientific(nil, "123e-2")
				Expect(err).ToNot(HaveOccurred())
				Expect(res.String()).To(Equal("1"))
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

		When("('0.0123E-1', 6) is passed", func() {
			It("Should return 1230", func() {
				res, err := ParseIntScientificExporter.ParseIntScientificDecimals(nil, "0.0123E-1", big.NewInt(6))
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

		When("'0.12345E+5' is passed", func() {
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

		When("'0.12345e5' is passed", func() {
			It("Should return 12345", func() {
				res, err := ParseIntScientificExporter.ParseIntScientific(nil, "0.12345e5")
				Expect(err).ToNot(HaveOccurred())
				Expect(res.String()).To(Equal("12345"))
			})
		})

		When("'0.12345E5' is passed", func() {
			It("Should return 12345", func() {
				res, err := ParseIntScientificExporter.ParseIntScientific(nil, "0.12345E5")
				Expect(err).ToNot(HaveOccurred())
				Expect(res.String()).To(Equal("12345"))
			})
		})

		When("'12345e5' is passed", func() {
			It("Should return 12345", func() {
				res, err := ParseIntScientificExporter.ParseIntScientific(nil, "12345e5")
				Expect(err).ToNot(HaveOccurred())
				Expect(res.String()).To(Equal("1234500000"))
			})
		})

		When("'12345E5' is passed", func() {
			It("Should return 12345", func() {
				res, err := ParseIntScientificExporter.ParseIntScientific(nil, "12345E5")
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

		When("('0.00123e3', 2) is passed", func() {
			It("Should return 123", func() {
				res, err := ParseIntScientificExporter.ParseIntScientificDecimals(nil, "0.00123e3", big.NewInt(2))
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

		When("('0.00123E3', 2) is passed", func() {
			It("Should return 123", func() {
				res, err := ParseIntScientificExporter.ParseIntScientificDecimals(nil, "0.00123E3", big.NewInt(2))
				Expect(err).ToNot(HaveOccurred())
				Expect(res.String()).To(Equal("123"))
			})
		})

		When("('0.00123E+2', 3) is passed", func() {
			It("Should return 123", func() {
				res, err := ParseIntScientificExporter.ParseIntScientificDecimals(nil, "0.00123E+2", big.NewInt(3))
				Expect(err).ToNot(HaveOccurred())
				Expect(res.String()).To(Equal("123"))
			})
		})

		When("('0.00123e2', 3) is passed", func() {
			It("Should return 123", func() {
				res, err := ParseIntScientificExporter.ParseIntScientificDecimals(nil, "0.00123e2", big.NewInt(3))
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

		When("('123.00123e3', 4) is passed", func() {
			It("Should return 1230012300", func() {
				res, err := ParseIntScientificExporter.ParseIntScientificDecimals(nil, "123.00123e3", big.NewInt(4))
				Expect(err).ToNot(HaveOccurred())
				Expect(res.String()).To(Equal("1230012300"))
			})
		})

		When("('0.000001', 6) is passed", func() {
			It("Should return 1", func() {
				res, err := ParseIntScientificExporter.ParseIntScientificDecimals(nil, "0.000001", big.NewInt(6))
				Expect(err).ToNot(HaveOccurred())
				Expect(res.String()).To(Equal("1"))
			})
		})

		When("('1e-7', 7) is passed", func() {
			It("Should return 1", func() {
				res, err := ParseIntScientificExporter.ParseIntScientificDecimals(nil, "1e-7", big.NewInt(7))
				Expect(err).ToNot(HaveOccurred())
				Expect(res.String()).To(Equal("1"))
			})
		})

		When("('100000000000000000000', 0) is passed", func() {
			It("Should return 100000000000000000000", func() {
				res, err := ParseIntScientificExporter.ParseIntScientificDecimals(nil, "100000000000000000000", big.NewInt(0))
				Expect(err).ToNot(HaveOccurred())
				Expect(res.String()).To(Equal("100000000000000000000"))
			})
		})

		When("('1e+21', 0) is passed", func() {
			It("Should return 1000000000000000000000", func() {
				res, err := ParseIntScientificExporter.ParseIntScientificDecimals(nil, "1e+21", big.NewInt(0))
				Expect(err).ToNot(HaveOccurred())
				Expect(res.String()).To(Equal("1000000000000000000000"))
			})
		})

		When("('999999999999999900000', 0) is passed", func() {
			It("Should return 999999999999999900000", func() {
				res, err := ParseIntScientificExporter.ParseIntScientificDecimals(nil, "999999999999999900000", big.NewInt(0))
				Expect(err).ToNot(HaveOccurred())
				Expect(res.String()).To(Equal("999999999999999900000"))
			})
		})

		When("('9007199254740992', 0) is passed", func() {
			It("Should return 9007199254740992", func() {
				res, err := ParseIntScientificExporter.ParseIntScientificDecimals(nil, "9007199254740992", big.NewInt(0))
				Expect(err).ToNot(HaveOccurred())
				Expect(res.String()).To(Equal("9007199254740992"))
			})
		})

		When("('9007199254740993', 0) is passed", func() {
			It("Should return 9007199254740992", func() {
				res, err := ParseIntScientificExporter.ParseIntScientificDecimals(nil, "9007199254740993", big.NewInt(0))
				Expect(err).ToNot(HaveOccurred())
				Expect(res.String()).To(Equal("9007199254740993"))
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

		When("there is an extra '.'", func() {
			It("Should revert", func() {
				_, err := ParseIntScientificExporter.ParseIntScientificDecimals(nil, "1.0123e-3.", big.NewInt(2))
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError(errors.New("abi: improperly formatted output")))
				// Expect(TestRig.LastExecuted()).To(MatchRegexp(`.*monthToNumber.*`))
			})
		})

		When("there is an extra '.'", func() {
			It("Should revert", func() {
				_, err := ParseIntScientificExporter.ParseIntScientificDecimals(nil, "1.0123e.-3", big.NewInt(2))
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError(errors.New("abi: improperly formatted output")))
				// Expect(TestRig.LastExecuted()).To(MatchRegexp(`.*monthToNumber.*`))
			})
		})

		When("there is an extra '.'", func() {
			It("Should revert", func() {
				_, err := ParseIntScientificExporter.ParseIntScientificDecimals(nil, "1.0123e-.3", big.NewInt(2))
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError(errors.New("abi: improperly formatted output")))
				// Expect(TestRig.LastExecuted()).To(MatchRegexp(`.*monthToNumber.*`))
			})
		})

		When("there is an extra exponent symbol", func() {
			It("Should revert", func() {
				_, err := ParseIntScientificExporter.ParseIntScientificDecimals(nil, "1.0123e-3e", big.NewInt(2))
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError(errors.New("abi: improperly formatted output")))
				// Expect(TestRig.LastExecuted()).To(MatchRegexp(`.*monthToNumber.*`))
			})
		})

		When("there is an extra exponent symbol", func() {
			It("Should revert", func() {
				_, err := ParseIntScientificExporter.ParseIntScientificDecimals(nil, "1.01E23e-3", big.NewInt(2))
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError(errors.New("abi: improperly formatted output")))
				// Expect(TestRig.LastExecuted()).To(MatchRegexp(`.*monthToNumber.*`))
			})
		})

		When("there is an extra exponent symbol", func() {
			It("Should revert", func() {
				_, err := ParseIntScientificExporter.ParseIntScientificDecimals(nil, "1.012e3e-3", big.NewInt(2))
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError(errors.New("abi: improperly formatted output")))
				// Expect(TestRig.LastExecuted()).To(MatchRegexp(`.*monthToNumber.*`))
			})
		})

		When("there is an extra exponent symbol", func() {
			It("Should revert", func() {
				_, err := ParseIntScientificExporter.ParseIntScientificDecimals(nil, "1.0123E-3e", big.NewInt(2))
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError(errors.New("abi: improperly formatted output")))
				// Expect(TestRig.LastExecuted()).To(MatchRegexp(`.*monthToNumber.*`))
			})
		})

		When("there is an extra exponent symbol", func() {
			It("Should revert", func() {
				_, err := ParseIntScientificExporter.ParseIntScientificDecimals(nil, "1.01e23E-3", big.NewInt(2))
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError(errors.New("abi: improperly formatted output")))
				// Expect(TestRig.LastExecuted()).To(MatchRegexp(`.*monthToNumber.*`))
			})
		})

		When("there is an extra exponent symbol", func() {
			It("Should revert", func() {
				_, err := ParseIntScientificExporter.ParseIntScientificDecimals(nil, "1.0123ee-3", big.NewInt(2))
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError(errors.New("abi: improperly formatted output")))
				// Expect(TestRig.LastExecuted()).To(MatchRegexp(`.*monthToNumber.*`))
			})
		})

		When("there is an extra exponent symbol", func() {
			It("Should revert", func() {
				_, err := ParseIntScientificExporter.ParseIntScientificDecimals(nil, "1.0123eE-3", big.NewInt(2))
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError(errors.New("abi: improperly formatted output")))
				// Expect(TestRig.LastExecuted()).To(MatchRegexp(`.*monthToNumber.*`))
			})
		})

		When("there is an extra exponent symbol", func() {
			It("Should revert", func() {
				_, err := ParseIntScientificExporter.ParseIntScientificDecimals(nil, "1.0123Ee-3", big.NewInt(2))
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError(errors.New("abi: improperly formatted output")))
				// Expect(TestRig.LastExecuted()).To(MatchRegexp(`.*monthToNumber.*`))
			})
		})

		When("there is an extra exponent symbol", func() {
			It("Should revert", func() {
				_, err := ParseIntScientificExporter.ParseIntScientificDecimals(nil, "1.0123EE-3", big.NewInt(2))
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

		When("there is a '-' before 'E+'", func() {
			It("Should revert", func() {
				_, err := ParseIntScientificExporter.ParseIntScientificDecimals(nil, "1.01-23E+3", big.NewInt(2))
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError(errors.New("abi: improperly formatted output")))
				// Expect(TestRig.LastExecuted()).To(MatchRegexp(`.*monthToNumber.*`))
			})
		})

		When("there is a '-' before 'E'", func() {
			It("Should revert", func() {
				_, err := ParseIntScientificExporter.ParseIntScientificDecimals(nil, "1.01-23E3", big.NewInt(2))
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError(errors.New("abi: improperly formatted output")))
				// Expect(TestRig.LastExecuted()).To(MatchRegexp(`.*monthToNumber.*`))
			})
		})

		When("there is a '-' before 'e'", func() {
			It("Should revert", func() {
				_, err := ParseIntScientificExporter.ParseIntScientificDecimals(nil, "1.01-23e3", big.NewInt(2))
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

		When("the exponent is not specified", func() {
			It("Should revert", func() {
				_, err := ParseIntScientificExporter.ParseIntScientificDecimals(nil, "1.0123e", big.NewInt(2))
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError(errors.New("abi: unmarshalling empty output")))
				// Expect(TestRig.LastExecuted()).To(MatchRegexp(`.*monthToNumber.*`))
			})
		})

		When("the exponent is not specified", func() {
			It("Should revert", func() {
				_, err := ParseIntScientificExporter.ParseIntScientificDecimals(nil, "1.0123e-", big.NewInt(2))
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError(errors.New("abi: unmarshalling empty output")))
				// Expect(TestRig.LastExecuted()).To(MatchRegexp(`.*monthToNumber.*`))
			})
		})

		When("the exponent is not specified", func() {
			It("Should revert", func() {
				_, err := ParseIntScientificExporter.ParseIntScientificDecimals(nil, "1.0123e+", big.NewInt(2))
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError(errors.New("abi: unmarshalling empty output")))
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

		When("there is an illegal character (non-digit,'e','-','+')", func() {
			It("Should revert", func() {
				_, err := ParseIntScientificExporter.ParseIntScientificDecimals(nil, "1.0123e+P", big.NewInt(2))
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError(errors.New("abi: improperly formatted output")))
				// Expect(TestRig.LastExecuted()).To(MatchRegexp(`.*monthToNumber.*`))
			})
		})

	})

})
