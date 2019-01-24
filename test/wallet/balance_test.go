package wallet_test

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/tokencard/contracts/pkg/bindings/mocks"
	. "github.com/tokencard/contracts/test/shared"
)

var _ = Describe("balance", func() {
	Context("When the contract has no balance", func() {
		It("should return 0", func() {
			b, err := Wallet.Balance(nil, common.HexToAddress("0x0"))
			Expect(err).ToNot(HaveOccurred())
			Expect(b.String()).To(Equal("0"))
		})
	})

	FContext("When contract has 1 ETH", func() {
		BeforeEach(func() {
			BankAccount.MustTransfer(Backend, WalletAddress, EthToWei(1))
		})

		It("should return 1 ETH", func() {
			b, err := Wallet.Balance(nil, common.HexToAddress("0x0"))
			Expect(err).ToNot(HaveOccurred())
			Expect(b.String()).To(Equal(EthToWei(1).String()))
		})
	})

	Context("When contract owns one TKN", func() {
		var t *mocks.Token
		var ta common.Address
		BeforeEach(func() {
			var err error
			ta, _, t, err = mocks.DeployToken(BankAccount.TransactOpts(), Backend)
			Expect(err).ToNot(HaveOccurred())

			tx, err := t.Credit(BankAccount.TransactOpts(), WalletAddress, big.NewInt(1))

			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()

			r, err := Backend.TransactionReceipt(context.Background(), tx.Hash())
			Expect(err).ToNot(HaveOccurred())
			Expect(r.Status).To(Equal(types.ReceiptStatusSuccessful))

		})

		It("Balance for that token should return 1", func() {
			b, err := Wallet.Balance(nil, ta)
			Expect(err).ToNot(HaveOccurred())
			Expect(b.String()).To(Equal("1"))
		})

	})

})
