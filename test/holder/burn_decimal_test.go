package tokenHolder_test
//
// import (
// 	// "errors"
// 	"math/big"
// 	"context"
//
//   "github.com/ethereum/go-ethereum/core/types"
//
// 	. "github.com/onsi/ginkgo"
// 	. "github.com/onsi/gomega"
// 	. "github.com/tokencard/contracts/test/shared"
// 	// "github.com/tokencard/ethertest"
// 	"github.com/ethereum/go-ethereum/common"
//
// )
//
// var _ = Describe("TokenHolder", func() {
//
// 	Context("A valid Hodler wants to burrrrrnnnn TKN", func() {
//
//     var tx *types.Transaction
//
//     When("one thousand TKN are credited to an external address", func() {
//
//
// 			BeforeEach(func() {
// 				var err error
// 				tx, err = TKNBurner.Credit(BankAccount.TransactOpts(), BankAccount.Address(), big.NewInt(583))
// 				Expect(err).ToNot(HaveOccurred())
// 				Backend.Commit()
// 				Expect(isSuccessful(tx)).To(BeTrue())
// 			})
//
//       It("should increase the total TKN supply to 583", func() {
//         s, err := TKNBurner.TotalSupply(nil)
//         Expect(err).ToNot(HaveOccurred())
//         Expect(s.String()).To(Equal("583"))
//       })
//
//       It("should increase TKN balance of the specific address", func() {
//         b, err := TKNBurner.BalanceOf(nil, BankAccount.Address())
//         Expect(err).ToNot(HaveOccurred())
//         Expect(b.String()).To(Equal("583"))
//       })
//
//       When("187 TKN are transfered to a random address", func() {
// 				BeforeEach(func() {
// 					tx, err := TKNBurner.Transfer(BankAccount.TransactOpts(), RandomAccount.Address(), big.NewInt(187))
// 					Expect(err).ToNot(HaveOccurred())
// 					Backend.Commit()
// 					Expect(isSuccessful(tx)).To(BeTrue())
// 				})
//
// 				It("should increase TKN balance of the random person", func() {
// 					b, err := TKNBurner.BalanceOf(nil, RandomAccount.Address())
// 					Expect(err).ToNot(HaveOccurred())
// 					Expect(b.String()).To(Equal("187"))
// 				})
//
// 				It("should decrease TKN balance of the initial address", func() {
// 					b, err := TKNBurner.BalanceOf(nil, BankAccount.Address())
// 					Expect(err).ToNot(HaveOccurred())
// 					Expect(b.String()).To(Equal("396"))
// 				})
//
//         It("should leave the total supply intact", func() {
//           s, err := TKNBurner.TotalSupply(nil)
//           Expect(err).ToNot(HaveOccurred())
//           Expect(s.String()).To(Equal("583"))
// 				})
//
// 				When("The holder contract owns two types of ERC20 tokens and 1 ETH", func() {
//
// 					//add the tokens to the list
// 					BeforeEach(func() {
// 						var err error
// 						tokens := []common.Address{common.HexToAddress("0x0"), ERC20Contract1Address, ERC20Contract2Address}
// 						tx, err = TokenHolder.SetTokens(Owner.TransactOpts(), tokens)
// 						Expect(err).ToNot(HaveOccurred())
// 						Backend.Commit()
// 						Expect(isSuccessful(tx)).To(BeTrue())
// 					})
//
// 					BeforeEach(func() {
// 						var err error
// 						tx, err = ERC20Contract1.Credit(BankAccount.TransactOpts(), TokenHolderAddress, big.NewInt(396))
// 						Expect(err).ToNot(HaveOccurred())
// 						Backend.Commit()
// 						Expect(isSuccessful(tx)).To(BeTrue())
// 					})
//
// 					BeforeEach(func() {
// 						var err error
// 						tx, err = ERC20Contract2.Credit(BankAccount.TransactOpts(), TokenHolderAddress, big.NewInt(765))
// 						Expect(err).ToNot(HaveOccurred())
// 						Backend.Commit()
// 						Expect(isSuccessful(tx)).To(BeTrue())
//
// 					})
//
// 					BeforeEach(func() {
// 						BankAccount.MustTransfer(Backend, TokenHolderAddress, EthToWei(1))
// 					})
//
// 					It("should increase the ETH balance of the holder contract address by 1 ETH", func() {
// 						b, e := Backend.BalanceAt(context.Background(), TokenHolderAddress, nil)
// 						Expect(e).ToNot(HaveOccurred())
// 						bStr := b.String()
// 						Expect(bStr).To(Equal("1000000000000000000"))
// 						Expect(len(bStr)).To(Equal(19)) //1 ETH = 10^18 wei
// 					})
//
// 					It("should increase the ERC20 type-1 balance of the holder contract by 396", func() {
// 						b, err := ERC20Contract1.BalanceOf(nil, TokenHolderAddress)
// 						Expect(err).ToNot(HaveOccurred())
// 						Expect(b.String()).To(Equal("396"))
// 					})
//
// 					It("should increase the ERC20 type-2 balance of the holder contract by 765", func() {
// 						b, err := ERC20Contract2.BalanceOf(nil, TokenHolderAddress)
// 						Expect(err).ToNot(HaveOccurred())
// 						Expect(b.String()).To(Equal("765"))
// 					})
//
// 					When("When the token holder (contract) is set by the owner", func() {
//
// 						BeforeEach(func() {
// 							tx, err := TKNBurner.SetTokenHolder(Owner.TransactOpts(), TokenHolderAddress)
// 							Expect(err).ToNot(HaveOccurred())
// 							Backend.Commit()
// 							Expect(isSuccessful(tx)).To(BeTrue())
// 						})
//
//
// 						When("When the random address burns 200 tokens (out of 1000)", func() {
//
// 							var initialBalance *big.Int
// 							txCost := new(big.Int)
//
// 							BeforeEach(func() {
// 								var err error
// 								initialBalance, err = Backend.BalanceAt(context.Background(), RandomAccount.Address(), nil)
// 								Expect(err).ToNot(HaveOccurred())
// 						 })
//
// 							BeforeEach(func() {
// 								tx, err := TKNBurner.Burn(RandomAccount.TransactOpts(), big.NewInt(200))
// 								Expect(err).ToNot(HaveOccurred())
// 								Backend.Commit()
// 								Expect(isSuccessful(tx)).To(BeTrue())
// 								r, err := Backend.TransactionReceipt(context.Background(), tx.Hash())
// 								Expect(err).ToNot(HaveOccurred())
// 								Expect(err).ToNot(HaveOccurred())
// 								txCost.Mul(tx.GasPrice(), big.NewInt(int64(r.GasUsed)))
// 							})
//
// 							FIt("should decrease the (ETH) balance of the holder contract by 3.11764705882%", func() {
// 								b, e := Backend.BalanceAt(context.Background(), TokenHolderAddress, nil)
// 								Expect(e).ToNot(HaveOccurred())
// 								finalBal := EthToWei(8)
// 								finalBal.Div(finalBal, big.NewInt(10))//the final balance should be 0.8 ETH -> 8/10
// 								Expect(b.String()).To(Equal(finalBal.String()))
// 							})
//
// 							It("should decrease the total supply by 200 (800 remaining)", func() {
// 			          s, err := TKNBurner.TotalSupply(nil)
// 			          Expect(err).ToNot(HaveOccurred())
// 	 	 	          Expect(s.String()).To(Equal("800"))
// 							})
//
// 							It("should decrease TKN balance of the random address", func() {
// 								b, err := TKNBurner.BalanceOf(nil, RandomAccount.Address())
// 								Expect(err).ToNot(HaveOccurred())
// 								Expect(b.String()).To(Equal("100"))
// 							})
//
// 							It("Should emit a Burn event", func() {
// 								it, err := TKNBurner.FilterBurn(nil)
// 								Expect(err).ToNot(HaveOccurred())
// 								Expect(it.Next()).To(BeTrue())
// 								evt := it.Event
// 								Expect(it.Next()).To(BeFalse())
// 								Expect(evt.Burner).To(Equal(RandomAccount.Address()))
// 								Expect(evt.Amount.String()).To(Equal("200"))
// 							})
//
// 							It("should increase the ERC20 type-1 balance of the random address by 200 (20%)", func() {
// 								b, err := ERC20Contract1.BalanceOf(nil, RandomAccount.Address())
// 								Expect(err).ToNot(HaveOccurred())
// 								Expect(b.String()).To(Equal("200"))
// 							})
//
// 							It("should increase the ERC20 type-2 balance of the random address by 200 (20%)", func() {
// 								b, err := ERC20Contract2.BalanceOf(nil, RandomAccount.Address())
// 								Expect(err).ToNot(HaveOccurred())
// 								Expect(b.String()).To(Equal("100"))
// 							})
//
// 							It("should increase the ETH balance of the random address by 20% * 1 ETH", func() {
// 								b, e := Backend.BalanceAt(context.Background(), RandomAccount.Address(), nil)
// 								Expect(e).ToNot(HaveOccurred())
// 								bStr := b.String()
// 								newBalance := initialBalance.Add(initialBalance, big.NewInt(200000000000000000)) //we first add the 20% of 1 ETH
// 								newBalance.Sub(newBalance, txCost) //we have to deduct the tx cost
// 								Expect(bStr).To(Equal(newBalance.String()))
// 							})
//
// 						})
// 					})
// 				}) //When("The holder contract has two types of ERC20 tokens"
//     })
//
//     // BeforeEach(func() {
//     //   var err error
//     //   tx, err = TKNBurner.Burn(RandomAccount.TransactOpts(), big.NewInt(500))
//     //   Expect(err).ToNot(HaveOccurred())
//     //   Backend.Commit()
//     // })
//     //
//     // It("Should succeed", func() {
//     //   Expect(isSuccessful(tx)).To(BeTrue())
//     // })
//
// 	 })
//   })
//
//
// })
