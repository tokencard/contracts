package wallet_test

import (
	"math/big"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("addController", func() {

	Context("When controller has 1 ETH", func() {

		BeforeEach(func() {
			bankWallet.Transfer(be, controller.Address(), ONE_ETH)
		})

		Context("When owner tries to remove controller", func() {
			It("Should fail", func() {
				to := owner.TransactOpts()
				to.GasLimit = 100000
				tx, err := w.RemoveController(to, randomPerson.Address())
				Expect(err).ToNot(HaveOccurred())
				be.Commit()
				Expect(isSuccessful(tx)).To(BeFalse())
			})
		})

		Context("When controller adds a random person as controller", func() {

			BeforeEach(func() {
				tx, err := w.AddController(controller.TransactOpts(), randomPerson.Address())
				Expect(err).ToNot(HaveOccurred())
				be.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})

			It("Should increase the controller count", func() {
				count, err := w.ControllerCount(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(count).To(Equal(big.NewInt(2)))
			})

			It("Should emit an AddController event", func() {
				it, err := w.FilterAddController(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(it.Next()).To(BeTrue())
				Expect(it.Next()).To(BeTrue())
				evt := it.Event
				Expect(it.Next()).To(BeFalse())
				Expect(evt.Account).To(Equal(randomPerson.Address()))
			})

			It("Should have added random person as controller", func() {
				isController, err := w.IsController(nil, randomPerson.Address())
				Expect(err).ToNot(HaveOccurred())
				Expect(isController).To(BeTrue())
			})
		})

	})

})

var _ = Describe("removeController", func() {

	Context("When controller has 1 ETH and random person is added as a controller", func() {

		BeforeEach(func() {
			bankWallet.Transfer(be, controller.Address(), ONE_ETH)
			tx, err := w.AddController(controller.TransactOpts(), randomPerson.Address())
			Expect(err).ToNot(HaveOccurred())
			be.Commit()
			Expect(isSuccessful(tx)).To(BeTrue())

		})

		Context("When owner tries to remove random person as controller", func() {
			It("Should fail", func() {
				to := owner.TransactOpts()
				to.GasLimit = 100000
				tx, err := w.RemoveController(to, randomPerson.Address())
				Expect(err).ToNot(HaveOccurred())
				be.Commit()
				Expect(isSuccessful(tx)).To(BeFalse())
			})
		})

		Context("When controller removes the random person as controller", func() {
			BeforeEach(func() {
				tx, err := w.RemoveController(controller.TransactOpts(), randomPerson.Address())
				Expect(err).ToNot(HaveOccurred())
				be.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})

			It("Should decrease the controller count", func() {
				count, err := w.ControllerCount(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(count).To(Equal(big.NewInt(1)))
			})

			It("Should emit an RemoveController event", func() {
				it, err := w.FilterRemoveController(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(it.Next()).To(BeTrue())
				evt := it.Event
				Expect(it.Next()).To(BeFalse())
				Expect(evt.Account).To(Equal(randomPerson.Address()))
			})

			It("Should have removed the random person as controller", func() {
				isController, err := w.IsController(nil, randomPerson.Address())
				Expect(err).ToNot(HaveOccurred())
				Expect(isController).To(BeFalse())
			})
		})

	})

})
