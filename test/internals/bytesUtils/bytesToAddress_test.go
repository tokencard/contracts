package bytesUtils_test

import (
    "math/big"

    "github.com/ethereum/go-ethereum/common"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("BytesToAddress", func() {

		When("the length is less than 20 bytes", func() {
			It("Should revert", func() {
                bytes := common.Hex2Bytes("00000000000000000000000000000000000000")
                _, err := BytesUtilsExporter.BytesToAddress(nil, bytes, big.NewInt(0))
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("slicing out of range"))
			})
		})

        When("the length is 20 bytes but the start index is not zero", func() {
			It("Should revert", func() {
                bytes := common.Hex2Bytes("0000000000000000000000000000000000000000")
				_, err := BytesUtilsExporter.BytesToAddress(nil, bytes, big.NewInt(1))
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("slicing out of range"))
			})
		})

        When("the there is an overflow and start position wraps around in the second addition", func() {
			It("Should revert", func() {
                bytes := common.Hex2Bytes("000000000000000000000000000000000000000001")
				_, err := BytesUtilsExporter.BytesToAddress(nil, bytes, big.NewInt(-20))
                Expect(err).To(HaveOccurred())
    			Expect(err.Error()).To(ContainSubstring("SafeMath: addition overflow"))
			})
		})

        When("the there is no overflow but the start position is a large number", func() {
			It("Should revert", func() {
                bytes := common.Hex2Bytes("000000000000000000000000000000000000000001")
				_, err := BytesUtilsExporter.BytesToAddress(nil, bytes, big.NewInt(-23))
                Expect(err).To(HaveOccurred())
    			Expect(err.Error()).To(ContainSubstring("slicing out of range"))
			})
		})

        It("should return Address(0)", func() {
            bytes := common.Hex2Bytes("0000000000000000000000000000000000000000")
			addr, err := BytesUtilsExporter.BytesToAddress(nil, bytes, big.NewInt(0))
			Expect(err).ToNot(HaveOccurred())
			Expect(addr).To(Equal(common.HexToAddress("0x0")))
		})

		It("should return Address(1)", func() {
            bytes := common.Hex2Bytes("0000000000000000000000000000000000000001")
			addr, err := BytesUtilsExporter.BytesToAddress(nil, bytes, big.NewInt(0))
			Expect(err).ToNot(HaveOccurred())
			Expect(addr).To(Equal(common.HexToAddress("0x1")))
		})

        It("should return Address(1)", func() {
            bytes := common.Hex2Bytes("ffffffffffffffffffffffffffffffffffffffff0000000000000000000000000000000000000001ffffffffffffff")
			addr, err := BytesUtilsExporter.BytesToAddress(nil, bytes, big.NewInt(20))
			Expect(err).ToNot(HaveOccurred())
			Expect(addr).To(Equal(common.HexToAddress("0x1")))
		})

        It("should return fff...", func() {
            bytes := common.Hex2Bytes("ffffffffffffffffffffffffffffffffffffffff")
			addr, err := BytesUtilsExporter.BytesToAddress(nil, bytes, big.NewInt(0))
			Expect(err).ToNot(HaveOccurred())
			Expect(addr).To(Equal(common.HexToAddress("0xffffffffffffffffffffffffffffffffffffffff")))
		})

        It("should return the given random address", func() {
            bytes := common.Hex2Bytes("1234567890abcdefabcd0123456789abcdefabcd")
			addr, err := BytesUtilsExporter.BytesToAddress(nil, bytes, big.NewInt(0))
			Expect(err).ToNot(HaveOccurred())
			Expect(addr).To(Equal(common.BytesToAddress(bytes)))
		})
})
