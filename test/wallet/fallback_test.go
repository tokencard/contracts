package wallet_test

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tokencard/contracts/test/shared"
)

var _ = Describe("fallback", func() {
	Context("When a random person deposits ETH to the wallet address", func() {
		BeforeEach(func() {
			RandomAccount.MustTransfer(Backend, WalletAddress, FinneyToWei(1))
		})

		It("should increase the wallet's ETH balance by the same amount", func() {
			balance, err := Backend.BalanceAt(context.Background(), WalletAddress, nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(balance.String()).To(Equal(FinneyToWei(1).String()))
		})

		It("should emit a deposit event", func() {
			it, err := Wallet.FilterReceived(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(it.Next()).To(BeTrue())
			evt := it.Event
			Expect(it.Next()).To(BeFalse())
			Expect(evt.From).To(Equal(RandomAccount.Address()))
			Expect(evt.Amount.String()).To(Equal(FinneyToWei(1).String()))
		})
	})

	When(" a random person calls a wallet method that doesn't exist", func() {
		var randomAddress common.Address
		BeforeEach(func() {
			privateKey, err := crypto.GenerateKey()
			randomAddress = crypto.PubkeyToAddress(privateKey.PublicKey)
			Expect(err).ToNot(HaveOccurred())
			BankAccount.MustTransfer(Backend, randomAddress, EthToWei(1))
			unsignedTx := types.NewTransaction(0, WalletAddress, FinneyToWei(1), 30000, GweiToWei(20), []byte{0xf8, 0xa8, 0xfd, 0xd6})
			signedTx, err := types.SignTx(unsignedTx, types.HomesteadSigner{}, privateKey)
			Expect(err).ToNot(HaveOccurred())
			err = Backend.SendTransaction(context.Background(), signedTx)
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(signedTx)).To(BeTrue())
		})

		It("should increase the wallet's ETH balance by one Finney", func() {
			balance, err := Backend.BalanceAt(context.Background(), WalletAddress, nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(balance.String()).To(Equal(FinneyToWei(1).String()))
		})

		It("should emit a deposit event", func() {
			it, err := Wallet.FilterReceived(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(it.Next()).To(BeTrue())
			evt := it.Event
			Expect(it.Next()).To(BeFalse())
			Expect(evt.From).To(Equal(randomAddress))
			Expect(evt.Amount.String()).To(Equal(FinneyToWei(1).String()))
		})
	})
})
