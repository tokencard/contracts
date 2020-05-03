package controller_test

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tokencard/contracts/v3/test/shared"
	"github.com/tokencard/ethertest"
)

var _ = Describe("AddController", func() {

	const gasLimit = 100000

	When("controller Admin calls AddController with a random address", func() {
		var tx *types.Transaction

		BeforeEach(func() {
			var err error
			tx, err = ControllerContract.AddController(ControllerAdmin.TransactOpts(), RandomAccount.Address())
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
		})

		It("should succeed", func() {
			Expect(isSuccessful(tx)).To(BeTrue())
		})

		It("should increase number of controllers", func() {
			count, err := ControllerContract.ControllerCount(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(count.String()).To(Equal("2"))
		})

		It("should emit AddController event", func() {
			it, err := ControllerContract.FilterAddedController(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(it.Next()).To(BeTrue())
			Expect(it.Next()).To(BeTrue())
			evt := it.Event
			Expect(it.Next()).To(BeFalse())
			Expect(evt.Sender).To(Equal(ControllerAdmin.Address()))
			Expect(evt.Controller).To(Equal(RandomAccount.Address()))
		})

		It("the new controller to be confirmed as a controller", func() {
			controller, err := ControllerContract.IsController(nil, RandomAccount.Address())
			Expect(err).ToNot(HaveOccurred())
			Expect(controller).To(BeTrue())
		})

		It("address(0x1) to not be controller", func() {
			controller, err := ControllerContract.IsController(nil, common.HexToAddress("0x1"))
			Expect(err).ToNot(HaveOccurred())
			Expect(controller).To(BeFalse())
		})

	})

	When("controller Admin calls AddController with it's own address", func() {
		It("should fail at the already owner requirenment", func() {
			tx, err := ControllerContract.AddController(ControllerAdmin.TransactOpts(ethertest.WithGasLimit(gasLimit)), ControllerAdmin.Address())
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeFalse())
			returnData, _ := ethCall(tx)
			Expect(string(returnData[len(returnData)-64:])).To(ContainSubstring("provided account is already an admin"))
		})
	})

	When("controller Admin calls AddController with Owner's address", func() {
		It("should fail at already controller requirenment", func() {
			tx, err := ControllerContract.AddController(ControllerAdmin.TransactOpts(ethertest.WithGasLimit(gasLimit)), ControllerOwner.Address())
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeFalse())
			returnData, _ := ethCall(tx)
			Expect(string(returnData[len(returnData)-64:])).To(ContainSubstring("provided account is already the owner"))
		})
	})

	When("controller Admin calls AddController with controller's address", func() {
		It("should fail at already controller requirenment", func() {
			tx, err := ControllerContract.AddController(ControllerAdmin.TransactOpts(ethertest.WithGasLimit(gasLimit)), Controller.Address())
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeFalse())
			returnData, _ := ethCall(tx)
			Expect(string(returnData[len(returnData)-64:])).To(ContainSubstring("provided account is already a controller"))
		})
	})

	When("controller Admin calls AddController with 0 address", func() {
		BeforeEach(func() {
			tx, err := ControllerContract.AddController(ControllerAdmin.TransactOpts(ethertest.WithGasLimit(gasLimit)), common.HexToAddress("0x0"))
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeFalse())
			returnData, _ := ethCall(tx)
			Expect(string(returnData[len(returnData)-64:])).To(ContainSubstring("provided account is the zero address"))
		})
	})

	When("Owner calls AddController with a random address", func() {
		It("should succeed", func() {
			tx, err := ControllerContract.AddController(ControllerOwner.TransactOpts(), RandomAccount.Address())
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeTrue())
		})
	})

	When("controller calls AddController with a random address", func() {
		It("should fail at the notAdmin requirenment", func() {
			tx, err := ControllerContract.AddController(Controller.TransactOpts(ethertest.WithGasLimit(gasLimit)), RandomAccount.Address())
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeFalse())
			returnData, _ := ethCall(tx)
			Expect(string(returnData[len(returnData)-64:])).To(ContainSubstring("sender is not admin or owner"))
		})
	})

})
