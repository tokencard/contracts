package licence_test

import (

  "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tokencard/contracts/test/shared"
	"github.com/tokencard/ethertest"
)

var _ = Describe("updateDao", func() {


  Context("Before updating the DAO contract address", func() {

    It("should be pointing to 0x0", func() {
      addr, err := Licence.DAO(nil)
      Expect(err).ToNot(HaveOccurred())
      Expect(addr).To(Equal(common.HexToAddress("0x0")))
    })

    It("should NOT be locked", func() {
      lock, err := Dao.IsLocked(nil)
      Expect(err).ToNot(HaveOccurred())
      Expect(lock).To(BeFalse())
    })
  })

	When("called by the owner", func() {

		var tx *types.Transaction

		BeforeEach(func() {
			var err error
			tx, err = Licence.UpdateDAO(Owner.TransactOpts(), DaoAddress)
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
		})

		It("Should succeed", func() {
			Expect(isSuccessful(tx)).To(BeTrue())
		})

		It("Should emit a UpdatedDAO event", func() {
			it, err := Licence.FilterUpdatedDAO(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(it.Next()).To(BeTrue())
			evt := it.Event
			Expect(it.Next()).To(BeFalse())
      Expect(evt.Sender).To(Equal(Owner.Address()))
			Expect(evt.NewDAO).To(Equal(DaoAddress))
		})

    Context("It DAO is not locked after the update", func(){

      BeforeEach(func() {
  			var err error
  			tx, err = Licence.UpdateDAO(Owner.TransactOpts(), common.HexToAddress("0x1"))
  			Expect(err).ToNot(HaveOccurred())
  			Backend.Commit()
  		})

  		It("Should be able to be updated again", func() {
  			Expect(isSuccessful(tx)).To(BeTrue())
  		})

  		It("Should emit a new UpdatedDAO event", func() {
  			it, err := Licence.FilterUpdatedDAO(nil)
  			Expect(err).ToNot(HaveOccurred())
  			Expect(it.Next()).To(BeTrue())
  			evt := it.Event
  			Expect(it.Next()).To(BeTrue())
        Expect(evt.Sender).To(Equal(Owner.Address()))
  			Expect(evt.NewDAO).To(Equal(DaoAddress))
        evt = it.Event
  			Expect(it.Next()).To(BeFalse())
        Expect(evt.Sender).To(Equal(Owner.Address()))
  			Expect(evt.NewDAO).To(Equal(common.HexToAddress("0x1")))
  		})

  	}) //not locked

    Context("The DAO is locked after the update", func(){

      BeforeEach(func() {
  			tx, err := Dao.Lock(Owner.TransactOpts())
  			Expect(err).ToNot(HaveOccurred())
  			Backend.Commit()
        Expect(isSuccessful(tx)).To(BeTrue())
  		})

      It("Should not be possible to update again", func() {
        tx, err := Licence.UpdateDAO(Owner.TransactOpts(ethertest.WithGasLimit(100000)), DaoAddress)
  			Expect(err).ToNot(HaveOccurred())
  			Backend.Commit()
        Expect(isGasExhausted(tx, 100000)).To(BeFalse())
  			Expect(isSuccessful(tx)).To(BeFalse())
  		})

  		It("Should emit a UpdatedDAO event", func() {
  			it, err := Licence.FilterUpdatedDAO(nil)
  			Expect(err).ToNot(HaveOccurred())
  			Expect(it.Next()).To(BeTrue())
  			evt := it.Event
  			Expect(it.Next()).To(BeFalse())
        Expect(evt.Sender).To(Equal(Owner.Address()))
  			Expect(evt.NewDAO).To(Equal(DaoAddress))
  		})

    }) //locked

  }) //called by the owner

	Context("When not called by the Ower", func() {
		It("Should fail", func() {
			tx, err := Licence.UpdateDAO(BankAccount.TransactOpts(ethertest.WithGasLimit(100000)), DaoAddress)
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isGasExhausted(tx, 100000)).To(BeFalse())
			Expect(isSuccessful(tx)).To(BeFalse())
		})
	})

})
