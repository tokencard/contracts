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

var _ = Describe("2FA", func() {

	It("should be true", func() {
		oo, err := Wallet.Monolith2FA(nil)
		Expect(err).ToNot(HaveOccurred())
		Expect(oo).To(BeTrue())
	})

	It("should NOT allow a non-owner to set Monolith 2FA", func() {
		tx, err := Wallet.SetMonolith2FA(Controller.TransactOpts(ethertest.WithGasLimit(60000)))
		Expect(err).ToNot(HaveOccurred())
		Backend.Commit()
		Expect(isSuccessful(tx)).To(BeFalse())
		returnData, _ := ethCall(tx)
		Expect(string(returnData[len(returnData)-64:])).To(ContainSubstring("sender is not owner"))
	})

	It("should fail if monolith 2FA is already enabled", func() {
		tx, err := Wallet.SetMonolith2FA(Owner.TransactOpts(ethertest.WithGasLimit(60000)))
		Expect(err).ToNot(HaveOccurred())
		Backend.Commit()
		Expect(isSuccessful(tx)).To(BeFalse())
		returnData, _ := ethCall(tx)
		Expect(string(returnData[len(returnData)-64:])).To(ContainSubstring("Monolith 2FA already enabled"))
	})

	It("should fail if not called by itself i.e. relayed Tx", func() {
		tx, err := Wallet.SetPersonal2FA(Controller.TransactOpts(ethertest.WithGasLimit(60000)), common.HexToAddress("0x1"))
		Expect(err).ToNot(HaveOccurred())
		Backend.Commit()
		Expect(isSuccessful(tx)).To(BeFalse())
		returnData, _ := ethCall(tx)
		Expect(string(returnData[len(returnData)-64:])).To(ContainSubstring("Only self"))
	})

	It("should fail if privileged mode is not used", func() {
		a, err := abi.JSON(strings.NewReader(WALLET_2FA_ABI))
		Expect(err).ToNot(HaveOccurred())
		data, err := a.Pack("setPersonal2FA", RandomAccount.Address())
		Expect(err).ToNot(HaveOccurred())

		batch := []byte(fmt.Sprintf("%s%s%s%s", WalletAddress, abi.U256(EthToWei(0)), abi.U256(big.NewInt(int64(len(data)))), data))

		nonce := big.NewInt(0)
		signature, err := SignData(nonce, batch, Owner.PrivKey())
		Expect(err).ToNot(HaveOccurred())

		tx, err := Wallet.ExecuteRelayedTransaction(Controller.TransactOpts(ethertest.WithGasLimit(500000)), nonce, batch, signature)
		Expect(err).ToNot(HaveOccurred())
		Backend.Commit()
		Expect(isSuccessful(tx)).To(BeFalse())

		returnData, _ := ethCall(tx)
		Expect(string(returnData[len(returnData)-94:])).To(ContainSubstring("Setting a personal 2FA requires privileged mode"))
	})

	It("should fail if trying to set 0x0 as the personal 2FA", func() {
		a, err := abi.JSON(strings.NewReader(WALLET_2FA_ABI))
		Expect(err).ToNot(HaveOccurred())
		data, err := a.Pack("setPersonal2FA", common.HexToAddress("0x0"))
		Expect(err).ToNot(HaveOccurred())

		batch := []byte(fmt.Sprintf("%s%s%s%s", WalletAddress, abi.U256(EthToWei(0)), abi.U256(big.NewInt(int64(len(data)))), data))

		nonce := big.NewInt(0)
		signature, err := SignData(nonce, batch, Owner.PrivKey())
		Expect(err).ToNot(HaveOccurred())

		tx, err := Wallet.ExecutePrivilegedRelayedTransaction(Controller.TransactOpts(ethertest.WithGasLimit(500000)), nonce, batch, signature)
		Expect(err).ToNot(HaveOccurred())
		Backend.Commit()
		Expect(isSuccessful(tx)).To(BeFalse())

		returnData, _ := ethCall(tx)
		Expect(string(returnData[len(returnData)-64:])).To(ContainSubstring("2FA cannot be the 0 address"))
	})

	It("should fail if trying to set the wallet address as the personal 2FA", func() {
		a, err := abi.JSON(strings.NewReader(WALLET_2FA_ABI))
		Expect(err).ToNot(HaveOccurred())
		data, err := a.Pack("setPersonal2FA", WalletAddress)
		Expect(err).ToNot(HaveOccurred())

		batch := []byte(fmt.Sprintf("%s%s%s%s", WalletAddress, abi.U256(EthToWei(0)), abi.U256(big.NewInt(int64(len(data)))), data))

		nonce := big.NewInt(0)
		signature, err := SignData(nonce, batch, Owner.PrivKey())
		Expect(err).ToNot(HaveOccurred())

		tx, err := Wallet.ExecutePrivilegedRelayedTransaction(Controller.TransactOpts(ethertest.WithGasLimit(500000)), nonce, batch, signature)
		Expect(err).ToNot(HaveOccurred())
		Backend.Commit()
		Expect(isSuccessful(tx)).To(BeFalse())

		returnData, _ := ethCall(tx)
		Expect(string(returnData[len(returnData)-64:])).To(ContainSubstring("2FA cannot be the wallet address"))
	})

	When("the owner wants to set the daily limit ", func() {
		It("should fail when trying to do it directly", func() {
			tx, err := Wallet.SetDailyLimit(Owner.TransactOpts(ethertest.WithGasLimit(80000)), MweiToWei(12000))
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeFalse())
			returnData, _ := ethCall(tx)
			Expect(string(returnData[len(returnData)-64:])).To(ContainSubstring("Only self"))
		})

		When("Monolith opt-out is NOT enabled and Monolith 2FA relays", func() {
			It("should succeed when privileged mode on", func() {
				a, err := abi.JSON(strings.NewReader(WALLET_ABI))
				Expect(err).ToNot(HaveOccurred())
				data, err := a.Pack("setDailyLimit", MweiToWei(12000))
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

			It("should fail when when privileged mode is off", func() {
				a, err := abi.JSON(strings.NewReader(WALLET_ABI))
				Expect(err).ToNot(HaveOccurred())
				data, err := a.Pack("setDailyLimit", MweiToWei(12000))
				Expect(err).ToNot(HaveOccurred())

				batch := []byte(fmt.Sprintf("%s%s%s%s", WalletAddress, abi.U256(EthToWei(0)), abi.U256(big.NewInt(int64(len(data)))), data))

				nonce := big.NewInt(0)
				signature, err := SignData(nonce, batch, Owner.PrivKey())
				Expect(err).ToNot(HaveOccurred())

				tx, err := Wallet.ExecuteRelayedTransaction(Controller.TransactOpts(ethertest.WithGasLimit(150000)), nonce, batch, signature)
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeFalse())
				returnData, _ := ethCall(tx)
				Expect(string(returnData[len(returnData)-94:])).To(ContainSubstring("Setting the daily limit requires privileged mode"))
			})
		})

		When("Monolith opt-out is enabled (personal2FA is set) and the random account is set as the personal 2FA", func() {
			BeforeEach(func() {
				a, err := abi.JSON(strings.NewReader(WALLET_2FA_ABI))
				Expect(err).ToNot(HaveOccurred())
				data, err := a.Pack("setPersonal2FA", RandomAccount.Address())
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

			It("should be false", func() {
				m2fa, err := Wallet.Monolith2FA(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(m2fa).To(BeFalse())
			})

			It("should set the random account address", func() {
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
				Expect(evt.Sender).To(Equal(WalletAddress))
				Expect(evt.P2FA).To(Equal(RandomAccount.Address()))
			})

			It("should fail when Monolith 2FA tries to set the limit", func() {
				a, err := abi.JSON(strings.NewReader(WALLET_ABI))
				Expect(err).ToNot(HaveOccurred())
				data, err := a.Pack("setDailyLimit", MweiToWei(12000))
				Expect(err).ToNot(HaveOccurred())

				batch := []byte(fmt.Sprintf("%s%s%s%s", WalletAddress, abi.U256(EthToWei(0)), abi.U256(big.NewInt(int64(len(data)))), data))

				nonce := big.NewInt(1)
				signature, err := SignData(nonce, batch, Owner.PrivKey())
				Expect(err).ToNot(HaveOccurred())

				tx, err := Wallet.ExecutePrivilegedRelayedTransaction(Controller.TransactOpts(ethertest.WithGasLimit(100000)), nonce, batch, signature)
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeFalse())
				returnData, _ := ethCall(tx)
				Expect(string(returnData[len(returnData)-64:])).To(ContainSubstring("Only personal 2FA"))
			})

			When("the personal 2FA sets the limit", func() {
				BeforeEach(func() {
					a, err := abi.JSON(strings.NewReader(WALLET_ABI))
					Expect(err).ToNot(HaveOccurred())
					data, err := a.Pack("setDailyLimit", MweiToWei(12000))
					Expect(err).ToNot(HaveOccurred())

					batch := []byte(fmt.Sprintf("%s%s%s%s", WalletAddress, abi.U256(EthToWei(0)), abi.U256(big.NewInt(int64(len(data)))), data))

					nonce := big.NewInt(1)
					signature, err := SignData(nonce, batch, Owner.PrivKey())
					Expect(err).ToNot(HaveOccurred())

					tx, err := Wallet.ExecutePrivilegedRelayedTransaction(RandomAccount.TransactOpts(), nonce, batch, signature)
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeTrue())
				})

				It("should have 12K $USD available for spending", func() {
					tl, err := Wallet.DailyLimitAvailable(nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(tl.String()).To(Equal(MweiToWei(12000).String()))
				})

			})

			When("the owner sets Monolith as 2FA again", func() {
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

				It("should succeed when Monolith 2FA sets the limit update", func() {
					a, err := abi.JSON(strings.NewReader(WALLET_ABI))
					Expect(err).ToNot(HaveOccurred())
					data, err := a.Pack("setDailyLimit", MweiToWei(11000))
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

				It("should fail when the random account (ex-personal2FA) tries to set the limit", func() {
					a, err := abi.JSON(strings.NewReader(WALLET_ABI))
					Expect(err).ToNot(HaveOccurred())
					data, err := a.Pack("setDailyLimit", MweiToWei(12000))
					Expect(err).ToNot(HaveOccurred())

					batch := []byte(fmt.Sprintf("%s%s%s%s", WalletAddress, abi.U256(EthToWei(0)), abi.U256(big.NewInt(int64(len(data)))), data))

					nonce := big.NewInt(0)
					signature, err := SignData(nonce, batch, Owner.PrivKey())
					Expect(err).ToNot(HaveOccurred())

					tx, err := Wallet.ExecutePrivilegedRelayedTransaction(RandomAccount.TransactOpts(ethertest.WithGasLimit(100000)), nonce, batch, signature)
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeFalse())
					returnData, _ := ethCall(tx)
					Expect(string(returnData[len(returnData)-64:])).To(ContainSubstring("Only Monolith 2FA"))
				})
			})

		})
	})

})

const WALLET_2FA_ABI = `[
    {
        "constant": false,
        "inputs": [
            {
                "name": "_p2FA",
                "type": "address"
            }
        ],
        "name": "setPersonal2FA",
        "outputs": [],
        "payable": false,
        "stateMutability": "nonpayable",
        "type": "function"
    }
]`
