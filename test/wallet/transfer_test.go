package wallet_test

import (
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tokencard/contracts/v3/test/shared"
	"github.com/tokencard/ethertest"
)

var _ = Describe("transfer", func() {

	Context("when the wallet has enough ETH", func() {
		BeforeEach(func() {
			BankAccount.MustTransfer(Backend, WalletAddress, EthToWei(101))
			BankAccount.MustTransfer(Backend, Controller.Address(), EthToWei(1))
		})

		var tx *types.Transaction

		Context("When I transfer 100 ETH to a randon person", func() {
			BeforeEach(func() {
				var err error
				tx, err = Wallet.Transfer(Owner.TransactOpts(ethertest.WithGasLimit(81000)), RandomAccount.Address(), common.HexToAddress("0x"), EthToWei(100))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
			})
			It("should succeed", func() {
				Expect(isSuccessful(tx)).To(BeTrue())
			})

			It("should reduce available transfer for today by 100 ETH", func() {
				av, err := Wallet.DailyLimitAvailable(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(av.String()).To(Equal("0"))
			})

			It("should update the available amount to 0", func() {
				av, err := Wallet.DailyLimitAvailable(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(av.String()).To(Equal("0"))
			})

			When("I transfer 1 extra wei to a randon person", func() {
				It("should fail", func() {
					var err error
					tx, err = Wallet.Transfer(Owner.TransactOpts(ethertest.WithGasLimit(81000)), RandomAccount.Address(), common.HexToAddress("0x"), big.NewInt(1))
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeFalse())
				})
			})

			Context("After I wait for 24 hours", func() {
				BeforeEach(func() {
					Backend.AdjustTime(time.Hour*24 + time.Second)
					Backend.Commit()
				})

				It("should update the available amount to 100 ETH", func() {
					av, err := Wallet.DailyLimitAvailable(nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(av.String()).To(Equal(EthToWei(100).String()))
				})

				When("I transfer 1 Finney  to a randon person", func() {
					BeforeEach(func() {
						var err error
						tx, err = Wallet.Transfer(Owner.TransactOpts(ethertest.WithGasLimit(81000)), RandomAccount.Address(), common.HexToAddress("0x"), FinneyToWei(1))
						Expect(err).ToNot(HaveOccurred())
						Backend.Commit()
					})

					It("should succeed", func() {
						Expect(isSuccessful(tx)).To(BeTrue())
					})

					It("should emit UpdatedAvailableDailyLimit event", func() {
						it, err := Wallet.FilterUpdatedAvailableDailyLimit(nil)
						Expect(err).ToNot(HaveOccurred())
						Expect(it.Next()).To(BeTrue())
						_ = it.Event
						Expect(it.Next()).To(BeFalse())
					})

					It("should reduce available transfer for today by 1 Finney", func() {
						av, err := Wallet.DailyLimitAvailable(nil)
						Expect(err).ToNot(HaveOccurred())
						Expect(av.String()).To(Equal("99999000000000000000"))
					})

				})
			})

		})

		Context("When I transfer 1 Finney to a randon person", func() {
			BeforeEach(func() {
				var err error
				tx, err = Wallet.Transfer(Owner.TransactOpts(ethertest.WithGasLimit(81000)), RandomAccount.Address(), common.HexToAddress("0x"), FinneyToWei(1))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
			})

			It("should succeed", func() {
				Expect(isSuccessful(tx)).To(BeTrue())
			})

			It("should reduce available transfer for today by 1 Finney", func() {
				av, err := Wallet.DailyLimitAvailable(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(av.String()).To(Equal("99999000000000000000"))
			})

		})

		Context("When I transfer 1 Finney to a the zero address", func() {
			BeforeEach(func() {
				var err error
				tx, err = Wallet.Transfer(Owner.TransactOpts(ethertest.WithGasLimit(81000)), common.HexToAddress("0x0"), common.HexToAddress("0x"), FinneyToWei(1))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
			})
			It("should fail", func() {
				Expect(isSuccessful(tx)).To(BeFalse())
			})

			It("should reduce available transfer for today by 0 Finney", func() {
				av, err := Wallet.DailyLimitAvailable(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(av.String()).To(Equal("100000000000000000000"))
			})

		})

		Context("When controller tries to transfer 1 Finney to a random person", func() {

			BeforeEach(func() {
				var err error
				tx, err = Wallet.Transfer(Controller.TransactOpts(ethertest.WithGasLimit(81000)), RandomAccount.Address(), common.HexToAddress("0x"), FinneyToWei(1))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
			})

			It("Should fail", func() {
				Expect(isSuccessful(tx)).To(BeFalse())
			})

		})

		Context("When a random person tries to transfer 1 Finney to a random person", func() {

			BeforeEach(func() {
				var err error
				tx, err = Wallet.Transfer(RandomAccount.TransactOpts(ethertest.WithGasLimit(81000)), RandomAccount.Address(), common.HexToAddress("0x"), FinneyToWei(1))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
			})

			It("should fail", func() {
				Expect(isSuccessful(tx)).To(BeFalse())
			})

		})

		Context("When I have one thousand tokens", func() {
			BeforeEach(func() {
				var err error
				tx, err = TKNBurner.Mint(BankAccount.TransactOpts(), WalletAddress, big.NewInt(1000))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())

			})

			Context("When I transfer 300 tokens to a random person", func() {
				BeforeEach(func() {
					tx, err := Wallet.Transfer(Owner.TransactOpts(), RandomAccount.Address(), TKNBurnerAddress, big.NewInt(300))
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeTrue())
				})

				It("should increase TKN balance of the random person", func() {
					b, err := TKNBurner.BalanceOf(nil, RandomAccount.Address())
					Expect(err).ToNot(HaveOccurred())
					Expect(b.String()).To(Equal("300"))
				})

				It("should decrease TKN balance of the wallet", func() {
                    b, err := TKNBurner.BalanceOf(nil, WalletAddress)
					Expect(err).ToNot(HaveOccurred())
					Expect(b.String()).To(Equal("700"))
				})

				It("should reduce the available daily balance", func() {
					av, err := Wallet.DailyLimitAvailable(nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(av.String()).To(AlmostEqual("99999999999951010000"))
				})
			})

			Context("When controller tries to transfer one token to a random person", func() {
				BeforeEach(func() {
					var err error
					tx, err = Wallet.Transfer(Controller.TransactOpts(ethertest.WithGasLimit(80000)), RandomAccount.Address(), TKNBurnerAddress, big.NewInt(1))
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
				})

				It("should fail", func() {
					Expect(isSuccessful(tx)).To(BeFalse())
				})
			})

			Context("When random person tries to transfer one token to a random person", func() {

				BeforeEach(func() {
					var err error
					tx, err = Wallet.Transfer(RandomAccount.TransactOpts(ethertest.WithGasLimit(80000)), RandomAccount.Address(), TKNBurnerAddress, big.NewInt(1))
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
				})

				It("should fail", func() {
					Expect(isSuccessful(tx)).To(BeFalse())
				})
			})

		})

		Context("When I have one thousand tokens that are not whitelisted", func() {
			BeforeEach(func() {
				var err error
				tx, err = ERC20Contract1.Credit(BankAccount.TransactOpts(), WalletAddress, big.NewInt(1000))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())

			})

			Context("When I transfer 300 tokens to a random person", func() {
				BeforeEach(func() {
					tx, err := Wallet.Transfer(Owner.TransactOpts(), RandomAccount.Address(), ERC20Contract1Address, big.NewInt(300))
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeTrue())
				})

				It("should increase token balance of the random person", func() {
					b, err := ERC20Contract1.BalanceOf(nil, RandomAccount.Address())
					Expect(err).ToNot(HaveOccurred())
					Expect(b.String()).To(Equal("300"))
				})

				It("should decrease token balance of the wallet", func() {
                    b, err := ERC20Contract1.BalanceOf(nil, WalletAddress)
					Expect(err).ToNot(HaveOccurred())
					Expect(b.String()).To(Equal("700"))
				})

				It("should not reduce the available daily balance", func() {
					av, err := Wallet.DailyLimitAvailable(nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(av.String()).To(Equal("100000000000000000000"))
				})
			})

			When("I have one thousand NonCompliantERC20 tokens that are not whitelisted", func() {
				BeforeEach(func() {
					var err error
					tx, err = NonCompliantERC20.Credit(BankAccount.TransactOpts(), WalletAddress, big.NewInt(1000))
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeTrue())

				})

				When("I transfer 300 tokens to a random person", func() {
					BeforeEach(func() {
						tx, err := Wallet.Transfer(Owner.TransactOpts(), RandomAccount.Address(), NonCompliantERC20Address, big.NewInt(300))
						Expect(err).ToNot(HaveOccurred())
						Backend.Commit()
						Expect(isSuccessful(tx)).To(BeTrue())
					})

					It("should increase token balance of the random person", func() {
						b, err := NonCompliantERC20.BalanceOf(nil, RandomAccount.Address())
						Expect(err).ToNot(HaveOccurred())
						Expect(b.String()).To(Equal("300"))
					})

					It("should decrease token balance of the wallet", func() {
                        b, err := NonCompliantERC20.BalanceOf(nil, WalletAddress)
						Expect(err).ToNot(HaveOccurred())
						Expect(b.String()).To(Equal("700"))
					})

					It("should not reduce the available daily balance", func() {
						av, err := Wallet.DailyLimitAvailable(nil)
						Expect(err).ToNot(HaveOccurred())
						Expect(av.String()).To(Equal("100000000000000000000"))
					})
				})
			})

			Context("When controller tries to transfer one token to a random person", func() {
				BeforeEach(func() {
					var err error
					tx, err = Wallet.Transfer(Controller.TransactOpts(ethertest.WithGasLimit(80000)), RandomAccount.Address(), ERC20Contract1Address, big.NewInt(1))
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
				})

				It("should fail", func() {
					Expect(isSuccessful(tx)).To(BeFalse())
				})
			})

			Context("When random person tries to transfer one token to a random person", func() {

				BeforeEach(func() {
					var err error
					tx, err = Wallet.Transfer(RandomAccount.TransactOpts(ethertest.WithGasLimit(80000)), RandomAccount.Address(), ERC20Contract1Address, big.NewInt(1))
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
				})

				It("should fail", func() {
					Expect(isSuccessful(tx)).To(BeFalse())
				})

			})

		})

	})

})
