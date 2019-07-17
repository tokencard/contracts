package controller_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tokencard/contracts/test/shared"
	"github.com/tokencard/ethertest"
)

var _ = Describe("Controller Stopping", func() {

	When("a random address tries to stop", func() {
		It("Should fail", func() {
			tx, err := ControllerContract.Stop(RandomAccount.TransactOpts(ethertest.WithGasLimit(1000000)))
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeFalse())
		})
	})

	When("controller tries to stop", func() {
		It("Should fail", func() {
			tx, err := ControllerContract.Stop(Controller.TransactOpts(ethertest.WithGasLimit(1000000)))
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeFalse())
		})
	})

	When("the Controller admin stops the controller", func() {

		BeforeEach(func() {
			tx, err := ControllerContract.Stop(ControllerAdmin.TransactOpts())
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeTrue())
		})

		It("The controller should be stopped", func() {
			stopped, err := ControllerContract.IsStopped(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(stopped).To(BeTrue())
		})

	})

	When("a random address tries to start", func() {
		It("Should fail", func() {
			tx, err := ControllerContract.Start(RandomAccount.TransactOpts(ethertest.WithGasLimit(1000000)))
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeFalse())
		})
	})

	When("controller tries to start", func() {
		It("Should fail", func() {
			tx, err := ControllerContract.Start(Controller.TransactOpts(ethertest.WithGasLimit(1000000)))
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeFalse())
		})
	})

	When("controller admin tries to start", func() {
		It("Should fail", func() {
			tx, err := ControllerContract.Start(ControllerAdmin.TransactOpts(ethertest.WithGasLimit(1000000)))
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeFalse())
		})
	})

	When("owner tries to start", func() {
		It("Should pass", func() {
			tx, err := ControllerContract.Start(ControllerOwner.TransactOpts())
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeTrue())
		})
	})

})
