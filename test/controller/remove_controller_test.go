package controller_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tokencard/contracts/v3/test/shared"
	"github.com/tokencard/ethertest"
)

var _ = Describe("removeController", func() {

	const gasLimit = 100000

	When("controller owner calls RemoveController with a controller admin address", func() {

		BeforeEach(func() {
			tx, err := ControllerContract.RemoveController(ControllerOwner.TransactOpts(ethertest.WithGasLimit(gasLimit)), ControllerAdmin.Address())
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeFalse())
			returnData, _ := ethCall(tx)
			Expect(string(returnData[len(returnData)-64:])).To(ContainSubstring("provided account is not a controller"))
		})

		It("should NOT decrease number of admins", func() {
			count, err := ControllerContract.AdminCount(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(count.String()).To(Equal("1"))
		})

		It("should NOT emit RemoveController event", func() {
			it, err := ControllerContract.FilterRemovedController(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(it.Next()).To(BeFalse())
		})

	})

	When("controller admin calls RemoveController with a controller address", func() {

		BeforeEach(func() {
			tx, err := ControllerContract.RemoveController(ControllerAdmin.TransactOpts(), Controller.Address())
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeTrue())
		})

		It("should decrease number of controllers", func() {
			count, err := ControllerContract.ControllerCount(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(count.String()).To(Equal("0"))
		})

		It("should emit RemovedController event", func() {
			it, err := ControllerContract.FilterRemovedController(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(it.Next()).To(BeTrue())
			evt := it.Event
			Expect(it.Next()).To(BeFalse())
			Expect(evt.Sender).To(Equal(ControllerAdmin.Address()))
			Expect(evt.Controller).To(Equal(Controller.Address()))
		})
	})

	When("controller admin calls RemoveController with a non controller address", func() {

		BeforeEach(func() {
			tx, err := ControllerContract.RemoveController(ControllerAdmin.TransactOpts(ethertest.WithGasLimit(gasLimit)), RandomAccount.Address())
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeFalse())
			returnData, _ := ethCall(tx)
			Expect(string(returnData[len(returnData)-64:])).To(ContainSubstring("provided account is not a controller"))
		})

		It("should NOT decrease number of controllers", func() {
			count, err := ControllerContract.ControllerCount(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(count.String()).To(Equal("1"))
		})

		It("should NOT emit RemovedController event", func() {
			it, err := ControllerContract.FilterRemovedController(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(it.Next()).To(BeFalse())
		})

	})

	When("a Random address calls RemoveController with a controller admin address", func() {
		It("should fail", func() {
			tx, err := ControllerContract.RemoveController(RandomAccount.TransactOpts(ethertest.WithGasLimit(gasLimit)), ControllerAdmin.Address())
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeFalse())
			returnData, _ := ethCall(tx)
			Expect(string(returnData[len(returnData)-64:])).To(ContainSubstring("sender is not admin or owner"))
		})

	})

})
