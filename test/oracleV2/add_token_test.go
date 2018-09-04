package oracle_v2_test

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("addToken", func() {

	Context("When called by the controller", func() {
		Context("When the token is not supported", func() {
			It("Should not fail", func() {
				tx, err := oracle.AddToken(controllerWallet.TransactOpts(), common.HexToAddress("0x0"), "ETH", 18)
				Expect(err).ToNot(HaveOccurred())
				be.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})
		})
		Context("When the token is added/supported", func() {
			BeforeEach(func() {
				tx, err := oracle.AddToken(controllerWallet.TransactOpts(), common.HexToAddress("0x0"), "ETH", 18)
				Expect(err).ToNot(HaveOccurred())
				be.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})

			It("Should emit an TokenAddition event", func() {
				// s := make([]common.Address)
				// s.append(common.HexToAddress("0x0"))
				it, err := oracle.FilterTokenAddition(nil, []common.Address{common.HexToAddress("0x0")})
				Expect(err).ToNot(HaveOccurred())
				Expect(it.Next()).To(BeTrue())
				evt := it.Event
				Expect(it.Next()).To(BeFalse())
				Expect(evt.TokenID).To(Equal(common.HexToAddress("0x0")))
				Expect(evt.Label).To(Equal("ETH"))
			})
			It("Should update ContactAddresses", func() {
				// index := big.NewInt(0)
				a, err := oracle.ContractAddresses(nil, big.NewInt(0))
				Expect(err).ToNot(HaveOccurred())
				Expect(a).To(Equal(common.HexToAddress("0x0")))
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
			Expect(isGasExausted(tx, to.GasLimit)).To(BeFalse())
			Expect(isSuccessful(tx)).To(BeFalse())

		})
	})

})
