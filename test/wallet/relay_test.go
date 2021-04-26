package wallet_test

import (
	"context"
	"crypto/ecdsa"
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

var _ = Describe("relay Tx", func() {

	BeforeEach(func() {
		BankAccount.MustTransfer(Backend, WalletProxyAddress, EthToWei(2))
	})

	When("relayNonce variable initialized to 0", func() {
		It("should succeed", func() {
			nonce, err := WalletProxy.RelayNonce(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(nonce.String()).To(Equal("0"))
		})
	})

	When("a non-owner account tries to increase the nonce", func() {
		It("should fail", func() {
			tx, err := WalletProxy.IncreaseRelayNonce(Controller.TransactOpts(ethertest.WithGasLimit(60000)))
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeFalse())
		})
	})

	When("a random account tries to relay", func() {
		It("should fail", func() {
			a, err := abi.JSON(strings.NewReader(WALLET_ABI))
			Expect(err).ToNot(HaveOccurred())
			privateKey, _ := crypto.GenerateKey()
			randomAddress := crypto.PubkeyToAddress(privateKey.PublicKey)
			data, err := a.Pack("transfer", randomAddress, common.HexToAddress("0x0"), EthToWei(1))
			Expect(err).ToNot(HaveOccurred())

			nonce := big.NewInt(0)
			chainId := big.NewInt(1337)
			signature, err := SignData(chainId, WalletProxyAddress, nonce, data, Owner.PrivKey())
			Expect(err).ToNot(HaveOccurred())

			tx, err := WalletProxy.ExecuteRelayedTransaction(RandomAccount.TransactOpts(ethertest.WithGasLimit(500000)), nonce, data, signature)
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeFalse())
		})
	})

	When("a controller tries to relay a transaction signed from the same owner but for another wallet", func() {
		It("should fail", func() {
			a, err := abi.JSON(strings.NewReader(WALLET_ABI))
			Expect(err).ToNot(HaveOccurred())
			privateKey, _ := crypto.GenerateKey()
			randomAddress := crypto.PubkeyToAddress(privateKey.PublicKey)
			data, err := a.Pack("transfer", randomAddress, common.HexToAddress("0x0"), EthToWei(1))
			Expect(err).ToNot(HaveOccurred())

			nonce := big.NewInt(0)
			chainId := big.NewInt(1) //the actual ethereum mainnet
			signature, err := SignData(chainId, WalletProxyAddress, nonce, data, Owner.PrivKey())
			Expect(err).ToNot(HaveOccurred())

			tx, err := WalletProxy.ExecuteRelayedTransaction(Controller.TransactOpts(ethertest.WithGasLimit(500000)), nonce, data, signature)
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeFalse())
			returnData, _ := ethCall(tx)
			Expect(string(returnData[len(returnData)-64:])).To(ContainSubstring("invalid signature"))
		})
	})

	When("a controller tries to relay a transaction signed from the same owner but for another wallet", func() {
		It("should fail", func() {
			a, err := abi.JSON(strings.NewReader(WALLET_ABI))
			Expect(err).ToNot(HaveOccurred())
			privateKey, _ := crypto.GenerateKey()
			randomAddress := crypto.PubkeyToAddress(privateKey.PublicKey)
			data, err := a.Pack("transfer", randomAddress, common.HexToAddress("0x0"), EthToWei(1))
			Expect(err).ToNot(HaveOccurred())

			nonce := big.NewInt(0)
			chainId := big.NewInt(1337)
			signature, err := SignData(chainId, randomAddress, nonce, data, Owner.PrivKey())
			Expect(err).ToNot(HaveOccurred())

			tx, err := WalletProxy.ExecuteRelayedTransaction(Controller.TransactOpts(ethertest.WithGasLimit(500000)), nonce, data, signature)
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeFalse())
			returnData, _ := ethCall(tx)
			Expect(string(returnData[len(returnData)-64:])).To(ContainSubstring("invalid signature"))
		})
	})

	When("a controller tries to relay a transaction signed by a random account", func() {
		It("should fail", func() {
			a, err := abi.JSON(strings.NewReader(WALLET_ABI))
			Expect(err).ToNot(HaveOccurred())
			privateKey, _ := crypto.GenerateKey()
			randomAddress := crypto.PubkeyToAddress(privateKey.PublicKey)
			data, err := a.Pack("transfer", randomAddress, common.HexToAddress("0x0"), EthToWei(1))
			Expect(err).ToNot(HaveOccurred())

			nonce := big.NewInt(0)
			chainId := big.NewInt(1337)
			signature, err := SignData(chainId, WalletProxyAddress, nonce, data, privateKey)
			Expect(err).ToNot(HaveOccurred())

			tx, err := WalletProxy.ExecuteRelayedTransaction(Controller.TransactOpts(ethertest.WithGasLimit(500000)), nonce, data, signature)
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeFalse())
			returnData, _ := ethCall(tx)
			Expect(string(returnData[len(returnData)-64:])).To(ContainSubstring("invalid signature"))
		})
	})

	When("a controller tries to relay an owner signed transaction", func() {
		var randomAddress common.Address
		var privateKey *ecdsa.PrivateKey

		//transfer ownership to a newly generated daddress
		BeforeEach(func() {
			privateKey, _ = crypto.GenerateKey()
			randomAddress = crypto.PubkeyToAddress(privateKey.PublicKey)
		})

		When("the call succeeds", func() {
			BeforeEach(func() {

				tx, err := WalletProxy.TransferOwnership(Owner.TransactOpts(), randomAddress, false)
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())

				a, err := abi.JSON(strings.NewReader(WALLET_ABI))
				Expect(err).ToNot(HaveOccurred())
				data, err := a.Pack("transfer", randomAddress, common.HexToAddress("0x0"), EthToWei(1))
				Expect(err).ToNot(HaveOccurred())

				batch := []byte(fmt.Sprintf("%s%s%s%s", WalletProxyAddress, abi.U256(EthToWei(0)), abi.U256(big.NewInt(int64(len(data)))), data))

				nonce := big.NewInt(0)
				chainId := big.NewInt(1337)
				signature, err := SignData(chainId, WalletProxyAddress, nonce, batch, privateKey)
				Expect(err).ToNot(HaveOccurred())

				tx, err = WalletProxy.ExecuteRelayedTransaction(Controller.TransactOpts(ethertest.WithGasLimit(500000)), nonce, batch, signature)
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})

			It("should change the owner", func() {
				o, err := WalletProxy.Owner(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(o).To(Equal(randomAddress))
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
				a, err := abi.JSON(strings.NewReader(WALLET_ABI))
				Expect(err).ToNot(HaveOccurred())
				data, err := a.Pack("transfer", randomAddress, common.HexToAddress("0x0"), EthToWei(1))
				batch := []byte(fmt.Sprintf("%s%s%s%s", WalletProxyAddress, abi.U256(EthToWei(0)), abi.U256(big.NewInt(int64(len(data)))), data))
				Expect(evt.Data).To(Equal(batch))
				Expect(evt.ReturnData).To(Equal([]byte{}))
			})

			It("should decrease the wallet's ETH balance ", func() {
				b, err := Backend.BalanceAt(context.Background(), WalletProxyAddress, nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(b.String()).To(Equal(EthToWei(1).String()))
			})

			It("should increase the random address' balance ", func() {
				b, err := Backend.BalanceAt(context.Background(), randomAddress, nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(b.String()).To(Equal(EthToWei(1).String()))
			})

			It("should fail when trying to replay", func() {
				a, err := abi.JSON(strings.NewReader(WALLET_ABI))
				Expect(err).ToNot(HaveOccurred())
				data, err := a.Pack("transfer", randomAddress, common.HexToAddress("0x0"), EthToWei(1))
				Expect(err).ToNot(HaveOccurred())

				nonce := big.NewInt(0)
				chainId := big.NewInt(1337)
				signature, err := SignData(chainId, WalletProxyAddress, nonce, data, privateKey)
				Expect(err).ToNot(HaveOccurred())

				tx, err := WalletProxy.ExecuteRelayedTransaction(Controller.TransactOpts(ethertest.WithGasLimit(500000)), nonce, data, signature)
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeFalse())
			})

			It("should succeed when increasing the nonce", func() {
				a, err := abi.JSON(strings.NewReader(WALLET_ABI))
				Expect(err).ToNot(HaveOccurred())
				data, err := a.Pack("transfer", randomAddress, common.HexToAddress("0x0"), EthToWei(1))
				Expect(err).ToNot(HaveOccurred())

				batch := []byte(fmt.Sprintf("%s%s%s%s", WalletProxyAddress, abi.U256(EthToWei(0)), abi.U256(big.NewInt(int64(len(data)))), data))

				nonce := big.NewInt(1)
				chainId := big.NewInt(1337)
				signature, err := SignData(chainId, WalletProxyAddress, nonce, batch, privateKey)
				Expect(err).ToNot(HaveOccurred())

				tx, err := WalletProxy.ExecuteRelayedTransaction(Controller.TransactOpts(), nonce, batch, signature)
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})
		})

		When("the call fails (transfer to 0x0 address)", func() {
			It("should return the error string emitted by require", func() {
				a, err := abi.JSON(strings.NewReader(WALLET_ABI))
				Expect(err).ToNot(HaveOccurred())
				data, err := a.Pack("transfer", common.HexToAddress("0x0"), common.HexToAddress("0x0"), EthToWei(1))
				Expect(err).ToNot(HaveOccurred())

				batch := []byte(fmt.Sprintf("%s%s%s%s", WalletProxyAddress, abi.U256(EthToWei(0)), abi.U256(big.NewInt(int64(len(data)))), data))

				nonce := big.NewInt(0)
				chainId := big.NewInt(1337)
				signature, err := SignData(chainId, WalletProxyAddress, nonce, batch, privateKey)
				Expect(err).ToNot(HaveOccurred())

				tx, err := WalletProxy.TransferOwnership(Owner.TransactOpts(), randomAddress, false)
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())

				tx, err = WalletProxy.ExecuteRelayedTransaction(Controller.TransactOpts(ethertest.WithGasLimit(500000)), nonce, batch, signature)
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
				signature, err := SignData(chainId, WalletProxyAddress, nonce, batch, privateKey)
				Expect(err).ToNot(HaveOccurred())

				tx, err := WalletProxy.IncreaseRelayNonce(Owner.TransactOpts())
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())

				tx, err = WalletProxy.IncreaseRelayNonce(Owner.TransactOpts())
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())

				tx, err = WalletProxy.TransferOwnership(Owner.TransactOpts(), randomAddress, false)
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
})
