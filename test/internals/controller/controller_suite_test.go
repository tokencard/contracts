package controller_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/core/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tokencard/contracts/test/shared"
)

func init() {
	TestRig.AddCoverageForContracts("../../../build/internals/controller/combined.json", "../../../contracts")
}

func TestWalletSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Controller Suite")
}

var _ = BeforeEach(func() {
	err := InitializeBackend()
	Expect(err).ToNot(HaveOccurred())
})

var someFailed bool

var _ = AfterEach(func() {
	if CurrentGinkgoTestDescription().Failed {
		someFailed = true
	}
})

var _ = AfterSuite(func() {
	if !someFailed {
		TestRig.ExpectMinimumCoverage("internals/controller.sol", 100.00)
		TestRig.PrintGasUsage(os.Stdout)
	}
})

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
})

// func isGasExhausted(tx *types.Transaction, gasLimit uint64) bool {
// 	r, err := Backend.TransactionReceipt(context.Background(), tx.Hash())
// 	Expect(err).ToNot(HaveOccurred())
// 	if r.Status == types.ReceiptStatusSuccessful {
// 		return false
// 	}
// 	return r.GasUsed == gasLimit
// }
