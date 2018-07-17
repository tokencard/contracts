package wallet_test

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("spendAvailable", func() {
	BeforeEach(func() {
		bankWallet.MustTransfer(be, controller.Address(), ethToWei(1))
	})

	Context("When wallet has been deployed", func() {
		It("Should have 100 ETH spend available", func() {
			av, err := w.SpendAvailable(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(av.String()).To(Equal(ethToWei(100).String()))
		})

		Context("When a random person tries to initialize the spend limit", func() {
			It("Should fail", func() {
				to := randomPerson.TransactOpts()
				to.GasLimit = 100000
				tx, err := w.InitializeSpendLimit(to, ethToWei(1))
				Expect(err).ToNot(HaveOccurred())
				be.Commit()
				Expect(isSuccessful(tx)).To(BeFalse())
			})
		})

		Context("When controller tries to initialize the spend limit", func() {
			It("Should fail", func() {
				to := controller.TransactOpts()
				to.GasLimit = 100000
				tx, err := w.InitializeSpendLimit(to, ethToWei(1))
				Expect(err).ToNot(HaveOccurred())
				be.Commit()
				Expect(isSuccessful(tx)).To(BeFalse())
			})
		})

		Context("When I initialize spend limit to 1 ETH", func() {
			BeforeEach(func() {
				tx, err := w.InitializeSpendLimit(owner.TransactOpts(), ethToWei(1))
				Expect(err).ToNot(HaveOccurred())
				be.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})

			It("Should lower the spend available to 1 ETH", func() {
				av, err := w.SpendAvailable(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(av.String()).To(Equal(ethToWei(1).String()))
			})

			Context("When I try to initialize the spending limit again", func() {
				It("Should fail", func() {
					to := owner.TransactOpts()
					to.GasLimit = 100000
					tx, err := w.InitializeSpendLimit(to, ethToWei(1))
					Expect(err).ToNot(HaveOccurred())
					be.Commit()
					Expect(isSuccessful(tx)).To(BeFalse())
				})
			})
		})

		Context("When I initialize spend limit to 1000 ETH", func() {
			BeforeEach(func() {
				tx, err := w.InitializeSpendLimit(owner.TransactOpts(), ethToWei(1000))
				Expect(err).ToNot(HaveOccurred())
				be.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})

			It("Should keep the spend available at 100 ETH", func() {
				av, err := w.SpendAvailable(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(av.String()).To(Equal(ethToWei(100).String()))
			})

			Context("After I wait for 24 hours", func() {
				BeforeEach(func() {
					be.AdjustTime(time.Hour*24 + time.Second)
					be.Commit()
				})

				It("Should increase the spend available at 1000 ETH", func() {
					av, err := w.SpendAvailable(nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(av.String()).To(Equal(ethToWei(1000).String()))
				})

			})
		})

		Context("When I set the spend limit to 1 ETH", func() {
			BeforeEach(func() {
				tx, err := w.SetSpendLimit(owner.TransactOpts(), ethToWei(1))
				Expect(err).ToNot(HaveOccurred())
				be.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})

			It("Should not reduce the spend available", func() {
				av, err := w.SpendAvailable(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(av.String()).To(Equal(ethToWei(100).String()))
			})

			It("Should set pendingSpendLimit to 1 ETH", func() {
				psl, err := w.PendingSpendLimit(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(psl.String()).To(Equal(ethToWei(1).String()))
			})

			Context("When the controller confirms the spend limit", func() {
				BeforeEach(func() {
					tx, err := w.SetSpendLimitConfirm(controller.TransactOpts())
					Expect(err).ToNot(HaveOccurred())
					be.Commit()
					Expect(isSuccessful(tx)).To(BeTrue())
				})

				It("Should emit SetSpendLimit event", func() {
					it, err := w.FilterSetSpendLimit(nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(it.Next()).To(BeTrue())
					evt := it.Event
					Expect(it.Next()).To(BeFalse())
					Expect(evt.Amount.String()).To(Equal(ethToWei(1).String()))
				})

				It("Should reduce the spend available to 1 ETH", func() {
					av, err := w.SpendAvailable(nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(av.String()).To(Equal(ethToWei(1).String()))
				})

				Context("When I try to initialize the spend limit", func() {
					It("Should fail", func() {
						to := owner.TransactOpts()
						to.GasLimit = 100000
						tx, err := w.InitializeSpendLimit(to, ethToWei(1000))
						Expect(err).ToNot(HaveOccurred())
						be.Commit()
						Expect(isSuccessful(tx)).To(BeFalse())
					})
				})

			})
			Context("When the controller cancels the spend limit change", func() {
				BeforeEach(func() {
					tx, err := w.SetSpendLimitCancel(controller.TransactOpts())
					Expect(err).ToNot(HaveOccurred())
					be.Commit()
					Expect(isSuccessful(tx)).To(BeTrue())
				})

				It("Should set pendingSpendLimit to 0", func() {
					psl, err := w.PendingSpendLimit(nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(psl.String()).To(Equal("0"))
				})

				It("Should make it possible to set the spending limit again", func() {
					tx, err := w.SetSpendLimit(owner.TransactOpts(), ethToWei(2))
					Expect(err).ToNot(HaveOccurred())
					be.Commit()
					Expect(isSuccessful(tx)).To(BeTrue())
				})

			})

			Context("When I try to set spending limit again", func() {
				It("Should fail", func() {
					to := owner.TransactOpts()
					to.GasLimit = 100000
					tx, err := w.SetSpendLimit(to, ethToWei(1))
					Expect(err).ToNot(HaveOccurred())
					be.Commit()
					Expect(isSuccessful(tx)).To(BeFalse())
				})
			})
		})
	})

})
