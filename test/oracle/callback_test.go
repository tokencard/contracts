package oracle_test

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto/sha3"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tokencard/ethertest"
)

var _ = Describe("callback", func() {
	Context("When called by oraclize", func() {
		Context("When a token exists and rates update has been requested", func() {
			BeforeEach(func() {
				tx, err := oracle.AddTokens(controller.TransactOpts(), []common.Address{common.HexToAddress("0xfe209bdE5CA32fa20E6728A005F26D651FFF5982")}, stringsToByte32("TKN"), []*big.Int{exponentiateDecimals(18)})
				Expect(err).ToNot(HaveOccurred())
				be.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})
			BeforeEach(func() {
				tx, err := oracle.UpdateTokenRates(controller.TransactOpts(WithValue(big.NewInt(100000000))))
				Expect(err).ToNot(HaveOccurred())
				be.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})

			Context("When a valid API pubKey is used", func() {
				Context("When a valid queryID is passed", func() {
					var id = [32]byte{}
					BeforeEach(func() {
						sha := sha3.NewKeccak256()
						_, err := sha.Write([]byte("https://min-api.cryptocompare.com/data/price?fsym=TKN&tsyms=ETH&sign=true"))
						Expect(err).ToNot(HaveOccurred())

						idSlice := sha.Sum(nil)
						Expect(len(idSlice)).To(Equal(32))

						copy(id[:], idSlice)
					})
					Context("When the proof is valid", func() {
						// var tx *types.Transaction
						// var err error
						//
						// BeforeEach(func() {
						// 	proof := common.Hex2Bytes("0041ed930d0cf64c73b82c3a04b958f2d27572c09ef7faacb14f062b2ce63eb78331a885fda74e113383ead579337b7e02cc414a214c3bd210142628087dcf5ded781c0060646174653a205765642c203033204f637420323031382031373a30303a323220474d540a6469676573743a205348412d3235363d36514d48744c664e677576362b63795a6133376d68513962776f394449482f6451672f54715a34467453383d")
						// 	tx, err = oracle.Callback(oraclizeCallbackAccount.TransactOpts(), id, "{\"ETH\":0.001702}", proof)
						// 	Expect(err).ToNot(HaveOccurred())
						// 	be.Commit()
						// })
						// It("Should succeed", func() {
						// 	Expect(isSuccessful(tx)).To(BeTrue())
						// })
						// It("Should emit a Verified Proof event", func() {
						// 	it, err := oracle.FilterVerifiedProof(nil)
						// 	Expect(err).ToNot(HaveOccurred())
						// 	Expect(it.Next()).To(BeTrue())
						// 	evt := it.Event
						// 	Expect(it.Next()).To(BeFalse())
						// 	Expect(evt.PublicKey).To(Equal(common.Hex2Bytes("a0f4f688350018ad1b9785991c0bde5f704b005dc79972b114dbed4a615a983710bfc647ebe5a320daa28771dce6a2d104f5efa2e4a85ba3760b76d46f8571ca")))
						// 	Expect(evt.Result).To(Equal("{\"ETH\":0.001702}"))
						// })
						// It("Should emit a TokenRateUpdate event", func() {
						// 	it, err := oracle.FilterUpdatedTokenRate(nil)
						// 	Expect(err).ToNot(HaveOccurred())
						// 	Expect(it.Next()).To(BeTrue())
						// 	evt := it.Event
						// 	Expect(it.Next()).To(BeFalse())
						// 	Expect(evt.Token).To(Equal(common.HexToAddress("0xfe209bdE5CA32fa20E6728A005F26D651FFF5982")))
						// 	Expect(evt.Rate.String()).To(Equal(big.NewInt(int64(0.001702 * math.Pow10(18))).String()))
						// })
						// It("Should update the token's rate and timestamp ", func() {
						// 	token, err := oracle.Tokens(nil, common.HexToAddress("0xfe209bdE5CA32fa20E6728A005F26D651FFF5982"))
						// 	Expect(err).ToNot(HaveOccurred())
						// 	Expect(token.Rate.String()).To(Equal(big.NewInt(int64(0.001702 * math.Pow10(18))).String()))
						// 	Expect(token.LastUpdate).NotTo(Equal(big.NewInt(0)))
						// })
						// It("Should fail when called again with the same (deleted, not valid) queryID", func() {
						// 	proof := common.Hex2Bytes("0041ed930d0cf64c73b82c3a04b958f2d27572c09ef7faacb14f062b2ce63eb78331a885fda74e113383ead579337b7e02cc414a214c3bd210142628087dcf5ded781c0060646174653a205765642c203033204f637420323031382031373a30303a323220474d540a6469676573743a205348412d3235363d36514d48744c664e677576362b63795a6133376d68513962776f394449482f6451672f54715a34467453383d")
						// 	tx, err = oracle.Callback(oraclizeCallbackAccount.TransactOpts(WithGasLimit(100000)), id, "{\"ETH\":0.001702}", proof)
						// 	Expect(err).ToNot(HaveOccurred())
						// 	be.Commit()
						// 	Expect(isGasExhausted(tx, 100000)).To(BeFalse())
						// 	Expect(isSuccessful(tx)).To(BeFalse())
						// })
					}) //valid proof
					Context("When the proof is not valid", func() {
						Context("When the date is invalid", func() {
							var timestamp *big.Int
							BeforeEach(func() {
								token, err := oracle.Tokens(nil, common.HexToAddress("0xfe209bdE5CA32fa20E6728A005F26D651FFF5982"))
								timestamp = token.LastUpdate
								Expect(err).ToNot(HaveOccurred())
								Expect(token.Rate.String()).To(Equal(big.NewInt(0).String()))
								Expect(token.LastUpdate).NotTo(Equal(big.NewInt(0)))
							})
							BeforeEach(func() {
								//year set to 2000
								proof := common.Hex2Bytes("0041ed930d0cf64c73b82c3a04b958f2d27572c09ef7faacb14f062b2ce63eb78331a885fda74e113383ead579337b7e02cc414a214c3bd210142628087dcf5ded781c0060646174653a205765642c203033204f637420323031382031373a30303a323220474d540a6469676573743a205348412d3235363d36514d48744c664e677576362b63795a6133376d68513962776f394449482f6451672f54715a34467453383d")
								tx, err := oracle.Callback(oraclizeCallbackAccount.TransactOpts(), id, "{\"ETH\":0.001702}", proof)
								Expect(err).ToNot(HaveOccurred())
								be.Commit()
								Expect(isSuccessful(tx)).To(BeTrue())
							})
							It("Should NOT update the token's rate and timestamp ", func() {
								token, err := oracle.Tokens(nil, common.HexToAddress("0xfe209bdE5CA32fa20E6728A005F26D651FFF5982"))
								Expect(err).ToNot(HaveOccurred())
								Expect(token.Rate.String()).To(Equal(big.NewInt(0).String()))
								Expect(token.LastUpdate.String()).To(Equal(timestamp.String()))
							})
							It("Should emit a Failed Proof Verification (date) event", func() {
								it, err := oracle.FilterFailedProofVerification(nil)
								Expect(err).ToNot(HaveOccurred())
								Expect(it.Next()).To(BeTrue())
								evt := it.Event
								Expect(it.Next()).To(BeFalse())
								Expect(evt.PublicKey).To(Equal(common.Hex2Bytes("a0f4f688350018ad1b9785991c0bde5f704b005dc79972b114dbed4a615a983710bfc647ebe5a320daa28771dce6a2d104f5efa2e4a85ba3760b76d46f8571ca")))
								Expect(evt.Result).To(Equal("{\"ETH\":0.001702}"))
								Expect(evt.Reason).To(Equal("date"))
							})
						})
						Context("When the result hash does not match", func() {
							BeforeEach(func() {
								//date has been tampered with (year 2019 instead of 2018)
								//result has changed (bytes[-3]-bytes[-2])
								//the date check will pass but the result hash will fail
								proof := common.Hex2Bytes("0041ed930d0cf64c73b82c3a04b958f2d27572c09ef7faacb14f062b2ce63eb78331a885fda74e113383ead579337b7e02cc414a214c3bd210142628087dcf5ded781c0060646174653a205765642c203033204f637420323031392031373a30303a323220474d540a6469676573743a205348412d3235363d36514d48744c664e677576362b63795a6133376d68513962776f394449482f6451672f54715a34467454393d")
								tx, err := oracle.Callback(oraclizeCallbackAccount.TransactOpts(), id, "{\"ETH\":0.001702}", proof)
								Expect(err).ToNot(HaveOccurred())
								be.Commit()
								Expect(isSuccessful(tx)).To(BeTrue())
							})
							It("Should emit a Failed Proof Verification event", func() {
								it, err := oracle.FilterFailedProofVerification(nil)
								Expect(err).ToNot(HaveOccurred())
								Expect(it.Next()).To(BeTrue())
								evt := it.Event
								Expect(it.Next()).To(BeFalse())
								Expect(evt.PublicKey).To(Equal(common.Hex2Bytes("a0f4f688350018ad1b9785991c0bde5f704b005dc79972b114dbed4a615a983710bfc647ebe5a320daa28771dce6a2d104f5efa2e4a85ba3760b76d46f8571ca")))
								Expect(evt.Result).To(Equal("{\"ETH\":0.001702}"))
								Expect(evt.Reason).To(Equal("hash"))
							})
						})

						Context("When the signature is invalid", func() {
							BeforeEach(func() {
								//change the 3rd-10th bytes to 0xff, beginning of the proof
								// proof := common.Hex2Bytes("0041ffffffffffffffff2c3a04b958f2d27572c09ef7faacb14f062b2ce63eb78331a885fda74e113383ead579337b7e02cc414a214c3bd210142628087dcf5ded781c0060646174653a205765642c203033204f637420323031382031373a30303a323220474d540a6469676573743a205348412d3235363d36514d48744c664e677576362b63795a6133376d68513962776f394449482f6451672f54715a34467453383d")
								//date has been tampered with (year 2019 instead of 2018), this invalidates the signature since the whole header is signed
								proof := common.Hex2Bytes("0041ed930d0cf64c73b82c3a04b958f2d27572c09ef7faacb14f062b2ce63eb78331a885fda74e113383ead579337b7e02cc414a214c3bd210142628087dcf5ded781c0060646174653a205765642c203033204f637420323031392031373a30303a323220474d540a6469676573743a205348412d3235363d36514d48744c664e677576362b63795a6133376d68513962776f394449482f6451672f54715a34467453383d")
								tx, err := oracle.Callback(oraclizeCallbackAccount.TransactOpts(), id, "{\"ETH\":0.001702}", proof)
								Expect(err).ToNot(HaveOccurred())
								be.Commit()
								Expect(isSuccessful(tx)).To(BeTrue())
							})
							It("Should emit a Failed Proof Verification event", func() {
								it, err := oracle.FilterFailedProofVerification(nil)
								Expect(err).ToNot(HaveOccurred())
								Expect(it.Next()).To(BeTrue())
								evt := it.Event
								Expect(it.Next()).To(BeFalse())
								Expect(evt.PublicKey).To(Equal(common.Hex2Bytes("a0f4f688350018ad1b9785991c0bde5f704b005dc79972b114dbed4a615a983710bfc647ebe5a320daa28771dce6a2d104f5efa2e4a85ba3760b76d46f8571ca")))
								Expect(evt.Result).To(Equal("{\"ETH\":0.001702}"))
								Expect(evt.Reason).To(Equal("signature"))
							})
						})
					})
				}) //valid query id
			})
		}) //update rates
	})

	Context("When called by a random address", func() {
		Context("When the token exists", func() {
			BeforeEach(func() {
				tx, err := oracle.AddTokens(controller.TransactOpts(), []common.Address{common.HexToAddress("0x1")}, stringsToByte32("ETH"), []*big.Int{exponentiateDecimals(18)})
				Expect(err).ToNot(HaveOccurred())
				be.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})
			It("Should fail", func() {
				tx, err := oracle.Callback(randomAccount.TransactOpts(WithGasLimit(100000)), [32]byte{}, "", []byte{})
				Expect(err).ToNot(HaveOccurred())
				be.Commit()
				Expect(isGasExhausted(tx, 100000)).To(BeFalse())
				Expect(isSuccessful(tx)).To(BeFalse())
			})
		})

		Context("When then token does not exist", func() {
			It("Should fail", func() {
				tx, err := oracle.Callback(randomAccount.TransactOpts(WithGasLimit(100000)), [32]byte{}, "", []byte{})
				Expect(err).ToNot(HaveOccurred())
				be.Commit()
				Expect(isGasExhausted(tx, 100000)).To(BeFalse())
				Expect(isSuccessful(tx)).To(BeFalse())
			})
		})
	})

})
