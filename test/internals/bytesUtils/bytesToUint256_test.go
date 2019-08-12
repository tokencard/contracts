package bytesUtils_test

import (
    "math/big"

    "github.com/ethereum/go-ethereum/common"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("BytesToUint256", func() {

		When("the length is less than 32 bytes", func() {
			It("Should revert", func() {
                bytes := common.Hex2Bytes("00000000000000000000000000000000000000000000000000000000000000")
                _, err := BytesUtilsExporter.BytesToUint256(nil, bytes, big.NewInt(0))
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("slicing out of range"))
			})
		})

        When("the length is 32 bytes but the start index is not zero", func() {
			It("Should revert", func() {
                bytes := common.Hex2Bytes("0000000000000000000000000000000000000000000000000000000000000000")
				_, err := BytesUtilsExporter.BytesToUint256(nil, bytes, big.NewInt(1))
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("slicing out of range"))
			})
		})

        When("the there is an overflow and start position wraps around", func() {
			It("Should revert", func() {
                bytes := common.Hex2Bytes("0000000000000000000000000000000000000001")
				_, err := BytesUtilsExporter.BytesToUint256(nil, bytes, big.NewInt(-32))
                Expect(err).To(HaveOccurred())
    			Expect(err.Error()).To(ContainSubstring("SafeMath: addition overflow"))
			})
		})

        It("should return 0", func() {
            bytes := common.Hex2Bytes("0000000000000000000000000000000000000000000000000000000000000000")
			u256, err := BytesUtilsExporter.BytesToUint256(nil, bytes, big.NewInt(0))
			Expect(err).ToNot(HaveOccurred())
			Expect(u256.String()).To(Equal(big.NewInt(0).String()))
		})

		It("should return 1", func() {
            bytes := common.Hex2Bytes("0000000000000000000000000000000000000000000000000000000000000001")
			u256, err := BytesUtilsExporter.BytesToUint256(nil, bytes, big.NewInt(0))
			Expect(err).ToNot(HaveOccurred())
			Expect(u256.String()).To(Equal(big.NewInt(1).String()))
		})

        It("should return 1", func() {
            bytes := common.Hex2Bytes("ffffffff0000000000000000000000000000000000000000000000000000000000000001ffffffff")
			u256, err := BytesUtilsExporter.BytesToUint256(nil, bytes, big.NewInt(4))
			Expect(err).ToNot(HaveOccurred())
			Expect(u256.String()).To(Equal(big.NewInt(1).String()))
		})

        It("should return  2^256-1", func() {
            bytes := common.Hex2Bytes("ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff")
			u256, err := BytesUtilsExporter.BytesToUint256(nil, bytes, big.NewInt(0))
			Expect(err).ToNot(HaveOccurred())
            n := new(big.Int)
            n, ok := n.SetString("ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff",16)
            Expect(ok).To(BeTrue())
            Expect(err).ToNot(HaveOccurred())
			Expect(u256.String()).To(Equal(n.String()))
		})

        It("should return the given random number", func() {
            bytes := common.Hex2Bytes("abcdef0123456789abcdef0123456789abcdef0123456789abcdef0123456789")
			u256, err := BytesUtilsExporter.BytesToUint256(nil, bytes, big.NewInt(0))
			Expect(err).ToNot(HaveOccurred())
            n := new(big.Int)
            n, ok := n.SetString("abcdef0123456789abcdef0123456789abcdef0123456789abcdef0123456789",16)
            Expect(ok).To(BeTrue())
            Expect(err).ToNot(HaveOccurred())
			Expect(u256.String()).To(Equal(n.String()))
		})

        It("should return the given random number", func() {
            bytes := common.Hex2Bytes("ffffffffffabcdef0123456789abcdef0123456789abcdef0123456789abcdef0123456789ffff")
			u256, err := BytesUtilsExporter.BytesToUint256(nil, bytes, big.NewInt(5))
			Expect(err).ToNot(HaveOccurred())
            n := new(big.Int)
            n, ok := n.SetString("abcdef0123456789abcdef0123456789abcdef0123456789abcdef0123456789",16)
            Expect(ok).To(BeTrue())
            Expect(err).ToNot(HaveOccurred())
			Expect(u256.String()).To(Equal(n.String()))
		})
})
