package bytesUtils_test

import (
	"context"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/tokencard/contracts/v2/pkg/bindings/mocks"
	. "github.com/tokencard/contracts/v2/test/shared"
)

var BytesUtilsExporter *mocks.BytesUtilsExporter
var BytesUtilsExporterAddress common.Address

func init() {
	TestRig.AddCoverageForContracts(
		"../../../build/mocks/bytesUtilsExporter/combined.json",
		"../../../contracts",
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
	BytesUtilsExporterAddress, tx, BytesUtilsExporter, err = mocks.DeployBytesUtilsExporter(RandomAccount.TransactOpts(), Backend)
	Expect(err).ToNot(HaveOccurred())
	Backend.Commit()
	Expect(isSuccessful(tx)).To(BeTrue())
})

var _ = AfterEach(func() {
	err := Backend.Close()
	Expect(err).ToNot(HaveOccurred())
})

var _ = AfterSuite(func() {
	TestRig.ExpectMinimumCoverage("internals/bytesUtils.sol", 100.0)
	TestRig.ExpectMinimumCoverage("mocks/bytesUtilsExporter.sol", 100.0)
	TestRig.PrintGasUsage(os.Stdout)
})

func isSuccessful(tx *types.Transaction) bool {
	r, err := Backend.TransactionReceipt(context.Background(), tx.Hash())
	Expect(err).ToNot(HaveOccurred())
	return r.Status == types.ReceiptStatusSuccessful
}
