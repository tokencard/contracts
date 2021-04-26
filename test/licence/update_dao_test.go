package licence_test

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tokencard/contracts/v3/test/shared"
	"github.com/tokencard/ethertest"
)

var _ = Describe("updateDao", func() {

	Context("Before updating the DAO contract address", func() {

		It("should be pointing to 0x0", func() {
			addr, err := Licence.LicenceDAO(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(addr).To(Equal(common.HexToAddress("0x0")))
		})

		It("should NOT be locked", func() {
			lock, err := Licence.LicenceDAOLocked(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(lock).To(BeFalse())
		})
	})

	When("called by the controller admin", func() {

		var tx *types.Transaction

		BeforeEach(func() {
			var err error
			tx, err = Licence.UpdateLicenceDAO(ControllerAdmin.TransactOpts(), DAO.Address())
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
		})

		It("Should succeed", func() {
			Expect(isSuccessful(tx)).To(BeTrue())
		})

		It("Should emit a UpdatedLicenceDAO event", func() {
			it, err := Licence.FilterUpdatedLicenceDAO(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(it.Next()).To(BeTrue())
			evt := it.Event
			Expect(it.Next()).To(BeFalse())
			Expect(evt.NewDAO).To(Equal(DAO.Address()))
		})

		Context("The DAO is not locked after the update", func() {

			BeforeEach(func() {
				var err error
				tx, err = Licence.UpdateLicenceDAO(ControllerAdmin.TransactOpts(), common.HexToAddress("0x1"))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
			})

			It("Should be able to be updated again", func() {
				Expect(isSuccessful(tx)).To(BeTrue())
			})

			It("Should emit a new UpdatedLicenceDAO event", func() {
				it, err := Licence.FilterUpdatedLicenceDAO(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(it.Next()).To(BeTrue())
				evt := it.Event
				Expect(it.Next()).To(BeTrue())
				Expect(evt.NewDAO).To(Equal(DAO.Address()))
				evt = it.Event
				Expect(it.Next()).To(BeFalse())
				Expect(evt.NewDAO).To(Equal(common.HexToAddress("0x1")))
			})

		}) //not locked

		Context("The DAO is locked after the update", func() {

			BeforeEach(func() {
				tx, err := Licence.LockLicenceDAO(ControllerAdmin.TransactOpts())
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})

			BeforeEach(func() {
				var err error
				tx, err = Licence.UpdateLicenceDAO(ControllerAdmin.TransactOpts(ethertest.WithGasLimit(100000)), common.HexToAddress("0x1"))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isGasExhausted(tx, 100000)).To(BeFalse())
			})

			It("Should not be possible to update again", func() {
				Expect(isSuccessful(tx)).To(BeFalse())
			})

			It("Should NOT emit a 2nd UpdatedLicenceDAO event", func() {
				it, err := Licence.FilterUpdatedLicenceDAO(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(it.Next()).To(BeTrue())
				evt := it.Event
				Expect(it.Next()).To(BeFalse())
				Expect(evt.NewDAO).To(Equal(DAO.Address()))
			})

		}) //locked

	}) //called by the owner

	Context("When not called by the Owner", func() {
		It("Should fail", func() {
			tx, err := Licence.UpdateLicenceDAO(BankAccount.TransactOpts(ethertest.WithGasLimit(100000)), DAO.Address())
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isGasExhausted(tx, 100000)).To(BeFalse())
			Expect(isSuccessful(tx)).To(BeFalse())
		})
	})

})
