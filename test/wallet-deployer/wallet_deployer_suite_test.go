package wallet_deployer_test

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"testing"
    "math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/tokencard/contracts/v3/pkg/bindings"
	"github.com/tokencard/contracts/v3/pkg/bindings/externals/upgradeability"
	. "github.com/tokencard/contracts/v3/test/shared"
)

var WalletImplementation *bindings.Wallet
var WalletImplementationAddress common.Address

var WalletCacheAddress common.Address
var WalletCache *bindings.WalletCache

var WalletDeployerAddress common.Address
var WalletDeployer *bindings.WalletDeployer

func deployInitProxy(owner common.Address, spendLimit *big.Int) common.Address {
	RandomProxyAddress, tx, _, err := upgradeability.DeployUpgradeabilityProxy(BankAccount.TransactOpts(), Backend, WalletImplementationAddress, nil)
	Expect(err).ToNot(HaveOccurred())
	Backend.Commit()
	Expect(isSuccessful(tx)).To(BeTrue())

	RandomProxy, err := bindings.NewWallet(RandomProxyAddress, Backend)
	tx, err = RandomProxy.InitializeWallet(BankAccount.TransactOpts(), owner, false, ENSRegistryAddress, TokenWhitelistNode, ControllerNode, LicenceNode, spendLimit)
	Expect(err).ToNot(HaveOccurred())
	Backend.Commit()
	Expect(isSuccessful(tx)).To(BeTrue())

	return RandomProxyAddress
}

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
		TestRig.ExpectMinimumCoverage("walletCache.sol", 80.00)
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

func isSuccessful(tx *types.Transaction) bool {
	r, err := Backend.TransactionReceipt(context.Background(), tx.Hash())
	Expect(err).ToNot(HaveOccurred())
	return r.Status == types.ReceiptStatusSuccessful
}

var WalletDeployerNode = EnsNode("wallet-deployer.v3.tokencard.eth")
var WalletCacheNode = EnsNode("wallet-cache.v3.tokencard.eth")

var _ = BeforeEach(func() {
	err := InitializeBackend()
	Expect(err).ToNot(HaveOccurred())
	var tx *types.Transaction
<<<<<<< HEAD
	// deploy wallet implementation
	WalletImplementationAddress, tx, WalletImplementation, err = bindings.DeployWallet(BankAccount.TransactOpts(), Backend)
	Expect(err).ToNot(HaveOccurred())
	Backend.Commit()
	Expect(isSuccessful(tx)).To(BeTrue())
	WalletCacheAddress, tx, WalletCache, err = bindings.DeployWalletCache(BankAccount.TransactOpts(), Backend, WalletImplementationAddress, ENSRegistryAddress, EthToWei(1), [32]byte{}, [32]byte{}, [32]byte{}, [32]byte{})
=======
	WalletCacheAddress, tx, WalletCache, err = bindings.DeployWalletCache(BankAccount.TransactOpts(), Backend, ENSRegistryAddress, big.NewInt(10000), [32]byte{}, [32]byte{}, [32]byte{}, [32]byte{})
>>>>>>> 511d3647... Use stablecoin as daily limit
	Expect(err).ToNot(HaveOccurred())
	WalletDeployerAddress, tx, WalletDeployer, err = bindings.DeployWalletDeployer(BankAccount.TransactOpts(), Backend, ENSRegistryAddress, [32]byte{}, [32]byte{})
	Expect(err).ToNot(HaveOccurred())
	Backend.Commit()
	Expect(isSuccessful(tx)).To(BeTrue())

	//Register with ENS
	{
		// set WalletDeployer node owner
		tx, err = ENSRegistry.SetSubnodeOwner(BankAccount.TransactOpts(), EnsNode("tokencard.eth"), LabelHash("v3"), BankAccount.Address())
		Expect(err).ToNot(HaveOccurred())
		Backend.Commit()
		tx, err = ENSRegistry.SetSubnodeOwner(BankAccount.TransactOpts(), EnsNode("v3.tokencard.eth"), LabelHash("wallet-deployer"), BankAccount.Address())
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
		tx, err = ENSRegistry.SetSubnodeOwner(BankAccount.TransactOpts(), EnsNode("tokencard.eth"), LabelHash("v3"), BankAccount.Address())
		Expect(err).ToNot(HaveOccurred())
		Backend.Commit()
		tx, err = ENSRegistry.SetSubnodeOwner(BankAccount.TransactOpts(), EnsNode("v3.tokencard.eth"), LabelHash("wallet-cache"), BankAccount.Address())
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
