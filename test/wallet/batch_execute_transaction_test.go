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

var _ = Describe("batchExecuteTransaction", func() {

	Context("when the wallet has enough ETH and ERC20 tokens", func() {
		BeforeEach(func() {
			BankAccount.MustTransfer(Backend, WalletProxyAddress, EthToWei(1))

			var err error
			tx, err := TKNBurner.Mint(BankAccount.TransactOpts(), WalletProxyAddress, big.NewInt(300))
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeTrue())
		})

		When("I batch a normal and an ERC20 transfer (both external calls)", func() {

			var randomAddress common.Address

			BeforeEach(func() {

				privateKey, _ := crypto.GenerateKey()
				randomAddress = crypto.PubkeyToAddress(privateKey.PublicKey)

				batch := fmt.Sprintf("%s%s%s", randomAddress, abi.U256(EthToWei(1)), abi.U256(big.NewInt(0)))

				a, err := abi.JSON(strings.NewReader(ERC20ABI))
				Expect(err).ToNot(HaveOccurred())
				data, err := a.Pack("transfer", randomAddress, big.NewInt(300))
				Expect(err).ToNot(HaveOccurred())
				batch = fmt.Sprintf("%s%s%s%s%s", batch, TKNBurnerAddress, abi.U256(big.NewInt(0)), abi.U256(big.NewInt(int64(len(data)))), data)

				tx, err := WalletProxy.BatchExecuteTransaction(Owner.TransactOpts(), []byte(batch))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})

			It("should increase random address' balance by the same amount", func() {
				b, e := Backend.BalanceAt(context.Background(), randomAddress, nil)
				Expect(e).ToNot(HaveOccurred())
				Expect(b.String()).To(Equal(EthToWei(1).String()))
			})

			It("should increase the TKN balance of the random address", func() {
				b, err := TKNBurner.BalanceOf(nil, randomAddress)
				Expect(err).ToNot(HaveOccurred())
				Expect(b.String()).To(Equal("300"))
			})

			It("should decrease the wallet's balance by the same amount", func() {
				b, e := Backend.BalanceAt(context.Background(), WalletProxyAddress, nil)
				Expect(e).ToNot(HaveOccurred())
				Expect(b.String()).To(Equal("0"))
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
				Expect(evt.Destination).To(Equal(randomAddress))
				Expect(evt.Value.String()).To(Equal(EthToWei(1).String()))
				Expect(evt.Data).To(Equal([]uint8{}))
				Expect(evt.ReturnData).To(Equal(common.Hex2Bytes("")))
				evt = it.Event
				Expect(it.Next()).To(BeFalse())
				Expect(evt.Destination).To(Equal(TKNBurnerAddress))
				Expect(evt.Value.String()).To(Equal("0"))
				a, _ := abi.JSON(strings.NewReader(ERC20ABI))
				d, _ := a.Pack("transfer", randomAddress, big.NewInt(300))
				Expect(evt.Data).To(Equal(d))
				Expect(evt.ReturnData).To(Equal(common.Hex2Bytes("0000000000000000000000000000000000000000000000000000000000000001")))
			})
		})

		When("I batch 3 transactions and the 3rd one fails (transfer 0 value)", func() {

			var randomAddress common.Address

			BeforeEach(func() {

				privateKey, _ := crypto.GenerateKey()
				randomAddress = crypto.PubkeyToAddress(privateKey.PublicKey)

				batch := fmt.Sprintf("%s%s%s", randomAddress, abi.U256(EthToWei(1)), abi.U256(big.NewInt(0)))

				a, err := abi.JSON(strings.NewReader(ERC20ABI))
				Expect(err).ToNot(HaveOccurred())
				data, err := a.Pack("transfer", randomAddress, big.NewInt(300))
				Expect(err).ToNot(HaveOccurred())
				batch = fmt.Sprintf("%s%s%s%s%s", batch, TKNBurnerAddress, abi.U256(big.NewInt(0)), abi.U256(big.NewInt(int64(len(data)))), data)

				a, err = abi.JSON(strings.NewReader(WALLET_ABI))
				Expect(err).ToNot(HaveOccurred())
				data, err = a.Pack("transfer", common.HexToAddress("0x0"), common.HexToAddress("0x0"), big.NewInt(0))
				Expect(err).ToNot(HaveOccurred())
				batch = fmt.Sprintf("%s%s%s%s%s", batch, WalletProxyAddress, abi.U256(big.NewInt(0)), abi.U256(big.NewInt(int64(len(data)))), data)

				tx, err := WalletProxy.BatchExecuteTransaction(Owner.TransactOpts(ethertest.WithGasLimit(1000000)), []byte(batch))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeFalse())

				returnData, _ := ethCall(tx)
				Expect(string(returnData[len(returnData)-64:])).To(ContainSubstring("value=0"))
			})

			It("should NOT increase random address' balance", func() {
				b, e := Backend.BalanceAt(context.Background(), randomAddress, nil)
				Expect(e).ToNot(HaveOccurred())
				Expect(b.String()).To(Equal("0"))
			})

			It("should NOT increase the TKN balance of the random address", func() {
				b, err := TKNBurner.BalanceOf(nil, randomAddress)
				Expect(err).ToNot(HaveOccurred())
				Expect(b.String()).To(Equal("0"))
			})

			It("should NOT decrease the wallet's balance", func() {
				b, e := Backend.BalanceAt(context.Background(), WalletProxyAddress, nil)
				Expect(e).ToNot(HaveOccurred())
				Expect(b.String()).To(Equal(EthToWei(1).String()))
			})

			It("should NOT decrease the TKN balance of the wallet", func() {
				b, err := TKNBurner.BalanceOf(nil, WalletProxyAddress)
				Expect(err).ToNot(HaveOccurred())
				Expect(b.String()).To(Equal("300"))
			})

		})

		When("I set the daily limit and the whitelist", func() {

			BeforeEach(func() {

				a, err := abi.JSON(strings.NewReader(WALLET_ABI))
				Expect(err).ToNot(HaveOccurred())
				data, err := a.Pack("setDailyLimit", MweiToWei(100))
				Expect(err).ToNot(HaveOccurred())
				batch := fmt.Sprintf("%s%s%s%s", WalletProxyAddress, abi.U256(big.NewInt(0)), abi.U256(big.NewInt(int64(len(data)))), data)

				data, err = a.Pack("addToWhitelist", []common.Address{RandomAccount.Address()})
				Expect(err).ToNot(HaveOccurred())
				batch = fmt.Sprintf("%s%s%s%s%s", batch, WalletProxyAddress, abi.U256(big.NewInt(0)), abi.U256(big.NewInt(int64(len(data)))), data)

				tx, err := WalletProxy.BatchExecuteTransaction(Owner.TransactOpts(), []byte(batch))
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
				Expect(evt.Sender).To(Equal(WalletProxyAddress))
				Expect(evt.Addresses).To(Equal([]common.Address{RandomAccount.Address()}))
			})

			It("should emit a limit set event", func() {
				it, err := Wallet.FilterSetDailyLimit(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(it.Next()).To(BeTrue())
				evt := it.Event
				Expect(it.Next()).To(BeFalse())
				Expect(evt.Sender).To(Equal(WalletProxyAddress))
				Expect(evt.Amount).To(Equal(EthToWei(1)))
			})

			It("should lower the available amount to 1 $USD", func() {
				av, err := Wallet.DailyLimitAvailable(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(av.String()).To(Equal(EthToWei(1).String()))
			})

			It("should have a limit of 1 $USD", func() {
				sl, err := Wallet.DailyLimitValue(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(sl.String()).To(Equal(EthToWei(1).String()))
			})
		})

		When("the datalength is a maliciously encoded", func() {

			var randomAddress common.Address

			It("should revert if the encoded length is larger than the actual batch size", func() {

				privateKey, _ := crypto.GenerateKey()
				randomAddress = crypto.PubkeyToAddress(privateKey.PublicKey)

				batch := fmt.Sprintf("%s%s%s", randomAddress, abi.U256(EthToWei(1)), abi.U256(big.NewInt(99999)))

				tx, err := WalletProxy.BatchExecuteTransaction(Owner.TransactOpts(ethertest.WithGasLimit(1000000)), []byte(batch))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeFalse())

				returnData, _ := ethCall(tx)
				Expect(string(returnData[len(returnData)-64:])).To(ContainSubstring("out of bounds"))
			})

			It("should revert if there is an index overflow", func() {

				privateKey, _ := crypto.GenerateKey()
				randomAddress = crypto.PubkeyToAddress(privateKey.PublicKey)

				batch := fmt.Sprintf("%s%s%s", randomAddress, abi.U256(EthToWei(1)), abi.U256(big.NewInt(-1)))

				tx, err := WalletProxy.BatchExecuteTransaction(Owner.TransactOpts(ethertest.WithGasLimit(1000000)), []byte(batch))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeFalse())

				returnData, _ := ethCall(tx)
				Expect(string(returnData[len(returnData)-64:])).To(ContainSubstring("SafeMath: addition overflow"))
			})
		})

		When("there are not enough bytes to encode the executeTransaction inputs", func() {

			var randomAddress common.Address

			It("should revert", func() {

				privateKey, _ := crypto.GenerateKey()
				randomAddress = crypto.PubkeyToAddress(privateKey.PublicKey)

				batch := fmt.Sprintf("%s%s", randomAddress, abi.U256(EthToWei(1)))

				tx, err := WalletProxy.BatchExecuteTransaction(Owner.TransactOpts(ethertest.WithGasLimit(1000000)), []byte(batch))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeFalse())

				returnData, _ := ethCall(tx)
				Expect(string(returnData[len(returnData)-64:])).To(ContainSubstring("SafeMath: subtraction overflow"))
			})
		})

// 	})
// })

const WALLET_ABI = `[
    {
        "constant": false,
        "inputs": [
            {
                "name": "_to",
                "type": "address"
            },
            {
                "name": "_asset",
                "type": "address"
            },
            {
                "name": "_amount",
                "type": "uint256"
            }
        ],
        "name": "transfer",
        "outputs": [],
        "payable": false,
        "stateMutability": "nonpayable",
        "type": "function"
    },
    {
        "constant": false,
        "inputs": [
            {
                "name": "_addresses",
                "type": "address[]"
            }
        ],
        "name": "addToWhitelist",
        "outputs": [],
        "payable": false,
        "stateMutability": "nonpayable",
        "type": "function"
    },
    {
        "constant": false,
        "inputs": [
            {
                "name": "_amount",
                "type": "uint256"
            }
        ],
        "name": "setDailyLimit",
        "outputs": [],
        "payable": false,
        "stateMutability": "nonpayable",
        "type": "function"
    }
]`
