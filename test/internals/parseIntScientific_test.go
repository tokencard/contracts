package internals_test

import (
	"errors"
	"math/big"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tokencard/contracts/test/shared"
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
			It("Should return 123", func() {
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
			It("Should return 123", func() {
				res, err := ParseIntScientificExporter.ParseIntScientific(nil, "123.0123")
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

		When("('1e+1', 0) is passed", func() {
			It("Should return 10", func() {
				res, err := ParseIntScientificExporter.ParseIntScientificDecimals(nil, "1e+1", big.NewInt(0))
				Expect(err).ToNot(HaveOccurred())
				Expect(res.String()).To(Equal("10"))
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
				res, err := ParseIntScientificExporter.ParseIntScientific(nil, "0.12345E+5")
				Expect(err).ToNot(HaveOccurred())
				Expect(res.String()).To(Equal("12345"))
			})
		})

		When("'12345e+5' is passed", func() {
			It("Should return 1234500000", func() {
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
			It("Should return 1234500000", func() {
				res, err := ParseIntScientificExporter.ParseIntScientific(nil, "12345e5")
				Expect(err).ToNot(HaveOccurred())
				Expect(res.String()).To(Equal("1234500000"))
			})
		})

		When("'12345E5' is passed", func() {
			It("Should return 1234500000", func() {
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

		When("('123.00123', 4) is passed", func() {
			It("Should return 1230012", func() {
				res, err := ParseIntScientificExporter.ParseIntScientificDecimals(nil, "123.001234", big.NewInt(4))
				Expect(err).ToNot(HaveOccurred())
				Expect(res.String()).To(Equal("1230012"))
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
				resStr := res.String()
				Expect(resStr).To(Equal("100000000000000000000"))
				Expect(len(resStr)).To(Equal(21))
			})
		})

		When("('1e+21', 0) is passed", func() {
			It("Should return 1000000000000000000000", func() {
				res, err := ParseIntScientificExporter.ParseIntScientificDecimals(nil, "1e+21", big.NewInt(0))
				Expect(err).ToNot(HaveOccurred())
				resStr := res.String()
				Expect(resStr).To(Equal("1000000000000000000000"))
				Expect(len(resStr)).To(Equal(22))
			})
		})

		When("('999999999999999900000', 0) is passed", func() {
			It("Should return 999999999999999900000", func() {
				res, err := ParseIntScientificExporter.ParseIntScientificDecimals(nil, "999999999999999900000", big.NewInt(0))
				Expect(err).ToNot(HaveOccurred())
				resStr := res.String()
				Expect(resStr).To(Equal("999999999999999900000"))
				Expect(len(resStr)).To(Equal(21))
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

		When("there is an extra '.' before 'e'", func() {
			It("Should revert", func() {
				_, err := ParseIntScientificExporter.ParseIntScientificDecimals(nil, "1.01.23e-3", big.NewInt(2))
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError(errors.New("abi: improperly formatted output")))
			})
		})

		When("there is an extra '.' after exponent", func() {
			It("Should revert", func() {
				_, err := ParseIntScientificExporter.ParseIntScientificDecimals(nil, "1.0123e-3.", big.NewInt(2))
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError(errors.New("abi: improperly formatted output")))
			})
		})

		When("there is an extra '.' immediately after 'e'", func() {
			It("Should revert", func() {
				_, err := ParseIntScientificExporter.ParseIntScientificDecimals(nil, "1.0123e.-3", big.NewInt(2))
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError(errors.New("abi: improperly formatted output")))
			})
		})

		When("there is an extra '.' after the exp sign", func() {
			It("Should revert", func() {
				_, err := ParseIntScientificExporter.ParseIntScientificDecimals(nil, "1.0123e-.3", big.NewInt(2))
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError(errors.New("abi: improperly formatted output")))
			})
		})

		When("there is a '.' immediately after the 'e'", func() {
			It("Should revert", func() {
				_, err := ParseIntScientificExporter.ParseIntScientificDecimals(nil, "1e.3", big.NewInt(0))
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError(errors.New("abi: improperly formatted output")))
			})
		})

		When("there is a '.' immediately after 'e' and before the exp sign", func() {
			It("Should revert", func() {
				_, err := ParseIntScientificExporter.ParseIntScientificDecimals(nil, "1e.+3", big.NewInt(0))
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError(errors.New("abi: improperly formatted output")))
			})
		})

		When("there is a '.' immediately after 'e' and before the exp sign", func() {
			It("Should revert", func() {
				_, err := ParseIntScientificExporter.ParseIntScientificDecimals(nil, "1e.-3", big.NewInt(2))
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError(errors.New("abi: improperly formatted output")))
			})
		})

		When("there is a '.' immediately after 'e' and before the exp sign", func() {
			It("Should revert", func() {
				_, err := ParseIntScientificExporter.ParseIntScientificDecimals(nil, "1e+.3", big.NewInt(0))
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError(errors.New("abi: improperly formatted output")))
			})
		})

		When("there is a '.' after the exp sign", func() {
			It("Should revert", func() {
				_, err := ParseIntScientificExporter.ParseIntScientificDecimals(nil, "1e-.3", big.NewInt(0))
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError(errors.New("abi: improperly formatted output")))
			})
		})

		When("there is a '.' after the exponent", func() {
			It("Should revert", func() {
				_, err := ParseIntScientificExporter.ParseIntScientificDecimals(nil, "1e3.", big.NewInt(0))
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError(errors.New("abi: improperly formatted output")))
			})
		})

		When("there is a '.' after a positive exponent", func() {
			It("Should revert", func() {
				_, err := ParseIntScientificExporter.ParseIntScientificDecimals(nil, "1e+3.", big.NewInt(0))
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError(errors.New("abi: improperly formatted output")))
			})
		})

		When("there is a '.' a negatiive exponent", func() {
			It("Should revert", func() {
				_, err := ParseIntScientificExporter.ParseIntScientificDecimals(nil, "1e-3.", big.NewInt(0))
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError(errors.New("abi: improperly formatted output")))
			})
		})

		When("there is an extra 'e' after 'e'", func() {
			It("Should revert", func() {
				_, err := ParseIntScientificExporter.ParseIntScientificDecimals(nil, "1.0123e-3e", big.NewInt(2))
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError(errors.New("abi: improperly formatted output")))
			})
		})

		When("there is an extra exponent after E", func() {
			It("Should revert", func() {
				_, err := ParseIntScientificExporter.ParseIntScientificDecimals(nil, "1.01E23e-3", big.NewInt(2))
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError(errors.New("abi: improperly formatted output")))

			})
		})

		When("there is an extra exponent after e", func() {
			It("Should revert", func() {
				_, err := ParseIntScientificExporter.ParseIntScientificDecimals(nil, "1.012e3e-3", big.NewInt(2))
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError(errors.New("abi: improperly formatted output")))
			})
		})

		When("there is an extra exponent symbol", func() {
			It("Should revert", func() {
				_, err := ParseIntScientificExporter.ParseIntScientificDecimals(nil, "1.0123E-3e", big.NewInt(2))
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError(errors.New("abi: improperly formatted output")))
			})
		})

		When("there is an extra exponent symbol", func() {
			It("Should revert", func() {
				_, err := ParseIntScientificExporter.ParseIntScientificDecimals(nil, "1.01e23E-3", big.NewInt(2))
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError(errors.New("abi: improperly formatted output")))
			})
		})

		When("there is an extra e adjacent to e", func() {
			It("Should revert", func() {
				_, err := ParseIntScientificExporter.ParseIntScientificDecimals(nil, "1.0123ee-3", big.NewInt(2))
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError(errors.New("abi: improperly formatted output")))

			})
		})

		When("there is an extra E adjacent to e", func() {
			It("Should revert", func() {
				_, err := ParseIntScientificExporter.ParseIntScientificDecimals(nil, "1.0123eE-3", big.NewInt(2))
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError(errors.New("abi: improperly formatted output")))

			})
		})

		When("there is an extra e adjacent to E", func() {
			It("Should revert", func() {
				_, err := ParseIntScientificExporter.ParseIntScientificDecimals(nil, "1.0123Ee-3", big.NewInt(2))
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError(errors.New("abi: improperly formatted output")))
			})
		})

		When("there is an an extra E adjacent to E", func() {
			It("Should revert", func() {
				_, err := ParseIntScientificExporter.ParseIntScientificDecimals(nil, "1.0123EE-3", big.NewInt(2))
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError(errors.New("abi: improperly formatted output")))

			})
		})

		When("there is a '+' before 'e-'", func() {
			It("Should revert", func() {
				_, err := ParseIntScientificExporter.ParseIntScientificDecimals(nil, "1.01+23e-3", big.NewInt(2))
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError(errors.New("abi: improperly formatted output")))

			})
		})

		When("there is a '+' before 'e+'", func() {
			It("Should revert", func() {
				_, err := ParseIntScientificExporter.ParseIntScientificDecimals(nil, "1.01+23e+3", big.NewInt(2))
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError(errors.New("abi: improperly formatted output")))

			})
		})

		When("there is a '-' before 'e-'", func() {
			It("Should revert", func() {
				_, err := ParseIntScientificExporter.ParseIntScientificDecimals(nil, "1.01-23e-3", big.NewInt(2))
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError(errors.New("abi: improperly formatted output")))

			})
		})

		When("there is a '-' before 'e+'", func() {
			It("Should revert", func() {
				_, err := ParseIntScientificExporter.ParseIntScientificDecimals(nil, "1.01-23e+3", big.NewInt(2))
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError(errors.New("abi: improperly formatted output")))

			})
		})

		When("there is a '-' before 'E+'", func() {
			It("Should revert", func() {
				_, err := ParseIntScientificExporter.ParseIntScientificDecimals(nil, "1.01-23E+3", big.NewInt(2))
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError(errors.New("abi: improperly formatted output")))

			})
		})

		When("there is a '-' before 'E'", func() {
			It("Should revert", func() {
				_, err := ParseIntScientificExporter.ParseIntScientificDecimals(nil, "1.01-23E3", big.NewInt(2))
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError(errors.New("abi: improperly formatted output")))

			})
		})

		When("there is a '-' before 'e'", func() {
			It("Should revert", func() {
				_, err := ParseIntScientificExporter.ParseIntScientificDecimals(nil, "1.01-23e3", big.NewInt(2))
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError(errors.New("abi: improperly formatted output")))

			})
		})

		When("there is an extra '-' immediately after '-'", func() {
			It("Should revert", func() {
				_, err := ParseIntScientificExporter.ParseIntScientificDecimals(nil, "1.0123e--3", big.NewInt(2))
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError(errors.New("abi: improperly formatted output")))

			})
		})

		When("there is an extra '-' after a negative exponent", func() {
			It("Should revert", func() {
				_, err := ParseIntScientificExporter.ParseIntScientificDecimals(nil, "1.0123e-3-", big.NewInt(2))
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError(errors.New("abi: improperly formatted output")))

			})
		})

		When("there is an extra '+' immediately after '-'", func() {
			It("Should revert", func() {
				_, err := ParseIntScientificExporter.ParseIntScientificDecimals(nil, "1.0123e-+3", big.NewInt(2))
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError(errors.New("abi: improperly formatted output")))

			})
		})

		When("there is an '+' after a negatie exponent", func() {
			It("Should revert", func() {
				_, err := ParseIntScientificExporter.ParseIntScientificDecimals(nil, "1.0123e-3+", big.NewInt(2))
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError(errors.New("abi: improperly formatted output")))

			})
		})

		When("the exponent is not specified", func() {
			It("Should revert", func() {
				_, err := ParseIntScientificExporter.ParseIntScientificDecimals(nil, "1.0123e", big.NewInt(2))
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError(errors.New("abi: unmarshalling empty output")))

			})
		})

		When("the negative exponent is not specified", func() {
			It("Should revert", func() {
				_, err := ParseIntScientificExporter.ParseIntScientificDecimals(nil, "1.0123e-", big.NewInt(2))
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError(errors.New("abi: unmarshalling empty output")))

			})
		})

		When("the positive exponent is not specified", func() {
			It("Should revert", func() {
				_, err := ParseIntScientificExporter.ParseIntScientificDecimals(nil, "1.0123e+", big.NewInt(2))
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError(errors.New("abi: unmarshalling empty output")))

			})
		})

		When("there is an illegal character 't' (non-digit,'e','-','+')", func() {
			It("Should revert", func() {
				_, err := ParseIntScientificExporter.ParseIntScientificDecimals(nil, "1.01t23e+", big.NewInt(2))
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError(errors.New("abi: improperly formatted output")))

			})
		})

		When("there is an illegal character '#' (non-digit,'e','-','+')", func() {
			It("Should revert", func() {
				_, err := ParseIntScientificExporter.ParseIntScientificDecimals(nil, "#1.01t23e+", big.NewInt(2))
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError(errors.New("abi: improperly formatted output")))

			})
		})

		When("there is an illegal character 'P' after the exponent sign(non-digit,'e','-','+')", func() {
			It("Should revert", func() {
				_, err := ParseIntScientificExporter.ParseIntScientificDecimals(nil, "1.0123e+P", big.NewInt(2))
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError(errors.New("abi: improperly formatted output")))

			})
		})

		When("there is an illegal character '&' after the exponent(non-digit,'e','-','+')", func() {
			It("Should revert", func() {
				_, err := ParseIntScientificExporter.ParseIntScientificDecimals(nil, "1.0123e+12&", big.NewInt(0))
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError(errors.New("abi: improperly formatted output")))

			})
		})

		When("there is no integral part", func() {
			It("Should revert", func() {
				_, err := ParseIntScientificExporter.ParseIntScientificDecimals(nil, "e1", big.NewInt(0))
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError(errors.New("abi: improperly formatted output")))
			})
		})

		When("there is no integral part before a positive eponent", func() {
			It("Should revert", func() {
				_, err := ParseIntScientificExporter.ParseIntScientificDecimals(nil, "e+1", big.NewInt(0))
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError(errors.New("abi: improperly formatted output")))
			})
		})

		When("there is no integral part before a negative eponent", func() {
			It("Should revert", func() {
				_, err := ParseIntScientificExporter.ParseIntScientificDecimals(nil, "e-1", big.NewInt(0))
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError(errors.New("abi: improperly formatted output")))
			})
		})

		When("there is no integral part before the decimal part", func() {
			It("Should revert", func() {
				_, err := ParseIntScientificExporter.ParseIntScientificDecimals(nil, ".1e-1", big.NewInt(0))
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError(errors.New("abi: improperly formatted output")))
			})
		})

		When("there is no integral part before the decimal part", func() {
			It("Should revert", func() {
				_, err := ParseIntScientificExporter.ParseIntScientificDecimals(nil, ".1e1", big.NewInt(2))
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError(errors.New("abi: improperly formatted output")))
			})
		})

		When("there is no integral part before the decimal part", func() {
			It("Should revert", func() {
				_, err := ParseIntScientificExporter.ParseIntScientificDecimals(nil, ".1234", big.NewInt(2))
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError(errors.New("abi: improperly formatted output")))
			})
		})

		When("the integral part is 2^256-1 (max_uint256_value)", func() {
			It("Should succeed", func() {
				//input = 2^256, 0 <= uint <=2^256-1
				res, err := ParseIntScientificExporter.ParseIntScientificDecimals(nil, "115792089237316195423570985008687907853269984665640564039457584007913129639935", big.NewInt(0))
				Expect(err).ToNot(HaveOccurred())
				Expect(res.String()).To(Equal("115792089237316195423570985008687907853269984665640564039457584007913129639935"))
			})
		})

		When("the integral part is 2^256 and an overflow occurs (max_uint256_value) = 2^256-1)", func() {
			It("Should revert", func() {
				//input 2^256, 0 <= uint <=2^256-1
				_, err := ParseIntScientificExporter.ParseIntScientificDecimals(nil, "115792089237316195423570985008687907853269984665640564039457584007913129639936", big.NewInt(0))
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError(errors.New("abi: unmarshalling empty output")))
			})
		})

		When("the exponent is 2^256 and an overflow occurs (max_uint256_value) = 2^256-1)", func() {
			It("Should revert", func() {
				//exp = 2^256, 0 <= uint <=2^256-1
				_, err := ParseIntScientificExporter.ParseIntScientificDecimals(nil, "1e115792089237316195423570985008687907853269984665640564039457584007913129639936", big.NewInt(0))
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError(errors.New("abi: unmarshalling empty output")))
			})
		})

		When("the decimal part is 2^256 and an overflow occurs (max_uint256_value) = 2^256-1)", func() {
			It("Should revert", func() {
				//dec = 2^256, 0 <= uint <=2^256-1
				_, err := ParseIntScientificExporter.ParseIntScientificDecimals(nil, "1.115792089237316195423570985008687907853269984665640564039457584007913129639936", big.NewInt(0))
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError(errors.New("abi: unmarshalling empty output")))
			})
		})

		When("an overflow occurs when adding the exponent and the magnitude (2^256)", func() {
			It("Should revert", func() {
				//exp = 2^255 - 1, _magnitudeMult = 1, 0 <= uint <=2^256-1
				_, err := ParseIntScientificExporter.ParseIntScientificDecimals(nil, "1e115792089237316195423570985008687907853269984665640564039457584007913129639935", big.NewInt(1))
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError(errors.New("abi: unmarshalling empty output")))
			})
		})

		When("an overflow occurs when adding the exponent and the magnitude (2^256)", func() {
			It("Should revert", func() {
				//exp = 2^255 - 2, _magnitudeMult = 2, 0 <= uint <=2^256-1
				_, err := ParseIntScientificExporter.ParseIntScientificDecimals(nil, "1e115792089237316195423570985008687907853269984665640564039457584007913129639934", big.NewInt(2))
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError(errors.New("abi: unmarshalling empty output")))
			})
		})

		When("an overflow occurs when the #decimals minted are 78 (> 77)", func() {
			It("Should revert", func() {
				//10^77 < 2^256 - 1 < 10^78, #decimals = 78
				_, err := ParseIntScientificExporter.ParseIntScientificDecimals(nil, "1.111111111111111111111111111111111111111111111111111111111111111111111111111111", big.NewInt(0))
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError(errors.New("abi: improperly formatted output")))
				// Expect(TestRig.LastExecuted()).To(MatchRegexp(`.*require\(second < 60, "second error"\);`))
			})
		})

		When("an overflow occurs when the integral part (10) is shifted 77 times", func() {
			It("Should revert", func() {
				//10^77 < 2^256 - 1 < 10^78
				_, err := ParseIntScientificExporter.ParseIntScientificDecimals(nil, "10", big.NewInt(77))
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError(errors.New("abi: unmarshalling empty output")))
			})
		})

		When("an overflow occurs when the integral part (1) is shifted 78 times", func() {
			It("Should revert", func() {
				//10^77 < 2^256 - 1 < 10^78
				_, err := ParseIntScientificExporter.ParseIntScientificDecimals(nil, "1", big.NewInt(78))
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError(errors.New("abi: improperly formatted output")))
			})
		})

		When("an overflow occurs when the integral part (10) is shifted 77 times", func() {
			It("Should revert", func() {
				//10^77 < 2^256 - 1 < 10^78
				_, err := ParseIntScientificExporter.ParseIntScientificDecimals(nil, "10.11111111111111111111111111111111111111111111111111111111111111111111111111111", big.NewInt(77))
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError(errors.New("abi: unmarshalling empty output")))
				Expect(TestRig.LastExecuted()).To(MatchRegexp(`.*require\(c / a == b\);`))
			})
		})

		When("an overflow occurs when the decimal part is added to the integral one", func() {
			It("Should revert (safeMath.add())", func() {
				//10^77 < 2^256 - 1 < 10^78
				_, err := ParseIntScientificExporter.ParseIntScientificDecimals(nil, "1.99999999999999999999999999999999999999999999999999999999999999999999999999999", big.NewInt(77))
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError(errors.New("abi: unmarshalling empty output")))
				Expect(TestRig.LastExecuted()).To(MatchRegexp(`.*require\(c >= a\);`))
			})
		})

		When("'0.00208' is passed to ParseIntScientificWei", func() {
			It("Should return 208", func() {
				res, err := ParseIntScientificExporter.ParseIntScientificWei(nil, "0.00208")
				Expect(err).ToNot(HaveOccurred())
				resStr := res.String()
				Expect(resStr).To(Equal("2080000000000000"))
				Expect(len(resStr)).To(Equal(16))
			})
		})

	})

})
