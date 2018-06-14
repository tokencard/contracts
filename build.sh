#!/usr/bin/env bash

function confirm {
    read -r -p "Would you like to generate Go bindings? [y/N]" response
    case "$response" in
        [yY][eE][sS]|[yY])
            echo "Generating Go bindings ..."
            ;;
        *)
            echo "Skipping Go bindings ..."
            exit 1
            ;;
    esac
}

# Compile solidity contracts and hex prefix bytecode string.
solc --overwrite --bin --abi wallet.sol     -o ./build/wallet     && bin=$(awk '{print "0x" $0}' ./build/wallet/TopUpWallet.bin)         && echo $bin > ./build/wallet/TopUpWallet.bin
solc --overwrite --bin --abi controller.sol -o ./build/controller && bin=$(awk '{print "0x" $0}' ./build/controller/Controller.bin) && echo $bin > ./build/controller/Controller.bin
solc --overwrite --bin --abi token.sol      -o ./build/token      && bin=$(awk '{print "0x" $0}' ./build/token/Token.bin)           && echo $bin > ./build/token/Token.bin
solc --overwrite --bin --abi card.sol       -o ./build/card       && bin=$(awk '{print "0x" $0}' ./build/card/Card.bin)             && echo $bin > ./build/card/Card.bin


# Generate Go bindings from solidity contracts.
confirm
abigen --abi ./build/wallet/TopUpWallet.abi    --bin ./build/wallet/TopUpWallet.bin    --pkg bindings --type=Wallet     --out ./pkg/bindings/wallet.go
abigen --abi ./build/card/Card.abi             --bin ./build/card/Card.bin             --pkg bindings --type=Card       --out ./pkg/bindings/card.go
abigen --abi ./build/controller/Controller.abi --bin ./build/controller/Controller.bin --pkg bindings --type=Controller --out ./pkg/bindings/controller.go
