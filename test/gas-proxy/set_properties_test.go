package gas_proxy_test

import (
	"math/big"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/tokencard/contracts/v3/pkg/bindings"
	. "github.com/tokencard/contracts/v3/test/shared"
)

var _ = Describe("SetProperties", func() {
	It("the initial value of the _gasToken property should be matching the initial GasToken", func() {
		gasTokenAddress, err := GasProxy.GasToken(nil)
		Expect(err).ToNot(HaveOccurred())
		Expect(gasTokenAddress.Hex()).To(Equal(GasTokenAddress.Hex()))
	})

	It("the initial value of the _gasTokenParameters property should be set to the initial values", func() {
		gasTokenParameters, err := GasProxy.GasTokenParameters(nil)
		Expect(err).ToNot(HaveOccurred())
		Expect(gasTokenParameters.FreeCallGasCost.String()).To(Equal("14154"))
		Expect(gasTokenParameters.GasRefundPerUnit.String()).To(Equal("35000"))
	})

	When("the _gasToken and _gasTokenParameters are set to new values by the admin", func() {
		BeforeEach(func() {
			tx, err := GasProxy.SetGasToken(ControllerAdmin.TransactOpts(), RandomAccount.Address(), bindings.GasRefundableGasTokenParameters{
				FreeCallGasCost:  big.NewInt(12345),
				GasRefundPerUnit: big.NewInt(55555),
			})
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeTrue())
		})

		It("should change the value of the _gasToken and _gasTokenParameters to that of the new values", func() {
			gasTokenAddress, err := GasProxy.GasToken(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(gasTokenAddress.Hex()).To(Equal(RandomAccount.Address().Hex()))
			gasTokenParameters, err := GasProxy.GasTokenParameters(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(gasTokenParameters.FreeCallGasCost.String()).To(Equal("12345"))
			Expect(gasTokenParameters.GasRefundPerUnit.String()).To(Equal("55555"))
		})
	})
})
