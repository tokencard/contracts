package upgrade_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/tokencard/contracts/v3/pkg/bindings/mocks"
	. "github.com/tokencard/contracts/v3/test/shared"
)

var _ = Describe("upgrade implementation", func() {

	It("should have spend limit of 1 ETH", func() {
		sl, err := ProxyWallet.SpendLimitValue(nil)
		Expect(err).ToNot(HaveOccurred())
		Expect(sl.String()).To(Equal(EthToWei(1).String()))
	})

	It("should have an owner", func() {
		o, err := ProxyWallet.Owner(nil)
		Expect(err).ToNot(HaveOccurred())
		Expect(o).To(Equal(Owner.Address()))
	})

	It("should have an owner", func() {
		rn, err := ProxyWallet.RelayNonce(nil)
		Expect(err).ToNot(HaveOccurred())
		Expect(rn.String()).To(Equal("0"))
	})

	When("we upgrade to a new mock imlementation", func() {

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
			})

			It("Should return the number of the beast", func() {
				value, err := ProxyWallet.SpendLimitValue(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(value.String()).To(Equal("666"))
			})

		})
	})

})
