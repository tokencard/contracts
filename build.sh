#!/usr/bin/env bash

set -e

SOLC="docker run -v$PWD:/solidity ethereum/solc:0.4.24"

# Compile solidity contracts and hex prefix bytecode string.
$SOLC --overwrite  --bin --abi wallet.sol     -o ./build/wallet     && bin=$(awk '{print "0x" $0}' ./build/wallet/Wallet.bin)         && echo $bin > ./build/wallet/Wallet.bin
$SOLC --overwrite  --bin --abi controller.sol -o ./build/controller && bin=$(awk '{print "0x" $0}' ./build/controller/Controller.bin) && echo $bin > ./build/controller/Controller.bin
$SOLC --overwrite  --bin --abi token.sol      -o ./build/token      && bin=$(awk '{print "0x" $0}' ./build/token/Token.bin)           && echo $bin > ./build/token/Token.bin
$SOLC --overwrite  --bin --abi card.sol       -o ./build/card       && bin=$(awk '{print "0x" $0}' ./build/card/Card.bin)             && echo $bin > ./build/card/Card.bin
$SOLC --overwrite  --bin --abi oracle.sol     -o ./build/oracle     && bin=$(awk '{print "0x" $0}' ./build/oracle/Oracle.bin)         && echo $bin > ./build/oracle/Oracle.bin

# create combined-json for code coverage
$SOLC --overwrite --combined-json bin-runtime,srcmap-runtime,ast,srcmap,bin -o ./build wallet.sol controller.sol token.sol card.sol oracle.sol

GE_PATH="$PWD"/vendor/github.com/ethereum/go-ethereum
if [ ! -d "$GE_PATH" ]
then
  GE_PATH=${GOPATH:-$HOME/go}/src/github.com/ethereum/go-ethereum
fi

# Generate Go bindings from solidity contracts.
ABIGEN="docker run --workdir /go/src/github/tokencard/contracts -e GOPATH=/go -v "$GE_PATH":/go/src/github.com/ethereum/go-ethereum -v "$PWD":/go/src/github/tokencard/contracts ethereum/client-go:alltools-v1.8.12 abigen"

$ABIGEN --abi ./build/wallet/Wallet.abi         --bin ./build/wallet/Wallet.bin         --pkg bindings --type=Wallet     --out ./pkg/bindings/wallet.go
$ABIGEN --abi ./build/card/Card.abi             --bin ./build/card/Card.bin             --pkg bindings --type=Card       --out ./pkg/bindings/card.go
$ABIGEN --abi ./build/oracle/Oracle.abi         --bin ./build/oracle/Oracle.bin         --pkg bindings --type=Oracle     --out ./pkg/bindings/oracle.go
$ABIGEN --abi ./build/controller/Controller.abi --bin ./build/controller/Controller.bin --pkg bindings --type=Controller --out ./pkg/bindings/controller.go
$ABIGEN --abi ./build/token/Token.abi           --bin ./build/token/Token.bin           --pkg bindings --type=Token      --out ./pkg/bindings/token.go
