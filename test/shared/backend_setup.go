package shared

import (
	"context"
	"fmt"
	"math"
	"math/big"
	"strings"
	"time"

	"encoding/hex"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	. "github.com/onsi/gomega"
	gtypes "github.com/onsi/gomega/types"
	"github.com/pkg/errors"
	"github.com/tokencard/contracts/pkg/bindings"
	"github.com/tokencard/contracts/pkg/bindings/externals/ens"
	"github.com/tokencard/contracts/pkg/bindings/internals"
	"github.com/tokencard/contracts/pkg/bindings/mocks"
	"github.com/tokencard/ethertest"
	"golang.org/x/crypto/sha3"
)

var ErrFailedTransaction = errors.New("transaction failed")

func EthToWei(amount int) *big.Int {
	r := big.NewInt(1000000000000000000)
	return r.Mul(r, big.NewInt(int64(amount)))
}

func FinneyToWei(amount int) *big.Int {
	r := big.NewInt(1000000000000000)
	return r.Mul(r, big.NewInt(int64(amount)))
}

func GweiToWei(amount int) *big.Int {
	r := big.NewInt(1000000000)
	return r.Mul(r, big.NewInt(int64(amount)))
}

// Lifted from https://github.com/tronprotocol/tron-demo/blob/master/demo/go-client-api/vendor/github.com/ethereum/go-ethereum/contracts/ens/ens.go
func EnsParentNode(name string) (common.Hash, common.Hash) {
	parts := strings.SplitN(name, ".", 2)
	label := crypto.Keccak256Hash([]byte(parts[0]))
	if len(parts) == 1 {
		return [32]byte{}, label
	} else {
		parentNode, parentLabel := EnsParentNode(parts[1])
		return crypto.Keccak256Hash(parentNode[:], parentLabel[:]), label
	}
}

func LabelHash(label string) common.Hash {
	return crypto.Keccak256Hash([]byte(label))
}

func EnsNode(name string) common.Hash {
	if name == "" {
		return common.Hash{}
	}
	parentNode, parentLabel := EnsParentNode(name)
	return crypto.Keccak256Hash(parentNode[:], parentLabel[:])
}

func stringToHashString(url string) string {
	sha := sha3.NewLegacyKeccak256()
	_, err := sha.Write([]byte(url))
	Expect(err).ToNot(HaveOccurred())

	idSlice := sha.Sum(nil)
	Expect(len(idSlice)).To(Equal(32))

	return hex.EncodeToString(idSlice)
}

var ENSResolver *ens.PublicResolver
var ENSResolverAddress common.Address

var ENSRegistryAddress common.Address
var ENSRegistry *ens.ENSRegistry

var ControllerContract *internals.Controller
var ControllerContractAddress common.Address

var TokenWhitelist *internals.TokenWhitelist
var TokenWhitelistAddress common.Address

var OraclizeResolver *mocks.OraclizeAddrResolver
var OraclizeResolverAddress common.Address

var OraclizeConnector *mocks.OraclizeConnector
var OraclizeConnectorAddress common.Address

var Oracle *bindings.Oracle
var OracleAddress common.Address

var TKNBurner *mocks.BurnerToken
var TKNBurnerAddress common.Address

var TokenHolder *bindings.Holder
var TokenHolderAddress common.Address

var CryptoFloatAddress common.Address

var Licence *bindings.Licence
var LicenceAddress common.Address

// var Wallet *bindings.Wallet
// var WalletAddress common.Address

var TKN *mocks.Token
var TKNAddress common.Address

var Stablecoin *mocks.Token
var StablecoinAddress common.Address

var RandomToken *mocks.Token
var RandomTokenAddress common.Address

var ERC20Contract1 *mocks.Token
var ERC20Contract1Address common.Address

var ERC20Contract2 *mocks.Token
var ERC20Contract2Address common.Address

var NonCompliantERC20 *mocks.NonCompliantToken
var NonCompliantERC20Address common.Address

var OracleName = EnsNode("oracle.tokencard.eth")
var ControllerName = EnsNode("controller.tokencard.eth")
var LicenceName = EnsNode("licence.tokencard.eth")
var TokenWhitelistName = EnsNode("token-whitelist.tokencard.eth")

var Owner *ethertest.Account
var Controller *ethertest.Account
var ControllerOwner *ethertest.Account
var ControllerAdmin *ethertest.Account
var RandomAccount *ethertest.Account
var BankAccount *ethertest.Account
var OraclizeConnectorOwner *ethertest.Account

