package wallet_test

import (
	"github.com/ethereum/go-ethereum/common"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tokencard/contracts/v2/test/shared"
	"github.com/tokencard/ethertest"
)

var _ = Describe("2FA", func() {

	It("should be true", func() {
		oo, err := Wallet.Monolith2FA(nil)
		Expect(err).ToNot(HaveOccurred())
		Expect(oo).To(BeTrue())
	})

	It("should NOT allow a non-owner to opt-out", func() {
		tx, err := Wallet.SetPersonal2FA(Controller.TransactOpts(ethertest.WithGasLimit(60000)), common.HexToAddress(("0x1")))
		Expect(err).ToNot(HaveOccurred())
		Backend.Commit()
		Expect(isSuccessful(tx)).To(BeFalse())
		returnData, _ := ethCall(tx)
		Expect(string(returnData[len(returnData)-64:])).To(ContainSubstring("sender is not owner"))
	})

	It("should revert if there is no change", func() {
		tx, err := Wallet.SetOptOption(Owner.TransactOpts(ethertest.WithGasLimit(60000)), false)
		Expect(err).ToNot(HaveOccurred())
		Backend.Commit()
		Expect(isSuccessful(tx)).To(BeFalse())
		returnData, _ := ethCall(tx)
		Expect(string(returnData[len(returnData)-64:])).To(ContainSubstring("no change in opt options"))
	})

	When("the owner submits a whitelist addition", func() {
		BeforeEach(func() {
			tx, err := Wallet.SetWhitelist(Owner.TransactOpts(), []common.Address{common.HexToAddress("0x1")})
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeTrue())

			tx, err = Wallet.SubmitWhitelistAddition(Owner.TransactOpts(), []common.Address{RandomAccount.Address()})
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeTrue())
		})

		When("the a random account tries to confirm the addition to the whitelist", func() {
			It("should fail", func() {
				pwl, err := Wallet.PendingWhitelistAddition(nil)
				Expect(err).ToNot(HaveOccurred())
				hash, err := Wallet.CalculateHash(nil, pwl)
				Expect(err).ToNot(HaveOccurred())
				tx, err := Wallet.CancelWhitelistAddition(RandomAccount.TransactOpts(ethertest.WithGasLimit(500000)), hash)
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeFalse())
				returnData, _ := ethCall(tx)
				Expect(string(returnData[len(returnData)-64:])).To(ContainSubstring("only owner or 2FA"))
			})
		})

		When("the controller tries to confirm the addition to the whitelist", func() {
			It("should succeed", func() {
				pwl, err := Wallet.PendingWhitelistAddition(nil)
				Expect(err).ToNot(HaveOccurred())
				hash, err := Wallet.CalculateHash(nil, pwl)
				Expect(err).ToNot(HaveOccurred())
				tx, err := Wallet.CancelWhitelistAddition(Controller.TransactOpts(), hash)
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})
		})

	})

	When("the owner submits a daily limit of 12K $USD", func() {
		BeforeEach(func() {
			tx, err := Wallet.SubmitDailyLimitUpdate(Owner.TransactOpts(), EthToWei(12000))
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeTrue())
		})

		When("opt-out is NOT set and controller confirms", func() {
			It("should succeed", func() {
				tx, err := Wallet.ConfirmDailyLimitUpdate(Controller.TransactOpts(), EthToWei(12000))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})

			It("should fail when a random account tries to confirm", func() {
				tx, err := Wallet.ConfirmDailyLimitUpdate(RandomAccount.TransactOpts(ethertest.WithGasLimit(80000)), EthToWei(12000))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeFalse())
				returnData, _ := ethCall(tx)
				Expect(string(returnData[len(returnData)-64:])).To(ContainSubstring("sender is not a controller"))
			})
		})

		When("opt-out is set", func() {
			BeforeEach(func() {
				tx, err := Wallet.SetOptOption(Owner.TransactOpts(), true)
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})

			It("should be true", func() {
				oo, err := Wallet.OptOut(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(oo).To(BeTrue())
			})

			It("Should emit a SetOptOption event", func() {
				it, err := Wallet.FilterSetOptOption(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(it.Next()).To(BeTrue())
				evt := it.Event
				Expect(it.Next()).To(BeFalse())
				Expect(evt.Sender).To(Equal(Owner.Address()))
				Expect(evt.OptOption).To(BeTrue())
			})

			It("should revert if there is no change", func() {
				tx, err := Wallet.SetOptOption(Owner.TransactOpts(ethertest.WithGasLimit(60000)), true)
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeFalse())
				returnData, _ := ethCall(tx)
				Expect(string(returnData[len(returnData)-64:])).To(ContainSubstring("no change in opt options"))
			})

			It("should fail when controller tries to confirm", func() {
				tx, err := Wallet.ConfirmDailyLimitUpdate(Controller.TransactOpts(ethertest.WithGasLimit(80000)), EthToWei(12000))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeFalse())
				returnData, _ := ethCall(tx)
				Expect(string(returnData[len(returnData)-64:])).To(ContainSubstring("sender is not personal 2FA"))
			})

			It("should fail when a random account tries to confirm", func() {
				tx, err := Wallet.ConfirmDailyLimitUpdate(RandomAccount.TransactOpts(ethertest.WithGasLimit(80000)), EthToWei(12000))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeFalse())
				returnData, _ := ethCall(tx)
				Expect(string(returnData[len(returnData)-64:])).To(ContainSubstring("sender is not personal 2FA"))
			})

			It("should fail when owner tries to set 0x0 as 2FA", func() {
				tx, err := Wallet.SetPersonal2FA(Owner.TransactOpts(ethertest.WithGasLimit(80000)), common.HexToAddress("0x0"))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeFalse())
				returnData, _ := ethCall(tx)
				Expect(string(returnData[len(returnData)-64:])).To(ContainSubstring("2FA cannot be set to zero"))
			})

			When("owner sets a random account as 2FA and it confirms the new limit", func() {
				BeforeEach(func() {
					tx, err := Wallet.SetPersonal2FA(Owner.TransactOpts(), RandomAccount.Address())
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeTrue())

					tx, err = Wallet.ConfirmDailyLimitUpdate(RandomAccount.TransactOpts(), EthToWei(12000))
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeTrue())
				})

				It("should have 12K $USD available for spending", func() {
					tl, err := Wallet.DailyLimitAvailable(nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(tl.String()).To(Equal(EthToWei(12000).String()))
				})

				It("Should emit a SetPersonal2FA event", func() {
					it, err := Wallet.FilterSetPersonal2FA(nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(it.Next()).To(BeTrue())
					evt := it.Event
					Expect(it.Next()).To(BeFalse())
					Expect(evt.Sender).To(Equal(Owner.Address()))
					Expect(evt.P2FA).To(Equal(RandomAccount.Address()))
				})

				When("the owner submits a whitelist addition", func() {
					BeforeEach(func() {
						tx, err := Wallet.SetWhitelist(Owner.TransactOpts(ethertest.WithGasLimit(500000)), []common.Address{common.HexToAddress("0x1")})
						Expect(err).ToNot(HaveOccurred())
						Backend.Commit()
						Expect(isSuccessful(tx)).To(BeTrue())

						tx, err = Wallet.SubmitWhitelistAddition(Owner.TransactOpts(ethertest.WithGasLimit(500000)), []common.Address{RandomAccount.Address()})
						Expect(err).ToNot(HaveOccurred())
						Backend.Commit()
						Expect(isSuccessful(tx)).To(BeTrue())
					})

					When("the controller tries to confirm the addition to the whitelist", func() {
						It("should fail", func() {
							pwl, err := Wallet.PendingWhitelistAddition(nil)
							Expect(err).ToNot(HaveOccurred())
							hash, err := Wallet.CalculateHash(nil, pwl)
							Expect(err).ToNot(HaveOccurred())
							tx, err := Wallet.CancelWhitelistAddition(Controller.TransactOpts(ethertest.WithGasLimit(500000)), hash)
							Expect(err).ToNot(HaveOccurred())
							Backend.Commit()
							Expect(isSuccessful(tx)).To(BeFalse())
							returnData, _ := ethCall(tx)
							Expect(string(returnData[len(returnData)-64:])).To(ContainSubstring("only owner or 2FA"))
						})
					})

					When("the random account (set as 2FA) tries to confirm the addition to the whitelist", func() {
						It("should succeed", func() {
							pwl, err := Wallet.PendingWhitelistAddition(nil)
							Expect(err).ToNot(HaveOccurred())
							hash, err := Wallet.CalculateHash(nil, pwl)
							Expect(err).ToNot(HaveOccurred())
							tx, err := Wallet.CancelWhitelistAddition(RandomAccount.TransactOpts(), hash)
							Expect(err).ToNot(HaveOccurred())
							Backend.Commit()
							Expect(isSuccessful(tx)).To(BeTrue())
						})
					})

				})

			})

		})
	})

})
