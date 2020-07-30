package gas_proxy_test

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"

	"github.com/ethereum/go-ethereum/common"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/tokencard/contracts/v3/pkg/bindings/mocks"
	. "github.com/tokencard/contracts/v3/test/shared"
)

var _ = Describe("ExecuteTransaction", func() {
	When("a controller transaction is executed in the gas proxy without gas tokens", func() {
		When("a no-op transaction is executed", func() {
			It("should be a successful transaction", func() {
				tx, err := GasProxy.ExecuteTransaction(Controller.TransactOpts(), common.Address{}, big.NewInt(0), []byte{})
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})
		})
		When("a wallet privileged operation is executed", func() {
			It("should be a successful transaction", func() {
				walletABI, err := abi.JSON(strings.NewReader(mocks.WalletABI))
				Expect(err).ToNot(HaveOccurred())
				data, err := walletABI.Pack("confirmOperation")
				Expect(err).ToNot(HaveOccurred())
				tx, err := GasProxy.ExecuteTransaction(Controller.TransactOpts(), WalletAddress, big.NewInt(0), data)
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})
		})
		When("a gas burn transaction is executed", func() {
			It("should be a successful transaction", func() {
				gasBurnerABI, err := abi.JSON(strings.NewReader(mocks.GasBurnerABI))
				Expect(err).ToNot(HaveOccurred())
				data, err := gasBurnerABI.Pack("burnGas", big.NewInt(100000))
				Expect(err).ToNot(HaveOccurred())
				tx, err := GasProxy.ExecuteTransaction(Controller.TransactOpts(), GasBurnerAddress, big.NewInt(0), data)
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})
		})
	})

	When("gas tokens are deposited into the gas proxy contract", func() {
		BeforeEach(func() {
			tx, err := GasToken.Transfer(BankAccount.TransactOpts(), GasProxyAddress, big.NewInt(20))
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeTrue())
		})
		It("the gas token balance of the gas proxy should equal to the amount ransferred", func() {
			balance, err := GasToken.BalanceOf(nil, GasProxyAddress)
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(balance.String()).To(Equal("20"))
		})
		When("a no-op transaction is executed", func() {
			It("should be a successful transaction", func() {
				tx, err := GasProxy.ExecuteTransaction(Controller.TransactOpts(), common.Address{}, big.NewInt(0), []byte{})
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})
		})
		When("a wallet privileged operation is executed", func() {
			It("should be a successful transaction", func() {
				walletABI, err := abi.JSON(strings.NewReader(mocks.WalletABI))
				Expect(err).ToNot(HaveOccurred())
				data, err := walletABI.Pack("confirmOperation")
				Expect(err).ToNot(HaveOccurred())
				tx, err := GasProxy.ExecuteTransaction(Controller.TransactOpts(), WalletAddress, big.NewInt(0), data)
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})
		})
		When("a gas burn transaction is executed", func() {
			It("should be a successful transaction", func() {
				gasBurnerABI, err := abi.JSON(strings.NewReader(mocks.GasBurnerABI))
				Expect(err).ToNot(HaveOccurred())
				data, err := gasBurnerABI.Pack("burnGas", big.NewInt(100000))
				Expect(err).ToNot(HaveOccurred())
				tx, err := GasProxy.ExecuteTransaction(Controller.TransactOpts(), GasBurnerAddress, big.NewInt(0), data)
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})
		})
	})
})
