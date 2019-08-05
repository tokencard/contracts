package licence_test

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tokencard/contracts/v2/test/shared"
	"github.com/tokencard/ethertest"
)

var _ = Describe("load TKN", func() {

	//Update it to show to the TKN contract deployed for testing
	BeforeEach(func() {
		tx, err := Licence.UpdateTKNContractAddress(ControllerAdmin.TransactOpts(), TKNBurnerAddress)
		Expect(err).ToNot(HaveOccurred())
		Backend.Commit()
		Expect(isSuccessful(tx)).To(BeTrue())
	})

	It("should be pointing to TKNBurnerAddress", func() {
		addr, err := Licence.TknContractAddress(nil)
		Expect(err).ToNot(HaveOccurred())
		Expect(addr).To(Equal(TKNBurnerAddress))
	})

	When("one thousand TKN are credited to a random acount", func() {

		BeforeEach(func() {
			tx, err := TKNBurner.Mint(BankAccount.TransactOpts(), RandomAccount.Address(), big.NewInt(1000))
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeTrue())
		})

		It("should increase the total TKN supply to 1000", func() {
			s, err := TKNBurner.TotalSupply(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(s.String()).To(Equal("1000"))
		})

		It("should increase TKN balance of the random account", func() {
			b, err := TKNBurner.BalanceOf(nil, RandomAccount.Address())
			Expect(err).ToNot(HaveOccurred())
			Expect(b.String()).To(Equal("1000"))
		})

		When("no approval is given", func() {

			It("Should revert", func() {
				tx, err := Licence.Load(RandomAccount.TransactOpts(ethertest.WithGasLimit(100000)), TKNBurnerAddress, big.NewInt(100))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isGasExhausted(tx, 100000)).To(BeFalse())
				Expect(isSuccessful(tx)).To(BeFalse())
			})

		})

		When("approval to Licence contract to transfer 444 tokens respectively is given", func() {

			BeforeEach(func() {
				tx, err := TKNBurner.Approve(RandomAccount.TransactOpts(), LicenceAddress, big.NewInt(444))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})

			It("Should emit 1 Approval event", func() {
				owner := []common.Address{RandomAccount.Address()}
				spender := []common.Address{LicenceAddress}
				it, err := TKNBurner.FilterApproval(nil, owner, spender)
				Expect(err).ToNot(HaveOccurred())
				Expect(it.Next()).To(BeTrue())
				evt := it.Event
				Expect(it.Next()).To(BeFalse())
				Expect(evt.Owner).To(Equal(RandomAccount.Address()))
				Expect(evt.Spender).To(Equal(LicenceAddress))
				Expect(evt.Value.String()).To(Equal("444"))
			})

			When("all the approved tokens are transfered ", func() {

				BeforeEach(func() {
					tx, err := Licence.Load(RandomAccount.TransactOpts(), TKNBurnerAddress, big.NewInt(444))
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeTrue())
				})

				It("Should NOT emit a TransferredToTokenHolder event", func() {
					it, err := Licence.FilterTransferredToTokenHolder(nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(it.Next()).To(BeFalse())
				})

				It("Should emit 1 TransferredToCryptoFloat event", func() {
					it, err := Licence.FilterTransferredToCryptoFloat(nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(it.Next()).To(BeTrue())
					evt := it.Event
					Expect(it.Next()).To(BeFalse())
					Expect(evt.From).To(Equal(RandomAccount.Address()))
					Expect(evt.To).To(Equal(CryptoFloatAddress))
					Expect(evt.Asset).To(Equal(TKNBurnerAddress))
					Expect(evt.Amount.String()).To(Equal("444"))
				})

				It("Should emit 1 TKNBurner Transfer event", func() {
					from := []common.Address{RandomAccount.Address()}
					var to []common.Address
					it, err := TKNBurner.FilterTransfer(nil, from, to)
					Expect(err).ToNot(HaveOccurred())
					Expect(it.Next()).To(BeTrue())
					evt := it.Event
					Expect(it.Next()).To(BeFalse())
					Expect(evt.From).To(Equal(RandomAccount.Address()))
					Expect(evt.To).To(Equal(CryptoFloatAddress))
					Expect(evt.Value.String()).To(Equal("444"))
				})

				It("should increase the TKN balance of the Float contract by 444", func() {
					b, err := TKNBurner.BalanceOf(nil, CryptoFloatAddress)
					Expect(err).ToNot(HaveOccurred())
					Expect(b.String()).To(Equal("444"))
				})

				It("should NOT increase the TKN balance of the Holder contract", func() {
					b, err := TKNBurner.BalanceOf(nil, TokenHolderAddress)
					Expect(err).ToNot(HaveOccurred())
					Expect(b.String()).To(Equal("0"))
				})

				It("should decrease the TKN balance of the RandomAccount by 444", func() {
					b, err := TKNBurner.BalanceOf(nil, RandomAccount.Address())
					Expect(err).ToNot(HaveOccurred())
					Expect(b.String()).To(Equal("556")) //1000-444
				})

			}) //equal to approval

			When("more tokens than approved are being transfered ", func() {

				It("Should revert", func() {
					tx, err := Licence.Load(RandomAccount.TransactOpts(ethertest.WithGasLimit(100000)), TKNAddress, big.NewInt(555))
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isGasExhausted(tx, 100000)).To(BeFalse())
					Expect(isSuccessful(tx)).To(BeFalse())
				})

			}) //more than approved

		}) //approval is given
	})

})
