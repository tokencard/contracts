package wallet_test

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tokencard/contracts/v3/test/shared"
	"github.com/tokencard/ethertest"
)

var _ = Describe("initializeWhitelist", func() {

	BeforeEach(func() {
		tx, err := Wallet.SubmitDailyLimitUpdate(Owner.TransactOpts(), EthToWei(1))
		Expect(err).ToNot(HaveOccurred())
		Backend.Commit()
		Expect(isSuccessful(tx)).To(BeTrue())

	})

	When("wallet has been deployed", func() {
		It("should have empty pending addition", func() {
			pa, err := WalletProxy.PendingWhitelistAddition(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(pa).To(HaveLen(0))
		})

		It("should have empty pending pendingRemoval", func() {
			pa, err := WalletProxy.PendingWhitelistRemoval(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(pa).To(HaveLen(0))
		})
	})

	When("I initialize whitelist with random account's address", func() {
		BeforeEach(func() {
			tx, err := WalletProxy.SetWhitelist(Owner.TransactOpts(), []common.Address{RandomAccount.Address()})
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeTrue())
		})

		It("should update the initializedWhitelist flag", func() {
			initialized, err := WalletProxy.IsSetWhitelist(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(initialized).To(BeTrue())
		})

		It("should add the random account to the whitelist", func() {
			isWhitelisted, err := WalletProxy.WhitelistMap(nil, RandomAccount.Address())
			Expect(err).ToNot(HaveOccurred())
			Expect(isWhitelisted).To(BeTrue())
		})

		It("should emit WhitelistAddition event", func() {
			it, err := WalletProxy.FilterAddedToWhitelist(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(it.Next()).To(BeTrue())
			evt := it.Event
			Expect(it.Next()).To(BeFalse())
			Expect(evt.Sender).To(Equal(Owner.Address()))
			Expect(evt.Addresses).To(Equal([]common.Address{RandomAccount.Address()}))
		})

		It("should add an entry to the Whitelist Array", func() {
			a, err := WalletProxy.WhitelistArray(nil, big.NewInt(0))
			Expect(err).ToNot(HaveOccurred())
			Expect(a).To(Equal(RandomAccount.Address()))
		})

		When("I try to initialize the whitelist again", func() {
			It("should fail", func() {
				tx, err := WalletProxy.SetWhitelist(Owner.TransactOpts(ethertest.WithGasLimit(100000)), []common.Address{RandomAccount.Address()})
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeFalse())
			})
		})

	})

	When("I initialize whitelist with duplicate entries", func() {
		BeforeEach(func() {
			tx, err := WalletProxy.SetWhitelist(Owner.TransactOpts(), []common.Address{common.HexToAddress("0x1"), common.HexToAddress("0x1")})
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeTrue())
		})

		It("should only persist 1 item in the array", func() {
			_, err := WalletProxy.WhitelistArray(nil, big.NewInt(1))
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("abi: attempting to unmarshall"))
		})

		It("should only persist 1 item in the array and it should be 0x1", func() {
			a, err := WalletProxy.WhitelistArray(nil, big.NewInt(0))
			Expect(err).ToNot(HaveOccurred())

			Expect(a).To(Equal(common.HexToAddress("0x1")))
		})
	})

	When("I initialize whitelist with the 0x0 address", func() {
		BeforeEach(func() {
			tx, err := WalletProxy.SetWhitelist(Owner.TransactOpts(ethertest.WithGasLimit(100000)), []common.Address{common.HexToAddress("0x0")})
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeFalse())
		})

		It("should NOT update the initializedWhitelist flag", func() {
			initialized, err := WalletProxy.IsSetWhitelist(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(initialized).To(BeFalse())
		})

		It("should NOT add the zero address to the whitelist", func() {
			isWhitelisted, err := WalletProxy.WhitelistMap(nil, common.HexToAddress("0x0"))
			Expect(err).ToNot(HaveOccurred())
			Expect(isWhitelisted).To(BeFalse())
		})

		It("should NOT emit WhitelistAddition event", func() {
			it, err := WalletProxy.FilterAddedToWhitelist(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(it.Next()).To(BeFalse())
		})

		When("I try to initialize the whitelist again", func() {
			It("should pass", func() {
				tx, err := WalletProxy.SetWhitelist(Owner.TransactOpts(ethertest.WithGasLimit(200000)), []common.Address{RandomAccount.Address()})
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})
		})

	})

	When("the controller tries to initialize the whitelist", func() {
		BeforeEach(func() {
			BankAccount.MustTransfer(Backend, Controller.Address(), EthToWei(1))
		})
		It("should fail", func() {
			tx, err := WalletProxy.SetWhitelist(Controller.TransactOpts(ethertest.WithGasLimit(65000)), []common.Address{RandomAccount.Address()})
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeFalse())
		})
	})

	When("a random account tries to initialize the whitelist", func() {
		It("should fail", func() {
			BankAccount.MustTransfer(Backend, RandomAccount.Address(), EthToWei(1))
			tx, err := WalletProxy.SetWhitelist(RandomAccount.TransactOpts(ethertest.WithGasLimit(100000)), []common.Address{RandomAccount.Address()})
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
	When("I add a random account to the whitelist before initialization", func() {
		It("should fail", func() {
			tx, err := WalletProxy.SubmitWhitelistAddition(Owner.TransactOpts(ethertest.WithGasLimit(1000000)), []common.Address{RandomAccount.Address()})
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeFalse())

			submitted, err := WalletProxy.SubmittedWhitelistAddition(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(submitted).To(BeFalse())

			initialized, err := WalletProxy.IsSetWhitelist(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(initialized).To(BeFalse())
		})
	})

	When("I add random account to the whitelist after initialization", func() {
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

		It("should update the whitelist submission flag", func() {
			submitted, err := WalletProxy.SubmittedWhitelistAddition(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(submitted).To(BeTrue())
		})

		It("should add the random account's address to the pending addition list", func() {
			pending, err := WalletProxy.PendingWhitelistAddition(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(pending).To(Equal([]common.Address{RandomAccount.Address()}))
		})

		It("should emit a whitelist submission event", func() {
			it, err := WalletProxy.FilterSubmittedWhitelistAddition(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(it.Next()).To(BeTrue())
			evt := it.Event
			Expect(it.Next()).To(BeFalse())
			Expect(evt.Addresses).To(Equal([]common.Address{RandomAccount.Address()}))
			hash, err := WalletProxy.CalculateHash(nil, []common.Address{RandomAccount.Address()})
			Expect(it.Next()).To(BeFalse())
			Expect(evt.Hash).To(Equal(hash))
		})

		When("I try to add another account to the whitelist", func() {
			It("should fail", func() {
				tx, err := WalletProxy.SubmitWhitelistAddition(Owner.TransactOpts(ethertest.WithGasLimit(500000)), []common.Address{RandomAccount.Address()})
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeFalse())
			})
		})

		When("I try to remove a account from the whitelist while addition is pending", func() {
			It("should fail", func() {
				tx, err := WalletProxy.SubmitWhitelistRemoval(Owner.TransactOpts(ethertest.WithGasLimit(500000)), []common.Address{RandomAccount.Address()})
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeFalse())
			})
		})

		When("controller tries to confirm whitelist addition with a wrong hash", func() {
			It("should fail", func() {
				tx, err := WalletProxy.ConfirmWhitelistAddition(Controller.TransactOpts(ethertest.WithGasLimit(500000)), common.BytesToHash(nil))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeFalse())
			})
		})

		When("the controller confirms the adding to the whitelist", func() {
			BeforeEach(func() {
				pwl, err := WalletProxy.PendingWhitelistAddition(nil)
				Expect(err).ToNot(HaveOccurred())
				hash, err := WalletProxy.CalculateHash(nil, pwl)
				Expect(err).ToNot(HaveOccurred())
				tx, err := WalletProxy.ConfirmWhitelistAddition(Controller.TransactOpts(ethertest.WithGasLimit(500000)), hash)
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})

			It("should add the random account's address to the whitelist", func() {
				isWhitelisted, err := WalletProxy.WhitelistMap(nil, RandomAccount.Address())
				Expect(err).ToNot(HaveOccurred())
				Expect(isWhitelisted).To(BeTrue())
			})

			It("should add the random account's address to the Whitelist Array", func() {
				ra, err := WalletProxy.WhitelistArray(nil, big.NewInt(1))
				Expect(err).ToNot(HaveOccurred())
				Expect(ra).To(Equal(RandomAccount.Address()))
			})

			It("should add the address to the whitelist map", func() {
				isWhitelisted, err := WalletProxy.WhitelistMap(nil, RandomAccount.Address())
				Expect(err).ToNot(HaveOccurred())
				Expect(isWhitelisted).To(BeTrue())
			})

			It("should emit whitelist addition event", func() {
				it, err := WalletProxy.FilterAddedToWhitelist(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(it.Next()).To(BeTrue())
				Expect(it.Next()).To(BeTrue())
				evt := it.Event
				Expect(it.Next()).To(BeFalse())
				Expect(evt.Sender).To(Equal(Controller.Address()))
				Expect(evt.Addresses).To(Equal([]common.Address{RandomAccount.Address()}))
			})

			When("I try to add another account to the whitelist", func() {
				It("should succeed", func() {
					tx, err := WalletProxy.SubmitWhitelistAddition(Owner.TransactOpts(), []common.Address{RandomAccount.Address()})
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeTrue())
				})
			})

			When("I try to initialize the whitelist", func() {
				It("should fail", func() {
					tx, err := WalletProxy.SetWhitelist(Owner.TransactOpts(ethertest.WithGasLimit(500000)), []common.Address{RandomAccount.Address()})
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeFalse())
				})
			})
		})

		When("controller tries to cancel whitelist addition with a wrong hash", func() {

			var txSuccessful bool

			BeforeEach(func() {
				tx, err := WalletProxy.CancelWhitelistAddition(Controller.TransactOpts(ethertest.WithGasLimit(500000)), common.BytesToHash(nil))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				txSuccessful = isSuccessful(tx)
			})

			It("should fail", func() {
				Expect(txSuccessful).To(BeFalse())
			})
		})

		When("the controller cancels adding to the whitelist", func() {
			var hash [32]uint8
			BeforeEach(func() {
				pwl, err := WalletProxy.PendingWhitelistAddition(nil)
				Expect(err).ToNot(HaveOccurred())
				hash, err = WalletProxy.CalculateHash(nil, pwl)
				Expect(err).ToNot(HaveOccurred())
				tx, err := WalletProxy.CancelWhitelistAddition(Controller.TransactOpts(), hash)
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})

			It("should emit CancelWhitelistAddition event", func() {
				it, err := WalletProxy.FilterCancelledWhitelistAddition(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(it.Next()).To(BeTrue())
				evt := it.Event
				Expect(it.Next()).To(BeFalse())
				Expect(evt.Sender).To(Equal(Controller.Address()))
				Expect(evt.Hash).To(Equal(hash))
			})

			It("should have empty pendingAddition list", func() {
				pending, err := WalletProxy.PendingWhitelistAddition(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(pending).To(HaveLen(0))
			})

			When("I try to initialize the whitelist", func() {
				It("should fail", func() {
					tx, err := WalletProxy.SetWhitelist(Owner.TransactOpts(ethertest.WithGasLimit(100000)), []common.Address{RandomAccount.Address()})
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeFalse())
				})
			})

			When("I try to add another account to the whitelist", func() {
				It("should succeed", func() {
					tx, err := WalletProxy.SubmitWhitelistAddition(Owner.TransactOpts(), []common.Address{RandomAccount.Address()})
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeTrue())
				})
			})
		})
	})

	When("I submit a list of addresses to Backend added to the whitelist after initialization", func() {
		BeforeEach(func() {
			tx, err := WalletProxy.SetWhitelist(Owner.TransactOpts(ethertest.WithGasLimit(500000)), []common.Address{common.HexToAddress("0x1")})
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeTrue())
		})

		It("should succeed", func() {
			addresses := []common.Address{}
			for i := 0; i < 20; i++ {
				addresses = append(addresses, ethertest.NewAccount().Address())
			}
			tx, err := WalletProxy.SubmitWhitelistAddition(Owner.TransactOpts(), addresses)
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeTrue())
		})
	})

	When("I submit an empty list of addresses to Backend added to the whitelist after initialization", func() {
		BeforeEach(func() {
			tx, err := WalletProxy.SetWhitelist(Owner.TransactOpts(ethertest.WithGasLimit(500000)), []common.Address{common.HexToAddress("0x1")})
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeTrue())
		})

		It("should fail", func() {
			addresses := []common.Address{}
			tx, err := WalletProxy.SubmitWhitelistAddition(Owner.TransactOpts(ethertest.WithGasLimit(100000)), addresses)
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeFalse())
		})
	})
})

