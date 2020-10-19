package wallet_test

import (
    "math/big"
	"time"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tokencard/contracts/v2/test/shared"
	"github.com/tokencard/ethertest"
)

var _ = Describe("DailyLimit", func() {

	BeforeEach(func() {
		BankAccount.MustTransfer(Backend, Controller.Address(), EthToWei(1))
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

        It("should emit a daily limit set event", func() {
            it, err := Wallet.FilterInitializedDailyLimit(nil)
            Expect(err).ToNot(HaveOccurred())
            Expect(it.Next()).To(BeTrue())
            evt := it.Event
            Expect(it.Next()).To(BeFalse())
            Expect(evt.Amount.String()).To(Equal(EthToWei(10000).String()))
            initTime := Backend.Blockchain().CurrentHeader().Time + 24*60*60 - 10
            Expect(evt.NextReset.String()).To(Equal(big.NewInt(int64(initTime)).String()))
        })
	})

    When("the controller tries to submit a new daily limit", func() {
        It("should fail", func() {
            tx, err := Wallet.SubmitDailyLimitUpdate(Controller.TransactOpts(ethertest.WithGasLimit(100000)), EthToWei(1))
            Expect(err).ToNot(HaveOccurred())
            Backend.Commit()
            Expect(isSuccessful(tx)).To(BeFalse())
        })
    })

    When("the controller tries to confirm with amount=0", func() {
        It("should fail", func() {
            tx, err := Wallet.ConfirmDailyLimitUpdate(Controller.TransactOpts(ethertest.WithGasLimit(100000)), big.NewInt(0))
            Expect(err).ToNot(HaveOccurred())
            Backend.Commit()
            Expect(isSuccessful(tx)).To(BeFalse())
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

	Describe("Changing daily limit", func() {

		When("I submit daily limit of 5 Finney", func() {
			It("should succeed", func() {
				tx, err := Wallet.SubmitDailyLimitUpdate(Owner.TransactOpts(), FinneyToWei(5))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
                Expect(isSuccessful(tx)).To(BeTrue())
			})
		})


        When("Owner sets the DailyLimit to 1000 $USD", func() {
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

			It("should lower the available amount to 1000 $USD", func() {
				av, err := Wallet.DailyLimitAvailable(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(av.String()).To(Equal(EthToWei(1000).String()))
			})

			It("should have a new limit of 1000 $USD", func() {
				sl, err := Wallet.DailyLimitValue(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(sl.String()).To(Equal(EthToWei(1000).String()))
			})

            It("should have a pending limit of 1000 $USD", func() {
				ptl, err := Wallet.DailyLimitPending(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(ptl.String()).To(Equal(EthToWei(1000).String()))
			})

            When("the controller tries to confirm with amount=0", func() {
                It("should fail", func() {
                    tx, err := Wallet.ConfirmDailyLimitUpdate(Controller.TransactOpts(ethertest.WithGasLimit(100000)), big.NewInt(0))
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeFalse())
                })
            })

            When("I transfer 1$", func() {
                BeforeEach(func() {

                    tx, err := Stablecoin.Credit(BankAccount.TransactOpts(), WalletAddress, EthToWei(1))
        			Expect(err).ToNot(HaveOccurred())
        			Backend.Commit()
        			Expect(isSuccessful(tx)).To(BeTrue())

                    tx, err = Wallet.Transfer(Owner.TransactOpts(), RandomAccount.Address(), StablecoinAddress, EthToWei(1))
                    Expect(err).ToNot(HaveOccurred())
                    Backend.Commit()
                    Expect(isSuccessful(tx)).To(BeTrue())
                })

                It("should have 9999 available", func() {
                    sa, err := Wallet.DailyLimitAvailable(nil)
                    Expect(err).ToNot(HaveOccurred())
                    Expect(sa.String()).To(Equal(EthToWei(999).String()))
                })

                When("the controller tries to confirm with amount=1000$", func() {
                    BeforeEach(func() {
                        tx, err := Wallet.ConfirmDailyLimitUpdate(Controller.TransactOpts(), EthToWei(1000))
    					Expect(err).ToNot(HaveOccurred())
    					Backend.Commit()
    					Expect(isSuccessful(tx)).To(BeTrue())
                    })

                    It("should not modify the avaiable amount", func() {
        				av, err := Wallet.DailyLimitAvailable(nil)
        				Expect(err).ToNot(HaveOccurred())
        				Expect(av.String()).To(Equal(EthToWei(999).String()))
        			})

        			It("should not modify the limit", func() {
        				sl, err := Wallet.DailyLimitValue(nil)
        				Expect(err).ToNot(HaveOccurred())
        				Expect(sl.String()).To(Equal(EthToWei(1000).String()))
        			})

                    It("should not modify the pending limit", func() {
        				ptl, err := Wallet.DailyLimitPending(nil)
        				Expect(err).ToNot(HaveOccurred())
        				Expect(ptl.String()).To(Equal(EthToWei(1000).String()))
        			})
                })
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
					Expect(isSuccessful(tx)).To(BeFalse())
				})
			})

			When("the controller confirms the limit using the wrong amount", func() {
				It("should fail", func() {
					tx, err := Wallet.ConfirmDailyLimitUpdate(Controller.TransactOpts(ethertest.WithGasLimit(100000)), EthToWei(12000))
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeFalse())
				})
			})

			When("the controller confirms with the right amount", func() {
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

                        When("I wait for 24 hours", func() {
            				BeforeEach(func() {
            					Backend.AdjustTime(time.Hour*24 + time.Second)
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
