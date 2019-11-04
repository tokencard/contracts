package wallet_test

import (
	"context"
	"math/big"
	"strings"
    "fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tokencard/contracts/v2/test/shared"
)

var _ = Describe("executeTransactions", func() {

	Context("when the wallet has enough ETH and ERC20 tokens", func() {
		BeforeEach(func() {
			BankAccount.MustTransfer(Backend, WalletAddress, EthToWei(1))

			var err error
			tx, err := TKNBurner.Mint(BankAccount.TransactOpts(), WalletAddress, big.NewInt(300))
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

				tx, err := Wallet.ExecuteTransactions(Owner.TransactOpts(), []byte(batch))
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
				b, e := Backend.BalanceAt(context.Background(), WalletAddress, nil)
				Expect(e).ToNot(HaveOccurred())
				Expect(b.String()).To(Equal("0"))
			})

            It("should decrease the TKN balance of the wallet", func() {
				b, err := TKNBurner.BalanceOf(nil, WalletAddress)
				Expect(err).ToNot(HaveOccurred())
				Expect(b.String()).To(Equal("0"))
			})

			It("should emit 2 ExecutedTransaction events", func() {
				it, err := Wallet.FilterExecutedTransaction(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(it.Next()).To(BeTrue())
				evt := it.Event
				Expect(it.Next()).To(BeTrue())
				Expect(evt.Destination).To(Equal(randomAddress))
				Expect(evt.Value.String()).To(Equal(EthToWei(1).String()))
				Expect(evt.Data).To(Equal([]uint8{}))
                Expect(evt.Returndata).To(Equal(common.Hex2Bytes("")))
                evt = it.Event
				Expect(it.Next()).To(BeFalse())
				Expect(evt.Destination).To(Equal(TKNBurnerAddress))
				Expect(evt.Value.String()).To(Equal("0"))
                a, _ := abi.JSON(strings.NewReader(ERC20ABI))
                d, _ := a.Pack("transfer", randomAddress, big.NewInt(300))
				Expect(evt.Data).To(Equal(d))
                Expect(evt.Returndata).To(Equal(common.Hex2Bytes("0000000000000000000000000000000000000000000000000000000000000001")))
			})
        })

        When("I top up gas and set the whitelist", func() {

			BeforeEach(func() {

                a, err := abi.JSON(strings.NewReader(WALLET_ABI))
				Expect(err).ToNot(HaveOccurred())
				data, err := a.Pack("setSpendLimit", EthToWei(1))
				Expect(err).ToNot(HaveOccurred())
                batch := fmt.Sprintf("%s%s%s%s", WalletAddress, abi.U256(big.NewInt(0)), abi.U256(big.NewInt(int64(len(data)))), data)

				data, err = a.Pack("setWhitelist", []common.Address{RandomAccount.Address()})
				Expect(err).ToNot(HaveOccurred())
                batch = fmt.Sprintf("%s%s%s%s%s", batch, WalletAddress, abi.U256(big.NewInt(0)), abi.U256(big.NewInt(int64(len(data)))), data)

				tx, err := Wallet.ExecuteTransactions(Owner.TransactOpts(), []byte(batch))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})

    		It("should update the initializedWhitelist flag", func() {
    			initialized, err := Wallet.IsSetWhitelist(nil)
    			Expect(err).ToNot(HaveOccurred())
    			Expect(initialized).To(BeTrue())
    		})

    		It("should add the random account to the whitelist", func() {
    			isWhitelisted, err := Wallet.WhitelistMap(nil, RandomAccount.Address())
    			Expect(err).ToNot(HaveOccurred())
    			Expect(isWhitelisted).To(BeTrue())
    		})

    		It("should emit WhitelistAddition event", func() {
    			it, err := Wallet.FilterAddedToWhitelist(nil)
    			Expect(err).ToNot(HaveOccurred())
    			Expect(it.Next()).To(BeTrue())
    			evt := it.Event
    			Expect(it.Next()).To(BeFalse())
    			Expect(evt.Sender).To(Equal(WalletAddress))
    			Expect(evt.Addresses).To(Equal([]common.Address{RandomAccount.Address()}))
    		})

            It("should emit a spend limit set event", func() {
				it, err := Wallet.FilterSetSpendLimit(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(it.Next()).To(BeTrue())
				evt := it.Event
				Expect(it.Next()).To(BeFalse())
				Expect(evt.Sender).To(Equal(WalletAddress))
				Expect(evt.Amount).To(Equal(EthToWei(1)))
			})

			It("should lower the spend available to 1 ETH", func() {
				av, err := Wallet.SpendLimitAvailable(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(av.String()).To(Equal(EthToWei(1).String()))
			})

			It("should have spend limit of 1 ETH", func() {
				sl, err := Wallet.SpendLimitValue(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(sl.String()).To(Equal(EthToWei(1).String()))
			})
        })
    })
})

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
                "name": "_amount",
                "type": "uint256"
            }
        ],
        "name": "topUpGas",
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
        "name": "setWhitelist",
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
        "name": "setSpendLimit",
        "outputs": [],
        "payable": false,
        "stateMutability": "nonpayable",
        "type": "function"
    },
    {
        "constant": false,
        "inputs": [
            {
                "name": "_transactionBatch",
                "type": "bytes"
            }
        ],
        "name": "executeTransactions",
        "outputs": [
            {
                "name": "",
                "type": "bool"
            }
        ],
        "payable": false,
        "stateMutability": "nonpayable",
        "type": "function"
    }
]`