var TestRig = ethertest.NewTestRig()
var Backend ethertest.TestBackend

func AlmostEqual(val string) gtypes.GomegaMatcher {
	return almostEqual(val)
}

type almostEqual string

func (m almostEqual) FailureMessage(actual interface{}) (message string) {
	return fmt.Sprintf("Actual value %v is not almost equal to expected %v", m, actual)
}

func (m almostEqual) NegatedFailureMessage(actual interface{}) (message string) {
	return fmt.Sprintf("Actual value %v is almost equal to expected %v", m, actual)
}

var maxDelta = big.NewInt(1000000)

func (m almostEqual) Match(actual interface{}) (success bool, err error) {

	expected, ok := big.NewInt(0).SetString(string(m), 10)
	if !ok {
		return false, fmt.Errorf("%q is not a decimal integer", expected)
	}

	actualString, ok := actual.(string)
	if !ok {
		return false, fmt.Errorf("actual value is not a string")
	}

	act, ok := big.NewInt(0).SetString(actualString, 10)
	if !ok {
		return false, fmt.Errorf("%q is not a decimal integer", actualString)
	}

	delta := big.NewInt(0)

	delta = delta.Sub(act, expected)

	return delta.Abs(delta).Cmp(maxDelta) == -1, nil
}

func StringsToByte32(names ...string) [][32]byte {
	r := [][32]byte{}
	for _, n := range names {
		nb := [32]byte{}
		copy(nb[:], []byte(n))
		r = append(r, nb)
	}
	return r
}

func ExponentiateDecimals(decimals uint8) *big.Int {
	base := big.NewInt(10)
	expDec := big.NewInt(1)
	for i := uint8(0); i < decimals; i++ {
		expDec.Mul(expDec, base)
	}
	return expDec
}

func DecimalsToMagnitude(decimals *big.Int) *big.Int {
	return new(big.Int).Exp(big.NewInt(10), decimals, nil)
}

