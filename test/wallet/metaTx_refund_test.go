package wallet_test

import (
	"context"
	"math/big"
	"strings"
    "fmt"
    "crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tokencard/contracts/v2/test/shared"
)

var _ = Describe("metaTx refund", func() {

	Context("when the wallet has enough ETH and ERC20 tokens", func() {
		BeforeEach(func() {
			BankAccount.MustTransfer(Backend, WalletAddress, EthToWei(2))

			var err error
			tx, err := TKNBurner.Mint(BankAccount.TransactOpts(), WalletAddress, big.NewInt(300))
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeTrue())
		})

		When("controller relays an ETH transfer and a refund", func() {

            var randomAddress common.Address
            var tokenBank common.Address
            var privateKey  *ecdsa.PrivateKey

            //transfer ownership to a newly generated daddress (makes signing easier)
            BeforeEach(func() {
                privateKey, _ = crypto.GenerateKey()
                tokenBank = crypto.PubkeyToAddress(privateKey.PublicKey)
                privateKey, _ = crypto.GenerateKey()
                randomAddress = crypto.PubkeyToAddress(privateKey.PublicKey)
                tx, err := Wallet.TransferOwnership(Owner.TransactOpts(), randomAddress, false)
                Expect(err).ToNot(HaveOccurred())
                Backend.Commit()
                Expect(isSuccessful(tx)).To(BeTrue())

            })

			BeforeEach(func() {

                batch := fmt.Sprintf("%s%s%s%s%s%s", randomAddress, abi.U256(EthToWei(1)), abi.U256(big.NewInt(0)),
                tokenBank, abi.U256(big.NewInt(1000)), abi.U256(big.NewInt(0)))

                a, err := abi.JSON(strings.NewReader(WALLET_ABI))
                Expect(err).ToNot(HaveOccurred())
                data, err := a.Pack("executeTransactions", []byte(batch))
                Expect(err).ToNot(HaveOccurred())

                nonce := big.NewInt(0)
                signature, err := SignData(nonce, data, privateKey)
                Expect(err).ToNot(HaveOccurred())

                tx, err := Wallet.ExecuteRelayedTransaction(Controller.TransactOpts(), nonce, data, signature)
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
    })
})
