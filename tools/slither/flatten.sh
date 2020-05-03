#!/bin/bash

if [ ! -d "contracts" ]; then 
	echo "error: script needs to be run from project root './tools/echidna/flatten.sh'"
	exit 1
fi

docker run --rm -v "$PWD":/contracts -it --workdir=/contracts trailofbits/eth-security-toolbox@sha256:e5e2ebbffcc4c3a334063ba3871ec25a4e1dc4915d77166aef1aa265e0a5978f -c 'solc-select 0.6.4 && find contracts -type f -name "*.sol" -exec slither-flat --convert-external {} \;'