func InitializeBackend() error {

	Owner = ethertest.NewAccount()
	ControllerOwner = ethertest.NewAccount()
	ControllerAdmin = ethertest.NewAccount()
	Controller = ethertest.NewAccount()
	RandomAccount = ethertest.NewAccount()
	BankAccount = ethertest.NewAccount()
	OraclizeConnectorOwner = ethertest.NewAccount()

	TestRig.AddGenesisAccountAllocation(ControllerOwner.Address(), EthToWei(1))
	TestRig.AddGenesisAccountAllocation(ControllerAdmin.Address(), EthToWei(1))
	TestRig.AddGenesisAccountAllocation(Controller.Address(), EthToWei(1000))
	TestRig.AddGenesisAccountAllocation(RandomAccount.Address(), EthToWei(1000))
	TestRig.AddGenesisAccountAllocation(OraclizeConnectorOwner.Address(), EthToWei(1000))
	TestRig.AddGenesisAccountAllocation(BankAccount.Address(), EthToWei(1000))
	TestRig.AddGenesisAccountAllocation(Owner.Address(), EthToWei(1))

	Backend = TestRig.NewTestBackend(ethertest.WithBlockchainTime(time.Date(2018, 9, 13, 15, 10, 0, 0, time.Local)))

	var err error
	var tx *types.Transaction

	// Deploy a stablecoin mock token.
	StablecoinAddress, tx, Stablecoin, err = mocks.DeployToken(BankAccount.TransactOpts(), Backend)
	if err != nil {
		return err
	}
	Backend.Commit()
	err = verifyTransaction(tx)
	if err != nil {
		return errors.Wrap(err, "deploying Stablecoin token contract")
	}

	ControllerContractAddress, tx, ControllerContract, err = internals.DeployController(BankAccount.TransactOpts(), Backend, ControllerOwner.Address(), false)
	if err != nil {
		return err
	}
	Backend.Commit()
	err = verifyTransaction(tx)
	if err != nil {
		return errors.Wrap(err, "deploying controller contract")
	}

	ControllerContract.AddAdmin(ControllerOwner.TransactOpts(), ControllerAdmin.Address())
	if err != nil {
		return err
	}
	Backend.Commit()
	err = verifyTransaction(tx)
	if err != nil {
		return errors.Wrap(err, "adding controller admin address")
	}

	tx, err = ControllerContract.AddController(ControllerAdmin.TransactOpts(), Controller.Address())
	if err != nil {
		return err
	}

	Backend.Commit()
	err = verifyTransaction(tx)
	if err != nil {
		return errors.Wrap(err, "adding controller address")
	}

	ENSRegistryAddress, tx, ENSRegistry, err = ens.DeployENSRegistry(BankAccount.TransactOpts(), Backend)
	if err != nil {
		return err
	}
	Backend.Commit()
	err = verifyTransaction(tx)
	if err != nil {
		return errors.Wrap(err, "deploying ENS registry")
	}

	tx, err = ENSRegistry.SetSubnodeOwner(BankAccount.TransactOpts(), EnsNode(""), LabelHash("eth"), BankAccount.Address())
	if err != nil {
		return err
	}
	Backend.Commit()
	err = verifyTransaction(tx)
	if err != nil {
		return errors.Wrap(err, "setting ENS 'eth' node owner")
	}

	tx, err = ENSRegistry.SetSubnodeOwner(BankAccount.TransactOpts(), EnsNode("eth"), LabelHash("tokencard"), BankAccount.Address())
	if err != nil {
		return err
	}

	Backend.Commit()
	err = verifyTransaction(tx)
	if err != nil {
		return errors.Wrap(err, "setting ENS 'tokencard' node owner")
	}

	tx, err = ENSRegistry.SetSubnodeOwner(BankAccount.TransactOpts(), EnsNode("tokencard.eth"), LabelHash("controller"), BankAccount.Address())
	if err != nil {
		return err
	}

	Backend.Commit()
	err = verifyTransaction(tx)
	if err != nil {
		return errors.Wrap(err, "setting ENS 'controller' node owner")
	}

	tx, err = ENSRegistry.SetSubnodeOwner(BankAccount.TransactOpts(), EnsNode("tokencard.eth"), LabelHash("oracle"), BankAccount.Address())
	if err != nil {
		return err
	}
	Backend.Commit()
	err = verifyTransaction(tx)
	if err != nil {
		return errors.Wrap(err, "setting ENS 'oracle' node owner")
	}

	tx, err = ENSRegistry.SetSubnodeOwner(BankAccount.TransactOpts(), EnsNode("tokencard.eth"), LabelHash("licence"), BankAccount.Address())
	if err != nil {
		return err
	}
	Backend.Commit()
	err = verifyTransaction(tx)
	if err != nil {
		return errors.Wrap(err, "setting ENS 'licence' node owner")
	}

	tx, err = ENSRegistry.SetSubnodeOwner(BankAccount.TransactOpts(), EnsNode("tokencard.eth"), LabelHash("token-whitelist"), BankAccount.Address())
	if err != nil {
		return err
	}
	Backend.Commit()
	err = verifyTransaction(tx)
	if err != nil {
		return errors.Wrap(err, "setting ENS 'token-whitelist' node owner")
	}

	ENSResolverAddress, tx, ENSResolver, err = ens.DeployPublicResolver(BankAccount.TransactOpts(), Backend, ENSRegistryAddress)
	if err != nil {
		return err
	}
	Backend.Commit()
	err = verifyTransaction(tx)
	if err != nil {
		return errors.Wrap(err, "deploying address resolver")
	}

	{
		// Register controller with ENS
		tx, err = ENSRegistry.SetResolver(BankAccount.TransactOpts(), ControllerName, ENSResolverAddress)
		if err != nil {
			return err
		}
		Backend.Commit()
		err = verifyTransaction(tx)
		if err != nil {
			return errors.Wrap(err, "setting controller ENS node resolver")
		}

		tx, err = ENSResolver.SetAddr(BankAccount.TransactOpts(), ControllerName, ControllerContractAddress)
		if err != nil {
			return err
		}
		Backend.Commit()
		err = verifyTransaction(tx)
		if err != nil {
			return errors.Wrap(err, "setting controller ENS node resolver's target address")
		}
	}

	OraclizeConnectorAddress, tx, OraclizeConnector, err = mocks.DeployOraclizeConnector(BankAccount.TransactOpts(), Backend, OraclizeConnectorOwner.Address())
	if err != nil {
		return err
	}
	Backend.Commit()
	err = verifyTransaction(tx)
	if err != nil {
		return errors.Wrap(err, "deploying Oraclize connector")
	}

	OraclizeResolverAddress, tx, OraclizeResolver, err = mocks.DeployOraclizeAddrResolver(BankAccount.TransactOpts(), Backend, OraclizeConnectorAddress)
	if err != nil {
		return err
	}
	Backend.Commit()
	err = verifyTransaction(tx)
	if err != nil {
		return errors.Wrap(err, "deploying Oraclize address resolver")
	}

	TokenWhitelistAddress, tx, TokenWhitelist, err = internals.DeployTokenWhitelist(BankAccount.TransactOpts(), Backend, ENSRegistryAddress, OracleName, ControllerName, Owner.Address(), true, StablecoinAddress)
	if err != nil {
		return err
	}
	Backend.Commit()
	err = verifyTransaction(tx)
	if err != nil {
		return errors.Wrap(err, "deploying oracle contract")
	}

	{
		// Register tokenWhitelist with ENS
		tx, err = ENSRegistry.SetResolver(BankAccount.TransactOpts(), TokenWhitelistName, ENSResolverAddress)
		if err != nil {
			return err
		}
		Backend.Commit()
		err = verifyTransaction(tx)
		if err != nil {
			return errors.Wrap(err, "setting tokenWhitelist ENS node resolver")
		}

		tx, err = ENSResolver.SetAddr(BankAccount.TransactOpts(), TokenWhitelistName, TokenWhitelistAddress)
		if err != nil {
			return err
		}
		Backend.Commit()
		err = verifyTransaction(tx)
		if err != nil {
			return errors.Wrap(err, "setting tokenWhitelist ENS node resolver's target address")
		}
	}

	// Deploy the Token oracle contract.
	OracleAddress, tx, Oracle, err = bindings.DeployOracle(BankAccount.TransactOpts(), Backend, OraclizeResolverAddress, ENSRegistryAddress, ControllerName, TokenWhitelistName)
	if err != nil {
		return err
	}
	Backend.Commit()
	err = verifyTransaction(tx)
	if err != nil {
		return errors.Wrap(err, "deploying oracle contract")
	}

	{
		// Register oracle with ENS

		tx, err = ENSRegistry.SetResolver(BankAccount.TransactOpts(), OracleName, ENSResolverAddress)
		if err != nil {
			return err
		}
		Backend.Commit()
		err = verifyTransaction(tx)
		if err != nil {
			return errors.Wrap(err, "setting oracle ENS node resolver")
		}

		tx, err = ENSResolver.SetAddr(BankAccount.TransactOpts(), OracleName, OracleAddress)
		if err != nil {
			return err
		}
		Backend.Commit()
		err = verifyTransaction(tx)
		if err != nil {
			return errors.Wrap(err, "setting oracle ENS node resolver's target address")
		}
	}

	// Deploy the real TKN contract with burner functionality.
	TKNBurnerAddress, tx, TKNBurner, err = mocks.DeployBurnerToken(Owner.TransactOpts(), Backend)
	if err != nil {
		return err
	}
	Backend.Commit()
	err = verifyTransaction(tx)
	if err != nil {
		return errors.Wrap(err, "deploying TKN contract")
	}

    // Deploy the Token holder contract.
	TokenHolderAddress, tx, TokenHolder, err = bindings.DeployHolder(Controller.TransactOpts(), Backend, Owner.Address(), false, TKNBurnerAddress, ENSRegistryAddress, TokenWhitelistName)
	if err != nil {
		return err
	}
	Backend.Commit()
	err = verifyTransaction(tx)
	if err != nil {
		return errors.Wrap(err, "deploying holder contract")
	}

	// Deploy the Token licence contract.
	CryptoFloatAddress = common.HexToAddress(stringToHashString("CryptoFloatAddress"))
	LicenceAddress, tx, Licence, err = bindings.DeployLicence(BankAccount.TransactOpts(), Backend, Owner.Address(), true, big.NewInt(10), CryptoFloatAddress, TokenHolderAddress, common.HexToAddress("0x0"))
	if err != nil {
		return err
	}

	Backend.Commit()
	err = verifyTransaction(tx)
	if err != nil {
		return errors.Wrap(err, "deploying licence contract")
	}

	{
		// Register licence with ENS

		tx, err = ENSRegistry.SetResolver(BankAccount.TransactOpts(), LicenceName, ENSResolverAddress)
		if err != nil {
			return err
		}

		Backend.Commit()
		err = verifyTransaction(tx)
		if err != nil {
			return errors.Wrap(err, "setting licence ENS node resolver")
		}

		tx, err = ENSResolver.SetAddr(BankAccount.TransactOpts(), LicenceName, LicenceAddress)
		if err != nil {
			return err
		}

		Backend.Commit()
		err = verifyTransaction(tx)
		if err != nil {
			return errors.Wrap(err, "setting licence ENS node resolver's target address")
		}
	}

	err = BankAccount.Transfer(Backend, Controller.Address(), EthToWei(1))
	if err != nil {
		return errors.Wrap(err, "crediting controller account with ETH")
	}

	// Deploy a TKN mock token.
	TKNAddress, tx, TKN, err = mocks.DeployToken(BankAccount.TransactOpts(), Backend)
	if err != nil {
		return err
	}
	Backend.Commit()
	err = verifyTransaction(tx)
	if err != nil {
		return errors.Wrap(err, "deploying TKN token contract")
	}

	// Deploy additional ERC20 mock tokens.
	ERC20Contract1Address, tx, ERC20Contract1, err = mocks.DeployToken(BankAccount.TransactOpts(), Backend)
	if err != nil {
		return err
	}
	Backend.Commit()
	err = verifyTransaction(tx)
	if err != nil {
		return errors.Wrap(err, "deploying ERC20-1 token contract")
	}

	ERC20Contract2Address, tx, ERC20Contract2, err = mocks.DeployToken(BankAccount.TransactOpts(), Backend)
	if err != nil {
		return err
	}
	Backend.Commit()
	err = verifyTransaction(tx)
	if err != nil {
		return errors.Wrap(err, "deploying ERC20-2 token contract")
	}

	NonCompliantERC20Address, tx, NonCompliantERC20, err = mocks.DeployNonCompliantToken(BankAccount.TransactOpts(), Backend)
	if err != nil {
		return err
	}
	Backend.Commit()
	err = verifyTransaction(tx)
	if err != nil {
		return errors.Wrap(err, "deploying NonCompliantERC20 token contract")
	}

	// Add the TKN token to the oracle list.
	tx, err = TokenWhitelist.AddTokens(Controller.TransactOpts(), []common.Address{TKNAddress}, StringsToByte32("TKN"), []*big.Int{ExponentiateDecimals(8)}, []bool{true}, []bool{true}, big.NewInt(20180913153211))
	if err != nil {
		return err
	}
	Backend.Commit()
	err = verifyTransaction(tx)
	if err != nil {
		return errors.Wrap(err, "adding TKN token to oracle")
	}

	// Add the Stablecoin token to the oracle list.
	tx, err = TokenWhitelist.AddTokens(Controller.TransactOpts(), []common.Address{StablecoinAddress}, StringsToByte32("DAI"), []*big.Int{ExponentiateDecimals(18)}, []bool{true}, []bool{true}, big.NewInt(20180913153211))
	if err != nil {
		return err
	}
	Backend.Commit()
	err = verifyTransaction(tx)
	if err != nil {
		return errors.Wrap(err, "adding Stablecoin token to oracle")
	}

	// Update the exchange rate of the TKN token.
	tx, err = TokenWhitelist.UpdateTokenRate(Controller.TransactOpts(), TKNAddress, big.NewInt(int64(0.00001633*math.Pow10(18))), big.NewInt(20180913153211))
	if err != nil {
		return err
	}
	Backend.Commit()
	err = verifyTransaction(tx)
	if err != nil {
		return errors.Wrap(err, "updating TKN token rate")
	}

	// Update the exchange rate of the Stablecoin token.
	tx, err = TokenWhitelist.UpdateTokenRate(Controller.TransactOpts(), StablecoinAddress, big.NewInt(int64(0.007462*math.Pow10(18))), big.NewInt(20180913153211))
	if err != nil {
		return err
	}
	Backend.Commit()
	err = verifyTransaction(tx)
	if err != nil {
		return errors.Wrap(err, "updating Stablecoin token rate")
	}

	return nil
}

func verifyTransaction(tx *types.Transaction) error {
	receipt, err := Backend.TransactionReceipt(context.Background(), tx.Hash())
	if err != nil {
		return err
	}
	if receipt.Status != types.ReceiptStatusSuccessful {
		return ErrFailedTransaction
	}
	return nil
}
