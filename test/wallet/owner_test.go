package wallet_test

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/tokencard/contracts/pkg/bindings"
	. "github.com/tokencard/ethertest"
)

var _ = Describe("ownable", func() {
	Context("When the contract with transferable ownership has been deployed", func() {
		It("should have an owner", func() {
			o, err := w.Owner(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(o).To(Equal(owner.Address()))
		})
		It("should not allow the ownership to be transferred to the 0x0 address", func() {
			tx, err := w.TransferOwnership(owner.TransactOpts(WithGasLimit(60000)), common.HexToAddress("0x0"))
			Expect(err).ToNot(HaveOccurred())
			be.Commit()
			Expect(isSuccessful(tx)).To(BeFalse())
			transferable, err := w.IsTransferable(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(transferable).To(BeTrue())
		})
		It("should allow the ownership to be transferred to any other address", func() {
			tx, err := w.TransferOwnership(owner.TransactOpts(), common.HexToAddress("0x1"))
			Expect(err).ToNot(HaveOccurred())
			be.Commit()
			Expect(isSuccessful(tx)).To(BeTrue())
			transferable, err := w.IsTransferable(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(transferable).To(BeFalse())
		})
	})
	Context("When a contract without transferable ownership has been deployed", func() {
		BeforeEach(func() {
			var err error
			var tx *types.Transaction
			wa, tx, w, err = bindings.DeployWallet(
				bankAccount.TransactOpts(),
				be,
				owner.Address(),
				false,
				ensAddress,
				oracleName,
				controllerName,
				ethToWei(100),
			)
			Expect(err).ToNot(HaveOccurred())
			be.Commit()
			Expect(isSuccessful(tx)).To(BeTrue())
		})
		It("should have an owner", func() {
			o, err := w.Owner(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(o).To(Equal(owner.Address()))
		})
		It("should not allow the ownership to be transferred", func() {
			tx, err := w.TransferOwnership(owner.TransactOpts(WithGasLimit(60000)), common.HexToAddress("0x1"))
			Expect(err).ToNot(HaveOccurred())
			be.Commit()
			Expect(isSuccessful(tx)).To(BeFalse())
			transferable, err := w.IsTransferable(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(transferable).To(BeFalse())
		})
	})
})
