package wallet_test

import (
	"math/big"
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tokencard/contracts/test/shared"
	"github.com/tokencard/ethertest"
)

var _ = Describe("bulk_transfer", func() {

	When("the wallet has enough ETH", func() {
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
			When("the size of the array of assets does NOT match the size of the array of amounts", func() {
				It("should fail", func() {
					tx, err := Wallet.BulkTransfer(Owner.TransactOpts(ethertest.WithGasLimit(81000)), RandomAccount.Address(), []common.Address{common.HexToAddress("0x0"), common.HexToAddress("0x2"), common.HexToAddress("0x3")}, []*big.Int{big.NewInt(1), big.NewInt(2)})
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isGasExhausted(tx, 81000)).To(BeFalse())
					Expect(isSuccessful(tx)).To(BeFalse())
				})
			})

			When("the size of the array of assets DOES match the size of the array of amounts", func() {

				var randomBalance *big.Int
				var assets []common.Address
				var amounts []*big.Int

				BeforeEach(func() {
					var err error
					randomBalance, err = Backend.BalanceAt(context.Background(), RandomAccount.Address(), nil)
					Expect(err).ToNot(HaveOccurred())
					assets = []common.Address{common.HexToAddress("0x0"), ERC20Contract1Address, ERC20Contract2Address}
					amounts = []*big.Int{EthToWei(1), big.NewInt(1000), big.NewInt(500)}
				})

				BeforeEach(func() {
					var err error
					tx, err = Wallet.BulkTransfer(Owner.TransactOpts(), RandomAccount.Address(), assets, amounts)
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
				It("should increase the balance of the RandomAccount by 1 eth", func() {
					newBalance, err := Backend.BalanceAt(context.Background(), RandomAccount.Address(), nil)
					Expect(err).ToNot(HaveOccurred())
					randomBalance.Add(randomBalance,EthToWei(1))
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
				Expect(evt.Amounts).To(Equal(amounts))
			})

			})
		})//called by the owner

		When("called by a random address", func() {
			It("should fail", func() {
				tx, err := Wallet.BulkTransfer(RandomAccount.TransactOpts(ethertest.WithGasLimit(81000)), RandomAccount.Address(), []common.Address{common.HexToAddress("0x0"), common.HexToAddress("0x2"), common.HexToAddress("0x3")}, []*big.Int{big.NewInt(1), big.NewInt(2)})
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isGasExhausted(tx, 81000)).To(BeFalse())
				Expect(isSuccessful(tx)).To(BeFalse())
			})
		})
	})
})
