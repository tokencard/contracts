package wallet_test

import (
	"context"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tokencard/contracts/v2/test/shared"
    "github.com/tokencard/contracts/v2/pkg/bindings"
	"github.com/tokencard/ethertest"
)

var _ = Describe("executeTransaction", func() {

	Context("when the wallet has enough ETH", func() {
		BeforeEach(func() {
			BankAccount.MustTransfer(Backend, WalletAddress, EthToWei(101))
		})

		var tx *types.Transaction

		When("I transfer 500 Finney to a random address using 'executeTransaction'", func() {

			var randomAddress common.Address
			var spendLimit *big.Int
			var err error

			When("the destination flag is set correctly (not a contract)", func() {
                When("the destination address is not whitelisted", func() {
    				BeforeEach(func() {
    					spendLimit, err = Wallet.SpendLimitAvailable(nil)
    					Expect(err).ToNot(HaveOccurred())
    					privateKey, err := crypto.GenerateKey()
    					randomAddress = crypto.PubkeyToAddress(privateKey.PublicKey)
    					tx, err = Wallet.ExecuteTransaction(Owner.TransactOpts(ethertest.WithGasLimit(100000)), randomAddress, FinneyToWei(500), nil, false)
    					Expect(err).ToNot(HaveOccurred())
    					Backend.Commit()
    					Expect(isSuccessful(tx)).To(BeTrue())
    				})

    				It("should increase random address' balance by the same amount", func() {
    					b, e := Backend.BalanceAt(context.Background(), randomAddress, nil)
    					Expect(e).ToNot(HaveOccurred())
    					Expect(b.String()).To(Equal(FinneyToWei(500).String()))
    				})

    				It("should reduce the available daily spend balance", func() {
    					spendLimit.Sub(spendLimit, FinneyToWei(500))
    					sl, err := Wallet.SpendLimitAvailable(nil)
    					Expect(err).ToNot(HaveOccurred())
    					Expect(sl.String()).To(Equal(spendLimit.String()))
    				})
                })

                When("the destination address is whitelisted", func() {

                    privateKey, _ := crypto.GenerateKey()
                    randomAddress = crypto.PubkeyToAddress(privateKey.PublicKey)

    				BeforeEach(func() {
    					tx, err := Wallet.SetWhitelist(Owner.TransactOpts(), []common.Address{randomAddress})
    					Expect(err).ToNot(HaveOccurred())
    					Backend.Commit()
    					Expect(isSuccessful(tx)).To(BeTrue())
    				})

                    BeforeEach(func() {
    					spendLimit, err = Wallet.SpendLimitAvailable(nil)
    					Expect(err).ToNot(HaveOccurred())
    					tx, err = Wallet.ExecuteTransaction(Owner.TransactOpts(ethertest.WithGasLimit(100000)), randomAddress, FinneyToWei(500), nil, false)
    					Expect(err).ToNot(HaveOccurred())
    					Backend.Commit()
    					Expect(isSuccessful(tx)).To(BeTrue())
    				})

    				It("should increase random address' balance by the same amount", func() {
    					b, e := Backend.BalanceAt(context.Background(), randomAddress, nil)
    					Expect(e).ToNot(HaveOccurred())
    					Expect(b.String()).To(Equal(FinneyToWei(500).String()))
    				})

    				It("should NOT reduce the available daily spend balance", func() {
    					sl, err := Wallet.SpendLimitAvailable(nil)
    					Expect(err).ToNot(HaveOccurred())
    					Expect(sl.String()).To(Equal(spendLimit.String()))
    				})
                })
			})

			When("the destination flag is ΝΟΤ set correctly (contract instead of EOA)", func() {
				It("should fail", func() {
					privateKey, err := crypto.GenerateKey()
					randomAddress = crypto.PubkeyToAddress(privateKey.PublicKey)
					tx, err = Wallet.ExecuteTransaction(Owner.TransactOpts(ethertest.WithGasLimit(100000)), randomAddress, FinneyToWei(500), nil, true)
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeFalse())
				})
			})
		})

		When("I have one thousand TKN tokens", func() {
			BeforeEach(func() {
				var err error
				tx, err = TKN.Credit(BankAccount.TransactOpts(), WalletAddress, big.NewInt(1000))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})

			When("I transfer 300 tokens to a random account using 'executeTransaction'", func() {
				BeforeEach(func() {
					a, err := abi.JSON(strings.NewReader(ERC20ABI))
					Expect(err).ToNot(HaveOccurred())
					data, err := a.Pack("transfer", RandomAccount.Address(), big.NewInt(300))
					Expect(err).ToNot(HaveOccurred())

					tx, err = Wallet.ExecuteTransaction(Owner.TransactOpts(), TKNAddress, big.NewInt(0), data, true)
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeTrue())
				})

				It("should increase the TKN balance of the random account", func() {
					b, err := TKN.BalanceOf(nil, RandomAccount.Address())
					Expect(err).ToNot(HaveOccurred())
					Expect(b.String()).To(Equal("300"))
				})

				It("should decrease the TKN balance of the wallet", func() {
                    b, err := TKN.BalanceOf(nil, WalletAddress)
					Expect(err).ToNot(HaveOccurred())
					Expect(b.String()).To(Equal("700"))
				})

				It("should reduce the available daily spend balance", func() {
					av, err := Wallet.SpendLimitAvailable(nil)
					Expect(err).ToNot(HaveOccurred())
                    eth, err := Wallet.ConvertToEther(nil, TKNAddress, big.NewInt(300))
                    Expect(err).ToNot(HaveOccurred())
                    av.Sub(av, eth) //subtract converted eth from dailySppendLimit
					Expect(av.String()).To(Equal(av.String()))
				})
			})

            When("I try to use a non-whitelisted method on a whitelisted/protected token address", func() {
				It("should fail", func() {
					a, err := abi.JSON(strings.NewReader(ERC20ABI))
					Expect(err).ToNot(HaveOccurred())
					data, err := a.Pack("increaseApproval", RandomAccount.Address(), big.NewInt(300))
					Expect(err).ToNot(HaveOccurred())

					tx, err = Wallet.ExecuteTransaction(Owner.TransactOpts(ethertest.WithGasLimit(100000)), TKNAddress, big.NewInt(0), data, true)
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeFalse())
				})
			})

            When("I try to use 'transfer' but data (i.e. value is corrupt) is missing", func() {
				It("should fail", func() {
					a, err := abi.JSON(strings.NewReader(ERC20ABI))
					Expect(err).ToNot(HaveOccurred())
					data, err := a.Pack("transfer", RandomAccount.Address(), big.NewInt(300))
					Expect(err).ToNot(HaveOccurred())
					tx, err = Wallet.ExecuteTransaction(Owner.TransactOpts(ethertest.WithGasLimit(100000)), TKNAddress, big.NewInt(0), data[:67], true)
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeFalse())
				})
			})

            When("I try to use 'approve' but data (i.e. value) is missing", func() {
				It("should fail", func() {
					a, err := abi.JSON(strings.NewReader(ERC20ABI))
					Expect(err).ToNot(HaveOccurred())
					data, err := a.Pack("approve", RandomAccount.Address(), big.NewInt(300))
					Expect(err).ToNot(HaveOccurred())
					tx, err = Wallet.ExecuteTransaction(Owner.TransactOpts(ethertest.WithGasLimit(100000)), TKNAddress, big.NewInt(0), data[:36], true)
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeFalse())
				})
			})

            When("I try to use 'transferFrom' but data (i.e. value is corrupt) is missing", func() {
				It("should fail", func() {
					a, err := abi.JSON(strings.NewReader(ERC20ABI))
					Expect(err).ToNot(HaveOccurred())
                    data, err := a.Pack("transferFrom", WalletAddress, Owner.Address(), big.NewInt(300))
					Expect(err).ToNot(HaveOccurred())
                    //transferFrom needs 100 bytes: 4(methodID) + 32 (from) + 32 (to) + 32 (value)
					tx, err = Wallet.ExecuteTransaction(Owner.TransactOpts(ethertest.WithGasLimit(100000)), TKNAddress, big.NewInt(0), data[:99], true)
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeFalse())
				})
			})

			When("I send data (transfer 300 tokens to a random account to an EOA using 'executeTransaction'", func() {
				It("should succeed", func() {
					a, err := abi.JSON(strings.NewReader(ERC20ABI))
					Expect(err).ToNot(HaveOccurred())
					data, err := a.Pack("transfer", RandomAccount.Address(), big.NewInt(300))
					Expect(err).ToNot(HaveOccurred())

					tx, err = Wallet.ExecuteTransaction(Owner.TransactOpts(), RandomAccount.Address(), big.NewInt(0), data, false)
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeTrue())
				})

			})

			When("I send data (transfer 300 tokens to a random account using 'executeTransaction' but the destination flag is not set correctly (isContract = true", func() {
				It("should fail", func() {
					a, err := abi.JSON(strings.NewReader(ERC20ABI))
					Expect(err).ToNot(HaveOccurred())
					data, err := a.Pack("transfer", RandomAccount.Address(), big.NewInt(300))
					Expect(err).ToNot(HaveOccurred())

					tx, err = Wallet.ExecuteTransaction(Owner.TransactOpts(ethertest.WithGasLimit(100000)), TKNAddress, big.NewInt(0), data, false)
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeFalse())
				})

			})

			When("random account is whitelisted", func() {
				BeforeEach(func() {
					tx, err := Wallet.SetWhitelist(Owner.TransactOpts(), []common.Address{RandomAccount.Address()})
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeTrue())
				})

				When("I transfer 300 tokens to a random account using 'executeTransaction'", func() {
					BeforeEach(func() {

						a, err := abi.JSON(strings.NewReader(ERC20ABI))
						Expect(err).ToNot(HaveOccurred())
						data, err := a.Pack("transfer", RandomAccount.Address(), big.NewInt(300))
						Expect(err).ToNot(HaveOccurred())

						tx, err = Wallet.ExecuteTransaction(Owner.TransactOpts(), TKNAddress, big.NewInt(0), data, true)
						Expect(err).ToNot(HaveOccurred())
						Backend.Commit()
						Expect(isSuccessful(tx)).To(BeTrue())
					})

					It("should increase the TKN balance of the random account", func() {
						b, err := TKN.BalanceOf(nil, RandomAccount.Address())
						Expect(err).ToNot(HaveOccurred())
						Expect(b.String()).To(Equal("300"))
					})

					It("should decrease the TKN balance of the wallet", func() {
						b, err := TKN.BalanceOf(nil, WalletAddress)
						Expect(err).ToNot(HaveOccurred())
						Expect(b.String()).To(Equal("700"))
					})

					It("should not reduce the available daily spend balance", func() {
						av, err := Wallet.SpendLimitAvailable(nil)
						Expect(err).ToNot(HaveOccurred())
						Expect(av.String()).To(Equal("100000000000000000000"))
					})
				})
			})

			When("I approve 300 tokens to a random wallet using 'executeTransaction'", func() {

                var RandomWallet *bindings.Wallet
                var RandomWalletAddress common.Address

                BeforeEach(func() {
                    var tx *types.Transaction
                    var err error
                    RandomWalletAddress, tx, RandomWallet, err = bindings.DeployWallet(BankAccount.TransactOpts(), Backend, RandomAccount.Address(), false, ENSRegistryAddress, TokenWhitelistName, ControllerName, LicenceName, EthToWei(1))
                	Expect(err).ToNot(HaveOccurred())
                	Backend.Commit()
                	Expect(isSuccessful(tx)).To(BeTrue())
                })

				BeforeEach(func() {
					a, err := abi.JSON(strings.NewReader(ERC20ABI))
					Expect(err).ToNot(HaveOccurred())
					data, err := a.Pack("approve", RandomWalletAddress, big.NewInt(300))
					Expect(err).ToNot(HaveOccurred())

					tx, err = Wallet.ExecuteTransaction(Owner.TransactOpts(), TKNAddress, big.NewInt(0), data, true)
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeTrue())
				})

				It("should not increase the TKN balance of the random wallet", func() {
					b, err := TKN.BalanceOf(nil, RandomWalletAddress)
					Expect(err).ToNot(HaveOccurred())
					Expect(b.String()).To(Equal("0"))
				})

				It("should not decrease the TKN balance of the wallet", func() {
					b, err := TKN.BalanceOf(nil, WalletAddress)
					Expect(err).ToNot(HaveOccurred())
					Expect(b.String()).To(Equal("1000"))
				})

				It("should reduce the available daily spend balance", func() {
					av, err := Wallet.SpendLimitAvailable(nil)
					Expect(err).ToNot(HaveOccurred())
                    eth, err := Wallet.ConvertToEther(nil, TKNAddress, big.NewInt(300))
                    Expect(err).ToNot(HaveOccurred())
                    av.Sub(av, eth) //subtract converted eth from dailySppendLimit
					Expect(av.String()).To(Equal(av.String()))
				})

                When("the approved random wallet 'transfersFrom' all the tokens to a whitelisted destination address using 'executeTransaction'", func() {

                    BeforeEach(func() {
    					tx, err := RandomWallet.SetWhitelist(RandomAccount.TransactOpts(), []common.Address{Owner.Address()})
    					Expect(err).ToNot(HaveOccurred())
    					Backend.Commit()
    					Expect(isSuccessful(tx)).To(BeTrue())
    				})

                    BeforeEach(func() {
    					a, err := abi.JSON(strings.NewReader(ERC20ABI))
    					Expect(err).ToNot(HaveOccurred())
    					data, err := a.Pack("transferFrom", WalletAddress, Owner.Address(), big.NewInt(300))
    					Expect(err).ToNot(HaveOccurred())

    					tx, err = RandomWallet.ExecuteTransaction(RandomAccount.TransactOpts(), TKNAddress, big.NewInt(0), data, true)
    					Expect(err).ToNot(HaveOccurred())
    					Backend.Commit()
    					Expect(isSuccessful(tx)).To(BeTrue())
    				})

                    It("should increase the TKN balance of the whitelisted destination", func() {
    					b, err := TKN.BalanceOf(nil, Owner.Address())
    					Expect(err).ToNot(HaveOccurred())
    					Expect(b.String()).To(Equal("300"))
    				})

    				It("should decrease the TKN balance of the wallet", func() {
    					b, err := TKN.BalanceOf(nil, WalletAddress)
    					Expect(err).ToNot(HaveOccurred())
    					Expect(b.String()).To(Equal("700"))
    				})

    				It("should NOT reduce the available daily spend balance", func() {
    					av, err := RandomWallet.SpendLimitAvailable(nil)
    					Expect(err).ToNot(HaveOccurred())
    					Expect(av.String()).To(Equal(EthToWei(1).String()))
    				})
                })

                When("the approved random wallet 'transfersFrom' all the tokens to itself using 'executeTransaction'", func() {
                    BeforeEach(func() {
    					a, err := abi.JSON(strings.NewReader(ERC20ABI))
    					Expect(err).ToNot(HaveOccurred())
    					data, err := a.Pack("transferFrom", WalletAddress, RandomWalletAddress, big.NewInt(300))
    					Expect(err).ToNot(HaveOccurred())

    					tx, err = RandomWallet.ExecuteTransaction(RandomAccount.TransactOpts(), TKNAddress, big.NewInt(0), data, true)
    					Expect(err).ToNot(HaveOccurred())
    					Backend.Commit()
    					Expect(isSuccessful(tx)).To(BeTrue())
    				})

                    It("should increase the TKN balance of the random wallet", func() {
    					b, err := TKN.BalanceOf(nil, RandomWalletAddress)
    					Expect(err).ToNot(HaveOccurred())
    					Expect(b.String()).To(Equal("300"))
    				})

    				It("should decrease the TKN balance of the wallet", func() {
    					b, err := TKN.BalanceOf(nil, WalletAddress)
    					Expect(err).ToNot(HaveOccurred())
    					Expect(b.String()).To(Equal("700"))
    				})

    				It("should reduce the available daily spend balance", func() {
    					av, err := RandomWallet.SpendLimitAvailable(nil)
    					Expect(err).ToNot(HaveOccurred())
                        eth, err := Wallet.ConvertToEther(nil, TKNAddress, big.NewInt(300))
                        Expect(err).ToNot(HaveOccurred())
                        av.Sub(av, eth) //subtract converted eth from dailySppendLimit
    					Expect(av.String()).To(Equal(av.String()))
    				})
                })

                When("the approved random wallet tries to 'transferFrom' more than the approved ammount to itself using 'executeTransaction'", func() {
                    It("should fail", func() {
    					a, err := abi.JSON(strings.NewReader(ERC20ABI))
    					Expect(err).ToNot(HaveOccurred())
    					data, err := a.Pack("transferFrom", WalletAddress, RandomWalletAddress, big.NewInt(301))
    					Expect(err).ToNot(HaveOccurred())

    					tx, err = RandomWallet.ExecuteTransaction(RandomAccount.TransactOpts(ethertest.WithGasLimit(500000)), TKNAddress, big.NewInt(0), data, true)
    					Expect(err).ToNot(HaveOccurred())
    					Backend.Commit()
    					Expect(isSuccessful(tx)).To(BeFalse())
    				})
                })
			})

			When("random account is whitelisted", func() {
				BeforeEach(func() {
					tx, err := Wallet.SetWhitelist(Owner.TransactOpts(), []common.Address{RandomAccount.Address()})
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeTrue())
				})

				When("I approve 300 tokens to a random account using 'executeTransaction'", func() {
					BeforeEach(func() {
						a, err := abi.JSON(strings.NewReader(ERC20ABI))
						Expect(err).ToNot(HaveOccurred())
						data, err := a.Pack("approve", RandomAccount.Address(), big.NewInt(300))
						Expect(err).ToNot(HaveOccurred())

						tx, err = Wallet.ExecuteTransaction(Owner.TransactOpts(), TKNAddress, big.NewInt(0), data, true)
						Expect(err).ToNot(HaveOccurred())
						Backend.Commit()
						Expect(isSuccessful(tx)).To(BeTrue())
					})

					It("should not increase TKN balance of the random account", func() {
						b, err := TKN.BalanceOf(nil, RandomAccount.Address())
						Expect(err).ToNot(HaveOccurred())
						Expect(b.String()).To(Equal("0"))
					})

					It("should not decrease TKN balance of the wallet", func() {
						b, err := TKN.BalanceOf(nil, WalletAddress)
						Expect(err).ToNot(HaveOccurred())
						Expect(b.String()).To(Equal("1000"))
					})

					It("should not reduce the available daily spend balance", func() {
						av, err := Wallet.SpendLimitAvailable(nil)
						Expect(err).ToNot(HaveOccurred())
						Expect(av.String()).To(Equal("100000000000000000000"))
					})
				})
			})
		})

        When("I have one thousand random tokens of a non-whitelisted ERC20 conctract", func() {
			BeforeEach(func() {
				var err error
				tx, err = ERC20Contract1.Credit(BankAccount.TransactOpts(), WalletAddress, big.NewInt(1000))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})

			When("I transfer 300 tokens and 1 ETH to a random account using 'executeTransaction'", func() {
				BeforeEach(func() {
					a, err := abi.JSON(strings.NewReader(ERC20ABI))
					Expect(err).ToNot(HaveOccurred())
					data, err := a.Pack("transfer", RandomAccount.Address(), big.NewInt(300))
					Expect(err).ToNot(HaveOccurred())

					tx, err = Wallet.ExecuteTransaction(Owner.TransactOpts(), ERC20Contract1Address, big.NewInt(0), data, true)
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeTrue())
				})

				It("should increase ERC20 balance of the random account", func() {
					b, err := ERC20Contract1.BalanceOf(nil, RandomAccount.Address())
					Expect(err).ToNot(HaveOccurred())
					Expect(b.String()).To(Equal("300"))
				})

				It("should decrease the ERC20 balance of the wallet", func() {
                    b, err := ERC20Contract1.BalanceOf(nil, WalletAddress)
					Expect(err).ToNot(HaveOccurred())
					Expect(b.String()).To(Equal("700"))
				})

				It("should NOT reduce the available daily spend balance", func() {
					av, err := Wallet.SpendLimitAvailable(nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(av.String()).To(Equal(EthToWei(100).String()))
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
        "constant": false,
        "inputs": [
            {
                "name": "_spender",
                "type": "address"
            },
            {
                "name": "_addedValue",
                "type": "uint256"
            }
        ],
        "name": "increaseApproval",
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
