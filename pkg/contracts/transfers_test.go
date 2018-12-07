package contracts_test

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tokencard/contracts/test/shared"
)

var _ = Describe("Wallet wrapper transfers", func() {
	Context("When the wallet receives one ETH", func() {
		var balance *big.Int
		BeforeEach(func() {
			err := BankAccount.Transfer(Backend, walletWrapperAddress, EthToWei(1))
			Expect(err).ToNot(HaveOccurred())
		})
		It("should increase the balance of the wallet by one ETH", func() {
			var err error
			balance, err = walletWrapper.Balance(context.Background(), nil, common.HexToAddress("0x0"))
			Expect(err).ToNot(HaveOccurred())
			Expect(balance).To(Equal(EthToWei(1)))
		})
		Context("When the Received events have been filtered", func() {
			It("should return a single Received event with the correct data", func() {
				events, err := walletWrapper.ReceivedEvents(context.Background(), nil, nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(events).To(HaveLen(1))
				Expect(events[0].Data).To(HaveLen(2))
				Expect(hexutil.Encode(events[0].Data[0])).To(Equal(BankAccount.Address().Hash().Hex()))
				Expect(hexutil.Encode(events[0].Data[1])).To(Equal("0x0000000000000000000000000000000000000000000000000de0b6b3a7640000"))
			})
		})
		Context("When two hundred Finney worth of gas has been topped up", func() {
			BeforeEach(func() {
				signedTx, err := walletWrapper.TopUpGas(opts, FinneyToWei(200))
				Expect(err).ToNot(HaveOccurred())
				Expect(Backend.SendTransaction(context.Background(), signedTx)).To(Succeed())
				Backend.Commit()
				Expect(isSuccessful(signedTx)).To(BeTrue())
			})
			It("should have decreased the wallet balance by two hundred Finney", func() {
				newBalance, err := walletWrapper.Balance(context.Background(), nil, common.HexToAddress("0x0"))
				Expect(err).ToNot(HaveOccurred())
				Expect(new(big.Int).Sub(balance, newBalance)).To(Equal(FinneyToWei(200)))
			})
			Context("When the ToppedUpGas events have been filtered", func() {
				It("should return a single ToppedUpGas event with the correct data", func() {
					events, err := walletWrapper.ToppedUpGasEvents(context.Background(), nil, nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(events).To(HaveLen(1))
					Expect(events[0].Data).To(HaveLen(3))
					Expect(hexutil.Encode(events[0].Data[0])).To(Equal(BankAccount.Address().Hash().Hex()))
					Expect(hexutil.Encode(events[0].Data[1])).To(Equal(BankAccount.Address().Hash().Hex()))
					Expect(hexutil.Encode(events[0].Data[2])).To(Equal("0x00000000000000000000000000000000000000000000000002c68af0bb140000"))
				})
			})
		})
		Context("When two hundred Finney has been transferred to a random account", func() {
			BeforeEach(func() {
				signedTx, err := walletWrapper.Transfer(opts, RandomAccount.Address(), common.HexToAddress("0x0"), FinneyToWei(200))
				Expect(err).ToNot(HaveOccurred())
				Expect(Backend.SendTransaction(context.Background(), signedTx)).To(Succeed())
				Backend.Commit()
				Expect(isSuccessful(signedTx)).To(BeTrue())
			})
			It("should have decreased the wallet balance by two hundred Finney", func() {
				newBalance, err := walletWrapper.Balance(context.Background(), nil, common.HexToAddress("0x0"))
				Expect(err).ToNot(HaveOccurred())
				Expect(new(big.Int).Sub(balance, newBalance)).To(Equal(FinneyToWei(200)))
			})
			Context("When the Transferred events have been filtered", func() {
				It("should return a single Transferred event with the correct data", func() {
					events, err := walletWrapper.TransferredEvents(context.Background(), nil, nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(events).To(HaveLen(1))
					Expect(events[0].Data).To(HaveLen(3))
					Expect(hexutil.Encode(events[0].Data[0])).To(Equal(RandomAccount.Address().Hash().Hex()))
					Expect(hexutil.Encode(events[0].Data[1])).To(Equal(common.HexToAddress("0x0").Hash().Hex()))
					Expect(hexutil.Encode(events[0].Data[2])).To(Equal("0x00000000000000000000000000000000000000000000000002c68af0bb140000"))
				})
			})
		})
	})
})
