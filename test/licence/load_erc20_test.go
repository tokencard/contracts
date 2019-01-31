package licence_test

import (

  "math/big"
  "context"

  "github.com/ethereum/go-ethereum/common"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tokencard/contracts/test/shared"
	"github.com/tokencard/ethertest"
)

var _ = Describe("load", func() {


    It("the initial balance of the Holder contract should be zero", func() {
      b, e := Backend.BalanceAt(context.Background(), CryptoFloatAddress, nil)
      Expect(e).ToNot(HaveOccurred())
      Expect(b.String()).To(Equal("0"))
    })

    It("the initial balance of CryptoFloatAddress address should be zero", func() {
      b, e := Backend.BalanceAt(context.Background(), TokenHolderAddress, nil)
      Expect(e).ToNot(HaveOccurred())
      Expect(b.String()).To(Equal("0"))
    })

    When("The RandomAccount is credited with two types of ERC20 tokens", func() {
      BeforeEach(func() {
        tx, err := ERC20Contract1.Credit(BankAccount.TransactOpts(), RandomAccount.Address(), big.NewInt(1000))
        Expect(err).ToNot(HaveOccurred())
        Backend.Commit()
        Expect(isSuccessful(tx)).To(BeTrue())
      })

      BeforeEach(func() {
        tx, err := ERC20Contract2.Credit(BankAccount.TransactOpts(), RandomAccount.Address(), big.NewInt(500))
        Expect(err).ToNot(HaveOccurred())
        Backend.Commit()
        Expect(isSuccessful(tx)).To(BeTrue())
      })

      When("no approval is given", func() {

        It("Should revert", func() {
    			tx, err := Licence.Load(RandomAccount.TransactOpts(ethertest.WithGasLimit(100000)), ERC20Contract1Address, big.NewInt(100))
          Expect(err).ToNot(HaveOccurred())
    			Backend.Commit()
    			Expect(isGasExhausted(tx, 100000)).To(BeFalse())
    			Expect(isSuccessful(tx)).To(BeFalse())
    		})

        It("Should revert", func() {
    			tx, err := Licence.Load(RandomAccount.TransactOpts(ethertest.WithGasLimit(100000)), ERC20Contract2Address, big.NewInt(50))
          Expect(err).ToNot(HaveOccurred())
    			Backend.Commit()
    			Expect(isGasExhausted(tx, 100000)).To(BeFalse())
    			Expect(isSuccessful(tx)).To(BeFalse())
    		})
      })

      When("approval to Licence contract to transfer 100 and 50 tokens respectively is given", func() {

        BeforeEach(func() {
          tx, err := ERC20Contract1.Approve(RandomAccount.TransactOpts(), LicenceAddress, big.NewInt(101))
          Expect(err).ToNot(HaveOccurred())
    			Backend.Commit()
    			Expect(isSuccessful(tx)).To(BeTrue())
        })

        BeforeEach(func() {
          tx, err := ERC20Contract2.Approve(RandomAccount.TransactOpts(), LicenceAddress, big.NewInt(2))
          Expect(err).ToNot(HaveOccurred())
    			Backend.Commit()
    			Expect(isSuccessful(tx)).To(BeTrue())
        })

        It("Should emit 1 ERC20Contract1 Approval event", func() {
          owner := []common.Address{RandomAccount.Address()}
          spender := []common.Address{LicenceAddress}
          it, err := ERC20Contract1.FilterApproval(nil,owner,spender)
          Expect(err).ToNot(HaveOccurred())
          Expect(it.Next()).To(BeTrue())
          evt := it.Event
          Expect(it.Next()).To(BeFalse())
          Expect(evt.Owner).To(Equal(RandomAccount.Address()))
          Expect(evt.Spender).To(Equal(LicenceAddress))
          Expect(evt.Value.String()).To(Equal("101"))
        })

        It("Should emit 1 ERC20Contract2 Approval event", func() {
          owner := []common.Address{RandomAccount.Address()}
          spender := []common.Address{LicenceAddress}
          it, err := ERC20Contract2.FilterApproval(nil,owner,spender)
          Expect(err).ToNot(HaveOccurred())
          Expect(it.Next()).To(BeTrue())
          evt := it.Event
          Expect(it.Next()).To(BeFalse())
          Expect(evt.Owner).To(Equal(RandomAccount.Address()))
          Expect(evt.Spender).To(Equal(LicenceAddress))
          Expect(evt.Value.String()).To(Equal("2"))
        })

        When("the exact approved amount is transfered ", func() {

          BeforeEach(func() {
      			tx, err := Licence.Load(RandomAccount.TransactOpts(), ERC20Contract1Address, big.NewInt(101))
            Expect(err).ToNot(HaveOccurred())
      			Backend.Commit()
      			Expect(isSuccessful(tx)).To(BeTrue())
      		})

          BeforeEach(func() {
      			tx, err := Licence.Load(RandomAccount.TransactOpts(), ERC20Contract2Address, big.NewInt(2))
            Expect(err).ToNot(HaveOccurred())
      			Backend.Commit()
      			Expect(isSuccessful(tx)).To(BeTrue())
      		})

          It("Should emit 2 TransferredToTokenHolder events", func() {
      			it, err := Licence.FilterTransferredToTokenHolder(nil)
      			Expect(err).ToNot(HaveOccurred())
      			Expect(it.Next()).To(BeTrue())
      			evt := it.Event
      			Expect(it.Next()).To(BeTrue())
            Expect(evt.From).To(Equal(RandomAccount.Address()))
      			Expect(evt.To).To(Equal(TokenHolderAddress))
            Expect(evt.Asset).To(Equal(ERC20Contract1Address)) //represents ETH
            Expect(evt.Amount.String()).To(Equal("1"))
            evt = it.Event
      			Expect(it.Next()).To(BeFalse())
            Expect(evt.From).To(Equal(RandomAccount.Address()))
      			Expect(evt.To).To(Equal(TokenHolderAddress))
            Expect(evt.Asset).To(Equal(ERC20Contract2Address)) //represents ETH
            Expect(evt.Amount.String()).To(Equal("1"))
      		})

          It("Should emit 2 TransferredToCryptoFloat events", func() {
      			it, err := Licence.FilterTransferredToCryptoFloat(nil)
      			Expect(err).ToNot(HaveOccurred())
      			Expect(it.Next()).To(BeTrue())
      			evt := it.Event
      			Expect(it.Next()).To(BeTrue())
            Expect(evt.From).To(Equal(RandomAccount.Address()))
      			Expect(evt.To).To(Equal(CryptoFloatAddress))
            Expect(evt.Asset).To(Equal(ERC20Contract1Address)) //represents ETH
            Expect(evt.Amount.String()).To(Equal("100"))
            evt = it.Event
      			Expect(it.Next()).To(BeFalse())
            Expect(evt.From).To(Equal(RandomAccount.Address()))
      			Expect(evt.To).To(Equal(CryptoFloatAddress))
            Expect(evt.Asset).To(Equal(ERC20Contract2Address)) //represents ETH
            Expect(evt.Amount.String()).To(Equal("1"))
      		})


          It("Should emit 2 ERC20Contract1 Transfer events", func() {
            from := []common.Address{RandomAccount.Address()}
            var to []common.Address
            it, err := ERC20Contract1.FilterTransfer(nil, from, to)
            Expect(err).ToNot(HaveOccurred())
            Expect(it.Next()).To(BeTrue())
            evt := it.Event
            Expect(it.Next()).To(BeTrue())
            Expect(evt.From).To(Equal(RandomAccount.Address()))
            Expect(evt.To).To(Equal(TokenHolderAddress))
            Expect(evt.Amount.String()).To(Equal("1"))
            evt = it.Event
      			Expect(it.Next()).To(BeFalse())
            Expect(evt.From).To(Equal(RandomAccount.Address()))
      			Expect(evt.To).To(Equal(CryptoFloatAddress))
            Expect(evt.Amount.String()).To(Equal("100"))
          })

          It("Should emit 1 ERC20Contract2 Transfer events", func() {
            from := []common.Address{RandomAccount.Address()}
            var to []common.Address
            it, err := ERC20Contract2.FilterTransfer(nil, from, to)
            Expect(err).ToNot(HaveOccurred())
            Expect(it.Next()).To(BeTrue())
            evt := it.Event
            Expect(it.Next()).To(BeTrue())
            Expect(evt.From).To(Equal(RandomAccount.Address()))
            Expect(evt.To).To(Equal(TokenHolderAddress))
            Expect(evt.Amount.String()).To(Equal("1"))
            evt = it.Event
      			Expect(it.Next()).To(BeFalse())
            Expect(evt.From).To(Equal(RandomAccount.Address()))
      			Expect(evt.To).To(Equal(CryptoFloatAddress))
            Expect(evt.Amount.String()).To(Equal("1"))
          })

          It("should increase the ERC20 type-1 balance of the Holder contract by 1", func() {
            b, err := ERC20Contract1.BalanceOf(nil, TokenHolderAddress)
            Expect(err).ToNot(HaveOccurred())
            Expect(b.String()).To(Equal("1"))
          })

          It("should increase the ERC20 type-2 balance of the Holder contract by 1", func() {
            b, err := ERC20Contract2.BalanceOf(nil, TokenHolderAddress)
            Expect(err).ToNot(HaveOccurred())
            Expect(b.String()).To(Equal("1"))
          })

          It("should increase the ERC20 type-1 balance of the Float contract by 100", func() {
            b, err := ERC20Contract1.BalanceOf(nil, CryptoFloatAddress)
            Expect(err).ToNot(HaveOccurred())
            Expect(b.String()).To(Equal("100"))
          })

          It("should increase the ERC20 type-2 balance of the Float contract by 1", func() {
            b, err := ERC20Contract2.BalanceOf(nil, CryptoFloatAddress)
            Expect(err).ToNot(HaveOccurred())
            Expect(b.String()).To(Equal("1"))
          })

          It("should decrease the ERC20 type-1 balance of the RandomAccount 101", func() {
            b, err := ERC20Contract1.BalanceOf(nil, RandomAccount.Address())
            Expect(err).ToNot(HaveOccurred())
            Expect(b.String()).To(Equal("899"))
          })

          It("should decrease the ERC20 type-2 balance of the RandomAccount by 2", func() {
            b, err := ERC20Contract2.BalanceOf(nil, RandomAccount.Address())
            Expect(err).ToNot(HaveOccurred())
            Expect(b.String()).To(Equal("498"))
          })

        }) //equal to approval


        When("a bigger amount than the approved one is tried to be transfered ", func() {

          It("Should revert", func() {
      			tx, err := Licence.Load(RandomAccount.TransactOpts(ethertest.WithGasLimit(100000)), ERC20Contract1Address, big.NewInt(102))
            Expect(err).ToNot(HaveOccurred())
      			Backend.Commit()
      			Expect(isGasExhausted(tx, 100000)).To(BeFalse())
      			Expect(isSuccessful(tx)).To(BeFalse())
      		})

          It("Should revert", func() {
      			tx, err := Licence.Load(RandomAccount.TransactOpts(ethertest.WithGasLimit(100000)), ERC20Contract2Address, big.NewInt(3))
            Expect(err).ToNot(HaveOccurred())
      			Backend.Commit()
      			Expect(isGasExhausted(tx, 100000)).To(BeFalse())
      			Expect(isSuccessful(tx)).To(BeFalse())
      		})

        }) //more than approved

      }) //approval is given


    }) //credited with two types of ERC20 tokens

})
