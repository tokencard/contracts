package licence_test

import (
	"github.com/ethereum/go-ethereum/common"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tokencard/contracts/test/shared"
	"github.com/tokencard/ethertest"
)

var _ = Describe("updateTKN", func() {

	Context("Before updating the TKN contract address", func() {

		It("should be pointing to 0xaAAf91D9b90dF800Df4F55c205fd6989c977E73a", func() {
			addr, err := Licence.TKNContractAddress(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(addr).To(Equal(common.HexToAddress("0xaAAf91D9b90dF800Df4F55c205fd6989c977E73a")))
		})

		It("should NOT be locked", func() {
			lock, err := Licence.TKNContractAddressLocked(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(lock).To(BeFalse())
		})
	})

	When("called by the owner", func() {

		BeforeEach(func() {
			var err error
			tx, err := Licence.UpdateTKNContractAddress(Owner.TransactOpts(), TKNBurnerAddress)
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeTrue())
		})

		It("Should emit a UpdatedTKNContractAddress event", func() {
			it, err := Licence.FilterUpdatedTKNContractAddress(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(it.Next()).To(BeTrue())
			evt := it.Event
			Expect(it.Next()).To(BeFalse())
			Expect(evt.NewTKN).To(Equal(TKNBurnerAddress))
		})

		Context("TKN is not locked after the update", func() {

			BeforeEach(func() {
				tx, err := Licence.UpdateTKNContractAddress(Owner.TransactOpts(), common.HexToAddress("0x1"))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})

			It("should be pointing to 0x1 address", func() {
				addr, err := Licence.TKNContractAddress(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(addr).To(Equal(common.HexToAddress("0x1")))
			})

			It("Should emit a new UpdatedTKNContractAddress event", func() {
				it, err := Licence.FilterUpdatedTKNContractAddress(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(it.Next()).To(BeTrue())
				evt := it.Event
				Expect(it.Next()).To(BeTrue())
				Expect(evt.NewTKN).To(Equal(TKNBurnerAddress))
				evt = it.Event
				Expect(it.Next()).To(BeFalse())
				Expect(evt.NewTKN).To(Equal(common.HexToAddress("0x1")))
			})

		}) //not locked

		Context("TKN is locked after the update", func() {

			BeforeEach(func() {
				tx, err := Licence.LockTKNContractAddress(Owner.TransactOpts())
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})

			BeforeEach(func() {
				tx, err := Licence.UpdateTKNContractAddress(Owner.TransactOpts(ethertest.WithGasLimit(100000)), common.HexToAddress("0x1"))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isGasExhausted(tx, 100000)).To(BeFalse())
				Expect(isSuccessful(tx)).To(BeFalse())
			})

			It("Should be locked", func() {
				lock, err := Licence.TKNContractAddressLocked(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(lock).To(BeTrue())
			})

			It("should be still pointing to TKNBurnerAddress", func() {
				addr, err := Licence.TKNContractAddress(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(addr).To(Equal(TKNBurnerAddress))
			})

			It("Should not emit a new UpdatedTKNContractAddress event", func() {
				it, err := Licence.FilterUpdatedTKNContractAddress(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(it.Next()).To(BeTrue())
				evt := it.Event
				Expect(it.Next()).To(BeFalse())
				Expect(evt.NewTKN).To(Equal(TKNBurnerAddress))
			})

		}) //locked

	}) //called by the owner

	Context("When not called by the Ower", func() {
		It("Should fail", func() {
			tx, err := Licence.UpdateTKNContractAddress(BankAccount.TransactOpts(ethertest.WithGasLimit(100000)), TKNBurnerAddress)
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isGasExhausted(tx, 100000)).To(BeFalse())
			Expect(isSuccessful(tx)).To(BeFalse())
		})
	})

	It("Should fail", func() {
		tx, err := Licence.LockTKNContractAddress(BankAccount.TransactOpts(ethertest.WithGasLimit(100000)))
		Expect(err).ToNot(HaveOccurred())
		Backend.Commit()
		Expect(isGasExhausted(tx, 100000)).To(BeFalse())
		Expect(isSuccessful(tx)).To(BeFalse())
	})

})
