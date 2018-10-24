package wallet_test

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/tokencard/contracts/pkg/bindings"
	. "github.com/tokencard/contracts/test/shared"
	"github.com/tokencard/ethertest"
)

var _ = Describe("ownable", func() {
	Context("When the contract with transferable ownership has been deployed", func() {
		It("should have an owner", func() {
			o, err := Wallet.Owner(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(o).To(Equal(Owner.Address()))
		})
		It("should not allow the ownership to Backend transferred to the 0x0 address", func() {
			tx, err := Wallet.TransferOwnership(Owner.TransactOpts(ethertest.WithGasLimit(60000)), common.HexToAddress("0x0"))
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeFalse())
			transferable, err := Wallet.IsTransferable(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(transferable).To(BeTrue())
		})
		It("should allow the ownership to Backend transferred to any other address", func() {
			tx, err := Wallet.TransferOwnership(Owner.TransactOpts(), common.HexToAddress("0x1"))
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeTrue())
			transferable, err := Wallet.IsTransferable(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(transferable).To(BeFalse())
		})
	})
	Context("When a contract without transferable ownership has been deployed", func() {
		BeforeEach(func() {
			var err error
			var tx *types.Transaction
			WalletAddress, tx, Wallet, err = bindings.DeployWallet(
				BankAccount.TransactOpts(),
				Backend,
				Owner.Address(),
				false,
				ENSRegistryAddress,
				OracleName,
				ControllerName,
				EthToWei(100),
			)
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeTrue())
		})
		It("should have an owner", func() {
			o, err := Wallet.Owner(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(o).To(Equal(Owner.Address()))
		})
		It("should not allow the ownership to Backend transferred", func() {
			tx, err := Wallet.TransferOwnership(Owner.TransactOpts(ethertest.WithGasLimit(60000)), common.HexToAddress("0x1"))
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeFalse())
			transferable, err := Wallet.IsTransferable(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(transferable).To(BeFalse())
		})
	})
})
