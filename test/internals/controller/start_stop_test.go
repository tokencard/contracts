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

        It("should emit Stopped event", func() {
			it, err := ControllerContract.FilterStopped(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(it.Next()).To(BeTrue())
			evt := it.Event
			Expect(it.Next()).To(BeFalse())
			Expect(evt.Sender).To(Equal(ControllerAdmin.Address()))
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

    When("the owner stops the controller", func() {

		BeforeEach(func() {
			tx, err := ControllerContract.Stop(ControllerOwner.TransactOpts())
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeTrue())
		})

        It("The controller should be stopped", func() {
			stopped, err := ControllerContract.IsStopped(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(stopped).To(BeTrue())
		})

        It("should emit Stopped event", func() {
			it, err := ControllerContract.FilterStopped(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(it.Next()).To(BeTrue())
			evt := it.Event
			Expect(it.Next()).To(BeFalse())
			Expect(evt.Sender).To(Equal(ControllerOwner.Address()))
		})

        When("owner tries to start", func() {
    		BeforeEach(func() {
    			tx, err := ControllerContract.Start(ControllerOwner.TransactOpts())
    			Expect(err).ToNot(HaveOccurred())
    			Backend.Commit()
    			Expect(isSuccessful(tx)).To(BeTrue())
    		})

            It("The controller should start again", func() {
    			stopped, err := ControllerContract.IsStopped(nil)
    			Expect(err).ToNot(HaveOccurred())
    			Expect(stopped).To(BeFalse())
    		})

            It("should emit Started event", func() {
    			it, err := ControllerContract.FilterStarted(nil)
    			Expect(err).ToNot(HaveOccurred())
    			Expect(it.Next()).To(BeTrue())
    			evt := it.Event
    			Expect(it.Next()).To(BeFalse())
    			Expect(evt.Sender).To(Equal(ControllerOwner.Address()))
    		})
    	})

    })



})
