package licence_test

import (

  "math/big"

	"github.com/ethereum/go-ethereum/core/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tokencard/contracts/test/shared"
	"github.com/tokencard/ethertest"
)

var _ = Describe("updateFee", func() {


  Context("Before the first fee update", func() {

    It("should be 1", func() {
      f, err := Licence.Fee(nil)
      Expect(err).ToNot(HaveOccurred())
      Expect(f.String()).To(Equal("1"))
    })
  })


	When("called by the DAO", func() {

		var tx *types.Transaction

    //update the address of the DAO contract, when the Licence contract is first deployed the interface points to 0x0
    BeforeEach(func() {
			var err error
			tx, err = Licence.UpdateDAO(Owner.TransactOpts(), DaoAddress)
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
      Expect(isSuccessful(tx)).To(BeTrue())
		})

    When("1 <= fee <= 100", func() {
  		BeforeEach(func() {
  			var err error
  			tx, err = Dao.UpdateLicenceFee(Owner.TransactOpts(), big.NewInt(100))
  			Expect(err).ToNot(HaveOccurred())
  			Backend.Commit()
  		})

  		It("Should succeed", func() {
  			Expect(isSuccessful(tx)).To(BeTrue())
  		})

  		It("Should emit a UpdatedFee event", func() {
  			it, err := Licence.FilterUpdatedFee(nil)
  			Expect(err).ToNot(HaveOccurred())
  			Expect(it.Next()).To(BeTrue())
  			evt := it.Event
  			Expect(it.Next()).To(BeFalse())
        Expect(evt.Sender).To(Equal(DaoAddress))
  			Expect(evt.NewFee).To(Equal(big.NewInt(100)))
  		})
  	})// 1 <= fee <= 100

    When("the fee value is invalid (0 || >100)", func() {

      It("Should fail", func() {
  			var err error
  			tx, err = Dao.UpdateLicenceFee(Owner.TransactOpts(ethertest.WithGasLimit(100000)), big.NewInt(101))
  			Expect(err).ToNot(HaveOccurred())
  			Backend.Commit()
        Expect(isGasExhausted(tx, 100000)).To(BeFalse())
  			Expect(isSuccessful(tx)).To(BeFalse())
  		})

      It("Should fail", func() {
  			var err error
  			tx, err = Dao.UpdateLicenceFee(Owner.TransactOpts(ethertest.WithGasLimit(100000)), big.NewInt(0))
  			Expect(err).ToNot(HaveOccurred())
  			Backend.Commit()
        Expect(isGasExhausted(tx, 100000)).To(BeFalse())
  			Expect(isSuccessful(tx)).To(BeFalse())
  		})

    })
  })


	When("not called by the DAO", func() {
		It("Should fail", func() {
			tx, err := Licence.UpdateFee(Owner.TransactOpts(ethertest.WithGasLimit(100000)), big.NewInt(1))
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isGasExhausted(tx, 100000)).To(BeFalse())
			Expect(isSuccessful(tx)).To(BeFalse())
		})
	})

})
