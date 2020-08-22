package wallet_test

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/tokencard/contracts/v3/pkg/bindings"
	"github.com/tokencard/contracts/v3/pkg/bindings/externals/upgradeability"
	. "github.com/tokencard/contracts/v3/test/shared"
	"github.com/tokencard/ethertest"
)

var _ = Describe("ownable", func() {

	When("the contract with transferable ownership has been deployed", func() {

		It("should have an owner", func() {
			o, err := WalletProxy.Owner(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(o).To(Equal(Owner.Address()))
		})

		It("should not allow the ownership to Backend transferred to the 0x0 address", func() {
			tx, err := WalletProxy.TransferOwnership(Owner.TransactOpts(ethertest.WithGasLimit(60000)), common.HexToAddress("0x0"), true)
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeFalse())
			transferable, err := WalletProxy.IsTransferable(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(transferable).To(BeTrue())
		})

		It("should allow the ownership to Backend transferred to any other address", func() {
			tx, err := WalletProxy.TransferOwnership(Owner.TransactOpts(), common.HexToAddress("0x1"), false)
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeTrue())
			transferable, err := WalletProxy.IsTransferable(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(transferable).To(BeFalse())
		})
	})
	When("a contract without transferable ownership has been deployed", func() {

		var RandomProxy *bindings.Wallet
		var RandomProxyAddress common.Address
		var tx *types.Transaction
		var err error

		BeforeEach(func() {
			RandomProxyAddress, tx, _, err = upgradeability.DeployUpgradeabilityProxy(BankAccount.TransactOpts(), Backend, WalletImplementationAddress, nil)
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeTrue())

			RandomProxy, err = bindings.NewWallet(RandomProxyAddress, Backend)
			tx, err = RandomProxy.InitializeWallet(BankAccount.TransactOpts(), Owner.Address(), false, ENSRegistryAddress, TokenWhitelistNode, ControllerNode, LicenceNode, EthToWei(100))
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeTrue())
		})

		It("should have an owner and be non-transferable", func() {
			o, err := RandomProxy.Owner(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(o).To(Equal(Owner.Address()))
			transferable, err := RandomProxy.IsTransferable(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(transferable).To(BeFalse())
		})

		It("should not allow the ownership to be transferred", func() {
			tx, err := RandomProxy.TransferOwnership(Owner.TransactOpts(ethertest.WithGasLimit(70000)), common.HexToAddress("0x1"), false)
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeFalse())
			returnData, _ := ethCall(tx)
			Expect(string(returnData[len(returnData)-64:])).To(ContainSubstring("ownership is not transferable"))
		})
	})
})
