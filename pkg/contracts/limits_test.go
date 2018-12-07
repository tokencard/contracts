package contracts_test

import (
	"context"

	"github.com/ethereum/go-ethereum/common/hexutil"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tokencard/contracts/test/shared"
)

var _ = Describe("Wallet wrapper limits", func() {
	Context("When the spend limit has been initialized to one ETH", func() {
		BeforeEach(func() {
			signedTx, err := walletWrapper.InitializeSpendLimit(opts, EthToWei(1))
			Expect(err).ToNot(HaveOccurred())
			Expect(Backend.SendTransaction(context.Background(), signedTx)).To(Succeed())
			Backend.Commit()
			Expect(isSuccessful(signedTx)).To(BeTrue())
		})
		It("should have set the spend limit to one ETH", func() {
			spendLimit, err := walletWrapper.SpendLimit(context.Background(), nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(spendLimit).To(Equal(EthToWei(1)))
		})
		Context("When the SetSpendLimit events have been filtered", func() {
			It("should return a single SetSpendLimit event with the correct data", func() {
				events, err := walletWrapper.SetSpendLimitEvents(context.Background(), nil, nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(events).To(HaveLen(1))
				Expect(events[0].Data).To(HaveLen(2))
				Expect(hexutil.Encode(events[0].Data[0])).To(Equal(BankAccount.Address().Hash().Hex()))
				Expect(hexutil.Encode(events[0].Data[1])).To(Equal("0x0000000000000000000000000000000000000000000000000de0b6b3a7640000"))
			})
		})
	})
	Context("When the top up limit has been initialized to two hundred Finney", func() {
		BeforeEach(func() {
			signedTx, err := walletWrapper.InitializeTopUpLimit(opts, FinneyToWei(200))
			Expect(err).ToNot(HaveOccurred())
			Expect(Backend.SendTransaction(context.Background(), signedTx)).To(Succeed())
			Backend.Commit()
			Expect(isSuccessful(signedTx)).To(BeTrue())
		})
		It("should have set the top up limit to two hundred Finney", func() {
			spendLimit, err := walletWrapper.TopUpLimit(context.Background(), nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(spendLimit).To(Equal(FinneyToWei(200)))
		})
		Context("When the SetTopUpLimit events have been filtered", func() {
			It("should return a single SetTopUpLimit event with the correct data", func() {
				events, err := walletWrapper.SetTopUpLimitEvents(context.Background(), nil, nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(events).To(HaveLen(1))
				Expect(events[0].Data).To(HaveLen(2))
				Expect(hexutil.Encode(events[0].Data[0])).To(Equal(BankAccount.Address().Hash().Hex()))
				Expect(hexutil.Encode(events[0].Data[1])).To(Equal("0x00000000000000000000000000000000000000000000000002c68af0bb140000"))
			})
		})
	})
})
