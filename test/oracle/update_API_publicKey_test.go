package oracle_test

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tokencard/contracts/test/shared"
	"github.com/tokencard/ethertest"
)

var _ = Describe("updateAPIPublicKey", func() {

	Context("When called by the controller", func() {

		var tx *types.Transaction

		BeforeEach(func() {
			var err error
			tx, err = Oracle.UpdateCryptoCompareAPIPublicKey(Controller.TransactOpts(), common.Hex2Bytes("fffffff"))
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
		})

		It("Should succeed", func() {
			Expect(isSuccessful(tx)).To(BeTrue())
		})

		It("Should emit a SetCryptoComparePublicKey event", func() {
			it, err := Oracle.FilterSetCryptoComparePublicKey(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(it.Next()).To(BeTrue())
			evt := it.Event
			Expect(it.Next()).To(BeFalse())
			Expect(evt.PublicKey).To(Equal(common.Hex2Bytes("fffffff")))
		})
	})

	Context("When not called by the controller", func() {
		It("Should fail", func() {
			tx, err := Oracle.UpdateCryptoCompareAPIPublicKey(RandomAccount.TransactOpts(ethertest.WithGasLimit(100000)), common.Hex2Bytes("fffffff"))
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isGasExhausted(tx, 100000)).To(BeFalse())
			Expect(isSuccessful(tx)).To(BeFalse())
		})
	})

})
