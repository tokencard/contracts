package oracle_test

import (
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tokencard/contracts/test/shared"
	"github.com/tokencard/ethertest"
)

var _ = Describe("setCustomGasPrice", func() {

	Context("When called by the controller", func() {
		var tx *types.Transaction

		BeforeEach(func() {
			var err error
			tx, err = Oracle.SetCustomGasPrice(Controller.TransactOpts(), big.NewInt(777))
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
		})

		It("Should succeed", func() {
			Expect(isSuccessful(tx)).To(BeTrue())
		})

		It("Should emit a SetGasPrice event", func() {
			it, err := Oracle.FilterSetGasPrice(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(it.Next()).To(BeTrue())
			evt := it.Event
			Expect(it.Next()).To(BeFalse())
			Expect(evt.GasPrice.String()).To(Equal(big.NewInt(777).String()))
		})
	})
	Context("When not called by the controller", func() {
		It("Should fail", func() {
			tx, err := Oracle.SetCustomGasPrice(RandomAccount.TransactOpts(ethertest.WithGasLimit(100000)), big.NewInt(777))
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isGasExhausted(tx, 100000)).To(BeFalse())
			Expect(isSuccessful(tx)).To(BeFalse())
		})
	})

})
