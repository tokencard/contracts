#!/bin/bash

set -e -o pipefail

SOLC="docker run --rm -u `id -u` -v $PWD:/solidity --workdir /solidity/contracts ethereum/solc:0.4.25 --optimize /=/"

compile_solidity() {
  echo "compiling ${1}"
  ${SOLC} --overwrite --bin --abi ${1}.sol -o /solidity/build/${1} --combined-json bin-runtime,srcmap-runtime,ast,srcmap,bin
}

contract_sources=(
  'wallet'
  'oracle'
  'holder'
  'tkn'
  'licence'
  'mocks/token'
  'mocks/oraclize-connector'
  'mocks/oraclize-resolver'
  'mocks/base64-exporter'
  'mocks/parseIntScientific-exporter'
  'mocks/tokenWhitelistableExporter'
  'internals/controller'
  'internals/tokenWhitelist'
  'internals/parseIntScientific'
  'externals/ens/PublicResolver'
  'externals/ens/ENSRegistry'
)

for c in "${contract_sources[@]}"
do
    compile_solidity $c
done

GE_PATH=${PWD}/vendor/github.com/ethereum/go-ethereum

# Generate Go bindings from solidity contracts.
ABIGEN="docker run --rm -u `id -u` --workdir /go/src/github/tokencard/contracts -e GOPATH=/go -v $GE_PATH:/go/src/github.com/ethereum/go-ethereum -v $PWD:/go/src/github/tokencard/contracts ethereum/client-go:alltools-v1.8.21 abigen"

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
  "holder/Holder holder.go Holder bindings"
  "tkn/Token token.go Token bindings"
  "licence/Licence licence.go Licence bindings"
  "mocks/token/Token mocks/token.go Token mocks"
  "mocks/oraclize-connector/Oraclize mocks/oraclize-connector.go Oraclize mocks"
  "mocks/oraclize-resolver/OraclizeAddrResolver mocks/oraclize-resolver.go OraclizeAddrResolver mocks"
  "mocks/base64-exporter/Base64Exporter mocks/base64-exporter.go Base64Exporter mocks"
  "mocks/parseIntScientific-exporter/ParseIntScientificExporter mocks/parseIntScientific-exporter.go ParseIntScientificExporter mocks"
  "mocks/tokenWhitelistableExporter/TokenWhitelistableExporter mocks/tokenWhitelistableExporter.go TokenWhitelistableExporter mocks"
  "internals/controller/Controller internals/controller.go Controller internals"
  "internals/tokenWhitelist/TokenWhitelist internals/tokenWhitelist.go TokenWhitelist internals"
  "internals/parseIntScientific/ParseIntScientific internals/parseIntScientific.go ParseIntScientific internals"
  "externals/ens/ENSRegistry/ENSRegistry externals/ens/ens-registry.go ENSRegistry ens"
  "externals/ens/PublicResolver/PublicResolver externals/ens/public-resolver.go PublicResolver ens"
)

for c in "${contracts[@]}"
do
    generate_binding "$c"
done

echo "done"
