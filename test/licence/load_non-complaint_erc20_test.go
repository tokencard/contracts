package licence_test

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tokencard/contracts/test/shared"
	"github.com/tokencard/ethertest"
)

var _ = Describe("load ERC20", func() {

	It("the initial balance of the Float contract should be zero", func() {
		b, e := Backend.BalanceAt(context.Background(), CryptoFloatAddress, nil)
		Expect(e).ToNot(HaveOccurred())
		Expect(b.String()).To(Equal("0"))
	})

	It("the initial balance of Holder address should be zero", func() {
		b, e := Backend.BalanceAt(context.Background(), TokenHolderAddress, nil)
		Expect(e).ToNot(HaveOccurred())
		Expect(b.String()).To(Equal("0"))
	})

	When("The RandomAccount is credited with two types of ERC20 tokens", func() {
		BeforeEach(func() {
			tx, err := NonCompliantERC20.Credit(BankAccount.TransactOpts(), RandomAccount.Address(), big.NewInt(1000))
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeTrue())
		})

		When("no approval is given", func() {
			It("Should revert", func() {
				tx, err := Licence.Load(RandomAccount.TransactOpts(ethertest.WithGasLimit(100000)), NonCompliantERC20Address, big.NewInt(100))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isGasExhausted(tx, 100000)).To(BeFalse())
				Expect(isSuccessful(tx)).To(BeFalse())
			})
		})

		When("approval to Licence contract to transfer 101 tokens is given", func() {

			BeforeEach(func() {
				tx, err := NonCompliantERC20.Approve(RandomAccount.TransactOpts(), LicenceAddress, big.NewInt(101))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})

			It("Should emit 1 NonCompliantERC20 Approval event", func() {
				owner := []common.Address{RandomAccount.Address()}
				spender := []common.Address{LicenceAddress}
				it, err := NonCompliantERC20.FilterApproval(nil, owner, spender)
				Expect(err).ToNot(HaveOccurred())
				Expect(it.Next()).To(BeTrue())
				evt := it.Event
				Expect(it.Next()).To(BeFalse())
				Expect(evt.Owner).To(Equal(RandomAccount.Address()))
				Expect(evt.Spender).To(Equal(LicenceAddress))
				Expect(evt.Value.String()).To(Equal("101"))
			})

			When("the exact approved amount is transfered ", func() {

				BeforeEach(func() {
					tx, err := Licence.Load(RandomAccount.TransactOpts(), NonCompliantERC20Address, big.NewInt(101))
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeTrue())
				})

				It("Should emit 1 TransferredToTokenHolder event", func() {
					it, err := Licence.FilterTransferredToTokenHolder(nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(it.Next()).To(BeTrue())
					evt := it.Event
					Expect(it.Next()).To(BeFalse())
					Expect(evt.From).To(Equal(RandomAccount.Address()))
					Expect(evt.To).To(Equal(TokenHolderAddress))
					Expect(evt.Asset).To(Equal(NonCompliantERC20Address))
					Expect(evt.Amount.String()).To(Equal("1"))
				})

				It("Should emit 1 TransferredToCryptoFloat event", func() {
					it, err := Licence.FilterTransferredToCryptoFloat(nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(it.Next()).To(BeTrue())
					evt := it.Event
					Expect(it.Next()).To(BeFalse())
					Expect(evt.From).To(Equal(RandomAccount.Address()))
					Expect(evt.To).To(Equal(CryptoFloatAddress))
					Expect(evt.Asset).To(Equal(NonCompliantERC20Address))
					Expect(evt.Amount.String()).To(Equal("100"))
				})

				It("Should emit 2 NonCompliantERC20 Transfer events", func() {
					from := []common.Address{RandomAccount.Address()}
					var to []common.Address
					it, err := NonCompliantERC20.FilterTransfer(nil, from, to)
					Expect(err).ToNot(HaveOccurred())
					Expect(it.Next()).To(BeTrue())
					evt := it.Event
					Expect(it.Next()).To(BeTrue())
					Expect(evt.From).To(Equal(RandomAccount.Address()))
					Expect(evt.To).To(Equal(TokenHolderAddress))
					Expect(evt.Amount.String()).To(Equal("1"))
					evt = it.Event
					Expect(it.Next()).To(BeFalse())
					Expect(evt.From).To(Equal(RandomAccount.Address()))
					Expect(evt.To).To(Equal(CryptoFloatAddress))
					Expect(evt.Amount.String()).To(Equal("100"))
				})

				It("should increase the NonCompliantERC20 balance of the Holder contract by 1", func() {
					b, err := NonCompliantERC20.BalanceOf(nil, TokenHolderAddress)
					Expect(err).ToNot(HaveOccurred())
					Expect(b.String()).To(Equal("1"))
				})

				It("should increase the NonCompliantERC20 balance of the Float contract by 100", func() {
					b, err := NonCompliantERC20.BalanceOf(nil, CryptoFloatAddress)
					Expect(err).ToNot(HaveOccurred())
					Expect(b.String()).To(Equal("100"))
				})

				It("should decrease the NonCompliantERC20 balance of the RandomAccount 101", func() {
					b, err := NonCompliantERC20.BalanceOf(nil, RandomAccount.Address())
					Expect(err).ToNot(HaveOccurred())
					Expect(b.String()).To(Equal("899"))
				})

			}) //equal to approval

			When("a bigger amount than the approved one is tried to be transfered ", func() {

				It("Should revert", func() {
					tx, err := Licence.Load(RandomAccount.TransactOpts(ethertest.WithGasLimit(100000)), NonCompliantERC20Address, big.NewInt(102))
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isGasExhausted(tx, 100000)).To(BeFalse())
					Expect(isSuccessful(tx)).To(BeFalse())
				})

			}) //more than approved

		}) //approval is given

	}) //credited with two types of ERC20 tokens

})
