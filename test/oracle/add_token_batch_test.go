package oracle_test

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tokencard/ethertest"
)

func stringsToByte32(names ...string) [][32]byte {
	r := [][32]byte{}
	for _, n := range names {
		nb := [32]byte{}
		copy(nb[:], []byte(n))
		r = append(r, nb)
	}
	return r
}

var _ = Describe("addTokens", func() {

	Context("When called by the controller", func() {
		Context("When the added tokens are not already supported", func() {
			var tx *types.Transaction
			BeforeEach(func() {
				var err error
				tokens := []common.Address{common.HexToAddress("0x0"), common.HexToAddress("0x1")}
				tx, err = oracle.AddTokens(controller.TransactOpts(), tokens, stringsToByte32("BNT", "TKN"), []uint8{18, 8})
				Expect(err).ToNot(HaveOccurred())
				be.Commit()
			})

			It("Should succeed", func() {
				Expect(isSuccessful(tx)).To(BeTrue())
			})

			It("Should update the tokens map", func() {
				token, err := oracle.Tokens(nil, common.HexToAddress("0x0"))
				Expect(err).ToNot(HaveOccurred())
				Expect(token.Exists).To(BeTrue())
				Expect(token.Decimals).To(Equal(uint8(18)))

				token, err = oracle.Tokens(nil, common.HexToAddress("0x1"))
				Expect(err).ToNot(HaveOccurred())
				Expect(token.Exists).To(BeTrue())
				Expect(token.Decimals).To(Equal(uint8(8)))
			})

			Context("When at least one of the added tokens is already supported", func() {
				BeforeEach(func() {
					var err error
					tokens := []common.Address{common.HexToAddress("0x0"), common.HexToAddress("0x2")}
					tx, err = oracle.AddTokens(controller.TransactOpts(WithGasLimit(200000)), tokens, stringsToByte32("BNT", "OMG"), []uint8{18, 18})
					Expect(err).ToNot(HaveOccurred())
					be.Commit()
				})

				It("Should fail", func() {
					Expect(isSuccessful(tx)).To(BeFalse())
				})
			})

			Context("When none of the added tokens are supported", func() {
				BeforeEach(func() {
					var err error
					tokens := []common.Address{common.HexToAddress("0x4"), common.HexToAddress("0x5")}
					tx, err = oracle.AddTokens(controller.TransactOpts(WithGasLimit(300000)), tokens, stringsToByte32("MKR", "MTL"), []uint8{1, 12})
					Expect(err).ToNot(HaveOccurred())
					be.Commit()
				})

				It("Should pass", func() {
					Expect(isSuccessful(tx)).To(BeTrue())
				})
			})

		})

		Context("When the parameters have different lengths", func() {
			It("Should fail", func() {
				tokens := []common.Address{common.HexToAddress("0x0"), common.HexToAddress("0x1")}
				tx, err := oracle.AddTokens(controller.TransactOpts(WithGasLimit(300000)), tokens, stringsToByte32("BNT", "TKN", "ZRX"), []uint8{18, 8})
				Expect(err).ToNot(HaveOccurred())
				be.Commit()
				Expect(isGasExhausted(tx, 300000)).To(BeFalse())
				Expect(isSuccessful(tx)).To(BeFalse())
			})
		})

	})

	Context("When called by a random address", func() {
		It("Should fail", func() {
			tokens := []common.Address{common.HexToAddress("0x0"), common.HexToAddress("0x1")}
			tx, err := oracle.AddTokens(randomAccount.TransactOpts(WithGasLimit(300000)), tokens, stringsToByte32("BNT", "TKN"), []uint8{18, 8})
			Expect(err).ToNot(HaveOccurred())
			be.Commit()
			Expect(isGasExhausted(tx, 300000)).To(BeFalse())
			Expect(isSuccessful(tx)).To(BeFalse())

		})
	})

})
