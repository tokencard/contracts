package wallet_test

import (

  "math/big"
  "context"

  "github.com/ethereum/go-ethereum/common"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tokencard/contracts/test/shared"
	"github.com/tokencard/ethertest"
)

var _ = Describe("wallet load eth", func() {


    When("no value is sent", func() {

      It("Should revert", func() {
		tx, err := Wallet.LoadTokenCard(Owner.TransactOpts(ethertest.WithGasLimit(100000)), common.HexToAddress("0x0"), big.NewInt(1000))
        Expect(err).ToNot(HaveOccurred())
		Backend.Commit()
		Expect(isGasExhausted(tx, 100000)).To(BeFalse())
		Expect(isSuccessful(tx)).To(BeFalse())
  		})
    })

    When("not called by the owner", func() {

      It("Should revert", func() {
          tx, err := Wallet.LoadTokenCard(RandomAccount.TransactOpts(ethertest.WithGasLimit(100000),ethertest.WithValue(EthToWei(1))), common.HexToAddress("0x0"), EthToWei(1))
          Expect(err).ToNot(HaveOccurred())
          Backend.Commit()
          Expect(isGasExhausted(tx, 100000)).To(BeFalse())
          Expect(isSuccessful(tx)).To(BeFalse())
      })
    })


    It("the initial balance of the Float contract should be zero", func() {
      b, e := Backend.BalanceAt(context.Background(), CryptoFloatAddress, nil)
      Expect(e).ToNot(HaveOccurred())
      Expect(b.String()).To(Equal("0"))
    })

    It("the initial balance of Holder contract address should be zero", func() {
      b, e := Backend.BalanceAt(context.Background(), TokenHolderAddress, nil)
      Expect(e).ToNot(HaveOccurred())
      Expect(b.String()).To(Equal("0"))
    })

    It("the initial balance of the Wallet should be zero", func() {
      b, e := Backend.BalanceAt(context.Background(), WalletAddress, nil)
      Expect(e).ToNot(HaveOccurred())
      Expect(b.String()).To(Equal("0"))
    })

    When("the amount sent is 101 ETH", func() {

        BeforeEach(func() {
            RandomAccount.MustTransfer(Backend, WalletAddress, EthToWei(102))
        })

        BeforeEach(func() {
		tx, err := Wallet.LoadTokenCard(Owner.TransactOpts(), common.HexToAddress("0x0"), EthToWei(101))
        Expect(err).ToNot(HaveOccurred())
		Backend.Commit()
		Expect(isSuccessful(tx)).To(BeTrue())
  		})

          It("Should emit a TransferredToTokenHolder event", func() {
    		it, err := Licence.FilterTransferredToTokenHolder(nil)
    		Expect(err).ToNot(HaveOccurred())
    		Expect(it.Next()).To(BeTrue())
    		evt := it.Event
    		Expect(it.Next()).To(BeFalse())
            Expect(evt.From).To(Equal(WalletAddress))
    		Expect(evt.To).To(Equal(TokenHolderAddress))
            Expect(evt.Asset).To(Equal(common.HexToAddress("0x0"))) //represents ETH
            Expect(evt.Amount.String()).To(Equal(EthToWei(1).String()))
      		})

          It("Should emit a TransferredToCryptoFloat event", func() {
    		it, err := Licence.FilterTransferredToCryptoFloat(nil)
    		Expect(err).ToNot(HaveOccurred())
    		Expect(it.Next()).To(BeTrue())
    		evt := it.Event
    		Expect(it.Next()).To(BeFalse())
            Expect(evt.From).To(Equal(WalletAddress))
    		Expect(evt.To).To(Equal(CryptoFloatAddress))
            Expect(evt.Asset).To(Equal(common.HexToAddress("0x0"))) //represents ETH
            Expect(evt.Amount.String()).To(Equal(EthToWei(100).String()))
      		})

        It("Should emit a LoadedTokenCard event", func() {
          it, err := Wallet.FilterLoadedTokenCard(nil)
          Expect(err).ToNot(HaveOccurred())
          Expect(it.Next()).To(BeTrue())
          evt := it.Event
          Expect(it.Next()).To(BeFalse())
          Expect(evt.Asset).To(Equal(common.HexToAddress("0x0"))) //represents ETH
          Expect(evt.Amount.String()).To(Equal(EthToWei(101).String()))
          })

          It("should increase the ETH balance of the holder contract address by 1 ETH", func() {
            b, e := Backend.BalanceAt(context.Background(), TokenHolderAddress, nil)
            Expect(e).ToNot(HaveOccurred())
            Expect(b.String()).To(Equal(EthToWei(1).String()))
          })

          It("should increase the ETH balance of the holder contract address by 100 ETH", func() {
            b, e := Backend.BalanceAt(context.Background(), CryptoFloatAddress, nil)
            Expect(e).ToNot(HaveOccurred())
            Expect(b.String()).To(Equal(EthToWei(100).String()))
          })
    })

})
