#!/bin/bash

if [ ! -d "contracts" ]; then 
	echo "error: script needs to be run from project root './tools/prettier/format.sh'"
	exit 1
fi

docker run --rm -v "$PWD":/contracts -it --workdir=/contracts --entrypoint=sh node@sha256:d2734cb9a3b16e1103b27bca3f9db410cf6834a11302d1b5304c1d8627fb9732 -c 'npm install --no-package-lock --no-save prettier@2.0.2 prettier-plugin-solidity@1.0.0-alpha.54 && npx prettier --write --plugin=prettier-plugin-solidity {contracts,contracts/mocks,contracts/internals,echidna/*}/*.sol'
