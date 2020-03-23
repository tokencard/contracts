package controller_test

import (
	"context"
	"math/big"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tokencard/contracts/v3/test/shared"
	"github.com/tokencard/ethertest"
)

var _ = Describe("Controller claim", func() {

	It("the initial balance of the Controller contract should be zero", func() {
		b, e := Backend.BalanceAt(context.Background(), ControllerContractAddress, nil)
		Expect(e).ToNot(HaveOccurred())
		Expect(b.String()).To(Equal("0"))
	})

	It("the initial ERC20 type-1 balance of the Controller contract should be zero", func() {
		b, err := ERC20Contract1.BalanceOf(nil, ControllerContractAddress)
		Expect(err).ToNot(HaveOccurred())
		Expect(b.String()).To(Equal("0"))
	})

	It("the initial ERC20 type-1 balance of the RandomAccount  should be zero", func() {
		b, err := ERC20Contract1.BalanceOf(nil, RandomAccount.Address())
		Expect(err).ToNot(HaveOccurred())
		Expect(b.String()).To(Equal("0"))
	})

	When("The Controller contract is credited with 333 of ERC20 type-1 tokens", func() {
		BeforeEach(func() {
			tx, err := ERC20Contract1.Credit(BankAccount.TransactOpts(), ControllerContractAddress, big.NewInt(333))
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeTrue())
		})

		It("should increase the ERC20 type-1 balance of the Controller contract by 333", func() {
			b, err := ERC20Contract1.BalanceOf(nil, ControllerContractAddress)
			Expect(err).ToNot(HaveOccurred())
			Expect(b.String()).To(Equal("333"))
		})

		When("the Controller admin withdraws 222 tokens", func() {

			BeforeEach(func() {
				tx, err := ControllerContract.Claim(ControllerAdmin.TransactOpts(), RandomAccount.Address(), ERC20Contract1Address, big.NewInt(222))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})

			It("Should emit 2 Transferred events", func() {
				it, err := ControllerContract.FilterClaimed(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(it.Next()).To(BeTrue())
				evt := it.Event
				Expect(it.Next()).To(BeFalse())
				Expect(evt.To).To(Equal(RandomAccount.Address()))
				Expect(evt.Asset).To(Equal(ERC20Contract1Address))
				Expect(evt.Amount.String()).To(Equal("222"))
			})

			It("should decrease the ERC20 type-1 balance of the Controller contract contract by 222", func() {
				b, err := ERC20Contract1.BalanceOf(nil, ControllerContractAddress)
				Expect(err).ToNot(HaveOccurred())
				Expect(b.String()).To(Equal("111"))
			})

			It("should increase the ERC20 type-1 balance of the RandomAccount by 222", func() {
				b, err := ERC20Contract1.BalanceOf(nil, RandomAccount.Address())
				Expect(err).ToNot(HaveOccurred())
				Expect(b.String()).To(Equal("222"))
			})

		}) //owner withdrawal

		When("controller tries to withdraw", func() {
			It("Should fail", func() {
				tx, err := ControllerContract.Claim(Controller.TransactOpts(ethertest.WithGasLimit(100000)), RandomAccount.Address(), ERC20Contract1Address, big.NewInt(222))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isGasExhausted(tx, 100000)).To(BeFalse())
				Expect(isSuccessful(tx)).To(BeFalse())
			})
		})

		When("controller owner tries to withdraw", func() {
			It("Should fail", func() {
				tx, err := ControllerContract.Claim(ControllerOwner.TransactOpts(ethertest.WithGasLimit(100000)), RandomAccount.Address(), ERC20Contract1Address, big.NewInt(222))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isGasExhausted(tx, 100000)).To(BeFalse())
				Expect(isSuccessful(tx)).To(BeFalse())
			})
		})

	})

})