var _ = Describe("whitelistRemoval", func() {

	BeforeEach(func() {
		BankAccount.MustTransfer(Backend, Controller.Address(), EthToWei(1))
	})

	When("3 addresses are in the whitelist", func() {

		BeforeEach(func() {
			tx, err := WalletProxy.SetWhitelist(Owner.TransactOpts(), []common.Address{RandomAccount.Address(), common.HexToAddress("0x1"), common.HexToAddress("0x2")})
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeTrue())
		})

		When("I submit removal of the random account from the whitelist", func() {

			BeforeEach(func() {
				tx, err := WalletProxy.SubmitWhitelistRemoval(Owner.TransactOpts(), []common.Address{RandomAccount.Address()})
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})

			It("should update the submittedWhitelistRemoval flag", func() {
				submitted, err := WalletProxy.SubmittedWhitelistRemoval(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(submitted).To(BeTrue())
			})

			It("should emit SubmitWhitelistRemoval event", func() {
				it, err := WalletProxy.FilterSubmittedWhitelistRemoval(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(it.Next()).To(BeTrue())
				evt := it.Event
				Expect(it.Next()).To(BeFalse())
				Expect(evt.Addresses).To(Equal([]common.Address{RandomAccount.Address()}))
				pwl, err := WalletProxy.PendingWhitelistRemoval(nil)
				Expect(err).ToNot(HaveOccurred())
				hash, err := WalletProxy.CalculateHash(nil, pwl)
				Expect(err).ToNot(HaveOccurred())
				Expect(evt.Hash).To(Equal(hash))
			})

			It("should add random account's address to pendingRemoval", func() {
				pending, err := WalletProxy.PendingWhitelistRemoval(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(pending).To(Equal([]common.Address{RandomAccount.Address()}))
			})

			When("I try to submit another list for removal", func() {

				It("should fail", func() {
					tx, err := WalletProxy.SubmitWhitelistRemoval(Owner.TransactOpts(ethertest.WithGasLimit(100000)), []common.Address{common.HexToAddress("0x3")})
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeFalse())
				})

			})

			When("I try to submit a list for addition while removal is pending", func() {

				It("should fail", func() {
					tx, err := WalletProxy.SubmitWhitelistAddition(Owner.TransactOpts(ethertest.WithGasLimit(100000)), []common.Address{common.HexToAddress("0x3")})
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeFalse())
				})

			})

			When("controller confirms the whitelist removal", func() {

				BeforeEach(func() {
					pwl, err := WalletProxy.PendingWhitelistRemoval(nil)
					Expect(err).ToNot(HaveOccurred())
					hash, err := WalletProxy.CalculateHash(nil, pwl)
					Expect(err).ToNot(HaveOccurred())
					tx, err := WalletProxy.ConfirmWhitelistRemoval(Controller.TransactOpts(), hash)
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeTrue())
				})

				It("should remove random account from the whitelist", func() {
					isWhitelisted, err := WalletProxy.WhitelistMap(nil, RandomAccount.Address())
					Expect(err).ToNot(HaveOccurred())
					Expect(isWhitelisted).To(BeFalse())
				})

				It("should emit WhitelistRemoval event", func() {
					it, err := WalletProxy.FilterRemovedFromWhitelist(nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(it.Next()).To(BeTrue())
					evt := it.Event
					Expect(it.Next()).To(BeFalse())
					Expect(evt.Sender).To(Equal(Controller.Address()))
					Expect(evt.Addresses).To(Equal([]common.Address{RandomAccount.Address()}))
				})

				It("should clear pendingRemoval", func() {
					pending, err := WalletProxy.PendingWhitelistRemoval(nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(pending).To(HaveLen(0))
				})

				It("should remove the random account's address from the Whitelist Array", func() {
					ra, err := WalletProxy.WhitelistArray(nil, big.NewInt(0))
					Expect(err).ToNot(HaveOccurred())
					Expect(ra).To(Equal(common.HexToAddress("0x2")))

					ra, err = WalletProxy.WhitelistArray(nil, big.NewInt(1))
					Expect(err).ToNot(HaveOccurred())
					Expect(ra).To(Equal(common.HexToAddress("0x1")))
				})

				When("I submit the last address for removal and controller confirms", func() {

					BeforeEach(func() {
						tx, err := WalletProxy.SubmitWhitelistRemoval(Owner.TransactOpts(), []common.Address{common.HexToAddress("0x1")})
						Expect(err).ToNot(HaveOccurred())
						Backend.Commit()
						Expect(isSuccessful(tx)).To(BeTrue())
					})

					BeforeEach(func() {
						pwl, err := WalletProxy.PendingWhitelistRemoval(nil)
						Expect(err).ToNot(HaveOccurred())
						hash, err := WalletProxy.CalculateHash(nil, pwl)
						Expect(err).ToNot(HaveOccurred())
						tx, err := WalletProxy.ConfirmWhitelistRemoval(Controller.TransactOpts(), hash)
						Expect(err).ToNot(HaveOccurred())
						Backend.Commit()
						Expect(isSuccessful(tx)).To(BeTrue())
					})

					It("should remove the 0x1 address from the whitelist", func() {
						isWhitelisted, err := WalletProxy.WhitelistMap(nil, common.HexToAddress("0x1"))
						Expect(err).ToNot(HaveOccurred())
						Expect(isWhitelisted).To(BeFalse())
					})

					It("should remove the 0x1 address from the Whitelist Array", func() {
						ra, err := WalletProxy.WhitelistArray(nil, big.NewInt(0))
						Expect(err).ToNot(HaveOccurred())
						Expect(ra).To(Equal(common.HexToAddress("0x2")))

					})

					It("should remove the 0x1 address from the Whitelist Array", func() {
						_, err := WalletProxy.WhitelistArray(nil, big.NewInt(1))
						Expect(err).To(HaveOccurred())
						Expect(err.Error()).To(ContainSubstring("abi: attempting to unmarshall"))
					})

				})

				When("controller tries to confirm whitelist removal with a wrong hash", func() {

					var txSuccessful bool

					BeforeEach(func() {
						tx, err := WalletProxy.ConfirmWhitelistRemoval(Controller.TransactOpts(ethertest.WithGasLimit(100000)), common.BytesToHash(nil))
						Expect(err).ToNot(HaveOccurred())
						Backend.Commit()
						txSuccessful = isSuccessful(tx)
					})

					It("should fail", func() {
						Expect(txSuccessful).To(BeFalse())
					})

				})

			})

			When("controller tries to cancel whitelist removal with a wrong hash", func() {

				var txSuccessful bool

				BeforeEach(func() {
					tx, err := WalletProxy.CancelWhitelistRemoval(Controller.TransactOpts(ethertest.WithGasLimit(100000)), common.BytesToHash(nil))
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					txSuccessful = isSuccessful(tx)
				})

				It("should fail", func() {
					Expect(txSuccessful).To(BeFalse())
				})
			})

			When("controller cancels the whitelist removal", func() {

				var hash [32]uint8
				BeforeEach(func() {
					pwl, err := WalletProxy.PendingWhitelistRemoval(nil)
					Expect(err).ToNot(HaveOccurred())
					hash, err = WalletProxy.CalculateHash(nil, pwl)
					Expect(err).ToNot(HaveOccurred())
					tx, err := WalletProxy.CancelWhitelistRemoval(Controller.TransactOpts(), hash)
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeTrue())
				})

				It("should emit CancelWhitelistRemoval event", func() {
					it, err := WalletProxy.FilterCancelledWhitelistRemoval(nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(it.Next()).To(BeTrue())
					evt := it.Event
					Expect(it.Next()).To(BeFalse())
					Expect(evt.Sender).To(Equal(Controller.Address()))
					Expect(evt.Hash).To(Equal(hash))

				})

				It("should have empty pendingRemoval list", func() {
					pending, err := WalletProxy.PendingWhitelistRemoval(nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(pending).To(HaveLen(0))
				})

				It("should keep the random account's address in the whitelist", func() {
					isWhitelisted, err := WalletProxy.WhitelistMap(nil, RandomAccount.Address())
					Expect(err).ToNot(HaveOccurred())
					Expect(isWhitelisted).To(BeTrue())
				})

				It("should keep the random account's address in the array", func() {
					ra, err := WalletProxy.WhitelistArray(nil, big.NewInt(0))
					Expect(err).ToNot(HaveOccurred())
					Expect(ra).To(Equal(RandomAccount.Address()))
				})

				When("I try to submit another list for removal", func() {

					It("should succeed", func() {
						tx, err := WalletProxy.SubmitWhitelistRemoval(Owner.TransactOpts(), []common.Address{RandomAccount.Address()})
						Expect(err).ToNot(HaveOccurred())
						Backend.Commit()
						Expect(isSuccessful(tx)).To(BeTrue())
					})

				})

			})

		})

	})

	When("I submit 20 addresses for removal from whitelist", func() {

		When("The Whitelist has been initialised", func() {

			BeforeEach(func() {
				tx, err := WalletProxy.SetWhitelist(Owner.TransactOpts(), []common.Address{RandomAccount.Address()})
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})

			It("should succeed", func() {
				addresses := []common.Address{}
				for i := 0; i < 20; i++ {
					addresses = append(addresses, ethertest.NewAccount().Address())
				}
				tx, err := WalletProxy.SubmitWhitelistRemoval(Owner.TransactOpts(), addresses)
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})

		})

		When("The Whitelist has NOT been initialised", func() {

			It("should fail", func() {
				addresses := []common.Address{}
				for i := 0; i < 20; i++ {
					addresses = append(addresses, ethertest.NewAccount().Address())
				}
				tx, err := WalletProxy.SubmitWhitelistRemoval(Owner.TransactOpts(ethertest.WithGasLimit(100000)), addresses)
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeFalse())
			})

		})

	})

})
