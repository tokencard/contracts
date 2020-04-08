#!/bin/bash

if [ ! -d "contracts" ]; then 
	echo "error: script needs to be run from project root './tools/echidna/echidna.sh'"
	exit 1
fi

if [ ! -d "crytic-export" ]; then
	echo "error: contracts need to be flattened before running echidna './tools/slither/flatten.sh'"
	exit 1
fi

docker run --rm -v "$PWD":/contracts -it --workdir=/contracts/tools/echidna trailofbits/eth-security-toolbox@sha256:3fb96e2d9de772f5e97f1c3c650c8a3d28660f8a64a60b76269da1ac19b86a28 -c 'solc-select 0.5.17 &&
echidna-test wallet/wallet.sol --config=wallet/wallet.yaml --contract=TEST &&
echidna-test wallet/addressWhitelist.sol --config=wallet/addressWhitelist.yaml --contract=TEST &&
echidna-test wallet/spendLimit.sol --config=wallet/spendLimit.yaml --contract=TEST &&
echidna-test wallet/gasTopUpLimit.sol --config=wallet/gasTopUpLimit.yaml --contract=TEST &&
echidna-test wallet/loadLimit.sol --config=wallet/loadLimit.yaml --contract=TEST &&
echidna-test controller/controller.sol --config=controller/controller.yaml --contract=TEST &&
echidna-test internals/ownable.sol --config=internals/ownable.yaml --contract=TEST'
