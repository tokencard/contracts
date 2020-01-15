package wallet_test

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tokencard/contracts/v2/test/shared"
	"github.com/tokencard/ethertest"
)

var _ = Describe("bulk_transfer", func() {

	When("the wallet has enough ETH and a DailyLimit of 2000$", func() {
		BeforeEach(func() {
			BankAccount.MustTransfer(Backend, WalletAddress, EthToWei(1))
			BankAccount.MustTransfer(Backend, Controller.Address(), EthToWei(1))
		})

		BeforeEach(func() {
			tx, err := ERC20Contract1.Credit(BankAccount.TransactOpts(), WalletAddress, big.NewInt(1000))
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeTrue())
		})

		BeforeEach(func() {
			tx, err := ERC20Contract2.Credit(BankAccount.TransactOpts(), WalletAddress, big.NewInt(500))
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeTrue())
		})

		BeforeEach(func() {
			tx, err := Wallet.SubmitDailyLimitUpdate(Owner.TransactOpts(), EthToWei(2000))
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeTrue())
		})

		It("should have a DailyLimit of 2000$", func() {
			av, err := Wallet.DailyLimitAvailable(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(av.String()).To(Equal(EthToWei(2000).String()))
		})

		It("the balance of the wallet should be 1 eth", func() {
			b, e := Backend.BalanceAt(context.Background(), WalletAddress, nil)
			Expect(e).ToNot(HaveOccurred())
			Expect(b.String()).To(Equal(EthToWei(1).String()))
		})

		It("should increase the ERC20 type-1 balance of the Wallet by 1000", func() {
			b, err := ERC20Contract1.BalanceOf(nil, WalletAddress)
			Expect(err).ToNot(HaveOccurred())
			Expect(b.String()).To(Equal("1000"))
		})

		It("should increase the ERC20 type-2 balance of the Wallet by 500", func() {
			b, err := ERC20Contract2.BalanceOf(nil, WalletAddress)
			Expect(err).ToNot(HaveOccurred())
			Expect(b.String()).To(Equal("500"))
		})

		var tx *types.Transaction
		When("called by the Owner", func() {
			When("the size of the array of assets is zero", func() {
				It("should fail", func() {
					tx, err := Wallet.BulkTransfer(Owner.TransactOpts(ethertest.WithGasLimit(81000)), RandomAccount.Address(), []common.Address{})
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isGasExhausted(tx, 81000)).To(BeFalse())
					Expect(isSuccessful(tx)).To(BeFalse())
				})
			})

			When("the size of the array of assets is greater than zero", func() {

				var randomBalance *big.Int
				var assets []common.Address

				BeforeEach(func() {
					var err error
					randomBalance, err = Backend.BalanceAt(context.Background(), RandomAccount.Address(), nil)
					Expect(err).ToNot(HaveOccurred())
					assets = []common.Address{common.HexToAddress("0x0"), ERC20Contract1Address, ERC20Contract2Address}
				})

				When("wallet balance <= DailyLimit", func() {
					When("the tokens are not in the TokeWhitelist i.e. not available, they are treated as transfering to whitelisted addresses", func() {
						//Balance is now 2 ETH
						BeforeEach(func() {
							BankAccount.MustTransfer(Backend, WalletAddress, EthToWei(1))
						})

						BeforeEach(func() {
							var err error
							tx, err = Wallet.BulkTransfer(Owner.TransactOpts(), RandomAccount.Address(), assets)
							Expect(err).ToNot(HaveOccurred())
							Backend.Commit()
							Expect(isSuccessful(tx)).To(BeTrue())
						})

						It("the balance of the wallet should be 0", func() {
							b, e := Backend.BalanceAt(context.Background(), WalletAddress, nil)
							Expect(e).ToNot(HaveOccurred())
							Expect(b.String()).To(Equal("0"))
						})
						It("should decrease the ERC20 type-1 balance of the Wallet by 1000", func() {
							b, err := ERC20Contract1.BalanceOf(nil, WalletAddress)
							Expect(err).ToNot(HaveOccurred())
							Expect(b.String()).To(Equal("0"))
						})

						It("should decrease the ERC20 type-2 balance of the Wallet by 500", func() {
							b, err := ERC20Contract2.BalanceOf(nil, WalletAddress)
							Expect(err).ToNot(HaveOccurred())
							Expect(b.String()).To(Equal("0"))
						})
						It("should increase the balance of the RandomAccount by 2 eth", func() {
							newBalance, err := Backend.BalanceAt(context.Background(), RandomAccount.Address(), nil)
							Expect(err).ToNot(HaveOccurred())
							randomBalance.Add(randomBalance, EthToWei(2))
							Expect(newBalance.String()).To(Equal(randomBalance.String()))
						})
						It("should increase the ERC20 type-1 balance of the RandomAccount by 1000", func() {
							b, err := ERC20Contract1.BalanceOf(nil, RandomAccount.Address())
							Expect(err).ToNot(HaveOccurred())
							Expect(b.String()).To(Equal("1000"))
						})
						It("should increase the ERC20 type-2 balance of the RandomAccount by 500", func() {
							b, err := ERC20Contract2.BalanceOf(nil, RandomAccount.Address())
							Expect(err).ToNot(HaveOccurred())
							Expect(b.String()).To(Equal("500"))
						})
						It("Should emit a BulkTransferred event", func() {
							it, err := Wallet.FilterBulkTransferred(nil)
							Expect(err).ToNot(HaveOccurred())
							Expect(it.Next()).To(BeTrue())
							evt := it.Event
							Expect(it.Next()).To(BeFalse())
							Expect(evt.To).To(Equal(RandomAccount.Address()))
							Expect(evt.Assets).To(Equal(assets))
						})

						It("Should emit 3 Transferred events", func() {
							it, err := Wallet.FilterTransferred(nil)
							Expect(err).ToNot(HaveOccurred())
							Expect(it.Next()).To(BeTrue())
							evt := it.Event
							Expect(it.Next()).To(BeTrue())
							Expect(evt.To).To(Equal(RandomAccount.Address()))
							Expect(evt.Asset).To(Equal(common.HexToAddress("0x0")))
							Expect(evt.Amount.String()).To(Equal(EthToWei(2).String()))
							evt = it.Event
							Expect(it.Next()).To(BeTrue())
							Expect(evt.To).To(Equal(RandomAccount.Address()))
							Expect(evt.Asset).To(Equal(ERC20Contract1Address))
							Expect(evt.Amount.String()).To(Equal("1000"))
							evt = it.Event
							Expect(it.Next()).To(BeFalse())
							Expect(evt.To).To(Equal(RandomAccount.Address()))
							Expect(evt.Asset).To(Equal(ERC20Contract2Address))
							Expect(evt.Amount.String()).To(Equal("500"))
						})

					})
					When("the tokens are added to the oracle", func() {

						BeforeEach(func() {
							tokens := []common.Address{ERC20Contract1Address, ERC20Contract2Address}
							tx, err := TokenWhitelist.AddTokens(
								ControllerAdmin.TransactOpts(),
								tokens,
								StringsToByte32(
									"ERC1",
									"ERC2",
								),
								[]*big.Int{
									DecimalsToMagnitude(big.NewInt(18)),
									DecimalsToMagnitude(big.NewInt(18)),
								},
								[]bool{true, true},
								[]bool{true, true},
								big.NewInt(20180913153211),
							)
							Expect(err).ToNot(HaveOccurred())
							Backend.Commit()
							Expect(isSuccessful(tx)).To(BeTrue())
						})
						When("rates are not updated (=0)", func() {
							It("Should revert", func() {
								tx, err := Wallet.BulkTransfer(Owner.TransactOpts(ethertest.WithGasLimit(1000000)), RandomAccount.Address(), assets)
								Expect(err).ToNot(HaveOccurred())
								Backend.Commit()
								Expect(isSuccessful(tx)).To(BeFalse())
                                returnData, _ := ethCall(tx)
                				Expect(string(returnData[len(returnData)-64:])).To(ContainSubstring("rate=0"))
							})
						})
						When("rates are updated with valid rates", func() {
							BeforeEach(func() {
								tx, err := TokenWhitelist.UpdateTokenRate(
									ControllerAdmin.TransactOpts(),
									ERC20Contract1Address,
									big.NewInt(2),
									big.NewInt(20180913153211),
								)
								Expect(err).ToNot(HaveOccurred())
								Backend.Commit()
								Expect(isSuccessful(tx)).To(BeTrue())
							})
							BeforeEach(func() {
								tx, err := TokenWhitelist.UpdateTokenRate(
									ControllerAdmin.TransactOpts(),
									ERC20Contract2Address,
									big.NewInt(1),
									big.NewInt(20180913153211),
								)
								Expect(err).ToNot(HaveOccurred())
								Backend.Commit()
								Expect(isSuccessful(tx)).To(BeTrue())
							})
							BeforeEach(func() {
								var err error
								tx, err = Wallet.BulkTransfer(Owner.TransactOpts(), RandomAccount.Address(), assets)
								Expect(err).ToNot(HaveOccurred())
								Backend.Commit()
								Expect(isSuccessful(tx)).To(BeTrue())
							})

							It("the balance of the wallet should be 0", func() {
								b, e := Backend.BalanceAt(context.Background(), WalletAddress, nil)
								Expect(e).ToNot(HaveOccurred())
								Expect(b.String()).To(Equal("0"))
							})
							It("should decrease the ERC20 type-1 balance of the Wallet by 1000", func() {
								b, err := ERC20Contract1.BalanceOf(nil, WalletAddress)
								Expect(err).ToNot(HaveOccurred())
								Expect(b.String()).To(Equal("0"))
							})

							It("should decrease the ERC20 type-2 balance of the Wallet by 500", func() {
								b, err := ERC20Contract2.BalanceOf(nil, WalletAddress)
								Expect(err).ToNot(HaveOccurred())
								Expect(b.String()).To(Equal("0"))
							})
						}) // valid rates, transfer does not exceed the DailyLimit

						When("rates are updated but the DailyLimit is ecxeeded", func() {
							BeforeEach(func() {
								tx, err := TokenWhitelist.UpdateTokenRate(
									ControllerAdmin.TransactOpts(),
									ERC20Contract1Address,
									EthToWei(100000000000000000), //10^17
									big.NewInt(20180913153211),
								)
								Expect(err).ToNot(HaveOccurred())
								Backend.Commit()
								Expect(isSuccessful(tx)).To(BeTrue())
							})
							BeforeEach(func() {
								tx, err := TokenWhitelist.UpdateTokenRate(
									ControllerAdmin.TransactOpts(),
									ERC20Contract2Address,
									EthToWei(1),
									big.NewInt(20180913153211),
								)
								Expect(err).ToNot(HaveOccurred())
								Backend.Commit()
								Expect(isSuccessful(tx)).To(BeTrue())
							})
							When("the destination address is not whitelisted", func() {
								It("Should revert", func() {
									tx, err := Wallet.BulkTransfer(Owner.TransactOpts(ethertest.WithGasLimit(1000000)), RandomAccount.Address(), assets)
									Expect(err).ToNot(HaveOccurred())
									Backend.Commit()
									Expect(isSuccessful(tx)).To(BeFalse())
                                    returnData, _ := ethCall(tx)
                    				Expect(string(returnData[len(returnData)-64:])).To(ContainSubstring("available<amount"))
								})
							})

							When("the 'sent' address is  whitelisted", func() {
								BeforeEach(func() {
									tx, err := Wallet.SetWhitelist(Owner.TransactOpts(), []common.Address{RandomAccount.Address()})
									Expect(err).ToNot(HaveOccurred())
									Backend.Commit()
									Expect(isSuccessful(tx)).To(BeTrue())
								})
								BeforeEach(func() {
									tx, err := Wallet.BulkTransfer(Owner.TransactOpts(), RandomAccount.Address(), assets)
									Expect(err).ToNot(HaveOccurred())
									Backend.Commit()
									Expect(isSuccessful(tx)).To(BeTrue())
								})
								It("the balance of the wallet should be 0", func() {
									b, e := Backend.BalanceAt(context.Background(), WalletAddress, nil)
									Expect(e).ToNot(HaveOccurred())
									Expect(b.String()).To(Equal("0"))
								})
								It("should decrease the ERC20 type-1 balance of the Wallet by 1000", func() {
									b, err := ERC20Contract1.BalanceOf(nil, WalletAddress)
									Expect(err).ToNot(HaveOccurred())
									Expect(b.String()).To(Equal("0"))
								})

								It("should decrease the ERC20 type-2 balance of the Wallet by 500", func() {
									b, err := ERC20Contract2.BalanceOf(nil, WalletAddress)
									Expect(err).ToNot(HaveOccurred())
									Expect(b.String()).To(Equal("0"))
								})

							})
						})

					}) //the tokens are added
				})

			})
		}) //called by the owner

		When("called by a random address", func() {
			It("should fail", func() {
				tx, err := Wallet.BulkTransfer(RandomAccount.TransactOpts(ethertest.WithGasLimit(81000)), RandomAccount.Address(), []common.Address{common.HexToAddress("0x0"), common.HexToAddress("0x2"), common.HexToAddress("0x3")})
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isGasExhausted(tx, 81000)).To(BeFalse())
				Expect(isSuccessful(tx)).To(BeFalse())
			})
		})
	})
})
