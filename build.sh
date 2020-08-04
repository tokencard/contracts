#!/bin/bash

set -e -o pipefail

SOLC_0_6="docker run --rm -u `id -u` -v $PWD:/solidity --workdir /solidity/contracts ethereum/solc:0.6.11 --optimize /=/"
SOLC_0_5="docker run --rm -u `id -u` -v $PWD:/solidity --workdir /solidity/contracts ethereum/solc:0.5.17 --optimize /=/"

compile_solidity() {
  echo "compiling ${2}"
  ${1} --overwrite --bin --abi ${2}.sol -o /solidity/build/${2} --combined-json bin-runtime,srcmap-runtime,ast,srcmap,bin
}

contracts_0_6=(
  'controller'
  'gasProxy'
  'holder'
  'licence'
  'oracle'
  'tokenWhitelist'
  'wallet'
  'walletCache'
  'walletDeployer'
  'externals/upgradeability/UpgradeabilityProxy'
  'mocks/base64Exporter'
  'mocks/burnerToken'
  'mocks/bytesUtilsExporter'
  'mocks/gasBurner'
  'mocks/gasToken'
  'mocks/isValidSignatureExporter'
  'mocks/nonCompliantToken'
  'mocks/oraclize'
  'mocks/parseIntScientificExporter'
  'mocks/token'
  'mocks/tokenWhitelistableExporter'
  'mocks/wallet'
)

contracts_0_5=(
  'externals/ens/ENSRegistry'
  'externals/ens/PublicResolver'
)

for c in "${contracts_0_6[@]}"
do
    compile_solidity "${SOLC_0_6}" $c
done

for c in "${contracts_0_5[@]}"
do
    compile_solidity "${SOLC_0_5}" $c
done


# Generate Go bindings from solidity contracts.
ABIGEN="docker run --rm -u `id -u` --workdir /go/src/github/tokencard/contracts -e GOPATH=/go -v $PWD:/go/src/github/tokencard/contracts ethereum/client-go:alltools-v1.9.12 abigen"

generate_binding() {
  contract=$(echo $1 | awk '{print $1}')
  go_source=$(echo $1 | awk '{print $2}')
  go_type=$(echo $1 | awk '{print $3}')
  package=$(echo $1 | awk '{print $4}')
  echo "Generating binding for ${go_type} (${contract})"
  ${ABIGEN} --abi ./build/${contract}.abi  --bin ./build/${contract}.bin --pkg ${package} --type=${go_type} --out ./pkg/bindings/${go_source}
}

contracts=(
  "controller/Controller controller.go Controller bindings"
  "gasProxy/GasProxy gasProxy.go GasProxy bindings"
  "holder/Holder holder.go Holder bindings"
  "licence/Licence licence.go Licence bindings"
  "oracle/Oracle oracle.go Oracle bindings"
  "tokenWhitelist/TokenWhitelist tokenWhitelist.go TokenWhitelist bindings"
  "wallet/Wallet wallet.go Wallet bindings"
  "walletCache/WalletCache walletCache.go WalletCache bindings"
  "walletDeployer/WalletDeployer walletDeployer.go WalletDeployer bindings"
  "externals/ens/ENSRegistry/ENSRegistry externals/ens/ENSRegistry.go ENSRegistry ens"
  "externals/ens/PublicResolver/PublicResolver externals/ens/PublicResolver.go PublicResolver ens"
  "externals/upgradeability/UpgradeabilityProxy/UpgradeabilityProxy externals/upgradeability/UpgradeabilityProxy.go UpgradeabilityProxy upgradeability"
  "mocks/base64Exporter/Base64Exporter mocks/base64Exporter.go Base64Exporter mocks"
  "mocks/burnerToken/BurnerToken mocks/burnerToken.go BurnerToken mocks"
  "mocks/bytesUtilsExporter/BytesUtilsExporter mocks/bytesUtilsExporter.go BytesUtilsExporter mocks"
  "mocks/isValidSignatureExporter/IsValidSignatureExporter mocks/isValidSignatureExporter.go IsValidSignatureExporter mocks"
  "mocks/nonCompliantToken/NonCompliantToken mocks/nonCompliantToken.go NonCompliantToken mocks"
  "mocks/oraclize/OraclizeConnector mocks/oraclizeConnector.go OraclizeConnector mocks"
  "mocks/oraclize/OraclizeAddrResolver mocks/oraclizeAddrResolver.go OraclizeAddrResolver mocks"
  "mocks/parseIntScientificExporter/ParseIntScientificExporter mocks/parseIntScientificExporter.go ParseIntScientificExporter mocks"
  "mocks/token/Token mocks/token.go Token mocks"
  "mocks/tokenWhitelistableExporter/TokenWhitelistableExporter mocks/tokenWhitelistableExporter.go TokenWhitelistableExporter mocks"
  "mocks/wallet/Wallet mocks/wallet.go Wallet mocks"
  "externals/ens/ENSRegistry/ENSRegistry externals/ens/ENSRegistry.go ENSRegistry ens"
  "externals/ens/PublicResolver/PublicResolver externals/ens/PublicResolver.go PublicResolver ens"
  "externals/upgradeability/UpgradeabilityProxy/UpgradeabilityProxy externals/upgradeability/UpgradeabilityProxy.go UpgradeabilityProxy upgradeability"
)

for c in "${contracts[@]}"
do
    generate_binding "$c"
done

echo "done"