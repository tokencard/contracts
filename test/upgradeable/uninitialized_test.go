package upgrade_test

import (
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/tokencard/contracts/v3/pkg/bindings"
	"github.com/tokencard/contracts/v3/pkg/bindings/externals/upgradeability"
	. "github.com/tokencard/contracts/v3/test/shared"
	"github.com/tokencard/ethertest"
)

var _ = Describe("uninitialized", func() {

	When("deploying a new proxy without initializing it", func() {
		var tx *types.Transaction
		var err error
		BeforeEach(func() {
			ProxyAddress, tx, Proxy, err = upgradeability.DeployUpgradeabilityProxy(BankAccount.TransactOpts(), Backend, WalletImplementationAddress, nil)
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeTrue())
			ProxyWallet, err = bindings.NewWallet(ProxyAddress, Backend)
		})

		It("Should fail when there's onlyOwnerOr2FA()", func() {
			tx, err = ProxyWallet.TopUpGas(Owner.TransactOpts(ethertest.WithGasLimit(100000)), EthToWei(100))
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeFalse())
			returnData, _ := ethCall(tx)
			Expect(string(returnData[len(returnData)-64:])).To(ContainSubstring("only owner or 2FA"))
		})

		It("Should fail it when there's onlyOwnerOrSelf()", func() {
			tx, err = ProxyWallet.SubmitDailyLimitUpdate(Owner.TransactOpts(ethertest.WithGasLimit(100000)), EthToWei(1000))
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeFalse())
			returnData, _ := ethCall(tx)
			Expect(string(returnData[len(returnData)-64:])).To(ContainSubstring("Not owner or self"))
		})

		It("Should fail it when there's only2FA()", func() {
			tx, err = ProxyWallet.ConfirmDailyLimitUpdate(Owner.TransactOpts(ethertest.WithGasLimit(100000)), big.NewInt(1000))
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeFalse())
			returnData, _ := ethCall(tx)
			Expect(string(returnData[len(returnData)-64:])).To(ContainSubstring("sender is not personal 2FA account"))
		})

		It("Should fail it when there's onlyOwner()", func() {
			tx, err = ProxyWallet.IncreaseRelayNonce(Owner.TransactOpts(ethertest.WithGasLimit(100000)))
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeFalse())
			returnData, _ := ethCall(tx)
			Expect(string(returnData[len(returnData)-64:])).To(ContainSubstring("sender is not owner"))
		})

	})

})
