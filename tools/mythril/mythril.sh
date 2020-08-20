#!/bin/bash

if [ ! -d "contracts" ]; then
	echo "error: script needs to be run from project root './tools/mythril/mythril.sh'"
	exit 1
fi

docker run --rm -v "$PWD":/contracts -it --workdir=/contracts/contracts --entrypoint=sh mythril/myth@sha256:48a0bbe632eca0312237cae70395e687c748f1cf2585494a305fee26fb685fc6 -c '
myth analyze --solv=0.6.12 controller.sol --execution-timeout=800 &&
myth analyze --solv=0.6.12 holder.sol --execution-timeout=800 &&
myth analyze --solv=0.6.12 licence.sol --execution-timeout=800 &&
myth analyze --solv=0.6.12 tokenWhitelist.sol --execution-timeout=800 &&
myth analyze --solv=0.6.12 walletDeployer.sol --execution-timeout=800 &&
myth analyze --solv=0.6.12 wallet.sol --execution-timeout=800 &&
myth analyze --solv=0.6.12 oracle.sol --execution-timeout=800'
