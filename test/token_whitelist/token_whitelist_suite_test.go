package token_whitelist_test

import (
	"context"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/tokencard/contracts/v3/pkg/bindings"
	"github.com/tokencard/contracts/v3/pkg/bindings/mocks"
	. "github.com/tokencard/contracts/v3/test/shared"
)

var TokenWhitelistableExporter *mocks.TokenWhitelistableExporter
var TokenWhitelistableExporterAddress common.Address

func init() {
	TestRig.AddCoverageForContracts(
		"../../build/tokenWhitelist/combined.json",
		"../../contracts",
	)

	TestRig.AddCoverageForContracts(
		"../../build/mocks/tokenWhitelistableExporter/combined.json",
		"../../contracts",
	)
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

func TestTokenWhitelistSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "TokenWhitelist Suite")
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

	TokenWhitelistAddress, tx, TokenWhitelist, err = bindings.DeployTokenWhitelist(BankAccount.TransactOpts(), Backend, ENSRegistryAddress, OracleName, ControllerName, StablecoinAddress)
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

	TokenWhitelistableExporterAddress, tx, TokenWhitelistableExporter, err = mocks.DeployTokenWhitelistableExporter(BankAccount.TransactOpts(), Backend, ENSRegistryAddress, TokenWhitelistName)
	Expect(err).ToNot(HaveOccurred())
	Backend.Commit()
	Expect(isSuccessful(tx)).To(BeTrue())

})

var _ = AfterEach(func() {
	err := Backend.Close()
	Expect(err).ToNot(HaveOccurred())

})

var _ = AfterSuite(func() {
	TestRig.ExpectMinimumCoverage("tokenWhitelist.sol", 98.00)
	TestRig.ExpectMinimumCoverage("internals/tokenWhitelistable.sol", 100.0)
	TestRig.PrintGasUsage(os.Stdout)
})

func isSuccessful(tx *types.Transaction) bool {
	r, err := Backend.TransactionReceipt(context.Background(), tx.Hash())
	Expect(err).ToNot(HaveOccurred())
	return r.Status == types.ReceiptStatusSuccessful
}
