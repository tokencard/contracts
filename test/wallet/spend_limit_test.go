package wallet_test

import (
	"time"

	"github.com/ethereum/go-ethereum/common"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tokencard/ethertest"
)

var _ = Describe("spendAvailable", func() {
	BeforeEach(func() {
		bankAccount.MustTransfer(be, controller.Address(), ethToWei(1))
	})

	Context("When wallet has been deployed", func() {
		It("should have 100 ETH spend available", func() {
			av, err := w.SpendAvailable(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(av.String()).To(Equal(ethToWei(100).String()))
		})

		Context("When a random person tries to initialize the spend limit", func() {
			It("should fail", func() {
				tx, err := w.InitializeSpendLimit(randomAccount.TransactOpts(WithGasLimit(100000)), ethToWei(1))
				Expect(err).ToNot(HaveOccurred())
				be.Commit()
				Expect(isSuccessful(tx)).To(BeFalse())
			})
		})

		Context("When controller tries to initialize the spend limit", func() {
			It("should fail", func() {
				tx, err := w.InitializeSpendLimit(controller.TransactOpts(WithGasLimit(100000)), ethToWei(1))
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

			It("should emit a spend limit set event", func() {
				it, err := w.FilterSetSpendLimit(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(it.Next()).To(BeTrue())
				evt := it.Event
				Expect(it.Next()).To(BeFalse())
				Expect(evt.Sender).To(Equal(owner.Address()))
				Expect(evt.Amount).To(Equal(ethToWei(1)))
			})

			It("should lower the spend available to 1 ETH", func() {
				av, err := w.SpendAvailable(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(av.String()).To(Equal(ethToWei(1).String()))
			})

			It("should have spend limit of 1 ETH", func() {
				sl, err := w.SpendLimit(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(sl.String()).To(Equal(ethToWei(1).String()))
			})

			Context("When I try to initialize the spending limit again", func() {
				It("should fail", func() {
					tx, err := w.InitializeSpendLimit(owner.TransactOpts(WithGasLimit(100000)), ethToWei(1))
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

			It("should update the initialization flag", func() {
				initialized, err := w.InitializedSpendLimit(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(initialized).To(BeTrue())
			})

			It("should keep the spend available at 100 ETH", func() {
				av, err := w.SpendAvailable(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(av.String()).To(Equal(ethToWei(100).String()))
			})

			Context("After I wait for 24 hours", func() {
				BeforeEach(func() {
					be.AdjustTime(time.Hour*24 + time.Second)
					be.Commit()
				})

				It("should increase the spend available to 1000 ETH", func() {
					av, err := w.SpendAvailable(nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(av.String()).To(Equal(ethToWei(1000).String()))
				})

				Context("when the contract has one eth", func() {
					BeforeEach(func() {
						bankAccount.Transfer(be, wa, ethToWei(1))
					})

					Context("When I transfer one ETH", func() {
						BeforeEach(func() {
							tx, err := w.Transfer(owner.TransactOpts(), randomAccount.Address(), common.HexToAddress("0x"), ethToWei(1))
							Expect(err).ToNot(HaveOccurred())
							be.Commit()
							Expect(isSuccessful(tx)).To(BeTrue())
						})
						It("should have 999 ETH spend available", func() {
							sa, err := w.SpendAvailable(nil)
							Expect(err).ToNot(HaveOccurred())
							Expect(sa.String()).To(Equal(ethToWei(999).String()))
						})
					})
				})
			})
		})

		Context("When I submit the spend limit of 1 ETH before initialization", func() {
			It("should fail", func() {
				tx, err := w.SubmitSpendLimit(owner.TransactOpts(WithGasLimit(100000)), ethToWei(1))
				Expect(err).ToNot(HaveOccurred())
				be.Commit()
				Expect(isSuccessful(tx)).To(BeFalse())

				submitted, err := w.SubmittedSpendLimit(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(submitted).To(BeFalse())
			})
		})

		Context("When I submit the spend limit of 1 ETH after initialization", func() {
			BeforeEach(func() {
				tx, err := w.InitializeSpendLimit(owner.TransactOpts(WithGasLimit(100000)), ethToWei(2))
				Expect(err).ToNot(HaveOccurred())
				be.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())

				tx, err = w.SubmitSpendLimit(owner.TransactOpts(WithGasLimit(100000)), ethToWei(1))
				Expect(err).ToNot(HaveOccurred())
				be.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})

			It("should update the submission flag", func() {
				submitted, err := w.SubmittedSpendLimit(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(submitted).To(BeTrue())
			})

			It("should emit the submission event", func() {
				it, err := w.FilterSubmittedSpendLimitChange(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(it.Next()).To(BeTrue())
				evt := it.Event
				Expect(it.Next()).To(BeFalse())
				Expect(evt.Amount).To(Equal(ethToWei(1)))
			})

			It("should not reduce the spend available", func() {
				av, err := w.SpendAvailable(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(av.String()).To(Equal(ethToWei(2).String()))
			})

			It("should set the pending spend limit to 1 ETH", func() {
				psl, err := w.PendingSpendLimit(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(psl.String()).To(Equal(ethToWei(1).String()))
			})

			Context("When the controller confirms the spend limit", func() {
				BeforeEach(func() {
					tx, err := w.ConfirmSpendLimit(controller.TransactOpts())
					Expect(err).ToNot(HaveOccurred())
					be.Commit()
					Expect(isSuccessful(tx)).To(BeTrue())
				})

				It("should emit a spend limit set event", func() {
					it, err := w.FilterSetSpendLimit(nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(it.Next()).To(BeTrue())
					Expect(it.Next()).To(BeTrue())
					evt := it.Event
					Expect(it.Next()).To(BeFalse())
					Expect(evt.Amount).To(Equal(ethToWei(1)))
				})

				It("should reduce the spend available to 1 ETH", func() {
					av, err := w.SpendAvailable(nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(av.String()).To(Equal(ethToWei(1).String()))
				})

				Context("When I try to initialize the spend limit", func() {
					It("should fail", func() {
						tx, err := w.InitializeSpendLimit(owner.TransactOpts(WithGasLimit(100000)), ethToWei(1000))
						Expect(err).ToNot(HaveOccurred())
						be.Commit()
						Expect(isSuccessful(tx)).To(BeFalse())
					})
				})

			})
			Context("When the controller cancels the spend limit change", func() {
				BeforeEach(func() {
					tx, err := w.CancelSpendLimit(controller.TransactOpts())
					Expect(err).ToNot(HaveOccurred())
					be.Commit()
					Expect(isSuccessful(tx)).To(BeTrue())
				})

				It("should emit a cancellation event", func() {
					it, err := w.FilterCancelledSpendLimitChange(nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(it.Next()).To(BeTrue())
					evt := it.Event
					Expect(it.Next()).To(BeFalse())
					Expect(evt.Sender).To(Equal(controller.Address()))
				})

				It("should set pending spend limit to 0", func() {
					psl, err := w.PendingSpendLimit(nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(psl.String()).To(Equal("0"))
				})

				It("should make it possible to set the spending limit again", func() {
					tx, err := w.SubmitSpendLimit(owner.TransactOpts(), ethToWei(2))
					Expect(err).ToNot(HaveOccurred())
					be.Commit()
					Expect(isSuccessful(tx)).To(BeTrue())
				})

			})

			Context("When I try to set spending limit again", func() {
				It("should fail", func() {
					tx, err := w.SubmitSpendLimit(owner.TransactOpts(WithGasLimit(100000)), ethToWei(1))
					Expect(err).ToNot(HaveOccurred())
					be.Commit()
					Expect(isSuccessful(tx)).To(BeFalse())
				})
			})
		})
	})

})
