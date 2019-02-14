package controller_test

import (
	"github.com/ethereum/go-ethereum/core/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tokencard/contracts/test/shared"

	"github.com/tokencard/ethertest"
)

var _ = Describe("addAdmin", func() {

	When("controller owner calls AddAdmin with a random address", func() {
		var err error
		var tx *types.Transaction
		BeforeEach(func() {
			tx, err = ControllerContract.AddAdmin(ControllerOwner.TransactOpts(ethertest.WithGasLimit(100000)), RandomAccount.Address())
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
		})

		It("should succeed", func() {
			Expect(isSuccessful(tx)).To(BeTrue())
		})

		It("should increase number of admins", func() {
			count, err := ControllerContract.AdminCount(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(count.String()).To(Equal("2"))
		})

		It("should emit AddedAdmin event", func() {
			it, err := ControllerContract.FilterAddedAdmin(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(it.Next()).To(BeTrue())
			Expect(it.Next()).To(BeTrue())
			evt := it.Event
			Expect(it.Next()).To(BeFalse())
			Expect(evt.Sender).To(Equal(ControllerOwner.Address()))
			Expect(evt.Admin).To(Equal(RandomAccount.Address()))
		})

	})

	When("controller owner calls AddAdmin with it's own address", func() {
		var err error
		var tx *types.Transaction
		BeforeEach(func() {
			tx, err = ControllerContract.AddAdmin(ControllerOwner.TransactOpts(ethertest.WithGasLimit(100000)), ControllerOwner.Address())
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
		})

		It("should not succeed", func() {
			Expect(isSuccessful(tx)).To(BeFalse())
		})

		It("should fail at the already owner requirenment", func() {
			Expect(TestRig.LastExecuted()).To(MatchRegexp(`require\(_account != _owner, "provided account is already the owner"\);`))
		})
	})

	When("controller owner calls AddAdmin with controller's address", func() {
		var err error
		var tx *types.Transaction
		BeforeEach(func() {
			tx, err = ControllerContract.AddAdmin(ControllerOwner.TransactOpts(ethertest.WithGasLimit(100000)), Controller.Address())
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
		})

		It("should not succeed", func() {
			Expect(isSuccessful(tx)).To(BeFalse())
		})

		It("should fail at already controller requirenment", func() {
			Expect(TestRig.LastExecuted()).To(MatchRegexp(`require\(!_isController\[_account\], "provided account is already a controller"\);`))
		})
	})

  When("admin calls AddAdmin with a random address", func() {
		var err error
		var tx *types.Transaction
		BeforeEach(func() {
			tx, err = ControllerContract.AddAdmin(ControllerAdmin.TransactOpts(ethertest.WithGasLimit(100000)), RandomAccount.Address())
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
		})

		It("should not succeed", func() {
			Expect(isSuccessful(tx)).To(BeFalse())
		})

		It("should fail at the not owner requirenment", func() {
			Expect(TestRig.LastExecuted()).To(MatchRegexp(`require\(msg.sender == _owner, "sender is not an owner"\);`))
		})
	})

  When("controller calls AddAdmin with a random address", func() {
		var err error
		var tx *types.Transaction
		BeforeEach(func() {
			tx, err = ControllerContract.AddAdmin(Controller.TransactOpts(ethertest.WithGasLimit(100000)), RandomAccount.Address())
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
		})

		It("should not succeed", func() {
			Expect(isSuccessful(tx)).To(BeFalse())
		})

		It("should fail at the not owner requirenment", func() {
			Expect(TestRig.LastExecuted()).To(MatchRegexp(`require\(msg.sender == _owner, "sender is not an owner"\);`))
		})
	})

})
