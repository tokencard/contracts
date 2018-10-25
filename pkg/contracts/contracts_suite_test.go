package contracts_test

import (
	"context"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/tokencard/contracts/pkg/contracts"
	. "github.com/tokencard/contracts/test/shared"
)

func TestContractsSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Contract Suite")
}

var walletWrapper *contracts.Wallet
var walletWrapperAddress common.Address
var opts *contracts.ConstructOpts

var _ = BeforeEach(func() {
	err := InitializeBackend()
	Expect(err).ToNot(HaveOccurred())

	opts = &contracts.ConstructOpts{
		From:     BankAccount.Address(),
		Nonce:    0,
		Value:    nil,
		GasPrice: nil,
		GasLimit: 0,
		Sign:     signTransaction,
		Context:  context.Background(),
	}

	var signedTx *types.Transaction
	walletWrapperAddress, signedTx, err = contracts.DeployWallet(opts, Backend, BankAccount.Address(), false, ENSRegistryAddress, OracleName, ControllerName, EthToWei(100))
	Expect(err).ToNot(HaveOccurred())

	Expect(Backend.SendTransaction(context.Background(), signedTx)).To(Succeed())
	Backend.Commit()
	Expect(isSuccessful(signedTx)).To(BeTrue())

	walletWrapper, err = contracts.NewWallet(Backend, walletWrapperAddress)
	Expect(err).ToNot(HaveOccurred())
})

func signTransaction(_ context.Context, tx *types.Transaction) (*types.Transaction, error) {
	return BankAccount.SignTransaction(Backend, tx)
}

func isSuccessful(tx *types.Transaction) bool {
	r, err := Backend.TransactionReceipt(context.Background(), tx.Hash())
	Expect(err).ToNot(HaveOccurred())
	return r.Status == types.ReceiptStatusSuccessful
}
