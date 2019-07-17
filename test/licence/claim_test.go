package licence_test

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tokencard/contracts/test/shared"
)

var _ = Describe("LicenceClaim", func() {

	It("the initial balance of the Licence contract should be zero", func() {
		b, e := Backend.BalanceAt(context.Background(), LicenceAddress, nil)
		Expect(e).ToNot(HaveOccurred())
		Expect(b.String()).To(Equal("0"))
	})

	It("the initial ERC20 type-1 balance of the Licence contract should be zero", func() {
		b, err := ERC20Contract1.BalanceOf(nil, LicenceAddress)
		Expect(err).ToNot(HaveOccurred())
		Expect(b.String()).To(Equal("0"))
	})

	It("the initial ERC20 type-1 balance of the RandomAccount  should be zero", func() {
		b, err := ERC20Contract1.BalanceOf(nil, RandomAccount.Address())
		Expect(err).ToNot(HaveOccurred())
		Expect(b.String()).To(Equal("0"))
	})

	When("The Licence contract is credited with 333 of ERC20 type-1 tokens and 1 ETH", func() {
		BeforeEach(func() {
			tx, err := ERC20Contract1.Credit(BankAccount.TransactOpts(), LicenceAddress, big.NewInt(333))
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeTrue())
		})

		BeforeEach(func() {
			BankAccount.MustTransfer(Backend, LicenceAddress, EthToWei(2))
		})

		It("should increase the ERC20 type-1 balance of the Licence contract by 333", func() {
			b, err := ERC20Contract1.BalanceOf(nil, LicenceAddress)
			Expect(err).ToNot(HaveOccurred())
			Expect(b.String()).To(Equal("333"))
		})

		It("should increase the ETH balance of the Licence contract by 2 ETH", func() {
			b, e := Backend.BalanceAt(context.Background(), LicenceAddress, nil)
			Expect(e).ToNot(HaveOccurred())
			Expect(b.String()).To(Equal(EthToWei(2).String()))
		})

		When("the controller admin withdraws 222 tokens and 1 ETH", func() {

			BeforeEach(func() {
				tx, err := Licence.Claim(ControllerAdmin.TransactOpts(), RandomAccount.Address(), ERC20Contract1Address, big.NewInt(222))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})

			BeforeEach(func() {
				tx, err := Licence.Claim(ControllerAdmin.TransactOpts(), RandomAccount.Address(), common.HexToAddress("0x0"), EthToWei(1))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})

			It("Should emit 2 Transferred events", func() {
				it, err := Licence.FilterClaimed(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(it.Next()).To(BeTrue())
				evt := it.Event
				Expect(it.Next()).To(BeTrue())
				Expect(evt.To).To(Equal(RandomAccount.Address()))
				Expect(evt.Asset).To(Equal(ERC20Contract1Address))
				Expect(evt.Amount.String()).To(Equal("222"))
				evt = it.Event
				Expect(it.Next()).To(BeFalse())
				Expect(evt.To).To(Equal(RandomAccount.Address()))
				Expect(evt.Asset).To(Equal(common.HexToAddress("0x0"))) //represents ETH
				Expect(evt.Amount.String()).To(Equal(EthToWei(1).String()))

			})

			It("should decrease the ERC20 type-1 balance of the Licence contract by 222", func() {
				b, err := ERC20Contract1.BalanceOf(nil, LicenceAddress)
				Expect(err).ToNot(HaveOccurred())
				Expect(b.String()).To(Equal("111"))
			})

			It("should decrease the ETH balance of the Licence contract by 1 ETH", func() {
				b, e := Backend.BalanceAt(context.Background(), LicenceAddress, nil)
				Expect(e).ToNot(HaveOccurred())
				Expect(b.String()).To(Equal(EthToWei(1).String()))
			})

			It("should increase the ERC20 type-1 balance of the RandomAccount by 222", func() {
				b, err := ERC20Contract1.BalanceOf(nil, RandomAccount.Address())
				Expect(err).ToNot(HaveOccurred())
				Expect(b.String()).To(Equal("222"))
			})

			It("should increase the ETH balance of the Random account by 1 ETH", func() { //initial balance of RandomAccount is 1000 ETH
				b, e := Backend.BalanceAt(context.Background(), RandomAccount.Address(), nil)
				Expect(e).ToNot(HaveOccurred())
				Expect(b.String()).To(Equal(EthToWei(1001).String()))
			})

		}) //owner withdrawal

	})

})
