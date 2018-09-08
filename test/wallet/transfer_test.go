package wallet_test

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/tokencard/ethertest"
)

var _ = Describe("transfer", func() {

	Context("when the wallet has enough ETH", func() {
		BeforeEach(func() {
			bankWallet.MustTransfer(be, wa, ONE_ETH)
			bankWallet.MustTransfer(be, controller.Address(), ONE_ETH)
		})

		var tx *types.Transaction
		var caller *ethertest.Wallet
		var amount *big.Int
		var asset common.Address
		var to common.Address

		BeforeEach(func() {
			caller = owner
			amount = ONE_FINNEY
			asset = common.HexToAddress("0x")
		})

		JustBeforeEach(func() {
			top := caller.TransactOpts()
			top.GasLimit = 81000
			var err error
			tx, err = w.Transfer(top, to, asset, amount)
			Expect(err).ToNot(HaveOccurred())
			be.Commit()
		})

		Context("When I transfer 1 Finney to a randon person", func() {
			BeforeEach(func() {
				to = randomPerson.Address()
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
				caller = controller
				to = randomPerson.Address()
			})
			It("Should fail", func() {
				Expect(isSuccessful(tx)).To(BeFalse())
			})

		})

		Context("When a random person tries to transfer 1 Finney to a random person", func() {

			BeforeEach(func() {
				caller = randomPerson
				bankWallet.MustTransfer(be, randomPerson.Address(), ONE_ETH)
				to = randomPerson.Address()
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
			Context("When I transfer one thousand tokens to a random person", func() {
				BeforeEach(func() {
					to = randomPerson.Address()
					asset = tkna
					amount = big.NewInt(1000)

					b, err := w.Balance(nil, asset)
					Expect(err).ToNot(HaveOccurred())
					Expect(b.String()).To(Equal("1000"))

				})

				It("should succeed", func() {
					Expect(isSuccessful(tx)).To(BeTrue())
				})

				It("should increase TKN balance of the random person", func() {
					b, err := tkn.BalanceOf(nil, randomPerson.Address())
					Expect(err).ToNot(HaveOccurred())
					Expect(b.String()).To(Equal("1000"))
				})

				It("should decrease TKN balance of the wallet", func() {
					b, err := w.Balance(nil, asset)
					Expect(err).ToNot(HaveOccurred())
					Expect(b.String()).To(Equal("0"))
				})

				It("should reduce the availabe balance for the day", func() {
					av, err := w.SpendAvailable(nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(av.String()).To(Equal("99999999983670000000"))
				})
			})

			Context("When controller tries to transfer one token to a random person", func() {
				BeforeEach(func() {
					to = randomPerson.Address()
					asset = tkna
					amount = big.NewInt(1)
					caller = controller
				})

				It("should fail", func() {
					Expect(isSuccessful(tx)).To(BeFalse())
				})
			})

			Context("When random person tries to transfer one token to a random person", func() {
				BeforeEach(func() {
					to = randomPerson.Address()
					asset = tkna
					amount = big.NewInt(1)
					caller = randomPerson
				})

				It("should fail", func() {
					Expect(isSuccessful(tx)).To(BeFalse())
				})
			})

		})

	})
})
