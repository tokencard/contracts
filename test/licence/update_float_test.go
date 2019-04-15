package licence_test

import (
	"github.com/ethereum/go-ethereum/core/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tokencard/contracts/test/shared"
	"github.com/tokencard/ethertest"
)

var _ = Describe("updateFloat", func() {

	Context("Before updating the float address", func() {

		It("should be pointing to CryptoFloat Address Address", func() {
			addr, err := Licence.CryptoFloat(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(addr).To(Equal(CryptoFloatAddress))
		})

		It("should NOT be locked", func() {
			lock, err := Licence.FloatLocked(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(lock).To(BeFalse())
		})
	})

	When("called by the owner", func() {

		var tx *types.Transaction

		BeforeEach(func() {
			var err error
			tx, err = Licence.UpdateFloat(Owner.TransactOpts(), TokenHolderAddress)
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
		})

		It("Should succeed", func() {
			Expect(isSuccessful(tx)).To(BeTrue())
		})

		It("Should emit an UpdatedCryptoFloat event", func() {
			it, err := Licence.FilterUpdatedCryptoFloat(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(it.Next()).To(BeTrue())
			evt := it.Event
			Expect(it.Next()).To(BeFalse())
			Expect(evt.NewFloat).To(Equal(TokenHolderAddress))
		})

		Context("CryptoFloat is not locked after the update", func() {

			BeforeEach(func() {
				var err error
				tx, err = Licence.UpdateFloat(Owner.TransactOpts(), CryptoFloatAddress)
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
			})

			It("Should be able to be updated again", func() {
				Expect(isSuccessful(tx)).To(BeTrue())
			})

			It("Should emit a new UpdatedCryptoFloat event", func() {
				it, err := Licence.FilterUpdatedCryptoFloat(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(it.Next()).To(BeTrue())
				evt := it.Event
				Expect(it.Next()).To(BeTrue())
				Expect(evt.NewFloat).To(Equal(TokenHolderAddress))
				evt = it.Event
				Expect(it.Next()).To(BeFalse())
				Expect(evt.NewFloat).To(Equal(CryptoFloatAddress))
			})

		}) //not locked

		Context("CryptoFloat is locked after the update", func() {

			BeforeEach(func() {
				tx, err := Licence.LockFloat(Owner.TransactOpts())
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})

			BeforeEach(func() {
				var err error
				tx, err = Licence.UpdateFloat(Owner.TransactOpts(ethertest.WithGasLimit(100000)), CryptoFloatAddress)
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isGasExhausted(tx, 100000)).To(BeFalse())
			})

			It("Should not be possible to update again", func() {
				Expect(isSuccessful(tx)).To(BeFalse())
			})

			It("Should NOT emit a 2nd UpdateFloat event", func() {
				it, err := Licence.FilterUpdatedCryptoFloat(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(it.Next()).To(BeTrue())
				evt := it.Event
				Expect(it.Next()).To(BeFalse())
				Expect(evt.NewFloat).To(Equal(TokenHolderAddress))
			})

		}) //locked

	}) //called by the owner

	Context("When not called by the Ower", func() {
		It("Should fail", func() {
			tx, err := Licence.UpdateHolder(BankAccount.TransactOpts(ethertest.WithGasLimit(100000)), CryptoFloatAddress)
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isGasExhausted(tx, 100000)).To(BeFalse())
			Expect(isSuccessful(tx)).To(BeFalse())
		})
	})

})
