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

# Compile solidity contracts.
solc --overwrite --bin --abi wallet.sol     -o ./build/wallet
solc --overwrite --bin --abi controller.sol -o ./build/controller
solc --overwrite --bin --abi token.sol      -o ./build/token
solc --overwrite --bin --abi card.sol       -o ./build/card


# Generate Go bindings from solidity contracts.
confirm
abigen --abi ./build/wallet/Wallet.abi         --bin ./build/wallet/Wallet.bin         --pkg bindings --type=Wallet     --out ./pkg/bindings/wallet.go
abigen --abi ./build/card/Card.abi             --bin ./build/card/Card.bin             --pkg bindings --type=Card       --out ./pkg/bindings/card.go
abigen --abi ./build/controller/Controller.abi --bin ./build/controller/Controller.bin --pkg bindings --type=Controller --out ./pkg/bindings/controller.go
