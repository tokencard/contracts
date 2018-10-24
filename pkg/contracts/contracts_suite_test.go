package contracts_test

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/tokencard/contracts/pkg/contracts"
	"github.com/tokencard/ethertest"
)

func TestContracts(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Contracts Suite")
}

var testRig = ethertest.NewTestRig()
var bankAccount = ethertest.NewAccount()

var _ = BeforeSuite(func() {
	testRig.AddGenesisAccountAllocation(bankAccount.Address(), ethToWei(200))
})

func ethToWei(amount int) *big.Int {
	r := big.NewInt(1000000000000000000)
	return r.Mul(r, big.NewInt(int64(amount)))
}
func finneyToWei(amount int) *big.Int {
	r := big.NewInt(1000000000000000)
	return r.Mul(r, big.NewInt(int64(amount)))
}

func gweiToWei(amount int) *big.Int {
	r := big.NewInt(1000000000)
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

var controllerContract *bindings.Controller
var controllerContractAddress common.Address

var _ = Describe("DeployWallet", func() {
	It("Deploys a wallet", func() {

		_, tx, err := contracts.DeployWallet(&contracts.ConstructOpts{}, be, bankAccount.Address(), false, common.Address{}, [32]byte{}, [32]byte{}, big.NewInt(10))
		Expect(err).ToNot(HaveOccurred())

		signed, err := bankAccount.SignTransaction(be, tx)
		Expect(err).ToNot(HaveOccurred())

		Expect(be.SendTransaction(context.Background(), signed)).To(Succeed())
		be.Commit()
		Expect(isSuccessful(signed)).To(BeTrue())

	})
})
