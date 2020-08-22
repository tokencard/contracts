package wallet_test

import (
	"github.com/ethereum/go-ethereum/common"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/tokencard/contracts/v3/pkg/bindings"
	"github.com/tokencard/contracts/v3/pkg/bindings/externals/upgradeability"
	. "github.com/tokencard/contracts/v3/test/shared"
	"github.com/tokencard/ethertest"
)

var _ = Describe("wallet initialization", func() {

	Context("the stablecoin should be in the whitelist", func() {
		It("Should update the tokens map", func() {
			symbol, _, _, available, loadable, _, _, err := TokenWhitelist.GetTokenInfo(nil, StablecoinAddress)
			Expect(err).ToNot(HaveOccurred())
			Expect(symbol).To(Equal("USDC"))
			Expect(available).To(BeTrue())
			Expect(loadable).To(BeTrue())
		})
	})

	When("the stablecoing is removed from the list", func() {

		BeforeEach(func() {
			tx, err := TokenWhitelist.RemoveTokens(ControllerAdmin.TransactOpts(), []common.Address{StablecoinAddress})
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeTrue())
		})

		It("Should NOT initialize a new wallet porxy", func() {
			RandomProxyAddress, tx, _, err := upgradeability.DeployUpgradeabilityProxy(BankAccount.TransactOpts(), Backend, WalletImplementationAddress, nil)
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeTrue())

			RandomProxy, err := bindings.NewWallet(RandomProxyAddress, Backend)
			tx, err = RandomProxy.InitializeWallet(BankAccount.TransactOpts(ethertest.WithGasLimit(7000000)), Owner.Address(), false, ENSRegistryAddress, TokenWhitelistNode, ControllerNode, LicenceNode, EthToWei(100))
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeFalse())
			returnData, _ := ethCall(tx)
			Expect(string(returnData[len(returnData)-64:])).To(ContainSubstring("no stablecoin"))
		})

	})

})
