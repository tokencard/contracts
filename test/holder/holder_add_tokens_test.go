package holder_test

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tokencard/contracts/test/shared"
	"github.com/tokencard/ethertest"
)

var _ = Describe("holder addTokens", func() {

	Context("When called by the owner", func() {

		Context("When the added tokens are not supported", func() {
			var tx *types.Transaction
			BeforeEach(func() {
				var err error
				tokens := []common.Address{common.HexToAddress("0x0"), common.HexToAddress("0x1")}
				tx, err = TokenHolder.AddTokens(Owner.TransactOpts(), tokens)
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
			})

			It("Should succeed", func() {
				Expect(isSuccessful(tx)).To(BeTrue())
			})

			It("Should update the tokenExists mapping", func() {
				exists, err := TokenHolder.TokenExists(nil, common.HexToAddress("0x0"))
				Expect(err).ToNot(HaveOccurred())
				Expect(exists).To(BeTrue())

				exists, err = TokenHolder.TokenExists(nil, common.HexToAddress("0x1"))
				Expect(err).ToNot(HaveOccurred())
				Expect(exists).To(BeTrue())
			})

			It("Should emit 2 AddedToken events", func() {
				it, err := TokenHolder.FilterAddedToken(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(it.Next()).To(BeTrue())
				evt := it.Event
				Expect(it.Next()).To(BeTrue())
        Expect(evt.Sender).To(Equal(Owner.Address()))
				Expect(evt.Token).To(Equal(common.HexToAddress("0x0")))
				evt = it.Event
				Expect(it.Next()).To(BeFalse())
        Expect(evt.Sender).To(Equal(Owner.Address()))
				Expect(evt.Token).To(Equal(common.HexToAddress("0x1")))
			})

			Context("When at least one of the added tokens is already supported", func() {
				It("Should fail", func() {
					var err error
					tokens := []common.Address{common.HexToAddress("0x0"), common.HexToAddress("0x2")}
					tx, err = TokenHolder.AddTokens(Owner.TransactOpts(ethertest.WithGasLimit(200000)), tokens)
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isGasExhausted(tx, 200000)).To(BeFalse())
					Expect(isSuccessful(tx)).To(BeFalse())
				})
			})

			Context("When none of the added tokens is supported", func() {
				It("Should pass", func() {
					var err error
					tokens := []common.Address{common.HexToAddress("0x4"), common.HexToAddress("0x5")}
					tx, err = TokenHolder.AddTokens(Owner.TransactOpts(), tokens)
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isGasExhausted(tx, 300000)).To(BeFalse())
					Expect(isSuccessful(tx)).To(BeTrue())
				})
			})

		})

		Context("When trying to add #0 tokens (empty list)", func() {

			var tx *types.Transaction
			BeforeEach(func() {
				var err error
        tx, err = TokenHolder.AddTokens(Owner.TransactOpts(), nil)
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
			})

			It("Should succeed", func() {
				Expect(isSuccessful(tx)).To(BeTrue())
			})

			It("Should NOT emit an AddedToken event", func() {
				it, err := TokenHolder.FilterAddedToken(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(it.Next()).To(BeFalse())
			})
		})

	})

	Context("When called by a random address", func() {
		It("Should fail", func() {
			tokens := []common.Address{common.HexToAddress("0x0"), common.HexToAddress("0x1")}
			tx, err := Oracle.AddTokens(
				RandomAccount.TransactOpts(ethertest.WithGasLimit(300000)),
				tokens,
				StringsToByte32(
					"BNT",
					"TKN",
				),
				[]*big.Int{
					DecimalsToMagnitude(big.NewInt(18)),
					DecimalsToMagnitude(big.NewInt(8)),
				},
				big.NewInt(20180913153211),
			)
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isGasExhausted(tx, 300000)).To(BeFalse())
			Expect(isSuccessful(tx)).To(BeFalse())
		})
	})

})
