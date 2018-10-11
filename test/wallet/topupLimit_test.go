package wallet_test

import (
	"time"

	"github.com/ethereum/go-ethereum/common"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tokencard/ethertest"
)

var _ = Describe("topupLimit", func() {
	BeforeEach(func() {
		bankWallet.MustTransfer(be, controller.Address(), finneyToWei(500))
	})

	Context("When the contract just has been deployed", func() {

		It("should have initial topup of 500 Finney available", func() {
			tl, err := w.TopupLimit(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(tl.String()).To(Equal(finneyToWei(500).String()))

			w.Balance(nil, common.HexToAddress("0x0"))

			tl, err = w.TopupAvailable(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(tl.String()).To(Equal(finneyToWei(500).String()))
		})

	})

	Describe("initializeTopupLimit", func() {
		var txSuccessful bool

		Context("When I try to initialize topup limit to one Gwei (below min topup limit)", func() {
			BeforeEach(func() {
				tx, err := w.InitializeTopupLimit(owner.TransactOpts(WithGasLimit(65000)), ONE_GWEI)
				Expect(err).ToNot(HaveOccurred())
				be.Commit()
				txSuccessful = isSuccessful(tx)
			})

			It("should fail", func() {
				Expect(txSuccessful).To(BeFalse())
			})

		})

		Context("When I try to initialize topup limit to one ETH (above max topup limit)", func() {
			BeforeEach(func() {
				tx, err := w.InitializeTopupLimit(owner.TransactOpts(WithGasLimit(65000)), ONE_ETH)
				Expect(err).ToNot(HaveOccurred())
				be.Commit()
				txSuccessful = isSuccessful(tx)
			})

			It("should fail", func() {
				Expect(txSuccessful).To(BeFalse())
			})

		})

		Context("When I initialize topup limit for the first time to one Finney", func() {
			BeforeEach(func() {
				tx, err := w.InitializeTopupLimit(owner.TransactOpts(WithGasLimit(65000)), ONE_FINNEY)
				Expect(err).ToNot(HaveOccurred())
				be.Commit()
				txSuccessful = isSuccessful(tx)
			})

			It("should succeed", func() {
				Expect(txSuccessful).To(BeTrue())
			})

			It("should update the initialization flag", func() {
				initialized, err := w.InitializedTopupLimit(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(initialized).To(BeTrue())
			})

			It("should emit topup limit set event", func() {
				it, err := w.FilterSetTopupLimit(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(it.Next()).To(BeTrue())
				evt := it.Event
				Expect(it.Next()).To(BeFalse())
				Expect(evt.Sender).To(Equal(owner.Address()))
				Expect(evt.Amount).To(Equal(ONE_FINNEY))
			})

			Context("When I try to initialize the limit again", func() {
				BeforeEach(func() {
					tx, err := w.InitializeTopupLimit(owner.TransactOpts(WithGasLimit(165000)), finneyToWei(500))
					Expect(err).ToNot(HaveOccurred())
					be.Commit()
					txSuccessful = isSuccessful(tx)
				})
				It("should fail", func() {
					Expect(txSuccessful).To(BeFalse())
				})

			})
		})

	})

	Describe("Changing daily topup limit", func() {

		Context("When I submit daily topup limit above 1 Finney before initialization", func() {
			It("should fail", func() {
				tx, err := w.SubmitTopupLimit(owner.TransactOpts(WithGasLimit(65000)), finneyToWei(5))
				Expect(err).ToNot(HaveOccurred())
				be.Commit()
				Expect(isSuccessful(tx)).To(BeFalse())
			})
		})

		Context("When I submit daily topup limit below 1 Finney after initialization", func() {
			It("should fail", func() {
				tx, err := w.InitializeTopupLimit(owner.TransactOpts(WithGasLimit(65000)), finneyToWei(5))
				Expect(err).ToNot(HaveOccurred())
				be.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())

				tx, err = w.SubmitTopupLimit(owner.TransactOpts(WithGasLimit(65000)), ONE_GWEI)
				Expect(err).ToNot(HaveOccurred())
				be.Commit()
				Expect(isSuccessful(tx)).To(BeFalse())
			})
		})

		Context("When I submit daily topup limit above 500 Finney after initialization", func() {
			It("should fail", func() {
				tx, err := w.InitializeTopupLimit(owner.TransactOpts(WithGasLimit(65000)), finneyToWei(5))
				Expect(err).ToNot(HaveOccurred())
				be.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())

				tx, err = w.SubmitTopupLimit(owner.TransactOpts(WithGasLimit(65000)), ONE_ETH)
				Expect(err).ToNot(HaveOccurred())
				be.Commit()
				Expect(isSuccessful(tx)).To(BeFalse())
			})
		})

		Context("When controller submits daily topup limit of 1 Finney after initialization", func() {
			It("should fail", func() {
				tx, err := w.InitializeTopupLimit(owner.TransactOpts(WithGasLimit(65000)), finneyToWei(5))
				Expect(err).ToNot(HaveOccurred())
				be.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())

				tx, err = w.SubmitTopupLimit(controller.TransactOpts(WithGasLimit(65000)), ONE_FINNEY)
				Expect(err).ToNot(HaveOccurred())
				be.Commit()
				Expect(isSuccessful(tx)).To(BeFalse())
			})

		})

		Context("When a random person submits daily topup limit of 1 Finney after initialization", func() {
			It("should fail", func() {
				tx, err := w.InitializeTopupLimit(owner.TransactOpts(WithGasLimit(65000)), finneyToWei(5))
				Expect(err).ToNot(HaveOccurred())
				be.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())

				tx, err = w.SubmitTopupLimit(randomPerson.TransactOpts(WithGasLimit(65000)), ONE_FINNEY)
				Expect(err).ToNot(HaveOccurred())
				be.Commit()
				Expect(isSuccessful(tx)).To(BeFalse())
			})

		})

		Context("When I submit topup limit of 1 Finney after initialization", func() {
			BeforeEach(func() {
				tx, err := w.InitializeTopupLimit(owner.TransactOpts(WithGasLimit(65000)), finneyToWei(5))
				Expect(err).ToNot(HaveOccurred())
				be.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())

				tx, err = w.SubmitTopupLimit(owner.TransactOpts(WithGasLimit(65000)), ONE_FINNEY)
				Expect(err).ToNot(HaveOccurred())
				be.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})

			It("should update the submission flag", func() {
				submitted, err := w.SubmittedTopupLimit(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(submitted).To(BeTrue())
			})

			It("should emit a submission event", func() {
				it, err := w.FilterSubmittedTopupLimitChange(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(it.Next()).To(BeTrue())
				evt := it.Event
				Expect(it.Next()).To(BeFalse())
				Expect(evt.Amount).To(Equal(ONE_FINNEY))
			})

			It("should have pending topup limit of 1 Finney", func() {
				ptl, err := w.PendingTopupLimit(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(ptl.String()).To(Equal(ONE_FINNEY.String()))
			})

			Context("When the controller cancels the limit change", func() {
				BeforeEach(func() {
					tx, err := w.CancelTopupLimit(controller.TransactOpts())
					Expect(err).ToNot(HaveOccurred())
					be.Commit()
					Expect(isSuccessful(tx)).To(BeTrue())
				})

				It("should emit a cancellation event", func() {
					it, err := w.FilterCancelledTopupLimitChange(nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(it.Next()).To(BeTrue())
					evt := it.Event
					Expect(it.Next()).To(BeFalse())
					Expect(evt.Sender).To(Equal(controller.Address()))
				})

				It("should set pending limit to 0", func() {
					psl, err := w.PendingTopupLimit(nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(psl.String()).To(Equal("0"))
				})
			})

			Context("When I try to cancel the limit change", func() {
				It("should fail", func() {
					tx, err := w.CancelTopupLimit(owner.TransactOpts(WithGasLimit(65000)))
					Expect(err).ToNot(HaveOccurred())
					be.Commit()
					Expect(isSuccessful(tx)).To(BeFalse())
				})
			})

			Context("When a random person tries to cancel the limit change", func() {
				It("should fail", func() {
					tx, err := w.CancelTopupLimit(randomPerson.TransactOpts(WithGasLimit(65000)))
					Expect(err).ToNot(HaveOccurred())
					be.Commit()
					Expect(isSuccessful(tx)).To(BeFalse())
				})
			})

			Context("When I try to submit a second topup limit of 500 Finney", func() {
				It("should fail", func() {
					tx, err := w.SubmitTopupLimit(owner.TransactOpts(WithGasLimit(65000)), finneyToWei(500))
					Expect(err).ToNot(HaveOccurred())
					be.Commit()
					Expect(isSuccessful(tx)).To(BeFalse())
				})

			})

			Context("When I try to confirm the topup limit", func() {
				It("should fail", func() {
					tx, err := w.ConfirmTopupLimit(owner.TransactOpts(WithGasLimit(65000)))
					Expect(err).ToNot(HaveOccurred())
					be.Commit()
					Expect(isSuccessful(tx)).To(BeFalse())
				})
			})

			Context("When a random person tries to confirm the topup limit", func() {
				It("should fail", func() {
					tx, err := w.ConfirmTopupLimit(randomPerson.TransactOpts(WithGasLimit(65000)))
					Expect(err).ToNot(HaveOccurred())
					be.Commit()
					Expect(isSuccessful(tx)).To(BeFalse())
				})
			})

			Context("When the controller confirms the topup limit", func() {
				BeforeEach(func() {
					tx, err := w.ConfirmTopupLimit(controller.TransactOpts())
					Expect(err).ToNot(HaveOccurred())
					be.Commit()
					Expect(isSuccessful(tx)).To(BeTrue())
				})

				It("should have 1 Finney available for topups", func() {
					tl, err := w.TopupAvailable(nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(tl.String()).To(Equal(ONE_FINNEY.String()))
				})

				Context("When I submit a second topup limit to 500 Finney", func() {
					BeforeEach(func() {
						tx, err := w.SubmitTopupLimit(owner.TransactOpts(WithGasLimit(65000)), finneyToWei(500))
						Expect(err).ToNot(HaveOccurred())
						be.Commit()
						Expect(isSuccessful(tx)).To(BeTrue())
					})

					Context("When the controller confirms the topup limit", func() {
						BeforeEach(func() {
							tx, err := w.ConfirmTopupLimit(controller.TransactOpts())
							Expect(err).ToNot(HaveOccurred())
							be.Commit()
							Expect(isSuccessful(tx)).To(BeTrue())
						})

						It("should have 1 Finney available for topups", func() {
							tl, err := w.TopupAvailable(nil)
							Expect(err).ToNot(HaveOccurred())
							Expect(tl.String()).To(Equal(ONE_FINNEY.String()))
						})
						Context("When I wait for longer than a day", func() {
							BeforeEach(func() {
								be.AdjustTime(time.Hour * 25)
								be.Commit()
							})

							It("should have 500 Finney available for topups", func() {
								tl, err := w.TopupAvailable(nil)
								Expect(err).ToNot(HaveOccurred())
								Expect(tl.String()).To(Equal(finneyToWei(500).String()))
							})

						})

					})

				})

			})

		})

	})

})
