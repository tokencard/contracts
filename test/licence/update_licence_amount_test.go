package licence_test

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tokencard/contracts/test/shared"
	"github.com/tokencard/ethertest"
)

var _ = Describe("updatedLicenceAmount", func() {

	Context("Before the first update", func() {

		It("should be 1", func() {
			f, err := Licence.LicenceAmountScaled(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(f.String()).To(Equal("10"))
		})

		It("should be pointing to 0x0", func() {
			addr, err := Licence.LicenceDAO(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(addr).To(Equal(common.HexToAddress("0x0")))
		})

	})

	When("called by the DAO", func() {

		var tx *types.Transaction

		BeforeEach(func() {
			var err error
			tx, err = Licence.UpdateLicenceDAO(ControllerAdmin.TransactOpts(), DAO.Address())
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeTrue())
		})

		It("should be updated to the DAO address", func() {
			addr, err := Licence.LicenceDAO(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(addr).To(Equal(DAO.Address()))
		})

		When("1 <= licenceAmount <= MAX_AMOUNT_SCALE", func() {
			BeforeEach(func() {
				var err error
				tx, err = Licence.UpdateLicenceAmount(DAO.TransactOpts(), big.NewInt(100))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
			})

			It("Should succeed", func() {
				Expect(isSuccessful(tx)).To(BeTrue())
			})

			It("Should emit a UpdatedLicenceAmount event", func() {
				it, err := Licence.FilterUpdatedLicenceAmount(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(it.Next()).To(BeTrue())
				evt := it.Event
				Expect(it.Next()).To(BeFalse())
				Expect(evt.NewAmount).To(Equal(big.NewInt(100)))
			})
		}) // 1 <= licenceAmount <= 100

		When("the licenceAmount value is invalid (0 || > MAX_AMOUNT_SCALE)", func() {

			It("Should fail", func() {
				var err error
				tx, err = Licence.UpdateLicenceAmount(DAO.TransactOpts(ethertest.WithGasLimit(100000)), big.NewInt(1001))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isGasExhausted(tx, 100000)).To(BeFalse())
				Expect(isSuccessful(tx)).To(BeFalse())
			})

			It("Should fail", func() {
				var err error
				tx, err = Licence.UpdateLicenceAmount(DAO.TransactOpts(ethertest.WithGasLimit(100000)), big.NewInt(0))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isGasExhausted(tx, 100000)).To(BeFalse())
				Expect(isSuccessful(tx)).To(BeFalse())
			})

		})
	})

	When("not called by the DAO", func() {
		It("Should fail", func() {
			tx, err := Licence.UpdateLicenceAmount(ControllerAdmin.TransactOpts(ethertest.WithGasLimit(100000)), big.NewInt(1))
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isGasExhausted(tx, 100000)).To(BeFalse())
			Expect(isSuccessful(tx)).To(BeFalse())
		})
	})

})
