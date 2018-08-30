package oracle_v2_test

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("updateRateManual", func() {

	Context("When the token is already supported", func() {
		BeforeEach(func() {
			tx, err := oracle.AddToken(controllerWallet.TransactOpts(), common.HexToAddress("0x1"), "ETH", 8)
			Expect(err).ToNot(HaveOccurred())
			be.Commit()
			Expect(isSuccessful(tx)).To(BeTrue())
		})
		Context("When called by the controller", func() {
			It("Should not fail", func() {
				tx, err := oracle.UpdateRateManual(controllerWallet.TransactOpts(), common.HexToAddress("0x1"), big.NewInt(666))
				Expect(err).ToNot(HaveOccurred())
				be.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})
		})
		Context("When not called by the controller", func() {
			It("Should fail", func() {
				to := randomWallet.TransactOpts()
				to.GasLimit = 100000
				tx, err := oracle.UpdateRateManual(to, common.HexToAddress("0x1"), big.NewInt(666))
				Expect(err).ToNot(HaveOccurred())
				be.Commit()
				Expect(isSuccessful(tx)).To(BeFalse())
			})
		})
	})

	Context("When the token is not supported", func() {
		Context("When called by the controller", func() {
			It("Should fail", func() {
				to := controllerWallet.TransactOpts()
				to.GasLimit = 100000
				tx, err := oracle.UpdateRateManual(to, common.HexToAddress("0x1"), big.NewInt(666))
				Expect(err).ToNot(HaveOccurred())
				be.Commit()
				Expect(isSuccessful(tx)).To(BeFalse())
			})
		})
		Context("When not called by the controller", func() {
			It("Should fail", func() {
				to := randomWallet.TransactOpts()
				to.GasLimit = 100000
				tx, err := oracle.UpdateRateManual(to, common.HexToAddress("0x1"), big.NewInt(666))
				Expect(err).ToNot(HaveOccurred())
				be.Commit()
				Expect(isSuccessful(tx)).To(BeFalse())
			})
		})
	})

})
