package wallet_test

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tokencard/contracts/test/shared"
	"github.com/tokencard/ethertest"
)

var _ = FDescribe("executeTransaction", func() {

	Context("when the wallet has enough ETH", func() {
		BeforeEach(func() {
			BankAccount.MustTransfer(Backend, WalletAddress, EthToWei(1))
			BankAccount.MustTransfer(Backend, Controller.Address(), EthToWei(1))
		})

		var tx *types.Transaction

		When("I transfer 500 Finney tp a random person using 'executeTransaction'", func() {
			BeforeEach(func() {
				tx, err := Wallet.ExecuteTransaction(Owner.TransactOpts(ethertest.WithGasLimit(100000)), RandomAccount.Address(), FinneyToWei(500), nil)
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})

			It("should transfer 500 Finney to random person", func() {
				bal := RandomAccount.Balance(Backend)
				Expect(bal.String()).To(Equal("1000500000000000000000"))
			})

			It("should reduce the available daily spend balance", func() {
				av, err := Wallet.SpendAvailable(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(av.String()).To(AlmostEqual("99500000000000000000"))
			})

		})

		When("I have one thousand tokens", func() {
			BeforeEach(func() {
				var err error
				tx, err = TKN.Credit(BankAccount.TransactOpts(), WalletAddress, big.NewInt(1000))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})

			When("I transfer 300 tokens to a random person using 'executeTransaction'", func() {
				BeforeEach(func() {
					a, err := abi.JSON(strings.NewReader(ERC20ABI))
					Expect(err).ToNot(HaveOccurred())
					data, err := a.Pack("transfer", RandomAccount.Address(), big.NewInt(300))
					Expect(err).ToNot(HaveOccurred())

					tx, err = Wallet.ExecuteTransaction(Owner.TransactOpts(), TKNAddress, big.NewInt(0), data)
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeTrue())
				})

				It("should increase TKN balance of the random person", func() {
					b, err := TKN.BalanceOf(nil, RandomAccount.Address())
					Expect(err).ToNot(HaveOccurred())
					Expect(b.String()).To(Equal("300"))
				})

				It("should decrease TKN balance of the wallet", func() {
					b, err := Wallet.Balance(nil, TKNAddress)
					Expect(err).ToNot(HaveOccurred())
					Expect(b.String()).To(Equal("700"))
				})

				It("should reduce the available daily spend balance", func() {
					av, err := Wallet.SpendAvailable(nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(av.String()).To(AlmostEqual("99999999999951010000"))
				})
			})

			When("random person is whitelisted", func() {
				BeforeEach(func() {
					tx, err := Wallet.InitializeWhitelist(Owner.TransactOpts(), []common.Address{RandomAccount.Address()})
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeTrue())
				})

				When("I transfer 300 tokens to a random person using 'executeTransaction'", func() {
					BeforeEach(func() {

						a, err := abi.JSON(strings.NewReader(ERC20ABI))
						Expect(err).ToNot(HaveOccurred())
						data, err := a.Pack("transfer", RandomAccount.Address(), big.NewInt(300))
						Expect(err).ToNot(HaveOccurred())

						tx, err = Wallet.ExecuteTransaction(Owner.TransactOpts(), TKNAddress, big.NewInt(0), data)
						Expect(err).ToNot(HaveOccurred())
						Backend.Commit()
						Expect(isSuccessful(tx)).To(BeTrue())
					})

					It("should increase TKN balance of the random person", func() {
						b, err := TKN.BalanceOf(nil, RandomAccount.Address())
						Expect(err).ToNot(HaveOccurred())
						Expect(b.String()).To(Equal("300"))
					})

					It("should decrease TKN balance of the wallet", func() {
						b, err := Wallet.Balance(nil, TKNAddress)
						Expect(err).ToNot(HaveOccurred())
						Expect(b.String()).To(Equal("700"))
					})

					It("should not reduce the available daily spend balance", func() {
						av, err := Wallet.SpendAvailable(nil)
						Expect(err).ToNot(HaveOccurred())
						Expect(av.String()).To(Equal("100000000000000000000"))
					})
				})
			})

			When("I approve 300 tokens to a random person using 'executeTransaction'", func() {
				BeforeEach(func() {
					a, err := abi.JSON(strings.NewReader(ERC20ABI))
					Expect(err).ToNot(HaveOccurred())
					data, err := a.Pack("approve", RandomAccount.Address(), big.NewInt(300))
					Expect(err).ToNot(HaveOccurred())

					tx, err = Wallet.ExecuteTransaction(Owner.TransactOpts(), TKNAddress, big.NewInt(0), data)
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeTrue())
				})

				It("should not increase TKN balance of the random person", func() {
					b, err := TKN.BalanceOf(nil, RandomAccount.Address())
					Expect(err).ToNot(HaveOccurred())
					Expect(b.String()).To(Equal("0"))
				})

				It("should not decrease TKN balance of the wallet", func() {
					b, err := Wallet.Balance(nil, TKNAddress)
					Expect(err).ToNot(HaveOccurred())
					Expect(b.String()).To(Equal("1000"))
				})

				It("should reduce the available daily spend balance", func() {
					av, err := Wallet.SpendAvailable(nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(av.String()).To(AlmostEqual("99999999999951010000"))
				})
			})

			When("random person is whitelisted", func() {
				BeforeEach(func() {
					tx, err := Wallet.InitializeWhitelist(Owner.TransactOpts(), []common.Address{RandomAccount.Address()})
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeTrue())
				})

				When("I approve 300 tokens to a random person using 'executeTransaction'", func() {
					BeforeEach(func() {
						a, err := abi.JSON(strings.NewReader(ERC20ABI))
						Expect(err).ToNot(HaveOccurred())
						data, err := a.Pack("approve", RandomAccount.Address(), big.NewInt(300))
						Expect(err).ToNot(HaveOccurred())

						tx, err = Wallet.ExecuteTransaction(Owner.TransactOpts(), TKNAddress, big.NewInt(0), data)
						Expect(err).ToNot(HaveOccurred())
						Backend.Commit()
						Expect(isSuccessful(tx)).To(BeTrue())
					})

					It("should not increase TKN balance of the random person", func() {
						b, err := TKN.BalanceOf(nil, RandomAccount.Address())
						Expect(err).ToNot(HaveOccurred())
						Expect(b.String()).To(Equal("0"))
					})

					It("should not decrease TKN balance of the wallet", func() {
						b, err := Wallet.Balance(nil, TKNAddress)
						Expect(err).ToNot(HaveOccurred())
						Expect(b.String()).To(Equal("1000"))
					})

					It("should not reduce the available daily spend balance", func() {
						av, err := Wallet.SpendAvailable(nil)
						Expect(err).ToNot(HaveOccurred())
						Expect(av.String()).To(Equal("100000000000000000000"))
					})
				})

			})
		})

	})
})

