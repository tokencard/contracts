#!/bin/bash

if [ ! -d "contracts" ]; then
	echo "error: script needs to be run from project root './tools/mythril/mythril.sh'"
	exit 1
fi

docker run --rm -v "$PWD":/contracts -it --workdir=/contracts/contracts --entrypoint=sh mythril/myth:latest -c '
myth analyze --solv=0.6.10 controller.sol --execution-timeout=800 &&
myth analyze --solv=0.6.10 holder.sol --execution-timeout=800 &&
myth analyze --solv=0.6.10 licence.sol --execution-timeout=800 &&
myth analyze --solv=0.6.10 tokenWhitelist.sol --execution-timeout=800 &&
myth analyze --solv=0.6.10 walletDeployer.sol --execution-timeout=800 &&
myth analyze --solv=0.6.10 wallet.sol --execution-timeout=800 &&
myth analyze --solv=0.6.10 oracle.sol --execution-timeout=800'
