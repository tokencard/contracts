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

		When("I batch a normal and an ERC20 transfer", func() {

			var randomAddress common.Address

			When("the destination address is not whitelisted", func() {
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
        })
    })
})
