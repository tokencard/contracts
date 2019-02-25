package wallet_test

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tokencard/contracts/test/shared"
	"github.com/tokencard/ethertest"
)

var _ = Describe("wallet load ERC20", func() {

	It("the initial balance of the Float contract should be zero", func() {
		b, e := Backend.BalanceAt(context.Background(), CryptoFloatAddress, nil)
		Expect(e).ToNot(HaveOccurred())
		Expect(b.String()).To(Equal("0"))
	})

	It("the initial balance of Holder contract address should be zero", func() {
		b, e := Backend.BalanceAt(context.Background(), TokenHolderAddress, nil)
		Expect(e).ToNot(HaveOccurred())
		Expect(b.String()).To(Equal("0"))
	})

	It("the initial balance of the Wallet should be zero", func() {
		b, e := Backend.BalanceAt(context.Background(), WalletAddress, nil)
		Expect(e).ToNot(HaveOccurred())
		Expect(b.String()).To(Equal("0"))
	})

	When("The Wallet contract is credited with two types of ERC20 tokens and  102 ETH", func() {
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

		// BeforeEach(func() {
		// 	RandomAccount.MustTransfer(Backend, WalletAddress, EthToWei(102))
		// })

		It("should increase the ERC20 type-1 balance of the Wallet by 1000", func() {
			b, err := ERC20Contract1.BalanceOf(nil, WalletAddress)
			Expect(err).ToNot(HaveOccurred())
			Expect(b.String()).To(Equal("1000"))
		})

		It("should increase the ERC20 type-2 balance of the Wallet by 2", func() {
			b, err := ERC20Contract2.BalanceOf(nil, WalletAddress)
			Expect(err).ToNot(HaveOccurred())
			Expect(b.String()).To(Equal("500"))
		})

		When("the tokens are laodable and the rates have been updated", func() {

			BeforeEach(func() {
				tokens := []common.Address{ERC20Contract1Address, ERC20Contract2Address}
				tx, err := TokenWhitelist.AddTokens(
					Controller.TransactOpts(),
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
					big.NewInt(20180913153211),
				)
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})

			BeforeEach(func() {
				tx, err := TokenWhitelist.UpdateTokenRate(
					Controller.TransactOpts(),
					ERC20Contract1Address,
					big.NewInt(555),
					big.NewInt(20180913153211),
				)
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})

			BeforeEach(func() {
				tx, err := TokenWhitelist.UpdateTokenRate(
					Controller.TransactOpts(),
					ERC20Contract2Address,
					big.NewInt(555),
					big.NewInt(20180913153211),
				)
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})

			When("a valid amount is transfered ", func() {

				BeforeEach(func() {
					tx, err := Wallet.LoadTokenCard(Owner.TransactOpts(), ERC20Contract1Address, big.NewInt(101))
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeTrue())
				})

				BeforeEach(func() {
					tx, err := Wallet.LoadTokenCard(Owner.TransactOpts(), ERC20Contract2Address, big.NewInt(2))
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeTrue())
				})

				It("Should emit 1 ERC20Contract1 Approval event", func() {
					owner := []common.Address{WalletAddress}
					spender := []common.Address{LicenceAddress}
					it, err := ERC20Contract1.FilterApproval(nil, owner, spender)
					Expect(err).ToNot(HaveOccurred())
					Expect(it.Next()).To(BeTrue())
					evt := it.Event
					Expect(it.Next()).To(BeFalse())
					Expect(evt.Owner).To(Equal(WalletAddress))
					Expect(evt.Spender).To(Equal(LicenceAddress))
					Expect(evt.Value.String()).To(Equal("101"))
				})

				It("Should emit 1 ERC20Contract2 Approval event", func() {
					owner := []common.Address{WalletAddress}
					spender := []common.Address{LicenceAddress}
					it, err := ERC20Contract2.FilterApproval(nil, owner, spender)
					Expect(err).ToNot(HaveOccurred())
					Expect(it.Next()).To(BeTrue())
					evt := it.Event
					Expect(it.Next()).To(BeFalse())
					Expect(evt.Owner).To(Equal(WalletAddress))
					Expect(evt.Spender).To(Equal(LicenceAddress))
					Expect(evt.Value.String()).To(Equal("2"))
				})

				It("Should emit 2 TransferredToTokenHolder events", func() {
					it, err := Licence.FilterTransferredToTokenHolder(nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(it.Next()).To(BeTrue())
					evt := it.Event
					Expect(it.Next()).To(BeTrue())
					Expect(evt.From).To(Equal(WalletAddress))
					Expect(evt.To).To(Equal(TokenHolderAddress))
					Expect(evt.Asset).To(Equal(ERC20Contract1Address))
					Expect(evt.Amount.String()).To(Equal("1"))
					evt = it.Event
					Expect(it.Next()).To(BeFalse())
					Expect(evt.From).To(Equal(WalletAddress))
					Expect(evt.To).To(Equal(TokenHolderAddress))
					Expect(evt.Asset).To(Equal(ERC20Contract2Address))
					Expect(evt.Amount.String()).To(Equal("1"))
				})

				It("Should emit 2 TransferredToCryptoFloat events", func() {
					it, err := Licence.FilterTransferredToCryptoFloat(nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(it.Next()).To(BeTrue())
					evt := it.Event
					Expect(it.Next()).To(BeTrue())
					Expect(evt.From).To(Equal(WalletAddress))
					Expect(evt.To).To(Equal(CryptoFloatAddress))
					Expect(evt.Asset).To(Equal(ERC20Contract1Address))
					Expect(evt.Amount.String()).To(Equal("100"))
					evt = it.Event
					Expect(it.Next()).To(BeFalse())
					Expect(evt.From).To(Equal(WalletAddress))
					Expect(evt.To).To(Equal(CryptoFloatAddress))
					Expect(evt.Asset).To(Equal(ERC20Contract2Address))
					Expect(evt.Amount.String()).To(Equal("1"))
				})

				It("Should emit 2 ERC20Contract1 Transfer events", func() {
					from := []common.Address{WalletAddress}
					var to []common.Address
					it, err := ERC20Contract1.FilterTransfer(nil, from, to)
					Expect(err).ToNot(HaveOccurred())
					Expect(it.Next()).To(BeTrue())
					evt := it.Event
					Expect(it.Next()).To(BeTrue())
					Expect(evt.From).To(Equal(WalletAddress))
					Expect(evt.To).To(Equal(TokenHolderAddress))
					Expect(evt.Amount.String()).To(Equal("1"))
					evt = it.Event
					Expect(it.Next()).To(BeFalse())
					Expect(evt.From).To(Equal(WalletAddress))
					Expect(evt.To).To(Equal(CryptoFloatAddress))
					Expect(evt.Amount.String()).To(Equal("100"))
				})

				It("Should emit 2 ERC20Contract2 Transfer events", func() {
					from := []common.Address{WalletAddress}
					var to []common.Address
					it, err := ERC20Contract2.FilterTransfer(nil, from, to)
					Expect(err).ToNot(HaveOccurred())
					Expect(it.Next()).To(BeTrue())
					evt := it.Event
					Expect(it.Next()).To(BeTrue())
					Expect(evt.From).To(Equal(WalletAddress))
					Expect(evt.To).To(Equal(TokenHolderAddress))
					Expect(evt.Amount.String()).To(Equal("1"))
					evt = it.Event
					Expect(it.Next()).To(BeFalse())
					Expect(evt.From).To(Equal(WalletAddress))
					Expect(evt.To).To(Equal(CryptoFloatAddress))
					Expect(evt.Amount.String()).To(Equal("1"))
				})

				It("Should emit 2 LoadedTokenCard events", func() {
					it, err := Wallet.FilterLoadedTokenCard(nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(it.Next()).To(BeTrue())
					evt := it.Event
					Expect(it.Next()).To(BeTrue())
					Expect(evt.Asset).To(Equal(ERC20Contract1Address))
					Expect(evt.Amount.String()).To(Equal("101"))
					evt = it.Event
					Expect(it.Next()).To(BeFalse())
					Expect(evt.Asset).To(Equal(ERC20Contract2Address))
					Expect(evt.Amount.String()).To(Equal("2"))
				})

				It("should increase the ERC20 type-1 balance of the Holder contract by 1", func() {
					b, err := ERC20Contract1.BalanceOf(nil, TokenHolderAddress)
					Expect(err).ToNot(HaveOccurred())
					Expect(b.String()).To(Equal("1"))
				})

				It("should increase the ERC20 type-2 balance of the Holder contract by 1", func() {
					b, err := ERC20Contract2.BalanceOf(nil, TokenHolderAddress)
					Expect(err).ToNot(HaveOccurred())
					Expect(b.String()).To(Equal("1"))
				})

				It("should increase the ERC20 type-1 balance of the Float contract by 100", func() {
					b, err := ERC20Contract1.BalanceOf(nil, CryptoFloatAddress)
					Expect(err).ToNot(HaveOccurred())
					Expect(b.String()).To(Equal("100"))
				})

				It("should increase the ERC20 type-2 balance of the Float contract by 1", func() {
					b, err := ERC20Contract2.BalanceOf(nil, CryptoFloatAddress)
					Expect(err).ToNot(HaveOccurred())
					Expect(b.String()).To(Equal("1"))
				})

				It("should decrease the ERC20 type-1 balance of the Wallet by 101", func() {
					b, err := ERC20Contract1.BalanceOf(nil, WalletAddress)
					Expect(err).ToNot(HaveOccurred())
					Expect(b.String()).To(Equal("899"))
				})

				It("should decrease the ERC20 type-2 balance of the Wallet by 2", func() {
					b, err := ERC20Contract2.BalanceOf(nil, WalletAddress)
					Expect(err).ToNot(HaveOccurred())
					Expect(b.String()).To(Equal("498"))
				})

			}) //equal to approval

			When("a bigger amount than the one owned is tried to be transfered ", func() {

				It("Should revert", func() {
					tx, err := Wallet.LoadTokenCard(Owner.TransactOpts(ethertest.WithGasLimit(300000)), ERC20Contract1Address, big.NewInt(1001))
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isGasExhausted(tx, 300000)).To(BeFalse())
					Expect(isSuccessful(tx)).To(BeFalse())
				})

				It("Should revert", func() {
					tx, err := Wallet.LoadTokenCard(Owner.TransactOpts(ethertest.WithGasLimit(300000)), ERC20Contract2Address, big.NewInt(501))
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isGasExhausted(tx, 300000)).To(BeFalse())
					Expect(isSuccessful(tx)).To(BeFalse())
				})

			}) //more than owned (and hence can be approved)

			When("a bigger amount than daily Load limit is tried to be loaded", func() {

				//change the exchange rate to be equal to daily Load limit + 1 wei
				BeforeEach(func() {

					limPlusOneWei, err := Wallet.LoadLimit(nil)
					Expect(err).ToNot(HaveOccurred())

					limPlusOneWei.Add(limPlusOneWei, big.NewInt(1))

					tx, err := TokenWhitelist.UpdateTokenRate(
						Controller.TransactOpts(),
						ERC20Contract1Address,
						limPlusOneWei,
						big.NewInt(20180913153211),
					)
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeTrue())
				})

				It("Should revert", func() {
					//1 token  = 101 eth + 1 wei => it should exceed the daily limit, 1 token =  1*magnitude (10^18)
					tx, err := Wallet.LoadTokenCard(Owner.TransactOpts(ethertest.WithGasLimit(100000)), ERC20Contract1Address, EthToWei(1))
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isGasExhausted(tx, 100000)).To(BeFalse())
					Expect(isSuccessful(tx)).To(BeFalse())
					Expect(TestRig.LastExecuted()).To(MatchRegexp(`.*_loadLimit.useAmount\(etherValue\);`))
				})

			}) //more than daily Load limit

		}) //tokens are loadable and rates have been updated

		When("the token is NOT loadable", func() {
			It("should fail (NOT loadable)", func() {
				tx, err := Wallet.LoadTokenCard(Owner.TransactOpts(ethertest.WithGasLimit(100000)), ERC20Contract1Address, big.NewInt(101))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isGasExhausted(tx, 100000)).To(BeFalse())
				Expect(isSuccessful(tx)).To(BeFalse())
				Expect(TestRig.LastExecuted()).To(MatchRegexp(`.*require\(_isTokenLoadable\(_asset\), "token not loadable"\);`))
			})
		})

		When("the token is loadable but the rate has NOT been updated", func() {
			BeforeEach(func() {
				tx, err := TokenWhitelist.AddTokens(
					Controller.TransactOpts(),
					[]common.Address{ERC20Contract1Address},
					StringsToByte32(
						"ERC1",
					),
					[]*big.Int{
						DecimalsToMagnitude(big.NewInt(18)),
					},
					[]bool{true},
					big.NewInt(20180913153211),
				)
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})

			It("should fail (rate == 0)", func() {
				tx, err := Wallet.LoadTokenCard(Owner.TransactOpts(ethertest.WithGasLimit(100000)), ERC20Contract1Address, big.NewInt(101))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isGasExhausted(tx, 100000)).To(BeFalse())
				Expect(isSuccessful(tx)).To(BeFalse())
				Expect(TestRig.LastExecuted()).To(MatchRegexp(`.*require\(rate != 0, "token rate is 0"\);`))
			})

		}) //the token is loadable but the rate has NOT been updated

	}) //credited with two types of ERC20 tokens

})
