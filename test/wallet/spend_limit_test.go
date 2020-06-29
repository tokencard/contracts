package wallet_test

import (
	"time"

	"github.com/ethereum/go-ethereum/common"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tokencard/contracts/v3/test/shared"
	"github.com/tokencard/ethertest"
)

var _ = Describe("spendAvailable", func() {
	BeforeEach(func() {
		BankAccount.MustTransfer(Backend, Controller.Address(), EthToWei(1))
	})

	When("wallet has been deployed", func() {
		It("should have 100 ETH spend available", func() {
			av, err := WalletProxy.SpendLimitAvailable(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(av.String()).To(Equal(EthToWei(100).String()))
		})

		When("a random person tries to set the spend limit", func() {
			It("should fail", func() {
				tx, err := WalletProxy.SetSpendLimit(RandomAccount.TransactOpts(ethertest.WithGasLimit(100000)), EthToWei(1))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isGasExhausted(tx, 100000)).To(BeFalse())
				Expect(isSuccessful(tx)).To(BeFalse())
			})
		})

		When("controller tries to set the spend limit", func() {
			It("should fail", func() {
				tx, err := WalletProxy.SetSpendLimit(Controller.TransactOpts(ethertest.WithGasLimit(100000)), EthToWei(1))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isGasExhausted(tx, 100000)).To(BeFalse())
				Expect(isSuccessful(tx)).To(BeFalse())
			})
		})

		When("Owner sets the spendLimit to 1 ETH", func() {
			BeforeEach(func() {
				tx, err := WalletProxy.SetSpendLimit(Owner.TransactOpts(), EthToWei(1))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})

			It("should emit a spend limit set event", func() {
				it, err := WalletProxy.FilterSetSpendLimit(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(it.Next()).To(BeTrue())
				evt := it.Event
				Expect(it.Next()).To(BeFalse())
				Expect(evt.Sender).To(Equal(Owner.Address()))
				Expect(evt.Amount).To(Equal(EthToWei(1)))
			})

			It("should lower the spend available to 1 ETH", func() {
				av, err := WalletProxy.SpendLimitAvailable(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(av.String()).To(Equal(EthToWei(1).String()))
			})

			It("should have spend limit of 1 ETH", func() {
				sl, err := WalletProxy.SpendLimitValue(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(sl.String()).To(Equal(EthToWei(1).String()))
			})

			When("the owner tries to initialize the spending limit again", func() {
				It("should fail", func() {
					tx, err := WalletProxy.SetSpendLimit(Owner.TransactOpts(ethertest.WithGasLimit(100000)), EthToWei(1))
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isGasExhausted(tx, 100000)).To(BeFalse())
					Expect(isSuccessful(tx)).To(BeFalse())
				})
			})
		})

		When("I set spend limit to 1000 ETH", func() {
			BeforeEach(func() {
				tx, err := WalletProxy.SetSpendLimit(Owner.TransactOpts(), EthToWei(1000))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})

			It("should update the set flag", func() {
				initialized, err := WalletProxy.SpendLimitControllerConfirmationRequired(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(initialized).To(BeTrue())
			})

			It("should keep the spend available at 100 ETH", func() {
				av, err := WalletProxy.SpendLimitAvailable(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(av.String()).To(Equal(EthToWei(100).String()))
			})

			When("I wait for 24 hours", func() {
				BeforeEach(func() {
					Backend.AdjustTime(time.Hour*24 + time.Second)
					Backend.Commit()
				})

				It("should increase the spend available to 1000 ETH", func() {
					av, err := WalletProxy.SpendLimitAvailable(nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(av.String()).To(Equal(EthToWei(1000).String()))
				})

				When("the contract has one eth", func() {
					BeforeEach(func() {
						BankAccount.Transfer(Backend, WalletProxyAddress, EthToWei(1))
					})

					When("I transfer one ETH", func() {
						BeforeEach(func() {
							tx, err := WalletProxy.Transfer(Owner.TransactOpts(), RandomAccount.Address(), common.HexToAddress("0x"), EthToWei(1))
							Expect(err).ToNot(HaveOccurred())
							Backend.Commit()
							Expect(isSuccessful(tx)).To(BeTrue())
						})
						It("should have 999 ETH spend available", func() {
							sa, err := WalletProxy.SpendLimitAvailable(nil)
							Expect(err).ToNot(HaveOccurred())
							Expect(sa.String()).To(Equal(EthToWei(999).String()))
						})
					})
				})
			})
		})

		When("I submit the spend limit of 1 ETH before initialization", func() {
			It("should fail", func() {
				tx, err := WalletProxy.SubmitSpendLimitUpdate(Owner.TransactOpts(ethertest.WithGasLimit(100000)), EthToWei(1))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isGasExhausted(tx, 100000)).To(BeFalse())
				Expect(isSuccessful(tx)).To(BeFalse())
			})
		})

		When("I submit 3  spend limit updates of 1 ETH after having set it", func() {
			BeforeEach(func() {
				tx, err := WalletProxy.SetSpendLimit(Owner.TransactOpts(ethertest.WithGasLimit(100000)), EthToWei(2))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isGasExhausted(tx, 100000)).To(BeFalse())
				Expect(isSuccessful(tx)).To(BeTrue())

				tx, err = WalletProxy.SubmitSpendLimitUpdate(Owner.TransactOpts(), EthToWei(3))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())

				tx, err = WalletProxy.SubmitSpendLimitUpdate(Owner.TransactOpts(), EthToWei(4))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())

				tx, err = WalletProxy.SubmitSpendLimitUpdate(Owner.TransactOpts(), EthToWei(1))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})

			It("should emit 3 submission events", func() {
				it, err := WalletProxy.FilterSubmittedSpendLimitUpdate(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(it.Next()).To(BeTrue())
				evt := it.Event
				Expect(it.Next()).To(BeTrue())
				Expect(evt.Amount).To(Equal(EthToWei(3)))
				evt = it.Event
				Expect(it.Next()).To(BeTrue())
				Expect(evt.Amount).To(Equal(EthToWei(4)))
				evt = it.Event
				Expect(it.Next()).To(BeFalse())
				Expect(evt.Amount).To(Equal(EthToWei(1)))
			})

			It("should not reduce the spend available", func() {
				av, err := WalletProxy.SpendLimitAvailable(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(av.String()).To(Equal(EthToWei(2).String()))
			})

			It("should set the pending spend limit to 1 ETH", func() {
				psl, err := WalletProxy.SpendLimitPending(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(psl.String()).To(Equal(EthToWei(1).String()))
			})

			When("the controller confirms the spend limit providing the wrong amount", func() {
				It("should fail", func() {
					tx, err := WalletProxy.ConfirmSpendLimitUpdate(Controller.TransactOpts(ethertest.WithGasLimit(100000)), EthToWei(2))
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isGasExhausted(tx, 100000)).To(BeFalse())
					Expect(isSuccessful(tx)).To(BeFalse())
				})
			})

			When("the controller confirms the spend limit", func() {
				BeforeEach(func() {
					tx, err := WalletProxy.ConfirmSpendLimitUpdate(Controller.TransactOpts(), EthToWei(1))
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeTrue())
				})

				It("should emit a spend limit set event", func() {
					it, err := WalletProxy.FilterSetSpendLimit(nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(it.Next()).To(BeTrue())
					Expect(it.Next()).To(BeTrue())
					evt := it.Event
					Expect(it.Next()).To(BeFalse())
					Expect(evt.Amount).To(Equal(EthToWei(1)))
				})

				It("should reduce the spend available to 1 ETH", func() {
					av, err := WalletProxy.SpendLimitAvailable(nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(av.String()).To(Equal(EthToWei(1).String()))
				})

				When("I try to set the spend limit", func() {
					It("should fail", func() {
						tx, err := WalletProxy.SetSpendLimit(Owner.TransactOpts(ethertest.WithGasLimit(100000)), EthToWei(1000))
						Expect(err).ToNot(HaveOccurred())
						Backend.Commit()
						Expect(isGasExhausted(tx, 100000)).To(BeFalse())
						Expect(isSuccessful(tx)).To(BeFalse())
					})
				})

			})

			When("I try to update the spendLimit again", func() {
				It("should succeed", func() {
					tx, err := WalletProxy.SubmitSpendLimitUpdate(Owner.TransactOpts(), EthToWei(1))
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeTrue())
				})
			})
		})
	})

})
