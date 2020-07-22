#!/bin/bash

if [ ! -d "contracts" ]; then 
	echo "error: script needs to be run from project root './tools/echidna/flatten.sh'"
	exit 1
fi

docker run --rm -v "$PWD":/contracts -it --workdir=/contracts trailofbits/eth-security-toolbox@sha256:6376ca0f1e01cfac40499650e3b5c3c430f7c6fee73fcd2ea71aad4d0fa0038b -c 'solc-select 0.6.11 && find contracts -type f -name "controller.sol" -exec slither-flat --convert-external {} \;'
