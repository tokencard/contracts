#!/bin/bash

if [ ! -d "contracts" ]; then 
	echo "error: script needs to be run from project root './tools/slither/slither.sh'"
	exit 1
fi

docker run --rm -v "$PWD":/contracts -it --workdir=/contracts/tools/slither trailofbits/eth-security-toolbox@sha256:e5e2ebbffcc4c3a334063ba3871ec25a4e1dc4915d77166aef1aa265e0a5978f -c 'solc-select 0.6.6 && slither --solc-args="--optimize" --triage-mode ../../contracts'
