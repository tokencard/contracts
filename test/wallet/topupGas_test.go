package wallet_test

import (
	"context"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/tokencard/ethertest"
)

var _ = Describe("topupGas", func() {

	Context("when the wallet has enough ETH", func() {
		BeforeEach(func() {
			bankWallet.MustTransfer(be, wa, TEN_ETH)
		})

		var tx *types.Transaction
		var err error
		var caller *ethertest.Wallet
		var value *big.Int

		BeforeEach(func() {
			value = ONE_FINNEY
		})

		JustBeforeEach(func() {
			to := caller.TransactOpts()
			to.GasLimit = 81000
			tx, err = w.TopupGas(to, value)
			be.Commit()
		})

		Context("When called by the wallet controller and is lower than topup limit", func() {
			BeforeEach(func() {
				caller = controller
				bankWallet.MustTransfer(be, controller.Address(), ONE_ETH)
			})

			It("Should not return error", func() {
				Expect(err).ToNot(HaveOccurred())
			})

			It("should succeed", func() {
				Expect(isSuccessful(tx)).To(BeTrue())
			})

			Context("When the value is above topup limit", func() {
				BeforeEach(func() {
					value = ONE_ETH
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

				Context("When daily limit has been exausted", func() {
					BeforeEach(func() {
						to := caller.TransactOpts()
						to.GasLimit = 165000
						tx, err = w.TopupGas(to, ONE_ETH)
						be.Commit()
						Expect(err).ToNot(HaveOccurred())
						Expect(isSuccessful(tx)).To(BeTrue())
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
			})

		})

		Context("When called by the wallet owner and is lower than topup limit", func() {
			BeforeEach(func() {
				caller = owner
			})

			It("Should not return error", func() {
				Expect(err).ToNot(HaveOccurred())
			})

			It("should succeed", func() {
				Expect(isSuccessful(tx)).To(BeTrue())
			})

			Context("When the value is above topup limit", func() {
				BeforeEach(func() {
					value = ONE_ETH
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

				Context("When daily limit has been exausted", func() {
					BeforeEach(func() {
						to := caller.TransactOpts()
						to.GasLimit = 165000
						tx, err = w.TopupGas(to, ONE_ETH)
						be.Commit()
						Expect(err).ToNot(HaveOccurred())
						Expect(isSuccessful(tx)).To(BeTrue())
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
		})

		Context("When called by some random address and is lower than topup limit", func() {
			BeforeEach(func() {
				caller = randomPerson
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
