package wallet_test

import (
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tokencard/contracts/v3/test/shared"
	"github.com/tokencard/ethertest"
)

var _ = Describe("DailyLimit", func() {

	BeforeEach(func() {
		BankAccount.MustTransfer(Backend, Controller.Address(), EthToWei(1))
	})

	When("the contract just has been deployed", func() {
		It("should have initial daily limit of 10000$", func() {
			ll, err := Wallet.DailyLimitValue(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(ll.String()).To(Equal(MweiToWei(10000).String()))

			ll, err = Wallet.DailyLimitAvailable(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(ll.String()).To(Equal(MweiToWei(10000).String()))
		})

		It("should emit a UpdatedAvailableDailyLimit event", func() {
			it, err := Wallet.FilterUpdatedAvailableDailyLimit(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(it.Next()).To(BeTrue())
			evt := it.Event
			Expect(it.Next()).To(BeFalse())
			Expect(evt.Amount.String()).To(Equal(MweiToWei(10000).String()))
			initTime := Backend.Blockchain().CurrentHeader().Time + 24*60*60 - 10
			Expect(evt.NextReset.String()).To(Equal(big.NewInt(int64(initTime)).String()))
		})
	})

	When("Wallet deployer is registered in ENS and is set to be a random address", func() {
		BeforeEach(func() {
			// set WalletDeployer node owner
			tx, err := ENSRegistry.SetSubnodeOwner(BankAccount.TransactOpts(), EnsNode("tokencard.eth"), LabelHash("v3"), BankAccount.Address())
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			tx, err = ENSRegistry.SetSubnodeOwner(BankAccount.TransactOpts(), EnsNode("v3.tokencard.eth"), LabelHash("wallet-deployer"), BankAccount.Address())
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			// Register WalletDeployer with ENS and set it to be the Random account
			tx, err = ENSRegistry.SetResolver(BankAccount.TransactOpts(), WalletDeployerNode, ENSResolverAddress)
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeTrue())
			tx, err = ENSResolver.SetAddr(BankAccount.TransactOpts(), WalletDeployerNode, RandomAccount.Address())
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeTrue())
		})
		It("should fail", func() {
			tx, err := Wallet.MigrateDailyLimit(Controller.TransactOpts(ethertest.WithGasLimit(100000)), MweiToWei(1))
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeFalse())
			returnData, _ := ethCall(tx)
			Expect(string(returnData[len(returnData)-64:])).To(ContainSubstring("Only walletDeployer"))
		})
		It("should succeed", func() {
			tx, err := Wallet.MigrateDailyLimit(RandomAccount.TransactOpts(), MweiToWei(1))
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeTrue())
		})
	})

	When("a controller or 2FA tries to call setDailyLimit", func() {
		It("should fail", func() {
			tx, err := Wallet.SetDailyLimit(Controller.TransactOpts(ethertest.WithGasLimit(100000)), MweiToWei(1))
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeFalse())
			returnData, _ := ethCall(tx)
			Expect(string(returnData[len(returnData)-64:])).To(ContainSubstring("Only self"))
		})
	})

	When("the owner calls setDailyLimit", func() {
		It("should fail", func() {
			tx, err := Wallet.SetDailyLimit(Owner.TransactOpts(ethertest.WithGasLimit(65000)), MweiToWei(1))
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeFalse())
			returnData, _ := ethCall(tx)
			Expect(string(returnData[len(returnData)-64:])).To(ContainSubstring("Only self"))
		})
	})

	Describe("Changing daily limit", func() {

		When("Owner sets the DailyLimit to 1000 $USD via privileged 2FA", func() {
			BeforeEach(func() {
				a, err := abi.JSON(strings.NewReader(WALLET_ABI))
				Expect(err).ToNot(HaveOccurred())
				data, err := a.Pack("setDailyLimit", MweiToWei(1000))
				Expect(err).ToNot(HaveOccurred())

				batch := []byte(fmt.Sprintf("%s%s%s%s", WalletAddress, abi.U256(EthToWei(0)), abi.U256(big.NewInt(int64(len(data)))), data))

				nonce := big.NewInt(0)
				signature, err := SignData(nonce, batch, Owner.PrivKey())
				Expect(err).ToNot(HaveOccurred())

				tx, err := Wallet.ExecutePrivilegedRelayedTransaction(Controller.TransactOpts(), nonce, batch, signature)
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})

			It("should emit a UpdatedAvailableDailyLimit event", func() {
				it, err := Wallet.FilterUpdatedAvailableDailyLimit(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(it.Next()).To(BeTrue())
				Expect(it.Next()).To(BeTrue())
				evt := it.Event
				Expect(it.Next()).To(BeFalse())
				Expect(evt.Amount.String()).To(Equal(MweiToWei(1000).String()))
				initTime := Backend.Blockchain().CurrentHeader().Time + 24*60*60
				Expect(evt.NextReset.String()).To(Equal(big.NewInt(int64(initTime)).String()))
			})

			It("should emit a daily limit set event", func() {
				it, err := Wallet.FilterSetDailyLimit(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(it.Next()).To(BeTrue())
				evt := it.Event
				Expect(it.Next()).To(BeFalse())
				Expect(evt.Sender).To(Equal(WalletAddress))
				Expect(evt.Amount.String()).To(Equal(MweiToWei(1000).String()))
			})

			It("should lower the available amount to 1000 $USD", func() {
				av, err := Wallet.DailyLimitAvailable(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(av.String()).To(Equal(MweiToWei(1000).String()))
			})

			It("should have a new limit of 1000 $USD", func() {
				sl, err := Wallet.DailyLimitValue(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(sl.String()).To(Equal(MweiToWei(1000).String()))
			})

			When("the owner transfers 1$", func() {
				BeforeEach(func() {

					tx, err := Stablecoin.Credit(BankAccount.TransactOpts(), WalletAddress, MweiToWei(1))
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeTrue())

					tx, err = Wallet.Transfer(Owner.TransactOpts(), RandomAccount.Address(), StablecoinAddress, MweiToWei(1))
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeTrue())
				})

				It("should emit a UpdatedAvailableDailyLimit event", func() {
					it, err := Wallet.FilterUpdatedAvailableDailyLimit(nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(it.Next()).To(BeTrue())
					Expect(it.Next()).To(BeTrue())
					Expect(it.Next()).To(BeTrue())
					evt := it.Event
					Expect(it.Next()).To(BeFalse())
					Expect(evt.Amount.String()).To(Equal(MweiToWei(999).String()))
					initTime := Backend.Blockchain().CurrentHeader().Time + 24*60*60 - 20
					Expect(evt.NextReset.String()).To(Equal(big.NewInt(int64(initTime)).String()))
				})

				It("should have 999 available", func() {
					sa, err := Wallet.DailyLimitAvailable(nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(sa.String()).To(Equal(MweiToWei(999).String()))
				})

				When("I wait for 24 hours", func() {
					BeforeEach(func() {
						Backend.AdjustTime(time.Hour*24 + time.Second)
						Backend.Commit()
					})

					It("should have 1K $USD available for spending", func() {
						ll, err := Wallet.DailyLimitAvailable(nil)
						Expect(err).ToNot(HaveOccurred())
						Expect(ll.String()).To(Equal(MweiToWei(1000).String()))
					})
				})

				When("I increase the daily limit", func() {
					BeforeEach(func() {
						// set to 11K USD
						a, err := abi.JSON(strings.NewReader(WALLET_ABI))
						Expect(err).ToNot(HaveOccurred())
						data, err := a.Pack("setDailyLimit", GweiToWei(11))
						Expect(err).ToNot(HaveOccurred())

						batch := []byte(fmt.Sprintf("%s%s%s%s", WalletAddress, abi.U256(EthToWei(0)), abi.U256(big.NewInt(int64(len(data)))), data))

						nonce := big.NewInt(1)
						signature, err := SignData(nonce, batch, Owner.PrivKey())
						Expect(err).ToNot(HaveOccurred())

						tx, err := Wallet.ExecutePrivilegedRelayedTransaction(Controller.TransactOpts(), nonce, batch, signature)
						Expect(err).ToNot(HaveOccurred())
						Backend.Commit()
						Expect(isSuccessful(tx)).To(BeTrue())
					})

					It("should have 11K $USD available", func() {
						ll, err := Wallet.DailyLimitAvailable(nil)
						Expect(err).ToNot(HaveOccurred())
						Expect(ll.String()).To(Equal(GweiToWei(11).String()))
					})

					It("should emit a UpdatedAvailableDailyLimit event", func() {
						it, err := Wallet.FilterUpdatedAvailableDailyLimit(nil)
						Expect(err).ToNot(HaveOccurred())
						Expect(it.Next()).To(BeTrue())
						Expect(it.Next()).To(BeTrue())
						Expect(it.Next()).To(BeTrue())
						Expect(it.Next()).To(BeTrue())
						evt := it.Event
						Expect(it.Next()).To(BeFalse())
						Expect(evt.Amount.String()).To(Equal(GweiToWei(11).String()))
						initTime := Backend.Blockchain().CurrentHeader().Time + 24*60*60
						Expect(evt.NextReset.String()).To(Equal(big.NewInt(int64(initTime)).String()))
					})
				})

				When("I decrease the daily limit", func() {
					BeforeEach(func() {
						// set to 100 USD
						a, err := abi.JSON(strings.NewReader(WALLET_ABI))
						Expect(err).ToNot(HaveOccurred())
						data, err := a.Pack("setDailyLimit", MweiToWei(100))
						Expect(err).ToNot(HaveOccurred())

						batch := []byte(fmt.Sprintf("%s%s%s%s", WalletAddress, abi.U256(EthToWei(0)), abi.U256(big.NewInt(int64(len(data)))), data))

						nonce := big.NewInt(1)
						signature, err := SignData(nonce, batch, Owner.PrivKey())
						Expect(err).ToNot(HaveOccurred())

						tx, err := Wallet.ExecutePrivilegedRelayedTransaction(Controller.TransactOpts(), nonce, batch, signature)
						Expect(err).ToNot(HaveOccurred())
						Backend.Commit()
						Expect(isSuccessful(tx)).To(BeTrue())
					})

					It("should have 100 $USD available", func() {
						ll, err := Wallet.DailyLimitAvailable(nil)
						Expect(err).ToNot(HaveOccurred())
						Expect(ll.String()).To(Equal(MweiToWei(100).String()))
					})

					It("should emit a UpdatedAvailableDailyLimit event", func() {
						it, err := Wallet.FilterUpdatedAvailableDailyLimit(nil)
						Expect(err).ToNot(HaveOccurred())
						Expect(it.Next()).To(BeTrue())
						Expect(it.Next()).To(BeTrue())
						Expect(it.Next()).To(BeTrue())
						Expect(it.Next()).To(BeTrue())
						evt := it.Event
						Expect(it.Next()).To(BeFalse())
						Expect(evt.Amount.String()).To(Equal(MweiToWei(100).String()))
						initTime := Backend.Blockchain().CurrentHeader().Time + 24*60*60
						Expect(evt.NextReset.String()).To(Equal(big.NewInt(int64(initTime)).String()))
					})

				})

			})
		})

	})

})
