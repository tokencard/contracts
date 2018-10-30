package wallet_test

import (
	"github.com/ethereum/go-ethereum/common"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tokencard/contracts/test/shared"
	"github.com/tokencard/ethertest"
)

var _ = Describe("initializeWhitelist", func() {
	BeforeEach(func() {
		tx, err := Wallet.InitializeSpendLimit(Owner.TransactOpts(), EthToWei(1))
		Expect(err).ToNot(HaveOccurred())
		Backend.Commit()
		Expect(isSuccessful(tx)).To(BeTrue())

	})

	Context("When wallet has been deployed", func() {
		It("should have empty pending addition", func() {
			pa, err := Wallet.PendingWhitelistAddition(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(pa).To(HaveLen(0))
		})

		It("should have empty pending pendingRemoval", func() {
			pa, err := Wallet.PendingWhitelistRemoval(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(pa).To(HaveLen(0))
		})
	})

	Context("When I initialize whitelist with random person's address", func() {
		BeforeEach(func() {
			tx, err := Wallet.InitializeWhitelist(Owner.TransactOpts(), []common.Address{RandomAccount.Address()})
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeTrue())
		})

		It("should update the initializedWhitelist flag", func() {
			initialized, err := Wallet.InitializedWhitelist(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(initialized).To(BeTrue())
		})

		It("should add the random person to the whitelist", func() {
			isWhitelisted, err := Wallet.IsWhitelisted(nil, RandomAccount.Address())
			Expect(err).ToNot(HaveOccurred())
			Expect(isWhitelisted).To(BeTrue())
		})

		It("should emit WhitelistAddition event", func() {
			it, err := Wallet.FilterAddedToWhitelist(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(it.Next()).To(BeTrue())
			evt := it.Event
			Expect(it.Next()).To(BeFalse())
			Expect(evt.Addresses).To(Equal([]common.Address{RandomAccount.Address()}))
		})

		Context("When I try to initialize the whitelist again", func() {
			It("should fail", func() {
				tx, err := Wallet.InitializeWhitelist(Owner.TransactOpts(ethertest.WithGasLimit(100000)), []common.Address{RandomAccount.Address()})
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeFalse())
			})
		})

	})

	Context("When I initialize whitelist with the 0x0 address", func() {
		BeforeEach(func() {
			tx, err := Wallet.InitializeWhitelist(Owner.TransactOpts(ethertest.WithGasLimit(100000)), []common.Address{common.HexToAddress("0x0")})
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeFalse())
		})

		It("should not update the initializedWhitelist flag", func() {
			initialized, err := Wallet.InitializedWhitelist(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(initialized).To(BeFalse())
		})

		It("should not add the zero address to the whitelist", func() {
			isWhitelisted, err := Wallet.IsWhitelisted(nil, common.HexToAddress("0x0"))
			Expect(err).ToNot(HaveOccurred())
			Expect(isWhitelisted).To(BeFalse())
		})

		It("should not emit WhitelistAddition event", func() {
			it, err := Wallet.FilterAddedToWhitelist(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(it.Next()).To(BeFalse())
		})

		Context("When I try to initialize the whitelist again", func() {
			It("should pass", func() {
				tx, err := Wallet.InitializeWhitelist(Owner.TransactOpts(ethertest.WithGasLimit(100000)), []common.Address{RandomAccount.Address()})
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})
		})

	})

	Context("When controller tries to initialize the whitelist", func() {
		BeforeEach(func() {
			BankAccount.MustTransfer(Backend, Controller.Address(), EthToWei(1))
		})
		It("should fail", func() {
			tx, err := Wallet.InitializeWhitelist(Controller.TransactOpts(ethertest.WithGasLimit(65000)), []common.Address{RandomAccount.Address()})
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeFalse())
		})
	})

	Context("When random person tries to initialize the whitelist", func() {
		BeforeEach(func() {
			BankAccount.MustTransfer(Backend, RandomAccount.Address(), EthToWei(1))
		})
		It("should fail", func() {
			tx, err := Wallet.InitializeWhitelist(RandomAccount.TransactOpts(ethertest.WithGasLimit(100000)), []common.Address{RandomAccount.Address()})
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeFalse())
		})
	})

})

var _ = Describe("whitelistAddition", func() {
	BeforeEach(func() {
		BankAccount.MustTransfer(Backend, Controller.Address(), EthToWei(1))
	})
	Context("When I add a random person to the whitelist before initialization", func() {
		It("should fail", func() {
			tx, err := Wallet.SubmitWhitelistAddition(Owner.TransactOpts(ethertest.WithGasLimit(1000000)), []common.Address{RandomAccount.Address()})
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeFalse())

			submitted, err := Wallet.SubmittedWhitelistAddition(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(submitted).To(BeFalse())

			initialized, err := Wallet.InitializedWhitelist(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(initialized).To(BeFalse())
		})
	})

	Context("When I add random person to the whitelist after initialization", func() {
		BeforeEach(func() {
			tx, err := Wallet.InitializeWhitelist(Owner.TransactOpts(ethertest.WithGasLimit(500000)), []common.Address{common.HexToAddress("0x1")})
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeTrue())

			tx, err = Wallet.SubmitWhitelistAddition(Owner.TransactOpts(ethertest.WithGasLimit(500000)), []common.Address{RandomAccount.Address()})
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeTrue())
		})

		It("should update the whitelist submission flag", func() {
			submitted, err := Wallet.SubmittedWhitelistAddition(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(submitted).To(BeTrue())
		})

		It("should add the random persons's address to the pending addition list", func() {
			pending, err := Wallet.PendingWhitelistAddition(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(pending).To(Equal([]common.Address{RandomAccount.Address()}))
		})

		It("should emit a whitelist submission event", func() {
			it, err := Wallet.FilterSubmittedWhitelistAddition(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(it.Next()).To(BeTrue())
			evt := it.Event
			Expect(it.Next()).To(BeFalse())
			Expect(evt.Addresses).To(Equal([]common.Address{RandomAccount.Address()}))
		})

		Context("When I try to add another person to the whitelist", func() {
			It("should fail", func() {
				tx, err := Wallet.SubmitWhitelistAddition(Owner.TransactOpts(ethertest.WithGasLimit(500000)), []common.Address{RandomAccount.Address()})
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeFalse())
			})
		})

		Context("When I try to remove a person from the whitelist while addition is pending", func() {
			It("should fail", func() {
				tx, err := Wallet.SubmitWhitelistRemoval(Owner.TransactOpts(ethertest.WithGasLimit(500000)), []common.Address{RandomAccount.Address()})
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeFalse())
			})
		})

		Context("When the controller confirms the adding to the whitelist", func() {
			BeforeEach(func() {
				tx, err := Wallet.ConfirmWhitelistAddition(Controller.TransactOpts(ethertest.WithGasLimit(500000)))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})

			It("should add the address to the whitelist", func() {
				isWhitelisted, err := Wallet.IsWhitelisted(nil, RandomAccount.Address())
				Expect(err).ToNot(HaveOccurred())
				Expect(isWhitelisted).To(BeTrue())
			})

			It("should emit whitelist addition event", func() {
				it, err := Wallet.FilterAddedToWhitelist(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(it.Next()).To(BeTrue())
				Expect(it.Next()).To(BeTrue())
				evt := it.Event
				Expect(it.Next()).To(BeFalse())
				Expect(evt.Addresses).To(Equal([]common.Address{RandomAccount.Address()}))
			})

			Context("When I try to add another person to the whitelist", func() {
				It("should succeed", func() {
					tx, err := Wallet.SubmitWhitelistAddition(Owner.TransactOpts(), []common.Address{RandomAccount.Address()})
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeTrue())
				})
			})

			Context("When I try to initialize the whitelist", func() {
				It("should fail", func() {
					tx, err := Wallet.InitializeWhitelist(Owner.TransactOpts(ethertest.WithGasLimit(500000)), []common.Address{RandomAccount.Address()})
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeFalse())
				})
			})
		})

		Context("When the controller cancels adding to the whitelist", func() {
			BeforeEach(func() {
				tx, err := Wallet.CancelWhitelistAddition(Controller.TransactOpts())
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})

			It("should emit CancelWhitelistAddition event", func() {
				it, err := Wallet.FilterCancelledWhitelistAddition(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(it.Next()).To(BeTrue())
				evt := it.Event
				Expect(it.Next()).To(BeFalse())
				Expect(evt.Sender).To(Equal(Controller.Address()))
			})

			It("should have empty pendingAddition list", func() {
				pending, err := Wallet.PendingWhitelistAddition(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(pending).To(HaveLen(0))
			})

			Context("When I try to initialize the whitelist", func() {
				It("should fail", func() {
					tx, err := Wallet.InitializeWhitelist(Owner.TransactOpts(ethertest.WithGasLimit(100000)), []common.Address{RandomAccount.Address()})
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeFalse())
				})
			})

			Context("When I try to add another person to the whitelist", func() {
				It("should succeed", func() {
					tx, err := Wallet.SubmitWhitelistAddition(Owner.TransactOpts(), []common.Address{RandomAccount.Address()})
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeTrue())
				})
			})
		})
	})

	Context("When I submit a list of addresses to Backend added to the whitelist after initialization", func() {
		BeforeEach(func() {
			tx, err := Wallet.InitializeWhitelist(Owner.TransactOpts(ethertest.WithGasLimit(500000)), []common.Address{common.HexToAddress("0x1")})
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeTrue())
		})

		It("should succeed", func() {
			addresses := []common.Address{}
			for i := 0; i < 20; i++ {
				addresses = append(addresses, ethertest.NewAccount().Address())
			}
			tx, err := Wallet.SubmitWhitelistAddition(Owner.TransactOpts(), addresses)
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeTrue())
		})
	})

	Context("When I submit an empty list of addresses to Backend added to the whitelist after initialization", func() {
		BeforeEach(func() {
			tx, err := Wallet.InitializeWhitelist(Owner.TransactOpts(ethertest.WithGasLimit(500000)), []common.Address{common.HexToAddress("0x1")})
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeTrue())
		})

		It("should succeed", func() {
			addresses := []common.Address{}
			tx, err := Wallet.SubmitWhitelistAddition(Owner.TransactOpts(), addresses)
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeTrue())
		})
	})
})

