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

var _ = Describe("Oracle claim", func() {

	It("the initial balance of the Oracle contract should be zero", func() {
		b, e := Backend.BalanceAt(context.Background(), OracleAddress, nil)
		Expect(e).ToNot(HaveOccurred())
		Expect(b.String()).To(Equal("0"))
	})

	It("the initial balance of the Float contract should be zero", func() {
		b, e := Backend.BalanceAt(context.Background(), CryptoFloatAddress, nil)
		Expect(e).ToNot(HaveOccurred())
		Expect(b.String()).To(Equal("0"))
	})

	Context("When the balance of the Oracle contract is increased by sending money through the updateRates function", func() {

		BeforeEach(func() {
			tokens := []common.Address{common.HexToAddress("0x0"), common.HexToAddress("0x1")}
			tx, err := TokenWhitelist.AddTokens(
				ControllerAdmin.TransactOpts(),
				tokens,
				StringsToByte32("BNT", "TKN"),
				[]*big.Int{DecimalsToMagnitude(big.NewInt(18)), DecimalsToMagnitude(big.NewInt(8))},
				[]bool{false, true},
				[]bool{false, true},
				big.NewInt(20180913153211),
			)
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeTrue())
		})

		BeforeEach(func() {
			tx, err := Oracle.UpdateTokenRates(Controller.TransactOpts(ethertest.WithValue(EthToWei(2))), big.NewInt(gasLimit))
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeTrue())
		})

		It("the  balance of the Oracle contract should increase by 2 ETH", func() {
			b, e := Backend.BalanceAt(context.Background(), OracleAddress, nil)
			Expect(e).ToNot(HaveOccurred())
			newbalance := new(big.Int).Sub(EthToWei(2), big.NewInt(2000000)) //10^6 is the mock price of the Oraclize query (x2 queries)
			Expect(b.String()).To(Equal(newbalance.String()))
		})

		Context("When called by the controller admin", func() {
			var tx *types.Transaction

			BeforeEach(func() {
				var err error
				tx, err = Oracle.Claim(ControllerAdmin.TransactOpts(), CryptoFloatAddress, common.HexToAddress("0x0"), EthToWei(1))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
			})

			It("Should succeed", func() {
				Expect(isSuccessful(tx)).To(BeTrue())
			})

			It("the  balance of the Oracle contract should decrease by 1 ETH", func() {
				b, e := Backend.BalanceAt(context.Background(), OracleAddress, nil)
				Expect(e).ToNot(HaveOccurred())
				newbalance := new(big.Int).Sub(EthToWei(2), big.NewInt(2000000)) //10^6 is the mock price of the Oraclize query (x2 queries)
				newbalance.Sub(newbalance, EthToWei(1))
				Expect(b.String()).To(Equal(newbalance.String()))
			})

			It("the  balance of the Float contract should increase by 1 ETH", func() {
				b, e := Backend.BalanceAt(context.Background(), CryptoFloatAddress, nil)
				Expect(e).ToNot(HaveOccurred())
				Expect(b.String()).To(Equal(EthToWei(1).String()))
			})

			It("Should emit a Claimed event", func() {
				it, err := Oracle.FilterClaimed(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(it.Next()).To(BeTrue())
				evt := it.Event
				Expect(it.Next()).To(BeFalse())
				Expect(evt.To).To(Equal(CryptoFloatAddress))
				Expect(evt.Asset).To(Equal(common.HexToAddress("0x0"))) //represents ETH
				Expect(evt.Amount.String()).To(Equal(EthToWei(1).String()))
			})
		})

		Context("When not called by the controller admin", func() {
			It("Should fail", func() {
				tx, err := Oracle.Claim(RandomAccount.TransactOpts(ethertest.WithGasLimit(100000)), CryptoFloatAddress, common.HexToAddress("0x0"), EthToWei(1))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isGasExhausted(tx, 100000)).To(BeFalse())
				Expect(isSuccessful(tx)).To(BeFalse())
			})
		})

	}) //Oracle balance: 2 ETH

})
