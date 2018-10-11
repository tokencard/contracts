package wallet_test

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("fallback", func() {
	Context("When a random person deposits ETH to the wallet address", func() {
		BeforeEach(func() {
			randomPerson.MustTransfer(be, wa, ONE_FINNEY)
		})

		It("should increase the wallet's ETH balance by the same amount", func() {
			balance, err := w.Balance(nil, common.HexToAddress("0x0"))
			Expect(err).ToNot(HaveOccurred())
			Expect(balance.String()).To(Equal(ONE_FINNEY.String()))
		})

		It("should emit a deposit event", func() {
			it, err := w.FilterReceived(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(it.Next()).To(BeTrue())
			evt := it.Event
			Expect(it.Next()).To(BeFalse())
			Expect(evt.From).To(Equal(randomPerson.Address()))
			Expect(evt.Amount.String()).To(Equal(ONE_FINNEY.String()))
		})
	})

	Context("When a random person calls a wallet method that doesn't exist", func() {
		It("should fail", func() {
			privateKey, err := crypto.GenerateKey()
			Expect(err).ToNot(HaveOccurred())
			bankWallet.MustTransfer(be, crypto.PubkeyToAddress(privateKey.PublicKey), ONE_ETH)
			unsignedTx := types.NewTransaction(0, wa, ONE_FINNEY, 30000, gweiToWei(20), []byte{0xf8, 0xa8, 0xfd, 0xd6})
			signedTx, err := types.SignTx(unsignedTx, types.HomesteadSigner{}, privateKey)
			Expect(err).ToNot(HaveOccurred())
			err = be.SendTransaction(context.Background(), signedTx)
			Expect(err).ToNot(HaveOccurred())
			be.Commit()
			Expect(isSuccessful(signedTx)).To(BeFalse())
		})
	})
})
