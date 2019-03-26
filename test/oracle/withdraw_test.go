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

var _ = Describe("withdraw", func() {

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

		It("the oracle should have 0 balance", func() {
			b, err := Backend.BalanceAt(context.Background(), OracleAddress, nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(b.String()).To(Equal("0"))

		})

		It("the 0x2 address should have 0 balance", func() {
			b, err := Backend.BalanceAt(context.Background(), common.HexToAddress("0x2"), nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(b.String()).To(Equal("0"))
		})

		Context("When called by the controller with value", func() {
			var tx *types.Transaction
			var err error

			BeforeEach(func() {
				tokenList := []common.Address{common.HexToAddress("0x0"), common.HexToAddress("0x2")}
				tx, err = Oracle.UpdateTokenRatesList(Controller.TransactOpts(ethertest.WithValue(big.NewInt(100000000))), big.NewInt(gasLimit), tokenList)
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})

			BeforeEach(func() {
				tx, err = Oracle.Withdraw(Controller.TransactOpts(), common.HexToAddress("0x2"), big.NewInt(1000000))
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
				newBalance := new(big.Int).Sub(big.NewInt(100000000), big.NewInt(2000000)) //10^6 is the mock price of the Oraclize query (x2 queries)
				newBalance.Sub(newBalance, big.NewInt(1000000))                            //10^6 is the mock price of the Oraclize query (x2 queries)
				Expect(b.String()).To(Equal(newBalance.String()))
			})

			It("Should withdraw from the oracle and transfer 1000000 wei to 0x2 address ", func() {
				b, err := Backend.BalanceAt(context.Background(), common.HexToAddress("0x2"), nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(b.String()).To(Equal("1000000"))
			})
		}) // called by the controller with value

		Context("When called by a Random account", func() {
			var tx *types.Transaction
			var err error

			BeforeEach(func() {
				tokenList := []common.Address{common.HexToAddress("0x0"), common.HexToAddress("0x2")}
				tx, err = Oracle.UpdateTokenRatesList(Controller.TransactOpts(ethertest.WithValue(big.NewInt(100000000))), big.NewInt(gasLimit), tokenList)
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})

			It("Should fail", func() {
				tx, err = Oracle.Withdraw(RandomAccount.TransactOpts(ethertest.WithGasLimit(100000)), common.HexToAddress("0x2"), big.NewInt(1000000))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isGasExhausted(tx, 100000)).To(BeFalse())
				Expect(isSuccessful(tx)).To(BeFalse())
			})
		})

	}) //tokens supported

}) //updateRates
