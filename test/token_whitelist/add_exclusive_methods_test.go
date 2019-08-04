package token_whitelist_test

import (
    "encoding/binary"
    "math/big"
	"github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/crypto"
    "github.com/tokencard/ethertest"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tokencard/contracts/v2/test/shared"
)

var _ = Describe("exclusive methods Ids", func() {

    When("TKN is added to the tokenWhitelist", func() {
        BeforeEach(func() {
            tx, err := TokenWhitelist.AddTokens(
                ControllerAdmin.TransactOpts(),
                []common.Address{TKNBurnerAddress},
                StringsToByte32("TKN"),
                []*big.Int{DecimalsToMagnitude(big.NewInt(8))},
                []bool{true},
                []bool{true},
                big.NewInt(20180913153211),
            )
            Expect(err).ToNot(HaveOccurred())
            Backend.Commit()
            Expect(isSuccessful(tx)).To(BeTrue())
        })

        When("the controller admin adds exclusive methods (one is in the general whitelist)", func() {

            var method1 uint32
            var method2 uint32

            BeforeEach(func() {
                methodID := crypto.Keccak256Hash([]byte("transfer(address,uint256)"))
                method1 = binary.BigEndian.Uint32(methodID[:4])

                methodID = crypto.Keccak256Hash([]byte("increaseApproval(address,uint256)"))
                method2 = binary.BigEndian.Uint32(methodID[:4])

                methods := []uint32{method1, method2}
                tx, err := TokenWhitelist.AddExclusiveMethods(ControllerAdmin.TransactOpts(),TKNBurnerAddress, methods)
                Expect(err).ToNot(HaveOccurred())
                Backend.Commit()
                Expect(isSuccessful(tx)).To(BeTrue())
            })

            It("should add 1 new method", func() {
                methodID := crypto.Keccak256Hash([]byte("increaseApproval(address,uint256)"))
                methodIdUint32 := binary.BigEndian.Uint32(methodID[:4])
                wl, err := TokenWhitelist.IsERC20MethodSupported(nil, TKNBurnerAddress, methodIdUint32)
                Expect(err).ToNot(HaveOccurred())
                Expect(wl).To(BeTrue())
            })

            It("Should emit an AddedExclusiveMethod event", func() {
                it, err := TokenWhitelist.FilterAddedExclusiveMethod(nil)
                Expect(err).ToNot(HaveOccurred())
                Expect(it.Next()).To(BeTrue())
                evt := it.Event
                Expect(it.Next()).To(BeFalse())
                Expect(evt.Token).To(Equal(TKNBurnerAddress))
                Expect(evt.MethodId).To(Equal(method2))
            })

            When("the controller admin removes methods", func() {
                BeforeEach(func() {
                    methodID := crypto.Keccak256Hash([]byte("transfer(address,uint256)"))
                    method1 = binary.BigEndian.Uint32(methodID[:4])

                    methodID = crypto.Keccak256Hash([]byte("increaseApproval(address,uint256)"))
                    method2 = binary.BigEndian.Uint32(methodID[:4])

                    methods := []uint32{method1, method2}
                    tx, err := TokenWhitelist.RemoveExclusiveMethods(ControllerAdmin.TransactOpts(), TKNBurnerAddress, methods)
                    Expect(err).ToNot(HaveOccurred())
                    Backend.Commit()
                    Expect(isSuccessful(tx)).To(BeTrue())
                })

                It("should remove 1 exclusive method", func() {
                    methodID := crypto.Keccak256Hash([]byte("increaseApproval(address,uint256)"))
                    methodIdUint32 := binary.BigEndian.Uint32(methodID[:4])
                    wl, err := TokenWhitelist.IsERC20MethodSupported(nil, TKNBurnerAddress, methodIdUint32)
                    Expect(err).ToNot(HaveOccurred())
                    Expect(wl).To(BeFalse())

                    methodID = crypto.Keccak256Hash([]byte("transfer(address,uint256)"))
                    methodIdUint32 = binary.BigEndian.Uint32(methodID[:4])
                    wl, err = TokenWhitelist.IsERC20MethodWhitelisted(nil, methodIdUint32)
                    Expect(err).ToNot(HaveOccurred())
                    Expect(wl).To(BeTrue())
                })

                It("Should emit 1 RemovedExclusiveMethod event", func() {
                    it, err := TokenWhitelist.FilterRemovedExclusiveMethod(nil)
                    Expect(err).ToNot(HaveOccurred())
                    Expect(it.Next()).To(BeTrue())
                    evt := it.Event
                    Expect(it.Next()).To(BeFalse())
                    Expect(evt.Token).To(Equal(TKNBurnerAddress))
                    Expect(evt.MethodId).To(Equal(method2))
                })
            })

            When("a non-admin account tries to remove exclusive methods", func() {
                It("should fail", func() {
                    methodID := crypto.Keccak256Hash([]byte("increaseApproval(address,uint256)"))
                    method1 := binary.BigEndian.Uint32(methodID[:4])
                    methods := []uint32{method1}
                    tx, err := TokenWhitelist.RemoveExclusiveMethods(Controller.TransactOpts(ethertest.WithGasLimit(100000)), TKNBurnerAddress, methods)
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
                tx, err := TokenWhitelist.AddExclusiveMethods(Controller.TransactOpts(ethertest.WithGasLimit(100000)), TKNBurnerAddress, methods)
                Expect(err).ToNot(HaveOccurred())
                Backend.Commit()
                Expect(isSuccessful(tx)).To(BeFalse())
            })
        })
    })

    When("no token exists in the tokenWhitelist", func() {

        It("should fail", func() {
            methodID := crypto.Keccak256Hash([]byte("increaseApproval(address,uint256)"))
            methodIdUint32 := binary.BigEndian.Uint32(methodID[:4])
            _, err := TokenWhitelist.IsERC20MethodSupported(nil, TKNBurnerAddress, methodIdUint32)
            Expect(err).To(HaveOccurred())
            Expect(err.Error()).To(ContainSubstring("non-existing token"))
        })

        When("the controller admin tries to add/remove exclusive methods", func() {

            methodID := crypto.Keccak256Hash([]byte("transfer(address,uint256)"))
            method := binary.BigEndian.Uint32(methodID[:4])
            methods := []uint32{method}

            It("should fail", func() {
                tx, err := TokenWhitelist.AddExclusiveMethods(ControllerAdmin.TransactOpts(ethertest.WithGasLimit(200000)),TKNBurnerAddress, methods)
                Expect(err).ToNot(HaveOccurred())
                Backend.Commit()
                Expect(isSuccessful(tx)).To(BeFalse())
            })

            It("should fail", func() {
                methodID := crypto.Keccak256Hash([]byte("transfer(address,uint256)"))
                method1 := binary.BigEndian.Uint32(methodID[:4])
                methods := []uint32{method1}
                tx, err := TokenWhitelist.RemoveExclusiveMethods(ControllerAdmin.TransactOpts(ethertest.WithGasLimit(200000)),TKNBurnerAddress, methods)
                Expect(err).ToNot(HaveOccurred())
                Backend.Commit()
                Expect(isSuccessful(tx)).To(BeFalse())
            })
        })

    })

})
