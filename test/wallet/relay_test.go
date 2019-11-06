package wallet_test

import (
    "strings"
    "crypto/ecdsa"
    "math/big"
    "fmt"
    "errors"
    "context"

	"github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/core/types"
    "github.com/ethereum/go-ethereum/accounts/abi"
    "github.com/ethereum/go-ethereum/crypto"
    "github.com/ethereum/go-ethereum"
    "github.com/tokencard/ethertest"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tokencard/contracts/v2/test/shared"
)


func SignData(nonce *big.Int, data []byte, prv *ecdsa.PrivateKey) ([]byte, error) {
    relayMessage := fmt.Sprintf("rlx:%s%s", abi.U256(nonce), data)
    hash := crypto.Keccak256([]byte(relayMessage))
    ethMessage := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(hash), hash)
    hash = crypto.Keccak256([]byte(ethMessage))
	sig, err := crypto.Sign(hash, prv)
	if err != nil {
		return nil, err
	}
    if len(sig) != 65 {
        return nil, errors.New("invalid sig len")
    }
    sig[64] += 27
	return sig, nil
}

var _ = Describe("relay Tx", func() {

    BeforeEach(func() {
        BankAccount.MustTransfer(Backend, WalletAddress, EthToWei(2))
    })

	When("a random acount tries to relay", func() {
        It("should fail", func() {
            a, err := abi.JSON(strings.NewReader(WALLET_ABI))
            Expect(err).ToNot(HaveOccurred())
            privateKey, _ := crypto.GenerateKey()
            randomAddress := crypto.PubkeyToAddress(privateKey.PublicKey)
            data, err := a.Pack("transfer", randomAddress, common.HexToAddress("0x0"), EthToWei(1))
            Expect(err).ToNot(HaveOccurred())

            nonce := big.NewInt(0)
            signature, err := SignData(nonce, data, Owner.PrivKey())
            Expect(err).ToNot(HaveOccurred())

            tx, err := Wallet.ExecuteRelayedTransaction(RandomAccount.TransactOpts(ethertest.WithGasLimit(500000)), nonce, data, signature)
            Expect(err).ToNot(HaveOccurred())
            Backend.Commit()
            Expect(isSuccessful(tx)).To(BeFalse())
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
            signature, err := SignData(nonce, data, privateKey)
            Expect(err).ToNot(HaveOccurred())

            tx, err := Wallet.ExecuteRelayedTransaction(Controller.TransactOpts(ethertest.WithGasLimit(500000)), nonce, data, signature)
            Expect(err).ToNot(HaveOccurred())
            Backend.Commit()
            Expect(isSuccessful(tx)).To(BeFalse())
        })
    })

    When("a controller tries to relay an owner signed transaction", func() {
        var randomAddress common.Address
        var privateKey  *ecdsa.PrivateKey

        //transfer ownership to a newly generated daddress
        BeforeEach(func() {
            privateKey, _ = crypto.GenerateKey()
            randomAddress = crypto.PubkeyToAddress(privateKey.PublicKey)
            tx, err := Wallet.TransferOwnership(Owner.TransactOpts(), randomAddress, false)
            Expect(err).ToNot(HaveOccurred())
            Backend.Commit()
            Expect(isSuccessful(tx)).To(BeTrue())

        })

        When("the call succeeds", func() {
            BeforeEach(func() {
                a, err := abi.JSON(strings.NewReader(WALLET_ABI))
                Expect(err).ToNot(HaveOccurred())
                data, err := a.Pack("transfer", randomAddress, common.HexToAddress("0x0"), EthToWei(1))
                Expect(err).ToNot(HaveOccurred())

                nonce := big.NewInt(0)
                signature, err := SignData(nonce, data, privateKey)
                Expect(err).ToNot(HaveOccurred())

                tx, err := Wallet.ExecuteRelayedTransaction(Controller.TransactOpts(), nonce, data, signature)
                Expect(err).ToNot(HaveOccurred())
                Backend.Commit()
                Expect(isSuccessful(tx)).To(BeTrue())
            })

            It("should change the owner", func() {
                o, err := Wallet.Owner(nil)
                Expect(err).ToNot(HaveOccurred())
                Expect(o).To(Equal(randomAddress))
            })

            It("should emit an ExecutedRelayedTransaction event", func() {
                it, err := Wallet.FilterExecutedRelayedTransaction(nil)
                Expect(err).ToNot(HaveOccurred())
                Expect(it.Next()).To(BeTrue())
                evt := it.Event
                Expect(it.Next()).To(BeFalse())
                a, err := abi.JSON(strings.NewReader(WALLET_ABI))
                Expect(err).ToNot(HaveOccurred())
                data, err := a.Pack("transfer", randomAddress, common.HexToAddress("0x0"), EthToWei(1))
                Expect(evt.Data).To(Equal(data))
                Expect(evt.Returndata).To(Equal([]byte{}))
            })

            It("should decrease the wallet's ETH balance ", func() {
    			b, err := Backend.BalanceAt(context.Background(), WalletAddress, nil)
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
                signature, err := SignData(nonce, data, privateKey)
                Expect(err).ToNot(HaveOccurred())

                tx, err := Wallet.ExecuteRelayedTransaction(Controller.TransactOpts(ethertest.WithGasLimit(500000)), nonce, data, signature)
                Expect(err).ToNot(HaveOccurred())
                Backend.Commit()
                Expect(isSuccessful(tx)).To(BeFalse())
            })

            It("should succeed when increasing the nonce", func() {
                a, err := abi.JSON(strings.NewReader(WALLET_ABI))
                Expect(err).ToNot(HaveOccurred())
                data, err := a.Pack("transfer", randomAddress, common.HexToAddress("0x0"), EthToWei(1))
                Expect(err).ToNot(HaveOccurred())

                nonce := big.NewInt(1)
                signature, err := SignData(nonce, data, privateKey)
                Expect(err).ToNot(HaveOccurred())

                tx, err := Wallet.ExecuteRelayedTransaction(Controller.TransactOpts(), nonce, data, signature)
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

                nonce := big.NewInt(0)
                signature, err := SignData(nonce, data, privateKey)
                Expect(err).ToNot(HaveOccurred())

                tx, err := Wallet.ExecuteRelayedTransaction(Controller.TransactOpts(ethertest.WithGasLimit(500000)), nonce, data, signature)
                Expect(err).ToNot(HaveOccurred())
                Backend.Commit()
                Expect(isSuccessful(tx)).To(BeFalse())

                msg, _ := tx.AsMessage(types.HomesteadSigner{})

                calMsg := ethereum.CallMsg {
                    From: msg.From(),
                    To: msg.To(),
                    Gas: msg.Gas(),
                    GasPrice: msg.GasPrice(),
                    Value: msg.Value(),
                    Data: msg.Data(),
                }

                r, err := Backend.TransactionReceipt(context.Background(), tx.Hash())
                returnData, _ := Backend.CallContract(context.Background(), calMsg, r.BlockNumber)
                Expect(string(returnData[len(returnData)-64:])).To(ContainSubstring("_to address cannot be set to 0x0"))
            })

        })

    })
})
