package token_whitelist_test

import (
	"math/big"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tokencard/contracts/v3/test/shared"
	"github.com/tokencard/ethertest"
)

var _ = Describe("TokenWhitelist claim", func() {

	It("the initial ERC20 type-1 balance of the TokenWhitelist contract should be zero", func() {
		b, err := ERC20Contract1.BalanceOf(nil, TokenWhitelistAddress)
		Expect(err).ToNot(HaveOccurred())
		Expect(b.String()).To(Equal("0"))
	})

	It("the initial ERC20 type-1 balance of the RandomAccount  should be zero", func() {
		b, err := ERC20Contract1.BalanceOf(nil, RandomAccount.Address())
		Expect(err).ToNot(HaveOccurred())
		Expect(b.String()).To(Equal("0"))
	})

	When("The TokenWhitelist contract is credited with 333 of ERC20 type-1 tokens", func() {
		BeforeEach(func() {
			tx, err := ERC20Contract1.Credit(BankAccount.TransactOpts(), TokenWhitelistAddress, big.NewInt(333))
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeTrue())
		})

		It("should increase the ERC20 type-1 balance of the TokenWhitelist contract by 333", func() {
			b, err := ERC20Contract1.BalanceOf(nil, TokenWhitelistAddress)
			Expect(err).ToNot(HaveOccurred())
			Expect(b.String()).To(Equal("333"))
		})

		When("the TokenWhitelistAddress owner withdraws 222 tokens", func() {

			BeforeEach(func() {
				tx, err := TokenWhitelist.Claim(ControllerAdmin.TransactOpts(), RandomAccount.Address(), ERC20Contract1Address, big.NewInt(222))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})

			It("Should emit 1 Claimed event", func() {
				it, err := TokenWhitelist.FilterClaimed(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(it.Next()).To(BeTrue())
				evt := it.Event
				Expect(it.Next()).To(BeFalse())
				Expect(evt.To).To(Equal(RandomAccount.Address()))
				Expect(evt.Asset).To(Equal(ERC20Contract1Address))
				Expect(evt.Amount.String()).To(Equal("222"))
			})

			It("should decrease the ERC20 type-1 balance of the TokenWhitelist contract contract by 222", func() {
				b, err := ERC20Contract1.BalanceOf(nil, TokenWhitelistAddress)
				Expect(err).ToNot(HaveOccurred())
				Expect(b.String()).To(Equal("111"))
			})

			It("should increase the ERC20 type-1 balance of the RandomAccount by 222", func() {
				b, err := ERC20Contract1.BalanceOf(nil, RandomAccount.Address())
				Expect(err).ToNot(HaveOccurred())
				Expect(b.String()).To(Equal("222"))
			})

		}) //owner withdrawal

		When("a random address tries to withdraw", func() {
			It("Should fail", func() {
				tx, err := TokenWhitelist.Claim(RandomAccount.TransactOpts(ethertest.WithGasLimit(100000)), RandomAccount.Address(), ERC20Contract1Address, big.NewInt(222))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeFalse())
			})
		})

	})

})
