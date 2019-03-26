package wallet_test

import (
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tokencard/contracts/test/shared"
	"github.com/tokencard/ethertest"
)

var _ = Describe("loadLimit", func() {

	BeforeEach(func() {
		BankAccount.MustTransfer(Backend, Controller.Address(), FinneyToWei(500))
	})

	When("the contract just has been deployed", func() {

		It("should have initial load limit of 10,000 USD /stablecoins", func() {
			ll, err := Wallet.LoadLimit(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(ll.String()).To(Equal(EthToWei(10000).String()))

			Wallet.Balance(nil, common.HexToAddress("0x0"))

			ll, err = Wallet.LoadAvailable(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(ll.String()).To(Equal(EthToWei(10000).String()))
		})

	})

	Describe("InitializeLoadLimit", func() {

		When("I try to initialize load limit to one Gwei (below min load limit)", func() {
			It("Should fail", func() {
				tx, err := Wallet.InitializeLoadLimit(Owner.TransactOpts(ethertest.WithGasLimit(100000)), GweiToWei(1))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isGasExhausted(tx, 100000)).To(BeFalse())
				Expect(isSuccessful(tx)).To(BeFalse())
			})

		})

		When("I try to initialize load limit to 10001 USD (above max load limit)", func() {
			It("Should fail", func() {
				tx, err := Wallet.InitializeLoadLimit(Owner.TransactOpts(ethertest.WithGasLimit(100000)), EthToWei(10001))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isGasExhausted(tx, 100000)).To(BeFalse())
				Expect(isSuccessful(tx)).To(BeFalse())
			})
		})

		When("I initialize load limit for the first time to one Finney", func() {
			var tx *types.Transaction
			BeforeEach(func() {
				var err error
				tx, err = Wallet.InitializeLoadLimit(Owner.TransactOpts(), FinneyToWei(1))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
			})

			It("should succeed", func() {
				Expect(isSuccessful(tx)).To(BeTrue())
			})

			It("should update the initialization flag", func() {
				initialized, err := Wallet.InitializedLoadLimit(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(initialized).To(BeTrue())
			})

			It("should emit a load limit set event", func() {
				it, err := Wallet.FilterSetLoadLimit(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(it.Next()).To(BeTrue())
				evt := it.Event
				Expect(it.Next()).To(BeFalse())
				Expect(evt.Sender).To(Equal(Owner.Address()))
				Expect(evt.Amount).To(Equal(FinneyToWei(1)))
			})

			When("I try to initialize the limit again", func() {
				It("should fail", func() {
					tx, err := Wallet.InitializeLoadLimit(Owner.TransactOpts(ethertest.WithGasLimit(100000)), FinneyToWei(500))
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isGasExhausted(tx, 100000)).To(BeFalse())
					Expect(isSuccessful(tx)).To(BeFalse())
				})

			})

		})

	})

	Describe("Changing daily Load limit", func() {

		When("I submit daily load limit of 5 Finney before initialization", func() {
			It("should fail", func() {
				tx, err := Wallet.SubmitLoadLimitUpdate(Owner.TransactOpts(ethertest.WithGasLimit(100000)), FinneyToWei(5))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isGasExhausted(tx, 100000)).To(BeFalse())
				Expect(isSuccessful(tx)).To(BeFalse())
			})
		})

		When("I submit daily load limit of 1 GWei (< min load limit) after initialization", func() {
			It("should fail", func() {
				tx, err := Wallet.InitializeLoadLimit(Owner.TransactOpts(ethertest.WithGasLimit(100000)), FinneyToWei(5))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isGasExhausted(tx, 100000)).To(BeFalse())
				Expect(isSuccessful(tx)).To(BeTrue())

				tx, err = Wallet.SubmitLoadLimitUpdate(Owner.TransactOpts(ethertest.WithGasLimit(100000)), GweiToWei(1))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isGasExhausted(tx, 100000)).To(BeFalse())
				Expect(isSuccessful(tx)).To(BeFalse())
			})
		})

		When("I submit daily load limit of 10002 eth (> max load eth) Finney after initialization", func() {
			It("should fail", func() {
				tx, err := Wallet.InitializeLoadLimit(Owner.TransactOpts(), FinneyToWei(5))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())

				tx, err = Wallet.SubmitLoadLimitUpdate(Owner.TransactOpts(ethertest.WithGasLimit(65000)), EthToWei(10002))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isGasExhausted(tx, 100000)).To(BeFalse())
				Expect(isSuccessful(tx)).To(BeFalse())
			})
		})

		When("controller submits daily load limit of 1 Finney after initialization", func() {
			It("should fail", func() {
				tx, err := Wallet.InitializeLoadLimit(Owner.TransactOpts(), FinneyToWei(5))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())

				tx, err = Wallet.SubmitLoadLimitUpdate(Controller.TransactOpts(ethertest.WithGasLimit(65000)), FinneyToWei(1))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isGasExhausted(tx, 65000)).To(BeFalse())
				Expect(isSuccessful(tx)).To(BeFalse())
			})

		})

		When("a random person submits daily load limit of 1 Finney after initialization", func() {
			It("should fail", func() {
				tx, err := Wallet.InitializeLoadLimit(Owner.TransactOpts(), FinneyToWei(5))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())

				tx, err = Wallet.SubmitLoadLimitUpdate(RandomAccount.TransactOpts(ethertest.WithGasLimit(65000)), FinneyToWei(1))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isGasExhausted(tx, 65000)).To(BeFalse())
				Expect(isSuccessful(tx)).To(BeFalse())
			})

		})

		When("I submit Load limit of 1 Finney after initialization", func() {
			BeforeEach(func() {
				tx, err := Wallet.InitializeLoadLimit(Owner.TransactOpts(), FinneyToWei(5))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())

				tx, err = Wallet.SubmitLoadLimitUpdate(Owner.TransactOpts(), FinneyToWei(1))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})

			It("should update the submission flag", func() {
				submitted, err := Wallet.SubmittedLoadLimit(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(submitted).To(BeTrue())
			})

			It("should emit a submission event", func() {
				it, err := Wallet.FilterSubmittedLoadLimitUpdate(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(it.Next()).To(BeTrue())
				evt := it.Event
				Expect(it.Next()).To(BeFalse())
				Expect(evt.Amount).To(Equal(FinneyToWei(1)))
			})

			It("should have pending load limit of 1 Finney", func() {
				ptl, err := Wallet.PendingLoadLimit(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(ptl.String()).To(Equal(FinneyToWei(1).String()))
			})

			When("the controller cancels the limit change", func() {
				BeforeEach(func() {
					tx, err := Wallet.CancelLoadLimitUpdate(Controller.TransactOpts(), FinneyToWei(1))
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeTrue())
				})

				It("should emit a cancellation event", func() {
					it, err := Wallet.FilterCancelledLoadLimitUpdate(nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(it.Next()).To(BeTrue())
					evt := it.Event
					Expect(it.Next()).To(BeFalse())
					Expect(evt.Sender).To(Equal(Controller.Address()))
				})

				It("should set pending load limit to 0", func() {
					psl, err := Wallet.PendingLoadLimit(nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(psl.String()).To(Equal("0"))
				})
			})

			When("the controller cancels the limit change with wrong amount", func() {
				It("should fail", func() {
					tx, err := Wallet.CancelLoadLimitUpdate(Controller.TransactOpts(ethertest.WithGasLimit(65000)), FinneyToWei(2))
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isGasExhausted(tx, 65000)).To(BeFalse())
					Expect(isSuccessful(tx)).To(BeFalse())
				})
			})

			When("the Owner tries to cancel the limit change", func() {
				It("should fail", func() {
					tx, err := Wallet.CancelLoadLimitUpdate(Owner.TransactOpts(ethertest.WithGasLimit(65000)), FinneyToWei(1))
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isGasExhausted(tx, 65000)).To(BeFalse())
					Expect(isSuccessful(tx)).To(BeFalse())
				})
			})

			When("a random person tries to cancel the limit change", func() {
				It("should fail", func() {
					tx, err := Wallet.CancelLoadLimitUpdate(RandomAccount.TransactOpts(ethertest.WithGasLimit(65000)), FinneyToWei(1))
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isGasExhausted(tx, 65000)).To(BeFalse())
					Expect(isSuccessful(tx)).To(BeFalse())
				})
			})

			When("the Owner tries to submit a second load limit of 500 Finney", func() {
				It("should fail", func() {
					tx, err := Wallet.SubmitLoadLimitUpdate(Owner.TransactOpts(ethertest.WithGasLimit(65000)), FinneyToWei(500))
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isGasExhausted(tx, 65000)).To(BeFalse())
					Expect(isSuccessful(tx)).To(BeFalse())
				})

			})

			When("I try to confirm the load limit", func() {
				It("should fail", func() {
					tx, err := Wallet.ConfirmLoadLimitUpdate(Owner.TransactOpts(ethertest.WithGasLimit(65000)), FinneyToWei(1))
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isGasExhausted(tx, 65000)).To(BeFalse())
					Expect(isSuccessful(tx)).To(BeFalse())
				})
			})

			When("a random person tries to confirm the load limit", func() {
				It("should fail", func() {
					tx, err := Wallet.ConfirmLoadLimitUpdate(RandomAccount.TransactOpts(ethertest.WithGasLimit(65000)), FinneyToWei(1))
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isGasExhausted(tx, 65000)).To(BeFalse())
					Expect(isSuccessful(tx)).To(BeFalse())
				})
			})

			When("the controller confirms the load limit using the wrong amount", func() {
				It("should fail", func() {
					tx, err := Wallet.ConfirmLoadLimitUpdate(Controller.TransactOpts(ethertest.WithGasLimit(100000)), FinneyToWei(2))
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isGasExhausted(tx, 100000)).To(BeFalse())
					Expect(isSuccessful(tx)).To(BeFalse())
				})
			})

			When("the controller confirms the load limit", func() {
				BeforeEach(func() {
					tx, err := Wallet.ConfirmLoadLimitUpdate(Controller.TransactOpts(), FinneyToWei(1))
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeTrue())
				})

				It("should have 1 Finney available for loading", func() {
					ll, err := Wallet.LoadAvailable(nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(ll.String()).To(Equal(FinneyToWei(1).String()))
				})

				When("I submit a second load limit to 500 Finney", func() {
					BeforeEach(func() {
						tx, err := Wallet.SubmitLoadLimitUpdate(Owner.TransactOpts(), FinneyToWei(500))
						Expect(err).ToNot(HaveOccurred())
						Backend.Commit()
						Expect(isSuccessful(tx)).To(BeTrue())
					})

					When("the controller confirms the load limit", func() {
						BeforeEach(func() {
							tx, err := Wallet.ConfirmLoadLimitUpdate(Controller.TransactOpts(), FinneyToWei(500))
							Expect(err).ToNot(HaveOccurred())
							Backend.Commit()
							Expect(isSuccessful(tx)).To(BeTrue())
						})

						It("should have 1 Finney available for loading", func() {
							tl, err := Wallet.LoadAvailable(nil)
							Expect(err).ToNot(HaveOccurred())
							Expect(tl.String()).To(Equal(FinneyToWei(1).String()))
						})
						When("I wait for longer than a day", func() {
							BeforeEach(func() {
								Backend.AdjustTime(time.Hour * 25)
								Backend.Commit()
							})

							It("should have 500 Finney available for further loading", func() {
								ll, err := Wallet.LoadAvailable(nil)
								Expect(err).ToNot(HaveOccurred())
								Expect(ll.String()).To(Equal(FinneyToWei(500).String()))
							})

						})

					})

				})

			})

		})

	})

})
