package wallet_test

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tokencard/contracts/test/shared"
	"github.com/tokencard/ethertest"
)

var _ = Describe("GasTopUpLimit", func() {
	BeforeEach(func() {
		BankAccount.MustTransfer(Backend, Controller.Address(), FinneyToWei(500))
	})

	When("the contract just has been deployed", func() {

		It("should have initial GasTopUp of 500 Finney available", func() {
			tl, err := Wallet.GasTopUpLimit(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(tl.String()).To(Equal(FinneyToWei(500).String()))

			tl, err = Wallet.GasTopUpAvailable(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(tl.String()).To(Equal(FinneyToWei(500).String()))
		})

	})

	Describe("initializeGasTopUpLimit", func() {
		var txSuccessful bool

		When("I try to initialize GasTopUp limit to one Gwei (below min GasTopUp limit)", func() {
			BeforeEach(func() {
				tx, err := Wallet.InitializeGasTopUpLimit(Owner.TransactOpts(ethertest.WithGasLimit(100000)), GweiToWei(1))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isGasExhausted(tx, 100000)).To(BeFalse())
				txSuccessful = isSuccessful(tx)
			})

			It("should fail", func() {
				Expect(txSuccessful).To(BeFalse())
			})

		})

		When("I try to initialize GasTopUp limit to one ETH (above max GasTopUp limit)", func() {
			BeforeEach(func() {
				tx, err := Wallet.InitializeGasTopUpLimit(Owner.TransactOpts(ethertest.WithGasLimit(100000)), EthToWei(1))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				ethertest.WithGasLimit(100000)
				txSuccessful = isSuccessful(tx)
			})

			It("should fail", func() {
				Expect(txSuccessful).To(BeFalse())
			})

		})

		When("I initialize GasTopUp limit for the first time to one Finney", func() {
			BeforeEach(func() {
				tx, err := Wallet.InitializeGasTopUpLimit(Owner.TransactOpts(ethertest.WithGasLimit(65000)), FinneyToWei(1))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				txSuccessful = isSuccessful(tx)
			})

			It("should succeed", func() {
				Expect(txSuccessful).To(BeTrue())
			})

			It("should update the initialization flag", func() {
				initialized, err := Wallet.InitializedGasTopUpLimit(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(initialized).To(BeTrue())
			})

			It("should emit GasTopUp limit set event", func() {
				it, err := Wallet.FilterSetGasTopUpLimit(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(it.Next()).To(BeTrue())
				evt := it.Event
				Expect(it.Next()).To(BeFalse())
				Expect(evt.Sender).To(Equal(Owner.Address()))
				Expect(evt.Amount).To(Equal(FinneyToWei(1)))
			})

			When("I try to initialize the limit again", func() {
				BeforeEach(func() {
					tx, err := Wallet.InitializeGasTopUpLimit(Owner.TransactOpts(ethertest.WithGasLimit(165000)), FinneyToWei(500))
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					txSuccessful = isSuccessful(tx)
				})
				It("should fail", func() {
					Expect(txSuccessful).To(BeFalse())
				})

			})
		})

	})

	Describe("Changing daily GasTopUp limit", func() {

		When("I submit daily GasTopUp limit above 1 Finney before initialization", func() {
			It("should fail", func() {
				tx, err := Wallet.SubmitGasTopUpLimit(Owner.TransactOpts(ethertest.WithGasLimit(65000)), FinneyToWei(5))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeFalse())
			})
		})

		When("I submit daily GasTopUp limit below 1 Finney after initialization", func() {
			It("should fail", func() {
				tx, err := Wallet.InitializeGasTopUpLimit(Owner.TransactOpts(ethertest.WithGasLimit(65000)), FinneyToWei(5))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())

				tx, err = Wallet.SubmitGasTopUpLimit(Owner.TransactOpts(ethertest.WithGasLimit(65000)), GweiToWei(1))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeFalse())
			})
		})

		When("I submit daily GasTopUp limit above 500 Finney after initialization", func() {
			It("should fail", func() {
				tx, err := Wallet.InitializeGasTopUpLimit(Owner.TransactOpts(ethertest.WithGasLimit(65000)), FinneyToWei(5))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())

				tx, err = Wallet.SubmitGasTopUpLimit(Owner.TransactOpts(ethertest.WithGasLimit(65000)), EthToWei(1))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeFalse())
			})
		})

		When("controller submits daily GasTopUp limit of 1 Finney after initialization", func() {
			It("should fail", func() {
				tx, err := Wallet.InitializeGasTopUpLimit(Owner.TransactOpts(ethertest.WithGasLimit(65000)), FinneyToWei(5))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())

				tx, err = Wallet.SubmitGasTopUpLimit(Controller.TransactOpts(ethertest.WithGasLimit(65000)), FinneyToWei(1))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeFalse())
			})

		})

		When("a random person submits daily GasTopUp limit of 1 Finney after initialization", func() {
			It("should fail", func() {
				tx, err := Wallet.InitializeGasTopUpLimit(Owner.TransactOpts(ethertest.WithGasLimit(65000)), FinneyToWei(5))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())

				tx, err = Wallet.SubmitGasTopUpLimit(RandomAccount.TransactOpts(ethertest.WithGasLimit(65000)), FinneyToWei(1))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeFalse())
			})

		})

		When("I submit GasTopUp limit of 1 Finney after initialization", func() {
			BeforeEach(func() {
				tx, err := Wallet.InitializeGasTopUpLimit(Owner.TransactOpts(ethertest.WithGasLimit(65000)), FinneyToWei(5))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())

				tx, err = Wallet.SubmitGasTopUpLimit(Owner.TransactOpts(ethertest.WithGasLimit(65000)), FinneyToWei(1))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})

			It("should update the submission flag", func() {
				submitted, err := Wallet.SubmittedGasTopUpLimit(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(submitted).To(BeTrue())
			})

			It("should emit a submission event", func() {
				it, err := Wallet.FilterSubmittedGasTopUpLimitChange(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(it.Next()).To(BeTrue())
				evt := it.Event
				Expect(it.Next()).To(BeFalse())
				Expect(evt.Amount).To(Equal(FinneyToWei(1)))
			})

			It("should have pending GasTopUp limit of 1 Finney", func() {
				ptl, err := Wallet.PendingGasTopUpLimit(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(ptl.String()).To(Equal(FinneyToWei(1).String()))
			})

			When("the controller cancels the limit change", func() {
				BeforeEach(func() {
					tx, err := Wallet.CancelGasTopUpLimitUpdate(Controller.TransactOpts(), FinneyToWei(1))
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeTrue())
				})

				It("should emit a cancellation event", func() {
					it, err := Wallet.FilterCancelledGasTopUpLimitUpdate(nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(it.Next()).To(BeTrue())
					evt := it.Event
					Expect(it.Next()).To(BeFalse())
					Expect(evt.Sender).To(Equal(Controller.Address()))
				})

				It("should set pending limit to 0", func() {
					psl, err := Wallet.PendingGasTopUpLimit(nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(psl.String()).To(Equal("0"))
				})
			})

			When("the controller cancels the limit change with wrong amount", func() {
				var txSuccessful bool
				BeforeEach(func() {
					tx, err := Wallet.CancelGasTopUpLimitUpdate(Controller.TransactOpts(ethertest.WithGasLimit(100000)), FinneyToWei(2))
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					txSuccessful = isSuccessful(tx)

				})

				It("should fail", func() {
					Expect(txSuccessful).To(BeFalse())
				})
			})

			When("I try to cancel the limit change", func() {
				It("should fail", func() {
					tx, err := Wallet.CancelGasTopUpLimitUpdate(Owner.TransactOpts(ethertest.WithGasLimit(65000)), FinneyToWei(1))
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeFalse())
				})
			})

			When("a random person tries to cancel the limit change", func() {
				It("should fail", func() {
					tx, err := Wallet.CancelGasTopUpLimitUpdate(RandomAccount.TransactOpts(ethertest.WithGasLimit(65000)), FinneyToWei(1))
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeFalse())
				})
			})

			When("I try to submit a second GasTopUp limit of 500 Finney", func() {
				It("should fail", func() {
					tx, err := Wallet.SubmitGasTopUpLimit(Owner.TransactOpts(ethertest.WithGasLimit(65000)), FinneyToWei(500))
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeFalse())
				})

			})

			When("I try to confirm the GasTopUp limit", func() {
				It("should fail", func() {
					tx, err := Wallet.ConfirmGasTopUpLimit(Owner.TransactOpts(ethertest.WithGasLimit(65000)), FinneyToWei(1))
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeFalse())
				})
			})

			When("a random person tries to confirm the GasTopUp limit", func() {
				It("should fail", func() {
					tx, err := Wallet.ConfirmGasTopUpLimit(RandomAccount.TransactOpts(ethertest.WithGasLimit(65000)), FinneyToWei(1))
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeFalse())
				})
			})

			When("the controller confirms the GasTopUp limit using the wrong amount", func() {
				var txSuccessful bool
				BeforeEach(func() {
					tx, err := Wallet.ConfirmGasTopUpLimit(Controller.TransactOpts(ethertest.WithGasLimit(100000)), FinneyToWei(2))
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					txSuccessful = isSuccessful(tx)
				})

				It("should fail", func() {
					Expect(txSuccessful).To(BeFalse())
				})
			})

			When("the controller confirms the GasTopUp limit", func() {
				BeforeEach(func() {
					tx, err := Wallet.ConfirmGasTopUpLimit(Controller.TransactOpts(), FinneyToWei(1))
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeTrue())
				})

				It("should have 1 Finney available for GasTopUps", func() {
					tl, err := Wallet.GasTopUpAvailable(nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(tl.String()).To(Equal(FinneyToWei(1).String()))
				})

				When("I submit a second GasTopUp limit to 500 Finney", func() {
					BeforeEach(func() {
						tx, err := Wallet.SubmitGasTopUpLimit(Owner.TransactOpts(ethertest.WithGasLimit(65000)), FinneyToWei(500))
						Expect(err).ToNot(HaveOccurred())
						Backend.Commit()
						Expect(isSuccessful(tx)).To(BeTrue())
					})

					When("the controller confirms the GasTopUp limit", func() {
						BeforeEach(func() {
							tx, err := Wallet.ConfirmGasTopUpLimit(Controller.TransactOpts(), FinneyToWei(500))
							Expect(err).ToNot(HaveOccurred())
							Backend.Commit()
							Expect(isSuccessful(tx)).To(BeTrue())
						})

						It("should have 1 Finney available for GasTopUps", func() {
							tl, err := Wallet.GasTopUpAvailable(nil)
							Expect(err).ToNot(HaveOccurred())
							Expect(tl.String()).To(Equal(FinneyToWei(1).String()))
						})
						When("I wait for longer than a day", func() {
							BeforeEach(func() {
								Backend.AdjustTime(time.Hour * 25)
								Backend.Commit()
							})

							It("should have 500 Finney available for GasTopUps", func() {
								tl, err := Wallet.GasTopUpAvailable(nil)
								Expect(err).ToNot(HaveOccurred())
								Expect(tl.String()).To(Equal(FinneyToWei(500).String()))
							})

						})

					})

				})

			})

		})

	})

})
