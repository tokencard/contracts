package controller_test

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tokencard/contracts/test/shared"
	"github.com/tokencard/ethertest"
)

var _ = Describe("addAdmin", func() {

	When("controller owner calls AddAdmin with a random address", func() {

		var tx *types.Transaction

		BeforeEach(func() {
			var err error
			tx, err = ControllerContract.AddAdmin(ControllerOwner.TransactOpts(), RandomAccount.Address())
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

		var tx *types.Transaction
		const gasLimit = 100000

		BeforeEach(func() {
			var err error
			tx, err = ControllerContract.AddAdmin(ControllerOwner.TransactOpts(ethertest.WithGasLimit(gasLimit)), ControllerOwner.Address())
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
		})

		It("should not succeed", func() {
			Expect(isSuccessful(tx)).To(BeFalse())
		})

		It("should not exaust gas", func() {
			Expect(isGasExhausted(tx, gasLimit)).To(BeFalse())
		})

		It("should fail at the already owner requirenment", func() {
			Expect(TestRig.LastExecuted()).To(MatchRegexp(`require\(!_isOwner\(_account\), "provided account is already the owner"\);`))
		})
	})

	When("controller owner calls AddAdmin with controller's address", func() {

		var tx *types.Transaction
		const gasLimit = 100000

		BeforeEach(func() {
			var err error
			tx, err = ControllerContract.AddAdmin(ControllerOwner.TransactOpts(ethertest.WithGasLimit(gasLimit)), Controller.Address())
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
		})

		It("should not succeed", func() {
			Expect(isSuccessful(tx)).To(BeFalse())
		})

		It("should not exaust gas", func() {
			Expect(isGasExhausted(tx, gasLimit)).To(BeFalse())
		})

		It("should fail at already controller requirenment", func() {
			Expect(TestRig.LastExecuted()).To(MatchRegexp(`require\(!_isController\[_account\], "provided account is already a controller"\);`))
		})
	})

	When("controller owner calls AddAdmin with admin's address", func() {

		var tx *types.Transaction
		const gasLimit = 100000

		BeforeEach(func() {
			var err error
			tx, err = ControllerContract.AddAdmin(ControllerOwner.TransactOpts(ethertest.WithGasLimit(gasLimit)), ControllerAdmin.Address())
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
		})

		It("should not succeed", func() {
			Expect(isSuccessful(tx)).To(BeFalse())
		})

		It("should not exaust gas", func() {
			Expect(isGasExhausted(tx, gasLimit)).To(BeFalse())
		})

		It("should fail at already controller requirenment", func() {
			Expect(TestRig.LastExecuted()).To(MatchRegexp(`require\(!_isAdmin\[_account\], "provided account is already an admin"\);`))
		})
	})

	When("controller owner calls AddAdmin with 0 address", func() {

		var tx *types.Transaction
		const gasLimit = 100000

		BeforeEach(func() {
			var err error
			tx, err = ControllerContract.AddAdmin(ControllerOwner.TransactOpts(ethertest.WithGasLimit(gasLimit)), common.HexToAddress("0x0"))
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
		})

		It("should not succeed", func() {
			Expect(isSuccessful(tx)).To(BeFalse())
		})

		It("should not exaust gas", func() {
			Expect(isGasExhausted(tx, gasLimit)).To(BeFalse())
		})

		It("should fail at already controller requirenment", func() {
			Expect(TestRig.LastExecuted()).To(MatchRegexp(`require\(_account != address\(0\), "provided account is the zero address"\);`))
		})
	})

	When("admin calls AddAdmin with a random address", func() {

		var tx *types.Transaction
		const gasLimit = 100000

		BeforeEach(func() {
			var err error
			tx, err = ControllerContract.AddAdmin(ControllerAdmin.TransactOpts(ethertest.WithGasLimit(gasLimit)), RandomAccount.Address())
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
		})

		It("should not succeed", func() {
			Expect(isSuccessful(tx)).To(BeFalse())
		})

		It("should not exaust gas", func() {
			Expect(isGasExhausted(tx, gasLimit)).To(BeFalse())
		})

		It("should fail at the not owner requirenment", func() {
			Expect(TestRig.LastExecuted()).To(MatchRegexp(`require\(_isOwner\(msg.sender\), "sender is not an owner"\);`))
		})
	})

	When("controller calls AddAdmin with a random address", func() {

		var tx *types.Transaction
		const gasLimit = 100000

		BeforeEach(func() {
			var err error
			tx, err = ControllerContract.AddAdmin(Controller.TransactOpts(ethertest.WithGasLimit(gasLimit)), RandomAccount.Address())
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
		})

		It("should not succeed", func() {
			Expect(isSuccessful(tx)).To(BeFalse())
		})

		It("should not exaust gas", func() {
			Expect(isGasExhausted(tx, gasLimit)).To(BeFalse())
		})

		It("should fail at the not owner requirenment", func() {
			Expect(TestRig.LastExecuted()).To(MatchRegexp(`require\(_isOwner\(msg.sender\), "sender is not an owner"\);`))
		})
	})

})
