#!/bin/bash

docker run --rm -v "$PWD":/contracts -it --workdir=/contracts/slither trailofbits/eth-security-toolbox@sha256:3fb96e2d9de772f5e97f1c3c650c8a3d28660f8a64a60b76269da1ac19b86a28 -c 'solc-select 0.5.15 && slither --triage-mode ../contracts'