var _ = Describe("whitelistRemoval", func() {
	BeforeEach(func() {
		BankAccount.MustTransfer(Backend, Controller.Address(), EthToWei(1))
	})
	Context("When random person is in the whitelist", func() {
		BeforeEach(func() {
			tx, err := Wallet.InitializeWhitelist(Owner.TransactOpts(), []common.Address{RandomAccount.Address()})
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeTrue())
		})

		Context("When I submit removal of random person from the whitelist", func() {
			BeforeEach(func() {
				tx, err := Wallet.SubmitWhitelistRemoval(Owner.TransactOpts(), []common.Address{RandomAccount.Address()})
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})

			It("should update the submittedWhitelistRemoval; flag", func() {
				submitted, err := Wallet.SubmittedWhitelistRemoval(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(submitted).To(BeTrue())
			})

			It("should emit SubmitWhitelistRemoval event", func() {
				it, err := Wallet.FilterSubmittedWhitelistRemoval(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(it.Next()).To(BeTrue())
				evt := it.Event
				Expect(it.Next()).To(BeFalse())
				Expect(evt.Addresses).To(Equal([]common.Address{RandomAccount.Address()}))
			})

			It("should add random person's address to pendingRemoval", func() {
				pending, err := Wallet.PendingWhitelistRemoval(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(pending).To(Equal([]common.Address{RandomAccount.Address()}))
			})

			Context("When I try to submit another list for removal", func() {
				It("should fail", func() {
					tx, err := Wallet.SubmitWhitelistRemoval(Owner.TransactOpts(ethertest.WithGasLimit(100000)), []common.Address{RandomAccount.Address()})
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeFalse())
				})
			})

			Context("When I try to submit a list for addition while removal is pending", func() {
				It("should fail", func() {
					tx, err := Wallet.SubmitWhitelistAddition(Owner.TransactOpts(ethertest.WithGasLimit(100000)), []common.Address{RandomAccount.Address()})
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeFalse())
				})
			})

			Context("When controller confirms the whitelist removal", func() {
				BeforeEach(func() {
					tx, err := Wallet.ConfirmWhitelistRemoval(Controller.TransactOpts())
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeTrue())
				})

				It("should remove random person from the whitelist", func() {
					isWhitelisted, err := Wallet.IsWhitelisted(nil, RandomAccount.Address())
					Expect(err).ToNot(HaveOccurred())
					Expect(isWhitelisted).To(BeFalse())
				})

				It("should emit WhitelistRemoval event", func() {
					it, err := Wallet.FilterRemovedFromWhitelist(nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(it.Next()).To(BeTrue())
					evt := it.Event
					Expect(it.Next()).To(BeFalse())
					Expect(evt.Addresses).To(Equal([]common.Address{RandomAccount.Address()}))
				})

				It("should clear pendingRemoval", func() {
					pending, err := Wallet.PendingWhitelistRemoval(nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(pending).To(HaveLen(0))
				})
				Context("When I try to submit addreses for removal", func() {
					It("should succeed", func() {
						tx, err := Wallet.SubmitWhitelistRemoval(Owner.TransactOpts(), []common.Address{RandomAccount.Address()})
						Expect(err).ToNot(HaveOccurred())
						Backend.Commit()
						Expect(isSuccessful(tx)).To(BeTrue())
					})
				})
			})

			Context("When controller cancels the whitelist removal", func() {
				BeforeEach(func() {
					tx, err := Wallet.CancelWhitelistRemoval(Controller.TransactOpts())
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeTrue())
				})

				It("should emit CancelWhitelistRemoval event", func() {
					it, err := Wallet.FilterCancelledWhitelistRemoval(nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(it.Next()).To(BeTrue())
					evt := it.Event
					Expect(it.Next()).To(BeFalse())
					Expect(evt.Sender).To(Equal(Controller.Address()))
				})

				It("should have empty pendingRemoval list", func() {
					pending, err := Wallet.PendingWhitelistRemoval(nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(pending).To(HaveLen(0))
				})

				It("should keep the random person's address in the whitelist", func() {
					isWhitelisted, err := Wallet.IsWhitelisted(nil, RandomAccount.Address())
					Expect(err).ToNot(HaveOccurred())
					Expect(isWhitelisted).To(BeTrue())
				})
				Context("When I try to submit another list for removal", func() {
					It("should succeed", func() {
						tx, err := Wallet.SubmitWhitelistRemoval(Owner.TransactOpts(), []common.Address{RandomAccount.Address()})
						Expect(err).ToNot(HaveOccurred())
						Backend.Commit()
						Expect(isSuccessful(tx)).To(BeTrue())
					})
				})

			})
		})

	})

	Context("When I submit 20 addresses for removal from whitelist", func() {
		It("should succeed", func() {
			addresses := []common.Address{}
			for i := 0; i < 20; i++ {
				addresses = append(addresses, ethertest.NewAccount().Address())
			}
			tx, err := Wallet.SubmitWhitelistRemoval(Owner.TransactOpts(), addresses)
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeTrue())
		})
	})

})
