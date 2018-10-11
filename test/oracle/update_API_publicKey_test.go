package oracle_test

import (
	"github.com/ethereum/go-ethereum/common"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("updateAPIPublicKey", func() {

	Context("When called by the controller", func() {
		It("Should not fail", func() {
			tx, err := oracle.UpdateAPIPublicKey(controllerWallet.TransactOpts(), common.Hex2Bytes("0xfffffff"))
			Expect(err).ToNot(HaveOccurred())
			be.Commit()
			Expect(isSuccessful(tx)).To(BeTrue())
		})
	})
	Context("When not called by the controller", func() {
		It("Should fail", func() {
			tx, err := oracle.UpdateAPIPublicKey(randomWallet.TransactOptsWithGasLimit(100000), common.Hex2Bytes("0xfffffff"))
			Expect(err).ToNot(HaveOccurred())
			be.Commit()
			Expect(isGasExhausted(tx, 100000)).To(BeFalse())
			Expect(isSuccessful(tx)).To(BeFalse())
		})
	})

})