const ERC20ABI = `[
    {
        "constant": true,
        "inputs": [],
        "name": "name",
        "outputs": [
            {
                "name": "",
                "type": "string"
            }
        ],
        "payable": false,
        "stateMutability": "view",
        "type": "function"
    },
    {
        "constant": false,
        "inputs": [
            {
                "name": "_spender",
                "type": "address"
            },
            {
                "name": "_value",
                "type": "uint256"
            }
        ],
        "name": "approve",
        "outputs": [
            {
                "name": "",
                "type": "bool"
            }
        ],
        "payable": false,
        "stateMutability": "nonpayable",
        "type": "function"
    },
    {
        "constant": true,
        "inputs": [],
        "name": "totalSupply",
        "outputs": [
            {
                "name": "",
                "type": "uint256"
            }
        ],
        "payable": false,
        "stateMutability": "view",
        "type": "function"
    },
    {
        "constant": false,
        "inputs": [
            {
                "name": "_from",
                "type": "address"
            },
            {
                "name": "_to",
                "type": "address"
            },
            {
                "name": "_value",
                "type": "uint256"
            }
        ],
        "name": "transferFrom",
        "outputs": [
            {
                "name": "",
                "type": "bool"
            }
        ],
        "payable": false,
        "stateMutability": "nonpayable",
        "type": "function"
    },
    {
        "constant": true,
        "inputs": [],
        "name": "decimals",
        "outputs": [
            {
                "name": "",
                "type": "uint8"
            }
        ],
        "payable": false,
        "stateMutability": "view",
        "type": "function"
    },
    {
        "constant": true,
        "inputs": [
            {
                "name": "_owner",
                "type": "address"
            }
        ],
        "name": "balanceOf",
        "outputs": [
            {
                "name": "balance",
                "type": "uint256"
            }
        ],
        "payable": false,
        "stateMutability": "view",
        "type": "function"
    },
    {
        "constant": true,
        "inputs": [],
        "name": "symbol",
        "outputs": [
            {
                "name": "",
                "type": "string"
            }
        ],
        "payable": false,
        "stateMutability": "view",
        "type": "function"
    },
    {
        "constant": false,
        "inputs": [
            {
                "name": "_to",
                "type": "address"
            },
            {
                "name": "_value",
                "type": "uint256"
            }
        ],
        "name": "transfer",
        "outputs": [
            {
                "name": "",
                "type": "bool"
            }
        ],
        "payable": false,
        "stateMutability": "nonpayable",
        "type": "function"
    },
    {
        "constant": true,
        "inputs": [
            {
                "name": "_owner",
                "type": "address"
            },
            {
                "name": "_spender",
                "type": "address"
            }
        ],
        "name": "allowance",
        "outputs": [
            {
                "name": "",
                "type": "uint256"
            }
        ],
        "payable": false,
        "stateMutability": "view",
        "type": "function"
    },
    {
        "payable": true,
        "stateMutability": "payable",
        "type": "fallback"
    },
    {
        "anonymous": false,
        "inputs": [
            {
                "indexed": true,
                "name": "owner",
                "type": "address"
            },
            {
                "indexed": true,
                "name": "spender",
                "type": "address"
            },
            {
                "indexed": false,
                "name": "value",
                "type": "uint256"
            }
        ],
        "name": "Approval",
        "type": "event"
    },
    {
        "anonymous": false,
        "inputs": [
            {
                "indexed": true,
                "name": "from",
                "type": "address"
            },
            {
                "indexed": true,
                "name": "to",
                "type": "address"
            },
            {
                "indexed": false,
                "name": "value",
                "type": "uint256"
            }
        ],
        "name": "Transfer",
        "type": "event"
    }
]`
