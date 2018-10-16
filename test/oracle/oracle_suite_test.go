package oracle_test

import (
	"context"
	"math/big"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/tokencard/contracts/pkg/bindings"
	"github.com/tokencard/contracts/pkg/bindings/external"
	"github.com/tokencard/contracts/pkg/bindings/mocks"
	"github.com/tokencard/ethertest"
)

func TestOracleSuite(t *testing.T) {
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

func ethToWei(amount int) *big.Int {
	r := big.NewInt(1000000000000000000)
	return r.Mul(r, big.NewInt(int64(amount)))
}

func finneyToWei(amount int) *big.Int {
	r := big.NewInt(1000000000000000)
	return r.Mul(r, big.NewInt(int64(amount)))
}

var testRig = ethertest.NewTestRig()
var controller = ethertest.NewAccount()
var randomAccount = ethertest.NewAccount()
var oraclizeCallbackAccount = ethertest.NewAccount()

var _ = BeforeSuite(func() {
	testRig.AddGenesisAccountAllocation(controller.Address(), ethToWei(1000))
	testRig.AddGenesisAccountAllocation(randomAccount.Address(), ethToWei(1000))
	testRig.AddGenesisAccountAllocation(oraclizeCallbackAccount.Address(), ethToWei(1000))
	testRig.AddCoverageForContracts("../../build/oracle/combined.json", "../../contracts/oracle.sol")
})

var _ = AfterSuite(func() {
	testRig.ExpectMinimumCoverage("oracle.sol:Oracle", 95.0)
	testRig.PrintGasUsage(os.Stdout)
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

func stringsToByte32(names ...string) [][32]byte {
	r := [][32]byte{}
	for _, n := range names {
		nb := [32]byte{}
		copy(nb[:], []byte(n))
		r = append(r, nb)
	}
	return r
}

func calculateMagnitude(decimals *big.Int) *big.Int {

	return new(big.Int).Exp(big.NewInt(10), decimals, nil)

}

var oraclizeMockAddrResolver *mocks.OraclizeAddrResolver
var oraclizeMockAddrResolverAddress common.Address

var oraclizeMock *mocks.Oraclize
var oraclizeMockAddress common.Address

var oracle *bindings.Oracle
var oracleAddress common.Address

var controllerContract *bindings.Controller
var controllerContractAddress common.Address

var ensResolver *bindings.Resolver
var ensResolverAddress common.Address

var ensAddress common.Address
var ens *external.ENSRegistry

var _ = BeforeEach(func() {
	be = testRig.NewTestBackend(ethertest.WithBlockchainTime(time.Date(2018, 9, 13, 15, 10, 0, 0, time.Local)))

	var err error
	var tx *types.Transaction

	ensAddress, tx, ens, err = external.DeployENSRegistry(controller.TransactOpts(), be)
	Expect(err).ToNot(HaveOccurred())
	be.Commit()
	Expect(isSuccessful(tx)).To(BeTrue())

	tx, err = ens.SetSubnodeOwner(controller.TransactOpts(), ensNode(""), labelHash("eth"), controller.Address())
	Expect(err).ToNot(HaveOccurred())
	be.Commit()
	Expect(isSuccessful(tx)).To(BeTrue())

	tx, err = ens.SetSubnodeOwner(controller.TransactOpts(), ensNode("eth"), labelHash("tokencard"), controller.Address())
	Expect(err).ToNot(HaveOccurred())
	be.Commit()
	Expect(isSuccessful(tx)).To(BeTrue())

	tx, err = ens.SetSubnodeOwner(controller.TransactOpts(), ensNode("tokencard.eth"), labelHash("controller"), controller.Address())
	Expect(err).ToNot(HaveOccurred())
	be.Commit()
	Expect(isSuccessful(tx)).To(BeTrue())

	ensResolverAddress, tx, ensResolver, err = bindings.DeployResolver(controller.TransactOpts(), be, ensAddress)
	Expect(err).ToNot(HaveOccurred())
	be.Commit()
	Expect(isSuccessful(tx)).To(BeTrue())

	tx, err = ens.SetResolver(controller.TransactOpts(), ensNode("controller.tokencard.eth"), ensResolverAddress)
	Expect(err).ToNot(HaveOccurred())
	be.Commit()
	Expect(isSuccessful(tx)).To(BeTrue())

	controllerContractAddress, tx, controllerContract, err = bindings.DeployController(controller.TransactOpts(), be, controller.Address())
	Expect(err).ToNot(HaveOccurred())
	be.Commit()
	Expect(isSuccessful(tx)).To(BeTrue())

	tx, err = ensResolver.SetAddr(controller.TransactOpts(), ensNode("controller.tokencard.eth"), controllerContractAddress)
	Expect(err).ToNot(HaveOccurred())
	be.Commit()
	Expect(isSuccessful(tx)).To(BeTrue())

	oraclizeMockAddress, tx, oraclizeMock, err = mocks.DeployOraclize(controller.TransactOpts(), be, oraclizeCallbackAccount.Address())
	Expect(err).ToNot(HaveOccurred())
	be.Commit()
	Expect(isSuccessful(tx)).To(BeTrue())

	oraclizeMockAddrResolverAddress, tx, _, err = mocks.DeployOraclizeAddrResolver(controller.TransactOpts(), be, oraclizeMockAddress)
	Expect(err).ToNot(HaveOccurred())
	be.Commit()
	Expect(isSuccessful(tx)).To(BeTrue())

	oracleAddress, tx, oracle, err = bindings.DeployOracle(
		controller.TransactOpts(),
		be,
		oraclizeMockAddrResolverAddress,
		ensAddress,
		ensNode("controller.tokencard.eth"),
	)
	Expect(err).ToNot(HaveOccurred())
	be.Commit()
	Expect(isSuccessful(tx)).To(BeTrue())

	isController, err := controllerContract.IsController(nil, controller.Address())
	Expect(err).ToNot(HaveOccurred())
	Expect(isController).To(BeTrue())
})
