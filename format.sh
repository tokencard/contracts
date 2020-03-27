#!/bin/bash

docker run --rm -v "$PWD":/contracts -it --workdir=/contracts quay.io/token/go-build:v0.3.7 \
    apk add -U nodejs npm &&
    npm install --no-package-lock prettier@2.0.2 prettier-plugin-solidity@1.0.0-alpha.47 &&
    npx prettier --write --plugin=prettier-plugin-solidity {contracts,contracts/mocks,contracts/internals}/*.sol
