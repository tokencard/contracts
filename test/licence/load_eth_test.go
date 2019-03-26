package licence_test

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tokencard/contracts/test/shared"
	"github.com/tokencard/ethertest"
)

var _ = Describe("load ETH", func() {

	When("no value is sent 0", func() {

		It("Should revert", func() {
			tx, err := Licence.Load(RandomAccount.TransactOpts(ethertest.WithGasLimit(100000)), common.HexToAddress("0x0"), big.NewInt(1000))
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

	It("the initial balance of Holder address should be zero", func() {
		b, e := Backend.BalanceAt(context.Background(), TokenHolderAddress, nil)
		Expect(e).ToNot(HaveOccurred())
		Expect(b.String()).To(Equal("0"))
	})

	When("the amount sent is 101 ETH", func() {

		BeforeEach(func() {
			tx, err := Licence.Load(RandomAccount.TransactOpts(ethertest.WithValue(EthToWei(101))), common.HexToAddress("0x0"), EthToWei(101))
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isGasExhausted(tx, 100000)).To(BeFalse())
			Expect(isSuccessful(tx)).To(BeTrue())
		})

		It("Should emit a TransferredToTokenHolder event", func() {
			it, err := Licence.FilterTransferredToTokenHolder(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(it.Next()).To(BeTrue())
			evt := it.Event
			Expect(it.Next()).To(BeFalse())
			Expect(evt.From).To(Equal(RandomAccount.Address()))
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
			Expect(evt.From).To(Equal(RandomAccount.Address()))
			Expect(evt.To).To(Equal(CryptoFloatAddress))
			Expect(evt.Asset).To(Equal(common.HexToAddress("0x0"))) //represents ETH
			Expect(evt.Amount.String()).To(Equal(EthToWei(100).String()))
		})

		It("should increase the ETH balance of the holder contract address by 1 ETH", func() {
			b, e := Backend.BalanceAt(context.Background(), TokenHolderAddress, nil)
			Expect(e).ToNot(HaveOccurred())
			Expect(b.String()).To(Equal(EthToWei(1).String()))
		})

		It("should increase the ETH balance of the Float contract address by 100 ETH", func() {
			b, e := Backend.BalanceAt(context.Background(), CryptoFloatAddress, nil)
			Expect(e).ToNot(HaveOccurred())
			Expect(b.String()).To(Equal(EthToWei(100).String()))
		})

	})

	When("the amount sent is 101 ETH + 1 wei", func() {

		var amount *big.Int

		BeforeEach(func() {
			v := big.NewInt(1) //value
			a := big.NewInt(1) //amount
			tx, err := Licence.Load(RandomAccount.TransactOpts(ethertest.WithValue(v.Add(v, EthToWei(101)))), common.HexToAddress("0x0"), a.Add(a, EthToWei(101)))
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isGasExhausted(tx, 100000)).To(BeFalse())
			Expect(isSuccessful(tx)).To(BeTrue())
		})

		It("Should emit a TransferredToTokenHolder event", func() {
			it, err := Licence.FilterTransferredToTokenHolder(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(it.Next()).To(BeTrue())
			evt := it.Event
			Expect(it.Next()).To(BeFalse())
			Expect(evt.From).To(Equal(RandomAccount.Address()))
			Expect(evt.To).To(Equal(TokenHolderAddress))
			Expect(evt.Asset).To(Equal(common.HexToAddress("0x0"))) //represents ETH
			amount = EthToWei(1)
			amount.Add(amount, big.NewInt(1))
			Expect(evt.Amount.String()).To(Equal(amount.String()))
		})

		It("Should emit a TransferredToCryptoFloat event", func() {
			it, err := Licence.FilterTransferredToCryptoFloat(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(it.Next()).To(BeTrue())
			evt := it.Event
			Expect(it.Next()).To(BeFalse())
			Expect(evt.From).To(Equal(RandomAccount.Address()))
			Expect(evt.To).To(Equal(CryptoFloatAddress))
			Expect(evt.Asset).To(Equal(common.HexToAddress("0x0"))) //represents ETH
			Expect(evt.Amount.String()).To(Equal(EthToWei(100).String()))
		})

		It("should increase the ETH balance of the holder contract address by 1 ETH + 1 wei", func() {
			b, e := Backend.BalanceAt(context.Background(), TokenHolderAddress, nil)
			Expect(e).ToNot(HaveOccurred())
			Expect(b.String()).To(Equal(amount.String()))
		})

		It("should increase the ETH balance of the holder contract address by 100 ETH", func() {
			b, e := Backend.BalanceAt(context.Background(), CryptoFloatAddress, nil)
			Expect(e).ToNot(HaveOccurred())
			Expect(b.String()).To(Equal(EthToWei(100).String()))
		})

	})

	When("the amount sent is 2 wei", func() {

		BeforeEach(func() {
			tx, err := Licence.Load(RandomAccount.TransactOpts(ethertest.WithValue(big.NewInt(2))), common.HexToAddress("0x0"), big.NewInt(2))
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isGasExhausted(tx, 100000)).To(BeFalse())
			Expect(isSuccessful(tx)).To(BeTrue())
		})

		It("Should emit a TransferredToTokenHolder event", func() {
			it, err := Licence.FilterTransferredToTokenHolder(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(it.Next()).To(BeTrue())
			evt := it.Event
			Expect(it.Next()).To(BeFalse())
			Expect(evt.From).To(Equal(RandomAccount.Address()))
			Expect(evt.To).To(Equal(TokenHolderAddress))
			Expect(evt.Asset).To(Equal(common.HexToAddress("0x0"))) //represents ETH
			Expect(evt.Amount.String()).To(Equal(big.NewInt(1).String()))
		})

		It("Should emit a TransferredToCryptoFloat event", func() {
			it, err := Licence.FilterTransferredToCryptoFloat(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(it.Next()).To(BeTrue())
			evt := it.Event
			Expect(it.Next()).To(BeFalse())
			Expect(evt.From).To(Equal(RandomAccount.Address()))
			Expect(evt.To).To(Equal(CryptoFloatAddress))
			Expect(evt.Asset).To(Equal(common.HexToAddress("0x0"))) //represents ETH
			Expect(evt.Amount.String()).To(Equal(big.NewInt(1).String()))
		})

		It("should increase the ETH balance of the holder contract address by 1 wei", func() {
			b, e := Backend.BalanceAt(context.Background(), TokenHolderAddress, nil)
			Expect(e).ToNot(HaveOccurred())
			Expect(b.String()).To(Equal("1"))
		})

		It("should increase the ETH balance of the holder contract address by 1 wei", func() {
			b, e := Backend.BalanceAt(context.Background(), CryptoFloatAddress, nil)
			Expect(e).ToNot(HaveOccurred())
			Expect(b.String()).To(Equal("1"))
		})
	})

	When("the amount sent is 101 ETH + 9 wei", func() {

		var amount, transferAmount *big.Int

		BeforeEach(func() {
			v := big.NewInt(9) //value
			a := big.NewInt(9) //amount
			tx, err := Licence.Load(RandomAccount.TransactOpts(ethertest.WithValue(v.Add(v, EthToWei(101)))), common.HexToAddress("0x0"), a.Add(a, EthToWei(101)))
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isGasExhausted(tx, 100000)).To(BeFalse())
			Expect(isSuccessful(tx)).To(BeTrue())
		})

		It("Should emit a TransferredToTokenHolder event", func() {
			it, err := Licence.FilterTransferredToTokenHolder(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(it.Next()).To(BeTrue())
			evt := it.Event
			Expect(it.Next()).To(BeFalse())
			Expect(evt.From).To(Equal(RandomAccount.Address()))
			Expect(evt.To).To(Equal(TokenHolderAddress))
			Expect(evt.Asset).To(Equal(common.HexToAddress("0x0"))) //represents ETH
			amount = EthToWei(1)
			amount.Add(amount, big.NewInt(1))
			Expect(evt.Amount.String()).To(Equal(amount.String()))
		})

		It("Should emit a TransferredToCryptoFloat event", func() {
			it, err := Licence.FilterTransferredToCryptoFloat(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(it.Next()).To(BeTrue())
			evt := it.Event
			Expect(it.Next()).To(BeFalse())
			Expect(evt.From).To(Equal(RandomAccount.Address()))
			Expect(evt.To).To(Equal(CryptoFloatAddress))
			Expect(evt.Asset).To(Equal(common.HexToAddress("0x0"))) //represents ETH
			transferAmount = EthToWei(100)
			transferAmount.Add(transferAmount, big.NewInt(8))
			Expect(evt.Amount.String()).To(Equal(transferAmount.String()))
		})

		It("should increase the ETH balance of the holder contract address by 1 ETH + 1 wei", func() {
			b, e := Backend.BalanceAt(context.Background(), TokenHolderAddress, nil)
			Expect(e).ToNot(HaveOccurred())
			Expect(b.String()).To(Equal(amount.String()))
		})

		It("should increase the ETH balance of the holder contract address by 100 ETH + 8 wei", func() {
			b, e := Backend.BalanceAt(context.Background(), CryptoFloatAddress, nil)
			Expect(e).ToNot(HaveOccurred())
			Expect(b.String()).To(Equal(transferAmount.String()))
		})

	})

})
