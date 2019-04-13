package token_whitelist_test

import (
	"math/big"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tokencard/contracts/test/shared"
	"github.com/tokencard/ethertest"
)

var _ = Describe("updateMaxStablecoinLoadLimit", func() {


	It("Should return 10000", func() {
		mll, err := TokenWhitelist.MaxStablecoinLoadLimit(nil)
		Expect(err).ToNot(HaveOccurred())
		Expect(mll.String()).To(Equal("10000"))
	})

	It("Should return 10000", func() {
		mll, err := TokenWhitelist.MaxStablecoinLoadLimit(nil)
		Expect(err).ToNot(HaveOccurred())
		Expect(mll.String()).To(Equal("10000"))
	})

	When("called by the controller", func() {
			BeforeEach(func() {
				tx, err := TokenWhitelist.UpdateMaxStablecoinLoadLimit(Controller.TransactOpts(),big.NewInt(15000))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})

			It("Should update the max load limit", func() {
				mll, err := TokenWhitelist.MaxStablecoinLoadLimit(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(mll.String()).To(Equal("15000"))
			})

			It("Should update the max load limit", func() {
				mll, err := TokenWhitelist.MaxStablecoinLoadLimit(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(mll.String()).To(Equal("15000"))
			})

			It("Should emit a MaxStablecoinLoadLimit event", func() {
				it, err := TokenWhitelist.FilterUpdatedMaxStablecoinLoadLimit(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(it.Next()).To(BeTrue())
				evt := it.Event
				Expect(it.Next()).To(BeFalse())
				Expect(evt.Sender).To(Equal(Controller.Address()))
				Expect(evt.NewLimit.String()).To(Equal("15000"))
			})
		})
		When("called by a random account", func() {
			It("Should fail", func() {
				tx, err := TokenWhitelistableExporter.UpdateMaxStablecoinLoadLimit(RandomAccount.TransactOpts(ethertest.WithGasLimit(100000)),big.NewInt(15000))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeFalse())
			})
		})
	})
