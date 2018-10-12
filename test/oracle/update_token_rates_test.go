package oracle_test

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("updateRates", func() {
	Context("When called by the controller", func() {
		Context("When the token is already supported", func() {
			BeforeEach(func() {
				tx, err := oracle.AddTokens(controller.TransactOpts(), []common.Address{common.HexToAddress("0x1")}, stringsToByte32("ETH"), []byte{8})
				Expect(err).ToNot(HaveOccurred())
				be.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})
			Context("When called by the controller", func() {
				var tx *types.Transaction
				var err error
				BeforeEach(func() {
					tx, err = oracle.UpdateTokenRates(controller.TransactOpts()) //TransactOptsWithValue(big.NewInt(1000000000)))
					Expect(err).ToNot(HaveOccurred())
					be.Commit()
				})
				It("Should succeed", func() {
					Expect(isSuccessful(tx)).To(BeTrue())
				})
				It("Should emit an Requested Update event", func() {
					it, err := oracle.FilterRequestedUpdate(nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(it.Next()).To(BeTrue())
					evt := it.Event
					Expect(it.Next()).To(BeFalse())
					Expect(evt.Label).To(Equal("ETH"))
				})
			})

		})
	})
})
