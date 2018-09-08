package wallet_test

import (
	"context"
	"fmt"
	"math/big"
	"testing"

	"math"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	gtypes "github.com/onsi/gomega/types"
	"github.com/tokencard/contracts/pkg/bindings"
	"github.com/tokencard/contracts/pkg/bindings/mocks"
	"github.com/tokencard/ethertest"
)

func TestWalletSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Contract Suite")
}

var HUNDRED_ETH *big.Int
var TEN_ETH *big.Int
var ONE_ETH *big.Int
var ONE_FINNEY *big.Int
var FIVE_HUNDRED_FINNEY *big.Int
var ONE_GWEI *big.Int

func init() {
	var s bool

	HUNDRED_ETH, s = big.NewInt(1).SetString("100000000000000000000", 10)
	if !s {
		panic("Could not parse one ETH")
	}

	TEN_ETH, s = big.NewInt(1).SetString("10000000000000000000", 10)
	if !s {
		panic("Could not parse one ETH")
	}

	ONE_ETH, s = big.NewInt(1).SetString("1000000000000000000", 10)
	if !s {
		panic("Could not parse one ETH")
	}

	ONE_FINNEY, s = big.NewInt(1).SetString("1000000000000000", 10)
	if !s {
		panic("Could not parse one ETH")
	}

	FIVE_HUNDRED_FINNEY, s = big.NewInt(1).SetString("500000000000000000", 10)
	if !s {
		panic("Could not parse one ETH")
	}

	ONE_GWEI, s = big.NewInt(1).SetString("1000000000", 10)
	if !s {
		panic("Could not parse one ETH")
	}

	testRig.AddGenesisAccountAllocation(bankWallet.Address(), HUNDRED_ETH)
	testRig.AddCoverageForContracts("../../build/wallet/combined.json", "../../contracts/wallet.sol")
}

var testRig = ethertest.NewTestRig()
var bankWallet = ethertest.NewWallet()

var _ = AfterSuite(func() {
	testRig.ExpectMinimumCoverage("wallet.sol:Wallet", 100.0)
})

func ethToWei(amount int) *big.Int {
	r := big.NewInt(1).Set(ONE_ETH)
	return r.Mul(r, big.NewInt(int64(amount)))
}

func isSuccessful(tx *types.Transaction) bool {
	r, err := be.TransactionReceipt(context.Background(), tx.Hash())
	Expect(err).ToNot(HaveOccurred())
	return r.Status == types.ReceiptStatusSuccessful
}

var be ethertest.TestBackend

var _ = BeforeEach(func() {
	be = testRig.NewTestBackend()
})

type almostEqual string

func AlmostEqual(val string) gtypes.GomegaMatcher {
	return almostEqual(val)
}

var maxDelta = big.NewInt(100000)

func (m almostEqual) Match(actual interface{}) (success bool, err error) {

	expected, ok := big.NewInt(0).SetString(string(m), 10)
	if !ok {
		return false, fmt.Errorf("%q is not a decimal integer", expected)
	}

	actualString, ok := actual.(string)
	if !ok {
		return false, fmt.Errorf("actual value is not a string")
	}

	act, ok := big.NewInt(0).SetString(actualString, 10)
	if !ok {
		return false, fmt.Errorf("%q is not a decimal integer", actualString)
	}

	delta := big.NewInt(0)

	delta = delta.Sub(act, expected)

	return delta.Abs(delta).Cmp(maxDelta) == -1, nil
}

func (m almostEqual) FailureMessage(actual interface{}) (message string) {
	return fmt.Sprintf("Actual value %v is not almost equal to expected %v", m, actual)
}

func (m almostEqual) NegatedFailureMessage(actual interface{}) (message string) {
	return fmt.Sprintf("Actual value %v is almost equal to expected %v", m, actual)
}

var owner *ethertest.Wallet
var controller *ethertest.Wallet
var randomPerson *ethertest.Wallet

var _ = BeforeEach(func() {
	owner = ethertest.NewWallet()
	bankWallet.MustTransfer(be, owner.Address(), ONE_ETH)
	controller = ethertest.NewWallet()
	randomPerson = ethertest.NewWallet()
	bankWallet.MustTransfer(be, randomPerson.Address(), FIVE_HUNDRED_FINNEY)
})

var oraclizeMockAddrResolver *mocks.OraclizeAddrResolver
var oraclizeMockAddrResolverAddress common.Address

var oraclizeMock *mocks.Oraclize
var oraclizeMockAddress common.Address

var oracle *bindings.Oracle
var oracleAddress common.Address

var _ = BeforeEach(func() {
	var err error

	oraclizeMockAddress, _, oraclizeMock, err = mocks.DeployOraclize(bankWallet.TransactOpts(), be, bankWallet.Address())
	Expect(err).ToNot(HaveOccurred())

	oraclizeMockAddrResolverAddress, _, oraclizeMockAddrResolver, err = mocks.DeployOraclizeAddrResolver(bankWallet.TransactOpts(), be, oraclizeMockAddress)
	Expect(err).ToNot(HaveOccurred())

	oracleAddress, _, oracle, err = bindings.DeployOracle(bankWallet.TransactOpts(), be, oraclizeMockAddrResolverAddress)
	Expect(err).ToNot(HaveOccurred())

	be.Commit()
})

var tkn *mocks.Token
var tkna common.Address
var _ = BeforeEach(func() {
	var err error
	tkna, _, tkn, err = mocks.DeployToken(bankWallet.TransactOpts(), be)
	Expect(err).ToNot(HaveOccurred())

	_, err = oracle.AddToken(bankWallet.TransactOpts(), tkna, "TKN", 8)
	Expect(err).ToNot(HaveOccurred())

	_, err = oracle.UpdateRateManual(bankWallet.TransactOpts(), tkna, big.NewInt(int64(0.001633*math.Pow10(18))))
	Expect(err).ToNot(HaveOccurred())

	be.Commit()
})

var w *bindings.Wallet
var wa common.Address

var _ = BeforeEach(func() {
	var err error
	wa, _, w, err = bindings.DeployWallet(
		bankWallet.TransactOpts(),
		be,
		owner.Address(),
		oracleAddress,
		[]common.Address{controller.Address()},
	)
	Expect(err).ToNot(HaveOccurred())
	be.Commit()
})
