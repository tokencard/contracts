package wallet_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/tokencard/contracts/v3/pkg/bindings"
	. "github.com/tokencard/contracts/v3/test/shared"
)

var _ = Describe("wallet constructor", func() {

	It("Should deploy a new wallet", func() {
		_, tx, _, err := bindings.DeployWallet(
			BankAccount.TransactOpts(),
			Backend,
			Owner.Address(),
			true,
			ENSRegistryAddress,
			TokenWhitelistNode,
			ControllerNode,
			LicenceNode,
			WalletDeployerNode,
			GweiToWei(15),
		)
		Expect(err).ToNot(HaveOccurred())
		Backend.Commit()
		Expect(isSuccessful(tx)).To(BeTrue())
	})

})
