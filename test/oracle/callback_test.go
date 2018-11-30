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
								Expect(TestRig.LastExecuted()).To(MatchRegexp(`.*token.rate = parseIntRevert\(parseRate\(_result\), 18\);`))
							})

							It("It should Fail", func() {
								//json not terminated properly, missing '}'
								proof := common.Hex2Bytes("004123ce60d99d27fa384793611661a7e4d061172201b0fac17afb5715da74633180ce9d6ac6c9a01df16c86f5bba227fbb15336045ca7efff4d85abdc382aceb8731c0060646174653a204672692c203136204e6f7620323031382031363a32323a313820474d540a6469676573743a205348412d3235363d435a307339584b44704f66494a54694e6e46696d34695a4275384c546572526a334135412b6d6a416b74733d")
								tx, err := Oracle.Callback(OraclizeConnectorOwner.TransactOpts(ethertest.WithGasLimit(500000)), id, "{\"ETH\"=0.003637 mpla mpla mpla mpla", proof)
								Expect(err).ToNot(HaveOccurred())
								Backend.Commit()
								Expect(isGasExhausted(tx, 500000)).To(BeFalse())
								Expect(isSuccessful(tx)).To(BeFalse())
								Expect(TestRig.LastExecuted()).To(MatchRegexp(`.*token.rate = parseIntRevert\(parseRate\(_result\), 18\);`))
							})

							It("It should Fail", func() {
								//no matchin prefix i.e. {\"ETH\":
								proof := common.Hex2Bytes("004132d410111b67765eddfb34e015bf070c27b1059ce92193136d476ff4735a6608225d6a9bc9ea1190ba79ba88eb8c58779a6a3a593574a36b7c365adce0dd2bd71c0060646174653a204672692c203136204e6f7620323031382031363a32323a313820474d540a6469676573743a205348412d3235363d34626531566b697051584b2f454b3453747a78706e2b63622b657673457a374c50507579533263737370493d")
								tx, err := Oracle.Callback(OraclizeConnectorOwner.TransactOpts(ethertest.WithGasLimit(500000)), id, "chancellor on brink of second bailout for banks", proof)
								Expect(err).ToNot(HaveOccurred())
								Backend.Commit()
								Expect(isGasExhausted(tx, 500000)).To(BeFalse())
								Expect(isSuccessful(tx)).To(BeFalse())
								Expect(TestRig.LastExecuted()).To(MatchRegexp(`.*token.rate = parseIntRevert\(parseRate\(_result\), 18\);`))
							})

							It("It should Fail", func() {
								//result is toooo long
								proof := common.Hex2Bytes("0041b9e2ae2711880db913d29f75eb424ba67cb1c9194ba215ee025f9639e03fde1610a530a20100730aea106f4f690eaa76c60660bb66a86c8214f0a8b500e3119f1b0060646174653a204672692c203136204e6f7620323031382031363a32323a313820474d540a6469676573743a205348412d3235363d56306678694339436e4a75566155563632416c554f4d52664c594c6950534179546e7958716f72696f46673d")
								tx, err := Oracle.Callback(OraclizeConnectorOwner.TransactOpts(ethertest.WithGasLimit(500000)), id, "{\"ETH\"=this result is too long and it is going to inject malicious code}", proof)
								Expect(err).ToNot(HaveOccurred())
								Backend.Commit()
								Expect(isGasExhausted(tx, 500000)).To(BeFalse())
								Expect(isSuccessful(tx)).To(BeFalse())
								Expect(TestRig.LastExecuted()).To(MatchRegexp(`.*token.rate = parseIntRevert\(parseRate\(_result\), 18\);`))
							})
						})

							Context("When the date is is misformated", func() {

								BeforeEach(func() {
									//update the public key, needed because we sign our own (misformated) results for the callback
									tx, err := Oracle.UpdateAPIPublicKey(Controller.TransactOpts(), common.Hex2Bytes("717580b4c7577ebe0a7c3c21213ffbfa1221d2c1fe455d4897800d86eb65d91f8fb6c2304a54d89ab5c13a690f03dce25f7d46af90f79908d6be8bcdcdf74c22"))
									Expect(err).ToNot(HaveOccurred())
									Backend.Commit()
									Expect(isSuccessful(tx)).To(BeTrue())
								})

								It("It should Fail", func() {
									//day set to 99, should revert: require(day > 0 && day < 32, "day error");
									proof := common.Hex2Bytes("0041150f3732c701235dbf7a0abf0f1f57c8e6901a9b987ea13c870678cd4477bc010c16c4dadf1c4a1bfaadc26157f596fed8e2617d0a2953fa37751820502b34011c0060646174653a204672692c203939204e6f7620323031382031363a32323a313020474d540a6469676573743a205348412d3235363d4459452b675a6c4147756c5630562f67774a4347452f78423171484b66516c42476a37586c3441496649383d")
									tx, err := Oracle.Callback(OraclizeConnectorOwner.TransactOpts(ethertest.WithGasLimit(300000)), id, "{\"ETH\":0.003637}", proof)
									Expect(err).ToNot(HaveOccurred())
									Backend.Commit()
									Expect(isGasExhausted(tx, 300000)).To(BeFalse())
									Expect(isSuccessful(tx)).To(BeFalse())
									Expect(TestRig.LastExecuted()).To(MatchRegexp(`.*require\(day > 0 && day < 32, "day error"\);`))
								})

								It("It should Fail", func() {
									//month set to Mov, should revert: require(month > 0 && month < 13, "month error");
									proof := common.Hex2Bytes("0041a963330e85d2dc73bd62c416399598da7230ce38f43801d77e46d373ad48bb4393fbba5ae1783c946e8288d358c4b94a51315c4ad3598d29e7868ed2f3984df71b0060646174653a204672692c203136204d6f7620323031382031363a32323a313020474d540a6469676573743a205348412d3235363d4459452b675a6c4147756c5630562f67774a4347452f78423171484b66516c42476a37586c3441496649383d")
									tx, err := Oracle.Callback(OraclizeConnectorOwner.TransactOpts(ethertest.WithGasLimit(300000)), id, "{\"ETH\":0.003637}", proof)
									Expect(err).ToNot(HaveOccurred())
									Backend.Commit()
									Expect(isGasExhausted(tx, 300000)).To(BeFalse())
									Expect(isSuccessful(tx)).To(BeFalse())
									Expect(TestRig.LastExecuted()).To(MatchRegexp(`.*monthToNumber.*`))
								})

								It("It should Fail", func() {
									//year set to 2017, should revert: require(year > 2017 && year < 3000, "year error");
									proof := common.Hex2Bytes("004167d7be41dfe1a7dfc59239d1c68f049c3d4ed3d6cc550ff48234009e224ae4463f18d6b1b5bfba929d3f20f754eeb1445653e127e5f46bb7e8602fd41dca58b41b0060646174653a204672692c203136204e6f7620323031372031363a32323a313820474d540a6469676573743a205348412d3235363d4459452b675a6c4147756c5630562f67774a4347452f78423171484b66516c42476a37586c3441496649383d")
									tx, err := Oracle.Callback(OraclizeConnectorOwner.TransactOpts(ethertest.WithGasLimit(300000)), id, "{\"ETH\":0.003637}", proof)
									Expect(err).ToNot(HaveOccurred())
									Backend.Commit()
									Expect(isGasExhausted(tx, 300000)).To(BeFalse())
									Expect(isSuccessful(tx)).To(BeFalse())
									Expect(TestRig.LastExecuted()).To(MatchRegexp(`.*require\(year > 2017 && year < 3000, "year error"\);`))
								})

								It("It should Fail", func() {
									//hour set to bb, should revert: revert(\"not a digit\");
									proof := common.Hex2Bytes("00414795d3cd436da3830e4675ec33b9986beec03a652e8343fe3d66c179cf0d533360fbb9822a566f99b27391da2d298601548b2b5e487fb3f1774a9bb8815474f71c0060646174653a204672692c203136204e6f7620323031382062623a32323a313820474d540a6469676573743a205348412d3235363d4459452b675a6c4147756c5630562f67774a4347452f78423171484b66516c42476a37586c3441496649383d")
									tx, err := Oracle.Callback(OraclizeConnectorOwner.TransactOpts(ethertest.WithGasLimit(300000)), id, "{\"ETH\":0.003637}", proof)
									Expect(err).ToNot(HaveOccurred())
									Backend.Commit()
									Expect(isGasExhausted(tx, 300000)).To(BeFalse())
									Expect(isSuccessful(tx)).To(BeFalse())
									Expect(TestRig.LastExecuted()).To(MatchRegexp(`.*revert\("not a digit"\);`))
								})

								It("It should Fail", func() {
									//hour set to 25, should revert: require(hour < 25, "hour error");
									proof := common.Hex2Bytes("0041e337e637f0dd990ea3c7a14e3b7836c6f71f511d2fa50d3c2322a171dfe94a997caf0ebcaef58eeedc99e121d64c2984f379bba7288846e2a10935e2da3834da1c0060646174653a204672692c203136204e6f7620323031382032353a32323a313820474d540a6469676573743a205348412d3235363d4459452b675a6c4147756c5630562f67774a4347452f78423171484b66516c42476a37586c3441496649383d")
									tx, err := Oracle.Callback(OraclizeConnectorOwner.TransactOpts(ethertest.WithGasLimit(300000)), id, "{\"ETH\":0.003637}", proof)
									Expect(err).ToNot(HaveOccurred())
									Backend.Commit()
									Expect(isGasExhausted(tx, 300000)).To(BeFalse())
									Expect(isSuccessful(tx)).To(BeFalse())
									Expect(TestRig.LastExecuted()).To(MatchRegexp(`.*require\(hour < 25, "hour error"\);`))
								})

								It("It should Fail", func() {
									//minute set to 60, should revert: require(minute < 60, "minute error");
									proof := common.Hex2Bytes("00419c41e305443d35209809205b8f754505390c1f39795d70755ef5e4825b27ffb5127cb6e8bf0b0b3ac083d502d7d29ce98ce2cca18721b027591cf575d7e5bd2a1b0060646174653a204672692c203136204e6f7620323031382031363a36303a313820474d540a6469676573743a205348412d3235363d4459452b675a6c4147756c5630562f67774a4347452f78423171484b66516c42476a37586c3441496649383d")
									tx, err := Oracle.Callback(OraclizeConnectorOwner.TransactOpts(ethertest.WithGasLimit(300000)), id, "{\"ETH\":0.003637}", proof)
									Expect(err).ToNot(HaveOccurred())
									Backend.Commit()
									Expect(isGasExhausted(tx, 300000)).To(BeFalse())
									Expect(isSuccessful(tx)).To(BeFalse())
									Expect(TestRig.LastExecuted()).To(MatchRegexp(`.*require\(minute < 60, "minute error"\);`))
								})

								It("It should Fail", func() {
									//second set to 60, should revert: require(second < 60, "second error");
									proof := common.Hex2Bytes("0041f608b1c041993b85bdc83fa061acf24aa2018b987fa275b3d28c140cc856bfbea0d35559ee8374f7b7f42ee28b6861dc8b345f3af353e48cdcd5cbe847d2918a1c0060646174653a204672692c203136204e6f7620323031382031363a32323a363020474d540a6469676573743a205348412d3235363d4459452b675a6c4147756c5630562f67774a4347452f78423171484b66516c42476a37586c3441496649383d")
									tx, err := Oracle.Callback(OraclizeConnectorOwner.TransactOpts(ethertest.WithGasLimit(300000)), id, "{\"ETH\":0.003637}", proof)
									Expect(err).ToNot(HaveOccurred())
									Backend.Commit()
									Expect(isGasExhausted(tx, 300000)).To(BeFalse())
									Expect(isSuccessful(tx)).To(BeFalse())
									Expect(TestRig.LastExecuted()).To(MatchRegexp(`.*require\(second < 60, "second error"\);`))
								})

								It("It should Fail", func() {
									//wrong delimiters, - instead of ' ', interpretes the whole string as the day
									proof := common.Hex2Bytes("0041f1bcafadb5a8cb52218f0dade45bade9024a2cef4d5aae363e57e41ef765c5b86980d0fb5ff62fec6b529352b92025fbb595f29b9441759b8932483d9f7b40061c0060646174653a204672692c2031362d4e6f762d323031382d31363a32323a363020474d540a6469676573743a205348412d3235363d4459452b675a6c4147756c5630562f67774a4347452f78423171484b66516c42476a37586c3441496649383d")
									tx, err := Oracle.Callback(OraclizeConnectorOwner.TransactOpts(ethertest.WithGasLimit(300000)), id, "{\"ETH\":0.003637}", proof)
									Expect(err).ToNot(HaveOccurred())
									Backend.Commit()
									Expect(isGasExhausted(tx, 300000)).To(BeFalse())
									Expect(isSuccessful(tx)).To(BeFalse())
									Expect(TestRig.LastExecuted()).To(MatchRegexp(`.*revert\("not a digit"\);`))
								})

								It("It should Fail", func() {
									//wrong delimiters, - instead of ' ' after month, month is concatenated with the rest of the string, monthToNumber() reverts
									proof := common.Hex2Bytes("0041cfb9355a630d4541c57b51319340854bde93f06b5095b7387b6e2b323f323d5f44bd6b69256844c76dd65e0989ae854cbd991927e984db220da15666b65d89e91b0060646174653a204672692c203136204e6f762d323031382031363a32323a313020474d540a6469676573743a205348412d3235363d4459452b675a6c4147756c5630562f67774a4347452f78423171484b66516c42476a37586c3441496649383d")
									tx, err := Oracle.Callback(OraclizeConnectorOwner.TransactOpts(ethertest.WithGasLimit(300000)), id, "{\"ETH\":0.003637}", proof)
									Expect(err).ToNot(HaveOccurred())
									Backend.Commit()
									Expect(isGasExhausted(tx, 300000)).To(BeFalse())
									Expect(isSuccessful(tx)).To(BeFalse())
									Expect(TestRig.LastExecuted()).To(MatchRegexp(`.*monthToNumber.*`))
								})

								It("It should Fail", func() {
									//wrong delimiters, - instead of ' ' after year, year is concatenated with the rest of the string, revert("not a digit");
									proof := common.Hex2Bytes("004187a94db9df9f4f4a8e1a07d3993ca247bd62cbd4d5ef1a3f3c072480cae3bfe8202286bb81a006e52c1f07d62952bf389684328a7b684f56d169ee9451fbcda91c0060646174653a204672692c203136204e6f7620323031382d31363a32323a313020474d540a6469676573743a205348412d3235363d4459452b675a6c4147756c5630562f67774a4347452f78423171484b66516c42476a37586c3441496649383d")
									tx, err := Oracle.Callback(OraclizeConnectorOwner.TransactOpts(ethertest.WithGasLimit(300000)), id, "{\"ETH\":0.003637}", proof)
									Expect(err).ToNot(HaveOccurred())
									Backend.Commit()
									Expect(isGasExhausted(tx, 300000)).To(BeFalse())
									Expect(isSuccessful(tx)).To(BeFalse())
									Expect(TestRig.LastExecuted()).To(MatchRegexp(`.*revert\("not a digit"\);`))
								})

								It("It should Fail", func() {
									//wrong delimiters, - instead of ' ' after hour, hour is concatenated with the rest of the string, revert("not a digit");
									proof := common.Hex2Bytes("0041e4eb35999bbf8de5410d7a35bfe39cbf630c0e10ef799e1defa6ad8de31e8a3be6b2d85e2ade74c61cfc0cdc68d9b87b964c33e02256e72fed2266681770721a1c0060646174653a204672692c203136204e6f7620323031382031362d32323a313020474d540a6469676573743a205348412d3235363d4459452b675a6c4147756c5630562f67774a4347452f78423171484b66516c42476a37586c3441496649383d")
									tx, err := Oracle.Callback(OraclizeConnectorOwner.TransactOpts(ethertest.WithGasLimit(300000)), id, "{\"ETH\":0.003637}", proof)
									Expect(err).ToNot(HaveOccurred())
									Backend.Commit()
									Expect(isGasExhausted(tx, 300000)).To(BeFalse())
									Expect(isSuccessful(tx)).To(BeFalse())
									Expect(TestRig.LastExecuted()).To(MatchRegexp(`.*revert\("not a digit"\);`))
								})

								It("It should Fail", func() {
									//wrong delimiters, - instead of ' ' after minute, minute is concatenated with the rest of the string, revert("not a digit");
									proof := common.Hex2Bytes("00412f9fc5e8d0623fd3e65361a076bbed3f1b4ed80c5bcc4ca3b7b7c39e8bb23cfeefcca511be20bfb36c121e2f6cb4cf52824bb6a127d6cd82e1662618148de8491c0060646174653a204672692c203136204e6f7620323031382031363a32322d313020474d540a6469676573743a205348412d3235363d4459452b675a6c4147756c5630562f67774a4347452f78423171484b66516c42476a37586c3441496649383d")
									tx, err := Oracle.Callback(OraclizeConnectorOwner.TransactOpts(ethertest.WithGasLimit(300000)), id, "{\"ETH\":0.003637}", proof)
									Expect(err).ToNot(HaveOccurred())
									Backend.Commit()
									Expect(isGasExhausted(tx, 300000)).To(BeFalse())
									Expect(isSuccessful(tx)).To(BeFalse())
									Expect(TestRig.LastExecuted()).To(MatchRegexp(`.*revert\("not a digit"\);`))
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
								//year set to Dec,2000 (could be set by oraclize), the date in the proof is different
								proof := common.Hex2Bytes("0041ed930d0cf64c73b82c3a04b958f2d27572c09ef7faacb14f062b2ce63eb78331a885fda74e113383ead579337b7e02cc414a214c3bd210142628087dcf5ded781c0060646174653a205765642c2030332044656320323030302031373a30303a323220474d540a6469676573743a205348412d3235363d36514d48744c664e677576362b63795a6133376d68513962776f394449482f6451672f54715a34467454393d")
								tx, err := Oracle.Callback(OraclizeConnectorOwner.TransactOpts(ethertest.WithGasLimit(300000)), id, "{\"ETH\":0.001702}", proof)
								Expect(err).ToNot(HaveOccurred())
								Backend.Commit()
								Expect(isGasExhausted(tx, 300000)).To(BeFalse())
								Expect(isSuccessful(tx)).To(BeFalse())
								Expect(TestRig.LastExecuted()).To(MatchRegexp(`.*revert\("invalid signature"\);`))
							})
							It("Should NOT update the token's rate and timestamp ", func() {
								token, err := Oracle.Tokens(nil, common.HexToAddress("0xfe209bdE5CA32fa20E6728A005F26D651FFF5982"))
								Expect(err).ToNot(HaveOccurred())
								Expect(token.Rate.String()).To(Equal(big.NewInt(0).String()))
								Expect(token.LastUpdate.String()).To(Equal(timestamp.String()))
							})
						})

						Context("When the result hash does not match", func() {
							It("Should fail", func() {
								//date has been tampered with (year 2019 instead of 2018)
								//result has changed (bytes[-3]-bytes[-2])
								//the signature verfication should fail, code does not reach the hash verification
								proof := common.Hex2Bytes("0041ed930d0cf64c73b82c3a04b958f2d27572c09ef7faacb14f062b2ce63eb78331a885fda74e113383ead579337b7e02cc414a214c3bd210142628087dcf5ded781c0060646174653a205765642c203033204f637420323031392031373a30303a323220474d540a6469676573743a205348412d3235363d36514d48744c664e677576362b63795a6133376d68513962776f394449482f6451672f54715a34467454393d")
								tx, err := Oracle.Callback(OraclizeConnectorOwner.TransactOpts(ethertest.WithGasLimit(200000)), id, "{\"ETH\":0.001702}", proof)
								Expect(err).ToNot(HaveOccurred())
								Backend.Commit()
								Expect(isGasExhausted(tx, 200000)).To(BeFalse())
								Expect(isSuccessful(tx)).To(BeFalse())
								Expect(TestRig.LastExecuted()).To(MatchRegexp(`.*revert\("invalid signature"\);`))
							})
						})

						Context("When the signature is invalid", func() {
							It("Should fail", func() {
								//change the 3rd-10th bytes to 0xff, beginning of the proof
								// proof := common.Hex2Bytes("0041ffffffffffffffff2c3a04b958f2d27572c09ef7faacb14f062b2ce63eb78331a885fda74e113383ead579337b7e02cc414a214c3bd210142628087dcf5ded781c0060646174653a205765642c203033204f637420323031382031373a30303a323220474d540a6469676573743a205348412d3235363d36514d48744c664e677576362b63795a6133376d68513962776f394449482f6451672f54715a34467453383d")
								//date has been tampered with (year 2019 instead of 2018), this invalidates the signature since the whole header is signed
								proof := common.Hex2Bytes("0041ed930d0cf64c73b82c3a04b958f2d27572c09ef7faacb14f062b2ce63eb78331a885fda74e113383ead579337b7e02cc414a214c3bd210142628087dcf5ded781c0060646174653a205765642c203033204f637420323031392031373a30303a323220474d540a6469676573743a205348412d3235363d36514d48744c664e677576362b63795a6133376d68513962776f394449482f6451672f54715a34467453383d")
								tx, err := Oracle.Callback(OraclizeConnectorOwner.TransactOpts(ethertest.WithGasLimit(200000)), id, "{\"ETH\":0.001702}", proof)
								Expect(err).ToNot(HaveOccurred())
								Backend.Commit()
								Expect(isGasExhausted(tx, 200000)).To(BeFalse())
								Expect(isSuccessful(tx)).To(BeFalse())
								Expect(TestRig.LastExecuted()).To(MatchRegexp(`.*revert\("invalid signature"\);`))
							})
						})

						Context("When the proof is empty", func() {
							It("Should fail", func() {
								proof := []byte{}
								tx, err := Oracle.Callback(OraclizeConnectorOwner.TransactOpts(ethertest.WithGasLimit(200000)), id, "{\"ETH\":0.001702}", proof)
								Expect(err).ToNot(HaveOccurred())
								Backend.Commit()
								Expect(isGasExhausted(tx, 200000)).To(BeFalse())
								Expect(isSuccessful(tx)).To(BeFalse())
								Expect(TestRig.LastExecuted()).To(MatchRegexp(`.*revert\("invalid proof length"\);`))
							})
							Context("When the signature is valid but the proof is not", func() {
								BeforeEach(func() {
									//update the public key, needed because we create our own proofs
									tx, err := Oracle.UpdateAPIPublicKey(Controller.TransactOpts(), common.Hex2Bytes("717580b4c7577ebe0a7c3c21213ffbfa1221d2c1fe455d4897800d86eb65d91f8fb6c2304a54d89ab5c13a690f03dce25f7d46af90f79908d6be8bcdcdf74c22"))
									Expect(err).ToNot(HaveOccurred())
									Backend.Commit()
									Expect(isSuccessful(tx)).To(BeTrue())
								})
								Context("When the sig length is not 65", func() {
									// the 2nd byte indicating the sig length should be set to 65
									It("It should Fail", func() {
										proof := common.Hex2Bytes("0040cfb9355a630d4541c57b51319340854bde93f06b5095b7387b6e2b323f323d5f44bd6b69256844c76dd65e0989ae854cbd991927e984db220da15666b65d89e91b0060646174653a204672692c203136204e6f762d323031382031363a32323a313020474d540a6469676573743a205348412d3235363d4459452b675a6c4147756c5630562f67774a4347452f78423171484b66516c42476a37586c3441496649383d")
										tx, err := Oracle.Callback(OraclizeConnectorOwner.TransactOpts(ethertest.WithGasLimit(300000)), id, "{\"ETH\":0.003637}", proof)
										Expect(err).ToNot(HaveOccurred())
										Backend.Commit()
										Expect(isGasExhausted(tx, 300000)).To(BeFalse())
										Expect(isSuccessful(tx)).To(BeFalse())
										Expect(TestRig.LastExecuted()).To(MatchRegexp(`.*revert\("invalid signature length"\);`))
									})
								})

								Context("When the headers length is not the expected one", func() {
									// it is set to HEADERS_LEN-1
									It("It should Fail", func() {
										proof := common.Hex2Bytes("0041cfb9355a630d4541c57b51319340854bde93f06b5095b7387b6e2b323f323d5f44bd6b69256844c76dd65e0989ae854cbd991927e984db220da15666b65d89e91b005f646174653a204672692c203136204e6f762d323031382031363a32323a313020474d540a6469676573743a205348412d3235363d4459452b675a6c4147756c5630562f67774a4347452f78423171484b66516c42476a37586c3441496649383d")
										tx, err := Oracle.Callback(OraclizeConnectorOwner.TransactOpts(ethertest.WithGasLimit(300000)), id, "{\"ETH\":0.003637}", proof)
										Expect(err).ToNot(HaveOccurred())
										Backend.Commit()
										Expect(isGasExhausted(tx, 300000)).To(BeFalse())
										Expect(isSuccessful(tx)).To(BeFalse())
										Expect(TestRig.LastExecuted()).To(MatchRegexp(`.*revert\("invalid headers length"\);`))
									})
								})

								Context("When the timestamp is not valid", func() {
									// older timestamp(Jan 2018) than the last updated one
									It("It should Fail", func() {
										proof := common.Hex2Bytes("004187d560a2484b126416445ae2b842a520b54ebab2e5ffdca301fa79fa9d50c0114cf82abd3338245bb2e0982f24563e5b1ec422ccec2e4452cc9c08e908d1635c1b0060646174653a204672692c203136204a616e20323031382031363a32323a313020474d540a6469676573743a205348412d3235363d4459452b675a6c4147756c5630562f67774a4347452f78423171484b66516c42476a37586c3441496649383d")
										tx, err := Oracle.Callback(OraclizeConnectorOwner.TransactOpts(ethertest.WithGasLimit(300000)), id, "{\"ETH\":0.003637}", proof)
										Expect(err).ToNot(HaveOccurred())
										Backend.Commit()
										Expect(isGasExhausted(tx, 300000)).To(BeFalse())
										Expect(isSuccessful(tx)).To(BeFalse())
										Expect(TestRig.LastExecuted()).To(MatchRegexp(`.*revert\("invalid date"\);`))
									})
								})

								Context("When the result hash is not matching", func() {
									// the result returned in the callback is not matching the one included in the proof/signature
									It("It should Fail", func() {
										proof := common.Hex2Bytes("00417b19518dc0850674278dfc47a7bef53fff574ccbcd8792d6e744e8b1db78d9439b812fffd03fcad8f5b3fd4e35c46729d6f4a2c8b8de9ce5cad4227c9f4f98691c0060646174653a204672692c203136204e6f7620323031382031363a32323a313020474d540a6469676573743a205348412d3235363d4459452b675a6c4147756c5630562f67774a4347452f78423171484b66516c42476a37586c3441496649383d")
										tx, err := Oracle.Callback(OraclizeConnectorOwner.TransactOpts(ethertest.WithGasLimit(300000)), id, "{\"ETH\":Moon}", proof)
										Expect(err).ToNot(HaveOccurred())
										Backend.Commit()
										Expect(isGasExhausted(tx, 300000)).To(BeFalse())
										Expect(isSuccessful(tx)).To(BeFalse())
										Expect(TestRig.LastExecuted()).To(MatchRegexp(`.*revert\("result hash not matching"\);`))
									})
								})

							})//When the signature is valid but the proof is not


						})

						Context("When the proof is cropped (signature)", func() {
							It("Should fail", func() {
								//last byte of sig is missing, sig[64] or proof[66] (0bytes + 1 byte for length)
								proof := common.Hex2Bytes("0041ed930d0cf64c73b82c3a04b958f2d27572c09ef7faacb14f062b2ce63eb78331a885fda74e113383ead579337b7e02cc414a214c3bd210142628087dcf5ded78")
								tx, err := Oracle.Callback(OraclizeConnectorOwner.TransactOpts(ethertest.WithGasLimit(200000)), id, "{\"ETH\":0.001702}", proof)
								Expect(err).ToNot(HaveOccurred())
								Backend.Commit()
								Expect(isGasExhausted(tx, 200000)).To(BeFalse())
								Expect(isSuccessful(tx)).To(BeFalse())
								Expect(TestRig.LastExecuted()).To(MatchRegexp(`.*revert\("invalid proof length"\);`))
							})
						})

						Context("When the proof is cropped (headers length)", func() {
							It("Should fail", func() {
								//proof length = 2 + sigLength + 1
								proof := common.Hex2Bytes("0041ed930d0cf64c73b82c3a04b958f2d27572c09ef7faacb14f062b2ce63eb78331a885fda74e113383ead579337b7e02cc414a214c3bd210142628087dcf5ded780000")
								tx, err := Oracle.Callback(OraclizeConnectorOwner.TransactOpts(ethertest.WithGasLimit(200000)), id, "{\"ETH\":0.001702}", proof)
								Expect(err).ToNot(HaveOccurred())
								Backend.Commit()
								Expect(isGasExhausted(tx, 200000)).To(BeFalse())
								Expect(isSuccessful(tx)).To(BeFalse())
								Expect(TestRig.LastExecuted()).To(MatchRegexp(`.*revert\("invalid proof length"\);`))
							})
						})

						Context("When the proof is cropped (headers)", func() {
							It("Should fail", func() {
								//last byte of proof (last headers byte) is missing
								proof := common.Hex2Bytes("0041ed930d0cf64c73b82c3a04b958f2d27572c09ef7faacb14f062b2ce63eb78331a885fda74e113383ead579337b7e02cc414a214c3bd210142628087dcf5ded781c0060646174653a205765642c203033204f637420323031382031373a30303a323220474d540a6469676573743a205348412d3235363d36514d48744c664e677576362b63795a6133376d68513962776f394449482f6451672f54715a3446745338")
								tx, err := Oracle.Callback(OraclizeConnectorOwner.TransactOpts(ethertest.WithGasLimit(200000)), id, "{\"ETH\":0.001702}", proof)
								Expect(err).ToNot(HaveOccurred())
								Backend.Commit()
								Expect(isGasExhausted(tx, 200000)).To(BeFalse())
								Expect(isSuccessful(tx)).To(BeFalse())
								Expect(TestRig.LastExecuted()).To(MatchRegexp(`.*revert\("invalid proof length"\);`))
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

					It("Should fail", func() {
						proof := common.Hex2Bytes("0041ed930d0cf64c73b82c3a04b958f2d27572c09ef7faacb14f062b2ce63eb78331a885fda74e113383ead579337b7e02cc414a214c3bd210142628087dcf5ded781c0060646174653a205765642c203033204f637420323031382031373a30303a323220474d540a6469676573743a205348412d3235363d36514d48744c664e677576362b63795a6133376d68513962776f394449482f6451672f54715a34467453383d")
						tx, err := Oracle.Callback(OraclizeConnectorOwner.TransactOpts(ethertest.WithGasLimit(200000)), id, "{\"ETH\":0.001702}", proof)
						Expect(err).ToNot(HaveOccurred())
						Backend.Commit()
						Expect(isGasExhausted(tx, 200000)).To(BeFalse())
						Expect(isSuccessful(tx)).To(BeFalse())
						Expect(TestRig.LastExecuted()).To(MatchRegexp(`.*revert\("invalid signature"\);`))
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
