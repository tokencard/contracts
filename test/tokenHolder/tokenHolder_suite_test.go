package tokenHolder_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/tokencard/contracts/pkg/bindings/mocks"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	. "github.com/tokencard/contracts/test/shared"
)

var ERC20Contract1 *mocks.Token
var ERC20Contract1Address common.Address

var ERC20Contract2 *mocks.Token
var ERC20Contract2Address common.Address

var TKNBurner *mocks.BurnerToken
var TKNBurnerAddress common.Address

func init() {
	TestRig.AddCoverageForContracts(
		"../../build/tokenHolder/combined.json",
		// "../../build/burnerToken/combined.json",
		 "../../contracts")
}

func TestTokenHolderSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Contract Suite")
}

var _ = BeforeEach(func() {
	err := InitializeBackend()
	Expect(err).ToNot(HaveOccurred())

	var tx *types.Transaction
	ERC20Contract1Address, tx, ERC20Contract1, err = mocks.DeployToken(RandomAccount.TransactOpts(), Backend)
	Expect(err).ToNot(HaveOccurred())
	Backend.Commit()
	Expect(isSuccessful(tx)).To(BeTrue())

	ERC20Contract2Address, tx, ERC20Contract2, err = mocks.DeployToken(RandomAccount.TransactOpts(), Backend)
	Expect(err).ToNot(HaveOccurred())
	Backend.Commit()
	Expect(isSuccessful(tx)).To(BeTrue())

	TKNBurnerAddress, tx, TKNBurner, err = mocks.DeployBurnerToken(Owner.TransactOpts(), Backend, Owner.Address(), true)
	Expect(err).ToNot(HaveOccurred())
	Backend.Commit()
	Expect(isSuccessful(tx)).To(BeTrue())

})

var _ = AfterSuite(func() {
	TestRig.ExpectMinimumCoverage("tokenHolder.sol", 10.00)
	// TestRig.ExpectMinimumCoverage("burnerToken.sol", 10.00)
	TestRig.PrintGasUsage(os.Stdout)
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
