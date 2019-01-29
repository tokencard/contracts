package internals_test

import (
	"context"
	"testing"
	"os"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/tokencard/contracts/pkg/bindings/mocks"
	. "github.com/tokencard/contracts/test/shared"
)

var ParseIntScientificExporter *mocks.ParseIntScientificExporter
var ParseIntScientificExporterAddress common.Address

func init() {
	TestRig.AddCoverageForContracts(
		// "../../build/internals/parseIntScientific/combined.json",
		"../../build/mocks/parseIntScientific-exporter/combined.json",
		"../../contracts",
	)
}

func TestParseIntScientificSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Contract Suite")
}

var _ = BeforeEach(func() {
	err := InitializeBackend()
	Expect(err).ToNot(HaveOccurred())

	var tx *types.Transaction
	ParseIntScientificExporterAddress, tx, ParseIntScientificExporter, err = mocks.DeployParseIntScientificExporter(RandomAccount.TransactOpts(), Backend)
	Expect(err).ToNot(HaveOccurred())
	Backend.Commit()
	Expect(isSuccessful(tx)).To(BeTrue())
})

var _ = AfterEach(func() {
	err := Backend.Close()
	Expect(err).ToNot(HaveOccurred())
})

var _ = AfterSuite(func() {
	TestRig.ExpectMinimumCoverage("internals/parseIntScientific.sol", 100.0)
	TestRig.ExpectMinimumCoverage("mocks/parseIntScientific-exporter.sol", 100.0)
	TestRig.PrintGasUsage(os.Stdout)
})

func isSuccessful(tx *types.Transaction) bool {
	r, err := Backend.TransactionReceipt(context.Background(), tx.Hash())
	Expect(err).ToNot(HaveOccurred())
	return r.Status == types.ReceiptStatusSuccessful
}
