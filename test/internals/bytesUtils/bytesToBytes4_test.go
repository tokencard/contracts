package bytesUtils_test

import (
    "math/big"
    "encoding/binary"
    "github.com/ethereum/go-ethereum/common"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("BytesToBytes4", func() {

        bytes4slice := make([]byte, 4)
        var bytes4 [4]byte

		When("the length is less than 4 bytes", func() {
			It("Should revert", func() {
                bytes := common.Hex2Bytes("000000")
                _, err := BytesUtilsExporter.BytesToBytes4(nil, bytes, big.NewInt(0))
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("slicing out of range"))
			})
		})

        When("the length is 4 bytes but the start index is not zero", func() {
			It("Should revert", func() {
                bytes := common.Hex2Bytes("00000000")
				_, err := BytesUtilsExporter.BytesToBytes4(nil, bytes, big.NewInt(1))
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("slicing out of range"))
			})
		})

        When("the there is an overflow and start position wraps around", func() {
			It("Should revert", func() {
                bytes := common.Hex2Bytes("0000000000000000000000000000000000000001")
				_, err := BytesUtilsExporter.BytesToBytes4(nil, bytes, big.NewInt(-4))
                Expect(err).To(HaveOccurred())
    			Expect(err.Error()).To(ContainSubstring("SafeMath: addition overflow"))
			})
		})

        It("should return 0", func() {
            bytes := common.Hex2Bytes("00000000")
			b4, err := BytesUtilsExporter.BytesToBytes4(nil, bytes, big.NewInt(0))
			Expect(err).ToNot(HaveOccurred())
            binary.BigEndian.PutUint32(bytes4slice, 0)
            copy(bytes4[:], bytes4slice)
			Expect(b4).To(Equal(bytes4))
		})

		It("should return 1", func() {
            bytes := common.Hex2Bytes("00000001")
			b4, err := BytesUtilsExporter.BytesToBytes4(nil, bytes, big.NewInt(0))
			Expect(err).ToNot(HaveOccurred())
            binary.BigEndian.PutUint32(bytes4slice, 1)
            copy(bytes4[:], bytes4slice)
			Expect(b4).To(Equal(bytes4))
		})

        It("should return 1", func() {
            bytes := common.Hex2Bytes("ffffffff00000001ffffffff")
			b4, err := BytesUtilsExporter.BytesToBytes4(nil, bytes, big.NewInt(4))
			Expect(err).ToNot(HaveOccurred())
            binary.BigEndian.PutUint32(bytes4slice, 1)
            copy(bytes4[:], bytes4slice)
			Expect(b4).To(Equal(bytes4))
		})

        It("should return  2^32-1", func() {
            bytes := common.Hex2Bytes("ffffffff")
			b4, err := BytesUtilsExporter.BytesToBytes4(nil, bytes, big.NewInt(0))
			Expect(err).ToNot(HaveOccurred())
            binary.BigEndian.PutUint32(bytes4slice, 4294967295)
            copy(bytes4[:], bytes4slice)
			Expect(b4).To(Equal(bytes4))
		})

        It("should return the given random number", func() {
            bytes := common.Hex2Bytes("abcdef89")
			b4, err := BytesUtilsExporter.BytesToBytes4(nil, bytes, big.NewInt(0))
			Expect(err).ToNot(HaveOccurred())
            binary.BigEndian.PutUint32(bytes4slice, 2882400137)
            copy(bytes4[:], bytes4slice)
			Expect(b4).To(Equal(bytes4))
		})

        It("should return the given random number", func() {
            bytes := common.Hex2Bytes("ffffffffffabcdef89ffff")
			b4, err := BytesUtilsExporter.BytesToBytes4(nil, bytes, big.NewInt(5))
			Expect(err).ToNot(HaveOccurred())
            binary.BigEndian.PutUint32(bytes4slice, 2882400137)
            copy(bytes4[:], bytes4slice)
			Expect(b4).To(Equal(bytes4))
		})
})
