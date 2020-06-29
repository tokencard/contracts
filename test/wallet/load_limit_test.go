package wallet_test

import (
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tokencard/contracts/v3/test/shared"
	"github.com/tokencard/ethertest"
)

var _ = Describe("loadLimit", func() {

	BeforeEach(func() {
		BankAccount.MustTransfer(Backend, Controller.Address(), FinneyToWei(500))

	})

	When("the contract just has been deployed", func() {

		It("should have initial MAX load limit of 10,000 USD /stablecoins", func() {
			ll, err := WalletProxy.LoadLimitValue(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(ll.String()).To(Equal(GweiToWei(10).String()))

			ll, err = WalletProxy.LoadLimitAvailable(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(ll.String()).To(Equal(GweiToWei(10).String()))
		})

	})

	Describe("SetLoadLimit", func() {

		When("I try to set load limit to 0 (below min load limit)", func() {
			It("Should succeed", func() {
				tx, err := WalletProxy.SetLoadLimit(Owner.TransactOpts(), GweiToWei(0))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})

		})

		When("I try to Set load limit to 10001 USD (above max load limit)", func() {
			It("Should fail", func() {
				tx, err := WalletProxy.SetLoadLimit(Owner.TransactOpts(ethertest.WithGasLimit(100000)), EthToWei(10001))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isGasExhausted(tx, 100000)).To(BeFalse())
				Expect(isSuccessful(tx)).To(BeFalse())
			})
		})

		When("I Set load limit for the first time to 1000 USD", func() {
			var tx *types.Transaction
			BeforeEach(func() {
				var err error
				tx, err = WalletProxy.SetLoadLimit(Owner.TransactOpts(), GweiToWei(1))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
			})

			It("should succeed", func() {
				Expect(isSuccessful(tx)).To(BeTrue())
			})

			It("should update the set flag", func() {
				set, err := WalletProxy.LoadLimitControllerConfirmationRequired(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(set).To(BeTrue())
			})

			It("should emit a load limit set event", func() {
				it, err := WalletProxy.FilterSetLoadLimit(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(it.Next()).To(BeTrue())
				evt := it.Event
				Expect(it.Next()).To(BeFalse())
				Expect(evt.Sender).To(Equal(Owner.Address()))
				Expect(evt.Amount.String()).To(Equal(GweiToWei(1).String()))
			})

			When("I try to set the limit again", func() {
				It("should fail", func() {
					tx, err := WalletProxy.SetLoadLimit(Owner.TransactOpts(ethertest.WithGasLimit(100000)), GweiToWei(5))
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isGasExhausted(tx, 100000)).To(BeFalse())
					Expect(isSuccessful(tx)).To(BeFalse())
				})

			})

		})

	})

	Describe("Changing daily Load limit", func() {

		When("I submit daily load limit of 5K USD before setting it", func() {
			It("should fail", func() {
				tx, err := WalletProxy.SubmitLoadLimitUpdate(Owner.TransactOpts(ethertest.WithGasLimit(100000)), GweiToWei(5))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isGasExhausted(tx, 100000)).To(BeFalse())
				Expect(isSuccessful(tx)).To(BeFalse())
			})
		})

		When("I submit daily load limit of 0 (< min load limit) after having set it", func() {
			It("should succeed", func() {
				tx, err := WalletProxy.SetLoadLimit(Owner.TransactOpts(), GweiToWei(10))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())

				tx, err = WalletProxy.SubmitLoadLimitUpdate(Owner.TransactOpts(), GweiToWei(0))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})
		})

		When("I submit daily load limit of 10001 USD (> max load limit) after having set it", func() {
			It("should fail", func() {
				tx, err := WalletProxy.SetLoadLimit(Owner.TransactOpts(), GweiToWei(5))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())

				tx, err = WalletProxy.SubmitLoadLimitUpdate(Owner.TransactOpts(ethertest.WithGasLimit(65000)), GweiToWei(10001))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isGasExhausted(tx, 100000)).To(BeFalse())
				Expect(isSuccessful(tx)).To(BeFalse())
			})
		})

		When("controller submits daily load limit of 1K USD(> max load limit) after having set it", func() {
			It("should fail", func() {
				tx, err := WalletProxy.SetLoadLimit(Owner.TransactOpts(), GweiToWei(5))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())

				tx, err = WalletProxy.SubmitLoadLimitUpdate(Controller.TransactOpts(ethertest.WithGasLimit(65000)), GweiToWei(1))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isGasExhausted(tx, 65000)).To(BeFalse())
				Expect(isSuccessful(tx)).To(BeFalse())
			})

		})

		When("a random person submits daily load limit of 0 USD after having set it", func() {
			It("should fail", func() {
				tx, err := WalletProxy.SetLoadLimit(Owner.TransactOpts(), GweiToWei(5))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())

				tx, err = WalletProxy.SubmitLoadLimitUpdate(RandomAccount.TransactOpts(ethertest.WithGasLimit(65000)), big.NewInt(0))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isGasExhausted(tx, 65000)).To(BeFalse())
				Expect(isSuccessful(tx)).To(BeFalse())
			})

		})

		When("I submit 2 Loadlimits of 2K and 1K USD respectively after having set it", func() {
			BeforeEach(func() {
				tx, err := WalletProxy.SetLoadLimit(Owner.TransactOpts(), GweiToWei(5))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())

				tx, err = WalletProxy.SubmitLoadLimitUpdate(Owner.TransactOpts(), GweiToWei(2))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())

				tx, err = WalletProxy.SubmitLoadLimitUpdate(Owner.TransactOpts(), GweiToWei(1))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})

			It("should emit 2 submission events", func() {
				it, err := WalletProxy.FilterSubmittedLoadLimitUpdate(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(it.Next()).To(BeTrue())
				evt := it.Event
				Expect(it.Next()).To(BeTrue())
				Expect(evt.Amount).To(Equal(GweiToWei(2)))
				evt = it.Event
				Expect(it.Next()).To(BeFalse())
				Expect(evt.Amount).To(Equal(GweiToWei(1)))
			})

			It("should have pending load limit of 1K USD", func() {
				ptl, err := WalletProxy.LoadLimitPending(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(ptl.String()).To(Equal(GweiToWei(1).String()))
			})

			When("the Owner tries to re-submit a load limit of 5K USD", func() {
				BeforeEach(func() {
					tx, err := WalletProxy.SubmitLoadLimitUpdate(Owner.TransactOpts(), GweiToWei(5))
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeTrue())
				})

				It("should set pending load limit to 5K", func() {
					psl, err := WalletProxy.LoadLimitPending(nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(psl.String()).To(Equal(GweiToWei(5).String()))
				})

				When("the Owner tries to submit a second load limit of 5K USD", func() {
					It("should succeed", func() {
						tx, err := WalletProxy.SubmitLoadLimitUpdate(Owner.TransactOpts(ethertest.WithGasLimit(65000)), GweiToWei(5))
						Expect(err).ToNot(HaveOccurred())
						Backend.Commit()
						Expect(isSuccessful(tx)).To(BeTrue())
					})
				})
			})

			When("I try to confirm the load limit", func() {
				It("should fail", func() {
					tx, err := WalletProxy.ConfirmLoadLimitUpdate(Owner.TransactOpts(ethertest.WithGasLimit(65000)), GweiToWei(1))
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeFalse())
				})
			})

			When("a random person tries to confirm the load limit", func() {
				It("should fail", func() {
					tx, err := WalletProxy.ConfirmLoadLimitUpdate(RandomAccount.TransactOpts(ethertest.WithGasLimit(65000)), GweiToWei(1))
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isGasExhausted(tx, 65000)).To(BeFalse())
					Expect(isSuccessful(tx)).To(BeFalse())
				})
			})

			When("the controller confirms the load limit using the wrong amount", func() {
				It("should fail", func() {
					tx, err := WalletProxy.ConfirmLoadLimitUpdate(Controller.TransactOpts(ethertest.WithGasLimit(100000)), GweiToWei(2))
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isGasExhausted(tx, 100000)).To(BeFalse())
					Expect(isSuccessful(tx)).To(BeFalse())
				})
			})

			When("the controller confirms the load limit", func() {
				BeforeEach(func() {
					tx, err := WalletProxy.ConfirmLoadLimitUpdate(Controller.TransactOpts(), GweiToWei(1))
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeTrue())
				})

				It("should have 1K USD available for loading", func() {
					ll, err := WalletProxy.LoadLimitAvailable(nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(ll.String()).To(Equal(GweiToWei(1).String()))
				})

				When("I submit a second load limit to 5K USD", func() {
					BeforeEach(func() {
						tx, err := WalletProxy.SubmitLoadLimitUpdate(Owner.TransactOpts(), GweiToWei(5))
						Expect(err).ToNot(HaveOccurred())
						Backend.Commit()
						Expect(isSuccessful(tx)).To(BeTrue())
					})

					When("the controller confirms the load limit", func() {
						BeforeEach(func() {
							tx, err := WalletProxy.ConfirmLoadLimitUpdate(Controller.TransactOpts(), GweiToWei(5))
							Expect(err).ToNot(HaveOccurred())
							Backend.Commit()
							Expect(isSuccessful(tx)).To(BeTrue())
						})

						It("should have 1K USD available for loading", func() {
							tl, err := WalletProxy.LoadLimitAvailable(nil)
							Expect(err).ToNot(HaveOccurred())
							Expect(tl.String()).To(Equal(GweiToWei(1).String()))
						})
						When("I wait for longer than a day", func() {
							BeforeEach(func() {
								Backend.AdjustTime(time.Hour * 25)
								Backend.Commit()
							})

							It("should have 5K USD available for further loading", func() {
								ll, err := WalletProxy.LoadLimitAvailable(nil)
								Expect(err).ToNot(HaveOccurred())
								Expect(ll.String()).To(Equal(GweiToWei(5).String()))
							})

						})

					})

				})

			})

		})

	})

})
