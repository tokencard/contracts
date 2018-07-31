package wallet_test

import (
	"time"

	"github.com/ethereum/go-ethereum/common"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("topupLimit", func() {
	BeforeEach(func() {
		bankWallet.MustTransfer(be, controller.Address(), FIVE_HUNDRED_FINNEY)
	})

	Context("When the contract just has been deployed", func() {

		It("should have initial topup of 500 Finney available", func() {
			tl, err := w.TopupLimit(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(tl.String()).To(Equal(FIVE_HUNDRED_FINNEY.String()))

			w.Balance(nil, common.HexToAddress("0x0"))

			tl, err = w.TopupAvailable(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(tl.String()).To(Equal(FIVE_HUNDRED_FINNEY.String()))
		})

	})

	Describe("initializeTopupLimit", func() {
		var txSuccessful bool

		Context("When I try to initialize topup limit to one Gwei (below min topup limit)", func() {
			BeforeEach(func() {
				op := owner.TransactOpts()
				op.GasLimit = 65000
				tx, err := w.InitializeTopupLimit(op, ONE_GWEI)
				Expect(err).ToNot(HaveOccurred())
				be.Commit()
				txSuccessful = isSuccessful(tx)
			})

			It("Should fail", func() {
				Expect(txSuccessful).To(BeFalse())
			})

		})

		Context("When I try to initialize topup limit to one ETH (aboce max topup limit)", func() {
			BeforeEach(func() {
				op := owner.TransactOpts()
				op.GasLimit = 65000
				tx, err := w.InitializeTopupLimit(op, ONE_ETH)
				Expect(err).ToNot(HaveOccurred())
				be.Commit()
				txSuccessful = isSuccessful(tx)
			})

			It("Should fail", func() {
				Expect(txSuccessful).To(BeFalse())
			})

		})

		Context("When I initialize topup limit for the first time to one Finney", func() {
			BeforeEach(func() {
				op := owner.TransactOpts()
				op.GasLimit = 65000
				tx, err := w.InitializeTopupLimit(op, ONE_FINNEY)
				Expect(err).ToNot(HaveOccurred())
				be.Commit()
				txSuccessful = isSuccessful(tx)
			})
			It("Should succeed", func() {
				Expect(txSuccessful).To(BeTrue())
			})

			Context("When I try to initialize the limit again", func() {
				BeforeEach(func() {
					op := owner.TransactOpts()
					op.GasLimit = 65000
					tx, err := w.InitializeTopupLimit(op, FIVE_HUNDRED_FINNEY)
					Expect(err).ToNot(HaveOccurred())
					be.Commit()
					txSuccessful = isSuccessful(tx)
				})
				It("Should fail", func() {
					Expect(txSuccessful).To(BeFalse())
				})

			})
		})

	})

	Describe("Changing daily topup limit", func() {

		Context("When I attempt to submit daily topup limit below 1 Finney", func() {
			It("The transaction should fail", func() {
				op := owner.TransactOpts()
				op.GasLimit = 65000
				tx, err := w.SetTopupLimit(op, ONE_GWEI)
				Expect(err).ToNot(HaveOccurred())
				be.Commit()
				Expect(isSuccessful(tx)).To(BeFalse())
			})
		})

		Context("When I attempt to submit daily topup limit above 500 Finney", func() {
			It("The transaction should fail", func() {
				op := owner.TransactOpts()
				op.GasLimit = 65000
				tx, err := w.SetTopupLimit(op, ONE_ETH)
				Expect(err).ToNot(HaveOccurred())
				be.Commit()
				Expect(isSuccessful(tx)).To(BeFalse())
			})
		})

		Context("When controller tries to submit daily topup limit of 1 Finney", func() {
			var txSuccessful bool
			BeforeEach(func() {
				op := controller.TransactOpts()
				op.GasLimit = 65000
				tx, err := w.SetTopupLimit(op, ONE_FINNEY)
				Expect(err).ToNot(HaveOccurred())
				be.Commit()
				txSuccessful = isSuccessful(tx)

			})

			It("Should fail", func() {
				Expect(txSuccessful).To(BeFalse())
			})

		})

		Context("When a random person tries to submit daily topup limit of 1 Finney", func() {
			var txSuccessful bool
			BeforeEach(func() {
				op := randomPerson.TransactOpts()
				op.GasLimit = 65000
				tx, err := w.SetTopupLimit(op, ONE_FINNEY)
				Expect(err).ToNot(HaveOccurred())
				be.Commit()
				txSuccessful = isSuccessful(tx)

			})

			It("Should fail", func() {
				Expect(txSuccessful).To(BeFalse())
			})

		})

		Context("When I submit topup limit change to 1 Finney", func() {

			var txSuccessful bool

			BeforeEach(func() {
				op := owner.TransactOpts()
				op.GasLimit = 65000
				tx, err := w.SetTopupLimit(op, ONE_FINNEY)
				Expect(err).ToNot(HaveOccurred())
				be.Commit()
				txSuccessful = isSuccessful(tx)

			})

			It("Should succeed", func() {
				Expect(txSuccessful).To(BeTrue())
			})

			It("Should have pending topup limit of 1 Finney", func() {
				ptl, err := w.PendingTopupLimit(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(ptl.String()).To(Equal(ONE_FINNEY.String()))
			})

			Context("When the controller cancells the limit change", func() {
				BeforeEach(func() {
					tx, err := w.SetTopupLimitCancel(controller.TransactOpts())
					Expect(err).ToNot(HaveOccurred())
					be.Commit()
					txSuccessful = isSuccessful(tx)
				})

				It("Should succeed", func() {
					Expect(txSuccessful).To(BeTrue())
				})
			})

			Context("When I try to cancel the limit change", func() {
				BeforeEach(func() {
					op := owner.TransactOpts()
					op.GasLimit = 65000
					tx, err := w.SetTopupLimitCancel(op)
					Expect(err).ToNot(HaveOccurred())
					be.Commit()
					txSuccessful = isSuccessful(tx)
				})

				It("Should fail", func() {
					Expect(txSuccessful).To(BeFalse())
				})
			})

			Context("When a random person tries to cancel the limit change", func() {
				BeforeEach(func() {
					op := randomPerson.TransactOpts()
					op.GasLimit = 65000
					tx, err := w.SetTopupLimitCancel(op)
					Expect(err).ToNot(HaveOccurred())
					be.Commit()
					txSuccessful = isSuccessful(tx)
				})

				It("Should fail", func() {
					Expect(txSuccessful).To(BeFalse())
				})
			})

			Context("When I try to submit a second topup limit to 500 Finney", func() {
				BeforeEach(func() {
					op := owner.TransactOpts()
					op.GasLimit = 65000
					tx, err := w.SetTopupLimit(op, FIVE_HUNDRED_FINNEY)
					Expect(err).ToNot(HaveOccurred())
					be.Commit()
					txSuccessful = isSuccessful(tx)
				})

				It("Should fail", func() {
					Expect(txSuccessful).To(BeFalse())
				})

			})

			Context("When I try to confirm the topup limit", func() {
				BeforeEach(func() {
					op := owner.TransactOpts()
					op.GasLimit = 65000
					tx, err := w.SetTopupLimitConfirm(op)
					Expect(err).ToNot(HaveOccurred())
					be.Commit()
					txSuccessful = isSuccessful(tx)
				})

				It("Should fail", func() {
					Expect(txSuccessful).To(BeFalse())
				})
			})

			Context("When a random person tries to confirm the topup limit", func() {
				BeforeEach(func() {
					op := randomPerson.TransactOpts()
					op.GasLimit = 65000
					tx, err := w.SetTopupLimitConfirm(op)
					Expect(err).ToNot(HaveOccurred())
					be.Commit()
					txSuccessful = isSuccessful(tx)
				})

				It("Should fail", func() {
					Expect(txSuccessful).To(BeFalse())
				})
			})

			Context("When the controller confirms the topup limit", func() {
				BeforeEach(func() {
					tx, err := w.SetTopupLimitConfirm(controller.TransactOpts())
					Expect(err).ToNot(HaveOccurred())
					be.Commit()
					txSuccessful = isSuccessful(tx)
				})

				It("Should succeed", func() {
					Expect(txSuccessful).To(BeTrue())
				})

				It("Should have 1 Finney available for topups", func() {
					tl, err := w.TopupAvailable(nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(tl.String()).To(Equal(ONE_FINNEY.String()))
				})

				Context("When I submit a second topup limit to 500 Finney", func() {
					BeforeEach(func() {
						op := owner.TransactOpts()
						op.GasLimit = 65000
						tx, err := w.SetTopupLimit(op, FIVE_HUNDRED_FINNEY)
						Expect(err).ToNot(HaveOccurred())
						be.Commit()
						txSuccessful = isSuccessful(tx)
					})

					It("Should succeed", func() {
						Expect(txSuccessful).To(BeTrue())
					})

					Context("When the controller confirms the topup limit", func() {
						BeforeEach(func() {
							tx, err := w.SetTopupLimitConfirm(controller.TransactOpts())
							Expect(err).ToNot(HaveOccurred())
							be.Commit()
							txSuccessful = isSuccessful(tx)
						})

						It("Should have 1 Finney available for topups", func() {
							tl, err := w.TopupAvailable(nil)
							Expect(err).ToNot(HaveOccurred())
							Expect(tl.String()).To(Equal(ONE_FINNEY.String()))
						})
						Context("When I wait for longer than a day", func() {
							BeforeEach(func() {
								be.AdjustTime(time.Hour * 25)
								be.Commit()
							})

							It("Should have 500 Finney available for topups", func() {
								tl, err := w.TopupAvailable(nil)
								Expect(err).ToNot(HaveOccurred())
								Expect(tl.String()).To(Equal(FIVE_HUNDRED_FINNEY.String()))
							})

						})

					})

				})

			})

		})

	})

})
