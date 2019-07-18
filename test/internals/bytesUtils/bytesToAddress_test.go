package bytesUtils_test

import (
    "math/big"

    "github.com/ethereum/go-ethereum/common"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = FDescribe("BytesUtils", func() {

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

        It("return Address(0)", func() {
            bytes := common.Hex2Bytes("0000000000000000000000000000000000000000")
			addr, err := BytesUtilsExporter.BytesToAddress(nil, bytes, big.NewInt(0))
			Expect(err).ToNot(HaveOccurred())
			Expect(addr).To(Equal(common.HexToAddress("0x0")))
		})

		It("return Address(1)", func() {
            bytes := common.Hex2Bytes("0000000000000000000000000000000000000001")
			addr, err := BytesUtilsExporter.BytesToAddress(nil, bytes, big.NewInt(0))
			Expect(err).ToNot(HaveOccurred())
			Expect(addr).To(Equal(common.HexToAddress("0x1")))
		})

        It("return the fff...", func() {
            bytes := common.Hex2Bytes("ffffffffffffffffffffffffffffffffffffffff")
			addr, err := BytesUtilsExporter.BytesToAddress(nil, bytes, big.NewInt(0))
			Expect(err).ToNot(HaveOccurred())
			Expect(addr).To(Equal(common.HexToAddress("0xffffffffffffffffffffffffffffffffffffffff")))
		})

        It("return the given random address", func() {
            bytes := common.Hex2Bytes("1234567890abcdefabcd0123456789abcdefabcd")
			addr, err := BytesUtilsExporter.BytesToAddress(nil, bytes, big.NewInt(0))
			Expect(err).ToNot(HaveOccurred())
			Expect(addr).To(Equal(common.HexToAddress("0x1234567890abcdefabcd0123456789abcdefabcd")))
		})
})
