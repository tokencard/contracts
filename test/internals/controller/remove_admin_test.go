package controller_test

import (
	"github.com/ethereum/go-ethereum/core/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tokencard/contracts/test/shared"
	"github.com/tokencard/ethertest"
)

var _ = Describe("removeAdmin", func() {

	When("controller owner calls RemoveAdmin with a controller admin address", func() {

		BeforeEach(func() {
			tx, err := ControllerContract.RemoveAdmin(ControllerOwner.TransactOpts(), ControllerAdmin.Address())
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
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

	When("controller owner calls RemoveAdmin with a non controller admin address", func() {

		var tx *types.Transaction
		const gasLimit = 100000

		BeforeEach(func() {
			var err error
			tx, err = ControllerContract.RemoveAdmin(ControllerOwner.TransactOpts(ethertest.WithGasLimit(gasLimit)), RandomAccount.Address())
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
		})

		It("should not succeed", func() {
			Expect(isSuccessful(tx)).To(BeFalse())
		})

		It("should not exaust gas", func() {
			Expect(isGasExhausted(tx, gasLimit)).To(BeFalse())
		})

		It("should NOT decrease number of admins", func() {
			count, err := ControllerContract.AdminCount(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(count.String()).To(Equal("1"))
		})

		It("should NOT emit RemovedAdmin event", func() {
			it, err := ControllerContract.FilterRemovedAdmin(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(it.Next()).To(BeFalse())
		})
	})

	When("controller owner calls RemoveAdmin with a non controller admin address", func() {

		var tx *types.Transaction
		const gasLimit = 100000

		BeforeEach(func() {
			var err error
			tx, err = ControllerContract.RemoveAdmin(ControllerOwner.TransactOpts(ethertest.WithGasLimit(gasLimit)), RandomAccount.Address())
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
		})

		It("should not succeed", func() {
			Expect(isSuccessful(tx)).To(BeFalse())
		})

		It("should not exaust gas", func() {
			Expect(isGasExhausted(tx, gasLimit)).To(BeFalse())
		})

		It("should NOT decrease number of admins", func() {
			count, err := ControllerContract.AdminCount(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(count.String()).To(Equal("1"))
		})

		It("should NOT emit RemovedAdmin event", func() {
			it, err := ControllerContract.FilterRemovedAdmin(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(it.Next()).To(BeFalse())
		})
	})

	When("a Random address calls RemoveAdmin with a controller admin address", func() {

		var tx *types.Transaction
		const gasLimit = 100000

		BeforeEach(func() {
			var err error
			tx, err = ControllerContract.RemoveAdmin(RandomAccount.TransactOpts(ethertest.WithGasLimit(gasLimit)), ControllerAdmin.Address())
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
		})

		It("should not succeed", func() {
			Expect(isSuccessful(tx)).To(BeFalse())
		})

		It("should not exaust gas", func() {
			Expect(isGasExhausted(tx, gasLimit)).To(BeFalse())
		})

		It("should NOT decrease number of admins", func() {
			count, err := ControllerContract.AdminCount(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(count.String()).To(Equal("1"))
		})

		It("should NOT emit RemovedAdmin event", func() {
			it, err := ControllerContract.FilterRemovedAdmin(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(it.Next()).To(BeFalse())
		})

	})

})
