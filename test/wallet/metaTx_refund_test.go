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
)

var _ = Describe("metaTx refund", func() {

	Context("when the wallet has enough ETH and ERC20 tokens", func() {
		BeforeEach(func() {
			BankAccount.MustTransfer(Backend, WalletProxyAddress, EthToWei(2))

			var err error
			tx, err := TKNBurner.Mint(BankAccount.TransactOpts(), WalletProxyAddress, big.NewInt(300))
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeTrue())
		})

		When("ownership is transfered to a newly generated address", func() {

			var randomAddress common.Address
			var tokenBank common.Address
			var privateKey *ecdsa.PrivateKey

			//transfer ownership to a newly generated address (makes signing easier)
			BeforeEach(func() {
				privateKey, _ = crypto.GenerateKey()
				tokenBank = crypto.PubkeyToAddress(privateKey.PublicKey)
				privateKey, _ = crypto.GenerateKey()
				randomAddress = crypto.PubkeyToAddress(privateKey.PublicKey)
				tx, err := WalletProxy.TransferOwnership(Owner.TransactOpts(), randomAddress, false)
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())

			})

			When("controller relays an ETH transfer and a refund", func() {

				BeforeEach(func() {
					batch := fmt.Sprintf("%s%s%s%s%s%s", randomAddress, abi.U256(EthToWei(1)), abi.U256(big.NewInt(0)),
						tokenBank, abi.U256(big.NewInt(1000)), abi.U256(big.NewInt(0)))

					a, err := abi.JSON(strings.NewReader(WALLET_ABI))
					Expect(err).ToNot(HaveOccurred())
					data, err := a.Pack("batchExecuteTransaction", []byte(batch))
					Expect(err).ToNot(HaveOccurred())

					nonce := big.NewInt(0)
					chainId := big.NewInt(1337)
					signature, err := SignData(chainId, WalletProxyAddress, nonce, data, privateKey)
					Expect(err).ToNot(HaveOccurred())

					tx, err := WalletProxy.ExecuteRelayedTransaction(Controller.TransactOpts(), nonce, data, signature)
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeTrue())
				})

				It("should increase random address' balance by the same amount", func() {
					b, e := Backend.BalanceAt(context.Background(), randomAddress, nil)
					Expect(e).ToNot(HaveOccurred())
					Expect(b.String()).To(Equal(EthToWei(1).String()))
				})

				It("should increase the Controller's balance by 1000 wei", func() {
					b, e := Backend.BalanceAt(context.Background(), tokenBank, nil)
					Expect(e).ToNot(HaveOccurred())
					Expect(b.String()).To(Equal("1000"))
				})
			})

			When("controller relays an ERC-20 transfer and an ERC-20 refund", func() {
				BeforeEach(func() {

					a, err := abi.JSON(strings.NewReader(ERC20ABI))
					Expect(err).ToNot(HaveOccurred())
					data, err := a.Pack("transfer", randomAddress, big.NewInt(290))
					Expect(err).ToNot(HaveOccurred())
					batch := fmt.Sprintf("%s%s%s%s", TKNBurnerAddress, abi.U256(big.NewInt(0)), abi.U256(big.NewInt(int64(len(data)))), data)

					data, err = a.Pack("transfer", tokenBank, big.NewInt(10))
					Expect(err).ToNot(HaveOccurred())
					batch = fmt.Sprintf("%s%s%s%s%s", batch, TKNBurnerAddress, abi.U256(big.NewInt(0)), abi.U256(big.NewInt(int64(len(data)))), data)

					a, err = abi.JSON(strings.NewReader(WALLET_ABI))
					Expect(err).ToNot(HaveOccurred())
					data, err = a.Pack("batchExecuteTransaction", []byte(batch))
					Expect(err).ToNot(HaveOccurred())

					nonce := big.NewInt(0)
					chainId := big.NewInt(1337)
					signature, err := SignData(chainId, WalletProxyAddress, nonce, data, privateKey)
					Expect(err).ToNot(HaveOccurred())

					tx, err := WalletProxy.ExecuteRelayedTransaction(Controller.TransactOpts(), nonce, data, signature)
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeTrue())
				})

				It("should increase the TKN balance of the random address", func() {
					b, err := TKNBurner.BalanceOf(nil, randomAddress)
					Expect(err).ToNot(HaveOccurred())
					Expect(b.String()).To(Equal("290"))
				})

				It("should increase the TKN balance of the tokenBank account", func() {
					b, err := TKNBurner.BalanceOf(nil, tokenBank)
					Expect(err).ToNot(HaveOccurred())
					Expect(b.String()).To(Equal("10"))
				})

				It("should decrease the TKN balance of the wallet", func() {
					b, err := TKNBurner.BalanceOf(nil, WalletProxyAddress)
					Expect(err).ToNot(HaveOccurred())
					Expect(b.String()).To(Equal("0"))
				})

				It("should emit 2 ExecutedTransaction events", func() {
					it, err := WalletProxy.FilterExecutedTransaction(nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(it.Next()).To(BeTrue())
					evt := it.Event
					Expect(it.Next()).To(BeTrue())
					Expect(evt.Destination).To(Equal(TKNBurnerAddress))
					Expect(evt.Value.String()).To(Equal("0"))
					a, _ := abi.JSON(strings.NewReader(ERC20ABI))
					d, _ := a.Pack("transfer", randomAddress, big.NewInt(290))
					Expect(evt.Data).To(Equal(d))
					Expect(evt.ReturnData).To(Equal(common.Hex2Bytes("0000000000000000000000000000000000000000000000000000000000000001")))
					evt = it.Event
					Expect(it.Next()).To(BeFalse())
					Expect(evt.Destination).To(Equal(TKNBurnerAddress))
					Expect(evt.Value.String()).To(Equal("0"))
					a, _ = abi.JSON(strings.NewReader(ERC20ABI))
					d, _ = a.Pack("transfer", tokenBank, big.NewInt(10))
					Expect(evt.Data).To(Equal(d))
					Expect(evt.ReturnData).To(Equal(common.Hex2Bytes("0000000000000000000000000000000000000000000000000000000000000001")))
				})

				It("should emit an ExecutedRelayedTransaction event", func() {
					it, err := WalletProxy.FilterExecutedRelayedTransaction(nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(it.Next()).To(BeTrue())
					evt := it.Event
					Expect(it.Next()).To(BeFalse())
					a, err := abi.JSON(strings.NewReader(ERC20ABI))
					Expect(err).ToNot(HaveOccurred())
					data, _ := a.Pack("transfer", randomAddress, big.NewInt(290))
					batch := fmt.Sprintf("%s%s%s%s", TKNBurnerAddress, abi.U256(big.NewInt(0)), abi.U256(big.NewInt(int64(len(data)))), data)
					data, _ = a.Pack("transfer", tokenBank, big.NewInt(10))
					batch = fmt.Sprintf("%s%s%s%s%s", batch, TKNBurnerAddress, abi.U256(big.NewInt(0)), abi.U256(big.NewInt(int64(len(data)))), data)
					a, err = abi.JSON(strings.NewReader(WALLET_ABI))
					data, _ = a.Pack("batchExecuteTransaction", []byte(batch))
					Expect(evt.Data).To(Equal(data))
					Expect(evt.ReturnData).To(Equal(common.Hex2Bytes("")))
				})
			})

		})

	})
})
