package token_whitelist_test

import (
	"math/big"
    "strings"
    "github.com/ethereum/go-ethereum/accounts/abi"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tokencard/contracts/test/shared"
	// "github.com/tokencard/ethertest"
)

var _ = Describe("get the ERC20 Recipient and Amount", func() {

    When("I try to use a non-whitelisted method on a whitelisted/protected token address", func() {
        It("should fail", func() {
            a, err := abi.JSON(strings.NewReader(ERC20ABI))
            Expect(err).ToNot(HaveOccurred())
            data, err := a.Pack("increaseApproval", RandomAccount.Address(), big.NewInt(300))
            Expect(err).ToNot(HaveOccurred())

            _, _, err = TokenWhitelistableExporter.GetERC20RecipientAndAmount(nil, data)
            Expect(err).To(HaveOccurred())
            Expect(err.Error()).To(ContainSubstring("unsupported method"))
        })
    })

    When("I try to use 'transfer' but data (i.e. value is corrupt) is missing", func() {
    	It("should fail", func() {
    		a, err := abi.JSON(strings.NewReader(ERC20ABI))
    		Expect(err).ToNot(HaveOccurred())
    		data, err := a.Pack("transfer", RandomAccount.Address(), big.NewInt(300))
    		Expect(err).ToNot(HaveOccurred())

            _, _, err = TokenWhitelistableExporter.GetERC20RecipientAndAmount(nil, data[:67])
            Expect(err).To(HaveOccurred())
            Expect(err.Error()).To(ContainSubstring("not enough method-encoding bytes"))
    	})
    })

    When("I try to use 'approve' but data (i.e. value) is missing", func() {
		It("should fail", func() {
			a, err := abi.JSON(strings.NewReader(ERC20ABI))
			Expect(err).ToNot(HaveOccurred())
			data, err := a.Pack("approve", RandomAccount.Address(), big.NewInt(300))
			Expect(err).ToNot(HaveOccurred())

            _, _, err = TokenWhitelistableExporter.GetERC20RecipientAndAmount(nil, data[:36])
            Expect(err).To(HaveOccurred())
            Expect(err.Error()).To(ContainSubstring("not enough method-encoding bytes"))
		})
    })

    When("I try to use 'transferFrom' but data (i.e. value is corrupt) is missing", func() {
		It("should fail", func() {
			a, err := abi.JSON(strings.NewReader(ERC20ABI))
			Expect(err).ToNot(HaveOccurred())
            data, err := a.Pack("transferFrom", RandomAccount.Address(), Owner.Address(), big.NewInt(300))
			Expect(err).ToNot(HaveOccurred())
            //transferFrom needs 100 bytes: 4(methodID) + 32 (from) + 32 (to) + 32 (value)
            _, _, err = TokenWhitelistableExporter.GetERC20RecipientAndAmount(nil, data[:99])
            Expect(err).To(HaveOccurred())
            Expect(err.Error()).To(ContainSubstring("not enough data for transferFrom"))
		})
	})

    When("I transfer 300 tokens to a random account", func() {
        It("should succeed", func() {
            a, err := abi.JSON(strings.NewReader(ERC20ABI))
            Expect(err).ToNot(HaveOccurred())
            data, err := a.Pack("transfer", RandomAccount.Address(), big.NewInt(300))
            Expect(err).ToNot(HaveOccurred())

            adr, amt, err := TokenWhitelistableExporter.GetERC20RecipientAndAmount(nil, data)
            Expect(err).ToNot(HaveOccurred())
            Expect(adr).To(Equal(RandomAccount.Address()))
            Expect(amt).To(Equal(big.NewInt(300)))
        })
    })

    When("I approve 300 tokens to a random account", func() {
        It("should succeed", func() {
            a, err := abi.JSON(strings.NewReader(ERC20ABI))
            Expect(err).ToNot(HaveOccurred())
            data, err := a.Pack("approve", RandomAccount.Address(), big.NewInt(300))
            Expect(err).ToNot(HaveOccurred())

            adr, amt, err := TokenWhitelistableExporter.GetERC20RecipientAndAmount(nil, data)
            Expect(err).ToNot(HaveOccurred())
            Expect(adr).To(Equal(RandomAccount.Address()))
            Expect(amt).To(Equal(big.NewInt(300)))
        })
    })

    When("I transferFrom from owner to a random account 300 tokens to a random account", func() {
        It("should succeed", func() {
            a, err := abi.JSON(strings.NewReader(ERC20ABI))
            Expect(err).ToNot(HaveOccurred())
            data, err := a.Pack("transferFrom", Owner.Address(), RandomAccount.Address(), big.NewInt(300))
            Expect(err).ToNot(HaveOccurred())

            adr, amt, err := TokenWhitelistableExporter.GetERC20RecipientAndAmount(nil, data)
            Expect(err).ToNot(HaveOccurred())
            Expect(adr).To(Equal(RandomAccount.Address()))
            Expect(amt).To(Equal(big.NewInt(300)))
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
