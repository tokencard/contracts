package wallet_test

import (
	"context"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/tokencard/contracts/v3/pkg/bindings"
	"github.com/tokencard/contracts/v3/pkg/bindings/externals/upgradeability"
	. "github.com/tokencard/contracts/v3/test/shared"
	"github.com/tokencard/ethertest"
)

var _ = Describe("executeTransaction", func() {

	Context("when the wallet has enough ETH, the daily limit is 100$ and the rate is 1", func() {
		BeforeEach(func() {
			BankAccount.MustTransfer(Backend, WalletProxyAddress, EthToWei(101))

			tx, err := TokenWhitelist.UpdateTokenRate(ControllerAdmin.TransactOpts(), StablecoinAddress, EthToWei(1), big.NewInt(20180913153211))
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeTrue())

			// Set the limit to 100$
			tx, err = WalletProxy.SubmitDailyLimitUpdate(Owner.TransactOpts(), EthToWei(100))
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeTrue())
		})

		var tx *types.Transaction

		When("I transfer 500 Finney to a random address using 'executeTransaction'", func() {

			var randomAddress common.Address
			var spendLimit *big.Int
			var err error

			When("the destination address is not whitelisted", func() {
				BeforeEach(func() {
					spendLimit, err = WalletProxy.DailyLimitAvailable(nil)
					Expect(err).ToNot(HaveOccurred())
					privateKey, _ := crypto.GenerateKey()
					randomAddress = crypto.PubkeyToAddress(privateKey.PublicKey)
					tx, err = WalletProxy.ExecuteTransaction(Owner.TransactOpts(), randomAddress, FinneyToWei(500), nil)
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeTrue())
				})

				It("should increase random address' balance by the same amount", func() {
					b, e := Backend.BalanceAt(context.Background(), randomAddress, nil)
					Expect(e).ToNot(HaveOccurred())
					Expect(b.String()).To(Equal(FinneyToWei(500).String()))
				})

				It("should reduce the available daily spend balance", func() {
					stableValue, err := WalletProxy.ConvertToStablecoin(nil, common.HexToAddress("0x0"), FinneyToWei(500))
					Expect(err).ToNot(HaveOccurred())
					spendLimit.Sub(spendLimit, stableValue)
					sl, err := WalletProxy.DailyLimitAvailable(nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(sl.String()).To(Equal(spendLimit.String()))
				})

				It("should emit an ExecutedTransaction event", func() {
					it, err := WalletProxy.FilterExecutedTransaction(nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(it.Next()).To(BeTrue())
					evt := it.Event
					Expect(it.Next()).To(BeFalse())
					Expect(evt.Destination).To(Equal(randomAddress))
					Expect(evt.Value.String()).To(Equal(FinneyToWei(500).String()))
					Expect(evt.Data).To(Equal([]uint8{}))
					Expect(evt.ReturnData).To(Equal(common.Hex2Bytes("")))
				})
			})

			When("I send data ('transfer' tx) to an random EOA using 'executeTransaction'", func() {
				BeforeEach(func() {
					a, err := abi.JSON(strings.NewReader(ERC20ABI))
					Expect(err).ToNot(HaveOccurred())
					privateKey, _ := crypto.GenerateKey()
					randomAddress = crypto.PubkeyToAddress(privateKey.PublicKey)
					data, err := a.Pack("transfer", randomAddress, big.NewInt(300))
					Expect(err).ToNot(HaveOccurred())

					tx, err = WalletProxy.ExecuteTransaction(Owner.TransactOpts(), randomAddress, FinneyToWei(500), data)
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeTrue())
				})

				It("should increase random EOA's balance by the same amount", func() {
					b, e := Backend.BalanceAt(context.Background(), randomAddress, nil)
					Expect(e).ToNot(HaveOccurred())
					Expect(b.String()).To(Equal(FinneyToWei(500).String()))
				})

				It("should reduce the available daily spend balance", func() {
					av, err := WalletProxy.DailyLimitAvailable(nil)
					Expect(err).ToNot(HaveOccurred())
					sl, err := WalletProxy.DailyLimitValue(nil)
					Expect(err).ToNot(HaveOccurred())
					sl.Sub(sl, FinneyToWei(500))
					Expect(av.String()).To(Equal(sl.String()))
				})

				It("should emit an ExecutedTransaction event", func() {
					it, err := WalletProxy.FilterExecutedTransaction(nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(it.Next()).To(BeTrue())
					evt := it.Event
					Expect(it.Next()).To(BeFalse())
					Expect(evt.Destination).To(Equal(randomAddress))
					Expect(evt.Value.String()).To(Equal(FinneyToWei(500).String()))
					a, err := abi.JSON(strings.NewReader(ERC20ABI))
					Expect(err).ToNot(HaveOccurred())
					ed, _ := a.Pack("transfer", randomAddress, big.NewInt(300))
					Expect(evt.Data).To(Equal(ed))
					Expect(evt.ReturnData).To(Equal(common.Hex2Bytes("")))
				})
			})

			When("the value sent is more than the daily limit", func() {
				It("should fail", func() {
					privateKey, _ := crypto.GenerateKey()
					randomAddress = crypto.PubkeyToAddress(privateKey.PublicKey)
					tx, err = WalletProxy.ExecuteTransaction(Owner.TransactOpts(ethertest.WithGasLimit(100000)), randomAddress, EthToWei(101), nil)
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeFalse())
				})
			})

			When("the value sent is exactly equal to the daily limit", func() {
				BeforeEach(func() {
					privateKey, _ := crypto.GenerateKey()
					randomAddress = crypto.PubkeyToAddress(privateKey.PublicKey)
					tx, err = WalletProxy.ExecuteTransaction(Owner.TransactOpts(), randomAddress, EthToWei(100), nil)
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeTrue())
				})

				It("should increase random EOA's balance by the same amount", func() {
					b, e := Backend.BalanceAt(context.Background(), randomAddress, nil)
					Expect(e).ToNot(HaveOccurred())
					Expect(b.String()).To(Equal(EthToWei(100).String()))
				})

				It("should decrease the wallet's balance by the same amount", func() {
					b, e := Backend.BalanceAt(context.Background(), WalletProxyAddress, nil)
					Expect(e).ToNot(HaveOccurred())
					Expect(b.String()).To(Equal(EthToWei(1).String()))
				})

				It("should reduce the available daily balance to 0", func() {
					av, err := WalletProxy.DailyLimitAvailable(nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(av.String()).To(Equal("0"))
				})

				It("should emit an ExecutedTransaction event", func() {
					it, err := WalletProxy.FilterExecutedTransaction(nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(it.Next()).To(BeTrue())
					evt := it.Event
					Expect(it.Next()).To(BeFalse())
					Expect(evt.Destination).To(Equal(randomAddress))
					Expect(evt.Value.String()).To(Equal(EthToWei(100).String()))
					Expect(evt.Data).To(Equal([]uint8{}))
					Expect(evt.ReturnData).To(Equal(common.Hex2Bytes("")))
				})
			})

			When("the destination address is whitelisted", func() {

				privateKey, _ := crypto.GenerateKey()
				randomAddress = crypto.PubkeyToAddress(privateKey.PublicKey)

				BeforeEach(func() {
					tx, err := WalletProxy.SetWhitelist(Owner.TransactOpts(), []common.Address{randomAddress})
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeTrue())
				})

				BeforeEach(func() {
					spendLimit, err = WalletProxy.DailyLimitAvailable(nil)
					Expect(err).ToNot(HaveOccurred())
					tx, err = WalletProxy.ExecuteTransaction(Owner.TransactOpts(ethertest.WithGasLimit(100000)), randomAddress, FinneyToWei(500), nil)
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeTrue())
				})

				It("should increase random address' balance by the same amount", func() {
					b, e := Backend.BalanceAt(context.Background(), randomAddress, nil)
					Expect(e).ToNot(HaveOccurred())
					Expect(b.String()).To(Equal(FinneyToWei(500).String()))
				})

				It("should NOT reduce the available daily balance", func() {
					sl, err := WalletProxy.DailyLimitAvailable(nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(sl.String()).To(Equal(spendLimit.String()))
				})

				It("should emit an ExecutedTransaction event", func() {
					it, err := WalletProxy.FilterExecutedTransaction(nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(it.Next()).To(BeTrue())
					evt := it.Event
					Expect(it.Next()).To(BeFalse())
					Expect(evt.Destination).To(Equal(randomAddress))
					Expect(evt.Value.String()).To(Equal(FinneyToWei(500).String()))
					Expect(evt.ReturnData).To(Equal(common.Hex2Bytes("")))
					Expect(evt.Data).To(Equal([]uint8{}))
					Expect(evt.ReturnData).To(Equal(common.Hex2Bytes("")))
				})
			})

		})

		When("I have one thousand TKN burner tokens", func() {

			BeforeEach(func() {
				tx, err := TKNBurner.Mint(BankAccount.TransactOpts(), WalletProxyAddress, big.NewInt(1000))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})

			When("I burn 300 tokens using 'executeTransaction'", func() {
				//first we have to set the Tokenholder because burner token is expecting to interact with Holder
				BeforeEach(func() {
					tx, err := TKNBurner.SetTokenHolder(Owner.TransactOpts(), TokenHolderAddress)
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeTrue())
				})

				BeforeEach(func() {
					a, err := abi.JSON(strings.NewReader(ERC20ABI))
					Expect(err).ToNot(HaveOccurred())
					data, err := a.Pack("burn", big.NewInt(300))
					Expect(err).ToNot(HaveOccurred())

					tx, err = WalletProxy.ExecuteTransaction(Owner.TransactOpts(), TKNBurnerAddress, big.NewInt(0), data)
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeTrue())
				})

				It("should decrease the total supply by 300 (800 remaining)", func() {
					s, err := TKNBurner.TotalSupply(nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(s.String()).To(Equal("700"))
				})

				It("should decrease the TKN balance of the wallet", func() {
					b, err := TKNBurner.BalanceOf(nil, WalletProxyAddress)
					Expect(err).ToNot(HaveOccurred())
					Expect(b.String()).To(Equal("700"))
				})

				It("should reduce the available daily balance", func() {
					av, err := WalletProxy.DailyLimitAvailable(nil)
					Expect(err).ToNot(HaveOccurred())
					eth, err := WalletProxy.ConvertToStablecoin(nil, TKNBurnerAddress, big.NewInt(300))
					Expect(err).ToNot(HaveOccurred())
					av.Sub(av, eth) //subtract converted eth from dailySppendLimit
					Expect(av.String()).To(Equal(av.String()))
				})

				It("should emit an ExecutedTransaction event", func() {
					it, err := WalletProxy.FilterExecutedTransaction(nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(it.Next()).To(BeTrue())
					evt := it.Event
					Expect(it.Next()).To(BeFalse())
					Expect(evt.Destination).To(Equal(TKNBurnerAddress))
					Expect(evt.Value.String()).To(Equal("0"))
					a, err := abi.JSON(strings.NewReader(ERC20ABI))
					Expect(err).ToNot(HaveOccurred())
					ed, _ := a.Pack("burn", big.NewInt(300))
					Expect(evt.Data).To(Equal(ed))
					Expect(evt.ReturnData).To(Equal(common.Hex2Bytes("0000000000000000000000000000000000000000000000000000000000000001")))
				})
			})

			When("I transfer 300 tokens to a random account using 'executeTransaction'", func() {
				BeforeEach(func() {
					a, err := abi.JSON(strings.NewReader(ERC20ABI))
					Expect(err).ToNot(HaveOccurred())
					data, err := a.Pack("transfer", RandomAccount.Address(), big.NewInt(300))
					Expect(err).ToNot(HaveOccurred())

					tx, err = WalletProxy.ExecuteTransaction(Owner.TransactOpts(), TKNBurnerAddress, big.NewInt(0), data)
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeTrue())
				})

				It("should increase the TKN balance of the random account", func() {
					b, err := TKNBurner.BalanceOf(nil, RandomAccount.Address())
					Expect(err).ToNot(HaveOccurred())
					Expect(b.String()).To(Equal("300"))
				})

				It("should decrease the TKN balance of the wallet", func() {
					b, err := TKNBurner.BalanceOf(nil, WalletProxyAddress)
					Expect(err).ToNot(HaveOccurred())
					Expect(b.String()).To(Equal("700"))
				})

				It("should reduce the available daily balance", func() {
					av, err := WalletProxy.DailyLimitAvailable(nil)
					Expect(err).ToNot(HaveOccurred())
					eth, err := WalletProxy.ConvertToStablecoin(nil, TKNBurnerAddress, big.NewInt(300))
					Expect(err).ToNot(HaveOccurred())
					av.Sub(av, eth) //subtract converted eth from dailySppendLimit
					Expect(av.String()).To(Equal(av.String()))
				})

				It("should emit an ExecutedTransaction event", func() {
					it, err := WalletProxy.FilterExecutedTransaction(nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(it.Next()).To(BeTrue())
					evt := it.Event
					Expect(it.Next()).To(BeFalse())
					Expect(evt.Destination).To(Equal(TKNBurnerAddress))
					Expect(evt.Value.String()).To(Equal("0"))
					a, err := abi.JSON(strings.NewReader(ERC20ABI))
					Expect(err).ToNot(HaveOccurred())
					ed, _ := a.Pack("transfer", RandomAccount.Address(), big.NewInt(300))
					Expect(evt.Data).To(Equal(ed))
					Expect(evt.ReturnData).To(Equal(common.Hex2Bytes("0000000000000000000000000000000000000000000000000000000000000001")))
				})
			})

			When("I try to use a non-whitelisted method on a whitelisted/protected token address", func() {
				It("should fail", func() {
					a, err := abi.JSON(strings.NewReader(ERC20ABI))
					Expect(err).ToNot(HaveOccurred())
					data, err := a.Pack("increaseApproval", RandomAccount.Address(), big.NewInt(300))
					Expect(err).ToNot(HaveOccurred())

					tx, err = WalletProxy.ExecuteTransaction(Owner.TransactOpts(ethertest.WithGasLimit(100000)), TKNBurnerAddress, big.NewInt(0), data)
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeFalse())
				})
			})

			When("I try to use 'transfer' but data (i.e. value is corrupt) is missing", func() {
				It("should fail", func() {
					a, err := abi.JSON(strings.NewReader(ERC20ABI))
					Expect(err).ToNot(HaveOccurred())
					data, err := a.Pack("transfer", RandomAccount.Address(), big.NewInt(300))
					Expect(err).ToNot(HaveOccurred())
					tx, err = WalletProxy.ExecuteTransaction(Owner.TransactOpts(ethertest.WithGasLimit(100000)), TKNBurnerAddress, big.NewInt(0), data[:67])
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeFalse())
				})
			})

			When("I try to use 'approve' but data (i.e. value) is missing", func() {
				It("should fail", func() {
					a, err := abi.JSON(strings.NewReader(ERC20ABI))
					Expect(err).ToNot(HaveOccurred())
					data, err := a.Pack("approve", RandomAccount.Address(), big.NewInt(300))
					Expect(err).ToNot(HaveOccurred())
					tx, err = WalletProxy.ExecuteTransaction(Owner.TransactOpts(ethertest.WithGasLimit(100000)), TKNBurnerAddress, big.NewInt(0), data[:36])
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeFalse())
				})
			})

			When("I try to use 'transferFrom' but data (i.e. value is corrupt) is missing", func() {
				It("should fail", func() {
					a, err := abi.JSON(strings.NewReader(ERC20ABI))
					Expect(err).ToNot(HaveOccurred())
					data, err := a.Pack("transferFrom", WalletProxyAddress, Owner.Address(), big.NewInt(300))
					Expect(err).ToNot(HaveOccurred())
					//transferFrom needs 100 bytes: 4(methodID) + 32 (from) + 32 (to) + 32 (value)
					tx, err = WalletProxy.ExecuteTransaction(Owner.TransactOpts(ethertest.WithGasLimit(100000)), TKNBurnerAddress, big.NewInt(0), data[:99])
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeFalse())
				})
			})

			When("the destinations addresses are whitelisted", func() {
				BeforeEach(func() {
					tx, err := WalletProxy.SetWhitelist(Owner.TransactOpts(), []common.Address{RandomAccount.Address(), TKNBurnerAddress})
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeTrue())
				})

				When("I approve 300 tokens to a random account using 'executeTransaction'", func() {
					BeforeEach(func() {
						a, err := abi.JSON(strings.NewReader(ERC20ABI))
						Expect(err).ToNot(HaveOccurred())
						data, err := a.Pack("approve", RandomAccount.Address(), big.NewInt(300))
						Expect(err).ToNot(HaveOccurred())

						tx, err = WalletProxy.ExecuteTransaction(Owner.TransactOpts(), TKNBurnerAddress, big.NewInt(0), data)
						Expect(err).ToNot(HaveOccurred())
						Backend.Commit()
						Expect(isSuccessful(tx)).To(BeTrue())
					})

					It("should not increase TKN balance of the random account", func() {
						b, err := TKNBurner.BalanceOf(nil, RandomAccount.Address())
						Expect(err).ToNot(HaveOccurred())
						Expect(b.String()).To(Equal("0"))
					})

					It("should not decrease TKN balance of the wallet", func() {
						b, err := TKNBurner.BalanceOf(nil, WalletProxyAddress)
						Expect(err).ToNot(HaveOccurred())
						Expect(b.String()).To(Equal("1000"))
					})

					It("should not reduce the available daily balance", func() {
						av, err := WalletProxy.DailyLimitAvailable(nil)
						Expect(err).ToNot(HaveOccurred())
						Expect(av.String()).To(Equal("100000000000000000000"))
					})
				})

				When("I transfer 300 tokens to a random account using 'executeTransaction'", func() {
					BeforeEach(func() {

						a, err := abi.JSON(strings.NewReader(ERC20ABI))
						Expect(err).ToNot(HaveOccurred())
						data, err := a.Pack("transfer", RandomAccount.Address(), big.NewInt(300))
						Expect(err).ToNot(HaveOccurred())

						tx, err = WalletProxy.ExecuteTransaction(Owner.TransactOpts(), TKNBurnerAddress, big.NewInt(0), data)
						Expect(err).ToNot(HaveOccurred())
						Backend.Commit()
						Expect(isSuccessful(tx)).To(BeTrue())
					})

					It("should increase the TKN balance of the random account", func() {
						b, err := TKNBurner.BalanceOf(nil, RandomAccount.Address())
						Expect(err).ToNot(HaveOccurred())
						Expect(b.String()).To(Equal("300"))
					})

					It("should decrease the TKN balance of the wallet", func() {
						b, err := TKNBurner.BalanceOf(nil, WalletProxyAddress)
						Expect(err).ToNot(HaveOccurred())
						Expect(b.String()).To(Equal("700"))
					})

					It("should not reduce the available daily balance", func() {
						av, err := WalletProxy.DailyLimitAvailable(nil)
						Expect(err).ToNot(HaveOccurred())
						Expect(av.String()).To(Equal("100000000000000000000"))
					})
				})

				When("I burn 300 tokens using 'executeTransaction'", func() {
					//first we have to set the Tokenholder because burner token is expecting to interact with Holder
					BeforeEach(func() {
						tx, err := TKNBurner.SetTokenHolder(Owner.TransactOpts(), TokenHolderAddress)
						Expect(err).ToNot(HaveOccurred())
						Backend.Commit()
						Expect(isSuccessful(tx)).To(BeTrue())
					})

					BeforeEach(func() {
						a, err := abi.JSON(strings.NewReader(ERC20ABI))
						Expect(err).ToNot(HaveOccurred())
						data, err := a.Pack("burn", big.NewInt(300))
						Expect(err).ToNot(HaveOccurred())

						tx, err = WalletProxy.ExecuteTransaction(Owner.TransactOpts(), TKNBurnerAddress, big.NewInt(0), data)
						Expect(err).ToNot(HaveOccurred())
						Backend.Commit()
						Expect(isSuccessful(tx)).To(BeTrue())
					})

					It("should decrease the total supply by 300 (800 remaining)", func() {
						s, err := TKNBurner.TotalSupply(nil)
						Expect(err).ToNot(HaveOccurred())
						Expect(s.String()).To(Equal("700"))
					})

					It("should decrease the TKN balance of the wallet", func() {
						b, err := TKNBurner.BalanceOf(nil, WalletProxyAddress)
						Expect(err).ToNot(HaveOccurred())
						Expect(b.String()).To(Equal("700"))
					})

					It("should not reduce the available daily spend balance", func() {
						av, err := WalletProxy.DailyLimitAvailable(nil)
						Expect(err).ToNot(HaveOccurred())
						Expect(av.String()).To(Equal("100000000000000000000"))
					})
				})
			})

			When("I approve 300 tokens to a random wallet using 'executeTransaction'", func() {

				var RandomWalletProxy *bindings.Wallet
				var RandomWalletProxyAddress common.Address

				BeforeEach(func() {
					var tx *types.Transaction
					var err error

					RandomWalletProxyAddress, tx, _, err = upgradeability.DeployUpgradeabilityProxy(BankAccount.TransactOpts(), Backend, WalletImplementationAddress, nil)
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeTrue())

					RandomWalletProxy, err = bindings.NewWallet(RandomWalletProxyAddress, Backend)
					tx, err = RandomWalletProxy.InitializeWallet(BankAccount.TransactOpts(), RandomAccount.Address(), false, ENSRegistryAddress, TokenWhitelistNode, ControllerNode, LicenceNode, big.NewInt(1000))
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeTrue())
				})

				BeforeEach(func() {
					a, err := abi.JSON(strings.NewReader(ERC20ABI))
					Expect(err).ToNot(HaveOccurred())
					data, err := a.Pack("approve", RandomWalletProxyAddress, big.NewInt(300))
					Expect(err).ToNot(HaveOccurred())

					tx, err = WalletProxy.ExecuteTransaction(Owner.TransactOpts(), TKNBurnerAddress, big.NewInt(0), data)
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeTrue())
				})

				It("should not increase the TKN balance of the random wallet", func() {
					b, err := TKNBurner.BalanceOf(nil, RandomWalletProxyAddress)
					Expect(err).ToNot(HaveOccurred())
					Expect(b.String()).To(Equal("0"))
				})

				It("should not decrease the TKN balance of the wallet", func() {
					b, err := TKNBurner.BalanceOf(nil, WalletProxyAddress)
					Expect(err).ToNot(HaveOccurred())
					Expect(b.String()).To(Equal("1000"))
				})

				It("should reduce the available daily balance", func() {
					av, err := WalletProxy.DailyLimitAvailable(nil)
					Expect(err).ToNot(HaveOccurred())
					sl, _ := WalletProxy.DailyLimitValue(nil)
					eth, err := WalletProxy.ConvertToStablecoin(nil, TKNBurnerAddress, big.NewInt(300))
					Expect(err).ToNot(HaveOccurred())
					sl.Sub(sl, eth) //subtract converted eth from dailySppendLimit
					Expect(av.String()).To(Equal(sl.String()))
				})

				When("the approved random wallet 'transfersFrom' all the tokens to a whitelisted destination address using 'executeTransaction'", func() {

					BeforeEach(func() {
						tx, err := RandomWalletProxy.SetWhitelist(RandomAccount.TransactOpts(), []common.Address{Owner.Address()})
						Expect(err).ToNot(HaveOccurred())
						Backend.Commit()
						Expect(isSuccessful(tx)).To(BeTrue())
					})

					BeforeEach(func() {
						a, err := abi.JSON(strings.NewReader(ERC20ABI))
						Expect(err).ToNot(HaveOccurred())
						data, err := a.Pack("transferFrom", WalletProxyAddress, Owner.Address(), big.NewInt(300))
						Expect(err).ToNot(HaveOccurred())

						tx, err = RandomWalletProxy.ExecuteTransaction(RandomAccount.TransactOpts(), TKNBurnerAddress, big.NewInt(0), data)
						Expect(err).ToNot(HaveOccurred())
						Backend.Commit()
						Expect(isSuccessful(tx)).To(BeTrue())
					})

					It("should increase the TKN balance of the whitelisted destination", func() {
						b, err := TKNBurner.BalanceOf(nil, Owner.Address())
						Expect(err).ToNot(HaveOccurred())
						Expect(b.String()).To(Equal("300"))
					})

					It("should decrease the TKN balance of the wallet", func() {
						b, err := TKNBurner.BalanceOf(nil, WalletProxyAddress)
						Expect(err).ToNot(HaveOccurred())
						Expect(b.String()).To(Equal("700"))
					})

					It("should NOT reduce the available daily balance", func() {
						av, err := RandomWalletProxy.DailyLimitAvailable(nil)
						Expect(err).ToNot(HaveOccurred())
						Expect(av.String()).To(Equal(EthToWei(1000).String()))
					})
				})

				When("the approved random wallet 'transfersFrom' all the tokens to itself using 'executeTransaction'", func() {
					BeforeEach(func() {
						a, err := abi.JSON(strings.NewReader(ERC20ABI))
						Expect(err).ToNot(HaveOccurred())
						data, err := a.Pack("transferFrom", WalletProxyAddress, RandomWalletProxyAddress, big.NewInt(300))
						Expect(err).ToNot(HaveOccurred())

						tx, err = RandomWalletProxy.ExecuteTransaction(RandomAccount.TransactOpts(), TKNBurnerAddress, big.NewInt(0), data)
						Expect(err).ToNot(HaveOccurred())
						Backend.Commit()
						Expect(isSuccessful(tx)).To(BeTrue())
					})

					It("should increase the TKN balance of the random wallet", func() {
						b, err := TKNBurner.BalanceOf(nil, RandomWalletProxyAddress)
						Expect(err).ToNot(HaveOccurred())
						Expect(b.String()).To(Equal("300"))
					})

					It("should decrease the TKN balance of the wallet", func() {
						b, err := TKNBurner.BalanceOf(nil, WalletProxyAddress)
						Expect(err).ToNot(HaveOccurred())
						Expect(b.String()).To(Equal("700"))
					})

					It("should reduce the available daily balance", func() {
						av, err := RandomWalletProxy.DailyLimitAvailable(nil)
						Expect(err).ToNot(HaveOccurred())
						sl, _ := RandomWalletProxy.DailyLimitValue(nil)
						eth, err := RandomWalletProxy.ConvertToStablecoin(nil, TKNBurnerAddress, big.NewInt(300))
						Expect(err).ToNot(HaveOccurred())
						sl.Sub(sl, eth) //subtract converted eth from dailySppendLimit
						Expect(av.String()).To(Equal(sl.String()))
					})
				})

				When("the approved random wallet tries to 'transferFrom' more than the approved ammount to itself using 'executeTransaction'", func() {
					It("should fail", func() {
						a, err := abi.JSON(strings.NewReader(ERC20ABI))
						Expect(err).ToNot(HaveOccurred())
						data, err := a.Pack("transferFrom", WalletProxyAddress, RandomWalletProxyAddress, big.NewInt(301))
						Expect(err).ToNot(HaveOccurred())

						tx, err = RandomWalletProxy.ExecuteTransaction(RandomAccount.TransactOpts(ethertest.WithGasLimit(500000)), TKNBurnerAddress, big.NewInt(0), data)
						Expect(err).ToNot(HaveOccurred())
						Backend.Commit()
						Expect(isSuccessful(tx)).To(BeFalse())
					})
				})
			})

		})

		When("I have one thousand random tokens of a non-whitelisted ERC20 conctract", func() {
			BeforeEach(func() {
				var err error
				tx, err = ERC20Contract1.Credit(BankAccount.TransactOpts(), WalletProxyAddress, big.NewInt(1000))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})

			When("I transfer 300 tokens to a random account using 'executeTransaction'", func() {
				BeforeEach(func() {
					a, err := abi.JSON(strings.NewReader(ERC20ABI))
					Expect(err).ToNot(HaveOccurred())
					data, err := a.Pack("transfer", RandomAccount.Address(), big.NewInt(300))
					Expect(err).ToNot(HaveOccurred())

					tx, err = WalletProxy.ExecuteTransaction(Owner.TransactOpts(), ERC20Contract1Address, big.NewInt(0), data)
					Expect(err).ToNot(HaveOccurred())
					Backend.Commit()
					Expect(isSuccessful(tx)).To(BeTrue())
				})

				It("should increase ERC20 balance of the random account", func() {
					b, err := ERC20Contract1.BalanceOf(nil, RandomAccount.Address())
					Expect(err).ToNot(HaveOccurred())
					Expect(b.String()).To(Equal("300"))
				})

				It("should decrease the ERC20 balance of the wallet", func() {
					b, err := ERC20Contract1.BalanceOf(nil, WalletProxyAddress)
					Expect(err).ToNot(HaveOccurred())
					Expect(b.String()).To(Equal("700"))
				})

				It("should NOT reduce the available daily spend balance", func() {
					av, err := WalletProxy.DailyLimitAvailable(nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(av.String()).To(Equal(EthToWei(100).String()))
				})

				It("should emit an ExecutedTransaction event", func() {
					it, err := WalletProxy.FilterExecutedTransaction(nil)
					Expect(err).ToNot(HaveOccurred())
					Expect(it.Next()).To(BeTrue())
					evt := it.Event
					Expect(it.Next()).To(BeFalse())
					Expect(evt.Destination).To(Equal(ERC20Contract1Address))
					Expect(evt.ReturnData).To(Equal(common.Hex2Bytes("0000000000000000000000000000000000000000000000000000000000000001")))
				})
			})
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
                "name": "_amount",
                "type": "uint256"
            }
        ],
        "name": "burn",
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
