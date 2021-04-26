package wallet_test

import (
	"context"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tokencard/contracts/v3/test/shared"
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

	When("The Wallet contract is credited with two types of ERC20 tokens", func() {
		BeforeEach(func() {
			tx, err := ERC20Contract1.Credit(BankAccount.TransactOpts(), WalletProxyAddress, big.NewInt(1000))
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeTrue())
		})

		BeforeEach(func() {
			tx, err := ERC20Contract2.Credit(BankAccount.TransactOpts(), WalletProxyAddress, big.NewInt(500))
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeTrue())
		})

		It("should increase the ERC20 type-1 balance of the Wallet by 1000", func() {
			b, err := ERC20Contract1.BalanceOf(nil, WalletProxyAddress)
			Expect(err).ToNot(HaveOccurred())
			Expect(b.String()).To(Equal("1000"))
		})

		It("should increase the ERC20 type-2 balance of the Wallet by 500", func() {
			b, err := ERC20Contract2.BalanceOf(nil, WalletProxyAddress)
			Expect(err).ToNot(HaveOccurred())
			Expect(b.String()).To(Equal("500"))
		})

		When("the tokens are loadable and the rates have been updated", func() {

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

			BeforeEach(func() {
				tx, err := TokenWhitelist.UpdateTokenRate(
					ControllerAdmin.TransactOpts(),
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
					ControllerAdmin.TransactOpts(),
					ERC20Contract2Address,
					big.NewInt(555),
					big.NewInt(20180913153211),
				)
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})

			When("a valid amount is transferred ", func() {

				BeforeEach(func() {
					tx, err := WalletProxy.LoadTokenCard(Owner.TransactOpts(), ERC20Contract1Address, big.NewInt(101))
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeTrue())
				})

				BeforeEach(func() {
					tx, err := WalletProxy.LoadTokenCard(Owner.TransactOpts(), ERC20Contract2Address, big.NewInt(2))
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeTrue())
				})

				It("Should emit one ERC20Contract1 Approval event", func() {
					owner := []common.Address{WalletProxyAddress}
					spender := []common.Address{LicenceAddress}
					it, err := ERC20Contract1.FilterApproval(nil, owner, spender)
					Expect(err).ToNot(HaveOccurred())
					Expect(it.Next()).To(BeTrue())
					evt := it.Event
					Expect(it.Next()).To(BeFalse())
					Expect(evt.Owner).To(Equal(WalletProxyAddress))
					Expect(evt.Spender).To(Equal(LicenceAddress))
					Expect(evt.Value.String()).To(Equal("101"))
				})

				It("Should emit one ERC20Contract2 Approval event", func() {
					owner := []common.Address{WalletProxyAddress}
					spender := []common.Address{LicenceAddress}
					it, err := ERC20Contract2.FilterApproval(nil, owner, spender)
					Expect(err).ToNot(HaveOccurred())
					Expect(it.Next()).To(BeTrue())
					evt := it.Event
					Expect(it.Next()).To(BeFalse())
					Expect(evt.Owner).To(Equal(WalletProxyAddress))
					Expect(evt.Spender).To(Equal(LicenceAddress))
					Expect(evt.Value.String()).To(Equal("2"))
				})

				It("Should emit two TransferredToTokenHolder events", func() {
					it, err := Licence.FilterTransferredToTokenHolder(nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(it.Next()).To(BeTrue())
					evt := it.Event
					Expect(it.Next()).To(BeTrue())
					Expect(evt.From).To(Equal(WalletProxyAddress))
					Expect(evt.To).To(Equal(TokenHolderAddress))
					Expect(evt.Asset).To(Equal(ERC20Contract1Address))
					Expect(evt.Amount.String()).To(Equal("1"))
					evt = it.Event
					Expect(it.Next()).To(BeFalse())
					Expect(evt.From).To(Equal(WalletProxyAddress))
					Expect(evt.To).To(Equal(TokenHolderAddress))
					Expect(evt.Asset).To(Equal(ERC20Contract2Address))
					Expect(evt.Amount.String()).To(Equal("1"))
				})

				It("Should emit two TransferredToCryptoFloat events", func() {
					it, err := Licence.FilterTransferredToCryptoFloat(nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(it.Next()).To(BeTrue())
					evt := it.Event
					Expect(it.Next()).To(BeTrue())
					Expect(evt.From).To(Equal(WalletProxyAddress))
					Expect(evt.To).To(Equal(CryptoFloatAddress))
					Expect(evt.Asset).To(Equal(ERC20Contract1Address))
					Expect(evt.Amount.String()).To(Equal("100"))
					evt = it.Event
					Expect(it.Next()).To(BeFalse())
					Expect(evt.From).To(Equal(WalletProxyAddress))
					Expect(evt.To).To(Equal(CryptoFloatAddress))
					Expect(evt.Asset).To(Equal(ERC20Contract2Address))
					Expect(evt.Amount.String()).To(Equal("1"))
				})

				It("Should emit two ERC20Contract1 Transfer events", func() {
					from := []common.Address{WalletProxyAddress}
					var to []common.Address
					it, err := ERC20Contract1.FilterTransfer(nil, from, to)
					Expect(err).ToNot(HaveOccurred())
					Expect(it.Next()).To(BeTrue())
					evt := it.Event
					Expect(it.Next()).To(BeTrue())
					Expect(evt.From).To(Equal(WalletProxyAddress))
					Expect(evt.To).To(Equal(TokenHolderAddress))
					Expect(evt.Amount.String()).To(Equal("1"))
					evt = it.Event
					Expect(it.Next()).To(BeFalse())
					Expect(evt.From).To(Equal(WalletProxyAddress))
					Expect(evt.To).To(Equal(CryptoFloatAddress))
					Expect(evt.Amount.String()).To(Equal("100"))
				})

				It("Should emit two ERC20Contract2 Transfer events", func() {
					from := []common.Address{WalletProxyAddress}
					var to []common.Address
					it, err := ERC20Contract2.FilterTransfer(nil, from, to)
					Expect(err).ToNot(HaveOccurred())
					Expect(it.Next()).To(BeTrue())
					evt := it.Event
					Expect(it.Next()).To(BeTrue())
					Expect(evt.From).To(Equal(WalletProxyAddress))
					Expect(evt.To).To(Equal(TokenHolderAddress))
					Expect(evt.Amount.String()).To(Equal("1"))
					evt = it.Event
					Expect(it.Next()).To(BeFalse())
					Expect(evt.From).To(Equal(WalletProxyAddress))
					Expect(evt.To).To(Equal(CryptoFloatAddress))
					Expect(evt.Amount.String()).To(Equal("1"))
				})

				It("Should emit two LoadedTokenCard events", func() {
					it, err := WalletProxy.FilterLoadedTokenCard(nil)
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
					b, err := ERC20Contract1.BalanceOf(nil, WalletProxyAddress)
					Expect(err).ToNot(HaveOccurred())
					Expect(b.String()).To(Equal("899"))
				})

				It("should decrease the ERC20 type-2 balance of the Wallet by 2", func() {
					b, err := ERC20Contract2.BalanceOf(nil, WalletProxyAddress)
					Expect(err).ToNot(HaveOccurred())
					Expect(b.String()).To(Equal("498"))
				})

			}) //equal to approval

			When("a bigger amount than the one owned is tried to be transferred ", func() {

				It("Should revert", func() {
					tx, err := WalletProxy.LoadTokenCard(Owner.TransactOpts(ethertest.WithGasLimit(300000)), ERC20Contract1Address, big.NewInt(1001))
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isGasExhausted(tx, 300000)).To(BeFalse())
					Expect(isSuccessful(tx)).To(BeFalse())
				})

				It("Should revert", func() {
					tx, err := WalletProxy.LoadTokenCard(Owner.TransactOpts(ethertest.WithGasLimit(300000)), ERC20Contract2Address, big.NewInt(501))
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isGasExhausted(tx, 300000)).To(BeFalse())
					Expect(isSuccessful(tx)).To(BeFalse())
				})

			}) //more than owned (and hence can be approved)

			When("a bigger amount than daily load limit is loaded", func() {

				BeforeEach(func() {
					tx, err := TokenWhitelist.UpdateTokenRate(
						ControllerAdmin.TransactOpts(),
						ERC20Contract1Address,
						big.NewInt(int64(math.Pow10(18))),
						big.NewInt(20180913153211),
					)
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeTrue())
				})

				It("Should revert", func() {
					tx, err := WalletProxy.LoadTokenCard(Owner.TransactOpts(ethertest.WithGasLimit(200000)), ERC20Contract1Address, EthToWei(1000))
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeFalse())
					returnData, _ := ethCall(tx)
					Expect(string(returnData[len(returnData)-64:])).To(ContainSubstring("available smaller than amount"))
				})

			}) //more than daily Load limit

		}) //tokens are loadable and rates have been updated

		When("the token is NOT loadable", func() {
			It("should fail (NOT loadable)", func() {
				tx, err := WalletProxy.LoadTokenCard(Owner.TransactOpts(ethertest.WithGasLimit(100000)), ERC20Contract1Address, big.NewInt(101))
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
					ControllerAdmin.TransactOpts(),
					[]common.Address{ERC20Contract1Address},
					StringsToByte32(
						"ERC1",
					),
					[]*big.Int{
						DecimalsToMagnitude(big.NewInt(18)),
					},
					[]bool{true},
					[]bool{true},
					big.NewInt(20180913153211),
				)
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})

			It("should fail (rate == 0)", func() {
				tx, err := WalletProxy.LoadTokenCard(Owner.TransactOpts(ethertest.WithGasLimit(100000)), ERC20Contract1Address, big.NewInt(101))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isGasExhausted(tx, 100000)).To(BeFalse())
				Expect(isSuccessful(tx)).To(BeFalse())
				returnData, _ := ethCall(tx)
				Expect(string(returnData[len(returnData)-64:])).To(ContainSubstring("rate=0"))
			})

		}) //the token is loadable but the rate has NOT been updated

	}) //credited with two types of ERC20 tokens

})
