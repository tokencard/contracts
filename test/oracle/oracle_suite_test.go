package oracle_test

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/tokencard/contracts/pkg/bindings"
	"github.com/tokencard/contracts/pkg/bindings/mocks"
	"github.com/tokencard/ethertest"
)

func TestOracleSuite(t *testing.T) {
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
var randomWallet = ethertest.NewWallet()
var oraclizeCallbackWallet = ethertest.NewWallet()

var _ = BeforeSuite(func() {
	testRig.AddGenesisAccountAllocation(controllerWallet.Address(), ethToWei(1000))
	testRig.AddGenesisAccountAllocation(randomWallet.Address(), ethToWei(1000))
	testRig.AddGenesisAccountAllocation(oraclizeCallbackWallet.Address(), ethToWei(1000))
	testRig.AddCoverageForContracts("../../build/oracle/combined.json", "../../contracts/oracle.sol")
})

var _ = AfterSuite(func() {
	testRig.ExpectMinimumCoverage("oracle.sol:Oracle", 40.07)
})

var be ethertest.TestBackend

func isSuccessful(tx *types.Transaction) bool {
	r, err := be.TransactionReceipt(context.Background(), tx.Hash())
	Expect(err).ToNot(HaveOccurred())
	return r.Status == types.ReceiptStatusSuccessful
}

func isGasExhausted(tx *types.Transaction, gasLimit uint64) bool {
	r, err := be.TransactionReceipt(context.Background(), tx.Hash())
	Expect(err).ToNot(HaveOccurred())
	if r.Status == types.ReceiptStatusSuccessful {
		return false
	}
	return r.GasUsed == gasLimit
}

var oraclizeMockAddrResolver *mocks.OraclizeAddrResolver
var oraclizeMockAddrResolverAddress common.Address

var oraclizeMock *mocks.Oraclize
var oraclizeMockAddress common.Address

var oracle *bindings.Oracle
var oracleAddress common.Address

var controllerContract *bindings.Controller
var controllerContractAddress common.Address

var _ = BeforeEach(func() {
	be = testRig.NewTestBackend()

	var err error
	var tx *types.Transaction

	controllerContractAddress, tx, controllerContract, err = bindings.DeployController(controllerWallet.TransactOpts(), be, controllerWallet.Address())
	Expect(err).ToNot(HaveOccurred())
	be.Commit()
	Expect(isSuccessful(tx)).To(BeTrue())

	oraclizeMockAddress, tx, oraclizeMock, err = mocks.DeployOraclize(controllerWallet.TransactOpts(), be, oraclizeCallbackWallet.Address())
	Expect(err).ToNot(HaveOccurred())
	be.Commit()
	Expect(isSuccessful(tx)).To(BeTrue())

	oraclizeMockAddrResolverAddress, tx, _, err = mocks.DeployOraclizeAddrResolver(controllerWallet.TransactOpts(), be, oraclizeMockAddress)
	Expect(err).ToNot(HaveOccurred())
	be.Commit()
	Expect(isSuccessful(tx)).To(BeTrue())

	oracleAddress, tx, oracle, err = bindings.DeployOracle(controllerWallet.TransactOpts(), be, oraclizeMockAddrResolverAddress, controllerContractAddress)
	Expect(err).ToNot(HaveOccurred())
	be.Commit()
	Expect(isSuccessful(tx)).To(BeTrue())
})
