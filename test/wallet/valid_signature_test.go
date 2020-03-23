package wallet_test

import (
	"strings"

    "github.com/tokencard/contracts/v3/pkg/bindings/mocks"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tokencard/contracts/v3/test/shared"
)

var _ = Describe("Valid Signature", func() {

    var IsValidSignatureExporter *mocks.IsValidSignatureExporter


    When("the exporter is deployed", func() {
        var tx *types.Transaction
        var err error
        BeforeEach(func() {
            _, tx, IsValidSignatureExporter, err = mocks.DeployIsValidSignatureExporter(BankAccount.TransactOpts(), Backend, WalletAddress)
        	Expect(err).ToNot(HaveOccurred())
        	Backend.Commit()
        	Expect(isSuccessful(tx)).To(BeTrue())
        })

    	When("a random account signs a message", func() {
            var data  []byte
            var hash  []byte
            var signature []byte
            var hash32  [32]byte

    		BeforeEach(func() {
    			a, err := abi.JSON(strings.NewReader(WALLET_ABI))
    			Expect(err).ToNot(HaveOccurred())
    			privateKey, _ := crypto.GenerateKey()
    			randomAddress := crypto.PubkeyToAddress(privateKey.PublicKey)
                data, err = a.Pack("transfer", randomAddress, common.HexToAddress("0x0"), EthToWei(1))
                Expect(err).ToNot(HaveOccurred())

                hash = crypto.Keccak256(data)
            	signature, err = SignMsg(data, privateKey)
    			Expect(err).ToNot(HaveOccurred())
    		})

            It("should revert", func() {
                copy(hash32[:], hash)
                _, err := Wallet.IsValidSignature(nil, hash32, signature)
                Expect(err).To(HaveOccurred())
                Expect(err.Error()).To(ContainSubstring("only owner"))
            })

            It("should revert", func() {
                _, err := IsValidSignatureExporter.IsValidSignature(nil, data, signature)
                Expect(err).To(HaveOccurred())
                Expect(err.Error()).To(ContainSubstring("only owner"))
            })
    	})

        When("the owner signs a message", func() {
            var data  []byte
            var hash  []byte
            var signature []byte
            var hash32  [32]byte

    		BeforeEach(func() {
    			a, err := abi.JSON(strings.NewReader(WALLET_ABI))
    			Expect(err).ToNot(HaveOccurred())
    			privateKey, _ := crypto.GenerateKey()
    			randomAddress := crypto.PubkeyToAddress(privateKey.PublicKey)
                data, err = a.Pack("transfer", randomAddress, common.HexToAddress("0x0"), EthToWei(1))
                Expect(err).ToNot(HaveOccurred())

                hash = crypto.Keccak256(data)
            	signature, err = SignMsg(data, Owner.PrivKey())
    			Expect(err).ToNot(HaveOccurred())
    		})

            It("should return 0x1626ba7e (_EIP_1654)", func() {
                copy(hash32[:], hash)
                ivs, err := Wallet.IsValidSignature(nil, hash32, signature)
                Expect(err).ToNot(HaveOccurred())
                Expect(ivs).To(Equal([4]byte{0x16,0x26,0xba,0x7e}))
            })

            It("should revert", func() {
                ivs, err := IsValidSignatureExporter.IsValidSignature(nil, data, signature)
                Expect(err).ToNot(HaveOccurred())
                Expect(ivs).To(Equal([4]byte{0x20,0xc1,0x3b,0x0b}))
            })
    	})
    })

})
