package externals_test

import (
	"context"
	"testing"
	// "os"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/tokencard/contracts/pkg/bindings/mocks"
	. "github.com/tokencard/contracts/test/shared"
)

var ParseIntScientificExporter *mocks.ParseIntScientificExporter
var ParseIntScientificExporterAddress common.Address

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

// var _ = AfterSuite(func() {
// 	TestRig.ExpectMinimumCoverage("parseIntScientific.sol:ParseIntScientific", 99.99)
// 	TestRig.PrintGasUsage(os.Stdout)
// })

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
