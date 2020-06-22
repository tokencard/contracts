package upgrade_test

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/i-stam/ethertest"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tokencard/contracts/v3/test/shared"
)

var _ = Describe("change Admin", func() {

	It("should have the wallet owner as the admin", func() {
		admin, err := Proxy.Admin(nil)
		Expect(err).ToNot(HaveOccurred())
		Expect(admin).To(Equal(Owner.Address()))
	})

	It("Should fail when trying to set it to address 0x0", func() {
		tx, err := Proxy.ChangeAdmin(Owner.TransactOpts(ethertest.WithGasLimit(100000)), common.HexToAddress("0x0"))
		Expect(err).ToNot(HaveOccurred())
		Backend.Commit()
		Expect(isSuccessful(tx)).To(BeFalse())
		returnData, _ := ethCall(tx)
		Expect(string(returnData[len(returnData)-64:])).To(ContainSubstring("Cannot change the admin of a proxy to address 0"))
	})

	When("trying to set a new admin and the wallet is transferable", func() {
		BeforeEach(func() {
			tx, err := Proxy.ChangeAdmin(Owner.TransactOpts(), RandomAccount.Address())
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeTrue())
		})

		It("should set the Random account as the admin", func() {
			admin, err := Proxy.Admin(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(admin).To(Equal(RandomAccount.Address()))
		})

	})

	When("trying to set a new admin and the wallet is not transferable", func() {

		// transferOwnership and set transferable to FALSE
		BeforeEach(func() {
			tx, err := ProxyWallet.TransferOwnership(Owner.TransactOpts(), RandomAccount.Address(), false)
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeTrue())
		})

		It("Should fail when trying to set it to address 0x0", func() {
			tx, err := Proxy.ChangeAdmin(Owner.TransactOpts(ethertest.WithGasLimit(100000)), common.HexToAddress("0x0"))
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeFalse())
			returnData, _ := ethCall(tx)
			Expect(string(returnData[len(returnData)-64:])).To(ContainSubstring("Cannot change the admin of a proxy to address 0"))
		})

		It("Should fail when trying to set it to a new address", func() {
			tx, err := Proxy.ChangeAdmin(Owner.TransactOpts(ethertest.WithGasLimit(100000)), RandomAccount.Address())
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeFalse())
			returnData, _ := ethCall(tx)
			Expect(string(returnData[len(returnData)-64:])).To(ContainSubstring("non-transferable proxy admin"))
		})

	})

})
