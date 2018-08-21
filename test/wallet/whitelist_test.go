package wallet_test

import (
	"github.com/ethereum/go-ethereum/common"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/tokencard/ethertest"
)

var _ = Describe("initializeWhitelist", func() {
	BeforeEach(func() {
		tx, err := w.InitializeSpendLimit(owner.TransactOpts(), ethToWei(1))
		Expect(err).ToNot(HaveOccurred())
		be.Commit()
		Expect(isSuccessful(tx)).To(BeTrue())

	})

	Context("When wallet has been deployed", func() {
		It("Shoud have empty pending addition", func() {
			pa, err := w.PendingAddition(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(pa).To(HaveLen(0))
		})

		It("Shoud have empty pending pendingRemoval", func() {
			pa, err := w.PendingRemoval(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(pa).To(HaveLen(0))
		})
	})

	Context("When I initialize whitelist with random person's address", func() {
		BeforeEach(func() {
			tx, err := w.InitializeWhitelist(owner.TransactOpts(), []common.Address{randomPerson.Address()})
			Expect(err).ToNot(HaveOccurred())
			be.Commit()
			Expect(isSuccessful(tx)).To(BeTrue())
		})

		It("should add the random person to the whitelist", func() {
			isWhitelisted, err := w.IsWhitelisted(nil, randomPerson.Address())
			Expect(err).ToNot(HaveOccurred())
			Expect(isWhitelisted).To(BeTrue())
		})

		It("Should emit WhitelistAddition event", func() {
			it, err := w.FilterWhitelistAddition(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(it.Next()).To(BeTrue())
			evt := it.Event
			Expect(it.Next()).To(BeFalse())
			Expect(evt.Addresses).To(Equal([]common.Address{randomPerson.Address()}))
		})

		Context("When I try to initialize the whitelist again", func() {
			It("Should fail", func() {
				to := owner.TransactOpts()
				to.GasLimit = 100000
				tx, err := w.InitializeWhitelist(to, []common.Address{randomPerson.Address()})
				Expect(err).ToNot(HaveOccurred())
				be.Commit()
				Expect(isSuccessful(tx)).To(BeFalse())
			})
		})

	})

	Context("When controller tries to initialize the whitelist again", func() {
		BeforeEach(func() {
			bankWallet.MustTransfer(be, controller.Address(), ethToWei(1))
		})
		It("Should fail", func() {
			to := controller.TransactOpts()
			to.GasLimit = 100000
			tx, err := w.InitializeWhitelist(to, []common.Address{randomPerson.Address()})
			Expect(err).ToNot(HaveOccurred())
			be.Commit()
			Expect(isSuccessful(tx)).To(BeFalse())
		})
	})

	Context("When random person tries to initialize the whitelist again", func() {
		BeforeEach(func() {
			bankWallet.MustTransfer(be, randomPerson.Address(), ethToWei(1))
		})
		It("Should fail", func() {
			to := randomPerson.TransactOpts()
			to.GasLimit = 100000
			tx, err := w.InitializeWhitelist(to, []common.Address{randomPerson.Address()})
			Expect(err).ToNot(HaveOccurred())
			be.Commit()
			Expect(isSuccessful(tx)).To(BeFalse())
		})
	})

})

var _ = Describe("addToWhitelist", func() {
	BeforeEach(func() {
		bankWallet.MustTransfer(be, controller.Address(), ethToWei(1))
	})
	Context("When I add random person to the white list", func() {
		BeforeEach(func() {
			tx, err := w.AddToWhitelist(owner.TransactOpts(), []common.Address{randomPerson.Address()})
			Expect(err).ToNot(HaveOccurred())
			be.Commit()
			Expect(isSuccessful(tx)).To(BeTrue())
		})

		It("Should add the random persons's address to the pending addition", func() {
			pending, err := w.PendingAddition(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(pending).To(Equal([]common.Address{randomPerson.Address()}))
		})

		Context("When I try to add another person to the whitelist", func() {
			It("Should fail", func() {
				to := owner.TransactOpts()
				to.GasLimit = 100000
				tx, err := w.AddToWhitelist(to, []common.Address{randomPerson.Address()})
				Expect(err).ToNot(HaveOccurred())
				be.Commit()
				Expect(isSuccessful(tx)).To(BeFalse())
			})
		})

		Context("When the controller confirms the adding to the whitelist", func() {
			BeforeEach(func() {
				to := controller.TransactOpts()
				tx, err := w.AddToWhitelistConfirm(to)
				Expect(err).ToNot(HaveOccurred())
				be.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})

			It("Should add the address to the whitelist", func() {
				isWhitelisted, err := w.IsWhitelisted(nil, randomPerson.Address())
				Expect(err).ToNot(HaveOccurred())
				Expect(isWhitelisted).To(BeTrue())
			})

			It("Should emit WhitelistAddition event", func() {
				it, err := w.FilterWhitelistAddition(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(it.Next()).To(BeTrue())
				evt := it.Event
				Expect(it.Next()).To(BeFalse())
				Expect(evt.Addresses).To(Equal([]common.Address{randomPerson.Address()}))
			})

			Context("When I try to add another person to the whitelist", func() {
				It("Should succeed", func() {
					tx, err := w.AddToWhitelist(owner.TransactOpts(), []common.Address{randomPerson.Address()})
					Expect(err).ToNot(HaveOccurred())
					be.Commit()
					Expect(isSuccessful(tx)).To(BeTrue())
				})
			})

			Context("When I try to initialize the whitelist", func() {
				It("Should fail", func() {
					to := owner.TransactOpts()
					to.GasLimit = 100000
					tx, err := w.InitializeWhitelist(to, []common.Address{randomPerson.Address()})
					Expect(err).ToNot(HaveOccurred())
					be.Commit()
					Expect(isSuccessful(tx)).To(BeFalse())
				})
			})

		})

		Context("When the controller cancels adding to the whitelist", func() {
			BeforeEach(func() {
				tx, err := w.AddToWhitelistCancel(controller.TransactOpts())
				Expect(err).ToNot(HaveOccurred())
				be.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})

			It("Should have empty pendingAddition list", func() {
				pending, err := w.PendingAddition(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(pending).To(HaveLen(0))
			})

			Context("When I try to initialize the whitelist", func() {
				It("Should fail", func() {
					to := owner.TransactOpts()
					to.GasLimit = 100000
					tx, err := w.InitializeWhitelist(to, []common.Address{randomPerson.Address()})
					Expect(err).ToNot(HaveOccurred())
					be.Commit()
					Expect(isSuccessful(tx)).To(BeFalse())
				})
			})

			Context("When I try to add another person to the whitelist", func() {
				It("Should succeed", func() {
					tx, err := w.AddToWhitelist(owner.TransactOpts(), []common.Address{randomPerson.Address()})
					Expect(err).ToNot(HaveOccurred())
					be.Commit()
					Expect(isSuccessful(tx)).To(BeTrue())
				})
			})
		})
	})

	Context("When I submit 20 addresses for adding to whitelist", func() {
		It("Should succeed", func() {
			addresses := []common.Address{}
			for i := 0; i < 20; i++ {
				addresses = append(addresses, ethertest.NewWallet().Address())
			}
			tx, err := w.AddToWhitelist(owner.TransactOpts(), addresses)
			Expect(err).ToNot(HaveOccurred())
			be.Commit()
			Expect(isSuccessful(tx)).To(BeTrue())
		})
	})

	Context("When I submit 21 addresses for adding to whitelist", func() {
		It("Should fail", func() {
			addresses := []common.Address{}
			for i := 0; i < 21; i++ {
				addresses = append(addresses, ethertest.NewWallet().Address())
			}
			to := owner.TransactOpts()
			to.GasLimit = 100000
			tx, err := w.AddToWhitelist(to, addresses)
			Expect(err).ToNot(HaveOccurred())
			be.Commit()
			Expect(isSuccessful(tx)).To(BeFalse())
		})
	})
})

var _ = Describe("removeFromWhitelist", func() {
	BeforeEach(func() {
		bankWallet.MustTransfer(be, controller.Address(), ethToWei(1))
	})
	Context("When random person is in the whitelist", func() {
		BeforeEach(func() {
			tx, err := w.InitializeWhitelist(owner.TransactOpts(), []common.Address{randomPerson.Address()})
			Expect(err).ToNot(HaveOccurred())
			be.Commit()
			Expect(isSuccessful(tx)).To(BeTrue())
		})

		Context("When I submit removal of random person from the whitelist", func() {
			BeforeEach(func() {
				tx, err := w.RemoveFromWhitelist(owner.TransactOpts(), []common.Address{randomPerson.Address()})
				Expect(err).ToNot(HaveOccurred())
				be.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})
			It("should add random person's address to pendingRemoval", func() {
				pending, err := w.PendingRemoval(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(pending).To(Equal([]common.Address{randomPerson.Address()}))
			})

			Context("When I try to submit another list for removal", func() {
				It("should fail", func() {
					to := owner.TransactOpts()
					to.GasLimit = 100000
					tx, err := w.RemoveFromWhitelist(to, []common.Address{randomPerson.Address()})
					Expect(err).ToNot(HaveOccurred())
					be.Commit()
					Expect(isSuccessful(tx)).To(BeFalse())
				})
			})

			Context("When controller confirms the whitelist removal", func() {
				BeforeEach(func() {
					tx, err := w.RemoveFromWhitelistConfirm(controller.TransactOpts())
					Expect(err).ToNot(HaveOccurred())
					be.Commit()
					Expect(isSuccessful(tx)).To(BeTrue())
				})

				It("Should remove random person from the whitelist", func() {
					isWhitelisted, err := w.IsWhitelisted(nil, randomPerson.Address())
					Expect(err).ToNot(HaveOccurred())
					Expect(isWhitelisted).To(BeFalse())
				})

				It("Should emit WhitelistRemoval event", func() {
					it, err := w.FilterWhitelistRemoval(nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(it.Next()).To(BeTrue())
					evt := it.Event
					Expect(it.Next()).To(BeFalse())
					Expect(evt.Addresses).To(Equal([]common.Address{randomPerson.Address()}))
				})

				It("should clear pendingRemoval", func() {
					pending, err := w.PendingRemoval(nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(pending).To(HaveLen(0))
				})
				Context("When I try to submit addreses for removal", func() {
					It("Should succeed", func() {
						tx, err := w.RemoveFromWhitelist(owner.TransactOpts(), []common.Address{randomPerson.Address()})
						Expect(err).ToNot(HaveOccurred())
						be.Commit()
						Expect(isSuccessful(tx)).To(BeTrue())
					})
				})
			})

			Context("When controller cancels the whitelist removal", func() {
				BeforeEach(func() {
					tx, err := w.RemoveFromWhitelistCancel(controller.TransactOpts())
					Expect(err).ToNot(HaveOccurred())
					be.Commit()
					Expect(isSuccessful(tx)).To(BeTrue())
				})

				It("Should have empty pendingRemoval list", func() {
					pending, err := w.PendingRemoval(nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(pending).To(HaveLen(0))
				})

				It("should keep the random person's address in the whitelist", func() {
					isWhitelisted, err := w.IsWhitelisted(nil, randomPerson.Address())
					Expect(err).ToNot(HaveOccurred())
					Expect(isWhitelisted).To(BeTrue())
				})
				Context("When I try to submit another list for removal", func() {
					It("should succeed", func() {
						tx, err := w.RemoveFromWhitelist(owner.TransactOpts(), []common.Address{randomPerson.Address()})
						Expect(err).ToNot(HaveOccurred())
						be.Commit()
						Expect(isSuccessful(tx)).To(BeTrue())
					})
				})

			})
		})

	})

	Context("When I submit 20 addresses for removal from whitelist", func() {
		It("Should succeed", func() {
			addresses := []common.Address{}
			for i := 0; i < 20; i++ {
				addresses = append(addresses, ethertest.NewWallet().Address())
			}
			tx, err := w.RemoveFromWhitelist(owner.TransactOpts(), addresses)
			Expect(err).ToNot(HaveOccurred())
			be.Commit()
			Expect(isSuccessful(tx)).To(BeTrue())
		})
	})

	Context("When I submit 21 addresses for removal from whitelist", func() {
		It("Should fail", func() {
			addresses := []common.Address{}
			for i := 0; i < 21; i++ {
				addresses = append(addresses, ethertest.NewWallet().Address())
			}
			to := owner.TransactOpts()
			to.GasLimit = 100000
			tx, err := w.RemoveFromWhitelist(to, addresses)
			Expect(err).ToNot(HaveOccurred())
			be.Commit()
			Expect(isSuccessful(tx)).To(BeFalse())
		})
	})

})
