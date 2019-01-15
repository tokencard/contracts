package tokenHolder_test

import (
	// "errors"
	"math/big"

  "github.com/ethereum/go-ethereum/core/types"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tokencard/contracts/test/shared"
	// "github.com/tokencard/ethertest"

)

var _ = Describe("TokenHolder", func() {

	Context("A valid Hodler wants to burrrrrnnnn TKN", func() {

    var tx *types.Transaction

    When("one thousand tokens are credited to an external address", func() {
			BeforeEach(func() {
				var err error
				tx, err = TKNBurner.Credit(BankAccount.TransactOpts(), BankAccount.Address(), big.NewInt(1000))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})

      It("should increase the total TKN supply to 1000", func() {
        s, err := TKNBurner.TotalSupply(nil)
        Expect(err).ToNot(HaveOccurred())
        Expect(s.String()).To(Equal("1000"))
      })

      It("should increase TKN balance of the specific address", func() {
        b, err := TKNBurner.BalanceOf(nil, BankAccount.Address())
        Expect(err).ToNot(HaveOccurred())
        Expect(b.String()).To(Equal("1000"))
      })

      When("300 tokens are transfered to a random address", func() {
				BeforeEach(func() {
					tx, err := TKNBurner.Transfer(BankAccount.TransactOpts(), RandomAccount.Address(), big.NewInt(300))
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeTrue())
				})

				It("should increase TKN balance of the random person", func() {
					b, err := TKNBurner.BalanceOf(nil, RandomAccount.Address())
					Expect(err).ToNot(HaveOccurred())
					Expect(b.String()).To(Equal("300"))
				})

				It("should decrease TKN balance of the initial address", func() {
					b, err := TKNBurner.BalanceOf(nil, BankAccount.Address())
					Expect(err).ToNot(HaveOccurred())
					Expect(b.String()).To(Equal("700"))
				})

        It("should leave the total supply intact", func() {
          s, err := TKNBurner.TotalSupply(nil)
          Expect(err).ToNot(HaveOccurred())
          Expect(s.String()).To(Equal("1000"))
				})

				When("When the token holder is set by the owner", func() {

					BeforeEach(func() {
						tx, err := TKNBurner.SetTokenHolder(Owner.TransactOpts(), TokenHolderAddress)
						Expect(err).ToNot(HaveOccurred())
						Backend.Commit()
						Expect(isSuccessful(tx)).To(BeTrue())
					})

					When("When the random address burns 200 tokens (out of 1000)", func() {

						BeforeEach(func() {
							tx, err := TKNBurner.Burn(RandomAccount.TransactOpts(), big.NewInt(200))
							Expect(err).ToNot(HaveOccurred())
							Backend.Commit()
							Expect(isSuccessful(tx)).To(BeTrue())
						})

						It("should decrease the total supply by 200 (800 remaining)", func() {
		          s, err := TKNBurner.TotalSupply(nil)
		          Expect(err).ToNot(HaveOccurred())
 	 	          Expect(s.String()).To(Equal("800"))
						})

						FIt("should decrease TKN balance of the random address", func() {
							b, err := TKNBurner.BalanceOf(nil, RandomAccount.Address())
							Expect(err).ToNot(HaveOccurred())
							Expect(b.String()).To(Equal("100"))
						})

						It("Should emit a Burn event", func() {
							it, err := TKNBurner.FilterBurn(nil)
							Expect(err).ToNot(HaveOccurred())
							Expect(it.Next()).To(BeTrue())
							evt := it.Event
							Expect(it.Next()).To(BeFalse())
							Expect(evt.Burner).To(Equal(RandomAccount.Address()))
							Expect(evt.Amount.String()).To(Equal("200"))
						})

					})
				})
    })

    // BeforeEach(func() {
    //   var err error
    //   tx, err = TKNBurner.Burn(RandomAccount.TransactOpts(), big.NewInt(500))
    //   Expect(err).ToNot(HaveOccurred())
    //   Backend.Commit()
    // })
    //
    // It("Should succeed", func() {
    //   Expect(isSuccessful(tx)).To(BeTrue())
    // })

	 })
  })


})
