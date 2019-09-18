package wallet_test

import (
    "strings"
    "crypto/ecdsa"
    "fmt"
    "errors"
    "context"

	"github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/accounts/abi"
    "github.com/ethereum/go-ethereum/crypto"
    "github.com/tokencard/ethertest"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tokencard/contracts/v2/test/shared"
)


func SignData(data []byte, prv *ecdsa.PrivateKey) ([]byte, error) {
    hash := crypto.Keccak256(data)
    EthMessage := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(hash), hash)
    hash = crypto.Keccak256([]byte(EthMessage))
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

	When("a random acount tries to relay", func() {
        It("should fail", func() {
            a, err := abi.JSON(strings.NewReader(WALLET_ABI))
            Expect(err).ToNot(HaveOccurred())
            privateKey, _ := crypto.GenerateKey()
            randomAddress := crypto.PubkeyToAddress(privateKey.PublicKey)
            data, err := a.Pack("transfer", randomAddress, common.HexToAddress("0x0"), EthToWei(1))
            Expect(err).ToNot(HaveOccurred())

            signature, err := SignData(data, Owner.PrivKey())
            Expect(err).ToNot(HaveOccurred())

            tx, err := Wallet.ExecuteRelayedTransaction(RandomAccount.TransactOpts(ethertest.WithGasLimit(500000)), data, signature)
            Expect(err).ToNot(HaveOccurred())
            Backend.Commit()
            Expect(isSuccessful(tx)).To(BeFalse())
        })
    })


    When("a controller tries to relay", func() {
        var randomAddress common.Address
        BeforeEach(func() {
            a, err := abi.JSON(strings.NewReader(WALLET_ABI))
            Expect(err).ToNot(HaveOccurred())
            privateKey, _ := crypto.GenerateKey()
            randomAddress = crypto.PubkeyToAddress(privateKey.PublicKey)
            data, err := a.Pack("transfer", randomAddress, common.HexToAddress("0x0"), EthToWei(1))
            Expect(err).ToNot(HaveOccurred())

            signature, err := SignData(data, Owner.PrivKey())
            Expect(err).ToNot(HaveOccurred())

            tx, err := Wallet.ExecuteRelayedTransaction(Controller.TransactOpts(), data, signature)
            Expect(err).ToNot(HaveOccurred())
            Backend.Commit()
            Expect(isSuccessful(tx)).To(BeTrue())
        })

        It("should emit an ExecutedRelayedTransaction event", func() {
            it, err := Wallet.FilterExecutedRelayedTransaction(nil)
            Expect(err).ToNot(HaveOccurred())
            Expect(it.Next()).To(BeTrue())
            evt := it.Event
            Expect(it.Next()).To(BeFalse())
            Expect(evt.From).To(Equal(Owner.Address()))
        })

        It("should decrease the wallet's ETH balance ", func() {
			b, err := Backend.BalanceAt(context.Background(), WalletAddress, nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(b.String()).To(Equal("0"))
		})

        // It("should increase the random address' balance ", func() {
		// 	b, err := Backend.BalanceAt(context.Background(), randomAddress, nil)
		// 	Expect(err).ToNot(HaveOccurred())
		// 	Expect(b.String()).To(Equal(EthToWei(1).String()))
		// })
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
    }
]`
