#!/bin/bash

docker run --rm -v "$PWD":/contracts -it --workdir=/contracts --entrypoint=sh node@sha256:d2734cb9a3b16e1103b27bca3f9db410cf6834a11302d1b5304c1d8627fb9732 -c 'npm install --no-package-lock --no-save prettier@2.0.4 prettier-plugin-solidity@1.0.0-alpha.48 && npx prettier --write --plugin=prettier-plugin-solidity {contracts,contracts/mocks,contracts/internals,echidna/*}/*.sol'
