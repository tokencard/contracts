package target

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Target is generic target interface
type Target func(context.Context, *ethclient.Client, *bind.TransactOpts, Resolve) (common.Address, error)

// Resolve is used by targets to resolve dependencies
type Resolve func(string) (common.Address, error)

const oraclizeConnectorName = "oraclize-connector.tokencard.eth"
const oraclizeResovlerName = "oraclize-resolver.tokencard.eth"
const tokenWhitelistName = "token-whitelist-v2.tokencard.eth"
const controllerName = "controller-v2.tokencard.eth"
const oracleName = "oracle-v2.tokencard.eth"
const holderName = "holder-v2.tokencard.eth"
const licenceName = "licence-v2.tokencard.eth"
const tknName = "TKN.tokencard.eth"
const WalletDeployerName = "wallet-deployer.tokencard.eth"
const stablecoinName = "DAI.tokencard.eth"

const ENSName = "ENSAddressParam"

const addTokensToTokenWhitelistName = "addTokensToTokenWhitelist"

// const controllerOwnerName = "controllerOwnerAddressParam"
const floatName = "floatAddressParam"

const DevEnvironment = "dev"

const DevControllerAddressName = "devController"

var targets = map[string]Target{
	oracleName:                    deployOracle,
	oraclizeConnectorName:         deployOraclizeConnector,
	oraclizeResovlerName:          deployOraclizeResolver,
	licenceName:                   deployLicence,
	controllerName:                deployController,
	holderName:                    deployHolder,
	tknName:                       deployTkn,
	WalletDeployerName:            deployWalletDeployer,
	tokenWhitelistName:            deployTokenWhitelist,
	DevEnvironment:                deployDev,
	addTokensToTokenWhitelistName: addTokensToTokenWhitelist,
	"add-eth":                     addETH,
}

var plainTokens = []string{
	"ETH",
	"EOS",
	"OMG",
	"BNT",
	"ZRX",
	"MKR",
	"DGD",
	"DGX",
	"GNO",
	"GEN",
	"BAT",
	"REP",
	"SNT",
	"GNT",
	"MANA",
	"LEND",
	"ANT",
	"LEO",
	"LINK",
	"USDC",
	"TUSD",
	"DAI",
}

func init() {
	// Add tokens to targets
	for _, tok := range plainTokens {
		targets[fmt.Sprintf("%s.tokencard.eth", tok)] = deployCoin
	}
}
