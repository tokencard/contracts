#!/bin/bash

if [ ! -d "contracts" ]; then 
	echo "error: script needs to be run from project root './tools/echidna/echidna.sh'"
	exit 1
fi

if [ ! -d "crytic-export" ]; then
	echo "error: contracts need to be flattened before running echidna './tools/slither/flatten.sh'"
	exit 1
fi

<<<<<<< HEAD
docker run --rm -v "$PWD":/contracts -it --workdir=/contracts trailofbits/eth-security-toolbox@sha256:6376ca0f1e01cfac40499650e3b5c3c430f7c6fee73fcd2ea71aad4d0fa0038b -c 'solc-select 0.6.11 &&
=======
docker run --rm -v "$PWD":/contracts -it --workdir=/contracts trailofbits/eth-security-toolbox@sha256:e5e2ebbffcc4c3a334063ba3871ec25a4e1dc4915d77166aef1aa265e0a5978f -c 'solc-select 0.6.4 &&
>>>>>>> 39b44442... Update tools to latest builds and use --optimize flag in slither
echidna-test tools/echidna/wallet/wallet.sol --config=tools/echidna/wallet/wallet.yaml --contract=TEST &&
echidna-test tools/echidna/wallet/addressWhitelist.sol --config=tools/echidna/wallet/addressWhitelist.yaml --contract=TEST &&
echidna-test tools/echidna/wallet/spendLimit.sol --config=tools/echidna/wallet/spendLimit.yaml --contract=TEST &&
echidna-test tools/echidna/wallet/gasTopUpLimit.sol --config=tools/echidna/wallet/gasTopUpLimit.yaml --contract=TEST &&
echidna-test tools/echidna/wallet/loadLimit.sol --config=tools/echidna/wallet/loadLimit.yaml --contract=TEST &&
echidna-test tools/echidna/controller/controller.sol --config=tools/echidna/controller/controller.yaml --contract=TEST &&
echidna-test tools/echidna/internals/ownable.sol --config=tools/echidna/internals/ownable.yaml --contract=TEST'
