package wallet_test

import (
	"context"
	"math"
	"math/big"
    "time"

	"github.com/ethereum/go-ethereum/common"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tokencard/contracts/v2/test/shared"
	"github.com/tokencard/ethertest"
)

var _ = Describe("wallet load eth", func() {

	It("Should return the licence ENS-registered node", func() {
		sa, err := Wallet.LicenceNode(nil)
		Expect(err).ToNot(HaveOccurred())
		Expect(sa).To(Equal([32]byte(LicenceName)))
	})


	When("the contract has no balance", func() {

		BeforeEach(func() {
			tx, err := TokenWhitelist.AddTokens(
				ControllerAdmin.TransactOpts(),
				[]common.Address{common.HexToAddress("0x0")},
				StringsToByte32("ETH"),
				[]*big.Int{DecimalsToMagnitude(big.NewInt(18))},
				[]bool{true},
				[]bool{true},
				big.NewInt(20180913153211),
			)
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeTrue())
		})

		When("no value is sent", func() {
			It("Should revert", func() {
				tx, err := Wallet.LoadTokenCard(Owner.TransactOpts(ethertest.WithGasLimit(1000000)), common.HexToAddress("0x0"), big.NewInt(1000))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isGasExhausted(tx, 1000000)).To(BeFalse())
				Expect(isSuccessful(tx)).To(BeFalse())
				Expect(TestRig.LastExecuted()).To(MatchRegexp(`.*load.value\(_amount\)\(_asset, _amount\);`))
			})
		})

		When("the right value is sent along with the transaction", func() {
			It("Should succeed", func() {
				tx, err := Wallet.LoadTokenCard(Owner.TransactOpts(ethertest.WithValue(big.NewInt(1000))), common.HexToAddress("0x0"), big.NewInt(1000))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})
		})

	})

	When("not called by the owner", func() {

		It("Should revert", func() {
			tx, err := Wallet.LoadTokenCard(RandomAccount.TransactOpts(ethertest.WithGasLimit(100000), ethertest.WithValue(EthToWei(1))), common.HexToAddress("0x0"), EthToWei(1))
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isGasExhausted(tx, 100000)).To(BeFalse())
			Expect(isSuccessful(tx)).To(BeFalse())
		})
	})

	It("the initial balance of the Float contract should be zero", func() {
		b, e := Backend.BalanceAt(context.Background(), CryptoFloatAddress, nil)
		Expect(e).ToNot(HaveOccurred())
		Expect(b.String()).To(Equal("0"))
	})

	It("the initial balance of Holder contract address should be zero", func() {
		b, e := Backend.BalanceAt(context.Background(), TokenHolderAddress, nil)
		Expect(e).ToNot(HaveOccurred())
		Expect(b.String()).To(Equal("0"))
	})

	It("the initial balance of the Wallet should be zero", func() {
		b, e := Backend.BalanceAt(context.Background(), WalletAddress, nil)
		Expect(e).ToNot(HaveOccurred())
		Expect(b.String()).To(Equal("0"))
	})

	When("the wallet has 102 ETH and the daily limit is set to 101 ETH", func() {

        BeforeEach(func() {
            tx, err := Wallet.SubmitDailyLimitUpdate(Owner.TransactOpts(), EthToWei(101))
            Expect(err).ToNot(HaveOccurred())
            Backend.Commit()
            Expect(isSuccessful(tx)).To(BeTrue())
            //increase the timestam by one day so the new (higher) limit takes effect
            Backend.AdjustTime(time.Hour*24 + time.Second)
            Backend.Commit()
        })

		BeforeEach(func() {
			RandomAccount.MustTransfer(Backend, WalletAddress, EthToWei(102))
		})

		When("ETH (0x0) is not in the token whitelist", func() {
			It("Should revert", func() {
				tx, err := Wallet.LoadTokenCard(Owner.TransactOpts(ethertest.WithGasLimit(100000)), common.HexToAddress("0x0"), big.NewInt(1000))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isGasExhausted(tx, 100000)).To(BeFalse())
				Expect(isSuccessful(tx)).To(BeFalse())
				Expect(TestRig.LastExecuted()).To(MatchRegexp(`.*_isTokenLoadable\(_asset\), "token not loadable"\);`))
			})
		})

		When("ETH (0x0) is added to the token whitelist", func() {
			BeforeEach(func() {
				tx, err := TokenWhitelist.AddTokens(
					ControllerAdmin.TransactOpts(),
					[]common.Address{common.HexToAddress("0x0")},
					StringsToByte32("ETH"),
					[]*big.Int{DecimalsToMagnitude(big.NewInt(18))},
					[]bool{true},
					[]bool{true},
					big.NewInt(20180913153211),
				)
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})

			When("a valid amount is loaded (101 ETH) and the stablecoin rate is 1", func() {

				BeforeEach(func() {
					tx, err := TokenWhitelist.UpdateTokenRate(
						ControllerAdmin.TransactOpts(),
						StablecoinAddress,
						EthToWei(1),
						big.NewInt(20180913153211),
					)
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeTrue())
				})

				BeforeEach(func() {
					tx, err := Wallet.LoadTokenCard(Owner.TransactOpts(), common.HexToAddress("0x0"), EthToWei(101))
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeTrue())
				})

				It("Should emit a TransferredToTokenHolder event", func() {
					it, err := Licence.FilterTransferredToTokenHolder(nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(it.Next()).To(BeTrue())
					evt := it.Event
					Expect(it.Next()).To(BeFalse())
					Expect(evt.From).To(Equal(WalletAddress))
					Expect(evt.To).To(Equal(TokenHolderAddress))
					Expect(evt.Asset).To(Equal(common.HexToAddress("0x0"))) //represents ETH
					Expect(evt.Amount.String()).To(Equal(EthToWei(1).String()))
				})

				It("Should emit a TransferredToCryptoFloat event", func() {
					it, err := Licence.FilterTransferredToCryptoFloat(nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(it.Next()).To(BeTrue())
					evt := it.Event
					Expect(it.Next()).To(BeFalse())
					Expect(evt.From).To(Equal(WalletAddress))
					Expect(evt.To).To(Equal(CryptoFloatAddress))
					Expect(evt.Asset).To(Equal(common.HexToAddress("0x0"))) //represents ETH
					Expect(evt.Amount.String()).To(Equal(EthToWei(100).String()))
				})

				It("Should emit a LoadedTokenCard event", func() {
					it, err := Wallet.FilterLoadedTokenCard(nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(it.Next()).To(BeTrue())
					evt := it.Event
					Expect(it.Next()).To(BeFalse())
					Expect(evt.Asset).To(Equal(common.HexToAddress("0x0"))) //represents ETH
					Expect(evt.Amount.String()).To(Equal(EthToWei(101).String()))
				})

				It("should increase the ETH balance of the holder contract address by 1 ETH", func() {
					b, e := Backend.BalanceAt(context.Background(), TokenHolderAddress, nil)
					Expect(e).ToNot(HaveOccurred())
					Expect(b.String()).To(Equal(EthToWei(1).String()))
				})

				It("should increase the ETH balance of the holder contract address by 100 ETH", func() {
					b, e := Backend.BalanceAt(context.Background(), CryptoFloatAddress, nil)
					Expect(e).ToNot(HaveOccurred())
					Expect(b.String()).To(Equal(EthToWei(100).String()))
				})
			}) //101 eth are transferred

			When("the stablecoin rate is 0.001", func() {

				BeforeEach(func() {
					tx, err := TokenWhitelist.UpdateTokenRate(
						ControllerAdmin.TransactOpts(),
						StablecoinAddress,
						big.NewInt(int64(0.001*math.Pow10(18))),
						big.NewInt(20180913153211),
					)
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeTrue())
				})

				When("a bigger amount than daily limit is loaded", func() {

					It("Should revert", func() {
						limPlusOneWei := EthToWei(10)
						limPlusOneWei.Add(limPlusOneWei, big.NewInt(1)) // 10 ETH * 1000 + 1= 10,001 stablecoins
						tx, err := Wallet.LoadTokenCard(Owner.TransactOpts(ethertest.WithGasLimit(100000)), common.HexToAddress("0x0"), limPlusOneWei)
						Expect(err).ToNot(HaveOccurred())
						Backend.Commit()
						Expect(isGasExhausted(tx, 100000)).To(BeFalse())
						Expect(isSuccessful(tx)).To(BeFalse())
						returnData, _ := ethCall(tx)
						Expect(string(returnData[len(returnData)-64:])).To(ContainSubstring("available<amount"))
					})
				}) //more daily Load limit

				When("the precise amount of the daily limit is loaded", func() {

					It("Should return 10000", func() {
						value, err := Wallet.ConvertToStablecoin(nil, common.HexToAddress("0x0"), EthToWei(10))
						Expect(err).ToNot(HaveOccurred())
						finalAmount := EthToWei(10)
						finalAmount.Mul(finalAmount, big.NewInt(1000))
						Expect(value.String()).To(Equal(finalAmount.String()))
					})

				}) //amount = spendLimit

			}) //stablecoin rate is 0.001

		}) //0x0 added to whitelist

	}) //102 eth are sent to the wallet

})
