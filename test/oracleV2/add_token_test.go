package oracle_v2_test

import (
	"github.com/ethereum/go-ethereum/common"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("addToken", func() {
	Context("When called by the controller", func() {
		Context("When the token is not already supported", func() {
			It("Should not fail", func() {
				tx, err := oracle.AddToken(controllerWallet.TransactOpts(), common.HexToAddress("0x1"), "ETH", 8)
				Expect(err).ToNot(HaveOccurred())
				be.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})
		})
		Context("When the token is already supported", func() {
			BeforeEach(func() {
				tx, err := oracle.AddToken(controllerWallet.TransactOpts(), common.HexToAddress("0x1"), "ETH", 8)
				Expect(err).ToNot(HaveOccurred())
				be.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})

			It("Should fail", func() {
				to := controllerWallet.TransactOpts()
				to.GasLimit = 100000
				tx, err := oracle.AddToken(to, common.HexToAddress("0x1"), "ETH", 8)
				Expect(err).ToNot(HaveOccurred())
				be.Commit()
				Expect(isSuccessful(tx)).To(BeFalse())
			})

		})
	})
})
