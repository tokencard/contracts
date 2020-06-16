package upgrade_test

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/Masterminds/semver"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/tokencard/contracts/v3/pkg/bindings"
	"github.com/tokencard/contracts/v3/pkg/bindings/externals/upgradeability"
	. "github.com/tokencard/contracts/v3/test/shared"
)

var Wallet *bindings.Wallet
var WalletAddress common.Address

var Proxy *upgradeability.AdminUpgradeabilityProxy
var ProxyAddress common.Address

var ProxyWallet *bindings.Wallet

func init() {
	TestRig.AddCoverageForContracts("../../build/wallet/combined.json", "../../contracts")
}

func TestWalletSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Upgradeable Suite")
}

var _ = BeforeEach(func() {
	err := InitializeBackend()
	Expect(err).ToNot(HaveOccurred())
	// Deploy the Token wallet contract.
	var tx *types.Transaction
	WalletAddress, tx, Wallet, err = bindings.DeployWallet(BankAccount.TransactOpts(), Backend)
	Expect(err).ToNot(HaveOccurred())
	Backend.Commit()
	Expect(isSuccessful(tx)).To(BeTrue())

	ProxyAddress, tx, Proxy, err = upgradeability.DeployAdminUpgradeabilityProxy(Owner.TransactOpts(), Backend, WalletAddress, Owner.Address(), nil)
	Expect(err).ToNot(HaveOccurred())
	Backend.Commit()
	Expect(isSuccessful(tx)).To(BeTrue())

	ProxyWallet, err = bindings.NewWallet(ProxyAddress, Backend)
	tx, err = ProxyWallet.InitializeWallet(Owner.TransactOpts(), Owner.Address(), true, ENSRegistryAddress, TokenWhitelistName, ControllerName, LicenceName, EthToWei(1))
	Expect(err).ToNot(HaveOccurred())
	Backend.Commit()
	Expect(isSuccessful(tx)).To(BeTrue())
})

var allPassed = true
var currentVersion = "3.2.0"

var _ = Describe("Wallet Version", func() {
	It("should return the current version", func() {
		v, err := ProxyWallet.WALLETVERSION(nil)
		Expect(err).ToNot(HaveOccurred())
		Expect(v).To(Equal(currentVersion))
	})

	It("should be a Semver", func() {
		v, err := ProxyWallet.WALLETVERSION(nil)
		Expect(err).ToNot(HaveOccurred())
		_, err = semver.NewVersion(v)
		Expect(err).ToNot(HaveOccurred())
	})

	It("should not start with a v prefix", func() {
		v, err := ProxyWallet.WALLETVERSION(nil)
		Expect(err).ToNot(HaveOccurred())
		Expect(strings.HasPrefix(v, "v")).To(BeFalse())
	})

	It("should have spend limit of 1 ETH", func() {
		sl, err := ProxyWallet.SpendLimitValue(nil)
		Expect(err).ToNot(HaveOccurred())
		Expect(sl.String()).To(Equal(EthToWei(1).String()))
	})
})

var _ = AfterEach(func() {
	td := CurrentGinkgoTestDescription()
	if td.Failed {
		allPassed = false
	}

})

// var _ = AfterSuite(func() {
// 	if allPassed {
// 		TestRig.ExpectMinimumCoverage("wallet.sol", 98.00)
// 		TestRig.PrintGasUsage(os.Stdout)
// 	}
// })

func isSuccessful(tx *types.Transaction) bool {
	r, err := Backend.TransactionReceipt(context.Background(), tx.Hash())
	Expect(err).ToNot(HaveOccurred())
	return r.Status == types.ReceiptStatusSuccessful
}

var _ = AfterEach(func() {
	td := CurrentGinkgoTestDescription()
	if td.Failed {
		fmt.Fprintf(GinkgoWriter, "\nLast Executed Smart Contract Line for %s:%d\n", td.FileName, td.LineNumber)
		fmt.Fprintln(GinkgoWriter, TestRig.LastExecuted())
	}
	err := Backend.Close()
	Expect(err).ToNot(HaveOccurred())
})
