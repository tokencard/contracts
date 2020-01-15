package wallet_test

import (
	"time"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tokencard/contracts/v2/test/shared"
	"github.com/tokencard/ethertest"
)

var _ = Describe("DailyLimit", func() {

	BeforeEach(func() {
		BankAccount.MustTransfer(Backend, Controller.Address(), FinneyToWei(500))
	})

	When("the contract just has been deployed", func() {

		It("should have initial daily limit of 10000$", func() {
			ll, err := Wallet.DailyLimitValue(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(ll.String()).To(Equal(EthToWei(10000).String()))

			ll, err = Wallet.DailyLimitAvailable(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(ll.String()).To(Equal(EthToWei(10000).String()))
		})

	})

	Describe("Changing daily limit", func() {

		When("I submit daily limit of 5 Finney", func() {
			It("should succeed", func() {
				tx, err := Wallet.SubmitDailyLimitUpdate(Owner.TransactOpts(), FinneyToWei(5))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
                Expect(isSuccessful(tx)).To(BeTrue())
			})
		})

		When("a random person submits a daily limit of 1 Finney", func() {
			It("should fail", func() {
				tx, err := Wallet.SubmitDailyLimitUpdate(RandomAccount.TransactOpts(ethertest.WithGasLimit(65000)), FinneyToWei(1))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeFalse())
			})
		})

		When("I submit 2 Dailylimits of 12K and 11K USD respectively", func() {
			BeforeEach(func() {
				tx, err := Wallet.SubmitDailyLimitUpdate(Owner.TransactOpts(), EthToWei(12000))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())

				tx, err = Wallet.SubmitDailyLimitUpdate(Owner.TransactOpts(), EthToWei(11000))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})

			It("should emit 2 submission events", func() {
				it, err := Wallet.FilterSubmittedDailyLimitUpdate(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(it.Next()).To(BeTrue())
				evt := it.Event
				Expect(it.Next()).To(BeTrue())
				Expect(evt.Amount).To(Equal(EthToWei(12000)))
				evt = it.Event
				Expect(it.Next()).To(BeFalse())
				Expect(evt.Amount).To(Equal(EthToWei(11000)))
			})

			It("should have a pending limit of 11K $", func() {
				ptl, err := Wallet.DailyLimitPending(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(ptl.String()).To(Equal(EthToWei(11000).String()))
			})

			When("the Owner tries to re-submit a limit of 13K $USD", func() {
				BeforeEach(func() {
					tx, err := Wallet.SubmitDailyLimitUpdate(Owner.TransactOpts(), EthToWei(13000))
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeTrue())
				})

				It("should set pending limit to 13K $USD", func() {
					psl, err := Wallet.DailyLimitPending(nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(psl.String()).To(Equal(EthToWei(13000).String()))
				})
			})


			When("the owner tries to confirm the daily limit", func() {
				It("should fail", func() {
					tx, err := Wallet.ConfirmDailyLimitUpdate(Owner.TransactOpts(ethertest.WithGasLimit(65000)), EthToWei(11000))
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeFalse())
				})
			})

			When("a random person tries to confirm the daily limit", func() {
				It("should fail", func() {
					tx, err := Wallet.ConfirmDailyLimitUpdate(RandomAccount.TransactOpts(ethertest.WithGasLimit(65000)), EthToWei(11000))
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isGasExhausted(tx, 65000)).To(BeFalse())
					Expect(isSuccessful(tx)).To(BeFalse())
				})
			})

			When("the controller confirms the limit using the wrong amount", func() {
				It("should fail", func() {
					tx, err := Wallet.ConfirmDailyLimitUpdate(Controller.TransactOpts(ethertest.WithGasLimit(100000)), EthToWei(12000))
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isGasExhausted(tx, 100000)).To(BeFalse())
					Expect(isSuccessful(tx)).To(BeFalse())
				})
			})

			When("the controller confirms the limit", func() {
				BeforeEach(func() {
					tx, err := Wallet.ConfirmDailyLimitUpdate(Controller.TransactOpts(), EthToWei(11000))
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeTrue())
				})

				It("should have 10K $USD available", func() {
					ll, err := Wallet.DailyLimitAvailable(nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(ll.String()).To(Equal(EthToWei(10000).String()))
				})

				When("I submit a second limit of 12K $USD", func() {
					BeforeEach(func() {
						tx, err := Wallet.SubmitDailyLimitUpdate(Owner.TransactOpts(), EthToWei(12000))
						Expect(err).ToNot(HaveOccurred())
						Backend.Commit()
						Expect(isSuccessful(tx)).To(BeTrue())
					})

					When("the controller confirms the submitted limit", func() {
						BeforeEach(func() {
							tx, err := Wallet.ConfirmDailyLimitUpdate(Controller.TransactOpts(), EthToWei(12000))
							Expect(err).ToNot(HaveOccurred())
							Backend.Commit()
							Expect(isSuccessful(tx)).To(BeTrue())
						})

						It("should have 10K $USD available for spending", func() {
							tl, err := Wallet.DailyLimitAvailable(nil)
							Expect(err).ToNot(HaveOccurred())
							Expect(tl.String()).To(Equal(EthToWei(10000).String()))
						})

						When("I wait for longer than a day", func() {
							BeforeEach(func() {
								Backend.AdjustTime(time.Hour * 25)
								Backend.Commit()
							})

							It("should have 12K $USD available for spending", func() {
								ll, err := Wallet.DailyLimitAvailable(nil)
								Expect(err).ToNot(HaveOccurred())
								Expect(ll.String()).To(Equal(EthToWei(12000).String()))
							})

						})

					})

				})

			})

		})

	})

})
