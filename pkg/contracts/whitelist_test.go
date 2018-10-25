package contracts_test

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/tokencard/contracts/pkg/contracts"
	. "github.com/tokencard/contracts/test/shared"
)

var _ = Describe("Wallet wrapper whitelist", func() {
	Context("When two addresses have been added to the whitelist", func() {
		BeforeEach(func() {
			signedTx, err := walletWrapper.InitializeWhitelist(opts, []common.Address{common.HexToAddress("0x123"), common.HexToAddress("0x124")})
			Expect(err).ToNot(HaveOccurred())
			Expect(Backend.SendTransaction(context.Background(), signedTx)).To(Succeed())
			Backend.Commit()
			Expect(isSuccessful(signedTx)).To(BeTrue())
		})

		It("should have two addresses in the whitelist", func() {
			whitelisted, err := walletWrapper.IsWhitelisted(context.Background(), nil, common.HexToAddress("0x123"))
			Expect(err).ToNot(HaveOccurred())
			Expect(whitelisted).To(BeTrue())

			whitelisted, err = walletWrapper.IsWhitelisted(context.Background(), nil, common.HexToAddress("0x124"))
			Expect(err).ToNot(HaveOccurred())
			Expect(whitelisted).To(BeTrue())
		})

		Context("When the AddedToWhitelist events have been filtered", func() {
			It("should return a single AddedToWhitelist event with the correct data", func() {
				events, err := walletWrapper.AddedToWhitelistEvents(context.Background(), nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(events).To(HaveLen(1))
				Expect(events[0].Data).To(HaveLen(5))
				Expect(hexutil.Encode(events[0].Data[0])).To(Equal(BankAccount.Address().Hash().Hex()))
				Expect(hexutil.Encode(events[0].Data[1])).To(Equal("0x0000000000000000000000000000000000000000000000000000000000000040"))
				Expect(hexutil.Encode(events[0].Data[2])).To(Equal("0x0000000000000000000000000000000000000000000000000000000000000002"))
				Expect(hexutil.Encode(events[0].Data[3])).To(Equal("0x0000000000000000000000000000000000000000000000000000000000000123"))
				Expect(hexutil.Encode(events[0].Data[4])).To(Equal("0x0000000000000000000000000000000000000000000000000000000000000124"))
			})
		})
		Context("When one address has been submitted to whitelist addition", func() {
			BeforeEach(func() {
				signedTx, err := walletWrapper.SubmitWhitelistAddition(opts, []common.Address{common.HexToAddress("0x125")})
				Expect(err).ToNot(HaveOccurred())
				Expect(Backend.SendTransaction(context.Background(), signedTx)).To(Succeed())
				Backend.Commit()
				Expect(isSuccessful(signedTx)).To(BeTrue())
			})
			It("should have one address pending whitelist addition", func() {
				pendingAddition, err := walletWrapper.PendingWhitelistAddition(context.Background(), nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(pendingAddition).To(HaveLen(1))
				Expect(pendingAddition[0]).To(Equal(common.HexToAddress("0x125")))
			})
			Context("When the SubmittedWhitelistAddition events have been filtered", func() {
				It("should return a single SubmittedWhitelistAddition event with the correct data", func() {
					events, err := walletWrapper.SubmittedWhitelistAdditionEvents(context.Background(), nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(events).To(HaveLen(1))
					Expect(events[0].Data).To(HaveLen(3))
					Expect(hexutil.Encode(events[0].Data[0])).To(Equal("0x0000000000000000000000000000000000000000000000000000000000000020"))
					Expect(hexutil.Encode(events[0].Data[1])).To(Equal("0x0000000000000000000000000000000000000000000000000000000000000001"))
					Expect(hexutil.Encode(events[0].Data[2])).To(Equal("0x0000000000000000000000000000000000000000000000000000000000000125"))
				})
			})
		})
		Context("When one address has been submitted to whitelist removal", func() {
			BeforeEach(func() {
				signedTx, err := walletWrapper.SubmitWhitelistRemoval(opts, []common.Address{common.HexToAddress("0x123")})
				Expect(err).ToNot(HaveOccurred())
				Expect(Backend.SendTransaction(context.Background(), signedTx)).To(Succeed())
				Backend.Commit()
				Expect(isSuccessful(signedTx)).To(BeTrue())
			})
			It("should have one address pending whitelist removal", func() {
				pendingRemoval, err := walletWrapper.PendingWhitelistRemoval(context.Background(), nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(pendingRemoval).To(HaveLen(1))
				Expect(pendingRemoval[0]).To(Equal(common.HexToAddress("0x123")))
			})
			Context("When the SubmittedWhitelistRemoval events have been filtered", func() {
				It("should return a single SubmittedWhitelistRemoval event with the correct data", func() {
					events, err := walletWrapper.SubmittedWhitelistRemovalEvents(context.Background(), nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(events).To(HaveLen(1))
					Expect(events[0].Data).To(HaveLen(3))
					Expect(hexutil.Encode(events[0].Data[0])).To(Equal("0x0000000000000000000000000000000000000000000000000000000000000020"))
					Expect(hexutil.Encode(events[0].Data[1])).To(Equal("0x0000000000000000000000000000000000000000000000000000000000000001"))
					Expect(hexutil.Encode(events[0].Data[2])).To(Equal("0x0000000000000000000000000000000000000000000000000000000000000123"))
				})
			})
			Context("When the address pending whitelist removal is confirmed", func() {
				BeforeEach(func() {
					opts = &contracts.ConstructOpts{
						From:     Controller.Address(),
						Nonce:    0,
						Value:    nil,
						GasPrice: nil,
						GasLimit: 0,
						Sign: func(_ context.Context, tx *types.Transaction) (*types.Transaction, error) {
							return Controller.SignTransaction(Backend, tx)
						},
						Context: context.Background(),
					}
					signedTx, err := walletWrapper.ConfirmWhitelistRemoval(opts)
					Expect(err).ToNot(HaveOccurred())
					Expect(Backend.SendTransaction(context.Background(), signedTx)).To(Succeed())
					Backend.Commit()
					Expect(isSuccessful(signedTx)).To(BeTrue())
				})
				It("should remove the address from the whitelist", func() {
					whitelisted, err := walletWrapper.IsWhitelisted(context.Background(), nil, common.HexToAddress("0x123"))
					Expect(err).ToNot(HaveOccurred())
					Expect(whitelisted).To(BeFalse())
				})
				Context("When the RemovedFromWhitelist events have been filtered", func() {
					It("should return a single RemovedFromWhitelist event with the correct data", func() {
						events, err := walletWrapper.RemovedFromWhitelistEvents(context.Background(), nil)
						Expect(err).ToNot(HaveOccurred())
						Expect(events).To(HaveLen(1))
						Expect(events[0].Data).To(HaveLen(4))
						Expect(hexutil.Encode(events[0].Data[0])).To(Equal(Controller.Address().Hash().Hex()))
						Expect(hexutil.Encode(events[0].Data[1])).To(Equal("0x0000000000000000000000000000000000000000000000000000000000000040"))
						Expect(hexutil.Encode(events[0].Data[2])).To(Equal("0x0000000000000000000000000000000000000000000000000000000000000001"))
						Expect(hexutil.Encode(events[0].Data[3])).To(Equal("0x0000000000000000000000000000000000000000000000000000000000000123"))
					})
				})
			})
		})
	})
})
