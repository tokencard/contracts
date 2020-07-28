package gas_proxy_test

import (
	"context"
	"math/big"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/tokencard/contracts/v3/pkg/bindings"
	"github.com/tokencard/contracts/v3/pkg/bindings/mocks"
	. "github.com/tokencard/contracts/v3/test/shared"
)

var GasToken *mocks.GasToken
var GasTokenAddress common.Address
var GasBurner *mocks.GasBurner
var GasBurnerAddress common.Address
var GasProxy *bindings.GasProxy
var GasProxyAddress common.Address

func init() {
	TestRig.AddCoverageForContracts(
		"../../build/gasProxy/combined.json",
		"../../contracts",
	)
}

func ethCall(tx *types.Transaction) ([]byte, error) {
	msg, _ := tx.AsMessage(types.HomesteadSigner{})

	calMsg := ethereum.CallMsg{
		From:     msg.From(),
		To:       msg.To(),
		Gas:      msg.Gas(),
		GasPrice: msg.GasPrice(),
		Value:    msg.Value(),
		Data:     msg.Data(),
	}

	return Backend.CallContract(context.Background(), calMsg, nil)
}

func TestTokenWhitelistSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "TokenWhitelist Suite")
}

var _ = BeforeEach(func() {
	err := InitializeBackend()
	Expect(err).ToNot(HaveOccurred())

	var tx *types.Transaction

	GasTokenAddress, tx, GasToken, err = mocks.DeployGasToken(BankAccount.TransactOpts(), Backend)
	Expect(err).ToNot(HaveOccurred())
	Backend.Commit()
	Expect(isSuccessful(tx)).To(BeTrue())

	GasBurnerAddress, tx, GasBurner, err = mocks.DeployGasBurner(BankAccount.TransactOpts(), Backend)
	Expect(err).ToNot(HaveOccurred())
	Backend.Commit()
	Expect(isSuccessful(tx)).To(BeTrue())

	GasProxyAddress, tx, GasProxy, err = bindings.DeployGasProxy(BankAccount.TransactOpts(), Backend, ENSRegistryAddress, ControllerName)
	Expect(err).ToNot(HaveOccurred())
	Backend.Commit()
	Expect(isSuccessful(tx)).To(BeTrue())

	tx, err = GasProxy.SetGasToken(ControllerAdmin.TransactOpts(), GasTokenAddress, bindings.GasRefundableGasTokenParameters{
		FreeCallGasCost:  big.NewInt(14154),
		GasRefundPerUnit: big.NewInt(41130),
	})
	Expect(err).ToNot(HaveOccurred())
	Backend.Commit()
	Expect(isSuccessful(tx)).To(BeTrue())

	tx, err = ControllerContract.AddController(ControllerAdmin.TransactOpts(), GasProxyAddress)
	Expect(err).ToNot(HaveOccurred())
	Backend.Commit()
	Expect(isSuccessful(tx)).To(BeTrue())

	tx, err = GasToken.Mint(BankAccount.TransactOpts(), big.NewInt(20))
	Expect(err).ToNot(HaveOccurred())
	Backend.Commit()
	Expect(isSuccessful(tx)).To(BeTrue())

	tx, err = GasToken.Transfer(BankAccount.TransactOpts(), GasProxyAddress, big.NewInt(20))
	Expect(err).ToNot(HaveOccurred())
	Backend.Commit()
	Expect(isSuccessful(tx)).To(BeTrue())
})

var _ = AfterEach(func() {
	err := Backend.Close()
	Expect(err).ToNot(HaveOccurred())

})

var _ = AfterSuite(func() {
	TestRig.ExpectMinimumCoverage("gasProxy.sol", 36.2)
	TestRig.ExpectMinimumCoverage("internals/gasRefundable.sol", 76.2)
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
