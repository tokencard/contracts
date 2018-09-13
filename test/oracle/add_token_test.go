package oracle_test

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("addToken", func() {

	Context("When called by the controller", func() {
		Context("When the added token is not already supported", func() {
			var tx *types.Transaction
			BeforeEach(func() {
				var err error
				tx, err = oracle.AddToken(controllerWallet.TransactOpts(), common.HexToAddress("0x0"), "ETH", 18)
				Expect(err).ToNot(HaveOccurred())
				be.Commit()
			})

			It("Should succeed", func() {
				Expect(isSuccessful(tx)).To(BeTrue())
			})

			It("Should emit a TokenAddition event", func() {
				it, err := oracle.FilterTokenAddition(nil, []common.Address{common.HexToAddress("0x0")})
				Expect(err).ToNot(HaveOccurred())
				Expect(it.Next()).To(BeTrue())
				evt := it.Event
				Expect(it.Next()).To(BeFalse())
				Expect(evt.Token).To(Equal(common.HexToAddress("0x0")))
			})

			It("Should update the contract addresses array", func() {
				addresses, err := oracle.ContractAddresses(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(addresses).To(HaveLen(1))
				Expect(addresses[0]).To(Equal(common.HexToAddress("0x0")))
			})

			It("Should update the tokens map", func() {
				token, err := oracle.Tokens(nil, common.HexToAddress("0x0"))
				Expect(err).ToNot(HaveOccurred())
				Expect(token.Supported).To(BeTrue())
				Expect(token.Decimals).To(Equal(uint8(18)))
			})
		})

		Context("When the token is already supported", func() {
			BeforeEach(func() {
				tx, err := oracle.AddToken(controllerWallet.TransactOpts(), common.HexToAddress("0x0"), "ETH", 18)
				Expect(err).ToNot(HaveOccurred())
				be.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})

			It("Should fail", func() {
				to := controllerWallet.TransactOpts()
				to.GasLimit = 100000
				tx, err := oracle.AddToken(to, common.HexToAddress("0x0"), "ETH", 18)
				Expect(err).ToNot(HaveOccurred())
				be.Commit()

				Expect(isSuccessful(tx)).To(BeFalse())

			})
		})
	})

	Context("When called by a random address", func() {
		It("Should fail", func() {
			to := randomWallet.TransactOpts()
			to.GasLimit = 300000
			tx, err := oracle.AddToken(to, common.HexToAddress("0x0"), "ETH", 18)
			Expect(err).ToNot(HaveOccurred())
			be.Commit()
			Expect(isGasExhausted(tx, to.GasLimit)).To(BeFalse())
			Expect(isSuccessful(tx)).To(BeFalse())

		})
	})

})
