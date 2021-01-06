package wallet_test

import (
	"github.com/ethereum/go-ethereum/common"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tokencard/contracts/v3/test/shared"

	"github.com/tokencard/ethertest"
)

var _ = Describe("2FA", func() {

	It("should be true when checking monolith 2FA variable value", func() {
		oo, err := WalletProxy.Monolith2FA(nil)
		Expect(err).ToNot(HaveOccurred())
		Expect(oo).To(BeTrue())
	})

	It("should NOT allow a non-owner to set Monolith 2FA", func() {
		tx, err := WalletProxy.SetMonolith2FA(Controller.TransactOpts(ethertest.WithGasLimit(60000)))
		Expect(err).ToNot(HaveOccurred())
		Backend.Commit()
		Expect(isSuccessful(tx)).To(BeFalse())
		returnData, _ := ethCall(tx)
		Expect(string(returnData[len(returnData)-64:])).To(ContainSubstring("sender is not owner"))
	})

	It("should fail to set Monolith 2FA if Monolith 2FA is already enabled", func() {
		tx, err := WalletProxy.SetMonolith2FA(Owner.TransactOpts(ethertest.WithGasLimit(60000)))
		Expect(err).ToNot(HaveOccurred())
		Backend.Commit()
		Expect(isSuccessful(tx)).To(BeFalse())
		returnData, _ := ethCall(tx)
		Expect(string(returnData[len(returnData)-64:])).To(ContainSubstring("monolith2FA already enabled"))
	})

	It("should NOT allow a non-owner to set personal 2FA", func() {
		tx, err := WalletProxy.SetPersonal2FA(Controller.TransactOpts(ethertest.WithGasLimit(60000)), common.HexToAddress(("0x1")))
		Expect(err).ToNot(HaveOccurred())
		Backend.Commit()
		Expect(isSuccessful(tx)).To(BeFalse())
		returnData, _ := ethCall(tx)
		Expect(string(returnData[len(returnData)-64:])).To(ContainSubstring("sender is not owner"))
	})

	It("should NOT allow a non-owner to set Monolith 2FA", func() {
		tx, err := WalletProxy.SetMonolith2FA(Controller.TransactOpts(ethertest.WithGasLimit(60000)))
		Expect(err).ToNot(HaveOccurred())
		Backend.Commit()
		Expect(isSuccessful(tx)).To(BeFalse())
		returnData, _ := ethCall(tx)
		Expect(string(returnData[len(returnData)-64:])).To(ContainSubstring("sender is not owner"))
	})

	It("should NOT allow an owner to set the personal 2FA to address 0x0", func() {
		tx, err := WalletProxy.SetPersonal2FA(Owner.TransactOpts(ethertest.WithGasLimit(60000)), common.HexToAddress(("0x0")))
		Expect(err).ToNot(HaveOccurred())
		Backend.Commit()
		Expect(isSuccessful(tx)).To(BeFalse())
		returnData, _ := ethCall(tx)
		Expect(string(returnData[len(returnData)-64:])).To(ContainSubstring("2FA cannot be set to zero"))
	})

	It("should NOT allow an owner to set the personal 2FA to the same address", func() {
		tx, err := WalletProxy.SetPersonal2FA(Owner.TransactOpts(ethertest.WithGasLimit(60000)), Owner.Address())
		Expect(err).ToNot(HaveOccurred())
		Backend.Commit()
		Expect(isSuccessful(tx)).To(BeTrue())
		returnData, _ := ethCall(tx)
		tx, err = WalletProxy.SetPersonal2FA(Owner.TransactOpts(ethertest.WithGasLimit(60000)), Owner.Address())
		Expect(err).ToNot(HaveOccurred())
		Backend.Commit()
		Expect(isSuccessful(tx)).To(BeFalse())
		returnData, _ = ethCall(tx)
		Expect(string(returnData[len(returnData)-64:])).To(ContainSubstring("address already set"))
	})

	//TODO add privileged test

	When("the owner submits an account whitelist addition", func() {
		BeforeEach(func() {
			tx, err := WalletProxy.SetWhitelist(Owner.TransactOpts(), []common.Address{common.HexToAddress("0x1")})
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeTrue())

			tx, err = WalletProxy.SubmitWhitelistAddition(Owner.TransactOpts(), []common.Address{RandomAccount.Address()})
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeTrue())
		})

		When("a random account tries to confirm the account addition to the whitelist", func() {
			It("should fail", func() {
				pwl, err := WalletProxy.PendingWhitelistAddition(nil)
				Expect(err).ToNot(HaveOccurred())
				hash, err := WalletProxy.CalculateHash(nil, pwl)
				Expect(err).ToNot(HaveOccurred())
				tx, err := WalletProxy.CancelWhitelistAddition(RandomAccount.TransactOpts(ethertest.WithGasLimit(500000)), hash)
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeFalse())
				returnData, _ := ethCall(tx)
				Expect(string(returnData[len(returnData)-64:])).To(ContainSubstring("only owner or 2FA"))
			})
		})

		When("the controller (Monolith 2FA) tries to confirm the account addition to the whitelist", func() {
			It("should succeed", func() {
				pwl, err := WalletProxy.PendingWhitelistAddition(nil)
				Expect(err).ToNot(HaveOccurred())
				hash, err := WalletProxy.CalculateHash(nil, pwl)
				Expect(err).ToNot(HaveOccurred())
				tx, err := WalletProxy.CancelWhitelistAddition(Controller.TransactOpts(), hash)
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})
		})
	})

	When("the owner submits a daily limit of 12k USD", func() {
		BeforeEach(func() {
			tx, err := WalletProxy.SubmitDailyLimitUpdate(Owner.TransactOpts(), EthToWei(12000))
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeTrue())
		})

		When("Monolith opt-out is NOT enabled and controller confirms", func() {
			It("should succeed", func() {
				tx, err := WalletProxy.ConfirmDailyLimitUpdate(Controller.TransactOpts(), EthToWei(12000))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})

			It("should fail when a random account tries to confirm", func() {
				tx, err := WalletProxy.ConfirmDailyLimitUpdate(RandomAccount.TransactOpts(ethertest.WithGasLimit(80000)), EthToWei(12000))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeFalse())
				returnData, _ := ethCall(tx)
				Expect(string(returnData[len(returnData)-64:])).To(ContainSubstring("sender is not a Monolith 2FA"))
			})
		})

		When("Monolith opt-out is enabled and 2FA is a random account", func() {
			BeforeEach(func() {
				tx, err := WalletProxy.SetPersonal2FA(Owner.TransactOpts(), RandomAccount.Address())
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})

			It("should return false when checking monolith2FA variable value", func() {
				m2fa, err := WalletProxy.Monolith2FA(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(m2fa).To(BeFalse())
			})

			It("should set the random account address as personal 2FA", func() {
				p2fa, err := WalletProxy.Personal2FA(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(p2fa).To(Equal(RandomAccount.Address()))
			})

			It("should emit a SetPersonal2Fa event", func() {
				it, err := WalletProxy.FilterSetPersonal2FA(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(it.Next()).To(BeTrue())
				evt := it.Event
				Expect(it.Next()).To(BeFalse())
				Expect(evt.Sender).To(Equal(Owner.Address()))
				Expect(evt.P2FA).To(Equal(RandomAccount.Address()))
			})

			It("should fail when controller (ex Monolith 2FA) tries to confirm the limit update", func() {
				tx, err := WalletProxy.ConfirmDailyLimitUpdate(Controller.TransactOpts(ethertest.WithGasLimit(80000)), EthToWei(12000))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeFalse())
				returnData, _ := ethCall(tx)
				Expect(string(returnData[len(returnData)-64:])).To(ContainSubstring("sender is not personal 2FA"))
			})

			When("the owner set Monlith 2FA again", func() {
				BeforeEach(func() {
					tx, err := WalletProxy.SetMonolith2FA(Owner.TransactOpts())
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeTrue())
				})

				It("should be true when checking monolith 2FA variable value", func() {
					m2fa, err := WalletProxy.Monolith2FA(nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(m2fa).To(BeTrue())
				})

				It("should set the 2FA address to 0x0 (setting to 0x0 can be done only when calling SetMonolith2FA)", func() {
					p2fa, err := WalletProxy.Personal2FA(nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(p2fa).To(Equal(common.HexToAddress("0x0")))
				})

				It("Should emit a SetMonolith2FA event", func() {
					it, err := WalletProxy.FilterSetMonolith2FA(nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(it.Next()).To(BeTrue())
					evt := it.Event
					Expect(it.Next()).To(BeFalse())
					Expect(evt.Sender).To(Equal(Owner.Address()))
				})

				It("should succeed when the controller (Monolith 2FA) confirms the limit update", func() {
					tx, err := WalletProxy.ConfirmDailyLimitUpdate(Controller.TransactOpts(), EthToWei(12000))
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeTrue())
				})

				It("should fail when a random account tries to confirm the daily limit update", func() {
					tx, err := WalletProxy.ConfirmDailyLimitUpdate(RandomAccount.TransactOpts(ethertest.WithGasLimit(80000)), EthToWei(12000))
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeFalse())
					returnData, _ := ethCall(tx)
					Expect(string(returnData[len(returnData)-64:])).To(ContainSubstring("sender is not a Monolith 2FA"))
				})
			})

			When("the random account confirms the new limit update", func() {
				BeforeEach(func() {
					tx, err := WalletProxy.ConfirmDailyLimitUpdate(RandomAccount.TransactOpts(), EthToWei(12000))
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeTrue())
				})

				It("should have 12k USD available for spending", func() {
					tl, err := WalletProxy.DailyLimitAvailable(nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(tl.String()).To(Equal(EthToWei(12000).String()))
				})

				It("should emit a SetPersonal2FA event", func() {
					it, err := WalletProxy.FilterSetPersonal2FA(nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(it.Next()).To(BeTrue())
					evt := it.Event
					Expect(it.Next()).To(BeFalse())
					Expect(evt.Sender).To(Equal(Owner.Address()))
					Expect(evt.P2FA).To(Equal(RandomAccount.Address()))
				})

				When("the owner submits a whitelist addition", func() {
					BeforeEach(func() {
						tx, err := WalletProxy.SetWhitelist(Owner.TransactOpts(ethertest.WithGasLimit(500000)), []common.Address{common.HexToAddress("0x1")})
						Expect(err).ToNot(HaveOccurred())
						Backend.Commit()
						Expect(isSuccessful(tx)).To(BeTrue())

						tx, err = WalletProxy.SubmitWhitelistAddition(Owner.TransactOpts(ethertest.WithGasLimit(500000)), []common.Address{RandomAccount.Address()})
						Expect(err).ToNot(HaveOccurred())
						Backend.Commit()
						Expect(isSuccessful(tx)).To(BeTrue())
					})

					When("the controller tries to cancel the account addition to the whitelist", func() {
						It("should fail", func() {
							pwl, err := WalletProxy.PendingWhitelistAddition(nil)
							Expect(err).ToNot(HaveOccurred())
							hash, err := WalletProxy.CalculateHash(nil, pwl)
							Expect(err).ToNot(HaveOccurred())
							tx, err := WalletProxy.CancelWhitelistAddition(Controller.TransactOpts(ethertest.WithGasLimit(500000)), hash)
							Expect(err).ToNot(HaveOccurred())
							Backend.Commit()
							Expect(isSuccessful(tx)).To(BeFalse())
							returnData, _ := ethCall(tx)
							Expect(string(returnData[len(returnData)-64:])).To(ContainSubstring("only owner or 2FA"))
						})
					})

					When("the random account (set now as 2FA) tries to cancel the account addition to the whitelist", func() {
						It("should succeed", func() {
							pwl, err := WalletProxy.PendingWhitelistAddition(nil)
							Expect(err).ToNot(HaveOccurred())
							hash, err := WalletProxy.CalculateHash(nil, pwl)
							Expect(err).ToNot(HaveOccurred())
							tx, err := WalletProxy.CancelWhitelistAddition(RandomAccount.TransactOpts(), hash)
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
