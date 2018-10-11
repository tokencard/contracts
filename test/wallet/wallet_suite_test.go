package wallet_test

import (
	"context"
	"fmt"
	"math/big"
	"os"
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

var ONE_ETH *big.Int
var ONE_FINNEY *big.Int
var ONE_GWEI *big.Int

func init() {
	var s bool

	ONE_ETH, s = big.NewInt(1).SetString("1000000000000000000", 10)
	if !s {
		panic("Could not parse one ETH")
	}

	ONE_FINNEY, s = big.NewInt(1).SetString("1000000000000000", 10)
	if !s {
		panic("Could not parse one ETH")
	}

	ONE_GWEI, s = big.NewInt(1).SetString("1000000000", 10)
	if !s {
		panic("Could not parse one ETH")
	}

	testRig.AddGenesisAccountAllocation(bankAccount.Address(), ethToWei(100))
	testRig.AddCoverageForContracts("../../build/wallet/combined.json", "../../contracts/wallet.sol")
}

var testRig = ethertest.NewTestRig()
var bankAccount = ethertest.NewAccount()

var _ = AfterSuite(func() {
	testRig.ExpectMinimumCoverage("wallet.sol:Wallet", 100.00)
	testRig.PrintGasUsage(os.Stdout)
})

func ethToWei(amount int) *big.Int {
	r := big.NewInt(1).Set(ONE_ETH)
	return r.Mul(r, big.NewInt(int64(amount)))
}

func finneyToWei(amount int) *big.Int {
	r := big.NewInt(1).Set(ONE_FINNEY)
	return r.Mul(r, big.NewInt(int64(amount)))
}

func gweiToWei(amount int) *big.Int {
	r := big.NewInt(1).Set(ONE_GWEI)
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

var maxDelta = big.NewInt(1000000)

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

var owner *ethertest.Account
var controller *ethertest.Account
var randomAccount *ethertest.Account

var _ = BeforeEach(func() {
	owner = ethertest.NewAccount()
	bankAccount.MustTransfer(be, owner.Address(), ONE_ETH)
	controller = ethertest.NewAccount()
	randomAccount = ethertest.NewAccount()
	bankAccount.MustTransfer(be, randomAccount.Address(), finneyToWei(500))
})

var controllerContract *bindings.Controller
var controllerContractAddress common.Address

var oraclizeMockAddrResolver *mocks.OraclizeAddrResolver
var oraclizeMockAddrResolverAddress common.Address

var oraclizeMock *mocks.Oraclize
var oraclizeMockAddress common.Address

var oracle *bindings.Oracle
var oracleAddress common.Address

var oracleResolver *bindings.Resolver
var oracleResolverAddress common.Address

var _ = BeforeEach(func() {
	var err error
	var tx *types.Transaction

	controllerContractAddress, tx, controllerContract, err = bindings.DeployController(bankAccount.TransactOpts(), be, bankAccount.Address())
	Expect(err).ToNot(HaveOccurred())
	be.Commit()
	Expect(isSuccessful(tx)).To(BeTrue())

	tx, err = controllerContract.AddController(bankAccount.TransactOpts(), controller.Address())
	Expect(err).ToNot(HaveOccurred())
	be.Commit()
	Expect(isSuccessful(tx)).To(BeTrue())

	Expect(bankAccount.Transfer(be, controller.Address(), ethToWei(1))).To(Succeed())

	oraclizeMockAddress, tx, oraclizeMock, err = mocks.DeployOraclize(bankAccount.TransactOpts(), be, bankAccount.Address())
	Expect(err).ToNot(HaveOccurred())
	be.Commit()
	Expect(isSuccessful(tx)).To(BeTrue())

	oraclizeMockAddrResolverAddress, tx, oraclizeMockAddrResolver, err = mocks.DeployOraclizeAddrResolver(bankAccount.TransactOpts(), be, oraclizeMockAddress)
	Expect(err).ToNot(HaveOccurred())
	be.Commit()
	Expect(isSuccessful(tx)).To(BeTrue())

	oracleAddress, tx, oracle, err = bindings.DeployOracle(bankAccount.TransactOpts(), be, oraclizeMockAddrResolverAddress, controllerContractAddress)
	Expect(err).ToNot(HaveOccurred())
	be.Commit()
	Expect(isSuccessful(tx)).To(BeTrue())

	oracleResolverAddress, tx, oracleResolver, err = bindings.DeployResolver(bankAccount.TransactOpts(), be, oracleAddress, controllerContractAddress)
	Expect(err).ToNot(HaveOccurred())
	be.Commit()
	Expect(isSuccessful(tx)).To(BeTrue())

})

func stringsToByte32(names ...string) [][32]byte {
	r := [][32]byte{}
	for _, n := range names {
		nb := [32]byte{}
		copy(nb[:], []byte(n))
		r = append(r, nb)
	}
	return r
}

var tkn *mocks.Token
var tkna common.Address
var _ = BeforeEach(func() {
	var err error
	var tx *types.Transaction
	tkna, tx, tkn, err = mocks.DeployToken(bankAccount.TransactOpts(), be)
	Expect(err).ToNot(HaveOccurred())
	be.Commit()
	Expect(isSuccessful(tx)).To(BeTrue())

	tx, err = oracle.AddTokens(controller.TransactOpts(), []common.Address{tkna}, stringsToByte32("TKN"), []byte{8})
	Expect(err).ToNot(HaveOccurred())
	be.Commit()
	Expect(isSuccessful(tx)).To(BeTrue())

	tx, err = oracle.UpdateTokenRate(controller.TransactOpts(), tkna, big.NewInt(int64(0.00001633*math.Pow10(18))))
	Expect(err).ToNot(HaveOccurred())
	be.Commit()
	Expect(isSuccessful(tx)).To(BeTrue())

})

var w *bindings.Wallet
var wa common.Address

var _ = BeforeEach(func() {
	var err error
	var tx *types.Transaction
	wa, tx, w, err = bindings.DeployWallet(
		bankAccount.TransactOpts(),
		be,
		owner.Address(),
		true,
		oracleResolverAddress,
		controllerContractAddress,
		ethToWei(100),
	)
	Expect(err).ToNot(HaveOccurred())
	be.Commit()
	Expect(isSuccessful(tx)).To(BeTrue())
})
