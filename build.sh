#!/bin/bash

set -e -o pipefail

SOLC="docker run --rm -u `id -u` -v $PWD:/solidity --workdir /solidity/contracts ethereum/solc:0.5.17 --optimize /=/"

compile_solidity() {
  echo "compiling ${1}"
  ${SOLC} --overwrite --bin --abi ${1}.sol -o /solidity/build/${1} --combined-json bin-runtime,srcmap-runtime,ast,srcmap,bin
}

contract_sources=(
  'controller'
  'gasProxy'
  'holder'
  'licence'
  'oracle'
  'tokenWhitelist'
  'wallet'
  'walletCache'
  'walletDeployer'
  'mocks/base64Exporter'
  'mocks/burnerToken'
  'mocks/bytesUtilsExporter'
  'mocks/isValidSignatureExporter'
  'mocks/nonCompliantToken'
  'mocks/oraclize'
  'mocks/parseIntScientificExporter'
  'mocks/token'
  'mocks/tokenWhitelistableExporter'
  'mocks/walletMock'
  'mocks/gasToken'
  'mocks/gasBurner'
  'externals/ens/ENSRegistry'
  'externals/ens/PublicResolver'
  'externals/upgradeability/UpgradeabilityProxy'
)

for c in "${contract_sources[@]}"
do
    compile_solidity $c
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
  "wallet/Wallet wallet.go Wallet bindings"
  "oracle/Oracle oracle.go Oracle bindings"
  "licence/Licence licence.go Licence bindings"
  "holder/Holder holder.go Holder bindings"
  "controller/Controller controller.go Controller bindings"
  "tokenWhitelist/TokenWhitelist tokenWhitelist.go TokenWhitelist bindings"
  "walletDeployer/WalletDeployer walletDeployer.go WalletDeployer bindings"
  "walletCache/WalletCache walletCache.go WalletCache bindings"
  "gasProxy/GasProxy gasProxy.go GasProxy bindings"
  "mocks/token/Token mocks/token.go Token mocks"
  "mocks/burnerToken/BurnerToken mocks/burnerToken.go BurnerToken mocks"
  "mocks/nonCompliantToken/NonCompliantToken mocks/nonCompliantToken.go NonCompliantToken mocks"
  "mocks/base64Exporter/Base64Exporter mocks/base64Exporter.go Base64Exporter mocks"
  "mocks/oraclize/OraclizeConnector mocks/oraclizeConnector.go OraclizeConnector mocks"
  "mocks/oraclize/OraclizeAddrResolver mocks/oraclizeAddrResolver.go OraclizeAddrResolver mocks"
  "mocks/bytesUtilsExporter/BytesUtilsExporter mocks/bytesUtilsExporter.go BytesUtilsExporter mocks"
  "mocks/isValidSignatureExporter/IsValidSignatureExporter mocks/isValidSignatureExporter.go IsValidSignatureExporter mocks"
  "mocks/parseIntScientificExporter/ParseIntScientificExporter mocks/parseIntScientificExporter.go ParseIntScientificExporter mocks"
  "mocks/tokenWhitelistableExporter/TokenWhitelistableExporter mocks/tokenWhitelistableExporter.go TokenWhitelistableExporter mocks"
  "mocks/gasToken/GasToken mocks/gasToken.go GasToken mocks"
  "mocks/gasBurner/GasBurner mocks/gasBurner.go GasBurner mocks"
  "mocks/walletMock/WalletMock mocks/walletMock.go WalletMock mocks"
  "externals/ens/ENSRegistry/ENSRegistry externals/ens/ENSRegistry.go ENSRegistry ens"
  "externals/ens/PublicResolver/PublicResolver externals/ens/PublicResolver.go PublicResolver ens"
  "externals/upgradeability/UpgradeabilityProxy/UpgradeabilityProxy externals/upgradeability/UpgradeabilityProxy.go UpgradeabilityProxy upgradeability"
)

for c in "${contracts[@]}"
do
    generate_binding "$c"
done

echo "done"
