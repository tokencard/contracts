#!/bin/bash

set -e -o pipefail

SOLC="docker run --rm -u `id -u` -v $PWD:/solidity --workdir /solidity/contracts ethereum/solc:0.4.24"

# Compile solidity contracts and hex prefix bytecode string.
${SOLC} --overwrite  --bin --abi wallet.sol         -o /solidity/build/wallet     --combined-json bin-runtime,srcmap-runtime,ast,srcmap,bin && bin=$(awk '{print "0x" $0}' ./build/wallet/Wallet.bin)         && echo $bin > ./build/wallet/Wallet.bin         && echo "Compiled wallet."     || exit 1
${SOLC} --overwrite  --bin --abi token.sol          -o /solidity/build/token      --combined-json bin-runtime,srcmap-runtime,ast,srcmap,bin && bin=$(awk '{print "0x" $0}' ./build/token/Token.bin)           && echo $bin > ./build/token/Token.bin           && echo "Compiled token."      || exit 1
${SOLC} --overwrite  --bin --abi oracle.sol         -o /solidity/build/oracle     --combined-json bin-runtime,srcmap-runtime,ast,srcmap,bin && bin=$(awk '{print "0x" $0}' ./build/oracle/Oracle.bin)         && echo $bin > ./build/oracle/Oracle.bin         && echo "Compiled oracle."     || exit 1
${SOLC} --overwrite  --bin --abi old/card.sol       -o /solidity/build/card       --combined-json bin-runtime,srcmap-runtime,ast,srcmap,bin && bin=$(awk '{print "0x" $0}' ./build/card/Card.bin)             && echo $bin > ./build/card/Card.bin             && echo "Compiled card."       || exit 1
${SOLC} --overwrite  --bin --abi old/controller.sol -o /solidity/build/controller --combined-json bin-runtime,srcmap-runtime,ast,srcmap,bin && bin=$(awk '{print "0x" $0}' ./build/controller/Controller.bin) && echo $bin > ./build/controller/Controller.bin && echo "Compiled controller." || exit 1

GE_PATH=${PWD}/vendor/github.com/ethereum/go-ethereum
if [ ! -d ${GE_PATH} ]
then
  GE_PATH=${GOPATH:-$HOME/go}/src/github.com/ethereum/go-ethereum
fi

# Generate Go bindings from solidity contracts.
ABIGEN="docker run --rm -u `id -u` --workdir /go/src/github/tokencard/contracts -e GOPATH=/go -v $GE_PATH:/go/src/github.com/ethereum/go-ethereum -v $PWD:/go/src/github/tokencard/contracts ethereum/client-go:alltools-v1.8.15 abigen"

${ABIGEN} --abi ./build/wallet/Wallet.abi         --bin ./build/wallet/Wallet.bin         --pkg bindings --type=Wallet     --out ./pkg/bindings/wallet.go     && echo "Generated wallet bindings."     || exit 1
${ABIGEN} --abi ./build/card/Card.abi             --bin ./build/card/Card.bin             --pkg bindings --type=Card       --out ./pkg/bindings/card.go       && echo "Generated card bindings."       || exit 1
${ABIGEN} --abi ./build/oracle/Oracle.abi         --bin ./build/oracle/Oracle.bin         --pkg bindings --type=Oracle     --out ./pkg/bindings/oracle.go     && echo "Generated oracle bindings."     || exit 1
${ABIGEN} --abi ./build/controller/Controller.abi --bin ./build/controller/Controller.bin --pkg bindings --type=Controller --out ./pkg/bindings/controller.go && echo "Generated controller bindings." || exit 1
${ABIGEN} --abi ./build/token/Token.abi           --bin ./build/token/Token.bin           --pkg bindings --type=Token      --out ./pkg/bindings/token.go      && echo "Generated token bindings."      || exit 1
