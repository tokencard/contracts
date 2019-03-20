package wallet_test

import (
	"context"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tokencard/contracts/test/shared"
	"github.com/tokencard/ethertest"
)

var _ = Describe("topUpGas", func() {

	Context("when the wallet has enough ETH", func() {
		BeforeEach(func() {
			BankAccount.MustTransfer(Backend, WalletAddress, EthToWei(10))
		})

		var tx *types.Transaction
		var err error
		var caller *ethertest.Account
		var ownerBalance *big.Int

		BeforeEach(func() {
			caller = Controller
			BankAccount.MustTransfer(Backend, Controller.Address(), EthToWei(1))
			ownerBalance, err = Backend.BalanceAt(context.Background(), Owner.Address(), nil)
			Expect(err).ToNot(HaveOccurred())
		})

		Context("When called by the wallet controller and is lower than top up limit", func() {

			BeforeEach(func() {
				tx, err = Wallet.TopUpGas(caller.TransactOpts(ethertest.WithGasLimit(181000)), FinneyToWei(1))
				Backend.Commit()
			})

			It("Should not return error", func() {
				Expect(err).ToNot(HaveOccurred())
			})

			It("should succeed", func() {
				Expect(isSuccessful(tx)).To(BeTrue())
			})

			It("should top up the gas succesfully", func() {
				b, e := Backend.BalanceAt(context.Background(), WalletAddress, nil)
				Expect(e).ToNot(HaveOccurred())
				newBalance := new(big.Int).Sub(EthToWei(10), FinneyToWei(1))
				Expect(b.String()).To(Equal(newBalance.String()))
				ownerBalance.Add(ownerBalance, FinneyToWei(1))
				Expect(Owner.Balance(Backend).String()).To(Equal(ownerBalance.String()))
			})

		})

		Context("When the value is above top up limit", func() {

			BeforeEach(func() {
				tx, err = Wallet.TopUpGas(caller.TransactOpts(ethertest.WithGasLimit(81000)), EthToWei(1))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
			})

			It("should not return error", func() {
				Expect(err).ToNot(HaveOccurred())
			})

			It("should fail", func() {
				Expect(isSuccessful(tx)).To(BeFalse())
			})

			It("should NOT top up the gas", func() {
				b, e := Backend.BalanceAt(context.Background(), WalletAddress, nil)
				Expect(e).ToNot(HaveOccurred())
				Expect(b.String()).To(Equal(EthToWei(10).String())) //Wallet address has initially 10 ETH
				Expect(Owner.Balance(Backend).String()).To(Equal(ownerBalance.String()))
			})

		})

		Context("When daily limit has been exausted", func() {
			BeforeEach(func() {
				caller = Controller
				BankAccount.MustTransfer(Backend, Controller.Address(), EthToWei(1))
			})

			BeforeEach(func() {
				tx, err = Wallet.TopUpGas(caller.TransactOpts(ethertest.WithGasLimit(81000)), FinneyToWei(500))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})

			BeforeEach(func() {
				tx, err = Wallet.TopUpGas(caller.TransactOpts(ethertest.WithGasLimit(81000)), FinneyToWei(1))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
			})

			It("should fail", func() {
				Expect(isSuccessful(tx)).To(BeFalse())
			})

			Context("When I wait for one day", func() {
				BeforeEach(func() {
					Backend.AdjustTime(time.Hour*24 + time.Second)
					Backend.Commit()
				})

				BeforeEach(func() {
					tx, err = Wallet.TopUpGas(caller.TransactOpts(ethertest.WithGasLimit(165000)), FinneyToWei(1))
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
				})

				It("should not return error", func() {
					Expect(err).ToNot(HaveOccurred())
				})

				It("should not fail", func() {
					Expect(isSuccessful(tx)).To(BeTrue())
				})

				It("should top up only the top up limit of the gas + 1 Finney", func() {

					b, e := Backend.BalanceAt(context.Background(), WalletAddress, nil)
					Expect(e).ToNot(HaveOccurred())
					newBalance := new(big.Int).Sub(EthToWei(10), FinneyToWei(501)) //Wallet address has initially 10 ETH
					Expect(b.String()).To(Equal(newBalance.String()))
					ownerBalance.Add(ownerBalance, FinneyToWei(501))
					Expect(Owner.Balance(Backend).String()).To(Equal(ownerBalance.String()))
				})

			})

		})

		Context("When called by the wallet owner", func() {

			var txCost *big.Int

			BeforeEach(func() {
				caller = Owner
			})

			Context("When the value is lower than top up limit", func() {

				BeforeEach(func() {
					tx, err = Wallet.TopUpGas(caller.TransactOpts(ethertest.WithGasLimit(81000)), FinneyToWei(500))
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeTrue())
					r, err := Backend.TransactionReceipt(context.Background(), tx.Hash())
					Expect(err).ToNot(HaveOccurred())
					txCost = new(big.Int).Mul(tx.GasPrice(), big.NewInt(int64(r.GasUsed)))
				})

				It("should top up succesfully", func() {
					b, e := Backend.BalanceAt(context.Background(), WalletAddress, nil)
					Expect(e).ToNot(HaveOccurred())
					newBalance := new(big.Int).Sub(EthToWei(10), FinneyToWei(500)) //Wallet address has initially 10 ETH
					Expect(b.String()).To(Equal(newBalance.String()))
					ownerBalance.Add(ownerBalance, FinneyToWei(500))
					ownerBalance.Sub(ownerBalance, txCost) //subtract the tx cost
					Expect(Owner.Balance(Backend).String()).To(Equal(ownerBalance.String()))
				})

			})

			Context("When the value is above top up limit", func() {

				BeforeEach(func() {
					tx, err = Wallet.TopUpGas(caller.TransactOpts(ethertest.WithGasLimit(81000)), FinneyToWei(800))
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeFalse())
					r, err := Backend.TransactionReceipt(context.Background(), tx.Hash())
					Expect(err).ToNot(HaveOccurred())
					txCost = new(big.Int).Mul(tx.GasPrice(), big.NewInt(int64(r.GasUsed)))
				})

				It("should NOT top up", func() {
					b, e := Backend.BalanceAt(context.Background(), WalletAddress, nil)
					Expect(e).ToNot(HaveOccurred())
					Expect(b.String()).To(Equal(EthToWei(10).String())) //Wallet address has initially 10 ETH
					ownerBalance.Sub(ownerBalance, txCost)              //subtract the tx cost
					Expect(Owner.Balance(Backend).String()).To(Equal(ownerBalance.String()))
				})

			})

			Context("When daily limit has been exausted", func() {

				BeforeEach(func() {
					tx, err = Wallet.TopUpGas(caller.TransactOpts(ethertest.WithGasLimit(81000)), FinneyToWei(500))
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					r, err := Backend.TransactionReceipt(context.Background(), tx.Hash())
					Expect(err).ToNot(HaveOccurred())
					txCost = new(big.Int).Mul(tx.GasPrice(), big.NewInt(int64(r.GasUsed)))
				})

				BeforeEach(func() {
					tx, err = Wallet.TopUpGas(caller.TransactOpts(ethertest.WithGasLimit(165000)), FinneyToWei(1))
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeFalse())
					r, err := Backend.TransactionReceipt(context.Background(), tx.Hash())
					Expect(err).ToNot(HaveOccurred())
					tempCost := new(big.Int).Mul(tx.GasPrice(), big.NewInt(int64(r.GasUsed)))
					txCost.Add(txCost, tempCost) //accumulate cost
				})

				Context("When I wait for one day", func() {
					BeforeEach(func() {
						Backend.AdjustTime(time.Hour*24 + time.Second)
						Backend.Commit()
					})

					BeforeEach(func() {
						tx, err = Wallet.TopUpGas(caller.TransactOpts(ethertest.WithGasLimit(165000)), FinneyToWei(1))
						Expect(err).ToNot(HaveOccurred())
						Backend.Commit()
						Expect(isSuccessful(tx)).To(BeTrue())
						r, err := Backend.TransactionReceipt(context.Background(), tx.Hash())
						Expect(err).ToNot(HaveOccurred())
						tempCost := new(big.Int).Mul(tx.GasPrice(), big.NewInt(int64(r.GasUsed)))
						txCost.Add(txCost, tempCost) //accumulate cost
					})

					It("should top up only the top up limit of the gas", func() {
						b, e := Backend.BalanceAt(context.Background(), WalletAddress, nil)
						Expect(e).ToNot(HaveOccurred())
						newBalance := new(big.Int).Sub(EthToWei(10), FinneyToWei(501)) //Wallet address has initially 10 ETH
						Expect(b.String()).To(Equal(newBalance.String()))
						ownerBalance.Add(ownerBalance, FinneyToWei(501))
						ownerBalance.Sub(ownerBalance, txCost) //subtract the tx cost
						Expect(Owner.Balance(Backend).String()).To(Equal(ownerBalance.String()))
					})

				})
			})
		})

		Context("When called by some random address and is lower than top up limit", func() {

			BeforeEach(func() {
				caller = RandomAccount
			})

			BeforeEach(func() {
				tx, err = Wallet.TopUpGas(caller.TransactOpts(ethertest.WithGasLimit(165000)), EthToWei(1))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
			})

			It("should not return error", func() {
				Expect(err).ToNot(HaveOccurred())
			})

			It("should fail", func() {
				Expect(isSuccessful(tx)).To(BeFalse())
			})
		})

	})
})
