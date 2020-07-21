#!/bin/bash

if [ ! -d "contracts" ]; then 
	echo "error: script needs to be run from project root './tools/echidna/echidna.sh'"
	exit 1
fi

if [ ! -d "crytic-export" ]; then
	echo "error: contracts need to be flattened before running echidna './tools/slither/flatten.sh'"
	exit 1
fi

docker run --rm -v "$PWD":/contracts -it --workdir=/contracts trailofbits/eth-security-toolbox@sha256:3fb96e2d9de772f5e97f1c3c650c8a3d28660f8a64a60b76269da1ac19b86a28 -c 'solc-select 0.6.11 &&
echidna-test tools/echidna/wallet/wallet.sol --config=tools/echidna/wallet/wallet.yaml --contract=TEST &&
echidna-test tools/echidna/wallet/addressWhitelist.sol --config=tools/echidna/wallet/addressWhitelist.yaml --contract=TEST &&
echidna-test tools/echidna/wallet/spendLimit.sol --config=tools/echidna/wallet/spendLimit.yaml --contract=TEST &&
echidna-test tools/echidna/wallet/gasTopUpLimit.sol --config=tools/echidna/wallet/gasTopUpLimit.yaml --contract=TEST &&
echidna-test tools/echidna/wallet/loadLimit.sol --config=tools/echidna/wallet/loadLimit.yaml --contract=TEST &&
echidna-test tools/echidna/controller/controller.sol --config=tools/echidna/controller/controller.yaml --contract=TEST &&
echidna-test tools/echidna/internals/ownable.sol --config=tools/echidna/internals/ownable.yaml --contract=TEST'
