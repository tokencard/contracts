package wallet_test

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tokencard/contracts/test/shared"
	"github.com/tokencard/ethertest"
)

var _ = Describe("bulk_transfer", func() {

	Context("when the wallet has enough ETH", func() {
		BeforeEach(func() {
			BankAccount.MustTransfer(Backend, WalletAddress, EthToWei(1))
			BankAccount.MustTransfer(Backend, Controller.Address(), EthToWei(1))
		})

		var tx *types.Transaction

		Context("When the size of the array of assets don't match the size of the array of amounts", func() {
			BeforeEach(func() {
				var err error
				tx, err = Wallet.BulkTransfer(Owner.TransactOpts(ethertest.WithGasLimit(81000)), RandomAccount.Address(), []common.Address{common.HexToAddress("0x0"), common.HexToAddress("0x2"), common.HexToAddress("0x3")}, []*big.Int{big.NewInt(1), big.NewInt(2)})
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
			})
			It("should fail", func() {
				Expect(isSuccessful(tx)).To(BeFalse())
			})
		})

		Context("When the size of the array of assets does match the size of the array of amounts", func() {
			BeforeEach(func() {
				var err error
				tx, err = Wallet.BulkTransfer(Owner.TransactOpts(ethertest.WithGasLimit(81000)), RandomAccount.Address(), []common.Address{common.HexToAddress("0x0")}, []*big.Int{big.NewInt(1)})
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
			})
			It("should pass", func() {
				Expect(isSuccessful(tx)).To(BeTrue())
			})
		})
	})
})
