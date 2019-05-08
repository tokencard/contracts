package licence_test

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/tokencard/contracts/pkg/bindings"
	. "github.com/tokencard/contracts/test/shared"
	"github.com/tokencard/ethertest"
)

// var DAO *bindings.Dao
// var DAOAddress common.Address

var DAO *ethertest.Account

func init() {
	TestRig.AddCoverageForContracts(
		"../../build/licence/combined.json",
		"../../contracts")
}

func TestLicenceSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Contract Suite")
}

var _ = BeforeEach(func() {
	err := InitializeBackend()
	Expect(err).ToNot(HaveOccurred())

	var tx *types.Transaction

	DAO = ethertest.NewAccount()
	err = BankAccount.Transfer(Backend, DAO.Address(), EthToWei(50))

	Expect(err).ToNot(HaveOccurred())

	LicenceAddress, tx, Licence, err = bindings.DeployLicence(BankAccount.TransactOpts(), Backend, Owner.Address(), true, big.NewInt(10), CryptoFloatAddress, TokenHolderAddress, common.HexToAddress("0x0"))
	Expect(err).ToNot(HaveOccurred())
	Backend.Commit()
	Expect(isSuccessful(tx)).To(BeTrue())

})

var _ = Describe("constructor is called with an out of range licence value", func() {
		It("should fail", func() {
			_, tx, _, err := bindings.DeployLicence(BankAccount.TransactOpts(ethertest.WithGasLimit(2000000)), Backend, Owner.Address(), true, big.NewInt(0), CryptoFloatAddress, TokenHolderAddress, common.HexToAddress("0x0"))
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeFalse())
		})

		It("should fail", func() {
			_, tx, _, err := bindings.DeployLicence(BankAccount.TransactOpts(ethertest.WithGasLimit(2000000)), Backend, Owner.Address(), true, big.NewInt(1001), CryptoFloatAddress, TokenHolderAddress, common.HexToAddress("0x0"))
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeFalse())
		})
})

var _ = AfterSuite(func() {
	TestRig.ExpectMinimumCoverage("licence.sol", 100.00)
	TestRig.PrintGasUsage(os.Stdout)
})

func isSuccessful(tx *types.Transaction) bool {
	r, err := Backend.TransactionReceipt(context.Background(), tx.Hash())
	Expect(err).ToNot(HaveOccurred())
	return r.Status == types.ReceiptStatusSuccessful
}

func isGasExhausted(tx *types.Transaction, gasLimit uint64) bool {
	r, err := Backend.TransactionReceipt(context.Background(), tx.Hash())
	Expect(err).ToNot(HaveOccurred())
	if r.Status == types.ReceiptStatusSuccessful {
		return false
	}
	return r.GasUsed == gasLimit
}

var _ = AfterEach(func() {
	td := CurrentGinkgoTestDescription()
	if td.Failed {
		fmt.Fprintf(GinkgoWriter, "\nLast Executed Smart Contract Line for %s:%d\n", td.FileName, td.LineNumber)
		fmt.Fprintln(GinkgoWriter, TestRig.LastExecuted())
	}
	err := Backend.Close()
	Expect(err).ToNot(HaveOccurred())
})
