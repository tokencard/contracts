package wallet_test

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tokencard/contracts/v2/test/shared"
)

var _ = Describe("Valid Signature", func() {


	When("a random account signs a message", func() {
        var signature []byte
        var hash  []byte
        var hash32  [32]byte
		BeforeEach(func() {
			a, err := abi.JSON(strings.NewReader(WALLET_ABI))
			Expect(err).ToNot(HaveOccurred())
			privateKey, _ := crypto.GenerateKey()
			randomAddress := crypto.PubkeyToAddress(privateKey.PublicKey)
            data, err := a.Pack("transfer", randomAddress, common.HexToAddress("0x0"), EthToWei(1))
            Expect(err).ToNot(HaveOccurred())

            hash = crypto.Keccak256(data)
        	signature, err = SignMsg(data, privateKey)
			Expect(err).ToNot(HaveOccurred())
		})

        It("should revert", func() {
            copy(hash32[:], hash)
            _, err := Wallet.IsValidSignature(nil, hash32, signature)
            Expect(err).To(HaveOccurred())
            Expect(err.Error()).To(ContainSubstring("not by owner"))
        })
	})

    When("the owner signs a message", func() {
        var signature []byte
        var hash  []byte
        var hash32  [32]byte

		BeforeEach(func() {
			a, err := abi.JSON(strings.NewReader(WALLET_ABI))
			Expect(err).ToNot(HaveOccurred())
			privateKey, _ := crypto.GenerateKey()
			randomAddress := crypto.PubkeyToAddress(privateKey.PublicKey)
            data, err := a.Pack("transfer", randomAddress, common.HexToAddress("0x0"), EthToWei(1))
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
	})

})
