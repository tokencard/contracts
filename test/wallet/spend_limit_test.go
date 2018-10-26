package wallet_test

import (
	"time"

	"github.com/ethereum/go-ethereum/common"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tokencard/contracts/test/shared"
	"github.com/tokencard/ethertest"
)

var _ = Describe("spendAvailable", func() {
	BeforeEach(func() {
		BankAccount.MustTransfer(Backend, Controller.Address(), EthToWei(1))
	})

	Context("When wallet has been deployed", func() {
		It("should have 100 ETH spend available", func() {
			av, err := Wallet.SpendAvailable(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(av.String()).To(Equal(EthToWei(100).String()))
		})

		Context("When a random person tries to initialize the spend limit", func() {
			It("should fail", func() {
				tx, err := Wallet.InitializeSpendLimit(RandomAccount.TransactOpts(ethertest.WithGasLimit(100000)), EthToWei(1))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeFalse())
			})
		})

		Context("When controller tries to initialize the spend limit", func() {
			It("should fail", func() {
				tx, err := Wallet.InitializeSpendLimit(Controller.TransactOpts(ethertest.WithGasLimit(100000)), EthToWei(1))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeFalse())
			})
		})

		Context("When I initialize spend limit to 1 ETH", func() {
			BeforeEach(func() {
				tx, err := Wallet.InitializeSpendLimit(Owner.TransactOpts(), EthToWei(1))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})

			It("should emit a spend limit set event", func() {
				it, err := Wallet.FilterSetSpendLimit(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(it.Next()).To(BeTrue())
				evt := it.Event
				Expect(it.Next()).To(BeFalse())
				Expect(evt.Sender).To(Equal(Owner.Address()))
				Expect(evt.Amount).To(Equal(EthToWei(1)))
			})

			It("should lower the spend available to 1 ETH", func() {
				av, err := Wallet.SpendAvailable(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(av.String()).To(Equal(EthToWei(1).String()))
			})

			It("should have spend limit of 1 ETH", func() {
				sl, err := Wallet.SpendLimit(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(sl.String()).To(Equal(EthToWei(1).String()))
			})

			Context("When I try to initialize the spending limit again", func() {
				It("should fail", func() {
					tx, err := Wallet.InitializeSpendLimit(Owner.TransactOpts(ethertest.WithGasLimit(100000)), EthToWei(1))
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeFalse())
				})
			})
		})

		Context("When I initialize spend limit to 1000 ETH", func() {
			BeforeEach(func() {
				tx, err := Wallet.InitializeSpendLimit(Owner.TransactOpts(), EthToWei(1000))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})

			It("should update the initialization flag", func() {
				initialized, err := Wallet.InitializedSpendLimit(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(initialized).To(BeTrue())
			})

			It("should keep the spend available at 100 ETH", func() {
				av, err := Wallet.SpendAvailable(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(av.String()).To(Equal(EthToWei(100).String()))
			})

			Context("After I wait for 24 hours", func() {
				BeforeEach(func() {
					Backend.AdjustTime(time.Hour*24 + time.Second)
					Backend.Commit()
				})

				It("should increase the spend available to 1000 ETH", func() {
					av, err := Wallet.SpendAvailable(nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(av.String()).To(Equal(EthToWei(1000).String()))
				})

				Context("when the contract has one eth", func() {
					BeforeEach(func() {
						BankAccount.Transfer(Backend, WalletAddress, EthToWei(1))
					})

					Context("When I transfer one ETH", func() {
						BeforeEach(func() {
							tx, err := Wallet.Transfer(Owner.TransactOpts(), RandomAccount.Address(), common.HexToAddress("0x"), EthToWei(1))
							Expect(err).ToNot(HaveOccurred())
							Backend.Commit()
							Expect(isSuccessful(tx)).To(BeTrue())
						})
						It("should have 999 ETH spend available", func() {
							sa, err := Wallet.SpendAvailable(nil)
							Expect(err).ToNot(HaveOccurred())
							Expect(sa.String()).To(Equal(EthToWei(999).String()))
						})
					})
				})
			})
		})

		Context("When I submit the spend limit of 1 ETH before initialization", func() {
			It("should fail", func() {
				tx, err := Wallet.SubmitSpendLimit(Owner.TransactOpts(ethertest.WithGasLimit(100000)), EthToWei(1))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeFalse())

				submitted, err := Wallet.SubmittedSpendLimit(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(submitted).To(BeFalse())
			})
		})

		Context("When I submit the spend limit of 1 ETH after initialization", func() {
			BeforeEach(func() {
				tx, err := Wallet.InitializeSpendLimit(Owner.TransactOpts(ethertest.WithGasLimit(100000)), EthToWei(2))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())

				tx, err = Wallet.SubmitSpendLimit(Owner.TransactOpts(ethertest.WithGasLimit(100000)), EthToWei(1))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})

			It("should update the submission flag", func() {
				submitted, err := Wallet.SubmittedSpendLimit(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(submitted).To(BeTrue())
			})

			It("should emit the submission event", func() {
				it, err := Wallet.FilterSubmittedSpendLimitChange(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(it.Next()).To(BeTrue())
				evt := it.Event
				Expect(it.Next()).To(BeFalse())
				Expect(evt.Amount).To(Equal(EthToWei(1)))
			})

			It("should not reduce the spend available", func() {
				av, err := Wallet.SpendAvailable(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(av.String()).To(Equal(EthToWei(2).String()))
			})

			It("should set the pending spend limit to 1 ETH", func() {
				psl, err := Wallet.PendingSpendLimit(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(psl.String()).To(Equal(EthToWei(1).String()))
			})

			Context("When the controller confirms the spend limit", func() {
				BeforeEach(func() {
					tx, err := Wallet.ConfirmSpendLimit(Controller.TransactOpts())
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeTrue())
				})

				It("should emit a spend limit set event", func() {
					it, err := Wallet.FilterSetSpendLimit(nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(it.Next()).To(BeTrue())
					Expect(it.Next()).To(BeTrue())
					evt := it.Event
					Expect(it.Next()).To(BeFalse())
					Expect(evt.Amount).To(Equal(EthToWei(1)))
				})

				It("should reduce the spend available to 1 ETH", func() {
					av, err := Wallet.SpendAvailable(nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(av.String()).To(Equal(EthToWei(1).String()))
				})

				Context("When I try to initialize the spend limit", func() {
					It("should fail", func() {
						tx, err := Wallet.InitializeSpendLimit(Owner.TransactOpts(ethertest.WithGasLimit(100000)), EthToWei(1000))
						Expect(err).ToNot(HaveOccurred())
						Backend.Commit()
						Expect(isSuccessful(tx)).To(BeFalse())
					})
				})

			})
			Context("When the controller cancels the spend limit change", func() {
				BeforeEach(func() {
					tx, err := Wallet.CancelSpendLimit(Controller.TransactOpts())
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeTrue())
				})

				It("should emit a cancellation event", func() {
					it, err := Wallet.FilterCancelledSpendLimitChange(nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(it.Next()).To(BeTrue())
					evt := it.Event
					Expect(it.Next()).To(BeFalse())
					Expect(evt.Sender).To(Equal(Controller.Address()))
				})

				It("should set pending spend limit to 0", func() {
					psl, err := Wallet.PendingSpendLimit(nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(psl.String()).To(Equal("0"))
				})

				It("should make it possible to set the spending limit again", func() {
					tx, err := Wallet.SubmitSpendLimit(Owner.TransactOpts(), EthToWei(2))
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeTrue())
				})

			})

			Context("When I try to set spending limit again", func() {
				It("should fail", func() {
					tx, err := Wallet.SubmitSpendLimit(Owner.TransactOpts(ethertest.WithGasLimit(100000)), EthToWei(1))
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeFalse())
				})
			})
		})
	})

})
