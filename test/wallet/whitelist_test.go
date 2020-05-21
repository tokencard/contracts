package wallet_test

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tokencard/contracts/v3/test/shared"
	"github.com/tokencard/ethertest"
)

var _ = Describe("initializeWhitelist", func() {

	When("a random account tries to add to the whitelist", func() {
		It("should fail", func() {
			tx, err := Wallet.AddToWhitelist(RandomAccount.TransactOpts(ethertest.WithGasLimit(100000)), []common.Address{RandomAccount.Address()})
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeFalse())
			returnData, _ := ethCall(tx)
			Expect(string(returnData[len(returnData)-64:])).To(ContainSubstring("Only self"))
		})
	})

	When("a random account tries to remove from the whitelist", func() {
		It("should fail", func() {
			tx, err := Wallet.RemoveFromWhitelist(RandomAccount.TransactOpts(ethertest.WithGasLimit(100000)), []common.Address{RandomAccount.Address()})
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeFalse())
			returnData, _ := ethCall(tx)
			Expect(string(returnData[len(returnData)-64:])).To(ContainSubstring("Only self"))
		})
	})

	When("I try to add to whitelist duplicate entries", func() {
		BeforeEach(func() {
			a, err := abi.JSON(strings.NewReader(WALLET_ABI))
			Expect(err).ToNot(HaveOccurred())
			data, err := a.Pack("addToWhitelist", []common.Address{common.HexToAddress("0x1"), common.HexToAddress("0x1")})
			Expect(err).ToNot(HaveOccurred())

			batch := []byte(fmt.Sprintf("%s%s%s%s", WalletAddress, abi.U256(EthToWei(0)), abi.U256(big.NewInt(int64(len(data)))), data))

			nonce := big.NewInt(0)
			signature, err := SignData(nonce, batch, Owner.PrivKey())
			Expect(err).ToNot(HaveOccurred())

			tx, err := Wallet.ExecutePrivilegedRelayedTransaction(Controller.TransactOpts(ethertest.WithGasLimit(100000)), nonce, batch, signature)
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeFalse())
			returnData, _ := ethCall(tx)
			Expect(string(returnData[len(returnData)-64:])).To(ContainSubstring("Address already whitelisted"))
		})
	})

	When("I try to add to whitelist the 0 address", func() {
		BeforeEach(func() {
			a, err := abi.JSON(strings.NewReader(WALLET_ABI))
			Expect(err).ToNot(HaveOccurred())
			data, err := a.Pack("addToWhitelist", []common.Address{common.HexToAddress("0x1"), common.HexToAddress("0x0")})
			Expect(err).ToNot(HaveOccurred())

			batch := []byte(fmt.Sprintf("%s%s%s%s", WalletAddress, abi.U256(EthToWei(0)), abi.U256(big.NewInt(int64(len(data)))), data))

			nonce := big.NewInt(0)
			signature, err := SignData(nonce, batch, Owner.PrivKey())
			Expect(err).ToNot(HaveOccurred())

			tx, err := Wallet.ExecutePrivilegedRelayedTransaction(Controller.TransactOpts(ethertest.WithGasLimit(100000)), nonce, batch, signature)
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeFalse())
			returnData, _ := ethCall(tx)
			Expect(string(returnData[len(returnData)-64:])).To(ContainSubstring("Contains 0 address"))
		})
	})

	When("I try to add to whitelist the owner address", func() {
		It("should fail", func() {
			a, err := abi.JSON(strings.NewReader(WALLET_ABI))
			Expect(err).ToNot(HaveOccurred())
			data, err := a.Pack("addToWhitelist", []common.Address{Owner.Address()})
			Expect(err).ToNot(HaveOccurred())

			batch := []byte(fmt.Sprintf("%s%s%s%s", WalletAddress, abi.U256(EthToWei(0)), abi.U256(big.NewInt(int64(len(data)))), data))

			nonce := big.NewInt(0)
			signature, err := SignData(nonce, batch, Owner.PrivKey())
			Expect(err).ToNot(HaveOccurred())

			tx, err := Wallet.ExecutePrivilegedRelayedTransaction(Controller.TransactOpts(ethertest.WithGasLimit(100000)), nonce, batch, signature)
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeFalse())
			returnData, _ := ethCall(tx)
			Expect(string(returnData[len(returnData)-64:])).To(ContainSubstring("Contains owner address"))
		})
	})

	When("I try to add or remove an empty list", func() {
		It("should fail", func() {
			a, err := abi.JSON(strings.NewReader(WALLET_ABI))
			Expect(err).ToNot(HaveOccurred())
			data, err := a.Pack("addToWhitelist", []common.Address{})
			Expect(err).ToNot(HaveOccurred())

			batch := []byte(fmt.Sprintf("%s%s%s%s", WalletAddress, abi.U256(EthToWei(0)), abi.U256(big.NewInt(int64(len(data)))), data))

			nonce := big.NewInt(0)
			signature, err := SignData(nonce, batch, Owner.PrivKey())
			Expect(err).ToNot(HaveOccurred())

			tx, err := Wallet.ExecutePrivilegedRelayedTransaction(Controller.TransactOpts(ethertest.WithGasLimit(100000)), nonce, batch, signature)
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeFalse())
			returnData, _ := ethCall(tx)
			Expect(string(returnData[len(returnData)-94:])).To(ContainSubstring("Empty list to be added to whitelist"))
		})

		It("should fail", func() {
			a, err := abi.JSON(strings.NewReader(WALLET_ABI))
			Expect(err).ToNot(HaveOccurred())
			data, err := a.Pack("removeFromWhitelist", []common.Address{})
			Expect(err).ToNot(HaveOccurred())

			batch := []byte(fmt.Sprintf("%s%s%s%s", WalletAddress, abi.U256(EthToWei(0)), abi.U256(big.NewInt(int64(len(data)))), data))

			nonce := big.NewInt(0)
			signature, err := SignData(nonce, batch, Owner.PrivKey())
			Expect(err).ToNot(HaveOccurred())

			tx, err := Wallet.ExecutePrivilegedRelayedTransaction(Controller.TransactOpts(ethertest.WithGasLimit(100000)), nonce, batch, signature)
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeFalse())
			returnData, _ := ethCall(tx)
			Expect(string(returnData[len(returnData)-94:])).To(ContainSubstring("Empty list to be removed from whitelist"))
		})
	})

	When("I try remove from an empty whiteList", func() {
		It("should fail", func() {
			a, err := abi.JSON(strings.NewReader(WALLET_ABI))
			Expect(err).ToNot(HaveOccurred())
			data, err := a.Pack("removeFromWhitelist", []common.Address{RandomAccount.Address()})
			Expect(err).ToNot(HaveOccurred())

			batch := []byte(fmt.Sprintf("%s%s%s%s", WalletAddress, abi.U256(EthToWei(0)), abi.U256(big.NewInt(int64(len(data)))), data))

			nonce := big.NewInt(0)
			signature, err := SignData(nonce, batch, Owner.PrivKey())
			Expect(err).ToNot(HaveOccurred())

			tx, err := Wallet.ExecutePrivilegedRelayedTransaction(Controller.TransactOpts(ethertest.WithGasLimit(100000)), nonce, batch, signature)
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeFalse())
			returnData, _ := ethCall(tx)
			Expect(string(returnData[len(returnData)-64:])).To(ContainSubstring("Whitelist is empty"))
		})
	})

	When("I add to the whitelist a random account's address", func() {
		BeforeEach(func() {
			a, err := abi.JSON(strings.NewReader(WALLET_ABI))
			Expect(err).ToNot(HaveOccurred())
			data, err := a.Pack("addToWhitelist", []common.Address{RandomAccount.Address()})
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

		It("should add the random account to the whitelist", func() {
			isWhitelisted, err := Wallet.WhitelistMap(nil, RandomAccount.Address())
			Expect(err).ToNot(HaveOccurred())
			Expect(isWhitelisted).To(BeTrue())
		})

		It("should emit AddedToWhitelist event", func() {
			it, err := Wallet.FilterAddedToWhitelist(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(it.Next()).To(BeTrue())
			evt := it.Event
			Expect(it.Next()).To(BeFalse())
			Expect(evt.Sender).To(Equal(WalletAddress))
			Expect(evt.Addresses).To(Equal([]common.Address{RandomAccount.Address()}))
		})

		It("should add an entry to the Whitelist Array", func() {
			a, err := Wallet.WhitelistArray(nil, big.NewInt(0))
			Expect(err).ToNot(HaveOccurred())
			Expect(a).To(Equal(RandomAccount.Address()))
		})

		When("I add more addresses to the whitelist", func() {

			BeforeEach(func() {
				a, err := abi.JSON(strings.NewReader(WALLET_ABI))
				Expect(err).ToNot(HaveOccurred())
				data, err := a.Pack("addToWhitelist", []common.Address{common.HexToAddress("0x1"), common.HexToAddress("0x2")})
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

			When("I try remove a non-existent address", func() {
				It("should fail", func() {
					a, err := abi.JSON(strings.NewReader(WALLET_ABI))
					Expect(err).ToNot(HaveOccurred())
					data, err := a.Pack("removeFromWhitelist", []common.Address{common.HexToAddress("0x0")})
					Expect(err).ToNot(HaveOccurred())

					batch := []byte(fmt.Sprintf("%s%s%s%s", WalletAddress, abi.U256(EthToWei(0)), abi.U256(big.NewInt(int64(len(data)))), data))

					nonce := big.NewInt(2)
					signature, err := SignData(nonce, batch, Owner.PrivKey())
					Expect(err).ToNot(HaveOccurred())

					tx, err := Wallet.ExecutePrivilegedRelayedTransaction(Controller.TransactOpts(ethertest.WithGasLimit(100000)), nonce, batch, signature)
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeFalse())
					returnData, _ := ethCall(tx)
					Expect(string(returnData[len(returnData)-64:])).To(ContainSubstring("Address not whitelisted"))
				})
			})

			When("I try to remove from the whitelist and provide duplicate entries", func() {
				BeforeEach(func() {
					a, err := abi.JSON(strings.NewReader(WALLET_ABI))
					Expect(err).ToNot(HaveOccurred())
					data, err := a.Pack("removeFromWhitelist", []common.Address{common.HexToAddress("0x1"), common.HexToAddress("0x1")})
					Expect(err).ToNot(HaveOccurred())

					batch := []byte(fmt.Sprintf("%s%s%s%s", WalletAddress, abi.U256(EthToWei(0)), abi.U256(big.NewInt(int64(len(data)))), data))

					nonce := big.NewInt(2)
					signature, err := SignData(nonce, batch, Owner.PrivKey())
					Expect(err).ToNot(HaveOccurred())

					tx, err := Wallet.ExecutePrivilegedRelayedTransaction(Controller.TransactOpts(ethertest.WithGasLimit(100000)), nonce, batch, signature)
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeFalse())
					returnData, _ := ethCall(tx)
					Expect(string(returnData[len(returnData)-64:])).To(ContainSubstring("Address not whitelisted"))
				})
			})

			When("I remove 2 addressses from the whitelist", func() {

				BeforeEach(func() {
					a, err := abi.JSON(strings.NewReader(WALLET_ABI))
					Expect(err).ToNot(HaveOccurred())
					data, err := a.Pack("removeFromWhitelist", []common.Address{common.HexToAddress("0x2"), RandomAccount.Address()})
					Expect(err).ToNot(HaveOccurred())

					batch := []byte(fmt.Sprintf("%s%s%s%s", WalletAddress, abi.U256(EthToWei(0)), abi.U256(big.NewInt(int64(len(data)))), data))

					nonce := big.NewInt(2)
					signature, err := SignData(nonce, batch, Owner.PrivKey())
					Expect(err).ToNot(HaveOccurred())

					tx, err := Wallet.ExecutePrivilegedRelayedTransaction(Controller.TransactOpts(), nonce, batch, signature)
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeTrue())
				})

				It("should leave 0x1 in the whitelist", func() {
					isWhitelisted, err := Wallet.WhitelistMap(nil, common.HexToAddress("0x1"))
					Expect(err).ToNot(HaveOccurred())
					Expect(isWhitelisted).To(BeTrue())
				})

				It("should keep 0x1 in the Whitelist Array", func() {
					ra, err := Wallet.WhitelistArray(nil, big.NewInt(0))
					Expect(err).ToNot(HaveOccurred())
					Expect(ra).To(Equal(common.HexToAddress("0x1")))
				})

				It("should emit a WhitelistRemoval event", func() {
					it, err := Wallet.FilterRemovedFromWhitelist(nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(it.Next()).To(BeTrue())
					evt := it.Event
					Expect(it.Next()).To(BeFalse())
					Expect(evt.Sender).To(Equal(WalletAddress))
					Expect(evt.Addresses).To(Equal([]common.Address{common.HexToAddress("0x2"), RandomAccount.Address()}))
				})

				It("should remove 0x2 address from the whitelist", func() {
					isWhitelisted, err := Wallet.WhitelistMap(nil, common.HexToAddress("0x2"))
					Expect(err).ToNot(HaveOccurred())
					Expect(isWhitelisted).To(BeFalse())
				})

				It("should remove the random account from the whitelist", func() {
					isWhitelisted, err := Wallet.WhitelistMap(nil, RandomAccount.Address())
					Expect(err).ToNot(HaveOccurred())
					Expect(isWhitelisted).To(BeFalse())
				})

				It("should reduce the array size", func() {
					_, err := Wallet.WhitelistArray(nil, big.NewInt(1))
					Expect(err).To(HaveOccurred())
					Expect(err.Error()).To(ContainSubstring("abi: attempting to unmarshall"))
				})

			})

		})

	})

})
