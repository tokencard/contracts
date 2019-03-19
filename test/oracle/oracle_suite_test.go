package oracle_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/core/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/tokencard/contracts/pkg/bindings"
	"github.com/tokencard/contracts/pkg/bindings/internals"
	. "github.com/tokencard/contracts/test/shared"
)

const GAS_LIMIT = 2000000

func init() {
	TestRig.AddCoverageForContracts(
		"../../build/oracle/combined.json",
		"../../contracts",
	)
}

func TestOracleSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Oracle Suite")
}

var _ = BeforeEach(func() {
	err := InitializeBackend()
	Expect(err).ToNot(HaveOccurred())

	var tx *types.Transaction
	OracleAddress, tx, Oracle, err = bindings.DeployOracle(BankAccount.TransactOpts(), Backend, OraclizeResolverAddress, ENSRegistryAddress, ControllerName, TokenWhitelistName)
	Expect(err).ToNot(HaveOccurred())
	Backend.Commit()
	Expect(isSuccessful(tx)).To(BeTrue())

	{
		// Register oracle with ENS

		tx, err = ENSRegistry.SetResolver(BankAccount.TransactOpts(), OracleName, ENSResolverAddress)
		Expect(err).ToNot(HaveOccurred())
		Backend.Commit()
		Expect(isSuccessful(tx)).To(BeTrue())
		tx, err = ENSResolver.SetAddr(BankAccount.TransactOpts(), OracleName, OracleAddress)
		Expect(err).ToNot(HaveOccurred())
		Backend.Commit()
		Expect(isSuccessful(tx)).To(BeTrue())
	}

	TokenWhitelistAddress, tx, TokenWhitelist, err = internals.DeployTokenWhitelist(BankAccount.TransactOpts(), Backend, ENSRegistryAddress, OracleName, ControllerName, Owner.Address(), true, StablecoinAddress)
	Expect(err).ToNot(HaveOccurred())
	Backend.Commit()
	Expect(isSuccessful(tx)).To(BeTrue())

	{
		// Register tokenWhitelist with ENS
		tx, err = ENSRegistry.SetResolver(BankAccount.TransactOpts(), TokenWhitelistName, ENSResolverAddress)
		Expect(err).ToNot(HaveOccurred())
		Backend.Commit()
		Expect(isSuccessful(tx)).To(BeTrue())
		tx, err = ENSResolver.SetAddr(BankAccount.TransactOpts(), TokenWhitelistName, TokenWhitelistAddress)
		Expect(err).ToNot(HaveOccurred())
		Backend.Commit()
		Expect(isSuccessful(tx)).To(BeTrue())
	}

})

var _ = AfterEach(func() {
	td := CurrentGinkgoTestDescription()

	if td.Failed {
		fmt.Fprintf(GinkgoWriter, "\nLast Executed Smart Contract Line for %s:%d\n", td.FileName, td.LineNumber)
		fmt.Fprintln(GinkgoWriter, TestRig.LastExecuted())
	}
	err := Backend.Close()
	Expect(err).ToNot(HaveOccurred())
})

var _ = AfterSuite(func() {
	TestRig.ExpectMinimumCoverage("oracle.sol", 100.00)
	TestRig.PrintGasUsage(os.Stdout)
})

func isSuccessful(tx *types.Transaction) bool {
	r, err := Backend.TransactionReceipt(context.Background(), tx.Hash())
	Expect(err).ToNot(HaveOccurred())
	return r.Status == types.ReceiptStatusSuccessful
}

func isGasExhausted(tx *types.Transaction, gasLimit uint64) bool {
	r, err := Backend.TransactionReceipt(context.Background(), tx.Hash())
	Expect(err).ToNot(HaveOccurred())
	if r.Status == types.ReceiptStatusSuccessful {
		return false
	}
	return r.GasUsed == gasLimit
}
