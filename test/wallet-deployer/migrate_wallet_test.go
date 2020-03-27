package wallet_deployer_test

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/tokencard/contracts/v3/pkg/bindings"
	. "github.com/tokencard/contracts/v3/test/shared"
	"github.com/tokencard/ethertest"
)

var _ = Describe("Migrate Wallet", func() {

	var MigratedWallet *bindings.Wallet

	When("no wallets are cached and a controller migrates a Wallet with non-initialized security settings", func() {

		BeforeEach(func() {

			RandomWalletAddress, tx, _, err := bindings.DeployWallet(BankAccount.TransactOpts(), Backend, Owner.Address(), false, ENSRegistryAddress, TokenWhitelistName, ControllerName, LicenceName, EthToWei(2))
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeTrue())

			tx, err = WalletDeployer.MigrateWallet(Controller.TransactOpts(), Owner.Address(), RandomWalletAddress, false, false, false, false, EthToWei(2), FinneyToWei(1), EthToWei(1), []common.Address{common.HexToAddress("0x1"), common.HexToAddress("0x2")})
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeTrue())

			MigratedWalletAddress, err := WalletDeployer.DeployedWallets(nil, Owner.Address())

			MigratedWallet, err = bindings.NewWallet(MigratedWalletAddress, Backend)
			Expect(err).ToNot(HaveOccurred())
		})

		It("should NOT set the whitelist flag", func() {
			initialized, err := MigratedWallet.IsSetWhitelist(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(initialized).To(BeFalse())
		})

		It("should NOT make the GasTopUpLimit updateable", func() {
			initialized, err := MigratedWallet.GasTopUpLimitControllerConfirmationRequired(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(initialized).To(BeFalse())
		})

		It("should NOT make SpendLimit updateable", func() {
			initialized, err := MigratedWallet.SpendLimitControllerConfirmationRequired(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(initialized).To(BeFalse())
		})

		It("should NOT make LoadLimit updateable", func() {
			initialized, err := MigratedWallet.LoadLimitControllerConfirmationRequired(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(initialized).To(BeFalse())
		})

		It("should NOT add the whitelisted addresses to the whitelist", func() {
			isWhitelisted, err := MigratedWallet.WhitelistMap(nil, common.HexToAddress("0x1"))
			Expect(err).ToNot(HaveOccurred())
			Expect(isWhitelisted).To(BeFalse())

			isWhitelisted, err = MigratedWallet.WhitelistMap(nil, common.HexToAddress("0x2"))
			Expect(err).ToNot(HaveOccurred())
			Expect(isWhitelisted).To(BeFalse())
		})

		It("should NOT update the spend limit to 2 ETH", func() {
			sl, err := MigratedWallet.SpendLimitValue(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(sl.String()).To(Equal(EthToWei(1).String()))
		})

		It("should NOT update the  gasTopUp limit to 1 finney", func() {
			sl, err := MigratedWallet.GasTopUpLimitValue(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(sl.String()).To(Equal(FinneyToWei(500).String()))
		})

		It("should NOT increase the loadLimit", func() {
			sl, err := MigratedWallet.LoadLimitValue(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(sl.String()).To(Equal(GweiToWei(10).String()))
		})

	})

	When("no wallets are cached and a controller migrates a Wallet and send 1000 wei", func() {

		var RandomWalletAddress common.Address
		var RandomOwner common.Address
		var tx *types.Transaction
		var err error

		BeforeEach(func() {
			privateKey, _ := crypto.GenerateKey()
			RandomOwner = crypto.PubkeyToAddress(privateKey.PublicKey)
			RandomWalletAddress, tx, _, err = bindings.DeployWallet(BankAccount.TransactOpts(), Backend, RandomOwner, false, ENSRegistryAddress, TokenWhitelistName, ControllerName, LicenceName, EthToWei(2))
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeTrue())

			tx, err = WalletDeployer.MigrateWallet(Controller.TransactOpts(ethertest.WithValue(big.NewInt(1000))), RandomOwner, RandomWalletAddress, true, true, true, true, EthToWei(2), FinneyToWei(1), GweiToWei(1), []common.Address{common.HexToAddress("0x1"), common.HexToAddress("0x2")})
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeTrue())
		})

		It("DeployedWallets should return the new Wallet address", func() {
			addr, err := WalletDeployer.DeployedWallets(nil, RandomOwner)
			Expect(err).ToNot(HaveOccurred())
			Expect(addr).ToNot(Equal(common.HexToAddress("0x0")))
		})

		It("should emit a MigratedWallet event", func() {

			addr, err := WalletDeployer.DeployedWallets(nil, RandomOwner)
			Expect(err).ToNot(HaveOccurred())

			it, err := WalletDeployer.FilterMigratedWallet(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(it.Next()).To(BeTrue())
			Expect(it.Event.Wallet).To(Equal(addr))
			Expect(it.Event.OldWallet).To(Equal(RandomWalletAddress))
			Expect(it.Event.Owner).To(Equal(RandomOwner))
			Expect(it.Event.Paid.String()).To(Equal("1000"))
		})

		It("should have cached Wallet count 0", func() {
			ccc, err := WalletCache.CachedWalletsCount(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(ccc.String()).To(Equal("0"))
		})

		It("should increase the owner's balance by 1000 wei", func() {
			b, e := Backend.BalanceAt(context.Background(), RandomOwner, nil)
			Expect(e).ToNot(HaveOccurred())
			Expect(b.String()).To(Equal("1000"))
		})

		It("should fail if a wallet is already deployed/migrated for this owner", func() {
			tx, err = WalletDeployer.MigrateWallet(Controller.TransactOpts(ethertest.WithGasLimit(5000000)), RandomOwner, RandomWalletAddress, true, true, true, true, EthToWei(2), FinneyToWei(1), EthToWei(1000), []common.Address{common.HexToAddress("0x1"), common.HexToAddress("0x2")})
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeFalse())
			returnData, _ := ethCall(tx)
			Expect(string(returnData[len(returnData)-64:])).To(ContainSubstring("wallet already deployed for owner"))
		})

		It("should emit CachedWallet event", func() {

			addr, err := WalletDeployer.DeployedWallets(nil, RandomOwner)
			Expect(err).ToNot(HaveOccurred())

			it, err := WalletCache.FilterCachedWallet(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(it.Next()).To(BeTrue())
			Expect(it.Event.Wallet).To(Equal(addr))
		})

		When("a new Wallet is migrated", func() {

			BeforeEach(func() {
				MigratedWalletAddress, err := WalletDeployer.DeployedWallets(nil, RandomOwner)
				MigratedWallet, err = bindings.NewWallet(MigratedWalletAddress, Backend)
				Expect(err).ToNot(HaveOccurred())
			})

			It("should emit a setSpendLimit set event", func() {
				it, err := MigratedWallet.FilterSetSpendLimit(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(it.Next()).To(BeTrue())
				evt := it.Event
				Expect(it.Next()).To(BeFalse())
				Expect(evt.Sender).To(Equal(WalletDeployerAddress))
				Expect(evt.Amount).To(Equal(EthToWei(2)))
			})

			It("should keep the available spend Limit  to 1 ETH", func() {
				av, err := MigratedWallet.SpendLimitAvailable(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(av.String()).To(Equal(EthToWei(1).String()))
			})

			It("should update the spend limit to 2 ETH", func() {
				sl, err := MigratedWallet.SpendLimitValue(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(sl.String()).To(Equal(EthToWei(2).String()))
			})

			It("should make SpendLimit updateable", func() {
				initialized, err := MigratedWallet.SpendLimitControllerConfirmationRequired(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(initialized).To(BeTrue())
			})

			It("should update the gasTopUp limit to 1 finney", func() {
				sl, err := MigratedWallet.GasTopUpLimitValue(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(sl.String()).To(Equal(FinneyToWei(1).String()))
			})

			It("should decrease the available  gasTopUpLimit  to 1 Finney", func() {
				av, err := MigratedWallet.GasTopUpLimitAvailable(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(av.String()).To(Equal(FinneyToWei(1).String()))
			})

			It("should make GasToUpLimit updateable", func() {
				initialized, err := MigratedWallet.GasTopUpLimitControllerConfirmationRequired(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(initialized).To(BeTrue())
			})

			It("should emit a setGasTopUpLimit event", func() {
				it, err := MigratedWallet.FilterSetGasTopUpLimit(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(it.Next()).To(BeTrue())
				evt := it.Event
				Expect(it.Next()).To(BeFalse())
				Expect(evt.Sender).To(Equal(WalletDeployerAddress))
				Expect(evt.Amount).To(Equal(FinneyToWei(1)))
			})

			It("should make LoadLimit updateable", func() {
				initialized, err := MigratedWallet.LoadLimitControllerConfirmationRequired(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(initialized).To(BeTrue())
			})

			It("should emit a setLoadLimit set event", func() {
				it, err := MigratedWallet.FilterSetLoadLimit(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(it.Next()).To(BeTrue())
				evt := it.Event
				Expect(it.Next()).To(BeFalse())
				Expect(evt.Sender).To(Equal(WalletDeployerAddress))
				Expect(evt.Amount).To(Equal(GweiToWei(1)))
			})

			It("should decrease the available loadLimit to 1000 USD", func() {
				av, err := MigratedWallet.LoadLimitAvailable(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(av.String()).To(Equal(GweiToWei(1).String()))
			})

			It("should decrease the loadLimit to 1000 USD", func() {
				sl, err := MigratedWallet.LoadLimitValue(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(sl.String()).To(Equal(GweiToWei(1).String()))
			})

			It("should update the Whitelist initializedWhitelist flag", func() {
				initialized, err := MigratedWallet.IsSetWhitelist(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(initialized).To(BeTrue())
			})

			It("should update the TopUpLimit initializedTopup flag", func() {
				initialized, err := MigratedWallet.GasTopUpLimitControllerConfirmationRequired(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(initialized).To(BeTrue())
			})

			It("should add the whitelisted addresses to the whitelist", func() {
				isWhitelisted, err := MigratedWallet.WhitelistMap(nil, common.HexToAddress("0x1"))
				Expect(err).ToNot(HaveOccurred())
				Expect(isWhitelisted).To(BeTrue())

				isWhitelisted, err = MigratedWallet.WhitelistMap(nil, common.HexToAddress("0x2"))
				Expect(err).ToNot(HaveOccurred())
				Expect(isWhitelisted).To(BeTrue())
			})

			It("should emit AddedToWhitelist event", func() {
				it, err := MigratedWallet.FilterAddedToWhitelist(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(it.Next()).To(BeTrue())
				evt := it.Event
				Expect(it.Next()).To(BeFalse())
				Expect(evt.Sender).To(Equal(WalletDeployerAddress))
				Expect(evt.Addresses).To(Equal([]common.Address{common.HexToAddress("0x1"), common.HexToAddress("0x2")}))
			})
		})

	}) //no wallets chached

	When("the worng owner is passed in", func() {
		It("should fail", func() {
			RandomWalletAddress, tx, _, err := bindings.DeployWallet(BankAccount.TransactOpts(), Backend, Owner.Address(), false, ENSRegistryAddress, TokenWhitelistName, ControllerName, LicenceName, EthToWei(2))
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeTrue())

			tx, err = WalletDeployer.MigrateWallet(Controller.TransactOpts(ethertest.WithGasLimit(5000000)), Controller.Address(), RandomWalletAddress, false, false, false, false, EthToWei(2), FinneyToWei(1), EthToWei(1000), []common.Address{common.HexToAddress("0x1"), common.HexToAddress("0x2")})
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeFalse())
			returnData, _ := ethCall(tx)
			Expect(string(returnData[len(returnData)-64:])).To(ContainSubstring("owner mismatch"))
		})
	})

	When("a random account tries to migrate a Wallet", func() {
		It("should fail", func() {
			tx, err := WalletDeployer.MigrateWallet(RandomAccount.TransactOpts(ethertest.WithGasLimit(5000000)), Owner.Address(), RandomAccount.Address(), false, false, false, false, EthToWei(1), FinneyToWei(2), EthToWei(1000), []common.Address{common.HexToAddress("0x1"), common.HexToAddress("0x2")})
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeFalse())
		})
	})

	When("one wallet is cached", func() {
		BeforeEach(func() {
			tx, err := WalletCache.CacheWallet(Controller.TransactOpts())
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeTrue())
		})

		When("a controller migrates a Wallet", func() {

			var RandomWalletAddress common.Address
			var tx *types.Transaction
			var err error

			BeforeEach(func() {

				RandomWalletAddress, tx, _, err = bindings.DeployWallet(BankAccount.TransactOpts(), Backend, Owner.Address(), false, ENSRegistryAddress, TokenWhitelistName, ControllerName, LicenceName, EthToWei(2))
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())

				tx, err = WalletDeployer.MigrateWallet(Controller.TransactOpts(), Owner.Address(), RandomWalletAddress, false, false, false, false, EthToWei(1), FinneyToWei(2), EthToWei(1000), []common.Address{common.HexToAddress("0x1"), common.HexToAddress("0x2")})
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})

			It("should emit MigratedWallet event", func() {

				addr, err := WalletDeployer.DeployedWallets(nil, Owner.Address())
				Expect(err).ToNot(HaveOccurred())

				it, err := WalletDeployer.FilterMigratedWallet(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(it.Next()).To(BeTrue())
				Expect(it.Event.Wallet).To(Equal(addr))
				Expect(it.Event.OldWallet).To(Equal(RandomWalletAddress))
				Expect(it.Event.Owner).To(Equal(Owner.Address()))
			})

			It("should have cached Wallet count 0", func() {
				ccc, err := WalletCache.CachedWalletsCount(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(ccc.String()).To(Equal("0"))
			})

			It("should emit CachedWallet event", func() {

				addr, err := WalletDeployer.DeployedWallets(nil, Owner.Address())
				Expect(err).ToNot(HaveOccurred())

				it, err := WalletCache.FilterCachedWallet(nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(it.Next()).To(BeTrue())
				Expect(it.Event.Wallet).To(Equal(addr))
			})

		})
	})

})
