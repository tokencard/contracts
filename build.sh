#!/bin/bash

set -e -o pipefail

SOLC="docker run --rm -u `id -u` -v $PWD:/solidity --workdir /solidity/contracts ethereum/solc:0.4.24 --optimize"

compile_solidity() {
  echo "compiling ${1}"
  ${SOLC} --overwrite --bin --abi ${1}.sol -o /solidity/build/${1} --combined-json bin-runtime,srcmap-runtime,ast,srcmap,bin
}

contract_sources=(
  'wallet'
  'mocks/token'
  'mocks/oraclize-connector'
  'mocks/oraclize-resolver'
  'oracle'
)

for c in "${contract_sources[@]}"
do
    compile_solidity $c
done

GE_PATH=${PWD}/vendor/github.com/ethereum/go-ethereum
if [ ! -d ${GE_PATH} ]
then
  GE_PATH=${GOPATH:-$HOME/go}/src/github.com/ethereum/go-ethereum
fi

# Generate Go bindings from solidity contracts.
ABIGEN="docker run --rm -u `id -u` --workdir /go/src/github/tokencard/contracts -e GOPATH=/go -v $GE_PATH:/go/src/github.com/ethereum/go-ethereum -v $PWD:/go/src/github/tokencard/contracts ethereum/client-go:alltools-v1.8.15 abigen"

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
  "mocks/token/Token mocks/token.go Token mocks"
  "mocks/oraclize-resolver/OraclizeAddrResolver mocks/oraclize-resolver.go OraclizeAddrResolver mocks"
  "mocks/oraclize-connector/Oraclize mocks/oraclize-connector.go Oraclize mocks"
)

for c in "${contracts[@]}"
do
    generate_binding "$c"
done

echo "done"
