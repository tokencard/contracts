package licence_test

import (
	"github.com/ethereum/go-ethereum/core/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tokencard/contracts/v3/test/shared"
	"github.com/tokencard/ethertest"
)

var _ = Describe("updateHolder", func() {

	Context("Before updating the holder contract address", func() {

		It("should be pointing to TokenHolder Address", func() {
			addr, err := Licence.TokenHolder(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(addr).To(Equal(TokenHolderAddress))
		})

		It("should NOT be locked", func() {
			lock, err := Licence.HolderLocked(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(lock).To(BeFalse())
		})
	})

	When("called by the controller admin", func() {

		var tx *types.Transaction

		BeforeEach(func() {
			var err error
			tx, err = Licence.UpdateHolder(ControllerAdmin.TransactOpts(), CryptoFloatAddress)
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
		})

		It("Should succeed", func() {
			Expect(isSuccessful(tx)).To(BeTrue())
		})

		It("Should emit an UpdatedTokenHolder event", func() {
			it, err := Licence.FilterUpdatedTokenHolder(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(it.Next()).To(BeTrue())
			evt := it.Event
			Expect(it.Next()).To(BeFalse())
			Expect(evt.NewHolder).To(Equal(CryptoFloatAddress))
		})

		Context("TokenHolder is not locked after the update", func() {

			BeforeEach(func() {
				var err error
				tx, err = Licence.UpdateHolder(ControllerAdmin.TransactOpts(), TokenHolderAddress)
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
			})

			It("Should be able to be updated again", func() {
				Expect(isSuccessful(tx)).To(BeTrue())
			})

			It("Should emit a new UpdatedTokenHolder event", func() {
				it, err := Licence.FilterUpdatedTokenHolder(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(it.Next()).To(BeTrue())
				evt := it.Event
				Expect(it.Next()).To(BeTrue())
				Expect(evt.NewHolder).To(Equal(CryptoFloatAddress))
				evt = it.Event
				Expect(it.Next()).To(BeFalse())
				Expect(evt.NewHolder).To(Equal(TokenHolderAddress))
			})

		}) //not locked

		Context("TokenHolder is locked after the update", func() {

			BeforeEach(func() {
				tx, err := Licence.LockHolder(ControllerAdmin.TransactOpts())
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})

			BeforeEach(func() {
				var err error
				tx, err = Licence.UpdateHolder(ControllerAdmin.TransactOpts(ethertest.WithGasLimit(100000)), TokenHolderAddress)
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isGasExhausted(tx, 100000)).To(BeFalse())
				Expect(isSuccessful(tx)).To(BeFalse())
			})

			It("Should not be possible to update again", func() {
				Expect(isSuccessful(tx)).To(BeFalse())
			})

			It("Should NOT emit a 2nd UpdateHolder event", func() {
				it, err := Licence.FilterUpdatedTokenHolder(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(it.Next()).To(BeTrue())
				evt := it.Event
				Expect(it.Next()).To(BeFalse())
				Expect(evt.NewHolder).To(Equal(CryptoFloatAddress))
			})

		}) //locked

	}) //called by the owner

	Context("When not called by the Owner", func() {
		It("Should fail", func() {
			tx, err := Licence.UpdateHolder(BankAccount.TransactOpts(ethertest.WithGasLimit(100000)), CryptoFloatAddress)
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isGasExhausted(tx, 100000)).To(BeFalse())
			Expect(isSuccessful(tx)).To(BeFalse())
		})
	})

})
