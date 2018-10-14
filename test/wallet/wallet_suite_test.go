package wallet_test

import (
	"context"
	"fmt"
	"math"
	"math/big"
	"os"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	gtypes "github.com/onsi/gomega/types"
	"github.com/tokencard/contracts/pkg/bindings"
	"github.com/tokencard/contracts/pkg/bindings/external"
	"github.com/tokencard/contracts/pkg/bindings/mocks"
	"github.com/tokencard/ethertest"
)

func TestWalletSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Contract Suite")
}

// lifted from https://github.com/tronprotocol/tron-demo/blob/master/demo/go-client-api/vendor/github.com/ethereum/go-ethereum/contracts/ens/ens.go
func ensParentNode(name string) (common.Hash, common.Hash) {
	parts := strings.SplitN(name, ".", 2)
	label := crypto.Keccak256Hash([]byte(parts[0]))
	if len(parts) == 1 {
		return [32]byte{}, label
	} else {
		parentNode, parentLabel := ensParentNode(parts[1])
		return crypto.Keccak256Hash(parentNode[:], parentLabel[:]), label
	}
}

func labelHash(label string) common.Hash {
	return crypto.Keccak256Hash([]byte(label))
}

func ensNode(name string) common.Hash {
	if name == "" {
		return common.Hash{}
	}
	parentNode, parentLabel := ensParentNode(name)
	return crypto.Keccak256Hash(parentNode[:], parentLabel[:])
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

var ensResolver *bindings.Resolver
var ensResolverAddress common.Address

var ensAddress common.Address
var ens *external.ENSRegistry

var oracleName = ensNode("oracle.tokencard.eth")
var controllerName = ensNode("controller.tokencard.eth")

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

	ensAddress, tx, ens, err = external.DeployENSRegistry(bankAccount.TransactOpts(), be)
	Expect(err).ToNot(HaveOccurred())
	be.Commit()
	Expect(isSuccessful(tx)).To(BeTrue())

	tx, err = ens.SetSubnodeOwner(bankAccount.TransactOpts(), ensNode(""), labelHash("eth"), bankAccount.Address())
	Expect(err).ToNot(HaveOccurred())
	be.Commit()
	Expect(isSuccessful(tx)).To(BeTrue())

	tx, err = ens.SetSubnodeOwner(bankAccount.TransactOpts(), ensNode("eth"), labelHash("tokencard"), bankAccount.Address())
	Expect(err).ToNot(HaveOccurred())
	be.Commit()
	Expect(isSuccessful(tx)).To(BeTrue())

	tx, err = ens.SetSubnodeOwner(bankAccount.TransactOpts(), ensNode("tokencard.eth"), labelHash("controller"), bankAccount.Address())
	Expect(err).ToNot(HaveOccurred())
	be.Commit()
	Expect(isSuccessful(tx)).To(BeTrue())

	tx, err = ens.SetSubnodeOwner(bankAccount.TransactOpts(), ensNode("tokencard.eth"), labelHash("oracle"), bankAccount.Address())
	Expect(err).ToNot(HaveOccurred())
	be.Commit()
	Expect(isSuccessful(tx)).To(BeTrue())

	ensResolverAddress, tx, ensResolver, err = bindings.DeployResolver(bankAccount.TransactOpts(), be, ensAddress)
	Expect(err).ToNot(HaveOccurred())
	be.Commit()
	Expect(isSuccessful(tx)).To(BeTrue())

	{
		// register controller with ENS

		tx, err = ens.SetResolver(bankAccount.TransactOpts(), controllerName, ensResolverAddress)
		Expect(err).ToNot(HaveOccurred())
		be.Commit()
		Expect(isSuccessful(tx)).To(BeTrue())

		tx, err = ensResolver.SetAddr(bankAccount.TransactOpts(), controllerName, controllerContractAddress)
		Expect(err).ToNot(HaveOccurred())
		be.Commit()
		Expect(isSuccessful(tx)).To(BeTrue())

	}

	oraclizeMockAddress, tx, oraclizeMock, err = mocks.DeployOraclize(bankAccount.TransactOpts(), be, bankAccount.Address())
	Expect(err).ToNot(HaveOccurred())
	be.Commit()
	Expect(isSuccessful(tx)).To(BeTrue())

	oraclizeMockAddrResolverAddress, tx, oraclizeMockAddrResolver, err = mocks.DeployOraclizeAddrResolver(bankAccount.TransactOpts(), be, oraclizeMockAddress)
	Expect(err).ToNot(HaveOccurred())
	be.Commit()
	Expect(isSuccessful(tx)).To(BeTrue())

	oracleAddress, tx, oracle, err = bindings.DeployOracle(
		bankAccount.TransactOpts(),
		be,
		oraclizeMockAddrResolverAddress,
		ensAddress,
		controllerName,
	)
	Expect(err).ToNot(HaveOccurred())
	be.Commit()
	Expect(isSuccessful(tx)).To(BeTrue())

	{

		// register oracle with ENS

		tx, err = ens.SetResolver(bankAccount.TransactOpts(), oracleName, ensResolverAddress)
		Expect(err).ToNot(HaveOccurred())
		be.Commit()
		Expect(isSuccessful(tx)).To(BeTrue())

		tx, err = ensResolver.SetAddr(bankAccount.TransactOpts(), oracleName, oracleAddress)
		Expect(err).ToNot(HaveOccurred())
		be.Commit()
		Expect(isSuccessful(tx)).To(BeTrue())

		Expect(bankAccount.Transfer(be, controller.Address(), ethToWei(1))).To(Succeed())

	}

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

func exponentiateDecimals(decimals uint8) *big.Int {
	base := big.NewInt(10)
	expDec := big.NewInt(1)
	for i := 0; i < 8; i++ {
		expDec.Mul(expDec, base)
	}
	return expDec
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

	tx, err = oracle.AddTokens(controller.TransactOpts(), []common.Address{tkna}, stringsToByte32("TKN"), []*big.Int{exponentiateDecimals(8)})
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
		ensAddress,
		oracleName,
		controllerName,
		ethToWei(100),
	)
	Expect(err).ToNot(HaveOccurred())
	be.Commit()
	Expect(isSuccessful(tx)).To(BeTrue())
})
