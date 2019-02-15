package controller_test

import (
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

		BeforeEach(func() {
			tx, err := ControllerContract.RemoveAdmin(ControllerOwner.TransactOpts(ethertest.WithGasLimit(100000)), RandomAccount.Address())
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isGasExhausted(tx, 100000)).To(BeFalse())
			Expect(isSuccessful(tx)).To(BeFalse())
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

		BeforeEach(func() {
			tx, err := ControllerContract.RemoveAdmin(ControllerOwner.TransactOpts(ethertest.WithGasLimit(100000)), RandomAccount.Address())
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isGasExhausted(tx, 100000)).To(BeFalse())
			Expect(isSuccessful(tx)).To(BeFalse())
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

		BeforeEach(func() {
			tx, err := ControllerContract.RemoveAdmin(RandomAccount.TransactOpts(ethertest.WithGasLimit(100000)), ControllerAdmin.Address())
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isGasExhausted(tx, 100000)).To(BeFalse())
			Expect(isSuccessful(tx)).To(BeFalse())
		})
	})

})
