package wallet_test

import (
	"context"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tokencard/contracts/v3/test/shared"
	"github.com/tokencard/ethertest"
)

var _ = Describe("executePrivilegedRelayedTransaction", func() {

	BeforeEach(func() {
		BankAccount.MustTransfer(Backend, WalletProxyAddress, EthToWei(4))
	})

	When("a random account tries to relay", func() {
		It("should fail", func() {
			a, err := abi.JSON(strings.NewReader(WALLET_ABI))
			Expect(err).ToNot(HaveOccurred())
			data, err := a.Pack("transfer", RandomAccount.Address(), common.HexToAddress("0x0"), EthToWei(1))
			Expect(err).ToNot(HaveOccurred())

			batch := []byte(fmt.Sprintf("%s%s%s%s", WalletProxyAddress, abi.U256(EthToWei(0)), abi.U256(big.NewInt(int64(len(data)))), data))

			nonce := big.NewInt(0)
			chainId := big.NewInt(1337)
			signature, err := SignData(chainId, WalletProxyAddress, nonce, batch, Owner.PrivKey())
			Expect(err).ToNot(HaveOccurred())

			tx, err := WalletProxy.ExecutePrivilegedRelayedTransaction(RandomAccount.TransactOpts(ethertest.WithGasLimit(500000)), nonce, batch, signature)
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeFalse())
			returnData, _ := ethCall(tx)
			Expect(string(returnData[len(returnData)-64:])).To(ContainSubstring("sender is not a Monolith 2FA"))
		})
	})

	When("2FA tries to relay a transaction signed by a random account", func() {
		It("should fail", func() {
			a, err := abi.JSON(strings.NewReader(WALLET_ABI))
			Expect(err).ToNot(HaveOccurred())
			privateKey, _ := crypto.GenerateKey()
			data, err := a.Pack("transfer", RandomAccount.Address(), common.HexToAddress("0x0"), EthToWei(1))
			Expect(err).ToNot(HaveOccurred())

			batch := []byte(fmt.Sprintf("%s%s%s%s", WalletProxyAddress, abi.U256(EthToWei(0)), abi.U256(big.NewInt(int64(len(data)))), data))

			nonce := big.NewInt(0)
			chainId := big.NewInt(1337)
			signature, err := SignData(chainId, WalletProxyAddress, nonce, batch, privateKey)
			Expect(err).ToNot(HaveOccurred())

			tx, err := WalletProxy.ExecutePrivilegedRelayedTransaction(Controller.TransactOpts(ethertest.WithGasLimit(500000)), nonce, batch, signature)
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeFalse())
			returnData, _ := ethCall(tx)
			Expect(string(returnData[len(returnData)-64:])).To(ContainSubstring("invalid signature"))
		})
	})

	When("2FA tries to relay two owner-signed transactions: send value(no data) + transfer() with a value above the daily limiy", func() {

		BeforeEach(func() {
			// Add  ERC20 token to tokenWhitelist and update token rate to be equal to 1 (1 token = 1 ETH)
			tx, err := TokenWhitelist.AddTokens(
				ControllerAdmin.TransactOpts(),
				[]common.Address{ERC20Contract1Address},
				StringsToByte32(
					"ERC1",
				),
				[]*big.Int{
					DecimalsToMagnitude(big.NewInt(18)),
				},
				[]bool{true},
				[]bool{true},
				big.NewInt(20180913153211),
			)
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeTrue())
			tx, err = TokenWhitelist.UpdateTokenRate(ControllerAdmin.TransactOpts(), ERC20Contract1Address, EthToWei(1), big.NewInt(20180913153211))
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeTrue())
			// Reduce the daily limit to 100$
			tx, err = WalletProxy.SubmitDailyLimitUpdate(Owner.TransactOpts(), MweiToWei(100))
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeTrue())
			// credit wallet with 1 token
			tx, err = ERC20Contract1.Credit(BankAccount.TransactOpts(), WalletProxyAddress, EthToWei(1))
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeTrue())
		})

		It("should reduce the available daily balance", func() {
			av, err := WalletProxy.DailyLimitAvailable(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(av.String()).To(Equal(MweiToWei(100).String()))
		})

		When("the the privileged mode is used they succeed", func() {
			BeforeEach(func() {
				// send value (1 ETH) > 100$
				data1 := fmt.Sprintf("%s%s%s", RandomAccount.Address(), abi.U256(EthToWei(1)), abi.U256(big.NewInt(0)))
				// use wallet's ERC20 transfer to transfer the token
				a, err := abi.JSON(strings.NewReader(WALLET_ABI))
				Expect(err).ToNot(HaveOccurred())
				data2, err := a.Pack("transfer", RandomAccount.Address(), ERC20Contract1Address, EthToWei(1))
				Expect(err).ToNot(HaveOccurred())

				batch := []byte(fmt.Sprintf("%s%s%s%s%s", data1, WalletProxyAddress, abi.U256(EthToWei(0)), abi.U256(big.NewInt(int64(len(data2)))), data2))

				nonce := big.NewInt(0)
				chainId := big.NewInt(1337)
				signature, err := SignData(chainId, WalletProxyAddress, nonce, batch, Owner.PrivKey())
				Expect(err).ToNot(HaveOccurred())

				tx, err := WalletProxy.ExecutePrivilegedRelayedTransaction(Controller.TransactOpts(), nonce, batch, signature)
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})

			It("should emit an ExecutedRelayedTransaction event", func() {
				it, err := WalletProxy.FilterExecutedRelayedTransaction(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(it.Next()).To(BeTrue())
				evt := it.Event
				Expect(it.Next()).To(BeFalse())
				data1 := []byte(fmt.Sprintf("%s%s%s", RandomAccount.Address(), abi.U256(EthToWei(1)), abi.U256(big.NewInt(0))))
				a, err := abi.JSON(strings.NewReader(WALLET_ABI))
				Expect(err).ToNot(HaveOccurred())
				data2, err := a.Pack("transfer", RandomAccount.Address(), ERC20Contract1Address, EthToWei(1))
				batchData := []byte(fmt.Sprintf("%s%s%s%s%s", data1, WalletProxyAddress, abi.U256(EthToWei(0)), abi.U256(big.NewInt(int64(len(data2)))), data2))
				Expect(evt.Data).To(Equal(batchData))
				Expect(evt.Privileged).To(Equal(true))
			})

			It("should emit 2 ExecutedTransaction events", func() {
				it, err := WalletProxy.FilterExecutedTransaction(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(it.Next()).To(BeTrue())
				evt := it.Event
				Expect(it.Next()).To(BeTrue())
				Expect(evt.Destination).To(Equal(RandomAccount.Address()))
				Expect(evt.Value.String()).To(Equal(EthToWei(1).String()))
				Expect(evt.Data).To(Equal(common.Hex2Bytes("")))
				Expect(evt.ReturnData).To(Equal(common.Hex2Bytes("")))
				evt = it.Event
				Expect(it.Next()).To(BeFalse())
				a, err := abi.JSON(strings.NewReader(WALLET_ABI))
				Expect(err).ToNot(HaveOccurred())
				data, err := a.Pack("transfer", RandomAccount.Address(), ERC20Contract1Address, EthToWei(1))
				Expect(evt.Destination).To(Equal(WalletProxyAddress))
				Expect(evt.Value.String()).To(Equal("0"))
				Expect(evt.Data).To(Equal(data))
				Expect(evt.ReturnData).To(Equal(common.Hex2Bytes("")))
			})

			It("should decrease the wallet's ETH balance ", func() {
				b, err := Backend.BalanceAt(context.Background(), WalletProxyAddress, nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(b.String()).To(Equal(EthToWei(3).String()))
			})

			It("should decrease TKN balance of the wallet", func() {
				b, err := ERC20Contract1.BalanceOf(nil, WalletProxyAddress)
				Expect(err).ToNot(HaveOccurred())
				Expect(b.String()).To(Equal("0"))
			})

			It("should reset privileged back to false", func() {
				p, err := WalletProxy.Privileged(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(p).To(BeFalse())
			})

			It("should fail when trying to replay", func() {
				a, err := abi.JSON(strings.NewReader(WALLET_ABI))
				Expect(err).ToNot(HaveOccurred())
				data, err := a.Pack("transfer", RandomAccount.Address(), common.HexToAddress("0x0"), EthToWei(1))
				Expect(err).ToNot(HaveOccurred())

				batch := []byte(fmt.Sprintf("%s%s%s%s", WalletProxyAddress, abi.U256(EthToWei(0)), abi.U256(big.NewInt(int64(len(data)))), data))

				nonce := big.NewInt(0)
				chainId := big.NewInt(1337)
				signature, err := SignData(chainId, WalletProxyAddress, nonce, batch, Owner.PrivKey())
				Expect(err).ToNot(HaveOccurred())

				tx, err := WalletProxy.ExecutePrivilegedRelayedTransaction(Controller.TransactOpts(ethertest.WithGasLimit(500000)), nonce, batch, signature)
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeFalse())
				returnData, _ := ethCall(tx)
				Expect(string(returnData[len(returnData)-64:])).To(ContainSubstring("tx replay"))
			})

			It("should succeed when increasing the nonce", func() {
				a, err := abi.JSON(strings.NewReader(WALLET_ABI))
				Expect(err).ToNot(HaveOccurred())
				data, err := a.Pack("transfer", RandomAccount.Address(), common.HexToAddress("0x0"), EthToWei(1))
				Expect(err).ToNot(HaveOccurred())

				batch := []byte(fmt.Sprintf("%s%s%s%s", WalletProxyAddress, abi.U256(EthToWei(0)), abi.U256(big.NewInt(int64(len(data)))), data))

				nonce := big.NewInt(1)
				chainId := big.NewInt(1337)
				signature, err := SignData(chainId, WalletProxyAddress, nonce, batch, Owner.PrivKey())
				Expect(err).ToNot(HaveOccurred())

				tx, err := WalletProxy.ExecutePrivilegedRelayedTransaction(Controller.TransactOpts(), nonce, batch, signature)
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})
		})

		When("the the non-privileged mode is used each one of them fails separately due to having exceeded the limit ", func() {
			It("should fail when transfering 1 ETH (>100$)", func() {
				// send value (1 ETH) > 100$
				a, err := abi.JSON(strings.NewReader(WALLET_ABI))
				Expect(err).ToNot(HaveOccurred())
				data, err := a.Pack("transfer", RandomAccount.Address(), common.HexToAddress("0x0"), EthToWei(1))
				Expect(err).ToNot(HaveOccurred())

				batch := []byte(fmt.Sprintf("%s%s%s%s", WalletProxyAddress, abi.U256(EthToWei(1)), abi.U256(big.NewInt(int64(len(data)))), data))

				nonce := big.NewInt(0)
				chainId := big.NewInt(1337)
				signature, err := SignData(chainId, WalletProxyAddress, nonce, batch, Owner.PrivKey())
				Expect(err).ToNot(HaveOccurred())

				tx, err := WalletProxy.ExecuteRelayedTransaction(Controller.TransactOpts(ethertest.WithGasLimit(200000)), nonce, batch, signature)
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeFalse())
				returnData, _ := ethCall(tx)
				Expect(string(returnData[len(returnData)-64:])).To(ContainSubstring("available smaller than amount"))
			})

			It("should fail when transfering 1 token (>100$)", func() {
				// use wallet's ERC20 transfer to transfer the token
				a, err := abi.JSON(strings.NewReader(WALLET_ABI))
				Expect(err).ToNot(HaveOccurred())
				data, err := a.Pack("transfer", RandomAccount.Address(), ERC20Contract1Address, EthToWei(1))
				Expect(err).ToNot(HaveOccurred())
				batch := []byte(fmt.Sprintf("%s%s%s%s", WalletProxyAddress, abi.U256(EthToWei(1)), abi.U256(big.NewInt(int64(len(data)))), data))

				nonce := big.NewInt(0)
				chainId := big.NewInt(1337)
				signature, err := SignData(chainId, WalletProxyAddress, nonce, batch, Owner.PrivKey())
				Expect(err).ToNot(HaveOccurred())

				tx, err := WalletProxy.ExecuteRelayedTransaction(Controller.TransactOpts(ethertest.WithGasLimit(200000)), nonce, batch, signature)
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeFalse())
				returnData, _ := ethCall(tx)
				Expect(string(returnData[len(returnData)-64:])).To(ContainSubstring("available smaller than amount"))

			})

		})
	})

	When("transfering to 0 address)", func() {
		It("should return the error string emitted by require", func() {
			a, err := abi.JSON(strings.NewReader(WALLET_ABI))
			Expect(err).ToNot(HaveOccurred())
			data, err := a.Pack("transfer", common.HexToAddress("0x0"), common.HexToAddress("0x0"), EthToWei(1))
			Expect(err).ToNot(HaveOccurred())

			batch := []byte(fmt.Sprintf("%s%s%s%s", WalletProxyAddress, abi.U256(EthToWei(0)), abi.U256(big.NewInt(int64(len(data)))), data))

			nonce := big.NewInt(0)
			chainId := big.NewInt(1337)
			signature, err := SignData(chainId, WalletProxyAddress, nonce, batch, Owner.PrivKey())
			Expect(err).ToNot(HaveOccurred())

			tx, err := WalletProxy.ExecutePrivilegedRelayedTransaction(Controller.TransactOpts(ethertest.WithGasLimit(500000)), nonce, batch, signature)
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeFalse())

			returnData, _ := ethCall(tx)
			Expect(string(returnData[len(returnData)-64:])).To(ContainSubstring("destination=0"))
		})
	})

	When("the owner increases the nonce before the relayer relayes the transaction", func() {
		BeforeEach(func() {
			a, err := abi.JSON(strings.NewReader(WALLET_ABI))
			Expect(err).ToNot(HaveOccurred())
			data, err := a.Pack("transfer", common.HexToAddress("0x0"), common.HexToAddress("0x0"), EthToWei(1))
			Expect(err).ToNot(HaveOccurred())

			batch := []byte(fmt.Sprintf("%s%s%s%s", WalletProxyAddress, abi.U256(EthToWei(0)), abi.U256(big.NewInt(int64(len(data)))), data))

			nonce := big.NewInt(0)
			chainId := big.NewInt(1337)
			signature, err := SignData(chainId, WalletProxyAddress, nonce, batch, Owner.PrivKey())
			Expect(err).ToNot(HaveOccurred())

			tx, err := WalletProxy.IncreaseRelayNonce(Owner.TransactOpts())
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeTrue())

			tx, err = WalletProxy.IncreaseRelayNonce(Owner.TransactOpts())
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeTrue())

			tx, err = WalletProxy.ExecutePrivilegedRelayedTransaction(Controller.TransactOpts(ethertest.WithGasLimit(500000)), nonce, batch, signature)
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeFalse())

			returnData, _ := ethCall(tx)
			Expect(string(returnData[len(returnData)-64:])).To(ContainSubstring("tx replay"))
		})

		It("should emit an IncreasedRelayNonce event", func() {
			it, err := WalletProxy.FilterIncreasedRelayNonce(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(it.Next()).To(BeTrue())
			evt := it.Event
			Expect(evt.Sender).To(Equal(Owner.Address()))
			Expect(evt.CurrentNonce.String()).To(Equal("1"))
			Expect(it.Next()).To(BeTrue())
			evt = it.Event
			Expect(it.Next()).To(BeFalse())
			Expect(evt.Sender).To(Equal(Owner.Address()))
			Expect(evt.CurrentNonce.String()).To(Equal("2"))
		})
	})

})
