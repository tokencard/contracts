package tokenHolder_test

import (
	// "errors"
	"math/big"

  "github.com/ethereum/go-ethereum/core/types"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tokencard/contracts/test/shared"
	// "github.com/tokencard/ethertest"
	"github.com/ethereum/go-ethereum/common"

)

var _ = Describe("TokenHolder", func() {

	Context("A valid Hodler wants to burrrrrnnnn TKN", func() {

    var tx *types.Transaction

    When("one thousand TKN are credited to an external address", func() {
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

      When("300 TKN are transfered to a random address", func() {
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

				When("The holder contract owns two types of ERC20 tokens", func() {

					//add the tokens to the list
					BeforeEach(func() {
						var err error
						tokens := []common.Address{ERC20Contract1Address, ERC20Contract2Address}
						tx, err = TokenHolder.SetTokens(Owner.TransactOpts(), tokens)
						Expect(err).ToNot(HaveOccurred())
						Backend.Commit()
						Expect(isSuccessful(tx)).To(BeTrue())
					})

					BeforeEach(func() {
						var err error
						tx, err = ERC20Contract1.Credit(BankAccount.TransactOpts(), TokenHolderAddress, big.NewInt(1000))
						Expect(err).ToNot(HaveOccurred())
						Backend.Commit()
						Expect(isSuccessful(tx)).To(BeTrue())
					})

					BeforeEach(func() {
						var err error
						tx, err = ERC20Contract2.Credit(BankAccount.TransactOpts(), TokenHolderAddress, big.NewInt(1000))
						Expect(err).ToNot(HaveOccurred())
						Backend.Commit()
						Expect(isSuccessful(tx)).To(BeTrue())

					})

					When("When the token holder (contract) is set by the owner", func() {

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

							It("should decrease TKN balance of the random address", func() {
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

							It("should increase the ERC20 type-1 balance of the random address by 200 (20%)", func() {
								b, err := ERC20Contract1.BalanceOf(nil, RandomAccount.Address())
								Expect(err).ToNot(HaveOccurred())
								Expect(b.String()).To(Equal("200"))
							})

							It("should increase the ERC20 type-2 balance of the random address by 200 (20%)", func() {
								b, err := ERC20Contract2.BalanceOf(nil, RandomAccount.Address())
								Expect(err).ToNot(HaveOccurred())
								Expect(b.String()).To(Equal("200"))
							})

						})
					})
				}) //When("The holder contract has two types of ERC20 tokens"
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
