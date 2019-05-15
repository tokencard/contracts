package wallet_test

import (
	"math/big"
	"strings"
    "math"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tokencard/contracts/test/shared"
)

var _ = Describe("executeTransaction", func() {

	Context("when the wallet has enough ETH", func() {
		BeforeEach(func() {
			BankAccount.MustTransfer(Backend, WalletAddress, EthToWei(1))
			BankAccount.MustTransfer(Backend, Controller.Address(), EthToWei(1))
		})

		var tx *types.Transaction

    When("the token exists and the rates have been updated", func() {
        BeforeEach(func() {
				tx, err := Oracle.AddTokens(Controller.TransactOpts(), []common.Address{NonCompliantERC20Address}, StringsToByte32("NON"), []*big.Int{ExponentiateDecimals(8)}, big.NewInt(20180913153211))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})

            BeforeEach(func() {
                tx, err := Oracle.UpdateTokenRate(Controller.TransactOpts(), NonCompliantERC20Address, big.NewInt(int64(0.00001633*math.Pow10(18))), big.NewInt(20180913153211))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})

		When("I have one thousand tokens", func() {
			BeforeEach(func() {
				var err error
				tx, err = NonCompliantERC20.Credit(BankAccount.TransactOpts(), WalletAddress, big.NewInt(1000))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})

			When("I transfer 300 tokens to a random person using 'executeTransaction'", func() {
				BeforeEach(func() {
					a, err := abi.JSON(strings.NewReader(ERC20ABI))
					Expect(err).ToNot(HaveOccurred())
					data, err := a.Pack("transfer", RandomAccount.Address(), big.NewInt(300))
					Expect(err).ToNot(HaveOccurred())

					tx, err = Wallet.ExecuteTransaction(Owner.TransactOpts(), NonCompliantERC20Address, big.NewInt(0), data)
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeTrue())
				})

                It("Should update the tokens map", func() {
    				token, err := Oracle.Tokens(nil, NonCompliantERC20Address)
    				Expect(err).ToNot(HaveOccurred())
    				Expect(token.Exists).To(BeTrue())
    				Expect(token.Rate).To(Equal(big.NewInt(int64(0.00001633*math.Pow10(18)))))
    			})

				It("should increase NonCompliantERC20 balance of the random person", func() {
					b, err := NonCompliantERC20.BalanceOf(nil, RandomAccount.Address())
					Expect(err).ToNot(HaveOccurred())
					Expect(b.String()).To(Equal("300"))
				})

				It("should decrease NonCompliantERC20 balance of the wallet", func() {
					b, err := Wallet.Balance(nil, NonCompliantERC20Address)
					Expect(err).ToNot(HaveOccurred())
					Expect(b.String()).To(Equal("700"))
				})

				It("should reduce the available daily spend balance", func() {
					av, err := Wallet.SpendAvailable(nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(av.String()).To(AlmostEqual("99999999999951010000"))
				})
			})

			When("random person is whitelisted", func() {
				BeforeEach(func() {
					tx, err := Wallet.InitializeWhitelist(Owner.TransactOpts(), []common.Address{RandomAccount.Address()})
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeTrue())
				})

				When("I transfer 300 tokens to a random person using 'executeTransaction'", func() {
					BeforeEach(func() {

						a, err := abi.JSON(strings.NewReader(ERC20ABI))
						Expect(err).ToNot(HaveOccurred())
						data, err := a.Pack("transfer", RandomAccount.Address(), big.NewInt(300))
						Expect(err).ToNot(HaveOccurred())

						tx, err = Wallet.ExecuteTransaction(Owner.TransactOpts(), NonCompliantERC20Address, big.NewInt(0), data)
						Expect(err).ToNot(HaveOccurred())
						Backend.Commit()
						Expect(isSuccessful(tx)).To(BeTrue())
					})

					It("should increase NonCompliantERC20 balance of the random person", func() {
						b, err := NonCompliantERC20.BalanceOf(nil, RandomAccount.Address())
						Expect(err).ToNot(HaveOccurred())
						Expect(b.String()).To(Equal("300"))
					})

					It("should decrease NonCompliantERC20 balance of the wallet", func() {
						b, err := Wallet.Balance(nil, NonCompliantERC20Address)
						Expect(err).ToNot(HaveOccurred())
						Expect(b.String()).To(Equal("700"))
					})

					It("should not reduce the available daily spend balance", func() {
						av, err := Wallet.SpendAvailable(nil)
						Expect(err).ToNot(HaveOccurred())
						Expect(av.String()).To(Equal("100000000000000000000"))
					})
				})
			})

			When("I approve 300 tokens to a random person using 'executeTransaction'", func() {
				BeforeEach(func() {
					a, err := abi.JSON(strings.NewReader(ERC20ABI))
					Expect(err).ToNot(HaveOccurred())
					data, err := a.Pack("approve", RandomAccount.Address(), big.NewInt(300))
					Expect(err).ToNot(HaveOccurred())

					tx, err = Wallet.ExecuteTransaction(Owner.TransactOpts(), NonCompliantERC20Address, big.NewInt(0), data)
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeTrue())
				})

				It("should not increase NonCompliantERC20 balance of the random person", func() {
					b, err := NonCompliantERC20.BalanceOf(nil, RandomAccount.Address())
					Expect(err).ToNot(HaveOccurred())
					Expect(b.String()).To(Equal("0"))
				})

				It("should not decrease NonCompliantERC20 balance of the wallet", func() {
					b, err := Wallet.Balance(nil, NonCompliantERC20Address)
					Expect(err).ToNot(HaveOccurred())
					Expect(b.String()).To(Equal("1000"))
				})

				It("should reduce the available daily spend balance", func() {
					av, err := Wallet.SpendAvailable(nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(av.String()).To(AlmostEqual("99999999999951010000"))
				})
			})

			When("random person is whitelisted", func() {
				BeforeEach(func() {
					tx, err := Wallet.InitializeWhitelist(Owner.TransactOpts(), []common.Address{RandomAccount.Address()})
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeTrue())
				})

				When("I approve 300 tokens to a random person using 'executeTransaction'", func() {
					BeforeEach(func() {
						a, err := abi.JSON(strings.NewReader(ERC20ABI))
						Expect(err).ToNot(HaveOccurred())
						data, err := a.Pack("approve", RandomAccount.Address(), big.NewInt(300))
						Expect(err).ToNot(HaveOccurred())

						tx, err = Wallet.ExecuteTransaction(Owner.TransactOpts(), NonCompliantERC20Address, big.NewInt(0), data)
						Expect(err).ToNot(HaveOccurred())
						Backend.Commit()
						Expect(isSuccessful(tx)).To(BeTrue())
					})

					It("should not increase NonCompliantERC20 balance of the random person", func() {
						b, err := NonCompliantERC20.BalanceOf(nil, RandomAccount.Address())
						Expect(err).ToNot(HaveOccurred())
						Expect(b.String()).To(Equal("0"))
					})

					It("should not decrease NonCompliantERC20 balance of the wallet", func() {
						b, err := Wallet.Balance(nil, NonCompliantERC20Address)
						Expect(err).ToNot(HaveOccurred())
						Expect(b.String()).To(Equal("1000"))
					})

					It("should not reduce the available daily spend balance", func() {
						av, err := Wallet.SpendAvailable(nil)
						Expect(err).ToNot(HaveOccurred())
						Expect(av.String()).To(Equal("100000000000000000000"))
					})
				})

			})
		})

    })
	})
})
