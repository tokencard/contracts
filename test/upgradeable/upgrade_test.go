package upgrade_test

import (
	"github.com/i-stam/ethertest"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/tokencard/contracts/v3/pkg/bindings/mocks"
	. "github.com/tokencard/contracts/v3/test/shared"
)

var _ = Describe("upgrade implementation", func() {

	It("should fail when someone tries to re-initialize", func() {
		tx, err := ProxyWallet.InitializeWallet(BankAccount.TransactOpts(ethertest.WithGasLimit(100000)), RandomAccount.Address(), true, ENSRegistryAddress, TokenWhitelistName, ControllerName, LicenceName, EthToWei(1))
		Expect(err).ToNot(HaveOccurred())
		Backend.Commit()
		Expect(isSuccessful(tx)).To(BeFalse())
		returnData, _ := ethCall(tx)
		Expect(string(returnData[len(returnData)-64:])).To(ContainSubstring("Contract instance has already been initialized"))
	})

	It("set the correct wallet implementations", func() {
		imp, err := Proxy.Implementation(nil)
		Expect(err).ToNot(HaveOccurred())
		Expect(imp).To(Equal(WalletImplementationAddress))
	})

	It("should set the wallet owner as the admin", func() {
		admin, err := Proxy.Admin(nil)
		Expect(err).ToNot(HaveOccurred())
		Expect(admin).To(Equal(Owner.Address()))
	})

	When("we increase the nonce", func() {

		BeforeEach(func() {
			tx, err := ProxyWallet.IncreaseRelayNonce(Owner.TransactOpts())
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeTrue())
		})

		It("should increase the nonce", func() {
			rn, err := ProxyWallet.RelayNonce(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(rn.String()).To(Equal("1"))
		})

		When("we upgrade to a new mock imlementation", func() {

			BeforeEach(func() {
				WalletMockAddress, tx, _, err := mocks.DeployWalletMock(BankAccount.TransactOpts(), Backend)
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())

				tx, err = Proxy.UpgradeTo(Owner.TransactOpts(), WalletMockAddress)
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())

				tx, err = ProxyWallet.IncreaseRelayNonce(Owner.TransactOpts()) //new implementation: nonce += 2;
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})

			It("Should return the previous nonce + 2", func() {
				rn, err := ProxyWallet.RelayNonce(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(rn.String()).To(Equal("3"))
			})

		})
	})

})
