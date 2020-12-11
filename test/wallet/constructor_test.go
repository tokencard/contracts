package wallet_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/tokencard/contracts/v3/pkg/bindings"
	"github.com/tokencard/contracts/v3/pkg/bindings/externals/upgradeability"
	. "github.com/tokencard/contracts/v3/test/shared"
	"github.com/tokencard/ethertest"
)

var _ = Describe("wallet initialization", func() {

	It("Should deploy a new wallet", func() {
		RandomProxyAddress, tx, _, err := upgradeability.DeployUpgradeabilityProxy(
			BankAccount.TransactOpts(),
			Backend,
			WalletImplementationAddress,
			nil,
		)
		Expect(err).ToNot(HaveOccurred())
		Backend.Commit()
		Expect(isSuccessful(tx)).To(BeTrue())

		RandomProxy, err := bindings.NewWallet(RandomProxyAddress, Backend)
		tx, err = RandomProxy.InitializeWallet(BankAccount.TransactOpts(ethertest.WithGasLimit(7000000)), Owner.Address(), false, ENSRegistryAddress, TokenWhitelistNode, ControllerNode, LicenceNode, EthToWei(100))
		Expect(err).ToNot(HaveOccurred())
		Backend.Commit()
		Expect(isSuccessful(tx)).To(BeTrue())
		//returnData, _ := ethCall(tx)
		//Expect(string(returnData[len(returnData)-64:])).To(ContainSubstring("no stablecoin"))
	})

})
