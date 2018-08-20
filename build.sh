#!/usr/bin/env bash

set -e

SOLC="docker run --rm -v $PWD:/solidity ethereum/solc:0.4.24"
echo ${SOLC}

# Compile solidity contracts and hex prefix bytecode string.
${SOLC} --overwrite  --bin --abi ./contracts/wallet.sol         -o ./build/wallet     && bin=$(awk '{print "0x" $0}' ./build/wallet/Wallet.bin)         && echo $bin > ./build/wallet/Wallet.bin && echo "success 1"
${SOLC} --overwrite  --bin --abi ./contracts/token.sol          -o ./build/token      && bin=$(awk '{print "0x" $0}' ./build/token/Token.bin)           && echo $bin > ./build/token/Token.bin && echo "success 2"
${SOLC} --overwrite  --bin --abi ./contracts/oracle.sol         -o ./build/oracle     && bin=$(awk '{print "0x" $0}' ./build/oracle/Oracle.bin)         && echo $bin > ./build/oracle/Oracle.bin && echo "success 3"
${SOLC} --overwrite  --bin --abi ./contracts/old/card.sol       -o ./build/card       && bin=$(awk '{print "0x" $0}' ./build/card/Card.bin)             && echo $bin > ./build/card/Card.bin && echo "success 4"
${SOLC} --overwrite  --bin --abi ./contracts/old/controller.sol -o ./build/controller && bin=$(awk '{print "0x" $0}' ./build/controller/Controller.bin) && echo $bin > ./build/controller/Controller.bin && echo "success 5"

# create combined-json for code coverage
${SOLC} --overwrite --combined-json bin-runtime,srcmap-runtime,ast,srcmap,bin -o ./build ./contracts/wallet.sol ./contracts/oracle.sol ./contracts/token.sol ./contracts/old/controller.sol ./contracts/old/card.sol && echo "success 5"

GE_PATH=${PWD}/vendor/github.com/ethereum/go-ethereum
if [ ! -d ${GE_PATH} ]
then
  echo "if statement"
  GE_PATH=${GOPATH:-$HOME/go}/src/github.com/ethereum/go-ethereum
fi

# Generate Go bindings from solidity contracts.
ABIGEN="docker run --rm --workdir /go/src/github/tokencard/contracts -e GOPATH=/go -v $GE_PATH:/go/src/github.com/ethereum/go-ethereum -v $PWD:/go/src/github/tokencard/contracts ethereum/client-go:alltools-v1.8.12 abigen"
echo ${ABIGEN}

${ABIGEN} --abi ./build/wallet/Wallet.abi         --bin ./build/wallet/Wallet.bin         --pkg bindings --type=Wallet     --out ./pkg/bindings/wallet.go && echo "success 6"
${ABIGEN} --abi ./build/card/Card.abi             --bin ./build/card/Card.bin             --pkg bindings --type=Card       --out ./pkg/bindings/card.go && echo "success 7"
${ABIGEN} --abi ./build/oracle/Oracle.abi         --bin ./build/oracle/Oracle.bin         --pkg bindings --type=Oracle     --out ./pkg/bindings/oracle.go && echo "success 8"
${ABIGEN} --abi ./build/controller/Controller.abi --bin ./build/controller/Controller.bin --pkg bindings --type=Controller --out ./pkg/bindings/controller.go && echo "success 9"
${ABIGEN} --abi ./build/token/Token.abi           --bin ./build/token/Token.bin           --pkg bindings --type=Token      --out ./pkg/bindings/token.go && echo "success 10"
