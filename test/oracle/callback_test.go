package oracle_test

import (
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto/sha3"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tokencard/contracts/test/shared"
	"github.com/tokencard/ethertest"
)

func stringToQueryId(url string) [32]byte {
	var id = [32]byte{}
	sha := sha3.NewKeccak256()
	_, err := sha.Write([]byte(url))
	Expect(err).ToNot(HaveOccurred())

	idSlice := sha.Sum(nil)
	Expect(len(idSlice)).To(Equal(32))

	copy(id[:], idSlice)

	return id
}

var _ = Describe("callback", func() {

	Context("When called by oraclize", func() {

		Context("When a token exists and rates update has been requested", func() {
			BeforeEach(func() {
				tx, err := Oracle.AddTokens(
					Controller.TransactOpts(),
					[]common.Address{common.HexToAddress("0xfe209bdE5CA32fa20E6728A005F26D651FFF5982")},
					StringsToByte32("TKN"),
					[]*big.Int{DecimalsToMagnitude(big.NewInt(18))},
					big.NewInt(20180913153211),
				)
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})
			BeforeEach(func() {
				tx, err := Oracle.UpdateTokenRates(Controller.TransactOpts(ethertest.WithValue(big.NewInt(100000000))))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})

			Context("When a valid API pubKey is used", func() {

				Context("When a valid queryID is passed", func() {

					var id [32]byte
					BeforeEach(func() {
						id = stringToQueryId("https://min-api.cryptocompare.com/data/price?fsym=TKN&tsyms=ETH&sign=true")
					})

					Context("When the proof is valid", func() {
						var tx *types.Transaction
						var err error
						Context("When the result has the expected format", func() {
							BeforeEach(func() {
								proof := common.Hex2Bytes("0041ed930d0cf64c73b82c3a04b958f2d27572c09ef7faacb14f062b2ce63eb78331a885fda74e113383ead579337b7e02cc414a214c3bd210142628087dcf5ded781c0060646174653a205765642c203033204f637420323031382031373a30303a323220474d540a6469676573743a205348412d3235363d36514d48744c664e677576362b63795a6133376d68513962776f394449482f6451672f54715a34467453383d")
								tx, err = Oracle.Callback(OraclizeConnectorOwner.TransactOpts(), id, "{\"ETH\":0.001702}", proof)
								Expect(err).ToNot(HaveOccurred())
								Backend.Commit()
							})
							It("Should succeed", func() {
								Expect(isSuccessful(tx)).To(BeTrue())
							})
							It("Should emit a Verified Proof event", func() {
								it, err := Oracle.FilterVerifiedProof(nil)
								Expect(err).ToNot(HaveOccurred())
								Expect(it.Next()).To(BeTrue())
								evt := it.Event
								Expect(it.Next()).To(BeFalse())
								Expect(evt.PublicKey).To(Equal(common.Hex2Bytes("a0f4f688350018ad1b9785991c0bde5f704b005dc79972b114dbed4a615a983710bfc647ebe5a320daa28771dce6a2d104f5efa2e4a85ba3760b76d46f8571ca")))
								Expect(evt.Result).To(Equal("{\"ETH\":0.001702}"))
							})
							It("Should emit a TokenRateUpdate event", func() {
								it, err := Oracle.FilterUpdatedTokenRate(nil)
								Expect(err).ToNot(HaveOccurred())
								Expect(it.Next()).To(BeTrue())
								evt := it.Event
								Expect(it.Next()).To(BeFalse())
								Expect(evt.Token).To(Equal(common.HexToAddress("0xfe209bdE5CA32fa20E6728A005F26D651FFF5982")))
								Expect(evt.Rate.String()).To(Equal(big.NewInt(int64(0.001702 * math.Pow10(18))).String()))
							})
							It("Should update the token's rate and timestamp ", func() {
								token, err := Oracle.Tokens(nil, common.HexToAddress("0xfe209bdE5CA32fa20E6728A005F26D651FFF5982"))
								Expect(err).ToNot(HaveOccurred())
								Expect(token.Rate.String()).To(Equal(big.NewInt(int64(0.001702 * math.Pow10(18))).String()))
								Expect(token.LastUpdate).NotTo(Equal(big.NewInt(0)))
							})
							It("Should fail when called again with the same (deleted, not valid) queryID", func() {
								proof := common.Hex2Bytes("0041ed930d0cf64c73b82c3a04b958f2d27572c09ef7faacb14f062b2ce63eb78331a885fda74e113383ead579337b7e02cc414a214c3bd210142628087dcf5ded781c0060646174653a205765642c203033204f637420323031382031373a30303a323220474d540a6469676573743a205348412d3235363d36514d48744c664e677576362b63795a6133376d68513962776f394449482f6451672f54715a34467453383d")
								tx, err = Oracle.Callback(OraclizeConnectorOwner.TransactOpts(ethertest.WithGasLimit(100000)), id, "{\"ETH\":0.001702}", proof)
								Expect(err).ToNot(HaveOccurred())
								Backend.Commit()
								Expect(isGasExhausted(tx, 100000)).To(BeFalse())
								Expect(isSuccessful(tx)).To(BeFalse())
							})
						})

						Context("When the result is is misformated", func() {

							BeforeEach(func() {
								//update the public key, needed because we sign our own (misformated) results for the callback
								tx, err := Oracle.UpdateAPIPublicKey(Controller.TransactOpts(), common.Hex2Bytes("717580b4c7577ebe0a7c3c21213ffbfa1221d2c1fe455d4897800d86eb65d91f8fb6c2304a54d89ab5c13a690f03dce25f7d46af90f79908d6be8bcdcdf74c22"))
								Expect(err).ToNot(HaveOccurred())
								Backend.Commit()
								Expect(isSuccessful(tx)).To(BeTrue())
							})

							It("It should Fail", func() {
								//'=' instead of ':'
								proof := common.Hex2Bytes("00418d0fbf90402017c9f5e4a92c4e09a05409c07efe6b26160c0935973ab79452330a44bee3b04fe01a784791a5a5bec02ddd1ef3cd80c08e8e551e034f456f48c21c0060646174653a204672692c203136204e6f7620323031382031363a32323a313820474d540a6469676573743a205348412d3235363d4259334f48496c5474497172744e725a69577a65315763657966496752364c31496b456c395070623651493d")
								tx, err := Oracle.Callback(OraclizeConnectorOwner.TransactOpts(ethertest.WithGasLimit(500000)), id, "{\"ETH\"=0.003637}", proof)
								Expect(err).ToNot(HaveOccurred())
								Backend.Commit()
								Expect(isGasExhausted(tx, 500000)).To(BeFalse())
								Expect(isSuccessful(tx)).To(BeFalse())
								Expect(TestRig.LastExecuted()).To(MatchRegexp(`.*token.rate = parseInt\(parseRate\(_result\), 18\);`))
							})

							It("It should Fail", func() {
								//json not terminated properly, missing '}'
								proof := common.Hex2Bytes("004123ce60d99d27fa384793611661a7e4d061172201b0fac17afb5715da74633180ce9d6ac6c9a01df16c86f5bba227fbb15336045ca7efff4d85abdc382aceb8731c0060646174653a204672692c203136204e6f7620323031382031363a32323a313820474d540a6469676573743a205348412d3235363d435a307339584b44704f66494a54694e6e46696d34695a4275384c546572526a334135412b6d6a416b74733d")
								tx, err := Oracle.Callback(OraclizeConnectorOwner.TransactOpts(ethertest.WithGasLimit(500000)), id, "{\"ETH\"=0.003637 mpla mpla mpla mpla", proof)
								Expect(err).ToNot(HaveOccurred())
								Backend.Commit()
								Expect(isGasExhausted(tx, 500000)).To(BeFalse())
								Expect(isSuccessful(tx)).To(BeFalse())
								Expect(TestRig.LastExecuted()).To(MatchRegexp(`.*token.rate = parseInt\(parseRate\(_result\), 18\);`))
							})

							It("It should Fail", func() {
								//no matchin prefix i.e. {\"ETH\":
								proof := common.Hex2Bytes("004132d410111b67765eddfb34e015bf070c27b1059ce92193136d476ff4735a6608225d6a9bc9ea1190ba79ba88eb8c58779a6a3a593574a36b7c365adce0dd2bd71c0060646174653a204672692c203136204e6f7620323031382031363a32323a313820474d540a6469676573743a205348412d3235363d34626531566b697051584b2f454b3453747a78706e2b63622b657673457a374c50507579533263737370493d")
								tx, err := Oracle.Callback(OraclizeConnectorOwner.TransactOpts(ethertest.WithGasLimit(500000)), id, "chancellor on brink of second bailout for banks", proof)
								Expect(err).ToNot(HaveOccurred())
								Backend.Commit()
								Expect(isGasExhausted(tx, 500000)).To(BeFalse())
								Expect(isSuccessful(tx)).To(BeFalse())
								Expect(TestRig.LastExecuted()).To(MatchRegexp(`.*token.rate = parseInt\(parseRate\(_result\), 18\);`))
							})

							It("It should Fail", func() {
								//result is toooo long
								proof := common.Hex2Bytes("0041b9e2ae2711880db913d29f75eb424ba67cb1c9194ba215ee025f9639e03fde1610a530a20100730aea106f4f690eaa76c60660bb66a86c8214f0a8b500e3119f1b0060646174653a204672692c203136204e6f7620323031382031363a32323a313820474d540a6469676573743a205348412d3235363d56306678694339436e4a75566155563632416c554f4d52664c594c6950534179546e7958716f72696f46673d")
								tx, err := Oracle.Callback(OraclizeConnectorOwner.TransactOpts(ethertest.WithGasLimit(500000)), id, "{\"ETH\"=this result is too long and it is going to inject malicious code}", proof)
								Expect(err).ToNot(HaveOccurred())
								Backend.Commit()
								Expect(isGasExhausted(tx, 500000)).To(BeFalse())
								Expect(isSuccessful(tx)).To(BeFalse())
								Expect(TestRig.LastExecuted()).To(MatchRegexp(`.*token.rate = parseInt\(parseRate\(_result\), 18\);`))
							})

						})

					}) //valid proof

					Context("When the proof is not valid", func() {

						Context("When the date is invalid", func() {
							var timestamp *big.Int
							BeforeEach(func() {
								token, err := Oracle.Tokens(nil, common.HexToAddress("0xfe209bdE5CA32fa20E6728A005F26D651FFF5982"))
								timestamp = token.LastUpdate
								Expect(err).ToNot(HaveOccurred())
								Expect(token.Rate.String()).To(Equal(big.NewInt(0).String()))
								Expect(token.LastUpdate).NotTo(Equal(big.NewInt(0)))
							})
							BeforeEach(func() {
								//year set to Dec,2000
								proof := common.Hex2Bytes("0041ed930d0cf64c73b82c3a04b958f2d27572c09ef7faacb14f062b2ce63eb78331a885fda74e113383ead579337b7e02cc414a214c3bd210142628087dcf5ded781c0060646174653a205765642c2030332044656320323030302031373a30303a323220474d540a6469676573743a205348412d3235363d36514d48744c664e677576362b63795a6133376d68513962776f394449482f6451672f54715a34467454393d")
								tx, err := Oracle.Callback(OraclizeConnectorOwner.TransactOpts(), id, "{\"ETH\":0.001702}", proof)
								Expect(err).ToNot(HaveOccurred())
								Backend.Commit()
								Expect(isSuccessful(tx)).To(BeTrue())
							})
							It("Should NOT update the token's rate and timestamp ", func() {
								token, err := Oracle.Tokens(nil, common.HexToAddress("0xfe209bdE5CA32fa20E6728A005F26D651FFF5982"))
								Expect(err).ToNot(HaveOccurred())
								Expect(token.Rate.String()).To(Equal(big.NewInt(0).String()))
								Expect(token.LastUpdate.String()).To(Equal(timestamp.String()))
							})
							It("Should emit a Failed Proof Verification (date) event", func() {
								it, err := Oracle.FilterFailedProofVerification(nil)
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
								tx, err := Oracle.Callback(OraclizeConnectorOwner.TransactOpts(), id, "{\"ETH\":0.001702}", proof)
								Expect(err).ToNot(HaveOccurred())
								Backend.Commit()
								Expect(isSuccessful(tx)).To(BeTrue())
							})
							It("Should emit a Failed Proof Verification event", func() {
								it, err := Oracle.FilterFailedProofVerification(nil)
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
								tx, err := Oracle.Callback(OraclizeConnectorOwner.TransactOpts(), id, "{\"ETH\":0.001702}", proof)
								Expect(err).ToNot(HaveOccurred())
								Backend.Commit()
								Expect(isSuccessful(tx)).To(BeTrue())
							})
							It("Should emit a Failed Proof Verification event", func() {
								it, err := Oracle.FilterFailedProofVerification(nil)
								Expect(err).ToNot(HaveOccurred())
								Expect(it.Next()).To(BeTrue())
								evt := it.Event
								Expect(it.Next()).To(BeFalse())
								Expect(evt.PublicKey).To(Equal(common.Hex2Bytes("a0f4f688350018ad1b9785991c0bde5f704b005dc79972b114dbed4a615a983710bfc647ebe5a320daa28771dce6a2d104f5efa2e4a85ba3760b76d46f8571ca")))
								Expect(evt.Result).To(Equal("{\"ETH\":0.001702}"))
								Expect(evt.Reason).To(Equal("signature"))
							})
						})
						Context("When the proof is empty", func() {
							BeforeEach(func() {
								proof := []byte{}
								tx, err := Oracle.Callback(OraclizeConnectorOwner.TransactOpts(), id, "{\"ETH\":0.001702}", proof)
								Expect(err).ToNot(HaveOccurred())
								Backend.Commit()
								Expect(isSuccessful(tx)).To(BeTrue())
							})
							It("Should emit a Failed Proof Verification event", func() {
								it, err := Oracle.FilterFailedProofVerification(nil)
								Expect(err).ToNot(HaveOccurred())
								Expect(it.Next()).To(BeTrue())
								evt := it.Event
								Expect(it.Next()).To(BeFalse())
								Expect(evt.PublicKey).To(Equal(common.Hex2Bytes("a0f4f688350018ad1b9785991c0bde5f704b005dc79972b114dbed4a615a983710bfc647ebe5a320daa28771dce6a2d104f5efa2e4a85ba3760b76d46f8571ca")))
								Expect(evt.Result).To(Equal("{\"ETH\":0.001702}"))
								Expect(evt.Reason).To(Equal("emptyProof"))
							})
						})

						Context("When the proof is cropped (signature)", func() {
							BeforeEach(func() {
								//last byte of sig is missing, sig[64] or proof[66] (0bytes + 1 byte for length)
								proof := common.Hex2Bytes("0041ed930d0cf64c73b82c3a04b958f2d27572c09ef7faacb14f062b2ce63eb78331a885fda74e113383ead579337b7e02cc414a214c3bd210142628087dcf5ded78")
								tx, err := Oracle.Callback(OraclizeConnectorOwner.TransactOpts(), id, "{\"ETH\":0.001702}", proof)
								Expect(err).ToNot(HaveOccurred())
								Backend.Commit()
								Expect(isSuccessful(tx)).To(BeTrue())
							})
							It("Should emit a Failed Proof Verification event", func() {
								it, err := Oracle.FilterFailedProofVerification(nil)
								Expect(err).ToNot(HaveOccurred())
								Expect(it.Next()).To(BeTrue())
								evt := it.Event
								Expect(it.Next()).To(BeFalse())
								Expect(evt.PublicKey).To(Equal(common.Hex2Bytes("a0f4f688350018ad1b9785991c0bde5f704b005dc79972b114dbed4a615a983710bfc647ebe5a320daa28771dce6a2d104f5efa2e4a85ba3760b76d46f8571ca")))
								Expect(evt.Result).To(Equal("{\"ETH\":0.001702}"))
								Expect(evt.Reason).To(Equal("croppedProof:sig"))
							})
						})

						Context("When the proof length is 1!", func() {
							BeforeEach(func() {
								//last byte of sig is missing, sig[64] or proof[66] (0bytes + 1 byte for length)
								proof := common.Hex2Bytes("00")
								tx, err := Oracle.Callback(OraclizeConnectorOwner.TransactOpts(), id, "{\"ETH\":0.001702}", proof)
								Expect(err).ToNot(HaveOccurred())
								Backend.Commit()
								Expect(isSuccessful(tx)).To(BeTrue())
							})
							It("Should emit a Failed Proof Verification event", func() {
								it, err := Oracle.FilterFailedProofVerification(nil)
								Expect(err).ToNot(HaveOccurred())
								Expect(it.Next()).To(BeTrue())
								evt := it.Event
								Expect(it.Next()).To(BeFalse())
								Expect(evt.PublicKey).To(Equal(common.Hex2Bytes("a0f4f688350018ad1b9785991c0bde5f704b005dc79972b114dbed4a615a983710bfc647ebe5a320daa28771dce6a2d104f5efa2e4a85ba3760b76d46f8571ca")))
								Expect(evt.Result).To(Equal("{\"ETH\":0.001702}"))
								Expect(evt.Reason).To(Equal("croppedProof:sig"))
							})
						})

						Context("When the proof is cropped (headers length)", func() {
							BeforeEach(func() {
								//proof length = 2 + sigLength + 1
								proof := common.Hex2Bytes("0041ed930d0cf64c73b82c3a04b958f2d27572c09ef7faacb14f062b2ce63eb78331a885fda74e113383ead579337b7e02cc414a214c3bd210142628087dcf5ded780000")
								tx, err := Oracle.Callback(OraclizeConnectorOwner.TransactOpts(), id, "{\"ETH\":0.001702}", proof)
								Expect(err).ToNot(HaveOccurred())
								Backend.Commit()
								Expect(isSuccessful(tx)).To(BeTrue())
							})
							It("Should emit a Failed Proof Verification event", func() {
								it, err := Oracle.FilterFailedProofVerification(nil)
								Expect(err).ToNot(HaveOccurred())
								Expect(it.Next()).To(BeTrue())
								evt := it.Event
								Expect(it.Next()).To(BeFalse())
								Expect(evt.PublicKey).To(Equal(common.Hex2Bytes("a0f4f688350018ad1b9785991c0bde5f704b005dc79972b114dbed4a615a983710bfc647ebe5a320daa28771dce6a2d104f5efa2e4a85ba3760b76d46f8571ca")))
								Expect(evt.Result).To(Equal("{\"ETH\":0.001702}"))
								Expect(evt.Reason).To(Equal("croppedProof:headers"))
							})
						})

						Context("When the proof is cropped (headers)", func() {
							BeforeEach(func() {
								//last byte of proof (last headers byte) is missing
								proof := common.Hex2Bytes("0041ed930d0cf64c73b82c3a04b958f2d27572c09ef7faacb14f062b2ce63eb78331a885fda74e113383ead579337b7e02cc414a214c3bd210142628087dcf5ded781c0060646174653a205765642c203033204f637420323031382031373a30303a323220474d540a6469676573743a205348412d3235363d36514d48744c664e677576362b63795a6133376d68513962776f394449482f6451672f54715a3446745338")
								tx, err := Oracle.Callback(OraclizeConnectorOwner.TransactOpts(), id, "{\"ETH\":0.001702}", proof)
								Expect(err).ToNot(HaveOccurred())
								Backend.Commit()
								Expect(isSuccessful(tx)).To(BeTrue())
							})
							It("Should emit a Failed Proof Verification event", func() {
								it, err := Oracle.FilterFailedProofVerification(nil)
								Expect(err).ToNot(HaveOccurred())
								Expect(it.Next()).To(BeTrue())
								evt := it.Event
								Expect(it.Next()).To(BeFalse())
								Expect(evt.PublicKey).To(Equal(common.Hex2Bytes("a0f4f688350018ad1b9785991c0bde5f704b005dc79972b114dbed4a615a983710bfc647ebe5a320daa28771dce6a2d104f5efa2e4a85ba3760b76d46f8571ca")))
								Expect(evt.Result).To(Equal("{\"ETH\":0.001702}"))
								Expect(evt.Reason).To(Equal("croppedProof:headers"))
							})
						})

					}) //invalid proof
				}) //valid query id

				Context("When an invalid queryID is passed", func() {

					var id [32]byte
					BeforeEach(func() {
						//change pair to EOS/ETH instead of TKN/ETH
						id = stringToQueryId("https://min-api.cryptocompare.com/data/price?fsym=EOS&tsyms=ETH&sign=true")
					})

					Context("When the proof is valid", func() {
						It("Should fail", func() {
							proof := common.Hex2Bytes("0041ed930d0cf64c73b82c3a04b958f2d27572c09ef7faacb14f062b2ce63eb78331a885fda74e113383ead579337b7e02cc414a214c3bd210142628087dcf5ded781c0060646174653a205765642c203033204f637420323031382031373a30303a323220474d540a6469676573743a205348412d3235363d36514d48744c664e677576362b63795a6133376d68513962776f394449482f6451672f54715a34467453383d")
							tx, err := Oracle.Callback(OraclizeConnectorOwner.TransactOpts(ethertest.WithGasLimit(200000)), id, "{\"ETH\":0.001702}", proof)
							Expect(err).ToNot(HaveOccurred())
							Backend.Commit()
							Expect(isGasExhausted(tx, 200000)).To(BeFalse())
							Expect(isSuccessful(tx)).To(BeFalse())
						})
					})

				}) //invalid query id

			}) //valid API key

			Context("When an invalid API pubKey is used", func() {

				var id [32]byte
				BeforeEach(func() {
					id = stringToQueryId("https://min-api.cryptocompare.com/data/price?fsym=TKN&tsyms=ETH&sign=true")
				})
				BeforeEach(func() {
					tx, err := Oracle.UpdateAPIPublicKey(Controller.TransactOpts(), common.Hex2Bytes("fffffff"))
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeTrue())
				})

				It("Should update public key", func() {
					key, err := Oracle.APIPublicKey(nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(key).To(Equal(common.Hex2Bytes("fffffff")))
				})

				Context("When the proof is valid", func() {

					BeforeEach(func() {
						proof := common.Hex2Bytes("0041ed930d0cf64c73b82c3a04b958f2d27572c09ef7faacb14f062b2ce63eb78331a885fda74e113383ead579337b7e02cc414a214c3bd210142628087dcf5ded781c0060646174653a205765642c203033204f637420323031382031373a30303a323220474d540a6469676573743a205348412d3235363d36514d48744c664e677576362b63795a6133376d68513962776f394449482f6451672f54715a34467453383d")
						tx, err := Oracle.Callback(OraclizeConnectorOwner.TransactOpts(ethertest.WithGasLimit(200000)), id, "{\"ETH\":0.001702}", proof)
						Expect(err).ToNot(HaveOccurred())
						Backend.Commit()
						Expect(isGasExhausted(tx, 200000)).To(BeFalse())
						Expect(isSuccessful(tx)).To(BeTrue())
					})

					It("Should emit a Failed Proof Verification event", func() {
						it, err := Oracle.FilterFailedProofVerification(nil)
						Expect(err).ToNot(HaveOccurred())
						Expect(it.Next()).To(BeTrue())
						evt := it.Event
						Expect(it.Next()).To(BeFalse())
						Expect(evt.PublicKey).To(Equal(common.Hex2Bytes("fffffff")))
						Expect(evt.Result).To(Equal("{\"ETH\":0.001702}"))
						Expect(evt.Reason).To(Equal("signature"))
					})

				})

			}) //invalid API key

		}) //update rates requested

		Context("When a token exists but rates update has NOT been requested", func() {

			BeforeEach(func() {
				tx, err := Oracle.AddTokens(
					Controller.TransactOpts(),
					[]common.Address{common.HexToAddress("0xfe209bdE5CA32fa20E6728A005F26D651FFF5982")},
					StringsToByte32("TKN"),
					[]*big.Int{DecimalsToMagnitude(big.NewInt(18))},
					big.NewInt(20180913153211),
				)
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})

			It("Should fail", func() {
				proof := common.Hex2Bytes("0041ed930d0cf64c73b82c3a04b958f2d27572c09ef7faacb14f062b2ce63eb78331a885fda74e113383ead579337b7e02cc414a214c3bd210142628087dcf5ded781c0060646174653a205765642c203033204f637420323031382031373a30303a323220474d540a6469676573743a205348412d3235363d36514d48744c664e677576362b63795a6133376d68513962776f394449482f6451672f54715a34467453383d")
				id := stringToQueryId("https://min-api.cryptocompare.com/data/price?fsym=TKN&tsyms=ETH&sign=true")
				tx, err := Oracle.Callback(OraclizeConnectorOwner.TransactOpts(ethertest.WithGasLimit(200000)), id, "{\"ETH\":0.001702}", proof)
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isGasExhausted(tx, 200000)).To(BeFalse())
				Expect(isSuccessful(tx)).To(BeFalse())
			})
		})

	})

	Context("When not called by Oraclize", func() {

		Context("When then token does not exist", func() {
			It("Should fail", func() {
				tx, err := Oracle.Callback(RandomAccount.TransactOpts(ethertest.WithGasLimit(100000)), [32]byte{}, "", []byte{})
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isGasExhausted(tx, 100000)).To(BeFalse())
				Expect(isSuccessful(tx)).To(BeFalse())
			})
		})

		Context("When a token exists and rates update has been requested", func() {
			BeforeEach(func() {
				tx, err := Oracle.AddTokens(
					Controller.TransactOpts(),
					[]common.Address{common.HexToAddress("0xfe209bdE5CA32fa20E6728A005F26D651FFF5982")},
					StringsToByte32("TKN"),
					[]*big.Int{DecimalsToMagnitude(big.NewInt(8))},
					big.NewInt(20180913153211),
				)
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})
			BeforeEach(func() {
				tx, err := Oracle.UpdateTokenRates(Controller.TransactOpts(ethertest.WithValue(big.NewInt(100000000))))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})
			It("Should fail", func() {
				proof := common.Hex2Bytes("0041ed930d0cf64c73b82c3a04b958f2d27572c09ef7faacb14f062b2ce63eb78331a885fda74e113383ead579337b7e02cc414a214c3bd210142628087dcf5ded781c0060646174653a205765642c203033204f637420323031382031373a30303a323220474d540a6469676573743a205348412d3235363d36514d48744c664e677576362b63795a6133376d68513962776f394449482f6451672f54715a34467453383d")
				id := stringToQueryId("https://min-api.cryptocompare.com/data/price?fsym=TKN&tsyms=ETH&sign=true")
				tx, err := Oracle.Callback(Controller.TransactOpts(ethertest.WithGasLimit(200000)), id, "{\"ETH\":0.001702}", proof)
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isGasExhausted(tx, 200000)).To(BeFalse())
				Expect(isSuccessful(tx)).To(BeFalse())
			})
		})

	}) //random address

})
