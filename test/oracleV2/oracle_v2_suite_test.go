package oracle_v2_test

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/tokencard/contracts/pkg/bindings"
	"github.com/tokencard/ethertest"

	"testing"
)

func TestOracleV2Suite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Contract Suite")
}

func ethToWei(amount int) *big.Int {
	r := big.NewInt(1000000000000000000)
	return r.Mul(r, big.NewInt(int64(amount)))
}

func finneyToWei(amount int) *big.Int {
	r := big.NewInt(1000000000000000)
	return r.Mul(r, big.NewInt(int64(amount)))
}

var testRig = ethertest.NewTestRig()
var controllerWallet = ethertest.NewWallet()
var oraclizeCallbackWallet = ethertest.NewWallet()

var _ = BeforeSuite(func() {
	testRig.AddGenesisAccountAllocation(controllerWallet.Address(), ethToWei(1000))
	// testRig.AddGenesisAccountAllocation(oraclizeCallbackWallet.Address(), ethToWei(1000))
	testRig.AddCoverageForContracts("../../build/oracleV2/combined.json", "../../contracts/oracleV2.sol")
})

var _ = AfterSuite(func() {
	testRig.ExpectMinimumCoverage("oracleV2.sol:Oracle", 50.0)
})

func balanceOf(a common.Address) *big.Int {
	b, err := be.BalanceAt(context.Background(), a, nil)
	Expect(err).ToNot(HaveOccurred())
	return b

}

var be ethertest.TestBackend

func isSuccessful(tx *types.Transaction) bool {
	r, err := be.TransactionReceipt(context.Background(), tx.Hash())
	Expect(err).ToNot(HaveOccurred())
	return r.Status == types.ReceiptStatusSuccessful
}

var oraclizeMockAddrResolver *bindings.MockOraclizeAddrResolver
var oraclizeMockAddrResolverAddress common.Address

var oraclizeMock *bindings.MockOraclize
var oraclizeMockAddress common.Address

var oracle *bindings.OracleV2
var oa common.Address

var _ = BeforeEach(func() {
	be = testRig.NewTestBackend()

	var err error

	oraclizeMockAddress, _, oraclizeMock, err = bindings.DeployMockOraclize(controllerWallet.TransactOpts(), be, oraclizeCallbackWallet.Address())
	Expect(err).ToNot(HaveOccurred())

	oraclizeMockAddrResolverAddress, _, oraclizeMockAddrResolver, err = bindings.DeployMockOraclizeAddrResolver(controllerWallet.TransactOpts(), be, oraclizeMockAddress)
	Expect(err).ToNot(HaveOccurred())

	oa, _, oracle, err = bindings.DeployOracleV2(controllerWallet.TransactOpts(), be, oraclizeMockAddrResolverAddress)
	Expect(err).ToNot(HaveOccurred())

	be.Commit()
})
