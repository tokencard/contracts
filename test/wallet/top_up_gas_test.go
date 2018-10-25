package wallet_test

import (
	"context"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tokencard/contracts/test/shared"
	"github.com/tokencard/ethertest"
)

var _ = Describe("topUpGas", func() {

	Context("when the wallet has enough ETH", func() {
		BeforeEach(func() {
			BankAccount.MustTransfer(Backend, WalletAddress, EthToWei(10))
		})

		var tx *types.Transaction
		var err error
		var caller *ethertest.Account

		BeforeEach(func() {
			caller = Controller
			BankAccount.MustTransfer(Backend, Controller.Address(), EthToWei(1))
		})

		Context("When called by the wallet controller and is lower than top up limit", func() {

			BeforeEach(func() {
				tx, err = Wallet.TopUpGas(caller.TransactOpts(ethertest.WithGasLimit(81000)), FinneyToWei(1))
				Backend.Commit()
			})

			It("Should not return error", func() {
				Expect(err).ToNot(HaveOccurred())
			})

			It("should succeed", func() {
				Expect(isSuccessful(tx)).To(BeTrue())
			})
		})

		Context("When the value is above top up limit", func() {

			BeforeEach(func() {
				tx, err = Wallet.TopUpGas(caller.TransactOpts(ethertest.WithGasLimit(81000)), EthToWei(1))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
			})

			It("should not return error", func() {
				Expect(err).ToNot(HaveOccurred())
			})

			It("should not fail", func() {
				Expect(isSuccessful(tx)).To(BeTrue())
			})

			It("should top up only the top up limit of the gas", func() {
				b, e := Backend.BalanceAt(context.Background(), WalletAddress, nil)
				Expect(e).ToNot(HaveOccurred())
				Expect(b.String()).To(Equal("9500000000000000000"))
				Expect(Owner.Balance(Backend).String()).To(Equal("1500000000000000000"))
			})

		})

		Context("When daily limit has been exausted", func() {
			BeforeEach(func() {
				caller = Controller
				BankAccount.MustTransfer(Backend, Controller.Address(), EthToWei(1))
			})

			BeforeEach(func() {
				tx, err = Wallet.TopUpGas(caller.TransactOpts(ethertest.WithGasLimit(81000)), EthToWei(1))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})

			BeforeEach(func() {
				tx, err = Wallet.TopUpGas(caller.TransactOpts(ethertest.WithGasLimit(165000)), EthToWei(1))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
			})

			It("should not return error", func() {
				Expect(err).ToNot(HaveOccurred())
			})

			It("should fail", func() {
				Expect(isSuccessful(tx)).To(BeFalse())
			})

			Context("When I wait for one day", func() {
				BeforeEach(func() {
					Backend.AdjustTime(time.Hour*24 + time.Second)
					Backend.Commit()
				})

				BeforeEach(func() {
					tx, err = Wallet.TopUpGas(caller.TransactOpts(ethertest.WithGasLimit(165000)), EthToWei(1))
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
				})

				It("should not return error", func() {
					Expect(err).ToNot(HaveOccurred())
				})

				It("should not fail", func() {
					Expect(isSuccessful(tx)).To(BeTrue())
				})

				It("should top up only the top up limit of the gas", func() {

					b, e := Backend.BalanceAt(context.Background(), WalletAddress, nil)
					Expect(e).ToNot(HaveOccurred())
					Expect(b.String()).To(Equal("9000000000000000000"))
					Expect(Owner.Balance(Backend).String()).To(Equal("2000000000000000000"))
				})

			})

		})

		Context("When called bu the wallet owner", func() {
			BeforeEach(func() {
				caller = Owner
			})

			Context("When the value is lower than top up limit", func() {

				BeforeEach(func() {
					tx, err = Wallet.TopUpGas(caller.TransactOpts(ethertest.WithGasLimit(81000)), FinneyToWei(500))
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
				})

				It("should not return error", func() {
					Expect(err).ToNot(HaveOccurred())
				})

				It("should succeed", func() {
					Expect(isSuccessful(tx)).To(BeTrue())
				})

			})

			Context("When the value is above top up limit", func() {

				BeforeEach(func() {
					tx, err = Wallet.TopUpGas(caller.TransactOpts(ethertest.WithGasLimit(81000)), FinneyToWei(800))
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
				})

				It("should not return error", func() {
					Expect(err).ToNot(HaveOccurred())
				})

				It("should not fail", func() {
					Expect(isSuccessful(tx)).To(BeTrue())
				})

				It("should top up only the top up limit of the gas", func() {
					b, e := Backend.BalanceAt(context.Background(), WalletAddress, nil)
					Expect(e).ToNot(HaveOccurred())
					Expect(b.String()).To(Equal("9500000000000000000"))
					Expect(Owner.Balance(Backend).String()).To(AlmostEqual("1500000000000000000"))
				})

			})

			Context("When daily limit has been exausted", func() {

				BeforeEach(func() {
					tx, err = Wallet.TopUpGas(caller.TransactOpts(ethertest.WithGasLimit(81000)), FinneyToWei(500))
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
				})

				BeforeEach(func() {
					tx, err = Wallet.TopUpGas(caller.TransactOpts(ethertest.WithGasLimit(165000)), EthToWei(1))
					Backend.Commit()
				})

				It("should not return error", func() {
					Expect(err).ToNot(HaveOccurred())
				})

				It("should fail", func() {
					Expect(isSuccessful(tx)).To(BeFalse())
				})

				Context("When I wait for one day", func() {
					BeforeEach(func() {
						Backend.AdjustTime(time.Hour*24 + time.Second)
						Backend.Commit()
					})

					BeforeEach(func() {
						tx, err = Wallet.TopUpGas(caller.TransactOpts(ethertest.WithGasLimit(165000)), EthToWei(1))
						Backend.Commit()
					})

					It("should not return error", func() {
						Expect(err).ToNot(HaveOccurred())
					})

					It("should not fail", func() {

						Expect(isSuccessful(tx)).To(BeTrue())
					})

					It("should top up only the top up limit of the gas", func() {
						b, e := Backend.BalanceAt(context.Background(), WalletAddress, nil)
						Expect(e).ToNot(HaveOccurred())
						Expect(b.String()).To(Equal("9000000000000000000"))
						Expect(Owner.Balance(Backend).String()).To(AlmostEqual("2000000000000000000"))
					})

				})
			})
		})

		Context("When called by some random address and is lower than top up limit", func() {

			BeforeEach(func() {
				caller = RandomAccount
			})

			BeforeEach(func() {
				tx, err = Wallet.TopUpGas(caller.TransactOpts(ethertest.WithGasLimit(165000)), EthToWei(1))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
			})

			It("should not return error", func() {
				Expect(err).ToNot(HaveOccurred())
			})

			It("should fail", func() {
				Expect(isSuccessful(tx)).To(BeFalse())
			})
		})

	})
})
