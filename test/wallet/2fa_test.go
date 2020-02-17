package wallet_test

import (
	"github.com/ethereum/go-ethereum/common"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tokencard/contracts/v2/test/shared"
	"github.com/tokencard/ethertest"
)

var _ = FDescribe("2FA", func() {

	It("should be true", func() {
		oo, err := Wallet.Monolith2FA(nil)
		Expect(err).ToNot(HaveOccurred())
		Expect(oo).To(BeTrue())
	})

	It("should NOT allow a non-owner to set Personal 2FA", func() {
		tx, err := Wallet.SetPersonal2FA(Controller.TransactOpts(ethertest.WithGasLimit(60000)), common.HexToAddress(("0x1")))
		Expect(err).ToNot(HaveOccurred())
		Backend.Commit()
		Expect(isSuccessful(tx)).To(BeFalse())
		returnData, _ := ethCall(tx)
		Expect(string(returnData[len(returnData)-64:])).To(ContainSubstring("sender is not owner"))
	})

    It("should NOT allow a non-owner to set Monolith 2FA", func() {
		tx, err := Wallet.SetMonolith2FA(Controller.TransactOpts(ethertest.WithGasLimit(60000)))
		Expect(err).ToNot(HaveOccurred())
		Backend.Commit()
		Expect(isSuccessful(tx)).To(BeFalse())
		returnData, _ := ethCall(tx)
		Expect(string(returnData[len(returnData)-64:])).To(ContainSubstring("sender is not owner"))
	})

    It("should NOT allow to set the personal 2FA to address 0x0", func() {
		tx, err := Wallet.SetPersonal2FA(Owner.TransactOpts(ethertest.WithGasLimit(60000)), common.HexToAddress(("0x0")))
		Expect(err).ToNot(HaveOccurred())
		Backend.Commit()
		Expect(isSuccessful(tx)).To(BeFalse())
		returnData, _ := ethCall(tx)
		Expect(string(returnData[len(returnData)-64:])).To(ContainSubstring("2FA cannot be set to zero"))
	})

    It("should fail if monolith 2FA is already enabled", func() {
		tx, err := Wallet.SetMonolith2FA(Owner.TransactOpts(ethertest.WithGasLimit(60000)))
		Expect(err).ToNot(HaveOccurred())
		Backend.Commit()
		Expect(isSuccessful(tx)).To(BeFalse())
		returnData, _ := ethCall(tx)
		Expect(string(returnData[len(returnData)-64:])).To(ContainSubstring("monolith2FA already enabled"))
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

		When("Monolith opt-out is NOT enabled and controller confirms", func() {
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

		When("Monolith opt-out is enabled and 2FA is a random account", func() {
			BeforeEach(func() {
				tx, err := Wallet.SetPersonal2FA(Owner.TransactOpts(), RandomAccount.Address())
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})

			It("should be false", func() {
				m2fa, err := Wallet.Monolith2FA(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(m2fa).To(BeFalse())
			})

            It("should set the 2FA address", func() {
				p2fa, err := Wallet.Personal2FA(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(p2fa).To(Equal(RandomAccount.Address()))
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

			It("should fail when controller tries to confirm", func() {
				tx, err := Wallet.ConfirmDailyLimitUpdate(Controller.TransactOpts(ethertest.WithGasLimit(80000)), EthToWei(12000))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeFalse())
				returnData, _ := ethCall(tx)
				Expect(string(returnData[len(returnData)-64:])).To(ContainSubstring("sender is not personal 2FA"))
			})

            When("the owner set Monlith 2FA again", func() {
				BeforeEach(func() {
                    tx, err := Wallet.SetMonolith2FA(Owner.TransactOpts())
            		Expect(err).ToNot(HaveOccurred())
            		Backend.Commit()
            		Expect(isSuccessful(tx)).To(BeTrue())
                })

                It("should be true", func() {
    				m2fa, err := Wallet.Monolith2FA(nil)
    				Expect(err).ToNot(HaveOccurred())
    				Expect(m2fa).To(BeTrue())
    			})

                It("should set the 2FA address to 0", func() {
    				p2fa, err := Wallet.Personal2FA(nil)
    				Expect(err).ToNot(HaveOccurred())
    				Expect(p2fa).To(Equal(common.HexToAddress("0x0")))
    			})

                It("Should emit a SetMonolith2FA event", func() {
    				it, err := Wallet.FilterSetMonolith2FA(nil)
    				Expect(err).ToNot(HaveOccurred())
    				Expect(it.Next()).To(BeTrue())
    				evt := it.Event
    				Expect(it.Next()).To(BeFalse())
    				Expect(evt.Sender).To(Equal(Owner.Address()))
    			})

                It("should succeed when the controller confirms the limit update", func() {
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

			When("the random account confirms the new limit", func() {
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

					When("the random account (set now as 2FA) tries to confirm the addition to the whitelist", func() {
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
