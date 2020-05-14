package token_whitelist_test

import (
	"github.com/ethereum/go-ethereum/common"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tokencard/contracts/v3/test/shared"
)

var _ = Describe("isERC20MethodWhitelisted", func() {

	var methoId [4]byte

	It("ERC20 approve should be whitelisted", func() {
		approve := common.Hex2Bytes("095ea7b3")
		copy(methoId[:], approve[:4])
		mw, err := TokenWhitelist.IsERC20MethodWhitelisted(nil, methoId)
		Expect(err).ToNot(HaveOccurred())
		Expect(mw).To(BeTrue())
	})

	It("ERC20 burn should be whitelisted", func() {
		burn := common.Hex2Bytes("42966c68")
		copy(methoId[:], burn[:4])
		mw, err := TokenWhitelist.IsERC20MethodWhitelisted(nil, methoId)
		Expect(err).ToNot(HaveOccurred())
		Expect(mw).To(BeTrue())
	})

	It("ERC20 transfer should be whitelisted", func() {
		transfer := common.Hex2Bytes("a9059cbb")
		copy(methoId[:], transfer[:4])
		mw, err := TokenWhitelist.IsERC20MethodWhitelisted(nil, methoId)
		Expect(err).ToNot(HaveOccurred())
		Expect(mw).To(BeTrue())
	})

	It("ERC20 transferFrom should be whitelisted", func() {
		transferFrom := common.Hex2Bytes("23b872dd")
		copy(methoId[:], transferFrom[:4])
		mw, err := TokenWhitelist.IsERC20MethodWhitelisted(nil, methoId)
		Expect(err).ToNot(HaveOccurred())
		Expect(mw).To(BeTrue())
	})

	It("All other IDs should not be whitelisted", func() {
		randomId := common.Hex2Bytes("00000000")
		copy(methoId[:], randomId[:4])
		mw, err := TokenWhitelist.IsERC20MethodWhitelisted(nil, methoId)
		Expect(err).ToNot(HaveOccurred())
		Expect(mw).To(BeFalse())

		randomId = common.Hex2Bytes("ffffffff")
		copy(methoId[:], randomId[:4])
		mw, err = TokenWhitelist.IsERC20MethodWhitelisted(nil, methoId)
		Expect(err).ToNot(HaveOccurred())
		Expect(mw).To(BeFalse())

		randomId = common.Hex2Bytes("12fa6cb")
		copy(methoId[:], randomId[:4])
		mw, err = TokenWhitelist.IsERC20MethodWhitelisted(nil, methoId)
		Expect(err).ToNot(HaveOccurred())
		Expect(mw).To(BeFalse())
	})

})
