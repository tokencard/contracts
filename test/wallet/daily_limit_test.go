package wallet_test

import (
	"math/big"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tokencard/contracts/v3/test/shared"
	"github.com/tokencard/ethertest"
)

var _ = Describe("Daily limit", func() {

	BeforeEach(func() {
		BankAccount.MustTransfer(Backend, Controller.Address(), EthToWei(1))
	})

	When("the contract just has been deployed", func() {

		It("should have initial daily limit of 10000$", func() {
			ll, err := WalletProxy.DailyLimitValue(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(ll.String()).To(Equal(EthToWei(10000).String()))

			ll, err = WalletProxy.DailyLimitAvailable(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(ll.String()).To(Equal(EthToWei(10000).String()))
		})

		It("should emit InitializedDailyLimit event", func() {
			it, err := WalletProxy.FilterInitializedDailyLimit(nil)
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
			tx, err := WalletProxy.SubmitDailyLimitUpdate(Controller.TransactOpts(ethertest.WithGasLimit(100000)), EthToWei(1))
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeFalse())
		})
	})

	When("the controller tries to confirm with amount=0", func() {
		It("should fail", func() {
			tx, err := WalletProxy.ConfirmDailyLimitUpdate(Controller.TransactOpts(ethertest.WithGasLimit(100000)), big.NewInt(0))
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeFalse())
		})
	})

	When("a random person submits a daily limit of 1 Finney", func() {
		It("should fail", func() {
			tx, err := WalletProxy.SubmitDailyLimitUpdate(RandomAccount.TransactOpts(ethertest.WithGasLimit(65000)), FinneyToWei(1))
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeFalse())
		})
	})

	Describe("Changing daily limit", func() {

		When("Owner sets the DailyLimit to 1000 $USD", func() {
			BeforeEach(func() {
				tx, err := WalletProxy.SubmitDailyLimitUpdate(Owner.TransactOpts(), EthToWei(1000))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})

			It("should emit a UpdatedAvailableDailyLimit event", func() {
				it, err := WalletProxy.FilterUpdatedAvailableDailyLimit(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(it.Next()).To(BeTrue())
				evt := it.Event
				Expect(it.Next()).To(BeFalse())
				Expect(evt.Amount.String()).To(Equal(EthToWei(1000).String()))
				initTime := Backend.Blockchain().CurrentHeader().Time + 24*60*60 - 20
				Expect(evt.NextReset.String()).To(Equal(big.NewInt(int64(initTime)).String()))
			})

			It("should emit a daily limit set event", func() {
				it, err := WalletProxy.FilterSetDailyLimit(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(it.Next()).To(BeTrue())
				evt := it.Event
				Expect(it.Next()).To(BeFalse())
				Expect(evt.Sender).To(Equal(Owner.Address()))
				Expect(evt.Amount).To(Equal(EthToWei(1000)))
			})

			It("should lower the available amount to 1000 $USD", func() {
				av, err := WalletProxy.DailyLimitAvailable(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(av.String()).To(Equal(EthToWei(1000).String()))
			})

			It("should have a new limit of 1000 $USD", func() {
				sl, err := WalletProxy.DailyLimitValue(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(sl.String()).To(Equal(EthToWei(1000).String()))
			})

			It("should have a pending limit of 1000 $USD", func() {
				ptl, err := WalletProxy.DailyLimitPending(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(ptl.String()).To(Equal(EthToWei(1000).String()))
			})

			When("the controller tries to confirm with amount=0", func() {
				It("should fail", func() {
					tx, err := WalletProxy.ConfirmDailyLimitUpdate(Controller.TransactOpts(ethertest.WithGasLimit(100000)), big.NewInt(0))
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeFalse())
				})
			})

			When("I transfer 1$", func() {
				BeforeEach(func() {

					tx, err := Stablecoin.Credit(BankAccount.TransactOpts(), WalletProxyAddress, EthToWei(1))
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeTrue())

					tx, err = WalletProxy.Transfer(Owner.TransactOpts(), RandomAccount.Address(), StablecoinAddress, EthToWei(1))
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeTrue())
				})

				It("should emit a UpdatedAvailableDailyLimit event", func() {
					it, err := WalletProxy.FilterUpdatedAvailableDailyLimit(nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(it.Next()).To(BeTrue())
					Expect(it.Next()).To(BeTrue())
					evt := it.Event
					Expect(it.Next()).To(BeFalse())
					Expect(evt.Amount.String()).To(Equal(EthToWei(999).String()))
					initTime := Backend.Blockchain().CurrentHeader().Time + 24*60*60 - 40
					Expect(evt.NextReset.String()).To(Equal(big.NewInt(int64(initTime)).String()))
				})

				It("should have 9999 available", func() {
					sa, err := WalletProxy.DailyLimitAvailable(nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(sa.String()).To(Equal(EthToWei(999).String()))
				})

				When("the controller tries to confirm with amount=1000$", func() {
					It("should fail", func() {
						tx, err := WalletProxy.ConfirmDailyLimitUpdate(Controller.TransactOpts(ethertest.WithGasLimit(65000)), EthToWei(1000))
						Expect(err).ToNot(HaveOccurred())
						Backend.Commit()
						Expect(isSuccessful(tx)).To(BeFalse())
						returnData, _ := ethCall(tx)
						Expect(string(returnData[len(returnData)-64:])).To(ContainSubstring("limit should be greater than current one"))
					})
				})
			})
		})

		When("I submit 2 Dailylimits of 12K and 11K USD respectively", func() {
			BeforeEach(func() {
				tx, err := WalletProxy.SubmitDailyLimitUpdate(Owner.TransactOpts(), EthToWei(12000))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())

				tx, err = WalletProxy.SubmitDailyLimitUpdate(Owner.TransactOpts(), EthToWei(11000))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})

			It("should emit 2 submission events", func() {
				it, err := WalletProxy.FilterSubmittedDailyLimitUpdate(nil)
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
				ptl, err := WalletProxy.DailyLimitPending(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(ptl.String()).To(Equal(EthToWei(11000).String()))
			})

			When("the Owner tries to re-submit a limit of 13K $USD", func() {
				BeforeEach(func() {
					tx, err := WalletProxy.SubmitDailyLimitUpdate(Owner.TransactOpts(), EthToWei(13000))
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeTrue())
				})

				It("should set pending limit to 13K $USD", func() {
					psl, err := WalletProxy.DailyLimitPending(nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(psl.String()).To(Equal(EthToWei(13000).String()))
				})
			})

			When("the owner tries to confirm the daily limit", func() {
				It("should fail", func() {
					tx, err := WalletProxy.ConfirmDailyLimitUpdate(Owner.TransactOpts(ethertest.WithGasLimit(65000)), EthToWei(11000))
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeFalse())
				})
			})

			When("a random person tries to confirm the daily limit", func() {
				It("should fail", func() {
					tx, err := WalletProxy.ConfirmDailyLimitUpdate(RandomAccount.TransactOpts(ethertest.WithGasLimit(65000)), EthToWei(11000))
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isGasExhausted(tx, 65000)).To(BeFalse())
					Expect(isSuccessful(tx)).To(BeFalse())
				})
			})

			When("the controller confirms the limit using the wrong amount", func() {
				It("should fail", func() {
					tx, err := WalletProxy.ConfirmDailyLimitUpdate(Controller.TransactOpts(ethertest.WithGasLimit(100000)), EthToWei(12000))
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isGasExhausted(tx, 100000)).To(BeFalse())
					Expect(isSuccessful(tx)).To(BeFalse())
					returnData, _ := ethCall(tx)
					Expect(string(returnData[len(returnData)-64:])).To(ContainSubstring("confirmed or submitted limit mismatch"))
				})
			})

			When("the controller confirms the limit with the right amount", func() {
				BeforeEach(func() {
					tx, err := WalletProxy.ConfirmDailyLimitUpdate(Controller.TransactOpts(), EthToWei(11000))
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeTrue())
				})

				It("should have 11K USD available", func() {
					ll, err := WalletProxy.DailyLimitAvailable(nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(ll.String()).To(Equal(EthToWei(11000).String()))
				})

				It("should emit a UpdatedAvailableDailyLimit event", func() {
					it, err := WalletProxy.FilterUpdatedAvailableDailyLimit(nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(it.Next()).To(BeTrue())
					evt := it.Event
					Expect(it.Next()).To(BeFalse())
					Expect(evt.Amount.String()).To(Equal(EthToWei(11000).String()))
					initTime := Backend.Blockchain().CurrentHeader().Time + 24*60*60
					Expect(evt.NextReset.String()).To(Equal(big.NewInt(int64(initTime)).String()))
				})

				When("the controller tries to re-confirm", func() {
					It("should fail", func() {
						tx, err := WalletProxy.ConfirmDailyLimitUpdate(Controller.TransactOpts(ethertest.WithGasLimit(65000)), EthToWei(11000))
						Expect(err).ToNot(HaveOccurred())
						Backend.Commit()
						Expect(isSuccessful(tx)).To(BeFalse())
						returnData, _ := ethCall(tx)
						Expect(string(returnData[len(returnData)-64:])).To(ContainSubstring("limit should be greater than current one"))
					})
				})

				When("I submit a second limit of 12k USD", func() {
					BeforeEach(func() {
						tx, err := WalletProxy.SubmitDailyLimitUpdate(Owner.TransactOpts(), EthToWei(12000))
						Expect(err).ToNot(HaveOccurred())
						Backend.Commit()
						Expect(isSuccessful(tx)).To(BeTrue())
					})

					When("the controller confirms the submitted limit", func() {
						BeforeEach(func() {
							tx, err := WalletProxy.ConfirmDailyLimitUpdate(Controller.TransactOpts(), EthToWei(12000))
							Expect(err).ToNot(HaveOccurred())
							Backend.Commit()
							Expect(isSuccessful(tx)).To(BeTrue())
						})

						It("should have 12k USD available for spending", func() {
							tl, err := WalletProxy.DailyLimitAvailable(nil)
							Expect(err).ToNot(HaveOccurred())
							Expect(tl.String()).To(Equal(EthToWei(12000).String()))
						})

						It("should emit a UpdatedAvailableDailyLimit event", func() {
							it, err := WalletProxy.FilterUpdatedAvailableDailyLimit(nil)
							Expect(err).ToNot(HaveOccurred())
							Expect(it.Next()).To(BeTrue())
							Expect(it.Next()).To(BeTrue())
							evt := it.Event
							Expect(it.Next()).To(BeFalse())
							Expect(evt.Amount.String()).To(Equal(EthToWei(12000).String()))
							initTime := Backend.Blockchain().CurrentHeader().Time + 24*60*60
							Expect(evt.NextReset.String()).To(Equal(big.NewInt(int64(initTime)).String()))
						})

						When("I wait for 24 hours", func() {
							BeforeEach(func() {
								Backend.AdjustTime(time.Hour*24 + time.Second)
								Backend.Commit()
							})

							It("should have 12k USD available for spending", func() {
								ll, err := WalletProxy.DailyLimitAvailable(nil)
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
