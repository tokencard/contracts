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

var _ = Describe("wallet load non-compliant ERC20", func() {

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

	When("The Wallet contract is credited with NonCompliantERC20 tokens ", func() {
		BeforeEach(func() {
			tx, err := NonCompliantERC20.Credit(BankAccount.TransactOpts(), WalletProxyAddress, EthToWei(1000))
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeTrue())
		})

		It("should increase the NonCompliantERC20 balance of the Wallet by 1000", func() {
			b, err := NonCompliantERC20.BalanceOf(nil, WalletProxyAddress)
			Expect(err).ToNot(HaveOccurred())
			Expect(b.String()).To(Equal(EthToWei(1000).String()))
		})

		When("the token is loadable and the rates have been updated", func() {

			BeforeEach(func() {
				tokens := []common.Address{NonCompliantERC20Address}
				tx, err := TokenWhitelist.AddTokens(
					ControllerAdmin.TransactOpts(),
					tokens,
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

			BeforeEach(func() {
				tx, err := TokenWhitelist.UpdateTokenRate(
					ControllerAdmin.TransactOpts(),
					NonCompliantERC20Address,
					big.NewInt(555),
					big.NewInt(20180913153211),
				)
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})

			When("a valid amount is transferred ", func() {

				BeforeEach(func() {
					tx, err := WalletProxy.LoadTokenCard(Owner.TransactOpts(), NonCompliantERC20Address, EthToWei(101))
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeTrue())
				})

				It("Should emit 1 NonCompliantERC20 Approval event", func() {
					owner := []common.Address{WalletProxyAddress}
					spender := []common.Address{LicenceAddress}
					it, err := NonCompliantERC20.FilterApproval(nil, owner, spender)
					Expect(err).ToNot(HaveOccurred())
					Expect(it.Next()).To(BeTrue())
					evt := it.Event
					Expect(it.Next()).To(BeFalse())
					Expect(evt.Owner).To(Equal(WalletProxyAddress))
					Expect(evt.Spender).To(Equal(LicenceAddress))
					Expect(evt.Value.String()).To(Equal(EthToWei(101).String()))
				})

				It("Should emit 1 TransferredToTokenHolder event", func() {
					it, err := Licence.FilterTransferredToTokenHolder(nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(it.Next()).To(BeTrue())
					evt := it.Event
					Expect(it.Next()).To(BeFalse())
					Expect(evt.From).To(Equal(WalletProxyAddress))
					Expect(evt.To).To(Equal(TokenHolderAddress))
					Expect(evt.Asset).To(Equal(NonCompliantERC20Address))
					Expect(evt.Amount.String()).To(Equal(EthToWei(1).String()))
				})

				It("Should emit 1 TransferredToCryptoFloat event", func() {
					it, err := Licence.FilterTransferredToCryptoFloat(nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(it.Next()).To(BeTrue())
					evt := it.Event
					Expect(it.Next()).To(BeFalse())
					Expect(evt.From).To(Equal(WalletProxyAddress))
					Expect(evt.To).To(Equal(CryptoFloatAddress))
					Expect(evt.Asset).To(Equal(NonCompliantERC20Address))
					Expect(evt.Amount.String()).To(Equal(EthToWei(100).String()))
				})

				It("Should emit 2 NonCompliantERC20 Transfer events", func() {
					from := []common.Address{WalletProxyAddress}
					var to []common.Address
					it, err := NonCompliantERC20.FilterTransfer(nil, from, to)
					Expect(err).ToNot(HaveOccurred())
					Expect(it.Next()).To(BeTrue())
					evt := it.Event
					Expect(it.Next()).To(BeTrue())
					Expect(evt.From).To(Equal(WalletProxyAddress))
					Expect(evt.To).To(Equal(TokenHolderAddress))
					Expect(evt.Amount.String()).To(Equal(EthToWei(1).String()))
					evt = it.Event
					Expect(it.Next()).To(BeFalse())
					Expect(evt.From).To(Equal(WalletProxyAddress))
					Expect(evt.To).To(Equal(CryptoFloatAddress))
					Expect(evt.Amount.String()).To(Equal(EthToWei(100).String()))
				})

				It("Should emit 1 LoadedTokenCard event", func() {
					it, err := WalletProxy.FilterLoadedTokenCard(nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(it.Next()).To(BeTrue())
					evt := it.Event
					Expect(it.Next()).To(BeFalse())
					Expect(evt.Asset).To(Equal(NonCompliantERC20Address))
					Expect(evt.Amount.String()).To(Equal(EthToWei(101).String()))
				})

				It("should increase the NonCompliantERC20 balance of the Holder contract by 1", func() {
					b, err := NonCompliantERC20.BalanceOf(nil, TokenHolderAddress)
					Expect(err).ToNot(HaveOccurred())
					Expect(b.String()).To(Equal(EthToWei(1).String()))
				})

				It("should increase the NonCompliantERC20 balance of the Float contract by 100", func() {
					b, err := NonCompliantERC20.BalanceOf(nil, CryptoFloatAddress)
					Expect(err).ToNot(HaveOccurred())
					Expect(b.String()).To(Equal(EthToWei(100).String()))
				})

				It("should decrease the NonCompliantERC201 balance of the Wallet by 101", func() {
					b, err := NonCompliantERC20.BalanceOf(nil, WalletProxyAddress)
					Expect(err).ToNot(HaveOccurred())
					Expect(b.String()).To(Equal(EthToWei(899).String()))
				})

			}) //equal to approval

			When("a bigger amount than the one owned is loaded", func() {

				It("Should revert", func() {
					tx, err := WalletProxy.LoadTokenCard(Owner.TransactOpts(ethertest.WithGasLimit(300000)), NonCompliantERC20Address, EthToWei(1001))
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeFalse())
					returnData, _ := ethCall(tx)
					Expect(string(returnData[len(returnData)-64:])).To(ContainSubstring("SafeERC20: low-level call failed"))
				})

			}) //more than owned (and hence can be approved)

			When("a bigger amount than daily limit is loaded", func() {

				BeforeEach(func() {
					tx, err := TokenWhitelist.UpdateTokenRate(
						ControllerAdmin.TransactOpts(),
						NonCompliantERC20Address,
						big.NewInt(int64(math.Pow10(18))),
						big.NewInt(20180913153211),
					)
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeTrue())
				})

				It("Should revert", func() {
					tx, err := WalletProxy.LoadTokenCard(Owner.TransactOpts(ethertest.WithGasLimit(200000)), NonCompliantERC20Address, EthToWei(1000))
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeFalse())
					returnData, _ := ethCall(tx)
					Expect(string(returnData[len(returnData)-64:])).To(ContainSubstring("available smaller than amount"))
				})

			}) //more than daily limit

		}) //tokens are loadable and rates have been updated

		When("the token is NOT loadable", func() {
			It("should fail (NOT loadable)", func() {
				tx, err := WalletProxy.LoadTokenCard(Owner.TransactOpts(ethertest.WithGasLimit(100000)), NonCompliantERC20Address, big.NewInt(101))
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
					[]common.Address{NonCompliantERC20Address},
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
				tx, err := WalletProxy.LoadTokenCard(Owner.TransactOpts(ethertest.WithGasLimit(100000)), NonCompliantERC20Address, big.NewInt(101))
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
