package wallet_test

import (
	"context"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/tokencard/ethertest"
	. "github.com/tokencard/ethertest"
)

var _ = Describe("topupGas", func() {

	Context("when the wallet has enough ETH", func() {
		BeforeEach(func() {
			bankAccount.MustTransfer(be, wa, ethToWei(10))
		})

		var tx *types.Transaction
		var err error
		var caller *ethertest.Account

		BeforeEach(func() {
			caller = controller
			bankAccount.MustTransfer(be, controller.Address(), ONE_ETH)
		})

		Context("When called by the wallet controller and is lower than topup limit", func() {

			BeforeEach(func() {
				tx, err = w.TopUpGas(caller.TransactOpts(WithGasLimit(81000)), ONE_FINNEY)
				be.Commit()
			})

			It("Should not return error", func() {
				Expect(err).ToNot(HaveOccurred())
			})

			It("should succeed", func() {
				Expect(isSuccessful(tx)).To(BeTrue())
			})
		})

		Context("When the value is above topup limit", func() {

			BeforeEach(func() {
				tx, err = w.TopUpGas(caller.TransactOpts(WithGasLimit(81000)), ONE_ETH)
				Expect(err).ToNot(HaveOccurred())
				be.Commit()
			})

			It("Should not return error", func() {
				Expect(err).ToNot(HaveOccurred())
			})

			It("should not fail", func() {
				Expect(isSuccessful(tx)).To(BeTrue())
			})

			It("Should topup only the topup limit of the gas", func() {
				b, e := be.BalanceAt(context.Background(), wa, nil)
				Expect(e).ToNot(HaveOccurred())
				Expect(b.String()).To(Equal("9500000000000000000"))
				Expect(owner.Balance(be).String()).To(Equal("1500000000000000000"))
			})

		})

		Context("When daily limit has been exausted", func() {
			BeforeEach(func() {
				caller = controller
				bankAccount.MustTransfer(be, controller.Address(), ONE_ETH)
			})

			BeforeEach(func() {
				tx, err = w.TopUpGas(caller.TransactOpts(WithGasLimit(81000)), ONE_ETH)
				Expect(err).ToNot(HaveOccurred())
				be.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})

			BeforeEach(func() {
				tx, err = w.TopUpGas(caller.TransactOpts(WithGasLimit(165000)), ONE_ETH)
				Expect(err).ToNot(HaveOccurred())
				be.Commit()
			})

			It("Should not return error", func() {
				Expect(err).ToNot(HaveOccurred())
			})

			It("should fail", func() {
				Expect(isSuccessful(tx)).To(BeFalse())
			})

			Context("When I wait for one day", func() {
				BeforeEach(func() {
					be.AdjustTime(time.Hour*24 + time.Second)
					be.Commit()
				})

				BeforeEach(func() {
					tx, err = w.TopUpGas(caller.TransactOpts(WithGasLimit(165000)), ONE_ETH)
					Expect(err).ToNot(HaveOccurred())
					be.Commit()
				})

				It("Should not return error", func() {
					Expect(err).ToNot(HaveOccurred())
				})

				It("should not fail", func() {
					Expect(isSuccessful(tx)).To(BeTrue())
				})

				It("Should topup only the topup limit of the gas", func() {

					b, e := be.BalanceAt(context.Background(), wa, nil)
					Expect(e).ToNot(HaveOccurred())
					Expect(b.String()).To(Equal("9000000000000000000"))
					Expect(owner.Balance(be).String()).To(Equal("2000000000000000000"))
				})

			})

		})

		Context("When called bu the wallet owner", func() {
			BeforeEach(func() {
				caller = owner
			})

			Context("when the value is lower than topup limit", func() {

				BeforeEach(func() {
					tx, err = w.TopUpGas(caller.TransactOpts(WithGasLimit(81000)), finneyToWei(500))
					Expect(err).ToNot(HaveOccurred())
					be.Commit()
				})

				It("Should not return error", func() {
					Expect(err).ToNot(HaveOccurred())
				})

				It("should succeed", func() {
					Expect(isSuccessful(tx)).To(BeTrue())
				})

			})

			Context("When the value is above topup limit", func() {

				BeforeEach(func() {
					tx, err = w.TopUpGas(caller.TransactOpts(WithGasLimit(81000)), finneyToWei(800))
					Expect(err).ToNot(HaveOccurred())
					be.Commit()
				})

				It("Should not return error", func() {
					Expect(err).ToNot(HaveOccurred())
				})

				It("should not fail", func() {
					Expect(isSuccessful(tx)).To(BeTrue())
				})

				It("Should topup only the topup limit of the gas", func() {
					b, e := be.BalanceAt(context.Background(), wa, nil)
					Expect(e).ToNot(HaveOccurred())
					Expect(b.String()).To(Equal("9500000000000000000"))
					Expect(owner.Balance(be).String()).To(AlmostEqual("1500000000000000000"))
				})

			})

			Context("When daily limit has been exausted", func() {

				BeforeEach(func() {
					tx, err = w.TopUpGas(caller.TransactOpts(WithGasLimit(81000)), finneyToWei(500))
					Expect(err).ToNot(HaveOccurred())
					be.Commit()
				})

				BeforeEach(func() {
					tx, err = w.TopUpGas(caller.TransactOpts(WithGasLimit(165000)), ONE_ETH)
					be.Commit()
				})

				It("Should not return error", func() {
					Expect(err).ToNot(HaveOccurred())
				})

				It("should fail", func() {
					Expect(isSuccessful(tx)).To(BeFalse())
				})

				Context("When I wait for one day", func() {
					BeforeEach(func() {
						be.AdjustTime(time.Hour*24 + time.Second)
						be.Commit()
					})

					BeforeEach(func() {
						tx, err = w.TopUpGas(caller.TransactOpts(WithGasLimit(165000)), ONE_ETH)
						be.Commit()
					})

					It("Should not return error", func() {
						Expect(err).ToNot(HaveOccurred())
					})

					It("should not fail", func() {

						Expect(isSuccessful(tx)).To(BeTrue())
					})

					It("Should topup only the topup limit of the gas", func() {
						b, e := be.BalanceAt(context.Background(), wa, nil)
						Expect(e).ToNot(HaveOccurred())
						Expect(b.String()).To(Equal("9000000000000000000"))
						Expect(owner.Balance(be).String()).To(AlmostEqual("2000000000000000000"))
					})

				})
			})
		})

		Context("When called by some random address and is lower than topup limit", func() {

			BeforeEach(func() {
				caller = randomAccount
			})

			BeforeEach(func() {
				tx, err = w.TopUpGas(caller.TransactOpts(WithGasLimit(165000)), ONE_ETH)
				Expect(err).ToNot(HaveOccurred())
				be.Commit()
			})

			It("Should not return error", func() {
				Expect(err).ToNot(HaveOccurred())
			})

			It("should fail", func() {
				Expect(isSuccessful(tx)).To(BeFalse())
			})
		})

	})
})
