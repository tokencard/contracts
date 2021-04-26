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

var _ = Describe("executeRelayedTransaction", func() {

	BeforeEach(func() {
		BankAccount.MustTransfer(Backend, WalletProxyAddress, EthToWei(4))
	})

	When("a non-owner account tries to increase the nonce", func() {
		It("should fail", func() {
			tx, err := WalletProxy.IncreaseRelayNonce(Controller.TransactOpts(ethertest.WithGasLimit(60000)))
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeFalse())
			returnData, _ := ethCall(tx)
			Expect(string(returnData[len(returnData)-64:])).To(ContainSubstring("sender is not owner"))
		})
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

			tx, err := WalletProxy.ExecuteRelayedTransaction(RandomAccount.TransactOpts(ethertest.WithGasLimit(500000)), nonce, batch, signature)
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeFalse())
			returnData, _ := ethCall(tx)
			Expect(string(returnData[len(returnData)-64:])).To(ContainSubstring("sender is not controller"))
		})
	})

	When("a controller tries to relay a transaction signed by a random account", func() {
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

			tx, err := WalletProxy.ExecuteRelayedTransaction(Controller.TransactOpts(ethertest.WithGasLimit(1000000)), nonce, batch, signature)
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeFalse())
			returnData, _ := ethCall(tx)
			Expect(string(returnData[len(returnData)-64:])).To(ContainSubstring("invalid signature"))
		})
	})

	When("a controller tries to relay two owner-signed transactions: send value(no data) + transfer()", func() {

		When("the call succeeds", func() {
			BeforeEach(func() {

				// send value (1 ETH)
				data1 := fmt.Sprintf("%s%s%s", RandomAccount.Address(), abi.U256(EthToWei(1)), abi.U256(big.NewInt(0)))
				// use wallet's transfer
				a, err := abi.JSON(strings.NewReader(WALLET_ABI))
				Expect(err).ToNot(HaveOccurred())
				data2, err := a.Pack("transfer", RandomAccount.Address(), common.HexToAddress("0x0"), EthToWei(1))
				Expect(err).ToNot(HaveOccurred())

				batch := []byte(fmt.Sprintf("%s%s%s%s%s", data1, WalletProxyAddress, abi.U256(EthToWei(0)), abi.U256(big.NewInt(int64(len(data2)))), data2))

				nonce := big.NewInt(0)
				chainId := big.NewInt(1337)
				signature, err := SignData(chainId, WalletProxyAddress, nonce, batch, Owner.PrivKey())
				Expect(err).ToNot(HaveOccurred())

				tx, err := WalletProxy.ExecuteRelayedTransaction(Controller.TransactOpts(ethertest.WithGasLimit(1000000)), nonce, batch, signature)
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})

			It("should emit an IncreasedRelayNonce event", func() {
				it, err := WalletProxy.FilterIncreasedRelayNonce(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(it.Next()).To(BeTrue())
				evt := it.Event
				Expect(it.Next()).To(BeFalse())
				Expect(evt.Sender).To(Equal(Controller.Address()))
				Expect(evt.CurrentNonce.String()).To(Equal("1"))
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
				data2, err := a.Pack("transfer", RandomAccount.Address(), common.HexToAddress("0x0"), EthToWei(1))
				batchData := []byte(fmt.Sprintf("%s%s%s%s%s", data1, WalletProxyAddress, abi.U256(EthToWei(0)), abi.U256(big.NewInt(int64(len(data2)))), data2))
				Expect(evt.Data).To(Equal(batchData))
				Expect(evt.Privileged).To(Equal(false))
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
				data, err := a.Pack("transfer", RandomAccount.Address(), common.HexToAddress("0x0"), EthToWei(1))
				Expect(evt.Destination).To(Equal(WalletProxyAddress))
				Expect(evt.Value.String()).To(Equal("0"))
				Expect(evt.Data).To(Equal(data))
				Expect(evt.ReturnData).To(Equal(common.Hex2Bytes("")))
			})

			It("should decrease the wallet's ETH balance ", func() {
				b, err := Backend.BalanceAt(context.Background(), WalletProxyAddress, nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(b.String()).To(Equal(EthToWei(2).String()))
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

				tx, err := WalletProxy.ExecuteRelayedTransaction(Controller.TransactOpts(ethertest.WithGasLimit(500000)), nonce, batch, signature)
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

				tx, err := WalletProxy.ExecuteRelayedTransaction(Controller.TransactOpts(), nonce, batch, signature)
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})
		})
	})

	When("transfering to 0 address", func() {
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

			tx, err := WalletProxy.ExecuteRelayedTransaction(Controller.TransactOpts(ethertest.WithGasLimit(500000)), nonce, batch, signature)
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

			tx, err = WalletProxy.ExecuteRelayedTransaction(Controller.TransactOpts(ethertest.WithGasLimit(500000)), nonce, batch, signature)
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
