package holder_test

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tokencard/contracts/test/shared"

	"github.com/ethereum/go-ethereum/common"
)

var _ = Describe("TokenHolder", func() {

	Context("A valid Hodler wants to burrrrrnnnn TKN", func() {

		var tx *types.Transaction

		When("one thousand TKN are credited(minted) to an external address", func() {

			BeforeEach(func() {
				tx, err := TKNBurner.Mint(Owner.TransactOpts(), BankAccount.Address(), big.NewInt(1000))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})

			It("should increase the total TKN supply to 1000", func() {
				s, err := TKNBurner.TotalSupply(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(s.String()).To(Equal("1000"))
			})

			It("should increase TKN balance of the specific address", func() {
				b, err := TKNBurner.BalanceOf(nil, BankAccount.Address())
				Expect(err).ToNot(HaveOccurred())
				Expect(b.String()).To(Equal("1000"))
			})

			It("Should emit a Transfer event", func() {
				from := []common.Address{common.HexToAddress("0x0")}
				to := []common.Address{BankAccount.Address()}
				it, err := TKNBurner.FilterTransfer(nil, from, to)
				Expect(err).ToNot(HaveOccurred())
				Expect(it.Next()).To(BeTrue())
				evt := it.Event
				Expect(it.Next()).To(BeFalse())
				Expect(evt.From).To(Equal(common.HexToAddress("0x0")))
				Expect(evt.To).To(Equal(BankAccount.Address()))
				Expect(evt.Value.String()).To(Equal("1000"))
			})

			When("300 TKN are transfered to a random address", func() {

				BeforeEach(func() {
					tx, err := TKNBurner.Transfer(BankAccount.TransactOpts(), RandomAccount.Address(), big.NewInt(300))
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeTrue())
				})

				It("should increase TKN balance of the random address", func() {
					b, err := TKNBurner.BalanceOf(nil, RandomAccount.Address())
					Expect(err).ToNot(HaveOccurred())
					Expect(b.String()).To(Equal("300"))
				})

				It("should decrease TKN balance of the initial address", func() {
					b, err := TKNBurner.BalanceOf(nil, BankAccount.Address())
					Expect(err).ToNot(HaveOccurred())
					Expect(b.String()).To(Equal("700"))
				})

				It("should leave the total supply intact", func() {
					s, err := TKNBurner.TotalSupply(nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(s.String()).To(Equal("1000"))
				})

				It("Should emit a Transfer event", func() {
					from := []common.Address{BankAccount.Address()}
					to := []common.Address{RandomAccount.Address()}
					it, err := TKNBurner.FilterTransfer(nil, from, to)
					Expect(err).ToNot(HaveOccurred())
					Expect(it.Next()).To(BeTrue())
					evt := it.Event
					Expect(it.Next()).To(BeFalse())
					Expect(evt.From).To(Equal(BankAccount.Address()))
					Expect(evt.To).To(Equal(RandomAccount.Address()))
					Expect(evt.Value.String()).To(Equal("300"))
				})

				When("The holder contract owns 3 types of ERC20 tokens (2/3 are redeemable) and 1 ETH", func() {

					//add the tokens to the list
					BeforeEach(func() {
						tokens := []common.Address{common.HexToAddress("0x0"), ERC20Contract1Address, ERC20Contract2Address, ERC20Contract3Address}
						tx, err := TokenWhitelist.AddTokens(
							ControllerAdmin.TransactOpts(),
							tokens,
							StringsToByte32(
								"ETH",
								"ERC1",
								"ERC2",
								"ERC3",
							),
							[]*big.Int{
								DecimalsToMagnitude(big.NewInt(18)),
								DecimalsToMagnitude(big.NewInt(18)),
								DecimalsToMagnitude(big.NewInt(18)),
								DecimalsToMagnitude(big.NewInt(18)),
							},
							[]bool{true, true, true, true},
							[]bool{true, true, true, false},
							big.NewInt(20180913153211),
						)
						Expect(err).ToNot(HaveOccurred())
						Backend.Commit()
						Expect(isSuccessful(tx)).To(BeTrue())
					})

					BeforeEach(func() {
						var err error
						tx, err = ERC20Contract1.Credit(BankAccount.TransactOpts(), TokenHolderAddress, big.NewInt(1000))
						Expect(err).ToNot(HaveOccurred())
						Backend.Commit()
						Expect(isSuccessful(tx)).To(BeTrue())
					})

					BeforeEach(func() {
						var err error
						tx, err = ERC20Contract2.Credit(BankAccount.TransactOpts(), TokenHolderAddress, big.NewInt(500))
						Expect(err).ToNot(HaveOccurred())
						Backend.Commit()
						Expect(isSuccessful(tx)).To(BeTrue())

					})

					BeforeEach(func() {
						var err error
						tx, err = ERC20Contract3.Credit(BankAccount.TransactOpts(), TokenHolderAddress, big.NewInt(1000))
						Expect(err).ToNot(HaveOccurred())
						Backend.Commit()
						Expect(isSuccessful(tx)).To(BeTrue())

					})

					BeforeEach(func() {
						BankAccount.MustTransfer(Backend, TokenHolderAddress, EthToWei(1))
					})

					It("should increase the ETH balance of the holder contract address by 1 ETH", func() {
						b, e := Backend.BalanceAt(context.Background(), TokenHolderAddress, nil)
						Expect(e).ToNot(HaveOccurred())
						bStr := b.String()
						Expect(bStr).To(Equal(EthToWei(1).String()))
					})

					It("should increase the ERC20 type-1 balance of the holder contract by 1000", func() {
						b, err := ERC20Contract1.BalanceOf(nil, TokenHolderAddress)
						Expect(err).ToNot(HaveOccurred())
						Expect(b.String()).To(Equal("1000"))
					})

					It("should increase the ERC20 type-2 balance of the holder contract by 500", func() {
						b, err := ERC20Contract2.BalanceOf(nil, TokenHolderAddress)
						Expect(err).ToNot(HaveOccurred())
						Expect(b.String()).To(Equal("500"))
					})

					It("should increase the ERC20 type-3 balance of the holder contract by 1000", func() {
						b, err := ERC20Contract3.BalanceOf(nil, TokenHolderAddress)
						Expect(err).ToNot(HaveOccurred())
						Expect(b.String()).To(Equal("1000"))
					})

					When("When the token holder (contract) is set by the owner", func() {

						BeforeEach(func() {
							tx, err := TKNBurner.SetTokenHolder(Owner.TransactOpts(), TokenHolderAddress)
							Expect(err).ToNot(HaveOccurred())
							Backend.Commit()
							Expect(isSuccessful(tx)).To(BeTrue())
						})

						When("When the random address tries to burn more TKN than it owns", func() {

							var initialBalance *big.Int
							txCost := new(big.Int)

							BeforeEach(func() {
								var err error
								initialBalance, err = Backend.BalanceAt(context.Background(), RandomAccount.Address(), nil)
								Expect(err).ToNot(HaveOccurred())
							})

							BeforeEach(func() {
								tx, err := TKNBurner.Burn(RandomAccount.TransactOpts(), big.NewInt(301))
								Expect(err).ToNot(HaveOccurred())
								Backend.Commit()
								Expect(isSuccessful(tx)).To(BeTrue())
								r, err := Backend.TransactionReceipt(context.Background(), tx.Hash())
								Expect(err).ToNot(HaveOccurred())
								txCost.Mul(tx.GasPrice(), big.NewInt(int64(r.GasUsed)))
							})

							It("should NOT have an impact on the (ETH) balance of the holder contract", func() {
								b, e := Backend.BalanceAt(context.Background(), TokenHolderAddress, nil)
								Expect(e).ToNot(HaveOccurred())
								finalBal := EthToWei(1)
								Expect(b.String()).To(Equal(finalBal.String()))
							})

							It("should leave the total supply intact", func() {
								s, err := TKNBurner.TotalSupply(nil)
								Expect(err).ToNot(HaveOccurred())
								Expect(s.String()).To(Equal("1000"))
							})

							It("should leave the TKN balance of the random address intact", func() {
								b, err := TKNBurner.BalanceOf(nil, RandomAccount.Address())
								Expect(err).ToNot(HaveOccurred())
								Expect(b.String()).To(Equal("300"))
							})

							It("should leave the ERC20 type-1 balance of the holder contract intact", func() {
								b, err := ERC20Contract1.BalanceOf(nil, TokenHolderAddress)
								Expect(err).ToNot(HaveOccurred())
								Expect(b.String()).To(Equal("1000"))
							})

							It("should leave the ERC20 type-2 balance of the holder contract intact", func() {
								b, err := ERC20Contract2.BalanceOf(nil, TokenHolderAddress)
								Expect(err).ToNot(HaveOccurred())
								Expect(b.String()).To(Equal("500"))
							})

							It("Should NOT emit a Transfer event", func() {
								from := []common.Address{RandomAccount.Address()}
								to := []common.Address{common.HexToAddress("0x0")}
								it, err := TKNBurner.FilterTransfer(nil, from, to)
								Expect(err).ToNot(HaveOccurred())
								Expect(it.Next()).To(BeFalse())
							})

							It("should NOT increase the ERC20 type-1 balance of the random address", func() {
								b, err := ERC20Contract1.BalanceOf(nil, RandomAccount.Address())
								Expect(err).ToNot(HaveOccurred())
								Expect(b.String()).To(Equal("0"))
							})

							It("shouldn't increase the ERC20 type-2 balance of the random address", func() {
								b, err := ERC20Contract2.BalanceOf(nil, RandomAccount.Address())
								Expect(err).ToNot(HaveOccurred())
								Expect(b.String()).To(Equal("0"))
							})

							It("should NOT increase the ETH balance of the random address", func() {
								b, e := Backend.BalanceAt(context.Background(), RandomAccount.Address(), nil)
								Expect(e).ToNot(HaveOccurred())
								bStr := b.String()
								newBalance := initialBalance
								newBalance.Sub(initialBalance, txCost) //we have to deduct the tx cost
								Expect(bStr).To(Equal(newBalance.String()))
							})

						})

						When("When the random address burns 200 tokens (out of 1000)", func() {

							var initialBalance *big.Int
							txCost := new(big.Int)

							BeforeEach(func() {
								var err error
								initialBalance, err = Backend.BalanceAt(context.Background(), RandomAccount.Address(), nil)
								Expect(err).ToNot(HaveOccurred())
							})

							BeforeEach(func() {
								tx, err := TKNBurner.Burn(RandomAccount.TransactOpts(), big.NewInt(200))
								Expect(err).ToNot(HaveOccurred())
								Backend.Commit()
								Expect(isSuccessful(tx)).To(BeTrue())
								r, err := Backend.TransactionReceipt(context.Background(), tx.Hash())
								Expect(err).ToNot(HaveOccurred())
								txCost.Mul(tx.GasPrice(), big.NewInt(int64(r.GasUsed)))
							})

							It("should decrease the (ETH) balance of the holder contract by 20%", func() {
								b, e := Backend.BalanceAt(context.Background(), TokenHolderAddress, nil)
								Expect(e).ToNot(HaveOccurred())
								finalBal := EthToWei(8)
								finalBal.Div(finalBal, big.NewInt(10)) //the final balance should be 0.8 ETH -> 8/10
								Expect(b.String()).To(Equal(finalBal.String()))
							})

							It("should decrease the total supply by 200 (800 remaining)", func() {
								s, err := TKNBurner.TotalSupply(nil)
								Expect(err).ToNot(HaveOccurred())
								Expect(s.String()).To(Equal("800"))
							})

							It("should decrease TKN balance of the random address", func() {
								b, err := TKNBurner.BalanceOf(nil, RandomAccount.Address())
								Expect(err).ToNot(HaveOccurred())
								Expect(b.String()).To(Equal("100"))
							})

							It("should decrease the ERC20 type-1 balance of the holder contract by 200", func() {
								b, err := ERC20Contract1.BalanceOf(nil, TokenHolderAddress)
								Expect(err).ToNot(HaveOccurred())
								Expect(b.String()).To(Equal("800"))
							})

							It("should decrease the ERC20 type-2 balance of the holder contract by 100", func() {
								b, err := ERC20Contract2.BalanceOf(nil, TokenHolderAddress)
								Expect(err).ToNot(HaveOccurred())
								Expect(b.String()).To(Equal("400"))
							})

							It("should NOT decrease the ERC20 type-3 balance of the holder contract", func() {
								b, err := ERC20Contract3.BalanceOf(nil, TokenHolderAddress)
								Expect(err).ToNot(HaveOccurred())
								Expect(b.String()).To(Equal("1000"))
							})

							It("Should emit a Transfer event", func() {
								from := []common.Address{RandomAccount.Address()}
								to := []common.Address{common.HexToAddress("0x0")}
								it, err := TKNBurner.FilterTransfer(nil, from, to)
								Expect(err).ToNot(HaveOccurred())
								Expect(it.Next()).To(BeTrue())
								evt := it.Event
								Expect(it.Next()).To(BeFalse())
								Expect(evt.From).To(Equal(RandomAccount.Address()))
								Expect(evt.To).To(Equal(common.HexToAddress("0x0")))
								Expect(evt.Value.String()).To(Equal("200"))
							})

							It("Should emit 3 Claimed events", func() {
								it, err := TokenHolder.FilterCashAndBurned(nil)
								Expect(err).ToNot(HaveOccurred())
								Expect(it.Next()).To(BeTrue())
								evt := it.Event
								Expect(it.Next()).To(BeTrue())
								Expect(evt.To).To(Equal(RandomAccount.Address()))
								Expect(evt.Asset).To(Equal(common.HexToAddress("0x0")))
								Expect(evt.Amount.String()).To(Equal("200000000000000000"))
								evt = it.Event
								Expect(it.Next()).To(BeTrue())
								Expect(evt.To).To(Equal(RandomAccount.Address()))
								Expect(evt.Asset).To(Equal(ERC20Contract1Address))
								Expect(evt.Amount.String()).To(Equal("200"))
								evt = it.Event
								Expect(it.Next()).To(BeFalse())
								Expect(evt.To).To(Equal(RandomAccount.Address()))
								Expect(evt.Asset).To(Equal(ERC20Contract2Address))
								Expect(evt.Amount.String()).To(Equal("100"))
							})

							It("should increase the ERC20 type-1 balance of the random address by 200 (20%)", func() {
								b, err := ERC20Contract1.BalanceOf(nil, RandomAccount.Address())
								Expect(err).ToNot(HaveOccurred())
								Expect(b.String()).To(Equal("200"))
							})

							It("should increase the ERC20 type-2 balance of the random address by 200 (20%)", func() {
								b, err := ERC20Contract2.BalanceOf(nil, RandomAccount.Address())
								Expect(err).ToNot(HaveOccurred())
								Expect(b.String()).To(Equal("100"))
							})

							It("should NOT increase the ERC20 type-3 balance of the random address", func() {
								b, err := ERC20Contract3.BalanceOf(nil, RandomAccount.Address())
								Expect(err).ToNot(HaveOccurred())
								Expect(b.String()).To(Equal("0"))
							})

							It("should increase the ETH balance of the random address by 20% * 1 ETH", func() {
								b, e := Backend.BalanceAt(context.Background(), RandomAccount.Address(), nil)
								Expect(e).ToNot(HaveOccurred())
								bStr := b.String()
								newBalance := initialBalance.Add(initialBalance, big.NewInt(200000000000000000)) //we first add the 20% of 1 ETH
								newBalance.Sub(newBalance, txCost)                                               //we have to deduct the tx cost
								Expect(bStr).To(Equal(newBalance.String()))
							})

						})

						When("When the random address tries to burn more TKN than it owns", func() {

							var initialBalance *big.Int
							txCost := new(big.Int)

							BeforeEach(func() {
								var err error
								initialBalance, err = Backend.BalanceAt(context.Background(), RandomAccount.Address(), nil)
								Expect(err).ToNot(HaveOccurred())
							})

							BeforeEach(func() {
								tx, err := TKNBurner.Burn(RandomAccount.TransactOpts(), big.NewInt(0))
								Expect(err).ToNot(HaveOccurred())
								Backend.Commit()
								Expect(isSuccessful(tx)).To(BeTrue())
								r, err := Backend.TransactionReceipt(context.Background(), tx.Hash())
								Expect(err).ToNot(HaveOccurred())
								txCost.Mul(tx.GasPrice(), big.NewInt(int64(r.GasUsed)))
							})

							It("should NOT have an impact on the (ETH) balance of the holder contract", func() {
								b, e := Backend.BalanceAt(context.Background(), TokenHolderAddress, nil)
								Expect(e).ToNot(HaveOccurred())
								finalBal := EthToWei(1)
								Expect(b.String()).To(Equal(finalBal.String()))
							})

							It("should leave the total supply intact", func() {
								s, err := TKNBurner.TotalSupply(nil)
								Expect(err).ToNot(HaveOccurred())
								Expect(s.String()).To(Equal("1000"))
							})

							It("should leave the TKN balance of the random address intact", func() {
								b, err := TKNBurner.BalanceOf(nil, RandomAccount.Address())
								Expect(err).ToNot(HaveOccurred())
								Expect(b.String()).To(Equal("300"))
							})

							It("should leave the ERC20 type-1 balance of the holder contract intact", func() {
								b, err := ERC20Contract1.BalanceOf(nil, TokenHolderAddress)
								Expect(err).ToNot(HaveOccurred())
								Expect(b.String()).To(Equal("1000"))
							})

							It("should leave the ERC20 type-2 balance of the holder contract intact", func() {
								b, err := ERC20Contract2.BalanceOf(nil, TokenHolderAddress)
								Expect(err).ToNot(HaveOccurred())
								Expect(b.String()).To(Equal("500"))
							})

							It("Should emit a Transfer event of 0 value", func() {
								from := []common.Address{RandomAccount.Address()}
								to := []common.Address{common.HexToAddress("0x0")}
								it, err := TKNBurner.FilterTransfer(nil, from, to)
								Expect(err).ToNot(HaveOccurred())
								Expect(it.Next()).To(BeTrue())
								evt := it.Event
								Expect(it.Next()).To(BeFalse())
								Expect(evt.From).To(Equal(RandomAccount.Address()))
								Expect(evt.To).To(Equal(common.HexToAddress("0x0")))
								Expect(evt.Value.String()).To(Equal("0"))
							})

							It("should NOT increase the ERC20 type-1 balance of the random address", func() {
								b, err := ERC20Contract1.BalanceOf(nil, RandomAccount.Address())
								Expect(err).ToNot(HaveOccurred())
								Expect(b.String()).To(Equal("0"))
							})

							It("shouldn't increase the ERC20 type-2 balance of the random address", func() {
								b, err := ERC20Contract2.BalanceOf(nil, RandomAccount.Address())
								Expect(err).ToNot(HaveOccurred())
								Expect(b.String()).To(Equal("0"))
							})

							It("should NOT increase the ETH balance of the random address", func() {
								b, e := Backend.BalanceAt(context.Background(), RandomAccount.Address(), nil)
								Expect(e).ToNot(HaveOccurred())
								bStr := b.String()
								newBalance := initialBalance
								newBalance.Sub(initialBalance, txCost) //we have to deduct the tx cost
								Expect(bStr).To(Equal(newBalance.String()))
							})

						})

					})
				}) //When("The holder contract has two types of ERC20 tokens"
			})
		})
	})

})
