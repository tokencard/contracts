package wallet_test

import (
	"github.com/ethereum/go-ethereum/common"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/tokencard/contracts/v2/pkg/bindings"
	. "github.com/tokencard/contracts/v2/test/shared"
	"github.com/tokencard/ethertest"
)

var _ = Describe("wallet constructor", func() {

	Context("the stablecoin should be in the whitelist", func() {
		It("Should update the tokens map", func() {
			symbol, _, _, available, loadable, _, _, err := TokenWhitelist.GetTokenInfo(nil, StablecoinAddress)
			Expect(err).ToNot(HaveOccurred())
			Expect(symbol).To(Equal("DAI"))
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

		It("Should NOT deploy a new wallet", func() {
			_, tx, _, err := bindings.DeployWallet(
				BankAccount.TransactOpts(ethertest.WithGasLimit(7000000)),
				Backend,
				Owner.Address(),
				true,
				ENSRegistryAddress,
				TokenWhitelistName,
				ControllerName,
				LicenceName,
				EthToWei(100),
			)
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isGasExhausted(tx, 7000000)).To(BeFalse())
			Expect(isSuccessful(tx)).To(BeFalse())
			Expect(TestRig.LastExecuted()).To(MatchRegexp(`.*require\(stablecoinMagnitude > 0, "stablecoin not set"\);`))
		})

	})

})
