package controller_test

import (
	"github.com/ethereum/go-ethereum/core/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tokencard/contracts/test/shared"
	"github.com/ethereum/go-ethereum/common"
	"github.com/tokencard/ethertest"
)

var _ = Describe("AddController", func() {

	When("controller Admin calls AddController with a random address", func() {
		var err error
		var tx *types.Transaction
		BeforeEach(func() {
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

	})

	When("controller Admin calls AddController with it's own address", func() {
		var err error
		var tx *types.Transaction
		BeforeEach(func() {
			tx, err = ControllerContract.AddController(ControllerAdmin.TransactOpts(ethertest.WithGasLimit(100000)), ControllerAdmin.Address())
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isGasExhausted(tx, 100000)).To(BeFalse())
		})

		It("should not succeed", func() {
			Expect(isSuccessful(tx)).To(BeFalse())
		})

		It("should fail at the already owner requirenment", func() {
			Expect(TestRig.LastExecuted()).To(MatchRegexp(`require\(!_isAdmin\[_account\], "provided account is already an admin"\);`))
		})
	})

	When("controller Admin calls AddController with Owner's address", func() {
		var err error
		var tx *types.Transaction
		BeforeEach(func() {
			tx, err = ControllerContract.AddController(ControllerAdmin.TransactOpts(ethertest.WithGasLimit(100000)), ControllerOwner.Address())
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isGasExhausted(tx, 100000)).To(BeFalse())
		})

		It("should not succeed", func() {
			Expect(isSuccessful(tx)).To(BeFalse())
		})

		It("should fail at already controller requirenment", func() {
			Expect(TestRig.LastExecuted()).To(MatchRegexp(`require\(_account != owner\(\), "provided account is already the owner"\);`))
		})
	})

  When("controller Admin calls AddController with controller's address", func() {
		var err error
		var tx *types.Transaction
		BeforeEach(func() {
			tx, err = ControllerContract.AddController(ControllerAdmin.TransactOpts(ethertest.WithGasLimit(100000)), Controller.Address())
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isGasExhausted(tx, 100000)).To(BeFalse())
		})

		It("should not succeed", func() {
			Expect(isSuccessful(tx)).To(BeFalse())
		})

		It("should fail at already controller requirenment", func() {
			Expect(TestRig.LastExecuted()).To(MatchRegexp(`require\(!_isController\[_account\], "provided account is already a controller"\);`))
		})
	})

	When("controller Admin calls AddController with 0 address", func() {
		var err error
		var tx *types.Transaction
		BeforeEach(func() {
			tx, err = ControllerContract.AddController(ControllerAdmin.TransactOpts(ethertest.WithGasLimit(100000)), common.HexToAddress("0x0"))
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isGasExhausted(tx, 100000)).To(BeFalse())
		})

		It("should not succeed", func() {
			Expect(isSuccessful(tx)).To(BeFalse())
		})

		It("should fail at already controller requirenment", func() {
			Expect(TestRig.LastExecuted()).To(MatchRegexp(`require\(_account != address\(0\), "provided account is the zero address"\);`))
		})
	})

  When("Owner calls AddController with a random address", func() {
		var err error
		var tx *types.Transaction
		BeforeEach(func() {
			tx, err = ControllerContract.AddController(ControllerOwner.TransactOpts(ethertest.WithGasLimit(100000)), RandomAccount.Address())
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
		})

		It("should not succeed", func() {
			Expect(isSuccessful(tx)).To(BeFalse())
		})

		It("should fail at the not owner requirenment", func() {
			Expect(TestRig.LastExecuted()).To(MatchRegexp(`require\(isAdmin\(msg.sender\), "sender is not an admin"\);`))
		})
	})

  When("controller calls AddController with a random address", func() {
		var err error
		var tx *types.Transaction
		BeforeEach(func() {
			tx, err = ControllerContract.AddController(Controller.TransactOpts(ethertest.WithGasLimit(100000)), RandomAccount.Address())
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isGasExhausted(tx, 100000)).To(BeFalse())
		})

		It("should not succeed", func() {
			Expect(isSuccessful(tx)).To(BeFalse())
		})

		It("should fail at the notAdmin requirenment", func() {
			Expect(TestRig.LastExecuted()).To(MatchRegexp(`require\(isAdmin\(msg.sender\), "sender is not an admin"\);`))
		})
	})

})
