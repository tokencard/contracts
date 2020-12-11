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

	var tx *types.Transaction

	Context("when the wallet has enough ETH, the daily limit is 100$ and the rate is 1", func() {
		BeforeEach(func() {
			BankAccount.MustTransfer(Backend, WalletProxyAddress, EthToWei(101))
			BankAccount.MustTransfer(Backend, Controller.Address(), EthToWei(1))

			tx, err := TokenWhitelist.UpdateTokenRate(ControllerAdmin.TransactOpts(), StablecoinAddress, EthToWei(1), big.NewInt(20180913153211))
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeTrue())

			tx, err = WalletProxy.SubmitDailyLimitUpdate(Owner.TransactOpts(), EthToWei(100))
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeTrue())
		})

		It("should emit a UpdatedAvailableDailyLimit event", func() {
			it, err := WalletProxy.FilterUpdatedAvailableDailyLimit(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(it.Next()).To(BeTrue())
			evt := it.Event
			Expect(it.Next()).To(BeFalse())
			Expect(evt.Amount.String()).To(Equal(EthToWei(100).String()))
		})

		Context("When I transfer 100 ETH to a random account", func() {
			BeforeEach(func() {
				tx, err := WalletProxy.Transfer(Owner.TransactOpts(ethertest.WithGasLimit(81000)), RandomAccount.Address(), common.HexToAddress("0x"), EthToWei(100))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})

			It("should emit a UpdatedAvailableDailyLimit event", func() {
				it, err := WalletProxy.FilterUpdatedAvailableDailyLimit(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(it.Next()).To(BeTrue())
				Expect(it.Next()).To(BeTrue())
				evt := it.Event
				Expect(it.Next()).To(BeFalse())
				Expect(evt.Amount.String()).To(Equal("0"))
			})

			When("I transfer 1 extra stablecoin to a randon address", func() {
				BeforeEach(func() {
					tx, err := Stablecoin.Credit(BankAccount.TransactOpts(), WalletProxyAddress, EthToWei(1))
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeTrue())
				})

				It("should fail", func() {
					tx, err := WalletProxy.Transfer(Owner.TransactOpts(ethertest.WithGasLimit(81000)), RandomAccount.Address(), common.HexToAddress("0x"), big.NewInt(1))
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

				It("should update the limit available to 100 USD", func() {
					av, err := WalletProxy.DailyLimitAvailable(nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(av.String()).To(Equal(EthToWei(100).String()))
				})

				When("I transfer 1 ETH (hence 1 USD due to 1-1 rate) to a random account", func() {
					BeforeEach(func() {
						tx, err := WalletProxy.Transfer(Owner.TransactOpts(), RandomAccount.Address(), common.HexToAddress("0x"), EthToWei(1))
						Expect(err).ToNot(HaveOccurred())
						Backend.Commit()
						Expect(isSuccessful(tx)).To(BeTrue())
					})

					It("should emit UpdatedAvailableDailyLimit event", func() {
						it, err := WalletProxy.FilterUpdatedAvailableDailyLimit(nil)
						Expect(err).ToNot(HaveOccurred())
						Expect(it.Next()).To(BeTrue())
						Expect(it.Next()).To(BeTrue())
						Expect(it.Next()).To(BeTrue())
						Expect(it.Next()).To(BeTrue())
						evt := it.Event
						Expect(it.Next()).To(BeFalse())
						Expect(evt.Amount.String()).To(Equal(EthToWei(99).String()))
						initTime := Backend.Blockchain().CurrentHeader().Time + 24*60*60
						Expect(evt.NextReset.String()).To(Equal(big.NewInt(int64(initTime)).String()))
					})

					It("should reduce available amount of 1 USD", func() {
						av, err := WalletProxy.DailyLimitAvailable(nil)
						Expect(err).ToNot(HaveOccurred())
						Expect(av.String()).To(Equal(EthToWei(99).String()))
					})

				})
			})

		})

		Context("When I transfer 1 Finney to a the zero address", func() {
			BeforeEach(func() {
				tx, err := WalletProxy.Transfer(Owner.TransactOpts(ethertest.WithGasLimit(81000)), common.HexToAddress("0x0"), common.HexToAddress("0x"), FinneyToWei(1))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeFalse())
				returnData, _ := ethCall(tx)
				Expect(string(returnData[len(returnData)-64:])).To(ContainSubstring("destination=0"))
			})

			It("should NOT reduce available amount", func() {
				av, err := WalletProxy.DailyLimitAvailable(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(av.String()).To(Equal(EthToWei(100).String()))
			})

		})

		Context("When controller tries to transfer 1 Finney to a random address", func() {
			It("Should fail", func() {
				tx, err := WalletProxy.Transfer(Controller.TransactOpts(ethertest.WithGasLimit(81000)), RandomAccount.Address(), common.HexToAddress("0x"), FinneyToWei(1))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeFalse())
				returnData, _ := ethCall(tx)
				Expect(string(returnData[len(returnData)-64:])).To(ContainSubstring("Only owner or self"))
			})

		})

		Context("When a random address tries to transfer 1 Finney to a random address", func() {
			It("should fail", func() {
				tx, err := WalletProxy.Transfer(RandomAccount.TransactOpts(ethertest.WithGasLimit(81000)), RandomAccount.Address(), common.HexToAddress("0x"), FinneyToWei(1))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeFalse())
				returnData, _ := ethCall(tx)
				Expect(string(returnData[len(returnData)-64:])).To(ContainSubstring("Only owner or self"))
			})
		})

		Context("When I have 1000 stablecoin tokens", func() {
			BeforeEach(func() {
				var err error
				tx, err := Stablecoin.Credit(BankAccount.TransactOpts(), WalletProxyAddress, EthToWei(1000))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})

			Context("When I transfer 100$ to a random address", func() {
				BeforeEach(func() {
					tx, err := WalletProxy.Transfer(Owner.TransactOpts(), RandomAccount.Address(), StablecoinAddress, EthToWei(100))
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeTrue())
				})

				It("should increase Stablecoin balance of the random address", func() {
					b, err := Stablecoin.BalanceOf(nil, RandomAccount.Address())
					Expect(err).ToNot(HaveOccurred())
					Expect(b.String()).To(Equal(EthToWei(100).String()))
				})

				It("should decrease Stablecoin balance of the wallet", func() {
					b, err := Stablecoin.BalanceOf(nil, WalletProxyAddress)
					Expect(err).ToNot(HaveOccurred())
					Expect(b.String()).To(Equal(EthToWei(900).String()))
				})

				It("should reduce the available daily spend balance", func() {
					av, err := WalletProxy.DailyLimitAvailable(nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(av.String()).To(Equal("0"))
				})
			})

			Context("When controller tries to transfer one stablecoin base unit to a random address", func() {
				It("should fail", func() {
					tx, err := WalletProxy.Transfer(Controller.TransactOpts(ethertest.WithGasLimit(80000)), RandomAccount.Address(), StablecoinAddress, big.NewInt(1))
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeFalse())
					returnData, _ := ethCall(tx)
					Expect(string(returnData[len(returnData)-64:])).To(ContainSubstring("Only owner or self"))
				})
			})

			Context("When random address tries to transfer one token to a random address", func() {
				It("should fail", func() {
					tx, err := WalletProxy.Transfer(RandomAccount.TransactOpts(ethertest.WithGasLimit(80000)), RandomAccount.Address(), StablecoinAddress, big.NewInt(1))
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeFalse())
					returnData, _ := ethCall(tx)
					Expect(string(returnData[len(returnData)-64:])).To(ContainSubstring("Only owner or self"))
				})
			})

		})

		Context("When I have one thousand tokens that are not in the TokenWhitelist contract", func() {
			BeforeEach(func() {
				tx, err := ERC20Contract1.Credit(BankAccount.TransactOpts(), WalletProxyAddress, big.NewInt(1000))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())

			})

			Context("When I transfer 300 tokens to a random address", func() {
				BeforeEach(func() {
					tx, err := WalletProxy.Transfer(Owner.TransactOpts(), RandomAccount.Address(), ERC20Contract1Address, big.NewInt(300))
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeTrue())
				})

				It("should increase token balance of the random address", func() {
					b, err := ERC20Contract1.BalanceOf(nil, RandomAccount.Address())
					Expect(err).ToNot(HaveOccurred())
					Expect(b.String()).To(Equal("300"))
				})

				It("should decrease token balance of the wallet", func() {
					b, err := ERC20Contract1.BalanceOf(nil, WalletProxyAddress)
					Expect(err).ToNot(HaveOccurred())
					Expect(b.String()).To(Equal("700"))
				})

				It("should not reduce the available daily balance", func() {
					av, err := WalletProxy.DailyLimitAvailable(nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(av.String()).To(Equal(EthToWei(100).String()))
				})
			})

			When("I have one thousand NonCompliantERC20 tokens that are not in the TokenWhitelist contract", func() {
				BeforeEach(func() {
					var err error
					tx, err = NonCompliantERC20.Credit(BankAccount.TransactOpts(), WalletProxyAddress, big.NewInt(1000))
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeTrue())

				})

				When("I transfer 300 tokens to a random address", func() {
					BeforeEach(func() {
						tx, err := WalletProxy.Transfer(Owner.TransactOpts(), RandomAccount.Address(), NonCompliantERC20Address, big.NewInt(300))
						Expect(err).ToNot(HaveOccurred())
						Backend.Commit()
						Expect(isSuccessful(tx)).To(BeTrue())
					})

					It("should increase token balance of the random address", func() {
						b, err := NonCompliantERC20.BalanceOf(nil, RandomAccount.Address())
						Expect(err).ToNot(HaveOccurred())
						Expect(b.String()).To(Equal("300"))
					})

					It("should decrease token balance of the wallet", func() {
						b, err := NonCompliantERC20.BalanceOf(nil, WalletProxyAddress)
						Expect(err).ToNot(HaveOccurred())
						Expect(b.String()).To(Equal("700"))
					})

					It("should not reduce the available daily balance", func() {
						av, err := WalletProxy.DailyLimitAvailable(nil)
						Expect(err).ToNot(HaveOccurred())
						Expect(av.String()).To(Equal(EthToWei(100).String()))
					})
				})
			})

		})

	})

})
