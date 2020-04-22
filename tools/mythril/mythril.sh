#!/bin/bash

if [ ! -d "contracts" ]; then
	echo "error: script needs to be run from project root './tools/mythril/mythril.sh'"
	exit 1
fi

docker run --rm -v "$PWD":/contracts -it --workdir=/contracts/contracts --entrypoint=sh mythril/myth@sha256:75605a2a7e848e416f471a67159d6812feaa45e03ddd352aac588e505d68e5c5 -c '
myth analyze --solv=0.6.4 controller.sol --execution-timeout=800 &&
myth analyze --solv=0.6.4 holder.sol --execution-timeout=800 &&
myth analyze --solv=0.6.4 licence.sol --execution-timeout=800 &&
myth analyze --solv=0.6.4 tokenWhitelist.sol --execution-timeout=800 &&
myth analyze --solv=0.6.4 walletDeployer.sol --execution-timeout=800 &&
myth analyze --solv=0.6.4 wallet.sol --execution-timeout=800 &&
myth analyze --solv=0.6.4 oracle.sol --execution-timeout=800'
