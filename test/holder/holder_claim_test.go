package holder_test

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"

	"github.com/ethereum/go-ethereum/common"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tokencard/contracts/test/shared"
	"github.com/tokencard/ethertest"
)

var _ = Describe("TokenHolder", func() {

	When("The holder contract owns 3 types of ERC20 tokens (1/3 are redeemable) and 1 ETH", func() {

		var tx *types.Transaction
		//add the tokens to the list
		BeforeEach(func() {
			tokens := []common.Address{common.HexToAddress("0x0"), ERC20Contract1Address, ERC20Contract2Address, ERC20Contract3Address}
			tx, err := TokenWhitelist.AddTokens(
				ControllerAdmin.TransactOpts(),
				tokens,
				StringsToByte32(
					"ETH",
					"ERC1",
					"ERC2",
					"ERC3",
				),
				[]*big.Int{
					DecimalsToMagnitude(big.NewInt(18)),
					DecimalsToMagnitude(big.NewInt(18)),
					DecimalsToMagnitude(big.NewInt(18)),
					DecimalsToMagnitude(big.NewInt(18)),
				},
				[]bool{true, true, true, true},
				[]bool{true, true, false, false},
				big.NewInt(20180913153211),
			)
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeTrue())
		})

		BeforeEach(func() {
			var err error
			tx, err = ERC20Contract1.Credit(BankAccount.TransactOpts(), TokenHolderAddress, big.NewInt(1000))
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeTrue())
		})

		BeforeEach(func() {
			var err error
			tx, err = ERC20Contract2.Credit(BankAccount.TransactOpts(), TokenHolderAddress, big.NewInt(500))
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeTrue())

		})

		BeforeEach(func() {
			var err error
			tx, err = ERC20Contract3.Credit(BankAccount.TransactOpts(), TokenHolderAddress, big.NewInt(1000))
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeTrue())

		})

		BeforeEach(func() {
			BankAccount.MustTransfer(Backend, TokenHolderAddress, EthToWei(1))
		})

		When("a random address tries to claim", func() {
			It("should fail", func() {
				tx, err := TokenHolder.NonRedeemableTokenClaim(RandomAccount.TransactOpts(ethertest.WithGasLimit(80000)), Owner.Address(), []common.Address{ERC20Contract3Address})
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeFalse())
			})

		})

		When("The owner tries to claim a redeemable token or ETH", func() {
			It("should fail", func() {
				tx, err := TokenHolder.NonRedeemableTokenClaim(ControllerAdmin.TransactOpts(ethertest.WithGasLimit(80000)), Owner.Address(), []common.Address{common.HexToAddress("0x0")})
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeFalse())
			})

			It("should fail", func() {
				tx, err := TokenHolder.NonRedeemableTokenClaim(ControllerAdmin.TransactOpts(ethertest.WithGasLimit(80000)), Owner.Address(), []common.Address{ERC20Contract1Address, ERC20Contract3Address})
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeFalse())
			})
		})

		When("The owner tries to claim all non-redeemable tokens", func() {
			BeforeEach(func() {
				tx, err := TokenHolder.NonRedeemableTokenClaim(ControllerAdmin.TransactOpts(), Owner.Address(), []common.Address{ERC20Contract2Address, ERC20Contract3Address})
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})

			It("Should emit 2 Claimed events", func() {
				it, err := TokenHolder.FilterClaimed(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(it.Next()).To(BeTrue())
				evt := it.Event
				Expect(it.Next()).To(BeTrue())
				Expect(evt.To).To(Equal(Owner.Address()))
				Expect(evt.Asset).To(Equal(ERC20Contract2Address))
				Expect(evt.Amount.String()).To(Equal("500"))
				evt = it.Event
				Expect(it.Next()).To(BeFalse())
				Expect(evt.To).To(Equal(Owner.Address()))
				Expect(evt.Asset).To(Equal(ERC20Contract3Address))
				Expect(evt.Amount.String()).To(Equal("1000"))
			})

			It("Should emit 2 Transfer events", func() {
				from := []common.Address{TokenHolderAddress}
				to := []common.Address{Owner.Address()}
				it, err := ERC20Contract2.FilterTransfer(nil, from, to)
				Expect(err).ToNot(HaveOccurred())
				Expect(it.Next()).To(BeTrue())
				evt := it.Event
				Expect(it.Next()).To(BeFalse())
				Expect(evt.From).To(Equal(TokenHolderAddress))
				Expect(evt.To).To(Equal(Owner.Address()))
				Expect(evt.Amount.String()).To(Equal("500"))
				it, err = ERC20Contract3.FilterTransfer(nil, from, to)
				Expect(err).ToNot(HaveOccurred())
				Expect(it.Next()).To(BeTrue())
				evt = it.Event
				Expect(evt.From).To(Equal(TokenHolderAddress))
				Expect(evt.To).To(Equal(Owner.Address()))
				Expect(evt.Amount.String()).To(Equal("1000"))
			})

			It("should increase the ERC20 type-2 balance of the owner by 500", func() {
				b, err := ERC20Contract2.BalanceOf(nil, Owner.Address())
				Expect(err).ToNot(HaveOccurred())
				Expect(b.String()).To(Equal("500"))
			})

			It("should increase the ERC20 type-3 balance of the owner by 1000", func() {
				b, err := ERC20Contract3.BalanceOf(nil, Owner.Address())
				Expect(err).ToNot(HaveOccurred())
				Expect(b.String()).To(Equal("1000"))
			})

			It("should decrease the ERC20 type-2 balance of the holder contract by 500", func() {
				b, err := ERC20Contract2.BalanceOf(nil, TokenHolderAddress)
				Expect(err).ToNot(HaveOccurred())
				Expect(b.String()).To(Equal("0"))
			})

			It("should decrease the ERC20 type-3 balance of the holder contract by 1000", func() {
				b, err := ERC20Contract3.BalanceOf(nil, TokenHolderAddress)
				Expect(err).ToNot(HaveOccurred())
				Expect(b.String()).To(Equal("0"))
			})

			It("should NOT have an impact on the (ETH) balance of the holder contract", func() {
				b, e := Backend.BalanceAt(context.Background(), TokenHolderAddress, nil)
				Expect(e).ToNot(HaveOccurred())
				Expect(b.String()).To(Equal(EthToWei(1).String()))
			})
		})

	})

})
