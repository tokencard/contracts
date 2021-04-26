package upgrade_test

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/tokencard/contracts/v3/pkg/bindings/mocks"
	. "github.com/tokencard/contracts/v3/test/shared"
)

var _ = Describe("fallback", func() {

	var Wallet *mocks.Wallet
	var WalletAddress common.Address

	When("trying transfer funds", func() {
		var tx *types.Transaction
		var err error
		BeforeEach(func() {
			WalletAddress, tx, Wallet, err = mocks.DeployWallet(Owner.TransactOpts(), Backend, ENSRegistryAddress, ControllerNode)
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeTrue())
			// deposit 1 wei
			BankAccount.MustTransfer(Backend, WalletAddress, big.NewInt(1))
		})

		When("using sendValue() aka call()", func() {

			BeforeEach(func() {
				// transfer 1 wei to proxy
				tx, err = Wallet.SendValue(Owner.TransactOpts(), ProxyAddress, big.NewInt(1))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})

			It("should return 1 wei", func() {
				b, e := Backend.BalanceAt(context.Background(), ProxyAddress, nil)
				Expect(e).ToNot(HaveOccurred())
				Expect(b.String()).To(Equal("1"))
			})
		})

		When("using transfer()", func() {
			BeforeEach(func() {
				// transfer 1 wei to proxy
				tx, err = Wallet.Transfer(Owner.TransactOpts(), ProxyAddress, big.NewInt(1))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})
			It("should return 1 wei", func() {
				b, e := Backend.BalanceAt(context.Background(), ProxyAddress, nil)
				Expect(e).ToNot(HaveOccurred())
				Expect(b.String()).To(Equal("1"))
			})
			It("should emit a deposit event", func() {
				it, err := Proxy.FilterReceived(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(it.Next()).To(BeTrue())
				evt := it.Event
				Expect(it.Next()).To(BeFalse())
				Expect(evt.From).To(Equal(WalletAddress))
				Expect(evt.Amount.String()).To(Equal("1"))
			})
		})

	})

})
