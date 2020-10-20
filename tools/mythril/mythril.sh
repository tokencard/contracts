#!/bin/bash

if [ ! -d "contracts" ]; then
	echo "error: script needs to be run from project root './tools/mythril/mythril.sh'"
	exit 1
fi

<<<<<<< HEAD
docker run --rm -v "$PWD":/contracts -it --workdir=/contracts/contracts --entrypoint=sh mythril/myth@sha256:48a0bbe632eca0312237cae70395e687c748f1cf2585494a305fee26fb685fc6 -c '
myth analyze --solv=0.6.12 controller.sol --execution-timeout=800 &&
myth analyze --solv=0.6.12 holder.sol --execution-timeout=800 &&
myth analyze --solv=0.6.12 licence.sol --execution-timeout=800 &&
myth analyze --solv=0.6.12 tokenWhitelist.sol --execution-timeout=800 &&
myth analyze --solv=0.6.12 walletDeployer.sol --execution-timeout=800 &&
myth analyze --solv=0.6.12 wallet.sol --execution-timeout=800 &&
myth analyze --solv=0.6.12 oracle.sol --execution-timeout=800'
=======
docker run --rm -v "$PWD":/contracts -it --workdir=/contracts/contracts --entrypoint=sh mythril/myth:latest -c '
myth analyze --solv=0.6.4 controller.sol --execution-timeout=800 &&
myth analyze --solv=0.6.4 holder.sol --execution-timeout=800 &&
myth analyze --solv=0.6.4 licence.sol --execution-timeout=800 &&
myth analyze --solv=0.6.4 tokenWhitelist.sol --execution-timeout=800 &&
myth analyze --solv=0.6.4 walletDeployer.sol --execution-timeout=800 &&
myth analyze --solv=0.6.4 wallet.sol --execution-timeout=800 &&
myth analyze --solv=0.6.4 oracle.sol --execution-timeout=800'
>>>>>>> 39b44442... Update tools to latest builds and use --optimize flag in slither
