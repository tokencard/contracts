package oracle_test

import (
	"context"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/tokencard/contracts/pkg/bindings/mocks"
	. "github.com/tokencard/contracts/test/shared"
)

var Base64Exporter *mocks.Base64Exporter
var Base64ExporterAddress common.Address

// func init() {
// 	TestRig.AddCoverageForContracts("../../build/mocks/base64-exporter/combined.json", "../../contracts/externals/base64.sol")
// }

func TestBase64Suite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Contract Suite")
}

var _ = BeforeEach(func() {
	err := InitializeBackend()
	Expect(err).ToNot(HaveOccurred())

	var tx *types.Transaction
	Base64ExporterAddress, tx, Base64Exporter, err = mocks.DeployBase64Exporter(RandomAccount.TransactOpts(), Backend)
	Expect(err).ToNot(HaveOccurred())
	Backend.Commit()
	Expect(isSuccessful(tx)).To(BeTrue())
})

// var _ = AfterSuite(func() {
// 	TestRig.ExpectMinimumCoverage("base64.sol:Base64", 10.0)
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
