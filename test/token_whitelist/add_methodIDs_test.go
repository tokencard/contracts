package token_whitelist_test

import (
    "encoding/binary"
    "github.com/ethereum/go-ethereum/crypto"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tokencard/contracts/test/shared"
	"github.com/tokencard/ethertest"
)

var _ = Describe("method Ids", func() {

    When("the tokenWhitelist is deployed", func() {
        It("it should have 3 methods (transfer, approve, transferFrom) whitelisted", func() {

            methodID := crypto.Keccak256Hash([]byte("transfer(address,uint256)"))
            methodIdUint32 := binary.BigEndian.Uint32(methodID[:4])
            wl, err := TokenWhitelist.MethodIdWhitelist(nil, methodIdUint32)
            Expect(err).ToNot(HaveOccurred())
            Expect(wl).To(BeTrue())

            wl, err = TokenWhitelistableExporter.IsMethodIdWhitelisted(nil, methodIdUint32)
            Expect(err).ToNot(HaveOccurred())
            Expect(wl).To(BeTrue())

            methodID = crypto.Keccak256Hash([]byte("approve(address,uint256)"))
            methodIdUint32 = binary.BigEndian.Uint32(methodID[:4])
            wl, err = TokenWhitelist.MethodIdWhitelist(nil, methodIdUint32)
            Expect(err).ToNot(HaveOccurred())
            Expect(wl).To(BeTrue())

            methodID = crypto.Keccak256Hash([]byte("transferFrom(address,address,uint256)"))
            methodIdUint32 = binary.BigEndian.Uint32(methodID[:4])
            wl, err = TokenWhitelist.MethodIdWhitelist(nil, methodIdUint32)
            Expect(err).ToNot(HaveOccurred())
            Expect(wl).To(BeTrue())
        })
    })

    When("the controller admin adds methods (one already existing)", func() {

        var method1 uint32
        var method2 uint32

        BeforeEach(func() {
            methodID := crypto.Keccak256Hash([]byte("transfer(address,uint256)"))
            method1 = binary.BigEndian.Uint32(methodID[:4])

            methodID = crypto.Keccak256Hash([]byte("increaseApproval(address,uint256)"))
            method2 = binary.BigEndian.Uint32(methodID[:4])

            methods := []uint32{method1, method2}
            tx, err := TokenWhitelist.AddMethodIds(ControllerAdmin.TransactOpts(), methods)
            Expect(err).ToNot(HaveOccurred())
            Backend.Commit()
            Expect(isSuccessful(tx)).To(BeTrue())
        })

        It("should add 1 new method", func() {
            methodID := crypto.Keccak256Hash([]byte("increaseApproval(address,uint256)"))
            methodIdUint32 := binary.BigEndian.Uint32(methodID[:4])
            wl, err := TokenWhitelist.MethodIdWhitelist(nil, methodIdUint32)
            Expect(err).ToNot(HaveOccurred())
            Expect(wl).To(BeTrue())
        })

        It("Should emit an AddedMethodId event", func() {
            it, err := TokenWhitelist.FilterAddedMethodId(nil)
            Expect(err).ToNot(HaveOccurred())
            Expect(it.Next()).To(BeTrue())
            evt := it.Event
            Expect(it.Next()).To(BeFalse())
            Expect(evt.MethodId).To(Equal(method2))
        })

        When("the controller admin removes methods", func() {
            BeforeEach(func() {
                methodID := crypto.Keccak256Hash([]byte("transfer(address,uint256)"))
                method1 = binary.BigEndian.Uint32(methodID[:4])

                methodID = crypto.Keccak256Hash([]byte("increaseApproval(address,uint256)"))
                method2 = binary.BigEndian.Uint32(methodID[:4])

                methods := []uint32{method1, method2}
                tx, err := TokenWhitelist.RemoveMethodIds(ControllerAdmin.TransactOpts(), methods)
                Expect(err).ToNot(HaveOccurred())
                Backend.Commit()
                Expect(isSuccessful(tx)).To(BeTrue())
            })

            It("should remove 2 methods", func() {
                methodID := crypto.Keccak256Hash([]byte("increaseApproval(address,uint256)"))
                methodIdUint32 := binary.BigEndian.Uint32(methodID[:4])
                wl, err := TokenWhitelist.MethodIdWhitelist(nil, methodIdUint32)
                Expect(err).ToNot(HaveOccurred())
                Expect(wl).To(BeFalse())

                methodID = crypto.Keccak256Hash([]byte("transfer(address,uint256)"))
                methodIdUint32 = binary.BigEndian.Uint32(methodID[:4])
                wl, err = TokenWhitelist.MethodIdWhitelist(nil, methodIdUint32)
                Expect(err).ToNot(HaveOccurred())
                Expect(wl).To(BeFalse())
            })

            It("Should emit 2 RemovedMethodId event", func() {
                it, err := TokenWhitelist.FilterRemovedMethodId(nil)
                Expect(err).ToNot(HaveOccurred())
                Expect(it.Next()).To(BeTrue())
                evt := it.Event
                Expect(it.Next()).To(BeTrue())
                Expect(evt.MethodId).To(Equal(method1))
                evt = it.Event
                Expect(it.Next()).To(BeFalse())
                Expect(evt.MethodId).To(Equal(method2))
            })
        })

        When("a non-admin account tries to remove methods", func() {
            It("should fail", func() {
                methodID := crypto.Keccak256Hash([]byte("transfer(address,uint256)"))
                method1 := binary.BigEndian.Uint32(methodID[:4])
                methods := []uint32{method1}
                tx, err := TokenWhitelist.RemoveMethodIds(Controller.TransactOpts(ethertest.WithGasLimit(100000)), methods)
                Expect(err).ToNot(HaveOccurred())
                Backend.Commit()
                Expect(isSuccessful(tx)).To(BeFalse())
            })
        })
    })

    When("a non-admin account tries to add methods", func() {
        It("should fail", func() {
            methodID := crypto.Keccak256Hash([]byte("transfer(address,uint256)"))
            method1 := binary.BigEndian.Uint32(methodID[:4])
            methods := []uint32{method1}
            tx, err := TokenWhitelist.AddMethodIds(Controller.TransactOpts(ethertest.WithGasLimit(100000)), methods)
            Expect(err).ToNot(HaveOccurred())
            Backend.Commit()
            Expect(isSuccessful(tx)).To(BeFalse())
        })
    })

})
