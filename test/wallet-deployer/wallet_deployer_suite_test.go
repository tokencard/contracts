package wallet_deployer_test

import (
	"context"
	"fmt"
	"os"
	"testing"

    "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
    "github.com/tokencard/contracts/v2/pkg/bindings"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tokencard/contracts/v2/test/shared"
)

func ethCall(tx *types.Transaction) ([]byte, error) {
	msg, _ := tx.AsMessage(types.HomesteadSigner{})

	calMsg := ethereum.CallMsg{
		From:     msg.From(),
		To:       msg.To(),
		Gas:      msg.Gas(),
		GasPrice: msg.GasPrice(),
		Value:    msg.Value(),
		Data:     msg.Data(),
	}

	return Backend.CallContract(context.Background(), calMsg, nil)
}

func TestWalletDeployer(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "WalletDeployer Suite")
}

var _ = BeforeSuite(func() {
	TestRig.AddCoverageForContracts(
		"../../build/walletDeployer/combined.json",
		"../../contracts",
	)
	TestRig.AddCoverageForContracts(
		"../../build/walletCache/combined.json",
		"../../contracts",
	)
})

var allPassed = true

var _ = AfterEach(func() {
	td := CurrentGinkgoTestDescription()
	if td.Failed {
		allPassed = false
	}

})

var _ = AfterSuite(func() {
	if allPassed {
		TestRig.ExpectMinimumCoverage("walletDeployer.sol", 92.00)
		TestRig.ExpectMinimumCoverage("walletCache.sol", 88.00)
		TestRig.PrintGasUsage(os.Stdout)
	}
})

var _ = AfterEach(func() {
	td := CurrentGinkgoTestDescription()

	if td.Failed {
		fmt.Fprintf(GinkgoWriter, "\nLast Executed Smart Contract Line for %s:%d\n", td.FileName, td.LineNumber)
		fmt.Fprintln(GinkgoWriter, TestRig.LastExecuted())
	}
})

var _ = AfterEach(func() {
	err := Backend.Close()
	Expect(err).ToNot(HaveOccurred())

})

var WalletCacheAddress common.Address
var WalletCache *bindings.WalletCache

var WalletDeployerAddress common.Address
var WalletDeployer *bindings.WalletDeployer

func isSuccessful(tx *types.Transaction) bool {
	r, err := Backend.TransactionReceipt(context.Background(), tx.Hash())
	Expect(err).ToNot(HaveOccurred())
	return r.Status == types.ReceiptStatusSuccessful
}

var WalletDeployerNode = EnsNode("wallet-deployer.tokencard.eth")
var WalletCacheNode = EnsNode("wallet-cache.tokencard.eth")

var _ = BeforeEach(func() {
	err := InitializeBackend()
	Expect(err).ToNot(HaveOccurred())
	var tx *types.Transaction
	WalletCacheAddress, tx, WalletCache, err = bindings.DeployWalletCache(BankAccount.TransactOpts(), Backend, ENSRegistryAddress, EthToWei(1), [32]byte{}, [32]byte{}, [32]byte{}, [32]byte{})
	Expect(err).ToNot(HaveOccurred())
	WalletDeployerAddress, tx, WalletDeployer, err = bindings.DeployWalletDeployer(BankAccount.TransactOpts(), Backend, ENSRegistryAddress, [32]byte{}, [32]byte{})
	Expect(err).ToNot(HaveOccurred())
	Backend.Commit()
	Expect(isSuccessful(tx)).To(BeTrue())

	//Register with ENS
	{
		// set WalletDeployer node owner
		tx, err = ENSRegistry.SetSubnodeOwner(BankAccount.TransactOpts(), EnsNode("tokencard.eth"), LabelHash("wallet-deployer"), BankAccount.Address())
		Expect(err).ToNot(HaveOccurred())
		Backend.Commit()
		// Register WalletDeployer with ENS
		tx, err = ENSRegistry.SetResolver(BankAccount.TransactOpts(), WalletDeployerNode, ENSResolverAddress)
		Expect(err).ToNot(HaveOccurred())
		Backend.Commit()
		Expect(isSuccessful(tx)).To(BeTrue())
		tx, err = ENSResolver.SetAddr(BankAccount.TransactOpts(), WalletDeployerNode, WalletDeployerAddress)
		Expect(err).ToNot(HaveOccurred())
		Backend.Commit()
		Expect(isSuccessful(tx)).To(BeTrue())
	}
	{
		// set WalletCache node owner
		tx, err = ENSRegistry.SetSubnodeOwner(BankAccount.TransactOpts(), EnsNode("tokencard.eth"), LabelHash("wallet-cache"), BankAccount.Address())
		Expect(err).ToNot(HaveOccurred())
		Backend.Commit()
		// Register WalletCache with ENS
		tx, err = ENSRegistry.SetResolver(BankAccount.TransactOpts(), WalletCacheNode, ENSResolverAddress)
		Expect(err).ToNot(HaveOccurred())
		Backend.Commit()
		Expect(isSuccessful(tx)).To(BeTrue())
		tx, err = ENSResolver.SetAddr(BankAccount.TransactOpts(), WalletCacheNode, WalletCacheAddress)
		Expect(err).ToNot(HaveOccurred())
		Backend.Commit()
		Expect(isSuccessful(tx)).To(BeTrue())
	}
})
