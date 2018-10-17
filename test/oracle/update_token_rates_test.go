package oracle_test

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tokencard/ethertest"
)

var _ = Describe("updateRates", func() {

	Context("When tokens are already supported", func() {
		BeforeEach(func() {
			tokens := []common.Address{common.HexToAddress("0x0"), common.HexToAddress("0x1")}
			tx, err := oracle.AddTokens(
				controller.TransactOpts(),
				tokens,
				stringsToByte32("BNT", "TKN"),
				[]*big.Int{calculateMagnitude(big.NewInt(18)), calculateMagnitude(big.NewInt(8))},
				big.NewInt(20180913153211),
			)
			Expect(err).ToNot(HaveOccurred())
			be.Commit()
			Expect(isSuccessful(tx)).To(BeTrue())
		})

		Context("When called by the controller with value", func() {
			var tx *types.Transaction
			var err error
			BeforeEach(func() {
				tx, err = oracle.UpdateTokenRates(controller.TransactOpts(WithValue(big.NewInt(100000000))))
				Expect(err).ToNot(HaveOccurred())
				be.Commit()
			})
			It("Should succeed", func() {
				Expect(isSuccessful(tx)).To(BeTrue())
			})
			It("Should emit an Requested Update event", func() {
				it, err := oracle.FilterRequestedUpdate(nil)
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
				b, err := be.BalanceAt(context.Background(), oracleAddress, nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(b.String()).NotTo(Equal("0"))
				Expect(b.String()).NotTo(Equal("100000000"))
			})
		}) // called by the controller with value

		Context("When called by the controller with 0 value", func() {
			var tx *types.Transaction
			var err error
			BeforeEach(func() {
				tx, err = oracle.UpdateTokenRates(controller.TransactOpts())
				Expect(err).ToNot(HaveOccurred())
				be.Commit()
			})
			It("Should succeed", func() {
				Expect(isSuccessful(tx)).To(BeTrue())
			})
			It("Should emit an Failed Update Request event", func() {
				it, err := oracle.FilterFailedUpdateRequest(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(it.Next()).To(BeTrue())
				evt := it.Event
				Expect(it.Next()).To(BeFalse())
				Expect(evt.Reason).To(Equal("zero balance"))
			})
			It("Oracle contract balance should be 0", func() {
				b, err := be.BalanceAt(context.Background(), oracleAddress, nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(b.String()).To(Equal("0"))
			})
		})

		Context("When called by a random address", func() {
			It("Should fail", func() {
				tx, err := oracle.UpdateTokenRates(randomAccount.TransactOpts(WithGasLimit(100000)))
				Expect(err).ToNot(HaveOccurred())
				be.Commit()
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
				tx, err = oracle.UpdateTokenRates(controller.TransactOpts(WithValue(big.NewInt(100000000))))
				Expect(err).ToNot(HaveOccurred())
				be.Commit()
			})
			It("Should succeed", func() {
				Expect(isSuccessful(tx)).To(BeTrue())
			})
			It("Should not make any query", func() {
				b, err := be.BalanceAt(context.Background(), oracleAddress, nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(b.String()).To(Equal("100000000"))
			})
		})
	})

}) //updateRates
