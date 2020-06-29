package wallet_test

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/tokencard/contracts/v3/pkg/bindings/mocks"
	. "github.com/tokencard/contracts/v3/test/shared"
)

var _ = Describe("balance", func() {
	Context("When the contract has no balance", func() {
		It("should return 0", func() {
			b, e := Backend.BalanceAt(context.Background(), WalletProxyAddress, nil)
			Expect(e).ToNot(HaveOccurred())
			Expect(b.String()).To(Equal("0"))
		})
	})

	Context("When contract has 1 ETH", func() {
		BeforeEach(func() {
			BankAccount.MustTransfer(Backend, WalletProxyAddress, EthToWei(1))
		})

		It("should return 1 ETH", func() {
			b, e := Backend.BalanceAt(context.Background(), WalletProxyAddress, nil)
			Expect(e).ToNot(HaveOccurred())
			Expect(b.String()).To(Equal(EthToWei(1).String()))
		})
	})

	Context("When contract owns one TKN", func() {
		var t *mocks.Token
		BeforeEach(func() {
			var err error
			_, _, t, err = mocks.DeployToken(BankAccount.TransactOpts(), Backend)
			Expect(err).ToNot(HaveOccurred())

			tx, err := t.Credit(BankAccount.TransactOpts(), WalletProxyAddress, big.NewInt(1))

			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()

			r, err := Backend.TransactionReceipt(context.Background(), tx.Hash())
			Expect(err).ToNot(HaveOccurred())
			Expect(r.Status).To(Equal(types.ReceiptStatusSuccessful))

		})

		It("Balance for that token should return 1", func() {
			b, err := t.BalanceOf(nil, WalletProxyAddress)
			Expect(err).ToNot(HaveOccurred())
			Expect(b.String()).To(Equal("1"))
		})

	})

})
