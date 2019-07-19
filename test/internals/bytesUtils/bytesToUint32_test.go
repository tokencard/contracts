package bytesUtils_test

import (
    "math/big"

    "github.com/ethereum/go-ethereum/common"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("BytesToUint32", func() {

		When("the length is less than 4 bytes", func() {
			It("Should revert", func() {
                bytes := common.Hex2Bytes("000000")
                _, err := BytesUtilsExporter.BytesToUint32(nil, bytes, big.NewInt(0))
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("slicing out of range"))
			})
		})

        When("the length is 4 bytes but the start index is not zero", func() {
			It("Should revert", func() {
                bytes := common.Hex2Bytes("00000000")
				_, err := BytesUtilsExporter.BytesToUint32(nil, bytes, big.NewInt(1))
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("slicing out of range"))
			})
		})

        It("should return 0", func() {
            bytes := common.Hex2Bytes("00000000")
			u32, err := BytesUtilsExporter.BytesToUint32(nil, bytes, big.NewInt(0))
			Expect(err).ToNot(HaveOccurred())
			Expect(u32).To(Equal(uint32(0)))
		})

		It("should return 1", func() {
            bytes := common.Hex2Bytes("00000001")
			u32, err := BytesUtilsExporter.BytesToUint32(nil, bytes, big.NewInt(0))
			Expect(err).ToNot(HaveOccurred())
			Expect(u32).To(Equal(uint32(1)))
		})

        It("should return 1", func() {
            bytes := common.Hex2Bytes("ffffffff00000001ffffffff")
			u32, err := BytesUtilsExporter.BytesToUint32(nil, bytes, big.NewInt(4))
			Expect(err).ToNot(HaveOccurred())
			Expect(u32).To(Equal(uint32(1)))
		})

        It("should return  2^32-1", func() {
            bytes := common.Hex2Bytes("ffffffff")
			u32, err := BytesUtilsExporter.BytesToUint32(nil, bytes, big.NewInt(0))
			Expect(err).ToNot(HaveOccurred())
			Expect(u32).To(Equal(uint32(4294967295)))
		})

        It("should return the given random number", func() {
            bytes := common.Hex2Bytes("abcdef89")
			u32, err := BytesUtilsExporter.BytesToUint32(nil, bytes, big.NewInt(0))
			Expect(err).ToNot(HaveOccurred())
			Expect(u32).To(Equal(uint32(2882400137)))
		})

        It("should return the given random number", func() {
            bytes := common.Hex2Bytes("ffffffffffabcdef89ffff")
			u32, err := BytesUtilsExporter.BytesToUint32(nil, bytes, big.NewInt(5))
			Expect(err).ToNot(HaveOccurred())
			Expect(u32).To(Equal(uint32(2882400137)))
		})
})
