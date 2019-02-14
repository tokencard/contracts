package controller_test

import (
	"github.com/ethereum/go-ethereum/core/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tokencard/contracts/test/shared"
	"github.com/tokencard/ethertest"
)

var _ = Describe("removeAdmin", func() {

	When("controller owner calls RemoveAdmin with a controlle admin address", func() {
		var err error
		var tx *types.Transaction
		BeforeEach(func() {
			tx, err = ControllerContract.RemoveAdmin(ControllerOwner.TransactOpts(ethertest.WithGasLimit(100000)), ControllerAdmin.Address())
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
		})

		It("should succeed", func() {
			Expect(isSuccessful(tx)).To(BeTrue())
		})

		It("should decrease number of admins", func() {
			count, err := ControllerContract.AdminCount(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(count.String()).To(Equal("0"))
		})

		It("should emit RemovedAdmin event", func() {
			it, err := ControllerContract.FilterRemovedAdmin(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(it.Next()).To(BeTrue())
			evt := it.Event
			Expect(it.Next()).To(BeFalse())
			Expect(evt.Sender).To(Equal(ControllerOwner.Address()))
			Expect(evt.Admin).To(Equal(ControllerAdmin.Address()))
		})
	})

})
