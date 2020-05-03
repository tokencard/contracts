package controller_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/core/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tokencard/contracts/v3/test/shared"
)

func init() {
	TestRig.AddCoverageForContracts("../../build/controller/combined.json", "../../contracts")
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

var _ = AfterEach(func() {
	err := Backend.Close()
	Expect(err).ToNot(HaveOccurred())
})

var _ = AfterSuite(func() {
	if !someFailed {
		TestRig.ExpectMinimumCoverage("controller.sol", 100.00)
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
