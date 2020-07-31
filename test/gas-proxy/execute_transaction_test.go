package gas_proxy_test

import (
	"context"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/ethereum/go-ethereum/common"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/tokencard/contracts/v3/pkg/bindings/mocks"
	. "github.com/tokencard/contracts/v3/test/shared"
)

var _ = Describe("ExecuteTransaction", func() {
	When("a controller transaction is executed in the context of the gas proxy with no gas tokens deposited", func() {
		When("a no-op transaction is executed", func() {
			var tx *types.Transaction
			BeforeEach(func() {
				err := error(nil)
				tx, err = GasProxy.ExecuteTransaction(Controller.TransactOpts(), common.Address{}, big.NewInt(0), []byte{})
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
			})
			It("should be a successful transaction", func() {
				Expect(isSuccessful(tx)).To(BeTrue())
			})
			It("should use the expected amout of gas", func() {
				receipt, err := Backend.TransactionReceipt(context.Background(), tx.Hash())
				Expect(err).ToNot(HaveOccurred())
				fmt.Fprintf(GinkgoWriter, "Gas used by no-op execute transaction: %d\n", receipt.GasUsed)
				Expect(receipt.GasUsed > 45000 && receipt.GasUsed < 60000).To(BeTrue())
			})
		})
		When("a wallet privileged operation is executed", func() {
			var tx *types.Transaction
			BeforeEach(func() {
				walletABI, err := abi.JSON(strings.NewReader(mocks.WalletABI))
				Expect(err).ToNot(HaveOccurred())
				data, err := walletABI.Pack("confirmOperation")
				Expect(err).ToNot(HaveOccurred())
				tx, err = GasProxy.ExecuteTransaction(Controller.TransactOpts(), WalletAddress, big.NewInt(0), data)
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
			})
			It("should be a successful transaction", func() {
				Expect(isSuccessful(tx)).To(BeTrue())
			})
			It("should use the expected amout of gas", func() {
				receipt, err := Backend.TransactionReceipt(context.Background(), tx.Hash())
				Expect(err).ToNot(HaveOccurred())
				fmt.Fprintf(GinkgoWriter, "Gas used by confirmOperation execute transaction: %d\n", receipt.GasUsed)
				Expect(receipt.GasUsed > 60000 && receipt.GasUsed < 75000).To(BeTrue())
			})
		})
		When("a gas burn transaction is executed", func() {
			var tx *types.Transaction
			BeforeEach(func() {
				gasBurnerABI, err := abi.JSON(strings.NewReader(mocks.GasBurnerABI))
				Expect(err).ToNot(HaveOccurred())
				data, err := gasBurnerABI.Pack("burnGas", big.NewInt(1000000))
				Expect(err).ToNot(HaveOccurred())
				tx, err = GasProxy.ExecuteTransaction(Controller.TransactOpts(), GasBurnerAddress, big.NewInt(0), data)
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
			})
			It("should be a successful transaction", func() {
				Expect(isSuccessful(tx)).To(BeTrue())
			})
			It("should use the expected amout of gas", func() {
				receipt, err := Backend.TransactionReceipt(context.Background(), tx.Hash())
				Expect(err).ToNot(HaveOccurred())
				fmt.Fprintf(GinkgoWriter, "Gas used by burnGas execute transaction: %d\n", receipt.GasUsed)
				Expect(receipt.GasUsed > 100000 && receipt.GasUsed < 1150000).To(BeTrue())
			})
		})
	})

	When("a controller transaction is executed in the context of the gas proxy with gas tokens deposited", func() {
		BeforeEach(func() {
			tx, err := GasToken.Transfer(BankAccount.TransactOpts(), GasProxyAddress, big.NewInt(140))
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeTrue())
		})
		It("the gas token balance of the gas proxy should equal to the expected amount", func() {
			balance, err := GasToken.BalanceOf(nil, GasProxyAddress)
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(balance.String()).To(Equal("140"))
		})
		When("a no-op transaction is executed", func() {
			var tx *types.Transaction
			var totalBurned *big.Int
			BeforeEach(func() {
				err := error(nil)
				totalBurned, err = GasToken.TotalBurned(nil)
				Expect(err).ToNot(HaveOccurred())
				tx, err = GasProxy.ExecuteTransaction(Controller.TransactOpts(), common.Address{}, big.NewInt(0), []byte{})
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
			})
			It("should be a successful transaction", func() {
				Expect(isSuccessful(tx)).To(BeTrue())
			})
			It("should use less gas due to the gas refund", func() {
				receipt, err := Backend.TransactionReceipt(context.Background(), tx.Hash())
				Expect(err).ToNot(HaveOccurred())
				fmt.Fprintf(GinkgoWriter, "Gas used by no-op execute transaction (with gas refund): %d\n", receipt.GasUsed)
				Expect(receipt.GasUsed > 20000 && receipt.GasUsed < 60000).To(BeTrue())
			})
			It("should burn the expected amount of gas tokens", func() {
				newTotalBurned, err := GasToken.TotalBurned(nil)
				Expect(err).ToNot(HaveOccurred())
				amountBurned := new(big.Int).Sub(newTotalBurned, totalBurned)
				fmt.Fprintf(GinkgoWriter, "Gas tokens burned by no-op execute transaction: %s\n", amountBurned.String())
				Expect(amountBurned.String()).To(Equal("1"))
			})
		})
		When("a wallet privileged operation is executed", func() {
			var tx *types.Transaction
			var totalBurned *big.Int
			BeforeEach(func() {
				err := error(nil)
				totalBurned, err = GasToken.TotalBurned(nil)
				Expect(err).ToNot(HaveOccurred())
				walletABI, err := abi.JSON(strings.NewReader(mocks.WalletABI))
				Expect(err).ToNot(HaveOccurred())
				data, err := walletABI.Pack("confirmOperation")
				Expect(err).ToNot(HaveOccurred())
				tx, err = GasProxy.ExecuteTransaction(Controller.TransactOpts(), WalletAddress, big.NewInt(0), data)
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
			})
			It("should be a successful transaction", func() {
				Expect(isSuccessful(tx)).To(BeTrue())
			})
			It("should use less gas due to the gas refund", func() {
				receipt, err := Backend.TransactionReceipt(context.Background(), tx.Hash())
				Expect(err).ToNot(HaveOccurred())
				fmt.Fprintf(GinkgoWriter, "Gas used by confirmOperation execute transaction (with gas refund): %d\n", receipt.GasUsed)
				Expect(receipt.GasUsed > 30000 && receipt.GasUsed < 75000).To(BeTrue())
			})
			It("should burn the expected amount of gas tokens", func() {
				newTotalBurned, err := GasToken.TotalBurned(nil)
				Expect(err).ToNot(HaveOccurred())
				amountBurned := new(big.Int).Sub(newTotalBurned, totalBurned)
				fmt.Fprintf(GinkgoWriter, "Gas tokens burned by confirmOperation execute transaction: %s\n", amountBurned.String())
				Expect(amountBurned.String()).To(Equal("1"))
			})
		})
		When("a gas burn transaction is executed", func() {
			var tx *types.Transaction
			var totalBurned *big.Int
			BeforeEach(func() {
				err := error(nil)
				totalBurned, err = GasToken.TotalBurned(nil)
				Expect(err).ToNot(HaveOccurred())
				gasBurnerABI, err := abi.JSON(strings.NewReader(mocks.GasBurnerABI))
				Expect(err).ToNot(HaveOccurred())
				data, err := gasBurnerABI.Pack("burnGas", big.NewInt(1000000))
				Expect(err).ToNot(HaveOccurred())
				tx, err = GasProxy.ExecuteTransaction(Controller.TransactOpts(), GasBurnerAddress, big.NewInt(0), data)
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
			})
			It("should be a successful transaction", func() {
				Expect(isSuccessful(tx)).To(BeTrue())
			})
			It("should use less gas due to the gas refund", func() {
				receipt, err := Backend.TransactionReceipt(context.Background(), tx.Hash())
				Expect(err).ToNot(HaveOccurred())
				fmt.Fprintf(GinkgoWriter, "Gas used by burnGas execute transaction (with gas refund): %d\n", receipt.GasUsed)
				Expect(receipt.GasUsed > 100000 && receipt.GasUsed < 1150000).To(BeTrue())
			})
			It("should burn the expected amount of gas tokens", func() {
				newTotalBurned, err := GasToken.TotalBurned(nil)
				Expect(err).ToNot(HaveOccurred())
				amountBurned := new(big.Int).Sub(newTotalBurned, totalBurned)
				fmt.Fprintf(GinkgoWriter, "Gas tokens burned by burnGas execute transaction: %s\n", amountBurned.String())
				Expect(amountBurned.String()).To(Equal("26"))
			})
		})
	})
})
