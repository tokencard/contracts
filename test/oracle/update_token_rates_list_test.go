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

var _ = Describe("updateRatesList", func() {

	Context("When tokens are already supported", func() {
		BeforeEach(func() {
			tokens := []common.Address{common.HexToAddress("0x0"), common.HexToAddress("0x1"), common.HexToAddress("0x2")}
			tx, err := TokenWhitelist.AddTokens(
				Controller.TransactOpts(),
				tokens,
				StringsToByte32("BNT", "TKN", "DGX"),
				[]*big.Int{DecimalsToMagnitude(big.NewInt(18)), DecimalsToMagnitude(big.NewInt(8)), DecimalsToMagnitude(big.NewInt(18))},
				[]bool{false, true, false},
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
				tokenList := []common.Address{common.HexToAddress("0x0"), common.HexToAddress("0x2")}
				tx, err = Oracle.UpdateTokenRatesList(Controller.TransactOpts(ethertest.WithValue(big.NewInt(100000000))), big.NewInt(GAS_LIMIT), tokenList)
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
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
				Expect(evt.Symbol).To(Equal("DGX"))
			})
			It("Should transfer value to oracle contract", func() {
				b, err := Backend.BalanceAt(context.Background(), OracleAddress, nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(b.String()).NotTo(Equal("0"))
				newbalance := new(big.Int).Sub(big.NewInt(100000000), big.NewInt(2000000)) //10^6 is the mock price of the Oraclize query (x2 queries)
				Expect(b.String()).To(Equal(newbalance.String()))

				Expect(b.String()).To(Equal(newbalance.String()))
			})
		}) // called by the controller with value

		Context("When called by the controller with 0 value", func() {
			var tx *types.Transaction
			var err error
			BeforeEach(func() {
				tokenList := []common.Address{common.HexToAddress("0x0"), common.HexToAddress("0x2")}
				tx, err = Oracle.UpdateTokenRatesList(Controller.TransactOpts(), big.NewInt(GAS_LIMIT), tokenList)
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
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
			It("Oracle contract balance should be 0", func() {
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

	Context("When the token list is empty", func() {

		Context("When called by the controller with value", func() {
			var tx *types.Transaction
			var err error
			BeforeEach(func() {
				tokenList := []common.Address{}
				tx, err = Oracle.UpdateTokenRatesList(Controller.TransactOpts(ethertest.WithValue(big.NewInt(100000000))), big.NewInt(GAS_LIMIT), tokenList)
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isGasExhausted(tx, 100000000)).To(BeFalse())
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
				Expect(evt.Reason).To(Equal("empty token list"))
			})
		})
	})

}) //updateRates
