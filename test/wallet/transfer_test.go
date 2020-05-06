package wallet_test

import (
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tokencard/contracts/v3/test/shared"
	"github.com/tokencard/ethertest"
)

var _ = Describe("transfer", func() {

	Context("when the wallet has enough ETH, the daily limit is 100$ and the rate is 1", func() {
		BeforeEach(func() {
			BankAccount.MustTransfer(Backend, WalletAddress, EthToWei(101))
			BankAccount.MustTransfer(Backend, Controller.Address(), EthToWei(1))
			// rate = 1
			tx, err := TokenWhitelist.UpdateTokenRate(ControllerAdmin.TransactOpts(), StablecoinAddress, EthToWei(1), big.NewInt(20180913153211))
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeTrue())

			tx, err = Wallet.SubmitDailyLimitUpdate(Owner.TransactOpts(), MweiToWei(100))
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeTrue())
		})

		It("should emit a UpdatedAvailableDailyLimit event", func() {
			it, err := Wallet.FilterUpdatedAvailableDailyLimit(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(it.Next()).To(BeTrue())
			Expect(it.Next()).To(BeTrue())
			evt := it.Event
			Expect(it.Next()).To(BeFalse())
			Expect(evt.Amount.String()).To(Equal(MweiToWei(100).String()))
			initTime := Backend.Blockchain().CurrentHeader().Time + 24*60*60 - 40
			Expect(evt.NextReset.String()).To(Equal(big.NewInt(int64(initTime)).String()))
		})

		Context("When I transfer 100 ETH to a randon acount", func() {
			BeforeEach(func() {
				tx, err := Wallet.Transfer(Owner.TransactOpts(ethertest.WithGasLimit(81000)), RandomAccount.Address(), common.HexToAddress("0x"), EthToWei(100))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})

			It("should emit a UpdatedAvailableDailyLimit event", func() {
				it, err := Wallet.FilterUpdatedAvailableDailyLimit(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(it.Next()).To(BeTrue())
				Expect(it.Next()).To(BeTrue())
				Expect(it.Next()).To(BeTrue())
				evt := it.Event
				Expect(it.Next()).To(BeFalse())
				Expect(evt.Amount.String()).To(Equal("0"))
				initTime := Backend.Blockchain().CurrentHeader().Time + 24*60*60 - 50
				Expect(evt.NextReset.String()).To(Equal(big.NewInt(int64(initTime)).String()))
			})

			When("I transfer 1 extra stablecoin base unit to a randon account", func() {
				BeforeEach(func() {
					tx, err := Stablecoin.Credit(BankAccount.TransactOpts(), WalletAddress, MweiToWei(1))
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeTrue())
				})
				It("should fail", func() {
					tx, err := Wallet.Transfer(Owner.TransactOpts(ethertest.WithGasLimit(81000)), RandomAccount.Address(), StablecoinAddress, big.NewInt(1))
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeFalse())
					returnData, _ := ethCall(tx)
					Expect(string(returnData[len(returnData)-64:])).To(ContainSubstring("available<amount"))
				})
			})

			Context("After I wait for 24 hours", func() {
				BeforeEach(func() {
					Backend.AdjustTime(time.Hour*24 + time.Second)
					Backend.Commit()
				})

				It("should update the available amount to 100$", func() {
					av, err := Wallet.DailyLimitAvailable(nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(av.String()).To(Equal(MweiToWei(100).String()))
				})

				When("I transfer 1 ETH (hence 1$ due to 1-1 rate) to a randon account", func() {
					BeforeEach(func() {
						tx, err := Wallet.Transfer(Owner.TransactOpts(), RandomAccount.Address(), common.HexToAddress("0x"), EthToWei(1))
						Expect(err).ToNot(HaveOccurred())
						Backend.Commit()
						Expect(isSuccessful(tx)).To(BeTrue())
					})

					It("should emit a UpdatedAvailableDailyLimit event", func() {
						it, err := Wallet.FilterUpdatedAvailableDailyLimit(nil)
						Expect(err).ToNot(HaveOccurred())
						Expect(it.Next()).To(BeTrue())
						Expect(it.Next()).To(BeTrue())
						Expect(it.Next()).To(BeTrue())
						Expect(it.Next()).To(BeTrue())
						evt := it.Event
						Expect(it.Next()).To(BeFalse())
						Expect(evt.Amount.String()).To(Equal(MweiToWei(99).String()))
						initTime := Backend.Blockchain().CurrentHeader().Time + 24*60*60
						Expect(evt.NextReset.String()).To(Equal(big.NewInt(int64(initTime)).String()))
					})

					It("should reduce available amount 1$", func() {
						av, err := Wallet.DailyLimitAvailable(nil)
						Expect(err).ToNot(HaveOccurred())
						Expect(av.String()).To(Equal(MweiToWei(99).String()))
					})

				})
			})

		})

		Context("When I transfer 1 Finney to a the zero address", func() {
			BeforeEach(func() {
				tx, err := Wallet.Transfer(Owner.TransactOpts(ethertest.WithGasLimit(81000)), common.HexToAddress("0x0"), common.HexToAddress("0x"), FinneyToWei(1))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeFalse())
				returnData, _ := ethCall(tx)
				Expect(string(returnData[len(returnData)-64:])).To(ContainSubstring("destination=0"))
			})

			It("should NOT reduce available amount", func() {
				av, err := Wallet.DailyLimitAvailable(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(av.String()).To(Equal(MweiToWei(100).String()))
			})

		})

		Context("When controller tries to transfer 1 Finney to a random account", func() {

			It("should NOT reduce available amount", func() {
				tx, err := Wallet.Transfer(Controller.TransactOpts(ethertest.WithGasLimit(81000)), RandomAccount.Address(), common.HexToAddress("0x"), FinneyToWei(1))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeFalse())
				returnData, _ := ethCall(tx)
				Expect(string(returnData[len(returnData)-64:])).To(ContainSubstring("Not owner or self"))
			})

		})

		Context("When a random acount tries to transfer 1 Finney to a random account", func() {
			It("should fail", func() {
				tx, err := Wallet.Transfer(RandomAccount.TransactOpts(ethertest.WithGasLimit(81000)), RandomAccount.Address(), common.HexToAddress("0x"), FinneyToWei(1))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeFalse())
				returnData, _ := ethCall(tx)
				Expect(string(returnData[len(returnData)-64:])).To(ContainSubstring("Not owner or self"))
			})
		})

		When("the wallet has one thousand stablecoins", func() {
			BeforeEach(func() {
				tx, err := Stablecoin.Credit(BankAccount.TransactOpts(), WalletAddress, MweiToWei(1000))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})

			When("we transfer 100$ to a random acount", func() {
				BeforeEach(func() {
					tx, err := Wallet.Transfer(Owner.TransactOpts(), RandomAccount.Address(), StablecoinAddress, MweiToWei(100))
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeTrue())
				})

				It("should increase Stablecoin balance of the random acount", func() {
					b, err := Stablecoin.BalanceOf(nil, RandomAccount.Address())
					Expect(err).ToNot(HaveOccurred())
					Expect(b.String()).To(Equal(MweiToWei(100).String()))
				})

				It("should decrease Stablecoin balance of the wallet", func() {
					b, err := Stablecoin.BalanceOf(nil, WalletAddress)
					Expect(err).ToNot(HaveOccurred())
					Expect(b.String()).To(Equal(MweiToWei(900).String()))
				})

				It("should reduce the available daily balance", func() {
					av, err := Wallet.DailyLimitAvailable(nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(av.String()).To(Equal("0"))
				})
			})

			Context("When controller tries to transfer one stablecoin base unit to a random account", func() {
				It("should fail", func() {
					tx, err := Wallet.Transfer(Controller.TransactOpts(ethertest.WithGasLimit(80000)), RandomAccount.Address(), StablecoinAddress, big.NewInt(1))
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeFalse())
					returnData, _ := ethCall(tx)
					Expect(string(returnData[len(returnData)-64:])).To(ContainSubstring("Not owner or self"))
				})
			})

			Context("When random acount tries to transfer one token to a random acount", func() {
				It("should fail", func() {
					tx, err := Wallet.Transfer(RandomAccount.TransactOpts(ethertest.WithGasLimit(80000)), RandomAccount.Address(), StablecoinAddress, big.NewInt(1))
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeFalse())
					returnData, _ := ethCall(tx)
					Expect(string(returnData[len(returnData)-64:])).To(ContainSubstring("Not owner or self"))
				})
			})

		})

		Context("When I have one thousand tokens that are not in the tokenWhitelist", func() {
			BeforeEach(func() {
				tx, err := ERC20Contract1.Credit(BankAccount.TransactOpts(), WalletAddress, big.NewInt(1000))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})

			Context("When I transfer 300 tokens to a random acount", func() {
				BeforeEach(func() {
					tx, err := Wallet.Transfer(Owner.TransactOpts(), RandomAccount.Address(), ERC20Contract1Address, big.NewInt(300))
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeTrue())
				})

				It("should increase token balance of the random acount", func() {
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
					Expect(av.String()).To(Equal(MweiToWei(100).String()))
				})
			})

			When("I have one thousand NonCompliantERC20 tokens that are not in the tokenWhitelist", func() {
				BeforeEach(func() {
					tx, err := NonCompliantERC20.Credit(BankAccount.TransactOpts(), WalletAddress, big.NewInt(1000))
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeTrue())
				})

				When("I transfer 300 tokens to a random acount", func() {
					BeforeEach(func() {
						tx, err := Wallet.Transfer(Owner.TransactOpts(), RandomAccount.Address(), NonCompliantERC20Address, big.NewInt(300))
						Expect(err).ToNot(HaveOccurred())
						Backend.Commit()
						Expect(isSuccessful(tx)).To(BeTrue())
					})

					It("should increase token balance of the random acount", func() {
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
						Expect(av.String()).To(Equal(MweiToWei(100).String()))
					})
				})
			})

		})

	})

})
