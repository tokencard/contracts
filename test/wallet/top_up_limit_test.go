package wallet_test

import (
	"time"

	// "github.com/ethereum/go-ethereum/core/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tokencard/contracts/v3/test/shared"
	"github.com/tokencard/ethertest"
)

var _ = Describe("GasTopUpLimit", func() {
	BeforeEach(func() {
		BankAccount.MustTransfer(Backend, Controller.Address(), FinneyToWei(500))
	})

	When("the contract just has been deployed", func() {

		It("should have initial GasTopUp of 500 Finney available", func() {
			tl, err := Wallet.GasTopUpLimitValue(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(tl.String()).To(Equal(FinneyToWei(500).String()))

			tl, err = Wallet.GasTopUpLimitAvailable(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(tl.String()).To(Equal(FinneyToWei(500).String()))
		})

	})

	Describe("SetGasTopUpLimit", func() {
		var txSuccessful bool

		When("I try to set GasTopUp limit to one Gwei (below min GasTopUp limit)", func() {
			BeforeEach(func() {
				tx, err := Wallet.SetGasTopUpLimit(Owner.TransactOpts(ethertest.WithGasLimit(100000)), GweiToWei(1))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isGasExhausted(tx, 100000)).To(BeFalse())
				txSuccessful = isSuccessful(tx)
			})

			It("should fail", func() {
				Expect(txSuccessful).To(BeFalse())
			})

		})

		When("I try to set GasTopUp limit to one ETH (above max GasTopUp limit)", func() {
			BeforeEach(func() {
				tx, err := Wallet.SetGasTopUpLimit(Owner.TransactOpts(ethertest.WithGasLimit(100000)), EthToWei(1))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				ethertest.WithGasLimit(100000)
				txSuccessful = isSuccessful(tx)
			})

			It("should fail", func() {
				Expect(txSuccessful).To(BeFalse())
			})

		})

		When("I set GasTopUp limit for the first time to one Finney", func() {
			BeforeEach(func() {
				tx, err := Wallet.SetGasTopUpLimit(Owner.TransactOpts(ethertest.WithGasLimit(65000)), FinneyToWei(1))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				txSuccessful = isSuccessful(tx)
			})

			It("should succeed", func() {
				Expect(txSuccessful).To(BeTrue())
			})

			It("should update the set flag", func() {
				initialized, err := Wallet.GasTopUpLimitUpdateable(nil)
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

			When("I try to set the limit again", func() {
				It("should fail", func() {
					tx, err := Wallet.SetGasTopUpLimit(Owner.TransactOpts(ethertest.WithGasLimit(65000)), FinneyToWei(500))
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeFalse())
				})

			})
		})

	})

	Describe("Changing daily GasTopUp limit", func() {

		When("I submit daily GasTopUp limit above 1 Finney before initialization", func() {
			It("should fail", func() {
				tx, err := Wallet.SubmitGasTopUpLimitUpdate(Owner.TransactOpts(ethertest.WithGasLimit(65000)), FinneyToWei(5))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeFalse())
			})
		})

		When("I submit daily GasTopUp limit below 1 Finney after having set it", func() {
			It("should fail", func() {
				tx, err := Wallet.SetGasTopUpLimit(Owner.TransactOpts(ethertest.WithGasLimit(65000)), FinneyToWei(5))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())

				tx, err = Wallet.SubmitGasTopUpLimitUpdate(Owner.TransactOpts(ethertest.WithGasLimit(65000)), GweiToWei(1))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeFalse())
			})
		})

		When("I submit daily GasTopUp limit above 500 Finney after having set it", func() {
			It("should fail", func() {
				tx, err := Wallet.SetGasTopUpLimit(Owner.TransactOpts(ethertest.WithGasLimit(65000)), FinneyToWei(5))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())

				tx, err = Wallet.SubmitGasTopUpLimitUpdate(Owner.TransactOpts(ethertest.WithGasLimit(65000)), EthToWei(1))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeFalse())
			})
		})

		When("controller submits daily GasTopUp limit of 1 Finney after having set it", func() {
			It("should fail", func() {
				tx, err := Wallet.SetGasTopUpLimit(Owner.TransactOpts(), FinneyToWei(5))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())

				tx, err = Wallet.SubmitGasTopUpLimitUpdate(Controller.TransactOpts(ethertest.WithGasLimit(65000)), FinneyToWei(1))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeFalse())
			})

		})

		When("a random person submits daily GasTopUp limit of 1 Finney after having set it", func() {
			It("should fail", func() {
				tx, err := Wallet.SetGasTopUpLimit(Owner.TransactOpts(), FinneyToWei(5))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())

				tx, err = Wallet.SubmitGasTopUpLimitUpdate(RandomAccount.TransactOpts(ethertest.WithGasLimit(60000)), FinneyToWei(1))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeFalse())
			})

		})

		When("I submit GasTopUp limit of 2 and 1 Finney after having set it", func() {
			BeforeEach(func() {
				tx, err := Wallet.SetGasTopUpLimit(Owner.TransactOpts(ethertest.WithGasLimit(65000)), FinneyToWei(5))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())

				tx, err = Wallet.SubmitGasTopUpLimitUpdate(Owner.TransactOpts(ethertest.WithGasLimit(65000)), FinneyToWei(2))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())

				tx, err = Wallet.SubmitGasTopUpLimitUpdate(Owner.TransactOpts(ethertest.WithGasLimit(65000)), FinneyToWei(1))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})

			It("should emit 2 submission events", func() {
				it, err := Wallet.FilterSubmittedGasTopUpLimitUpdate(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(it.Next()).To(BeTrue())
				evt := it.Event
				Expect(it.Next()).To(BeTrue())
				Expect(evt.Amount).To(Equal(FinneyToWei(2)))
				evt = it.Event
				Expect(it.Next()).To(BeFalse())
				Expect(evt.Amount).To(Equal(FinneyToWei(1)))
			})

			It("should have pending GasTopUp limit of 1 Finney", func() {
				ptl, err := Wallet.GasTopUpLimitPending(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(ptl.String()).To(Equal(FinneyToWei(1).String()))
			})

			When("I try to submit a second GasTopUp limit of 500 Finney", func() {
				It("should succeed", func() {
					tx, err := Wallet.SubmitGasTopUpLimitUpdate(Owner.TransactOpts(), FinneyToWei(500))
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeTrue())
				})

			})

			When("I try to confirm the TopUpGas limit", func() {
				It("should fail", func() {
					tx, err := Wallet.ConfirmGasTopUpLimitUpdate(Owner.TransactOpts(ethertest.WithGasLimit(65000)), FinneyToWei(1))
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeFalse())
				})
			})

			When("a random person tries to confirm the GasTopUp limit", func() {
				It("should fail", func() {
					tx, err := Wallet.ConfirmGasTopUpLimitUpdate(RandomAccount.TransactOpts(ethertest.WithGasLimit(65000)), FinneyToWei(1))
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeFalse())
				})
			})

			When("the controller confirms the GasTopUp limit using the wrong amount", func() {
				It("should fail", func() {
					tx, err := Wallet.ConfirmGasTopUpLimitUpdate(Controller.TransactOpts(ethertest.WithGasLimit(100000)), FinneyToWei(2))
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeFalse())
				})
			})

			When("the controller confirms the GasTopUp limit", func() {
				BeforeEach(func() {
					tx, err := Wallet.ConfirmGasTopUpLimitUpdate(Controller.TransactOpts(), FinneyToWei(1))
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeTrue())
				})

				It("should have 1 Finney available for GasTopUps", func() {
					tl, err := Wallet.GasTopUpLimitAvailable(nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(tl.String()).To(Equal(FinneyToWei(1).String()))
				})

				When("I submit 2 GasTopUps 250 and 500 Finney repsectively", func() {
					BeforeEach(func() {
						tx, err := Wallet.SubmitGasTopUpLimitUpdate(Owner.TransactOpts(ethertest.WithGasLimit(65000)), FinneyToWei(250))
						Expect(err).ToNot(HaveOccurred())
						Backend.Commit()
						Expect(isSuccessful(tx)).To(BeTrue())

						tx, err = Wallet.SubmitGasTopUpLimitUpdate(Owner.TransactOpts(ethertest.WithGasLimit(65000)), FinneyToWei(500))
						Expect(err).ToNot(HaveOccurred())
						Backend.Commit()
						Expect(isSuccessful(tx)).To(BeTrue())
					})

					When("the controller confirms the last GasTopUp", func() {
						BeforeEach(func() {
							tx, err := Wallet.ConfirmGasTopUpLimitUpdate(Controller.TransactOpts(), FinneyToWei(500))
							Expect(err).ToNot(HaveOccurred())
							Backend.Commit()
							Expect(isSuccessful(tx)).To(BeTrue())
						})

						It("should have 1 Finney available for GasTopUps", func() {
							tl, err := Wallet.GasTopUpLimitAvailable(nil)
							Expect(err).ToNot(HaveOccurred())
							Expect(tl.String()).To(Equal(FinneyToWei(1).String()))
						})
						When("I wait for longer than a day", func() {
							BeforeEach(func() {
								Backend.AdjustTime(time.Hour * 25)
								Backend.Commit()
							})

							It("should have 500 Finney available for GasTopUps", func() {
								tl, err := Wallet.GasTopUpLimitAvailable(nil)
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
