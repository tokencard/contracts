package wallet_test

import (
	"time"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tokencard/contracts/v2/test/shared"
	"github.com/tokencard/ethertest"
)

var _ = FDescribe("Daily limit", func() {
	BeforeEach(func() {
		BankAccount.MustTransfer(Backend, Controller.Address(), EthToWei(1))
	})

	When("wallet has been deployed", func() {
		It("should have 10000 USD available", func() {
			av, err := Wallet.DailyLimitAvailable(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(av.String()).To(Equal(EthToWei(10000).String()))
		})

		When("a random person tries to submit a new daily limit", func() {
			It("should fail", func() {
				tx, err := Wallet.SubmitDailyLimitUpdate(RandomAccount.TransactOpts(ethertest.WithGasLimit(100000)), EthToWei(1))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isGasExhausted(tx, 100000)).To(BeFalse())
				Expect(isSuccessful(tx)).To(BeFalse())
			})
		})

		When("controller tries to submit a new daily limit", func() {
			It("should fail", func() {
				tx, err := Wallet.SubmitDailyLimitUpdate(Controller.TransactOpts(ethertest.WithGasLimit(100000)), EthToWei(1))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isGasExhausted(tx, 100000)).To(BeFalse())
				Expect(isSuccessful(tx)).To(BeFalse())
			})
		})

		When("Owner sets the DailyLimit to 1000 USD", func() {
			BeforeEach(func() {
				tx, err := Wallet.SubmitDailyLimitUpdate(Owner.TransactOpts(), EthToWei(1000))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})

			It("should emit a daily limit set event", func() {
				it, err := Wallet.FilterSetDailyLimit(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(it.Next()).To(BeTrue())
				evt := it.Event
				Expect(it.Next()).To(BeFalse())
				Expect(evt.Sender).To(Equal(Owner.Address()))
				Expect(evt.Amount).To(Equal(EthToWei(1000)))
			})

			It("should lower the available amount to 1000 USD", func() {
				av, err := Wallet.DailyLimitAvailable(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(av.String()).To(Equal(EthToWei(1000).String()))
			})

			It("should have a new limit of 1000 USD", func() {
				sl, err := Wallet.DailyLimitValue(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(sl.String()).To(Equal(EthToWei(1000).String()))
			})

		})

		When("I submit a new limit of 20000 USD and the controller confirms", func() {
			BeforeEach(func() {
				tx, err := Wallet.SubmitDailyLimitUpdate(Owner.TransactOpts(), EthToWei(20000))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())

                tx, err = Wallet.ConfirmDailyLimitUpdate(Controller.TransactOpts(), EthToWei(20000))
                Expect(err).ToNot(HaveOccurred())
                Backend.Commit()
                Expect(isSuccessful(tx)).To(BeTrue())
			})

			It("should keep the available amount at 10000 USD", func() {
				av, err := Wallet.DailyLimitAvailable(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(av.String()).To(Equal(EthToWei(10000).String()))
			})

			When("I wait for 24 hours", func() {
				BeforeEach(func() {
					Backend.AdjustTime(time.Hour*24 + time.Second)
					Backend.Commit()
				})

				It("should increase the available amount to 20000 USD", func() {
					av, err := Wallet.DailyLimitAvailable(nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(av.String()).To(Equal(EthToWei(20000).String()))
				})

				When("the contract has 1$", func() {
					BeforeEach(func() {
                        tx, err := Stablecoin.Credit(BankAccount.TransactOpts(), WalletAddress, EthToWei(1))
            			Expect(err).ToNot(HaveOccurred())
            			Backend.Commit()
            			Expect(isSuccessful(tx)).To(BeTrue())
					})

					When("I transfer 1$", func() {
						BeforeEach(func() {
							tx, err := Wallet.Transfer(Owner.TransactOpts(), RandomAccount.Address(), StablecoinAddress, EthToWei(1))
							Expect(err).ToNot(HaveOccurred())
							Backend.Commit()
							Expect(isSuccessful(tx)).To(BeTrue())
						})
						It("should have 9999 available", func() {
							sa, err := Wallet.DailyLimitAvailable(nil)
							Expect(err).ToNot(HaveOccurred())
							Expect(sa.String()).To(Equal(EthToWei(19999).String()))
						})
					})
				})
			})
		})

		When("I submit 3 limit updates", func() {
			BeforeEach(func() {
				tx, err := Wallet.SubmitDailyLimitUpdate(Owner.TransactOpts(), EthToWei(13000))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())

				tx, err = Wallet.SubmitDailyLimitUpdate(Owner.TransactOpts(), EthToWei(14000))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())

				tx, err = Wallet.SubmitDailyLimitUpdate(Owner.TransactOpts(), EthToWei(11000))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})

			It("should emit 3 submission events", func() {
				it, err := Wallet.FilterSubmittedDailyLimitUpdate(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(it.Next()).To(BeTrue())
				evt := it.Event
				Expect(it.Next()).To(BeTrue())
				Expect(evt.Amount).To(Equal(EthToWei(13000)))
				evt = it.Event
				Expect(it.Next()).To(BeTrue())
				Expect(evt.Amount).To(Equal(EthToWei(14000)))
				evt = it.Event
				Expect(it.Next()).To(BeFalse())
				Expect(evt.Amount).To(Equal(EthToWei(11000)))
			})

			It("should not reduce the available amount", func() {
				av, err := Wallet.DailyLimitAvailable(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(av.String()).To(Equal(EthToWei(10000).String()))
			})

			It("should set the pending limit to 11K $USD", func() {
				psl, err := Wallet.DailyLimitPending(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(psl.String()).To(Equal(EthToWei(11000).String()))
			})

			When("the controller confirms the limit providing the wrong amount", func() {
				It("should fail", func() {
					tx, err := Wallet.ConfirmDailyLimitUpdate(Controller.TransactOpts(ethertest.WithGasLimit(100000)), EthToWei(13000))
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeFalse())
				})
			})

			When("the controller confirms the right limit", func() {
				BeforeEach(func() {
					tx, err := Wallet.ConfirmDailyLimitUpdate(Controller.TransactOpts(), EthToWei(11000))
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeTrue())
				})

				It("should emit a limit set event", func() {
					it, err := Wallet.FilterSetDailyLimit(nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(it.Next()).To(BeTrue())
					evt := it.Event
					Expect(it.Next()).To(BeFalse())
					Expect(evt.Amount).To(Equal(EthToWei(11000)))
				})

				It("should keep the same available amount", func() {
					av, err := Wallet.DailyLimitAvailable(nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(av.String()).To(Equal(EthToWei(10000).String()))
				})
			})

			When("I try to update the limit again", func() {
				It("should succeed", func() {
					tx, err := Wallet.SubmitDailyLimitUpdate(Owner.TransactOpts(), EthToWei(1))
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeTrue())
				})
			})
		})
	})

})
