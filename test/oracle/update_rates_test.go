package oracle_test

import (
	"github.com/ethereum/go-ethereum/common"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("updateRates", func() {
	Context("When called by the controller", func() {
		Context("When the token is already supported", func() {
			BeforeEach(func() {
				tx, err := oracle.AddTokens(controllerWallet.TransactOpts(), []common.Address{common.HexToAddress("0x1")}, stringsToByte32("ETH"), []byte{8})
				Expect(err).ToNot(HaveOccurred())
				be.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})

			It("Should do stuff", func() {
				tx, err := oracle.UpdateTokenRates(controllerWallet.TransactOpts())
				Expect(err).ToNot(HaveOccurred())
				be.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})

		})
	})
})
