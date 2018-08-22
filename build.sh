#!/bin/bash

set -e -o pipefail

SOLC="docker run --rm -u `id -u` -v $PWD:/solidity --workdir /solidity/contracts ethereum/solc:0.4.24"

compile_solidity() {
  echo "compiling ${1}"
  ${SOLC} --overwrite --bin --abi ${1}.sol -o /solidity/build/${1} --combined-json bin-runtime,srcmap-runtime,ast,srcmap,bin
}

contract_sources=(
  'wallet'
  'token'
  'oracle'
  'oracleV2'
  'oraclize'
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
  echo "Generating binding for ${go_type} (${contract})"
  bin=$(awk '{print "0x" $0}' ./build/${contract}.bin)
  echo $bin > ./build/${contract}.bin
  ${ABIGEN} --abi ./build/${contract}.abi  --bin ./build/${contract}.bin --pkg bindings --type=$go_type --out ./pkg/bindings/$go_source
  # rm temp_bin_file
}

contracts=(
  "wallet/Wallet wallet.go Wallet"
  "oracle/Oracle oracle.go Oracle"
  "token/Token token.go Token"
  "oracleV2/Oracle oracleV2.go OracleV2"
  "oraclize/Oraclize mock_oraclize.go MockOraclize"
  "oraclize/OraclizeAddrResolver mock_oraclize_addr_resolver.go MockOraclizeAddrResolver"
)

for c in "${contracts[@]}"
do
    generate_binding "$c"
done

echo "done."
