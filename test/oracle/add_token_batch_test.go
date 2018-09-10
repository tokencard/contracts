package oracle_test

import (
	"github.com/ethereum/go-ethereum/common"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/ethereum/go-ethereum/core/types"
)

var _ = Describe("addTokenBatch", func() {

	Context("When called by the controller", func() {
		Context("When the added tokens are not already supported", func() {
			var tx *types.Transaction
			BeforeEach(func() {
				var err error
				tokens := []common.Address{common.HexToAddress("0x0"), common.HexToAddress("0x1")}
				tx, err = oracle.AddTokenBatch(controllerWallet.TransactOpts(), tokens, "BNT.TKN", []uint8{18,8})
				Expect(err).ToNot(HaveOccurred())
				be.Commit()
			})

			It("Should succeed", func() {
				Expect(isSuccessful(tx)).To(BeTrue())
			})

			It("Should update contract addresses array", func() {
				addresses, err := oracle.ContractAddresses(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(addresses).To(HaveLen(2))
				Expect(addresses[0]).To(Equal(common.HexToAddress("0x0")))
				Expect(addresses[1]).To(Equal(common.HexToAddress("0x1")))
			})

			It("Should update the tokens map", func() {
				token, err := oracle.Tokens(nil, common.HexToAddress("0x0"))
				Expect(err).ToNot(HaveOccurred())
				Expect(token.Supported).To(BeTrue())
				Expect(token.Decimals).To(Equal(uint8(18)))

				token, err = oracle.Tokens(nil, common.HexToAddress("0x1"))
				Expect(err).ToNot(HaveOccurred())
				Expect(token.Supported).To(BeTrue())
				Expect(token.Decimals).To(Equal(uint8(8)))
			})
		})

		Context("When the parameters have different lengths", func() {
			It("Should fail", func() {
				to := controllerWallet.TransactOpts()
				to.GasLimit = 300000
				tokens := []common.Address{common.HexToAddress("0x0"), common.HexToAddress("0x1")}
				tx, err := oracle.AddTokenBatch(to, tokens, "BNT.TKN.ZRX", []uint8{18,8})
				Expect(err).ToNot(HaveOccurred())
				be.Commit()
				Expect(isGasExhausted(tx, to.GasLimit)).To(BeFalse())
				Expect(isSuccessful(tx)).To(BeFalse())
			})
		})
	})

	Context("When called by a random address", func() {
		It("Should fail", func() {
			to := randomWallet.TransactOpts()
			to.GasLimit = 300000
			tokens := []common.Address{common.HexToAddress("0x0"), common.HexToAddress("0x1")}
			tx, err := oracle.AddTokenBatch(to, tokens, "BNT.TKN", []uint8{18,8})
			Expect(err).ToNot(HaveOccurred())
			be.Commit()
			Expect(isGasExhausted(tx, to.GasLimit)).To(BeFalse())
			Expect(isSuccessful(tx)).To(BeFalse())

		})
	})

})
