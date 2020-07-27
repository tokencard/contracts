package gas_proxy_test

import (
	"math/big"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tokencard/contracts/v3/test/shared"
)

var _ = Describe("SetProperties", func() {
	It("the initial value of the _gasToken property should be matching the GasToken address", func() {
		gasTokenAddress, err := GasProxy.GasToken(nil)
		Expect(err).ToNot(HaveOccurred())
		Expect(gasTokenAddress.Hex()).To(Equal(GasTokenAddress.Hex()))
	})

	It("the initial value of the _freeCallGasCost property should be set to 14154", func() {
		freeCallGasCost, err := GasProxy.FreeCallGasCost(nil)
		Expect(err).ToNot(HaveOccurred())
		Expect(freeCallGasCost.String()).To(Equal("14154"))
	})

	It("the initial value of the _gasRefundPerUnit property should be set to 41130", func() {
		gasRefundPerUnit, err := GasProxy.GasRefundPerUnit(nil)
		Expect(err).ToNot(HaveOccurred())
		Expect(gasRefundPerUnit.String()).To(Equal("41130"))
	})

	When("the _gasToken property is set to a new gas token address by the controller", func() {
		BeforeEach(func() {
			tx, err := GasProxy.SetGasToken(Controller.TransactOpts(), RandomAccount.Address())
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeTrue())
		})

		It("should change the value of the _gasToken property to that of the new gas token address", func() {
			gasTokenAddress, err := GasProxy.GasToken(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(gasTokenAddress.Hex()).To(Equal(RandomAccount.Address().Hex()))
		})
	})

	When("the _freeCallGasCost property is set to a new value by the controller", func() {
		BeforeEach(func() {
			tx, err := GasProxy.SetFreeCallGasCost(Controller.TransactOpts(), big.NewInt(12345))
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeTrue())
		})

		It("should change the value of the _freeCallGasCost property to that of the new value", func() {
			freeCallGasCost, err := GasProxy.FreeCallGasCost(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(freeCallGasCost.String()).To(Equal("12345"))
		})

	})

	When("the _gasRefundPerUnit property is set to a new value by the controller", func() {
		BeforeEach(func() {
			tx, err := GasProxy.SetGasRefundPerUnit(Controller.TransactOpts(), big.NewInt(55555))
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeTrue())
		})

		It("should change the value of the _gasRefundPerUnit property to that of the new value", func() {
			gasRefundPerUnit, err := GasProxy.GasRefundPerUnit(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(gasRefundPerUnit.String()).To(Equal("55555"))
		})

	})

})
