package wallet_test

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("transfer", func() {

	Context("when the wallet has enough ETH", func() {
		BeforeEach(func() {
			bankWallet.MustTransfer(be, wa, ONE_ETH)
			bankWallet.MustTransfer(be, controller.Address(), ONE_ETH)
		})

		var tx *types.Transaction

		Context("When I transfer 1 Finney to a randon person", func() {
			BeforeEach(func() {
				var err error
				tx, err = w.Transfer(owner.TransactOptsWithGasLimit(81000), randomPerson.Address(), common.HexToAddress("0x"), ONE_FINNEY)
				Expect(err).ToNot(HaveOccurred())
				be.Commit()
			})
			It("Should succeed", func() {
				Expect(isSuccessful(tx)).To(BeTrue())
			})

			It("should reduce available transfer for today by 1 Finney", func() {
				av, err := w.SpendAvailable(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(av.String()).To(Equal("99999000000000000000"))
			})

		})

		Context("When controller tries to transfer 1 Finney to a random person", func() {

			BeforeEach(func() {
				var err error
				tx, err = w.Transfer(controller.TransactOptsWithGasLimit(81000), randomPerson.Address(), common.HexToAddress("0x"), ONE_FINNEY)
				Expect(err).ToNot(HaveOccurred())
				be.Commit()
			})

			It("Should fail", func() {
				Expect(isSuccessful(tx)).To(BeFalse())
			})

		})

		Context("When a random person tries to transfer 1 Finney to a random person", func() {

			BeforeEach(func() {
				var err error
				tx, err = w.Transfer(randomPerson.TransactOptsWithGasLimit(81000), randomPerson.Address(), common.HexToAddress("0x"), ONE_FINNEY)
				Expect(err).ToNot(HaveOccurred())
				be.Commit()
			})

			It("Should fail", func() {
				Expect(isSuccessful(tx)).To(BeFalse())
			})

		})

		Context("When I have one thousand tokens", func() {
			BeforeEach(func() {
				var err error
				tx, err = tkn.Credit(bankWallet.TransactOpts(), wa, big.NewInt(1000))
				Expect(err).ToNot(HaveOccurred())
				be.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())

			})

			Context("When I transfer 300 tokens to a random person", func() {
				BeforeEach(func() {
					var err error
					tx, err = w.SubmitWhitelistAddition(owner.TransactOpts(), []common.Address{randomPerson.Address()})
					Expect(err).ToNot(HaveOccurred())
					be.Commit()
					Expect(isSuccessful(tx)).To(BeTrue())

					tx, err = w.Transfer(owner.TransactOpts(), randomPerson.Address(), tkna, big.NewInt(300))
					Expect(err).ToNot(HaveOccurred())
					be.Commit()
				})

				It("should succeed", func() {
					Expect(isSuccessful(tx)).To(BeTrue())
				})

				It("should increase TKN balance of the random person", func() {
					b, err := tkn.BalanceOf(nil, randomPerson.Address())
					Expect(err).ToNot(HaveOccurred())
					Expect(b.String()).To(Equal("300"))
				})

				It("should decrease TKN balance of the wallet", func() {
					b, err := w.Balance(nil, tkna)
					Expect(err).ToNot(HaveOccurred())
					Expect(b.String()).To(Equal("700"))
				})

				It("should reduce the availabe balance for the day", func() {
					av, err := w.SpendAvailable(nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(av.String()).To(AlmostEqual("99999999999951010000"))
				})
			})

			Context("When controller tries to transfer one token to a random person", func() {
				BeforeEach(func() {
					var err error
					tx, err = w.Transfer(controller.TransactOptsWithGasLimit(80000), randomPerson.Address(), tkna, big.NewInt(1))
					Expect(err).ToNot(HaveOccurred())
					be.Commit()
				})

				It("should fail", func() {
					Expect(isSuccessful(tx)).To(BeFalse())
				})
			})

			Context("When random person tries to transfer one token to a random person", func() {

				BeforeEach(func() {
					var err error
					tx, err = w.Transfer(randomPerson.TransactOptsWithGasLimit(80000), randomPerson.Address(), tkna, big.NewInt(1))
					Expect(err).ToNot(HaveOccurred())
					be.Commit()
				})

				It("should fail", func() {
					Expect(isSuccessful(tx)).To(BeFalse())
				})
			})

		})

	})
})
