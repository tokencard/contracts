package gas_proxy_test

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/ethereum/go-ethereum/common"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/tokencard/contracts/v3/pkg/bindings"
	"github.com/tokencard/contracts/v3/pkg/bindings/externals/upgradeability"
	. "github.com/tokencard/contracts/v3/test/shared"
)

var WalletImplementationAddress common.Address

var Proxy *upgradeability.UpgradeabilityProxy
var ProxyAddress common.Address

var ProxyWallet *bindings.Wallet

func SignData(chainId *big.Int, address common.Address, nonce *big.Int, data []byte, prv *ecdsa.PrivateKey) ([]byte, error) {
	relayMessage := fmt.Sprintf("monolith:%s%s%s%s", abi.U256(chainId), address, abi.U256(nonce), data)
	hash := crypto.Keccak256([]byte(relayMessage))
	ethMessage := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(hash), hash)
	hash = crypto.Keccak256([]byte(ethMessage))
	sig, err := crypto.Sign(hash, prv)
	if err != nil {
		return nil, err
	}
	if len(sig) != 65 {
		return nil, errors.New("invalid sig len")
	}
	sig[64] += 27
	return sig, nil
}

var _ = Describe("ExecuteMetaTransaction", func() {

	When("a wallet proxy is deployed and eth is deposited into it", func() {
		BeforeEach(func() {
			WalletImplementationAddress, tx, _, err := bindings.DeployWallet(BankAccount.TransactOpts(), Backend)
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeTrue())

			ProxyAddress, tx, Proxy, err = upgradeability.DeployUpgradeabilityProxy(BankAccount.TransactOpts(), Backend, WalletImplementationAddress, nil)
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeTrue())

			ProxyWallet, err = bindings.NewWallet(ProxyAddress, Backend)
			tx, err = ProxyWallet.InitializeWallet(Owner.TransactOpts(), Owner.Address(), true, ENSRegistryAddress, TokenWhitelistNode, ControllerNode, LicenceNode, EthToWei(1))
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeTrue())
			// deposit eth
			BankAccount.MustTransfer(Backend, ProxyAddress, EthToWei(2))
		})

		When("the controller proxy relays an owner signed transaction", func() {
			BeforeEach(func() {
				// create TX data and sign with owner's key
				walletABI, err := abi.JSON(strings.NewReader(bindings.WalletABI))
				Expect(err).ToNot(HaveOccurred())
				data, err := walletABI.Pack("transfer", RandomAccount.Address(), common.HexToAddress("0x0"), EthToWei(1))
				Expect(err).ToNot(HaveOccurred())
				nonce := big.NewInt(0)
				chainId := big.NewInt(1337)
				signature, err := SignData(chainId, ProxyAddress, nonce, data, Owner.PrivKey())
				Expect(err).ToNot(HaveOccurred())
				// create relayed transaction data
				data, err = walletABI.Pack("executeRelayedTransaction", nonce, data, signature)
				Expect(err).ToNot(HaveOccurred())

				tx, err := GasProxy.ExecuteTransaction(Controller.TransactOpts(), ProxyAddress, big.NewInt(0), data)
				Expect(err).ToNot(HaveOccurred())
				Backend.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})

			It("should decrease the wallet's ETH balance ", func() {
				b, err := Backend.BalanceAt(context.Background(), ProxyAddress, nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(b.String()).To(Equal(EthToWei(1).String()))
			})

		})
	})

})
