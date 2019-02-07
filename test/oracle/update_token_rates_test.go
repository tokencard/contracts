package oracle_test

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tokencard/contracts/test/shared"
	"github.com/tokencard/ethertest"
)

var _ = Describe("updateRates", func() {

	Context("When tokens are already supported", func() {
		BeforeEach(func() {
			tokens := []common.Address{common.HexToAddress("0x0"), common.HexToAddress("0x1")}
			tx, err := TokenWhitelist.AddTokens(
				Controller.TransactOpts(),
				tokens,
				StringsToByte32("BNT", "TKN"),
				[]*big.Int{DecimalsToMagnitude(big.NewInt(18)), DecimalsToMagnitude(big.NewInt(8))},
				[]bool{false, true},
				big.NewInt(20180913153211),
			)
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeTrue())
		})

		Context("When called by the controller with value", func() {
			var tx *types.Transaction
			var err error
			BeforeEach(func() {
				tx, err = Oracle.UpdateTokenRates(Controller.TransactOpts(ethertest.WithValue(big.NewInt(100000000))), big.NewInt(GAS_LIMIT))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
			})
			It("Should succeed", func() {
				Expect(isSuccessful(tx)).To(BeTrue())
			})
			It("Should emit an Requested Update event", func() {
				it, err := Oracle.FilterRequestedUpdate(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(it.Next()).To(BeTrue())
				evt := it.Event
				Expect(it.Next()).To(BeTrue())
				Expect(evt.Symbol).To(Equal("BNT"))
				evt = it.Event
				Expect(it.Next()).To(BeFalse())
				Expect(evt.Symbol).To(Equal("TKN"))
			})
			It("Should transfer value to oracle contract", func() {
				b, err := Backend.BalanceAt(context.Background(), OracleAddress, nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(b.String()).NotTo(Equal("0"))
				Expect(b.String()).NotTo(Equal("100000000"))
			})
		}) // called by the controller with value

		Context("When called by the controller with 0 value", func() {
			var tx *types.Transaction
			var err error
			BeforeEach(func() {
				tx, err = Oracle.UpdateTokenRates(Controller.TransactOpts(), big.NewInt(GAS_LIMIT))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
			})
			It("Should succeed", func() {
				Expect(isSuccessful(tx)).To(BeTrue())
			})
			It("Should emit an Failed Update Request event", func() {
				it, err := Oracle.FilterFailedUpdateRequest(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(it.Next()).To(BeTrue())
				evt := it.Event
				Expect(it.Next()).To(BeFalse())
				Expect(evt.Reason).To(Equal("insufficient balance"))
			})
			It("Oracle contract balance should Backend 0", func() {
				b, err := Backend.BalanceAt(context.Background(), OracleAddress, nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(b.String()).To(Equal("0"))
			})
		})

		Context("When called by a random address", func() {
			It("Should fail", func() {
				tx, err := Oracle.UpdateTokenRates(RandomAccount.TransactOpts(ethertest.WithGasLimit(100000)), big.NewInt(GAS_LIMIT))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isGasExhausted(tx, 100000)).To(BeFalse())
				Expect(isSuccessful(tx)).To(BeFalse())
			})
		})

	}) //tokens supported

	Context("When no tokens exist", func() {

		Context("When called by the controller with value", func() {
			var tx *types.Transaction
			var err error
			BeforeEach(func() {
				tx, err = Oracle.UpdateTokenRates(Controller.TransactOpts(ethertest.WithValue(big.NewInt(100000000))), big.NewInt(GAS_LIMIT))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
			})
			It("Should succeed", func() {
				Expect(isSuccessful(tx)).To(BeTrue())
			})
			It("Should not make any query", func() {
				b, err := Backend.BalanceAt(context.Background(), OracleAddress, nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(b.String()).To(Equal("100000000"))
			})
			It("Should emit an Failed Update Request event", func() {
				it, err := Oracle.FilterFailedUpdateRequest(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(it.Next()).To(BeTrue())
				evt := it.Event
				Expect(it.Next()).To(BeFalse())
				Expect(evt.Reason).To(Equal("no tokens"))
			})
		})
	})

}) //updateRates
